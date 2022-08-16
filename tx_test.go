package mongox

import (
	"context"
	"testing"
)

type Login struct {
	Id       string `bson:"_id"`
	Passport string
	Ip       string
	Addr     string
}

func (l *Login) DocumentName() string {
	return "login"
}

func TestTxInsert(t *testing.T) {
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
	t.Log(tx.Insert(context.Background(), &Login{Id: "001", Ip: "127.0.0.1", Addr: ":8001"}))
}
