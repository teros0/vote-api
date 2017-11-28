package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"pool/common"
	"pool/data"
	"pool/models"
)

func CreateQuestion(w http.ResponseWriter, r *http.Request) {
	var dataResource QuestionResource
	err := json.NewDecoder(r.Body).Decode(&dataResource)
	if err != nil {
		common.DisplayAppError(w, err, "Invalid Question data", 500)
		return
	}
	questionModel := dataResource.Data
	question := &models.Question{
		PoolId:      bson.ObjectIdHex(questionModel.PoolId),
		Description: questionModel.Description,
	}
	context := NewContext()
	defer context.Close()
	col := context.DbCollection("questions")
	repo := &data.QuestionRepository{C: col}
	repo.Create(question)
	j, err := json.Marshal(question)
	if err != nil {
		common.DisplayAppError(w, err, "An unexpected error has occurred", 500)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(j)
}

func GetQuestionsByPool(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	context := NewContext()
	defer context.Close()
	col := context.DbCollection("questions")
	repo := &data.QuestionRepository{C: col}
	questions := repo.GetByPool(id)
	j, err := json.Marshal(QuestionsResource{Data: questions})
	if err != nil {
		common.DisplayAppError(w, err, "An unexpected error has occurred", 500)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

func GetQuestions(w http.ResponseWriter, r *http.Request) {
	context := NewContext()
	defer context.Close()
	col := context.DbCollection("questions")
	repo := &data.QuestionRepository{C: col}
	questions := repo.GetAll()
	j, err := json.Marshal(QuestionsResource{Data: questions})
	if err != nil {
		common.DisplayAppError(w, err, "An unexpected error has occurred", 500)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

func GetQuestionById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	context := NewContext()
	defer context.Close()
	col := context.DbCollection("questions")
	repo := &data.QuestionRepository{C: col}
	question, err := repo.GetById(id)
	if err != nil {
		if err == mgo.ErrNotFound {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		common.DisplayAppError(w, err, "An unexpected error has occurred", 500)
		return

	}
	j, err := json.Marshal(question)
	if err != nil {
		common.DisplayAppError(w, err, "An unexpected error has occurred", 500)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

func UpdateQuestion(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := bson.ObjectIdHex(vars["id"])
	var dataResource QuestionResource
	err := json.NewDecoder(r.Body).Decode(&dataResource)
	if err != nil {
		common.DisplayAppError(w, err, "Invalid Question data", 500)
		return
	}
	questionModel := dataResource.Data
	question := &models.Question{
		Id:          id,
		Description: questionModel.Description,
	}
	context := NewContext()
	defer context.Close()
	col := context.DbCollection("questions")
	repo := &data.QuestionRepository{C: col}
	if err := repo.Update(question); err != nil {
		common.DisplayAppError(w, err, "An unexpected error has occurred", 500)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func DeleteQuestion(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	context := NewContext()
	defer context.Close()
	col := context.DbCollection("questions")
	repo := &data.QuestionRepository{C: col}
	err := repo.Delete(id)
	if err != nil {
		common.DisplayAppError(w, err, "An unexpected error has occurred", 500)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
