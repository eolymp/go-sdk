// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v3.21.12
// source: eolymp/universe/universe.proto

package universe

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

var File_eolymp_universe_universe_proto protoreflect.FileDescriptor

var file_eolymp_universe_universe_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x75, 0x6e, 0x69, 0x76, 0x65, 0x72, 0x73,
	0x65, 0x2f, 0x75, 0x6e, 0x69, 0x76, 0x65, 0x72, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0f, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x75, 0x6e, 0x69, 0x76, 0x65, 0x72, 0x73,
	0x65, 0x1a, 0x1d, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x22, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x72, 0x61, 0x74, 0x65, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x23, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x75, 0x6e, 0x69,
	0x76, 0x65, 0x72, 0x73, 0x65, 0x2f, 0x73, 0x70, 0x61, 0x63, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x8c, 0x09, 0x0a, 0x08, 0x55, 0x6e,
	0x69, 0x76, 0x65, 0x72, 0x73, 0x65, 0x12, 0x83, 0x01, 0x0a, 0x0b, 0x4c, 0x6f, 0x6f, 0x6b, 0x75,
	0x70, 0x53, 0x70, 0x61, 0x63, 0x65, 0x12, 0x21, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e,
	0x75, 0x6e, 0x69, 0x76, 0x65, 0x72, 0x73, 0x65, 0x2e, 0x4c, 0x6f, 0x6f, 0x6b, 0x75, 0x70, 0x53,
	0x70, 0x61, 0x63, 0x65, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x22, 0x2e, 0x65, 0x6f, 0x6c, 0x79,
	0x6d, 0x70, 0x2e, 0x75, 0x6e, 0x69, 0x76, 0x65, 0x72, 0x73, 0x65, 0x2e, 0x4c, 0x6f, 0x6f, 0x6b,
	0x75, 0x70, 0x53, 0x70, 0x61, 0x63, 0x65, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x22, 0x2d, 0xea,
	0xe2, 0x0a, 0x0b, 0xf5, 0xe2, 0x0a, 0x00, 0x00, 0x20, 0x41, 0xf8, 0xe2, 0x0a, 0x64, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x18, 0x12, 0x16, 0x2f, 0x73, 0x70, 0x61, 0x63, 0x65, 0x73, 0x2f, 0x5f, 0x5f,
	0x6c, 0x6f, 0x6f, 0x6b, 0x75, 0x70, 0x2f, 0x7b, 0x6b, 0x65, 0x79, 0x7d, 0x12, 0x90, 0x01, 0x0a,
	0x0b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x70, 0x61, 0x63, 0x65, 0x12, 0x21, 0x2e, 0x65,
	0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x75, 0x6e, 0x69, 0x76, 0x65, 0x72, 0x73, 0x65, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x70, 0x61, 0x63, 0x65, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a,
	0x22, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x75, 0x6e, 0x69, 0x76, 0x65, 0x72, 0x73,
	0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x70, 0x61, 0x63, 0x65, 0x4f, 0x75, 0x74,
	0x70, 0x75, 0x74, 0x22, 0x3a, 0xea, 0xe2, 0x0a, 0x0b, 0xf5, 0xe2, 0x0a, 0x00, 0x00, 0x80, 0x3f,
	0xf8, 0xe2, 0x0a, 0x05, 0x82, 0xe3, 0x0a, 0x18, 0x8a, 0xe3, 0x0a, 0x14, 0x75, 0x6e, 0x69, 0x76,
	0x65, 0x72, 0x73, 0x65, 0x3a, 0x73, 0x70, 0x61, 0x63, 0x65, 0x3a, 0x77, 0x72, 0x69, 0x74, 0x65,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x09, 0x22, 0x07, 0x2f, 0x73, 0x70, 0x61, 0x63, 0x65, 0x73, 0x12,
	0x9b, 0x01, 0x0a, 0x0b, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x70, 0x61, 0x63, 0x65, 0x12,
	0x21, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x75, 0x6e, 0x69, 0x76, 0x65, 0x72, 0x73,
	0x65, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x70, 0x61, 0x63, 0x65, 0x49, 0x6e, 0x70,
	0x75, 0x74, 0x1a, 0x22, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x75, 0x6e, 0x69, 0x76,
	0x65, 0x72, 0x73, 0x65, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x70, 0x61, 0x63, 0x65,
	0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x22, 0x45, 0xea, 0xe2, 0x0a, 0x0b, 0xf5, 0xe2, 0x0a, 0x00,
	0x00, 0x80, 0x3f, 0xf8, 0xe2, 0x0a, 0x05, 0x82, 0xe3, 0x0a, 0x18, 0x8a, 0xe3, 0x0a, 0x14, 0x75,
	0x6e, 0x69, 0x76, 0x65, 0x72, 0x73, 0x65, 0x3a, 0x73, 0x70, 0x61, 0x63, 0x65, 0x3a, 0x77, 0x72,
	0x69, 0x74, 0x65, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x14, 0x1a, 0x12, 0x2f, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x73, 0x2f, 0x7b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x7d, 0x12, 0x9b, 0x01,
	0x0a, 0x0b, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x70, 0x61, 0x63, 0x65, 0x12, 0x21, 0x2e,
	0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x75, 0x6e, 0x69, 0x76, 0x65, 0x72, 0x73, 0x65, 0x2e,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x70, 0x61, 0x63, 0x65, 0x49, 0x6e, 0x70, 0x75, 0x74,
	0x1a, 0x22, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x75, 0x6e, 0x69, 0x76, 0x65, 0x72,
	0x73, 0x65, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x70, 0x61, 0x63, 0x65, 0x4f, 0x75,
	0x74, 0x70, 0x75, 0x74, 0x22, 0x45, 0xea, 0xe2, 0x0a, 0x0b, 0xf5, 0xe2, 0x0a, 0x00, 0x00, 0x00,
	0x40, 0xf8, 0xe2, 0x0a, 0x0a, 0x82, 0xe3, 0x0a, 0x18, 0x8a, 0xe3, 0x0a, 0x14, 0x75, 0x6e, 0x69,
	0x76, 0x65, 0x72, 0x73, 0x65, 0x3a, 0x73, 0x70, 0x61, 0x63, 0x65, 0x3a, 0x77, 0x72, 0x69, 0x74,
	0x65, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x14, 0x2a, 0x12, 0x2f, 0x73, 0x70, 0x61, 0x63, 0x65, 0x73,
	0x2f, 0x7b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x7d, 0x12, 0x85, 0x01, 0x0a, 0x0d,
	0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x53, 0x70, 0x61, 0x63, 0x65, 0x12, 0x23, 0x2e,
	0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x75, 0x6e, 0x69, 0x76, 0x65, 0x72, 0x73, 0x65, 0x2e,
	0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x53, 0x70, 0x61, 0x63, 0x65, 0x49, 0x6e, 0x70,
	0x75, 0x74, 0x1a, 0x24, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x75, 0x6e, 0x69, 0x76,
	0x65, 0x72, 0x73, 0x65, 0x2e, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x53, 0x70, 0x61,
	0x63, 0x65, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x22, 0x29, 0xea, 0xe2, 0x0a, 0x0b, 0xf5, 0xe2,
	0x0a, 0x00, 0x00, 0x20, 0x41, 0xf8, 0xe2, 0x0a, 0x64, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x14, 0x12,
	0x12, 0x2f, 0x73, 0x70, 0x61, 0x63, 0x65, 0x73, 0x2f, 0x7b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x5f,
	0x69, 0x64, 0x7d, 0x12, 0x8b, 0x01, 0x0a, 0x0d, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65,
	0x51, 0x75, 0x6f, 0x74, 0x61, 0x12, 0x23, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x75,
	0x6e, 0x69, 0x76, 0x65, 0x72, 0x73, 0x65, 0x2e, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65,
	0x51, 0x75, 0x6f, 0x74, 0x61, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x24, 0x2e, 0x65, 0x6f, 0x6c,
	0x79, 0x6d, 0x70, 0x2e, 0x75, 0x6e, 0x69, 0x76, 0x65, 0x72, 0x73, 0x65, 0x2e, 0x44, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x62, 0x65, 0x51, 0x75, 0x6f, 0x74, 0x61, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74,
	0x22, 0x2f, 0xea, 0xe2, 0x0a, 0x0b, 0xf5, 0xe2, 0x0a, 0x00, 0x00, 0x20, 0x41, 0xf8, 0xe2, 0x0a,
	0x64, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1a, 0x12, 0x18, 0x2f, 0x73, 0x70, 0x61, 0x63, 0x65, 0x73,
	0x2f, 0x7b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x7d, 0x2f, 0x71, 0x75, 0x6f, 0x74,
	0x61, 0x12, 0x85, 0x01, 0x0a, 0x0b, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x51, 0x75, 0x6f, 0x74,
	0x61, 0x12, 0x21, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x75, 0x6e, 0x69, 0x76, 0x65,
	0x72, 0x73, 0x65, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x51, 0x75, 0x6f, 0x74, 0x61, 0x49,
	0x6e, 0x70, 0x75, 0x74, 0x1a, 0x22, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x75, 0x6e,
	0x69, 0x76, 0x65, 0x72, 0x73, 0x65, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x51, 0x75, 0x6f,
	0x74, 0x61, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x22, 0x2f, 0xea, 0xe2, 0x0a, 0x0b, 0xf5, 0xe2,
	0x0a, 0x00, 0x00, 0x20, 0x41, 0xf8, 0xe2, 0x0a, 0x64, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1a, 0x1a,
	0x18, 0x2f, 0x73, 0x70, 0x61, 0x63, 0x65, 0x73, 0x2f, 0x7b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x5f,
	0x69, 0x64, 0x7d, 0x2f, 0x71, 0x75, 0x6f, 0x74, 0x61, 0x12, 0x8c, 0x01, 0x0a, 0x0a, 0x4c, 0x69,
	0x73, 0x74, 0x53, 0x70, 0x61, 0x63, 0x65, 0x73, 0x12, 0x20, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d,
	0x70, 0x2e, 0x75, 0x6e, 0x69, 0x76, 0x65, 0x72, 0x73, 0x65, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x53,
	0x70, 0x61, 0x63, 0x65, 0x73, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x21, 0x2e, 0x65, 0x6f, 0x6c,
	0x79, 0x6d, 0x70, 0x2e, 0x75, 0x6e, 0x69, 0x76, 0x65, 0x72, 0x73, 0x65, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x53, 0x70, 0x61, 0x63, 0x65, 0x73, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x22, 0x39, 0xea,
	0xe2, 0x0a, 0x0b, 0xf5, 0xe2, 0x0a, 0x00, 0x00, 0xa0, 0x40, 0xf8, 0xe2, 0x0a, 0x14, 0x82, 0xe3,
	0x0a, 0x17, 0x8a, 0xe3, 0x0a, 0x13, 0x75, 0x6e, 0x69, 0x76, 0x65, 0x72, 0x73, 0x65, 0x3a, 0x73,
	0x70, 0x61, 0x63, 0x65, 0x3a, 0x72, 0x65, 0x61, 0x64, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x09, 0x12,
	0x07, 0x2f, 0x73, 0x70, 0x61, 0x63, 0x65, 0x73, 0x42, 0x33, 0x5a, 0x31, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x67, 0x6f,
	0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x75, 0x6e, 0x69, 0x76,
	0x65, 0x72, 0x73, 0x65, 0x3b, 0x75, 0x6e, 0x69, 0x76, 0x65, 0x72, 0x73, 0x65, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_eolymp_universe_universe_proto_goTypes = []any{
	(*LookupSpaceInput)(nil),    // 0: eolymp.universe.LookupSpaceInput
	(*CreateSpaceInput)(nil),    // 1: eolymp.universe.CreateSpaceInput
	(*UpdateSpaceInput)(nil),    // 2: eolymp.universe.UpdateSpaceInput
	(*DeleteSpaceInput)(nil),    // 3: eolymp.universe.DeleteSpaceInput
	(*DescribeSpaceInput)(nil),  // 4: eolymp.universe.DescribeSpaceInput
	(*DescribeQuotaInput)(nil),  // 5: eolymp.universe.DescribeQuotaInput
	(*UpdateQuotaInput)(nil),    // 6: eolymp.universe.UpdateQuotaInput
	(*ListSpacesInput)(nil),     // 7: eolymp.universe.ListSpacesInput
	(*LookupSpaceOutput)(nil),   // 8: eolymp.universe.LookupSpaceOutput
	(*CreateSpaceOutput)(nil),   // 9: eolymp.universe.CreateSpaceOutput
	(*UpdateSpaceOutput)(nil),   // 10: eolymp.universe.UpdateSpaceOutput
	(*DeleteSpaceOutput)(nil),   // 11: eolymp.universe.DeleteSpaceOutput
	(*DescribeSpaceOutput)(nil), // 12: eolymp.universe.DescribeSpaceOutput
	(*DescribeQuotaOutput)(nil), // 13: eolymp.universe.DescribeQuotaOutput
	(*UpdateQuotaOutput)(nil),   // 14: eolymp.universe.UpdateQuotaOutput
	(*ListSpacesOutput)(nil),    // 15: eolymp.universe.ListSpacesOutput
}
var file_eolymp_universe_universe_proto_depIdxs = []int32{
	0,  // 0: eolymp.universe.Universe.LookupSpace:input_type -> eolymp.universe.LookupSpaceInput
	1,  // 1: eolymp.universe.Universe.CreateSpace:input_type -> eolymp.universe.CreateSpaceInput
	2,  // 2: eolymp.universe.Universe.UpdateSpace:input_type -> eolymp.universe.UpdateSpaceInput
	3,  // 3: eolymp.universe.Universe.DeleteSpace:input_type -> eolymp.universe.DeleteSpaceInput
	4,  // 4: eolymp.universe.Universe.DescribeSpace:input_type -> eolymp.universe.DescribeSpaceInput
	5,  // 5: eolymp.universe.Universe.DescribeQuota:input_type -> eolymp.universe.DescribeQuotaInput
	6,  // 6: eolymp.universe.Universe.UpdateQuota:input_type -> eolymp.universe.UpdateQuotaInput
	7,  // 7: eolymp.universe.Universe.ListSpaces:input_type -> eolymp.universe.ListSpacesInput
	8,  // 8: eolymp.universe.Universe.LookupSpace:output_type -> eolymp.universe.LookupSpaceOutput
	9,  // 9: eolymp.universe.Universe.CreateSpace:output_type -> eolymp.universe.CreateSpaceOutput
	10, // 10: eolymp.universe.Universe.UpdateSpace:output_type -> eolymp.universe.UpdateSpaceOutput
	11, // 11: eolymp.universe.Universe.DeleteSpace:output_type -> eolymp.universe.DeleteSpaceOutput
	12, // 12: eolymp.universe.Universe.DescribeSpace:output_type -> eolymp.universe.DescribeSpaceOutput
	13, // 13: eolymp.universe.Universe.DescribeQuota:output_type -> eolymp.universe.DescribeQuotaOutput
	14, // 14: eolymp.universe.Universe.UpdateQuota:output_type -> eolymp.universe.UpdateQuotaOutput
	15, // 15: eolymp.universe.Universe.ListSpaces:output_type -> eolymp.universe.ListSpacesOutput
	8,  // [8:16] is the sub-list for method output_type
	0,  // [0:8] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_eolymp_universe_universe_proto_init() }
func file_eolymp_universe_universe_proto_init() {
	if File_eolymp_universe_universe_proto != nil {
		return
	}
	file_eolymp_universe_space_service_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_eolymp_universe_universe_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_eolymp_universe_universe_proto_goTypes,
		DependencyIndexes: file_eolymp_universe_universe_proto_depIdxs,
	}.Build()
	File_eolymp_universe_universe_proto = out.File
	file_eolymp_universe_universe_proto_rawDesc = nil
	file_eolymp_universe_universe_proto_goTypes = nil
	file_eolymp_universe_universe_proto_depIdxs = nil
}
