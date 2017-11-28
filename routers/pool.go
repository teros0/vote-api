package routers

import (
	"pool/common"
	"pool/controllers"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

func SetPoolRoutes(r *mux.Router) *mux.Router {
	poolRouter := mux.NewRouter()
	poolRouter.HandleFunc("/pools", controllers.CreatePool).Methods("POST")
	poolRouter.HandleFunc("/pools/{id}", controllers.UpdatePool).Methods("PUT")
	poolRouter.HandleFunc("/pools", controllers.GetPools).Methods("Get")
	poolRouter.HandleFunc("/pools/{id}", controllers.GetPoolById).Methods("GET")
	poolRouter.HandleFunc("/pools/users/{id}", controllers.GetPoolsByUser).Methods("GET")
	poolRouter.HandleFunc("/pools/{id}", controllers.DeletePool).Methods("DELETE")
	r.PathPrefix("/tasks").Handler(negroni.New(
		negroni.HandlerFunc(common.Middleware.Authorize),
		negroni.Wrap(poolRouter),
	))
	return r
}
