// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.2
// source: eolymp/helpdesk/helpdesk.proto

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

// HelpdeskClient is the client API for Helpdesk service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HelpdeskClient interface {
	DescribeDocument(ctx context.Context, in *DescribeDocumentInput, opts ...grpc.CallOption) (*DescribeDocumentOutput, error)
	ListDocuments(ctx context.Context, in *ListDocumentsInput, opts ...grpc.CallOption) (*ListDocumentsOutput, error)
	CreateDocument(ctx context.Context, in *CreateDocumentInput, opts ...grpc.CallOption) (*CreateDocumentOutput, error)
	UpdateDocument(ctx context.Context, in *UpdateDocumentInput, opts ...grpc.CallOption) (*UpdateDocumentOutput, error)
	DeleteDocument(ctx context.Context, in *DeleteDocumentInput, opts ...grpc.CallOption) (*DeleteDocumentOutput, error)
	DescribePath(ctx context.Context, in *DescribePathInput, opts ...grpc.CallOption) (*DescribePathOutput, error)
	ListPaths(ctx context.Context, in *ListPathsInput, opts ...grpc.CallOption) (*ListPathsOutput, error)
	ListParents(ctx context.Context, in *ListParentsInput, opts ...grpc.CallOption) (*ListParentsOutput, error)
}

type helpdeskClient struct {
	cc grpc.ClientConnInterface
}

func NewHelpdeskClient(cc grpc.ClientConnInterface) HelpdeskClient {
	return &helpdeskClient{cc}
}

