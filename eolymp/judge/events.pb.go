// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v3.21.12
// source: eolymp/judge/events.proto

package judge

import (
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

type SubmissionCompletedEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ContestId  string      `protobuf:"bytes,10,opt,name=contest_id,json=contestId,proto3" json:"contest_id,omitempty"`
	Submission *Submission `protobuf:"bytes,1,opt,name=submission,proto3" json:"submission,omitempty"`
}

func (x *SubmissionCompletedEvent) Reset() {
	*x = SubmissionCompletedEvent{}
	mi := &file_eolymp_judge_events_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SubmissionCompletedEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubmissionCompletedEvent) ProtoMessage() {}

func (x *SubmissionCompletedEvent) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_judge_events_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubmissionCompletedEvent.ProtoReflect.Descriptor instead.
func (*SubmissionCompletedEvent) Descriptor() ([]byte, []int) {
	return file_eolymp_judge_events_proto_rawDescGZIP(), []int{0}
}

func (x *SubmissionCompletedEvent) GetContestId() string {
	if x != nil {
		return x.ContestId
	}
	return ""
}

func (x *SubmissionCompletedEvent) GetSubmission() *Submission {
	if x != nil {
		return x.Submission
	}
	return nil
}

type RebuildScoreEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ContestId  string `protobuf:"bytes,1,opt,name=contest_id,json=contestId,proto3" json:"contest_id,omitempty"`
	ActivityId string `protobuf:"bytes,2,opt,name=activity_id,json=activityId,proto3" json:"activity_id,omitempty"`
}

func (x *RebuildScoreEvent) Reset() {
	*x = RebuildScoreEvent{}
	mi := &file_eolymp_judge_events_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RebuildScoreEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RebuildScoreEvent) ProtoMessage() {}

func (x *RebuildScoreEvent) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_judge_events_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RebuildScoreEvent.ProtoReflect.Descriptor instead.
func (*RebuildScoreEvent) Descriptor() ([]byte, []int) {
	return file_eolymp_judge_events_proto_rawDescGZIP(), []int{1}
}

func (x *RebuildScoreEvent) GetContestId() string {
	if x != nil {
		return x.ContestId
	}
	return ""
}

func (x *RebuildScoreEvent) GetActivityId() string {
	if x != nil {
		return x.ActivityId
	}
	return ""
}

type ScoreChangedEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ContestId     string `protobuf:"bytes,1,opt,name=contest_id,json=contestId,proto3" json:"contest_id,omitempty"`
	ParticipantId string `protobuf:"bytes,2,opt,name=participant_id,json=participantId,proto3" json:"participant_id,omitempty"`
	Unofficial    bool   `protobuf:"varint,4,opt,name=unofficial,proto3" json:"unofficial,omitempty"`
	Score         *Score `protobuf:"bytes,3,opt,name=score,proto3" json:"score,omitempty"`
}

func (x *ScoreChangedEvent) Reset() {
	*x = ScoreChangedEvent{}
	mi := &file_eolymp_judge_events_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ScoreChangedEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ScoreChangedEvent) ProtoMessage() {}

func (x *ScoreChangedEvent) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_judge_events_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ScoreChangedEvent.ProtoReflect.Descriptor instead.
func (*ScoreChangedEvent) Descriptor() ([]byte, []int) {
	return file_eolymp_judge_events_proto_rawDescGZIP(), []int{2}
}

func (x *ScoreChangedEvent) GetContestId() string {
	if x != nil {
		return x.ContestId
	}
	return ""
}

func (x *ScoreChangedEvent) GetParticipantId() string {
	if x != nil {
		return x.ParticipantId
	}
	return ""
}

func (x *ScoreChangedEvent) GetUnofficial() bool {
	if x != nil {
		return x.Unofficial
	}
	return false
}

func (x *ScoreChangedEvent) GetScore() *Score {
	if x != nil {
		return x.Score
	}
	return nil
}

type ParticipantChangedEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ContestId string       `protobuf:"bytes,10,opt,name=contest_id,json=contestId,proto3" json:"contest_id,omitempty"`
	Before    *Participant `protobuf:"bytes,1,opt,name=before,proto3" json:"before,omitempty"`
	After     *Participant `protobuf:"bytes,2,opt,name=after,proto3" json:"after,omitempty"`
}

func (x *ParticipantChangedEvent) Reset() {
	*x = ParticipantChangedEvent{}
	mi := &file_eolymp_judge_events_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ParticipantChangedEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ParticipantChangedEvent) ProtoMessage() {}

func (x *ParticipantChangedEvent) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_judge_events_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ParticipantChangedEvent.ProtoReflect.Descriptor instead.
func (*ParticipantChangedEvent) Descriptor() ([]byte, []int) {
	return file_eolymp_judge_events_proto_rawDescGZIP(), []int{3}
}

func (x *ParticipantChangedEvent) GetContestId() string {
	if x != nil {
		return x.ContestId
	}
	return ""
}

func (x *ParticipantChangedEvent) GetBefore() *Participant {
	if x != nil {
		return x.Before
	}
	return nil
}

