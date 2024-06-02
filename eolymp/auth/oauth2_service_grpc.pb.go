// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v3.21.12
// source: eolymp/auth/oauth2_service.proto

package auth

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
	OAuth2Service_IssueToken_FullMethodName      = "/eolymp.auth.OAuth2Service/IssueToken"
	OAuth2Service_IntrospectToken_FullMethodName = "/eolymp.auth.OAuth2Service/IntrospectToken"
	OAuth2Service_RevokeToken_FullMethodName     = "/eolymp.auth.OAuth2Service/RevokeToken"
	OAuth2Service_RequestAuth_FullMethodName     = "/eolymp.auth.OAuth2Service/RequestAuth"
	OAuth2Service_UserInfo_FullMethodName        = "/eolymp.auth.OAuth2Service/UserInfo"
)

// OAuth2ServiceClient is the client API for OAuth2Service service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OAuth2ServiceClient interface {
	// issue a new token
	IssueToken(ctx context.Context, in *IssueTokenInput, opts ...grpc.CallOption) (*IssueTokenOutput, error)
	// introspect token
	IntrospectToken(ctx context.Context, in *IntrospectTokenInput, opts ...grpc.CallOption) (*IntrospectTokenOutput, error)
	// revokes a given token
	RevokeToken(ctx context.Context, in *RevokeTokenInput, opts ...grpc.CallOption) (*RevokeTokenOutput, error)
	// issues a new authorization code for the authenticated user
	RequestAuth(ctx context.Context, in *RequestAuthInput, opts ...grpc.CallOption) (*RequestAuthOutput, error)
	// returns claims for currently authenticated user
	UserInfo(ctx context.Context, in *UserInfoInput, opts ...grpc.CallOption) (*UserInfoOutput, error)
}

type oAuth2ServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewOAuth2ServiceClient(cc grpc.ClientConnInterface) OAuth2ServiceClient {
	return &oAuth2ServiceClient{cc}
}

func (c *oAuth2ServiceClient) IssueToken(ctx context.Context, in *IssueTokenInput, opts ...grpc.CallOption) (*IssueTokenOutput, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(IssueTokenOutput)
	err := c.cc.Invoke(ctx, OAuth2Service_IssueToken_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *oAuth2ServiceClient) IntrospectToken(ctx context.Context, in *IntrospectTokenInput, opts ...grpc.CallOption) (*IntrospectTokenOutput, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(IntrospectTokenOutput)
	err := c.cc.Invoke(ctx, OAuth2Service_IntrospectToken_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *oAuth2ServiceClient) RevokeToken(ctx context.Context, in *RevokeTokenInput, opts ...grpc.CallOption) (*RevokeTokenOutput, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RevokeTokenOutput)
	err := c.cc.Invoke(ctx, OAuth2Service_RevokeToken_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *oAuth2ServiceClient) RequestAuth(ctx context.Context, in *RequestAuthInput, opts ...grpc.CallOption) (*RequestAuthOutput, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RequestAuthOutput)
	err := c.cc.Invoke(ctx, OAuth2Service_RequestAuth_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *oAuth2ServiceClient) UserInfo(ctx context.Context, in *UserInfoInput, opts ...grpc.CallOption) (*UserInfoOutput, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UserInfoOutput)
	err := c.cc.Invoke(ctx, OAuth2Service_UserInfo_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OAuth2ServiceServer is the server API for OAuth2Service service.
// All implementations should embed UnimplementedOAuth2ServiceServer
// for forward compatibility
type OAuth2ServiceServer interface {
	// issue a new token
	IssueToken(context.Context, *IssueTokenInput) (*IssueTokenOutput, error)
	// introspect token
	IntrospectToken(context.Context, *IntrospectTokenInput) (*IntrospectTokenOutput, error)
	// revokes a given token
	RevokeToken(context.Context, *RevokeTokenInput) (*RevokeTokenOutput, error)
	// issues a new authorization code for the authenticated user
	RequestAuth(context.Context, *RequestAuthInput) (*RequestAuthOutput, error)
	// returns claims for currently authenticated user
	UserInfo(context.Context, *UserInfoInput) (*UserInfoOutput, error)
}

// UnimplementedOAuth2ServiceServer should be embedded to have forward compatible implementations.
type UnimplementedOAuth2ServiceServer struct {
}

func (UnimplementedOAuth2ServiceServer) IssueToken(context.Context, *IssueTokenInput) (*IssueTokenOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IssueToken not implemented")
}
func (UnimplementedOAuth2ServiceServer) IntrospectToken(context.Context, *IntrospectTokenInput) (*IntrospectTokenOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IntrospectToken not implemented")
}
func (UnimplementedOAuth2ServiceServer) RevokeToken(context.Context, *RevokeTokenInput) (*RevokeTokenOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RevokeToken not implemented")
}
func (UnimplementedOAuth2ServiceServer) RequestAuth(context.Context, *RequestAuthInput) (*RequestAuthOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RequestAuth not implemented")
}
func (UnimplementedOAuth2ServiceServer) UserInfo(context.Context, *UserInfoInput) (*UserInfoOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserInfo not implemented")
}

// UnsafeOAuth2ServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OAuth2ServiceServer will
// result in compilation errors.
type UnsafeOAuth2ServiceServer interface {
	mustEmbedUnimplementedOAuth2ServiceServer()
}

func RegisterOAuth2ServiceServer(s grpc.ServiceRegistrar, srv OAuth2ServiceServer) {
	s.RegisterService(&OAuth2Service_ServiceDesc, srv)
}

func _OAuth2Service_IssueToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IssueTokenInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OAuth2ServiceServer).IssueToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OAuth2Service_IssueToken_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OAuth2ServiceServer).IssueToken(ctx, req.(*IssueTokenInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _OAuth2Service_IntrospectToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IntrospectTokenInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OAuth2ServiceServer).IntrospectToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OAuth2Service_IntrospectToken_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OAuth2ServiceServer).IntrospectToken(ctx, req.(*IntrospectTokenInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _OAuth2Service_RevokeToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RevokeTokenInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OAuth2ServiceServer).RevokeToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OAuth2Service_RevokeToken_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OAuth2ServiceServer).RevokeToken(ctx, req.(*RevokeTokenInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _OAuth2Service_RequestAuth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestAuthInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OAuth2ServiceServer).RequestAuth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OAuth2Service_RequestAuth_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OAuth2ServiceServer).RequestAuth(ctx, req.(*RequestAuthInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _OAuth2Service_UserInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserInfoInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OAuth2ServiceServer).UserInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OAuth2Service_UserInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OAuth2ServiceServer).UserInfo(ctx, req.(*UserInfoInput))
	}
	return interceptor(ctx, in, info, handler)
}

// OAuth2Service_ServiceDesc is the grpc.ServiceDesc for OAuth2Service service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OAuth2Service_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "eolymp.auth.OAuth2Service",
	HandlerType: (*OAuth2ServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "IssueToken",
			Handler:    _OAuth2Service_IssueToken_Handler,
		},
		{
			MethodName: "IntrospectToken",
			Handler:    _OAuth2Service_IntrospectToken_Handler,
		},
		{
			MethodName: "RevokeToken",
			Handler:    _OAuth2Service_RevokeToken_Handler,
		},
		{
			MethodName: "RequestAuth",
			Handler:    _OAuth2Service_RequestAuth_Handler,
		},
		{
			MethodName: "UserInfo",
			Handler:    _OAuth2Service_UserInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "eolymp/auth/oauth2_service.proto",
}
