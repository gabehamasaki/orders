-- Write your migrate up statements here
CREATE TABLE IF NOT EXISTS products (
    "id" uuid PRIMARY KEY DEFAULT gen_random_uuid (),
    "name" VARCHAR NOT NULL,
    "description" TEXT NULL,
    "price" DECIMAL(10, 2) NOT NULL,
    "image_url" TEXT NULL,
    "created_at" TIMESTAMP NOT NULL DEFAULT now(),
    "updated_at" TIMESTAMP NOT NULL DEFAULT now()
);
---- create above / drop below ----
DROP TABLE IF EXISTS products;

-- Write your migrate down statements here. If this migration is irreversible
-- Then delete the separator line above.
