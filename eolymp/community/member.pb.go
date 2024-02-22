// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v3.18.1
// source: eolymp/community/member.proto

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

type Member_Extra int32

const (
	Member_NO_EXTRA   Member_Extra = 0
	Member_TIER       Member_Extra = 1
	Member_STATS      Member_Extra = 2
	Member_GROUPS     Member_Extra = 3
	Member_ATTRIBUTES Member_Extra = 4
)

// Enum value maps for Member_Extra.
var (
	Member_Extra_name = map[int32]string{
		0: "NO_EXTRA",
		1: "TIER",
		2: "STATS",
		3: "GROUPS",
		4: "ATTRIBUTES",
	}
	Member_Extra_value = map[string]int32{
		"NO_EXTRA":   0,
		"TIER":       1,
		"STATS":      2,
		"GROUPS":     3,
		"ATTRIBUTES": 4,
	}
)

func (x Member_Extra) Enum() *Member_Extra {
	p := new(Member_Extra)
	*p = x
	return p
}

func (x Member_Extra) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Member_Extra) Descriptor() protoreflect.EnumDescriptor {
	return file_eolymp_community_member_proto_enumTypes[0].Descriptor()
}

func (Member_Extra) Type() protoreflect.EnumType {
	return &file_eolymp_community_member_proto_enumTypes[0]
}

func (x Member_Extra) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Member_Extra.Descriptor instead.
func (Member_Extra) EnumDescriptor() ([]byte, []int) {
	return file_eolymp_community_member_proto_rawDescGZIP(), []int{0, 0}
}

