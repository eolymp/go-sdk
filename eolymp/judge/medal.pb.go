// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v4.24.4
// source: eolymp/judge/medal.proto

package judge

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

const file_eolymp_judge_medal_proto_rawDesc = "" +
	"\n" +
	"\x18eolymp/judge/medal.proto\x12\feolymp.judge*`\n" +
	"\x05Medal\x12\f\n" +
	"\bNO_MEDAL\x10\x00\x12\x0e\n" +
	"\n" +
	"GOLD_MEDAL\x10\x01\x12\x10\n" +
	"\fSILVER_MEDAL\x10\x02\x12\x10\n" +
	"\fBRONZE_MEDAL\x10\x03\x12\x15\n" +
	"\x11HONORABLE_MENTION\x10\x04B-Z+github.com/eolymp/go-sdk/eolymp/judge;judgeb\x06proto3"

var (
	file_eolymp_judge_medal_proto_rawDescOnce sync.Once
	file_eolymp_judge_medal_proto_rawDescData []byte
)

func file_eolymp_judge_medal_proto_rawDescGZIP() []byte {
	file_eolymp_judge_medal_proto_rawDescOnce.Do(func() {
		file_eolymp_judge_medal_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_eolymp_judge_medal_proto_rawDesc), len(file_eolymp_judge_medal_proto_rawDesc)))
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
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_eolymp_judge_medal_proto_rawDesc), len(file_eolymp_judge_medal_proto_rawDesc)),
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
	file_eolymp_judge_medal_proto_goTypes = nil
	file_eolymp_judge_medal_proto_depIdxs = nil
}
