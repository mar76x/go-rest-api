package db

import (
	"context"
	"testing"

	db "github.com/mar76x/go-rest-api/db/sqlc"
	"github.com/mar76x/go-rest-api/util"
	"github.com/stretchr/testify/require"
)

func createRandomBranch(t *testing.T) db.Branch {
	company := createRandomCompany(t)
	arg := db.CreateBranchParams{
		ID:          util.RandomNumber(),
		CompanyID:   company.ID,
		Name:        util.RandomName(),
		Description: util.RandomName(),
	}
	branch, err := testQueries.CreateBranch(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, branch)
	require.Equal(t, arg.ID, branch.ID)
	require.Equal(t, arg.Name, branch.Name)
	require.Equal(t, arg.Description, branch.Description)
	require.NotZero(t, branch.ID)
	require.NotZero(t, branch.CreatedAt)
	return branch
}

func TestCreateBranch(t *testing.T) {
	createRandomBranch(t)
}

func TestGetBranch(t *testing.T) {
	branch := createRandomBranch(t)
	res, err := testQueries.GetBranch(context.Background(), branch.ID)
	require.NoError(t, err)
	require.NotEmpty(t, res)
	require.Equal(t, branch.ID, res.ID)
	require.Equal(t, branch.Name, res.Name)
	require.Equal(t, branch.Description, res.Description)
	require.NotZero(t, res.ID)
	require.NotZero(t, res.CreatedAt)
}

func TestUpdateBranch(t *testing.T) {
	branch := createRandomBranch(t)
	arg := db.UpdateBranchParams{
		ID:          branch.ID,
		Name:        util.RandomName(),
		Description: util.RandomName(),
	}
	err1 := testQueries.UpdateBranch(context.Background(), arg)
	require.NoError(t, err1)
	res, err2 := testQueries.GetBranch(context.Background(), branch.ID)
	require.NoError(t, err2)
	require.NotEmpty(t, res)
	require.Equal(t, branch.ID, res.ID)
	require.Equal(t, arg.Name, res.Name)
	require.Equal(t, arg.Description, res.Description)
	require.NotZero(t, res.ID)
	require.NotZero(t, res.CreatedAt)
}

func TestDeleteBranch(t *testing.T) {
	branch := createRandomBranch(t)
	err := testQueries.DeleteBranch(context.Background(), branch.ID)
	require.NoError(t, err)
	res, err := testQueries.GetBranch(context.Background(), branch.ID)
	require.Error(t, err)
	require.EqualError(t, err, "no rows in result set")
	require.Empty(t, res)
}

func TestListBranches(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomBranch(t)
	}
	arg := db.ListBranchesParams{
		Limit:  5,
		Offset: 5,
	}
	companies, err := testQueries.ListBranches(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, companies, 5)
	for _, c := range companies {
		require.NotEmpty(t, c)
	}
}
