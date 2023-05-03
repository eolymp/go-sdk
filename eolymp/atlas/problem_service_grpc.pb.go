// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.18.1
// source: eolymp/atlas/problem_service.proto

package atlas

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
	ProblemService_DeleteProblem_FullMethodName    = "/eolymp.atlas.ProblemService/DeleteProblem"
	ProblemService_UpdateProblem_FullMethodName    = "/eolymp.atlas.ProblemService/UpdateProblem"
	ProblemService_DescribeProblem_FullMethodName  = "/eolymp.atlas.ProblemService/DescribeProblem"
	ProblemService_UpdateVisibility_FullMethodName = "/eolymp.atlas.ProblemService/UpdateVisibility"
	ProblemService_UpdatePrivacy_FullMethodName    = "/eolymp.atlas.ProblemService/UpdatePrivacy"
	ProblemService_ListVersions_FullMethodName     = "/eolymp.atlas.ProblemService/ListVersions"
)

// ProblemServiceClient is the client API for ProblemService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProblemServiceClient interface {
	DeleteProblem(ctx context.Context, in *DeleteProblemInput, opts ...grpc.CallOption) (*DeleteProblemOutput, error)
	UpdateProblem(ctx context.Context, in *UpdateProblemInput, opts ...grpc.CallOption) (*UpdateProblemOutput, error)
	DescribeProblem(ctx context.Context, in *DescribeProblemInput, opts ...grpc.CallOption) (*DescribeProblemOutput, error)
	// deprecated: use UpdateProblem instead
	UpdateVisibility(ctx context.Context, in *UpdateVisibilityInput, opts ...grpc.CallOption) (*UpdateVisibilityOutput, error)
	// deprecated: use UpdateProblem instead
	UpdatePrivacy(ctx context.Context, in *UpdatePrivacyInput, opts ...grpc.CallOption) (*UpdatePrivacyOutput, error)
	ListVersions(ctx context.Context, in *ListVersionsInput, opts ...grpc.CallOption) (*ListVersionsOutput, error)
}

type problemServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewProblemServiceClient(cc grpc.ClientConnInterface) ProblemServiceClient {
	return &problemServiceClient{cc}
}

