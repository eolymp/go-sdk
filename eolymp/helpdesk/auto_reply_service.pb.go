// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.29.4
// source: eolymp/helpdesk/auto_reply_service.proto

package helpdesk

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

type CreateAutoReplyInput struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Reply         *AutoReply             `protobuf:"bytes,1,opt,name=reply,proto3" json:"reply,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateAutoReplyInput) Reset() {
	*x = CreateAutoReplyInput{}
	mi := &file_eolymp_helpdesk_auto_reply_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateAutoReplyInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAutoReplyInput) ProtoMessage() {}

func (x *CreateAutoReplyInput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_helpdesk_auto_reply_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateAutoReplyInput.ProtoReflect.Descriptor instead.
func (*CreateAutoReplyInput) Descriptor() ([]byte, []int) {
	return file_eolymp_helpdesk_auto_reply_service_proto_rawDescGZIP(), []int{0}
}

func (x *CreateAutoReplyInput) GetReply() *AutoReply {
	if x != nil {
		return x.Reply
	}
	return nil
}

type CreateAutoReplyOutput struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ReplyId       string                 `protobuf:"bytes,1,opt,name=reply_id,json=replyId,proto3" json:"reply_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateAutoReplyOutput) Reset() {
	*x = CreateAutoReplyOutput{}
	mi := &file_eolymp_helpdesk_auto_reply_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateAutoReplyOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAutoReplyOutput) ProtoMessage() {}

func (x *CreateAutoReplyOutput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_helpdesk_auto_reply_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateAutoReplyOutput.ProtoReflect.Descriptor instead.
func (*CreateAutoReplyOutput) Descriptor() ([]byte, []int) {
	return file_eolymp_helpdesk_auto_reply_service_proto_rawDescGZIP(), []int{1}
}

func (x *CreateAutoReplyOutput) GetReplyId() string {
	if x != nil {
		return x.ReplyId
	}
	return ""
}

type UpdateAutoReplyInput struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ReplyId       string                 `protobuf:"bytes,1,opt,name=reply_id,json=replyId,proto3" json:"reply_id,omitempty"`
	Reply         *AutoReply             `protobuf:"bytes,2,opt,name=reply,proto3" json:"reply,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateAutoReplyInput) Reset() {
	*x = UpdateAutoReplyInput{}
	mi := &file_eolymp_helpdesk_auto_reply_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateAutoReplyInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateAutoReplyInput) ProtoMessage() {}

func (x *UpdateAutoReplyInput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_helpdesk_auto_reply_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateAutoReplyInput.ProtoReflect.Descriptor instead.
func (*UpdateAutoReplyInput) Descriptor() ([]byte, []int) {
	return file_eolymp_helpdesk_auto_reply_service_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateAutoReplyInput) GetReplyId() string {
	if x != nil {
		return x.ReplyId
	}
	return ""
}

func (x *UpdateAutoReplyInput) GetReply() *AutoReply {
	if x != nil {
		return x.Reply
	}
	return nil
}

type UpdateAutoReplyOutput struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateAutoReplyOutput) Reset() {
	*x = UpdateAutoReplyOutput{}
	mi := &file_eolymp_helpdesk_auto_reply_service_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateAutoReplyOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateAutoReplyOutput) ProtoMessage() {}

func (x *UpdateAutoReplyOutput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_helpdesk_auto_reply_service_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateAutoReplyOutput.ProtoReflect.Descriptor instead.
func (*UpdateAutoReplyOutput) Descriptor() ([]byte, []int) {
	return file_eolymp_helpdesk_auto_reply_service_proto_rawDescGZIP(), []int{3}
}

type DeleteAutoReplyInput struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ReplyId       string                 `protobuf:"bytes,1,opt,name=reply_id,json=replyId,proto3" json:"reply_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteAutoReplyInput) Reset() {
	*x = DeleteAutoReplyInput{}
	mi := &file_eolymp_helpdesk_auto_reply_service_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteAutoReplyInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteAutoReplyInput) ProtoMessage() {}

