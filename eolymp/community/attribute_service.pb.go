// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v4.24.4
// source: eolymp/community/attribute_service.proto

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

type CreateAttributeInput struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	AttributeKey  string                 `protobuf:"bytes,1,opt,name=attribute_key,json=attributeKey,proto3" json:"attribute_key,omitempty"`
	Attribute     *Attribute             `protobuf:"bytes,2,opt,name=attribute,proto3" json:"attribute,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateAttributeInput) Reset() {
	*x = CreateAttributeInput{}
	mi := &file_eolymp_community_attribute_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateAttributeInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAttributeInput) ProtoMessage() {}

func (x *CreateAttributeInput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_community_attribute_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateAttributeInput.ProtoReflect.Descriptor instead.
func (*CreateAttributeInput) Descriptor() ([]byte, []int) {
	return file_eolymp_community_attribute_service_proto_rawDescGZIP(), []int{0}
}

func (x *CreateAttributeInput) GetAttributeKey() string {
	if x != nil {
		return x.AttributeKey
	}
	return ""
}

func (x *CreateAttributeInput) GetAttribute() *Attribute {
	if x != nil {
		return x.Attribute
	}
	return nil
}

type CreateAttributeOutput struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateAttributeOutput) Reset() {
	*x = CreateAttributeOutput{}
	mi := &file_eolymp_community_attribute_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateAttributeOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAttributeOutput) ProtoMessage() {}

func (x *CreateAttributeOutput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_community_attribute_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateAttributeOutput.ProtoReflect.Descriptor instead.
func (*CreateAttributeOutput) Descriptor() ([]byte, []int) {
	return file_eolymp_community_attribute_service_proto_rawDescGZIP(), []int{1}
}

type UpdateAttributeInput struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Patch         []Attribute_Patch      `protobuf:"varint,3,rep,packed,name=patch,proto3,enum=eolymp.community.Attribute_Patch" json:"patch,omitempty"`
	AttributeKey  string                 `protobuf:"bytes,1,opt,name=attribute_key,json=attributeKey,proto3" json:"attribute_key,omitempty"`
	Attribute     *Attribute             `protobuf:"bytes,2,opt,name=attribute,proto3" json:"attribute,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateAttributeInput) Reset() {
	*x = UpdateAttributeInput{}
	mi := &file_eolymp_community_attribute_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateAttributeInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateAttributeInput) ProtoMessage() {}

func (x *UpdateAttributeInput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_community_attribute_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateAttributeInput.ProtoReflect.Descriptor instead.
func (*UpdateAttributeInput) Descriptor() ([]byte, []int) {
	return file_eolymp_community_attribute_service_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateAttributeInput) GetPatch() []Attribute_Patch {
	if x != nil {
		return x.Patch
	}
	return nil
}

func (x *UpdateAttributeInput) GetAttributeKey() string {
	if x != nil {
		return x.AttributeKey
	}
	return ""
}

func (x *UpdateAttributeInput) GetAttribute() *Attribute {
	if x != nil {
		return x.Attribute
	}
	return nil
}

type UpdateAttributeOutput struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateAttributeOutput) Reset() {
	*x = UpdateAttributeOutput{}
	mi := &file_eolymp_community_attribute_service_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateAttributeOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateAttributeOutput) ProtoMessage() {}

