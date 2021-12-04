package mongo

import (
	"context"
	"fmt"
	"grpc-practice/internal/server/config"
	"grpc-practice/internal/server/models"

	"go.mongodb.org/mongo-driver/bson"
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
	clientOptions := options.Client().ApplyURI(connectUri)

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

func (m *MongoDB) GetItemByName(name string) (*models.Item, error) {

	var result models.Item
	if err := m.collection.FindOne(context.TODO(), bson.M{"name": name}).Decode(&result); err != nil {
		return nil, err
	}
	if result.Name == "" && result.Price == 0 {
		return nil, nil
	}
	return &result, nil
}

func (m *MongoDB) CreateItem(item models.Item) error {
	_, err := m.collection.InsertOne(context.TODO(), bson.D{
		{"name", item.Name},
		{"price", item.Price},
	})
	if err != nil {
		return err
	}
	return nil
}

func (m *MongoDB) UpdatePriceByName(item models.Item) error {
	_, err := m.collection.UpdateOne(
		context.TODO(),
		bson.M{"name": item.Name},
		bson.D{
			{"$set", bson.D{{"price", item.Price}}},
		},
	)
	if err != nil {
		return err
	}
	return nil
}

func (m *MongoDB) GetItemsWithSort(start int32, limit int32, sortType string, orderType string) ([]models.Item, error) {
	var fieldName string
	var orderNum int

	if sortType == "NAME" {
		fieldName = "name"
	} else if sortType == "PRICE" {
		fieldName = "price"
	}

	if orderType == "ASC" {
		orderNum = 1
	} else if orderType == "DESC" {
		orderNum = -1
	}

	queryOptions := options.FindOptions{}
	queryOptions.SetSkip(int64(start))
	queryOptions.SetLimit(int64(limit))
	queryOptions.SetSort(bson.D{{fieldName, orderNum}})

	queryResult, err := m.collection.Find(context.TODO(), bson.M{}, &queryOptions)
	if err != nil {
		return []models.Item{}, err
	}

	results := make([]models.Item, 1)
	err = queryResult.All(context.Background(), &results)
	if err != nil {
		return []models.Item{}, err
	}

	return results, nil
}
