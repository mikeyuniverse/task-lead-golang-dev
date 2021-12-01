package server

import (
	"fmt"
	"grpc-server/internal/config"
	"grpc-server/internal/services"
	"net"

	"google.golang.org/grpc"
)

type Server struct {
	Host     string
	Port     string
	services *services.Services
}

func New(cfg *config.GRPC, services *services.Services) (*Server, error) {
	return &Server{
		Host:     cfg.Host,
		Port:     cfg.Port,
		services: services,
	}, nil
}

func (s *Server) Start() error {
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%s", s.Host, s.Port))
	if err != nil {
		return err
	}

	grpcServer := grpc.NewServer()
	// notification.RegisterNotificationServiceServer(grpcServer, &server{})
	if err := grpcServer.Serve(lis); err != nil {
		return err
	}
	return nil
}
