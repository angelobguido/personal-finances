package storage

import (
	"database/sql"

	"github.com/angelobguido/personal-finances/internal/types"
)

var Db *sql.DB

func GetFinances() ([]types.Finance, error) {
	finances := []types.Finance{}

	rows, err := Db.Query("SELECT id, name, category, amount FROM finance ORDER BY id")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var finance types.Finance
		if err := rows.Scan(&finance.Id, &finance.Name, &finance.Category, &finance.Amount); err != nil {
			return nil, err
		}
		finances = append(finances, finance)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return finances, err
}

func CreateFinance(name string, amount float64, category string) (*types.Finance, error) {

	var finance = types.Finance{}

	if err := Db.QueryRow("INSERT INTO finance(name, amount, category) VALUES ($1, $2, $3) RETURNING id, name, category, amount", name, amount, category).Scan(&finance.Id, &finance.Name, &finance.Category, &finance.Amount); err != nil {
		return nil, err
	}

	return &finance, nil
}

func GetFinanceById(id string) (*types.Finance, error) {

	finance := types.Finance{}

	if err := Db.QueryRow("SELECT id, name, category, amount FROM finance WHERE id=$1", id).Scan(&finance.Id, &finance.Name, &finance.Category, &finance.Amount); err != nil {
		return nil, err
	}

	return &finance, nil
}

func UpdateFinanceById(id string, name *string, amount *float64, category *string) (*types.Finance, error) {

	finance := types.Finance{}

	if err := Db.QueryRow("UPDATE finance SET name = COALESCE($1, name), category = COALESCE($2, category), amount = COALESCE($3, amount) WHERE id=$4 RETURNING id, name, category, amount", name, category, amount, id).Scan(&finance.Id, &finance.Name, &finance.Category, &finance.Amount); err != nil {
		return nil, err
	}

	return &finance, nil

}

func DeleteFinanceById(id string) error {

	if err := Db.QueryRow("DELETE FROM finance WHERE id=$1", id).Err(); err != nil {
		return err
	}

	return nil

}
