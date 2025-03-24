// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v4.24.4
// source: eolymp/atlas/attachment_service.proto

package atlas

import (
	_ "github.com/eolymp/go-sdk/eolymp/annotations"
	wellknown "github.com/eolymp/go-sdk/eolymp/wellknown"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
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

type CreateAttachmentInput struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Attachment    *Attachment            `protobuf:"bytes,2,opt,name=attachment,proto3" json:"attachment,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateAttachmentInput) Reset() {
	*x = CreateAttachmentInput{}
	mi := &file_eolymp_atlas_attachment_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateAttachmentInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAttachmentInput) ProtoMessage() {}

func (x *CreateAttachmentInput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_atlas_attachment_service_proto_msgTypes[0]
	if x != nil {
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

func (x *CreateAttachmentInput) GetAttachment() *Attachment {
	if x != nil {
		return x.Attachment
	}
	return nil
}

type CreateAttachmentOutput struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	AttachmentId  string                 `protobuf:"bytes,1,opt,name=attachment_id,json=attachmentId,proto3" json:"attachment_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateAttachmentOutput) Reset() {
	*x = CreateAttachmentOutput{}
	mi := &file_eolymp_atlas_attachment_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateAttachmentOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAttachmentOutput) ProtoMessage() {}

