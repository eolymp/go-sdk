// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.4
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
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	ProblemService_CreateProblem_FullMethodName   = "/eolymp.atlas.ProblemService/CreateProblem"
	ProblemService_UpdateProblem_FullMethodName   = "/eolymp.atlas.ProblemService/UpdateProblem"
	ProblemService_DeleteProblem_FullMethodName   = "/eolymp.atlas.ProblemService/DeleteProblem"
	ProblemService_DescribeProblem_FullMethodName = "/eolymp.atlas.ProblemService/DescribeProblem"
	ProblemService_ListProblems_FullMethodName    = "/eolymp.atlas.ProblemService/ListProblems"
	ProblemService_SyncProblem_FullMethodName     = "/eolymp.atlas.ProblemService/SyncProblem"
	ProblemService_VoteProblem_FullMethodName     = "/eolymp.atlas.ProblemService/VoteProblem"
	ProblemService_ListVersions_FullMethodName    = "/eolymp.atlas.ProblemService/ListVersions"
	ProblemService_ListRuntimes_FullMethodName    = "/eolymp.atlas.ProblemService/ListRuntimes"
	ProblemService_ExportProblem_FullMethodName   = "/eolymp.atlas.ProblemService/ExportProblem"
)

// ProblemServiceClient is the client API for ProblemService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProblemServiceClient interface {
	CreateProblem(ctx context.Context, in *CreateProblemInput, opts ...grpc.CallOption) (*CreateProblemOutput, error)
	UpdateProblem(ctx context.Context, in *UpdateProblemInput, opts ...grpc.CallOption) (*UpdateProblemOutput, error)
	DeleteProblem(ctx context.Context, in *DeleteProblemInput, opts ...grpc.CallOption) (*DeleteProblemOutput, error)
	DescribeProblem(ctx context.Context, in *DescribeProblemInput, opts ...grpc.CallOption) (*DescribeProblemOutput, error)
	ListProblems(ctx context.Context, in *ListProblemsInput, opts ...grpc.CallOption) (*ListProblemsOutput, error)
	SyncProblem(ctx context.Context, in *SyncProblemInput, opts ...grpc.CallOption) (*SyncProblemOutput, error)
	VoteProblem(ctx context.Context, in *VoteProblemInput, opts ...grpc.CallOption) (*VoteProblemOutput, error)
	ListVersions(ctx context.Context, in *ListVersionsInput, opts ...grpc.CallOption) (*ListVersionsOutput, error)
	ListRuntimes(ctx context.Context, in *ListRuntimesInput, opts ...grpc.CallOption) (*ListRuntimesOutput, error)
	ExportProblem(ctx context.Context, in *ExportProblemInput, opts ...grpc.CallOption) (*ExportProblemOutput, error)
}

type problemServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewProblemServiceClient(cc grpc.ClientConnInterface) ProblemServiceClient {
	return &problemServiceClient{cc}
}

