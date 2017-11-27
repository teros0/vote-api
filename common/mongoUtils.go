package common

import (
	"log"

	"gopkg.in/mgo.v2"
)

var session *mgo.Session

func GetSession() *mgo.Session {
	if session == nil {
		var err error
		session, err = mgo.DialWithInfo(&mgo.DialInfo{
			Addrs: []string{C.MongoAddress},
		})
		if err != nil {
			log.Fatalf("GetSession: %s\n", err)
		}
	}
	return session.Copy()
}

func createDbSession() {
	var err error
	session, err = mgo.DialWithInfo(&mgo.DialInfo{
		Addrs: []string{C.MongoAddress},
	})
	if err != nil {
		log.Fatalf("GetSession: %s\n", err)
	}
}

func addIndexes() {
	var err error
	userIndex := mgo.Index{
		Key:        []string{"email"},
		Unique:     true,
		Background: true,
		Sparse:     true,
	}
	poolIndex := mgo.Index{
		Key:        []string{"createdby"},
		Unique:     false,
		Background: true,
		Sparse:     true,
	}
	questionIndex := mgo.Index{
		Key:        []string{"questionid"},
		Unique:     false,
		Background: true,
		Sparse:     true,
	}
	voteIndex := mgo.Index{
		Key:        []string{"questionid"},
		Unique:     false,
		Background: true,
		Sparse:     true,
	}
	session := GetSession()
	defer session.Close()

	userCol := session.DB("").C("users")
	poolCol := session.DB("").C("pools")
	questionCol := session.DB("").C("questions")
	voteCol := session.DB("").C("votes")
	err = userCol.EnsureIndex(userIndex)
	if err != nil {
		log.Println("addIndex error", err)
	}
	err = poolCol.EnsureIndex(poolIndex)
	if err != nil {
		log.Println("addIndex error", err)
	}
	err = questionCol.EnsureIndex(questionIndex)
	if err != nil {
		log.Println("addIndex error", err)
	}
	err = voteCol.EnsureIndex(voteIndex)
	if err != nil {
		log.Println("addIndex error", err)
	}
}
