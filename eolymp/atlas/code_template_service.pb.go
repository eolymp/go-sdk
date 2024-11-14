// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v3.21.12
// source: eolymp/atlas/code_template_service.proto

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

type UpdateCodeTemplateInput_Patch int32

const (
	UpdateCodeTemplateInput_ALL     UpdateCodeTemplateInput_Patch = 0
	UpdateCodeTemplateInput_RUNTIME UpdateCodeTemplateInput_Patch = 1
	UpdateCodeTemplateInput_SOURCE  UpdateCodeTemplateInput_Patch = 2
	UpdateCodeTemplateInput_HEADER  UpdateCodeTemplateInput_Patch = 3
	UpdateCodeTemplateInput_FOOTER  UpdateCodeTemplateInput_Patch = 4
	UpdateCodeTemplateInput_SECRET  UpdateCodeTemplateInput_Patch = 5
	UpdateCodeTemplateInput_FILES   UpdateCodeTemplateInput_Patch = 6
)

// Enum value maps for UpdateCodeTemplateInput_Patch.
var (
	UpdateCodeTemplateInput_Patch_name = map[int32]string{
		0: "ALL",
		1: "RUNTIME",
		2: "SOURCE",
		3: "HEADER",
		4: "FOOTER",
		5: "SECRET",
		6: "FILES",
	}
	UpdateCodeTemplateInput_Patch_value = map[string]int32{
		"ALL":     0,
		"RUNTIME": 1,
		"SOURCE":  2,
		"HEADER":  3,
		"FOOTER":  4,
		"SECRET":  5,
		"FILES":   6,
	}
)

func (x UpdateCodeTemplateInput_Patch) Enum() *UpdateCodeTemplateInput_Patch {
	p := new(UpdateCodeTemplateInput_Patch)
	*p = x
	return p
}

func (x UpdateCodeTemplateInput_Patch) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (UpdateCodeTemplateInput_Patch) Descriptor() protoreflect.EnumDescriptor {
	return file_eolymp_atlas_code_template_service_proto_enumTypes[0].Descriptor()
}

func (UpdateCodeTemplateInput_Patch) Type() protoreflect.EnumType {
	return &file_eolymp_atlas_code_template_service_proto_enumTypes[0]
}

func (x UpdateCodeTemplateInput_Patch) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use UpdateCodeTemplateInput_Patch.Descriptor instead.
func (UpdateCodeTemplateInput_Patch) EnumDescriptor() ([]byte, []int) {
	return file_eolymp_atlas_code_template_service_proto_rawDescGZIP(), []int{2, 0}
}

type CreateCodeTemplateInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Template *Template `protobuf:"bytes,2,opt,name=template,proto3" json:"template,omitempty"`
}

func (x *CreateCodeTemplateInput) Reset() {
	*x = CreateCodeTemplateInput{}
	mi := &file_eolymp_atlas_code_template_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateCodeTemplateInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCodeTemplateInput) ProtoMessage() {}

func (x *CreateCodeTemplateInput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_atlas_code_template_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCodeTemplateInput.ProtoReflect.Descriptor instead.
func (*CreateCodeTemplateInput) Descriptor() ([]byte, []int) {
	return file_eolymp_atlas_code_template_service_proto_rawDescGZIP(), []int{0}
}

func (x *CreateCodeTemplateInput) GetTemplate() *Template {
	if x != nil {
		return x.Template
	}
	return nil
}

type CreateCodeTemplateOutput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TemplateId string `protobuf:"bytes,1,opt,name=template_id,json=templateId,proto3" json:"template_id,omitempty"`
}

func (x *CreateCodeTemplateOutput) Reset() {
	*x = CreateCodeTemplateOutput{}
	mi := &file_eolymp_atlas_code_template_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateCodeTemplateOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCodeTemplateOutput) ProtoMessage() {}

