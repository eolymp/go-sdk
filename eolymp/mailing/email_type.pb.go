// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.21.12
// source: eolymp/mailing/email_type.proto

package mailing

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

type EmailType int32

const (
	EmailType_UNKNOWN_TYPE EmailType = 0
	EmailType_ACCOUNT      EmailType = 1
	EmailType_NEWSLETTER   EmailType = 2
	EmailType_UPDATES      EmailType = 3
	EmailType_COMMERCIAL   EmailType = 4
	EmailType_OTHER        EmailType = 5
)

// Enum value maps for EmailType.
var (
	EmailType_name = map[int32]string{
		0: "UNKNOWN_TYPE",
		1: "ACCOUNT",
		2: "NEWSLETTER",
		3: "UPDATES",
		4: "COMMERCIAL",
		5: "OTHER",
	}
	EmailType_value = map[string]int32{
		"UNKNOWN_TYPE": 0,
		"ACCOUNT":      1,
		"NEWSLETTER":   2,
		"UPDATES":      3,
		"COMMERCIAL":   4,
		"OTHER":        5,
	}
)

func (x EmailType) Enum() *EmailType {
	p := new(EmailType)
	*p = x
	return p
}

func (x EmailType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EmailType) Descriptor() protoreflect.EnumDescriptor {
	return file_eolymp_mailing_email_type_proto_enumTypes[0].Descriptor()
}

func (EmailType) Type() protoreflect.EnumType {
	return &file_eolymp_mailing_email_type_proto_enumTypes[0]
}

func (x EmailType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use EmailType.Descriptor instead.
func (EmailType) EnumDescriptor() ([]byte, []int) {
	return file_eolymp_mailing_email_type_proto_rawDescGZIP(), []int{0}
}

var File_eolymp_mailing_email_type_proto protoreflect.FileDescriptor

var file_eolymp_mailing_email_type_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x6d, 0x61, 0x69, 0x6c, 0x69, 0x6e, 0x67,
	0x2f, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x6d, 0x61, 0x69, 0x6c, 0x69, 0x6e,
	0x67, 0x2a, 0x62, 0x0a, 0x09, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x12, 0x10,
	0x0a, 0x0c, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x10, 0x00,
	0x12, 0x0b, 0x0a, 0x07, 0x41, 0x43, 0x43, 0x4f, 0x55, 0x4e, 0x54, 0x10, 0x01, 0x12, 0x0e, 0x0a,
	0x0a, 0x4e, 0x45, 0x57, 0x53, 0x4c, 0x45, 0x54, 0x54, 0x45, 0x52, 0x10, 0x02, 0x12, 0x0b, 0x0a,
	0x07, 0x55, 0x50, 0x44, 0x41, 0x54, 0x45, 0x53, 0x10, 0x03, 0x12, 0x0e, 0x0a, 0x0a, 0x43, 0x4f,
	0x4d, 0x4d, 0x45, 0x52, 0x43, 0x49, 0x41, 0x4c, 0x10, 0x04, 0x12, 0x09, 0x0a, 0x05, 0x4f, 0x54,
	0x48, 0x45, 0x52, 0x10, 0x05, 0x42, 0x31, 0x5a, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x67, 0x6f, 0x2d, 0x73, 0x64,
	0x6b, 0x2f, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x6d, 0x61, 0x69, 0x6c, 0x69, 0x6e, 0x67,
	0x3b, 0x6d, 0x61, 0x69, 0x6c, 0x69, 0x6e, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_eolymp_mailing_email_type_proto_rawDescOnce sync.Once
	file_eolymp_mailing_email_type_proto_rawDescData = file_eolymp_mailing_email_type_proto_rawDesc
)

func file_eolymp_mailing_email_type_proto_rawDescGZIP() []byte {
	file_eolymp_mailing_email_type_proto_rawDescOnce.Do(func() {
		file_eolymp_mailing_email_type_proto_rawDescData = protoimpl.X.CompressGZIP(file_eolymp_mailing_email_type_proto_rawDescData)
	})
	return file_eolymp_mailing_email_type_proto_rawDescData
}

var file_eolymp_mailing_email_type_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_eolymp_mailing_email_type_proto_goTypes = []any{
	(EmailType)(0), // 0: eolymp.mailing.EmailType
}
var file_eolymp_mailing_email_type_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_eolymp_mailing_email_type_proto_init() }
func file_eolymp_mailing_email_type_proto_init() {
	if File_eolymp_mailing_email_type_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_eolymp_mailing_email_type_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_eolymp_mailing_email_type_proto_goTypes,
		DependencyIndexes: file_eolymp_mailing_email_type_proto_depIdxs,
		EnumInfos:         file_eolymp_mailing_email_type_proto_enumTypes,
	}.Build()
	File_eolymp_mailing_email_type_proto = out.File
	file_eolymp_mailing_email_type_proto_rawDesc = nil
	file_eolymp_mailing_email_type_proto_goTypes = nil
	file_eolymp_mailing_email_type_proto_depIdxs = nil
}
