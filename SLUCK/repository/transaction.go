package repository

import (
	"context"
	"database/sql"
)

var txKey = struct{}{}

// コンテキスト内にあるトランザクションを識別するキーを取得して、対応するトランザクションのオブジェクトを返す
func GetTx(ctx context.Context) (*sql.Tx, bool) {

	tx, ok := ctx.Value(&txKey).(*sql.Tx)

	return tx, ok

}
