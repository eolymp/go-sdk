// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v3.21.12
// source: eolymp/executor/language_service.proto

package executor

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

type DescribeLanguageInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LanguageId string `protobuf:"bytes,1,opt,name=language_id,json=languageId,proto3" json:"language_id,omitempty"`
}

func (x *DescribeLanguageInput) Reset() {
	*x = DescribeLanguageInput{}
	mi := &file_eolymp_executor_language_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DescribeLanguageInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescribeLanguageInput) ProtoMessage() {}

func (x *DescribeLanguageInput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_executor_language_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DescribeLanguageInput.ProtoReflect.Descriptor instead.
func (*DescribeLanguageInput) Descriptor() ([]byte, []int) {
	return file_eolymp_executor_language_service_proto_rawDescGZIP(), []int{0}
}

func (x *DescribeLanguageInput) GetLanguageId() string {
	if x != nil {
		return x.LanguageId
	}
	return ""
}

type DescribeLanguageOutput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Language *Language `protobuf:"bytes,1,opt,name=language,proto3" json:"language,omitempty"`
}

func (x *DescribeLanguageOutput) Reset() {
	*x = DescribeLanguageOutput{}
	mi := &file_eolymp_executor_language_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DescribeLanguageOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescribeLanguageOutput) ProtoMessage() {}

func (x *DescribeLanguageOutput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_executor_language_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DescribeLanguageOutput.ProtoReflect.Descriptor instead.
func (*DescribeLanguageOutput) Descriptor() ([]byte, []int) {
	return file_eolymp_executor_language_service_proto_rawDescGZIP(), []int{1}
}

func (x *DescribeLanguageOutput) GetLanguage() *Language {
	if x != nil {
		return x.Language
	}
	return nil
}

type ListLanguagesInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Filters *ListLanguagesInput_Filter `protobuf:"bytes,40,opt,name=filters,proto3" json:"filters,omitempty"`
}

func (x *ListLanguagesInput) Reset() {
	*x = ListLanguagesInput{}
	mi := &file_eolymp_executor_language_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListLanguagesInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListLanguagesInput) ProtoMessage() {}

