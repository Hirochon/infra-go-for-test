// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.14.0
// source: user.sql

package sqlcgenmodel

import (
	"context"
)

const getUser = `-- name: GetUser :one
SELECT id
FROM user
WHERE id = ?
`

func (q *Queries) GetUser(ctx context.Context, id string) (string, error) {
	row := q.db.QueryRowContext(ctx, getUser, id)
	err := row.Scan(&id)
	return id, err
}
