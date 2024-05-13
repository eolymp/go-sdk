// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v3.21.12
// source: eolymp/universe/billing.proto

package universe

import (
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

type Billing_Status int32

const (
	Billing_UNKNOWN_STATUS Billing_Status = 0
	Billing_TRIAL          Billing_Status = 1
	Billing_ACTIVE         Billing_Status = 2
	Billing_CANCELLED      Billing_Status = 3
)

// Enum value maps for Billing_Status.
var (
	Billing_Status_name = map[int32]string{
		0: "UNKNOWN_STATUS",
		1: "TRIAL",
		2: "ACTIVE",
		3: "CANCELLED",
	}
	Billing_Status_value = map[string]int32{
		"UNKNOWN_STATUS": 0,
		"TRIAL":          1,
		"ACTIVE":         2,
		"CANCELLED":      3,
	}
)

func (x Billing_Status) Enum() *Billing_Status {
	p := new(Billing_Status)
	*p = x
	return p
}

func (x Billing_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Billing_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_eolymp_universe_billing_proto_enumTypes[0].Descriptor()
}

func (Billing_Status) Type() protoreflect.EnumType {
	return &file_eolymp_universe_billing_proto_enumTypes[0]
}

func (x Billing_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Billing_Status.Descriptor instead.
func (Billing_Status) EnumDescriptor() ([]byte, []int) {
	return file_eolymp_universe_billing_proto_rawDescGZIP(), []int{0, 0}
}

type Billing_Recurrence int32

const (
	Billing_UNKNOWN_RECURRENCE Billing_Recurrence = 0
	Billing_MONTHLY            Billing_Recurrence = 1
	Billing_YEARLY             Billing_Recurrence = 2
)

// Enum value maps for Billing_Recurrence.
var (
	Billing_Recurrence_name = map[int32]string{
		0: "UNKNOWN_RECURRENCE",
		1: "MONTHLY",
		2: "YEARLY",
	}
	Billing_Recurrence_value = map[string]int32{
		"UNKNOWN_RECURRENCE": 0,
		"MONTHLY":            1,
		"YEARLY":             2,
	}
)

func (x Billing_Recurrence) Enum() *Billing_Recurrence {
	p := new(Billing_Recurrence)
	*p = x
	return p
}

func (x Billing_Recurrence) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Billing_Recurrence) Descriptor() protoreflect.EnumDescriptor {
	return file_eolymp_universe_billing_proto_enumTypes[1].Descriptor()
}

func (Billing_Recurrence) Type() protoreflect.EnumType {
	return &file_eolymp_universe_billing_proto_enumTypes[1]
}

func (x Billing_Recurrence) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Billing_Recurrence.Descriptor instead.
func (Billing_Recurrence) EnumDescriptor() ([]byte, []int) {
	return file_eolymp_universe_billing_proto_rawDescGZIP(), []int{0, 1}
}

type Billing struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Billing) Reset() {
	*x = Billing{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_universe_billing_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Billing) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Billing) ProtoMessage() {}

func (x *Billing) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_universe_billing_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Billing.ProtoReflect.Descriptor instead.
func (*Billing) Descriptor() ([]byte, []int) {
	return file_eolymp_universe_billing_proto_rawDescGZIP(), []int{0}
}

type Billing_Information struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name       string            `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Email      string            `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	Phone      string            `protobuf:"bytes,3,opt,name=phone,proto3" json:"phone,omitempty"`
	TaxIdType  string            `protobuf:"bytes,10,opt,name=tax_id_type,json=taxIdType,proto3" json:"tax_id_type,omitempty"`
	TaxIdValue string            `protobuf:"bytes,11,opt,name=tax_id_value,json=taxIdValue,proto3" json:"tax_id_value,omitempty"`
	Address    *commerce.Address `protobuf:"bytes,20,opt,name=address,proto3" json:"address,omitempty"`
	Currency   string            `protobuf:"bytes,90,opt,name=currency,proto3" json:"currency,omitempty"`
	Language   string            `protobuf:"bytes,91,opt,name=language,proto3" json:"language,omitempty"`
}

func (x *Billing_Information) Reset() {
	*x = Billing_Information{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_universe_billing_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Billing_Information) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Billing_Information) ProtoMessage() {}