func (x *ListLanguagesInput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_executor_language_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListLanguagesInput.ProtoReflect.Descriptor instead.
func (*ListLanguagesInput) Descriptor() ([]byte, []int) {
	return file_eolymp_executor_language_service_proto_rawDescGZIP(), []int{2}
}

func (x *ListLanguagesInput) GetFilters() *ListLanguagesInput_Filter {
	if x != nil {
		return x.Filters
	}
	return nil
}

type ListLanguagesOutput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items []*Language `protobuf:"bytes,2,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *ListLanguagesOutput) Reset() {
	*x = ListLanguagesOutput{}
	mi := &file_eolymp_executor_language_service_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListLanguagesOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListLanguagesOutput) ProtoMessage() {}

func (x *ListLanguagesOutput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_executor_language_service_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListLanguagesOutput.ProtoReflect.Descriptor instead.
func (*ListLanguagesOutput) Descriptor() ([]byte, []int) {
	return file_eolymp_executor_language_service_proto_rawDescGZIP(), []int{3}
}

func (x *ListLanguagesOutput) GetItems() []*Language {
	if x != nil {
		return x.Items
	}
	return nil
}

type DescribeRuntimeInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RuntimeId string `protobuf:"bytes,1,opt,name=runtime_id,json=runtimeId,proto3" json:"runtime_id,omitempty"`
}

func (x *DescribeRuntimeInput) Reset() {
	*x = DescribeRuntimeInput{}
	mi := &file_eolymp_executor_language_service_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DescribeRuntimeInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescribeRuntimeInput) ProtoMessage() {}

func (x *DescribeRuntimeInput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_executor_language_service_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DescribeRuntimeInput.ProtoReflect.Descriptor instead.
func (*DescribeRuntimeInput) Descriptor() ([]byte, []int) {
	return file_eolymp_executor_language_service_proto_rawDescGZIP(), []int{4}
}

func (x *DescribeRuntimeInput) GetRuntimeId() string {
	if x != nil {
		return x.RuntimeId
	}
	return ""
}

type DescribeRuntimeOutput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Runtime *Runtime `protobuf:"bytes,1,opt,name=runtime,proto3" json:"runtime,omitempty"`
}

func (x *DescribeRuntimeOutput) Reset() {
	*x = DescribeRuntimeOutput{}
	mi := &file_eolymp_executor_language_service_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DescribeRuntimeOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescribeRuntimeOutput) ProtoMessage() {}

func (x *DescribeRuntimeOutput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_executor_language_service_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DescribeRuntimeOutput.ProtoReflect.Descriptor instead.
func (*DescribeRuntimeOutput) Descriptor() ([]byte, []int) {
	return file_eolymp_executor_language_service_proto_rawDescGZIP(), []int{5}
}

func (x *DescribeRuntimeOutput) GetRuntime() *Runtime {
	if x != nil {
		return x.Runtime
	}
	return nil
}

type ListRuntimeInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Filters *ListRuntimeInput_Filter `protobuf:"bytes,40,opt,name=filters,proto3" json:"filters,omitempty"`
}

func (x *ListRuntimeInput) Reset() {
	*x = ListRuntimeInput{}
	mi := &file_eolymp_executor_language_service_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListRuntimeInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListRuntimeInput) ProtoMessage() {}

func (x *ListRuntimeInput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_executor_language_service_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListRuntimeInput.ProtoReflect.Descriptor instead.
func (*ListRuntimeInput) Descriptor() ([]byte, []int) {
	return file_eolymp_executor_language_service_proto_rawDescGZIP(), []int{6}
}

func (x *ListRuntimeInput) GetFilters() *ListRuntimeInput_Filter {
	if x != nil {
		return x.Filters
	}
	return nil
}

type ListRuntimeOutput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items []*Runtime `protobuf:"bytes,2,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *ListRuntimeOutput) Reset() {
	*x = ListRuntimeOutput{}
	mi := &file_eolymp_executor_language_service_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListRuntimeOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListRuntimeOutput) ProtoMessage() {}

func (x *ListRuntimeOutput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_executor_language_service_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListRuntimeOutput.ProtoReflect.Descriptor instead.
func (*ListRuntimeOutput) Descriptor() ([]byte, []int) {
	return file_eolymp_executor_language_service_proto_rawDescGZIP(), []int{7}
}

func (x *ListRuntimeOutput) GetItems() []*Runtime {
	if x != nil {
		return x.Items
	}
	return nil
}

type DescribeCodeTemplateInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RuntimeId string `protobuf:"bytes,1,opt,name=runtime_id,json=runtimeId,proto3" json:"runtime_id,omitempty"`
}

func (x *DescribeCodeTemplateInput) Reset() {
	*x = DescribeCodeTemplateInput{}
	mi := &file_eolymp_executor_language_service_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DescribeCodeTemplateInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescribeCodeTemplateInput) ProtoMessage() {}

func (x *DescribeCodeTemplateInput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_executor_language_service_proto_msgTypes[8]
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
	return file_eolymp_executor_language_service_proto_rawDescGZIP(), []int{8}
}

func (x *DescribeCodeTemplateInput) GetRuntimeId() string {
	if x != nil {
		return x.RuntimeId
	}
	return ""
}

type DescribeCodeTemplateOutput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Template string `protobuf:"bytes,1,opt,name=template,proto3" json:"template,omitempty"`
}

func (x *DescribeCodeTemplateOutput) Reset() {
	*x = DescribeCodeTemplateOutput{}
	mi := &file_eolymp_executor_language_service_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DescribeCodeTemplateOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescribeCodeTemplateOutput) ProtoMessage() {}

func (x *DescribeCodeTemplateOutput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_executor_language_service_proto_msgTypes[9]
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
	return file_eolymp_executor_language_service_proto_rawDescGZIP(), []int{9}
}

