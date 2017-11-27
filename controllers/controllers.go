package controllers

import (
	"fmt"
	"net/http"
)

func TestController(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "CONTROLL")
}
