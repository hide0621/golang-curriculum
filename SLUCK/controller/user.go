package controller

import (
	"fmt"
	"sluck/usecase" // ディレクトリ名ではなくモジュール名でパスを指定する

	"github.com/labstack/echo/v4"
)

// コントローラー層だけechoのcontextを使う
type UserController interface {
	Create(ctx echo.Context) error
}

type userController struct {
	u usecase.UserUsecase
}

func NewUserController(u usecase.UserUsecase) UserController {
	return &userController{u: u}
}

func (c *userController) Create(ctx echo.Context) error {
	fmt.Println("creating ... ")
	return nil
}
