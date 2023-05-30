package handler

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	db "github.com/mar76x/go-rest-api/db/sqlc"
)

func (h *Handler) ListCompanies(w http.ResponseWriter, r *http.Request) {
	limit := chi.URLParam(r, "limit")
	offset := chi.URLParam(r, "offset")

	if limit == "" {
		limit = "10"
	}
	l, err := strconv.ParseInt(limit, 10, 32)
	if err != nil {
		log.Println("invalid 'limit' query param\n", err)
		return
	}
	o, err := strconv.ParseInt(offset, 10, 32)
	if err != nil && o != 0 {
		log.Println("invalid 'offset' query param\n", err)
		return
	}

	arg := db.ListCompaniesParams{Limit: int32(l), Offset: int32(o)}
	companyList, err := h.Repository.ListCompanies(context.Background(), arg)
	if err != nil {
		log.Println(err)
		http.Error(w, "unable to retrive data from database", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(companyList); err != nil {
		log.Println(err)
		http.Error(w, "unable to encode json", http.StatusInternalServerError)
		return
	}
}

func (h *Handler) GetCompany(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		log.Println(err)
		http.Error(w, "invalid id format", http.StatusBadRequest)
		return
	}

	company, err := h.Repository.GetCompany(context.Background(), id)
	if err != nil {
		log.Println(err)
		http.Error(w, http.StatusText(404), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(company); err != nil {
		log.Println(err)
		http.Error(w, "unable to encode json", http.StatusInternalServerError)
		return
	}
}

func (h *Handler) CreateCompany(w http.ResponseWriter, r *http.Request) {
	var company db.Company
	if err := json.NewDecoder(r.Body).Decode(&company); err != nil {
		log.Println(err)
		http.Error(w, "error decoding json body", http.StatusBadRequest)
		return
	}

	arg := db.CreateCompanyParams{Name: company.Name, Description: company.Description}
	h.Repository.CreateCompany(context.Background(), arg)
}

func (h *Handler) UpdateCompany(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		log.Println(err)
		http.Error(w, "Invalid id format", http.StatusBadRequest)
		return
	}

	var company db.Company
	if err := json.NewDecoder(r.Body).Decode(&company); err != nil {
		log.Println(err)
		http.Error(w, "error decoding json body", http.StatusBadRequest)
		return
	}

	arg := db.UpdateCompanyParams{ID: id, Name: company.Name, Description: company.Description}
	h.Repository.UpdateCompany(context.Background(), arg)
}

func (h *Handler) DeleteCompany(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		log.Println(err)
		http.Error(w, "Invalid id format", http.StatusBadRequest)
		return
	}

	if err := h.Repository.DeleteCompany(context.Background(), id); err != nil {
		log.Println(err)
		http.Error(w, http.StatusText(404), http.StatusNotFound)
		return
	}
}
