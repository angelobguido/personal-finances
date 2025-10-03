package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/angelobguido/personal-finances/internal/api"
	"github.com/angelobguido/personal-finances/internal/renderer"
	"github.com/angelobguido/personal-finances/internal/storage"
	"github.com/angelobguido/personal-finances/internal/utils"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {

	renderer.Templates = template.Must(template.ParseGlob("templates/*.html"))

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

	mux.HandleFunc("GET /{$}", api.HealthCheck)
	mux.HandleFunc("GET /finances/{Id}", api.GetFinanceById)
	mux.HandleFunc("PATCH /finances/{Id}", api.UpdateFinanceById)
	mux.HandleFunc("DELETE /finances/{Id}", api.DeleteFinanceById)
	mux.HandleFunc("GET /finances", api.GetFinances)
	mux.HandleFunc("POST /finances", api.CreateFinance)

	mux.HandleFunc("GET /home", renderer.RenderHome)
	mux.HandleFunc("POST /render/finances", renderer.CreateFinance)
	mux.HandleFunc("GET /render/finances/edit/{Id}", renderer.RenderEditFinance)
	mux.HandleFunc("PATCH /render/finances/{Id}", renderer.UpdateFinance)
	mux.HandleFunc("GET /render/finances/{Id}", renderer.RenderFinance)

	fmt.Printf("Starting server at port 8090\n")

	log.Fatal(http.ListenAndServe(":8090", mux))
}
