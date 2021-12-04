package repo

import (
	"grpc-practice/internal/server/models"
	"grpc-practice/internal/server/repo/mongo"
	"log"
)

type Repo struct {
	db *mongo.MongoDB
}

func New(db *mongo.MongoDB) (*Repo, error) {
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
			log.Println(err)
			continue
		}

		if itemDB == nil {
			err = r.db.CreateItem(item)
			if err != nil {
				log.Println(err)
			}
			continue
		}

		if item.Price != price {
			log.Printf("CHANGE PRICE %d --> %d", itemDB.Price, item.Price)
			err = r.db.UpdatePriceByName(item)
			if err != nil {
				log.Println(err)
			}
			continue
		}
	}
	return nil
}

func (r *Repo) GetItemsWithSort(start int32, limit int32, sortType string, orderType string) ([]models.Item, error) {

	return r.db.GetItemsWithSort(start, limit, sortType, orderType)
}
