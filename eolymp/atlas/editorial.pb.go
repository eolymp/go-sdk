// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.18.1
// source: eolymp/atlas/editorial.proto

package atlas

import (
	ecm "github.com/eolymp/go-sdk/eolymp/ecm"
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

type Editorial struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Unique identifier, assigned when statement is created. Keep empty when creating new statement.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Problem where this statement belongs. Keep empty when creating new statement.
	ProblemId string `protobuf:"bytes,2,opt,name=problem_id,json=problemId,proto3" json:"problem_id,omitempty"`
	// Statement locale code, should consist of two lowercase latin letters.
	Locale string `protobuf:"bytes,3,opt,name=locale,proto3" json:"locale,omitempty"`
	// Statement content.
	//
	// Types that are assignable to Content:
	//	*Editorial_ContentHtml
	//	*Editorial_ContentLatex
	//	*Editorial_ContentMarkdown
	//	*Editorial_ContentEcm
	Content isEditorial_Content `protobuf_oneof:"content"`
	// Statement download link, allows to download statement in original format.
	DownloadLink string `protobuf:"bytes,7,opt,name=download_link,json=downloadLink,proto3" json:"download_link,omitempty"`
}

func (x *Editorial) Reset() {
	*x = Editorial{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_atlas_editorial_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Editorial) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Editorial) ProtoMessage() {}

func (x *Editorial) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_atlas_editorial_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Editorial.ProtoReflect.Descriptor instead.
func (*Editorial) Descriptor() ([]byte, []int) {
	return file_eolymp_atlas_editorial_proto_rawDescGZIP(), []int{0}
}

func (x *Editorial) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Editorial) GetProblemId() string {
	if x != nil {
		return x.ProblemId
	}
	return ""
}

func (x *Editorial) GetLocale() string {
	if x != nil {
		return x.Locale
	}
	return ""
}

func (m *Editorial) GetContent() isEditorial_Content {
	if m != nil {
		return m.Content
	}
	return nil
}

func (x *Editorial) GetContentHtml() string {
	if x, ok := x.GetContent().(*Editorial_ContentHtml); ok {
		return x.ContentHtml
	}
	return ""
}

func (x *Editorial) GetContentLatex() string {
	if x, ok := x.GetContent().(*Editorial_ContentLatex); ok {
		return x.ContentLatex
	}
	return ""
}

func (x *Editorial) GetContentMarkdown() string {
	if x, ok := x.GetContent().(*Editorial_ContentMarkdown); ok {
		return x.ContentMarkdown
	}
	return ""
}

func (x *Editorial) GetContentEcm() *ecm.Node {
	if x, ok := x.GetContent().(*Editorial_ContentEcm); ok {
		return x.ContentEcm
	}
	return nil
}

func (x *Editorial) GetDownloadLink() string {
	if x != nil {
		return x.DownloadLink
	}
	return ""
}

type isEditorial_Content interface {
	isEditorial_Content()
}

type Editorial_ContentHtml struct {
	ContentHtml string `protobuf:"bytes,50,opt,name=content_html,json=contentHtml,proto3,oneof"`
}

type Editorial_ContentLatex struct {
	ContentLatex string `protobuf:"bytes,51,opt,name=content_latex,json=contentLatex,proto3,oneof"`
}

type Editorial_ContentMarkdown struct {
	ContentMarkdown string `protobuf:"bytes,52,opt,name=content_markdown,json=contentMarkdown,proto3,oneof"`
}

type Editorial_ContentEcm struct {
	ContentEcm *ecm.Node `protobuf:"bytes,53,opt,name=content_ecm,json=contentEcm,proto3,oneof"`
}

func (*Editorial_ContentHtml) isEditorial_Content() {}

func (*Editorial_ContentLatex) isEditorial_Content() {}

func (*Editorial_ContentMarkdown) isEditorial_Content() {}

func (*Editorial_ContentEcm) isEditorial_Content() {}

var File_eolymp_atlas_editorial_proto protoreflect.FileDescriptor

