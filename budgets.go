package main

import (
	"encoding/json"
	"net/http"

	"github.com/azbshiri/common/db"
	"github.com/pquerna/ffjson/ffjson"
)

type budget struct {
	db.Model
	Amount float64 `json:"amount"`
}

func (s *server) getBudgets(w http.ResponseWriter, r *http.Request) {
	var budgets []*budget
	err := s.db.Table(&budget).Select()
	if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(DatabaseError)
			return
	}

	ffjson.NewEncoder(w).Encode(budgets)
}


func (s *server) createBudget(w http.ResponseWriter, r *http.Request) {
	var param struct {
			Amount float64 `json:"amount"`
	}
	err := ffjson.NewDecoder().DecodeReader(r.Body, &param)
	if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(BadParamError)
			return
	}

	budget := budget{Amount: param.Amount}
	err = s.db.Insert(&budget)
	if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(DatabaseError)
			return
	}

	ffjson.NewEncoder(w).Encode(budget)
}