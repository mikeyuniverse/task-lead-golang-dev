package services

import (
	"context"
	"errors"
	"grpc-practice/internal/server/models"
	"grpc-practice/internal/server/repo"
	"grpc-practice/pkg/proto/transport"
)

//go:generate mockgen -source=services.go -destination=mocks/mock.go

type Getter interface {
	GetItemsByURL(url string) ([]models.Item, error)
}

type Services struct {
	repo   *repo.Repo
	getter Getter
}

func New(repo *repo.Repo, getter Getter) *Services {
	return &Services{
		repo: repo,
	}
}

func (s *Services) Fetch(url string) error {
	items, err := s.getter.GetItemsByURL(url)
	if err != nil {
		return err
	}
	return s.repo.Products.UpdateItems(items)
}

func (s *Services) List(context context.Context, params repo.ListParams) ([]*transport.Item, error) {
	start := params.Start
	limit := params.Limit
	sortType := params.SortType
	orderType := params.OrderType

	if start < 0 || limit <= 0 {
		return []*transport.Item{}, errors.New("pagging params must be greater than 0")
	}

	if orderType != "ASC" && orderType != "DESC" {
		return []*transport.Item{}, errors.New("unknown ordering params")
	}

	if sortType != "NAME" && sortType != "PRICE" {
		return []*transport.Item{}, errors.New("unknown sorting column name")
	}

	items, err := s.repo.Products.GetItemsWithSort(params)
	if err != nil {
		return []*transport.Item{}, err
	}

	result := make([]*transport.Item, len(items))
	for _, item := range items {
		result = append(result, item.ToPB())
	}
	return result, nil
}