func (c *problemServiceClient) DeleteProblem(ctx context.Context, in *DeleteProblemInput, opts ...grpc.CallOption) (*DeleteProblemOutput, error) {
	out := new(DeleteProblemOutput)
	err := c.cc.Invoke(ctx, ProblemService_DeleteProblem_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *problemServiceClient) UpdateProblem(ctx context.Context, in *UpdateProblemInput, opts ...grpc.CallOption) (*UpdateProblemOutput, error) {
	out := new(UpdateProblemOutput)
	err := c.cc.Invoke(ctx, ProblemService_UpdateProblem_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *problemServiceClient) DescribeProblem(ctx context.Context, in *DescribeProblemInput, opts ...grpc.CallOption) (*DescribeProblemOutput, error) {
	out := new(DescribeProblemOutput)
	err := c.cc.Invoke(ctx, ProblemService_DescribeProblem_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *problemServiceClient) UpdateVisibility(ctx context.Context, in *UpdateVisibilityInput, opts ...grpc.CallOption) (*UpdateVisibilityOutput, error) {
	out := new(UpdateVisibilityOutput)
	err := c.cc.Invoke(ctx, ProblemService_UpdateVisibility_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *problemServiceClient) UpdatePrivacy(ctx context.Context, in *UpdatePrivacyInput, opts ...grpc.CallOption) (*UpdatePrivacyOutput, error) {
	out := new(UpdatePrivacyOutput)
	err := c.cc.Invoke(ctx, ProblemService_UpdatePrivacy_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *problemServiceClient) ListVersions(ctx context.Context, in *ListVersionsInput, opts ...grpc.CallOption) (*ListVersionsOutput, error) {
	out := new(ListVersionsOutput)
	err := c.cc.Invoke(ctx, ProblemService_ListVersions_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProblemServiceServer is the server API for ProblemService service.
// All implementations should embed UnimplementedProblemServiceServer
// for forward compatibility
type ProblemServiceServer interface {
	DeleteProblem(context.Context, *DeleteProblemInput) (*DeleteProblemOutput, error)
	UpdateProblem(context.Context, *UpdateProblemInput) (*UpdateProblemOutput, error)
	DescribeProblem(context.Context, *DescribeProblemInput) (*DescribeProblemOutput, error)
	// deprecated: use UpdateProblem instead
	UpdateVisibility(context.Context, *UpdateVisibilityInput) (*UpdateVisibilityOutput, error)
	// deprecated: use UpdateProblem instead
	UpdatePrivacy(context.Context, *UpdatePrivacyInput) (*UpdatePrivacyOutput, error)
	ListVersions(context.Context, *ListVersionsInput) (*ListVersionsOutput, error)
}

// UnimplementedProblemServiceServer should be embedded to have forward compatible implementations.
type UnimplementedProblemServiceServer struct {
}

func (UnimplementedProblemServiceServer) DeleteProblem(context.Context, *DeleteProblemInput) (*DeleteProblemOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteProblem not implemented")
}
func (UnimplementedProblemServiceServer) UpdateProblem(context.Context, *UpdateProblemInput) (*UpdateProblemOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateProblem not implemented")
}
func (UnimplementedProblemServiceServer) DescribeProblem(context.Context, *DescribeProblemInput) (*DescribeProblemOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DescribeProblem not implemented")
}
func (UnimplementedProblemServiceServer) UpdateVisibility(context.Context, *UpdateVisibilityInput) (*UpdateVisibilityOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateVisibility not implemented")
}
func (UnimplementedProblemServiceServer) UpdatePrivacy(context.Context, *UpdatePrivacyInput) (*UpdatePrivacyOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePrivacy not implemented")
}
func (UnimplementedProblemServiceServer) ListVersions(context.Context, *ListVersionsInput) (*ListVersionsOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListVersions not implemented")
}

// UnsafeProblemServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProblemServiceServer will
// result in compilation errors.
type UnsafeProblemServiceServer interface {
	mustEmbedUnimplementedProblemServiceServer()
}

func RegisterProblemServiceServer(s grpc.ServiceRegistrar, srv ProblemServiceServer) {
	s.RegisterService(&ProblemService_ServiceDesc, srv)
}

func _ProblemService_DeleteProblem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteProblemInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProblemServiceServer).DeleteProblem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProblemService_DeleteProblem_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProblemServiceServer).DeleteProblem(ctx, req.(*DeleteProblemInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProblemService_UpdateProblem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateProblemInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProblemServiceServer).UpdateProblem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProblemService_UpdateProblem_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProblemServiceServer).UpdateProblem(ctx, req.(*UpdateProblemInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProblemService_DescribeProblem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DescribeProblemInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProblemServiceServer).DescribeProblem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProblemService_DescribeProblem_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProblemServiceServer).DescribeProblem(ctx, req.(*DescribeProblemInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProblemService_UpdateVisibility_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateVisibilityInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProblemServiceServer).UpdateVisibility(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProblemService_UpdateVisibility_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProblemServiceServer).UpdateVisibility(ctx, req.(*UpdateVisibilityInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProblemService_UpdatePrivacy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePrivacyInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProblemServiceServer).UpdatePrivacy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProblemService_UpdatePrivacy_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProblemServiceServer).UpdatePrivacy(ctx, req.(*UpdatePrivacyInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProblemService_ListVersions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListVersionsInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProblemServiceServer).ListVersions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProblemService_ListVersions_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProblemServiceServer).ListVersions(ctx, req.(*ListVersionsInput))
	}
	return interceptor(ctx, in, info, handler)
}

// ProblemService_ServiceDesc is the grpc.ServiceDesc for ProblemService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ProblemService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "eolymp.atlas.ProblemService",
	HandlerType: (*ProblemServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DeleteProblem",
			Handler:    _ProblemService_DeleteProblem_Handler,
		},
		{
			MethodName: "UpdateProblem",
			Handler:    _ProblemService_UpdateProblem_Handler,
		},
		{
			MethodName: "DescribeProblem",
			Handler:    _ProblemService_DescribeProblem_Handler,
		},
		{
			MethodName: "UpdateVisibility",
			Handler:    _ProblemService_UpdateVisibility_Handler,
		},
		{
			MethodName: "UpdatePrivacy",
			Handler:    _ProblemService_UpdatePrivacy_Handler,
		},
		{
			MethodName: "ListVersions",
			Handler:    _ProblemService_ListVersions_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "eolymp/atlas/problem_service.proto",
}