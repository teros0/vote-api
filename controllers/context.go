package controllers

import (
	"gopkg.in/mgo.v2"
	"pool/common"
)

type Context struct {
	MongoSession *mgo.Session
}

func (c *Context) Close() {
	c.MongoSession.Close()
}

func (c *Context) DbCollection(name string) *mgo.Collection {
	return c.MongoSession.DB("").C(name)
}

func NewContext() *Context {
	context := &Context {
		MongoSession: common.GetSession(),
	}
	return context
}