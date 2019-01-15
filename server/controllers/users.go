package controllers

import (
	"fmt"
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/edwintcloud/GasMap/server/models"
	"github.com/edwintcloud/GasMap/server/utils"
	"github.com/globalsign/mgo/bson"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// UserController is our user controller struct
type UserController struct {
	E *echo.Echo
}

// Register registers our user controller routes
func (c *UserController) Register() {

	// authorization not needed to create a user
	c.E.POST("/api/v1/users/create", c.createUser)

	routes := c.E.Group("/api/v1/users")
	// jwt middleware for these routes, you must be authorized!
	routes.Use(middleware.JWT([]byte(utils.JwtSecret)))
	{
		routes.GET("", c.getUser)
	}
}

func (c *UserController) getUser(e echo.Context) error {
	user := models.User{}

	// get user id from jwt
	userJWT := e.Get("user").(*jwt.Token)
	claims := userJWT.Claims.(jwt.MapClaims)
	id := claims["id"].(string)

	// set user ID
	user.ID = bson.ObjectIdHex(id)

	// find user by id
	err := user.FindByID()
	if err != nil {
		return e.JSON(http.StatusBadRequest, models.ResponseError{Error: err.Error()})
	}

	// populate vehicles slice with vehicle information
	for i, v := range user.Vehicles {
		vehicle := models.Vehicle{}
		vehicle.ID = v.(bson.ObjectId)
		err = vehicle.FindByID()
		if err != nil {
			return e.JSON(http.StatusBadRequest, models.ResponseError{Error: err.Error()})
		}
		user.Vehicles[i] = vehicle
	}

	// populate trips slice with trip information
	for i, v := range user.Trips {
		trip := models.Trip{}
		trip.ID = v.(bson.ObjectId)
		err = trip.FindByID()
		if err != nil {
			return e.JSON(http.StatusBadRequest, models.ResponseError{Error: err.Error()})
		}
		user.Trips[i] = trip
	}

	// return user
	return e.JSON(http.StatusOK, user)
}

// Creates user if not exists and generates jwt for user and returns the user and jwt
func (c *UserController) createUser(e echo.Context) error {
	user := models.User{}

	// bind request data to user struct
	err := e.Bind(&user)
	if err != nil {
		return e.JSON(http.StatusBadRequest, models.ResponseError{Error: err.Error()})
	}

	// create user in db
	err = user.Create()
	if err != nil {
		return e.JSON(http.StatusBadRequest, models.ResponseError{Error: err.Error()})
	}

	// find user in db by email, also generates jwt for Token field
	// password is set to ""
	err = user.FindByEmail()
	if err != nil {
		return e.JSON(http.StatusBadRequest, models.ResponseError{Error: err.Error()})
	}

	// populate vehicles slice with vehicle information
	for i, v := range user.Vehicles {
		vehicle := models.Vehicle{}
		vehicle.ID = v.(bson.ObjectId)
		err = vehicle.FindByID()
		if err != nil {
			return e.JSON(http.StatusBadRequest, models.ResponseError{Error: err.Error()})
		}
		user.Vehicles[i] = vehicle
	}

	// populate trips slice with trip information
	for i, v := range user.Trips {
		trip := models.Trip{}
		trip.ID = v.(bson.ObjectId)
		err = trip.FindByID()
		if err != nil {
			return e.JSON(http.StatusBadRequest, models.ResponseError{Error: err.Error()})
			fmt.Println(err)
		}
		user.Trips[i] = trip
	}

	// return user
	return e.JSON(http.StatusOK, user)

}
