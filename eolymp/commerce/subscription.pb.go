// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.18.1
// source: eolymp/commerce/subscription.proto

package commerce

import (
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

type Subscription_Status int32

const (
	Subscription_UNKNOWN_STATUS     Subscription_Status = 0
	Subscription_INCOMPLETE         Subscription_Status = 1
	Subscription_INCOMPLETE_EXPIRED Subscription_Status = 2
	Subscription_TRIALING           Subscription_Status = 3
	Subscription_ACTIVE             Subscription_Status = 4
	Subscription_PAST_DUE           Subscription_Status = 5
	Subscription_CANCELLED          Subscription_Status = 6
	Subscription_UNPAID             Subscription_Status = 7
)

// Enum value maps for Subscription_Status.
var (
	Subscription_Status_name = map[int32]string{
		0: "UNKNOWN_STATUS",
		1: "INCOMPLETE",
		2: "INCOMPLETE_EXPIRED",
		3: "TRIALING",
		4: "ACTIVE",
		5: "PAST_DUE",
		6: "CANCELLED",
		7: "UNPAID",
	}
	Subscription_Status_value = map[string]int32{
		"UNKNOWN_STATUS":     0,
		"INCOMPLETE":         1,
		"INCOMPLETE_EXPIRED": 2,
		"TRIALING":           3,
		"ACTIVE":             4,
		"PAST_DUE":           5,
		"CANCELLED":          6,
		"UNPAID":             7,
	}
)

func (x Subscription_Status) Enum() *Subscription_Status {
	p := new(Subscription_Status)
	*p = x
	return p
}

func (x Subscription_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Subscription_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_eolymp_commerce_subscription_proto_enumTypes[0].Descriptor()
}

func (Subscription_Status) Type() protoreflect.EnumType {
	return &file_eolymp_commerce_subscription_proto_enumTypes[0]
}

func (x Subscription_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Subscription_Status.Descriptor instead.
func (Subscription_Status) EnumDescriptor() ([]byte, []int) {
	return file_eolymp_commerce_subscription_proto_rawDescGZIP(), []int{0, 0}
}

type Subscription struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                   string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Status               Subscription_Status    `protobuf:"varint,2,opt,name=status,proto3,enum=eolymp.commerce.Subscription_Status" json:"status,omitempty"`
	CustomerId           string                 `protobuf:"bytes,3,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	Description          string                 `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	CreatedAt            *timestamppb.Timestamp `protobuf:"bytes,10,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`       // time when subscription was created
	StartedAt            *timestamppb.Timestamp `protobuf:"bytes,11,opt,name=started_at,json=startedAt,proto3" json:"started_at,omitempty"`       // time when subscription was started (activated first time)
	CancelAt             *timestamppb.Timestamp `protobuf:"bytes,13,opt,name=cancel_at,json=cancelAt,proto3" json:"cancel_at,omitempty"`          // time when subscription was/will be cancelled
	CancelledAt          *timestamppb.Timestamp `protobuf:"bytes,14,opt,name=cancelled_at,json=cancelledAt,proto3" json:"cancelled_at,omitempty"` // time when subscription was/will be cancelled
	EndedAt              *timestamppb.Timestamp `protobuf:"bytes,15,opt,name=ended_at,json=endedAt,proto3" json:"ended_at,omitempty"`             // time when subscription was/will be cancelled
	PeriodStart          *timestamppb.Timestamp `protobuf:"bytes,40,opt,name=period_start,json=periodStart,proto3" json:"period_start,omitempty"` // current billing period start (same as last renew date)
	PeriodEnd            *timestamppb.Timestamp `protobuf:"bytes,41,opt,name=period_end,json=periodEnd,proto3" json:"period_end,omitempty"`       // current billing period end (same as renew date for next period)
	TrialStart           *timestamppb.Timestamp `protobuf:"bytes,70,opt,name=trial_start,json=trialStart,proto3" json:"trial_start,omitempty"`    // time when subscription was/will be cancelled
	TrialEnd             *timestamppb.Timestamp `protobuf:"bytes,71,opt,name=trial_end,json=trialEnd,proto3" json:"trial_end,omitempty"`          // time when subscription was/will be cancelled
	CancellationComment  string                 `protobuf:"bytes,50,opt,name=cancellation_comment,json=cancellationComment,proto3" json:"cancellation_comment,omitempty"`
	CancellationFeedback string                 `protobuf:"bytes,51,opt,name=cancellation_feedback,json=cancellationFeedback,proto3" json:"cancellation_feedback,omitempty"`
	CancellationReason   string                 `protobuf:"bytes,52,opt,name=cancellation_reason,json=cancellationReason,proto3" json:"cancellation_reason,omitempty"`
	Currency             string                 `protobuf:"bytes,31,opt,name=currency,proto3" json:"currency,omitempty"`
	Items                []*Subscription_Item   `protobuf:"bytes,999,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *Subscription) Reset() {
	*x = Subscription{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_commerce_subscription_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Subscription) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Subscription) ProtoMessage() {}

func (x *Subscription) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_commerce_subscription_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Subscription.ProtoReflect.Descriptor instead.
func (*Subscription) Descriptor() ([]byte, []int) {
	return file_eolymp_commerce_subscription_proto_rawDescGZIP(), []int{0}
}

func (x *Subscription) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Subscription) GetStatus() Subscription_Status {
	if x != nil {
		return x.Status
	}
	return Subscription_UNKNOWN_STATUS
}

func (x *Subscription) GetCustomerId() string {
	if x != nil {
		return x.CustomerId
	}
	return ""
}

func (x *Subscription) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Subscription) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Subscription) GetStartedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.StartedAt
	}
	return nil
}

func (x *Subscription) GetCancelAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CancelAt
	}
	return nil
}

func (x *Subscription) GetCancelledAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CancelledAt
	}
	return nil
}

func (x *Subscription) GetEndedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.EndedAt
	}
	return nil
}

func (x *Subscription) GetPeriodStart() *timestamppb.Timestamp {
	if x != nil {
		return x.PeriodStart
	}
	return nil
}

func (x *Subscription) GetPeriodEnd() *timestamppb.Timestamp {
	if x != nil {
		return x.PeriodEnd
	}
	return nil
}

func (x *Subscription) GetTrialStart() *timestamppb.Timestamp {
	if x != nil {
		return x.TrialStart
	}
	return nil
}

func (x *Subscription) GetTrialEnd() *timestamppb.Timestamp {
	if x != nil {
		return x.TrialEnd
	}
	return nil
}

func (x *Subscription) GetCancellationComment() string {
	if x != nil {
		return x.CancellationComment
	}
	return ""
}

func (x *Subscription) GetCancellationFeedback() string {
	if x != nil {
		return x.CancellationFeedback
	}
	return ""
}

func (x *Subscription) GetCancellationReason() string {
	if x != nil {
		return x.CancellationReason
	}
	return ""
}

func (x *Subscription) GetCurrency() string {
	if x != nil {
		return x.Currency
	}
	return ""
}

func (x *Subscription) GetItems() []*Subscription_Item {
	if x != nil {
		return x.Items
	}
	return nil
}

type Subscription_Item struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Quantity  uint32 `protobuf:"varint,2,opt,name=quantity,proto3" json:"quantity,omitempty"`
	ProductId string `protobuf:"bytes,3,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
	Price     *Price `protobuf:"bytes,4,opt,name=price,proto3" json:"price,omitempty"`
}

func (x *Subscription_Item) Reset() {
	*x = Subscription_Item{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_commerce_subscription_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Subscription_Item) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Subscription_Item) ProtoMessage() {}

func (x *Subscription_Item) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_commerce_subscription_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Subscription_Item.ProtoReflect.Descriptor instead.
func (*Subscription_Item) Descriptor() ([]byte, []int) {
	return file_eolymp_commerce_subscription_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Subscription_Item) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Subscription_Item) GetQuantity() uint32 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

