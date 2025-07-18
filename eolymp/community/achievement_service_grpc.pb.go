// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.4
// source: eolymp/community/achievement_service.proto

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
	AchievementService_AssignAchievement_FullMethodName   = "/eolymp.community.AchievementService/AssignAchievement"
	AchievementService_UnassignAchievement_FullMethodName = "/eolymp.community.AchievementService/UnassignAchievement"
	AchievementService_ListAchievements_FullMethodName    = "/eolymp.community.AchievementService/ListAchievements"
)

// AchievementServiceClient is the client API for AchievementService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AchievementServiceClient interface {
	AssignAchievement(ctx context.Context, in *AssignAchievementInput, opts ...grpc.CallOption) (*AssignAchievementOutput, error)
	UnassignAchievement(ctx context.Context, in *UnassignAchievementInput, opts ...grpc.CallOption) (*UnassignAchievementOutput, error)
	ListAchievements(ctx context.Context, in *ListAchievementsInput, opts ...grpc.CallOption) (*ListAchievementsOutput, error)
}

type achievementServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAchievementServiceClient(cc grpc.ClientConnInterface) AchievementServiceClient {
	return &achievementServiceClient{cc}
}

func (c *achievementServiceClient) AssignAchievement(ctx context.Context, in *AssignAchievementInput, opts ...grpc.CallOption) (*AssignAchievementOutput, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AssignAchievementOutput)
	err := c.cc.Invoke(ctx, AchievementService_AssignAchievement_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *achievementServiceClient) UnassignAchievement(ctx context.Context, in *UnassignAchievementInput, opts ...grpc.CallOption) (*UnassignAchievementOutput, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UnassignAchievementOutput)
	err := c.cc.Invoke(ctx, AchievementService_UnassignAchievement_FullMethodName, in, out, cOpts...)
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
// for forward compatibility.
type AchievementServiceServer interface {
	AssignAchievement(context.Context, *AssignAchievementInput) (*AssignAchievementOutput, error)
	UnassignAchievement(context.Context, *UnassignAchievementInput) (*UnassignAchievementOutput, error)
	ListAchievements(context.Context, *ListAchievementsInput) (*ListAchievementsOutput, error)
}

// UnimplementedAchievementServiceServer should be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedAchievementServiceServer struct{}

func (UnimplementedAchievementServiceServer) AssignAchievement(context.Context, *AssignAchievementInput) (*AssignAchievementOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AssignAchievement not implemented")
}
func (UnimplementedAchievementServiceServer) UnassignAchievement(context.Context, *UnassignAchievementInput) (*UnassignAchievementOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnassignAchievement not implemented")
}
func (UnimplementedAchievementServiceServer) ListAchievements(context.Context, *ListAchievementsInput) (*ListAchievementsOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListAchievements not implemented")
}
func (UnimplementedAchievementServiceServer) testEmbeddedByValue() {}

// UnsafeAchievementServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AchievementServiceServer will
// result in compilation errors.
type UnsafeAchievementServiceServer interface {
	mustEmbedUnimplementedAchievementServiceServer()
}

func RegisterAchievementServiceServer(s grpc.ServiceRegistrar, srv AchievementServiceServer) {
	// If the following call pancis, it indicates UnimplementedAchievementServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&AchievementService_ServiceDesc, srv)
}

func _AchievementService_AssignAchievement_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AssignAchievementInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AchievementServiceServer).AssignAchievement(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AchievementService_AssignAchievement_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AchievementServiceServer).AssignAchievement(ctx, req.(*AssignAchievementInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _AchievementService_UnassignAchievement_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UnassignAchievementInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AchievementServiceServer).UnassignAchievement(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AchievementService_UnassignAchievement_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AchievementServiceServer).UnassignAchievement(ctx, req.(*UnassignAchievementInput))
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
	ServiceName: "eolymp.community.AchievementService",
	HandlerType: (*AchievementServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AssignAchievement",
			Handler:    _AchievementService_AssignAchievement_Handler,
		},
		{
			MethodName: "UnassignAchievement",
			Handler:    _AchievementService_UnassignAchievement_Handler,
		},
		{
			MethodName: "ListAchievements",
			Handler:    _AchievementService_ListAchievements_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "eolymp/community/achievement_service.proto",
}
