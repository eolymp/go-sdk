// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v4.24.4
// source: eolymp/atlas/statement.proto

package atlas

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

type Statement_Extra int32

const (
	Statement_UNKNOWN_EXTRA  Statement_Extra = 0
	Statement_CONTENT_RENDER Statement_Extra = 1
	Statement_CONTENT_VALUE  Statement_Extra = 2
)

// Enum value maps for Statement_Extra.
var (
	Statement_Extra_name = map[int32]string{
		0: "UNKNOWN_EXTRA",
		1: "CONTENT_RENDER",
		2: "CONTENT_VALUE",
	}
	Statement_Extra_value = map[string]int32{
		"UNKNOWN_EXTRA":  0,
		"CONTENT_RENDER": 1,
		"CONTENT_VALUE":  2,
	}
)

func (x Statement_Extra) Enum() *Statement_Extra {
	p := new(Statement_Extra)
	*p = x
	return p
}

func (x Statement_Extra) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Statement_Extra) Descriptor() protoreflect.EnumDescriptor {
	return file_eolymp_atlas_statement_proto_enumTypes[0].Descriptor()
}

func (Statement_Extra) Type() protoreflect.EnumType {
	return &file_eolymp_atlas_statement_proto_enumTypes[0]
}

func (x Statement_Extra) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Statement_Extra.Descriptor instead.
func (Statement_Extra) EnumDescriptor() ([]byte, []int) {
	return file_eolymp_atlas_statement_proto_rawDescGZIP(), []int{0, 0}
}

type Statement_Patch int32

const (
	Statement_UNKNOWN_PATCH       Statement_Patch = 0
	Statement_PATCH_ALL           Statement_Patch = 1
	Statement_PATCH_LOCALE        Statement_Patch = 2
	Statement_PATCH_AUTOMATIC     Statement_Patch = 8
	Statement_PATCH_DRAFT         Statement_Patch = 9
	Statement_PATCH_TITLE         Statement_Patch = 3
	Statement_PATCH_CONTENT       Statement_Patch = 4
	Statement_PATCH_DOWNLOAD_LINK Statement_Patch = 5
	Statement_PATCH_AUTHOR        Statement_Patch = 6
	Statement_PATCH_SOURCE        Statement_Patch = 7
	Statement_PATCH_AUTHOR_ID     Statement_Patch = 10
)

// Enum value maps for Statement_Patch.
var (
	Statement_Patch_name = map[int32]string{
		0:  "UNKNOWN_PATCH",
		1:  "PATCH_ALL",
		2:  "PATCH_LOCALE",
		8:  "PATCH_AUTOMATIC",
		9:  "PATCH_DRAFT",
		3:  "PATCH_TITLE",
		4:  "PATCH_CONTENT",
		5:  "PATCH_DOWNLOAD_LINK",
		6:  "PATCH_AUTHOR",
		7:  "PATCH_SOURCE",
		10: "PATCH_AUTHOR_ID",
	}
	Statement_Patch_value = map[string]int32{
		"UNKNOWN_PATCH":       0,
		"PATCH_ALL":           1,
		"PATCH_LOCALE":        2,
		"PATCH_AUTOMATIC":     8,
		"PATCH_DRAFT":         9,
		"PATCH_TITLE":         3,
		"PATCH_CONTENT":       4,
		"PATCH_DOWNLOAD_LINK": 5,
		"PATCH_AUTHOR":        6,
		"PATCH_SOURCE":        7,
		"PATCH_AUTHOR_ID":     10,
	}
)

func (x Statement_Patch) Enum() *Statement_Patch {
	p := new(Statement_Patch)
	*p = x
	return p
}

func (x Statement_Patch) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Statement_Patch) Descriptor() protoreflect.EnumDescriptor {
	return file_eolymp_atlas_statement_proto_enumTypes[1].Descriptor()
}

func (Statement_Patch) Type() protoreflect.EnumType {
	return &file_eolymp_atlas_statement_proto_enumTypes[1]
}

func (x Statement_Patch) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Statement_Patch.Descriptor instead.
func (Statement_Patch) EnumDescriptor() ([]byte, []int) {
	return file_eolymp_atlas_statement_proto_rawDescGZIP(), []int{0, 1}
}

