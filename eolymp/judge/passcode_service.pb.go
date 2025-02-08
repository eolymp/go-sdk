// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v4.24.4
// source: eolymp/judge/passcode_service.proto

package judge

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

type VerifyPasscodeInput struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ContestId     string                 `protobuf:"bytes,1,opt,name=contest_id,json=contestId,proto3" json:"contest_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *VerifyPasscodeInput) Reset() {
	*x = VerifyPasscodeInput{}
	mi := &file_eolymp_judge_passcode_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *VerifyPasscodeInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VerifyPasscodeInput) ProtoMessage() {}

func (x *VerifyPasscodeInput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_judge_passcode_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VerifyPasscodeInput.ProtoReflect.Descriptor instead.
func (*VerifyPasscodeInput) Descriptor() ([]byte, []int) {
	return file_eolymp_judge_passcode_service_proto_rawDescGZIP(), []int{0}
}

func (x *VerifyPasscodeInput) GetContestId() string {
	if x != nil {
		return x.ContestId
	}
	return ""
}

type VerifyPasscodeOutput struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Required      bool                   `protobuf:"varint,1,opt,name=required,proto3" json:"required,omitempty"`
	Valid         bool                   `protobuf:"varint,2,opt,name=valid,proto3" json:"valid,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *VerifyPasscodeOutput) Reset() {
	*x = VerifyPasscodeOutput{}
	mi := &file_eolymp_judge_passcode_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *VerifyPasscodeOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VerifyPasscodeOutput) ProtoMessage() {}

func (x *VerifyPasscodeOutput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_judge_passcode_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VerifyPasscodeOutput.ProtoReflect.Descriptor instead.
func (*VerifyPasscodeOutput) Descriptor() ([]byte, []int) {
	return file_eolymp_judge_passcode_service_proto_rawDescGZIP(), []int{1}
}

func (x *VerifyPasscodeOutput) GetRequired() bool {
	if x != nil {
		return x.Required
	}
	return false
}

func (x *VerifyPasscodeOutput) GetValid() bool {
	if x != nil {
		return x.Valid
	}
	return false
}

type EnterPasscodeInput struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ContestId     string                 `protobuf:"bytes,1,opt,name=contest_id,json=contestId,proto3" json:"contest_id,omitempty"`
	Passcode      string                 `protobuf:"bytes,2,opt,name=passcode,proto3" json:"passcode,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *EnterPasscodeInput) Reset() {
	*x = EnterPasscodeInput{}
	mi := &file_eolymp_judge_passcode_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *EnterPasscodeInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EnterPasscodeInput) ProtoMessage() {}

func (x *EnterPasscodeInput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_judge_passcode_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EnterPasscodeInput.ProtoReflect.Descriptor instead.
func (*EnterPasscodeInput) Descriptor() ([]byte, []int) {
	return file_eolymp_judge_passcode_service_proto_rawDescGZIP(), []int{2}
}

func (x *EnterPasscodeInput) GetContestId() string {
	if x != nil {
		return x.ContestId
	}
	return ""
}

func (x *EnterPasscodeInput) GetPasscode() string {
	if x != nil {
		return x.Passcode
	}
	return ""
}

type EnterPasscodeOutput struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *EnterPasscodeOutput) Reset() {
	*x = EnterPasscodeOutput{}
	mi := &file_eolymp_judge_passcode_service_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *EnterPasscodeOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EnterPasscodeOutput) ProtoMessage() {}

func (x *EnterPasscodeOutput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_judge_passcode_service_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EnterPasscodeOutput.ProtoReflect.Descriptor instead.
func (*EnterPasscodeOutput) Descriptor() ([]byte, []int) {
	return file_eolymp_judge_passcode_service_proto_rawDescGZIP(), []int{3}
}

type ResetPasscodeInput struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ContestId     string                 `protobuf:"bytes,1,opt,name=contest_id,json=contestId,proto3" json:"contest_id,omitempty"`
	ParticipantId string                 `protobuf:"bytes,2,opt,name=participant_id,json=participantId,proto3" json:"participant_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ResetPasscodeInput) Reset() {
	*x = ResetPasscodeInput{}
	mi := &file_eolymp_judge_passcode_service_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ResetPasscodeInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResetPasscodeInput) ProtoMessage() {}

