// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.1
// 	protoc        v4.24.4
// source: eolymp/resolver/resolver.proto

package resolver

import (
	_ "github.com/eolymp/go-sdk/eolymp/annotations"
	judge "github.com/eolymp/go-sdk/eolymp/judge"
	ranker "github.com/eolymp/go-sdk/eolymp/ranker"
	universe "github.com/eolymp/go-sdk/eolymp/universe"
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

type Record_Type int32

const (
	Record_UNSPECIFIED Record_Type = 0
	Record_SPACE       Record_Type = 1
	Record_CONTEST     Record_Type = 2
	Record_SCOREBOARD  Record_Type = 3
)

// Enum value maps for Record_Type.
var (
	Record_Type_name = map[int32]string{
		0: "UNSPECIFIED",
		1: "SPACE",
		2: "CONTEST",
		3: "SCOREBOARD",
	}
	Record_Type_value = map[string]int32{
		"UNSPECIFIED": 0,
		"SPACE":       1,
		"CONTEST":     2,
		"SCOREBOARD":  3,
	}
)

func (x Record_Type) Enum() *Record_Type {
	p := new(Record_Type)
	*p = x
	return p
}

func (x Record_Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Record_Type) Descriptor() protoreflect.EnumDescriptor {
	return file_eolymp_resolver_resolver_proto_enumTypes[0].Descriptor()
}

func (Record_Type) Type() protoreflect.EnumType {
	return &file_eolymp_resolver_resolver_proto_enumTypes[0]
}

func (x Record_Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Record_Type.Descriptor instead.
func (Record_Type) EnumDescriptor() ([]byte, []int) {
	return file_eolymp_resolver_resolver_proto_rawDescGZIP(), []int{1, 0}
}

type Authorization struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Authorization) Reset() {
	*x = Authorization{}
	mi := &file_eolymp_resolver_resolver_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Authorization) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Authorization) ProtoMessage() {}

func (x *Authorization) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_resolver_resolver_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Authorization.ProtoReflect.Descriptor instead.
func (*Authorization) Descriptor() ([]byte, []int) {
	return file_eolymp_resolver_resolver_proto_rawDescGZIP(), []int{0}
}

type Record struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Target        *Record_Target         `protobuf:"bytes,20,opt,name=target,proto3" json:"target,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Record) Reset() {
	*x = Record{}
	mi := &file_eolymp_resolver_resolver_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Record) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Record) ProtoMessage() {}

func (x *Record) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_resolver_resolver_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Record.ProtoReflect.Descriptor instead.
func (*Record) Descriptor() ([]byte, []int) {
	return file_eolymp_resolver_resolver_proto_rawDescGZIP(), []int{1}
}

func (x *Record) GetTarget() *Record_Target {
	if x != nil {
		return x.Target
	}
	return nil
}

type ResolveNameInput struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Name          string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ResolveNameInput) Reset() {
	*x = ResolveNameInput{}
	mi := &file_eolymp_resolver_resolver_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ResolveNameInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResolveNameInput) ProtoMessage() {}

func (x *ResolveNameInput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_resolver_resolver_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResolveNameInput.ProtoReflect.Descriptor instead.
func (*ResolveNameInput) Descriptor() ([]byte, []int) {
	return file_eolymp_resolver_resolver_proto_rawDescGZIP(), []int{2}
}

func (x *ResolveNameInput) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type ResolveNameOutput struct {
	state  protoimpl.MessageState `protogen:"open.v1"`
	Target *Record_Target         `protobuf:"bytes,1,opt,name=target,proto3" json:"target,omitempty"`
	// Types that are valid to be assigned to Targetx:
	//
	//	*ResolveNameOutput_Space
	//	*ResolveNameOutput_Contest
	//	*ResolveNameOutput_Scoreboard
	Targetx isResolveNameOutput_Targetx `protobuf_oneof:"targetx"`
	// Types that are valid to be assigned to Auth:
	//
	//	*ResolveNameOutput_Oauth2
	Auth          isResolveNameOutput_Auth `protobuf_oneof:"auth"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ResolveNameOutput) Reset() {
	*x = ResolveNameOutput{}
	mi := &file_eolymp_resolver_resolver_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ResolveNameOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResolveNameOutput) ProtoMessage() {}

