// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v4.24.4
// source: eolymp/atlas/editor_service.proto

package atlas

import (
	_ "github.com/eolymp/go-sdk/eolymp/annotations"
	runtime "github.com/eolymp/go-sdk/eolymp/runtime"
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

type Editor_Feature int32

const (
	Editor_UNKNOWN_FEATURE Editor_Feature = 0
	Editor_PRINT_CODE      Editor_Feature = 1
)

// Enum value maps for Editor_Feature.
var (
	Editor_Feature_name = map[int32]string{
		0: "UNKNOWN_FEATURE",
		1: "PRINT_CODE",
	}
	Editor_Feature_value = map[string]int32{
		"UNKNOWN_FEATURE": 0,
		"PRINT_CODE":      1,
	}
)

func (x Editor_Feature) Enum() *Editor_Feature {
	p := new(Editor_Feature)
	*p = x
	return p
}

func (x Editor_Feature) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Editor_Feature) Descriptor() protoreflect.EnumDescriptor {
	return file_eolymp_atlas_editor_service_proto_enumTypes[0].Descriptor()
}

func (Editor_Feature) Type() protoreflect.EnumType {
	return &file_eolymp_atlas_editor_service_proto_enumTypes[0]
}

func (x Editor_Feature) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Editor_Feature.Descriptor instead.
func (Editor_Feature) EnumDescriptor() ([]byte, []int) {
	return file_eolymp_atlas_editor_service_proto_rawDescGZIP(), []int{0, 0}
}

type Editor struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	State         *Editor_State          `protobuf:"bytes,1,opt,name=state,proto3" json:"state,omitempty"`                                                 // current state of the editor
	Features      []Editor_Feature       `protobuf:"varint,10,rep,packed,name=features,proto3,enum=eolymp.atlas.Editor_Feature" json:"features,omitempty"` // list of enabled features
	Runtimes      []*runtime.Runtime     `protobuf:"bytes,11,rep,name=runtimes,proto3" json:"runtimes,omitempty"`                                          // list of available runtimes
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Editor) Reset() {
	*x = Editor{}
	mi := &file_eolymp_atlas_editor_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Editor) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Editor) ProtoMessage() {}

func (x *Editor) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_atlas_editor_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Editor.ProtoReflect.Descriptor instead.
func (*Editor) Descriptor() ([]byte, []int) {
	return file_eolymp_atlas_editor_service_proto_rawDescGZIP(), []int{0}
}

func (x *Editor) GetState() *Editor_State {
	if x != nil {
		return x.State
	}
	return nil
}

func (x *Editor) GetFeatures() []Editor_Feature {
	if x != nil {
		return x.Features
	}
	return nil
}

func (x *Editor) GetRuntimes() []*runtime.Runtime {
	if x != nil {
		return x.Runtimes
	}
	return nil
}

type DescribeEditorInput struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DescribeEditorInput) Reset() {
	*x = DescribeEditorInput{}
	mi := &file_eolymp_atlas_editor_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DescribeEditorInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescribeEditorInput) ProtoMessage() {}

func (x *DescribeEditorInput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_atlas_editor_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DescribeEditorInput.ProtoReflect.Descriptor instead.
func (*DescribeEditorInput) Descriptor() ([]byte, []int) {
	return file_eolymp_atlas_editor_service_proto_rawDescGZIP(), []int{1}
}

type DescribeEditorOutput struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Editor        *Editor                `protobuf:"bytes,1,opt,name=editor,proto3" json:"editor,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DescribeEditorOutput) Reset() {
	*x = DescribeEditorOutput{}
	mi := &file_eolymp_atlas_editor_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DescribeEditorOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescribeEditorOutput) ProtoMessage() {}

func (x *DescribeEditorOutput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_atlas_editor_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DescribeEditorOutput.ProtoReflect.Descriptor instead.
func (*DescribeEditorOutput) Descriptor() ([]byte, []int) {
	return file_eolymp_atlas_editor_service_proto_rawDescGZIP(), []int{2}
}

func (x *DescribeEditorOutput) GetEditor() *Editor {
	if x != nil {
		return x.Editor
	}
	return nil
}

type DescribeEditorStateInput struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DescribeEditorStateInput) Reset() {
	*x = DescribeEditorStateInput{}
	mi := &file_eolymp_atlas_editor_service_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DescribeEditorStateInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescribeEditorStateInput) ProtoMessage() {}

func (x *DescribeEditorStateInput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_atlas_editor_service_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DescribeEditorStateInput.ProtoReflect.Descriptor instead.
func (*DescribeEditorStateInput) Descriptor() ([]byte, []int) {
	return file_eolymp_atlas_editor_service_proto_rawDescGZIP(), []int{3}
}

type DescribeEditorStateOutput struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Runtime       string                 `protobuf:"bytes,1,opt,name=runtime,proto3" json:"runtime,omitempty"`
	SourceCode    string                 `protobuf:"bytes,2,opt,name=source_code,json=sourceCode,proto3" json:"source_code,omitempty"`
	InputData     string                 `protobuf:"bytes,3,opt,name=input_data,json=inputData,proto3" json:"input_data,omitempty"`
	Features      []Editor_Feature       `protobuf:"varint,10,rep,packed,name=features,proto3,enum=eolymp.atlas.Editor_Feature" json:"features,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DescribeEditorStateOutput) Reset() {
	*x = DescribeEditorStateOutput{}
	mi := &file_eolymp_atlas_editor_service_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DescribeEditorStateOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescribeEditorStateOutput) ProtoMessage() {}

