package server

import (
	"context"
	"fmt"
	"grpc-practice/internal/server/config"
	"grpc-practice/internal/server/models"
	"net/http"

	// "grpc-server/pkg/proto/transport"

	"net"

	"grpc-practice/pkg/proto/transport"

	"google.golang.org/grpc"
)

type Server struct {
	Host       string
	Port       string
	services   services
	httpServer *http.Server
	transport.UnimplementedFetchServiceServer
}

type services interface {
	Fetch(req *transport.FetchRequest) error
	List(context context.Context, start int32, limit int32, sortType string, orderType string) ([]*transport.Item, error)
	GetDataByURL(url string) ([]models.Item, error)
	UpdateDB([]models.Item) error
}

func New(cfg *config.GRPC, services services) (*Server, error) {
	return &Server{
		Host:     cfg.Host,
		Port:     cfg.Port,
		services: services,
		httpServer: &http.Server{
			Addr: "localhost:8080",
		},
	}, nil
}

func (srv *Server) Start() error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", srv.Port))
	if err != nil {
		return err
	}

	grpcServer := grpc.NewServer()
	transport.RegisterFetchServiceServer(grpcServer, srv)
	if err := grpcServer.Serve(lis); err != nil {
		return err
	}
	return nil
}

func (s *Server) Fetch(ctx context.Context, req *transport.FetchRequest) (*transport.Empty, error) {
	err := s.services.Fetch(req)
	if err != nil {
		return new(transport.Empty), err
	}
	return new(transport.Empty), nil
}

func (s *Server) List(ctx context.Context, req *transport.ListRequest) (*transport.ListResponse, error) {
	start := req.Start
	limit := req.Limit
	sortType := req.Sort.String()
	orderType := req.Pagging.String()
	response, err := s.services.List(ctx, start, limit, sortType, orderType)
	if err != nil {
		return nil, err
	}
	return &transport.ListResponse{Item: response}, nil
}
