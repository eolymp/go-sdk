// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.29.0
// 	protoc        v3.18.1
// source: eolymp/atlas/permission.proto

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

type Permission_Role int32

const (
	Permission_NONE   Permission_Role = 0
	Permission_OWNER  Permission_Role = 1
	Permission_EDITOR Permission_Role = 2
)

// Enum value maps for Permission_Role.
var (
	Permission_Role_name = map[int32]string{
		0: "NONE",
		1: "OWNER",
		2: "EDITOR",
	}
	Permission_Role_value = map[string]int32{
		"NONE":   0,
		"OWNER":  1,
		"EDITOR": 2,
	}
)

func (x Permission_Role) Enum() *Permission_Role {
	p := new(Permission_Role)
	*p = x
	return p
}

func (x Permission_Role) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Permission_Role) Descriptor() protoreflect.EnumDescriptor {
	return file_eolymp_atlas_permission_proto_enumTypes[0].Descriptor()
}

func (Permission_Role) Type() protoreflect.EnumType {
	return &file_eolymp_atlas_permission_proto_enumTypes[0]
}

func (x Permission_Role) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Permission_Role.Descriptor instead.
func (Permission_Role) EnumDescriptor() ([]byte, []int) {
	return file_eolymp_atlas_permission_proto_rawDescGZIP(), []int{0, 0}
}

type Permission struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProblemId string          `protobuf:"bytes,1,opt,name=problem_id,json=problemId,proto3" json:"problem_id,omitempty"`
	UserId    string          `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Role      Permission_Role `protobuf:"varint,3,opt,name=role,proto3,enum=eolymp.atlas.Permission_Role" json:"role,omitempty"`
}

func (x *Permission) Reset() {
	*x = Permission{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_atlas_permission_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Permission) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Permission) ProtoMessage() {}

func (x *Permission) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_atlas_permission_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Permission.ProtoReflect.Descriptor instead.
func (*Permission) Descriptor() ([]byte, []int) {
	return file_eolymp_atlas_permission_proto_rawDescGZIP(), []int{0}
}

func (x *Permission) GetProblemId() string {
	if x != nil {
		return x.ProblemId
	}
	return ""
}

func (x *Permission) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *Permission) GetRole() Permission_Role {
	if x != nil {
		return x.Role
	}
	return Permission_NONE
}

var File_eolymp_atlas_permission_proto protoreflect.FileDescriptor

var file_eolymp_atlas_permission_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x61, 0x74, 0x6c, 0x61, 0x73, 0x2f, 0x70,
	0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0c, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x61, 0x74, 0x6c, 0x61, 0x73, 0x22, 0xa0, 0x01,
	0x0a, 0x0a, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1d, 0x0a, 0x0a,
	0x70, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x70, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75,
	0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x12, 0x31, 0x0a, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x1d, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x61, 0x74, 0x6c, 0x61,
	0x73, 0x2e, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x52, 0x6f, 0x6c,
	0x65, 0x52, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x22, 0x27, 0x0a, 0x04, 0x52, 0x6f, 0x6c, 0x65, 0x12,
	0x08, 0x0a, 0x04, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x4f, 0x57, 0x4e,
	0x45, 0x52, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x45, 0x44, 0x49, 0x54, 0x4f, 0x52, 0x10, 0x02,
	0x42, 0x2d, 0x5a, 0x2b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65,
	0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x67, 0x6f, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x65, 0x6f, 0x6c,
	0x79, 0x6d, 0x70, 0x2f, 0x61, 0x74, 0x6c, 0x61, 0x73, 0x3b, 0x61, 0x74, 0x6c, 0x61, 0x73, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_eolymp_atlas_permission_proto_rawDescOnce sync.Once
	file_eolymp_atlas_permission_proto_rawDescData = file_eolymp_atlas_permission_proto_rawDesc
)

func file_eolymp_atlas_permission_proto_rawDescGZIP() []byte {
	file_eolymp_atlas_permission_proto_rawDescOnce.Do(func() {
		file_eolymp_atlas_permission_proto_rawDescData = protoimpl.X.CompressGZIP(file_eolymp_atlas_permission_proto_rawDescData)
	})
	return file_eolymp_atlas_permission_proto_rawDescData
}

var file_eolymp_atlas_permission_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_eolymp_atlas_permission_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_eolymp_atlas_permission_proto_goTypes = []interface{}{
	(Permission_Role)(0), // 0: eolymp.atlas.Permission.Role
	(*Permission)(nil),   // 1: eolymp.atlas.Permission
}
var file_eolymp_atlas_permission_proto_depIdxs = []int32{
	0, // 0: eolymp.atlas.Permission.role:type_name -> eolymp.atlas.Permission.Role
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_eolymp_atlas_permission_proto_init() }
func file_eolymp_atlas_permission_proto_init() {
	if File_eolymp_atlas_permission_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_eolymp_atlas_permission_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Permission); i {
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
			RawDescriptor: file_eolymp_atlas_permission_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_eolymp_atlas_permission_proto_goTypes,
		DependencyIndexes: file_eolymp_atlas_permission_proto_depIdxs,
		EnumInfos:         file_eolymp_atlas_permission_proto_enumTypes,
		MessageInfos:      file_eolymp_atlas_permission_proto_msgTypes,
	}.Build()
	File_eolymp_atlas_permission_proto = out.File
	file_eolymp_atlas_permission_proto_rawDesc = nil
	file_eolymp_atlas_permission_proto_goTypes = nil
	file_eolymp_atlas_permission_proto_depIdxs = nil
}
