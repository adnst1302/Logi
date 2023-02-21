package controller

import (
	"Logi/api/model/resp"
	"Logi/api/model/scm"
	"Logi/api/service"
	"Logi/app"
	"github.com/labstack/echo/v4"
	"net/http"
)

func CheckBalance(c echo.Context) error {
	respCheckBalance := service.CheckBalance(c)
	return c.JSON(http.StatusOK, respCheckBalance)
}

func GetAllMembers(c echo.Context) error {
	var members []scm.Member
	app.Instance.Find(&members)

	respGetAllMember := resp.GetAllMember{
		Success: true,
		Code:    200,
		Data:    members,
	}

	return c.JSON(http.StatusOK, respGetAllMember)
}
