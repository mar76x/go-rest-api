package db

import (
	"context"
	"testing"

	db "github.com/mar76x/go-rest-api/db/sqlc"
	"github.com/mar76x/go-rest-api/util"
	"github.com/stretchr/testify/require"
)

func createRandomCompany(t *testing.T) db.Company {
	arg := db.CreateCompanyParams{
		Name:        util.RandomName(),
		Description: util.RandomName(),
	}
	company, err := testQueries.CreateCompany(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, company)
	require.Equal(t, arg.Name, company.Name)
	require.Equal(t, arg.Description, company.Description)
	require.NotZero(t, company.ID)
	require.NotZero(t, company.CreatedAt)
	return company
}

func TestCreateCompany(t *testing.T) {
	createRandomCompany(t)
}

func TestGetCompany(t *testing.T) {
	company := createRandomCompany(t)
	res, err := testQueries.GetCompany(context.Background(), company.ID)
	require.NoError(t, err)
	require.NotEmpty(t, res)
	require.Equal(t, company.ID, res.ID)
	require.Equal(t, company.Name, res.Name)
	require.Equal(t, company.Description, res.Description)
	require.NotZero(t, res.ID)
	require.NotZero(t, res.CreatedAt)
}

func TestUpdateCompany(t *testing.T) {
	company := createRandomCompany(t)
	arg := db.UpdateCompanyParams{
		ID:          company.ID,
		Name:        util.RandomName(),
		Description: util.RandomName(),
	}
	err1 := testQueries.UpdateCompany(context.Background(), arg)
	require.NoError(t, err1)
	res, err2 := testQueries.GetCompany(context.Background(), company.ID)
	require.NoError(t, err2)
	require.NotEmpty(t, res)
	require.Equal(t, company.ID, res.ID)
	require.Equal(t, arg.Name, res.Name)
	require.Equal(t, arg.Description, res.Description)
	require.NotZero(t, res.ID)
	require.NotZero(t, res.CreatedAt)
}

func TestDeleteCompany(t *testing.T) {
	company := createRandomCompany(t)
	err := testQueries.DeleteCompany(context.Background(), company.ID)
	require.NoError(t, err)
	res, err := testQueries.GetCompany(context.Background(), company.ID)
	require.Error(t, err)
	require.EqualError(t, err, "no rows in result set")
	require.Empty(t, res)
}

func TestListCompanies(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomCompany(t)
	}
	arg := db.ListCompaniesParams{
		Limit:  5,
		Offset: 5,
	}
	companies, err := testQueries.ListCompanies(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, companies, 5)
	for _, c := range companies {
		require.NotEmpty(t, c)
	}
}
