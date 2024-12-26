// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.1
// 	protoc        v4.24.4
// source: eolymp/discussion/post_type.proto

package discussion

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

type PostType_Extra int32

const (
	PostType_UNKNOWN_EXTRA PostType_Extra = 0
	PostType_VARIANTS      PostType_Extra = 1 // include post name translations
)

// Enum value maps for PostType_Extra.
var (
	PostType_Extra_name = map[int32]string{
		0: "UNKNOWN_EXTRA",
		1: "VARIANTS",
	}
	PostType_Extra_value = map[string]int32{
		"UNKNOWN_EXTRA": 0,
		"VARIANTS":      1,
	}
)

func (x PostType_Extra) Enum() *PostType_Extra {
	p := new(PostType_Extra)
	*p = x
	return p
}

func (x PostType_Extra) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PostType_Extra) Descriptor() protoreflect.EnumDescriptor {
	return file_eolymp_discussion_post_type_proto_enumTypes[0].Descriptor()
}

func (PostType_Extra) Type() protoreflect.EnumType {
	return &file_eolymp_discussion_post_type_proto_enumTypes[0]
}

func (x PostType_Extra) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PostType_Extra.Descriptor instead.
func (PostType_Extra) EnumDescriptor() ([]byte, []int) {
	return file_eolymp_discussion_post_type_proto_rawDescGZIP(), []int{0, 0}
}

type PostType struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Hidden        bool                   `protobuf:"varint,3,opt,name=hidden,proto3" json:"hidden,omitempty"`
	Order         int32                  `protobuf:"varint,10,opt,name=order,proto3" json:"order,omitempty"`
	Variants      []*PostType_Variant    `protobuf:"bytes,100,rep,name=variants,proto3" json:"variants,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PostType) Reset() {
	*x = PostType{}
	mi := &file_eolymp_discussion_post_type_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PostType) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostType) ProtoMessage() {}

func (x *PostType) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_discussion_post_type_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PostType.ProtoReflect.Descriptor instead.
func (*PostType) Descriptor() ([]byte, []int) {
	return file_eolymp_discussion_post_type_proto_rawDescGZIP(), []int{0}
}

func (x *PostType) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *PostType) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *PostType) GetHidden() bool {
	if x != nil {
		return x.Hidden
	}
	return false
}

func (x *PostType) GetOrder() int32 {
	if x != nil {
		return x.Order
	}
	return 0
}

func (x *PostType) GetVariants() []*PostType_Variant {
	if x != nil {
		return x.Variants
	}
	return nil
}

type PostType_Variant struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Locale        string                 `protobuf:"bytes,1,opt,name=locale,proto3" json:"locale,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PostType_Variant) Reset() {
	*x = PostType_Variant{}
	mi := &file_eolymp_discussion_post_type_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PostType_Variant) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostType_Variant) ProtoMessage() {}

func (x *PostType_Variant) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_discussion_post_type_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PostType_Variant.ProtoReflect.Descriptor instead.
func (*PostType_Variant) Descriptor() ([]byte, []int) {
	return file_eolymp_discussion_post_type_proto_rawDescGZIP(), []int{0, 0}
}

func (x *PostType_Variant) GetLocale() string {
	if x != nil {
		return x.Locale
	}
	return ""
}

func (x *PostType_Variant) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_eolymp_discussion_post_type_proto protoreflect.FileDescriptor

var file_eolymp_discussion_post_type_proto_rawDesc = []byte{
	0x0a, 0x21, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x64, 0x69, 0x73, 0x63, 0x75, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x2f, 0x70, 0x6f, 0x73, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x11, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x64, 0x69, 0x73, 0x63,
	0x75, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0xfe, 0x01, 0x0a, 0x08, 0x50, 0x6f, 0x73, 0x74, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x68, 0x69, 0x64, 0x64, 0x65,
	0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x68, 0x69, 0x64, 0x64, 0x65, 0x6e, 0x12,
	0x14, 0x0a, 0x05, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x3f, 0x0a, 0x08, 0x76, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x74,
	0x73, 0x18, 0x64, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70,
	0x2e, 0x64, 0x69, 0x73, 0x63, 0x75, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x50, 0x6f, 0x73, 0x74,
	0x54, 0x79, 0x70, 0x65, 0x2e, 0x56, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x74, 0x52, 0x08, 0x76, 0x61,
	0x72, 0x69, 0x61, 0x6e, 0x74, 0x73, 0x1a, 0x35, 0x0a, 0x07, 0x56, 0x61, 0x72, 0x69, 0x61, 0x6e,
	0x74, 0x12, 0x16, 0x0a, 0x06, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x28, 0x0a,
	0x05, 0x45, 0x78, 0x74, 0x72, 0x61, 0x12, 0x11, 0x0a, 0x0d, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57,
	0x4e, 0x5f, 0x45, 0x58, 0x54, 0x52, 0x41, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x56, 0x41, 0x52,
	0x49, 0x41, 0x4e, 0x54, 0x53, 0x10, 0x01, 0x42, 0x37, 0x5a, 0x35, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x67, 0x6f, 0x2d,
	0x73, 0x64, 0x6b, 0x2f, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x64, 0x69, 0x73, 0x63, 0x75,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x3b, 0x64, 0x69, 0x73, 0x63, 0x75, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_eolymp_discussion_post_type_proto_rawDescOnce sync.Once
	file_eolymp_discussion_post_type_proto_rawDescData = file_eolymp_discussion_post_type_proto_rawDesc
)

func file_eolymp_discussion_post_type_proto_rawDescGZIP() []byte {
	file_eolymp_discussion_post_type_proto_rawDescOnce.Do(func() {
		file_eolymp_discussion_post_type_proto_rawDescData = protoimpl.X.CompressGZIP(file_eolymp_discussion_post_type_proto_rawDescData)
	})
	return file_eolymp_discussion_post_type_proto_rawDescData
}

var file_eolymp_discussion_post_type_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_eolymp_discussion_post_type_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_eolymp_discussion_post_type_proto_goTypes = []any{
	(PostType_Extra)(0),      // 0: eolymp.discussion.PostType.Extra
	(*PostType)(nil),         // 1: eolymp.discussion.PostType
	(*PostType_Variant)(nil), // 2: eolymp.discussion.PostType.Variant
}
var file_eolymp_discussion_post_type_proto_depIdxs = []int32{
	2, // 0: eolymp.discussion.PostType.variants:type_name -> eolymp.discussion.PostType.Variant
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_eolymp_discussion_post_type_proto_init() }
func file_eolymp_discussion_post_type_proto_init() {
	if File_eolymp_discussion_post_type_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_eolymp_discussion_post_type_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_eolymp_discussion_post_type_proto_goTypes,
		DependencyIndexes: file_eolymp_discussion_post_type_proto_depIdxs,
		EnumInfos:         file_eolymp_discussion_post_type_proto_enumTypes,
		MessageInfos:      file_eolymp_discussion_post_type_proto_msgTypes,
	}.Build()
	File_eolymp_discussion_post_type_proto = out.File
	file_eolymp_discussion_post_type_proto_rawDesc = nil
	file_eolymp_discussion_post_type_proto_goTypes = nil
	file_eolymp_discussion_post_type_proto_depIdxs = nil
}
