// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v4.24.4
// source: eolymp/judge/contest_service.proto

package judge

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
	ContestService_CreateContest_FullMethodName        = "/eolymp.judge.ContestService/CreateContest"
	ContestService_DeleteContest_FullMethodName        = "/eolymp.judge.ContestService/DeleteContest"
	ContestService_UpdateContest_FullMethodName        = "/eolymp.judge.ContestService/UpdateContest"
	ContestService_CopyContest_FullMethodName          = "/eolymp.judge.ContestService/CopyContest"
	ContestService_DescribeContest_FullMethodName      = "/eolymp.judge.ContestService/DescribeContest"
	ContestService_ListContests_FullMethodName         = "/eolymp.judge.ContestService/ListContests"
	ContestService_OpenContest_FullMethodName          = "/eolymp.judge.ContestService/OpenContest"
	ContestService_CloseContest_FullMethodName         = "/eolymp.judge.ContestService/CloseContest"
	ContestService_SuspendContest_FullMethodName       = "/eolymp.judge.ContestService/SuspendContest"
	ContestService_FreezeContest_FullMethodName        = "/eolymp.judge.ContestService/FreezeContest"
	ContestService_FinalizeContest_FullMethodName      = "/eolymp.judge.ContestService/FinalizeContest"
	ContestService_ResumeContest_FullMethodName        = "/eolymp.judge.ContestService/ResumeContest"
	ContestService_WatchContest_FullMethodName         = "/eolymp.judge.ContestService/WatchContest"
	ContestService_ListActivities_FullMethodName       = "/eolymp.judge.ContestService/ListActivities"
	ContestService_DescribeContestUsage_FullMethodName = "/eolymp.judge.ContestService/DescribeContestUsage"
)

// ContestServiceClient is the client API for ContestService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ContestServiceClient interface {
	CreateContest(ctx context.Context, in *CreateContestInput, opts ...grpc.CallOption) (*CreateContestOutput, error)
	DeleteContest(ctx context.Context, in *DeleteContestInput, opts ...grpc.CallOption) (*DeleteContestOutput, error)
	UpdateContest(ctx context.Context, in *UpdateContestInput, opts ...grpc.CallOption) (*UpdateContestOutput, error)
	CopyContest(ctx context.Context, in *CopyContestInput, opts ...grpc.CallOption) (*CopyContestOutput, error)
	DescribeContest(ctx context.Context, in *DescribeContestInput, opts ...grpc.CallOption) (*DescribeContestOutput, error)
	ListContests(ctx context.Context, in *ListContestsInput, opts ...grpc.CallOption) (*ListContestsOutput, error)
	// Force-starts scheduled contest, this call also automatically changes starts_at to current time and adjusts
	// ends_at to match original date range of the contest.
	OpenContest(ctx context.Context, in *OpenContestInput, opts ...grpc.CallOption) (*OpenContestOutput, error)
	// Force-finishes open contest, this method automatically changes ends_at to current time.
	CloseContest(ctx context.Context, in *CloseContestInput, opts ...grpc.CallOption) (*CloseContestOutput, error)
	// Temporarily stop contest and block participant's interface
	// Use ResumeContest to switch contest back to a normal mode.
	SuspendContest(ctx context.Context, in *SuspendContestInput, opts ...grpc.CallOption) (*SuspendContestOutput, error)
	// Temporarily restrict submission function
	// Use ResumeContest to switch contest back to a normal mode.
	FreezeContest(ctx context.Context, in *FreezeContestInput, opts ...grpc.CallOption) (*FreezeContestOutput, error)
	// Finalize contest results.
	// This action finalizes contest standings, issues participation certificates and awards medals etc.
	// After finalizing a contest it is not possible to change any parameters or results.
	// This action is irreversible.
	FinalizeContest(ctx context.Context, in *FinalizeContestInput, opts ...grpc.CallOption) (*FinalizeContestOutput, error)
	// Re-start suspended or frozen contest
	ResumeContest(ctx context.Context, in *ResumeContestInput, opts ...grpc.CallOption) (*ResumeContestOutput, error)
	WatchContest(ctx context.Context, in *WatchContestInput, opts ...grpc.CallOption) (grpc.ServerStreamingClient[WatchContestOutput], error)
	ListActivities(ctx context.Context, in *ListActivitiesInput, opts ...grpc.CallOption) (*ListActivitiesOutput, error)
	DescribeContestUsage(ctx context.Context, in *DescribeContestUsageInput, opts ...grpc.CallOption) (*DescribeContestUsageOutput, error)
}

type contestServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewContestServiceClient(cc grpc.ClientConnInterface) ContestServiceClient {
	return &contestServiceClient{cc}
}

