package repo

import (
	"grpc-practice/internal/server/models"

	"go.mongodb.org/mongo-driver/mongo"
)

var collectionProducts = "products"

type Products interface {
	UpdateItems(items []models.Item) error
	GetItemByName(name string) (*models.Item, error)
	CreateItem(item models.Item) error
	UpdatePriceByName(item models.Item) error
	GetItemsWithSort(params ListParams) ([]models.Item, error)
}

type Repo struct {
	Products Products
}

func New(db *mongo.Database) (*Repo, error) {
	return &Repo{
		Products: newProductsRepo(db),
	}, nil
}
