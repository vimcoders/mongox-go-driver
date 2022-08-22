package mongox

import (
	"context"
	"fmt"
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

func TestTxQuery(t *testing.T) {
	type TestAccount struct {
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
	tx.Insert(context.Background(), &TestAccount{Passport: "123", Addr: ":8001"})
	tx.Insert(context.Background(), &TestAccount{Passport: "456", Addr: ":8002"})
	v, err := tx.Query(context.Background(), &TestAccount{Passport: "123"})
	if err != nil {
		t.Error(err)
		return
	}
	for _, vv := range v {
		t.Log(fmt.Printf("%+v", vv.(*TestAccount)))
	}
}