func (x *Billing_Information) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_universe_billing_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Billing_Information.ProtoReflect.Descriptor instead.
func (*Billing_Information) Descriptor() ([]byte, []int) {
	return file_eolymp_universe_billing_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Billing_Information) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Billing_Information) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *Billing_Information) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *Billing_Information) GetTaxIdType() string {
	if x != nil {
		return x.TaxIdType
	}
	return ""
}

func (x *Billing_Information) GetTaxIdValue() string {
	if x != nil {
		return x.TaxIdValue
	}
	return ""
}

func (x *Billing_Information) GetAddress() *commerce.Address {
	if x != nil {
		return x.Address
	}
	return nil
}

func (x *Billing_Information) GetCurrency() string {
	if x != nil {
		return x.Currency
	}
	return ""
}

func (x *Billing_Information) GetLanguage() string {
	if x != nil {
		return x.Language
	}
	return ""
}

type Billing_Subscription struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     string         `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Status Billing_Status `protobuf:"varint,2,opt,name=status,proto3,enum=eolymp.universe.Billing_Status" json:"status,omitempty"`
	// Irregular structure flag
	//
	// If this flag is set to true, show warning to the customer because subscription might have multiple items,
	// unrecognized changes or other irregularities. It might be due to an error or special manual configuration.
	//
	// Updating irregular subscriptions is not possible without intervention from support.
	Irregular bool `protobuf:"varint,3,opt,name=irregular,proto3" json:"irregular,omitempty"`
	// Defines if payment method has been added to the subscription
	HasPaymentMethod bool                   `protobuf:"varint,4,opt,name=has_payment_method,json=hasPaymentMethod,proto3" json:"has_payment_method,omitempty"`
	Plan             *Plan                  `protobuf:"bytes,10,opt,name=plan,proto3" json:"plan,omitempty"`
	Variant          *Plan_Variant          `protobuf:"bytes,11,opt,name=variant,proto3" json:"variant,omitempty"`
	Seats            uint32                 `protobuf:"varint,12,opt,name=seats,proto3" json:"seats,omitempty"`
	CreatedAt        *timestamppb.Timestamp `protobuf:"bytes,210,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`       // time when subscription was created
	StartedAt        *timestamppb.Timestamp `protobuf:"bytes,211,opt,name=started_at,json=startedAt,proto3" json:"started_at,omitempty"`       // time when subscription was started (activated first time)
	CancelAt         *timestamppb.Timestamp `protobuf:"bytes,213,opt,name=cancel_at,json=cancelAt,proto3" json:"cancel_at,omitempty"`          // time when subscription was/will be cancelled
	CancelledAt      *timestamppb.Timestamp `protobuf:"bytes,214,opt,name=cancelled_at,json=cancelledAt,proto3" json:"cancelled_at,omitempty"` // time when subscription was/will be cancelled
	EndedAt          *timestamppb.Timestamp `protobuf:"bytes,215,opt,name=ended_at,json=endedAt,proto3" json:"ended_at,omitempty"`             // time when subscription was/will be cancelled
	PeriodStart      *timestamppb.Timestamp `protobuf:"bytes,240,opt,name=period_start,json=periodStart,proto3" json:"period_start,omitempty"` // current billing period start
	PeriodEnd        *timestamppb.Timestamp `protobuf:"bytes,241,opt,name=period_end,json=periodEnd,proto3" json:"period_end,omitempty"`       // current billing period end
	TrialStart       *timestamppb.Timestamp `protobuf:"bytes,270,opt,name=trial_start,json=trialStart,proto3" json:"trial_start,omitempty"`    // time when trial has started
	TrialEnd         *timestamppb.Timestamp `protobuf:"bytes,271,opt,name=trial_end,json=trialEnd,proto3" json:"trial_end,omitempty"`          // time when trial will be over
}

func (x *Billing_Subscription) Reset() {
	*x = Billing_Subscription{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_universe_billing_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Billing_Subscription) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Billing_Subscription) ProtoMessage() {}

func (x *Billing_Subscription) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_universe_billing_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Billing_Subscription.ProtoReflect.Descriptor instead.
func (*Billing_Subscription) Descriptor() ([]byte, []int) {
	return file_eolymp_universe_billing_proto_rawDescGZIP(), []int{0, 1}
}

func (x *Billing_Subscription) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Billing_Subscription) GetStatus() Billing_Status {
	if x != nil {
		return x.Status
	}
	return Billing_UNKNOWN_STATUS
}

