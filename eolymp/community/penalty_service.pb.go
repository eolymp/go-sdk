// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v4.24.4
// source: eolymp/community/penalty_service.proto

package community

import (
	_ "github.com/eolymp/go-sdk/eolymp/annotations"
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

type UpdatePenaltyInput_Patch int32

const (
	UpdatePenaltyInput_ALL         UpdatePenaltyInput_Patch = 0
	UpdatePenaltyInput_SUMMARY     UpdatePenaltyInput_Patch = 1
	UpdatePenaltyInput_DESCRIPTION UpdatePenaltyInput_Patch = 2
	UpdatePenaltyInput_SCOPE       UpdatePenaltyInput_Patch = 3
	UpdatePenaltyInput_EXPIRES_AT  UpdatePenaltyInput_Patch = 4
)

// Enum value maps for UpdatePenaltyInput_Patch.
var (
	UpdatePenaltyInput_Patch_name = map[int32]string{
		0: "ALL",
		1: "SUMMARY",
		2: "DESCRIPTION",
		3: "SCOPE",
		4: "EXPIRES_AT",
	}
	UpdatePenaltyInput_Patch_value = map[string]int32{
		"ALL":         0,
		"SUMMARY":     1,
		"DESCRIPTION": 2,
		"SCOPE":       3,
		"EXPIRES_AT":  4,
	}
)

func (x UpdatePenaltyInput_Patch) Enum() *UpdatePenaltyInput_Patch {
	p := new(UpdatePenaltyInput_Patch)
	*p = x
	return p
}

func (x UpdatePenaltyInput_Patch) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (UpdatePenaltyInput_Patch) Descriptor() protoreflect.EnumDescriptor {
	return file_eolymp_community_penalty_service_proto_enumTypes[0].Descriptor()
}

func (UpdatePenaltyInput_Patch) Type() protoreflect.EnumType {
	return &file_eolymp_community_penalty_service_proto_enumTypes[0]
}

func (x UpdatePenaltyInput_Patch) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use UpdatePenaltyInput_Patch.Descriptor instead.
func (UpdatePenaltyInput_Patch) EnumDescriptor() ([]byte, []int) {
	return file_eolymp_community_penalty_service_proto_rawDescGZIP(), []int{2, 0}
}

type CreatePenaltyInput struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Penalty       *Penalty               `protobuf:"bytes,1,opt,name=penalty,proto3" json:"penalty,omitempty"`
	DontNotify    bool                   `protobuf:"varint,2,opt,name=dont_notify,json=dontNotify,proto3" json:"dont_notify,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreatePenaltyInput) Reset() {
	*x = CreatePenaltyInput{}
	mi := &file_eolymp_community_penalty_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreatePenaltyInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePenaltyInput) ProtoMessage() {}

func (x *CreatePenaltyInput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_community_penalty_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePenaltyInput.ProtoReflect.Descriptor instead.
func (*CreatePenaltyInput) Descriptor() ([]byte, []int) {
	return file_eolymp_community_penalty_service_proto_rawDescGZIP(), []int{0}
}

func (x *CreatePenaltyInput) GetPenalty() *Penalty {
	if x != nil {
		return x.Penalty
	}
	return nil
}

func (x *CreatePenaltyInput) GetDontNotify() bool {
	if x != nil {
		return x.DontNotify
	}
	return false
}

type CreatePenaltyOutput struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	PenaltyId     string                 `protobuf:"bytes,1,opt,name=penalty_id,json=penaltyId,proto3" json:"penalty_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreatePenaltyOutput) Reset() {
	*x = CreatePenaltyOutput{}
	mi := &file_eolymp_community_penalty_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreatePenaltyOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePenaltyOutput) ProtoMessage() {}

