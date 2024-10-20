-- name: InsertProduct :one
INSERT INTO
    products (
        name,
        description,
        price,
        image_url,
        client_id
    )
VALUES ($1, $2, $3, $4, $5) RETURNING id;

-- name: GetProduct :one
SELECT * FROM products WHERE id = $1 AND client_id = $2;

-- name: GetProducts :many
WITH product_count AS (
    SELECT COUNT(*) AS total FROM products
)
SELECT p.id, p.name, p.description, p.price, p.image_url, p.created_at, p.updated_at,
       pc.total,
       CEIL(pc.total::float / $1::int) AS total_pages
FROM products p, product_count pc
WHERE p.client_id = $3
ORDER BY p.created_at
LIMIT $1
OFFSET $2;
