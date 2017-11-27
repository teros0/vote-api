package routers

import (
	"pool/common"
	"pool/controllers"

	"github.com/codegangsta/negroni"

	"github.com/gorilla/mux"
)

func SetUserRoutes(router *mux.Router) *mux.Router {
	router.HandleFunc("/users/register", controllers.User.Register).Methods("POST")
	router.HandleFunc("/users/login", controllers.User.Login).Methods("POST")
	router.Handle("/users/logout", negroni.New(
		common.Middleware.Authorize,
		negroni.Wrap(controllers.User.Logout)),
	).Methods("POST")
	return router
}
