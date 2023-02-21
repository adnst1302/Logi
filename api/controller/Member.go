package controller

import (
	"Logi/api/model/pyl"
	"Logi/api/model/resp"
	"Logi/api/model/scm"
	"Logi/app"
	"Logi/helper"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
)

func CreateMember(c echo.Context) error {
	pylCreateMember := new(pyl.CreateMember)
	err := c.Bind(pylCreateMember)
	helper.Pie(err)

	respCreateMember := resp.CreateMember{}
	respSucessCreateMember := resp.SuccessCreateMember{}

	scmCreateMember := scm.Member{
		UserId:   pylCreateMember.UserId,
		Email:    pylCreateMember.Email,
		Role:     pylCreateMember.Role,
		Password: pylCreateMember.Password,
	}

	respSucessCreateMember.Message = "Success create member"

	app.Instance.Create(&scmCreateMember)

	respCreateMember.Data = respSucessCreateMember
	respCreateMember.Code = 200
	respCreateMember.Success = true
	log.Print("Create member sukses")
	return c.JSON(http.StatusOK, respCreateMember)
}
