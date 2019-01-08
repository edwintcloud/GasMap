package main

import (
	"net/http"

	"github.com/labstack/echo"
)

func main() {

	// Create new instance of echo
	e := echo.New()

	// test route
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello")
	})

	// Start http server, wrap with fatal helper func
	e.Logger.Fatal(e.Start(":9000"))

}
