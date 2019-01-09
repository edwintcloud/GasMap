package models

import "github.com/globalsign/mgo/bson"

// Trip is our trip model
type Trip struct {
	ID         bson.ObjectId   `json:"_id,omitempty" bson:"_id,omitempty"`
	Name       string          `json:"name" bson:"name"`
	VehicleID  bson.ObjectId   `json:"vehicleId,omitempty" bson:"vehicleId,omitempty"`
	CurrentMte string          `json:"currentMte" bson:"currentMte"`
	Status     string          `json:"status" bson:"status"`
	From       Location        `json:"from" bson:"from"`
	To         Location        `json:"to" bson:"to"`
	Gallons    string          `json:"gallons" bson:"gallons"`
	Price      string          `json:"price" bson:"price"`
	Stations   []bson.ObjectId `json:"stations,omitempty" bson:"stations,omitempty"`
}
