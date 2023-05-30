package db

import (
	"context"
	"testing"

	db "github.com/mar76x/go-rest-api/db/sqlc"
	"github.com/mar76x/go-rest-api/util"
	"github.com/stretchr/testify/require"
)

func createRandomRole(t *testing.T) db.Role {
	arg := db.CreateRoleParams{
		Name:        util.RandomName(),
		Description: util.RandomName(),
	}
	role, err := testQueries.CreateRole(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, role)
	require.Equal(t, arg.Name, role.Name)
	require.Equal(t, arg.Description, role.Description)
	require.NotZero(t, role.ID)
	require.NotZero(t, role.CreatedAt)
	return role
}

func TestCreateRole(t *testing.T) {
	createRandomRole(t)
}

func TestGetRole(t *testing.T) {
	role := createRandomRole(t)
	res, err := testQueries.GetRole(context.Background(), role.ID)
	require.NoError(t, err)
	require.NotEmpty(t, res)
	require.Equal(t, role.ID, res.ID)
	require.Equal(t, role.Name, res.Name)
	require.Equal(t, role.Description, res.Description)
	require.NotZero(t, res.ID)
	require.NotZero(t, res.CreatedAt)
}

func TestUpdateRole(t *testing.T) {
	role := createRandomRole(t)
	arg := db.UpdateRoleParams{
		ID:          role.ID,
		Name:        util.RandomName(),
		Description: util.RandomName(),
	}
	err1 := testQueries.UpdateRole(context.Background(), arg)
	require.NoError(t, err1)
	res, err2 := testQueries.GetRole(context.Background(), role.ID)
	require.NoError(t, err2)
	require.NotEmpty(t, res)
	require.Equal(t, arg.ID, res.ID)
	require.Equal(t, arg.Name, res.Name)
	require.Equal(t, arg.Description, res.Description)
	require.NotZero(t, res.ID)
	require.NotZero(t, res.CreatedAt)
}

func TestDeleteRole(t *testing.T) {
	role := createRandomRole(t)
	err := testQueries.DeleteRole(context.Background(), role.ID)
	require.NoError(t, err)
	res, err := testQueries.GetRole(context.Background(), role.ID)
	require.Error(t, err)
	require.EqualError(t, err, "no rows in result set")
	require.Empty(t, res)
}

func TestListRoles(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomRole(t)
	}
	arg := db.ListRolesParams{
		Limit:  5,
		Offset: 5,
	}
	roles, err := testQueries.ListRoles(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, roles, 5)
	for _, c := range roles {
		require.NotEmpty(t, c)
	}
}