func (c *helpdeskClient) DescribeDocument(ctx context.Context, in *DescribeDocumentInput, opts ...grpc.CallOption) (*DescribeDocumentOutput, error) {
	out := new(DescribeDocumentOutput)
	err := c.cc.Invoke(ctx, "/eolymp.helpdesk.Helpdesk/DescribeDocument", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *helpdeskClient) ListDocuments(ctx context.Context, in *ListDocumentsInput, opts ...grpc.CallOption) (*ListDocumentsOutput, error) {
	out := new(ListDocumentsOutput)
	err := c.cc.Invoke(ctx, "/eolymp.helpdesk.Helpdesk/ListDocuments", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *helpdeskClient) CreateDocument(ctx context.Context, in *CreateDocumentInput, opts ...grpc.CallOption) (*CreateDocumentOutput, error) {
	out := new(CreateDocumentOutput)
	err := c.cc.Invoke(ctx, "/eolymp.helpdesk.Helpdesk/CreateDocument", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *helpdeskClient) UpdateDocument(ctx context.Context, in *UpdateDocumentInput, opts ...grpc.CallOption) (*UpdateDocumentOutput, error) {
	out := new(UpdateDocumentOutput)
	err := c.cc.Invoke(ctx, "/eolymp.helpdesk.Helpdesk/UpdateDocument", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *helpdeskClient) DeleteDocument(ctx context.Context, in *DeleteDocumentInput, opts ...grpc.CallOption) (*DeleteDocumentOutput, error) {
	out := new(DeleteDocumentOutput)
	err := c.cc.Invoke(ctx, "/eolymp.helpdesk.Helpdesk/DeleteDocument", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *helpdeskClient) DescribePath(ctx context.Context, in *DescribePathInput, opts ...grpc.CallOption) (*DescribePathOutput, error) {
	out := new(DescribePathOutput)
	err := c.cc.Invoke(ctx, "/eolymp.helpdesk.Helpdesk/DescribePath", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *helpdeskClient) ListPaths(ctx context.Context, in *ListPathsInput, opts ...grpc.CallOption) (*ListPathsOutput, error) {
	out := new(ListPathsOutput)
	err := c.cc.Invoke(ctx, "/eolymp.helpdesk.Helpdesk/ListPaths", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *helpdeskClient) ListParents(ctx context.Context, in *ListParentsInput, opts ...grpc.CallOption) (*ListParentsOutput, error) {
	out := new(ListParentsOutput)
	err := c.cc.Invoke(ctx, "/eolymp.helpdesk.Helpdesk/ListParents", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HelpdeskServer is the server API for Helpdesk service.
// All implementations must embed UnimplementedHelpdeskServer
// for forward compatibility
type HelpdeskServer interface {
	DescribeDocument(context.Context, *DescribeDocumentInput) (*DescribeDocumentOutput, error)
	ListDocuments(context.Context, *ListDocumentsInput) (*ListDocumentsOutput, error)
	CreateDocument(context.Context, *CreateDocumentInput) (*CreateDocumentOutput, error)
	UpdateDocument(context.Context, *UpdateDocumentInput) (*UpdateDocumentOutput, error)
	DeleteDocument(context.Context, *DeleteDocumentInput) (*DeleteDocumentOutput, error)
	DescribePath(context.Context, *DescribePathInput) (*DescribePathOutput, error)
	ListPaths(context.Context, *ListPathsInput) (*ListPathsOutput, error)
	ListParents(context.Context, *ListParentsInput) (*ListParentsOutput, error)
	mustEmbedUnimplementedHelpdeskServer()
}

// UnimplementedHelpdeskServer must be embedded to have forward compatible implementations.
type UnimplementedHelpdeskServer struct {
}

func (UnimplementedHelpdeskServer) DescribeDocument(context.Context, *DescribeDocumentInput) (*DescribeDocumentOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DescribeDocument not implemented")
}
func (UnimplementedHelpdeskServer) ListDocuments(context.Context, *ListDocumentsInput) (*ListDocumentsOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListDocuments not implemented")
}
func (UnimplementedHelpdeskServer) CreateDocument(context.Context, *CreateDocumentInput) (*CreateDocumentOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateDocument not implemented")
}
func (UnimplementedHelpdeskServer) UpdateDocument(context.Context, *UpdateDocumentInput) (*UpdateDocumentOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateDocument not implemented")
}
func (UnimplementedHelpdeskServer) DeleteDocument(context.Context, *DeleteDocumentInput) (*DeleteDocumentOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteDocument not implemented")
}
func (UnimplementedHelpdeskServer) DescribePath(context.Context, *DescribePathInput) (*DescribePathOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DescribePath not implemented")
}
func (UnimplementedHelpdeskServer) ListPaths(context.Context, *ListPathsInput) (*ListPathsOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListPaths not implemented")
}
func (UnimplementedHelpdeskServer) ListParents(context.Context, *ListParentsInput) (*ListParentsOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListParents not implemented")
}
func (UnimplementedHelpdeskServer) mustEmbedUnimplementedHelpdeskServer() {}

// UnsafeHelpdeskServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HelpdeskServer will
// result in compilation errors.
type UnsafeHelpdeskServer interface {
	mustEmbedUnimplementedHelpdeskServer()
}

func RegisterHelpdeskServer(s grpc.ServiceRegistrar, srv HelpdeskServer) {
	s.RegisterService(&Helpdesk_ServiceDesc, srv)
}

func _Helpdesk_DescribeDocument_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DescribeDocumentInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HelpdeskServer).DescribeDocument(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/eolymp.helpdesk.Helpdesk/DescribeDocument",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HelpdeskServer).DescribeDocument(ctx, req.(*DescribeDocumentInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _Helpdesk_ListDocuments_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListDocumentsInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HelpdeskServer).ListDocuments(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/eolymp.helpdesk.Helpdesk/ListDocuments",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HelpdeskServer).ListDocuments(ctx, req.(*ListDocumentsInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _Helpdesk_CreateDocument_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateDocumentInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HelpdeskServer).CreateDocument(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/eolymp.helpdesk.Helpdesk/CreateDocument",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HelpdeskServer).CreateDocument(ctx, req.(*CreateDocumentInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _Helpdesk_UpdateDocument_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateDocumentInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HelpdeskServer).UpdateDocument(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/eolymp.helpdesk.Helpdesk/UpdateDocument",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HelpdeskServer).UpdateDocument(ctx, req.(*UpdateDocumentInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _Helpdesk_DeleteDocument_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteDocumentInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HelpdeskServer).DeleteDocument(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/eolymp.helpdesk.Helpdesk/DeleteDocument",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HelpdeskServer).DeleteDocument(ctx, req.(*DeleteDocumentInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _Helpdesk_DescribePath_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DescribePathInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HelpdeskServer).DescribePath(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/eolymp.helpdesk.Helpdesk/DescribePath",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HelpdeskServer).DescribePath(ctx, req.(*DescribePathInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _Helpdesk_ListPaths_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListPathsInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HelpdeskServer).ListPaths(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/eolymp.helpdesk.Helpdesk/ListPaths",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HelpdeskServer).ListPaths(ctx, req.(*ListPathsInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _Helpdesk_ListParents_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListParentsInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HelpdeskServer).ListParents(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/eolymp.helpdesk.Helpdesk/ListParents",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HelpdeskServer).ListParents(ctx, req.(*ListParentsInput))
	}
	return interceptor(ctx, in, info, handler)
}

// Helpdesk_ServiceDesc is the grpc.ServiceDesc for Helpdesk service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Helpdesk_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "eolymp.helpdesk.Helpdesk",
	HandlerType: (*HelpdeskServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DescribeDocument",
			Handler:    _Helpdesk_DescribeDocument_Handler,
		},
		{
			MethodName: "ListDocuments",
			Handler:    _Helpdesk_ListDocuments_Handler,
		},
		{
			MethodName: "CreateDocument",
			Handler:    _Helpdesk_CreateDocument_Handler,
		},
		{
			MethodName: "UpdateDocument",
			Handler:    _Helpdesk_UpdateDocument_Handler,
		},
		{
			MethodName: "DeleteDocument",
			Handler:    _Helpdesk_DeleteDocument_Handler,
		},
		{
			MethodName: "DescribePath",
			Handler:    _Helpdesk_DescribePath_Handler,
		},
		{
			MethodName: "ListPaths",
			Handler:    _Helpdesk_ListPaths_Handler,
		},
		{
			MethodName: "ListParents",
			Handler:    _Helpdesk_ListParents_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "eolymp/helpdesk/helpdesk.proto",
}