func (x *UpdateAttributeOutput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_community_attribute_service_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateAttributeOutput.ProtoReflect.Descriptor instead.
func (*UpdateAttributeOutput) Descriptor() ([]byte, []int) {
	return file_eolymp_community_attribute_service_proto_rawDescGZIP(), []int{3}
}

type RemoveAttributeInput struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	AttributeKey  string                 `protobuf:"bytes,1,opt,name=attribute_key,json=attributeKey,proto3" json:"attribute_key,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RemoveAttributeInput) Reset() {
	*x = RemoveAttributeInput{}
	mi := &file_eolymp_community_attribute_service_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RemoveAttributeInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveAttributeInput) ProtoMessage() {}

func (x *RemoveAttributeInput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_community_attribute_service_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveAttributeInput.ProtoReflect.Descriptor instead.
func (*RemoveAttributeInput) Descriptor() ([]byte, []int) {
	return file_eolymp_community_attribute_service_proto_rawDescGZIP(), []int{4}
}

func (x *RemoveAttributeInput) GetAttributeKey() string {
	if x != nil {
		return x.AttributeKey
	}
	return ""
}

type RemoveAttributeOutput struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RemoveAttributeOutput) Reset() {
	*x = RemoveAttributeOutput{}
	mi := &file_eolymp_community_attribute_service_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RemoveAttributeOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveAttributeOutput) ProtoMessage() {}

func (x *RemoveAttributeOutput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_community_attribute_service_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveAttributeOutput.ProtoReflect.Descriptor instead.
func (*RemoveAttributeOutput) Descriptor() ([]byte, []int) {
	return file_eolymp_community_attribute_service_proto_rawDescGZIP(), []int{5}
}

type DescribeAttributeInput struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	AttributeKey  string                 `protobuf:"bytes,1,opt,name=attribute_key,json=attributeKey,proto3" json:"attribute_key,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DescribeAttributeInput) Reset() {
	*x = DescribeAttributeInput{}
	mi := &file_eolymp_community_attribute_service_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DescribeAttributeInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescribeAttributeInput) ProtoMessage() {}

func (x *DescribeAttributeInput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_community_attribute_service_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DescribeAttributeInput.ProtoReflect.Descriptor instead.
func (*DescribeAttributeInput) Descriptor() ([]byte, []int) {
	return file_eolymp_community_attribute_service_proto_rawDescGZIP(), []int{6}
}

func (x *DescribeAttributeInput) GetAttributeKey() string {
	if x != nil {
		return x.AttributeKey
	}
	return ""
}

type DescribeAttributeOutput struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Attribute     *Attribute             `protobuf:"bytes,1,opt,name=attribute,proto3" json:"attribute,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DescribeAttributeOutput) Reset() {
	*x = DescribeAttributeOutput{}
	mi := &file_eolymp_community_attribute_service_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DescribeAttributeOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescribeAttributeOutput) ProtoMessage() {}

func (x *DescribeAttributeOutput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_community_attribute_service_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DescribeAttributeOutput.ProtoReflect.Descriptor instead.
func (*DescribeAttributeOutput) Descriptor() ([]byte, []int) {
	return file_eolymp_community_attribute_service_proto_rawDescGZIP(), []int{7}
}

func (x *DescribeAttributeOutput) GetAttribute() *Attribute {
	if x != nil {
		return x.Attribute
	}
	return nil
}

type ListAttributesInput struct {
	state         protoimpl.MessageState      `protogen:"open.v1"`
	Offset        int32                       `protobuf:"varint,10,opt,name=offset,proto3" json:"offset,omitempty"`
	Size          int32                       `protobuf:"varint,11,opt,name=size,proto3" json:"size,omitempty"`
	Filters       *ListAttributesInput_Filter `protobuf:"bytes,40,opt,name=filters,proto3" json:"filters,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListAttributesInput) Reset() {
	*x = ListAttributesInput{}
	mi := &file_eolymp_community_attribute_service_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListAttributesInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListAttributesInput) ProtoMessage() {}

func (x *ListAttributesInput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_community_attribute_service_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListAttributesInput.ProtoReflect.Descriptor instead.
func (*ListAttributesInput) Descriptor() ([]byte, []int) {
	return file_eolymp_community_attribute_service_proto_rawDescGZIP(), []int{8}
}

func (x *ListAttributesInput) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *ListAttributesInput) GetSize() int32 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *ListAttributesInput) GetFilters() *ListAttributesInput_Filter {
	if x != nil {
		return x.Filters
	}
	return nil
}