func (x *ResolveNameOutput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_resolver_resolver_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResolveNameOutput.ProtoReflect.Descriptor instead.
func (*ResolveNameOutput) Descriptor() ([]byte, []int) {
	return file_eolymp_resolver_resolver_proto_rawDescGZIP(), []int{3}
}

func (x *ResolveNameOutput) GetTarget() *Record_Target {
	if x != nil {
		return x.Target
	}
	return nil
}

func (x *ResolveNameOutput) GetTargetx() isResolveNameOutput_Targetx {
	if x != nil {
		return x.Targetx
	}
	return nil
}

func (x *ResolveNameOutput) GetSpace() *universe.Space {
	if x != nil {
		if x, ok := x.Targetx.(*ResolveNameOutput_Space); ok {
			return x.Space
		}
	}
	return nil
}

func (x *ResolveNameOutput) GetContest() *judge.Contest {
	if x != nil {
		if x, ok := x.Targetx.(*ResolveNameOutput_Contest); ok {
			return x.Contest
		}
	}
	return nil
}

func (x *ResolveNameOutput) GetScoreboard() *ranker.Scoreboard {
	if x != nil {
		if x, ok := x.Targetx.(*ResolveNameOutput_Scoreboard); ok {
			return x.Scoreboard
		}
	}
	return nil
}

func (x *ResolveNameOutput) GetAuth() isResolveNameOutput_Auth {
	if x != nil {
		return x.Auth
	}
	return nil
}

func (x *ResolveNameOutput) GetOauth2() *Authorization_OAuth2 {
	if x != nil {
		if x, ok := x.Auth.(*ResolveNameOutput_Oauth2); ok {
			return x.Oauth2
		}
	}
	return nil
}

type isResolveNameOutput_Targetx interface {
	isResolveNameOutput_Targetx()
}

type ResolveNameOutput_Space struct {
	Space *universe.Space `protobuf:"bytes,10,opt,name=space,proto3,oneof"`
}

type ResolveNameOutput_Contest struct {
	Contest *judge.Contest `protobuf:"bytes,11,opt,name=contest,proto3,oneof"`
}

type ResolveNameOutput_Scoreboard struct {
	Scoreboard *ranker.Scoreboard `protobuf:"bytes,12,opt,name=scoreboard,proto3,oneof"`
}

func (*ResolveNameOutput_Space) isResolveNameOutput_Targetx() {}

func (*ResolveNameOutput_Contest) isResolveNameOutput_Targetx() {}

func (*ResolveNameOutput_Scoreboard) isResolveNameOutput_Targetx() {}

type isResolveNameOutput_Auth interface {
	isResolveNameOutput_Auth()
}

type ResolveNameOutput_Oauth2 struct {
	Oauth2 *Authorization_OAuth2 `protobuf:"bytes,20,opt,name=oauth2,proto3,oneof"`
}

func (*ResolveNameOutput_Oauth2) isResolveNameOutput_Auth() {}

type Authorization_OAuth2 struct {
	state             protoimpl.MessageState `protogen:"open.v1"`
	ClientId          string                 `protobuf:"bytes,1,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	TokenEndpoint     string                 `protobuf:"bytes,2,opt,name=token_endpoint,json=tokenEndpoint,proto3" json:"token_endpoint,omitempty"`
	AuthorizeEndpoint string                 `protobuf:"bytes,3,opt,name=authorize_endpoint,json=authorizeEndpoint,proto3" json:"authorize_endpoint,omitempty"`
	UserinfoEndpoint  string                 `protobuf:"bytes,4,opt,name=userinfo_endpoint,json=userinfoEndpoint,proto3" json:"userinfo_endpoint,omitempty"`
	SignoutEndpoint   string                 `protobuf:"bytes,5,opt,name=signout_endpoint,json=signoutEndpoint,proto3" json:"signout_endpoint,omitempty"`
	unknownFields     protoimpl.UnknownFields
	sizeCache         protoimpl.SizeCache
}

func (x *Authorization_OAuth2) Reset() {
	*x = Authorization_OAuth2{}
	mi := &file_eolymp_resolver_resolver_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Authorization_OAuth2) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Authorization_OAuth2) ProtoMessage() {}

func (x *Authorization_OAuth2) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_resolver_resolver_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Authorization_OAuth2.ProtoReflect.Descriptor instead.
func (*Authorization_OAuth2) Descriptor() ([]byte, []int) {
	return file_eolymp_resolver_resolver_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Authorization_OAuth2) GetClientId() string {
	if x != nil {
		return x.ClientId
	}
	return ""
}

func (x *Authorization_OAuth2) GetTokenEndpoint() string {
	if x != nil {
		return x.TokenEndpoint
	}
	return ""
}

func (x *Authorization_OAuth2) GetAuthorizeEndpoint() string {
	if x != nil {
		return x.AuthorizeEndpoint
	}
	return ""
}

func (x *Authorization_OAuth2) GetUserinfoEndpoint() string {
	if x != nil {
		return x.UserinfoEndpoint
	}
	return ""
}

func (x *Authorization_OAuth2) GetSignoutEndpoint() string {
	if x != nil {
		return x.SignoutEndpoint
	}
	return ""
}

type Record_Target struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Type          Record_Type            `protobuf:"varint,1,opt,name=type,proto3,enum=eolymp.resolver.Record_Type" json:"type,omitempty"`
	Url           string                 `protobuf:"bytes,2,opt,name=url,proto3" json:"url,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Record_Target) Reset() {
	*x = Record_Target{}
	mi := &file_eolymp_resolver_resolver_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Record_Target) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Record_Target) ProtoMessage() {}