func (x *CreatePenaltyOutput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_community_penalty_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePenaltyOutput.ProtoReflect.Descriptor instead.
func (*CreatePenaltyOutput) Descriptor() ([]byte, []int) {
	return file_eolymp_community_penalty_service_proto_rawDescGZIP(), []int{1}
}

func (x *CreatePenaltyOutput) GetPenaltyId() string {
	if x != nil {
		return x.PenaltyId
	}
	return ""
}

type UpdatePenaltyInput struct {
	state         protoimpl.MessageState     `protogen:"open.v1"`
	Patch         []UpdatePenaltyInput_Patch `protobuf:"varint,3,rep,packed,name=patch,proto3,enum=eolymp.community.UpdatePenaltyInput_Patch" json:"patch,omitempty"`
	PenaltyId     string                     `protobuf:"bytes,1,opt,name=penalty_id,json=penaltyId,proto3" json:"penalty_id,omitempty"`
	Penalty       *Penalty                   `protobuf:"bytes,2,opt,name=penalty,proto3" json:"penalty,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdatePenaltyInput) Reset() {
	*x = UpdatePenaltyInput{}
	mi := &file_eolymp_community_penalty_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdatePenaltyInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdatePenaltyInput) ProtoMessage() {}

func (x *UpdatePenaltyInput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_community_penalty_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdatePenaltyInput.ProtoReflect.Descriptor instead.
func (*UpdatePenaltyInput) Descriptor() ([]byte, []int) {
	return file_eolymp_community_penalty_service_proto_rawDescGZIP(), []int{2}
}

func (x *UpdatePenaltyInput) GetPatch() []UpdatePenaltyInput_Patch {
	if x != nil {
		return x.Patch
	}
	return nil
}

func (x *UpdatePenaltyInput) GetPenaltyId() string {
	if x != nil {
		return x.PenaltyId
	}
	return ""
}

func (x *UpdatePenaltyInput) GetPenalty() *Penalty {
	if x != nil {
		return x.Penalty
	}
	return nil
}

type UpdatePenaltyOutput struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdatePenaltyOutput) Reset() {
	*x = UpdatePenaltyOutput{}
	mi := &file_eolymp_community_penalty_service_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdatePenaltyOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdatePenaltyOutput) ProtoMessage() {}

func (x *UpdatePenaltyOutput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_community_penalty_service_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdatePenaltyOutput.ProtoReflect.Descriptor instead.
func (*UpdatePenaltyOutput) Descriptor() ([]byte, []int) {
	return file_eolymp_community_penalty_service_proto_rawDescGZIP(), []int{3}
}

type DeletePenaltyInput struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	PenaltyId     string                 `protobuf:"bytes,1,opt,name=penalty_id,json=penaltyId,proto3" json:"penalty_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeletePenaltyInput) Reset() {
	*x = DeletePenaltyInput{}
	mi := &file_eolymp_community_penalty_service_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeletePenaltyInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeletePenaltyInput) ProtoMessage() {}

func (x *DeletePenaltyInput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_community_penalty_service_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeletePenaltyInput.ProtoReflect.Descriptor instead.
func (*DeletePenaltyInput) Descriptor() ([]byte, []int) {
	return file_eolymp_community_penalty_service_proto_rawDescGZIP(), []int{4}
}

func (x *DeletePenaltyInput) GetPenaltyId() string {
	if x != nil {
		return x.PenaltyId
	}
	return ""
}

type DeletePenaltyOutput struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeletePenaltyOutput) Reset() {
	*x = DeletePenaltyOutput{}
	mi := &file_eolymp_community_penalty_service_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeletePenaltyOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeletePenaltyOutput) ProtoMessage() {}

func (x *DeletePenaltyOutput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_community_penalty_service_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeletePenaltyOutput.ProtoReflect.Descriptor instead.
func (*DeletePenaltyOutput) Descriptor() ([]byte, []int) {
	return file_eolymp_community_penalty_service_proto_rawDescGZIP(), []int{5}
}

type DescribePenaltyInput struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	PenaltyId     string                 `protobuf:"bytes,1,opt,name=penalty_id,json=penaltyId,proto3" json:"penalty_id,omitempty"`
	Extra         []Penalty_Extra        `protobuf:"varint,1123,rep,packed,name=extra,proto3,enum=eolymp.community.Penalty_Extra" json:"extra,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DescribePenaltyInput) Reset() {
	*x = DescribePenaltyInput{}
	mi := &file_eolymp_community_penalty_service_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DescribePenaltyInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescribePenaltyInput) ProtoMessage() {}

func (x *DescribePenaltyInput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_community_penalty_service_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DescribePenaltyInput.ProtoReflect.Descriptor instead.
func (*DescribePenaltyInput) Descriptor() ([]byte, []int) {
	return file_eolymp_community_penalty_service_proto_rawDescGZIP(), []int{6}
}

func (x *DescribePenaltyInput) GetPenaltyId() string {
	if x != nil {
		return x.PenaltyId
	}
	return ""
}

func (x *DescribePenaltyInput) GetExtra() []Penalty_Extra {
	if x != nil {
		return x.Extra
	}
	return nil
}

type DescribePenaltyOutput struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Penalty       *Penalty               `protobuf:"bytes,1,opt,name=penalty,proto3" json:"penalty,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DescribePenaltyOutput) Reset() {
	*x = DescribePenaltyOutput{}
	mi := &file_eolymp_community_penalty_service_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DescribePenaltyOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescribePenaltyOutput) ProtoMessage() {}

func (x *DescribePenaltyOutput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_community_penalty_service_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DescribePenaltyOutput.ProtoReflect.Descriptor instead.
func (*DescribePenaltyOutput) Descriptor() ([]byte, []int) {
	return file_eolymp_community_penalty_service_proto_rawDescGZIP(), []int{7}
}

func (x *DescribePenaltyOutput) GetPenalty() *Penalty {
	if x != nil {
		return x.Penalty
	}
	return nil
}

type ListPenaltiesInput struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Offset        int32                  `protobuf:"varint,10,opt,name=offset,proto3" json:"offset,omitempty"`
	Size          int32                  `protobuf:"varint,11,opt,name=size,proto3" json:"size,omitempty"`
	Extra         []Penalty_Extra        `protobuf:"varint,1123,rep,packed,name=extra,proto3,enum=eolymp.community.Penalty_Extra" json:"extra,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListPenaltiesInput) Reset() {
	*x = ListPenaltiesInput{}
	mi := &file_eolymp_community_penalty_service_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListPenaltiesInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListPenaltiesInput) ProtoMessage() {}

func (x *ListPenaltiesInput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_community_penalty_service_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListPenaltiesInput.ProtoReflect.Descriptor instead.
func (*ListPenaltiesInput) Descriptor() ([]byte, []int) {
	return file_eolymp_community_penalty_service_proto_rawDescGZIP(), []int{8}
}

func (x *ListPenaltiesInput) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *ListPenaltiesInput) GetSize() int32 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *ListPenaltiesInput) GetExtra() []Penalty_Extra {
	if x != nil {
		return x.Extra
	}
	return nil
}

type ListPenaltiesOutput struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Total         int32                  `protobuf:"varint,1,opt,name=total,proto3" json:"total,omitempty"`
	Items         []*Penalty             `protobuf:"bytes,2,rep,name=items,proto3" json:"items,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListPenaltiesOutput) Reset() {
	*x = ListPenaltiesOutput{}
	mi := &file_eolymp_community_penalty_service_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListPenaltiesOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListPenaltiesOutput) ProtoMessage() {}

func (x *ListPenaltiesOutput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_community_penalty_service_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListPenaltiesOutput.ProtoReflect.Descriptor instead.
func (*ListPenaltiesOutput) Descriptor() ([]byte, []int) {
	return file_eolymp_community_penalty_service_proto_rawDescGZIP(), []int{9}
}

func (x *ListPenaltiesOutput) GetTotal() int32 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *ListPenaltiesOutput) GetItems() []*Penalty {
	if x != nil {
		return x.Items
	}
	return nil
}

var File_eolymp_community_penalty_service_proto protoreflect.FileDescriptor

var file_eolymp_community_penalty_service_proto_rawDesc = string([]byte{
	0x0a, 0x26, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69,
	0x74, 0x79, 0x2f, 0x70, 0x65, 0x6e, 0x61, 0x6c, 0x74, 0x79, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x1a, 0x1d, 0x65, 0x6f, 0x6c, 0x79,
	0x6d, 0x70, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x68,
	0x74, 0x74, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x22, 0x65, 0x6f, 0x6c, 0x79, 0x6d,
	0x70, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x72, 0x61,
	0x74, 0x65, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x65,
	0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2f, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x65,
	0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x2f,
	0x70, 0x65, 0x6e, 0x61, 0x6c, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x6a, 0x0a,
	0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x65, 0x6e, 0x61, 0x6c, 0x74, 0x79, 0x49, 0x6e,
	0x70, 0x75, 0x74, 0x12, 0x33, 0x0a, 0x07, 0x70, 0x65, 0x6e, 0x61, 0x6c, 0x74, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x2e, 0x50, 0x65, 0x6e, 0x61, 0x6c, 0x74, 0x79, 0x52,
	0x07, 0x70, 0x65, 0x6e, 0x61, 0x6c, 0x74, 0x79, 0x12, 0x1f, 0x0a, 0x0b, 0x64, 0x6f, 0x6e, 0x74,
	0x5f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x64,
	0x6f, 0x6e, 0x74, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x22, 0x34, 0x0a, 0x13, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x50, 0x65, 0x6e, 0x61, 0x6c, 0x74, 0x79, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74,
	0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x65, 0x6e, 0x61, 0x6c, 0x74, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x65, 0x6e, 0x61, 0x6c, 0x74, 0x79, 0x49, 0x64, 0x22,
	0xf5, 0x01, 0x0a, 0x12, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x65, 0x6e, 0x61, 0x6c, 0x74,
	0x79, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x40, 0x0a, 0x05, 0x70, 0x61, 0x74, 0x63, 0x68, 0x18,
	0x03, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x2a, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50,
	0x65, 0x6e, 0x61, 0x6c, 0x74, 0x79, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x2e, 0x50, 0x61, 0x74, 0x63,
	0x68, 0x52, 0x05, 0x70, 0x61, 0x74, 0x63, 0x68, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x65, 0x6e, 0x61,
	0x6c, 0x74, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x65,
	0x6e, 0x61, 0x6c, 0x74, 0x79, 0x49, 0x64, 0x12, 0x33, 0x0a, 0x07, 0x70, 0x65, 0x6e, 0x61, 0x6c,
	0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d,
	0x70, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x2e, 0x50, 0x65, 0x6e, 0x61,
	0x6c, 0x74, 0x79, 0x52, 0x07, 0x70, 0x65, 0x6e, 0x61, 0x6c, 0x74, 0x79, 0x22, 0x49, 0x0a, 0x05,
	0x50, 0x61, 0x74, 0x63, 0x68, 0x12, 0x07, 0x0a, 0x03, 0x41, 0x4c, 0x4c, 0x10, 0x00, 0x12, 0x0b,
	0x0a, 0x07, 0x53, 0x55, 0x4d, 0x4d, 0x41, 0x52, 0x59, 0x10, 0x01, 0x12, 0x0f, 0x0a, 0x0b, 0x44,
	0x45, 0x53, 0x43, 0x52, 0x49, 0x50, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x02, 0x12, 0x09, 0x0a, 0x05,
	0x53, 0x43, 0x4f, 0x50, 0x45, 0x10, 0x03, 0x12, 0x0e, 0x0a, 0x0a, 0x45, 0x58, 0x50, 0x49, 0x52,
	0x45, 0x53, 0x5f, 0x41, 0x54, 0x10, 0x04, 0x22, 0x15, 0x0a, 0x13, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x50, 0x65, 0x6e, 0x61, 0x6c, 0x74, 0x79, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x22, 0x33,
	0x0a, 0x12, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x65, 0x6e, 0x61, 0x6c, 0x74, 0x79, 0x49,
	0x6e, 0x70, 0x75, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x65, 0x6e, 0x61, 0x6c, 0x74, 0x79, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x65, 0x6e, 0x61, 0x6c, 0x74,
	0x79, 0x49, 0x64, 0x22, 0x15, 0x0a, 0x13, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x65, 0x6e,
	0x61, 0x6c, 0x74, 0x79, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x22, 0x6d, 0x0a, 0x14, 0x44, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x50, 0x65, 0x6e, 0x61, 0x6c, 0x74, 0x79, 0x49, 0x6e, 0x70,
	0x75, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x65, 0x6e, 0x61, 0x6c, 0x74, 0x79, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x65, 0x6e, 0x61, 0x6c, 0x74, 0x79, 0x49,
	0x64, 0x12, 0x36, 0x0a, 0x05, 0x65, 0x78, 0x74, 0x72, 0x61, 0x18, 0xe3, 0x08, 0x20, 0x03, 0x28,
	0x0e, 0x32, 0x1f, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x75,
	0x6e, 0x69, 0x74, 0x79, 0x2e, 0x50, 0x65, 0x6e, 0x61, 0x6c, 0x74, 0x79, 0x2e, 0x45, 0x78, 0x74,
	0x72, 0x61, 0x52, 0x05, 0x65, 0x78, 0x74, 0x72, 0x61, 0x22, 0x4c, 0x0a, 0x15, 0x44, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x62, 0x65, 0x50, 0x65, 0x6e, 0x61, 0x6c, 0x74, 0x79, 0x4f, 0x75, 0x74, 0x70,
	0x75, 0x74, 0x12, 0x33, 0x0a, 0x07, 0x70, 0x65, 0x6e, 0x61, 0x6c, 0x74, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x2e, 0x50, 0x65, 0x6e, 0x61, 0x6c, 0x74, 0x79, 0x52, 0x07,
	0x70, 0x65, 0x6e, 0x61, 0x6c, 0x74, 0x79, 0x22, 0x78, 0x0a, 0x12, 0x4c, 0x69, 0x73, 0x74, 0x50,
	0x65, 0x6e, 0x61, 0x6c, 0x74, 0x69, 0x65, 0x73, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x16, 0x0a,
	0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6f,
	0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x0b, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x36, 0x0a, 0x05, 0x65, 0x78, 0x74,
	0x72, 0x61, 0x18, 0xe3, 0x08, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x1f, 0x2e, 0x65, 0x6f, 0x6c, 0x79,
	0x6d, 0x70, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x2e, 0x50, 0x65, 0x6e,
	0x61, 0x6c, 0x74, 0x79, 0x2e, 0x45, 0x78, 0x74, 0x72, 0x61, 0x52, 0x05, 0x65, 0x78, 0x74, 0x72,
	0x61, 0x22, 0x5c, 0x0a, 0x13, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x65, 0x6e, 0x61, 0x6c, 0x74, 0x69,
	0x65, 0x73, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61,
	0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x2f,
	0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e,
	0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79,
	0x2e, 0x50, 0x65, 0x6e, 0x61, 0x6c, 0x74, 0x79, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x32,
	0xdb, 0x06, 0x0a, 0x0e, 0x50, 0x65, 0x6e, 0x61, 0x6c, 0x74, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x9d, 0x01, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x65, 0x6e,
	0x61, 0x6c, 0x74, 0x79, 0x12, 0x24, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x65,
	0x6e, 0x61, 0x6c, 0x74, 0x79, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x25, 0x2e, 0x65, 0x6f, 0x6c,
	0x79, 0x6d, 0x70, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x50, 0x65, 0x6e, 0x61, 0x6c, 0x74, 0x79, 0x4f, 0x75, 0x74, 0x70, 0x75,
	0x74, 0x22, 0x3f, 0xea, 0xe2, 0x0a, 0x0b, 0xf5, 0xe2, 0x0a, 0x00, 0x00, 0x20, 0x41, 0xf8, 0xe2,
	0x0a, 0x64, 0x82, 0xe3, 0x0a, 0x1a, 0x8a, 0xe3, 0x0a, 0x16, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e,
	0x69, 0x74, 0x79, 0x3a, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x3a, 0x77, 0x72, 0x69, 0x74, 0x65,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0c, 0x22, 0x0a, 0x2f, 0x70, 0x65, 0x6e, 0x61, 0x6c, 0x74, 0x69,
	0x65, 0x73, 0x12, 0xaa, 0x01, 0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x65, 0x6e,
	0x61, 0x6c, 0x74, 0x79, 0x12, 0x24, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x65,
	0x6e, 0x61, 0x6c, 0x74, 0x79, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x25, 0x2e, 0x65, 0x6f, 0x6c,
	0x79, 0x6d, 0x70, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x2e, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x50, 0x65, 0x6e, 0x61, 0x6c, 0x74, 0x79, 0x4f, 0x75, 0x74, 0x70, 0x75,
	0x74, 0x22, 0x4c, 0xea, 0xe2, 0x0a, 0x0b, 0xf5, 0xe2, 0x0a, 0x00, 0x00, 0xa0, 0x40, 0xf8, 0xe2,
	0x0a, 0x14, 0x82, 0xe3, 0x0a, 0x1a, 0x8a, 0xe3, 0x0a, 0x16, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e,
	0x69, 0x74, 0x79, 0x3a, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x3a, 0x77, 0x72, 0x69, 0x74, 0x65,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x19, 0x22, 0x17, 0x2f, 0x70, 0x65, 0x6e, 0x61, 0x6c, 0x74, 0x69,
	0x65, 0x73, 0x2f, 0x7b, 0x70, 0x65, 0x6e, 0x61, 0x6c, 0x74, 0x79, 0x5f, 0x69, 0x64, 0x7d, 0x12,
	0xaa, 0x01, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x65, 0x6e, 0x61, 0x6c, 0x74,
	0x79, 0x12, 0x24, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x75,
	0x6e, 0x69, 0x74, 0x79, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x65, 0x6e, 0x61, 0x6c,
	0x74, 0x79, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x25, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x50, 0x65, 0x6e, 0x61, 0x6c, 0x74, 0x79, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x22, 0x4c,
	0xea, 0xe2, 0x0a, 0x0b, 0xf5, 0xe2, 0x0a, 0x00, 0x00, 0x20, 0x41, 0xf8, 0xe2, 0x0a, 0x64, 0x82,
	0xe3, 0x0a, 0x1a, 0x8a, 0xe3, 0x0a, 0x16, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79,
	0x3a, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x3a, 0x77, 0x72, 0x69, 0x74, 0x65, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x19, 0x2a, 0x17, 0x2f, 0x70, 0x65, 0x6e, 0x61, 0x6c, 0x74, 0x69, 0x65, 0x73, 0x2f,
	0x7b, 0x70, 0x65, 0x6e, 0x61, 0x6c, 0x74, 0x79, 0x5f, 0x69, 0x64, 0x7d, 0x12, 0xaf, 0x01, 0x0a,
	0x0f, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x50, 0x65, 0x6e, 0x61, 0x6c, 0x74, 0x79,
	0x12, 0x26, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e,
	0x69, 0x74, 0x79, 0x2e, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x50, 0x65, 0x6e, 0x61,
	0x6c, 0x74, 0x79, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x27, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d,
	0x70, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x2e, 0x44, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x62, 0x65, 0x50, 0x65, 0x6e, 0x61, 0x6c, 0x74, 0x79, 0x4f, 0x75, 0x74, 0x70, 0x75,
	0x74, 0x22, 0x4b, 0xea, 0xe2, 0x0a, 0x0b, 0xf5, 0xe2, 0x0a, 0x00, 0x00, 0xa0, 0x40, 0xf8, 0xe2,
	0x0a, 0x14, 0x82, 0xe3, 0x0a, 0x19, 0x8a, 0xe3, 0x0a, 0x15, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e,
	0x69, 0x74, 0x79, 0x3a, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x3a, 0x72, 0x65, 0x61, 0x64, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x19, 0x12, 0x17, 0x2f, 0x70, 0x65, 0x6e, 0x61, 0x6c, 0x74, 0x69, 0x65,
	0x73, 0x2f, 0x7b, 0x70, 0x65, 0x6e, 0x61, 0x6c, 0x74, 0x79, 0x5f, 0x69, 0x64, 0x7d, 0x12, 0x9c,
	0x01, 0x0a, 0x0d, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x65, 0x6e, 0x61, 0x6c, 0x74, 0x69, 0x65, 0x73,
	0x12, 0x24, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e,
	0x69, 0x74, 0x79, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x65, 0x6e, 0x61, 0x6c, 0x74, 0x69, 0x65,
	0x73, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x25, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x65,
	0x6e, 0x61, 0x6c, 0x74, 0x69, 0x65, 0x73, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x22, 0x3e, 0xea,
	0xe2, 0x0a, 0x0b, 0xf5, 0xe2, 0x0a, 0x00, 0x00, 0xa0, 0x40, 0xf8, 0xe2, 0x0a, 0x14, 0x82, 0xe3,
	0x0a, 0x19, 0x8a, 0xe3, 0x0a, 0x15, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x3a,
	0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x3a, 0x72, 0x65, 0x61, 0x64, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x0c, 0x12, 0x0a, 0x2f, 0x70, 0x65, 0x6e, 0x61, 0x6c, 0x74, 0x69, 0x65, 0x73, 0x42, 0x35, 0x5a,
	0x33, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6f, 0x6c, 0x79,
	0x6d, 0x70, 0x2f, 0x67, 0x6f, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70,
	0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x3b, 0x63, 0x6f, 0x6d, 0x6d, 0x75,
	0x6e, 0x69, 0x74, 0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_eolymp_community_penalty_service_proto_rawDescOnce sync.Once
	file_eolymp_community_penalty_service_proto_rawDescData []byte
)

func file_eolymp_community_penalty_service_proto_rawDescGZIP() []byte {
	file_eolymp_community_penalty_service_proto_rawDescOnce.Do(func() {
		file_eolymp_community_penalty_service_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_eolymp_community_penalty_service_proto_rawDesc), len(file_eolymp_community_penalty_service_proto_rawDesc)))
	})
	return file_eolymp_community_penalty_service_proto_rawDescData
}

