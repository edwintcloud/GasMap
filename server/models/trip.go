package models

import (
	"github.com/edwintcloud/GasMap/server/utils"
	"github.com/globalsign/mgo/bson"
)

// Trip is our trip model
type Trip struct {
	ID         bson.ObjectId   `json:"_id,omitempty" bson:"_id,omitempty"`
	Name       string          `json:"name" bson:"name"`
	VehicleID  bson.ObjectId   `json:"vehicleId,omitempty" bson:"vehicleId,omitempty"`
	CurrentMte string          `json:"currentMte" bson:"currentMte"`
	Status     string          `json:"status,omitempty" bson:"status,omitempty"`
	From       Location        `json:"from" bson:"from"`
	To         Location        `json:"to" bson:"to"`
	Gallons    string          `json:"gallons,omitempty" bson:"gallons,omitempty"`
	Price      string          `json:"price,omitempty" bson:"price,omitempty"`
	Stations   []bson.ObjectId `json:"stations,omitempty" bson:"stations,omitempty"`
}

// Create creates a new trip in the db
func (t *Trip) Create() error {

	// create new trip in db
	err := utils.DB.C("trips").Insert(t)
	if err != nil {
		return err
	}

	// if all went well, return nil
	return nil
}

// FindByID finds a trip in the db by id
func (t *Trip) FindByID() error {

	// find trip in db by id
	err := utils.DB.C("trips").FindId(t).One(&t)
	if err != nil {
		return err
	}

	// if all went well, return nil
	return nil
}

// RemoveByID removes a trip from the db by id
func (t *Trip) RemoveByID() error {

	// delete trip from db by id
	err := utils.DB.C("trips").RemoveId(t.ID)
	if err != nil {
		return err
	}

	// if all went well, return nil
	return nil
}
