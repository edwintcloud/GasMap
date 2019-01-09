package main

import (
	"fmt"
	"net/http"

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

	// Create new instance of echo
	e := echo.New()

	// Connect to mongodb or panic
	session, err := utils.ConnectToDb(viper.GetString("database.url"), viper.GetString("database.name"))
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
	e.Logger.Fatal(e.Start(":9000"))

}
