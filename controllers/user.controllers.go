package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// GetUser godoc
// @Summary Show the status of server.
// @Description get the status of server.
// @Tags root
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router / [get]
func GetUser(c echo.Context) error {
	// User ID from path `users/:id`
	id := c.Param("id")
	return c.String(http.StatusOK, id)
}

func GetAllUser(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!\nAll Users")
}
