package controllers

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/edwintcloud/GasMap/server/models"
	"github.com/edwintcloud/GasMap/server/utils"
	"github.com/labstack/echo"
	uuid "github.com/satori/go.uuid"
)

// UserController is our user controller struct
type UserController struct {
	E *echo.Echo
}

// Register registers our user controller routes
func (c *UserController) Register() {

	routes := c.E.Group("/api/v1/users")
	{
		routes.GET("", c.get)
		routes.GET("/auth/google/login", c.googleLogin)
		routes.GET("/auth/google/callback", c.googleCallback)
	}
}

// Get is our user controller get route
func (c *UserController) get(e echo.Context) error {
	return e.JSON(http.StatusOK, models.ResponseMsg{Message: "Hello"})
}

// GoogleLogin is our user controller google login route
func (c *UserController) googleLogin(e echo.Context) error {

	// generate random uuid to prevent csrf attacks
	state := uuid.NewV4().String()

	// generate url for Google consent page
	url := utils.GoogleOauth.AuthCodeURL(state)

	// redirect to url
	return e.Redirect(http.StatusTemporaryRedirect, url)
}

// GoogleCallback is our user controller google callback route
func (c *UserController) googleCallback(e echo.Context) error {

	// get code from query params
	code := e.QueryParam("code")

	// Handle the exchange code to initiate a transport.
	token, err := utils.GoogleOauth.Exchange(context.TODO(), code)
	if err != nil {
		return e.JSON(http.StatusBadRequest, models.ResponseError{Error: err.Error()})
	}

	// Construct the client
	client := utils.GoogleOauth.Client(context.TODO(), token)

	// Get user info
	resp, err := client.Get("https://www.googleapis.com/oauth2/v3/userinfo")
	if err != nil {
		return e.JSON(http.StatusBadRequest, models.ResponseError{Error: err.Error()})
	}

	// defer resp body to close when request finishes
	defer resp.Body.Close()

	// Read user info from resp body
	data, _ := ioutil.ReadAll(resp.Body)

	// unmarshal json resp.body into GoogleProfile struct
	profile := models.GoogleProfile{}
	err = json.Unmarshal([]byte(data), &profile)
	if err != nil {
		return e.JSON(http.StatusBadRequest, models.ResponseError{Error: err.Error()})
	}

	// create user in db
	user := models.User{
		FirstName: profile.FirstName,
		LastName:  profile.LastName,
		Email:     profile.Email,
		Password:  code,
	}
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
