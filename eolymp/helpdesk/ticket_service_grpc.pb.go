// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: eolymp/helpdesk/ticket_service.proto

package helpdesk

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
	TicketService_CreateTicket_FullMethodName     = "/eolymp.helpdesk.TicketService/CreateTicket"
	TicketService_UpdateTicket_FullMethodName     = "/eolymp.helpdesk.TicketService/UpdateTicket"
	TicketService_DeleteTicket_FullMethodName     = "/eolymp.helpdesk.TicketService/DeleteTicket"
	TicketService_DescribeTicket_FullMethodName   = "/eolymp.helpdesk.TicketService/DescribeTicket"
	TicketService_ListTickets_FullMethodName      = "/eolymp.helpdesk.TicketService/ListTickets"
	TicketService_ApproveTicket_FullMethodName    = "/eolymp.helpdesk.TicketService/ApproveTicket"
	TicketService_RejectTicket_FullMethodName     = "/eolymp.helpdesk.TicketService/RejectTicket"
	TicketService_CloseTicket_FullMethodName      = "/eolymp.helpdesk.TicketService/CloseTicket"
	TicketService_AddComment_FullMethodName       = "/eolymp.helpdesk.TicketService/AddComment"
	TicketService_UpdateComment_FullMethodName    = "/eolymp.helpdesk.TicketService/UpdateComment"
	TicketService_DeleteComment_FullMethodName    = "/eolymp.helpdesk.TicketService/DeleteComment"
	TicketService_ListComments_FullMethodName     = "/eolymp.helpdesk.TicketService/ListComments"
	TicketService_DescribeComment_FullMethodName  = "/eolymp.helpdesk.TicketService/DescribeComment"
	TicketService_UploadAttachment_FullMethodName = "/eolymp.helpdesk.TicketService/UploadAttachment"
)

// TicketServiceClient is the client API for TicketService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TicketServiceClient interface {
	CreateTicket(ctx context.Context, in *CreateTicketInput, opts ...grpc.CallOption) (*CreateTicketOutput, error)
	UpdateTicket(ctx context.Context, in *UpdateTicketInput, opts ...grpc.CallOption) (*UpdateTicketOutput, error)
	DeleteTicket(ctx context.Context, in *DeleteTicketInput, opts ...grpc.CallOption) (*DeleteTicketOutput, error)
	DescribeTicket(ctx context.Context, in *DescribeTicketInput, opts ...grpc.CallOption) (*DescribeTicketOutput, error)
	ListTickets(ctx context.Context, in *ListTicketsInput, opts ...grpc.CallOption) (*ListTicketsOutput, error)
	ApproveTicket(ctx context.Context, in *ApproveTicketInput, opts ...grpc.CallOption) (*ApproveTicketOutput, error)
	RejectTicket(ctx context.Context, in *RejectTicketInput, opts ...grpc.CallOption) (*RejectTicketOutput, error)
	CloseTicket(ctx context.Context, in *CloseTicketInput, opts ...grpc.CallOption) (*CloseTicketOutput, error)
	AddComment(ctx context.Context, in *AddCommentInput, opts ...grpc.CallOption) (*AddCommentOutput, error)
	UpdateComment(ctx context.Context, in *UpdateCommentInput, opts ...grpc.CallOption) (*UpdateCommentOutput, error)
	DeleteComment(ctx context.Context, in *DeleteCommentInput, opts ...grpc.CallOption) (*DeleteCommentOutput, error)
	ListComments(ctx context.Context, in *ListCommentsInput, opts ...grpc.CallOption) (*ListCommentsOutput, error)
	DescribeComment(ctx context.Context, in *DescribeCommentInput, opts ...grpc.CallOption) (*DescribeCommentOutput, error)
	UploadAttachment(ctx context.Context, in *UploadAttachmentInput, opts ...grpc.CallOption) (*UploadAttachmentOutput, error)
}

type ticketServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTicketServiceClient(cc grpc.ClientConnInterface) TicketServiceClient {
	return &ticketServiceClient{cc}
}

