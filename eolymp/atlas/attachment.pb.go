// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v4.24.4
// source: eolymp/atlas/attachment.proto

package atlas

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

type Attachment struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	ProblemId     string                 `protobuf:"bytes,2,opt,name=problem_id,json=problemId,proto3" json:"problem_id,omitempty"` // deprecate
	Name          string                 `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Link          string                 `protobuf:"bytes,4,opt,name=link,proto3" json:"link,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Attachment) Reset() {
	*x = Attachment{}
	mi := &file_eolymp_atlas_attachment_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Attachment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Attachment) ProtoMessage() {}

func (x *Attachment) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_atlas_attachment_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Attachment.ProtoReflect.Descriptor instead.
func (*Attachment) Descriptor() ([]byte, []int) {
	return file_eolymp_atlas_attachment_proto_rawDescGZIP(), []int{0}
}

func (x *Attachment) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Attachment) GetProblemId() string {
	if x != nil {
		return x.ProblemId
	}
	return ""
}

func (x *Attachment) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Attachment) GetLink() string {
	if x != nil {
		return x.Link
	}
	return ""
}

var File_eolymp_atlas_attachment_proto protoreflect.FileDescriptor

const file_eolymp_atlas_attachment_proto_rawDesc = "" +
	"\n" +
	"\x1deolymp/atlas/attachment.proto\x12\feolymp.atlas\"c\n" +
	"\n" +
	"Attachment\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x12\x1d\n" +
	"\n" +
	"problem_id\x18\x02 \x01(\tR\tproblemId\x12\x12\n" +
	"\x04name\x18\x03 \x01(\tR\x04name\x12\x12\n" +
	"\x04link\x18\x04 \x01(\tR\x04linkB-Z+github.com/eolymp/go-sdk/eolymp/atlas;atlasb\x06proto3"

var (
	file_eolymp_atlas_attachment_proto_rawDescOnce sync.Once
	file_eolymp_atlas_attachment_proto_rawDescData []byte
)

func file_eolymp_atlas_attachment_proto_rawDescGZIP() []byte {
	file_eolymp_atlas_attachment_proto_rawDescOnce.Do(func() {
		file_eolymp_atlas_attachment_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_eolymp_atlas_attachment_proto_rawDesc), len(file_eolymp_atlas_attachment_proto_rawDesc)))
	})
	return file_eolymp_atlas_attachment_proto_rawDescData
}

var file_eolymp_atlas_attachment_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_eolymp_atlas_attachment_proto_goTypes = []any{
	(*Attachment)(nil), // 0: eolymp.atlas.Attachment
}
var file_eolymp_atlas_attachment_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_eolymp_atlas_attachment_proto_init() }
func file_eolymp_atlas_attachment_proto_init() {
	if File_eolymp_atlas_attachment_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_eolymp_atlas_attachment_proto_rawDesc), len(file_eolymp_atlas_attachment_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_eolymp_atlas_attachment_proto_goTypes,
		DependencyIndexes: file_eolymp_atlas_attachment_proto_depIdxs,
		MessageInfos:      file_eolymp_atlas_attachment_proto_msgTypes,
	}.Build()
	File_eolymp_atlas_attachment_proto = out.File
	file_eolymp_atlas_attachment_proto_goTypes = nil
	file_eolymp_atlas_attachment_proto_depIdxs = nil
}
