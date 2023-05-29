-- name: ListBranches :many
SELECT * FROM branch
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: GetBranch :one
SELECT * FROM branch
WHERE id = $1;

-- name: CreateBranch :one
INSERT INTO branch
(id, company_id, name, description)
VALUES ($1, $2, $3, $4) RETURNING *;

-- name: UpdateBranch :exec
UPDATE branch
    set name = $2, description = $3
WHERE id = $1;

-- name: DeleteBranch :exec
DELETE FROM branch
WHERE id = $1;
