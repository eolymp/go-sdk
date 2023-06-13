// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.18.1
// source: eolymp/discussion/forum_service.proto

package atlas

import (
	_ "github.com/eolymp/go-sdk/eolymp/annotations"
	discussion "github.com/eolymp/go-sdk/eolymp/discussion"
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

type DescribeForumInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ForumId string `protobuf:"bytes,1,opt,name=forum_id,json=forumId,proto3" json:"forum_id,omitempty"`
	Render  bool   `protobuf:"varint,2,opt,name=render,proto3" json:"render,omitempty"`
}

func (x *DescribeForumInput) Reset() {
	*x = DescribeForumInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_discussion_forum_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DescribeForumInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescribeForumInput) ProtoMessage() {}

func (x *DescribeForumInput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_discussion_forum_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DescribeForumInput.ProtoReflect.Descriptor instead.
func (*DescribeForumInput) Descriptor() ([]byte, []int) {
	return file_eolymp_discussion_forum_service_proto_rawDescGZIP(), []int{0}
}

func (x *DescribeForumInput) GetForumId() string {
	if x != nil {
		return x.ForumId
	}
	return ""
}

func (x *DescribeForumInput) GetRender() bool {
	if x != nil {
		return x.Render
	}
	return false
}

type DescribeForumOutput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Forum *discussion.Forum `protobuf:"bytes,1,opt,name=forum,proto3" json:"forum,omitempty"`
}

func (x *DescribeForumOutput) Reset() {
	*x = DescribeForumOutput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_discussion_forum_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DescribeForumOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescribeForumOutput) ProtoMessage() {}

func (x *DescribeForumOutput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_discussion_forum_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DescribeForumOutput.ProtoReflect.Descriptor instead.
func (*DescribeForumOutput) Descriptor() ([]byte, []int) {
	return file_eolymp_discussion_forum_service_proto_rawDescGZIP(), []int{1}
}

func (x *DescribeForumOutput) GetForum() *discussion.Forum {
	if x != nil {
		return x.Forum
	}
	return nil
}

type ListForumsInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Render bool `protobuf:"varint,1,opt,name=render,proto3" json:"render,omitempty"`
	// pagination
	Offset int32 `protobuf:"varint,10,opt,name=offset,proto3" json:"offset,omitempty"`
	Size   int32 `protobuf:"varint,11,opt,name=size,proto3" json:"size,omitempty"`
	// data filters
	Filters *ListForumsInput_Filter `protobuf:"bytes,40,opt,name=filters,proto3" json:"filters,omitempty"`
}

func (x *ListForumsInput) Reset() {
	*x = ListForumsInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_discussion_forum_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListForumsInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListForumsInput) ProtoMessage() {}

func (x *ListForumsInput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_discussion_forum_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListForumsInput.ProtoReflect.Descriptor instead.
func (*ListForumsInput) Descriptor() ([]byte, []int) {
	return file_eolymp_discussion_forum_service_proto_rawDescGZIP(), []int{2}
}

func (x *ListForumsInput) GetRender() bool {
	if x != nil {
		return x.Render
	}
	return false
}

func (x *ListForumsInput) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *ListForumsInput) GetSize() int32 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *ListForumsInput) GetFilters() *ListForumsInput_Filter {
	if x != nil {
		return x.Filters
	}
	return nil
}

