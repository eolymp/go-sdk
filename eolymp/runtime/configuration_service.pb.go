// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v3.21.12
// source: eolymp/runtime/configuration_service.proto

package runtime

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

type DescribeRuntimeConfigInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DescribeRuntimeConfigInput) Reset() {
	*x = DescribeRuntimeConfigInput{}
	mi := &file_eolymp_runtime_configuration_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DescribeRuntimeConfigInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescribeRuntimeConfigInput) ProtoMessage() {}

func (x *DescribeRuntimeConfigInput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_runtime_configuration_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DescribeRuntimeConfigInput.ProtoReflect.Descriptor instead.
func (*DescribeRuntimeConfigInput) Descriptor() ([]byte, []int) {
	return file_eolymp_runtime_configuration_service_proto_rawDescGZIP(), []int{0}
}

type DescribeRuntimeConfigOutput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AllowedRuntime []string `protobuf:"bytes,1,rep,name=allowed_runtime,json=allowedRuntime,proto3" json:"allowed_runtime,omitempty"`
}

func (x *DescribeRuntimeConfigOutput) Reset() {
	*x = DescribeRuntimeConfigOutput{}
	mi := &file_eolymp_runtime_configuration_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DescribeRuntimeConfigOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescribeRuntimeConfigOutput) ProtoMessage() {}

func (x *DescribeRuntimeConfigOutput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_runtime_configuration_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DescribeRuntimeConfigOutput.ProtoReflect.Descriptor instead.
func (*DescribeRuntimeConfigOutput) Descriptor() ([]byte, []int) {
	return file_eolymp_runtime_configuration_service_proto_rawDescGZIP(), []int{1}
}

func (x *DescribeRuntimeConfigOutput) GetAllowedRuntime() []string {
	if x != nil {
		return x.AllowedRuntime
	}
	return nil
}

type ConfigureRuntimeConfigInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AllowedRuntime []string `protobuf:"bytes,1,rep,name=allowed_runtime,json=allowedRuntime,proto3" json:"allowed_runtime,omitempty"`
}

func (x *ConfigureRuntimeConfigInput) Reset() {
	*x = ConfigureRuntimeConfigInput{}
	mi := &file_eolymp_runtime_configuration_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ConfigureRuntimeConfigInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfigureRuntimeConfigInput) ProtoMessage() {}

func (x *ConfigureRuntimeConfigInput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_runtime_configuration_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfigureRuntimeConfigInput.ProtoReflect.Descriptor instead.
func (*ConfigureRuntimeConfigInput) Descriptor() ([]byte, []int) {
	return file_eolymp_runtime_configuration_service_proto_rawDescGZIP(), []int{2}
}

func (x *ConfigureRuntimeConfigInput) GetAllowedRuntime() []string {
	if x != nil {
		return x.AllowedRuntime
	}
	return nil
}

type ConfigureRuntimeConfigOutput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ConfigureRuntimeConfigOutput) Reset() {
	*x = ConfigureRuntimeConfigOutput{}
	mi := &file_eolymp_runtime_configuration_service_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ConfigureRuntimeConfigOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfigureRuntimeConfigOutput) ProtoMessage() {}

