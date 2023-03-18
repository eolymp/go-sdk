// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.18.1
// source: eolymp/acl/acl.proto

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
	Acl_GrantPermission_FullMethodName      = "/eolymp.acl.Acl/GrantPermission"
	Acl_RevokePermission_FullMethodName     = "/eolymp.acl.Acl/RevokePermission"
	Acl_DescribePermission_FullMethodName   = "/eolymp.acl.Acl/DescribePermission"
	Acl_ListPermission_FullMethodName       = "/eolymp.acl.Acl/ListPermission"
	Acl_IntrospectPermission_FullMethodName = "/eolymp.acl.Acl/IntrospectPermission"
)

// AclClient is the client API for Acl service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AclClient interface {
	GrantPermission(ctx context.Context, in *GrantPermissionInput, opts ...grpc.CallOption) (*GrantPermissionOutput, error)
	RevokePermission(ctx context.Context, in *RevokePermissionInput, opts ...grpc.CallOption) (*RevokePermissionOutput, error)
	DescribePermission(ctx context.Context, in *DescribePermissionInput, opts ...grpc.CallOption) (*DescribePermissionOutput, error)
	ListPermission(ctx context.Context, in *ListPermissionInput, opts ...grpc.CallOption) (*ListPermissionOutput, error)
	IntrospectPermission(ctx context.Context, in *IntrospectPermissionInput, opts ...grpc.CallOption) (*IntrospectPermissionOutput, error)
}

type aclClient struct {
	cc grpc.ClientConnInterface
}

func NewAclClient(cc grpc.ClientConnInterface) AclClient {
	return &aclClient{cc}
}

func (c *aclClient) GrantPermission(ctx context.Context, in *GrantPermissionInput, opts ...grpc.CallOption) (*GrantPermissionOutput, error) {
	out := new(GrantPermissionOutput)
	err := c.cc.Invoke(ctx, Acl_GrantPermission_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aclClient) RevokePermission(ctx context.Context, in *RevokePermissionInput, opts ...grpc.CallOption) (*RevokePermissionOutput, error) {
	out := new(RevokePermissionOutput)
	err := c.cc.Invoke(ctx, Acl_RevokePermission_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aclClient) DescribePermission(ctx context.Context, in *DescribePermissionInput, opts ...grpc.CallOption) (*DescribePermissionOutput, error) {
	out := new(DescribePermissionOutput)
	err := c.cc.Invoke(ctx, Acl_DescribePermission_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aclClient) ListPermission(ctx context.Context, in *ListPermissionInput, opts ...grpc.CallOption) (*ListPermissionOutput, error) {
	out := new(ListPermissionOutput)
	err := c.cc.Invoke(ctx, Acl_ListPermission_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aclClient) IntrospectPermission(ctx context.Context, in *IntrospectPermissionInput, opts ...grpc.CallOption) (*IntrospectPermissionOutput, error) {
	out := new(IntrospectPermissionOutput)
	err := c.cc.Invoke(ctx, Acl_IntrospectPermission_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AclServer is the server API for Acl service.
// All implementations should embed UnimplementedAclServer
// for forward compatibility
type AclServer interface {
	GrantPermission(context.Context, *GrantPermissionInput) (*GrantPermissionOutput, error)
	RevokePermission(context.Context, *RevokePermissionInput) (*RevokePermissionOutput, error)
	DescribePermission(context.Context, *DescribePermissionInput) (*DescribePermissionOutput, error)
	ListPermission(context.Context, *ListPermissionInput) (*ListPermissionOutput, error)
	IntrospectPermission(context.Context, *IntrospectPermissionInput) (*IntrospectPermissionOutput, error)
}

// UnimplementedAclServer should be embedded to have forward compatible implementations.
type UnimplementedAclServer struct {
}

func (UnimplementedAclServer) GrantPermission(context.Context, *GrantPermissionInput) (*GrantPermissionOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GrantPermission not implemented")
}
func (UnimplementedAclServer) RevokePermission(context.Context, *RevokePermissionInput) (*RevokePermissionOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RevokePermission not implemented")
}
func (UnimplementedAclServer) DescribePermission(context.Context, *DescribePermissionInput) (*DescribePermissionOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DescribePermission not implemented")
}
func (UnimplementedAclServer) ListPermission(context.Context, *ListPermissionInput) (*ListPermissionOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListPermission not implemented")
}
func (UnimplementedAclServer) IntrospectPermission(context.Context, *IntrospectPermissionInput) (*IntrospectPermissionOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IntrospectPermission not implemented")
}

// UnsafeAclServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AclServer will
// result in compilation errors.
type UnsafeAclServer interface {
	mustEmbedUnimplementedAclServer()
}

func RegisterAclServer(s grpc.ServiceRegistrar, srv AclServer) {
	s.RegisterService(&Acl_ServiceDesc, srv)
}

func _Acl_GrantPermission_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GrantPermissionInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AclServer).GrantPermission(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Acl_GrantPermission_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AclServer).GrantPermission(ctx, req.(*GrantPermissionInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _Acl_RevokePermission_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RevokePermissionInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AclServer).RevokePermission(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Acl_RevokePermission_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AclServer).RevokePermission(ctx, req.(*RevokePermissionInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _Acl_DescribePermission_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DescribePermissionInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AclServer).DescribePermission(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Acl_DescribePermission_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AclServer).DescribePermission(ctx, req.(*DescribePermissionInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _Acl_ListPermission_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListPermissionInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AclServer).ListPermission(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Acl_ListPermission_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AclServer).ListPermission(ctx, req.(*ListPermissionInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _Acl_IntrospectPermission_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IntrospectPermissionInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AclServer).IntrospectPermission(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Acl_IntrospectPermission_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AclServer).IntrospectPermission(ctx, req.(*IntrospectPermissionInput))
	}
	return interceptor(ctx, in, info, handler)
}

// Acl_ServiceDesc is the grpc.ServiceDesc for Acl service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Acl_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "eolymp.acl.Acl",
	HandlerType: (*AclServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GrantPermission",
			Handler:    _Acl_GrantPermission_Handler,
		},
		{
			MethodName: "RevokePermission",
			Handler:    _Acl_RevokePermission_Handler,
		},
		{
			MethodName: "DescribePermission",
			Handler:    _Acl_DescribePermission_Handler,
		},
		{
			MethodName: "ListPermission",
			Handler:    _Acl_ListPermission_Handler,
		},
		{
			MethodName: "IntrospectPermission",
			Handler:    _Acl_IntrospectPermission_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "eolymp/acl/acl.proto",
}
