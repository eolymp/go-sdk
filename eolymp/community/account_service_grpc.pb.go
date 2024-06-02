// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v3.21.12
// source: eolymp/community/account_service.proto

package community

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
	AccountService_CreateAccount_FullMethodName                   = "/eolymp.community.AccountService/CreateAccount"
	AccountService_DescribeAccount_FullMethodName                 = "/eolymp.community.AccountService/DescribeAccount"
	AccountService_UpdateAccount_FullMethodName                   = "/eolymp.community.AccountService/UpdateAccount"
	AccountService_UploadPicture_FullMethodName                   = "/eolymp.community.AccountService/UploadPicture"
	AccountService_DeleteAccount_FullMethodName                   = "/eolymp.community.AccountService/DeleteAccount"
	AccountService_ResendVerification_FullMethodName              = "/eolymp.community.AccountService/ResendVerification"
	AccountService_CompleteVerification_FullMethodName            = "/eolymp.community.AccountService/CompleteVerification"
	AccountService_StartRecovery_FullMethodName                   = "/eolymp.community.AccountService/StartRecovery"
	AccountService_CompleteRecovery_FullMethodName                = "/eolymp.community.AccountService/CompleteRecovery"
	AccountService_DescribeNotificationPreferences_FullMethodName = "/eolymp.community.AccountService/DescribeNotificationPreferences"
	AccountService_UpdateNotificationPreferences_FullMethodName   = "/eolymp.community.AccountService/UpdateNotificationPreferences"
)

// AccountServiceClient is the client API for AccountService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// AccountService provides API to create and manage your own account.
type AccountServiceClient interface {
	CreateAccount(ctx context.Context, in *CreateAccountInput, opts ...grpc.CallOption) (*CreateAccountOutput, error)
	DescribeAccount(ctx context.Context, in *DescribeAccountInput, opts ...grpc.CallOption) (*DescribeAccountOutput, error)
	UpdateAccount(ctx context.Context, in *UpdateAccountInput, opts ...grpc.CallOption) (*UpdateAccountOutput, error)
	UploadPicture(ctx context.Context, in *UploadPictureInput, opts ...grpc.CallOption) (*UploadPictureOutput, error)
	DeleteAccount(ctx context.Context, in *DeleteAccountInput, opts ...grpc.CallOption) (*DeleteAccountOutput, error)
	ResendVerification(ctx context.Context, in *ResendVerificationInput, opts ...grpc.CallOption) (*ResendVerificationOutput, error)
	CompleteVerification(ctx context.Context, in *CompleteVerificationInput, opts ...grpc.CallOption) (*CompleteVerificationOutput, error)
	StartRecovery(ctx context.Context, in *StartRecoveryInput, opts ...grpc.CallOption) (*StartRecoveryOutput, error)
	CompleteRecovery(ctx context.Context, in *CompleteRecoverInput, opts ...grpc.CallOption) (*CompleteRecoverOutput, error)
	DescribeNotificationPreferences(ctx context.Context, in *DescribeNotificationPreferencesInput, opts ...grpc.CallOption) (*DescribeNotificationPreferencesOutput, error)
	UpdateNotificationPreferences(ctx context.Context, in *UpdateNotificationPreferencesInput, opts ...grpc.CallOption) (*UpdateNotificationPreferencesOutput, error)
}

type accountServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAccountServiceClient(cc grpc.ClientConnInterface) AccountServiceClient {
	return &accountServiceClient{cc}
}

