// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.18.1
// source: eolymp/atlas/problem_solver.proto

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
	ProblemSolver_ListStatements_FullMethodName     = "/eolymp.atlas.ProblemSolver/ListStatements"
	ProblemSolver_ListExamples_FullMethodName       = "/eolymp.atlas.ProblemSolver/ListExamples"
	ProblemSolver_CreateSubmission_FullMethodName   = "/eolymp.atlas.ProblemSolver/CreateSubmission"
	ProblemSolver_ListSubmissions_FullMethodName    = "/eolymp.atlas.ProblemSolver/ListSubmissions"
	ProblemSolver_DescribeSubmission_FullMethodName = "/eolymp.atlas.ProblemSolver/DescribeSubmission"
	ProblemSolver_WatchSubmission_FullMethodName    = "/eolymp.atlas.ProblemSolver/WatchSubmission"
	ProblemSolver_DescribeScore_FullMethodName      = "/eolymp.atlas.ProblemSolver/DescribeScore"
	ProblemSolver_LookupCodeTemplate_FullMethodName = "/eolymp.atlas.ProblemSolver/LookupCodeTemplate"
)

// ProblemSolverClient is the client API for ProblemSolver service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProblemSolverClient interface {
	ListStatements(ctx context.Context, in *ListStatementsInput, opts ...grpc.CallOption) (*ListStatementsOutput, error)
	ListExamples(ctx context.Context, in *ListExamplesInput, opts ...grpc.CallOption) (*ListExamplesOutput, error)
	CreateSubmission(ctx context.Context, in *CreateSubmissionInput, opts ...grpc.CallOption) (*CreateSubmissionOutput, error)
	ListSubmissions(ctx context.Context, in *ListSubmissionsInput, opts ...grpc.CallOption) (*ListSubmissionsOutput, error)
	DescribeSubmission(ctx context.Context, in *DescribeSubmissionInput, opts ...grpc.CallOption) (*DescribeSubmissionOutput, error)
	WatchSubmission(ctx context.Context, in *WatchSubmissionInput, opts ...grpc.CallOption) (ProblemSolver_WatchSubmissionClient, error)
	DescribeScore(ctx context.Context, in *DescribeScoreInput, opts ...grpc.CallOption) (*DescribeScoreOutput, error)
	LookupCodeTemplate(ctx context.Context, in *LookupCodeTemplateInput, opts ...grpc.CallOption) (*LookupCodeTemplateOutput, error)
}

type problemSolverClient struct {
	cc grpc.ClientConnInterface
}

func NewProblemSolverClient(cc grpc.ClientConnInterface) ProblemSolverClient {
	return &problemSolverClient{cc}
}

