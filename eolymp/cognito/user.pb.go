// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.2
// source: eolymp/cognito/user.proto

package cognito

import (
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type User_RankTrend int32

const (
	User_NONE User_RankTrend = 0
	User_UP   User_RankTrend = 1
	User_DOWN User_RankTrend = 2
)

// Enum value maps for User_RankTrend.
var (
	User_RankTrend_name = map[int32]string{
		0: "NONE",
		1: "UP",
		2: "DOWN",
	}
	User_RankTrend_value = map[string]int32{
		"NONE": 0,
		"UP":   1,
		"DOWN": 2,
	}
)

func (x User_RankTrend) Enum() *User_RankTrend {
	p := new(User_RankTrend)
	*p = x
	return p
}

func (x User_RankTrend) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (User_RankTrend) Descriptor() protoreflect.EnumDescriptor {
	return file_eolymp_cognito_user_proto_enumTypes[0].Descriptor()
}

func (User_RankTrend) Type() protoreflect.EnumType {
	return &file_eolymp_cognito_user_proto_enumTypes[0]
}

func (x User_RankTrend) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use User_RankTrend.Descriptor instead.
func (User_RankTrend) EnumDescriptor() ([]byte, []int) {
	return file_eolymp_cognito_user_proto_rawDescGZIP(), []int{0, 0}
}

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           string               `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`                                                                     // Unique identifier
	Username     string               `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`                                                         // Username (handler)
	Email        string               `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`                                                               // Email address (requires VIEW_PRIVATE_DATA entitlement)
	Active       bool                 `protobuf:"varint,4,opt,name=active,proto3" json:"active,omitempty"`                                                            // Account is active, user can login
	Rank         uint32               `protobuf:"varint,41,opt,name=rank,proto3" json:"rank,omitempty"`                                                               // Rank
	RankTrend    User_RankTrend       `protobuf:"varint,42,opt,name=rank_trend,json=rankTrend,proto3,enum=eolymp.cognito.User_RankTrend" json:"rank_trend,omitempty"` // Rank trend (up or down)
	Name         string               `protobuf:"bytes,10,opt,name=name,proto3" json:"name,omitempty"`                                                                // Full name
	Picture      string               `protobuf:"bytes,11,opt,name=picture,proto3" json:"picture,omitempty"`                                                          // Profile picture
	Company      string               `protobuf:"bytes,12,opt,name=company,proto3" json:"company,omitempty"`                                                          // Company
	Occupation   string               `protobuf:"bytes,13,opt,name=occupation,proto3" json:"occupation,omitempty"`                                                    // Occupation
	Country      string               `protobuf:"bytes,21,opt,name=country,proto3" json:"country,omitempty"`                                                          // Country code
	City         string               `protobuf:"bytes,22,opt,name=city,proto3" json:"city,omitempty"`                                                                // City
	EmailStatus  string               `protobuf:"bytes,32,opt,name=email_status,json=emailStatus,proto3" json:"email_status,omitempty"`                               // Email confirmation status (requires VIEW_PRIVATE_DATA entitlement)
	Birthday     string               `protobuf:"bytes,33,opt,name=birthday,proto3" json:"birthday,omitempty"`                                                        // Birthday (requires VIEW_PRIVATE_DATA entitlement)
	RegisteredOn *timestamp.Timestamp `protobuf:"bytes,34,opt,name=registered_on,json=registeredOn,proto3" json:"registered_on,omitempty"`                            // Exact time when user registered (requires VIEW_PRIVATE_DATA entitlement)
	LastActivity *timestamp.Timestamp `protobuf:"bytes,35,opt,name=last_activity,json=lastActivity,proto3" json:"last_activity,omitempty"`                            // Exact time when user was last active (requires VIEW_PRIVATE_DATA entitlement)
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_cognito_user_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_cognito_user_proto_msgTypes[0]
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
	return file_eolymp_cognito_user_proto_rawDescGZIP(), []int{0}
}

func (x *User) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *User) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *User) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *User) GetActive() bool {
	if x != nil {
		return x.Active
	}
	return false
}

func (x *User) GetRank() uint32 {
	if x != nil {
		return x.Rank
	}
	return 0
}

