// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.18.1
// source: eolymp/worker/worker_service.proto

package worker

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
	Worker_CreateJob_FullMethodName   = "/eolymp.worker.Worker/CreateJob"
	Worker_DescribeJob_FullMethodName = "/eolymp.worker.Worker/DescribeJob"
	Worker_ListJobs_FullMethodName    = "/eolymp.worker.Worker/ListJobs"
)

// WorkerClient is the client API for Worker service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WorkerClient interface {
	CreateJob(ctx context.Context, in *CreateJobInput, opts ...grpc.CallOption) (*CreateJobOutput, error)
	DescribeJob(ctx context.Context, in *DescribeJobInput, opts ...grpc.CallOption) (*DescribeJobOutput, error)
	ListJobs(ctx context.Context, in *ListJobsInput, opts ...grpc.CallOption) (*ListJobsOutput, error)
}

type workerClient struct {
	cc grpc.ClientConnInterface
}

func NewWorkerClient(cc grpc.ClientConnInterface) WorkerClient {
	return &workerClient{cc}
}

func (c *workerClient) CreateJob(ctx context.Context, in *CreateJobInput, opts ...grpc.CallOption) (*CreateJobOutput, error) {
	out := new(CreateJobOutput)
	err := c.cc.Invoke(ctx, Worker_CreateJob_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workerClient) DescribeJob(ctx context.Context, in *DescribeJobInput, opts ...grpc.CallOption) (*DescribeJobOutput, error) {
	out := new(DescribeJobOutput)
	err := c.cc.Invoke(ctx, Worker_DescribeJob_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workerClient) ListJobs(ctx context.Context, in *ListJobsInput, opts ...grpc.CallOption) (*ListJobsOutput, error) {
	out := new(ListJobsOutput)
	err := c.cc.Invoke(ctx, Worker_ListJobs_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WorkerServer is the server API for Worker service.
// All implementations should embed UnimplementedWorkerServer
// for forward compatibility
type WorkerServer interface {
	CreateJob(context.Context, *CreateJobInput) (*CreateJobOutput, error)
	DescribeJob(context.Context, *DescribeJobInput) (*DescribeJobOutput, error)
	ListJobs(context.Context, *ListJobsInput) (*ListJobsOutput, error)
}

// UnimplementedWorkerServer should be embedded to have forward compatible implementations.
type UnimplementedWorkerServer struct {
}

func (UnimplementedWorkerServer) CreateJob(context.Context, *CreateJobInput) (*CreateJobOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateJob not implemented")
}
func (UnimplementedWorkerServer) DescribeJob(context.Context, *DescribeJobInput) (*DescribeJobOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DescribeJob not implemented")
}
func (UnimplementedWorkerServer) ListJobs(context.Context, *ListJobsInput) (*ListJobsOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListJobs not implemented")
}

// UnsafeWorkerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WorkerServer will
// result in compilation errors.
type UnsafeWorkerServer interface {
	mustEmbedUnimplementedWorkerServer()
}

func RegisterWorkerServer(s grpc.ServiceRegistrar, srv WorkerServer) {
	s.RegisterService(&Worker_ServiceDesc, srv)
}

func _Worker_CreateJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateJobInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkerServer).CreateJob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Worker_CreateJob_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkerServer).CreateJob(ctx, req.(*CreateJobInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _Worker_DescribeJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DescribeJobInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkerServer).DescribeJob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Worker_DescribeJob_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkerServer).DescribeJob(ctx, req.(*DescribeJobInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _Worker_ListJobs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListJobsInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkerServer).ListJobs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Worker_ListJobs_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkerServer).ListJobs(ctx, req.(*ListJobsInput))
	}
	return interceptor(ctx, in, info, handler)
}

// Worker_ServiceDesc is the grpc.ServiceDesc for Worker service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Worker_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "eolymp.worker.Worker",
	HandlerType: (*WorkerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateJob",
			Handler:    _Worker_CreateJob_Handler,
		},
		{
			MethodName: "DescribeJob",
			Handler:    _Worker_DescribeJob_Handler,
		},
		{
			MethodName: "ListJobs",
			Handler:    _Worker_ListJobs_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "eolymp/worker/worker_service.proto",
}

