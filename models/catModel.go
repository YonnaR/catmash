package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

/*Cat is db representation of a cat */
type Cat struct {
	ID       bson.ObjectId `bson:"_id" json:"_id" binding:"required"`
	FileName string        `bson:"fileName" json:"fileName"`
	Name     string        `bson:"name" json:"name"`
	Img      string        `bson:"img" json:"img"`
	Vote     int           `bson:"vote" json:"vote"`
	Created  time.Time     `bson:"created" json:"created"`
}
