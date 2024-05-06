// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.0
// 	protoc        v3.21.12
// source: eolymp/atlas/attachment_service.proto

package atlas

import (
	_ "github.com/eolymp/go-sdk/eolymp/annotations"
	wellknown "github.com/eolymp/go-sdk/eolymp/wellknown"
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

type CreateAttachmentInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProblemId  string      `protobuf:"bytes,1,opt,name=problem_id,json=problemId,proto3" json:"problem_id,omitempty"`
	Attachment *Attachment `protobuf:"bytes,2,opt,name=attachment,proto3" json:"attachment,omitempty"`
}

func (x *CreateAttachmentInput) Reset() {
	*x = CreateAttachmentInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_atlas_attachment_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateAttachmentInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAttachmentInput) ProtoMessage() {}

func (x *CreateAttachmentInput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_atlas_attachment_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateAttachmentInput.ProtoReflect.Descriptor instead.
func (*CreateAttachmentInput) Descriptor() ([]byte, []int) {
	return file_eolymp_atlas_attachment_service_proto_rawDescGZIP(), []int{0}
}

func (x *CreateAttachmentInput) GetProblemId() string {
	if x != nil {
		return x.ProblemId
	}
	return ""
}

func (x *CreateAttachmentInput) GetAttachment() *Attachment {
	if x != nil {
		return x.Attachment
	}
	return nil
}

type CreateAttachmentOutput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AttachmentId string `protobuf:"bytes,1,opt,name=attachment_id,json=attachmentId,proto3" json:"attachment_id,omitempty"`
}

func (x *CreateAttachmentOutput) Reset() {
	*x = CreateAttachmentOutput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_atlas_attachment_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateAttachmentOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAttachmentOutput) ProtoMessage() {}

func (x *CreateAttachmentOutput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_atlas_attachment_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateAttachmentOutput.ProtoReflect.Descriptor instead.
func (*CreateAttachmentOutput) Descriptor() ([]byte, []int) {
	return file_eolymp_atlas_attachment_service_proto_rawDescGZIP(), []int{1}
}

func (x *CreateAttachmentOutput) GetAttachmentId() string {
	if x != nil {
		return x.AttachmentId
	}
	return ""
}

type UpdateAttachmentInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProblemId    string      `protobuf:"bytes,1,opt,name=problem_id,json=problemId,proto3" json:"problem_id,omitempty"`
	AttachmentId string      `protobuf:"bytes,2,opt,name=attachment_id,json=attachmentId,proto3" json:"attachment_id,omitempty"`
	Attachment   *Attachment `protobuf:"bytes,3,opt,name=attachment,proto3" json:"attachment,omitempty"`
}

func (x *UpdateAttachmentInput) Reset() {
	*x = UpdateAttachmentInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_atlas_attachment_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateAttachmentInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateAttachmentInput) ProtoMessage() {}

func (x *UpdateAttachmentInput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_atlas_attachment_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateAttachmentInput.ProtoReflect.Descriptor instead.
func (*UpdateAttachmentInput) Descriptor() ([]byte, []int) {
	return file_eolymp_atlas_attachment_service_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateAttachmentInput) GetProblemId() string {
	if x != nil {
		return x.ProblemId
	}
	return ""
}

func (x *UpdateAttachmentInput) GetAttachmentId() string {
	if x != nil {
		return x.AttachmentId
	}
	return ""
}

func (x *UpdateAttachmentInput) GetAttachment() *Attachment {
	if x != nil {
		return x.Attachment
	}
	return nil
}

type UpdateAttachmentOutput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateAttachmentOutput) Reset() {
	*x = UpdateAttachmentOutput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_atlas_attachment_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateAttachmentOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateAttachmentOutput) ProtoMessage() {}

