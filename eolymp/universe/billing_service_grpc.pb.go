// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.21.12
// source: eolymp/universe/billing_service.proto

package universe

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
	BillingService_DescribeBillingInformation_FullMethodName = "/eolymp.universe.BillingService/DescribeBillingInformation"
	BillingService_UpdateBillingInformation_FullMethodName   = "/eolymp.universe.BillingService/UpdateBillingInformation"
	BillingService_DescribeSubscription_FullMethodName       = "/eolymp.universe.BillingService/DescribeSubscription"
	BillingService_CreateSubscription_FullMethodName         = "/eolymp.universe.BillingService/CreateSubscription"
	BillingService_UpdateSubscription_FullMethodName         = "/eolymp.universe.BillingService/UpdateSubscription"
	BillingService_CancelSubscription_FullMethodName         = "/eolymp.universe.BillingService/CancelSubscription"
	BillingService_StartSubscriptionTrial_FullMethodName     = "/eolymp.universe.BillingService/StartSubscriptionTrial"
	BillingService_EndSubscriptionTrial_FullMethodName       = "/eolymp.universe.BillingService/EndSubscriptionTrial"
	BillingService_SimulateSubscription_FullMethodName       = "/eolymp.universe.BillingService/SimulateSubscription"
	BillingService_CreatePortalLink_FullMethodName           = "/eolymp.universe.BillingService/CreatePortalLink"
	BillingService_UpcomingInvoice_FullMethodName            = "/eolymp.universe.BillingService/UpcomingInvoice"
	BillingService_DescribeInvoice_FullMethodName            = "/eolymp.universe.BillingService/DescribeInvoice"
	BillingService_PayInvoice_FullMethodName                 = "/eolymp.universe.BillingService/PayInvoice"
	BillingService_ListInvoices_FullMethodName               = "/eolymp.universe.BillingService/ListInvoices"
)

// BillingServiceClient is the client API for BillingService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BillingServiceClient interface {
	DescribeBillingInformation(ctx context.Context, in *DescribeBillingInformationInput, opts ...grpc.CallOption) (*DescribeBillingInformationOutput, error)
	UpdateBillingInformation(ctx context.Context, in *UpdateBillingInformationInput, opts ...grpc.CallOption) (*UpdateBillingInformationOutput, error)
	DescribeSubscription(ctx context.Context, in *DescribeSubscriptionInput, opts ...grpc.CallOption) (*DescribeSubscriptionOutput, error)
	CreateSubscription(ctx context.Context, in *CreateSubscriptionInput, opts ...grpc.CallOption) (*CreateSubscriptionOutput, error)
	UpdateSubscription(ctx context.Context, in *UpdateSubscriptionInput, opts ...grpc.CallOption) (*UpdateSubscriptionOutput, error)
	CancelSubscription(ctx context.Context, in *CancelSubscriptionInput, opts ...grpc.CallOption) (*CancelSubscriptionOutput, error)
	StartSubscriptionTrial(ctx context.Context, in *StartSubscriptionTrialInput, opts ...grpc.CallOption) (*StartSubscriptionTrialOutput, error)
	EndSubscriptionTrial(ctx context.Context, in *EndSubscriptionTrialInput, opts ...grpc.CallOption) (*EndSubscriptionTrialOutput, error)
	SimulateSubscription(ctx context.Context, in *SimulateSubscriptionInput, opts ...grpc.CallOption) (*SimulateSubscriptionOutput, error)
	CreatePortalLink(ctx context.Context, in *CreatePortalLinkInput, opts ...grpc.CallOption) (*CreatePortalLinkOutput, error)
	UpcomingInvoice(ctx context.Context, in *UpcomingInvoiceInput, opts ...grpc.CallOption) (*UpcomingInvoiceOutput, error)
	DescribeInvoice(ctx context.Context, in *DescribeInvoiceInput, opts ...grpc.CallOption) (*DescribeInvoiceOutput, error)
	PayInvoice(ctx context.Context, in *PayInvoiceInput, opts ...grpc.CallOption) (*PayInvoiceOutput, error)
	ListInvoices(ctx context.Context, in *ListInvoicesInput, opts ...grpc.CallOption) (*ListInvoicesOutput, error)
}

type billingServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBillingServiceClient(cc grpc.ClientConnInterface) BillingServiceClient {
	return &billingServiceClient{cc}
}