func (x *DescribeCodeTemplateOutput) GetTemplate() string {
	if x != nil {
		return x.Template
	}
	return ""
}

type ListLanguagesInput_Filter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         []*wellknown.ExpressionID     `protobuf:"bytes,1,rep,name=id,proto3" json:"id,omitempty"`
	Name       []*wellknown.ExpressionString `protobuf:"bytes,2,rep,name=name,proto3" json:"name,omitempty"`
	Deprecated []*wellknown.ExpressionBool   `protobuf:"bytes,3,rep,name=deprecated,proto3" json:"deprecated,omitempty"`
}

func (x *ListLanguagesInput_Filter) Reset() {
	*x = ListLanguagesInput_Filter{}
	mi := &file_eolymp_executor_language_service_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListLanguagesInput_Filter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListLanguagesInput_Filter) ProtoMessage() {}

func (x *ListLanguagesInput_Filter) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_executor_language_service_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListLanguagesInput_Filter.ProtoReflect.Descriptor instead.
func (*ListLanguagesInput_Filter) Descriptor() ([]byte, []int) {
	return file_eolymp_executor_language_service_proto_rawDescGZIP(), []int{2, 0}
}

func (x *ListLanguagesInput_Filter) GetId() []*wellknown.ExpressionID {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *ListLanguagesInput_Filter) GetName() []*wellknown.ExpressionString {
	if x != nil {
		return x.Name
	}
	return nil
}

func (x *ListLanguagesInput_Filter) GetDeprecated() []*wellknown.ExpressionBool {
	if x != nil {
		return x.Deprecated
	}
	return nil
}

type ListRuntimeInput_Filter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         []*wellknown.ExpressionID     `protobuf:"bytes,1,rep,name=id,proto3" json:"id,omitempty"`
	Lang       []*wellknown.ExpressionID     `protobuf:"bytes,2,rep,name=lang,proto3" json:"lang,omitempty"`
	Version    []*wellknown.ExpressionID     `protobuf:"bytes,3,rep,name=version,proto3" json:"version,omitempty"`
	Name       []*wellknown.ExpressionString `protobuf:"bytes,4,rep,name=name,proto3" json:"name,omitempty"`
	Deprecated []*wellknown.ExpressionBool   `protobuf:"bytes,5,rep,name=deprecated,proto3" json:"deprecated,omitempty"`
}

func (x *ListRuntimeInput_Filter) Reset() {
	*x = ListRuntimeInput_Filter{}
	mi := &file_eolymp_executor_language_service_proto_msgTypes[11]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListRuntimeInput_Filter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListRuntimeInput_Filter) ProtoMessage() {}

func (x *ListRuntimeInput_Filter) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_executor_language_service_proto_msgTypes[11]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListRuntimeInput_Filter.ProtoReflect.Descriptor instead.
func (*ListRuntimeInput_Filter) Descriptor() ([]byte, []int) {
	return file_eolymp_executor_language_service_proto_rawDescGZIP(), []int{6, 0}
}

func (x *ListRuntimeInput_Filter) GetId() []*wellknown.ExpressionID {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *ListRuntimeInput_Filter) GetLang() []*wellknown.ExpressionID {
	if x != nil {
		return x.Lang
	}
	return nil
}

func (x *ListRuntimeInput_Filter) GetVersion() []*wellknown.ExpressionID {
	if x != nil {
		return x.Version
	}
	return nil
}

func (x *ListRuntimeInput_Filter) GetName() []*wellknown.ExpressionString {
	if x != nil {
		return x.Name
	}
	return nil
}

func (x *ListRuntimeInput_Filter) GetDeprecated() []*wellknown.ExpressionBool {
	if x != nil {
		return x.Deprecated
	}
	return nil
}

var File_eolymp_executor_language_service_proto protoreflect.FileDescriptor

