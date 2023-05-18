-- Company queries
-- name: ListCompanies :many
SELECT * FROM companies;

-- name: GetCompany :one
SELECT * FROM companies 
WHERE id = $1;

-- name: CreateCompany :one
INSERT INTO companies 
(id, name, description)
VALUES ($1, $2, $3) RETURNING *;

-- name: UpdateCompany :exec
UPDATE companies
    set name = $2, description = $3
WHERE id = $1;

-- name: DeleteCompany :exec
DELETE FROM companies
WHERE id = $1;

--
--

-- Branch queries
-- name: ListBranches :many
SELECT * FROM branches;

-- name: GetBranch :one
SELECT * FROM branches
WHERE id = $1;

-- name: CreateBranch :one
INSERT INTO branches 
(id, name, description)
VALUES ($1, $2, $3) RETURNING *;

-- name: UpdateBranch :exec
UPDATE branches
    set name = $2, description = $3
WHERE id = $1;

-- name: DeleteBranch :exec
DELETE FROM branches
WHERE id = $1;

--
--

-- Area queries
-- name: ListAreas :many
SELECT * FROM areas;

-- name: GetArea :one
SELECT * FROM areas
WHERE id = $1;

-- name: CreateArea :one
INSERT INTO areas
(id, name, description)
VALUES ($1, $2, $3) RETURNING *;

-- name: UpdateArea :exec
UPDATE areas
    set name = $2, description = $3
WHERE id = $1;

-- name: DeleteArea :exec
DELETE FROM companies
WHERE id = $1;

--
--

-- Department queries
-- name: ListDepartments :many
SELECT * FROM departments;

-- name: GetDepartment :one
SELECT * FROM departments
WHERE id = $1;

-- name: CreateDepartments :one
INSERT INTO departments 
(id, name, description)
VALUES ($1, $2, $3) RETURNING *;

-- name: UpdateDepartment :exec
UPDATE departments
    set name = $2, description = $3
WHERE id = $1;

-- name: DeleteDepartments :exec
DELETE FROM departments
WHERE id = $1;

--
--

-- Role queries
-- name: ListRoles :many
SELECT * FROM roles;

-- name: GetRole :one
SELECT * FROM roles 
WHERE id = $1;

-- name: CreateRole :one
INSERT INTO roles
(id, name, description)
VALUES ($1, $2, $3) RETURNING *;

-- name: UpdateRoles :exec
UPDATE roles
    set name = $2, description = $3
WHERE id = $1;

-- name: DeleteRoles :exec
DELETE FROM roles
WHERE id = $1;
