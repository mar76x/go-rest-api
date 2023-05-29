package handler

import (
	"github.com/go-chi/chi"
	db "github.com/mar76x/go-rest-api/db/sqlc"
)

type Handler struct {
	Mux        *chi.Mux
	Repository *db.Queries
}

func NewHandler(mux *chi.Mux, query *db.Queries) *Handler {
	h := &Handler{
		Mux:        mux,
		Repository: query,
	}
	return h
}
