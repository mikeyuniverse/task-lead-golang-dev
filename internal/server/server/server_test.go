package server

import (
	"context"
	"errors"
	"grpc-practice/internal/server/config"
	"grpc-practice/internal/server/repo"
	mock_server "grpc-practice/internal/server/server/mocks"
	"grpc-practice/pkg/proto/transport"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestServer_Fetch(t *testing.T) {
	testTable := []struct {
		name       string
		requestUrl string
		wantError  error
		wantErr    bool
	}{
		{
			name:       "OK",
			requestUrl: "http://164.92.251.245:8080/api/v1/products",
			wantError:  nil,
			wantErr:    false,
		},
		{
			name:       "EMPTY URL",
			requestUrl: "",
			wantError:  errors.New("empty url"),
			wantErr:    true,
		},
	}

	for _, testCase := range testTable {
		// Initialize dependencies
		ctx := context.Background()
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		services := mock_server.NewMockservicer(ctrl)
		services.EXPECT().Fetch(testCase.requestUrl).Return(testCase.wantError)
		server, err := New(&ctx, config.GRPC{Port: "9000"}, services)

		if err != nil {
			t.Fatal("error with server initialize")
		}

		t.Run(testCase.name, func(t *testing.T) {
			// TODO Check answer from request
			_, err = server.Fetch(ctx, &transport.FetchRequest{
				Url: testCase.requestUrl,
			})

			if err != nil && !testCase.wantErr {
				t.Fatalf("Fetch error. Error - %s\n", err.Error())
			}

		})
	}

}

func TestServer_List(t *testing.T) {
	// TODO How do I test pagination and sorting?
	testTable := []struct {
		name    string
		data_in struct {
			sort    transport.SortType
			pagging transport.Pagging
		}
		data_want struct {
			itemCount int
			wantErr   bool
			items     transport.ListResponse
			errorText error
		}
	}{
		{
			name: "OK",
			data_in: struct {
				sort    transport.SortType
				pagging transport.Pagging
			}{
				sort: transport.SortType_NAME,
				pagging: transport.Pagging{
					Limit: 5,
					Start: 0,
				},
			},
			data_want: struct {
				itemCount int
				wantErr   bool
				items     transport.ListResponse
				errorText error
			}{
				itemCount: 5,
				wantErr:   false,
				items: transport.ListResponse{
					Item: []*transport.Item{
						{Name: "Item 1", Price: 100},
						{Name: "Item 2", Price: 200},
						{Name: "Item 3", Price: 300},
						{Name: "Item 4", Price: 400},
						{Name: "Item 5", Price: 500},
					},
				},
				errorText: nil,
			},
		},
	}

	for _, test := range testTable {
		// Initialize dependencies
		ctx := context.Background()
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		services := mock_server.NewMockservicer(ctrl)
		services.EXPECT().List(context.TODO(), repo.ListParams{
			Start:     int64(test.data_in.pagging.Start),
			Limit:     int64(test.data_in.pagging.Limit),
			SortType:  test.data_in.sort.String(),
			OrderType: transport.SortOrder_DESC.String(),
		}).Return(test.data_want.items.Item, test.data_want.errorText)
		server, err := New(&ctx, config.GRPC{Port: "9000"}, services)

		if err != nil {
			t.Error("error init server")
		}

		response, err := server.List(ctx, &transport.ListRequest{
			Start:   test.data_in.pagging.Start,
			Limit:   test.data_in.pagging.Limit,
			Sort:    test.data_in.sort,
			Pagging: transport.SortOrder_DESC,
		})

		if err != nil && !test.data_want.wantErr {
			t.Fatalf("request error - %s\n", err)
		}

		if len(response.Item) != test.data_want.itemCount {
			t.Fatalf("the expected quantity of goods does not match what was received\nWant items - %d\nGot - %d\n", test.data_in.pagging.Limit, len(response.Item))
		}

	}

}