type ListForumsOutput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Total int32               `protobuf:"varint,1,opt,name=total,proto3" json:"total,omitempty"`
	Items []*discussion.Forum `protobuf:"bytes,2,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *ListForumsOutput) Reset() {
	*x = ListForumsOutput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_discussion_forum_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListForumsOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListForumsOutput) ProtoMessage() {}

func (x *ListForumsOutput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_discussion_forum_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListForumsOutput.ProtoReflect.Descriptor instead.
func (*ListForumsOutput) Descriptor() ([]byte, []int) {
	return file_eolymp_discussion_forum_service_proto_rawDescGZIP(), []int{3}
}

func (x *ListForumsOutput) GetTotal() int32 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *ListForumsOutput) GetItems() []*discussion.Forum {
	if x != nil {
		return x.Items
	}
	return nil
}

type CreateForumInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Forum *discussion.Forum `protobuf:"bytes,1,opt,name=forum,proto3" json:"forum,omitempty"`
}

func (x *CreateForumInput) Reset() {
	*x = CreateForumInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_discussion_forum_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateForumInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateForumInput) ProtoMessage() {}

func (x *CreateForumInput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_discussion_forum_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateForumInput.ProtoReflect.Descriptor instead.
func (*CreateForumInput) Descriptor() ([]byte, []int) {
	return file_eolymp_discussion_forum_service_proto_rawDescGZIP(), []int{4}
}

func (x *CreateForumInput) GetForum() *discussion.Forum {
	if x != nil {
		return x.Forum
	}
	return nil
}

type CreateForumOutput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ForumId string `protobuf:"bytes,1,opt,name=forum_id,json=forumId,proto3" json:"forum_id,omitempty"`
}

func (x *CreateForumOutput) Reset() {
	*x = CreateForumOutput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_discussion_forum_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateForumOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateForumOutput) ProtoMessage() {}

func (x *CreateForumOutput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_discussion_forum_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateForumOutput.ProtoReflect.Descriptor instead.
func (*CreateForumOutput) Descriptor() ([]byte, []int) {
	return file_eolymp_discussion_forum_service_proto_rawDescGZIP(), []int{5}
}

func (x *CreateForumOutput) GetForumId() string {
	if x != nil {
		return x.ForumId
	}
	return ""
}

type UpdateForumInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ForumId string            `protobuf:"bytes,1,opt,name=forum_id,json=forumId,proto3" json:"forum_id,omitempty"`
	Forum   *discussion.Forum `protobuf:"bytes,2,opt,name=forum,proto3" json:"forum,omitempty"`
}

func (x *UpdateForumInput) Reset() {
	*x = UpdateForumInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_discussion_forum_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateForumInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateForumInput) ProtoMessage() {}

func (x *UpdateForumInput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_discussion_forum_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateForumInput.ProtoReflect.Descriptor instead.
func (*UpdateForumInput) Descriptor() ([]byte, []int) {
	return file_eolymp_discussion_forum_service_proto_rawDescGZIP(), []int{6}
}

func (x *UpdateForumInput) GetForumId() string {
	if x != nil {
		return x.ForumId
	}
	return ""
}

func (x *UpdateForumInput) GetForum() *discussion.Forum {
	if x != nil {
		return x.Forum
	}
	return nil
}

type UpdateForumOutput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateForumOutput) Reset() {
	*x = UpdateForumOutput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_discussion_forum_service_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateForumOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateForumOutput) ProtoMessage() {}

func (x *UpdateForumOutput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_discussion_forum_service_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateForumOutput.ProtoReflect.Descriptor instead.
func (*UpdateForumOutput) Descriptor() ([]byte, []int) {
	return file_eolymp_discussion_forum_service_proto_rawDescGZIP(), []int{7}
}

type DeleteForumInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ForumId string `protobuf:"bytes,1,opt,name=forum_id,json=forumId,proto3" json:"forum_id,omitempty"`
}

func (x *DeleteForumInput) Reset() {
	*x = DeleteForumInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_discussion_forum_service_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteForumInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteForumInput) ProtoMessage() {}

func (x *DeleteForumInput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_discussion_forum_service_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteForumInput.ProtoReflect.Descriptor instead.
func (*DeleteForumInput) Descriptor() ([]byte, []int) {
	return file_eolymp_discussion_forum_service_proto_rawDescGZIP(), []int{8}
}

func (x *DeleteForumInput) GetForumId() string {
	if x != nil {
		return x.ForumId
	}
	return ""
}

type DeleteForumOutput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteForumOutput) Reset() {
	*x = DeleteForumOutput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_discussion_forum_service_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteForumOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteForumOutput) ProtoMessage() {}

func (x *DeleteForumOutput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_discussion_forum_service_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteForumOutput.ProtoReflect.Descriptor instead.
func (*DeleteForumOutput) Descriptor() ([]byte, []int) {
	return file_eolymp_discussion_forum_service_proto_rawDescGZIP(), []int{9}
}

type ListForumsInput_Filter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Query    string                    `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
	Id       []*wellknown.ExpressionID `protobuf:"bytes,2,rep,name=id,proto3" json:"id,omitempty"`
	ParentId []*wellknown.ExpressionID `protobuf:"bytes,3,rep,name=parent_id,json=parentId,proto3" json:"parent_id,omitempty"`
}

