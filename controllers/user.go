package controllers

import (
	"encoding/json"
	"net/http"
	"pool/common"
	"pool/data"
	"pool/models"
)

type user struct{}

var User user

func (u *user) Register(w http.ResponseWriter, r *http.Request) {
	var dataResource UserResource
	if err := json.NewDecoder(r.Body).Decode(&dataResource); err != nil {
		common.DisplayAppError(w, err, "Invalid user data", 500)
		return
	}
	user := &dataResource.Data
	context := NewContext()
	defer context.Close()
	c := context.DbCollection("users")
	repo := &data.UserRepository{c}
	// Insert User Document
	repo.CreateUser(user)
	// Eliminating hash pass from response
	user.HashPassword = nil
	if j, err := json.Marshal(UserResource{Data: *user}); err != nil {
		common.DisplayAppError(w, err, "An unexpected error has occured", 500)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(j)
}

func (u *user) Login(w http.ResponseWriter, r *http.Request) {
	var dataResource LoginResource
	var token string
	if err := json.NewDecoder(r.Body).Decode(&dataResource); err != nil {
		common.DisplayAppError(w, err, "Invalid login data", 500)
		return
	}
	loginModel := dataResource.Data
	loginUser := models.User {
		Email: loginModel.Email,
		Password: loginModel.Password,
	}
	context := NewContext()
	defer context.Close()
	c := context.DbCollection("users")
	repo := &data.UserRepository{c}

	if user, err := repo.Login(loginUser); err != nil {
		common.DisplayAppError(w, err, "Invalid login credentials", 401)
		return
	}

	token, err = common.GenerateJWT(user.Email, "member")
	if err != nil {
		common.DisplayAppError(w, err, "Error while generating the access token", 500)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	user.HashPassword = nil
	authUser := AuthUserModel{
		User: user,
		Token: token,
	}
	j, err := json.Marshal(AuthUserResource{Data: authUser})
	if err != nil {
		common.DisplayAppError(w, err, "An unexpected error has occured", 500)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

