package controllers

import (
	"encoding/json"
	"github.com/ericpai/prophet/iaas"
	"net/http"
)

// GetAccountsHandler handles the account list of prophet
func GetAccountsHandler(w http.ResponseWriter, r *http.Request) {
	rvObj := iaas.ListAccounts()
	rv, _ := json.Marshal(rvObj)
	w.WriteHeader(http.StatusOK)
	w.Write(rv)
}
