// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.18.1
// source: eolymp/judge/score_service.proto

package judge

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
	ScoreService_IntrospectScore_FullMethodName = "/eolymp.judge.ScoreService/IntrospectScore"
	ScoreService_WatchScore_FullMethodName      = "/eolymp.judge.ScoreService/WatchScore"
	ScoreService_DescribeScore_FullMethodName   = "/eolymp.judge.ScoreService/DescribeScore"
	ScoreService_ImportScore_FullMethodName     = "/eolymp.judge.ScoreService/ImportScore"
	ScoreService_ExportScore_FullMethodName     = "/eolymp.judge.ScoreService/ExportScore"
	ScoreService_ListResult_FullMethodName      = "/eolymp.judge.ScoreService/ListResult"
	ScoreService_RebuildScore_FullMethodName    = "/eolymp.judge.ScoreService/RebuildScore"
)

// ScoreServiceClient is the client API for ScoreService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ScoreServiceClient interface {
	IntrospectScore(ctx context.Context, in *IntrospectScoreInput, opts ...grpc.CallOption) (*IntrospectScoreOutput, error)
	WatchScore(ctx context.Context, in *WatchScoreInput, opts ...grpc.CallOption) (ScoreService_WatchScoreClient, error)
	DescribeScore(ctx context.Context, in *DescribeScoreInput, opts ...grpc.CallOption) (*DescribeScoreOutput, error)
	// ImportScore for ghost participants
	ImportScore(ctx context.Context, in *ImportScoreInput, opts ...grpc.CallOption) (*ImportScoreOutput, error)
	// ExportScore for ghost participants
	ExportScore(ctx context.Context, in *ExportScoreInput, opts ...grpc.CallOption) (*ExportScoreOutput, error)
	// ListResult retrieves scoreboard
	ListResult(ctx context.Context, in *ListResultInput, opts ...grpc.CallOption) (*ListResultOutput, error)
	// Rebuild scoreboard
	RebuildScore(ctx context.Context, in *RebuildScoreInput, opts ...grpc.CallOption) (*RebuildScoreOutput, error)
}

type scoreServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewScoreServiceClient(cc grpc.ClientConnInterface) ScoreServiceClient {
	return &scoreServiceClient{cc}
}

