package controllers

import (
	"pool/models"
)

type (
	// 4 POST - user/register
	UserResource struct {
		Data models.User `json:"data"`
	}
	// 4 POST - user/login
	LoginResource struct {
		Data LoginModel `json:"data"`
	}
	// Responce for authorized user POST - /user/login
	AuthUserResource struct {
		Data AuthUserModel `json:"data"`
	}
	// Model for authentication
	LoginModel struct {
		Email string `json"email"`
		Password string `json:"password"`
	}
	// Model for authorized user with access token
	AuthUserModel struct {
		User models.User `json:"user"`
		Token string `json:""token`
	}
	// For POST/PUT - /pools
	// For GET - pools/id
	PoolResource struct {
		Data models.Pool `json:"data"`
	}
	// For Get - /pools
	PoolsResource struct {
		Data []models.Task `json:"data"`
	}	
)