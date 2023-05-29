package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/jackc/pgx/v5"
	"github.com/mar76x/go-rest-api/cmd/handler"
	db "github.com/mar76x/go-rest-api/db/sqlc"
	"github.com/mar76x/go-rest-api/util"
)

func main() {
	fmt.Println("GO REST API")
	ctx := context.Background()

	conn, err := connectDB()
	if err != nil {
		panic(err)
	}
	defer conn.Close(ctx)

	q := db.New(conn)
	cr := chi.NewRouter()
	cr.Use(middleware.RequestID)
	cr.Use(middleware.RealIP)
	cr.Use(middleware.Logger)
	cr.Use(middleware.Recoverer)
	cr.Use(middleware.Timeout(60 * time.Second))
	h := handler.NewHandler(cr, q)

	cr.Route("/company", func(cr chi.Router) {
		cr.Get("/", h.ListCompanies)
		cr.Post("/", h.CreateCompany)
	})

	s := &http.Server{
		Addr:         ":8080",
		Handler:      cr,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	go func() {
		fmt.Println("starting server on port 8080")
		if err := s.ListenAndServe(); err != nil {
			fmt.Printf("error starting server : %s\n", err)
			os.Exit(1)
		}
	}()

	// trap sigterm or interrupt and gracefully shutdown the server
	sigChannel := make(chan os.Signal)
	signal.Notify(sigChannel, os.Interrupt)
	signal.Notify(sigChannel, os.Kill)

	sig := <-sigChannel
	fmt.Println("recieved terminate, graceful shutdown", sig)

	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(tc)

}

func connectDB() (*pgx.Conn, error) {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config: ", err)
	}
	url := fmt.Sprint("postgres://", config.DB.User, ":", config.DB.Password, "@", config.DB.Host, ":", config.DB.Port, "/", config.DB.Name, "?sslmode=", config.DB.SSL)
	fmt.Println("connecting to database...")
	conn, err := pgx.Connect(context.Background(), url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "unable to connect to database: %v\n", err)
		return nil, err
	}
	fmt.Println("successfully connected to database")
	return conn, nil
}
