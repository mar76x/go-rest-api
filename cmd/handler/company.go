package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	db "github.com/mar76x/go-rest-api/db/sqlc"
)

func (h *Handler) ListCompanies(w http.ResponseWriter, r *http.Request) {
	args := db.ListCompaniesParams{Limit: 10, Offset: 2}
	companyList, err := h.Repository.ListCompanies(context.Background(), args)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "unable to retrive data from database", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(companyList); err != nil {
		http.Error(w, "unable to encode json", http.StatusInternalServerError)
		return
	}
}

func (h *Handler) CreateCompany(w http.ResponseWriter, r *http.Request) {
	var company db.Company
	if err := json.NewDecoder(r.Body).Decode(&company); err != nil {
		log.Println("error: ", err)
		http.Error(w, "error decoding json body", http.StatusBadRequest)
		return
	}
	// company := db.CreateCompanyParams{company.Name}
	// h.Repository.CreateCompany(context.Background())
}
