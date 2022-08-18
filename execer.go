package mongox

import (
	"context"
	"reflect"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Execer struct {
	*mongo.Database
	mongo.Session
}

func (e *Execer) Insert(ctx context.Context, doc interface{}) (interface{}, error) {
	document, ok := doc.(Document)
	if !ok {
		return nil, nil
	}
	return e.Collection(document.DocumentName()).InsertOne(ctx, doc)
}

func (e *Execer) Delete(ctx context.Context, doc interface{}) (interface{}, error) {
	document, ok := doc.(Document)
	if !ok {
		return nil, nil
	}
	if iface, ok := doc.(Identify); ok {
		return e.Collection(document.DocumentName()).DeleteOne(ctx, bson.M{"_id": iface.Identify()})
	}
	return nil, nil
}

func (e *Execer) Query(ctx context.Context, doc interface{}) ([]interface{}, error) {
	document, ok := doc.(Document)
	if !ok {
		return nil, nil
	}
	cur, err := e.Collection(document.DocumentName()).Find(ctx, bson.M{})
	if err != nil {
		return nil, err

	}
	defer cur.Close()
	for cur.Next(ctx) {
		v := reflect.New(reflect.TypeOf(value)).Interface()
		if err := cur.Decode(v); err != nil {
			return nil, err
		}
		result = append(result, v)
	}
	return result, nil
}

func (e *Execer) Update(ctx context.Context, doc interface{}) (interface{}, error) {
	document, ok := doc.(Document)
	if !ok {
		return nil, nil
	}
	if iface, ok := doc.(Identify); ok {
		return e.Collection(document.DocumentName()).UpdateByID(ctx, iface.Identify(), bson.M{"$set": doc})
	}
	return nil, nil
}