func (x *Record_Target) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_resolver_resolver_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Record_Target.ProtoReflect.Descriptor instead.
func (*Record_Target) Descriptor() ([]byte, []int) {
	return file_eolymp_resolver_resolver_proto_rawDescGZIP(), []int{1, 0}
}

func (x *Record_Target) GetType() Record_Type {
	if x != nil {
		return x.Type
	}
	return Record_UNSPECIFIED
}

func (x *Record_Target) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

var File_eolymp_resolver_resolver_proto protoreflect.FileDescriptor

var file_eolymp_resolver_resolver_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65,
	0x72, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0f, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65,
	0x72, 0x1a, 0x1d, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x22, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x72, 0x61, 0x74, 0x65, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1a, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x6a, 0x75, 0x64,
	0x67, 0x65, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x72, 0x61, 0x6e, 0x6b, 0x65, 0x72, 0x2f,
	0x73, 0x63, 0x6f, 0x72, 0x65, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1b, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x75, 0x6e, 0x69, 0x76, 0x65, 0x72, 0x73,
	0x65, 0x2f, 0x73, 0x70, 0x61, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe5, 0x01,
	0x0a, 0x0d, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x1a,
	0xd3, 0x01, 0x0a, 0x06, 0x4f, 0x41, 0x75, 0x74, 0x68, 0x32, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x25, 0x0a, 0x0e, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x5f, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0d, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x2d,
	0x0a, 0x12, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x5f, 0x65, 0x6e, 0x64, 0x70,
	0x6f, 0x69, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x61, 0x75, 0x74, 0x68,
	0x6f, 0x72, 0x69, 0x7a, 0x65, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x2b, 0x0a,
	0x11, 0x75, 0x73, 0x65, 0x72, 0x69, 0x6e, 0x66, 0x6f, 0x5f, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69,
	0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x75, 0x73, 0x65, 0x72, 0x69, 0x6e,
	0x66, 0x6f, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x29, 0x0a, 0x10, 0x73, 0x69,
	0x67, 0x6e, 0x6f, 0x75, 0x74, 0x5f, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x73, 0x69, 0x67, 0x6e, 0x6f, 0x75, 0x74, 0x45, 0x6e, 0x64,
	0x70, 0x6f, 0x69, 0x6e, 0x74, 0x22, 0xcf, 0x01, 0x0a, 0x06, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64,
	0x12, 0x36, 0x0a, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x18, 0x14, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1e, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x6c, 0x76,
	0x65, 0x72, 0x2e, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x2e, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74,
	0x52, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x1a, 0x4c, 0x0a, 0x06, 0x54, 0x61, 0x72, 0x67,
	0x65, 0x74, 0x12, 0x30, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x1c, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x6c, 0x76,
	0x65, 0x72, 0x2e, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x22, 0x3f, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0f,
	0x0a, 0x0b, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12,
	0x09, 0x0a, 0x05, 0x53, 0x50, 0x41, 0x43, 0x45, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x43, 0x4f,
	0x4e, 0x54, 0x45, 0x53, 0x54, 0x10, 0x02, 0x12, 0x0e, 0x0a, 0x0a, 0x53, 0x43, 0x4f, 0x52, 0x45,
	0x42, 0x4f, 0x41, 0x52, 0x44, 0x10, 0x03, 0x22, 0x26, 0x0a, 0x10, 0x52, 0x65, 0x73, 0x6f, 0x6c,
	0x76, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22,
	0xbf, 0x02, 0x0a, 0x11, 0x52, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x4f,
	0x75, 0x74, 0x70, 0x75, 0x74, 0x12, 0x36, 0x0a, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x72,
	0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x2e, 0x54,
	0x61, 0x72, 0x67, 0x65, 0x74, 0x52, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x12, 0x2e, 0x0a,
	0x05, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x65,
	0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x75, 0x6e, 0x69, 0x76, 0x65, 0x72, 0x73, 0x65, 0x2e, 0x53,
	0x70, 0x61, 0x63, 0x65, 0x48, 0x00, 0x52, 0x05, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x31, 0x0a,
	0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x73, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15,
	0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x6a, 0x75, 0x64, 0x67, 0x65, 0x2e, 0x43, 0x6f,
	0x6e, 0x74, 0x65, 0x73, 0x74, 0x48, 0x00, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x73, 0x74,
	0x12, 0x3b, 0x0a, 0x0a, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x18, 0x0c,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x72, 0x61,
	0x6e, 0x6b, 0x65, 0x72, 0x2e, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x48,
	0x00, 0x52, 0x0a, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x12, 0x3f, 0x0a,
	0x06, 0x6f, 0x61, 0x75, 0x74, 0x68, 0x32, 0x18, 0x14, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e,
	0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x72, 0x2e,
	0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x4f, 0x41,
	0x75, 0x74, 0x68, 0x32, 0x48, 0x01, 0x52, 0x06, 0x6f, 0x61, 0x75, 0x74, 0x68, 0x32, 0x42, 0x09,
	0x0a, 0x07, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x78, 0x42, 0x06, 0x0a, 0x04, 0x61, 0x75, 0x74,
	0x68, 0x32, 0x86, 0x01, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x72, 0x12, 0x7a,
	0x0a, 0x0b, 0x52, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x2e,
	0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x72, 0x2e,
	0x52, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x49, 0x6e, 0x70, 0x75, 0x74,
	0x1a, 0x22, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x6c, 0x76,
	0x65, 0x72, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x4f, 0x75,
	0x74, 0x70, 0x75, 0x74, 0x22, 0x24, 0xea, 0xe2, 0x0a, 0x0b, 0xf5, 0xe2, 0x0a, 0x00, 0x00, 0x48,
	0x42, 0xf8, 0xe2, 0x0a, 0x64, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0f, 0x12, 0x0d, 0x2f, 0x6e, 0x61,
	0x6d, 0x65, 0x73, 0x2f, 0x7b, 0x6e, 0x61, 0x6d, 0x65, 0x7d, 0x42, 0x33, 0x5a, 0x31, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f,
	0x67, 0x6f, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x72, 0x65,
	0x73, 0x6f, 0x6c, 0x76, 0x65, 0x72, 0x3b, 0x72, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x72, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_eolymp_resolver_resolver_proto_rawDescOnce sync.Once
	file_eolymp_resolver_resolver_proto_rawDescData = file_eolymp_resolver_resolver_proto_rawDesc
)

