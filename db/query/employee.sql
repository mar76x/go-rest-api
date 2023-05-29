-- name: ListEmployees :many
SELECT * FROM employee
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: GetEmployee :one
SELECT * FROM employee
WHERE id = $1;

-- name: CreateEmployee :one
INSERT INTO employee
(id, number, name, surname, birthdate, dni, cuil, marital_status)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING *;

-- name: UpdateEmployee :exec
UPDATE employee
    set number = $2, name = $3, surname = $4, birthdate = $5, dni = $6, cuil = $7, marital_status = $8
WHERE id = $1;

-- name: DeleteEmployee :exec
DELETE FROM employee
WHERE id = $1;
