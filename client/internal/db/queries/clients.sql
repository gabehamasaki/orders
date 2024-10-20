-- name: InsertClient :one
INSERT INTO
    clients (
        "name",
        "brand_name",
        "logo_url"
    )
VALUES ($1, $2, $3) RETURNING "id";

-- name: UpdateClient :exec
UPDATE clients
SET
    "name" = $2,
    "brand_name" = $3,
    "logo_url" = $4
WHERE
    id = $1;

-- name: DeleteClient :exec
DELETE FROM clients WHERE id = $1;

-- name: GetClients :many
WITH client_count AS (
    SELECT COUNT(*) AS total FROM clients
)
SELECT c.id, c.name, c.brand_name, c.logo_url,
       cc.total,
       CEIL(cc.total::float / $1::int) AS total_pages
FROM clients c, client_count cc
ORDER BY c.created_at
LIMIT $1
OFFSET $2;

-- name: FindClientById :one
SELECT * FROM clients WHERE id = $1;
