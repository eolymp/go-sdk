// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.18.1
// source: eolymp/guardian/guardian.proto

package guardian

import (
	_ "github.com/eolymp/go-sdk/eolymp/annotations"
	universe "github.com/eolymp/go-sdk/eolymp/universe"
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

type ListPoliciesInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListPoliciesInput) Reset() {
	*x = ListPoliciesInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_guardian_guardian_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListPoliciesInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListPoliciesInput) ProtoMessage() {}

func (x *ListPoliciesInput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_guardian_guardian_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListPoliciesInput.ProtoReflect.Descriptor instead.
func (*ListPoliciesInput) Descriptor() ([]byte, []int) {
	return file_eolymp_guardian_guardian_proto_rawDescGZIP(), []int{0}
}

type ListPoliciesOutput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Total int32              `protobuf:"varint,1,opt,name=total,proto3" json:"total,omitempty"`
	Items []*universe.Policy `protobuf:"bytes,2,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *ListPoliciesOutput) Reset() {
	*x = ListPoliciesOutput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_guardian_guardian_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListPoliciesOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListPoliciesOutput) ProtoMessage() {}

func (x *ListPoliciesOutput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_guardian_guardian_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListPoliciesOutput.ProtoReflect.Descriptor instead.
func (*ListPoliciesOutput) Descriptor() ([]byte, []int) {
	return file_eolymp_guardian_guardian_proto_rawDescGZIP(), []int{1}
}

func (x *ListPoliciesOutput) GetTotal() int32 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *ListPoliciesOutput) GetItems() []*universe.Policy {
	if x != nil {
		return x.Items
	}
	return nil
}

type DescribePolicyInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *DescribePolicyInput) Reset() {
	*x = DescribePolicyInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_guardian_guardian_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DescribePolicyInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescribePolicyInput) ProtoMessage() {}

func (x *DescribePolicyInput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_guardian_guardian_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DescribePolicyInput.ProtoReflect.Descriptor instead.
func (*DescribePolicyInput) Descriptor() ([]byte, []int) {
	return file_eolymp_guardian_guardian_proto_rawDescGZIP(), []int{2}
}

func (x *DescribePolicyInput) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type DescribePolicyOutput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Policy *universe.Policy `protobuf:"bytes,1,opt,name=policy,proto3" json:"policy,omitempty"`
}

func (x *DescribePolicyOutput) Reset() {
	*x = DescribePolicyOutput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_guardian_guardian_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DescribePolicyOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescribePolicyOutput) ProtoMessage() {}

func (x *DescribePolicyOutput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_guardian_guardian_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DescribePolicyOutput.ProtoReflect.Descriptor instead.
func (*DescribePolicyOutput) Descriptor() ([]byte, []int) {
	return file_eolymp_guardian_guardian_proto_rawDescGZIP(), []int{3}
}

func (x *DescribePolicyOutput) GetPolicy() *universe.Policy {
	if x != nil {
		return x.Policy
	}
	return nil
}

type DefinePolicyInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name   string           `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Policy *universe.Policy `protobuf:"bytes,2,opt,name=policy,proto3" json:"policy,omitempty"`
}

func (x *DefinePolicyInput) Reset() {
	*x = DefinePolicyInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_guardian_guardian_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DefinePolicyInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DefinePolicyInput) ProtoMessage() {}

func (x *DefinePolicyInput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_guardian_guardian_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DefinePolicyInput.ProtoReflect.Descriptor instead.
func (*DefinePolicyInput) Descriptor() ([]byte, []int) {
	return file_eolymp_guardian_guardian_proto_rawDescGZIP(), []int{4}
}

func (x *DefinePolicyInput) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *DefinePolicyInput) GetPolicy() *universe.Policy {
	if x != nil {
		return x.Policy
	}
	return nil
}

type DefinePolicyOutput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DefinePolicyOutput) Reset() {
	*x = DefinePolicyOutput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_guardian_guardian_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DefinePolicyOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DefinePolicyOutput) ProtoMessage() {}

