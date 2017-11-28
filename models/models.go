package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type (
	User struct {
		Id           bson.ObjectId `bson:"_id, omitempty" json:"id"`
		FirstName    string        `json:"firstname"`
		LastName     string        `json:"lastname"`
		Email        string        `json:"email"`
		Password     string        `json:"password, omitempty"`
		HashPassword []byte        `json:"hashpassword, omitempty"`
	}
	Pool struct {
		Id          bson.ObjectId `bson:"_id,omitempty" json:"id"`
		UserId      bson.ObjectId `json:"userid"`
		Name        string        `json:"name"`
		Description string        `json:"description"`
		CreatedOn   time.Time     `json:"createdon,omitempty"`
		Due         time.Time     `json:"due,omitempty"`
		Status      string        `json:"status,omitempty"`
		Tags        []string      `json:"tags,omitempty"`
	}
	Question struct {
		Id          bson.ObjectId `bson:"_id,omitempty" json:"id"`
		PoolId      bson.ObjectId `json:"poolid"`
		Description string        `json:"description"`
		CreatedOn   time.Time     `json:"createdon,omitempty"`
	}
	Vote struct {
		Id         bson.ObjectId `bson:"_id,omitempty" json:"id"`
		UserId     bson.ObjectId `json:"userid"`
		QuestionId bson.ObjectId `json:"questionid"`
	}
)
