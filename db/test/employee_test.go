package db

import (
	"context"
	"testing"

	"github.com/google/uuid"
	db "github.com/mar76x/go-rest-api/db/sqlc"
	"github.com/mar76x/go-rest-api/util"
	"github.com/stretchr/testify/require"
)

func createRandomEmployee(t *testing.T) db.Employee {
	arg := db.CreateEmployeeParams{
		ID:            uuid.New(),
		Number:        util.RandomNumber(),
		Name:          util.RandomName(),
		Surname:       util.RandomName(),
		Birthdate:     util.RandomString(10),
		Dni:           util.RandomName(),
		Cuil:          util.RandomName(),
		MaritalStatus: util.RandomMaritalStatus(),
	}
	employee, err := testQueries.CreateEmployee(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, employee)
	require.Equal(t, arg.ID, employee.ID)
	require.Equal(t, arg.Number, employee.Number)
	require.Equal(t, arg.Name, employee.Name)
	require.Equal(t, arg.Surname, employee.Surname)
	require.Equal(t, arg.Birthdate, employee.Birthdate)
	require.Equal(t, arg.Dni, employee.Dni)
	require.Equal(t, arg.Cuil, employee.Cuil)
	require.Equal(t, arg.MaritalStatus, employee.MaritalStatus)
	require.NotZero(t, employee.ID)
	require.NotZero(t, employee.CreatedAt)
	return employee
}

func TestCreateEmployee(t *testing.T) {
	createRandomEmployee(t)
}

func TestGetEmployee(t *testing.T) {
	employee := createRandomEmployee(t)
	res, err := testQueries.GetEmployee(context.Background(), employee.ID)
	require.NoError(t, err)
	require.NotEmpty(t, res)
	require.Equal(t, employee.ID, res.ID)
	require.Equal(t, employee.Number, res.Number)
	require.Equal(t, employee.Name, res.Name)
	require.Equal(t, employee.Surname, res.Surname)
	require.Equal(t, employee.Birthdate, res.Birthdate)
	require.Equal(t, employee.Dni, res.Dni)
	require.Equal(t, employee.Cuil, res.Cuil)
	require.Equal(t, employee.MaritalStatus, res.MaritalStatus)
	require.NotZero(t, res.ID)
	require.NotZero(t, res.CreatedAt)
}

func TestUpdateEmployee(t *testing.T) {
	employee := createRandomEmployee(t)
	arg := db.UpdateEmployeeParams{
		ID:            employee.ID,
		Number:        util.RandomNumber(),
		Name:          util.RandomName(),
		Surname:       util.RandomName(),
		Birthdate:     util.RandomString(10),
		Dni:           util.RandomName(),
		Cuil:          util.RandomName(),
		MaritalStatus: util.RandomMaritalStatus(),
	}
	err1 := testQueries.UpdateEmployee(context.Background(), arg)
	require.NoError(t, err1)
	res, err2 := testQueries.GetEmployee(context.Background(), employee.ID)
	require.NoError(t, err2)
	require.NotEmpty(t, res)
	require.Equal(t, arg.ID, res.ID)
	require.Equal(t, arg.Number, res.Number)
	require.Equal(t, arg.Name, res.Name)
	require.Equal(t, arg.Surname, res.Surname)
	require.Equal(t, arg.Birthdate, res.Birthdate)
	require.Equal(t, arg.Dni, res.Dni)
	require.Equal(t, arg.Cuil, res.Cuil)
	require.Equal(t, arg.MaritalStatus, res.MaritalStatus)
	require.NotZero(t, res.ID)
	require.NotZero(t, res.CreatedAt)
}

func TestDeleteEmployee(t *testing.T) {
	employee := createRandomEmployee(t)
	err := testQueries.DeleteEmployee(context.Background(), employee.ID)
	require.NoError(t, err)
	res, err := testQueries.GetEmployee(context.Background(), employee.ID)
	require.Error(t, err)
	require.EqualError(t, err, "no rows in result set")
	require.Empty(t, res)
}

func TestListEmployees(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomEmployee(t)
	}
	arg := db.ListEmployeesParams{
		Limit:  5,
		Offset: 5,
	}
	companies, err := testQueries.ListEmployees(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, companies, 5)
	for _, c := range companies {
		require.NotEmpty(t, c)
	}
}
