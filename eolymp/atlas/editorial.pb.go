// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.29.4
// source: eolymp/atlas/editorial.proto

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

type Editorial_Extra int32

const (
	Editorial_NO_EXTRA       Editorial_Extra = 0
	Editorial_CONTENT_RENDER Editorial_Extra = 1
	Editorial_CONTENT_VALUE  Editorial_Extra = 2
)

// Enum value maps for Editorial_Extra.
var (
	Editorial_Extra_name = map[int32]string{
		0: "NO_EXTRA",
		1: "CONTENT_RENDER",
		2: "CONTENT_VALUE",
	}
	Editorial_Extra_value = map[string]int32{
		"NO_EXTRA":       0,
		"CONTENT_RENDER": 1,
		"CONTENT_VALUE":  2,
	}
)

func (x Editorial_Extra) Enum() *Editorial_Extra {
	p := new(Editorial_Extra)
	*p = x
	return p
}

func (x Editorial_Extra) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Editorial_Extra) Descriptor() protoreflect.EnumDescriptor {
	return file_eolymp_atlas_editorial_proto_enumTypes[0].Descriptor()
}

func (Editorial_Extra) Type() protoreflect.EnumType {
	return &file_eolymp_atlas_editorial_proto_enumTypes[0]
}

func (x Editorial_Extra) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Editorial_Extra.Descriptor instead.
func (Editorial_Extra) EnumDescriptor() ([]byte, []int) {
	return file_eolymp_atlas_editorial_proto_rawDescGZIP(), []int{0, 0}
}

type Editorial_Patch int32

const (
	Editorial_UNKNOWN_PATCH       Editorial_Patch = 0
	Editorial_PATCH_ALL           Editorial_Patch = 1
	Editorial_PATCH_LOCALE        Editorial_Patch = 2
	Editorial_PATCH_AUTOMATIC     Editorial_Patch = 3
	Editorial_PATCH_DRAFT         Editorial_Patch = 4
	Editorial_PATCH_CONTENT       Editorial_Patch = 5
	Editorial_PATCH_DOWNLOAD_LINK Editorial_Patch = 6
	Editorial_PATCH_AUTHOR_ID     Editorial_Patch = 7
)

// Enum value maps for Editorial_Patch.
var (
	Editorial_Patch_name = map[int32]string{
		0: "UNKNOWN_PATCH",
		1: "PATCH_ALL",
		2: "PATCH_LOCALE",
		3: "PATCH_AUTOMATIC",
		4: "PATCH_DRAFT",
		5: "PATCH_CONTENT",
		6: "PATCH_DOWNLOAD_LINK",
		7: "PATCH_AUTHOR_ID",
	}
	Editorial_Patch_value = map[string]int32{
		"UNKNOWN_PATCH":       0,
		"PATCH_ALL":           1,
		"PATCH_LOCALE":        2,
		"PATCH_AUTOMATIC":     3,
		"PATCH_DRAFT":         4,
		"PATCH_CONTENT":       5,
		"PATCH_DOWNLOAD_LINK": 6,
		"PATCH_AUTHOR_ID":     7,
	}
)

func (x Editorial_Patch) Enum() *Editorial_Patch {
	p := new(Editorial_Patch)
	*p = x
	return p
}

func (x Editorial_Patch) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Editorial_Patch) Descriptor() protoreflect.EnumDescriptor {
	return file_eolymp_atlas_editorial_proto_enumTypes[1].Descriptor()
}

func (Editorial_Patch) Type() protoreflect.EnumType {
	return &file_eolymp_atlas_editorial_proto_enumTypes[1]
}

func (x Editorial_Patch) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Editorial_Patch.Descriptor instead.
func (Editorial_Patch) EnumDescriptor() ([]byte, []int) {
	return file_eolymp_atlas_editorial_proto_rawDescGZIP(), []int{0, 1}
}

type Editorial struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`                // Unique identifier, assigned when editorial is created. Keep empty when creating new editorial.
	Locale        string                 `protobuf:"bytes,3,opt,name=locale,proto3" json:"locale,omitempty"`        // Editorial locale code, should consist of two lowercase latin letters.
	Automatic     bool                   `protobuf:"varint,8,opt,name=automatic,proto3" json:"automatic,omitempty"` // if true means the editorial was created automatically (probably translation from another language)
	Draft         bool                   `protobuf:"varint,10,opt,name=draft,proto3" json:"draft,omitempty"`
	Content       *ecm.Content           `protobuf:"bytes,6,opt,name=content,proto3" json:"content,omitempty"`
	DownloadLink  string                 `protobuf:"bytes,7,opt,name=download_link,json=downloadLink,proto3" json:"download_link,omitempty"`
	AuthorId      string                 `protobuf:"bytes,103,opt,name=author_id,json=authorId,proto3" json:"author_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Editorial) Reset() {
	*x = Editorial{}
	mi := &file_eolymp_atlas_editorial_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Editorial) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Editorial) ProtoMessage() {}

