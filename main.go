package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/angelobguido/personal-finances/internal/api"
	"github.com/angelobguido/personal-finances/internal/storage"
	"github.com/angelobguido/personal-finances/internal/utils"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {

	fs := http.FileServer(http.Dir("static"))
	front := http.FileServer(http.Dir("frontend/dist"))

	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, relying on system environment variables.")
	}

	dbConnectionString := utils.GetEnv("DB_CONNECTION_STRING", "postgres://postgres:localpassword@db:5432/postgres?sslmode=disable")

	var err error
	storage.Db, err = sql.Open("postgres", dbConnectionString)
	if err != nil {
		log.Fatalf("FATAL: Error connecting to database (%v): %v", dbConnectionString, err)
	}
	defer storage.Db.Close()

	if err := storage.Db.Ping(); err != nil {
		log.Fatalf("FATAL: Could not ping database (%v): %v", dbConnectionString, err)
	}
	log.Println("Successfully connected to the database.")

	mux := http.NewServeMux()

	// Transaction endpoints
	mux.HandleFunc("GET /transactions/{Id}", api.GetTransactionById)
	mux.HandleFunc("PATCH /transactions/{Id}", api.UpdateTransactionById)
	mux.HandleFunc("DELETE /transactions/{Id}", api.DeleteTransactionById)
	mux.HandleFunc("GET /transactions", api.GetTransactions)
	mux.HandleFunc("POST /transactions", api.CreateTransaction)

	// Category endpoints
	mux.HandleFunc("GET /categories/{Id}", api.GetCategoryById)
	mux.HandleFunc("PATCH /categories/{Id}", api.UpdateCategoryById)
	mux.HandleFunc("DELETE /categories/{Id}", api.DeleteCategoryById)
	mux.HandleFunc("GET /categories", api.GetCategories)
	mux.HandleFunc("POST /categories", api.CreateCategory)

	// Report endpoint
	mux.HandleFunc("GET /report", api.GetReport)

	mux.Handle("/static/", http.StripPrefix("/static/", fs))
	mux.Handle("/", front)

	fmt.Printf("Starting server at port 5000\n")

	log.Fatal(http.ListenAndServe(":5000", mux))
}
