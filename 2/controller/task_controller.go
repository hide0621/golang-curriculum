package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type TaskController struct {
}

func (t *TaskController) Get(c echo.Context) error {

	// tasks, err := usecase.GetTasks()

	return c.JSON(http.StatusOK, nil)

}
