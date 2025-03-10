// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v4.24.4
// source: eolymp/newsletter/campaign.proto

package newsletter

import (
	community "github.com/eolymp/go-sdk/eolymp/community"
	ecm "github.com/eolymp/go-sdk/eolymp/ecm"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
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

type Campaign_Patch int32

const (
	Campaign_ALL     Campaign_Patch = 0
	Campaign_TYPE    Campaign_Patch = 1
	Campaign_NAME    Campaign_Patch = 5
	Campaign_SUBJECT Campaign_Patch = 2
	Campaign_CONTENT Campaign_Patch = 3
	Campaign_LOCALE  Campaign_Patch = 4
)

// Enum value maps for Campaign_Patch.
var (
	Campaign_Patch_name = map[int32]string{
		0: "ALL",
		1: "TYPE",
		5: "NAME",
		2: "SUBJECT",
		3: "CONTENT",
		4: "LOCALE",
	}
	Campaign_Patch_value = map[string]int32{
		"ALL":     0,
		"TYPE":    1,
		"NAME":    5,
		"SUBJECT": 2,
		"CONTENT": 3,
		"LOCALE":  4,
	}
)

func (x Campaign_Patch) Enum() *Campaign_Patch {
	p := new(Campaign_Patch)
	*p = x
	return p
}

func (x Campaign_Patch) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Campaign_Patch) Descriptor() protoreflect.EnumDescriptor {
	return file_eolymp_newsletter_campaign_proto_enumTypes[0].Descriptor()
}

func (Campaign_Patch) Type() protoreflect.EnumType {
	return &file_eolymp_newsletter_campaign_proto_enumTypes[0]
}

func (x Campaign_Patch) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Campaign_Patch.Descriptor instead.
func (Campaign_Patch) EnumDescriptor() ([]byte, []int) {
	return file_eolymp_newsletter_campaign_proto_rawDescGZIP(), []int{0, 0}
}

type Campaign_Extra int32

const (
	Campaign_UNKNOWN_EXTRA  Campaign_Extra = 0
	Campaign_CONTENT_VALUE  Campaign_Extra = 1
	Campaign_CONTENT_RENDER Campaign_Extra = 2
)

// Enum value maps for Campaign_Extra.
var (
	Campaign_Extra_name = map[int32]string{
		0: "UNKNOWN_EXTRA",
		1: "CONTENT_VALUE",
		2: "CONTENT_RENDER",
	}
	Campaign_Extra_value = map[string]int32{
		"UNKNOWN_EXTRA":  0,
		"CONTENT_VALUE":  1,
		"CONTENT_RENDER": 2,
	}
)

func (x Campaign_Extra) Enum() *Campaign_Extra {
	p := new(Campaign_Extra)
	*p = x
	return p
}

func (x Campaign_Extra) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Campaign_Extra) Descriptor() protoreflect.EnumDescriptor {
	return file_eolymp_newsletter_campaign_proto_enumTypes[1].Descriptor()
}

func (Campaign_Extra) Type() protoreflect.EnumType {
	return &file_eolymp_newsletter_campaign_proto_enumTypes[1]
}

func (x Campaign_Extra) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Campaign_Extra.Descriptor instead.
func (Campaign_Extra) EnumDescriptor() ([]byte, []int) {
	return file_eolymp_newsletter_campaign_proto_rawDescGZIP(), []int{0, 1}
}

