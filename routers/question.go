package routers

import (
	"pool/common"
	"pool/controllers"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

func SetQuestionRoutes(r *mux.Router) *mux.Router {
	questionRouter := mux.NewRouter()
	questionRouter.HandleFunc("/questions", controllers.Question.Create).Methods("POST")
	questionRouter.HandleFunc("/questions/{id}", controllers.Question.Update).Methods("PUT")
	questionRouter.HandleFunc("/questions", controllers.Question.Get).Methods("Get")
	questionRouter.HandleFunc("/questions/{id}", controllers.Question.GetById).Methods("GET")
	questionRouter.HandleFunc("/questions/pools/{id}", controllers.Question.GetByUser).Methods("GET")
	questionRouter.HandleFunc("/questions/{id}", controllers.Question.Delete).Methods("DELETE")
	r.PathPrefix("/tasks").Handler(negroni.New(
		negroni.HandlerFunc(common.Middleware.Authorize),
		negroni.Wrap(questionRouter),
	))
	return router
}
