// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.18.1
// source: eolymp/community/group_service.proto

package community

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

type UpdateGroupInput_Patch int32

const (
	UpdateGroupInput_ALL         UpdateGroupInput_Patch = 0 // change all properties (same as an empty patch)
	UpdateGroupInput_NAME        UpdateGroupInput_Patch = 1 // change active/inactive flag
	UpdateGroupInput_DESCRIPTION UpdateGroupInput_Patch = 2 // change official/unofficial flag
)

// Enum value maps for UpdateGroupInput_Patch.
var (
	UpdateGroupInput_Patch_name = map[int32]string{
		0: "ALL",
		1: "NAME",
		2: "DESCRIPTION",
	}
	UpdateGroupInput_Patch_value = map[string]int32{
		"ALL":         0,
		"NAME":        1,
		"DESCRIPTION": 2,
	}
)

func (x UpdateGroupInput_Patch) Enum() *UpdateGroupInput_Patch {
	p := new(UpdateGroupInput_Patch)
	*p = x
	return p
}

func (x UpdateGroupInput_Patch) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (UpdateGroupInput_Patch) Descriptor() protoreflect.EnumDescriptor {
	return file_eolymp_community_group_service_proto_enumTypes[0].Descriptor()
}

func (UpdateGroupInput_Patch) Type() protoreflect.EnumType {
	return &file_eolymp_community_group_service_proto_enumTypes[0]
}

func (x UpdateGroupInput_Patch) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use UpdateGroupInput_Patch.Descriptor instead.
func (UpdateGroupInput_Patch) EnumDescriptor() ([]byte, []int) {
	return file_eolymp_community_group_service_proto_rawDescGZIP(), []int{2, 0}
}

type CreateGroupInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Group *Group `protobuf:"bytes,1,opt,name=group,proto3" json:"group,omitempty"`
}

func (x *CreateGroupInput) Reset() {
	*x = CreateGroupInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_community_group_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateGroupInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateGroupInput) ProtoMessage() {}

func (x *CreateGroupInput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_community_group_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateGroupInput.ProtoReflect.Descriptor instead.
func (*CreateGroupInput) Descriptor() ([]byte, []int) {
	return file_eolymp_community_group_service_proto_rawDescGZIP(), []int{0}
}

func (x *CreateGroupInput) GetGroup() *Group {
	if x != nil {
		return x.Group
	}
	return nil
}

type CreateGroupOutput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GroupId string `protobuf:"bytes,1,opt,name=group_id,json=groupId,proto3" json:"group_id,omitempty"`
}

func (x *CreateGroupOutput) Reset() {
	*x = CreateGroupOutput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_community_group_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateGroupOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateGroupOutput) ProtoMessage() {}

func (x *CreateGroupOutput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_community_group_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateGroupOutput.ProtoReflect.Descriptor instead.
func (*CreateGroupOutput) Descriptor() ([]byte, []int) {
	return file_eolymp_community_group_service_proto_rawDescGZIP(), []int{1}
}

func (x *CreateGroupOutput) GetGroupId() string {
	if x != nil {
		return x.GroupId
	}
	return ""
}

type UpdateGroupInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Patch   []UpdateGroupInput_Patch `protobuf:"varint,1,rep,packed,name=patch,proto3,enum=eolymp.community.UpdateGroupInput_Patch" json:"patch,omitempty"` // defines group props to be updated, empty means update everything
	GroupId string                   `protobuf:"bytes,2,opt,name=group_id,json=groupId,proto3" json:"group_id,omitempty"`
	Group   *Group                   `protobuf:"bytes,3,opt,name=group,proto3" json:"group,omitempty"`
}

func (x *UpdateGroupInput) Reset() {
	*x = UpdateGroupInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_community_group_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateGroupInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateGroupInput) ProtoMessage() {}

func (x *UpdateGroupInput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_community_group_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateGroupInput.ProtoReflect.Descriptor instead.
func (*UpdateGroupInput) Descriptor() ([]byte, []int) {
	return file_eolymp_community_group_service_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateGroupInput) GetPatch() []UpdateGroupInput_Patch {
	if x != nil {
		return x.Patch
	}
	return nil
}

func (x *UpdateGroupInput) GetGroupId() string {
	if x != nil {
		return x.GroupId
	}
	return ""
}

func (x *UpdateGroupInput) GetGroup() *Group {
	if x != nil {
		return x.Group
	}
	return nil
}

type UpdateGroupOutput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateGroupOutput) Reset() {
	*x = UpdateGroupOutput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_community_group_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateGroupOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateGroupOutput) ProtoMessage() {}

