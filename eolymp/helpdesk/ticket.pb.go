// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v3.21.12
// source: eolymp/helpdesk/ticket.proto

package helpdesk

import (
	ecm "github.com/eolymp/go-sdk/eolymp/ecm"
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

type Ticket_Type int32

const (
	Ticket_NONE                  Ticket_Type = 0
	Ticket_QUESTION              Ticket_Type = 1
	Ticket_QUOTA_INCREASE        Ticket_Type = 2
	Ticket_FEEDBACK              Ticket_Type = 3
	Ticket_ACADEMIC_PLAN_REQUEST Ticket_Type = 4
	Ticket_SALES_REQUEST         Ticket_Type = 5
)

// Enum value maps for Ticket_Type.
var (
	Ticket_Type_name = map[int32]string{
		0: "NONE",
		1: "QUESTION",
		2: "QUOTA_INCREASE",
		3: "FEEDBACK",
		4: "ACADEMIC_PLAN_REQUEST",
		5: "SALES_REQUEST",
	}
	Ticket_Type_value = map[string]int32{
		"NONE":                  0,
		"QUESTION":              1,
		"QUOTA_INCREASE":        2,
		"FEEDBACK":              3,
		"ACADEMIC_PLAN_REQUEST": 4,
		"SALES_REQUEST":         5,
	}
)

func (x Ticket_Type) Enum() *Ticket_Type {
	p := new(Ticket_Type)
	*p = x
	return p
}

func (x Ticket_Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Ticket_Type) Descriptor() protoreflect.EnumDescriptor {
	return file_eolymp_helpdesk_ticket_proto_enumTypes[0].Descriptor()
}

func (Ticket_Type) Type() protoreflect.EnumType {
	return &file_eolymp_helpdesk_ticket_proto_enumTypes[0]
}

func (x Ticket_Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Ticket_Type.Descriptor instead.
func (Ticket_Type) EnumDescriptor() ([]byte, []int) {
	return file_eolymp_helpdesk_ticket_proto_rawDescGZIP(), []int{0, 0}
}

type Ticket_Status int32

const (
	Ticket_UNKNOWN  Ticket_Status = 0
	Ticket_PENDING  Ticket_Status = 1 // pending review by agent
	Ticket_AWAITING Ticket_Status = 2 // pending action from customer
	Ticket_CLOSED   Ticket_Status = 3 // ticket is resolved
	Ticket_APPROVED Ticket_Status = 4 // request is approved
	Ticket_REJECTED Ticket_Status = 5 // request is rejected
)

// Enum value maps for Ticket_Status.
var (
	Ticket_Status_name = map[int32]string{
		0: "UNKNOWN",
		1: "PENDING",
		2: "AWAITING",
		3: "CLOSED",
		4: "APPROVED",
		5: "REJECTED",
	}
	Ticket_Status_value = map[string]int32{
		"UNKNOWN":  0,
		"PENDING":  1,
		"AWAITING": 2,
		"CLOSED":   3,
		"APPROVED": 4,
		"REJECTED": 5,
	}
)

func (x Ticket_Status) Enum() *Ticket_Status {
	p := new(Ticket_Status)
	*p = x
	return p
}

func (x Ticket_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Ticket_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_eolymp_helpdesk_ticket_proto_enumTypes[1].Descriptor()
}

func (Ticket_Status) Type() protoreflect.EnumType {
	return &file_eolymp_helpdesk_ticket_proto_enumTypes[1]
}

func (x Ticket_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Ticket_Status.Descriptor instead.
func (Ticket_Status) EnumDescriptor() ([]byte, []int) {
	return file_eolymp_helpdesk_ticket_proto_rawDescGZIP(), []int{0, 1}
}

type Ticket struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Type      Ticket_Type            `protobuf:"varint,2,opt,name=type,proto3,enum=eolymp.helpdesk.Ticket_Type" json:"type,omitempty"`
	UserId    string                 `protobuf:"bytes,3,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	UserEmail string                 `protobuf:"bytes,4,opt,name=user_email,json=userEmail,proto3" json:"user_email,omitempty"`
	Metadata  map[string]string      `protobuf:"bytes,5,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Status    Ticket_Status          `protobuf:"varint,6,opt,name=status,proto3,enum=eolymp.helpdesk.Ticket_Status" json:"status,omitempty"`
	Locale    string                 `protobuf:"bytes,9,opt,name=locale,proto3" json:"locale,omitempty"`
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,10,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt *timestamppb.Timestamp `protobuf:"bytes,11,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	Secret    string                 `protobuf:"bytes,12,opt,name=secret,proto3" json:"secret,omitempty"`
	Subject   string                 `protobuf:"bytes,20,opt,name=subject,proto3" json:"subject,omitempty"`
	Message   *ecm.Content           `protobuf:"bytes,50,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *Ticket) Reset() {
	*x = Ticket{}
	mi := &file_eolymp_helpdesk_ticket_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Ticket) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ticket) ProtoMessage() {}

func (x *Ticket) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_helpdesk_ticket_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Ticket.ProtoReflect.Descriptor instead.
func (*Ticket) Descriptor() ([]byte, []int) {
	return file_eolymp_helpdesk_ticket_proto_rawDescGZIP(), []int{0}
}

func (x *Ticket) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Ticket) GetType() Ticket_Type {
	if x != nil {
		return x.Type
	}
	return Ticket_NONE
}

