package repo

import (
	"context"
	"grpc-practice/internal/server/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ProductsRepo struct {
	db *mongo.Collection
}

func newProductsRepo(db *mongo.Database) *ProductsRepo {
	return &ProductsRepo{db: db.Collection(collectionProducts)}
}

func (r *ProductsRepo) UpdateItems(items []models.Item) error {
	for _, item := range items {

		if item.Name == "" && item.Price == 0 {
			continue
		}

		name := item.Name
		price := item.Price

		itemDB, err := r.GetItemByName(name)
		if err != nil {
			continue
		}

		if itemDB == nil {
			r.CreateItem(item)
			continue
		}

		if item.Price != price {
			r.UpdatePriceByName(item)
			continue
		}
	}
	return nil
}

func (r *ProductsRepo) GetItemByName(name string) (*models.Item, error) {
	var result models.Item

	if err := r.db.FindOne(context.TODO(), bson.M{"name": name}).Decode(&result); err != nil {
		return nil, err
	}

	if result.Name == "" && result.Price == 0 {
		return nil, nil
	}

	return &result, nil
}

func (r *ProductsRepo) CreateItem(item models.Item) error {
	_, err := r.db.InsertOne(context.TODO(), bson.D{
		{"name", item.Name},
		{"price", item.Price},
	})
	return err
}

func (r *ProductsRepo) UpdatePriceByName(item models.Item) error {
	_, err := r.db.UpdateOne(
		context.TODO(),
		bson.M{"name": item.Name},
		bson.D{
			{"$set", bson.D{{"price", item.Price}}},
		},
	)
	return err
}

type ListParams struct {
	Start     int64
	Limit     int64
	SortType  string
	OrderType string
}

func (r *ProductsRepo) GetItemsWithSort(params ListParams) ([]models.Item, error) {
	var fieldName string
	var orderNum int

	if params.SortType == "NAME" {
		fieldName = "name"
	} else if params.SortType == "PRICE" {
		fieldName = "price"
	}

	if params.OrderType == "ASC" {
		orderNum = 1
	} else if params.OrderType == "DESC" {
		orderNum = -1
	}

	queryOptions := options.FindOptions{}
	queryOptions.SetSkip(params.Start)
	queryOptions.SetLimit(params.Limit)
	queryOptions.SetSort(bson.D{{fieldName, orderNum}})

	queryResult, err := r.db.Find(context.TODO(), bson.M{}, &queryOptions)
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