func (x *DeleteAutoReplyInput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_helpdesk_auto_reply_service_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteAutoReplyInput.ProtoReflect.Descriptor instead.
func (*DeleteAutoReplyInput) Descriptor() ([]byte, []int) {
	return file_eolymp_helpdesk_auto_reply_service_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteAutoReplyInput) GetReplyId() string {
	if x != nil {
		return x.ReplyId
	}
	return ""
}

type DeleteAutoReplyOutput struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteAutoReplyOutput) Reset() {
	*x = DeleteAutoReplyOutput{}
	mi := &file_eolymp_helpdesk_auto_reply_service_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteAutoReplyOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteAutoReplyOutput) ProtoMessage() {}

func (x *DeleteAutoReplyOutput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_helpdesk_auto_reply_service_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteAutoReplyOutput.ProtoReflect.Descriptor instead.
func (*DeleteAutoReplyOutput) Descriptor() ([]byte, []int) {
	return file_eolymp_helpdesk_auto_reply_service_proto_rawDescGZIP(), []int{5}
}

type ListAutoRepliesInput struct {
	state  protoimpl.MessageState `protogen:"open.v1"`
	Render bool                   `protobuf:"varint,1,opt,name=render,proto3" json:"render,omitempty"`
	// pagination
	Offset int32 `protobuf:"varint,10,opt,name=offset,proto3" json:"offset,omitempty"`
	Size   int32 `protobuf:"varint,11,opt,name=size,proto3" json:"size,omitempty"`
	// data filters
	Filters       *ListAutoRepliesInput_Filter `protobuf:"bytes,40,opt,name=filters,proto3" json:"filters,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListAutoRepliesInput) Reset() {
	*x = ListAutoRepliesInput{}
	mi := &file_eolymp_helpdesk_auto_reply_service_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListAutoRepliesInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListAutoRepliesInput) ProtoMessage() {}

func (x *ListAutoRepliesInput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_helpdesk_auto_reply_service_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListAutoRepliesInput.ProtoReflect.Descriptor instead.
func (*ListAutoRepliesInput) Descriptor() ([]byte, []int) {
	return file_eolymp_helpdesk_auto_reply_service_proto_rawDescGZIP(), []int{6}
}

func (x *ListAutoRepliesInput) GetRender() bool {
	if x != nil {
		return x.Render
	}
	return false
}

func (x *ListAutoRepliesInput) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *ListAutoRepliesInput) GetSize() int32 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *ListAutoRepliesInput) GetFilters() *ListAutoRepliesInput_Filter {
	if x != nil {
		return x.Filters
	}
	return nil
}

type ListAutoRepliesOutput struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Total         int32                  `protobuf:"varint,1,opt,name=total,proto3" json:"total,omitempty"`
	Items         []*AutoReply           `protobuf:"bytes,2,rep,name=items,proto3" json:"items,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListAutoRepliesOutput) Reset() {
	*x = ListAutoRepliesOutput{}
	mi := &file_eolymp_helpdesk_auto_reply_service_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListAutoRepliesOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListAutoRepliesOutput) ProtoMessage() {}

func (x *ListAutoRepliesOutput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_helpdesk_auto_reply_service_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListAutoRepliesOutput.ProtoReflect.Descriptor instead.
func (*ListAutoRepliesOutput) Descriptor() ([]byte, []int) {
	return file_eolymp_helpdesk_auto_reply_service_proto_rawDescGZIP(), []int{7}
}

func (x *ListAutoRepliesOutput) GetTotal() int32 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *ListAutoRepliesOutput) GetItems() []*AutoReply {
	if x != nil {
		return x.Items
	}
	return nil
}

type DescribeAutoReplyInput struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ReplyId       string                 `protobuf:"bytes,1,opt,name=reply_id,json=replyId,proto3" json:"reply_id,omitempty"`
	Render        bool                   `protobuf:"varint,2,opt,name=render,proto3" json:"render,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DescribeAutoReplyInput) Reset() {
	*x = DescribeAutoReplyInput{}
	mi := &file_eolymp_helpdesk_auto_reply_service_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DescribeAutoReplyInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescribeAutoReplyInput) ProtoMessage() {}

func (x *DescribeAutoReplyInput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_helpdesk_auto_reply_service_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DescribeAutoReplyInput.ProtoReflect.Descriptor instead.
func (*DescribeAutoReplyInput) Descriptor() ([]byte, []int) {
	return file_eolymp_helpdesk_auto_reply_service_proto_rawDescGZIP(), []int{8}
}

func (x *DescribeAutoReplyInput) GetReplyId() string {
	if x != nil {
		return x.ReplyId
	}
	return ""
}

func (x *DescribeAutoReplyInput) GetRender() bool {
	if x != nil {
		return x.Render
	}
	return false
}

type DescribeAutoReplyOutput struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Reply         *AutoReply             `protobuf:"bytes,1,opt,name=reply,proto3" json:"reply,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DescribeAutoReplyOutput) Reset() {
	*x = DescribeAutoReplyOutput{}
	mi := &file_eolymp_helpdesk_auto_reply_service_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DescribeAutoReplyOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescribeAutoReplyOutput) ProtoMessage() {}

