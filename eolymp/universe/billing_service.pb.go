// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.18.1
// source: eolymp/universe/billing_service.proto

package universe

import (
	_ "github.com/eolymp/go-sdk/eolymp/annotations"
	commerce "github.com/eolymp/go-sdk/eolymp/commerce"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type DescribeBillingInformationInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DescribeBillingInformationInput) Reset() {
	*x = DescribeBillingInformationInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_universe_billing_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DescribeBillingInformationInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescribeBillingInformationInput) ProtoMessage() {}

func (x *DescribeBillingInformationInput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_universe_billing_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DescribeBillingInformationInput.ProtoReflect.Descriptor instead.
func (*DescribeBillingInformationInput) Descriptor() ([]byte, []int) {
	return file_eolymp_universe_billing_service_proto_rawDescGZIP(), []int{0}
}

type DescribeBillingInformationOutput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *Billing_Information `protobuf:"bytes,1,opt,name=info,proto3" json:"info,omitempty"`
}

func (x *DescribeBillingInformationOutput) Reset() {
	*x = DescribeBillingInformationOutput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_universe_billing_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DescribeBillingInformationOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescribeBillingInformationOutput) ProtoMessage() {}

func (x *DescribeBillingInformationOutput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_universe_billing_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DescribeBillingInformationOutput.ProtoReflect.Descriptor instead.
func (*DescribeBillingInformationOutput) Descriptor() ([]byte, []int) {
	return file_eolymp_universe_billing_service_proto_rawDescGZIP(), []int{1}
}

func (x *DescribeBillingInformationOutput) GetInfo() *Billing_Information {
	if x != nil {
		return x.Info
	}
	return nil
}

type UpdateBillingInformationInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *Billing_Information `protobuf:"bytes,1,opt,name=info,proto3" json:"info,omitempty"`
}

func (x *UpdateBillingInformationInput) Reset() {
	*x = UpdateBillingInformationInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_universe_billing_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateBillingInformationInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateBillingInformationInput) ProtoMessage() {}

func (x *UpdateBillingInformationInput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_universe_billing_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateBillingInformationInput.ProtoReflect.Descriptor instead.
func (*UpdateBillingInformationInput) Descriptor() ([]byte, []int) {
	return file_eolymp_universe_billing_service_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateBillingInformationInput) GetInfo() *Billing_Information {
	if x != nil {
		return x.Info
	}
	return nil
}

type UpdateBillingInformationOutput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateBillingInformationOutput) Reset() {
	*x = UpdateBillingInformationOutput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_universe_billing_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateBillingInformationOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateBillingInformationOutput) ProtoMessage() {}

func (x *UpdateBillingInformationOutput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_universe_billing_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateBillingInformationOutput.ProtoReflect.Descriptor instead.
func (*UpdateBillingInformationOutput) Descriptor() ([]byte, []int) {
	return file_eolymp_universe_billing_service_proto_rawDescGZIP(), []int{3}
}

type DescribeCurrentPlanInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DescribeCurrentPlanInput) Reset() {
	*x = DescribeCurrentPlanInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_universe_billing_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DescribeCurrentPlanInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescribeCurrentPlanInput) ProtoMessage() {}

