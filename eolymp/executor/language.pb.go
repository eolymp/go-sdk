// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v3.21.12
// source: eolymp/executor/language.proto

package executor

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

type Language_Type int32

const (
	Language_UNKNOWN_TYPE Language_Type = 0
	Language_PROGRAMMING  Language_Type = 1
	Language_SQL          Language_Type = 3
	Language_OTHER        Language_Type = 999
)

// Enum value maps for Language_Type.
var (
	Language_Type_name = map[int32]string{
		0:   "UNKNOWN_TYPE",
		1:   "PROGRAMMING",
		3:   "SQL",
		999: "OTHER",
	}
	Language_Type_value = map[string]int32{
		"UNKNOWN_TYPE": 0,
		"PROGRAMMING":  1,
		"SQL":          3,
		"OTHER":        999,
	}
)

func (x Language_Type) Enum() *Language_Type {
	p := new(Language_Type)
	*p = x
	return p
}

func (x Language_Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Language_Type) Descriptor() protoreflect.EnumDescriptor {
	return file_eolymp_executor_language_proto_enumTypes[0].Descriptor()
}

func (Language_Type) Type() protoreflect.EnumType {
	return &file_eolymp_executor_language_proto_enumTypes[0]
}

func (x Language_Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Language_Type.Descriptor instead.
func (Language_Type) EnumDescriptor() ([]byte, []int) {
	return file_eolymp_executor_language_proto_rawDescGZIP(), []int{0, 0}
}

type Language struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Language unique identifier
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Human friendly name
	Name string `protobuf:"bytes,10,opt,name=name,proto3" json:"name,omitempty"`
	// Language type
	Type Language_Type `protobuf:"varint,12,opt,name=type,proto3,enum=eolymp.executor.Language_Type" json:"type,omitempty"`
	// Deprecated, means this language should be avoided
	Deprecated bool `protobuf:"varint,11,opt,name=deprecated,proto3" json:"deprecated,omitempty"`
}

func (x *Language) Reset() {
	*x = Language{}
	mi := &file_eolymp_executor_language_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Language) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Language) ProtoMessage() {}

func (x *Language) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_executor_language_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Language.ProtoReflect.Descriptor instead.
func (*Language) Descriptor() ([]byte, []int) {
	return file_eolymp_executor_language_proto_rawDescGZIP(), []int{0}
}

func (x *Language) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Language) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Language) GetType() Language_Type {
	if x != nil {
		return x.Type
	}
	return Language_UNKNOWN_TYPE
}

func (x *Language) GetDeprecated() bool {
	if x != nil {
		return x.Deprecated
	}
	return false
}

var File_eolymp_executor_language_proto protoreflect.FileDescriptor

var file_eolymp_executor_language_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x6f,
	0x72, 0x2f, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0f, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x6f,
	0x72, 0x22, 0xc2, 0x01, 0x0a, 0x08, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x32, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x1e, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74,
	0x6f, 0x72, 0x2e, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x2e, 0x54, 0x79, 0x70, 0x65,
	0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x64, 0x65, 0x70, 0x72, 0x65, 0x63,
	0x61, 0x74, 0x65, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x64, 0x65, 0x70, 0x72,
	0x65, 0x63, 0x61, 0x74, 0x65, 0x64, 0x22, 0x3e, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x10,
	0x0a, 0x0c, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x10, 0x00,
	0x12, 0x0f, 0x0a, 0x0b, 0x50, 0x52, 0x4f, 0x47, 0x52, 0x41, 0x4d, 0x4d, 0x49, 0x4e, 0x47, 0x10,
	0x01, 0x12, 0x07, 0x0a, 0x03, 0x53, 0x51, 0x4c, 0x10, 0x03, 0x12, 0x0a, 0x0a, 0x05, 0x4f, 0x54,
	0x48, 0x45, 0x52, 0x10, 0xe7, 0x07, 0x42, 0x33, 0x5a, 0x31, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x67, 0x6f, 0x2d, 0x73,
	0x64, 0x6b, 0x2f, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74,
	0x6f, 0x72, 0x3b, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x6f, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_eolymp_executor_language_proto_rawDescOnce sync.Once
	file_eolymp_executor_language_proto_rawDescData = file_eolymp_executor_language_proto_rawDesc
)

func file_eolymp_executor_language_proto_rawDescGZIP() []byte {
	file_eolymp_executor_language_proto_rawDescOnce.Do(func() {
		file_eolymp_executor_language_proto_rawDescData = protoimpl.X.CompressGZIP(file_eolymp_executor_language_proto_rawDescData)
	})
	return file_eolymp_executor_language_proto_rawDescData
}

var file_eolymp_executor_language_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_eolymp_executor_language_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_eolymp_executor_language_proto_goTypes = []any{
	(Language_Type)(0), // 0: eolymp.executor.Language.Type
	(*Language)(nil),   // 1: eolymp.executor.Language
}
var file_eolymp_executor_language_proto_depIdxs = []int32{
	0, // 0: eolymp.executor.Language.type:type_name -> eolymp.executor.Language.Type
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_eolymp_executor_language_proto_init() }
func file_eolymp_executor_language_proto_init() {
	if File_eolymp_executor_language_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_eolymp_executor_language_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_eolymp_executor_language_proto_goTypes,
		DependencyIndexes: file_eolymp_executor_language_proto_depIdxs,
		EnumInfos:         file_eolymp_executor_language_proto_enumTypes,
		MessageInfos:      file_eolymp_executor_language_proto_msgTypes,
	}.Build()
	File_eolymp_executor_language_proto = out.File
	file_eolymp_executor_language_proto_rawDesc = nil
	file_eolymp_executor_language_proto_goTypes = nil
	file_eolymp_executor_language_proto_depIdxs = nil
}
