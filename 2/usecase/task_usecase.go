package usecase

import "2/repository"

type TaskUsecase interface {
	CreateTask(title string) error
	GetTask(id int) (Task, error)
	UpdateTask(id int, title string) error
	DeleteTask(id int) error
}

type taskUsecase struct {
	r repository.TaskRepository
}

func NewTaskUsecase(r repository.TaskRepository) TaskUsecase {
	return &taskUsecase{r: r}
}
