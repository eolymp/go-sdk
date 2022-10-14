// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.18.1
// source: eolymp/resolver/resolver.proto

package resolver

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
	return file_eolymp_resolver_resolver_proto_rawDescGZIP(), []int{0, 0}
}

type Record struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key     string         `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	SpaceId string         `protobuf:"bytes,2,opt,name=space_id,json=spaceId,proto3" json:"space_id,omitempty"`
	Summary string         `protobuf:"bytes,3,opt,name=summary,proto3" json:"summary,omitempty"`
	Auth    *Record_Auth   `protobuf:"bytes,10,opt,name=auth,proto3" json:"auth,omitempty"`
	Target  *Record_Target `protobuf:"bytes,20,opt,name=target,proto3" json:"target,omitempty"`
}

func (x *Record) Reset() {
	*x = Record{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_resolver_resolver_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Record) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Record) ProtoMessage() {}

func (x *Record) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_resolver_resolver_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
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
	return file_eolymp_resolver_resolver_proto_rawDescGZIP(), []int{0}
}

func (x *Record) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *Record) GetSpaceId() string {
	if x != nil {
		return x.SpaceId
	}
	return ""
}

func (x *Record) GetSummary() string {
	if x != nil {
		return x.Summary
	}
	return ""
}

func (x *Record) GetAuth() *Record_Auth {
	if x != nil {
		return x.Auth
	}
	return nil
}

func (x *Record) GetTarget() *Record_Target {
	if x != nil {
		return x.Target
	}
	return nil
}

type ResolveNameInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *ResolveNameInput) Reset() {
	*x = ResolveNameInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_resolver_resolver_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResolveNameInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResolveNameInput) ProtoMessage() {}

func (x *ResolveNameInput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_resolver_resolver_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
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
	return file_eolymp_resolver_resolver_proto_rawDescGZIP(), []int{1}
}

func (x *ResolveNameInput) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type ResolveNameOutput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Auth   *Record_Auth   `protobuf:"bytes,1,opt,name=auth,proto3" json:"auth,omitempty"`
	Target *Record_Target `protobuf:"bytes,2,opt,name=target,proto3" json:"target,omitempty"`
}

func (x *ResolveNameOutput) Reset() {
	*x = ResolveNameOutput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_resolver_resolver_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResolveNameOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResolveNameOutput) ProtoMessage() {}

func (x *ResolveNameOutput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_resolver_resolver_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
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
	return file_eolymp_resolver_resolver_proto_rawDescGZIP(), []int{2}
}

func (x *ResolveNameOutput) GetAuth() *Record_Auth {
	if x != nil {
		return x.Auth
	}
	return nil
}

func (x *ResolveNameOutput) GetTarget() *Record_Target {
	if x != nil {
		return x.Target
	}
	return nil
}

type Record_Auth struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClientId     string `protobuf:"bytes,10,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	ClientSecret string `protobuf:"bytes,11,opt,name=client_secret,json=clientSecret,proto3" json:"client_secret,omitempty"`
	Scopes       string `protobuf:"bytes,12,opt,name=scopes,proto3" json:"scopes,omitempty"`
	TokenUrl     string `protobuf:"bytes,20,opt,name=token_url,json=tokenUrl,proto3" json:"token_url,omitempty"`
	AuthorizeUrl string `protobuf:"bytes,30,opt,name=authorize_url,json=authorizeUrl,proto3" json:"authorize_url,omitempty"`
}

func (x *Record_Auth) Reset() {
	*x = Record_Auth{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_resolver_resolver_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Record_Auth) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Record_Auth) ProtoMessage() {}

func (x *Record_Auth) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_resolver_resolver_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Record_Auth.ProtoReflect.Descriptor instead.
func (*Record_Auth) Descriptor() ([]byte, []int) {
	return file_eolymp_resolver_resolver_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Record_Auth) GetClientId() string {
	if x != nil {
		return x.ClientId
	}
	return ""
}

func (x *Record_Auth) GetClientSecret() string {
	if x != nil {
		return x.ClientSecret
	}
	return ""
}

func (x *Record_Auth) GetScopes() string {
	if x != nil {
		return x.Scopes
	}
	return ""
}

func (x *Record_Auth) GetTokenUrl() string {
	if x != nil {
		return x.TokenUrl
	}
	return ""
}

func (x *Record_Auth) GetAuthorizeUrl() string {
	if x != nil {
		return x.AuthorizeUrl
	}
	return ""
}

type Record_Target struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type Record_Type `protobuf:"varint,1,opt,name=type,proto3,enum=eolymp.resolver.Record_Type" json:"type,omitempty"`
	Url  string      `protobuf:"bytes,2,opt,name=url,proto3" json:"url,omitempty"`
}

func (x *Record_Target) Reset() {
	*x = Record_Target{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_resolver_resolver_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Record_Target) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Record_Target) ProtoMessage() {}

