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
	ID        bson.ObjectId   `json:"_id,omitempty" bson:"_id,omitempty"`
	Email     string          `json:"email" bson:"email"`
	Password  string          `json:"password,omitempty" bson:"password,omitempty"`
	Token     string          `json:"token,omitempty" bson:"-"`
	FirstName string          `json:"firstName" bson:"firstName"`
	LastName  string          `json:"lastName" bson:"lastName"`
	Vehicles  []interface{}   `json:"vehicles,omitempty" bson:"vehicles,omitempty"`
	Trips     []bson.ObjectId `json:"trips,omitempty" bson:"trips,omitempty"`
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

	// set password to "" so we don't expose it
	u.Password = ""

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

// FindByID finds a user by id in the db
func (u *User) FindByID() error {

	// find user by id in db
	err := utils.DB.C("users").FindId(u.ID).One(&u)
	if err != nil {
		return err
	}

	// set password to "" so we don't expose it
	u.Password = ""

	// if all went well, return nil
	return nil
}
