// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.18.1
// source: eolymp/atlas/asset_service.proto

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
	AssetService_UploadFile_FullMethodName              = "/eolymp.atlas.AssetService/UploadFile"
	AssetService_StartMultipartUpload_FullMethodName    = "/eolymp.atlas.AssetService/StartMultipartUpload"
	AssetService_UploadPart_FullMethodName              = "/eolymp.atlas.AssetService/UploadPart"
	AssetService_CompleteMultipartUpload_FullMethodName = "/eolymp.atlas.AssetService/CompleteMultipartUpload"
)

// AssetServiceClient is the client API for AssetService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AssetServiceClient interface {
	UploadFile(ctx context.Context, in *UploadFileInput, opts ...grpc.CallOption) (*UploadFileOutput, error)
	StartMultipartUpload(ctx context.Context, in *StartMultipartUploadInput, opts ...grpc.CallOption) (*StartMultipartUploadOutput, error)
	UploadPart(ctx context.Context, in *UploadPartInput, opts ...grpc.CallOption) (*UploadPartOutput, error)
	CompleteMultipartUpload(ctx context.Context, in *CompleteMultipartUploadInput, opts ...grpc.CallOption) (*CompleteMultipartUploadOutput, error)
}

type assetServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAssetServiceClient(cc grpc.ClientConnInterface) AssetServiceClient {
	return &assetServiceClient{cc}
}

func (c *assetServiceClient) UploadFile(ctx context.Context, in *UploadFileInput, opts ...grpc.CallOption) (*UploadFileOutput, error) {
	out := new(UploadFileOutput)
	err := c.cc.Invoke(ctx, AssetService_UploadFile_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *assetServiceClient) StartMultipartUpload(ctx context.Context, in *StartMultipartUploadInput, opts ...grpc.CallOption) (*StartMultipartUploadOutput, error) {
	out := new(StartMultipartUploadOutput)
	err := c.cc.Invoke(ctx, AssetService_StartMultipartUpload_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *assetServiceClient) UploadPart(ctx context.Context, in *UploadPartInput, opts ...grpc.CallOption) (*UploadPartOutput, error) {
	out := new(UploadPartOutput)
	err := c.cc.Invoke(ctx, AssetService_UploadPart_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *assetServiceClient) CompleteMultipartUpload(ctx context.Context, in *CompleteMultipartUploadInput, opts ...grpc.CallOption) (*CompleteMultipartUploadOutput, error) {
	out := new(CompleteMultipartUploadOutput)
	err := c.cc.Invoke(ctx, AssetService_CompleteMultipartUpload_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AssetServiceServer is the server API for AssetService service.
// All implementations should embed UnimplementedAssetServiceServer
// for forward compatibility
type AssetServiceServer interface {
	UploadFile(context.Context, *UploadFileInput) (*UploadFileOutput, error)
	StartMultipartUpload(context.Context, *StartMultipartUploadInput) (*StartMultipartUploadOutput, error)
	UploadPart(context.Context, *UploadPartInput) (*UploadPartOutput, error)
	CompleteMultipartUpload(context.Context, *CompleteMultipartUploadInput) (*CompleteMultipartUploadOutput, error)
}

// UnimplementedAssetServiceServer should be embedded to have forward compatible implementations.
type UnimplementedAssetServiceServer struct {
}

func (UnimplementedAssetServiceServer) UploadFile(context.Context, *UploadFileInput) (*UploadFileOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UploadFile not implemented")
}
func (UnimplementedAssetServiceServer) StartMultipartUpload(context.Context, *StartMultipartUploadInput) (*StartMultipartUploadOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StartMultipartUpload not implemented")
}
func (UnimplementedAssetServiceServer) UploadPart(context.Context, *UploadPartInput) (*UploadPartOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UploadPart not implemented")
}
func (UnimplementedAssetServiceServer) CompleteMultipartUpload(context.Context, *CompleteMultipartUploadInput) (*CompleteMultipartUploadOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CompleteMultipartUpload not implemented")
}

// UnsafeAssetServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AssetServiceServer will
// result in compilation errors.
type UnsafeAssetServiceServer interface {
	mustEmbedUnimplementedAssetServiceServer()
}

func RegisterAssetServiceServer(s grpc.ServiceRegistrar, srv AssetServiceServer) {
	s.RegisterService(&AssetService_ServiceDesc, srv)
}

func _AssetService_UploadFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UploadFileInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AssetServiceServer).UploadFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AssetService_UploadFile_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AssetServiceServer).UploadFile(ctx, req.(*UploadFileInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _AssetService_StartMultipartUpload_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartMultipartUploadInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AssetServiceServer).StartMultipartUpload(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AssetService_StartMultipartUpload_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AssetServiceServer).StartMultipartUpload(ctx, req.(*StartMultipartUploadInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _AssetService_UploadPart_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UploadPartInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AssetServiceServer).UploadPart(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AssetService_UploadPart_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AssetServiceServer).UploadPart(ctx, req.(*UploadPartInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _AssetService_CompleteMultipartUpload_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CompleteMultipartUploadInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AssetServiceServer).CompleteMultipartUpload(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AssetService_CompleteMultipartUpload_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AssetServiceServer).CompleteMultipartUpload(ctx, req.(*CompleteMultipartUploadInput))
	}
	return interceptor(ctx, in, info, handler)
}

// AssetService_ServiceDesc is the grpc.ServiceDesc for AssetService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AssetService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "eolymp.atlas.AssetService",
	HandlerType: (*AssetServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UploadFile",
			Handler:    _AssetService_UploadFile_Handler,
		},
		{
			MethodName: "StartMultipartUpload",
			Handler:    _AssetService_StartMultipartUpload_Handler,
		},
		{
			MethodName: "UploadPart",
			Handler:    _AssetService_UploadPart_Handler,
		},
		{
			MethodName: "CompleteMultipartUpload",
			Handler:    _AssetService_CompleteMultipartUpload_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "eolymp/atlas/asset_service.proto",
}
