package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello\n")
}
func main() {
	db, err := sql.Open("postgres", "postgresql://test:test@postgres:5432/simple?sslmode=disable")
	if err != nil {
		log.Fatal("cannot establish connection to database: ", err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal("cannot establish connection to database: ", err)
	}
	http.HandleFunc("/hello", hello)
	http.ListenAndServe(":3000", nil)
}
