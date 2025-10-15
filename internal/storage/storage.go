package storage

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/angelobguido/personal-finances/internal/types"
)

var Db *sql.DB

func GetFinances() ([]types.Finance, error) {
	finances := []types.Finance{}

	rows, err := Db.Query("SELECT id, name, category, amount, created_at FROM finance ORDER BY id")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var finance types.Finance
		if err := rows.Scan(&finance.Id, &finance.Name, &finance.Category, &finance.Amount, &finance.CreatedAt); err != nil {
			return nil, err
		}
		finances = append(finances, finance)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return finances, err
}

func CreateFinance(name string, amount float64, category string, createdAt time.Time) (*types.Finance, error) {

	var finance = types.Finance{}

	if err := Db.QueryRow("INSERT INTO finance(name, amount, category, created_at) VALUES ($1, $2, $3, $4) RETURNING id, name, category, amount, created_at", name, amount, category, createdAt).Scan(&finance.Id, &finance.Name, &finance.Category, &finance.Amount, &finance.CreatedAt); err != nil {
		return nil, err
	}

	return &finance, nil
}

func GetFinanceById(id string) (*types.Finance, error) {

	finance := types.Finance{}

	if err := Db.QueryRow("SELECT id, name, category, amount, created_at FROM finance WHERE id=$1", id).Scan(&finance.Id, &finance.Name, &finance.Category, &finance.Amount, &finance.CreatedAt); err != nil {
		return nil, err
	}

	return &finance, nil
}

func UpdateFinanceById(id string, name *string, amount *float64, category *string, createdAt *time.Time) (*types.Finance, error) {

	finance := types.Finance{}

	if err := Db.QueryRow("UPDATE finance SET name = COALESCE($1, name), category = COALESCE($2, category), amount = COALESCE($3, amount), created_at = COALESCE($4, created_at) WHERE id=$5 RETURNING id, name, category, amount, created_at", name, category, amount, createdAt, id).Scan(&finance.Id, &finance.Name, &finance.Category, &finance.Amount, &finance.CreatedAt); err != nil {
		fmt.Printf("Error %v", err.Error())
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

func GetCategoriesReport() ([]types.CategorySummary, error) {
	categories := []types.CategorySummary{}

	rows, err := Db.Query("SELECT SUM(amount), category as total FROM finance GROUP BY category")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var total float64
		var category string
		if err := rows.Scan(&total, &category); err != nil {
			return nil, err
		}
		categories = append(categories, types.CategorySummary{Category: category, Total: total})
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return categories, nil
}
