// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.29.0
// source: feed_follows.sql

package database

import (
	"context"
	"time"
)

const createFeedFollows = `-- name: CreateFeedFollows :exec
INSERT INTO feeds_follows (id,created_at,updated_at,feed_id,user_id) 
VALUES (?,?,?,?,?)
`

type CreateFeedFollowsParams struct {
	ID        string
	CreatedAt time.Time
	UpdatedAt time.Time
	FeedID    string
	UserID    string
}

func (q *Queries) CreateFeedFollows(ctx context.Context, arg CreateFeedFollowsParams) error {
	_, err := q.db.ExecContext(ctx, createFeedFollows,
		arg.ID,
		arg.CreatedAt,
		arg.UpdatedAt,
		arg.FeedID,
		arg.UserID,
	)
	return err
}

const deleteFeedFollows = `-- name: DeleteFeedFollows :exec
DELETE FROM feeds_follows WHERE id =? AND user_id =?
`

type DeleteFeedFollowsParams struct {
	ID     string
	UserID string
}

func (q *Queries) DeleteFeedFollows(ctx context.Context, arg DeleteFeedFollowsParams) error {
	_, err := q.db.ExecContext(ctx, deleteFeedFollows, arg.ID, arg.UserID)
	return err
}

const getFeedFollows = `-- name: GetFeedFollows :one
SELECT id, created_at, updated_at, feed_id, user_id FROM feeds_follows WHERE id = ?
`

func (q *Queries) GetFeedFollows(ctx context.Context, id string) (FeedsFollow, error) {
	row := q.db.QueryRowContext(ctx, getFeedFollows, id)
	var i FeedsFollow
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.FeedID,
		&i.UserID,
	)
	return i, err
}

const getFeedsFollows = `-- name: GetFeedsFollows :many
SELECT id, created_at, updated_at, feed_id, user_id FROM feeds_follows WHERE user_id =?
`

func (q *Queries) GetFeedsFollows(ctx context.Context, userID string) ([]FeedsFollow, error) {
	rows, err := q.db.QueryContext(ctx, getFeedsFollows, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []FeedsFollow
	for rows.Next() {
		var i FeedsFollow
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.FeedID,
			&i.UserID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
