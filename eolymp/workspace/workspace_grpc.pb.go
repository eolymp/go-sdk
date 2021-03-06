// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.18.1
// source: eolymp/workspace/workspace.proto

package workspace

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

// WorkspaceClient is the client API for Workspace service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WorkspaceClient interface {
	DescribeProject(ctx context.Context, in *DescribeProjectInput, opts ...grpc.CallOption) (*DescribeProjectOutput, error)
	ListProjects(ctx context.Context, in *ListProjectsInput, opts ...grpc.CallOption) (*ListProjectsOutput, error)
	CreateProject(ctx context.Context, in *CreateProjectInput, opts ...grpc.CallOption) (*CreateProjectOutput, error)
	UpdateProject(ctx context.Context, in *UpdateProjectInput, opts ...grpc.CallOption) (*UpdateProjectOutput, error)
	DeleteProject(ctx context.Context, in *DeleteProjectInput, opts ...grpc.CallOption) (*DeleteProjectOutput, error)
	ListFiles(ctx context.Context, in *ListFilesInput, opts ...grpc.CallOption) (*ListFilesOutput, error)
	DescribeFile(ctx context.Context, in *DescribeFileInput, opts ...grpc.CallOption) (*DescribeFileOutput, error)
	UploadFile(ctx context.Context, in *UploadFileInput, opts ...grpc.CallOption) (*UploadFileOutput, error)
	RemoveFile(ctx context.Context, in *RemoveFileInput, opts ...grpc.CallOption) (*RemoveFileOutput, error)
}

type workspaceClient struct {
	cc grpc.ClientConnInterface
}

func NewWorkspaceClient(cc grpc.ClientConnInterface) WorkspaceClient {
	return &workspaceClient{cc}
}

