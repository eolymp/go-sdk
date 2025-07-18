// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.4
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
	StudentService_CreateStudent_FullMethodName           = "/eolymp.course.StudentService/CreateStudent"
	StudentService_UpdateStudent_FullMethodName           = "/eolymp.course.StudentService/UpdateStudent"
	StudentService_DeleteStudent_FullMethodName           = "/eolymp.course.StudentService/DeleteStudent"
	StudentService_DescribeStudent_FullMethodName         = "/eolymp.course.StudentService/DescribeStudent"
	StudentService_ListStudents_FullMethodName            = "/eolymp.course.StudentService/ListStudents"
	StudentService_WatchStudent_FullMethodName            = "/eolymp.course.StudentService/WatchStudent"
	StudentService_JoinCourse_FullMethodName              = "/eolymp.course.StudentService/JoinCourse"
	StudentService_DescribeViewer_FullMethodName          = "/eolymp.course.StudentService/DescribeViewer"
	StudentService_ListStudentAssignments_FullMethodName  = "/eolymp.course.StudentService/ListStudentAssignments"
	StudentService_UpdateStudentAssignment_FullMethodName = "/eolymp.course.StudentService/UpdateStudentAssignment"
	StudentService_DeleteStudentAssignment_FullMethodName = "/eolymp.course.StudentService/DeleteStudentAssignment"
	StudentService_ListStudentGrades_FullMethodName       = "/eolymp.course.StudentService/ListStudentGrades"
	StudentService_ListModuleGrades_FullMethodName        = "/eolymp.course.StudentService/ListModuleGrades"
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
	WatchStudent(ctx context.Context, in *WatchStudentInput, opts ...grpc.CallOption) (grpc.ServerStreamingClient[WatchStudentOutput], error)
	JoinCourse(ctx context.Context, in *JoinCourseInput, opts ...grpc.CallOption) (*JoinCourseOutput, error)
	DescribeViewer(ctx context.Context, in *DescribeViewerInput, opts ...grpc.CallOption) (*DescribeViewerOutput, error)
	ListStudentAssignments(ctx context.Context, in *ListStudentAssignmentsInput, opts ...grpc.CallOption) (*ListStudentAssignmentsOutput, error)
	UpdateStudentAssignment(ctx context.Context, in *UpdateStudentAssignmentInput, opts ...grpc.CallOption) (*UpdateStudentAssignmentOutput, error)
	DeleteStudentAssignment(ctx context.Context, in *DeleteStudentAssignmentInput, opts ...grpc.CallOption) (*DeleteStudentAssignmentOutput, error)
	ListStudentGrades(ctx context.Context, in *ListStudentGradesInput, opts ...grpc.CallOption) (*ListStudentGradesOutput, error)
	ListModuleGrades(ctx context.Context, in *ListModuleGradesInput, opts ...grpc.CallOption) (*ListModuleGradesOutput, error)
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

