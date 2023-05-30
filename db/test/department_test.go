package db

import (
	"context"
	"testing"

	db "github.com/mar76x/go-rest-api/db/sqlc"
	"github.com/mar76x/go-rest-api/util"
	"github.com/stretchr/testify/require"
)

func createRandomDepartment(t *testing.T) db.Department {
	arg := db.CreateDepartmentParams{
		Name:        util.RandomName(),
		Description: util.RandomName(),
	}
	department, err := testQueries.CreateDepartment(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, department)
	require.Equal(t, arg.Name, department.Name)
	require.Equal(t, arg.Description, department.Description)
	require.NotZero(t, department.ID)
	require.NotZero(t, department.CreatedAt)
	return department
}

func TestCreateDepartment(t *testing.T) {
	createRandomDepartment(t)
}

func TestGetDepartment(t *testing.T) {
	department := createRandomDepartment(t)
	res, err := testQueries.GetDepartment(context.Background(), department.ID)
	require.NoError(t, err)
	require.NotEmpty(t, res)
	require.Equal(t, department.ID, res.ID)
	require.Equal(t, department.Name, res.Name)
	require.Equal(t, department.Description, res.Description)
	require.NotZero(t, res.ID)
	require.NotZero(t, res.CreatedAt)
}

func TestUpdateDepartment(t *testing.T) {
	department := createRandomDepartment(t)
	arg := db.UpdateDepartmentParams{
		ID:          department.ID,
		Name:        util.RandomName(),
		Description: util.RandomName(),
	}
	err1 := testQueries.UpdateDepartment(context.Background(), arg)
	require.NoError(t, err1)
	res, err2 := testQueries.GetDepartment(context.Background(), department.ID)
	require.NoError(t, err2)
	require.NotEmpty(t, res)
	require.Equal(t, department.ID, res.ID)
	require.Equal(t, arg.Name, res.Name)
	require.Equal(t, arg.Description, res.Description)
	require.NotZero(t, res.ID)
	require.NotZero(t, res.CreatedAt)
}

func TestDeleteDepartment(t *testing.T) {
	department := createRandomDepartment(t)
	err := testQueries.DeleteDepartment(context.Background(), department.ID)
	require.NoError(t, err)
	res, err := testQueries.GetDepartment(context.Background(), department.ID)
	require.Error(t, err)
	require.EqualError(t, err, "no rows in result set")
	require.Empty(t, res)
}

func TestListDepartments(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomDepartment(t)
	}
	arg := db.ListDepartmentsParams{
		Limit:  5,
		Offset: 5,
	}
	companies, err := testQueries.ListDepartments(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, companies, 5)
	for _, c := range companies {
		require.NotEmpty(t, c)
	}
}
