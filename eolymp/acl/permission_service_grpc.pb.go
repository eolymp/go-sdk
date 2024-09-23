// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.21.12
// source: eolymp/acl/permission_service.proto

package acl

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	PermissionService_IntrospectPermissions_FullMethodName = "/eolymp.acl.PermissionService/IntrospectPermissions"
)

// PermissionServiceClient is the client API for PermissionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PermissionServiceClient interface {
	IntrospectPermissions(ctx context.Context, in *IntrospectPermissionsInput, opts ...grpc.CallOption) (*IntrospectPermissionsOutput, error)
}

type permissionServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPermissionServiceClient(cc grpc.ClientConnInterface) PermissionServiceClient {
	return &permissionServiceClient{cc}
}

func (c *permissionServiceClient) IntrospectPermissions(ctx context.Context, in *IntrospectPermissionsInput, opts ...grpc.CallOption) (*IntrospectPermissionsOutput, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(IntrospectPermissionsOutput)
	err := c.cc.Invoke(ctx, PermissionService_IntrospectPermissions_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PermissionServiceServer is the server API for PermissionService service.
// All implementations should embed UnimplementedPermissionServiceServer
// for forward compatibility.
type PermissionServiceServer interface {
	IntrospectPermissions(context.Context, *IntrospectPermissionsInput) (*IntrospectPermissionsOutput, error)
}

// UnimplementedPermissionServiceServer should be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedPermissionServiceServer struct{}

func (UnimplementedPermissionServiceServer) IntrospectPermissions(context.Context, *IntrospectPermissionsInput) (*IntrospectPermissionsOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IntrospectPermissions not implemented")
}
func (UnimplementedPermissionServiceServer) testEmbeddedByValue() {}

// UnsafePermissionServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PermissionServiceServer will
// result in compilation errors.
type UnsafePermissionServiceServer interface {
	mustEmbedUnimplementedPermissionServiceServer()
}

func RegisterPermissionServiceServer(s grpc.ServiceRegistrar, srv PermissionServiceServer) {
	// If the following call pancis, it indicates UnimplementedPermissionServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&PermissionService_ServiceDesc, srv)
}

func _PermissionService_IntrospectPermissions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IntrospectPermissionsInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PermissionServiceServer).IntrospectPermissions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PermissionService_IntrospectPermissions_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PermissionServiceServer).IntrospectPermissions(ctx, req.(*IntrospectPermissionsInput))
	}
	return interceptor(ctx, in, info, handler)
}

// PermissionService_ServiceDesc is the grpc.ServiceDesc for PermissionService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PermissionService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "eolymp.acl.PermissionService",
	HandlerType: (*PermissionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "IntrospectPermissions",
			Handler:    _PermissionService_IntrospectPermissions_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "eolymp/acl/permission_service.proto",
}