package repository

import (
	"context"
	"database/sql"
)

var txKey = struct{}{}

// コンテキスト内にあるトランザクションを識別するキーを取得する
func GetTx(ctx context.Context) (*sql.Tx, bool) {

	tx, ok := ctx.Value(&txKey).(*sql.Tx)

	return tx, ok

}
