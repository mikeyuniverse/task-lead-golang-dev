package mongo

import (
	"context"
	"fmt"
	"grpc-server/internal/config"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDB struct {
	client         *mongo.Client
	collection     *mongo.Collection
	dbName         string
	collectionName string
}

func Connect(cfg *config.MongoConfig) (*MongoDB, error) {
	connectUri := fmt.Sprintf("mongodb://%s:%s", cfg.Host, cfg.Port)
	clientOptions := options.Client().
		ApplyURI(connectUri).
		SetAuth(options.Credential{
			Username: cfg.User,
			Password: cfg.Password,
		})

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return nil, err
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return nil, err
	}

	collection := client.Database(cfg.DBName).Collection(cfg.CollectionName)

	return &MongoDB{
		client:         client,
		collection:     collection,
		dbName:         cfg.DBName,
		collectionName: cfg.CollectionName,
	}, nil
}
