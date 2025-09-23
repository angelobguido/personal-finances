package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"slices"
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

	newId := len(finances) + 1
	finance.Id = newId
	finance.Amount = *financeRequest.Amount
	finance.Name = *financeRequest.Name
	finance.Type = *financeRequest.Type

	finances = append(finances, finance)

	encode(w, &map[string]Finance{"data": finance}, http.StatusCreated)
}

func getFinanceById(w http.ResponseWriter, r *http.Request) {

	id := r.PathValue("Id")

	for i, finance := range finances {
		if fmt.Sprintf("%v", finance.Id) == id {
			encode(w, &map[string]Finance{"data": finances[i]}, http.StatusOK)
			return
		}
	}

	encode(w, &map[string]string{"error": fmt.Sprintf("Finance with id %v doesn't exist!", id)}, http.StatusNotFound)

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
