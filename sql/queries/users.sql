-- name: CreateUser :exec
INSERT INTO users (id,created_at,updated_at,name) 
VALUES (?,?,?,?);

-- name: GetUser :one
SELECT * FROM users WHERE id = ?;
