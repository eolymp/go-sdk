// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.2
// source: eolymp/universe/universe.proto

package universe

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

// UniverseClient is the client API for Universe service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UniverseClient interface {
	// Create a space
	CreateSpace(ctx context.Context, in *CreateSpaceInput, opts ...grpc.CallOption) (*CreateSpaceOutput, error)
	// Update existing space
	UpdateSpace(ctx context.Context, in *UpdateSpaceInput, opts ...grpc.CallOption) (*UpdateSpaceOutput, error)
	// Delete space
	DeleteSpace(ctx context.Context, in *DeleteSpaceInput, opts ...grpc.CallOption) (*DeleteSpaceOutput, error)
	// Lookup space by domain key
	LookupSpace(ctx context.Context, in *LookupSpaceInput, opts ...grpc.CallOption) (*LookupSpaceOutput, error)
	// Describe space
	DescribeSpace(ctx context.Context, in *DescribeSpaceInput, opts ...grpc.CallOption) (*DescribeSpaceOutput, error)
	// Describe quota
	DescribeQuota(ctx context.Context, in *DescribeQuotaInput, opts ...grpc.CallOption) (*DescribeQuotaOutput, error)
	// List spaces of a contest
	ListSpaces(ctx context.Context, in *ListSpacesInput, opts ...grpc.CallOption) (*ListSpacesOutput, error)
	// Add space permission
	GrantPermission(ctx context.Context, in *GrantPermissionInput, opts ...grpc.CallOption) (*GrantPermissionOutput, error)
	// Delete space permission
	RevokePermission(ctx context.Context, in *RevokePermissionInput, opts ...grpc.CallOption) (*RevokePermissionOutput, error)
	// Describe space permission
	DescribePermission(ctx context.Context, in *DescribePermissionInput, opts ...grpc.CallOption) (*DescribePermissionOutput, error)
	// Describe space permission
	IntrospectPermission(ctx context.Context, in *IntrospectPermissionInput, opts ...grpc.CallOption) (*IntrospectPermissionOutput, error)
	// List permissions in a space
	ListPermissions(ctx context.Context, in *ListPermissionsInput, opts ...grpc.CallOption) (*ListPermissionsOutput, error)
}

type universeClient struct {
	cc grpc.ClientConnInterface
}

func NewUniverseClient(cc grpc.ClientConnInterface) UniverseClient {
	return &universeClient{cc}
}

