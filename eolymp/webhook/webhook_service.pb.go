// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.29.4
// source: eolymp/webhook/webhook_service.proto

package webhook

import (
	_ "github.com/eolymp/go-sdk/eolymp/annotations"
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

type CreateWebhookInput struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Webhook       *Webhook               `protobuf:"bytes,1,opt,name=webhook,proto3" json:"webhook,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateWebhookInput) Reset() {
	*x = CreateWebhookInput{}
	mi := &file_eolymp_webhook_webhook_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateWebhookInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateWebhookInput) ProtoMessage() {}

func (x *CreateWebhookInput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_webhook_webhook_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateWebhookInput.ProtoReflect.Descriptor instead.
func (*CreateWebhookInput) Descriptor() ([]byte, []int) {
	return file_eolymp_webhook_webhook_service_proto_rawDescGZIP(), []int{0}
}

func (x *CreateWebhookInput) GetWebhook() *Webhook {
	if x != nil {
		return x.Webhook
	}
	return nil
}

type CreateWebhookOutput struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	WebhookId     string                 `protobuf:"bytes,1,opt,name=webhook_id,json=webhookId,proto3" json:"webhook_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateWebhookOutput) Reset() {
	*x = CreateWebhookOutput{}
	mi := &file_eolymp_webhook_webhook_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateWebhookOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateWebhookOutput) ProtoMessage() {}

func (x *CreateWebhookOutput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_webhook_webhook_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateWebhookOutput.ProtoReflect.Descriptor instead.
func (*CreateWebhookOutput) Descriptor() ([]byte, []int) {
	return file_eolymp_webhook_webhook_service_proto_rawDescGZIP(), []int{1}
}

func (x *CreateWebhookOutput) GetWebhookId() string {
	if x != nil {
		return x.WebhookId
	}
	return ""
}

type UpdateWebhookInput struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Patch         []Webhook_Patch        `protobuf:"varint,1,rep,packed,name=patch,proto3,enum=eolymp.webhook.Webhook_Patch" json:"patch,omitempty"`
	WebhookId     string                 `protobuf:"bytes,2,opt,name=webhook_id,json=webhookId,proto3" json:"webhook_id,omitempty"`
	Webhook       *Webhook               `protobuf:"bytes,3,opt,name=webhook,proto3" json:"webhook,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateWebhookInput) Reset() {
	*x = UpdateWebhookInput{}
	mi := &file_eolymp_webhook_webhook_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateWebhookInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateWebhookInput) ProtoMessage() {}

func (x *UpdateWebhookInput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_webhook_webhook_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateWebhookInput.ProtoReflect.Descriptor instead.
func (*UpdateWebhookInput) Descriptor() ([]byte, []int) {
	return file_eolymp_webhook_webhook_service_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateWebhookInput) GetPatch() []Webhook_Patch {
	if x != nil {
		return x.Patch
	}
	return nil
}

func (x *UpdateWebhookInput) GetWebhookId() string {
	if x != nil {
		return x.WebhookId
	}
	return ""
}

func (x *UpdateWebhookInput) GetWebhook() *Webhook {
	if x != nil {
		return x.Webhook
	}
	return nil
}

type UpdateWebhookOutput struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateWebhookOutput) Reset() {
	*x = UpdateWebhookOutput{}
	mi := &file_eolymp_webhook_webhook_service_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateWebhookOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateWebhookOutput) ProtoMessage() {}

func (x *UpdateWebhookOutput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_webhook_webhook_service_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateWebhookOutput.ProtoReflect.Descriptor instead.
func (*UpdateWebhookOutput) Descriptor() ([]byte, []int) {
	return file_eolymp_webhook_webhook_service_proto_rawDescGZIP(), []int{3}
}

type DeleteWebhookInput struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	WebhookId     string                 `protobuf:"bytes,1,opt,name=webhook_id,json=webhookId,proto3" json:"webhook_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteWebhookInput) Reset() {
	*x = DeleteWebhookInput{}
	mi := &file_eolymp_webhook_webhook_service_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteWebhookInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteWebhookInput) ProtoMessage() {}