func (x *CreateCodeTemplateOutput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_atlas_code_template_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCodeTemplateOutput.ProtoReflect.Descriptor instead.
func (*CreateCodeTemplateOutput) Descriptor() ([]byte, []int) {
	return file_eolymp_atlas_code_template_service_proto_rawDescGZIP(), []int{1}
}

func (x *CreateCodeTemplateOutput) GetTemplateId() string {
	if x != nil {
		return x.TemplateId
	}
	return ""
}

type UpdateCodeTemplateInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TemplateId string    `protobuf:"bytes,2,opt,name=template_id,json=templateId,proto3" json:"template_id,omitempty"`
	Template   *Template `protobuf:"bytes,3,opt,name=template,proto3" json:"template,omitempty"`
}

func (x *UpdateCodeTemplateInput) Reset() {
	*x = UpdateCodeTemplateInput{}
	mi := &file_eolymp_atlas_code_template_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateCodeTemplateInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateCodeTemplateInput) ProtoMessage() {}

func (x *UpdateCodeTemplateInput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_atlas_code_template_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateCodeTemplateInput.ProtoReflect.Descriptor instead.
func (*UpdateCodeTemplateInput) Descriptor() ([]byte, []int) {
	return file_eolymp_atlas_code_template_service_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateCodeTemplateInput) GetTemplateId() string {
	if x != nil {
		return x.TemplateId
	}
	return ""
}

func (x *UpdateCodeTemplateInput) GetTemplate() *Template {
	if x != nil {
		return x.Template
	}
	return nil
}

type UpdateCodeTemplateOutput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateCodeTemplateOutput) Reset() {
	*x = UpdateCodeTemplateOutput{}
	mi := &file_eolymp_atlas_code_template_service_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateCodeTemplateOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateCodeTemplateOutput) ProtoMessage() {}

func (x *UpdateCodeTemplateOutput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_atlas_code_template_service_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateCodeTemplateOutput.ProtoReflect.Descriptor instead.
func (*UpdateCodeTemplateOutput) Descriptor() ([]byte, []int) {
	return file_eolymp_atlas_code_template_service_proto_rawDescGZIP(), []int{3}
}

type DeleteCodeTemplateInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TemplateId string `protobuf:"bytes,2,opt,name=template_id,json=templateId,proto3" json:"template_id,omitempty"`
}

func (x *DeleteCodeTemplateInput) Reset() {
	*x = DeleteCodeTemplateInput{}
	mi := &file_eolymp_atlas_code_template_service_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteCodeTemplateInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteCodeTemplateInput) ProtoMessage() {}

func (x *DeleteCodeTemplateInput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_atlas_code_template_service_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteCodeTemplateInput.ProtoReflect.Descriptor instead.
func (*DeleteCodeTemplateInput) Descriptor() ([]byte, []int) {
	return file_eolymp_atlas_code_template_service_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteCodeTemplateInput) GetTemplateId() string {
	if x != nil {
		return x.TemplateId
	}
	return ""
}

type DeleteCodeTemplateOutput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteCodeTemplateOutput) Reset() {
	*x = DeleteCodeTemplateOutput{}
	mi := &file_eolymp_atlas_code_template_service_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteCodeTemplateOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteCodeTemplateOutput) ProtoMessage() {}

func (x *DeleteCodeTemplateOutput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_atlas_code_template_service_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteCodeTemplateOutput.ProtoReflect.Descriptor instead.
func (*DeleteCodeTemplateOutput) Descriptor() ([]byte, []int) {
	return file_eolymp_atlas_code_template_service_proto_rawDescGZIP(), []int{5}
}

type ListCodeTemplatesInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Offset  int32  `protobuf:"varint,10,opt,name=offset,proto3" json:"offset,omitempty"`
	Size    int32  `protobuf:"varint,11,opt,name=size,proto3" json:"size,omitempty"`
	Version uint32 `protobuf:"varint,100,opt,name=version,proto3" json:"version,omitempty"` // request data for specific problem version
}

func (x *ListCodeTemplatesInput) Reset() {
	*x = ListCodeTemplatesInput{}
	mi := &file_eolymp_atlas_code_template_service_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListCodeTemplatesInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListCodeTemplatesInput) ProtoMessage() {}

