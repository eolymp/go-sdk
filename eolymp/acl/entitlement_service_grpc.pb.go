// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: eolymp/acl/entitlement_service.proto

package acl

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
	EntitlementService_IntrospectEntitlements_FullMethodName = "/eolymp.acl.EntitlementService/IntrospectEntitlements"
)

// EntitlementServiceClient is the client API for EntitlementService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EntitlementServiceClient interface {
	IntrospectEntitlements(ctx context.Context, in *IntrospectEntitlementsInput, opts ...grpc.CallOption) (*IntrospectEntitlementsOutput, error)
}

type entitlementServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewEntitlementServiceClient(cc grpc.ClientConnInterface) EntitlementServiceClient {
	return &entitlementServiceClient{cc}
}

func (c *entitlementServiceClient) IntrospectEntitlements(ctx context.Context, in *IntrospectEntitlementsInput, opts ...grpc.CallOption) (*IntrospectEntitlementsOutput, error) {
	out := new(IntrospectEntitlementsOutput)
	err := c.cc.Invoke(ctx, EntitlementService_IntrospectEntitlements_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EntitlementServiceServer is the server API for EntitlementService service.
// All implementations should embed UnimplementedEntitlementServiceServer
// for forward compatibility
type EntitlementServiceServer interface {
	IntrospectEntitlements(context.Context, *IntrospectEntitlementsInput) (*IntrospectEntitlementsOutput, error)
}

// UnimplementedEntitlementServiceServer should be embedded to have forward compatible implementations.
type UnimplementedEntitlementServiceServer struct {
}

func (UnimplementedEntitlementServiceServer) IntrospectEntitlements(context.Context, *IntrospectEntitlementsInput) (*IntrospectEntitlementsOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IntrospectEntitlements not implemented")
}

// UnsafeEntitlementServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EntitlementServiceServer will
// result in compilation errors.
type UnsafeEntitlementServiceServer interface {
	mustEmbedUnimplementedEntitlementServiceServer()
}

func RegisterEntitlementServiceServer(s grpc.ServiceRegistrar, srv EntitlementServiceServer) {
	s.RegisterService(&EntitlementService_ServiceDesc, srv)
}

func _EntitlementService_IntrospectEntitlements_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IntrospectEntitlementsInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EntitlementServiceServer).IntrospectEntitlements(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EntitlementService_IntrospectEntitlements_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EntitlementServiceServer).IntrospectEntitlements(ctx, req.(*IntrospectEntitlementsInput))
	}
	return interceptor(ctx, in, info, handler)
}

// EntitlementService_ServiceDesc is the grpc.ServiceDesc for EntitlementService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var EntitlementService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "eolymp.acl.EntitlementService",
	HandlerType: (*EntitlementServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "IntrospectEntitlements",
			Handler:    _EntitlementService_IntrospectEntitlements_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "eolymp/acl/entitlement_service.proto",
}
