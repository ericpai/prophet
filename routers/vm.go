package routers

import (
	"github.com/ericpai/prophet/controllers"
	"github.com/gorilla/mux"
)

func initAccountRouters(apiRouter *mux.Router) {
	vmRouter := apiRouter.PathPrefix("/vm").Subrouter()
	vmRouter.HandleFunc("/overview", controllers.GetVMInstancesHandler).Methods("GET")
	vmRouter.HandleFunc("/offerings", controllers.GetVMOfferingsHandler).Methods("GET")
}