type ListAttributesOutput struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Total         int32                  `protobuf:"varint,1,opt,name=total,proto3" json:"total,omitempty"`
	Items         []*Attribute           `protobuf:"bytes,2,rep,name=items,proto3" json:"items,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListAttributesOutput) Reset() {
	*x = ListAttributesOutput{}
	mi := &file_eolymp_community_attribute_service_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListAttributesOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListAttributesOutput) ProtoMessage() {}

func (x *ListAttributesOutput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_community_attribute_service_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListAttributesOutput.ProtoReflect.Descriptor instead.
func (*ListAttributesOutput) Descriptor() ([]byte, []int) {
	return file_eolymp_community_attribute_service_proto_rawDescGZIP(), []int{9}
}

func (x *ListAttributesOutput) GetTotal() int32 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *ListAttributesOutput) GetItems() []*Attribute {
	if x != nil {
		return x.Items
	}
	return nil
}

type ListAttributesInput_Filter struct {
	state         protoimpl.MessageState      `protogen:"open.v1"`
	Id            []*wellknown.ExpressionID   `protobuf:"bytes,1,rep,name=id,proto3" json:"id,omitempty"`
	Key           []*wellknown.ExpressionEnum `protobuf:"bytes,2,rep,name=key,proto3" json:"key,omitempty"`
	Hidden        []*wellknown.ExpressionBool `protobuf:"bytes,3,rep,name=hidden,proto3" json:"hidden,omitempty"`
	Required      []*wellknown.ExpressionBool `protobuf:"bytes,4,rep,name=required,proto3" json:"required,omitempty"`
	Type          []*wellknown.ExpressionEnum `protobuf:"bytes,5,rep,name=type,proto3" json:"type,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListAttributesInput_Filter) Reset() {
	*x = ListAttributesInput_Filter{}
	mi := &file_eolymp_community_attribute_service_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListAttributesInput_Filter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListAttributesInput_Filter) ProtoMessage() {}

func (x *ListAttributesInput_Filter) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_community_attribute_service_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListAttributesInput_Filter.ProtoReflect.Descriptor instead.
func (*ListAttributesInput_Filter) Descriptor() ([]byte, []int) {
	return file_eolymp_community_attribute_service_proto_rawDescGZIP(), []int{8, 0}
}

func (x *ListAttributesInput_Filter) GetId() []*wellknown.ExpressionID {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *ListAttributesInput_Filter) GetKey() []*wellknown.ExpressionEnum {
	if x != nil {
		return x.Key
	}
	return nil
}

func (x *ListAttributesInput_Filter) GetHidden() []*wellknown.ExpressionBool {
	if x != nil {
		return x.Hidden
	}
	return nil
}

func (x *ListAttributesInput_Filter) GetRequired() []*wellknown.ExpressionBool {
	if x != nil {
		return x.Required
	}
	return nil
}

func (x *ListAttributesInput_Filter) GetType() []*wellknown.ExpressionEnum {
	if x != nil {
		return x.Type
	}
	return nil
}

var File_eolymp_community_attribute_service_proto protoreflect.FileDescriptor

