package main

import (
	"log"
	"net/http"

	"github.com/codegangsta/negroni"
	"github.com/greatontime/gotask-api/common"
	"github.com/greatontime/gotask-api/routers"
)

func main() {
	common.StartUp()
	router := routers.InitRoutes()
	n := negroni.Classic()
	n.UseHandler(router)
	server := &http.Server{
		Addr:    common.AppConfig.Server,
		Handler: n,
	}
	log.Println("Listening ...")
	server.ListenAndServe()
}
