// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.18.1
// source: eolymp/playground/playground.proto

package playground

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
	Playground_CreateRun_FullMethodName   = "/eolymp.playground.Playground/CreateRun"
	Playground_DescribeRun_FullMethodName = "/eolymp.playground.Playground/DescribeRun"
)

// PlaygroundClient is the client API for Playground service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PlaygroundClient interface {
	CreateRun(ctx context.Context, in *CreateRunInput, opts ...grpc.CallOption) (*CreateRunOutput, error)
	DescribeRun(ctx context.Context, in *DescribeRunInput, opts ...grpc.CallOption) (*DescribeRunOutput, error)
}

type playgroundClient struct {
	cc grpc.ClientConnInterface
}

func NewPlaygroundClient(cc grpc.ClientConnInterface) PlaygroundClient {
	return &playgroundClient{cc}
}

func (c *playgroundClient) CreateRun(ctx context.Context, in *CreateRunInput, opts ...grpc.CallOption) (*CreateRunOutput, error) {
	out := new(CreateRunOutput)
	err := c.cc.Invoke(ctx, Playground_CreateRun_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *playgroundClient) DescribeRun(ctx context.Context, in *DescribeRunInput, opts ...grpc.CallOption) (*DescribeRunOutput, error) {
	out := new(DescribeRunOutput)
	err := c.cc.Invoke(ctx, Playground_DescribeRun_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PlaygroundServer is the server API for Playground service.
// All implementations should embed UnimplementedPlaygroundServer
// for forward compatibility
type PlaygroundServer interface {
	CreateRun(context.Context, *CreateRunInput) (*CreateRunOutput, error)
	DescribeRun(context.Context, *DescribeRunInput) (*DescribeRunOutput, error)
}

// UnimplementedPlaygroundServer should be embedded to have forward compatible implementations.
type UnimplementedPlaygroundServer struct {
}

func (UnimplementedPlaygroundServer) CreateRun(context.Context, *CreateRunInput) (*CreateRunOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateRun not implemented")
}
func (UnimplementedPlaygroundServer) DescribeRun(context.Context, *DescribeRunInput) (*DescribeRunOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DescribeRun not implemented")
}

// UnsafePlaygroundServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PlaygroundServer will
// result in compilation errors.
type UnsafePlaygroundServer interface {
	mustEmbedUnimplementedPlaygroundServer()
}

func RegisterPlaygroundServer(s grpc.ServiceRegistrar, srv PlaygroundServer) {
	s.RegisterService(&Playground_ServiceDesc, srv)
}

func _Playground_CreateRun_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRunInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlaygroundServer).CreateRun(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Playground_CreateRun_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlaygroundServer).CreateRun(ctx, req.(*CreateRunInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _Playground_DescribeRun_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DescribeRunInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlaygroundServer).DescribeRun(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Playground_DescribeRun_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlaygroundServer).DescribeRun(ctx, req.(*DescribeRunInput))
	}
	return interceptor(ctx, in, info, handler)
}

// Playground_ServiceDesc is the grpc.ServiceDesc for Playground service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Playground_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "eolymp.playground.Playground",
	HandlerType: (*PlaygroundServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateRun",
			Handler:    _Playground_CreateRun_Handler,
		},
		{
			MethodName: "DescribeRun",
			Handler:    _Playground_DescribeRun_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "eolymp/playground/playground.proto",
}
