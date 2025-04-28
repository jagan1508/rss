-- name: CreateFeedFollows :exec
INSERT INTO feeds_follows (id,created_at,updated_at,feed_id,user_id) 
VALUES (?,?,?,?,?);

-- name: GetFeedFollows :one
SELECT * FROM feeds_follows WHERE id = ?;

-- name: GetFeedsFollows :many
SELECT * FROM feeds_follows WHERE user_id =?;

-- name: DeleteFeedFollows :exec
DELETE FROM feeds_follows WHERE id =? AND user_id =?;

