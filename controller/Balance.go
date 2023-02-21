package controller

import (
	"Logi/helper"
	"Logi/model/pyl"
	"Logi/model/resp"
	"github.com/labstack/echo/v4"
	"net/http"
)

func CheckBalance(c echo.Context) error {
	pylCheckBalance := new(pyl.CheckBalance)
	err := c.Bind(pylCheckBalance)

	helper.Pie(err)

	//var balance int

	balance := 0
	if pylCheckBalance.UserId == "adit" {
		balance = 40
	}

	respBalance := resp.CheckBalance{
		Success: true,
		Code:    200,
		Balance: float32(balance),
		Error:   false,
	}
	return c.JSON(http.StatusOK, respBalance)
}
