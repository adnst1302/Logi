package cmd

import (
	"Logi/api/controller"
	"Logi/app"
	"github.com/labstack/echo/v4"
	"log"
)

func Execute() {
	log.Print("Starting system ....")

	app.Connect()
	app.Migrate()

	rt := echo.New()
	rtV1 := rt.Group("/v1")
	// member
	rtV1.POST("/member/create", controller.CreateMember)
	rtV1.GET("/member/list", controller.GetAllMembers)
	// balance
	rtV1.POST("/balance/check", controller.CheckBalance)

	rt.Logger.Fatal(rt.Start(":8001"))
}
