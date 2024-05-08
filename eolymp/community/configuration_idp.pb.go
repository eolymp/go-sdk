// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v3.21.12
// source: eolymp/community/configuration_idp.proto

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

type IdentityProvider struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *IdentityProvider) Reset() {
	*x = IdentityProvider{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_community_configuration_idp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IdentityProvider) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IdentityProvider) ProtoMessage() {}

func (x *IdentityProvider) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_community_configuration_idp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IdentityProvider.ProtoReflect.Descriptor instead.
func (*IdentityProvider) Descriptor() ([]byte, []int) {
	return file_eolymp_community_configuration_idp_proto_rawDescGZIP(), []int{0}
}

type IdentityProvider_Basecamp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *IdentityProvider_Basecamp) Reset() {
	*x = IdentityProvider_Basecamp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_community_configuration_idp_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IdentityProvider_Basecamp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IdentityProvider_Basecamp) ProtoMessage() {}

func (x *IdentityProvider_Basecamp) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_community_configuration_idp_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IdentityProvider_Basecamp.ProtoReflect.Descriptor instead.
func (*IdentityProvider_Basecamp) Descriptor() ([]byte, []int) {
	return file_eolymp_community_configuration_idp_proto_rawDescGZIP(), []int{0, 0}
}

type IdentityProvider_Local struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AllowModifyNickname bool `protobuf:"varint,1,opt,name=allow_modify_nickname,json=allowModifyNickname,proto3" json:"allow_modify_nickname,omitempty"`
	AllowModifyBasics   bool `protobuf:"varint,2,opt,name=allow_modify_basics,json=allowModifyBasics,proto3" json:"allow_modify_basics,omitempty"`
	AllowModifyEmail    bool `protobuf:"varint,3,opt,name=allow_modify_email,json=allowModifyEmail,proto3" json:"allow_modify_email,omitempty"`
	AllowModifyPassword bool `protobuf:"varint,4,opt,name=allow_modify_password,json=allowModifyPassword,proto3" json:"allow_modify_password,omitempty"`
}

func (x *IdentityProvider_Local) Reset() {
	*x = IdentityProvider_Local{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_community_configuration_idp_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IdentityProvider_Local) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IdentityProvider_Local) ProtoMessage() {}

func (x *IdentityProvider_Local) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_community_configuration_idp_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IdentityProvider_Local.ProtoReflect.Descriptor instead.
func (*IdentityProvider_Local) Descriptor() ([]byte, []int) {
	return file_eolymp_community_configuration_idp_proto_rawDescGZIP(), []int{0, 1}
}

func (x *IdentityProvider_Local) GetAllowModifyNickname() bool {
	if x != nil {
		return x.AllowModifyNickname
	}
	return false
}

func (x *IdentityProvider_Local) GetAllowModifyBasics() bool {
	if x != nil {
		return x.AllowModifyBasics
	}
	return false
}

func (x *IdentityProvider_Local) GetAllowModifyEmail() bool {
	if x != nil {
		return x.AllowModifyEmail
	}
	return false
}

func (x *IdentityProvider_Local) GetAllowModifyPassword() bool {
	if x != nil {
		return x.AllowModifyPassword
	}
	return false
}

