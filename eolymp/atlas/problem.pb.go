// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.12.2
// source: eolymp/atlas/problem.proto

package atlas

import (
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

type Problem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Unique identifier.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Index in the public problem catalog.
	Index int32 `protobuf:"varint,10,opt,name=index,proto3" json:"index,omitempty"`
	// Problem is visible to users in public catalog.
	Visible bool `protobuf:"varint,11,opt,name=visible,proto3" json:"visible,omitempty"`
	// Problem is only accessible to users who were specifically granted access to it.
	Private bool `protobuf:"varint,12,opt,name=private,proto3" json:"private,omitempty"`
}

func (x *Problem) Reset() {
	*x = Problem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_atlas_problem_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Problem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Problem) ProtoMessage() {}

func (x *Problem) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_atlas_problem_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Problem.ProtoReflect.Descriptor instead.
func (*Problem) Descriptor() ([]byte, []int) {
	return file_eolymp_atlas_problem_proto_rawDescGZIP(), []int{0}
}

func (x *Problem) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Problem) GetIndex() int32 {
	if x != nil {
		return x.Index
	}
	return 0
}

func (x *Problem) GetVisible() bool {
	if x != nil {
		return x.Visible
	}
	return false
}

func (x *Problem) GetPrivate() bool {
	if x != nil {
		return x.Private
	}
	return false
}

var File_eolymp_atlas_problem_proto protoreflect.FileDescriptor

var file_eolymp_atlas_problem_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x61, 0x74, 0x6c, 0x61, 0x73, 0x2f, 0x70,
	0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x65, 0x6f,
	0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x61, 0x74, 0x6c, 0x61, 0x73, 0x22, 0x63, 0x0a, 0x07, 0x50, 0x72,
	0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x0a,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x18, 0x0a, 0x07, 0x76,
	0x69, 0x73, 0x69, 0x62, 0x6c, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x76, 0x69,
	0x73, 0x69, 0x62, 0x6c, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65,
	0x18, 0x0c, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x42,
	0x33, 0x5a, 0x31, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6f,
	0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x73, 0x2f, 0x67,
	0x6f, 0x2f, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x61, 0x74, 0x6c, 0x61, 0x73, 0x3b, 0x61,
	0x74, 0x6c, 0x61, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_eolymp_atlas_problem_proto_rawDescOnce sync.Once
	file_eolymp_atlas_problem_proto_rawDescData = file_eolymp_atlas_problem_proto_rawDesc
)

func file_eolymp_atlas_problem_proto_rawDescGZIP() []byte {
	file_eolymp_atlas_problem_proto_rawDescOnce.Do(func() {
		file_eolymp_atlas_problem_proto_rawDescData = protoimpl.X.CompressGZIP(file_eolymp_atlas_problem_proto_rawDescData)
	})
	return file_eolymp_atlas_problem_proto_rawDescData
}

var file_eolymp_atlas_problem_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_eolymp_atlas_problem_proto_goTypes = []interface{}{
	(*Problem)(nil), // 0: eolymp.atlas.Problem
}
var file_eolymp_atlas_problem_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_eolymp_atlas_problem_proto_init() }
func file_eolymp_atlas_problem_proto_init() {
	if File_eolymp_atlas_problem_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_eolymp_atlas_problem_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Problem); i {
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
			RawDescriptor: file_eolymp_atlas_problem_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_eolymp_atlas_problem_proto_goTypes,
		DependencyIndexes: file_eolymp_atlas_problem_proto_depIdxs,
		MessageInfos:      file_eolymp_atlas_problem_proto_msgTypes,
	}.Build()
	File_eolymp_atlas_problem_proto = out.File
	file_eolymp_atlas_problem_proto_rawDesc = nil
	file_eolymp_atlas_problem_proto_goTypes = nil
	file_eolymp_atlas_problem_proto_depIdxs = nil
}
