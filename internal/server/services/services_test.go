package services

import (
	"context"
	"errors"
	"fmt"
	"grpc-practice/internal/server/models"
	repo "grpc-practice/internal/server/repo"
	mock_repo "grpc-practice/internal/server/repo/mocks"
	getter "grpc-practice/pkg/proto/getterURL"
	"grpc-practice/pkg/proto/transport"
	"testing"

	"github.com/golang/mock/gomock"
)

type testTableFetch struct {
	name   string
	dataIn struct {
		url string
	}
	dataOut struct {
		getItemsByUrl struct {
			items    []models.Item
			errorObj error
		}
		updateItems struct {
			errorObj error
		}
	}
}

func Test_Fetch(t *testing.T) {

}

type testTableList struct {
	name string
	In   repo.ListParams
	Out  struct {
		wantErr              bool
		wantErrBeforeDBQuery bool
		errorObject          error
		itemCount            int
		itemsFromRepo        []models.Item
		itemsResult          []*transport.Item
	}
}

func Test_List(t *testing.T) {
	testTable := []testTableList{
		{
			name: "OK",
			In: repo.ListParams{
				Start:     2,
				Limit:     7,
				SortType:  "NAME",
				OrderType: "DESC",
			},
			Out: struct {
				wantErr              bool
				wantErrBeforeDBQuery bool
				errorObject          error
				itemCount            int
				itemsFromRepo        []models.Item
				itemsResult          []*transport.Item
			}{
				wantErr:              false,
				wantErrBeforeDBQuery: false,
				errorObject:          nil,
				itemCount:            14,
				itemsFromRepo: []models.Item{
					{Name: "Item 1", Price: 100},
					{Name: "Item 2", Price: 200},
					{Name: "Item 3", Price: 300},
					{Name: "Item 4", Price: 400},
					{Name: "Item 5", Price: 500},
					{Name: "Item 6", Price: 600},
					{Name: "Item 7", Price: 600},
				},
				itemsResult: []*transport.Item{
					{Name: "Item 1", Price: 100},
					{Name: "Item 2", Price: 200},
					{Name: "Item 3", Price: 300},
					{Name: "Item 4", Price: 400},
					{Name: "Item 5", Price: 500},
					{Name: "Item 6", Price: 600},
					{Name: "Item 7", Price: 600},
				},
			},
		},
		{
			name: "Error in sort type",
			In: repo.ListParams{
				Start:     0,
				Limit:     5,
				SortType:  "PRAME",
				OrderType: "DESC",
			},
			Out: struct {
				wantErr              bool
				wantErrBeforeDBQuery bool
				errorObject          error
				itemCount            int
				itemsFromRepo        []models.Item
				itemsResult          []*transport.Item
			}{
				wantErr:              true,
				wantErrBeforeDBQuery: true,
				errorObject:          errors.New("unknown sorting column name"),
				itemCount:            0,
				itemsFromRepo:        []models.Item{},
				itemsResult:          []*transport.Item{},
			},
		},
	}

	for _, test := range testTable {
		t.Run(test.name, func(t *testing.T) {
			// Init dependencies
			controller := gomock.NewController(t)
			repoProducts := mock_repo.NewMockProducter(controller)

			getter := getter.New()
			repository := repo.Repo{Products: repoProducts}
			services := New(&repository, getter)

			if !test.Out.wantErrBeforeDBQuery {
				repoProducts.EXPECT().GetItemsWithSort(test.In).Return(test.Out.itemsFromRepo, test.Out.errorObject)
			}
			// Action
			ctx := context.Background()
			items, err := services.List(ctx, test.In)

			// Assert
			if err != nil && !test.Out.wantErr {
				t.Fatalf("list error.\nerror - %s\n", err.Error())
			}

			if err != nil && test.Out.wantErr {
				if err.Error() != test.Out.errorObject.Error() {
					t.Fatalf("The expected error does not match the one received.\nExpected: %s\n", err.Error())
				}
			}

			if len(items) != test.Out.itemCount {
				fmt.Println(items)
				t.Fatalf("the expected quantity of goods does not match the received quantity\nNeed Items count - %d\nGot Items - %d\n", test.Out.itemCount, len(items))
			}
		})

	}
}

func Test_getItemsByURL(t *testing.T) {}
