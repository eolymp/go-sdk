// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.18.1
// source: eolymp/judge/acl.proto

package judge

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

type Permission_Role int32

const (
	Permission_NONE   Permission_Role = 0
	Permission_ADMIN  Permission_Role = 1
	Permission_VIEWER Permission_Role = 2
)

// Enum value maps for Permission_Role.
var (
	Permission_Role_name = map[int32]string{
		0: "NONE",
		1: "ADMIN",
		2: "VIEWER",
	}
	Permission_Role_value = map[string]int32{
		"NONE":   0,
		"ADMIN":  1,
		"VIEWER": 2,
	}
)

func (x Permission_Role) Enum() *Permission_Role {
	p := new(Permission_Role)
	*p = x
	return p
}

func (x Permission_Role) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Permission_Role) Descriptor() protoreflect.EnumDescriptor {
	return file_eolymp_judge_acl_proto_enumTypes[0].Descriptor()
}

func (Permission_Role) Type() protoreflect.EnumType {
	return &file_eolymp_judge_acl_proto_enumTypes[0]
}

func (x Permission_Role) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Permission_Role.Descriptor instead.
func (Permission_Role) EnumDescriptor() ([]byte, []int) {
	return file_eolymp_judge_acl_proto_rawDescGZIP(), []int{0, 0}
}

type Permission struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     string          `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Role   Permission_Role `protobuf:"varint,2,opt,name=role,proto3,enum=eolymp.judge.Permission_Role" json:"role,omitempty"`
	UserId string          `protobuf:"bytes,3,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *Permission) Reset() {
	*x = Permission{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_judge_acl_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Permission) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Permission) ProtoMessage() {}

func (x *Permission) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_judge_acl_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Permission.ProtoReflect.Descriptor instead.
func (*Permission) Descriptor() ([]byte, []int) {
	return file_eolymp_judge_acl_proto_rawDescGZIP(), []int{0}
}

func (x *Permission) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Permission) GetRole() Permission_Role {
	if x != nil {
		return x.Role
	}
	return Permission_NONE
}

func (x *Permission) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type GrantPermissionInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Role   string `protobuf:"bytes,2,opt,name=role,proto3" json:"role,omitempty"`
}

func (x *GrantPermissionInput) Reset() {
	*x = GrantPermissionInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_judge_acl_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GrantPermissionInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GrantPermissionInput) ProtoMessage() {}

func (x *GrantPermissionInput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_judge_acl_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GrantPermissionInput.ProtoReflect.Descriptor instead.
func (*GrantPermissionInput) Descriptor() ([]byte, []int) {
	return file_eolymp_judge_acl_proto_rawDescGZIP(), []int{1}
}

func (x *GrantPermissionInput) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *GrantPermissionInput) GetRole() string {
	if x != nil {
		return x.Role
	}
	return ""
}

type GrantPermissionOutput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GrantPermissionOutput) Reset() {
	*x = GrantPermissionOutput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_judge_acl_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GrantPermissionOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GrantPermissionOutput) ProtoMessage() {}

func (x *GrantPermissionOutput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_judge_acl_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GrantPermissionOutput.ProtoReflect.Descriptor instead.
func (*GrantPermissionOutput) Descriptor() ([]byte, []int) {
	return file_eolymp_judge_acl_proto_rawDescGZIP(), []int{2}
}

type RevokePermissionInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *RevokePermissionInput) Reset() {
	*x = RevokePermissionInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_judge_acl_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RevokePermissionInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RevokePermissionInput) ProtoMessage() {}

func (x *RevokePermissionInput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_judge_acl_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RevokePermissionInput.ProtoReflect.Descriptor instead.
func (*RevokePermissionInput) Descriptor() ([]byte, []int) {
	return file_eolymp_judge_acl_proto_rawDescGZIP(), []int{3}
}

func (x *RevokePermissionInput) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type RevokePermissionOutput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *RevokePermissionOutput) Reset() {
	*x = RevokePermissionOutput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_judge_acl_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RevokePermissionOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RevokePermissionOutput) ProtoMessage() {}

func (x *RevokePermissionOutput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_judge_acl_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RevokePermissionOutput.ProtoReflect.Descriptor instead.
func (*RevokePermissionOutput) Descriptor() ([]byte, []int) {
	return file_eolymp_judge_acl_proto_rawDescGZIP(), []int{4}
}

type DescribePermissionInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *DescribePermissionInput) Reset() {
	*x = DescribePermissionInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_judge_acl_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DescribePermissionInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescribePermissionInput) ProtoMessage() {}