func (x *DefinePolicyOutput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_guardian_guardian_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DefinePolicyOutput.ProtoReflect.Descriptor instead.
func (*DefinePolicyOutput) Descriptor() ([]byte, []int) {
	return file_eolymp_guardian_guardian_proto_rawDescGZIP(), []int{5}
}

type RemovePolicyInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *RemovePolicyInput) Reset() {
	*x = RemovePolicyInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_guardian_guardian_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemovePolicyInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemovePolicyInput) ProtoMessage() {}

func (x *RemovePolicyInput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_guardian_guardian_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemovePolicyInput.ProtoReflect.Descriptor instead.
func (*RemovePolicyInput) Descriptor() ([]byte, []int) {
	return file_eolymp_guardian_guardian_proto_rawDescGZIP(), []int{6}
}

func (x *RemovePolicyInput) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type RemovePolicyOutput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *RemovePolicyOutput) Reset() {
	*x = RemovePolicyOutput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_guardian_guardian_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemovePolicyOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemovePolicyOutput) ProtoMessage() {}

func (x *RemovePolicyOutput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_guardian_guardian_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemovePolicyOutput.ProtoReflect.Descriptor instead.
func (*RemovePolicyOutput) Descriptor() ([]byte, []int) {
	return file_eolymp_guardian_guardian_proto_rawDescGZIP(), []int{7}
}

var File_eolymp_guardian_guardian_proto protoreflect.FileDescriptor

var file_eolymp_guardian_guardian_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x67, 0x75, 0x61, 0x72, 0x64, 0x69, 0x61,
	0x6e, 0x2f, 0x67, 0x75, 0x61, 0x72, 0x64, 0x69, 0x61, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0f, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x75, 0x6e, 0x69, 0x76, 0x65, 0x72, 0x73,
	0x65, 0x1a, 0x1d, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x22, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x72, 0x61, 0x74, 0x65, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x67, 0x75, 0x61,
	0x72, 0x64, 0x69, 0x61, 0x6e, 0x2f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x13, 0x0a, 0x11, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x69,
	0x65, 0x73, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x22, 0x59, 0x0a, 0x12, 0x4c, 0x69, 0x73, 0x74, 0x50,
	0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65, 0x73, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x12, 0x14, 0x0a,
	0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x74, 0x6f,
	0x74, 0x61, 0x6c, 0x12, 0x2d, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x17, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x75, 0x6e, 0x69, 0x76,
	0x65, 0x72, 0x73, 0x65, 0x2e, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x05, 0x69, 0x74, 0x65,
	0x6d, 0x73, 0x22, 0x29, 0x0a, 0x13, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x50, 0x6f,
	0x6c, 0x69, 0x63, 0x79, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x47, 0x0a,
	0x14, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x4f,
	0x75, 0x74, 0x70, 0x75, 0x74, 0x12, 0x2f, 0x0a, 0x06, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x75,
	0x6e, 0x69, 0x76, 0x65, 0x72, 0x73, 0x65, 0x2e, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x06,
	0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x22, 0x58, 0x0a, 0x11, 0x44, 0x65, 0x66, 0x69, 0x6e, 0x65,
	0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x2f, 0x0a, 0x06, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x17, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x75, 0x6e, 0x69, 0x76, 0x65, 0x72, 0x73,
	0x65, 0x2e, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x06, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79,
	0x22, 0x14, 0x0a, 0x12, 0x44, 0x65, 0x66, 0x69, 0x6e, 0x65, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79,
	0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x22, 0x27, 0x0a, 0x11, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65,
	0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22,
	0x14, 0x0a, 0x12, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x4f,
	0x75, 0x74, 0x70, 0x75, 0x74, 0x32, 0x94, 0x04, 0x0a, 0x08, 0x47, 0x75, 0x61, 0x72, 0x64, 0x69,
	0x61, 0x6e, 0x12, 0x79, 0x0a, 0x0c, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x69,
	0x65, 0x73, 0x12, 0x22, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x75, 0x6e, 0x69, 0x76,
	0x65, 0x72, 0x73, 0x65, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65,
	0x73, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x23, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e,
	0x75, 0x6e, 0x69, 0x76, 0x65, 0x72, 0x73, 0x65, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x6f, 0x6c,
	0x69, 0x63, 0x69, 0x65, 0x73, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x22, 0x20, 0xea, 0xe2, 0x0a,
	0x0b, 0xf5, 0xe2, 0x0a, 0x00, 0x00, 0x20, 0x41, 0xf8, 0xe2, 0x0a, 0x64, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x0b, 0x12, 0x09, 0x2f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65, 0x73, 0x12, 0x86, 0x01,
	0x0a, 0x0e, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79,
	0x12, 0x24, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x75, 0x6e, 0x69, 0x76, 0x65, 0x72,
	0x73, 0x65, 0x2e, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x50, 0x6f, 0x6c, 0x69, 0x63,
	0x79, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x25, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e,
	0x75, 0x6e, 0x69, 0x76, 0x65, 0x72, 0x73, 0x65, 0x2e, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62,
	0x65, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x22, 0x27, 0xea,
	0xe2, 0x0a, 0x0b, 0xf5, 0xe2, 0x0a, 0x00, 0x00, 0x20, 0x41, 0xf8, 0xe2, 0x0a, 0x64, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x12, 0x12, 0x10, 0x2f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65, 0x73, 0x2f,
	0x7b, 0x6e, 0x61, 0x6d, 0x65, 0x7d, 0x12, 0x80, 0x01, 0x0a, 0x0c, 0x44, 0x65, 0x66, 0x69, 0x6e,
	0x65, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0x22, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70,
	0x2e, 0x75, 0x6e, 0x69, 0x76, 0x65, 0x72, 0x73, 0x65, 0x2e, 0x44, 0x65, 0x66, 0x69, 0x6e, 0x65,
	0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x23, 0x2e, 0x65, 0x6f,
	0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x75, 0x6e, 0x69, 0x76, 0x65, 0x72, 0x73, 0x65, 0x2e, 0x44, 0x65,
	0x66, 0x69, 0x6e, 0x65, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74,
	0x22, 0x27, 0xea, 0xe2, 0x0a, 0x0b, 0xf5, 0xe2, 0x0a, 0x00, 0x00, 0x20, 0x41, 0xf8, 0xe2, 0x0a,
	0x64, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x12, 0x1a, 0x10, 0x2f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x69,
	0x65, 0x73, 0x2f, 0x7b, 0x6e, 0x61, 0x6d, 0x65, 0x7d, 0x12, 0x80, 0x01, 0x0a, 0x0c, 0x52, 0x65,
	0x6d, 0x6f, 0x76, 0x65, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0x22, 0x2e, 0x65, 0x6f, 0x6c,
	0x79, 0x6d, 0x70, 0x2e, 0x75, 0x6e, 0x69, 0x76, 0x65, 0x72, 0x73, 0x65, 0x2e, 0x52, 0x65, 0x6d,
	0x6f, 0x76, 0x65, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x23,
	0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x75, 0x6e, 0x69, 0x76, 0x65, 0x72, 0x73, 0x65,
	0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x4f, 0x75, 0x74,
	0x70, 0x75, 0x74, 0x22, 0x27, 0xea, 0xe2, 0x0a, 0x0b, 0xf5, 0xe2, 0x0a, 0x00, 0x00, 0x20, 0x41,
	0xf8, 0xe2, 0x0a, 0x64, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x12, 0x2a, 0x10, 0x2f, 0x70, 0x6f, 0x6c,
	0x69, 0x63, 0x69, 0x65, 0x73, 0x2f, 0x7b, 0x6e, 0x61, 0x6d, 0x65, 0x7d, 0x42, 0x33, 0x5a, 0x31,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6f, 0x6c, 0x79, 0x6d,
	0x70, 0x2f, 0x67, 0x6f, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f,
	0x67, 0x75, 0x61, 0x72, 0x64, 0x69, 0x61, 0x6e, 0x3b, 0x67, 0x75, 0x61, 0x72, 0x64, 0x69, 0x61,
	0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_eolymp_guardian_guardian_proto_rawDescOnce sync.Once
	file_eolymp_guardian_guardian_proto_rawDescData = file_eolymp_guardian_guardian_proto_rawDesc
)