func file_eolymp_resolver_resolver_proto_rawDescGZIP() []byte {
	file_eolymp_resolver_resolver_proto_rawDescOnce.Do(func() {
		file_eolymp_resolver_resolver_proto_rawDescData = protoimpl.X.CompressGZIP(file_eolymp_resolver_resolver_proto_rawDescData)
	})
	return file_eolymp_resolver_resolver_proto_rawDescData
}

var file_eolymp_resolver_resolver_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_eolymp_resolver_resolver_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_eolymp_resolver_resolver_proto_goTypes = []any{
	(Record_Type)(0),             // 0: eolymp.resolver.Record.Type
	(*Authorization)(nil),        // 1: eolymp.resolver.Authorization
	(*Record)(nil),               // 2: eolymp.resolver.Record
	(*ResolveNameInput)(nil),     // 3: eolymp.resolver.ResolveNameInput
	(*ResolveNameOutput)(nil),    // 4: eolymp.resolver.ResolveNameOutput
	(*Authorization_OAuth2)(nil), // 5: eolymp.resolver.Authorization.OAuth2
	(*Record_Target)(nil),        // 6: eolymp.resolver.Record.Target
	(*universe.Space)(nil),       // 7: eolymp.universe.Space
	(*judge.Contest)(nil),        // 8: eolymp.judge.Contest
	(*ranker.Scoreboard)(nil),    // 9: eolymp.ranker.Scoreboard
}
var file_eolymp_resolver_resolver_proto_depIdxs = []int32{
	6, // 0: eolymp.resolver.Record.target:type_name -> eolymp.resolver.Record.Target
	6, // 1: eolymp.resolver.ResolveNameOutput.target:type_name -> eolymp.resolver.Record.Target
	7, // 2: eolymp.resolver.ResolveNameOutput.space:type_name -> eolymp.universe.Space
	8, // 3: eolymp.resolver.ResolveNameOutput.contest:type_name -> eolymp.judge.Contest
	9, // 4: eolymp.resolver.ResolveNameOutput.scoreboard:type_name -> eolymp.ranker.Scoreboard
	5, // 5: eolymp.resolver.ResolveNameOutput.oauth2:type_name -> eolymp.resolver.Authorization.OAuth2
	0, // 6: eolymp.resolver.Record.Target.type:type_name -> eolymp.resolver.Record.Type
	3, // 7: eolymp.resolver.Resolver.ResolveName:input_type -> eolymp.resolver.ResolveNameInput
	4, // 8: eolymp.resolver.Resolver.ResolveName:output_type -> eolymp.resolver.ResolveNameOutput
	8, // [8:9] is the sub-list for method output_type
	7, // [7:8] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_eolymp_resolver_resolver_proto_init() }
func file_eolymp_resolver_resolver_proto_init() {
	if File_eolymp_resolver_resolver_proto != nil {
		return
	}
	file_eolymp_resolver_resolver_proto_msgTypes[3].OneofWrappers = []any{
		(*ResolveNameOutput_Space)(nil),
		(*ResolveNameOutput_Contest)(nil),
		(*ResolveNameOutput_Scoreboard)(nil),
		(*ResolveNameOutput_Oauth2)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_eolymp_resolver_resolver_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_eolymp_resolver_resolver_proto_goTypes,
		DependencyIndexes: file_eolymp_resolver_resolver_proto_depIdxs,
		EnumInfos:         file_eolymp_resolver_resolver_proto_enumTypes,
		MessageInfos:      file_eolymp_resolver_resolver_proto_msgTypes,
	}.Build()
	File_eolymp_resolver_resolver_proto = out.File
	file_eolymp_resolver_resolver_proto_rawDesc = nil
	file_eolymp_resolver_resolver_proto_goTypes = nil
	file_eolymp_resolver_resolver_proto_depIdxs = nil
}