func (x *DescribePermissionInput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_judge_acl_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DescribePermissionInput.ProtoReflect.Descriptor instead.
func (*DescribePermissionInput) Descriptor() ([]byte, []int) {
	return file_eolymp_judge_acl_proto_rawDescGZIP(), []int{5}
}

func (x *DescribePermissionInput) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type DescribePermissionOutput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Permission *Permission `protobuf:"bytes,1,opt,name=permission,proto3" json:"permission,omitempty"`
}

func (x *DescribePermissionOutput) Reset() {
	*x = DescribePermissionOutput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_judge_acl_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DescribePermissionOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescribePermissionOutput) ProtoMessage() {}

func (x *DescribePermissionOutput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_judge_acl_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DescribePermissionOutput.ProtoReflect.Descriptor instead.
func (*DescribePermissionOutput) Descriptor() ([]byte, []int) {
	return file_eolymp_judge_acl_proto_rawDescGZIP(), []int{6}
}

func (x *DescribePermissionOutput) GetPermission() *Permission {
	if x != nil {
		return x.Permission
	}
	return nil
}

type ListPermissionsInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SpaceId string                       `protobuf:"bytes,1,opt,name=space_id,json=spaceId,proto3" json:"space_id,omitempty"`
	Offset  int32                        `protobuf:"varint,10,opt,name=offset,proto3" json:"offset,omitempty"`
	Size    int32                        `protobuf:"varint,11,opt,name=size,proto3" json:"size,omitempty"`
	Filters *ListPermissionsInput_Filter `protobuf:"bytes,40,opt,name=filters,proto3" json:"filters,omitempty"`
}

func (x *ListPermissionsInput) Reset() {
	*x = ListPermissionsInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_judge_acl_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListPermissionsInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListPermissionsInput) ProtoMessage() {}

func (x *ListPermissionsInput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_judge_acl_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListPermissionsInput.ProtoReflect.Descriptor instead.
func (*ListPermissionsInput) Descriptor() ([]byte, []int) {
	return file_eolymp_judge_acl_proto_rawDescGZIP(), []int{7}
}

func (x *ListPermissionsInput) GetSpaceId() string {
	if x != nil {
		return x.SpaceId
	}
	return ""
}

func (x *ListPermissionsInput) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *ListPermissionsInput) GetSize() int32 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *ListPermissionsInput) GetFilters() *ListPermissionsInput_Filter {
	if x != nil {
		return x.Filters
	}
	return nil
}

type ListPermissionsOutput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Total int32         `protobuf:"varint,1,opt,name=total,proto3" json:"total,omitempty"`
	Items []*Permission `protobuf:"bytes,2,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *ListPermissionsOutput) Reset() {
	*x = ListPermissionsOutput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_judge_acl_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListPermissionsOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListPermissionsOutput) ProtoMessage() {}

func (x *ListPermissionsOutput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_judge_acl_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListPermissionsOutput.ProtoReflect.Descriptor instead.
func (*ListPermissionsOutput) Descriptor() ([]byte, []int) {
	return file_eolymp_judge_acl_proto_rawDescGZIP(), []int{8}
}

func (x *ListPermissionsOutput) GetTotal() int32 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *ListPermissionsOutput) GetItems() []*Permission {
	if x != nil {
		return x.Items
	}
	return nil
}

type ListPermissionsInput_Filter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     []*wellknown.ExpressionID   `protobuf:"bytes,1,rep,name=id,proto3" json:"id,omitempty"`
	UserId []*wellknown.ExpressionID   `protobuf:"bytes,2,rep,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Role   []*wellknown.ExpressionEnum `protobuf:"bytes,3,rep,name=role,proto3" json:"role,omitempty"`
}

func (x *ListPermissionsInput_Filter) Reset() {
	*x = ListPermissionsInput_Filter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_judge_acl_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListPermissionsInput_Filter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListPermissionsInput_Filter) ProtoMessage() {}

func (x *ListPermissionsInput_Filter) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_judge_acl_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListPermissionsInput_Filter.ProtoReflect.Descriptor instead.
func (*ListPermissionsInput_Filter) Descriptor() ([]byte, []int) {
	return file_eolymp_judge_acl_proto_rawDescGZIP(), []int{7, 0}
}

func (x *ListPermissionsInput_Filter) GetId() []*wellknown.ExpressionID {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *ListPermissionsInput_Filter) GetUserId() []*wellknown.ExpressionID {
	if x != nil {
		return x.UserId
	}
	return nil
}

