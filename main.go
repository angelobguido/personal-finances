package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

var dbConnectionString = getEnv("DB_CONNECTION_STRING", "postgres://postgres:localpassword@db:5432/postgres?sslmode=disable")

var db *sql.DB

func healthCheck(w http.ResponseWriter, r *http.Request) {

	encode(w, &map[string]string{"message": "The server is running!"}, http.StatusOK)

}

func main() {

	var err error
	db, err = sql.Open("postgres", dbConnectionString)
	if err != nil {
		log.Fatalf("FATAL: Error connecting to database: %v", err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatalf("FATAL: Could not ping database: %v", err)
	}
	log.Println("Successfully connected to the database.")

	mux := http.NewServeMux()

	mux.HandleFunc("GET /{$}", healthCheck)
	mux.HandleFunc("GET /finances/{Id}", getFinanceById)
	mux.HandleFunc("PATCH /finances/{Id}", updateFinanceById)
	mux.HandleFunc("DELETE /finances/{Id}", deleteFinanceById)
	mux.HandleFunc("GET /finances", getFinances)
	mux.HandleFunc("POST /finances", createFinance)

	fmt.Printf("Starting server at port 8090\n")

	log.Fatal(http.ListenAndServe(":8090", mux))
}
