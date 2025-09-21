package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"slices"
)

// type FinanceType string

// const (
// 	FixedCost        FinanceType = "fixed_cost"
// 	Comfort          FinanceType = "comfort"
// 	Goals            FinanceType = "goals"
// 	Pleasures        FinanceType = "pleasures"
// 	FinancialFreedom FinanceType = "financial_freedom"
// 	Knowledge        FinanceType = "knowledge"
// )

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

func healthCheck(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	err := json.NewEncoder(w).Encode(map[string]string{"message": "Health Check OK"})

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func getFinances(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	err := json.NewEncoder(w).Encode(map[string]Finances{"data": finances})

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

func createFinance(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var financeRequest = FinanceRequest{}

	err := json.NewDecoder(r.Body).Decode(&financeRequest)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if financeRequest.Amount == nil || financeRequest.Name == nil || financeRequest.Type == nil {
		http.Error(w, "All fields are required!", http.StatusBadRequest)
		return
	}

	var finance = Finance{}

	newId := len(finances) + 1
	finance.Id = newId
	finance.Amount = *financeRequest.Amount
	finance.Name = *financeRequest.Name
	finance.Type = *financeRequest.Type

	finances = append(finances, finance)

	err = json.NewEncoder(w).Encode(map[string]Finance{"data": finance})

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func getFinanceById(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	Id := r.PathValue("Id")

	for i, finance := range finances {
		if fmt.Sprintf("%v", finance.Id) == Id {

			err := json.NewEncoder(w).Encode(map[string]Finance{"data": finances[i]})
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			return
		}
	}

	http.NotFound(w, r)
}

func updateFinanceById(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	Id := r.PathValue("Id")

	var financeRequest = FinanceRequest{}

	err := json.NewDecoder(r.Body).Decode(&financeRequest)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	for i, finance := range finances {
		if fmt.Sprintf("%v", finance.Id) == Id {
			if financeRequest.Name != nil {
				finances[i].Name = *financeRequest.Name
			}
			if financeRequest.Type != nil {
				finances[i].Type = *financeRequest.Type
			}
			if financeRequest.Amount != nil {
				finances[i].Amount = *financeRequest.Amount
			}

			err := json.NewEncoder(w).Encode(map[string]Finance{"data": finances[i]})
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			return
		}
	}

	http.NotFound(w, r)

}

func deleteFinanceById(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	id := r.PathValue("Id")

	for i, finance := range finances {
		if fmt.Sprintf("%v", finance.Id) == id {

			finances = slices.Delete(finances, i, i+1)
			err := json.NewEncoder(w).Encode(map[string]string{"message": fmt.Sprintf("Finance with id %v deleted!", id)})

			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			return
		}
	}

	http.NotFound(w, r)

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
