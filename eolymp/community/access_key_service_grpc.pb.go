// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.4
// source: eolymp/community/access_key_service.proto

package community

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
	AccessKeyService_CreateAccessKey_FullMethodName = "/eolymp.community.AccessKeyService/CreateAccessKey"
	AccessKeyService_DeleteAccessKey_FullMethodName = "/eolymp.community.AccessKeyService/DeleteAccessKey"
	AccessKeyService_ListAccessKeys_FullMethodName  = "/eolymp.community.AccessKeyService/ListAccessKeys"
)

// AccessKeyServiceClient is the client API for AccessKeyService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AccessKeyServiceClient interface {
	// Create API key.
	CreateAccessKey(ctx context.Context, in *CreateAccessKeyInput, opts ...grpc.CallOption) (*CreateAccessKeyOutput, error)
	// Delete API key.
	DeleteAccessKey(ctx context.Context, in *DeleteAccessKeyInput, opts ...grpc.CallOption) (*DeleteAccessKeyOutput, error)
	ListAccessKeys(ctx context.Context, in *ListAccessKeysInput, opts ...grpc.CallOption) (*ListAccessKeysOutput, error)
}

type accessKeyServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAccessKeyServiceClient(cc grpc.ClientConnInterface) AccessKeyServiceClient {
	return &accessKeyServiceClient{cc}
}

func (c *accessKeyServiceClient) CreateAccessKey(ctx context.Context, in *CreateAccessKeyInput, opts ...grpc.CallOption) (*CreateAccessKeyOutput, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateAccessKeyOutput)
	err := c.cc.Invoke(ctx, AccessKeyService_CreateAccessKey_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accessKeyServiceClient) DeleteAccessKey(ctx context.Context, in *DeleteAccessKeyInput, opts ...grpc.CallOption) (*DeleteAccessKeyOutput, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteAccessKeyOutput)
	err := c.cc.Invoke(ctx, AccessKeyService_DeleteAccessKey_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accessKeyServiceClient) ListAccessKeys(ctx context.Context, in *ListAccessKeysInput, opts ...grpc.CallOption) (*ListAccessKeysOutput, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListAccessKeysOutput)
	err := c.cc.Invoke(ctx, AccessKeyService_ListAccessKeys_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AccessKeyServiceServer is the server API for AccessKeyService service.
// All implementations should embed UnimplementedAccessKeyServiceServer
// for forward compatibility.
type AccessKeyServiceServer interface {
	// Create API key.
	CreateAccessKey(context.Context, *CreateAccessKeyInput) (*CreateAccessKeyOutput, error)
	// Delete API key.
	DeleteAccessKey(context.Context, *DeleteAccessKeyInput) (*DeleteAccessKeyOutput, error)
	ListAccessKeys(context.Context, *ListAccessKeysInput) (*ListAccessKeysOutput, error)
}

// UnimplementedAccessKeyServiceServer should be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedAccessKeyServiceServer struct{}

func (UnimplementedAccessKeyServiceServer) CreateAccessKey(context.Context, *CreateAccessKeyInput) (*CreateAccessKeyOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAccessKey not implemented")
}
func (UnimplementedAccessKeyServiceServer) DeleteAccessKey(context.Context, *DeleteAccessKeyInput) (*DeleteAccessKeyOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteAccessKey not implemented")
}
func (UnimplementedAccessKeyServiceServer) ListAccessKeys(context.Context, *ListAccessKeysInput) (*ListAccessKeysOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListAccessKeys not implemented")
}
func (UnimplementedAccessKeyServiceServer) testEmbeddedByValue() {}

// UnsafeAccessKeyServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AccessKeyServiceServer will
// result in compilation errors.
type UnsafeAccessKeyServiceServer interface {
	mustEmbedUnimplementedAccessKeyServiceServer()
}

func RegisterAccessKeyServiceServer(s grpc.ServiceRegistrar, srv AccessKeyServiceServer) {
	// If the following call pancis, it indicates UnimplementedAccessKeyServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&AccessKeyService_ServiceDesc, srv)
}

func _AccessKeyService_CreateAccessKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAccessKeyInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccessKeyServiceServer).CreateAccessKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AccessKeyService_CreateAccessKey_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccessKeyServiceServer).CreateAccessKey(ctx, req.(*CreateAccessKeyInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccessKeyService_DeleteAccessKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteAccessKeyInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccessKeyServiceServer).DeleteAccessKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AccessKeyService_DeleteAccessKey_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccessKeyServiceServer).DeleteAccessKey(ctx, req.(*DeleteAccessKeyInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccessKeyService_ListAccessKeys_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListAccessKeysInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccessKeyServiceServer).ListAccessKeys(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AccessKeyService_ListAccessKeys_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccessKeyServiceServer).ListAccessKeys(ctx, req.(*ListAccessKeysInput))
	}
	return interceptor(ctx, in, info, handler)
}

// AccessKeyService_ServiceDesc is the grpc.ServiceDesc for AccessKeyService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AccessKeyService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "eolymp.community.AccessKeyService",
	HandlerType: (*AccessKeyServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateAccessKey",
			Handler:    _AccessKeyService_CreateAccessKey_Handler,
		},
		{
			MethodName: "DeleteAccessKey",
			Handler:    _AccessKeyService_DeleteAccessKey_Handler,
		},
		{
			MethodName: "ListAccessKeys",
			Handler:    _AccessKeyService_ListAccessKeys_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "eolymp/community/access_key_service.proto",
}
