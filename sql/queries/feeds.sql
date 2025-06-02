-- name: CreateFeed :exec
INSERT INTO feeds (id,created_at,updated_at,name,url,user_id) 
VALUES (?,?,?,?,?,?);

-- name: GetFeed :one
SELECT * FROM feeds WHERE id = ?;

-- name: GetFeeds :many
SELECT * FROM feeds;

-- name: GetNextFeedToFetch :many
SELECT * FROM feeds ORDER BY last_fetched_at ASC LIMIT ?;   

-- name: MarkFeedAsFetched :exec
UPDATE feeds SET last_fetched_at = NOW(),
updated_at = NOW()
WHERE id =?;

-- name: GetFeedAsFetched :one
SELECT * FROM feeds WHERE id =?;