const (
	WorkerService_CreateJob_FullMethodName   = "/eolymp.worker.WorkerService/CreateJob"
	WorkerService_DescribeJob_FullMethodName = "/eolymp.worker.WorkerService/DescribeJob"
	WorkerService_ListJobs_FullMethodName    = "/eolymp.worker.WorkerService/ListJobs"
)

// WorkerServiceClient is the client API for WorkerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WorkerServiceClient interface {
	CreateJob(ctx context.Context, in *CreateJobInput, opts ...grpc.CallOption) (*CreateJobOutput, error)
	DescribeJob(ctx context.Context, in *DescribeJobInput, opts ...grpc.CallOption) (*DescribeJobOutput, error)
	ListJobs(ctx context.Context, in *ListJobsInput, opts ...grpc.CallOption) (*ListJobsOutput, error)
}

type workerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewWorkerServiceClient(cc grpc.ClientConnInterface) WorkerServiceClient {
	return &workerServiceClient{cc}
}

func (c *workerServiceClient) CreateJob(ctx context.Context, in *CreateJobInput, opts ...grpc.CallOption) (*CreateJobOutput, error) {
	out := new(CreateJobOutput)
	err := c.cc.Invoke(ctx, WorkerService_CreateJob_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workerServiceClient) DescribeJob(ctx context.Context, in *DescribeJobInput, opts ...grpc.CallOption) (*DescribeJobOutput, error) {
	out := new(DescribeJobOutput)
	err := c.cc.Invoke(ctx, WorkerService_DescribeJob_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workerServiceClient) ListJobs(ctx context.Context, in *ListJobsInput, opts ...grpc.CallOption) (*ListJobsOutput, error) {
	out := new(ListJobsOutput)
	err := c.cc.Invoke(ctx, WorkerService_ListJobs_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WorkerServiceServer is the server API for WorkerService service.
// All implementations should embed UnimplementedWorkerServiceServer
// for forward compatibility
type WorkerServiceServer interface {
	CreateJob(context.Context, *CreateJobInput) (*CreateJobOutput, error)
	DescribeJob(context.Context, *DescribeJobInput) (*DescribeJobOutput, error)
	ListJobs(context.Context, *ListJobsInput) (*ListJobsOutput, error)
}

// UnimplementedWorkerServiceServer should be embedded to have forward compatible implementations.
type UnimplementedWorkerServiceServer struct {
}

func (UnimplementedWorkerServiceServer) CreateJob(context.Context, *CreateJobInput) (*CreateJobOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateJob not implemented")
}
func (UnimplementedWorkerServiceServer) DescribeJob(context.Context, *DescribeJobInput) (*DescribeJobOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DescribeJob not implemented")
}
func (UnimplementedWorkerServiceServer) ListJobs(context.Context, *ListJobsInput) (*ListJobsOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListJobs not implemented")
}

// UnsafeWorkerServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WorkerServiceServer will
// result in compilation errors.
type UnsafeWorkerServiceServer interface {
	mustEmbedUnimplementedWorkerServiceServer()
}

func RegisterWorkerServiceServer(s grpc.ServiceRegistrar, srv WorkerServiceServer) {
	s.RegisterService(&WorkerService_ServiceDesc, srv)
}

func _WorkerService_CreateJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateJobInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkerServiceServer).CreateJob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WorkerService_CreateJob_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkerServiceServer).CreateJob(ctx, req.(*CreateJobInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _WorkerService_DescribeJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DescribeJobInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkerServiceServer).DescribeJob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WorkerService_DescribeJob_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkerServiceServer).DescribeJob(ctx, req.(*DescribeJobInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _WorkerService_ListJobs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListJobsInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkerServiceServer).ListJobs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WorkerService_ListJobs_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkerServiceServer).ListJobs(ctx, req.(*ListJobsInput))
	}
	return interceptor(ctx, in, info, handler)
}

// WorkerService_ServiceDesc is the grpc.ServiceDesc for WorkerService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var WorkerService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "eolymp.worker.WorkerService",
	HandlerType: (*WorkerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateJob",
			Handler:    _WorkerService_CreateJob_Handler,
		},
		{
			MethodName: "DescribeJob",
			Handler:    _WorkerService_DescribeJob_Handler,
		},
		{
			MethodName: "ListJobs",
			Handler:    _WorkerService_ListJobs_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "eolymp/worker/worker_service.proto",
}
