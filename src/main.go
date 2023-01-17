package main

import (
	"github.com/labstack/echo/v4"
	"henotech.net/di"
	"henotech.net/util"
)

func main() {
	db := util.NewDB()

	taskHandler := di.Init(db)

	e := echo.New()

	e.POST("/task", taskHandler.Post())
	e.GET("/task", taskHandler.FindAll())
	e.GET("/task/:id", taskHandler.Get())
	e.PUT("/task/:id", taskHandler.Put())
	e.DELETE("/task/:id", taskHandler.Delete())

	e.Logger.Fatal(e.Start(":8080"))
}
