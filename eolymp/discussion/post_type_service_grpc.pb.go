// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v3.21.12
// source: eolymp/discussion/post_type_service.proto

package discussion

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	PostTypeService_DescribePostType_FullMethodName = "/eolymp.discussion.PostTypeService/DescribePostType"
	PostTypeService_ListPostTypes_FullMethodName    = "/eolymp.discussion.PostTypeService/ListPostTypes"
	PostTypeService_CreatePostType_FullMethodName   = "/eolymp.discussion.PostTypeService/CreatePostType"
	PostTypeService_UpdatePostType_FullMethodName   = "/eolymp.discussion.PostTypeService/UpdatePostType"
	PostTypeService_DeletePostType_FullMethodName   = "/eolymp.discussion.PostTypeService/DeletePostType"
)

// PostTypeServiceClient is the client API for PostTypeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PostTypeServiceClient interface {
	DescribePostType(ctx context.Context, in *DescribePostTypeInput, opts ...grpc.CallOption) (*DescribePostTypeOutput, error)
	ListPostTypes(ctx context.Context, in *ListPostTypesInput, opts ...grpc.CallOption) (*ListPostTypesOutput, error)
	CreatePostType(ctx context.Context, in *CreatePostTypeInput, opts ...grpc.CallOption) (*CreatePostTypeOutput, error)
	UpdatePostType(ctx context.Context, in *UpdatePostTypeInput, opts ...grpc.CallOption) (*UpdatePostTypeOutput, error)
	DeletePostType(ctx context.Context, in *DeletePostTypeInput, opts ...grpc.CallOption) (*DeletePostTypeOutput, error)
}

type postTypeServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPostTypeServiceClient(cc grpc.ClientConnInterface) PostTypeServiceClient {
	return &postTypeServiceClient{cc}
}

func (c *postTypeServiceClient) DescribePostType(ctx context.Context, in *DescribePostTypeInput, opts ...grpc.CallOption) (*DescribePostTypeOutput, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DescribePostTypeOutput)
	err := c.cc.Invoke(ctx, PostTypeService_DescribePostType_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postTypeServiceClient) ListPostTypes(ctx context.Context, in *ListPostTypesInput, opts ...grpc.CallOption) (*ListPostTypesOutput, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListPostTypesOutput)
	err := c.cc.Invoke(ctx, PostTypeService_ListPostTypes_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postTypeServiceClient) CreatePostType(ctx context.Context, in *CreatePostTypeInput, opts ...grpc.CallOption) (*CreatePostTypeOutput, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreatePostTypeOutput)
	err := c.cc.Invoke(ctx, PostTypeService_CreatePostType_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postTypeServiceClient) UpdatePostType(ctx context.Context, in *UpdatePostTypeInput, opts ...grpc.CallOption) (*UpdatePostTypeOutput, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdatePostTypeOutput)
	err := c.cc.Invoke(ctx, PostTypeService_UpdatePostType_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postTypeServiceClient) DeletePostType(ctx context.Context, in *DeletePostTypeInput, opts ...grpc.CallOption) (*DeletePostTypeOutput, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeletePostTypeOutput)
	err := c.cc.Invoke(ctx, PostTypeService_DeletePostType_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PostTypeServiceServer is the server API for PostTypeService service.
// All implementations should embed UnimplementedPostTypeServiceServer
// for forward compatibility
type PostTypeServiceServer interface {
	DescribePostType(context.Context, *DescribePostTypeInput) (*DescribePostTypeOutput, error)
	ListPostTypes(context.Context, *ListPostTypesInput) (*ListPostTypesOutput, error)
	CreatePostType(context.Context, *CreatePostTypeInput) (*CreatePostTypeOutput, error)
	UpdatePostType(context.Context, *UpdatePostTypeInput) (*UpdatePostTypeOutput, error)
	DeletePostType(context.Context, *DeletePostTypeInput) (*DeletePostTypeOutput, error)
}

// UnimplementedPostTypeServiceServer should be embedded to have forward compatible implementations.
type UnimplementedPostTypeServiceServer struct {
}

func (UnimplementedPostTypeServiceServer) DescribePostType(context.Context, *DescribePostTypeInput) (*DescribePostTypeOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DescribePostType not implemented")
}
func (UnimplementedPostTypeServiceServer) ListPostTypes(context.Context, *ListPostTypesInput) (*ListPostTypesOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListPostTypes not implemented")
}
func (UnimplementedPostTypeServiceServer) CreatePostType(context.Context, *CreatePostTypeInput) (*CreatePostTypeOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePostType not implemented")
}
func (UnimplementedPostTypeServiceServer) UpdatePostType(context.Context, *UpdatePostTypeInput) (*UpdatePostTypeOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePostType not implemented")
}
func (UnimplementedPostTypeServiceServer) DeletePostType(context.Context, *DeletePostTypeInput) (*DeletePostTypeOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeletePostType not implemented")
}

// UnsafePostTypeServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PostTypeServiceServer will
// result in compilation errors.
type UnsafePostTypeServiceServer interface {
	mustEmbedUnimplementedPostTypeServiceServer()
}

func RegisterPostTypeServiceServer(s grpc.ServiceRegistrar, srv PostTypeServiceServer) {
	s.RegisterService(&PostTypeService_ServiceDesc, srv)
}

func _PostTypeService_DescribePostType_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DescribePostTypeInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostTypeServiceServer).DescribePostType(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PostTypeService_DescribePostType_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostTypeServiceServer).DescribePostType(ctx, req.(*DescribePostTypeInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _PostTypeService_ListPostTypes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListPostTypesInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostTypeServiceServer).ListPostTypes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PostTypeService_ListPostTypes_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostTypeServiceServer).ListPostTypes(ctx, req.(*ListPostTypesInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _PostTypeService_CreatePostType_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePostTypeInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostTypeServiceServer).CreatePostType(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PostTypeService_CreatePostType_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostTypeServiceServer).CreatePostType(ctx, req.(*CreatePostTypeInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _PostTypeService_UpdatePostType_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePostTypeInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostTypeServiceServer).UpdatePostType(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PostTypeService_UpdatePostType_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostTypeServiceServer).UpdatePostType(ctx, req.(*UpdatePostTypeInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _PostTypeService_DeletePostType_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeletePostTypeInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostTypeServiceServer).DeletePostType(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PostTypeService_DeletePostType_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostTypeServiceServer).DeletePostType(ctx, req.(*DeletePostTypeInput))
	}
	return interceptor(ctx, in, info, handler)
}

// PostTypeService_ServiceDesc is the grpc.ServiceDesc for PostTypeService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PostTypeService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "eolymp.discussion.PostTypeService",
	HandlerType: (*PostTypeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DescribePostType",
			Handler:    _PostTypeService_DescribePostType_Handler,
		},
		{
			MethodName: "ListPostTypes",
			Handler:    _PostTypeService_ListPostTypes_Handler,
		},
		{
			MethodName: "CreatePostType",
			Handler:    _PostTypeService_CreatePostType_Handler,
		},
		{
			MethodName: "UpdatePostType",
			Handler:    _PostTypeService_UpdatePostType_Handler,
		},
		{
			MethodName: "DeletePostType",
			Handler:    _PostTypeService_DeletePostType_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "eolymp/discussion/post_type_service.proto",
}
