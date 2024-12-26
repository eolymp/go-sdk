// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v4.24.4
// source: eolymp/newsletter/newsletter_service.proto

package newsletter

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
	NewsletterService_DescribeNewsletter_FullMethodName            = "/eolymp.newsletter.NewsletterService/DescribeNewsletter"
	NewsletterService_ListNewsletters_FullMethodName               = "/eolymp.newsletter.NewsletterService/ListNewsletters"
	NewsletterService_CreateNewsletter_FullMethodName              = "/eolymp.newsletter.NewsletterService/CreateNewsletter"
	NewsletterService_UpdateNewsletter_FullMethodName              = "/eolymp.newsletter.NewsletterService/UpdateNewsletter"
	NewsletterService_DeleteNewsletter_FullMethodName              = "/eolymp.newsletter.NewsletterService/DeleteNewsletter"
	NewsletterService_SendNewsletter_FullMethodName                = "/eolymp.newsletter.NewsletterService/SendNewsletter"
	NewsletterService_TestNewsletter_FullMethodName                = "/eolymp.newsletter.NewsletterService/TestNewsletter"
	NewsletterService_DescribeNewsletterTranslation_FullMethodName = "/eolymp.newsletter.NewsletterService/DescribeNewsletterTranslation"
	NewsletterService_ListNewsletterTranslations_FullMethodName    = "/eolymp.newsletter.NewsletterService/ListNewsletterTranslations"
	NewsletterService_CreateNewsletterTranslation_FullMethodName   = "/eolymp.newsletter.NewsletterService/CreateNewsletterTranslation"
	NewsletterService_UpdateNewsletterTranslation_FullMethodName   = "/eolymp.newsletter.NewsletterService/UpdateNewsletterTranslation"
	NewsletterService_DeleteNewsletterTranslation_FullMethodName   = "/eolymp.newsletter.NewsletterService/DeleteNewsletterTranslation"
)

// NewsletterServiceClient is the client API for NewsletterService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type NewsletterServiceClient interface {
	DescribeNewsletter(ctx context.Context, in *DescribeNewsletterInput, opts ...grpc.CallOption) (*DescribeNewsletterOutput, error)
	ListNewsletters(ctx context.Context, in *ListNewslettersInput, opts ...grpc.CallOption) (*ListNewslettersOutput, error)
	CreateNewsletter(ctx context.Context, in *CreateNewsletterInput, opts ...grpc.CallOption) (*CreateNewsletterOutput, error)
	UpdateNewsletter(ctx context.Context, in *UpdateNewsletterInput, opts ...grpc.CallOption) (*UpdateNewsletterOutput, error)
	DeleteNewsletter(ctx context.Context, in *DeleteNewsletterInput, opts ...grpc.CallOption) (*DeleteNewsletterOutput, error)
	SendNewsletter(ctx context.Context, in *SendNewsletterInput, opts ...grpc.CallOption) (*SendNewsletterOutput, error)
	TestNewsletter(ctx context.Context, in *TestNewsletterInput, opts ...grpc.CallOption) (*TestNewsletterOutput, error)
	DescribeNewsletterTranslation(ctx context.Context, in *DescribeNewsletterTranslationInput, opts ...grpc.CallOption) (*DescribeNewsletterTranslationOutput, error)
	ListNewsletterTranslations(ctx context.Context, in *ListNewsletterTranslationsInput, opts ...grpc.CallOption) (*ListNewsletterTranslationsOutput, error)
	CreateNewsletterTranslation(ctx context.Context, in *CreateNewsletterTranslationInput, opts ...grpc.CallOption) (*CreateNewsletterTranslationOutput, error)
	UpdateNewsletterTranslation(ctx context.Context, in *UpdateNewsletterTranslationInput, opts ...grpc.CallOption) (*UpdateNewsletterTranslationOutput, error)
	DeleteNewsletterTranslation(ctx context.Context, in *DeleteNewsletterTranslationInput, opts ...grpc.CallOption) (*DeleteNewsletterTranslationOutput, error)
}

type newsletterServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewNewsletterServiceClient(cc grpc.ClientConnInterface) NewsletterServiceClient {
	return &newsletterServiceClient{cc}
}