func (x *UpdateAttachmentOutput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_atlas_attachment_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateAttachmentOutput.ProtoReflect.Descriptor instead.
func (*UpdateAttachmentOutput) Descriptor() ([]byte, []int) {
	return file_eolymp_atlas_attachment_service_proto_rawDescGZIP(), []int{3}
}

type DeleteAttachmentInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProblemId    string `protobuf:"bytes,1,opt,name=problem_id,json=problemId,proto3" json:"problem_id,omitempty"`
	AttachmentId string `protobuf:"bytes,2,opt,name=attachment_id,json=attachmentId,proto3" json:"attachment_id,omitempty"`
}

func (x *DeleteAttachmentInput) Reset() {
	*x = DeleteAttachmentInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_atlas_attachment_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteAttachmentInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteAttachmentInput) ProtoMessage() {}

func (x *DeleteAttachmentInput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_atlas_attachment_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteAttachmentInput.ProtoReflect.Descriptor instead.
func (*DeleteAttachmentInput) Descriptor() ([]byte, []int) {
	return file_eolymp_atlas_attachment_service_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteAttachmentInput) GetProblemId() string {
	if x != nil {
		return x.ProblemId
	}
	return ""
}

func (x *DeleteAttachmentInput) GetAttachmentId() string {
	if x != nil {
		return x.AttachmentId
	}
	return ""
}

type DeleteAttachmentOutput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteAttachmentOutput) Reset() {
	*x = DeleteAttachmentOutput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_atlas_attachment_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteAttachmentOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteAttachmentOutput) ProtoMessage() {}

func (x *DeleteAttachmentOutput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_atlas_attachment_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteAttachmentOutput.ProtoReflect.Descriptor instead.
func (*DeleteAttachmentOutput) Descriptor() ([]byte, []int) {
	return file_eolymp_atlas_attachment_service_proto_rawDescGZIP(), []int{5}
}

type ListAttachmentsInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProblemId string `protobuf:"bytes,1,opt,name=problem_id,json=problemId,proto3" json:"problem_id,omitempty"`
	// pagination
	Offset int32 `protobuf:"varint,10,opt,name=offset,proto3" json:"offset,omitempty"`
	Size   int32 `protobuf:"varint,11,opt,name=size,proto3" json:"size,omitempty"`
	// data filters
	Filters *ListAttachmentsInput_Filter `protobuf:"bytes,40,opt,name=filters,proto3" json:"filters,omitempty"`
	// request data for specific problem version
	Version uint32 `protobuf:"varint,100,opt,name=version,proto3" json:"version,omitempty"`
}

func (x *ListAttachmentsInput) Reset() {
	*x = ListAttachmentsInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_atlas_attachment_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListAttachmentsInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListAttachmentsInput) ProtoMessage() {}

func (x *ListAttachmentsInput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_atlas_attachment_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListAttachmentsInput.ProtoReflect.Descriptor instead.
func (*ListAttachmentsInput) Descriptor() ([]byte, []int) {
	return file_eolymp_atlas_attachment_service_proto_rawDescGZIP(), []int{6}
}

func (x *ListAttachmentsInput) GetProblemId() string {
	if x != nil {
		return x.ProblemId
	}
	return ""
}

func (x *ListAttachmentsInput) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *ListAttachmentsInput) GetSize() int32 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *ListAttachmentsInput) GetFilters() *ListAttachmentsInput_Filter {
	if x != nil {
		return x.Filters
	}
	return nil
}

func (x *ListAttachmentsInput) GetVersion() uint32 {
	if x != nil {
		return x.Version
	}
	return 0
}

