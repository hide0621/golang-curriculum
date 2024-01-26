package main

import (
	"net/http"
	"sluck/controller"

	"sluck/infra"
	"sluck/repository"
	"sluck/usecase"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

// echoのバリデーションをgo-playground/validatorを使ったカスタムバリデーションにする
// シングルトンインスタンスで使う
type CustomValidator struct {
	validator *validator.Validate
}

// echoのValidateメソッドをWrapしたメソッド(echoのものと紐付く構造体は違うが同じシグネチャなのでダックタイピングができる)
func (cv *CustomValidator) Validate(i any) error {

	if err := cv.validator.Struct(i); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil

}

func main() {

	e := echo.New()

	e.Validator = &CustomValidator{validator: validator.New()}

	db := infra.Connect()

	defer db.Close()

	ur := repository.NewUserRepository(db)
	uu := usecase.NewUserUsecase(ur)
	uc := controller.NewUserController(uu)

	e.POST("/users", uc.Create)
	e.GET("/users/:id", uc.Get)
	e.PUT("/users", uc.Update)
	e.DELETE("/users/:id", uc.Delete)

	e.Start(":8080")

}