func (x *ListPermissionsInput_Filter) GetRole() []*wellknown.ExpressionEnum {
	if x != nil {
		return x.Role
	}
	return nil
}

var File_eolymp_judge_acl_proto protoreflect.FileDescriptor

var file_eolymp_judge_acl_proto_rawDesc = []byte{
	0x0a, 0x16, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x6a, 0x75, 0x64, 0x67, 0x65, 0x2f, 0x61,
	0x63, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70,
	0x2e, 0x6a, 0x75, 0x64, 0x67, 0x65, 0x1a, 0x1d, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x61,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x68, 0x74, 0x74, 0x70, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x22, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x61, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x72, 0x61, 0x74, 0x65, 0x6c, 0x69,
	0x6d, 0x69, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x21, 0x65, 0x6f, 0x6c, 0x79, 0x6d,
	0x70, 0x2f, 0x77, 0x65, 0x6c, 0x6c, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x2f, 0x65, 0x78, 0x70, 0x72,
	0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x91, 0x01, 0x0a,
	0x0a, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x31, 0x0a, 0x04, 0x72,
	0x6f, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1d, 0x2e, 0x65, 0x6f, 0x6c, 0x79,
	0x6d, 0x70, 0x2e, 0x6a, 0x75, 0x64, 0x67, 0x65, 0x2e, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x12, 0x17,
	0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x27, 0x0a, 0x04, 0x52, 0x6f, 0x6c, 0x65, 0x12,
	0x08, 0x0a, 0x04, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x41, 0x44, 0x4d,
	0x49, 0x4e, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x56, 0x49, 0x45, 0x57, 0x45, 0x52, 0x10, 0x02,
	0x22, 0x43, 0x0a, 0x14, 0x47, 0x72, 0x61, 0x6e, 0x74, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x72, 0x6f, 0x6c, 0x65, 0x22, 0x17, 0x0a, 0x15, 0x47, 0x72, 0x61, 0x6e, 0x74, 0x50, 0x65,
	0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x22, 0x30,
	0x0a, 0x15, 0x52, 0x65, 0x76, 0x6f, 0x6b, 0x65, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x22, 0x18, 0x0a, 0x16, 0x52, 0x65, 0x76, 0x6f, 0x6b, 0x65, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x22, 0x32, 0x0a, 0x17, 0x44, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x54,
	0x0a, 0x18, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x12, 0x38, 0x0a, 0x0a, 0x70, 0x65,
	0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18,
	0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x6a, 0x75, 0x64, 0x67, 0x65, 0x2e, 0x50, 0x65,
	0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x22, 0xcc, 0x02, 0x0a, 0x14, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x65, 0x72,
	0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x19, 0x0a,
	0x08, 0x73, 0x70, 0x61, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x73, 0x70, 0x61, 0x63, 0x65, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73,
	0x65, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74,
	0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04,
	0x73, 0x69, 0x7a, 0x65, 0x12, 0x43, 0x0a, 0x07, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x18,
	0x28, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x6a,
	0x75, 0x64, 0x67, 0x65, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x73, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x2e, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72,
	0x52, 0x07, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x1a, 0xa7, 0x01, 0x0a, 0x06, 0x46, 0x69,
	0x6c, 0x74, 0x65, 0x72, 0x12, 0x2e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x1e, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x77, 0x65, 0x6c, 0x6c, 0x6b, 0x6e,
	0x6f, 0x77, 0x6e, 0x2e, 0x45, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x44,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x37, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x77,
	0x65, 0x6c, 0x6c, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x2e, 0x45, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x49, 0x44, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x34, 0x0a,
	0x04, 0x72, 0x6f, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x65, 0x6f,
	0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x77, 0x65, 0x6c, 0x6c, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x2e, 0x45,
	0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x45, 0x6e, 0x75, 0x6d, 0x52, 0x04, 0x72,
	0x6f, 0x6c, 0x65, 0x22, 0x5d, 0x0a, 0x15, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x65, 0x72, 0x6d, 0x69,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x12, 0x14, 0x0a, 0x05,
	0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x74, 0x6f, 0x74,
	0x61, 0x6c, 0x12, 0x2e, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x18, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x6a, 0x75, 0x64, 0x67, 0x65,
	0x2e, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x05, 0x69, 0x74, 0x65,
	0x6d, 0x73, 0x32, 0xb6, 0x04, 0x0a, 0x03, 0x41, 0x63, 0x6c, 0x12, 0x89, 0x01, 0x0a, 0x0f, 0x47,
	0x72, 0x61, 0x6e, 0x74, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x22,
	0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x6a, 0x75, 0x64, 0x67, 0x65, 0x2e, 0x47, 0x72,
	0x61, 0x6e, 0x74, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x70,
	0x75, 0x74, 0x1a, 0x23, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x6a, 0x75, 0x64, 0x67,
	0x65, 0x2e, 0x47, 0x72, 0x61, 0x6e, 0x74, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x22, 0x2d, 0xea, 0xe2, 0x0a, 0x0b, 0xf5, 0xe2, 0x0a,
	0x00, 0x00, 0xa0, 0x40, 0xf8, 0xe2, 0x0a, 0x14, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x18, 0x1a, 0x16,
	0x2f, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x7b, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x7d, 0x12, 0x8c, 0x01, 0x0a, 0x10, 0x52, 0x65, 0x76, 0x6f, 0x6b,
	0x65, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x23, 0x2e, 0x65, 0x6f,
	0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x6a, 0x75, 0x64, 0x67, 0x65, 0x2e, 0x52, 0x65, 0x76, 0x6f, 0x6b,
	0x65, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x70, 0x75, 0x74,
	0x1a, 0x24, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x6a, 0x75, 0x64, 0x67, 0x65, 0x2e,
	0x52, 0x65, 0x76, 0x6f, 0x6b, 0x65, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x22, 0x2d, 0xea, 0xe2, 0x0a, 0x0b, 0xf5, 0xe2, 0x0a, 0x00,
	0x00, 0xa0, 0x40, 0xf8, 0xe2, 0x0a, 0x14, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x18, 0x2a, 0x16, 0x2f,
	0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x7b, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x7d, 0x12, 0x92, 0x01, 0x0a, 0x12, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x62, 0x65, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x25, 0x2e, 0x65,
	0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x6a, 0x75, 0x64, 0x67, 0x65, 0x2e, 0x44, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x62, 0x65, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x6e,
	0x70, 0x75, 0x74, 0x1a, 0x26, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x6a, 0x75, 0x64,
	0x67, 0x65, 0x2e, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x50, 0x65, 0x72, 0x6d, 0x69,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x22, 0x2d, 0xea, 0xe2, 0x0a,
	0x0b, 0xf5, 0xe2, 0x0a, 0x00, 0x00, 0x20, 0x41, 0xf8, 0xe2, 0x0a, 0x14, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x18, 0x12, 0x16, 0x2f, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73,
	0x2f, 0x7b, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x7d, 0x12, 0x7f, 0x0a, 0x0f, 0x4c, 0x69,
	0x73, 0x74, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x22, 0x2e,
	0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x6a, 0x75, 0x64, 0x67, 0x65, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x49, 0x6e, 0x70, 0x75,
	0x74, 0x1a, 0x23, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x6a, 0x75, 0x64, 0x67, 0x65,
	0x2e, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73,
	0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x22, 0x23, 0xea, 0xe2, 0x0a, 0x0b, 0xf5, 0xe2, 0x0a, 0x00,
	0x00, 0x20, 0x41, 0xf8, 0xe2, 0x0a, 0x14, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0e, 0x12, 0x0c, 0x2f,
	0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x42, 0x2d, 0x5a, 0x2b, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70,
	0x2f, 0x67, 0x6f, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x6a,
	0x75, 0x64, 0x67, 0x65, 0x3b, 0x6a, 0x75, 0x64, 0x67, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_eolymp_judge_acl_proto_rawDescOnce sync.Once
	file_eolymp_judge_acl_proto_rawDescData = file_eolymp_judge_acl_proto_rawDesc
)

