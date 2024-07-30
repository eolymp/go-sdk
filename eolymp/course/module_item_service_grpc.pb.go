// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.21.12
// source: eolymp/course/module_item_service.proto

package course

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
	ModuleItemService_CreateModuleItem_FullMethodName   = "/eolymp.course.ModuleItemService/CreateModuleItem"
	ModuleItemService_UpdateModuleItem_FullMethodName   = "/eolymp.course.ModuleItemService/UpdateModuleItem"
	ModuleItemService_MoveModuleItem_FullMethodName     = "/eolymp.course.ModuleItemService/MoveModuleItem"
	ModuleItemService_DeleteModuleItem_FullMethodName   = "/eolymp.course.ModuleItemService/DeleteModuleItem"
	ModuleItemService_DescribeModuleItem_FullMethodName = "/eolymp.course.ModuleItemService/DescribeModuleItem"
	ModuleItemService_ListModuleItems_FullMethodName    = "/eolymp.course.ModuleItemService/ListModuleItems"
)

// ModuleItemServiceClient is the client API for ModuleItemService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ModuleItemServiceClient interface {
	CreateModuleItem(ctx context.Context, in *CreateModuleItemInput, opts ...grpc.CallOption) (*CreateModuleItemOutput, error)
	UpdateModuleItem(ctx context.Context, in *UpdateModuleItemInput, opts ...grpc.CallOption) (*UpdateModuleItemOutput, error)
	MoveModuleItem(ctx context.Context, in *MoveModuleItemInput, opts ...grpc.CallOption) (*MoveModuleItemOutput, error)
	DeleteModuleItem(ctx context.Context, in *DeleteModuleItemInput, opts ...grpc.CallOption) (*DeleteModuleItemOutput, error)
	DescribeModuleItem(ctx context.Context, in *DescribeModuleItemInput, opts ...grpc.CallOption) (*DescribeModuleItemOutput, error)
	ListModuleItems(ctx context.Context, in *ListModuleItemsInput, opts ...grpc.CallOption) (*ListModuleItemsOutput, error)
}

type moduleItemServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewModuleItemServiceClient(cc grpc.ClientConnInterface) ModuleItemServiceClient {
	return &moduleItemServiceClient{cc}
}