type Member struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id             string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name           string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"` // display name, readonly, users nickname, ghosts name or teams name
	Url            string                 `protobuf:"bytes,3,opt,name=url,proto3" json:"url,omitempty"`
	Rank           int32                  `protobuf:"varint,70,opt,name=rank,proto3" json:"rank,omitempty"`
	Score          int32                  `protobuf:"varint,71,opt,name=score,proto3" json:"score,omitempty"`
	Active         bool                   `protobuf:"varint,10,opt,name=active,proto3" json:"active,omitempty"`
	Incomplete     bool                   `protobuf:"varint,20,opt,name=incomplete,proto3" json:"incomplete,omitempty"` // member profile (attributes) is missing some information
	Unofficial     bool                   `protobuf:"varint,30,opt,name=unofficial,proto3" json:"unofficial,omitempty"` // member participates in all competitions unofficially
	Secret         bool                   `protobuf:"varint,40,opt,name=secret,proto3" json:"secret,omitempty"`         // member is secret and does not appear on anywhere (for example, an admin who performs testing)
	TierId         string                 `protobuf:"bytes,50,opt,name=tier_id,json=tierId,proto3" json:"tier_id,omitempty"`
	FallbackTierId string                 `protobuf:"bytes,51,opt,name=fallback_tier_id,json=fallbackTierId,proto3" json:"fallback_tier_id,omitempty"`
	CreatedAt      *timestamppb.Timestamp `protobuf:"bytes,60,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	// Types that are assignable to Account:
	//	*Member_User
	//	*Member_Team
	//	*Member_Ghost
	Account    isMember_Account   `protobuf_oneof:"account"`
	Stats      *Member_Stats      `protobuf:"bytes,300,opt,name=stats,proto3" json:"stats,omitempty"`
	Groups     []string           `protobuf:"bytes,200,rep,name=groups,proto3" json:"groups,omitempty"`
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

func (x *Member) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *Member) GetRank() int32 {
	if x != nil {
		return x.Rank
	}
	return 0
}

func (x *Member) GetScore() int32 {
	if x != nil {
		return x.Score
	}
	return 0
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

func (x *Member) GetTierId() string {
	if x != nil {
		return x.TierId
	}
	return ""
}

func (x *Member) GetFallbackTierId() string {
	if x != nil {
		return x.FallbackTierId
	}
	return ""
}

func (x *Member) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
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

func (x *Member) GetStats() *Member_Stats {
	if x != nil {
		return x.Stats
	}
	return nil
}

func (x *Member) GetGroups() []string {
	if x != nil {
		return x.Groups
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

type Member_Stats struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Streak              int32 `protobuf:"varint,10,opt,name=streak,proto3" json:"streak,omitempty"`
	ProblemsSolved      int32 `protobuf:"varint,20,opt,name=problems_solved,json=problemsSolved,proto3" json:"problems_solved,omitempty"`
	SubmissionsAccepted int32 `protobuf:"varint,30,opt,name=submissions_accepted,json=submissionsAccepted,proto3" json:"submissions_accepted,omitempty"`
	SubmissionsTotal    int32 `protobuf:"varint,41,opt,name=submissions_total,json=submissionsTotal,proto3" json:"submissions_total,omitempty"`
}

func (x *Member_Stats) Reset() {
	*x = Member_Stats{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_community_member_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Member_Stats) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Member_Stats) ProtoMessage() {}

func (x *Member_Stats) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use Member_Stats.ProtoReflect.Descriptor instead.
func (*Member_Stats) Descriptor() ([]byte, []int) {
	return file_eolymp_community_member_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Member_Stats) GetStreak() int32 {
	if x != nil {
		return x.Streak
	}
	return 0
}

func (x *Member_Stats) GetProblemsSolved() int32 {
	if x != nil {
		return x.ProblemsSolved
	}
	return 0
}

func (x *Member_Stats) GetSubmissionsAccepted() int32 {
	if x != nil {
		return x.SubmissionsAccepted
	}
	return 0
}

func (x *Member_Stats) GetSubmissionsTotal() int32 {
	if x != nil {
		return x.SubmissionsTotal
	}
	return 0
}

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
	0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xf5, 0x06, 0x0a, 0x06, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75,
	0x72, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x61, 0x6e, 0x6b, 0x18, 0x46, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x04, 0x72, 0x61, 0x6e, 0x6b, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x18,
	0x47, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x61, 0x63,
	0x74, 0x69, 0x76, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x69, 0x6e, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65,
	0x74, 0x65, 0x18, 0x14, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x69, 0x6e, 0x63, 0x6f, 0x6d, 0x70,
	0x6c, 0x65, 0x74, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x75, 0x6e, 0x6f, 0x66, 0x66, 0x69, 0x63, 0x69,
	0x61, 0x6c, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x75, 0x6e, 0x6f, 0x66, 0x66, 0x69,
	0x63, 0x69, 0x61, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x18, 0x28,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x12, 0x17, 0x0a, 0x07,
	0x74, 0x69, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x32, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74,
	0x69, 0x65, 0x72, 0x49, 0x64, 0x12, 0x28, 0x0a, 0x10, 0x66, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63,
	0x6b, 0x5f, 0x74, 0x69, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x33, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0e, 0x66, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x54, 0x69, 0x65, 0x72, 0x49, 0x64, 0x12,
	0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x3c, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x2c, 0x0a, 0x04, 0x75, 0x73,
	0x65, 0x72, 0x18, 0x64, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d,
	0x70, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x2e, 0x55, 0x73, 0x65, 0x72,
	0x48, 0x00, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x2c, 0x0a, 0x04, 0x74, 0x65, 0x61, 0x6d,
	0x18, 0x65, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x2e, 0x54, 0x65, 0x61, 0x6d, 0x48, 0x00,
	0x52, 0x04, 0x74, 0x65, 0x61, 0x6d, 0x12, 0x2f, 0x0a, 0x05, 0x67, 0x68, 0x6f, 0x73, 0x74, 0x18,
	0x66, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x2e, 0x47, 0x68, 0x6f, 0x73, 0x74, 0x48, 0x00,
	0x52, 0x05, 0x67, 0x68, 0x6f, 0x73, 0x74, 0x12, 0x35, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x73,
	0x18, 0xac, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x2e, 0x4d, 0x65, 0x6d, 0x62, 0x65,
	0x72, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x73, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x73, 0x12, 0x17,
	0x0a, 0x06, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x18, 0xc8, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x06, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x12, 0x42, 0x0a, 0x0a, 0x61, 0x74, 0x74, 0x72, 0x69,
	0x62, 0x75, 0x74, 0x65, 0x73, 0x18, 0x84, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x65,
	0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x2e,
	0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x2e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52,
	0x0a, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x1a, 0xa8, 0x01, 0x0a, 0x05,
	0x53, 0x74, 0x61, 0x74, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6b, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6b, 0x12, 0x27, 0x0a,
	0x0f, 0x70, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x73, 0x5f, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x64,
	0x18, 0x14, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0e, 0x70, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x73,
	0x53, 0x6f, 0x6c, 0x76, 0x65, 0x64, 0x12, 0x31, 0x0a, 0x14, 0x73, 0x75, 0x62, 0x6d, 0x69, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x73, 0x5f, 0x61, 0x63, 0x63, 0x65, 0x70, 0x74, 0x65, 0x64, 0x18, 0x1e,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x13, 0x73, 0x75, 0x62, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x73, 0x41, 0x63, 0x63, 0x65, 0x70, 0x74, 0x65, 0x64, 0x12, 0x2b, 0x0a, 0x11, 0x73, 0x75, 0x62,
	0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x5f, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x29,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x10, 0x73, 0x75, 0x62, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x73, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x22, 0x46, 0x0a, 0x05, 0x45, 0x78, 0x74, 0x72, 0x61, 0x12,
	0x0c, 0x0a, 0x08, 0x4e, 0x4f, 0x5f, 0x45, 0x58, 0x54, 0x52, 0x41, 0x10, 0x00, 0x12, 0x08, 0x0a,
	0x04, 0x54, 0x49, 0x45, 0x52, 0x10, 0x01, 0x12, 0x09, 0x0a, 0x05, 0x53, 0x54, 0x41, 0x54, 0x53,
	0x10, 0x02, 0x12, 0x0a, 0x0a, 0x06, 0x47, 0x52, 0x4f, 0x55, 0x50, 0x53, 0x10, 0x03, 0x12, 0x0e,
	0x0a, 0x0a, 0x41, 0x54, 0x54, 0x52, 0x49, 0x42, 0x55, 0x54, 0x45, 0x53, 0x10, 0x04, 0x42, 0x09,
	0x0a, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x42, 0x35, 0x5a, 0x33, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x67,
	0x6f, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x63, 0x6f, 0x6d,
	0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x3b, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
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
var file_eolymp_community_member_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_eolymp_community_member_proto_goTypes = []interface{}{
	(Member_Extra)(0),             // 0: eolymp.community.Member.Extra
	(*Member)(nil),                // 1: eolymp.community.Member
	(*Member_Stats)(nil),          // 2: eolymp.community.Member.Stats
	(*timestamppb.Timestamp)(nil), // 3: google.protobuf.Timestamp
	(*User)(nil),                  // 4: eolymp.community.User
	(*Team)(nil),                  // 5: eolymp.community.Team
	(*Ghost)(nil),                 // 6: eolymp.community.Ghost
	(*Attribute_Value)(nil),       // 7: eolymp.community.Attribute.Value
}
var file_eolymp_community_member_proto_depIdxs = []int32{
	3, // 0: eolymp.community.Member.created_at:type_name -> google.protobuf.Timestamp
	4, // 1: eolymp.community.Member.user:type_name -> eolymp.community.User
	5, // 2: eolymp.community.Member.team:type_name -> eolymp.community.Team
	6, // 3: eolymp.community.Member.ghost:type_name -> eolymp.community.Ghost
	2, // 4: eolymp.community.Member.stats:type_name -> eolymp.community.Member.Stats
	7, // 5: eolymp.community.Member.attributes:type_name -> eolymp.community.Attribute.Value
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
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
		file_eolymp_community_member_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Member_Stats); i {
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
			NumEnums:      1,
			NumMessages:   2,
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
