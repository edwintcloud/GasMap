package models

import (
	"errors"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/edwintcloud/GasMap/server/utils"
	"github.com/globalsign/mgo/bson"
)

// User is our user model
type User struct {
	ID        bson.ObjectId `json:"_id,omitempty" bson:"_id,omitempty"`
	Email     string        `json:"email" bson:"email"`
	Password  string        `json:"password,omitempty" bson:"password,omitempty"`
	Token     string        `json:"token,omitempty" bson:"-"`
	FirstName string        `json:"firstName" bson:"firstName"`
	LastName  string        `json:"lastName" bson:"lastName"`
	Vehicles  []interface{} `json:"vehicles,omitempty" bson:"vehicles,omitempty"`
	Trips     []interface{} `json:"trips,omitempty" bson:"trips,omitempty"`
}

// Create is our create method for users
func (u *User) Create() error {

	// check to make sure we have an email
	if len(u.Email) < 3 {
		return errors.New("email must be specified")
	}

	// create user in db
	err := utils.DB.C("users").Insert(u)
	// if error is duplicate key because the user was already created.. continue
	if err != nil && !strings.Contains(err.Error(), "E11000") {
		return err
	}

	// if all went well, return nil
	return nil
}

// FindByEmail finds a user by email
func (u *User) FindByEmail() error {

	// find user by email in db
	err := utils.DB.C("users").Find(bson.M{"email": u.Email}).One(&u)
	if err != nil {
		return err
	}

	// generate jwt for u
	err = u.generateJwt()
	if err != nil {
		return err
	}

	// if all went well, return nil
	return nil
}

// generate jwt token
func (u *User) generateJwt() error {
	var err error

	// create token
	token := jwt.New(jwt.SigningMethodHS256)

	// set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = u.ID.Hex()
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	// generate encoded token and set Token field
	u.Token, err = token.SignedString([]byte(utils.JwtSecret))
	if err != nil {
		return err
	}

	// if all went well, return nil
	return nil
}

// AddVehicle adds a vehicle to user
func (u *User) AddVehicle(v *Vehicle) error {

	// append vehicle id to user struct
	u.Vehicles = append(u.Vehicles, v.ID)

	// update user in db
	err := utils.DB.C("users").UpdateId(u.ID, u)
	if err != nil {
		return err
	}

	// if all went well, return nil
	return nil
}

// AddTrip adds a trip to user
func (u *User) AddTrip(t *Trip) error {

	// append trip id to user struct
	u.Trips = append(u.Trips, t.ID)

	// update user in db
	err := utils.DB.C("users").UpdateId(u.ID, u)
	if err != nil {
		return err
	}

	// if all went well, return nil
	return nil
}

// FindByID finds a user by id in the db
func (u *User) FindByID() error {

	// find user by id in db
	err := utils.DB.C("users").FindId(u.ID).One(&u)
	if err != nil {
		return err
	}

	// generate jwt for u
	err = u.generateJwt()
	if err != nil {
		return err
	}

	// if all went well, return nil
	return nil
}

// RemoveByID removes a user from the db by id
func (u *User) RemoveByID() error {

	// delete user from db by id
	err := utils.DB.C("users").RemoveId(u.ID)
	if err != nil {
		return err
	}

	// if all went well, return nil
	return nil
}

// RemoveTrip removes a trip from a user
func (u *User) RemoveTrip(t *Trip) error {
	var newTrips []interface{}

	// append elements not matching t to newTrips
	for _, v := range u.Trips {
		if v.(bson.ObjectId) != t.ID {
			newTrips = append(newTrips, v)
		}
	}

	// replace current u.Trips with newTrips, if newTrips is not empty
	if len(newTrips) > 0 {
		u.Trips = newTrips
	} else {
		u.Trips = nil
	}

	// update user in db
	err := utils.DB.C("users").UpdateId(u.ID, u)
	if err != nil {
		return err
	}

	// if all went well, return nil
	return nil
}
