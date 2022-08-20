package mongox

import (
	"context"

	"github.com/vimcoders/go-driver"
	"go.mongodb.org/mongo-driver/mongo"
)

type Tx struct {
	driver.Execer
	commit   func(ctx context.Context) error
	rollback func(ctx context.Context) error
}

func NewTx(d *mongo.Database, s mongo.Session) driver.Tx {
	return &Tx{
		Execer: NewExecer(d, s),
		commit: func(ctx context.Context) error {
			return s.CommitTransaction(ctx)
		},
		rollback: func(ctx context.Context) error {
			return s.AbortTransaction(ctx)
		},
	}
}

func (tx *Tx) Commit(ctx context.Context) error {
	defer tx.Close(ctx)
	return tx.commit(ctx)
}

func (tx *Tx) Rollback(ctx context.Context) error {
	defer tx.Close(ctx)
	return tx.rollback(ctx)
}