var file_eolymp_executor_language_service_proto_rawDesc = []byte{
	0x0a, 0x26, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x6f,
	0x72, 0x2f, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70,
	0x2e, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x6f, 0x72, 0x1a, 0x1d, 0x65, 0x6f, 0x6c, 0x79, 0x6d,
	0x70, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x68, 0x74,
	0x74, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x22, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70,
	0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x72, 0x61, 0x74,
	0x65, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x65, 0x6f,
	0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2f, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x65, 0x6f,
	0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x6f, 0x72, 0x2f, 0x6c, 0x61,
	0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x65, 0x6f,
	0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x6f, 0x72, 0x2f, 0x72, 0x75,
	0x6e, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x21, 0x65, 0x6f, 0x6c,
	0x79, 0x6d, 0x70, 0x2f, 0x77, 0x65, 0x6c, 0x6c, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x2f, 0x65, 0x78,
	0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x38,
	0x0a, 0x15, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61,
	0x67, 0x65, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x6c, 0x61, 0x6e, 0x67, 0x75,
	0x61, 0x67, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6c, 0x61,
	0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x49, 0x64, 0x22, 0x4f, 0x0a, 0x16, 0x44, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x62, 0x65, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x4f, 0x75, 0x74, 0x70,
	0x75, 0x74, 0x12, 0x35, 0x0a, 0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x65, 0x78,
	0x65, 0x63, 0x75, 0x74, 0x6f, 0x72, 0x2e, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x52,
	0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x22, 0x8f, 0x02, 0x0a, 0x12, 0x4c, 0x69,
	0x73, 0x74, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x73, 0x49, 0x6e, 0x70, 0x75, 0x74,
	0x12, 0x44, 0x0a, 0x07, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x18, 0x28, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x2a, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x65, 0x78, 0x65, 0x63, 0x75,
	0x74, 0x6f, 0x72, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65,
	0x73, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x2e, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x52, 0x07, 0x66,
	0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x1a, 0xb2, 0x01, 0x0a, 0x06, 0x46, 0x69, 0x6c, 0x74, 0x65,
	0x72, 0x12, 0x2e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e,
	0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x77, 0x65, 0x6c, 0x6c, 0x6b, 0x6e, 0x6f, 0x77, 0x6e,
	0x2e, 0x45, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x36, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x22, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x77, 0x65, 0x6c, 0x6c, 0x6b, 0x6e, 0x6f,
	0x77, 0x6e, 0x2e, 0x45, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x72,
	0x69, 0x6e, 0x67, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x40, 0x0a, 0x0a, 0x64, 0x65, 0x70,
	0x72, 0x65, 0x63, 0x61, 0x74, 0x65, 0x64, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e,
	0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x77, 0x65, 0x6c, 0x6c, 0x6b, 0x6e, 0x6f, 0x77, 0x6e,
	0x2e, 0x45, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x42, 0x6f, 0x6f, 0x6c, 0x52,
	0x0a, 0x64, 0x65, 0x70, 0x72, 0x65, 0x63, 0x61, 0x74, 0x65, 0x64, 0x22, 0x46, 0x0a, 0x13, 0x4c,
	0x69, 0x73, 0x74, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x73, 0x4f, 0x75, 0x74, 0x70,
	0x75, 0x74, 0x12, 0x2f, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x19, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x65, 0x78, 0x65, 0x63, 0x75,
	0x74, 0x6f, 0x72, 0x2e, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x52, 0x05, 0x69, 0x74,
	0x65, 0x6d, 0x73, 0x22, 0x35, 0x0a, 0x14, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x52,
	0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x72,
	0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x49, 0x64, 0x22, 0x4b, 0x0a, 0x15, 0x44, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x52, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x4f, 0x75, 0x74,
	0x70, 0x75, 0x74, 0x12, 0x32, 0x0a, 0x07, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x65, 0x78,
	0x65, 0x63, 0x75, 0x74, 0x6f, 0x72, 0x2e, 0x52, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x52, 0x07,
	0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x22, 0xf9, 0x02, 0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74,
	0x52, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x42, 0x0a, 0x07,
	0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x18, 0x28, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e,
	0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x6f, 0x72, 0x2e,
	0x4c, 0x69, 0x73, 0x74, 0x52, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x49, 0x6e, 0x70, 0x75, 0x74,
	0x2e, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x52, 0x07, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73,
	0x1a, 0xa0, 0x02, 0x0a, 0x06, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x2e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70,
	0x2e, 0x77, 0x65, 0x6c, 0x6c, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x2e, 0x45, 0x78, 0x70, 0x72, 0x65,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x52, 0x02, 0x69, 0x64, 0x12, 0x32, 0x0a, 0x04, 0x6c,
	0x61, 0x6e, 0x67, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x65, 0x6f, 0x6c, 0x79,
	0x6d, 0x70, 0x2e, 0x77, 0x65, 0x6c, 0x6c, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x2e, 0x45, 0x78, 0x70,
	0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x52, 0x04, 0x6c, 0x61, 0x6e, 0x67, 0x12,
	0x38, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x1e, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x77, 0x65, 0x6c, 0x6c, 0x6b, 0x6e,
	0x6f, 0x77, 0x6e, 0x2e, 0x45, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x44,
	0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x36, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70,
	0x2e, 0x77, 0x65, 0x6c, 0x6c, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x2e, 0x45, 0x78, 0x70, 0x72, 0x65,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x40, 0x0a, 0x0a, 0x64, 0x65, 0x70, 0x72, 0x65, 0x63, 0x61, 0x74, 0x65, 0x64, 0x18,
	0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x77,
	0x65, 0x6c, 0x6c, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x2e, 0x45, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x42, 0x6f, 0x6f, 0x6c, 0x52, 0x0a, 0x64, 0x65, 0x70, 0x72, 0x65, 0x63, 0x61,
	0x74, 0x65, 0x64, 0x22, 0x43, 0x0a, 0x11, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x75, 0x6e, 0x74, 0x69,
	0x6d, 0x65, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x12, 0x2e, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d,
	0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70,
	0x2e, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x6f, 0x72, 0x2e, 0x52, 0x75, 0x6e, 0x74, 0x69, 0x6d,
	0x65, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x22, 0x3a, 0x0a, 0x19, 0x44, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x62, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65,
	0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x72, 0x75, 0x6e, 0x74, 0x69,
	0x6d, 0x65, 0x49, 0x64, 0x22, 0x38, 0x0a, 0x1a, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65,
	0x43, 0x6f, 0x64, 0x65, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x4f, 0x75, 0x74, 0x70,
	0x75, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x32, 0x89,
	0x07, 0x0a, 0x0f, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0xb7, 0x01, 0x0a, 0x10, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x4c,
	0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x12, 0x26, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70,
	0x2e, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x6f, 0x72, 0x2e, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x62, 0x65, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a,
	0x27, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x6f,
	0x72, 0x2e, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61,
	0x67, 0x65, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x22, 0x52, 0xea, 0xe2, 0x0a, 0x0c, 0xf5, 0xe2,
	0x0a, 0x00, 0x00, 0xa0, 0x41, 0xf8, 0xe2, 0x0a, 0xc8, 0x01, 0x82, 0xe3, 0x0a, 0x19, 0x8a, 0xe3,
	0x0a, 0x15, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x6f, 0x72, 0x3a, 0x72, 0x75, 0x6e, 0x74, 0x69,
	0x6d, 0x65, 0x3a, 0x72, 0x65, 0x61, 0x64, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1f, 0x12, 0x1d, 0x2f,
	0x65, 0x78, 0x65, 0x63, 0x2f, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x73, 0x2f, 0x7b,
	0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x5f, 0x69, 0x64, 0x7d, 0x12, 0xa0, 0x01, 0x0a,
	0x0d, 0x4c, 0x69, 0x73, 0x74, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x73, 0x12, 0x23,
	0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x6f, 0x72,
	0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x73, 0x49, 0x6e,
	0x70, 0x75, 0x74, 0x1a, 0x24, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x65, 0x78, 0x65,
	0x63, 0x75, 0x74, 0x6f, 0x72, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61,
	0x67, 0x65, 0x73, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x22, 0x44, 0xea, 0xe2, 0x0a, 0x0c, 0xf5,
	0xe2, 0x0a, 0x00, 0x00, 0xa0, 0x41, 0xf8, 0xe2, 0x0a, 0xc8, 0x01, 0x82, 0xe3, 0x0a, 0x19, 0x8a,
	0xe3, 0x0a, 0x15, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x6f, 0x72, 0x3a, 0x72, 0x75, 0x6e, 0x74,
	0x69, 0x6d, 0x65, 0x3a, 0x72, 0x65, 0x61, 0x64, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x11, 0x12, 0x0f,
	0x2f, 0x65, 0x78, 0x65, 0x63, 0x2f, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x73, 0x12,
	0xb1, 0x01, 0x0a, 0x0f, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x52, 0x75, 0x6e, 0x74,
	0x69, 0x6d, 0x65, 0x12, 0x25, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x65, 0x78, 0x65,
	0x63, 0x75, 0x74, 0x6f, 0x72, 0x2e, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x52, 0x75,
	0x6e, 0x74, 0x69, 0x6d, 0x65, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x26, 0x2e, 0x65, 0x6f, 0x6c,
	0x79, 0x6d, 0x70, 0x2e, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x6f, 0x72, 0x2e, 0x44, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x62, 0x65, 0x52, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x4f, 0x75, 0x74, 0x70,
	0x75, 0x74, 0x22, 0x4f, 0xea, 0xe2, 0x0a, 0x0c, 0xf5, 0xe2, 0x0a, 0x00, 0x00, 0xa0, 0x41, 0xf8,
	0xe2, 0x0a, 0xc8, 0x01, 0x82, 0xe3, 0x0a, 0x19, 0x8a, 0xe3, 0x0a, 0x15, 0x65, 0x78, 0x65, 0x63,
	0x75, 0x74, 0x6f, 0x72, 0x3a, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x3a, 0x72, 0x65, 0x61,
	0x64, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1c, 0x12, 0x1a, 0x2f, 0x65, 0x78, 0x65, 0x63, 0x2f, 0x72,
	0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x2f, 0x7b, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x5f,
	0x69, 0x64, 0x7d, 0x12, 0x98, 0x01, 0x0a, 0x0b, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x75, 0x6e, 0x74,
	0x69, 0x6d, 0x65, 0x12, 0x21, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x65, 0x78, 0x65,
	0x63, 0x75, 0x74, 0x6f, 0x72, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x75, 0x6e, 0x74, 0x69, 0x6d,
	0x65, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x22, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e,
	0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x6f, 0x72, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x75, 0x6e,
	0x74, 0x69, 0x6d, 0x65, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x22, 0x42, 0xea, 0xe2, 0x0a, 0x0c,
	0xf5, 0xe2, 0x0a, 0x00, 0x00, 0xa0, 0x41, 0xf8, 0xe2, 0x0a, 0xc8, 0x01, 0x82, 0xe3, 0x0a, 0x19,
	0x8a, 0xe3, 0x0a, 0x15, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x6f, 0x72, 0x3a, 0x72, 0x75, 0x6e,
	0x74, 0x69, 0x6d, 0x65, 0x3a, 0x72, 0x65, 0x61, 0x64, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0f, 0x12,
	0x0d, 0x2f, 0x65, 0x78, 0x65, 0x63, 0x2f, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x12, 0xc9,
	0x01, 0x0a, 0x14, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x54,
	0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x12, 0x2a, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70,
	0x2e, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x6f, 0x72, 0x2e, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x62, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x49, 0x6e,
	0x70, 0x75, 0x74, 0x1a, 0x2b, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x65, 0x78, 0x65,
	0x63, 0x75, 0x74, 0x6f, 0x72, 0x2e, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x43, 0x6f,
	0x64, 0x65, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74,
	0x22, 0x58, 0xea, 0xe2, 0x0a, 0x0c, 0xf5, 0xe2, 0x0a, 0x00, 0x00, 0xa0, 0x41, 0xf8, 0xe2, 0x0a,
	0xc8, 0x01, 0x82, 0xe3, 0x0a, 0x19, 0x8a, 0xe3, 0x0a, 0x15, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74,
	0x6f, 0x72, 0x3a, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x3a, 0x72, 0x65, 0x61, 0x64, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x25, 0x12, 0x23, 0x2f, 0x65, 0x78, 0x65, 0x63, 0x2f, 0x72, 0x75, 0x6e,
	0x74, 0x69, 0x6d, 0x65, 0x2f, 0x7b, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x69, 0x64,
	0x7d, 0x2f, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x42, 0x33, 0x5a, 0x31, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f,
	0x67, 0x6f, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x65, 0x78,
	0x65, 0x63, 0x75, 0x74, 0x6f, 0x72, 0x3b, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x6f, 0x72, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_eolymp_executor_language_service_proto_rawDescOnce sync.Once
	file_eolymp_executor_language_service_proto_rawDescData = file_eolymp_executor_language_service_proto_rawDesc
)

