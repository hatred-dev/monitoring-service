-- name: GetIPs :many
SELECT *
FROM ip_address
WHERE project_id = $1;

-- name: GetAllIPs :many
SELECT *
FROM ip_address;

-- name: CreateIP :one
INSERT INTO ip_address (project_id, ip)
VALUES ($1, $2)
RETURNING *;

-- name: UpdateIP :exec
UPDATE ip_address
SET ip = $2
WHERE id = $1;

