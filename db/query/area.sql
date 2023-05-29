-- name: ListAreas :many
SELECT * FROM area
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: GetArea :one
SELECT * FROM area
WHERE id = $1;

-- name: CreateArea :one
INSERT INTO area
(id, name, description)
VALUES ($1, $2, $3) RETURNING *;

-- name: UpdateArea :exec
UPDATE area
    set name = $2, description = $3
WHERE id = $1;

-- name: DeleteArea :exec
DELETE FROM area
WHERE id = $1;
