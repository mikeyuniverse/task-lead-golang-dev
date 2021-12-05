package server

import (
	"context"
	"fmt"
	"grpc-practice/internal/server/config"
	"grpc-practice/pkg/proto/transport"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type services interface {
	Fetch(req *transport.FetchRequest) error
	List(context context.Context, start int32, limit int32, sortType string, orderType string) ([]*transport.Item, error)
}

type Server struct {
	Port     string
	services services
	transport.UnimplementedFetchServiceServer
}

func New(cfg *config.GRPC, services services) (*Server, error) {
	return &Server{
		Port:     cfg.Port,
		services: services,
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
	s.services.Fetch(req)
	return new(transport.Empty), nil
}

func (s *Server) List(ctx context.Context, req *transport.ListRequest) (*transport.ListResponse, error) {
	start := req.Start
	limit := req.Limit
	sortType := req.Sort.String()
	orderType := req.Pagging.String()

	response, err := s.services.List(ctx, start, limit, sortType, orderType)
	if err != nil {
		return nil, status.Error(codes.Canceled, err.Error())
	}

	return &transport.ListResponse{Item: response}, nil
}
