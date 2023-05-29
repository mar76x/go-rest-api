-- name: ListRoles :many
SELECT * FROM role
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: GetRole :one
SELECT * FROM role
WHERE id = $1;

-- name: CreateRole :one
INSERT INTO role
(id, name, description)
VALUES ($1, $2, $3) RETURNING *;

-- name: UpdateRole :exec
UPDATE role
    set name = $2, description = $3
WHERE id = $1;

-- name: DeleteRole :exec
DELETE FROM role
WHERE id = $1;
