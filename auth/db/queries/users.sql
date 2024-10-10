-- name: InserUser :one
INSERT INTO users 
  ("name", "email", "password", "created_at", "updated_at") VALUES
  ( $1, $2, $3, now(), now() )
RETURNING "id";