func (x *Subscription_Item) GetProductId() string {
	if x != nil {
		return x.ProductId
	}
	return ""
}

func (x *Subscription_Item) GetPrice() *Price {
	if x != nil {
		return x.Price
	}
	return nil
}

var File_eolymp_commerce_subscription_proto protoreflect.FileDescriptor

var file_eolymp_commerce_subscription_proto_rawDesc = []byte{
	0x0a, 0x22, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x63,
	0x65, 0x2f, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x65, 0x72, 0x63, 0x65, 0x1a, 0x1b, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x63, 0x6f,
	0x6d, 0x6d, 0x65, 0x72, 0x63, 0x65, 0x2f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xaf, 0x09, 0x0a, 0x0c, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x3c, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x24, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x65, 0x72, 0x63, 0x65, 0x2e, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x5f, 0x61, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x12, 0x39, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0b,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x37, 0x0a, 0x09, 0x63,
	0x61, 0x6e, 0x63, 0x65, 0x6c, 0x5f, 0x61, 0x74, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x08, 0x63, 0x61, 0x6e, 0x63,
	0x65, 0x6c, 0x41, 0x74, 0x12, 0x3d, 0x0a, 0x0c, 0x63, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x6c, 0x65,
	0x64, 0x5f, 0x61, 0x74, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0b, 0x63, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x6c, 0x65,
	0x64, 0x41, 0x74, 0x12, 0x35, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18,
	0x0f, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x65, 0x64, 0x41, 0x74, 0x12, 0x3d, 0x0a, 0x0c, 0x70, 0x65,
	0x72, 0x69, 0x6f, 0x64, 0x5f, 0x73, 0x74, 0x61, 0x72, 0x74, 0x18, 0x28, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0b, 0x70, 0x65,
	0x72, 0x69, 0x6f, 0x64, 0x53, 0x74, 0x61, 0x72, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x70, 0x65, 0x72,
	0x69, 0x6f, 0x64, 0x5f, 0x65, 0x6e, 0x64, 0x18, 0x29, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x70, 0x65, 0x72, 0x69, 0x6f,
	0x64, 0x45, 0x6e, 0x64, 0x12, 0x3b, 0x0a, 0x0b, 0x74, 0x72, 0x69, 0x61, 0x6c, 0x5f, 0x73, 0x74,
	0x61, 0x72, 0x74, 0x18, 0x46, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x74, 0x72, 0x69, 0x61, 0x6c, 0x53, 0x74, 0x61, 0x72,
	0x74, 0x12, 0x37, 0x0a, 0x09, 0x74, 0x72, 0x69, 0x61, 0x6c, 0x5f, 0x65, 0x6e, 0x64, 0x18, 0x47,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x08, 0x74, 0x72, 0x69, 0x61, 0x6c, 0x45, 0x6e, 0x64, 0x12, 0x31, 0x0a, 0x14, 0x63, 0x61,
	0x6e, 0x63, 0x65, 0x6c, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x65,
	0x6e, 0x74, 0x18, 0x32, 0x20, 0x01, 0x28, 0x09, 0x52, 0x13, 0x63, 0x61, 0x6e, 0x63, 0x65, 0x6c,
	0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x33, 0x0a,
	0x15, 0x63, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x66, 0x65,
	0x65, 0x64, 0x62, 0x61, 0x63, 0x6b, 0x18, 0x33, 0x20, 0x01, 0x28, 0x09, 0x52, 0x14, 0x63, 0x61,
	0x6e, 0x63, 0x65, 0x6c, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x46, 0x65, 0x65, 0x64, 0x62, 0x61,
	0x63, 0x6b, 0x12, 0x2f, 0x0a, 0x13, 0x63, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x6c, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x18, 0x34, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x12, 0x63, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x61,
	0x73, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x18,
	0x1f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x12,
	0x39, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0xe7, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x22, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x63,
	0x65, 0x2e, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x49,
	0x74, 0x65, 0x6d, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x1a, 0x7f, 0x0a, 0x04, 0x49, 0x74,
	0x65, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x1d,
	0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x12, 0x2c, 0x0a,
	0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x65,
	0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x63, 0x65, 0x2e, 0x50,
	0x72, 0x69, 0x63, 0x65, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x22, 0x87, 0x01, 0x0a, 0x06,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x12, 0x0a, 0x0e, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57,
	0x4e, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x10, 0x00, 0x12, 0x0e, 0x0a, 0x0a, 0x49, 0x4e,
	0x43, 0x4f, 0x4d, 0x50, 0x4c, 0x45, 0x54, 0x45, 0x10, 0x01, 0x12, 0x16, 0x0a, 0x12, 0x49, 0x4e,
	0x43, 0x4f, 0x4d, 0x50, 0x4c, 0x45, 0x54, 0x45, 0x5f, 0x45, 0x58, 0x50, 0x49, 0x52, 0x45, 0x44,
	0x10, 0x02, 0x12, 0x0c, 0x0a, 0x08, 0x54, 0x52, 0x49, 0x41, 0x4c, 0x49, 0x4e, 0x47, 0x10, 0x03,
	0x12, 0x0a, 0x0a, 0x06, 0x41, 0x43, 0x54, 0x49, 0x56, 0x45, 0x10, 0x04, 0x12, 0x0c, 0x0a, 0x08,
	0x50, 0x41, 0x53, 0x54, 0x5f, 0x44, 0x55, 0x45, 0x10, 0x05, 0x12, 0x0d, 0x0a, 0x09, 0x43, 0x41,
	0x4e, 0x43, 0x45, 0x4c, 0x4c, 0x45, 0x44, 0x10, 0x06, 0x12, 0x0a, 0x0a, 0x06, 0x55, 0x4e, 0x50,
	0x41, 0x49, 0x44, 0x10, 0x07, 0x42, 0x33, 0x5a, 0x31, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x67, 0x6f, 0x2d, 0x73, 0x64,
	0x6b, 0x2f, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x63,
	0x65, 0x3b, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_eolymp_commerce_subscription_proto_rawDescOnce sync.Once
	file_eolymp_commerce_subscription_proto_rawDescData = file_eolymp_commerce_subscription_proto_rawDesc
)

