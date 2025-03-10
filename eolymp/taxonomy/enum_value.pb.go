// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v4.24.4
// source: eolymp/taxonomy/enum_value.proto

package taxonomy

import (
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

type Value struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`         // entry name
	Summary       string                 `protobuf:"bytes,3,opt,name=summary,proto3" json:"summary,omitempty"`   // entry summary (short description)
	Abbr          string                 `protobuf:"bytes,5,opt,name=abbr,proto3" json:"abbr,omitempty"`         // abbreviation
	Image         string                 `protobuf:"bytes,6,opt,name=image,proto3" json:"image,omitempty"`       // icon
	Keywords      []string               `protobuf:"bytes,4,rep,name=keywords,proto3" json:"keywords,omitempty"` // alternative names and keywords related to the entry (used for search)
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Value) Reset() {
	*x = Value{}
	mi := &file_eolymp_taxonomy_enum_value_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Value) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Value) ProtoMessage() {}

func (x *Value) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_taxonomy_enum_value_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Value.ProtoReflect.Descriptor instead.
func (*Value) Descriptor() ([]byte, []int) {
	return file_eolymp_taxonomy_enum_value_proto_rawDescGZIP(), []int{0}
}

func (x *Value) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Value) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Value) GetSummary() string {
	if x != nil {
		return x.Summary
	}
	return ""
}

func (x *Value) GetAbbr() string {
	if x != nil {
		return x.Abbr
	}
	return ""
}

func (x *Value) GetImage() string {
	if x != nil {
		return x.Image
	}
	return ""
}

func (x *Value) GetKeywords() []string {
	if x != nil {
		return x.Keywords
	}
	return nil
}

type Value_Translation struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Locale        string                 `protobuf:"bytes,1,opt,name=locale,proto3" json:"locale,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`         // entry name
	Abbr          string                 `protobuf:"bytes,5,opt,name=abbr,proto3" json:"abbr,omitempty"`         // abbreviation
	Summary       string                 `protobuf:"bytes,3,opt,name=summary,proto3" json:"summary,omitempty"`   // topic summary (short description)
	Keywords      []string               `protobuf:"bytes,4,rep,name=keywords,proto3" json:"keywords,omitempty"` // alternative names and keywords related to the entry (used for search)
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Value_Translation) Reset() {
	*x = Value_Translation{}
	mi := &file_eolymp_taxonomy_enum_value_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Value_Translation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Value_Translation) ProtoMessage() {}

func (x *Value_Translation) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_taxonomy_enum_value_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Value_Translation.ProtoReflect.Descriptor instead.
func (*Value_Translation) Descriptor() ([]byte, []int) {
	return file_eolymp_taxonomy_enum_value_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Value_Translation) GetLocale() string {
	if x != nil {
		return x.Locale
	}
	return ""
}

func (x *Value_Translation) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Value_Translation) GetAbbr() string {
	if x != nil {
		return x.Abbr
	}
	return ""
}

func (x *Value_Translation) GetSummary() string {
	if x != nil {
		return x.Summary
	}
	return ""
}

func (x *Value_Translation) GetKeywords() []string {
	if x != nil {
		return x.Keywords
	}
	return nil
}

var File_eolymp_taxonomy_enum_value_proto protoreflect.FileDescriptor

var file_eolymp_taxonomy_enum_value_proto_rawDesc = string([]byte{
	0x0a, 0x20, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x74, 0x61, 0x78, 0x6f, 0x6e, 0x6f, 0x6d,
	0x79, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0f, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x74, 0x61, 0x78, 0x6f, 0x6e,
	0x6f, 0x6d, 0x79, 0x22, 0x91, 0x02, 0x0a, 0x05, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x61,
	0x62, 0x62, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x61, 0x62, 0x62, 0x72, 0x12,
	0x14, 0x0a, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x69, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64,
	0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64,
	0x73, 0x1a, 0x83, 0x01, 0x0a, 0x0b, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x6c, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x16, 0x0a, 0x06, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x61, 0x62, 0x62, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x61, 0x62, 0x62,
	0x72, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x6b,
	0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x6b,
	0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x42, 0x33, 0x5a, 0x31, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x67, 0x6f, 0x2d,
	0x73, 0x64, 0x6b, 0x2f, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x74, 0x61, 0x78, 0x6f, 0x6e,
	0x6f, 0x6d, 0x79, 0x3b, 0x74, 0x61, 0x78, 0x6f, 0x6e, 0x6f, 0x6d, 0x79, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_eolymp_taxonomy_enum_value_proto_rawDescOnce sync.Once
	file_eolymp_taxonomy_enum_value_proto_rawDescData []byte
)

func file_eolymp_taxonomy_enum_value_proto_rawDescGZIP() []byte {
	file_eolymp_taxonomy_enum_value_proto_rawDescOnce.Do(func() {
		file_eolymp_taxonomy_enum_value_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_eolymp_taxonomy_enum_value_proto_rawDesc), len(file_eolymp_taxonomy_enum_value_proto_rawDesc)))
	})
	return file_eolymp_taxonomy_enum_value_proto_rawDescData
}

var file_eolymp_taxonomy_enum_value_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_eolymp_taxonomy_enum_value_proto_goTypes = []any{
	(*Value)(nil),             // 0: eolymp.taxonomy.Value
	(*Value_Translation)(nil), // 1: eolymp.taxonomy.Value.Translation
}
var file_eolymp_taxonomy_enum_value_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_eolymp_taxonomy_enum_value_proto_init() }
func file_eolymp_taxonomy_enum_value_proto_init() {
	if File_eolymp_taxonomy_enum_value_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_eolymp_taxonomy_enum_value_proto_rawDesc), len(file_eolymp_taxonomy_enum_value_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_eolymp_taxonomy_enum_value_proto_goTypes,
		DependencyIndexes: file_eolymp_taxonomy_enum_value_proto_depIdxs,
		MessageInfos:      file_eolymp_taxonomy_enum_value_proto_msgTypes,
	}.Build()
	File_eolymp_taxonomy_enum_value_proto = out.File
	file_eolymp_taxonomy_enum_value_proto_goTypes = nil
	file_eolymp_taxonomy_enum_value_proto_depIdxs = nil
}