type Campaign struct {
	state           protoimpl.MessageState `protogen:"open.v1"`
	Id              string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Type            community.EmailType    `protobuf:"varint,2,opt,name=type,proto3,enum=eolymp.community.EmailType" json:"type,omitempty"`
	CreatedAt       *timestamppb.Timestamp `protobuf:"bytes,20,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	Name            string                 `protobuf:"bytes,10,opt,name=name,proto3" json:"name,omitempty"`
	Subject         string                 `protobuf:"bytes,11,opt,name=subject,proto3" json:"subject,omitempty"`
	Content         *ecm.Content           `protobuf:"bytes,12,opt,name=content,proto3" json:"content,omitempty"`
	RecipientsCount uint32                 `protobuf:"varint,30,opt,name=recipients_count,json=recipientsCount,proto3" json:"recipients_count,omitempty"`
	PendingCount    uint32                 `protobuf:"varint,31,opt,name=pending_count,json=pendingCount,proto3" json:"pending_count,omitempty"`
	SentCount       uint32                 `protobuf:"varint,32,opt,name=sent_count,json=sentCount,proto3" json:"sent_count,omitempty"`
	ErrorCount      uint32                 `protobuf:"varint,33,opt,name=error_count,json=errorCount,proto3" json:"error_count,omitempty"`
	unknownFields   protoimpl.UnknownFields
	sizeCache       protoimpl.SizeCache
}

func (x *Campaign) Reset() {
	*x = Campaign{}
	mi := &file_eolymp_newsletter_campaign_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Campaign) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Campaign) ProtoMessage() {}

func (x *Campaign) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_newsletter_campaign_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Campaign.ProtoReflect.Descriptor instead.
func (*Campaign) Descriptor() ([]byte, []int) {
	return file_eolymp_newsletter_campaign_proto_rawDescGZIP(), []int{0}
}

func (x *Campaign) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Campaign) GetType() community.EmailType {
	if x != nil {
		return x.Type
	}
	return community.EmailType(0)
}

func (x *Campaign) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Campaign) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Campaign) GetSubject() string {
	if x != nil {
		return x.Subject
	}
	return ""
}

func (x *Campaign) GetContent() *ecm.Content {
	if x != nil {
		return x.Content
	}
	return nil
}

func (x *Campaign) GetRecipientsCount() uint32 {
	if x != nil {
		return x.RecipientsCount
	}
	return 0
}

func (x *Campaign) GetPendingCount() uint32 {
	if x != nil {
		return x.PendingCount
	}
	return 0
}

func (x *Campaign) GetSentCount() uint32 {
	if x != nil {
		return x.SentCount
	}
	return 0
}

func (x *Campaign) GetErrorCount() uint32 {
	if x != nil {
		return x.ErrorCount
	}
	return 0
}

type Campaign_Translation struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Locale        string                 `protobuf:"bytes,102,opt,name=locale,proto3" json:"locale,omitempty"`
	Subject       string                 `protobuf:"bytes,103,opt,name=subject,proto3" json:"subject,omitempty"`
	Content       *ecm.Content           `protobuf:"bytes,104,opt,name=content,proto3" json:"content,omitempty"`
	Automatic     bool                   `protobuf:"varint,105,opt,name=automatic,proto3" json:"automatic,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Campaign_Translation) Reset() {
	*x = Campaign_Translation{}
	mi := &file_eolymp_newsletter_campaign_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Campaign_Translation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Campaign_Translation) ProtoMessage() {}

func (x *Campaign_Translation) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_newsletter_campaign_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Campaign_Translation.ProtoReflect.Descriptor instead.
func (*Campaign_Translation) Descriptor() ([]byte, []int) {
	return file_eolymp_newsletter_campaign_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Campaign_Translation) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Campaign_Translation) GetLocale() string {
	if x != nil {
		return x.Locale
	}
	return ""
}

func (x *Campaign_Translation) GetSubject() string {
	if x != nil {
		return x.Subject
	}
	return ""
}

func (x *Campaign_Translation) GetContent() *ecm.Content {
	if x != nil {
		return x.Content
	}
	return nil
}

func (x *Campaign_Translation) GetAutomatic() bool {
	if x != nil {
		return x.Automatic
	}
	return false
}

var File_eolymp_newsletter_campaign_proto protoreflect.FileDescriptor

