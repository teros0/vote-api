package routers

import (
	"github.com/gorilla/mux"
)

func InitRouter() *mux.Router {
	r := mux.NewRouter()
	r = SetUserRoutes(r)
	r = SetQuestionRoutes(r)
	r = SetPoolRoutes(r)
	return r
}
