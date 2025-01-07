// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.2
// 	protoc        v4.24.4
// source: eolymp/newsletter/recipient_service.proto

package newsletter

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

type CreateRecipientInput struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Recipient     *Recipient             `protobuf:"bytes,10,opt,name=recipient,proto3" json:"recipient,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateRecipientInput) Reset() {
	*x = CreateRecipientInput{}
	mi := &file_eolymp_newsletter_recipient_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateRecipientInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRecipientInput) ProtoMessage() {}

func (x *CreateRecipientInput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_newsletter_recipient_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRecipientInput.ProtoReflect.Descriptor instead.
func (*CreateRecipientInput) Descriptor() ([]byte, []int) {
	return file_eolymp_newsletter_recipient_service_proto_rawDescGZIP(), []int{0}
}

func (x *CreateRecipientInput) GetRecipient() *Recipient {
	if x != nil {
		return x.Recipient
	}
	return nil
}

type CreateRecipientOutput struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	RecipientId   string                 `protobuf:"bytes,1,opt,name=recipient_id,json=recipientId,proto3" json:"recipient_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateRecipientOutput) Reset() {
	*x = CreateRecipientOutput{}
	mi := &file_eolymp_newsletter_recipient_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateRecipientOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRecipientOutput) ProtoMessage() {}

func (x *CreateRecipientOutput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_newsletter_recipient_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRecipientOutput.ProtoReflect.Descriptor instead.
func (*CreateRecipientOutput) Descriptor() ([]byte, []int) {
	return file_eolymp_newsletter_recipient_service_proto_rawDescGZIP(), []int{1}
}

func (x *CreateRecipientOutput) GetRecipientId() string {
	if x != nil {
		return x.RecipientId
	}
	return ""
}

type ImportRecipientsInput struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Types that are valid to be assigned to Source:
	//
	//	*ImportRecipientsInput_GroupId
	Source        isImportRecipientsInput_Source `protobuf_oneof:"source"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ImportRecipientsInput) Reset() {
	*x = ImportRecipientsInput{}
	mi := &file_eolymp_newsletter_recipient_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ImportRecipientsInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ImportRecipientsInput) ProtoMessage() {}

func (x *ImportRecipientsInput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_newsletter_recipient_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ImportRecipientsInput.ProtoReflect.Descriptor instead.
func (*ImportRecipientsInput) Descriptor() ([]byte, []int) {
	return file_eolymp_newsletter_recipient_service_proto_rawDescGZIP(), []int{2}
}

func (x *ImportRecipientsInput) GetSource() isImportRecipientsInput_Source {
	if x != nil {
		return x.Source
	}
	return nil
}

func (x *ImportRecipientsInput) GetGroupId() string {
	if x != nil {
		if x, ok := x.Source.(*ImportRecipientsInput_GroupId); ok {
			return x.GroupId
		}
	}
	return ""
}

type isImportRecipientsInput_Source interface {
	isImportRecipientsInput_Source()
}

type ImportRecipientsInput_GroupId struct {
	GroupId string `protobuf:"bytes,11,opt,name=group_id,json=groupId,proto3,oneof"`
}

func (*ImportRecipientsInput_GroupId) isImportRecipientsInput_Source() {}

type ImportRecipientsOutput struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ImportedTotal string                 `protobuf:"bytes,1,opt,name=imported_total,json=importedTotal,proto3" json:"imported_total,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ImportRecipientsOutput) Reset() {
	*x = ImportRecipientsOutput{}
	mi := &file_eolymp_newsletter_recipient_service_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ImportRecipientsOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ImportRecipientsOutput) ProtoMessage() {}

