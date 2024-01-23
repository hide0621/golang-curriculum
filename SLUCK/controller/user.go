package controller

import (
	"net/http"
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

	var req UserRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}

	// ここでコントローラー層からモデル層への変換を行う（DTOのような役割）
	u := toModel(req)

	// ここでユースケース層の関数を呼び出す
	c.u.Create(ctx.Request().Context(), u)

	return nil

}
