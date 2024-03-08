// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v3.18.1
// source: eolymp/feed/feed_service.proto

package feed

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

type ListEntriesInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Size  int32  `protobuf:"varint,11,opt,name=size,proto3" json:"size,omitempty"`
	After string `protobuf:"bytes,12,opt,name=after,proto3" json:"after,omitempty"`
}

func (x *ListEntriesInput) Reset() {
	*x = ListEntriesInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_feed_feed_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListEntriesInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListEntriesInput) ProtoMessage() {}

func (x *ListEntriesInput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_feed_feed_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListEntriesInput.ProtoReflect.Descriptor instead.
func (*ListEntriesInput) Descriptor() ([]byte, []int) {
	return file_eolymp_feed_feed_service_proto_rawDescGZIP(), []int{0}
}

func (x *ListEntriesInput) GetSize() int32 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *ListEntriesInput) GetAfter() string {
	if x != nil {
		return x.After
	}
	return ""
}

type ListEntriesOutput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Total          int32    `protobuf:"varint,1,opt,name=total,proto3" json:"total,omitempty"`
	Items          []*Entry `protobuf:"bytes,2,rep,name=items,proto3" json:"items,omitempty"`
	NextPageCursor string   `protobuf:"bytes,3,opt,name=next_page_cursor,json=nextPageCursor,proto3" json:"next_page_cursor,omitempty"`
}

func (x *ListEntriesOutput) Reset() {
	*x = ListEntriesOutput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_feed_feed_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListEntriesOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListEntriesOutput) ProtoMessage() {}

func (x *ListEntriesOutput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_feed_feed_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListEntriesOutput.ProtoReflect.Descriptor instead.
func (*ListEntriesOutput) Descriptor() ([]byte, []int) {
	return file_eolymp_feed_feed_service_proto_rawDescGZIP(), []int{1}
}

func (x *ListEntriesOutput) GetTotal() int32 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *ListEntriesOutput) GetItems() []*Entry {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *ListEntriesOutput) GetNextPageCursor() string {
	if x != nil {
		return x.NextPageCursor
	}
	return ""
}

var File_eolymp_feed_feed_service_proto protoreflect.FileDescriptor

var file_eolymp_feed_feed_service_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x66, 0x65, 0x65, 0x64, 0x2f, 0x66, 0x65,
	0x65, 0x64, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0b, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x66, 0x65, 0x65, 0x64, 0x1a, 0x1d, 0x65,
	0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2f, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x22, 0x65, 0x6f,
	0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2f, 0x72, 0x61, 0x74, 0x65, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x17, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x66, 0x65, 0x65, 0x64, 0x2f, 0x65, 0x6e,
	0x74, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x3c, 0x0a, 0x10, 0x4c, 0x69, 0x73,
	0x74, 0x45, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x12, 0x0a,
	0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x73, 0x69, 0x7a,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x66, 0x74, 0x65, 0x72, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x61, 0x66, 0x74, 0x65, 0x72, 0x22, 0x7d, 0x0a, 0x11, 0x4c, 0x69, 0x73, 0x74, 0x45,
	0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x12, 0x14, 0x0a, 0x05,
	0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x74, 0x6f, 0x74,
	0x61, 0x6c, 0x12, 0x28, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x12, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x66, 0x65, 0x65, 0x64, 0x2e,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x28, 0x0a, 0x10,
	0x6e, 0x65, 0x78, 0x74, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x63, 0x75, 0x72, 0x73, 0x6f, 0x72,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x6e, 0x65, 0x78, 0x74, 0x50, 0x61, 0x67, 0x65,
	0x43, 0x75, 0x72, 0x73, 0x6f, 0x72, 0x32, 0x79, 0x0a, 0x0b, 0x46, 0x65, 0x65, 0x64, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x6a, 0x0a, 0x0b, 0x4c, 0x69, 0x73, 0x74, 0x45, 0x6e, 0x74,
	0x72, 0x69, 0x65, 0x73, 0x12, 0x1d, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x66, 0x65,
	0x65, 0x64, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x49, 0x6e,
	0x70, 0x75, 0x74, 0x1a, 0x1e, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x66, 0x65, 0x65,
	0x64, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x4f, 0x75, 0x74,
	0x70, 0x75, 0x74, 0x22, 0x1c, 0xea, 0xe2, 0x0a, 0x0b, 0xf5, 0xe2, 0x0a, 0x00, 0x00, 0xa0, 0x40,
	0xf8, 0xe2, 0x0a, 0x14, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x07, 0x12, 0x05, 0x2f, 0x66, 0x65, 0x65,
	0x64, 0x42, 0x2b, 0x5a, 0x29, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x67, 0x6f, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x65, 0x6f,
	0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x66, 0x65, 0x65, 0x64, 0x3b, 0x66, 0x65, 0x65, 0x64, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_eolymp_feed_feed_service_proto_rawDescOnce sync.Once
	file_eolymp_feed_feed_service_proto_rawDescData = file_eolymp_feed_feed_service_proto_rawDesc
)

func file_eolymp_feed_feed_service_proto_rawDescGZIP() []byte {
	file_eolymp_feed_feed_service_proto_rawDescOnce.Do(func() {
		file_eolymp_feed_feed_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_eolymp_feed_feed_service_proto_rawDescData)
	})
	return file_eolymp_feed_feed_service_proto_rawDescData
}

var file_eolymp_feed_feed_service_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_eolymp_feed_feed_service_proto_goTypes = []interface{}{
	(*ListEntriesInput)(nil),  // 0: eolymp.feed.ListEntriesInput
	(*ListEntriesOutput)(nil), // 1: eolymp.feed.ListEntriesOutput
	(*Entry)(nil),             // 2: eolymp.feed.Entry
}
var file_eolymp_feed_feed_service_proto_depIdxs = []int32{
	2, // 0: eolymp.feed.ListEntriesOutput.items:type_name -> eolymp.feed.Entry
	0, // 1: eolymp.feed.FeedService.ListEntries:input_type -> eolymp.feed.ListEntriesInput
	1, // 2: eolymp.feed.FeedService.ListEntries:output_type -> eolymp.feed.ListEntriesOutput
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_eolymp_feed_feed_service_proto_init() }
func file_eolymp_feed_feed_service_proto_init() {
	if File_eolymp_feed_feed_service_proto != nil {
		return
	}
	file_eolymp_feed_entry_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_eolymp_feed_feed_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListEntriesInput); i {
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
		file_eolymp_feed_feed_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListEntriesOutput); i {
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
			RawDescriptor: file_eolymp_feed_feed_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_eolymp_feed_feed_service_proto_goTypes,
		DependencyIndexes: file_eolymp_feed_feed_service_proto_depIdxs,
		MessageInfos:      file_eolymp_feed_feed_service_proto_msgTypes,
	}.Build()
	File_eolymp_feed_feed_service_proto = out.File
	file_eolymp_feed_feed_service_proto_rawDesc = nil
	file_eolymp_feed_feed_service_proto_goTypes = nil
	file_eolymp_feed_feed_service_proto_depIdxs = nil
}
