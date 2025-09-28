package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var db *sql.DB

func healthCheck(w http.ResponseWriter, r *http.Request) {

	encode(w, &map[string]string{"message": "The server is running!"}, http.StatusOK)

}

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func main() {

	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, relying on system environment variables.")
	}

	dbConnectionString := getEnv("DB_CONNECTION_STRING", "postgres://postgres:localpassword@db:5432/postgres?sslmode=disable")

	var err error
	db, err = sql.Open("postgres", dbConnectionString)
	if err != nil {
		log.Fatalf("FATAL: Error connecting to database (%v): %v", dbConnectionString, err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatalf("FATAL: Could not ping database (%v): %v", dbConnectionString, err)
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

	log.Fatal(http.ListenAndServe(":8090", corsMiddleware(mux)))
}
