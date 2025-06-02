-- name: CreatePost :exec
INSERT INTO posts (
    id,
    created_at, 
    updated_at,
    title,
    description,
    published_at,
    url,
    feed_id
) 
VALUES (?,?,?,?,?,?,?,?);

-- name: GetPost :one
SELECT * FROM users WHERE id = ?;



