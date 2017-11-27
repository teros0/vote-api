package main

import (
	"net/http"
	"pool/config"
	"pool/routers"
)

func main() {
	r := routers.InitRouter()
	s := &http.Server{
		Addr:    config.ServerAddress,
		Handler: r,
	}
	s.ListenAndServe()
}
