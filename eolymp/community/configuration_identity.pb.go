// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v3.21.12
// source: eolymp/community/configuration_identity.proto

package community

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

type IdentityConfig_DisplayNameType int32

const (
	IdentityConfig_UNKNOWN_DISPLAY_NAME IdentityConfig_DisplayNameType = 0
	IdentityConfig_NICKNAME             IdentityConfig_DisplayNameType = 1
	IdentityConfig_NAME                 IdentityConfig_DisplayNameType = 2
	IdentityConfig_ATTRIBUTE            IdentityConfig_DisplayNameType = 3
)

// Enum value maps for IdentityConfig_DisplayNameType.
var (
	IdentityConfig_DisplayNameType_name = map[int32]string{
		0: "UNKNOWN_DISPLAY_NAME",
		1: "NICKNAME",
		2: "NAME",
		3: "ATTRIBUTE",
	}
	IdentityConfig_DisplayNameType_value = map[string]int32{
		"UNKNOWN_DISPLAY_NAME": 0,
		"NICKNAME":             1,
		"NAME":                 2,
		"ATTRIBUTE":            3,
	}
)

func (x IdentityConfig_DisplayNameType) Enum() *IdentityConfig_DisplayNameType {
	p := new(IdentityConfig_DisplayNameType)
	*p = x
	return p
}