func file_eolymp_judge_acl_proto_rawDescGZIP() []byte {
	file_eolymp_judge_acl_proto_rawDescOnce.Do(func() {
		file_eolymp_judge_acl_proto_rawDescData = protoimpl.X.CompressGZIP(file_eolymp_judge_acl_proto_rawDescData)
	})
	return file_eolymp_judge_acl_proto_rawDescData
}

var file_eolymp_judge_acl_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_eolymp_judge_acl_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_eolymp_judge_acl_proto_goTypes = []interface{}{
	(Permission_Role)(0),                // 0: eolymp.judge.Permission.Role
	(*Permission)(nil),                  // 1: eolymp.judge.Permission
	(*GrantPermissionInput)(nil),        // 2: eolymp.judge.GrantPermissionInput
	(*GrantPermissionOutput)(nil),       // 3: eolymp.judge.GrantPermissionOutput
	(*RevokePermissionInput)(nil),       // 4: eolymp.judge.RevokePermissionInput
	(*RevokePermissionOutput)(nil),      // 5: eolymp.judge.RevokePermissionOutput
	(*DescribePermissionInput)(nil),     // 6: eolymp.judge.DescribePermissionInput
	(*DescribePermissionOutput)(nil),    // 7: eolymp.judge.DescribePermissionOutput
	(*ListPermissionsInput)(nil),        // 8: eolymp.judge.ListPermissionsInput
	(*ListPermissionsOutput)(nil),       // 9: eolymp.judge.ListPermissionsOutput
	(*ListPermissionsInput_Filter)(nil), // 10: eolymp.judge.ListPermissionsInput.Filter
	(*wellknown.ExpressionID)(nil),      // 11: eolymp.wellknown.ExpressionID
	(*wellknown.ExpressionEnum)(nil),    // 12: eolymp.wellknown.ExpressionEnum
}
var file_eolymp_judge_acl_proto_depIdxs = []int32{
	0,  // 0: eolymp.judge.Permission.role:type_name -> eolymp.judge.Permission.Role
	1,  // 1: eolymp.judge.DescribePermissionOutput.permission:type_name -> eolymp.judge.Permission
	10, // 2: eolymp.judge.ListPermissionsInput.filters:type_name -> eolymp.judge.ListPermissionsInput.Filter
	1,  // 3: eolymp.judge.ListPermissionsOutput.items:type_name -> eolymp.judge.Permission
	11, // 4: eolymp.judge.ListPermissionsInput.Filter.id:type_name -> eolymp.wellknown.ExpressionID
	11, // 5: eolymp.judge.ListPermissionsInput.Filter.user_id:type_name -> eolymp.wellknown.ExpressionID
	12, // 6: eolymp.judge.ListPermissionsInput.Filter.role:type_name -> eolymp.wellknown.ExpressionEnum
	2,  // 7: eolymp.judge.Acl.GrantPermission:input_type -> eolymp.judge.GrantPermissionInput
	4,  // 8: eolymp.judge.Acl.RevokePermission:input_type -> eolymp.judge.RevokePermissionInput
	6,  // 9: eolymp.judge.Acl.DescribePermission:input_type -> eolymp.judge.DescribePermissionInput
	8,  // 10: eolymp.judge.Acl.ListPermissions:input_type -> eolymp.judge.ListPermissionsInput
	3,  // 11: eolymp.judge.Acl.GrantPermission:output_type -> eolymp.judge.GrantPermissionOutput
	5,  // 12: eolymp.judge.Acl.RevokePermission:output_type -> eolymp.judge.RevokePermissionOutput
	7,  // 13: eolymp.judge.Acl.DescribePermission:output_type -> eolymp.judge.DescribePermissionOutput
	9,  // 14: eolymp.judge.Acl.ListPermissions:output_type -> eolymp.judge.ListPermissionsOutput
	11, // [11:15] is the sub-list for method output_type
	7,  // [7:11] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_eolymp_judge_acl_proto_init() }
func file_eolymp_judge_acl_proto_init() {
	if File_eolymp_judge_acl_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_eolymp_judge_acl_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Permission); i {
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
		file_eolymp_judge_acl_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GrantPermissionInput); i {
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
		file_eolymp_judge_acl_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GrantPermissionOutput); i {
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
		file_eolymp_judge_acl_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RevokePermissionInput); i {
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
		file_eolymp_judge_acl_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RevokePermissionOutput); i {
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
		file_eolymp_judge_acl_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DescribePermissionInput); i {
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
		file_eolymp_judge_acl_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DescribePermissionOutput); i {
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
		file_eolymp_judge_acl_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListPermissionsInput); i {
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
		file_eolymp_judge_acl_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListPermissionsOutput); i {
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
		file_eolymp_judge_acl_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListPermissionsInput_Filter); i {
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
			RawDescriptor: file_eolymp_judge_acl_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_eolymp_judge_acl_proto_goTypes,
		DependencyIndexes: file_eolymp_judge_acl_proto_depIdxs,
		EnumInfos:         file_eolymp_judge_acl_proto_enumTypes,
		MessageInfos:      file_eolymp_judge_acl_proto_msgTypes,
	}.Build()
	File_eolymp_judge_acl_proto = out.File
	file_eolymp_judge_acl_proto_rawDesc = nil
	file_eolymp_judge_acl_proto_goTypes = nil
	file_eolymp_judge_acl_proto_depIdxs = nil
}
