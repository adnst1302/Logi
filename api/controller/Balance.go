package controller

import (
	"Logi/api/service"
	"github.com/labstack/echo/v4"
	"net/http"
)

func CheckBalance(c echo.Context) error {
	respCheckBalance := service.CheckBalance(c)
	return c.JSON(http.StatusOK, respCheckBalance)
}
