package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/edwintcloud/GasMap/server/models"
	"github.com/edwintcloud/GasMap/server/utils"
	"github.com/labstack/echo"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

var (
	testUser = models.User{
		Email:     "test@email.com",
		FirstName: "test",
		LastName:  "user",
	}
	testUserJSON = `{"email":"test@email.com","firstName":"test","lastName":"user"}`
	db           = dbInfo{}
)

type dbInfo struct {
	url  string
	name string
}

// get db connection info before running tests
func init() {

	// set config file for viper to load
	viper.SetConfigFile(`../config.json`)

	// load config file, panic on error
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	db.url = viper.GetString("database.url")
	db.name = viper.GetString("database.name")

}

func TestCreateUser(t *testing.T) {
	var ok bool

	// Connect to mongodb or die
	session, err := utils.ConnectToDb(db.url, db.name)
	if err != nil {
		t.Fatal("Unable to connect to db: ", err)
	}

	// close db session when test finishes
	defer session.Close()

	// create new echo instance
	e := echo.New()

	// create a new request
	req := httptest.NewRequest(http.MethodPost, "/api/v1/users/create", strings.NewReader(testUserJSON))

	// set req headers to application/json
	req.Header.Set("Content-Type", "application/json")

	// create new recorder
	rec := httptest.NewRecorder()

	// create new echo context with request and recorder
	c := e.NewContext(req, rec)

	// create new instance of our user controller and register with echo
	h := UserController{E: e}
	h.Register()

	// create the user using the handler
	ok = assert.NoError(t, h.createUser(c))
	if !ok {
		t.Fatal("Handler failure")
	}

	// did we get a http status ok?
	ok = assert.Equal(t, http.StatusOK, rec.Code, "Expected http status to be 200 OK")
	if !ok {
		t.FailNow()
	}

	// read response body into user struct
	user := models.User{}
	body, _ := ioutil.ReadAll(rec.Result().Body)
	json.Unmarshal(body, &user)

	// delete test user when test finishes
	defer user.RemoveByID()

	// did we get a token?
	ok = assert.NotEmpty(t, user.Token, "Expected Token to not be Empty")
	if !ok {
		t.FailNow()
	}

}
