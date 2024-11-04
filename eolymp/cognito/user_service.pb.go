// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v3.21.12
// source: eolymp/cognito/user_service.proto

package cognito

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

type UserChangedEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Before *User `protobuf:"bytes,1,opt,name=before,proto3" json:"before,omitempty"`
	After  *User `protobuf:"bytes,2,opt,name=after,proto3" json:"after,omitempty"`
}

func (x *UserChangedEvent) Reset() {
	*x = UserChangedEvent{}
	mi := &file_eolymp_cognito_user_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UserChangedEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserChangedEvent) ProtoMessage() {}

func (x *UserChangedEvent) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_cognito_user_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserChangedEvent.ProtoReflect.Descriptor instead.
func (*UserChangedEvent) Descriptor() ([]byte, []int) {
	return file_eolymp_cognito_user_service_proto_rawDescGZIP(), []int{0}
}

func (x *UserChangedEvent) GetBefore() *User {
	if x != nil {
		return x.Before
	}
	return nil
}

func (x *UserChangedEvent) GetAfter() *User {
	if x != nil {
		return x.After
	}
	return nil
}

type IntrospectUserInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *IntrospectUserInput) Reset() {
	*x = IntrospectUserInput{}
	mi := &file_eolymp_cognito_user_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *IntrospectUserInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IntrospectUserInput) ProtoMessage() {}

func (x *IntrospectUserInput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_cognito_user_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IntrospectUserInput.ProtoReflect.Descriptor instead.
func (*IntrospectUserInput) Descriptor() ([]byte, []int) {
	return file_eolymp_cognito_user_service_proto_rawDescGZIP(), []int{1}
}

type IntrospectUserOutput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User *User `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *IntrospectUserOutput) Reset() {
	*x = IntrospectUserOutput{}
	mi := &file_eolymp_cognito_user_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *IntrospectUserOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IntrospectUserOutput) ProtoMessage() {}

func (x *IntrospectUserOutput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_cognito_user_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IntrospectUserOutput.ProtoReflect.Descriptor instead.
func (*IntrospectUserOutput) Descriptor() ([]byte, []int) {
	return file_eolymp_cognito_user_service_proto_rawDescGZIP(), []int{2}
}

func (x *IntrospectUserOutput) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

type DescribeUserInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *DescribeUserInput) Reset() {
	*x = DescribeUserInput{}
	mi := &file_eolymp_cognito_user_service_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DescribeUserInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescribeUserInput) ProtoMessage() {}

func (x *DescribeUserInput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_cognito_user_service_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DescribeUserInput.ProtoReflect.Descriptor instead.
func (*DescribeUserInput) Descriptor() ([]byte, []int) {
	return file_eolymp_cognito_user_service_proto_rawDescGZIP(), []int{3}
}

func (x *DescribeUserInput) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type DescribeUserOutput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User *User `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *DescribeUserOutput) Reset() {
	*x = DescribeUserOutput{}
	mi := &file_eolymp_cognito_user_service_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DescribeUserOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescribeUserOutput) ProtoMessage() {}

func (x *DescribeUserOutput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_cognito_user_service_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DescribeUserOutput.ProtoReflect.Descriptor instead.
func (*DescribeUserOutput) Descriptor() ([]byte, []int) {
	return file_eolymp_cognito_user_service_proto_rawDescGZIP(), []int{4}
}

func (x *DescribeUserOutput) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

type ListUsersInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// pagination
	Offset int32 `protobuf:"varint,10,opt,name=offset,proto3" json:"offset,omitempty"`
	Size   int32 `protobuf:"varint,11,opt,name=size,proto3" json:"size,omitempty"`
	// data filters
	Filters *ListUsersInput_Filter `protobuf:"bytes,40,opt,name=filters,proto3" json:"filters,omitempty"`
}

func (x *ListUsersInput) Reset() {
	*x = ListUsersInput{}
	mi := &file_eolymp_cognito_user_service_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListUsersInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListUsersInput) ProtoMessage() {}

func (x *ListUsersInput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_cognito_user_service_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListUsersInput.ProtoReflect.Descriptor instead.
func (*ListUsersInput) Descriptor() ([]byte, []int) {
	return file_eolymp_cognito_user_service_proto_rawDescGZIP(), []int{5}
}

func (x *ListUsersInput) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *ListUsersInput) GetSize() int32 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *ListUsersInput) GetFilters() *ListUsersInput_Filter {
	if x != nil {
		return x.Filters
	}
	return nil
}

