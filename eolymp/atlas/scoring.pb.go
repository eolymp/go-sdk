// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.12.2
// source: eolymp/atlas/scoring.proto

package atlas

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

type ScoringMode int32

const (
	ScoringMode_NO_SCORE ScoringMode = 0 // no score is awarded for passing tests from this testset
	ScoringMode_EACH     ScoringMode = 1 // scores are awarded for passing each individual test
	ScoringMode_ALL      ScoringMode = 2 // scores are only awarded if all tests are passing
)

// Enum value maps for ScoringMode.
var (
	ScoringMode_name = map[int32]string{
		0: "NO_SCORE",
		1: "EACH",
		2: "ALL",
	}
	ScoringMode_value = map[string]int32{
		"NO_SCORE": 0,
		"EACH":     1,
		"ALL":      2,
	}
)

func (x ScoringMode) Enum() *ScoringMode {
	p := new(ScoringMode)
	*p = x
	return p
}

func (x ScoringMode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ScoringMode) Descriptor() protoreflect.EnumDescriptor {
	return file_eolymp_atlas_scoring_proto_enumTypes[0].Descriptor()
}

func (ScoringMode) Type() protoreflect.EnumType {
	return &file_eolymp_atlas_scoring_proto_enumTypes[0]
}

func (x ScoringMode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ScoringMode.Descriptor instead.
func (ScoringMode) EnumDescriptor() ([]byte, []int) {
	return file_eolymp_atlas_scoring_proto_rawDescGZIP(), []int{0}
}

var File_eolymp_atlas_scoring_proto protoreflect.FileDescriptor

var file_eolymp_atlas_scoring_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x61, 0x74, 0x6c, 0x61, 0x73, 0x2f, 0x73,
	0x63, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x65, 0x6f,
	0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x61, 0x74, 0x6c, 0x61, 0x73, 0x2a, 0x2e, 0x0a, 0x0b, 0x53, 0x63,
	0x6f, 0x72, 0x69, 0x6e, 0x67, 0x4d, 0x6f, 0x64, 0x65, 0x12, 0x0c, 0x0a, 0x08, 0x4e, 0x4f, 0x5f,
	0x53, 0x43, 0x4f, 0x52, 0x45, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x45, 0x41, 0x43, 0x48, 0x10,
	0x01, 0x12, 0x07, 0x0a, 0x03, 0x41, 0x4c, 0x4c, 0x10, 0x02, 0x42, 0x33, 0x5a, 0x31, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f,
	0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x73, 0x2f, 0x67, 0x6f, 0x2f, 0x65, 0x6f, 0x6c,
	0x79, 0x6d, 0x70, 0x2f, 0x61, 0x74, 0x6c, 0x61, 0x73, 0x3b, 0x61, 0x74, 0x6c, 0x61, 0x73, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_eolymp_atlas_scoring_proto_rawDescOnce sync.Once
	file_eolymp_atlas_scoring_proto_rawDescData = file_eolymp_atlas_scoring_proto_rawDesc
)

func file_eolymp_atlas_scoring_proto_rawDescGZIP() []byte {
	file_eolymp_atlas_scoring_proto_rawDescOnce.Do(func() {
		file_eolymp_atlas_scoring_proto_rawDescData = protoimpl.X.CompressGZIP(file_eolymp_atlas_scoring_proto_rawDescData)
	})
	return file_eolymp_atlas_scoring_proto_rawDescData
}

var file_eolymp_atlas_scoring_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_eolymp_atlas_scoring_proto_goTypes = []interface{}{
	(ScoringMode)(0), // 0: eolymp.atlas.ScoringMode
}
var file_eolymp_atlas_scoring_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_eolymp_atlas_scoring_proto_init() }
func file_eolymp_atlas_scoring_proto_init() {
	if File_eolymp_atlas_scoring_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_eolymp_atlas_scoring_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_eolymp_atlas_scoring_proto_goTypes,
		DependencyIndexes: file_eolymp_atlas_scoring_proto_depIdxs,
		EnumInfos:         file_eolymp_atlas_scoring_proto_enumTypes,
	}.Build()
	File_eolymp_atlas_scoring_proto = out.File
	file_eolymp_atlas_scoring_proto_rawDesc = nil
	file_eolymp_atlas_scoring_proto_goTypes = nil
	file_eolymp_atlas_scoring_proto_depIdxs = nil
}
