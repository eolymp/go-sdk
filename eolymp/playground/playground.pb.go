// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.0
// 	protoc        v3.21.12
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

	Runtime string `protobuf:"bytes,2,opt,name=runtime,proto3" json:"runtime,omitempty"`
	Source  string `protobuf:"bytes,100,opt,name=source,proto3" json:"source,omitempty"`
	// Types that are assignable to Input:
	//
	//	*CreateRunInput_InputData
	//	*CreateRunInput_InputRef
	Input isCreateRunInput_Input `protobuf_oneof:"input"`
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

func (m *CreateRunInput) GetInput() isCreateRunInput_Input {
	if m != nil {
		return m.Input
	}
	return nil
}

func (x *CreateRunInput) GetInputData() []byte {
	if x, ok := x.GetInput().(*CreateRunInput_InputData); ok {
		return x.InputData
	}
	return nil
}

func (x *CreateRunInput) GetInputRef() string {
	if x, ok := x.GetInput().(*CreateRunInput_InputRef); ok {
		return x.InputRef
	}
	return ""
}

type isCreateRunInput_Input interface {
	isCreateRunInput_Input()
}

type CreateRunInput_InputData struct {
	InputData []byte `protobuf:"bytes,101,opt,name=input_data,json=inputData,proto3,oneof"`
}

type CreateRunInput_InputRef struct {
	InputRef string `protobuf:"bytes,102,opt,name=input_ref,json=inputRef,proto3,oneof"`
}

func (*CreateRunInput_InputData) isCreateRunInput_Input() {}

func (*CreateRunInput_InputRef) isCreateRunInput_Input() {}

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

type WatchRunInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RunId string `protobuf:"bytes,1,opt,name=run_id,json=runId,proto3" json:"run_id,omitempty"`
}

func (x *WatchRunInput) Reset() {
	*x = WatchRunInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_playground_playground_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WatchRunInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WatchRunInput) ProtoMessage() {}

func (x *WatchRunInput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_playground_playground_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WatchRunInput.ProtoReflect.Descriptor instead.
func (*WatchRunInput) Descriptor() ([]byte, []int) {
	return file_eolymp_playground_playground_proto_rawDescGZIP(), []int{4}
}

func (x *WatchRunInput) GetRunId() string {
	if x != nil {
		return x.RunId
	}
	return ""
}

type WatchRunOutput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Run *Run `protobuf:"bytes,1,opt,name=run,proto3" json:"run,omitempty"`
}

func (x *WatchRunOutput) Reset() {
	*x = WatchRunOutput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_playground_playground_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WatchRunOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WatchRunOutput) ProtoMessage() {}

