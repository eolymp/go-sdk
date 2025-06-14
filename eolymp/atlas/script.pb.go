// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.29.4
// source: eolymp/atlas/script.proto

package atlas

import (
	executor "github.com/eolymp/go-sdk/eolymp/executor"
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

type Script_Extra int32

const (
	Script_NO_EXTRA Script_Extra = 0
	Script_SOURCE   Script_Extra = 1
)

// Enum value maps for Script_Extra.
var (
	Script_Extra_name = map[int32]string{
		0: "NO_EXTRA",
		1: "SOURCE",
	}
	Script_Extra_value = map[string]int32{
		"NO_EXTRA": 0,
		"SOURCE":   1,
	}
)

func (x Script_Extra) Enum() *Script_Extra {
	p := new(Script_Extra)
	*p = x
	return p
}

func (x Script_Extra) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Script_Extra) Descriptor() protoreflect.EnumDescriptor {
	return file_eolymp_atlas_script_proto_enumTypes[0].Descriptor()
}

func (Script_Extra) Type() protoreflect.EnumType {
	return &file_eolymp_atlas_script_proto_enumTypes[0]
}

func (x Script_Extra) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Script_Extra.Descriptor instead.
func (Script_Extra) EnumDescriptor() ([]byte, []int) {
	return file_eolymp_atlas_script_proto_rawDescGZIP(), []int{0, 0}
}

type Script struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	Id    string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name  string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// Secret means code and configuration must not be exposed to users
	Secret bool `protobuf:"varint,7,opt,name=secret,proto3" json:"secret,omitempty"`
	// The runtime to execute the script
	Runtime string `protobuf:"bytes,10,opt,name=runtime,proto3" json:"runtime,omitempty"`
	// Source code for script
	Source string `protobuf:"bytes,11,opt,name=source,proto3" json:"source,omitempty"`
	// Additional files placed into workdir during compilation and execution
	Files         []*executor.File `protobuf:"bytes,20,rep,name=files,proto3" json:"files,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Script) Reset() {
	*x = Script{}
	mi := &file_eolymp_atlas_script_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Script) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Script) ProtoMessage() {}

func (x *Script) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_atlas_script_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Script.ProtoReflect.Descriptor instead.
func (*Script) Descriptor() ([]byte, []int) {
	return file_eolymp_atlas_script_proto_rawDescGZIP(), []int{0}
}

func (x *Script) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Script) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Script) GetSecret() bool {
	if x != nil {
		return x.Secret
	}
	return false
}

func (x *Script) GetRuntime() string {
	if x != nil {
		return x.Runtime
	}
	return ""
}

func (x *Script) GetSource() string {
	if x != nil {
		return x.Source
	}
	return ""
}

func (x *Script) GetFiles() []*executor.File {
	if x != nil {
		return x.Files
	}
	return nil
}

var File_eolymp_atlas_script_proto protoreflect.FileDescriptor

const file_eolymp_atlas_script_proto_rawDesc = "" +
	"\n" +
	"\x19eolymp/atlas/script.proto\x12\feolymp.atlas\x1a\x1aeolymp/executor/file.proto\"\xc6\x01\n" +
	"\x06Script\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x12\x12\n" +
	"\x04name\x18\x02 \x01(\tR\x04name\x12\x16\n" +
	"\x06secret\x18\a \x01(\bR\x06secret\x12\x18\n" +
	"\aruntime\x18\n" +
	" \x01(\tR\aruntime\x12\x16\n" +
	"\x06source\x18\v \x01(\tR\x06source\x12+\n" +
	"\x05files\x18\x14 \x03(\v2\x15.eolymp.executor.FileR\x05files\"!\n" +
	"\x05Extra\x12\f\n" +
	"\bNO_EXTRA\x10\x00\x12\n" +
	"\n" +
	"\x06SOURCE\x10\x01B-Z+github.com/eolymp/go-sdk/eolymp/atlas;atlasb\x06proto3"

var (
	file_eolymp_atlas_script_proto_rawDescOnce sync.Once
	file_eolymp_atlas_script_proto_rawDescData []byte
)

func file_eolymp_atlas_script_proto_rawDescGZIP() []byte {
	file_eolymp_atlas_script_proto_rawDescOnce.Do(func() {
		file_eolymp_atlas_script_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_eolymp_atlas_script_proto_rawDesc), len(file_eolymp_atlas_script_proto_rawDesc)))
	})
	return file_eolymp_atlas_script_proto_rawDescData
}

var file_eolymp_atlas_script_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_eolymp_atlas_script_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_eolymp_atlas_script_proto_goTypes = []any{
	(Script_Extra)(0),     // 0: eolymp.atlas.Script.Extra
	(*Script)(nil),        // 1: eolymp.atlas.Script
	(*executor.File)(nil), // 2: eolymp.executor.File
}
var file_eolymp_atlas_script_proto_depIdxs = []int32{
	2, // 0: eolymp.atlas.Script.files:type_name -> eolymp.executor.File
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_eolymp_atlas_script_proto_init() }
func file_eolymp_atlas_script_proto_init() {
	if File_eolymp_atlas_script_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_eolymp_atlas_script_proto_rawDesc), len(file_eolymp_atlas_script_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_eolymp_atlas_script_proto_goTypes,
		DependencyIndexes: file_eolymp_atlas_script_proto_depIdxs,
		EnumInfos:         file_eolymp_atlas_script_proto_enumTypes,
		MessageInfos:      file_eolymp_atlas_script_proto_msgTypes,
	}.Build()
	File_eolymp_atlas_script_proto = out.File
	file_eolymp_atlas_script_proto_goTypes = nil
	file_eolymp_atlas_script_proto_depIdxs = nil
}
