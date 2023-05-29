-- name: ListDepartments :many
SELECT * FROM department
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: GetDepartment :one
SELECT * FROM department
WHERE id = $1;

-- name: CreateDepartment :one
INSERT INTO department 
(id, name, description)
VALUES ($1, $2, $3) RETURNING *;

-- name: UpdateDepartment :exec
UPDATE department
    set name = $2, description = $3
WHERE id = $1;

-- name: DeleteDepartment :exec
DELETE FROM department
WHERE id = $1;
