// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.18.1
// source: eolymp/atlas/suggestion_service.proto

package atlas

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
	SuggestionService_CreateSuggestion_FullMethodName   = "/eolymp.atlas.SuggestionService/CreateSuggestion"
	SuggestionService_UpdateSuggestion_FullMethodName   = "/eolymp.atlas.SuggestionService/UpdateSuggestion"
	SuggestionService_ReviewSuggestion_FullMethodName   = "/eolymp.atlas.SuggestionService/ReviewSuggestion"
	SuggestionService_ResubmitSuggestion_FullMethodName = "/eolymp.atlas.SuggestionService/ResubmitSuggestion"
	SuggestionService_DeleteSuggestion_FullMethodName   = "/eolymp.atlas.SuggestionService/DeleteSuggestion"
	SuggestionService_ListSuggestions_FullMethodName    = "/eolymp.atlas.SuggestionService/ListSuggestions"
	SuggestionService_DescribeSuggestion_FullMethodName = "/eolymp.atlas.SuggestionService/DescribeSuggestion"
)

// SuggestionServiceClient is the client API for SuggestionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SuggestionServiceClient interface {
	CreateSuggestion(ctx context.Context, in *CreateSuggestionInput, opts ...grpc.CallOption) (*CreateSuggestionOutput, error)
	UpdateSuggestion(ctx context.Context, in *UpdateSuggestionInput, opts ...grpc.CallOption) (*UpdateSuggestionOutput, error)
	ReviewSuggestion(ctx context.Context, in *ReviewSuggestionInput, opts ...grpc.CallOption) (*ReviewSuggestionOutput, error)
	ResubmitSuggestion(ctx context.Context, in *ResubmitSuggestionInput, opts ...grpc.CallOption) (*ResubmitSuggestionOutput, error)
	DeleteSuggestion(ctx context.Context, in *DeleteSuggestionInput, opts ...grpc.CallOption) (*DeleteSuggestionOutput, error)
	ListSuggestions(ctx context.Context, in *ListSuggestionsInput, opts ...grpc.CallOption) (*ListSuggestionsOutput, error)
	DescribeSuggestion(ctx context.Context, in *DescribeSuggestionInput, opts ...grpc.CallOption) (*DescribeSuggestionOutput, error)
}

type suggestionServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSuggestionServiceClient(cc grpc.ClientConnInterface) SuggestionServiceClient {
	return &suggestionServiceClient{cc}
}

