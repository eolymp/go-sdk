// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v3.21.12
// source: eolymp/atlas/bookmark_service.proto

package atlas

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	BookmarkService_GetBookmark_FullMethodName = "/eolymp.atlas.BookmarkService/GetBookmark"
	BookmarkService_SetBookmark_FullMethodName = "/eolymp.atlas.BookmarkService/SetBookmark"
)

// BookmarkServiceClient is the client API for BookmarkService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BookmarkServiceClient interface {
	GetBookmark(ctx context.Context, in *GetBookmarkInput, opts ...grpc.CallOption) (*GetBookmarkOutput, error)
	SetBookmark(ctx context.Context, in *SetBookmarkInput, opts ...grpc.CallOption) (*SetBookmarkOutput, error)
}

type bookmarkServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBookmarkServiceClient(cc grpc.ClientConnInterface) BookmarkServiceClient {
	return &bookmarkServiceClient{cc}
}

func (c *bookmarkServiceClient) GetBookmark(ctx context.Context, in *GetBookmarkInput, opts ...grpc.CallOption) (*GetBookmarkOutput, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetBookmarkOutput)
	err := c.cc.Invoke(ctx, BookmarkService_GetBookmark_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookmarkServiceClient) SetBookmark(ctx context.Context, in *SetBookmarkInput, opts ...grpc.CallOption) (*SetBookmarkOutput, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SetBookmarkOutput)
	err := c.cc.Invoke(ctx, BookmarkService_SetBookmark_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BookmarkServiceServer is the server API for BookmarkService service.
// All implementations should embed UnimplementedBookmarkServiceServer
// for forward compatibility
type BookmarkServiceServer interface {
	GetBookmark(context.Context, *GetBookmarkInput) (*GetBookmarkOutput, error)
	SetBookmark(context.Context, *SetBookmarkInput) (*SetBookmarkOutput, error)
}

// UnimplementedBookmarkServiceServer should be embedded to have forward compatible implementations.
type UnimplementedBookmarkServiceServer struct {
}

func (UnimplementedBookmarkServiceServer) GetBookmark(context.Context, *GetBookmarkInput) (*GetBookmarkOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBookmark not implemented")
}
func (UnimplementedBookmarkServiceServer) SetBookmark(context.Context, *SetBookmarkInput) (*SetBookmarkOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetBookmark not implemented")
}

// UnsafeBookmarkServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BookmarkServiceServer will
// result in compilation errors.
type UnsafeBookmarkServiceServer interface {
	mustEmbedUnimplementedBookmarkServiceServer()
}

func RegisterBookmarkServiceServer(s grpc.ServiceRegistrar, srv BookmarkServiceServer) {
	s.RegisterService(&BookmarkService_ServiceDesc, srv)
}

func _BookmarkService_GetBookmark_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBookmarkInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookmarkServiceServer).GetBookmark(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BookmarkService_GetBookmark_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookmarkServiceServer).GetBookmark(ctx, req.(*GetBookmarkInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookmarkService_SetBookmark_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetBookmarkInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookmarkServiceServer).SetBookmark(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BookmarkService_SetBookmark_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookmarkServiceServer).SetBookmark(ctx, req.(*SetBookmarkInput))
	}
	return interceptor(ctx, in, info, handler)
}

// BookmarkService_ServiceDesc is the grpc.ServiceDesc for BookmarkService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BookmarkService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "eolymp.atlas.BookmarkService",
	HandlerType: (*BookmarkServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetBookmark",
			Handler:    _BookmarkService_GetBookmark_Handler,
		},
		{
			MethodName: "SetBookmark",
			Handler:    _BookmarkService_SetBookmark_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "eolymp/atlas/bookmark_service.proto",
}
