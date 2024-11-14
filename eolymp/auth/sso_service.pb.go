// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v3.21.12
// source: eolymp/auth/sso_service.proto

package auth

import (
	_ "github.com/eolymp/go-sdk/eolymp/annotations"
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

type AuthorizeRequestInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClientId            string `protobuf:"bytes,1,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	CodeChallenge       string `protobuf:"bytes,2,opt,name=code_challenge,json=codeChallenge,proto3" json:"code_challenge,omitempty"`
	CodeChallengeMethod string `protobuf:"bytes,3,opt,name=code_challenge_method,json=codeChallengeMethod,proto3" json:"code_challenge_method,omitempty"`
	RedirectUri         string `protobuf:"bytes,4,opt,name=redirect_uri,json=redirectUri,proto3" json:"redirect_uri,omitempty"`
	ResponseType        string `protobuf:"bytes,5,opt,name=response_type,json=responseType,proto3" json:"response_type,omitempty"`
	Scope               string `protobuf:"bytes,6,opt,name=scope,proto3" json:"scope,omitempty"`
	State               string `protobuf:"bytes,7,opt,name=state,proto3" json:"state,omitempty"`
}

func (x *AuthorizeRequestInput) Reset() {
	*x = AuthorizeRequestInput{}
	mi := &file_eolymp_auth_sso_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AuthorizeRequestInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthorizeRequestInput) ProtoMessage() {}

func (x *AuthorizeRequestInput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_auth_sso_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthorizeRequestInput.ProtoReflect.Descriptor instead.
func (*AuthorizeRequestInput) Descriptor() ([]byte, []int) {
	return file_eolymp_auth_sso_service_proto_rawDescGZIP(), []int{0}
}

func (x *AuthorizeRequestInput) GetClientId() string {
	if x != nil {
		return x.ClientId
	}
	return ""
}

func (x *AuthorizeRequestInput) GetCodeChallenge() string {
	if x != nil {
		return x.CodeChallenge
	}
	return ""
}

func (x *AuthorizeRequestInput) GetCodeChallengeMethod() string {
	if x != nil {
		return x.CodeChallengeMethod
	}
	return ""
}

func (x *AuthorizeRequestInput) GetRedirectUri() string {
	if x != nil {
		return x.RedirectUri
	}
	return ""
}

func (x *AuthorizeRequestInput) GetResponseType() string {
	if x != nil {
		return x.ResponseType
	}
	return ""
}

func (x *AuthorizeRequestInput) GetScope() string {
	if x != nil {
		return x.Scope
	}
	return ""
}

func (x *AuthorizeRequestInput) GetState() string {
	if x != nil {
		return x.State
	}
	return ""
}

type AuthorizeRequestOutput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RedirectUri string `protobuf:"bytes,1,opt,name=redirect_uri,json=redirectUri,proto3" json:"redirect_uri,omitempty"`
}

func (x *AuthorizeRequestOutput) Reset() {
	*x = AuthorizeRequestOutput{}
	mi := &file_eolymp_auth_sso_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AuthorizeRequestOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthorizeRequestOutput) ProtoMessage() {}

func (x *AuthorizeRequestOutput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_auth_sso_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthorizeRequestOutput.ProtoReflect.Descriptor instead.
func (*AuthorizeRequestOutput) Descriptor() ([]byte, []int) {
	return file_eolymp_auth_sso_service_proto_rawDescGZIP(), []int{1}
}

func (x *AuthorizeRequestOutput) GetRedirectUri() string {
	if x != nil {
		return x.RedirectUri
	}
	return ""
}

type AuthorizeCallbackInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code  string `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	State string `protobuf:"bytes,2,opt,name=state,proto3" json:"state,omitempty"`
}

func (x *AuthorizeCallbackInput) Reset() {
	*x = AuthorizeCallbackInput{}
	mi := &file_eolymp_auth_sso_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AuthorizeCallbackInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthorizeCallbackInput) ProtoMessage() {}

func (x *AuthorizeCallbackInput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_auth_sso_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthorizeCallbackInput.ProtoReflect.Descriptor instead.
func (*AuthorizeCallbackInput) Descriptor() ([]byte, []int) {
	return file_eolymp_auth_sso_service_proto_rawDescGZIP(), []int{2}
}

func (x *AuthorizeCallbackInput) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *AuthorizeCallbackInput) GetState() string {
	if x != nil {
		return x.State
	}
	return ""
}

type AuthorizeCallbackOutput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RedirectUri string `protobuf:"bytes,1,opt,name=redirect_uri,json=redirectUri,proto3" json:"redirect_uri,omitempty"`
}

func (x *AuthorizeCallbackOutput) Reset() {
	*x = AuthorizeCallbackOutput{}
	mi := &file_eolymp_auth_sso_service_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AuthorizeCallbackOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthorizeCallbackOutput) ProtoMessage() {}

