// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.21.12
// source: eolymp/atlas/script_service.proto

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
	ScriptService_CreateScript_FullMethodName   = "/eolymp.atlas.ScriptService/CreateScript"
	ScriptService_UpdateScript_FullMethodName   = "/eolymp.atlas.ScriptService/UpdateScript"
	ScriptService_DeleteScript_FullMethodName   = "/eolymp.atlas.ScriptService/DeleteScript"
	ScriptService_DescribeScript_FullMethodName = "/eolymp.atlas.ScriptService/DescribeScript"
	ScriptService_ListScripts_FullMethodName    = "/eolymp.atlas.ScriptService/ListScripts"
)

// ScriptServiceClient is the client API for ScriptService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ScriptServiceClient interface {
	CreateScript(ctx context.Context, in *CreateScriptInput, opts ...grpc.CallOption) (*CreateScriptOutput, error)
	UpdateScript(ctx context.Context, in *UpdateScriptInput, opts ...grpc.CallOption) (*UpdateScriptOutput, error)
	DeleteScript(ctx context.Context, in *DeleteScriptInput, opts ...grpc.CallOption) (*DeleteScriptOutput, error)
	DescribeScript(ctx context.Context, in *DescribeScriptInput, opts ...grpc.CallOption) (*DescribeScriptOutput, error)
	ListScripts(ctx context.Context, in *ListScriptsInput, opts ...grpc.CallOption) (*ListScriptsOutput, error)
}

type scriptServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewScriptServiceClient(cc grpc.ClientConnInterface) ScriptServiceClient {
	return &scriptServiceClient{cc}
}

func (c *scriptServiceClient) CreateScript(ctx context.Context, in *CreateScriptInput, opts ...grpc.CallOption) (*CreateScriptOutput, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateScriptOutput)
	err := c.cc.Invoke(ctx, ScriptService_CreateScript_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *scriptServiceClient) UpdateScript(ctx context.Context, in *UpdateScriptInput, opts ...grpc.CallOption) (*UpdateScriptOutput, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateScriptOutput)
	err := c.cc.Invoke(ctx, ScriptService_UpdateScript_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *scriptServiceClient) DeleteScript(ctx context.Context, in *DeleteScriptInput, opts ...grpc.CallOption) (*DeleteScriptOutput, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteScriptOutput)
	err := c.cc.Invoke(ctx, ScriptService_DeleteScript_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *scriptServiceClient) DescribeScript(ctx context.Context, in *DescribeScriptInput, opts ...grpc.CallOption) (*DescribeScriptOutput, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DescribeScriptOutput)
	err := c.cc.Invoke(ctx, ScriptService_DescribeScript_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *scriptServiceClient) ListScripts(ctx context.Context, in *ListScriptsInput, opts ...grpc.CallOption) (*ListScriptsOutput, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListScriptsOutput)
	err := c.cc.Invoke(ctx, ScriptService_ListScripts_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ScriptServiceServer is the server API for ScriptService service.
// All implementations should embed UnimplementedScriptServiceServer
// for forward compatibility.
type ScriptServiceServer interface {
	CreateScript(context.Context, *CreateScriptInput) (*CreateScriptOutput, error)
	UpdateScript(context.Context, *UpdateScriptInput) (*UpdateScriptOutput, error)
	DeleteScript(context.Context, *DeleteScriptInput) (*DeleteScriptOutput, error)
	DescribeScript(context.Context, *DescribeScriptInput) (*DescribeScriptOutput, error)
	ListScripts(context.Context, *ListScriptsInput) (*ListScriptsOutput, error)
}

// UnimplementedScriptServiceServer should be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedScriptServiceServer struct{}

func (UnimplementedScriptServiceServer) CreateScript(context.Context, *CreateScriptInput) (*CreateScriptOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateScript not implemented")
}
func (UnimplementedScriptServiceServer) UpdateScript(context.Context, *UpdateScriptInput) (*UpdateScriptOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateScript not implemented")
}
func (UnimplementedScriptServiceServer) DeleteScript(context.Context, *DeleteScriptInput) (*DeleteScriptOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteScript not implemented")
}
func (UnimplementedScriptServiceServer) DescribeScript(context.Context, *DescribeScriptInput) (*DescribeScriptOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DescribeScript not implemented")
}
func (UnimplementedScriptServiceServer) ListScripts(context.Context, *ListScriptsInput) (*ListScriptsOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListScripts not implemented")
}
func (UnimplementedScriptServiceServer) testEmbeddedByValue() {}

// UnsafeScriptServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ScriptServiceServer will
// result in compilation errors.
type UnsafeScriptServiceServer interface {
	mustEmbedUnimplementedScriptServiceServer()
}

func RegisterScriptServiceServer(s grpc.ServiceRegistrar, srv ScriptServiceServer) {
	// If the following call pancis, it indicates UnimplementedScriptServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&ScriptService_ServiceDesc, srv)
}

func _ScriptService_CreateScript_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateScriptInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ScriptServiceServer).CreateScript(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ScriptService_CreateScript_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ScriptServiceServer).CreateScript(ctx, req.(*CreateScriptInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _ScriptService_UpdateScript_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateScriptInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ScriptServiceServer).UpdateScript(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ScriptService_UpdateScript_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ScriptServiceServer).UpdateScript(ctx, req.(*UpdateScriptInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _ScriptService_DeleteScript_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteScriptInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ScriptServiceServer).DeleteScript(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ScriptService_DeleteScript_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ScriptServiceServer).DeleteScript(ctx, req.(*DeleteScriptInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _ScriptService_DescribeScript_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DescribeScriptInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ScriptServiceServer).DescribeScript(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ScriptService_DescribeScript_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ScriptServiceServer).DescribeScript(ctx, req.(*DescribeScriptInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _ScriptService_ListScripts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListScriptsInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ScriptServiceServer).ListScripts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ScriptService_ListScripts_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ScriptServiceServer).ListScripts(ctx, req.(*ListScriptsInput))
	}
	return interceptor(ctx, in, info, handler)
}

// ScriptService_ServiceDesc is the grpc.ServiceDesc for ScriptService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ScriptService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "eolymp.atlas.ScriptService",
	HandlerType: (*ScriptServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateScript",
			Handler:    _ScriptService_CreateScript_Handler,
		},
		{
			MethodName: "UpdateScript",
			Handler:    _ScriptService_UpdateScript_Handler,
		},
		{
			MethodName: "DeleteScript",
			Handler:    _ScriptService_DeleteScript_Handler,
		},
		{
			MethodName: "DescribeScript",
			Handler:    _ScriptService_DescribeScript_Handler,
		},
		{
			MethodName: "ListScripts",
			Handler:    _ScriptService_ListScripts_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "eolymp/atlas/script_service.proto",
}
