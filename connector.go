package mongox

import (
	"context"

	"github.com/vimcoders/go-driver"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Config struct {
	Addr       string
	DB         string
}

type Connector struct {
	*mongo.Database
}

func Connect(cfg *Config) (driver.Connector, error) {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(cfg.Addr))
	if err != nil {
		return nil, err
	}
	return &Connector{client.Database("test")}, nil
}

func (c *Connector) Tx(ctx context.Context) (driver.Tx, error) {
	sess, err := c.Database.Client().StartSession()
	if err != nil {
		return nil, nil
	}
	sess.StartTransaction()
	return &Tx{c.Database, sess}, nil
}

func (c *Connector) SetMaxOpenConns(n int) {
	//c.db.SetMaxOpenConns(n)
}

func (c *Connector) Close() (err error) {
	return nil
	//return c.db.Close()
}
