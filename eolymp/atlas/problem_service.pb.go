// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.21.12
// source: eolymp/atlas/problem_service.proto

package atlas

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

type ListVersionsInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProblemId string `protobuf:"bytes,1,opt,name=problem_id,json=problemId,proto3" json:"problem_id,omitempty"`
	// pagination
	Offset int32 `protobuf:"varint,10,opt,name=offset,proto3" json:"offset,omitempty"`
	Size   int32 `protobuf:"varint,11,opt,name=size,proto3" json:"size,omitempty"`
	// data filters
	Filters *ListVersionsInput_Filter `protobuf:"bytes,40,opt,name=filters,proto3" json:"filters,omitempty"`
}

func (x *ListVersionsInput) Reset() {
	*x = ListVersionsInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_atlas_problem_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListVersionsInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListVersionsInput) ProtoMessage() {}

func (x *ListVersionsInput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_atlas_problem_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListVersionsInput.ProtoReflect.Descriptor instead.
func (*ListVersionsInput) Descriptor() ([]byte, []int) {
	return file_eolymp_atlas_problem_service_proto_rawDescGZIP(), []int{0}
}

func (x *ListVersionsInput) GetProblemId() string {
	if x != nil {
		return x.ProblemId
	}
	return ""
}

func (x *ListVersionsInput) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *ListVersionsInput) GetSize() int32 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *ListVersionsInput) GetFilters() *ListVersionsInput_Filter {
	if x != nil {
		return x.Filters
	}
	return nil
}

type ListVersionsOutput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Total int32      `protobuf:"varint,1,opt,name=total,proto3" json:"total,omitempty"`
	Items []*Version `protobuf:"bytes,2,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *ListVersionsOutput) Reset() {
	*x = ListVersionsOutput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_atlas_problem_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListVersionsOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListVersionsOutput) ProtoMessage() {}

func (x *ListVersionsOutput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_atlas_problem_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListVersionsOutput.ProtoReflect.Descriptor instead.
func (*ListVersionsOutput) Descriptor() ([]byte, []int) {
	return file_eolymp_atlas_problem_service_proto_rawDescGZIP(), []int{1}
}

func (x *ListVersionsOutput) GetTotal() int32 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *ListVersionsOutput) GetItems() []*Version {
	if x != nil {
		return x.Items
	}
	return nil
}

type ListVersionsInput_Filter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Number     []*wellknown.ExpressionInt       `protobuf:"bytes,1,rep,name=number,proto3" json:"number,omitempty"`
	CreatedBy  []*wellknown.ExpressionID        `protobuf:"bytes,2,rep,name=created_by,json=createdBy,proto3" json:"created_by,omitempty"`
	CreatedAt  []*wellknown.ExpressionTimestamp `protobuf:"bytes,3,rep,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	ChangeOp   []*wellknown.ExpressionEnum      `protobuf:"bytes,4,rep,name=change_op,json=changeOp,proto3" json:"change_op,omitempty"`
	ChangePath []*wellknown.ExpressionString    `protobuf:"bytes,5,rep,name=change_path,json=changePath,proto3" json:"change_path,omitempty"`
}

func (x *ListVersionsInput_Filter) Reset() {
	*x = ListVersionsInput_Filter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_atlas_problem_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListVersionsInput_Filter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListVersionsInput_Filter) ProtoMessage() {}

func (x *ListVersionsInput_Filter) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_atlas_problem_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListVersionsInput_Filter.ProtoReflect.Descriptor instead.
func (*ListVersionsInput_Filter) Descriptor() ([]byte, []int) {
	return file_eolymp_atlas_problem_service_proto_rawDescGZIP(), []int{0, 0}
}

func (x *ListVersionsInput_Filter) GetNumber() []*wellknown.ExpressionInt {
	if x != nil {
		return x.Number
	}
	return nil
}

func (x *ListVersionsInput_Filter) GetCreatedBy() []*wellknown.ExpressionID {
	if x != nil {
		return x.CreatedBy
	}
	return nil
}

func (x *ListVersionsInput_Filter) GetCreatedAt() []*wellknown.ExpressionTimestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *ListVersionsInput_Filter) GetChangeOp() []*wellknown.ExpressionEnum {
	if x != nil {
		return x.ChangeOp
	}
	return nil
}

func (x *ListVersionsInput_Filter) GetChangePath() []*wellknown.ExpressionString {
	if x != nil {
		return x.ChangePath
	}
	return nil
}

var File_eolymp_atlas_problem_service_proto protoreflect.FileDescriptor

