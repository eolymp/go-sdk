// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.29.4
// source: eolymp/discussion/subscription_service.proto

package discussion

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

type DescribeSubscriptionInput struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DescribeSubscriptionInput) Reset() {
	*x = DescribeSubscriptionInput{}
	mi := &file_eolymp_discussion_subscription_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DescribeSubscriptionInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescribeSubscriptionInput) ProtoMessage() {}

func (x *DescribeSubscriptionInput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_discussion_subscription_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DescribeSubscriptionInput.ProtoReflect.Descriptor instead.
func (*DescribeSubscriptionInput) Descriptor() ([]byte, []int) {
	return file_eolymp_discussion_subscription_service_proto_rawDescGZIP(), []int{0}
}

type DescribeSubscriptionOutput struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Subscription  *Subscription          `protobuf:"bytes,1,opt,name=subscription,proto3" json:"subscription,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DescribeSubscriptionOutput) Reset() {
	*x = DescribeSubscriptionOutput{}
	mi := &file_eolymp_discussion_subscription_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DescribeSubscriptionOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescribeSubscriptionOutput) ProtoMessage() {}

func (x *DescribeSubscriptionOutput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_discussion_subscription_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DescribeSubscriptionOutput.ProtoReflect.Descriptor instead.
func (*DescribeSubscriptionOutput) Descriptor() ([]byte, []int) {
	return file_eolymp_discussion_subscription_service_proto_rawDescGZIP(), []int{1}
}

func (x *DescribeSubscriptionOutput) GetSubscription() *Subscription {
	if x != nil {
		return x.Subscription
	}
	return nil
}

type UpdateSubscriptionInput struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Subscription  *Subscription          `protobuf:"bytes,1,opt,name=subscription,proto3" json:"subscription,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateSubscriptionInput) Reset() {
	*x = UpdateSubscriptionInput{}
	mi := &file_eolymp_discussion_subscription_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateSubscriptionInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateSubscriptionInput) ProtoMessage() {}

func (x *UpdateSubscriptionInput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_discussion_subscription_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateSubscriptionInput.ProtoReflect.Descriptor instead.
func (*UpdateSubscriptionInput) Descriptor() ([]byte, []int) {
	return file_eolymp_discussion_subscription_service_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateSubscriptionInput) GetSubscription() *Subscription {
	if x != nil {
		return x.Subscription
	}
	return nil
}

type UpdateSubscriptionOutput struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateSubscriptionOutput) Reset() {
	*x = UpdateSubscriptionOutput{}
	mi := &file_eolymp_discussion_subscription_service_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateSubscriptionOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateSubscriptionOutput) ProtoMessage() {}

