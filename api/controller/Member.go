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

func UpdateProfileMember(c echo.Context) error {
	respUpdateProfileMember := service.UpdateProfileMember(c)
	return c.JSON(http.StatusOK, respUpdateProfileMember)
}

func DeleteMember(c echo.Context) error {
	respDeleteMember := service.DeleteMember(c)
	return c.JSON(http.StatusOK, respDeleteMember)
}

func DetailMember(c echo.Context) error {
	respDetailMember := service.DetailMember(c)
	return c.JSON(http.StatusOK, respDetailMember)
}