var file_eolymp_atlas_problem_service_proto_rawDesc = []byte{
	0x0a, 0x22, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x61, 0x74, 0x6c, 0x61, 0x73, 0x2f, 0x70,
	0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x61, 0x74, 0x6c,
	0x61, 0x73, 0x1a, 0x1d, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x22, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x72, 0x61, 0x74, 0x65, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x61, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x22, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x61, 0x74,
	0x6c, 0x61, 0x73, 0x2f, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1a, 0x65, 0x6f, 0x6c, 0x79, 0x6d,
	0x70, 0x2f, 0x61, 0x74, 0x6c, 0x61, 0x73, 0x2f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x21, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x77, 0x65,
	0x6c, 0x6c, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x2f, 0x65, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xed, 0x03, 0x0a, 0x11, 0x4c, 0x69, 0x73,
	0x74, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x1d,
	0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x49, 0x64, 0x12, 0x16, 0x0a,
	0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6f,
	0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x0b, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x40, 0x0a, 0x07, 0x66, 0x69, 0x6c,
	0x74, 0x65, 0x72, 0x73, 0x18, 0x28, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x65, 0x6f, 0x6c,
	0x79, 0x6d, 0x70, 0x2e, 0x61, 0x74, 0x6c, 0x61, 0x73, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x56, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x2e, 0x46, 0x69, 0x6c, 0x74,
	0x65, 0x72, 0x52, 0x07, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x1a, 0xca, 0x02, 0x0a, 0x06,
	0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x37, 0x0a, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e,
	0x77, 0x65, 0x6c, 0x6c, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x2e, 0x45, 0x78, 0x70, 0x72, 0x65, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x74, 0x52, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12,
	0x3d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x62, 0x79, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x77, 0x65, 0x6c,
	0x6c, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x2e, 0x45, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x49, 0x44, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x12, 0x44,
	0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x03, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x25, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x77, 0x65, 0x6c, 0x6c,
	0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x2e, 0x45, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x12, 0x3d, 0x0a, 0x09, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x5f, 0x6f,
	0x70, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70,
	0x2e, 0x77, 0x65, 0x6c, 0x6c, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x2e, 0x45, 0x78, 0x70, 0x72, 0x65,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x45, 0x6e, 0x75, 0x6d, 0x52, 0x08, 0x63, 0x68, 0x61, 0x6e, 0x67,
	0x65, 0x4f, 0x70, 0x12, 0x43, 0x0a, 0x0b, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x5f, 0x70, 0x61,
	0x74, 0x68, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d,
	0x70, 0x2e, 0x77, 0x65, 0x6c, 0x6c, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x2e, 0x45, 0x78, 0x70, 0x72,
	0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x52, 0x0a, 0x63, 0x68,
	0x61, 0x6e, 0x67, 0x65, 0x50, 0x61, 0x74, 0x68, 0x22, 0x57, 0x0a, 0x12, 0x4c, 0x69, 0x73, 0x74,
	0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x12, 0x14,
	0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x74,
	0x6f, 0x74, 0x61, 0x6c, 0x12, 0x2b, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x61, 0x74, 0x6c,
	0x61, 0x73, 0x2e, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d,
	0x73, 0x32, 0xf3, 0x07, 0x0a, 0x0e, 0x50, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x88, 0x01, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50,
	0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x12, 0x20, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e,
	0x61, 0x74, 0x6c, 0x61, 0x73, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x62,
	0x6c, 0x65, 0x6d, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x21, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d,
	0x70, 0x2e, 0x61, 0x74, 0x6c, 0x61, 0x73, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x72,
	0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x22, 0x32, 0xea, 0xe2, 0x0a,
	0x0b, 0xf5, 0xe2, 0x0a, 0x0a, 0xd7, 0x23, 0x3e, 0xf8, 0xe2, 0x0a, 0x05, 0x82, 0xe3, 0x0a, 0x17,
	0x8a, 0xe3, 0x0a, 0x13, 0x61, 0x74, 0x6c, 0x61, 0x73, 0x3a, 0x70, 0x72, 0x6f, 0x62, 0x6c, 0x65,
	0x6d, 0x3a, 0x77, 0x72, 0x69, 0x74, 0x65, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x02, 0x2a, 0x00, 0x12,
	0x88, 0x01, 0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x62, 0x6c, 0x65,
	0x6d, 0x12, 0x20, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x61, 0x74, 0x6c, 0x61, 0x73,
	0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x49, 0x6e,
	0x70, 0x75, 0x74, 0x1a, 0x21, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x61, 0x74, 0x6c,
	0x61, 0x73, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d,
	0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x22, 0x32, 0xea, 0xe2, 0x0a, 0x0b, 0xf5, 0xe2, 0x0a, 0x00,
	0x00, 0xa0, 0x41, 0xf8, 0xe2, 0x0a, 0x64, 0x82, 0xe3, 0x0a, 0x17, 0x8a, 0xe3, 0x0a, 0x13, 0x61,
	0x74, 0x6c, 0x61, 0x73, 0x3a, 0x70, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x3a, 0x77, 0x72, 0x69,
	0x74, 0x65, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x02, 0x1a, 0x00, 0x12, 0x87, 0x01, 0x0a, 0x0b, 0x53,
	0x79, 0x6e, 0x63, 0x50, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x12, 0x1e, 0x2e, 0x65, 0x6f, 0x6c,
	0x79, 0x6d, 0x70, 0x2e, 0x61, 0x74, 0x6c, 0x61, 0x73, 0x2e, 0x53, 0x79, 0x6e, 0x63, 0x50, 0x72,
	0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x1f, 0x2e, 0x65, 0x6f, 0x6c,
	0x79, 0x6d, 0x70, 0x2e, 0x61, 0x74, 0x6c, 0x61, 0x73, 0x2e, 0x53, 0x79, 0x6e, 0x63, 0x50, 0x72,
	0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x22, 0x37, 0xea, 0xe2, 0x0a,
	0x0b, 0xf5, 0xe2, 0x0a, 0x0a, 0xd7, 0x23, 0x3e, 0xf8, 0xe2, 0x0a, 0x05, 0x82, 0xe3, 0x0a, 0x17,
	0x8a, 0xe3, 0x0a, 0x13, 0x61, 0x74, 0x6c, 0x61, 0x73, 0x3a, 0x70, 0x72, 0x6f, 0x62, 0x6c, 0x65,
	0x6d, 0x3a, 0x77, 0x72, 0x69, 0x74, 0x65, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x07, 0x22, 0x05, 0x2f,
	0x73, 0x79, 0x6e, 0x63, 0x12, 0x8d, 0x01, 0x0a, 0x0f, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62,
	0x65, 0x50, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x12, 0x22, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d,
	0x70, 0x2e, 0x61, 0x74, 0x6c, 0x61, 0x73, 0x2e, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65,
	0x50, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x23, 0x2e, 0x65,
	0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x61, 0x74, 0x6c, 0x61, 0x73, 0x2e, 0x44, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x62, 0x65, 0x50, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x4f, 0x75, 0x74, 0x70, 0x75,
	0x74, 0x22, 0x31, 0xea, 0xe2, 0x0a, 0x0b, 0xf5, 0xe2, 0x0a, 0x00, 0x00, 0xa0, 0x41, 0xf8, 0xe2,
	0x0a, 0x64, 0x82, 0xe3, 0x0a, 0x16, 0x8a, 0xe3, 0x0a, 0x12, 0x61, 0x74, 0x6c, 0x61, 0x73, 0x3a,
	0x70, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x3a, 0x72, 0x65, 0x61, 0x64, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x02, 0x12, 0x00, 0x12, 0x9c, 0x01, 0x0a, 0x10, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x56,
	0x69, 0x73, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x12, 0x23, 0x2e, 0x65, 0x6f, 0x6c, 0x79,
	0x6d, 0x70, 0x2e, 0x61, 0x74, 0x6c, 0x61, 0x73, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x56,
	0x69, 0x73, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x24,
	0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x61, 0x74, 0x6c, 0x61, 0x73, 0x2e, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x56, 0x69, 0x73, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x4f, 0x75,
	0x74, 0x70, 0x75, 0x74, 0x22, 0x3d, 0xea, 0xe2, 0x0a, 0x0b, 0xf5, 0xe2, 0x0a, 0x00, 0x00, 0xa0,
	0x41, 0xf8, 0xe2, 0x0a, 0x64, 0x82, 0xe3, 0x0a, 0x17, 0x8a, 0xe3, 0x0a, 0x13, 0x61, 0x74, 0x6c,
	0x61, 0x73, 0x3a, 0x70, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x3a, 0x77, 0x72, 0x69, 0x74, 0x65,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0d, 0x22, 0x0b, 0x2f, 0x76, 0x69, 0x73, 0x69, 0x62, 0x69, 0x6c,
	0x69, 0x74, 0x79, 0x12, 0x90, 0x01, 0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x72,
	0x69, 0x76, 0x61, 0x63, 0x79, 0x12, 0x20, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x61,
	0x74, 0x6c, 0x61, 0x73, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x72, 0x69, 0x76, 0x61,
	0x63, 0x79, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x21, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70,
	0x2e, 0x61, 0x74, 0x6c, 0x61, 0x73, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x72, 0x69,
	0x76, 0x61, 0x63, 0x79, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x22, 0x3a, 0xea, 0xe2, 0x0a, 0x0b,
	0xf5, 0xe2, 0x0a, 0x00, 0x00, 0xa0, 0x41, 0xf8, 0xe2, 0x0a, 0x64, 0x82, 0xe3, 0x0a, 0x17, 0x8a,
	0xe3, 0x0a, 0x13, 0x61, 0x74, 0x6c, 0x61, 0x73, 0x3a, 0x70, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d,
	0x3a, 0x77, 0x72, 0x69, 0x74, 0x65, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0a, 0x22, 0x08, 0x2f, 0x70,
	0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x12, 0x7f, 0x0a, 0x0c, 0x4c, 0x69, 0x73, 0x74, 0x56, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x1f, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e,
	0x61, 0x74, 0x6c, 0x61, 0x73, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x20, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70,
	0x2e, 0x61, 0x74, 0x6c, 0x61, 0x73, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x56, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x73, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x22, 0x2c, 0x82, 0xe3, 0x0a, 0x17, 0x8a,
	0xe3, 0x0a, 0x13, 0x61, 0x74, 0x6c, 0x61, 0x73, 0x3a, 0x70, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d,
	0x3a, 0x77, 0x72, 0x69, 0x74, 0x65, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0b, 0x12, 0x09, 0x2f, 0x76,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x42, 0x2d, 0x5a, 0x2b, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x67, 0x6f, 0x2d,
	0x73, 0x64, 0x6b, 0x2f, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x61, 0x74, 0x6c, 0x61, 0x73,
	0x3b, 0x61, 0x74, 0x6c, 0x61, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_eolymp_atlas_problem_service_proto_rawDescOnce sync.Once
	file_eolymp_atlas_problem_service_proto_rawDescData = file_eolymp_atlas_problem_service_proto_rawDesc
)

