// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v4.24.4
// source: eolymp/universe/billing.proto

package universe

import (
	commerce "github.com/eolymp/go-sdk/eolymp/commerce"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
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
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Billing) Reset() {
	*x = Billing{}
	mi := &file_eolymp_universe_billing_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Billing) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Billing) ProtoMessage() {}

func (x *Billing) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_universe_billing_proto_msgTypes[0]
	if x != nil {
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
	state         protoimpl.MessageState `protogen:"open.v1"`
	Name          string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Email         string                 `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	Phone         string                 `protobuf:"bytes,3,opt,name=phone,proto3" json:"phone,omitempty"`
	TaxIdType     string                 `protobuf:"bytes,10,opt,name=tax_id_type,json=taxIdType,proto3" json:"tax_id_type,omitempty"`
	TaxIdValue    string                 `protobuf:"bytes,11,opt,name=tax_id_value,json=taxIdValue,proto3" json:"tax_id_value,omitempty"`
	Address       *commerce.Address      `protobuf:"bytes,20,opt,name=address,proto3" json:"address,omitempty"`
	Currency      string                 `protobuf:"bytes,90,opt,name=currency,proto3" json:"currency,omitempty"`
	Language      string                 `protobuf:"bytes,91,opt,name=language,proto3" json:"language,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Billing_Information) Reset() {
	*x = Billing_Information{}
	mi := &file_eolymp_universe_billing_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Billing_Information) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Billing_Information) ProtoMessage() {}

func (x *Billing_Information) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_universe_billing_proto_msgTypes[1]
	if x != nil {
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
	state  protoimpl.MessageState `protogen:"open.v1"`
	Id     string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Status Billing_Status         `protobuf:"varint,2,opt,name=status,proto3,enum=eolymp.universe.Billing_Status" json:"status,omitempty"`
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
	unknownFields    protoimpl.UnknownFields
	sizeCache        protoimpl.SizeCache
}

func (x *Billing_Subscription) Reset() {
	*x = Billing_Subscription{}
	mi := &file_eolymp_universe_billing_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Billing_Subscription) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Billing_Subscription) ProtoMessage() {}

func (x *Billing_Subscription) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_universe_billing_proto_msgTypes[2]
	if x != nil {
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

const file_eolymp_universe_billing_proto_rawDesc = "" +
	"\n" +
	"\x1deolymp/universe/billing.proto\x12\x0feolymp.universe\x1a\x1deolymp/commerce/address.proto\x1a\x1aeolymp/universe/plan.proto\x1a\x1fgoogle/protobuf/timestamp.proto\"\xc8\t\n" +
	"\aBilling\x1a\xfb\x01\n" +
	"\vInformation\x12\x12\n" +
	"\x04name\x18\x01 \x01(\tR\x04name\x12\x14\n" +
	"\x05email\x18\x02 \x01(\tR\x05email\x12\x14\n" +
	"\x05phone\x18\x03 \x01(\tR\x05phone\x12\x1e\n" +
	"\vtax_id_type\x18\n" +
	" \x01(\tR\ttaxIdType\x12 \n" +
	"\ftax_id_value\x18\v \x01(\tR\n" +
	"taxIdValue\x122\n" +
	"\aaddress\x18\x14 \x01(\v2\x18.eolymp.commerce.AddressR\aaddress\x12\x1a\n" +
	"\bcurrency\x18Z \x01(\tR\bcurrency\x12\x1a\n" +
	"\blanguage\x18[ \x01(\tR\blanguage\x1a\xbb\x06\n" +
	"\fSubscription\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x127\n" +
	"\x06status\x18\x02 \x01(\x0e2\x1f.eolymp.universe.Billing.StatusR\x06status\x12\x1c\n" +
	"\tirregular\x18\x03 \x01(\bR\tirregular\x12,\n" +
	"\x12has_payment_method\x18\x04 \x01(\bR\x10hasPaymentMethod\x12)\n" +
	"\x04plan\x18\n" +
	" \x01(\v2\x15.eolymp.universe.PlanR\x04plan\x127\n" +
	"\avariant\x18\v \x01(\v2\x1d.eolymp.universe.Plan.VariantR\avariant\x12\x14\n" +
	"\x05seats\x18\f \x01(\rR\x05seats\x12:\n" +
	"\n" +
	"created_at\x18\xd2\x01 \x01(\v2\x1a.google.protobuf.TimestampR\tcreatedAt\x12:\n" +
	"\n" +
	"started_at\x18\xd3\x01 \x01(\v2\x1a.google.protobuf.TimestampR\tstartedAt\x128\n" +
	"\tcancel_at\x18\xd5\x01 \x01(\v2\x1a.google.protobuf.TimestampR\bcancelAt\x12>\n" +
	"\fcancelled_at\x18\xd6\x01 \x01(\v2\x1a.google.protobuf.TimestampR\vcancelledAt\x126\n" +
	"\bended_at\x18\xd7\x01 \x01(\v2\x1a.google.protobuf.TimestampR\aendedAt\x12>\n" +
	"\fperiod_start\x18\xf0\x01 \x01(\v2\x1a.google.protobuf.TimestampR\vperiodStart\x12:\n" +
	"\n" +
	"period_end\x18\xf1\x01 \x01(\v2\x1a.google.protobuf.TimestampR\tperiodEnd\x12<\n" +
	"\vtrial_start\x18\x8e\x02 \x01(\v2\x1a.google.protobuf.TimestampR\n" +
	"trialStart\x128\n" +
	"\ttrial_end\x18\x8f\x02 \x01(\v2\x1a.google.protobuf.TimestampR\btrialEnd\"B\n" +
	"\x06Status\x12\x12\n" +
	"\x0eUNKNOWN_STATUS\x10\x00\x12\t\n" +
	"\x05TRIAL\x10\x01\x12\n" +
	"\n" +
	"\x06ACTIVE\x10\x02\x12\r\n" +
	"\tCANCELLED\x10\x03\"=\n" +
	"\n" +
	"Recurrence\x12\x16\n" +
	"\x12UNKNOWN_RECURRENCE\x10\x00\x12\v\n" +
	"\aMONTHLY\x10\x01\x12\n" +
	"\n" +
	"\x06YEARLY\x10\x02B3Z1github.com/eolymp/go-sdk/eolymp/universe;universeb\x06proto3"

var (
	file_eolymp_universe_billing_proto_rawDescOnce sync.Once
	file_eolymp_universe_billing_proto_rawDescData []byte
)

func file_eolymp_universe_billing_proto_rawDescGZIP() []byte {
	file_eolymp_universe_billing_proto_rawDescOnce.Do(func() {
		file_eolymp_universe_billing_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_eolymp_universe_billing_proto_rawDesc), len(file_eolymp_universe_billing_proto_rawDesc)))
	})
	return file_eolymp_universe_billing_proto_rawDescData
}

var file_eolymp_universe_billing_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_eolymp_universe_billing_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_eolymp_universe_billing_proto_goTypes = []any{
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_eolymp_universe_billing_proto_rawDesc), len(file_eolymp_universe_billing_proto_rawDesc)),
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
	file_eolymp_universe_billing_proto_goTypes = nil
	file_eolymp_universe_billing_proto_depIdxs = nil
}
