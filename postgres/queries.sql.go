// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0
// source: queries.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createArea = `-- name: CreateArea :one
INSERT INTO areas
(id, name, description)
VALUES ($1, $2, $3) RETURNING id, name, description
`

type CreateAreaParams struct {
	ID          pgtype.UUID
	Name        string
	Description string
}

func (q *Queries) CreateArea(ctx context.Context, arg CreateAreaParams) (Area, error) {
	row := q.db.QueryRow(ctx, createArea, arg.ID, arg.Name, arg.Description)
	var i Area
	err := row.Scan(&i.ID, &i.Name, &i.Description)
	return i, err
}

const createBranch = `-- name: CreateBranch :one
INSERT INTO branches 
(id, name, description)
VALUES ($1, $2, $3) RETURNING id, name, description
`

type CreateBranchParams struct {
	ID          pgtype.UUID
	Name        string
	Description string
}

func (q *Queries) CreateBranch(ctx context.Context, arg CreateBranchParams) (Branch, error) {
	row := q.db.QueryRow(ctx, createBranch, arg.ID, arg.Name, arg.Description)
	var i Branch
	err := row.Scan(&i.ID, &i.Name, &i.Description)
	return i, err
}

const createCompany = `-- name: CreateCompany :one
INSERT INTO companies 
(id, name, description)
VALUES ($1, $2, $3) RETURNING id, name, description
`

type CreateCompanyParams struct {
	ID          pgtype.UUID
	Name        string
	Description string
}

func (q *Queries) CreateCompany(ctx context.Context, arg CreateCompanyParams) (Company, error) {
	row := q.db.QueryRow(ctx, createCompany, arg.ID, arg.Name, arg.Description)
	var i Company
	err := row.Scan(&i.ID, &i.Name, &i.Description)
	return i, err
}

const createDepartments = `-- name: CreateDepartments :one
INSERT INTO departments 
(id, name, description)
VALUES ($1, $2, $3) RETURNING id, name, description
`

type CreateDepartmentsParams struct {
	ID          pgtype.UUID
	Name        string
	Description string
}

func (q *Queries) CreateDepartments(ctx context.Context, arg CreateDepartmentsParams) (Department, error) {
	row := q.db.QueryRow(ctx, createDepartments, arg.ID, arg.Name, arg.Description)
	var i Department
	err := row.Scan(&i.ID, &i.Name, &i.Description)
	return i, err
}

const createRole = `-- name: CreateRole :one
INSERT INTO roles
(id, name, description)
VALUES ($1, $2, $3) RETURNING id, name, description
`

type CreateRoleParams struct {
	ID          pgtype.UUID
	Name        string
	Description string
}

func (q *Queries) CreateRole(ctx context.Context, arg CreateRoleParams) (Role, error) {
	row := q.db.QueryRow(ctx, createRole, arg.ID, arg.Name, arg.Description)
	var i Role
	err := row.Scan(&i.ID, &i.Name, &i.Description)
	return i, err
}

const deleteArea = `-- name: DeleteArea :exec
DELETE FROM companies
WHERE id = $1
`

func (q *Queries) DeleteArea(ctx context.Context, id pgtype.UUID) error {
	_, err := q.db.Exec(ctx, deleteArea, id)
	return err
}

const deleteBranch = `-- name: DeleteBranch :exec
DELETE FROM branches
WHERE id = $1
`

func (q *Queries) DeleteBranch(ctx context.Context, id pgtype.UUID) error {
	_, err := q.db.Exec(ctx, deleteBranch, id)
	return err
}

const deleteCompany = `-- name: DeleteCompany :exec
DELETE FROM companies
WHERE id = $1
`

func (q *Queries) DeleteCompany(ctx context.Context, id pgtype.UUID) error {
	_, err := q.db.Exec(ctx, deleteCompany, id)
	return err
}

const deleteDepartments = `-- name: DeleteDepartments :exec
DELETE FROM departments
WHERE id = $1
`

func (q *Queries) DeleteDepartments(ctx context.Context, id pgtype.UUID) error {
	_, err := q.db.Exec(ctx, deleteDepartments, id)
	return err
}

const deleteRoles = `-- name: DeleteRoles :exec
DELETE FROM roles
WHERE id = $1
`

func (q *Queries) DeleteRoles(ctx context.Context, id pgtype.UUID) error {
	_, err := q.db.Exec(ctx, deleteRoles, id)
	return err
}

const getArea = `-- name: GetArea :one
SELECT id, name, description FROM areas
WHERE id = $1
`

func (q *Queries) GetArea(ctx context.Context, id pgtype.UUID) (Area, error) {
	row := q.db.QueryRow(ctx, getArea, id)
	var i Area
	err := row.Scan(&i.ID, &i.Name, &i.Description)
	return i, err
}

const getBranch = `-- name: GetBranch :one
SELECT id, name, description FROM branches
WHERE id = $1
`

func (q *Queries) GetBranch(ctx context.Context, id pgtype.UUID) (Branch, error) {
	row := q.db.QueryRow(ctx, getBranch, id)
	var i Branch
	err := row.Scan(&i.ID, &i.Name, &i.Description)
	return i, err
}

const getCompany = `-- name: GetCompany :one
SELECT id, name, description FROM companies 
WHERE id = $1
`

func (q *Queries) GetCompany(ctx context.Context, id pgtype.UUID) (Company, error) {
	row := q.db.QueryRow(ctx, getCompany, id)
	var i Company
	err := row.Scan(&i.ID, &i.Name, &i.Description)
	return i, err
}

const getDepartment = `-- name: GetDepartment :one
SELECT id, name, description FROM departments
WHERE id = $1
`

func (q *Queries) GetDepartment(ctx context.Context, id pgtype.UUID) (Department, error) {
	row := q.db.QueryRow(ctx, getDepartment, id)
	var i Department
	err := row.Scan(&i.ID, &i.Name, &i.Description)
	return i, err
}

const getRole = `-- name: GetRole :one
SELECT id, name, description FROM roles 
WHERE id = $1
`

func (q *Queries) GetRole(ctx context.Context, id pgtype.UUID) (Role, error) {
	row := q.db.QueryRow(ctx, getRole, id)
	var i Role
	err := row.Scan(&i.ID, &i.Name, &i.Description)
	return i, err
}

const listAreas = `-- name: ListAreas :many

