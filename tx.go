package mongox

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type Tx struct {
	*mongo.Client
}

func (tx *Tx) Exec(list ...interface{}) (interface{}, error) {
	session, err := tx.StartSession()

	if err != nil {
		return nil, err
	}

	defer session.EndSession(context.Background())

	if err := session.StartTransaction(); err != nil {
		return nil, err
	}

	if err := mongo.WithSession(context.Background(), session, func(sc mongo.SessionContext) error {
		return nil
	}); err != nil {
		return nil, err
	}

	return nil, nil
}
