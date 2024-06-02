// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v3.21.12
// source: eolymp/community/member_user.proto

package community

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

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Issuer                string                 `protobuf:"bytes,1,opt,name=issuer,proto3" json:"issuer,omitempty"`
	Subject               string                 `protobuf:"bytes,2,opt,name=subject,proto3" json:"subject,omitempty"`
	Nickname              string                 `protobuf:"bytes,10,opt,name=nickname,proto3" json:"nickname,omitempty"`
	NicknameChangeTimeout uint32                 `protobuf:"varint,11,opt,name=nickname_change_timeout,json=nicknameChangeTimeout,proto3" json:"nickname_change_timeout,omitempty"`
	Email                 string                 `protobuf:"bytes,20,opt,name=email,proto3" json:"email,omitempty"`
	EmailVerified         bool                   `protobuf:"varint,21,opt,name=email_verified,json=emailVerified,proto3" json:"email_verified,omitempty"`
	Password              string                 `protobuf:"bytes,30,opt,name=password,proto3" json:"password,omitempty"`
	PasswordAge           uint32                 `protobuf:"varint,31,opt,name=password_age,json=passwordAge,proto3" json:"password_age,omitempty"`
	Name                  string                 `protobuf:"bytes,40,opt,name=name,proto3" json:"name,omitempty"`
	Picture               string                 `protobuf:"bytes,50,opt,name=picture,proto3" json:"picture,omitempty"`
	Birthday              *timestamppb.Timestamp `protobuf:"bytes,60,opt,name=birthday,proto3" json:"birthday,omitempty"`
	Country               string                 `protobuf:"bytes,70,opt,name=country,proto3" json:"country,omitempty"`
	City                  string                 `protobuf:"bytes,73,opt,name=city,proto3" json:"city,omitempty"`
	TeamId                string                 `protobuf:"bytes,100,opt,name=team_id,json=teamId,proto3" json:"team_id,omitempty"`
	Preferences           *User_Preferences      `protobuf:"bytes,800,opt,name=preferences,proto3" json:"preferences,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_community_member_user_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_community_member_user_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_eolymp_community_member_user_proto_rawDescGZIP(), []int{0}
}

func (x *User) GetIssuer() string {
	if x != nil {
		return x.Issuer
	}
	return ""
}

func (x *User) GetSubject() string {
	if x != nil {
		return x.Subject
	}
	return ""
}

func (x *User) GetNickname() string {
	if x != nil {
		return x.Nickname
	}
	return ""
}

func (x *User) GetNicknameChangeTimeout() uint32 {
	if x != nil {
		return x.NicknameChangeTimeout
	}
	return 0
}

func (x *User) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *User) GetEmailVerified() bool {
	if x != nil {
		return x.EmailVerified
	}
	return false
}

func (x *User) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *User) GetPasswordAge() uint32 {
	if x != nil {
		return x.PasswordAge
	}
	return 0
}

func (x *User) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *User) GetPicture() string {
	if x != nil {
		return x.Picture
	}
	return ""
}

func (x *User) GetBirthday() *timestamppb.Timestamp {
	if x != nil {
		return x.Birthday
	}
	return nil
}

func (x *User) GetCountry() string {
	if x != nil {
		return x.Country
	}
	return ""
}

func (x *User) GetCity() string {
	if x != nil {
		return x.City
	}
	return ""
}

func (x *User) GetTeamId() string {
	if x != nil {
		return x.TeamId
	}
	return ""
}

func (x *User) GetPreferences() *User_Preferences {
	if x != nil {
		return x.Preferences
	}
	return nil
}

type User_Preferences struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Locale        string   `protobuf:"bytes,1,opt,name=locale,proto3" json:"locale,omitempty"`
	Timezone      string   `protobuf:"bytes,2,opt,name=timezone,proto3" json:"timezone,omitempty"`
	Runtime       string   `protobuf:"bytes,10,opt,name=runtime,proto3" json:"runtime,omitempty"`
	Notifications []string `protobuf:"bytes,20,rep,name=notifications,proto3" json:"notifications,omitempty"`
}

func (x *User_Preferences) Reset() {
	*x = User_Preferences{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_community_member_user_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User_Preferences) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User_Preferences) ProtoMessage() {}

func (x *User_Preferences) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_community_member_user_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User_Preferences.ProtoReflect.Descriptor instead.
func (*User_Preferences) Descriptor() ([]byte, []int) {
	return file_eolymp_community_member_user_proto_rawDescGZIP(), []int{0, 0}
}

func (x *User_Preferences) GetLocale() string {
	if x != nil {
		return x.Locale
	}
	return ""
}

func (x *User_Preferences) GetTimezone() string {
	if x != nil {
		return x.Timezone
	}
	return ""
}

func (x *User_Preferences) GetRuntime() string {
	if x != nil {
		return x.Runtime
	}
	return ""
}

func (x *User_Preferences) GetNotifications() []string {
	if x != nil {
		return x.Notifications
	}
	return nil
}

var File_eolymp_community_member_user_proto protoreflect.FileDescriptor

var file_eolymp_community_member_user_proto_rawDesc = []byte{
	0x0a, 0x22, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69,
	0x74, 0x79, 0x2f, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x80, 0x05, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72,
	0x12, 0x16, 0x0a, 0x06, 0x69, 0x73, 0x73, 0x75, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x69, 0x73, 0x73, 0x75, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x62, 0x6a,
	0x65, 0x63, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x75, 0x62, 0x6a, 0x65,
	0x63, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x0a,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x36,
	0x0a, 0x17, 0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x5f, 0x63, 0x68, 0x61, 0x6e, 0x67,
	0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x15, 0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x54,
	0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18,
	0x14, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x25, 0x0a, 0x0e,
	0x65, 0x6d, 0x61, 0x69, 0x6c, 0x5f, 0x76, 0x65, 0x72, 0x69, 0x66, 0x69, 0x65, 0x64, 0x18, 0x15,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x0d, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x56, 0x65, 0x72, 0x69, 0x66,
	0x69, 0x65, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18,
	0x1e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12,
	0x21, 0x0a, 0x0c, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x5f, 0x61, 0x67, 0x65, 0x18,
	0x1f, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x41,
	0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x28, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x69, 0x63, 0x74, 0x75, 0x72,
	0x65, 0x18, 0x32, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x69, 0x63, 0x74, 0x75, 0x72, 0x65,
	0x12, 0x36, 0x0a, 0x08, 0x62, 0x69, 0x72, 0x74, 0x68, 0x64, 0x61, 0x79, 0x18, 0x3c, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x08,
	0x62, 0x69, 0x72, 0x74, 0x68, 0x64, 0x61, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x72, 0x79, 0x18, 0x46, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x72, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x69, 0x74, 0x79, 0x18, 0x49, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x63, 0x69, 0x74, 0x79, 0x12, 0x17, 0x0a, 0x07, 0x74, 0x65, 0x61, 0x6d, 0x5f, 0x69,
	0x64, 0x18, 0x64, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x65, 0x61, 0x6d, 0x49, 0x64, 0x12,
	0x45, 0x0a, 0x0b, 0x70, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x18, 0xa0,
	0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x2e, 0x50, 0x72,
	0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x52, 0x0b, 0x70, 0x72, 0x65, 0x66, 0x65,
	0x72, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x1a, 0x81, 0x01, 0x0a, 0x0b, 0x50, 0x72, 0x65, 0x66, 0x65,
	0x72, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x65, 0x12, 0x1a,
	0x0a, 0x08, 0x74, 0x69, 0x6d, 0x65, 0x7a, 0x6f, 0x6e, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x74, 0x69, 0x6d, 0x65, 0x7a, 0x6f, 0x6e, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x75,
	0x6e, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x72, 0x75, 0x6e,
	0x74, 0x69, 0x6d, 0x65, 0x12, 0x24, 0x0a, 0x0d, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x14, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0d, 0x6e, 0x6f, 0x74,
	0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x42, 0x35, 0x5a, 0x33, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f,
	0x67, 0x6f, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x63, 0x6f,
	0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x3b, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74,
	0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_eolymp_community_member_user_proto_rawDescOnce sync.Once
	file_eolymp_community_member_user_proto_rawDescData = file_eolymp_community_member_user_proto_rawDesc
)

func file_eolymp_community_member_user_proto_rawDescGZIP() []byte {
	file_eolymp_community_member_user_proto_rawDescOnce.Do(func() {
		file_eolymp_community_member_user_proto_rawDescData = protoimpl.X.CompressGZIP(file_eolymp_community_member_user_proto_rawDescData)
	})
	return file_eolymp_community_member_user_proto_rawDescData
}

var file_eolymp_community_member_user_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_eolymp_community_member_user_proto_goTypes = []interface{}{
	(*User)(nil),                  // 0: eolymp.community.User
	(*User_Preferences)(nil),      // 1: eolymp.community.User.Preferences
	(*timestamppb.Timestamp)(nil), // 2: google.protobuf.Timestamp
}
var file_eolymp_community_member_user_proto_depIdxs = []int32{
	2, // 0: eolymp.community.User.birthday:type_name -> google.protobuf.Timestamp
	1, // 1: eolymp.community.User.preferences:type_name -> eolymp.community.User.Preferences
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_eolymp_community_member_user_proto_init() }
func file_eolymp_community_member_user_proto_init() {
	if File_eolymp_community_member_user_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_eolymp_community_member_user_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
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
		file_eolymp_community_member_user_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User_Preferences); i {
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
			RawDescriptor: file_eolymp_community_member_user_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_eolymp_community_member_user_proto_goTypes,
		DependencyIndexes: file_eolymp_community_member_user_proto_depIdxs,
		MessageInfos:      file_eolymp_community_member_user_proto_msgTypes,
	}.Build()
	File_eolymp_community_member_user_proto = out.File
	file_eolymp_community_member_user_proto_rawDesc = nil
	file_eolymp_community_member_user_proto_goTypes = nil
	file_eolymp_community_member_user_proto_depIdxs = nil
}