func (x *UpdateGroupOutput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_community_group_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateGroupOutput.ProtoReflect.Descriptor instead.
func (*UpdateGroupOutput) Descriptor() ([]byte, []int) {
	return file_eolymp_community_group_service_proto_rawDescGZIP(), []int{3}
}

type DeleteGroupInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GroupId string `protobuf:"bytes,1,opt,name=group_id,json=groupId,proto3" json:"group_id,omitempty"`
}

func (x *DeleteGroupInput) Reset() {
	*x = DeleteGroupInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_community_group_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteGroupInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteGroupInput) ProtoMessage() {}

func (x *DeleteGroupInput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_community_group_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteGroupInput.ProtoReflect.Descriptor instead.
func (*DeleteGroupInput) Descriptor() ([]byte, []int) {
	return file_eolymp_community_group_service_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteGroupInput) GetGroupId() string {
	if x != nil {
		return x.GroupId
	}
	return ""
}

type DeleteGroupOutput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteGroupOutput) Reset() {
	*x = DeleteGroupOutput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_community_group_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteGroupOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteGroupOutput) ProtoMessage() {}

func (x *DeleteGroupOutput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_community_group_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteGroupOutput.ProtoReflect.Descriptor instead.
func (*DeleteGroupOutput) Descriptor() ([]byte, []int) {
	return file_eolymp_community_group_service_proto_rawDescGZIP(), []int{5}
}

type DescribeGroupInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GroupId string `protobuf:"bytes,1,opt,name=group_id,json=groupId,proto3" json:"group_id,omitempty"`
}

func (x *DescribeGroupInput) Reset() {
	*x = DescribeGroupInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_community_group_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DescribeGroupInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescribeGroupInput) ProtoMessage() {}

func (x *DescribeGroupInput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_community_group_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DescribeGroupInput.ProtoReflect.Descriptor instead.
func (*DescribeGroupInput) Descriptor() ([]byte, []int) {
	return file_eolymp_community_group_service_proto_rawDescGZIP(), []int{6}
}

func (x *DescribeGroupInput) GetGroupId() string {
	if x != nil {
		return x.GroupId
	}
	return ""
}

type DescribeGroupOutput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Group *Group `protobuf:"bytes,1,opt,name=group,proto3" json:"group,omitempty"`
}

func (x *DescribeGroupOutput) Reset() {
	*x = DescribeGroupOutput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_community_group_service_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DescribeGroupOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescribeGroupOutput) ProtoMessage() {}

func (x *DescribeGroupOutput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_community_group_service_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DescribeGroupOutput.ProtoReflect.Descriptor instead.
func (*DescribeGroupOutput) Descriptor() ([]byte, []int) {
	return file_eolymp_community_group_service_proto_rawDescGZIP(), []int{7}
}

func (x *DescribeGroupOutput) GetGroup() *Group {
	if x != nil {
		return x.Group
	}
	return nil
}

type ListGroupsInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Offset  int32                   `protobuf:"varint,10,opt,name=offset,proto3" json:"offset,omitempty"`
	Size    int32                   `protobuf:"varint,11,opt,name=size,proto3" json:"size,omitempty"`
	Filters *ListGroupsInput_Filter `protobuf:"bytes,40,opt,name=filters,proto3" json:"filters,omitempty"`
}

func (x *ListGroupsInput) Reset() {
	*x = ListGroupsInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_community_group_service_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListGroupsInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListGroupsInput) ProtoMessage() {}

func (x *ListGroupsInput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_community_group_service_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListGroupsInput.ProtoReflect.Descriptor instead.
func (*ListGroupsInput) Descriptor() ([]byte, []int) {
	return file_eolymp_community_group_service_proto_rawDescGZIP(), []int{8}
}

func (x *ListGroupsInput) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *ListGroupsInput) GetSize() int32 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *ListGroupsInput) GetFilters() *ListGroupsInput_Filter {
	if x != nil {
		return x.Filters
	}
	return nil
}

