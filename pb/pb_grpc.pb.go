// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: pb/pb.proto

package pb

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

// JwtTokenServiceClient is the client API for JwtTokenService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type JwtTokenServiceClient interface {
	CheckToken(ctx context.Context, in *JwtRequest, opts ...grpc.CallOption) (*JwtResponse, error)
}

type jwtTokenServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewJwtTokenServiceClient(cc grpc.ClientConnInterface) JwtTokenServiceClient {
	return &jwtTokenServiceClient{cc}
}

func (c *jwtTokenServiceClient) CheckToken(ctx context.Context, in *JwtRequest, opts ...grpc.CallOption) (*JwtResponse, error) {
	out := new(JwtResponse)
	err := c.cc.Invoke(ctx, "/pb.JwtTokenService/CheckToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// JwtTokenServiceServer is the server API for JwtTokenService service.
// All implementations must embed UnimplementedJwtTokenServiceServer
// for forward compatibility
type JwtTokenServiceServer interface {
	CheckToken(context.Context, *JwtRequest) (*JwtResponse, error)
	mustEmbedUnimplementedJwtTokenServiceServer()
}

// UnimplementedJwtTokenServiceServer must be embedded to have forward compatible implementations.
type UnimplementedJwtTokenServiceServer struct {
}

func (UnimplementedJwtTokenServiceServer) CheckToken(context.Context, *JwtRequest) (*JwtResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckToken not implemented")
}
func (UnimplementedJwtTokenServiceServer) mustEmbedUnimplementedJwtTokenServiceServer() {}

// UnsafeJwtTokenServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to JwtTokenServiceServer will
// result in compilation errors.
type UnsafeJwtTokenServiceServer interface {
	mustEmbedUnimplementedJwtTokenServiceServer()
}

func RegisterJwtTokenServiceServer(s grpc.ServiceRegistrar, srv JwtTokenServiceServer) {
	s.RegisterService(&JwtTokenService_ServiceDesc, srv)
}

func _JwtTokenService_CheckToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JwtRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JwtTokenServiceServer).CheckToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.JwtTokenService/CheckToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JwtTokenServiceServer).CheckToken(ctx, req.(*JwtRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// JwtTokenService_ServiceDesc is the grpc.ServiceDesc for JwtTokenService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var JwtTokenService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.JwtTokenService",
	HandlerType: (*JwtTokenServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CheckToken",
			Handler:    _JwtTokenService_CheckToken_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pb/pb.proto",
}
