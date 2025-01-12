// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.2
// 	protoc        v4.24.4
// source: eolymp/atlas/editor_service.proto

package atlas

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

type PrintCodeInput struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Source        string                 `protobuf:"bytes,1,opt,name=source,proto3" json:"source,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PrintCodeInput) Reset() {
	*x = PrintCodeInput{}
	mi := &file_eolymp_atlas_editor_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PrintCodeInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PrintCodeInput) ProtoMessage() {}

func (x *PrintCodeInput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_atlas_editor_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PrintCodeInput.ProtoReflect.Descriptor instead.
func (*PrintCodeInput) Descriptor() ([]byte, []int) {
	return file_eolymp_atlas_editor_service_proto_rawDescGZIP(), []int{0}
}

func (x *PrintCodeInput) GetSource() string {
	if x != nil {
		return x.Source
	}
	return ""
}

type PrintCodeOutput struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PrintCodeOutput) Reset() {
	*x = PrintCodeOutput{}
	mi := &file_eolymp_atlas_editor_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PrintCodeOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PrintCodeOutput) ProtoMessage() {}

func (x *PrintCodeOutput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_atlas_editor_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PrintCodeOutput.ProtoReflect.Descriptor instead.
func (*PrintCodeOutput) Descriptor() ([]byte, []int) {
	return file_eolymp_atlas_editor_service_proto_rawDescGZIP(), []int{1}
}

var File_eolymp_atlas_editor_service_proto protoreflect.FileDescriptor

var file_eolymp_atlas_editor_service_proto_rawDesc = []byte{
	0x0a, 0x21, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x61, 0x74, 0x6c, 0x61, 0x73, 0x2f, 0x65,
	0x64, 0x69, 0x74, 0x6f, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x61, 0x74, 0x6c, 0x61,
	0x73, 0x1a, 0x1d, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x22, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x72, 0x61, 0x74, 0x65, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x28, 0x0a, 0x0e, 0x50, 0x72, 0x69, 0x6e, 0x74, 0x43, 0x6f, 0x64,
	0x65, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x22, 0x11,
	0x0a, 0x0f, 0x50, 0x72, 0x69, 0x6e, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x4f, 0x75, 0x74, 0x70, 0x75,
	0x74, 0x32, 0x9e, 0x01, 0x0a, 0x0d, 0x45, 0x64, 0x69, 0x74, 0x6f, 0x72, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x8c, 0x01, 0x0a, 0x09, 0x50, 0x72, 0x69, 0x6e, 0x74, 0x43, 0x6f, 0x64,
	0x65, 0x12, 0x1c, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x61, 0x74, 0x6c, 0x61, 0x73,
	0x2e, 0x50, 0x72, 0x69, 0x6e, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a,
	0x1d, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x61, 0x74, 0x6c, 0x61, 0x73, 0x2e, 0x50,
	0x72, 0x69, 0x6e, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x22, 0x42,
	0xea, 0xe2, 0x0a, 0x0b, 0xf5, 0xe2, 0x0a, 0x00, 0x00, 0x80, 0x3f, 0xf8, 0xe2, 0x0a, 0x02, 0x82,
	0xe3, 0x0a, 0x1a, 0x8a, 0xe3, 0x0a, 0x16, 0x61, 0x74, 0x6c, 0x61, 0x73, 0x3a, 0x73, 0x75, 0x62,
	0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x3a, 0x77, 0x72, 0x69, 0x74, 0x65, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x0f, 0x22, 0x0d, 0x2f, 0x65, 0x64, 0x69, 0x74, 0x6f, 0x72, 0x3a, 0x70, 0x72, 0x69,
	0x6e, 0x74, 0x42, 0x2d, 0x5a, 0x2b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x67, 0x6f, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x65,
	0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x61, 0x74, 0x6c, 0x61, 0x73, 0x3b, 0x61, 0x74, 0x6c, 0x61,
	0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_eolymp_atlas_editor_service_proto_rawDescOnce sync.Once
	file_eolymp_atlas_editor_service_proto_rawDescData = file_eolymp_atlas_editor_service_proto_rawDesc
)

func file_eolymp_atlas_editor_service_proto_rawDescGZIP() []byte {
	file_eolymp_atlas_editor_service_proto_rawDescOnce.Do(func() {
		file_eolymp_atlas_editor_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_eolymp_atlas_editor_service_proto_rawDescData)
	})
	return file_eolymp_atlas_editor_service_proto_rawDescData
}

var file_eolymp_atlas_editor_service_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_eolymp_atlas_editor_service_proto_goTypes = []any{
	(*PrintCodeInput)(nil),  // 0: eolymp.atlas.PrintCodeInput
	(*PrintCodeOutput)(nil), // 1: eolymp.atlas.PrintCodeOutput
}
var file_eolymp_atlas_editor_service_proto_depIdxs = []int32{
	0, // 0: eolymp.atlas.EditorService.PrintCode:input_type -> eolymp.atlas.PrintCodeInput
	1, // 1: eolymp.atlas.EditorService.PrintCode:output_type -> eolymp.atlas.PrintCodeOutput
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_eolymp_atlas_editor_service_proto_init() }
func file_eolymp_atlas_editor_service_proto_init() {
	if File_eolymp_atlas_editor_service_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_eolymp_atlas_editor_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_eolymp_atlas_editor_service_proto_goTypes,
		DependencyIndexes: file_eolymp_atlas_editor_service_proto_depIdxs,
		MessageInfos:      file_eolymp_atlas_editor_service_proto_msgTypes,
	}.Build()
	File_eolymp_atlas_editor_service_proto = out.File
	file_eolymp_atlas_editor_service_proto_rawDesc = nil
	file_eolymp_atlas_editor_service_proto_goTypes = nil
	file_eolymp_atlas_editor_service_proto_depIdxs = nil
}
