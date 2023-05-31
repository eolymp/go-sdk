// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.18.1
// source: eolymp/l10n/project_service.proto

package l10n

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

type ListProjectsInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// pagination
	Offset int32 `protobuf:"varint,10,opt,name=offset,proto3" json:"offset,omitempty"`
	Size   int32 `protobuf:"varint,11,opt,name=size,proto3" json:"size,omitempty"`
	// data filters
	Filters *ListProjectsInput_Filter `protobuf:"bytes,40,opt,name=filters,proto3" json:"filters,omitempty"`
}

func (x *ListProjectsInput) Reset() {
	*x = ListProjectsInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_l10n_project_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListProjectsInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListProjectsInput) ProtoMessage() {}

func (x *ListProjectsInput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_l10n_project_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListProjectsInput.ProtoReflect.Descriptor instead.
func (*ListProjectsInput) Descriptor() ([]byte, []int) {
	return file_eolymp_l10n_project_service_proto_rawDescGZIP(), []int{0}
}

func (x *ListProjectsInput) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *ListProjectsInput) GetSize() int32 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *ListProjectsInput) GetFilters() *ListProjectsInput_Filter {
	if x != nil {
		return x.Filters
	}
	return nil
}

type ListProjectsOutput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Total int32      `protobuf:"varint,1,opt,name=total,proto3" json:"total,omitempty"`
	Items []*Project `protobuf:"bytes,2,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *ListProjectsOutput) Reset() {
	*x = ListProjectsOutput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_l10n_project_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListProjectsOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListProjectsOutput) ProtoMessage() {}

func (x *ListProjectsOutput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_l10n_project_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListProjectsOutput.ProtoReflect.Descriptor instead.
func (*ListProjectsOutput) Descriptor() ([]byte, []int) {
	return file_eolymp_l10n_project_service_proto_rawDescGZIP(), []int{1}
}

func (x *ListProjectsOutput) GetTotal() int32 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *ListProjectsOutput) GetItems() []*Project {
	if x != nil {
		return x.Items
	}
	return nil
}

type DescribeProjectInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProjectId string `protobuf:"bytes,1,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
}

func (x *DescribeProjectInput) Reset() {
	*x = DescribeProjectInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_l10n_project_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DescribeProjectInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescribeProjectInput) ProtoMessage() {}

func (x *DescribeProjectInput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_l10n_project_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DescribeProjectInput.ProtoReflect.Descriptor instead.
func (*DescribeProjectInput) Descriptor() ([]byte, []int) {
	return file_eolymp_l10n_project_service_proto_rawDescGZIP(), []int{2}
}

func (x *DescribeProjectInput) GetProjectId() string {
	if x != nil {
		return x.ProjectId
	}
	return ""
}

type DescribeProjectOutput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Project *Project `protobuf:"bytes,1,opt,name=project,proto3" json:"project,omitempty"`
}

func (x *DescribeProjectOutput) Reset() {
	*x = DescribeProjectOutput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_l10n_project_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DescribeProjectOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescribeProjectOutput) ProtoMessage() {}

func (x *DescribeProjectOutput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_l10n_project_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DescribeProjectOutput.ProtoReflect.Descriptor instead.
func (*DescribeProjectOutput) Descriptor() ([]byte, []int) {
	return file_eolymp_l10n_project_service_proto_rawDescGZIP(), []int{3}
}

func (x *DescribeProjectOutput) GetProject() *Project {
	if x != nil {
		return x.Project
	}
	return nil
}

type ListProjectsInput_Filter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id []*wellknown.ExpressionID `protobuf:"bytes,1,rep,name=id,proto3" json:"id,omitempty"`
}

func (x *ListProjectsInput_Filter) Reset() {
	*x = ListProjectsInput_Filter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_l10n_project_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListProjectsInput_Filter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListProjectsInput_Filter) ProtoMessage() {}

func (x *ListProjectsInput_Filter) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_l10n_project_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListProjectsInput_Filter.ProtoReflect.Descriptor instead.
func (*ListProjectsInput_Filter) Descriptor() ([]byte, []int) {
	return file_eolymp_l10n_project_service_proto_rawDescGZIP(), []int{0, 0}
}

