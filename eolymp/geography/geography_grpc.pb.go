// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.21.12
// source: eolymp/geography/geography.proto

package geography

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
	Geography_DescribeCountry_FullMethodName       = "/eolymp.geography.Geography/DescribeCountry"
	Geography_ListCountries_FullMethodName         = "/eolymp.geography.Geography/ListCountries"
	Geography_DescribeRegion_FullMethodName        = "/eolymp.geography.Geography/DescribeRegion"
	Geography_ListRegions_FullMethodName           = "/eolymp.geography.Geography/ListRegions"
	Geography_DeprecatedListRegions_FullMethodName = "/eolymp.geography.Geography/DeprecatedListRegions"
)

// GeographyClient is the client API for Geography service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GeographyClient interface {
	DescribeCountry(ctx context.Context, in *DescribeCountryInput, opts ...grpc.CallOption) (*DescribeCountryOutput, error)
	ListCountries(ctx context.Context, in *ListCountriesInput, opts ...grpc.CallOption) (*ListCountriesOutput, error)
	DescribeRegion(ctx context.Context, in *DescribeRegionInput, opts ...grpc.CallOption) (*DescribeRegionOutput, error)
	ListRegions(ctx context.Context, in *ListRegionsInput, opts ...grpc.CallOption) (*ListRegionsOutput, error)
	DeprecatedListRegions(ctx context.Context, in *ListRegionsInput, opts ...grpc.CallOption) (*ListRegionsOutput, error)
}

type geographyClient struct {
	cc grpc.ClientConnInterface
}

func NewGeographyClient(cc grpc.ClientConnInterface) GeographyClient {
	return &geographyClient{cc}
}

func (c *geographyClient) DescribeCountry(ctx context.Context, in *DescribeCountryInput, opts ...grpc.CallOption) (*DescribeCountryOutput, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DescribeCountryOutput)
	err := c.cc.Invoke(ctx, Geography_DescribeCountry_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *geographyClient) ListCountries(ctx context.Context, in *ListCountriesInput, opts ...grpc.CallOption) (*ListCountriesOutput, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListCountriesOutput)
	err := c.cc.Invoke(ctx, Geography_ListCountries_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *geographyClient) DescribeRegion(ctx context.Context, in *DescribeRegionInput, opts ...grpc.CallOption) (*DescribeRegionOutput, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DescribeRegionOutput)
	err := c.cc.Invoke(ctx, Geography_DescribeRegion_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *geographyClient) ListRegions(ctx context.Context, in *ListRegionsInput, opts ...grpc.CallOption) (*ListRegionsOutput, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListRegionsOutput)
	err := c.cc.Invoke(ctx, Geography_ListRegions_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *geographyClient) DeprecatedListRegions(ctx context.Context, in *ListRegionsInput, opts ...grpc.CallOption) (*ListRegionsOutput, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListRegionsOutput)
	err := c.cc.Invoke(ctx, Geography_DeprecatedListRegions_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GeographyServer is the server API for Geography service.
// All implementations should embed UnimplementedGeographyServer
// for forward compatibility.
type GeographyServer interface {
	DescribeCountry(context.Context, *DescribeCountryInput) (*DescribeCountryOutput, error)
	ListCountries(context.Context, *ListCountriesInput) (*ListCountriesOutput, error)
	DescribeRegion(context.Context, *DescribeRegionInput) (*DescribeRegionOutput, error)
	ListRegions(context.Context, *ListRegionsInput) (*ListRegionsOutput, error)
	DeprecatedListRegions(context.Context, *ListRegionsInput) (*ListRegionsOutput, error)
}

// UnimplementedGeographyServer should be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedGeographyServer struct{}

func (UnimplementedGeographyServer) DescribeCountry(context.Context, *DescribeCountryInput) (*DescribeCountryOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DescribeCountry not implemented")
}
func (UnimplementedGeographyServer) ListCountries(context.Context, *ListCountriesInput) (*ListCountriesOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListCountries not implemented")
}
func (UnimplementedGeographyServer) DescribeRegion(context.Context, *DescribeRegionInput) (*DescribeRegionOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DescribeRegion not implemented")
}
func (UnimplementedGeographyServer) ListRegions(context.Context, *ListRegionsInput) (*ListRegionsOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListRegions not implemented")
}
func (UnimplementedGeographyServer) DeprecatedListRegions(context.Context, *ListRegionsInput) (*ListRegionsOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeprecatedListRegions not implemented")
}
func (UnimplementedGeographyServer) testEmbeddedByValue() {}

// UnsafeGeographyServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GeographyServer will
// result in compilation errors.
type UnsafeGeographyServer interface {
	mustEmbedUnimplementedGeographyServer()
}

func RegisterGeographyServer(s grpc.ServiceRegistrar, srv GeographyServer) {
	// If the following call pancis, it indicates UnimplementedGeographyServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
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
		FullMethod: Geography_DescribeCountry_FullMethodName,
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
		FullMethod: Geography_ListCountries_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GeographyServer).ListCountries(ctx, req.(*ListCountriesInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _Geography_DescribeRegion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DescribeRegionInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GeographyServer).DescribeRegion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Geography_DescribeRegion_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GeographyServer).DescribeRegion(ctx, req.(*DescribeRegionInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _Geography_ListRegions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRegionsInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GeographyServer).ListRegions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Geography_ListRegions_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GeographyServer).ListRegions(ctx, req.(*ListRegionsInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _Geography_DeprecatedListRegions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRegionsInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GeographyServer).DeprecatedListRegions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Geography_DeprecatedListRegions_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GeographyServer).DeprecatedListRegions(ctx, req.(*ListRegionsInput))
	}
	return interceptor(ctx, in, info, handler)
}

// Geography_ServiceDesc is the grpc.ServiceDesc for Geography service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Geography_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "eolymp.geography.Geography",
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
		{
			MethodName: "DescribeRegion",
			Handler:    _Geography_DescribeRegion_Handler,
		},
		{
			MethodName: "ListRegions",
			Handler:    _Geography_ListRegions_Handler,
		},
		{
			MethodName: "DeprecatedListRegions",
			Handler:    _Geography_DeprecatedListRegions_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "eolymp/geography/geography.proto",
}
