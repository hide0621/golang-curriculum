package model

import "time"

// 講師は実装していなかったが必要だと思われるので追加した
type Message struct {
	ID        int       `json:"id"`
	Content   string    `json:"content"`
	UserID    int       `json:"user_id"`
	ChannelID int       `json:"channel_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeleteAt  time.Time `json:"delete_at"`
}
