package mongox

import (
	"context"
	"reflect"

	"github.com/vimcoders/go-driver"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Execer struct {
	*mongo.Database
	mongo.Session
}

func NewExecer(d *mongo.Database, s mongo.Session) driver.Execer {
	return &Execer{d, s}
}

func (e *Execer) Insert(ctx context.Context, doc interface{}) (interface{}, error) {
	docName := reflect.TypeOf(doc).Elem().Name()
	return e.Collection(docName).InsertOne(ctx, doc)
}

func (e *Execer) Delete(ctx context.Context, doc interface{}) (interface{}, error) {
	t := reflect.TypeOf(doc).Elem()
	if iface, ok := doc.(Identify); ok {
		return e.Collection(t.Name()).DeleteOne(ctx, bson.M{"_id": iface.Identify()})
	}
	return nil, nil
}

func (e *Execer) Query(ctx context.Context, doc interface{}) (result []interface{}, err error) {
	filter := bson.M{}
	t, v := reflect.TypeOf(doc).Elem(), reflect.ValueOf(doc).Elem()
	for i := 0; i < t.NumField(); i++ {
		if ok := v.Field(i).IsZero(); ok {
			continue
		}
		filter[t.Field(i).Tag.Get("bson")] = v.Field(i).Interface()
	}
	docName := t.Name()
	cur, err := e.Collection(docName).Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)
	for cur.Next(ctx) {
		v := reflect.New(reflect.TypeOf(doc).Elem()).Interface()
		if err := cur.Decode(v); err != nil {
			return nil, err
		}
		result = append(result, v)
	}
	return result, nil
}

func (e *Execer) Update(ctx context.Context, doc interface{}) (interface{}, error) {
	docName := reflect.TypeOf(doc).Elem().Name()
	if iface, ok := doc.(Identify); ok {
		return e.Collection(docName).UpdateByID(ctx, iface.Identify(), bson.M{"$set": doc})
	}
	return nil, nil
}

func (e *Execer) Close(ctx context.Context) error {
	e.EndSession(ctx)
	return nil
}
