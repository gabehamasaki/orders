package server

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/gabehamasaki/orders/auth/internal/db"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

// createToken generates a new JWT token for the given user ID or returns an existing valid token
func (s *Server) createToken(userID string, clientID string) (string, error) {
	// Set token expiration time to 24 hours from now
	expirationTime := time.Now().Add(time.Hour * 24)
	ext := expirationTime.Unix()
	iat := time.Now().Unix()

	// Check if there's an existing token for the user
	alreadyExists, err := s.DB.GetTokenIsExpiredByUserId(context.Background(), uuid.MustParse(userID))
	if err != nil {
		// If there's an error, try to find a valid token
		validToken, err := s.DB.FindTokenByUserId(context.Background(), uuid.MustParse(userID))
		if err == nil {
			// If a valid token is found, return it
			return validToken.Token, nil
		}
	} else {
		// If an expired token exists, delete it
		err = s.DB.DeletedTokenById(context.Background(), alreadyExists.ID)
		if err != nil {
			return "", fmt.Errorf("failed to delete expired token: %w", err)
		}
	}

	// Create a new token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": fmt.Sprintf("%s|%s", userID, clientID), // Subject (user ID)
		"iss": "orders-app",                           // Issuer
		"aud": "user",                                 // Audience
		"exp": ext,                                    // Expiration time
		"iat": iat,                                    // Issued at
	})

	// Sign the token with the server's secret key
	tokenString, err := token.SignedString(s.SecretKey)
	if err != nil {
		return "", fmt.Errorf("failed to sign token: %w", err)
	}

	// Insert the new token into the database
	_, err = s.DB.InsertToken(context.Background(), db.InsertTokenParams{
		UserID: uuid.MustParse(userID),
		Token:  tokenString,
		ExpiresAt: pgtype.Timestamp{
			Time:  expirationTime,
			Valid: true,
		},
	})
	if err != nil {
		return "", fmt.Errorf("failed to insert token into database: %w", err)
	}

	return tokenString, nil
}

type VerifyTokenResponse struct {
	UserID   string
	ClientID string
	Valid    bool
}

// verifyToken validates the given JWT token and returns the associated user ID
func (s *Server) verifyToken(ctx context.Context, tokenString string) (*VerifyTokenResponse, error) {
	var userID uuid.UUID
	var clientID uuid.UUID

	// Parse and validate the token
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		// Verify the signing method
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}

		// Extract the subject (user ID) from the token claims
		sub, err := t.Claims.GetSubject()
		if err != nil {
			return nil, fmt.Errorf("failed to get subject from token: %w", err)
		}

		IDS := strings.Split(sub, "|")

		// Parse the user ID
		userID, err = uuid.Parse(IDS[0])
		if err != nil {
			return nil, fmt.Errorf("invalid user ID in token: %w", err)
		}

		// Parse the client ID
		clientID, err = uuid.Parse(IDS[1])
		if err != nil {
			return nil, fmt.Errorf("invalid client ID in token: %w", err)
		}

		// Verify that the user exists in the database
		_, err = s.DB.FindUserById(ctx, db.FindUserByIdParams{
			ID: userID,
			ClientID: pgtype.UUID{
				Bytes: clientID,
				Valid: IDS[1] != "",
			},
		})
		if err != nil {
			return nil, fmt.Errorf("user not found in database: %w", err)
		}

		existToken, err := s.DB.FindTokenByUserId(ctx, userID)
		if err == nil {
			if existToken.Token != tokenString {
				return nil, fmt.Errorf("token not found in database")
			}
			return s.SecretKey, nil
		}

		// Return the secret key for token validation
		return s.SecretKey, nil
	})

	// Handle parsing errors
	if err != nil {
		return nil, fmt.Errorf("failed to parse token: %w", err)
	}

	// Check if the token is valid
	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	// Return the validated user ID
	return &VerifyTokenResponse{
		UserID:   userID.String(),
		ClientID: clientID.String(),
	}, nil
}