func (x *UpdateSubscriptionOutput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_discussion_subscription_service_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateSubscriptionOutput.ProtoReflect.Descriptor instead.
func (*UpdateSubscriptionOutput) Descriptor() ([]byte, []int) {
	return file_eolymp_discussion_subscription_service_proto_rawDescGZIP(), []int{3}
}

var File_eolymp_discussion_subscription_service_proto protoreflect.FileDescriptor

const file_eolymp_discussion_subscription_service_proto_rawDesc = "" +
	"\n" +
	",eolymp/discussion/subscription_service.proto\x12\x11eolymp.discussion\x1a\x1deolymp/annotations/http.proto\x1a\"eolymp/annotations/ratelimit.proto\x1a\x1eeolymp/annotations/scope.proto\x1a$eolymp/discussion/subscription.proto\"\x1b\n" +
	"\x19DescribeSubscriptionInput\"a\n" +
	"\x1aDescribeSubscriptionOutput\x12C\n" +
	"\fsubscription\x18\x01 \x01(\v2\x1f.eolymp.discussion.SubscriptionR\fsubscription\"^\n" +
	"\x17UpdateSubscriptionInput\x12C\n" +
	"\fsubscription\x18\x01 \x01(\v2\x1f.eolymp.discussion.SubscriptionR\fsubscription\"\x1a\n" +
	"\x18UpdateSubscriptionOutput2\x91\x03\n" +
	"\x13SubscriptionService\x12\xbe\x01\n" +
	"\x14DescribeSubscription\x12,.eolymp.discussion.DescribeSubscriptionInput\x1a-.eolymp.discussion.DescribeSubscriptionOutput\"I\xea\xe2\n" +
	"\f\xf5\xe2\n" +
	"\x00\x00\xa0A\xf8\xe2\n" +
	"\xf4\x03\x82\xe3\n" +
	" \x8a\xe3\n" +
	"\x1cdiscussion:subscription:read\x82\xd3\xe4\x93\x02\x0f\x12\r/subscription\x12\xb8\x01\n" +
	"\x12UpdateSubscription\x12*.eolymp.discussion.UpdateSubscriptionInput\x1a+.eolymp.discussion.UpdateSubscriptionOutput\"I\xea\xe2\n" +
	"\v\xf5\xe2\n" +
	"\x00\x00\x80?\xf8\xe2\n" +
	"\x05\x82\xe3\n" +
	"!\x8a\xe3\n" +
	"\x1ddiscussion:subscription:write\x82\xd3\xe4\x93\x02\x0f\x1a\r/subscriptionB7Z5github.com/eolymp/go-sdk/eolymp/discussion;discussionb\x06proto3"

var (
	file_eolymp_discussion_subscription_service_proto_rawDescOnce sync.Once
	file_eolymp_discussion_subscription_service_proto_rawDescData []byte
)

func file_eolymp_discussion_subscription_service_proto_rawDescGZIP() []byte {
	file_eolymp_discussion_subscription_service_proto_rawDescOnce.Do(func() {
		file_eolymp_discussion_subscription_service_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_eolymp_discussion_subscription_service_proto_rawDesc), len(file_eolymp_discussion_subscription_service_proto_rawDesc)))
	})
	return file_eolymp_discussion_subscription_service_proto_rawDescData
}

var file_eolymp_discussion_subscription_service_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_eolymp_discussion_subscription_service_proto_goTypes = []any{
	(*DescribeSubscriptionInput)(nil),  // 0: eolymp.discussion.DescribeSubscriptionInput
	(*DescribeSubscriptionOutput)(nil), // 1: eolymp.discussion.DescribeSubscriptionOutput
	(*UpdateSubscriptionInput)(nil),    // 2: eolymp.discussion.UpdateSubscriptionInput
	(*UpdateSubscriptionOutput)(nil),   // 3: eolymp.discussion.UpdateSubscriptionOutput
	(*Subscription)(nil),               // 4: eolymp.discussion.Subscription
}
var file_eolymp_discussion_subscription_service_proto_depIdxs = []int32{
	4, // 0: eolymp.discussion.DescribeSubscriptionOutput.subscription:type_name -> eolymp.discussion.Subscription
	4, // 1: eolymp.discussion.UpdateSubscriptionInput.subscription:type_name -> eolymp.discussion.Subscription
	0, // 2: eolymp.discussion.SubscriptionService.DescribeSubscription:input_type -> eolymp.discussion.DescribeSubscriptionInput
	2, // 3: eolymp.discussion.SubscriptionService.UpdateSubscription:input_type -> eolymp.discussion.UpdateSubscriptionInput
	1, // 4: eolymp.discussion.SubscriptionService.DescribeSubscription:output_type -> eolymp.discussion.DescribeSubscriptionOutput
	3, // 5: eolymp.discussion.SubscriptionService.UpdateSubscription:output_type -> eolymp.discussion.UpdateSubscriptionOutput
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_eolymp_discussion_subscription_service_proto_init() }
func file_eolymp_discussion_subscription_service_proto_init() {
	if File_eolymp_discussion_subscription_service_proto != nil {
		return
	}
	file_eolymp_discussion_subscription_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_eolymp_discussion_subscription_service_proto_rawDesc), len(file_eolymp_discussion_subscription_service_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_eolymp_discussion_subscription_service_proto_goTypes,
		DependencyIndexes: file_eolymp_discussion_subscription_service_proto_depIdxs,
		MessageInfos:      file_eolymp_discussion_subscription_service_proto_msgTypes,
	}.Build()
	File_eolymp_discussion_subscription_service_proto = out.File
	file_eolymp_discussion_subscription_service_proto_goTypes = nil
	file_eolymp_discussion_subscription_service_proto_depIdxs = nil
}
