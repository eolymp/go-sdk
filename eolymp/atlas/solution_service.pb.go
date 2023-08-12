// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.18.1
// source: eolymp/atlas/solution_service.proto

package atlas

import (
	_ "github.com/eolymp/go-sdk/eolymp/annotations"
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

type UpdateSolutionInput_Patch int32

const (
	UpdateSolutionInput_ALL     UpdateSolutionInput_Patch = 0
	UpdateSolutionInput_NAME    UpdateSolutionInput_Patch = 1
	UpdateSolutionInput_TYPE    UpdateSolutionInput_Patch = 2
	UpdateSolutionInput_RUNTIME UpdateSolutionInput_Patch = 3
	UpdateSolutionInput_SOURCE  UpdateSolutionInput_Patch = 4
)

// Enum value maps for UpdateSolutionInput_Patch.
var (
	UpdateSolutionInput_Patch_name = map[int32]string{
		0: "ALL",
		1: "NAME",
		2: "TYPE",
		3: "RUNTIME",
		4: "SOURCE",
	}
	UpdateSolutionInput_Patch_value = map[string]int32{
		"ALL":     0,
		"NAME":    1,
		"TYPE":    2,
		"RUNTIME": 3,
		"SOURCE":  4,
	}
)

func (x UpdateSolutionInput_Patch) Enum() *UpdateSolutionInput_Patch {
	p := new(UpdateSolutionInput_Patch)
	*p = x
	return p
}

func (x UpdateSolutionInput_Patch) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (UpdateSolutionInput_Patch) Descriptor() protoreflect.EnumDescriptor {
	return file_eolymp_atlas_solution_service_proto_enumTypes[0].Descriptor()
}

func (UpdateSolutionInput_Patch) Type() protoreflect.EnumType {
	return &file_eolymp_atlas_solution_service_proto_enumTypes[0]
}

func (x UpdateSolutionInput_Patch) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use UpdateSolutionInput_Patch.Descriptor instead.
func (UpdateSolutionInput_Patch) EnumDescriptor() ([]byte, []int) {
	return file_eolymp_atlas_solution_service_proto_rawDescGZIP(), []int{6, 0}
}

type ListSolutionsInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProblemId string `protobuf:"bytes,1,opt,name=problem_id,json=problemId,proto3" json:"problem_id,omitempty"`
	Version   uint32 `protobuf:"varint,100,opt,name=version,proto3" json:"version,omitempty"` // request data for specific problem version
}

func (x *ListSolutionsInput) Reset() {
	*x = ListSolutionsInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_atlas_solution_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListSolutionsInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListSolutionsInput) ProtoMessage() {}