func (x *ImportRecipientsOutput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_newsletter_recipient_service_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ImportRecipientsOutput.ProtoReflect.Descriptor instead.
func (*ImportRecipientsOutput) Descriptor() ([]byte, []int) {
	return file_eolymp_newsletter_recipient_service_proto_rawDescGZIP(), []int{3}
}

func (x *ImportRecipientsOutput) GetImportedTotal() string {
	if x != nil {
		return x.ImportedTotal
	}
	return ""
}

type RemoveRecipientInput struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	RecipientId   string                 `protobuf:"bytes,2,opt,name=recipient_id,json=recipientId,proto3" json:"recipient_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RemoveRecipientInput) Reset() {
	*x = RemoveRecipientInput{}
	mi := &file_eolymp_newsletter_recipient_service_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RemoveRecipientInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveRecipientInput) ProtoMessage() {}

func (x *RemoveRecipientInput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_newsletter_recipient_service_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveRecipientInput.ProtoReflect.Descriptor instead.
func (*RemoveRecipientInput) Descriptor() ([]byte, []int) {
	return file_eolymp_newsletter_recipient_service_proto_rawDescGZIP(), []int{4}
}

func (x *RemoveRecipientInput) GetRecipientId() string {
	if x != nil {
		return x.RecipientId
	}
	return ""
}

type RemoveRecipientOutput struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RemoveRecipientOutput) Reset() {
	*x = RemoveRecipientOutput{}
	mi := &file_eolymp_newsletter_recipient_service_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RemoveRecipientOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveRecipientOutput) ProtoMessage() {}

func (x *RemoveRecipientOutput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_newsletter_recipient_service_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveRecipientOutput.ProtoReflect.Descriptor instead.
func (*RemoveRecipientOutput) Descriptor() ([]byte, []int) {
	return file_eolymp_newsletter_recipient_service_proto_rawDescGZIP(), []int{5}
}

type ListRecipientsInput struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Offset        int32                  `protobuf:"varint,10,opt,name=offset,proto3" json:"offset,omitempty"`
	Size          int32                  `protobuf:"varint,11,opt,name=size,proto3" json:"size,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListRecipientsInput) Reset() {
	*x = ListRecipientsInput{}
	mi := &file_eolymp_newsletter_recipient_service_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListRecipientsInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListRecipientsInput) ProtoMessage() {}

func (x *ListRecipientsInput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_newsletter_recipient_service_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListRecipientsInput.ProtoReflect.Descriptor instead.
func (*ListRecipientsInput) Descriptor() ([]byte, []int) {
	return file_eolymp_newsletter_recipient_service_proto_rawDescGZIP(), []int{6}
}

func (x *ListRecipientsInput) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *ListRecipientsInput) GetSize() int32 {
	if x != nil {
		return x.Size
	}
	return 0
}

type ListRecipientsOutput struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Total         int32                  `protobuf:"varint,1,opt,name=total,proto3" json:"total,omitempty"`
	Items         []*Recipient           `protobuf:"bytes,2,rep,name=items,proto3" json:"items,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListRecipientsOutput) Reset() {
	*x = ListRecipientsOutput{}
	mi := &file_eolymp_newsletter_recipient_service_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListRecipientsOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListRecipientsOutput) ProtoMessage() {}

func (x *ListRecipientsOutput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_newsletter_recipient_service_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListRecipientsOutput.ProtoReflect.Descriptor instead.
func (*ListRecipientsOutput) Descriptor() ([]byte, []int) {
	return file_eolymp_newsletter_recipient_service_proto_rawDescGZIP(), []int{7}
}

func (x *ListRecipientsOutput) GetTotal() int32 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *ListRecipientsOutput) GetItems() []*Recipient {
	if x != nil {
		return x.Items
	}
	return nil
}

type DescribeRecipientInput struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	RecipientId   string                 `protobuf:"bytes,1,opt,name=recipient_id,json=recipientId,proto3" json:"recipient_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DescribeRecipientInput) Reset() {
	*x = DescribeRecipientInput{}
	mi := &file_eolymp_newsletter_recipient_service_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DescribeRecipientInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescribeRecipientInput) ProtoMessage() {}

