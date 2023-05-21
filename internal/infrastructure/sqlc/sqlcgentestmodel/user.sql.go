// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.14.0
// source: user.sql

package sqlcgentestmodel

import (
	"context"
)

const testCreateUser = `-- name: TestCreateUser :exec
INSERT INTO user (id)
VALUES
    ('01H0YHVK6NREDFDEK81YT1ERKR'),
    ('01H0YHVK6N4DJVRX7HXS9KSF4V'),
    ('01H0YHVK6MK92QZ1SZD2HAXEZ2')
`

func (q *Queries) TestCreateUser(ctx context.Context) error {
	_, err := q.db.ExecContext(ctx, testCreateUser)
	return err
}

const testDeleteUser = `-- name: TestDeleteUser :exec
DELETE FROM user
WHERE id IN ('01H0YHVK6NREDFDEK81YT1ERKR', '01H0YHVK6N4DJVRX7HXS9KSF4V', '01H0YHVK6MK92QZ1SZD2HAXEZ2')
`

func (q *Queries) TestDeleteUser(ctx context.Context) error {
	_, err := q.db.ExecContext(ctx, testDeleteUser)
	return err
}