func (x *Record_Target) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_resolver_resolver_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
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
	return file_eolymp_resolver_resolver_proto_rawDescGZIP(), []int{0, 1}
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
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xed, 0x03, 0x0a, 0x06, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x12,
	0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x70, 0x61, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x70, 0x61, 0x63, 0x65, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07,
	0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73,
	0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x12, 0x30, 0x0a, 0x04, 0x61, 0x75, 0x74, 0x68, 0x18, 0x0a,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x72, 0x65,
	0x73, 0x6f, 0x6c, 0x76, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x2e, 0x41, 0x75,
	0x74, 0x68, 0x52, 0x04, 0x61, 0x75, 0x74, 0x68, 0x12, 0x36, 0x0a, 0x06, 0x74, 0x61, 0x72, 0x67,
	0x65, 0x74, 0x18, 0x14, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d,
	0x70, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x63, 0x6f, 0x72,
	0x64, 0x2e, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x52, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74,
	0x1a, 0xa2, 0x01, 0x0a, 0x04, 0x41, 0x75, 0x74, 0x68, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x5f, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x63,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73,
	0x63, 0x6f, 0x70, 0x65, 0x73, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x63, 0x6f,
	0x70, 0x65, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x75, 0x72, 0x6c,
	0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x55, 0x72, 0x6c,
	0x12, 0x23, 0x0a, 0x0d, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x5f, 0x75, 0x72,
	0x6c, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69,
	0x7a, 0x65, 0x55, 0x72, 0x6c, 0x1a, 0x4c, 0x0a, 0x06, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x12,
	0x30, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1c, 0x2e,
	0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x72, 0x2e,
	0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x75, 0x72, 0x6c, 0x22, 0x3f, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0f, 0x0a, 0x0b, 0x55,
	0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05,
	0x53, 0x50, 0x41, 0x43, 0x45, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x43, 0x4f, 0x4e, 0x54, 0x45,
	0x53, 0x54, 0x10, 0x02, 0x12, 0x0e, 0x0a, 0x0a, 0x53, 0x43, 0x4f, 0x52, 0x45, 0x42, 0x4f, 0x41,
	0x52, 0x44, 0x10, 0x03, 0x22, 0x26, 0x0a, 0x10, 0x52, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x4e,
	0x61, 0x6d, 0x65, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x7d, 0x0a, 0x11,
	0x52, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x4f, 0x75, 0x74, 0x70, 0x75,
	0x74, 0x12, 0x30, 0x0a, 0x04, 0x61, 0x75, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1c, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65,
	0x72, 0x2e, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x52, 0x04, 0x61,
	0x75, 0x74, 0x68, 0x12, 0x36, 0x0a, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x72, 0x65, 0x73,
	0x6f, 0x6c, 0x76, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x2e, 0x54, 0x61, 0x72,
	0x67, 0x65, 0x74, 0x52, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x32, 0x86, 0x01, 0x0a, 0x08,
	0x52, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x72, 0x12, 0x7a, 0x0a, 0x0b, 0x52, 0x65, 0x73, 0x6f,
	0x6c, 0x76, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70,
	0x2e, 0x72, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x6c, 0x76,
	0x65, 0x4e, 0x61, 0x6d, 0x65, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x22, 0x2e, 0x65, 0x6f, 0x6c,
	0x79, 0x6d, 0x70, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x73,
	0x6f, 0x6c, 0x76, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x22, 0x24,
	0xea, 0xe2, 0x0a, 0x0b, 0xf5, 0xe2, 0x0a, 0x00, 0x00, 0x48, 0x42, 0xf8, 0xe2, 0x0a, 0x64, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x0f, 0x12, 0x0d, 0x2f, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x2f, 0x7b, 0x6e,
	0x61, 0x6d, 0x65, 0x7d, 0x42, 0x33, 0x5a, 0x31, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x67, 0x6f, 0x2d, 0x73, 0x64, 0x6b,
	0x2f, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x72,
	0x3b, 0x72, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
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
var file_eolymp_resolver_resolver_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_eolymp_resolver_resolver_proto_goTypes = []interface{}{
	(Record_Type)(0),          // 0: eolymp.resolver.Record.Type
	(*Record)(nil),            // 1: eolymp.resolver.Record
	(*ResolveNameInput)(nil),  // 2: eolymp.resolver.ResolveNameInput
	(*ResolveNameOutput)(nil), // 3: eolymp.resolver.ResolveNameOutput
	(*Record_Auth)(nil),       // 4: eolymp.resolver.Record.Auth
	(*Record_Target)(nil),     // 5: eolymp.resolver.Record.Target
}
var file_eolymp_resolver_resolver_proto_depIdxs = []int32{
	4, // 0: eolymp.resolver.Record.auth:type_name -> eolymp.resolver.Record.Auth
	5, // 1: eolymp.resolver.Record.target:type_name -> eolymp.resolver.Record.Target
	4, // 2: eolymp.resolver.ResolveNameOutput.auth:type_name -> eolymp.resolver.Record.Auth
	5, // 3: eolymp.resolver.ResolveNameOutput.target:type_name -> eolymp.resolver.Record.Target
	0, // 4: eolymp.resolver.Record.Target.type:type_name -> eolymp.resolver.Record.Type
	2, // 5: eolymp.resolver.Resolver.ResolveName:input_type -> eolymp.resolver.ResolveNameInput
	3, // 6: eolymp.resolver.Resolver.ResolveName:output_type -> eolymp.resolver.ResolveNameOutput
	6, // [6:7] is the sub-list for method output_type
	5, // [5:6] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_eolymp_resolver_resolver_proto_init() }
func file_eolymp_resolver_resolver_proto_init() {
	if File_eolymp_resolver_resolver_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_eolymp_resolver_resolver_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Record); i {
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
		file_eolymp_resolver_resolver_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResolveNameInput); i {
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
		file_eolymp_resolver_resolver_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResolveNameOutput); i {
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
		file_eolymp_resolver_resolver_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Record_Auth); i {
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
		file_eolymp_resolver_resolver_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Record_Target); i {
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
			RawDescriptor: file_eolymp_resolver_resolver_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
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
