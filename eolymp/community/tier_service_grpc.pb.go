// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.18.1
// source: eolymp/community/tier_service.proto

package community

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
	TierService_DescribeTier_FullMethodName   = "/eolymp.community.TierService/DescribeTier"
	TierService_ListTiers_FullMethodName      = "/eolymp.community.TierService/ListTiers"
	TierService_ListCurrencies_FullMethodName = "/eolymp.community.TierService/ListCurrencies"
)

// TierServiceClient is the client API for TierService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TierServiceClient interface {
	DescribeTier(ctx context.Context, in *DescribeTierInput, opts ...grpc.CallOption) (*DescribeTierOutput, error)
	ListTiers(ctx context.Context, in *ListTiersInput, opts ...grpc.CallOption) (*ListTiersOutput, error)
	ListCurrencies(ctx context.Context, in *ListCurrenciesInput, opts ...grpc.CallOption) (*ListCurrenciesOutput, error)
}

type tierServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTierServiceClient(cc grpc.ClientConnInterface) TierServiceClient {
	return &tierServiceClient{cc}
}

func (c *tierServiceClient) DescribeTier(ctx context.Context, in *DescribeTierInput, opts ...grpc.CallOption) (*DescribeTierOutput, error) {
	out := new(DescribeTierOutput)
	err := c.cc.Invoke(ctx, TierService_DescribeTier_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tierServiceClient) ListTiers(ctx context.Context, in *ListTiersInput, opts ...grpc.CallOption) (*ListTiersOutput, error) {
	out := new(ListTiersOutput)
	err := c.cc.Invoke(ctx, TierService_ListTiers_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tierServiceClient) ListCurrencies(ctx context.Context, in *ListCurrenciesInput, opts ...grpc.CallOption) (*ListCurrenciesOutput, error) {
	out := new(ListCurrenciesOutput)
	err := c.cc.Invoke(ctx, TierService_ListCurrencies_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TierServiceServer is the server API for TierService service.
// All implementations should embed UnimplementedTierServiceServer
// for forward compatibility
type TierServiceServer interface {
	DescribeTier(context.Context, *DescribeTierInput) (*DescribeTierOutput, error)
	ListTiers(context.Context, *ListTiersInput) (*ListTiersOutput, error)
	ListCurrencies(context.Context, *ListCurrenciesInput) (*ListCurrenciesOutput, error)
}

// UnimplementedTierServiceServer should be embedded to have forward compatible implementations.
type UnimplementedTierServiceServer struct {
}

func (UnimplementedTierServiceServer) DescribeTier(context.Context, *DescribeTierInput) (*DescribeTierOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DescribeTier not implemented")
}
func (UnimplementedTierServiceServer) ListTiers(context.Context, *ListTiersInput) (*ListTiersOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListTiers not implemented")
}
func (UnimplementedTierServiceServer) ListCurrencies(context.Context, *ListCurrenciesInput) (*ListCurrenciesOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListCurrencies not implemented")
}

// UnsafeTierServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TierServiceServer will
// result in compilation errors.
type UnsafeTierServiceServer interface {
	mustEmbedUnimplementedTierServiceServer()
}

func RegisterTierServiceServer(s grpc.ServiceRegistrar, srv TierServiceServer) {
	s.RegisterService(&TierService_ServiceDesc, srv)
}

func _TierService_DescribeTier_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DescribeTierInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TierServiceServer).DescribeTier(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TierService_DescribeTier_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TierServiceServer).DescribeTier(ctx, req.(*DescribeTierInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _TierService_ListTiers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListTiersInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TierServiceServer).ListTiers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TierService_ListTiers_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TierServiceServer).ListTiers(ctx, req.(*ListTiersInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _TierService_ListCurrencies_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListCurrenciesInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TierServiceServer).ListCurrencies(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TierService_ListCurrencies_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TierServiceServer).ListCurrencies(ctx, req.(*ListCurrenciesInput))
	}
	return interceptor(ctx, in, info, handler)
}

// TierService_ServiceDesc is the grpc.ServiceDesc for TierService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TierService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "eolymp.community.TierService",
	HandlerType: (*TierServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DescribeTier",
			Handler:    _TierService_DescribeTier_Handler,
		},
		{
			MethodName: "ListTiers",
			Handler:    _TierService_ListTiers_Handler,
		},
		{
			MethodName: "ListCurrencies",
			Handler:    _TierService_ListCurrencies_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "eolymp/community/tier_service.proto",
}
