// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.21.12
// source: eolymp/course/material_service.proto

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
	MaterialService_CreateMaterial_FullMethodName   = "/eolymp.course.MaterialService/CreateMaterial"
	MaterialService_UpdateMaterial_FullMethodName   = "/eolymp.course.MaterialService/UpdateMaterial"
	MaterialService_MoveMaterial_FullMethodName     = "/eolymp.course.MaterialService/MoveMaterial"
	MaterialService_DeleteMaterial_FullMethodName   = "/eolymp.course.MaterialService/DeleteMaterial"
	MaterialService_DescribeMaterial_FullMethodName = "/eolymp.course.MaterialService/DescribeMaterial"
	MaterialService_ListMaterials_FullMethodName    = "/eolymp.course.MaterialService/ListMaterials"
)

// MaterialServiceClient is the client API for MaterialService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MaterialServiceClient interface {
	CreateMaterial(ctx context.Context, in *CreateMaterialInput, opts ...grpc.CallOption) (*CreateMaterialOutput, error)
	UpdateMaterial(ctx context.Context, in *UpdateMaterialInput, opts ...grpc.CallOption) (*UpdateMaterialOutput, error)
	MoveMaterial(ctx context.Context, in *MoveMaterialInput, opts ...grpc.CallOption) (*MoveMaterialOutput, error)
	DeleteMaterial(ctx context.Context, in *DeleteMaterialInput, opts ...grpc.CallOption) (*DeleteMaterialOutput, error)
	DescribeMaterial(ctx context.Context, in *DescribeMaterialInput, opts ...grpc.CallOption) (*DescribeMaterialOutput, error)
	ListMaterials(ctx context.Context, in *ListMaterialsInput, opts ...grpc.CallOption) (*ListMaterialsOutput, error)
}

type materialServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMaterialServiceClient(cc grpc.ClientConnInterface) MaterialServiceClient {
	return &materialServiceClient{cc}
}

