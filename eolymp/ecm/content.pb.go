// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v4.24.4
// source: eolymp/ecm/content.proto

package ecm

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

type Content struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Types that are valid to be assigned to Value:
	//
	//	*Content_Html
	//	*Content_Latex
	//	*Content_Markdown
	//	*Content_Ecm
	Value         isContent_Value `protobuf_oneof:"value"`
	Render        *Node           `protobuf:"bytes,99,opt,name=render,proto3" json:"render,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Content) Reset() {
	*x = Content{}
	mi := &file_eolymp_ecm_content_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Content) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Content) ProtoMessage() {}

func (x *Content) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_ecm_content_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Content.ProtoReflect.Descriptor instead.
func (*Content) Descriptor() ([]byte, []int) {
	return file_eolymp_ecm_content_proto_rawDescGZIP(), []int{0}
}

func (x *Content) GetValue() isContent_Value {
	if x != nil {
		return x.Value
	}
	return nil
}

func (x *Content) GetHtml() string {
	if x != nil {
		if x, ok := x.Value.(*Content_Html); ok {
			return x.Html
		}
	}
	return ""
}

func (x *Content) GetLatex() string {
	if x != nil {
		if x, ok := x.Value.(*Content_Latex); ok {
			return x.Latex
		}
	}
	return ""
}

func (x *Content) GetMarkdown() string {
	if x != nil {
		if x, ok := x.Value.(*Content_Markdown); ok {
			return x.Markdown
		}
	}
	return ""
}

func (x *Content) GetEcm() *Node {
	if x != nil {
		if x, ok := x.Value.(*Content_Ecm); ok {
			return x.Ecm
		}
	}
	return nil
}

func (x *Content) GetRender() *Node {
	if x != nil {
		return x.Render
	}
	return nil
}

type isContent_Value interface {
	isContent_Value()
}

type Content_Html struct {
	Html string `protobuf:"bytes,1,opt,name=html,proto3,oneof"`
}

type Content_Latex struct {
	Latex string `protobuf:"bytes,2,opt,name=latex,proto3,oneof"`
}

type Content_Markdown struct {
	Markdown string `protobuf:"bytes,3,opt,name=markdown,proto3,oneof"`
}

type Content_Ecm struct {
	Ecm *Node `protobuf:"bytes,4,opt,name=ecm,proto3,oneof"`
}

func (*Content_Html) isContent_Value() {}

func (*Content_Latex) isContent_Value() {}

func (*Content_Markdown) isContent_Value() {}

func (*Content_Ecm) isContent_Value() {}

var File_eolymp_ecm_content_proto protoreflect.FileDescriptor

var file_eolymp_ecm_content_proto_rawDesc = string([]byte{
	0x0a, 0x18, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x65, 0x63, 0x6d, 0x2f, 0x63, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x65, 0x6f, 0x6c, 0x79,
	0x6d, 0x70, 0x2e, 0x65, 0x63, 0x6d, 0x1a, 0x15, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x65,
	0x63, 0x6d, 0x2f, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xae, 0x01,
	0x0a, 0x07, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x14, 0x0a, 0x04, 0x68, 0x74, 0x6d,
	0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x04, 0x68, 0x74, 0x6d, 0x6c, 0x12,
	0x16, 0x0a, 0x05, 0x6c, 0x61, 0x74, 0x65, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00,
	0x52, 0x05, 0x6c, 0x61, 0x74, 0x65, 0x78, 0x12, 0x1c, 0x0a, 0x08, 0x6d, 0x61, 0x72, 0x6b, 0x64,
	0x6f, 0x77, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x08, 0x6d, 0x61, 0x72,
	0x6b, 0x64, 0x6f, 0x77, 0x6e, 0x12, 0x24, 0x0a, 0x03, 0x65, 0x63, 0x6d, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x10, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x65, 0x63, 0x6d, 0x2e,
	0x4e, 0x6f, 0x64, 0x65, 0x48, 0x00, 0x52, 0x03, 0x65, 0x63, 0x6d, 0x12, 0x28, 0x0a, 0x06, 0x72,
	0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x63, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x65, 0x6f,
	0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x65, 0x63, 0x6d, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x06, 0x72,
	0x65, 0x6e, 0x64, 0x65, 0x72, 0x42, 0x07, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x29,
	0x5a, 0x27, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6f, 0x6c,
	0x79, 0x6d, 0x70, 0x2f, 0x67, 0x6f, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x65, 0x6f, 0x6c, 0x79, 0x6d,
	0x70, 0x2f, 0x65, 0x63, 0x6d, 0x3b, 0x65, 0x63, 0x6d, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
})

var (
	file_eolymp_ecm_content_proto_rawDescOnce sync.Once
	file_eolymp_ecm_content_proto_rawDescData []byte
)

func file_eolymp_ecm_content_proto_rawDescGZIP() []byte {
	file_eolymp_ecm_content_proto_rawDescOnce.Do(func() {
		file_eolymp_ecm_content_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_eolymp_ecm_content_proto_rawDesc), len(file_eolymp_ecm_content_proto_rawDesc)))
	})
	return file_eolymp_ecm_content_proto_rawDescData
}

var file_eolymp_ecm_content_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_eolymp_ecm_content_proto_goTypes = []any{
	(*Content)(nil), // 0: eolymp.ecm.Content
	(*Node)(nil),    // 1: eolymp.ecm.Node
}
var file_eolymp_ecm_content_proto_depIdxs = []int32{
	1, // 0: eolymp.ecm.Content.ecm:type_name -> eolymp.ecm.Node
	1, // 1: eolymp.ecm.Content.render:type_name -> eolymp.ecm.Node
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_eolymp_ecm_content_proto_init() }
func file_eolymp_ecm_content_proto_init() {
	if File_eolymp_ecm_content_proto != nil {
		return
	}
	file_eolymp_ecm_node_proto_init()
	file_eolymp_ecm_content_proto_msgTypes[0].OneofWrappers = []any{
		(*Content_Html)(nil),
		(*Content_Latex)(nil),
		(*Content_Markdown)(nil),
		(*Content_Ecm)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_eolymp_ecm_content_proto_rawDesc), len(file_eolymp_ecm_content_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_eolymp_ecm_content_proto_goTypes,
		DependencyIndexes: file_eolymp_ecm_content_proto_depIdxs,
		MessageInfos:      file_eolymp_ecm_content_proto_msgTypes,
	}.Build()
	File_eolymp_ecm_content_proto = out.File
	file_eolymp_ecm_content_proto_goTypes = nil
	file_eolymp_ecm_content_proto_depIdxs = nil
}
