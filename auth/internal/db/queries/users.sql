-- name: InsertUser :one
INSERT INTO users 
  ("name", "email", "password", "created_at", "updated_at") VALUES
  ( $1, $2, $3, now(), now() )
RETURNING "id";

-- name: FindUserByEmail :one
SELECT * FROM users WHERE email = $1;

-- name: FindUserById :one
SELECT * FROM users WHERE id = $1;