func file_eolymp_commerce_subscription_proto_rawDescGZIP() []byte {
	file_eolymp_commerce_subscription_proto_rawDescOnce.Do(func() {
		file_eolymp_commerce_subscription_proto_rawDescData = protoimpl.X.CompressGZIP(file_eolymp_commerce_subscription_proto_rawDescData)
	})
	return file_eolymp_commerce_subscription_proto_rawDescData
}

var file_eolymp_commerce_subscription_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_eolymp_commerce_subscription_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_eolymp_commerce_subscription_proto_goTypes = []interface{}{
	(Subscription_Status)(0),      // 0: eolymp.commerce.Subscription.Status
	(*Subscription)(nil),          // 1: eolymp.commerce.Subscription
	(*Subscription_Item)(nil),     // 2: eolymp.commerce.Subscription.Item
	(*timestamppb.Timestamp)(nil), // 3: google.protobuf.Timestamp
	(*Price)(nil),                 // 4: eolymp.commerce.Price
}
var file_eolymp_commerce_subscription_proto_depIdxs = []int32{
	0,  // 0: eolymp.commerce.Subscription.status:type_name -> eolymp.commerce.Subscription.Status
	3,  // 1: eolymp.commerce.Subscription.created_at:type_name -> google.protobuf.Timestamp
	3,  // 2: eolymp.commerce.Subscription.started_at:type_name -> google.protobuf.Timestamp
	3,  // 3: eolymp.commerce.Subscription.cancel_at:type_name -> google.protobuf.Timestamp
	3,  // 4: eolymp.commerce.Subscription.cancelled_at:type_name -> google.protobuf.Timestamp
	3,  // 5: eolymp.commerce.Subscription.ended_at:type_name -> google.protobuf.Timestamp
	3,  // 6: eolymp.commerce.Subscription.period_start:type_name -> google.protobuf.Timestamp
	3,  // 7: eolymp.commerce.Subscription.period_end:type_name -> google.protobuf.Timestamp
	3,  // 8: eolymp.commerce.Subscription.trial_start:type_name -> google.protobuf.Timestamp
	3,  // 9: eolymp.commerce.Subscription.trial_end:type_name -> google.protobuf.Timestamp
	2,  // 10: eolymp.commerce.Subscription.items:type_name -> eolymp.commerce.Subscription.Item
	4,  // 11: eolymp.commerce.Subscription.Item.price:type_name -> eolymp.commerce.Price
	12, // [12:12] is the sub-list for method output_type
	12, // [12:12] is the sub-list for method input_type
	12, // [12:12] is the sub-list for extension type_name
	12, // [12:12] is the sub-list for extension extendee
	0,  // [0:12] is the sub-list for field type_name
}

func init() { file_eolymp_commerce_subscription_proto_init() }
func file_eolymp_commerce_subscription_proto_init() {
	if File_eolymp_commerce_subscription_proto != nil {
		return
	}
	file_eolymp_commerce_price_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_eolymp_commerce_subscription_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Subscription); i {
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
		file_eolymp_commerce_subscription_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Subscription_Item); i {
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
			RawDescriptor: file_eolymp_commerce_subscription_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_eolymp_commerce_subscription_proto_goTypes,
		DependencyIndexes: file_eolymp_commerce_subscription_proto_depIdxs,
		EnumInfos:         file_eolymp_commerce_subscription_proto_enumTypes,
		MessageInfos:      file_eolymp_commerce_subscription_proto_msgTypes,
	}.Build()
	File_eolymp_commerce_subscription_proto = out.File
	file_eolymp_commerce_subscription_proto_rawDesc = nil
	file_eolymp_commerce_subscription_proto_goTypes = nil
	file_eolymp_commerce_subscription_proto_depIdxs = nil
}
