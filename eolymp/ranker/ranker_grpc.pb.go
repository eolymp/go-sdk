// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.4
// source: eolymp/ranker/ranker.proto

package ranker

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
	Ranker_CreateScoreboard_FullMethodName         = "/eolymp.ranker.Ranker/CreateScoreboard"
	Ranker_UpdateScoreboard_FullMethodName         = "/eolymp.ranker.Ranker/UpdateScoreboard"
	Ranker_RebuildScoreboard_FullMethodName        = "/eolymp.ranker.Ranker/RebuildScoreboard"
	Ranker_DeleteScoreboard_FullMethodName         = "/eolymp.ranker.Ranker/DeleteScoreboard"
	Ranker_DescribeScoreboard_FullMethodName       = "/eolymp.ranker.Ranker/DescribeScoreboard"
	Ranker_ListScoreboards_FullMethodName          = "/eolymp.ranker.Ranker/ListScoreboards"
	Ranker_DescribeScoreboardRow_FullMethodName    = "/eolymp.ranker.Ranker/DescribeScoreboardRow"
	Ranker_ListScoreboardRows_FullMethodName       = "/eolymp.ranker.Ranker/ListScoreboardRows"
	Ranker_AddScoreboardColumn_FullMethodName      = "/eolymp.ranker.Ranker/AddScoreboardColumn"
	Ranker_UpdateScoreboardColumn_FullMethodName   = "/eolymp.ranker.Ranker/UpdateScoreboardColumn"
	Ranker_DeleteScoreboardColumn_FullMethodName   = "/eolymp.ranker.Ranker/DeleteScoreboardColumn"
	Ranker_DescribeScoreboardColumn_FullMethodName = "/eolymp.ranker.Ranker/DescribeScoreboardColumn"
	Ranker_ListScoreboardColumns_FullMethodName    = "/eolymp.ranker.Ranker/ListScoreboardColumns"
	Ranker_ListActivities_FullMethodName           = "/eolymp.ranker.Ranker/ListActivities"
	Ranker_ScheduleAction_FullMethodName           = "/eolymp.ranker.Ranker/ScheduleAction"
	Ranker_UnscheduleAction_FullMethodName         = "/eolymp.ranker.Ranker/UnscheduleAction"
	Ranker_ListScheduledActions_FullMethodName     = "/eolymp.ranker.Ranker/ListScheduledActions"
)

