// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0

package db

import (
	uuid "github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

type Area struct {
	ID          int64              `json:"id"`
	Name        string             `json:"name"`
	Description string             `json:"description"`
	CreatedAt   pgtype.Timestamptz `json:"created_at"`
	UpdatedAt   pgtype.Timestamptz `json:"updated_at"`
	DeletedAt   pgtype.Timestamptz `json:"deleted_at"`
}

type Branch struct {
	ID          int64              `json:"id"`
	CompanyID   int64              `json:"company_id"`
	Name        string             `json:"name"`
	Description string             `json:"description"`
	CreatedAt   pgtype.Timestamptz `json:"created_at"`
	UpdatedAt   pgtype.Timestamptz `json:"updated_at"`
	DeletedAt   pgtype.Timestamptz `json:"deleted_at"`
}

type Company struct {
	ID          int64              `json:"id"`
	Name        string             `json:"name"`
	Description string             `json:"description"`
	CreatedAt   pgtype.Timestamptz `json:"created_at"`
	UpdatedAt   pgtype.Timestamptz `json:"updated_at"`
	DeletedAt   pgtype.Timestamptz `json:"deleted_at"`
}

type Contract struct {
	ID           int64              `json:"id"`
	Type         string             `json:"type"`
	StartDate    string             `json:"start_date"`
	EmployeeID   uuid.UUID          `json:"employee_id"`
	CompanyID    int64              `json:"company_id"`
	BranchID     int64              `json:"branch_id"`
	AreaID       int64              `json:"area_id"`
	DepartmentID int64              `json:"department_id"`
	RoleID       int64              `json:"role_id"`
	CreatedAt    pgtype.Timestamptz `json:"created_at"`
	UpdatedAt    pgtype.Timestamptz `json:"updated_at"`
	DeletedAt    pgtype.Timestamptz `json:"deleted_at"`
}

type Department struct {
	ID          int64              `json:"id"`
	Name        string             `json:"name"`
	Description string             `json:"description"`
	CreatedAt   pgtype.Timestamptz `json:"created_at"`
	UpdatedAt   pgtype.Timestamptz `json:"updated_at"`
	DeletedAt   pgtype.Timestamptz `json:"deleted_at"`
}

type Employee struct {
	ID            uuid.UUID          `json:"id"`
	Number        int64              `json:"number"`
	Name          string             `json:"name"`
	Surname       string             `json:"surname"`
	Birthdate     string             `json:"birthdate"`
	Dni           string             `json:"dni"`
	Cuil          string             `json:"cuil"`
	MaritalStatus string             `json:"marital_status"`
	CreatedAt     pgtype.Timestamptz `json:"created_at"`
	UpdatedAt     pgtype.Timestamptz `json:"updated_at"`
	DeletedAt     pgtype.Timestamptz `json:"deleted_at"`
}

type Paycheck struct {
	ID          int64              `json:"id"`
	Type        string             `json:"type"`
	Filename    string             `json:"filename"`
	Description string             `json:"description"`
	Folder      string             `json:"folder"`
	Path        string             `json:"path"`
	Read        bool               `json:"read"`
	Signed      bool               `json:"signed"`
	EmployeeID  uuid.UUID          `json:"employee_id"`
	CreatedAt   pgtype.Timestamptz `json:"created_at"`
	UpdatedAt   pgtype.Timestamptz `json:"updated_at"`
	DeletedAt   pgtype.Timestamptz `json:"deleted_at"`
}

type Role struct {
	ID          int64              `json:"id"`
	Name        string             `json:"name"`
	Description string             `json:"description"`
	CreatedAt   pgtype.Timestamptz `json:"created_at"`
	UpdatedAt   pgtype.Timestamptz `json:"updated_at"`
	DeletedAt   pgtype.Timestamptz `json:"deleted_at"`
}