SELECT id, name, description FROM areas
`

// Area queries
func (q *Queries) ListAreas(ctx context.Context) ([]Area, error) {
	rows, err := q.db.Query(ctx, listAreas)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Area
	for rows.Next() {
		var i Area
		if err := rows.Scan(&i.ID, &i.Name, &i.Description); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listBranches = `-- name: ListBranches :many

SELECT id, name, description FROM branches
`

// Branch queries
func (q *Queries) ListBranches(ctx context.Context) ([]Branch, error) {
	rows, err := q.db.Query(ctx, listBranches)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Branch
	for rows.Next() {
		var i Branch
		if err := rows.Scan(&i.ID, &i.Name, &i.Description); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listCompanies = `-- name: ListCompanies :many
SELECT id, name, description FROM companies
`

// Company queries
func (q *Queries) ListCompanies(ctx context.Context) ([]Company, error) {
	rows, err := q.db.Query(ctx, listCompanies)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Company
	for rows.Next() {
		var i Company
		if err := rows.Scan(&i.ID, &i.Name, &i.Description); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listDepartments = `-- name: ListDepartments :many

SELECT id, name, description FROM departments
`

// Department queries
func (q *Queries) ListDepartments(ctx context.Context) ([]Department, error) {
	rows, err := q.db.Query(ctx, listDepartments)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Department
	for rows.Next() {
		var i Department
		if err := rows.Scan(&i.ID, &i.Name, &i.Description); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listRoles = `-- name: ListRoles :many

SELECT id, name, description FROM roles
`

// Role queries
func (q *Queries) ListRoles(ctx context.Context) ([]Role, error) {
	rows, err := q.db.Query(ctx, listRoles)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Role
	for rows.Next() {
		var i Role
		if err := rows.Scan(&i.ID, &i.Name, &i.Description); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateArea = `-- name: UpdateArea :exec
UPDATE areas
    set name = $2, description = $3
WHERE id = $1
`

type UpdateAreaParams struct {
	ID          pgtype.UUID
	Name        string
	Description string
}

func (q *Queries) UpdateArea(ctx context.Context, arg UpdateAreaParams) error {
	_, err := q.db.Exec(ctx, updateArea, arg.ID, arg.Name, arg.Description)
	return err
}

const updateBranch = `-- name: UpdateBranch :exec
UPDATE branches
    set name = $2, description = $3
WHERE id = $1
`

type UpdateBranchParams struct {
	ID          pgtype.UUID
	Name        string
	Description string
}

func (q *Queries) UpdateBranch(ctx context.Context, arg UpdateBranchParams) error {
	_, err := q.db.Exec(ctx, updateBranch, arg.ID, arg.Name, arg.Description)
	return err
}

const updateCompany = `-- name: UpdateCompany :exec
UPDATE companies
    set name = $2, description = $3
WHERE id = $1
`

type UpdateCompanyParams struct {
	ID          pgtype.UUID
	Name        string
	Description string
}

func (q *Queries) UpdateCompany(ctx context.Context, arg UpdateCompanyParams) error {
	_, err := q.db.Exec(ctx, updateCompany, arg.ID, arg.Name, arg.Description)
	return err
}

const updateDepartment = `-- name: UpdateDepartment :exec
UPDATE departments
    set name = $2, description = $3
WHERE id = $1
`

type UpdateDepartmentParams struct {
	ID          pgtype.UUID
	Name        string
	Description string
}

func (q *Queries) UpdateDepartment(ctx context.Context, arg UpdateDepartmentParams) error {
	_, err := q.db.Exec(ctx, updateDepartment, arg.ID, arg.Name, arg.Description)
	return err
}

const updateRoles = `-- name: UpdateRoles :exec
UPDATE roles
    set name = $2, description = $3
WHERE id = $1
`

type UpdateRolesParams struct {
	ID          pgtype.UUID
	Name        string
	Description string
}

func (q *Queries) UpdateRoles(ctx context.Context, arg UpdateRolesParams) error {
	_, err := q.db.Exec(ctx, updateRoles, arg.ID, arg.Name, arg.Description)
	return err
}
