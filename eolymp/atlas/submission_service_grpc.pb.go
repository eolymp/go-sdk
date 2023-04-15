// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.18.1
// source: eolymp/atlas/submission_service.proto

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
	SubmissionService_CreateSubmission_FullMethodName   = "/eolymp.atlas.SubmissionService/CreateSubmission"
	SubmissionService_RetestSubmission_FullMethodName   = "/eolymp.atlas.SubmissionService/RetestSubmission"
	SubmissionService_DescribeSubmission_FullMethodName = "/eolymp.atlas.SubmissionService/DescribeSubmission"
)

// SubmissionServiceClient is the client API for SubmissionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SubmissionServiceClient interface {
	CreateSubmission(ctx context.Context, in *CreateSubmissionInput, opts ...grpc.CallOption) (*CreateSubmissionOutput, error)
	RetestSubmission(ctx context.Context, in *RetestSubmissionInput, opts ...grpc.CallOption) (*RetestSubmissionOutput, error)
	DescribeSubmission(ctx context.Context, in *DescribeSubmissionInput, opts ...grpc.CallOption) (*DescribeSubmissionOutput, error)
}

type submissionServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSubmissionServiceClient(cc grpc.ClientConnInterface) SubmissionServiceClient {
	return &submissionServiceClient{cc}
}

func (c *submissionServiceClient) CreateSubmission(ctx context.Context, in *CreateSubmissionInput, opts ...grpc.CallOption) (*CreateSubmissionOutput, error) {
	out := new(CreateSubmissionOutput)
	err := c.cc.Invoke(ctx, SubmissionService_CreateSubmission_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *submissionServiceClient) RetestSubmission(ctx context.Context, in *RetestSubmissionInput, opts ...grpc.CallOption) (*RetestSubmissionOutput, error) {
	out := new(RetestSubmissionOutput)
	err := c.cc.Invoke(ctx, SubmissionService_RetestSubmission_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *submissionServiceClient) DescribeSubmission(ctx context.Context, in *DescribeSubmissionInput, opts ...grpc.CallOption) (*DescribeSubmissionOutput, error) {
	out := new(DescribeSubmissionOutput)
	err := c.cc.Invoke(ctx, SubmissionService_DescribeSubmission_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SubmissionServiceServer is the server API for SubmissionService service.
// All implementations should embed UnimplementedSubmissionServiceServer
// for forward compatibility
type SubmissionServiceServer interface {
	CreateSubmission(context.Context, *CreateSubmissionInput) (*CreateSubmissionOutput, error)
	RetestSubmission(context.Context, *RetestSubmissionInput) (*RetestSubmissionOutput, error)
	DescribeSubmission(context.Context, *DescribeSubmissionInput) (*DescribeSubmissionOutput, error)
}

// UnimplementedSubmissionServiceServer should be embedded to have forward compatible implementations.
type UnimplementedSubmissionServiceServer struct {
}

func (UnimplementedSubmissionServiceServer) CreateSubmission(context.Context, *CreateSubmissionInput) (*CreateSubmissionOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateSubmission not implemented")
}
func (UnimplementedSubmissionServiceServer) RetestSubmission(context.Context, *RetestSubmissionInput) (*RetestSubmissionOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RetestSubmission not implemented")
}
func (UnimplementedSubmissionServiceServer) DescribeSubmission(context.Context, *DescribeSubmissionInput) (*DescribeSubmissionOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DescribeSubmission not implemented")
}

// UnsafeSubmissionServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SubmissionServiceServer will
// result in compilation errors.
type UnsafeSubmissionServiceServer interface {
	mustEmbedUnimplementedSubmissionServiceServer()
}

func RegisterSubmissionServiceServer(s grpc.ServiceRegistrar, srv SubmissionServiceServer) {
	s.RegisterService(&SubmissionService_ServiceDesc, srv)
}

func _SubmissionService_CreateSubmission_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateSubmissionInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SubmissionServiceServer).CreateSubmission(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SubmissionService_CreateSubmission_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SubmissionServiceServer).CreateSubmission(ctx, req.(*CreateSubmissionInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _SubmissionService_RetestSubmission_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RetestSubmissionInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SubmissionServiceServer).RetestSubmission(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SubmissionService_RetestSubmission_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SubmissionServiceServer).RetestSubmission(ctx, req.(*RetestSubmissionInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _SubmissionService_DescribeSubmission_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DescribeSubmissionInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SubmissionServiceServer).DescribeSubmission(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SubmissionService_DescribeSubmission_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SubmissionServiceServer).DescribeSubmission(ctx, req.(*DescribeSubmissionInput))
	}
	return interceptor(ctx, in, info, handler)
}

// SubmissionService_ServiceDesc is the grpc.ServiceDesc for SubmissionService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SubmissionService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "eolymp.atlas.SubmissionService",
	HandlerType: (*SubmissionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateSubmission",
			Handler:    _SubmissionService_CreateSubmission_Handler,
		},
		{
			MethodName: "RetestSubmission",
			Handler:    _SubmissionService_RetestSubmission_Handler,
		},
		{
			MethodName: "DescribeSubmission",
			Handler:    _SubmissionService_DescribeSubmission_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "eolymp/atlas/submission_service.proto",
}
