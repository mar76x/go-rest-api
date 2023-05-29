package db

import (
	"context"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/jackc/pgx/v5"
	db "github.com/mar76x/go-rest-api/db/sqlc"
	"github.com/mar76x/go-rest-api/util"
)

var testDB *pgx.Conn // to create Tx
var testQueries *db.Queries

func TestMain(m *testing.M) {
	config, err := util.LoadConfig("../..")
	if err != nil {
		log.Fatal("cannot load config: ", err)
	}
	url := fmt.Sprint("postgres://", config.DB.User, ":", config.DB.Password, "@", config.DB.Host, ":", config.DB.Port, "/", config.DB.Name, "?sslmode=", config.DB.SSL)
	conn, err := pgx.Connect(context.Background(), url)
	if err != nil {
		log.Fatal("unable to connect to database:", err)
	}
	testDB = conn
	testQueries = db.New(conn)
	os.Exit(m.Run())
}
