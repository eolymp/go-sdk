// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v3.21.12
// source: eolymp/community/membership_service.proto

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
	MembershipService_DescribeMembership_FullMethodName = "/eolymp.community.MembershipService/DescribeMembership"
	MembershipService_UpdateMembership_FullMethodName   = "/eolymp.community.MembershipService/UpdateMembership"
)

// MembershipServiceClient is the client API for MembershipService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// MembershipService provides API to manage own membership in space.
type MembershipServiceClient interface {
	// Describe member profile for authenticated user
	DescribeMembership(ctx context.Context, in *DescribeMembershipInput, opts ...grpc.CallOption) (*DescribeMembershipOutput, error)
	// Update member profile for authenticated user
	UpdateMembership(ctx context.Context, in *UpdateMembershipInput, opts ...grpc.CallOption) (*UpdateMembershipOutput, error)
}

type membershipServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMembershipServiceClient(cc grpc.ClientConnInterface) MembershipServiceClient {
	return &membershipServiceClient{cc}
}

func (c *membershipServiceClient) DescribeMembership(ctx context.Context, in *DescribeMembershipInput, opts ...grpc.CallOption) (*DescribeMembershipOutput, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DescribeMembershipOutput)
	err := c.cc.Invoke(ctx, MembershipService_DescribeMembership_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *membershipServiceClient) UpdateMembership(ctx context.Context, in *UpdateMembershipInput, opts ...grpc.CallOption) (*UpdateMembershipOutput, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateMembershipOutput)
	err := c.cc.Invoke(ctx, MembershipService_UpdateMembership_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MembershipServiceServer is the server API for MembershipService service.
// All implementations should embed UnimplementedMembershipServiceServer
// for forward compatibility
//
// MembershipService provides API to manage own membership in space.
type MembershipServiceServer interface {
	// Describe member profile for authenticated user
	DescribeMembership(context.Context, *DescribeMembershipInput) (*DescribeMembershipOutput, error)
	// Update member profile for authenticated user
	UpdateMembership(context.Context, *UpdateMembershipInput) (*UpdateMembershipOutput, error)
}

// UnimplementedMembershipServiceServer should be embedded to have forward compatible implementations.
type UnimplementedMembershipServiceServer struct {
}

func (UnimplementedMembershipServiceServer) DescribeMembership(context.Context, *DescribeMembershipInput) (*DescribeMembershipOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DescribeMembership not implemented")
}
func (UnimplementedMembershipServiceServer) UpdateMembership(context.Context, *UpdateMembershipInput) (*UpdateMembershipOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateMembership not implemented")
}

// UnsafeMembershipServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MembershipServiceServer will
// result in compilation errors.
type UnsafeMembershipServiceServer interface {
	mustEmbedUnimplementedMembershipServiceServer()
}

func RegisterMembershipServiceServer(s grpc.ServiceRegistrar, srv MembershipServiceServer) {
	s.RegisterService(&MembershipService_ServiceDesc, srv)
}

func _MembershipService_DescribeMembership_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DescribeMembershipInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MembershipServiceServer).DescribeMembership(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MembershipService_DescribeMembership_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MembershipServiceServer).DescribeMembership(ctx, req.(*DescribeMembershipInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _MembershipService_UpdateMembership_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateMembershipInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MembershipServiceServer).UpdateMembership(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MembershipService_UpdateMembership_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MembershipServiceServer).UpdateMembership(ctx, req.(*UpdateMembershipInput))
	}
	return interceptor(ctx, in, info, handler)
}

// MembershipService_ServiceDesc is the grpc.ServiceDesc for MembershipService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MembershipService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "eolymp.community.MembershipService",
	HandlerType: (*MembershipServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DescribeMembership",
			Handler:    _MembershipService_DescribeMembership_Handler,
		},
		{
			MethodName: "UpdateMembership",
			Handler:    _MembershipService_UpdateMembership_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "eolymp/community/membership_service.proto",
}