func (c *ticketServiceClient) CreateTicket(ctx context.Context, in *CreateTicketInput, opts ...grpc.CallOption) (*CreateTicketOutput, error) {
	out := new(CreateTicketOutput)
	err := c.cc.Invoke(ctx, TicketService_CreateTicket_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ticketServiceClient) UpdateTicket(ctx context.Context, in *UpdateTicketInput, opts ...grpc.CallOption) (*UpdateTicketOutput, error) {
	out := new(UpdateTicketOutput)
	err := c.cc.Invoke(ctx, TicketService_UpdateTicket_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ticketServiceClient) DeleteTicket(ctx context.Context, in *DeleteTicketInput, opts ...grpc.CallOption) (*DeleteTicketOutput, error) {
	out := new(DeleteTicketOutput)
	err := c.cc.Invoke(ctx, TicketService_DeleteTicket_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ticketServiceClient) DescribeTicket(ctx context.Context, in *DescribeTicketInput, opts ...grpc.CallOption) (*DescribeTicketOutput, error) {
	out := new(DescribeTicketOutput)
	err := c.cc.Invoke(ctx, TicketService_DescribeTicket_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ticketServiceClient) ListTickets(ctx context.Context, in *ListTicketsInput, opts ...grpc.CallOption) (*ListTicketsOutput, error) {
	out := new(ListTicketsOutput)
	err := c.cc.Invoke(ctx, TicketService_ListTickets_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ticketServiceClient) ApproveTicket(ctx context.Context, in *ApproveTicketInput, opts ...grpc.CallOption) (*ApproveTicketOutput, error) {
	out := new(ApproveTicketOutput)
	err := c.cc.Invoke(ctx, TicketService_ApproveTicket_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ticketServiceClient) RejectTicket(ctx context.Context, in *RejectTicketInput, opts ...grpc.CallOption) (*RejectTicketOutput, error) {
	out := new(RejectTicketOutput)
	err := c.cc.Invoke(ctx, TicketService_RejectTicket_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ticketServiceClient) CloseTicket(ctx context.Context, in *CloseTicketInput, opts ...grpc.CallOption) (*CloseTicketOutput, error) {
	out := new(CloseTicketOutput)
	err := c.cc.Invoke(ctx, TicketService_CloseTicket_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ticketServiceClient) AddComment(ctx context.Context, in *AddCommentInput, opts ...grpc.CallOption) (*AddCommentOutput, error) {
	out := new(AddCommentOutput)
	err := c.cc.Invoke(ctx, TicketService_AddComment_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ticketServiceClient) UpdateComment(ctx context.Context, in *UpdateCommentInput, opts ...grpc.CallOption) (*UpdateCommentOutput, error) {
	out := new(UpdateCommentOutput)
	err := c.cc.Invoke(ctx, TicketService_UpdateComment_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ticketServiceClient) DeleteComment(ctx context.Context, in *DeleteCommentInput, opts ...grpc.CallOption) (*DeleteCommentOutput, error) {
	out := new(DeleteCommentOutput)
	err := c.cc.Invoke(ctx, TicketService_DeleteComment_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ticketServiceClient) ListComments(ctx context.Context, in *ListCommentsInput, opts ...grpc.CallOption) (*ListCommentsOutput, error) {
	out := new(ListCommentsOutput)
	err := c.cc.Invoke(ctx, TicketService_ListComments_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ticketServiceClient) DescribeComment(ctx context.Context, in *DescribeCommentInput, opts ...grpc.CallOption) (*DescribeCommentOutput, error) {
	out := new(DescribeCommentOutput)
	err := c.cc.Invoke(ctx, TicketService_DescribeComment_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ticketServiceClient) UploadAttachment(ctx context.Context, in *UploadAttachmentInput, opts ...grpc.CallOption) (*UploadAttachmentOutput, error) {
	out := new(UploadAttachmentOutput)
	err := c.cc.Invoke(ctx, TicketService_UploadAttachment_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TicketServiceServer is the server API for TicketService service.
// All implementations should embed UnimplementedTicketServiceServer
// for forward compatibility
type TicketServiceServer interface {
	CreateTicket(context.Context, *CreateTicketInput) (*CreateTicketOutput, error)
	UpdateTicket(context.Context, *UpdateTicketInput) (*UpdateTicketOutput, error)
	DeleteTicket(context.Context, *DeleteTicketInput) (*DeleteTicketOutput, error)
	DescribeTicket(context.Context, *DescribeTicketInput) (*DescribeTicketOutput, error)
	ListTickets(context.Context, *ListTicketsInput) (*ListTicketsOutput, error)
	ApproveTicket(context.Context, *ApproveTicketInput) (*ApproveTicketOutput, error)
	RejectTicket(context.Context, *RejectTicketInput) (*RejectTicketOutput, error)
	CloseTicket(context.Context, *CloseTicketInput) (*CloseTicketOutput, error)
	AddComment(context.Context, *AddCommentInput) (*AddCommentOutput, error)
	UpdateComment(context.Context, *UpdateCommentInput) (*UpdateCommentOutput, error)
	DeleteComment(context.Context, *DeleteCommentInput) (*DeleteCommentOutput, error)
	ListComments(context.Context, *ListCommentsInput) (*ListCommentsOutput, error)
	DescribeComment(context.Context, *DescribeCommentInput) (*DescribeCommentOutput, error)
	UploadAttachment(context.Context, *UploadAttachmentInput) (*UploadAttachmentOutput, error)
}

// UnimplementedTicketServiceServer should be embedded to have forward compatible implementations.
type UnimplementedTicketServiceServer struct {
}

func (UnimplementedTicketServiceServer) CreateTicket(context.Context, *CreateTicketInput) (*CreateTicketOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTicket not implemented")
}
func (UnimplementedTicketServiceServer) UpdateTicket(context.Context, *UpdateTicketInput) (*UpdateTicketOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateTicket not implemented")
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
func (UnimplementedTicketServiceServer) ApproveTicket(context.Context, *ApproveTicketInput) (*ApproveTicketOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ApproveTicket not implemented")
}
func (UnimplementedTicketServiceServer) RejectTicket(context.Context, *RejectTicketInput) (*RejectTicketOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RejectTicket not implemented")
}
func (UnimplementedTicketServiceServer) CloseTicket(context.Context, *CloseTicketInput) (*CloseTicketOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CloseTicket not implemented")
}
func (UnimplementedTicketServiceServer) AddComment(context.Context, *AddCommentInput) (*AddCommentOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddComment not implemented")
}
func (UnimplementedTicketServiceServer) UpdateComment(context.Context, *UpdateCommentInput) (*UpdateCommentOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateComment not implemented")
}
func (UnimplementedTicketServiceServer) DeleteComment(context.Context, *DeleteCommentInput) (*DeleteCommentOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteComment not implemented")
}
func (UnimplementedTicketServiceServer) ListComments(context.Context, *ListCommentsInput) (*ListCommentsOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListComments not implemented")
}
func (UnimplementedTicketServiceServer) DescribeComment(context.Context, *DescribeCommentInput) (*DescribeCommentOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DescribeComment not implemented")
}
func (UnimplementedTicketServiceServer) UploadAttachment(context.Context, *UploadAttachmentInput) (*UploadAttachmentOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UploadAttachment not implemented")
}

// UnsafeTicketServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TicketServiceServer will
// result in compilation errors.
type UnsafeTicketServiceServer interface {
	mustEmbedUnimplementedTicketServiceServer()
}

func RegisterTicketServiceServer(s grpc.ServiceRegistrar, srv TicketServiceServer) {
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

func _TicketService_ApproveTicket_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ApproveTicketInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TicketServiceServer).ApproveTicket(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TicketService_ApproveTicket_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TicketServiceServer).ApproveTicket(ctx, req.(*ApproveTicketInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _TicketService_RejectTicket_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RejectTicketInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TicketServiceServer).RejectTicket(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TicketService_RejectTicket_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TicketServiceServer).RejectTicket(ctx, req.(*RejectTicketInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _TicketService_CloseTicket_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CloseTicketInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TicketServiceServer).CloseTicket(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TicketService_CloseTicket_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TicketServiceServer).CloseTicket(ctx, req.(*CloseTicketInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _TicketService_AddComment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddCommentInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TicketServiceServer).AddComment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TicketService_AddComment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TicketServiceServer).AddComment(ctx, req.(*AddCommentInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _TicketService_UpdateComment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateCommentInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TicketServiceServer).UpdateComment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TicketService_UpdateComment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TicketServiceServer).UpdateComment(ctx, req.(*UpdateCommentInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _TicketService_DeleteComment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteCommentInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TicketServiceServer).DeleteComment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TicketService_DeleteComment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TicketServiceServer).DeleteComment(ctx, req.(*DeleteCommentInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _TicketService_ListComments_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListCommentsInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TicketServiceServer).ListComments(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TicketService_ListComments_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TicketServiceServer).ListComments(ctx, req.(*ListCommentsInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _TicketService_DescribeComment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DescribeCommentInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TicketServiceServer).DescribeComment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TicketService_DescribeComment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TicketServiceServer).DescribeComment(ctx, req.(*DescribeCommentInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _TicketService_UploadAttachment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UploadAttachmentInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TicketServiceServer).UploadAttachment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TicketService_UploadAttachment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TicketServiceServer).UploadAttachment(ctx, req.(*UploadAttachmentInput))
	}
	return interceptor(ctx, in, info, handler)
}

// TicketService_ServiceDesc is the grpc.ServiceDesc for TicketService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TicketService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "eolymp.helpdesk.TicketService",
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
			MethodName: "ApproveTicket",
			Handler:    _TicketService_ApproveTicket_Handler,
		},
		{
			MethodName: "RejectTicket",
			Handler:    _TicketService_RejectTicket_Handler,
		},
		{
			MethodName: "CloseTicket",
			Handler:    _TicketService_CloseTicket_Handler,
		},
		{
			MethodName: "AddComment",
			Handler:    _TicketService_AddComment_Handler,
		},
		{
			MethodName: "UpdateComment",
			Handler:    _TicketService_UpdateComment_Handler,
		},
		{
			MethodName: "DeleteComment",
			Handler:    _TicketService_DeleteComment_Handler,
		},
		{
			MethodName: "ListComments",
			Handler:    _TicketService_ListComments_Handler,
		},
		{
			MethodName: "DescribeComment",
			Handler:    _TicketService_DescribeComment_Handler,
		},
		{
			MethodName: "UploadAttachment",
			Handler:    _TicketService_UploadAttachment_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "eolymp/helpdesk/ticket_service.proto",
}
