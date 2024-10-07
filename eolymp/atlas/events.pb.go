// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v3.21.12
// source: eolymp/atlas/events.proto

package atlas

import (
	acl "github.com/eolymp/go-sdk/eolymp/acl"
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

// SubmissionCompleteEvent is dispatched every time submission is fully judged (got final result).
type SubmissionCompleteEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Submission *Submission `protobuf:"bytes,1,opt,name=submission,proto3" json:"submission,omitempty"` // submission
	Update     bool        `protobuf:"varint,2,opt,name=update,proto3" json:"update,omitempty"`        // true, if submission has been completed before, this flag is set when submission is being retested
}

func (x *SubmissionCompleteEvent) Reset() {
	*x = SubmissionCompleteEvent{}
	mi := &file_eolymp_atlas_events_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SubmissionCompleteEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubmissionCompleteEvent) ProtoMessage() {}

func (x *SubmissionCompleteEvent) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_atlas_events_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubmissionCompleteEvent.ProtoReflect.Descriptor instead.
func (*SubmissionCompleteEvent) Descriptor() ([]byte, []int) {
	return file_eolymp_atlas_events_proto_rawDescGZIP(), []int{0}
}

func (x *SubmissionCompleteEvent) GetSubmission() *Submission {
	if x != nil {
		return x.Submission
	}
	return nil
}

func (x *SubmissionCompleteEvent) GetUpdate() bool {
	if x != nil {
		return x.Update
	}
	return false
}

type SubmissionChangedEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Before *Submission `protobuf:"bytes,1,opt,name=before,proto3" json:"before,omitempty"`
	After  *Submission `protobuf:"bytes,2,opt,name=after,proto3" json:"after,omitempty"`
}

func (x *SubmissionChangedEvent) Reset() {
	*x = SubmissionChangedEvent{}
	mi := &file_eolymp_atlas_events_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SubmissionChangedEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubmissionChangedEvent) ProtoMessage() {}

func (x *SubmissionChangedEvent) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_atlas_events_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubmissionChangedEvent.ProtoReflect.Descriptor instead.
func (*SubmissionChangedEvent) Descriptor() ([]byte, []int) {
	return file_eolymp_atlas_events_proto_rawDescGZIP(), []int{1}
}

func (x *SubmissionChangedEvent) GetBefore() *Submission {
	if x != nil {
		return x.Before
	}
	return nil
}

func (x *SubmissionChangedEvent) GetAfter() *Submission {
	if x != nil {
		return x.After
	}
	return nil
}

// ScoreUpdatedEvent is dispatched when score for a given problem and user has been updated (increased).
type ScoreUpdatedEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Score *Score `protobuf:"bytes,1,opt,name=score,proto3" json:"score,omitempty"`
}

func (x *ScoreUpdatedEvent) Reset() {
	*x = ScoreUpdatedEvent{}
	mi := &file_eolymp_atlas_events_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ScoreUpdatedEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ScoreUpdatedEvent) ProtoMessage() {}

func (x *ScoreUpdatedEvent) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_atlas_events_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ScoreUpdatedEvent.ProtoReflect.Descriptor instead.
func (*ScoreUpdatedEvent) Descriptor() ([]byte, []int) {
	return file_eolymp_atlas_events_proto_rawDescGZIP(), []int{2}
}

func (x *ScoreUpdatedEvent) GetScore() *Score {
	if x != nil {
		return x.Score
	}
	return nil
}

type ProblemChangedEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Before *Problem `protobuf:"bytes,1,opt,name=before,proto3" json:"before,omitempty"`
	After  *Problem `protobuf:"bytes,2,opt,name=after,proto3" json:"after,omitempty"`
}

func (x *ProblemChangedEvent) Reset() {
	*x = ProblemChangedEvent{}
	mi := &file_eolymp_atlas_events_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ProblemChangedEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProblemChangedEvent) ProtoMessage() {}

func (x *ProblemChangedEvent) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_atlas_events_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProblemChangedEvent.ProtoReflect.Descriptor instead.
func (*ProblemChangedEvent) Descriptor() ([]byte, []int) {
	return file_eolymp_atlas_events_proto_rawDescGZIP(), []int{3}
}

func (x *ProblemChangedEvent) GetBefore() *Problem {
	if x != nil {
		return x.Before
	}
	return nil
}

