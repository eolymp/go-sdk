// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.18.1
// source: eolymp/oauth2/oauth2.proto

package oauth2

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

// OAuth2Client is the client API for OAuth2 service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OAuth2Client interface {
	Token(ctx context.Context, in *TokenInput, opts ...grpc.CallOption) (*TokenOutput, error)
	Authorize(ctx context.Context, in *AuthorizeInput, opts ...grpc.CallOption) (*AuthorizeOutput, error)
}

type oAuth2Client struct {
	cc grpc.ClientConnInterface
}

func NewOAuth2Client(cc grpc.ClientConnInterface) OAuth2Client {
	return &oAuth2Client{cc}
}

func (c *oAuth2Client) Token(ctx context.Context, in *TokenInput, opts ...grpc.CallOption) (*TokenOutput, error) {
	out := new(TokenOutput)
	err := c.cc.Invoke(ctx, "/eolymp.oauth2.OAuth2/Token", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *oAuth2Client) Authorize(ctx context.Context, in *AuthorizeInput, opts ...grpc.CallOption) (*AuthorizeOutput, error) {
	out := new(AuthorizeOutput)
	err := c.cc.Invoke(ctx, "/eolymp.oauth2.OAuth2/Authorize", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OAuth2Server is the server API for OAuth2 service.
// All implementations should embed UnimplementedOAuth2Server
// for forward compatibility
type OAuth2Server interface {
	Token(context.Context, *TokenInput) (*TokenOutput, error)
	Authorize(context.Context, *AuthorizeInput) (*AuthorizeOutput, error)
}

// UnimplementedOAuth2Server should be embedded to have forward compatible implementations.
type UnimplementedOAuth2Server struct {
}

func (UnimplementedOAuth2Server) Token(context.Context, *TokenInput) (*TokenOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Token not implemented")
}
func (UnimplementedOAuth2Server) Authorize(context.Context, *AuthorizeInput) (*AuthorizeOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Authorize not implemented")
}

// UnsafeOAuth2Server may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OAuth2Server will
// result in compilation errors.
type UnsafeOAuth2Server interface {
	mustEmbedUnimplementedOAuth2Server()
}

func RegisterOAuth2Server(s grpc.ServiceRegistrar, srv OAuth2Server) {
	s.RegisterService(&OAuth2_ServiceDesc, srv)
}

func _OAuth2_Token_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TokenInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OAuth2Server).Token(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/eolymp.oauth2.OAuth2/Token",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OAuth2Server).Token(ctx, req.(*TokenInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _OAuth2_Authorize_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthorizeInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OAuth2Server).Authorize(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/eolymp.oauth2.OAuth2/Authorize",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OAuth2Server).Authorize(ctx, req.(*AuthorizeInput))
	}
	return interceptor(ctx, in, info, handler)
}

// OAuth2_ServiceDesc is the grpc.ServiceDesc for OAuth2 service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OAuth2_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "eolymp.oauth2.OAuth2",
	HandlerType: (*OAuth2Server)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Token",
			Handler:    _OAuth2_Token_Handler,
		},
		{
			MethodName: "Authorize",
			Handler:    _OAuth2_Authorize_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "eolymp/oauth2/oauth2.proto",
}
