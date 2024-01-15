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

	stmt := `INSERT INTO tasks (title) VALUES (?) RETURNING id`

	err := r.db.QueryRow(stmt, task.Title).Scan(&task.ID)

	return task.ID, err

}

func (r *taskRepositoryImpl) Read(id int) (*Task, error) {

	stmt := `SELECT id, title FROM tasks WHERE id = ?`

	task := Task{}

	err := r.db.QueryRow(stmt, id).Scan(&task.ID, &task.Title)

	return &task, err

}

func (r *taskRepositoryImpl) Update(task *Task) error {

	stmt := `UPDATE tasks SET title = ? WHERE id = ?`

	rows, err := r.db.Exec(stmt, task.Title, task.ID)
	if err != nil {
		return err
	}

	rowsAffected, err := rows.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return sql.ErrNoRows
	}

	return err

}
