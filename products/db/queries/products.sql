-- name: InsertProduct :one
INSERT INTO
    products (
        name,
        description,
        price,
        image_url
    )
VALUES ($1, $2, $3, $4) RETURNING id;

-- name: GetProduct :one
SELECT * FROM products WHERE id = $1;