func (c *contestServiceClient) CreateContest(ctx context.Context, in *CreateContestInput, opts ...grpc.CallOption) (*CreateContestOutput, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateContestOutput)
	err := c.cc.Invoke(ctx, ContestService_CreateContest_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contestServiceClient) DeleteContest(ctx context.Context, in *DeleteContestInput, opts ...grpc.CallOption) (*DeleteContestOutput, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteContestOutput)
	err := c.cc.Invoke(ctx, ContestService_DeleteContest_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contestServiceClient) UpdateContest(ctx context.Context, in *UpdateContestInput, opts ...grpc.CallOption) (*UpdateContestOutput, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateContestOutput)
	err := c.cc.Invoke(ctx, ContestService_UpdateContest_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contestServiceClient) CopyContest(ctx context.Context, in *CopyContestInput, opts ...grpc.CallOption) (*CopyContestOutput, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CopyContestOutput)
	err := c.cc.Invoke(ctx, ContestService_CopyContest_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contestServiceClient) DescribeContest(ctx context.Context, in *DescribeContestInput, opts ...grpc.CallOption) (*DescribeContestOutput, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DescribeContestOutput)
	err := c.cc.Invoke(ctx, ContestService_DescribeContest_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contestServiceClient) ListContests(ctx context.Context, in *ListContestsInput, opts ...grpc.CallOption) (*ListContestsOutput, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListContestsOutput)
	err := c.cc.Invoke(ctx, ContestService_ListContests_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contestServiceClient) OpenContest(ctx context.Context, in *OpenContestInput, opts ...grpc.CallOption) (*OpenContestOutput, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(OpenContestOutput)
	err := c.cc.Invoke(ctx, ContestService_OpenContest_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contestServiceClient) CloseContest(ctx context.Context, in *CloseContestInput, opts ...grpc.CallOption) (*CloseContestOutput, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CloseContestOutput)
	err := c.cc.Invoke(ctx, ContestService_CloseContest_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contestServiceClient) SuspendContest(ctx context.Context, in *SuspendContestInput, opts ...grpc.CallOption) (*SuspendContestOutput, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SuspendContestOutput)
	err := c.cc.Invoke(ctx, ContestService_SuspendContest_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contestServiceClient) FreezeContest(ctx context.Context, in *FreezeContestInput, opts ...grpc.CallOption) (*FreezeContestOutput, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(FreezeContestOutput)
	err := c.cc.Invoke(ctx, ContestService_FreezeContest_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contestServiceClient) FinalizeContest(ctx context.Context, in *FinalizeContestInput, opts ...grpc.CallOption) (*FinalizeContestOutput, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(FinalizeContestOutput)
	err := c.cc.Invoke(ctx, ContestService_FinalizeContest_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contestServiceClient) ResumeContest(ctx context.Context, in *ResumeContestInput, opts ...grpc.CallOption) (*ResumeContestOutput, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ResumeContestOutput)
	err := c.cc.Invoke(ctx, ContestService_ResumeContest_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contestServiceClient) WatchContest(ctx context.Context, in *WatchContestInput, opts ...grpc.CallOption) (grpc.ServerStreamingClient[WatchContestOutput], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &ContestService_ServiceDesc.Streams[0], ContestService_WatchContest_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[WatchContestInput, WatchContestOutput]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type ContestService_WatchContestClient = grpc.ServerStreamingClient[WatchContestOutput]

func (c *contestServiceClient) ListActivities(ctx context.Context, in *ListActivitiesInput, opts ...grpc.CallOption) (*ListActivitiesOutput, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListActivitiesOutput)
	err := c.cc.Invoke(ctx, ContestService_ListActivities_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contestServiceClient) DescribeContestUsage(ctx context.Context, in *DescribeContestUsageInput, opts ...grpc.CallOption) (*DescribeContestUsageOutput, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DescribeContestUsageOutput)
	err := c.cc.Invoke(ctx, ContestService_DescribeContestUsage_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ContestServiceServer is the server API for ContestService service.
// All implementations should embed UnimplementedContestServiceServer
// for forward compatibility.
type ContestServiceServer interface {
	CreateContest(context.Context, *CreateContestInput) (*CreateContestOutput, error)
	DeleteContest(context.Context, *DeleteContestInput) (*DeleteContestOutput, error)
	UpdateContest(context.Context, *UpdateContestInput) (*UpdateContestOutput, error)
	CopyContest(context.Context, *CopyContestInput) (*CopyContestOutput, error)
	DescribeContest(context.Context, *DescribeContestInput) (*DescribeContestOutput, error)
	ListContests(context.Context, *ListContestsInput) (*ListContestsOutput, error)
	// Force-starts scheduled contest, this call also automatically changes starts_at to current time and adjusts
	// ends_at to match original date range of the contest.
	OpenContest(context.Context, *OpenContestInput) (*OpenContestOutput, error)
	// Force-finishes open contest, this method automatically changes ends_at to current time.
	CloseContest(context.Context, *CloseContestInput) (*CloseContestOutput, error)
	// Temporarily stop contest and block participant's interface
	// Use ResumeContest to switch contest back to a normal mode.
	SuspendContest(context.Context, *SuspendContestInput) (*SuspendContestOutput, error)
	// Temporarily restrict submission function
	// Use ResumeContest to switch contest back to a normal mode.
	FreezeContest(context.Context, *FreezeContestInput) (*FreezeContestOutput, error)
	// Finalize contest results.
	// This action finalizes contest standings, issues participation certificates and awards medals etc.
	// After finalizing a contest it is not possible to change any parameters or results.
	// This action is irreversible.
	FinalizeContest(context.Context, *FinalizeContestInput) (*FinalizeContestOutput, error)
	// Re-start suspended or frozen contest
	ResumeContest(context.Context, *ResumeContestInput) (*ResumeContestOutput, error)
	WatchContest(*WatchContestInput, grpc.ServerStreamingServer[WatchContestOutput]) error
	ListActivities(context.Context, *ListActivitiesInput) (*ListActivitiesOutput, error)
	DescribeContestUsage(context.Context, *DescribeContestUsageInput) (*DescribeContestUsageOutput, error)
}

// UnimplementedContestServiceServer should be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedContestServiceServer struct{}

func (UnimplementedContestServiceServer) CreateContest(context.Context, *CreateContestInput) (*CreateContestOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateContest not implemented")
}
func (UnimplementedContestServiceServer) DeleteContest(context.Context, *DeleteContestInput) (*DeleteContestOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteContest not implemented")
}
func (UnimplementedContestServiceServer) UpdateContest(context.Context, *UpdateContestInput) (*UpdateContestOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateContest not implemented")
}
func (UnimplementedContestServiceServer) CopyContest(context.Context, *CopyContestInput) (*CopyContestOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CopyContest not implemented")
}
func (UnimplementedContestServiceServer) DescribeContest(context.Context, *DescribeContestInput) (*DescribeContestOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DescribeContest not implemented")
}
func (UnimplementedContestServiceServer) ListContests(context.Context, *ListContestsInput) (*ListContestsOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListContests not implemented")
}
func (UnimplementedContestServiceServer) OpenContest(context.Context, *OpenContestInput) (*OpenContestOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OpenContest not implemented")
}
func (UnimplementedContestServiceServer) CloseContest(context.Context, *CloseContestInput) (*CloseContestOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CloseContest not implemented")
}
func (UnimplementedContestServiceServer) SuspendContest(context.Context, *SuspendContestInput) (*SuspendContestOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SuspendContest not implemented")
}
func (UnimplementedContestServiceServer) FreezeContest(context.Context, *FreezeContestInput) (*FreezeContestOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FreezeContest not implemented")
}
func (UnimplementedContestServiceServer) FinalizeContest(context.Context, *FinalizeContestInput) (*FinalizeContestOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FinalizeContest not implemented")
}
func (UnimplementedContestServiceServer) ResumeContest(context.Context, *ResumeContestInput) (*ResumeContestOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ResumeContest not implemented")
}
func (UnimplementedContestServiceServer) WatchContest(*WatchContestInput, grpc.ServerStreamingServer[WatchContestOutput]) error {
	return status.Errorf(codes.Unimplemented, "method WatchContest not implemented")
}
func (UnimplementedContestServiceServer) ListActivities(context.Context, *ListActivitiesInput) (*ListActivitiesOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListActivities not implemented")
}
func (UnimplementedContestServiceServer) DescribeContestUsage(context.Context, *DescribeContestUsageInput) (*DescribeContestUsageOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DescribeContestUsage not implemented")
}
func (UnimplementedContestServiceServer) testEmbeddedByValue() {}

// UnsafeContestServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ContestServiceServer will
// result in compilation errors.
type UnsafeContestServiceServer interface {
	mustEmbedUnimplementedContestServiceServer()
}

func RegisterContestServiceServer(s grpc.ServiceRegistrar, srv ContestServiceServer) {
	// If the following call pancis, it indicates UnimplementedContestServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&ContestService_ServiceDesc, srv)
}

func _ContestService_CreateContest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateContestInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContestServiceServer).CreateContest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ContestService_CreateContest_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContestServiceServer).CreateContest(ctx, req.(*CreateContestInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContestService_DeleteContest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteContestInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContestServiceServer).DeleteContest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ContestService_DeleteContest_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContestServiceServer).DeleteContest(ctx, req.(*DeleteContestInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContestService_UpdateContest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateContestInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContestServiceServer).UpdateContest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ContestService_UpdateContest_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContestServiceServer).UpdateContest(ctx, req.(*UpdateContestInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContestService_CopyContest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CopyContestInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContestServiceServer).CopyContest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ContestService_CopyContest_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContestServiceServer).CopyContest(ctx, req.(*CopyContestInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContestService_DescribeContest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DescribeContestInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContestServiceServer).DescribeContest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ContestService_DescribeContest_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContestServiceServer).DescribeContest(ctx, req.(*DescribeContestInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContestService_ListContests_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListContestsInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContestServiceServer).ListContests(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ContestService_ListContests_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContestServiceServer).ListContests(ctx, req.(*ListContestsInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContestService_OpenContest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OpenContestInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContestServiceServer).OpenContest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ContestService_OpenContest_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContestServiceServer).OpenContest(ctx, req.(*OpenContestInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContestService_CloseContest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CloseContestInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContestServiceServer).CloseContest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ContestService_CloseContest_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContestServiceServer).CloseContest(ctx, req.(*CloseContestInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContestService_SuspendContest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SuspendContestInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContestServiceServer).SuspendContest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ContestService_SuspendContest_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContestServiceServer).SuspendContest(ctx, req.(*SuspendContestInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContestService_FreezeContest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FreezeContestInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContestServiceServer).FreezeContest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ContestService_FreezeContest_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContestServiceServer).FreezeContest(ctx, req.(*FreezeContestInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContestService_FinalizeContest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FinalizeContestInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContestServiceServer).FinalizeContest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ContestService_FinalizeContest_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContestServiceServer).FinalizeContest(ctx, req.(*FinalizeContestInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContestService_ResumeContest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResumeContestInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContestServiceServer).ResumeContest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ContestService_ResumeContest_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContestServiceServer).ResumeContest(ctx, req.(*ResumeContestInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContestService_WatchContest_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(WatchContestInput)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ContestServiceServer).WatchContest(m, &grpc.GenericServerStream[WatchContestInput, WatchContestOutput]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type ContestService_WatchContestServer = grpc.ServerStreamingServer[WatchContestOutput]

func _ContestService_ListActivities_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListActivitiesInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContestServiceServer).ListActivities(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ContestService_ListActivities_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContestServiceServer).ListActivities(ctx, req.(*ListActivitiesInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContestService_DescribeContestUsage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DescribeContestUsageInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContestServiceServer).DescribeContestUsage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ContestService_DescribeContestUsage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContestServiceServer).DescribeContestUsage(ctx, req.(*DescribeContestUsageInput))
	}
	return interceptor(ctx, in, info, handler)
}

// ContestService_ServiceDesc is the grpc.ServiceDesc for ContestService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ContestService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "eolymp.judge.ContestService",
	HandlerType: (*ContestServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateContest",
			Handler:    _ContestService_CreateContest_Handler,
		},
		{
			MethodName: "DeleteContest",
			Handler:    _ContestService_DeleteContest_Handler,
		},
		{
			MethodName: "UpdateContest",
			Handler:    _ContestService_UpdateContest_Handler,
		},
		{
			MethodName: "CopyContest",
			Handler:    _ContestService_CopyContest_Handler,
		},
		{
			MethodName: "DescribeContest",
			Handler:    _ContestService_DescribeContest_Handler,
		},
		{
			MethodName: "ListContests",
			Handler:    _ContestService_ListContests_Handler,
		},
		{
			MethodName: "OpenContest",
			Handler:    _ContestService_OpenContest_Handler,
		},
		{
			MethodName: "CloseContest",
			Handler:    _ContestService_CloseContest_Handler,
		},
		{
			MethodName: "SuspendContest",
			Handler:    _ContestService_SuspendContest_Handler,
		},
		{
			MethodName: "FreezeContest",
			Handler:    _ContestService_FreezeContest_Handler,
		},
		{
			MethodName: "FinalizeContest",
			Handler:    _ContestService_FinalizeContest_Handler,
		},
		{
			MethodName: "ResumeContest",
			Handler:    _ContestService_ResumeContest_Handler,
		},
		{
			MethodName: "ListActivities",
			Handler:    _ContestService_ListActivities_Handler,
		},
		{
			MethodName: "DescribeContestUsage",
			Handler:    _ContestService_DescribeContestUsage_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "WatchContest",
			Handler:       _ContestService_WatchContest_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "eolymp/judge/contest_service.proto",
}
