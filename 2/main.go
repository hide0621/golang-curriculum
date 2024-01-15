package main

import (
	"2/controller"
	"database/sql"
	"fmt"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func initDB() (*sql.DB, error) {

	db, err := sql.Open("sqlite3", "./test.db")
	return db, err

}

func main() {

	db, err := initDB()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(db)

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// メソッドのレシーバーがポインタ型で定義されている場合、呼び出し時にそのオブジェクト（レシーバー）が値型であっても呼び出すことは可能。
	// レシーバーが値型で定義されていて、呼び出し時はポインタ型のオブジェクト（レシーバー）でも同様に呼び出せる
	// ただし揃えないと気持ち悪いし、挙動の誤解を起こしそう...

	// taskController := controller.TaskController{}

	taskController := &controller.TaskController{}

	e.GET("/tasks", taskController.Get)

	e.POST("/tasks", taskController.Create)

	e.Start(":8080")

}
