// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.18.1
// source: eolymp/judge/announcement_service.proto

package judge

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
	AnnouncementService_CreateAnnouncement_FullMethodName         = "/eolymp.judge.AnnouncementService/CreateAnnouncement"
	AnnouncementService_UpdateAnnouncement_FullMethodName         = "/eolymp.judge.AnnouncementService/UpdateAnnouncement"
	AnnouncementService_DeleteAnnouncement_FullMethodName         = "/eolymp.judge.AnnouncementService/DeleteAnnouncement"
	AnnouncementService_ReadAnnouncement_FullMethodName           = "/eolymp.judge.AnnouncementService/ReadAnnouncement"
	AnnouncementService_DescribeAnnouncement_FullMethodName       = "/eolymp.judge.AnnouncementService/DescribeAnnouncement"
	AnnouncementService_DescribeAnnouncementStatus_FullMethodName = "/eolymp.judge.AnnouncementService/DescribeAnnouncementStatus"
	AnnouncementService_ListAnnouncements_FullMethodName          = "/eolymp.judge.AnnouncementService/ListAnnouncements"
)

// AnnouncementServiceClient is the client API for AnnouncementService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AnnouncementServiceClient interface {
	// Create announcement for a contest
	CreateAnnouncement(ctx context.Context, in *CreateAnnouncementInput, opts ...grpc.CallOption) (*CreateAnnouncementOutput, error)
	// Update existing announcement in a contest
	UpdateAnnouncement(ctx context.Context, in *UpdateAnnouncementInput, opts ...grpc.CallOption) (*UpdateAnnouncementOutput, error)
	// Delete announcement
	DeleteAnnouncement(ctx context.Context, in *DeleteAnnouncementInput, opts ...grpc.CallOption) (*DeleteAnnouncementOutput, error)
	// Mark announcement as read by authenticated participant
	ReadAnnouncement(ctx context.Context, in *ReadAnnouncementInput, opts ...grpc.CallOption) (*ReadAnnouncementOutput, error)
	// Describe announcement
	DescribeAnnouncement(ctx context.Context, in *DescribeAnnouncementInput, opts ...grpc.CallOption) (*DescribeAnnouncementOutput, error)
	// Describe announcement status
	DescribeAnnouncementStatus(ctx context.Context, in *DescribeAnnouncementStatusInput, opts ...grpc.CallOption) (*DescribeAnnouncementStatusOutput, error)
	// List announcements of a contest
	ListAnnouncements(ctx context.Context, in *ListAnnouncementsInput, opts ...grpc.CallOption) (*ListAnnouncementsOutput, error)
}

type announcementServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAnnouncementServiceClient(cc grpc.ClientConnInterface) AnnouncementServiceClient {
	return &announcementServiceClient{cc}
}

