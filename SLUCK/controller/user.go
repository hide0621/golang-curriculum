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

	/*
		main関数内でe.Validatorを作成しているのでctx.Validateメソッドが使える
		実行されるのはCustomValidatorのValidateメソッド
		Error()をつけることでエラーの内容を文字列で返す
		例："code=400, message=Key: 'UserRequest.Name' Error:Field validation for 'Name' failed on the 'required' tag"
	*/
	if err := ctx.Validate(req); err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}

	// ここでコントローラー層からモデル層への変換を行う（DTOのような役割）
	u := toModel(req)

	// ここでユースケース層の関数を呼び出す
	c.u.Create(ctx.Request().Context(), u)

	return nil

}
