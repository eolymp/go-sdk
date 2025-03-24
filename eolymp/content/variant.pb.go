// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v4.24.4
// source: eolymp/content/variant.proto

package content

import (
	ecm "github.com/eolymp/go-sdk/eolymp/ecm"
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

type Variant_Extra int32

const (
	Variant_NO_EXTRA       Variant_Extra = 0
	Variant_CONTENT_RENDER Variant_Extra = 1
	Variant_CONTENT_VALUE  Variant_Extra = 2
)

// Enum value maps for Variant_Extra.
var (
	Variant_Extra_name = map[int32]string{
		0: "NO_EXTRA",
		1: "CONTENT_RENDER",
		2: "CONTENT_VALUE",
	}
	Variant_Extra_value = map[string]int32{
		"NO_EXTRA":       0,
		"CONTENT_RENDER": 1,
		"CONTENT_VALUE":  2,
	}
)

func (x Variant_Extra) Enum() *Variant_Extra {
	p := new(Variant_Extra)
	*p = x
	return p
}

func (x Variant_Extra) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Variant_Extra) Descriptor() protoreflect.EnumDescriptor {
	return file_eolymp_content_variant_proto_enumTypes[0].Descriptor()
}

func (Variant_Extra) Type() protoreflect.EnumType {
	return &file_eolymp_content_variant_proto_enumTypes[0]
}

func (x Variant_Extra) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Variant_Extra.Descriptor instead.
func (Variant_Extra) EnumDescriptor() ([]byte, []int) {
	return file_eolymp_content_variant_proto_rawDescGZIP(), []int{0, 0}
}

type Variant struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Locale        string                 `protobuf:"bytes,11,opt,name=locale,proto3" json:"locale,omitempty"`
	Title         string                 `protobuf:"bytes,12,opt,name=title,proto3" json:"title,omitempty"`
	Content       *ecm.Content           `protobuf:"bytes,51,opt,name=content,proto3" json:"content,omitempty"`
	Automatic     bool                   `protobuf:"varint,8,opt,name=automatic,proto3" json:"automatic,omitempty"` // if true means the variant was created automatically (probably translation from another language)
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Variant) Reset() {
	*x = Variant{}
	mi := &file_eolymp_content_variant_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Variant) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Variant) ProtoMessage() {}

func (x *Variant) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_content_variant_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Variant.ProtoReflect.Descriptor instead.
func (*Variant) Descriptor() ([]byte, []int) {
	return file_eolymp_content_variant_proto_rawDescGZIP(), []int{0}
}

func (x *Variant) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Variant) GetLocale() string {
	if x != nil {
		return x.Locale
	}
	return ""
}

func (x *Variant) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Variant) GetContent() *ecm.Content {
	if x != nil {
		return x.Content
	}
	return nil
}

func (x *Variant) GetAutomatic() bool {
	if x != nil {
		return x.Automatic
	}
	return false
}

var File_eolymp_content_variant_proto protoreflect.FileDescriptor

const file_eolymp_content_variant_proto_rawDesc = "" +
	"\n" +
	"\x1ceolymp/content/variant.proto\x12\x0eeolymp.content\x1a\x18eolymp/ecm/content.proto\"\xd2\x01\n" +
	"\aVariant\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x12\x16\n" +
	"\x06locale\x18\v \x01(\tR\x06locale\x12\x14\n" +
	"\x05title\x18\f \x01(\tR\x05title\x12-\n" +
	"\acontent\x183 \x01(\v2\x13.eolymp.ecm.ContentR\acontent\x12\x1c\n" +
	"\tautomatic\x18\b \x01(\bR\tautomatic\"<\n" +
	"\x05Extra\x12\f\n" +
	"\bNO_EXTRA\x10\x00\x12\x12\n" +
	"\x0eCONTENT_RENDER\x10\x01\x12\x11\n" +
	"\rCONTENT_VALUE\x10\x02B1Z/github.com/eolymp/go-sdk/eolymp/content;contentb\x06proto3"

var (
	file_eolymp_content_variant_proto_rawDescOnce sync.Once
	file_eolymp_content_variant_proto_rawDescData []byte
)

func file_eolymp_content_variant_proto_rawDescGZIP() []byte {
	file_eolymp_content_variant_proto_rawDescOnce.Do(func() {
		file_eolymp_content_variant_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_eolymp_content_variant_proto_rawDesc), len(file_eolymp_content_variant_proto_rawDesc)))
	})
	return file_eolymp_content_variant_proto_rawDescData
}

var file_eolymp_content_variant_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_eolymp_content_variant_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_eolymp_content_variant_proto_goTypes = []any{
	(Variant_Extra)(0),  // 0: eolymp.content.Variant.Extra
	(*Variant)(nil),     // 1: eolymp.content.Variant
	(*ecm.Content)(nil), // 2: eolymp.ecm.Content
}
var file_eolymp_content_variant_proto_depIdxs = []int32{
	2, // 0: eolymp.content.Variant.content:type_name -> eolymp.ecm.Content
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_eolymp_content_variant_proto_init() }
func file_eolymp_content_variant_proto_init() {
	if File_eolymp_content_variant_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_eolymp_content_variant_proto_rawDesc), len(file_eolymp_content_variant_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_eolymp_content_variant_proto_goTypes,
		DependencyIndexes: file_eolymp_content_variant_proto_depIdxs,
		EnumInfos:         file_eolymp_content_variant_proto_enumTypes,
		MessageInfos:      file_eolymp_content_variant_proto_msgTypes,
	}.Build()
	File_eolymp_content_variant_proto = out.File
	file_eolymp_content_variant_proto_goTypes = nil
	file_eolymp_content_variant_proto_depIdxs = nil
}
