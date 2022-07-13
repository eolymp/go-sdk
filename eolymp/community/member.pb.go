// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.12.2
// source: eolymp/community/member.proto

package community

import (
	_ "github.com/eolymp/go-sdk/eolymp/annotations"
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

type Member_Status int32

const (
	Member_UNKNOWN_STATUS Member_Status = 0
	Member_UNSTAFFED      Member_Status = 1
	Member_UNREGISTERED   Member_Status = 2
	Member_ACTIVE         Member_Status = 3
	Member_INACTIVE       Member_Status = 4
	Member_GHOST          Member_Status = 5
)

// Enum value maps for Member_Status.
var (
	Member_Status_name = map[int32]string{
		0: "UNKNOWN_STATUS",
		1: "UNSTAFFED",
		2: "UNREGISTERED",
		3: "ACTIVE",
		4: "INACTIVE",
		5: "GHOST",
	}
	Member_Status_value = map[string]int32{
		"UNKNOWN_STATUS": 0,
		"UNSTAFFED":      1,
		"UNREGISTERED":   2,
		"ACTIVE":         3,
		"INACTIVE":       4,
		"GHOST":          5,
	}
)

func (x Member_Status) Enum() *Member_Status {
	p := new(Member_Status)
	*p = x
	return p
}

func (x Member_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Member_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_eolymp_community_member_proto_enumTypes[0].Descriptor()
}

func (Member_Status) Type() protoreflect.EnumType {
	return &file_eolymp_community_member_proto_enumTypes[0]
}

func (x Member_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Member_Status.Descriptor instead.
func (Member_Status) EnumDescriptor() ([]byte, []int) {
	return file_eolymp_community_member_proto_rawDescGZIP(), []int{0, 0}
}

type Member struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         string          `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Uri        string          `protobuf:"bytes,9999,opt,name=uri,proto3" json:"uri,omitempty"`
	Name       string          `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Disabled   bool            `protobuf:"varint,3,opt,name=disabled,proto3" json:"disabled,omitempty"`
	Registered bool            `protobuf:"varint,4,opt,name=registered,proto3" json:"registered,omitempty"`
	Staffed    bool            `protobuf:"varint,5,opt,name=staffed,proto3" json:"staffed,omitempty"`
	Ghost      bool            `protobuf:"varint,7,opt,name=ghost,proto3" json:"ghost,omitempty"`
	Status     Member_Status   `protobuf:"varint,6,opt,name=status,proto3,enum=eolymp.community.Member_Status" json:"status,omitempty"`
	Users      []*Member_User  `protobuf:"bytes,10,rep,name=users,proto3" json:"users,omitempty"`
	Values     []*Member_Value `protobuf:"bytes,20,rep,name=values,proto3" json:"values,omitempty"`
}

func (x *Member) Reset() {
	*x = Member{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_community_member_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Member) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Member) ProtoMessage() {}

func (x *Member) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_community_member_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Member.ProtoReflect.Descriptor instead.
func (*Member) Descriptor() ([]byte, []int) {
	return file_eolymp_community_member_proto_rawDescGZIP(), []int{0}
}

func (x *Member) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Member) GetUri() string {
	if x != nil {
		return x.Uri
	}
	return ""
}

func (x *Member) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Member) GetDisabled() bool {
	if x != nil {
		return x.Disabled
	}
	return false
}

func (x *Member) GetRegistered() bool {
	if x != nil {
		return x.Registered
	}
	return false
}

func (x *Member) GetStaffed() bool {
	if x != nil {
		return x.Staffed
	}
	return false
}

func (x *Member) GetGhost() bool {
	if x != nil {
		return x.Ghost
	}
	return false
}

func (x *Member) GetStatus() Member_Status {
	if x != nil {
		return x.Status
	}
	return Member_UNKNOWN_STATUS
}

func (x *Member) GetUsers() []*Member_User {
	if x != nil {
		return x.Users
	}
	return nil
}

func (x *Member) GetValues() []*Member_Value {
	if x != nil {
		return x.Values
	}
	return nil
}

type Member_User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *Member_User) Reset() {
	*x = Member_User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_community_member_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Member_User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Member_User) ProtoMessage() {}

func (x *Member_User) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_community_member_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Member_User.ProtoReflect.Descriptor instead.
func (*Member_User) Descriptor() ([]byte, []int) {
	return file_eolymp_community_member_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Member_User) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type Member_Value struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AttributeKey string `protobuf:"bytes,1,opt,name=attribute_key,json=attributeKey,proto3" json:"attribute_key,omitempty"`
	ValueString  string `protobuf:"bytes,2,opt,name=value_string,json=valueString,proto3" json:"value_string,omitempty"`
	ValueNumber  int32  `protobuf:"varint,3,opt,name=value_number,json=valueNumber,proto3" json:"value_number,omitempty"`
}

func (x *Member_Value) Reset() {
	*x = Member_Value{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_community_member_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Member_Value) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Member_Value) ProtoMessage() {}

func (x *Member_Value) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_community_member_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Member_Value.ProtoReflect.Descriptor instead.
func (*Member_Value) Descriptor() ([]byte, []int) {
	return file_eolymp_community_member_proto_rawDescGZIP(), []int{0, 1}
}

func (x *Member_Value) GetAttributeKey() string {
	if x != nil {
		return x.AttributeKey
	}
	return ""
}

func (x *Member_Value) GetValueString() string {
	if x != nil {
		return x.ValueString
	}
	return ""
}

func (x *Member_Value) GetValueNumber() int32 {
	if x != nil {
		return x.ValueNumber
	}
	return 0
}

