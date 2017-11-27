package data

import (
	"pool/models"

	"golang.org/x/crypto/bcrypt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type UserRepository struct {
	C *mgo.Collection
}

func (r *UserRepository) CreateUser(user *models.User) error {
	user.Id = bson.NewObjectId()
	hpass, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.HashPassword = hpass
	user.Password = ""
	err = r.C.Insert(&user)
	return err
}

func (r *UserRepository) Login(user models.User) (u models.User, err error) {
	if err = r.C.Find(bson.M{"email": user.Email}).One(&u); err != nil {
		return
	}
	if err = bcrypt.CompareHashAndPassword(u.HashPassword, []byte(user.Password)); err != nil {
		u = models.User{}
	}
	return
}