func (x *Ticket) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *Ticket) GetUserEmail() string {
	if x != nil {
		return x.UserEmail
	}
	return ""
}

func (x *Ticket) GetMetadata() map[string]string {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *Ticket) GetStatus() Ticket_Status {
	if x != nil {
		return x.Status
	}
	return Ticket_UNKNOWN
}

func (x *Ticket) GetLocale() string {
	if x != nil {
		return x.Locale
	}
	return ""
}

func (x *Ticket) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Ticket) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *Ticket) GetSecret() string {
	if x != nil {
		return x.Secret
	}
	return ""
}

func (x *Ticket) GetSubject() string {
	if x != nil {
		return x.Subject
	}
	return ""
}

func (x *Ticket) GetMessage() *ecm.Content {
	if x != nil {
		return x.Message
	}
	return nil
}

type Ticket_Comment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId    string                 `protobuf:"bytes,3,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	UserEmail string                 `protobuf:"bytes,4,opt,name=user_email,json=userEmail,proto3" json:"user_email,omitempty"`
	Metadata  map[string]string      `protobuf:"bytes,5,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,10,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt *timestamppb.Timestamp `protobuf:"bytes,11,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	Message   *ecm.Content           `protobuf:"bytes,50,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *Ticket_Comment) Reset() {
	*x = Ticket_Comment{}
	mi := &file_eolymp_helpdesk_ticket_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Ticket_Comment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ticket_Comment) ProtoMessage() {}

func (x *Ticket_Comment) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_helpdesk_ticket_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Ticket_Comment.ProtoReflect.Descriptor instead.
func (*Ticket_Comment) Descriptor() ([]byte, []int) {
	return file_eolymp_helpdesk_ticket_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Ticket_Comment) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Ticket_Comment) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *Ticket_Comment) GetUserEmail() string {
	if x != nil {
		return x.UserEmail
	}
	return ""
}

func (x *Ticket_Comment) GetMetadata() map[string]string {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *Ticket_Comment) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Ticket_Comment) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *Ticket_Comment) GetMessage() *ecm.Content {
	if x != nil {
		return x.Message
	}
	return nil
}

var File_eolymp_helpdesk_ticket_proto protoreflect.FileDescriptor

var file_eolymp_helpdesk_ticket_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x68, 0x65, 0x6c, 0x70, 0x64, 0x65, 0x73,
	0x6b, 0x2f, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f,
	0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x68, 0x65, 0x6c, 0x70, 0x64, 0x65, 0x73, 0x6b, 0x1a,
	0x18, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x65, 0x63, 0x6d, 0x2f, 0x63, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf4, 0x08, 0x0a, 0x06, 0x54,
	0x69, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x30, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x1c, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x68, 0x65, 0x6c,
	0x70, 0x64, 0x65, 0x73, 0x6b, 0x2e, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x54, 0x79, 0x70,
	0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x73, 0x65, 0x72, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12,
	0x41, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x05, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x25, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x68, 0x65, 0x6c, 0x70, 0x64,
	0x65, 0x73, 0x6b, 0x2e, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x12, 0x36, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x1e, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x68, 0x65, 0x6c, 0x70,
	0x64, 0x65, 0x73, 0x6b, 0x2e, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x6c, 0x6f,
	0x63, 0x61, 0x6c, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6c, 0x6f, 0x63, 0x61,
	0x6c, 0x65, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x39, 0x0a,
	0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x63, 0x72,
	0x65, 0x74, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74,
	0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x14, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x2d, 0x0a, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x32, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x65, 0x6f,
	0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x65, 0x63, 0x6d, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0xfe, 0x02, 0x0a, 0x07, 0x43, 0x6f,
	0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1d,
	0x0a, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x75, 0x73, 0x65, 0x72, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x49, 0x0a,
	0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x2d, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x68, 0x65, 0x6c, 0x70, 0x64, 0x65, 0x73,
	0x6b, 0x2e, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74,
	0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x08,
	0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61,
	0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x2d,
	0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x32, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x13, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x65, 0x63, 0x6d, 0x2e, 0x43, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x3b, 0x0a,
	0x0d, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10,
	0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79,
	0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x3b, 0x0a, 0x0d, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x6e, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x08, 0x0a, 0x04, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x51, 0x55, 0x45,
	0x53, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x01, 0x12, 0x12, 0x0a, 0x0e, 0x51, 0x55, 0x4f, 0x54, 0x41,
	0x5f, 0x49, 0x4e, 0x43, 0x52, 0x45, 0x41, 0x53, 0x45, 0x10, 0x02, 0x12, 0x0c, 0x0a, 0x08, 0x46,
	0x45, 0x45, 0x44, 0x42, 0x41, 0x43, 0x4b, 0x10, 0x03, 0x12, 0x19, 0x0a, 0x15, 0x41, 0x43, 0x41,
	0x44, 0x45, 0x4d, 0x49, 0x43, 0x5f, 0x50, 0x4c, 0x41, 0x4e, 0x5f, 0x52, 0x45, 0x51, 0x55, 0x45,
	0x53, 0x54, 0x10, 0x04, 0x12, 0x11, 0x0a, 0x0d, 0x53, 0x41, 0x4c, 0x45, 0x53, 0x5f, 0x52, 0x45,
	0x51, 0x55, 0x45, 0x53, 0x54, 0x10, 0x05, 0x22, 0x58, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x0b,
	0x0a, 0x07, 0x50, 0x45, 0x4e, 0x44, 0x49, 0x4e, 0x47, 0x10, 0x01, 0x12, 0x0c, 0x0a, 0x08, 0x41,
	0x57, 0x41, 0x49, 0x54, 0x49, 0x4e, 0x47, 0x10, 0x02, 0x12, 0x0a, 0x0a, 0x06, 0x43, 0x4c, 0x4f,
	0x53, 0x45, 0x44, 0x10, 0x03, 0x12, 0x0c, 0x0a, 0x08, 0x41, 0x50, 0x50, 0x52, 0x4f, 0x56, 0x45,
	0x44, 0x10, 0x04, 0x12, 0x0c, 0x0a, 0x08, 0x52, 0x45, 0x4a, 0x45, 0x43, 0x54, 0x45, 0x44, 0x10,
	0x05, 0x42, 0x33, 0x5a, 0x31, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x67, 0x6f, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x65, 0x6f,
	0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x68, 0x65, 0x6c, 0x70, 0x64, 0x65, 0x73, 0x6b, 0x3b, 0x68, 0x65,
	0x6c, 0x70, 0x64, 0x65, 0x73, 0x6b, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_eolymp_helpdesk_ticket_proto_rawDescOnce sync.Once
	file_eolymp_helpdesk_ticket_proto_rawDescData = file_eolymp_helpdesk_ticket_proto_rawDesc
)

func file_eolymp_helpdesk_ticket_proto_rawDescGZIP() []byte {
	file_eolymp_helpdesk_ticket_proto_rawDescOnce.Do(func() {
		file_eolymp_helpdesk_ticket_proto_rawDescData = protoimpl.X.CompressGZIP(file_eolymp_helpdesk_ticket_proto_rawDescData)
	})
	return file_eolymp_helpdesk_ticket_proto_rawDescData
}

var file_eolymp_helpdesk_ticket_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_eolymp_helpdesk_ticket_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_eolymp_helpdesk_ticket_proto_goTypes = []any{
	(Ticket_Type)(0),              // 0: eolymp.helpdesk.Ticket.Type
	(Ticket_Status)(0),            // 1: eolymp.helpdesk.Ticket.Status
	(*Ticket)(nil),                // 2: eolymp.helpdesk.Ticket
	(*Ticket_Comment)(nil),        // 3: eolymp.helpdesk.Ticket.Comment
	nil,                           // 4: eolymp.helpdesk.Ticket.MetadataEntry
	nil,                           // 5: eolymp.helpdesk.Ticket.Comment.MetadataEntry
	(*timestamppb.Timestamp)(nil), // 6: google.protobuf.Timestamp
	(*ecm.Content)(nil),           // 7: eolymp.ecm.Content
}
var file_eolymp_helpdesk_ticket_proto_depIdxs = []int32{
	0,  // 0: eolymp.helpdesk.Ticket.type:type_name -> eolymp.helpdesk.Ticket.Type
	4,  // 1: eolymp.helpdesk.Ticket.metadata:type_name -> eolymp.helpdesk.Ticket.MetadataEntry
	1,  // 2: eolymp.helpdesk.Ticket.status:type_name -> eolymp.helpdesk.Ticket.Status
	6,  // 3: eolymp.helpdesk.Ticket.created_at:type_name -> google.protobuf.Timestamp
	6,  // 4: eolymp.helpdesk.Ticket.updated_at:type_name -> google.protobuf.Timestamp
	7,  // 5: eolymp.helpdesk.Ticket.message:type_name -> eolymp.ecm.Content
	5,  // 6: eolymp.helpdesk.Ticket.Comment.metadata:type_name -> eolymp.helpdesk.Ticket.Comment.MetadataEntry
	6,  // 7: eolymp.helpdesk.Ticket.Comment.created_at:type_name -> google.protobuf.Timestamp
	6,  // 8: eolymp.helpdesk.Ticket.Comment.updated_at:type_name -> google.protobuf.Timestamp
	7,  // 9: eolymp.helpdesk.Ticket.Comment.message:type_name -> eolymp.ecm.Content
	10, // [10:10] is the sub-list for method output_type
	10, // [10:10] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_eolymp_helpdesk_ticket_proto_init() }
func file_eolymp_helpdesk_ticket_proto_init() {
	if File_eolymp_helpdesk_ticket_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_eolymp_helpdesk_ticket_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_eolymp_helpdesk_ticket_proto_goTypes,
		DependencyIndexes: file_eolymp_helpdesk_ticket_proto_depIdxs,
		EnumInfos:         file_eolymp_helpdesk_ticket_proto_enumTypes,
		MessageInfos:      file_eolymp_helpdesk_ticket_proto_msgTypes,
	}.Build()
	File_eolymp_helpdesk_ticket_proto = out.File
	file_eolymp_helpdesk_ticket_proto_rawDesc = nil
	file_eolymp_helpdesk_ticket_proto_goTypes = nil
	file_eolymp_helpdesk_ticket_proto_depIdxs = nil
}
