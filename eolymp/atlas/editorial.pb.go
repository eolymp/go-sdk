// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v3.21.12
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
	Editorial_PATCH_LOCALE        Editorial_Patch = 1
	Editorial_PATCH_AUTOMATIC     Editorial_Patch = 2
	Editorial_PATCH_DRAFT         Editorial_Patch = 3
	Editorial_PATCH_CONTENT       Editorial_Patch = 4
	Editorial_PATCH_DOWNLOAD_LINK Editorial_Patch = 5
)

// Enum value maps for Editorial_Patch.
var (
	Editorial_Patch_name = map[int32]string{
		0: "UNKNOWN_PATCH",
		1: "PATCH_LOCALE",
		2: "PATCH_AUTOMATIC",
		3: "PATCH_DRAFT",
		4: "PATCH_CONTENT",
		5: "PATCH_DOWNLOAD_LINK",
	}
	Editorial_Patch_value = map[string]int32{
		"UNKNOWN_PATCH":       0,
		"PATCH_LOCALE":        1,
		"PATCH_AUTOMATIC":     2,
		"PATCH_DRAFT":         3,
		"PATCH_CONTENT":       4,
		"PATCH_DOWNLOAD_LINK": 5,
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
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           string       `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`                // Unique identifier, assigned when editorial is created. Keep empty when creating new editorial.
	Locale       string       `protobuf:"bytes,3,opt,name=locale,proto3" json:"locale,omitempty"`        // Editorial locale code, should consist of two lowercase latin letters.
	Automatic    bool         `protobuf:"varint,8,opt,name=automatic,proto3" json:"automatic,omitempty"` // if true means the editorial was created automatically (probably translation from another language)
	Draft        bool         `protobuf:"varint,10,opt,name=draft,proto3" json:"draft,omitempty"`
	Content      *ecm.Content `protobuf:"bytes,6,opt,name=content,proto3" json:"content,omitempty"`
	DownloadLink string       `protobuf:"bytes,7,opt,name=download_link,json=downloadLink,proto3" json:"download_link,omitempty"`
	AuthorId     string       `protobuf:"bytes,103,opt,name=author_id,json=authorId,proto3" json:"author_id,omitempty"`
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

var file_eolymp_atlas_editorial_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x61, 0x74, 0x6c, 0x61, 0x73, 0x2f, 0x65,
	0x64, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c,
	0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x61, 0x74, 0x6c, 0x61, 0x73, 0x1a, 0x18, 0x65, 0x6f,
	0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x65, 0x63, 0x6d, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x96, 0x03, 0x0a, 0x09, 0x45, 0x64, 0x69, 0x74, 0x6f,
	0x72, 0x69, 0x61, 0x6c, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x65, 0x12, 0x1c, 0x0a, 0x09,
	0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x69, 0x63, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x09, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x69, 0x63, 0x12, 0x14, 0x0a, 0x05, 0x64, 0x72,
	0x61, 0x66, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x64, 0x72, 0x61, 0x66, 0x74,
	0x12, 0x2d, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x13, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x65, 0x63, 0x6d, 0x2e, 0x43,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12,
	0x23, 0x0a, 0x0d, 0x64, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x6c, 0x69, 0x6e, 0x6b,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x64, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64,
	0x4c, 0x69, 0x6e, 0x6b, 0x12, 0x1b, 0x0a, 0x09, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x67, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x49,
	0x64, 0x22, 0x3c, 0x0a, 0x05, 0x45, 0x78, 0x74, 0x72, 0x61, 0x12, 0x0c, 0x0a, 0x08, 0x4e, 0x4f,
	0x5f, 0x45, 0x58, 0x54, 0x52, 0x41, 0x10, 0x00, 0x12, 0x12, 0x0a, 0x0e, 0x43, 0x4f, 0x4e, 0x54,
	0x45, 0x4e, 0x54, 0x5f, 0x52, 0x45, 0x4e, 0x44, 0x45, 0x52, 0x10, 0x01, 0x12, 0x11, 0x0a, 0x0d,
	0x43, 0x4f, 0x4e, 0x54, 0x45, 0x4e, 0x54, 0x5f, 0x56, 0x41, 0x4c, 0x55, 0x45, 0x10, 0x02, 0x22,
	0x7e, 0x0a, 0x05, 0x50, 0x61, 0x74, 0x63, 0x68, 0x12, 0x11, 0x0a, 0x0d, 0x55, 0x4e, 0x4b, 0x4e,
	0x4f, 0x57, 0x4e, 0x5f, 0x50, 0x41, 0x54, 0x43, 0x48, 0x10, 0x00, 0x12, 0x10, 0x0a, 0x0c, 0x50,
	0x41, 0x54, 0x43, 0x48, 0x5f, 0x4c, 0x4f, 0x43, 0x41, 0x4c, 0x45, 0x10, 0x01, 0x12, 0x13, 0x0a,
	0x0f, 0x50, 0x41, 0x54, 0x43, 0x48, 0x5f, 0x41, 0x55, 0x54, 0x4f, 0x4d, 0x41, 0x54, 0x49, 0x43,
	0x10, 0x02, 0x12, 0x0f, 0x0a, 0x0b, 0x50, 0x41, 0x54, 0x43, 0x48, 0x5f, 0x44, 0x52, 0x41, 0x46,
	0x54, 0x10, 0x03, 0x12, 0x11, 0x0a, 0x0d, 0x50, 0x41, 0x54, 0x43, 0x48, 0x5f, 0x43, 0x4f, 0x4e,
	0x54, 0x45, 0x4e, 0x54, 0x10, 0x04, 0x12, 0x17, 0x0a, 0x13, 0x50, 0x41, 0x54, 0x43, 0x48, 0x5f,
	0x44, 0x4f, 0x57, 0x4e, 0x4c, 0x4f, 0x41, 0x44, 0x5f, 0x4c, 0x49, 0x4e, 0x4b, 0x10, 0x05, 0x42,
	0x2d, 0x5a, 0x2b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6f,
	0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x67, 0x6f, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x65, 0x6f, 0x6c, 0x79,
	0x6d, 0x70, 0x2f, 0x61, 0x74, 0x6c, 0x61, 0x73, 0x3b, 0x61, 0x74, 0x6c, 0x61, 0x73, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
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
			RawDescriptor: file_eolymp_atlas_editorial_proto_rawDesc,
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
	file_eolymp_atlas_editorial_proto_rawDesc = nil
	file_eolymp_atlas_editorial_proto_goTypes = nil
	file_eolymp_atlas_editorial_proto_depIdxs = nil
}
