// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.4
// 	protoc        v4.24.4
// source: eolymp/community/linked_account_service.proto

package community

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

type DescribeLinkedAccountInput struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	LinkId        string                 `protobuf:"bytes,1,opt,name=link_id,json=linkId,proto3" json:"link_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DescribeLinkedAccountInput) Reset() {
	*x = DescribeLinkedAccountInput{}
	mi := &file_eolymp_community_linked_account_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DescribeLinkedAccountInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescribeLinkedAccountInput) ProtoMessage() {}

func (x *DescribeLinkedAccountInput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_community_linked_account_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DescribeLinkedAccountInput.ProtoReflect.Descriptor instead.
func (*DescribeLinkedAccountInput) Descriptor() ([]byte, []int) {
	return file_eolymp_community_linked_account_service_proto_rawDescGZIP(), []int{0}
}

func (x *DescribeLinkedAccountInput) GetLinkId() string {
	if x != nil {
		return x.LinkId
	}
	return ""
}

type DescribeLinkedAccountOutput struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Link          *LinkedAccount         `protobuf:"bytes,1,opt,name=link,proto3" json:"link,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DescribeLinkedAccountOutput) Reset() {
	*x = DescribeLinkedAccountOutput{}
	mi := &file_eolymp_community_linked_account_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DescribeLinkedAccountOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescribeLinkedAccountOutput) ProtoMessage() {}

func (x *DescribeLinkedAccountOutput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_community_linked_account_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DescribeLinkedAccountOutput.ProtoReflect.Descriptor instead.
func (*DescribeLinkedAccountOutput) Descriptor() ([]byte, []int) {
	return file_eolymp_community_linked_account_service_proto_rawDescGZIP(), []int{1}
}

func (x *DescribeLinkedAccountOutput) GetLink() *LinkedAccount {
	if x != nil {
		return x.Link
	}
	return nil
}

type ListLinkedAccountsInput struct {
	state         protoimpl.MessageState          `protogen:"open.v1"`
	Offset        int32                           `protobuf:"varint,10,opt,name=offset,proto3" json:"offset,omitempty"`
	Size          int32                           `protobuf:"varint,11,opt,name=size,proto3" json:"size,omitempty"`
	Filters       *ListLinkedAccountsInput_Filter `protobuf:"bytes,40,opt,name=filters,proto3" json:"filters,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListLinkedAccountsInput) Reset() {
	*x = ListLinkedAccountsInput{}
	mi := &file_eolymp_community_linked_account_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListLinkedAccountsInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListLinkedAccountsInput) ProtoMessage() {}

func (x *ListLinkedAccountsInput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_community_linked_account_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListLinkedAccountsInput.ProtoReflect.Descriptor instead.
func (*ListLinkedAccountsInput) Descriptor() ([]byte, []int) {
	return file_eolymp_community_linked_account_service_proto_rawDescGZIP(), []int{2}
}

func (x *ListLinkedAccountsInput) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *ListLinkedAccountsInput) GetSize() int32 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *ListLinkedAccountsInput) GetFilters() *ListLinkedAccountsInput_Filter {
	if x != nil {
		return x.Filters
	}
	return nil
}

type ListLinkedAccountsOutput struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Total         int32                  `protobuf:"varint,1,opt,name=total,proto3" json:"total,omitempty"`
	Items         []*LinkedAccount       `protobuf:"bytes,2,rep,name=items,proto3" json:"items,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListLinkedAccountsOutput) Reset() {
	*x = ListLinkedAccountsOutput{}
	mi := &file_eolymp_community_linked_account_service_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListLinkedAccountsOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListLinkedAccountsOutput) ProtoMessage() {}

func (x *ListLinkedAccountsOutput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_community_linked_account_service_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListLinkedAccountsOutput.ProtoReflect.Descriptor instead.
func (*ListLinkedAccountsOutput) Descriptor() ([]byte, []int) {
	return file_eolymp_community_linked_account_service_proto_rawDescGZIP(), []int{3}
}

func (x *ListLinkedAccountsOutput) GetTotal() int32 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *ListLinkedAccountsOutput) GetItems() []*LinkedAccount {
	if x != nil {
		return x.Items
	}
	return nil
}

type DeleteLinkedAccountInput struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	LinkId        string                 `protobuf:"bytes,1,opt,name=link_id,json=linkId,proto3" json:"link_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteLinkedAccountInput) Reset() {
	*x = DeleteLinkedAccountInput{}
	mi := &file_eolymp_community_linked_account_service_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteLinkedAccountInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteLinkedAccountInput) ProtoMessage() {}

func (x *DeleteLinkedAccountInput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_community_linked_account_service_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteLinkedAccountInput.ProtoReflect.Descriptor instead.
func (*DeleteLinkedAccountInput) Descriptor() ([]byte, []int) {
	return file_eolymp_community_linked_account_service_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteLinkedAccountInput) GetLinkId() string {
	if x != nil {
		return x.LinkId
	}
	return ""
}

type DeleteLinkedAccountOutput struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteLinkedAccountOutput) Reset() {
	*x = DeleteLinkedAccountOutput{}
	mi := &file_eolymp_community_linked_account_service_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteLinkedAccountOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteLinkedAccountOutput) ProtoMessage() {}

func (x *DeleteLinkedAccountOutput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_community_linked_account_service_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteLinkedAccountOutput.ProtoReflect.Descriptor instead.
func (*DeleteLinkedAccountOutput) Descriptor() ([]byte, []int) {
	return file_eolymp_community_linked_account_service_proto_rawDescGZIP(), []int{5}
}

type ListLinkedAccountsInput_Filter struct {
	state         protoimpl.MessageState      `protogen:"open.v1"`
	Id            []*wellknown.ExpressionID   `protobuf:"bytes,2,rep,name=id,proto3" json:"id,omitempty"`
	Type          []*wellknown.ExpressionEnum `protobuf:"bytes,3,rep,name=type,proto3" json:"type,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListLinkedAccountsInput_Filter) Reset() {
	*x = ListLinkedAccountsInput_Filter{}
	mi := &file_eolymp_community_linked_account_service_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListLinkedAccountsInput_Filter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListLinkedAccountsInput_Filter) ProtoMessage() {}

