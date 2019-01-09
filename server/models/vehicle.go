package models

import (
	"github.com/globalsign/mgo/bson"
)

// Vehicle is our vehicle model
type Vehicle struct {
	ID          bson.ObjectId `json:"_id,omitempty" bson:"_id,omitempty"`
	Year        string        `json:"year" bson:"year"`
	Make        string        `json:"make" bson:"make"`
	Model       string        `json:"model" bson:"model"`
	ImgURL      string        `json:"imgUrl,omitempty" bson:"imgUrl,omitempty"`
	Mpg         string        `json:"mpg" bson:"mpg"`
	TankSize    string        `json:"tankSize" bson:"taskSize"`
	FuelQuality string        `json:"fuelQuality" bson:"fuelQuality"`
}