type ListGroupsOutput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Total int32    `protobuf:"varint,1,opt,name=total,proto3" json:"total,omitempty"`
	Items []*Group `protobuf:"bytes,2,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *ListGroupsOutput) Reset() {
	*x = ListGroupsOutput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_community_group_service_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListGroupsOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListGroupsOutput) ProtoMessage() {}

func (x *ListGroupsOutput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_community_group_service_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListGroupsOutput.ProtoReflect.Descriptor instead.
func (*ListGroupsOutput) Descriptor() ([]byte, []int) {
	return file_eolymp_community_group_service_proto_rawDescGZIP(), []int{9}
}

func (x *ListGroupsOutput) GetTotal() int32 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *ListGroupsOutput) GetItems() []*Group {
	if x != nil {
		return x.Items
	}
	return nil
}

type ListGroupsInput_Filter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Query string                        `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
	Id    []*wellknown.ExpressionID     `protobuf:"bytes,2,rep,name=id,proto3" json:"id,omitempty"`
	Name  []*wellknown.ExpressionString `protobuf:"bytes,3,rep,name=name,proto3" json:"name,omitempty"`
}

func (x *ListGroupsInput_Filter) Reset() {
	*x = ListGroupsInput_Filter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_community_group_service_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListGroupsInput_Filter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListGroupsInput_Filter) ProtoMessage() {}

func (x *ListGroupsInput_Filter) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_community_group_service_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListGroupsInput_Filter.ProtoReflect.Descriptor instead.
func (*ListGroupsInput_Filter) Descriptor() ([]byte, []int) {
	return file_eolymp_community_group_service_proto_rawDescGZIP(), []int{8, 0}
}

func (x *ListGroupsInput_Filter) GetQuery() string {
	if x != nil {
		return x.Query
	}
	return ""
}

func (x *ListGroupsInput_Filter) GetId() []*wellknown.ExpressionID {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *ListGroupsInput_Filter) GetName() []*wellknown.ExpressionString {
	if x != nil {
		return x.Name
	}
	return nil
}

var File_eolymp_community_group_service_proto protoreflect.FileDescriptor

