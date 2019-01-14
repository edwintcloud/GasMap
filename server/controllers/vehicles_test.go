package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/edwintcloud/GasMap/server/models"
	"github.com/edwintcloud/GasMap/server/utils"
	"github.com/globalsign/mgo/bson"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

var testVehicleJSON = `{
	"year":"2012",
	"make":"Pontiac",
	"model":"Firebird",
	"mpg":"23",
	"tankSize":"20",
	"fuelQuality":"Regular"
}`

func TestCreateVehicle(t *testing.T) {
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
	req := httptest.NewRequest(http.MethodPost, "/api/v1/vehicles", strings.NewReader(testVehicleJSON))

	// set req headers to application/json
	req.Header.Set("Content-Type", "application/json")

	// create new recorder
	rec := httptest.NewRecorder()

	// create new echo context with request and recorder
	c := e.NewContext(req, rec)

	// decode testJWT into token
	token, err := jwt.Parse(testJWT, func(t *jwt.Token) (interface{}, error) {
		// Check the signing method
		if t.Method.Alg() != "HS256" {
			return nil, fmt.Errorf("unexpected jwt signing method=%v", t.Header["alg"])
		}
		return []byte(utils.JwtSecret), nil
	})
	if err != nil {
		t.Fatal("Unable to parse jwt: ", err)
	}

	// Store user information from token into context.
	c.Set("user", token)

	// create new instance of our vehicle controller and register with echo
	h := VehicleController{E: e}
	h.Register()

	// create the vehicle using the handler
	ok = assert.NoError(t, h.createVehicle(c))
	if !ok {
		t.Fatal("Handler failure")
	}

	// did we get a http status ok?
	ok = assert.Equal(t, http.StatusOK, rec.Code, "Expected http status to be 200 OK")
	if !ok {
		t.FailNow()
	}

	// read response body into vehicle struct
	vehicle := models.Vehicle{}
	body, _ := ioutil.ReadAll(rec.Result().Body)
	json.Unmarshal(body, &vehicle)

	// does the model returned match the model we created?
	ok = assert.Equal(t, vehicle.Model, "Firebird", "Expected returned vehicle to match created vehicle")
	if !ok {
		t.FailNow()
	}

	// get current user from db
	claims := token.Claims.(jwt.MapClaims)
	user := models.User{ID: bson.ObjectIdHex(claims["id"].(string))}
	user.FindByID()

	// did the vehicle get added to current user?
	ok = assert.Equal(t, vehicle.ID.Hex(), user.Vehicles[0].(bson.ObjectId).Hex(), "Expected vehicle to be added to current user")
	if !ok {
		t.FailNow()
	}

	// delete test user and vehicle when we are done
	user.RemoveByID()
	vehicle.RemoveByID()
}
