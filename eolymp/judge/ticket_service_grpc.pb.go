// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.21.12
// source: eolymp/judge/ticket_service.proto

package judge

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
	TicketService_CreateTicket_FullMethodName       = "/eolymp.judge.TicketService/CreateTicket"
	TicketService_UpdateTicket_FullMethodName       = "/eolymp.judge.TicketService/UpdateTicket"
	TicketService_ReadTicket_FullMethodName         = "/eolymp.judge.TicketService/ReadTicket"
	TicketService_DeleteTicket_FullMethodName       = "/eolymp.judge.TicketService/DeleteTicket"
	TicketService_DescribeTicket_FullMethodName     = "/eolymp.judge.TicketService/DescribeTicket"
	TicketService_ListTickets_FullMethodName        = "/eolymp.judge.TicketService/ListTickets"
	TicketService_ReplyTicket_FullMethodName        = "/eolymp.judge.TicketService/ReplyTicket"
	TicketService_WatchTicket_FullMethodName        = "/eolymp.judge.TicketService/WatchTicket"
	TicketService_WatchTickets_FullMethodName       = "/eolymp.judge.TicketService/WatchTickets"
	TicketService_WatchTicketSummary_FullMethodName = "/eolymp.judge.TicketService/WatchTicketSummary"
	TicketService_ListReplies_FullMethodName        = "/eolymp.judge.TicketService/ListReplies"
	TicketService_DeleteReply_FullMethodName        = "/eolymp.judge.TicketService/DeleteReply"
	TicketService_UpdateReply_FullMethodName        = "/eolymp.judge.TicketService/UpdateReply"
	TicketService_WatchReplies_FullMethodName       = "/eolymp.judge.TicketService/WatchReplies"
)

// TicketServiceClient is the client API for TicketService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TicketServiceClient interface {
	CreateTicket(ctx context.Context, in *CreateTicketInput, opts ...grpc.CallOption) (*CreateTicketOutput, error)
	UpdateTicket(ctx context.Context, in *UpdateTicketInput, opts ...grpc.CallOption) (*UpdateTicketOutput, error)
	// ReadTicket marks ticket as read by participant (sets is_read flag to true).
	ReadTicket(ctx context.Context, in *ReadTicketInput, opts ...grpc.CallOption) (*ReadTicketOutput, error)
	DeleteTicket(ctx context.Context, in *DeleteTicketInput, opts ...grpc.CallOption) (*DeleteTicketOutput, error)
	DescribeTicket(ctx context.Context, in *DescribeTicketInput, opts ...grpc.CallOption) (*DescribeTicketOutput, error)
	// ListTickets fetches tickets matching criteria in the input parameter.
	ListTickets(ctx context.Context, in *ListTicketsInput, opts ...grpc.CallOption) (*ListTicketsOutput, error)
	// ReplyTicket allows to add reply to a ticket. If reply is added by participant it sets is_read and needs_reply to
	// true, otherwise, if reply added by contest administrator, this method sets these flags to false.
	ReplyTicket(ctx context.Context, in *ReplyTicketInput, opts ...grpc.CallOption) (*ReplyTicketOutput, error)
	WatchTicket(ctx context.Context, in *WatchTicketInput, opts ...grpc.CallOption) (grpc.ServerStreamingClient[WatchTicketOutput], error)
	WatchTickets(ctx context.Context, in *WatchTicketsInput, opts ...grpc.CallOption) (grpc.ServerStreamingClient[WatchTicketsOutput], error)
	WatchTicketSummary(ctx context.Context, in *WatchTicketSummaryInput, opts ...grpc.CallOption) (grpc.ServerStreamingClient[WatchTicketSummaryOutput], error)
	// ListReplies fetches replies for a particular ticket.
	ListReplies(ctx context.Context, in *ListRepliesInput, opts ...grpc.CallOption) (*ListRepliesOutput, error)
	// DeleteReply allows author to delete his own reply.
	DeleteReply(ctx context.Context, in *DeleteReplyInput, opts ...grpc.CallOption) (*DeleteReplyOutput, error)
	// UpdateReply allows author to update his own reply.
	UpdateReply(ctx context.Context, in *UpdateReplyInput, opts ...grpc.CallOption) (*UpdateReplyOutput, error)
	WatchReplies(ctx context.Context, in *WatchRepliesInput, opts ...grpc.CallOption) (grpc.ServerStreamingClient[WatchRepliesOutput], error)
}

type ticketServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTicketServiceClient(cc grpc.ClientConnInterface) TicketServiceClient {
	return &ticketServiceClient{cc}
}

