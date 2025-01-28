// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.4
// 	protoc        v4.24.4
// source: eolymp/annotations/ratelimit.proto

package annotations

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
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

type RateLimit struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Limit         float32                `protobuf:"fixed32,22062,opt,name=limit,proto3" json:"limit,omitempty"` // number of requests per second
	Burst         int32                  `protobuf:"varint,22063,opt,name=burst,proto3" json:"burst,omitempty"`  // allowed burst per second
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RateLimit) Reset() {
	*x = RateLimit{}
	mi := &file_eolymp_annotations_ratelimit_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RateLimit) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RateLimit) ProtoMessage() {}

func (x *RateLimit) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_annotations_ratelimit_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RateLimit.ProtoReflect.Descriptor instead.
func (*RateLimit) Descriptor() ([]byte, []int) {
	return file_eolymp_annotations_ratelimit_proto_rawDescGZIP(), []int{0}
}

func (x *RateLimit) GetLimit() float32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *RateLimit) GetBurst() int32 {
	if x != nil {
		return x.Burst
	}
	return 0
}

var file_eolymp_annotations_ratelimit_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptorpb.MethodOptions)(nil),
		ExtensionType: (*RateLimit)(nil),
		Field:         22061,
		Name:          "eolymp.api.ratelimit",
		Tag:           "bytes,22061,opt,name=ratelimit",
		Filename:      "eolymp/annotations/ratelimit.proto",
	},
}

// Extension fields to descriptorpb.MethodOptions.
var (
	// optional eolymp.api.RateLimit ratelimit = 22061;
	E_Ratelimit = &file_eolymp_annotations_ratelimit_proto_extTypes[0]
)

var File_eolymp_annotations_ratelimit_proto protoreflect.FileDescriptor

var file_eolymp_annotations_ratelimit_proto_rawDesc = string([]byte{
	0x0a, 0x22, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x72, 0x61, 0x74, 0x65, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x61, 0x70, 0x69,
	0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x3b, 0x0a, 0x09, 0x52, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x12,
	0x16, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0xae, 0xac, 0x01, 0x20, 0x01, 0x28, 0x02,
	0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x16, 0x0a, 0x05, 0x62, 0x75, 0x72, 0x73, 0x74,
	0x18, 0xaf, 0xac, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x62, 0x75, 0x72, 0x73, 0x74, 0x3a,
	0x55, 0x0a, 0x09, 0x72, 0x61, 0x74, 0x65, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x1e, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4d,
	0x65, 0x74, 0x68, 0x6f, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xad, 0xac, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x52, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x52, 0x09, 0x72, 0x61, 0x74,
	0x65, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x42, 0x39, 0x5a, 0x37, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x67, 0x6f, 0x2d, 0x73,
	0x64, 0x6b, 0x2f, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x3b, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_eolymp_annotations_ratelimit_proto_rawDescOnce sync.Once
	file_eolymp_annotations_ratelimit_proto_rawDescData []byte
)

func file_eolymp_annotations_ratelimit_proto_rawDescGZIP() []byte {
	file_eolymp_annotations_ratelimit_proto_rawDescOnce.Do(func() {
		file_eolymp_annotations_ratelimit_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_eolymp_annotations_ratelimit_proto_rawDesc), len(file_eolymp_annotations_ratelimit_proto_rawDesc)))
	})
	return file_eolymp_annotations_ratelimit_proto_rawDescData
}

var file_eolymp_annotations_ratelimit_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_eolymp_annotations_ratelimit_proto_goTypes = []any{
	(*RateLimit)(nil),                  // 0: eolymp.api.RateLimit
	(*descriptorpb.MethodOptions)(nil), // 1: google.protobuf.MethodOptions
}
var file_eolymp_annotations_ratelimit_proto_depIdxs = []int32{
	1, // 0: eolymp.api.ratelimit:extendee -> google.protobuf.MethodOptions
	0, // 1: eolymp.api.ratelimit:type_name -> eolymp.api.RateLimit
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	1, // [1:2] is the sub-list for extension type_name
	0, // [0:1] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_eolymp_annotations_ratelimit_proto_init() }
func file_eolymp_annotations_ratelimit_proto_init() {
	if File_eolymp_annotations_ratelimit_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_eolymp_annotations_ratelimit_proto_rawDesc), len(file_eolymp_annotations_ratelimit_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 1,
			NumServices:   0,
		},
		GoTypes:           file_eolymp_annotations_ratelimit_proto_goTypes,
		DependencyIndexes: file_eolymp_annotations_ratelimit_proto_depIdxs,
		MessageInfos:      file_eolymp_annotations_ratelimit_proto_msgTypes,
		ExtensionInfos:    file_eolymp_annotations_ratelimit_proto_extTypes,
	}.Build()
	File_eolymp_annotations_ratelimit_proto = out.File
	file_eolymp_annotations_ratelimit_proto_goTypes = nil
	file_eolymp_annotations_ratelimit_proto_depIdxs = nil
}
