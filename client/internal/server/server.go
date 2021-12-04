package server

import (
	"context"
	"grpc-client/internal/models"
	"grpc-client/pkg/proto/transport"

	"google.golang.org/grpc"
)

type Server struct {
	grpcClient transport.FetchServiceClient
}

func New() (*Server, error) {
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	client := transport.NewFetchServiceClient(conn)

	return &Server{
		grpcClient: client,
	}, nil
}

func (s *Server) Fetch(url string) error {
	_, err := s.grpcClient.Fetch(context.Background(), &transport.FetchRequest{Url: url})
	return err
}

func (s *Server) List(pagg *models.Pagging, sort *models.Sorting) ([]string, error) {
	return []string{}, nil
}
