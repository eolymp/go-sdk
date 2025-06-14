// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.4
// source: eolymp/community/group_service.proto

package community

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
	GroupService_CreateGroup_FullMethodName   = "/eolymp.community.GroupService/CreateGroup"
	GroupService_UpdateGroup_FullMethodName   = "/eolymp.community.GroupService/UpdateGroup"
	GroupService_DeleteGroup_FullMethodName   = "/eolymp.community.GroupService/DeleteGroup"
	GroupService_DescribeGroup_FullMethodName = "/eolymp.community.GroupService/DescribeGroup"
	GroupService_ListGroups_FullMethodName    = "/eolymp.community.GroupService/ListGroups"
)

// GroupServiceClient is the client API for GroupService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// GroupService provides API to manage space groups.
type GroupServiceClient interface {
	CreateGroup(ctx context.Context, in *CreateGroupInput, opts ...grpc.CallOption) (*CreateGroupOutput, error)
	UpdateGroup(ctx context.Context, in *UpdateGroupInput, opts ...grpc.CallOption) (*UpdateGroupOutput, error)
	DeleteGroup(ctx context.Context, in *DeleteGroupInput, opts ...grpc.CallOption) (*DeleteGroupOutput, error)
	DescribeGroup(ctx context.Context, in *DescribeGroupInput, opts ...grpc.CallOption) (*DescribeGroupOutput, error)
	ListGroups(ctx context.Context, in *ListGroupsInput, opts ...grpc.CallOption) (*ListGroupsOutput, error)
}

type groupServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewGroupServiceClient(cc grpc.ClientConnInterface) GroupServiceClient {
	return &groupServiceClient{cc}
}

func (c *groupServiceClient) CreateGroup(ctx context.Context, in *CreateGroupInput, opts ...grpc.CallOption) (*CreateGroupOutput, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateGroupOutput)
	err := c.cc.Invoke(ctx, GroupService_CreateGroup_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupServiceClient) UpdateGroup(ctx context.Context, in *UpdateGroupInput, opts ...grpc.CallOption) (*UpdateGroupOutput, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateGroupOutput)
	err := c.cc.Invoke(ctx, GroupService_UpdateGroup_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupServiceClient) DeleteGroup(ctx context.Context, in *DeleteGroupInput, opts ...grpc.CallOption) (*DeleteGroupOutput, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteGroupOutput)
	err := c.cc.Invoke(ctx, GroupService_DeleteGroup_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupServiceClient) DescribeGroup(ctx context.Context, in *DescribeGroupInput, opts ...grpc.CallOption) (*DescribeGroupOutput, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DescribeGroupOutput)
	err := c.cc.Invoke(ctx, GroupService_DescribeGroup_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupServiceClient) ListGroups(ctx context.Context, in *ListGroupsInput, opts ...grpc.CallOption) (*ListGroupsOutput, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListGroupsOutput)
	err := c.cc.Invoke(ctx, GroupService_ListGroups_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GroupServiceServer is the server API for GroupService service.
// All implementations should embed UnimplementedGroupServiceServer
// for forward compatibility.
//
// GroupService provides API to manage space groups.
type GroupServiceServer interface {
	CreateGroup(context.Context, *CreateGroupInput) (*CreateGroupOutput, error)
	UpdateGroup(context.Context, *UpdateGroupInput) (*UpdateGroupOutput, error)
	DeleteGroup(context.Context, *DeleteGroupInput) (*DeleteGroupOutput, error)
	DescribeGroup(context.Context, *DescribeGroupInput) (*DescribeGroupOutput, error)
	ListGroups(context.Context, *ListGroupsInput) (*ListGroupsOutput, error)
}

// UnimplementedGroupServiceServer should be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedGroupServiceServer struct{}

func (UnimplementedGroupServiceServer) CreateGroup(context.Context, *CreateGroupInput) (*CreateGroupOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateGroup not implemented")
}
func (UnimplementedGroupServiceServer) UpdateGroup(context.Context, *UpdateGroupInput) (*UpdateGroupOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateGroup not implemented")
}
func (UnimplementedGroupServiceServer) DeleteGroup(context.Context, *DeleteGroupInput) (*DeleteGroupOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteGroup not implemented")
}
func (UnimplementedGroupServiceServer) DescribeGroup(context.Context, *DescribeGroupInput) (*DescribeGroupOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DescribeGroup not implemented")
}
func (UnimplementedGroupServiceServer) ListGroups(context.Context, *ListGroupsInput) (*ListGroupsOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListGroups not implemented")
}
func (UnimplementedGroupServiceServer) testEmbeddedByValue() {}

// UnsafeGroupServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GroupServiceServer will
// result in compilation errors.
type UnsafeGroupServiceServer interface {
	mustEmbedUnimplementedGroupServiceServer()
}

func RegisterGroupServiceServer(s grpc.ServiceRegistrar, srv GroupServiceServer) {
	// If the following call pancis, it indicates UnimplementedGroupServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&GroupService_ServiceDesc, srv)
}

func _GroupService_CreateGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateGroupInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupServiceServer).CreateGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GroupService_CreateGroup_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupServiceServer).CreateGroup(ctx, req.(*CreateGroupInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _GroupService_UpdateGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateGroupInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupServiceServer).UpdateGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GroupService_UpdateGroup_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupServiceServer).UpdateGroup(ctx, req.(*UpdateGroupInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _GroupService_DeleteGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteGroupInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupServiceServer).DeleteGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GroupService_DeleteGroup_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupServiceServer).DeleteGroup(ctx, req.(*DeleteGroupInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _GroupService_DescribeGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DescribeGroupInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupServiceServer).DescribeGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GroupService_DescribeGroup_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupServiceServer).DescribeGroup(ctx, req.(*DescribeGroupInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _GroupService_ListGroups_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListGroupsInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupServiceServer).ListGroups(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GroupService_ListGroups_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupServiceServer).ListGroups(ctx, req.(*ListGroupsInput))
	}
	return interceptor(ctx, in, info, handler)
}

// GroupService_ServiceDesc is the grpc.ServiceDesc for GroupService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GroupService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "eolymp.community.GroupService",
	HandlerType: (*GroupServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateGroup",
			Handler:    _GroupService_CreateGroup_Handler,
		},
		{
			MethodName: "UpdateGroup",
			Handler:    _GroupService_UpdateGroup_Handler,
		},
		{
			MethodName: "DeleteGroup",
			Handler:    _GroupService_DeleteGroup_Handler,
		},
		{
			MethodName: "DescribeGroup",
			Handler:    _GroupService_DescribeGroup_Handler,
		},
		{
			MethodName: "ListGroups",
			Handler:    _GroupService_ListGroups_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "eolymp/community/group_service.proto",
}
