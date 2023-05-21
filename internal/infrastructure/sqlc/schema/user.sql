-- name: GetUser :one
SELECT id
FROM user
WHERE id = ?
;