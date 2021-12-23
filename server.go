package main

import (
	"net/http"
	"fmt"
	"database/sql"
	"github.com/gorilla/mux"
)

type server struct {
	db *sql.DB
	mux *mux.Router
}

func newServer(db *sql.DB, mux *mux.Router) *server {
	fmt.Println("newServer")
	s := server{db, mux}
	s.routes() // register handlers.
	fmt.Println("newServer ok")
	return &s
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.mux.ServeHTTP(w, r)
}