func (x *ProblemChangedEvent) GetAfter() *Problem {
	if x != nil {
		return x.After
	}
	return nil
}

type StatementChangedEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProblemId string     `protobuf:"bytes,1,opt,name=problem_id,json=problemId,proto3" json:"problem_id,omitempty"`
	Before    *Statement `protobuf:"bytes,2,opt,name=before,proto3" json:"before,omitempty"`
	After     *Statement `protobuf:"bytes,3,opt,name=after,proto3" json:"after,omitempty"`
}

func (x *StatementChangedEvent) Reset() {
	*x = StatementChangedEvent{}
	mi := &file_eolymp_atlas_events_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *StatementChangedEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatementChangedEvent) ProtoMessage() {}

func (x *StatementChangedEvent) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_atlas_events_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatementChangedEvent.ProtoReflect.Descriptor instead.
func (*StatementChangedEvent) Descriptor() ([]byte, []int) {
	return file_eolymp_atlas_events_proto_rawDescGZIP(), []int{4}
}

func (x *StatementChangedEvent) GetProblemId() string {
	if x != nil {
		return x.ProblemId
	}
	return ""
}

func (x *StatementChangedEvent) GetBefore() *Statement {
	if x != nil {
		return x.Before
	}
	return nil
}

func (x *StatementChangedEvent) GetAfter() *Statement {
	if x != nil {
		return x.After
	}
	return nil
}

type BookmarkChangedEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProblemId string `protobuf:"bytes,1,opt,name=problem_id,json=problemId,proto3" json:"problem_id,omitempty"`
	MemberId  string `protobuf:"bytes,2,opt,name=member_id,json=memberId,proto3" json:"member_id,omitempty"`
	Before    bool   `protobuf:"varint,3,opt,name=before,proto3" json:"before,omitempty"`
	After     bool   `protobuf:"varint,4,opt,name=after,proto3" json:"after,omitempty"`
}

func (x *BookmarkChangedEvent) Reset() {
	*x = BookmarkChangedEvent{}
	mi := &file_eolymp_atlas_events_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BookmarkChangedEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BookmarkChangedEvent) ProtoMessage() {}

func (x *BookmarkChangedEvent) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_atlas_events_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BookmarkChangedEvent.ProtoReflect.Descriptor instead.
func (*BookmarkChangedEvent) Descriptor() ([]byte, []int) {
	return file_eolymp_atlas_events_proto_rawDescGZIP(), []int{5}
}

func (x *BookmarkChangedEvent) GetProblemId() string {
	if x != nil {
		return x.ProblemId
	}
	return ""
}

func (x *BookmarkChangedEvent) GetMemberId() string {
	if x != nil {
		return x.MemberId
	}
	return ""
}

func (x *BookmarkChangedEvent) GetBefore() bool {
	if x != nil {
		return x.Before
	}
	return false
}

func (x *BookmarkChangedEvent) GetAfter() bool {
	if x != nil {
		return x.After
	}
	return false
}

type PermissionChangedEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProblemId string          `protobuf:"bytes,1,opt,name=problem_id,json=problemId,proto3" json:"problem_id,omitempty"`
	UserId    string          `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Before    *acl.Permission `protobuf:"bytes,3,opt,name=before,proto3" json:"before,omitempty"`
	After     *acl.Permission `protobuf:"bytes,4,opt,name=after,proto3" json:"after,omitempty"`
}

func (x *PermissionChangedEvent) Reset() {
	*x = PermissionChangedEvent{}
	mi := &file_eolymp_atlas_events_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PermissionChangedEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PermissionChangedEvent) ProtoMessage() {}

func (x *PermissionChangedEvent) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_atlas_events_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PermissionChangedEvent.ProtoReflect.Descriptor instead.
func (*PermissionChangedEvent) Descriptor() ([]byte, []int) {
	return file_eolymp_atlas_events_proto_rawDescGZIP(), []int{6}
}

func (x *PermissionChangedEvent) GetProblemId() string {
	if x != nil {
		return x.ProblemId
	}
	return ""
}

func (x *PermissionChangedEvent) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *PermissionChangedEvent) GetBefore() *acl.Permission {
	if x != nil {
		return x.Before
	}
	return nil
}

func (x *PermissionChangedEvent) GetAfter() *acl.Permission {
	if x != nil {
		return x.After
	}
	return nil
}

type SuggestionChangedEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProblemId string      `protobuf:"bytes,1,opt,name=problem_id,json=problemId,proto3" json:"problem_id,omitempty"`
	Before    *Suggestion `protobuf:"bytes,2,opt,name=before,proto3" json:"before,omitempty"`
	After     *Suggestion `protobuf:"bytes,3,opt,name=after,proto3" json:"after,omitempty"`
}

func (x *SuggestionChangedEvent) Reset() {
	*x = SuggestionChangedEvent{}
	mi := &file_eolymp_atlas_events_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SuggestionChangedEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SuggestionChangedEvent) ProtoMessage() {}

func (x *SuggestionChangedEvent) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_atlas_events_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SuggestionChangedEvent.ProtoReflect.Descriptor instead.
func (*SuggestionChangedEvent) Descriptor() ([]byte, []int) {
	return file_eolymp_atlas_events_proto_rawDescGZIP(), []int{7}
}

func (x *SuggestionChangedEvent) GetProblemId() string {
	if x != nil {
		return x.ProblemId
	}
	return ""
}

func (x *SuggestionChangedEvent) GetBefore() *Suggestion {
	if x != nil {
		return x.Before
	}
	return nil
}

func (x *SuggestionChangedEvent) GetAfter() *Suggestion {
	if x != nil {
		return x.After
	}
	return nil
}

var File_eolymp_atlas_events_proto protoreflect.FileDescriptor

var file_eolymp_atlas_events_proto_rawDesc = []byte{
	0x0a, 0x19, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x61, 0x74, 0x6c, 0x61, 0x73, 0x2f, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x65, 0x6f, 0x6c,
	0x79, 0x6d, 0x70, 0x2e, 0x61, 0x74, 0x6c, 0x61, 0x73, 0x1a, 0x1c, 0x65, 0x6f, 0x6c, 0x79, 0x6d,
	0x70, 0x2f, 0x61, 0x63, 0x6c, 0x2f, 0x61, 0x63, 0x6c, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1a, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f,
	0x61, 0x74, 0x6c, 0x61, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x61, 0x74, 0x6c, 0x61,
	0x73, 0x2f, 0x73, 0x63, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x61, 0x74,
	0x6c, 0x61, 0x73, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x61, 0x74, 0x6c, 0x61,
	0x73, 0x2f, 0x73, 0x75, 0x62, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1d, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x61, 0x74, 0x6c, 0x61, 0x73,
	0x2f, 0x73, 0x75, 0x67, 0x67, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x6b, 0x0a, 0x17, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x43,
	0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x38, 0x0a, 0x0a,
	0x73, 0x75, 0x62, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x18, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x61, 0x74, 0x6c, 0x61, 0x73, 0x2e,
	0x53, 0x75, 0x62, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x73, 0x75, 0x62, 0x6d,
	0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x22, 0x7a,
	0x0a, 0x16, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x43, 0x68, 0x61, 0x6e,
	0x67, 0x65, 0x64, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x30, 0x0a, 0x06, 0x62, 0x65, 0x66, 0x6f,
	0x72, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d,
	0x70, 0x2e, 0x61, 0x74, 0x6c, 0x61, 0x73, 0x2e, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x52, 0x06, 0x62, 0x65, 0x66, 0x6f, 0x72, 0x65, 0x12, 0x2e, 0x0a, 0x05, 0x61, 0x66,
	0x74, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x65, 0x6f, 0x6c, 0x79,
	0x6d, 0x70, 0x2e, 0x61, 0x74, 0x6c, 0x61, 0x73, 0x2e, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x52, 0x05, 0x61, 0x66, 0x74, 0x65, 0x72, 0x22, 0x3e, 0x0a, 0x11, 0x53, 0x63,
	0x6f, 0x72, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12,
	0x29, 0x0a, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13,
	0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x61, 0x74, 0x6c, 0x61, 0x73, 0x2e, 0x53, 0x63,
	0x6f, 0x72, 0x65, 0x52, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x22, 0x71, 0x0a, 0x13, 0x50, 0x72,
	0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x64, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x12, 0x2d, 0x0a, 0x06, 0x62, 0x65, 0x66, 0x6f, 0x72, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x15, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x61, 0x74, 0x6c, 0x61, 0x73,
	0x2e, 0x50, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x52, 0x06, 0x62, 0x65, 0x66, 0x6f, 0x72, 0x65,
	0x12, 0x2b, 0x0a, 0x05, 0x61, 0x66, 0x74, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x15, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x61, 0x74, 0x6c, 0x61, 0x73, 0x2e, 0x50,
	0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x52, 0x05, 0x61, 0x66, 0x74, 0x65, 0x72, 0x22, 0x96, 0x01,
	0x0a, 0x15, 0x53, 0x74, 0x61, 0x74, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x43, 0x68, 0x61, 0x6e, 0x67,
	0x65, 0x64, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x62, 0x6c,
	0x65, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f,
	0x62, 0x6c, 0x65, 0x6d, 0x49, 0x64, 0x12, 0x2f, 0x0a, 0x06, 0x62, 0x65, 0x66, 0x6f, 0x72, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e,
	0x61, 0x74, 0x6c, 0x61, 0x73, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x52,
	0x06, 0x62, 0x65, 0x66, 0x6f, 0x72, 0x65, 0x12, 0x2d, 0x0a, 0x05, 0x61, 0x66, 0x74, 0x65, 0x72,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e,
	0x61, 0x74, 0x6c, 0x61, 0x73, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x52,
	0x05, 0x61, 0x66, 0x74, 0x65, 0x72, 0x22, 0x80, 0x01, 0x0a, 0x14, 0x42, 0x6f, 0x6f, 0x6b, 0x6d,
	0x61, 0x72, 0x6b, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x64, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12,
	0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x49, 0x64, 0x12, 0x1b,
	0x0a, 0x09, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x62,
	0x65, 0x66, 0x6f, 0x72, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x62, 0x65, 0x66,
	0x6f, 0x72, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x66, 0x74, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x05, 0x61, 0x66, 0x74, 0x65, 0x72, 0x22, 0xae, 0x01, 0x0a, 0x16, 0x50, 0x65,
	0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x64, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x62, 0x6c, 0x65,
	0x6d, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x2e, 0x0a, 0x06,
	0x62, 0x65, 0x66, 0x6f, 0x72, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x65,
	0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x61, 0x63, 0x6c, 0x2e, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x52, 0x06, 0x62, 0x65, 0x66, 0x6f, 0x72, 0x65, 0x12, 0x2c, 0x0a, 0x05,
	0x61, 0x66, 0x74, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x65, 0x6f,
	0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x61, 0x63, 0x6c, 0x2e, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x52, 0x05, 0x61, 0x66, 0x74, 0x65, 0x72, 0x22, 0x99, 0x01, 0x0a, 0x16, 0x53,
	0x75, 0x67, 0x67, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x64,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x62, 0x6c,
	0x65, 0x6d, 0x49, 0x64, 0x12, 0x30, 0x0a, 0x06, 0x62, 0x65, 0x66, 0x6f, 0x72, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x61, 0x74,
	0x6c, 0x61, 0x73, 0x2e, 0x53, 0x75, 0x67, 0x67, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x06,
	0x62, 0x65, 0x66, 0x6f, 0x72, 0x65, 0x12, 0x2e, 0x0a, 0x05, 0x61, 0x66, 0x74, 0x65, 0x72, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x61,
	0x74, 0x6c, 0x61, 0x73, 0x2e, 0x53, 0x75, 0x67, 0x67, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x05, 0x61, 0x66, 0x74, 0x65, 0x72, 0x42, 0x2d, 0x5a, 0x2b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x67, 0x6f, 0x2d, 0x73,
	0x64, 0x6b, 0x2f, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x61, 0x74, 0x6c, 0x61, 0x73, 0x3b,
	0x61, 0x74, 0x6c, 0x61, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_eolymp_atlas_events_proto_rawDescOnce sync.Once
	file_eolymp_atlas_events_proto_rawDescData = file_eolymp_atlas_events_proto_rawDesc
)

func file_eolymp_atlas_events_proto_rawDescGZIP() []byte {
	file_eolymp_atlas_events_proto_rawDescOnce.Do(func() {
		file_eolymp_atlas_events_proto_rawDescData = protoimpl.X.CompressGZIP(file_eolymp_atlas_events_proto_rawDescData)
	})
	return file_eolymp_atlas_events_proto_rawDescData
}

var file_eolymp_atlas_events_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_eolymp_atlas_events_proto_goTypes = []any{
	(*SubmissionCompleteEvent)(nil), // 0: eolymp.atlas.SubmissionCompleteEvent
	(*SubmissionChangedEvent)(nil),  // 1: eolymp.atlas.SubmissionChangedEvent
	(*ScoreUpdatedEvent)(nil),       // 2: eolymp.atlas.ScoreUpdatedEvent
	(*ProblemChangedEvent)(nil),     // 3: eolymp.atlas.ProblemChangedEvent
	(*StatementChangedEvent)(nil),   // 4: eolymp.atlas.StatementChangedEvent
	(*BookmarkChangedEvent)(nil),    // 5: eolymp.atlas.BookmarkChangedEvent
	(*PermissionChangedEvent)(nil),  // 6: eolymp.atlas.PermissionChangedEvent
	(*SuggestionChangedEvent)(nil),  // 7: eolymp.atlas.SuggestionChangedEvent
	(*Submission)(nil),              // 8: eolymp.atlas.Submission
	(*Score)(nil),                   // 9: eolymp.atlas.Score
	(*Problem)(nil),                 // 10: eolymp.atlas.Problem
	(*Statement)(nil),               // 11: eolymp.atlas.Statement
	(*acl.Permission)(nil),          // 12: eolymp.acl.Permission
	(*Suggestion)(nil),              // 13: eolymp.atlas.Suggestion
}
var file_eolymp_atlas_events_proto_depIdxs = []int32{
	8,  // 0: eolymp.atlas.SubmissionCompleteEvent.submission:type_name -> eolymp.atlas.Submission
	8,  // 1: eolymp.atlas.SubmissionChangedEvent.before:type_name -> eolymp.atlas.Submission
	8,  // 2: eolymp.atlas.SubmissionChangedEvent.after:type_name -> eolymp.atlas.Submission
	9,  // 3: eolymp.atlas.ScoreUpdatedEvent.score:type_name -> eolymp.atlas.Score
	10, // 4: eolymp.atlas.ProblemChangedEvent.before:type_name -> eolymp.atlas.Problem
	10, // 5: eolymp.atlas.ProblemChangedEvent.after:type_name -> eolymp.atlas.Problem
	11, // 6: eolymp.atlas.StatementChangedEvent.before:type_name -> eolymp.atlas.Statement
	11, // 7: eolymp.atlas.StatementChangedEvent.after:type_name -> eolymp.atlas.Statement
	12, // 8: eolymp.atlas.PermissionChangedEvent.before:type_name -> eolymp.acl.Permission
	12, // 9: eolymp.atlas.PermissionChangedEvent.after:type_name -> eolymp.acl.Permission
	13, // 10: eolymp.atlas.SuggestionChangedEvent.before:type_name -> eolymp.atlas.Suggestion
	13, // 11: eolymp.atlas.SuggestionChangedEvent.after:type_name -> eolymp.atlas.Suggestion
	12, // [12:12] is the sub-list for method output_type
	12, // [12:12] is the sub-list for method input_type
	12, // [12:12] is the sub-list for extension type_name
	12, // [12:12] is the sub-list for extension extendee
	0,  // [0:12] is the sub-list for field type_name
}

func init() { file_eolymp_atlas_events_proto_init() }
func file_eolymp_atlas_events_proto_init() {
	if File_eolymp_atlas_events_proto != nil {
		return
	}
	file_eolymp_atlas_problem_proto_init()
	file_eolymp_atlas_scoring_score_proto_init()
	file_eolymp_atlas_statement_proto_init()
	file_eolymp_atlas_submission_proto_init()
	file_eolymp_atlas_suggestion_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_eolymp_atlas_events_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_eolymp_atlas_events_proto_goTypes,
		DependencyIndexes: file_eolymp_atlas_events_proto_depIdxs,
		MessageInfos:      file_eolymp_atlas_events_proto_msgTypes,
	}.Build()
	File_eolymp_atlas_events_proto = out.File
	file_eolymp_atlas_events_proto_rawDesc = nil
	file_eolymp_atlas_events_proto_goTypes = nil
	file_eolymp_atlas_events_proto_depIdxs = nil
}