func (c *newsletterServiceClient) DescribeNewsletter(ctx context.Context, in *DescribeNewsletterInput, opts ...grpc.CallOption) (*DescribeNewsletterOutput, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DescribeNewsletterOutput)
	err := c.cc.Invoke(ctx, NewsletterService_DescribeNewsletter_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *newsletterServiceClient) ListNewsletters(ctx context.Context, in *ListNewslettersInput, opts ...grpc.CallOption) (*ListNewslettersOutput, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListNewslettersOutput)
	err := c.cc.Invoke(ctx, NewsletterService_ListNewsletters_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *newsletterServiceClient) CreateNewsletter(ctx context.Context, in *CreateNewsletterInput, opts ...grpc.CallOption) (*CreateNewsletterOutput, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateNewsletterOutput)
	err := c.cc.Invoke(ctx, NewsletterService_CreateNewsletter_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *newsletterServiceClient) UpdateNewsletter(ctx context.Context, in *UpdateNewsletterInput, opts ...grpc.CallOption) (*UpdateNewsletterOutput, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateNewsletterOutput)
	err := c.cc.Invoke(ctx, NewsletterService_UpdateNewsletter_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *newsletterServiceClient) DeleteNewsletter(ctx context.Context, in *DeleteNewsletterInput, opts ...grpc.CallOption) (*DeleteNewsletterOutput, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteNewsletterOutput)
	err := c.cc.Invoke(ctx, NewsletterService_DeleteNewsletter_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *newsletterServiceClient) SendNewsletter(ctx context.Context, in *SendNewsletterInput, opts ...grpc.CallOption) (*SendNewsletterOutput, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SendNewsletterOutput)
	err := c.cc.Invoke(ctx, NewsletterService_SendNewsletter_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *newsletterServiceClient) TestNewsletter(ctx context.Context, in *TestNewsletterInput, opts ...grpc.CallOption) (*TestNewsletterOutput, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(TestNewsletterOutput)
	err := c.cc.Invoke(ctx, NewsletterService_TestNewsletter_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *newsletterServiceClient) DescribeNewsletterTranslation(ctx context.Context, in *DescribeNewsletterTranslationInput, opts ...grpc.CallOption) (*DescribeNewsletterTranslationOutput, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DescribeNewsletterTranslationOutput)
	err := c.cc.Invoke(ctx, NewsletterService_DescribeNewsletterTranslation_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *newsletterServiceClient) ListNewsletterTranslations(ctx context.Context, in *ListNewsletterTranslationsInput, opts ...grpc.CallOption) (*ListNewsletterTranslationsOutput, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListNewsletterTranslationsOutput)
	err := c.cc.Invoke(ctx, NewsletterService_ListNewsletterTranslations_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *newsletterServiceClient) CreateNewsletterTranslation(ctx context.Context, in *CreateNewsletterTranslationInput, opts ...grpc.CallOption) (*CreateNewsletterTranslationOutput, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateNewsletterTranslationOutput)
	err := c.cc.Invoke(ctx, NewsletterService_CreateNewsletterTranslation_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *newsletterServiceClient) UpdateNewsletterTranslation(ctx context.Context, in *UpdateNewsletterTranslationInput, opts ...grpc.CallOption) (*UpdateNewsletterTranslationOutput, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateNewsletterTranslationOutput)
	err := c.cc.Invoke(ctx, NewsletterService_UpdateNewsletterTranslation_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *newsletterServiceClient) DeleteNewsletterTranslation(ctx context.Context, in *DeleteNewsletterTranslationInput, opts ...grpc.CallOption) (*DeleteNewsletterTranslationOutput, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteNewsletterTranslationOutput)
	err := c.cc.Invoke(ctx, NewsletterService_DeleteNewsletterTranslation_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NewsletterServiceServer is the server API for NewsletterService service.
// All implementations should embed UnimplementedNewsletterServiceServer
// for forward compatibility.
type NewsletterServiceServer interface {
	DescribeNewsletter(context.Context, *DescribeNewsletterInput) (*DescribeNewsletterOutput, error)
	ListNewsletters(context.Context, *ListNewslettersInput) (*ListNewslettersOutput, error)
	CreateNewsletter(context.Context, *CreateNewsletterInput) (*CreateNewsletterOutput, error)
	UpdateNewsletter(context.Context, *UpdateNewsletterInput) (*UpdateNewsletterOutput, error)
	DeleteNewsletter(context.Context, *DeleteNewsletterInput) (*DeleteNewsletterOutput, error)
	SendNewsletter(context.Context, *SendNewsletterInput) (*SendNewsletterOutput, error)
	TestNewsletter(context.Context, *TestNewsletterInput) (*TestNewsletterOutput, error)
	DescribeNewsletterTranslation(context.Context, *DescribeNewsletterTranslationInput) (*DescribeNewsletterTranslationOutput, error)
	ListNewsletterTranslations(context.Context, *ListNewsletterTranslationsInput) (*ListNewsletterTranslationsOutput, error)
	CreateNewsletterTranslation(context.Context, *CreateNewsletterTranslationInput) (*CreateNewsletterTranslationOutput, error)
	UpdateNewsletterTranslation(context.Context, *UpdateNewsletterTranslationInput) (*UpdateNewsletterTranslationOutput, error)
	DeleteNewsletterTranslation(context.Context, *DeleteNewsletterTranslationInput) (*DeleteNewsletterTranslationOutput, error)
}

// UnimplementedNewsletterServiceServer should be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedNewsletterServiceServer struct{}

func (UnimplementedNewsletterServiceServer) DescribeNewsletter(context.Context, *DescribeNewsletterInput) (*DescribeNewsletterOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DescribeNewsletter not implemented")
}
func (UnimplementedNewsletterServiceServer) ListNewsletters(context.Context, *ListNewslettersInput) (*ListNewslettersOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListNewsletters not implemented")
}
func (UnimplementedNewsletterServiceServer) CreateNewsletter(context.Context, *CreateNewsletterInput) (*CreateNewsletterOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateNewsletter not implemented")
}
func (UnimplementedNewsletterServiceServer) UpdateNewsletter(context.Context, *UpdateNewsletterInput) (*UpdateNewsletterOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateNewsletter not implemented")
}
func (UnimplementedNewsletterServiceServer) DeleteNewsletter(context.Context, *DeleteNewsletterInput) (*DeleteNewsletterOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteNewsletter not implemented")
}
func (UnimplementedNewsletterServiceServer) SendNewsletter(context.Context, *SendNewsletterInput) (*SendNewsletterOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendNewsletter not implemented")
}
func (UnimplementedNewsletterServiceServer) TestNewsletter(context.Context, *TestNewsletterInput) (*TestNewsletterOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TestNewsletter not implemented")
}
func (UnimplementedNewsletterServiceServer) DescribeNewsletterTranslation(context.Context, *DescribeNewsletterTranslationInput) (*DescribeNewsletterTranslationOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DescribeNewsletterTranslation not implemented")
}
func (UnimplementedNewsletterServiceServer) ListNewsletterTranslations(context.Context, *ListNewsletterTranslationsInput) (*ListNewsletterTranslationsOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListNewsletterTranslations not implemented")
}
func (UnimplementedNewsletterServiceServer) CreateNewsletterTranslation(context.Context, *CreateNewsletterTranslationInput) (*CreateNewsletterTranslationOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateNewsletterTranslation not implemented")
}
func (UnimplementedNewsletterServiceServer) UpdateNewsletterTranslation(context.Context, *UpdateNewsletterTranslationInput) (*UpdateNewsletterTranslationOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateNewsletterTranslation not implemented")
}
func (UnimplementedNewsletterServiceServer) DeleteNewsletterTranslation(context.Context, *DeleteNewsletterTranslationInput) (*DeleteNewsletterTranslationOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteNewsletterTranslation not implemented")
}
func (UnimplementedNewsletterServiceServer) testEmbeddedByValue() {}

// UnsafeNewsletterServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NewsletterServiceServer will
// result in compilation errors.
type UnsafeNewsletterServiceServer interface {
	mustEmbedUnimplementedNewsletterServiceServer()
}

func RegisterNewsletterServiceServer(s grpc.ServiceRegistrar, srv NewsletterServiceServer) {
	// If the following call pancis, it indicates UnimplementedNewsletterServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&NewsletterService_ServiceDesc, srv)
}

func _NewsletterService_DescribeNewsletter_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DescribeNewsletterInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NewsletterServiceServer).DescribeNewsletter(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NewsletterService_DescribeNewsletter_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NewsletterServiceServer).DescribeNewsletter(ctx, req.(*DescribeNewsletterInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _NewsletterService_ListNewsletters_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListNewslettersInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NewsletterServiceServer).ListNewsletters(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NewsletterService_ListNewsletters_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NewsletterServiceServer).ListNewsletters(ctx, req.(*ListNewslettersInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _NewsletterService_CreateNewsletter_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateNewsletterInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NewsletterServiceServer).CreateNewsletter(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NewsletterService_CreateNewsletter_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NewsletterServiceServer).CreateNewsletter(ctx, req.(*CreateNewsletterInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _NewsletterService_UpdateNewsletter_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateNewsletterInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NewsletterServiceServer).UpdateNewsletter(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NewsletterService_UpdateNewsletter_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NewsletterServiceServer).UpdateNewsletter(ctx, req.(*UpdateNewsletterInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _NewsletterService_DeleteNewsletter_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteNewsletterInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NewsletterServiceServer).DeleteNewsletter(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NewsletterService_DeleteNewsletter_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NewsletterServiceServer).DeleteNewsletter(ctx, req.(*DeleteNewsletterInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _NewsletterService_SendNewsletter_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendNewsletterInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NewsletterServiceServer).SendNewsletter(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NewsletterService_SendNewsletter_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NewsletterServiceServer).SendNewsletter(ctx, req.(*SendNewsletterInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _NewsletterService_TestNewsletter_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TestNewsletterInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NewsletterServiceServer).TestNewsletter(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NewsletterService_TestNewsletter_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NewsletterServiceServer).TestNewsletter(ctx, req.(*TestNewsletterInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _NewsletterService_DescribeNewsletterTranslation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DescribeNewsletterTranslationInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NewsletterServiceServer).DescribeNewsletterTranslation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NewsletterService_DescribeNewsletterTranslation_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NewsletterServiceServer).DescribeNewsletterTranslation(ctx, req.(*DescribeNewsletterTranslationInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _NewsletterService_ListNewsletterTranslations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListNewsletterTranslationsInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NewsletterServiceServer).ListNewsletterTranslations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NewsletterService_ListNewsletterTranslations_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NewsletterServiceServer).ListNewsletterTranslations(ctx, req.(*ListNewsletterTranslationsInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _NewsletterService_CreateNewsletterTranslation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateNewsletterTranslationInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NewsletterServiceServer).CreateNewsletterTranslation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NewsletterService_CreateNewsletterTranslation_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NewsletterServiceServer).CreateNewsletterTranslation(ctx, req.(*CreateNewsletterTranslationInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _NewsletterService_UpdateNewsletterTranslation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateNewsletterTranslationInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NewsletterServiceServer).UpdateNewsletterTranslation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NewsletterService_UpdateNewsletterTranslation_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NewsletterServiceServer).UpdateNewsletterTranslation(ctx, req.(*UpdateNewsletterTranslationInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _NewsletterService_DeleteNewsletterTranslation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteNewsletterTranslationInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NewsletterServiceServer).DeleteNewsletterTranslation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NewsletterService_DeleteNewsletterTranslation_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NewsletterServiceServer).DeleteNewsletterTranslation(ctx, req.(*DeleteNewsletterTranslationInput))
	}
	return interceptor(ctx, in, info, handler)
}

// NewsletterService_ServiceDesc is the grpc.ServiceDesc for NewsletterService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var NewsletterService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "eolymp.newsletter.NewsletterService",
	HandlerType: (*NewsletterServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DescribeNewsletter",
			Handler:    _NewsletterService_DescribeNewsletter_Handler,
		},
		{
			MethodName: "ListNewsletters",
			Handler:    _NewsletterService_ListNewsletters_Handler,
		},
		{
			MethodName: "CreateNewsletter",
			Handler:    _NewsletterService_CreateNewsletter_Handler,
		},
		{
			MethodName: "UpdateNewsletter",
			Handler:    _NewsletterService_UpdateNewsletter_Handler,
		},
		{
			MethodName: "DeleteNewsletter",
			Handler:    _NewsletterService_DeleteNewsletter_Handler,
		},
		{
			MethodName: "SendNewsletter",
			Handler:    _NewsletterService_SendNewsletter_Handler,
		},
		{
			MethodName: "TestNewsletter",
			Handler:    _NewsletterService_TestNewsletter_Handler,
		},
		{
			MethodName: "DescribeNewsletterTranslation",
			Handler:    _NewsletterService_DescribeNewsletterTranslation_Handler,
		},
		{
			MethodName: "ListNewsletterTranslations",
			Handler:    _NewsletterService_ListNewsletterTranslations_Handler,
		},
		{
			MethodName: "CreateNewsletterTranslation",
			Handler:    _NewsletterService_CreateNewsletterTranslation_Handler,
		},
		{
			MethodName: "UpdateNewsletterTranslation",
			Handler:    _NewsletterService_UpdateNewsletterTranslation_Handler,
		},
		{
			MethodName: "DeleteNewsletterTranslation",
			Handler:    _NewsletterService_DeleteNewsletterTranslation_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "eolymp/newsletter/newsletter_service.proto",
}
