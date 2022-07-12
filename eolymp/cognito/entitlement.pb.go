// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.12.2
// source: eolymp/cognito/entitlement.proto

package cognito

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

type Entitlement int32

const (
	Entitlement_VIEW_PUBLIC_PROFILE  Entitlement = 0 // View public information about active user.
	Entitlement_VIEW_BLOCKED_PROFILE Entitlement = 1 // View public information about blocked users.
	Entitlement_VIEW_PRIVATE_DATA    Entitlement = 2 // View private information about user: email, birthday etc.
	Entitlement_MANAGE_ROLES         Entitlement = 3 // Update user's roles.
	Entitlement_BLOCK_USERS          Entitlement = 4 // Activate and de-activate user profiles.
)

// Enum value maps for Entitlement.
var (
	Entitlement_name = map[int32]string{
		0: "VIEW_PUBLIC_PROFILE",
		1: "VIEW_BLOCKED_PROFILE",
		2: "VIEW_PRIVATE_DATA",
		3: "MANAGE_ROLES",
		4: "BLOCK_USERS",
	}
	Entitlement_value = map[string]int32{
		"VIEW_PUBLIC_PROFILE":  0,
		"VIEW_BLOCKED_PROFILE": 1,
		"VIEW_PRIVATE_DATA":    2,
		"MANAGE_ROLES":         3,
		"BLOCK_USERS":          4,
	}
)

func (x Entitlement) Enum() *Entitlement {
	p := new(Entitlement)
	*p = x
	return p
}

func (x Entitlement) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Entitlement) Descriptor() protoreflect.EnumDescriptor {
	return file_eolymp_cognito_entitlement_proto_enumTypes[0].Descriptor()
}

func (Entitlement) Type() protoreflect.EnumType {
	return &file_eolymp_cognito_entitlement_proto_enumTypes[0]
}

func (x Entitlement) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Entitlement.Descriptor instead.
func (Entitlement) EnumDescriptor() ([]byte, []int) {
	return file_eolymp_cognito_entitlement_proto_rawDescGZIP(), []int{0}
}

var File_eolymp_cognito_entitlement_proto protoreflect.FileDescriptor

var file_eolymp_cognito_entitlement_proto_rawDesc = []byte{
	0x0a, 0x20, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x63, 0x6f, 0x67, 0x6e, 0x69, 0x74, 0x6f,
	0x2f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x63, 0x6f, 0x67, 0x6e, 0x69,
	0x74, 0x6f, 0x2a, 0x7a, 0x0a, 0x0b, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x6d, 0x65, 0x6e,
	0x74, 0x12, 0x17, 0x0a, 0x13, 0x56, 0x49, 0x45, 0x57, 0x5f, 0x50, 0x55, 0x42, 0x4c, 0x49, 0x43,
	0x5f, 0x50, 0x52, 0x4f, 0x46, 0x49, 0x4c, 0x45, 0x10, 0x00, 0x12, 0x18, 0x0a, 0x14, 0x56, 0x49,
	0x45, 0x57, 0x5f, 0x42, 0x4c, 0x4f, 0x43, 0x4b, 0x45, 0x44, 0x5f, 0x50, 0x52, 0x4f, 0x46, 0x49,
	0x4c, 0x45, 0x10, 0x01, 0x12, 0x15, 0x0a, 0x11, 0x56, 0x49, 0x45, 0x57, 0x5f, 0x50, 0x52, 0x49,
	0x56, 0x41, 0x54, 0x45, 0x5f, 0x44, 0x41, 0x54, 0x41, 0x10, 0x02, 0x12, 0x10, 0x0a, 0x0c, 0x4d,
	0x41, 0x4e, 0x41, 0x47, 0x45, 0x5f, 0x52, 0x4f, 0x4c, 0x45, 0x53, 0x10, 0x03, 0x12, 0x0f, 0x0a,
	0x0b, 0x42, 0x4c, 0x4f, 0x43, 0x4b, 0x5f, 0x55, 0x53, 0x45, 0x52, 0x53, 0x10, 0x04, 0x42, 0x31,
	0x5a, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6f, 0x6c,
	0x79, 0x6d, 0x70, 0x2f, 0x67, 0x6f, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x65, 0x6f, 0x6c, 0x79, 0x6d,
	0x70, 0x2f, 0x63, 0x6f, 0x67, 0x6e, 0x69, 0x74, 0x6f, 0x3b, 0x63, 0x6f, 0x67, 0x6e, 0x69, 0x74,
	0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_eolymp_cognito_entitlement_proto_rawDescOnce sync.Once
	file_eolymp_cognito_entitlement_proto_rawDescData = file_eolymp_cognito_entitlement_proto_rawDesc
)

func file_eolymp_cognito_entitlement_proto_rawDescGZIP() []byte {
	file_eolymp_cognito_entitlement_proto_rawDescOnce.Do(func() {
		file_eolymp_cognito_entitlement_proto_rawDescData = protoimpl.X.CompressGZIP(file_eolymp_cognito_entitlement_proto_rawDescData)
	})
	return file_eolymp_cognito_entitlement_proto_rawDescData
}

var file_eolymp_cognito_entitlement_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_eolymp_cognito_entitlement_proto_goTypes = []interface{}{
	(Entitlement)(0), // 0: eolymp.cognito.Entitlement
}
var file_eolymp_cognito_entitlement_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_eolymp_cognito_entitlement_proto_init() }
func file_eolymp_cognito_entitlement_proto_init() {
	if File_eolymp_cognito_entitlement_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_eolymp_cognito_entitlement_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_eolymp_cognito_entitlement_proto_goTypes,
		DependencyIndexes: file_eolymp_cognito_entitlement_proto_depIdxs,
		EnumInfos:         file_eolymp_cognito_entitlement_proto_enumTypes,
	}.Build()
	File_eolymp_cognito_entitlement_proto = out.File
	file_eolymp_cognito_entitlement_proto_rawDesc = nil
	file_eolymp_cognito_entitlement_proto_goTypes = nil
	file_eolymp_cognito_entitlement_proto_depIdxs = nil
}