var file_eolymp_community_penalty_service_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_eolymp_community_penalty_service_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_eolymp_community_penalty_service_proto_goTypes = []any{
	(UpdatePenaltyInput_Patch)(0), // 0: eolymp.community.UpdatePenaltyInput.Patch
	(*CreatePenaltyInput)(nil),    // 1: eolymp.community.CreatePenaltyInput
	(*CreatePenaltyOutput)(nil),   // 2: eolymp.community.CreatePenaltyOutput
	(*UpdatePenaltyInput)(nil),    // 3: eolymp.community.UpdatePenaltyInput
	(*UpdatePenaltyOutput)(nil),   // 4: eolymp.community.UpdatePenaltyOutput
	(*DeletePenaltyInput)(nil),    // 5: eolymp.community.DeletePenaltyInput
	(*DeletePenaltyOutput)(nil),   // 6: eolymp.community.DeletePenaltyOutput
	(*DescribePenaltyInput)(nil),  // 7: eolymp.community.DescribePenaltyInput
	(*DescribePenaltyOutput)(nil), // 8: eolymp.community.DescribePenaltyOutput
	(*ListPenaltiesInput)(nil),    // 9: eolymp.community.ListPenaltiesInput
	(*ListPenaltiesOutput)(nil),   // 10: eolymp.community.ListPenaltiesOutput
	(*Penalty)(nil),               // 11: eolymp.community.Penalty
	(Penalty_Extra)(0),            // 12: eolymp.community.Penalty.Extra
}
var file_eolymp_community_penalty_service_proto_depIdxs = []int32{
	11, // 0: eolymp.community.CreatePenaltyInput.penalty:type_name -> eolymp.community.Penalty
	0,  // 1: eolymp.community.UpdatePenaltyInput.patch:type_name -> eolymp.community.UpdatePenaltyInput.Patch
	11, // 2: eolymp.community.UpdatePenaltyInput.penalty:type_name -> eolymp.community.Penalty
	12, // 3: eolymp.community.DescribePenaltyInput.extra:type_name -> eolymp.community.Penalty.Extra
	11, // 4: eolymp.community.DescribePenaltyOutput.penalty:type_name -> eolymp.community.Penalty
	12, // 5: eolymp.community.ListPenaltiesInput.extra:type_name -> eolymp.community.Penalty.Extra
	11, // 6: eolymp.community.ListPenaltiesOutput.items:type_name -> eolymp.community.Penalty
	1,  // 7: eolymp.community.PenaltyService.CreatePenalty:input_type -> eolymp.community.CreatePenaltyInput
	3,  // 8: eolymp.community.PenaltyService.UpdatePenalty:input_type -> eolymp.community.UpdatePenaltyInput
	5,  // 9: eolymp.community.PenaltyService.DeletePenalty:input_type -> eolymp.community.DeletePenaltyInput
	7,  // 10: eolymp.community.PenaltyService.DescribePenalty:input_type -> eolymp.community.DescribePenaltyInput
	9,  // 11: eolymp.community.PenaltyService.ListPenalties:input_type -> eolymp.community.ListPenaltiesInput
	2,  // 12: eolymp.community.PenaltyService.CreatePenalty:output_type -> eolymp.community.CreatePenaltyOutput
	4,  // 13: eolymp.community.PenaltyService.UpdatePenalty:output_type -> eolymp.community.UpdatePenaltyOutput
	6,  // 14: eolymp.community.PenaltyService.DeletePenalty:output_type -> eolymp.community.DeletePenaltyOutput
	8,  // 15: eolymp.community.PenaltyService.DescribePenalty:output_type -> eolymp.community.DescribePenaltyOutput
	10, // 16: eolymp.community.PenaltyService.ListPenalties:output_type -> eolymp.community.ListPenaltiesOutput
	12, // [12:17] is the sub-list for method output_type
	7,  // [7:12] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_eolymp_community_penalty_service_proto_init() }
func file_eolymp_community_penalty_service_proto_init() {
	if File_eolymp_community_penalty_service_proto != nil {
		return
	}
	file_eolymp_community_penalty_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_eolymp_community_penalty_service_proto_rawDesc), len(file_eolymp_community_penalty_service_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_eolymp_community_penalty_service_proto_goTypes,
		DependencyIndexes: file_eolymp_community_penalty_service_proto_depIdxs,
		EnumInfos:         file_eolymp_community_penalty_service_proto_enumTypes,
		MessageInfos:      file_eolymp_community_penalty_service_proto_msgTypes,
	}.Build()
	File_eolymp_community_penalty_service_proto = out.File
	file_eolymp_community_penalty_service_proto_goTypes = nil
	file_eolymp_community_penalty_service_proto_depIdxs = nil
}
