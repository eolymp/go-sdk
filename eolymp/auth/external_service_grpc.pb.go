// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.18.1
// source: eolymp/auth/external_service.proto

package auth

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
	ExternalService_AuthorizeRequest_FullMethodName  = "/eolymp.auth.ExternalService/AuthorizeRequest"
	ExternalService_AuthorizeCallback_FullMethodName = "/eolymp.auth.ExternalService/AuthorizeCallback"
)

// ExternalServiceClient is the client API for ExternalService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ExternalServiceClient interface {
	AuthorizeRequest(ctx context.Context, in *AuthorizeRequestInput, opts ...grpc.CallOption) (*AuthorizeRequestOutput, error)
	AuthorizeCallback(ctx context.Context, in *AuthorizeCallbackInput, opts ...grpc.CallOption) (*AuthorizeCallbackOutput, error)
}

type externalServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewExternalServiceClient(cc grpc.ClientConnInterface) ExternalServiceClient {
	return &externalServiceClient{cc}
}

func (c *externalServiceClient) AuthorizeRequest(ctx context.Context, in *AuthorizeRequestInput, opts ...grpc.CallOption) (*AuthorizeRequestOutput, error) {
	out := new(AuthorizeRequestOutput)
	err := c.cc.Invoke(ctx, ExternalService_AuthorizeRequest_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *externalServiceClient) AuthorizeCallback(ctx context.Context, in *AuthorizeCallbackInput, opts ...grpc.CallOption) (*AuthorizeCallbackOutput, error) {
	out := new(AuthorizeCallbackOutput)
	err := c.cc.Invoke(ctx, ExternalService_AuthorizeCallback_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ExternalServiceServer is the server API for ExternalService service.
// All implementations should embed UnimplementedExternalServiceServer
// for forward compatibility
type ExternalServiceServer interface {
	AuthorizeRequest(context.Context, *AuthorizeRequestInput) (*AuthorizeRequestOutput, error)
	AuthorizeCallback(context.Context, *AuthorizeCallbackInput) (*AuthorizeCallbackOutput, error)
}

// UnimplementedExternalServiceServer should be embedded to have forward compatible implementations.
type UnimplementedExternalServiceServer struct {
}

func (UnimplementedExternalServiceServer) AuthorizeRequest(context.Context, *AuthorizeRequestInput) (*AuthorizeRequestOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AuthorizeRequest not implemented")
}
func (UnimplementedExternalServiceServer) AuthorizeCallback(context.Context, *AuthorizeCallbackInput) (*AuthorizeCallbackOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AuthorizeCallback not implemented")
}

// UnsafeExternalServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ExternalServiceServer will
// result in compilation errors.
type UnsafeExternalServiceServer interface {
	mustEmbedUnimplementedExternalServiceServer()
}

func RegisterExternalServiceServer(s grpc.ServiceRegistrar, srv ExternalServiceServer) {
	s.RegisterService(&ExternalService_ServiceDesc, srv)
}

func _ExternalService_AuthorizeRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthorizeRequestInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExternalServiceServer).AuthorizeRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ExternalService_AuthorizeRequest_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExternalServiceServer).AuthorizeRequest(ctx, req.(*AuthorizeRequestInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExternalService_AuthorizeCallback_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthorizeCallbackInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExternalServiceServer).AuthorizeCallback(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ExternalService_AuthorizeCallback_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExternalServiceServer).AuthorizeCallback(ctx, req.(*AuthorizeCallbackInput))
	}
	return interceptor(ctx, in, info, handler)
}

// ExternalService_ServiceDesc is the grpc.ServiceDesc for ExternalService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ExternalService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "eolymp.auth.ExternalService",
	HandlerType: (*ExternalServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AuthorizeRequest",
			Handler:    _ExternalService_AuthorizeRequest_Handler,
		},
		{
			MethodName: "AuthorizeCallback",
			Handler:    _ExternalService_AuthorizeCallback_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "eolymp/auth/external_service.proto",
}
