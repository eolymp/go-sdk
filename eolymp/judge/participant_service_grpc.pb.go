// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.21.12
// source: eolymp/judge/participant_service.proto

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
	ParticipantService_AddParticipant_FullMethodName        = "/eolymp.judge.ParticipantService/AddParticipant"
	ParticipantService_EnableParticipant_FullMethodName     = "/eolymp.judge.ParticipantService/EnableParticipant"
	ParticipantService_DisableParticipant_FullMethodName    = "/eolymp.judge.ParticipantService/DisableParticipant"
	ParticipantService_DisqualifyParticipant_FullMethodName = "/eolymp.judge.ParticipantService/DisqualifyParticipant"
	ParticipantService_UpdateParticipant_FullMethodName     = "/eolymp.judge.ParticipantService/UpdateParticipant"
	ParticipantService_RemoveParticipant_FullMethodName     = "/eolymp.judge.ParticipantService/RemoveParticipant"
	ParticipantService_ListParticipants_FullMethodName      = "/eolymp.judge.ParticipantService/ListParticipants"
	ParticipantService_DescribeParticipant_FullMethodName   = "/eolymp.judge.ParticipantService/DescribeParticipant"
	ParticipantService_IntrospectParticipant_FullMethodName = "/eolymp.judge.ParticipantService/IntrospectParticipant"
	ParticipantService_WatchParticipant_FullMethodName      = "/eolymp.judge.ParticipantService/WatchParticipant"
	ParticipantService_JoinContest_FullMethodName           = "/eolymp.judge.ParticipantService/JoinContest"
	ParticipantService_StartContest_FullMethodName          = "/eolymp.judge.ParticipantService/StartContest"
	ParticipantService_VerifyPasscode_FullMethodName        = "/eolymp.judge.ParticipantService/VerifyPasscode"
	ParticipantService_EnterPasscode_FullMethodName         = "/eolymp.judge.ParticipantService/EnterPasscode"
	ParticipantService_ResetPasscode_FullMethodName         = "/eolymp.judge.ParticipantService/ResetPasscode"
	ParticipantService_SetPasscode_FullMethodName           = "/eolymp.judge.ParticipantService/SetPasscode"
	ParticipantService_RemovePasscode_FullMethodName        = "/eolymp.judge.ParticipantService/RemovePasscode"
)