func (c *accountServiceClient) CreateAccount(ctx context.Context, in *CreateAccountInput, opts ...grpc.CallOption) (*CreateAccountOutput, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateAccountOutput)
	err := c.cc.Invoke(ctx, AccountService_CreateAccount_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountServiceClient) DescribeAccount(ctx context.Context, in *DescribeAccountInput, opts ...grpc.CallOption) (*DescribeAccountOutput, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DescribeAccountOutput)
	err := c.cc.Invoke(ctx, AccountService_DescribeAccount_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountServiceClient) UpdateAccount(ctx context.Context, in *UpdateAccountInput, opts ...grpc.CallOption) (*UpdateAccountOutput, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateAccountOutput)
	err := c.cc.Invoke(ctx, AccountService_UpdateAccount_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountServiceClient) UploadPicture(ctx context.Context, in *UploadPictureInput, opts ...grpc.CallOption) (*UploadPictureOutput, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UploadPictureOutput)
	err := c.cc.Invoke(ctx, AccountService_UploadPicture_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountServiceClient) DeleteAccount(ctx context.Context, in *DeleteAccountInput, opts ...grpc.CallOption) (*DeleteAccountOutput, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteAccountOutput)
	err := c.cc.Invoke(ctx, AccountService_DeleteAccount_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountServiceClient) ResendVerification(ctx context.Context, in *ResendVerificationInput, opts ...grpc.CallOption) (*ResendVerificationOutput, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ResendVerificationOutput)
	err := c.cc.Invoke(ctx, AccountService_ResendVerification_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountServiceClient) CompleteVerification(ctx context.Context, in *CompleteVerificationInput, opts ...grpc.CallOption) (*CompleteVerificationOutput, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CompleteVerificationOutput)
	err := c.cc.Invoke(ctx, AccountService_CompleteVerification_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountServiceClient) StartRecovery(ctx context.Context, in *StartRecoveryInput, opts ...grpc.CallOption) (*StartRecoveryOutput, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(StartRecoveryOutput)
	err := c.cc.Invoke(ctx, AccountService_StartRecovery_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountServiceClient) CompleteRecovery(ctx context.Context, in *CompleteRecoverInput, opts ...grpc.CallOption) (*CompleteRecoverOutput, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CompleteRecoverOutput)
	err := c.cc.Invoke(ctx, AccountService_CompleteRecovery_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountServiceClient) DescribeNotificationPreferences(ctx context.Context, in *DescribeNotificationPreferencesInput, opts ...grpc.CallOption) (*DescribeNotificationPreferencesOutput, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DescribeNotificationPreferencesOutput)
	err := c.cc.Invoke(ctx, AccountService_DescribeNotificationPreferences_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountServiceClient) UpdateNotificationPreferences(ctx context.Context, in *UpdateNotificationPreferencesInput, opts ...grpc.CallOption) (*UpdateNotificationPreferencesOutput, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateNotificationPreferencesOutput)
	err := c.cc.Invoke(ctx, AccountService_UpdateNotificationPreferences_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AccountServiceServer is the server API for AccountService service.
// All implementations should embed UnimplementedAccountServiceServer
// for forward compatibility
//
// AccountService provides API to create and manage your own account.
type AccountServiceServer interface {
	CreateAccount(context.Context, *CreateAccountInput) (*CreateAccountOutput, error)
	DescribeAccount(context.Context, *DescribeAccountInput) (*DescribeAccountOutput, error)
	UpdateAccount(context.Context, *UpdateAccountInput) (*UpdateAccountOutput, error)
	UploadPicture(context.Context, *UploadPictureInput) (*UploadPictureOutput, error)
	DeleteAccount(context.Context, *DeleteAccountInput) (*DeleteAccountOutput, error)
	ResendVerification(context.Context, *ResendVerificationInput) (*ResendVerificationOutput, error)
	CompleteVerification(context.Context, *CompleteVerificationInput) (*CompleteVerificationOutput, error)
	StartRecovery(context.Context, *StartRecoveryInput) (*StartRecoveryOutput, error)
	CompleteRecovery(context.Context, *CompleteRecoverInput) (*CompleteRecoverOutput, error)
	DescribeNotificationPreferences(context.Context, *DescribeNotificationPreferencesInput) (*DescribeNotificationPreferencesOutput, error)
	UpdateNotificationPreferences(context.Context, *UpdateNotificationPreferencesInput) (*UpdateNotificationPreferencesOutput, error)
}

// UnimplementedAccountServiceServer should be embedded to have forward compatible implementations.
type UnimplementedAccountServiceServer struct {
}

func (UnimplementedAccountServiceServer) CreateAccount(context.Context, *CreateAccountInput) (*CreateAccountOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAccount not implemented")
}
func (UnimplementedAccountServiceServer) DescribeAccount(context.Context, *DescribeAccountInput) (*DescribeAccountOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DescribeAccount not implemented")
}
func (UnimplementedAccountServiceServer) UpdateAccount(context.Context, *UpdateAccountInput) (*UpdateAccountOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAccount not implemented")
}
func (UnimplementedAccountServiceServer) UploadPicture(context.Context, *UploadPictureInput) (*UploadPictureOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UploadPicture not implemented")
}
func (UnimplementedAccountServiceServer) DeleteAccount(context.Context, *DeleteAccountInput) (*DeleteAccountOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteAccount not implemented")
}
func (UnimplementedAccountServiceServer) ResendVerification(context.Context, *ResendVerificationInput) (*ResendVerificationOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ResendVerification not implemented")
}
func (UnimplementedAccountServiceServer) CompleteVerification(context.Context, *CompleteVerificationInput) (*CompleteVerificationOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CompleteVerification not implemented")
}
func (UnimplementedAccountServiceServer) StartRecovery(context.Context, *StartRecoveryInput) (*StartRecoveryOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StartRecovery not implemented")
}
func (UnimplementedAccountServiceServer) CompleteRecovery(context.Context, *CompleteRecoverInput) (*CompleteRecoverOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CompleteRecovery not implemented")
}
func (UnimplementedAccountServiceServer) DescribeNotificationPreferences(context.Context, *DescribeNotificationPreferencesInput) (*DescribeNotificationPreferencesOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DescribeNotificationPreferences not implemented")
}
func (UnimplementedAccountServiceServer) UpdateNotificationPreferences(context.Context, *UpdateNotificationPreferencesInput) (*UpdateNotificationPreferencesOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateNotificationPreferences not implemented")
}

// UnsafeAccountServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AccountServiceServer will
// result in compilation errors.
type UnsafeAccountServiceServer interface {
	mustEmbedUnimplementedAccountServiceServer()
}

func RegisterAccountServiceServer(s grpc.ServiceRegistrar, srv AccountServiceServer) {
	s.RegisterService(&AccountService_ServiceDesc, srv)
}

func _AccountService_CreateAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAccountInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServiceServer).CreateAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AccountService_CreateAccount_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServiceServer).CreateAccount(ctx, req.(*CreateAccountInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountService_DescribeAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DescribeAccountInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServiceServer).DescribeAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AccountService_DescribeAccount_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServiceServer).DescribeAccount(ctx, req.(*DescribeAccountInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountService_UpdateAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateAccountInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServiceServer).UpdateAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AccountService_UpdateAccount_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServiceServer).UpdateAccount(ctx, req.(*UpdateAccountInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountService_UploadPicture_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UploadPictureInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServiceServer).UploadPicture(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AccountService_UploadPicture_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServiceServer).UploadPicture(ctx, req.(*UploadPictureInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountService_DeleteAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteAccountInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServiceServer).DeleteAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AccountService_DeleteAccount_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServiceServer).DeleteAccount(ctx, req.(*DeleteAccountInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountService_ResendVerification_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResendVerificationInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServiceServer).ResendVerification(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AccountService_ResendVerification_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServiceServer).ResendVerification(ctx, req.(*ResendVerificationInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountService_CompleteVerification_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CompleteVerificationInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServiceServer).CompleteVerification(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AccountService_CompleteVerification_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServiceServer).CompleteVerification(ctx, req.(*CompleteVerificationInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountService_StartRecovery_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartRecoveryInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServiceServer).StartRecovery(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AccountService_StartRecovery_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServiceServer).StartRecovery(ctx, req.(*StartRecoveryInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountService_CompleteRecovery_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CompleteRecoverInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServiceServer).CompleteRecovery(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AccountService_CompleteRecovery_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServiceServer).CompleteRecovery(ctx, req.(*CompleteRecoverInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountService_DescribeNotificationPreferences_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DescribeNotificationPreferencesInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServiceServer).DescribeNotificationPreferences(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AccountService_DescribeNotificationPreferences_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServiceServer).DescribeNotificationPreferences(ctx, req.(*DescribeNotificationPreferencesInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountService_UpdateNotificationPreferences_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateNotificationPreferencesInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServiceServer).UpdateNotificationPreferences(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AccountService_UpdateNotificationPreferences_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServiceServer).UpdateNotificationPreferences(ctx, req.(*UpdateNotificationPreferencesInput))
	}
	return interceptor(ctx, in, info, handler)
}

// AccountService_ServiceDesc is the grpc.ServiceDesc for AccountService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AccountService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "eolymp.community.AccountService",
	HandlerType: (*AccountServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateAccount",
			Handler:    _AccountService_CreateAccount_Handler,
		},
		{
			MethodName: "DescribeAccount",
			Handler:    _AccountService_DescribeAccount_Handler,
		},
		{
			MethodName: "UpdateAccount",
			Handler:    _AccountService_UpdateAccount_Handler,
		},
		{
			MethodName: "UploadPicture",
			Handler:    _AccountService_UploadPicture_Handler,
		},
		{
			MethodName: "DeleteAccount",
			Handler:    _AccountService_DeleteAccount_Handler,
		},
		{
			MethodName: "ResendVerification",
			Handler:    _AccountService_ResendVerification_Handler,
		},
		{
			MethodName: "CompleteVerification",
			Handler:    _AccountService_CompleteVerification_Handler,
		},
		{
			MethodName: "StartRecovery",
			Handler:    _AccountService_StartRecovery_Handler,
		},
		{
			MethodName: "CompleteRecovery",
			Handler:    _AccountService_CompleteRecovery_Handler,
		},
		{
			MethodName: "DescribeNotificationPreferences",
			Handler:    _AccountService_DescribeNotificationPreferences_Handler,
		},
		{
			MethodName: "UpdateNotificationPreferences",
			Handler:    _AccountService_UpdateNotificationPreferences_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "eolymp/community/account_service.proto",
}
