package server

import (
	"context"
	"fmt"
	"grpc-practice/internal/server/config"
	"grpc-practice/internal/server/repo"
	"grpc-practice/pkg/proto/transport"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

//go:generate mockgen -source=server.go -destination=mocks/mock.go

type servicer interface {
	Fetch(url string) error
	List(context context.Context, params repo.ListParams) ([]*transport.Item, error)
}

type Server struct {
	Port     string
	services servicer
	transport.UnimplementedFetchServiceServer
	context *context.Context
}

func New(context *context.Context, cfg config.GRPC, services servicer) (*Server, error) {
	return &Server{
		Port:     cfg.Port,
		services: services,
		context:  context,
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
	err := s.services.Fetch(req.Url)
	if err != nil {
		return new(transport.Empty), status.Error(codes.Canceled, err.Error())
	}
	return new(transport.Empty), nil
}

func (s *Server) List(ctx context.Context, req *transport.ListRequest) (*transport.ListResponse, error) {
	response, err := s.services.List(ctx, repo.ListParams{
		Start:     int64(req.Start),
		Limit:     int64(req.Limit),
		SortType:  req.Sort.String(),
		OrderType: req.Pagging.String(),
	})
	if err != nil {
		return nil, status.Error(codes.Canceled, err.Error())
	}

	return &transport.ListResponse{Item: response}, nil
}