func (c *problemSolverClient) ListStatements(ctx context.Context, in *ListStatementsInput, opts ...grpc.CallOption) (*ListStatementsOutput, error) {
	out := new(ListStatementsOutput)
	err := c.cc.Invoke(ctx, ProblemSolver_ListStatements_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *problemSolverClient) ListExamples(ctx context.Context, in *ListExamplesInput, opts ...grpc.CallOption) (*ListExamplesOutput, error) {
	out := new(ListExamplesOutput)
	err := c.cc.Invoke(ctx, ProblemSolver_ListExamples_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *problemSolverClient) CreateSubmission(ctx context.Context, in *CreateSubmissionInput, opts ...grpc.CallOption) (*CreateSubmissionOutput, error) {
	out := new(CreateSubmissionOutput)
	err := c.cc.Invoke(ctx, ProblemSolver_CreateSubmission_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *problemSolverClient) ListSubmissions(ctx context.Context, in *ListSubmissionsInput, opts ...grpc.CallOption) (*ListSubmissionsOutput, error) {
	out := new(ListSubmissionsOutput)
	err := c.cc.Invoke(ctx, ProblemSolver_ListSubmissions_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *problemSolverClient) DescribeSubmission(ctx context.Context, in *DescribeSubmissionInput, opts ...grpc.CallOption) (*DescribeSubmissionOutput, error) {
	out := new(DescribeSubmissionOutput)
	err := c.cc.Invoke(ctx, ProblemSolver_DescribeSubmission_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *problemSolverClient) WatchSubmission(ctx context.Context, in *WatchSubmissionInput, opts ...grpc.CallOption) (ProblemSolver_WatchSubmissionClient, error) {
	stream, err := c.cc.NewStream(ctx, &ProblemSolver_ServiceDesc.Streams[0], ProblemSolver_WatchSubmission_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &problemSolverWatchSubmissionClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ProblemSolver_WatchSubmissionClient interface {
	Recv() (*WatchSubmissionOutput, error)
	grpc.ClientStream
}

type problemSolverWatchSubmissionClient struct {
	grpc.ClientStream
}

func (x *problemSolverWatchSubmissionClient) Recv() (*WatchSubmissionOutput, error) {
	m := new(WatchSubmissionOutput)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *problemSolverClient) DescribeScore(ctx context.Context, in *DescribeScoreInput, opts ...grpc.CallOption) (*DescribeScoreOutput, error) {
	out := new(DescribeScoreOutput)
	err := c.cc.Invoke(ctx, ProblemSolver_DescribeScore_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *problemSolverClient) LookupCodeTemplate(ctx context.Context, in *LookupCodeTemplateInput, opts ...grpc.CallOption) (*LookupCodeTemplateOutput, error) {
	out := new(LookupCodeTemplateOutput)
	err := c.cc.Invoke(ctx, ProblemSolver_LookupCodeTemplate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProblemSolverServer is the server API for ProblemSolver service.
// All implementations should embed UnimplementedProblemSolverServer
// for forward compatibility
type ProblemSolverServer interface {
	ListStatements(context.Context, *ListStatementsInput) (*ListStatementsOutput, error)
	ListExamples(context.Context, *ListExamplesInput) (*ListExamplesOutput, error)
	CreateSubmission(context.Context, *CreateSubmissionInput) (*CreateSubmissionOutput, error)
	ListSubmissions(context.Context, *ListSubmissionsInput) (*ListSubmissionsOutput, error)
	DescribeSubmission(context.Context, *DescribeSubmissionInput) (*DescribeSubmissionOutput, error)
	WatchSubmission(*WatchSubmissionInput, ProblemSolver_WatchSubmissionServer) error
	DescribeScore(context.Context, *DescribeScoreInput) (*DescribeScoreOutput, error)
	LookupCodeTemplate(context.Context, *LookupCodeTemplateInput) (*LookupCodeTemplateOutput, error)
}

// UnimplementedProblemSolverServer should be embedded to have forward compatible implementations.
type UnimplementedProblemSolverServer struct {
}

func (UnimplementedProblemSolverServer) ListStatements(context.Context, *ListStatementsInput) (*ListStatementsOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListStatements not implemented")
}
func (UnimplementedProblemSolverServer) ListExamples(context.Context, *ListExamplesInput) (*ListExamplesOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListExamples not implemented")
}
func (UnimplementedProblemSolverServer) CreateSubmission(context.Context, *CreateSubmissionInput) (*CreateSubmissionOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateSubmission not implemented")
}
func (UnimplementedProblemSolverServer) ListSubmissions(context.Context, *ListSubmissionsInput) (*ListSubmissionsOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListSubmissions not implemented")
}
func (UnimplementedProblemSolverServer) DescribeSubmission(context.Context, *DescribeSubmissionInput) (*DescribeSubmissionOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DescribeSubmission not implemented")
}
func (UnimplementedProblemSolverServer) WatchSubmission(*WatchSubmissionInput, ProblemSolver_WatchSubmissionServer) error {
	return status.Errorf(codes.Unimplemented, "method WatchSubmission not implemented")
}
func (UnimplementedProblemSolverServer) DescribeScore(context.Context, *DescribeScoreInput) (*DescribeScoreOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DescribeScore not implemented")
}
func (UnimplementedProblemSolverServer) LookupCodeTemplate(context.Context, *LookupCodeTemplateInput) (*LookupCodeTemplateOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LookupCodeTemplate not implemented")
}

// UnsafeProblemSolverServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProblemSolverServer will
// result in compilation errors.
type UnsafeProblemSolverServer interface {
	mustEmbedUnimplementedProblemSolverServer()
}

func RegisterProblemSolverServer(s grpc.ServiceRegistrar, srv ProblemSolverServer) {
	s.RegisterService(&ProblemSolver_ServiceDesc, srv)
}

func _ProblemSolver_ListStatements_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListStatementsInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProblemSolverServer).ListStatements(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProblemSolver_ListStatements_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProblemSolverServer).ListStatements(ctx, req.(*ListStatementsInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProblemSolver_ListExamples_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListExamplesInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProblemSolverServer).ListExamples(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProblemSolver_ListExamples_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProblemSolverServer).ListExamples(ctx, req.(*ListExamplesInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProblemSolver_CreateSubmission_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateSubmissionInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProblemSolverServer).CreateSubmission(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProblemSolver_CreateSubmission_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProblemSolverServer).CreateSubmission(ctx, req.(*CreateSubmissionInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProblemSolver_ListSubmissions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListSubmissionsInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProblemSolverServer).ListSubmissions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProblemSolver_ListSubmissions_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProblemSolverServer).ListSubmissions(ctx, req.(*ListSubmissionsInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProblemSolver_DescribeSubmission_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DescribeSubmissionInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProblemSolverServer).DescribeSubmission(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProblemSolver_DescribeSubmission_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProblemSolverServer).DescribeSubmission(ctx, req.(*DescribeSubmissionInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProblemSolver_WatchSubmission_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(WatchSubmissionInput)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ProblemSolverServer).WatchSubmission(m, &problemSolverWatchSubmissionServer{stream})
}

type ProblemSolver_WatchSubmissionServer interface {
	Send(*WatchSubmissionOutput) error
	grpc.ServerStream
}

type problemSolverWatchSubmissionServer struct {
	grpc.ServerStream
}

func (x *problemSolverWatchSubmissionServer) Send(m *WatchSubmissionOutput) error {
	return x.ServerStream.SendMsg(m)
}

func _ProblemSolver_DescribeScore_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DescribeScoreInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProblemSolverServer).DescribeScore(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProblemSolver_DescribeScore_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProblemSolverServer).DescribeScore(ctx, req.(*DescribeScoreInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProblemSolver_LookupCodeTemplate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LookupCodeTemplateInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProblemSolverServer).LookupCodeTemplate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProblemSolver_LookupCodeTemplate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProblemSolverServer).LookupCodeTemplate(ctx, req.(*LookupCodeTemplateInput))
	}
	return interceptor(ctx, in, info, handler)
}

// ProblemSolver_ServiceDesc is the grpc.ServiceDesc for ProblemSolver service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ProblemSolver_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "eolymp.atlas.ProblemSolver",
	HandlerType: (*ProblemSolverServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListStatements",
			Handler:    _ProblemSolver_ListStatements_Handler,
		},
		{
			MethodName: "ListExamples",
			Handler:    _ProblemSolver_ListExamples_Handler,
		},
		{
			MethodName: "CreateSubmission",
			Handler:    _ProblemSolver_CreateSubmission_Handler,
		},
		{
			MethodName: "ListSubmissions",
			Handler:    _ProblemSolver_ListSubmissions_Handler,
		},
		{
			MethodName: "DescribeSubmission",
			Handler:    _ProblemSolver_DescribeSubmission_Handler,
		},
		{
			MethodName: "DescribeScore",
			Handler:    _ProblemSolver_DescribeScore_Handler,
		},
		{
			MethodName: "LookupCodeTemplate",
			Handler:    _ProblemSolver_LookupCodeTemplate_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "WatchSubmission",
			Handler:       _ProblemSolver_WatchSubmission_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "eolymp/atlas/problem_solver.proto",
}
