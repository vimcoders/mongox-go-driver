package mongox

import (
	"context"
	"reflect"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Tx struct {
	*mongo.Database
	mongo.Session
}

func (tx *Tx) Delete(ctx context.Context, value interface{}) (interface{}, error) {
	document, ok := value.(Document)
	if !ok {
		return nil, nil
	}
	if iface, ok := value.(Identify); ok {
		return tx.Collection(document.DocumentName()).DeleteOne(ctx, bson.D{{"_id", iface.Id()}})
	}
	t, v := reflect.TypeOf(value), reflect.ValueOf(value)
	for i := 0; i < t.NumField(); i++ {
		if t.Field(i).Name != "id" {
			continue
		}
		return tx.Collection(document.DocumentName()).DeleteOne(ctx, bson.D{{"id", v.Field(i).String()}})
	}
	return nil, nil
}

func (tx *Tx) Update(ctx context.Context, value interface{}) (interface{}, error) {
	document, ok := value.(Document)
	if !ok {
		return nil, nil
	}
	if iface, ok := value.(Identify); ok {
		return tx.Collection(document.DocumentName()).UpdateOne(ctx, bson.D{{"id", iface.Id()}}, value)
	}
	t, v := reflect.TypeOf(value), reflect.ValueOf(value)
	for i := 0; i < t.NumField(); i++ {
		if t.Field(i).Name != "id" {
			continue
		}
		return tx.Collection(document.DocumentName()).UpdateOne(ctx, bson.D{{"id", v.Field(i).String()}}, value)
	}
	return nil, nil
}

func (tx *Tx) Query(ctx context.Context, value interface{}) (result []interface{}, err error) {
	document, ok := value.(Document)
	if !ok {
		return nil, nil
	}
	cur, err := tx.Collection(document.DocumentName()).Find(ctx, bson.M{}, nil)
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)
	for cur.Next(ctx) {
		v := reflect.New(reflect.TypeOf(value)).Interface()
		if err := cur.Decode(v); err != nil {
			return nil, err
		}
		result = append(result, v)
	}
	return result, nil
}

func (tx *Tx) Insert(ctx context.Context, value interface{}) (interface{}, error) {
	if document, ok := value.(Document); ok {
		return tx.Collection(document.DocumentName()).InsertOne(ctx, value)
	}
	return nil, nil
}

func (tx *Tx) Commit(ctx context.Context) error {
	defer tx.EndSession(ctx)
	return tx.Session.CommitTransaction(ctx)
}

func (tx *Tx) Rollback(ctx context.Context) error {
	defer tx.EndSession(ctx)
	return tx.Session.AbortTransaction(ctx)
}
