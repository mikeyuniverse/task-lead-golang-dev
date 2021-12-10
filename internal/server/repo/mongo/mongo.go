package mongo

import (
	"context"
	"fmt"
	"grpc-practice/internal/server/config"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDB struct {
	db             *mongo.Database
	dbName         string
	collectionName string
}

func Connect(cfg *config.MongoConfig) (*mongo.Database, error) {
	connectUri := fmt.Sprintf("mongodb://%s:%s", cfg.Host, cfg.Port)
	clientOptions := options.Client().ApplyURI(connectUri)

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return nil, err
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return nil, err
	}

	return client.Database(cfg.DBName), nil
}
