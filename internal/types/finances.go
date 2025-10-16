package types

import "time"

type Finance struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	Category  string    `json:"category"`
	Amount    float64   `json:"amount"`
	CreatedAt time.Time `json:"created_at"`
}

type Finances []Finance

type FinanceRequest struct {
	Name      *string    `json:"name"`
	Category  *string    `json:"category"`
	Amount    *float64   `json:"amount"`
	CreatedAt *time.Time `json:"created_at"`
}

type Report struct {
	TotalIncome     float64          `json:"total_income"`
	TotalExpense    float64          `json:"total_expense"`
	NetTotal        float64          `json:"net_total"`
	ExpensesSummary []ExpenseSummary `json:"expenses_summary"`
}

type ExpenseSummary struct {
	Category     string  `json:"category"`
	Total        float64 `json:"total"`
	ExpenseRatio float64 `json:"expense_ratio"`
}

type CategoryTotal struct {
	Category string  `json:"category"`
	Total    float64 `json:"total"`
}