func (x *DescribeEditorStateOutput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_atlas_editor_service_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DescribeEditorStateOutput.ProtoReflect.Descriptor instead.
func (*DescribeEditorStateOutput) Descriptor() ([]byte, []int) {
	return file_eolymp_atlas_editor_service_proto_rawDescGZIP(), []int{4}
}

func (x *DescribeEditorStateOutput) GetRuntime() string {
	if x != nil {
		return x.Runtime
	}
	return ""
}

func (x *DescribeEditorStateOutput) GetSourceCode() string {
	if x != nil {
		return x.SourceCode
	}
	return ""
}

func (x *DescribeEditorStateOutput) GetInputData() string {
	if x != nil {
		return x.InputData
	}
	return ""
}

func (x *DescribeEditorStateOutput) GetFeatures() []Editor_Feature {
	if x != nil {
		return x.Features
	}
	return nil
}

type UpdateEditorStateInput struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Runtime       string                 `protobuf:"bytes,1,opt,name=runtime,proto3" json:"runtime,omitempty"`
	SourceCode    string                 `protobuf:"bytes,2,opt,name=source_code,json=sourceCode,proto3" json:"source_code,omitempty"`
	InputData     string                 `protobuf:"bytes,3,opt,name=input_data,json=inputData,proto3" json:"input_data,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateEditorStateInput) Reset() {
	*x = UpdateEditorStateInput{}
	mi := &file_eolymp_atlas_editor_service_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateEditorStateInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateEditorStateInput) ProtoMessage() {}

func (x *UpdateEditorStateInput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_atlas_editor_service_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateEditorStateInput.ProtoReflect.Descriptor instead.
func (*UpdateEditorStateInput) Descriptor() ([]byte, []int) {
	return file_eolymp_atlas_editor_service_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateEditorStateInput) GetRuntime() string {
	if x != nil {
		return x.Runtime
	}
	return ""
}

func (x *UpdateEditorStateInput) GetSourceCode() string {
	if x != nil {
		return x.SourceCode
	}
	return ""
}

func (x *UpdateEditorStateInput) GetInputData() string {
	if x != nil {
		return x.InputData
	}
	return ""
}

type UpdateEditorStateOutput struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateEditorStateOutput) Reset() {
	*x = UpdateEditorStateOutput{}
	mi := &file_eolymp_atlas_editor_service_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateEditorStateOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateEditorStateOutput) ProtoMessage() {}

func (x *UpdateEditorStateOutput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_atlas_editor_service_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateEditorStateOutput.ProtoReflect.Descriptor instead.
func (*UpdateEditorStateOutput) Descriptor() ([]byte, []int) {
	return file_eolymp_atlas_editor_service_proto_rawDescGZIP(), []int{6}
}

type PrintEditorCodeInput struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Runtime       string                 `protobuf:"bytes,1,opt,name=runtime,proto3" json:"runtime,omitempty"`
	SourceCode    string                 `protobuf:"bytes,2,opt,name=source_code,json=sourceCode,proto3" json:"source_code,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PrintEditorCodeInput) Reset() {
	*x = PrintEditorCodeInput{}
	mi := &file_eolymp_atlas_editor_service_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PrintEditorCodeInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PrintEditorCodeInput) ProtoMessage() {}