func (x *ConfigureRuntimeConfigOutput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_runtime_configuration_service_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfigureRuntimeConfigOutput.ProtoReflect.Descriptor instead.
func (*ConfigureRuntimeConfigOutput) Descriptor() ([]byte, []int) {
	return file_eolymp_runtime_configuration_service_proto_rawDescGZIP(), []int{3}
}

var File_eolymp_runtime_configuration_service_proto protoreflect.FileDescriptor

var file_eolymp_runtime_configuration_service_proto_rawDesc = []byte{
	0x0a, 0x2a, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65,
	0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x65, 0x6f,
	0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x1a, 0x1d, 0x65, 0x6f,
	0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2f, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x22, 0x65, 0x6f, 0x6c,
	0x79, 0x6d, 0x70, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f,
	0x72, 0x61, 0x74, 0x65, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x1c, 0x0a, 0x1a, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x52, 0x75, 0x6e, 0x74, 0x69,
	0x6d, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x22, 0x46, 0x0a,
	0x1b, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x52, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x12, 0x27, 0x0a, 0x0f,
	0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64, 0x5f, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0e, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64, 0x52, 0x75,
	0x6e, 0x74, 0x69, 0x6d, 0x65, 0x22, 0x46, 0x0a, 0x1b, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75,
	0x72, 0x65, 0x52, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x49,
	0x6e, 0x70, 0x75, 0x74, 0x12, 0x27, 0x0a, 0x0f, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64, 0x5f,
	0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0e, 0x61,
	0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64, 0x52, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x22, 0x1e, 0x0a,
	0x1c, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x65, 0x52, 0x75, 0x6e, 0x74, 0x69, 0x6d,
	0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x32, 0xee, 0x02,
	0x0a, 0x14, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x99, 0x01, 0x0a, 0x15, 0x44, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x62, 0x65, 0x52, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x12, 0x2a, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d,
	0x65, 0x2e, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x52, 0x75, 0x6e, 0x74, 0x69, 0x6d,
	0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x2b, 0x2e, 0x65,
	0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x44, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x52, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x22, 0x27, 0xea, 0xe2, 0x0a, 0x0b, 0xf5,
	0xe2, 0x0a, 0x00, 0x00, 0x20, 0x41, 0xf8, 0xe2, 0x0a, 0x64, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x12,
	0x12, 0x10, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x2f, 0x72, 0x75, 0x6e, 0x74, 0x69,
	0x6d, 0x65, 0x12, 0xb9, 0x01, 0x0a, 0x16, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x65,
	0x52, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x2b, 0x2e,
	0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x65, 0x52, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x2c, 0x2e, 0x65, 0x6f, 0x6c,
	0x79, 0x6d, 0x70, 0x2e, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x75, 0x72, 0x65, 0x52, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x22, 0x44, 0xea, 0xe2, 0x0a, 0x0b, 0xf5, 0xe2,
	0x0a, 0x00, 0x00, 0x80, 0x3f, 0xf8, 0xe2, 0x0a, 0x05, 0x82, 0xe3, 0x0a, 0x19, 0x8a, 0xe3, 0x0a,
	0x15, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x3a, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65,
	0x3a, 0x77, 0x72, 0x69, 0x74, 0x65, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x12, 0x1a, 0x10, 0x2f, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x2f, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x42, 0x31,
	0x5a, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6f, 0x6c,
	0x79, 0x6d, 0x70, 0x2f, 0x67, 0x6f, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x65, 0x6f, 0x6c, 0x79, 0x6d,
	0x70, 0x2f, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x3b, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d,
	0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_eolymp_runtime_configuration_service_proto_rawDescOnce sync.Once
	file_eolymp_runtime_configuration_service_proto_rawDescData = file_eolymp_runtime_configuration_service_proto_rawDesc
)

func file_eolymp_runtime_configuration_service_proto_rawDescGZIP() []byte {
	file_eolymp_runtime_configuration_service_proto_rawDescOnce.Do(func() {
		file_eolymp_runtime_configuration_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_eolymp_runtime_configuration_service_proto_rawDescData)
	})
	return file_eolymp_runtime_configuration_service_proto_rawDescData
}

var file_eolymp_runtime_configuration_service_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_eolymp_runtime_configuration_service_proto_goTypes = []any{
	(*DescribeRuntimeConfigInput)(nil),   // 0: eolymp.runtime.DescribeRuntimeConfigInput
	(*DescribeRuntimeConfigOutput)(nil),  // 1: eolymp.runtime.DescribeRuntimeConfigOutput
	(*ConfigureRuntimeConfigInput)(nil),  // 2: eolymp.runtime.ConfigureRuntimeConfigInput
	(*ConfigureRuntimeConfigOutput)(nil), // 3: eolymp.runtime.ConfigureRuntimeConfigOutput
}
var file_eolymp_runtime_configuration_service_proto_depIdxs = []int32{
	0, // 0: eolymp.runtime.ConfigurationService.DescribeRuntimeConfig:input_type -> eolymp.runtime.DescribeRuntimeConfigInput
	2, // 1: eolymp.runtime.ConfigurationService.ConfigureRuntimeConfig:input_type -> eolymp.runtime.ConfigureRuntimeConfigInput
	1, // 2: eolymp.runtime.ConfigurationService.DescribeRuntimeConfig:output_type -> eolymp.runtime.DescribeRuntimeConfigOutput
	3, // 3: eolymp.runtime.ConfigurationService.ConfigureRuntimeConfig:output_type -> eolymp.runtime.ConfigureRuntimeConfigOutput
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_eolymp_runtime_configuration_service_proto_init() }
func file_eolymp_runtime_configuration_service_proto_init() {
	if File_eolymp_runtime_configuration_service_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_eolymp_runtime_configuration_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_eolymp_runtime_configuration_service_proto_goTypes,
		DependencyIndexes: file_eolymp_runtime_configuration_service_proto_depIdxs,
		MessageInfos:      file_eolymp_runtime_configuration_service_proto_msgTypes,
	}.Build()
	File_eolymp_runtime_configuration_service_proto = out.File
	file_eolymp_runtime_configuration_service_proto_rawDesc = nil
	file_eolymp_runtime_configuration_service_proto_goTypes = nil
	file_eolymp_runtime_configuration_service_proto_depIdxs = nil
}
