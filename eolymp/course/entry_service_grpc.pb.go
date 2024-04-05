// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.18.1
// source: eolymp/course/entry_service.proto

package course

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
	EntryService_CreateEntry_FullMethodName      = "/eolymp.course.EntryService/CreateEntry"
	EntryService_UpdateEntry_FullMethodName      = "/eolymp.course.EntryService/UpdateEntry"
	EntryService_RenameEntry_FullMethodName      = "/eolymp.course.EntryService/RenameEntry"
	EntryService_MoveEntry_FullMethodName        = "/eolymp.course.EntryService/MoveEntry"
	EntryService_DeleteEntry_FullMethodName      = "/eolymp.course.EntryService/DeleteEntry"
	EntryService_DescribeEntry_FullMethodName    = "/eolymp.course.EntryService/DescribeEntry"
	EntryService_ListEntries_FullMethodName      = "/eolymp.course.EntryService/ListEntries"
	EntryService_DescribeTOC_FullMethodName      = "/eolymp.course.EntryService/DescribeTOC"
	EntryService_ListParents_FullMethodName      = "/eolymp.course.EntryService/ListParents"
	EntryService_DescribeProgress_FullMethodName = "/eolymp.course.EntryService/DescribeProgress"
	EntryService_ReportProgress_FullMethodName   = "/eolymp.course.EntryService/ReportProgress"
	EntryService_WatchProgress_FullMethodName    = "/eolymp.course.EntryService/WatchProgress"
)

// EntryServiceClient is the client API for EntryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EntryServiceClient interface {
	CreateEntry(ctx context.Context, in *CreateEntryInput, opts ...grpc.CallOption) (*CreateEntryOutput, error)
	UpdateEntry(ctx context.Context, in *UpdateEntryInput, opts ...grpc.CallOption) (*UpdateEntryOutput, error)
	RenameEntry(ctx context.Context, in *RenameEntryInput, opts ...grpc.CallOption) (*RenameEntryOutput, error)
	MoveEntry(ctx context.Context, in *MoveEntryInput, opts ...grpc.CallOption) (*MoveEntryOutput, error)
	DeleteEntry(ctx context.Context, in *DeleteEntryInput, opts ...grpc.CallOption) (*DeleteEntryOutput, error)
	DescribeEntry(ctx context.Context, in *DescribeEntryInput, opts ...grpc.CallOption) (*DescribeEntryOutput, error)
	ListEntries(ctx context.Context, in *ListEntriesInput, opts ...grpc.CallOption) (*ListEntriesOutput, error)
	DescribeTOC(ctx context.Context, in *DescribeTOCInput, opts ...grpc.CallOption) (*DescribeTOCOutput, error)
	ListParents(ctx context.Context, in *ListParentsInput, opts ...grpc.CallOption) (*ListParentsOutput, error)
	DescribeProgress(ctx context.Context, in *DescribeProgressInput, opts ...grpc.CallOption) (*DescribeProgressOutput, error)
	ReportProgress(ctx context.Context, in *ReportProgressInput, opts ...grpc.CallOption) (*ReportProgressOutput, error)
	WatchProgress(ctx context.Context, in *WatchProgressInput, opts ...grpc.CallOption) (EntryService_WatchProgressClient, error)
}

type entryServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewEntryServiceClient(cc grpc.ClientConnInterface) EntryServiceClient {
	return &entryServiceClient{cc}
}