func (c *materialServiceClient) CreateMaterial(ctx context.Context, in *CreateMaterialInput, opts ...grpc.CallOption) (*CreateMaterialOutput, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateMaterialOutput)
	err := c.cc.Invoke(ctx, MaterialService_CreateMaterial_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *materialServiceClient) UpdateMaterial(ctx context.Context, in *UpdateMaterialInput, opts ...grpc.CallOption) (*UpdateMaterialOutput, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateMaterialOutput)
	err := c.cc.Invoke(ctx, MaterialService_UpdateMaterial_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *materialServiceClient) MoveMaterial(ctx context.Context, in *MoveMaterialInput, opts ...grpc.CallOption) (*MoveMaterialOutput, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(MoveMaterialOutput)
	err := c.cc.Invoke(ctx, MaterialService_MoveMaterial_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *materialServiceClient) DeleteMaterial(ctx context.Context, in *DeleteMaterialInput, opts ...grpc.CallOption) (*DeleteMaterialOutput, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteMaterialOutput)
	err := c.cc.Invoke(ctx, MaterialService_DeleteMaterial_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *materialServiceClient) DescribeMaterial(ctx context.Context, in *DescribeMaterialInput, opts ...grpc.CallOption) (*DescribeMaterialOutput, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DescribeMaterialOutput)
	err := c.cc.Invoke(ctx, MaterialService_DescribeMaterial_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *materialServiceClient) ListMaterials(ctx context.Context, in *ListMaterialsInput, opts ...grpc.CallOption) (*ListMaterialsOutput, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListMaterialsOutput)
	err := c.cc.Invoke(ctx, MaterialService_ListMaterials_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MaterialServiceServer is the server API for MaterialService service.
// All implementations should embed UnimplementedMaterialServiceServer
// for forward compatibility.
type MaterialServiceServer interface {
	CreateMaterial(context.Context, *CreateMaterialInput) (*CreateMaterialOutput, error)
	UpdateMaterial(context.Context, *UpdateMaterialInput) (*UpdateMaterialOutput, error)
	MoveMaterial(context.Context, *MoveMaterialInput) (*MoveMaterialOutput, error)
	DeleteMaterial(context.Context, *DeleteMaterialInput) (*DeleteMaterialOutput, error)
	DescribeMaterial(context.Context, *DescribeMaterialInput) (*DescribeMaterialOutput, error)
	ListMaterials(context.Context, *ListMaterialsInput) (*ListMaterialsOutput, error)
}

// UnimplementedMaterialServiceServer should be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedMaterialServiceServer struct{}

func (UnimplementedMaterialServiceServer) CreateMaterial(context.Context, *CreateMaterialInput) (*CreateMaterialOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateMaterial not implemented")
}
func (UnimplementedMaterialServiceServer) UpdateMaterial(context.Context, *UpdateMaterialInput) (*UpdateMaterialOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateMaterial not implemented")
}
func (UnimplementedMaterialServiceServer) MoveMaterial(context.Context, *MoveMaterialInput) (*MoveMaterialOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MoveMaterial not implemented")
}
func (UnimplementedMaterialServiceServer) DeleteMaterial(context.Context, *DeleteMaterialInput) (*DeleteMaterialOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteMaterial not implemented")
}
func (UnimplementedMaterialServiceServer) DescribeMaterial(context.Context, *DescribeMaterialInput) (*DescribeMaterialOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DescribeMaterial not implemented")
}
func (UnimplementedMaterialServiceServer) ListMaterials(context.Context, *ListMaterialsInput) (*ListMaterialsOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListMaterials not implemented")
}
func (UnimplementedMaterialServiceServer) testEmbeddedByValue() {}

// UnsafeMaterialServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MaterialServiceServer will
// result in compilation errors.
type UnsafeMaterialServiceServer interface {
	mustEmbedUnimplementedMaterialServiceServer()
}

func RegisterMaterialServiceServer(s grpc.ServiceRegistrar, srv MaterialServiceServer) {
	// If the following call pancis, it indicates UnimplementedMaterialServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&MaterialService_ServiceDesc, srv)
}

func _MaterialService_CreateMaterial_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateMaterialInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MaterialServiceServer).CreateMaterial(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MaterialService_CreateMaterial_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MaterialServiceServer).CreateMaterial(ctx, req.(*CreateMaterialInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _MaterialService_UpdateMaterial_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateMaterialInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MaterialServiceServer).UpdateMaterial(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MaterialService_UpdateMaterial_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MaterialServiceServer).UpdateMaterial(ctx, req.(*UpdateMaterialInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _MaterialService_MoveMaterial_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MoveMaterialInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MaterialServiceServer).MoveMaterial(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MaterialService_MoveMaterial_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MaterialServiceServer).MoveMaterial(ctx, req.(*MoveMaterialInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _MaterialService_DeleteMaterial_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteMaterialInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MaterialServiceServer).DeleteMaterial(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MaterialService_DeleteMaterial_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MaterialServiceServer).DeleteMaterial(ctx, req.(*DeleteMaterialInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _MaterialService_DescribeMaterial_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DescribeMaterialInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MaterialServiceServer).DescribeMaterial(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MaterialService_DescribeMaterial_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MaterialServiceServer).DescribeMaterial(ctx, req.(*DescribeMaterialInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _MaterialService_ListMaterials_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListMaterialsInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MaterialServiceServer).ListMaterials(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MaterialService_ListMaterials_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MaterialServiceServer).ListMaterials(ctx, req.(*ListMaterialsInput))
	}
	return interceptor(ctx, in, info, handler)
}

// MaterialService_ServiceDesc is the grpc.ServiceDesc for MaterialService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MaterialService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "eolymp.course.MaterialService",
	HandlerType: (*MaterialServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateMaterial",
			Handler:    _MaterialService_CreateMaterial_Handler,
		},
		{
			MethodName: "UpdateMaterial",
			Handler:    _MaterialService_UpdateMaterial_Handler,
		},
		{
			MethodName: "MoveMaterial",
			Handler:    _MaterialService_MoveMaterial_Handler,
		},
		{
			MethodName: "DeleteMaterial",
			Handler:    _MaterialService_DeleteMaterial_Handler,
		},
		{
			MethodName: "DescribeMaterial",
			Handler:    _MaterialService_DescribeMaterial_Handler,
		},
		{
			MethodName: "ListMaterials",
			Handler:    _MaterialService_ListMaterials_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "eolymp/course/material_service.proto",
}