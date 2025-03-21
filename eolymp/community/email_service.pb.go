// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v4.24.4
// source: eolymp/community/email_service.proto

package community

import (
	_ "github.com/eolymp/go-sdk/eolymp/annotations"
	ecm "github.com/eolymp/go-sdk/eolymp/ecm"
	_ "github.com/eolymp/go-sdk/eolymp/wellknown"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/timestamppb"
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

type SendEmailInput struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	MemberId      string                 `protobuf:"bytes,1,opt,name=member_id,json=memberId,proto3" json:"member_id,omitempty"`
	DryRun        bool                   `protobuf:"varint,20,opt,name=dry_run,json=dryRun,proto3" json:"dry_run,omitempty"` // send a test email to the current user instead of the member
	BccMe         bool                   `protobuf:"varint,21,opt,name=bcc_me,json=bccMe,proto3" json:"bcc_me,omitempty"`    // send a copy of the email to the current user
	Type          EmailType              `protobuf:"varint,10,opt,name=type,proto3,enum=eolymp.community.EmailType" json:"type,omitempty"`
	Campaign      string                 `protobuf:"bytes,15,opt,name=campaign,proto3" json:"campaign,omitempty"` // for internal use, campaign id for tracking feedback
	Locale        string                 `protobuf:"bytes,11,opt,name=locale,proto3" json:"locale,omitempty"`     // locale for the content
	Subject       string                 `protobuf:"bytes,12,opt,name=subject,proto3" json:"subject,omitempty"`
	Content       *ecm.Content           `protobuf:"bytes,13,opt,name=content,proto3" json:"content,omitempty"`
	Parameters    map[string]string      `protobuf:"bytes,14,rep,name=parameters,proto3" json:"parameters,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SendEmailInput) Reset() {
	*x = SendEmailInput{}
	mi := &file_eolymp_community_email_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SendEmailInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendEmailInput) ProtoMessage() {}

func (x *SendEmailInput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_community_email_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendEmailInput.ProtoReflect.Descriptor instead.
func (*SendEmailInput) Descriptor() ([]byte, []int) {
	return file_eolymp_community_email_service_proto_rawDescGZIP(), []int{0}
}

func (x *SendEmailInput) GetMemberId() string {
	if x != nil {
		return x.MemberId
	}
	return ""
}

func (x *SendEmailInput) GetDryRun() bool {
	if x != nil {
		return x.DryRun
	}
	return false
}

func (x *SendEmailInput) GetBccMe() bool {
	if x != nil {
		return x.BccMe
	}
	return false
}

func (x *SendEmailInput) GetType() EmailType {
	if x != nil {
		return x.Type
	}
	return EmailType_UNKNOWN_TYPE
}

func (x *SendEmailInput) GetCampaign() string {
	if x != nil {
		return x.Campaign
	}
	return ""
}

func (x *SendEmailInput) GetLocale() string {
	if x != nil {
		return x.Locale
	}
	return ""
}

func (x *SendEmailInput) GetSubject() string {
	if x != nil {
		return x.Subject
	}
	return ""
}

func (x *SendEmailInput) GetContent() *ecm.Content {
	if x != nil {
		return x.Content
	}
	return nil
}

func (x *SendEmailInput) GetParameters() map[string]string {
	if x != nil {
		return x.Parameters
	}
	return nil
}

type SendEmailOutput struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	MessageId     string                 `protobuf:"bytes,1,opt,name=message_id,json=messageId,proto3" json:"message_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SendEmailOutput) Reset() {
	*x = SendEmailOutput{}
	mi := &file_eolymp_community_email_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SendEmailOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendEmailOutput) ProtoMessage() {}

func (x *SendEmailOutput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_community_email_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendEmailOutput.ProtoReflect.Descriptor instead.
func (*SendEmailOutput) Descriptor() ([]byte, []int) {
	return file_eolymp_community_email_service_proto_rawDescGZIP(), []int{1}
}

func (x *SendEmailOutput) GetMessageId() string {
	if x != nil {
		return x.MessageId
	}
	return ""
}

type DescribeEmailUsageInput struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DescribeEmailUsageInput) Reset() {
	*x = DescribeEmailUsageInput{}
	mi := &file_eolymp_community_email_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DescribeEmailUsageInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescribeEmailUsageInput) ProtoMessage() {}