func (c *moduleItemServiceClient) CreateModuleItem(ctx context.Context, in *CreateModuleItemInput, opts ...grpc.CallOption) (*CreateModuleItemOutput, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateModuleItemOutput)
	err := c.cc.Invoke(ctx, ModuleItemService_CreateModuleItem_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *moduleItemServiceClient) UpdateModuleItem(ctx context.Context, in *UpdateModuleItemInput, opts ...grpc.CallOption) (*UpdateModuleItemOutput, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateModuleItemOutput)
	err := c.cc.Invoke(ctx, ModuleItemService_UpdateModuleItem_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *moduleItemServiceClient) MoveModuleItem(ctx context.Context, in *MoveModuleItemInput, opts ...grpc.CallOption) (*MoveModuleItemOutput, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(MoveModuleItemOutput)
	err := c.cc.Invoke(ctx, ModuleItemService_MoveModuleItem_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *moduleItemServiceClient) DeleteModuleItem(ctx context.Context, in *DeleteModuleItemInput, opts ...grpc.CallOption) (*DeleteModuleItemOutput, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteModuleItemOutput)
	err := c.cc.Invoke(ctx, ModuleItemService_DeleteModuleItem_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *moduleItemServiceClient) DescribeModuleItem(ctx context.Context, in *DescribeModuleItemInput, opts ...grpc.CallOption) (*DescribeModuleItemOutput, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DescribeModuleItemOutput)
	err := c.cc.Invoke(ctx, ModuleItemService_DescribeModuleItem_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *moduleItemServiceClient) ListModuleItems(ctx context.Context, in *ListModuleItemsInput, opts ...grpc.CallOption) (*ListModuleItemsOutput, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListModuleItemsOutput)
	err := c.cc.Invoke(ctx, ModuleItemService_ListModuleItems_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ModuleItemServiceServer is the server API for ModuleItemService service.
// All implementations should embed UnimplementedModuleItemServiceServer
// for forward compatibility.
type ModuleItemServiceServer interface {
	CreateModuleItem(context.Context, *CreateModuleItemInput) (*CreateModuleItemOutput, error)
	UpdateModuleItem(context.Context, *UpdateModuleItemInput) (*UpdateModuleItemOutput, error)
	MoveModuleItem(context.Context, *MoveModuleItemInput) (*MoveModuleItemOutput, error)
	DeleteModuleItem(context.Context, *DeleteModuleItemInput) (*DeleteModuleItemOutput, error)
	DescribeModuleItem(context.Context, *DescribeModuleItemInput) (*DescribeModuleItemOutput, error)
	ListModuleItems(context.Context, *ListModuleItemsInput) (*ListModuleItemsOutput, error)
}

// UnimplementedModuleItemServiceServer should be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedModuleItemServiceServer struct{}

func (UnimplementedModuleItemServiceServer) CreateModuleItem(context.Context, *CreateModuleItemInput) (*CreateModuleItemOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateModuleItem not implemented")
}
func (UnimplementedModuleItemServiceServer) UpdateModuleItem(context.Context, *UpdateModuleItemInput) (*UpdateModuleItemOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateModuleItem not implemented")
}
func (UnimplementedModuleItemServiceServer) MoveModuleItem(context.Context, *MoveModuleItemInput) (*MoveModuleItemOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MoveModuleItem not implemented")
}
func (UnimplementedModuleItemServiceServer) DeleteModuleItem(context.Context, *DeleteModuleItemInput) (*DeleteModuleItemOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteModuleItem not implemented")
}
func (UnimplementedModuleItemServiceServer) DescribeModuleItem(context.Context, *DescribeModuleItemInput) (*DescribeModuleItemOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DescribeModuleItem not implemented")
}
func (UnimplementedModuleItemServiceServer) ListModuleItems(context.Context, *ListModuleItemsInput) (*ListModuleItemsOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListModuleItems not implemented")
}
func (UnimplementedModuleItemServiceServer) testEmbeddedByValue() {}

// UnsafeModuleItemServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ModuleItemServiceServer will
// result in compilation errors.
type UnsafeModuleItemServiceServer interface {
	mustEmbedUnimplementedModuleItemServiceServer()
}

func RegisterModuleItemServiceServer(s grpc.ServiceRegistrar, srv ModuleItemServiceServer) {
	// If the following call pancis, it indicates UnimplementedModuleItemServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&ModuleItemService_ServiceDesc, srv)
}

func _ModuleItemService_CreateModuleItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateModuleItemInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ModuleItemServiceServer).CreateModuleItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ModuleItemService_CreateModuleItem_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ModuleItemServiceServer).CreateModuleItem(ctx, req.(*CreateModuleItemInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _ModuleItemService_UpdateModuleItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateModuleItemInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ModuleItemServiceServer).UpdateModuleItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ModuleItemService_UpdateModuleItem_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ModuleItemServiceServer).UpdateModuleItem(ctx, req.(*UpdateModuleItemInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _ModuleItemService_MoveModuleItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MoveModuleItemInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ModuleItemServiceServer).MoveModuleItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ModuleItemService_MoveModuleItem_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ModuleItemServiceServer).MoveModuleItem(ctx, req.(*MoveModuleItemInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _ModuleItemService_DeleteModuleItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteModuleItemInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ModuleItemServiceServer).DeleteModuleItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ModuleItemService_DeleteModuleItem_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ModuleItemServiceServer).DeleteModuleItem(ctx, req.(*DeleteModuleItemInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _ModuleItemService_DescribeModuleItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DescribeModuleItemInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ModuleItemServiceServer).DescribeModuleItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ModuleItemService_DescribeModuleItem_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ModuleItemServiceServer).DescribeModuleItem(ctx, req.(*DescribeModuleItemInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _ModuleItemService_ListModuleItems_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListModuleItemsInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ModuleItemServiceServer).ListModuleItems(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ModuleItemService_ListModuleItems_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ModuleItemServiceServer).ListModuleItems(ctx, req.(*ListModuleItemsInput))
	}
	return interceptor(ctx, in, info, handler)
}

// ModuleItemService_ServiceDesc is the grpc.ServiceDesc for ModuleItemService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ModuleItemService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "eolymp.course.ModuleItemService",
	HandlerType: (*ModuleItemServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateModuleItem",
			Handler:    _ModuleItemService_CreateModuleItem_Handler,
		},
		{
			MethodName: "UpdateModuleItem",
			Handler:    _ModuleItemService_UpdateModuleItem_Handler,
		},
		{
			MethodName: "MoveModuleItem",
			Handler:    _ModuleItemService_MoveModuleItem_Handler,
		},
		{
			MethodName: "DeleteModuleItem",
			Handler:    _ModuleItemService_DeleteModuleItem_Handler,
		},
		{
			MethodName: "DescribeModuleItem",
			Handler:    _ModuleItemService_DescribeModuleItem_Handler,
		},
		{
			MethodName: "ListModuleItems",
			Handler:    _ModuleItemService_ListModuleItems_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "eolymp/course/module_item_service.proto",
}
