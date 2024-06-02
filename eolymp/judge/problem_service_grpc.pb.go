// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v3.21.12
// source: eolymp/judge/problem_service.proto

package judge

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	ProblemService_ImportProblem_FullMethodName        = "/eolymp.judge.ProblemService/ImportProblem"
	ProblemService_SyncProblem_FullMethodName          = "/eolymp.judge.ProblemService/SyncProblem"
	ProblemService_UpdateProblem_FullMethodName        = "/eolymp.judge.ProblemService/UpdateProblem"
	ProblemService_ListProblems_FullMethodName         = "/eolymp.judge.ProblemService/ListProblems"
	ProblemService_DescribeProblem_FullMethodName      = "/eolymp.judge.ProblemService/DescribeProblem"
	ProblemService_DeleteProblem_FullMethodName        = "/eolymp.judge.ProblemService/DeleteProblem"
	ProblemService_LookupCodeTemplate_FullMethodName   = "/eolymp.judge.ProblemService/LookupCodeTemplate"
	ProblemService_DescribeCodeTemplate_FullMethodName = "/eolymp.judge.ProblemService/DescribeCodeTemplate"
	ProblemService_ListStatements_FullMethodName       = "/eolymp.judge.ProblemService/ListStatements"
	ProblemService_ListAttachments_FullMethodName      = "/eolymp.judge.ProblemService/ListAttachments"
	ProblemService_ListExamples_FullMethodName         = "/eolymp.judge.ProblemService/ListExamples"
)

// ProblemServiceClient is the client API for ProblemService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProblemServiceClient interface {
	// ImportProblem from Atlas (problem catalog)
	ImportProblem(ctx context.Context, in *ImportProblemInput, opts ...grpc.CallOption) (*ImportProblemOutput, error)
	// SyncProblem with Atlas (problem catalog)
	SyncProblem(ctx context.Context, in *SyncProblemInput, opts ...grpc.CallOption) (*SyncProblemOutput, error)
	UpdateProblem(ctx context.Context, in *UpdateProblemInput, opts ...grpc.CallOption) (*UpdateProblemOutput, error)
	ListProblems(ctx context.Context, in *ListProblemsInput, opts ...grpc.CallOption) (*ListProblemsOutput, error)
	DescribeProblem(ctx context.Context, in *DescribeProblemInput, opts ...grpc.CallOption) (*DescribeProblemOutput, error)
	DeleteProblem(ctx context.Context, in *DeleteProblemInput, opts ...grpc.CallOption) (*DeleteProblemOutput, error)
	// Lookup template for a given runtime/language
	LookupCodeTemplate(ctx context.Context, in *LookupCodeTemplateInput, opts ...grpc.CallOption) (*LookupCodeTemplateOutput, error)
	// Return code template for problem
	DescribeCodeTemplate(ctx context.Context, in *DescribeCodeTemplateInput, opts ...grpc.CallOption) (*DescribeCodeTemplateOutput, error)
	ListStatements(ctx context.Context, in *ListStatementsInput, opts ...grpc.CallOption) (*ListStatementsOutput, error)
	ListAttachments(ctx context.Context, in *ListAttachmentsInput, opts ...grpc.CallOption) (*ListAttachmentsOutput, error)
	ListExamples(ctx context.Context, in *ListExamplesInput, opts ...grpc.CallOption) (*ListExamplesOutput, error)
}

type problemServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewProblemServiceClient(cc grpc.ClientConnInterface) ProblemServiceClient {
	return &problemServiceClient{cc}
}

