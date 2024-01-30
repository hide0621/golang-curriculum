package transaction

import (
	"context"
	"database/sql"
)

type Transaction interface {
	DoInTx(context.Context, func(context.Context) (any, error)) (any, error)
}

type tx struct {
	db *sql.DB
}

func NewTransaction(db *sql.DB) Transaction {
	return &tx{db: db}
}

// 引数fにはDB操作を行う関数を渡す
func (t *tx) DoInTx(ctx context.Context, f func(context.Context) (any, error)) (any, error) {

	// トランザクションを開始する（無名のトランザクションのオブジェクトの作成）
	tx, err := t.db.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return nil, err
	}

	/*
		上記で作成した無名のトランザクションに対して「tx」というキーを与えてコンテキストを作成する
		ctxの中にはtxトランザクションが入っている
	*/
	ctx = context.WithValue(ctx, "tx", tx)

	/*
		上記で作成したtxトランザクションでDB操作を行う
		txトランザクション中にエラーが発生した場合はロールバックする
	*/
	v, err := f(ctx)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	/*
		上記でエラーが発生しなかった場合はコミットする
		コミット出来なかった場合はロールバックする
	*/
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	return v, nil
}
