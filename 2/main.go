package main

import (
	"2/controller"
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	taskController := &controller.TaskController{}

	e.GET("/tasks", taskController.Get)

	e.POST("/tasks", func(c echo.Context) error {
		fmt.Println("create task")
		return c.String(200, "create task")
	})

	e.Start(":8080")

}
