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
	rtV1.GET("/member/list", controller.GetAllMembers)
	rtV1.POST("/member/create", controller.CreateMember)
	rtV1.POST("/member/profile/update", controller.UpdateProfileMember)
	rtV1.POST("/member/delete", controller.DeleteMember)
	rtV1.POST("/member/detail", controller.DetailMember)
	// balance & transaction
	rtV1.POST("/transaction/add", controller.AddTransaction)
	rtV1.POST("/balance/check", controller.CheckBalance)

	rt.Logger.Fatal(rt.Start(":8001"))
}
