-- +goose Up
CREATE TABLE posts (
    id VARCHAR(40) PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    title TEXT NOT NULL,
    description TEXT,
    published_at TIMESTAMP NOT NULL,
    url VARCHAR(500) UNIQUE NOT NULL,
    feed_id VARCHAR(40) NOT NULL,
    FOREIGN KEY (feed_id) REFERENCES feeds(id) ON DELETE CASCADE 
);

-- +goose Down
DROP TABLE posts;