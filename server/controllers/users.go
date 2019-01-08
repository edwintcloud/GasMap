package controllers

import (
	"net/http"

	"github.com/edwintcloud/GasMap/server/models"
	"github.com/labstack/echo"
)

// UserController is our user controller struct
type UserController struct {
	E *echo.Echo
}

// Register registers our user controller routes
func (c *UserController) Register() {

	routes := c.E.Group("/api/v1/users")
	{
		routes.GET("", c.Get)
	}
}

// Get is our user controller get route
func (c *UserController) Get(e echo.Context) error {
	return e.JSON(http.StatusOK, models.ResponseMsg{Message: "Hello"})
}
