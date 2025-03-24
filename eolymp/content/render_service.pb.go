// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v4.24.4
// source: eolymp/content/render_service.proto

package content

import (
	_ "github.com/eolymp/go-sdk/eolymp/annotations"
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

type RenderContentInput struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Content       *ecm.Content           `protobuf:"bytes,1,opt,name=content,proto3" json:"content,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RenderContentInput) Reset() {
	*x = RenderContentInput{}
	mi := &file_eolymp_content_render_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RenderContentInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RenderContentInput) ProtoMessage() {}

func (x *RenderContentInput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_content_render_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RenderContentInput.ProtoReflect.Descriptor instead.
func (*RenderContentInput) Descriptor() ([]byte, []int) {
	return file_eolymp_content_render_service_proto_rawDescGZIP(), []int{0}
}

func (x *RenderContentInput) GetContent() *ecm.Content {
	if x != nil {
		return x.Content
	}
	return nil
}

type RenderContentOutput struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Render        *ecm.Node              `protobuf:"bytes,1,opt,name=render,proto3" json:"render,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RenderContentOutput) Reset() {
	*x = RenderContentOutput{}
	mi := &file_eolymp_content_render_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RenderContentOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RenderContentOutput) ProtoMessage() {}

func (x *RenderContentOutput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_content_render_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RenderContentOutput.ProtoReflect.Descriptor instead.
func (*RenderContentOutput) Descriptor() ([]byte, []int) {
	return file_eolymp_content_render_service_proto_rawDescGZIP(), []int{1}
}

func (x *RenderContentOutput) GetRender() *ecm.Node {
	if x != nil {
		return x.Render
	}
	return nil
}

var File_eolymp_content_render_service_proto protoreflect.FileDescriptor

const file_eolymp_content_render_service_proto_rawDesc = "" +
	"\n" +
	"#eolymp/content/render_service.proto\x12\x0eeolymp.content\x1a\x1deolymp/annotations/http.proto\x1a\"eolymp/annotations/ratelimit.proto\x1a\x18eolymp/ecm/content.proto\x1a\x15eolymp/ecm/node.proto\"C\n" +
	"\x12RenderContentInput\x12-\n" +
	"\acontent\x18\x01 \x01(\v2\x13.eolymp.ecm.ContentR\acontent\"?\n" +
	"\x13RenderContentOutput\x12(\n" +
	"\x06render\x18\x01 \x01(\v2\x10.eolymp.ecm.NodeR\x06render2\x9f\x01\n" +
	"\rRenderService\x12\x8d\x01\n" +
	"\rRenderContent\x12\".eolymp.content.RenderContentInput\x1a#.eolymp.content.RenderContentOutput\"3\xea\xe2\n" +
	"\v\xf5\xe2\n" +
	"\x00\x00\xc8A\xf8\xe2\n" +
	"d\x82\xd3\xe4\x93\x02\x1eZ\v\"\t/renderer\"\x0f/content:renderB1Z/github.com/eolymp/go-sdk/eolymp/content;contentb\x06proto3"

var (
	file_eolymp_content_render_service_proto_rawDescOnce sync.Once
	file_eolymp_content_render_service_proto_rawDescData []byte
)

func file_eolymp_content_render_service_proto_rawDescGZIP() []byte {
	file_eolymp_content_render_service_proto_rawDescOnce.Do(func() {
		file_eolymp_content_render_service_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_eolymp_content_render_service_proto_rawDesc), len(file_eolymp_content_render_service_proto_rawDesc)))
	})
	return file_eolymp_content_render_service_proto_rawDescData
}

var file_eolymp_content_render_service_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_eolymp_content_render_service_proto_goTypes = []any{
	(*RenderContentInput)(nil),  // 0: eolymp.content.RenderContentInput
	(*RenderContentOutput)(nil), // 1: eolymp.content.RenderContentOutput
	(*ecm.Content)(nil),         // 2: eolymp.ecm.Content
	(*ecm.Node)(nil),            // 3: eolymp.ecm.Node
}
var file_eolymp_content_render_service_proto_depIdxs = []int32{
	2, // 0: eolymp.content.RenderContentInput.content:type_name -> eolymp.ecm.Content
	3, // 1: eolymp.content.RenderContentOutput.render:type_name -> eolymp.ecm.Node
	0, // 2: eolymp.content.RenderService.RenderContent:input_type -> eolymp.content.RenderContentInput
	1, // 3: eolymp.content.RenderService.RenderContent:output_type -> eolymp.content.RenderContentOutput
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_eolymp_content_render_service_proto_init() }
func file_eolymp_content_render_service_proto_init() {
	if File_eolymp_content_render_service_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_eolymp_content_render_service_proto_rawDesc), len(file_eolymp_content_render_service_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_eolymp_content_render_service_proto_goTypes,
		DependencyIndexes: file_eolymp_content_render_service_proto_depIdxs,
		MessageInfos:      file_eolymp_content_render_service_proto_msgTypes,
	}.Build()
	File_eolymp_content_render_service_proto = out.File
	file_eolymp_content_render_service_proto_goTypes = nil
	file_eolymp_content_render_service_proto_depIdxs = nil
}
