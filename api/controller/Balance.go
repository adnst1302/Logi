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

func AddTransaction(c echo.Context) error {
	respAddTransaction := service.AddTransaction(c)
	return c.JSON(http.StatusOK, respAddTransaction)
}
