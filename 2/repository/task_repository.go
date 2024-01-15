package repository

import "database/sql"

type TaskRepository interface {
	Create(task *Task) error
	Read(id int) (*Task, error)
	Update(task *Task) error
	Delete(id int) error
}

type taskRepositoryImpl struct {
	db *sql.DB
}

func NewTaskRepository(db *sql.DB) TaskRepository {
	return &taskRepositoryImpl{db: db}
}
