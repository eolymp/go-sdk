// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.21.12
// source: eolymp/discussion/post_service.proto

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
	PostService_DescribePost_FullMethodName            = "/eolymp.discussion.PostService/DescribePost"
	PostService_ListPosts_FullMethodName               = "/eolymp.discussion.PostService/ListPosts"
	PostService_CreatePost_FullMethodName              = "/eolymp.discussion.PostService/CreatePost"
	PostService_UpdatePost_FullMethodName              = "/eolymp.discussion.PostService/UpdatePost"
	PostService_PublishPost_FullMethodName             = "/eolymp.discussion.PostService/PublishPost"
	PostService_UnpublishPost_FullMethodName           = "/eolymp.discussion.PostService/UnpublishPost"
	PostService_ModeratePost_FullMethodName            = "/eolymp.discussion.PostService/ModeratePost"
	PostService_DeletePost_FullMethodName              = "/eolymp.discussion.PostService/DeletePost"
	PostService_VotePost_FullMethodName                = "/eolymp.discussion.PostService/VotePost"
	PostService_DescribePostTranslation_FullMethodName = "/eolymp.discussion.PostService/DescribePostTranslation"
	PostService_ListPostTranslations_FullMethodName    = "/eolymp.discussion.PostService/ListPostTranslations"
	PostService_CreatePostTranslation_FullMethodName   = "/eolymp.discussion.PostService/CreatePostTranslation"
	PostService_UpdatePostTranslation_FullMethodName   = "/eolymp.discussion.PostService/UpdatePostTranslation"
	PostService_DeletePostTranslation_FullMethodName   = "/eolymp.discussion.PostService/DeletePostTranslation"
)

// PostServiceClient is the client API for PostService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PostServiceClient interface {
	DescribePost(ctx context.Context, in *DescribePostInput, opts ...grpc.CallOption) (*DescribePostOutput, error)
	ListPosts(ctx context.Context, in *ListPostsInput, opts ...grpc.CallOption) (*ListPostsOutput, error)
	CreatePost(ctx context.Context, in *CreatePostInput, opts ...grpc.CallOption) (*CreatePostOutput, error)
	UpdatePost(ctx context.Context, in *UpdatePostInput, opts ...grpc.CallOption) (*UpdatePostOutput, error)
	PublishPost(ctx context.Context, in *PublishPostInput, opts ...grpc.CallOption) (*PublishPostOutput, error)
	UnpublishPost(ctx context.Context, in *UnpublishPostInput, opts ...grpc.CallOption) (*UnpublishPostOutput, error)
	ModeratePost(ctx context.Context, in *ModeratePostInput, opts ...grpc.CallOption) (*ModeratePostOutput, error)
	DeletePost(ctx context.Context, in *DeletePostInput, opts ...grpc.CallOption) (*DeletePostOutput, error)
	VotePost(ctx context.Context, in *VotePostInput, opts ...grpc.CallOption) (*VotePostOutput, error)
	DescribePostTranslation(ctx context.Context, in *DescribePostTranslationInput, opts ...grpc.CallOption) (*DescribePostTranslationOutput, error)
	ListPostTranslations(ctx context.Context, in *ListPostTranslationsInput, opts ...grpc.CallOption) (*ListPostTranslationsOutput, error)
	CreatePostTranslation(ctx context.Context, in *CreatePostTranslationInput, opts ...grpc.CallOption) (*CreatePostTranslationOutput, error)
	UpdatePostTranslation(ctx context.Context, in *UpdatePostTranslationInput, opts ...grpc.CallOption) (*UpdatePostTranslationOutput, error)
	DeletePostTranslation(ctx context.Context, in *DeletePostTranslationInput, opts ...grpc.CallOption) (*DeletePostTranslationOutput, error)
}

type postServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPostServiceClient(cc grpc.ClientConnInterface) PostServiceClient {
	return &postServiceClient{cc}
}