func (x *ListCodeTemplatesInput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_atlas_code_template_service_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListCodeTemplatesInput.ProtoReflect.Descriptor instead.
func (*ListCodeTemplatesInput) Descriptor() ([]byte, []int) {
	return file_eolymp_atlas_code_template_service_proto_rawDescGZIP(), []int{6}
}

func (x *ListCodeTemplatesInput) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *ListCodeTemplatesInput) GetSize() int32 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *ListCodeTemplatesInput) GetVersion() uint32 {
	if x != nil {
		return x.Version
	}
	return 0
}

type ListCodeTemplatesOutput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Total int32       `protobuf:"varint,1,opt,name=total,proto3" json:"total,omitempty"`
	Items []*Template `protobuf:"bytes,2,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *ListCodeTemplatesOutput) Reset() {
	*x = ListCodeTemplatesOutput{}
	mi := &file_eolymp_atlas_code_template_service_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListCodeTemplatesOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListCodeTemplatesOutput) ProtoMessage() {}

func (x *ListCodeTemplatesOutput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_atlas_code_template_service_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListCodeTemplatesOutput.ProtoReflect.Descriptor instead.
func (*ListCodeTemplatesOutput) Descriptor() ([]byte, []int) {
	return file_eolymp_atlas_code_template_service_proto_rawDescGZIP(), []int{7}
}

func (x *ListCodeTemplatesOutput) GetTotal() int32 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *ListCodeTemplatesOutput) GetItems() []*Template {
	if x != nil {
		return x.Items
	}
	return nil
}

type DescribeCodeTemplateInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TemplateId string `protobuf:"bytes,2,opt,name=template_id,json=templateId,proto3" json:"template_id,omitempty"`
	Version    uint32 `protobuf:"varint,100,opt,name=version,proto3" json:"version,omitempty"` // request data for specific problem version
}

func (x *DescribeCodeTemplateInput) Reset() {
	*x = DescribeCodeTemplateInput{}
	mi := &file_eolymp_atlas_code_template_service_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DescribeCodeTemplateInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescribeCodeTemplateInput) ProtoMessage() {}

func (x *DescribeCodeTemplateInput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_atlas_code_template_service_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DescribeCodeTemplateInput.ProtoReflect.Descriptor instead.
func (*DescribeCodeTemplateInput) Descriptor() ([]byte, []int) {
	return file_eolymp_atlas_code_template_service_proto_rawDescGZIP(), []int{8}
}

func (x *DescribeCodeTemplateInput) GetTemplateId() string {
	if x != nil {
		return x.TemplateId
	}
	return ""
}

func (x *DescribeCodeTemplateInput) GetVersion() uint32 {
	if x != nil {
		return x.Version
	}
	return 0
}

type DescribeCodeTemplateOutput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Template *Template `protobuf:"bytes,1,opt,name=template,proto3" json:"template,omitempty"`
}

func (x *DescribeCodeTemplateOutput) Reset() {
	*x = DescribeCodeTemplateOutput{}
	mi := &file_eolymp_atlas_code_template_service_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DescribeCodeTemplateOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescribeCodeTemplateOutput) ProtoMessage() {}

func (x *DescribeCodeTemplateOutput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_atlas_code_template_service_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DescribeCodeTemplateOutput.ProtoReflect.Descriptor instead.
func (*DescribeCodeTemplateOutput) Descriptor() ([]byte, []int) {
	return file_eolymp_atlas_code_template_service_proto_rawDescGZIP(), []int{9}
}

func (x *DescribeCodeTemplateOutput) GetTemplate() *Template {
	if x != nil {
		return x.Template
	}
	return nil
}

type LookupCodeTemplateInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Runtime string `protobuf:"bytes,2,opt,name=runtime,proto3" json:"runtime,omitempty"`
}

func (x *LookupCodeTemplateInput) Reset() {
	*x = LookupCodeTemplateInput{}
	mi := &file_eolymp_atlas_code_template_service_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LookupCodeTemplateInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LookupCodeTemplateInput) ProtoMessage() {}

func (x *LookupCodeTemplateInput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_atlas_code_template_service_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LookupCodeTemplateInput.ProtoReflect.Descriptor instead.
func (*LookupCodeTemplateInput) Descriptor() ([]byte, []int) {
	return file_eolymp_atlas_code_template_service_proto_rawDescGZIP(), []int{10}
}

func (x *LookupCodeTemplateInput) GetRuntime() string {
	if x != nil {
		return x.Runtime
	}
	return ""
}

type LookupCodeTemplateOutput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Template *Template `protobuf:"bytes,1,opt,name=template,proto3" json:"template,omitempty"`
}

func (x *LookupCodeTemplateOutput) Reset() {
	*x = LookupCodeTemplateOutput{}
	mi := &file_eolymp_atlas_code_template_service_proto_msgTypes[11]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LookupCodeTemplateOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LookupCodeTemplateOutput) ProtoMessage() {}

func (x *LookupCodeTemplateOutput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_atlas_code_template_service_proto_msgTypes[11]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LookupCodeTemplateOutput.ProtoReflect.Descriptor instead.
func (*LookupCodeTemplateOutput) Descriptor() ([]byte, []int) {
	return file_eolymp_atlas_code_template_service_proto_rawDescGZIP(), []int{11}
}

func (x *LookupCodeTemplateOutput) GetTemplate() *Template {
	if x != nil {
		return x.Template
	}
	return nil
}

var File_eolymp_atlas_code_template_service_proto protoreflect.FileDescriptor

var file_eolymp_atlas_code_template_service_proto_rawDesc = []byte{
	0x0a, 0x28, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x61, 0x74, 0x6c, 0x61, 0x73, 0x2f, 0x63,
	0x6f, 0x64, 0x65, 0x5f, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x65, 0x6f, 0x6c, 0x79,
	0x6d, 0x70, 0x2e, 0x61, 0x74, 0x6c, 0x61, 0x73, 0x1a, 0x1d, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70,
	0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x68, 0x74, 0x74,
	0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x22, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x72, 0x61, 0x74, 0x65,
	0x6c, 0x69, 0x6d, 0x69, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x65, 0x6f, 0x6c,
	0x79, 0x6d, 0x70, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f,
	0x73, 0x63, 0x6f, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x65, 0x6f, 0x6c,
	0x79, 0x6d, 0x70, 0x2f, 0x61, 0x74, 0x6c, 0x61, 0x73, 0x2f, 0x63, 0x6f, 0x64, 0x65, 0x5f, 0x74,
	0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4d, 0x0a,
	0x17, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x54, 0x65, 0x6d, 0x70, 0x6c,
	0x61, 0x74, 0x65, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x32, 0x0a, 0x08, 0x74, 0x65, 0x6d, 0x70,
	0x6c, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x65, 0x6f, 0x6c,
	0x79, 0x6d, 0x70, 0x2e, 0x61, 0x74, 0x6c, 0x61, 0x73, 0x2e, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61,
	0x74, 0x65, 0x52, 0x08, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x22, 0x3b, 0x0a, 0x18,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61,
	0x74, 0x65, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x65, 0x6d, 0x70,
	0x6c, 0x61, 0x74, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x74,
	0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x49, 0x64, 0x22, 0xc8, 0x01, 0x0a, 0x17, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65,
	0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74,
	0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x74, 0x65, 0x6d, 0x70,
	0x6c, 0x61, 0x74, 0x65, 0x49, 0x64, 0x12, 0x32, 0x0a, 0x08, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61,
	0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d,
	0x70, 0x2e, 0x61, 0x74, 0x6c, 0x61, 0x73, 0x2e, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65,
	0x52, 0x08, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x22, 0x58, 0x0a, 0x05, 0x50, 0x61,
	0x74, 0x63, 0x68, 0x12, 0x07, 0x0a, 0x03, 0x41, 0x4c, 0x4c, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07,
	0x52, 0x55, 0x4e, 0x54, 0x49, 0x4d, 0x45, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x53, 0x4f, 0x55,
	0x52, 0x43, 0x45, 0x10, 0x02, 0x12, 0x0a, 0x0a, 0x06, 0x48, 0x45, 0x41, 0x44, 0x45, 0x52, 0x10,
	0x03, 0x12, 0x0a, 0x0a, 0x06, 0x46, 0x4f, 0x4f, 0x54, 0x45, 0x52, 0x10, 0x04, 0x12, 0x0a, 0x0a,
	0x06, 0x53, 0x45, 0x43, 0x52, 0x45, 0x54, 0x10, 0x05, 0x12, 0x09, 0x0a, 0x05, 0x46, 0x49, 0x4c,
	0x45, 0x53, 0x10, 0x06, 0x22, 0x1a, 0x0a, 0x18, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6f,
	0x64, 0x65, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74,
	0x22, 0x3a, 0x0a, 0x17, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x54, 0x65,
	0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x74,
	0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x49, 0x64, 0x22, 0x1a, 0x0a, 0x18,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61,
	0x74, 0x65, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x22, 0x5e, 0x0a, 0x16, 0x4c, 0x69, 0x73, 0x74,
	0x43, 0x6f, 0x64, 0x65, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x49, 0x6e, 0x70,
	0x75, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69,
	0x7a, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x64, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x5d, 0x0a, 0x17, 0x4c, 0x69, 0x73, 0x74,
	0x43, 0x6f, 0x64, 0x65, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x4f, 0x75, 0x74,
	0x70, 0x75, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x2c, 0x0a, 0x05, 0x69, 0x74, 0x65,
	0x6d, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d,
	0x70, 0x2e, 0x61, 0x74, 0x6c, 0x61, 0x73, 0x2e, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65,
	0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x22, 0x56, 0x0a, 0x19, 0x44, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x62, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x49,
	0x6e, 0x70, 0x75, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x74, 0x65, 0x6d, 0x70, 0x6c,
	0x61, 0x74, 0x65, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x18, 0x64, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x22,
	0x50, 0x0a, 0x1a, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x54,
	0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x12, 0x32, 0x0a,
	0x08, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x16, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x61, 0x74, 0x6c, 0x61, 0x73, 0x2e, 0x54,
	0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x08, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74,
	0x65, 0x22, 0x33, 0x0a, 0x17, 0x4c, 0x6f, 0x6f, 0x6b, 0x75, 0x70, 0x43, 0x6f, 0x64, 0x65, 0x54,
	0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x18, 0x0a, 0x07,
	0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x72,
	0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x22, 0x4e, 0x0a, 0x18, 0x4c, 0x6f, 0x6f, 0x6b, 0x75, 0x70,
	0x43, 0x6f, 0x64, 0x65, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x4f, 0x75, 0x74, 0x70,
	0x75, 0x74, 0x12, 0x32, 0x0a, 0x08, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x61, 0x74,
	0x6c, 0x61, 0x73, 0x2e, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x08, 0x74, 0x65,
	0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x32, 0x96, 0x08, 0x0a, 0x13, 0x43, 0x6f, 0x64, 0x65, 0x54,
	0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0xa1,
	0x01, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x54, 0x65, 0x6d,
	0x70, 0x6c, 0x61, 0x74, 0x65, 0x12, 0x25, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x61,
	0x74, 0x6c, 0x61, 0x73, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x54,
	0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x26, 0x2e, 0x65,
	0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x61, 0x74, 0x6c, 0x61, 0x73, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x4f, 0x75,
	0x74, 0x70, 0x75, 0x74, 0x22, 0x3c, 0xea, 0xe2, 0x0a, 0x0b, 0xf5, 0xe2, 0x0a, 0x0a, 0xd7, 0x23,
	0x3e, 0xf8, 0xe2, 0x0a, 0x05, 0x82, 0xe3, 0x0a, 0x17, 0x8a, 0xe3, 0x0a, 0x13, 0x61, 0x74, 0x6c,
	0x61, 0x73, 0x3a, 0x70, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x3a, 0x77, 0x72, 0x69, 0x74, 0x65,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0c, 0x22, 0x0a, 0x2f, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74,
	0x65, 0x73, 0x12, 0xaf, 0x01, 0x0a, 0x12, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x64,
	0x65, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x12, 0x25, 0x2e, 0x65, 0x6f, 0x6c, 0x79,
	0x6d, 0x70, 0x2e, 0x61, 0x74, 0x6c, 0x61, 0x73, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43,
	0x6f, 0x64, 0x65, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x70, 0x75, 0x74,
	0x1a, 0x26, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x61, 0x74, 0x6c, 0x61, 0x73, 0x2e,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61,
	0x74, 0x65, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x22, 0x4a, 0xea, 0xe2, 0x0a, 0x0b, 0xf5, 0xe2,
	0x0a, 0x0a, 0xd7, 0x23, 0x3e, 0xf8, 0xe2, 0x0a, 0x05, 0x82, 0xe3, 0x0a, 0x17, 0x8a, 0xe3, 0x0a,
	0x13, 0x61, 0x74, 0x6c, 0x61, 0x73, 0x3a, 0x70, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x3a, 0x77,
	0x72, 0x69, 0x74, 0x65, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1a, 0x22, 0x18, 0x2f, 0x74, 0x65, 0x6d,
	0x70, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x2f, 0x7b, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65,
	0x5f, 0x69, 0x64, 0x7d, 0x12, 0xaf, 0x01, 0x0a, 0x12, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43,
	0x6f, 0x64, 0x65, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x12, 0x25, 0x2e, 0x65, 0x6f,
	0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x61, 0x74, 0x6c, 0x61, 0x73, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x43, 0x6f, 0x64, 0x65, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x70,
	0x75, 0x74, 0x1a, 0x26, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x61, 0x74, 0x6c, 0x61,
	0x73, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x54, 0x65, 0x6d, 0x70,
	0x6c, 0x61, 0x74, 0x65, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x22, 0x4a, 0xea, 0xe2, 0x0a, 0x0b,
	0xf5, 0xe2, 0x0a, 0x0a, 0xd7, 0x23, 0x3e, 0xf8, 0xe2, 0x0a, 0x05, 0x82, 0xe3, 0x0a, 0x17, 0x8a,
	0xe3, 0x0a, 0x13, 0x61, 0x74, 0x6c, 0x61, 0x73, 0x3a, 0x70, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d,
	0x3a, 0x77, 0x72, 0x69, 0x74, 0x65, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1a, 0x2a, 0x18, 0x2f, 0x74,
	0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x2f, 0x7b, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61,
	0x74, 0x65, 0x5f, 0x69, 0x64, 0x7d, 0x12, 0x9d, 0x01, 0x0a, 0x11, 0x4c, 0x69, 0x73, 0x74, 0x43,
	0x6f, 0x64, 0x65, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x12, 0x24, 0x2e, 0x65,
	0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x61, 0x74, 0x6c, 0x61, 0x73, 0x2e, 0x4c, 0x69, 0x73, 0x74,
	0x43, 0x6f, 0x64, 0x65, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x49, 0x6e, 0x70,
	0x75, 0x74, 0x1a, 0x25, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x61, 0x74, 0x6c, 0x61,
	0x73, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61,
	0x74, 0x65, 0x73, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x22, 0x3b, 0xea, 0xe2, 0x0a, 0x0b, 0xf5,
	0xe2, 0x0a, 0x00, 0x00, 0xa0, 0x41, 0xf8, 0xe2, 0x0a, 0x64, 0x82, 0xe3, 0x0a, 0x16, 0x8a, 0xe3,
	0x0a, 0x12, 0x61, 0x74, 0x6c, 0x61, 0x73, 0x3a, 0x70, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x3a,
	0x72, 0x65, 0x61, 0x64, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0c, 0x12, 0x0a, 0x2f, 0x74, 0x65, 0x6d,
	0x70, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x12, 0xb4, 0x01, 0x0a, 0x14, 0x44, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x62, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x12,
	0x27, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x61, 0x74, 0x6c, 0x61, 0x73, 0x2e, 0x44,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x54, 0x65, 0x6d, 0x70, 0x6c,
	0x61, 0x74, 0x65, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x28, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d,
	0x70, 0x2e, 0x61, 0x74, 0x6c, 0x61, 0x73, 0x2e, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65,
	0x43, 0x6f, 0x64, 0x65, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x4f, 0x75, 0x74, 0x70,
	0x75, 0x74, 0x22, 0x49, 0xea, 0xe2, 0x0a, 0x0b, 0xf5, 0xe2, 0x0a, 0x00, 0x00, 0xa0, 0x41, 0xf8,
	0xe2, 0x0a, 0x64, 0x82, 0xe3, 0x0a, 0x16, 0x8a, 0xe3, 0x0a, 0x12, 0x61, 0x74, 0x6c, 0x61, 0x73,
	0x3a, 0x70, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x3a, 0x72, 0x65, 0x61, 0x64, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x1a, 0x12, 0x18, 0x2f, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x2f,
	0x7b, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x5f, 0x69, 0x64, 0x7d, 0x12, 0x9f, 0x01,
	0x0a, 0x12, 0x4c, 0x6f, 0x6f, 0x6b, 0x75, 0x70, 0x43, 0x6f, 0x64, 0x65, 0x54, 0x65, 0x6d, 0x70,
	0x6c, 0x61, 0x74, 0x65, 0x12, 0x25, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x61, 0x74,
	0x6c, 0x61, 0x73, 0x2e, 0x4c, 0x6f, 0x6f, 0x6b, 0x75, 0x70, 0x43, 0x6f, 0x64, 0x65, 0x54, 0x65,
	0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x26, 0x2e, 0x65, 0x6f,
	0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x61, 0x74, 0x6c, 0x61, 0x73, 0x2e, 0x4c, 0x6f, 0x6f, 0x6b, 0x75,
	0x70, 0x43, 0x6f, 0x64, 0x65, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x4f, 0x75, 0x74,
	0x70, 0x75, 0x74, 0x22, 0x3a, 0xea, 0xe2, 0x0a, 0x0b, 0xf5, 0xe2, 0x0a, 0x00, 0x00, 0xa0, 0x40,
	0xf8, 0xe2, 0x0a, 0x14, 0x82, 0xe3, 0x0a, 0x16, 0x8a, 0xe3, 0x0a, 0x12, 0x61, 0x74, 0x6c, 0x61,
	0x73, 0x3a, 0x70, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x3a, 0x72, 0x65, 0x61, 0x64, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x0b, 0x12, 0x09, 0x2f, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x42,
	0x2d, 0x5a, 0x2b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6f,
	0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x67, 0x6f, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x65, 0x6f, 0x6c, 0x79,
	0x6d, 0x70, 0x2f, 0x61, 0x74, 0x6c, 0x61, 0x73, 0x3b, 0x61, 0x74, 0x6c, 0x61, 0x73, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_eolymp_atlas_code_template_service_proto_rawDescOnce sync.Once
	file_eolymp_atlas_code_template_service_proto_rawDescData = file_eolymp_atlas_code_template_service_proto_rawDesc
)

func file_eolymp_atlas_code_template_service_proto_rawDescGZIP() []byte {
	file_eolymp_atlas_code_template_service_proto_rawDescOnce.Do(func() {
		file_eolymp_atlas_code_template_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_eolymp_atlas_code_template_service_proto_rawDescData)
	})
	return file_eolymp_atlas_code_template_service_proto_rawDescData
}

var file_eolymp_atlas_code_template_service_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_eolymp_atlas_code_template_service_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_eolymp_atlas_code_template_service_proto_goTypes = []any{
	(UpdateCodeTemplateInput_Patch)(0), // 0: eolymp.atlas.UpdateCodeTemplateInput.Patch
	(*CreateCodeTemplateInput)(nil),    // 1: eolymp.atlas.CreateCodeTemplateInput
	(*CreateCodeTemplateOutput)(nil),   // 2: eolymp.atlas.CreateCodeTemplateOutput
	(*UpdateCodeTemplateInput)(nil),    // 3: eolymp.atlas.UpdateCodeTemplateInput
	(*UpdateCodeTemplateOutput)(nil),   // 4: eolymp.atlas.UpdateCodeTemplateOutput
	(*DeleteCodeTemplateInput)(nil),    // 5: eolymp.atlas.DeleteCodeTemplateInput
	(*DeleteCodeTemplateOutput)(nil),   // 6: eolymp.atlas.DeleteCodeTemplateOutput
	(*ListCodeTemplatesInput)(nil),     // 7: eolymp.atlas.ListCodeTemplatesInput
	(*ListCodeTemplatesOutput)(nil),    // 8: eolymp.atlas.ListCodeTemplatesOutput
	(*DescribeCodeTemplateInput)(nil),  // 9: eolymp.atlas.DescribeCodeTemplateInput
	(*DescribeCodeTemplateOutput)(nil), // 10: eolymp.atlas.DescribeCodeTemplateOutput
	(*LookupCodeTemplateInput)(nil),    // 11: eolymp.atlas.LookupCodeTemplateInput
	(*LookupCodeTemplateOutput)(nil),   // 12: eolymp.atlas.LookupCodeTemplateOutput
	(*Template)(nil),                   // 13: eolymp.atlas.Template
}
var file_eolymp_atlas_code_template_service_proto_depIdxs = []int32{
	13, // 0: eolymp.atlas.CreateCodeTemplateInput.template:type_name -> eolymp.atlas.Template
	13, // 1: eolymp.atlas.UpdateCodeTemplateInput.template:type_name -> eolymp.atlas.Template
	13, // 2: eolymp.atlas.ListCodeTemplatesOutput.items:type_name -> eolymp.atlas.Template
	13, // 3: eolymp.atlas.DescribeCodeTemplateOutput.template:type_name -> eolymp.atlas.Template
	13, // 4: eolymp.atlas.LookupCodeTemplateOutput.template:type_name -> eolymp.atlas.Template
	1,  // 5: eolymp.atlas.CodeTemplateService.CreateCodeTemplate:input_type -> eolymp.atlas.CreateCodeTemplateInput
	3,  // 6: eolymp.atlas.CodeTemplateService.UpdateCodeTemplate:input_type -> eolymp.atlas.UpdateCodeTemplateInput
	5,  // 7: eolymp.atlas.CodeTemplateService.DeleteCodeTemplate:input_type -> eolymp.atlas.DeleteCodeTemplateInput
	7,  // 8: eolymp.atlas.CodeTemplateService.ListCodeTemplates:input_type -> eolymp.atlas.ListCodeTemplatesInput
	9,  // 9: eolymp.atlas.CodeTemplateService.DescribeCodeTemplate:input_type -> eolymp.atlas.DescribeCodeTemplateInput
	11, // 10: eolymp.atlas.CodeTemplateService.LookupCodeTemplate:input_type -> eolymp.atlas.LookupCodeTemplateInput
	2,  // 11: eolymp.atlas.CodeTemplateService.CreateCodeTemplate:output_type -> eolymp.atlas.CreateCodeTemplateOutput
	4,  // 12: eolymp.atlas.CodeTemplateService.UpdateCodeTemplate:output_type -> eolymp.atlas.UpdateCodeTemplateOutput
	6,  // 13: eolymp.atlas.CodeTemplateService.DeleteCodeTemplate:output_type -> eolymp.atlas.DeleteCodeTemplateOutput
	8,  // 14: eolymp.atlas.CodeTemplateService.ListCodeTemplates:output_type -> eolymp.atlas.ListCodeTemplatesOutput
	10, // 15: eolymp.atlas.CodeTemplateService.DescribeCodeTemplate:output_type -> eolymp.atlas.DescribeCodeTemplateOutput
	12, // 16: eolymp.atlas.CodeTemplateService.LookupCodeTemplate:output_type -> eolymp.atlas.LookupCodeTemplateOutput
	11, // [11:17] is the sub-list for method output_type
	5,  // [5:11] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_eolymp_atlas_code_template_service_proto_init() }
func file_eolymp_atlas_code_template_service_proto_init() {
	if File_eolymp_atlas_code_template_service_proto != nil {
		return
	}
	file_eolymp_atlas_code_template_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_eolymp_atlas_code_template_service_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_eolymp_atlas_code_template_service_proto_goTypes,
		DependencyIndexes: file_eolymp_atlas_code_template_service_proto_depIdxs,
		EnumInfos:         file_eolymp_atlas_code_template_service_proto_enumTypes,
		MessageInfos:      file_eolymp_atlas_code_template_service_proto_msgTypes,
	}.Build()
	File_eolymp_atlas_code_template_service_proto = out.File
	file_eolymp_atlas_code_template_service_proto_rawDesc = nil
	file_eolymp_atlas_code_template_service_proto_goTypes = nil
	file_eolymp_atlas_code_template_service_proto_depIdxs = nil
}
