// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.4
// source: eolymp/atlas/attachment_service.proto

package atlas

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
	AttachmentService_CreateAttachment_FullMethodName   = "/eolymp.atlas.AttachmentService/CreateAttachment"
	AttachmentService_UpdateAttachment_FullMethodName   = "/eolymp.atlas.AttachmentService/UpdateAttachment"
	AttachmentService_DeleteAttachment_FullMethodName   = "/eolymp.atlas.AttachmentService/DeleteAttachment"
	AttachmentService_ListAttachments_FullMethodName    = "/eolymp.atlas.AttachmentService/ListAttachments"
	AttachmentService_DescribeAttachment_FullMethodName = "/eolymp.atlas.AttachmentService/DescribeAttachment"
)

// AttachmentServiceClient is the client API for AttachmentService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AttachmentServiceClient interface {
	CreateAttachment(ctx context.Context, in *CreateAttachmentInput, opts ...grpc.CallOption) (*CreateAttachmentOutput, error)
	UpdateAttachment(ctx context.Context, in *UpdateAttachmentInput, opts ...grpc.CallOption) (*UpdateAttachmentOutput, error)
	DeleteAttachment(ctx context.Context, in *DeleteAttachmentInput, opts ...grpc.CallOption) (*DeleteAttachmentOutput, error)
	ListAttachments(ctx context.Context, in *ListAttachmentsInput, opts ...grpc.CallOption) (*ListAttachmentsOutput, error)
	DescribeAttachment(ctx context.Context, in *DescribeAttachmentInput, opts ...grpc.CallOption) (*DescribeAttachmentOutput, error)
}

type attachmentServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAttachmentServiceClient(cc grpc.ClientConnInterface) AttachmentServiceClient {
	return &attachmentServiceClient{cc}
}

func (c *attachmentServiceClient) CreateAttachment(ctx context.Context, in *CreateAttachmentInput, opts ...grpc.CallOption) (*CreateAttachmentOutput, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateAttachmentOutput)
	err := c.cc.Invoke(ctx, AttachmentService_CreateAttachment_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *attachmentServiceClient) UpdateAttachment(ctx context.Context, in *UpdateAttachmentInput, opts ...grpc.CallOption) (*UpdateAttachmentOutput, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateAttachmentOutput)
	err := c.cc.Invoke(ctx, AttachmentService_UpdateAttachment_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *attachmentServiceClient) DeleteAttachment(ctx context.Context, in *DeleteAttachmentInput, opts ...grpc.CallOption) (*DeleteAttachmentOutput, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteAttachmentOutput)
	err := c.cc.Invoke(ctx, AttachmentService_DeleteAttachment_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *attachmentServiceClient) ListAttachments(ctx context.Context, in *ListAttachmentsInput, opts ...grpc.CallOption) (*ListAttachmentsOutput, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListAttachmentsOutput)
	err := c.cc.Invoke(ctx, AttachmentService_ListAttachments_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *attachmentServiceClient) DescribeAttachment(ctx context.Context, in *DescribeAttachmentInput, opts ...grpc.CallOption) (*DescribeAttachmentOutput, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DescribeAttachmentOutput)
	err := c.cc.Invoke(ctx, AttachmentService_DescribeAttachment_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AttachmentServiceServer is the server API for AttachmentService service.
// All implementations should embed UnimplementedAttachmentServiceServer
// for forward compatibility.
type AttachmentServiceServer interface {
	CreateAttachment(context.Context, *CreateAttachmentInput) (*CreateAttachmentOutput, error)
	UpdateAttachment(context.Context, *UpdateAttachmentInput) (*UpdateAttachmentOutput, error)
	DeleteAttachment(context.Context, *DeleteAttachmentInput) (*DeleteAttachmentOutput, error)
	ListAttachments(context.Context, *ListAttachmentsInput) (*ListAttachmentsOutput, error)
	DescribeAttachment(context.Context, *DescribeAttachmentInput) (*DescribeAttachmentOutput, error)
}

// UnimplementedAttachmentServiceServer should be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedAttachmentServiceServer struct{}

func (UnimplementedAttachmentServiceServer) CreateAttachment(context.Context, *CreateAttachmentInput) (*CreateAttachmentOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAttachment not implemented")
}
func (UnimplementedAttachmentServiceServer) UpdateAttachment(context.Context, *UpdateAttachmentInput) (*UpdateAttachmentOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAttachment not implemented")
}
func (UnimplementedAttachmentServiceServer) DeleteAttachment(context.Context, *DeleteAttachmentInput) (*DeleteAttachmentOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteAttachment not implemented")
}
func (UnimplementedAttachmentServiceServer) ListAttachments(context.Context, *ListAttachmentsInput) (*ListAttachmentsOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListAttachments not implemented")
}
func (UnimplementedAttachmentServiceServer) DescribeAttachment(context.Context, *DescribeAttachmentInput) (*DescribeAttachmentOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DescribeAttachment not implemented")
}
func (UnimplementedAttachmentServiceServer) testEmbeddedByValue() {}

// UnsafeAttachmentServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AttachmentServiceServer will
// result in compilation errors.
type UnsafeAttachmentServiceServer interface {
	mustEmbedUnimplementedAttachmentServiceServer()
}

func RegisterAttachmentServiceServer(s grpc.ServiceRegistrar, srv AttachmentServiceServer) {
	// If the following call pancis, it indicates UnimplementedAttachmentServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&AttachmentService_ServiceDesc, srv)
}

func _AttachmentService_CreateAttachment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAttachmentInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AttachmentServiceServer).CreateAttachment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AttachmentService_CreateAttachment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AttachmentServiceServer).CreateAttachment(ctx, req.(*CreateAttachmentInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _AttachmentService_UpdateAttachment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateAttachmentInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AttachmentServiceServer).UpdateAttachment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AttachmentService_UpdateAttachment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AttachmentServiceServer).UpdateAttachment(ctx, req.(*UpdateAttachmentInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _AttachmentService_DeleteAttachment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteAttachmentInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AttachmentServiceServer).DeleteAttachment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AttachmentService_DeleteAttachment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AttachmentServiceServer).DeleteAttachment(ctx, req.(*DeleteAttachmentInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _AttachmentService_ListAttachments_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListAttachmentsInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AttachmentServiceServer).ListAttachments(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AttachmentService_ListAttachments_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AttachmentServiceServer).ListAttachments(ctx, req.(*ListAttachmentsInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _AttachmentService_DescribeAttachment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DescribeAttachmentInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AttachmentServiceServer).DescribeAttachment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AttachmentService_DescribeAttachment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AttachmentServiceServer).DescribeAttachment(ctx, req.(*DescribeAttachmentInput))
	}
	return interceptor(ctx, in, info, handler)
}

// AttachmentService_ServiceDesc is the grpc.ServiceDesc for AttachmentService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AttachmentService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "eolymp.atlas.AttachmentService",
	HandlerType: (*AttachmentServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateAttachment",
			Handler:    _AttachmentService_CreateAttachment_Handler,
		},
		{
			MethodName: "UpdateAttachment",
			Handler:    _AttachmentService_UpdateAttachment_Handler,
		},
		{
			MethodName: "DeleteAttachment",
			Handler:    _AttachmentService_DeleteAttachment_Handler,
		},
		{
			MethodName: "ListAttachments",
			Handler:    _AttachmentService_ListAttachments_Handler,
		},
		{
			MethodName: "DescribeAttachment",
			Handler:    _AttachmentService_DescribeAttachment_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "eolymp/atlas/attachment_service.proto",
}
