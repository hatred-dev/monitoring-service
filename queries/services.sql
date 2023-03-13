-- name: GetServices :many
SELECT *
FROM services
WHERE project_id = $1;

-- name: CreateService :exec
INSERT INTO services (project_id, service_name, url)
VALUES ($1, $2, $3)
RETURNING *;

-- name: UpdateService :exec
UPDATE services
SET service_name = $2,
    url          = $3
WHERE project_id = $1;

-- name: DeleteService :exec
DELETE
FROM services
WHERE project_id = $1
  AND service_name = $2;

-- name: GetService :one
SELECT *
FROM services
WHERE project_id = $1
  AND service_name = $2
LIMIT 1;