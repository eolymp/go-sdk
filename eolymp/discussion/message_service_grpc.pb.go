// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v4.24.4
// source: eolymp/discussion/message_service.proto

package discussion

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
	MessageService_DescribeMessage_FullMethodName    = "/eolymp.discussion.MessageService/DescribeMessage"
	MessageService_ListMessages_FullMethodName       = "/eolymp.discussion.MessageService/ListMessages"
	MessageService_PostMessage_FullMethodName        = "/eolymp.discussion.MessageService/PostMessage"
	MessageService_UpdateMessage_FullMethodName      = "/eolymp.discussion.MessageService/UpdateMessage"
	MessageService_DeleteMessage_FullMethodName      = "/eolymp.discussion.MessageService/DeleteMessage"
	MessageService_VoteMessage_FullMethodName        = "/eolymp.discussion.MessageService/VoteMessage"
	MessageService_ListMessageChanges_FullMethodName = "/eolymp.discussion.MessageService/ListMessageChanges"
)

// MessageServiceClient is the client API for MessageService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MessageServiceClient interface {
	DescribeMessage(ctx context.Context, in *DescribeMessageInput, opts ...grpc.CallOption) (*DescribeMessageOutput, error)
	ListMessages(ctx context.Context, in *ListMessagesInput, opts ...grpc.CallOption) (*ListMessagesOutput, error)
	PostMessage(ctx context.Context, in *PostMessageInput, opts ...grpc.CallOption) (*PostMessageOutput, error)
	UpdateMessage(ctx context.Context, in *UpdateMessageInput, opts ...grpc.CallOption) (*UpdateMessageOutput, error)
	DeleteMessage(ctx context.Context, in *DeleteMessageInput, opts ...grpc.CallOption) (*DeleteMessageOutput, error)
	VoteMessage(ctx context.Context, in *VoteMessageInput, opts ...grpc.CallOption) (*VoteMessageOutput, error)
	ListMessageChanges(ctx context.Context, in *ListMessageChangesInput, opts ...grpc.CallOption) (*ListMessageChangesOutput, error)
}

type messageServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMessageServiceClient(cc grpc.ClientConnInterface) MessageServiceClient {
	return &messageServiceClient{cc}
}