func (x *ResetPasscodeInput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_judge_passcode_service_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResetPasscodeInput.ProtoReflect.Descriptor instead.
func (*ResetPasscodeInput) Descriptor() ([]byte, []int) {
	return file_eolymp_judge_passcode_service_proto_rawDescGZIP(), []int{4}
}

func (x *ResetPasscodeInput) GetContestId() string {
	if x != nil {
		return x.ContestId
	}
	return ""
}

func (x *ResetPasscodeInput) GetParticipantId() string {
	if x != nil {
		return x.ParticipantId
	}
	return ""
}

type ResetPasscodeOutput struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Passcode      string                 `protobuf:"bytes,1,opt,name=passcode,proto3" json:"passcode,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ResetPasscodeOutput) Reset() {
	*x = ResetPasscodeOutput{}
	mi := &file_eolymp_judge_passcode_service_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ResetPasscodeOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResetPasscodeOutput) ProtoMessage() {}

func (x *ResetPasscodeOutput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_judge_passcode_service_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResetPasscodeOutput.ProtoReflect.Descriptor instead.
func (*ResetPasscodeOutput) Descriptor() ([]byte, []int) {
	return file_eolymp_judge_passcode_service_proto_rawDescGZIP(), []int{5}
}

func (x *ResetPasscodeOutput) GetPasscode() string {
	if x != nil {
		return x.Passcode
	}
	return ""
}

type SetPasscodeInput struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ContestId     string                 `protobuf:"bytes,1,opt,name=contest_id,json=contestId,proto3" json:"contest_id,omitempty"`
	ParticipantId string                 `protobuf:"bytes,2,opt,name=participant_id,json=participantId,proto3" json:"participant_id,omitempty"`
	Passcode      string                 `protobuf:"bytes,3,opt,name=passcode,proto3" json:"passcode,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SetPasscodeInput) Reset() {
	*x = SetPasscodeInput{}
	mi := &file_eolymp_judge_passcode_service_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SetPasscodeInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetPasscodeInput) ProtoMessage() {}

func (x *SetPasscodeInput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_judge_passcode_service_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetPasscodeInput.ProtoReflect.Descriptor instead.
func (*SetPasscodeInput) Descriptor() ([]byte, []int) {
	return file_eolymp_judge_passcode_service_proto_rawDescGZIP(), []int{6}
}

func (x *SetPasscodeInput) GetContestId() string {
	if x != nil {
		return x.ContestId
	}
	return ""
}

func (x *SetPasscodeInput) GetParticipantId() string {
	if x != nil {
		return x.ParticipantId
	}
	return ""
}

func (x *SetPasscodeInput) GetPasscode() string {
	if x != nil {
		return x.Passcode
	}
	return ""
}

type SetPasscodeOutput struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SetPasscodeOutput) Reset() {
	*x = SetPasscodeOutput{}
	mi := &file_eolymp_judge_passcode_service_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SetPasscodeOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetPasscodeOutput) ProtoMessage() {}

