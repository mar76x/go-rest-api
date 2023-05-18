package main

import (
	"context"
	"fmt"
	_ "net/http/pprof"
	"os"

	"github.com/jackc/pgx/v5"
)

// Run - is responsible for the instantation and startup of the app
func Run() error {
	conn, err := pgx.Connect(context.Background(), os.Getenv("DB_URL"))
	if err != nil {
		return err
	}
	defer conn.Close(context.Background())
	fmt.Println("Successfully connected to database")
	fmt.Println("Starting up the application")
	return nil
}

// Application start-point
func main() {
	fmt.Println("Go REST API")
	if err := Run(); err != nil {
		// custom response if error
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
}
