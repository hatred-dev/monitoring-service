-- name: GetServices :many
SELECT s.id, p.project_name, s.service_name, s.url
from services s
         INNER JOIN projects p on p.id = s.project_id;

-- name: GetServicesByProjectName :many
SELECT *
FROM services
WHERE project_id = (SELECT project_id FROM projects WHERE project_name = $1);

-- name: CreateService :one
INSERT INTO services (project_id, service_name, url)
VALUES ((SELECT id FROM projects WHERE project_name = $1), $2, $3)
RETURNING *;

-- name: UpdateService :exec
UPDATE services
SET service_name = $2,
    url          = $3
WHERE project_id = (SELECT project_id FROM projects WHERE project_name = $1);

-- name: DeleteService :exec
DELETE
FROM services
WHERE project_id = (SELECT project_id FROM projects WHERE project_name = $1)
  AND service_name = $2;

-- name: GetService :one
SELECT *
FROM services
WHERE project_id = $1
  AND service_name = $2
LIMIT 1;

-- name: ServiceExists :one
SELECT EXISTS(SELECT 1
              FROM services s
              WHERE s.service_name = $2
                AND project_id = (SELECT id FROM projects p WHERE p.project_name = $1));