func (x *ListProjectsInput_Filter) GetId() []*wellknown.ExpressionID {
	if x != nil {
		return x.Id
	}
	return nil
}

var File_eolymp_l10n_project_service_proto protoreflect.FileDescriptor

var file_eolymp_l10n_project_service_proto_rawDesc = []byte{
	0x0a, 0x21, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x6c, 0x31, 0x30, 0x6e, 0x2f, 0x70, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x6c, 0x31, 0x30, 0x6e,
	0x1a, 0x1d, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x22, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2f, 0x72, 0x61, 0x74, 0x65, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x6c, 0x31, 0x30, 0x6e,
	0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x21,
	0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x77, 0x65, 0x6c, 0x6c, 0x6b, 0x6e, 0x6f, 0x77, 0x6e,
	0x2f, 0x65, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xba, 0x01, 0x0a, 0x11, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x73, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65,
	0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12,
	0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x73,
	0x69, 0x7a, 0x65, 0x12, 0x3f, 0x0a, 0x07, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x18, 0x28,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x6c, 0x31,
	0x30, 0x6e, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x49,
	0x6e, 0x70, 0x75, 0x74, 0x2e, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x52, 0x07, 0x66, 0x69, 0x6c,
	0x74, 0x65, 0x72, 0x73, 0x1a, 0x38, 0x0a, 0x06, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x2e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x65, 0x6f, 0x6c,
	0x79, 0x6d, 0x70, 0x2e, 0x77, 0x65, 0x6c, 0x6c, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x2e, 0x45, 0x78,
	0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x52, 0x02, 0x69, 0x64, 0x22, 0x56,
	0x0a, 0x12, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x4f, 0x75,
	0x74, 0x70, 0x75, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x2a, 0x0a, 0x05, 0x69, 0x74,
	0x65, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x65, 0x6f, 0x6c, 0x79,
	0x6d, 0x70, 0x2e, 0x6c, 0x31, 0x30, 0x6e, 0x2e, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52,
	0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x22, 0x35, 0x0a, 0x14, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x62, 0x65, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x1d,
	0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x22, 0x47, 0x0a,
	0x15, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x12, 0x2e, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70,
	0x2e, 0x6c, 0x31, 0x30, 0x6e, 0x2e, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x07, 0x70,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x32, 0xc0, 0x02, 0x0a, 0x0e, 0x50, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x8a, 0x01, 0x0a, 0x0c, 0x4c, 0x69,
	0x73, 0x74, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x12, 0x1e, 0x2e, 0x65, 0x6f, 0x6c,
	0x79, 0x6d, 0x70, 0x2e, 0x6c, 0x31, 0x30, 0x6e, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x73, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x1f, 0x2e, 0x65, 0x6f, 0x6c,
	0x79, 0x6d, 0x70, 0x2e, 0x6c, 0x31, 0x30, 0x6e, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x73, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x22, 0x39, 0xea, 0xe2, 0x0a,
	0x0b, 0xf5, 0xe2, 0x0a, 0x00, 0x00, 0xa0, 0x40, 0xf8, 0xe2, 0x0a, 0x32, 0x82, 0xe3, 0x0a, 0x15,
	0x8a, 0xe3, 0x0a, 0x11, 0x6c, 0x31, 0x30, 0x6e, 0x3a, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x3a, 0x72, 0x65, 0x61, 0x64, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0b, 0x12, 0x09, 0x2f, 0x70, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x12, 0xa0, 0x01, 0x0a, 0x0f, 0x44, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x62, 0x65, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x21, 0x2e, 0x65, 0x6f, 0x6c,
	0x79, 0x6d, 0x70, 0x2e, 0x6c, 0x31, 0x30, 0x6e, 0x2e, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62,
	0x65, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x22, 0x2e,
	0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x6c, 0x31, 0x30, 0x6e, 0x2e, 0x44, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x62, 0x65, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x4f, 0x75, 0x74, 0x70, 0x75,
	0x74, 0x22, 0x46, 0xea, 0xe2, 0x0a, 0x0b, 0xf5, 0xe2, 0x0a, 0x00, 0x00, 0xa0, 0x40, 0xf8, 0xe2,
	0x0a, 0x32, 0x82, 0xe3, 0x0a, 0x15, 0x8a, 0xe3, 0x0a, 0x11, 0x6c, 0x31, 0x30, 0x6e, 0x3a, 0x70,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x3a, 0x72, 0x65, 0x61, 0x64, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x18, 0x12, 0x16, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x7b, 0x70, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x7d, 0x42, 0x2b, 0x5a, 0x29, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x67,
	0x6f, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x6c, 0x31, 0x30,
	0x6e, 0x3b, 0x6c, 0x31, 0x30, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_eolymp_l10n_project_service_proto_rawDescOnce sync.Once
	file_eolymp_l10n_project_service_proto_rawDescData = file_eolymp_l10n_project_service_proto_rawDesc
)

