// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.2
// source: eolymp/wellknown/geography.proto

package playground

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

// GeographyClient is the client API for Geography service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GeographyClient interface {
	DescribeCountry(ctx context.Context, in *DescribeCountryInput, opts ...grpc.CallOption) (*DescribeCountryOutput, error)
	ListCountries(ctx context.Context, in *ListCountriesInput, opts ...grpc.CallOption) (*ListCountriesOutput, error)
}

type geographyClient struct {
	cc grpc.ClientConnInterface
}

func NewGeographyClient(cc grpc.ClientConnInterface) GeographyClient {
	return &geographyClient{cc}
}

func (c *geographyClient) DescribeCountry(ctx context.Context, in *DescribeCountryInput, opts ...grpc.CallOption) (*DescribeCountryOutput, error) {
	out := new(DescribeCountryOutput)
	err := c.cc.Invoke(ctx, "/eolymp.playground.Geography/DescribeCountry", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *geographyClient) ListCountries(ctx context.Context, in *ListCountriesInput, opts ...grpc.CallOption) (*ListCountriesOutput, error) {
	out := new(ListCountriesOutput)
	err := c.cc.Invoke(ctx, "/eolymp.playground.Geography/ListCountries", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GeographyServer is the server API for Geography service.
// All implementations must embed UnimplementedGeographyServer
// for forward compatibility
type GeographyServer interface {
	DescribeCountry(context.Context, *DescribeCountryInput) (*DescribeCountryOutput, error)
	ListCountries(context.Context, *ListCountriesInput) (*ListCountriesOutput, error)
	mustEmbedUnimplementedGeographyServer()
}

// UnimplementedGeographyServer must be embedded to have forward compatible implementations.
type UnimplementedGeographyServer struct {
}

func (UnimplementedGeographyServer) DescribeCountry(context.Context, *DescribeCountryInput) (*DescribeCountryOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DescribeCountry not implemented")
}
func (UnimplementedGeographyServer) ListCountries(context.Context, *ListCountriesInput) (*ListCountriesOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListCountries not implemented")
}
func (UnimplementedGeographyServer) mustEmbedUnimplementedGeographyServer() {}

// UnsafeGeographyServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GeographyServer will
// result in compilation errors.
type UnsafeGeographyServer interface {
	mustEmbedUnimplementedGeographyServer()
}

func RegisterGeographyServer(s grpc.ServiceRegistrar, srv GeographyServer) {
	s.RegisterService(&Geography_ServiceDesc, srv)
}

func _Geography_DescribeCountry_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DescribeCountryInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GeographyServer).DescribeCountry(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/eolymp.playground.Geography/DescribeCountry",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GeographyServer).DescribeCountry(ctx, req.(*DescribeCountryInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _Geography_ListCountries_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListCountriesInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GeographyServer).ListCountries(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/eolymp.playground.Geography/ListCountries",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GeographyServer).ListCountries(ctx, req.(*ListCountriesInput))
	}
	return interceptor(ctx, in, info, handler)
}

// Geography_ServiceDesc is the grpc.ServiceDesc for Geography service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Geography_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "eolymp.playground.Geography",
	HandlerType: (*GeographyServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DescribeCountry",
			Handler:    _Geography_DescribeCountry_Handler,
		},
		{
			MethodName: "ListCountries",
			Handler:    _Geography_ListCountries_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "eolymp/wellknown/geography.proto",
}