func (c *entryServiceClient) CreateEntry(ctx context.Context, in *CreateEntryInput, opts ...grpc.CallOption) (*CreateEntryOutput, error) {
	out := new(CreateEntryOutput)
	err := c.cc.Invoke(ctx, EntryService_CreateEntry_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *entryServiceClient) UpdateEntry(ctx context.Context, in *UpdateEntryInput, opts ...grpc.CallOption) (*UpdateEntryOutput, error) {
	out := new(UpdateEntryOutput)
	err := c.cc.Invoke(ctx, EntryService_UpdateEntry_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *entryServiceClient) RenameEntry(ctx context.Context, in *RenameEntryInput, opts ...grpc.CallOption) (*RenameEntryOutput, error) {
	out := new(RenameEntryOutput)
	err := c.cc.Invoke(ctx, EntryService_RenameEntry_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *entryServiceClient) MoveEntry(ctx context.Context, in *MoveEntryInput, opts ...grpc.CallOption) (*MoveEntryOutput, error) {
	out := new(MoveEntryOutput)
	err := c.cc.Invoke(ctx, EntryService_MoveEntry_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *entryServiceClient) DeleteEntry(ctx context.Context, in *DeleteEntryInput, opts ...grpc.CallOption) (*DeleteEntryOutput, error) {
	out := new(DeleteEntryOutput)
	err := c.cc.Invoke(ctx, EntryService_DeleteEntry_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *entryServiceClient) DescribeEntry(ctx context.Context, in *DescribeEntryInput, opts ...grpc.CallOption) (*DescribeEntryOutput, error) {
	out := new(DescribeEntryOutput)
	err := c.cc.Invoke(ctx, EntryService_DescribeEntry_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *entryServiceClient) ListEntries(ctx context.Context, in *ListEntriesInput, opts ...grpc.CallOption) (*ListEntriesOutput, error) {
	out := new(ListEntriesOutput)
	err := c.cc.Invoke(ctx, EntryService_ListEntries_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *entryServiceClient) DescribeTOC(ctx context.Context, in *DescribeTOCInput, opts ...grpc.CallOption) (*DescribeTOCOutput, error) {
	out := new(DescribeTOCOutput)
	err := c.cc.Invoke(ctx, EntryService_DescribeTOC_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *entryServiceClient) ListParents(ctx context.Context, in *ListParentsInput, opts ...grpc.CallOption) (*ListParentsOutput, error) {
	out := new(ListParentsOutput)
	err := c.cc.Invoke(ctx, EntryService_ListParents_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *entryServiceClient) DescribeProgress(ctx context.Context, in *DescribeProgressInput, opts ...grpc.CallOption) (*DescribeProgressOutput, error) {
	out := new(DescribeProgressOutput)
	err := c.cc.Invoke(ctx, EntryService_DescribeProgress_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *entryServiceClient) ReportProgress(ctx context.Context, in *ReportProgressInput, opts ...grpc.CallOption) (*ReportProgressOutput, error) {
	out := new(ReportProgressOutput)
	err := c.cc.Invoke(ctx, EntryService_ReportProgress_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *entryServiceClient) WatchProgress(ctx context.Context, in *WatchProgressInput, opts ...grpc.CallOption) (EntryService_WatchProgressClient, error) {
	stream, err := c.cc.NewStream(ctx, &EntryService_ServiceDesc.Streams[0], EntryService_WatchProgress_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &entryServiceWatchProgressClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type EntryService_WatchProgressClient interface {
	Recv() (*WatchProgressOutput, error)
	grpc.ClientStream
}

type entryServiceWatchProgressClient struct {
	grpc.ClientStream
}

func (x *entryServiceWatchProgressClient) Recv() (*WatchProgressOutput, error) {
	m := new(WatchProgressOutput)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// EntryServiceServer is the server API for EntryService service.
// All implementations should embed UnimplementedEntryServiceServer
// for forward compatibility
type EntryServiceServer interface {
	CreateEntry(context.Context, *CreateEntryInput) (*CreateEntryOutput, error)
	UpdateEntry(context.Context, *UpdateEntryInput) (*UpdateEntryOutput, error)
	RenameEntry(context.Context, *RenameEntryInput) (*RenameEntryOutput, error)
	MoveEntry(context.Context, *MoveEntryInput) (*MoveEntryOutput, error)
	DeleteEntry(context.Context, *DeleteEntryInput) (*DeleteEntryOutput, error)
	DescribeEntry(context.Context, *DescribeEntryInput) (*DescribeEntryOutput, error)
	ListEntries(context.Context, *ListEntriesInput) (*ListEntriesOutput, error)
	DescribeTOC(context.Context, *DescribeTOCInput) (*DescribeTOCOutput, error)
	ListParents(context.Context, *ListParentsInput) (*ListParentsOutput, error)
	DescribeProgress(context.Context, *DescribeProgressInput) (*DescribeProgressOutput, error)
	ReportProgress(context.Context, *ReportProgressInput) (*ReportProgressOutput, error)
	WatchProgress(*WatchProgressInput, EntryService_WatchProgressServer) error
}

// UnimplementedEntryServiceServer should be embedded to have forward compatible implementations.
type UnimplementedEntryServiceServer struct {
}

func (UnimplementedEntryServiceServer) CreateEntry(context.Context, *CreateEntryInput) (*CreateEntryOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateEntry not implemented")
}
func (UnimplementedEntryServiceServer) UpdateEntry(context.Context, *UpdateEntryInput) (*UpdateEntryOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateEntry not implemented")
}
func (UnimplementedEntryServiceServer) RenameEntry(context.Context, *RenameEntryInput) (*RenameEntryOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RenameEntry not implemented")
}
func (UnimplementedEntryServiceServer) MoveEntry(context.Context, *MoveEntryInput) (*MoveEntryOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MoveEntry not implemented")
}
func (UnimplementedEntryServiceServer) DeleteEntry(context.Context, *DeleteEntryInput) (*DeleteEntryOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteEntry not implemented")
}
func (UnimplementedEntryServiceServer) DescribeEntry(context.Context, *DescribeEntryInput) (*DescribeEntryOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DescribeEntry not implemented")
}
func (UnimplementedEntryServiceServer) ListEntries(context.Context, *ListEntriesInput) (*ListEntriesOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListEntries not implemented")
}
func (UnimplementedEntryServiceServer) DescribeTOC(context.Context, *DescribeTOCInput) (*DescribeTOCOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DescribeTOC not implemented")
}
func (UnimplementedEntryServiceServer) ListParents(context.Context, *ListParentsInput) (*ListParentsOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListParents not implemented")
}
func (UnimplementedEntryServiceServer) DescribeProgress(context.Context, *DescribeProgressInput) (*DescribeProgressOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DescribeProgress not implemented")
}
func (UnimplementedEntryServiceServer) ReportProgress(context.Context, *ReportProgressInput) (*ReportProgressOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReportProgress not implemented")
}
func (UnimplementedEntryServiceServer) WatchProgress(*WatchProgressInput, EntryService_WatchProgressServer) error {
	return status.Errorf(codes.Unimplemented, "method WatchProgress not implemented")
}

// UnsafeEntryServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EntryServiceServer will
// result in compilation errors.
type UnsafeEntryServiceServer interface {
	mustEmbedUnimplementedEntryServiceServer()
}

func RegisterEntryServiceServer(s grpc.ServiceRegistrar, srv EntryServiceServer) {
	s.RegisterService(&EntryService_ServiceDesc, srv)
}

func _EntryService_CreateEntry_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateEntryInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EntryServiceServer).CreateEntry(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EntryService_CreateEntry_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EntryServiceServer).CreateEntry(ctx, req.(*CreateEntryInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _EntryService_UpdateEntry_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateEntryInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EntryServiceServer).UpdateEntry(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EntryService_UpdateEntry_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EntryServiceServer).UpdateEntry(ctx, req.(*UpdateEntryInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _EntryService_RenameEntry_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RenameEntryInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EntryServiceServer).RenameEntry(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EntryService_RenameEntry_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EntryServiceServer).RenameEntry(ctx, req.(*RenameEntryInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _EntryService_MoveEntry_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MoveEntryInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EntryServiceServer).MoveEntry(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EntryService_MoveEntry_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EntryServiceServer).MoveEntry(ctx, req.(*MoveEntryInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _EntryService_DeleteEntry_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteEntryInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EntryServiceServer).DeleteEntry(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EntryService_DeleteEntry_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EntryServiceServer).DeleteEntry(ctx, req.(*DeleteEntryInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _EntryService_DescribeEntry_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DescribeEntryInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EntryServiceServer).DescribeEntry(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EntryService_DescribeEntry_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EntryServiceServer).DescribeEntry(ctx, req.(*DescribeEntryInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _EntryService_ListEntries_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListEntriesInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EntryServiceServer).ListEntries(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EntryService_ListEntries_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EntryServiceServer).ListEntries(ctx, req.(*ListEntriesInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _EntryService_DescribeTOC_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DescribeTOCInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EntryServiceServer).DescribeTOC(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EntryService_DescribeTOC_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EntryServiceServer).DescribeTOC(ctx, req.(*DescribeTOCInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _EntryService_ListParents_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListParentsInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EntryServiceServer).ListParents(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EntryService_ListParents_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EntryServiceServer).ListParents(ctx, req.(*ListParentsInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _EntryService_DescribeProgress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DescribeProgressInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EntryServiceServer).DescribeProgress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EntryService_DescribeProgress_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EntryServiceServer).DescribeProgress(ctx, req.(*DescribeProgressInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _EntryService_ReportProgress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReportProgressInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EntryServiceServer).ReportProgress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EntryService_ReportProgress_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EntryServiceServer).ReportProgress(ctx, req.(*ReportProgressInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _EntryService_WatchProgress_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(WatchProgressInput)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(EntryServiceServer).WatchProgress(m, &entryServiceWatchProgressServer{stream})
}

type EntryService_WatchProgressServer interface {
	Send(*WatchProgressOutput) error
	grpc.ServerStream
}

type entryServiceWatchProgressServer struct {
	grpc.ServerStream
}

func (x *entryServiceWatchProgressServer) Send(m *WatchProgressOutput) error {
	return x.ServerStream.SendMsg(m)
}

// EntryService_ServiceDesc is the grpc.ServiceDesc for EntryService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var EntryService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "eolymp.course.EntryService",
	HandlerType: (*EntryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateEntry",
			Handler:    _EntryService_CreateEntry_Handler,
		},
		{
			MethodName: "UpdateEntry",
			Handler:    _EntryService_UpdateEntry_Handler,
		},
		{
			MethodName: "RenameEntry",
			Handler:    _EntryService_RenameEntry_Handler,
		},
		{
			MethodName: "MoveEntry",
			Handler:    _EntryService_MoveEntry_Handler,
		},
		{
			MethodName: "DeleteEntry",
			Handler:    _EntryService_DeleteEntry_Handler,
		},
		{
			MethodName: "DescribeEntry",
			Handler:    _EntryService_DescribeEntry_Handler,
		},
		{
			MethodName: "ListEntries",
			Handler:    _EntryService_ListEntries_Handler,
		},
		{
			MethodName: "DescribeTOC",
			Handler:    _EntryService_DescribeTOC_Handler,
		},
		{
			MethodName: "ListParents",
			Handler:    _EntryService_ListParents_Handler,
		},
		{
			MethodName: "DescribeProgress",
			Handler:    _EntryService_DescribeProgress_Handler,
		},
		{
			MethodName: "ReportProgress",
			Handler:    _EntryService_ReportProgress_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "WatchProgress",
			Handler:       _EntryService_WatchProgress_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "eolymp/course/entry_service.proto",
}