type IdentityProvider_OIDC struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClientId          string `protobuf:"bytes,10,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	ClientSecret      string `protobuf:"bytes,11,opt,name=client_secret,json=clientSecret,proto3" json:"client_secret,omitempty"`
	Issuer            string `protobuf:"bytes,20,opt,name=issuer,proto3" json:"issuer,omitempty"`
	AuthorizeEndpoint string `protobuf:"bytes,21,opt,name=authorize_endpoint,json=authorizeEndpoint,proto3" json:"authorize_endpoint,omitempty"`
	TokenEndpoint     string `protobuf:"bytes,22,opt,name=token_endpoint,json=tokenEndpoint,proto3" json:"token_endpoint,omitempty"`
	KeysEndpoint      string `protobuf:"bytes,23,opt,name=keys_endpoint,json=keysEndpoint,proto3" json:"keys_endpoint,omitempty"`
	UserinfoEndpoint  string `protobuf:"bytes,24,opt,name=userinfo_endpoint,json=userinfoEndpoint,proto3" json:"userinfo_endpoint,omitempty"`
	RedirectUri       string `protobuf:"bytes,25,opt,name=redirect_uri,json=redirectUri,proto3" json:"redirect_uri,omitempty"`
}

func (x *IdentityProvider_OIDC) Reset() {
	*x = IdentityProvider_OIDC{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_community_configuration_idp_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IdentityProvider_OIDC) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IdentityProvider_OIDC) ProtoMessage() {}

func (x *IdentityProvider_OIDC) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_community_configuration_idp_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IdentityProvider_OIDC.ProtoReflect.Descriptor instead.
func (*IdentityProvider_OIDC) Descriptor() ([]byte, []int) {
	return file_eolymp_community_configuration_idp_proto_rawDescGZIP(), []int{0, 2}
}

func (x *IdentityProvider_OIDC) GetClientId() string {
	if x != nil {
		return x.ClientId
	}
	return ""
}

func (x *IdentityProvider_OIDC) GetClientSecret() string {
	if x != nil {
		return x.ClientSecret
	}
	return ""
}

func (x *IdentityProvider_OIDC) GetIssuer() string {
	if x != nil {
		return x.Issuer
	}
	return ""
}

func (x *IdentityProvider_OIDC) GetAuthorizeEndpoint() string {
	if x != nil {
		return x.AuthorizeEndpoint
	}
	return ""
}

func (x *IdentityProvider_OIDC) GetTokenEndpoint() string {
	if x != nil {
		return x.TokenEndpoint
	}
	return ""
}

func (x *IdentityProvider_OIDC) GetKeysEndpoint() string {
	if x != nil {
		return x.KeysEndpoint
	}
	return ""
}

func (x *IdentityProvider_OIDC) GetUserinfoEndpoint() string {
	if x != nil {
		return x.UserinfoEndpoint
	}
	return ""
}

func (x *IdentityProvider_OIDC) GetRedirectUri() string {
	if x != nil {
		return x.RedirectUri
	}
	return ""
}

var File_eolymp_community_configuration_idp_proto protoreflect.FileDescriptor

var file_eolymp_community_configuration_idp_proto_rawDesc = []byte{
	0x0a, 0x28, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69,
	0x74, 0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x69, 0x64, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x65, 0x6f, 0x6c, 0x79,
	0x6d, 0x70, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x22, 0x9c, 0x04, 0x0a,
	0x10, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65,
	0x72, 0x1a, 0x0a, 0x0a, 0x08, 0x42, 0x61, 0x73, 0x65, 0x63, 0x61, 0x6d, 0x70, 0x1a, 0xcd, 0x01,
	0x0a, 0x05, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x12, 0x32, 0x0a, 0x15, 0x61, 0x6c, 0x6c, 0x6f, 0x77,
	0x5f, 0x6d, 0x6f, 0x64, 0x69, 0x66, 0x79, 0x5f, 0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x13, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x4d, 0x6f, 0x64,
	0x69, 0x66, 0x79, 0x4e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x2e, 0x0a, 0x13, 0x61,
	0x6c, 0x6c, 0x6f, 0x77, 0x5f, 0x6d, 0x6f, 0x64, 0x69, 0x66, 0x79, 0x5f, 0x62, 0x61, 0x73, 0x69,
	0x63, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x11, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x4d,
	0x6f, 0x64, 0x69, 0x66, 0x79, 0x42, 0x61, 0x73, 0x69, 0x63, 0x73, 0x12, 0x2c, 0x0a, 0x12, 0x61,
	0x6c, 0x6c, 0x6f, 0x77, 0x5f, 0x6d, 0x6f, 0x64, 0x69, 0x66, 0x79, 0x5f, 0x65, 0x6d, 0x61, 0x69,
	0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x10, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x4d, 0x6f,
	0x64, 0x69, 0x66, 0x79, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x32, 0x0a, 0x15, 0x61, 0x6c, 0x6c,
	0x6f, 0x77, 0x5f, 0x6d, 0x6f, 0x64, 0x69, 0x66, 0x79, 0x5f, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f,
	0x72, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x13, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x4d,
	0x6f, 0x64, 0x69, 0x66, 0x79, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x1a, 0xab, 0x02,
	0x0a, 0x04, 0x4f, 0x49, 0x44, 0x43, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x5f, 0x69, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x49, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x65,
	0x63, 0x72, 0x65, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x69, 0x73, 0x73, 0x75,
	0x65, 0x72, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x69, 0x73, 0x73, 0x75, 0x65, 0x72,
	0x12, 0x2d, 0x0a, 0x12, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x5f, 0x65, 0x6e,
	0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x15, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x61, 0x75,
	0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x12,
	0x25, 0x0a, 0x0e, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e,
	0x74, 0x18, 0x16, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x45, 0x6e,
	0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x6b, 0x65, 0x79, 0x73, 0x5f, 0x65,
	0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x17, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x6b,
	0x65, 0x79, 0x73, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x2b, 0x0a, 0x11, 0x75,
	0x73, 0x65, 0x72, 0x69, 0x6e, 0x66, 0x6f, 0x5f, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74,
	0x18, 0x18, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x75, 0x73, 0x65, 0x72, 0x69, 0x6e, 0x66, 0x6f,
	0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x72, 0x65, 0x64, 0x69,
	0x72, 0x65, 0x63, 0x74, 0x5f, 0x75, 0x72, 0x69, 0x18, 0x19, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x72, 0x65, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x55, 0x72, 0x69, 0x42, 0x35, 0x5a, 0x33, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70,
	0x2f, 0x67, 0x6f, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x63,
	0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x3b, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69,
	0x74, 0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_eolymp_community_configuration_idp_proto_rawDescOnce sync.Once
	file_eolymp_community_configuration_idp_proto_rawDescData = file_eolymp_community_configuration_idp_proto_rawDesc
)

func file_eolymp_community_configuration_idp_proto_rawDescGZIP() []byte {
	file_eolymp_community_configuration_idp_proto_rawDescOnce.Do(func() {
		file_eolymp_community_configuration_idp_proto_rawDescData = protoimpl.X.CompressGZIP(file_eolymp_community_configuration_idp_proto_rawDescData)
	})
	return file_eolymp_community_configuration_idp_proto_rawDescData
}

var file_eolymp_community_configuration_idp_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_eolymp_community_configuration_idp_proto_goTypes = []interface{}{
	(*IdentityProvider)(nil),          // 0: eolymp.community.IdentityProvider
	(*IdentityProvider_Basecamp)(nil), // 1: eolymp.community.IdentityProvider.Basecamp
	(*IdentityProvider_Local)(nil),    // 2: eolymp.community.IdentityProvider.Local
	(*IdentityProvider_OIDC)(nil),     // 3: eolymp.community.IdentityProvider.OIDC
}
var file_eolymp_community_configuration_idp_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_eolymp_community_configuration_idp_proto_init() }
func file_eolymp_community_configuration_idp_proto_init() {
	if File_eolymp_community_configuration_idp_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_eolymp_community_configuration_idp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IdentityProvider); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_eolymp_community_configuration_idp_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IdentityProvider_Basecamp); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_eolymp_community_configuration_idp_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IdentityProvider_Local); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_eolymp_community_configuration_idp_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IdentityProvider_OIDC); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_eolymp_community_configuration_idp_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_eolymp_community_configuration_idp_proto_goTypes,
		DependencyIndexes: file_eolymp_community_configuration_idp_proto_depIdxs,
		MessageInfos:      file_eolymp_community_configuration_idp_proto_msgTypes,
	}.Build()
	File_eolymp_community_configuration_idp_proto = out.File
	file_eolymp_community_configuration_idp_proto_rawDesc = nil
	file_eolymp_community_configuration_idp_proto_goTypes = nil
	file_eolymp_community_configuration_idp_proto_depIdxs = nil
}
