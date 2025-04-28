-- +goose Up
CREATE TABLE feeds_follows (
    id VARCHAR(40) PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    feed_id VARCHAR(40) NOT NULL,
    user_id VARCHAR(40) NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE ,
    FOREIGN KEY (feed_id) REFERENCES feeds(id) ON DELETE CASCADE ,
    UNIQUE (feed_id,user_id)
);

-- +goose Down
DROP TABLE feeds_follows;