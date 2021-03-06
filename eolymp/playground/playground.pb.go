// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.18.1
// source: eolymp/playground/playground.proto

package playground

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

type CreateRunInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// deprecated, use runtime instead
	Lang    string `protobuf:"bytes,1,opt,name=lang,proto3" json:"lang,omitempty"`
	Runtime string `protobuf:"bytes,2,opt,name=runtime,proto3" json:"runtime,omitempty"`
	// Source code
	// deprecated, use source_ern instead
	Source string `protobuf:"bytes,31,opt,name=source,proto3" json:"source,omitempty"`
	// Source code ERN (data up to 5KB or blob)
	SourceErn string `protobuf:"bytes,30,opt,name=source_ern,json=sourceErn,proto3" json:"source_ern,omitempty"`
	// Input ERN (data up to 5KB or blob)
	InputErn string `protobuf:"bytes,10,opt,name=input_ern,json=inputErn,proto3" json:"input_ern,omitempty"`
	// Problem ERN allows to reuse template from Atlas or Judge problem.
	ProblemErn string `protobuf:"bytes,20,opt,name=problem_ern,json=problemErn,proto3" json:"problem_ern,omitempty"`
	// deprecated, use input ern instead
	Input string `protobuf:"bytes,3,opt,name=input,proto3" json:"input,omitempty"`
	// deprecated use input ern instead
	//
	// Types that are assignable to InputData:
	//	*CreateRunInput_InputContent
	//	*CreateRunInput_InputObjectId
	InputData isCreateRunInput_InputData `protobuf_oneof:"input_data"`
}

func (x *CreateRunInput) Reset() {
	*x = CreateRunInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_playground_playground_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateRunInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRunInput) ProtoMessage() {}

func (x *CreateRunInput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_playground_playground_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRunInput.ProtoReflect.Descriptor instead.
func (*CreateRunInput) Descriptor() ([]byte, []int) {
	return file_eolymp_playground_playground_proto_rawDescGZIP(), []int{0}
}

func (x *CreateRunInput) GetLang() string {
	if x != nil {
		return x.Lang
	}
	return ""
}

func (x *CreateRunInput) GetRuntime() string {
	if x != nil {
		return x.Runtime
	}
	return ""
}

func (x *CreateRunInput) GetSource() string {
	if x != nil {
		return x.Source
	}
	return ""
}

func (x *CreateRunInput) GetSourceErn() string {
	if x != nil {
		return x.SourceErn
	}
	return ""
}

func (x *CreateRunInput) GetInputErn() string {
	if x != nil {
		return x.InputErn
	}
	return ""
}

func (x *CreateRunInput) GetProblemErn() string {
	if x != nil {
		return x.ProblemErn
	}
	return ""
}

func (x *CreateRunInput) GetInput() string {
	if x != nil {
		return x.Input
	}
	return ""
}

func (m *CreateRunInput) GetInputData() isCreateRunInput_InputData {
	if m != nil {
		return m.InputData
	}
	return nil
}

func (x *CreateRunInput) GetInputContent() []byte {
	if x, ok := x.GetInputData().(*CreateRunInput_InputContent); ok {
		return x.InputContent
	}
	return nil
}

func (x *CreateRunInput) GetInputObjectId() string {
	if x, ok := x.GetInputData().(*CreateRunInput_InputObjectId); ok {
		return x.InputObjectId
	}
	return ""
}

type isCreateRunInput_InputData interface {
	isCreateRunInput_InputData()
}

type CreateRunInput_InputContent struct {
	InputContent []byte `protobuf:"bytes,4,opt,name=input_content,json=inputContent,proto3,oneof"`
}

type CreateRunInput_InputObjectId struct {
	InputObjectId string `protobuf:"bytes,5,opt,name=input_object_id,json=inputObjectId,proto3,oneof"`
}

func (*CreateRunInput_InputContent) isCreateRunInput_InputData() {}

func (*CreateRunInput_InputObjectId) isCreateRunInput_InputData() {}

type CreateRunOutput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RunId string `protobuf:"bytes,1,opt,name=run_id,json=runId,proto3" json:"run_id,omitempty"`
}

func (x *CreateRunOutput) Reset() {
	*x = CreateRunOutput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_playground_playground_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateRunOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRunOutput) ProtoMessage() {}

func (x *CreateRunOutput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_playground_playground_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRunOutput.ProtoReflect.Descriptor instead.
func (*CreateRunOutput) Descriptor() ([]byte, []int) {
	return file_eolymp_playground_playground_proto_rawDescGZIP(), []int{1}
}

func (x *CreateRunOutput) GetRunId() string {
	if x != nil {
		return x.RunId
	}
	return ""
}

type DescribeRunInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RunId string `protobuf:"bytes,1,opt,name=run_id,json=runId,proto3" json:"run_id,omitempty"`
}

func (x *DescribeRunInput) Reset() {
	*x = DescribeRunInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_playground_playground_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DescribeRunInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescribeRunInput) ProtoMessage() {}