func (c *universeClient) CreateSpace(ctx context.Context, in *CreateSpaceInput, opts ...grpc.CallOption) (*CreateSpaceOutput, error) {
	out := new(CreateSpaceOutput)
	err := c.cc.Invoke(ctx, "/eolymp.universe.Universe/CreateSpace", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *universeClient) UpdateSpace(ctx context.Context, in *UpdateSpaceInput, opts ...grpc.CallOption) (*UpdateSpaceOutput, error) {
	out := new(UpdateSpaceOutput)
	err := c.cc.Invoke(ctx, "/eolymp.universe.Universe/UpdateSpace", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *universeClient) DeleteSpace(ctx context.Context, in *DeleteSpaceInput, opts ...grpc.CallOption) (*DeleteSpaceOutput, error) {
	out := new(DeleteSpaceOutput)
	err := c.cc.Invoke(ctx, "/eolymp.universe.Universe/DeleteSpace", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *universeClient) LookupSpace(ctx context.Context, in *LookupSpaceInput, opts ...grpc.CallOption) (*LookupSpaceOutput, error) {
	out := new(LookupSpaceOutput)
	err := c.cc.Invoke(ctx, "/eolymp.universe.Universe/LookupSpace", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *universeClient) DescribeSpace(ctx context.Context, in *DescribeSpaceInput, opts ...grpc.CallOption) (*DescribeSpaceOutput, error) {
	out := new(DescribeSpaceOutput)
	err := c.cc.Invoke(ctx, "/eolymp.universe.Universe/DescribeSpace", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *universeClient) DescribeQuota(ctx context.Context, in *DescribeQuotaInput, opts ...grpc.CallOption) (*DescribeQuotaOutput, error) {
	out := new(DescribeQuotaOutput)
	err := c.cc.Invoke(ctx, "/eolymp.universe.Universe/DescribeQuota", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *universeClient) ListSpaces(ctx context.Context, in *ListSpacesInput, opts ...grpc.CallOption) (*ListSpacesOutput, error) {
	out := new(ListSpacesOutput)
	err := c.cc.Invoke(ctx, "/eolymp.universe.Universe/ListSpaces", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *universeClient) GrantPermission(ctx context.Context, in *GrantPermissionInput, opts ...grpc.CallOption) (*GrantPermissionOutput, error) {
	out := new(GrantPermissionOutput)
	err := c.cc.Invoke(ctx, "/eolymp.universe.Universe/GrantPermission", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *universeClient) RevokePermission(ctx context.Context, in *RevokePermissionInput, opts ...grpc.CallOption) (*RevokePermissionOutput, error) {
	out := new(RevokePermissionOutput)
	err := c.cc.Invoke(ctx, "/eolymp.universe.Universe/RevokePermission", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *universeClient) DescribePermission(ctx context.Context, in *DescribePermissionInput, opts ...grpc.CallOption) (*DescribePermissionOutput, error) {
	out := new(DescribePermissionOutput)
	err := c.cc.Invoke(ctx, "/eolymp.universe.Universe/DescribePermission", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *universeClient) IntrospectPermission(ctx context.Context, in *IntrospectPermissionInput, opts ...grpc.CallOption) (*IntrospectPermissionOutput, error) {
	out := new(IntrospectPermissionOutput)
	err := c.cc.Invoke(ctx, "/eolymp.universe.Universe/IntrospectPermission", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *universeClient) ListPermissions(ctx context.Context, in *ListPermissionsInput, opts ...grpc.CallOption) (*ListPermissionsOutput, error) {
	out := new(ListPermissionsOutput)
	err := c.cc.Invoke(ctx, "/eolymp.universe.Universe/ListPermissions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UniverseServer is the server API for Universe service.
// All implementations must embed UnimplementedUniverseServer
// for forward compatibility
type UniverseServer interface {
	// Create a space
	CreateSpace(context.Context, *CreateSpaceInput) (*CreateSpaceOutput, error)
	// Update existing space
	UpdateSpace(context.Context, *UpdateSpaceInput) (*UpdateSpaceOutput, error)
	// Delete space
	DeleteSpace(context.Context, *DeleteSpaceInput) (*DeleteSpaceOutput, error)
	// Lookup space by domain key
	LookupSpace(context.Context, *LookupSpaceInput) (*LookupSpaceOutput, error)
	// Describe space
	DescribeSpace(context.Context, *DescribeSpaceInput) (*DescribeSpaceOutput, error)
	// Describe quota
	DescribeQuota(context.Context, *DescribeQuotaInput) (*DescribeQuotaOutput, error)
	// List spaces of a contest
	ListSpaces(context.Context, *ListSpacesInput) (*ListSpacesOutput, error)
	// Add space permission
	GrantPermission(context.Context, *GrantPermissionInput) (*GrantPermissionOutput, error)
	// Delete space permission
	RevokePermission(context.Context, *RevokePermissionInput) (*RevokePermissionOutput, error)
	// Describe space permission
	DescribePermission(context.Context, *DescribePermissionInput) (*DescribePermissionOutput, error)
	// Describe space permission
	IntrospectPermission(context.Context, *IntrospectPermissionInput) (*IntrospectPermissionOutput, error)
	// List permissions in a space
	ListPermissions(context.Context, *ListPermissionsInput) (*ListPermissionsOutput, error)
	mustEmbedUnimplementedUniverseServer()
}

// UnimplementedUniverseServer must be embedded to have forward compatible implementations.
type UnimplementedUniverseServer struct {
}

func (UnimplementedUniverseServer) CreateSpace(context.Context, *CreateSpaceInput) (*CreateSpaceOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateSpace not implemented")
}
func (UnimplementedUniverseServer) UpdateSpace(context.Context, *UpdateSpaceInput) (*UpdateSpaceOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateSpace not implemented")
}
func (UnimplementedUniverseServer) DeleteSpace(context.Context, *DeleteSpaceInput) (*DeleteSpaceOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteSpace not implemented")
}
func (UnimplementedUniverseServer) LookupSpace(context.Context, *LookupSpaceInput) (*LookupSpaceOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LookupSpace not implemented")
}
func (UnimplementedUniverseServer) DescribeSpace(context.Context, *DescribeSpaceInput) (*DescribeSpaceOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DescribeSpace not implemented")
}
func (UnimplementedUniverseServer) DescribeQuota(context.Context, *DescribeQuotaInput) (*DescribeQuotaOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DescribeQuota not implemented")
}
func (UnimplementedUniverseServer) ListSpaces(context.Context, *ListSpacesInput) (*ListSpacesOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListSpaces not implemented")
}
func (UnimplementedUniverseServer) GrantPermission(context.Context, *GrantPermissionInput) (*GrantPermissionOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GrantPermission not implemented")
}
func (UnimplementedUniverseServer) RevokePermission(context.Context, *RevokePermissionInput) (*RevokePermissionOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RevokePermission not implemented")
}
func (UnimplementedUniverseServer) DescribePermission(context.Context, *DescribePermissionInput) (*DescribePermissionOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DescribePermission not implemented")
}
func (UnimplementedUniverseServer) IntrospectPermission(context.Context, *IntrospectPermissionInput) (*IntrospectPermissionOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IntrospectPermission not implemented")
}
func (UnimplementedUniverseServer) ListPermissions(context.Context, *ListPermissionsInput) (*ListPermissionsOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListPermissions not implemented")
}
func (UnimplementedUniverseServer) mustEmbedUnimplementedUniverseServer() {}

// UnsafeUniverseServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UniverseServer will
// result in compilation errors.
type UnsafeUniverseServer interface {
	mustEmbedUnimplementedUniverseServer()
}

func RegisterUniverseServer(s grpc.ServiceRegistrar, srv UniverseServer) {
	s.RegisterService(&Universe_ServiceDesc, srv)
}

func _Universe_CreateSpace_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateSpaceInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UniverseServer).CreateSpace(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/eolymp.universe.Universe/CreateSpace",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UniverseServer).CreateSpace(ctx, req.(*CreateSpaceInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _Universe_UpdateSpace_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateSpaceInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UniverseServer).UpdateSpace(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/eolymp.universe.Universe/UpdateSpace",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UniverseServer).UpdateSpace(ctx, req.(*UpdateSpaceInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _Universe_DeleteSpace_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteSpaceInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UniverseServer).DeleteSpace(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/eolymp.universe.Universe/DeleteSpace",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UniverseServer).DeleteSpace(ctx, req.(*DeleteSpaceInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _Universe_LookupSpace_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LookupSpaceInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UniverseServer).LookupSpace(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/eolymp.universe.Universe/LookupSpace",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UniverseServer).LookupSpace(ctx, req.(*LookupSpaceInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _Universe_DescribeSpace_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DescribeSpaceInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UniverseServer).DescribeSpace(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/eolymp.universe.Universe/DescribeSpace",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UniverseServer).DescribeSpace(ctx, req.(*DescribeSpaceInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _Universe_DescribeQuota_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DescribeQuotaInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UniverseServer).DescribeQuota(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/eolymp.universe.Universe/DescribeQuota",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UniverseServer).DescribeQuota(ctx, req.(*DescribeQuotaInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _Universe_ListSpaces_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListSpacesInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UniverseServer).ListSpaces(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/eolymp.universe.Universe/ListSpaces",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UniverseServer).ListSpaces(ctx, req.(*ListSpacesInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _Universe_GrantPermission_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GrantPermissionInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UniverseServer).GrantPermission(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/eolymp.universe.Universe/GrantPermission",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UniverseServer).GrantPermission(ctx, req.(*GrantPermissionInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _Universe_RevokePermission_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RevokePermissionInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UniverseServer).RevokePermission(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/eolymp.universe.Universe/RevokePermission",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UniverseServer).RevokePermission(ctx, req.(*RevokePermissionInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _Universe_DescribePermission_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DescribePermissionInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UniverseServer).DescribePermission(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/eolymp.universe.Universe/DescribePermission",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UniverseServer).DescribePermission(ctx, req.(*DescribePermissionInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _Universe_IntrospectPermission_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IntrospectPermissionInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UniverseServer).IntrospectPermission(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/eolymp.universe.Universe/IntrospectPermission",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UniverseServer).IntrospectPermission(ctx, req.(*IntrospectPermissionInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _Universe_ListPermissions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListPermissionsInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UniverseServer).ListPermissions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/eolymp.universe.Universe/ListPermissions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UniverseServer).ListPermissions(ctx, req.(*ListPermissionsInput))
	}
	return interceptor(ctx, in, info, handler)
}

// Universe_ServiceDesc is the grpc.ServiceDesc for Universe service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Universe_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "eolymp.universe.Universe",
	HandlerType: (*UniverseServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateSpace",
			Handler:    _Universe_CreateSpace_Handler,
		},
		{
			MethodName: "UpdateSpace",
			Handler:    _Universe_UpdateSpace_Handler,
		},
		{
			MethodName: "DeleteSpace",
			Handler:    _Universe_DeleteSpace_Handler,
		},
		{
			MethodName: "LookupSpace",
			Handler:    _Universe_LookupSpace_Handler,
		},
		{
			MethodName: "DescribeSpace",
			Handler:    _Universe_DescribeSpace_Handler,
		},
		{
			MethodName: "DescribeQuota",
			Handler:    _Universe_DescribeQuota_Handler,
		},
		{
			MethodName: "ListSpaces",
			Handler:    _Universe_ListSpaces_Handler,
		},
		{
			MethodName: "GrantPermission",
			Handler:    _Universe_GrantPermission_Handler,
		},
		{
			MethodName: "RevokePermission",
			Handler:    _Universe_RevokePermission_Handler,
		},
		{
			MethodName: "DescribePermission",
			Handler:    _Universe_DescribePermission_Handler,
		},
		{
			MethodName: "IntrospectPermission",
			Handler:    _Universe_IntrospectPermission_Handler,
		},
		{
			MethodName: "ListPermissions",
			Handler:    _Universe_ListPermissions_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "eolymp/universe/universe.proto",
}