var file_eolymp_newsletter_campaign_proto_rawDesc = string([]byte{
	0x0a, 0x20, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x6e, 0x65, 0x77, 0x73, 0x6c, 0x65, 0x74,
	0x74, 0x65, 0x72, 0x2f, 0x63, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x11, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x6e, 0x65, 0x77, 0x73, 0x6c,
	0x65, 0x74, 0x74, 0x65, 0x72, 0x1a, 0x21, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x63, 0x6f,
	0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x2f, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x5f, 0x74, 0x79,
	0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x18, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70,
	0x2f, 0x65, 0x63, 0x6d, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xa1, 0x05, 0x0a, 0x08, 0x43, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x2f, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1b,
	0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74,
	0x79, 0x2e, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18,
	0x14, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x2d, 0x0a, 0x07, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x65, 0x6f,
	0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x65, 0x63, 0x6d, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x29, 0x0a, 0x10, 0x72, 0x65, 0x63,
	0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x1e, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x0f, 0x72, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x43,
	0x6f, 0x75, 0x6e, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x70, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x5f,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x1f, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c, 0x70, 0x65, 0x6e,
	0x64, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x65, 0x6e,
	0x74, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x20, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x73,
	0x65, 0x6e, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x21, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x1a, 0x9c, 0x01, 0x0a, 0x0b, 0x54, 0x72,
	0x61, 0x6e, 0x73, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x6c, 0x6f, 0x63,
	0x61, 0x6c, 0x65, 0x18, 0x66, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6c, 0x6f, 0x63, 0x61, 0x6c,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x67, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x2d, 0x0a, 0x07, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x68, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x65,
	0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x65, 0x63, 0x6d, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x61, 0x75,
	0x74, 0x6f, 0x6d, 0x61, 0x74, 0x69, 0x63, 0x18, 0x69, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x61,
	0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x69, 0x63, 0x22, 0x4a, 0x0a, 0x05, 0x50, 0x61, 0x74, 0x63,
	0x68, 0x12, 0x07, 0x0a, 0x03, 0x41, 0x4c, 0x4c, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x54, 0x59,
	0x50, 0x45, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x41, 0x4d, 0x45, 0x10, 0x05, 0x12, 0x0b,
	0x0a, 0x07, 0x53, 0x55, 0x42, 0x4a, 0x45, 0x43, 0x54, 0x10, 0x02, 0x12, 0x0b, 0x0a, 0x07, 0x43,
	0x4f, 0x4e, 0x54, 0x45, 0x4e, 0x54, 0x10, 0x03, 0x12, 0x0a, 0x0a, 0x06, 0x4c, 0x4f, 0x43, 0x41,
	0x4c, 0x45, 0x10, 0x04, 0x22, 0x41, 0x0a, 0x05, 0x45, 0x78, 0x74, 0x72, 0x61, 0x12, 0x11, 0x0a,
	0x0d, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x5f, 0x45, 0x58, 0x54, 0x52, 0x41, 0x10, 0x00,
	0x12, 0x11, 0x0a, 0x0d, 0x43, 0x4f, 0x4e, 0x54, 0x45, 0x4e, 0x54, 0x5f, 0x56, 0x41, 0x4c, 0x55,
	0x45, 0x10, 0x01, 0x12, 0x12, 0x0a, 0x0e, 0x43, 0x4f, 0x4e, 0x54, 0x45, 0x4e, 0x54, 0x5f, 0x52,
	0x45, 0x4e, 0x44, 0x45, 0x52, 0x10, 0x02, 0x42, 0x37, 0x5a, 0x35, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x67, 0x6f, 0x2d,
	0x73, 0x64, 0x6b, 0x2f, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x6e, 0x65, 0x77, 0x73, 0x6c,
	0x65, 0x74, 0x74, 0x65, 0x72, 0x3b, 0x6e, 0x65, 0x77, 0x73, 0x6c, 0x65, 0x74, 0x74, 0x65, 0x72,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_eolymp_newsletter_campaign_proto_rawDescOnce sync.Once
	file_eolymp_newsletter_campaign_proto_rawDescData []byte
)

func file_eolymp_newsletter_campaign_proto_rawDescGZIP() []byte {
	file_eolymp_newsletter_campaign_proto_rawDescOnce.Do(func() {
		file_eolymp_newsletter_campaign_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_eolymp_newsletter_campaign_proto_rawDesc), len(file_eolymp_newsletter_campaign_proto_rawDesc)))
	})
	return file_eolymp_newsletter_campaign_proto_rawDescData
}

var file_eolymp_newsletter_campaign_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_eolymp_newsletter_campaign_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_eolymp_newsletter_campaign_proto_goTypes = []any{
	(Campaign_Patch)(0),           // 0: eolymp.newsletter.Campaign.Patch
	(Campaign_Extra)(0),           // 1: eolymp.newsletter.Campaign.Extra
	(*Campaign)(nil),              // 2: eolymp.newsletter.Campaign
	(*Campaign_Translation)(nil),  // 3: eolymp.newsletter.Campaign.Translation
	(community.EmailType)(0),      // 4: eolymp.community.EmailType
	(*timestamppb.Timestamp)(nil), // 5: google.protobuf.Timestamp
	(*ecm.Content)(nil),           // 6: eolymp.ecm.Content
}
var file_eolymp_newsletter_campaign_proto_depIdxs = []int32{
	4, // 0: eolymp.newsletter.Campaign.type:type_name -> eolymp.community.EmailType
	5, // 1: eolymp.newsletter.Campaign.created_at:type_name -> google.protobuf.Timestamp
	6, // 2: eolymp.newsletter.Campaign.content:type_name -> eolymp.ecm.Content
	6, // 3: eolymp.newsletter.Campaign.Translation.content:type_name -> eolymp.ecm.Content
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_eolymp_newsletter_campaign_proto_init() }
func file_eolymp_newsletter_campaign_proto_init() {
	if File_eolymp_newsletter_campaign_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_eolymp_newsletter_campaign_proto_rawDesc), len(file_eolymp_newsletter_campaign_proto_rawDesc)),
			NumEnums:      2,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_eolymp_newsletter_campaign_proto_goTypes,
		DependencyIndexes: file_eolymp_newsletter_campaign_proto_depIdxs,
		EnumInfos:         file_eolymp_newsletter_campaign_proto_enumTypes,
		MessageInfos:      file_eolymp_newsletter_campaign_proto_msgTypes,
	}.Build()
	File_eolymp_newsletter_campaign_proto = out.File
	file_eolymp_newsletter_campaign_proto_goTypes = nil
	file_eolymp_newsletter_campaign_proto_depIdxs = nil
}
