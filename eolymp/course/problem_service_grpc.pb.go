// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: eolymp/course/problem_service.proto

package course

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
	ProblemService_ListStatements_FullMethodName     = "/eolymp.course.ProblemService/ListStatements"
	ProblemService_ListExamples_FullMethodName       = "/eolymp.course.ProblemService/ListExamples"
	ProblemService_CreateSubmission_FullMethodName   = "/eolymp.course.ProblemService/CreateSubmission"
	ProblemService_ListSubmissions_FullMethodName    = "/eolymp.course.ProblemService/ListSubmissions"
	ProblemService_DescribeSubmission_FullMethodName = "/eolymp.course.ProblemService/DescribeSubmission"
	ProblemService_WatchSubmission_FullMethodName    = "/eolymp.course.ProblemService/WatchSubmission"
	ProblemService_LookupCodeTemplate_FullMethodName = "/eolymp.course.ProblemService/LookupCodeTemplate"
	ProblemService_CreateRun_FullMethodName          = "/eolymp.course.ProblemService/CreateRun"
	ProblemService_DescribeRun_FullMethodName        = "/eolymp.course.ProblemService/DescribeRun"
	ProblemService_WatchRun_FullMethodName           = "/eolymp.course.ProblemService/WatchRun"
)

// ProblemServiceClient is the client API for ProblemService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProblemServiceClient interface {
	ListStatements(ctx context.Context, in *ListStatementsInput, opts ...grpc.CallOption) (*ListStatementsOutput, error)
	ListExamples(ctx context.Context, in *ListExamplesInput, opts ...grpc.CallOption) (*ListExamplesOutput, error)
	CreateSubmission(ctx context.Context, in *CreateSubmissionInput, opts ...grpc.CallOption) (*CreateSubmissionOutput, error)
	ListSubmissions(ctx context.Context, in *ListSubmissionsInput, opts ...grpc.CallOption) (*ListSubmissionsOutput, error)
	DescribeSubmission(ctx context.Context, in *DescribeSubmissionInput, opts ...grpc.CallOption) (*DescribeSubmissionOutput, error)
	WatchSubmission(ctx context.Context, in *WatchSubmissionInput, opts ...grpc.CallOption) (ProblemService_WatchSubmissionClient, error)
	LookupCodeTemplate(ctx context.Context, in *LookupCodeTemplateInput, opts ...grpc.CallOption) (*LookupCodeTemplateOutput, error)
	CreateRun(ctx context.Context, in *CreateRunInput, opts ...grpc.CallOption) (*CreateRunOutput, error)
	DescribeRun(ctx context.Context, in *DescribeRunInput, opts ...grpc.CallOption) (*DescribeRunOutput, error)
	WatchRun(ctx context.Context, in *WatchRunInput, opts ...grpc.CallOption) (ProblemService_WatchRunClient, error)
}

type problemServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewProblemServiceClient(cc grpc.ClientConnInterface) ProblemServiceClient {
	return &problemServiceClient{cc}
}

