-- name: GetProject :one
SELECT *
FROM projects
WHERE project_name = $1
LIMIT 1;

-- name: ListProjects :many
SELECT *
FROM projects
ORDER BY project_name;

-- name: CreateProject :one
INSERT INTO projects (project_name, active)
VALUES ($1, $2)
RETURNING *;

-- name: DeleteProject :exec
DELETE
FROM projects
WHERE project_name = $1;

-- name: UpdateProject :exec
UPDATE projects
SET project_name = $2,
    active       = $3
WHERE project_name = $1;

-- name: ProjectSetActive :exec
UPDATE projects
SET active = $2
WHERE project_name = $1;