func (x *ParticipantChangedEvent) GetAfter() *Participant {
	if x != nil {
		return x.After
	}
	return nil
}

type ParticipantJoinedEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ContestId   string       `protobuf:"bytes,10,opt,name=contest_id,json=contestId,proto3" json:"contest_id,omitempty"`
	Participant *Participant `protobuf:"bytes,2,opt,name=participant,proto3" json:"participant,omitempty"`
}

func (x *ParticipantJoinedEvent) Reset() {
	*x = ParticipantJoinedEvent{}
	mi := &file_eolymp_judge_events_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ParticipantJoinedEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ParticipantJoinedEvent) ProtoMessage() {}

func (x *ParticipantJoinedEvent) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_judge_events_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ParticipantJoinedEvent.ProtoReflect.Descriptor instead.
func (*ParticipantJoinedEvent) Descriptor() ([]byte, []int) {
	return file_eolymp_judge_events_proto_rawDescGZIP(), []int{4}
}

func (x *ParticipantJoinedEvent) GetContestId() string {
	if x != nil {
		return x.ContestId
	}
	return ""
}

func (x *ParticipantJoinedEvent) GetParticipant() *Participant {
	if x != nil {
		return x.Participant
	}
	return nil
}

type RetestProblemEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ContestId  string `protobuf:"bytes,1,opt,name=contest_id,json=contestId,proto3" json:"contest_id,omitempty"`
	ProblemId  string `protobuf:"bytes,2,opt,name=problem_id,json=problemId,proto3" json:"problem_id,omitempty"`
	ActivityId string `protobuf:"bytes,3,opt,name=activity_id,json=activityId,proto3" json:"activity_id,omitempty"`
}

func (x *RetestProblemEvent) Reset() {
	*x = RetestProblemEvent{}
	mi := &file_eolymp_judge_events_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RetestProblemEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RetestProblemEvent) ProtoMessage() {}

func (x *RetestProblemEvent) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_judge_events_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RetestProblemEvent.ProtoReflect.Descriptor instead.
func (*RetestProblemEvent) Descriptor() ([]byte, []int) {
	return file_eolymp_judge_events_proto_rawDescGZIP(), []int{5}
}

func (x *RetestProblemEvent) GetContestId() string {
	if x != nil {
		return x.ContestId
	}
	return ""
}

func (x *RetestProblemEvent) GetProblemId() string {
	if x != nil {
		return x.ProblemId
	}
	return ""
}

func (x *RetestProblemEvent) GetActivityId() string {
	if x != nil {
		return x.ActivityId
	}
	return ""
}

var File_eolymp_judge_events_proto protoreflect.FileDescriptor

var file_eolymp_judge_events_proto_rawDesc = []byte{
	0x0a, 0x19, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x6a, 0x75, 0x64, 0x67, 0x65, 0x2f, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x65, 0x6f, 0x6c,
	0x79, 0x6d, 0x70, 0x2e, 0x6a, 0x75, 0x64, 0x67, 0x65, 0x1a, 0x1e, 0x65, 0x6f, 0x6c, 0x79, 0x6d,
	0x70, 0x2f, 0x6a, 0x75, 0x64, 0x67, 0x65, 0x2f, 0x70, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70,
	0x61, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x18, 0x65, 0x6f, 0x6c, 0x79, 0x6d,
	0x70, 0x2f, 0x6a, 0x75, 0x64, 0x67, 0x65, 0x2f, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x6a, 0x75, 0x64, 0x67,
	0x65, 0x2f, 0x73, 0x75, 0x62, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x73, 0x0a, 0x18, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x1d,
	0x0a, 0x0a, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x73, 0x74, 0x49, 0x64, 0x12, 0x38, 0x0a,
	0x0a, 0x73, 0x75, 0x62, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x18, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x6a, 0x75, 0x64, 0x67, 0x65,
	0x2e, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x73, 0x75, 0x62,
	0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x53, 0x0a, 0x11, 0x52, 0x65, 0x62, 0x75, 0x69,
	0x6c, 0x64, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x1d, 0x0a, 0x0a,
	0x63, 0x6f, 0x6e, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x73, 0x74, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x61,
	0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x49, 0x64, 0x22, 0xa4, 0x01, 0x0a,
	0x11, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x64, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x73, 0x74, 0x49,
	0x64, 0x12, 0x25, 0x0a, 0x0e, 0x70, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x70, 0x61, 0x72, 0x74, 0x69,
	0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x75, 0x6e, 0x6f, 0x66,
	0x66, 0x69, 0x63, 0x69, 0x61, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x75, 0x6e,
	0x6f, 0x66, 0x66, 0x69, 0x63, 0x69, 0x61, 0x6c, 0x12, 0x29, 0x0a, 0x05, 0x73, 0x63, 0x6f, 0x72,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70,
	0x2e, 0x6a, 0x75, 0x64, 0x67, 0x65, 0x2e, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x52, 0x05, 0x73, 0x63,
	0x6f, 0x72, 0x65, 0x22, 0x9c, 0x01, 0x0a, 0x17, 0x50, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70,
	0x61, 0x6e, 0x74, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x64, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12,
	0x1d, 0x0a, 0x0a, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x73, 0x74, 0x49, 0x64, 0x12, 0x31,
	0x0a, 0x06, 0x62, 0x65, 0x66, 0x6f, 0x72, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19,
	0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x6a, 0x75, 0x64, 0x67, 0x65, 0x2e, 0x50, 0x61,
	0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x52, 0x06, 0x62, 0x65, 0x66, 0x6f, 0x72,
	0x65, 0x12, 0x2f, 0x0a, 0x05, 0x61, 0x66, 0x74, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x19, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x6a, 0x75, 0x64, 0x67, 0x65, 0x2e,
	0x50, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x52, 0x05, 0x61, 0x66, 0x74,
	0x65, 0x72, 0x22, 0x74, 0x0a, 0x16, 0x50, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e,
	0x74, 0x4a, 0x6f, 0x69, 0x6e, 0x65, 0x64, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x1d, 0x0a, 0x0a,
	0x63, 0x6f, 0x6e, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x73, 0x74, 0x49, 0x64, 0x12, 0x3b, 0x0a, 0x0b, 0x70,
	0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x19, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x6a, 0x75, 0x64, 0x67, 0x65, 0x2e,
	0x50, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x52, 0x0b, 0x70, 0x61, 0x72,
	0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x22, 0x73, 0x0a, 0x12, 0x52, 0x65, 0x74, 0x65,
	0x73, 0x74, 0x50, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x1d,
	0x0a, 0x0a, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x73, 0x74, 0x49, 0x64, 0x12, 0x1d, 0x0a,
	0x0a, 0x70, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b,
	0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x49, 0x64, 0x42, 0x2d, 0x5a,
	0x2b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6f, 0x6c, 0x79,
	0x6d, 0x70, 0x2f, 0x67, 0x6f, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70,
	0x2f, 0x6a, 0x75, 0x64, 0x67, 0x65, 0x3b, 0x6a, 0x75, 0x64, 0x67, 0x65, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_eolymp_judge_events_proto_rawDescOnce sync.Once
	file_eolymp_judge_events_proto_rawDescData = file_eolymp_judge_events_proto_rawDesc
)