func (c *studentServiceClient) JoinCourse(ctx context.Context, in *JoinCourseInput, opts ...grpc.CallOption) (*JoinCourseOutput, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(JoinCourseOutput)
	err := c.cc.Invoke(ctx, StudentService_JoinCourse_FullMethodName, in, out, cOpts...)
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

func (c *studentServiceClient) ListStudentAssignments(ctx context.Context, in *ListStudentAssignmentsInput, opts ...grpc.CallOption) (*ListStudentAssignmentsOutput, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListStudentAssignmentsOutput)
	err := c.cc.Invoke(ctx, StudentService_ListStudentAssignments_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *studentServiceClient) UpdateStudentAssignment(ctx context.Context, in *UpdateStudentAssignmentInput, opts ...grpc.CallOption) (*UpdateStudentAssignmentOutput, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateStudentAssignmentOutput)
	err := c.cc.Invoke(ctx, StudentService_UpdateStudentAssignment_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *studentServiceClient) DeleteStudentAssignment(ctx context.Context, in *DeleteStudentAssignmentInput, opts ...grpc.CallOption) (*DeleteStudentAssignmentOutput, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteStudentAssignmentOutput)
	err := c.cc.Invoke(ctx, StudentService_DeleteStudentAssignment_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *studentServiceClient) ListStudentGrades(ctx context.Context, in *ListStudentGradesInput, opts ...grpc.CallOption) (*ListStudentGradesOutput, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListStudentGradesOutput)
	err := c.cc.Invoke(ctx, StudentService_ListStudentGrades_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *studentServiceClient) ListModuleGrades(ctx context.Context, in *ListModuleGradesInput, opts ...grpc.CallOption) (*ListModuleGradesOutput, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListModuleGradesOutput)
	err := c.cc.Invoke(ctx, StudentService_ListModuleGrades_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StudentServiceServer is the server API for StudentService service.
// All implementations should embed UnimplementedStudentServiceServer
// for forward compatibility.
type StudentServiceServer interface {
	CreateStudent(context.Context, *CreateStudentInput) (*CreateStudentOutput, error)
	UpdateStudent(context.Context, *UpdateStudentInput) (*UpdateStudentOutput, error)
	DeleteStudent(context.Context, *DeleteStudentInput) (*DeleteStudentOutput, error)
	DescribeStudent(context.Context, *DescribeStudentInput) (*DescribeStudentOutput, error)
	ListStudents(context.Context, *ListStudentsInput) (*ListStudentsOutput, error)
	WatchStudent(*WatchStudentInput, grpc.ServerStreamingServer[WatchStudentOutput]) error
	JoinCourse(context.Context, *JoinCourseInput) (*JoinCourseOutput, error)
	DescribeViewer(context.Context, *DescribeViewerInput) (*DescribeViewerOutput, error)
	ListStudentAssignments(context.Context, *ListStudentAssignmentsInput) (*ListStudentAssignmentsOutput, error)
	UpdateStudentAssignment(context.Context, *UpdateStudentAssignmentInput) (*UpdateStudentAssignmentOutput, error)
	DeleteStudentAssignment(context.Context, *DeleteStudentAssignmentInput) (*DeleteStudentAssignmentOutput, error)
	ListStudentGrades(context.Context, *ListStudentGradesInput) (*ListStudentGradesOutput, error)
	ListModuleGrades(context.Context, *ListModuleGradesInput) (*ListModuleGradesOutput, error)
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
func (UnimplementedStudentServiceServer) DescribeStudent(context.Context, *DescribeStudentInput) (*DescribeStudentOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DescribeStudent not implemented")
}
func (UnimplementedStudentServiceServer) ListStudents(context.Context, *ListStudentsInput) (*ListStudentsOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListStudents not implemented")
}
func (UnimplementedStudentServiceServer) WatchStudent(*WatchStudentInput, grpc.ServerStreamingServer[WatchStudentOutput]) error {
	return status.Errorf(codes.Unimplemented, "method WatchStudent not implemented")
}
func (UnimplementedStudentServiceServer) JoinCourse(context.Context, *JoinCourseInput) (*JoinCourseOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method JoinCourse not implemented")
}
func (UnimplementedStudentServiceServer) DescribeViewer(context.Context, *DescribeViewerInput) (*DescribeViewerOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DescribeViewer not implemented")
}
func (UnimplementedStudentServiceServer) ListStudentAssignments(context.Context, *ListStudentAssignmentsInput) (*ListStudentAssignmentsOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListStudentAssignments not implemented")
}
func (UnimplementedStudentServiceServer) UpdateStudentAssignment(context.Context, *UpdateStudentAssignmentInput) (*UpdateStudentAssignmentOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateStudentAssignment not implemented")
}
func (UnimplementedStudentServiceServer) DeleteStudentAssignment(context.Context, *DeleteStudentAssignmentInput) (*DeleteStudentAssignmentOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteStudentAssignment not implemented")
}
func (UnimplementedStudentServiceServer) ListStudentGrades(context.Context, *ListStudentGradesInput) (*ListStudentGradesOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListStudentGrades not implemented")
}
func (UnimplementedStudentServiceServer) ListModuleGrades(context.Context, *ListModuleGradesInput) (*ListModuleGradesOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListModuleGrades not implemented")
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

func _StudentService_JoinCourse_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JoinCourseInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StudentServiceServer).JoinCourse(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StudentService_JoinCourse_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StudentServiceServer).JoinCourse(ctx, req.(*JoinCourseInput))
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

func _StudentService_ListStudentAssignments_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListStudentAssignmentsInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StudentServiceServer).ListStudentAssignments(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StudentService_ListStudentAssignments_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StudentServiceServer).ListStudentAssignments(ctx, req.(*ListStudentAssignmentsInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _StudentService_UpdateStudentAssignment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateStudentAssignmentInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StudentServiceServer).UpdateStudentAssignment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StudentService_UpdateStudentAssignment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StudentServiceServer).UpdateStudentAssignment(ctx, req.(*UpdateStudentAssignmentInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _StudentService_DeleteStudentAssignment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteStudentAssignmentInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StudentServiceServer).DeleteStudentAssignment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StudentService_DeleteStudentAssignment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StudentServiceServer).DeleteStudentAssignment(ctx, req.(*DeleteStudentAssignmentInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _StudentService_ListStudentGrades_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListStudentGradesInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StudentServiceServer).ListStudentGrades(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StudentService_ListStudentGrades_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StudentServiceServer).ListStudentGrades(ctx, req.(*ListStudentGradesInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _StudentService_ListModuleGrades_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListModuleGradesInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StudentServiceServer).ListModuleGrades(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StudentService_ListModuleGrades_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StudentServiceServer).ListModuleGrades(ctx, req.(*ListModuleGradesInput))
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
			MethodName: "JoinCourse",
			Handler:    _StudentService_JoinCourse_Handler,
		},
		{
			MethodName: "DescribeViewer",
			Handler:    _StudentService_DescribeViewer_Handler,
		},
		{
			MethodName: "ListStudentAssignments",
			Handler:    _StudentService_ListStudentAssignments_Handler,
		},
		{
			MethodName: "UpdateStudentAssignment",
			Handler:    _StudentService_UpdateStudentAssignment_Handler,
		},
		{
			MethodName: "DeleteStudentAssignment",
			Handler:    _StudentService_DeleteStudentAssignment_Handler,
		},
		{
			MethodName: "ListStudentGrades",
			Handler:    _StudentService_ListStudentGrades_Handler,
		},
		{
			MethodName: "ListModuleGrades",
			Handler:    _StudentService_ListModuleGrades_Handler,
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