func (x *ListLinkedAccountsInput_Filter) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_community_linked_account_service_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListLinkedAccountsInput_Filter.ProtoReflect.Descriptor instead.
func (*ListLinkedAccountsInput_Filter) Descriptor() ([]byte, []int) {
	return file_eolymp_community_linked_account_service_proto_rawDescGZIP(), []int{2, 0}
}

func (x *ListLinkedAccountsInput_Filter) GetId() []*wellknown.ExpressionID {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *ListLinkedAccountsInput_Filter) GetType() []*wellknown.ExpressionEnum {
	if x != nil {
		return x.Type
	}
	return nil
}

var File_eolymp_community_linked_account_service_proto protoreflect.FileDescriptor

var file_eolymp_community_linked_account_service_proto_rawDesc = string([]byte{
	0x0a, 0x2d, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69,
	0x74, 0x79, 0x2f, 0x6c, 0x69, 0x6e, 0x6b, 0x65, 0x64, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x10, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74,
	0x79, 0x1a, 0x1d, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x22, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x72, 0x61, 0x74, 0x65, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x25, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x63, 0x6f, 0x6d,
	0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x2f, 0x6c, 0x69, 0x6e, 0x6b, 0x65, 0x64, 0x5f, 0x61, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x21, 0x65, 0x6f, 0x6c,
	0x79, 0x6d, 0x70, 0x2f, 0x77, 0x65, 0x6c, 0x6c, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x2f, 0x65, 0x78,
	0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x35,
	0x0a, 0x1a, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x4c, 0x69, 0x6e, 0x6b, 0x65, 0x64,
	0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x17, 0x0a, 0x07,
	0x6c, 0x69, 0x6e, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6c,
	0x69, 0x6e, 0x6b, 0x49, 0x64, 0x22, 0x52, 0x0a, 0x1b, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62,
	0x65, 0x4c, 0x69, 0x6e, 0x6b, 0x65, 0x64, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4f, 0x75,
	0x74, 0x70, 0x75, 0x74, 0x12, 0x33, 0x0a, 0x04, 0x6c, 0x69, 0x6e, 0x6b, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x75, 0x6e, 0x69, 0x74, 0x79, 0x2e, 0x4c, 0x69, 0x6e, 0x6b, 0x65, 0x64, 0x41, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x52, 0x04, 0x6c, 0x69, 0x6e, 0x6b, 0x22, 0x81, 0x02, 0x0a, 0x17, 0x4c, 0x69,
	0x73, 0x74, 0x4c, 0x69, 0x6e, 0x6b, 0x65, 0x64, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73,
	0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x12, 0x0a,
	0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x73, 0x69, 0x7a,
	0x65, 0x12, 0x4a, 0x0a, 0x07, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x18, 0x28, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x30, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x75, 0x6e, 0x69, 0x74, 0x79, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4c, 0x69, 0x6e, 0x6b, 0x65, 0x64,
	0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x2e, 0x46, 0x69,
	0x6c, 0x74, 0x65, 0x72, 0x52, 0x07, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x1a, 0x6e, 0x0a,
	0x06, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x2e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x77, 0x65, 0x6c,
	0x6c, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x2e, 0x45, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x49, 0x44, 0x52, 0x02, 0x69, 0x64, 0x12, 0x34, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x77,
	0x65, 0x6c, 0x6c, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x2e, 0x45, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x45, 0x6e, 0x75, 0x6d, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0x67, 0x0a,
	0x18, 0x4c, 0x69, 0x73, 0x74, 0x4c, 0x69, 0x6e, 0x6b, 0x65, 0x64, 0x41, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x73, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74,
	0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12,
	0x35, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f,
	0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74,
	0x79, 0x2e, 0x4c, 0x69, 0x6e, 0x6b, 0x65, 0x64, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52,
	0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x22, 0x33, 0x0a, 0x18, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x4c, 0x69, 0x6e, 0x6b, 0x65, 0x64, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x6e, 0x70,
	0x75, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x6c, 0x69, 0x6e, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x6c, 0x69, 0x6e, 0x6b, 0x49, 0x64, 0x22, 0x1b, 0x0a, 0x19, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x4c, 0x69, 0x6e, 0x6b, 0x65, 0x64, 0x41, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x32, 0xfb, 0x03, 0x0a, 0x14, 0x4c, 0x69, 0x6e,
	0x6b, 0x65, 0x64, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0xa7, 0x01, 0x0a, 0x15, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x4c, 0x69,
	0x6e, 0x6b, 0x65, 0x64, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x2c, 0x2e, 0x65, 0x6f,
	0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x2e, 0x44,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x4c, 0x69, 0x6e, 0x6b, 0x65, 0x64, 0x41, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x2d, 0x2e, 0x65, 0x6f, 0x6c, 0x79,
	0x6d, 0x70, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x2e, 0x44, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x62, 0x65, 0x4c, 0x69, 0x6e, 0x6b, 0x65, 0x64, 0x41, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x22, 0x31, 0xea, 0xe2, 0x0a, 0x0b, 0xf5, 0xe2,
	0x0a, 0x00, 0x00, 0xa0, 0x40, 0xf8, 0xe2, 0x0a, 0x14, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1c, 0x12,
	0x1a, 0x2f, 0x6c, 0x69, 0x6e, 0x6b, 0x65, 0x64, 0x2d, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x73, 0x2f, 0x7b, 0x6c, 0x69, 0x6e, 0x6b, 0x5f, 0x69, 0x64, 0x7d, 0x12, 0x94, 0x01, 0x0a, 0x12,
	0x4c, 0x69, 0x73, 0x74, 0x4c, 0x69, 0x6e, 0x6b, 0x65, 0x64, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x73, 0x12, 0x29, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x75, 0x6e, 0x69, 0x74, 0x79, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4c, 0x69, 0x6e, 0x6b, 0x65, 0x64,
	0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x2a, 0x2e,
	0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79,
	0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4c, 0x69, 0x6e, 0x6b, 0x65, 0x64, 0x41, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x73, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x22, 0x27, 0xea, 0xe2, 0x0a, 0x0b, 0xf5,
	0xe2, 0x0a, 0x00, 0x00, 0xa0, 0x40, 0xf8, 0xe2, 0x0a, 0x14, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x12,
	0x12, 0x10, 0x2f, 0x6c, 0x69, 0x6e, 0x6b, 0x65, 0x64, 0x2d, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x73, 0x12, 0xa1, 0x01, 0x0a, 0x13, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4c, 0x69, 0x6e,
	0x6b, 0x65, 0x64, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x2a, 0x2e, 0x65, 0x6f, 0x6c,
	0x79, 0x6d, 0x70, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x2e, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x4c, 0x69, 0x6e, 0x6b, 0x65, 0x64, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x2b, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x4c, 0x69, 0x6e, 0x6b, 0x65, 0x64, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4f, 0x75, 0x74,
	0x70, 0x75, 0x74, 0x22, 0x31, 0xea, 0xe2, 0x0a, 0x0b, 0xf5, 0xe2, 0x0a, 0x00, 0x00, 0xa0, 0x40,
	0xf8, 0xe2, 0x0a, 0x14, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1c, 0x2a, 0x1a, 0x2f, 0x6c, 0x69, 0x6e,
	0x6b, 0x65, 0x64, 0x2d, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x2f, 0x7b, 0x6c, 0x69,
	0x6e, 0x6b, 0x5f, 0x69, 0x64, 0x7d, 0x42, 0x35, 0x5a, 0x33, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x67, 0x6f, 0x2d, 0x73,
	0x64, 0x6b, 0x2f, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e,
	0x69, 0x74, 0x79, 0x3b, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_eolymp_community_linked_account_service_proto_rawDescOnce sync.Once
	file_eolymp_community_linked_account_service_proto_rawDescData []byte
)

func file_eolymp_community_linked_account_service_proto_rawDescGZIP() []byte {
	file_eolymp_community_linked_account_service_proto_rawDescOnce.Do(func() {
		file_eolymp_community_linked_account_service_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_eolymp_community_linked_account_service_proto_rawDesc), len(file_eolymp_community_linked_account_service_proto_rawDesc)))
	})
	return file_eolymp_community_linked_account_service_proto_rawDescData
}

