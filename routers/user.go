package routers

import (
	"pool/controllers"

	"github.com/gorilla/mux"
)

func SetUserRoutes(router *mux.Router) *mux.Router {
	router.HandleFunc("/users/register", controllers.User.Register).Methods("POST")
	router.HandleFunc("/users/login", controllers.User.Login).Methods("POST")
	return router
}
