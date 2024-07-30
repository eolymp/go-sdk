// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.21.12
// source: eolymp/atlas/code_template_service.proto

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
	CodeTemplateService_CreateCodeTemplate_FullMethodName   = "/eolymp.atlas.CodeTemplateService/CreateCodeTemplate"
	CodeTemplateService_UpdateCodeTemplate_FullMethodName   = "/eolymp.atlas.CodeTemplateService/UpdateCodeTemplate"
	CodeTemplateService_DeleteCodeTemplate_FullMethodName   = "/eolymp.atlas.CodeTemplateService/DeleteCodeTemplate"
	CodeTemplateService_ListCodeTemplates_FullMethodName    = "/eolymp.atlas.CodeTemplateService/ListCodeTemplates"
	CodeTemplateService_DescribeCodeTemplate_FullMethodName = "/eolymp.atlas.CodeTemplateService/DescribeCodeTemplate"
	CodeTemplateService_LookupCodeTemplate_FullMethodName   = "/eolymp.atlas.CodeTemplateService/LookupCodeTemplate"
)

// CodeTemplateServiceClient is the client API for CodeTemplateService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CodeTemplateServiceClient interface {
	CreateCodeTemplate(ctx context.Context, in *CreateCodeTemplateInput, opts ...grpc.CallOption) (*CreateCodeTemplateOutput, error)
	UpdateCodeTemplate(ctx context.Context, in *UpdateCodeTemplateInput, opts ...grpc.CallOption) (*UpdateCodeTemplateOutput, error)
	DeleteCodeTemplate(ctx context.Context, in *DeleteCodeTemplateInput, opts ...grpc.CallOption) (*DeleteCodeTemplateOutput, error)
	ListCodeTemplates(ctx context.Context, in *ListCodeTemplatesInput, opts ...grpc.CallOption) (*ListCodeTemplatesOutput, error)
	DescribeCodeTemplate(ctx context.Context, in *DescribeCodeTemplateInput, opts ...grpc.CallOption) (*DescribeCodeTemplateOutput, error)
	LookupCodeTemplate(ctx context.Context, in *LookupCodeTemplateInput, opts ...grpc.CallOption) (*LookupCodeTemplateOutput, error)
}

type codeTemplateServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCodeTemplateServiceClient(cc grpc.ClientConnInterface) CodeTemplateServiceClient {
	return &codeTemplateServiceClient{cc}
}

