package controller

import (
	"2/model"
	"2/usecase"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type TaskController interface {
	Get(c echo.Context) error
	Create(c echo.Context) error
}

type taskController struct {
	u usecase.TaskUsecase
}

func NewTaskController(u usecase.TaskUsecase) TaskController {
	return &taskController{u: u}
}

func (t *taskController) Get(c echo.Context) error {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		msg := fmt.Errorf("parse error: %v", err.Error())
		// curl http://localhost:8080/tasks/hoge とすると "parse error: strconv.Atoi: parsing \"hoge\": invalid syntax" と返ってくる
		return c.JSON(http.StatusBadRequest, msg.Error())
	}

	task, err := t.u.GetTask(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, task)

}

func (t *taskController) Create(c echo.Context) error {

	var task model.Task

	if err := c.Bind(&task); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	createdID, err := t.u.CreateTask(task.Title)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, createdID)

}
