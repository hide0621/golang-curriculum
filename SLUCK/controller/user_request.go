package controller

import (
	"sluck/model"
	"time"
)

// コントローラー層で扱うデータ型(リクエストでさばくでデータ型)
type UserRequest struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email"`
}

// コントローラー層からモデル層への変換はコントローラー層で行う（DTOのような役割を持つ関数）
func toModel(req UserRequest) *model.User {

	now := time.Now() // サーバー側で生成した現在時刻を入れる

	return &model.User{
		Name:      req.Name,
		Age:       req.Age,
		Email:     req.Email,
		CreatedAt: now,
		UpdatedAt: now,
	}

}
