package controller

import (
	"sluck/model"
	"time"
)

// コントローラー層で扱うデータ型(リクエストでさばくでデータ型)
type UserRequest struct {
	ID    int    `json:"id"`
	Name  string `json:"name" validate:"required"` // echoのバリデーションを使って、nameをリクエストで必須にする
	Age   int    `json:"age"`
	Email string `json:"email"`
}

// コントローラー層からモデル層への変換はコントローラー層で行う（DTOのような役割を持つ関数）
func toModel(req UserRequest) *model.User {

	now := time.Now() // サーバー側で生成した現在時刻を入れる

	return &model.User{
		ID:        req.ID,
		Name:      req.Name,
		Age:       req.Age,
		Email:     req.Email,
		CreatedAt: now,
		UpdatedAt: now,
	}

}