func (c *postServiceClient) DescribePost(ctx context.Context, in *DescribePostInput, opts ...grpc.CallOption) (*DescribePostOutput, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DescribePostOutput)
	err := c.cc.Invoke(ctx, PostService_DescribePost_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postServiceClient) ListPosts(ctx context.Context, in *ListPostsInput, opts ...grpc.CallOption) (*ListPostsOutput, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListPostsOutput)
	err := c.cc.Invoke(ctx, PostService_ListPosts_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postServiceClient) CreatePost(ctx context.Context, in *CreatePostInput, opts ...grpc.CallOption) (*CreatePostOutput, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreatePostOutput)
	err := c.cc.Invoke(ctx, PostService_CreatePost_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postServiceClient) UpdatePost(ctx context.Context, in *UpdatePostInput, opts ...grpc.CallOption) (*UpdatePostOutput, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdatePostOutput)
	err := c.cc.Invoke(ctx, PostService_UpdatePost_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postServiceClient) PublishPost(ctx context.Context, in *PublishPostInput, opts ...grpc.CallOption) (*PublishPostOutput, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PublishPostOutput)
	err := c.cc.Invoke(ctx, PostService_PublishPost_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postServiceClient) UnpublishPost(ctx context.Context, in *UnpublishPostInput, opts ...grpc.CallOption) (*UnpublishPostOutput, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UnpublishPostOutput)
	err := c.cc.Invoke(ctx, PostService_UnpublishPost_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postServiceClient) ModeratePost(ctx context.Context, in *ModeratePostInput, opts ...grpc.CallOption) (*ModeratePostOutput, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ModeratePostOutput)
	err := c.cc.Invoke(ctx, PostService_ModeratePost_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postServiceClient) DeletePost(ctx context.Context, in *DeletePostInput, opts ...grpc.CallOption) (*DeletePostOutput, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeletePostOutput)
	err := c.cc.Invoke(ctx, PostService_DeletePost_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postServiceClient) VotePost(ctx context.Context, in *VotePostInput, opts ...grpc.CallOption) (*VotePostOutput, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(VotePostOutput)
	err := c.cc.Invoke(ctx, PostService_VotePost_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postServiceClient) DescribePostTranslation(ctx context.Context, in *DescribePostTranslationInput, opts ...grpc.CallOption) (*DescribePostTranslationOutput, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DescribePostTranslationOutput)
	err := c.cc.Invoke(ctx, PostService_DescribePostTranslation_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postServiceClient) ListPostTranslations(ctx context.Context, in *ListPostTranslationsInput, opts ...grpc.CallOption) (*ListPostTranslationsOutput, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListPostTranslationsOutput)
	err := c.cc.Invoke(ctx, PostService_ListPostTranslations_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postServiceClient) CreatePostTranslation(ctx context.Context, in *CreatePostTranslationInput, opts ...grpc.CallOption) (*CreatePostTranslationOutput, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreatePostTranslationOutput)
	err := c.cc.Invoke(ctx, PostService_CreatePostTranslation_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postServiceClient) UpdatePostTranslation(ctx context.Context, in *UpdatePostTranslationInput, opts ...grpc.CallOption) (*UpdatePostTranslationOutput, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdatePostTranslationOutput)
	err := c.cc.Invoke(ctx, PostService_UpdatePostTranslation_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postServiceClient) DeletePostTranslation(ctx context.Context, in *DeletePostTranslationInput, opts ...grpc.CallOption) (*DeletePostTranslationOutput, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeletePostTranslationOutput)
	err := c.cc.Invoke(ctx, PostService_DeletePostTranslation_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PostServiceServer is the server API for PostService service.
// All implementations should embed UnimplementedPostServiceServer
// for forward compatibility.
type PostServiceServer interface {
	DescribePost(context.Context, *DescribePostInput) (*DescribePostOutput, error)
	ListPosts(context.Context, *ListPostsInput) (*ListPostsOutput, error)
	CreatePost(context.Context, *CreatePostInput) (*CreatePostOutput, error)
	UpdatePost(context.Context, *UpdatePostInput) (*UpdatePostOutput, error)
	PublishPost(context.Context, *PublishPostInput) (*PublishPostOutput, error)
	UnpublishPost(context.Context, *UnpublishPostInput) (*UnpublishPostOutput, error)
	ModeratePost(context.Context, *ModeratePostInput) (*ModeratePostOutput, error)
	DeletePost(context.Context, *DeletePostInput) (*DeletePostOutput, error)
	VotePost(context.Context, *VotePostInput) (*VotePostOutput, error)
	DescribePostTranslation(context.Context, *DescribePostTranslationInput) (*DescribePostTranslationOutput, error)
	ListPostTranslations(context.Context, *ListPostTranslationsInput) (*ListPostTranslationsOutput, error)
	CreatePostTranslation(context.Context, *CreatePostTranslationInput) (*CreatePostTranslationOutput, error)
	UpdatePostTranslation(context.Context, *UpdatePostTranslationInput) (*UpdatePostTranslationOutput, error)
	DeletePostTranslation(context.Context, *DeletePostTranslationInput) (*DeletePostTranslationOutput, error)
}

// UnimplementedPostServiceServer should be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedPostServiceServer struct{}

func (UnimplementedPostServiceServer) DescribePost(context.Context, *DescribePostInput) (*DescribePostOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DescribePost not implemented")
}
func (UnimplementedPostServiceServer) ListPosts(context.Context, *ListPostsInput) (*ListPostsOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListPosts not implemented")
}
func (UnimplementedPostServiceServer) CreatePost(context.Context, *CreatePostInput) (*CreatePostOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePost not implemented")
}
func (UnimplementedPostServiceServer) UpdatePost(context.Context, *UpdatePostInput) (*UpdatePostOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePost not implemented")
}
func (UnimplementedPostServiceServer) PublishPost(context.Context, *PublishPostInput) (*PublishPostOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PublishPost not implemented")
}
func (UnimplementedPostServiceServer) UnpublishPost(context.Context, *UnpublishPostInput) (*UnpublishPostOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnpublishPost not implemented")
}
func (UnimplementedPostServiceServer) ModeratePost(context.Context, *ModeratePostInput) (*ModeratePostOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ModeratePost not implemented")
}
func (UnimplementedPostServiceServer) DeletePost(context.Context, *DeletePostInput) (*DeletePostOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeletePost not implemented")
}
func (UnimplementedPostServiceServer) VotePost(context.Context, *VotePostInput) (*VotePostOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VotePost not implemented")
}
func (UnimplementedPostServiceServer) DescribePostTranslation(context.Context, *DescribePostTranslationInput) (*DescribePostTranslationOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DescribePostTranslation not implemented")
}
func (UnimplementedPostServiceServer) ListPostTranslations(context.Context, *ListPostTranslationsInput) (*ListPostTranslationsOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListPostTranslations not implemented")
}
func (UnimplementedPostServiceServer) CreatePostTranslation(context.Context, *CreatePostTranslationInput) (*CreatePostTranslationOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePostTranslation not implemented")
}
func (UnimplementedPostServiceServer) UpdatePostTranslation(context.Context, *UpdatePostTranslationInput) (*UpdatePostTranslationOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePostTranslation not implemented")
}
func (UnimplementedPostServiceServer) DeletePostTranslation(context.Context, *DeletePostTranslationInput) (*DeletePostTranslationOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeletePostTranslation not implemented")
}
func (UnimplementedPostServiceServer) testEmbeddedByValue() {}

// UnsafePostServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PostServiceServer will
// result in compilation errors.
type UnsafePostServiceServer interface {
	mustEmbedUnimplementedPostServiceServer()
}

func RegisterPostServiceServer(s grpc.ServiceRegistrar, srv PostServiceServer) {
	// If the following call pancis, it indicates UnimplementedPostServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&PostService_ServiceDesc, srv)
}

func _PostService_DescribePost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DescribePostInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostServiceServer).DescribePost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PostService_DescribePost_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostServiceServer).DescribePost(ctx, req.(*DescribePostInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _PostService_ListPosts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListPostsInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostServiceServer).ListPosts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PostService_ListPosts_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostServiceServer).ListPosts(ctx, req.(*ListPostsInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _PostService_CreatePost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePostInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostServiceServer).CreatePost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PostService_CreatePost_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostServiceServer).CreatePost(ctx, req.(*CreatePostInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _PostService_UpdatePost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePostInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostServiceServer).UpdatePost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PostService_UpdatePost_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostServiceServer).UpdatePost(ctx, req.(*UpdatePostInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _PostService_PublishPost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PublishPostInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostServiceServer).PublishPost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PostService_PublishPost_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostServiceServer).PublishPost(ctx, req.(*PublishPostInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _PostService_UnpublishPost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UnpublishPostInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostServiceServer).UnpublishPost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PostService_UnpublishPost_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostServiceServer).UnpublishPost(ctx, req.(*UnpublishPostInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _PostService_ModeratePost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ModeratePostInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostServiceServer).ModeratePost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PostService_ModeratePost_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostServiceServer).ModeratePost(ctx, req.(*ModeratePostInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _PostService_DeletePost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeletePostInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostServiceServer).DeletePost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PostService_DeletePost_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostServiceServer).DeletePost(ctx, req.(*DeletePostInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _PostService_VotePost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VotePostInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostServiceServer).VotePost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PostService_VotePost_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostServiceServer).VotePost(ctx, req.(*VotePostInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _PostService_DescribePostTranslation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DescribePostTranslationInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostServiceServer).DescribePostTranslation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PostService_DescribePostTranslation_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostServiceServer).DescribePostTranslation(ctx, req.(*DescribePostTranslationInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _PostService_ListPostTranslations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListPostTranslationsInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostServiceServer).ListPostTranslations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PostService_ListPostTranslations_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostServiceServer).ListPostTranslations(ctx, req.(*ListPostTranslationsInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _PostService_CreatePostTranslation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePostTranslationInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostServiceServer).CreatePostTranslation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PostService_CreatePostTranslation_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostServiceServer).CreatePostTranslation(ctx, req.(*CreatePostTranslationInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _PostService_UpdatePostTranslation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePostTranslationInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostServiceServer).UpdatePostTranslation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PostService_UpdatePostTranslation_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostServiceServer).UpdatePostTranslation(ctx, req.(*UpdatePostTranslationInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _PostService_DeletePostTranslation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeletePostTranslationInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostServiceServer).DeletePostTranslation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PostService_DeletePostTranslation_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostServiceServer).DeletePostTranslation(ctx, req.(*DeletePostTranslationInput))
	}
	return interceptor(ctx, in, info, handler)
}

