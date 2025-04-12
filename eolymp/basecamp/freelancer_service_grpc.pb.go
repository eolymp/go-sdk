// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v4.24.4
// source: eolymp/basecamp/freelancer_service.proto

package basecamp

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
	FreelancerService_DescribeFreelancerStatus_FullMethodName = "/eolymp.basecamp.FreelancerService/DescribeFreelancerStatus"
	FreelancerService_UpdateFreelancerStatus_FullMethodName   = "/eolymp.basecamp.FreelancerService/UpdateFreelancerStatus"
)

// FreelancerServiceClient is the client API for FreelancerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FreelancerServiceClient interface {
	DescribeFreelancerStatus(ctx context.Context, in *DescribeFreelancerStatusInput, opts ...grpc.CallOption) (*DescribeFreelancerStatusOutput, error)
	UpdateFreelancerStatus(ctx context.Context, in *UpdateFreelancerStatusInput, opts ...grpc.CallOption) (*UpdateFreelancerStatusOutput, error)
}

type freelancerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewFreelancerServiceClient(cc grpc.ClientConnInterface) FreelancerServiceClient {
	return &freelancerServiceClient{cc}
}

func (c *freelancerServiceClient) DescribeFreelancerStatus(ctx context.Context, in *DescribeFreelancerStatusInput, opts ...grpc.CallOption) (*DescribeFreelancerStatusOutput, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DescribeFreelancerStatusOutput)
	err := c.cc.Invoke(ctx, FreelancerService_DescribeFreelancerStatus_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *freelancerServiceClient) UpdateFreelancerStatus(ctx context.Context, in *UpdateFreelancerStatusInput, opts ...grpc.CallOption) (*UpdateFreelancerStatusOutput, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateFreelancerStatusOutput)
	err := c.cc.Invoke(ctx, FreelancerService_UpdateFreelancerStatus_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FreelancerServiceServer is the server API for FreelancerService service.
// All implementations should embed UnimplementedFreelancerServiceServer
// for forward compatibility.
type FreelancerServiceServer interface {
	DescribeFreelancerStatus(context.Context, *DescribeFreelancerStatusInput) (*DescribeFreelancerStatusOutput, error)
	UpdateFreelancerStatus(context.Context, *UpdateFreelancerStatusInput) (*UpdateFreelancerStatusOutput, error)
}

// UnimplementedFreelancerServiceServer should be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedFreelancerServiceServer struct{}

func (UnimplementedFreelancerServiceServer) DescribeFreelancerStatus(context.Context, *DescribeFreelancerStatusInput) (*DescribeFreelancerStatusOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DescribeFreelancerStatus not implemented")
}
func (UnimplementedFreelancerServiceServer) UpdateFreelancerStatus(context.Context, *UpdateFreelancerStatusInput) (*UpdateFreelancerStatusOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateFreelancerStatus not implemented")
}
func (UnimplementedFreelancerServiceServer) testEmbeddedByValue() {}

// UnsafeFreelancerServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FreelancerServiceServer will
// result in compilation errors.
type UnsafeFreelancerServiceServer interface {
	mustEmbedUnimplementedFreelancerServiceServer()
}

func RegisterFreelancerServiceServer(s grpc.ServiceRegistrar, srv FreelancerServiceServer) {
	// If the following call pancis, it indicates UnimplementedFreelancerServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&FreelancerService_ServiceDesc, srv)
}

func _FreelancerService_DescribeFreelancerStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DescribeFreelancerStatusInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FreelancerServiceServer).DescribeFreelancerStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FreelancerService_DescribeFreelancerStatus_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FreelancerServiceServer).DescribeFreelancerStatus(ctx, req.(*DescribeFreelancerStatusInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _FreelancerService_UpdateFreelancerStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateFreelancerStatusInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FreelancerServiceServer).UpdateFreelancerStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FreelancerService_UpdateFreelancerStatus_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FreelancerServiceServer).UpdateFreelancerStatus(ctx, req.(*UpdateFreelancerStatusInput))
	}
	return interceptor(ctx, in, info, handler)
}

// FreelancerService_ServiceDesc is the grpc.ServiceDesc for FreelancerService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FreelancerService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "eolymp.basecamp.FreelancerService",
	HandlerType: (*FreelancerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DescribeFreelancerStatus",
			Handler:    _FreelancerService_DescribeFreelancerStatus_Handler,
		},
		{
			MethodName: "UpdateFreelancerStatus",
			Handler:    _FreelancerService_UpdateFreelancerStatus_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "eolymp/basecamp/freelancer_service.proto",
}
