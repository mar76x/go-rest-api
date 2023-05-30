package db

import (
	"context"
	"testing"

	db "github.com/mar76x/go-rest-api/db/sqlc"
	"github.com/mar76x/go-rest-api/util"
	"github.com/stretchr/testify/require"
)

func createRandomArea(t *testing.T) db.Area {
	arg := db.CreateAreaParams{
		Name:        util.RandomName(),
		Description: util.RandomName(),
	}
	area, err := testQueries.CreateArea(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, area)
	require.Equal(t, arg.Name, area.Name)
	require.Equal(t, arg.Description, area.Description)
	require.NotZero(t, area.ID)
	require.NotZero(t, area.CreatedAt)
	return area
}

func TestCreateArea(t *testing.T) {
	createRandomArea(t)
}

func TestGetArea(t *testing.T) {
	area := createRandomArea(t)
	res, err := testQueries.GetArea(context.Background(), area.ID)
	require.NoError(t, err)
	require.NotEmpty(t, res)
	require.Equal(t, area.ID, res.ID)
	require.Equal(t, area.Name, res.Name)
	require.Equal(t, area.Description, res.Description)
	require.NotZero(t, res.ID)
	require.NotZero(t, res.CreatedAt)
}

func TestUpdateArea(t *testing.T) {
	area := createRandomArea(t)
	arg := db.UpdateAreaParams{
		ID:          area.ID,
		Name:        util.RandomName(),
		Description: util.RandomName(),
	}
	err1 := testQueries.UpdateArea(context.Background(), arg)
	require.NoError(t, err1)
	res, err2 := testQueries.GetArea(context.Background(), area.ID)
	require.NoError(t, err2)
	require.NotEmpty(t, res)
	require.Equal(t, area.ID, res.ID)
	require.Equal(t, arg.Name, res.Name)
	require.Equal(t, arg.Description, res.Description)
	require.NotZero(t, res.ID)
	require.NotZero(t, res.CreatedAt)
}

func TestDeleteArea(t *testing.T) {
	area := createRandomArea(t)
	err := testQueries.DeleteArea(context.Background(), area.ID)
	require.NoError(t, err)
	res, err := testQueries.GetArea(context.Background(), area.ID)
	require.Error(t, err)
	require.EqualError(t, err, "no rows in result set")
	require.Empty(t, res)
}

func TestListAreas(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomArea(t)
	}
	arg := db.ListAreasParams{
		Limit:  5,
		Offset: 5,
	}
	companies, err := testQueries.ListAreas(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, companies, 5)
	for _, c := range companies {
		require.NotEmpty(t, c)
	}
}
