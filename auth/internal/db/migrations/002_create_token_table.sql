-- Write your migrate up statements here
CREATE TABLE IF NOT EXISTS user_token (
    "id" uuid PRIMARY KEY DEFAULT gen_random_uuid (),
    "user_id" uuid NOT NULL REFERENCES users ("id") ON DELETE CASCADE,
    "token" TEXT UNIQUE NOT NULL,
    "created_at" TIMESTAMP NOT NULL DEFAULT now(),
    "expires_at" TIMESTAMP NOT NULL
);

CREATE INDEX idx_user_token_user_id ON user_token ("user_id");

CREATE INDEX idx_user_token_token ON user_token ("token");

---- create above / drop below ----
DROP TABLE IF EXISTS user_token;

-- Write your migrate down statements here. If this migration is irreversible
-- Then delete the separator line above.
