package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type FinanceType string

const (
	FixedCost        FinanceType = "fixed_cost"
	Comfort          FinanceType = "comfort"
	Goals            FinanceType = "goals"
	Pleasures        FinanceType = "pleasures"
	FinancialFreedom FinanceType = "financial_freedom"
	Knowledge        FinanceType = "knowledge"
)

type Finance struct {
	Id     int         `json:"id"`
	Name   string      `json:"name"`
	Type   FinanceType `json:"type"`
	Amount float64     `json:"amount"`
}

type Finances []Finance

var finances = Finances{
	{1, "Aluguel", FixedCost, 373.94},
	{2, "Crunchyroll", Pleasures, 14.99},
	{3, "Investimento em Tesouro Direto", FinancialFreedom, 1200.00},
}

func healthCheck(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Health Check OK\n")
}

func getFinances(w http.ResponseWriter, r *http.Request) {

	for _, finance := range finances {
		fmt.Fprintf(w, "ID: %v, Name: %v, Type: %v, Amount: %v\n", finance.Id, finance.Name, finance.Type, finance.Amount)
	}
}

func createFinance(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var finance = Finance{}

	err := json.NewDecoder(r.Body).Decode(&finance)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	newId := len(finances) + 1
	finance.Id = newId

	finances = append(finances, finance)

	err = json.NewEncoder(w).Encode(&finance)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func getFinanceById(w http.ResponseWriter, r *http.Request) {

	Id := r.PathValue("Id")

	for _, finance := range finances {
		if fmt.Sprintf("%v", finance.Id) == Id {
			fmt.Fprintf(w, "ID: %v, Name: %v, Type: %v, Amount: %v\n", finance.Id, finance.Name, finance.Type, finance.Amount)
			return
		}
	}

	http.NotFound(w, r)
}

// func updateFinanceById(w http.ResponseWriter, r *http.Request) {

// 	Id := r.PathValue("Id")

// 	for _, finance := range finances {
// 		if fmt.Sprintf("%v", finance.Id) == Id {
// 			fmt.Fprintf(w, "ID: %v, Name: %v, Type: %v, Amount: %v\n", finance.Id, finance.Name, finance.Type, finance.Amount)
// 			return
// 		}
// 	}

// 	http.NotFound(w, r)

// }

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("GET /", healthCheck)
	mux.HandleFunc("GET /finances/{Id}/", getFinanceById)
	//mux.HandleFunc("PATCH /finances/{Id}/", updateFinanceById)
	mux.HandleFunc("GET /finances/", getFinances)
	mux.HandleFunc("POST /finances/", createFinance)

	fmt.Printf("Starting server at port 8090\n")

	log.Fatal(http.ListenAndServe(":8090", mux))
}
