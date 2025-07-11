// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.4
// source: eolymp/community/attribute_service.proto

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
	AttributeService_CreateAttribute_FullMethodName   = "/eolymp.community.AttributeService/CreateAttribute"
	AttributeService_UpdateAttribute_FullMethodName   = "/eolymp.community.AttributeService/UpdateAttribute"
	AttributeService_RemoveAttribute_FullMethodName   = "/eolymp.community.AttributeService/RemoveAttribute"
	AttributeService_DescribeAttribute_FullMethodName = "/eolymp.community.AttributeService/DescribeAttribute"
	AttributeService_ListAttributes_FullMethodName    = "/eolymp.community.AttributeService/ListAttributes"
)

// AttributeServiceClient is the client API for AttributeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AttributeServiceClient interface {
	CreateAttribute(ctx context.Context, in *CreateAttributeInput, opts ...grpc.CallOption) (*CreateAttributeOutput, error)
	UpdateAttribute(ctx context.Context, in *UpdateAttributeInput, opts ...grpc.CallOption) (*UpdateAttributeOutput, error)
	RemoveAttribute(ctx context.Context, in *RemoveAttributeInput, opts ...grpc.CallOption) (*RemoveAttributeOutput, error)
	DescribeAttribute(ctx context.Context, in *DescribeAttributeInput, opts ...grpc.CallOption) (*DescribeAttributeOutput, error)
	ListAttributes(ctx context.Context, in *ListAttributesInput, opts ...grpc.CallOption) (*ListAttributesOutput, error)
}

type attributeServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAttributeServiceClient(cc grpc.ClientConnInterface) AttributeServiceClient {
	return &attributeServiceClient{cc}
}

func (c *attributeServiceClient) CreateAttribute(ctx context.Context, in *CreateAttributeInput, opts ...grpc.CallOption) (*CreateAttributeOutput, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateAttributeOutput)
	err := c.cc.Invoke(ctx, AttributeService_CreateAttribute_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *attributeServiceClient) UpdateAttribute(ctx context.Context, in *UpdateAttributeInput, opts ...grpc.CallOption) (*UpdateAttributeOutput, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateAttributeOutput)
	err := c.cc.Invoke(ctx, AttributeService_UpdateAttribute_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *attributeServiceClient) RemoveAttribute(ctx context.Context, in *RemoveAttributeInput, opts ...grpc.CallOption) (*RemoveAttributeOutput, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RemoveAttributeOutput)
	err := c.cc.Invoke(ctx, AttributeService_RemoveAttribute_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *attributeServiceClient) DescribeAttribute(ctx context.Context, in *DescribeAttributeInput, opts ...grpc.CallOption) (*DescribeAttributeOutput, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DescribeAttributeOutput)
	err := c.cc.Invoke(ctx, AttributeService_DescribeAttribute_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *attributeServiceClient) ListAttributes(ctx context.Context, in *ListAttributesInput, opts ...grpc.CallOption) (*ListAttributesOutput, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListAttributesOutput)
	err := c.cc.Invoke(ctx, AttributeService_ListAttributes_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AttributeServiceServer is the server API for AttributeService service.
// All implementations should embed UnimplementedAttributeServiceServer
// for forward compatibility.
type AttributeServiceServer interface {
	CreateAttribute(context.Context, *CreateAttributeInput) (*CreateAttributeOutput, error)
	UpdateAttribute(context.Context, *UpdateAttributeInput) (*UpdateAttributeOutput, error)
	RemoveAttribute(context.Context, *RemoveAttributeInput) (*RemoveAttributeOutput, error)
	DescribeAttribute(context.Context, *DescribeAttributeInput) (*DescribeAttributeOutput, error)
	ListAttributes(context.Context, *ListAttributesInput) (*ListAttributesOutput, error)
}

// UnimplementedAttributeServiceServer should be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedAttributeServiceServer struct{}

func (UnimplementedAttributeServiceServer) CreateAttribute(context.Context, *CreateAttributeInput) (*CreateAttributeOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAttribute not implemented")
}
func (UnimplementedAttributeServiceServer) UpdateAttribute(context.Context, *UpdateAttributeInput) (*UpdateAttributeOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAttribute not implemented")
}
func (UnimplementedAttributeServiceServer) RemoveAttribute(context.Context, *RemoveAttributeInput) (*RemoveAttributeOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveAttribute not implemented")
}
func (UnimplementedAttributeServiceServer) DescribeAttribute(context.Context, *DescribeAttributeInput) (*DescribeAttributeOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DescribeAttribute not implemented")
}
func (UnimplementedAttributeServiceServer) ListAttributes(context.Context, *ListAttributesInput) (*ListAttributesOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListAttributes not implemented")
}
func (UnimplementedAttributeServiceServer) testEmbeddedByValue() {}

// UnsafeAttributeServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AttributeServiceServer will
// result in compilation errors.
type UnsafeAttributeServiceServer interface {
	mustEmbedUnimplementedAttributeServiceServer()
}

func RegisterAttributeServiceServer(s grpc.ServiceRegistrar, srv AttributeServiceServer) {
	// If the following call pancis, it indicates UnimplementedAttributeServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&AttributeService_ServiceDesc, srv)
}

func _AttributeService_CreateAttribute_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAttributeInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AttributeServiceServer).CreateAttribute(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AttributeService_CreateAttribute_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AttributeServiceServer).CreateAttribute(ctx, req.(*CreateAttributeInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _AttributeService_UpdateAttribute_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateAttributeInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AttributeServiceServer).UpdateAttribute(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AttributeService_UpdateAttribute_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AttributeServiceServer).UpdateAttribute(ctx, req.(*UpdateAttributeInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _AttributeService_RemoveAttribute_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveAttributeInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AttributeServiceServer).RemoveAttribute(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AttributeService_RemoveAttribute_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AttributeServiceServer).RemoveAttribute(ctx, req.(*RemoveAttributeInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _AttributeService_DescribeAttribute_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DescribeAttributeInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AttributeServiceServer).DescribeAttribute(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AttributeService_DescribeAttribute_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AttributeServiceServer).DescribeAttribute(ctx, req.(*DescribeAttributeInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _AttributeService_ListAttributes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListAttributesInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AttributeServiceServer).ListAttributes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AttributeService_ListAttributes_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AttributeServiceServer).ListAttributes(ctx, req.(*ListAttributesInput))
	}
	return interceptor(ctx, in, info, handler)
}

// AttributeService_ServiceDesc is the grpc.ServiceDesc for AttributeService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AttributeService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "eolymp.community.AttributeService",
	HandlerType: (*AttributeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateAttribute",
			Handler:    _AttributeService_CreateAttribute_Handler,
		},
		{
			MethodName: "UpdateAttribute",
			Handler:    _AttributeService_UpdateAttribute_Handler,
		},
		{
			MethodName: "RemoveAttribute",
			Handler:    _AttributeService_RemoveAttribute_Handler,
		},
		{
			MethodName: "DescribeAttribute",
			Handler:    _AttributeService_DescribeAttribute_Handler,
		},
		{
			MethodName: "ListAttributes",
			Handler:    _AttributeService_ListAttributes_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "eolymp/community/attribute_service.proto",
}
