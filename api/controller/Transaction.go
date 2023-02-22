package controller

import (
	"Logi/api/service"
	"github.com/labstack/echo/v4"
	"net/http"
)

func AddTransaction(c echo.Context) error {
	respAddTransaction := service.AddTransaction(c)
	return c.JSON(http.StatusOK, respAddTransaction)
}
