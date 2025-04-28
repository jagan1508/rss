-- +goose Up

ALTER TABLE users ADD COLUMN api_key VARCHAR(64) UNIQUE;
UPDATE users SET api_key = SHA2(UUID(), 256);
ALTER TABLE users MODIFY COLUMN api_key VARCHAR(64) NOT NULL;


-- +goose Down
ALTER TABLE users DROP COLUMN api_key;

