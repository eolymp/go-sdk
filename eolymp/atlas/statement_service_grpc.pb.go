// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.4
// source: eolymp/atlas/statement_service.proto

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
	StatementService_CreateStatement_FullMethodName     = "/eolymp.atlas.StatementService/CreateStatement"
	StatementService_UpdateStatement_FullMethodName     = "/eolymp.atlas.StatementService/UpdateStatement"
	StatementService_DeleteStatement_FullMethodName     = "/eolymp.atlas.StatementService/DeleteStatement"
	StatementService_DescribeStatement_FullMethodName   = "/eolymp.atlas.StatementService/DescribeStatement"
	StatementService_LookupStatement_FullMethodName     = "/eolymp.atlas.StatementService/LookupStatement"
	StatementService_PreviewStatement_FullMethodName    = "/eolymp.atlas.StatementService/PreviewStatement"
	StatementService_ListStatements_FullMethodName      = "/eolymp.atlas.StatementService/ListStatements"
	StatementService_TranslateStatements_FullMethodName = "/eolymp.atlas.StatementService/TranslateStatements"
	StatementService_ExportStatement_FullMethodName     = "/eolymp.atlas.StatementService/ExportStatement"
)

// StatementServiceClient is the client API for StatementService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StatementServiceClient interface {
	CreateStatement(ctx context.Context, in *CreateStatementInput, opts ...grpc.CallOption) (*CreateStatementOutput, error)
	UpdateStatement(ctx context.Context, in *UpdateStatementInput, opts ...grpc.CallOption) (*UpdateStatementOutput, error)
	DeleteStatement(ctx context.Context, in *DeleteStatementInput, opts ...grpc.CallOption) (*DeleteStatementOutput, error)
	// DescribeStatement returns statement.
	DescribeStatement(ctx context.Context, in *DescribeStatementInput, opts ...grpc.CallOption) (*DescribeStatementOutput, error)
	// LookupStatement finds a statement in one of the requested languages.
	LookupStatement(ctx context.Context, in *LookupStatementInput, opts ...grpc.CallOption) (*LookupStatementOutput, error)
	// PreviewStatement renders unsaved statement.
	//
	// This method can be used to render statement before it has been saved.
	PreviewStatement(ctx context.Context, in *PreviewStatementInput, opts ...grpc.CallOption) (*PreviewStatementOutput, error)
	ListStatements(ctx context.Context, in *ListStatementsInput, opts ...grpc.CallOption) (*ListStatementsOutput, error)
	TranslateStatements(ctx context.Context, in *TranslateStatementsInput, opts ...grpc.CallOption) (*TranslateStatementsOutput, error)
	ExportStatement(ctx context.Context, in *ExportStatementInput, opts ...grpc.CallOption) (*ExportStatementOutput, error)
}

type statementServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewStatementServiceClient(cc grpc.ClientConnInterface) StatementServiceClient {
	return &statementServiceClient{cc}
}

