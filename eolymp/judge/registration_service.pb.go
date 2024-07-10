// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.21.12
// source: eolymp/judge/registration_service.proto

package judge

import (
	_ "github.com/eolymp/go-sdk/eolymp/annotations"
	community "github.com/eolymp/go-sdk/eolymp/community"
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

type DescribeRegistrationInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ContestId string `protobuf:"bytes,1,opt,name=contest_id,json=contestId,proto3" json:"contest_id,omitempty"`
}

func (x *DescribeRegistrationInput) Reset() {
	*x = DescribeRegistrationInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_judge_registration_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DescribeRegistrationInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescribeRegistrationInput) ProtoMessage() {}

func (x *DescribeRegistrationInput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_judge_registration_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DescribeRegistrationInput.ProtoReflect.Descriptor instead.
func (*DescribeRegistrationInput) Descriptor() ([]byte, []int) {
	return file_eolymp_judge_registration_service_proto_rawDescGZIP(), []int{0}
}

func (x *DescribeRegistrationInput) GetContestId() string {
	if x != nil {
		return x.ContestId
	}
	return ""
}

type DescribeRegistrationOutput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Attributes []*community.Attribute       `protobuf:"bytes,20,rep,name=attributes,proto3" json:"attributes,omitempty"`
	Values     []*community.Attribute_Value `protobuf:"bytes,21,rep,name=values,proto3" json:"values,omitempty"`
}

func (x *DescribeRegistrationOutput) Reset() {
	*x = DescribeRegistrationOutput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_judge_registration_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DescribeRegistrationOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescribeRegistrationOutput) ProtoMessage() {}

func (x *DescribeRegistrationOutput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_judge_registration_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DescribeRegistrationOutput.ProtoReflect.Descriptor instead.
func (*DescribeRegistrationOutput) Descriptor() ([]byte, []int) {
	return file_eolymp_judge_registration_service_proto_rawDescGZIP(), []int{1}
}

func (x *DescribeRegistrationOutput) GetAttributes() []*community.Attribute {
	if x != nil {
		return x.Attributes
	}
	return nil
}

func (x *DescribeRegistrationOutput) GetValues() []*community.Attribute_Value {
	if x != nil {
		return x.Values
	}
	return nil
}

type SubmitRegistrationInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ContestId string                       `protobuf:"bytes,1,opt,name=contest_id,json=contestId,proto3" json:"contest_id,omitempty"`
	Values    []*community.Attribute_Value `protobuf:"bytes,20,rep,name=values,proto3" json:"values,omitempty"`
}

func (x *SubmitRegistrationInput) Reset() {
	*x = SubmitRegistrationInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_judge_registration_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubmitRegistrationInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubmitRegistrationInput) ProtoMessage() {}

func (x *SubmitRegistrationInput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_judge_registration_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubmitRegistrationInput.ProtoReflect.Descriptor instead.
func (*SubmitRegistrationInput) Descriptor() ([]byte, []int) {
	return file_eolymp_judge_registration_service_proto_rawDescGZIP(), []int{2}
}

func (x *SubmitRegistrationInput) GetContestId() string {
	if x != nil {
		return x.ContestId
	}
	return ""
}

func (x *SubmitRegistrationInput) GetValues() []*community.Attribute_Value {
	if x != nil {
		return x.Values
	}
	return nil
}

type SubmitRegistrationOutput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SubmitRegistrationOutput) Reset() {
	*x = SubmitRegistrationOutput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_judge_registration_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubmitRegistrationOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubmitRegistrationOutput) ProtoMessage() {}

