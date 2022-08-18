package mongox

import (
	"context"
	"testing"
)

type Test struct {
	Id       string `bson:"_id"`
	Passport string `bson:"_passport"`
	Ip       string `bson:"_ip"`
	Addr     string `bson:"_addr"`
}

func (t *Test) DocumentName() string {
	return "login"
}

func (t *Test) Identify() string {
	return t.Id
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
	defer tx.Rollback(context.Background())
	t.Log(tx.Insert(context.Background(), &Test{Id: "006", Ip: "127.0.0.1", Addr: ":8001"}))
}

func TestTxDelete(t *testing.T) {
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
	defer tx.Rollback(context.Background())
	t.Log(tx.Delete(context.Background(), &Test{Id: "005", Ip: "127.0.0.1", Addr: ":8001"}))
}

func TestTxQuery(t *testing.T) {
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
	defer tx.Rollback(context.Background())
	t.Log(tx.Query(context.Background(), &Test{Id: "005", Ip: "127.0.0.1", Addr: ":8001"}))
}

func TestTxUpdate(t *testing.T) {
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
	defer tx.Rollback(context.Background())
	t.Log(tx.Update(context.Background(), &Test{Id: "005", Ip: "127.0.0.1", Addr: ":8001"}))
}
