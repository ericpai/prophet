package routers

import (
	"github.com/ericpai/prophet/controllers"
	"github.com/gorilla/mux"
)

func initVMRouters(apiRouter *mux.Router) {
	vmRouter := apiRouter.PathPrefix("/accounts").Subrouter()
	vmRouter.HandleFunc("/", controllers.GetAccountsHandler).Methods("GET")
}
