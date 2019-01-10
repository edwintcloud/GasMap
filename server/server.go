package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/edwintcloud/GasMap/server/controllers"
	"github.com/edwintcloud/GasMap/server/models"
	"github.com/edwintcloud/GasMap/server/utils"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/spf13/viper"
)

func init() {

	// set config file for viper to load
	viper.SetConfigFile(`config.json`)

	// load config file, panic on error
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
}

func main() {
	dbURL := viper.GetString("database.url")
	dbName := viper.GetString("database.name")
	port := viper.GetString("server.address")

	// Create new instance of echo
	e := echo.New()

	// if mongodb_uri is an env var, use that to connect instead, and set our port
	if v, ok := os.LookupEnv("MONGODB_URI"); ok {
		dbURL = v
		database := strings.Split(v, "/")
		dbName = database[len(database)-1]
		port = fmt.Sprintf(":%s", os.Getenv("PORT"))
	}

	// Connect to mongodb or panic
	session, err := utils.ConnectToDb(dbURL, dbName)
	if err != nil {
		panic(err)
	}

	// defer mongo session to close when app closes
	defer session.Close()

	// register our middlewares
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${method}  ${uri}  ${latency_human}  ${status}\n",
	}))

	// Register controller routes with echo
	userController := controllers.UserController{E: e}
	userController.Register()

	// catch all route
	e.Any("*", func(c echo.Context) error {
		err := fmt.Sprintf("Bad Request - %s %s", c.Request().Method, c.Request().RequestURI)
		return c.JSON(http.StatusBadRequest, models.ResponseError{Error: err})
	})

	// Start http server, wrap with fatal helper func

	e.Logger.Fatal(e.Start(port))

}
