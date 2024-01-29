package repository

import (
	"context"
	"database/sql"
)

type MessageRepository interface {
	Delete(ctx context.Context, userID int) error
}

type messageRepository struct {
	db *sql.DB
}

func NewMessageRepository(db *sql.DB) MessageRepository {
	return &messageRepository{db: db}
}

func (r *messageRepository) Delete(ctx context.Context, userID int) error {

	_, err := r.db.ExecContext(ctx, "DELETE FROM message WHERE user_id = ?", userID)
	if err != nil {
		return err
	}

	return nil

}
