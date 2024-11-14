// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v3.21.12
// source: eolymp/atlas/version.proto

package atlas

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Version struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         string                 `protobuf:"bytes,6,opt,name=id,proto3" json:"id,omitempty"`
	Number     uint32                 `protobuf:"varint,1,opt,name=number,proto3" json:"number,omitempty"`
	CreatedAt  *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	CreatedBy  string                 `protobuf:"bytes,3,opt,name=created_by,json=createdBy,proto3" json:"created_by,omitempty"`
	ChangeOp   string                 `protobuf:"bytes,4,opt,name=change_op,json=changeOp,proto3" json:"change_op,omitempty"`
	ChangePath string                 `protobuf:"bytes,5,opt,name=change_path,json=changePath,proto3" json:"change_path,omitempty"`
	Cursor     string                 `protobuf:"bytes,100,opt,name=cursor,proto3" json:"cursor,omitempty"`
}

func (x *Version) Reset() {
	*x = Version{}
	mi := &file_eolymp_atlas_version_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Version) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Version) ProtoMessage() {}

func (x *Version) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_atlas_version_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Version.ProtoReflect.Descriptor instead.
func (*Version) Descriptor() ([]byte, []int) {
	return file_eolymp_atlas_version_proto_rawDescGZIP(), []int{0}
}

func (x *Version) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Version) GetNumber() uint32 {
	if x != nil {
		return x.Number
	}
	return 0
}

func (x *Version) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Version) GetCreatedBy() string {
	if x != nil {
		return x.CreatedBy
	}
	return ""
}

func (x *Version) GetChangeOp() string {
	if x != nil {
		return x.ChangeOp
	}
	return ""
}

func (x *Version) GetChangePath() string {
	if x != nil {
		return x.ChangePath
	}
	return ""
}

func (x *Version) GetCursor() string {
	if x != nil {
		return x.Cursor
	}
	return ""
}

var File_eolymp_atlas_version_proto protoreflect.FileDescriptor

var file_eolymp_atlas_version_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x61, 0x74, 0x6c, 0x61, 0x73, 0x2f, 0x76,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x65, 0x6f,
	0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x61, 0x74, 0x6c, 0x61, 0x73, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe1, 0x01, 0x0a, 0x07,
	0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12,
	0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x62, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x68, 0x61,
	0x6e, 0x67, 0x65, 0x5f, 0x6f, 0x70, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x68,
	0x61, 0x6e, 0x67, 0x65, 0x4f, 0x70, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65,
	0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x68, 0x61,
	0x6e, 0x67, 0x65, 0x50, 0x61, 0x74, 0x68, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x75, 0x72, 0x73, 0x6f,
	0x72, 0x18, 0x64, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x63, 0x75, 0x72, 0x73, 0x6f, 0x72, 0x42,
	0x2d, 0x5a, 0x2b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6f,
	0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x67, 0x6f, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x65, 0x6f, 0x6c, 0x79,
	0x6d, 0x70, 0x2f, 0x61, 0x74, 0x6c, 0x61, 0x73, 0x3b, 0x61, 0x74, 0x6c, 0x61, 0x73, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_eolymp_atlas_version_proto_rawDescOnce sync.Once
	file_eolymp_atlas_version_proto_rawDescData = file_eolymp_atlas_version_proto_rawDesc
)

func file_eolymp_atlas_version_proto_rawDescGZIP() []byte {
	file_eolymp_atlas_version_proto_rawDescOnce.Do(func() {
		file_eolymp_atlas_version_proto_rawDescData = protoimpl.X.CompressGZIP(file_eolymp_atlas_version_proto_rawDescData)
	})
	return file_eolymp_atlas_version_proto_rawDescData
}

var file_eolymp_atlas_version_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_eolymp_atlas_version_proto_goTypes = []any{
	(*Version)(nil),               // 0: eolymp.atlas.Version
	(*timestamppb.Timestamp)(nil), // 1: google.protobuf.Timestamp
}
var file_eolymp_atlas_version_proto_depIdxs = []int32{
	1, // 0: eolymp.atlas.Version.created_at:type_name -> google.protobuf.Timestamp
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_eolymp_atlas_version_proto_init() }
func file_eolymp_atlas_version_proto_init() {
	if File_eolymp_atlas_version_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_eolymp_atlas_version_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_eolymp_atlas_version_proto_goTypes,
		DependencyIndexes: file_eolymp_atlas_version_proto_depIdxs,
		MessageInfos:      file_eolymp_atlas_version_proto_msgTypes,
	}.Build()
	File_eolymp_atlas_version_proto = out.File
	file_eolymp_atlas_version_proto_rawDesc = nil
	file_eolymp_atlas_version_proto_goTypes = nil
	file_eolymp_atlas_version_proto_depIdxs = nil
}