func (x *DescribeEmailUsageInput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_community_email_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DescribeEmailUsageInput.ProtoReflect.Descriptor instead.
func (*DescribeEmailUsageInput) Descriptor() ([]byte, []int) {
	return file_eolymp_community_email_service_proto_rawDescGZIP(), []int{2}
}

type DescribeEmailUsageOutput struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	DailyEmails   uint32                 `protobuf:"varint,2,opt,name=daily_emails,json=dailyEmails,proto3" json:"daily_emails,omitempty"`
	MonthlyEmails uint32                 `protobuf:"varint,3,opt,name=monthly_emails,json=monthlyEmails,proto3" json:"monthly_emails,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DescribeEmailUsageOutput) Reset() {
	*x = DescribeEmailUsageOutput{}
	mi := &file_eolymp_community_email_service_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DescribeEmailUsageOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescribeEmailUsageOutput) ProtoMessage() {}

func (x *DescribeEmailUsageOutput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_community_email_service_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DescribeEmailUsageOutput.ProtoReflect.Descriptor instead.
func (*DescribeEmailUsageOutput) Descriptor() ([]byte, []int) {
	return file_eolymp_community_email_service_proto_rawDescGZIP(), []int{3}
}

func (x *DescribeEmailUsageOutput) GetDailyEmails() uint32 {
	if x != nil {
		return x.DailyEmails
	}
	return 0
}

func (x *DescribeEmailUsageOutput) GetMonthlyEmails() uint32 {
	if x != nil {
		return x.MonthlyEmails
	}
	return 0
}

var File_eolymp_community_email_service_proto protoreflect.FileDescriptor

var file_eolymp_community_email_service_proto_rawDesc = string([]byte{
	0x0a, 0x24, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69,
	0x74, 0x79, 0x2f, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x1a, 0x1d, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70,
	0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x68, 0x74, 0x74,
	0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x22, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x72, 0x61, 0x74, 0x65,
	0x6c, 0x69, 0x6d, 0x69, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x65, 0x6f, 0x6c,
	0x79, 0x6d, 0x70, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f,
	0x73, 0x63, 0x6f, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x21, 0x65, 0x6f, 0x6c,
	0x79, 0x6d, 0x70, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x2f, 0x65, 0x6d,
	0x61, 0x69, 0x6c, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d,
	0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79,
	0x2f, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x18, 0x65,
	0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x65, 0x63, 0x6d, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f,
	0x77, 0x65, 0x6c, 0x6c, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x2f, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x21, 0x65, 0x6f, 0x6c, 0x79, 0x6d,
	0x70, 0x2f, 0x77, 0x65, 0x6c, 0x6c, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x2f, 0x65, 0x78, 0x70, 0x72,
	0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9c, 0x03,
	0x0a, 0x0e, 0x53, 0x65, 0x6e, 0x64, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x49, 0x6e, 0x70, 0x75, 0x74,
	0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x49, 0x64, 0x12, 0x17, 0x0a,
	0x07, 0x64, 0x72, 0x79, 0x5f, 0x72, 0x75, 0x6e, 0x18, 0x14, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06,
	0x64, 0x72, 0x79, 0x52, 0x75, 0x6e, 0x12, 0x15, 0x0a, 0x06, 0x62, 0x63, 0x63, 0x5f, 0x6d, 0x65,
	0x18, 0x15, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x62, 0x63, 0x63, 0x4d, 0x65, 0x12, 0x2f, 0x0a,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1b, 0x2e, 0x65, 0x6f,
	0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x2e, 0x45,
	0x6d, 0x61, 0x69, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1a,
	0x0a, 0x08, 0x63, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x63, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x6c, 0x6f,
	0x63, 0x61, 0x6c, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6c, 0x6f, 0x63, 0x61,
	0x6c, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x0c, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x2d, 0x0a, 0x07,
	0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e,
	0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x65, 0x63, 0x6d, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x50, 0x0a, 0x0a, 0x70,
	0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x18, 0x0e, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x30, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69,
	0x74, 0x79, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x49, 0x6e, 0x70, 0x75,
	0x74, 0x2e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x52, 0x0a, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x1a, 0x3d, 0x0a,
	0x0f, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x30, 0x0a, 0x0f,
	0x53, 0x65, 0x6e, 0x64, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x12,
	0x1d, 0x0a, 0x0a, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x49, 0x64, 0x22, 0x19,
	0x0a, 0x17, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x55,
	0x73, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x22, 0x64, 0x0a, 0x18, 0x44, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x62, 0x65, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x55, 0x73, 0x61, 0x67, 0x65, 0x4f,
	0x75, 0x74, 0x70, 0x75, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x64, 0x61, 0x69, 0x6c, 0x79, 0x5f, 0x65,
	0x6d, 0x61, 0x69, 0x6c, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x64, 0x61, 0x69,
	0x6c, 0x79, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x25, 0x0a, 0x0e, 0x6d, 0x6f, 0x6e, 0x74,
	0x68, 0x6c, 0x79, 0x5f, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x0d, 0x6d, 0x6f, 0x6e, 0x74, 0x68, 0x6c, 0x79, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x73, 0x32,
	0xe4, 0x02, 0x0a, 0x0c, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0xa2, 0x01, 0x0a, 0x09, 0x53, 0x65, 0x6e, 0x64, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x20,
	0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74,
	0x79, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x49, 0x6e, 0x70, 0x75, 0x74,
	0x1a, 0x21, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e,
	0x69, 0x74, 0x79, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x4f, 0x75, 0x74,
	0x70, 0x75, 0x74, 0x22, 0x50, 0xea, 0xe2, 0x0a, 0x0b, 0xf5, 0xe2, 0x0a, 0x00, 0x00, 0xa0, 0x40,
	0xf8, 0xe2, 0x0a, 0x14, 0x82, 0xe3, 0x0a, 0x1a, 0x8a, 0xe3, 0x0a, 0x16, 0x63, 0x6f, 0x6d, 0x6d,
	0x75, 0x6e, 0x69, 0x74, 0x79, 0x3a, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x3a, 0x77, 0x72, 0x69,
	0x74, 0x65, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1d, 0x22, 0x1b, 0x2f, 0x6d, 0x65, 0x6d, 0x62, 0x65,
	0x72, 0x73, 0x2f, 0x7b, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x7d, 0x2f, 0x6e,
	0x6f, 0x74, 0x69, 0x66, 0x79, 0x12, 0xae, 0x01, 0x0a, 0x12, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x62, 0x65, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x55, 0x73, 0x61, 0x67, 0x65, 0x12, 0x29, 0x2e, 0x65,
	0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x2e,
	0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x55, 0x73, 0x61,
	0x67, 0x65, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x2a, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x2e, 0x44, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x62, 0x65, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x55, 0x73, 0x61, 0x67, 0x65, 0x4f, 0x75, 0x74,
	0x70, 0x75, 0x74, 0x22, 0x41, 0xea, 0xe2, 0x0a, 0x0b, 0xf5, 0xe2, 0x0a, 0x00, 0x00, 0xa0, 0x40,
	0xf8, 0xe2, 0x0a, 0x14, 0x82, 0xe3, 0x0a, 0x19, 0x8a, 0xe3, 0x0a, 0x15, 0x63, 0x6f, 0x6d, 0x6d,
	0x75, 0x6e, 0x69, 0x74, 0x79, 0x3a, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x3a, 0x72, 0x65, 0x61,
	0x64, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0f, 0x12, 0x0d, 0x2f, 0x75, 0x73, 0x61, 0x67, 0x65, 0x2f,
	0x65, 0x6d, 0x61, 0x69, 0x6c, 0x73, 0x42, 0x35, 0x5a, 0x33, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x67, 0x6f, 0x2d, 0x73,
	0x64, 0x6b, 0x2f, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e,
	0x69, 0x74, 0x79, 0x3b, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_eolymp_community_email_service_proto_rawDescOnce sync.Once
	file_eolymp_community_email_service_proto_rawDescData []byte
)

func file_eolymp_community_email_service_proto_rawDescGZIP() []byte {
	file_eolymp_community_email_service_proto_rawDescOnce.Do(func() {
		file_eolymp_community_email_service_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_eolymp_community_email_service_proto_rawDesc), len(file_eolymp_community_email_service_proto_rawDesc)))
	})
	return file_eolymp_community_email_service_proto_rawDescData
}

var file_eolymp_community_email_service_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_eolymp_community_email_service_proto_goTypes = []any{
	(*SendEmailInput)(nil),           // 0: eolymp.community.SendEmailInput
	(*SendEmailOutput)(nil),          // 1: eolymp.community.SendEmailOutput
	(*DescribeEmailUsageInput)(nil),  // 2: eolymp.community.DescribeEmailUsageInput
	(*DescribeEmailUsageOutput)(nil), // 3: eolymp.community.DescribeEmailUsageOutput
	nil,                              // 4: eolymp.community.SendEmailInput.ParametersEntry
	(EmailType)(0),                   // 5: eolymp.community.EmailType
	(*ecm.Content)(nil),              // 6: eolymp.ecm.Content
}
var file_eolymp_community_email_service_proto_depIdxs = []int32{
	5, // 0: eolymp.community.SendEmailInput.type:type_name -> eolymp.community.EmailType
	6, // 1: eolymp.community.SendEmailInput.content:type_name -> eolymp.ecm.Content
	4, // 2: eolymp.community.SendEmailInput.parameters:type_name -> eolymp.community.SendEmailInput.ParametersEntry
	0, // 3: eolymp.community.EmailService.SendEmail:input_type -> eolymp.community.SendEmailInput
	2, // 4: eolymp.community.EmailService.DescribeEmailUsage:input_type -> eolymp.community.DescribeEmailUsageInput
	1, // 5: eolymp.community.EmailService.SendEmail:output_type -> eolymp.community.SendEmailOutput
	3, // 6: eolymp.community.EmailService.DescribeEmailUsage:output_type -> eolymp.community.DescribeEmailUsageOutput
	5, // [5:7] is the sub-list for method output_type
	3, // [3:5] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_eolymp_community_email_service_proto_init() }
func file_eolymp_community_email_service_proto_init() {
	if File_eolymp_community_email_service_proto != nil {
		return
	}
	file_eolymp_community_email_type_proto_init()
	file_eolymp_community_member_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_eolymp_community_email_service_proto_rawDesc), len(file_eolymp_community_email_service_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_eolymp_community_email_service_proto_goTypes,
		DependencyIndexes: file_eolymp_community_email_service_proto_depIdxs,
		MessageInfos:      file_eolymp_community_email_service_proto_msgTypes,
	}.Build()
	File_eolymp_community_email_service_proto = out.File
	file_eolymp_community_email_service_proto_goTypes = nil
	file_eolymp_community_email_service_proto_depIdxs = nil
}
