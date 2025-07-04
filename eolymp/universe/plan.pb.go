// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.29.4
// source: eolymp/universe/plan.proto

package universe

import (
	ecm "github.com/eolymp/go-sdk/eolymp/ecm"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
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

type Plan_Extra int32

const (
	Plan_NO_EXTRA           Plan_Extra = 0
	Plan_DESCRIPTION_RENDER Plan_Extra = 1
	Plan_DESCRIPTION_VALUE  Plan_Extra = 2
)

// Enum value maps for Plan_Extra.
var (
	Plan_Extra_name = map[int32]string{
		0: "NO_EXTRA",
		1: "DESCRIPTION_RENDER",
		2: "DESCRIPTION_VALUE",
	}
	Plan_Extra_value = map[string]int32{
		"NO_EXTRA":           0,
		"DESCRIPTION_RENDER": 1,
		"DESCRIPTION_VALUE":  2,
	}
)

func (x Plan_Extra) Enum() *Plan_Extra {
	p := new(Plan_Extra)
	*p = x
	return p
}

func (x Plan_Extra) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Plan_Extra) Descriptor() protoreflect.EnumDescriptor {
	return file_eolymp_universe_plan_proto_enumTypes[0].Descriptor()
}

func (Plan_Extra) Type() protoreflect.EnumType {
	return &file_eolymp_universe_plan_proto_enumTypes[0]
}

func (x Plan_Extra) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Plan_Extra.Descriptor instead.
func (Plan_Extra) EnumDescriptor() ([]byte, []int) {
	return file_eolymp_universe_plan_proto_rawDescGZIP(), []int{0, 0}
}

type Plan_Recurrence int32

const (
	Plan_UNKNOWN_RECURRENCE Plan_Recurrence = 0
	Plan_ONETIME            Plan_Recurrence = 1
	Plan_MONTHLY            Plan_Recurrence = 2
	Plan_YEARLY             Plan_Recurrence = 3
)

// Enum value maps for Plan_Recurrence.
var (
	Plan_Recurrence_name = map[int32]string{
		0: "UNKNOWN_RECURRENCE",
		1: "ONETIME",
		2: "MONTHLY",
		3: "YEARLY",
	}
	Plan_Recurrence_value = map[string]int32{
		"UNKNOWN_RECURRENCE": 0,
		"ONETIME":            1,
		"MONTHLY":            2,
		"YEARLY":             3,
	}
)

func (x Plan_Recurrence) Enum() *Plan_Recurrence {
	p := new(Plan_Recurrence)
	*p = x
	return p
}

func (x Plan_Recurrence) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Plan_Recurrence) Descriptor() protoreflect.EnumDescriptor {
	return file_eolymp_universe_plan_proto_enumTypes[1].Descriptor()
}

func (Plan_Recurrence) Type() protoreflect.EnumType {
	return &file_eolymp_universe_plan_proto_enumTypes[1]
}

func (x Plan_Recurrence) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Plan_Recurrence.Descriptor instead.
func (Plan_Recurrence) EnumDescriptor() ([]byte, []int) {
	return file_eolymp_universe_plan_proto_rawDescGZIP(), []int{0, 1}
}