func (x *User) GetRankTrend() User_RankTrend {
	if x != nil {
		return x.RankTrend
	}
	return User_NONE
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

func (x *User) GetCompany() string {
	if x != nil {
		return x.Company
	}
	return ""
}

func (x *User) GetOccupation() string {
	if x != nil {
		return x.Occupation
	}
	return ""
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

func (x *User) GetEmailStatus() string {
	if x != nil {
		return x.EmailStatus
	}
	return ""
}

func (x *User) GetBirthday() string {
	if x != nil {
		return x.Birthday
	}
	return ""
}

func (x *User) GetRegisteredOn() *timestamp.Timestamp {
	if x != nil {
		return x.RegisteredOn
	}
	return nil
}

func (x *User) GetLastActivity() *timestamp.Timestamp {
	if x != nil {
		return x.LastActivity
	}
	return nil
}

var File_eolymp_cognito_user_proto protoreflect.FileDescriptor

var file_eolymp_cognito_user_proto_rawDesc = []byte{
	0x0a, 0x19, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x63, 0x6f, 0x67, 0x6e, 0x69, 0x74, 0x6f,
	0x2f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x65, 0x6f, 0x6c,
	0x79, 0x6d, 0x70, 0x2e, 0x63, 0x6f, 0x67, 0x6e, 0x69, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb3, 0x04, 0x0a,
	0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x76,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x72, 0x61, 0x6e, 0x6b, 0x18, 0x29, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x72,
	0x61, 0x6e, 0x6b, 0x12, 0x3d, 0x0a, 0x0a, 0x72, 0x61, 0x6e, 0x6b, 0x5f, 0x74, 0x72, 0x65, 0x6e,
	0x64, 0x18, 0x2a, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1e, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70,
	0x2e, 0x63, 0x6f, 0x67, 0x6e, 0x69, 0x74, 0x6f, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x2e, 0x52, 0x61,
	0x6e, 0x6b, 0x54, 0x72, 0x65, 0x6e, 0x64, 0x52, 0x09, 0x72, 0x61, 0x6e, 0x6b, 0x54, 0x72, 0x65,
	0x6e, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x69, 0x63, 0x74, 0x75, 0x72,
	0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x69, 0x63, 0x74, 0x75, 0x72, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x18, 0x0c, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x12, 0x1e, 0x0a, 0x0a, 0x6f, 0x63,
	0x63, 0x75, 0x70, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x6f, 0x63, 0x63, 0x75, 0x70, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x15, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x72, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x69, 0x74, 0x79, 0x18, 0x16, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x63, 0x69, 0x74, 0x79, 0x12, 0x21, 0x0a, 0x0c, 0x65, 0x6d, 0x61, 0x69,
	0x6c, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x20, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x65, 0x6d, 0x61, 0x69, 0x6c, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x62,
	0x69, 0x72, 0x74, 0x68, 0x64, 0x61, 0x79, 0x18, 0x21, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x62,
	0x69, 0x72, 0x74, 0x68, 0x64, 0x61, 0x79, 0x12, 0x3f, 0x0a, 0x0d, 0x72, 0x65, 0x67, 0x69, 0x73,
	0x74, 0x65, 0x72, 0x65, 0x64, 0x5f, 0x6f, 0x6e, 0x18, 0x22, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0c, 0x72, 0x65, 0x67, 0x69,
	0x73, 0x74, 0x65, 0x72, 0x65, 0x64, 0x4f, 0x6e, 0x12, 0x3f, 0x0a, 0x0d, 0x6c, 0x61, 0x73, 0x74,
	0x5f, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x18, 0x23, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0c, 0x6c, 0x61, 0x73,
	0x74, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x22, 0x27, 0x0a, 0x09, 0x52, 0x61, 0x6e,
	0x6b, 0x54, 0x72, 0x65, 0x6e, 0x64, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00,
	0x12, 0x06, 0x0a, 0x02, 0x55, 0x50, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x44, 0x4f, 0x57, 0x4e,
	0x10, 0x02, 0x42, 0x37, 0x5a, 0x35, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74,
	0x73, 0x2f, 0x67, 0x6f, 0x2f, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x63, 0x6f, 0x67, 0x6e,
	0x69, 0x74, 0x6f, 0x3b, 0x63, 0x6f, 0x67, 0x6e, 0x69, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_eolymp_cognito_user_proto_rawDescOnce sync.Once
	file_eolymp_cognito_user_proto_rawDescData = file_eolymp_cognito_user_proto_rawDesc
)

func file_eolymp_cognito_user_proto_rawDescGZIP() []byte {
	file_eolymp_cognito_user_proto_rawDescOnce.Do(func() {
		file_eolymp_cognito_user_proto_rawDescData = protoimpl.X.CompressGZIP(file_eolymp_cognito_user_proto_rawDescData)
	})
	return file_eolymp_cognito_user_proto_rawDescData
}

var file_eolymp_cognito_user_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_eolymp_cognito_user_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_eolymp_cognito_user_proto_goTypes = []interface{}{
	(User_RankTrend)(0),         // 0: eolymp.cognito.User.RankTrend
	(*User)(nil),                // 1: eolymp.cognito.User
	(*timestamp.Timestamp)(nil), // 2: google.protobuf.Timestamp
}
var file_eolymp_cognito_user_proto_depIdxs = []int32{
	0, // 0: eolymp.cognito.User.rank_trend:type_name -> eolymp.cognito.User.RankTrend
	2, // 1: eolymp.cognito.User.registered_on:type_name -> google.protobuf.Timestamp
	2, // 2: eolymp.cognito.User.last_activity:type_name -> google.protobuf.Timestamp
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_eolymp_cognito_user_proto_init() }
func file_eolymp_cognito_user_proto_init() {
	if File_eolymp_cognito_user_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_eolymp_cognito_user_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_eolymp_cognito_user_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_eolymp_cognito_user_proto_goTypes,
		DependencyIndexes: file_eolymp_cognito_user_proto_depIdxs,
		EnumInfos:         file_eolymp_cognito_user_proto_enumTypes,
		MessageInfos:      file_eolymp_cognito_user_proto_msgTypes,
	}.Build()
	File_eolymp_cognito_user_proto = out.File
	file_eolymp_cognito_user_proto_rawDesc = nil
	file_eolymp_cognito_user_proto_goTypes = nil
	file_eolymp_cognito_user_proto_depIdxs = nil
}
