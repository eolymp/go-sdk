// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
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

const (
	OAuth2_Token_FullMethodName      = "/eolymp.oauth2.OAuth2/Token"
	OAuth2_UserInfo_FullMethodName   = "/eolymp.oauth2.OAuth2/UserInfo"
	OAuth2_Introspect_FullMethodName = "/eolymp.oauth2.OAuth2/Introspect"
	OAuth2_Revoke_FullMethodName     = "/eolymp.oauth2.OAuth2/Revoke"
	OAuth2_AuthCode_FullMethodName   = "/eolymp.oauth2.OAuth2/AuthCode"
	OAuth2_Authorize_FullMethodName  = "/eolymp.oauth2.OAuth2/Authorize"
	OAuth2_Callback_FullMethodName   = "/eolymp.oauth2.OAuth2/Callback"
)

// OAuth2Client is the client API for OAuth2 service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OAuth2Client interface {
	// issue a new token
	Token(ctx context.Context, in *TokenInput, opts ...grpc.CallOption) (*TokenOutput, error)
	// fetch identity claims for authenticated user
	UserInfo(ctx context.Context, in *UserInfoInput, opts ...grpc.CallOption) (*UserInfoOutput, error)
	// fetch access token data for authenticated user
	Introspect(ctx context.Context, in *IntrospectInput, opts ...grpc.CallOption) (*IntrospectOutput, error)
	// revokes a given token
	Revoke(ctx context.Context, in *RevokeInput, opts ...grpc.CallOption) (*RevokeOutput, error)
	// issues a new authorization code for authenticated user
	AuthCode(ctx context.Context, in *AuthCodeInput, opts ...grpc.CallOption) (*AuthCodeOutput, error)
	// performs redirect to identity provider
	Authorize(ctx context.Context, in *AuthorizeInput, opts ...grpc.CallOption) (*AuthorizeOutput, error)
	// callback for identity provider
	Callback(ctx context.Context, in *CallbackInput, opts ...grpc.CallOption) (*CallbackOutput, error)
}

type oAuth2Client struct {
	cc grpc.ClientConnInterface
}

func NewOAuth2Client(cc grpc.ClientConnInterface) OAuth2Client {
	return &oAuth2Client{cc}
}