type Statement struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Locale        string                 `protobuf:"bytes,3,opt,name=locale,proto3" json:"locale,omitempty"`
	Automatic     bool                   `protobuf:"varint,8,opt,name=automatic,proto3" json:"automatic,omitempty"` // if true means the statement was created automatically (probably translation from another language)
	Draft         bool                   `protobuf:"varint,9,opt,name=draft,proto3" json:"draft,omitempty"`
	Title         string                 `protobuf:"bytes,4,opt,name=title,proto3" json:"title,omitempty"`
	Content       *ecm.Content           `protobuf:"bytes,6,opt,name=content,proto3" json:"content,omitempty"`
	DownloadLink  string                 `protobuf:"bytes,7,opt,name=download_link,json=downloadLink,proto3" json:"download_link,omitempty"` // Statement download link, allows to download statement in original format.
	Author        string                 `protobuf:"bytes,101,opt,name=author,proto3" json:"author,omitempty"`                               // Problem author name.
	Source        string                 `protobuf:"bytes,102,opt,name=source,proto3" json:"source,omitempty"`                               // Problem source, name of the contest or olympiad where this problem was initially published.
	AuthorId      string                 `protobuf:"bytes,103,opt,name=author_id,json=authorId,proto3" json:"author_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Statement) Reset() {
	*x = Statement{}
	mi := &file_eolymp_atlas_statement_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Statement) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Statement) ProtoMessage() {}

func (x *Statement) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_atlas_statement_proto_msgTypes[0]
	if x != nil {
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

func (x *Statement) GetLocale() string {
	if x != nil {
		return x.Locale
	}
	return ""
}

func (x *Statement) GetAutomatic() bool {
	if x != nil {
		return x.Automatic
	}
	return false
}

func (x *Statement) GetDraft() bool {
	if x != nil {
		return x.Draft
	}
	return false
}

func (x *Statement) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Statement) GetContent() *ecm.Content {
	if x != nil {
		return x.Content
	}
	return nil
}

func (x *Statement) GetDownloadLink() string {
	if x != nil {
		return x.DownloadLink
	}
	return ""
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

func (x *Statement) GetAuthorId() string {
	if x != nil {
		return x.AuthorId
	}
	return ""
}

var File_eolymp_atlas_statement_proto protoreflect.FileDescriptor

var file_eolymp_atlas_statement_proto_rawDesc = string([]byte{
	0x0a, 0x1c, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x61, 0x74, 0x6c, 0x61, 0x73, 0x2f, 0x73,
	0x74, 0x61, 0x74, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c,
	0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x61, 0x74, 0x6c, 0x61, 0x73, 0x1a, 0x18, 0x65, 0x6f,
	0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x65, 0x63, 0x6d, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xbb, 0x04, 0x0a, 0x09, 0x53, 0x74, 0x61, 0x74, 0x65,
	0x6d, 0x65, 0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x65, 0x12, 0x1c, 0x0a, 0x09,
	0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x69, 0x63, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x09, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x69, 0x63, 0x12, 0x14, 0x0a, 0x05, 0x64, 0x72,
	0x61, 0x66, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x64, 0x72, 0x61, 0x66, 0x74,
	0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x2d, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70,
	0x2e, 0x65, 0x63, 0x6d, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x52, 0x07, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x64, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61,
	0x64, 0x5f, 0x6c, 0x69, 0x6e, 0x6b, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x64, 0x6f,
	0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x4c, 0x69, 0x6e, 0x6b, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x75,
	0x74, 0x68, 0x6f, 0x72, 0x18, 0x65, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x75, 0x74, 0x68,
	0x6f, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x66, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x61, 0x75,
	0x74, 0x68, 0x6f, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x67, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x61,
	0x75, 0x74, 0x68, 0x6f, 0x72, 0x49, 0x64, 0x22, 0x41, 0x0a, 0x05, 0x45, 0x78, 0x74, 0x72, 0x61,
	0x12, 0x11, 0x0a, 0x0d, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x5f, 0x45, 0x58, 0x54, 0x52,
	0x41, 0x10, 0x00, 0x12, 0x12, 0x0a, 0x0e, 0x43, 0x4f, 0x4e, 0x54, 0x45, 0x4e, 0x54, 0x5f, 0x52,
	0x45, 0x4e, 0x44, 0x45, 0x52, 0x10, 0x01, 0x12, 0x11, 0x0a, 0x0d, 0x43, 0x4f, 0x4e, 0x54, 0x45,
	0x4e, 0x54, 0x5f, 0x56, 0x41, 0x4c, 0x55, 0x45, 0x10, 0x02, 0x22, 0xd7, 0x01, 0x0a, 0x05, 0x50,
	0x61, 0x74, 0x63, 0x68, 0x12, 0x11, 0x0a, 0x0d, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x5f,
	0x50, 0x41, 0x54, 0x43, 0x48, 0x10, 0x00, 0x12, 0x0d, 0x0a, 0x09, 0x50, 0x41, 0x54, 0x43, 0x48,
	0x5f, 0x41, 0x4c, 0x4c, 0x10, 0x01, 0x12, 0x10, 0x0a, 0x0c, 0x50, 0x41, 0x54, 0x43, 0x48, 0x5f,
	0x4c, 0x4f, 0x43, 0x41, 0x4c, 0x45, 0x10, 0x02, 0x12, 0x13, 0x0a, 0x0f, 0x50, 0x41, 0x54, 0x43,
	0x48, 0x5f, 0x41, 0x55, 0x54, 0x4f, 0x4d, 0x41, 0x54, 0x49, 0x43, 0x10, 0x08, 0x12, 0x0f, 0x0a,
	0x0b, 0x50, 0x41, 0x54, 0x43, 0x48, 0x5f, 0x44, 0x52, 0x41, 0x46, 0x54, 0x10, 0x09, 0x12, 0x0f,
	0x0a, 0x0b, 0x50, 0x41, 0x54, 0x43, 0x48, 0x5f, 0x54, 0x49, 0x54, 0x4c, 0x45, 0x10, 0x03, 0x12,
	0x11, 0x0a, 0x0d, 0x50, 0x41, 0x54, 0x43, 0x48, 0x5f, 0x43, 0x4f, 0x4e, 0x54, 0x45, 0x4e, 0x54,
	0x10, 0x04, 0x12, 0x17, 0x0a, 0x13, 0x50, 0x41, 0x54, 0x43, 0x48, 0x5f, 0x44, 0x4f, 0x57, 0x4e,
	0x4c, 0x4f, 0x41, 0x44, 0x5f, 0x4c, 0x49, 0x4e, 0x4b, 0x10, 0x05, 0x12, 0x10, 0x0a, 0x0c, 0x50,
	0x41, 0x54, 0x43, 0x48, 0x5f, 0x41, 0x55, 0x54, 0x48, 0x4f, 0x52, 0x10, 0x06, 0x12, 0x10, 0x0a,
	0x0c, 0x50, 0x41, 0x54, 0x43, 0x48, 0x5f, 0x53, 0x4f, 0x55, 0x52, 0x43, 0x45, 0x10, 0x07, 0x12,
	0x13, 0x0a, 0x0f, 0x50, 0x41, 0x54, 0x43, 0x48, 0x5f, 0x41, 0x55, 0x54, 0x48, 0x4f, 0x52, 0x5f,
	0x49, 0x44, 0x10, 0x0a, 0x42, 0x2d, 0x5a, 0x2b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x67, 0x6f, 0x2d, 0x73, 0x64, 0x6b,
	0x2f, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x61, 0x74, 0x6c, 0x61, 0x73, 0x3b, 0x61, 0x74,
	0x6c, 0x61, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_eolymp_atlas_statement_proto_rawDescOnce sync.Once
	file_eolymp_atlas_statement_proto_rawDescData []byte
)

func file_eolymp_atlas_statement_proto_rawDescGZIP() []byte {
	file_eolymp_atlas_statement_proto_rawDescOnce.Do(func() {
		file_eolymp_atlas_statement_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_eolymp_atlas_statement_proto_rawDesc), len(file_eolymp_atlas_statement_proto_rawDesc)))
	})
	return file_eolymp_atlas_statement_proto_rawDescData
}

var file_eolymp_atlas_statement_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_eolymp_atlas_statement_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_eolymp_atlas_statement_proto_goTypes = []any{
	(Statement_Extra)(0), // 0: eolymp.atlas.Statement.Extra
	(Statement_Patch)(0), // 1: eolymp.atlas.Statement.Patch
	(*Statement)(nil),    // 2: eolymp.atlas.Statement
	(*ecm.Content)(nil),  // 3: eolymp.ecm.Content
}
var file_eolymp_atlas_statement_proto_depIdxs = []int32{
	3, // 0: eolymp.atlas.Statement.content:type_name -> eolymp.ecm.Content
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_eolymp_atlas_statement_proto_init() }
func file_eolymp_atlas_statement_proto_init() {
	if File_eolymp_atlas_statement_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_eolymp_atlas_statement_proto_rawDesc), len(file_eolymp_atlas_statement_proto_rawDesc)),
			NumEnums:      2,
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
	file_eolymp_atlas_statement_proto_goTypes = nil
	file_eolymp_atlas_statement_proto_depIdxs = nil
}