// ParticipantServiceClient is the client API for ParticipantService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ParticipantServiceClient interface {
	AddParticipant(ctx context.Context, in *AddParticipantInput, opts ...grpc.CallOption) (*AddParticipantOutput, error)
	EnableParticipant(ctx context.Context, in *EnableParticipantInput, opts ...grpc.CallOption) (*EnableParticipantOutput, error)
	DisableParticipant(ctx context.Context, in *DisableParticipantInput, opts ...grpc.CallOption) (*DisableParticipantOutput, error)
	DisqualifyParticipant(ctx context.Context, in *DisqualifyParticipantInput, opts ...grpc.CallOption) (*DisqualifyParticipantOutput, error)
	UpdateParticipant(ctx context.Context, in *UpdateParticipantInput, opts ...grpc.CallOption) (*UpdateParticipantOutput, error)
	RemoveParticipant(ctx context.Context, in *RemoveParticipantInput, opts ...grpc.CallOption) (*RemoveParticipantOutput, error)
	ListParticipants(ctx context.Context, in *ListParticipantsInput, opts ...grpc.CallOption) (*ListParticipantsOutput, error)
	DescribeParticipant(ctx context.Context, in *DescribeParticipantInput, opts ...grpc.CallOption) (*DescribeParticipantOutput, error)
	// IntrospectParticipant allows to fetch participant data for a currently authorized user.
	IntrospectParticipant(ctx context.Context, in *IntrospectParticipantInput, opts ...grpc.CallOption) (*IntrospectParticipantOutput, error)
	WatchParticipant(ctx context.Context, in *WatchParticipantInput, opts ...grpc.CallOption) (grpc.ServerStreamingClient[WatchParticipantOutput], error)
	// Allows a participant (currently authorized user) to join (add himself to) a public contest.
	// deprecated: use registration service instead
	JoinContest(ctx context.Context, in *JoinContestInput, opts ...grpc.CallOption) (*JoinContestOutput, error)
	// Allows a participant (currently authorized user) to start participating in the contest, see problems and submit solutions.
	StartContest(ctx context.Context, in *StartContestInput, opts ...grpc.CallOption) (*StartContestOutput, error)
	// Verify if passcode is required for the contest and if authenticated token has entered the passcode.
	VerifyPasscode(ctx context.Context, in *VerifyPasscodeInput, opts ...grpc.CallOption) (*VerifyPasscodeOutput, error)
	// Enter passcode marks current session as one authenticated by passcode.
	EnterPasscode(ctx context.Context, in *EnterPasscodeInput, opts ...grpc.CallOption) (*EnterPasscodeOutput, error)
	// Set a new passcode to the participant, if passcode was not set it will be now required
	ResetPasscode(ctx context.Context, in *ResetPasscodeInput, opts ...grpc.CallOption) (*ResetPasscodeOutput, error)
	// Set a new passcode to the participant, if passcode was not set it will be now required
	SetPasscode(ctx context.Context, in *SetPasscodeInput, opts ...grpc.CallOption) (*SetPasscodeOutput, error)
	// Remove passcode from participant and allow her to enter contest without passcode.
	RemovePasscode(ctx context.Context, in *RemovePasscodeInput, opts ...grpc.CallOption) (*RemovePasscodeOutput, error)
}

type participantServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewParticipantServiceClient(cc grpc.ClientConnInterface) ParticipantServiceClient {
	return &participantServiceClient{cc}
}

