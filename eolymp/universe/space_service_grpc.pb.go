// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.4
// source: eolymp/universe/space_service.proto

package universe

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
	SpaceService_LookupSpace_FullMethodName   = "/eolymp.universe.SpaceService/LookupSpace"
	SpaceService_CreateSpace_FullMethodName   = "/eolymp.universe.SpaceService/CreateSpace"
	SpaceService_UpdateSpace_FullMethodName   = "/eolymp.universe.SpaceService/UpdateSpace"
	SpaceService_DeleteSpace_FullMethodName   = "/eolymp.universe.SpaceService/DeleteSpace"
	SpaceService_DescribeSpace_FullMethodName = "/eolymp.universe.SpaceService/DescribeSpace"
	SpaceService_ListSpaces_FullMethodName    = "/eolymp.universe.SpaceService/ListSpaces"
)

// SpaceServiceClient is the client API for SpaceService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SpaceServiceClient interface {
	// Lookup space by domain key
	LookupSpace(ctx context.Context, in *LookupSpaceInput, opts ...grpc.CallOption) (*LookupSpaceOutput, error)
	// Create a space
	CreateSpace(ctx context.Context, in *CreateSpaceInput, opts ...grpc.CallOption) (*CreateSpaceOutput, error)
	// Update existing space
	UpdateSpace(ctx context.Context, in *UpdateSpaceInput, opts ...grpc.CallOption) (*UpdateSpaceOutput, error)
	// Delete space
	DeleteSpace(ctx context.Context, in *DeleteSpaceInput, opts ...grpc.CallOption) (*DeleteSpaceOutput, error)
	// Describe space
	DescribeSpace(ctx context.Context, in *DescribeSpaceInput, opts ...grpc.CallOption) (*DescribeSpaceOutput, error)
	// List spaces of a contest
	ListSpaces(ctx context.Context, in *ListSpacesInput, opts ...grpc.CallOption) (*ListSpacesOutput, error)
}

type spaceServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSpaceServiceClient(cc grpc.ClientConnInterface) SpaceServiceClient {
	return &spaceServiceClient{cc}
}

