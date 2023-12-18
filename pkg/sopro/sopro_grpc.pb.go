// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: sopro/sopro.proto

package sopropb

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

const (
	SoproService_PingPong_FullMethodName = "/sopropb.SoproService/PingPong"
)

// SoproServiceClient is the client API for SoproService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SoproServiceClient interface {
	PingPong(ctx context.Context, in *HealthRequest, opts ...grpc.CallOption) (*HealthResponse, error)
}

type soproServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSoproServiceClient(cc grpc.ClientConnInterface) SoproServiceClient {
	return &soproServiceClient{cc}
}

func (c *soproServiceClient) PingPong(ctx context.Context, in *HealthRequest, opts ...grpc.CallOption) (*HealthResponse, error) {
	out := new(HealthResponse)
	err := c.cc.Invoke(ctx, SoproService_PingPong_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SoproServiceServer is the server API for SoproService service.
// All implementations must embed UnimplementedSoproServiceServer
// for forward compatibility
type SoproServiceServer interface {
	PingPong(context.Context, *HealthRequest) (*HealthResponse, error)
	mustEmbedUnimplementedSoproServiceServer()
}

// UnimplementedSoproServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSoproServiceServer struct {
}

func (UnimplementedSoproServiceServer) PingPong(context.Context, *HealthRequest) (*HealthResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PingPong not implemented")
}
func (UnimplementedSoproServiceServer) mustEmbedUnimplementedSoproServiceServer() {}

// UnsafeSoproServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SoproServiceServer will
// result in compilation errors.
type UnsafeSoproServiceServer interface {
	mustEmbedUnimplementedSoproServiceServer()
}

func RegisterSoproServiceServer(s grpc.ServiceRegistrar, srv SoproServiceServer) {
	s.RegisterService(&SoproService_ServiceDesc, srv)
}

func _SoproService_PingPong_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HealthRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SoproServiceServer).PingPong(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SoproService_PingPong_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SoproServiceServer).PingPong(ctx, req.(*HealthRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SoproService_ServiceDesc is the grpc.ServiceDesc for SoproService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SoproService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "sopropb.SoproService",
	HandlerType: (*SoproServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PingPong",
			Handler:    _SoproService_PingPong_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "sopro/sopro.proto",
}