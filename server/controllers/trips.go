package controllers

import (
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/edwintcloud/GasMap/server/models"
	"github.com/edwintcloud/GasMap/server/utils"
	"github.com/globalsign/mgo/bson"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// TripController is our trip controller struct
type TripController struct {
	E *echo.Echo
}

// Register registers our trip controller routes with echo
func (c *TripController) Register() {

	routes := c.E.Group("/api/v1/trips")
	// jwt middleware for these routes, you must be authorized!
	routes.Use(middleware.JWT([]byte(utils.JwtSecret)))
	{
		routes.POST("", c.createTrip)
		routes.DELETE("", c.deleteTrip)
	}
}

func (c *TripController) createTrip(e echo.Context) error {
	trip := models.Trip{}
	user := models.User{}

	// get user id from jwt
	userJWT := e.Get("user").(*jwt.Token)
	claims := userJWT.Claims.(jwt.MapClaims)
	id := claims["id"].(string)

	// set user ID
	user.ID = bson.ObjectIdHex(id)

	// bind request data to trip struct
	err := e.Bind(&trip)
	if err != nil {
		return e.JSON(http.StatusBadRequest, models.ResponseError{Error: err.Error()})
	}

	// create trip in db
	err = trip.Create()
	if err != nil {
		return e.JSON(http.StatusBadRequest, models.ResponseError{Error: err.Error()})
	}

	// find user by id
	err = user.FindByID()
	if err != nil {
		return e.JSON(http.StatusBadRequest, models.ResponseError{Error: err.Error()})
	}

	// add trip to user in db
	err = user.AddTrip(&trip)
	if err != nil {
		return e.JSON(http.StatusBadRequest, models.ResponseError{Error: err.Error()})
	}

	// return trip
	return e.JSON(http.StatusOK, trip)
}

func (c *TripController) deleteTrip(e echo.Context) error {
	trip := models.Trip{}
	user := models.User{}

	// get user id from jwt
	userJWT := e.Get("user").(*jwt.Token)
	claims := userJWT.Claims.(jwt.MapClaims)
	id := claims["id"].(string)

	// set user ID
	user.ID = bson.ObjectIdHex(id)

	// get trip id from query param
	tripID := e.QueryParam("id")
	if !bson.IsObjectIdHex(tripID) {
		return e.JSON(http.StatusBadRequest, models.ResponseError{Error: "Invalid object id"})
	}
	trip.ID = bson.ObjectIdHex(tripID)

	// delete trip from db
	err := trip.RemoveByID()
	if err != nil {
		return e.JSON(http.StatusBadRequest, models.ResponseError{Error: err.Error()})
	}

	// find user by id
	err = user.FindByID()
	if err != nil {
		return e.JSON(http.StatusBadRequest, models.ResponseError{Error: err.Error()})
	}

	// remove trip from current user trips slice
	err = user.RemoveTrip(&trip)
	if err != nil {
		return e.JSON(http.StatusBadRequest, models.ResponseError{Error: err.Error()})
	}

	// if all went well, return success and http status 200 OK
	return e.JSON(http.StatusOK, models.ResponseMsg{Message: "OK"})
}