type ListUsersOutput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Total int32   `protobuf:"varint,1,opt,name=total,proto3" json:"total,omitempty"`
	Items []*User `protobuf:"bytes,2,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *ListUsersOutput) Reset() {
	*x = ListUsersOutput{}
	mi := &file_eolymp_cognito_user_service_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListUsersOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListUsersOutput) ProtoMessage() {}

func (x *ListUsersOutput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_cognito_user_service_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListUsersOutput.ProtoReflect.Descriptor instead.
func (*ListUsersOutput) Descriptor() ([]byte, []int) {
	return file_eolymp_cognito_user_service_proto_rawDescGZIP(), []int{6}
}

func (x *ListUsersOutput) GetTotal() int32 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *ListUsersOutput) GetItems() []*User {
	if x != nil {
		return x.Items
	}
	return nil
}

type ListUsersInput_Filter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       []*wellknown.ExpressionID     `protobuf:"bytes,1,rep,name=id,proto3" json:"id,omitempty"`
	Username []*wellknown.ExpressionString `protobuf:"bytes,2,rep,name=username,proto3" json:"username,omitempty"`
}

func (x *ListUsersInput_Filter) Reset() {
	*x = ListUsersInput_Filter{}
	mi := &file_eolymp_cognito_user_service_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListUsersInput_Filter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListUsersInput_Filter) ProtoMessage() {}

func (x *ListUsersInput_Filter) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_cognito_user_service_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListUsersInput_Filter.ProtoReflect.Descriptor instead.
func (*ListUsersInput_Filter) Descriptor() ([]byte, []int) {
	return file_eolymp_cognito_user_service_proto_rawDescGZIP(), []int{5, 0}
}

func (x *ListUsersInput_Filter) GetId() []*wellknown.ExpressionID {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *ListUsersInput_Filter) GetUsername() []*wellknown.ExpressionString {
	if x != nil {
		return x.Username
	}
	return nil
}

var File_eolymp_cognito_user_service_proto protoreflect.FileDescriptor

var file_eolymp_cognito_user_service_proto_rawDesc = []byte{
	0x0a, 0x21, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x63, 0x6f, 0x67, 0x6e, 0x69, 0x74, 0x6f,
	0x2f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x63, 0x6f, 0x67, 0x6e,
	0x69, 0x74, 0x6f, 0x1a, 0x1d, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x22, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x72, 0x61, 0x74, 0x65, 0x6c, 0x69, 0x6d, 0x69, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x61,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x63, 0x6f, 0x70, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x63,
	0x6f, 0x67, 0x6e, 0x69, 0x74, 0x6f, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x21, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x77, 0x65, 0x6c, 0x6c, 0x6b, 0x6e,
	0x6f, 0x77, 0x6e, 0x2f, 0x65, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x6c, 0x0a, 0x10, 0x55, 0x73, 0x65, 0x72, 0x43, 0x68, 0x61, 0x6e,
	0x67, 0x65, 0x64, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x2c, 0x0a, 0x06, 0x62, 0x65, 0x66, 0x6f,
	0x72, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d,
	0x70, 0x2e, 0x63, 0x6f, 0x67, 0x6e, 0x69, 0x74, 0x6f, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x06,
	0x62, 0x65, 0x66, 0x6f, 0x72, 0x65, 0x12, 0x2a, 0x0a, 0x05, 0x61, 0x66, 0x74, 0x65, 0x72, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x63,
	0x6f, 0x67, 0x6e, 0x69, 0x74, 0x6f, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x05, 0x61, 0x66, 0x74,
	0x65, 0x72, 0x22, 0x15, 0x0a, 0x13, 0x49, 0x6e, 0x74, 0x72, 0x6f, 0x73, 0x70, 0x65, 0x63, 0x74,
	0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x22, 0x40, 0x0a, 0x14, 0x49, 0x6e, 0x74,
	0x72, 0x6f, 0x73, 0x70, 0x65, 0x63, 0x74, 0x55, 0x73, 0x65, 0x72, 0x4f, 0x75, 0x74, 0x70, 0x75,
	0x74, 0x12, 0x28, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x14, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x63, 0x6f, 0x67, 0x6e, 0x69, 0x74, 0x6f,
	0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x22, 0x2c, 0x0a, 0x11, 0x44,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x70, 0x75, 0x74,
	0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x3e, 0x0a, 0x12, 0x44, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x62, 0x65, 0x55, 0x73, 0x65, 0x72, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x12,
	0x28, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e,
	0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x63, 0x6f, 0x67, 0x6e, 0x69, 0x74, 0x6f, 0x2e, 0x55,
	0x73, 0x65, 0x72, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x22, 0xf7, 0x01, 0x0a, 0x0e, 0x4c, 0x69,
	0x73, 0x74, 0x55, 0x73, 0x65, 0x72, 0x73, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x16, 0x0a, 0x06,
	0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6f, 0x66,
	0x66, 0x73, 0x65, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x0b, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x3f, 0x0a, 0x07, 0x66, 0x69, 0x6c, 0x74,
	0x65, 0x72, 0x73, 0x18, 0x28, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x65, 0x6f, 0x6c, 0x79,
	0x6d, 0x70, 0x2e, 0x63, 0x6f, 0x67, 0x6e, 0x69, 0x74, 0x6f, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x55,
	0x73, 0x65, 0x72, 0x73, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x2e, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72,
	0x52, 0x07, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x1a, 0x78, 0x0a, 0x06, 0x46, 0x69, 0x6c,
	0x74, 0x65, 0x72, 0x12, 0x2e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x1e, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x77, 0x65, 0x6c, 0x6c, 0x6b, 0x6e, 0x6f,
	0x77, 0x6e, 0x2e, 0x45, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x3e, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x77,
	0x65, 0x6c, 0x6c, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x2e, 0x45, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e,
	0x61, 0x6d, 0x65, 0x22, 0x53, 0x0a, 0x0f, 0x4c, 0x69, 0x73, 0x74, 0x55, 0x73, 0x65, 0x72, 0x73,
	0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x2a, 0x0a, 0x05,
	0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x65, 0x6f,
	0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x63, 0x6f, 0x67, 0x6e, 0x69, 0x74, 0x6f, 0x2e, 0x55, 0x73, 0x65,
	0x72, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x32, 0xaf, 0x02, 0x0a, 0x0b, 0x55, 0x73, 0x65,
	0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x98, 0x01, 0x0a, 0x0c, 0x44, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x62, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x21, 0x2e, 0x65, 0x6f, 0x6c, 0x79,
	0x6d, 0x70, 0x2e, 0x63, 0x6f, 0x67, 0x6e, 0x69, 0x74, 0x6f, 0x2e, 0x44, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x62, 0x65, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x22, 0x2e, 0x65,
	0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x63, 0x6f, 0x67, 0x6e, 0x69, 0x74, 0x6f, 0x2e, 0x44, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x55, 0x73, 0x65, 0x72, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74,
	0x22, 0x41, 0xea, 0xe2, 0x0a, 0x0c, 0xf5, 0xe2, 0x0a, 0x00, 0x00, 0x48, 0x43, 0xf8, 0xe2, 0x0a,
	0xf4, 0x03, 0x82, 0xe3, 0x0a, 0x15, 0x8a, 0xe3, 0x0a, 0x11, 0x63, 0x6f, 0x67, 0x6e, 0x69, 0x74,
	0x6f, 0x3a, 0x75, 0x73, 0x65, 0x72, 0x3a, 0x72, 0x65, 0x61, 0x64, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x12, 0x12, 0x10, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2f, 0x7b, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x7d, 0x12, 0x84, 0x01, 0x0a, 0x09, 0x4c, 0x69, 0x73, 0x74, 0x55, 0x73, 0x65, 0x72,
	0x73, 0x12, 0x1e, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x63, 0x6f, 0x67, 0x6e, 0x69,
	0x74, 0x6f, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x55, 0x73, 0x65, 0x72, 0x73, 0x49, 0x6e, 0x70, 0x75,
	0x74, 0x1a, 0x1f, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x63, 0x6f, 0x67, 0x6e, 0x69,
	0x74, 0x6f, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x55, 0x73, 0x65, 0x72, 0x73, 0x4f, 0x75, 0x74, 0x70,
	0x75, 0x74, 0x22, 0x36, 0xea, 0xe2, 0x0a, 0x0b, 0xf5, 0xe2, 0x0a, 0x00, 0x00, 0x70, 0x41, 0xf8,
	0xe2, 0x0a, 0x64, 0x82, 0xe3, 0x0a, 0x15, 0x8a, 0xe3, 0x0a, 0x11, 0x63, 0x6f, 0x67, 0x6e, 0x69,
	0x74, 0x6f, 0x3a, 0x75, 0x73, 0x65, 0x72, 0x3a, 0x72, 0x65, 0x61, 0x64, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x08, 0x12, 0x06, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x42, 0x31, 0x5a, 0x2f, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f,
	0x67, 0x6f, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x63, 0x6f,
	0x67, 0x6e, 0x69, 0x74, 0x6f, 0x3b, 0x63, 0x6f, 0x67, 0x6e, 0x69, 0x74, 0x6f, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_eolymp_cognito_user_service_proto_rawDescOnce sync.Once
	file_eolymp_cognito_user_service_proto_rawDescData = file_eolymp_cognito_user_service_proto_rawDesc
)

func file_eolymp_cognito_user_service_proto_rawDescGZIP() []byte {
	file_eolymp_cognito_user_service_proto_rawDescOnce.Do(func() {
		file_eolymp_cognito_user_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_eolymp_cognito_user_service_proto_rawDescData)
	})
	return file_eolymp_cognito_user_service_proto_rawDescData
}

var file_eolymp_cognito_user_service_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_eolymp_cognito_user_service_proto_goTypes = []any{
	(*UserChangedEvent)(nil),           // 0: eolymp.cognito.UserChangedEvent
	(*IntrospectUserInput)(nil),        // 1: eolymp.cognito.IntrospectUserInput
	(*IntrospectUserOutput)(nil),       // 2: eolymp.cognito.IntrospectUserOutput
	(*DescribeUserInput)(nil),          // 3: eolymp.cognito.DescribeUserInput
	(*DescribeUserOutput)(nil),         // 4: eolymp.cognito.DescribeUserOutput
	(*ListUsersInput)(nil),             // 5: eolymp.cognito.ListUsersInput
	(*ListUsersOutput)(nil),            // 6: eolymp.cognito.ListUsersOutput
	(*ListUsersInput_Filter)(nil),      // 7: eolymp.cognito.ListUsersInput.Filter
	(*User)(nil),                       // 8: eolymp.cognito.User
	(*wellknown.ExpressionID)(nil),     // 9: eolymp.wellknown.ExpressionID
	(*wellknown.ExpressionString)(nil), // 10: eolymp.wellknown.ExpressionString
}
var file_eolymp_cognito_user_service_proto_depIdxs = []int32{
	8,  // 0: eolymp.cognito.UserChangedEvent.before:type_name -> eolymp.cognito.User
	8,  // 1: eolymp.cognito.UserChangedEvent.after:type_name -> eolymp.cognito.User
	8,  // 2: eolymp.cognito.IntrospectUserOutput.user:type_name -> eolymp.cognito.User
	8,  // 3: eolymp.cognito.DescribeUserOutput.user:type_name -> eolymp.cognito.User
	7,  // 4: eolymp.cognito.ListUsersInput.filters:type_name -> eolymp.cognito.ListUsersInput.Filter
	8,  // 5: eolymp.cognito.ListUsersOutput.items:type_name -> eolymp.cognito.User
	9,  // 6: eolymp.cognito.ListUsersInput.Filter.id:type_name -> eolymp.wellknown.ExpressionID
	10, // 7: eolymp.cognito.ListUsersInput.Filter.username:type_name -> eolymp.wellknown.ExpressionString
	3,  // 8: eolymp.cognito.UserService.DescribeUser:input_type -> eolymp.cognito.DescribeUserInput
	5,  // 9: eolymp.cognito.UserService.ListUsers:input_type -> eolymp.cognito.ListUsersInput
	4,  // 10: eolymp.cognito.UserService.DescribeUser:output_type -> eolymp.cognito.DescribeUserOutput
	6,  // 11: eolymp.cognito.UserService.ListUsers:output_type -> eolymp.cognito.ListUsersOutput
	10, // [10:12] is the sub-list for method output_type
	8,  // [8:10] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_eolymp_cognito_user_service_proto_init() }
func file_eolymp_cognito_user_service_proto_init() {
	if File_eolymp_cognito_user_service_proto != nil {
		return
	}
	file_eolymp_cognito_user_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_eolymp_cognito_user_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_eolymp_cognito_user_service_proto_goTypes,
		DependencyIndexes: file_eolymp_cognito_user_service_proto_depIdxs,
		MessageInfos:      file_eolymp_cognito_user_service_proto_msgTypes,
	}.Build()
	File_eolymp_cognito_user_service_proto = out.File
	file_eolymp_cognito_user_service_proto_rawDesc = nil
	file_eolymp_cognito_user_service_proto_goTypes = nil
	file_eolymp_cognito_user_service_proto_depIdxs = nil
}
