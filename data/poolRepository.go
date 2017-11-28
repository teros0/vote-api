package data

import (
	"pool/models"
	"time"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type PoolRepository struct {
	C *mgo.Collection
}

func (r *PoolRepository) Create(pool *models.Pool) error {
	obj_id := bson.NewObjectId()
	pool.Id = obj_id
	pool.CreatedOn = time.Now()
	pool.Status = "Created"
	err := r.C.Insert(&pool)
	return err
}

func (r *PoolRepository) Update(pool *models.Pool) error {
	err := r.C.Update(bson.M{"_id": pool.Id},
		bson.M{"$set": bson.M{
			"name":        pool.Name,
			"description": pool.Description,
			"due":         pool.Due,
			"status":      pool.Status,
			"tags":        pool.Tags,
		}})
	return err
}

func (r *PoolRepository) Delete(id string) error {
	err := r.C.Remove(bson.M{"_id": bson.ObjectIdHex(id)})
	return err
}

func (r *PoolRepository) GetAll() []models.Pool {
	var pools []models.Pool
	iter := r.C.Find(nil).Iter()
	result := models.Pool{}
	for iter.Next(&result) {
		pools = append(pools, result)
	}
	return pools
}

func (r *PoolRepository) GetById(id string) (pool models.Pool, err error) {
	err = r.C.FindId(bson.ObjectIdHex(id)).One(&pool)
	return
}

func (r *PoolRepository) GetByUser(userId string) []models.Pool {
	var pools []models.Pool
	iter := r.C.Find(bson.M{"createdby": userId}).Iter()
	res := models.Pool{}
	for iter.Next(&res) {
		pools = append(pools, res)
	}
	return pools
}
