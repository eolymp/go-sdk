// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v4.24.4
// source: eolymp/universe/plan_service.proto

package universe

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
	PlanService_DescribePlan_FullMethodName = "/eolymp.universe.PlanService/DescribePlan"
	PlanService_ListPlans_FullMethodName    = "/eolymp.universe.PlanService/ListPlans"
)

// PlanServiceClient is the client API for PlanService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PlanServiceClient interface {
	DescribePlan(ctx context.Context, in *DescribePlanInput, opts ...grpc.CallOption) (*DescribePlanOutput, error)
	ListPlans(ctx context.Context, in *ListPlansInput, opts ...grpc.CallOption) (*ListPlansOutput, error)
}

type planServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPlanServiceClient(cc grpc.ClientConnInterface) PlanServiceClient {
	return &planServiceClient{cc}
}

func (c *planServiceClient) DescribePlan(ctx context.Context, in *DescribePlanInput, opts ...grpc.CallOption) (*DescribePlanOutput, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DescribePlanOutput)
	err := c.cc.Invoke(ctx, PlanService_DescribePlan_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *planServiceClient) ListPlans(ctx context.Context, in *ListPlansInput, opts ...grpc.CallOption) (*ListPlansOutput, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListPlansOutput)
	err := c.cc.Invoke(ctx, PlanService_ListPlans_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PlanServiceServer is the server API for PlanService service.
// All implementations should embed UnimplementedPlanServiceServer
// for forward compatibility.
type PlanServiceServer interface {
	DescribePlan(context.Context, *DescribePlanInput) (*DescribePlanOutput, error)
	ListPlans(context.Context, *ListPlansInput) (*ListPlansOutput, error)
}

// UnimplementedPlanServiceServer should be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedPlanServiceServer struct{}

func (UnimplementedPlanServiceServer) DescribePlan(context.Context, *DescribePlanInput) (*DescribePlanOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DescribePlan not implemented")
}
func (UnimplementedPlanServiceServer) ListPlans(context.Context, *ListPlansInput) (*ListPlansOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListPlans not implemented")
}
func (UnimplementedPlanServiceServer) testEmbeddedByValue() {}

// UnsafePlanServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PlanServiceServer will
// result in compilation errors.
type UnsafePlanServiceServer interface {
	mustEmbedUnimplementedPlanServiceServer()
}

func RegisterPlanServiceServer(s grpc.ServiceRegistrar, srv PlanServiceServer) {
	// If the following call pancis, it indicates UnimplementedPlanServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&PlanService_ServiceDesc, srv)
}

func _PlanService_DescribePlan_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DescribePlanInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlanServiceServer).DescribePlan(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PlanService_DescribePlan_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlanServiceServer).DescribePlan(ctx, req.(*DescribePlanInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _PlanService_ListPlans_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListPlansInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlanServiceServer).ListPlans(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PlanService_ListPlans_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlanServiceServer).ListPlans(ctx, req.(*ListPlansInput))
	}
	return interceptor(ctx, in, info, handler)
}

// PlanService_ServiceDesc is the grpc.ServiceDesc for PlanService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PlanService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "eolymp.universe.PlanService",
	HandlerType: (*PlanServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DescribePlan",
			Handler:    _PlanService_DescribePlan_Handler,
		},
		{
			MethodName: "ListPlans",
			Handler:    _PlanService_ListPlans_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "eolymp/universe/plan_service.proto",
}
