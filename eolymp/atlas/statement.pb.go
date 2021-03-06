// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.18.1
// source: eolymp/atlas/statement.proto

package atlas

import (
	_ "github.com/eolymp/go-sdk/eolymp/annotations"
	typewriter "github.com/eolymp/go-sdk/eolymp/typewriter"
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

type Statement_Format int32

const (
	Statement_NONE     Statement_Format = 0
	Statement_TEX      Statement_Format = 1
	Statement_MARKDOWN Statement_Format = 2
	Statement_HTML     Statement_Format = 3
	Statement_RICH     Statement_Format = 4
)

// Enum value maps for Statement_Format.
var (
	Statement_Format_name = map[int32]string{
		0: "NONE",
		1: "TEX",
		2: "MARKDOWN",
		3: "HTML",
		4: "RICH",
	}
	Statement_Format_value = map[string]int32{
		"NONE":     0,
		"TEX":      1,
		"MARKDOWN": 2,
		"HTML":     3,
		"RICH":     4,
	}
)

func (x Statement_Format) Enum() *Statement_Format {
	p := new(Statement_Format)
	*p = x
	return p
}

func (x Statement_Format) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Statement_Format) Descriptor() protoreflect.EnumDescriptor {
	return file_eolymp_atlas_statement_proto_enumTypes[0].Descriptor()
}

func (Statement_Format) Type() protoreflect.EnumType {
	return &file_eolymp_atlas_statement_proto_enumTypes[0]
}

func (x Statement_Format) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Statement_Format.Descriptor instead.
func (Statement_Format) EnumDescriptor() ([]byte, []int) {
	return file_eolymp_atlas_statement_proto_rawDescGZIP(), []int{0, 0}
}

type Statement struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Unique identifier, assigned when statement is created. Keep empty when creating new statement.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Problem where this statement belongs. Keep empty when creating new statement.
	ProblemId string `protobuf:"bytes,2,opt,name=problem_id,json=problemId,proto3" json:"problem_id,omitempty"`
	// Statement locale code, should consist of two lowercase latin letters.
	Locale string `protobuf:"bytes,3,opt,name=locale,proto3" json:"locale,omitempty"`
	// Statement title.
	Title string `protobuf:"bytes,4,opt,name=title,proto3" json:"title,omitempty"`
	// Statement content. Might be defined using LaTeX, Markdown or HTML formats, see format field.
	Content string `protobuf:"bytes,5,opt,name=content,proto3" json:"content,omitempty"`
	// Statement content (in rich format).
	ContentRich *typewriter.Container `protobuf:"bytes,8,opt,name=content_rich,json=contentRich,proto3" json:"content_rich,omitempty"`
	// Statement download link, allows to download statement in original format.
	DownloadLink string `protobuf:"bytes,7,opt,name=download_link,json=downloadLink,proto3" json:"download_link,omitempty"`
	// Statement content format. Preferred format is LaTeX, avoid using HTML format as it is deprecated.
	Format Statement_Format `protobuf:"varint,6,opt,name=format,proto3,enum=eolymp.atlas.Statement_Format" json:"format,omitempty"`
	// Problem author name.
	Author string `protobuf:"bytes,101,opt,name=author,proto3" json:"author,omitempty"`
	// Problem source, name of the contest or olympiad where this problem was initially published.
	Source string `protobuf:"bytes,102,opt,name=source,proto3" json:"source,omitempty"`
}

func (x *Statement) Reset() {
	*x = Statement{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_atlas_statement_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Statement) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Statement) ProtoMessage() {}

func (x *Statement) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_atlas_statement_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Statement.ProtoReflect.Descriptor instead.
func (*Statement) Descriptor() ([]byte, []int) {
	return file_eolymp_atlas_statement_proto_rawDescGZIP(), []int{0}
}

func (x *Statement) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Statement) GetProblemId() string {
	if x != nil {
		return x.ProblemId
	}
	return ""
}

func (x *Statement) GetLocale() string {
	if x != nil {
		return x.Locale
	}
	return ""
}

func (x *Statement) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Statement) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *Statement) GetContentRich() *typewriter.Container {
	if x != nil {
		return x.ContentRich
	}
	return nil
}

func (x *Statement) GetDownloadLink() string {
	if x != nil {
		return x.DownloadLink
	}
	return ""
}

func (x *Statement) GetFormat() Statement_Format {
	if x != nil {
		return x.Format
	}
	return Statement_NONE
}

func (x *Statement) GetAuthor() string {
	if x != nil {
		return x.Author
	}
	return ""
}

func (x *Statement) GetSource() string {
	if x != nil {
		return x.Source
	}
	return ""
}

var File_eolymp_atlas_statement_proto protoreflect.FileDescriptor

