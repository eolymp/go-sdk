// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v3.21.12
// source: eolymp/commerce/price.proto

package commerce

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Price_Recurrence int32

const (
	Price_UNKNOWN_RECURRENCE Price_Recurrence = 0
	Price_ONETIME            Price_Recurrence = 1
	Price_MONTHLY            Price_Recurrence = 2
	Price_YEARLY             Price_Recurrence = 3
)

// Enum value maps for Price_Recurrence.
var (
	Price_Recurrence_name = map[int32]string{
		0: "UNKNOWN_RECURRENCE",
		1: "ONETIME",
		2: "MONTHLY",
		3: "YEARLY",
	}
	Price_Recurrence_value = map[string]int32{
		"UNKNOWN_RECURRENCE": 0,
		"ONETIME":            1,
		"MONTHLY":            2,
		"YEARLY":             3,
	}
)

func (x Price_Recurrence) Enum() *Price_Recurrence {
	p := new(Price_Recurrence)
	*p = x
	return p
}

func (x Price_Recurrence) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Price_Recurrence) Descriptor() protoreflect.EnumDescriptor {
	return file_eolymp_commerce_price_proto_enumTypes[0].Descriptor()
}

func (Price_Recurrence) Type() protoreflect.EnumType {
	return &file_eolymp_commerce_price_proto_enumTypes[0]
}

func (x Price_Recurrence) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Price_Recurrence.Descriptor instead.
func (Price_Recurrence) EnumDescriptor() ([]byte, []int) {
	return file_eolymp_commerce_price_proto_rawDescGZIP(), []int{0, 0}
}

type Price struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         string           `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Recurrence Price_Recurrence `protobuf:"varint,3,opt,name=recurrence,proto3,enum=eolymp.commerce.Price_Recurrence" json:"recurrence,omitempty"`
	Currency   string           `protobuf:"bytes,30,opt,name=currency,proto3" json:"currency,omitempty"`
	UnitAmount int32            `protobuf:"varint,31,opt,name=unit_amount,json=unitAmount,proto3" json:"unit_amount,omitempty"`
}

func (x *Price) Reset() {
	*x = Price{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_commerce_price_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Price) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Price) ProtoMessage() {}

func (x *Price) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_commerce_price_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Price.ProtoReflect.Descriptor instead.
func (*Price) Descriptor() ([]byte, []int) {
	return file_eolymp_commerce_price_proto_rawDescGZIP(), []int{0}
}

func (x *Price) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Price) GetRecurrence() Price_Recurrence {
	if x != nil {
		return x.Recurrence
	}
	return Price_UNKNOWN_RECURRENCE
}

func (x *Price) GetCurrency() string {
	if x != nil {
		return x.Currency
	}
	return ""
}

func (x *Price) GetUnitAmount() int32 {
	if x != nil {
		return x.UnitAmount
	}
	return 0
}

var File_eolymp_commerce_price_proto protoreflect.FileDescriptor

var file_eolymp_commerce_price_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x63,
	0x65, 0x2f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x65,
	0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x63, 0x65, 0x22, 0xe3,
	0x01, 0x0a, 0x05, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x41, 0x0a, 0x0a, 0x72, 0x65, 0x63, 0x75,
	0x72, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x21, 0x2e, 0x65,
	0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x63, 0x65, 0x2e, 0x50,
	0x72, 0x69, 0x63, 0x65, 0x2e, 0x52, 0x65, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x52,
	0x0a, 0x72, 0x65, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x63,
	0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63,
	0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x1f, 0x0a, 0x0b, 0x75, 0x6e, 0x69, 0x74, 0x5f,
	0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x1f, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x75, 0x6e,
	0x69, 0x74, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x4a, 0x0a, 0x0a, 0x52, 0x65, 0x63, 0x75,
	0x72, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x16, 0x0a, 0x12, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57,
	0x4e, 0x5f, 0x52, 0x45, 0x43, 0x55, 0x52, 0x52, 0x45, 0x4e, 0x43, 0x45, 0x10, 0x00, 0x12, 0x0b,
	0x0a, 0x07, 0x4f, 0x4e, 0x45, 0x54, 0x49, 0x4d, 0x45, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x4d,
	0x4f, 0x4e, 0x54, 0x48, 0x4c, 0x59, 0x10, 0x02, 0x12, 0x0a, 0x0a, 0x06, 0x59, 0x45, 0x41, 0x52,
	0x4c, 0x59, 0x10, 0x03, 0x42, 0x33, 0x5a, 0x31, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x67, 0x6f, 0x2d, 0x73, 0x64, 0x6b,
	0x2f, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x63, 0x65,
	0x3b, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_eolymp_commerce_price_proto_rawDescOnce sync.Once
	file_eolymp_commerce_price_proto_rawDescData = file_eolymp_commerce_price_proto_rawDesc
)

func file_eolymp_commerce_price_proto_rawDescGZIP() []byte {
	file_eolymp_commerce_price_proto_rawDescOnce.Do(func() {
		file_eolymp_commerce_price_proto_rawDescData = protoimpl.X.CompressGZIP(file_eolymp_commerce_price_proto_rawDescData)
	})
	return file_eolymp_commerce_price_proto_rawDescData
}

var file_eolymp_commerce_price_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_eolymp_commerce_price_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_eolymp_commerce_price_proto_goTypes = []interface{}{
	(Price_Recurrence)(0), // 0: eolymp.commerce.Price.Recurrence
	(*Price)(nil),         // 1: eolymp.commerce.Price
}
var file_eolymp_commerce_price_proto_depIdxs = []int32{
	0, // 0: eolymp.commerce.Price.recurrence:type_name -> eolymp.commerce.Price.Recurrence
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_eolymp_commerce_price_proto_init() }
func file_eolymp_commerce_price_proto_init() {
	if File_eolymp_commerce_price_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_eolymp_commerce_price_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Price); i {
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
			RawDescriptor: file_eolymp_commerce_price_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_eolymp_commerce_price_proto_goTypes,
		DependencyIndexes: file_eolymp_commerce_price_proto_depIdxs,
		EnumInfos:         file_eolymp_commerce_price_proto_enumTypes,
		MessageInfos:      file_eolymp_commerce_price_proto_msgTypes,
	}.Build()
	File_eolymp_commerce_price_proto = out.File
	file_eolymp_commerce_price_proto_rawDesc = nil
	file_eolymp_commerce_price_proto_goTypes = nil
	file_eolymp_commerce_price_proto_depIdxs = nil
}
