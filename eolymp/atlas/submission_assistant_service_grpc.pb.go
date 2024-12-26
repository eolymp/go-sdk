// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v4.24.4
// source: eolymp/atlas/submission_assistant_service.proto

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
	SubmissionAssistantService_RequestDebugAssistance_FullMethodName  = "/eolymp.atlas.SubmissionAssistantService/RequestDebugAssistance"
	SubmissionAssistantService_DescribeDebugAssistance_FullMethodName = "/eolymp.atlas.SubmissionAssistantService/DescribeDebugAssistance"
	SubmissionAssistantService_RateDebugAssistance_FullMethodName     = "/eolymp.atlas.SubmissionAssistantService/RateDebugAssistance"
)

// SubmissionAssistantServiceClient is the client API for SubmissionAssistantService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SubmissionAssistantServiceClient interface {
	RequestDebugAssistance(ctx context.Context, in *RequestDebugAssistanceInput, opts ...grpc.CallOption) (*RequestDebugAssistanceOutput, error)
	DescribeDebugAssistance(ctx context.Context, in *DescribeDebugAssistanceInput, opts ...grpc.CallOption) (*DescribeDebugAssistanceOutput, error)
	RateDebugAssistance(ctx context.Context, in *RateDebugAssistanceInput, opts ...grpc.CallOption) (*RateDebugAssistanceOutput, error)
}

type submissionAssistantServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSubmissionAssistantServiceClient(cc grpc.ClientConnInterface) SubmissionAssistantServiceClient {
	return &submissionAssistantServiceClient{cc}
}

func (c *submissionAssistantServiceClient) RequestDebugAssistance(ctx context.Context, in *RequestDebugAssistanceInput, opts ...grpc.CallOption) (*RequestDebugAssistanceOutput, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RequestDebugAssistanceOutput)
	err := c.cc.Invoke(ctx, SubmissionAssistantService_RequestDebugAssistance_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *submissionAssistantServiceClient) DescribeDebugAssistance(ctx context.Context, in *DescribeDebugAssistanceInput, opts ...grpc.CallOption) (*DescribeDebugAssistanceOutput, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DescribeDebugAssistanceOutput)
	err := c.cc.Invoke(ctx, SubmissionAssistantService_DescribeDebugAssistance_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *submissionAssistantServiceClient) RateDebugAssistance(ctx context.Context, in *RateDebugAssistanceInput, opts ...grpc.CallOption) (*RateDebugAssistanceOutput, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RateDebugAssistanceOutput)
	err := c.cc.Invoke(ctx, SubmissionAssistantService_RateDebugAssistance_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SubmissionAssistantServiceServer is the server API for SubmissionAssistantService service.
// All implementations should embed UnimplementedSubmissionAssistantServiceServer
// for forward compatibility.
type SubmissionAssistantServiceServer interface {
	RequestDebugAssistance(context.Context, *RequestDebugAssistanceInput) (*RequestDebugAssistanceOutput, error)
	DescribeDebugAssistance(context.Context, *DescribeDebugAssistanceInput) (*DescribeDebugAssistanceOutput, error)
	RateDebugAssistance(context.Context, *RateDebugAssistanceInput) (*RateDebugAssistanceOutput, error)
}

// UnimplementedSubmissionAssistantServiceServer should be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedSubmissionAssistantServiceServer struct{}

func (UnimplementedSubmissionAssistantServiceServer) RequestDebugAssistance(context.Context, *RequestDebugAssistanceInput) (*RequestDebugAssistanceOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RequestDebugAssistance not implemented")
}
func (UnimplementedSubmissionAssistantServiceServer) DescribeDebugAssistance(context.Context, *DescribeDebugAssistanceInput) (*DescribeDebugAssistanceOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DescribeDebugAssistance not implemented")
}
func (UnimplementedSubmissionAssistantServiceServer) RateDebugAssistance(context.Context, *RateDebugAssistanceInput) (*RateDebugAssistanceOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RateDebugAssistance not implemented")
}
func (UnimplementedSubmissionAssistantServiceServer) testEmbeddedByValue() {}

// UnsafeSubmissionAssistantServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SubmissionAssistantServiceServer will
// result in compilation errors.
type UnsafeSubmissionAssistantServiceServer interface {
	mustEmbedUnimplementedSubmissionAssistantServiceServer()
}

func RegisterSubmissionAssistantServiceServer(s grpc.ServiceRegistrar, srv SubmissionAssistantServiceServer) {
	// If the following call pancis, it indicates UnimplementedSubmissionAssistantServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&SubmissionAssistantService_ServiceDesc, srv)
}

func _SubmissionAssistantService_RequestDebugAssistance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestDebugAssistanceInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SubmissionAssistantServiceServer).RequestDebugAssistance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SubmissionAssistantService_RequestDebugAssistance_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SubmissionAssistantServiceServer).RequestDebugAssistance(ctx, req.(*RequestDebugAssistanceInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _SubmissionAssistantService_DescribeDebugAssistance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DescribeDebugAssistanceInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SubmissionAssistantServiceServer).DescribeDebugAssistance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SubmissionAssistantService_DescribeDebugAssistance_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SubmissionAssistantServiceServer).DescribeDebugAssistance(ctx, req.(*DescribeDebugAssistanceInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _SubmissionAssistantService_RateDebugAssistance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RateDebugAssistanceInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SubmissionAssistantServiceServer).RateDebugAssistance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SubmissionAssistantService_RateDebugAssistance_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SubmissionAssistantServiceServer).RateDebugAssistance(ctx, req.(*RateDebugAssistanceInput))
	}
	return interceptor(ctx, in, info, handler)
}

// SubmissionAssistantService_ServiceDesc is the grpc.ServiceDesc for SubmissionAssistantService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SubmissionAssistantService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "eolymp.atlas.SubmissionAssistantService",
	HandlerType: (*SubmissionAssistantServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RequestDebugAssistance",
			Handler:    _SubmissionAssistantService_RequestDebugAssistance_Handler,
		},
		{
			MethodName: "DescribeDebugAssistance",
			Handler:    _SubmissionAssistantService_DescribeDebugAssistance_Handler,
		},
		{
			MethodName: "RateDebugAssistance",
			Handler:    _SubmissionAssistantService_RateDebugAssistance_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "eolymp/atlas/submission_assistant_service.proto",
}