// RankerClient is the client API for Ranker service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RankerClient interface {
	// Create a scoreboard
	CreateScoreboard(ctx context.Context, in *CreateScoreboardInput, opts ...grpc.CallOption) (*CreateScoreboardOutput, error)
	// Update existing scoreboard in a contest
	UpdateScoreboard(ctx context.Context, in *UpdateScoreboardInput, opts ...grpc.CallOption) (*UpdateScoreboardOutput, error)
	// Rebuild scoreboard: recalculate score for all participants
	RebuildScoreboard(ctx context.Context, in *RebuildScoreboardInput, opts ...grpc.CallOption) (*RebuildScoreboardOutput, error)
	// Delete scoreboard
	DeleteScoreboard(ctx context.Context, in *DeleteScoreboardInput, opts ...grpc.CallOption) (*DeleteScoreboardOutput, error)
	// Describe scoreboard
	DescribeScoreboard(ctx context.Context, in *DescribeScoreboardInput, opts ...grpc.CallOption) (*DescribeScoreboardOutput, error)
	// List scoreboards of a contest
	ListScoreboards(ctx context.Context, in *ListScoreboardsInput, opts ...grpc.CallOption) (*ListScoreboardsOutput, error)
	// List scoreboards of a contest
	DescribeScoreboardRow(ctx context.Context, in *DescribeScoreboardRowInput, opts ...grpc.CallOption) (*DescribeScoreboardRowOutput, error)
	// List scoreboards of a contest
	ListScoreboardRows(ctx context.Context, in *ListScoreboardRowsInput, opts ...grpc.CallOption) (*ListScoreboardRowsOutput, error)
	AddScoreboardColumn(ctx context.Context, in *AddScoreboardColumnInput, opts ...grpc.CallOption) (*AddScoreboardColumnOutput, error)
	UpdateScoreboardColumn(ctx context.Context, in *UpdateScoreboardColumnInput, opts ...grpc.CallOption) (*UpdateScoreboardColumnOutput, error)
	DeleteScoreboardColumn(ctx context.Context, in *DeleteScoreboardColumnInput, opts ...grpc.CallOption) (*DeleteScoreboardColumnOutput, error)
	DescribeScoreboardColumn(ctx context.Context, in *DescribeScoreboardColumnInput, opts ...grpc.CallOption) (*DescribeScoreboardColumnOutput, error)
	// List scoreboards of a contest
	ListScoreboardColumns(ctx context.Context, in *ListScoreboardColumnsInput, opts ...grpc.CallOption) (*ListScoreboardColumnsOutput, error)
	ListActivities(ctx context.Context, in *ListActivitiesInput, opts ...grpc.CallOption) (*ListActivitiesOutput, error)
	ScheduleAction(ctx context.Context, in *ScheduleActionInput, opts ...grpc.CallOption) (*ScheduleActionOutput, error)
	UnscheduleAction(ctx context.Context, in *UnscheduleActionInput, opts ...grpc.CallOption) (*UnscheduleActionOutput, error)
	ListScheduledActions(ctx context.Context, in *ListScheduledActionsInput, opts ...grpc.CallOption) (*ListScheduledActionsOutput, error)
}

type rankerClient struct {
	cc grpc.ClientConnInterface
}

func NewRankerClient(cc grpc.ClientConnInterface) RankerClient {
	return &rankerClient{cc}
}

