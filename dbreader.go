package dbreader

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type MongoReader struct {
	S *mgo.Session
	D string
	C string
	Q bson.M
}

func (mr MongoReader) readOne(response interface{}) error {
	s := mr.S
	d := mr.D
	c := mr.C
	q := mr.Q
	err := s.DB(d).C(c).Find(q).One(response)
	if err != nil {
		return err
	}
	return nil
}

type Reader interface {
	readOne(response interface{}) error
}

func readOne(r Reader, response interface{}) error {
	err := r.readOne(response)
	return err
}