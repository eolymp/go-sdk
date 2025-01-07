// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.2
// 	protoc        v4.24.4
// source: eolymp/atlas/testing_config.proto

package atlas

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

type TestingConfig struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	RunCount      uint32                 `protobuf:"varint,1,opt,name=run_count,json=runCount,proto3" json:"run_count,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *TestingConfig) Reset() {
	*x = TestingConfig{}
	mi := &file_eolymp_atlas_testing_config_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TestingConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestingConfig) ProtoMessage() {}

func (x *TestingConfig) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_atlas_testing_config_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestingConfig.ProtoReflect.Descriptor instead.
func (*TestingConfig) Descriptor() ([]byte, []int) {
	return file_eolymp_atlas_testing_config_proto_rawDescGZIP(), []int{0}
}

func (x *TestingConfig) GetRunCount() uint32 {
	if x != nil {
		return x.RunCount
	}
	return 0
}

var File_eolymp_atlas_testing_config_proto protoreflect.FileDescriptor

var file_eolymp_atlas_testing_config_proto_rawDesc = []byte{
	0x0a, 0x21, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x61, 0x74, 0x6c, 0x61, 0x73, 0x2f, 0x74,
	0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x61, 0x74, 0x6c, 0x61,
	0x73, 0x22, 0x2c, 0x0a, 0x0d, 0x54, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x12, 0x1b, 0x0a, 0x09, 0x72, 0x75, 0x6e, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x72, 0x75, 0x6e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x42,
	0x2d, 0x5a, 0x2b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6f,
	0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x67, 0x6f, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x65, 0x6f, 0x6c, 0x79,
	0x6d, 0x70, 0x2f, 0x61, 0x74, 0x6c, 0x61, 0x73, 0x3b, 0x61, 0x74, 0x6c, 0x61, 0x73, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_eolymp_atlas_testing_config_proto_rawDescOnce sync.Once
	file_eolymp_atlas_testing_config_proto_rawDescData = file_eolymp_atlas_testing_config_proto_rawDesc
)

func file_eolymp_atlas_testing_config_proto_rawDescGZIP() []byte {
	file_eolymp_atlas_testing_config_proto_rawDescOnce.Do(func() {
		file_eolymp_atlas_testing_config_proto_rawDescData = protoimpl.X.CompressGZIP(file_eolymp_atlas_testing_config_proto_rawDescData)
	})
	return file_eolymp_atlas_testing_config_proto_rawDescData
}

var file_eolymp_atlas_testing_config_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_eolymp_atlas_testing_config_proto_goTypes = []any{
	(*TestingConfig)(nil), // 0: eolymp.atlas.TestingConfig
}
var file_eolymp_atlas_testing_config_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_eolymp_atlas_testing_config_proto_init() }
func file_eolymp_atlas_testing_config_proto_init() {
	if File_eolymp_atlas_testing_config_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_eolymp_atlas_testing_config_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_eolymp_atlas_testing_config_proto_goTypes,
		DependencyIndexes: file_eolymp_atlas_testing_config_proto_depIdxs,
		MessageInfos:      file_eolymp_atlas_testing_config_proto_msgTypes,
	}.Build()
	File_eolymp_atlas_testing_config_proto = out.File
	file_eolymp_atlas_testing_config_proto_rawDesc = nil
	file_eolymp_atlas_testing_config_proto_goTypes = nil
	file_eolymp_atlas_testing_config_proto_depIdxs = nil
}
