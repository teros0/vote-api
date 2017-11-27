package main

import (
	"log"
	"net/http"
	"pool/common"
	"pool/routers"

	"github.com/codegangsta/negroni"
)

func main() {
	//Initializing global variables
	common.StartUp()
	r := routers.InitRouter()
	n := negroni.Classic()
	n.UseHandler(r)
	s := &http.Server{
		Addr:    config.ServerAddress,
		Handler: n,
	}
	s.ListenAndServe()
	log.Println("App successfully started on address", common.C.ServerAddress)
}
