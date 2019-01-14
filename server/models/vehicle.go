package models

import (
	"github.com/edwintcloud/GasMap/server/utils"
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

// Create creates a new vehicle in our db
func (v *Vehicle) Create() error {

	// generate object id for vehicle
	v.ID = bson.NewObjectId()

	// create vehicle in database
	err := utils.DB.C("vehicles").Insert(v)
	if err != nil {
		return err
	}

	// if all went well, return nil
	return nil
}

// FindByID finds a vehicle by ID
func (v *Vehicle) FindByID() error {

	// find vehicle by id in db
	err := utils.DB.C("vehicles").FindId(v.ID).One(&v)
	if err != nil {
		return err
	}

	// if all went well, return nil
	return nil
}

// RemoveByID removes a vehicle in db by id
func (v *Vehicle) RemoveByID() error {

	// delete vehicle from the db by id
	err := utils.DB.C("vehicles").RemoveId(v.ID)
	if err != nil {
		return err
	}

	// if all went well, return nil
	return nil
}