// PostService_ServiceDesc is the grpc.ServiceDesc for PostService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PostService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "eolymp.discussion.PostService",
	HandlerType: (*PostServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DescribePost",
			Handler:    _PostService_DescribePost_Handler,
		},
		{
			MethodName: "ListPosts",
			Handler:    _PostService_ListPosts_Handler,
		},
		{
			MethodName: "CreatePost",
			Handler:    _PostService_CreatePost_Handler,
		},
		{
			MethodName: "UpdatePost",
			Handler:    _PostService_UpdatePost_Handler,
		},
		{
			MethodName: "PublishPost",
			Handler:    _PostService_PublishPost_Handler,
		},
		{
			MethodName: "UnpublishPost",
			Handler:    _PostService_UnpublishPost_Handler,
		},
		{
			MethodName: "ModeratePost",
			Handler:    _PostService_ModeratePost_Handler,
		},
		{
			MethodName: "DeletePost",
			Handler:    _PostService_DeletePost_Handler,
		},
		{
			MethodName: "VotePost",
			Handler:    _PostService_VotePost_Handler,
		},
		{
			MethodName: "DescribePostTranslation",
			Handler:    _PostService_DescribePostTranslation_Handler,
		},
		{
			MethodName: "ListPostTranslations",
			Handler:    _PostService_ListPostTranslations_Handler,
		},
		{
			MethodName: "CreatePostTranslation",
			Handler:    _PostService_CreatePostTranslation_Handler,
		},
		{
			MethodName: "UpdatePostTranslation",
			Handler:    _PostService_UpdatePostTranslation_Handler,
		},
		{
			MethodName: "DeletePostTranslation",
			Handler:    _PostService_DeletePostTranslation_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "eolymp/discussion/post_service.proto",
}