func (c *oAuth2Client) Token(ctx context.Context, in *TokenInput, opts ...grpc.CallOption) (*TokenOutput, error) {
	out := new(TokenOutput)
	err := c.cc.Invoke(ctx, OAuth2_Token_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *oAuth2Client) UserInfo(ctx context.Context, in *UserInfoInput, opts ...grpc.CallOption) (*UserInfoOutput, error) {
	out := new(UserInfoOutput)
	err := c.cc.Invoke(ctx, OAuth2_UserInfo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *oAuth2Client) Introspect(ctx context.Context, in *IntrospectInput, opts ...grpc.CallOption) (*IntrospectOutput, error) {
	out := new(IntrospectOutput)
	err := c.cc.Invoke(ctx, OAuth2_Introspect_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *oAuth2Client) Revoke(ctx context.Context, in *RevokeInput, opts ...grpc.CallOption) (*RevokeOutput, error) {
	out := new(RevokeOutput)
	err := c.cc.Invoke(ctx, OAuth2_Revoke_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *oAuth2Client) AuthCode(ctx context.Context, in *AuthCodeInput, opts ...grpc.CallOption) (*AuthCodeOutput, error) {
	out := new(AuthCodeOutput)
	err := c.cc.Invoke(ctx, OAuth2_AuthCode_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *oAuth2Client) Authorize(ctx context.Context, in *AuthorizeInput, opts ...grpc.CallOption) (*AuthorizeOutput, error) {
	out := new(AuthorizeOutput)
	err := c.cc.Invoke(ctx, OAuth2_Authorize_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *oAuth2Client) Callback(ctx context.Context, in *CallbackInput, opts ...grpc.CallOption) (*CallbackOutput, error) {
	out := new(CallbackOutput)
	err := c.cc.Invoke(ctx, OAuth2_Callback_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OAuth2Server is the server API for OAuth2 service.
// All implementations should embed UnimplementedOAuth2Server
// for forward compatibility
type OAuth2Server interface {
	// issue a new token
	Token(context.Context, *TokenInput) (*TokenOutput, error)
	// fetch identity claims for authenticated user
	UserInfo(context.Context, *UserInfoInput) (*UserInfoOutput, error)
	// fetch access token data for authenticated user
	Introspect(context.Context, *IntrospectInput) (*IntrospectOutput, error)
	// revokes a given token
	Revoke(context.Context, *RevokeInput) (*RevokeOutput, error)
	// issues a new authorization code for authenticated user
	AuthCode(context.Context, *AuthCodeInput) (*AuthCodeOutput, error)
	// performs redirect to identity provider
	Authorize(context.Context, *AuthorizeInput) (*AuthorizeOutput, error)
	// callback for identity provider
	Callback(context.Context, *CallbackInput) (*CallbackOutput, error)
}

// UnimplementedOAuth2Server should be embedded to have forward compatible implementations.
type UnimplementedOAuth2Server struct {
}

func (UnimplementedOAuth2Server) Token(context.Context, *TokenInput) (*TokenOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Token not implemented")
}
func (UnimplementedOAuth2Server) UserInfo(context.Context, *UserInfoInput) (*UserInfoOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserInfo not implemented")
}
func (UnimplementedOAuth2Server) Introspect(context.Context, *IntrospectInput) (*IntrospectOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Introspect not implemented")
}
func (UnimplementedOAuth2Server) Revoke(context.Context, *RevokeInput) (*RevokeOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Revoke not implemented")
}
func (UnimplementedOAuth2Server) AuthCode(context.Context, *AuthCodeInput) (*AuthCodeOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AuthCode not implemented")
}
func (UnimplementedOAuth2Server) Authorize(context.Context, *AuthorizeInput) (*AuthorizeOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Authorize not implemented")
}
func (UnimplementedOAuth2Server) Callback(context.Context, *CallbackInput) (*CallbackOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Callback not implemented")
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
		FullMethod: OAuth2_Token_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OAuth2Server).Token(ctx, req.(*TokenInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _OAuth2_UserInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserInfoInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OAuth2Server).UserInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OAuth2_UserInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OAuth2Server).UserInfo(ctx, req.(*UserInfoInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _OAuth2_Introspect_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IntrospectInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OAuth2Server).Introspect(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OAuth2_Introspect_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OAuth2Server).Introspect(ctx, req.(*IntrospectInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _OAuth2_Revoke_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RevokeInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OAuth2Server).Revoke(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OAuth2_Revoke_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OAuth2Server).Revoke(ctx, req.(*RevokeInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _OAuth2_AuthCode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthCodeInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OAuth2Server).AuthCode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OAuth2_AuthCode_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OAuth2Server).AuthCode(ctx, req.(*AuthCodeInput))
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
		FullMethod: OAuth2_Authorize_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OAuth2Server).Authorize(ctx, req.(*AuthorizeInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _OAuth2_Callback_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CallbackInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OAuth2Server).Callback(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OAuth2_Callback_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OAuth2Server).Callback(ctx, req.(*CallbackInput))
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
			MethodName: "UserInfo",
			Handler:    _OAuth2_UserInfo_Handler,
		},
		{
			MethodName: "Introspect",
			Handler:    _OAuth2_Introspect_Handler,
		},
		{
			MethodName: "Revoke",
			Handler:    _OAuth2_Revoke_Handler,
		},
		{
			MethodName: "AuthCode",
			Handler:    _OAuth2_AuthCode_Handler,
		},
		{
			MethodName: "Authorize",
			Handler:    _OAuth2_Authorize_Handler,
		},
		{
			MethodName: "Callback",
			Handler:    _OAuth2_Callback_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "eolymp/oauth2/oauth2.proto",
}
