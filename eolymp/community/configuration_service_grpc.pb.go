// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v4.24.4
// source: eolymp/community/configuration_service.proto

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
	ConfigurationService_DescribeIdentityConfig_FullMethodName  = "/eolymp.community.ConfigurationService/DescribeIdentityConfig"
	ConfigurationService_ConfigureIdentityConfig_FullMethodName = "/eolymp.community.ConfigurationService/ConfigureIdentityConfig"
)

// ConfigurationServiceClient is the client API for ConfigurationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ConfigurationServiceClient interface {
	// Describe identity provider configuration
	DescribeIdentityConfig(ctx context.Context, in *DescribeIdentityConfigInput, opts ...grpc.CallOption) (*DescribeIdentityConfigOutput, error)
	// Update identity provider configuration
	ConfigureIdentityConfig(ctx context.Context, in *ConfigureIdentityConfigInput, opts ...grpc.CallOption) (*ConfigureIdentityConfigOutput, error)
}

type configurationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewConfigurationServiceClient(cc grpc.ClientConnInterface) ConfigurationServiceClient {
	return &configurationServiceClient{cc}
}

func (c *configurationServiceClient) DescribeIdentityConfig(ctx context.Context, in *DescribeIdentityConfigInput, opts ...grpc.CallOption) (*DescribeIdentityConfigOutput, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DescribeIdentityConfigOutput)
	err := c.cc.Invoke(ctx, ConfigurationService_DescribeIdentityConfig_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *configurationServiceClient) ConfigureIdentityConfig(ctx context.Context, in *ConfigureIdentityConfigInput, opts ...grpc.CallOption) (*ConfigureIdentityConfigOutput, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ConfigureIdentityConfigOutput)
	err := c.cc.Invoke(ctx, ConfigurationService_ConfigureIdentityConfig_FullMethodName, in, out, cOpts...)
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
	DescribeIdentityConfig(context.Context, *DescribeIdentityConfigInput) (*DescribeIdentityConfigOutput, error)
	// Update identity provider configuration
	ConfigureIdentityConfig(context.Context, *ConfigureIdentityConfigInput) (*ConfigureIdentityConfigOutput, error)
}

// UnimplementedConfigurationServiceServer should be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedConfigurationServiceServer struct{}

func (UnimplementedConfigurationServiceServer) DescribeIdentityConfig(context.Context, *DescribeIdentityConfigInput) (*DescribeIdentityConfigOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DescribeIdentityConfig not implemented")
}
func (UnimplementedConfigurationServiceServer) ConfigureIdentityConfig(context.Context, *ConfigureIdentityConfigInput) (*ConfigureIdentityConfigOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConfigureIdentityConfig not implemented")
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

func _ConfigurationService_DescribeIdentityConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DescribeIdentityConfigInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConfigurationServiceServer).DescribeIdentityConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ConfigurationService_DescribeIdentityConfig_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConfigurationServiceServer).DescribeIdentityConfig(ctx, req.(*DescribeIdentityConfigInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConfigurationService_ConfigureIdentityConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConfigureIdentityConfigInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConfigurationServiceServer).ConfigureIdentityConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ConfigurationService_ConfigureIdentityConfig_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConfigurationServiceServer).ConfigureIdentityConfig(ctx, req.(*ConfigureIdentityConfigInput))
	}
	return interceptor(ctx, in, info, handler)
}

// ConfigurationService_ServiceDesc is the grpc.ServiceDesc for ConfigurationService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ConfigurationService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "eolymp.community.ConfigurationService",
	HandlerType: (*ConfigurationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DescribeIdentityConfig",
			Handler:    _ConfigurationService_DescribeIdentityConfig_Handler,
		},
		{
			MethodName: "ConfigureIdentityConfig",
			Handler:    _ConfigurationService_ConfigureIdentityConfig_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "eolymp/community/configuration_service.proto",
}
