// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v4.24.4
// source: eolymp/certificate/certificate_service.proto

package certificate

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
	CertificateService_CreateCertificate_FullMethodName   = "/eolymp.certificate.CertificateService/CreateCertificate"
	CertificateService_VoidCertificate_FullMethodName     = "/eolymp.certificate.CertificateService/VoidCertificate"
	CertificateService_DescribeCertificate_FullMethodName = "/eolymp.certificate.CertificateService/DescribeCertificate"
	CertificateService_ListCertificates_FullMethodName    = "/eolymp.certificate.CertificateService/ListCertificates"
)

// CertificateServiceClient is the client API for CertificateService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CertificateServiceClient interface {
	CreateCertificate(ctx context.Context, in *CreateCertificateInput, opts ...grpc.CallOption) (*CreateCertificateOutput, error)
	VoidCertificate(ctx context.Context, in *VoidCertificateInput, opts ...grpc.CallOption) (*VoidCertificateOutput, error)
	DescribeCertificate(ctx context.Context, in *DescribeCertificateInput, opts ...grpc.CallOption) (*DescribeCertificateOutput, error)
	ListCertificates(ctx context.Context, in *ListCertificatesInput, opts ...grpc.CallOption) (*ListCertificatesOutput, error)
}

type certificateServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCertificateServiceClient(cc grpc.ClientConnInterface) CertificateServiceClient {
	return &certificateServiceClient{cc}
}

func (c *certificateServiceClient) CreateCertificate(ctx context.Context, in *CreateCertificateInput, opts ...grpc.CallOption) (*CreateCertificateOutput, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateCertificateOutput)
	err := c.cc.Invoke(ctx, CertificateService_CreateCertificate_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *certificateServiceClient) VoidCertificate(ctx context.Context, in *VoidCertificateInput, opts ...grpc.CallOption) (*VoidCertificateOutput, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(VoidCertificateOutput)
	err := c.cc.Invoke(ctx, CertificateService_VoidCertificate_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *certificateServiceClient) DescribeCertificate(ctx context.Context, in *DescribeCertificateInput, opts ...grpc.CallOption) (*DescribeCertificateOutput, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DescribeCertificateOutput)
	err := c.cc.Invoke(ctx, CertificateService_DescribeCertificate_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *certificateServiceClient) ListCertificates(ctx context.Context, in *ListCertificatesInput, opts ...grpc.CallOption) (*ListCertificatesOutput, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListCertificatesOutput)
	err := c.cc.Invoke(ctx, CertificateService_ListCertificates_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CertificateServiceServer is the server API for CertificateService service.
// All implementations should embed UnimplementedCertificateServiceServer
// for forward compatibility.
type CertificateServiceServer interface {
	CreateCertificate(context.Context, *CreateCertificateInput) (*CreateCertificateOutput, error)
	VoidCertificate(context.Context, *VoidCertificateInput) (*VoidCertificateOutput, error)
	DescribeCertificate(context.Context, *DescribeCertificateInput) (*DescribeCertificateOutput, error)
	ListCertificates(context.Context, *ListCertificatesInput) (*ListCertificatesOutput, error)
}

// UnimplementedCertificateServiceServer should be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedCertificateServiceServer struct{}

func (UnimplementedCertificateServiceServer) CreateCertificate(context.Context, *CreateCertificateInput) (*CreateCertificateOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCertificate not implemented")
}
func (UnimplementedCertificateServiceServer) VoidCertificate(context.Context, *VoidCertificateInput) (*VoidCertificateOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VoidCertificate not implemented")
}
func (UnimplementedCertificateServiceServer) DescribeCertificate(context.Context, *DescribeCertificateInput) (*DescribeCertificateOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DescribeCertificate not implemented")
}
func (UnimplementedCertificateServiceServer) ListCertificates(context.Context, *ListCertificatesInput) (*ListCertificatesOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListCertificates not implemented")
}
func (UnimplementedCertificateServiceServer) testEmbeddedByValue() {}

// UnsafeCertificateServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CertificateServiceServer will
// result in compilation errors.
type UnsafeCertificateServiceServer interface {
	mustEmbedUnimplementedCertificateServiceServer()
}

func RegisterCertificateServiceServer(s grpc.ServiceRegistrar, srv CertificateServiceServer) {
	// If the following call pancis, it indicates UnimplementedCertificateServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&CertificateService_ServiceDesc, srv)
}

func _CertificateService_CreateCertificate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCertificateInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CertificateServiceServer).CreateCertificate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CertificateService_CreateCertificate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CertificateServiceServer).CreateCertificate(ctx, req.(*CreateCertificateInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _CertificateService_VoidCertificate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VoidCertificateInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CertificateServiceServer).VoidCertificate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CertificateService_VoidCertificate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CertificateServiceServer).VoidCertificate(ctx, req.(*VoidCertificateInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _CertificateService_DescribeCertificate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DescribeCertificateInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CertificateServiceServer).DescribeCertificate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CertificateService_DescribeCertificate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CertificateServiceServer).DescribeCertificate(ctx, req.(*DescribeCertificateInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _CertificateService_ListCertificates_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListCertificatesInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CertificateServiceServer).ListCertificates(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CertificateService_ListCertificates_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CertificateServiceServer).ListCertificates(ctx, req.(*ListCertificatesInput))
	}
	return interceptor(ctx, in, info, handler)
}

// CertificateService_ServiceDesc is the grpc.ServiceDesc for CertificateService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CertificateService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "eolymp.certificate.CertificateService",
	HandlerType: (*CertificateServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateCertificate",
			Handler:    _CertificateService_CreateCertificate_Handler,
		},
		{
			MethodName: "VoidCertificate",
			Handler:    _CertificateService_VoidCertificate_Handler,
		},
		{
			MethodName: "DescribeCertificate",
			Handler:    _CertificateService_DescribeCertificate_Handler,
		},
		{
			MethodName: "ListCertificates",
			Handler:    _CertificateService_ListCertificates_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "eolymp/certificate/certificate_service.proto",
}