var File_eolymp_community_member_proto protoreflect.FileDescriptor

var file_eolymp_community_member_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69,
	0x74, 0x79, 0x2f, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x10, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74,
	0x79, 0x1a, 0x21, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf3, 0x04, 0x0a, 0x06, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x11, 0x0a, 0x03, 0x75, 0x72, 0x69, 0x18, 0x8f, 0x4e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75,
	0x72, 0x69, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c,
	0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c,
	0x65, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x65, 0x64,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72,
	0x65, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x74, 0x61, 0x66, 0x66, 0x65, 0x64, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x74, 0x61, 0x66, 0x66, 0x65, 0x64, 0x12, 0x14, 0x0a, 0x05,
	0x67, 0x68, 0x6f, 0x73, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x67, 0x68, 0x6f,
	0x73, 0x74, 0x12, 0x37, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x1f, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x75, 0x6e, 0x69, 0x74, 0x79, 0x2e, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x2e, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x33, 0x0a, 0x05, 0x75,
	0x73, 0x65, 0x72, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x65, 0x6f, 0x6c,
	0x79, 0x6d, 0x70, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x2e, 0x4d, 0x65,
	0x6d, 0x62, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73,
	0x12, 0x36, 0x0a, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x14, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x1e, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e,
	0x69, 0x74, 0x79, 0x2e, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x2e, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x52, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x1a, 0x1f, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72,
	0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x1a, 0x72, 0x0a, 0x05, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x5f,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x61, 0x74, 0x74, 0x72, 0x69,
	0x62, 0x75, 0x74, 0x65, 0x4b, 0x65, 0x79, 0x12, 0x21, 0x0a, 0x0c, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x5f, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x12, 0x21, 0x0a, 0x0c, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0b, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x22, 0x62, 0x0a,
	0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x12, 0x0a, 0x0e, 0x55, 0x4e, 0x4b, 0x4e, 0x4f,
	0x57, 0x4e, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x10, 0x00, 0x12, 0x0d, 0x0a, 0x09, 0x55,
	0x4e, 0x53, 0x54, 0x41, 0x46, 0x46, 0x45, 0x44, 0x10, 0x01, 0x12, 0x10, 0x0a, 0x0c, 0x55, 0x4e,
	0x52, 0x45, 0x47, 0x49, 0x53, 0x54, 0x45, 0x52, 0x45, 0x44, 0x10, 0x02, 0x12, 0x0a, 0x0a, 0x06,
	0x41, 0x43, 0x54, 0x49, 0x56, 0x45, 0x10, 0x03, 0x12, 0x0c, 0x0a, 0x08, 0x49, 0x4e, 0x41, 0x43,
	0x54, 0x49, 0x56, 0x45, 0x10, 0x04, 0x12, 0x09, 0x0a, 0x05, 0x47, 0x48, 0x4f, 0x53, 0x54, 0x10,
	0x05, 0x3a, 0x27, 0xb2, 0xe3, 0x0a, 0x23, 0xba, 0xe3, 0x0a, 0x06, 0x6d, 0x65, 0x6d, 0x62, 0x65,
	0x72, 0xc2, 0xe3, 0x0a, 0x15, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x75, 0x6e, 0x69, 0x76,
	0x65, 0x72, 0x73, 0x65, 0x2e, 0x53, 0x70, 0x61, 0x63, 0x65, 0x42, 0x35, 0x5a, 0x33, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f,
	0x67, 0x6f, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x63, 0x6f,
	0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x3b, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74,
	0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_eolymp_community_member_proto_rawDescOnce sync.Once
	file_eolymp_community_member_proto_rawDescData = file_eolymp_community_member_proto_rawDesc
)

func file_eolymp_community_member_proto_rawDescGZIP() []byte {
	file_eolymp_community_member_proto_rawDescOnce.Do(func() {
		file_eolymp_community_member_proto_rawDescData = protoimpl.X.CompressGZIP(file_eolymp_community_member_proto_rawDescData)
	})
	return file_eolymp_community_member_proto_rawDescData
}

var file_eolymp_community_member_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_eolymp_community_member_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_eolymp_community_member_proto_goTypes = []interface{}{
	(Member_Status)(0),   // 0: eolymp.community.Member.Status
	(*Member)(nil),       // 1: eolymp.community.Member
	(*Member_User)(nil),  // 2: eolymp.community.Member.User
	(*Member_Value)(nil), // 3: eolymp.community.Member.Value
}
var file_eolymp_community_member_proto_depIdxs = []int32{
	0, // 0: eolymp.community.Member.status:type_name -> eolymp.community.Member.Status
	2, // 1: eolymp.community.Member.users:type_name -> eolymp.community.Member.User
	3, // 2: eolymp.community.Member.values:type_name -> eolymp.community.Member.Value
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_eolymp_community_member_proto_init() }
func file_eolymp_community_member_proto_init() {
	if File_eolymp_community_member_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_eolymp_community_member_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Member); i {
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
		file_eolymp_community_member_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Member_User); i {
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
		file_eolymp_community_member_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Member_Value); i {
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
			RawDescriptor: file_eolymp_community_member_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_eolymp_community_member_proto_goTypes,
		DependencyIndexes: file_eolymp_community_member_proto_depIdxs,
		EnumInfos:         file_eolymp_community_member_proto_enumTypes,
		MessageInfos:      file_eolymp_community_member_proto_msgTypes,
	}.Build()
	File_eolymp_community_member_proto = out.File
	file_eolymp_community_member_proto_rawDesc = nil
	file_eolymp_community_member_proto_goTypes = nil
	file_eolymp_community_member_proto_depIdxs = nil
}