func (c *ticketServiceClient) CreateTicket(ctx context.Context, in *CreateTicketInput, opts ...grpc.CallOption) (*CreateTicketOutput, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateTicketOutput)
	err := c.cc.Invoke(ctx, TicketService_CreateTicket_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ticketServiceClient) UpdateTicket(ctx context.Context, in *UpdateTicketInput, opts ...grpc.CallOption) (*UpdateTicketOutput, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateTicketOutput)
	err := c.cc.Invoke(ctx, TicketService_UpdateTicket_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ticketServiceClient) ReadTicket(ctx context.Context, in *ReadTicketInput, opts ...grpc.CallOption) (*ReadTicketOutput, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ReadTicketOutput)
	err := c.cc.Invoke(ctx, TicketService_ReadTicket_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ticketServiceClient) DeleteTicket(ctx context.Context, in *DeleteTicketInput, opts ...grpc.CallOption) (*DeleteTicketOutput, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteTicketOutput)
	err := c.cc.Invoke(ctx, TicketService_DeleteTicket_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ticketServiceClient) DescribeTicket(ctx context.Context, in *DescribeTicketInput, opts ...grpc.CallOption) (*DescribeTicketOutput, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DescribeTicketOutput)
	err := c.cc.Invoke(ctx, TicketService_DescribeTicket_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ticketServiceClient) ListTickets(ctx context.Context, in *ListTicketsInput, opts ...grpc.CallOption) (*ListTicketsOutput, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListTicketsOutput)
	err := c.cc.Invoke(ctx, TicketService_ListTickets_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ticketServiceClient) ReplyTicket(ctx context.Context, in *ReplyTicketInput, opts ...grpc.CallOption) (*ReplyTicketOutput, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ReplyTicketOutput)
	err := c.cc.Invoke(ctx, TicketService_ReplyTicket_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ticketServiceClient) WatchTicket(ctx context.Context, in *WatchTicketInput, opts ...grpc.CallOption) (grpc.ServerStreamingClient[WatchTicketOutput], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &TicketService_ServiceDesc.Streams[0], TicketService_WatchTicket_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[WatchTicketInput, WatchTicketOutput]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type TicketService_WatchTicketClient = grpc.ServerStreamingClient[WatchTicketOutput]

func (c *ticketServiceClient) WatchTickets(ctx context.Context, in *WatchTicketsInput, opts ...grpc.CallOption) (grpc.ServerStreamingClient[WatchTicketsOutput], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &TicketService_ServiceDesc.Streams[1], TicketService_WatchTickets_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[WatchTicketsInput, WatchTicketsOutput]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type TicketService_WatchTicketsClient = grpc.ServerStreamingClient[WatchTicketsOutput]

func (c *ticketServiceClient) WatchTicketSummary(ctx context.Context, in *WatchTicketSummaryInput, opts ...grpc.CallOption) (grpc.ServerStreamingClient[WatchTicketSummaryOutput], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &TicketService_ServiceDesc.Streams[2], TicketService_WatchTicketSummary_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[WatchTicketSummaryInput, WatchTicketSummaryOutput]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type TicketService_WatchTicketSummaryClient = grpc.ServerStreamingClient[WatchTicketSummaryOutput]

func (c *ticketServiceClient) ListReplies(ctx context.Context, in *ListRepliesInput, opts ...grpc.CallOption) (*ListRepliesOutput, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListRepliesOutput)
	err := c.cc.Invoke(ctx, TicketService_ListReplies_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ticketServiceClient) DeleteReply(ctx context.Context, in *DeleteReplyInput, opts ...grpc.CallOption) (*DeleteReplyOutput, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteReplyOutput)
	err := c.cc.Invoke(ctx, TicketService_DeleteReply_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ticketServiceClient) UpdateReply(ctx context.Context, in *UpdateReplyInput, opts ...grpc.CallOption) (*UpdateReplyOutput, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateReplyOutput)
	err := c.cc.Invoke(ctx, TicketService_UpdateReply_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ticketServiceClient) WatchReplies(ctx context.Context, in *WatchRepliesInput, opts ...grpc.CallOption) (grpc.ServerStreamingClient[WatchRepliesOutput], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &TicketService_ServiceDesc.Streams[3], TicketService_WatchReplies_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[WatchRepliesInput, WatchRepliesOutput]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type TicketService_WatchRepliesClient = grpc.ServerStreamingClient[WatchRepliesOutput]

// TicketServiceServer is the server API for TicketService service.
// All implementations should embed UnimplementedTicketServiceServer
// for forward compatibility.
type TicketServiceServer interface {
	CreateTicket(context.Context, *CreateTicketInput) (*CreateTicketOutput, error)
	UpdateTicket(context.Context, *UpdateTicketInput) (*UpdateTicketOutput, error)
	// ReadTicket marks ticket as read by participant (sets is_read flag to true).
	ReadTicket(context.Context, *ReadTicketInput) (*ReadTicketOutput, error)
	DeleteTicket(context.Context, *DeleteTicketInput) (*DeleteTicketOutput, error)
	DescribeTicket(context.Context, *DescribeTicketInput) (*DescribeTicketOutput, error)
	// ListTickets fetches tickets matching criteria in the input parameter.
	ListTickets(context.Context, *ListTicketsInput) (*ListTicketsOutput, error)
	// ReplyTicket allows to add reply to a ticket. If reply is added by participant it sets is_read and needs_reply to
	// true, otherwise, if reply added by contest administrator, this method sets these flags to false.
	ReplyTicket(context.Context, *ReplyTicketInput) (*ReplyTicketOutput, error)
	WatchTicket(*WatchTicketInput, grpc.ServerStreamingServer[WatchTicketOutput]) error
	WatchTickets(*WatchTicketsInput, grpc.ServerStreamingServer[WatchTicketsOutput]) error
	WatchTicketSummary(*WatchTicketSummaryInput, grpc.ServerStreamingServer[WatchTicketSummaryOutput]) error
	// ListReplies fetches replies for a particular ticket.
	ListReplies(context.Context, *ListRepliesInput) (*ListRepliesOutput, error)
	// DeleteReply allows author to delete his own reply.
	DeleteReply(context.Context, *DeleteReplyInput) (*DeleteReplyOutput, error)
	// UpdateReply allows author to update his own reply.
	UpdateReply(context.Context, *UpdateReplyInput) (*UpdateReplyOutput, error)
	WatchReplies(*WatchRepliesInput, grpc.ServerStreamingServer[WatchRepliesOutput]) error
}

// UnimplementedTicketServiceServer should be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedTicketServiceServer struct{}

func (UnimplementedTicketServiceServer) CreateTicket(context.Context, *CreateTicketInput) (*CreateTicketOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTicket not implemented")
}
func (UnimplementedTicketServiceServer) UpdateTicket(context.Context, *UpdateTicketInput) (*UpdateTicketOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateTicket not implemented")
}
func (UnimplementedTicketServiceServer) ReadTicket(context.Context, *ReadTicketInput) (*ReadTicketOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadTicket not implemented")
}
func (UnimplementedTicketServiceServer) DeleteTicket(context.Context, *DeleteTicketInput) (*DeleteTicketOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTicket not implemented")
}
func (UnimplementedTicketServiceServer) DescribeTicket(context.Context, *DescribeTicketInput) (*DescribeTicketOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DescribeTicket not implemented")
}
func (UnimplementedTicketServiceServer) ListTickets(context.Context, *ListTicketsInput) (*ListTicketsOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListTickets not implemented")
}
func (UnimplementedTicketServiceServer) ReplyTicket(context.Context, *ReplyTicketInput) (*ReplyTicketOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReplyTicket not implemented")
}
func (UnimplementedTicketServiceServer) WatchTicket(*WatchTicketInput, grpc.ServerStreamingServer[WatchTicketOutput]) error {
	return status.Errorf(codes.Unimplemented, "method WatchTicket not implemented")
}
func (UnimplementedTicketServiceServer) WatchTickets(*WatchTicketsInput, grpc.ServerStreamingServer[WatchTicketsOutput]) error {
	return status.Errorf(codes.Unimplemented, "method WatchTickets not implemented")
}
func (UnimplementedTicketServiceServer) WatchTicketSummary(*WatchTicketSummaryInput, grpc.ServerStreamingServer[WatchTicketSummaryOutput]) error {
	return status.Errorf(codes.Unimplemented, "method WatchTicketSummary not implemented")
}
func (UnimplementedTicketServiceServer) ListReplies(context.Context, *ListRepliesInput) (*ListRepliesOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListReplies not implemented")
}
func (UnimplementedTicketServiceServer) DeleteReply(context.Context, *DeleteReplyInput) (*DeleteReplyOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteReply not implemented")
}
func (UnimplementedTicketServiceServer) UpdateReply(context.Context, *UpdateReplyInput) (*UpdateReplyOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateReply not implemented")
}
func (UnimplementedTicketServiceServer) WatchReplies(*WatchRepliesInput, grpc.ServerStreamingServer[WatchRepliesOutput]) error {
	return status.Errorf(codes.Unimplemented, "method WatchReplies not implemented")
}
func (UnimplementedTicketServiceServer) testEmbeddedByValue() {}

// UnsafeTicketServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TicketServiceServer will
// result in compilation errors.
type UnsafeTicketServiceServer interface {
	mustEmbedUnimplementedTicketServiceServer()
}

func RegisterTicketServiceServer(s grpc.ServiceRegistrar, srv TicketServiceServer) {
	// If the following call pancis, it indicates UnimplementedTicketServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&TicketService_ServiceDesc, srv)
}

func _TicketService_CreateTicket_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTicketInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TicketServiceServer).CreateTicket(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TicketService_CreateTicket_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TicketServiceServer).CreateTicket(ctx, req.(*CreateTicketInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _TicketService_UpdateTicket_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateTicketInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TicketServiceServer).UpdateTicket(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TicketService_UpdateTicket_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TicketServiceServer).UpdateTicket(ctx, req.(*UpdateTicketInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _TicketService_ReadTicket_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadTicketInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TicketServiceServer).ReadTicket(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TicketService_ReadTicket_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TicketServiceServer).ReadTicket(ctx, req.(*ReadTicketInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _TicketService_DeleteTicket_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteTicketInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TicketServiceServer).DeleteTicket(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TicketService_DeleteTicket_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TicketServiceServer).DeleteTicket(ctx, req.(*DeleteTicketInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _TicketService_DescribeTicket_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DescribeTicketInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TicketServiceServer).DescribeTicket(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TicketService_DescribeTicket_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TicketServiceServer).DescribeTicket(ctx, req.(*DescribeTicketInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _TicketService_ListTickets_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListTicketsInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TicketServiceServer).ListTickets(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TicketService_ListTickets_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TicketServiceServer).ListTickets(ctx, req.(*ListTicketsInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _TicketService_ReplyTicket_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReplyTicketInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TicketServiceServer).ReplyTicket(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TicketService_ReplyTicket_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TicketServiceServer).ReplyTicket(ctx, req.(*ReplyTicketInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _TicketService_WatchTicket_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(WatchTicketInput)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(TicketServiceServer).WatchTicket(m, &grpc.GenericServerStream[WatchTicketInput, WatchTicketOutput]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type TicketService_WatchTicketServer = grpc.ServerStreamingServer[WatchTicketOutput]

func _TicketService_WatchTickets_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(WatchTicketsInput)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(TicketServiceServer).WatchTickets(m, &grpc.GenericServerStream[WatchTicketsInput, WatchTicketsOutput]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type TicketService_WatchTicketsServer = grpc.ServerStreamingServer[WatchTicketsOutput]

func _TicketService_WatchTicketSummary_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(WatchTicketSummaryInput)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(TicketServiceServer).WatchTicketSummary(m, &grpc.GenericServerStream[WatchTicketSummaryInput, WatchTicketSummaryOutput]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type TicketService_WatchTicketSummaryServer = grpc.ServerStreamingServer[WatchTicketSummaryOutput]

func _TicketService_ListReplies_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRepliesInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TicketServiceServer).ListReplies(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TicketService_ListReplies_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TicketServiceServer).ListReplies(ctx, req.(*ListRepliesInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _TicketService_DeleteReply_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteReplyInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TicketServiceServer).DeleteReply(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TicketService_DeleteReply_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TicketServiceServer).DeleteReply(ctx, req.(*DeleteReplyInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _TicketService_UpdateReply_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateReplyInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TicketServiceServer).UpdateReply(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TicketService_UpdateReply_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TicketServiceServer).UpdateReply(ctx, req.(*UpdateReplyInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _TicketService_WatchReplies_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(WatchRepliesInput)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(TicketServiceServer).WatchReplies(m, &grpc.GenericServerStream[WatchRepliesInput, WatchRepliesOutput]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type TicketService_WatchRepliesServer = grpc.ServerStreamingServer[WatchRepliesOutput]

// TicketService_ServiceDesc is the grpc.ServiceDesc for TicketService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TicketService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "eolymp.judge.TicketService",
	HandlerType: (*TicketServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateTicket",
			Handler:    _TicketService_CreateTicket_Handler,
		},
		{
			MethodName: "UpdateTicket",
			Handler:    _TicketService_UpdateTicket_Handler,
		},
		{
			MethodName: "ReadTicket",
			Handler:    _TicketService_ReadTicket_Handler,
		},
		{
			MethodName: "DeleteTicket",
			Handler:    _TicketService_DeleteTicket_Handler,
		},
		{
			MethodName: "DescribeTicket",
			Handler:    _TicketService_DescribeTicket_Handler,
		},
		{
			MethodName: "ListTickets",
			Handler:    _TicketService_ListTickets_Handler,
		},
		{
			MethodName: "ReplyTicket",
			Handler:    _TicketService_ReplyTicket_Handler,
		},
		{
			MethodName: "ListReplies",
			Handler:    _TicketService_ListReplies_Handler,
		},
		{
			MethodName: "DeleteReply",
			Handler:    _TicketService_DeleteReply_Handler,
		},
		{
			MethodName: "UpdateReply",
			Handler:    _TicketService_UpdateReply_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "WatchTicket",
			Handler:       _TicketService_WatchTicket_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "WatchTickets",
			Handler:       _TicketService_WatchTickets_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "WatchTicketSummary",
			Handler:       _TicketService_WatchTicketSummary_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "WatchReplies",
			Handler:       _TicketService_WatchReplies_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "eolymp/judge/ticket_service.proto",
}
