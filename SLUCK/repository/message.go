package repository

import (
	"context"
	"database/sql"
	"fmt"
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

	/*
		トランザクション（今回はtxトランザクション）の中のDBオブジェクトを使用する
		こうすることで同一トランザクション内で複数のDB操作を行うことができる
		例：ユーザーを削除すると同時にそのユーザーのメッセージも削除する
	*/
	db, ok := GetTx(ctx)
	if !ok {
		return fmt.Errorf("failed to get tx. transaction tx is not found")
	}

	_, err := db.Exec("DELETE FROM message WHERE user_id = ?", userID)
	if err != nil {
		return err
	}

	return nil

}
