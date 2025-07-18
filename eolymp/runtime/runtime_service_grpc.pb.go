// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.4
// source: eolymp/runtime/runtime_service.proto

package runtime

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
	RuntimeService_DescribeLanguage_FullMethodName     = "/eolymp.runtime.RuntimeService/DescribeLanguage"
	RuntimeService_ListLanguages_FullMethodName        = "/eolymp.runtime.RuntimeService/ListLanguages"
	RuntimeService_DescribeRuntime_FullMethodName      = "/eolymp.runtime.RuntimeService/DescribeRuntime"
	RuntimeService_ListRuntimes_FullMethodName         = "/eolymp.runtime.RuntimeService/ListRuntimes"
	RuntimeService_DescribeCodeTemplate_FullMethodName = "/eolymp.runtime.RuntimeService/DescribeCodeTemplate"
)

// RuntimeServiceClient is the client API for RuntimeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RuntimeServiceClient interface {
	DescribeLanguage(ctx context.Context, in *DescribeLanguageInput, opts ...grpc.CallOption) (*DescribeLanguageOutput, error)
	ListLanguages(ctx context.Context, in *ListLanguagesInput, opts ...grpc.CallOption) (*ListLanguagesOutput, error)
	DescribeRuntime(ctx context.Context, in *DescribeRuntimeInput, opts ...grpc.CallOption) (*DescribeRuntimeOutput, error)
	ListRuntimes(ctx context.Context, in *ListRuntimesInput, opts ...grpc.CallOption) (*ListRuntimesOutput, error)
	DescribeCodeTemplate(ctx context.Context, in *DescribeCodeTemplateInput, opts ...grpc.CallOption) (*DescribeCodeTemplateOutput, error)
}

type runtimeServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRuntimeServiceClient(cc grpc.ClientConnInterface) RuntimeServiceClient {
	return &runtimeServiceClient{cc}
}

