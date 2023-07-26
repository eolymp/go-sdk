// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.18.1
// source: eolymp/community/member.proto

package community

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

type Member struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name       string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"` // readonly, users nickname, ghosts name or teams name
	Active     bool   `protobuf:"varint,10,opt,name=active,proto3" json:"active,omitempty"`
	Incomplete bool   `protobuf:"varint,20,opt,name=incomplete,proto3" json:"incomplete,omitempty"` // member profile (attributes) is missing some information
	Unofficial bool   `protobuf:"varint,30,opt,name=unofficial,proto3" json:"unofficial,omitempty"` // member participates in all competitions unofficially
	Secret     bool   `protobuf:"varint,40,opt,name=secret,proto3" json:"secret,omitempty"`         // member is secret and does not appear on anywhere (for example, an admin who performs testing)
	// Types that are assignable to Account:
	//	*Member_User
	//	*Member_Team
	//	*Member_Ghost
	Account    isMember_Account   `protobuf_oneof:"account"`
	Attributes []*Attribute_Value `protobuf:"bytes,900,rep,name=attributes,proto3" json:"attributes,omitempty"`
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

func (x *Member) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Member) GetActive() bool {
	if x != nil {
		return x.Active
	}
	return false
}

func (x *Member) GetIncomplete() bool {
	if x != nil {
		return x.Incomplete
	}
	return false
}

func (x *Member) GetUnofficial() bool {
	if x != nil {
		return x.Unofficial
	}
	return false
}

func (x *Member) GetSecret() bool {
	if x != nil {
		return x.Secret
	}
	return false
}

func (m *Member) GetAccount() isMember_Account {
	if m != nil {
		return m.Account
	}
	return nil
}

func (x *Member) GetUser() *User {
	if x, ok := x.GetAccount().(*Member_User); ok {
		return x.User
	}
	return nil
}

func (x *Member) GetTeam() *Team {
	if x, ok := x.GetAccount().(*Member_Team); ok {
		return x.Team
	}
	return nil
}

func (x *Member) GetGhost() *Ghost {
	if x, ok := x.GetAccount().(*Member_Ghost); ok {
		return x.Ghost
	}
	return nil
}

func (x *Member) GetAttributes() []*Attribute_Value {
	if x != nil {
		return x.Attributes
	}
	return nil
}

type isMember_Account interface {
	isMember_Account()
}

type Member_User struct {
	User *User `protobuf:"bytes,100,opt,name=user,proto3,oneof"`
}

type Member_Team struct {
	Team *Team `protobuf:"bytes,101,opt,name=team,proto3,oneof"`
}

type Member_Ghost struct {
	Ghost *Ghost `protobuf:"bytes,102,opt,name=ghost,proto3,oneof"`
}

func (*Member_User) isMember_Account() {}

func (*Member_Team) isMember_Account() {}

func (*Member_Ghost) isMember_Account() {}

var File_eolymp_community_member_proto protoreflect.FileDescriptor

var file_eolymp_community_member_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69,
	0x74, 0x79, 0x2f, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x10, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74,
	0x79, 0x1a, 0x20, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e,
	0x69, 0x74, 0x79, 0x2f, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x23, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x63, 0x6f, 0x6d, 0x6d,
	0x75, 0x6e, 0x69, 0x74, 0x79, 0x2f, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x5f, 0x67, 0x68, 0x6f,
	0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x22, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70,
	0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x2f, 0x6d, 0x65, 0x6d, 0x62, 0x65,
	0x72, 0x5f, 0x74, 0x65, 0x61, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x22, 0x65, 0x6f,
	0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x2f, 0x6d,
	0x65, 0x6d, 0x62, 0x65, 0x72, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xf8, 0x02, 0x0a, 0x06, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x06, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x69, 0x6e, 0x63, 0x6f, 0x6d,
	0x70, 0x6c, 0x65, 0x74, 0x65, 0x18, 0x14, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x69, 0x6e, 0x63,
	0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x75, 0x6e, 0x6f, 0x66, 0x66,
	0x69, 0x63, 0x69, 0x61, 0x6c, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x75, 0x6e, 0x6f,
	0x66, 0x66, 0x69, 0x63, 0x69, 0x61, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x63, 0x72, 0x65,
	0x74, 0x18, 0x28, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x12,
	0x2c, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x64, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e,
	0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79,
	0x2e, 0x55, 0x73, 0x65, 0x72, 0x48, 0x00, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x2c, 0x0a,
	0x04, 0x74, 0x65, 0x61, 0x6d, 0x18, 0x65, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x65, 0x6f,
	0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x2e, 0x54,
	0x65, 0x61, 0x6d, 0x48, 0x00, 0x52, 0x04, 0x74, 0x65, 0x61, 0x6d, 0x12, 0x2f, 0x0a, 0x05, 0x67,
	0x68, 0x6f, 0x73, 0x74, 0x18, 0x66, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x65, 0x6f, 0x6c,
	0x79, 0x6d, 0x70, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x2e, 0x47, 0x68,
	0x6f, 0x73, 0x74, 0x48, 0x00, 0x52, 0x05, 0x67, 0x68, 0x6f, 0x73, 0x74, 0x12, 0x42, 0x0a, 0x0a,
	0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x18, 0x84, 0x07, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x21, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x75,
	0x6e, 0x69, 0x74, 0x79, 0x2e, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x2e, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x52, 0x0a, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73,
	0x42, 0x09, 0x0a, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x42, 0x35, 0x5a, 0x33, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70,
	0x2f, 0x67, 0x6f, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x63,
	0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x3b, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69,
	0x74, 0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
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

var file_eolymp_community_member_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_eolymp_community_member_proto_goTypes = []interface{}{
	(*Member)(nil),          // 0: eolymp.community.Member
	(*User)(nil),            // 1: eolymp.community.User
	(*Team)(nil),            // 2: eolymp.community.Team
	(*Ghost)(nil),           // 3: eolymp.community.Ghost
	(*Attribute_Value)(nil), // 4: eolymp.community.Attribute.Value
}
var file_eolymp_community_member_proto_depIdxs = []int32{
	1, // 0: eolymp.community.Member.user:type_name -> eolymp.community.User
	2, // 1: eolymp.community.Member.team:type_name -> eolymp.community.Team
	3, // 2: eolymp.community.Member.ghost:type_name -> eolymp.community.Ghost
	4, // 3: eolymp.community.Member.attributes:type_name -> eolymp.community.Attribute.Value
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_eolymp_community_member_proto_init() }
func file_eolymp_community_member_proto_init() {
	if File_eolymp_community_member_proto != nil {
		return
	}
	file_eolymp_community_attribute_proto_init()
	file_eolymp_community_member_ghost_proto_init()
	file_eolymp_community_member_team_proto_init()
	file_eolymp_community_member_user_proto_init()
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
	}
	file_eolymp_community_member_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*Member_User)(nil),
		(*Member_Team)(nil),
		(*Member_Ghost)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_eolymp_community_member_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_eolymp_community_member_proto_goTypes,
		DependencyIndexes: file_eolymp_community_member_proto_depIdxs,
		MessageInfos:      file_eolymp_community_member_proto_msgTypes,
	}.Build()
	File_eolymp_community_member_proto = out.File
	file_eolymp_community_member_proto_rawDesc = nil
	file_eolymp_community_member_proto_goTypes = nil
	file_eolymp_community_member_proto_depIdxs = nil
}
