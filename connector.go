package mongox

import (
	"context"

	"github.com/vimcoders/go-driver"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Config struct {
	DriverName string
	Usr        string
	Pwd        string
	Addr       string
	DB         string
}

type Connector struct {
	*mongo.Client
}

func Connect(cfg *Config) (driver.Connector, error) {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(cfg.Addr))

	if err != nil {
		return nil, err
	}

	return &Connector{client}, nil
}

func (c *Connector) Conn(ctx context.Context) (driver.Conn, error) {
	return nil, nil
	//dc, err := c.db.Conn(ctx)

	//if err != nil {
	//	return nil, err
	//}

	//return &connect{
	//	conn: dc,
	//}, nil
}

func (c *Connector) Tx(ctx context.Context) (driver.Tx, error) {
	return nil, nil
	//tx, err := c.db.BeginTx(ctx, &sql.TxOptions{})

	//if err != nil {
	//	return nil, err
	//}

	//return &trans{tx}, nil
}

func (c *Connector) SetMaxOpenConns(n int) {
	//c.db.SetMaxOpenConns(n)
}

func (c *Connector) Close() (err error) {
	return nil
	//return c.db.Close()
}
