// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v4.24.4
// source: eolymp/community/rating_service.proto

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
	RatingService_SetRating_FullMethodName    = "/eolymp.community.RatingService/SetRating"
	RatingService_DeleteRating_FullMethodName = "/eolymp.community.RatingService/DeleteRating"
)

// RatingServiceClient is the client API for RatingService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RatingServiceClient interface {
	// Create or Update rating point
	// If point with the same id or ref exists, it will be updated, otherwise created.
	SetRating(ctx context.Context, in *SetRatingInput, opts ...grpc.CallOption) (*SetRatingOutput, error)
	DeleteRating(ctx context.Context, in *DeleteRatingInput, opts ...grpc.CallOption) (*DeleteRatingOutput, error)
}

type ratingServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRatingServiceClient(cc grpc.ClientConnInterface) RatingServiceClient {
	return &ratingServiceClient{cc}
}

func (c *ratingServiceClient) SetRating(ctx context.Context, in *SetRatingInput, opts ...grpc.CallOption) (*SetRatingOutput, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SetRatingOutput)
	err := c.cc.Invoke(ctx, RatingService_SetRating_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ratingServiceClient) DeleteRating(ctx context.Context, in *DeleteRatingInput, opts ...grpc.CallOption) (*DeleteRatingOutput, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteRatingOutput)
	err := c.cc.Invoke(ctx, RatingService_DeleteRating_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RatingServiceServer is the server API for RatingService service.
// All implementations should embed UnimplementedRatingServiceServer
// for forward compatibility.
type RatingServiceServer interface {
	// Create or Update rating point
	// If point with the same id or ref exists, it will be updated, otherwise created.
	SetRating(context.Context, *SetRatingInput) (*SetRatingOutput, error)
	DeleteRating(context.Context, *DeleteRatingInput) (*DeleteRatingOutput, error)
}

// UnimplementedRatingServiceServer should be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedRatingServiceServer struct{}

func (UnimplementedRatingServiceServer) SetRating(context.Context, *SetRatingInput) (*SetRatingOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetRating not implemented")
}
func (UnimplementedRatingServiceServer) DeleteRating(context.Context, *DeleteRatingInput) (*DeleteRatingOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteRating not implemented")
}
func (UnimplementedRatingServiceServer) testEmbeddedByValue() {}

// UnsafeRatingServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RatingServiceServer will
// result in compilation errors.
type UnsafeRatingServiceServer interface {
	mustEmbedUnimplementedRatingServiceServer()
}

func RegisterRatingServiceServer(s grpc.ServiceRegistrar, srv RatingServiceServer) {
	// If the following call pancis, it indicates UnimplementedRatingServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&RatingService_ServiceDesc, srv)
}

func _RatingService_SetRating_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetRatingInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RatingServiceServer).SetRating(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RatingService_SetRating_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RatingServiceServer).SetRating(ctx, req.(*SetRatingInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _RatingService_DeleteRating_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRatingInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RatingServiceServer).DeleteRating(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RatingService_DeleteRating_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RatingServiceServer).DeleteRating(ctx, req.(*DeleteRatingInput))
	}
	return interceptor(ctx, in, info, handler)
}

// RatingService_ServiceDesc is the grpc.ServiceDesc for RatingService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RatingService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "eolymp.community.RatingService",
	HandlerType: (*RatingServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SetRating",
			Handler:    _RatingService_SetRating_Handler,
		},
		{
			MethodName: "DeleteRating",
			Handler:    _RatingService_DeleteRating_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "eolymp/community/rating_service.proto",
}