var file_eolymp_community_group_service_proto_rawDesc = []byte{
	0x0a, 0x24, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69,
	0x74, 0x79, 0x2f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x1a, 0x1d, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70,
	0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x68, 0x74, 0x74,
	0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x22, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x72, 0x61, 0x74, 0x65,
	0x6c, 0x69, 0x6d, 0x69, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x65, 0x6f, 0x6c,
	0x79, 0x6d, 0x70, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f,
	0x73, 0x63, 0x6f, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x65, 0x6f, 0x6c,
	0x79, 0x6d, 0x70, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x2f, 0x67, 0x72,
	0x6f, 0x75, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x21, 0x65, 0x6f, 0x6c, 0x79, 0x6d,
	0x70, 0x2f, 0x77, 0x65, 0x6c, 0x6c, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x2f, 0x65, 0x78, 0x70, 0x72,
	0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x41, 0x0a, 0x10,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x6e, 0x70, 0x75, 0x74,
	0x12, 0x2d, 0x0a, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x17, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69,
	0x74, 0x79, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x22,
	0x2e, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x4f, 0x75,
	0x74, 0x70, 0x75, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x22,
	0xc9, 0x01, 0x0a, 0x10, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49,
	0x6e, 0x70, 0x75, 0x74, 0x12, 0x3e, 0x0a, 0x05, 0x70, 0x61, 0x74, 0x63, 0x68, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0e, 0x32, 0x28, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x2e, 0x50, 0x61, 0x74, 0x63, 0x68, 0x52, 0x05, 0x70,
	0x61, 0x74, 0x63, 0x68, 0x12, 0x19, 0x0a, 0x08, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x12,
	0x2d, 0x0a, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17,
	0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74,
	0x79, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x22, 0x2b,
	0x0a, 0x05, 0x50, 0x61, 0x74, 0x63, 0x68, 0x12, 0x07, 0x0a, 0x03, 0x41, 0x4c, 0x4c, 0x10, 0x00,
	0x12, 0x08, 0x0a, 0x04, 0x4e, 0x41, 0x4d, 0x45, 0x10, 0x01, 0x12, 0x0f, 0x0a, 0x0b, 0x44, 0x45,
	0x53, 0x43, 0x52, 0x49, 0x50, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x02, 0x22, 0x13, 0x0a, 0x11, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74,
	0x22, 0x2d, 0x0a, 0x10, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49,
	0x6e, 0x70, 0x75, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x22,
	0x13, 0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x4f, 0x75,
	0x74, 0x70, 0x75, 0x74, 0x22, 0x2f, 0x0a, 0x12, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x67, 0x72,
	0x6f, 0x75, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x67, 0x72,
	0x6f, 0x75, 0x70, 0x49, 0x64, 0x22, 0x44, 0x0a, 0x13, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62,
	0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x12, 0x2d, 0x0a, 0x05,
	0x67, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x65, 0x6f,
	0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x2e, 0x47,
	0x72, 0x6f, 0x75, 0x70, 0x52, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x22, 0x8a, 0x02, 0x0a, 0x0f,
	0x4c, 0x69, 0x73, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12,
	0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18,
	0x0b, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x42, 0x0a, 0x07, 0x66,
	0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x18, 0x28, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x65,
	0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x2e,
	0x4c, 0x69, 0x73, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x2e,
	0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x52, 0x07, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x1a,
	0x86, 0x01, 0x0a, 0x06, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x71, 0x75,
	0x65, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79,
	0x12, 0x2e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x65,
	0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x77, 0x65, 0x6c, 0x6c, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x2e,
	0x45, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x36, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22,
	0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x77, 0x65, 0x6c, 0x6c, 0x6b, 0x6e, 0x6f, 0x77,
	0x6e, 0x2e, 0x45, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x72, 0x69,
	0x6e, 0x67, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x57, 0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x12, 0x14, 0x0a, 0x05,
	0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x74, 0x6f, 0x74,
	0x61, 0x6c, 0x12, 0x2d, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x17, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x75,
	0x6e, 0x69, 0x74, 0x79, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d,
	0x73, 0x32, 0x9e, 0x06, 0x0a, 0x0c, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x93, 0x01, 0x0a, 0x0b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x12, 0x22, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x75, 0x6e, 0x69, 0x74, 0x79, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x23, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x22, 0x3b, 0xea, 0xe2, 0x0a,
	0x0b, 0xf5, 0xe2, 0x0a, 0x00, 0x00, 0xa0, 0x40, 0xf8, 0xe2, 0x0a, 0x14, 0x82, 0xe3, 0x0a, 0x19,
	0x8a, 0xe3, 0x0a, 0x15, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x3a, 0x67, 0x72,
	0x6f, 0x75, 0x70, 0x3a, 0x77, 0x72, 0x69, 0x74, 0x65, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x09, 0x22,
	0x07, 0x2f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x12, 0x9e, 0x01, 0x0a, 0x0b, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x22, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d,
	0x70, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x2e, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x23, 0x2e, 0x65,
	0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x2e,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x4f, 0x75, 0x74, 0x70, 0x75,
	0x74, 0x22, 0x46, 0xea, 0xe2, 0x0a, 0x0b, 0xf5, 0xe2, 0x0a, 0x00, 0x00, 0xa0, 0x40, 0xf8, 0xe2,
	0x0a, 0x14, 0x82, 0xe3, 0x0a, 0x19, 0x8a, 0xe3, 0x0a, 0x15, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e,
	0x69, 0x74, 0x79, 0x3a, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x3a, 0x77, 0x72, 0x69, 0x74, 0x65, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x14, 0x22, 0x12, 0x2f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x2f, 0x7b,
	0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x69, 0x64, 0x7d, 0x12, 0x9e, 0x01, 0x0a, 0x0b, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x22, 0x2e, 0x65, 0x6f, 0x6c, 0x79,
	0x6d, 0x70, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x2e, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x23, 0x2e,
	0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79,
	0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x4f, 0x75, 0x74, 0x70,
	0x75, 0x74, 0x22, 0x46, 0xea, 0xe2, 0x0a, 0x0b, 0xf5, 0xe2, 0x0a, 0x00, 0x00, 0xa0, 0x40, 0xf8,
	0xe2, 0x0a, 0x14, 0x82, 0xe3, 0x0a, 0x19, 0x8a, 0xe3, 0x0a, 0x15, 0x63, 0x6f, 0x6d, 0x6d, 0x75,
	0x6e, 0x69, 0x74, 0x79, 0x3a, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x3a, 0x77, 0x72, 0x69, 0x74, 0x65,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x14, 0x2a, 0x12, 0x2f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x2f,
	0x7b, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x69, 0x64, 0x7d, 0x12, 0xa3, 0x01, 0x0a, 0x0d, 0x44,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x24, 0x2e, 0x65,
	0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x2e,
	0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x6e, 0x70,
	0x75, 0x74, 0x1a, 0x25, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x75, 0x6e, 0x69, 0x74, 0x79, 0x2e, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x22, 0x45, 0xea, 0xe2, 0x0a, 0x0b, 0xf5,
	0xe2, 0x0a, 0x00, 0x00, 0xa0, 0x40, 0xf8, 0xe2, 0x0a, 0x14, 0x82, 0xe3, 0x0a, 0x18, 0x8a, 0xe3,
	0x0a, 0x14, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x3a, 0x67, 0x72, 0x6f, 0x75,
	0x70, 0x3a, 0x72, 0x65, 0x61, 0x64, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x14, 0x12, 0x12, 0x2f, 0x67,
	0x72, 0x6f, 0x75, 0x70, 0x73, 0x2f, 0x7b, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x69, 0x64, 0x7d,
	0x12, 0x8f, 0x01, 0x0a, 0x0a, 0x4c, 0x69, 0x73, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x12,
	0x21, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69,
	0x74, 0x79, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x49, 0x6e, 0x70,
	0x75, 0x74, 0x1a, 0x22, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x75, 0x6e, 0x69, 0x74, 0x79, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73,
	0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x22, 0x3a, 0xea, 0xe2, 0x0a, 0x0b, 0xf5, 0xe2, 0x0a, 0x00,
	0x00, 0xa0, 0x40, 0xf8, 0xe2, 0x0a, 0x14, 0x82, 0xe3, 0x0a, 0x18, 0x8a, 0xe3, 0x0a, 0x14, 0x63,
	0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x3a, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x3a, 0x72,
	0x65, 0x61, 0x64, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x09, 0x12, 0x07, 0x2f, 0x67, 0x72, 0x6f, 0x75,
	0x70, 0x73, 0x42, 0x35, 0x5a, 0x33, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x67, 0x6f, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x65,
	0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x3b,
	0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_eolymp_community_group_service_proto_rawDescOnce sync.Once
	file_eolymp_community_group_service_proto_rawDescData = file_eolymp_community_group_service_proto_rawDesc
)

