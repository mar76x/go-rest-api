package handler

import (
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/jackc/pgx/v5"
	db "github.com/mar76x/go-rest-api/db/sqlc"
)

type Handler struct {
	Mux        *chi.Mux
	Repository *db.Queries
}

func NewHandler(conn *pgx.Conn) *Handler {
	h := &Handler{
		Mux:        chi.NewMux(),
		Repository: db.New(conn),
	}

	h.Mux.Use(middleware.RequestID)
	h.Mux.Use(middleware.RealIP)
	h.Mux.Use(middleware.Logger)
	h.Mux.Use(middleware.Recoverer)
	h.Mux.Use(middleware.Timeout(60 * time.Second))
	h.Mux.Route("/company", func(r chi.Router) {
		r.Get("/", h.ListCompanies)
		r.Get("/{id}", h.GetCompany)
		r.Post("/", h.CreateCompany)
		r.Put("/{id}", h.UpdateCompany)
		r.Delete("/{id}", h.DeleteCompany)
	})

	return h
}
