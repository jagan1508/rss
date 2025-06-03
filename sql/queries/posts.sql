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

-- name: GetPostsForUser :many
SELECT posts.* FROM posts 
JOIN feeds_follows ON posts.feed_id =feeds_follows.feed_id
WHERE feeds_follows.user_id = ? ORDER BY posts.published_at DESC LIMIT ?;


