// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v4.24.4
// source: eolymp/notify/notification_service.proto

package notify

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
	NotificationService_CreateNotification_FullMethodName   = "/eolymp.notify.NotificationService/CreateNotification"
	NotificationService_DescribeNotification_FullMethodName = "/eolymp.notify.NotificationService/DescribeNotification"
	NotificationService_ReadNotification_FullMethodName     = "/eolymp.notify.NotificationService/ReadNotification"
	NotificationService_DeleteNotification_FullMethodName   = "/eolymp.notify.NotificationService/DeleteNotification"
	NotificationService_ListNotifications_FullMethodName    = "/eolymp.notify.NotificationService/ListNotifications"
	NotificationService_DescribePreferences_FullMethodName  = "/eolymp.notify.NotificationService/DescribePreferences"
	NotificationService_UpdatePreferences_FullMethodName    = "/eolymp.notify.NotificationService/UpdatePreferences"
)

// NotificationServiceClient is the client API for NotificationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type NotificationServiceClient interface {
	CreateNotification(ctx context.Context, in *CreateNotificationInput, opts ...grpc.CallOption) (*CreateNotificationOutput, error)
	DescribeNotification(ctx context.Context, in *DescribeNotificationInput, opts ...grpc.CallOption) (*DescribeNotificationOutput, error)
	ReadNotification(ctx context.Context, in *ReadNotificationInput, opts ...grpc.CallOption) (*ReadNotificationOutput, error)
	DeleteNotification(ctx context.Context, in *DeleteNotificationInput, opts ...grpc.CallOption) (*DeleteNotificationOutput, error)
	ListNotifications(ctx context.Context, in *ListNotificationsInput, opts ...grpc.CallOption) (*ListNotificationsOutput, error)
	DescribePreferences(ctx context.Context, in *DescribePreferencesInput, opts ...grpc.CallOption) (*DescribePreferencesOutput, error)
	UpdatePreferences(ctx context.Context, in *UpdatePreferencesInput, opts ...grpc.CallOption) (*UpdatePreferencesOutput, error)
}

type notificationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewNotificationServiceClient(cc grpc.ClientConnInterface) NotificationServiceClient {
	return &notificationServiceClient{cc}
}

