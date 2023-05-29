-- name: ListCompanies :many
SELECT * FROM company
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: GetCompany :one
SELECT * FROM company
WHERE id = $1;

-- name: CreateCompany :one
INSERT INTO company
(id, name, description)
VALUES ($1, $2, $3) RETURNING *;

-- name: UpdateCompany :exec
UPDATE company
    set name = $2, description = $3
WHERE id = $1;

-- name: DeleteCompany :exec
DELETE FROM company
WHERE id = $1;
