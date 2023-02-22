package service

import (
	"Logi/api/model/pyl"
	"Logi/api/model/resp"
	"Logi/api/model/scm"
	"Logi/app"
	"Logi/helper"
	"github.com/labstack/echo/v4"
)

func CheckBalance(c echo.Context) resp.CheckBalance {
	pylCheckBalance := new(pyl.CheckBalance)
	err := c.Bind(pylCheckBalance)
	helper.Pie(err)
	checkMember := app.CheckMember(pylCheckBalance.UserId)
	respCheckBalance := resp.CheckBalance{}

	if checkMember == false {
		respCheckBalance.Balance = 0
		respCheckBalance.Code = 400
		respCheckBalance.Success = false
		respCheckBalance.Message = "Failed check balance, userid not found!!!"
	} else {

		respCheckBalance.Balance = SubGetBalanceMember(pylCheckBalance.UserId)
		respCheckBalance.Code = 200
		respCheckBalance.Success = true
		respCheckBalance.Message = "Check balance success."
	}

	return respCheckBalance
}

func SubGetBalanceMember(userId string) float64 {
	var balances []scm.BalanceTransaction
	app.Instance.Where("user_id = ?", userId).Find(&balances)

	var fb float64

	for x := 0; x < len(balances); x++ {
		if balances[x].Flow == "CREDIT" {
			fb += balances[x].Amount
		} else {
			fb -= balances[x].Amount
		}
	}
	return fb
}