func (c *problemServiceClient) ImportProblem(ctx context.Context, in *ImportProblemInput, opts ...grpc.CallOption) (*ImportProblemOutput, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ImportProblemOutput)
	err := c.cc.Invoke(ctx, ProblemService_ImportProblem_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *problemServiceClient) SyncProblem(ctx context.Context, in *SyncProblemInput, opts ...grpc.CallOption) (*SyncProblemOutput, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SyncProblemOutput)
	err := c.cc.Invoke(ctx, ProblemService_SyncProblem_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *problemServiceClient) UpdateProblem(ctx context.Context, in *UpdateProblemInput, opts ...grpc.CallOption) (*UpdateProblemOutput, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateProblemOutput)
	err := c.cc.Invoke(ctx, ProblemService_UpdateProblem_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *problemServiceClient) ListProblems(ctx context.Context, in *ListProblemsInput, opts ...grpc.CallOption) (*ListProblemsOutput, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListProblemsOutput)
	err := c.cc.Invoke(ctx, ProblemService_ListProblems_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *problemServiceClient) DescribeProblem(ctx context.Context, in *DescribeProblemInput, opts ...grpc.CallOption) (*DescribeProblemOutput, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DescribeProblemOutput)
	err := c.cc.Invoke(ctx, ProblemService_DescribeProblem_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *problemServiceClient) DeleteProblem(ctx context.Context, in *DeleteProblemInput, opts ...grpc.CallOption) (*DeleteProblemOutput, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteProblemOutput)
	err := c.cc.Invoke(ctx, ProblemService_DeleteProblem_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *problemServiceClient) LookupCodeTemplate(ctx context.Context, in *LookupCodeTemplateInput, opts ...grpc.CallOption) (*LookupCodeTemplateOutput, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(LookupCodeTemplateOutput)
	err := c.cc.Invoke(ctx, ProblemService_LookupCodeTemplate_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *problemServiceClient) DescribeCodeTemplate(ctx context.Context, in *DescribeCodeTemplateInput, opts ...grpc.CallOption) (*DescribeCodeTemplateOutput, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DescribeCodeTemplateOutput)
	err := c.cc.Invoke(ctx, ProblemService_DescribeCodeTemplate_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *problemServiceClient) ListStatements(ctx context.Context, in *ListStatementsInput, opts ...grpc.CallOption) (*ListStatementsOutput, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListStatementsOutput)
	err := c.cc.Invoke(ctx, ProblemService_ListStatements_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *problemServiceClient) ListAttachments(ctx context.Context, in *ListAttachmentsInput, opts ...grpc.CallOption) (*ListAttachmentsOutput, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListAttachmentsOutput)
	err := c.cc.Invoke(ctx, ProblemService_ListAttachments_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *problemServiceClient) ListExamples(ctx context.Context, in *ListExamplesInput, opts ...grpc.CallOption) (*ListExamplesOutput, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListExamplesOutput)
	err := c.cc.Invoke(ctx, ProblemService_ListExamples_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProblemServiceServer is the server API for ProblemService service.
// All implementations should embed UnimplementedProblemServiceServer
// for forward compatibility
type ProblemServiceServer interface {
	// ImportProblem from Atlas (problem catalog)
	ImportProblem(context.Context, *ImportProblemInput) (*ImportProblemOutput, error)
	// SyncProblem with Atlas (problem catalog)
	SyncProblem(context.Context, *SyncProblemInput) (*SyncProblemOutput, error)
	UpdateProblem(context.Context, *UpdateProblemInput) (*UpdateProblemOutput, error)
	ListProblems(context.Context, *ListProblemsInput) (*ListProblemsOutput, error)
	DescribeProblem(context.Context, *DescribeProblemInput) (*DescribeProblemOutput, error)
	DeleteProblem(context.Context, *DeleteProblemInput) (*DeleteProblemOutput, error)
	// Lookup template for a given runtime/language
	LookupCodeTemplate(context.Context, *LookupCodeTemplateInput) (*LookupCodeTemplateOutput, error)
	// Return code template for problem
	DescribeCodeTemplate(context.Context, *DescribeCodeTemplateInput) (*DescribeCodeTemplateOutput, error)
	ListStatements(context.Context, *ListStatementsInput) (*ListStatementsOutput, error)
	ListAttachments(context.Context, *ListAttachmentsInput) (*ListAttachmentsOutput, error)
	ListExamples(context.Context, *ListExamplesInput) (*ListExamplesOutput, error)
}

// UnimplementedProblemServiceServer should be embedded to have forward compatible implementations.
type UnimplementedProblemServiceServer struct {
}

func (UnimplementedProblemServiceServer) ImportProblem(context.Context, *ImportProblemInput) (*ImportProblemOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ImportProblem not implemented")
}
func (UnimplementedProblemServiceServer) SyncProblem(context.Context, *SyncProblemInput) (*SyncProblemOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SyncProblem not implemented")
}
func (UnimplementedProblemServiceServer) UpdateProblem(context.Context, *UpdateProblemInput) (*UpdateProblemOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateProblem not implemented")
}
func (UnimplementedProblemServiceServer) ListProblems(context.Context, *ListProblemsInput) (*ListProblemsOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListProblems not implemented")
}
func (UnimplementedProblemServiceServer) DescribeProblem(context.Context, *DescribeProblemInput) (*DescribeProblemOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DescribeProblem not implemented")
}
func (UnimplementedProblemServiceServer) DeleteProblem(context.Context, *DeleteProblemInput) (*DeleteProblemOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteProblem not implemented")
}
func (UnimplementedProblemServiceServer) LookupCodeTemplate(context.Context, *LookupCodeTemplateInput) (*LookupCodeTemplateOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LookupCodeTemplate not implemented")
}
func (UnimplementedProblemServiceServer) DescribeCodeTemplate(context.Context, *DescribeCodeTemplateInput) (*DescribeCodeTemplateOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DescribeCodeTemplate not implemented")
}
func (UnimplementedProblemServiceServer) ListStatements(context.Context, *ListStatementsInput) (*ListStatementsOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListStatements not implemented")
}
func (UnimplementedProblemServiceServer) ListAttachments(context.Context, *ListAttachmentsInput) (*ListAttachmentsOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListAttachments not implemented")
}
func (UnimplementedProblemServiceServer) ListExamples(context.Context, *ListExamplesInput) (*ListExamplesOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListExamples not implemented")
}

// UnsafeProblemServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProblemServiceServer will
// result in compilation errors.
type UnsafeProblemServiceServer interface {
	mustEmbedUnimplementedProblemServiceServer()
}

func RegisterProblemServiceServer(s grpc.ServiceRegistrar, srv ProblemServiceServer) {
	s.RegisterService(&ProblemService_ServiceDesc, srv)
}

func _ProblemService_ImportProblem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ImportProblemInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProblemServiceServer).ImportProblem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProblemService_ImportProblem_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProblemServiceServer).ImportProblem(ctx, req.(*ImportProblemInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProblemService_SyncProblem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SyncProblemInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProblemServiceServer).SyncProblem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProblemService_SyncProblem_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProblemServiceServer).SyncProblem(ctx, req.(*SyncProblemInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProblemService_UpdateProblem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateProblemInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProblemServiceServer).UpdateProblem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProblemService_UpdateProblem_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProblemServiceServer).UpdateProblem(ctx, req.(*UpdateProblemInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProblemService_ListProblems_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListProblemsInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProblemServiceServer).ListProblems(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProblemService_ListProblems_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProblemServiceServer).ListProblems(ctx, req.(*ListProblemsInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProblemService_DescribeProblem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DescribeProblemInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProblemServiceServer).DescribeProblem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProblemService_DescribeProblem_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProblemServiceServer).DescribeProblem(ctx, req.(*DescribeProblemInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProblemService_DeleteProblem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteProblemInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProblemServiceServer).DeleteProblem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProblemService_DeleteProblem_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProblemServiceServer).DeleteProblem(ctx, req.(*DeleteProblemInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProblemService_LookupCodeTemplate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LookupCodeTemplateInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProblemServiceServer).LookupCodeTemplate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProblemService_LookupCodeTemplate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProblemServiceServer).LookupCodeTemplate(ctx, req.(*LookupCodeTemplateInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProblemService_DescribeCodeTemplate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DescribeCodeTemplateInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProblemServiceServer).DescribeCodeTemplate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProblemService_DescribeCodeTemplate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProblemServiceServer).DescribeCodeTemplate(ctx, req.(*DescribeCodeTemplateInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProblemService_ListStatements_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListStatementsInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProblemServiceServer).ListStatements(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProblemService_ListStatements_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProblemServiceServer).ListStatements(ctx, req.(*ListStatementsInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProblemService_ListAttachments_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListAttachmentsInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProblemServiceServer).ListAttachments(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProblemService_ListAttachments_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProblemServiceServer).ListAttachments(ctx, req.(*ListAttachmentsInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProblemService_ListExamples_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListExamplesInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProblemServiceServer).ListExamples(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProblemService_ListExamples_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProblemServiceServer).ListExamples(ctx, req.(*ListExamplesInput))
	}
	return interceptor(ctx, in, info, handler)
}

// ProblemService_ServiceDesc is the grpc.ServiceDesc for ProblemService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ProblemService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "eolymp.judge.ProblemService",
	HandlerType: (*ProblemServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ImportProblem",
			Handler:    _ProblemService_ImportProblem_Handler,
		},
		{
			MethodName: "SyncProblem",
			Handler:    _ProblemService_SyncProblem_Handler,
		},
		{
			MethodName: "UpdateProblem",
			Handler:    _ProblemService_UpdateProblem_Handler,
		},
		{
			MethodName: "ListProblems",
			Handler:    _ProblemService_ListProblems_Handler,
		},
		{
			MethodName: "DescribeProblem",
			Handler:    _ProblemService_DescribeProblem_Handler,
		},
		{
			MethodName: "DeleteProblem",
			Handler:    _ProblemService_DeleteProblem_Handler,
		},
		{
			MethodName: "LookupCodeTemplate",
			Handler:    _ProblemService_LookupCodeTemplate_Handler,
		},
		{
			MethodName: "DescribeCodeTemplate",
			Handler:    _ProblemService_DescribeCodeTemplate_Handler,
		},
		{
			MethodName: "ListStatements",
			Handler:    _ProblemService_ListStatements_Handler,
		},
		{
			MethodName: "ListAttachments",
			Handler:    _ProblemService_ListAttachments_Handler,
		},
		{
			MethodName: "ListExamples",
			Handler:    _ProblemService_ListExamples_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "eolymp/judge/problem_service.proto",
}
