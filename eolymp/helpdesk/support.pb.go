// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.4
// 	protoc        v4.24.4
// source: eolymp/helpdesk/support.proto

package helpdesk

import (
	_ "github.com/eolymp/go-sdk/eolymp/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_eolymp_helpdesk_support_proto protoreflect.FileDescriptor

var file_eolymp_helpdesk_support_proto_rawDesc = string([]byte{
	0x0a, 0x1d, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x68, 0x65, 0x6c, 0x70, 0x64, 0x65, 0x73,
	0x6b, 0x2f, 0x73, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0f, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x68, 0x65, 0x6c, 0x70, 0x64, 0x65, 0x73, 0x6b,
	0x1a, 0x1d, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x22, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2f, 0x72, 0x61, 0x74, 0x65, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x28, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x68, 0x65, 0x6c, 0x70,
	0x64, 0x65, 0x73, 0x6b, 0x2f, 0x61, 0x75, 0x74, 0x6f, 0x5f, 0x72, 0x65, 0x70, 0x6c, 0x79, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x24, 0x65,
	0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x68, 0x65, 0x6c, 0x70, 0x64, 0x65, 0x73, 0x6b, 0x2f, 0x74,
	0x69, 0x63, 0x6b, 0x65, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x32, 0xfd, 0x1a, 0x0a, 0x07, 0x53, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x12,
	0x9e, 0x01, 0x0a, 0x0c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74,
	0x12, 0x22, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x68, 0x65, 0x6c, 0x70, 0x64, 0x65,
	0x73, 0x6b, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x49,
	0x6e, 0x70, 0x75, 0x74, 0x1a, 0x23, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x68, 0x65,
	0x6c, 0x70, 0x64, 0x65, 0x73, 0x6b, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x63,
	0x6b, 0x65, 0x74, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x22, 0x45, 0xea, 0xe2, 0x0a, 0x0b, 0xf5,
	0xe2, 0x0a, 0x00, 0x00, 0x00, 0x40, 0xf8, 0xe2, 0x0a, 0x0a, 0x82, 0xe3, 0x0a, 0x19, 0x8a, 0xe3,
	0x0a, 0x15, 0x68, 0x65, 0x6c, 0x70, 0x64, 0x65, 0x73, 0x6b, 0x3a, 0x74, 0x69, 0x63, 0x6b, 0x65,
	0x74, 0x3a, 0x77, 0x72, 0x69, 0x74, 0x65, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x13, 0x22, 0x11, 0x2f,
	0x68, 0x65, 0x6c, 0x70, 0x64, 0x65, 0x73, 0x6b, 0x2f, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x73,
	0x12, 0xaa, 0x01, 0x0a, 0x0c, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x63, 0x6b, 0x65,
	0x74, 0x12, 0x22, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x68, 0x65, 0x6c, 0x70, 0x64,
	0x65, 0x73, 0x6b, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74,
	0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x23, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x68,
	0x65, 0x6c, 0x70, 0x64, 0x65, 0x73, 0x6b, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69,
	0x63, 0x6b, 0x65, 0x74, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x22, 0x51, 0xea, 0xe2, 0x0a, 0x0b,
	0xf5, 0xe2, 0x0a, 0x00, 0x00, 0x00, 0x40, 0xf8, 0xe2, 0x0a, 0x0a, 0x82, 0xe3, 0x0a, 0x19, 0x8a,
	0xe3, 0x0a, 0x15, 0x68, 0x65, 0x6c, 0x70, 0x64, 0x65, 0x73, 0x6b, 0x3a, 0x74, 0x69, 0x63, 0x6b,
	0x65, 0x74, 0x3a, 0x77, 0x72, 0x69, 0x74, 0x65, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1f, 0x1a, 0x1d,
	0x2f, 0x68, 0x65, 0x6c, 0x70, 0x64, 0x65, 0x73, 0x6b, 0x2f, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74,
	0x73, 0x2f, 0x7b, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x5f, 0x69, 0x64, 0x7d, 0x12, 0xaa, 0x01,
	0x0a, 0x0c, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x22,
	0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x68, 0x65, 0x6c, 0x70, 0x64, 0x65, 0x73, 0x6b,
	0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x49, 0x6e, 0x70,
	0x75, 0x74, 0x1a, 0x23, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x68, 0x65, 0x6c, 0x70,
	0x64, 0x65, 0x73, 0x6b, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x69, 0x63, 0x6b, 0x65,
	0x74, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x22, 0x51, 0xea, 0xe2, 0x0a, 0x0b, 0xf5, 0xe2, 0x0a,
	0x00, 0x00, 0x00, 0x40, 0xf8, 0xe2, 0x0a, 0x0a, 0x82, 0xe3, 0x0a, 0x19, 0x8a, 0xe3, 0x0a, 0x15,
	0x68, 0x65, 0x6c, 0x70, 0x64, 0x65, 0x73, 0x6b, 0x3a, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x3a,
	0x77, 0x72, 0x69, 0x74, 0x65, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1f, 0x2a, 0x1d, 0x2f, 0x68, 0x65,
	0x6c, 0x70, 0x64, 0x65, 0x73, 0x6b, 0x2f, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x2f, 0x7b,
	0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x5f, 0x69, 0x64, 0x7d, 0x12, 0xaf, 0x01, 0x0a, 0x0e, 0x44,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x24, 0x2e,
	0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x68, 0x65, 0x6c, 0x70, 0x64, 0x65, 0x73, 0x6b, 0x2e,
	0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x49, 0x6e,
	0x70, 0x75, 0x74, 0x1a, 0x25, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x68, 0x65, 0x6c,
	0x70, 0x64, 0x65, 0x73, 0x6b, 0x2e, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x54, 0x69,
	0x63, 0x6b, 0x65, 0x74, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x22, 0x50, 0xea, 0xe2, 0x0a, 0x0b,
	0xf5, 0xe2, 0x0a, 0x00, 0x00, 0x00, 0x40, 0xf8, 0xe2, 0x0a, 0x0a, 0x82, 0xe3, 0x0a, 0x18, 0x8a,
	0xe3, 0x0a, 0x14, 0x68, 0x65, 0x6c, 0x70, 0x64, 0x65, 0x73, 0x6b, 0x3a, 0x74, 0x69, 0x63, 0x6b,
	0x65, 0x74, 0x3a, 0x72, 0x65, 0x61, 0x64, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1f, 0x12, 0x1d, 0x2f,
	0x68, 0x65, 0x6c, 0x70, 0x64, 0x65, 0x73, 0x6b, 0x2f, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x73,
	0x2f, 0x7b, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x5f, 0x69, 0x64, 0x7d, 0x12, 0x9a, 0x01, 0x0a,
	0x0b, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x12, 0x21, 0x2e, 0x65,
	0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x68, 0x65, 0x6c, 0x70, 0x64, 0x65, 0x73, 0x6b, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a,
	0x22, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x68, 0x65, 0x6c, 0x70, 0x64, 0x65, 0x73,
	0x6b, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x4f, 0x75, 0x74,
	0x70, 0x75, 0x74, 0x22, 0x44, 0xea, 0xe2, 0x0a, 0x0b, 0xf5, 0xe2, 0x0a, 0x00, 0x00, 0x00, 0x40,
	0xf8, 0xe2, 0x0a, 0x0a, 0x82, 0xe3, 0x0a, 0x18, 0x8a, 0xe3, 0x0a, 0x14, 0x68, 0x65, 0x6c, 0x70,
	0x64, 0x65, 0x73, 0x6b, 0x3a, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x3a, 0x72, 0x65, 0x61, 0x64,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x13, 0x12, 0x11, 0x2f, 0x68, 0x65, 0x6c, 0x70, 0x64, 0x65, 0x73,
	0x6b, 0x2f, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x12, 0xb5, 0x01, 0x0a, 0x0d, 0x41, 0x70,
	0x70, 0x72, 0x6f, 0x76, 0x65, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x23, 0x2e, 0x65, 0x6f,
	0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x68, 0x65, 0x6c, 0x70, 0x64, 0x65, 0x73, 0x6b, 0x2e, 0x41, 0x70,
	0x70, 0x72, 0x6f, 0x76, 0x65, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x49, 0x6e, 0x70, 0x75, 0x74,
	0x1a, 0x24, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x68, 0x65, 0x6c, 0x70, 0x64, 0x65,
	0x73, 0x6b, 0x2e, 0x41, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x65, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74,
	0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x22, 0x59, 0xea, 0xe2, 0x0a, 0x0b, 0xf5, 0xe2, 0x0a, 0x00,
	0x00, 0x00, 0x40, 0xf8, 0xe2, 0x0a, 0x0a, 0x82, 0xe3, 0x0a, 0x19, 0x8a, 0xe3, 0x0a, 0x15, 0x68,
	0x65, 0x6c, 0x70, 0x64, 0x65, 0x73, 0x6b, 0x3a, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x3a, 0x77,
	0x72, 0x69, 0x74, 0x65, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x27, 0x22, 0x25, 0x2f, 0x68, 0x65, 0x6c,
	0x70, 0x64, 0x65, 0x73, 0x6b, 0x2f, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x2f, 0x7b, 0x74,
	0x69, 0x63, 0x6b, 0x65, 0x74, 0x5f, 0x69, 0x64, 0x7d, 0x2f, 0x61, 0x70, 0x70, 0x72, 0x6f, 0x76,
	0x65, 0x12, 0xb1, 0x01, 0x0a, 0x0c, 0x52, 0x65, 0x6a, 0x65, 0x63, 0x74, 0x54, 0x69, 0x63, 0x6b,
	0x65, 0x74, 0x12, 0x22, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x68, 0x65, 0x6c, 0x70,
	0x64, 0x65, 0x73, 0x6b, 0x2e, 0x52, 0x65, 0x6a, 0x65, 0x63, 0x74, 0x54, 0x69, 0x63, 0x6b, 0x65,
	0x74, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x23, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e,
	0x68, 0x65, 0x6c, 0x70, 0x64, 0x65, 0x73, 0x6b, 0x2e, 0x52, 0x65, 0x6a, 0x65, 0x63, 0x74, 0x54,
	0x69, 0x63, 0x6b, 0x65, 0x74, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x22, 0x58, 0xea, 0xe2, 0x0a,
	0x0b, 0xf5, 0xe2, 0x0a, 0x00, 0x00, 0x00, 0x40, 0xf8, 0xe2, 0x0a, 0x0a, 0x82, 0xe3, 0x0a, 0x19,
	0x8a, 0xe3, 0x0a, 0x15, 0x68, 0x65, 0x6c, 0x70, 0x64, 0x65, 0x73, 0x6b, 0x3a, 0x74, 0x69, 0x63,
	0x6b, 0x65, 0x74, 0x3a, 0x77, 0x72, 0x69, 0x74, 0x65, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x26, 0x22,
	0x24, 0x2f, 0x68, 0x65, 0x6c, 0x70, 0x64, 0x65, 0x73, 0x6b, 0x2f, 0x74, 0x69, 0x63, 0x6b, 0x65,
	0x74, 0x73, 0x2f, 0x7b, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x5f, 0x69, 0x64, 0x7d, 0x2f, 0x72,
	0x65, 0x6a, 0x65, 0x63, 0x74, 0x12, 0xad, 0x01, 0x0a, 0x0b, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x54,
	0x69, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x21, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x68,
	0x65, 0x6c, 0x70, 0x64, 0x65, 0x73, 0x6b, 0x2e, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x54, 0x69, 0x63,
	0x6b, 0x65, 0x74, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x22, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d,
	0x70, 0x2e, 0x68, 0x65, 0x6c, 0x70, 0x64, 0x65, 0x73, 0x6b, 0x2e, 0x43, 0x6c, 0x6f, 0x73, 0x65,
	0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x22, 0x57, 0xea, 0xe2,
	0x0a, 0x0b, 0xf5, 0xe2, 0x0a, 0x00, 0x00, 0x00, 0x40, 0xf8, 0xe2, 0x0a, 0x0a, 0x82, 0xe3, 0x0a,
	0x19, 0x8a, 0xe3, 0x0a, 0x15, 0x68, 0x65, 0x6c, 0x70, 0x64, 0x65, 0x73, 0x6b, 0x3a, 0x74, 0x69,
	0x63, 0x6b, 0x65, 0x74, 0x3a, 0x77, 0x72, 0x69, 0x74, 0x65, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x25,
	0x22, 0x23, 0x2f, 0x68, 0x65, 0x6c, 0x70, 0x64, 0x65, 0x73, 0x6b, 0x2f, 0x74, 0x69, 0x63, 0x6b,
	0x65, 0x74, 0x73, 0x2f, 0x7b, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x5f, 0x69, 0x64, 0x7d, 0x2f,
	0x63, 0x6c, 0x6f, 0x73, 0x65, 0x12, 0xad, 0x01, 0x0a, 0x0a, 0x41, 0x64, 0x64, 0x43, 0x6f, 0x6d,
	0x6d, 0x65, 0x6e, 0x74, 0x12, 0x20, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x68, 0x65,
	0x6c, 0x70, 0x64, 0x65, 0x73, 0x6b, 0x2e, 0x41, 0x64, 0x64, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e,
	0x74, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x21, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e,
	0x68, 0x65, 0x6c, 0x70, 0x64, 0x65, 0x73, 0x6b, 0x2e, 0x41, 0x64, 0x64, 0x43, 0x6f, 0x6d, 0x6d,
	0x65, 0x6e, 0x74, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x22, 0x5a, 0xea, 0xe2, 0x0a, 0x0b, 0xf5,
	0xe2, 0x0a, 0x00, 0x00, 0x00, 0x40, 0xf8, 0xe2, 0x0a, 0x0a, 0x82, 0xe3, 0x0a, 0x19, 0x8a, 0xe3,
	0x0a, 0x15, 0x68, 0x65, 0x6c, 0x70, 0x64, 0x65, 0x73, 0x6b, 0x3a, 0x74, 0x69, 0x63, 0x6b, 0x65,
	0x74, 0x3a, 0x77, 0x72, 0x69, 0x74, 0x65, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x28, 0x22, 0x26, 0x2f,
	0x68, 0x65, 0x6c, 0x70, 0x64, 0x65, 0x73, 0x6b, 0x2f, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x73,
	0x2f, 0x7b, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x5f, 0x69, 0x64, 0x7d, 0x2f, 0x63, 0x6f, 0x6d,
	0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12, 0xc3, 0x01, 0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x23, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70,
	0x2e, 0x68, 0x65, 0x6c, 0x70, 0x64, 0x65, 0x73, 0x6b, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x24, 0x2e, 0x65,
	0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x68, 0x65, 0x6c, 0x70, 0x64, 0x65, 0x73, 0x6b, 0x2e, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x4f, 0x75, 0x74, 0x70,
	0x75, 0x74, 0x22, 0x67, 0xea, 0xe2, 0x0a, 0x0b, 0xf5, 0xe2, 0x0a, 0x00, 0x00, 0x00, 0x40, 0xf8,
	0xe2, 0x0a, 0x0a, 0x82, 0xe3, 0x0a, 0x19, 0x8a, 0xe3, 0x0a, 0x15, 0x68, 0x65, 0x6c, 0x70, 0x64,
	0x65, 0x73, 0x6b, 0x3a, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x3a, 0x77, 0x72, 0x69, 0x74, 0x65,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x35, 0x1a, 0x33, 0x2f, 0x68, 0x65, 0x6c, 0x70, 0x64, 0x65, 0x73,
	0x6b, 0x2f, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x2f, 0x7b, 0x74, 0x69, 0x63, 0x6b, 0x65,
	0x74, 0x5f, 0x69, 0x64, 0x7d, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x7b,
	0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x7d, 0x12, 0xc3, 0x01, 0x0a, 0x0d,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x23, 0x2e,
	0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x68, 0x65, 0x6c, 0x70, 0x64, 0x65, 0x73, 0x6b, 0x2e,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x70,
	0x75, 0x74, 0x1a, 0x24, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x68, 0x65, 0x6c, 0x70,
	0x64, 0x65, 0x73, 0x6b, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x65,
	0x6e, 0x74, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x22, 0x67, 0xea, 0xe2, 0x0a, 0x0b, 0xf5, 0xe2,
	0x0a, 0x00, 0x00, 0x00, 0x40, 0xf8, 0xe2, 0x0a, 0x0a, 0x82, 0xe3, 0x0a, 0x19, 0x8a, 0xe3, 0x0a,
	0x15, 0x68, 0x65, 0x6c, 0x70, 0x64, 0x65, 0x73, 0x6b, 0x3a, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74,
	0x3a, 0x77, 0x72, 0x69, 0x74, 0x65, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x35, 0x2a, 0x33, 0x2f, 0x68,
	0x65, 0x6c, 0x70, 0x64, 0x65, 0x73, 0x6b, 0x2f, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x2f,
	0x7b, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x5f, 0x69, 0x64, 0x7d, 0x2f, 0x63, 0x6f, 0x6d, 0x6d,
	0x65, 0x6e, 0x74, 0x73, 0x2f, 0x7b, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64,
	0x7d, 0x12, 0xb2, 0x01, 0x0a, 0x0c, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e,
	0x74, 0x73, 0x12, 0x22, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x68, 0x65, 0x6c, 0x70,
	0x64, 0x65, 0x73, 0x6b, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74,
	0x73, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x23, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e,
	0x68, 0x65, 0x6c, 0x70, 0x64, 0x65, 0x73, 0x6b, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f, 0x6d,
	0x6d, 0x65, 0x6e, 0x74, 0x73, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x22, 0x59, 0xea, 0xe2, 0x0a,
	0x0b, 0xf5, 0xe2, 0x0a, 0x00, 0x00, 0xa0, 0x40, 0xf8, 0xe2, 0x0a, 0x14, 0x82, 0xe3, 0x0a, 0x18,
	0x8a, 0xe3, 0x0a, 0x14, 0x68, 0x65, 0x6c, 0x70, 0x64, 0x65, 0x73, 0x6b, 0x3a, 0x74, 0x69, 0x63,
	0x6b, 0x65, 0x74, 0x3a, 0x72, 0x65, 0x61, 0x64, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x28, 0x12, 0x26,
	0x2f, 0x68, 0x65, 0x6c, 0x70, 0x64, 0x65, 0x73, 0x6b, 0x2f, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74,
	0x73, 0x2f, 0x7b, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x5f, 0x69, 0x64, 0x7d, 0x2f, 0x63, 0x6f,
	0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12, 0xc9, 0x01, 0x0a, 0x0f, 0x44, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x62, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x25, 0x2e, 0x65, 0x6f, 0x6c,
	0x79, 0x6d, 0x70, 0x2e, 0x68, 0x65, 0x6c, 0x70, 0x64, 0x65, 0x73, 0x6b, 0x2e, 0x44, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x62, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x70, 0x75,
	0x74, 0x1a, 0x26, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x68, 0x65, 0x6c, 0x70, 0x64,
	0x65, 0x73, 0x6b, 0x2e, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x43, 0x6f, 0x6d, 0x6d,
	0x65, 0x6e, 0x74, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x22, 0x67, 0xea, 0xe2, 0x0a, 0x0b, 0xf5,
	0xe2, 0x0a, 0x00, 0x00, 0xa0, 0x40, 0xf8, 0xe2, 0x0a, 0x32, 0x82, 0xe3, 0x0a, 0x19, 0x8a, 0xe3,
	0x0a, 0x15, 0x68, 0x65, 0x6c, 0x70, 0x64, 0x65, 0x73, 0x6b, 0x3a, 0x74, 0x69, 0x63, 0x6b, 0x65,
	0x74, 0x3a, 0x77, 0x72, 0x69, 0x74, 0x65, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x35, 0x12, 0x33, 0x2f,
	0x68, 0x65, 0x6c, 0x70, 0x64, 0x65, 0x73, 0x6b, 0x2f, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x73,
	0x2f, 0x7b, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x5f, 0x69, 0x64, 0x7d, 0x2f, 0x63, 0x6f, 0x6d,
	0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x7b, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69,
	0x64, 0x7d, 0x12, 0xaf, 0x01, 0x0a, 0x0f, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x75, 0x74,
	0x6f, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x25, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e,
	0x68, 0x65, 0x6c, 0x70, 0x64, 0x65, 0x73, 0x6b, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41,
	0x75, 0x74, 0x6f, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x26, 0x2e,
	0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x68, 0x65, 0x6c, 0x70, 0x64, 0x65, 0x73, 0x6b, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x75, 0x74, 0x6f, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x4f,
	0x75, 0x74, 0x70, 0x75, 0x74, 0x22, 0x4d, 0xea, 0xe2, 0x0a, 0x0b, 0xf5, 0xe2, 0x0a, 0x00, 0x00,
	0x00, 0x40, 0xf8, 0xe2, 0x0a, 0x0a, 0x82, 0xe3, 0x0a, 0x1d, 0x8a, 0xe3, 0x0a, 0x19, 0x68, 0x65,
	0x6c, 0x70, 0x64, 0x65, 0x73, 0x6b, 0x3a, 0x61, 0x75, 0x74, 0x6f, 0x2d, 0x72, 0x65, 0x70, 0x6c,
	0x79, 0x3a, 0x77, 0x72, 0x69, 0x74, 0x65, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x17, 0x22, 0x15, 0x2f,
	0x68, 0x65, 0x6c, 0x70, 0x64, 0x65, 0x73, 0x6b, 0x2f, 0x61, 0x75, 0x74, 0x6f, 0x72, 0x65, 0x70,
	0x6c, 0x69, 0x65, 0x73, 0x12, 0xba, 0x01, 0x0a, 0x0f, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41,
	0x75, 0x74, 0x6f, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x25, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d,
	0x70, 0x2e, 0x68, 0x65, 0x6c, 0x70, 0x64, 0x65, 0x73, 0x6b, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x41, 0x75, 0x74, 0x6f, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a,
	0x26, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x68, 0x65, 0x6c, 0x70, 0x64, 0x65, 0x73,
	0x6b, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x75, 0x74, 0x6f, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x22, 0x58, 0xea, 0xe2, 0x0a, 0x0b, 0xf5, 0xe2, 0x0a,
	0x00, 0x00, 0x00, 0x40, 0xf8, 0xe2, 0x0a, 0x0a, 0x82, 0xe3, 0x0a, 0x1d, 0x8a, 0xe3, 0x0a, 0x19,
	0x68, 0x65, 0x6c, 0x70, 0x64, 0x65, 0x73, 0x6b, 0x3a, 0x61, 0x75, 0x74, 0x6f, 0x2d, 0x72, 0x65,
	0x70, 0x6c, 0x79, 0x3a, 0x77, 0x72, 0x69, 0x74, 0x65, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x22, 0x1a,
	0x20, 0x2f, 0x68, 0x65, 0x6c, 0x70, 0x64, 0x65, 0x73, 0x6b, 0x2f, 0x61, 0x75, 0x74, 0x6f, 0x72,
	0x65, 0x70, 0x6c, 0x69, 0x65, 0x73, 0x2f, 0x7b, 0x72, 0x65, 0x70, 0x6c, 0x79, 0x5f, 0x69, 0x64,
	0x7d, 0x12, 0xba, 0x01, 0x0a, 0x0f, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x75, 0x74, 0x6f,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x25, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x68,
	0x65, 0x6c, 0x70, 0x64, 0x65, 0x73, 0x6b, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x75,
	0x74, 0x6f, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x26, 0x2e, 0x65,
	0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x68, 0x65, 0x6c, 0x70, 0x64, 0x65, 0x73, 0x6b, 0x2e, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x75, 0x74, 0x6f, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x4f, 0x75,
	0x74, 0x70, 0x75, 0x74, 0x22, 0x58, 0xea, 0xe2, 0x0a, 0x0b, 0xf5, 0xe2, 0x0a, 0x00, 0x00, 0x00,
	0x40, 0xf8, 0xe2, 0x0a, 0x0a, 0x82, 0xe3, 0x0a, 0x1d, 0x8a, 0xe3, 0x0a, 0x19, 0x68, 0x65, 0x6c,
	0x70, 0x64, 0x65, 0x73, 0x6b, 0x3a, 0x61, 0x75, 0x74, 0x6f, 0x2d, 0x72, 0x65, 0x70, 0x6c, 0x79,
	0x3a, 0x77, 0x72, 0x69, 0x74, 0x65, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x22, 0x2a, 0x20, 0x2f, 0x68,
	0x65, 0x6c, 0x70, 0x64, 0x65, 0x73, 0x6b, 0x2f, 0x61, 0x75, 0x74, 0x6f, 0x72, 0x65, 0x70, 0x6c,
	0x69, 0x65, 0x73, 0x2f, 0x7b, 0x72, 0x65, 0x70, 0x6c, 0x79, 0x5f, 0x69, 0x64, 0x7d, 0x12, 0xae,
	0x01, 0x0a, 0x0f, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x75, 0x74, 0x6f, 0x52, 0x65, 0x70, 0x6c, 0x69,
	0x65, 0x73, 0x12, 0x25, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x68, 0x65, 0x6c, 0x70,
	0x64, 0x65, 0x73, 0x6b, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x75, 0x74, 0x6f, 0x52, 0x65, 0x70,
	0x6c, 0x69, 0x65, 0x73, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x26, 0x2e, 0x65, 0x6f, 0x6c, 0x79,
	0x6d, 0x70, 0x2e, 0x68, 0x65, 0x6c, 0x70, 0x64, 0x65, 0x73, 0x6b, 0x2e, 0x4c, 0x69, 0x73, 0x74,
	0x41, 0x75, 0x74, 0x6f, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x65, 0x73, 0x4f, 0x75, 0x74, 0x70, 0x75,
	0x74, 0x22, 0x4c, 0xea, 0xe2, 0x0a, 0x0b, 0xf5, 0xe2, 0x0a, 0x00, 0x00, 0xa0, 0x40, 0xf8, 0xe2,
	0x0a, 0x14, 0x82, 0xe3, 0x0a, 0x1c, 0x8a, 0xe3, 0x0a, 0x18, 0x68, 0x65, 0x6c, 0x70, 0x64, 0x65,
	0x73, 0x6b, 0x3a, 0x61, 0x75, 0x74, 0x6f, 0x2d, 0x72, 0x65, 0x70, 0x6c, 0x79, 0x3a, 0x72, 0x65,
	0x61, 0x64, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x17, 0x12, 0x15, 0x2f, 0x68, 0x65, 0x6c, 0x70, 0x64,
	0x65, 0x73, 0x6b, 0x2f, 0x61, 0x75, 0x74, 0x6f, 0x72, 0x65, 0x70, 0x6c, 0x69, 0x65, 0x73, 0x12,
	0xc0, 0x01, 0x0a, 0x11, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x41, 0x75, 0x74, 0x6f,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x27, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x68,
	0x65, 0x6c, 0x70, 0x64, 0x65, 0x73, 0x6b, 0x2e, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65,
	0x41, 0x75, 0x74, 0x6f, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x28,
	0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x68, 0x65, 0x6c, 0x70, 0x64, 0x65, 0x73, 0x6b,
	0x2e, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x41, 0x75, 0x74, 0x6f, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x22, 0x58, 0xea, 0xe2, 0x0a, 0x0b, 0xf5, 0xe2,
	0x0a, 0x00, 0x00, 0xa0, 0x40, 0xf8, 0xe2, 0x0a, 0x32, 0x82, 0xe3, 0x0a, 0x1d, 0x8a, 0xe3, 0x0a,
	0x19, 0x68, 0x65, 0x6c, 0x70, 0x64, 0x65, 0x73, 0x6b, 0x3a, 0x61, 0x75, 0x74, 0x6f, 0x2d, 0x72,
	0x65, 0x70, 0x6c, 0x79, 0x3a, 0x77, 0x72, 0x69, 0x74, 0x65, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x22,
	0x12, 0x20, 0x2f, 0x68, 0x65, 0x6c, 0x70, 0x64, 0x65, 0x73, 0x6b, 0x2f, 0x61, 0x75, 0x74, 0x6f,
	0x72, 0x65, 0x70, 0x6c, 0x69, 0x65, 0x73, 0x2f, 0x7b, 0x72, 0x65, 0x70, 0x6c, 0x79, 0x5f, 0x69,
	0x64, 0x7d, 0x12, 0xae, 0x01, 0x0a, 0x10, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x41, 0x74, 0x74,
	0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x26, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70,
	0x2e, 0x68, 0x65, 0x6c, 0x70, 0x64, 0x65, 0x73, 0x6b, 0x2e, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64,
	0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a,
	0x27, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x68, 0x65, 0x6c, 0x70, 0x64, 0x65, 0x73,
	0x6b, 0x2e, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65,
	0x6e, 0x74, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x22, 0x49, 0xea, 0xe2, 0x0a, 0x0b, 0xf5, 0xe2,
	0x0a, 0x00, 0x00, 0xa0, 0x40, 0xf8, 0xe2, 0x0a, 0x32, 0x82, 0xe3, 0x0a, 0x19, 0x8a, 0xe3, 0x0a,
	0x15, 0x68, 0x65, 0x6c, 0x70, 0x64, 0x65, 0x73, 0x6b, 0x3a, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74,
	0x3a, 0x77, 0x72, 0x69, 0x74, 0x65, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x17, 0x22, 0x15, 0x2f, 0x68,
	0x65, 0x6c, 0x70, 0x64, 0x65, 0x73, 0x6b, 0x2f, 0x61, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65,
	0x6e, 0x74, 0x73, 0x42, 0x33, 0x5a, 0x31, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x67, 0x6f, 0x2d, 0x73, 0x64, 0x6b, 0x2f,
	0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x68, 0x65, 0x6c, 0x70, 0x64, 0x65, 0x73, 0x6b, 0x3b,
	0x68, 0x65, 0x6c, 0x70, 0x64, 0x65, 0x73, 0x6b, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var file_eolymp_helpdesk_support_proto_goTypes = []any{
	(*CreateTicketInput)(nil),       // 0: eolymp.helpdesk.CreateTicketInput
	(*UpdateTicketInput)(nil),       // 1: eolymp.helpdesk.UpdateTicketInput
	(*DeleteTicketInput)(nil),       // 2: eolymp.helpdesk.DeleteTicketInput
	(*DescribeTicketInput)(nil),     // 3: eolymp.helpdesk.DescribeTicketInput
	(*ListTicketsInput)(nil),        // 4: eolymp.helpdesk.ListTicketsInput
	(*ApproveTicketInput)(nil),      // 5: eolymp.helpdesk.ApproveTicketInput
	(*RejectTicketInput)(nil),       // 6: eolymp.helpdesk.RejectTicketInput
	(*CloseTicketInput)(nil),        // 7: eolymp.helpdesk.CloseTicketInput
	(*AddCommentInput)(nil),         // 8: eolymp.helpdesk.AddCommentInput
	(*UpdateCommentInput)(nil),      // 9: eolymp.helpdesk.UpdateCommentInput
	(*DeleteCommentInput)(nil),      // 10: eolymp.helpdesk.DeleteCommentInput
	(*ListCommentsInput)(nil),       // 11: eolymp.helpdesk.ListCommentsInput
	(*DescribeCommentInput)(nil),    // 12: eolymp.helpdesk.DescribeCommentInput
	(*CreateAutoReplyInput)(nil),    // 13: eolymp.helpdesk.CreateAutoReplyInput
	(*UpdateAutoReplyInput)(nil),    // 14: eolymp.helpdesk.UpdateAutoReplyInput
	(*DeleteAutoReplyInput)(nil),    // 15: eolymp.helpdesk.DeleteAutoReplyInput
	(*ListAutoRepliesInput)(nil),    // 16: eolymp.helpdesk.ListAutoRepliesInput
	(*DescribeAutoReplyInput)(nil),  // 17: eolymp.helpdesk.DescribeAutoReplyInput
	(*UploadAttachmentInput)(nil),   // 18: eolymp.helpdesk.UploadAttachmentInput
	(*CreateTicketOutput)(nil),      // 19: eolymp.helpdesk.CreateTicketOutput
	(*UpdateTicketOutput)(nil),      // 20: eolymp.helpdesk.UpdateTicketOutput
	(*DeleteTicketOutput)(nil),      // 21: eolymp.helpdesk.DeleteTicketOutput
	(*DescribeTicketOutput)(nil),    // 22: eolymp.helpdesk.DescribeTicketOutput
	(*ListTicketsOutput)(nil),       // 23: eolymp.helpdesk.ListTicketsOutput
	(*ApproveTicketOutput)(nil),     // 24: eolymp.helpdesk.ApproveTicketOutput
	(*RejectTicketOutput)(nil),      // 25: eolymp.helpdesk.RejectTicketOutput
	(*CloseTicketOutput)(nil),       // 26: eolymp.helpdesk.CloseTicketOutput
	(*AddCommentOutput)(nil),        // 27: eolymp.helpdesk.AddCommentOutput
	(*UpdateCommentOutput)(nil),     // 28: eolymp.helpdesk.UpdateCommentOutput
	(*DeleteCommentOutput)(nil),     // 29: eolymp.helpdesk.DeleteCommentOutput
	(*ListCommentsOutput)(nil),      // 30: eolymp.helpdesk.ListCommentsOutput
	(*DescribeCommentOutput)(nil),   // 31: eolymp.helpdesk.DescribeCommentOutput
	(*CreateAutoReplyOutput)(nil),   // 32: eolymp.helpdesk.CreateAutoReplyOutput
	(*UpdateAutoReplyOutput)(nil),   // 33: eolymp.helpdesk.UpdateAutoReplyOutput
	(*DeleteAutoReplyOutput)(nil),   // 34: eolymp.helpdesk.DeleteAutoReplyOutput
	(*ListAutoRepliesOutput)(nil),   // 35: eolymp.helpdesk.ListAutoRepliesOutput
	(*DescribeAutoReplyOutput)(nil), // 36: eolymp.helpdesk.DescribeAutoReplyOutput
	(*UploadAttachmentOutput)(nil),  // 37: eolymp.helpdesk.UploadAttachmentOutput
}
var file_eolymp_helpdesk_support_proto_depIdxs = []int32{
	0,  // 0: eolymp.helpdesk.Support.CreateTicket:input_type -> eolymp.helpdesk.CreateTicketInput
	1,  // 1: eolymp.helpdesk.Support.UpdateTicket:input_type -> eolymp.helpdesk.UpdateTicketInput
	2,  // 2: eolymp.helpdesk.Support.DeleteTicket:input_type -> eolymp.helpdesk.DeleteTicketInput
	3,  // 3: eolymp.helpdesk.Support.DescribeTicket:input_type -> eolymp.helpdesk.DescribeTicketInput
	4,  // 4: eolymp.helpdesk.Support.ListTickets:input_type -> eolymp.helpdesk.ListTicketsInput
	5,  // 5: eolymp.helpdesk.Support.ApproveTicket:input_type -> eolymp.helpdesk.ApproveTicketInput
	6,  // 6: eolymp.helpdesk.Support.RejectTicket:input_type -> eolymp.helpdesk.RejectTicketInput
	7,  // 7: eolymp.helpdesk.Support.CloseTicket:input_type -> eolymp.helpdesk.CloseTicketInput
	8,  // 8: eolymp.helpdesk.Support.AddComment:input_type -> eolymp.helpdesk.AddCommentInput
	9,  // 9: eolymp.helpdesk.Support.UpdateComment:input_type -> eolymp.helpdesk.UpdateCommentInput
	10, // 10: eolymp.helpdesk.Support.DeleteComment:input_type -> eolymp.helpdesk.DeleteCommentInput
	11, // 11: eolymp.helpdesk.Support.ListComments:input_type -> eolymp.helpdesk.ListCommentsInput
	12, // 12: eolymp.helpdesk.Support.DescribeComment:input_type -> eolymp.helpdesk.DescribeCommentInput
	13, // 13: eolymp.helpdesk.Support.CreateAutoReply:input_type -> eolymp.helpdesk.CreateAutoReplyInput
	14, // 14: eolymp.helpdesk.Support.UpdateAutoReply:input_type -> eolymp.helpdesk.UpdateAutoReplyInput
	15, // 15: eolymp.helpdesk.Support.DeleteAutoReply:input_type -> eolymp.helpdesk.DeleteAutoReplyInput
	16, // 16: eolymp.helpdesk.Support.ListAutoReplies:input_type -> eolymp.helpdesk.ListAutoRepliesInput
	17, // 17: eolymp.helpdesk.Support.DescribeAutoReply:input_type -> eolymp.helpdesk.DescribeAutoReplyInput
	18, // 18: eolymp.helpdesk.Support.UploadAttachment:input_type -> eolymp.helpdesk.UploadAttachmentInput
	19, // 19: eolymp.helpdesk.Support.CreateTicket:output_type -> eolymp.helpdesk.CreateTicketOutput
	20, // 20: eolymp.helpdesk.Support.UpdateTicket:output_type -> eolymp.helpdesk.UpdateTicketOutput
	21, // 21: eolymp.helpdesk.Support.DeleteTicket:output_type -> eolymp.helpdesk.DeleteTicketOutput
	22, // 22: eolymp.helpdesk.Support.DescribeTicket:output_type -> eolymp.helpdesk.DescribeTicketOutput
	23, // 23: eolymp.helpdesk.Support.ListTickets:output_type -> eolymp.helpdesk.ListTicketsOutput
	24, // 24: eolymp.helpdesk.Support.ApproveTicket:output_type -> eolymp.helpdesk.ApproveTicketOutput
	25, // 25: eolymp.helpdesk.Support.RejectTicket:output_type -> eolymp.helpdesk.RejectTicketOutput
	26, // 26: eolymp.helpdesk.Support.CloseTicket:output_type -> eolymp.helpdesk.CloseTicketOutput
	27, // 27: eolymp.helpdesk.Support.AddComment:output_type -> eolymp.helpdesk.AddCommentOutput
	28, // 28: eolymp.helpdesk.Support.UpdateComment:output_type -> eolymp.helpdesk.UpdateCommentOutput
	29, // 29: eolymp.helpdesk.Support.DeleteComment:output_type -> eolymp.helpdesk.DeleteCommentOutput
	30, // 30: eolymp.helpdesk.Support.ListComments:output_type -> eolymp.helpdesk.ListCommentsOutput
	31, // 31: eolymp.helpdesk.Support.DescribeComment:output_type -> eolymp.helpdesk.DescribeCommentOutput
	32, // 32: eolymp.helpdesk.Support.CreateAutoReply:output_type -> eolymp.helpdesk.CreateAutoReplyOutput
	33, // 33: eolymp.helpdesk.Support.UpdateAutoReply:output_type -> eolymp.helpdesk.UpdateAutoReplyOutput
	34, // 34: eolymp.helpdesk.Support.DeleteAutoReply:output_type -> eolymp.helpdesk.DeleteAutoReplyOutput
	35, // 35: eolymp.helpdesk.Support.ListAutoReplies:output_type -> eolymp.helpdesk.ListAutoRepliesOutput
	36, // 36: eolymp.helpdesk.Support.DescribeAutoReply:output_type -> eolymp.helpdesk.DescribeAutoReplyOutput
	37, // 37: eolymp.helpdesk.Support.UploadAttachment:output_type -> eolymp.helpdesk.UploadAttachmentOutput
	19, // [19:38] is the sub-list for method output_type
	0,  // [0:19] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_eolymp_helpdesk_support_proto_init() }
func file_eolymp_helpdesk_support_proto_init() {
	if File_eolymp_helpdesk_support_proto != nil {
		return
	}
	file_eolymp_helpdesk_auto_reply_service_proto_init()
	file_eolymp_helpdesk_ticket_service_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_eolymp_helpdesk_support_proto_rawDesc), len(file_eolymp_helpdesk_support_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_eolymp_helpdesk_support_proto_goTypes,
		DependencyIndexes: file_eolymp_helpdesk_support_proto_depIdxs,
	}.Build()
	File_eolymp_helpdesk_support_proto = out.File
	file_eolymp_helpdesk_support_proto_goTypes = nil
	file_eolymp_helpdesk_support_proto_depIdxs = nil
}
