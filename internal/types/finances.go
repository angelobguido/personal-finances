package types

type Finance struct {
	Id       int     `json:"id"`
	Name     string  `json:"name"`
	Category string  `json:"category"`
	Amount   float64 `json:"amount"`
}

type Finances []Finance

type FinanceRequest struct {
	Name     *string  `json:"name"`
	Category *string  `json:"category"`
	Amount   *float64 `json:"amount"`
}