func (c *workspaceClient) DescribeProject(ctx context.Context, in *DescribeProjectInput, opts ...grpc.CallOption) (*DescribeProjectOutput, error) {
	out := new(DescribeProjectOutput)
	err := c.cc.Invoke(ctx, "/eolymp.workspace.Workspace/DescribeProject", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workspaceClient) ListProjects(ctx context.Context, in *ListProjectsInput, opts ...grpc.CallOption) (*ListProjectsOutput, error) {
	out := new(ListProjectsOutput)
	err := c.cc.Invoke(ctx, "/eolymp.workspace.Workspace/ListProjects", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workspaceClient) CreateProject(ctx context.Context, in *CreateProjectInput, opts ...grpc.CallOption) (*CreateProjectOutput, error) {
	out := new(CreateProjectOutput)
	err := c.cc.Invoke(ctx, "/eolymp.workspace.Workspace/CreateProject", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workspaceClient) UpdateProject(ctx context.Context, in *UpdateProjectInput, opts ...grpc.CallOption) (*UpdateProjectOutput, error) {
	out := new(UpdateProjectOutput)
	err := c.cc.Invoke(ctx, "/eolymp.workspace.Workspace/UpdateProject", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workspaceClient) DeleteProject(ctx context.Context, in *DeleteProjectInput, opts ...grpc.CallOption) (*DeleteProjectOutput, error) {
	out := new(DeleteProjectOutput)
	err := c.cc.Invoke(ctx, "/eolymp.workspace.Workspace/DeleteProject", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workspaceClient) ListFiles(ctx context.Context, in *ListFilesInput, opts ...grpc.CallOption) (*ListFilesOutput, error) {
	out := new(ListFilesOutput)
	err := c.cc.Invoke(ctx, "/eolymp.workspace.Workspace/ListFiles", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workspaceClient) DescribeFile(ctx context.Context, in *DescribeFileInput, opts ...grpc.CallOption) (*DescribeFileOutput, error) {
	out := new(DescribeFileOutput)
	err := c.cc.Invoke(ctx, "/eolymp.workspace.Workspace/DescribeFile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workspaceClient) UploadFile(ctx context.Context, in *UploadFileInput, opts ...grpc.CallOption) (*UploadFileOutput, error) {
	out := new(UploadFileOutput)
	err := c.cc.Invoke(ctx, "/eolymp.workspace.Workspace/UploadFile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workspaceClient) RemoveFile(ctx context.Context, in *RemoveFileInput, opts ...grpc.CallOption) (*RemoveFileOutput, error) {
	out := new(RemoveFileOutput)
	err := c.cc.Invoke(ctx, "/eolymp.workspace.Workspace/RemoveFile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WorkspaceServer is the server API for Workspace service.
// All implementations should embed UnimplementedWorkspaceServer
// for forward compatibility
type WorkspaceServer interface {
	DescribeProject(context.Context, *DescribeProjectInput) (*DescribeProjectOutput, error)
	ListProjects(context.Context, *ListProjectsInput) (*ListProjectsOutput, error)
	CreateProject(context.Context, *CreateProjectInput) (*CreateProjectOutput, error)
	UpdateProject(context.Context, *UpdateProjectInput) (*UpdateProjectOutput, error)
	DeleteProject(context.Context, *DeleteProjectInput) (*DeleteProjectOutput, error)
	ListFiles(context.Context, *ListFilesInput) (*ListFilesOutput, error)
	DescribeFile(context.Context, *DescribeFileInput) (*DescribeFileOutput, error)
	UploadFile(context.Context, *UploadFileInput) (*UploadFileOutput, error)
	RemoveFile(context.Context, *RemoveFileInput) (*RemoveFileOutput, error)
}

// UnimplementedWorkspaceServer should be embedded to have forward compatible implementations.
type UnimplementedWorkspaceServer struct {
}

func (UnimplementedWorkspaceServer) DescribeProject(context.Context, *DescribeProjectInput) (*DescribeProjectOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DescribeProject not implemented")
}
func (UnimplementedWorkspaceServer) ListProjects(context.Context, *ListProjectsInput) (*ListProjectsOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListProjects not implemented")
}
func (UnimplementedWorkspaceServer) CreateProject(context.Context, *CreateProjectInput) (*CreateProjectOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateProject not implemented")
}
func (UnimplementedWorkspaceServer) UpdateProject(context.Context, *UpdateProjectInput) (*UpdateProjectOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateProject not implemented")
}
func (UnimplementedWorkspaceServer) DeleteProject(context.Context, *DeleteProjectInput) (*DeleteProjectOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteProject not implemented")
}
func (UnimplementedWorkspaceServer) ListFiles(context.Context, *ListFilesInput) (*ListFilesOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListFiles not implemented")
}
func (UnimplementedWorkspaceServer) DescribeFile(context.Context, *DescribeFileInput) (*DescribeFileOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DescribeFile not implemented")
}
func (UnimplementedWorkspaceServer) UploadFile(context.Context, *UploadFileInput) (*UploadFileOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UploadFile not implemented")
}
func (UnimplementedWorkspaceServer) RemoveFile(context.Context, *RemoveFileInput) (*RemoveFileOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveFile not implemented")
}

// UnsafeWorkspaceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WorkspaceServer will
// result in compilation errors.
type UnsafeWorkspaceServer interface {
	mustEmbedUnimplementedWorkspaceServer()
}

func RegisterWorkspaceServer(s grpc.ServiceRegistrar, srv WorkspaceServer) {
	s.RegisterService(&Workspace_ServiceDesc, srv)
}

func _Workspace_DescribeProject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DescribeProjectInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkspaceServer).DescribeProject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/eolymp.workspace.Workspace/DescribeProject",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkspaceServer).DescribeProject(ctx, req.(*DescribeProjectInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _Workspace_ListProjects_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListProjectsInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkspaceServer).ListProjects(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/eolymp.workspace.Workspace/ListProjects",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkspaceServer).ListProjects(ctx, req.(*ListProjectsInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _Workspace_CreateProject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateProjectInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkspaceServer).CreateProject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/eolymp.workspace.Workspace/CreateProject",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkspaceServer).CreateProject(ctx, req.(*CreateProjectInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _Workspace_UpdateProject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateProjectInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkspaceServer).UpdateProject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/eolymp.workspace.Workspace/UpdateProject",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkspaceServer).UpdateProject(ctx, req.(*UpdateProjectInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _Workspace_DeleteProject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteProjectInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkspaceServer).DeleteProject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/eolymp.workspace.Workspace/DeleteProject",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkspaceServer).DeleteProject(ctx, req.(*DeleteProjectInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _Workspace_ListFiles_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListFilesInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkspaceServer).ListFiles(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/eolymp.workspace.Workspace/ListFiles",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkspaceServer).ListFiles(ctx, req.(*ListFilesInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _Workspace_DescribeFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DescribeFileInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkspaceServer).DescribeFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/eolymp.workspace.Workspace/DescribeFile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkspaceServer).DescribeFile(ctx, req.(*DescribeFileInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _Workspace_UploadFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UploadFileInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkspaceServer).UploadFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/eolymp.workspace.Workspace/UploadFile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkspaceServer).UploadFile(ctx, req.(*UploadFileInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _Workspace_RemoveFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveFileInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkspaceServer).RemoveFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/eolymp.workspace.Workspace/RemoveFile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkspaceServer).RemoveFile(ctx, req.(*RemoveFileInput))
	}
	return interceptor(ctx, in, info, handler)
}

// Workspace_ServiceDesc is the grpc.ServiceDesc for Workspace service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Workspace_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "eolymp.workspace.Workspace",
	HandlerType: (*WorkspaceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DescribeProject",
			Handler:    _Workspace_DescribeProject_Handler,
		},
		{
			MethodName: "ListProjects",
			Handler:    _Workspace_ListProjects_Handler,
		},
		{
			MethodName: "CreateProject",
			Handler:    _Workspace_CreateProject_Handler,
		},
		{
			MethodName: "UpdateProject",
			Handler:    _Workspace_UpdateProject_Handler,
		},
		{
			MethodName: "DeleteProject",
			Handler:    _Workspace_DeleteProject_Handler,
		},
		{
			MethodName: "ListFiles",
			Handler:    _Workspace_ListFiles_Handler,
		},
		{
			MethodName: "DescribeFile",
			Handler:    _Workspace_DescribeFile_Handler,
		},
		{
			MethodName: "UploadFile",
			Handler:    _Workspace_UploadFile_Handler,
		},
		{
			MethodName: "RemoveFile",
			Handler:    _Workspace_RemoveFile_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "eolymp/workspace/workspace.proto",
}
