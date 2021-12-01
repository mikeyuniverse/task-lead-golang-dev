package server

import "grpc-client/internal/models"

type Server struct{}

func New() *Server {
	return &Server{}
}

func (s *Server) Fetch(url string) {}

func (s *Server) List(pagg *models.Pagging, sort *models.Sorting) ([]string, error) {
	return []string{}, nil
}
