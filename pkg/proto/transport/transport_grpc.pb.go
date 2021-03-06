// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package transport

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// FetchServiceClient is the client API for FetchService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FetchServiceClient interface {
	Fetch(ctx context.Context, in *FetchRequest, opts ...grpc.CallOption) (*Empty, error)
	List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error)
}

type fetchServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewFetchServiceClient(cc grpc.ClientConnInterface) FetchServiceClient {
	return &fetchServiceClient{cc}
}

func (c *fetchServiceClient) Fetch(ctx context.Context, in *FetchRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/FetchService/Fetch", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fetchServiceClient) List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error) {
	out := new(ListResponse)
	err := c.cc.Invoke(ctx, "/FetchService/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FetchServiceServer is the server API for FetchService service.
// All implementations must embed UnimplementedFetchServiceServer
// for forward compatibility
type FetchServiceServer interface {
	Fetch(context.Context, *FetchRequest) (*Empty, error)
	List(context.Context, *ListRequest) (*ListResponse, error)
	mustEmbedUnimplementedFetchServiceServer()
}

// UnimplementedFetchServiceServer must be embedded to have forward compatible implementations.
type UnimplementedFetchServiceServer struct {
}

func (UnimplementedFetchServiceServer) Fetch(context.Context, *FetchRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Fetch not implemented")
}
func (UnimplementedFetchServiceServer) List(context.Context, *ListRequest) (*ListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedFetchServiceServer) mustEmbedUnimplementedFetchServiceServer() {}

// UnsafeFetchServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FetchServiceServer will
// result in compilation errors.
type UnsafeFetchServiceServer interface {
	mustEmbedUnimplementedFetchServiceServer()
}

func RegisterFetchServiceServer(s grpc.ServiceRegistrar, srv FetchServiceServer) {
	s.RegisterService(&FetchService_ServiceDesc, srv)
}

func _FetchService_Fetch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FetchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FetchServiceServer).Fetch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/FetchService/Fetch",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FetchServiceServer).Fetch(ctx, req.(*FetchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FetchService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FetchServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/FetchService/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FetchServiceServer).List(ctx, req.(*ListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// FetchService_ServiceDesc is the grpc.ServiceDesc for FetchService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FetchService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "FetchService",
	HandlerType: (*FetchServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Fetch",
			Handler:    _FetchService_Fetch_Handler,
		},
		{
			MethodName: "List",
			Handler:    _FetchService_List_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "transport.proto",
}