var file_eolymp_community_attribute_service_proto_rawDesc = string([]byte{
	0x0a, 0x28, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69,
	0x74, 0x79, 0x2f, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x65, 0x6f, 0x6c, 0x79,
	0x6d, 0x70, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x1a, 0x1d, 0x65, 0x6f,
	0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2f, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x22, 0x65, 0x6f, 0x6c,
	0x79, 0x6d, 0x70, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f,
	0x72, 0x61, 0x74, 0x65, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x20, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74,
	0x79, 0x2f, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x21, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x77, 0x65, 0x6c, 0x6c, 0x6b, 0x6e,
	0x6f, 0x77, 0x6e, 0x2f, 0x65, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x76, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x74,
	0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x23, 0x0a, 0x0d,
	0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0c, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x4b, 0x65,
	0x79, 0x12, 0x39, 0x0a, 0x09, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x2e, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74,
	0x65, 0x52, 0x09, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x22, 0x17, 0x0a, 0x15,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x4f,
	0x75, 0x74, 0x70, 0x75, 0x74, 0x22, 0xaf, 0x01, 0x0a, 0x14, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x37,
	0x0a, 0x05, 0x70, 0x61, 0x74, 0x63, 0x68, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x21, 0x2e,
	0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79,
	0x2e, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x2e, 0x50, 0x61, 0x74, 0x63, 0x68,
	0x52, 0x05, 0x70, 0x61, 0x74, 0x63, 0x68, 0x12, 0x23, 0x0a, 0x0d, 0x61, 0x74, 0x74, 0x72, 0x69,
	0x62, 0x75, 0x74, 0x65, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c,
	0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x4b, 0x65, 0x79, 0x12, 0x39, 0x0a, 0x09,
	0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1b, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69,
	0x74, 0x79, 0x2e, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x52, 0x09, 0x61, 0x74,
	0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x22, 0x17, 0x0a, 0x15, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74,
	0x22, 0x3b, 0x0a, 0x14, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62,
	0x75, 0x74, 0x65, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x61, 0x74, 0x74, 0x72,
	0x69, 0x62, 0x75, 0x74, 0x65, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0c, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x4b, 0x65, 0x79, 0x22, 0x17, 0x0a,
	0x15, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65,
	0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x22, 0x3d, 0x0a, 0x16, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x62, 0x65, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x49, 0x6e, 0x70, 0x75, 0x74,
	0x12, 0x23, 0x0a, 0x0d, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x5f, 0x6b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75,
	0x74, 0x65, 0x4b, 0x65, 0x79, 0x22, 0x54, 0x0a, 0x17, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62,
	0x65, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74,
	0x12, 0x39, 0x0a, 0x09, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x2e, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65,
	0x52, 0x09, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x22, 0xa6, 0x03, 0x0a, 0x13,
	0x4c, 0x69, 0x73, 0x74, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x49, 0x6e,
	0x70, 0x75, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x73,
	0x69, 0x7a, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x12,
	0x46, 0x0a, 0x07, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x18, 0x28, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x2c, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e,
	0x69, 0x74, 0x79, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74,
	0x65, 0x73, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x2e, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x52, 0x07,
	0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x1a, 0x9a, 0x02, 0x0a, 0x06, 0x46, 0x69, 0x6c, 0x74,
	0x65, 0x72, 0x12, 0x2e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e,
	0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x77, 0x65, 0x6c, 0x6c, 0x6b, 0x6e, 0x6f, 0x77,
	0x6e, 0x2e, 0x45, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x32, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x20, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x77, 0x65, 0x6c, 0x6c, 0x6b, 0x6e, 0x6f,
	0x77, 0x6e, 0x2e, 0x45, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x45, 0x6e, 0x75,
	0x6d, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x38, 0x0a, 0x06, 0x68, 0x69, 0x64, 0x64, 0x65, 0x6e,
	0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e,
	0x77, 0x65, 0x6c, 0x6c, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x2e, 0x45, 0x78, 0x70, 0x72, 0x65, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x42, 0x6f, 0x6f, 0x6c, 0x52, 0x06, 0x68, 0x69, 0x64, 0x64, 0x65, 0x6e,
	0x12, 0x3c, 0x0a, 0x08, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x18, 0x04, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x20, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x77, 0x65, 0x6c, 0x6c,
	0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x2e, 0x45, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x42, 0x6f, 0x6f, 0x6c, 0x52, 0x08, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x12, 0x34,
	0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x65,
	0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x77, 0x65, 0x6c, 0x6c, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x2e,
	0x45, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x45, 0x6e, 0x75, 0x6d, 0x52, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x22, 0x5f, 0x0a, 0x14, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x74, 0x74, 0x72,
	0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x12, 0x14, 0x0a, 0x05,
	0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x74, 0x6f, 0x74,
	0x61, 0x6c, 0x12, 0x31, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x1b, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x75,
	0x6e, 0x69, 0x74, 0x79, 0x2e, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x52, 0x05,
	0x69, 0x74, 0x65, 0x6d, 0x73, 0x32, 0x95, 0x07, 0x0a, 0x10, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62,
	0x75, 0x74, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0xa7, 0x01, 0x0a, 0x0f, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x12, 0x26,
	0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74,
	0x79, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74,
	0x65, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x27, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x22,
	0x43, 0xea, 0xe2, 0x0a, 0x0b, 0xf5, 0xe2, 0x0a, 0x00, 0x00, 0xa0, 0x40, 0xf8, 0xe2, 0x0a, 0x14,
	0x82, 0xe3, 0x0a, 0x1d, 0x8a, 0xe3, 0x0a, 0x19, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74,
	0x79, 0x3a, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x3a, 0x77, 0x72, 0x69, 0x74,
	0x65, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0d, 0x22, 0x0b, 0x2f, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62,
	0x75, 0x74, 0x65, 0x73, 0x12, 0xb7, 0x01, 0x0a, 0x0f, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41,
	0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x12, 0x26, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d,
	0x70, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x2e, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x49, 0x6e, 0x70, 0x75, 0x74,
	0x1a, 0x27, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e,
	0x69, 0x74, 0x79, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62,
	0x75, 0x74, 0x65, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x22, 0x53, 0xea, 0xe2, 0x0a, 0x0b, 0xf5,
	0xe2, 0x0a, 0x00, 0x00, 0xa0, 0x40, 0xf8, 0xe2, 0x0a, 0x14, 0x82, 0xe3, 0x0a, 0x1d, 0x8a, 0xe3,
	0x0a, 0x19, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x3a, 0x61, 0x74, 0x74, 0x72,
	0x69, 0x62, 0x75, 0x74, 0x65, 0x3a, 0x77, 0x72, 0x69, 0x74, 0x65, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x1d, 0x22, 0x1b, 0x2f, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x2f, 0x7b,
	0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x5f, 0x6b, 0x65, 0x79, 0x7d, 0x12, 0xb7,
	0x01, 0x0a, 0x0f, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75,
	0x74, 0x65, 0x12, 0x26, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x75, 0x6e, 0x69, 0x74, 0x79, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x41, 0x74, 0x74, 0x72,
	0x69, 0x62, 0x75, 0x74, 0x65, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x27, 0x2e, 0x65, 0x6f, 0x6c,
	0x79, 0x6d, 0x70, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x2e, 0x52, 0x65,
	0x6d, 0x6f, 0x76, 0x65, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x4f, 0x75, 0x74,
	0x70, 0x75, 0x74, 0x22, 0x53, 0xea, 0xe2, 0x0a, 0x0b, 0xf5, 0xe2, 0x0a, 0x00, 0x00, 0xa0, 0x40,
	0xf8, 0xe2, 0x0a, 0x14, 0x82, 0xe3, 0x0a, 0x1d, 0x8a, 0xe3, 0x0a, 0x19, 0x63, 0x6f, 0x6d, 0x6d,
	0x75, 0x6e, 0x69, 0x74, 0x79, 0x3a, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x3a,
	0x77, 0x72, 0x69, 0x74, 0x65, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1d, 0x2a, 0x1b, 0x2f, 0x61, 0x74,
	0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x2f, 0x7b, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62,
	0x75, 0x74, 0x65, 0x5f, 0x6b, 0x65, 0x79, 0x7d, 0x12, 0xbc, 0x01, 0x0a, 0x11, 0x44, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x62, 0x65, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x12, 0x28,
	0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74,
	0x79, 0x2e, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62,
	0x75, 0x74, 0x65, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x29, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d,
	0x70, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x2e, 0x44, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x62, 0x65, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x4f, 0x75, 0x74,
	0x70, 0x75, 0x74, 0x22, 0x52, 0xea, 0xe2, 0x0a, 0x0b, 0xf5, 0xe2, 0x0a, 0x00, 0x00, 0xa0, 0x40,
	0xf8, 0xe2, 0x0a, 0x14, 0x82, 0xe3, 0x0a, 0x1c, 0x8a, 0xe3, 0x0a, 0x18, 0x63, 0x6f, 0x6d, 0x6d,
	0x75, 0x6e, 0x69, 0x74, 0x79, 0x3a, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x3a,
	0x72, 0x65, 0x61, 0x64, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1d, 0x12, 0x1b, 0x2f, 0x61, 0x74, 0x74,
	0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x2f, 0x7b, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75,
	0x74, 0x65, 0x5f, 0x6b, 0x65, 0x79, 0x7d, 0x12, 0xa3, 0x01, 0x0a, 0x0e, 0x4c, 0x69, 0x73, 0x74,
	0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x12, 0x25, 0x2e, 0x65, 0x6f, 0x6c,
	0x79, 0x6d, 0x70, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x2e, 0x4c, 0x69,
	0x73, 0x74, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x49, 0x6e, 0x70, 0x75,
	0x74, 0x1a, 0x26, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x75,
	0x6e, 0x69, 0x74, 0x79, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75,
	0x74, 0x65, 0x73, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x22, 0x42, 0xea, 0xe2, 0x0a, 0x0b, 0xf5,
	0xe2, 0x0a, 0x00, 0x00, 0xa0, 0x40, 0xf8, 0xe2, 0x0a, 0x14, 0x82, 0xe3, 0x0a, 0x1c, 0x8a, 0xe3,
	0x0a, 0x18, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x3a, 0x61, 0x74, 0x74, 0x72,
	0x69, 0x62, 0x75, 0x74, 0x65, 0x3a, 0x72, 0x65, 0x61, 0x64, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0d,
	0x12, 0x0b, 0x2f, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x42, 0x35, 0x5a,
	0x33, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6f, 0x6c, 0x79,
	0x6d, 0x70, 0x2f, 0x67, 0x6f, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70,
	0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x3b, 0x63, 0x6f, 0x6d, 0x6d, 0x75,
	0x6e, 0x69, 0x74, 0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_eolymp_community_attribute_service_proto_rawDescOnce sync.Once
	file_eolymp_community_attribute_service_proto_rawDescData []byte
)

func file_eolymp_community_attribute_service_proto_rawDescGZIP() []byte {
	file_eolymp_community_attribute_service_proto_rawDescOnce.Do(func() {
		file_eolymp_community_attribute_service_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_eolymp_community_attribute_service_proto_rawDesc), len(file_eolymp_community_attribute_service_proto_rawDesc)))
	})
	return file_eolymp_community_attribute_service_proto_rawDescData
}