func (x *SetPasscodeOutput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_judge_passcode_service_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetPasscodeOutput.ProtoReflect.Descriptor instead.
func (*SetPasscodeOutput) Descriptor() ([]byte, []int) {
	return file_eolymp_judge_passcode_service_proto_rawDescGZIP(), []int{7}
}

type RemovePasscodeInput struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ContestId     string                 `protobuf:"bytes,1,opt,name=contest_id,json=contestId,proto3" json:"contest_id,omitempty"`
	ParticipantId string                 `protobuf:"bytes,2,opt,name=participant_id,json=participantId,proto3" json:"participant_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RemovePasscodeInput) Reset() {
	*x = RemovePasscodeInput{}
	mi := &file_eolymp_judge_passcode_service_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RemovePasscodeInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemovePasscodeInput) ProtoMessage() {}

func (x *RemovePasscodeInput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_judge_passcode_service_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemovePasscodeInput.ProtoReflect.Descriptor instead.
func (*RemovePasscodeInput) Descriptor() ([]byte, []int) {
	return file_eolymp_judge_passcode_service_proto_rawDescGZIP(), []int{8}
}

func (x *RemovePasscodeInput) GetContestId() string {
	if x != nil {
		return x.ContestId
	}
	return ""
}

func (x *RemovePasscodeInput) GetParticipantId() string {
	if x != nil {
		return x.ParticipantId
	}
	return ""
}

type RemovePasscodeOutput struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RemovePasscodeOutput) Reset() {
	*x = RemovePasscodeOutput{}
	mi := &file_eolymp_judge_passcode_service_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RemovePasscodeOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemovePasscodeOutput) ProtoMessage() {}

func (x *RemovePasscodeOutput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_judge_passcode_service_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemovePasscodeOutput.ProtoReflect.Descriptor instead.
func (*RemovePasscodeOutput) Descriptor() ([]byte, []int) {
	return file_eolymp_judge_passcode_service_proto_rawDescGZIP(), []int{9}
}

var File_eolymp_judge_passcode_service_proto protoreflect.FileDescriptor

var file_eolymp_judge_passcode_service_proto_rawDesc = string([]byte{
	0x0a, 0x23, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x6a, 0x75, 0x64, 0x67, 0x65, 0x2f, 0x70,
	0x61, 0x73, 0x73, 0x63, 0x6f, 0x64, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x6a, 0x75,
	0x64, 0x67, 0x65, 0x1a, 0x1d, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x22, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x72, 0x61, 0x74, 0x65, 0x6c, 0x69, 0x6d, 0x69, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x61,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x63, 0x6f, 0x70, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x34, 0x0a, 0x13, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79,
	0x50, 0x61, 0x73, 0x73, 0x63, 0x6f, 0x64, 0x65, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x1d, 0x0a,
	0x0a, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x73, 0x74, 0x49, 0x64, 0x22, 0x48, 0x0a, 0x14,
	0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x50, 0x61, 0x73, 0x73, 0x63, 0x6f, 0x64, 0x65, 0x4f, 0x75,
	0x74, 0x70, 0x75, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64,
	0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x22, 0x4f, 0x0a, 0x12, 0x45, 0x6e, 0x74, 0x65, 0x72, 0x50,
	0x61, 0x73, 0x73, 0x63, 0x6f, 0x64, 0x65, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x1d, 0x0a, 0x0a,
	0x63, 0x6f, 0x6e, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x73, 0x74, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x70,
	0x61, 0x73, 0x73, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70,
	0x61, 0x73, 0x73, 0x63, 0x6f, 0x64, 0x65, 0x22, 0x15, 0x0a, 0x13, 0x45, 0x6e, 0x74, 0x65, 0x72,
	0x50, 0x61, 0x73, 0x73, 0x63, 0x6f, 0x64, 0x65, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x22, 0x5a,
	0x0a, 0x12, 0x52, 0x65, 0x73, 0x65, 0x74, 0x50, 0x61, 0x73, 0x73, 0x63, 0x6f, 0x64, 0x65, 0x49,
	0x6e, 0x70, 0x75, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x73, 0x74, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x73,
	0x74, 0x49, 0x64, 0x12, 0x25, 0x0a, 0x0e, 0x70, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61,
	0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x70, 0x61, 0x72,
	0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x49, 0x64, 0x22, 0x31, 0x0a, 0x13, 0x52, 0x65,
	0x73, 0x65, 0x74, 0x50, 0x61, 0x73, 0x73, 0x63, 0x6f, 0x64, 0x65, 0x4f, 0x75, 0x74, 0x70, 0x75,
	0x74, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x63, 0x6f, 0x64, 0x65, 0x22, 0x74, 0x0a,
	0x10, 0x53, 0x65, 0x74, 0x50, 0x61, 0x73, 0x73, 0x63, 0x6f, 0x64, 0x65, 0x49, 0x6e, 0x70, 0x75,
	0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x73, 0x74, 0x49, 0x64,
	0x12, 0x25, 0x0a, 0x0e, 0x70, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x70, 0x61, 0x72, 0x74, 0x69, 0x63,
	0x69, 0x70, 0x61, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x63,
	0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x63,
	0x6f, 0x64, 0x65, 0x22, 0x13, 0x0a, 0x11, 0x53, 0x65, 0x74, 0x50, 0x61, 0x73, 0x73, 0x63, 0x6f,
	0x64, 0x65, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x22, 0x5b, 0x0a, 0x13, 0x52, 0x65, 0x6d, 0x6f,
	0x76, 0x65, 0x50, 0x61, 0x73, 0x73, 0x63, 0x6f, 0x64, 0x65, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12,
	0x1d, 0x0a, 0x0a, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x73, 0x74, 0x49, 0x64, 0x12, 0x25,
	0x0a, 0x0e, 0x70, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x5f, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x70, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70,
	0x61, 0x6e, 0x74, 0x49, 0x64, 0x22, 0x16, 0x0a, 0x14, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x50,
	0x61, 0x73, 0x73, 0x63, 0x6f, 0x64, 0x65, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x32, 0xe1, 0x06,
	0x0a, 0x0f, 0x50, 0x61, 0x73, 0x73, 0x63, 0x6f, 0x64, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x9a, 0x01, 0x0a, 0x0e, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x50, 0x61, 0x73, 0x73,
	0x63, 0x6f, 0x64, 0x65, 0x12, 0x21, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x6a, 0x75,
	0x64, 0x67, 0x65, 0x2e, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x50, 0x61, 0x73, 0x73, 0x63, 0x6f,
	0x64, 0x65, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x22, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70,
	0x2e, 0x6a, 0x75, 0x64, 0x67, 0x65, 0x2e, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x50, 0x61, 0x73,
	0x73, 0x63, 0x6f, 0x64, 0x65, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x22, 0x41, 0xea, 0xe2, 0x0a,
	0x0b, 0xf5, 0xe2, 0x0a, 0x00, 0x00, 0x20, 0x41, 0xf8, 0xe2, 0x0a, 0x32, 0x82, 0xe3, 0x0a, 0x16,
	0x8a, 0xe3, 0x0a, 0x12, 0x6a, 0x75, 0x64, 0x67, 0x65, 0x3a, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x73,
	0x74, 0x3a, 0x72, 0x65, 0x61, 0x64, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x12, 0x22, 0x10, 0x2f, 0x76,
	0x65, 0x72, 0x69, 0x66, 0x79, 0x2d, 0x70, 0x61, 0x73, 0x73, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x9d,
	0x01, 0x0a, 0x0d, 0x45, 0x6e, 0x74, 0x65, 0x72, 0x50, 0x61, 0x73, 0x73, 0x63, 0x6f, 0x64, 0x65,
	0x12, 0x20, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x6a, 0x75, 0x64, 0x67, 0x65, 0x2e,
	0x45, 0x6e, 0x74, 0x65, 0x72, 0x50, 0x61, 0x73, 0x73, 0x63, 0x6f, 0x64, 0x65, 0x49, 0x6e, 0x70,
	0x75, 0x74, 0x1a, 0x21, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x6a, 0x75, 0x64, 0x67,
	0x65, 0x2e, 0x45, 0x6e, 0x74, 0x65, 0x72, 0x50, 0x61, 0x73, 0x73, 0x63, 0x6f, 0x64, 0x65, 0x4f,
	0x75, 0x74, 0x70, 0x75, 0x74, 0x22, 0x47, 0xea, 0xe2, 0x0a, 0x0b, 0xf5, 0xe2, 0x0a, 0x00, 0x00,
	0x20, 0x41, 0xf8, 0xe2, 0x0a, 0x32, 0x82, 0xe3, 0x0a, 0x1d, 0x8a, 0xe3, 0x0a, 0x19, 0x6a, 0x75,
	0x64, 0x67, 0x65, 0x3a, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x73, 0x74, 0x3a, 0x70, 0x61, 0x72, 0x74,
	0x69, 0x63, 0x69, 0x70, 0x61, 0x74, 0x65, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x11, 0x22, 0x0f, 0x2f,
	0x65, 0x6e, 0x74, 0x65, 0x72, 0x2d, 0x70, 0x61, 0x73, 0x73, 0x63, 0x6f, 0x64, 0x65, 0x12, 0xaf,
	0x01, 0x0a, 0x0d, 0x52, 0x65, 0x73, 0x65, 0x74, 0x50, 0x61, 0x73, 0x73, 0x63, 0x6f, 0x64, 0x65,
	0x12, 0x20, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x6a, 0x75, 0x64, 0x67, 0x65, 0x2e,
	0x52, 0x65, 0x73, 0x65, 0x74, 0x50, 0x61, 0x73, 0x73, 0x63, 0x6f, 0x64, 0x65, 0x49, 0x6e, 0x70,
	0x75, 0x74, 0x1a, 0x21, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x6a, 0x75, 0x64, 0x67,
	0x65, 0x2e, 0x52, 0x65, 0x73, 0x65, 0x74, 0x50, 0x61, 0x73, 0x73, 0x63, 0x6f, 0x64, 0x65, 0x4f,
	0x75, 0x74, 0x70, 0x75, 0x74, 0x22, 0x59, 0xea, 0xe2, 0x0a, 0x0b, 0xf5, 0xe2, 0x0a, 0x00, 0x00,
	0x20, 0x41, 0xf8, 0xe2, 0x0a, 0x32, 0x82, 0xe3, 0x0a, 0x17, 0x8a, 0xe3, 0x0a, 0x13, 0x6a, 0x75,
	0x64, 0x67, 0x65, 0x3a, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x73, 0x74, 0x3a, 0x77, 0x72, 0x69, 0x74,
	0x65, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x29, 0x22, 0x27, 0x2f, 0x70, 0x61, 0x72, 0x74, 0x69, 0x63,
	0x69, 0x70, 0x61, 0x6e, 0x74, 0x73, 0x2f, 0x7b, 0x70, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70,
	0x61, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x7d, 0x2f, 0x70, 0x61, 0x73, 0x73, 0x63, 0x6f, 0x64, 0x65,
	0x12, 0xa9, 0x01, 0x0a, 0x0b, 0x53, 0x65, 0x74, 0x50, 0x61, 0x73, 0x73, 0x63, 0x6f, 0x64, 0x65,
	0x12, 0x1e, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x6a, 0x75, 0x64, 0x67, 0x65, 0x2e,
	0x53, 0x65, 0x74, 0x50, 0x61, 0x73, 0x73, 0x63, 0x6f, 0x64, 0x65, 0x49, 0x6e, 0x70, 0x75, 0x74,
	0x1a, 0x1f, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x6a, 0x75, 0x64, 0x67, 0x65, 0x2e,
	0x53, 0x65, 0x74, 0x50, 0x61, 0x73, 0x73, 0x63, 0x6f, 0x64, 0x65, 0x4f, 0x75, 0x74, 0x70, 0x75,
	0x74, 0x22, 0x59, 0xea, 0xe2, 0x0a, 0x0b, 0xf5, 0xe2, 0x0a, 0x00, 0x00, 0x20, 0x41, 0xf8, 0xe2,
	0x0a, 0x32, 0x82, 0xe3, 0x0a, 0x17, 0x8a, 0xe3, 0x0a, 0x13, 0x6a, 0x75, 0x64, 0x67, 0x65, 0x3a,
	0x63, 0x6f, 0x6e, 0x74, 0x65, 0x73, 0x74, 0x3a, 0x77, 0x72, 0x69, 0x74, 0x65, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x29, 0x1a, 0x27, 0x2f, 0x70, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e,
	0x74, 0x73, 0x2f, 0x7b, 0x70, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x5f,
	0x69, 0x64, 0x7d, 0x2f, 0x70, 0x61, 0x73, 0x73, 0x63, 0x6f, 0x64, 0x65, 0x12, 0xb2, 0x01, 0x0a,
	0x0e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x50, 0x61, 0x73, 0x73, 0x63, 0x6f, 0x64, 0x65, 0x12,
	0x21, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x6a, 0x75, 0x64, 0x67, 0x65, 0x2e, 0x52,
	0x65, 0x6d, 0x6f, 0x76, 0x65, 0x50, 0x61, 0x73, 0x73, 0x63, 0x6f, 0x64, 0x65, 0x49, 0x6e, 0x70,
	0x75, 0x74, 0x1a, 0x22, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x6a, 0x75, 0x64, 0x67,
	0x65, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x50, 0x61, 0x73, 0x73, 0x63, 0x6f, 0x64, 0x65,
	0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x22, 0x59, 0xea, 0xe2, 0x0a, 0x0b, 0xf5, 0xe2, 0x0a, 0x00,
	0x00, 0x20, 0x41, 0xf8, 0xe2, 0x0a, 0x32, 0x82, 0xe3, 0x0a, 0x17, 0x8a, 0xe3, 0x0a, 0x13, 0x6a,
	0x75, 0x64, 0x67, 0x65, 0x3a, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x73, 0x74, 0x3a, 0x77, 0x72, 0x69,
	0x74, 0x65, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x29, 0x2a, 0x27, 0x2f, 0x70, 0x61, 0x72, 0x74, 0x69,
	0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x73, 0x2f, 0x7b, 0x70, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69,
	0x70, 0x61, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x7d, 0x2f, 0x70, 0x61, 0x73, 0x73, 0x63, 0x6f, 0x64,
	0x65, 0x42, 0x2d, 0x5a, 0x2b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x67, 0x6f, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x65, 0x6f,
	0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x6a, 0x75, 0x64, 0x67, 0x65, 0x3b, 0x6a, 0x75, 0x64, 0x67, 0x65,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_eolymp_judge_passcode_service_proto_rawDescOnce sync.Once
	file_eolymp_judge_passcode_service_proto_rawDescData []byte
)

func file_eolymp_judge_passcode_service_proto_rawDescGZIP() []byte {
	file_eolymp_judge_passcode_service_proto_rawDescOnce.Do(func() {
		file_eolymp_judge_passcode_service_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_eolymp_judge_passcode_service_proto_rawDesc), len(file_eolymp_judge_passcode_service_proto_rawDesc)))
	})
	return file_eolymp_judge_passcode_service_proto_rawDescData
}

var file_eolymp_judge_passcode_service_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_eolymp_judge_passcode_service_proto_goTypes = []any{
	(*VerifyPasscodeInput)(nil),  // 0: eolymp.judge.VerifyPasscodeInput
	(*VerifyPasscodeOutput)(nil), // 1: eolymp.judge.VerifyPasscodeOutput
	(*EnterPasscodeInput)(nil),   // 2: eolymp.judge.EnterPasscodeInput
	(*EnterPasscodeOutput)(nil),  // 3: eolymp.judge.EnterPasscodeOutput
	(*ResetPasscodeInput)(nil),   // 4: eolymp.judge.ResetPasscodeInput
	(*ResetPasscodeOutput)(nil),  // 5: eolymp.judge.ResetPasscodeOutput
	(*SetPasscodeInput)(nil),     // 6: eolymp.judge.SetPasscodeInput
	(*SetPasscodeOutput)(nil),    // 7: eolymp.judge.SetPasscodeOutput
	(*RemovePasscodeInput)(nil),  // 8: eolymp.judge.RemovePasscodeInput
	(*RemovePasscodeOutput)(nil), // 9: eolymp.judge.RemovePasscodeOutput
}
var file_eolymp_judge_passcode_service_proto_depIdxs = []int32{
	0, // 0: eolymp.judge.PasscodeService.VerifyPasscode:input_type -> eolymp.judge.VerifyPasscodeInput
	2, // 1: eolymp.judge.PasscodeService.EnterPasscode:input_type -> eolymp.judge.EnterPasscodeInput
	4, // 2: eolymp.judge.PasscodeService.ResetPasscode:input_type -> eolymp.judge.ResetPasscodeInput
	6, // 3: eolymp.judge.PasscodeService.SetPasscode:input_type -> eolymp.judge.SetPasscodeInput
	8, // 4: eolymp.judge.PasscodeService.RemovePasscode:input_type -> eolymp.judge.RemovePasscodeInput
	1, // 5: eolymp.judge.PasscodeService.VerifyPasscode:output_type -> eolymp.judge.VerifyPasscodeOutput
	3, // 6: eolymp.judge.PasscodeService.EnterPasscode:output_type -> eolymp.judge.EnterPasscodeOutput
	5, // 7: eolymp.judge.PasscodeService.ResetPasscode:output_type -> eolymp.judge.ResetPasscodeOutput
	7, // 8: eolymp.judge.PasscodeService.SetPasscode:output_type -> eolymp.judge.SetPasscodeOutput
	9, // 9: eolymp.judge.PasscodeService.RemovePasscode:output_type -> eolymp.judge.RemovePasscodeOutput
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_eolymp_judge_passcode_service_proto_init() }
func file_eolymp_judge_passcode_service_proto_init() {
	if File_eolymp_judge_passcode_service_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_eolymp_judge_passcode_service_proto_rawDesc), len(file_eolymp_judge_passcode_service_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_eolymp_judge_passcode_service_proto_goTypes,
		DependencyIndexes: file_eolymp_judge_passcode_service_proto_depIdxs,
		MessageInfos:      file_eolymp_judge_passcode_service_proto_msgTypes,
	}.Build()
	File_eolymp_judge_passcode_service_proto = out.File
	file_eolymp_judge_passcode_service_proto_goTypes = nil
	file_eolymp_judge_passcode_service_proto_depIdxs = nil
}
