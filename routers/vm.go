package routers

import (
	"github.com/ericpai/prophet/controllers"
	"github.com/gorilla/mux"
)

func initVMRouters(apiRouter *mux.Router) {
	vmRouter := apiRouter.PathPrefix("/vm").Subrouter()
	vmRouter.HandleFunc("/overview", controllers.GetVMInstancesHandler).Methods("GET")
}
