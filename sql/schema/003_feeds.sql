-- +goose Up
CREATE TABLE feeds (
    id VARCHAR(40) PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    name TEXT NOT NULL,
    url VARCHAR(500) UNIQUE NOT NULL,
    user_id VARCHAR(40) NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE 
);

-- +goose Down
DROP TABLE feeds;