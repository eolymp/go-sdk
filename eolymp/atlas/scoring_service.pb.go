// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.18.1
// source: eolymp/atlas/scoring_service.proto

package atlas

import (
	_ "github.com/eolymp/go-sdk/eolymp/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_eolymp_atlas_scoring_service_proto protoreflect.FileDescriptor

var file_eolymp_atlas_scoring_service_proto_rawDesc = []byte{
	0x0a, 0x22, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x61, 0x74, 0x6c, 0x61, 0x73, 0x2f, 0x73,
	0x63, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x61, 0x74, 0x6c,
	0x61, 0x73, 0x1a, 0x1d, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x18, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x61, 0x74, 0x6c, 0x61, 0x73, 0x2f,
	0x61, 0x74, 0x6c, 0x61, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xbe, 0x03, 0x0a, 0x0e,
	0x53, 0x63, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x8c,
	0x01, 0x0a, 0x0d, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x53, 0x63, 0x6f, 0x72, 0x65,
	0x12, 0x20, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x61, 0x74, 0x6c, 0x61, 0x73, 0x2e,
	0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x49, 0x6e, 0x70,
	0x75, 0x74, 0x1a, 0x21, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x61, 0x74, 0x6c, 0x61,
	0x73, 0x2e, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x4f,
	0x75, 0x74, 0x70, 0x75, 0x74, 0x22, 0x36, 0x82, 0xe3, 0x0a, 0x19, 0x8a, 0xe3, 0x0a, 0x15, 0x61,
	0x74, 0x6c, 0x61, 0x73, 0x3a, 0x73, 0x75, 0x62, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x3a,
	0x72, 0x65, 0x61, 0x64, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x13, 0x12, 0x11, 0x2f, 0x73, 0x63, 0x6f,
	0x72, 0x65, 0x73, 0x2f, 0x7b, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x7d, 0x12, 0x9b, 0x01,
	0x0a, 0x16, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x50, 0x72, 0x6f, 0x62, 0x6c, 0x65,
	0x6d, 0x47, 0x72, 0x61, 0x64, 0x69, 0x6e, 0x67, 0x12, 0x29, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d,
	0x70, 0x2e, 0x61, 0x74, 0x6c, 0x61, 0x73, 0x2e, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65,
	0x50, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x47, 0x72, 0x61, 0x64, 0x69, 0x6e, 0x67, 0x49, 0x6e,
	0x70, 0x75, 0x74, 0x1a, 0x2a, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x61, 0x74, 0x6c,
	0x61, 0x73, 0x2e, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x50, 0x72, 0x6f, 0x62, 0x6c,
	0x65, 0x6d, 0x47, 0x72, 0x61, 0x64, 0x69, 0x6e, 0x67, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x22,
	0x2a, 0x82, 0xe3, 0x0a, 0x16, 0x8a, 0xe3, 0x0a, 0x12, 0x61, 0x74, 0x6c, 0x61, 0x73, 0x3a, 0x70,
	0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x3a, 0x72, 0x65, 0x61, 0x64, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x0a, 0x12, 0x08, 0x2f, 0x67, 0x72, 0x61, 0x64, 0x69, 0x6e, 0x67, 0x12, 0x7f, 0x0a, 0x0e, 0x4c,
	0x69, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x54, 0x6f, 0x70, 0x12, 0x21, 0x2e,
	0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x61, 0x74, 0x6c, 0x61, 0x73, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x50, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x54, 0x6f, 0x70, 0x49, 0x6e, 0x70, 0x75, 0x74,
	0x1a, 0x22, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x61, 0x74, 0x6c, 0x61, 0x73, 0x2e,
	0x4c, 0x69, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x54, 0x6f, 0x70, 0x4f, 0x75,
	0x74, 0x70, 0x75, 0x74, 0x22, 0x26, 0x82, 0xe3, 0x0a, 0x16, 0x8a, 0xe3, 0x0a, 0x12, 0x61, 0x74,
	0x6c, 0x61, 0x73, 0x3a, 0x70, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x3a, 0x72, 0x65, 0x61, 0x64,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x06, 0x12, 0x04, 0x2f, 0x74, 0x6f, 0x70, 0x42, 0x2d, 0x5a, 0x2b,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6f, 0x6c, 0x79, 0x6d,
	0x70, 0x2f, 0x67, 0x6f, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f,
	0x61, 0x74, 0x6c, 0x61, 0x73, 0x3b, 0x61, 0x74, 0x6c, 0x61, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var file_eolymp_atlas_scoring_service_proto_goTypes = []interface{}{
	(*DescribeScoreInput)(nil),           // 0: eolymp.atlas.DescribeScoreInput
	(*DescribeProblemGradingInput)(nil),  // 1: eolymp.atlas.DescribeProblemGradingInput
	(*ListProblemTopInput)(nil),          // 2: eolymp.atlas.ListProblemTopInput
	(*DescribeScoreOutput)(nil),          // 3: eolymp.atlas.DescribeScoreOutput
	(*DescribeProblemGradingOutput)(nil), // 4: eolymp.atlas.DescribeProblemGradingOutput
	(*ListProblemTopOutput)(nil),         // 5: eolymp.atlas.ListProblemTopOutput
}
var file_eolymp_atlas_scoring_service_proto_depIdxs = []int32{
	0, // 0: eolymp.atlas.ScoringService.DescribeScore:input_type -> eolymp.atlas.DescribeScoreInput
	1, // 1: eolymp.atlas.ScoringService.DescribeProblemGrading:input_type -> eolymp.atlas.DescribeProblemGradingInput
	2, // 2: eolymp.atlas.ScoringService.ListProblemTop:input_type -> eolymp.atlas.ListProblemTopInput
	3, // 3: eolymp.atlas.ScoringService.DescribeScore:output_type -> eolymp.atlas.DescribeScoreOutput
	4, // 4: eolymp.atlas.ScoringService.DescribeProblemGrading:output_type -> eolymp.atlas.DescribeProblemGradingOutput
	5, // 5: eolymp.atlas.ScoringService.ListProblemTop:output_type -> eolymp.atlas.ListProblemTopOutput
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_eolymp_atlas_scoring_service_proto_init() }
func file_eolymp_atlas_scoring_service_proto_init() {
	if File_eolymp_atlas_scoring_service_proto != nil {
		return
	}
	file_eolymp_atlas_atlas_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_eolymp_atlas_scoring_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_eolymp_atlas_scoring_service_proto_goTypes,
		DependencyIndexes: file_eolymp_atlas_scoring_service_proto_depIdxs,
	}.Build()
	File_eolymp_atlas_scoring_service_proto = out.File
	file_eolymp_atlas_scoring_service_proto_rawDesc = nil
	file_eolymp_atlas_scoring_service_proto_goTypes = nil
	file_eolymp_atlas_scoring_service_proto_depIdxs = nil
}