func (c *spaceServiceClient) LookupSpace(ctx context.Context, in *LookupSpaceInput, opts ...grpc.CallOption) (*LookupSpaceOutput, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(LookupSpaceOutput)
	err := c.cc.Invoke(ctx, SpaceService_LookupSpace_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *spaceServiceClient) CreateSpace(ctx context.Context, in *CreateSpaceInput, opts ...grpc.CallOption) (*CreateSpaceOutput, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateSpaceOutput)
	err := c.cc.Invoke(ctx, SpaceService_CreateSpace_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *spaceServiceClient) UpdateSpace(ctx context.Context, in *UpdateSpaceInput, opts ...grpc.CallOption) (*UpdateSpaceOutput, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateSpaceOutput)
	err := c.cc.Invoke(ctx, SpaceService_UpdateSpace_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *spaceServiceClient) DeleteSpace(ctx context.Context, in *DeleteSpaceInput, opts ...grpc.CallOption) (*DeleteSpaceOutput, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteSpaceOutput)
	err := c.cc.Invoke(ctx, SpaceService_DeleteSpace_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *spaceServiceClient) DescribeSpace(ctx context.Context, in *DescribeSpaceInput, opts ...grpc.CallOption) (*DescribeSpaceOutput, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DescribeSpaceOutput)
	err := c.cc.Invoke(ctx, SpaceService_DescribeSpace_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *spaceServiceClient) ListSpaces(ctx context.Context, in *ListSpacesInput, opts ...grpc.CallOption) (*ListSpacesOutput, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListSpacesOutput)
	err := c.cc.Invoke(ctx, SpaceService_ListSpaces_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SpaceServiceServer is the server API for SpaceService service.
// All implementations should embed UnimplementedSpaceServiceServer
// for forward compatibility.
type SpaceServiceServer interface {
	// Lookup space by domain key
	LookupSpace(context.Context, *LookupSpaceInput) (*LookupSpaceOutput, error)
	// Create a space
	CreateSpace(context.Context, *CreateSpaceInput) (*CreateSpaceOutput, error)
	// Update existing space
	UpdateSpace(context.Context, *UpdateSpaceInput) (*UpdateSpaceOutput, error)
	// Delete space
	DeleteSpace(context.Context, *DeleteSpaceInput) (*DeleteSpaceOutput, error)
	// Describe space
	DescribeSpace(context.Context, *DescribeSpaceInput) (*DescribeSpaceOutput, error)
	// List spaces of a contest
	ListSpaces(context.Context, *ListSpacesInput) (*ListSpacesOutput, error)
}

// UnimplementedSpaceServiceServer should be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedSpaceServiceServer struct{}

func (UnimplementedSpaceServiceServer) LookupSpace(context.Context, *LookupSpaceInput) (*LookupSpaceOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LookupSpace not implemented")
}
func (UnimplementedSpaceServiceServer) CreateSpace(context.Context, *CreateSpaceInput) (*CreateSpaceOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateSpace not implemented")
}
func (UnimplementedSpaceServiceServer) UpdateSpace(context.Context, *UpdateSpaceInput) (*UpdateSpaceOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateSpace not implemented")
}
func (UnimplementedSpaceServiceServer) DeleteSpace(context.Context, *DeleteSpaceInput) (*DeleteSpaceOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteSpace not implemented")
}
func (UnimplementedSpaceServiceServer) DescribeSpace(context.Context, *DescribeSpaceInput) (*DescribeSpaceOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DescribeSpace not implemented")
}
func (UnimplementedSpaceServiceServer) ListSpaces(context.Context, *ListSpacesInput) (*ListSpacesOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListSpaces not implemented")
}
func (UnimplementedSpaceServiceServer) testEmbeddedByValue() {}

// UnsafeSpaceServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SpaceServiceServer will
// result in compilation errors.
type UnsafeSpaceServiceServer interface {
	mustEmbedUnimplementedSpaceServiceServer()
}

func RegisterSpaceServiceServer(s grpc.ServiceRegistrar, srv SpaceServiceServer) {
	// If the following call pancis, it indicates UnimplementedSpaceServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&SpaceService_ServiceDesc, srv)
}

func _SpaceService_LookupSpace_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LookupSpaceInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SpaceServiceServer).LookupSpace(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SpaceService_LookupSpace_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SpaceServiceServer).LookupSpace(ctx, req.(*LookupSpaceInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _SpaceService_CreateSpace_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateSpaceInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SpaceServiceServer).CreateSpace(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SpaceService_CreateSpace_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SpaceServiceServer).CreateSpace(ctx, req.(*CreateSpaceInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _SpaceService_UpdateSpace_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateSpaceInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SpaceServiceServer).UpdateSpace(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SpaceService_UpdateSpace_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SpaceServiceServer).UpdateSpace(ctx, req.(*UpdateSpaceInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _SpaceService_DeleteSpace_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteSpaceInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SpaceServiceServer).DeleteSpace(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SpaceService_DeleteSpace_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SpaceServiceServer).DeleteSpace(ctx, req.(*DeleteSpaceInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _SpaceService_DescribeSpace_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DescribeSpaceInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SpaceServiceServer).DescribeSpace(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SpaceService_DescribeSpace_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SpaceServiceServer).DescribeSpace(ctx, req.(*DescribeSpaceInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _SpaceService_ListSpaces_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListSpacesInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SpaceServiceServer).ListSpaces(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SpaceService_ListSpaces_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SpaceServiceServer).ListSpaces(ctx, req.(*ListSpacesInput))
	}
	return interceptor(ctx, in, info, handler)
}

// SpaceService_ServiceDesc is the grpc.ServiceDesc for SpaceService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SpaceService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "eolymp.universe.SpaceService",
	HandlerType: (*SpaceServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "LookupSpace",
			Handler:    _SpaceService_LookupSpace_Handler,
		},
		{
			MethodName: "CreateSpace",
			Handler:    _SpaceService_CreateSpace_Handler,
		},
		{
			MethodName: "UpdateSpace",
			Handler:    _SpaceService_UpdateSpace_Handler,
		},
		{
			MethodName: "DeleteSpace",
			Handler:    _SpaceService_DeleteSpace_Handler,
		},
		{
			MethodName: "DescribeSpace",
			Handler:    _SpaceService_DescribeSpace_Handler,
		},
		{
			MethodName: "ListSpaces",
			Handler:    _SpaceService_ListSpaces_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "eolymp/universe/space_service.proto",
}
