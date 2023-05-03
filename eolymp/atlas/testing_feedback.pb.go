// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.18.1
// source: eolymp/atlas/testing_feedback.proto

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

type FeedbackPolicy int32

const (
	FeedbackPolicy_COMPLETE      FeedbackPolicy = 0 // show each test individually
	FeedbackPolicy_ICPC          FeedbackPolicy = 1 // show results as in ICPC contest (only first not accepted test from the set)
	FeedbackPolicy_ICPC_EXPANDED FeedbackPolicy = 2 // show results as in ICPC contest including number of the test which failed
)

// Enum value maps for FeedbackPolicy.
var (
	FeedbackPolicy_name = map[int32]string{
		0: "COMPLETE",
		1: "ICPC",
		2: "ICPC_EXPANDED",
	}
	FeedbackPolicy_value = map[string]int32{
		"COMPLETE":      0,
		"ICPC":          1,
		"ICPC_EXPANDED": 2,
	}
)

func (x FeedbackPolicy) Enum() *FeedbackPolicy {
	p := new(FeedbackPolicy)
	*p = x
	return p
}

func (x FeedbackPolicy) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (FeedbackPolicy) Descriptor() protoreflect.EnumDescriptor {
	return file_eolymp_atlas_testing_feedback_proto_enumTypes[0].Descriptor()
}

func (FeedbackPolicy) Type() protoreflect.EnumType {
	return &file_eolymp_atlas_testing_feedback_proto_enumTypes[0]
}

func (x FeedbackPolicy) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use FeedbackPolicy.Descriptor instead.
func (FeedbackPolicy) EnumDescriptor() ([]byte, []int) {
	return file_eolymp_atlas_testing_feedback_proto_rawDescGZIP(), []int{0}
}

var File_eolymp_atlas_testing_feedback_proto protoreflect.FileDescriptor

var file_eolymp_atlas_testing_feedback_proto_rawDesc = []byte{
	0x0a, 0x23, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x61, 0x74, 0x6c, 0x61, 0x73, 0x2f, 0x74,
	0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x66, 0x65, 0x65, 0x64, 0x62, 0x61, 0x63, 0x6b, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x61, 0x74,
	0x6c, 0x61, 0x73, 0x2a, 0x3b, 0x0a, 0x0e, 0x46, 0x65, 0x65, 0x64, 0x62, 0x61, 0x63, 0x6b, 0x50,
	0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0x0c, 0x0a, 0x08, 0x43, 0x4f, 0x4d, 0x50, 0x4c, 0x45, 0x54,
	0x45, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x49, 0x43, 0x50, 0x43, 0x10, 0x01, 0x12, 0x11, 0x0a,
	0x0d, 0x49, 0x43, 0x50, 0x43, 0x5f, 0x45, 0x58, 0x50, 0x41, 0x4e, 0x44, 0x45, 0x44, 0x10, 0x02,
	0x42, 0x2d, 0x5a, 0x2b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65,
	0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x67, 0x6f, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x65, 0x6f, 0x6c,
	0x79, 0x6d, 0x70, 0x2f, 0x61, 0x74, 0x6c, 0x61, 0x73, 0x3b, 0x61, 0x74, 0x6c, 0x61, 0x73, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_eolymp_atlas_testing_feedback_proto_rawDescOnce sync.Once
	file_eolymp_atlas_testing_feedback_proto_rawDescData = file_eolymp_atlas_testing_feedback_proto_rawDesc
)

func file_eolymp_atlas_testing_feedback_proto_rawDescGZIP() []byte {
	file_eolymp_atlas_testing_feedback_proto_rawDescOnce.Do(func() {
		file_eolymp_atlas_testing_feedback_proto_rawDescData = protoimpl.X.CompressGZIP(file_eolymp_atlas_testing_feedback_proto_rawDescData)
	})
	return file_eolymp_atlas_testing_feedback_proto_rawDescData
}

var file_eolymp_atlas_testing_feedback_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_eolymp_atlas_testing_feedback_proto_goTypes = []interface{}{
	(FeedbackPolicy)(0), // 0: eolymp.atlas.FeedbackPolicy
}
var file_eolymp_atlas_testing_feedback_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_eolymp_atlas_testing_feedback_proto_init() }
func file_eolymp_atlas_testing_feedback_proto_init() {
	if File_eolymp_atlas_testing_feedback_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_eolymp_atlas_testing_feedback_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_eolymp_atlas_testing_feedback_proto_goTypes,
		DependencyIndexes: file_eolymp_atlas_testing_feedback_proto_depIdxs,
		EnumInfos:         file_eolymp_atlas_testing_feedback_proto_enumTypes,
	}.Build()
	File_eolymp_atlas_testing_feedback_proto = out.File
	file_eolymp_atlas_testing_feedback_proto_rawDesc = nil
	file_eolymp_atlas_testing_feedback_proto_goTypes = nil
	file_eolymp_atlas_testing_feedback_proto_depIdxs = nil
}