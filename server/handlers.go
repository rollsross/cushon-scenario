package server

import (
	"encoding/json"
	"net/http"

	"github.com/rodionross/cushon-scenario/storage"
)

type CreateAccoutAndFundBody struct {
	AccountTypeId string `json:"accountTypeId" example:"018ef16a-31a7-7e11-a77d-78b2eea91e2f"`
	FundId        string `json:"fundId" example:"018ef16a-31a7-7e11-a77d-78b2eea91e2f"`
	Balance       int    `json:"balance" example:"2500000"`
}

func HandleCreateAccountAndFund(s storage.Repository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userId := r.PathValue("id")

		var body CreateAccoutAndFundBody

		err := json.NewDecoder(r.Body).Decode(&body)
		if err != nil {
			http.Error(w, "invalid jason", http.StatusBadRequest)
			return
		}

		err = s.CreateAccoutAndFund(userId, body.AccountTypeId, body.FundId, body.Balance)
		if err != nil {
			http.Error(w, "failed to create and account and fund", http.StatusBadRequest)
			return
		}
	}
}

func HandleGetAccountAndFund(s storage.Repository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userId := r.PathValue("id")

		data, err := s.GetAccountAndFund(userId)
		if err != nil {
			http.Error(w, "user not found", http.StatusBadRequest)
			return
		}

		res, err := json.Marshal(data)
		if err != nil {
			http.Error(w, "invalid json", http.StatusBadRequest)
			return
		}

		w.Header().Add("Content-Type", "application/json")
		w.Write(res)
	}
}
