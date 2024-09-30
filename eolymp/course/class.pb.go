// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.21.12
// source: eolymp/course/class.proto

package course

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

type Class struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	GroupId     string `protobuf:"bytes,3,opt,name=group_id,json=groupId,proto3" json:"group_id,omitempty"`
	ModuleCount uint32 `protobuf:"varint,6,opt,name=module_count,json=moduleCount,proto3" json:"module_count,omitempty"` // number of assigned modules (in all statuses)
}

func (x *Class) Reset() {
	*x = Class{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_course_class_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Class) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Class) ProtoMessage() {}

func (x *Class) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_course_class_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Class.ProtoReflect.Descriptor instead.
func (*Class) Descriptor() ([]byte, []int) {
	return file_eolymp_course_class_proto_rawDescGZIP(), []int{0}
}

func (x *Class) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Class) GetGroupId() string {
	if x != nil {
		return x.GroupId
	}
	return ""
}

func (x *Class) GetModuleCount() uint32 {
	if x != nil {
		return x.ModuleCount
	}
	return 0
}

var File_eolymp_course_class_proto protoreflect.FileDescriptor

var file_eolymp_course_class_proto_rawDesc = []byte{
	0x0a, 0x19, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x2f,
	0x63, 0x6c, 0x61, 0x73, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x65, 0x6f, 0x6c,
	0x79, 0x6d, 0x70, 0x2e, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x22, 0x55, 0x0a, 0x05, 0x43, 0x6c,
	0x61, 0x73, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x69, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x12, 0x21,
	0x0a, 0x0c, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x43, 0x6f, 0x75, 0x6e,
	0x74, 0x42, 0x2f, 0x5a, 0x2d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x67, 0x6f, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x65, 0x6f,
	0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x3b, 0x63, 0x6f, 0x75, 0x72,
	0x73, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_eolymp_course_class_proto_rawDescOnce sync.Once
	file_eolymp_course_class_proto_rawDescData = file_eolymp_course_class_proto_rawDesc
)

func file_eolymp_course_class_proto_rawDescGZIP() []byte {
	file_eolymp_course_class_proto_rawDescOnce.Do(func() {
		file_eolymp_course_class_proto_rawDescData = protoimpl.X.CompressGZIP(file_eolymp_course_class_proto_rawDescData)
	})
	return file_eolymp_course_class_proto_rawDescData
}

var file_eolymp_course_class_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_eolymp_course_class_proto_goTypes = []any{
	(*Class)(nil), // 0: eolymp.course.Class
}
var file_eolymp_course_class_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_eolymp_course_class_proto_init() }
func file_eolymp_course_class_proto_init() {
	if File_eolymp_course_class_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_eolymp_course_class_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*Class); i {
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
			RawDescriptor: file_eolymp_course_class_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_eolymp_course_class_proto_goTypes,
		DependencyIndexes: file_eolymp_course_class_proto_depIdxs,
		MessageInfos:      file_eolymp_course_class_proto_msgTypes,
	}.Build()
	File_eolymp_course_class_proto = out.File
	file_eolymp_course_class_proto_rawDesc = nil
	file_eolymp_course_class_proto_goTypes = nil
	file_eolymp_course_class_proto_depIdxs = nil
}
