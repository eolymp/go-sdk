// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v4.24.4
// source: eolymp/community/activity_graph_service.proto

package community

import (
	_ "github.com/eolymp/go-sdk/eolymp/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
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

type DescribeActivityGraphInput struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	After         *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=after,proto3" json:"after,omitempty"`
	Before        *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=before,proto3" json:"before,omitempty"`
	Metric        string                 `protobuf:"bytes,3,opt,name=metric,proto3" json:"metric,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DescribeActivityGraphInput) Reset() {
	*x = DescribeActivityGraphInput{}
	mi := &file_eolymp_community_activity_graph_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DescribeActivityGraphInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescribeActivityGraphInput) ProtoMessage() {}

func (x *DescribeActivityGraphInput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_community_activity_graph_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DescribeActivityGraphInput.ProtoReflect.Descriptor instead.
func (*DescribeActivityGraphInput) Descriptor() ([]byte, []int) {
	return file_eolymp_community_activity_graph_service_proto_rawDescGZIP(), []int{0}
}

func (x *DescribeActivityGraphInput) GetAfter() *timestamppb.Timestamp {
	if x != nil {
		return x.After
	}
	return nil
}

func (x *DescribeActivityGraphInput) GetBefore() *timestamppb.Timestamp {
	if x != nil {
		return x.Before
	}
	return nil
}

func (x *DescribeActivityGraphInput) GetMetric() string {
	if x != nil {
		return x.Metric
	}
	return ""
}

type DescribeActivityGraphOutput struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Values        []int32                `protobuf:"varint,1,rep,packed,name=values,proto3" json:"values,omitempty"`
	MaxValue      int32                  `protobuf:"varint,10,opt,name=max_value,json=maxValue,proto3" json:"max_value,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DescribeActivityGraphOutput) Reset() {
	*x = DescribeActivityGraphOutput{}
	mi := &file_eolymp_community_activity_graph_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DescribeActivityGraphOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescribeActivityGraphOutput) ProtoMessage() {}