func (c *codeTemplateServiceClient) CreateCodeTemplate(ctx context.Context, in *CreateCodeTemplateInput, opts ...grpc.CallOption) (*CreateCodeTemplateOutput, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateCodeTemplateOutput)
	err := c.cc.Invoke(ctx, CodeTemplateService_CreateCodeTemplate_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *codeTemplateServiceClient) UpdateCodeTemplate(ctx context.Context, in *UpdateCodeTemplateInput, opts ...grpc.CallOption) (*UpdateCodeTemplateOutput, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateCodeTemplateOutput)
	err := c.cc.Invoke(ctx, CodeTemplateService_UpdateCodeTemplate_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *codeTemplateServiceClient) DeleteCodeTemplate(ctx context.Context, in *DeleteCodeTemplateInput, opts ...grpc.CallOption) (*DeleteCodeTemplateOutput, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteCodeTemplateOutput)
	err := c.cc.Invoke(ctx, CodeTemplateService_DeleteCodeTemplate_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *codeTemplateServiceClient) ListCodeTemplates(ctx context.Context, in *ListCodeTemplatesInput, opts ...grpc.CallOption) (*ListCodeTemplatesOutput, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListCodeTemplatesOutput)
	err := c.cc.Invoke(ctx, CodeTemplateService_ListCodeTemplates_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *codeTemplateServiceClient) DescribeCodeTemplate(ctx context.Context, in *DescribeCodeTemplateInput, opts ...grpc.CallOption) (*DescribeCodeTemplateOutput, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DescribeCodeTemplateOutput)
	err := c.cc.Invoke(ctx, CodeTemplateService_DescribeCodeTemplate_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *codeTemplateServiceClient) LookupCodeTemplate(ctx context.Context, in *LookupCodeTemplateInput, opts ...grpc.CallOption) (*LookupCodeTemplateOutput, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(LookupCodeTemplateOutput)
	err := c.cc.Invoke(ctx, CodeTemplateService_LookupCodeTemplate_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CodeTemplateServiceServer is the server API for CodeTemplateService service.
// All implementations should embed UnimplementedCodeTemplateServiceServer
// for forward compatibility.
type CodeTemplateServiceServer interface {
	CreateCodeTemplate(context.Context, *CreateCodeTemplateInput) (*CreateCodeTemplateOutput, error)
	UpdateCodeTemplate(context.Context, *UpdateCodeTemplateInput) (*UpdateCodeTemplateOutput, error)
	DeleteCodeTemplate(context.Context, *DeleteCodeTemplateInput) (*DeleteCodeTemplateOutput, error)
	ListCodeTemplates(context.Context, *ListCodeTemplatesInput) (*ListCodeTemplatesOutput, error)
	DescribeCodeTemplate(context.Context, *DescribeCodeTemplateInput) (*DescribeCodeTemplateOutput, error)
	LookupCodeTemplate(context.Context, *LookupCodeTemplateInput) (*LookupCodeTemplateOutput, error)
}

// UnimplementedCodeTemplateServiceServer should be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedCodeTemplateServiceServer struct{}

func (UnimplementedCodeTemplateServiceServer) CreateCodeTemplate(context.Context, *CreateCodeTemplateInput) (*CreateCodeTemplateOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCodeTemplate not implemented")
}
func (UnimplementedCodeTemplateServiceServer) UpdateCodeTemplate(context.Context, *UpdateCodeTemplateInput) (*UpdateCodeTemplateOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCodeTemplate not implemented")
}
func (UnimplementedCodeTemplateServiceServer) DeleteCodeTemplate(context.Context, *DeleteCodeTemplateInput) (*DeleteCodeTemplateOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCodeTemplate not implemented")
}
func (UnimplementedCodeTemplateServiceServer) ListCodeTemplates(context.Context, *ListCodeTemplatesInput) (*ListCodeTemplatesOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListCodeTemplates not implemented")
}
func (UnimplementedCodeTemplateServiceServer) DescribeCodeTemplate(context.Context, *DescribeCodeTemplateInput) (*DescribeCodeTemplateOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DescribeCodeTemplate not implemented")
}
func (UnimplementedCodeTemplateServiceServer) LookupCodeTemplate(context.Context, *LookupCodeTemplateInput) (*LookupCodeTemplateOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LookupCodeTemplate not implemented")
}
func (UnimplementedCodeTemplateServiceServer) testEmbeddedByValue() {}

// UnsafeCodeTemplateServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CodeTemplateServiceServer will
// result in compilation errors.
type UnsafeCodeTemplateServiceServer interface {
	mustEmbedUnimplementedCodeTemplateServiceServer()
}

func RegisterCodeTemplateServiceServer(s grpc.ServiceRegistrar, srv CodeTemplateServiceServer) {
	// If the following call pancis, it indicates UnimplementedCodeTemplateServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&CodeTemplateService_ServiceDesc, srv)
}

func _CodeTemplateService_CreateCodeTemplate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCodeTemplateInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CodeTemplateServiceServer).CreateCodeTemplate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CodeTemplateService_CreateCodeTemplate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CodeTemplateServiceServer).CreateCodeTemplate(ctx, req.(*CreateCodeTemplateInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _CodeTemplateService_UpdateCodeTemplate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateCodeTemplateInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CodeTemplateServiceServer).UpdateCodeTemplate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CodeTemplateService_UpdateCodeTemplate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CodeTemplateServiceServer).UpdateCodeTemplate(ctx, req.(*UpdateCodeTemplateInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _CodeTemplateService_DeleteCodeTemplate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteCodeTemplateInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CodeTemplateServiceServer).DeleteCodeTemplate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CodeTemplateService_DeleteCodeTemplate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CodeTemplateServiceServer).DeleteCodeTemplate(ctx, req.(*DeleteCodeTemplateInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _CodeTemplateService_ListCodeTemplates_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListCodeTemplatesInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CodeTemplateServiceServer).ListCodeTemplates(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CodeTemplateService_ListCodeTemplates_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CodeTemplateServiceServer).ListCodeTemplates(ctx, req.(*ListCodeTemplatesInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _CodeTemplateService_DescribeCodeTemplate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DescribeCodeTemplateInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CodeTemplateServiceServer).DescribeCodeTemplate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CodeTemplateService_DescribeCodeTemplate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CodeTemplateServiceServer).DescribeCodeTemplate(ctx, req.(*DescribeCodeTemplateInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _CodeTemplateService_LookupCodeTemplate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LookupCodeTemplateInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CodeTemplateServiceServer).LookupCodeTemplate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CodeTemplateService_LookupCodeTemplate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CodeTemplateServiceServer).LookupCodeTemplate(ctx, req.(*LookupCodeTemplateInput))
	}
	return interceptor(ctx, in, info, handler)
}

// CodeTemplateService_ServiceDesc is the grpc.ServiceDesc for CodeTemplateService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CodeTemplateService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "eolymp.atlas.CodeTemplateService",
	HandlerType: (*CodeTemplateServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateCodeTemplate",
			Handler:    _CodeTemplateService_CreateCodeTemplate_Handler,
		},
		{
			MethodName: "UpdateCodeTemplate",
			Handler:    _CodeTemplateService_UpdateCodeTemplate_Handler,
		},
		{
			MethodName: "DeleteCodeTemplate",
			Handler:    _CodeTemplateService_DeleteCodeTemplate_Handler,
		},
		{
			MethodName: "ListCodeTemplates",
			Handler:    _CodeTemplateService_ListCodeTemplates_Handler,
		},
		{
			MethodName: "DescribeCodeTemplate",
			Handler:    _CodeTemplateService_DescribeCodeTemplate_Handler,
		},
		{
			MethodName: "LookupCodeTemplate",
			Handler:    _CodeTemplateService_LookupCodeTemplate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "eolymp/atlas/code_template_service.proto",
}