func (c *problemServiceClient) ListStatements(ctx context.Context, in *ListStatementsInput, opts ...grpc.CallOption) (*ListStatementsOutput, error) {
	out := new(ListStatementsOutput)
	err := c.cc.Invoke(ctx, ProblemService_ListStatements_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *problemServiceClient) ListExamples(ctx context.Context, in *ListExamplesInput, opts ...grpc.CallOption) (*ListExamplesOutput, error) {
	out := new(ListExamplesOutput)
	err := c.cc.Invoke(ctx, ProblemService_ListExamples_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *problemServiceClient) CreateSubmission(ctx context.Context, in *CreateSubmissionInput, opts ...grpc.CallOption) (*CreateSubmissionOutput, error) {
	out := new(CreateSubmissionOutput)
	err := c.cc.Invoke(ctx, ProblemService_CreateSubmission_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *problemServiceClient) ListSubmissions(ctx context.Context, in *ListSubmissionsInput, opts ...grpc.CallOption) (*ListSubmissionsOutput, error) {
	out := new(ListSubmissionsOutput)
	err := c.cc.Invoke(ctx, ProblemService_ListSubmissions_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *problemServiceClient) DescribeSubmission(ctx context.Context, in *DescribeSubmissionInput, opts ...grpc.CallOption) (*DescribeSubmissionOutput, error) {
	out := new(DescribeSubmissionOutput)
	err := c.cc.Invoke(ctx, ProblemService_DescribeSubmission_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *problemServiceClient) WatchSubmission(ctx context.Context, in *WatchSubmissionInput, opts ...grpc.CallOption) (ProblemService_WatchSubmissionClient, error) {
	stream, err := c.cc.NewStream(ctx, &ProblemService_ServiceDesc.Streams[0], ProblemService_WatchSubmission_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &problemServiceWatchSubmissionClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ProblemService_WatchSubmissionClient interface {
	Recv() (*WatchSubmissionOutput, error)
	grpc.ClientStream
}

type problemServiceWatchSubmissionClient struct {
	grpc.ClientStream
}

func (x *problemServiceWatchSubmissionClient) Recv() (*WatchSubmissionOutput, error) {
	m := new(WatchSubmissionOutput)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *problemServiceClient) LookupCodeTemplate(ctx context.Context, in *LookupCodeTemplateInput, opts ...grpc.CallOption) (*LookupCodeTemplateOutput, error) {
	out := new(LookupCodeTemplateOutput)
	err := c.cc.Invoke(ctx, ProblemService_LookupCodeTemplate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *problemServiceClient) CreateRun(ctx context.Context, in *CreateRunInput, opts ...grpc.CallOption) (*CreateRunOutput, error) {
	out := new(CreateRunOutput)
	err := c.cc.Invoke(ctx, ProblemService_CreateRun_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *problemServiceClient) DescribeRun(ctx context.Context, in *DescribeRunInput, opts ...grpc.CallOption) (*DescribeRunOutput, error) {
	out := new(DescribeRunOutput)
	err := c.cc.Invoke(ctx, ProblemService_DescribeRun_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *problemServiceClient) WatchRun(ctx context.Context, in *WatchRunInput, opts ...grpc.CallOption) (ProblemService_WatchRunClient, error) {
	stream, err := c.cc.NewStream(ctx, &ProblemService_ServiceDesc.Streams[1], ProblemService_WatchRun_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &problemServiceWatchRunClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ProblemService_WatchRunClient interface {
	Recv() (*WatchRunOutput, error)
	grpc.ClientStream
}

type problemServiceWatchRunClient struct {
	grpc.ClientStream
}

func (x *problemServiceWatchRunClient) Recv() (*WatchRunOutput, error) {
	m := new(WatchRunOutput)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ProblemServiceServer is the server API for ProblemService service.
// All implementations should embed UnimplementedProblemServiceServer
// for forward compatibility
type ProblemServiceServer interface {
	ListStatements(context.Context, *ListStatementsInput) (*ListStatementsOutput, error)
	ListExamples(context.Context, *ListExamplesInput) (*ListExamplesOutput, error)
	CreateSubmission(context.Context, *CreateSubmissionInput) (*CreateSubmissionOutput, error)
	ListSubmissions(context.Context, *ListSubmissionsInput) (*ListSubmissionsOutput, error)
	DescribeSubmission(context.Context, *DescribeSubmissionInput) (*DescribeSubmissionOutput, error)
	WatchSubmission(*WatchSubmissionInput, ProblemService_WatchSubmissionServer) error
	LookupCodeTemplate(context.Context, *LookupCodeTemplateInput) (*LookupCodeTemplateOutput, error)
	CreateRun(context.Context, *CreateRunInput) (*CreateRunOutput, error)
	DescribeRun(context.Context, *DescribeRunInput) (*DescribeRunOutput, error)
	WatchRun(*WatchRunInput, ProblemService_WatchRunServer) error
}

// UnimplementedProblemServiceServer should be embedded to have forward compatible implementations.
type UnimplementedProblemServiceServer struct {
}

func (UnimplementedProblemServiceServer) ListStatements(context.Context, *ListStatementsInput) (*ListStatementsOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListStatements not implemented")
}
func (UnimplementedProblemServiceServer) ListExamples(context.Context, *ListExamplesInput) (*ListExamplesOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListExamples not implemented")
}
func (UnimplementedProblemServiceServer) CreateSubmission(context.Context, *CreateSubmissionInput) (*CreateSubmissionOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateSubmission not implemented")
}
func (UnimplementedProblemServiceServer) ListSubmissions(context.Context, *ListSubmissionsInput) (*ListSubmissionsOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListSubmissions not implemented")
}
func (UnimplementedProblemServiceServer) DescribeSubmission(context.Context, *DescribeSubmissionInput) (*DescribeSubmissionOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DescribeSubmission not implemented")
}
func (UnimplementedProblemServiceServer) WatchSubmission(*WatchSubmissionInput, ProblemService_WatchSubmissionServer) error {
	return status.Errorf(codes.Unimplemented, "method WatchSubmission not implemented")
}
func (UnimplementedProblemServiceServer) LookupCodeTemplate(context.Context, *LookupCodeTemplateInput) (*LookupCodeTemplateOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LookupCodeTemplate not implemented")
}
func (UnimplementedProblemServiceServer) CreateRun(context.Context, *CreateRunInput) (*CreateRunOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateRun not implemented")
}
func (UnimplementedProblemServiceServer) DescribeRun(context.Context, *DescribeRunInput) (*DescribeRunOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DescribeRun not implemented")
}
func (UnimplementedProblemServiceServer) WatchRun(*WatchRunInput, ProblemService_WatchRunServer) error {
	return status.Errorf(codes.Unimplemented, "method WatchRun not implemented")
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

func _ProblemService_CreateSubmission_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateSubmissionInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProblemServiceServer).CreateSubmission(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProblemService_CreateSubmission_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProblemServiceServer).CreateSubmission(ctx, req.(*CreateSubmissionInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProblemService_ListSubmissions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListSubmissionsInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProblemServiceServer).ListSubmissions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProblemService_ListSubmissions_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProblemServiceServer).ListSubmissions(ctx, req.(*ListSubmissionsInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProblemService_DescribeSubmission_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DescribeSubmissionInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProblemServiceServer).DescribeSubmission(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProblemService_DescribeSubmission_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProblemServiceServer).DescribeSubmission(ctx, req.(*DescribeSubmissionInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProblemService_WatchSubmission_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(WatchSubmissionInput)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ProblemServiceServer).WatchSubmission(m, &problemServiceWatchSubmissionServer{stream})
}

type ProblemService_WatchSubmissionServer interface {
	Send(*WatchSubmissionOutput) error
	grpc.ServerStream
}

type problemServiceWatchSubmissionServer struct {
	grpc.ServerStream
}

func (x *problemServiceWatchSubmissionServer) Send(m *WatchSubmissionOutput) error {
	return x.ServerStream.SendMsg(m)
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

func _ProblemService_CreateRun_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRunInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProblemServiceServer).CreateRun(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProblemService_CreateRun_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProblemServiceServer).CreateRun(ctx, req.(*CreateRunInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProblemService_DescribeRun_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DescribeRunInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProblemServiceServer).DescribeRun(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProblemService_DescribeRun_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProblemServiceServer).DescribeRun(ctx, req.(*DescribeRunInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProblemService_WatchRun_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(WatchRunInput)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ProblemServiceServer).WatchRun(m, &problemServiceWatchRunServer{stream})
}

type ProblemService_WatchRunServer interface {
	Send(*WatchRunOutput) error
	grpc.ServerStream
}

type problemServiceWatchRunServer struct {
	grpc.ServerStream
}

func (x *problemServiceWatchRunServer) Send(m *WatchRunOutput) error {
	return x.ServerStream.SendMsg(m)
}

// ProblemService_ServiceDesc is the grpc.ServiceDesc for ProblemService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ProblemService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "eolymp.course.ProblemService",
	HandlerType: (*ProblemServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListStatements",
			Handler:    _ProblemService_ListStatements_Handler,
		},
		{
			MethodName: "ListExamples",
			Handler:    _ProblemService_ListExamples_Handler,
		},
		{
			MethodName: "CreateSubmission",
			Handler:    _ProblemService_CreateSubmission_Handler,
		},
		{
			MethodName: "ListSubmissions",
			Handler:    _ProblemService_ListSubmissions_Handler,
		},
		{
			MethodName: "DescribeSubmission",
			Handler:    _ProblemService_DescribeSubmission_Handler,
		},
		{
			MethodName: "LookupCodeTemplate",
			Handler:    _ProblemService_LookupCodeTemplate_Handler,
		},
		{
			MethodName: "CreateRun",
			Handler:    _ProblemService_CreateRun_Handler,
		},
		{
			MethodName: "DescribeRun",
			Handler:    _ProblemService_DescribeRun_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "WatchSubmission",
			Handler:       _ProblemService_WatchSubmission_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "WatchRun",
			Handler:       _ProblemService_WatchRun_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "eolymp/course/problem_service.proto",
}