func (x *Billing_Subscription) GetIrregular() bool {
	if x != nil {
		return x.Irregular
	}
	return false
}

func (x *Billing_Subscription) GetHasPaymentMethod() bool {
	if x != nil {
		return x.HasPaymentMethod
	}
	return false
}

func (x *Billing_Subscription) GetPlan() *Plan {
	if x != nil {
		return x.Plan
	}
	return nil
}

func (x *Billing_Subscription) GetVariant() *Plan_Variant {
	if x != nil {
		return x.Variant
	}
	return nil
}

func (x *Billing_Subscription) GetSeats() uint32 {
	if x != nil {
		return x.Seats
	}
	return 0
}

func (x *Billing_Subscription) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Billing_Subscription) GetStartedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.StartedAt
	}
	return nil
}

func (x *Billing_Subscription) GetCancelAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CancelAt
	}
	return nil
}

func (x *Billing_Subscription) GetCancelledAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CancelledAt
	}
	return nil
}

func (x *Billing_Subscription) GetEndedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.EndedAt
	}
	return nil
}

func (x *Billing_Subscription) GetPeriodStart() *timestamppb.Timestamp {
	if x != nil {
		return x.PeriodStart
	}
	return nil
}

func (x *Billing_Subscription) GetPeriodEnd() *timestamppb.Timestamp {
	if x != nil {
		return x.PeriodEnd
	}
	return nil
}

func (x *Billing_Subscription) GetTrialStart() *timestamppb.Timestamp {
	if x != nil {
		return x.TrialStart
	}
	return nil
}

func (x *Billing_Subscription) GetTrialEnd() *timestamppb.Timestamp {
	if x != nil {
		return x.TrialEnd
	}
	return nil
}

var File_eolymp_universe_billing_proto protoreflect.FileDescriptor

