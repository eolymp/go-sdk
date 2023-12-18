// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.18.1
// source: eolymp/universe/plan.proto

package universe

import (
	commerce "github.com/eolymp/go-sdk/eolymp/commerce"
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

type Plan struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     string            `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name   string            `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Prices []*commerce.Price `protobuf:"bytes,100,rep,name=prices,proto3" json:"prices,omitempty"`
}

func (x *Plan) Reset() {
	*x = Plan{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_universe_plan_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Plan) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Plan) ProtoMessage() {}

func (x *Plan) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_universe_plan_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
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

func (x *Plan) GetPrices() []*commerce.Price {
	if x != nil {
		return x.Prices
	}
	return nil
}

var File_eolymp_universe_plan_proto protoreflect.FileDescriptor

var file_eolymp_universe_plan_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x75, 0x6e, 0x69, 0x76, 0x65, 0x72, 0x73,
	0x65, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x65, 0x6f,
	0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x75, 0x6e, 0x69, 0x76, 0x65, 0x72, 0x73, 0x65, 0x1a, 0x1b, 0x65,
	0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x63, 0x65, 0x2f, 0x70,
	0x72, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa0, 0x01, 0x0a, 0x04, 0x50,
	0x6c, 0x61, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x2e, 0x0a, 0x06, 0x70, 0x72, 0x69, 0x63, 0x65,
	0x73, 0x18, 0x64, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x63, 0x65, 0x2e, 0x50, 0x72, 0x69, 0x63, 0x65, 0x52,
	0x06, 0x70, 0x72, 0x69, 0x63, 0x65, 0x73, 0x22, 0x44, 0x0a, 0x05, 0x45, 0x78, 0x74, 0x72, 0x61,
	0x12, 0x0c, 0x0a, 0x08, 0x4e, 0x4f, 0x5f, 0x45, 0x58, 0x54, 0x52, 0x41, 0x10, 0x00, 0x12, 0x16,
	0x0a, 0x12, 0x44, 0x45, 0x53, 0x43, 0x52, 0x49, 0x50, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x52, 0x45,
	0x4e, 0x44, 0x45, 0x52, 0x10, 0x01, 0x12, 0x15, 0x0a, 0x11, 0x44, 0x45, 0x53, 0x43, 0x52, 0x49,
	0x50, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x56, 0x41, 0x4c, 0x55, 0x45, 0x10, 0x02, 0x42, 0x33, 0x5a,
	0x31, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6f, 0x6c, 0x79,
	0x6d, 0x70, 0x2f, 0x67, 0x6f, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70,
	0x2f, 0x75, 0x6e, 0x69, 0x76, 0x65, 0x72, 0x73, 0x65, 0x3b, 0x75, 0x6e, 0x69, 0x76, 0x65, 0x72,
	0x73, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_eolymp_universe_plan_proto_rawDescOnce sync.Once
	file_eolymp_universe_plan_proto_rawDescData = file_eolymp_universe_plan_proto_rawDesc
)

func file_eolymp_universe_plan_proto_rawDescGZIP() []byte {
	file_eolymp_universe_plan_proto_rawDescOnce.Do(func() {
		file_eolymp_universe_plan_proto_rawDescData = protoimpl.X.CompressGZIP(file_eolymp_universe_plan_proto_rawDescData)
	})
	return file_eolymp_universe_plan_proto_rawDescData
}

var file_eolymp_universe_plan_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_eolymp_universe_plan_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_eolymp_universe_plan_proto_goTypes = []interface{}{
	(Plan_Extra)(0),        // 0: eolymp.universe.Plan.Extra
	(*Plan)(nil),           // 1: eolymp.universe.Plan
	(*commerce.Price)(nil), // 2: eolymp.commerce.Price
}
var file_eolymp_universe_plan_proto_depIdxs = []int32{
	2, // 0: eolymp.universe.Plan.prices:type_name -> eolymp.commerce.Price
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_eolymp_universe_plan_proto_init() }
func file_eolymp_universe_plan_proto_init() {
	if File_eolymp_universe_plan_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_eolymp_universe_plan_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Plan); i {
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
			RawDescriptor: file_eolymp_universe_plan_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_eolymp_universe_plan_proto_goTypes,
		DependencyIndexes: file_eolymp_universe_plan_proto_depIdxs,
		EnumInfos:         file_eolymp_universe_plan_proto_enumTypes,
		MessageInfos:      file_eolymp_universe_plan_proto_msgTypes,
	}.Build()
	File_eolymp_universe_plan_proto = out.File
	file_eolymp_universe_plan_proto_rawDesc = nil
	file_eolymp_universe_plan_proto_goTypes = nil
	file_eolymp_universe_plan_proto_depIdxs = nil
}
