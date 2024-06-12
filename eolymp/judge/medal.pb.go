// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.21.12
// source: eolymp/judge/medal.proto

package judge

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

type Medal int32

const (
	Medal_NO_MEDAL          Medal = 0
	Medal_GOLD_MEDAL        Medal = 1
	Medal_SILVER_MEDAL      Medal = 2
	Medal_BRONZE_MEDAL      Medal = 3
	Medal_HONORABLE_MENTION Medal = 4
)

// Enum value maps for Medal.
var (
	Medal_name = map[int32]string{
		0: "NO_MEDAL",
		1: "GOLD_MEDAL",
		2: "SILVER_MEDAL",
		3: "BRONZE_MEDAL",
		4: "HONORABLE_MENTION",
	}
	Medal_value = map[string]int32{
		"NO_MEDAL":          0,
		"GOLD_MEDAL":        1,
		"SILVER_MEDAL":      2,
		"BRONZE_MEDAL":      3,
		"HONORABLE_MENTION": 4,
	}
)

func (x Medal) Enum() *Medal {
	p := new(Medal)
	*p = x
	return p
}

func (x Medal) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Medal) Descriptor() protoreflect.EnumDescriptor {
	return file_eolymp_judge_medal_proto_enumTypes[0].Descriptor()
}

func (Medal) Type() protoreflect.EnumType {
	return &file_eolymp_judge_medal_proto_enumTypes[0]
}

func (x Medal) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Medal.Descriptor instead.
func (Medal) EnumDescriptor() ([]byte, []int) {
	return file_eolymp_judge_medal_proto_rawDescGZIP(), []int{0}
}

var File_eolymp_judge_medal_proto protoreflect.FileDescriptor

var file_eolymp_judge_medal_proto_rawDesc = []byte{
	0x0a, 0x18, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x6a, 0x75, 0x64, 0x67, 0x65, 0x2f, 0x6d,
	0x65, 0x64, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x65, 0x6f, 0x6c, 0x79,
	0x6d, 0x70, 0x2e, 0x6a, 0x75, 0x64, 0x67, 0x65, 0x2a, 0x60, 0x0a, 0x05, 0x4d, 0x65, 0x64, 0x61,
	0x6c, 0x12, 0x0c, 0x0a, 0x08, 0x4e, 0x4f, 0x5f, 0x4d, 0x45, 0x44, 0x41, 0x4c, 0x10, 0x00, 0x12,
	0x0e, 0x0a, 0x0a, 0x47, 0x4f, 0x4c, 0x44, 0x5f, 0x4d, 0x45, 0x44, 0x41, 0x4c, 0x10, 0x01, 0x12,
	0x10, 0x0a, 0x0c, 0x53, 0x49, 0x4c, 0x56, 0x45, 0x52, 0x5f, 0x4d, 0x45, 0x44, 0x41, 0x4c, 0x10,
	0x02, 0x12, 0x10, 0x0a, 0x0c, 0x42, 0x52, 0x4f, 0x4e, 0x5a, 0x45, 0x5f, 0x4d, 0x45, 0x44, 0x41,
	0x4c, 0x10, 0x03, 0x12, 0x15, 0x0a, 0x11, 0x48, 0x4f, 0x4e, 0x4f, 0x52, 0x41, 0x42, 0x4c, 0x45,
	0x5f, 0x4d, 0x45, 0x4e, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x04, 0x42, 0x2d, 0x5a, 0x2b, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f,
	0x67, 0x6f, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x6a, 0x75,
	0x64, 0x67, 0x65, 0x3b, 0x6a, 0x75, 0x64, 0x67, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_eolymp_judge_medal_proto_rawDescOnce sync.Once
	file_eolymp_judge_medal_proto_rawDescData = file_eolymp_judge_medal_proto_rawDesc
)

func file_eolymp_judge_medal_proto_rawDescGZIP() []byte {
	file_eolymp_judge_medal_proto_rawDescOnce.Do(func() {
		file_eolymp_judge_medal_proto_rawDescData = protoimpl.X.CompressGZIP(file_eolymp_judge_medal_proto_rawDescData)
	})
	return file_eolymp_judge_medal_proto_rawDescData
}

var file_eolymp_judge_medal_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_eolymp_judge_medal_proto_goTypes = []any{
	(Medal)(0), // 0: eolymp.judge.Medal
}
var file_eolymp_judge_medal_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_eolymp_judge_medal_proto_init() }
func file_eolymp_judge_medal_proto_init() {
	if File_eolymp_judge_medal_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_eolymp_judge_medal_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_eolymp_judge_medal_proto_goTypes,
		DependencyIndexes: file_eolymp_judge_medal_proto_depIdxs,
		EnumInfos:         file_eolymp_judge_medal_proto_enumTypes,
	}.Build()
	File_eolymp_judge_medal_proto = out.File
	file_eolymp_judge_medal_proto_rawDesc = nil
	file_eolymp_judge_medal_proto_goTypes = nil
	file_eolymp_judge_medal_proto_depIdxs = nil
}
