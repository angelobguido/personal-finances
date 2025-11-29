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

func GetTransactions(w http.ResponseWriter, r *http.Request) {

	transactions, err := storage.GetTransactions()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	utils.Encode(w, &map[string][]types.Transaction{"data": transactions}, http.StatusOK)

}

func CreateTransaction(w http.ResponseWriter, r *http.Request) {

	transactionRequest, err := utils.Decode[types.TransactionRequestData](r)

	if err != nil {
		utils.Encode(w, &map[string]string{"error": "The body structure is wrong!"}, http.StatusBadRequest)
		return
	}

	transaction, err := storage.CreateTransaction(transactionRequest)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	utils.Encode(w, transaction, http.StatusCreated)
}

func GetTransactionById(w http.ResponseWriter, r *http.Request) {

	id := r.PathValue("Id")

	transaction, err := storage.GetTransactionById(id)

	if err != nil {

		if err == sql.ErrNoRows {
			utils.Encode(w, &map[string]string{"error": fmt.Sprintf("Transaction with id %v doesn't exist!", id)}, http.StatusNotFound)
			return
		}

		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	utils.Encode(w, transaction, http.StatusOK)
}

func UpdateTransactionById(w http.ResponseWriter, r *http.Request) {

	id := r.PathValue("Id")

	transactionRequest, err := utils.Decode[types.TransactionRequestData](r)

	if err != nil {
		utils.Encode(w, &map[string]string{"error": "The body structure is wrong!"}, http.StatusBadRequest)
		return
	}

	transaction, err := storage.UpdateTransactionById(id, transactionRequest)

	if err != nil {

		if err == sql.ErrNoRows {
			utils.Encode(w, &map[string]string{"error": fmt.Sprintf("Finance with id %v doesn't exist!", id)}, http.StatusNotFound)
			return
		}

		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	utils.Encode(w, transaction, http.StatusOK)

}

func DeleteTransactionById(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("Id")

	if err := storage.DeleteTransactionById(id); err != nil {

		if err == sql.ErrNoRows {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)

}

func GetCategories(w http.ResponseWriter, r *http.Request) {

	categories, err := storage.GetCategories()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	utils.Encode(w, &map[string][]types.Category{"data": categories}, http.StatusOK)

}

func CreateCategory(w http.ResponseWriter, r *http.Request) {

	categoryRequest, err := utils.Decode[types.CategoryRequestData](r)
	if err != nil {
		utils.Encode(w, &map[string]string{"error": "The body structure is wrong!"}, http.StatusBadRequest)
		return
	}

	category, err := storage.CreateCategory(categoryRequest)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	utils.Encode(w, category, http.StatusCreated)
}

func GetCategoryById(w http.ResponseWriter, r *http.Request) {

	id := r.PathValue("Id")

	category, err := storage.GetCategoryById(id)

	if err != nil {
		if err == sql.ErrNoRows {
			utils.Encode(w, &map[string]string{"error": fmt.Sprintf("Category with id %v doesn't exist!", id)}, http.StatusNotFound)
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	utils.Encode(w, category, http.StatusOK)
}

func UpdateCategoryById(w http.ResponseWriter, r *http.Request) {

	id := r.PathValue("Id")

	categoryRequest, err := utils.Decode[types.CategoryRequestData](r)

	if err != nil {
		utils.Encode(w, &map[string]string{"error": "The body structure is wrong!"}, http.StatusBadRequest)
		return
	}

	category, err := storage.UpdateCategoryById(id, categoryRequest)

	if err != nil {
		if err == sql.ErrNoRows {
			utils.Encode(w, &map[string]string{"error": fmt.Sprintf("Category with id %v doesn't exist!", id)}, http.StatusNotFound)
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	utils.Encode(w, category, http.StatusOK)
}

func DeleteCategoryById(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("Id")
	if err := storage.DeleteCategoryById(id); err != nil {

		if err == sql.ErrNoRows {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func GetReport(w http.ResponseWriter, r *http.Request) {
	startParam := r.URL.Query().Get("start")
	endParam := r.URL.Query().Get("end")

	if startParam == "" || endParam == "" {
		utils.Encode(w, &map[string]any{"data": nil}, http.StatusOK)
		return
	}

	start, err := utils.ParseDate(startParam)
	if err != nil {
		utils.Encode(w, &map[string]string{"error": "Invalid start date format. Use YYYY-MM-DD"}, http.StatusBadRequest)
		return
	}

	end, err := utils.ParseDate(endParam)
	if err != nil {
		utils.Encode(w, &map[string]string{"error": "Invalid end date format. Use YYYY-MM-DD"}, http.StatusBadRequest)
		return
	}

	report, err := storage.GetReport(start, end)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	utils.Encode(w, &map[string]*types.Report{"data": report}, http.StatusOK)
}
