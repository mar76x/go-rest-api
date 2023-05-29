-- name: ListContract :many
SELECT * FROM contract
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: GetContract :one
SELECT * FROM contract
WHERE id = $1;

-- name: CreateContract :one
INSERT INTO contract
(id, type, start_date, employee_id, company_id, branch_id, area_id, department_id, role_id)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING *;

-- name: UpdateContract :exec
UPDATE contract
    set type = $2, start_date = $3, employee_id = $4, company_id = $5, branch_id = $6, area_id = $7, department_id = $8, role_id = $9
WHERE id = $1;

-- name: DeleteContract :exec
DELETE FROM contract
WHERE id = $1;
