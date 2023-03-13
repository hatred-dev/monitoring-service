-- name: GetProjectByName :one
SELECT *
FROM projects
WHERE project_name = $1
LIMIT 1;


-- name: GetProjects :many
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
WHERE id = $1;

-- name: ProjectSetState :exec
UPDATE projects
SET active = $2
WHERE project_name = $1;