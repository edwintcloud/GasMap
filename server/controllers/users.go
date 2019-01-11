package controllers

import (
	"net/http"

	"github.com/edwintcloud/GasMap/server/models"
	"github.com/edwintcloud/GasMap/server/utils"
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
	c.E.POST("/api/v1/users/create", c.post)

	routes := c.E.Group("/api/v1/users")
	// jwt middleware for these routes, you must be authorized!
	routes.Use(middleware.JWT([]byte(utils.JwtSecret)))
	{
		routes.GET("", c.get)
	}
}

func (c *UserController) get(e echo.Context) error {
	return e.JSON(http.StatusOK, models.ResponseMsg{Message: "Hello"})
}

func (c *UserController) post(e echo.Context) error {
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

	// return user
	return e.JSON(http.StatusOK, user)

}