func (c *runtimeServiceClient) DescribeLanguage(ctx context.Context, in *DescribeLanguageInput, opts ...grpc.CallOption) (*DescribeLanguageOutput, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DescribeLanguageOutput)
	err := c.cc.Invoke(ctx, RuntimeService_DescribeLanguage_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *runtimeServiceClient) ListLanguages(ctx context.Context, in *ListLanguagesInput, opts ...grpc.CallOption) (*ListLanguagesOutput, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListLanguagesOutput)
	err := c.cc.Invoke(ctx, RuntimeService_ListLanguages_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *runtimeServiceClient) DescribeRuntime(ctx context.Context, in *DescribeRuntimeInput, opts ...grpc.CallOption) (*DescribeRuntimeOutput, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DescribeRuntimeOutput)
	err := c.cc.Invoke(ctx, RuntimeService_DescribeRuntime_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *runtimeServiceClient) ListRuntimes(ctx context.Context, in *ListRuntimesInput, opts ...grpc.CallOption) (*ListRuntimesOutput, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListRuntimesOutput)
	err := c.cc.Invoke(ctx, RuntimeService_ListRuntimes_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *runtimeServiceClient) DescribeCodeTemplate(ctx context.Context, in *DescribeCodeTemplateInput, opts ...grpc.CallOption) (*DescribeCodeTemplateOutput, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DescribeCodeTemplateOutput)
	err := c.cc.Invoke(ctx, RuntimeService_DescribeCodeTemplate_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RuntimeServiceServer is the server API for RuntimeService service.
// All implementations should embed UnimplementedRuntimeServiceServer
// for forward compatibility.
type RuntimeServiceServer interface {
	DescribeLanguage(context.Context, *DescribeLanguageInput) (*DescribeLanguageOutput, error)
	ListLanguages(context.Context, *ListLanguagesInput) (*ListLanguagesOutput, error)
	DescribeRuntime(context.Context, *DescribeRuntimeInput) (*DescribeRuntimeOutput, error)
	ListRuntimes(context.Context, *ListRuntimesInput) (*ListRuntimesOutput, error)
	DescribeCodeTemplate(context.Context, *DescribeCodeTemplateInput) (*DescribeCodeTemplateOutput, error)
}

// UnimplementedRuntimeServiceServer should be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedRuntimeServiceServer struct{}

func (UnimplementedRuntimeServiceServer) DescribeLanguage(context.Context, *DescribeLanguageInput) (*DescribeLanguageOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DescribeLanguage not implemented")
}
func (UnimplementedRuntimeServiceServer) ListLanguages(context.Context, *ListLanguagesInput) (*ListLanguagesOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListLanguages not implemented")
}
func (UnimplementedRuntimeServiceServer) DescribeRuntime(context.Context, *DescribeRuntimeInput) (*DescribeRuntimeOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DescribeRuntime not implemented")
}
func (UnimplementedRuntimeServiceServer) ListRuntimes(context.Context, *ListRuntimesInput) (*ListRuntimesOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListRuntimes not implemented")
}
func (UnimplementedRuntimeServiceServer) DescribeCodeTemplate(context.Context, *DescribeCodeTemplateInput) (*DescribeCodeTemplateOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DescribeCodeTemplate not implemented")
}
func (UnimplementedRuntimeServiceServer) testEmbeddedByValue() {}

// UnsafeRuntimeServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RuntimeServiceServer will
// result in compilation errors.
type UnsafeRuntimeServiceServer interface {
	mustEmbedUnimplementedRuntimeServiceServer()
}

func RegisterRuntimeServiceServer(s grpc.ServiceRegistrar, srv RuntimeServiceServer) {
	// If the following call pancis, it indicates UnimplementedRuntimeServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&RuntimeService_ServiceDesc, srv)
}

func _RuntimeService_DescribeLanguage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DescribeLanguageInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RuntimeServiceServer).DescribeLanguage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RuntimeService_DescribeLanguage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RuntimeServiceServer).DescribeLanguage(ctx, req.(*DescribeLanguageInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _RuntimeService_ListLanguages_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListLanguagesInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RuntimeServiceServer).ListLanguages(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RuntimeService_ListLanguages_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RuntimeServiceServer).ListLanguages(ctx, req.(*ListLanguagesInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _RuntimeService_DescribeRuntime_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DescribeRuntimeInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RuntimeServiceServer).DescribeRuntime(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RuntimeService_DescribeRuntime_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RuntimeServiceServer).DescribeRuntime(ctx, req.(*DescribeRuntimeInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _RuntimeService_ListRuntimes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRuntimesInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RuntimeServiceServer).ListRuntimes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RuntimeService_ListRuntimes_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RuntimeServiceServer).ListRuntimes(ctx, req.(*ListRuntimesInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _RuntimeService_DescribeCodeTemplate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DescribeCodeTemplateInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RuntimeServiceServer).DescribeCodeTemplate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RuntimeService_DescribeCodeTemplate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RuntimeServiceServer).DescribeCodeTemplate(ctx, req.(*DescribeCodeTemplateInput))
	}
	return interceptor(ctx, in, info, handler)
}

// RuntimeService_ServiceDesc is the grpc.ServiceDesc for RuntimeService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RuntimeService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "eolymp.runtime.RuntimeService",
	HandlerType: (*RuntimeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DescribeLanguage",
			Handler:    _RuntimeService_DescribeLanguage_Handler,
		},
		{
			MethodName: "ListLanguages",
			Handler:    _RuntimeService_ListLanguages_Handler,
		},
		{
			MethodName: "DescribeRuntime",
			Handler:    _RuntimeService_DescribeRuntime_Handler,
		},
		{
			MethodName: "ListRuntimes",
			Handler:    _RuntimeService_ListRuntimes_Handler,
		},
		{
			MethodName: "DescribeCodeTemplate",
			Handler:    _RuntimeService_DescribeCodeTemplate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "eolymp/runtime/runtime_service.proto",
}
