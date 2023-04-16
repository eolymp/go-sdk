// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.18.1
// source: eolymp/taxonomy/topic.proto

package taxonomy

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

type Topic struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name     string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`         // topic name
	Summary  string   `protobuf:"bytes,3,opt,name=summary,proto3" json:"summary,omitempty"`   // topic summary (short description)
	Keywords []string `protobuf:"bytes,4,rep,name=keywords,proto3" json:"keywords,omitempty"` // alternative names and keywords related to the topic (used for search)
}

func (x *Topic) Reset() {
	*x = Topic{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_taxonomy_topic_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Topic) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Topic) ProtoMessage() {}

func (x *Topic) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_taxonomy_topic_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Topic.ProtoReflect.Descriptor instead.
func (*Topic) Descriptor() ([]byte, []int) {
	return file_eolymp_taxonomy_topic_proto_rawDescGZIP(), []int{0}
}

func (x *Topic) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Topic) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Topic) GetSummary() string {
	if x != nil {
		return x.Summary
	}
	return ""
}

func (x *Topic) GetKeywords() []string {
	if x != nil {
		return x.Keywords
	}
	return nil
}

type Topic_Translation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Locale   string   `protobuf:"bytes,1,opt,name=locale,proto3" json:"locale,omitempty"`
	Name     string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`         // topic name
	Summary  string   `protobuf:"bytes,3,opt,name=summary,proto3" json:"summary,omitempty"`   // topic summary (short description)
	Keywords []string `protobuf:"bytes,4,rep,name=keywords,proto3" json:"keywords,omitempty"` // alternative names and keywords related to the topic (used for search)
}

func (x *Topic_Translation) Reset() {
	*x = Topic_Translation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_taxonomy_topic_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Topic_Translation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Topic_Translation) ProtoMessage() {}

func (x *Topic_Translation) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_taxonomy_topic_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Topic_Translation.ProtoReflect.Descriptor instead.
func (*Topic_Translation) Descriptor() ([]byte, []int) {
	return file_eolymp_taxonomy_topic_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Topic_Translation) GetLocale() string {
	if x != nil {
		return x.Locale
	}
	return ""
}

func (x *Topic_Translation) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Topic_Translation) GetSummary() string {
	if x != nil {
		return x.Summary
	}
	return ""
}

func (x *Topic_Translation) GetKeywords() []string {
	if x != nil {
		return x.Keywords
	}
	return nil
}

var File_eolymp_taxonomy_topic_proto protoreflect.FileDescriptor

var file_eolymp_taxonomy_topic_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x74, 0x61, 0x78, 0x6f, 0x6e, 0x6f, 0x6d,
	0x79, 0x2f, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x65,
	0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x74, 0x61, 0x78, 0x6f, 0x6e, 0x6f, 0x6d, 0x79, 0x22, 0xd2,
	0x01, 0x0a, 0x05, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73,
	0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72,
	0x64, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72,
	0x64, 0x73, 0x1a, 0x6f, 0x0a, 0x0b, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x6c, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x16, 0x0a, 0x06, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x6b, 0x65, 0x79, 0x77, 0x6f,
	0x72, 0x64, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x6b, 0x65, 0x79, 0x77, 0x6f,
	0x72, 0x64, 0x73, 0x42, 0x33, 0x5a, 0x31, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x67, 0x6f, 0x2d, 0x73, 0x64, 0x6b, 0x2f,
	0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x74, 0x61, 0x78, 0x6f, 0x6e, 0x6f, 0x6d, 0x79, 0x3b,
	0x74, 0x61, 0x78, 0x6f, 0x6e, 0x6f, 0x6d, 0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_eolymp_taxonomy_topic_proto_rawDescOnce sync.Once
	file_eolymp_taxonomy_topic_proto_rawDescData = file_eolymp_taxonomy_topic_proto_rawDesc
)

func file_eolymp_taxonomy_topic_proto_rawDescGZIP() []byte {
	file_eolymp_taxonomy_topic_proto_rawDescOnce.Do(func() {
		file_eolymp_taxonomy_topic_proto_rawDescData = protoimpl.X.CompressGZIP(file_eolymp_taxonomy_topic_proto_rawDescData)
	})
	return file_eolymp_taxonomy_topic_proto_rawDescData
}

var file_eolymp_taxonomy_topic_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_eolymp_taxonomy_topic_proto_goTypes = []interface{}{
	(*Topic)(nil),             // 0: eolymp.taxonomy.Topic
	(*Topic_Translation)(nil), // 1: eolymp.taxonomy.Topic.Translation
}
var file_eolymp_taxonomy_topic_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_eolymp_taxonomy_topic_proto_init() }
func file_eolymp_taxonomy_topic_proto_init() {
	if File_eolymp_taxonomy_topic_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_eolymp_taxonomy_topic_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Topic); i {
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
		file_eolymp_taxonomy_topic_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Topic_Translation); i {
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
			RawDescriptor: file_eolymp_taxonomy_topic_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_eolymp_taxonomy_topic_proto_goTypes,
		DependencyIndexes: file_eolymp_taxonomy_topic_proto_depIdxs,
		MessageInfos:      file_eolymp_taxonomy_topic_proto_msgTypes,
	}.Build()
	File_eolymp_taxonomy_topic_proto = out.File
	file_eolymp_taxonomy_topic_proto_rawDesc = nil
	file_eolymp_taxonomy_topic_proto_goTypes = nil
	file_eolymp_taxonomy_topic_proto_depIdxs = nil
}