func file_eolymp_executor_language_service_proto_rawDescGZIP() []byte {
	file_eolymp_executor_language_service_proto_rawDescOnce.Do(func() {
		file_eolymp_executor_language_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_eolymp_executor_language_service_proto_rawDescData)
	})
	return file_eolymp_executor_language_service_proto_rawDescData
}

var file_eolymp_executor_language_service_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_eolymp_executor_language_service_proto_goTypes = []any{
	(*DescribeLanguageInput)(nil),      // 0: eolymp.executor.DescribeLanguageInput
	(*DescribeLanguageOutput)(nil),     // 1: eolymp.executor.DescribeLanguageOutput
	(*ListLanguagesInput)(nil),         // 2: eolymp.executor.ListLanguagesInput
	(*ListLanguagesOutput)(nil),        // 3: eolymp.executor.ListLanguagesOutput
	(*DescribeRuntimeInput)(nil),       // 4: eolymp.executor.DescribeRuntimeInput
	(*DescribeRuntimeOutput)(nil),      // 5: eolymp.executor.DescribeRuntimeOutput
	(*ListRuntimeInput)(nil),           // 6: eolymp.executor.ListRuntimeInput
	(*ListRuntimeOutput)(nil),          // 7: eolymp.executor.ListRuntimeOutput
	(*DescribeCodeTemplateInput)(nil),  // 8: eolymp.executor.DescribeCodeTemplateInput
	(*DescribeCodeTemplateOutput)(nil), // 9: eolymp.executor.DescribeCodeTemplateOutput
	(*ListLanguagesInput_Filter)(nil),  // 10: eolymp.executor.ListLanguagesInput.Filter
	(*ListRuntimeInput_Filter)(nil),    // 11: eolymp.executor.ListRuntimeInput.Filter
	(*Language)(nil),                   // 12: eolymp.executor.Language
	(*Runtime)(nil),                    // 13: eolymp.executor.Runtime
	(*wellknown.ExpressionID)(nil),     // 14: eolymp.wellknown.ExpressionID
	(*wellknown.ExpressionString)(nil), // 15: eolymp.wellknown.ExpressionString
	(*wellknown.ExpressionBool)(nil),   // 16: eolymp.wellknown.ExpressionBool
}
var file_eolymp_executor_language_service_proto_depIdxs = []int32{
	12, // 0: eolymp.executor.DescribeLanguageOutput.language:type_name -> eolymp.executor.Language
	10, // 1: eolymp.executor.ListLanguagesInput.filters:type_name -> eolymp.executor.ListLanguagesInput.Filter
	12, // 2: eolymp.executor.ListLanguagesOutput.items:type_name -> eolymp.executor.Language
	13, // 3: eolymp.executor.DescribeRuntimeOutput.runtime:type_name -> eolymp.executor.Runtime
	11, // 4: eolymp.executor.ListRuntimeInput.filters:type_name -> eolymp.executor.ListRuntimeInput.Filter
	13, // 5: eolymp.executor.ListRuntimeOutput.items:type_name -> eolymp.executor.Runtime
	14, // 6: eolymp.executor.ListLanguagesInput.Filter.id:type_name -> eolymp.wellknown.ExpressionID
	15, // 7: eolymp.executor.ListLanguagesInput.Filter.name:type_name -> eolymp.wellknown.ExpressionString
	16, // 8: eolymp.executor.ListLanguagesInput.Filter.deprecated:type_name -> eolymp.wellknown.ExpressionBool
	14, // 9: eolymp.executor.ListRuntimeInput.Filter.id:type_name -> eolymp.wellknown.ExpressionID
	14, // 10: eolymp.executor.ListRuntimeInput.Filter.lang:type_name -> eolymp.wellknown.ExpressionID
	14, // 11: eolymp.executor.ListRuntimeInput.Filter.version:type_name -> eolymp.wellknown.ExpressionID
	15, // 12: eolymp.executor.ListRuntimeInput.Filter.name:type_name -> eolymp.wellknown.ExpressionString
	16, // 13: eolymp.executor.ListRuntimeInput.Filter.deprecated:type_name -> eolymp.wellknown.ExpressionBool
	0,  // 14: eolymp.executor.LanguageService.DescribeLanguage:input_type -> eolymp.executor.DescribeLanguageInput
	2,  // 15: eolymp.executor.LanguageService.ListLanguages:input_type -> eolymp.executor.ListLanguagesInput
	4,  // 16: eolymp.executor.LanguageService.DescribeRuntime:input_type -> eolymp.executor.DescribeRuntimeInput
	6,  // 17: eolymp.executor.LanguageService.ListRuntime:input_type -> eolymp.executor.ListRuntimeInput
	8,  // 18: eolymp.executor.LanguageService.DescribeCodeTemplate:input_type -> eolymp.executor.DescribeCodeTemplateInput
	1,  // 19: eolymp.executor.LanguageService.DescribeLanguage:output_type -> eolymp.executor.DescribeLanguageOutput
	3,  // 20: eolymp.executor.LanguageService.ListLanguages:output_type -> eolymp.executor.ListLanguagesOutput
	5,  // 21: eolymp.executor.LanguageService.DescribeRuntime:output_type -> eolymp.executor.DescribeRuntimeOutput
	7,  // 22: eolymp.executor.LanguageService.ListRuntime:output_type -> eolymp.executor.ListRuntimeOutput
	9,  // 23: eolymp.executor.LanguageService.DescribeCodeTemplate:output_type -> eolymp.executor.DescribeCodeTemplateOutput
	19, // [19:24] is the sub-list for method output_type
	14, // [14:19] is the sub-list for method input_type
	14, // [14:14] is the sub-list for extension type_name
	14, // [14:14] is the sub-list for extension extendee
	0,  // [0:14] is the sub-list for field type_name
}

func init() { file_eolymp_executor_language_service_proto_init() }
func file_eolymp_executor_language_service_proto_init() {
	if File_eolymp_executor_language_service_proto != nil {
		return
	}
	file_eolymp_executor_language_proto_init()
	file_eolymp_executor_runtime_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_eolymp_executor_language_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_eolymp_executor_language_service_proto_goTypes,
		DependencyIndexes: file_eolymp_executor_language_service_proto_depIdxs,
		MessageInfos:      file_eolymp_executor_language_service_proto_msgTypes,
	}.Build()
	File_eolymp_executor_language_service_proto = out.File
	file_eolymp_executor_language_service_proto_rawDesc = nil
	file_eolymp_executor_language_service_proto_goTypes = nil
	file_eolymp_executor_language_service_proto_depIdxs = nil
}
