-- name: AddAccess :one
INSERT INTO accesses (api_key, refresh_token, role)
VALUES ($1, $2, $3)
RETURNING *;

-- name: ListAccesses :many
SELECT *
FROM accesses
ORDER BY created_at DESC;

-- name: GetAccess :one
SELECT * FROM accesses
WHERE id = $1 LIMIT 1;

-- name: EditAccess :one
UPDATE accesses
SET refresh_token = $2,
  updated_at = now()
WHERE id = $1
RETURNING *;

-- name: RemoveAccess :exec
DELETE FROM accesses
WHERE id = $1;
