// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.18.1
// source: eolymp/community/subscription_service.proto

package community

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

type ListSubscriptionsInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Offset int32 `protobuf:"varint,10,opt,name=offset,proto3" json:"offset,omitempty"`
	Size   int32 `protobuf:"varint,11,opt,name=size,proto3" json:"size,omitempty"`
}

func (x *ListSubscriptionsInput) Reset() {
	*x = ListSubscriptionsInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_community_subscription_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListSubscriptionsInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListSubscriptionsInput) ProtoMessage() {}

func (x *ListSubscriptionsInput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_community_subscription_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListSubscriptionsInput.ProtoReflect.Descriptor instead.
func (*ListSubscriptionsInput) Descriptor() ([]byte, []int) {
	return file_eolymp_community_subscription_service_proto_rawDescGZIP(), []int{0}
}

func (x *ListSubscriptionsInput) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *ListSubscriptionsInput) GetSize() int32 {
	if x != nil {
		return x.Size
	}
	return 0
}

type ListSubscriptionsOutput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Total int32           `protobuf:"varint,1,opt,name=total,proto3" json:"total,omitempty"`
	Items []*Subscription `protobuf:"bytes,2,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *ListSubscriptionsOutput) Reset() {
	*x = ListSubscriptionsOutput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_community_subscription_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListSubscriptionsOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListSubscriptionsOutput) ProtoMessage() {}

func (x *ListSubscriptionsOutput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_community_subscription_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListSubscriptionsOutput.ProtoReflect.Descriptor instead.
func (*ListSubscriptionsOutput) Descriptor() ([]byte, []int) {
	return file_eolymp_community_subscription_service_proto_rawDescGZIP(), []int{1}
}

func (x *ListSubscriptionsOutput) GetTotal() int32 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *ListSubscriptionsOutput) GetItems() []*Subscription {
	if x != nil {
		return x.Items
	}
	return nil
}

type DescribeSubscriptionInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SubscriptionId string `protobuf:"bytes,1,opt,name=subscription_id,json=subscriptionId,proto3" json:"subscription_id,omitempty"`
}

func (x *DescribeSubscriptionInput) Reset() {
	*x = DescribeSubscriptionInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_community_subscription_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DescribeSubscriptionInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescribeSubscriptionInput) ProtoMessage() {}

func (x *DescribeSubscriptionInput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_community_subscription_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
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
	return file_eolymp_community_subscription_service_proto_rawDescGZIP(), []int{2}
}

func (x *DescribeSubscriptionInput) GetSubscriptionId() string {
	if x != nil {
		return x.SubscriptionId
	}
	return ""
}

type DescribeSubscriptionOutput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Subscription *Subscription `protobuf:"bytes,1,opt,name=subscription,proto3" json:"subscription,omitempty"`
}

func (x *DescribeSubscriptionOutput) Reset() {
	*x = DescribeSubscriptionOutput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_community_subscription_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DescribeSubscriptionOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescribeSubscriptionOutput) ProtoMessage() {}

func (x *DescribeSubscriptionOutput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_community_subscription_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
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
	return file_eolymp_community_subscription_service_proto_rawDescGZIP(), []int{3}
}

func (x *DescribeSubscriptionOutput) GetSubscription() *Subscription {
	if x != nil {
		return x.Subscription
	}
	return nil
}

var File_eolymp_community_subscription_service_proto protoreflect.FileDescriptor

var file_eolymp_community_subscription_service_proto_rawDesc = []byte{
	0x0a, 0x2b, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69,
	0x74, 0x79, 0x2f, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x65,
	0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x1a,
	0x1d, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2f, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x22,
	0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x72, 0x61, 0x74, 0x65, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x23, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x75,
	0x6e, 0x69, 0x74, 0x79, 0x2f, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x44, 0x0a, 0x16, 0x4c, 0x69, 0x73, 0x74, 0x53,
	0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x49, 0x6e, 0x70, 0x75,
	0x74, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a,
	0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x22, 0x65, 0x0a,
	0x17, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61,
	0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x34,
	0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e,
	0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79,
	0x2e, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x05, 0x69,
	0x74, 0x65, 0x6d, 0x73, 0x22, 0x44, 0x0a, 0x19, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65,
	0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x70, 0x75,
	0x74, 0x12, 0x27, 0x0a, 0x0f, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x73, 0x75, 0x62, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x22, 0x60, 0x0a, 0x1a, 0x44, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x12, 0x42, 0x0a, 0x0c, 0x73, 0x75, 0x62, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e,
	0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74,
	0x79, 0x2e, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0c,
	0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x32, 0x9a, 0x03, 0x0a,
	0x13, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0xcd, 0x01, 0x0a, 0x14, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62,
	0x65, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2b, 0x2e,
	0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79,
	0x2e, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x2c, 0x2e, 0x65, 0x6f, 0x6c,
	0x79, 0x6d, 0x70, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x2e, 0x44, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x22, 0x5a, 0xea, 0xe2, 0x0a, 0x0b, 0xf5, 0xe2,
	0x0a, 0x00, 0x00, 0xa0, 0x40, 0xf8, 0xe2, 0x0a, 0x32, 0x82, 0xe3, 0x0a, 0x1f, 0x8a, 0xe3, 0x0a,
	0x1b, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x3a, 0x73, 0x75, 0x62, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x3a, 0x72, 0x65, 0x61, 0x64, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x22, 0x12, 0x20, 0x2f, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x7b, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x69, 0x64, 0x7d, 0x12, 0xb2, 0x01, 0x0a, 0x11, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x75, 0x62,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x28, 0x2e, 0x65, 0x6f, 0x6c,
	0x79, 0x6d, 0x70, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x2e, 0x4c, 0x69,
	0x73, 0x74, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x49,
	0x6e, 0x70, 0x75, 0x74, 0x1a, 0x29, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x75, 0x62, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x22,
	0x48, 0xea, 0xe2, 0x0a, 0x0b, 0xf5, 0xe2, 0x0a, 0x00, 0x00, 0xa0, 0x40, 0xf8, 0xe2, 0x0a, 0x14,
	0x82, 0xe3, 0x0a, 0x1f, 0x8a, 0xe3, 0x0a, 0x1b, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74,
	0x79, 0x3a, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x3a, 0x72,
	0x65, 0x61, 0x64, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x10, 0x12, 0x0e, 0x2f, 0x73, 0x75, 0x62, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x42, 0x35, 0x5a, 0x33, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x67,
	0x6f, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x63, 0x6f, 0x6d,
	0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x3b, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_eolymp_community_subscription_service_proto_rawDescOnce sync.Once
	file_eolymp_community_subscription_service_proto_rawDescData = file_eolymp_community_subscription_service_proto_rawDesc
)

func file_eolymp_community_subscription_service_proto_rawDescGZIP() []byte {
	file_eolymp_community_subscription_service_proto_rawDescOnce.Do(func() {
		file_eolymp_community_subscription_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_eolymp_community_subscription_service_proto_rawDescData)
	})
	return file_eolymp_community_subscription_service_proto_rawDescData
}

var file_eolymp_community_subscription_service_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_eolymp_community_subscription_service_proto_goTypes = []interface{}{
	(*ListSubscriptionsInput)(nil),     // 0: eolymp.community.ListSubscriptionsInput
	(*ListSubscriptionsOutput)(nil),    // 1: eolymp.community.ListSubscriptionsOutput
	(*DescribeSubscriptionInput)(nil),  // 2: eolymp.community.DescribeSubscriptionInput
	(*DescribeSubscriptionOutput)(nil), // 3: eolymp.community.DescribeSubscriptionOutput
	(*Subscription)(nil),               // 4: eolymp.community.Subscription
}
var file_eolymp_community_subscription_service_proto_depIdxs = []int32{
	4, // 0: eolymp.community.ListSubscriptionsOutput.items:type_name -> eolymp.community.Subscription
	4, // 1: eolymp.community.DescribeSubscriptionOutput.subscription:type_name -> eolymp.community.Subscription
	2, // 2: eolymp.community.SubscriptionService.DescribeSubscription:input_type -> eolymp.community.DescribeSubscriptionInput
	0, // 3: eolymp.community.SubscriptionService.ListSubscriptions:input_type -> eolymp.community.ListSubscriptionsInput
	3, // 4: eolymp.community.SubscriptionService.DescribeSubscription:output_type -> eolymp.community.DescribeSubscriptionOutput
	1, // 5: eolymp.community.SubscriptionService.ListSubscriptions:output_type -> eolymp.community.ListSubscriptionsOutput
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_eolymp_community_subscription_service_proto_init() }
func file_eolymp_community_subscription_service_proto_init() {
	if File_eolymp_community_subscription_service_proto != nil {
		return
	}
	file_eolymp_community_subscription_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_eolymp_community_subscription_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListSubscriptionsInput); i {
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
		file_eolymp_community_subscription_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListSubscriptionsOutput); i {
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
		file_eolymp_community_subscription_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DescribeSubscriptionInput); i {
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
		file_eolymp_community_subscription_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DescribeSubscriptionOutput); i {
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
			RawDescriptor: file_eolymp_community_subscription_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_eolymp_community_subscription_service_proto_goTypes,
		DependencyIndexes: file_eolymp_community_subscription_service_proto_depIdxs,
		MessageInfos:      file_eolymp_community_subscription_service_proto_msgTypes,
	}.Build()
	File_eolymp_community_subscription_service_proto = out.File
	file_eolymp_community_subscription_service_proto_rawDesc = nil
	file_eolymp_community_subscription_service_proto_goTypes = nil
	file_eolymp_community_subscription_service_proto_depIdxs = nil
}