var file_eolymp_community_attribute_service_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_eolymp_community_attribute_service_proto_goTypes = []any{
	(*CreateAttributeInput)(nil),       // 0: eolymp.community.CreateAttributeInput
	(*CreateAttributeOutput)(nil),      // 1: eolymp.community.CreateAttributeOutput
	(*UpdateAttributeInput)(nil),       // 2: eolymp.community.UpdateAttributeInput
	(*UpdateAttributeOutput)(nil),      // 3: eolymp.community.UpdateAttributeOutput
	(*RemoveAttributeInput)(nil),       // 4: eolymp.community.RemoveAttributeInput
	(*RemoveAttributeOutput)(nil),      // 5: eolymp.community.RemoveAttributeOutput
	(*DescribeAttributeInput)(nil),     // 6: eolymp.community.DescribeAttributeInput
	(*DescribeAttributeOutput)(nil),    // 7: eolymp.community.DescribeAttributeOutput
	(*ListAttributesInput)(nil),        // 8: eolymp.community.ListAttributesInput
	(*ListAttributesOutput)(nil),       // 9: eolymp.community.ListAttributesOutput
	(*ListAttributesInput_Filter)(nil), // 10: eolymp.community.ListAttributesInput.Filter
	(*Attribute)(nil),                  // 11: eolymp.community.Attribute
	(Attribute_Patch)(0),               // 12: eolymp.community.Attribute.Patch
	(*wellknown.ExpressionID)(nil),     // 13: eolymp.wellknown.ExpressionID
	(*wellknown.ExpressionEnum)(nil),   // 14: eolymp.wellknown.ExpressionEnum
	(*wellknown.ExpressionBool)(nil),   // 15: eolymp.wellknown.ExpressionBool
}
var file_eolymp_community_attribute_service_proto_depIdxs = []int32{
	11, // 0: eolymp.community.CreateAttributeInput.attribute:type_name -> eolymp.community.Attribute
	12, // 1: eolymp.community.UpdateAttributeInput.patch:type_name -> eolymp.community.Attribute.Patch
	11, // 2: eolymp.community.UpdateAttributeInput.attribute:type_name -> eolymp.community.Attribute
	11, // 3: eolymp.community.DescribeAttributeOutput.attribute:type_name -> eolymp.community.Attribute
	10, // 4: eolymp.community.ListAttributesInput.filters:type_name -> eolymp.community.ListAttributesInput.Filter
	11, // 5: eolymp.community.ListAttributesOutput.items:type_name -> eolymp.community.Attribute
	13, // 6: eolymp.community.ListAttributesInput.Filter.id:type_name -> eolymp.wellknown.ExpressionID
	14, // 7: eolymp.community.ListAttributesInput.Filter.key:type_name -> eolymp.wellknown.ExpressionEnum
	15, // 8: eolymp.community.ListAttributesInput.Filter.hidden:type_name -> eolymp.wellknown.ExpressionBool
	15, // 9: eolymp.community.ListAttributesInput.Filter.required:type_name -> eolymp.wellknown.ExpressionBool
	14, // 10: eolymp.community.ListAttributesInput.Filter.type:type_name -> eolymp.wellknown.ExpressionEnum
	0,  // 11: eolymp.community.AttributeService.CreateAttribute:input_type -> eolymp.community.CreateAttributeInput
	2,  // 12: eolymp.community.AttributeService.UpdateAttribute:input_type -> eolymp.community.UpdateAttributeInput
	4,  // 13: eolymp.community.AttributeService.RemoveAttribute:input_type -> eolymp.community.RemoveAttributeInput
	6,  // 14: eolymp.community.AttributeService.DescribeAttribute:input_type -> eolymp.community.DescribeAttributeInput
	8,  // 15: eolymp.community.AttributeService.ListAttributes:input_type -> eolymp.community.ListAttributesInput
	1,  // 16: eolymp.community.AttributeService.CreateAttribute:output_type -> eolymp.community.CreateAttributeOutput
	3,  // 17: eolymp.community.AttributeService.UpdateAttribute:output_type -> eolymp.community.UpdateAttributeOutput
	5,  // 18: eolymp.community.AttributeService.RemoveAttribute:output_type -> eolymp.community.RemoveAttributeOutput
	7,  // 19: eolymp.community.AttributeService.DescribeAttribute:output_type -> eolymp.community.DescribeAttributeOutput
	9,  // 20: eolymp.community.AttributeService.ListAttributes:output_type -> eolymp.community.ListAttributesOutput
	16, // [16:21] is the sub-list for method output_type
	11, // [11:16] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_eolymp_community_attribute_service_proto_init() }
func file_eolymp_community_attribute_service_proto_init() {
	if File_eolymp_community_attribute_service_proto != nil {
		return
	}
	file_eolymp_community_attribute_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_eolymp_community_attribute_service_proto_rawDesc), len(file_eolymp_community_attribute_service_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_eolymp_community_attribute_service_proto_goTypes,
		DependencyIndexes: file_eolymp_community_attribute_service_proto_depIdxs,
		MessageInfos:      file_eolymp_community_attribute_service_proto_msgTypes,
	}.Build()
	File_eolymp_community_attribute_service_proto = out.File
	file_eolymp_community_attribute_service_proto_goTypes = nil
	file_eolymp_community_attribute_service_proto_depIdxs = nil
}
