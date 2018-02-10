package routers

import "github.com/gorilla/mux"

var MainRouter *mux.Router

func init() {
	MainRouter = mux.NewRouter()
	apiRouter := MainRouter.PathPrefix("/api").Subrouter()
	initVMRouters(apiRouter)
}