func (c *messageServiceClient) DescribeMessage(ctx context.Context, in *DescribeMessageInput, opts ...grpc.CallOption) (*DescribeMessageOutput, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DescribeMessageOutput)
	err := c.cc.Invoke(ctx, MessageService_DescribeMessage_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageServiceClient) ListMessages(ctx context.Context, in *ListMessagesInput, opts ...grpc.CallOption) (*ListMessagesOutput, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListMessagesOutput)
	err := c.cc.Invoke(ctx, MessageService_ListMessages_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageServiceClient) PostMessage(ctx context.Context, in *PostMessageInput, opts ...grpc.CallOption) (*PostMessageOutput, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PostMessageOutput)
	err := c.cc.Invoke(ctx, MessageService_PostMessage_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageServiceClient) UpdateMessage(ctx context.Context, in *UpdateMessageInput, opts ...grpc.CallOption) (*UpdateMessageOutput, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateMessageOutput)
	err := c.cc.Invoke(ctx, MessageService_UpdateMessage_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageServiceClient) DeleteMessage(ctx context.Context, in *DeleteMessageInput, opts ...grpc.CallOption) (*DeleteMessageOutput, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteMessageOutput)
	err := c.cc.Invoke(ctx, MessageService_DeleteMessage_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageServiceClient) VoteMessage(ctx context.Context, in *VoteMessageInput, opts ...grpc.CallOption) (*VoteMessageOutput, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(VoteMessageOutput)
	err := c.cc.Invoke(ctx, MessageService_VoteMessage_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageServiceClient) ListMessageChanges(ctx context.Context, in *ListMessageChangesInput, opts ...grpc.CallOption) (*ListMessageChangesOutput, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListMessageChangesOutput)
	err := c.cc.Invoke(ctx, MessageService_ListMessageChanges_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MessageServiceServer is the server API for MessageService service.
// All implementations should embed UnimplementedMessageServiceServer
// for forward compatibility.
type MessageServiceServer interface {
	DescribeMessage(context.Context, *DescribeMessageInput) (*DescribeMessageOutput, error)
	ListMessages(context.Context, *ListMessagesInput) (*ListMessagesOutput, error)
	PostMessage(context.Context, *PostMessageInput) (*PostMessageOutput, error)
	UpdateMessage(context.Context, *UpdateMessageInput) (*UpdateMessageOutput, error)
	DeleteMessage(context.Context, *DeleteMessageInput) (*DeleteMessageOutput, error)
	VoteMessage(context.Context, *VoteMessageInput) (*VoteMessageOutput, error)
	ListMessageChanges(context.Context, *ListMessageChangesInput) (*ListMessageChangesOutput, error)
}

// UnimplementedMessageServiceServer should be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedMessageServiceServer struct{}

func (UnimplementedMessageServiceServer) DescribeMessage(context.Context, *DescribeMessageInput) (*DescribeMessageOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DescribeMessage not implemented")
}
func (UnimplementedMessageServiceServer) ListMessages(context.Context, *ListMessagesInput) (*ListMessagesOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListMessages not implemented")
}
func (UnimplementedMessageServiceServer) PostMessage(context.Context, *PostMessageInput) (*PostMessageOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PostMessage not implemented")
}
func (UnimplementedMessageServiceServer) UpdateMessage(context.Context, *UpdateMessageInput) (*UpdateMessageOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateMessage not implemented")
}
func (UnimplementedMessageServiceServer) DeleteMessage(context.Context, *DeleteMessageInput) (*DeleteMessageOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteMessage not implemented")
}
func (UnimplementedMessageServiceServer) VoteMessage(context.Context, *VoteMessageInput) (*VoteMessageOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VoteMessage not implemented")
}
func (UnimplementedMessageServiceServer) ListMessageChanges(context.Context, *ListMessageChangesInput) (*ListMessageChangesOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListMessageChanges not implemented")
}
func (UnimplementedMessageServiceServer) testEmbeddedByValue() {}

// UnsafeMessageServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MessageServiceServer will
// result in compilation errors.
type UnsafeMessageServiceServer interface {
	mustEmbedUnimplementedMessageServiceServer()
}

func RegisterMessageServiceServer(s grpc.ServiceRegistrar, srv MessageServiceServer) {
	// If the following call pancis, it indicates UnimplementedMessageServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&MessageService_ServiceDesc, srv)
}

func _MessageService_DescribeMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DescribeMessageInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageServiceServer).DescribeMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MessageService_DescribeMessage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageServiceServer).DescribeMessage(ctx, req.(*DescribeMessageInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _MessageService_ListMessages_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListMessagesInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageServiceServer).ListMessages(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MessageService_ListMessages_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageServiceServer).ListMessages(ctx, req.(*ListMessagesInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _MessageService_PostMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PostMessageInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageServiceServer).PostMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MessageService_PostMessage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageServiceServer).PostMessage(ctx, req.(*PostMessageInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _MessageService_UpdateMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateMessageInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageServiceServer).UpdateMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MessageService_UpdateMessage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageServiceServer).UpdateMessage(ctx, req.(*UpdateMessageInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _MessageService_DeleteMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteMessageInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageServiceServer).DeleteMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MessageService_DeleteMessage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageServiceServer).DeleteMessage(ctx, req.(*DeleteMessageInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _MessageService_VoteMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VoteMessageInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageServiceServer).VoteMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MessageService_VoteMessage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageServiceServer).VoteMessage(ctx, req.(*VoteMessageInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _MessageService_ListMessageChanges_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListMessageChangesInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageServiceServer).ListMessageChanges(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MessageService_ListMessageChanges_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageServiceServer).ListMessageChanges(ctx, req.(*ListMessageChangesInput))
	}
	return interceptor(ctx, in, info, handler)
}

// MessageService_ServiceDesc is the grpc.ServiceDesc for MessageService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MessageService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "eolymp.discussion.MessageService",
	HandlerType: (*MessageServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DescribeMessage",
			Handler:    _MessageService_DescribeMessage_Handler,
		},
		{
			MethodName: "ListMessages",
			Handler:    _MessageService_ListMessages_Handler,
		},
		{
			MethodName: "PostMessage",
			Handler:    _MessageService_PostMessage_Handler,
		},
		{
			MethodName: "UpdateMessage",
			Handler:    _MessageService_UpdateMessage_Handler,
		},
		{
			MethodName: "DeleteMessage",
			Handler:    _MessageService_DeleteMessage_Handler,
		},
		{
			MethodName: "VoteMessage",
			Handler:    _MessageService_VoteMessage_Handler,
		},
		{
			MethodName: "ListMessageChanges",
			Handler:    _MessageService_ListMessageChanges_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "eolymp/discussion/message_service.proto",
}
