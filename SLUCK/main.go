package main

import (
	"fmt"
	"sluck/controller"
	"sluck/repository"
	"sluck/usecase"

	"github.com/labstack/echo/v4"
)

func main() {

	e := echo.New()

	fmt.Println("Hello, World!")

	ur := repository.NewUserRepository(nil)
	uu := usecase.NewUserUsecase(ur)
	uc := controller.NewUserController(uu)

	e.POST("/users", uc.Create)

	e.Start(":8080")

}
