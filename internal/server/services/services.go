package services

import (
	"context"
	"encoding/csv"
	"errors"
	"fmt"
	"grpc-practice/internal/server/models"
	"grpc-practice/pkg/proto/transport"
	"io"
	"net/http"
	"strconv"
	"strings"
)

type Repository interface {
	GetItemsWithSort(int32, int32, string, string) ([]models.Item, error)
	UpdateItems([]models.Item) error
}

type Services struct {
	repo Repository
}

func New(repo Repository) *Services {
	return &Services{
		repo: repo,
	}
}

func (s *Services) Fetch(req *transport.FetchRequest) error {
	items, err := s.GetItemsByURL(req.Url)
	if err != nil {
		return err
	}
	return s.updateAllItems(items)
}

func (s *Services) List(context context.Context, start int32, limit int32, sortType string, orderType string) ([]*transport.Item, error) {
	if start < 0 || limit <= 0 {
		fmt.Println(start, limit, sortType, orderType)
		return []*transport.Item{}, errors.New("pagging params must be greater than 0")
	}

	if orderType != "ASC" && orderType != "DESC" {
		fmt.Println(start, limit, sortType, orderType)
		return []*transport.Item{}, errors.New("unknown ordering params")
	}

	if sortType != "NAME" && sortType != "PRICE" {
		fmt.Println(start, limit, sortType, orderType)
		return []*transport.Item{}, errors.New("unknown sorting column name")
	}

	items, err := s.repo.GetItemsWithSort(start, limit, sortType, orderType)
	if err != nil {
		return []*transport.Item{}, err
	}

	result := make([]*transport.Item, len(items))
	for _, item := range items {
		result = append(result, item.ToPB())
	}
	return result, nil
}

func (s *Services) GetItemsByURL(url string) ([]models.Item, error) {
	resp, err := http.Get(url)
	if err != nil {
		return []models.Item{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return []models.Item{}, fmt.Errorf("request error. code - %d", resp.StatusCode)
	}

	r := csv.NewReader(resp.Body)

	data := make([]models.Item, 1)

	for {
		record, err := r.Read()

		if err != nil {
			if err == io.EOF {
				break
			}
			return []models.Item{}, err
		}

		record = strings.Split(record[0], ";")
		if len(record) != 2 {
			continue
		}

		name := record[0]
		price, err := strconv.Atoi(record[1])
		if err != nil {
			return []models.Item{}, err
		}

		data = append(data, models.Item{
			Name:  name,
			Price: price,
		})
	}

	return data, nil
}

func (s *Services) updateAllItems(items []models.Item) error {
	return s.repo.UpdateItems(items)
}