type Plan struct {
	state            protoimpl.MessageState `protogen:"open.v1"`
	Id               string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name             string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description      *ecm.Content           `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Quota            *Quota                 `protobuf:"bytes,4,opt,name=quota,proto3" json:"quota,omitempty"`
	Labels           []string               `protobuf:"bytes,5,rep,name=labels,proto3" json:"labels,omitempty"`
	RequiresApproval bool                   `protobuf:"varint,7,opt,name=requires_approval,json=requiresApproval,proto3" json:"requires_approval,omitempty"` // special plan which requires approval
	MinSeats         uint32                 `protobuf:"varint,10,opt,name=min_seats,json=minSeats,proto3" json:"min_seats,omitempty"`
	MaxSeats         uint32                 `protobuf:"varint,11,opt,name=max_seats,json=maxSeats,proto3" json:"max_seats,omitempty"`
	Variants         []*Plan_Variant        `protobuf:"bytes,100,rep,name=variants,proto3" json:"variants,omitempty"`
	unknownFields    protoimpl.UnknownFields
	sizeCache        protoimpl.SizeCache
}

func (x *Plan) Reset() {
	*x = Plan{}
	mi := &file_eolymp_universe_plan_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Plan) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Plan) ProtoMessage() {}

func (x *Plan) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_universe_plan_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Plan.ProtoReflect.Descriptor instead.
func (*Plan) Descriptor() ([]byte, []int) {
	return file_eolymp_universe_plan_proto_rawDescGZIP(), []int{0}
}

func (x *Plan) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Plan) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Plan) GetDescription() *ecm.Content {
	if x != nil {
		return x.Description
	}
	return nil
}

func (x *Plan) GetQuota() *Quota {
	if x != nil {
		return x.Quota
	}
	return nil
}

func (x *Plan) GetLabels() []string {
	if x != nil {
		return x.Labels
	}
	return nil
}

func (x *Plan) GetRequiresApproval() bool {
	if x != nil {
		return x.RequiresApproval
	}
	return false
}

func (x *Plan) GetMinSeats() uint32 {
	if x != nil {
		return x.MinSeats
	}
	return 0
}

func (x *Plan) GetMaxSeats() uint32 {
	if x != nil {
		return x.MaxSeats
	}
	return 0
}

func (x *Plan) GetVariants() []*Plan_Variant {
	if x != nil {
		return x.Variants
	}
	return nil
}

type Plan_Variant struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Recurrence    Plan_Recurrence        `protobuf:"varint,3,opt,name=recurrence,proto3,enum=eolymp.universe.Plan_Recurrence" json:"recurrence,omitempty"`
	Currency      string                 `protobuf:"bytes,30,opt,name=currency,proto3" json:"currency,omitempty"`
	UnitAmount    int32                  `protobuf:"varint,31,opt,name=unit_amount,json=unitAmount,proto3" json:"unit_amount,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Plan_Variant) Reset() {
	*x = Plan_Variant{}
	mi := &file_eolymp_universe_plan_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Plan_Variant) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Plan_Variant) ProtoMessage() {}

func (x *Plan_Variant) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_universe_plan_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Plan_Variant.ProtoReflect.Descriptor instead.
func (*Plan_Variant) Descriptor() ([]byte, []int) {
	return file_eolymp_universe_plan_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Plan_Variant) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Plan_Variant) GetRecurrence() Plan_Recurrence {
	if x != nil {
		return x.Recurrence
	}
	return Plan_UNKNOWN_RECURRENCE
}

func (x *Plan_Variant) GetCurrency() string {
	if x != nil {
		return x.Currency
	}
	return ""
}

func (x *Plan_Variant) GetUnitAmount() int32 {
	if x != nil {
		return x.UnitAmount
	}
	return 0
}

var File_eolymp_universe_plan_proto protoreflect.FileDescriptor

