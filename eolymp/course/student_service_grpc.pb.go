// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.21.12
// source: eolymp/course/student_service.proto

package course

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
	StudentService_CreateStudent_FullMethodName   = "/eolymp.course.StudentService/CreateStudent"
	StudentService_UpdateStudent_FullMethodName   = "/eolymp.course.StudentService/UpdateStudent"
	StudentService_DeleteStudent_FullMethodName   = "/eolymp.course.StudentService/DeleteStudent"
	StudentService_DescribeViewer_FullMethodName  = "/eolymp.course.StudentService/DescribeViewer"
	StudentService_DescribeStudent_FullMethodName = "/eolymp.course.StudentService/DescribeStudent"
	StudentService_ListStudents_FullMethodName    = "/eolymp.course.StudentService/ListStudents"
	StudentService_WatchStudent_FullMethodName    = "/eolymp.course.StudentService/WatchStudent"
)

// StudentServiceClient is the client API for StudentService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StudentServiceClient interface {
	CreateStudent(ctx context.Context, in *CreateStudentInput, opts ...grpc.CallOption) (*CreateStudentOutput, error)
	UpdateStudent(ctx context.Context, in *UpdateStudentInput, opts ...grpc.CallOption) (*UpdateStudentOutput, error)
	DeleteStudent(ctx context.Context, in *DeleteStudentInput, opts ...grpc.CallOption) (*DeleteStudentOutput, error)
	DescribeViewer(ctx context.Context, in *DescribeViewerInput, opts ...grpc.CallOption) (*DescribeViewerOutput, error)
	DescribeStudent(ctx context.Context, in *DescribeStudentInput, opts ...grpc.CallOption) (*DescribeStudentOutput, error)
	ListStudents(ctx context.Context, in *ListStudentsInput, opts ...grpc.CallOption) (*ListStudentsOutput, error)
	WatchStudent(ctx context.Context, in *WatchStudentInput, opts ...grpc.CallOption) (grpc.ServerStreamingClient[WatchStudentOutput], error)
}

type studentServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewStudentServiceClient(cc grpc.ClientConnInterface) StudentServiceClient {
	return &studentServiceClient{cc}
}