type ListAttachmentsOutput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Total int32         `protobuf:"varint,1,opt,name=total,proto3" json:"total,omitempty"`
	Items []*Attachment `protobuf:"bytes,2,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *ListAttachmentsOutput) Reset() {
	*x = ListAttachmentsOutput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_atlas_attachment_service_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListAttachmentsOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListAttachmentsOutput) ProtoMessage() {}

func (x *ListAttachmentsOutput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_atlas_attachment_service_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListAttachmentsOutput.ProtoReflect.Descriptor instead.
func (*ListAttachmentsOutput) Descriptor() ([]byte, []int) {
	return file_eolymp_atlas_attachment_service_proto_rawDescGZIP(), []int{7}
}

func (x *ListAttachmentsOutput) GetTotal() int32 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *ListAttachmentsOutput) GetItems() []*Attachment {
	if x != nil {
		return x.Items
	}
	return nil
}

type DescribeAttachmentInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProblemId    string `protobuf:"bytes,1,opt,name=problem_id,json=problemId,proto3" json:"problem_id,omitempty"`
	AttachmentId string `protobuf:"bytes,2,opt,name=attachment_id,json=attachmentId,proto3" json:"attachment_id,omitempty"`
	Version      uint32 `protobuf:"varint,100,opt,name=version,proto3" json:"version,omitempty"` // request data for specific problem version
}

func (x *DescribeAttachmentInput) Reset() {
	*x = DescribeAttachmentInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_atlas_attachment_service_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DescribeAttachmentInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescribeAttachmentInput) ProtoMessage() {}

func (x *DescribeAttachmentInput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_atlas_attachment_service_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DescribeAttachmentInput.ProtoReflect.Descriptor instead.
func (*DescribeAttachmentInput) Descriptor() ([]byte, []int) {
	return file_eolymp_atlas_attachment_service_proto_rawDescGZIP(), []int{8}
}

func (x *DescribeAttachmentInput) GetProblemId() string {
	if x != nil {
		return x.ProblemId
	}
	return ""
}

func (x *DescribeAttachmentInput) GetAttachmentId() string {
	if x != nil {
		return x.AttachmentId
	}
	return ""
}

func (x *DescribeAttachmentInput) GetVersion() uint32 {
	if x != nil {
		return x.Version
	}
	return 0
}

type DescribeAttachmentOutput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Attachment *Attachment `protobuf:"bytes,1,opt,name=attachment,proto3" json:"attachment,omitempty"`
}

func (x *DescribeAttachmentOutput) Reset() {
	*x = DescribeAttachmentOutput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_atlas_attachment_service_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DescribeAttachmentOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescribeAttachmentOutput) ProtoMessage() {}

func (x *DescribeAttachmentOutput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_atlas_attachment_service_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DescribeAttachmentOutput.ProtoReflect.Descriptor instead.
func (*DescribeAttachmentOutput) Descriptor() ([]byte, []int) {
	return file_eolymp_atlas_attachment_service_proto_rawDescGZIP(), []int{9}
}

func (x *DescribeAttachmentOutput) GetAttachment() *Attachment {
	if x != nil {
		return x.Attachment
	}
	return nil
}

type ListAttachmentsInput_Filter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   []*wellknown.ExpressionID     `protobuf:"bytes,1,rep,name=id,proto3" json:"id,omitempty"`
	Name []*wellknown.ExpressionString `protobuf:"bytes,2,rep,name=name,proto3" json:"name,omitempty"`
}

func (x *ListAttachmentsInput_Filter) Reset() {
	*x = ListAttachmentsInput_Filter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_atlas_attachment_service_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListAttachmentsInput_Filter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListAttachmentsInput_Filter) ProtoMessage() {}

func (x *ListAttachmentsInput_Filter) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_atlas_attachment_service_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListAttachmentsInput_Filter.ProtoReflect.Descriptor instead.
func (*ListAttachmentsInput_Filter) Descriptor() ([]byte, []int) {
	return file_eolymp_atlas_attachment_service_proto_rawDescGZIP(), []int{6, 0}
}

func (x *ListAttachmentsInput_Filter) GetId() []*wellknown.ExpressionID {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *ListAttachmentsInput_Filter) GetName() []*wellknown.ExpressionString {
	if x != nil {
		return x.Name
	}
	return nil
}

var File_eolymp_atlas_attachment_service_proto protoreflect.FileDescriptor

var file_eolymp_atlas_attachment_service_proto_rawDesc = []byte{
	0x0a, 0x25, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x61, 0x74, 0x6c, 0x61, 0x73, 0x2f, 0x61,
	0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e,
	0x61, 0x74, 0x6c, 0x61, 0x73, 0x1a, 0x1d, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x61, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x22, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x72, 0x61, 0x74, 0x65, 0x6c, 0x69, 0x6d,
	0x69, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70,
	0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x63, 0x6f,
	0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70,
	0x2f, 0x61, 0x74, 0x6c, 0x61, 0x73, 0x2f, 0x61, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x21, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f,
	0x77, 0x65, 0x6c, 0x6c, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x2f, 0x65, 0x78, 0x70, 0x72, 0x65, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x70, 0x0a, 0x15, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x6e,
	0x70, 0x75, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d,
	0x49, 0x64, 0x12, 0x38, 0x0a, 0x0a, 0x61, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e,
	0x61, 0x74, 0x6c, 0x61, 0x73, 0x2e, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74,
	0x52, 0x0a, 0x61, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0x3d, 0x0a, 0x16,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74,
	0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x61, 0x74, 0x74, 0x61, 0x63, 0x68,
	0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x61,
	0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x22, 0x95, 0x01, 0x0a, 0x15,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74,
	0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x62, 0x6c,
	0x65, 0x6d, 0x49, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x61, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65,
	0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x61, 0x74, 0x74,
	0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x38, 0x0a, 0x0a, 0x61, 0x74, 0x74,
	0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e,
	0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x61, 0x74, 0x6c, 0x61, 0x73, 0x2e, 0x41, 0x74, 0x74,
	0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x0a, 0x61, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d,
	0x65, 0x6e, 0x74, 0x22, 0x18, 0x0a, 0x16, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x74, 0x74,
	0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x22, 0x5b, 0x0a,
	0x15, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e,
	0x74, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x62, 0x6c, 0x65,
	0x6d, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x62,
	0x6c, 0x65, 0x6d, 0x49, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x61, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d,
	0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x61, 0x74,
	0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x22, 0x18, 0x0a, 0x16, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x4f, 0x75,
	0x74, 0x70, 0x75, 0x74, 0x22, 0xb2, 0x02, 0x0a, 0x14, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x74, 0x74,
	0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x1d, 0x0a,
	0x0a, 0x70, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06,
	0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6f, 0x66,
	0x66, 0x73, 0x65, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x0b, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x43, 0x0a, 0x07, 0x66, 0x69, 0x6c, 0x74,
	0x65, 0x72, 0x73, 0x18, 0x28, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x65, 0x6f, 0x6c, 0x79,
	0x6d, 0x70, 0x2e, 0x61, 0x74, 0x6c, 0x61, 0x73, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x74, 0x74,
	0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x2e, 0x46, 0x69,
	0x6c, 0x74, 0x65, 0x72, 0x52, 0x07, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x12, 0x18, 0x0a,
	0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x64, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07,
	0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x1a, 0x70, 0x0a, 0x06, 0x46, 0x69, 0x6c, 0x74, 0x65,
	0x72, 0x12, 0x2e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e,
	0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x77, 0x65, 0x6c, 0x6c, 0x6b, 0x6e, 0x6f, 0x77, 0x6e,
	0x2e, 0x45, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x36, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x22, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x77, 0x65, 0x6c, 0x6c, 0x6b, 0x6e, 0x6f,
	0x77, 0x6e, 0x2e, 0x45, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x72,
	0x69, 0x6e, 0x67, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x5d, 0x0a, 0x15, 0x4c, 0x69, 0x73,
	0x74, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x4f, 0x75, 0x74, 0x70,
	0x75, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x2e, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d,
	0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70,
	0x2e, 0x61, 0x74, 0x6c, 0x61, 0x73, 0x2e, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e,
	0x74, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x22, 0x77, 0x0a, 0x17, 0x44, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x62, 0x65, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x6e,
	0x70, 0x75, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d,
	0x49, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x61, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x61, 0x74, 0x74, 0x61, 0x63,
	0x68, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x18, 0x64, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x22, 0x54, 0x0a, 0x18, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x41, 0x74, 0x74,
	0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x12, 0x38, 0x0a,
	0x0a, 0x61, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x18, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x61, 0x74, 0x6c, 0x61, 0x73,
	0x2e, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x0a, 0x61, 0x74, 0x74,
	0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x32, 0xe4, 0x06, 0x0a, 0x11, 0x41, 0x74, 0x74, 0x61,
	0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x9d, 0x01,
	0x0a, 0x10, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65,
	0x6e, 0x74, 0x12, 0x23, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x61, 0x74, 0x6c, 0x61,
	0x73, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65,
	0x6e, 0x74, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x24, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70,
	0x2e, 0x61, 0x74, 0x6c, 0x61, 0x73, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x74, 0x74,
	0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x22, 0x3e, 0xea,
	0xe2, 0x0a, 0x0b, 0xf5, 0xe2, 0x0a, 0x0a, 0xd7, 0x23, 0x3e, 0xf8, 0xe2, 0x0a, 0x05, 0x82, 0xe3,
	0x0a, 0x17, 0x8a, 0xe3, 0x0a, 0x13, 0x61, 0x74, 0x6c, 0x61, 0x73, 0x3a, 0x70, 0x72, 0x6f, 0x62,
	0x6c, 0x65, 0x6d, 0x3a, 0x77, 0x72, 0x69, 0x74, 0x65, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0e, 0x22,
	0x0c, 0x2f, 0x61, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12, 0xad, 0x01,
	0x0a, 0x10, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65,
	0x6e, 0x74, 0x12, 0x23, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x61, 0x74, 0x6c, 0x61,
	0x73, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65,
	0x6e, 0x74, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x24, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70,
	0x2e, 0x61, 0x74, 0x6c, 0x61, 0x73, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x74, 0x74,
	0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x22, 0x4e, 0xea,
	0xe2, 0x0a, 0x0b, 0xf5, 0xe2, 0x0a, 0x0a, 0xd7, 0x23, 0x3e, 0xf8, 0xe2, 0x0a, 0x05, 0x82, 0xe3,
	0x0a, 0x17, 0x8a, 0xe3, 0x0a, 0x13, 0x61, 0x74, 0x6c, 0x61, 0x73, 0x3a, 0x70, 0x72, 0x6f, 0x62,
	0x6c, 0x65, 0x6d, 0x3a, 0x77, 0x72, 0x69, 0x74, 0x65, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1e, 0x22,
	0x1c, 0x2f, 0x61, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x7b, 0x61,
	0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x7d, 0x12, 0xad, 0x01,
	0x0a, 0x10, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65,
	0x6e, 0x74, 0x12, 0x23, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x61, 0x74, 0x6c, 0x61,
	0x73, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65,
	0x6e, 0x74, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x24, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70,
	0x2e, 0x61, 0x74, 0x6c, 0x61, 0x73, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x74, 0x74,
	0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x22, 0x4e, 0xea,
	0xe2, 0x0a, 0x0b, 0xf5, 0xe2, 0x0a, 0x0a, 0xd7, 0x23, 0x3e, 0xf8, 0xe2, 0x0a, 0x05, 0x82, 0xe3,
	0x0a, 0x17, 0x8a, 0xe3, 0x0a, 0x13, 0x61, 0x74, 0x6c, 0x61, 0x73, 0x3a, 0x70, 0x72, 0x6f, 0x62,
	0x6c, 0x65, 0x6d, 0x3a, 0x77, 0x72, 0x69, 0x74, 0x65, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1e, 0x2a,
	0x1c, 0x2f, 0x61, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x7b, 0x61,
	0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x7d, 0x12, 0x99, 0x01,
	0x0a, 0x0f, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74,
	0x73, 0x12, 0x22, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x61, 0x74, 0x6c, 0x61, 0x73,
	0x2e, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x73,
	0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x23, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x61,
	0x74, 0x6c, 0x61, 0x73, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d,
	0x65, 0x6e, 0x74, 0x73, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x22, 0x3d, 0xea, 0xe2, 0x0a, 0x0b,
	0xf5, 0xe2, 0x0a, 0x00, 0x00, 0xa0, 0x41, 0xf8, 0xe2, 0x0a, 0x64, 0x82, 0xe3, 0x0a, 0x16, 0x8a,
	0xe3, 0x0a, 0x12, 0x61, 0x74, 0x6c, 0x61, 0x73, 0x3a, 0x70, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d,
	0x3a, 0x72, 0x65, 0x61, 0x64, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0e, 0x12, 0x0c, 0x2f, 0x61, 0x74,
	0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12, 0xb2, 0x01, 0x0a, 0x12, 0x44, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74,
	0x12, 0x25, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x61, 0x74, 0x6c, 0x61, 0x73, 0x2e,
	0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65,
	0x6e, 0x74, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x26, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70,
	0x2e, 0x61, 0x74, 0x6c, 0x61, 0x73, 0x2e, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x41,
	0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x22,
	0x4d, 0xea, 0xe2, 0x0a, 0x0b, 0xf5, 0xe2, 0x0a, 0x00, 0x00, 0xa0, 0x41, 0xf8, 0xe2, 0x0a, 0x64,
	0x82, 0xe3, 0x0a, 0x16, 0x8a, 0xe3, 0x0a, 0x12, 0x61, 0x74, 0x6c, 0x61, 0x73, 0x3a, 0x70, 0x72,
	0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x3a, 0x72, 0x65, 0x61, 0x64, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1e,
	0x12, 0x1c, 0x2f, 0x61, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x7b,
	0x61, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x7d, 0x42, 0x2d,
	0x5a, 0x2b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6f, 0x6c,
	0x79, 0x6d, 0x70, 0x2f, 0x67, 0x6f, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x65, 0x6f, 0x6c, 0x79, 0x6d,
	0x70, 0x2f, 0x61, 0x74, 0x6c, 0x61, 0x73, 0x3b, 0x61, 0x74, 0x6c, 0x61, 0x73, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_eolymp_atlas_attachment_service_proto_rawDescOnce sync.Once
	file_eolymp_atlas_attachment_service_proto_rawDescData = file_eolymp_atlas_attachment_service_proto_rawDesc
)

func file_eolymp_atlas_attachment_service_proto_rawDescGZIP() []byte {
	file_eolymp_atlas_attachment_service_proto_rawDescOnce.Do(func() {
		file_eolymp_atlas_attachment_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_eolymp_atlas_attachment_service_proto_rawDescData)
	})
	return file_eolymp_atlas_attachment_service_proto_rawDescData
}

var file_eolymp_atlas_attachment_service_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_eolymp_atlas_attachment_service_proto_goTypes = []interface{}{
	(*CreateAttachmentInput)(nil),       // 0: eolymp.atlas.CreateAttachmentInput
	(*CreateAttachmentOutput)(nil),      // 1: eolymp.atlas.CreateAttachmentOutput
	(*UpdateAttachmentInput)(nil),       // 2: eolymp.atlas.UpdateAttachmentInput
	(*UpdateAttachmentOutput)(nil),      // 3: eolymp.atlas.UpdateAttachmentOutput
	(*DeleteAttachmentInput)(nil),       // 4: eolymp.atlas.DeleteAttachmentInput
	(*DeleteAttachmentOutput)(nil),      // 5: eolymp.atlas.DeleteAttachmentOutput
	(*ListAttachmentsInput)(nil),        // 6: eolymp.atlas.ListAttachmentsInput
	(*ListAttachmentsOutput)(nil),       // 7: eolymp.atlas.ListAttachmentsOutput
	(*DescribeAttachmentInput)(nil),     // 8: eolymp.atlas.DescribeAttachmentInput
	(*DescribeAttachmentOutput)(nil),    // 9: eolymp.atlas.DescribeAttachmentOutput
	(*ListAttachmentsInput_Filter)(nil), // 10: eolymp.atlas.ListAttachmentsInput.Filter
	(*Attachment)(nil),                  // 11: eolymp.atlas.Attachment
	(*wellknown.ExpressionID)(nil),      // 12: eolymp.wellknown.ExpressionID
	(*wellknown.ExpressionString)(nil),  // 13: eolymp.wellknown.ExpressionString
}
var file_eolymp_atlas_attachment_service_proto_depIdxs = []int32{
	11, // 0: eolymp.atlas.CreateAttachmentInput.attachment:type_name -> eolymp.atlas.Attachment
	11, // 1: eolymp.atlas.UpdateAttachmentInput.attachment:type_name -> eolymp.atlas.Attachment
	10, // 2: eolymp.atlas.ListAttachmentsInput.filters:type_name -> eolymp.atlas.ListAttachmentsInput.Filter
	11, // 3: eolymp.atlas.ListAttachmentsOutput.items:type_name -> eolymp.atlas.Attachment
	11, // 4: eolymp.atlas.DescribeAttachmentOutput.attachment:type_name -> eolymp.atlas.Attachment
	12, // 5: eolymp.atlas.ListAttachmentsInput.Filter.id:type_name -> eolymp.wellknown.ExpressionID
	13, // 6: eolymp.atlas.ListAttachmentsInput.Filter.name:type_name -> eolymp.wellknown.ExpressionString
	0,  // 7: eolymp.atlas.AttachmentService.CreateAttachment:input_type -> eolymp.atlas.CreateAttachmentInput
	2,  // 8: eolymp.atlas.AttachmentService.UpdateAttachment:input_type -> eolymp.atlas.UpdateAttachmentInput
	4,  // 9: eolymp.atlas.AttachmentService.DeleteAttachment:input_type -> eolymp.atlas.DeleteAttachmentInput
	6,  // 10: eolymp.atlas.AttachmentService.ListAttachments:input_type -> eolymp.atlas.ListAttachmentsInput
	8,  // 11: eolymp.atlas.AttachmentService.DescribeAttachment:input_type -> eolymp.atlas.DescribeAttachmentInput
	1,  // 12: eolymp.atlas.AttachmentService.CreateAttachment:output_type -> eolymp.atlas.CreateAttachmentOutput
	3,  // 13: eolymp.atlas.AttachmentService.UpdateAttachment:output_type -> eolymp.atlas.UpdateAttachmentOutput
	5,  // 14: eolymp.atlas.AttachmentService.DeleteAttachment:output_type -> eolymp.atlas.DeleteAttachmentOutput
	7,  // 15: eolymp.atlas.AttachmentService.ListAttachments:output_type -> eolymp.atlas.ListAttachmentsOutput
	9,  // 16: eolymp.atlas.AttachmentService.DescribeAttachment:output_type -> eolymp.atlas.DescribeAttachmentOutput
	12, // [12:17] is the sub-list for method output_type
	7,  // [7:12] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_eolymp_atlas_attachment_service_proto_init() }
func file_eolymp_atlas_attachment_service_proto_init() {
	if File_eolymp_atlas_attachment_service_proto != nil {
		return
	}
	file_eolymp_atlas_attachment_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_eolymp_atlas_attachment_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateAttachmentInput); i {
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
		file_eolymp_atlas_attachment_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateAttachmentOutput); i {
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
		file_eolymp_atlas_attachment_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateAttachmentInput); i {
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
		file_eolymp_atlas_attachment_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateAttachmentOutput); i {
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
		file_eolymp_atlas_attachment_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteAttachmentInput); i {
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
		file_eolymp_atlas_attachment_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteAttachmentOutput); i {
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
		file_eolymp_atlas_attachment_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListAttachmentsInput); i {
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
		file_eolymp_atlas_attachment_service_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListAttachmentsOutput); i {
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
		file_eolymp_atlas_attachment_service_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DescribeAttachmentInput); i {
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
		file_eolymp_atlas_attachment_service_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DescribeAttachmentOutput); i {
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
		file_eolymp_atlas_attachment_service_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListAttachmentsInput_Filter); i {
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
			RawDescriptor: file_eolymp_atlas_attachment_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_eolymp_atlas_attachment_service_proto_goTypes,
		DependencyIndexes: file_eolymp_atlas_attachment_service_proto_depIdxs,
		MessageInfos:      file_eolymp_atlas_attachment_service_proto_msgTypes,
	}.Build()
	File_eolymp_atlas_attachment_service_proto = out.File
	file_eolymp_atlas_attachment_service_proto_rawDesc = nil
	file_eolymp_atlas_attachment_service_proto_goTypes = nil
	file_eolymp_atlas_attachment_service_proto_depIdxs = nil
}
