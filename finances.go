package main

import (
	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/lib/pq"
)

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

func getFinances(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

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

	w.Header().Set("Access-Control-Allow-Origin", "*")

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

	w.Header().Set("Access-Control-Allow-Origin", "*")

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

	w.Header().Set("Access-Control-Allow-Origin", "*")

	id := r.PathValue("Id")

	financeRequest, err := decode[FinanceRequest](r)

	if err != nil {
		encode(w, &map[string]string{"error": "The body structure is wrong!"}, http.StatusBadRequest)
		return
	}

	finance := Finance{}

	if err := db.QueryRow("UPDATE finance SET name = COALESCE($1, name), type = COALESCE($2, type), amount = COALESCE($3, amount) WHERE id=$4 RETURNING id, name, type, amount", financeRequest.Name, financeRequest.Type, financeRequest.Amount, id).Scan(&finance.Id, &finance.Name, &finance.Type, &finance.Amount); err != nil {

		if err == sql.ErrNoRows {
			encode(w, &map[string]string{"error": fmt.Sprintf("Finance with id %v doesn't exist!", id)}, http.StatusNotFound)
			return
		}

		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	encode(w, &map[string]Finance{"data": finance}, http.StatusOK)

}

func deleteFinanceById(w http.ResponseWriter, r *http.Request) {

	id := r.PathValue("Id")

	if err := db.QueryRow("DELETE FROM finance WHERE id=$1", id).Err(); err != nil {

		if err == sql.ErrNoRows {
			w.WriteHeader(http.StatusNoContent)
			w.Header().Set("Access-Control-Allow-Origin", "*")

			return
		}

		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Set("Access-Control-Allow-Origin", "*")

		return
	}

	w.WriteHeader(http.StatusNoContent)
	w.Header().Set("Access-Control-Allow-Origin", "*")

}
