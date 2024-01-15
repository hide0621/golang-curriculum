package controller

import (
	"2/model"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type TaskController struct {
}

func (t *TaskController) Get(c echo.Context) error {

	// tasks, err := usecase.GetTasks()

	return c.JSON(http.StatusOK, nil)

}

func (t *TaskController) Create(c echo.Context) error {

	var task model.Task

	if err := c.Bind(&task); err != nil {
		return c.JSON(http.StatusBadRequest, nil)
	}

	fmt.Println(task.Title)

	// created, err := usecase.CreateTask(task)
	// if err != nil {
	// 	return c.JSON(http.StatusInternalServerError, nil)
	// }

	// return c.JSON(http.StatusOK, created)

	return c.JSON(http.StatusOK, nil)

}