func file_eolymp_guardian_guardian_proto_rawDescGZIP() []byte {
	file_eolymp_guardian_guardian_proto_rawDescOnce.Do(func() {
		file_eolymp_guardian_guardian_proto_rawDescData = protoimpl.X.CompressGZIP(file_eolymp_guardian_guardian_proto_rawDescData)
	})
	return file_eolymp_guardian_guardian_proto_rawDescData
}

var file_eolymp_guardian_guardian_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_eolymp_guardian_guardian_proto_goTypes = []interface{}{
	(*ListPoliciesInput)(nil),    // 0: eolymp.universe.ListPoliciesInput
	(*ListPoliciesOutput)(nil),   // 1: eolymp.universe.ListPoliciesOutput
	(*DescribePolicyInput)(nil),  // 2: eolymp.universe.DescribePolicyInput
	(*DescribePolicyOutput)(nil), // 3: eolymp.universe.DescribePolicyOutput
	(*DefinePolicyInput)(nil),    // 4: eolymp.universe.DefinePolicyInput
	(*DefinePolicyOutput)(nil),   // 5: eolymp.universe.DefinePolicyOutput
	(*RemovePolicyInput)(nil),    // 6: eolymp.universe.RemovePolicyInput
	(*RemovePolicyOutput)(nil),   // 7: eolymp.universe.RemovePolicyOutput
	(*universe.Policy)(nil),      // 8: eolymp.universe.Policy
}
var file_eolymp_guardian_guardian_proto_depIdxs = []int32{
	8, // 0: eolymp.universe.ListPoliciesOutput.items:type_name -> eolymp.universe.Policy
	8, // 1: eolymp.universe.DescribePolicyOutput.policy:type_name -> eolymp.universe.Policy
	8, // 2: eolymp.universe.DefinePolicyInput.policy:type_name -> eolymp.universe.Policy
	0, // 3: eolymp.universe.Guardian.ListPolicies:input_type -> eolymp.universe.ListPoliciesInput
	2, // 4: eolymp.universe.Guardian.DescribePolicy:input_type -> eolymp.universe.DescribePolicyInput
	4, // 5: eolymp.universe.Guardian.DefinePolicy:input_type -> eolymp.universe.DefinePolicyInput
	6, // 6: eolymp.universe.Guardian.RemovePolicy:input_type -> eolymp.universe.RemovePolicyInput
	1, // 7: eolymp.universe.Guardian.ListPolicies:output_type -> eolymp.universe.ListPoliciesOutput
	3, // 8: eolymp.universe.Guardian.DescribePolicy:output_type -> eolymp.universe.DescribePolicyOutput
	5, // 9: eolymp.universe.Guardian.DefinePolicy:output_type -> eolymp.universe.DefinePolicyOutput
	7, // 10: eolymp.universe.Guardian.RemovePolicy:output_type -> eolymp.universe.RemovePolicyOutput
	7, // [7:11] is the sub-list for method output_type
	3, // [3:7] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_eolymp_guardian_guardian_proto_init() }
func file_eolymp_guardian_guardian_proto_init() {
	if File_eolymp_guardian_guardian_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_eolymp_guardian_guardian_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListPoliciesInput); i {
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
		file_eolymp_guardian_guardian_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListPoliciesOutput); i {
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
		file_eolymp_guardian_guardian_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DescribePolicyInput); i {
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
		file_eolymp_guardian_guardian_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DescribePolicyOutput); i {
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
		file_eolymp_guardian_guardian_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DefinePolicyInput); i {
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
		file_eolymp_guardian_guardian_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DefinePolicyOutput); i {
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
		file_eolymp_guardian_guardian_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemovePolicyInput); i {
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
		file_eolymp_guardian_guardian_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemovePolicyOutput); i {
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
			RawDescriptor: file_eolymp_guardian_guardian_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_eolymp_guardian_guardian_proto_goTypes,
		DependencyIndexes: file_eolymp_guardian_guardian_proto_depIdxs,
		MessageInfos:      file_eolymp_guardian_guardian_proto_msgTypes,
	}.Build()
	File_eolymp_guardian_guardian_proto = out.File
	file_eolymp_guardian_guardian_proto_rawDesc = nil
	file_eolymp_guardian_guardian_proto_goTypes = nil
	file_eolymp_guardian_guardian_proto_depIdxs = nil
}
