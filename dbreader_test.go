package dbreader

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"testing"
	"fmt"
)

type res struct {
	ComplaintId int `json:"complaint_id" bson:"complaint_id"`
}

func TestReadOne(t *testing.T) {
	session, _ := mgo.Dial("127.0.0.1")
	var r res
	mr := MongoReader{session,"companies","complaints",bson.M{"state":"FL"}}
	err := readOne(mr, &r)
	if err != nil {
		t.Error(err)
		t.Fail()
	}
	fmt.Print(r)
}
