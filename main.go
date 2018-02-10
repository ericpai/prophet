package main

import (
	"log"
	"net/http"
	"time"

	"github.com/ericpai/prophet/controllers"
	"github.com/ericpai/prophet/iaas"
	"github.com/ericpai/prophet/routers"
)

func main() {
	iaas.InitConfig("account-secret.yaml")
	controllers.InitSecretary()
	srv := &http.Server{
		Handler:      routers.MainRouter,
		Addr:         "127.0.0.1:8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}
