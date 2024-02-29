// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.18.1
// source: eolymp/feed/feed_service.proto

package feed

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
	FeedService_ListEntries_FullMethodName = "/eolymp.feed.FeedService/ListEntries"
)

// FeedServiceClient is the client API for FeedService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FeedServiceClient interface {
	ListEntries(ctx context.Context, in *ListEntriesInput, opts ...grpc.CallOption) (*ListEntriesOutput, error)
}

type feedServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewFeedServiceClient(cc grpc.ClientConnInterface) FeedServiceClient {
	return &feedServiceClient{cc}
}

func (c *feedServiceClient) ListEntries(ctx context.Context, in *ListEntriesInput, opts ...grpc.CallOption) (*ListEntriesOutput, error) {
	out := new(ListEntriesOutput)
	err := c.cc.Invoke(ctx, FeedService_ListEntries_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FeedServiceServer is the server API for FeedService service.
// All implementations should embed UnimplementedFeedServiceServer
// for forward compatibility
type FeedServiceServer interface {
	ListEntries(context.Context, *ListEntriesInput) (*ListEntriesOutput, error)
}

// UnimplementedFeedServiceServer should be embedded to have forward compatible implementations.
type UnimplementedFeedServiceServer struct {
}

func (UnimplementedFeedServiceServer) ListEntries(context.Context, *ListEntriesInput) (*ListEntriesOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListEntries not implemented")
}

// UnsafeFeedServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FeedServiceServer will
// result in compilation errors.
type UnsafeFeedServiceServer interface {
	mustEmbedUnimplementedFeedServiceServer()
}

func RegisterFeedServiceServer(s grpc.ServiceRegistrar, srv FeedServiceServer) {
	s.RegisterService(&FeedService_ServiceDesc, srv)
}

func _FeedService_ListEntries_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListEntriesInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FeedServiceServer).ListEntries(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FeedService_ListEntries_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FeedServiceServer).ListEntries(ctx, req.(*ListEntriesInput))
	}
	return interceptor(ctx, in, info, handler)
}

// FeedService_ServiceDesc is the grpc.ServiceDesc for FeedService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FeedService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "eolymp.feed.FeedService",
	HandlerType: (*FeedServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListEntries",
			Handler:    _FeedService_ListEntries_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "eolymp/feed/feed_service.proto",
}