var file_eolymp_atlas_editorial_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x61, 0x74, 0x6c, 0x61, 0x73, 0x2f, 0x65,
	0x64, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c,
	0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x61, 0x74, 0x6c, 0x61, 0x73, 0x1a, 0x15, 0x65, 0x6f,
	0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x65, 0x63, 0x6d, 0x2f, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xb0, 0x02, 0x0a, 0x09, 0x45, 0x64, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x61,
	0x6c, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x5f, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x49, 0x64,
	0x12, 0x16, 0x0a, 0x06, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x65, 0x12, 0x23, 0x0a, 0x0c, 0x63, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x5f, 0x68, 0x74, 0x6d, 0x6c, 0x18, 0x32, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00,
	0x52, 0x0b, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x48, 0x74, 0x6d, 0x6c, 0x12, 0x25, 0x0a,
	0x0d, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x5f, 0x6c, 0x61, 0x74, 0x65, 0x78, 0x18, 0x33,
	0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0c, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x4c,
	0x61, 0x74, 0x65, 0x78, 0x12, 0x2b, 0x0a, 0x10, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x5f,
	0x6d, 0x61, 0x72, 0x6b, 0x64, 0x6f, 0x77, 0x6e, 0x18, 0x34, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00,
	0x52, 0x0f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x4d, 0x61, 0x72, 0x6b, 0x64, 0x6f, 0x77,
	0x6e, 0x12, 0x33, 0x0a, 0x0b, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x5f, 0x65, 0x63, 0x6d,
	0x18, 0x35, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e,
	0x65, 0x63, 0x6d, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x48, 0x00, 0x52, 0x0a, 0x63, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x45, 0x63, 0x6d, 0x12, 0x23, 0x0a, 0x0d, 0x64, 0x6f, 0x77, 0x6e, 0x6c, 0x6f,
	0x61, 0x64, 0x5f, 0x6c, 0x69, 0x6e, 0x6b, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x64,
	0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x4c, 0x69, 0x6e, 0x6b, 0x42, 0x09, 0x0a, 0x07, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x42, 0x2d, 0x5a, 0x2b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x67, 0x6f, 0x2d, 0x73,
	0x64, 0x6b, 0x2f, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x61, 0x74, 0x6c, 0x61, 0x73, 0x3b,
	0x61, 0x74, 0x6c, 0x61, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_eolymp_atlas_editorial_proto_rawDescOnce sync.Once
	file_eolymp_atlas_editorial_proto_rawDescData = file_eolymp_atlas_editorial_proto_rawDesc
)

func file_eolymp_atlas_editorial_proto_rawDescGZIP() []byte {
	file_eolymp_atlas_editorial_proto_rawDescOnce.Do(func() {
		file_eolymp_atlas_editorial_proto_rawDescData = protoimpl.X.CompressGZIP(file_eolymp_atlas_editorial_proto_rawDescData)
	})
	return file_eolymp_atlas_editorial_proto_rawDescData
}

var file_eolymp_atlas_editorial_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_eolymp_atlas_editorial_proto_goTypes = []interface{}{
	(*Editorial)(nil), // 0: eolymp.atlas.Editorial
	(*ecm.Node)(nil),  // 1: eolymp.ecm.Node
}
var file_eolymp_atlas_editorial_proto_depIdxs = []int32{
	1, // 0: eolymp.atlas.Editorial.content_ecm:type_name -> eolymp.ecm.Node
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_eolymp_atlas_editorial_proto_init() }
func file_eolymp_atlas_editorial_proto_init() {
	if File_eolymp_atlas_editorial_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_eolymp_atlas_editorial_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Editorial); i {
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
	file_eolymp_atlas_editorial_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*Editorial_ContentHtml)(nil),
		(*Editorial_ContentLatex)(nil),
		(*Editorial_ContentMarkdown)(nil),
		(*Editorial_ContentEcm)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_eolymp_atlas_editorial_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_eolymp_atlas_editorial_proto_goTypes,
		DependencyIndexes: file_eolymp_atlas_editorial_proto_depIdxs,
		MessageInfos:      file_eolymp_atlas_editorial_proto_msgTypes,
	}.Build()
	File_eolymp_atlas_editorial_proto = out.File
	file_eolymp_atlas_editorial_proto_rawDesc = nil
	file_eolymp_atlas_editorial_proto_goTypes = nil
	file_eolymp_atlas_editorial_proto_depIdxs = nil
}