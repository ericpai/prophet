package routers

import (
	"github.com/gorilla/mux"
	"net/http"
)

var MainRouter *mux.Router

func init() {
	MainRouter = mux.NewRouter()
	apiRouter := MainRouter.PathPrefix("/api").Subrouter()
	MainRouter.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/")))
	initVMRouters(apiRouter)
	initAccountRouters(apiRouter)
}