func (c *announcementServiceClient) CreateAnnouncement(ctx context.Context, in *CreateAnnouncementInput, opts ...grpc.CallOption) (*CreateAnnouncementOutput, error) {
	out := new(CreateAnnouncementOutput)
	err := c.cc.Invoke(ctx, AnnouncementService_CreateAnnouncement_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *announcementServiceClient) UpdateAnnouncement(ctx context.Context, in *UpdateAnnouncementInput, opts ...grpc.CallOption) (*UpdateAnnouncementOutput, error) {
	out := new(UpdateAnnouncementOutput)
	err := c.cc.Invoke(ctx, AnnouncementService_UpdateAnnouncement_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *announcementServiceClient) DeleteAnnouncement(ctx context.Context, in *DeleteAnnouncementInput, opts ...grpc.CallOption) (*DeleteAnnouncementOutput, error) {
	out := new(DeleteAnnouncementOutput)
	err := c.cc.Invoke(ctx, AnnouncementService_DeleteAnnouncement_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *announcementServiceClient) ReadAnnouncement(ctx context.Context, in *ReadAnnouncementInput, opts ...grpc.CallOption) (*ReadAnnouncementOutput, error) {
	out := new(ReadAnnouncementOutput)
	err := c.cc.Invoke(ctx, AnnouncementService_ReadAnnouncement_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *announcementServiceClient) DescribeAnnouncement(ctx context.Context, in *DescribeAnnouncementInput, opts ...grpc.CallOption) (*DescribeAnnouncementOutput, error) {
	out := new(DescribeAnnouncementOutput)
	err := c.cc.Invoke(ctx, AnnouncementService_DescribeAnnouncement_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *announcementServiceClient) DescribeAnnouncementStatus(ctx context.Context, in *DescribeAnnouncementStatusInput, opts ...grpc.CallOption) (*DescribeAnnouncementStatusOutput, error) {
	out := new(DescribeAnnouncementStatusOutput)
	err := c.cc.Invoke(ctx, AnnouncementService_DescribeAnnouncementStatus_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *announcementServiceClient) ListAnnouncements(ctx context.Context, in *ListAnnouncementsInput, opts ...grpc.CallOption) (*ListAnnouncementsOutput, error) {
	out := new(ListAnnouncementsOutput)
	err := c.cc.Invoke(ctx, AnnouncementService_ListAnnouncements_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AnnouncementServiceServer is the server API for AnnouncementService service.
// All implementations should embed UnimplementedAnnouncementServiceServer
// for forward compatibility
type AnnouncementServiceServer interface {
	// Create announcement for a contest
	CreateAnnouncement(context.Context, *CreateAnnouncementInput) (*CreateAnnouncementOutput, error)
	// Update existing announcement in a contest
	UpdateAnnouncement(context.Context, *UpdateAnnouncementInput) (*UpdateAnnouncementOutput, error)
	// Delete announcement
	DeleteAnnouncement(context.Context, *DeleteAnnouncementInput) (*DeleteAnnouncementOutput, error)
	// Mark announcement as read by authenticated participant
	ReadAnnouncement(context.Context, *ReadAnnouncementInput) (*ReadAnnouncementOutput, error)
	// Describe announcement
	DescribeAnnouncement(context.Context, *DescribeAnnouncementInput) (*DescribeAnnouncementOutput, error)
	// Describe announcement status
	DescribeAnnouncementStatus(context.Context, *DescribeAnnouncementStatusInput) (*DescribeAnnouncementStatusOutput, error)
	// List announcements of a contest
	ListAnnouncements(context.Context, *ListAnnouncementsInput) (*ListAnnouncementsOutput, error)
}

// UnimplementedAnnouncementServiceServer should be embedded to have forward compatible implementations.
type UnimplementedAnnouncementServiceServer struct {
}

func (UnimplementedAnnouncementServiceServer) CreateAnnouncement(context.Context, *CreateAnnouncementInput) (*CreateAnnouncementOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAnnouncement not implemented")
}
func (UnimplementedAnnouncementServiceServer) UpdateAnnouncement(context.Context, *UpdateAnnouncementInput) (*UpdateAnnouncementOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAnnouncement not implemented")
}
func (UnimplementedAnnouncementServiceServer) DeleteAnnouncement(context.Context, *DeleteAnnouncementInput) (*DeleteAnnouncementOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteAnnouncement not implemented")
}
func (UnimplementedAnnouncementServiceServer) ReadAnnouncement(context.Context, *ReadAnnouncementInput) (*ReadAnnouncementOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadAnnouncement not implemented")
}
func (UnimplementedAnnouncementServiceServer) DescribeAnnouncement(context.Context, *DescribeAnnouncementInput) (*DescribeAnnouncementOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DescribeAnnouncement not implemented")
}
func (UnimplementedAnnouncementServiceServer) DescribeAnnouncementStatus(context.Context, *DescribeAnnouncementStatusInput) (*DescribeAnnouncementStatusOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DescribeAnnouncementStatus not implemented")
}
func (UnimplementedAnnouncementServiceServer) ListAnnouncements(context.Context, *ListAnnouncementsInput) (*ListAnnouncementsOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListAnnouncements not implemented")
}

// UnsafeAnnouncementServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AnnouncementServiceServer will
// result in compilation errors.
type UnsafeAnnouncementServiceServer interface {
	mustEmbedUnimplementedAnnouncementServiceServer()
}

func RegisterAnnouncementServiceServer(s grpc.ServiceRegistrar, srv AnnouncementServiceServer) {
	s.RegisterService(&AnnouncementService_ServiceDesc, srv)
}

func _AnnouncementService_CreateAnnouncement_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAnnouncementInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AnnouncementServiceServer).CreateAnnouncement(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AnnouncementService_CreateAnnouncement_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AnnouncementServiceServer).CreateAnnouncement(ctx, req.(*CreateAnnouncementInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _AnnouncementService_UpdateAnnouncement_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateAnnouncementInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AnnouncementServiceServer).UpdateAnnouncement(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AnnouncementService_UpdateAnnouncement_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AnnouncementServiceServer).UpdateAnnouncement(ctx, req.(*UpdateAnnouncementInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _AnnouncementService_DeleteAnnouncement_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteAnnouncementInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AnnouncementServiceServer).DeleteAnnouncement(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AnnouncementService_DeleteAnnouncement_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AnnouncementServiceServer).DeleteAnnouncement(ctx, req.(*DeleteAnnouncementInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _AnnouncementService_ReadAnnouncement_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadAnnouncementInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AnnouncementServiceServer).ReadAnnouncement(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AnnouncementService_ReadAnnouncement_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AnnouncementServiceServer).ReadAnnouncement(ctx, req.(*ReadAnnouncementInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _AnnouncementService_DescribeAnnouncement_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DescribeAnnouncementInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AnnouncementServiceServer).DescribeAnnouncement(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AnnouncementService_DescribeAnnouncement_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AnnouncementServiceServer).DescribeAnnouncement(ctx, req.(*DescribeAnnouncementInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _AnnouncementService_DescribeAnnouncementStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DescribeAnnouncementStatusInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AnnouncementServiceServer).DescribeAnnouncementStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AnnouncementService_DescribeAnnouncementStatus_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AnnouncementServiceServer).DescribeAnnouncementStatus(ctx, req.(*DescribeAnnouncementStatusInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _AnnouncementService_ListAnnouncements_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListAnnouncementsInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AnnouncementServiceServer).ListAnnouncements(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AnnouncementService_ListAnnouncements_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AnnouncementServiceServer).ListAnnouncements(ctx, req.(*ListAnnouncementsInput))
	}
	return interceptor(ctx, in, info, handler)
}

// AnnouncementService_ServiceDesc is the grpc.ServiceDesc for AnnouncementService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AnnouncementService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "eolymp.judge.AnnouncementService",
	HandlerType: (*AnnouncementServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateAnnouncement",
			Handler:    _AnnouncementService_CreateAnnouncement_Handler,
		},
		{
			MethodName: "UpdateAnnouncement",
			Handler:    _AnnouncementService_UpdateAnnouncement_Handler,
		},
		{
			MethodName: "DeleteAnnouncement",
			Handler:    _AnnouncementService_DeleteAnnouncement_Handler,
		},
		{
			MethodName: "ReadAnnouncement",
			Handler:    _AnnouncementService_ReadAnnouncement_Handler,
		},
		{
			MethodName: "DescribeAnnouncement",
			Handler:    _AnnouncementService_DescribeAnnouncement_Handler,
		},
		{
			MethodName: "DescribeAnnouncementStatus",
			Handler:    _AnnouncementService_DescribeAnnouncementStatus_Handler,
		},
		{
			MethodName: "ListAnnouncements",
			Handler:    _AnnouncementService_ListAnnouncements_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "eolymp/judge/announcement_service.proto",
}