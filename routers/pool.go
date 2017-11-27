package routers

import (
	"pool/common"
	"pool/controllers"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

func SetPoolRoutes(r *mux.Router) *mux.Router {
	poolRouter := mux.NewRouter()
	poolRouter.HandleFunc("/pools", controllers.Pool.Create).Methods("POST")
	poolRouter.HandleFunc("/pools/{id}", controllers.Pool.Update).Methods("PUT")
	poolRouter.HandleFunc("/pools", controllers.Pool.Get).Methods("Get")
	poolRouter.HandleFunc("/pools/{id}", controllers.Pool.GetById).Methods("GET")
	poolRouter.HandleFunc("/pools/users/{id}", controllers.Pool.GetByUser).Methods("GET")
	poolRouter.HandleFunc("/pools/{id}", controllers.Pool.Delete).Methods("DELETE")
	r.PathPrefix("/tasks").Handler(negroni.New(
		negroni.HandlerFunc(common.Middleware.Authorize),
		negroni.Wrap(poolRouter),
	))
	return router
}