func file_eolymp_judge_events_proto_rawDescGZIP() []byte {
	file_eolymp_judge_events_proto_rawDescOnce.Do(func() {
		file_eolymp_judge_events_proto_rawDescData = protoimpl.X.CompressGZIP(file_eolymp_judge_events_proto_rawDescData)
	})
	return file_eolymp_judge_events_proto_rawDescData
}

var file_eolymp_judge_events_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_eolymp_judge_events_proto_goTypes = []any{
	(*SubmissionCompletedEvent)(nil), // 0: eolymp.judge.SubmissionCompletedEvent
	(*RebuildScoreEvent)(nil),        // 1: eolymp.judge.RebuildScoreEvent
	(*ScoreChangedEvent)(nil),        // 2: eolymp.judge.ScoreChangedEvent
	(*ParticipantChangedEvent)(nil),  // 3: eolymp.judge.ParticipantChangedEvent
	(*ParticipantJoinedEvent)(nil),   // 4: eolymp.judge.ParticipantJoinedEvent
	(*RetestProblemEvent)(nil),       // 5: eolymp.judge.RetestProblemEvent
	(*Submission)(nil),               // 6: eolymp.judge.Submission
	(*Score)(nil),                    // 7: eolymp.judge.Score
	(*Participant)(nil),              // 8: eolymp.judge.Participant
}
var file_eolymp_judge_events_proto_depIdxs = []int32{
	6, // 0: eolymp.judge.SubmissionCompletedEvent.submission:type_name -> eolymp.judge.Submission
	7, // 1: eolymp.judge.ScoreChangedEvent.score:type_name -> eolymp.judge.Score
	8, // 2: eolymp.judge.ParticipantChangedEvent.before:type_name -> eolymp.judge.Participant
	8, // 3: eolymp.judge.ParticipantChangedEvent.after:type_name -> eolymp.judge.Participant
	8, // 4: eolymp.judge.ParticipantJoinedEvent.participant:type_name -> eolymp.judge.Participant
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_eolymp_judge_events_proto_init() }
func file_eolymp_judge_events_proto_init() {
	if File_eolymp_judge_events_proto != nil {
		return
	}
	file_eolymp_judge_participant_proto_init()
	file_eolymp_judge_score_proto_init()
	file_eolymp_judge_submission_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_eolymp_judge_events_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_eolymp_judge_events_proto_goTypes,
		DependencyIndexes: file_eolymp_judge_events_proto_depIdxs,
		MessageInfos:      file_eolymp_judge_events_proto_msgTypes,
	}.Build()
	File_eolymp_judge_events_proto = out.File
	file_eolymp_judge_events_proto_rawDesc = nil
	file_eolymp_judge_events_proto_goTypes = nil
	file_eolymp_judge_events_proto_depIdxs = nil
}
