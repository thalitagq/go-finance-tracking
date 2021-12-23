package main

import (
	"net/http"
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
	"github.com/gorilla/mux"
)

func main() {
	
	fmt.Println("connecting........")
	// these details match the docker-compose.yml file.
	postgresInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		"postgres", 5432, "user", "password", "budget")
	db, err := sql.Open("postgres", postgresInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	start := time.Now()
	for db.Ping() != nil {
		if start.After(start.Add(10 * time.Second)) {
			fmt.Println("failed to connect after 10 secs.")
			break
		}
	}
	fmt.Println("connected:", db.Ping() == nil)

	mux := mux.NewRouter()
	server := newServer(db, mux)
	http.ListenAndServe(":8080", server)
}