package controller

import (
	"fmt"
	"net/http"
	"sluck/usecase" // ディレクトリ名ではなくモジュール名でパスを指定する
	"strconv"

	"github.com/labstack/echo/v4"
)

// コントローラー層だけechoのcontextを使う
type UserController interface {
	Get(ctx echo.Context) error
	Create(ctx echo.Context) error
	Update(ctx echo.Context) error
	Delete(ctx echo.Context) error
}

type userController struct {
	u usecase.UserUsecase
}

func NewUserController(u usecase.UserUsecase) UserController {
	return &userController{u: u}
}

func (c *userController) Get(ctx echo.Context) error {

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		msg := fmt.Errorf("parse error: %v", err.Error())
		// curl http://localhost:8080/tasks/hoge とすると "parse error: strconv.Atoi: parsing \"hoge\": invalid syntax" と返ってくる
		return ctx.JSON(http.StatusBadRequest, msg.Error())
	}

	u, err := c.u.GetByID(ctx.Request().Context(), id)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}
	return ctx.JSON(http.StatusOK, u)

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

func (c *userController) Update(ctx echo.Context) error {

	var req UserRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}

	if err := ctx.Validate(req); err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}

	u := toModel(req)

	c.u.Update(ctx.Request().Context(), u)

	return nil

}

func (c *userController) Delete(ctx echo.Context) error {

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		msg := fmt.Errorf("parse error: %v", err.Error())
		// curl http://localhost:8080/tasks/hoge とすると "parse error: strconv.Atoi: parsing \"hoge\": invalid syntax" と返ってくる
		return ctx.JSON(http.StatusBadRequest, msg.Error())
	}

	if err := c.u.Delete(ctx.Request().Context(), id); err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}

	return nil

}