func (x *AuthorizeCallbackOutput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_auth_sso_service_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthorizeCallbackOutput.ProtoReflect.Descriptor instead.
func (*AuthorizeCallbackOutput) Descriptor() ([]byte, []int) {
	return file_eolymp_auth_sso_service_proto_rawDescGZIP(), []int{3}
}

func (x *AuthorizeCallbackOutput) GetRedirectUri() string {
	if x != nil {
		return x.RedirectUri
	}
	return ""
}

var File_eolymp_auth_sso_service_proto protoreflect.FileDescriptor

var file_eolymp_auth_sso_service_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x73, 0x73,
	0x6f, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0b, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x1a, 0x22, 0x65, 0x6f,
	0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2f, 0x72, 0x61, 0x74, 0x65, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x83, 0x02, 0x0a, 0x15, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x25, 0x0a, 0x0e, 0x63, 0x6f, 0x64, 0x65, 0x5f,
	0x63, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0d, 0x63, 0x6f, 0x64, 0x65, 0x43, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x12, 0x32,
	0x0a, 0x15, 0x63, 0x6f, 0x64, 0x65, 0x5f, 0x63, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65,
	0x5f, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x13, 0x63,
	0x6f, 0x64, 0x65, 0x43, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x4d, 0x65, 0x74, 0x68,
	0x6f, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x72, 0x65, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x5f, 0x75,
	0x72, 0x69, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x72, 0x65, 0x64, 0x69, 0x72, 0x65,
	0x63, 0x74, 0x55, 0x72, 0x69, 0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x72, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x63,
	0x6f, 0x70, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x63, 0x6f, 0x70, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x22, 0x3b, 0x0a, 0x16, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72,
	0x69, 0x7a, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74,
	0x12, 0x21, 0x0a, 0x0c, 0x72, 0x65, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x5f, 0x75, 0x72, 0x69,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x72, 0x65, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74,
	0x55, 0x72, 0x69, 0x22, 0x42, 0x0a, 0x16, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65,
	0x43, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x12, 0x0a,
	0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x22, 0x3c, 0x0a, 0x17, 0x41, 0x75, 0x74, 0x68, 0x6f,
	0x72, 0x69, 0x7a, 0x65, 0x43, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x4f, 0x75, 0x74, 0x70,
	0x75, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x72, 0x65, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x5f, 0x75,
	0x72, 0x69, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x72, 0x65, 0x64, 0x69, 0x72, 0x65,
	0x63, 0x74, 0x55, 0x72, 0x69, 0x32, 0xeb, 0x01, 0x0a, 0x0a, 0x53, 0x53, 0x4f, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x6c, 0x0a, 0x10, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x22, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d,
	0x70, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x23, 0x2e, 0x65,
	0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x6f,
	0x72, 0x69, 0x7a, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4f, 0x75, 0x74, 0x70, 0x75,
	0x74, 0x22, 0x0f, 0xea, 0xe2, 0x0a, 0x0b, 0xf5, 0xe2, 0x0a, 0x00, 0x00, 0x20, 0x41, 0xf8, 0xe2,
	0x0a, 0x32, 0x12, 0x6f, 0x0a, 0x11, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x43,
	0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x12, 0x23, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70,
	0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x43,
	0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x24, 0x2e, 0x65,
	0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x6f,
	0x72, 0x69, 0x7a, 0x65, 0x43, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x4f, 0x75, 0x74, 0x70,
	0x75, 0x74, 0x22, 0x0f, 0xea, 0xe2, 0x0a, 0x0b, 0xf5, 0xe2, 0x0a, 0x00, 0x00, 0x20, 0x41, 0xf8,
	0xe2, 0x0a, 0x32, 0x42, 0x2b, 0x5a, 0x29, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x67, 0x6f, 0x2d, 0x73, 0x64, 0x6b, 0x2f,
	0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x3b, 0x61, 0x75, 0x74, 0x68,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_eolymp_auth_sso_service_proto_rawDescOnce sync.Once
	file_eolymp_auth_sso_service_proto_rawDescData = file_eolymp_auth_sso_service_proto_rawDesc
)

func file_eolymp_auth_sso_service_proto_rawDescGZIP() []byte {
	file_eolymp_auth_sso_service_proto_rawDescOnce.Do(func() {
		file_eolymp_auth_sso_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_eolymp_auth_sso_service_proto_rawDescData)
	})
	return file_eolymp_auth_sso_service_proto_rawDescData
}

var file_eolymp_auth_sso_service_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_eolymp_auth_sso_service_proto_goTypes = []any{
	(*AuthorizeRequestInput)(nil),   // 0: eolymp.auth.AuthorizeRequestInput
	(*AuthorizeRequestOutput)(nil),  // 1: eolymp.auth.AuthorizeRequestOutput
	(*AuthorizeCallbackInput)(nil),  // 2: eolymp.auth.AuthorizeCallbackInput
	(*AuthorizeCallbackOutput)(nil), // 3: eolymp.auth.AuthorizeCallbackOutput
}
var file_eolymp_auth_sso_service_proto_depIdxs = []int32{
	0, // 0: eolymp.auth.SSOService.AuthorizeRequest:input_type -> eolymp.auth.AuthorizeRequestInput
	2, // 1: eolymp.auth.SSOService.AuthorizeCallback:input_type -> eolymp.auth.AuthorizeCallbackInput
	1, // 2: eolymp.auth.SSOService.AuthorizeRequest:output_type -> eolymp.auth.AuthorizeRequestOutput
	3, // 3: eolymp.auth.SSOService.AuthorizeCallback:output_type -> eolymp.auth.AuthorizeCallbackOutput
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_eolymp_auth_sso_service_proto_init() }
func file_eolymp_auth_sso_service_proto_init() {
	if File_eolymp_auth_sso_service_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_eolymp_auth_sso_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_eolymp_auth_sso_service_proto_goTypes,
		DependencyIndexes: file_eolymp_auth_sso_service_proto_depIdxs,
		MessageInfos:      file_eolymp_auth_sso_service_proto_msgTypes,
	}.Build()
	File_eolymp_auth_sso_service_proto = out.File
	file_eolymp_auth_sso_service_proto_rawDesc = nil
	file_eolymp_auth_sso_service_proto_goTypes = nil
	file_eolymp_auth_sso_service_proto_depIdxs = nil
}
