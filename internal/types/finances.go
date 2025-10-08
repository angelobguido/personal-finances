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
