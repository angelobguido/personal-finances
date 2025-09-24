package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"slices"

	_ "github.com/lib/pq"
)

var dbConnectionString = getEnv("DB_CONNECTION_STRING", "postgres://postgres:localpassword@db:5432/postgres?sslmode=disable")

var db *sql.DB

type Finance struct {
	Id     int     `json:"id"`
	Name   string  `json:"name"`
	Type   string  `json:"type"`
	Amount float64 `json:"amount"`
}

type Finances []Finance

type FinanceRequest struct {
	Name   *string  `json:"name"`
	Type   *string  `json:"type"`
	Amount *float64 `json:"amount"`
}

var finances = Finances{
	{1, "Aluguel", "FixedCost", 373.94},
	{2, "Crunchyroll", "Pleasures", 14.99},
	{3, "Investimento em Tesouro Direto", "FinancialFreedom", 1200.00},
}

func encode[T any](w http.ResponseWriter, v *T, status int) error {

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	if err := json.NewEncoder(w).Encode(*v); err != nil {
		return fmt.Errorf("encode json: %w", err)
	}
	return nil

}

func decode[T any](r *http.Request) (T, error) {

	var v T
	if err := json.NewDecoder(r.Body).Decode(&v); err != nil {
		return v, fmt.Errorf("decode json: %w", err)
	}
	return v, nil

}

func healthCheck(w http.ResponseWriter, r *http.Request) {

	encode(w, &map[string]string{"message": "The server is running!"}, http.StatusOK)

}

func getFinances(w http.ResponseWriter, r *http.Request) {

	finances := []Finance{}

	rows, err := db.Query("SELECT id, name, type, amount FROM finance")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	defer rows.Close()

	for rows.Next() {
		var finance Finance
		if err := rows.Scan(&finance.Id, &finance.Name, &finance.Type, &finance.Amount); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		finances = append(finances, finance)
	}

	if err := rows.Err(); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	encode(w, &map[string]Finances{"data": finances}, http.StatusOK)

}

func createFinance(w http.ResponseWriter, r *http.Request) {

	financeRequest, err := decode[FinanceRequest](r)

	if err != nil {
		encode(w, &map[string]string{"error": "The body structure is wrong!"}, http.StatusBadRequest)
		return
	}

	if financeRequest.Amount == nil || financeRequest.Name == nil || financeRequest.Type == nil {
		encode(w, &map[string]string{"error": "All fields are required!"}, http.StatusBadRequest)
		return
	}

	var finance = Finance{}

	if err := db.QueryRow("INSERT INTO finance(name, amount, type) VALUES ($1, $2, $3) RETURNING id, name, type, amount", *financeRequest.Name, *financeRequest.Amount, *financeRequest.Type).Scan(&finance.Id, &finance.Name, &finance.Type, &finance.Amount); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	encode(w, &map[string]Finance{"data": finance}, http.StatusCreated)
}

func getFinanceById(w http.ResponseWriter, r *http.Request) {

	id := r.PathValue("Id")

	finance := Finance{}

	if err := db.QueryRow("SELECT id, name, type, amount FROM finance WHERE id=$1", id).Scan(&finance.Id, &finance.Name, &finance.Type, &finance.Amount); err != nil {

		if err == sql.ErrNoRows {
			encode(w, &map[string]string{"error": fmt.Sprintf("Finance with id %v doesn't exist!", id)}, http.StatusNotFound)
			return
		}

		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	encode(w, &map[string]Finance{"data": finance}, http.StatusOK)
}

func updateFinanceById(w http.ResponseWriter, r *http.Request) {

	id := r.PathValue("Id")

	financeRequest, err := decode[FinanceRequest](r)

	if err != nil {
		encode(w, &map[string]string{"error": "The body structure is wrong!"}, http.StatusBadRequest)
		return
	}

	for i, finance := range finances {
		if fmt.Sprintf("%v", finance.Id) == id {
			if financeRequest.Name != nil {
				finances[i].Name = *financeRequest.Name
			}
			if financeRequest.Type != nil {
				finances[i].Type = *financeRequest.Type
			}
			if financeRequest.Amount != nil {
				finances[i].Amount = *financeRequest.Amount
			}

			encode(w, &map[string]Finance{"data": finances[i]}, http.StatusOK)
			return
		}
	}

	encode(w, &map[string]string{"error": fmt.Sprintf("Finance with id %v doesn't exist!", id)}, http.StatusNotFound)

}

func deleteFinanceById(w http.ResponseWriter, r *http.Request) {

	id := r.PathValue("Id")

	for i, finance := range finances {
		if fmt.Sprintf("%v", finance.Id) == id {

			finances = slices.Delete(finances, i, i+1)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}

	w.WriteHeader(http.StatusNoContent)

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

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