func (x IdentityConfig_DisplayNameType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (IdentityConfig_DisplayNameType) Descriptor() protoreflect.EnumDescriptor {
	return file_eolymp_community_configuration_identity_proto_enumTypes[0].Descriptor()
}

func (IdentityConfig_DisplayNameType) Type() protoreflect.EnumType {
	return &file_eolymp_community_configuration_identity_proto_enumTypes[0]
}

func (x IdentityConfig_DisplayNameType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use IdentityConfig_DisplayNameType.Descriptor instead.
func (IdentityConfig_DisplayNameType) EnumDescriptor() ([]byte, []int) {
	return file_eolymp_community_configuration_identity_proto_rawDescGZIP(), []int{0, 0}
}

type IdentityConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Provider:
	//
	//	*IdentityConfig_Local
	//	*IdentityConfig_Basecamp
	//	*IdentityConfig_Oidc
	Provider             isIdentityConfig_Provider      `protobuf_oneof:"provider"`
	DisplayNameType      IdentityConfig_DisplayNameType `protobuf:"varint,100,opt,name=display_name_type,json=displayNameType,proto3,enum=eolymp.community.IdentityConfig_DisplayNameType" json:"display_name_type,omitempty"`
	DisplayNameAttribute string                         `protobuf:"bytes,101,opt,name=display_name_attribute,json=displayNameAttribute,proto3" json:"display_name_attribute,omitempty"`
	AllowSignUp          bool                           `protobuf:"varint,102,opt,name=allow_sign_up,json=allowSignUp,proto3" json:"allow_sign_up,omitempty"`                            // users can join on their own
	RequireEmailVerified bool                           `protobuf:"varint,103,opt,name=require_email_verified,json=requireEmailVerified,proto3" json:"require_email_verified,omitempty"` // users must verify email to use the site, setting this flag to false will override "email_verified" to true for all members
}

func (x *IdentityConfig) Reset() {
	*x = IdentityConfig{}
	mi := &file_eolymp_community_configuration_identity_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *IdentityConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IdentityConfig) ProtoMessage() {}

func (x *IdentityConfig) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_community_configuration_identity_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IdentityConfig.ProtoReflect.Descriptor instead.
func (*IdentityConfig) Descriptor() ([]byte, []int) {
	return file_eolymp_community_configuration_identity_proto_rawDescGZIP(), []int{0}
}

func (m *IdentityConfig) GetProvider() isIdentityConfig_Provider {
	if m != nil {
		return m.Provider
	}
	return nil
}

func (x *IdentityConfig) GetLocal() *IdentityProvider_Local {
	if x, ok := x.GetProvider().(*IdentityConfig_Local); ok {
		return x.Local
	}
	return nil
}

func (x *IdentityConfig) GetBasecamp() *IdentityProvider_Basecamp {
	if x, ok := x.GetProvider().(*IdentityConfig_Basecamp); ok {
		return x.Basecamp
	}
	return nil
}

func (x *IdentityConfig) GetOidc() *IdentityProvider_OIDC {
	if x, ok := x.GetProvider().(*IdentityConfig_Oidc); ok {
		return x.Oidc
	}
	return nil
}

func (x *IdentityConfig) GetDisplayNameType() IdentityConfig_DisplayNameType {
	if x != nil {
		return x.DisplayNameType
	}
	return IdentityConfig_UNKNOWN_DISPLAY_NAME
}

func (x *IdentityConfig) GetDisplayNameAttribute() string {
	if x != nil {
		return x.DisplayNameAttribute
	}
	return ""
}

func (x *IdentityConfig) GetAllowSignUp() bool {
	if x != nil {
		return x.AllowSignUp
	}
	return false
}

func (x *IdentityConfig) GetRequireEmailVerified() bool {
	if x != nil {
		return x.RequireEmailVerified
	}
	return false
}

type isIdentityConfig_Provider interface {
	isIdentityConfig_Provider()
}

type IdentityConfig_Local struct {
	Local *IdentityProvider_Local `protobuf:"bytes,1,opt,name=local,proto3,oneof"`
}

type IdentityConfig_Basecamp struct {
	Basecamp *IdentityProvider_Basecamp `protobuf:"bytes,2,opt,name=basecamp,proto3,oneof"`
}

type IdentityConfig_Oidc struct {
	Oidc *IdentityProvider_OIDC `protobuf:"bytes,3,opt,name=oidc,proto3,oneof"`
}

func (*IdentityConfig_Local) isIdentityConfig_Provider() {}

func (*IdentityConfig_Basecamp) isIdentityConfig_Provider() {}

func (*IdentityConfig_Oidc) isIdentityConfig_Provider() {}

var File_eolymp_community_configuration_identity_proto protoreflect.FileDescriptor

var file_eolymp_community_configuration_identity_proto_rawDesc = []byte{
	0x0a, 0x2d, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69,
	0x74, 0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x10, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74,
	0x79, 0x1a, 0x28, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e,
	0x69, 0x74, 0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x69, 0x64, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xaa, 0x04, 0x0a, 0x0e,
	0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x40,
	0x0a, 0x05, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e,
	0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79,
	0x2e, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65,
	0x72, 0x2e, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x48, 0x00, 0x52, 0x05, 0x6c, 0x6f, 0x63, 0x61, 0x6c,
	0x12, 0x49, 0x0a, 0x08, 0x62, 0x61, 0x73, 0x65, 0x63, 0x61, 0x6d, 0x70, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x75, 0x6e, 0x69, 0x74, 0x79, 0x2e, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x50, 0x72,
	0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x63, 0x61, 0x6d, 0x70, 0x48,
	0x00, 0x52, 0x08, 0x62, 0x61, 0x73, 0x65, 0x63, 0x61, 0x6d, 0x70, 0x12, 0x3d, 0x0a, 0x04, 0x6f,
	0x69, 0x64, 0x63, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x65, 0x6f, 0x6c, 0x79,
	0x6d, 0x70, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x2e, 0x49, 0x64, 0x65,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x4f, 0x49,
	0x44, 0x43, 0x48, 0x00, 0x52, 0x04, 0x6f, 0x69, 0x64, 0x63, 0x12, 0x5c, 0x0a, 0x11, 0x64, 0x69,
	0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x64, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x30, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x2e, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x44, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x4e,
	0x61, 0x6d, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0f, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79,
	0x4e, 0x61, 0x6d, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x34, 0x0a, 0x16, 0x64, 0x69, 0x73, 0x70,
	0x6c, 0x61, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x5f, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75,
	0x74, 0x65, 0x18, 0x65, 0x20, 0x01, 0x28, 0x09, 0x52, 0x14, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61,
	0x79, 0x4e, 0x61, 0x6d, 0x65, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x12, 0x22,
	0x0a, 0x0d, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x5f, 0x73, 0x69, 0x67, 0x6e, 0x5f, 0x75, 0x70, 0x18,
	0x66, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x53, 0x69, 0x67, 0x6e,
	0x55, 0x70, 0x12, 0x34, 0x0a, 0x16, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x5f, 0x65, 0x6d,
	0x61, 0x69, 0x6c, 0x5f, 0x76, 0x65, 0x72, 0x69, 0x66, 0x69, 0x65, 0x64, 0x18, 0x67, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x14, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x45, 0x6d, 0x61, 0x69, 0x6c,
	0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x65, 0x64, 0x22, 0x52, 0x0a, 0x0f, 0x44, 0x69, 0x73, 0x70,
	0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x14, 0x55,
	0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x5f, 0x44, 0x49, 0x53, 0x50, 0x4c, 0x41, 0x59, 0x5f, 0x4e,
	0x41, 0x4d, 0x45, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x4e, 0x49, 0x43, 0x4b, 0x4e, 0x41, 0x4d,
	0x45, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x41, 0x4d, 0x45, 0x10, 0x02, 0x12, 0x0d, 0x0a,
	0x09, 0x41, 0x54, 0x54, 0x52, 0x49, 0x42, 0x55, 0x54, 0x45, 0x10, 0x03, 0x42, 0x0a, 0x0a, 0x08,
	0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x42, 0x35, 0x5a, 0x33, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x67, 0x6f,
	0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x63, 0x6f, 0x6d, 0x6d,
	0x75, 0x6e, 0x69, 0x74, 0x79, 0x3b, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_eolymp_community_configuration_identity_proto_rawDescOnce sync.Once
	file_eolymp_community_configuration_identity_proto_rawDescData = file_eolymp_community_configuration_identity_proto_rawDesc
)

func file_eolymp_community_configuration_identity_proto_rawDescGZIP() []byte {
	file_eolymp_community_configuration_identity_proto_rawDescOnce.Do(func() {
		file_eolymp_community_configuration_identity_proto_rawDescData = protoimpl.X.CompressGZIP(file_eolymp_community_configuration_identity_proto_rawDescData)
	})
	return file_eolymp_community_configuration_identity_proto_rawDescData
}

var file_eolymp_community_configuration_identity_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_eolymp_community_configuration_identity_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_eolymp_community_configuration_identity_proto_goTypes = []any{
	(IdentityConfig_DisplayNameType)(0), // 0: eolymp.community.IdentityConfig.DisplayNameType
	(*IdentityConfig)(nil),              // 1: eolymp.community.IdentityConfig
	(*IdentityProvider_Local)(nil),      // 2: eolymp.community.IdentityProvider.Local
	(*IdentityProvider_Basecamp)(nil),   // 3: eolymp.community.IdentityProvider.Basecamp
	(*IdentityProvider_OIDC)(nil),       // 4: eolymp.community.IdentityProvider.OIDC
}
var file_eolymp_community_configuration_identity_proto_depIdxs = []int32{
	2, // 0: eolymp.community.IdentityConfig.local:type_name -> eolymp.community.IdentityProvider.Local
	3, // 1: eolymp.community.IdentityConfig.basecamp:type_name -> eolymp.community.IdentityProvider.Basecamp
	4, // 2: eolymp.community.IdentityConfig.oidc:type_name -> eolymp.community.IdentityProvider.OIDC
	0, // 3: eolymp.community.IdentityConfig.display_name_type:type_name -> eolymp.community.IdentityConfig.DisplayNameType
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_eolymp_community_configuration_identity_proto_init() }
func file_eolymp_community_configuration_identity_proto_init() {
	if File_eolymp_community_configuration_identity_proto != nil {
		return
	}
	file_eolymp_community_configuration_idp_proto_init()
	file_eolymp_community_configuration_identity_proto_msgTypes[0].OneofWrappers = []any{
		(*IdentityConfig_Local)(nil),
		(*IdentityConfig_Basecamp)(nil),
		(*IdentityConfig_Oidc)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_eolymp_community_configuration_identity_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_eolymp_community_configuration_identity_proto_goTypes,
		DependencyIndexes: file_eolymp_community_configuration_identity_proto_depIdxs,
		EnumInfos:         file_eolymp_community_configuration_identity_proto_enumTypes,
		MessageInfos:      file_eolymp_community_configuration_identity_proto_msgTypes,
	}.Build()
	File_eolymp_community_configuration_identity_proto = out.File
	file_eolymp_community_configuration_identity_proto_rawDesc = nil
	file_eolymp_community_configuration_identity_proto_goTypes = nil
	file_eolymp_community_configuration_identity_proto_depIdxs = nil
}
