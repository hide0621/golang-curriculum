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

	// メソッドのレシーバーがポインタ型で定義されている場合、呼び出し時にそのオブジェクト（レシーバー）が値型であっても呼び出すことは可能。
	// レシーバーが値型で定義されていて、呼び出し時はポインタ型のオブジェクト（レシーバー）でも同様に呼び出せる
	// ただし揃えないと気持ち悪いし、挙動の誤解を起こしそう...

	// taskController := controller.TaskController{}

	taskController := &controller.TaskController{}

	e.GET("/tasks", taskController.Get)

	e.POST("/tasks", func(c echo.Context) error {
		fmt.Println("create task")
		return c.String(200, "create task")
	})

	e.Start(":8080")

}