func (c *billingServiceClient) DescribeBillingInformation(ctx context.Context, in *DescribeBillingInformationInput, opts ...grpc.CallOption) (*DescribeBillingInformationOutput, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DescribeBillingInformationOutput)
	err := c.cc.Invoke(ctx, BillingService_DescribeBillingInformation_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *billingServiceClient) UpdateBillingInformation(ctx context.Context, in *UpdateBillingInformationInput, opts ...grpc.CallOption) (*UpdateBillingInformationOutput, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateBillingInformationOutput)
	err := c.cc.Invoke(ctx, BillingService_UpdateBillingInformation_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *billingServiceClient) DescribeSubscription(ctx context.Context, in *DescribeSubscriptionInput, opts ...grpc.CallOption) (*DescribeSubscriptionOutput, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DescribeSubscriptionOutput)
	err := c.cc.Invoke(ctx, BillingService_DescribeSubscription_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *billingServiceClient) CreateSubscription(ctx context.Context, in *CreateSubscriptionInput, opts ...grpc.CallOption) (*CreateSubscriptionOutput, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateSubscriptionOutput)
	err := c.cc.Invoke(ctx, BillingService_CreateSubscription_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *billingServiceClient) UpdateSubscription(ctx context.Context, in *UpdateSubscriptionInput, opts ...grpc.CallOption) (*UpdateSubscriptionOutput, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateSubscriptionOutput)
	err := c.cc.Invoke(ctx, BillingService_UpdateSubscription_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *billingServiceClient) CancelSubscription(ctx context.Context, in *CancelSubscriptionInput, opts ...grpc.CallOption) (*CancelSubscriptionOutput, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CancelSubscriptionOutput)
	err := c.cc.Invoke(ctx, BillingService_CancelSubscription_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *billingServiceClient) StartSubscriptionTrial(ctx context.Context, in *StartSubscriptionTrialInput, opts ...grpc.CallOption) (*StartSubscriptionTrialOutput, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(StartSubscriptionTrialOutput)
	err := c.cc.Invoke(ctx, BillingService_StartSubscriptionTrial_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *billingServiceClient) EndSubscriptionTrial(ctx context.Context, in *EndSubscriptionTrialInput, opts ...grpc.CallOption) (*EndSubscriptionTrialOutput, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(EndSubscriptionTrialOutput)
	err := c.cc.Invoke(ctx, BillingService_EndSubscriptionTrial_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *billingServiceClient) SimulateSubscription(ctx context.Context, in *SimulateSubscriptionInput, opts ...grpc.CallOption) (*SimulateSubscriptionOutput, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SimulateSubscriptionOutput)
	err := c.cc.Invoke(ctx, BillingService_SimulateSubscription_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *billingServiceClient) CreatePortalLink(ctx context.Context, in *CreatePortalLinkInput, opts ...grpc.CallOption) (*CreatePortalLinkOutput, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreatePortalLinkOutput)
	err := c.cc.Invoke(ctx, BillingService_CreatePortalLink_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *billingServiceClient) UpcomingInvoice(ctx context.Context, in *UpcomingInvoiceInput, opts ...grpc.CallOption) (*UpcomingInvoiceOutput, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpcomingInvoiceOutput)
	err := c.cc.Invoke(ctx, BillingService_UpcomingInvoice_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *billingServiceClient) DescribeInvoice(ctx context.Context, in *DescribeInvoiceInput, opts ...grpc.CallOption) (*DescribeInvoiceOutput, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DescribeInvoiceOutput)
	err := c.cc.Invoke(ctx, BillingService_DescribeInvoice_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *billingServiceClient) PayInvoice(ctx context.Context, in *PayInvoiceInput, opts ...grpc.CallOption) (*PayInvoiceOutput, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PayInvoiceOutput)
	err := c.cc.Invoke(ctx, BillingService_PayInvoice_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *billingServiceClient) ListInvoices(ctx context.Context, in *ListInvoicesInput, opts ...grpc.CallOption) (*ListInvoicesOutput, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListInvoicesOutput)
	err := c.cc.Invoke(ctx, BillingService_ListInvoices_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BillingServiceServer is the server API for BillingService service.
// All implementations should embed UnimplementedBillingServiceServer
// for forward compatibility.
type BillingServiceServer interface {
	DescribeBillingInformation(context.Context, *DescribeBillingInformationInput) (*DescribeBillingInformationOutput, error)
	UpdateBillingInformation(context.Context, *UpdateBillingInformationInput) (*UpdateBillingInformationOutput, error)
	DescribeSubscription(context.Context, *DescribeSubscriptionInput) (*DescribeSubscriptionOutput, error)
	CreateSubscription(context.Context, *CreateSubscriptionInput) (*CreateSubscriptionOutput, error)
	UpdateSubscription(context.Context, *UpdateSubscriptionInput) (*UpdateSubscriptionOutput, error)
	CancelSubscription(context.Context, *CancelSubscriptionInput) (*CancelSubscriptionOutput, error)
	StartSubscriptionTrial(context.Context, *StartSubscriptionTrialInput) (*StartSubscriptionTrialOutput, error)
	EndSubscriptionTrial(context.Context, *EndSubscriptionTrialInput) (*EndSubscriptionTrialOutput, error)
	SimulateSubscription(context.Context, *SimulateSubscriptionInput) (*SimulateSubscriptionOutput, error)
	CreatePortalLink(context.Context, *CreatePortalLinkInput) (*CreatePortalLinkOutput, error)
	UpcomingInvoice(context.Context, *UpcomingInvoiceInput) (*UpcomingInvoiceOutput, error)
	DescribeInvoice(context.Context, *DescribeInvoiceInput) (*DescribeInvoiceOutput, error)
	PayInvoice(context.Context, *PayInvoiceInput) (*PayInvoiceOutput, error)
	ListInvoices(context.Context, *ListInvoicesInput) (*ListInvoicesOutput, error)
}

// UnimplementedBillingServiceServer should be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedBillingServiceServer struct{}

func (UnimplementedBillingServiceServer) DescribeBillingInformation(context.Context, *DescribeBillingInformationInput) (*DescribeBillingInformationOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DescribeBillingInformation not implemented")
}
func (UnimplementedBillingServiceServer) UpdateBillingInformation(context.Context, *UpdateBillingInformationInput) (*UpdateBillingInformationOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateBillingInformation not implemented")
}
func (UnimplementedBillingServiceServer) DescribeSubscription(context.Context, *DescribeSubscriptionInput) (*DescribeSubscriptionOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DescribeSubscription not implemented")
}
func (UnimplementedBillingServiceServer) CreateSubscription(context.Context, *CreateSubscriptionInput) (*CreateSubscriptionOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateSubscription not implemented")
}
func (UnimplementedBillingServiceServer) UpdateSubscription(context.Context, *UpdateSubscriptionInput) (*UpdateSubscriptionOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateSubscription not implemented")
}
func (UnimplementedBillingServiceServer) CancelSubscription(context.Context, *CancelSubscriptionInput) (*CancelSubscriptionOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CancelSubscription not implemented")
}
func (UnimplementedBillingServiceServer) StartSubscriptionTrial(context.Context, *StartSubscriptionTrialInput) (*StartSubscriptionTrialOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StartSubscriptionTrial not implemented")
}
func (UnimplementedBillingServiceServer) EndSubscriptionTrial(context.Context, *EndSubscriptionTrialInput) (*EndSubscriptionTrialOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EndSubscriptionTrial not implemented")
}
func (UnimplementedBillingServiceServer) SimulateSubscription(context.Context, *SimulateSubscriptionInput) (*SimulateSubscriptionOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SimulateSubscription not implemented")
}
func (UnimplementedBillingServiceServer) CreatePortalLink(context.Context, *CreatePortalLinkInput) (*CreatePortalLinkOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePortalLink not implemented")
}
func (UnimplementedBillingServiceServer) UpcomingInvoice(context.Context, *UpcomingInvoiceInput) (*UpcomingInvoiceOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpcomingInvoice not implemented")
}
func (UnimplementedBillingServiceServer) DescribeInvoice(context.Context, *DescribeInvoiceInput) (*DescribeInvoiceOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DescribeInvoice not implemented")
}
func (UnimplementedBillingServiceServer) PayInvoice(context.Context, *PayInvoiceInput) (*PayInvoiceOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PayInvoice not implemented")
}
func (UnimplementedBillingServiceServer) ListInvoices(context.Context, *ListInvoicesInput) (*ListInvoicesOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListInvoices not implemented")
}
func (UnimplementedBillingServiceServer) testEmbeddedByValue() {}

// UnsafeBillingServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BillingServiceServer will
// result in compilation errors.
type UnsafeBillingServiceServer interface {
	mustEmbedUnimplementedBillingServiceServer()
}

func RegisterBillingServiceServer(s grpc.ServiceRegistrar, srv BillingServiceServer) {
	// If the following call pancis, it indicates UnimplementedBillingServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&BillingService_ServiceDesc, srv)
}

func _BillingService_DescribeBillingInformation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DescribeBillingInformationInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BillingServiceServer).DescribeBillingInformation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BillingService_DescribeBillingInformation_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BillingServiceServer).DescribeBillingInformation(ctx, req.(*DescribeBillingInformationInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _BillingService_UpdateBillingInformation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateBillingInformationInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BillingServiceServer).UpdateBillingInformation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BillingService_UpdateBillingInformation_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BillingServiceServer).UpdateBillingInformation(ctx, req.(*UpdateBillingInformationInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _BillingService_DescribeSubscription_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DescribeSubscriptionInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BillingServiceServer).DescribeSubscription(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BillingService_DescribeSubscription_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BillingServiceServer).DescribeSubscription(ctx, req.(*DescribeSubscriptionInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _BillingService_CreateSubscription_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateSubscriptionInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BillingServiceServer).CreateSubscription(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BillingService_CreateSubscription_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BillingServiceServer).CreateSubscription(ctx, req.(*CreateSubscriptionInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _BillingService_UpdateSubscription_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateSubscriptionInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BillingServiceServer).UpdateSubscription(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BillingService_UpdateSubscription_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BillingServiceServer).UpdateSubscription(ctx, req.(*UpdateSubscriptionInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _BillingService_CancelSubscription_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CancelSubscriptionInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BillingServiceServer).CancelSubscription(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BillingService_CancelSubscription_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BillingServiceServer).CancelSubscription(ctx, req.(*CancelSubscriptionInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _BillingService_StartSubscriptionTrial_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartSubscriptionTrialInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BillingServiceServer).StartSubscriptionTrial(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BillingService_StartSubscriptionTrial_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BillingServiceServer).StartSubscriptionTrial(ctx, req.(*StartSubscriptionTrialInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _BillingService_EndSubscriptionTrial_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EndSubscriptionTrialInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BillingServiceServer).EndSubscriptionTrial(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BillingService_EndSubscriptionTrial_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BillingServiceServer).EndSubscriptionTrial(ctx, req.(*EndSubscriptionTrialInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _BillingService_SimulateSubscription_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SimulateSubscriptionInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BillingServiceServer).SimulateSubscription(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BillingService_SimulateSubscription_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BillingServiceServer).SimulateSubscription(ctx, req.(*SimulateSubscriptionInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _BillingService_CreatePortalLink_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePortalLinkInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BillingServiceServer).CreatePortalLink(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BillingService_CreatePortalLink_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BillingServiceServer).CreatePortalLink(ctx, req.(*CreatePortalLinkInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _BillingService_UpcomingInvoice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpcomingInvoiceInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BillingServiceServer).UpcomingInvoice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BillingService_UpcomingInvoice_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BillingServiceServer).UpcomingInvoice(ctx, req.(*UpcomingInvoiceInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _BillingService_DescribeInvoice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DescribeInvoiceInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BillingServiceServer).DescribeInvoice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BillingService_DescribeInvoice_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BillingServiceServer).DescribeInvoice(ctx, req.(*DescribeInvoiceInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _BillingService_PayInvoice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PayInvoiceInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BillingServiceServer).PayInvoice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BillingService_PayInvoice_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BillingServiceServer).PayInvoice(ctx, req.(*PayInvoiceInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _BillingService_ListInvoices_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListInvoicesInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BillingServiceServer).ListInvoices(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BillingService_ListInvoices_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BillingServiceServer).ListInvoices(ctx, req.(*ListInvoicesInput))
	}
	return interceptor(ctx, in, info, handler)
}

// BillingService_ServiceDesc is the grpc.ServiceDesc for BillingService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BillingService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "eolymp.universe.BillingService",
	HandlerType: (*BillingServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DescribeBillingInformation",
			Handler:    _BillingService_DescribeBillingInformation_Handler,
		},
		{
			MethodName: "UpdateBillingInformation",
			Handler:    _BillingService_UpdateBillingInformation_Handler,
		},
		{
			MethodName: "DescribeSubscription",
			Handler:    _BillingService_DescribeSubscription_Handler,
		},
		{
			MethodName: "CreateSubscription",
			Handler:    _BillingService_CreateSubscription_Handler,
		},
		{
			MethodName: "UpdateSubscription",
			Handler:    _BillingService_UpdateSubscription_Handler,
		},
		{
			MethodName: "CancelSubscription",
			Handler:    _BillingService_CancelSubscription_Handler,
		},
		{
			MethodName: "StartSubscriptionTrial",
			Handler:    _BillingService_StartSubscriptionTrial_Handler,
		},
		{
			MethodName: "EndSubscriptionTrial",
			Handler:    _BillingService_EndSubscriptionTrial_Handler,
		},
		{
			MethodName: "SimulateSubscription",
			Handler:    _BillingService_SimulateSubscription_Handler,
		},
		{
			MethodName: "CreatePortalLink",
			Handler:    _BillingService_CreatePortalLink_Handler,
		},
		{
			MethodName: "UpcomingInvoice",
			Handler:    _BillingService_UpcomingInvoice_Handler,
		},
		{
			MethodName: "DescribeInvoice",
			Handler:    _BillingService_DescribeInvoice_Handler,
		},
		{
			MethodName: "PayInvoice",
			Handler:    _BillingService_PayInvoice_Handler,
		},
		{
			MethodName: "ListInvoices",
			Handler:    _BillingService_ListInvoices_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "eolymp/universe/billing_service.proto",
}
