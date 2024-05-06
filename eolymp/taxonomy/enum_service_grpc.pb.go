// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: eolymp/taxonomy/enum_service.proto

package taxonomy

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
	EnumService_CreateEnum_FullMethodName        = "/eolymp.taxonomy.EnumService/CreateEnum"
	EnumService_DeleteEnum_FullMethodName        = "/eolymp.taxonomy.EnumService/DeleteEnum"
	EnumService_UpdateEnum_FullMethodName        = "/eolymp.taxonomy.EnumService/UpdateEnum"
	EnumService_DescribeEnum_FullMethodName      = "/eolymp.taxonomy.EnumService/DescribeEnum"
	EnumService_ListEnums_FullMethodName         = "/eolymp.taxonomy.EnumService/ListEnums"
	EnumService_CreateValue_FullMethodName       = "/eolymp.taxonomy.EnumService/CreateValue"
	EnumService_DeleteValue_FullMethodName       = "/eolymp.taxonomy.EnumService/DeleteValue"
	EnumService_UpdateValue_FullMethodName       = "/eolymp.taxonomy.EnumService/UpdateValue"
	EnumService_DescribeValue_FullMethodName     = "/eolymp.taxonomy.EnumService/DescribeValue"
	EnumService_ListValues_FullMethodName        = "/eolymp.taxonomy.EnumService/ListValues"
	EnumService_TranslateValue_FullMethodName    = "/eolymp.taxonomy.EnumService/TranslateValue"
	EnumService_DeleteTranslation_FullMethodName = "/eolymp.taxonomy.EnumService/DeleteTranslation"
	EnumService_ListTranslations_FullMethodName  = "/eolymp.taxonomy.EnumService/ListTranslations"
)

// EnumServiceClient is the client API for EnumService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EnumServiceClient interface {
	CreateEnum(ctx context.Context, in *CreateEnumInput, opts ...grpc.CallOption) (*CreateEnumOutput, error)
	DeleteEnum(ctx context.Context, in *DeleteEnumInput, opts ...grpc.CallOption) (*DeleteEnumOutput, error)
	UpdateEnum(ctx context.Context, in *UpdateEnumInput, opts ...grpc.CallOption) (*UpdateEnumOutput, error)
	DescribeEnum(ctx context.Context, in *DescribeEnumInput, opts ...grpc.CallOption) (*DescribeEnumOutput, error)
	ListEnums(ctx context.Context, in *ListEnumsInput, opts ...grpc.CallOption) (*ListEnumsOutput, error)
	CreateValue(ctx context.Context, in *CreateValueInput, opts ...grpc.CallOption) (*CreateValueOutput, error)
	DeleteValue(ctx context.Context, in *DeleteValueInput, opts ...grpc.CallOption) (*DeleteValueOutput, error)
	UpdateValue(ctx context.Context, in *UpdateValueInput, opts ...grpc.CallOption) (*UpdateValueOutput, error)
	DescribeValue(ctx context.Context, in *DescribeValueInput, opts ...grpc.CallOption) (*DescribeValueOutput, error)
	ListValues(ctx context.Context, in *ListValuesInput, opts ...grpc.CallOption) (*ListValuesOutput, error)
	TranslateValue(ctx context.Context, in *TranslateValueInput, opts ...grpc.CallOption) (*TranslateValueOutput, error)
	DeleteTranslation(ctx context.Context, in *DeleteTranslationInput, opts ...grpc.CallOption) (*DeleteTranslationOutput, error)
	ListTranslations(ctx context.Context, in *ListTranslationsInput, opts ...grpc.CallOption) (*ListTranslationsOutput, error)
}

type enumServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewEnumServiceClient(cc grpc.ClientConnInterface) EnumServiceClient {
	return &enumServiceClient{cc}
}

