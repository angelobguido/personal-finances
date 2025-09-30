package api

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/angelobguido/personal-finances/internal/storage"
	"github.com/angelobguido/personal-finances/internal/types"
	"github.com/angelobguido/personal-finances/internal/utils"
	_ "github.com/lib/pq"
)

func HealthCheck(w http.ResponseWriter, r *http.Request) {

	utils.Encode(w, &map[string]string{"message": "The server is running!"}, http.StatusOK)

}

func GetFinancesList() ([]types.Finance, error) {
	finances := []types.Finance{}

	rows, err := storage.Db.Query("SELECT id, name, type, amount FROM finance")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var finance types.Finance
		if err := rows.Scan(&finance.Id, &finance.Name, &finance.Type, &finance.Amount); err != nil {
			return nil, err
		}
		finances = append(finances, finance)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return finances, err
}

func GetFinances(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	finances := []types.Finance{}

	rows, err := storage.Db.Query("SELECT id, name, type, amount FROM finance")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	defer rows.Close()

	for rows.Next() {
		var finance types.Finance
		if err := rows.Scan(&finance.Id, &finance.Name, &finance.Type, &finance.Amount); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		finances = append(finances, finance)
	}

	if err := rows.Err(); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	utils.Encode(w, &map[string]types.Finances{"data": finances}, http.StatusOK)

}

func CreateFinance(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")

	financeRequest, err := utils.Decode[types.FinanceRequest](r)

	if err != nil {
		utils.Encode(w, &map[string]string{"error": "The body structure is wrong!"}, http.StatusBadRequest)
		return
	}

	if financeRequest.Amount == nil || financeRequest.Name == nil || financeRequest.Type == nil {
		utils.Encode(w, &map[string]string{"error": "All fields are required!"}, http.StatusBadRequest)
		return
	}

	var finance = types.Finance{}

	if err := storage.Db.QueryRow("INSERT INTO finance(name, amount, type) VALUES ($1, $2, $3) RETURNING id, name, type, amount", *financeRequest.Name, *financeRequest.Amount, *financeRequest.Type).Scan(&finance.Id, &finance.Name, &finance.Type, &finance.Amount); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	utils.Encode(w, &map[string]types.Finance{"data": finance}, http.StatusCreated)
}

func GetFinanceById(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")

	id := r.PathValue("Id")

	finance := types.Finance{}

	if err := storage.Db.QueryRow("SELECT id, name, type, amount FROM finance WHERE id=$1", id).Scan(&finance.Id, &finance.Name, &finance.Type, &finance.Amount); err != nil {

		if err == sql.ErrNoRows {
			utils.Encode(w, &map[string]string{"error": fmt.Sprintf("Finance with id %v doesn't exist!", id)}, http.StatusNotFound)
			return
		}

		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	utils.Encode(w, &map[string]types.Finance{"data": finance}, http.StatusOK)
}

func UpdateFinanceById(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")

	id := r.PathValue("Id")

	financeRequest, err := utils.Decode[types.FinanceRequest](r)

	if err != nil {
		utils.Encode(w, &map[string]string{"error": "The body structure is wrong!"}, http.StatusBadRequest)
		return
	}

	finance := types.Finance{}

	if err := storage.Db.QueryRow("UPDATE finance SET name = COALESCE($1, name), type = COALESCE($2, type), amount = COALESCE($3, amount) WHERE id=$4 RETURNING id, name, type, amount", financeRequest.Name, financeRequest.Type, financeRequest.Amount, id).Scan(&finance.Id, &finance.Name, &finance.Type, &finance.Amount); err != nil {

		if err == sql.ErrNoRows {
			utils.Encode(w, &map[string]string{"error": fmt.Sprintf("Finance with id %v doesn't exist!", id)}, http.StatusNotFound)
			return
		}

		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	utils.Encode(w, &map[string]types.Finance{"data": finance}, http.StatusOK)

}

func DeleteFinanceById(w http.ResponseWriter, r *http.Request) {

	id := r.PathValue("Id")

	if err := storage.Db.QueryRow("DELETE FROM finance WHERE id=$1", id).Err(); err != nil {

		if err == sql.ErrNoRows {
			w.WriteHeader(http.StatusNoContent)
			w.Header().Set("Access-Control-Allow-Origin", "*")

			return
		}

		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Set("Access-Control-Allow-Origin", "*")

		return
	}

	w.WriteHeader(http.StatusNoContent)
	w.Header().Set("Access-Control-Allow-Origin", "*")

}
