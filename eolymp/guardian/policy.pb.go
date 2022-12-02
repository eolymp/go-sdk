// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.18.1
// source: eolymp/guardian/policy.proto

package universe

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

type Statement_Effect int32

const (
	Statement_NONE  Statement_Effect = 0
	Statement_DENY  Statement_Effect = 1
	Statement_ALLOW Statement_Effect = 2
)

// Enum value maps for Statement_Effect.
var (
	Statement_Effect_name = map[int32]string{
		0: "NONE",
		1: "DENY",
		2: "ALLOW",
	}
	Statement_Effect_value = map[string]int32{
		"NONE":  0,
		"DENY":  1,
		"ALLOW": 2,
	}
)

func (x Statement_Effect) Enum() *Statement_Effect {
	p := new(Statement_Effect)
	*p = x
	return p
}

func (x Statement_Effect) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Statement_Effect) Descriptor() protoreflect.EnumDescriptor {
	return file_eolymp_guardian_policy_proto_enumTypes[0].Descriptor()
}

func (Statement_Effect) Type() protoreflect.EnumType {
	return &file_eolymp_guardian_policy_proto_enumTypes[0]
}

func (x Statement_Effect) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Statement_Effect.Descriptor instead.
func (Statement_Effect) EnumDescriptor() ([]byte, []int) {
	return file_eolymp_guardian_policy_proto_rawDescGZIP(), []int{1, 0}
}

type Policy struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name       string       `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Statements []*Statement `protobuf:"bytes,2,rep,name=statements,proto3" json:"statements,omitempty"`
}

func (x *Policy) Reset() {
	*x = Policy{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_guardian_policy_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Policy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Policy) ProtoMessage() {}

func (x *Policy) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_guardian_policy_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
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
	return file_eolymp_guardian_policy_proto_rawDescGZIP(), []int{0}
}

func (x *Policy) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Policy) GetStatements() []*Statement {
	if x != nil {
		return x.Statements
	}
	return nil
}

type Statement struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string           `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Effect    Statement_Effect `protobuf:"varint,2,opt,name=effect,proto3,enum=eolymp.universe.Statement_Effect" json:"effect,omitempty"`
	Principal string           `protobuf:"bytes,3,opt,name=principal,proto3" json:"principal,omitempty"`
	Resources []string         `protobuf:"bytes,10,rep,name=resources,proto3" json:"resources,omitempty"`
	Actions   []string         `protobuf:"bytes,20,rep,name=actions,proto3" json:"actions,omitempty"`
}

func (x *Statement) Reset() {
	*x = Statement{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_guardian_policy_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Statement) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Statement) ProtoMessage() {}

func (x *Statement) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_guardian_policy_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Statement.ProtoReflect.Descriptor instead.
func (*Statement) Descriptor() ([]byte, []int) {
	return file_eolymp_guardian_policy_proto_rawDescGZIP(), []int{1}
}

func (x *Statement) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Statement) GetEffect() Statement_Effect {
	if x != nil {
		return x.Effect
	}
	return Statement_NONE
}

func (x *Statement) GetPrincipal() string {
	if x != nil {
		return x.Principal
	}
	return ""
}

func (x *Statement) GetResources() []string {
	if x != nil {
		return x.Resources
	}
	return nil
}

func (x *Statement) GetActions() []string {
	if x != nil {
		return x.Actions
	}
	return nil
}

var File_eolymp_guardian_policy_proto protoreflect.FileDescriptor

