package routers

import (
	"github.com/greatontime/controller"
	"github.com/gorilla/mux"
)

func SetUserRoutes(router *mux.Router) *mux.Router {
	router.HandleFunc("/users/register", controller.Register).Methods("POST")
	router.HandleFunc("/users/login", controller.Login).Methods("POST")
	return router
}