func (x *SubmitRegistrationOutput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_judge_registration_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubmitRegistrationOutput.ProtoReflect.Descriptor instead.
func (*SubmitRegistrationOutput) Descriptor() ([]byte, []int) {
	return file_eolymp_judge_registration_service_proto_rawDescGZIP(), []int{3}
}

var File_eolymp_judge_registration_service_proto protoreflect.FileDescriptor

var file_eolymp_judge_registration_service_proto_rawDesc = []byte{
	0x0a, 0x27, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x6a, 0x75, 0x64, 0x67, 0x65, 0x2f, 0x72,
	0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x65, 0x6f, 0x6c, 0x79, 0x6d,
	0x70, 0x2e, 0x6a, 0x75, 0x64, 0x67, 0x65, 0x1a, 0x1d, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x68, 0x74, 0x74, 0x70,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x22, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x61,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x72, 0x61, 0x74, 0x65, 0x6c,
	0x69, 0x6d, 0x69, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x65, 0x6f, 0x6c, 0x79,
	0x6d, 0x70, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73,
	0x63, 0x6f, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x65, 0x6f, 0x6c, 0x79,
	0x6d, 0x70, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x2f, 0x61, 0x74, 0x74,
	0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x3a, 0x0a, 0x19,
	0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6f, 0x6e,
	0x74, 0x65, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x73, 0x74, 0x49, 0x64, 0x22, 0x94, 0x01, 0x0a, 0x1a, 0x44, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x62, 0x65, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x12, 0x3b, 0x0a, 0x0a, 0x61, 0x74, 0x74, 0x72, 0x69,
	0x62, 0x75, 0x74, 0x65, 0x73, 0x18, 0x14, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x65, 0x6f,
	0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x2e, 0x41,
	0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x52, 0x0a, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62,
	0x75, 0x74, 0x65, 0x73, 0x12, 0x39, 0x0a, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x15,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x2e, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74,
	0x65, 0x2e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x22,
	0x73, 0x0a, 0x17, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x63, 0x6f, 0x6e, 0x74, 0x65, 0x73, 0x74, 0x49, 0x64, 0x12, 0x39, 0x0a, 0x06, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x73, 0x18, 0x14, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x65, 0x6f, 0x6c, 0x79,
	0x6d, 0x70, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x2e, 0x41, 0x74, 0x74,
	0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x2e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x06, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x73, 0x22, 0x1a, 0x0a, 0x18, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x52, 0x65,
	0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74,
	0x32, 0xf5, 0x02, 0x0a, 0x13, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0xb0, 0x01, 0x0a, 0x14, 0x44, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x62, 0x65, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x27, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x6a, 0x75, 0x64, 0x67, 0x65,
	0x2e, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x28, 0x2e, 0x65, 0x6f, 0x6c,
	0x79, 0x6d, 0x70, 0x2e, 0x6a, 0x75, 0x64, 0x67, 0x65, 0x2e, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x62, 0x65, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x75,
	0x74, 0x70, 0x75, 0x74, 0x22, 0x45, 0xea, 0xe2, 0x0a, 0x0b, 0xf5, 0xe2, 0x0a, 0x00, 0x00, 0xa0,
	0x40, 0xf8, 0xe2, 0x0a, 0x0a, 0x82, 0xe3, 0x0a, 0x1d, 0x8a, 0xe3, 0x0a, 0x19, 0x6a, 0x75, 0x64,
	0x67, 0x65, 0x3a, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x73, 0x74, 0x3a, 0x70, 0x61, 0x72, 0x74, 0x69,
	0x63, 0x69, 0x70, 0x61, 0x74, 0x65, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0f, 0x12, 0x0d, 0x2f, 0x72,
	0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0xaa, 0x01, 0x0a, 0x12,
	0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x25, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x6a, 0x75, 0x64, 0x67,
	0x65, 0x2e, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x26, 0x2e, 0x65, 0x6f, 0x6c, 0x79,
	0x6d, 0x70, 0x2e, 0x6a, 0x75, 0x64, 0x67, 0x65, 0x2e, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x52,
	0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x75, 0x74, 0x70, 0x75,
	0x74, 0x22, 0x45, 0xea, 0xe2, 0x0a, 0x0b, 0xf5, 0xe2, 0x0a, 0x00, 0x00, 0x80, 0x3f, 0xf8, 0xe2,
	0x0a, 0x03, 0x82, 0xe3, 0x0a, 0x1d, 0x8a, 0xe3, 0x0a, 0x19, 0x6a, 0x75, 0x64, 0x67, 0x65, 0x3a,
	0x63, 0x6f, 0x6e, 0x74, 0x65, 0x73, 0x74, 0x3a, 0x70, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70,
	0x61, 0x74, 0x65, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0f, 0x22, 0x0d, 0x2f, 0x72, 0x65, 0x67, 0x69,
	0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x2d, 0x5a, 0x2b, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x67, 0x6f,
	0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x6a, 0x75, 0x64, 0x67,
	0x65, 0x3b, 0x6a, 0x75, 0x64, 0x67, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_eolymp_judge_registration_service_proto_rawDescOnce sync.Once
	file_eolymp_judge_registration_service_proto_rawDescData = file_eolymp_judge_registration_service_proto_rawDesc
)

func file_eolymp_judge_registration_service_proto_rawDescGZIP() []byte {
	file_eolymp_judge_registration_service_proto_rawDescOnce.Do(func() {
		file_eolymp_judge_registration_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_eolymp_judge_registration_service_proto_rawDescData)
	})
	return file_eolymp_judge_registration_service_proto_rawDescData
}