func (x *DeleteWebhookInput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_webhook_webhook_service_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteWebhookInput.ProtoReflect.Descriptor instead.
func (*DeleteWebhookInput) Descriptor() ([]byte, []int) {
	return file_eolymp_webhook_webhook_service_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteWebhookInput) GetWebhookId() string {
	if x != nil {
		return x.WebhookId
	}
	return ""
}

type DeleteWebhookOutput struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteWebhookOutput) Reset() {
	*x = DeleteWebhookOutput{}
	mi := &file_eolymp_webhook_webhook_service_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteWebhookOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteWebhookOutput) ProtoMessage() {}

func (x *DeleteWebhookOutput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_webhook_webhook_service_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteWebhookOutput.ProtoReflect.Descriptor instead.
func (*DeleteWebhookOutput) Descriptor() ([]byte, []int) {
	return file_eolymp_webhook_webhook_service_proto_rawDescGZIP(), []int{5}
}

type DescribeWebhookInput struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	WebhookId     string                 `protobuf:"bytes,1,opt,name=webhook_id,json=webhookId,proto3" json:"webhook_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DescribeWebhookInput) Reset() {
	*x = DescribeWebhookInput{}
	mi := &file_eolymp_webhook_webhook_service_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DescribeWebhookInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescribeWebhookInput) ProtoMessage() {}

func (x *DescribeWebhookInput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_webhook_webhook_service_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DescribeWebhookInput.ProtoReflect.Descriptor instead.
func (*DescribeWebhookInput) Descriptor() ([]byte, []int) {
	return file_eolymp_webhook_webhook_service_proto_rawDescGZIP(), []int{6}
}

func (x *DescribeWebhookInput) GetWebhookId() string {
	if x != nil {
		return x.WebhookId
	}
	return ""
}

type DescribeWebhookOutput struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Webhook       *Webhook               `protobuf:"bytes,1,opt,name=webhook,proto3" json:"webhook,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DescribeWebhookOutput) Reset() {
	*x = DescribeWebhookOutput{}
	mi := &file_eolymp_webhook_webhook_service_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DescribeWebhookOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescribeWebhookOutput) ProtoMessage() {}

func (x *DescribeWebhookOutput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_webhook_webhook_service_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DescribeWebhookOutput.ProtoReflect.Descriptor instead.
func (*DescribeWebhookOutput) Descriptor() ([]byte, []int) {
	return file_eolymp_webhook_webhook_service_proto_rawDescGZIP(), []int{7}
}

func (x *DescribeWebhookOutput) GetWebhook() *Webhook {
	if x != nil {
		return x.Webhook
	}
	return nil
}

type ListWebhooksInput struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Size          int32                  `protobuf:"varint,11,opt,name=size,proto3" json:"size,omitempty"`
	Offset        int32                  `protobuf:"varint,10,opt,name=offset,proto3" json:"offset,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListWebhooksInput) Reset() {
	*x = ListWebhooksInput{}
	mi := &file_eolymp_webhook_webhook_service_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListWebhooksInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListWebhooksInput) ProtoMessage() {}

func (x *ListWebhooksInput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_webhook_webhook_service_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListWebhooksInput.ProtoReflect.Descriptor instead.
func (*ListWebhooksInput) Descriptor() ([]byte, []int) {
	return file_eolymp_webhook_webhook_service_proto_rawDescGZIP(), []int{8}
}

func (x *ListWebhooksInput) GetSize() int32 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *ListWebhooksInput) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

type ListWebhooksOutput struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Total         int32                  `protobuf:"varint,1,opt,name=total,proto3" json:"total,omitempty"`
	Items         []*Webhook             `protobuf:"bytes,2,rep,name=items,proto3" json:"items,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListWebhooksOutput) Reset() {
	*x = ListWebhooksOutput{}
	mi := &file_eolymp_webhook_webhook_service_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListWebhooksOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListWebhooksOutput) ProtoMessage() {}

func (x *ListWebhooksOutput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_webhook_webhook_service_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListWebhooksOutput.ProtoReflect.Descriptor instead.
func (*ListWebhooksOutput) Descriptor() ([]byte, []int) {
	return file_eolymp_webhook_webhook_service_proto_rawDescGZIP(), []int{9}
}

func (x *ListWebhooksOutput) GetTotal() int32 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *ListWebhooksOutput) GetItems() []*Webhook {
	if x != nil {
		return x.Items
	}
	return nil
}

