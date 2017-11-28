package controllers

import (
	"encoding/json"
	"net/http"
	"pool/common"
	"pool/data"

	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func CreatePool(w http.ResponseWriter, r *http.Request) {
	var dataResource PoolResource
	err := json.NewDecoder(r.Body).Decode(&dataResource)
	if err != nil {
		common.DisplayAppError(w, err, "Invalid pool data", http.StatusBadRequest)
		return
	}
	pool := &dataResource.Data
	context := NewContext()
	defer context.Close()
	c := context.DbCollection("pools")
	repo := &data.PoolRepository{c}
	repo.Create(pool)
	j, err := json.Marshal(PoolResource{Data: *pool})
	if err != nil {
		common.DisplayAppError(w, err, "An unexpected error has occured", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(j)
}

func GetPools(w http.ResponseWriter, r *http.Request) {
	context := NewContext()
	defer context.Close()
	c := context.DbCollection("pools")
	repo := &data.PoolRepository{c}
	pools := repo.GetAll()
	j, err := json.Marshal(PoolsResource{Data: pools})
	if err != nil {
		common.DisplayAppError(w, err, "An unexpected error has occured", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(j)
}

func GetPoolById(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	context := NewContext()
	defer context.Close()
	c := context.DbCollection("pools")
	repo := &data.PoolRepository{c}
	pool, err := repo.GetById(id)
	if err != nil {
		if err == mgo.ErrNotFound {
			common.DisplayAppError(w, err, "No resource found", http.StatusNoContent)
		} else {
			common.DisplayAppError(w, err, "An unexpected error has occured", http.StatusInternalServerError)
		}
		return
	}
	j, err := json.Marshal(PoolResource{Data: pool})
	if err != nil {
		common.DisplayAppError(w, err, "An unexpected error has occured", http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(j)
}

func GetPoolsByUser(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	context := NewContext()
	defer context.Close()
	c := context.DbCollection("pools")
	repo := &data.PoolRepository{c}
	pools := repo.GetByUser(id)
	j, err := json.Marshal(PoolsResource{Data: pools})
	if err != nil {
		common.DisplayAppError(w, err, "An unexpected error has occured", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(j)
}

func UpdatePool(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := bson.ObjectIdHex(vars["id"])
	var dataResource PoolResource
	err := json.NewDecoder(r.Body).Decode(&dataResource)
	if err != nil {
		common.DisplayAppError(w, err, "Invalid pool data", http.StatusBadRequest)
		return
	}
	pool := &dataResource.Data
	pool.Id = id
	context := NewContext()
	defer context.Close()
	c := context.DbCollection("pools")
	repo := &data.PoolRepository{C: c}
	if err := repo.Update(pool); err != nil {
		common.DisplayAppError(w, err, "An unexpected error has occured", http.StatusInternalServerError)
		return
	}
}

func DeletePool(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	context := NewContext()
	defer context.Close()
	c := context.DbCollection("pools")
	repo := &data.PoolRepository{c}
	if err := repo.Delete(id); err != nil {
		common.DisplayAppError(w, err, "An unexpected errror has occured", http.StatusInternalServerError)
		return
	}
}
