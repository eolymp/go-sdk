// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.18.1
// source: eolymp/atlas/taxonomy_service.proto

package atlas

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
	TopicService_CreateTopic_FullMethodName       = "/eolymp.atlas.TopicService/CreateTopic"
	TopicService_DeleteTopic_FullMethodName       = "/eolymp.atlas.TopicService/DeleteTopic"
	TopicService_UpdateTopic_FullMethodName       = "/eolymp.atlas.TopicService/UpdateTopic"
	TopicService_DescribeTopic_FullMethodName     = "/eolymp.atlas.TopicService/DescribeTopic"
	TopicService_ListTopics_FullMethodName        = "/eolymp.atlas.TopicService/ListTopics"
	TopicService_TranslateTopic_FullMethodName    = "/eolymp.atlas.TopicService/TranslateTopic"
	TopicService_DeleteTranslation_FullMethodName = "/eolymp.atlas.TopicService/DeleteTranslation"
	TopicService_ListTranslations_FullMethodName  = "/eolymp.atlas.TopicService/ListTranslations"
)

// TopicServiceClient is the client API for TopicService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TopicServiceClient interface {
	CreateTopic(ctx context.Context, in *CreateTopicInput, opts ...grpc.CallOption) (*CreateTopicOutput, error)
	DeleteTopic(ctx context.Context, in *DeleteTopicInput, opts ...grpc.CallOption) (*DeleteTopicOutput, error)
	UpdateTopic(ctx context.Context, in *UpdateTopicInput, opts ...grpc.CallOption) (*UpdateTopicOutput, error)
	DescribeTopic(ctx context.Context, in *DescribeTopicInput, opts ...grpc.CallOption) (*DescribeTopicOutput, error)
	ListTopics(ctx context.Context, in *ListTopicsInput, opts ...grpc.CallOption) (*ListTopicsOutput, error)
	TranslateTopic(ctx context.Context, in *TranslateTopicInput, opts ...grpc.CallOption) (*TranslateTopicOutput, error)
	DeleteTranslation(ctx context.Context, in *DeleteTranslationInput, opts ...grpc.CallOption) (*DeleteTranslationOutput, error)
	ListTranslations(ctx context.Context, in *ListTranslationsInput, opts ...grpc.CallOption) (*ListTranslationsOutput, error)
}

type topicServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTopicServiceClient(cc grpc.ClientConnInterface) TopicServiceClient {
	return &topicServiceClient{cc}
}