func (c *scoreServiceClient) IntrospectScore(ctx context.Context, in *IntrospectScoreInput, opts ...grpc.CallOption) (*IntrospectScoreOutput, error) {
	out := new(IntrospectScoreOutput)
	err := c.cc.Invoke(ctx, ScoreService_IntrospectScore_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *scoreServiceClient) WatchScore(ctx context.Context, in *WatchScoreInput, opts ...grpc.CallOption) (ScoreService_WatchScoreClient, error) {
	stream, err := c.cc.NewStream(ctx, &ScoreService_ServiceDesc.Streams[0], ScoreService_WatchScore_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &scoreServiceWatchScoreClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ScoreService_WatchScoreClient interface {
	Recv() (*WatchScoreOutput, error)
	grpc.ClientStream
}

type scoreServiceWatchScoreClient struct {
	grpc.ClientStream
}

func (x *scoreServiceWatchScoreClient) Recv() (*WatchScoreOutput, error) {
	m := new(WatchScoreOutput)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *scoreServiceClient) DescribeScore(ctx context.Context, in *DescribeScoreInput, opts ...grpc.CallOption) (*DescribeScoreOutput, error) {
	out := new(DescribeScoreOutput)
	err := c.cc.Invoke(ctx, ScoreService_DescribeScore_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *scoreServiceClient) ImportScore(ctx context.Context, in *ImportScoreInput, opts ...grpc.CallOption) (*ImportScoreOutput, error) {
	out := new(ImportScoreOutput)
	err := c.cc.Invoke(ctx, ScoreService_ImportScore_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *scoreServiceClient) ExportScore(ctx context.Context, in *ExportScoreInput, opts ...grpc.CallOption) (*ExportScoreOutput, error) {
	out := new(ExportScoreOutput)
	err := c.cc.Invoke(ctx, ScoreService_ExportScore_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *scoreServiceClient) ListResult(ctx context.Context, in *ListResultInput, opts ...grpc.CallOption) (*ListResultOutput, error) {
	out := new(ListResultOutput)
	err := c.cc.Invoke(ctx, ScoreService_ListResult_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *scoreServiceClient) RebuildScore(ctx context.Context, in *RebuildScoreInput, opts ...grpc.CallOption) (*RebuildScoreOutput, error) {
	out := new(RebuildScoreOutput)
	err := c.cc.Invoke(ctx, ScoreService_RebuildScore_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ScoreServiceServer is the server API for ScoreService service.
// All implementations should embed UnimplementedScoreServiceServer
// for forward compatibility
type ScoreServiceServer interface {
	IntrospectScore(context.Context, *IntrospectScoreInput) (*IntrospectScoreOutput, error)
	WatchScore(*WatchScoreInput, ScoreService_WatchScoreServer) error
	DescribeScore(context.Context, *DescribeScoreInput) (*DescribeScoreOutput, error)
	// ImportScore for ghost participants
	ImportScore(context.Context, *ImportScoreInput) (*ImportScoreOutput, error)
	// ExportScore for ghost participants
	ExportScore(context.Context, *ExportScoreInput) (*ExportScoreOutput, error)
	// ListResult retrieves scoreboard
	ListResult(context.Context, *ListResultInput) (*ListResultOutput, error)
	// Rebuild scoreboard
	RebuildScore(context.Context, *RebuildScoreInput) (*RebuildScoreOutput, error)
}

// UnimplementedScoreServiceServer should be embedded to have forward compatible implementations.
type UnimplementedScoreServiceServer struct {
}

func (UnimplementedScoreServiceServer) IntrospectScore(context.Context, *IntrospectScoreInput) (*IntrospectScoreOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IntrospectScore not implemented")
}
func (UnimplementedScoreServiceServer) WatchScore(*WatchScoreInput, ScoreService_WatchScoreServer) error {
	return status.Errorf(codes.Unimplemented, "method WatchScore not implemented")
}
func (UnimplementedScoreServiceServer) DescribeScore(context.Context, *DescribeScoreInput) (*DescribeScoreOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DescribeScore not implemented")
}
func (UnimplementedScoreServiceServer) ImportScore(context.Context, *ImportScoreInput) (*ImportScoreOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ImportScore not implemented")
}
func (UnimplementedScoreServiceServer) ExportScore(context.Context, *ExportScoreInput) (*ExportScoreOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExportScore not implemented")
}
func (UnimplementedScoreServiceServer) ListResult(context.Context, *ListResultInput) (*ListResultOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListResult not implemented")
}
func (UnimplementedScoreServiceServer) RebuildScore(context.Context, *RebuildScoreInput) (*RebuildScoreOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RebuildScore not implemented")
}

// UnsafeScoreServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ScoreServiceServer will
// result in compilation errors.
type UnsafeScoreServiceServer interface {
	mustEmbedUnimplementedScoreServiceServer()
}

func RegisterScoreServiceServer(s grpc.ServiceRegistrar, srv ScoreServiceServer) {
	s.RegisterService(&ScoreService_ServiceDesc, srv)
}

func _ScoreService_IntrospectScore_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IntrospectScoreInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ScoreServiceServer).IntrospectScore(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ScoreService_IntrospectScore_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ScoreServiceServer).IntrospectScore(ctx, req.(*IntrospectScoreInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _ScoreService_WatchScore_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(WatchScoreInput)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ScoreServiceServer).WatchScore(m, &scoreServiceWatchScoreServer{stream})
}

type ScoreService_WatchScoreServer interface {
	Send(*WatchScoreOutput) error
	grpc.ServerStream
}

type scoreServiceWatchScoreServer struct {
	grpc.ServerStream
}

func (x *scoreServiceWatchScoreServer) Send(m *WatchScoreOutput) error {
	return x.ServerStream.SendMsg(m)
}

func _ScoreService_DescribeScore_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DescribeScoreInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ScoreServiceServer).DescribeScore(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ScoreService_DescribeScore_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ScoreServiceServer).DescribeScore(ctx, req.(*DescribeScoreInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _ScoreService_ImportScore_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ImportScoreInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ScoreServiceServer).ImportScore(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ScoreService_ImportScore_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ScoreServiceServer).ImportScore(ctx, req.(*ImportScoreInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _ScoreService_ExportScore_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExportScoreInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ScoreServiceServer).ExportScore(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ScoreService_ExportScore_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ScoreServiceServer).ExportScore(ctx, req.(*ExportScoreInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _ScoreService_ListResult_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListResultInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ScoreServiceServer).ListResult(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ScoreService_ListResult_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ScoreServiceServer).ListResult(ctx, req.(*ListResultInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _ScoreService_RebuildScore_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RebuildScoreInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ScoreServiceServer).RebuildScore(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ScoreService_RebuildScore_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ScoreServiceServer).RebuildScore(ctx, req.(*RebuildScoreInput))
	}
	return interceptor(ctx, in, info, handler)
}

// ScoreService_ServiceDesc is the grpc.ServiceDesc for ScoreService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ScoreService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "eolymp.judge.ScoreService",
	HandlerType: (*ScoreServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "IntrospectScore",
			Handler:    _ScoreService_IntrospectScore_Handler,
		},
		{
			MethodName: "DescribeScore",
			Handler:    _ScoreService_DescribeScore_Handler,
		},
		{
			MethodName: "ImportScore",
			Handler:    _ScoreService_ImportScore_Handler,
		},
		{
			MethodName: "ExportScore",
			Handler:    _ScoreService_ExportScore_Handler,
		},
		{
			MethodName: "ListResult",
			Handler:    _ScoreService_ListResult_Handler,
		},
		{
			MethodName: "RebuildScore",
			Handler:    _ScoreService_RebuildScore_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "WatchScore",
			Handler:       _ScoreService_WatchScore_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "eolymp/judge/score_service.proto",
}
