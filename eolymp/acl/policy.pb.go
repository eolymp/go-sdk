// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.4
// 	protoc        v4.24.4
// source: eolymp/acl/policy.proto

package acl

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

type Policy struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Principal     string                 `protobuf:"bytes,3,opt,name=principal,proto3" json:"principal,omitempty"`                           // policy principal (user id)
	Resource      string                 `protobuf:"bytes,10,opt,name=resource,proto3" json:"resource,omitempty"`                            // policy resource url, for example /contests/xyz, empty means policy applies globally
	AllowAll      bool                   `protobuf:"varint,12,opt,name=allow_all,json=allowAll,proto3" json:"allow_all,omitempty"`           // allow all actions
	Allows        []Action               `protobuf:"varint,11,rep,packed,name=allows,proto3,enum=eolymp.acl.Action" json:"allows,omitempty"` // list of allowed actions
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Policy) Reset() {
	*x = Policy{}
	mi := &file_eolymp_acl_policy_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Policy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Policy) ProtoMessage() {}

func (x *Policy) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_acl_policy_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Policy.ProtoReflect.Descriptor instead.
func (*Policy) Descriptor() ([]byte, []int) {
	return file_eolymp_acl_policy_proto_rawDescGZIP(), []int{0}
}

func (x *Policy) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Policy) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Policy) GetPrincipal() string {
	if x != nil {
		return x.Principal
	}
	return ""
}

func (x *Policy) GetResource() string {
	if x != nil {
		return x.Resource
	}
	return ""
}

func (x *Policy) GetAllowAll() bool {
	if x != nil {
		return x.AllowAll
	}
	return false
}

func (x *Policy) GetAllows() []Action {
	if x != nil {
		return x.Allows
	}
	return nil
}

var File_eolymp_acl_policy_proto protoreflect.FileDescriptor

var file_eolymp_acl_policy_proto_rawDesc = string([]byte{
	0x0a, 0x17, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x61, 0x63, 0x6c, 0x2f, 0x70, 0x6f, 0x6c,
	0x69, 0x63, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x65, 0x6f, 0x6c, 0x79, 0x6d,
	0x70, 0x2e, 0x61, 0x63, 0x6c, 0x1a, 0x17, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x61, 0x63,
	0x6c, 0x2f, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xaf,
	0x01, 0x0a, 0x06, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a,
	0x09, 0x70, 0x72, 0x69, 0x6e, 0x63, 0x69, 0x70, 0x61, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x70, 0x72, 0x69, 0x6e, 0x63, 0x69, 0x70, 0x61, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x61, 0x6c, 0x6c, 0x6f, 0x77,
	0x5f, 0x61, 0x6c, 0x6c, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x61, 0x6c, 0x6c, 0x6f,
	0x77, 0x41, 0x6c, 0x6c, 0x12, 0x2a, 0x0a, 0x06, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x73, 0x18, 0x0b,
	0x20, 0x03, 0x28, 0x0e, 0x32, 0x12, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x61, 0x63,
	0x6c, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x06, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x73,
	0x42, 0x29, 0x5a, 0x27, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65,
	0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x67, 0x6f, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x65, 0x6f, 0x6c,
	0x79, 0x6d, 0x70, 0x2f, 0x61, 0x63, 0x6c, 0x3b, 0x61, 0x63, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
})

var (
	file_eolymp_acl_policy_proto_rawDescOnce sync.Once
	file_eolymp_acl_policy_proto_rawDescData []byte
)

func file_eolymp_acl_policy_proto_rawDescGZIP() []byte {
	file_eolymp_acl_policy_proto_rawDescOnce.Do(func() {
		file_eolymp_acl_policy_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_eolymp_acl_policy_proto_rawDesc), len(file_eolymp_acl_policy_proto_rawDesc)))
	})
	return file_eolymp_acl_policy_proto_rawDescData
}

var file_eolymp_acl_policy_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_eolymp_acl_policy_proto_goTypes = []any{
	(*Policy)(nil), // 0: eolymp.acl.Policy
	(Action)(0),    // 1: eolymp.acl.Action
}
var file_eolymp_acl_policy_proto_depIdxs = []int32{
	1, // 0: eolymp.acl.Policy.allows:type_name -> eolymp.acl.Action
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_eolymp_acl_policy_proto_init() }
func file_eolymp_acl_policy_proto_init() {
	if File_eolymp_acl_policy_proto != nil {
		return
	}
	file_eolymp_acl_action_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_eolymp_acl_policy_proto_rawDesc), len(file_eolymp_acl_policy_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_eolymp_acl_policy_proto_goTypes,
		DependencyIndexes: file_eolymp_acl_policy_proto_depIdxs,
		MessageInfos:      file_eolymp_acl_policy_proto_msgTypes,
	}.Build()
	File_eolymp_acl_policy_proto = out.File
	file_eolymp_acl_policy_proto_goTypes = nil
	file_eolymp_acl_policy_proto_depIdxs = nil
}
