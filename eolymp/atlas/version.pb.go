// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v4.24.4
// source: eolymp/atlas/version.proto

package atlas

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
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

type Version_Operation int32

const (
	Version_UNKNOWN_OPERATION Version_Operation = 0
	Version_ADD               Version_Operation = 1
	Version_UPDATE            Version_Operation = 2
	Version_DELETE            Version_Operation = 3
)

// Enum value maps for Version_Operation.
var (
	Version_Operation_name = map[int32]string{
		0: "UNKNOWN_OPERATION",
		1: "ADD",
		2: "UPDATE",
		3: "DELETE",
	}
	Version_Operation_value = map[string]int32{
		"UNKNOWN_OPERATION": 0,
		"ADD":               1,
		"UPDATE":            2,
		"DELETE":            3,
	}
)

func (x Version_Operation) Enum() *Version_Operation {
	p := new(Version_Operation)
	*p = x
	return p
}

func (x Version_Operation) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Version_Operation) Descriptor() protoreflect.EnumDescriptor {
	return file_eolymp_atlas_version_proto_enumTypes[0].Descriptor()
}

func (Version_Operation) Type() protoreflect.EnumType {
	return &file_eolymp_atlas_version_proto_enumTypes[0]
}

func (x Version_Operation) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Version_Operation.Descriptor instead.
func (Version_Operation) EnumDescriptor() ([]byte, []int) {
	return file_eolymp_atlas_version_proto_rawDescGZIP(), []int{0, 0}
}

type Version struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,6,opt,name=id,proto3" json:"id,omitempty"`
	Number        uint32                 `protobuf:"varint,1,opt,name=number,proto3" json:"number,omitempty"`
	CreatedAt     *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	CreatedBy     string                 `protobuf:"bytes,3,opt,name=created_by,json=createdBy,proto3" json:"created_by,omitempty"`
	Summary       string                 `protobuf:"bytes,9,opt,name=summary,proto3" json:"summary,omitempty"`
	Changes       []*Version_Change      `protobuf:"bytes,10,rep,name=changes,proto3" json:"changes,omitempty"`
	Cursor        string                 `protobuf:"bytes,100,opt,name=cursor,proto3" json:"cursor,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
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

func (x *Version) GetSummary() string {
	if x != nil {
		return x.Summary
	}
	return ""
}

func (x *Version) GetChanges() []*Version_Change {
	if x != nil {
		return x.Changes
	}
	return nil
}

func (x *Version) GetCursor() string {
	if x != nil {
		return x.Cursor
	}
	return ""
}

type Version_Change struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Op            Version_Operation      `protobuf:"varint,1,opt,name=op,proto3,enum=eolymp.atlas.Version_Operation" json:"op,omitempty"`
	Path          string                 `protobuf:"bytes,2,opt,name=path,proto3" json:"path,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Version_Change) Reset() {
	*x = Version_Change{}
	mi := &file_eolymp_atlas_version_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Version_Change) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Version_Change) ProtoMessage() {}

func (x *Version_Change) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_atlas_version_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Version_Change.ProtoReflect.Descriptor instead.
func (*Version_Change) Descriptor() ([]byte, []int) {
	return file_eolymp_atlas_version_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Version_Change) GetOp() Version_Operation {
	if x != nil {
		return x.Op
	}
	return Version_UNKNOWN_OPERATION
}

func (x *Version_Change) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

var File_eolymp_atlas_version_proto protoreflect.FileDescriptor

const file_eolymp_atlas_version_proto_rawDesc = "" +
	"\n" +
	"\x1aeolymp/atlas/version.proto\x12\feolymp.atlas\x1a\x1fgoogle/protobuf/timestamp.proto\"\x89\x03\n" +
	"\aVersion\x12\x0e\n" +
	"\x02id\x18\x06 \x01(\tR\x02id\x12\x16\n" +
	"\x06number\x18\x01 \x01(\rR\x06number\x129\n" +
	"\n" +
	"created_at\x18\x02 \x01(\v2\x1a.google.protobuf.TimestampR\tcreatedAt\x12\x1d\n" +
	"\n" +
	"created_by\x18\x03 \x01(\tR\tcreatedBy\x12\x18\n" +
	"\asummary\x18\t \x01(\tR\asummary\x126\n" +
	"\achanges\x18\n" +
	" \x03(\v2\x1c.eolymp.atlas.Version.ChangeR\achanges\x12\x16\n" +
	"\x06cursor\x18d \x01(\tR\x06cursor\x1aM\n" +
	"\x06Change\x12/\n" +
	"\x02op\x18\x01 \x01(\x0e2\x1f.eolymp.atlas.Version.OperationR\x02op\x12\x12\n" +
	"\x04path\x18\x02 \x01(\tR\x04path\"C\n" +
	"\tOperation\x12\x15\n" +
	"\x11UNKNOWN_OPERATION\x10\x00\x12\a\n" +
	"\x03ADD\x10\x01\x12\n" +
	"\n" +
	"\x06UPDATE\x10\x02\x12\n" +
	"\n" +
	"\x06DELETE\x10\x03B-Z+github.com/eolymp/go-sdk/eolymp/atlas;atlasb\x06proto3"

var (
	file_eolymp_atlas_version_proto_rawDescOnce sync.Once
	file_eolymp_atlas_version_proto_rawDescData []byte
)

func file_eolymp_atlas_version_proto_rawDescGZIP() []byte {
	file_eolymp_atlas_version_proto_rawDescOnce.Do(func() {
		file_eolymp_atlas_version_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_eolymp_atlas_version_proto_rawDesc), len(file_eolymp_atlas_version_proto_rawDesc)))
	})
	return file_eolymp_atlas_version_proto_rawDescData
}

var file_eolymp_atlas_version_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_eolymp_atlas_version_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_eolymp_atlas_version_proto_goTypes = []any{
	(Version_Operation)(0),        // 0: eolymp.atlas.Version.Operation
	(*Version)(nil),               // 1: eolymp.atlas.Version
	(*Version_Change)(nil),        // 2: eolymp.atlas.Version.Change
	(*timestamppb.Timestamp)(nil), // 3: google.protobuf.Timestamp
}
var file_eolymp_atlas_version_proto_depIdxs = []int32{
	3, // 0: eolymp.atlas.Version.created_at:type_name -> google.protobuf.Timestamp
	2, // 1: eolymp.atlas.Version.changes:type_name -> eolymp.atlas.Version.Change
	0, // 2: eolymp.atlas.Version.Change.op:type_name -> eolymp.atlas.Version.Operation
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
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
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_eolymp_atlas_version_proto_rawDesc), len(file_eolymp_atlas_version_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_eolymp_atlas_version_proto_goTypes,
		DependencyIndexes: file_eolymp_atlas_version_proto_depIdxs,
		EnumInfos:         file_eolymp_atlas_version_proto_enumTypes,
		MessageInfos:      file_eolymp_atlas_version_proto_msgTypes,
	}.Build()
	File_eolymp_atlas_version_proto = out.File
	file_eolymp_atlas_version_proto_goTypes = nil
	file_eolymp_atlas_version_proto_depIdxs = nil
}
