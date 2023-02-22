package service

import (
	"Logi/api/model/pyl"
	"Logi/api/model/resp"
	"Logi/helper"
	"github.com/labstack/echo/v4"
)

func CheckBalance(c echo.Context) resp.CheckBalance {
	pylCheckBalance := new(pyl.CheckBalance)
	err := c.Bind(pylCheckBalance)
	helper.Pie(err)

	//var balance int
	respBalance := resp.CheckBalance{
		Success: true,
		Code:    200,
		Balance: 900,
	}

	return respBalance
}

func AddTransaction(c echo.Context) resp.AddTransaction {
	pylAddTransaction := new(pyl.AddTransaction)
	err := c.Bind(pylAddTransaction)
	helper.Pie(err)

	respAddTransaction := resp.AddTransaction{
		Success: true,
		Code:    200,
		Message: "Success add transaction with trxId",
	}
	return respAddTransaction
}