func (c *enumServiceClient) CreateEnum(ctx context.Context, in *CreateEnumInput, opts ...grpc.CallOption) (*CreateEnumOutput, error) {
	out := new(CreateEnumOutput)
	err := c.cc.Invoke(ctx, EnumService_CreateEnum_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *enumServiceClient) DeleteEnum(ctx context.Context, in *DeleteEnumInput, opts ...grpc.CallOption) (*DeleteEnumOutput, error) {
	out := new(DeleteEnumOutput)
	err := c.cc.Invoke(ctx, EnumService_DeleteEnum_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *enumServiceClient) UpdateEnum(ctx context.Context, in *UpdateEnumInput, opts ...grpc.CallOption) (*UpdateEnumOutput, error) {
	out := new(UpdateEnumOutput)
	err := c.cc.Invoke(ctx, EnumService_UpdateEnum_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *enumServiceClient) DescribeEnum(ctx context.Context, in *DescribeEnumInput, opts ...grpc.CallOption) (*DescribeEnumOutput, error) {
	out := new(DescribeEnumOutput)
	err := c.cc.Invoke(ctx, EnumService_DescribeEnum_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *enumServiceClient) ListEnums(ctx context.Context, in *ListEnumsInput, opts ...grpc.CallOption) (*ListEnumsOutput, error) {
	out := new(ListEnumsOutput)
	err := c.cc.Invoke(ctx, EnumService_ListEnums_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *enumServiceClient) CreateValue(ctx context.Context, in *CreateValueInput, opts ...grpc.CallOption) (*CreateValueOutput, error) {
	out := new(CreateValueOutput)
	err := c.cc.Invoke(ctx, EnumService_CreateValue_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *enumServiceClient) DeleteValue(ctx context.Context, in *DeleteValueInput, opts ...grpc.CallOption) (*DeleteValueOutput, error) {
	out := new(DeleteValueOutput)
	err := c.cc.Invoke(ctx, EnumService_DeleteValue_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *enumServiceClient) UpdateValue(ctx context.Context, in *UpdateValueInput, opts ...grpc.CallOption) (*UpdateValueOutput, error) {
	out := new(UpdateValueOutput)
	err := c.cc.Invoke(ctx, EnumService_UpdateValue_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *enumServiceClient) DescribeValue(ctx context.Context, in *DescribeValueInput, opts ...grpc.CallOption) (*DescribeValueOutput, error) {
	out := new(DescribeValueOutput)
	err := c.cc.Invoke(ctx, EnumService_DescribeValue_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *enumServiceClient) ListValues(ctx context.Context, in *ListValuesInput, opts ...grpc.CallOption) (*ListValuesOutput, error) {
	out := new(ListValuesOutput)
	err := c.cc.Invoke(ctx, EnumService_ListValues_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *enumServiceClient) TranslateValue(ctx context.Context, in *TranslateValueInput, opts ...grpc.CallOption) (*TranslateValueOutput, error) {
	out := new(TranslateValueOutput)
	err := c.cc.Invoke(ctx, EnumService_TranslateValue_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *enumServiceClient) DeleteTranslation(ctx context.Context, in *DeleteTranslationInput, opts ...grpc.CallOption) (*DeleteTranslationOutput, error) {
	out := new(DeleteTranslationOutput)
	err := c.cc.Invoke(ctx, EnumService_DeleteTranslation_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *enumServiceClient) ListTranslations(ctx context.Context, in *ListTranslationsInput, opts ...grpc.CallOption) (*ListTranslationsOutput, error) {
	out := new(ListTranslationsOutput)
	err := c.cc.Invoke(ctx, EnumService_ListTranslations_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EnumServiceServer is the server API for EnumService service.
// All implementations should embed UnimplementedEnumServiceServer
// for forward compatibility
type EnumServiceServer interface {
	CreateEnum(context.Context, *CreateEnumInput) (*CreateEnumOutput, error)
	DeleteEnum(context.Context, *DeleteEnumInput) (*DeleteEnumOutput, error)
	UpdateEnum(context.Context, *UpdateEnumInput) (*UpdateEnumOutput, error)
	DescribeEnum(context.Context, *DescribeEnumInput) (*DescribeEnumOutput, error)
	ListEnums(context.Context, *ListEnumsInput) (*ListEnumsOutput, error)
	CreateValue(context.Context, *CreateValueInput) (*CreateValueOutput, error)
	DeleteValue(context.Context, *DeleteValueInput) (*DeleteValueOutput, error)
	UpdateValue(context.Context, *UpdateValueInput) (*UpdateValueOutput, error)
	DescribeValue(context.Context, *DescribeValueInput) (*DescribeValueOutput, error)
	ListValues(context.Context, *ListValuesInput) (*ListValuesOutput, error)
	TranslateValue(context.Context, *TranslateValueInput) (*TranslateValueOutput, error)
	DeleteTranslation(context.Context, *DeleteTranslationInput) (*DeleteTranslationOutput, error)
	ListTranslations(context.Context, *ListTranslationsInput) (*ListTranslationsOutput, error)
}

// UnimplementedEnumServiceServer should be embedded to have forward compatible implementations.
type UnimplementedEnumServiceServer struct {
}

func (UnimplementedEnumServiceServer) CreateEnum(context.Context, *CreateEnumInput) (*CreateEnumOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateEnum not implemented")
}
func (UnimplementedEnumServiceServer) DeleteEnum(context.Context, *DeleteEnumInput) (*DeleteEnumOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteEnum not implemented")
}
func (UnimplementedEnumServiceServer) UpdateEnum(context.Context, *UpdateEnumInput) (*UpdateEnumOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateEnum not implemented")
}
func (UnimplementedEnumServiceServer) DescribeEnum(context.Context, *DescribeEnumInput) (*DescribeEnumOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DescribeEnum not implemented")
}
func (UnimplementedEnumServiceServer) ListEnums(context.Context, *ListEnumsInput) (*ListEnumsOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListEnums not implemented")
}
func (UnimplementedEnumServiceServer) CreateValue(context.Context, *CreateValueInput) (*CreateValueOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateValue not implemented")
}
func (UnimplementedEnumServiceServer) DeleteValue(context.Context, *DeleteValueInput) (*DeleteValueOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteValue not implemented")
}
func (UnimplementedEnumServiceServer) UpdateValue(context.Context, *UpdateValueInput) (*UpdateValueOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateValue not implemented")
}
func (UnimplementedEnumServiceServer) DescribeValue(context.Context, *DescribeValueInput) (*DescribeValueOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DescribeValue not implemented")
}
func (UnimplementedEnumServiceServer) ListValues(context.Context, *ListValuesInput) (*ListValuesOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListValues not implemented")
}
func (UnimplementedEnumServiceServer) TranslateValue(context.Context, *TranslateValueInput) (*TranslateValueOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TranslateValue not implemented")
}
func (UnimplementedEnumServiceServer) DeleteTranslation(context.Context, *DeleteTranslationInput) (*DeleteTranslationOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTranslation not implemented")
}
func (UnimplementedEnumServiceServer) ListTranslations(context.Context, *ListTranslationsInput) (*ListTranslationsOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListTranslations not implemented")
}

// UnsafeEnumServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EnumServiceServer will
// result in compilation errors.
type UnsafeEnumServiceServer interface {
	mustEmbedUnimplementedEnumServiceServer()
}

func RegisterEnumServiceServer(s grpc.ServiceRegistrar, srv EnumServiceServer) {
	s.RegisterService(&EnumService_ServiceDesc, srv)
}

func _EnumService_CreateEnum_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateEnumInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EnumServiceServer).CreateEnum(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EnumService_CreateEnum_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EnumServiceServer).CreateEnum(ctx, req.(*CreateEnumInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _EnumService_DeleteEnum_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteEnumInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EnumServiceServer).DeleteEnum(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EnumService_DeleteEnum_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EnumServiceServer).DeleteEnum(ctx, req.(*DeleteEnumInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _EnumService_UpdateEnum_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateEnumInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EnumServiceServer).UpdateEnum(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EnumService_UpdateEnum_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EnumServiceServer).UpdateEnum(ctx, req.(*UpdateEnumInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _EnumService_DescribeEnum_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DescribeEnumInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EnumServiceServer).DescribeEnum(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EnumService_DescribeEnum_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EnumServiceServer).DescribeEnum(ctx, req.(*DescribeEnumInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _EnumService_ListEnums_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListEnumsInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EnumServiceServer).ListEnums(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EnumService_ListEnums_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EnumServiceServer).ListEnums(ctx, req.(*ListEnumsInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _EnumService_CreateValue_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateValueInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EnumServiceServer).CreateValue(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EnumService_CreateValue_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EnumServiceServer).CreateValue(ctx, req.(*CreateValueInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _EnumService_DeleteValue_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteValueInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EnumServiceServer).DeleteValue(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EnumService_DeleteValue_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EnumServiceServer).DeleteValue(ctx, req.(*DeleteValueInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _EnumService_UpdateValue_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateValueInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EnumServiceServer).UpdateValue(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EnumService_UpdateValue_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EnumServiceServer).UpdateValue(ctx, req.(*UpdateValueInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _EnumService_DescribeValue_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DescribeValueInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EnumServiceServer).DescribeValue(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EnumService_DescribeValue_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EnumServiceServer).DescribeValue(ctx, req.(*DescribeValueInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _EnumService_ListValues_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListValuesInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EnumServiceServer).ListValues(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EnumService_ListValues_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EnumServiceServer).ListValues(ctx, req.(*ListValuesInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _EnumService_TranslateValue_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TranslateValueInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EnumServiceServer).TranslateValue(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EnumService_TranslateValue_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EnumServiceServer).TranslateValue(ctx, req.(*TranslateValueInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _EnumService_DeleteTranslation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteTranslationInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EnumServiceServer).DeleteTranslation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EnumService_DeleteTranslation_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EnumServiceServer).DeleteTranslation(ctx, req.(*DeleteTranslationInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _EnumService_ListTranslations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListTranslationsInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EnumServiceServer).ListTranslations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EnumService_ListTranslations_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EnumServiceServer).ListTranslations(ctx, req.(*ListTranslationsInput))
	}
	return interceptor(ctx, in, info, handler)
}

// EnumService_ServiceDesc is the grpc.ServiceDesc for EnumService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var EnumService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "eolymp.taxonomy.EnumService",
	HandlerType: (*EnumServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateEnum",
			Handler:    _EnumService_CreateEnum_Handler,
		},
		{
			MethodName: "DeleteEnum",
			Handler:    _EnumService_DeleteEnum_Handler,
		},
		{
			MethodName: "UpdateEnum",
			Handler:    _EnumService_UpdateEnum_Handler,
		},
		{
			MethodName: "DescribeEnum",
			Handler:    _EnumService_DescribeEnum_Handler,
		},
		{
			MethodName: "ListEnums",
			Handler:    _EnumService_ListEnums_Handler,
		},
		{
			MethodName: "CreateValue",
			Handler:    _EnumService_CreateValue_Handler,
		},
		{
			MethodName: "DeleteValue",
			Handler:    _EnumService_DeleteValue_Handler,
		},
		{
			MethodName: "UpdateValue",
			Handler:    _EnumService_UpdateValue_Handler,
		},
		{
			MethodName: "DescribeValue",
			Handler:    _EnumService_DescribeValue_Handler,
		},
		{
			MethodName: "ListValues",
			Handler:    _EnumService_ListValues_Handler,
		},
		{
			MethodName: "TranslateValue",
			Handler:    _EnumService_TranslateValue_Handler,
		},
		{
			MethodName: "DeleteTranslation",
			Handler:    _EnumService_DeleteTranslation_Handler,
		},
		{
			MethodName: "ListTranslations",
			Handler:    _EnumService_ListTranslations_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "eolymp/taxonomy/enum_service.proto",
}