func (x *DescribeAutoReplyOutput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_helpdesk_auto_reply_service_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DescribeAutoReplyOutput.ProtoReflect.Descriptor instead.
func (*DescribeAutoReplyOutput) Descriptor() ([]byte, []int) {
	return file_eolymp_helpdesk_auto_reply_service_proto_rawDescGZIP(), []int{9}
}

func (x *DescribeAutoReplyOutput) GetReply() *AutoReply {
	if x != nil {
		return x.Reply
	}
	return nil
}

type ListAutoRepliesInput_Filter struct {
	state         protoimpl.MessageState      `protogen:"open.v1"`
	Query         string                      `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
	Id            []*wellknown.ExpressionID   `protobuf:"bytes,2,rep,name=id,proto3" json:"id,omitempty"`
	Locale        []*wellknown.ExpressionEnum `protobuf:"bytes,3,rep,name=locale,proto3" json:"locale,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListAutoRepliesInput_Filter) Reset() {
	*x = ListAutoRepliesInput_Filter{}
	mi := &file_eolymp_helpdesk_auto_reply_service_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListAutoRepliesInput_Filter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListAutoRepliesInput_Filter) ProtoMessage() {}

func (x *ListAutoRepliesInput_Filter) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_helpdesk_auto_reply_service_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListAutoRepliesInput_Filter.ProtoReflect.Descriptor instead.
func (*ListAutoRepliesInput_Filter) Descriptor() ([]byte, []int) {
	return file_eolymp_helpdesk_auto_reply_service_proto_rawDescGZIP(), []int{6, 0}
}

func (x *ListAutoRepliesInput_Filter) GetQuery() string {
	if x != nil {
		return x.Query
	}
	return ""
}

func (x *ListAutoRepliesInput_Filter) GetId() []*wellknown.ExpressionID {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *ListAutoRepliesInput_Filter) GetLocale() []*wellknown.ExpressionEnum {
	if x != nil {
		return x.Locale
	}
	return nil
}

var File_eolymp_helpdesk_auto_reply_service_proto protoreflect.FileDescriptor

const file_eolymp_helpdesk_auto_reply_service_proto_rawDesc = "" +
	"\n" +
	"(eolymp/helpdesk/auto_reply_service.proto\x12\x0feolymp.helpdesk\x1a\x1deolymp/annotations/http.proto\x1a\"eolymp/annotations/ratelimit.proto\x1a\x1eeolymp/annotations/scope.proto\x1a eolymp/helpdesk/auto_reply.proto\x1a!eolymp/wellknown/expression.proto\"H\n" +
	"\x14CreateAutoReplyInput\x120\n" +
	"\x05reply\x18\x01 \x01(\v2\x1a.eolymp.helpdesk.AutoReplyR\x05reply\"2\n" +
	"\x15CreateAutoReplyOutput\x12\x19\n" +
	"\breply_id\x18\x01 \x01(\tR\areplyId\"c\n" +
	"\x14UpdateAutoReplyInput\x12\x19\n" +
	"\breply_id\x18\x01 \x01(\tR\areplyId\x120\n" +
	"\x05reply\x18\x02 \x01(\v2\x1a.eolymp.helpdesk.AutoReplyR\x05reply\"\x17\n" +
	"\x15UpdateAutoReplyOutput\"1\n" +
	"\x14DeleteAutoReplyInput\x12\x19\n" +
	"\breply_id\x18\x01 \x01(\tR\areplyId\"\x17\n" +
	"\x15DeleteAutoReplyOutput\"\xad\x02\n" +
	"\x14ListAutoRepliesInput\x12\x16\n" +
	"\x06render\x18\x01 \x01(\bR\x06render\x12\x16\n" +
	"\x06offset\x18\n" +
	" \x01(\x05R\x06offset\x12\x12\n" +
	"\x04size\x18\v \x01(\x05R\x04size\x12F\n" +
	"\afilters\x18( \x01(\v2,.eolymp.helpdesk.ListAutoRepliesInput.FilterR\afilters\x1a\x88\x01\n" +
	"\x06Filter\x12\x14\n" +
	"\x05query\x18\x01 \x01(\tR\x05query\x12.\n" +
	"\x02id\x18\x02 \x03(\v2\x1e.eolymp.wellknown.ExpressionIDR\x02id\x128\n" +
	"\x06locale\x18\x03 \x03(\v2 .eolymp.wellknown.ExpressionEnumR\x06locale\"_\n" +
	"\x15ListAutoRepliesOutput\x12\x14\n" +
	"\x05total\x18\x01 \x01(\x05R\x05total\x120\n" +
	"\x05items\x18\x02 \x03(\v2\x1a.eolymp.helpdesk.AutoReplyR\x05items\"K\n" +
	"\x16DescribeAutoReplyInput\x12\x19\n" +
	"\breply_id\x18\x01 \x01(\tR\areplyId\x12\x16\n" +
	"\x06render\x18\x02 \x01(\bR\x06render\"K\n" +
	"\x17DescribeAutoReplyOutput\x120\n" +
	"\x05reply\x18\x01 \x01(\v2\x1a.eolymp.helpdesk.AutoReplyR\x05reply2\xb7\a\n" +
	"\x10AutoReplyService\x12\xb0\x01\n" +
	"\x0fCreateAutoReply\x12%.eolymp.helpdesk.CreateAutoReplyInput\x1a&.eolymp.helpdesk.CreateAutoReplyOutput\"N\xea\xe2\n" +
	"\v\xf5\xe2\n" +
	"\x00\x00\x00@\xf8\xe2\n" +
	"\n" +
	"\x82\xe3\n" +
	"\x1d\x8a\xe3\n" +
	"\x19helpdesk:auto-reply:write\x82\xd3\xe4\x93\x02\x18\"\x16/helpdesk/auto-replies\x12\xbb\x01\n" +
	"\x0fUpdateAutoReply\x12%.eolymp.helpdesk.UpdateAutoReplyInput\x1a&.eolymp.helpdesk.UpdateAutoReplyOutput\"Y\xea\xe2\n" +
	"\v\xf5\xe2\n" +
	"\x00\x00\x00@\xf8\xe2\n" +
	"\n" +
	"\x82\xe3\n" +
	"\x1d\x8a\xe3\n" +
	"\x19helpdesk:auto-reply:write\x82\xd3\xe4\x93\x02#\x1a!/helpdesk/auto-replies/{reply_id}\x12\xbb\x01\n" +
	"\x0fDeleteAutoReply\x12%.eolymp.helpdesk.DeleteAutoReplyInput\x1a&.eolymp.helpdesk.DeleteAutoReplyOutput\"Y\xea\xe2\n" +
	"\v\xf5\xe2\n" +
	"\x00\x00\x00@\xf8\xe2\n" +
	"\n" +
	"\x82\xe3\n" +
	"\x1d\x8a\xe3\n" +
	"\x19helpdesk:auto-reply:write\x82\xd3\xe4\x93\x02#*!/helpdesk/auto-replies/{reply_id}\x12\xaf\x01\n" +
	"\x0fListAutoReplies\x12%.eolymp.helpdesk.ListAutoRepliesInput\x1a&.eolymp.helpdesk.ListAutoRepliesOutput\"M\xea\xe2\n" +
	"\v\xf5\xe2\n" +
	"\x00\x00\xa0@\xf8\xe2\n" +
	"\x14\x82\xe3\n" +
	"\x1c\x8a\xe3\n" +
	"\x18helpdesk:auto-reply:read\x82\xd3\xe4\x93\x02\x18\x12\x16/helpdesk/auto-replies\x12\xc1\x01\n" +
	"\x11DescribeAutoReply\x12'.eolymp.helpdesk.DescribeAutoReplyInput\x1a(.eolymp.helpdesk.DescribeAutoReplyOutput\"Y\xea\xe2\n" +
	"\v\xf5\xe2\n" +
	"\x00\x00\xa0@\xf8\xe2\n" +
	"2\x82\xe3\n" +
	"\x1d\x8a\xe3\n" +
	"\x19helpdesk:auto-reply:write\x82\xd3\xe4\x93\x02#\x12!/helpdesk/auto-replies/{reply_id}B3Z1github.com/eolymp/go-sdk/eolymp/helpdesk;helpdeskb\x06proto3"

var (
	file_eolymp_helpdesk_auto_reply_service_proto_rawDescOnce sync.Once
	file_eolymp_helpdesk_auto_reply_service_proto_rawDescData []byte
)

func file_eolymp_helpdesk_auto_reply_service_proto_rawDescGZIP() []byte {
	file_eolymp_helpdesk_auto_reply_service_proto_rawDescOnce.Do(func() {
		file_eolymp_helpdesk_auto_reply_service_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_eolymp_helpdesk_auto_reply_service_proto_rawDesc), len(file_eolymp_helpdesk_auto_reply_service_proto_rawDesc)))
	})
	return file_eolymp_helpdesk_auto_reply_service_proto_rawDescData
}