func (x *DescribeRecipientInput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_newsletter_recipient_service_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DescribeRecipientInput.ProtoReflect.Descriptor instead.
func (*DescribeRecipientInput) Descriptor() ([]byte, []int) {
	return file_eolymp_newsletter_recipient_service_proto_rawDescGZIP(), []int{8}
}

func (x *DescribeRecipientInput) GetRecipientId() string {
	if x != nil {
		return x.RecipientId
	}
	return ""
}

type DescribeRecipientOutput struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Recipient     *Recipient             `protobuf:"bytes,1,opt,name=recipient,proto3" json:"recipient,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DescribeRecipientOutput) Reset() {
	*x = DescribeRecipientOutput{}
	mi := &file_eolymp_newsletter_recipient_service_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DescribeRecipientOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescribeRecipientOutput) ProtoMessage() {}

func (x *DescribeRecipientOutput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_newsletter_recipient_service_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DescribeRecipientOutput.ProtoReflect.Descriptor instead.
func (*DescribeRecipientOutput) Descriptor() ([]byte, []int) {
	return file_eolymp_newsletter_recipient_service_proto_rawDescGZIP(), []int{9}
}

func (x *DescribeRecipientOutput) GetRecipient() *Recipient {
	if x != nil {
		return x.Recipient
	}
	return nil
}

var File_eolymp_newsletter_recipient_service_proto protoreflect.FileDescriptor

var file_eolymp_newsletter_recipient_service_proto_rawDesc = []byte{
	0x0a, 0x29, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x6e, 0x65, 0x77, 0x73, 0x6c, 0x65, 0x74,
	0x74, 0x65, 0x72, 0x2f, 0x72, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11, 0x65, 0x6f, 0x6c,
	0x79, 0x6d, 0x70, 0x2e, 0x6e, 0x65, 0x77, 0x73, 0x6c, 0x65, 0x74, 0x74, 0x65, 0x72, 0x1a, 0x1d,
	0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x22, 0x65,
	0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2f, 0x72, 0x61, 0x74, 0x65, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x21, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x6e, 0x65, 0x77, 0x73, 0x6c, 0x65,
	0x74, 0x74, 0x65, 0x72, 0x2f, 0x72, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x52, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x3a, 0x0a, 0x09,
	0x72, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1c, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x6e, 0x65, 0x77, 0x73, 0x6c, 0x65, 0x74,
	0x74, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x09, 0x72,
	0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x22, 0x3a, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x4f, 0x75, 0x74, 0x70, 0x75,
	0x74, 0x12, 0x21, 0x0a, 0x0c, 0x72, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x72, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65,
	0x6e, 0x74, 0x49, 0x64, 0x22, 0x3e, 0x0a, 0x15, 0x49, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65,
	0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x1b, 0x0a,
	0x08, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x48,
	0x00, 0x52, 0x07, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x42, 0x08, 0x0a, 0x06, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x22, 0x3f, 0x0a, 0x16, 0x49, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65,
	0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x12, 0x25,
	0x0a, 0x0e, 0x69, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x64, 0x5f, 0x74, 0x6f, 0x74, 0x61, 0x6c,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x69, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x64,
	0x54, 0x6f, 0x74, 0x61, 0x6c, 0x22, 0x39, 0x0a, 0x14, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x52,
	0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x21, 0x0a,
	0x0c, 0x72, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x72, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64,
	0x22, 0x17, 0x0a, 0x15, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x52, 0x65, 0x63, 0x69, 0x70, 0x69,
	0x65, 0x6e, 0x74, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x22, 0x41, 0x0a, 0x13, 0x4c, 0x69, 0x73,
	0x74, 0x52, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x49, 0x6e, 0x70, 0x75, 0x74,
	0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65,
	0x18, 0x0b, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x22, 0x60, 0x0a, 0x14,
	0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x4f, 0x75,
	0x74, 0x70, 0x75, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x32, 0x0a, 0x05, 0x69, 0x74,
	0x65, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x65, 0x6f, 0x6c, 0x79,
	0x6d, 0x70, 0x2e, 0x6e, 0x65, 0x77, 0x73, 0x6c, 0x65, 0x74, 0x74, 0x65, 0x72, 0x2e, 0x52, 0x65,
	0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x22, 0x3b,
	0x0a, 0x16, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x52, 0x65, 0x63, 0x69, 0x70, 0x69,
	0x65, 0x6e, 0x74, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x72, 0x65, 0x63, 0x69,
	0x70, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x72, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x22, 0x55, 0x0a, 0x17, 0x44,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x52, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74,
	0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x12, 0x3a, 0x0a, 0x09, 0x72, 0x65, 0x63, 0x69, 0x70, 0x69,
	0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x65, 0x6f, 0x6c, 0x79,
	0x6d, 0x70, 0x2e, 0x6e, 0x65, 0x77, 0x73, 0x6c, 0x65, 0x74, 0x74, 0x65, 0x72, 0x2e, 0x52, 0x65,
	0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x09, 0x72, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65,
	0x6e, 0x74, 0x32, 0xa4, 0x07, 0x0a, 0x10, 0x52, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0xc0, 0x01, 0x0a, 0x11, 0x44, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x62, 0x65, 0x52, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x12, 0x29, 0x2e,
	0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x6e, 0x65, 0x77, 0x73, 0x6c, 0x65, 0x74, 0x74, 0x65,
	0x72, 0x2e, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x52, 0x65, 0x63, 0x69, 0x70, 0x69,
	0x65, 0x6e, 0x74, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x2a, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d,
	0x70, 0x2e, 0x6e, 0x65, 0x77, 0x73, 0x6c, 0x65, 0x74, 0x74, 0x65, 0x72, 0x2e, 0x44, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x62, 0x65, 0x52, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x4f, 0x75,
	0x74, 0x70, 0x75, 0x74, 0x22, 0x54, 0xea, 0xe2, 0x0a, 0x0b, 0xf5, 0xe2, 0x0a, 0x00, 0x00, 0xa0,
	0x40, 0xf8, 0xe2, 0x0a, 0x32, 0x82, 0xe3, 0x0a, 0x1f, 0x8a, 0xe3, 0x0a, 0x1b, 0x6e, 0x65, 0x77,
	0x73, 0x6c, 0x65, 0x74, 0x74, 0x65, 0x72, 0x3a, 0x6e, 0x65, 0x77, 0x73, 0x6c, 0x65, 0x74, 0x74,
	0x65, 0x72, 0x3a, 0x77, 0x72, 0x69, 0x74, 0x65, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1c, 0x12, 0x1a,
	0x2f, 0x72, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x7b, 0x72, 0x65, 0x63,
	0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x7d, 0x12, 0xb6, 0x01, 0x0a, 0x10, 0x49,
	0x6d, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x12,
	0x28, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x6e, 0x65, 0x77, 0x73, 0x6c, 0x65, 0x74,
	0x74, 0x65, 0x72, 0x2e, 0x49, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x63, 0x69, 0x70, 0x69,
	0x65, 0x6e, 0x74, 0x73, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x29, 0x2e, 0x65, 0x6f, 0x6c, 0x79,
	0x6d, 0x70, 0x2e, 0x6e, 0x65, 0x77, 0x73, 0x6c, 0x65, 0x74, 0x74, 0x65, 0x72, 0x2e, 0x49, 0x6d,
	0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x4f, 0x75,
	0x74, 0x70, 0x75, 0x74, 0x22, 0x4d, 0xea, 0xe2, 0x0a, 0x0b, 0xf5, 0xe2, 0x0a, 0x00, 0x00, 0xa0,
	0x40, 0xf8, 0xe2, 0x0a, 0x32, 0x82, 0xe3, 0x0a, 0x1f, 0x8a, 0xe3, 0x0a, 0x1b, 0x6e, 0x65, 0x77,
	0x73, 0x6c, 0x65, 0x74, 0x74, 0x65, 0x72, 0x3a, 0x6e, 0x65, 0x77, 0x73, 0x6c, 0x65, 0x74, 0x74,
	0x65, 0x72, 0x3a, 0x77, 0x72, 0x69, 0x74, 0x65, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x15, 0x22, 0x13,
	0x2f, 0x69, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x2f, 0x72, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65,
	0x6e, 0x74, 0x73, 0x12, 0xa8, 0x01, 0x0a, 0x0e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x63, 0x69,
	0x70, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x26, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e,
	0x6e, 0x65, 0x77, 0x73, 0x6c, 0x65, 0x74, 0x74, 0x65, 0x72, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52,
	0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x27,
	0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x6e, 0x65, 0x77, 0x73, 0x6c, 0x65, 0x74, 0x74,
	0x65, 0x72, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74,
	0x73, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x22, 0x45, 0xea, 0xe2, 0x0a, 0x0b, 0xf5, 0xe2, 0x0a,
	0x00, 0x00, 0xa0, 0x40, 0xf8, 0xe2, 0x0a, 0x32, 0x82, 0xe3, 0x0a, 0x1f, 0x8a, 0xe3, 0x0a, 0x1b,
	0x6e, 0x65, 0x77, 0x73, 0x6c, 0x65, 0x74, 0x74, 0x65, 0x72, 0x3a, 0x6e, 0x65, 0x77, 0x73, 0x6c,
	0x65, 0x74, 0x74, 0x65, 0x72, 0x3a, 0x77, 0x72, 0x69, 0x74, 0x65, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x0d, 0x12, 0x0b, 0x2f, 0x72, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x12, 0xab,
	0x01, 0x0a, 0x0f, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65,
	0x6e, 0x74, 0x12, 0x27, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x6e, 0x65, 0x77, 0x73,
	0x6c, 0x65, 0x74, 0x74, 0x65, 0x72, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x63,
	0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x28, 0x2e, 0x65, 0x6f,
	0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x6e, 0x65, 0x77, 0x73, 0x6c, 0x65, 0x74, 0x74, 0x65, 0x72, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x4f,
	0x75, 0x74, 0x70, 0x75, 0x74, 0x22, 0x45, 0xea, 0xe2, 0x0a, 0x0b, 0xf5, 0xe2, 0x0a, 0x00, 0x00,
	0xa0, 0x40, 0xf8, 0xe2, 0x0a, 0x32, 0x82, 0xe3, 0x0a, 0x1f, 0x8a, 0xe3, 0x0a, 0x1b, 0x6e, 0x65,
	0x77, 0x73, 0x6c, 0x65, 0x74, 0x74, 0x65, 0x72, 0x3a, 0x6e, 0x65, 0x77, 0x73, 0x6c, 0x65, 0x74,
	0x74, 0x65, 0x72, 0x3a, 0x77, 0x72, 0x69, 0x74, 0x65, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0d, 0x22,
	0x0b, 0x2f, 0x72, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x12, 0xba, 0x01, 0x0a,
	0x0f, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x52, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74,
	0x12, 0x27, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x6e, 0x65, 0x77, 0x73, 0x6c, 0x65,
	0x74, 0x74, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x52, 0x65, 0x63, 0x69, 0x70,
	0x69, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x28, 0x2e, 0x65, 0x6f, 0x6c, 0x79,
	0x6d, 0x70, 0x2e, 0x6e, 0x65, 0x77, 0x73, 0x6c, 0x65, 0x74, 0x74, 0x65, 0x72, 0x2e, 0x52, 0x65,
	0x6d, 0x6f, 0x76, 0x65, 0x52, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x4f, 0x75, 0x74,
	0x70, 0x75, 0x74, 0x22, 0x54, 0xea, 0xe2, 0x0a, 0x0b, 0xf5, 0xe2, 0x0a, 0x00, 0x00, 0xa0, 0x40,
	0xf8, 0xe2, 0x0a, 0x32, 0x82, 0xe3, 0x0a, 0x1f, 0x8a, 0xe3, 0x0a, 0x1b, 0x6e, 0x65, 0x77, 0x73,
	0x6c, 0x65, 0x74, 0x74, 0x65, 0x72, 0x3a, 0x6e, 0x65, 0x77, 0x73, 0x6c, 0x65, 0x74, 0x74, 0x65,
	0x72, 0x3a, 0x77, 0x72, 0x69, 0x74, 0x65, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1c, 0x2a, 0x1a, 0x2f,
	0x72, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x7b, 0x72, 0x65, 0x63, 0x69,
	0x70, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x7d, 0x42, 0x37, 0x5a, 0x35, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x67,
	0x6f, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x6e, 0x65, 0x77,
	0x73, 0x6c, 0x65, 0x74, 0x74, 0x65, 0x72, 0x3b, 0x6e, 0x65, 0x77, 0x73, 0x6c, 0x65, 0x74, 0x74,
	0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_eolymp_newsletter_recipient_service_proto_rawDescOnce sync.Once
	file_eolymp_newsletter_recipient_service_proto_rawDescData = file_eolymp_newsletter_recipient_service_proto_rawDesc
)

func file_eolymp_newsletter_recipient_service_proto_rawDescGZIP() []byte {
	file_eolymp_newsletter_recipient_service_proto_rawDescOnce.Do(func() {
		file_eolymp_newsletter_recipient_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_eolymp_newsletter_recipient_service_proto_rawDescData)
	})
	return file_eolymp_newsletter_recipient_service_proto_rawDescData
}

var file_eolymp_newsletter_recipient_service_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_eolymp_newsletter_recipient_service_proto_goTypes = []any{
	(*CreateRecipientInput)(nil),    // 0: eolymp.newsletter.CreateRecipientInput
	(*CreateRecipientOutput)(nil),   // 1: eolymp.newsletter.CreateRecipientOutput
	(*ImportRecipientsInput)(nil),   // 2: eolymp.newsletter.ImportRecipientsInput
	(*ImportRecipientsOutput)(nil),  // 3: eolymp.newsletter.ImportRecipientsOutput
	(*RemoveRecipientInput)(nil),    // 4: eolymp.newsletter.RemoveRecipientInput
	(*RemoveRecipientOutput)(nil),   // 5: eolymp.newsletter.RemoveRecipientOutput
	(*ListRecipientsInput)(nil),     // 6: eolymp.newsletter.ListRecipientsInput
	(*ListRecipientsOutput)(nil),    // 7: eolymp.newsletter.ListRecipientsOutput
	(*DescribeRecipientInput)(nil),  // 8: eolymp.newsletter.DescribeRecipientInput
	(*DescribeRecipientOutput)(nil), // 9: eolymp.newsletter.DescribeRecipientOutput
	(*Recipient)(nil),               // 10: eolymp.newsletter.Recipient
}
var file_eolymp_newsletter_recipient_service_proto_depIdxs = []int32{
	10, // 0: eolymp.newsletter.CreateRecipientInput.recipient:type_name -> eolymp.newsletter.Recipient
	10, // 1: eolymp.newsletter.ListRecipientsOutput.items:type_name -> eolymp.newsletter.Recipient
	10, // 2: eolymp.newsletter.DescribeRecipientOutput.recipient:type_name -> eolymp.newsletter.Recipient
	8,  // 3: eolymp.newsletter.RecipientService.DescribeRecipient:input_type -> eolymp.newsletter.DescribeRecipientInput
	2,  // 4: eolymp.newsletter.RecipientService.ImportRecipients:input_type -> eolymp.newsletter.ImportRecipientsInput
	6,  // 5: eolymp.newsletter.RecipientService.ListRecipients:input_type -> eolymp.newsletter.ListRecipientsInput
	0,  // 6: eolymp.newsletter.RecipientService.CreateRecipient:input_type -> eolymp.newsletter.CreateRecipientInput
	4,  // 7: eolymp.newsletter.RecipientService.RemoveRecipient:input_type -> eolymp.newsletter.RemoveRecipientInput
	9,  // 8: eolymp.newsletter.RecipientService.DescribeRecipient:output_type -> eolymp.newsletter.DescribeRecipientOutput
	3,  // 9: eolymp.newsletter.RecipientService.ImportRecipients:output_type -> eolymp.newsletter.ImportRecipientsOutput
	7,  // 10: eolymp.newsletter.RecipientService.ListRecipients:output_type -> eolymp.newsletter.ListRecipientsOutput
	1,  // 11: eolymp.newsletter.RecipientService.CreateRecipient:output_type -> eolymp.newsletter.CreateRecipientOutput
	5,  // 12: eolymp.newsletter.RecipientService.RemoveRecipient:output_type -> eolymp.newsletter.RemoveRecipientOutput
	8,  // [8:13] is the sub-list for method output_type
	3,  // [3:8] is the sub-list for method input_type
	3,  // [3:3] is the sub-list for extension type_name
	3,  // [3:3] is the sub-list for extension extendee
	0,  // [0:3] is the sub-list for field type_name
}

func init() { file_eolymp_newsletter_recipient_service_proto_init() }
func file_eolymp_newsletter_recipient_service_proto_init() {
	if File_eolymp_newsletter_recipient_service_proto != nil {
		return
	}
	file_eolymp_newsletter_recipient_proto_init()
	file_eolymp_newsletter_recipient_service_proto_msgTypes[2].OneofWrappers = []any{
		(*ImportRecipientsInput_GroupId)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_eolymp_newsletter_recipient_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_eolymp_newsletter_recipient_service_proto_goTypes,
		DependencyIndexes: file_eolymp_newsletter_recipient_service_proto_depIdxs,
		MessageInfos:      file_eolymp_newsletter_recipient_service_proto_msgTypes,
	}.Build()
	File_eolymp_newsletter_recipient_service_proto = out.File
	file_eolymp_newsletter_recipient_service_proto_rawDesc = nil
	file_eolymp_newsletter_recipient_service_proto_goTypes = nil
	file_eolymp_newsletter_recipient_service_proto_depIdxs = nil
}
