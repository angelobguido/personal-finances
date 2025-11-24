package types

import "time"

type Transaction struct {
	Id         int       `json:"id"`
	Name       string    `json:"name"`
	CategoryId int       `json:"category_id"`
	Amount     float64   `json:"amount"`
	Data       *any      `json:"data"`
	CreatedAt  time.Time `json:"created_at"`
}

type Category struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	IsIncome  bool      `json:"is_income"`
	Data      *any      `json:"data"`
	CreatedAt time.Time `json:"created_at"`
}

type TransactionCreateData struct {
	Name       string    `json:"name"`
	CategoryId int       `json:"category_id"`
	Amount     float64   `json:"amount"`
	CreatedAt  time.Time `json:"created_at"`
	Data       *any      `json:"data"`
}

type TransactionUpdateData struct {
	Name       *string    `json:"name"`
	CategoryId *int       `json:"category_id"`
	Amount     *float64   `json:"amount"`
	CreatedAt  *time.Time `json:"created_at"`
	Data       *any       `json:"data"`
}

type CategoryCreateData struct {
	Name     string `json:"name"`
	IsIncome bool   `json:"is_income"`
	Data     *any   `json:"data"`
}

type CategoryUpdateData struct {
	Name     *string `json:"name"`
	IsIncome *bool   `json:"is_income"`
	Data     *any    `json:"data"`
}

type Report struct {
	Categories []CategorySummary `json:"categories"`
}

type CategorySummary struct {
	Id       int     `json:"id"`
	Name     string  `json:"category"`
	IsIncome bool    `json:"is_income"`
	Total    float64 `json:"total"`
}
