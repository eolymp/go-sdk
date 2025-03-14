// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v4.24.4
// source: eolymp/community/configuration_service.proto

package community

import (
	_ "github.com/eolymp/go-sdk/eolymp/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type DescribeIdentityConfigInput struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DescribeIdentityConfigInput) Reset() {
	*x = DescribeIdentityConfigInput{}
	mi := &file_eolymp_community_configuration_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DescribeIdentityConfigInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescribeIdentityConfigInput) ProtoMessage() {}

func (x *DescribeIdentityConfigInput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_community_configuration_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DescribeIdentityConfigInput.ProtoReflect.Descriptor instead.
func (*DescribeIdentityConfigInput) Descriptor() ([]byte, []int) {
	return file_eolymp_community_configuration_service_proto_rawDescGZIP(), []int{0}
}

type DescribeIdentityConfigOutput struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Config        *IdentityConfig        `protobuf:"bytes,1,opt,name=config,proto3" json:"config,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DescribeIdentityConfigOutput) Reset() {
	*x = DescribeIdentityConfigOutput{}
	mi := &file_eolymp_community_configuration_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DescribeIdentityConfigOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescribeIdentityConfigOutput) ProtoMessage() {}

func (x *DescribeIdentityConfigOutput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_community_configuration_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DescribeIdentityConfigOutput.ProtoReflect.Descriptor instead.
func (*DescribeIdentityConfigOutput) Descriptor() ([]byte, []int) {
	return file_eolymp_community_configuration_service_proto_rawDescGZIP(), []int{1}
}

func (x *DescribeIdentityConfigOutput) GetConfig() *IdentityConfig {
	if x != nil {
		return x.Config
	}
	return nil
}

type ConfigureIdentityConfigInput struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Config        *IdentityConfig        `protobuf:"bytes,1,opt,name=config,proto3" json:"config,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ConfigureIdentityConfigInput) Reset() {
	*x = ConfigureIdentityConfigInput{}
	mi := &file_eolymp_community_configuration_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ConfigureIdentityConfigInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfigureIdentityConfigInput) ProtoMessage() {}

func (x *ConfigureIdentityConfigInput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_community_configuration_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfigureIdentityConfigInput.ProtoReflect.Descriptor instead.
func (*ConfigureIdentityConfigInput) Descriptor() ([]byte, []int) {
	return file_eolymp_community_configuration_service_proto_rawDescGZIP(), []int{2}
}

func (x *ConfigureIdentityConfigInput) GetConfig() *IdentityConfig {
	if x != nil {
		return x.Config
	}
	return nil
}

type ConfigureIdentityConfigOutput struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ConfigureIdentityConfigOutput) Reset() {
	*x = ConfigureIdentityConfigOutput{}
	mi := &file_eolymp_community_configuration_service_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ConfigureIdentityConfigOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfigureIdentityConfigOutput) ProtoMessage() {}

