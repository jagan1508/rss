-- name: CreateUser :exec
INSERT INTO users (id,created_at,updated_at,name,api_key) 
VALUES (?,?,?,?,SHA2(UUID(), 256));

-- name: GetUser :one
SELECT * FROM users WHERE id = ?;

-- name: GetUserByApiKey :one
SELECT * FROM users WHERE api_key = ?;