func (x *ListForumsInput_Filter) Reset() {
	*x = ListForumsInput_Filter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_discussion_forum_service_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListForumsInput_Filter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListForumsInput_Filter) ProtoMessage() {}

func (x *ListForumsInput_Filter) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_discussion_forum_service_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListForumsInput_Filter.ProtoReflect.Descriptor instead.
func (*ListForumsInput_Filter) Descriptor() ([]byte, []int) {
	return file_eolymp_discussion_forum_service_proto_rawDescGZIP(), []int{2, 0}
}

func (x *ListForumsInput_Filter) GetQuery() string {
	if x != nil {
		return x.Query
	}
	return ""
}

func (x *ListForumsInput_Filter) GetId() []*wellknown.ExpressionID {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *ListForumsInput_Filter) GetParentId() []*wellknown.ExpressionID {
	if x != nil {
		return x.ParentId
	}
	return nil
}

var File_eolymp_discussion_forum_service_proto protoreflect.FileDescriptor

var file_eolymp_discussion_forum_service_proto_rawDesc = []byte{
	0x0a, 0x25, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x64, 0x69, 0x73, 0x63, 0x75, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x2f, 0x66, 0x6f, 0x72, 0x75, 0x6d, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e,
	0x64, 0x69, 0x73, 0x63, 0x75, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x1a, 0x1d, 0x65, 0x6f, 0x6c, 0x79,
	0x6d, 0x70, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x68,
	0x74, 0x74, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x22, 0x65, 0x6f, 0x6c, 0x79, 0x6d,
	0x70, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x72, 0x61,
	0x74, 0x65, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x65,
	0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2f, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x65,
	0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x64, 0x69, 0x73, 0x63, 0x75, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x2f, 0x66, 0x6f, 0x72, 0x75, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x21, 0x65, 0x6f,
	0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x77, 0x65, 0x6c, 0x6c, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x2f, 0x65,
	0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x47, 0x0a, 0x12, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x46, 0x6f, 0x72, 0x75, 0x6d,
	0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x66, 0x6f, 0x72, 0x75, 0x6d, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x66, 0x6f, 0x72, 0x75, 0x6d, 0x49, 0x64,
	0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x06, 0x72, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x22, 0x45, 0x0a, 0x13, 0x44, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x62, 0x65, 0x46, 0x6f, 0x72, 0x75, 0x6d, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x12,
	0x2e, 0x0a, 0x05, 0x66, 0x6f, 0x72, 0x75, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18,
	0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x75, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x2e, 0x46, 0x6f, 0x72, 0x75, 0x6d, 0x52, 0x05, 0x66, 0x6f, 0x72, 0x75, 0x6d, 0x22,
	0xa8, 0x02, 0x0a, 0x0f, 0x4c, 0x69, 0x73, 0x74, 0x46, 0x6f, 0x72, 0x75, 0x6d, 0x73, 0x49, 0x6e,
	0x70, 0x75, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x06, 0x72, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x6f,
	0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6f, 0x66, 0x66,
	0x73, 0x65, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x43, 0x0a, 0x07, 0x66, 0x69, 0x6c, 0x74, 0x65,
	0x72, 0x73, 0x18, 0x28, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d,
	0x70, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x75, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x46, 0x6f, 0x72, 0x75, 0x6d, 0x73, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x2e, 0x46, 0x69, 0x6c,
	0x74, 0x65, 0x72, 0x52, 0x07, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x1a, 0x8b, 0x01, 0x0a,
	0x06, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x12, 0x2e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x65, 0x6f, 0x6c, 0x79,
	0x6d, 0x70, 0x2e, 0x77, 0x65, 0x6c, 0x6c, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x2e, 0x45, 0x78, 0x70,
	0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x52, 0x02, 0x69, 0x64, 0x12, 0x3b, 0x0a,
	0x09, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x1e, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x77, 0x65, 0x6c, 0x6c, 0x6b, 0x6e,
	0x6f, 0x77, 0x6e, 0x2e, 0x45, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x44,
	0x52, 0x08, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x22, 0x58, 0x0a, 0x10, 0x4c, 0x69,
	0x73, 0x74, 0x46, 0x6f, 0x72, 0x75, 0x6d, 0x73, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x12, 0x14,
	0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x74,
	0x6f, 0x74, 0x61, 0x6c, 0x12, 0x2e, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x64, 0x69, 0x73,
	0x63, 0x75, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x46, 0x6f, 0x72, 0x75, 0x6d, 0x52, 0x05, 0x69,
	0x74, 0x65, 0x6d, 0x73, 0x22, 0x42, 0x0a, 0x10, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x46, 0x6f,
	0x72, 0x75, 0x6d, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x2e, 0x0a, 0x05, 0x66, 0x6f, 0x72, 0x75,
	0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70,
	0x2e, 0x64, 0x69, 0x73, 0x63, 0x75, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x46, 0x6f, 0x72, 0x75,
	0x6d, 0x52, 0x05, 0x66, 0x6f, 0x72, 0x75, 0x6d, 0x22, 0x2e, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x46, 0x6f, 0x72, 0x75, 0x6d, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x12, 0x19, 0x0a,
	0x08, 0x66, 0x6f, 0x72, 0x75, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x66, 0x6f, 0x72, 0x75, 0x6d, 0x49, 0x64, 0x22, 0x5d, 0x0a, 0x10, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x46, 0x6f, 0x72, 0x75, 0x6d, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x19, 0x0a, 0x08,
	0x66, 0x6f, 0x72, 0x75, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x66, 0x6f, 0x72, 0x75, 0x6d, 0x49, 0x64, 0x12, 0x2e, 0x0a, 0x05, 0x66, 0x6f, 0x72, 0x75, 0x6d,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e,
	0x64, 0x69, 0x73, 0x63, 0x75, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x46, 0x6f, 0x72, 0x75, 0x6d,
	0x52, 0x05, 0x66, 0x6f, 0x72, 0x75, 0x6d, 0x22, 0x13, 0x0a, 0x11, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x46, 0x6f, 0x72, 0x75, 0x6d, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x22, 0x2d, 0x0a, 0x10,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x46, 0x6f, 0x72, 0x75, 0x6d, 0x49, 0x6e, 0x70, 0x75, 0x74,
	0x12, 0x19, 0x0a, 0x08, 0x66, 0x6f, 0x72, 0x75, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x66, 0x6f, 0x72, 0x75, 0x6d, 0x49, 0x64, 0x22, 0x13, 0x0a, 0x11, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x46, 0x6f, 0x72, 0x75, 0x6d, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74,
	0x32, 0xae, 0x06, 0x0a, 0x0c, 0x46, 0x6f, 0x72, 0x75, 0x6d, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0xa7, 0x01, 0x0a, 0x0d, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x46, 0x6f,
	0x72, 0x75, 0x6d, 0x12, 0x25, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x64, 0x69, 0x73,
	0x63, 0x75, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65,
	0x46, 0x6f, 0x72, 0x75, 0x6d, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x26, 0x2e, 0x65, 0x6f, 0x6c,
	0x79, 0x6d, 0x70, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x75, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x44,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x46, 0x6f, 0x72, 0x75, 0x6d, 0x4f, 0x75, 0x74, 0x70,
	0x75, 0x74, 0x22, 0x47, 0xea, 0xe2, 0x0a, 0x0c, 0xf5, 0xe2, 0x0a, 0x00, 0x00, 0xa0, 0x41, 0xf8,
	0xe2, 0x0a, 0xf4, 0x03, 0x82, 0xe3, 0x0a, 0x19, 0x8a, 0xe3, 0x0a, 0x15, 0x64, 0x69, 0x73, 0x63,
	0x75, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x3a, 0x66, 0x6f, 0x72, 0x75, 0x6d, 0x3a, 0x72, 0x65, 0x61,
	0x64, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x14, 0x12, 0x12, 0x2f, 0x66, 0x6f, 0x72, 0x75, 0x6d, 0x73,
	0x2f, 0x7b, 0x66, 0x6f, 0x72, 0x75, 0x6d, 0x5f, 0x69, 0x64, 0x7d, 0x12, 0x92, 0x01, 0x0a, 0x0a,
	0x4c, 0x69, 0x73, 0x74, 0x46, 0x6f, 0x72, 0x75, 0x6d, 0x73, 0x12, 0x22, 0x2e, 0x65, 0x6f, 0x6c,
	0x79, 0x6d, 0x70, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x75, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x46, 0x6f, 0x72, 0x75, 0x6d, 0x73, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x23,
	0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x75, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x46, 0x6f, 0x72, 0x75, 0x6d, 0x73, 0x4f, 0x75, 0x74,
	0x70, 0x75, 0x74, 0x22, 0x3b, 0xea, 0xe2, 0x0a, 0x0b, 0xf5, 0xe2, 0x0a, 0x00, 0x00, 0xa0, 0x41,
	0xf8, 0xe2, 0x0a, 0x64, 0x82, 0xe3, 0x0a, 0x19, 0x8a, 0xe3, 0x0a, 0x15, 0x64, 0x69, 0x73, 0x63,
	0x75, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x3a, 0x66, 0x6f, 0x72, 0x75, 0x6d, 0x3a, 0x72, 0x65, 0x61,
	0x64, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x09, 0x12, 0x07, 0x2f, 0x66, 0x6f, 0x72, 0x75, 0x6d, 0x73,
	0x12, 0x96, 0x01, 0x0a, 0x0b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x46, 0x6f, 0x72, 0x75, 0x6d,
	0x12, 0x23, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x75, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x46, 0x6f, 0x72, 0x75, 0x6d,
	0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x24, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x64,
	0x69, 0x73, 0x63, 0x75, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x46, 0x6f, 0x72, 0x75, 0x6d, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x22, 0x3c, 0xea, 0xe2, 0x0a,
	0x0b, 0xf5, 0xe2, 0x0a, 0x00, 0x00, 0xa0, 0x40, 0xf8, 0xe2, 0x0a, 0x32, 0x82, 0xe3, 0x0a, 0x1a,
	0x8a, 0xe3, 0x0a, 0x16, 0x64, 0x69, 0x73, 0x63, 0x75, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x3a, 0x66,
	0x6f, 0x72, 0x75, 0x6d, 0x3a, 0x77, 0x72, 0x69, 0x74, 0x65, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x09,
	0x22, 0x07, 0x2f, 0x66, 0x6f, 0x72, 0x75, 0x6d, 0x73, 0x12, 0xa1, 0x01, 0x0a, 0x0b, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x46, 0x6f, 0x72, 0x75, 0x6d, 0x12, 0x23, 0x2e, 0x65, 0x6f, 0x6c, 0x79,
	0x6d, 0x70, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x75, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x46, 0x6f, 0x72, 0x75, 0x6d, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x24,
	0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x75, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x46, 0x6f, 0x72, 0x75, 0x6d, 0x4f, 0x75,
	0x74, 0x70, 0x75, 0x74, 0x22, 0x47, 0xea, 0xe2, 0x0a, 0x0b, 0xf5, 0xe2, 0x0a, 0x00, 0x00, 0xa0,
	0x40, 0xf8, 0xe2, 0x0a, 0x32, 0x82, 0xe3, 0x0a, 0x1a, 0x8a, 0xe3, 0x0a, 0x16, 0x64, 0x69, 0x73,
	0x63, 0x75, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x3a, 0x66, 0x6f, 0x72, 0x75, 0x6d, 0x3a, 0x77, 0x72,
	0x69, 0x74, 0x65, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x14, 0x1a, 0x12, 0x2f, 0x66, 0x6f, 0x72, 0x75,
	0x6d, 0x73, 0x2f, 0x7b, 0x66, 0x6f, 0x72, 0x75, 0x6d, 0x5f, 0x69, 0x64, 0x7d, 0x12, 0xa1, 0x01,
	0x0a, 0x0b, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x46, 0x6f, 0x72, 0x75, 0x6d, 0x12, 0x23, 0x2e,
	0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x75, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x46, 0x6f, 0x72, 0x75, 0x6d, 0x49, 0x6e, 0x70,
	0x75, 0x74, 0x1a, 0x24, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x64, 0x69, 0x73, 0x63,
	0x75, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x46, 0x6f, 0x72,
	0x75, 0x6d, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x22, 0x47, 0xea, 0xe2, 0x0a, 0x0b, 0xf5, 0xe2,
	0x0a, 0x00, 0x00, 0xa0, 0x40, 0xf8, 0xe2, 0x0a, 0x32, 0x82, 0xe3, 0x0a, 0x1a, 0x8a, 0xe3, 0x0a,
	0x16, 0x64, 0x69, 0x73, 0x63, 0x75, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x3a, 0x66, 0x6f, 0x72, 0x75,
	0x6d, 0x3a, 0x77, 0x72, 0x69, 0x74, 0x65, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x14, 0x2a, 0x12, 0x2f,
	0x66, 0x6f, 0x72, 0x75, 0x6d, 0x73, 0x2f, 0x7b, 0x66, 0x6f, 0x72, 0x75, 0x6d, 0x5f, 0x69, 0x64,
	0x7d, 0x42, 0x2d, 0x5a, 0x2b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x67, 0x6f, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x65, 0x6f,
	0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x61, 0x74, 0x6c, 0x61, 0x73, 0x3b, 0x61, 0x74, 0x6c, 0x61, 0x73,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_eolymp_discussion_forum_service_proto_rawDescOnce sync.Once
	file_eolymp_discussion_forum_service_proto_rawDescData = file_eolymp_discussion_forum_service_proto_rawDesc
)

