-- Write your migrate up statements here
CREATE TABLE IF NOT EXISTS clients (
    "id" uuid PRIMARY KEY DEFAULT gen_random_uuid (),
    "name" VARCHAR(255) NOT NULL,
    "brand_name" VARCHAR(255) NOT NULL,
    "logo_url" VARCHAR NULL,
    "created_at" TIMESTAMP NOT NULL DEFAULT now(),
    "updated_at" TIMESTAMP NOT NULL DEFAULT now()
);
---- create above / drop below ----

DROP TABLE IF EXISTS clients;

-- Write your migrate down statements here. If this migration is irreversible
-- Then delete the separator line above.
