// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.3
// 	protoc        v4.24.4
// source: eolymp/discussion/subscription_service.proto

package discussion

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

var file_eolymp_discussion_subscription_service_proto_rawDesc = []byte{
	0x0a, 0x2c, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x64, 0x69, 0x73, 0x63, 0x75, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x2f, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11,
	0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x75, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x1a, 0x1d, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x22, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x72, 0x61, 0x74, 0x65, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x24, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x64, 0x69, 0x73,
	0x63, 0x75, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2f, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x1b, 0x0a, 0x19, 0x44, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x22, 0x61, 0x0a, 0x1a, 0x44, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x62, 0x65, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x4f,
	0x75, 0x74, 0x70, 0x75, 0x74, 0x12, 0x43, 0x0a, 0x0c, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x65, 0x6f,
	0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x75, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e,
	0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0c, 0x73, 0x75,
	0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x5e, 0x0a, 0x17, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x43, 0x0a, 0x0c, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x65, 0x6f,
	0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x75, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e,
	0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0c, 0x73, 0x75,
	0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x1a, 0x0a, 0x18, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x32, 0x91, 0x03, 0x0a, 0x13, 0x53, 0x75, 0x62, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0xbe,
	0x01, 0x0a, 0x14, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x53, 0x75, 0x62, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2c, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70,
	0x2e, 0x64, 0x69, 0x73, 0x63, 0x75, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x44, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x62, 0x65, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x2d, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x64,
	0x69, 0x73, 0x63, 0x75, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x62, 0x65, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x75,
	0x74, 0x70, 0x75, 0x74, 0x22, 0x49, 0xea, 0xe2, 0x0a, 0x0c, 0xf5, 0xe2, 0x0a, 0x00, 0x00, 0xa0,
	0x41, 0xf8, 0xe2, 0x0a, 0xf4, 0x03, 0x82, 0xe3, 0x0a, 0x20, 0x8a, 0xe3, 0x0a, 0x1c, 0x64, 0x69,
	0x73, 0x63, 0x75, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x3a, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x3a, 0x72, 0x65, 0x61, 0x64, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0f,
	0x12, 0x0d, 0x2f, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0xb8, 0x01, 0x0a, 0x12, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2a, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e,
	0x64, 0x69, 0x73, 0x63, 0x75, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x70,
	0x75, 0x74, 0x1a, 0x2b, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x64, 0x69, 0x73, 0x63,
	0x75, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x75, 0x62,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x22,
	0x49, 0xea, 0xe2, 0x0a, 0x0b, 0xf5, 0xe2, 0x0a, 0x00, 0x00, 0x80, 0x3f, 0xf8, 0xe2, 0x0a, 0x05,
	0x82, 0xe3, 0x0a, 0x21, 0x8a, 0xe3, 0x0a, 0x1d, 0x64, 0x69, 0x73, 0x63, 0x75, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x3a, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x3a,
	0x77, 0x72, 0x69, 0x74, 0x65, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0f, 0x1a, 0x0d, 0x2f, 0x73, 0x75,
	0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x37, 0x5a, 0x35, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f,
	0x67, 0x6f, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x64, 0x69,
	0x73, 0x63, 0x75, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x3b, 0x64, 0x69, 0x73, 0x63, 0x75, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_eolymp_discussion_subscription_service_proto_rawDescOnce sync.Once
	file_eolymp_discussion_subscription_service_proto_rawDescData = file_eolymp_discussion_subscription_service_proto_rawDesc
)

func file_eolymp_discussion_subscription_service_proto_rawDescGZIP() []byte {
	file_eolymp_discussion_subscription_service_proto_rawDescOnce.Do(func() {
		file_eolymp_discussion_subscription_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_eolymp_discussion_subscription_service_proto_rawDescData)
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
			RawDescriptor: file_eolymp_discussion_subscription_service_proto_rawDesc,
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
	file_eolymp_discussion_subscription_service_proto_rawDesc = nil
	file_eolymp_discussion_subscription_service_proto_goTypes = nil
	file_eolymp_discussion_subscription_service_proto_depIdxs = nil
}
