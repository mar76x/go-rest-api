-- name: ListPaychecks :many
SELECT * FROM paycheck
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: GetPaycheck :one
SELECT * FROM paycheck
WHERE id = $1;

-- name: CreatePaycheck :one
INSERT INTO paycheck
(type, filename, description, folder, path, read, signed, employee_id)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING *;

-- name: UpdatePaycheck :exec
UPDATE paycheck
    set type = $2, filename = $3, description = $4, folder = $5, path = $6, read = $7, signed = $8, employee_id = $9
WHERE id = $1;

-- name: DeletePaycheck :exec
DELETE FROM paycheck
WHERE id = $1;
