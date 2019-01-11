package controllers

import (
	"net/http"

	"github.com/globalsign/mgo/bson"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/edwintcloud/GasMap/server/utils"

	"github.com/edwintcloud/GasMap/server/models"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// VehicleController is our vehicle controller struct
type VehicleController struct {
	E *echo.Echo
}

// Register registers our vehicle controller routes
func (c *VehicleController) Register() {

	routes := c.E.Group("/api/v1/vehicles")
	// jwt middleware for these routes, you must be authorized!
	routes.Use(middleware.JWT([]byte(utils.JwtSecret)))
	{
		routes.GET("", c.get)
		routes.POST("", c.post)
	}
}

func (c *VehicleController) get(e echo.Context) error {
	user := e.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	id := claims["id"].(string)
	return e.JSON(http.StatusOK, models.ResponseMsg{Message: id})
}

func (c *VehicleController) post(e echo.Context) error {
	vehicle := models.Vehicle{}
	user := models.User{}

	// get user id from jwt
	userJWT := e.Get("user").(*jwt.Token)
	claims := userJWT.Claims.(jwt.MapClaims)
	id := claims["id"].(string)

	// set user ID
	user.ID = bson.ObjectIdHex(id)

	// bind request data to vehicle struct
	err := e.Bind(&vehicle)
	if err != nil {
		return e.JSON(http.StatusBadRequest, models.ResponseError{Error: err.Error()})
	}

	// create vehicle in db
	err = vehicle.Create()
	if err != nil {
		return e.JSON(http.StatusBadRequest, models.ResponseError{Error: err.Error()})
	}

	// find user in db by id
	err = user.FindByID()
	if err != nil {
		return e.JSON(http.StatusBadRequest, models.ResponseError{Error: err.Error()})
	}

	// add vehicle to user in db
	err = user.AddVehicle(&vehicle)
	if err != nil {
		return e.JSON(http.StatusBadRequest, models.ResponseError{Error: err.Error()})
	}

	// return vehicle
	return e.JSON(http.StatusOK, vehicle)

}