func (c *statementServiceClient) CreateStatement(ctx context.Context, in *CreateStatementInput, opts ...grpc.CallOption) (*CreateStatementOutput, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateStatementOutput)
	err := c.cc.Invoke(ctx, StatementService_CreateStatement_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *statementServiceClient) UpdateStatement(ctx context.Context, in *UpdateStatementInput, opts ...grpc.CallOption) (*UpdateStatementOutput, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateStatementOutput)
	err := c.cc.Invoke(ctx, StatementService_UpdateStatement_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *statementServiceClient) DeleteStatement(ctx context.Context, in *DeleteStatementInput, opts ...grpc.CallOption) (*DeleteStatementOutput, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteStatementOutput)
	err := c.cc.Invoke(ctx, StatementService_DeleteStatement_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *statementServiceClient) DescribeStatement(ctx context.Context, in *DescribeStatementInput, opts ...grpc.CallOption) (*DescribeStatementOutput, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DescribeStatementOutput)
	err := c.cc.Invoke(ctx, StatementService_DescribeStatement_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *statementServiceClient) LookupStatement(ctx context.Context, in *LookupStatementInput, opts ...grpc.CallOption) (*LookupStatementOutput, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(LookupStatementOutput)
	err := c.cc.Invoke(ctx, StatementService_LookupStatement_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *statementServiceClient) PreviewStatement(ctx context.Context, in *PreviewStatementInput, opts ...grpc.CallOption) (*PreviewStatementOutput, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PreviewStatementOutput)
	err := c.cc.Invoke(ctx, StatementService_PreviewStatement_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *statementServiceClient) ListStatements(ctx context.Context, in *ListStatementsInput, opts ...grpc.CallOption) (*ListStatementsOutput, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListStatementsOutput)
	err := c.cc.Invoke(ctx, StatementService_ListStatements_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *statementServiceClient) TranslateStatements(ctx context.Context, in *TranslateStatementsInput, opts ...grpc.CallOption) (*TranslateStatementsOutput, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(TranslateStatementsOutput)
	err := c.cc.Invoke(ctx, StatementService_TranslateStatements_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *statementServiceClient) ExportStatement(ctx context.Context, in *ExportStatementInput, opts ...grpc.CallOption) (*ExportStatementOutput, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ExportStatementOutput)
	err := c.cc.Invoke(ctx, StatementService_ExportStatement_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StatementServiceServer is the server API for StatementService service.
// All implementations should embed UnimplementedStatementServiceServer
// for forward compatibility.
type StatementServiceServer interface {
	CreateStatement(context.Context, *CreateStatementInput) (*CreateStatementOutput, error)
	UpdateStatement(context.Context, *UpdateStatementInput) (*UpdateStatementOutput, error)
	DeleteStatement(context.Context, *DeleteStatementInput) (*DeleteStatementOutput, error)
	// DescribeStatement returns statement.
	DescribeStatement(context.Context, *DescribeStatementInput) (*DescribeStatementOutput, error)
	// LookupStatement finds a statement in one of the requested languages.
	LookupStatement(context.Context, *LookupStatementInput) (*LookupStatementOutput, error)
	// PreviewStatement renders unsaved statement.
	//
	// This method can be used to render statement before it has been saved.
	PreviewStatement(context.Context, *PreviewStatementInput) (*PreviewStatementOutput, error)
	ListStatements(context.Context, *ListStatementsInput) (*ListStatementsOutput, error)
	TranslateStatements(context.Context, *TranslateStatementsInput) (*TranslateStatementsOutput, error)
	ExportStatement(context.Context, *ExportStatementInput) (*ExportStatementOutput, error)
}

// UnimplementedStatementServiceServer should be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedStatementServiceServer struct{}

func (UnimplementedStatementServiceServer) CreateStatement(context.Context, *CreateStatementInput) (*CreateStatementOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateStatement not implemented")
}
func (UnimplementedStatementServiceServer) UpdateStatement(context.Context, *UpdateStatementInput) (*UpdateStatementOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateStatement not implemented")
}
func (UnimplementedStatementServiceServer) DeleteStatement(context.Context, *DeleteStatementInput) (*DeleteStatementOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteStatement not implemented")
}
func (UnimplementedStatementServiceServer) DescribeStatement(context.Context, *DescribeStatementInput) (*DescribeStatementOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DescribeStatement not implemented")
}
func (UnimplementedStatementServiceServer) LookupStatement(context.Context, *LookupStatementInput) (*LookupStatementOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LookupStatement not implemented")
}
func (UnimplementedStatementServiceServer) PreviewStatement(context.Context, *PreviewStatementInput) (*PreviewStatementOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PreviewStatement not implemented")
}
func (UnimplementedStatementServiceServer) ListStatements(context.Context, *ListStatementsInput) (*ListStatementsOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListStatements not implemented")
}
func (UnimplementedStatementServiceServer) TranslateStatements(context.Context, *TranslateStatementsInput) (*TranslateStatementsOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TranslateStatements not implemented")
}
func (UnimplementedStatementServiceServer) ExportStatement(context.Context, *ExportStatementInput) (*ExportStatementOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExportStatement not implemented")
}
func (UnimplementedStatementServiceServer) testEmbeddedByValue() {}

// UnsafeStatementServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StatementServiceServer will
// result in compilation errors.
type UnsafeStatementServiceServer interface {
	mustEmbedUnimplementedStatementServiceServer()
}

func RegisterStatementServiceServer(s grpc.ServiceRegistrar, srv StatementServiceServer) {
	// If the following call pancis, it indicates UnimplementedStatementServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&StatementService_ServiceDesc, srv)
}

func _StatementService_CreateStatement_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateStatementInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StatementServiceServer).CreateStatement(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StatementService_CreateStatement_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StatementServiceServer).CreateStatement(ctx, req.(*CreateStatementInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _StatementService_UpdateStatement_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateStatementInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StatementServiceServer).UpdateStatement(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StatementService_UpdateStatement_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StatementServiceServer).UpdateStatement(ctx, req.(*UpdateStatementInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _StatementService_DeleteStatement_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteStatementInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StatementServiceServer).DeleteStatement(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StatementService_DeleteStatement_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StatementServiceServer).DeleteStatement(ctx, req.(*DeleteStatementInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _StatementService_DescribeStatement_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DescribeStatementInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StatementServiceServer).DescribeStatement(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StatementService_DescribeStatement_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StatementServiceServer).DescribeStatement(ctx, req.(*DescribeStatementInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _StatementService_LookupStatement_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LookupStatementInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StatementServiceServer).LookupStatement(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StatementService_LookupStatement_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StatementServiceServer).LookupStatement(ctx, req.(*LookupStatementInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _StatementService_PreviewStatement_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PreviewStatementInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StatementServiceServer).PreviewStatement(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StatementService_PreviewStatement_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StatementServiceServer).PreviewStatement(ctx, req.(*PreviewStatementInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _StatementService_ListStatements_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListStatementsInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StatementServiceServer).ListStatements(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StatementService_ListStatements_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StatementServiceServer).ListStatements(ctx, req.(*ListStatementsInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _StatementService_TranslateStatements_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TranslateStatementsInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StatementServiceServer).TranslateStatements(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StatementService_TranslateStatements_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StatementServiceServer).TranslateStatements(ctx, req.(*TranslateStatementsInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _StatementService_ExportStatement_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExportStatementInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StatementServiceServer).ExportStatement(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StatementService_ExportStatement_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StatementServiceServer).ExportStatement(ctx, req.(*ExportStatementInput))
	}
	return interceptor(ctx, in, info, handler)
}

// StatementService_ServiceDesc is the grpc.ServiceDesc for StatementService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var StatementService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "eolymp.atlas.StatementService",
	HandlerType: (*StatementServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateStatement",
			Handler:    _StatementService_CreateStatement_Handler,
		},
		{
			MethodName: "UpdateStatement",
			Handler:    _StatementService_UpdateStatement_Handler,
		},
		{
			MethodName: "DeleteStatement",
			Handler:    _StatementService_DeleteStatement_Handler,
		},
		{
			MethodName: "DescribeStatement",
			Handler:    _StatementService_DescribeStatement_Handler,
		},
		{
			MethodName: "LookupStatement",
			Handler:    _StatementService_LookupStatement_Handler,
		},
		{
			MethodName: "PreviewStatement",
			Handler:    _StatementService_PreviewStatement_Handler,
		},
		{
			MethodName: "ListStatements",
			Handler:    _StatementService_ListStatements_Handler,
		},
		{
			MethodName: "TranslateStatements",
			Handler:    _StatementService_TranslateStatements_Handler,
		},
		{
			MethodName: "ExportStatement",
			Handler:    _StatementService_ExportStatement_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "eolymp/atlas/statement_service.proto",
}