var file_eolymp_judge_registration_service_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_eolymp_judge_registration_service_proto_goTypes = []any{
	(*DescribeRegistrationInput)(nil),  // 0: eolymp.judge.DescribeRegistrationInput
	(*DescribeRegistrationOutput)(nil), // 1: eolymp.judge.DescribeRegistrationOutput
	(*SubmitRegistrationInput)(nil),    // 2: eolymp.judge.SubmitRegistrationInput
	(*SubmitRegistrationOutput)(nil),   // 3: eolymp.judge.SubmitRegistrationOutput
	(*community.Attribute)(nil),        // 4: eolymp.community.Attribute
	(*community.Attribute_Value)(nil),  // 5: eolymp.community.Attribute.Value
}
var file_eolymp_judge_registration_service_proto_depIdxs = []int32{
	4, // 0: eolymp.judge.DescribeRegistrationOutput.attributes:type_name -> eolymp.community.Attribute
	5, // 1: eolymp.judge.DescribeRegistrationOutput.values:type_name -> eolymp.community.Attribute.Value
	5, // 2: eolymp.judge.SubmitRegistrationInput.values:type_name -> eolymp.community.Attribute.Value
	0, // 3: eolymp.judge.RegistrationService.DescribeRegistration:input_type -> eolymp.judge.DescribeRegistrationInput
	2, // 4: eolymp.judge.RegistrationService.SubmitRegistration:input_type -> eolymp.judge.SubmitRegistrationInput
	1, // 5: eolymp.judge.RegistrationService.DescribeRegistration:output_type -> eolymp.judge.DescribeRegistrationOutput
	3, // 6: eolymp.judge.RegistrationService.SubmitRegistration:output_type -> eolymp.judge.SubmitRegistrationOutput
	5, // [5:7] is the sub-list for method output_type
	3, // [3:5] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_eolymp_judge_registration_service_proto_init() }
func file_eolymp_judge_registration_service_proto_init() {
	if File_eolymp_judge_registration_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_eolymp_judge_registration_service_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*DescribeRegistrationInput); i {
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
		file_eolymp_judge_registration_service_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*DescribeRegistrationOutput); i {
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
		file_eolymp_judge_registration_service_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*SubmitRegistrationInput); i {
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
		file_eolymp_judge_registration_service_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*SubmitRegistrationOutput); i {
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
			RawDescriptor: file_eolymp_judge_registration_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_eolymp_judge_registration_service_proto_goTypes,
		DependencyIndexes: file_eolymp_judge_registration_service_proto_depIdxs,
		MessageInfos:      file_eolymp_judge_registration_service_proto_msgTypes,
	}.Build()
	File_eolymp_judge_registration_service_proto = out.File
	file_eolymp_judge_registration_service_proto_rawDesc = nil
	file_eolymp_judge_registration_service_proto_goTypes = nil
	file_eolymp_judge_registration_service_proto_depIdxs = nil
}