var file_eolymp_helpdesk_auto_reply_service_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_eolymp_helpdesk_auto_reply_service_proto_goTypes = []any{
	(*CreateAutoReplyInput)(nil),        // 0: eolymp.helpdesk.CreateAutoReplyInput
	(*CreateAutoReplyOutput)(nil),       // 1: eolymp.helpdesk.CreateAutoReplyOutput
	(*UpdateAutoReplyInput)(nil),        // 2: eolymp.helpdesk.UpdateAutoReplyInput
	(*UpdateAutoReplyOutput)(nil),       // 3: eolymp.helpdesk.UpdateAutoReplyOutput
	(*DeleteAutoReplyInput)(nil),        // 4: eolymp.helpdesk.DeleteAutoReplyInput
	(*DeleteAutoReplyOutput)(nil),       // 5: eolymp.helpdesk.DeleteAutoReplyOutput
	(*ListAutoRepliesInput)(nil),        // 6: eolymp.helpdesk.ListAutoRepliesInput
	(*ListAutoRepliesOutput)(nil),       // 7: eolymp.helpdesk.ListAutoRepliesOutput
	(*DescribeAutoReplyInput)(nil),      // 8: eolymp.helpdesk.DescribeAutoReplyInput
	(*DescribeAutoReplyOutput)(nil),     // 9: eolymp.helpdesk.DescribeAutoReplyOutput
	(*ListAutoRepliesInput_Filter)(nil), // 10: eolymp.helpdesk.ListAutoRepliesInput.Filter
	(*AutoReply)(nil),                   // 11: eolymp.helpdesk.AutoReply
	(*wellknown.ExpressionID)(nil),      // 12: eolymp.wellknown.ExpressionID
	(*wellknown.ExpressionEnum)(nil),    // 13: eolymp.wellknown.ExpressionEnum
}
var file_eolymp_helpdesk_auto_reply_service_proto_depIdxs = []int32{
	11, // 0: eolymp.helpdesk.CreateAutoReplyInput.reply:type_name -> eolymp.helpdesk.AutoReply
	11, // 1: eolymp.helpdesk.UpdateAutoReplyInput.reply:type_name -> eolymp.helpdesk.AutoReply
	10, // 2: eolymp.helpdesk.ListAutoRepliesInput.filters:type_name -> eolymp.helpdesk.ListAutoRepliesInput.Filter
	11, // 3: eolymp.helpdesk.ListAutoRepliesOutput.items:type_name -> eolymp.helpdesk.AutoReply
	11, // 4: eolymp.helpdesk.DescribeAutoReplyOutput.reply:type_name -> eolymp.helpdesk.AutoReply
	12, // 5: eolymp.helpdesk.ListAutoRepliesInput.Filter.id:type_name -> eolymp.wellknown.ExpressionID
	13, // 6: eolymp.helpdesk.ListAutoRepliesInput.Filter.locale:type_name -> eolymp.wellknown.ExpressionEnum
	0,  // 7: eolymp.helpdesk.AutoReplyService.CreateAutoReply:input_type -> eolymp.helpdesk.CreateAutoReplyInput
	2,  // 8: eolymp.helpdesk.AutoReplyService.UpdateAutoReply:input_type -> eolymp.helpdesk.UpdateAutoReplyInput
	4,  // 9: eolymp.helpdesk.AutoReplyService.DeleteAutoReply:input_type -> eolymp.helpdesk.DeleteAutoReplyInput
	6,  // 10: eolymp.helpdesk.AutoReplyService.ListAutoReplies:input_type -> eolymp.helpdesk.ListAutoRepliesInput
	8,  // 11: eolymp.helpdesk.AutoReplyService.DescribeAutoReply:input_type -> eolymp.helpdesk.DescribeAutoReplyInput
	1,  // 12: eolymp.helpdesk.AutoReplyService.CreateAutoReply:output_type -> eolymp.helpdesk.CreateAutoReplyOutput
	3,  // 13: eolymp.helpdesk.AutoReplyService.UpdateAutoReply:output_type -> eolymp.helpdesk.UpdateAutoReplyOutput
	5,  // 14: eolymp.helpdesk.AutoReplyService.DeleteAutoReply:output_type -> eolymp.helpdesk.DeleteAutoReplyOutput
	7,  // 15: eolymp.helpdesk.AutoReplyService.ListAutoReplies:output_type -> eolymp.helpdesk.ListAutoRepliesOutput
	9,  // 16: eolymp.helpdesk.AutoReplyService.DescribeAutoReply:output_type -> eolymp.helpdesk.DescribeAutoReplyOutput
	12, // [12:17] is the sub-list for method output_type
	7,  // [7:12] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_eolymp_helpdesk_auto_reply_service_proto_init() }
func file_eolymp_helpdesk_auto_reply_service_proto_init() {
	if File_eolymp_helpdesk_auto_reply_service_proto != nil {
		return
	}
	file_eolymp_helpdesk_auto_reply_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_eolymp_helpdesk_auto_reply_service_proto_rawDesc), len(file_eolymp_helpdesk_auto_reply_service_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_eolymp_helpdesk_auto_reply_service_proto_goTypes,
		DependencyIndexes: file_eolymp_helpdesk_auto_reply_service_proto_depIdxs,
		MessageInfos:      file_eolymp_helpdesk_auto_reply_service_proto_msgTypes,
	}.Build()
	File_eolymp_helpdesk_auto_reply_service_proto = out.File
	file_eolymp_helpdesk_auto_reply_service_proto_goTypes = nil
	file_eolymp_helpdesk_auto_reply_service_proto_depIdxs = nil
}
