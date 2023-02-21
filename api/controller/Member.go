package controller

import (
	"Logi/api/service"
	"github.com/labstack/echo/v4"
	"net/http"
)

func CreateMember(c echo.Context) error {
	respCreateMember := service.CreateMember(c)
	return c.JSON(http.StatusOK, respCreateMember)
}
