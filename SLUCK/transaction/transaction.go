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

func (t *tx) DoInTx(ctx context.Context, f func(context.Context) (any, error)) (any, error) {

	return nil, nil
}