func (c *suggestionServiceClient) CreateSuggestion(ctx context.Context, in *CreateSuggestionInput, opts ...grpc.CallOption) (*CreateSuggestionOutput, error) {
	out := new(CreateSuggestionOutput)
	err := c.cc.Invoke(ctx, SuggestionService_CreateSuggestion_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *suggestionServiceClient) UpdateSuggestion(ctx context.Context, in *UpdateSuggestionInput, opts ...grpc.CallOption) (*UpdateSuggestionOutput, error) {
	out := new(UpdateSuggestionOutput)
	err := c.cc.Invoke(ctx, SuggestionService_UpdateSuggestion_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *suggestionServiceClient) ReviewSuggestion(ctx context.Context, in *ReviewSuggestionInput, opts ...grpc.CallOption) (*ReviewSuggestionOutput, error) {
	out := new(ReviewSuggestionOutput)
	err := c.cc.Invoke(ctx, SuggestionService_ReviewSuggestion_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *suggestionServiceClient) ResubmitSuggestion(ctx context.Context, in *ResubmitSuggestionInput, opts ...grpc.CallOption) (*ResubmitSuggestionOutput, error) {
	out := new(ResubmitSuggestionOutput)
	err := c.cc.Invoke(ctx, SuggestionService_ResubmitSuggestion_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *suggestionServiceClient) DeleteSuggestion(ctx context.Context, in *DeleteSuggestionInput, opts ...grpc.CallOption) (*DeleteSuggestionOutput, error) {
	out := new(DeleteSuggestionOutput)
	err := c.cc.Invoke(ctx, SuggestionService_DeleteSuggestion_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *suggestionServiceClient) ListSuggestions(ctx context.Context, in *ListSuggestionsInput, opts ...grpc.CallOption) (*ListSuggestionsOutput, error) {
	out := new(ListSuggestionsOutput)
	err := c.cc.Invoke(ctx, SuggestionService_ListSuggestions_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *suggestionServiceClient) DescribeSuggestion(ctx context.Context, in *DescribeSuggestionInput, opts ...grpc.CallOption) (*DescribeSuggestionOutput, error) {
	out := new(DescribeSuggestionOutput)
	err := c.cc.Invoke(ctx, SuggestionService_DescribeSuggestion_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SuggestionServiceServer is the server API for SuggestionService service.
// All implementations should embed UnimplementedSuggestionServiceServer
// for forward compatibility
type SuggestionServiceServer interface {
	CreateSuggestion(context.Context, *CreateSuggestionInput) (*CreateSuggestionOutput, error)
	UpdateSuggestion(context.Context, *UpdateSuggestionInput) (*UpdateSuggestionOutput, error)
	ReviewSuggestion(context.Context, *ReviewSuggestionInput) (*ReviewSuggestionOutput, error)
	ResubmitSuggestion(context.Context, *ResubmitSuggestionInput) (*ResubmitSuggestionOutput, error)
	DeleteSuggestion(context.Context, *DeleteSuggestionInput) (*DeleteSuggestionOutput, error)
	ListSuggestions(context.Context, *ListSuggestionsInput) (*ListSuggestionsOutput, error)
	DescribeSuggestion(context.Context, *DescribeSuggestionInput) (*DescribeSuggestionOutput, error)
}

// UnimplementedSuggestionServiceServer should be embedded to have forward compatible implementations.
type UnimplementedSuggestionServiceServer struct {
}

func (UnimplementedSuggestionServiceServer) CreateSuggestion(context.Context, *CreateSuggestionInput) (*CreateSuggestionOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateSuggestion not implemented")
}
func (UnimplementedSuggestionServiceServer) UpdateSuggestion(context.Context, *UpdateSuggestionInput) (*UpdateSuggestionOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateSuggestion not implemented")
}
func (UnimplementedSuggestionServiceServer) ReviewSuggestion(context.Context, *ReviewSuggestionInput) (*ReviewSuggestionOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReviewSuggestion not implemented")
}
func (UnimplementedSuggestionServiceServer) ResubmitSuggestion(context.Context, *ResubmitSuggestionInput) (*ResubmitSuggestionOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ResubmitSuggestion not implemented")
}
func (UnimplementedSuggestionServiceServer) DeleteSuggestion(context.Context, *DeleteSuggestionInput) (*DeleteSuggestionOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteSuggestion not implemented")
}
func (UnimplementedSuggestionServiceServer) ListSuggestions(context.Context, *ListSuggestionsInput) (*ListSuggestionsOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListSuggestions not implemented")
}
func (UnimplementedSuggestionServiceServer) DescribeSuggestion(context.Context, *DescribeSuggestionInput) (*DescribeSuggestionOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DescribeSuggestion not implemented")
}

// UnsafeSuggestionServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SuggestionServiceServer will
// result in compilation errors.
type UnsafeSuggestionServiceServer interface {
	mustEmbedUnimplementedSuggestionServiceServer()
}

func RegisterSuggestionServiceServer(s grpc.ServiceRegistrar, srv SuggestionServiceServer) {
	s.RegisterService(&SuggestionService_ServiceDesc, srv)
}

func _SuggestionService_CreateSuggestion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateSuggestionInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SuggestionServiceServer).CreateSuggestion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SuggestionService_CreateSuggestion_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SuggestionServiceServer).CreateSuggestion(ctx, req.(*CreateSuggestionInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _SuggestionService_UpdateSuggestion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateSuggestionInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SuggestionServiceServer).UpdateSuggestion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SuggestionService_UpdateSuggestion_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SuggestionServiceServer).UpdateSuggestion(ctx, req.(*UpdateSuggestionInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _SuggestionService_ReviewSuggestion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReviewSuggestionInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SuggestionServiceServer).ReviewSuggestion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SuggestionService_ReviewSuggestion_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SuggestionServiceServer).ReviewSuggestion(ctx, req.(*ReviewSuggestionInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _SuggestionService_ResubmitSuggestion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResubmitSuggestionInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SuggestionServiceServer).ResubmitSuggestion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SuggestionService_ResubmitSuggestion_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SuggestionServiceServer).ResubmitSuggestion(ctx, req.(*ResubmitSuggestionInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _SuggestionService_DeleteSuggestion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteSuggestionInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SuggestionServiceServer).DeleteSuggestion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SuggestionService_DeleteSuggestion_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SuggestionServiceServer).DeleteSuggestion(ctx, req.(*DeleteSuggestionInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _SuggestionService_ListSuggestions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListSuggestionsInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SuggestionServiceServer).ListSuggestions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SuggestionService_ListSuggestions_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SuggestionServiceServer).ListSuggestions(ctx, req.(*ListSuggestionsInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _SuggestionService_DescribeSuggestion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DescribeSuggestionInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SuggestionServiceServer).DescribeSuggestion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SuggestionService_DescribeSuggestion_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SuggestionServiceServer).DescribeSuggestion(ctx, req.(*DescribeSuggestionInput))
	}
	return interceptor(ctx, in, info, handler)
}

// SuggestionService_ServiceDesc is the grpc.ServiceDesc for SuggestionService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SuggestionService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "eolymp.atlas.SuggestionService",
	HandlerType: (*SuggestionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateSuggestion",
			Handler:    _SuggestionService_CreateSuggestion_Handler,
		},
		{
			MethodName: "UpdateSuggestion",
			Handler:    _SuggestionService_UpdateSuggestion_Handler,
		},
		{
			MethodName: "ReviewSuggestion",
			Handler:    _SuggestionService_ReviewSuggestion_Handler,
		},
		{
			MethodName: "ResubmitSuggestion",
			Handler:    _SuggestionService_ResubmitSuggestion_Handler,
		},
		{
			MethodName: "DeleteSuggestion",
			Handler:    _SuggestionService_DeleteSuggestion_Handler,
		},
		{
			MethodName: "ListSuggestions",
			Handler:    _SuggestionService_ListSuggestions_Handler,
		},
		{
			MethodName: "DescribeSuggestion",
			Handler:    _SuggestionService_DescribeSuggestion_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "eolymp/atlas/suggestion_service.proto",
}
