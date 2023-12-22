// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.18.1
// source: eolymp/content/render_service.proto

package content

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
	RenderService_RenderContent_FullMethodName = "/eolymp.content.RenderService/RenderContent"
)

// RenderServiceClient is the client API for RenderService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RenderServiceClient interface {
	RenderContent(ctx context.Context, in *RenderContentInput, opts ...grpc.CallOption) (*RenderContentOutput, error)
}

type renderServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRenderServiceClient(cc grpc.ClientConnInterface) RenderServiceClient {
	return &renderServiceClient{cc}
}

func (c *renderServiceClient) RenderContent(ctx context.Context, in *RenderContentInput, opts ...grpc.CallOption) (*RenderContentOutput, error) {
	out := new(RenderContentOutput)
	err := c.cc.Invoke(ctx, RenderService_RenderContent_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RenderServiceServer is the server API for RenderService service.
// All implementations should embed UnimplementedRenderServiceServer
// for forward compatibility
type RenderServiceServer interface {
	RenderContent(context.Context, *RenderContentInput) (*RenderContentOutput, error)
}

// UnimplementedRenderServiceServer should be embedded to have forward compatible implementations.
type UnimplementedRenderServiceServer struct {
}

func (UnimplementedRenderServiceServer) RenderContent(context.Context, *RenderContentInput) (*RenderContentOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RenderContent not implemented")
}

// UnsafeRenderServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RenderServiceServer will
// result in compilation errors.
type UnsafeRenderServiceServer interface {
	mustEmbedUnimplementedRenderServiceServer()
}

func RegisterRenderServiceServer(s grpc.ServiceRegistrar, srv RenderServiceServer) {
	s.RegisterService(&RenderService_ServiceDesc, srv)
}

func _RenderService_RenderContent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RenderContentInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RenderServiceServer).RenderContent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RenderService_RenderContent_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RenderServiceServer).RenderContent(ctx, req.(*RenderContentInput))
	}
	return interceptor(ctx, in, info, handler)
}

// RenderService_ServiceDesc is the grpc.ServiceDesc for RenderService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RenderService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "eolymp.content.RenderService",
	HandlerType: (*RenderServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RenderContent",
			Handler:    _RenderService_RenderContent_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "eolymp/content/render_service.proto",
}