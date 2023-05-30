package db

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/jackc/pgx/v5"
	db "github.com/mar76x/go-rest-api/db/sqlc"
)

var testDB *pgx.Conn // to create Tx
var testQueries *db.Queries

func TestMain(m *testing.M) {
	url := "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable"
	conn, err := pgx.Connect(context.Background(), url)
	if err != nil {
		log.Fatal("unable to connect to database:\n", err)
	}
	testDB = conn
	testQueries = db.New(conn)
	os.Exit(m.Run())
}