var file_eolymp_guardian_policy_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x67, 0x75, 0x61, 0x72, 0x64, 0x69, 0x61,
	0x6e, 0x2f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f,
	0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x75, 0x6e, 0x69, 0x76, 0x65, 0x72, 0x73, 0x65, 0x22,
	0x58, 0x0a, 0x06, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x3a, 0x0a,
	0x0a, 0x73, 0x74, 0x61, 0x74, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x75, 0x6e, 0x69, 0x76, 0x65,
	0x72, 0x73, 0x65, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x0a, 0x73,
	0x74, 0x61, 0x74, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x22, 0xd5, 0x01, 0x0a, 0x09, 0x53, 0x74,
	0x61, 0x74, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x39, 0x0a, 0x06, 0x65, 0x66, 0x66, 0x65, 0x63,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x21, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70,
	0x2e, 0x75, 0x6e, 0x69, 0x76, 0x65, 0x72, 0x73, 0x65, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x6d,
	0x65, 0x6e, 0x74, 0x2e, 0x45, 0x66, 0x66, 0x65, 0x63, 0x74, 0x52, 0x06, 0x65, 0x66, 0x66, 0x65,
	0x63, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x72, 0x69, 0x6e, 0x63, 0x69, 0x70, 0x61, 0x6c, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x69, 0x6e, 0x63, 0x69, 0x70, 0x61, 0x6c,
	0x12, 0x1c, 0x0a, 0x09, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x18, 0x0a, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x09, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x12, 0x18,
	0x0a, 0x07, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x14, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x07, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x27, 0x0a, 0x06, 0x45, 0x66, 0x66, 0x65,
	0x63, 0x74, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04,
	0x44, 0x45, 0x4e, 0x59, 0x10, 0x01, 0x12, 0x09, 0x0a, 0x05, 0x41, 0x4c, 0x4c, 0x4f, 0x57, 0x10,
	0x02, 0x42, 0x33, 0x5a, 0x31, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x67, 0x6f, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x65, 0x6f,
	0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x75, 0x6e, 0x69, 0x76, 0x65, 0x72, 0x73, 0x65, 0x3b, 0x75, 0x6e,
	0x69, 0x76, 0x65, 0x72, 0x73, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_eolymp_guardian_policy_proto_rawDescOnce sync.Once
	file_eolymp_guardian_policy_proto_rawDescData = file_eolymp_guardian_policy_proto_rawDesc
)

func file_eolymp_guardian_policy_proto_rawDescGZIP() []byte {
	file_eolymp_guardian_policy_proto_rawDescOnce.Do(func() {
		file_eolymp_guardian_policy_proto_rawDescData = protoimpl.X.CompressGZIP(file_eolymp_guardian_policy_proto_rawDescData)
	})
	return file_eolymp_guardian_policy_proto_rawDescData
}

var file_eolymp_guardian_policy_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_eolymp_guardian_policy_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_eolymp_guardian_policy_proto_goTypes = []interface{}{
	(Statement_Effect)(0), // 0: eolymp.universe.Statement.Effect
	(*Policy)(nil),        // 1: eolymp.universe.Policy
	(*Statement)(nil),     // 2: eolymp.universe.Statement
}
var file_eolymp_guardian_policy_proto_depIdxs = []int32{
	2, // 0: eolymp.universe.Policy.statements:type_name -> eolymp.universe.Statement
	0, // 1: eolymp.universe.Statement.effect:type_name -> eolymp.universe.Statement.Effect
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_eolymp_guardian_policy_proto_init() }
func file_eolymp_guardian_policy_proto_init() {
	if File_eolymp_guardian_policy_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_eolymp_guardian_policy_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Policy); i {
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
		file_eolymp_guardian_policy_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Statement); i {
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
			RawDescriptor: file_eolymp_guardian_policy_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_eolymp_guardian_policy_proto_goTypes,
		DependencyIndexes: file_eolymp_guardian_policy_proto_depIdxs,
		EnumInfos:         file_eolymp_guardian_policy_proto_enumTypes,
		MessageInfos:      file_eolymp_guardian_policy_proto_msgTypes,
	}.Build()
	File_eolymp_guardian_policy_proto = out.File
	file_eolymp_guardian_policy_proto_rawDesc = nil
	file_eolymp_guardian_policy_proto_goTypes = nil
	file_eolymp_guardian_policy_proto_depIdxs = nil
}
