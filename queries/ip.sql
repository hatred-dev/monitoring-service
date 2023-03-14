-- name: GetIPs :many
SELECT *
FROM ip_address
WHERE project_id = $1;

-- name: GetAllIPs :many
SELECT *
FROM ip_address;

-- name: CreateIP :one
INSERT INTO ip_address (project_id, ip)
VALUES ((SELECT p.id FROM projects p WHERE p.project_name = $1), $2)
RETURNING *;

-- name: IPExists :one
SELECT EXISTS(SELECT 1 from ip_address where ip = $1);


-- name: UpdateIP :exec
UPDATE ip_address
SET ip = $2
WHERE id = (SELECT i.id FROM ip_address i WHERE i.ip = $1);

