package mongox

import (
	"context"

	"github.com/mongo-driver/bson"
	"github.com/mongo-driver/mongo"
	"github.com/mongo-go-driver/mongo"
)

type Tx struct {
	*mongo.Database
	mongo.Session
}

func (tx *Tx) Delete(ctx context.Context, i interface{}) (interface{}, error) {
	col := tx.Collection("test")
	return col.DeleteOne(ctx, bson.D{"id", i})
}

func (tx *Tx) Update(ctx context.Context) (interface{}, error) {
	col := tx.Collection("test")
	return col.UpdateOne(ctx, bson.D{"id", i})
}

func (tx *Tx) Query(ctx context.Context, i interface{}) ([]interface{}, error) {
	col := tx.Collection("test")
	col.Find(ctx, nil)
	return nil, nil
}

func (tx *Tx) Insert(ctx context.Context, i interface{}) (interface{}, error) {
	col := tx.Collection("test")
	return col.InsertOne(ctx, i)
}

func (tx *Tx) Commit(ctx context.Context) error {
	defer tx.EndSession(ctx)
	return tx.Session.CommitTransaction(ctx)
}

func (tx *Tx) Rollback(ctx context.Context) error {
	defer tx.EndSession(ctx)
	return tx.Session.AbortTransaction(ctx)
}
