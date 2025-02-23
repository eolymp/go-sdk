// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v4.24.4
// source: eolymp/universe/events.proto

package universe

import (
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

type SpaceChangeRecord_Operation int32

const (
	SpaceChangeRecord_NO_OPERATION SpaceChangeRecord_Operation = 0
	SpaceChangeRecord_CREATE       SpaceChangeRecord_Operation = 1
	SpaceChangeRecord_UPDATE       SpaceChangeRecord_Operation = 2
	SpaceChangeRecord_DELETE       SpaceChangeRecord_Operation = 3
	SpaceChangeRecord_SOFT_DELETE  SpaceChangeRecord_Operation = 4
)

// Enum value maps for SpaceChangeRecord_Operation.
var (
	SpaceChangeRecord_Operation_name = map[int32]string{
		0: "NO_OPERATION",
		1: "CREATE",
		2: "UPDATE",
		3: "DELETE",
		4: "SOFT_DELETE",
	}
	SpaceChangeRecord_Operation_value = map[string]int32{
		"NO_OPERATION": 0,
		"CREATE":       1,
		"UPDATE":       2,
		"DELETE":       3,
		"SOFT_DELETE":  4,
	}
)

func (x SpaceChangeRecord_Operation) Enum() *SpaceChangeRecord_Operation {
	p := new(SpaceChangeRecord_Operation)
	*p = x
	return p
}

func (x SpaceChangeRecord_Operation) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SpaceChangeRecord_Operation) Descriptor() protoreflect.EnumDescriptor {
	return file_eolymp_universe_events_proto_enumTypes[0].Descriptor()
}

func (SpaceChangeRecord_Operation) Type() protoreflect.EnumType {
	return &file_eolymp_universe_events_proto_enumTypes[0]
}

func (x SpaceChangeRecord_Operation) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SpaceChangeRecord_Operation.Descriptor instead.
func (SpaceChangeRecord_Operation) EnumDescriptor() ([]byte, []int) {
	return file_eolymp_universe_events_proto_rawDescGZIP(), []int{0, 0}
}

