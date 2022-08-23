package mongox

import (
	"context"

	"github.com/vimcoders/go-driver"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type Config struct {
	Addr        string
	DB          string
	MaxPoolSize uint64
}

type Connector struct {
	*mongo.Database
}

func Connect(cfg *Config) (driver.Connector, error) {
	option := options.Client().ApplyURI(cfg.Addr)
	if cfg.MaxPoolSize > 0 {
		option.SetMaxPoolSize(cfg.MaxPoolSize)
	}
	client, err := mongo.Connect(context.Background(), option)
	if err != nil {
		return nil, err
	}
	if err := client.Ping(context.Background(), readpref.Primary()); err != nil {
		return nil, err
	}
	return &Connector{client.Database(cfg.DB)}, nil
}

func (c *Connector) Tx(ctx context.Context) (driver.Tx, error) {
	sess, err := c.Database.Client().StartSession()
	if err != nil {
		return nil, nil
	}
	sess.StartTransaction()
	return NewTx(c.Database, sess), nil
}

func (c *Connector) Execer(ctx context.Context) (driver.Execer, error) {
	sess, err := c.Database.Client().StartSession()
	if err != nil {
		return nil, nil
	}
	return &Execer{c.Database, sess}, nil
}

func (c *Connector) SetMaxOpenConns(n int) {
	//c.db.SetMaxOpenConns(n)
}

func (c *Connector) Close() (err error) {
	return c.Database.Client().Disconnect(context.Background())
}