func (x *Editorial) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_atlas_editorial_proto_msgTypes[0]
	if x != nil {
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

func (x *Editorial) GetLocale() string {
	if x != nil {
		return x.Locale
	}
	return ""
}

func (x *Editorial) GetAutomatic() bool {
	if x != nil {
		return x.Automatic
	}
	return false
}

func (x *Editorial) GetDraft() bool {
	if x != nil {
		return x.Draft
	}
	return false
}

func (x *Editorial) GetContent() *ecm.Content {
	if x != nil {
		return x.Content
	}
	return nil
}

func (x *Editorial) GetDownloadLink() string {
	if x != nil {
		return x.DownloadLink
	}
	return ""
}

func (x *Editorial) GetAuthorId() string {
	if x != nil {
		return x.AuthorId
	}
	return ""
}

var File_eolymp_atlas_editorial_proto protoreflect.FileDescriptor

const file_eolymp_atlas_editorial_proto_rawDesc = "" +
	"\n" +
	"\x1ceolymp/atlas/editorial.proto\x12\feolymp.atlas\x1a\x18eolymp/ecm/content.proto\"\xbb\x03\n" +
	"\tEditorial\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x12\x16\n" +
	"\x06locale\x18\x03 \x01(\tR\x06locale\x12\x1c\n" +
	"\tautomatic\x18\b \x01(\bR\tautomatic\x12\x14\n" +
	"\x05draft\x18\n" +
	" \x01(\bR\x05draft\x12-\n" +
	"\acontent\x18\x06 \x01(\v2\x13.eolymp.ecm.ContentR\acontent\x12#\n" +
	"\rdownload_link\x18\a \x01(\tR\fdownloadLink\x12\x1b\n" +
	"\tauthor_id\x18g \x01(\tR\bauthorId\"<\n" +
	"\x05Extra\x12\f\n" +
	"\bNO_EXTRA\x10\x00\x12\x12\n" +
	"\x0eCONTENT_RENDER\x10\x01\x12\x11\n" +
	"\rCONTENT_VALUE\x10\x02\"\xa2\x01\n" +
	"\x05Patch\x12\x11\n" +
	"\rUNKNOWN_PATCH\x10\x00\x12\r\n" +
	"\tPATCH_ALL\x10\x01\x12\x10\n" +
	"\fPATCH_LOCALE\x10\x02\x12\x13\n" +
	"\x0fPATCH_AUTOMATIC\x10\x03\x12\x0f\n" +
	"\vPATCH_DRAFT\x10\x04\x12\x11\n" +
	"\rPATCH_CONTENT\x10\x05\x12\x17\n" +
	"\x13PATCH_DOWNLOAD_LINK\x10\x06\x12\x13\n" +
	"\x0fPATCH_AUTHOR_ID\x10\aB-Z+github.com/eolymp/go-sdk/eolymp/atlas;atlasb\x06proto3"

var (
	file_eolymp_atlas_editorial_proto_rawDescOnce sync.Once
	file_eolymp_atlas_editorial_proto_rawDescData []byte
)

func file_eolymp_atlas_editorial_proto_rawDescGZIP() []byte {
	file_eolymp_atlas_editorial_proto_rawDescOnce.Do(func() {
		file_eolymp_atlas_editorial_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_eolymp_atlas_editorial_proto_rawDesc), len(file_eolymp_atlas_editorial_proto_rawDesc)))
	})
	return file_eolymp_atlas_editorial_proto_rawDescData
}

var file_eolymp_atlas_editorial_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_eolymp_atlas_editorial_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_eolymp_atlas_editorial_proto_goTypes = []any{
	(Editorial_Extra)(0), // 0: eolymp.atlas.Editorial.Extra
	(Editorial_Patch)(0), // 1: eolymp.atlas.Editorial.Patch
	(*Editorial)(nil),    // 2: eolymp.atlas.Editorial
	(*ecm.Content)(nil),  // 3: eolymp.ecm.Content
}
var file_eolymp_atlas_editorial_proto_depIdxs = []int32{
	3, // 0: eolymp.atlas.Editorial.content:type_name -> eolymp.ecm.Content
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_eolymp_atlas_editorial_proto_rawDesc), len(file_eolymp_atlas_editorial_proto_rawDesc)),
			NumEnums:      2,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_eolymp_atlas_editorial_proto_goTypes,
		DependencyIndexes: file_eolymp_atlas_editorial_proto_depIdxs,
		EnumInfos:         file_eolymp_atlas_editorial_proto_enumTypes,
		MessageInfos:      file_eolymp_atlas_editorial_proto_msgTypes,
	}.Build()
	File_eolymp_atlas_editorial_proto = out.File
	file_eolymp_atlas_editorial_proto_goTypes = nil
	file_eolymp_atlas_editorial_proto_depIdxs = nil
}