type SpaceChangeRecord struct {
	state         protoimpl.MessageState      `protogen:"open.v1"`
	Op            SpaceChangeRecord_Operation `protobuf:"varint,1,opt,name=op,proto3,enum=eolymp.universe.SpaceChangeRecord_Operation" json:"op,omitempty"`
	Space         *Space                      `protobuf:"bytes,2,opt,name=space,proto3" json:"space,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SpaceChangeRecord) Reset() {
	*x = SpaceChangeRecord{}
	mi := &file_eolymp_universe_events_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SpaceChangeRecord) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SpaceChangeRecord) ProtoMessage() {}

func (x *SpaceChangeRecord) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_universe_events_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SpaceChangeRecord.ProtoReflect.Descriptor instead.
func (*SpaceChangeRecord) Descriptor() ([]byte, []int) {
	return file_eolymp_universe_events_proto_rawDescGZIP(), []int{0}
}

func (x *SpaceChangeRecord) GetOp() SpaceChangeRecord_Operation {
	if x != nil {
		return x.Op
	}
	return SpaceChangeRecord_NO_OPERATION
}

func (x *SpaceChangeRecord) GetSpace() *Space {
	if x != nil {
		return x.Space
	}
	return nil
}

var File_eolymp_universe_events_proto protoreflect.FileDescriptor

var file_eolymp_universe_events_proto_rawDesc = string([]byte{
	0x0a, 0x1c, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x75, 0x6e, 0x69, 0x76, 0x65, 0x72, 0x73,
	0x65, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f,
	0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x75, 0x6e, 0x69, 0x76, 0x65, 0x72, 0x73, 0x65, 0x1a,
	0x1b, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x75, 0x6e, 0x69, 0x76, 0x65, 0x72, 0x73, 0x65,
	0x2f, 0x73, 0x70, 0x61, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd3, 0x01, 0x0a,
	0x11, 0x53, 0x70, 0x61, 0x63, 0x65, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x65, 0x63, 0x6f,
	0x72, 0x64, 0x12, 0x3c, 0x0a, 0x02, 0x6f, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2c,
	0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x75, 0x6e, 0x69, 0x76, 0x65, 0x72, 0x73, 0x65,
	0x2e, 0x53, 0x70, 0x61, 0x63, 0x65, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x65, 0x63, 0x6f,
	0x72, 0x64, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x02, 0x6f, 0x70,
	0x12, 0x2c, 0x0a, 0x05, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x16, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x75, 0x6e, 0x69, 0x76, 0x65, 0x72, 0x73,
	0x65, 0x2e, 0x53, 0x70, 0x61, 0x63, 0x65, 0x52, 0x05, 0x73, 0x70, 0x61, 0x63, 0x65, 0x22, 0x52,
	0x0a, 0x09, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x10, 0x0a, 0x0c, 0x4e,
	0x4f, 0x5f, 0x4f, 0x50, 0x45, 0x52, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x00, 0x12, 0x0a, 0x0a,
	0x06, 0x43, 0x52, 0x45, 0x41, 0x54, 0x45, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x55, 0x50, 0x44,
	0x41, 0x54, 0x45, 0x10, 0x02, 0x12, 0x0a, 0x0a, 0x06, 0x44, 0x45, 0x4c, 0x45, 0x54, 0x45, 0x10,
	0x03, 0x12, 0x0f, 0x0a, 0x0b, 0x53, 0x4f, 0x46, 0x54, 0x5f, 0x44, 0x45, 0x4c, 0x45, 0x54, 0x45,
	0x10, 0x04, 0x42, 0x33, 0x5a, 0x31, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x67, 0x6f, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x65,
	0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x75, 0x6e, 0x69, 0x76, 0x65, 0x72, 0x73, 0x65, 0x3b, 0x75,
	0x6e, 0x69, 0x76, 0x65, 0x72, 0x73, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_eolymp_universe_events_proto_rawDescOnce sync.Once
	file_eolymp_universe_events_proto_rawDescData []byte
)

func file_eolymp_universe_events_proto_rawDescGZIP() []byte {
	file_eolymp_universe_events_proto_rawDescOnce.Do(func() {
		file_eolymp_universe_events_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_eolymp_universe_events_proto_rawDesc), len(file_eolymp_universe_events_proto_rawDesc)))
	})
	return file_eolymp_universe_events_proto_rawDescData
}

var file_eolymp_universe_events_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_eolymp_universe_events_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_eolymp_universe_events_proto_goTypes = []any{
	(SpaceChangeRecord_Operation)(0), // 0: eolymp.universe.SpaceChangeRecord.Operation
	(*SpaceChangeRecord)(nil),        // 1: eolymp.universe.SpaceChangeRecord
	(*Space)(nil),                    // 2: eolymp.universe.Space
}
var file_eolymp_universe_events_proto_depIdxs = []int32{
	0, // 0: eolymp.universe.SpaceChangeRecord.op:type_name -> eolymp.universe.SpaceChangeRecord.Operation
	2, // 1: eolymp.universe.SpaceChangeRecord.space:type_name -> eolymp.universe.Space
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_eolymp_universe_events_proto_init() }
func file_eolymp_universe_events_proto_init() {
	if File_eolymp_universe_events_proto != nil {
		return
	}
	file_eolymp_universe_space_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_eolymp_universe_events_proto_rawDesc), len(file_eolymp_universe_events_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_eolymp_universe_events_proto_goTypes,
		DependencyIndexes: file_eolymp_universe_events_proto_depIdxs,
		EnumInfos:         file_eolymp_universe_events_proto_enumTypes,
		MessageInfos:      file_eolymp_universe_events_proto_msgTypes,
	}.Build()
	File_eolymp_universe_events_proto = out.File
	file_eolymp_universe_events_proto_goTypes = nil
	file_eolymp_universe_events_proto_depIdxs = nil
}