func (x *DescribeRunInput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_playground_playground_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DescribeRunInput.ProtoReflect.Descriptor instead.
func (*DescribeRunInput) Descriptor() ([]byte, []int) {
	return file_eolymp_playground_playground_proto_rawDescGZIP(), []int{2}
}

func (x *DescribeRunInput) GetRunId() string {
	if x != nil {
		return x.RunId
	}
	return ""
}

type DescribeRunOutput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Run *Run `protobuf:"bytes,1,opt,name=run,proto3" json:"run,omitempty"`
}

func (x *DescribeRunOutput) Reset() {
	*x = DescribeRunOutput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_playground_playground_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DescribeRunOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescribeRunOutput) ProtoMessage() {}

func (x *DescribeRunOutput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_playground_playground_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DescribeRunOutput.ProtoReflect.Descriptor instead.
func (*DescribeRunOutput) Descriptor() ([]byte, []int) {
	return file_eolymp_playground_playground_proto_rawDescGZIP(), []int{3}
}

func (x *DescribeRunOutput) GetRun() *Run {
	if x != nil {
		return x.Run
	}
	return nil
}

var File_eolymp_playground_playground_proto protoreflect.FileDescriptor

var file_eolymp_playground_playground_proto_rawDesc = []byte{
	0x0a, 0x22, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x70, 0x6c, 0x61, 0x79, 0x67, 0x72, 0x6f,
	0x75, 0x6e, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x79, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x70, 0x6c, 0x61,
	0x79, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x1a, 0x1d, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x68, 0x74, 0x74, 0x70,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x22, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x61,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x72, 0x61, 0x74, 0x65, 0x6c,
	0x69, 0x6d, 0x69, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x65, 0x6f, 0x6c, 0x79,
	0x6d, 0x70, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73,
	0x63, 0x6f, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x65, 0x6f, 0x6c, 0x79,
	0x6d, 0x70, 0x2f, 0x70, 0x6c, 0x61, 0x79, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x2f, 0x72, 0x75,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa8, 0x02, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x52, 0x75, 0x6e, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x61,
	0x6e, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6c, 0x61, 0x6e, 0x67, 0x12, 0x18,
	0x0a, 0x07, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x18, 0x1f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x65, 0x72, 0x6e, 0x18, 0x1e,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x45, 0x72, 0x6e, 0x12,
	0x1b, 0x0a, 0x09, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x5f, 0x65, 0x72, 0x6e, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x45, 0x72, 0x6e, 0x12, 0x1f, 0x0a, 0x0b,
	0x70, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x5f, 0x65, 0x72, 0x6e, 0x18, 0x14, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x70, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x45, 0x72, 0x6e, 0x12, 0x14, 0x0a,
	0x05, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x69, 0x6e,
	0x70, 0x75, 0x74, 0x12, 0x25, 0x0a, 0x0d, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x5f, 0x63, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x48, 0x00, 0x52, 0x0c, 0x69, 0x6e,
	0x70, 0x75, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x28, 0x0a, 0x0f, 0x69, 0x6e,
	0x70, 0x75, 0x74, 0x5f, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0d, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x4f, 0x62, 0x6a, 0x65,
	0x63, 0x74, 0x49, 0x64, 0x42, 0x0c, 0x0a, 0x0a, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x5f, 0x64, 0x61,
	0x74, 0x61, 0x22, 0x28, 0x0a, 0x0f, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x75, 0x6e, 0x4f,
	0x75, 0x74, 0x70, 0x75, 0x74, 0x12, 0x15, 0x0a, 0x06, 0x72, 0x75, 0x6e, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x72, 0x75, 0x6e, 0x49, 0x64, 0x22, 0x29, 0x0a, 0x10,
	0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x52, 0x75, 0x6e, 0x49, 0x6e, 0x70, 0x75, 0x74,
	0x12, 0x15, 0x0a, 0x06, 0x72, 0x75, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x72, 0x75, 0x6e, 0x49, 0x64, 0x22, 0x3d, 0x0a, 0x11, 0x44, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x62, 0x65, 0x52, 0x75, 0x6e, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x12, 0x28, 0x0a, 0x03,
	0x72, 0x75, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x65, 0x6f, 0x6c, 0x79,
	0x6d, 0x70, 0x2e, 0x70, 0x6c, 0x61, 0x79, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x2e, 0x52, 0x75,
	0x6e, 0x52, 0x03, 0x72, 0x75, 0x6e, 0x32, 0xd3, 0x02, 0x0a, 0x0a, 0x50, 0x6c, 0x61, 0x79, 0x67,
	0x72, 0x6f, 0x75, 0x6e, 0x64, 0x12, 0x97, 0x01, 0x0a, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x52, 0x75, 0x6e, 0x12, 0x21, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x70, 0x6c, 0x61,
	0x79, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x75,
	0x6e, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x22, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e,
	0x70, 0x6c, 0x61, 0x79, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x52, 0x75, 0x6e, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x22, 0x43, 0x82, 0xe3, 0x0a, 0x18,
	0x8a, 0xe3, 0x0a, 0x14, 0x70, 0x6c, 0x61, 0x79, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x3a, 0x72,
	0x75, 0x6e, 0x3a, 0x77, 0x72, 0x69, 0x74, 0x65, 0xea, 0xe2, 0x0a, 0x0b, 0xf5, 0xe2, 0x0a, 0x0a,
	0xd7, 0x23, 0x3e, 0xf8, 0xe2, 0x0a, 0x05, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x12, 0x22, 0x10, 0x2f,
	0x70, 0x6c, 0x61, 0x79, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x2f, 0x72, 0x75, 0x6e, 0x73, 0x12,
	0xaa, 0x01, 0x0a, 0x0b, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x52, 0x75, 0x6e, 0x12,
	0x23, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x70, 0x6c, 0x61, 0x79, 0x67, 0x72, 0x6f,
	0x75, 0x6e, 0x64, 0x2e, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x52, 0x75, 0x6e, 0x49,
	0x6e, 0x70, 0x75, 0x74, 0x1a, 0x24, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x70, 0x6c,
	0x61, 0x79, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x2e, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62,
	0x65, 0x52, 0x75, 0x6e, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x22, 0x50, 0x82, 0xe3, 0x0a, 0x17,
	0x8a, 0xe3, 0x0a, 0x13, 0x70, 0x6c, 0x61, 0x79, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x3a, 0x72,
	0x75, 0x6e, 0x3a, 0x72, 0x65, 0x61, 0x64, 0xea, 0xe2, 0x0a, 0x0b, 0xf5, 0xe2, 0x0a, 0x00, 0x00,
	0x00, 0x40, 0xf8, 0xe2, 0x0a, 0x0a, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x20, 0x12, 0x19, 0x2f, 0x70,
	0x6c, 0x61, 0x79, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x2f, 0x72, 0x75, 0x6e, 0x73, 0x2f, 0x7b,
	0x72, 0x75, 0x6e, 0x5f, 0x69, 0x64, 0x7d, 0x62, 0x03, 0x72, 0x75, 0x6e, 0x42, 0x37, 0x5a, 0x35,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6f, 0x6c, 0x79, 0x6d,
	0x70, 0x2f, 0x67, 0x6f, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f,
	0x70, 0x6c, 0x61, 0x79, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x3b, 0x70, 0x6c, 0x61, 0x79, 0x67,
	0x72, 0x6f, 0x75, 0x6e, 0x64, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_eolymp_playground_playground_proto_rawDescOnce sync.Once
	file_eolymp_playground_playground_proto_rawDescData = file_eolymp_playground_playground_proto_rawDesc
)

func file_eolymp_playground_playground_proto_rawDescGZIP() []byte {
	file_eolymp_playground_playground_proto_rawDescOnce.Do(func() {
		file_eolymp_playground_playground_proto_rawDescData = protoimpl.X.CompressGZIP(file_eolymp_playground_playground_proto_rawDescData)
	})
	return file_eolymp_playground_playground_proto_rawDescData
}

var file_eolymp_playground_playground_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_eolymp_playground_playground_proto_goTypes = []interface{}{
	(*CreateRunInput)(nil),    // 0: eolymp.playground.CreateRunInput
	(*CreateRunOutput)(nil),   // 1: eolymp.playground.CreateRunOutput
	(*DescribeRunInput)(nil),  // 2: eolymp.playground.DescribeRunInput
	(*DescribeRunOutput)(nil), // 3: eolymp.playground.DescribeRunOutput
	(*Run)(nil),               // 4: eolymp.playground.Run
}
var file_eolymp_playground_playground_proto_depIdxs = []int32{
	4, // 0: eolymp.playground.DescribeRunOutput.run:type_name -> eolymp.playground.Run
	0, // 1: eolymp.playground.Playground.CreateRun:input_type -> eolymp.playground.CreateRunInput
	2, // 2: eolymp.playground.Playground.DescribeRun:input_type -> eolymp.playground.DescribeRunInput
	1, // 3: eolymp.playground.Playground.CreateRun:output_type -> eolymp.playground.CreateRunOutput
	3, // 4: eolymp.playground.Playground.DescribeRun:output_type -> eolymp.playground.DescribeRunOutput
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_eolymp_playground_playground_proto_init() }
func file_eolymp_playground_playground_proto_init() {
	if File_eolymp_playground_playground_proto != nil {
		return
	}
	file_eolymp_playground_run_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_eolymp_playground_playground_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateRunInput); i {
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
		file_eolymp_playground_playground_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateRunOutput); i {
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
		file_eolymp_playground_playground_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DescribeRunInput); i {
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
		file_eolymp_playground_playground_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DescribeRunOutput); i {
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
	file_eolymp_playground_playground_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*CreateRunInput_InputContent)(nil),
		(*CreateRunInput_InputObjectId)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_eolymp_playground_playground_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_eolymp_playground_playground_proto_goTypes,
		DependencyIndexes: file_eolymp_playground_playground_proto_depIdxs,
		MessageInfos:      file_eolymp_playground_playground_proto_msgTypes,
	}.Build()
	File_eolymp_playground_playground_proto = out.File
	file_eolymp_playground_playground_proto_rawDesc = nil
	file_eolymp_playground_playground_proto_goTypes = nil
	file_eolymp_playground_playground_proto_depIdxs = nil
}
