package controllers

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/edwintcloud/GasMap/server/models"
	"github.com/labstack/echo"
)

var (
	testUser = models.User{
		Email:     "test@email.com",
		FirstName: "test",
		LastName:  "user",
	}
	testUserJSON = `{"email":"test@email.com","firstName":"test","lastName":"user"}`
)

func TestCreateUser(t *testing.T) {

	// create new echo instance
	e := echo.New()

	// create a new request
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(testUserJSON))

	// set req headers to application/json
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	// create new recorder
	rec := httptest.NewRecorder()

	// create new echo context with request and recorder
	c := e.NewContext(req, rec)
	c.SetPath("/api/v1/users/create")

	// create new instance of our user controller and register with echo
	h := &UserController{E: e}
	h.Register()

	// this throws the error **
	h.CreateUser(c)
	// make our assertions using the createUser method
	// if assert.NoError(t, h.CreateUser(c)) {
	// 	// did we get a http status ok?
	// 	assert.Equal(t, http.StatusOK, rec.Code)
	// }
}
