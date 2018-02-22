package controllers

import (
	"encoding/json"
	"github.com/ericpai/prophet/libs"
	"net/http"
)

// GetVMInstancesHandler handles the overview of used instances
func GetVMInstancesHandler(w http.ResponseWriter, r *http.Request) {
	queries := r.URL.Query()
	if !libs.CheckKeysExist(queries, "account", "provider") {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}
	vmManager, err := getVMManager(queries.Get("provider"))
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte(err.Error()))
		return
	}
	rvObj, err := vmManager.OverviewInstances(queries.Get("account"))
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte(err.Error()))
		return
	}
	rv, _ := json.Marshal(rvObj)
	w.WriteHeader(http.StatusOK)
	w.Write(rv)
}

func GetVMOfferingsHandler(w http.ResponseWriter, r *http.Request) {
	queries := r.URL.Query()
	if !libs.CheckKeysExist(queries, "account", "provider") {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}
	vmManager, err := getVMManager(queries.Get("provider"))
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte(err.Error()))
		return
	}
	rvObj, err := vmManager.OverviewOfferings(queries.Get("account"))
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte(err.Error()))
		return
	}
	rv, _ := json.Marshal(rvObj)
	w.WriteHeader(http.StatusOK)
	w.Write(rv)
}