func (c *problemServiceClient) CreateProblem(ctx context.Context, in *CreateProblemInput, opts ...grpc.CallOption) (*CreateProblemOutput, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateProblemOutput)
	err := c.cc.Invoke(ctx, ProblemService_CreateProblem_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *problemServiceClient) UpdateProblem(ctx context.Context, in *UpdateProblemInput, opts ...grpc.CallOption) (*UpdateProblemOutput, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateProblemOutput)
	err := c.cc.Invoke(ctx, ProblemService_UpdateProblem_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *problemServiceClient) DeleteProblem(ctx context.Context, in *DeleteProblemInput, opts ...grpc.CallOption) (*DeleteProblemOutput, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteProblemOutput)
	err := c.cc.Invoke(ctx, ProblemService_DeleteProblem_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *problemServiceClient) DescribeProblem(ctx context.Context, in *DescribeProblemInput, opts ...grpc.CallOption) (*DescribeProblemOutput, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DescribeProblemOutput)
	err := c.cc.Invoke(ctx, ProblemService_DescribeProblem_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *problemServiceClient) ListProblems(ctx context.Context, in *ListProblemsInput, opts ...grpc.CallOption) (*ListProblemsOutput, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListProblemsOutput)
	err := c.cc.Invoke(ctx, ProblemService_ListProblems_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *problemServiceClient) SyncProblem(ctx context.Context, in *SyncProblemInput, opts ...grpc.CallOption) (*SyncProblemOutput, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SyncProblemOutput)
	err := c.cc.Invoke(ctx, ProblemService_SyncProblem_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *problemServiceClient) VoteProblem(ctx context.Context, in *VoteProblemInput, opts ...grpc.CallOption) (*VoteProblemOutput, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(VoteProblemOutput)
	err := c.cc.Invoke(ctx, ProblemService_VoteProblem_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *problemServiceClient) ListVersions(ctx context.Context, in *ListVersionsInput, opts ...grpc.CallOption) (*ListVersionsOutput, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListVersionsOutput)
	err := c.cc.Invoke(ctx, ProblemService_ListVersions_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *problemServiceClient) ListRuntimes(ctx context.Context, in *ListRuntimesInput, opts ...grpc.CallOption) (*ListRuntimesOutput, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListRuntimesOutput)
	err := c.cc.Invoke(ctx, ProblemService_ListRuntimes_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *problemServiceClient) ExportProblem(ctx context.Context, in *ExportProblemInput, opts ...grpc.CallOption) (*ExportProblemOutput, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ExportProblemOutput)
	err := c.cc.Invoke(ctx, ProblemService_ExportProblem_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProblemServiceServer is the server API for ProblemService service.
// All implementations should embed UnimplementedProblemServiceServer
// for forward compatibility.
type ProblemServiceServer interface {
	CreateProblem(context.Context, *CreateProblemInput) (*CreateProblemOutput, error)
	UpdateProblem(context.Context, *UpdateProblemInput) (*UpdateProblemOutput, error)
	DeleteProblem(context.Context, *DeleteProblemInput) (*DeleteProblemOutput, error)
	DescribeProblem(context.Context, *DescribeProblemInput) (*DescribeProblemOutput, error)
	ListProblems(context.Context, *ListProblemsInput) (*ListProblemsOutput, error)
	SyncProblem(context.Context, *SyncProblemInput) (*SyncProblemOutput, error)
	VoteProblem(context.Context, *VoteProblemInput) (*VoteProblemOutput, error)
	ListVersions(context.Context, *ListVersionsInput) (*ListVersionsOutput, error)
	ListRuntimes(context.Context, *ListRuntimesInput) (*ListRuntimesOutput, error)
	ExportProblem(context.Context, *ExportProblemInput) (*ExportProblemOutput, error)
}

// UnimplementedProblemServiceServer should be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedProblemServiceServer struct{}

func (UnimplementedProblemServiceServer) CreateProblem(context.Context, *CreateProblemInput) (*CreateProblemOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateProblem not implemented")
}
func (UnimplementedProblemServiceServer) UpdateProblem(context.Context, *UpdateProblemInput) (*UpdateProblemOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateProblem not implemented")
}
func (UnimplementedProblemServiceServer) DeleteProblem(context.Context, *DeleteProblemInput) (*DeleteProblemOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteProblem not implemented")
}
func (UnimplementedProblemServiceServer) DescribeProblem(context.Context, *DescribeProblemInput) (*DescribeProblemOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DescribeProblem not implemented")
}
func (UnimplementedProblemServiceServer) ListProblems(context.Context, *ListProblemsInput) (*ListProblemsOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListProblems not implemented")
}
func (UnimplementedProblemServiceServer) SyncProblem(context.Context, *SyncProblemInput) (*SyncProblemOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SyncProblem not implemented")
}
func (UnimplementedProblemServiceServer) VoteProblem(context.Context, *VoteProblemInput) (*VoteProblemOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VoteProblem not implemented")
}
func (UnimplementedProblemServiceServer) ListVersions(context.Context, *ListVersionsInput) (*ListVersionsOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListVersions not implemented")
}
func (UnimplementedProblemServiceServer) ListRuntimes(context.Context, *ListRuntimesInput) (*ListRuntimesOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListRuntimes not implemented")
}
func (UnimplementedProblemServiceServer) ExportProblem(context.Context, *ExportProblemInput) (*ExportProblemOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExportProblem not implemented")
}
func (UnimplementedProblemServiceServer) testEmbeddedByValue() {}

// UnsafeProblemServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProblemServiceServer will
// result in compilation errors.
type UnsafeProblemServiceServer interface {
	mustEmbedUnimplementedProblemServiceServer()
}

func RegisterProblemServiceServer(s grpc.ServiceRegistrar, srv ProblemServiceServer) {
	// If the following call pancis, it indicates UnimplementedProblemServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&ProblemService_ServiceDesc, srv)
}

func _ProblemService_CreateProblem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateProblemInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProblemServiceServer).CreateProblem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProblemService_CreateProblem_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProblemServiceServer).CreateProblem(ctx, req.(*CreateProblemInput))
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

func _ProblemService_ListProblems_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListProblemsInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProblemServiceServer).ListProblems(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProblemService_ListProblems_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProblemServiceServer).ListProblems(ctx, req.(*ListProblemsInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProblemService_SyncProblem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SyncProblemInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProblemServiceServer).SyncProblem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProblemService_SyncProblem_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProblemServiceServer).SyncProblem(ctx, req.(*SyncProblemInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProblemService_VoteProblem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VoteProblemInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProblemServiceServer).VoteProblem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProblemService_VoteProblem_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProblemServiceServer).VoteProblem(ctx, req.(*VoteProblemInput))
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

func _ProblemService_ListRuntimes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRuntimesInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProblemServiceServer).ListRuntimes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProblemService_ListRuntimes_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProblemServiceServer).ListRuntimes(ctx, req.(*ListRuntimesInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProblemService_ExportProblem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExportProblemInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProblemServiceServer).ExportProblem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProblemService_ExportProblem_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProblemServiceServer).ExportProblem(ctx, req.(*ExportProblemInput))
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
			MethodName: "CreateProblem",
			Handler:    _ProblemService_CreateProblem_Handler,
		},
		{
			MethodName: "UpdateProblem",
			Handler:    _ProblemService_UpdateProblem_Handler,
		},
		{
			MethodName: "DeleteProblem",
			Handler:    _ProblemService_DeleteProblem_Handler,
		},
		{
			MethodName: "DescribeProblem",
			Handler:    _ProblemService_DescribeProblem_Handler,
		},
		{
			MethodName: "ListProblems",
			Handler:    _ProblemService_ListProblems_Handler,
		},
		{
			MethodName: "SyncProblem",
			Handler:    _ProblemService_SyncProblem_Handler,
		},
		{
			MethodName: "VoteProblem",
			Handler:    _ProblemService_VoteProblem_Handler,
		},
		{
			MethodName: "ListVersions",
			Handler:    _ProblemService_ListVersions_Handler,
		},
		{
			MethodName: "ListRuntimes",
			Handler:    _ProblemService_ListRuntimes_Handler,
		},
		{
			MethodName: "ExportProblem",
			Handler:    _ProblemService_ExportProblem_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "eolymp/atlas/problem_service.proto",
}
