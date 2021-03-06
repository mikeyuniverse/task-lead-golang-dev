package server

import (
	"context"
	"fmt"
	"grpc-practice/pkg/proto/transport"

	"google.golang.org/grpc"
)

type Server struct {
	grpcClient transport.FetchServiceClient
}

func New(grpcPort string) (*Server, error) {
	conn, err := grpc.Dial(fmt.Sprintf(":%s", grpcPort), grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	client := transport.NewFetchServiceClient(conn)

	return &Server{
		grpcClient: client,
	}, nil
}

func (s *Server) Fetch(url string) error {
	req := transport.FetchRequest{Url: url}
	_, err := s.grpcClient.Fetch(context.Background(), &req, []grpc.CallOption{}...)
	return err
}

func (s *Server) List(pagg *transport.Pagging, sort *transport.Sorting) ([][]string, error) {
	req := transport.ListRequest{
		Start:   pagg.Start,
		Limit:   pagg.Limit,
		Sort:    transport.SortType_NAME,
		Pagging: transport.SortOrder_DESC,
	}

	if sort.SortingType.GetType() == "NAME" {
		req.Sort = transport.SortType_NAME
	}

	if sort.SortingType.GetType() == "PRICE" {
		req.Sort = transport.SortType_PRICE
	}

	if sort.OrderType.Ascending {
		req.Pagging = transport.SortOrder_ASC
	}

	if sort.OrderType.Descending {
		req.Pagging = transport.SortOrder_DESC
	}

	items, err := s.grpcClient.List(context.TODO(), &req)
	if err != nil {
		return [][]string{}, err
	}

	data := itemsFormating(items)
	return data, err
}

func itemsFormating(items *transport.ListResponse) [][]string {
	itemList := items.Item
	result := make([][]string, len(itemList))
	for _, item := range itemList {
		result = append(result, []string{item.Name, string(item.Price)})
	}
	return result
}
