// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.4
// source: eolymp/judge/scoreboard_service.proto

package judge

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
	ScoreboardService_DescribeScoreboard_FullMethodName = "/eolymp.judge.ScoreboardService/DescribeScoreboard"
	ScoreboardService_ListScoreboardRows_FullMethodName = "/eolymp.judge.ScoreboardService/ListScoreboardRows"
)

// ScoreboardServiceClient is the client API for ScoreboardService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ScoreboardServiceClient interface {
	DescribeScoreboard(ctx context.Context, in *DescribeScoreboardInput, opts ...grpc.CallOption) (*DescribeScoreboardOutput, error)
	ListScoreboardRows(ctx context.Context, in *ListScoreboardRowsInput, opts ...grpc.CallOption) (*ListScoreboardRowsOutput, error)
}

type scoreboardServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewScoreboardServiceClient(cc grpc.ClientConnInterface) ScoreboardServiceClient {
	return &scoreboardServiceClient{cc}
}

func (c *scoreboardServiceClient) DescribeScoreboard(ctx context.Context, in *DescribeScoreboardInput, opts ...grpc.CallOption) (*DescribeScoreboardOutput, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DescribeScoreboardOutput)
	err := c.cc.Invoke(ctx, ScoreboardService_DescribeScoreboard_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *scoreboardServiceClient) ListScoreboardRows(ctx context.Context, in *ListScoreboardRowsInput, opts ...grpc.CallOption) (*ListScoreboardRowsOutput, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListScoreboardRowsOutput)
	err := c.cc.Invoke(ctx, ScoreboardService_ListScoreboardRows_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ScoreboardServiceServer is the server API for ScoreboardService service.
// All implementations should embed UnimplementedScoreboardServiceServer
// for forward compatibility.
type ScoreboardServiceServer interface {
	DescribeScoreboard(context.Context, *DescribeScoreboardInput) (*DescribeScoreboardOutput, error)
	ListScoreboardRows(context.Context, *ListScoreboardRowsInput) (*ListScoreboardRowsOutput, error)
}

// UnimplementedScoreboardServiceServer should be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedScoreboardServiceServer struct{}

func (UnimplementedScoreboardServiceServer) DescribeScoreboard(context.Context, *DescribeScoreboardInput) (*DescribeScoreboardOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DescribeScoreboard not implemented")
}
func (UnimplementedScoreboardServiceServer) ListScoreboardRows(context.Context, *ListScoreboardRowsInput) (*ListScoreboardRowsOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListScoreboardRows not implemented")
}
func (UnimplementedScoreboardServiceServer) testEmbeddedByValue() {}

// UnsafeScoreboardServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ScoreboardServiceServer will
// result in compilation errors.
type UnsafeScoreboardServiceServer interface {
	mustEmbedUnimplementedScoreboardServiceServer()
}

func RegisterScoreboardServiceServer(s grpc.ServiceRegistrar, srv ScoreboardServiceServer) {
	// If the following call pancis, it indicates UnimplementedScoreboardServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&ScoreboardService_ServiceDesc, srv)
}

func _ScoreboardService_DescribeScoreboard_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DescribeScoreboardInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ScoreboardServiceServer).DescribeScoreboard(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ScoreboardService_DescribeScoreboard_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ScoreboardServiceServer).DescribeScoreboard(ctx, req.(*DescribeScoreboardInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _ScoreboardService_ListScoreboardRows_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListScoreboardRowsInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ScoreboardServiceServer).ListScoreboardRows(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ScoreboardService_ListScoreboardRows_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ScoreboardServiceServer).ListScoreboardRows(ctx, req.(*ListScoreboardRowsInput))
	}
	return interceptor(ctx, in, info, handler)
}

// ScoreboardService_ServiceDesc is the grpc.ServiceDesc for ScoreboardService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ScoreboardService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "eolymp.judge.ScoreboardService",
	HandlerType: (*ScoreboardServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DescribeScoreboard",
			Handler:    _ScoreboardService_DescribeScoreboard_Handler,
		},
		{
			MethodName: "ListScoreboardRows",
			Handler:    _ScoreboardService_ListScoreboardRows_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "eolymp/judge/scoreboard_service.proto",
}
