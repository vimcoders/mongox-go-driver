package mongox

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type Tx struct {
	*mongo.Client
}

func (tx *Tx) Exec(list ...interface{}) (interface{}, error) {
	return tx.ExecContext(context.Background(), list)
}

func (tx *Tx) ExecContext(ctx context.Context, list ...interface{}) (interface{}, error) {
	session, err := tx.StartSession()

	if err != nil {
		return nil, err
	}

	defer session.EndSession(ctx)

	if err := session.StartTransaction(); err != nil {
		return nil, err
	}

	if err := mongo.WithSession(ctx, session, func(sc mongo.SessionContext) error {
		//if err := session.StartTransaction(); err != nil {
		//	return err
		//}

		//coll := tx.Database("db").Collection("coll")
		//res, err := coll.InsertOne(sc, bson.D{{"x", 1}})
		//if err != nil {
		//	return err
		//}

		//var result bson.M
		//if err = coll.FindOne(sc, bson.D{{"_id", res.InsertedID}}).Decode(result); err != nil {
		//	return err
		//}

		//return session.CommitTransaction(context.Background())
		return nil
	}); err != nil {
		return nil, err
	}

	return nil, nil
}
