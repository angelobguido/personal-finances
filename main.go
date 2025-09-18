package main

import (
	"fmt"
	"log"
	"net/http"
)

type FinanceType string

const (
	Investment FinanceType = "investment"
	Expense    FinanceType = "expense"
)

type Finance struct {
	id     int
	name   string
	f_type FinanceType
	amount float64
}

type Finances []Finance

var finances = Finances{
	{1, "Rent", Expense, 1500.00},
	{2, "Groceries", Expense, 300.00},
	{3, "Investment", Investment, 200.00},
	{4, "Utilities", Expense, 200.00},
	{5, "Entertainment", Expense, 150.00},
}

func health_check(w http.ResponseWriter, req *http.Request) {

	fmt.Fprintf(w, "Health Check OK\n")
}

func get_all_finances(w http.ResponseWriter, req *http.Request) {

	for _, finance := range finances {
		fmt.Fprintf(w, "ID: %v, Name: %v, Type: %v, Amount: %v\n", finance.id, finance.name, finance.f_type, finance.amount)
	}
}

func get_finance_by_id(w http.ResponseWriter, req *http.Request) {

	id := req.PathValue("id")

	for _, finance := range finances {
		if fmt.Sprintf("%v", finance.id) == id {
			fmt.Fprintf(w, "ID: %v, Name: %v, Type: %v, Amount: %v\n", finance.id, finance.name, finance.f_type, finance.amount)
			return
		}
	}

	http.NotFound(w, req)
}

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("GET /", health_check)
	mux.HandleFunc("GET /finance/{id}/", get_finance_by_id)
	mux.HandleFunc("GET /finance/", get_all_finances)

	fmt.Printf("Starting server at port 8090\n")

	log.Fatal(http.ListenAndServe(":8090", mux))
}
