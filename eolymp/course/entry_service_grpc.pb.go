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
	EntryService_CreateSection_FullMethodName  = "/eolymp.course.EntryService/CreateSection"
	EntryService_UpdateSection_FullMethodName  = "/eolymp.course.EntryService/UpdateSection"
	EntryService_CreateDocument_FullMethodName = "/eolymp.course.EntryService/CreateDocument"
	EntryService_UpdateDocument_FullMethodName = "/eolymp.course.EntryService/UpdateDocument"
	EntryService_RenameEntry_FullMethodName    = "/eolymp.course.EntryService/RenameEntry"
	EntryService_DeleteEntry_FullMethodName    = "/eolymp.course.EntryService/DeleteEntry"
	EntryService_DescribeEntry_FullMethodName  = "/eolymp.course.EntryService/DescribeEntry"
	EntryService_ListEntries_FullMethodName    = "/eolymp.course.EntryService/ListEntries"
)

// EntryServiceClient is the client API for EntryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EntryServiceClient interface {
	CreateSection(ctx context.Context, in *CreateSectionInput, opts ...grpc.CallOption) (*CreateSectionOutput, error)
	UpdateSection(ctx context.Context, in *UpdateSectionInput, opts ...grpc.CallOption) (*UpdateSectionOutput, error)
	CreateDocument(ctx context.Context, in *CreateDocumentInput, opts ...grpc.CallOption) (*CreateDocumentOutput, error)
	UpdateDocument(ctx context.Context, in *UpdateDocumentInput, opts ...grpc.CallOption) (*UpdateDocumentOutput, error)
	RenameEntry(ctx context.Context, in *RenameEntryInput, opts ...grpc.CallOption) (*RenameEntryOutput, error)
	DeleteEntry(ctx context.Context, in *DeleteEntryInput, opts ...grpc.CallOption) (*DeleteEntryOutput, error)
	DescribeEntry(ctx context.Context, in *DescribeEntryInput, opts ...grpc.CallOption) (*DescribeEntryOutput, error)
	ListEntries(ctx context.Context, in *ListEntriesInput, opts ...grpc.CallOption) (*ListEntriesOutput, error)
}

type entryServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewEntryServiceClient(cc grpc.ClientConnInterface) EntryServiceClient {
	return &entryServiceClient{cc}
}

func (c *entryServiceClient) CreateSection(ctx context.Context, in *CreateSectionInput, opts ...grpc.CallOption) (*CreateSectionOutput, error) {
	out := new(CreateSectionOutput)
	err := c.cc.Invoke(ctx, EntryService_CreateSection_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *entryServiceClient) UpdateSection(ctx context.Context, in *UpdateSectionInput, opts ...grpc.CallOption) (*UpdateSectionOutput, error) {
	out := new(UpdateSectionOutput)
	err := c.cc.Invoke(ctx, EntryService_UpdateSection_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *entryServiceClient) CreateDocument(ctx context.Context, in *CreateDocumentInput, opts ...grpc.CallOption) (*CreateDocumentOutput, error) {
	out := new(CreateDocumentOutput)
	err := c.cc.Invoke(ctx, EntryService_CreateDocument_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *entryServiceClient) UpdateDocument(ctx context.Context, in *UpdateDocumentInput, opts ...grpc.CallOption) (*UpdateDocumentOutput, error) {
	out := new(UpdateDocumentOutput)
	err := c.cc.Invoke(ctx, EntryService_UpdateDocument_FullMethodName, in, out, opts...)
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

// EntryServiceServer is the server API for EntryService service.
// All implementations should embed UnimplementedEntryServiceServer
// for forward compatibility
type EntryServiceServer interface {
	CreateSection(context.Context, *CreateSectionInput) (*CreateSectionOutput, error)
	UpdateSection(context.Context, *UpdateSectionInput) (*UpdateSectionOutput, error)
	CreateDocument(context.Context, *CreateDocumentInput) (*CreateDocumentOutput, error)
	UpdateDocument(context.Context, *UpdateDocumentInput) (*UpdateDocumentOutput, error)
	RenameEntry(context.Context, *RenameEntryInput) (*RenameEntryOutput, error)
	DeleteEntry(context.Context, *DeleteEntryInput) (*DeleteEntryOutput, error)
	DescribeEntry(context.Context, *DescribeEntryInput) (*DescribeEntryOutput, error)
	ListEntries(context.Context, *ListEntriesInput) (*ListEntriesOutput, error)
}

// UnimplementedEntryServiceServer should be embedded to have forward compatible implementations.
type UnimplementedEntryServiceServer struct {
}

func (UnimplementedEntryServiceServer) CreateSection(context.Context, *CreateSectionInput) (*CreateSectionOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateSection not implemented")
}
func (UnimplementedEntryServiceServer) UpdateSection(context.Context, *UpdateSectionInput) (*UpdateSectionOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateSection not implemented")
}
func (UnimplementedEntryServiceServer) CreateDocument(context.Context, *CreateDocumentInput) (*CreateDocumentOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateDocument not implemented")
}
func (UnimplementedEntryServiceServer) UpdateDocument(context.Context, *UpdateDocumentInput) (*UpdateDocumentOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateDocument not implemented")
}
func (UnimplementedEntryServiceServer) RenameEntry(context.Context, *RenameEntryInput) (*RenameEntryOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RenameEntry not implemented")
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

// UnsafeEntryServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EntryServiceServer will
// result in compilation errors.
type UnsafeEntryServiceServer interface {
	mustEmbedUnimplementedEntryServiceServer()
}

func RegisterEntryServiceServer(s grpc.ServiceRegistrar, srv EntryServiceServer) {
	s.RegisterService(&EntryService_ServiceDesc, srv)
}

func _EntryService_CreateSection_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateSectionInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EntryServiceServer).CreateSection(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EntryService_CreateSection_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EntryServiceServer).CreateSection(ctx, req.(*CreateSectionInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _EntryService_UpdateSection_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateSectionInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EntryServiceServer).UpdateSection(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EntryService_UpdateSection_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EntryServiceServer).UpdateSection(ctx, req.(*UpdateSectionInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _EntryService_CreateDocument_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateDocumentInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EntryServiceServer).CreateDocument(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EntryService_CreateDocument_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EntryServiceServer).CreateDocument(ctx, req.(*CreateDocumentInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _EntryService_UpdateDocument_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateDocumentInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EntryServiceServer).UpdateDocument(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EntryService_UpdateDocument_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EntryServiceServer).UpdateDocument(ctx, req.(*UpdateDocumentInput))
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

// EntryService_ServiceDesc is the grpc.ServiceDesc for EntryService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var EntryService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "eolymp.course.EntryService",
	HandlerType: (*EntryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateSection",
			Handler:    _EntryService_CreateSection_Handler,
		},
		{
			MethodName: "UpdateSection",
			Handler:    _EntryService_UpdateSection_Handler,
		},
		{
			MethodName: "CreateDocument",
			Handler:    _EntryService_CreateDocument_Handler,
		},
		{
			MethodName: "UpdateDocument",
			Handler:    _EntryService_UpdateDocument_Handler,
		},
		{
			MethodName: "RenameEntry",
			Handler:    _EntryService_RenameEntry_Handler,
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
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "eolymp/course/entry_service.proto",
}
