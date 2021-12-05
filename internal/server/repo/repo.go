package repo

import (
	"grpc-practice/internal/server/models"
)

type Database interface {
	GetItemByName(name string) (*models.Item, error)
	CreateItem(item models.Item) error
	UpdatePriceByName(item models.Item) error
	GetItemsWithSort(start int32, limit int32, sortType string, orderType string) ([]models.Item, error)
}

type Repo struct {
	db Database
}

func New(db Database) (*Repo, error) {
	return &Repo{db: db}, nil
}

func (r *Repo) UpdateItems(items []models.Item) error {
	for _, item := range items {

		if item.Name == "" && item.Price == 0 {
			continue
		}

		name := item.Name
		price := item.Price

		itemDB, err := r.db.GetItemByName(name)
		if err != nil {
			continue
		}

		if itemDB == nil {
			r.db.CreateItem(item)
			continue
		}

		if item.Price != price {
			r.db.UpdatePriceByName(item)
			continue
		}
	}
	return nil
}

func (r *Repo) GetItemsWithSort(start int32, limit int32, sortType string, orderType string) ([]models.Item, error) {
	return r.db.GetItemsWithSort(start, limit, sortType, orderType)
}
