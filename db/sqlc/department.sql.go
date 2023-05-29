// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0
// source: department.sql

package db

import (
	"context"
)

const createDepartment = `-- name: CreateDepartment :one
INSERT INTO department 
(id, name, description)
VALUES ($1, $2, $3) RETURNING id, name, description, created_at, updated_at, deleted_at
`

type CreateDepartmentParams struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (q *Queries) CreateDepartment(ctx context.Context, arg CreateDepartmentParams) (Department, error) {
	row := q.db.QueryRow(ctx, createDepartment, arg.ID, arg.Name, arg.Description)
	var i Department
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Description,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}

const deleteDepartment = `-- name: DeleteDepartment :exec
DELETE FROM department
WHERE id = $1
`

func (q *Queries) DeleteDepartment(ctx context.Context, id int64) error {
	_, err := q.db.Exec(ctx, deleteDepartment, id)
	return err
}

const getDepartment = `-- name: GetDepartment :one
SELECT id, name, description, created_at, updated_at, deleted_at FROM department
WHERE id = $1
`

func (q *Queries) GetDepartment(ctx context.Context, id int64) (Department, error) {
	row := q.db.QueryRow(ctx, getDepartment, id)
	var i Department
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Description,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}

const listDepartments = `-- name: ListDepartments :many
SELECT id, name, description, created_at, updated_at, deleted_at FROM department
ORDER BY id
LIMIT $1
OFFSET $2
`

type ListDepartmentsParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListDepartments(ctx context.Context, arg ListDepartmentsParams) ([]Department, error) {
	rows, err := q.db.Query(ctx, listDepartments, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Department
	for rows.Next() {
		var i Department
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Description,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.DeletedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateDepartment = `-- name: UpdateDepartment :exec
UPDATE department
    set name = $2, description = $3
WHERE id = $1
`

type UpdateDepartmentParams struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (q *Queries) UpdateDepartment(ctx context.Context, arg UpdateDepartmentParams) error {
	_, err := q.db.Exec(ctx, updateDepartment, arg.ID, arg.Name, arg.Description)
	return err
}