func (x *WatchRunOutput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_playground_playground_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WatchRunOutput.ProtoReflect.Descriptor instead.
func (*WatchRunOutput) Descriptor() ([]byte, []int) {
	return file_eolymp_playground_playground_proto_rawDescGZIP(), []int{5}
}

func (x *WatchRunOutput) GetRun() *Run {
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
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8b, 0x01, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x52, 0x75, 0x6e, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x75,
	0x6e, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x72, 0x75, 0x6e,
	0x74, 0x69, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x64,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x1f, 0x0a, 0x0a,
	0x69, 0x6e, 0x70, 0x75, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x65, 0x20, 0x01, 0x28, 0x0c,
	0x48, 0x00, 0x52, 0x09, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x44, 0x61, 0x74, 0x61, 0x12, 0x1d, 0x0a,
	0x09, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x5f, 0x72, 0x65, 0x66, 0x18, 0x66, 0x20, 0x01, 0x28, 0x09,
	0x48, 0x00, 0x52, 0x08, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x52, 0x65, 0x66, 0x42, 0x07, 0x0a, 0x05,
	0x69, 0x6e, 0x70, 0x75, 0x74, 0x22, 0x28, 0x0a, 0x0f, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52,
	0x75, 0x6e, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x12, 0x15, 0x0a, 0x06, 0x72, 0x75, 0x6e, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x72, 0x75, 0x6e, 0x49, 0x64, 0x22,
	0x29, 0x0a, 0x10, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x52, 0x75, 0x6e, 0x49, 0x6e,
	0x70, 0x75, 0x74, 0x12, 0x15, 0x0a, 0x06, 0x72, 0x75, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x72, 0x75, 0x6e, 0x49, 0x64, 0x22, 0x3d, 0x0a, 0x11, 0x44, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x52, 0x75, 0x6e, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x12,
	0x28, 0x0a, 0x03, 0x72, 0x75, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x65,
	0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x70, 0x6c, 0x61, 0x79, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64,
	0x2e, 0x52, 0x75, 0x6e, 0x52, 0x03, 0x72, 0x75, 0x6e, 0x22, 0x26, 0x0a, 0x0d, 0x57, 0x61, 0x74,
	0x63, 0x68, 0x52, 0x75, 0x6e, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x15, 0x0a, 0x06, 0x72, 0x75,
	0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x72, 0x75, 0x6e, 0x49,
	0x64, 0x22, 0x3a, 0x0a, 0x0e, 0x57, 0x61, 0x74, 0x63, 0x68, 0x52, 0x75, 0x6e, 0x4f, 0x75, 0x74,
	0x70, 0x75, 0x74, 0x12, 0x28, 0x0a, 0x03, 0x72, 0x75, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x16, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x70, 0x6c, 0x61, 0x79, 0x67, 0x72,
	0x6f, 0x75, 0x6e, 0x64, 0x2e, 0x52, 0x75, 0x6e, 0x52, 0x03, 0x72, 0x75, 0x6e, 0x32, 0xd4, 0x03,
	0x0a, 0x0a, 0x50, 0x6c, 0x61, 0x79, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x12, 0x8c, 0x01, 0x0a,
	0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x75, 0x6e, 0x12, 0x21, 0x2e, 0x65, 0x6f, 0x6c,
	0x79, 0x6d, 0x70, 0x2e, 0x70, 0x6c, 0x61, 0x79, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x75, 0x6e, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x22, 0x2e,
	0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x70, 0x6c, 0x61, 0x79, 0x67, 0x72, 0x6f, 0x75, 0x6e,
	0x64, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x75, 0x6e, 0x4f, 0x75, 0x74, 0x70, 0x75,
	0x74, 0x22, 0x38, 0xea, 0xe2, 0x0a, 0x0b, 0xf5, 0xe2, 0x0a, 0x0a, 0xd7, 0x23, 0x3e, 0xf8, 0xe2,
	0x0a, 0x05, 0x82, 0xe3, 0x0a, 0x18, 0x8a, 0xe3, 0x0a, 0x14, 0x70, 0x6c, 0x61, 0x79, 0x67, 0x72,
	0x6f, 0x75, 0x6e, 0x64, 0x3a, 0x72, 0x75, 0x6e, 0x3a, 0x77, 0x72, 0x69, 0x74, 0x65, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x07, 0x22, 0x05, 0x2f, 0x72, 0x75, 0x6e, 0x73, 0x12, 0x9a, 0x01, 0x0a, 0x0b,
	0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x52, 0x75, 0x6e, 0x12, 0x23, 0x2e, 0x65, 0x6f,
	0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x70, 0x6c, 0x61, 0x79, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x2e,
	0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x52, 0x75, 0x6e, 0x49, 0x6e, 0x70, 0x75, 0x74,
	0x1a, 0x24, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x70, 0x6c, 0x61, 0x79, 0x67, 0x72,
	0x6f, 0x75, 0x6e, 0x64, 0x2e, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x52, 0x75, 0x6e,
	0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x22, 0x40, 0xea, 0xe2, 0x0a, 0x0b, 0xf5, 0xe2, 0x0a, 0x00,
	0x00, 0x00, 0x40, 0xf8, 0xe2, 0x0a, 0x0a, 0x82, 0xe3, 0x0a, 0x17, 0x8a, 0xe3, 0x0a, 0x13, 0x70,
	0x6c, 0x61, 0x79, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x3a, 0x72, 0x75, 0x6e, 0x3a, 0x72, 0x65,
	0x61, 0x64, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x10, 0x12, 0x0e, 0x2f, 0x72, 0x75, 0x6e, 0x73, 0x2f,
	0x7b, 0x72, 0x75, 0x6e, 0x5f, 0x69, 0x64, 0x7d, 0x12, 0x99, 0x01, 0x0a, 0x08, 0x57, 0x61, 0x74,
	0x63, 0x68, 0x52, 0x75, 0x6e, 0x12, 0x20, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x70,
	0x6c, 0x61, 0x79, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x2e, 0x57, 0x61, 0x74, 0x63, 0x68, 0x52,
	0x75, 0x6e, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x21, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70,
	0x2e, 0x70, 0x6c, 0x61, 0x79, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x2e, 0x57, 0x61, 0x74, 0x63,
	0x68, 0x52, 0x75, 0x6e, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x22, 0x46, 0xea, 0xe2, 0x0a, 0x0b,
	0xf5, 0xe2, 0x0a, 0x00, 0x00, 0x00, 0x40, 0xf8, 0xe2, 0x0a, 0x0a, 0x82, 0xe3, 0x0a, 0x17, 0x8a,
	0xe3, 0x0a, 0x13, 0x70, 0x6c, 0x61, 0x79, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x3a, 0x72, 0x75,
	0x6e, 0x3a, 0x72, 0x65, 0x61, 0x64, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x16, 0x12, 0x14, 0x2f, 0x72,
	0x75, 0x6e, 0x73, 0x2f, 0x7b, 0x72, 0x75, 0x6e, 0x5f, 0x69, 0x64, 0x7d, 0x2f, 0x77, 0x61, 0x74,
	0x63, 0x68, 0x30, 0x01, 0x42, 0x37, 0x5a, 0x35, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x67, 0x6f, 0x2d, 0x73, 0x64, 0x6b,
	0x2f, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x70, 0x6c, 0x61, 0x79, 0x67, 0x72, 0x6f, 0x75,
	0x6e, 0x64, 0x3b, 0x70, 0x6c, 0x61, 0x79, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
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

var file_eolymp_playground_playground_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_eolymp_playground_playground_proto_goTypes = []interface{}{
	(*CreateRunInput)(nil),    // 0: eolymp.playground.CreateRunInput
	(*CreateRunOutput)(nil),   // 1: eolymp.playground.CreateRunOutput
	(*DescribeRunInput)(nil),  // 2: eolymp.playground.DescribeRunInput
	(*DescribeRunOutput)(nil), // 3: eolymp.playground.DescribeRunOutput
	(*WatchRunInput)(nil),     // 4: eolymp.playground.WatchRunInput
	(*WatchRunOutput)(nil),    // 5: eolymp.playground.WatchRunOutput
	(*Run)(nil),               // 6: eolymp.playground.Run
}
var file_eolymp_playground_playground_proto_depIdxs = []int32{
	6, // 0: eolymp.playground.DescribeRunOutput.run:type_name -> eolymp.playground.Run
	6, // 1: eolymp.playground.WatchRunOutput.run:type_name -> eolymp.playground.Run
	0, // 2: eolymp.playground.Playground.CreateRun:input_type -> eolymp.playground.CreateRunInput
	2, // 3: eolymp.playground.Playground.DescribeRun:input_type -> eolymp.playground.DescribeRunInput
	4, // 4: eolymp.playground.Playground.WatchRun:input_type -> eolymp.playground.WatchRunInput
	1, // 5: eolymp.playground.Playground.CreateRun:output_type -> eolymp.playground.CreateRunOutput
	3, // 6: eolymp.playground.Playground.DescribeRun:output_type -> eolymp.playground.DescribeRunOutput
	5, // 7: eolymp.playground.Playground.WatchRun:output_type -> eolymp.playground.WatchRunOutput
	5, // [5:8] is the sub-list for method output_type
	2, // [2:5] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
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
		file_eolymp_playground_playground_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WatchRunInput); i {
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
		file_eolymp_playground_playground_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WatchRunOutput); i {
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
		(*CreateRunInput_InputData)(nil),
		(*CreateRunInput_InputRef)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_eolymp_playground_playground_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
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