func file_eolymp_atlas_problem_service_proto_rawDescGZIP() []byte {
	file_eolymp_atlas_problem_service_proto_rawDescOnce.Do(func() {
		file_eolymp_atlas_problem_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_eolymp_atlas_problem_service_proto_rawDescData)
	})
	return file_eolymp_atlas_problem_service_proto_rawDescData
}

var file_eolymp_atlas_problem_service_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_eolymp_atlas_problem_service_proto_goTypes = []any{
	(*ListVersionsInput)(nil),             // 0: eolymp.atlas.ListVersionsInput
	(*ListVersionsOutput)(nil),            // 1: eolymp.atlas.ListVersionsOutput
	(*ListVersionsInput_Filter)(nil),      // 2: eolymp.atlas.ListVersionsInput.Filter
	(*Version)(nil),                       // 3: eolymp.atlas.Version
	(*wellknown.ExpressionInt)(nil),       // 4: eolymp.wellknown.ExpressionInt
	(*wellknown.ExpressionID)(nil),        // 5: eolymp.wellknown.ExpressionID
	(*wellknown.ExpressionTimestamp)(nil), // 6: eolymp.wellknown.ExpressionTimestamp
	(*wellknown.ExpressionEnum)(nil),      // 7: eolymp.wellknown.ExpressionEnum
	(*wellknown.ExpressionString)(nil),    // 8: eolymp.wellknown.ExpressionString
	(*DeleteProblemInput)(nil),            // 9: eolymp.atlas.DeleteProblemInput
	(*UpdateProblemInput)(nil),            // 10: eolymp.atlas.UpdateProblemInput
	(*SyncProblemInput)(nil),              // 11: eolymp.atlas.SyncProblemInput
	(*DescribeProblemInput)(nil),          // 12: eolymp.atlas.DescribeProblemInput
	(*UpdateVisibilityInput)(nil),         // 13: eolymp.atlas.UpdateVisibilityInput
	(*UpdatePrivacyInput)(nil),            // 14: eolymp.atlas.UpdatePrivacyInput
	(*DeleteProblemOutput)(nil),           // 15: eolymp.atlas.DeleteProblemOutput
	(*UpdateProblemOutput)(nil),           // 16: eolymp.atlas.UpdateProblemOutput
	(*SyncProblemOutput)(nil),             // 17: eolymp.atlas.SyncProblemOutput
	(*DescribeProblemOutput)(nil),         // 18: eolymp.atlas.DescribeProblemOutput
	(*UpdateVisibilityOutput)(nil),        // 19: eolymp.atlas.UpdateVisibilityOutput
	(*UpdatePrivacyOutput)(nil),           // 20: eolymp.atlas.UpdatePrivacyOutput
}
var file_eolymp_atlas_problem_service_proto_depIdxs = []int32{
	2,  // 0: eolymp.atlas.ListVersionsInput.filters:type_name -> eolymp.atlas.ListVersionsInput.Filter
	3,  // 1: eolymp.atlas.ListVersionsOutput.items:type_name -> eolymp.atlas.Version
	4,  // 2: eolymp.atlas.ListVersionsInput.Filter.number:type_name -> eolymp.wellknown.ExpressionInt
	5,  // 3: eolymp.atlas.ListVersionsInput.Filter.created_by:type_name -> eolymp.wellknown.ExpressionID
	6,  // 4: eolymp.atlas.ListVersionsInput.Filter.created_at:type_name -> eolymp.wellknown.ExpressionTimestamp
	7,  // 5: eolymp.atlas.ListVersionsInput.Filter.change_op:type_name -> eolymp.wellknown.ExpressionEnum
	8,  // 6: eolymp.atlas.ListVersionsInput.Filter.change_path:type_name -> eolymp.wellknown.ExpressionString
	9,  // 7: eolymp.atlas.ProblemService.DeleteProblem:input_type -> eolymp.atlas.DeleteProblemInput
	10, // 8: eolymp.atlas.ProblemService.UpdateProblem:input_type -> eolymp.atlas.UpdateProblemInput
	11, // 9: eolymp.atlas.ProblemService.SyncProblem:input_type -> eolymp.atlas.SyncProblemInput
	12, // 10: eolymp.atlas.ProblemService.DescribeProblem:input_type -> eolymp.atlas.DescribeProblemInput
	13, // 11: eolymp.atlas.ProblemService.UpdateVisibility:input_type -> eolymp.atlas.UpdateVisibilityInput
	14, // 12: eolymp.atlas.ProblemService.UpdatePrivacy:input_type -> eolymp.atlas.UpdatePrivacyInput
	0,  // 13: eolymp.atlas.ProblemService.ListVersions:input_type -> eolymp.atlas.ListVersionsInput
	15, // 14: eolymp.atlas.ProblemService.DeleteProblem:output_type -> eolymp.atlas.DeleteProblemOutput
	16, // 15: eolymp.atlas.ProblemService.UpdateProblem:output_type -> eolymp.atlas.UpdateProblemOutput
	17, // 16: eolymp.atlas.ProblemService.SyncProblem:output_type -> eolymp.atlas.SyncProblemOutput
	18, // 17: eolymp.atlas.ProblemService.DescribeProblem:output_type -> eolymp.atlas.DescribeProblemOutput
	19, // 18: eolymp.atlas.ProblemService.UpdateVisibility:output_type -> eolymp.atlas.UpdateVisibilityOutput
	20, // 19: eolymp.atlas.ProblemService.UpdatePrivacy:output_type -> eolymp.atlas.UpdatePrivacyOutput
	1,  // 20: eolymp.atlas.ProblemService.ListVersions:output_type -> eolymp.atlas.ListVersionsOutput
	14, // [14:21] is the sub-list for method output_type
	7,  // [7:14] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_eolymp_atlas_problem_service_proto_init() }
func file_eolymp_atlas_problem_service_proto_init() {
	if File_eolymp_atlas_problem_service_proto != nil {
		return
	}
	file_eolymp_atlas_library_service_proto_init()
	file_eolymp_atlas_version_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_eolymp_atlas_problem_service_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*ListVersionsInput); i {
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
		file_eolymp_atlas_problem_service_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*ListVersionsOutput); i {
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
		file_eolymp_atlas_problem_service_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*ListVersionsInput_Filter); i {
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
			RawDescriptor: file_eolymp_atlas_problem_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_eolymp_atlas_problem_service_proto_goTypes,
		DependencyIndexes: file_eolymp_atlas_problem_service_proto_depIdxs,
		MessageInfos:      file_eolymp_atlas_problem_service_proto_msgTypes,
	}.Build()
	File_eolymp_atlas_problem_service_proto = out.File
	file_eolymp_atlas_problem_service_proto_rawDesc = nil
	file_eolymp_atlas_problem_service_proto_goTypes = nil
	file_eolymp_atlas_problem_service_proto_depIdxs = nil
}
