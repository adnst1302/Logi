package controller

import (
	"Logi/api/service"
	"github.com/labstack/echo/v4"
	"net/http"
)

func CreateMember(c echo.Context) error {
	respServCreateMember := service.CreateMember(c)
	return c.JSON(http.StatusOK, respServCreateMember)
}

func GetAllMembers(c echo.Context) error {
	respServGetAllMember := service.GetAllMembers(c)
	return c.JSON(http.StatusOK, respServGetAllMember)
}