func (c *notificationServiceClient) CreateNotification(ctx context.Context, in *CreateNotificationInput, opts ...grpc.CallOption) (*CreateNotificationOutput, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateNotificationOutput)
	err := c.cc.Invoke(ctx, NotificationService_CreateNotification_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notificationServiceClient) DescribeNotification(ctx context.Context, in *DescribeNotificationInput, opts ...grpc.CallOption) (*DescribeNotificationOutput, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DescribeNotificationOutput)
	err := c.cc.Invoke(ctx, NotificationService_DescribeNotification_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notificationServiceClient) ReadNotification(ctx context.Context, in *ReadNotificationInput, opts ...grpc.CallOption) (*ReadNotificationOutput, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ReadNotificationOutput)
	err := c.cc.Invoke(ctx, NotificationService_ReadNotification_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notificationServiceClient) DeleteNotification(ctx context.Context, in *DeleteNotificationInput, opts ...grpc.CallOption) (*DeleteNotificationOutput, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteNotificationOutput)
	err := c.cc.Invoke(ctx, NotificationService_DeleteNotification_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notificationServiceClient) ListNotifications(ctx context.Context, in *ListNotificationsInput, opts ...grpc.CallOption) (*ListNotificationsOutput, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListNotificationsOutput)
	err := c.cc.Invoke(ctx, NotificationService_ListNotifications_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notificationServiceClient) DescribePreferences(ctx context.Context, in *DescribePreferencesInput, opts ...grpc.CallOption) (*DescribePreferencesOutput, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DescribePreferencesOutput)
	err := c.cc.Invoke(ctx, NotificationService_DescribePreferences_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notificationServiceClient) UpdatePreferences(ctx context.Context, in *UpdatePreferencesInput, opts ...grpc.CallOption) (*UpdatePreferencesOutput, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdatePreferencesOutput)
	err := c.cc.Invoke(ctx, NotificationService_UpdatePreferences_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NotificationServiceServer is the server API for NotificationService service.
// All implementations should embed UnimplementedNotificationServiceServer
// for forward compatibility.
type NotificationServiceServer interface {
	CreateNotification(context.Context, *CreateNotificationInput) (*CreateNotificationOutput, error)
	DescribeNotification(context.Context, *DescribeNotificationInput) (*DescribeNotificationOutput, error)
	ReadNotification(context.Context, *ReadNotificationInput) (*ReadNotificationOutput, error)
	DeleteNotification(context.Context, *DeleteNotificationInput) (*DeleteNotificationOutput, error)
	ListNotifications(context.Context, *ListNotificationsInput) (*ListNotificationsOutput, error)
	DescribePreferences(context.Context, *DescribePreferencesInput) (*DescribePreferencesOutput, error)
	UpdatePreferences(context.Context, *UpdatePreferencesInput) (*UpdatePreferencesOutput, error)
}

// UnimplementedNotificationServiceServer should be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedNotificationServiceServer struct{}

func (UnimplementedNotificationServiceServer) CreateNotification(context.Context, *CreateNotificationInput) (*CreateNotificationOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateNotification not implemented")
}
func (UnimplementedNotificationServiceServer) DescribeNotification(context.Context, *DescribeNotificationInput) (*DescribeNotificationOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DescribeNotification not implemented")
}
func (UnimplementedNotificationServiceServer) ReadNotification(context.Context, *ReadNotificationInput) (*ReadNotificationOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadNotification not implemented")
}
func (UnimplementedNotificationServiceServer) DeleteNotification(context.Context, *DeleteNotificationInput) (*DeleteNotificationOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteNotification not implemented")
}
func (UnimplementedNotificationServiceServer) ListNotifications(context.Context, *ListNotificationsInput) (*ListNotificationsOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListNotifications not implemented")
}
func (UnimplementedNotificationServiceServer) DescribePreferences(context.Context, *DescribePreferencesInput) (*DescribePreferencesOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DescribePreferences not implemented")
}
func (UnimplementedNotificationServiceServer) UpdatePreferences(context.Context, *UpdatePreferencesInput) (*UpdatePreferencesOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePreferences not implemented")
}
func (UnimplementedNotificationServiceServer) testEmbeddedByValue() {}

// UnsafeNotificationServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NotificationServiceServer will
// result in compilation errors.
type UnsafeNotificationServiceServer interface {
	mustEmbedUnimplementedNotificationServiceServer()
}

func RegisterNotificationServiceServer(s grpc.ServiceRegistrar, srv NotificationServiceServer) {
	// If the following call pancis, it indicates UnimplementedNotificationServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&NotificationService_ServiceDesc, srv)
}

func _NotificationService_CreateNotification_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateNotificationInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotificationServiceServer).CreateNotification(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NotificationService_CreateNotification_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotificationServiceServer).CreateNotification(ctx, req.(*CreateNotificationInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _NotificationService_DescribeNotification_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DescribeNotificationInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotificationServiceServer).DescribeNotification(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NotificationService_DescribeNotification_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotificationServiceServer).DescribeNotification(ctx, req.(*DescribeNotificationInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _NotificationService_ReadNotification_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadNotificationInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotificationServiceServer).ReadNotification(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NotificationService_ReadNotification_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotificationServiceServer).ReadNotification(ctx, req.(*ReadNotificationInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _NotificationService_DeleteNotification_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteNotificationInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotificationServiceServer).DeleteNotification(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NotificationService_DeleteNotification_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotificationServiceServer).DeleteNotification(ctx, req.(*DeleteNotificationInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _NotificationService_ListNotifications_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListNotificationsInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotificationServiceServer).ListNotifications(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NotificationService_ListNotifications_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotificationServiceServer).ListNotifications(ctx, req.(*ListNotificationsInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _NotificationService_DescribePreferences_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DescribePreferencesInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotificationServiceServer).DescribePreferences(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NotificationService_DescribePreferences_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotificationServiceServer).DescribePreferences(ctx, req.(*DescribePreferencesInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _NotificationService_UpdatePreferences_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePreferencesInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotificationServiceServer).UpdatePreferences(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NotificationService_UpdatePreferences_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotificationServiceServer).UpdatePreferences(ctx, req.(*UpdatePreferencesInput))
	}
	return interceptor(ctx, in, info, handler)
}

// NotificationService_ServiceDesc is the grpc.ServiceDesc for NotificationService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var NotificationService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "eolymp.notify.NotificationService",
	HandlerType: (*NotificationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateNotification",
			Handler:    _NotificationService_CreateNotification_Handler,
		},
		{
			MethodName: "DescribeNotification",
			Handler:    _NotificationService_DescribeNotification_Handler,
		},
		{
			MethodName: "ReadNotification",
			Handler:    _NotificationService_ReadNotification_Handler,
		},
		{
			MethodName: "DeleteNotification",
			Handler:    _NotificationService_DeleteNotification_Handler,
		},
		{
			MethodName: "ListNotifications",
			Handler:    _NotificationService_ListNotifications_Handler,
		},
		{
			MethodName: "DescribePreferences",
			Handler:    _NotificationService_DescribePreferences_Handler,
		},
		{
			MethodName: "UpdatePreferences",
			Handler:    _NotificationService_UpdatePreferences_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "eolymp/notify/notification_service.proto",
}
