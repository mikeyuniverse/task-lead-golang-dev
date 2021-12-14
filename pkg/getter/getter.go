package getter

import (
	"encoding/csv"
	"fmt"
	"grpc-practice/internal/server/models"
	"io"
	"net/http"
	"strconv"
	"strings"
)

type Getter struct{}

func New() *Getter {
	return &Getter{}
}

func (s *Getter) GetItemsByURL(url string) ([]models.Item, error) {
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