func file_eolymp_l10n_project_service_proto_rawDescGZIP() []byte {
	file_eolymp_l10n_project_service_proto_rawDescOnce.Do(func() {
		file_eolymp_l10n_project_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_eolymp_l10n_project_service_proto_rawDescData)
	})
	return file_eolymp_l10n_project_service_proto_rawDescData
}

var file_eolymp_l10n_project_service_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_eolymp_l10n_project_service_proto_goTypes = []interface{}{
	(*ListProjectsInput)(nil),        // 0: eolymp.l10n.ListProjectsInput
	(*ListProjectsOutput)(nil),       // 1: eolymp.l10n.ListProjectsOutput
	(*DescribeProjectInput)(nil),     // 2: eolymp.l10n.DescribeProjectInput
	(*DescribeProjectOutput)(nil),    // 3: eolymp.l10n.DescribeProjectOutput
	(*ListProjectsInput_Filter)(nil), // 4: eolymp.l10n.ListProjectsInput.Filter
	(*Project)(nil),                  // 5: eolymp.l10n.Project
	(*wellknown.ExpressionID)(nil),   // 6: eolymp.wellknown.ExpressionID
}
var file_eolymp_l10n_project_service_proto_depIdxs = []int32{
	4, // 0: eolymp.l10n.ListProjectsInput.filters:type_name -> eolymp.l10n.ListProjectsInput.Filter
	5, // 1: eolymp.l10n.ListProjectsOutput.items:type_name -> eolymp.l10n.Project
	5, // 2: eolymp.l10n.DescribeProjectOutput.project:type_name -> eolymp.l10n.Project
	6, // 3: eolymp.l10n.ListProjectsInput.Filter.id:type_name -> eolymp.wellknown.ExpressionID
	0, // 4: eolymp.l10n.ProjectService.ListProjects:input_type -> eolymp.l10n.ListProjectsInput
	2, // 5: eolymp.l10n.ProjectService.DescribeProject:input_type -> eolymp.l10n.DescribeProjectInput
	1, // 6: eolymp.l10n.ProjectService.ListProjects:output_type -> eolymp.l10n.ListProjectsOutput
	3, // 7: eolymp.l10n.ProjectService.DescribeProject:output_type -> eolymp.l10n.DescribeProjectOutput
	6, // [6:8] is the sub-list for method output_type
	4, // [4:6] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_eolymp_l10n_project_service_proto_init() }
func file_eolymp_l10n_project_service_proto_init() {
	if File_eolymp_l10n_project_service_proto != nil {
		return
	}
	file_eolymp_l10n_project_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_eolymp_l10n_project_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListProjectsInput); i {
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
		file_eolymp_l10n_project_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListProjectsOutput); i {
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
		file_eolymp_l10n_project_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DescribeProjectInput); i {
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
		file_eolymp_l10n_project_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DescribeProjectOutput); i {
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
		file_eolymp_l10n_project_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListProjectsInput_Filter); i {
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
			RawDescriptor: file_eolymp_l10n_project_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_eolymp_l10n_project_service_proto_goTypes,
		DependencyIndexes: file_eolymp_l10n_project_service_proto_depIdxs,
		MessageInfos:      file_eolymp_l10n_project_service_proto_msgTypes,
	}.Build()
	File_eolymp_l10n_project_service_proto = out.File
	file_eolymp_l10n_project_service_proto_rawDesc = nil
	file_eolymp_l10n_project_service_proto_goTypes = nil
	file_eolymp_l10n_project_service_proto_depIdxs = nil
}