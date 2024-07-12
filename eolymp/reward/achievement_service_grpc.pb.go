// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v3.21.12
// source: eolymp/reward/achievement_service.proto

package reward

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
	AchievementService_CreateAchievement_FullMethodName   = "/eolymp.reward.AchievementService/CreateAchievement"
	AchievementService_UpdateAchievement_FullMethodName   = "/eolymp.reward.AchievementService/UpdateAchievement"
	AchievementService_RemoveAchievement_FullMethodName   = "/eolymp.reward.AchievementService/RemoveAchievement"
	AchievementService_DescribeAchievement_FullMethodName = "/eolymp.reward.AchievementService/DescribeAchievement"
	AchievementService_ListAchievements_FullMethodName    = "/eolymp.reward.AchievementService/ListAchievements"
)

// AchievementServiceClient is the client API for AchievementService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AchievementServiceClient interface {
	CreateAchievement(ctx context.Context, in *CreateAchievementInput, opts ...grpc.CallOption) (*CreateAchievementOutput, error)
	UpdateAchievement(ctx context.Context, in *UpdateAchievementInput, opts ...grpc.CallOption) (*UpdateAchievementOutput, error)
	RemoveAchievement(ctx context.Context, in *RemoveAchievementInput, opts ...grpc.CallOption) (*RemoveAchievementOutput, error)
	DescribeAchievement(ctx context.Context, in *DescribeAchievementInput, opts ...grpc.CallOption) (*DescribeAchievementOutput, error)
	ListAchievements(ctx context.Context, in *ListAchievementsInput, opts ...grpc.CallOption) (*ListAchievementsOutput, error)
}

type achievementServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAchievementServiceClient(cc grpc.ClientConnInterface) AchievementServiceClient {
	return &achievementServiceClient{cc}
}

func (c *achievementServiceClient) CreateAchievement(ctx context.Context, in *CreateAchievementInput, opts ...grpc.CallOption) (*CreateAchievementOutput, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateAchievementOutput)
	err := c.cc.Invoke(ctx, AchievementService_CreateAchievement_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *achievementServiceClient) UpdateAchievement(ctx context.Context, in *UpdateAchievementInput, opts ...grpc.CallOption) (*UpdateAchievementOutput, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateAchievementOutput)
	err := c.cc.Invoke(ctx, AchievementService_UpdateAchievement_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *achievementServiceClient) RemoveAchievement(ctx context.Context, in *RemoveAchievementInput, opts ...grpc.CallOption) (*RemoveAchievementOutput, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RemoveAchievementOutput)
	err := c.cc.Invoke(ctx, AchievementService_RemoveAchievement_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *achievementServiceClient) DescribeAchievement(ctx context.Context, in *DescribeAchievementInput, opts ...grpc.CallOption) (*DescribeAchievementOutput, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DescribeAchievementOutput)
	err := c.cc.Invoke(ctx, AchievementService_DescribeAchievement_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *achievementServiceClient) ListAchievements(ctx context.Context, in *ListAchievementsInput, opts ...grpc.CallOption) (*ListAchievementsOutput, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListAchievementsOutput)
	err := c.cc.Invoke(ctx, AchievementService_ListAchievements_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AchievementServiceServer is the server API for AchievementService service.
// All implementations should embed UnimplementedAchievementServiceServer
// for forward compatibility
type AchievementServiceServer interface {
	CreateAchievement(context.Context, *CreateAchievementInput) (*CreateAchievementOutput, error)
	UpdateAchievement(context.Context, *UpdateAchievementInput) (*UpdateAchievementOutput, error)
	RemoveAchievement(context.Context, *RemoveAchievementInput) (*RemoveAchievementOutput, error)
	DescribeAchievement(context.Context, *DescribeAchievementInput) (*DescribeAchievementOutput, error)
	ListAchievements(context.Context, *ListAchievementsInput) (*ListAchievementsOutput, error)
}

// UnimplementedAchievementServiceServer should be embedded to have forward compatible implementations.
type UnimplementedAchievementServiceServer struct {
}

func (UnimplementedAchievementServiceServer) CreateAchievement(context.Context, *CreateAchievementInput) (*CreateAchievementOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAchievement not implemented")
}
func (UnimplementedAchievementServiceServer) UpdateAchievement(context.Context, *UpdateAchievementInput) (*UpdateAchievementOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAchievement not implemented")
}
func (UnimplementedAchievementServiceServer) RemoveAchievement(context.Context, *RemoveAchievementInput) (*RemoveAchievementOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveAchievement not implemented")
}
func (UnimplementedAchievementServiceServer) DescribeAchievement(context.Context, *DescribeAchievementInput) (*DescribeAchievementOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DescribeAchievement not implemented")
}
func (UnimplementedAchievementServiceServer) ListAchievements(context.Context, *ListAchievementsInput) (*ListAchievementsOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListAchievements not implemented")
}

// UnsafeAchievementServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AchievementServiceServer will
// result in compilation errors.
type UnsafeAchievementServiceServer interface {
	mustEmbedUnimplementedAchievementServiceServer()
}

func RegisterAchievementServiceServer(s grpc.ServiceRegistrar, srv AchievementServiceServer) {
	s.RegisterService(&AchievementService_ServiceDesc, srv)
}

func _AchievementService_CreateAchievement_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAchievementInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AchievementServiceServer).CreateAchievement(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AchievementService_CreateAchievement_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AchievementServiceServer).CreateAchievement(ctx, req.(*CreateAchievementInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _AchievementService_UpdateAchievement_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateAchievementInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AchievementServiceServer).UpdateAchievement(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AchievementService_UpdateAchievement_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AchievementServiceServer).UpdateAchievement(ctx, req.(*UpdateAchievementInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _AchievementService_RemoveAchievement_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveAchievementInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AchievementServiceServer).RemoveAchievement(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AchievementService_RemoveAchievement_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AchievementServiceServer).RemoveAchievement(ctx, req.(*RemoveAchievementInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _AchievementService_DescribeAchievement_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DescribeAchievementInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AchievementServiceServer).DescribeAchievement(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AchievementService_DescribeAchievement_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AchievementServiceServer).DescribeAchievement(ctx, req.(*DescribeAchievementInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _AchievementService_ListAchievements_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListAchievementsInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AchievementServiceServer).ListAchievements(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AchievementService_ListAchievements_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AchievementServiceServer).ListAchievements(ctx, req.(*ListAchievementsInput))
	}
	return interceptor(ctx, in, info, handler)
}

// AchievementService_ServiceDesc is the grpc.ServiceDesc for AchievementService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AchievementService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "eolymp.reward.AchievementService",
	HandlerType: (*AchievementServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateAchievement",
			Handler:    _AchievementService_CreateAchievement_Handler,
		},
		{
			MethodName: "UpdateAchievement",
			Handler:    _AchievementService_UpdateAchievement_Handler,
		},
		{
			MethodName: "RemoveAchievement",
			Handler:    _AchievementService_RemoveAchievement_Handler,
		},
		{
			MethodName: "DescribeAchievement",
			Handler:    _AchievementService_DescribeAchievement_Handler,
		},
		{
			MethodName: "ListAchievements",
			Handler:    _AchievementService_ListAchievements_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "eolymp/reward/achievement_service.proto",
}