type TestWebhookInput struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	WebhookId     string                 `protobuf:"bytes,1,opt,name=webhook_id,json=webhookId,proto3" json:"webhook_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *TestWebhookInput) Reset() {
	*x = TestWebhookInput{}
	mi := &file_eolymp_webhook_webhook_service_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TestWebhookInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestWebhookInput) ProtoMessage() {}

func (x *TestWebhookInput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_webhook_webhook_service_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestWebhookInput.ProtoReflect.Descriptor instead.
func (*TestWebhookInput) Descriptor() ([]byte, []int) {
	return file_eolymp_webhook_webhook_service_proto_rawDescGZIP(), []int{10}
}

func (x *TestWebhookInput) GetWebhookId() string {
	if x != nil {
		return x.WebhookId
	}
	return ""
}

type TestWebhookOutput struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Status        int32                  `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Response      string                 `protobuf:"bytes,2,opt,name=response,proto3" json:"response,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *TestWebhookOutput) Reset() {
	*x = TestWebhookOutput{}
	mi := &file_eolymp_webhook_webhook_service_proto_msgTypes[11]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TestWebhookOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestWebhookOutput) ProtoMessage() {}

func (x *TestWebhookOutput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_webhook_webhook_service_proto_msgTypes[11]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestWebhookOutput.ProtoReflect.Descriptor instead.
func (*TestWebhookOutput) Descriptor() ([]byte, []int) {
	return file_eolymp_webhook_webhook_service_proto_rawDescGZIP(), []int{11}
}

func (x *TestWebhookOutput) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *TestWebhookOutput) GetResponse() string {
	if x != nil {
		return x.Response
	}
	return ""
}

var File_eolymp_webhook_webhook_service_proto protoreflect.FileDescriptor

const file_eolymp_webhook_webhook_service_proto_rawDesc = "" +
	"\n" +
	"$eolymp/webhook/webhook_service.proto\x12\x0eeolymp.webhook\x1a\x1deolymp/annotations/http.proto\x1a\"eolymp/annotations/ratelimit.proto\x1a\x1eeolymp/annotations/scope.proto\x1a\x1ceolymp/webhook/webhook.proto\"G\n" +
	"\x12CreateWebhookInput\x121\n" +
	"\awebhook\x18\x01 \x01(\v2\x17.eolymp.webhook.WebhookR\awebhook\"4\n" +
	"\x13CreateWebhookOutput\x12\x1d\n" +
	"\n" +
	"webhook_id\x18\x01 \x01(\tR\twebhookId\"\x9b\x01\n" +
	"\x12UpdateWebhookInput\x123\n" +
	"\x05patch\x18\x01 \x03(\x0e2\x1d.eolymp.webhook.Webhook.PatchR\x05patch\x12\x1d\n" +
	"\n" +
	"webhook_id\x18\x02 \x01(\tR\twebhookId\x121\n" +
	"\awebhook\x18\x03 \x01(\v2\x17.eolymp.webhook.WebhookR\awebhook\"\x15\n" +
	"\x13UpdateWebhookOutput\"3\n" +
	"\x12DeleteWebhookInput\x12\x1d\n" +
	"\n" +
	"webhook_id\x18\x01 \x01(\tR\twebhookId\"\x15\n" +
	"\x13DeleteWebhookOutput\"5\n" +
	"\x14DescribeWebhookInput\x12\x1d\n" +
	"\n" +
	"webhook_id\x18\x01 \x01(\tR\twebhookId\"J\n" +
	"\x15DescribeWebhookOutput\x121\n" +
	"\awebhook\x18\x01 \x01(\v2\x17.eolymp.webhook.WebhookR\awebhook\"?\n" +
	"\x11ListWebhooksInput\x12\x12\n" +
	"\x04size\x18\v \x01(\x05R\x04size\x12\x16\n" +
	"\x06offset\x18\n" +
	" \x01(\x05R\x06offset\"Y\n" +
	"\x12ListWebhooksOutput\x12\x14\n" +
	"\x05total\x18\x01 \x01(\x05R\x05total\x12-\n" +
	"\x05items\x18\x02 \x03(\v2\x17.eolymp.webhook.WebhookR\x05items\"1\n" +
	"\x10TestWebhookInput\x12\x1d\n" +
	"\n" +
	"webhook_id\x18\x01 \x01(\tR\twebhookId\"G\n" +
	"\x11TestWebhookOutput\x12\x16\n" +
	"\x06status\x18\x01 \x01(\x05R\x06status\x12\x1a\n" +
	"\bresponse\x18\x02 \x01(\tR\bresponse2\xe0\a\n" +
	"\x0eWebhookService\x12\x97\x01\n" +
	"\rCreateWebhook\x12\".eolymp.webhook.CreateWebhookInput\x1a#.eolymp.webhook.CreateWebhookOutput\"=\xea\xe2\n" +
	"\v\xf5\xe2\n" +
	"\x00\x00\xa0@\xf8\xe2\n" +
	"\x14\x82\xe3\n" +
	"\x19\x8a\xe3\n" +
	"\x15webhook:webhook:write\x82\xd3\xe4\x93\x02\v\"\t/webhooks\x12\xa4\x01\n" +
	"\rUpdateWebhook\x12\".eolymp.webhook.UpdateWebhookInput\x1a#.eolymp.webhook.UpdateWebhookOutput\"J\xea\xe2\n" +
	"\v\xf5\xe2\n" +
	"\x00\x00\xa0@\xf8\xe2\n" +
	"\x14\x82\xe3\n" +
	"\x19\x8a\xe3\n" +
	"\x15webhook:webhook:write\x82\xd3\xe4\x93\x02\x18\"\x16/webhooks/{webhook_id}\x12\xa4\x01\n" +
	"\rDeleteWebhook\x12\".eolymp.webhook.DeleteWebhookInput\x1a#.eolymp.webhook.DeleteWebhookOutput\"J\xea\xe2\n" +
	"\v\xf5\xe2\n" +
	"\x00\x00\xa0@\xf8\xe2\n" +
	"\x14\x82\xe3\n" +
	"\x19\x8a\xe3\n" +
	"\x15webhook:webhook:write\x82\xd3\xe4\x93\x02\x18*\x16/webhooks/{webhook_id}\x12\xa9\x01\n" +
	"\x0fDescribeWebhook\x12$.eolymp.webhook.DescribeWebhookInput\x1a%.eolymp.webhook.DescribeWebhookOutput\"I\xea\xe2\n" +
	"\v\xf5\xe2\n" +
	"\x00\x00\xa0@\xf8\xe2\n" +
	"\x14\x82\xe3\n" +
	"\x18\x8a\xe3\n" +
	"\x14webhook:webhook:read\x82\xd3\xe4\x93\x02\x18\x12\x16/webhooks/{webhook_id}\x12\x93\x01\n" +
	"\fListWebhooks\x12!.eolymp.webhook.ListWebhooksInput\x1a\".eolymp.webhook.ListWebhooksOutput\"<\xea\xe2\n" +
	"\v\xf5\xe2\n" +
	"\x00\x00\xa0@\xf8\xe2\n" +
	"\x14\x82\xe3\n" +
	"\x18\x8a\xe3\n" +
	"\x14webhook:webhook:read\x82\xd3\xe4\x93\x02\v\x12\t/webhooks\x12\xa3\x01\n" +
	"\vTestWebhook\x12 .eolymp.webhook.TestWebhookInput\x1a!.eolymp.webhook.TestWebhookOutput\"O\xea\xe2\n" +
	"\v\xf5\xe2\n" +
	"\x00\x00\xa0@\xf8\xe2\n" +
	"\x14\x82\xe3\n" +
	"\x19\x8a\xe3\n" +
	"\x15webhook:webhook:write\x82\xd3\xe4\x93\x02\x1d\"\x1b/webhooks/{webhook_id}/testB1Z/github.com/eolymp/go-sdk/eolymp/webhook;webhookb\x06proto3"

var (
	file_eolymp_webhook_webhook_service_proto_rawDescOnce sync.Once
	file_eolymp_webhook_webhook_service_proto_rawDescData []byte
)

func file_eolymp_webhook_webhook_service_proto_rawDescGZIP() []byte {
	file_eolymp_webhook_webhook_service_proto_rawDescOnce.Do(func() {
		file_eolymp_webhook_webhook_service_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_eolymp_webhook_webhook_service_proto_rawDesc), len(file_eolymp_webhook_webhook_service_proto_rawDesc)))
	})
	return file_eolymp_webhook_webhook_service_proto_rawDescData
}

var file_eolymp_webhook_webhook_service_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_eolymp_webhook_webhook_service_proto_goTypes = []any{
	(*CreateWebhookInput)(nil),    // 0: eolymp.webhook.CreateWebhookInput
	(*CreateWebhookOutput)(nil),   // 1: eolymp.webhook.CreateWebhookOutput
	(*UpdateWebhookInput)(nil),    // 2: eolymp.webhook.UpdateWebhookInput
	(*UpdateWebhookOutput)(nil),   // 3: eolymp.webhook.UpdateWebhookOutput
	(*DeleteWebhookInput)(nil),    // 4: eolymp.webhook.DeleteWebhookInput
	(*DeleteWebhookOutput)(nil),   // 5: eolymp.webhook.DeleteWebhookOutput
	(*DescribeWebhookInput)(nil),  // 6: eolymp.webhook.DescribeWebhookInput
	(*DescribeWebhookOutput)(nil), // 7: eolymp.webhook.DescribeWebhookOutput
	(*ListWebhooksInput)(nil),     // 8: eolymp.webhook.ListWebhooksInput
	(*ListWebhooksOutput)(nil),    // 9: eolymp.webhook.ListWebhooksOutput
	(*TestWebhookInput)(nil),      // 10: eolymp.webhook.TestWebhookInput
	(*TestWebhookOutput)(nil),     // 11: eolymp.webhook.TestWebhookOutput
	(*Webhook)(nil),               // 12: eolymp.webhook.Webhook
	(Webhook_Patch)(0),            // 13: eolymp.webhook.Webhook.Patch
}
var file_eolymp_webhook_webhook_service_proto_depIdxs = []int32{
	12, // 0: eolymp.webhook.CreateWebhookInput.webhook:type_name -> eolymp.webhook.Webhook
	13, // 1: eolymp.webhook.UpdateWebhookInput.patch:type_name -> eolymp.webhook.Webhook.Patch
	12, // 2: eolymp.webhook.UpdateWebhookInput.webhook:type_name -> eolymp.webhook.Webhook
	12, // 3: eolymp.webhook.DescribeWebhookOutput.webhook:type_name -> eolymp.webhook.Webhook
	12, // 4: eolymp.webhook.ListWebhooksOutput.items:type_name -> eolymp.webhook.Webhook
	0,  // 5: eolymp.webhook.WebhookService.CreateWebhook:input_type -> eolymp.webhook.CreateWebhookInput
	2,  // 6: eolymp.webhook.WebhookService.UpdateWebhook:input_type -> eolymp.webhook.UpdateWebhookInput
	4,  // 7: eolymp.webhook.WebhookService.DeleteWebhook:input_type -> eolymp.webhook.DeleteWebhookInput
	6,  // 8: eolymp.webhook.WebhookService.DescribeWebhook:input_type -> eolymp.webhook.DescribeWebhookInput
	8,  // 9: eolymp.webhook.WebhookService.ListWebhooks:input_type -> eolymp.webhook.ListWebhooksInput
	10, // 10: eolymp.webhook.WebhookService.TestWebhook:input_type -> eolymp.webhook.TestWebhookInput
	1,  // 11: eolymp.webhook.WebhookService.CreateWebhook:output_type -> eolymp.webhook.CreateWebhookOutput
	3,  // 12: eolymp.webhook.WebhookService.UpdateWebhook:output_type -> eolymp.webhook.UpdateWebhookOutput
	5,  // 13: eolymp.webhook.WebhookService.DeleteWebhook:output_type -> eolymp.webhook.DeleteWebhookOutput
	7,  // 14: eolymp.webhook.WebhookService.DescribeWebhook:output_type -> eolymp.webhook.DescribeWebhookOutput
	9,  // 15: eolymp.webhook.WebhookService.ListWebhooks:output_type -> eolymp.webhook.ListWebhooksOutput
	11, // 16: eolymp.webhook.WebhookService.TestWebhook:output_type -> eolymp.webhook.TestWebhookOutput
	11, // [11:17] is the sub-list for method output_type
	5,  // [5:11] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_eolymp_webhook_webhook_service_proto_init() }
func file_eolymp_webhook_webhook_service_proto_init() {
	if File_eolymp_webhook_webhook_service_proto != nil {
		return
	}
	file_eolymp_webhook_webhook_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_eolymp_webhook_webhook_service_proto_rawDesc), len(file_eolymp_webhook_webhook_service_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_eolymp_webhook_webhook_service_proto_goTypes,
		DependencyIndexes: file_eolymp_webhook_webhook_service_proto_depIdxs,
		MessageInfos:      file_eolymp_webhook_webhook_service_proto_msgTypes,
	}.Build()
	File_eolymp_webhook_webhook_service_proto = out.File
	file_eolymp_webhook_webhook_service_proto_goTypes = nil
	file_eolymp_webhook_webhook_service_proto_depIdxs = nil
}