func (x *ListSolutionsInput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_atlas_solution_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListSolutionsInput.ProtoReflect.Descriptor instead.
func (*ListSolutionsInput) Descriptor() ([]byte, []int) {
	return file_eolymp_atlas_solution_service_proto_rawDescGZIP(), []int{0}
}

func (x *ListSolutionsInput) GetProblemId() string {
	if x != nil {
		return x.ProblemId
	}
	return ""
}

func (x *ListSolutionsInput) GetVersion() uint32 {
	if x != nil {
		return x.Version
	}
	return 0
}

type ListSolutionsOutput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Total int32       `protobuf:"varint,1,opt,name=total,proto3" json:"total,omitempty"`
	Items []*Solution `protobuf:"bytes,2,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *ListSolutionsOutput) Reset() {
	*x = ListSolutionsOutput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_atlas_solution_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListSolutionsOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListSolutionsOutput) ProtoMessage() {}

func (x *ListSolutionsOutput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_atlas_solution_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListSolutionsOutput.ProtoReflect.Descriptor instead.
func (*ListSolutionsOutput) Descriptor() ([]byte, []int) {
	return file_eolymp_atlas_solution_service_proto_rawDescGZIP(), []int{1}
}

func (x *ListSolutionsOutput) GetTotal() int32 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *ListSolutionsOutput) GetItems() []*Solution {
	if x != nil {
		return x.Items
	}
	return nil
}

type DescribeSolutionInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProblemId  string `protobuf:"bytes,1,opt,name=problem_id,json=problemId,proto3" json:"problem_id,omitempty"`
	SolutionId string `protobuf:"bytes,2,opt,name=solution_id,json=solutionId,proto3" json:"solution_id,omitempty"`
	Version    uint32 `protobuf:"varint,100,opt,name=version,proto3" json:"version,omitempty"` // request data for specific problem version
}

func (x *DescribeSolutionInput) Reset() {
	*x = DescribeSolutionInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_atlas_solution_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DescribeSolutionInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescribeSolutionInput) ProtoMessage() {}

func (x *DescribeSolutionInput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_atlas_solution_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DescribeSolutionInput.ProtoReflect.Descriptor instead.
func (*DescribeSolutionInput) Descriptor() ([]byte, []int) {
	return file_eolymp_atlas_solution_service_proto_rawDescGZIP(), []int{2}
}

func (x *DescribeSolutionInput) GetProblemId() string {
	if x != nil {
		return x.ProblemId
	}
	return ""
}

func (x *DescribeSolutionInput) GetSolutionId() string {
	if x != nil {
		return x.SolutionId
	}
	return ""
}

func (x *DescribeSolutionInput) GetVersion() uint32 {
	if x != nil {
		return x.Version
	}
	return 0
}

type DescribeSolutionOutput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Solution *Solution `protobuf:"bytes,1,opt,name=solution,proto3" json:"solution,omitempty"`
}

func (x *DescribeSolutionOutput) Reset() {
	*x = DescribeSolutionOutput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_atlas_solution_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DescribeSolutionOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescribeSolutionOutput) ProtoMessage() {}

func (x *DescribeSolutionOutput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_atlas_solution_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DescribeSolutionOutput.ProtoReflect.Descriptor instead.
func (*DescribeSolutionOutput) Descriptor() ([]byte, []int) {
	return file_eolymp_atlas_solution_service_proto_rawDescGZIP(), []int{3}
}

func (x *DescribeSolutionOutput) GetSolution() *Solution {
	if x != nil {
		return x.Solution
	}
	return nil
}

type CreateSolutionInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProblemId string    `protobuf:"bytes,1,opt,name=problem_id,json=problemId,proto3" json:"problem_id,omitempty"`
	Solution  *Solution `protobuf:"bytes,2,opt,name=solution,proto3" json:"solution,omitempty"`
}

func (x *CreateSolutionInput) Reset() {
	*x = CreateSolutionInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_atlas_solution_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateSolutionInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateSolutionInput) ProtoMessage() {}

func (x *CreateSolutionInput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_atlas_solution_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateSolutionInput.ProtoReflect.Descriptor instead.
func (*CreateSolutionInput) Descriptor() ([]byte, []int) {
	return file_eolymp_atlas_solution_service_proto_rawDescGZIP(), []int{4}
}

func (x *CreateSolutionInput) GetProblemId() string {
	if x != nil {
		return x.ProblemId
	}
	return ""
}

func (x *CreateSolutionInput) GetSolution() *Solution {
	if x != nil {
		return x.Solution
	}
	return nil
}

type CreateSolutionOutput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SolutionId string `protobuf:"bytes,1,opt,name=solution_id,json=solutionId,proto3" json:"solution_id,omitempty"`
}

func (x *CreateSolutionOutput) Reset() {
	*x = CreateSolutionOutput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_atlas_solution_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateSolutionOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateSolutionOutput) ProtoMessage() {}

func (x *CreateSolutionOutput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_atlas_solution_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateSolutionOutput.ProtoReflect.Descriptor instead.
func (*CreateSolutionOutput) Descriptor() ([]byte, []int) {
	return file_eolymp_atlas_solution_service_proto_rawDescGZIP(), []int{5}
}

func (x *CreateSolutionOutput) GetSolutionId() string {
	if x != nil {
		return x.SolutionId
	}
	return ""
}

type UpdateSolutionInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// specify list of fields to update, if empty all fields are updated
	Patch      []UpdateSolutionInput_Patch `protobuf:"varint,1,rep,packed,name=patch,proto3,enum=eolymp.atlas.UpdateSolutionInput_Patch" json:"patch,omitempty"`
	ProblemId  string                      `protobuf:"bytes,2,opt,name=problem_id,json=problemId,proto3" json:"problem_id,omitempty"`
	SolutionId string                      `protobuf:"bytes,3,opt,name=solution_id,json=solutionId,proto3" json:"solution_id,omitempty"`
	Solution   *Solution                   `protobuf:"bytes,4,opt,name=solution,proto3" json:"solution,omitempty"`
}

func (x *UpdateSolutionInput) Reset() {
	*x = UpdateSolutionInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_atlas_solution_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateSolutionInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateSolutionInput) ProtoMessage() {}

func (x *UpdateSolutionInput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_atlas_solution_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateSolutionInput.ProtoReflect.Descriptor instead.
func (*UpdateSolutionInput) Descriptor() ([]byte, []int) {
	return file_eolymp_atlas_solution_service_proto_rawDescGZIP(), []int{6}
}

func (x *UpdateSolutionInput) GetPatch() []UpdateSolutionInput_Patch {
	if x != nil {
		return x.Patch
	}
	return nil
}

func (x *UpdateSolutionInput) GetProblemId() string {
	if x != nil {
		return x.ProblemId
	}
	return ""
}

func (x *UpdateSolutionInput) GetSolutionId() string {
	if x != nil {
		return x.SolutionId
	}
	return ""
}

func (x *UpdateSolutionInput) GetSolution() *Solution {
	if x != nil {
		return x.Solution
	}
	return nil
}

type UpdateSolutionOutput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateSolutionOutput) Reset() {
	*x = UpdateSolutionOutput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_atlas_solution_service_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateSolutionOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateSolutionOutput) ProtoMessage() {}

func (x *UpdateSolutionOutput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_atlas_solution_service_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateSolutionOutput.ProtoReflect.Descriptor instead.
func (*UpdateSolutionOutput) Descriptor() ([]byte, []int) {
	return file_eolymp_atlas_solution_service_proto_rawDescGZIP(), []int{7}
}

type DeleteSolutionInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProblemId  string `protobuf:"bytes,1,opt,name=problem_id,json=problemId,proto3" json:"problem_id,omitempty"`
	SolutionId string `protobuf:"bytes,2,opt,name=solution_id,json=solutionId,proto3" json:"solution_id,omitempty"`
}

func (x *DeleteSolutionInput) Reset() {
	*x = DeleteSolutionInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_atlas_solution_service_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteSolutionInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteSolutionInput) ProtoMessage() {}

func (x *DeleteSolutionInput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_atlas_solution_service_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteSolutionInput.ProtoReflect.Descriptor instead.
func (*DeleteSolutionInput) Descriptor() ([]byte, []int) {
	return file_eolymp_atlas_solution_service_proto_rawDescGZIP(), []int{8}
}

func (x *DeleteSolutionInput) GetProblemId() string {
	if x != nil {
		return x.ProblemId
	}
	return ""
}

func (x *DeleteSolutionInput) GetSolutionId() string {
	if x != nil {
		return x.SolutionId
	}
	return ""
}

type DeleteSolutionOutput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteSolutionOutput) Reset() {
	*x = DeleteSolutionOutput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_atlas_solution_service_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteSolutionOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteSolutionOutput) ProtoMessage() {}

func (x *DeleteSolutionOutput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_atlas_solution_service_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteSolutionOutput.ProtoReflect.Descriptor instead.
func (*DeleteSolutionOutput) Descriptor() ([]byte, []int) {
	return file_eolymp_atlas_solution_service_proto_rawDescGZIP(), []int{9}
}

var File_eolymp_atlas_solution_service_proto protoreflect.FileDescriptor

var file_eolymp_atlas_solution_service_proto_rawDesc = []byte{
	0x0a, 0x23, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x61, 0x74, 0x6c, 0x61, 0x73, 0x2f, 0x73,
	0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x61, 0x74,
	0x6c, 0x61, 0x73, 0x1a, 0x1d, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x22, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x72, 0x61, 0x74, 0x65, 0x6c, 0x69, 0x6d, 0x69, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x61,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x63, 0x6f, 0x70, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x61,
	0x74, 0x6c, 0x61, 0x73, 0x2f, 0x73, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x4d, 0x0a, 0x12, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x6f, 0x6c, 0x75, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f,
	0x62, 0x6c, 0x65, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70,
	0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x18, 0x64, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x22, 0x59, 0x0a, 0x13, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x6f, 0x6c, 0x75, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74,
	0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12,
	0x2c, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16,
	0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x61, 0x74, 0x6c, 0x61, 0x73, 0x2e, 0x53, 0x6f,
	0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x22, 0x71, 0x0a,
	0x15, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x53, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f,
	0x6e, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x62, 0x6c, 0x65,
	0x6d, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x62,
	0x6c, 0x65, 0x6d, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x6f, 0x6c, 0x75,
	0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x18, 0x64, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x22, 0x4c, 0x0a, 0x16, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x53, 0x6f, 0x6c, 0x75,
	0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x12, 0x32, 0x0a, 0x08, 0x73, 0x6f,
	0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x65,
	0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x61, 0x74, 0x6c, 0x61, 0x73, 0x2e, 0x53, 0x6f, 0x6c, 0x75,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x08, 0x73, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x68,
	0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e,
	0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x62, 0x6c,
	0x65, 0x6d, 0x49, 0x64, 0x12, 0x32, 0x0a, 0x08, 0x73, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e,
	0x61, 0x74, 0x6c, 0x61, 0x73, 0x2e, 0x53, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x08,
	0x73, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x37, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x53, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74,
	0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x49,
	0x64, 0x22, 0x87, 0x02, 0x0a, 0x13, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x6f, 0x6c, 0x75,
	0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x3d, 0x0a, 0x05, 0x70, 0x61, 0x74,
	0x63, 0x68, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x27, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d,
	0x70, 0x2e, 0x61, 0x74, 0x6c, 0x61, 0x73, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x6f,
	0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x2e, 0x50, 0x61, 0x74, 0x63,
	0x68, 0x52, 0x05, 0x70, 0x61, 0x74, 0x63, 0x68, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x62,
	0x6c, 0x65, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72,
	0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x6f, 0x6c, 0x75, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x6f,
	0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x32, 0x0a, 0x08, 0x73, 0x6f, 0x6c, 0x75,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x65, 0x6f, 0x6c,
	0x79, 0x6d, 0x70, 0x2e, 0x61, 0x74, 0x6c, 0x61, 0x73, 0x2e, 0x53, 0x6f, 0x6c, 0x75, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x08, 0x73, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x3d, 0x0a, 0x05,
	0x50, 0x61, 0x74, 0x63, 0x68, 0x12, 0x07, 0x0a, 0x03, 0x41, 0x4c, 0x4c, 0x10, 0x00, 0x12, 0x08,
	0x0a, 0x04, 0x4e, 0x41, 0x4d, 0x45, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x54, 0x59, 0x50, 0x45,
	0x10, 0x02, 0x12, 0x0b, 0x0a, 0x07, 0x52, 0x55, 0x4e, 0x54, 0x49, 0x4d, 0x45, 0x10, 0x03, 0x12,
	0x0a, 0x0a, 0x06, 0x53, 0x4f, 0x55, 0x52, 0x43, 0x45, 0x10, 0x04, 0x22, 0x16, 0x0a, 0x14, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x75, 0x74,
	0x70, 0x75, 0x74, 0x22, 0x55, 0x0a, 0x13, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x6f, 0x6c,
	0x75, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72,
	0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x70, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x6f, 0x6c,
	0x75, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x73, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x22, 0x16, 0x0a, 0x14, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x53, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x75, 0x74, 0x70,
	0x75, 0x74, 0x32, 0xb4, 0x06, 0x0a, 0x0f, 0x53, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x95, 0x01, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x53, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x21, 0x2e, 0x65, 0x6f, 0x6c, 0x79,
	0x6d, 0x70, 0x2e, 0x61, 0x74, 0x6c, 0x61, 0x73, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53,
	0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x22, 0x2e, 0x65,
	0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x61, 0x74, 0x6c, 0x61, 0x73, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x53, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74,
	0x22, 0x3c, 0xea, 0xe2, 0x0a, 0x0b, 0xf5, 0xe2, 0x0a, 0x00, 0x00, 0x80, 0x3f, 0xf8, 0xe2, 0x0a,
	0x05, 0x82, 0xe3, 0x0a, 0x17, 0x8a, 0xe3, 0x0a, 0x13, 0x61, 0x74, 0x6c, 0x61, 0x73, 0x3a, 0x70,
	0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x3a, 0x77, 0x72, 0x69, 0x74, 0x65, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x0c, 0x1a, 0x0a, 0x2f, 0x73, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0xa3,
	0x01, 0x0a, 0x0e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x21, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x61, 0x74, 0x6c, 0x61, 0x73,
	0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x49,
	0x6e, 0x70, 0x75, 0x74, 0x1a, 0x22, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x61, 0x74,
	0x6c, 0x61, 0x73, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x6f, 0x6c, 0x75, 0x74, 0x69,
	0x6f, 0x6e, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x22, 0x4a, 0xea, 0xe2, 0x0a, 0x0b, 0xf5, 0xe2,
	0x0a, 0x00, 0x00, 0x80, 0x3f, 0xf8, 0xe2, 0x0a, 0x05, 0x82, 0xe3, 0x0a, 0x17, 0x8a, 0xe3, 0x0a,
	0x13, 0x61, 0x74, 0x6c, 0x61, 0x73, 0x3a, 0x70, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x3a, 0x77,
	0x72, 0x69, 0x74, 0x65, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1a, 0x1a, 0x18, 0x2f, 0x73, 0x6f, 0x6c,
	0x75, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x7b, 0x73, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x69, 0x64, 0x7d, 0x12, 0xa3, 0x01, 0x0a, 0x0e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53,
	0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x21, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70,
	0x2e, 0x61, 0x74, 0x6c, 0x61, 0x73, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x6f, 0x6c,
	0x75, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x22, 0x2e, 0x65, 0x6f, 0x6c,
	0x79, 0x6d, 0x70, 0x2e, 0x61, 0x74, 0x6c, 0x61, 0x73, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x53, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x22, 0x4a,
	0xea, 0xe2, 0x0a, 0x0b, 0xf5, 0xe2, 0x0a, 0x00, 0x00, 0x80, 0x3f, 0xf8, 0xe2, 0x0a, 0x05, 0x82,
	0xe3, 0x0a, 0x17, 0x8a, 0xe3, 0x0a, 0x13, 0x61, 0x74, 0x6c, 0x61, 0x73, 0x3a, 0x70, 0x72, 0x6f,
	0x62, 0x6c, 0x65, 0x6d, 0x3a, 0x77, 0x72, 0x69, 0x74, 0x65, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1a,
	0x2a, 0x18, 0x2f, 0x73, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x7b, 0x73, 0x6f,
	0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x7d, 0x12, 0xa8, 0x01, 0x0a, 0x10, 0x44,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x53, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x23, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x61, 0x74, 0x6c, 0x61, 0x73, 0x2e, 0x44,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x53, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x49,
	0x6e, 0x70, 0x75, 0x74, 0x1a, 0x24, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x61, 0x74,
	0x6c, 0x61, 0x73, 0x2e, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x53, 0x6f, 0x6c, 0x75,
	0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x22, 0x49, 0xea, 0xe2, 0x0a, 0x0b,
	0xf5, 0xe2, 0x0a, 0x00, 0x00, 0xa0, 0x41, 0xf8, 0xe2, 0x0a, 0x64, 0x82, 0xe3, 0x0a, 0x16, 0x8a,
	0xe3, 0x0a, 0x12, 0x61, 0x74, 0x6c, 0x61, 0x73, 0x3a, 0x70, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d,
	0x3a, 0x72, 0x65, 0x61, 0x64, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1a, 0x12, 0x18, 0x2f, 0x73, 0x6f,
	0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x7b, 0x73, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x69, 0x64, 0x7d, 0x12, 0x91, 0x01, 0x0a, 0x0d, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x6f,
	0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x20, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70,
	0x2e, 0x61, 0x74, 0x6c, 0x61, 0x73, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x6f, 0x6c, 0x75, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x21, 0x2e, 0x65, 0x6f, 0x6c, 0x79,
	0x6d, 0x70, 0x2e, 0x61, 0x74, 0x6c, 0x61, 0x73, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x6f, 0x6c,
	0x75, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x22, 0x3b, 0xea, 0xe2,
	0x0a, 0x0b, 0xf5, 0xe2, 0x0a, 0x00, 0x00, 0xa0, 0x41, 0xf8, 0xe2, 0x0a, 0x64, 0x82, 0xe3, 0x0a,
	0x16, 0x8a, 0xe3, 0x0a, 0x12, 0x61, 0x74, 0x6c, 0x61, 0x73, 0x3a, 0x70, 0x72, 0x6f, 0x62, 0x6c,
	0x65, 0x6d, 0x3a, 0x72, 0x65, 0x61, 0x64, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0c, 0x12, 0x0a, 0x2f,
	0x73, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x42, 0x2d, 0x5a, 0x2b, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x67,
	0x6f, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x61, 0x74, 0x6c,
	0x61, 0x73, 0x3b, 0x61, 0x74, 0x6c, 0x61, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_eolymp_atlas_solution_service_proto_rawDescOnce sync.Once
	file_eolymp_atlas_solution_service_proto_rawDescData = file_eolymp_atlas_solution_service_proto_rawDesc
)

func file_eolymp_atlas_solution_service_proto_rawDescGZIP() []byte {
	file_eolymp_atlas_solution_service_proto_rawDescOnce.Do(func() {
		file_eolymp_atlas_solution_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_eolymp_atlas_solution_service_proto_rawDescData)
	})
	return file_eolymp_atlas_solution_service_proto_rawDescData
}

var file_eolymp_atlas_solution_service_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_eolymp_atlas_solution_service_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_eolymp_atlas_solution_service_proto_goTypes = []interface{}{
	(UpdateSolutionInput_Patch)(0), // 0: eolymp.atlas.UpdateSolutionInput.Patch
	(*ListSolutionsInput)(nil),     // 1: eolymp.atlas.ListSolutionsInput
	(*ListSolutionsOutput)(nil),    // 2: eolymp.atlas.ListSolutionsOutput
	(*DescribeSolutionInput)(nil),  // 3: eolymp.atlas.DescribeSolutionInput
	(*DescribeSolutionOutput)(nil), // 4: eolymp.atlas.DescribeSolutionOutput
	(*CreateSolutionInput)(nil),    // 5: eolymp.atlas.CreateSolutionInput
	(*CreateSolutionOutput)(nil),   // 6: eolymp.atlas.CreateSolutionOutput
	(*UpdateSolutionInput)(nil),    // 7: eolymp.atlas.UpdateSolutionInput
	(*UpdateSolutionOutput)(nil),   // 8: eolymp.atlas.UpdateSolutionOutput
	(*DeleteSolutionInput)(nil),    // 9: eolymp.atlas.DeleteSolutionInput
	(*DeleteSolutionOutput)(nil),   // 10: eolymp.atlas.DeleteSolutionOutput
	(*Solution)(nil),               // 11: eolymp.atlas.Solution
}
var file_eolymp_atlas_solution_service_proto_depIdxs = []int32{
	11, // 0: eolymp.atlas.ListSolutionsOutput.items:type_name -> eolymp.atlas.Solution
	11, // 1: eolymp.atlas.DescribeSolutionOutput.solution:type_name -> eolymp.atlas.Solution
	11, // 2: eolymp.atlas.CreateSolutionInput.solution:type_name -> eolymp.atlas.Solution
	0,  // 3: eolymp.atlas.UpdateSolutionInput.patch:type_name -> eolymp.atlas.UpdateSolutionInput.Patch
	11, // 4: eolymp.atlas.UpdateSolutionInput.solution:type_name -> eolymp.atlas.Solution
	5,  // 5: eolymp.atlas.SolutionService.CreateSolution:input_type -> eolymp.atlas.CreateSolutionInput
	7,  // 6: eolymp.atlas.SolutionService.UpdateSolution:input_type -> eolymp.atlas.UpdateSolutionInput
	9,  // 7: eolymp.atlas.SolutionService.DeleteSolution:input_type -> eolymp.atlas.DeleteSolutionInput
	3,  // 8: eolymp.atlas.SolutionService.DescribeSolution:input_type -> eolymp.atlas.DescribeSolutionInput
	1,  // 9: eolymp.atlas.SolutionService.ListSolutions:input_type -> eolymp.atlas.ListSolutionsInput
	6,  // 10: eolymp.atlas.SolutionService.CreateSolution:output_type -> eolymp.atlas.CreateSolutionOutput
	8,  // 11: eolymp.atlas.SolutionService.UpdateSolution:output_type -> eolymp.atlas.UpdateSolutionOutput
	10, // 12: eolymp.atlas.SolutionService.DeleteSolution:output_type -> eolymp.atlas.DeleteSolutionOutput
	4,  // 13: eolymp.atlas.SolutionService.DescribeSolution:output_type -> eolymp.atlas.DescribeSolutionOutput
	2,  // 14: eolymp.atlas.SolutionService.ListSolutions:output_type -> eolymp.atlas.ListSolutionsOutput
	10, // [10:15] is the sub-list for method output_type
	5,  // [5:10] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_eolymp_atlas_solution_service_proto_init() }
func file_eolymp_atlas_solution_service_proto_init() {
	if File_eolymp_atlas_solution_service_proto != nil {
		return
	}
	file_eolymp_atlas_solution_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_eolymp_atlas_solution_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListSolutionsInput); i {
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
		file_eolymp_atlas_solution_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListSolutionsOutput); i {
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
		file_eolymp_atlas_solution_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DescribeSolutionInput); i {
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
		file_eolymp_atlas_solution_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DescribeSolutionOutput); i {
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
		file_eolymp_atlas_solution_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateSolutionInput); i {
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
		file_eolymp_atlas_solution_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateSolutionOutput); i {
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
		file_eolymp_atlas_solution_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateSolutionInput); i {
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
		file_eolymp_atlas_solution_service_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateSolutionOutput); i {
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
		file_eolymp_atlas_solution_service_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteSolutionInput); i {
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
		file_eolymp_atlas_solution_service_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteSolutionOutput); i {
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
			RawDescriptor: file_eolymp_atlas_solution_service_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_eolymp_atlas_solution_service_proto_goTypes,
		DependencyIndexes: file_eolymp_atlas_solution_service_proto_depIdxs,
		EnumInfos:         file_eolymp_atlas_solution_service_proto_enumTypes,
		MessageInfos:      file_eolymp_atlas_solution_service_proto_msgTypes,
	}.Build()
	File_eolymp_atlas_solution_service_proto = out.File
	file_eolymp_atlas_solution_service_proto_rawDesc = nil
	file_eolymp_atlas_solution_service_proto_goTypes = nil
	file_eolymp_atlas_solution_service_proto_depIdxs = nil
}