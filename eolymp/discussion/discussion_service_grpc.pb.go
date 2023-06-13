// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.18.1
// source: eolymp/discussion/discussion_service.proto

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
	DiscussionService_DescribeDiscussion_FullMethodName = "/eolymp.discussion.DiscussionService/DescribeDiscussion"
	DiscussionService_ListDiscussions_FullMethodName    = "/eolymp.discussion.DiscussionService/ListDiscussions"
	DiscussionService_CreateDiscussion_FullMethodName   = "/eolymp.discussion.DiscussionService/CreateDiscussion"
	DiscussionService_UpdateDiscussion_FullMethodName   = "/eolymp.discussion.DiscussionService/UpdateDiscussion"
	DiscussionService_DeleteDiscussion_FullMethodName   = "/eolymp.discussion.DiscussionService/DeleteDiscussion"
	DiscussionService_VoteDiscussion_FullMethodName     = "/eolymp.discussion.DiscussionService/VoteDiscussion"
)

// DiscussionServiceClient is the client API for DiscussionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DiscussionServiceClient interface {
	DescribeDiscussion(ctx context.Context, in *DescribeDiscussionInput, opts ...grpc.CallOption) (*DescribeDiscussionOutput, error)
	ListDiscussions(ctx context.Context, in *ListDiscussionsInput, opts ...grpc.CallOption) (*ListDiscussionsOutput, error)
	CreateDiscussion(ctx context.Context, in *CreateDiscussionInput, opts ...grpc.CallOption) (*CreateDiscussionOutput, error)
	UpdateDiscussion(ctx context.Context, in *UpdateDiscussionInput, opts ...grpc.CallOption) (*UpdateDiscussionOutput, error)
	DeleteDiscussion(ctx context.Context, in *DeleteDiscussionInput, opts ...grpc.CallOption) (*DeleteDiscussionOutput, error)
	VoteDiscussion(ctx context.Context, in *VoteDiscussionInput, opts ...grpc.CallOption) (*VoteDiscussionOutput, error)
}

type discussionServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDiscussionServiceClient(cc grpc.ClientConnInterface) DiscussionServiceClient {
	return &discussionServiceClient{cc}
}

