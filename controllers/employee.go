package controllers

import (
	"github.com/ilhamarrouf/echo-example/models"
	"github.com/labstack/echo"
	"net/http"
)

func GetEmployees(c echo.Context) error {
	employees := models.GetEmployee()

	return c.JSON(http.StatusOK, employees)
}