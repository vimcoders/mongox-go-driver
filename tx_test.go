package mongox

import (
	"context"
	"reflect"
	"testing"
)

func TestTxReflect(t *testing.T) {
	type TestInsert struct {
		Id       string `bson:"_id"`
		Passport string `bson:"_passport"`
		Ip       string `bson:"_ip"`
		Addr     string `bson:"_addr"`
	}
	t.Log(reflect.TypeOf(&TestInsert{}).Elem().Name())
}

func TestTxInsert(t *testing.T) {
	type TestInsert struct {
		Id       string `bson:"_id"`
		Passport string `bson:"_passport"`
		Ip       string `bson:"_ip"`
		Addr     string `bson:"_addr"`
	}
	c, err := Connect(&Config{
		Addr: "mongodb://127.0.0.1:27017",
		DB:   "test",
	})
	if err != nil {
		t.Error(err)
		return
	}
	tx, err := c.Tx(context.Background())
	if err != nil {
		t.Error(err)
		return
	}
	for i := 0; i < 10000000; i++ {
		t.Log(tx.Insert(context.Background(), &TestInsert{Ip: "127.0.0.1", Addr: ":8001"}))
	}
}