func (c *participantServiceClient) AddParticipant(ctx context.Context, in *AddParticipantInput, opts ...grpc.CallOption) (*AddParticipantOutput, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AddParticipantOutput)
	err := c.cc.Invoke(ctx, ParticipantService_AddParticipant_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *participantServiceClient) EnableParticipant(ctx context.Context, in *EnableParticipantInput, opts ...grpc.CallOption) (*EnableParticipantOutput, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(EnableParticipantOutput)
	err := c.cc.Invoke(ctx, ParticipantService_EnableParticipant_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *participantServiceClient) DisableParticipant(ctx context.Context, in *DisableParticipantInput, opts ...grpc.CallOption) (*DisableParticipantOutput, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DisableParticipantOutput)
	err := c.cc.Invoke(ctx, ParticipantService_DisableParticipant_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *participantServiceClient) DisqualifyParticipant(ctx context.Context, in *DisqualifyParticipantInput, opts ...grpc.CallOption) (*DisqualifyParticipantOutput, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DisqualifyParticipantOutput)
	err := c.cc.Invoke(ctx, ParticipantService_DisqualifyParticipant_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *participantServiceClient) UpdateParticipant(ctx context.Context, in *UpdateParticipantInput, opts ...grpc.CallOption) (*UpdateParticipantOutput, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateParticipantOutput)
	err := c.cc.Invoke(ctx, ParticipantService_UpdateParticipant_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *participantServiceClient) RemoveParticipant(ctx context.Context, in *RemoveParticipantInput, opts ...grpc.CallOption) (*RemoveParticipantOutput, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RemoveParticipantOutput)
	err := c.cc.Invoke(ctx, ParticipantService_RemoveParticipant_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *participantServiceClient) ListParticipants(ctx context.Context, in *ListParticipantsInput, opts ...grpc.CallOption) (*ListParticipantsOutput, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListParticipantsOutput)
	err := c.cc.Invoke(ctx, ParticipantService_ListParticipants_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *participantServiceClient) DescribeParticipant(ctx context.Context, in *DescribeParticipantInput, opts ...grpc.CallOption) (*DescribeParticipantOutput, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DescribeParticipantOutput)
	err := c.cc.Invoke(ctx, ParticipantService_DescribeParticipant_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *participantServiceClient) IntrospectParticipant(ctx context.Context, in *IntrospectParticipantInput, opts ...grpc.CallOption) (*IntrospectParticipantOutput, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(IntrospectParticipantOutput)
	err := c.cc.Invoke(ctx, ParticipantService_IntrospectParticipant_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *participantServiceClient) WatchParticipant(ctx context.Context, in *WatchParticipantInput, opts ...grpc.CallOption) (grpc.ServerStreamingClient[WatchParticipantOutput], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &ParticipantService_ServiceDesc.Streams[0], ParticipantService_WatchParticipant_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[WatchParticipantInput, WatchParticipantOutput]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type ParticipantService_WatchParticipantClient = grpc.ServerStreamingClient[WatchParticipantOutput]

func (c *participantServiceClient) JoinContest(ctx context.Context, in *JoinContestInput, opts ...grpc.CallOption) (*JoinContestOutput, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(JoinContestOutput)
	err := c.cc.Invoke(ctx, ParticipantService_JoinContest_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *participantServiceClient) StartContest(ctx context.Context, in *StartContestInput, opts ...grpc.CallOption) (*StartContestOutput, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(StartContestOutput)
	err := c.cc.Invoke(ctx, ParticipantService_StartContest_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *participantServiceClient) VerifyPasscode(ctx context.Context, in *VerifyPasscodeInput, opts ...grpc.CallOption) (*VerifyPasscodeOutput, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(VerifyPasscodeOutput)
	err := c.cc.Invoke(ctx, ParticipantService_VerifyPasscode_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *participantServiceClient) EnterPasscode(ctx context.Context, in *EnterPasscodeInput, opts ...grpc.CallOption) (*EnterPasscodeOutput, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(EnterPasscodeOutput)
	err := c.cc.Invoke(ctx, ParticipantService_EnterPasscode_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *participantServiceClient) ResetPasscode(ctx context.Context, in *ResetPasscodeInput, opts ...grpc.CallOption) (*ResetPasscodeOutput, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ResetPasscodeOutput)
	err := c.cc.Invoke(ctx, ParticipantService_ResetPasscode_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *participantServiceClient) SetPasscode(ctx context.Context, in *SetPasscodeInput, opts ...grpc.CallOption) (*SetPasscodeOutput, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SetPasscodeOutput)
	err := c.cc.Invoke(ctx, ParticipantService_SetPasscode_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *participantServiceClient) RemovePasscode(ctx context.Context, in *RemovePasscodeInput, opts ...grpc.CallOption) (*RemovePasscodeOutput, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RemovePasscodeOutput)
	err := c.cc.Invoke(ctx, ParticipantService_RemovePasscode_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ParticipantServiceServer is the server API for ParticipantService service.
// All implementations should embed UnimplementedParticipantServiceServer
// for forward compatibility.
type ParticipantServiceServer interface {
	AddParticipant(context.Context, *AddParticipantInput) (*AddParticipantOutput, error)
	EnableParticipant(context.Context, *EnableParticipantInput) (*EnableParticipantOutput, error)
	DisableParticipant(context.Context, *DisableParticipantInput) (*DisableParticipantOutput, error)
	DisqualifyParticipant(context.Context, *DisqualifyParticipantInput) (*DisqualifyParticipantOutput, error)
	UpdateParticipant(context.Context, *UpdateParticipantInput) (*UpdateParticipantOutput, error)
	RemoveParticipant(context.Context, *RemoveParticipantInput) (*RemoveParticipantOutput, error)
	ListParticipants(context.Context, *ListParticipantsInput) (*ListParticipantsOutput, error)
	DescribeParticipant(context.Context, *DescribeParticipantInput) (*DescribeParticipantOutput, error)
	// IntrospectParticipant allows to fetch participant data for a currently authorized user.
	IntrospectParticipant(context.Context, *IntrospectParticipantInput) (*IntrospectParticipantOutput, error)
	WatchParticipant(*WatchParticipantInput, grpc.ServerStreamingServer[WatchParticipantOutput]) error
	// Allows a participant (currently authorized user) to join (add himself to) a public contest.
	// deprecated: use registration service instead
	JoinContest(context.Context, *JoinContestInput) (*JoinContestOutput, error)
	// Allows a participant (currently authorized user) to start participating in the contest, see problems and submit solutions.
	StartContest(context.Context, *StartContestInput) (*StartContestOutput, error)
	// Verify if passcode is required for the contest and if authenticated token has entered the passcode.
	VerifyPasscode(context.Context, *VerifyPasscodeInput) (*VerifyPasscodeOutput, error)
	// Enter passcode marks current session as one authenticated by passcode.
	EnterPasscode(context.Context, *EnterPasscodeInput) (*EnterPasscodeOutput, error)
	// Set a new passcode to the participant, if passcode was not set it will be now required
	ResetPasscode(context.Context, *ResetPasscodeInput) (*ResetPasscodeOutput, error)
	// Set a new passcode to the participant, if passcode was not set it will be now required
	SetPasscode(context.Context, *SetPasscodeInput) (*SetPasscodeOutput, error)
	// Remove passcode from participant and allow her to enter contest without passcode.
	RemovePasscode(context.Context, *RemovePasscodeInput) (*RemovePasscodeOutput, error)
}

// UnimplementedParticipantServiceServer should be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedParticipantServiceServer struct{}

func (UnimplementedParticipantServiceServer) AddParticipant(context.Context, *AddParticipantInput) (*AddParticipantOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddParticipant not implemented")
}
func (UnimplementedParticipantServiceServer) EnableParticipant(context.Context, *EnableParticipantInput) (*EnableParticipantOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EnableParticipant not implemented")
}
func (UnimplementedParticipantServiceServer) DisableParticipant(context.Context, *DisableParticipantInput) (*DisableParticipantOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DisableParticipant not implemented")
}
func (UnimplementedParticipantServiceServer) DisqualifyParticipant(context.Context, *DisqualifyParticipantInput) (*DisqualifyParticipantOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DisqualifyParticipant not implemented")
}
func (UnimplementedParticipantServiceServer) UpdateParticipant(context.Context, *UpdateParticipantInput) (*UpdateParticipantOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateParticipant not implemented")
}
func (UnimplementedParticipantServiceServer) RemoveParticipant(context.Context, *RemoveParticipantInput) (*RemoveParticipantOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveParticipant not implemented")
}
func (UnimplementedParticipantServiceServer) ListParticipants(context.Context, *ListParticipantsInput) (*ListParticipantsOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListParticipants not implemented")
}
func (UnimplementedParticipantServiceServer) DescribeParticipant(context.Context, *DescribeParticipantInput) (*DescribeParticipantOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DescribeParticipant not implemented")
}
func (UnimplementedParticipantServiceServer) IntrospectParticipant(context.Context, *IntrospectParticipantInput) (*IntrospectParticipantOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IntrospectParticipant not implemented")
}
func (UnimplementedParticipantServiceServer) WatchParticipant(*WatchParticipantInput, grpc.ServerStreamingServer[WatchParticipantOutput]) error {
	return status.Errorf(codes.Unimplemented, "method WatchParticipant not implemented")
}
func (UnimplementedParticipantServiceServer) JoinContest(context.Context, *JoinContestInput) (*JoinContestOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method JoinContest not implemented")
}
func (UnimplementedParticipantServiceServer) StartContest(context.Context, *StartContestInput) (*StartContestOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StartContest not implemented")
}
func (UnimplementedParticipantServiceServer) VerifyPasscode(context.Context, *VerifyPasscodeInput) (*VerifyPasscodeOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VerifyPasscode not implemented")
}
func (UnimplementedParticipantServiceServer) EnterPasscode(context.Context, *EnterPasscodeInput) (*EnterPasscodeOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EnterPasscode not implemented")
}
func (UnimplementedParticipantServiceServer) ResetPasscode(context.Context, *ResetPasscodeInput) (*ResetPasscodeOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ResetPasscode not implemented")
}
func (UnimplementedParticipantServiceServer) SetPasscode(context.Context, *SetPasscodeInput) (*SetPasscodeOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetPasscode not implemented")
}
func (UnimplementedParticipantServiceServer) RemovePasscode(context.Context, *RemovePasscodeInput) (*RemovePasscodeOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemovePasscode not implemented")
}
func (UnimplementedParticipantServiceServer) testEmbeddedByValue() {}

// UnsafeParticipantServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ParticipantServiceServer will
// result in compilation errors.
type UnsafeParticipantServiceServer interface {
	mustEmbedUnimplementedParticipantServiceServer()
}

func RegisterParticipantServiceServer(s grpc.ServiceRegistrar, srv ParticipantServiceServer) {
	// If the following call pancis, it indicates UnimplementedParticipantServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&ParticipantService_ServiceDesc, srv)
}

func _ParticipantService_AddParticipant_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddParticipantInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ParticipantServiceServer).AddParticipant(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ParticipantService_AddParticipant_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ParticipantServiceServer).AddParticipant(ctx, req.(*AddParticipantInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _ParticipantService_EnableParticipant_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EnableParticipantInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ParticipantServiceServer).EnableParticipant(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ParticipantService_EnableParticipant_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ParticipantServiceServer).EnableParticipant(ctx, req.(*EnableParticipantInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _ParticipantService_DisableParticipant_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DisableParticipantInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ParticipantServiceServer).DisableParticipant(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ParticipantService_DisableParticipant_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ParticipantServiceServer).DisableParticipant(ctx, req.(*DisableParticipantInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _ParticipantService_DisqualifyParticipant_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DisqualifyParticipantInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ParticipantServiceServer).DisqualifyParticipant(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ParticipantService_DisqualifyParticipant_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ParticipantServiceServer).DisqualifyParticipant(ctx, req.(*DisqualifyParticipantInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _ParticipantService_UpdateParticipant_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateParticipantInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ParticipantServiceServer).UpdateParticipant(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ParticipantService_UpdateParticipant_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ParticipantServiceServer).UpdateParticipant(ctx, req.(*UpdateParticipantInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _ParticipantService_RemoveParticipant_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveParticipantInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ParticipantServiceServer).RemoveParticipant(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ParticipantService_RemoveParticipant_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ParticipantServiceServer).RemoveParticipant(ctx, req.(*RemoveParticipantInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _ParticipantService_ListParticipants_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListParticipantsInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ParticipantServiceServer).ListParticipants(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ParticipantService_ListParticipants_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ParticipantServiceServer).ListParticipants(ctx, req.(*ListParticipantsInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _ParticipantService_DescribeParticipant_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DescribeParticipantInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ParticipantServiceServer).DescribeParticipant(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ParticipantService_DescribeParticipant_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ParticipantServiceServer).DescribeParticipant(ctx, req.(*DescribeParticipantInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _ParticipantService_IntrospectParticipant_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IntrospectParticipantInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ParticipantServiceServer).IntrospectParticipant(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ParticipantService_IntrospectParticipant_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ParticipantServiceServer).IntrospectParticipant(ctx, req.(*IntrospectParticipantInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _ParticipantService_WatchParticipant_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(WatchParticipantInput)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ParticipantServiceServer).WatchParticipant(m, &grpc.GenericServerStream[WatchParticipantInput, WatchParticipantOutput]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type ParticipantService_WatchParticipantServer = grpc.ServerStreamingServer[WatchParticipantOutput]

func _ParticipantService_JoinContest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JoinContestInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ParticipantServiceServer).JoinContest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ParticipantService_JoinContest_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ParticipantServiceServer).JoinContest(ctx, req.(*JoinContestInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _ParticipantService_StartContest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartContestInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ParticipantServiceServer).StartContest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ParticipantService_StartContest_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ParticipantServiceServer).StartContest(ctx, req.(*StartContestInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _ParticipantService_VerifyPasscode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VerifyPasscodeInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ParticipantServiceServer).VerifyPasscode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ParticipantService_VerifyPasscode_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ParticipantServiceServer).VerifyPasscode(ctx, req.(*VerifyPasscodeInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _ParticipantService_EnterPasscode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EnterPasscodeInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ParticipantServiceServer).EnterPasscode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ParticipantService_EnterPasscode_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ParticipantServiceServer).EnterPasscode(ctx, req.(*EnterPasscodeInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _ParticipantService_ResetPasscode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResetPasscodeInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ParticipantServiceServer).ResetPasscode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ParticipantService_ResetPasscode_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ParticipantServiceServer).ResetPasscode(ctx, req.(*ResetPasscodeInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _ParticipantService_SetPasscode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetPasscodeInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ParticipantServiceServer).SetPasscode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ParticipantService_SetPasscode_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ParticipantServiceServer).SetPasscode(ctx, req.(*SetPasscodeInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _ParticipantService_RemovePasscode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemovePasscodeInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ParticipantServiceServer).RemovePasscode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ParticipantService_RemovePasscode_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ParticipantServiceServer).RemovePasscode(ctx, req.(*RemovePasscodeInput))
	}
	return interceptor(ctx, in, info, handler)
}

// ParticipantService_ServiceDesc is the grpc.ServiceDesc for ParticipantService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ParticipantService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "eolymp.judge.ParticipantService",
	HandlerType: (*ParticipantServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddParticipant",
			Handler:    _ParticipantService_AddParticipant_Handler,
		},
		{
			MethodName: "EnableParticipant",
			Handler:    _ParticipantService_EnableParticipant_Handler,
		},
		{
			MethodName: "DisableParticipant",
			Handler:    _ParticipantService_DisableParticipant_Handler,
		},
		{
			MethodName: "DisqualifyParticipant",
			Handler:    _ParticipantService_DisqualifyParticipant_Handler,
		},
		{
			MethodName: "UpdateParticipant",
			Handler:    _ParticipantService_UpdateParticipant_Handler,
		},
		{
			MethodName: "RemoveParticipant",
			Handler:    _ParticipantService_RemoveParticipant_Handler,
		},
		{
			MethodName: "ListParticipants",
			Handler:    _ParticipantService_ListParticipants_Handler,
		},
		{
			MethodName: "DescribeParticipant",
			Handler:    _ParticipantService_DescribeParticipant_Handler,
		},
		{
			MethodName: "IntrospectParticipant",
			Handler:    _ParticipantService_IntrospectParticipant_Handler,
		},
		{
			MethodName: "JoinContest",
			Handler:    _ParticipantService_JoinContest_Handler,
		},
		{
			MethodName: "StartContest",
			Handler:    _ParticipantService_StartContest_Handler,
		},
		{
			MethodName: "VerifyPasscode",
			Handler:    _ParticipantService_VerifyPasscode_Handler,
		},
		{
			MethodName: "EnterPasscode",
			Handler:    _ParticipantService_EnterPasscode_Handler,
		},
		{
			MethodName: "ResetPasscode",
			Handler:    _ParticipantService_ResetPasscode_Handler,
		},
		{
			MethodName: "SetPasscode",
			Handler:    _ParticipantService_SetPasscode_Handler,
		},
		{
			MethodName: "RemovePasscode",
			Handler:    _ParticipantService_RemovePasscode_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "WatchParticipant",
			Handler:       _ParticipantService_WatchParticipant_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "eolymp/judge/participant_service.proto",
}
