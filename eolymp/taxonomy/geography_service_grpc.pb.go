// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v4.24.4
// source: eolymp/taxonomy/geography_service.proto

package taxonomy

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
	GeographyService_ListCountries_FullMethodName   = "/eolymp.taxonomy.GeographyService/ListCountries"
	GeographyService_DescribeCountry_FullMethodName = "/eolymp.taxonomy.GeographyService/DescribeCountry"
	GeographyService_ListRegions_FullMethodName     = "/eolymp.taxonomy.GeographyService/ListRegions"
	GeographyService_DescribeRegion_FullMethodName  = "/eolymp.taxonomy.GeographyService/DescribeRegion"
)

// GeographyServiceClient is the client API for GeographyService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GeographyServiceClient interface {
	ListCountries(ctx context.Context, in *ListCountriesInput, opts ...grpc.CallOption) (*ListCountriesOutput, error)
	DescribeCountry(ctx context.Context, in *DescribeCountryInput, opts ...grpc.CallOption) (*DescribeCountryOutput, error)
	ListRegions(ctx context.Context, in *ListRegionsInput, opts ...grpc.CallOption) (*ListRegionsOutput, error)
	DescribeRegion(ctx context.Context, in *DescribeRegionInput, opts ...grpc.CallOption) (*DescribeRegionOutput, error)
}

type geographyServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewGeographyServiceClient(cc grpc.ClientConnInterface) GeographyServiceClient {
	return &geographyServiceClient{cc}
}

func (c *geographyServiceClient) ListCountries(ctx context.Context, in *ListCountriesInput, opts ...grpc.CallOption) (*ListCountriesOutput, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListCountriesOutput)
	err := c.cc.Invoke(ctx, GeographyService_ListCountries_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *geographyServiceClient) DescribeCountry(ctx context.Context, in *DescribeCountryInput, opts ...grpc.CallOption) (*DescribeCountryOutput, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DescribeCountryOutput)
	err := c.cc.Invoke(ctx, GeographyService_DescribeCountry_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *geographyServiceClient) ListRegions(ctx context.Context, in *ListRegionsInput, opts ...grpc.CallOption) (*ListRegionsOutput, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListRegionsOutput)
	err := c.cc.Invoke(ctx, GeographyService_ListRegions_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *geographyServiceClient) DescribeRegion(ctx context.Context, in *DescribeRegionInput, opts ...grpc.CallOption) (*DescribeRegionOutput, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DescribeRegionOutput)
	err := c.cc.Invoke(ctx, GeographyService_DescribeRegion_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GeographyServiceServer is the server API for GeographyService service.
// All implementations should embed UnimplementedGeographyServiceServer
// for forward compatibility.
type GeographyServiceServer interface {
	ListCountries(context.Context, *ListCountriesInput) (*ListCountriesOutput, error)
	DescribeCountry(context.Context, *DescribeCountryInput) (*DescribeCountryOutput, error)
	ListRegions(context.Context, *ListRegionsInput) (*ListRegionsOutput, error)
	DescribeRegion(context.Context, *DescribeRegionInput) (*DescribeRegionOutput, error)
}

// UnimplementedGeographyServiceServer should be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedGeographyServiceServer struct{}

func (UnimplementedGeographyServiceServer) ListCountries(context.Context, *ListCountriesInput) (*ListCountriesOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListCountries not implemented")
}
func (UnimplementedGeographyServiceServer) DescribeCountry(context.Context, *DescribeCountryInput) (*DescribeCountryOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DescribeCountry not implemented")
}
func (UnimplementedGeographyServiceServer) ListRegions(context.Context, *ListRegionsInput) (*ListRegionsOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListRegions not implemented")
}
func (UnimplementedGeographyServiceServer) DescribeRegion(context.Context, *DescribeRegionInput) (*DescribeRegionOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DescribeRegion not implemented")
}
func (UnimplementedGeographyServiceServer) testEmbeddedByValue() {}

// UnsafeGeographyServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GeographyServiceServer will
// result in compilation errors.
type UnsafeGeographyServiceServer interface {
	mustEmbedUnimplementedGeographyServiceServer()
}

func RegisterGeographyServiceServer(s grpc.ServiceRegistrar, srv GeographyServiceServer) {
	// If the following call pancis, it indicates UnimplementedGeographyServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&GeographyService_ServiceDesc, srv)
}

func _GeographyService_ListCountries_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListCountriesInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GeographyServiceServer).ListCountries(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GeographyService_ListCountries_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GeographyServiceServer).ListCountries(ctx, req.(*ListCountriesInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _GeographyService_DescribeCountry_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DescribeCountryInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GeographyServiceServer).DescribeCountry(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GeographyService_DescribeCountry_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GeographyServiceServer).DescribeCountry(ctx, req.(*DescribeCountryInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _GeographyService_ListRegions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRegionsInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GeographyServiceServer).ListRegions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GeographyService_ListRegions_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GeographyServiceServer).ListRegions(ctx, req.(*ListRegionsInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _GeographyService_DescribeRegion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DescribeRegionInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GeographyServiceServer).DescribeRegion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GeographyService_DescribeRegion_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GeographyServiceServer).DescribeRegion(ctx, req.(*DescribeRegionInput))
	}
	return interceptor(ctx, in, info, handler)
}

// GeographyService_ServiceDesc is the grpc.ServiceDesc for GeographyService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GeographyService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "eolymp.taxonomy.GeographyService",
	HandlerType: (*GeographyServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListCountries",
			Handler:    _GeographyService_ListCountries_Handler,
		},
		{
			MethodName: "DescribeCountry",
			Handler:    _GeographyService_DescribeCountry_Handler,
		},
		{
			MethodName: "ListRegions",
			Handler:    _GeographyService_ListRegions_Handler,
		},
		{
			MethodName: "DescribeRegion",
			Handler:    _GeographyService_DescribeRegion_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "eolymp/taxonomy/geography_service.proto",
}
