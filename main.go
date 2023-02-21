package main

import (
	"Logi/apps"
	"Logi/controller"
	"github.com/labstack/echo/v4"
)

func main() {
	apps.Connect()
	apps.Migrate()

	rt := echo.New()
	rtV1 := rt.Group("/v1")
	// member
	
	// balance
	rtV1.POST("/balance/check", controller.CheckBalance)

	rt.Logger.Fatal(rt.Start(":8001"))
}
