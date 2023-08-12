// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.18.1
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
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	StudentService_CreateStudent_FullMethodName   = "/eolymp.course.StudentService/CreateStudent"
	StudentService_UpdateStudent_FullMethodName   = "/eolymp.course.StudentService/UpdateStudent"
	StudentService_DeleteStudent_FullMethodName   = "/eolymp.course.StudentService/DeleteStudent"
	StudentService_DescribeStudent_FullMethodName = "/eolymp.course.StudentService/DescribeStudent"
	StudentService_ListStudents_FullMethodName    = "/eolymp.course.StudentService/ListStudents"
	StudentService_StartCourse_FullMethodName     = "/eolymp.course.StudentService/StartCourse"
	StudentService_DescribeViewer_FullMethodName  = "/eolymp.course.StudentService/DescribeViewer"
)

// StudentServiceClient is the client API for StudentService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StudentServiceClient interface {
	CreateStudent(ctx context.Context, in *CreateStudentInput, opts ...grpc.CallOption) (*CreateStudentOutput, error)
	UpdateStudent(ctx context.Context, in *UpdateStudentInput, opts ...grpc.CallOption) (*UpdateStudentOutput, error)
	DeleteStudent(ctx context.Context, in *DeleteStudentInput, opts ...grpc.CallOption) (*DeleteStudentOutput, error)
	DescribeStudent(ctx context.Context, in *DescribeStudentInput, opts ...grpc.CallOption) (*DescribeStudentOutput, error)
	ListStudents(ctx context.Context, in *ListStudentsInput, opts ...grpc.CallOption) (*ListStudentsOutput, error)
	StartCourse(ctx context.Context, in *StartCourseInput, opts ...grpc.CallOption) (*StartCourseOutput, error)
	DescribeViewer(ctx context.Context, in *DescribeViewerInput, opts ...grpc.CallOption) (*DescribeViewerOutput, error)
}

type studentServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewStudentServiceClient(cc grpc.ClientConnInterface) StudentServiceClient {
	return &studentServiceClient{cc}
}

func (c *studentServiceClient) CreateStudent(ctx context.Context, in *CreateStudentInput, opts ...grpc.CallOption) (*CreateStudentOutput, error) {
	out := new(CreateStudentOutput)
	err := c.cc.Invoke(ctx, StudentService_CreateStudent_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *studentServiceClient) UpdateStudent(ctx context.Context, in *UpdateStudentInput, opts ...grpc.CallOption) (*UpdateStudentOutput, error) {
	out := new(UpdateStudentOutput)
	err := c.cc.Invoke(ctx, StudentService_UpdateStudent_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *studentServiceClient) DeleteStudent(ctx context.Context, in *DeleteStudentInput, opts ...grpc.CallOption) (*DeleteStudentOutput, error) {
	out := new(DeleteStudentOutput)
	err := c.cc.Invoke(ctx, StudentService_DeleteStudent_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *studentServiceClient) DescribeStudent(ctx context.Context, in *DescribeStudentInput, opts ...grpc.CallOption) (*DescribeStudentOutput, error) {
	out := new(DescribeStudentOutput)
	err := c.cc.Invoke(ctx, StudentService_DescribeStudent_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *studentServiceClient) ListStudents(ctx context.Context, in *ListStudentsInput, opts ...grpc.CallOption) (*ListStudentsOutput, error) {
	out := new(ListStudentsOutput)
	err := c.cc.Invoke(ctx, StudentService_ListStudents_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *studentServiceClient) StartCourse(ctx context.Context, in *StartCourseInput, opts ...grpc.CallOption) (*StartCourseOutput, error) {
	out := new(StartCourseOutput)
	err := c.cc.Invoke(ctx, StudentService_StartCourse_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *studentServiceClient) DescribeViewer(ctx context.Context, in *DescribeViewerInput, opts ...grpc.CallOption) (*DescribeViewerOutput, error) {
	out := new(DescribeViewerOutput)
	err := c.cc.Invoke(ctx, StudentService_DescribeViewer_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StudentServiceServer is the server API for StudentService service.
// All implementations should embed UnimplementedStudentServiceServer
// for forward compatibility
type StudentServiceServer interface {
	CreateStudent(context.Context, *CreateStudentInput) (*CreateStudentOutput, error)
	UpdateStudent(context.Context, *UpdateStudentInput) (*UpdateStudentOutput, error)
	DeleteStudent(context.Context, *DeleteStudentInput) (*DeleteStudentOutput, error)
	DescribeStudent(context.Context, *DescribeStudentInput) (*DescribeStudentOutput, error)
	ListStudents(context.Context, *ListStudentsInput) (*ListStudentsOutput, error)
	StartCourse(context.Context, *StartCourseInput) (*StartCourseOutput, error)
	DescribeViewer(context.Context, *DescribeViewerInput) (*DescribeViewerOutput, error)
}

// UnimplementedStudentServiceServer should be embedded to have forward compatible implementations.
type UnimplementedStudentServiceServer struct {
}

func (UnimplementedStudentServiceServer) CreateStudent(context.Context, *CreateStudentInput) (*CreateStudentOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateStudent not implemented")
}
func (UnimplementedStudentServiceServer) UpdateStudent(context.Context, *UpdateStudentInput) (*UpdateStudentOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateStudent not implemented")
}
func (UnimplementedStudentServiceServer) DeleteStudent(context.Context, *DeleteStudentInput) (*DeleteStudentOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteStudent not implemented")
}
func (UnimplementedStudentServiceServer) DescribeStudent(context.Context, *DescribeStudentInput) (*DescribeStudentOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DescribeStudent not implemented")
}
func (UnimplementedStudentServiceServer) ListStudents(context.Context, *ListStudentsInput) (*ListStudentsOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListStudents not implemented")
}
func (UnimplementedStudentServiceServer) StartCourse(context.Context, *StartCourseInput) (*StartCourseOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StartCourse not implemented")
}
func (UnimplementedStudentServiceServer) DescribeViewer(context.Context, *DescribeViewerInput) (*DescribeViewerOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DescribeViewer not implemented")
}

// UnsafeStudentServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StudentServiceServer will
// result in compilation errors.
type UnsafeStudentServiceServer interface {
	mustEmbedUnimplementedStudentServiceServer()
}

func RegisterStudentServiceServer(s grpc.ServiceRegistrar, srv StudentServiceServer) {
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

func _StudentService_StartCourse_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartCourseInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StudentServiceServer).StartCourse(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StudentService_StartCourse_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StudentServiceServer).StartCourse(ctx, req.(*StartCourseInput))
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
			MethodName: "DescribeStudent",
			Handler:    _StudentService_DescribeStudent_Handler,
		},
		{
			MethodName: "ListStudents",
			Handler:    _StudentService_ListStudents_Handler,
		},
		{
			MethodName: "StartCourse",
			Handler:    _StudentService_StartCourse_Handler,
		},
		{
			MethodName: "DescribeViewer",
			Handler:    _StudentService_DescribeViewer_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "eolymp/course/student_service.proto",
}
