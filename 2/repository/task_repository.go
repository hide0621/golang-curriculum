package repository

import "database/sql"

type TaskRepository interface {
	Create(task *Task) (int, error)
	Read(id int) (*Task, error)
	Update(task *Task) error
	Delete(id int) error
}

type taskRepositoryImpl struct {
	db *sql.DB
}

// コントローラー層に実装されていて２重実装になるが、一旦ここにも実装
type Task struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

func NewTaskRepository(db *sql.DB) TaskRepository {
	return &taskRepositoryImpl{db: db}
}

func (r *taskRepositoryImpl) Create(task *Task) (int, error) {

	// ここでSQLを実行
	stmt := `INSERT INTO tasks (title) VALUES (?) RETURNING id`
	err := r.db.QueryRow(stmt, task.Title).Scan(&task.ID)

	return task.ID, err

}
