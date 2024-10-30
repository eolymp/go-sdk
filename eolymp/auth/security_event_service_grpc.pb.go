// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.21.12
// source: eolymp/auth/security_event_service.proto

package auth

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
	SecurityEventService_HandleSecurityEvent_FullMethodName = "/eolymp.auth.SecurityEventService/HandleSecurityEvent"
)

// SecurityEventServiceClient is the client API for SecurityEventService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SecurityEventServiceClient interface {
	HandleSecurityEvent(ctx context.Context, in *HandleSecurityEventInput, opts ...grpc.CallOption) (*HandleSecurityEventOutput, error)
}

type securityEventServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSecurityEventServiceClient(cc grpc.ClientConnInterface) SecurityEventServiceClient {
	return &securityEventServiceClient{cc}
}

func (c *securityEventServiceClient) HandleSecurityEvent(ctx context.Context, in *HandleSecurityEventInput, opts ...grpc.CallOption) (*HandleSecurityEventOutput, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(HandleSecurityEventOutput)
	err := c.cc.Invoke(ctx, SecurityEventService_HandleSecurityEvent_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SecurityEventServiceServer is the server API for SecurityEventService service.
// All implementations should embed UnimplementedSecurityEventServiceServer
// for forward compatibility.
type SecurityEventServiceServer interface {
	HandleSecurityEvent(context.Context, *HandleSecurityEventInput) (*HandleSecurityEventOutput, error)
}

// UnimplementedSecurityEventServiceServer should be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedSecurityEventServiceServer struct{}

func (UnimplementedSecurityEventServiceServer) HandleSecurityEvent(context.Context, *HandleSecurityEventInput) (*HandleSecurityEventOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HandleSecurityEvent not implemented")
}
func (UnimplementedSecurityEventServiceServer) testEmbeddedByValue() {}

// UnsafeSecurityEventServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SecurityEventServiceServer will
// result in compilation errors.
type UnsafeSecurityEventServiceServer interface {
	mustEmbedUnimplementedSecurityEventServiceServer()
}

func RegisterSecurityEventServiceServer(s grpc.ServiceRegistrar, srv SecurityEventServiceServer) {
	// If the following call pancis, it indicates UnimplementedSecurityEventServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&SecurityEventService_ServiceDesc, srv)
}

func _SecurityEventService_HandleSecurityEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HandleSecurityEventInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SecurityEventServiceServer).HandleSecurityEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SecurityEventService_HandleSecurityEvent_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SecurityEventServiceServer).HandleSecurityEvent(ctx, req.(*HandleSecurityEventInput))
	}
	return interceptor(ctx, in, info, handler)
}

// SecurityEventService_ServiceDesc is the grpc.ServiceDesc for SecurityEventService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SecurityEventService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "eolymp.auth.SecurityEventService",
	HandlerType: (*SecurityEventServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "HandleSecurityEvent",
			Handler:    _SecurityEventService_HandleSecurityEvent_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "eolymp/auth/security_event_service.proto",
}