func (x *DescribeCurrentPlanInput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_universe_billing_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DescribeCurrentPlanInput.ProtoReflect.Descriptor instead.
func (*DescribeCurrentPlanInput) Descriptor() ([]byte, []int) {
	return file_eolymp_universe_billing_service_proto_rawDescGZIP(), []int{4}
}

type DescribeCurrentPlanOutput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PlanId          string                 `protobuf:"bytes,1,opt,name=plan_id,json=planId,proto3" json:"plan_id,omitempty"`
	Quantity        uint32                 `protobuf:"varint,2,opt,name=quantity,proto3" json:"quantity,omitempty"` // current quantity, might be different from payment quantity if customer requested changes which apply in the next billing cycle
	Status          Billing_Status         `protobuf:"varint,3,opt,name=status,proto3,enum=eolymp.universe.Billing_Status" json:"status,omitempty"`
	Price           *commerce.Price        `protobuf:"bytes,10,opt,name=price,proto3" json:"price,omitempty"`
	InvoiceQuantity uint32                 `protobuf:"varint,102,opt,name=invoice_quantity,json=invoiceQuantity,proto3" json:"invoice_quantity,omitempty"` // quantity as configured in the upcoming payment
	InvoiceSubtotal uint32                 `protobuf:"varint,104,opt,name=invoice_subtotal,json=invoiceSubtotal,proto3" json:"invoice_subtotal,omitempty"` // total amount (might not be unit_amount * quantity if discount is applied)
	CreatedAt       *timestamppb.Timestamp `protobuf:"bytes,200,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	StartedAt       *timestamppb.Timestamp `protobuf:"bytes,201,opt,name=started_at,json=startedAt,proto3" json:"started_at,omitempty"`
	RenewedAt       *timestamppb.Timestamp `protobuf:"bytes,202,opt,name=renewed_at,json=renewedAt,proto3" json:"renewed_at,omitempty"`
	RenewsAt        *timestamppb.Timestamp `protobuf:"bytes,203,opt,name=renews_at,json=renewsAt,proto3" json:"renews_at,omitempty"`
	CancelAt        *timestamppb.Timestamp `protobuf:"bytes,204,opt,name=cancel_at,json=cancelAt,proto3" json:"cancel_at,omitempty"`
	CancelledAt     *timestamppb.Timestamp `protobuf:"bytes,205,opt,name=cancelled_at,json=cancelledAt,proto3" json:"cancelled_at,omitempty"`
	PausedAt        *timestamppb.Timestamp `protobuf:"bytes,206,opt,name=paused_at,json=pausedAt,proto3" json:"paused_at,omitempty"`
}

func (x *DescribeCurrentPlanOutput) Reset() {
	*x = DescribeCurrentPlanOutput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_universe_billing_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DescribeCurrentPlanOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescribeCurrentPlanOutput) ProtoMessage() {}

func (x *DescribeCurrentPlanOutput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_universe_billing_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DescribeCurrentPlanOutput.ProtoReflect.Descriptor instead.
func (*DescribeCurrentPlanOutput) Descriptor() ([]byte, []int) {
	return file_eolymp_universe_billing_service_proto_rawDescGZIP(), []int{5}
}

func (x *DescribeCurrentPlanOutput) GetPlanId() string {
	if x != nil {
		return x.PlanId
	}
	return ""
}

func (x *DescribeCurrentPlanOutput) GetQuantity() uint32 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

func (x *DescribeCurrentPlanOutput) GetStatus() Billing_Status {
	if x != nil {
		return x.Status
	}
	return Billing_UNKNOWN_STATUS
}

func (x *DescribeCurrentPlanOutput) GetPrice() *commerce.Price {
	if x != nil {
		return x.Price
	}
	return nil
}

func (x *DescribeCurrentPlanOutput) GetInvoiceQuantity() uint32 {
	if x != nil {
		return x.InvoiceQuantity
	}
	return 0
}

func (x *DescribeCurrentPlanOutput) GetInvoiceSubtotal() uint32 {
	if x != nil {
		return x.InvoiceSubtotal
	}
	return 0
}

func (x *DescribeCurrentPlanOutput) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *DescribeCurrentPlanOutput) GetStartedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.StartedAt
	}
	return nil
}

func (x *DescribeCurrentPlanOutput) GetRenewedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.RenewedAt
	}
	return nil
}

func (x *DescribeCurrentPlanOutput) GetRenewsAt() *timestamppb.Timestamp {
	if x != nil {
		return x.RenewsAt
	}
	return nil
}

func (x *DescribeCurrentPlanOutput) GetCancelAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CancelAt
	}
	return nil
}

func (x *DescribeCurrentPlanOutput) GetCancelledAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CancelledAt
	}
	return nil
}

func (x *DescribeCurrentPlanOutput) GetPausedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.PausedAt
	}
	return nil
}

type UpdateCurrentPlanInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PlanId   string `protobuf:"bytes,1,opt,name=plan_id,json=planId,proto3" json:"plan_id,omitempty"`
	PriceId  string `protobuf:"bytes,2,opt,name=price_id,json=priceId,proto3" json:"price_id,omitempty"`
	Quantity uint32 `protobuf:"varint,3,opt,name=quantity,proto3" json:"quantity,omitempty"`
}

func (x *UpdateCurrentPlanInput) Reset() {
	*x = UpdateCurrentPlanInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_universe_billing_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateCurrentPlanInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateCurrentPlanInput) ProtoMessage() {}

func (x *UpdateCurrentPlanInput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_universe_billing_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateCurrentPlanInput.ProtoReflect.Descriptor instead.
func (*UpdateCurrentPlanInput) Descriptor() ([]byte, []int) {
	return file_eolymp_universe_billing_service_proto_rawDescGZIP(), []int{6}
}

func (x *UpdateCurrentPlanInput) GetPlanId() string {
	if x != nil {
		return x.PlanId
	}
	return ""
}

func (x *UpdateCurrentPlanInput) GetPriceId() string {
	if x != nil {
		return x.PriceId
	}
	return ""
}

func (x *UpdateCurrentPlanInput) GetQuantity() uint32 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

type UpdateCurrentPlanOutput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CheckoutUrl string `protobuf:"bytes,1,opt,name=checkout_url,json=checkoutUrl,proto3" json:"checkout_url,omitempty"`
}

func (x *UpdateCurrentPlanOutput) Reset() {
	*x = UpdateCurrentPlanOutput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_universe_billing_service_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateCurrentPlanOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateCurrentPlanOutput) ProtoMessage() {}

func (x *UpdateCurrentPlanOutput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_universe_billing_service_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateCurrentPlanOutput.ProtoReflect.Descriptor instead.
func (*UpdateCurrentPlanOutput) Descriptor() ([]byte, []int) {
	return file_eolymp_universe_billing_service_proto_rawDescGZIP(), []int{7}
}

func (x *UpdateCurrentPlanOutput) GetCheckoutUrl() string {
	if x != nil {
		return x.CheckoutUrl
	}
	return ""
}

type CancelCurrentPlanInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CancelCurrentPlanInput) Reset() {
	*x = CancelCurrentPlanInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_universe_billing_service_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CancelCurrentPlanInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CancelCurrentPlanInput) ProtoMessage() {}

func (x *CancelCurrentPlanInput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_universe_billing_service_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CancelCurrentPlanInput.ProtoReflect.Descriptor instead.
func (*CancelCurrentPlanInput) Descriptor() ([]byte, []int) {
	return file_eolymp_universe_billing_service_proto_rawDescGZIP(), []int{8}
}

type CancelCurrentPlanOutput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CancelCurrentPlanOutput) Reset() {
	*x = CancelCurrentPlanOutput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_universe_billing_service_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CancelCurrentPlanOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CancelCurrentPlanOutput) ProtoMessage() {}

func (x *CancelCurrentPlanOutput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_universe_billing_service_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CancelCurrentPlanOutput.ProtoReflect.Descriptor instead.
func (*CancelCurrentPlanOutput) Descriptor() ([]byte, []int) {
	return file_eolymp_universe_billing_service_proto_rawDescGZIP(), []int{9}
}

var File_eolymp_universe_billing_service_proto protoreflect.FileDescriptor

var file_eolymp_universe_billing_service_proto_rawDesc = []byte{
	0x0a, 0x25, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x75, 0x6e, 0x69, 0x76, 0x65, 0x72, 0x73,
	0x65, 0x2f, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e,
	0x75, 0x6e, 0x69, 0x76, 0x65, 0x72, 0x73, 0x65, 0x1a, 0x1d, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70,
	0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x68, 0x74, 0x74,
	0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x22, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x72, 0x61, 0x74, 0x65,
	0x6c, 0x69, 0x6d, 0x69, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x65, 0x6f, 0x6c,
	0x79, 0x6d, 0x70, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x63, 0x65, 0x2f, 0x70, 0x72, 0x69,
	0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70,
	0x2f, 0x75, 0x6e, 0x69, 0x76, 0x65, 0x72, 0x73, 0x65, 0x2f, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e,
	0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x21, 0x0a, 0x1f, 0x44, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x62, 0x65, 0x42, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x49, 0x6e, 0x66, 0x6f, 0x72,
	0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x22, 0x5c, 0x0a, 0x20, 0x44,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x42, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x49, 0x6e,
	0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x12,
	0x38, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e,
	0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x75, 0x6e, 0x69, 0x76, 0x65, 0x72, 0x73, 0x65, 0x2e,
	0x42, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x22, 0x59, 0x0a, 0x1d, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x42, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x49, 0x6e, 0x66, 0x6f, 0x72, 0x6d,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x38, 0x0a, 0x04, 0x69, 0x6e,
	0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d,
	0x70, 0x2e, 0x75, 0x6e, 0x69, 0x76, 0x65, 0x72, 0x73, 0x65, 0x2e, 0x42, 0x69, 0x6c, 0x6c, 0x69,
	0x6e, 0x67, 0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x04,
	0x69, 0x6e, 0x66, 0x6f, 0x22, 0x20, 0x0a, 0x1e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x69,
	0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x49, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x22, 0x1a, 0x0a, 0x18, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x62, 0x65, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x50, 0x6c, 0x61, 0x6e, 0x49, 0x6e, 0x70,
	0x75, 0x74, 0x22, 0xaf, 0x05, 0x0a, 0x19, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x43,
	0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x50, 0x6c, 0x61, 0x6e, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74,
	0x12, 0x17, 0x0a, 0x07, 0x70, 0x6c, 0x61, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x70, 0x6c, 0x61, 0x6e, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x71, 0x75, 0x61,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x71, 0x75, 0x61,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x37, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1f, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x75,
	0x6e, 0x69, 0x76, 0x65, 0x72, 0x73, 0x65, 0x2e, 0x42, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2e,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x2c,
	0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e,
	0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x63, 0x65, 0x2e,
	0x50, 0x72, 0x69, 0x63, 0x65, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x29, 0x0a, 0x10,
	0x69, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x5f, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x18, 0x66, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0f, 0x69, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x51,
	0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x29, 0x0a, 0x10, 0x69, 0x6e, 0x76, 0x6f, 0x69,
	0x63, 0x65, 0x5f, 0x73, 0x75, 0x62, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x68, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x0f, 0x69, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x53, 0x75, 0x62, 0x74, 0x6f, 0x74,
	0x61, 0x6c, 0x12, 0x3a, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74,
	0x18, 0xc8, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x3a,
	0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0xc9, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x3a, 0x0a, 0x0a, 0x72, 0x65,
	0x6e, 0x65, 0x77, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0xca, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x72, 0x65, 0x6e,
	0x65, 0x77, 0x65, 0x64, 0x41, 0x74, 0x12, 0x38, 0x0a, 0x09, 0x72, 0x65, 0x6e, 0x65, 0x77, 0x73,
	0x5f, 0x61, 0x74, 0x18, 0xcb, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x08, 0x72, 0x65, 0x6e, 0x65, 0x77, 0x73, 0x41, 0x74,
	0x12, 0x38, 0x0a, 0x09, 0x63, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x5f, 0x61, 0x74, 0x18, 0xcc, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x08, 0x63, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x41, 0x74, 0x12, 0x3e, 0x0a, 0x0c, 0x63, 0x61,
	0x6e, 0x63, 0x65, 0x6c, 0x6c, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0xcd, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0b, 0x63,
	0x61, 0x6e, 0x63, 0x65, 0x6c, 0x6c, 0x65, 0x64, 0x41, 0x74, 0x12, 0x38, 0x0a, 0x09, 0x70, 0x61,
	0x75, 0x73, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0xce, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x08, 0x70, 0x61, 0x75, 0x73,
	0x65, 0x64, 0x41, 0x74, 0x22, 0x68, 0x0a, 0x16, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x75,
	0x72, 0x72, 0x65, 0x6e, 0x74, 0x50, 0x6c, 0x61, 0x6e, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x17,
	0x0a, 0x07, 0x70, 0x6c, 0x61, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x70, 0x6c, 0x61, 0x6e, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x70, 0x72, 0x69, 0x63, 0x65,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x72, 0x69, 0x63, 0x65,
	0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x22, 0x3c,
	0x0a, 0x17, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x50,
	0x6c, 0x61, 0x6e, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x68, 0x65,
	0x63, 0x6b, 0x6f, 0x75, 0x74, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x6f, 0x75, 0x74, 0x55, 0x72, 0x6c, 0x22, 0x18, 0x0a, 0x16,
	0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x50, 0x6c, 0x61,
	0x6e, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x22, 0x19, 0x0a, 0x17, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c,
	0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x50, 0x6c, 0x61, 0x6e, 0x4f, 0x75, 0x74, 0x70, 0x75,
	0x74, 0x32, 0x91, 0x06, 0x0a, 0x0e, 0x42, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0xa7, 0x01, 0x0a, 0x1a, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62,
	0x65, 0x42, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x49, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x30, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x75, 0x6e, 0x69,
	0x76, 0x65, 0x72, 0x73, 0x65, 0x2e, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x42, 0x69,
	0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x49, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x31, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x75,
	0x6e, 0x69, 0x76, 0x65, 0x72, 0x73, 0x65, 0x2e, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65,
	0x42, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x49, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x22, 0x24, 0xea, 0xe2, 0x0a, 0x0b, 0xf5, 0xe2,
	0x0a, 0x00, 0x00, 0xa0, 0x40, 0xf8, 0xe2, 0x0a, 0x32, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0f, 0x12,
	0x0d, 0x2f, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2f, 0x69, 0x6e, 0x66, 0x6f, 0x12, 0xa1,
	0x01, 0x0a, 0x18, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67,
	0x49, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2e, 0x2e, 0x65, 0x6f,
	0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x75, 0x6e, 0x69, 0x76, 0x65, 0x72, 0x73, 0x65, 0x2e, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x42, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x49, 0x6e, 0x66, 0x6f, 0x72,
	0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x2f, 0x2e, 0x65, 0x6f,
	0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x75, 0x6e, 0x69, 0x76, 0x65, 0x72, 0x73, 0x65, 0x2e, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x42, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x49, 0x6e, 0x66, 0x6f, 0x72,
	0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x22, 0x24, 0xea, 0xe2,
	0x0a, 0x0b, 0xf5, 0xe2, 0x0a, 0x00, 0x00, 0xa0, 0x40, 0xf8, 0xe2, 0x0a, 0x32, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x0f, 0x1a, 0x0d, 0x2f, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2f, 0x69, 0x6e,
	0x66, 0x6f, 0x12, 0x92, 0x01, 0x0a, 0x13, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x43,
	0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x50, 0x6c, 0x61, 0x6e, 0x12, 0x29, 0x2e, 0x65, 0x6f, 0x6c,
	0x79, 0x6d, 0x70, 0x2e, 0x75, 0x6e, 0x69, 0x76, 0x65, 0x72, 0x73, 0x65, 0x2e, 0x44, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x62, 0x65, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x50, 0x6c, 0x61, 0x6e,
	0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x2a, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x75,
	0x6e, 0x69, 0x76, 0x65, 0x72, 0x73, 0x65, 0x2e, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65,
	0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x50, 0x6c, 0x61, 0x6e, 0x4f, 0x75, 0x74, 0x70, 0x75,
	0x74, 0x22, 0x24, 0xea, 0xe2, 0x0a, 0x0b, 0xf5, 0xe2, 0x0a, 0x00, 0x00, 0xa0, 0x40, 0xf8, 0xe2,
	0x0a, 0x32, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0f, 0x12, 0x0d, 0x2f, 0x62, 0x69, 0x6c, 0x6c, 0x69,
	0x6e, 0x67, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x12, 0x8c, 0x01, 0x0a, 0x11, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x50, 0x6c, 0x61, 0x6e, 0x12, 0x27, 0x2e,
	0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x75, 0x6e, 0x69, 0x76, 0x65, 0x72, 0x73, 0x65, 0x2e,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x50, 0x6c, 0x61,
	0x6e, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x28, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e,
	0x75, 0x6e, 0x69, 0x76, 0x65, 0x72, 0x73, 0x65, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43,
	0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x50, 0x6c, 0x61, 0x6e, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74,
	0x22, 0x24, 0xea, 0xe2, 0x0a, 0x0b, 0xf5, 0xe2, 0x0a, 0x00, 0x00, 0xa0, 0x40, 0xf8, 0xe2, 0x0a,
	0x32, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0f, 0x1a, 0x0d, 0x2f, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e,
	0x67, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x12, 0x8c, 0x01, 0x0a, 0x11, 0x43, 0x61, 0x6e, 0x63, 0x65,
	0x6c, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x50, 0x6c, 0x61, 0x6e, 0x12, 0x27, 0x2e, 0x65,
	0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x75, 0x6e, 0x69, 0x76, 0x65, 0x72, 0x73, 0x65, 0x2e, 0x43,
	0x61, 0x6e, 0x63, 0x65, 0x6c, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x50, 0x6c, 0x61, 0x6e,
	0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x28, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x75,
	0x6e, 0x69, 0x76, 0x65, 0x72, 0x73, 0x65, 0x2e, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x43, 0x75,
	0x72, 0x72, 0x65, 0x6e, 0x74, 0x50, 0x6c, 0x61, 0x6e, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x22,
	0x24, 0xea, 0xe2, 0x0a, 0x0b, 0xf5, 0xe2, 0x0a, 0x00, 0x00, 0xa0, 0x40, 0xf8, 0xe2, 0x0a, 0x32,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0f, 0x2a, 0x0d, 0x2f, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67,
	0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x42, 0x33, 0x5a, 0x31, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x67, 0x6f, 0x2d, 0x73, 0x64,
	0x6b, 0x2f, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x75, 0x6e, 0x69, 0x76, 0x65, 0x72, 0x73,
	0x65, 0x3b, 0x75, 0x6e, 0x69, 0x76, 0x65, 0x72, 0x73, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_eolymp_universe_billing_service_proto_rawDescOnce sync.Once
	file_eolymp_universe_billing_service_proto_rawDescData = file_eolymp_universe_billing_service_proto_rawDesc
)

func file_eolymp_universe_billing_service_proto_rawDescGZIP() []byte {
	file_eolymp_universe_billing_service_proto_rawDescOnce.Do(func() {
		file_eolymp_universe_billing_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_eolymp_universe_billing_service_proto_rawDescData)
	})
	return file_eolymp_universe_billing_service_proto_rawDescData
}

var file_eolymp_universe_billing_service_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_eolymp_universe_billing_service_proto_goTypes = []interface{}{
	(*DescribeBillingInformationInput)(nil),  // 0: eolymp.universe.DescribeBillingInformationInput
	(*DescribeBillingInformationOutput)(nil), // 1: eolymp.universe.DescribeBillingInformationOutput
	(*UpdateBillingInformationInput)(nil),    // 2: eolymp.universe.UpdateBillingInformationInput
	(*UpdateBillingInformationOutput)(nil),   // 3: eolymp.universe.UpdateBillingInformationOutput
	(*DescribeCurrentPlanInput)(nil),         // 4: eolymp.universe.DescribeCurrentPlanInput
	(*DescribeCurrentPlanOutput)(nil),        // 5: eolymp.universe.DescribeCurrentPlanOutput
	(*UpdateCurrentPlanInput)(nil),           // 6: eolymp.universe.UpdateCurrentPlanInput
	(*UpdateCurrentPlanOutput)(nil),          // 7: eolymp.universe.UpdateCurrentPlanOutput
	(*CancelCurrentPlanInput)(nil),           // 8: eolymp.universe.CancelCurrentPlanInput
	(*CancelCurrentPlanOutput)(nil),          // 9: eolymp.universe.CancelCurrentPlanOutput
	(*Billing_Information)(nil),              // 10: eolymp.universe.Billing.Information
	(Billing_Status)(0),                      // 11: eolymp.universe.Billing.Status
	(*commerce.Price)(nil),                   // 12: eolymp.commerce.Price
	(*timestamppb.Timestamp)(nil),            // 13: google.protobuf.Timestamp
}
var file_eolymp_universe_billing_service_proto_depIdxs = []int32{
	10, // 0: eolymp.universe.DescribeBillingInformationOutput.info:type_name -> eolymp.universe.Billing.Information
	10, // 1: eolymp.universe.UpdateBillingInformationInput.info:type_name -> eolymp.universe.Billing.Information
	11, // 2: eolymp.universe.DescribeCurrentPlanOutput.status:type_name -> eolymp.universe.Billing.Status
	12, // 3: eolymp.universe.DescribeCurrentPlanOutput.price:type_name -> eolymp.commerce.Price
	13, // 4: eolymp.universe.DescribeCurrentPlanOutput.created_at:type_name -> google.protobuf.Timestamp
	13, // 5: eolymp.universe.DescribeCurrentPlanOutput.started_at:type_name -> google.protobuf.Timestamp
	13, // 6: eolymp.universe.DescribeCurrentPlanOutput.renewed_at:type_name -> google.protobuf.Timestamp
	13, // 7: eolymp.universe.DescribeCurrentPlanOutput.renews_at:type_name -> google.protobuf.Timestamp
	13, // 8: eolymp.universe.DescribeCurrentPlanOutput.cancel_at:type_name -> google.protobuf.Timestamp
	13, // 9: eolymp.universe.DescribeCurrentPlanOutput.cancelled_at:type_name -> google.protobuf.Timestamp
	13, // 10: eolymp.universe.DescribeCurrentPlanOutput.paused_at:type_name -> google.protobuf.Timestamp
	0,  // 11: eolymp.universe.BillingService.DescribeBillingInformation:input_type -> eolymp.universe.DescribeBillingInformationInput
	2,  // 12: eolymp.universe.BillingService.UpdateBillingInformation:input_type -> eolymp.universe.UpdateBillingInformationInput
	4,  // 13: eolymp.universe.BillingService.DescribeCurrentPlan:input_type -> eolymp.universe.DescribeCurrentPlanInput
	6,  // 14: eolymp.universe.BillingService.UpdateCurrentPlan:input_type -> eolymp.universe.UpdateCurrentPlanInput
	8,  // 15: eolymp.universe.BillingService.CancelCurrentPlan:input_type -> eolymp.universe.CancelCurrentPlanInput
	1,  // 16: eolymp.universe.BillingService.DescribeBillingInformation:output_type -> eolymp.universe.DescribeBillingInformationOutput
	3,  // 17: eolymp.universe.BillingService.UpdateBillingInformation:output_type -> eolymp.universe.UpdateBillingInformationOutput
	5,  // 18: eolymp.universe.BillingService.DescribeCurrentPlan:output_type -> eolymp.universe.DescribeCurrentPlanOutput
	7,  // 19: eolymp.universe.BillingService.UpdateCurrentPlan:output_type -> eolymp.universe.UpdateCurrentPlanOutput
	9,  // 20: eolymp.universe.BillingService.CancelCurrentPlan:output_type -> eolymp.universe.CancelCurrentPlanOutput
	16, // [16:21] is the sub-list for method output_type
	11, // [11:16] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_eolymp_universe_billing_service_proto_init() }
func file_eolymp_universe_billing_service_proto_init() {
	if File_eolymp_universe_billing_service_proto != nil {
		return
	}
	file_eolymp_universe_billing_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_eolymp_universe_billing_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DescribeBillingInformationInput); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_eolymp_universe_billing_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DescribeBillingInformationOutput); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_eolymp_universe_billing_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateBillingInformationInput); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_eolymp_universe_billing_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateBillingInformationOutput); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_eolymp_universe_billing_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DescribeCurrentPlanInput); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_eolymp_universe_billing_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DescribeCurrentPlanOutput); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_eolymp_universe_billing_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateCurrentPlanInput); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_eolymp_universe_billing_service_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateCurrentPlanOutput); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_eolymp_universe_billing_service_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CancelCurrentPlanInput); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_eolymp_universe_billing_service_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CancelCurrentPlanOutput); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_eolymp_universe_billing_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_eolymp_universe_billing_service_proto_goTypes,
		DependencyIndexes: file_eolymp_universe_billing_service_proto_depIdxs,
		MessageInfos:      file_eolymp_universe_billing_service_proto_msgTypes,
	}.Build()
	File_eolymp_universe_billing_service_proto = out.File
	file_eolymp_universe_billing_service_proto_rawDesc = nil
	file_eolymp_universe_billing_service_proto_goTypes = nil
	file_eolymp_universe_billing_service_proto_depIdxs = nil
}
