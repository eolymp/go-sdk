// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.2
// source: eolymp/executor/executor.proto

package executor

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

// ExecutorClient is the client API for Executor service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ExecutorClient interface {
	DescribeLanguage(ctx context.Context, in *DescribeLanguageInput, opts ...grpc.CallOption) (*DescribeLanguageOutput, error)
	ListLanguages(ctx context.Context, in *ListLanguagesInput, opts ...grpc.CallOption) (*ListLanguagesOutput, error)
	DescribeRuntime(ctx context.Context, in *DescribeRuntimeInput, opts ...grpc.CallOption) (*DescribeRuntimeOutput, error)
	ListRuntime(ctx context.Context, in *ListRuntimeInput, opts ...grpc.CallOption) (*ListRuntimeOutput, error)
	DescribeCodeTemplate(ctx context.Context, in *DescribeCodeTemplateInput, opts ...grpc.CallOption) (*DescribeCodeTemplateOutput, error)
}

type executorClient struct {
	cc grpc.ClientConnInterface
}

func NewExecutorClient(cc grpc.ClientConnInterface) ExecutorClient {
	return &executorClient{cc}
}

func (c *executorClient) DescribeLanguage(ctx context.Context, in *DescribeLanguageInput, opts ...grpc.CallOption) (*DescribeLanguageOutput, error) {
	out := new(DescribeLanguageOutput)
	err := c.cc.Invoke(ctx, "/eolymp.executor.Executor/DescribeLanguage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *executorClient) ListLanguages(ctx context.Context, in *ListLanguagesInput, opts ...grpc.CallOption) (*ListLanguagesOutput, error) {
	out := new(ListLanguagesOutput)
	err := c.cc.Invoke(ctx, "/eolymp.executor.Executor/ListLanguages", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *executorClient) DescribeRuntime(ctx context.Context, in *DescribeRuntimeInput, opts ...grpc.CallOption) (*DescribeRuntimeOutput, error) {
	out := new(DescribeRuntimeOutput)
	err := c.cc.Invoke(ctx, "/eolymp.executor.Executor/DescribeRuntime", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *executorClient) ListRuntime(ctx context.Context, in *ListRuntimeInput, opts ...grpc.CallOption) (*ListRuntimeOutput, error) {
	out := new(ListRuntimeOutput)
	err := c.cc.Invoke(ctx, "/eolymp.executor.Executor/ListRuntime", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *executorClient) DescribeCodeTemplate(ctx context.Context, in *DescribeCodeTemplateInput, opts ...grpc.CallOption) (*DescribeCodeTemplateOutput, error) {
	out := new(DescribeCodeTemplateOutput)
	err := c.cc.Invoke(ctx, "/eolymp.executor.Executor/DescribeCodeTemplate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ExecutorServer is the server API for Executor service.
// All implementations must embed UnimplementedExecutorServer
// for forward compatibility
type ExecutorServer interface {
	DescribeLanguage(context.Context, *DescribeLanguageInput) (*DescribeLanguageOutput, error)
	ListLanguages(context.Context, *ListLanguagesInput) (*ListLanguagesOutput, error)
	DescribeRuntime(context.Context, *DescribeRuntimeInput) (*DescribeRuntimeOutput, error)
	ListRuntime(context.Context, *ListRuntimeInput) (*ListRuntimeOutput, error)
	DescribeCodeTemplate(context.Context, *DescribeCodeTemplateInput) (*DescribeCodeTemplateOutput, error)
	mustEmbedUnimplementedExecutorServer()
}

// UnimplementedExecutorServer must be embedded to have forward compatible implementations.
type UnimplementedExecutorServer struct {
}

func (UnimplementedExecutorServer) DescribeLanguage(context.Context, *DescribeLanguageInput) (*DescribeLanguageOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DescribeLanguage not implemented")
}
func (UnimplementedExecutorServer) ListLanguages(context.Context, *ListLanguagesInput) (*ListLanguagesOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListLanguages not implemented")
}
func (UnimplementedExecutorServer) DescribeRuntime(context.Context, *DescribeRuntimeInput) (*DescribeRuntimeOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DescribeRuntime not implemented")
}
func (UnimplementedExecutorServer) ListRuntime(context.Context, *ListRuntimeInput) (*ListRuntimeOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListRuntime not implemented")
}
func (UnimplementedExecutorServer) DescribeCodeTemplate(context.Context, *DescribeCodeTemplateInput) (*DescribeCodeTemplateOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DescribeCodeTemplate not implemented")
}
func (UnimplementedExecutorServer) mustEmbedUnimplementedExecutorServer() {}

// UnsafeExecutorServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ExecutorServer will
// result in compilation errors.
type UnsafeExecutorServer interface {
	mustEmbedUnimplementedExecutorServer()
}

func RegisterExecutorServer(s grpc.ServiceRegistrar, srv ExecutorServer) {
	s.RegisterService(&Executor_ServiceDesc, srv)
}

func _Executor_DescribeLanguage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DescribeLanguageInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExecutorServer).DescribeLanguage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/eolymp.executor.Executor/DescribeLanguage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExecutorServer).DescribeLanguage(ctx, req.(*DescribeLanguageInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _Executor_ListLanguages_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListLanguagesInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExecutorServer).ListLanguages(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/eolymp.executor.Executor/ListLanguages",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExecutorServer).ListLanguages(ctx, req.(*ListLanguagesInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _Executor_DescribeRuntime_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DescribeRuntimeInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExecutorServer).DescribeRuntime(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/eolymp.executor.Executor/DescribeRuntime",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExecutorServer).DescribeRuntime(ctx, req.(*DescribeRuntimeInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _Executor_ListRuntime_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRuntimeInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExecutorServer).ListRuntime(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/eolymp.executor.Executor/ListRuntime",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExecutorServer).ListRuntime(ctx, req.(*ListRuntimeInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _Executor_DescribeCodeTemplate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DescribeCodeTemplateInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExecutorServer).DescribeCodeTemplate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/eolymp.executor.Executor/DescribeCodeTemplate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExecutorServer).DescribeCodeTemplate(ctx, req.(*DescribeCodeTemplateInput))
	}
	return interceptor(ctx, in, info, handler)
}

// Executor_ServiceDesc is the grpc.ServiceDesc for Executor service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Executor_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "eolymp.executor.Executor",
	HandlerType: (*ExecutorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DescribeLanguage",
			Handler:    _Executor_DescribeLanguage_Handler,
		},
		{
			MethodName: "ListLanguages",
			Handler:    _Executor_ListLanguages_Handler,
		},
		{
			MethodName: "DescribeRuntime",
			Handler:    _Executor_DescribeRuntime_Handler,
		},
		{
			MethodName: "ListRuntime",
			Handler:    _Executor_ListRuntime_Handler,
		},
		{
			MethodName: "DescribeCodeTemplate",
			Handler:    _Executor_DescribeCodeTemplate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "eolymp/executor/executor.proto",
}
