package api

import (
	"encoding/json"
	"net/http"

	"app/service"
)

func SharedAccountHandler(w http.ResponseWriter, r *http.Request) {
	accounts, err := service.GetSharedAccounts()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(accounts)
}
