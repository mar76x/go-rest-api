package db

import (
	"context"
	"testing"

	db "github.com/mar76x/go-rest-api/db/sqlc"
	"github.com/mar76x/go-rest-api/util"
	"github.com/stretchr/testify/require"
)

func createRandomPaycheck(t *testing.T) db.Paycheck {
	employee := createRandomEmployee(t)
	arg := db.CreatePaycheckParams{
		Type:        util.RandomName(),
		Filename:    util.RandomName(),
		Description: util.RandomName(),
		Folder:      util.RandomName(),
		Path:        util.RandomName(),
		Read:        true,
		Signed:      true,
		EmployeeID:  employee.ID,
	}
	paycheck, err := testQueries.CreatePaycheck(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, paycheck)
	require.Equal(t, arg.Type, paycheck.Type)
	require.Equal(t, arg.Filename, paycheck.Filename)
	require.Equal(t, arg.Description, paycheck.Description)
	require.Equal(t, arg.Folder, paycheck.Folder)
	require.Equal(t, arg.Path, paycheck.Path)
	require.Equal(t, arg.Read, paycheck.Read)
	require.Equal(t, arg.Signed, paycheck.Signed)
	require.Equal(t, arg.EmployeeID, paycheck.EmployeeID)
	require.NotZero(t, paycheck.ID)
	require.NotZero(t, paycheck.CreatedAt)
	return paycheck
}

func TestCreatePaycheck(t *testing.T) {
	createRandomPaycheck(t)
}

func TestGetPaycheck(t *testing.T) {
	paycheck := createRandomPaycheck(t)
	res, err := testQueries.GetPaycheck(context.Background(), paycheck.ID)
	require.NoError(t, err)
	require.NotEmpty(t, res)
	require.Equal(t, paycheck.ID, res.ID)
	require.Equal(t, paycheck.Type, res.Type)
	require.Equal(t, paycheck.Filename, res.Filename)
	require.Equal(t, paycheck.Description, res.Description)
	require.Equal(t, paycheck.Folder, res.Folder)
	require.Equal(t, paycheck.Path, res.Path)
	require.Equal(t, paycheck.Read, res.Read)
	require.Equal(t, paycheck.Signed, res.Signed)
	require.Equal(t, paycheck.EmployeeID, res.EmployeeID)
	require.NotZero(t, res.ID)
	require.NotZero(t, res.CreatedAt)
}

func TestUpdatePaycheck(t *testing.T) {
	employee := createRandomEmployee(t)
	paycheck := createRandomPaycheck(t)
	arg := db.UpdatePaycheckParams{
		ID:          paycheck.ID,
		Type:        util.RandomName(),
		Filename:    util.RandomName(),
		Description: util.RandomName(),
		Folder:      util.RandomName(),
		Path:        util.RandomName(),
		Read:        true,
		Signed:      true,
		EmployeeID:  employee.ID,
	}
	err1 := testQueries.UpdatePaycheck(context.Background(), arg)
	require.NoError(t, err1)
	res, err2 := testQueries.GetPaycheck(context.Background(), paycheck.ID)
	require.NoError(t, err2)
	require.NotEmpty(t, res)
	require.Equal(t, paycheck.ID, res.ID)
	require.Equal(t, arg.Type, res.Type)
	require.Equal(t, arg.Filename, res.Filename)
	require.Equal(t, arg.Description, res.Description)
	require.Equal(t, arg.Folder, res.Folder)
	require.Equal(t, arg.Path, res.Path)
	require.Equal(t, arg.Read, res.Read)
	require.Equal(t, arg.Signed, res.Signed)
	require.Equal(t, arg.EmployeeID, res.EmployeeID)
	require.NotZero(t, res.ID)
	require.NotZero(t, res.CreatedAt)
}

func TestDeletePaycheck(t *testing.T) {
	paycheck := createRandomPaycheck(t)
	err := testQueries.DeletePaycheck(context.Background(), paycheck.ID)
	require.NoError(t, err)
	res, err := testQueries.GetPaycheck(context.Background(), paycheck.ID)
	require.Error(t, err)
	require.EqualError(t, err, "no rows in result set")
	require.Empty(t, res)
}

func TestListPaychecks(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomPaycheck(t)
	}
	arg := db.ListPaychecksParams{
		Limit:  5,
		Offset: 5,
	}
	companies, err := testQueries.ListPaychecks(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, companies, 5)
	for _, c := range companies {
		require.NotEmpty(t, c)
	}
}