func (x *PrintEditorCodeInput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_atlas_editor_service_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PrintEditorCodeInput.ProtoReflect.Descriptor instead.
func (*PrintEditorCodeInput) Descriptor() ([]byte, []int) {
	return file_eolymp_atlas_editor_service_proto_rawDescGZIP(), []int{7}
}

func (x *PrintEditorCodeInput) GetRuntime() string {
	if x != nil {
		return x.Runtime
	}
	return ""
}

func (x *PrintEditorCodeInput) GetSourceCode() string {
	if x != nil {
		return x.SourceCode
	}
	return ""
}

type PrintEditorCodeOutput struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PrintEditorCodeOutput) Reset() {
	*x = PrintEditorCodeOutput{}
	mi := &file_eolymp_atlas_editor_service_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PrintEditorCodeOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PrintEditorCodeOutput) ProtoMessage() {}

func (x *PrintEditorCodeOutput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_atlas_editor_service_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PrintEditorCodeOutput.ProtoReflect.Descriptor instead.
func (*PrintEditorCodeOutput) Descriptor() ([]byte, []int) {
	return file_eolymp_atlas_editor_service_proto_rawDescGZIP(), []int{8}
}

type Editor_State struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Runtime       string                 `protobuf:"bytes,1,opt,name=runtime,proto3" json:"runtime,omitempty"`
	SourceCode    string                 `protobuf:"bytes,2,opt,name=source_code,json=sourceCode,proto3" json:"source_code,omitempty"`
	InputData     string                 `protobuf:"bytes,3,opt,name=input_data,json=inputData,proto3" json:"input_data,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Editor_State) Reset() {
	*x = Editor_State{}
	mi := &file_eolymp_atlas_editor_service_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Editor_State) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Editor_State) ProtoMessage() {}

func (x *Editor_State) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_atlas_editor_service_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Editor_State.ProtoReflect.Descriptor instead.
func (*Editor_State) Descriptor() ([]byte, []int) {
	return file_eolymp_atlas_editor_service_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Editor_State) GetRuntime() string {
	if x != nil {
		return x.Runtime
	}
	return ""
}

func (x *Editor_State) GetSourceCode() string {
	if x != nil {
		return x.SourceCode
	}
	return ""
}

func (x *Editor_State) GetInputData() string {
	if x != nil {
		return x.InputData
	}
	return ""
}

var File_eolymp_atlas_editor_service_proto protoreflect.FileDescriptor

var file_eolymp_atlas_editor_service_proto_rawDesc = string([]byte{
	0x0a, 0x21, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x61, 0x74, 0x6c, 0x61, 0x73, 0x2f, 0x65,
	0x64, 0x69, 0x74, 0x6f, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x61, 0x74, 0x6c, 0x61,
	0x73, 0x1a, 0x1d, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x22, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x72, 0x61, 0x74, 0x65, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x72, 0x75, 0x6e,
	0x74, 0x69, 0x6d, 0x65, 0x2f, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xbc, 0x02, 0x0a, 0x06, 0x45, 0x64, 0x69, 0x74, 0x6f, 0x72, 0x12, 0x30, 0x0a,
	0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x65,
	0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x61, 0x74, 0x6c, 0x61, 0x73, 0x2e, 0x45, 0x64, 0x69, 0x74,
	0x6f, 0x72, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12,
	0x38, 0x0a, 0x08, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28,
	0x0e, 0x32, 0x1c, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x61, 0x74, 0x6c, 0x61, 0x73,
	0x2e, 0x45, 0x64, 0x69, 0x74, 0x6f, 0x72, 0x2e, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x52,
	0x08, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x12, 0x33, 0x0a, 0x08, 0x72, 0x75, 0x6e,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x18, 0x0b, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x65, 0x6f,
	0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x52, 0x75, 0x6e,
	0x74, 0x69, 0x6d, 0x65, 0x52, 0x08, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x1a, 0x61,
	0x0a, 0x05, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x75, 0x6e, 0x74, 0x69,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d,
	0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x63, 0x6f, 0x64, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x43, 0x6f,
	0x64, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x61,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x44, 0x61, 0x74,
	0x61, 0x22, 0x2e, 0x0a, 0x07, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x13, 0x0a, 0x0f,
	0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x5f, 0x46, 0x45, 0x41, 0x54, 0x55, 0x52, 0x45, 0x10,
	0x00, 0x12, 0x0e, 0x0a, 0x0a, 0x50, 0x52, 0x49, 0x4e, 0x54, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x10,
	0x01, 0x22, 0x15, 0x0a, 0x13, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x45, 0x64, 0x69,
	0x74, 0x6f, 0x72, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x22, 0x44, 0x0a, 0x14, 0x44, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x62, 0x65, 0x45, 0x64, 0x69, 0x74, 0x6f, 0x72, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74,
	0x12, 0x2c, 0x0a, 0x06, 0x65, 0x64, 0x69, 0x74, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x14, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x61, 0x74, 0x6c, 0x61, 0x73, 0x2e,
	0x45, 0x64, 0x69, 0x74, 0x6f, 0x72, 0x52, 0x06, 0x65, 0x64, 0x69, 0x74, 0x6f, 0x72, 0x22, 0x1a,
	0x0a, 0x18, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x45, 0x64, 0x69, 0x74, 0x6f, 0x72,
	0x53, 0x74, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x22, 0xaf, 0x01, 0x0a, 0x19, 0x44,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x45, 0x64, 0x69, 0x74, 0x6f, 0x72, 0x53, 0x74, 0x61,
	0x74, 0x65, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x75, 0x6e, 0x74,
	0x69, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x72, 0x75, 0x6e, 0x74, 0x69,
	0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x63, 0x6f, 0x64,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x43,
	0x6f, 0x64, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x5f, 0x64, 0x61, 0x74,
	0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x44, 0x61,
	0x74, 0x61, 0x12, 0x38, 0x0a, 0x08, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x18, 0x0a,
	0x20, 0x03, 0x28, 0x0e, 0x32, 0x1c, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x61, 0x74,
	0x6c, 0x61, 0x73, 0x2e, 0x45, 0x64, 0x69, 0x74, 0x6f, 0x72, 0x2e, 0x46, 0x65, 0x61, 0x74, 0x75,
	0x72, 0x65, 0x52, 0x08, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x22, 0x72, 0x0a, 0x16,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x45, 0x64, 0x69, 0x74, 0x6f, 0x72, 0x53, 0x74, 0x61, 0x74,
	0x65, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65,
	0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x43, 0x6f, 0x64,
	0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x44, 0x61, 0x74, 0x61,
	0x22, 0x19, 0x0a, 0x17, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x45, 0x64, 0x69, 0x74, 0x6f, 0x72,
	0x53, 0x74, 0x61, 0x74, 0x65, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x22, 0x51, 0x0a, 0x14, 0x50,
	0x72, 0x69, 0x6e, 0x74, 0x45, 0x64, 0x69, 0x74, 0x6f, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x49, 0x6e,
	0x70, 0x75, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x1f, 0x0a,
	0x0b, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x22, 0x17,
	0x0a, 0x15, 0x50, 0x72, 0x69, 0x6e, 0x74, 0x45, 0x64, 0x69, 0x74, 0x6f, 0x72, 0x43, 0x6f, 0x64,
	0x65, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x32, 0x9a, 0x05, 0x0a, 0x0d, 0x45, 0x64, 0x69, 0x74,
	0x6f, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x94, 0x01, 0x0a, 0x0e, 0x44, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x45, 0x64, 0x69, 0x74, 0x6f, 0x72, 0x12, 0x21, 0x2e, 0x65,
	0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x61, 0x74, 0x6c, 0x61, 0x73, 0x2e, 0x44, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x62, 0x65, 0x45, 0x64, 0x69, 0x74, 0x6f, 0x72, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a,
	0x22, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x61, 0x74, 0x6c, 0x61, 0x73, 0x2e, 0x44,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x45, 0x64, 0x69, 0x74, 0x6f, 0x72, 0x4f, 0x75, 0x74,
	0x70, 0x75, 0x74, 0x22, 0x3b, 0xea, 0xe2, 0x0a, 0x0b, 0xf5, 0xe2, 0x0a, 0x00, 0x00, 0xa0, 0x40,
	0xf8, 0xe2, 0x0a, 0x0a, 0x82, 0xe3, 0x0a, 0x19, 0x8a, 0xe3, 0x0a, 0x15, 0x61, 0x74, 0x6c, 0x61,
	0x73, 0x3a, 0x73, 0x75, 0x62, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x3a, 0x72, 0x65, 0x61,
	0x64, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x09, 0x12, 0x07, 0x2f, 0x65, 0x64, 0x69, 0x74, 0x6f, 0x72,
	0x12, 0xa9, 0x01, 0x0a, 0x13, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x45, 0x64, 0x69,
	0x74, 0x6f, 0x72, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x26, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d,
	0x70, 0x2e, 0x61, 0x74, 0x6c, 0x61, 0x73, 0x2e, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65,
	0x45, 0x64, 0x69, 0x74, 0x6f, 0x72, 0x53, 0x74, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x70, 0x75, 0x74,
	0x1a, 0x27, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x61, 0x74, 0x6c, 0x61, 0x73, 0x2e,
	0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x45, 0x64, 0x69, 0x74, 0x6f, 0x72, 0x53, 0x74,
	0x61, 0x74, 0x65, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x22, 0x41, 0xea, 0xe2, 0x0a, 0x0b, 0xf5,
	0xe2, 0x0a, 0x00, 0x00, 0xa0, 0x40, 0xf8, 0xe2, 0x0a, 0x0a, 0x82, 0xe3, 0x0a, 0x19, 0x8a, 0xe3,
	0x0a, 0x15, 0x61, 0x74, 0x6c, 0x61, 0x73, 0x3a, 0x73, 0x75, 0x62, 0x6d, 0x69, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x3a, 0x72, 0x65, 0x61, 0x64, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0f, 0x12, 0x0d, 0x2f,
	0x65, 0x64, 0x69, 0x74, 0x6f, 0x72, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0xa4, 0x01, 0x0a,
	0x11, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x45, 0x64, 0x69, 0x74, 0x6f, 0x72, 0x53, 0x74, 0x61,
	0x74, 0x65, 0x12, 0x24, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x61, 0x74, 0x6c, 0x61,
	0x73, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x45, 0x64, 0x69, 0x74, 0x6f, 0x72, 0x53, 0x74,
	0x61, 0x74, 0x65, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x25, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d,
	0x70, 0x2e, 0x61, 0x74, 0x6c, 0x61, 0x73, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x45, 0x64,
	0x69, 0x74, 0x6f, 0x72, 0x53, 0x74, 0x61, 0x74, 0x65, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x22,
	0x42, 0xea, 0xe2, 0x0a, 0x0b, 0xf5, 0xe2, 0x0a, 0x00, 0x00, 0x80, 0x3f, 0xf8, 0xe2, 0x0a, 0x02,
	0x82, 0xe3, 0x0a, 0x1a, 0x8a, 0xe3, 0x0a, 0x16, 0x61, 0x74, 0x6c, 0x61, 0x73, 0x3a, 0x73, 0x75,
	0x62, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x3a, 0x77, 0x72, 0x69, 0x74, 0x65, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x0f, 0x22, 0x0d, 0x2f, 0x65, 0x64, 0x69, 0x74, 0x6f, 0x72, 0x2f, 0x73, 0x74,
	0x61, 0x74, 0x65, 0x12, 0x9e, 0x01, 0x0a, 0x0f, 0x50, 0x72, 0x69, 0x6e, 0x74, 0x45, 0x64, 0x69,
	0x74, 0x6f, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x22, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70,
	0x2e, 0x61, 0x74, 0x6c, 0x61, 0x73, 0x2e, 0x50, 0x72, 0x69, 0x6e, 0x74, 0x45, 0x64, 0x69, 0x74,
	0x6f, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x23, 0x2e, 0x65, 0x6f,
	0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x61, 0x74, 0x6c, 0x61, 0x73, 0x2e, 0x50, 0x72, 0x69, 0x6e, 0x74,
	0x45, 0x64, 0x69, 0x74, 0x6f, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74,
	0x22, 0x42, 0xea, 0xe2, 0x0a, 0x0b, 0xf5, 0xe2, 0x0a, 0x00, 0x00, 0x80, 0x3f, 0xf8, 0xe2, 0x0a,
	0x02, 0x82, 0xe3, 0x0a, 0x1a, 0x8a, 0xe3, 0x0a, 0x16, 0x61, 0x74, 0x6c, 0x61, 0x73, 0x3a, 0x73,
	0x75, 0x62, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x3a, 0x77, 0x72, 0x69, 0x74, 0x65, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x0f, 0x22, 0x0d, 0x2f, 0x65, 0x64, 0x69, 0x74, 0x6f, 0x72, 0x2f, 0x70,
	0x72, 0x69, 0x6e, 0x74, 0x42, 0x2d, 0x5a, 0x2b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x67, 0x6f, 0x2d, 0x73, 0x64, 0x6b,
	0x2f, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x61, 0x74, 0x6c, 0x61, 0x73, 0x3b, 0x61, 0x74,
	0x6c, 0x61, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_eolymp_atlas_editor_service_proto_rawDescOnce sync.Once
	file_eolymp_atlas_editor_service_proto_rawDescData []byte
)

func file_eolymp_atlas_editor_service_proto_rawDescGZIP() []byte {
	file_eolymp_atlas_editor_service_proto_rawDescOnce.Do(func() {
		file_eolymp_atlas_editor_service_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_eolymp_atlas_editor_service_proto_rawDesc), len(file_eolymp_atlas_editor_service_proto_rawDesc)))
	})
	return file_eolymp_atlas_editor_service_proto_rawDescData
}

var file_eolymp_atlas_editor_service_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_eolymp_atlas_editor_service_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_eolymp_atlas_editor_service_proto_goTypes = []any{
	(Editor_Feature)(0),               // 0: eolymp.atlas.Editor.Feature
	(*Editor)(nil),                    // 1: eolymp.atlas.Editor
	(*DescribeEditorInput)(nil),       // 2: eolymp.atlas.DescribeEditorInput
	(*DescribeEditorOutput)(nil),      // 3: eolymp.atlas.DescribeEditorOutput
	(*DescribeEditorStateInput)(nil),  // 4: eolymp.atlas.DescribeEditorStateInput
	(*DescribeEditorStateOutput)(nil), // 5: eolymp.atlas.DescribeEditorStateOutput
	(*UpdateEditorStateInput)(nil),    // 6: eolymp.atlas.UpdateEditorStateInput
	(*UpdateEditorStateOutput)(nil),   // 7: eolymp.atlas.UpdateEditorStateOutput
	(*PrintEditorCodeInput)(nil),      // 8: eolymp.atlas.PrintEditorCodeInput
	(*PrintEditorCodeOutput)(nil),     // 9: eolymp.atlas.PrintEditorCodeOutput
	(*Editor_State)(nil),              // 10: eolymp.atlas.Editor.State
	(*runtime.Runtime)(nil),           // 11: eolymp.runtime.Runtime
}
var file_eolymp_atlas_editor_service_proto_depIdxs = []int32{
	10, // 0: eolymp.atlas.Editor.state:type_name -> eolymp.atlas.Editor.State
	0,  // 1: eolymp.atlas.Editor.features:type_name -> eolymp.atlas.Editor.Feature
	11, // 2: eolymp.atlas.Editor.runtimes:type_name -> eolymp.runtime.Runtime
	1,  // 3: eolymp.atlas.DescribeEditorOutput.editor:type_name -> eolymp.atlas.Editor
	0,  // 4: eolymp.atlas.DescribeEditorStateOutput.features:type_name -> eolymp.atlas.Editor.Feature
	2,  // 5: eolymp.atlas.EditorService.DescribeEditor:input_type -> eolymp.atlas.DescribeEditorInput
	4,  // 6: eolymp.atlas.EditorService.DescribeEditorState:input_type -> eolymp.atlas.DescribeEditorStateInput
	6,  // 7: eolymp.atlas.EditorService.UpdateEditorState:input_type -> eolymp.atlas.UpdateEditorStateInput
	8,  // 8: eolymp.atlas.EditorService.PrintEditorCode:input_type -> eolymp.atlas.PrintEditorCodeInput
	3,  // 9: eolymp.atlas.EditorService.DescribeEditor:output_type -> eolymp.atlas.DescribeEditorOutput
	5,  // 10: eolymp.atlas.EditorService.DescribeEditorState:output_type -> eolymp.atlas.DescribeEditorStateOutput
	7,  // 11: eolymp.atlas.EditorService.UpdateEditorState:output_type -> eolymp.atlas.UpdateEditorStateOutput
	9,  // 12: eolymp.atlas.EditorService.PrintEditorCode:output_type -> eolymp.atlas.PrintEditorCodeOutput
	9,  // [9:13] is the sub-list for method output_type
	5,  // [5:9] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_eolymp_atlas_editor_service_proto_init() }
func file_eolymp_atlas_editor_service_proto_init() {
	if File_eolymp_atlas_editor_service_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_eolymp_atlas_editor_service_proto_rawDesc), len(file_eolymp_atlas_editor_service_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_eolymp_atlas_editor_service_proto_goTypes,
		DependencyIndexes: file_eolymp_atlas_editor_service_proto_depIdxs,
		EnumInfos:         file_eolymp_atlas_editor_service_proto_enumTypes,
		MessageInfos:      file_eolymp_atlas_editor_service_proto_msgTypes,
	}.Build()
	File_eolymp_atlas_editor_service_proto = out.File
	file_eolymp_atlas_editor_service_proto_goTypes = nil
	file_eolymp_atlas_editor_service_proto_depIdxs = nil
}