func (c *rankerClient) CreateScoreboard(ctx context.Context, in *CreateScoreboardInput, opts ...grpc.CallOption) (*CreateScoreboardOutput, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateScoreboardOutput)
	err := c.cc.Invoke(ctx, Ranker_CreateScoreboard_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rankerClient) UpdateScoreboard(ctx context.Context, in *UpdateScoreboardInput, opts ...grpc.CallOption) (*UpdateScoreboardOutput, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateScoreboardOutput)
	err := c.cc.Invoke(ctx, Ranker_UpdateScoreboard_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rankerClient) RebuildScoreboard(ctx context.Context, in *RebuildScoreboardInput, opts ...grpc.CallOption) (*RebuildScoreboardOutput, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RebuildScoreboardOutput)
	err := c.cc.Invoke(ctx, Ranker_RebuildScoreboard_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rankerClient) DeleteScoreboard(ctx context.Context, in *DeleteScoreboardInput, opts ...grpc.CallOption) (*DeleteScoreboardOutput, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteScoreboardOutput)
	err := c.cc.Invoke(ctx, Ranker_DeleteScoreboard_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rankerClient) DescribeScoreboard(ctx context.Context, in *DescribeScoreboardInput, opts ...grpc.CallOption) (*DescribeScoreboardOutput, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DescribeScoreboardOutput)
	err := c.cc.Invoke(ctx, Ranker_DescribeScoreboard_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rankerClient) ListScoreboards(ctx context.Context, in *ListScoreboardsInput, opts ...grpc.CallOption) (*ListScoreboardsOutput, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListScoreboardsOutput)
	err := c.cc.Invoke(ctx, Ranker_ListScoreboards_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rankerClient) DescribeScoreboardRow(ctx context.Context, in *DescribeScoreboardRowInput, opts ...grpc.CallOption) (*DescribeScoreboardRowOutput, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DescribeScoreboardRowOutput)
	err := c.cc.Invoke(ctx, Ranker_DescribeScoreboardRow_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rankerClient) ListScoreboardRows(ctx context.Context, in *ListScoreboardRowsInput, opts ...grpc.CallOption) (*ListScoreboardRowsOutput, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListScoreboardRowsOutput)
	err := c.cc.Invoke(ctx, Ranker_ListScoreboardRows_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rankerClient) AddScoreboardColumn(ctx context.Context, in *AddScoreboardColumnInput, opts ...grpc.CallOption) (*AddScoreboardColumnOutput, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AddScoreboardColumnOutput)
	err := c.cc.Invoke(ctx, Ranker_AddScoreboardColumn_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rankerClient) UpdateScoreboardColumn(ctx context.Context, in *UpdateScoreboardColumnInput, opts ...grpc.CallOption) (*UpdateScoreboardColumnOutput, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateScoreboardColumnOutput)
	err := c.cc.Invoke(ctx, Ranker_UpdateScoreboardColumn_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rankerClient) DeleteScoreboardColumn(ctx context.Context, in *DeleteScoreboardColumnInput, opts ...grpc.CallOption) (*DeleteScoreboardColumnOutput, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteScoreboardColumnOutput)
	err := c.cc.Invoke(ctx, Ranker_DeleteScoreboardColumn_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rankerClient) DescribeScoreboardColumn(ctx context.Context, in *DescribeScoreboardColumnInput, opts ...grpc.CallOption) (*DescribeScoreboardColumnOutput, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DescribeScoreboardColumnOutput)
	err := c.cc.Invoke(ctx, Ranker_DescribeScoreboardColumn_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rankerClient) ListScoreboardColumns(ctx context.Context, in *ListScoreboardColumnsInput, opts ...grpc.CallOption) (*ListScoreboardColumnsOutput, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListScoreboardColumnsOutput)
	err := c.cc.Invoke(ctx, Ranker_ListScoreboardColumns_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rankerClient) ListActivities(ctx context.Context, in *ListActivitiesInput, opts ...grpc.CallOption) (*ListActivitiesOutput, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListActivitiesOutput)
	err := c.cc.Invoke(ctx, Ranker_ListActivities_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rankerClient) ScheduleAction(ctx context.Context, in *ScheduleActionInput, opts ...grpc.CallOption) (*ScheduleActionOutput, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ScheduleActionOutput)
	err := c.cc.Invoke(ctx, Ranker_ScheduleAction_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rankerClient) UnscheduleAction(ctx context.Context, in *UnscheduleActionInput, opts ...grpc.CallOption) (*UnscheduleActionOutput, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UnscheduleActionOutput)
	err := c.cc.Invoke(ctx, Ranker_UnscheduleAction_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rankerClient) ListScheduledActions(ctx context.Context, in *ListScheduledActionsInput, opts ...grpc.CallOption) (*ListScheduledActionsOutput, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListScheduledActionsOutput)
	err := c.cc.Invoke(ctx, Ranker_ListScheduledActions_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RankerServer is the server API for Ranker service.
// All implementations should embed UnimplementedRankerServer
// for forward compatibility.
type RankerServer interface {
	// Create a scoreboard
	CreateScoreboard(context.Context, *CreateScoreboardInput) (*CreateScoreboardOutput, error)
	// Update existing scoreboard in a contest
	UpdateScoreboard(context.Context, *UpdateScoreboardInput) (*UpdateScoreboardOutput, error)
	// Rebuild scoreboard: recalculate score for all participants
	RebuildScoreboard(context.Context, *RebuildScoreboardInput) (*RebuildScoreboardOutput, error)
	// Delete scoreboard
	DeleteScoreboard(context.Context, *DeleteScoreboardInput) (*DeleteScoreboardOutput, error)
	// Describe scoreboard
	DescribeScoreboard(context.Context, *DescribeScoreboardInput) (*DescribeScoreboardOutput, error)
	// List scoreboards of a contest
	ListScoreboards(context.Context, *ListScoreboardsInput) (*ListScoreboardsOutput, error)
	// List scoreboards of a contest
	DescribeScoreboardRow(context.Context, *DescribeScoreboardRowInput) (*DescribeScoreboardRowOutput, error)
	// List scoreboards of a contest
	ListScoreboardRows(context.Context, *ListScoreboardRowsInput) (*ListScoreboardRowsOutput, error)
	AddScoreboardColumn(context.Context, *AddScoreboardColumnInput) (*AddScoreboardColumnOutput, error)
	UpdateScoreboardColumn(context.Context, *UpdateScoreboardColumnInput) (*UpdateScoreboardColumnOutput, error)
	DeleteScoreboardColumn(context.Context, *DeleteScoreboardColumnInput) (*DeleteScoreboardColumnOutput, error)
	DescribeScoreboardColumn(context.Context, *DescribeScoreboardColumnInput) (*DescribeScoreboardColumnOutput, error)
	// List scoreboards of a contest
	ListScoreboardColumns(context.Context, *ListScoreboardColumnsInput) (*ListScoreboardColumnsOutput, error)
	ListActivities(context.Context, *ListActivitiesInput) (*ListActivitiesOutput, error)
	ScheduleAction(context.Context, *ScheduleActionInput) (*ScheduleActionOutput, error)
	UnscheduleAction(context.Context, *UnscheduleActionInput) (*UnscheduleActionOutput, error)
	ListScheduledActions(context.Context, *ListScheduledActionsInput) (*ListScheduledActionsOutput, error)
}

// UnimplementedRankerServer should be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedRankerServer struct{}

func (UnimplementedRankerServer) CreateScoreboard(context.Context, *CreateScoreboardInput) (*CreateScoreboardOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateScoreboard not implemented")
}
func (UnimplementedRankerServer) UpdateScoreboard(context.Context, *UpdateScoreboardInput) (*UpdateScoreboardOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateScoreboard not implemented")
}
func (UnimplementedRankerServer) RebuildScoreboard(context.Context, *RebuildScoreboardInput) (*RebuildScoreboardOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RebuildScoreboard not implemented")
}
func (UnimplementedRankerServer) DeleteScoreboard(context.Context, *DeleteScoreboardInput) (*DeleteScoreboardOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteScoreboard not implemented")
}
func (UnimplementedRankerServer) DescribeScoreboard(context.Context, *DescribeScoreboardInput) (*DescribeScoreboardOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DescribeScoreboard not implemented")
}
func (UnimplementedRankerServer) ListScoreboards(context.Context, *ListScoreboardsInput) (*ListScoreboardsOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListScoreboards not implemented")
}
func (UnimplementedRankerServer) DescribeScoreboardRow(context.Context, *DescribeScoreboardRowInput) (*DescribeScoreboardRowOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DescribeScoreboardRow not implemented")
}
func (UnimplementedRankerServer) ListScoreboardRows(context.Context, *ListScoreboardRowsInput) (*ListScoreboardRowsOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListScoreboardRows not implemented")
}
func (UnimplementedRankerServer) AddScoreboardColumn(context.Context, *AddScoreboardColumnInput) (*AddScoreboardColumnOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddScoreboardColumn not implemented")
}
func (UnimplementedRankerServer) UpdateScoreboardColumn(context.Context, *UpdateScoreboardColumnInput) (*UpdateScoreboardColumnOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateScoreboardColumn not implemented")
}
func (UnimplementedRankerServer) DeleteScoreboardColumn(context.Context, *DeleteScoreboardColumnInput) (*DeleteScoreboardColumnOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteScoreboardColumn not implemented")
}
func (UnimplementedRankerServer) DescribeScoreboardColumn(context.Context, *DescribeScoreboardColumnInput) (*DescribeScoreboardColumnOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DescribeScoreboardColumn not implemented")
}
func (UnimplementedRankerServer) ListScoreboardColumns(context.Context, *ListScoreboardColumnsInput) (*ListScoreboardColumnsOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListScoreboardColumns not implemented")
}
func (UnimplementedRankerServer) ListActivities(context.Context, *ListActivitiesInput) (*ListActivitiesOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListActivities not implemented")
}
func (UnimplementedRankerServer) ScheduleAction(context.Context, *ScheduleActionInput) (*ScheduleActionOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ScheduleAction not implemented")
}
func (UnimplementedRankerServer) UnscheduleAction(context.Context, *UnscheduleActionInput) (*UnscheduleActionOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnscheduleAction not implemented")
}
func (UnimplementedRankerServer) ListScheduledActions(context.Context, *ListScheduledActionsInput) (*ListScheduledActionsOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListScheduledActions not implemented")
}
func (UnimplementedRankerServer) testEmbeddedByValue() {}

// UnsafeRankerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RankerServer will
// result in compilation errors.
type UnsafeRankerServer interface {
	mustEmbedUnimplementedRankerServer()
}

func RegisterRankerServer(s grpc.ServiceRegistrar, srv RankerServer) {
	// If the following call pancis, it indicates UnimplementedRankerServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Ranker_ServiceDesc, srv)
}

func _Ranker_CreateScoreboard_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateScoreboardInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RankerServer).CreateScoreboard(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Ranker_CreateScoreboard_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RankerServer).CreateScoreboard(ctx, req.(*CreateScoreboardInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _Ranker_UpdateScoreboard_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateScoreboardInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RankerServer).UpdateScoreboard(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Ranker_UpdateScoreboard_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RankerServer).UpdateScoreboard(ctx, req.(*UpdateScoreboardInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _Ranker_RebuildScoreboard_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RebuildScoreboardInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RankerServer).RebuildScoreboard(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Ranker_RebuildScoreboard_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RankerServer).RebuildScoreboard(ctx, req.(*RebuildScoreboardInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _Ranker_DeleteScoreboard_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteScoreboardInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RankerServer).DeleteScoreboard(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Ranker_DeleteScoreboard_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RankerServer).DeleteScoreboard(ctx, req.(*DeleteScoreboardInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _Ranker_DescribeScoreboard_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DescribeScoreboardInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RankerServer).DescribeScoreboard(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Ranker_DescribeScoreboard_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RankerServer).DescribeScoreboard(ctx, req.(*DescribeScoreboardInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _Ranker_ListScoreboards_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListScoreboardsInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RankerServer).ListScoreboards(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Ranker_ListScoreboards_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RankerServer).ListScoreboards(ctx, req.(*ListScoreboardsInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _Ranker_DescribeScoreboardRow_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DescribeScoreboardRowInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RankerServer).DescribeScoreboardRow(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Ranker_DescribeScoreboardRow_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RankerServer).DescribeScoreboardRow(ctx, req.(*DescribeScoreboardRowInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _Ranker_ListScoreboardRows_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListScoreboardRowsInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RankerServer).ListScoreboardRows(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Ranker_ListScoreboardRows_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RankerServer).ListScoreboardRows(ctx, req.(*ListScoreboardRowsInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _Ranker_AddScoreboardColumn_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddScoreboardColumnInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RankerServer).AddScoreboardColumn(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Ranker_AddScoreboardColumn_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RankerServer).AddScoreboardColumn(ctx, req.(*AddScoreboardColumnInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _Ranker_UpdateScoreboardColumn_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateScoreboardColumnInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RankerServer).UpdateScoreboardColumn(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Ranker_UpdateScoreboardColumn_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RankerServer).UpdateScoreboardColumn(ctx, req.(*UpdateScoreboardColumnInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _Ranker_DeleteScoreboardColumn_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteScoreboardColumnInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RankerServer).DeleteScoreboardColumn(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Ranker_DeleteScoreboardColumn_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RankerServer).DeleteScoreboardColumn(ctx, req.(*DeleteScoreboardColumnInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _Ranker_DescribeScoreboardColumn_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DescribeScoreboardColumnInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RankerServer).DescribeScoreboardColumn(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Ranker_DescribeScoreboardColumn_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RankerServer).DescribeScoreboardColumn(ctx, req.(*DescribeScoreboardColumnInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _Ranker_ListScoreboardColumns_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListScoreboardColumnsInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RankerServer).ListScoreboardColumns(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Ranker_ListScoreboardColumns_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RankerServer).ListScoreboardColumns(ctx, req.(*ListScoreboardColumnsInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _Ranker_ListActivities_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListActivitiesInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RankerServer).ListActivities(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Ranker_ListActivities_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RankerServer).ListActivities(ctx, req.(*ListActivitiesInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _Ranker_ScheduleAction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ScheduleActionInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RankerServer).ScheduleAction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Ranker_ScheduleAction_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RankerServer).ScheduleAction(ctx, req.(*ScheduleActionInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _Ranker_UnscheduleAction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UnscheduleActionInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RankerServer).UnscheduleAction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Ranker_UnscheduleAction_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RankerServer).UnscheduleAction(ctx, req.(*UnscheduleActionInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _Ranker_ListScheduledActions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListScheduledActionsInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RankerServer).ListScheduledActions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Ranker_ListScheduledActions_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RankerServer).ListScheduledActions(ctx, req.(*ListScheduledActionsInput))
	}
	return interceptor(ctx, in, info, handler)
}

// Ranker_ServiceDesc is the grpc.ServiceDesc for Ranker service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Ranker_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "eolymp.ranker.Ranker",
	HandlerType: (*RankerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateScoreboard",
			Handler:    _Ranker_CreateScoreboard_Handler,
		},
		{
			MethodName: "UpdateScoreboard",
			Handler:    _Ranker_UpdateScoreboard_Handler,
		},
		{
			MethodName: "RebuildScoreboard",
			Handler:    _Ranker_RebuildScoreboard_Handler,
		},
		{
			MethodName: "DeleteScoreboard",
			Handler:    _Ranker_DeleteScoreboard_Handler,
		},
		{
			MethodName: "DescribeScoreboard",
			Handler:    _Ranker_DescribeScoreboard_Handler,
		},
		{
			MethodName: "ListScoreboards",
			Handler:    _Ranker_ListScoreboards_Handler,
		},
		{
			MethodName: "DescribeScoreboardRow",
			Handler:    _Ranker_DescribeScoreboardRow_Handler,
		},
		{
			MethodName: "ListScoreboardRows",
			Handler:    _Ranker_ListScoreboardRows_Handler,
		},
		{
			MethodName: "AddScoreboardColumn",
			Handler:    _Ranker_AddScoreboardColumn_Handler,
		},
		{
			MethodName: "UpdateScoreboardColumn",
			Handler:    _Ranker_UpdateScoreboardColumn_Handler,
		},
		{
			MethodName: "DeleteScoreboardColumn",
			Handler:    _Ranker_DeleteScoreboardColumn_Handler,
		},
		{
			MethodName: "DescribeScoreboardColumn",
			Handler:    _Ranker_DescribeScoreboardColumn_Handler,
		},
		{
			MethodName: "ListScoreboardColumns",
			Handler:    _Ranker_ListScoreboardColumns_Handler,
		},
		{
			MethodName: "ListActivities",
			Handler:    _Ranker_ListActivities_Handler,
		},
		{
			MethodName: "ScheduleAction",
			Handler:    _Ranker_ScheduleAction_Handler,
		},
		{
			MethodName: "UnscheduleAction",
			Handler:    _Ranker_UnscheduleAction_Handler,
		},
		{
			MethodName: "ListScheduledActions",
			Handler:    _Ranker_ListScheduledActions_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "eolymp/ranker/ranker.proto",
}
