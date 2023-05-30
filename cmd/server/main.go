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

	"github.com/jackc/pgx/v5"
	"github.com/mar76x/go-rest-api/cmd/handler"
	"github.com/mar76x/go-rest-api/util"
)

func main() {
	fmt.Println("GO REST API")
	ctx := context.Background()

	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("ERROR: cannot load config from file\n", err)
	}

	conn, err := connectDB(config)
	if err != nil {
		log.Fatal("ERROR:", err)
	}
	defer conn.Close(ctx)

	h := handler.NewHandler(conn)

	s := &http.Server{
		Addr:         ":" + config.Server.Port,
		Handler:      h.Mux,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	go func() {
		log.Printf("starting server on port %s", s.Addr)
		if err := s.ListenAndServe(); err != nil {
			log.Fatal("ERROR: cannot start server\n", err)
		}
	}()

	// trap sigterm or interrupt and gracefully shutdown the server
	sigChannel := make(chan os.Signal)
	signal.Notify(sigChannel, os.Interrupt)
	signal.Notify(sigChannel, os.Kill)

	sig := <-sigChannel
	log.Println("INFO: recieved terminate, graceful shutdown", sig)

	c, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(c)
}

func connectDB(config util.Config) (*pgx.Conn, error) {
	url := fmt.Sprint("postgres://", config.DB.User, ":", config.DB.Password, "@", config.DB.Host, ":", config.DB.Port, "/", config.DB.Name, "?sslmode=", config.DB.SSL)
	log.Println("connecting to database...")
	conn, err := pgx.Connect(context.Background(), url)
	if err != nil {
		log.Fatal("ERROR: unable to connect to database", err)
		return nil, err
	}
	log.Println("INFO: successfully connected to database")
	return conn, nil
}
