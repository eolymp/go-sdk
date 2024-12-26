// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v4.24.4
// source: eolymp/discussion/configuration_service.proto

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
	ConfigurationService_DescribeDiscussionConfig_FullMethodName = "/eolymp.discussion.ConfigurationService/DescribeDiscussionConfig"
	ConfigurationService_UpdateDiscussionConfig_FullMethodName   = "/eolymp.discussion.ConfigurationService/UpdateDiscussionConfig"
)

// ConfigurationServiceClient is the client API for ConfigurationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ConfigurationServiceClient interface {
	// Describe identity provider configuration
	DescribeDiscussionConfig(ctx context.Context, in *DescribeDiscussionConfigInput, opts ...grpc.CallOption) (*DescribeDiscussionConfigOutput, error)
	// Update identity provider configuration
	UpdateDiscussionConfig(ctx context.Context, in *UpdateDiscussionConfigInput, opts ...grpc.CallOption) (*UpdateDiscussionConfigOutput, error)
}

type configurationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewConfigurationServiceClient(cc grpc.ClientConnInterface) ConfigurationServiceClient {
	return &configurationServiceClient{cc}
}

func (c *configurationServiceClient) DescribeDiscussionConfig(ctx context.Context, in *DescribeDiscussionConfigInput, opts ...grpc.CallOption) (*DescribeDiscussionConfigOutput, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DescribeDiscussionConfigOutput)
	err := c.cc.Invoke(ctx, ConfigurationService_DescribeDiscussionConfig_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *configurationServiceClient) UpdateDiscussionConfig(ctx context.Context, in *UpdateDiscussionConfigInput, opts ...grpc.CallOption) (*UpdateDiscussionConfigOutput, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateDiscussionConfigOutput)
	err := c.cc.Invoke(ctx, ConfigurationService_UpdateDiscussionConfig_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ConfigurationServiceServer is the server API for ConfigurationService service.
// All implementations should embed UnimplementedConfigurationServiceServer
// for forward compatibility.
type ConfigurationServiceServer interface {
	// Describe identity provider configuration
	DescribeDiscussionConfig(context.Context, *DescribeDiscussionConfigInput) (*DescribeDiscussionConfigOutput, error)
	// Update identity provider configuration
	UpdateDiscussionConfig(context.Context, *UpdateDiscussionConfigInput) (*UpdateDiscussionConfigOutput, error)
}

// UnimplementedConfigurationServiceServer should be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedConfigurationServiceServer struct{}

func (UnimplementedConfigurationServiceServer) DescribeDiscussionConfig(context.Context, *DescribeDiscussionConfigInput) (*DescribeDiscussionConfigOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DescribeDiscussionConfig not implemented")
}
func (UnimplementedConfigurationServiceServer) UpdateDiscussionConfig(context.Context, *UpdateDiscussionConfigInput) (*UpdateDiscussionConfigOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateDiscussionConfig not implemented")
}
func (UnimplementedConfigurationServiceServer) testEmbeddedByValue() {}

// UnsafeConfigurationServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ConfigurationServiceServer will
// result in compilation errors.
type UnsafeConfigurationServiceServer interface {
	mustEmbedUnimplementedConfigurationServiceServer()
}

func RegisterConfigurationServiceServer(s grpc.ServiceRegistrar, srv ConfigurationServiceServer) {
	// If the following call pancis, it indicates UnimplementedConfigurationServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&ConfigurationService_ServiceDesc, srv)
}

func _ConfigurationService_DescribeDiscussionConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DescribeDiscussionConfigInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConfigurationServiceServer).DescribeDiscussionConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ConfigurationService_DescribeDiscussionConfig_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConfigurationServiceServer).DescribeDiscussionConfig(ctx, req.(*DescribeDiscussionConfigInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConfigurationService_UpdateDiscussionConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateDiscussionConfigInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConfigurationServiceServer).UpdateDiscussionConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ConfigurationService_UpdateDiscussionConfig_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConfigurationServiceServer).UpdateDiscussionConfig(ctx, req.(*UpdateDiscussionConfigInput))
	}
	return interceptor(ctx, in, info, handler)
}

// ConfigurationService_ServiceDesc is the grpc.ServiceDesc for ConfigurationService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ConfigurationService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "eolymp.discussion.ConfigurationService",
	HandlerType: (*ConfigurationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DescribeDiscussionConfig",
			Handler:    _ConfigurationService_DescribeDiscussionConfig_Handler,
		},
		{
			MethodName: "UpdateDiscussionConfig",
			Handler:    _ConfigurationService_UpdateDiscussionConfig_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "eolymp/discussion/configuration_service.proto",
}
