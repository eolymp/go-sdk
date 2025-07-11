// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.4
// source: eolymp/community/activity_graph_service.proto

package community

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
	ActivityGraphService_DescribeActivityGraph_FullMethodName = "/eolymp.community.ActivityGraphService/DescribeActivityGraph"
)

// ActivityGraphServiceClient is the client API for ActivityGraphService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ActivityGraphServiceClient interface {
	DescribeActivityGraph(ctx context.Context, in *DescribeActivityGraphInput, opts ...grpc.CallOption) (*DescribeActivityGraphOutput, error)
}

type activityGraphServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewActivityGraphServiceClient(cc grpc.ClientConnInterface) ActivityGraphServiceClient {
	return &activityGraphServiceClient{cc}
}

func (c *activityGraphServiceClient) DescribeActivityGraph(ctx context.Context, in *DescribeActivityGraphInput, opts ...grpc.CallOption) (*DescribeActivityGraphOutput, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DescribeActivityGraphOutput)
	err := c.cc.Invoke(ctx, ActivityGraphService_DescribeActivityGraph_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ActivityGraphServiceServer is the server API for ActivityGraphService service.
// All implementations should embed UnimplementedActivityGraphServiceServer
// for forward compatibility.
type ActivityGraphServiceServer interface {
	DescribeActivityGraph(context.Context, *DescribeActivityGraphInput) (*DescribeActivityGraphOutput, error)
}

// UnimplementedActivityGraphServiceServer should be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedActivityGraphServiceServer struct{}

func (UnimplementedActivityGraphServiceServer) DescribeActivityGraph(context.Context, *DescribeActivityGraphInput) (*DescribeActivityGraphOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DescribeActivityGraph not implemented")
}
func (UnimplementedActivityGraphServiceServer) testEmbeddedByValue() {}

// UnsafeActivityGraphServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ActivityGraphServiceServer will
// result in compilation errors.
type UnsafeActivityGraphServiceServer interface {
	mustEmbedUnimplementedActivityGraphServiceServer()
}

func RegisterActivityGraphServiceServer(s grpc.ServiceRegistrar, srv ActivityGraphServiceServer) {
	// If the following call pancis, it indicates UnimplementedActivityGraphServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&ActivityGraphService_ServiceDesc, srv)
}

func _ActivityGraphService_DescribeActivityGraph_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DescribeActivityGraphInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActivityGraphServiceServer).DescribeActivityGraph(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ActivityGraphService_DescribeActivityGraph_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActivityGraphServiceServer).DescribeActivityGraph(ctx, req.(*DescribeActivityGraphInput))
	}
	return interceptor(ctx, in, info, handler)
}

// ActivityGraphService_ServiceDesc is the grpc.ServiceDesc for ActivityGraphService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ActivityGraphService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "eolymp.community.ActivityGraphService",
	HandlerType: (*ActivityGraphServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DescribeActivityGraph",
			Handler:    _ActivityGraphService_DescribeActivityGraph_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "eolymp/community/activity_graph_service.proto",
}