const file_eolymp_universe_plan_proto_rawDesc = "" +
	"\n" +
	"\x1aeolymp/universe/plan.proto\x12\x0feolymp.universe\x1a\x18eolymp/ecm/content.proto\x1a\x1beolymp/universe/quota.proto\"\xf6\x04\n" +
	"\x04Plan\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x12\x12\n" +
	"\x04name\x18\x02 \x01(\tR\x04name\x125\n" +
	"\vdescription\x18\x03 \x01(\v2\x13.eolymp.ecm.ContentR\vdescription\x12,\n" +
	"\x05quota\x18\x04 \x01(\v2\x16.eolymp.universe.QuotaR\x05quota\x12\x16\n" +
	"\x06labels\x18\x05 \x03(\tR\x06labels\x12+\n" +
	"\x11requires_approval\x18\a \x01(\bR\x10requiresApproval\x12\x1b\n" +
	"\tmin_seats\x18\n" +
	" \x01(\rR\bminSeats\x12\x1b\n" +
	"\tmax_seats\x18\v \x01(\rR\bmaxSeats\x129\n" +
	"\bvariants\x18d \x03(\v2\x1d.eolymp.universe.Plan.VariantR\bvariants\x1a\x98\x01\n" +
	"\aVariant\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x12@\n" +
	"\n" +
	"recurrence\x18\x03 \x01(\x0e2 .eolymp.universe.Plan.RecurrenceR\n" +
	"recurrence\x12\x1a\n" +
	"\bcurrency\x18\x1e \x01(\tR\bcurrency\x12\x1f\n" +
	"\vunit_amount\x18\x1f \x01(\x05R\n" +
	"unitAmount\"D\n" +
	"\x05Extra\x12\f\n" +
	"\bNO_EXTRA\x10\x00\x12\x16\n" +
	"\x12DESCRIPTION_RENDER\x10\x01\x12\x15\n" +
	"\x11DESCRIPTION_VALUE\x10\x02\"J\n" +
	"\n" +
	"Recurrence\x12\x16\n" +
	"\x12UNKNOWN_RECURRENCE\x10\x00\x12\v\n" +
	"\aONETIME\x10\x01\x12\v\n" +
	"\aMONTHLY\x10\x02\x12\n" +
	"\n" +
	"\x06YEARLY\x10\x03B3Z1github.com/eolymp/go-sdk/eolymp/universe;universeb\x06proto3"

var (
	file_eolymp_universe_plan_proto_rawDescOnce sync.Once
	file_eolymp_universe_plan_proto_rawDescData []byte
)

func file_eolymp_universe_plan_proto_rawDescGZIP() []byte {
	file_eolymp_universe_plan_proto_rawDescOnce.Do(func() {
		file_eolymp_universe_plan_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_eolymp_universe_plan_proto_rawDesc), len(file_eolymp_universe_plan_proto_rawDesc)))
	})
	return file_eolymp_universe_plan_proto_rawDescData
}

var file_eolymp_universe_plan_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_eolymp_universe_plan_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_eolymp_universe_plan_proto_goTypes = []any{
	(Plan_Extra)(0),      // 0: eolymp.universe.Plan.Extra
	(Plan_Recurrence)(0), // 1: eolymp.universe.Plan.Recurrence
	(*Plan)(nil),         // 2: eolymp.universe.Plan
	(*Plan_Variant)(nil), // 3: eolymp.universe.Plan.Variant
	(*ecm.Content)(nil),  // 4: eolymp.ecm.Content
	(*Quota)(nil),        // 5: eolymp.universe.Quota
}
var file_eolymp_universe_plan_proto_depIdxs = []int32{
	4, // 0: eolymp.universe.Plan.description:type_name -> eolymp.ecm.Content
	5, // 1: eolymp.universe.Plan.quota:type_name -> eolymp.universe.Quota
	3, // 2: eolymp.universe.Plan.variants:type_name -> eolymp.universe.Plan.Variant
	1, // 3: eolymp.universe.Plan.Variant.recurrence:type_name -> eolymp.universe.Plan.Recurrence
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_eolymp_universe_plan_proto_init() }
func file_eolymp_universe_plan_proto_init() {
	if File_eolymp_universe_plan_proto != nil {
		return
	}
	file_eolymp_universe_quota_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_eolymp_universe_plan_proto_rawDesc), len(file_eolymp_universe_plan_proto_rawDesc)),
			NumEnums:      2,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_eolymp_universe_plan_proto_goTypes,
		DependencyIndexes: file_eolymp_universe_plan_proto_depIdxs,
		EnumInfos:         file_eolymp_universe_plan_proto_enumTypes,
		MessageInfos:      file_eolymp_universe_plan_proto_msgTypes,
	}.Build()
	File_eolymp_universe_plan_proto = out.File
	file_eolymp_universe_plan_proto_goTypes = nil
	file_eolymp_universe_plan_proto_depIdxs = nil
}