var file_eolymp_atlas_statement_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x61, 0x74, 0x6c, 0x61, 0x73, 0x2f, 0x73,
	0x74, 0x61, 0x74, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c,
	0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x61, 0x74, 0x6c, 0x61, 0x73, 0x1a, 0x21, 0x65, 0x6f,
	0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1d, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x77, 0x72, 0x69, 0x74,
	0x65, 0x72, 0x2f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xba,
	0x03, 0x0a, 0x09, 0x53, 0x74, 0x61, 0x74, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1d, 0x0a, 0x0a,
	0x70, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x70, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x6c,
	0x6f, 0x63, 0x61, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6c, 0x6f, 0x63,
	0x61, 0x6c, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x12, 0x3f, 0x0a, 0x0c, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x5f, 0x72,
	0x69, 0x63, 0x68, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x65, 0x6f, 0x6c, 0x79,
	0x6d, 0x70, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x77, 0x72, 0x69, 0x74, 0x65, 0x72, 0x2e, 0x43, 0x6f,
	0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x52, 0x0b, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x52, 0x69, 0x63, 0x68, 0x12, 0x23, 0x0a, 0x0d, 0x64, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64,
	0x5f, 0x6c, 0x69, 0x6e, 0x6b, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x64, 0x6f, 0x77,
	0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x4c, 0x69, 0x6e, 0x6b, 0x12, 0x36, 0x0a, 0x06, 0x66, 0x6f, 0x72,
	0x6d, 0x61, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1e, 0x2e, 0x65, 0x6f, 0x6c, 0x79,
	0x6d, 0x70, 0x2e, 0x61, 0x74, 0x6c, 0x61, 0x73, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x6d, 0x65,
	0x6e, 0x74, 0x2e, 0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x52, 0x06, 0x66, 0x6f, 0x72, 0x6d, 0x61,
	0x74, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x18, 0x65, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x18, 0x66, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x22, 0x3d, 0x0a, 0x06, 0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x12, 0x08, 0x0a, 0x04, 0x4e,
	0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x07, 0x0a, 0x03, 0x54, 0x45, 0x58, 0x10, 0x01, 0x12, 0x0c,
	0x0a, 0x08, 0x4d, 0x41, 0x52, 0x4b, 0x44, 0x4f, 0x57, 0x4e, 0x10, 0x02, 0x12, 0x08, 0x0a, 0x04,
	0x48, 0x54, 0x4d, 0x4c, 0x10, 0x03, 0x12, 0x08, 0x0a, 0x04, 0x52, 0x49, 0x43, 0x48, 0x10, 0x04,
	0x3a, 0x29, 0xb2, 0xe3, 0x0a, 0x25, 0xba, 0xe3, 0x0a, 0x09, 0x73, 0x74, 0x61, 0x74, 0x65, 0x6d,
	0x65, 0x6e, 0x74, 0xc2, 0xe3, 0x0a, 0x14, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x61, 0x74,
	0x6c, 0x61, 0x73, 0x2e, 0x50, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x42, 0x2d, 0x5a, 0x2b, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70,
	0x2f, 0x67, 0x6f, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x61,
	0x74, 0x6c, 0x61, 0x73, 0x3b, 0x61, 0x74, 0x6c, 0x61, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_eolymp_atlas_statement_proto_rawDescOnce sync.Once
	file_eolymp_atlas_statement_proto_rawDescData = file_eolymp_atlas_statement_proto_rawDesc
)

func file_eolymp_atlas_statement_proto_rawDescGZIP() []byte {
	file_eolymp_atlas_statement_proto_rawDescOnce.Do(func() {
		file_eolymp_atlas_statement_proto_rawDescData = protoimpl.X.CompressGZIP(file_eolymp_atlas_statement_proto_rawDescData)
	})
	return file_eolymp_atlas_statement_proto_rawDescData
}

var file_eolymp_atlas_statement_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_eolymp_atlas_statement_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_eolymp_atlas_statement_proto_goTypes = []interface{}{
	(Statement_Format)(0),        // 0: eolymp.atlas.Statement.Format
	(*Statement)(nil),            // 1: eolymp.atlas.Statement
	(*typewriter.Container)(nil), // 2: eolymp.typewriter.Container
}
var file_eolymp_atlas_statement_proto_depIdxs = []int32{
	2, // 0: eolymp.atlas.Statement.content_rich:type_name -> eolymp.typewriter.Container
	0, // 1: eolymp.atlas.Statement.format:type_name -> eolymp.atlas.Statement.Format
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_eolymp_atlas_statement_proto_init() }
func file_eolymp_atlas_statement_proto_init() {
	if File_eolymp_atlas_statement_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_eolymp_atlas_statement_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Statement); i {
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
			RawDescriptor: file_eolymp_atlas_statement_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_eolymp_atlas_statement_proto_goTypes,
		DependencyIndexes: file_eolymp_atlas_statement_proto_depIdxs,
		EnumInfos:         file_eolymp_atlas_statement_proto_enumTypes,
		MessageInfos:      file_eolymp_atlas_statement_proto_msgTypes,
	}.Build()
	File_eolymp_atlas_statement_proto = out.File
	file_eolymp_atlas_statement_proto_rawDesc = nil
	file_eolymp_atlas_statement_proto_goTypes = nil
	file_eolymp_atlas_statement_proto_depIdxs = nil
}
