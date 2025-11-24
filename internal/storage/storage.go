package storage

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/angelobguido/personal-finances/internal/types"
)

var Db *sql.DB

func GetTransactions() ([]types.Transaction, error) {
	transactions := []types.Transaction{}

	rows, err := Db.Query("SELECT id, name, category_id, amount, created_at, data FROM transaction ORDER BY created_at DESC")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var transaction types.Transaction
		if err := rows.Scan(&transaction.Id, &transaction.Name, &transaction.CategoryId, &transaction.Amount, &transaction.CreatedAt, &transaction.Data); err != nil {
			return nil, err
		}
		transactions = append(transactions, transaction)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return transactions, err
}

func CreateTransaction(data types.TransactionCreateData) (*types.Transaction, error) {
	var transaction = types.Transaction{}

	if err := Db.QueryRow("INSERT INTO transaction(name, amount, category_id, created_at, data) VALUES ($1, $2, $3, $4, $5) RETURNING id, name, category_id, amount, created_at, data", data.Name, data.Amount, data.CategoryId, data.CreatedAt, data.Data).Scan(&transaction.Id, &transaction.Name, &transaction.CategoryId, &transaction.Amount, &transaction.CreatedAt, &transaction.Data); err != nil {
		return nil, err
	}

	return &transaction, nil
}

func GetTransactionById(id int) (*types.Transaction, error) {

	transaction := types.Transaction{}

	if err := Db.QueryRow("SELECT id, name, category_id, amount, created_at, data FROM transaction WHERE id=$1", id).Scan(&transaction.Id, &transaction.Name, &transaction.CategoryId, &transaction.Amount, &transaction.CreatedAt, &transaction.Data); err != nil {
		return nil, err
	}

	return &transaction, nil
}

func UpdateTransactionById(id int, data types.TransactionUpdateData) (*types.Transaction, error) {

	transaction := types.Transaction{}

	if err := Db.QueryRow("UPDATE transaction SET name = COALESCE($1, name), category_id = COALESCE($2, category_id), amount = COALESCE($3, amount), created_at = COALESCE($4, created_at), data = COALESCE($5, data) WHERE id=$6 RETURNING id, name, category_id, amount, created_at, data", data.Name, data.CategoryId, data.Amount, data.CreatedAt, data.Data, id).Scan(&transaction.Id, &transaction.Name, &transaction.CategoryId, &transaction.Amount, &transaction.CreatedAt, &transaction.Data); err != nil {
		fmt.Printf("Error %v", err.Error())
		return nil, err
	}

	return &transaction, nil

}

func DeleteTransactionById(id int) error {

	if err := Db.QueryRow("DELETE FROM transaction WHERE id=$1", id).Err(); err != nil {
		return err
	}

	return nil

}

func GetCategories() ([]types.Category, error) {

	categories := []types.Category{}
	rows, err := Db.Query("SELECT id, name, is_income, data, created_at FROM category ORDER BY name ASC")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var category types.Category
		if err := rows.Scan(&category.Id, &category.Name, &category.IsIncome, &category.Data, &category.CreatedAt); err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return categories, nil

}

func CreateCategory(data types.CategoryCreateData) (*types.Category, error) {

	category := types.Category{}
	if err := Db.QueryRow("INSERT INTO category(name, is_income, data) VALUES ($1, $2, $3) RETURNING id, name, is_income, data, created_at", data.Name, data.IsIncome, data.Data).Scan(&category.Id, &category.Name, &category.IsIncome, &category.Data, &category.CreatedAt); err != nil {
		return nil, err
	}
	return &category, nil
}

func GetCategoryById(id int) (*types.Category, error) {

	category := types.Category{}
	if err := Db.QueryRow("SELECT id, name, is_income, data, created_at FROM category WHERE id=$1", id).Scan(&category.Id, &category.Name, &category.IsIncome, &category.Data, &category.CreatedAt); err != nil {
		return nil, err
	}

	return &category, nil

}

func UpdateCategoryById(id int, data types.CategoryUpdateData) (*types.Category, error) {

	category := types.Category{}
	if err := Db.QueryRow("UPDATE category SET name = COALESCE($1, name), is_income = COALESCE($2, is_income), data = COALESCE($3, data) WHERE id=$4 RETURNING id, name, is_income, data, created_at", data.Name, data.IsIncome, data.Data, id).Scan(&category.Id, &category.Name, &category.IsIncome, &category.Data, &category.CreatedAt); err != nil {
		return nil, err
	}

	return &category, nil

}

func DeleteCategoryById(id int) error {

	if err := Db.QueryRow("DELETE FROM category WHERE id=$1", id).Err(); err != nil {
		return err
	}

	return nil
}

func GetReport(start time.Time, end time.Time) (*types.Report, error) {

	categories := []types.CategorySummary{}

	rows, err := Db.Query("SELECT SUM(t.amount) as total, c.id, c.name, c.is_income FROM transaction t JOIN category c ON t.category_id = c.id WHERE t.created_at >= $1 AND t.created_at <= $2 GROUP BY c.id, c.name, c.is_income", start, end)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		category := types.CategorySummary{}

		if err := rows.Scan(&category.Total, &category.Id, &category.Name, &category.IsIncome); err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	report := types.Report{}
	report.Categories = categories

	return &report, nil
}