func (x *ConfigureIdentityConfigOutput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_community_configuration_service_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfigureIdentityConfigOutput.ProtoReflect.Descriptor instead.
func (*ConfigureIdentityConfigOutput) Descriptor() ([]byte, []int) {
	return file_eolymp_community_configuration_service_proto_rawDescGZIP(), []int{3}
}

var File_eolymp_community_configuration_service_proto protoreflect.FileDescriptor

var file_eolymp_community_configuration_service_proto_rawDesc = string([]byte{
	0x0a, 0x2c, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69,
	0x74, 0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10,
	0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79,
	0x1a, 0x1d, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x22, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2f, 0x72, 0x61, 0x74, 0x65, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x2d, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x63, 0x6f, 0x6d, 0x6d,
	0x75, 0x6e, 0x69, 0x74, 0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x1d, 0x0a, 0x1b, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x49, 0x64,
	0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x49, 0x6e, 0x70, 0x75,
	0x74, 0x22, 0x58, 0x0a, 0x1c, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x49, 0x64, 0x65,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x4f, 0x75, 0x74, 0x70, 0x75,
	0x74, 0x12, 0x38, 0x0a, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x20, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x75,
	0x6e, 0x69, 0x74, 0x79, 0x2e, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x52, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x22, 0x58, 0x0a, 0x1c, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x65, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x38, 0x0a, 0x06, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x65, 0x6f,
	0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x2e, 0x49,
	0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x06, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x22, 0x1f, 0x0a, 0x1d, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75,
	0x72, 0x65, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x32, 0xff, 0x02, 0x0a, 0x14, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0xa1, 0x01, 0x0a, 0x16, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x49, 0x64, 0x65, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x2d, 0x2e, 0x65, 0x6f, 0x6c,
	0x79, 0x6d, 0x70, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x2e, 0x44, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x2e, 0x2e, 0x65, 0x6f, 0x6c, 0x79,
	0x6d, 0x70, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x2e, 0x44, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x62, 0x65, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x22, 0x28, 0xea, 0xe2, 0x0a, 0x0b, 0xf5,
	0xe2, 0x0a, 0x00, 0x00, 0x20, 0x41, 0xf8, 0xe2, 0x0a, 0x64, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x13,
	0x12, 0x11, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x2f, 0x69, 0x64, 0x65, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x12, 0xc2, 0x01, 0x0a, 0x17, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72,
	0x65, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12,
	0x2e, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69,
	0x74, 0x79, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x65, 0x49, 0x64, 0x65, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a,
	0x2f, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69,
	0x74, 0x79, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x65, 0x49, 0x64, 0x65, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74,
	0x22, 0x46, 0xea, 0xe2, 0x0a, 0x0b, 0xf5, 0xe2, 0x0a, 0x00, 0x00, 0x80, 0x3f, 0xf8, 0xe2, 0x0a,
	0x05, 0x82, 0xe3, 0x0a, 0x1a, 0x8a, 0xe3, 0x0a, 0x16, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69,
	0x74, 0x79, 0x3a, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x3a, 0x77, 0x72, 0x69, 0x74, 0x65, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x13, 0x1a, 0x11, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x2f,
	0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x42, 0x35, 0x5a, 0x33, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x67, 0x6f,
	0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x63, 0x6f, 0x6d, 0x6d,
	0x75, 0x6e, 0x69, 0x74, 0x79, 0x3b, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_eolymp_community_configuration_service_proto_rawDescOnce sync.Once
	file_eolymp_community_configuration_service_proto_rawDescData []byte
)

func file_eolymp_community_configuration_service_proto_rawDescGZIP() []byte {
	file_eolymp_community_configuration_service_proto_rawDescOnce.Do(func() {
		file_eolymp_community_configuration_service_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_eolymp_community_configuration_service_proto_rawDesc), len(file_eolymp_community_configuration_service_proto_rawDesc)))
	})
	return file_eolymp_community_configuration_service_proto_rawDescData
}

var file_eolymp_community_configuration_service_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_eolymp_community_configuration_service_proto_goTypes = []any{
	(*DescribeIdentityConfigInput)(nil),   // 0: eolymp.community.DescribeIdentityConfigInput
	(*DescribeIdentityConfigOutput)(nil),  // 1: eolymp.community.DescribeIdentityConfigOutput
	(*ConfigureIdentityConfigInput)(nil),  // 2: eolymp.community.ConfigureIdentityConfigInput
	(*ConfigureIdentityConfigOutput)(nil), // 3: eolymp.community.ConfigureIdentityConfigOutput
	(*IdentityConfig)(nil),                // 4: eolymp.community.IdentityConfig
}
var file_eolymp_community_configuration_service_proto_depIdxs = []int32{
	4, // 0: eolymp.community.DescribeIdentityConfigOutput.config:type_name -> eolymp.community.IdentityConfig
	4, // 1: eolymp.community.ConfigureIdentityConfigInput.config:type_name -> eolymp.community.IdentityConfig
	0, // 2: eolymp.community.ConfigurationService.DescribeIdentityConfig:input_type -> eolymp.community.DescribeIdentityConfigInput
	2, // 3: eolymp.community.ConfigurationService.ConfigureIdentityConfig:input_type -> eolymp.community.ConfigureIdentityConfigInput
	1, // 4: eolymp.community.ConfigurationService.DescribeIdentityConfig:output_type -> eolymp.community.DescribeIdentityConfigOutput
	3, // 5: eolymp.community.ConfigurationService.ConfigureIdentityConfig:output_type -> eolymp.community.ConfigureIdentityConfigOutput
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_eolymp_community_configuration_service_proto_init() }
func file_eolymp_community_configuration_service_proto_init() {
	if File_eolymp_community_configuration_service_proto != nil {
		return
	}
	file_eolymp_community_configuration_identity_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_eolymp_community_configuration_service_proto_rawDesc), len(file_eolymp_community_configuration_service_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_eolymp_community_configuration_service_proto_goTypes,
		DependencyIndexes: file_eolymp_community_configuration_service_proto_depIdxs,
		MessageInfos:      file_eolymp_community_configuration_service_proto_msgTypes,
	}.Build()
	File_eolymp_community_configuration_service_proto = out.File
	file_eolymp_community_configuration_service_proto_goTypes = nil
	file_eolymp_community_configuration_service_proto_depIdxs = nil
}
