// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: users.sql

package db

import (
	"context"

	"github.com/google/uuid"
)

const findUserByEmail = `-- name: FindUserByEmail :one
SELECT id, name, email, password, created_at, updated_at FROM users WHERE email = $1
`

func (q *Queries) FindUserByEmail(ctx context.Context, email string) (User, error) {
	row := q.db.QueryRow(ctx, findUserByEmail, email)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Email,
		&i.Password,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const findUserById = `-- name: FindUserById :one
SELECT id, name, email, password, created_at, updated_at FROM users WHERE id = $1
`

func (q *Queries) FindUserById(ctx context.Context, id uuid.UUID) (User, error) {
	row := q.db.QueryRow(ctx, findUserById, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Email,
		&i.Password,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const insertUser = `-- name: InsertUser :one
INSERT INTO users 
  ("name", "email", "password", "created_at", "updated_at") VALUES
  ( $1, $2, $3, now(), now() )
RETURNING "id"
`

type InsertUserParams struct {
	Name     string
	Email    string
	Password string
}

func (q *Queries) InsertUser(ctx context.Context, arg InsertUserParams) (uuid.UUID, error) {
	row := q.db.QueryRow(ctx, insertUser, arg.Name, arg.Email, arg.Password)
	var id uuid.UUID
	err := row.Scan(&id)
	return id, err
}
