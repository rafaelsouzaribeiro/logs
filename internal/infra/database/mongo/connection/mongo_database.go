package connection

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetMongoDataBase(DbName string, DB_HOST string) (*mongo.Database, error) {
	ctx := context.Background()
	mongoClient, err := mongo.Connect(ctx, options.Client().ApplyURI(DB_HOST))

	if err != nil {
		return nil, err
	}

	return mongoClient.Database(DbName), nil
}
