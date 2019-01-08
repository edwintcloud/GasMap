package models

import "github.com/globalsign/mgo/bson"

// User is our user model
type User struct {
	ID        bson.ObjectId `json:"_id,omitempty" bson:"_id,omitempty"`
	Email     string        `json:"email" bson:"email"`
	Password  string        `json:"password" bson:"password"`
	FirstName string        `json:"firstName" bson:"firstName"`
	LastName  string        `json:"lastName" bson:"lastName"`
	// []vehicles
	// []trips
}
