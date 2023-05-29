package db

import (
	"context"
	"testing"

	db "github.com/mar76x/go-rest-api/db/sqlc"
	"github.com/mar76x/go-rest-api/util"
	"github.com/stretchr/testify/require"
)

func createRandomContract(t *testing.T) db.Contract {
	company := createRandomCompany(t)
	branch := createRandomBranch(t)
	area := createRandomArea(t)
	department := createRandomDepartment(t)
	role := createRandomRole(t)
	employee := createRandomEmployee(t)
	arg := db.CreateContractParams{
		ID:           util.RandomNumber(),
		Type:         util.RandomName(),
		StartDate:    util.RandomString(10),
		EmployeeID:   employee.ID,
		CompanyID:    company.ID,
		BranchID:     branch.ID,
		AreaID:       area.ID,
		DepartmentID: department.ID,
		RoleID:       role.ID,
	}
	contract, err := testQueries.CreateContract(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, contract)
	require.Equal(t, arg.ID, contract.ID)
	require.Equal(t, arg.Type, contract.Type)
	require.Equal(t, arg.StartDate, contract.StartDate)
	require.Equal(t, arg.EmployeeID, contract.EmployeeID)
	require.Equal(t, arg.CompanyID, contract.CompanyID)
	require.Equal(t, arg.BranchID, contract.BranchID)
	require.Equal(t, arg.AreaID, contract.AreaID)
	require.Equal(t, arg.DepartmentID, contract.DepartmentID)
	require.Equal(t, arg.RoleID, contract.RoleID)
	require.NotZero(t, contract.ID)
	require.NotZero(t, contract.StartDate)
	require.NotZero(t, contract.CreatedAt)
	return contract
}

func TestCreateContract(t *testing.T) {
	createRandomContract(t)
}

func TestGetContract(t *testing.T) {
	contract := createRandomContract(t)
	res, err := testQueries.GetContract(context.Background(), contract.ID)
	require.NoError(t, err)
	require.NotEmpty(t, res)
	require.Equal(t, contract.ID, res.ID)
	require.Equal(t, contract.Type, res.Type)
	require.Equal(t, contract.StartDate, res.StartDate)
	require.Equal(t, contract.EmployeeID, res.EmployeeID)
	require.Equal(t, contract.CompanyID, res.CompanyID)
	require.Equal(t, contract.BranchID, res.BranchID)
	require.Equal(t, contract.AreaID, res.AreaID)
	require.Equal(t, contract.DepartmentID, res.DepartmentID)
	require.Equal(t, contract.RoleID, res.RoleID)
	require.NotZero(t, res.ID)
	require.NotZero(t, res.StartDate)
	require.NotZero(t, res.CreatedAt)
}

func TestUpdateContract(t *testing.T) {
	contract := createRandomContract(t)
	company := createRandomCompany(t)
	branch := createRandomBranch(t)
	area := createRandomArea(t)
	department := createRandomDepartment(t)
	role := createRandomRole(t)
	employee := createRandomEmployee(t)
	arg := db.UpdateContractParams{
		ID:           contract.ID,
		Type:         util.RandomName(),
		StartDate:    util.RandomString(10),
		EmployeeID:   employee.ID,
		CompanyID:    company.ID,
		BranchID:     branch.ID,
		AreaID:       area.ID,
		DepartmentID: department.ID,
		RoleID:       role.ID,
	}
	err1 := testQueries.UpdateContract(context.Background(), arg)
	require.NoError(t, err1)
	res, err2 := testQueries.GetContract(context.Background(), contract.ID)
	require.NoError(t, err2)
	require.NotEmpty(t, res)
	require.Equal(t, arg.ID, res.ID)
	require.Equal(t, arg.Type, res.Type)
	require.Equal(t, arg.StartDate, res.StartDate)
	require.Equal(t, arg.EmployeeID, res.EmployeeID)
	require.Equal(t, arg.CompanyID, res.CompanyID)
	require.Equal(t, arg.BranchID, res.BranchID)
	require.Equal(t, arg.AreaID, res.AreaID)
	require.Equal(t, arg.DepartmentID, res.DepartmentID)
	require.Equal(t, arg.RoleID, res.RoleID)
	require.NotZero(t, res.ID)
	require.NotZero(t, res.StartDate)
	require.NotZero(t, res.CreatedAt)
}

func TestDeleteContract(t *testing.T) {
	contract := createRandomContract(t)
	err := testQueries.DeleteContract(context.Background(), contract.ID)
	require.NoError(t, err)
	res, err := testQueries.GetContract(context.Background(), contract.ID)
	require.Error(t, err)
	require.EqualError(t, err, "no rows in result set")
	require.Empty(t, res)
}

func TestListContracts(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomContract(t)
	}
	arg := db.ListContractParams{
		Limit:  5,
		Offset: 5,
	}
	contracts, err := testQueries.ListContract(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, contracts, 5)
	for _, c := range contracts {
		require.NotEmpty(t, c)
	}
}