func (c *discussionServiceClient) DescribeDiscussion(ctx context.Context, in *DescribeDiscussionInput, opts ...grpc.CallOption) (*DescribeDiscussionOutput, error) {
	out := new(DescribeDiscussionOutput)
	err := c.cc.Invoke(ctx, DiscussionService_DescribeDiscussion_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *discussionServiceClient) ListDiscussions(ctx context.Context, in *ListDiscussionsInput, opts ...grpc.CallOption) (*ListDiscussionsOutput, error) {
	out := new(ListDiscussionsOutput)
	err := c.cc.Invoke(ctx, DiscussionService_ListDiscussions_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *discussionServiceClient) CreateDiscussion(ctx context.Context, in *CreateDiscussionInput, opts ...grpc.CallOption) (*CreateDiscussionOutput, error) {
	out := new(CreateDiscussionOutput)
	err := c.cc.Invoke(ctx, DiscussionService_CreateDiscussion_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *discussionServiceClient) UpdateDiscussion(ctx context.Context, in *UpdateDiscussionInput, opts ...grpc.CallOption) (*UpdateDiscussionOutput, error) {
	out := new(UpdateDiscussionOutput)
	err := c.cc.Invoke(ctx, DiscussionService_UpdateDiscussion_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *discussionServiceClient) DeleteDiscussion(ctx context.Context, in *DeleteDiscussionInput, opts ...grpc.CallOption) (*DeleteDiscussionOutput, error) {
	out := new(DeleteDiscussionOutput)
	err := c.cc.Invoke(ctx, DiscussionService_DeleteDiscussion_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *discussionServiceClient) VoteDiscussion(ctx context.Context, in *VoteDiscussionInput, opts ...grpc.CallOption) (*VoteDiscussionOutput, error) {
	out := new(VoteDiscussionOutput)
	err := c.cc.Invoke(ctx, DiscussionService_VoteDiscussion_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DiscussionServiceServer is the server API for DiscussionService service.
// All implementations should embed UnimplementedDiscussionServiceServer
// for forward compatibility
type DiscussionServiceServer interface {
	DescribeDiscussion(context.Context, *DescribeDiscussionInput) (*DescribeDiscussionOutput, error)
	ListDiscussions(context.Context, *ListDiscussionsInput) (*ListDiscussionsOutput, error)
	CreateDiscussion(context.Context, *CreateDiscussionInput) (*CreateDiscussionOutput, error)
	UpdateDiscussion(context.Context, *UpdateDiscussionInput) (*UpdateDiscussionOutput, error)
	DeleteDiscussion(context.Context, *DeleteDiscussionInput) (*DeleteDiscussionOutput, error)
	VoteDiscussion(context.Context, *VoteDiscussionInput) (*VoteDiscussionOutput, error)
}

// UnimplementedDiscussionServiceServer should be embedded to have forward compatible implementations.
type UnimplementedDiscussionServiceServer struct {
}

func (UnimplementedDiscussionServiceServer) DescribeDiscussion(context.Context, *DescribeDiscussionInput) (*DescribeDiscussionOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DescribeDiscussion not implemented")
}
func (UnimplementedDiscussionServiceServer) ListDiscussions(context.Context, *ListDiscussionsInput) (*ListDiscussionsOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListDiscussions not implemented")
}
func (UnimplementedDiscussionServiceServer) CreateDiscussion(context.Context, *CreateDiscussionInput) (*CreateDiscussionOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateDiscussion not implemented")
}
func (UnimplementedDiscussionServiceServer) UpdateDiscussion(context.Context, *UpdateDiscussionInput) (*UpdateDiscussionOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateDiscussion not implemented")
}
func (UnimplementedDiscussionServiceServer) DeleteDiscussion(context.Context, *DeleteDiscussionInput) (*DeleteDiscussionOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteDiscussion not implemented")
}
func (UnimplementedDiscussionServiceServer) VoteDiscussion(context.Context, *VoteDiscussionInput) (*VoteDiscussionOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VoteDiscussion not implemented")
}

// UnsafeDiscussionServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DiscussionServiceServer will
// result in compilation errors.
type UnsafeDiscussionServiceServer interface {
	mustEmbedUnimplementedDiscussionServiceServer()
}

func RegisterDiscussionServiceServer(s grpc.ServiceRegistrar, srv DiscussionServiceServer) {
	s.RegisterService(&DiscussionService_ServiceDesc, srv)
}

func _DiscussionService_DescribeDiscussion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DescribeDiscussionInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DiscussionServiceServer).DescribeDiscussion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DiscussionService_DescribeDiscussion_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DiscussionServiceServer).DescribeDiscussion(ctx, req.(*DescribeDiscussionInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _DiscussionService_ListDiscussions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListDiscussionsInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DiscussionServiceServer).ListDiscussions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DiscussionService_ListDiscussions_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DiscussionServiceServer).ListDiscussions(ctx, req.(*ListDiscussionsInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _DiscussionService_CreateDiscussion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateDiscussionInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DiscussionServiceServer).CreateDiscussion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DiscussionService_CreateDiscussion_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DiscussionServiceServer).CreateDiscussion(ctx, req.(*CreateDiscussionInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _DiscussionService_UpdateDiscussion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateDiscussionInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DiscussionServiceServer).UpdateDiscussion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DiscussionService_UpdateDiscussion_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DiscussionServiceServer).UpdateDiscussion(ctx, req.(*UpdateDiscussionInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _DiscussionService_DeleteDiscussion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteDiscussionInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DiscussionServiceServer).DeleteDiscussion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DiscussionService_DeleteDiscussion_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DiscussionServiceServer).DeleteDiscussion(ctx, req.(*DeleteDiscussionInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _DiscussionService_VoteDiscussion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VoteDiscussionInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DiscussionServiceServer).VoteDiscussion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DiscussionService_VoteDiscussion_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DiscussionServiceServer).VoteDiscussion(ctx, req.(*VoteDiscussionInput))
	}
	return interceptor(ctx, in, info, handler)
}

// DiscussionService_ServiceDesc is the grpc.ServiceDesc for DiscussionService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DiscussionService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "eolymp.discussion.DiscussionService",
	HandlerType: (*DiscussionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DescribeDiscussion",
			Handler:    _DiscussionService_DescribeDiscussion_Handler,
		},
		{
			MethodName: "ListDiscussions",
			Handler:    _DiscussionService_ListDiscussions_Handler,
		},
		{
			MethodName: "CreateDiscussion",
			Handler:    _DiscussionService_CreateDiscussion_Handler,
		},
		{
			MethodName: "UpdateDiscussion",
			Handler:    _DiscussionService_UpdateDiscussion_Handler,
		},
		{
			MethodName: "DeleteDiscussion",
			Handler:    _DiscussionService_DeleteDiscussion_Handler,
		},
		{
			MethodName: "VoteDiscussion",
			Handler:    _DiscussionService_VoteDiscussion_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "eolymp/discussion/discussion_service.proto",
}