func (c *topicServiceClient) CreateTopic(ctx context.Context, in *CreateTopicInput, opts ...grpc.CallOption) (*CreateTopicOutput, error) {
	out := new(CreateTopicOutput)
	err := c.cc.Invoke(ctx, TopicService_CreateTopic_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *topicServiceClient) DeleteTopic(ctx context.Context, in *DeleteTopicInput, opts ...grpc.CallOption) (*DeleteTopicOutput, error) {
	out := new(DeleteTopicOutput)
	err := c.cc.Invoke(ctx, TopicService_DeleteTopic_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *topicServiceClient) UpdateTopic(ctx context.Context, in *UpdateTopicInput, opts ...grpc.CallOption) (*UpdateTopicOutput, error) {
	out := new(UpdateTopicOutput)
	err := c.cc.Invoke(ctx, TopicService_UpdateTopic_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *topicServiceClient) DescribeTopic(ctx context.Context, in *DescribeTopicInput, opts ...grpc.CallOption) (*DescribeTopicOutput, error) {
	out := new(DescribeTopicOutput)
	err := c.cc.Invoke(ctx, TopicService_DescribeTopic_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *topicServiceClient) ListTopics(ctx context.Context, in *ListTopicsInput, opts ...grpc.CallOption) (*ListTopicsOutput, error) {
	out := new(ListTopicsOutput)
	err := c.cc.Invoke(ctx, TopicService_ListTopics_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *topicServiceClient) TranslateTopic(ctx context.Context, in *TranslateTopicInput, opts ...grpc.CallOption) (*TranslateTopicOutput, error) {
	out := new(TranslateTopicOutput)
	err := c.cc.Invoke(ctx, TopicService_TranslateTopic_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *topicServiceClient) DeleteTranslation(ctx context.Context, in *DeleteTranslationInput, opts ...grpc.CallOption) (*DeleteTranslationOutput, error) {
	out := new(DeleteTranslationOutput)
	err := c.cc.Invoke(ctx, TopicService_DeleteTranslation_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *topicServiceClient) ListTranslations(ctx context.Context, in *ListTranslationsInput, opts ...grpc.CallOption) (*ListTranslationsOutput, error) {
	out := new(ListTranslationsOutput)
	err := c.cc.Invoke(ctx, TopicService_ListTranslations_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TopicServiceServer is the server API for TopicService service.
// All implementations should embed UnimplementedTopicServiceServer
// for forward compatibility
type TopicServiceServer interface {
	CreateTopic(context.Context, *CreateTopicInput) (*CreateTopicOutput, error)
	DeleteTopic(context.Context, *DeleteTopicInput) (*DeleteTopicOutput, error)
	UpdateTopic(context.Context, *UpdateTopicInput) (*UpdateTopicOutput, error)
	DescribeTopic(context.Context, *DescribeTopicInput) (*DescribeTopicOutput, error)
	ListTopics(context.Context, *ListTopicsInput) (*ListTopicsOutput, error)
	TranslateTopic(context.Context, *TranslateTopicInput) (*TranslateTopicOutput, error)
	DeleteTranslation(context.Context, *DeleteTranslationInput) (*DeleteTranslationOutput, error)
	ListTranslations(context.Context, *ListTranslationsInput) (*ListTranslationsOutput, error)
}

// UnimplementedTopicServiceServer should be embedded to have forward compatible implementations.
type UnimplementedTopicServiceServer struct {
}

func (UnimplementedTopicServiceServer) CreateTopic(context.Context, *CreateTopicInput) (*CreateTopicOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTopic not implemented")
}
func (UnimplementedTopicServiceServer) DeleteTopic(context.Context, *DeleteTopicInput) (*DeleteTopicOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTopic not implemented")
}
func (UnimplementedTopicServiceServer) UpdateTopic(context.Context, *UpdateTopicInput) (*UpdateTopicOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateTopic not implemented")
}
func (UnimplementedTopicServiceServer) DescribeTopic(context.Context, *DescribeTopicInput) (*DescribeTopicOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DescribeTopic not implemented")
}
func (UnimplementedTopicServiceServer) ListTopics(context.Context, *ListTopicsInput) (*ListTopicsOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListTopics not implemented")
}
func (UnimplementedTopicServiceServer) TranslateTopic(context.Context, *TranslateTopicInput) (*TranslateTopicOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TranslateTopic not implemented")
}
func (UnimplementedTopicServiceServer) DeleteTranslation(context.Context, *DeleteTranslationInput) (*DeleteTranslationOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTranslation not implemented")
}
func (UnimplementedTopicServiceServer) ListTranslations(context.Context, *ListTranslationsInput) (*ListTranslationsOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListTranslations not implemented")
}

// UnsafeTopicServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TopicServiceServer will
// result in compilation errors.
type UnsafeTopicServiceServer interface {
	mustEmbedUnimplementedTopicServiceServer()
}

func RegisterTopicServiceServer(s grpc.ServiceRegistrar, srv TopicServiceServer) {
	s.RegisterService(&TopicService_ServiceDesc, srv)
}

func _TopicService_CreateTopic_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTopicInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TopicServiceServer).CreateTopic(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TopicService_CreateTopic_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TopicServiceServer).CreateTopic(ctx, req.(*CreateTopicInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _TopicService_DeleteTopic_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteTopicInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TopicServiceServer).DeleteTopic(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TopicService_DeleteTopic_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TopicServiceServer).DeleteTopic(ctx, req.(*DeleteTopicInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _TopicService_UpdateTopic_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateTopicInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TopicServiceServer).UpdateTopic(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TopicService_UpdateTopic_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TopicServiceServer).UpdateTopic(ctx, req.(*UpdateTopicInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _TopicService_DescribeTopic_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DescribeTopicInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TopicServiceServer).DescribeTopic(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TopicService_DescribeTopic_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TopicServiceServer).DescribeTopic(ctx, req.(*DescribeTopicInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _TopicService_ListTopics_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListTopicsInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TopicServiceServer).ListTopics(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TopicService_ListTopics_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TopicServiceServer).ListTopics(ctx, req.(*ListTopicsInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _TopicService_TranslateTopic_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TranslateTopicInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TopicServiceServer).TranslateTopic(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TopicService_TranslateTopic_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TopicServiceServer).TranslateTopic(ctx, req.(*TranslateTopicInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _TopicService_DeleteTranslation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteTranslationInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TopicServiceServer).DeleteTranslation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TopicService_DeleteTranslation_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TopicServiceServer).DeleteTranslation(ctx, req.(*DeleteTranslationInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _TopicService_ListTranslations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListTranslationsInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TopicServiceServer).ListTranslations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TopicService_ListTranslations_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TopicServiceServer).ListTranslations(ctx, req.(*ListTranslationsInput))
	}
	return interceptor(ctx, in, info, handler)
}

// TopicService_ServiceDesc is the grpc.ServiceDesc for TopicService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TopicService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "eolymp.atlas.TopicService",
	HandlerType: (*TopicServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateTopic",
			Handler:    _TopicService_CreateTopic_Handler,
		},
		{
			MethodName: "DeleteTopic",
			Handler:    _TopicService_DeleteTopic_Handler,
		},
		{
			MethodName: "UpdateTopic",
			Handler:    _TopicService_UpdateTopic_Handler,
		},
		{
			MethodName: "DescribeTopic",
			Handler:    _TopicService_DescribeTopic_Handler,
		},
		{
			MethodName: "ListTopics",
			Handler:    _TopicService_ListTopics_Handler,
		},
		{
			MethodName: "TranslateTopic",
			Handler:    _TopicService_TranslateTopic_Handler,
		},
		{
			MethodName: "DeleteTranslation",
			Handler:    _TopicService_DeleteTranslation_Handler,
		},
		{
			MethodName: "ListTranslations",
			Handler:    _TopicService_ListTranslations_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "eolymp/atlas/taxonomy_service.proto",
}
