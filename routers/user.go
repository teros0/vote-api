package routers

import (
	"pool/controllers"

	"github.com/gorilla/mux"
)

func SetUserRoutes(router *mux.Router) *mux.Router {
	router.HandleFunc("/users/login", controllers.User.Login).Methods("POST")
	router.HandleFunc("/users/logout", controllers.User.Logout).Methods("POST")
	return router
}
