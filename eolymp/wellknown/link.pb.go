// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v4.24.4
// source: eolymp/wellknown/link.proto

package wellknown

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

type Link struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Rel           string                 `protobuf:"bytes,1,opt,name=rel,proto3" json:"rel,omitempty"`   // relation to the object
	Ref           string                 `protobuf:"bytes,2,opt,name=ref,proto3" json:"ref,omitempty"`   // the location of linked object
	Type          string                 `protobuf:"bytes,3,opt,name=type,proto3" json:"type,omitempty"` // type of the linked object
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Link) Reset() {
	*x = Link{}
	mi := &file_eolymp_wellknown_link_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Link) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Link) ProtoMessage() {}

func (x *Link) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_wellknown_link_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Link.ProtoReflect.Descriptor instead.
func (*Link) Descriptor() ([]byte, []int) {
	return file_eolymp_wellknown_link_proto_rawDescGZIP(), []int{0}
}

func (x *Link) GetRel() string {
	if x != nil {
		return x.Rel
	}
	return ""
}

func (x *Link) GetRef() string {
	if x != nil {
		return x.Ref
	}
	return ""
}

func (x *Link) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

var File_eolymp_wellknown_link_proto protoreflect.FileDescriptor

const file_eolymp_wellknown_link_proto_rawDesc = "" +
	"\n" +
	"\x1beolymp/wellknown/link.proto\x12\x10eolymp.wellknown\">\n" +
	"\x04Link\x12\x10\n" +
	"\x03rel\x18\x01 \x01(\tR\x03rel\x12\x10\n" +
	"\x03ref\x18\x02 \x01(\tR\x03ref\x12\x12\n" +
	"\x04type\x18\x03 \x01(\tR\x04typeB5Z3github.com/eolymp/go-sdk/eolymp/wellknown;wellknownb\x06proto3"

var (
	file_eolymp_wellknown_link_proto_rawDescOnce sync.Once
	file_eolymp_wellknown_link_proto_rawDescData []byte
)

func file_eolymp_wellknown_link_proto_rawDescGZIP() []byte {
	file_eolymp_wellknown_link_proto_rawDescOnce.Do(func() {
		file_eolymp_wellknown_link_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_eolymp_wellknown_link_proto_rawDesc), len(file_eolymp_wellknown_link_proto_rawDesc)))
	})
	return file_eolymp_wellknown_link_proto_rawDescData
}

var file_eolymp_wellknown_link_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_eolymp_wellknown_link_proto_goTypes = []any{
	(*Link)(nil), // 0: eolymp.wellknown.Link
}
var file_eolymp_wellknown_link_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_eolymp_wellknown_link_proto_init() }
func file_eolymp_wellknown_link_proto_init() {
	if File_eolymp_wellknown_link_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_eolymp_wellknown_link_proto_rawDesc), len(file_eolymp_wellknown_link_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_eolymp_wellknown_link_proto_goTypes,
		DependencyIndexes: file_eolymp_wellknown_link_proto_depIdxs,
		MessageInfos:      file_eolymp_wellknown_link_proto_msgTypes,
	}.Build()
	File_eolymp_wellknown_link_proto = out.File
	file_eolymp_wellknown_link_proto_goTypes = nil
	file_eolymp_wellknown_link_proto_depIdxs = nil
}