var file_eolymp_universe_billing_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x75, 0x6e, 0x69, 0x76, 0x65, 0x72, 0x73,
	0x65, 0x2f, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0f, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x75, 0x6e, 0x69, 0x76, 0x65, 0x72, 0x73, 0x65,
	0x1a, 0x1d, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x63,
	0x65, 0x2f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1a, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x75, 0x6e, 0x69, 0x76, 0x65, 0x72, 0x73, 0x65,
	0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc8, 0x09, 0x0a,
	0x07, 0x42, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x1a, 0xfb, 0x01, 0x0a, 0x0b, 0x49, 0x6e, 0x66,
	0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61,
	0x69, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x1e, 0x0a, 0x0b, 0x74, 0x61, 0x78, 0x5f,
	0x69, 0x64, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74,
	0x61, 0x78, 0x49, 0x64, 0x54, 0x79, 0x70, 0x65, 0x12, 0x20, 0x0a, 0x0c, 0x74, 0x61, 0x78, 0x5f,
	0x69, 0x64, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x74, 0x61, 0x78, 0x49, 0x64, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x32, 0x0a, 0x07, 0x61, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x14, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x65, 0x6f,
	0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x63, 0x65, 0x2e, 0x41, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1a,
	0x0a, 0x08, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x5a, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61,
	0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x18, 0x5b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61,
	0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x1a, 0xbb, 0x06, 0x0a, 0x0c, 0x53, 0x75, 0x62, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x37, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1f, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70,
	0x2e, 0x75, 0x6e, 0x69, 0x76, 0x65, 0x72, 0x73, 0x65, 0x2e, 0x42, 0x69, 0x6c, 0x6c, 0x69, 0x6e,
	0x67, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x1c, 0x0a, 0x09, 0x69, 0x72, 0x72, 0x65, 0x67, 0x75, 0x6c, 0x61, 0x72, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x09, 0x69, 0x72, 0x72, 0x65, 0x67, 0x75, 0x6c, 0x61, 0x72, 0x12, 0x2c,
	0x0a, 0x12, 0x68, 0x61, 0x73, 0x5f, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x6d, 0x65,
	0x74, 0x68, 0x6f, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x10, 0x68, 0x61, 0x73, 0x50,
	0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x29, 0x0a, 0x04,
	0x70, 0x6c, 0x61, 0x6e, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x65, 0x6f, 0x6c,
	0x79, 0x6d, 0x70, 0x2e, 0x75, 0x6e, 0x69, 0x76, 0x65, 0x72, 0x73, 0x65, 0x2e, 0x50, 0x6c, 0x61,
	0x6e, 0x52, 0x04, 0x70, 0x6c, 0x61, 0x6e, 0x12, 0x37, 0x0a, 0x07, 0x76, 0x61, 0x72, 0x69, 0x61,
	0x6e, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d,
	0x70, 0x2e, 0x75, 0x6e, 0x69, 0x76, 0x65, 0x72, 0x73, 0x65, 0x2e, 0x50, 0x6c, 0x61, 0x6e, 0x2e,
	0x56, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x74, 0x52, 0x07, 0x76, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x74,
	0x12, 0x14, 0x0a, 0x05, 0x73, 0x65, 0x61, 0x74, 0x73, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x05, 0x73, 0x65, 0x61, 0x74, 0x73, 0x12, 0x3a, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x5f, 0x61, 0x74, 0x18, 0xd2, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x12, 0x3a, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74,
	0x18, 0xd3, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x38,
	0x0a, 0x09, 0x63, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x5f, 0x61, 0x74, 0x18, 0xd5, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x08,
	0x63, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x41, 0x74, 0x12, 0x3e, 0x0a, 0x0c, 0x63, 0x61, 0x6e, 0x63,
	0x65, 0x6c, 0x6c, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0xd6, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0b, 0x63, 0x61, 0x6e,
	0x63, 0x65, 0x6c, 0x6c, 0x65, 0x64, 0x41, 0x74, 0x12, 0x36, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x65,
	0x64, 0x5f, 0x61, 0x74, 0x18, 0xd7, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x65, 0x64, 0x41, 0x74,
	0x12, 0x3e, 0x0a, 0x0c, 0x70, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x5f, 0x73, 0x74, 0x61, 0x72, 0x74,
	0x18, 0xf0, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x0b, 0x70, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x53, 0x74, 0x61, 0x72, 0x74,
	0x12, 0x3a, 0x0a, 0x0a, 0x70, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x5f, 0x65, 0x6e, 0x64, 0x18, 0xf1,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x09, 0x70, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x45, 0x6e, 0x64, 0x12, 0x3c, 0x0a, 0x0b,
	0x74, 0x72, 0x69, 0x61, 0x6c, 0x5f, 0x73, 0x74, 0x61, 0x72, 0x74, 0x18, 0x8e, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a,
	0x74, 0x72, 0x69, 0x61, 0x6c, 0x53, 0x74, 0x61, 0x72, 0x74, 0x12, 0x38, 0x0a, 0x09, 0x74, 0x72,
	0x69, 0x61, 0x6c, 0x5f, 0x65, 0x6e, 0x64, 0x18, 0x8f, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x08, 0x74, 0x72, 0x69, 0x61,
	0x6c, 0x45, 0x6e, 0x64, 0x22, 0x42, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x12,
	0x0a, 0x0e, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53,
	0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x54, 0x52, 0x49, 0x41, 0x4c, 0x10, 0x01, 0x12, 0x0a, 0x0a,
	0x06, 0x41, 0x43, 0x54, 0x49, 0x56, 0x45, 0x10, 0x02, 0x12, 0x0d, 0x0a, 0x09, 0x43, 0x41, 0x4e,
	0x43, 0x45, 0x4c, 0x4c, 0x45, 0x44, 0x10, 0x03, 0x22, 0x3d, 0x0a, 0x0a, 0x52, 0x65, 0x63, 0x75,
	0x72, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x16, 0x0a, 0x12, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57,
	0x4e, 0x5f, 0x52, 0x45, 0x43, 0x55, 0x52, 0x52, 0x45, 0x4e, 0x43, 0x45, 0x10, 0x00, 0x12, 0x0b,
	0x0a, 0x07, 0x4d, 0x4f, 0x4e, 0x54, 0x48, 0x4c, 0x59, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x59,
	0x45, 0x41, 0x52, 0x4c, 0x59, 0x10, 0x02, 0x42, 0x33, 0x5a, 0x31, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x67, 0x6f, 0x2d,
	0x73, 0x64, 0x6b, 0x2f, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x75, 0x6e, 0x69, 0x76, 0x65,
	0x72, 0x73, 0x65, 0x3b, 0x75, 0x6e, 0x69, 0x76, 0x65, 0x72, 0x73, 0x65, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_eolymp_universe_billing_proto_rawDescOnce sync.Once
	file_eolymp_universe_billing_proto_rawDescData = file_eolymp_universe_billing_proto_rawDesc
)

