package mongosh

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"test/config"
)

func Connect(ctx context.Context) (*mongo.Database, error) {
	conf := config.Load()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(conf.Mongo.MDB_ADDRESS))
	if err != nil {
		return nil, err
	}

	db := client.Database(conf.Mongo.MDB_NAME)

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return nil, err
	}

	return db, nil
}
