package main

import (
	"encoding/json"
	"fmt"
	"io"
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

func health_check(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Health Check OK\n")
}

func get_all_finances(w http.ResponseWriter, r *http.Request) {

	for _, finance := range finances {
		fmt.Fprintf(w, "ID: %v, Name: %v, Type: %v, Amount: %v\n", finance.Id, finance.Name, finance.Type, finance.Amount)
	}
}

func create_finance(w http.ResponseWriter, r *http.Request) {

	defer r.Body.Close()

	var finance = Finance{}

	body, _ := io.ReadAll(r.Body)
	if err := json.Unmarshal(body, &finance); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	newId := len(finances) + 1
	finance.Id = newId

	fmt.Fprintf(w, "ID: %v, Name: %v, Type: %v, Amount: %v\n", finance.Id, finance.Name, finance.Type, finance.Amount)

	finances = append(finances, finance)
}

func get_finance_by_id(w http.ResponseWriter, r *http.Request) {

	Id := r.PathValue("Id")

	for _, finance := range finances {
		if fmt.Sprintf("%v", finance.Id) == Id {
			fmt.Fprintf(w, "ID: %v, Name: %v, Type: %v, Amount: %v\n", finance.Id, finance.Name, finance.Type, finance.Amount)
			return
		}
	}

	http.NotFound(w, r)
}

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("GET /", health_check)
	mux.HandleFunc("GET /finances/{Id}/", get_finance_by_id)
	mux.HandleFunc("GET /finances/", get_all_finances)
	mux.HandleFunc("POST /finances/", create_finance)

	fmt.Printf("Starting server at port 8090\n")

	log.Fatal(http.ListenAndServe(":8090", mux))
}
