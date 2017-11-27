package controllers

import (
	"fmt"
	"net/http"
)

type user struct{}

var User user

func (u *user) Login(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "LOGIN")
}

func (u *user) Logout(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "LOGOUT")
}
