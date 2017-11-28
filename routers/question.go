package routers

import (
	"pool/common"
	"pool/controllers"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

func SetQuestionRoutes(r *mux.Router) *mux.Router {
	questionRouter := mux.NewRouter()
	questionRouter.HandleFunc("/questions", controllers.CreateQuestion).Methods("POST")
	questionRouter.HandleFunc("/questions/{id}", controllers.UpdateQuestion).Methods("PUT")
	questionRouter.HandleFunc("/questions", controllers.GetQuestions).Methods("Get")
	questionRouter.HandleFunc("/questions/{id}", controllers.GetQuestionById).Methods("GET")
	questionRouter.HandleFunc("/questions/pools/{id}", controllers.GetQuestionsByPool).Methods("GET")
	questionRouter.HandleFunc("/questions/{id}", controllers.DeleteQuestion).Methods("DELETE")
	r.PathPrefix("/tasks").Handler(negroni.New(
		negroni.HandlerFunc(common.Middleware.Authorize),
		negroni.Wrap(questionRouter),
	))
	return r
}
