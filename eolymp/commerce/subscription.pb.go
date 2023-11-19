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
	Subscription_UNKNOWN_STATUS Subscription_Status = 0
	Subscription_PENDING        Subscription_Status = 1
	Subscription_ACTIVE         Subscription_Status = 2
	Subscription_PAUSED         Subscription_Status = 3
	Subscription_CANCELLED      Subscription_Status = 4
)

// Enum value maps for Subscription_Status.
var (
	Subscription_Status_name = map[int32]string{
		0: "UNKNOWN_STATUS",
		1: "PENDING",
		2: "ACTIVE",
		3: "PAUSED",
		4: "CANCELLED",
	}
	Subscription_Status_value = map[string]int32{
		"UNKNOWN_STATUS": 0,
		"PENDING":        1,
		"ACTIVE":         2,
		"PAUSED":         3,
		"CANCELLED":      4,
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

type Subscription_Recurrence int32

const (
	Subscription_UNKNOWN_RECURRENCE Subscription_Recurrence = 0
	Subscription_MONTHLY            Subscription_Recurrence = 1
	Subscription_YEARLY             Subscription_Recurrence = 2
)

// Enum value maps for Subscription_Recurrence.
var (
	Subscription_Recurrence_name = map[int32]string{
		0: "UNKNOWN_RECURRENCE",
		1: "MONTHLY",
		2: "YEARLY",
	}
	Subscription_Recurrence_value = map[string]int32{
		"UNKNOWN_RECURRENCE": 0,
		"MONTHLY":            1,
		"YEARLY":             2,
	}
)

func (x Subscription_Recurrence) Enum() *Subscription_Recurrence {
	p := new(Subscription_Recurrence)
	*p = x
	return p
}

func (x Subscription_Recurrence) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Subscription_Recurrence) Descriptor() protoreflect.EnumDescriptor {
	return file_eolymp_commerce_subscription_proto_enumTypes[1].Descriptor()
}

func (Subscription_Recurrence) Type() protoreflect.EnumType {
	return &file_eolymp_commerce_subscription_proto_enumTypes[1]
}

func (x Subscription_Recurrence) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Subscription_Recurrence.Descriptor instead.
func (Subscription_Recurrence) EnumDescriptor() ([]byte, []int) {
	return file_eolymp_commerce_subscription_proto_rawDescGZIP(), []int{0, 1}
}

type Subscription struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      string              `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Status  Subscription_Status `protobuf:"varint,2,opt,name=status,proto3,enum=eolymp.commerce.Subscription_Status" json:"status,omitempty"`
	SpaceId string              `protobuf:"bytes,3,opt,name=space_id,json=spaceId,proto3" json:"space_id,omitempty"`
	// Types that are assignable to Payer:
	//	*Subscription_UserId
	//	*Subscription_MemberId
	Payer                isSubscription_Payer   `protobuf_oneof:"payer"`
	PayerEmail           string                 `protobuf:"bytes,7,opt,name=payer_email,json=payerEmail,proto3" json:"payer_email,omitempty"`
	CreatedAt            *timestamppb.Timestamp `protobuf:"bytes,10,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`       // time when subscription was created
	StartedAt            *timestamppb.Timestamp `protobuf:"bytes,11,opt,name=started_at,json=startedAt,proto3" json:"started_at,omitempty"`       // time when subscription was started (activated first time)
	CancelledAt          *timestamppb.Timestamp `protobuf:"bytes,12,opt,name=cancelled_at,json=cancelledAt,proto3" json:"cancelled_at,omitempty"` // time when subscription was/will be cancelled
	PeriodStart          *timestamppb.Timestamp `protobuf:"bytes,40,opt,name=period_start,json=periodStart,proto3" json:"period_start,omitempty"` // current billing period start (same as last renew date)
	PeriodEnd            *timestamppb.Timestamp `protobuf:"bytes,41,opt,name=period_end,json=periodEnd,proto3" json:"period_end,omitempty"`       // current billing period end (same as renew date for next period)
	PaymentMethod        string                 `protobuf:"bytes,20,opt,name=payment_method,json=paymentMethod,proto3" json:"payment_method,omitempty"`
	CancellationComment  string                 `protobuf:"bytes,50,opt,name=cancellation_comment,json=cancellationComment,proto3" json:"cancellation_comment,omitempty"`
	CancellationFeedback string                 `protobuf:"bytes,51,opt,name=cancellation_feedback,json=cancellationFeedback,proto3" json:"cancellation_feedback,omitempty"`
	CancellationReason   string                 `protobuf:"bytes,52,opt,name=cancellation_reason,json=cancellationReason,proto3" json:"cancellation_reason,omitempty"`
	Currency             string                 `protobuf:"bytes,31,opt,name=currency,proto3" json:"currency,omitempty"`
	TotalAmount          uint32                 `protobuf:"varint,32,opt,name=total_amount,json=totalAmount,proto3" json:"total_amount,omitempty"`
	SuccessUrl           string                 `protobuf:"bytes,60,opt,name=success_url,json=successUrl,proto3" json:"success_url,omitempty"`
	CancelUrl            string                 `protobuf:"bytes,61,opt,name=cancel_url,json=cancelUrl,proto3" json:"cancel_url,omitempty"`
	CheckoutUrl          string                 `protobuf:"bytes,62,opt,name=checkout_url,json=checkoutUrl,proto3" json:"checkout_url,omitempty"`
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

func (x *Subscription) GetSpaceId() string {
	if x != nil {
		return x.SpaceId
	}
	return ""
}

func (m *Subscription) GetPayer() isSubscription_Payer {
	if m != nil {
		return m.Payer
	}
	return nil
}

func (x *Subscription) GetUserId() string {
	if x, ok := x.GetPayer().(*Subscription_UserId); ok {
		return x.UserId
	}
	return ""
}

func (x *Subscription) GetMemberId() string {
	if x, ok := x.GetPayer().(*Subscription_MemberId); ok {
		return x.MemberId
	}
	return ""
}

func (x *Subscription) GetPayerEmail() string {
	if x != nil {
		return x.PayerEmail
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

func (x *Subscription) GetCancelledAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CancelledAt
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

func (x *Subscription) GetPaymentMethod() string {
	if x != nil {
		return x.PaymentMethod
	}
	return ""
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

func (x *Subscription) GetTotalAmount() uint32 {
	if x != nil {
		return x.TotalAmount
	}
	return 0
}

func (x *Subscription) GetSuccessUrl() string {
	if x != nil {
		return x.SuccessUrl
	}
	return ""
}

func (x *Subscription) GetCancelUrl() string {
	if x != nil {
		return x.CancelUrl
	}
	return ""
}

func (x *Subscription) GetCheckoutUrl() string {
	if x != nil {
		return x.CheckoutUrl
	}
	return ""
}

func (x *Subscription) GetItems() []*Subscription_Item {
	if x != nil {
		return x.Items
	}
	return nil
}

type isSubscription_Payer interface {
	isSubscription_Payer()
}

type Subscription_UserId struct {
	UserId string `protobuf:"bytes,5,opt,name=user_id,json=userId,proto3,oneof"`
}

type Subscription_MemberId struct {
	MemberId string `protobuf:"bytes,6,opt,name=member_id,json=memberId,proto3,oneof"`
}

func (*Subscription_UserId) isSubscription_Payer() {}

func (*Subscription_MemberId) isSubscription_Payer() {}

type Subscription_Item struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                 string                  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Active             bool                    `protobuf:"varint,2,opt,name=active,proto3" json:"active,omitempty"`
	ProductId          string                  `protobuf:"bytes,20,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
	ProductName        string                  `protobuf:"bytes,21,opt,name=product_name,json=productName,proto3" json:"product_name,omitempty"`
	ProductImage       string                  `protobuf:"bytes,22,opt,name=product_image,json=productImage,proto3" json:"product_image,omitempty"`
	ProductDescription string                  `protobuf:"bytes,23,opt,name=product_description,json=productDescription,proto3" json:"product_description,omitempty"`
	Recurrence         Subscription_Recurrence `protobuf:"varint,10,opt,name=recurrence,proto3,enum=eolymp.commerce.Subscription_Recurrence" json:"recurrence,omitempty"`
	UnitAmount         uint32                  `protobuf:"varint,11,opt,name=unit_amount,json=unitAmount,proto3" json:"unit_amount,omitempty"`
	Quantity           uint32                  `protobuf:"varint,12,opt,name=quantity,proto3" json:"quantity,omitempty"`
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

func (x *Subscription_Item) GetActive() bool {
	if x != nil {
		return x.Active
	}
	return false
}

func (x *Subscription_Item) GetProductId() string {
	if x != nil {
		return x.ProductId
	}
	return ""
}

func (x *Subscription_Item) GetProductName() string {
	if x != nil {
		return x.ProductName
	}
	return ""
}

func (x *Subscription_Item) GetProductImage() string {
	if x != nil {
		return x.ProductImage
	}
	return ""
}

func (x *Subscription_Item) GetProductDescription() string {
	if x != nil {
		return x.ProductDescription
	}
	return ""
}

func (x *Subscription_Item) GetRecurrence() Subscription_Recurrence {
	if x != nil {
		return x.Recurrence
	}
	return Subscription_UNKNOWN_RECURRENCE
}

func (x *Subscription_Item) GetUnitAmount() uint32 {
	if x != nil {
		return x.UnitAmount
	}
	return 0
}

func (x *Subscription_Item) GetQuantity() uint32 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

var File_eolymp_commerce_subscription_proto protoreflect.FileDescriptor

var file_eolymp_commerce_subscription_proto_rawDesc = []byte{
	0x0a, 0x22, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x63,
	0x65, 0x2f, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x65, 0x72, 0x63, 0x65, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x88, 0x0b, 0x0a, 0x0c, 0x53, 0x75, 0x62, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x3c, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x24, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x63, 0x65, 0x2e, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x70, 0x61, 0x63, 0x65, 0x5f, 0x69,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x70, 0x61, 0x63, 0x65, 0x49, 0x64,
	0x12, 0x19, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x48, 0x00, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x09, 0x6d,
	0x65, 0x6d, 0x62, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00,
	0x52, 0x08, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x61,
	0x79, 0x65, 0x72, 0x5f, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x70, 0x61, 0x79, 0x65, 0x72, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x39, 0x0a, 0x0a, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x65,
	0x64, 0x5f, 0x61, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x12, 0x3d, 0x0a, 0x0c, 0x63, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x6c, 0x65, 0x64, 0x5f, 0x61,
	0x74, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x0b, 0x63, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x6c, 0x65, 0x64, 0x41, 0x74,
	0x12, 0x3d, 0x0a, 0x0c, 0x70, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x5f, 0x73, 0x74, 0x61, 0x72, 0x74,
	0x18, 0x28, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x0b, 0x70, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x53, 0x74, 0x61, 0x72, 0x74, 0x12,
	0x39, 0x0a, 0x0a, 0x70, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x5f, 0x65, 0x6e, 0x64, 0x18, 0x29, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x09, 0x70, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x45, 0x6e, 0x64, 0x12, 0x25, 0x0a, 0x0e, 0x70, 0x61,
	0x79, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0x14, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0d, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x74, 0x68, 0x6f,
	0x64, 0x12, 0x31, 0x0a, 0x14, 0x63, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x6c, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x32, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x13, 0x63, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6d,
	0x6d, 0x65, 0x6e, 0x74, 0x12, 0x33, 0x0a, 0x15, 0x63, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x6c, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x66, 0x65, 0x65, 0x64, 0x62, 0x61, 0x63, 0x6b, 0x18, 0x33, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x14, 0x63, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x6c, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x46, 0x65, 0x65, 0x64, 0x62, 0x61, 0x63, 0x6b, 0x12, 0x2f, 0x0a, 0x13, 0x63, 0x61, 0x6e,
	0x63, 0x65, 0x6c, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e,
	0x18, 0x34, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x63, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x6c, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x75,
	0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x1f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x75,
	0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x21, 0x0a, 0x0c, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f,
	0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x20, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x74, 0x6f,
	0x74, 0x61, 0x6c, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x75, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x3c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x55, 0x72, 0x6c, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x61,
	0x6e, 0x63, 0x65, 0x6c, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x3d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x63, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x55, 0x72, 0x6c, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x68, 0x65,
	0x63, 0x6b, 0x6f, 0x75, 0x74, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x3e, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x6f, 0x75, 0x74, 0x55, 0x72, 0x6c, 0x12, 0x39, 0x0a, 0x05,
	0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0xe7, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x65,
	0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x63, 0x65, 0x2e, 0x53,
	0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x49, 0x74, 0x65, 0x6d,
	0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x1a, 0xcd, 0x02, 0x0a, 0x04, 0x49, 0x74, 0x65, 0x6d,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x16, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x15, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x70, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x16, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0c, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x12,
	0x2f, 0x0a, 0x13, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x17, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x70, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x48, 0x0a, 0x0a, 0x72, 0x65, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x18, 0x0a,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x28, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x65, 0x72, 0x63, 0x65, 0x2e, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x52, 0x0a,
	0x72, 0x65, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x75, 0x6e,
	0x69, 0x74, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x0a, 0x75, 0x6e, 0x69, 0x74, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x71,
	0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x71,
	0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x22, 0x50, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x12, 0x0a, 0x0e, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x5f, 0x53, 0x54, 0x41,
	0x54, 0x55, 0x53, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x50, 0x45, 0x4e, 0x44, 0x49, 0x4e, 0x47,
	0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x41, 0x43, 0x54, 0x49, 0x56, 0x45, 0x10, 0x02, 0x12, 0x0a,
	0x0a, 0x06, 0x50, 0x41, 0x55, 0x53, 0x45, 0x44, 0x10, 0x03, 0x12, 0x0d, 0x0a, 0x09, 0x43, 0x41,
	0x4e, 0x43, 0x45, 0x4c, 0x4c, 0x45, 0x44, 0x10, 0x04, 0x22, 0x3d, 0x0a, 0x0a, 0x52, 0x65, 0x63,
	0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x16, 0x0a, 0x12, 0x55, 0x4e, 0x4b, 0x4e, 0x4f,
	0x57, 0x4e, 0x5f, 0x52, 0x45, 0x43, 0x55, 0x52, 0x52, 0x45, 0x4e, 0x43, 0x45, 0x10, 0x00, 0x12,
	0x0b, 0x0a, 0x07, 0x4d, 0x4f, 0x4e, 0x54, 0x48, 0x4c, 0x59, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06,
	0x59, 0x45, 0x41, 0x52, 0x4c, 0x59, 0x10, 0x02, 0x42, 0x07, 0x0a, 0x05, 0x70, 0x61, 0x79, 0x65,
	0x72, 0x42, 0x33, 0x5a, 0x31, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x67, 0x6f, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x65, 0x6f,
	0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x63, 0x65, 0x3b, 0x63, 0x6f,
	0x6d, 0x6d, 0x65, 0x72, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
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

var file_eolymp_commerce_subscription_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_eolymp_commerce_subscription_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_eolymp_commerce_subscription_proto_goTypes = []interface{}{
	(Subscription_Status)(0),      // 0: eolymp.commerce.Subscription.Status
	(Subscription_Recurrence)(0),  // 1: eolymp.commerce.Subscription.Recurrence
	(*Subscription)(nil),          // 2: eolymp.commerce.Subscription
	(*Subscription_Item)(nil),     // 3: eolymp.commerce.Subscription.Item
	(*timestamppb.Timestamp)(nil), // 4: google.protobuf.Timestamp
}
var file_eolymp_commerce_subscription_proto_depIdxs = []int32{
	0, // 0: eolymp.commerce.Subscription.status:type_name -> eolymp.commerce.Subscription.Status
	4, // 1: eolymp.commerce.Subscription.created_at:type_name -> google.protobuf.Timestamp
	4, // 2: eolymp.commerce.Subscription.started_at:type_name -> google.protobuf.Timestamp
	4, // 3: eolymp.commerce.Subscription.cancelled_at:type_name -> google.protobuf.Timestamp
	4, // 4: eolymp.commerce.Subscription.period_start:type_name -> google.protobuf.Timestamp
	4, // 5: eolymp.commerce.Subscription.period_end:type_name -> google.protobuf.Timestamp
	3, // 6: eolymp.commerce.Subscription.items:type_name -> eolymp.commerce.Subscription.Item
	1, // 7: eolymp.commerce.Subscription.Item.recurrence:type_name -> eolymp.commerce.Subscription.Recurrence
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_eolymp_commerce_subscription_proto_init() }
func file_eolymp_commerce_subscription_proto_init() {
	if File_eolymp_commerce_subscription_proto != nil {
		return
	}
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
	file_eolymp_commerce_subscription_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*Subscription_UserId)(nil),
		(*Subscription_MemberId)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_eolymp_commerce_subscription_proto_rawDesc,
			NumEnums:      2,
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