var file_eolymp_community_linked_account_service_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_eolymp_community_linked_account_service_proto_goTypes = []any{
	(*DescribeLinkedAccountInput)(nil),     // 0: eolymp.community.DescribeLinkedAccountInput
	(*DescribeLinkedAccountOutput)(nil),    // 1: eolymp.community.DescribeLinkedAccountOutput
	(*ListLinkedAccountsInput)(nil),        // 2: eolymp.community.ListLinkedAccountsInput
	(*ListLinkedAccountsOutput)(nil),       // 3: eolymp.community.ListLinkedAccountsOutput
	(*DeleteLinkedAccountInput)(nil),       // 4: eolymp.community.DeleteLinkedAccountInput
	(*DeleteLinkedAccountOutput)(nil),      // 5: eolymp.community.DeleteLinkedAccountOutput
	(*ListLinkedAccountsInput_Filter)(nil), // 6: eolymp.community.ListLinkedAccountsInput.Filter
	(*LinkedAccount)(nil),                  // 7: eolymp.community.LinkedAccount
	(*wellknown.ExpressionID)(nil),         // 8: eolymp.wellknown.ExpressionID
	(*wellknown.ExpressionEnum)(nil),       // 9: eolymp.wellknown.ExpressionEnum
}
var file_eolymp_community_linked_account_service_proto_depIdxs = []int32{
	7, // 0: eolymp.community.DescribeLinkedAccountOutput.link:type_name -> eolymp.community.LinkedAccount
	6, // 1: eolymp.community.ListLinkedAccountsInput.filters:type_name -> eolymp.community.ListLinkedAccountsInput.Filter
	7, // 2: eolymp.community.ListLinkedAccountsOutput.items:type_name -> eolymp.community.LinkedAccount
	8, // 3: eolymp.community.ListLinkedAccountsInput.Filter.id:type_name -> eolymp.wellknown.ExpressionID
	9, // 4: eolymp.community.ListLinkedAccountsInput.Filter.type:type_name -> eolymp.wellknown.ExpressionEnum
	0, // 5: eolymp.community.LinkedAccountService.DescribeLinkedAccount:input_type -> eolymp.community.DescribeLinkedAccountInput
	2, // 6: eolymp.community.LinkedAccountService.ListLinkedAccounts:input_type -> eolymp.community.ListLinkedAccountsInput
	4, // 7: eolymp.community.LinkedAccountService.DeleteLinkedAccount:input_type -> eolymp.community.DeleteLinkedAccountInput
	1, // 8: eolymp.community.LinkedAccountService.DescribeLinkedAccount:output_type -> eolymp.community.DescribeLinkedAccountOutput
	3, // 9: eolymp.community.LinkedAccountService.ListLinkedAccounts:output_type -> eolymp.community.ListLinkedAccountsOutput
	5, // 10: eolymp.community.LinkedAccountService.DeleteLinkedAccount:output_type -> eolymp.community.DeleteLinkedAccountOutput
	8, // [8:11] is the sub-list for method output_type
	5, // [5:8] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_eolymp_community_linked_account_service_proto_init() }
func file_eolymp_community_linked_account_service_proto_init() {
	if File_eolymp_community_linked_account_service_proto != nil {
		return
	}
	file_eolymp_community_linked_account_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_eolymp_community_linked_account_service_proto_rawDesc), len(file_eolymp_community_linked_account_service_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_eolymp_community_linked_account_service_proto_goTypes,
		DependencyIndexes: file_eolymp_community_linked_account_service_proto_depIdxs,
		MessageInfos:      file_eolymp_community_linked_account_service_proto_msgTypes,
	}.Build()
	File_eolymp_community_linked_account_service_proto = out.File
	file_eolymp_community_linked_account_service_proto_goTypes = nil
	file_eolymp_community_linked_account_service_proto_depIdxs = nil
}
