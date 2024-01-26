package repository

import (
	"context"
	"database/sql"
	"fmt"
	"sluck/model"
	"strconv"
)

type UserRepository interface {
	Create(ctx context.Context, user *model.User) (string, error)
	Read(ctx context.Context, id int) (*model.User, error)
	Update(ctx context.Context, user *model.User) error
	Delete(ctx context.Context, id int) error
}

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Create(ctx context.Context, user *model.User) (string, error) {

	result, err := r.db.Exec("INSERT INTO users (name, email, age) VALUES (?, ?, ?)", user.Name, user.Email, user.Age)
	if err != nil {
		return "", err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return "", err
	}

	idStr := strconv.FormatInt(id, 10)

	return idStr, nil

}

func (r *userRepository) Read(ctx context.Context, id int) (*model.User, error) {

	var user model.User

	err := r.db.QueryRow("SELECT id, name, email, age FROM users WHERE id = ?", id).Scan(&user.ID, &user.Name, &user.Email, &user.Age)
	if err != nil {
		return nil, err
	}

	return &user, nil

}

func (r *userRepository) Update(ctx context.Context, user *model.User) error {

	result, err := r.db.Exec("UPDATE users SET name = ?, email = ?, age = ? WHERE id = ?", user.Name, user.Email, user.Age, user.ID)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("no rows affected for id: %d", user.ID)
	}

	return nil

}

func (r *userRepository) Delete(ctx context.Context, id int) error {

	result, err := r.db.Exec("DELETE FROM users WHERE id = ?", id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("no rows affected for id: %d", id)
	}

	return nil

}