func file_eolymp_discussion_forum_service_proto_rawDescGZIP() []byte {
	file_eolymp_discussion_forum_service_proto_rawDescOnce.Do(func() {
		file_eolymp_discussion_forum_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_eolymp_discussion_forum_service_proto_rawDescData)
	})
	return file_eolymp_discussion_forum_service_proto_rawDescData
}

var file_eolymp_discussion_forum_service_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_eolymp_discussion_forum_service_proto_goTypes = []interface{}{
	(*DescribeForumInput)(nil),     // 0: eolymp.discussion.DescribeForumInput
	(*DescribeForumOutput)(nil),    // 1: eolymp.discussion.DescribeForumOutput
	(*ListForumsInput)(nil),        // 2: eolymp.discussion.ListForumsInput
	(*ListForumsOutput)(nil),       // 3: eolymp.discussion.ListForumsOutput
	(*CreateForumInput)(nil),       // 4: eolymp.discussion.CreateForumInput
	(*CreateForumOutput)(nil),      // 5: eolymp.discussion.CreateForumOutput
	(*UpdateForumInput)(nil),       // 6: eolymp.discussion.UpdateForumInput
	(*UpdateForumOutput)(nil),      // 7: eolymp.discussion.UpdateForumOutput
	(*DeleteForumInput)(nil),       // 8: eolymp.discussion.DeleteForumInput
	(*DeleteForumOutput)(nil),      // 9: eolymp.discussion.DeleteForumOutput
	(*ListForumsInput_Filter)(nil), // 10: eolymp.discussion.ListForumsInput.Filter
	(*discussion.Forum)(nil),       // 11: eolymp.discussion.Forum
	(*wellknown.ExpressionID)(nil), // 12: eolymp.wellknown.ExpressionID
}
var file_eolymp_discussion_forum_service_proto_depIdxs = []int32{
	11, // 0: eolymp.discussion.DescribeForumOutput.forum:type_name -> eolymp.discussion.Forum
	10, // 1: eolymp.discussion.ListForumsInput.filters:type_name -> eolymp.discussion.ListForumsInput.Filter
	11, // 2: eolymp.discussion.ListForumsOutput.items:type_name -> eolymp.discussion.Forum
	11, // 3: eolymp.discussion.CreateForumInput.forum:type_name -> eolymp.discussion.Forum
	11, // 4: eolymp.discussion.UpdateForumInput.forum:type_name -> eolymp.discussion.Forum
	12, // 5: eolymp.discussion.ListForumsInput.Filter.id:type_name -> eolymp.wellknown.ExpressionID
	12, // 6: eolymp.discussion.ListForumsInput.Filter.parent_id:type_name -> eolymp.wellknown.ExpressionID
	0,  // 7: eolymp.discussion.ForumService.DescribeForum:input_type -> eolymp.discussion.DescribeForumInput
	2,  // 8: eolymp.discussion.ForumService.ListForums:input_type -> eolymp.discussion.ListForumsInput
	4,  // 9: eolymp.discussion.ForumService.CreateForum:input_type -> eolymp.discussion.CreateForumInput
	6,  // 10: eolymp.discussion.ForumService.UpdateForum:input_type -> eolymp.discussion.UpdateForumInput
	8,  // 11: eolymp.discussion.ForumService.DeleteForum:input_type -> eolymp.discussion.DeleteForumInput
	1,  // 12: eolymp.discussion.ForumService.DescribeForum:output_type -> eolymp.discussion.DescribeForumOutput
	3,  // 13: eolymp.discussion.ForumService.ListForums:output_type -> eolymp.discussion.ListForumsOutput
	5,  // 14: eolymp.discussion.ForumService.CreateForum:output_type -> eolymp.discussion.CreateForumOutput
	7,  // 15: eolymp.discussion.ForumService.UpdateForum:output_type -> eolymp.discussion.UpdateForumOutput
	9,  // 16: eolymp.discussion.ForumService.DeleteForum:output_type -> eolymp.discussion.DeleteForumOutput
	12, // [12:17] is the sub-list for method output_type
	7,  // [7:12] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_eolymp_discussion_forum_service_proto_init() }
func file_eolymp_discussion_forum_service_proto_init() {
	if File_eolymp_discussion_forum_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_eolymp_discussion_forum_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DescribeForumInput); i {
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
		file_eolymp_discussion_forum_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DescribeForumOutput); i {
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
		file_eolymp_discussion_forum_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListForumsInput); i {
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
		file_eolymp_discussion_forum_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListForumsOutput); i {
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
		file_eolymp_discussion_forum_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateForumInput); i {
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
		file_eolymp_discussion_forum_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateForumOutput); i {
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
		file_eolymp_discussion_forum_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateForumInput); i {
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
		file_eolymp_discussion_forum_service_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateForumOutput); i {
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
		file_eolymp_discussion_forum_service_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteForumInput); i {
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
		file_eolymp_discussion_forum_service_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteForumOutput); i {
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
		file_eolymp_discussion_forum_service_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListForumsInput_Filter); i {
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
			RawDescriptor: file_eolymp_discussion_forum_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_eolymp_discussion_forum_service_proto_goTypes,
		DependencyIndexes: file_eolymp_discussion_forum_service_proto_depIdxs,
		MessageInfos:      file_eolymp_discussion_forum_service_proto_msgTypes,
	}.Build()
	File_eolymp_discussion_forum_service_proto = out.File
	file_eolymp_discussion_forum_service_proto_rawDesc = nil
	file_eolymp_discussion_forum_service_proto_goTypes = nil
	file_eolymp_discussion_forum_service_proto_depIdxs = nil
}