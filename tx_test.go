package mongox

import (
	"context"
	"testing"
)

func TestTxDelete(t *testing.T) {
	c, err := Connect(nil)
	if err != nil {
		t.Error(err)
		return
	}
	tx, err := c.Tx(context.Background())
	if err != nil {
		t.Error(err)
		return
	}
	type Login struct {
		Passport string
		Ip       string
		Addr     string
	}
	t.Log(tx.Insert(&Login{Ip: "127.0.0.1", Addr: ":8001"}))
}
