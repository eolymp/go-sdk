// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.18.1
// source: eolymp/typewriter/typewriter.proto

package typewriter

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

type UploadAssetInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Filename string `protobuf:"bytes,1,opt,name=filename,proto3" json:"filename,omitempty"`
	Data     []byte `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *UploadAssetInput) Reset() {
	*x = UploadAssetInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_typewriter_typewriter_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UploadAssetInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadAssetInput) ProtoMessage() {}

func (x *UploadAssetInput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_typewriter_typewriter_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadAssetInput.ProtoReflect.Descriptor instead.
func (*UploadAssetInput) Descriptor() ([]byte, []int) {
	return file_eolymp_typewriter_typewriter_proto_rawDescGZIP(), []int{0}
}

func (x *UploadAssetInput) GetFilename() string {
	if x != nil {
		return x.Filename
	}
	return ""
}

func (x *UploadAssetInput) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

type UploadAssetOutput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Link string `protobuf:"bytes,1,opt,name=link,proto3" json:"link,omitempty"`
}

func (x *UploadAssetOutput) Reset() {
	*x = UploadAssetOutput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_typewriter_typewriter_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UploadAssetOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadAssetOutput) ProtoMessage() {}

func (x *UploadAssetOutput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_typewriter_typewriter_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadAssetOutput.ProtoReflect.Descriptor instead.
func (*UploadAssetOutput) Descriptor() ([]byte, []int) {
	return file_eolymp_typewriter_typewriter_proto_rawDescGZIP(), []int{1}
}

func (x *UploadAssetOutput) GetLink() string {
	if x != nil {
		return x.Link
	}
	return ""
}

var File_eolymp_typewriter_typewriter_proto protoreflect.FileDescriptor

var file_eolymp_typewriter_typewriter_proto_rawDesc = []byte{
	0x0a, 0x22, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x77, 0x72, 0x69,
	0x74, 0x65, 0x72, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x77, 0x72, 0x69, 0x74, 0x65, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x74, 0x79, 0x70,
	0x65, 0x77, 0x72, 0x69, 0x74, 0x65, 0x72, 0x1a, 0x22, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x72, 0x61, 0x74, 0x65,
	0x6c, 0x69, 0x6d, 0x69, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x65, 0x6f, 0x6c,
	0x79, 0x6d, 0x70, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f,
	0x73, 0x63, 0x6f, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x42, 0x0a, 0x10, 0x55,
	0x70, 0x6c, 0x6f, 0x61, 0x64, 0x41, 0x73, 0x73, 0x65, 0x74, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12,
	0x1a, 0x0a, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22,
	0x27, 0x0a, 0x11, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x41, 0x73, 0x73, 0x65, 0x74, 0x4f, 0x75,
	0x74, 0x70, 0x75, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x69, 0x6e, 0x6b, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6c, 0x69, 0x6e, 0x6b, 0x32, 0x96, 0x01, 0x0a, 0x0a, 0x54, 0x79, 0x70,
	0x65, 0x77, 0x72, 0x69, 0x74, 0x65, 0x72, 0x12, 0x87, 0x01, 0x0a, 0x0b, 0x55, 0x70, 0x6c, 0x6f,
	0x61, 0x64, 0x41, 0x73, 0x73, 0x65, 0x74, 0x12, 0x23, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70,
	0x2e, 0x74, 0x79, 0x70, 0x65, 0x77, 0x72, 0x69, 0x74, 0x65, 0x72, 0x2e, 0x55, 0x70, 0x6c, 0x6f,
	0x61, 0x64, 0x41, 0x73, 0x73, 0x65, 0x74, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x24, 0x2e, 0x65,
	0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x77, 0x72, 0x69, 0x74, 0x65, 0x72,
	0x2e, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x41, 0x73, 0x73, 0x65, 0x74, 0x4f, 0x75, 0x74, 0x70,
	0x75, 0x74, 0x22, 0x2d, 0x82, 0xe3, 0x0a, 0x1a, 0x8a, 0xe3, 0x0a, 0x16, 0x74, 0x79, 0x70, 0x65,
	0x77, 0x72, 0x69, 0x74, 0x65, 0x72, 0x3a, 0x61, 0x73, 0x73, 0x65, 0x74, 0x3a, 0x77, 0x72, 0x69,
	0x74, 0x65, 0xea, 0xe2, 0x0a, 0x0b, 0xf5, 0xe2, 0x0a, 0x00, 0x00, 0x00, 0x40, 0xf8, 0xe2, 0x0a,
	0x64, 0x42, 0x37, 0x5a, 0x35, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x67, 0x6f, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x65, 0x6f,
	0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x77, 0x72, 0x69, 0x74, 0x65, 0x72, 0x3b,
	0x74, 0x79, 0x70, 0x65, 0x77, 0x72, 0x69, 0x74, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_eolymp_typewriter_typewriter_proto_rawDescOnce sync.Once
	file_eolymp_typewriter_typewriter_proto_rawDescData = file_eolymp_typewriter_typewriter_proto_rawDesc
)

func file_eolymp_typewriter_typewriter_proto_rawDescGZIP() []byte {
	file_eolymp_typewriter_typewriter_proto_rawDescOnce.Do(func() {
		file_eolymp_typewriter_typewriter_proto_rawDescData = protoimpl.X.CompressGZIP(file_eolymp_typewriter_typewriter_proto_rawDescData)
	})
	return file_eolymp_typewriter_typewriter_proto_rawDescData
}

var file_eolymp_typewriter_typewriter_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_eolymp_typewriter_typewriter_proto_goTypes = []interface{}{
	(*UploadAssetInput)(nil),  // 0: eolymp.typewriter.UploadAssetInput
	(*UploadAssetOutput)(nil), // 1: eolymp.typewriter.UploadAssetOutput
}
var file_eolymp_typewriter_typewriter_proto_depIdxs = []int32{
	0, // 0: eolymp.typewriter.Typewriter.UploadAsset:input_type -> eolymp.typewriter.UploadAssetInput
	1, // 1: eolymp.typewriter.Typewriter.UploadAsset:output_type -> eolymp.typewriter.UploadAssetOutput
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_eolymp_typewriter_typewriter_proto_init() }
func file_eolymp_typewriter_typewriter_proto_init() {
	if File_eolymp_typewriter_typewriter_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_eolymp_typewriter_typewriter_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UploadAssetInput); i {
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
		file_eolymp_typewriter_typewriter_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UploadAssetOutput); i {
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
			RawDescriptor: file_eolymp_typewriter_typewriter_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_eolymp_typewriter_typewriter_proto_goTypes,
		DependencyIndexes: file_eolymp_typewriter_typewriter_proto_depIdxs,
		MessageInfos:      file_eolymp_typewriter_typewriter_proto_msgTypes,
	}.Build()
	File_eolymp_typewriter_typewriter_proto = out.File
	file_eolymp_typewriter_typewriter_proto_rawDesc = nil
	file_eolymp_typewriter_typewriter_proto_goTypes = nil
	file_eolymp_typewriter_typewriter_proto_depIdxs = nil
}