func (x *DescribeActivityGraphOutput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_community_activity_graph_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DescribeActivityGraphOutput.ProtoReflect.Descriptor instead.
func (*DescribeActivityGraphOutput) Descriptor() ([]byte, []int) {
	return file_eolymp_community_activity_graph_service_proto_rawDescGZIP(), []int{1}
}

func (x *DescribeActivityGraphOutput) GetValues() []int32 {
	if x != nil {
		return x.Values
	}
	return nil
}

func (x *DescribeActivityGraphOutput) GetMaxValue() int32 {
	if x != nil {
		return x.MaxValue
	}
	return 0
}

var File_eolymp_community_activity_graph_service_proto protoreflect.FileDescriptor

var file_eolymp_community_activity_graph_service_proto_rawDesc = string([]byte{
	0x0a, 0x2d, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69,
	0x74, 0x79, 0x2f, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x5f, 0x67, 0x72, 0x61, 0x70,
	0x68, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x10, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74,
	0x79, 0x1a, 0x1d, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x22, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x72, 0x61, 0x74, 0x65, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9a, 0x01, 0x0a, 0x1a, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x62, 0x65, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x47, 0x72, 0x61, 0x70, 0x68, 0x49,
	0x6e, 0x70, 0x75, 0x74, 0x12, 0x30, 0x0a, 0x05, 0x61, 0x66, 0x74, 0x65, 0x72, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x05, 0x61, 0x66, 0x74, 0x65, 0x72, 0x12, 0x32, 0x0a, 0x06, 0x62, 0x65, 0x66, 0x6f, 0x72, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x06, 0x62, 0x65, 0x66, 0x6f, 0x72, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x65,
	0x74, 0x72, 0x69, 0x63, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x65, 0x74, 0x72,
	0x69, 0x63, 0x22, 0x52, 0x0a, 0x1b, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x41, 0x63,
	0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x47, 0x72, 0x61, 0x70, 0x68, 0x4f, 0x75, 0x74, 0x70, 0x75,
	0x74, 0x12, 0x16, 0x0a, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x05, 0x52, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x61, 0x78,
	0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x6d, 0x61,
	0x78, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x32, 0xd2, 0x01, 0x0a, 0x14, 0x41, 0x63, 0x74, 0x69, 0x76,
	0x69, 0x74, 0x79, 0x47, 0x72, 0x61, 0x70, 0x68, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0xb9, 0x01, 0x0a, 0x15, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x41, 0x63, 0x74, 0x69,
	0x76, 0x69, 0x74, 0x79, 0x47, 0x72, 0x61, 0x70, 0x68, 0x12, 0x2c, 0x2e, 0x65, 0x6f, 0x6c, 0x79,
	0x6d, 0x70, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x2e, 0x44, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x62, 0x65, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x47, 0x72, 0x61,
	0x70, 0x68, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x2d, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x2e, 0x44, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x62, 0x65, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x47, 0x72, 0x61, 0x70, 0x68,
	0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x22, 0x43, 0xea, 0xe2, 0x0a, 0x0b, 0xf5, 0xe2, 0x0a, 0x00,
	0x00, 0xa0, 0x40, 0xf8, 0xe2, 0x0a, 0x32, 0x82, 0xe3, 0x0a, 0x19, 0x8a, 0xe3, 0x0a, 0x15, 0x63,
	0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x3a, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x3a,
	0x72, 0x65, 0x61, 0x64, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x11, 0x12, 0x0f, 0x2f, 0x61, 0x63, 0x74,
	0x69, 0x76, 0x69, 0x74, 0x79, 0x2d, 0x67, 0x72, 0x61, 0x70, 0x68, 0x42, 0x35, 0x5a, 0x33, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70,
	0x2f, 0x67, 0x6f, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x63,
	0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x3b, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69,
	0x74, 0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_eolymp_community_activity_graph_service_proto_rawDescOnce sync.Once
	file_eolymp_community_activity_graph_service_proto_rawDescData []byte
)

func file_eolymp_community_activity_graph_service_proto_rawDescGZIP() []byte {
	file_eolymp_community_activity_graph_service_proto_rawDescOnce.Do(func() {
		file_eolymp_community_activity_graph_service_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_eolymp_community_activity_graph_service_proto_rawDesc), len(file_eolymp_community_activity_graph_service_proto_rawDesc)))
	})
	return file_eolymp_community_activity_graph_service_proto_rawDescData
}

var file_eolymp_community_activity_graph_service_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_eolymp_community_activity_graph_service_proto_goTypes = []any{
	(*DescribeActivityGraphInput)(nil),  // 0: eolymp.community.DescribeActivityGraphInput
	(*DescribeActivityGraphOutput)(nil), // 1: eolymp.community.DescribeActivityGraphOutput
	(*timestamppb.Timestamp)(nil),       // 2: google.protobuf.Timestamp
}
var file_eolymp_community_activity_graph_service_proto_depIdxs = []int32{
	2, // 0: eolymp.community.DescribeActivityGraphInput.after:type_name -> google.protobuf.Timestamp
	2, // 1: eolymp.community.DescribeActivityGraphInput.before:type_name -> google.protobuf.Timestamp
	0, // 2: eolymp.community.ActivityGraphService.DescribeActivityGraph:input_type -> eolymp.community.DescribeActivityGraphInput
	1, // 3: eolymp.community.ActivityGraphService.DescribeActivityGraph:output_type -> eolymp.community.DescribeActivityGraphOutput
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_eolymp_community_activity_graph_service_proto_init() }
func file_eolymp_community_activity_graph_service_proto_init() {
	if File_eolymp_community_activity_graph_service_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_eolymp_community_activity_graph_service_proto_rawDesc), len(file_eolymp_community_activity_graph_service_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_eolymp_community_activity_graph_service_proto_goTypes,
		DependencyIndexes: file_eolymp_community_activity_graph_service_proto_depIdxs,
		MessageInfos:      file_eolymp_community_activity_graph_service_proto_msgTypes,
	}.Build()
	File_eolymp_community_activity_graph_service_proto = out.File
	file_eolymp_community_activity_graph_service_proto_goTypes = nil
	file_eolymp_community_activity_graph_service_proto_depIdxs = nil
}