func file_eolymp_community_group_service_proto_rawDescGZIP() []byte {
	file_eolymp_community_group_service_proto_rawDescOnce.Do(func() {
		file_eolymp_community_group_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_eolymp_community_group_service_proto_rawDescData)
	})
	return file_eolymp_community_group_service_proto_rawDescData
}

var file_eolymp_community_group_service_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_eolymp_community_group_service_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_eolymp_community_group_service_proto_goTypes = []interface{}{
	(UpdateGroupInput_Patch)(0),        // 0: eolymp.community.UpdateGroupInput.Patch
	(*CreateGroupInput)(nil),           // 1: eolymp.community.CreateGroupInput
	(*CreateGroupOutput)(nil),          // 2: eolymp.community.CreateGroupOutput
	(*UpdateGroupInput)(nil),           // 3: eolymp.community.UpdateGroupInput
	(*UpdateGroupOutput)(nil),          // 4: eolymp.community.UpdateGroupOutput
	(*DeleteGroupInput)(nil),           // 5: eolymp.community.DeleteGroupInput
	(*DeleteGroupOutput)(nil),          // 6: eolymp.community.DeleteGroupOutput
	(*DescribeGroupInput)(nil),         // 7: eolymp.community.DescribeGroupInput
	(*DescribeGroupOutput)(nil),        // 8: eolymp.community.DescribeGroupOutput
	(*ListGroupsInput)(nil),            // 9: eolymp.community.ListGroupsInput
	(*ListGroupsOutput)(nil),           // 10: eolymp.community.ListGroupsOutput
	(*ListGroupsInput_Filter)(nil),     // 11: eolymp.community.ListGroupsInput.Filter
	(*Group)(nil),                      // 12: eolymp.community.Group
	(*wellknown.ExpressionID)(nil),     // 13: eolymp.wellknown.ExpressionID
	(*wellknown.ExpressionString)(nil), // 14: eolymp.wellknown.ExpressionString
}
var file_eolymp_community_group_service_proto_depIdxs = []int32{
	12, // 0: eolymp.community.CreateGroupInput.group:type_name -> eolymp.community.Group
	0,  // 1: eolymp.community.UpdateGroupInput.patch:type_name -> eolymp.community.UpdateGroupInput.Patch
	12, // 2: eolymp.community.UpdateGroupInput.group:type_name -> eolymp.community.Group
	12, // 3: eolymp.community.DescribeGroupOutput.group:type_name -> eolymp.community.Group
	11, // 4: eolymp.community.ListGroupsInput.filters:type_name -> eolymp.community.ListGroupsInput.Filter
	12, // 5: eolymp.community.ListGroupsOutput.items:type_name -> eolymp.community.Group
	13, // 6: eolymp.community.ListGroupsInput.Filter.id:type_name -> eolymp.wellknown.ExpressionID
	14, // 7: eolymp.community.ListGroupsInput.Filter.name:type_name -> eolymp.wellknown.ExpressionString
	1,  // 8: eolymp.community.GroupService.CreateGroup:input_type -> eolymp.community.CreateGroupInput
	3,  // 9: eolymp.community.GroupService.UpdateGroup:input_type -> eolymp.community.UpdateGroupInput
	5,  // 10: eolymp.community.GroupService.DeleteGroup:input_type -> eolymp.community.DeleteGroupInput
	7,  // 11: eolymp.community.GroupService.DescribeGroup:input_type -> eolymp.community.DescribeGroupInput
	9,  // 12: eolymp.community.GroupService.ListGroups:input_type -> eolymp.community.ListGroupsInput
	2,  // 13: eolymp.community.GroupService.CreateGroup:output_type -> eolymp.community.CreateGroupOutput
	4,  // 14: eolymp.community.GroupService.UpdateGroup:output_type -> eolymp.community.UpdateGroupOutput
	6,  // 15: eolymp.community.GroupService.DeleteGroup:output_type -> eolymp.community.DeleteGroupOutput
	8,  // 16: eolymp.community.GroupService.DescribeGroup:output_type -> eolymp.community.DescribeGroupOutput
	10, // 17: eolymp.community.GroupService.ListGroups:output_type -> eolymp.community.ListGroupsOutput
	13, // [13:18] is the sub-list for method output_type
	8,  // [8:13] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_eolymp_community_group_service_proto_init() }
func file_eolymp_community_group_service_proto_init() {
	if File_eolymp_community_group_service_proto != nil {
		return
	}
	file_eolymp_community_group_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_eolymp_community_group_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateGroupInput); i {
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
		file_eolymp_community_group_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateGroupOutput); i {
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
		file_eolymp_community_group_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateGroupInput); i {
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
		file_eolymp_community_group_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateGroupOutput); i {
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
		file_eolymp_community_group_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteGroupInput); i {
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
		file_eolymp_community_group_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteGroupOutput); i {
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
		file_eolymp_community_group_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DescribeGroupInput); i {
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
		file_eolymp_community_group_service_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DescribeGroupOutput); i {
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
		file_eolymp_community_group_service_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListGroupsInput); i {
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
		file_eolymp_community_group_service_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListGroupsOutput); i {
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
		file_eolymp_community_group_service_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListGroupsInput_Filter); i {
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
			RawDescriptor: file_eolymp_community_group_service_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_eolymp_community_group_service_proto_goTypes,
		DependencyIndexes: file_eolymp_community_group_service_proto_depIdxs,
		EnumInfos:         file_eolymp_community_group_service_proto_enumTypes,
		MessageInfos:      file_eolymp_community_group_service_proto_msgTypes,
	}.Build()
	File_eolymp_community_group_service_proto = out.File
	file_eolymp_community_group_service_proto_rawDesc = nil
	file_eolymp_community_group_service_proto_goTypes = nil
	file_eolymp_community_group_service_proto_depIdxs = nil
}
