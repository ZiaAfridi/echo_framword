package main

import (
	"echo/controllers"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	// routes
	e.GET("/v1", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	// route
	e.GET("/v1/user/:id", controllers.GetUser)

	e.Logger.Fatal(e.Start(":5000"))
}
