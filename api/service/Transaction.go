package service

import (
	"Logi/api/model/pyl"
	"Logi/api/model/resp"
	"Logi/api/model/scm"
	"Logi/app"
	"Logi/helper"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"time"
)

func AddTransaction(c echo.Context) resp.AddTransaction {
	pylAddTransaction := new(pyl.AddTransaction)
	err := c.Bind(pylAddTransaction)
	helper.Pie(err)

	checkMember := app.CheckMember(pylAddTransaction.UserId)
	respAddTransaction := resp.AddTransaction{}

	if checkMember == false {
		respAddTransaction.Code = 400
		respAddTransaction.Message = "Failed add transaction, user id not found!!!"
		respAddTransaction.Success = false
	} else {
		trxId := uuid.New().String()
		scmBalTrans := scm.BalanceTransaction{
			TrxId:     trxId,
			UserId:    pylAddTransaction.UserId,
			Amount:    pylAddTransaction.Amount,
			Flow:      pylAddTransaction.Flow,
			TypeTrans: pylAddTransaction.TypeTrans,
			CreateBy:  pylAddTransaction.CreateBy,
			ApproveBy: pylAddTransaction.ApproveBy,
			Desc:      pylAddTransaction.Desc,
			TrxTime:   time.Now(),
		}
		app.Instance.Create(&scmBalTrans)
		respAddTransaction.Code = 200
		respAddTransaction.Message = "Success add transaction with trxId." + trxId
		respAddTransaction.Success = true
	}

	return respAddTransaction
}