func (x *CreateAttachmentOutput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_atlas_attachment_service_proto_msgTypes[1]
	if x != nil {
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
	state         protoimpl.MessageState `protogen:"open.v1"`
	AttachmentId  string                 `protobuf:"bytes,2,opt,name=attachment_id,json=attachmentId,proto3" json:"attachment_id,omitempty"`
	Attachment    *Attachment            `protobuf:"bytes,3,opt,name=attachment,proto3" json:"attachment,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateAttachmentInput) Reset() {
	*x = UpdateAttachmentInput{}
	mi := &file_eolymp_atlas_attachment_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateAttachmentInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateAttachmentInput) ProtoMessage() {}

func (x *UpdateAttachmentInput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_atlas_attachment_service_proto_msgTypes[2]
	if x != nil {
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
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateAttachmentOutput) Reset() {
	*x = UpdateAttachmentOutput{}
	mi := &file_eolymp_atlas_attachment_service_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateAttachmentOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateAttachmentOutput) ProtoMessage() {}

func (x *UpdateAttachmentOutput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_atlas_attachment_service_proto_msgTypes[3]
	if x != nil {
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
	state         protoimpl.MessageState `protogen:"open.v1"`
	AttachmentId  string                 `protobuf:"bytes,2,opt,name=attachment_id,json=attachmentId,proto3" json:"attachment_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteAttachmentInput) Reset() {
	*x = DeleteAttachmentInput{}
	mi := &file_eolymp_atlas_attachment_service_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteAttachmentInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteAttachmentInput) ProtoMessage() {}

func (x *DeleteAttachmentInput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_atlas_attachment_service_proto_msgTypes[4]
	if x != nil {
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

func (x *DeleteAttachmentInput) GetAttachmentId() string {
	if x != nil {
		return x.AttachmentId
	}
	return ""
}

type DeleteAttachmentOutput struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteAttachmentOutput) Reset() {
	*x = DeleteAttachmentOutput{}
	mi := &file_eolymp_atlas_attachment_service_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteAttachmentOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteAttachmentOutput) ProtoMessage() {}

func (x *DeleteAttachmentOutput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_atlas_attachment_service_proto_msgTypes[5]
	if x != nil {
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
	state protoimpl.MessageState `protogen:"open.v1"`
	// pagination
	Offset int32 `protobuf:"varint,10,opt,name=offset,proto3" json:"offset,omitempty"`
	Size   int32 `protobuf:"varint,11,opt,name=size,proto3" json:"size,omitempty"`
	// data filters
	Filters *ListAttachmentsInput_Filter `protobuf:"bytes,40,opt,name=filters,proto3" json:"filters,omitempty"`
	// request data for specific problem version
	Version       uint32 `protobuf:"varint,100,opt,name=version,proto3" json:"version,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListAttachmentsInput) Reset() {
	*x = ListAttachmentsInput{}
	mi := &file_eolymp_atlas_attachment_service_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListAttachmentsInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListAttachmentsInput) ProtoMessage() {}

func (x *ListAttachmentsInput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_atlas_attachment_service_proto_msgTypes[6]
	if x != nil {
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
	state         protoimpl.MessageState `protogen:"open.v1"`
	Total         int32                  `protobuf:"varint,1,opt,name=total,proto3" json:"total,omitempty"`
	Items         []*Attachment          `protobuf:"bytes,2,rep,name=items,proto3" json:"items,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListAttachmentsOutput) Reset() {
	*x = ListAttachmentsOutput{}
	mi := &file_eolymp_atlas_attachment_service_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListAttachmentsOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListAttachmentsOutput) ProtoMessage() {}

func (x *ListAttachmentsOutput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_atlas_attachment_service_proto_msgTypes[7]
	if x != nil {
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
	state         protoimpl.MessageState `protogen:"open.v1"`
	AttachmentId  string                 `protobuf:"bytes,2,opt,name=attachment_id,json=attachmentId,proto3" json:"attachment_id,omitempty"`
	Version       uint32                 `protobuf:"varint,100,opt,name=version,proto3" json:"version,omitempty"` // request data for specific problem version
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DescribeAttachmentInput) Reset() {
	*x = DescribeAttachmentInput{}
	mi := &file_eolymp_atlas_attachment_service_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DescribeAttachmentInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescribeAttachmentInput) ProtoMessage() {}

func (x *DescribeAttachmentInput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_atlas_attachment_service_proto_msgTypes[8]
	if x != nil {
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
	state         protoimpl.MessageState `protogen:"open.v1"`
	Attachment    *Attachment            `protobuf:"bytes,1,opt,name=attachment,proto3" json:"attachment,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DescribeAttachmentOutput) Reset() {
	*x = DescribeAttachmentOutput{}
	mi := &file_eolymp_atlas_attachment_service_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DescribeAttachmentOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescribeAttachmentOutput) ProtoMessage() {}

func (x *DescribeAttachmentOutput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_atlas_attachment_service_proto_msgTypes[9]
	if x != nil {
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
	state         protoimpl.MessageState        `protogen:"open.v1"`
	Id            []*wellknown.ExpressionID     `protobuf:"bytes,1,rep,name=id,proto3" json:"id,omitempty"`
	Name          []*wellknown.ExpressionString `protobuf:"bytes,2,rep,name=name,proto3" json:"name,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListAttachmentsInput_Filter) Reset() {
	*x = ListAttachmentsInput_Filter{}
	mi := &file_eolymp_atlas_attachment_service_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListAttachmentsInput_Filter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListAttachmentsInput_Filter) ProtoMessage() {}

func (x *ListAttachmentsInput_Filter) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_atlas_attachment_service_proto_msgTypes[10]
	if x != nil {
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

const file_eolymp_atlas_attachment_service_proto_rawDesc = "" +
	"\n" +
	"%eolymp/atlas/attachment_service.proto\x12\feolymp.atlas\x1a\x1deolymp/annotations/http.proto\x1a\"eolymp/annotations/ratelimit.proto\x1a\x1eeolymp/annotations/scope.proto\x1a\x1deolymp/atlas/attachment.proto\x1a!eolymp/wellknown/expression.proto\"Q\n" +
	"\x15CreateAttachmentInput\x128\n" +
	"\n" +
	"attachment\x18\x02 \x01(\v2\x18.eolymp.atlas.AttachmentR\n" +
	"attachment\"=\n" +
	"\x16CreateAttachmentOutput\x12#\n" +
	"\rattachment_id\x18\x01 \x01(\tR\fattachmentId\"v\n" +
	"\x15UpdateAttachmentInput\x12#\n" +
	"\rattachment_id\x18\x02 \x01(\tR\fattachmentId\x128\n" +
	"\n" +
	"attachment\x18\x03 \x01(\v2\x18.eolymp.atlas.AttachmentR\n" +
	"attachment\"\x18\n" +
	"\x16UpdateAttachmentOutput\"<\n" +
	"\x15DeleteAttachmentInput\x12#\n" +
	"\rattachment_id\x18\x02 \x01(\tR\fattachmentId\"\x18\n" +
	"\x16DeleteAttachmentOutput\"\x93\x02\n" +
	"\x14ListAttachmentsInput\x12\x16\n" +
	"\x06offset\x18\n" +
	" \x01(\x05R\x06offset\x12\x12\n" +
	"\x04size\x18\v \x01(\x05R\x04size\x12C\n" +
	"\afilters\x18( \x01(\v2).eolymp.atlas.ListAttachmentsInput.FilterR\afilters\x12\x18\n" +
	"\aversion\x18d \x01(\rR\aversion\x1ap\n" +
	"\x06Filter\x12.\n" +
	"\x02id\x18\x01 \x03(\v2\x1e.eolymp.wellknown.ExpressionIDR\x02id\x126\n" +
	"\x04name\x18\x02 \x03(\v2\".eolymp.wellknown.ExpressionStringR\x04name\"]\n" +
	"\x15ListAttachmentsOutput\x12\x14\n" +
	"\x05total\x18\x01 \x01(\x05R\x05total\x12.\n" +
	"\x05items\x18\x02 \x03(\v2\x18.eolymp.atlas.AttachmentR\x05items\"X\n" +
	"\x17DescribeAttachmentInput\x12#\n" +
	"\rattachment_id\x18\x02 \x01(\tR\fattachmentId\x12\x18\n" +
	"\aversion\x18d \x01(\rR\aversion\"T\n" +
	"\x18DescribeAttachmentOutput\x128\n" +
	"\n" +
	"attachment\x18\x01 \x01(\v2\x18.eolymp.atlas.AttachmentR\n" +
	"attachment2\xe4\x06\n" +
	"\x11AttachmentService\x12\x9d\x01\n" +
	"\x10CreateAttachment\x12#.eolymp.atlas.CreateAttachmentInput\x1a$.eolymp.atlas.CreateAttachmentOutput\">\xea\xe2\n" +
	"\v\xf5\xe2\n" +
	"\n" +
	"\xd7#>\xf8\xe2\n" +
	"\x05\x82\xe3\n" +
	"\x17\x8a\xe3\n" +
	"\x13atlas:problem:write\x82\xd3\xe4\x93\x02\x0e\"\f/attachments\x12\xad\x01\n" +
	"\x10UpdateAttachment\x12#.eolymp.atlas.UpdateAttachmentInput\x1a$.eolymp.atlas.UpdateAttachmentOutput\"N\xea\xe2\n" +
	"\v\xf5\xe2\n" +
	"\n" +
	"\xd7#>\xf8\xe2\n" +
	"\x05\x82\xe3\n" +
	"\x17\x8a\xe3\n" +
	"\x13atlas:problem:write\x82\xd3\xe4\x93\x02\x1e\"\x1c/attachments/{attachment_id}\x12\xad\x01\n" +
	"\x10DeleteAttachment\x12#.eolymp.atlas.DeleteAttachmentInput\x1a$.eolymp.atlas.DeleteAttachmentOutput\"N\xea\xe2\n" +
	"\v\xf5\xe2\n" +
	"\n" +
	"\xd7#>\xf8\xe2\n" +
	"\x05\x82\xe3\n" +
	"\x17\x8a\xe3\n" +
	"\x13atlas:problem:write\x82\xd3\xe4\x93\x02\x1e*\x1c/attachments/{attachment_id}\x12\x99\x01\n" +
	"\x0fListAttachments\x12\".eolymp.atlas.ListAttachmentsInput\x1a#.eolymp.atlas.ListAttachmentsOutput\"=\xea\xe2\n" +
	"\v\xf5\xe2\n" +
	"\x00\x00\xa0A\xf8\xe2\n" +
	"d\x82\xe3\n" +
	"\x16\x8a\xe3\n" +
	"\x12atlas:problem:read\x82\xd3\xe4\x93\x02\x0e\x12\f/attachments\x12\xb2\x01\n" +
	"\x12DescribeAttachment\x12%.eolymp.atlas.DescribeAttachmentInput\x1a&.eolymp.atlas.DescribeAttachmentOutput\"M\xea\xe2\n" +
	"\v\xf5\xe2\n" +
	"\x00\x00\xa0A\xf8\xe2\n" +
	"d\x82\xe3\n" +
	"\x16\x8a\xe3\n" +
	"\x12atlas:problem:read\x82\xd3\xe4\x93\x02\x1e\x12\x1c/attachments/{attachment_id}B-Z+github.com/eolymp/go-sdk/eolymp/atlas;atlasb\x06proto3"

var (
	file_eolymp_atlas_attachment_service_proto_rawDescOnce sync.Once
	file_eolymp_atlas_attachment_service_proto_rawDescData []byte
)

func file_eolymp_atlas_attachment_service_proto_rawDescGZIP() []byte {
	file_eolymp_atlas_attachment_service_proto_rawDescOnce.Do(func() {
		file_eolymp_atlas_attachment_service_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_eolymp_atlas_attachment_service_proto_rawDesc), len(file_eolymp_atlas_attachment_service_proto_rawDesc)))
	})
	return file_eolymp_atlas_attachment_service_proto_rawDescData
}

var file_eolymp_atlas_attachment_service_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_eolymp_atlas_attachment_service_proto_goTypes = []any{
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_eolymp_atlas_attachment_service_proto_rawDesc), len(file_eolymp_atlas_attachment_service_proto_rawDesc)),
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
	file_eolymp_atlas_attachment_service_proto_goTypes = nil
	file_eolymp_atlas_attachment_service_proto_depIdxs = nil
}
