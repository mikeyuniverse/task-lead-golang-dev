package server

import (
	"context"
	"errors"
	"grpc-practice/internal/server/config"
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

func TestServer_List(t *testing.T) {}
