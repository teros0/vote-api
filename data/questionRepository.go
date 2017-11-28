package data

import (
	"time"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"pool/models"
)

type QuestionRepository struct {
	C *mgo.Collection
}

func (r *QuestionRepository) Create(question *models.Question) error {
	obj_id := bson.NewObjectId()
	question.Id = obj_id
	question.CreatedOn = time.Now()
	err := r.C.Insert(&question)
	return err
}

func (r *QuestionRepository) Update(question *models.Question) error {
	// partial update on MogoDB
	err := r.C.Update(bson.M{"_id": question.Id},
		bson.M{"$set": bson.M{
			"description": question.Description,
		}})
	return err
}
func (r *QuestionRepository) Delete(id string) error {
	err := r.C.Remove(bson.M{"_id": bson.ObjectIdHex(id)})
	return err
}
func (r *QuestionRepository) GetByPool(id string) []models.Question {
	var questions []models.Question
	poolid := bson.ObjectIdHex(id)
	iter := r.C.Find(bson.M{"poolid": poolid}).Iter()
	result := models.Question{}
	for iter.Next(&result) {
		questions = append(questions, result)
	}
	return questions
}
func (r *QuestionRepository) GetAll() []models.Question {
	var questions []models.Question
	iter := r.C.Find(nil).Iter()
	result := models.Question{}
	for iter.Next(&result) {
		questions = append(questions, result)
	}
	return questions
}
func (r *QuestionRepository) GetById(id string) (question models.Question, err error) {
	err = r.C.FindId(bson.ObjectIdHex(id)).One(&question)
	return
}
