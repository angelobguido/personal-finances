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

func GetFinances(w http.ResponseWriter, r *http.Request) {

	finances, err := storage.GetFinances()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	utils.Encode(w, &map[string]types.Finances{"data": finances}, http.StatusOK)

}

func CreateFinance(w http.ResponseWriter, r *http.Request) {

	financeRequest, err := utils.Decode[types.FinanceRequest](r)

	if err != nil {
		utils.Encode(w, &map[string]string{"error": "The body structure is wrong!"}, http.StatusBadRequest)
		return
	}

	if financeRequest.Amount == nil || financeRequest.Name == nil || financeRequest.Category == nil {
		utils.Encode(w, &map[string]string{"error": "All fields are required!"}, http.StatusBadRequest)
		return
	}

	finance, err := storage.CreateFinance(*financeRequest.Name, *financeRequest.Amount, *financeRequest.Category)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	utils.Encode(w, &map[string]types.Finance{"data": *finance}, http.StatusCreated)
}

func GetFinanceById(w http.ResponseWriter, r *http.Request) {

	id := r.PathValue("Id")

	finance, err := storage.GetFinanceById(id)

	if err != nil {

		if err == sql.ErrNoRows {
			utils.Encode(w, &map[string]string{"error": fmt.Sprintf("Finance with id %v doesn't exist!", id)}, http.StatusNotFound)
			return
		}

		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	utils.Encode(w, &map[string]types.Finance{"data": *finance}, http.StatusOK)
}

func UpdateFinanceById(w http.ResponseWriter, r *http.Request) {

	id := r.PathValue("Id")

	financeRequest, err := utils.Decode[types.FinanceRequest](r)

	if err != nil {
		utils.Encode(w, &map[string]string{"error": "The body structure is wrong!"}, http.StatusBadRequest)
		return
	}

	finance, err := storage.UpdateFinanceById(id, financeRequest.Name, financeRequest.Amount, financeRequest.Category)

	if err != nil {

		if err == sql.ErrNoRows {
			utils.Encode(w, &map[string]string{"error": fmt.Sprintf("Finance with id %v doesn't exist!", id)}, http.StatusNotFound)
			return
		}

		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	utils.Encode(w, &map[string]types.Finance{"data": *finance}, http.StatusOK)

}

func DeleteFinanceById(w http.ResponseWriter, r *http.Request) {

	id := r.PathValue("Id")

	if err := storage.DeleteFinanceById(id); err != nil {

		if err == sql.ErrNoRows {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)

}
