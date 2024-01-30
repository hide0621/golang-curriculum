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