func file_eolymp_universe_billing_proto_rawDescGZIP() []byte {
	file_eolymp_universe_billing_proto_rawDescOnce.Do(func() {
		file_eolymp_universe_billing_proto_rawDescData = protoimpl.X.CompressGZIP(file_eolymp_universe_billing_proto_rawDescData)
	})
	return file_eolymp_universe_billing_proto_rawDescData
}

var file_eolymp_universe_billing_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_eolymp_universe_billing_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_eolymp_universe_billing_proto_goTypes = []interface{}{
	(Billing_Status)(0),           // 0: eolymp.universe.Billing.Status
	(Billing_Recurrence)(0),       // 1: eolymp.universe.Billing.Recurrence
	(*Billing)(nil),               // 2: eolymp.universe.Billing
	(*Billing_Information)(nil),   // 3: eolymp.universe.Billing.Information
	(*Billing_Subscription)(nil),  // 4: eolymp.universe.Billing.Subscription
	(*commerce.Address)(nil),      // 5: eolymp.commerce.Address
	(*Plan)(nil),                  // 6: eolymp.universe.Plan
	(*Plan_Variant)(nil),          // 7: eolymp.universe.Plan.Variant
	(*timestamppb.Timestamp)(nil), // 8: google.protobuf.Timestamp
}
var file_eolymp_universe_billing_proto_depIdxs = []int32{
	5,  // 0: eolymp.universe.Billing.Information.address:type_name -> eolymp.commerce.Address
	0,  // 1: eolymp.universe.Billing.Subscription.status:type_name -> eolymp.universe.Billing.Status
	6,  // 2: eolymp.universe.Billing.Subscription.plan:type_name -> eolymp.universe.Plan
	7,  // 3: eolymp.universe.Billing.Subscription.variant:type_name -> eolymp.universe.Plan.Variant
	8,  // 4: eolymp.universe.Billing.Subscription.created_at:type_name -> google.protobuf.Timestamp
	8,  // 5: eolymp.universe.Billing.Subscription.started_at:type_name -> google.protobuf.Timestamp
	8,  // 6: eolymp.universe.Billing.Subscription.cancel_at:type_name -> google.protobuf.Timestamp
	8,  // 7: eolymp.universe.Billing.Subscription.cancelled_at:type_name -> google.protobuf.Timestamp
	8,  // 8: eolymp.universe.Billing.Subscription.ended_at:type_name -> google.protobuf.Timestamp
	8,  // 9: eolymp.universe.Billing.Subscription.period_start:type_name -> google.protobuf.Timestamp
	8,  // 10: eolymp.universe.Billing.Subscription.period_end:type_name -> google.protobuf.Timestamp
	8,  // 11: eolymp.universe.Billing.Subscription.trial_start:type_name -> google.protobuf.Timestamp
	8,  // 12: eolymp.universe.Billing.Subscription.trial_end:type_name -> google.protobuf.Timestamp
	13, // [13:13] is the sub-list for method output_type
	13, // [13:13] is the sub-list for method input_type
	13, // [13:13] is the sub-list for extension type_name
	13, // [13:13] is the sub-list for extension extendee
	0,  // [0:13] is the sub-list for field type_name
}

func init() { file_eolymp_universe_billing_proto_init() }
func file_eolymp_universe_billing_proto_init() {
	if File_eolymp_universe_billing_proto != nil {
		return
	}
	file_eolymp_universe_plan_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_eolymp_universe_billing_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Billing); i {
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
		file_eolymp_universe_billing_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Billing_Information); i {
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
		file_eolymp_universe_billing_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Billing_Subscription); i {
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
			RawDescriptor: file_eolymp_universe_billing_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_eolymp_universe_billing_proto_goTypes,
		DependencyIndexes: file_eolymp_universe_billing_proto_depIdxs,
		EnumInfos:         file_eolymp_universe_billing_proto_enumTypes,
		MessageInfos:      file_eolymp_universe_billing_proto_msgTypes,
	}.Build()
	File_eolymp_universe_billing_proto = out.File
	file_eolymp_universe_billing_proto_rawDesc = nil
	file_eolymp_universe_billing_proto_goTypes = nil
	file_eolymp_universe_billing_proto_depIdxs = nil
}