func (c *studentServiceClient) CreateStudent(ctx context.Context, in *CreateStudentInput, opts ...grpc.CallOption) (*CreateStudentOutput, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateStudentOutput)
	err := c.cc.Invoke(ctx, StudentService_CreateStudent_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *studentServiceClient) UpdateStudent(ctx context.Context, in *UpdateStudentInput, opts ...grpc.CallOption) (*UpdateStudentOutput, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateStudentOutput)
	err := c.cc.Invoke(ctx, StudentService_UpdateStudent_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *studentServiceClient) DeleteStudent(ctx context.Context, in *DeleteStudentInput, opts ...grpc.CallOption) (*DeleteStudentOutput, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteStudentOutput)
	err := c.cc.Invoke(ctx, StudentService_DeleteStudent_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *studentServiceClient) DescribeViewer(ctx context.Context, in *DescribeViewerInput, opts ...grpc.CallOption) (*DescribeViewerOutput, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DescribeViewerOutput)
	err := c.cc.Invoke(ctx, StudentService_DescribeViewer_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *studentServiceClient) DescribeStudent(ctx context.Context, in *DescribeStudentInput, opts ...grpc.CallOption) (*DescribeStudentOutput, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DescribeStudentOutput)
	err := c.cc.Invoke(ctx, StudentService_DescribeStudent_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *studentServiceClient) ListStudents(ctx context.Context, in *ListStudentsInput, opts ...grpc.CallOption) (*ListStudentsOutput, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListStudentsOutput)
	err := c.cc.Invoke(ctx, StudentService_ListStudents_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *studentServiceClient) WatchStudent(ctx context.Context, in *WatchStudentInput, opts ...grpc.CallOption) (grpc.ServerStreamingClient[WatchStudentOutput], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &StudentService_ServiceDesc.Streams[0], StudentService_WatchStudent_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[WatchStudentInput, WatchStudentOutput]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type StudentService_WatchStudentClient = grpc.ServerStreamingClient[WatchStudentOutput]

// StudentServiceServer is the server API for StudentService service.
// All implementations should embed UnimplementedStudentServiceServer
// for forward compatibility.
type StudentServiceServer interface {
	CreateStudent(context.Context, *CreateStudentInput) (*CreateStudentOutput, error)
	UpdateStudent(context.Context, *UpdateStudentInput) (*UpdateStudentOutput, error)
	DeleteStudent(context.Context, *DeleteStudentInput) (*DeleteStudentOutput, error)
	DescribeViewer(context.Context, *DescribeViewerInput) (*DescribeViewerOutput, error)
	DescribeStudent(context.Context, *DescribeStudentInput) (*DescribeStudentOutput, error)
	ListStudents(context.Context, *ListStudentsInput) (*ListStudentsOutput, error)
	WatchStudent(*WatchStudentInput, grpc.ServerStreamingServer[WatchStudentOutput]) error
}

// UnimplementedStudentServiceServer should be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedStudentServiceServer struct{}

func (UnimplementedStudentServiceServer) CreateStudent(context.Context, *CreateStudentInput) (*CreateStudentOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateStudent not implemented")
}
func (UnimplementedStudentServiceServer) UpdateStudent(context.Context, *UpdateStudentInput) (*UpdateStudentOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateStudent not implemented")
}
func (UnimplementedStudentServiceServer) DeleteStudent(context.Context, *DeleteStudentInput) (*DeleteStudentOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteStudent not implemented")
}
func (UnimplementedStudentServiceServer) DescribeViewer(context.Context, *DescribeViewerInput) (*DescribeViewerOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DescribeViewer not implemented")
}
func (UnimplementedStudentServiceServer) DescribeStudent(context.Context, *DescribeStudentInput) (*DescribeStudentOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DescribeStudent not implemented")
}
func (UnimplementedStudentServiceServer) ListStudents(context.Context, *ListStudentsInput) (*ListStudentsOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListStudents not implemented")
}
func (UnimplementedStudentServiceServer) WatchStudent(*WatchStudentInput, grpc.ServerStreamingServer[WatchStudentOutput]) error {
	return status.Errorf(codes.Unimplemented, "method WatchStudent not implemented")
}
func (UnimplementedStudentServiceServer) testEmbeddedByValue() {}

// UnsafeStudentServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StudentServiceServer will
// result in compilation errors.
type UnsafeStudentServiceServer interface {
	mustEmbedUnimplementedStudentServiceServer()
}

func RegisterStudentServiceServer(s grpc.ServiceRegistrar, srv StudentServiceServer) {
	// If the following call pancis, it indicates UnimplementedStudentServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&StudentService_ServiceDesc, srv)
}

func _StudentService_CreateStudent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateStudentInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StudentServiceServer).CreateStudent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StudentService_CreateStudent_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StudentServiceServer).CreateStudent(ctx, req.(*CreateStudentInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _StudentService_UpdateStudent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateStudentInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StudentServiceServer).UpdateStudent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StudentService_UpdateStudent_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StudentServiceServer).UpdateStudent(ctx, req.(*UpdateStudentInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _StudentService_DeleteStudent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteStudentInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StudentServiceServer).DeleteStudent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StudentService_DeleteStudent_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StudentServiceServer).DeleteStudent(ctx, req.(*DeleteStudentInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _StudentService_DescribeViewer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DescribeViewerInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StudentServiceServer).DescribeViewer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StudentService_DescribeViewer_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StudentServiceServer).DescribeViewer(ctx, req.(*DescribeViewerInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _StudentService_DescribeStudent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DescribeStudentInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StudentServiceServer).DescribeStudent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StudentService_DescribeStudent_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StudentServiceServer).DescribeStudent(ctx, req.(*DescribeStudentInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _StudentService_ListStudents_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListStudentsInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StudentServiceServer).ListStudents(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StudentService_ListStudents_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StudentServiceServer).ListStudents(ctx, req.(*ListStudentsInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _StudentService_WatchStudent_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(WatchStudentInput)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(StudentServiceServer).WatchStudent(m, &grpc.GenericServerStream[WatchStudentInput, WatchStudentOutput]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type StudentService_WatchStudentServer = grpc.ServerStreamingServer[WatchStudentOutput]

// StudentService_ServiceDesc is the grpc.ServiceDesc for StudentService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var StudentService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "eolymp.course.StudentService",
	HandlerType: (*StudentServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateStudent",
			Handler:    _StudentService_CreateStudent_Handler,
		},
		{
			MethodName: "UpdateStudent",
			Handler:    _StudentService_UpdateStudent_Handler,
		},
		{
			MethodName: "DeleteStudent",
			Handler:    _StudentService_DeleteStudent_Handler,
		},
		{
			MethodName: "DescribeViewer",
			Handler:    _StudentService_DescribeViewer_Handler,
		},
		{
			MethodName: "DescribeStudent",
			Handler:    _StudentService_DescribeStudent_Handler,
		},
		{
			MethodName: "ListStudents",
			Handler:    _StudentService_ListStudents_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "WatchStudent",
			Handler:       _StudentService_WatchStudent_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "eolymp/course/student_service.proto",
}
