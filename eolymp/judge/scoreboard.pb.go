// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.12.2
// source: eolymp/judge/scoreboard.proto

package judge

import (
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type Scoreboard_Status int32

const (
	Scoreboard_UNKNOWN    Scoreboard_Status = 0 // Scoreboard is up-to-date and being updated.
	Scoreboard_ACTIVE     Scoreboard_Status = 1 // Scoreboard is up-to-date and being updated.
	Scoreboard_FROZEN     Scoreboard_Status = 2 // Scoreboard is frozen and not updated anymore.
	Scoreboard_REBUILDING Scoreboard_Status = 3 // Scoreboard is being rebuild (results are recalculated).
	Scoreboard_FAILED     Scoreboard_Status = 4 // Scoreboard rebuild process has failed (needs to be restarted).
)

// Enum value maps for Scoreboard_Status.
var (
	Scoreboard_Status_name = map[int32]string{
		0: "UNKNOWN",
		1: "ACTIVE",
		2: "FROZEN",
		3: "REBUILDING",
		4: "FAILED",
	}
	Scoreboard_Status_value = map[string]int32{
		"UNKNOWN":    0,
		"ACTIVE":     1,
		"FROZEN":     2,
		"REBUILDING": 3,
		"FAILED":     4,
	}
)

func (x Scoreboard_Status) Enum() *Scoreboard_Status {
	p := new(Scoreboard_Status)
	*p = x
	return p
}

func (x Scoreboard_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Scoreboard_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_eolymp_judge_scoreboard_proto_enumTypes[0].Descriptor()
}

func (Scoreboard_Status) Type() protoreflect.EnumType {
	return &file_eolymp_judge_scoreboard_proto_enumTypes[0]
}

func (x Scoreboard_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Scoreboard_Status.Descriptor instead.
func (Scoreboard_Status) EnumDescriptor() ([]byte, []int) {
	return file_eolymp_judge_scoreboard_proto_rawDescGZIP(), []int{0, 0}
}

type Scoreboard struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id             string               `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name           string               `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	Default        bool                 `protobuf:"varint,2,opt,name=default,proto3" json:"default,omitempty"`
	Visible        bool                 `protobuf:"varint,3,opt,name=visible,proto3" json:"visible,omitempty"`
	AttemptPenalty uint32               `protobuf:"varint,30,opt,name=attempt_penalty,json=attemptPenalty,proto3" json:"attempt_penalty,omitempty"` // Amount of penalty awarded for each submission
	Status         Scoreboard_Status    `protobuf:"varint,11,opt,name=status,proto3,enum=eolymp.judge.Scoreboard_Status" json:"status,omitempty"`
	FreezeTime     *timestamp.Timestamp `protobuf:"bytes,21,opt,name=freeze_time,json=freezeTime,proto3" json:"freeze_time,omitempty"`
	FreezeIn       uint32               `protobuf:"varint,22,opt,name=freeze_in,json=freezeIn,proto3" json:"freeze_in,omitempty"`
}

func (x *Scoreboard) Reset() {
	*x = Scoreboard{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_judge_scoreboard_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Scoreboard) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Scoreboard) ProtoMessage() {}

func (x *Scoreboard) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_judge_scoreboard_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Scoreboard.ProtoReflect.Descriptor instead.
func (*Scoreboard) Descriptor() ([]byte, []int) {
	return file_eolymp_judge_scoreboard_proto_rawDescGZIP(), []int{0}
}

func (x *Scoreboard) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Scoreboard) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Scoreboard) GetDefault() bool {
	if x != nil {
		return x.Default
	}
	return false
}

func (x *Scoreboard) GetVisible() bool {
	if x != nil {
		return x.Visible
	}
	return false
}

func (x *Scoreboard) GetAttemptPenalty() uint32 {
	if x != nil {
		return x.AttemptPenalty
	}
	return 0
}

func (x *Scoreboard) GetStatus() Scoreboard_Status {
	if x != nil {
		return x.Status
	}
	return Scoreboard_UNKNOWN
}

func (x *Scoreboard) GetFreezeTime() *timestamp.Timestamp {
	if x != nil {
		return x.FreezeTime
	}
	return nil
}

func (x *Scoreboard) GetFreezeIn() uint32 {
	if x != nil {
		return x.FreezeIn
	}
	return 0
}

type Scoreboard_Row struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ParticipantId string              `protobuf:"bytes,1,opt,name=participant_id,json=participantId,proto3" json:"participant_id,omitempty"`
	ScoreboardId  string              `protobuf:"bytes,2,opt,name=scoreboard_id,json=scoreboardId,proto3" json:"scoreboard_id,omitempty"`
	Score         float32             `protobuf:"fixed32,10,opt,name=score,proto3" json:"score,omitempty"`                         // total score
	Penalty       float32             `protobuf:"fixed32,11,opt,name=penalty,proto3" json:"penalty,omitempty"`                     // total penalty
	Rank          uint32              `protobuf:"varint,12,opt,name=rank,proto3" json:"rank,omitempty"`                            // rank
	RankLower     uint32              `protobuf:"varint,13,opt,name=rank_lower,json=rankLower,proto3" json:"rank_lower,omitempty"` // lower rank
	Breakdown     []*Scoreboard_Score `protobuf:"bytes,20,rep,name=breakdown,proto3" json:"breakdown,omitempty"`                   // breakdown of score by problem
}

func (x *Scoreboard_Row) Reset() {
	*x = Scoreboard_Row{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_judge_scoreboard_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Scoreboard_Row) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Scoreboard_Row) ProtoMessage() {}

func (x *Scoreboard_Row) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_judge_scoreboard_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Scoreboard_Row.ProtoReflect.Descriptor instead.
func (*Scoreboard_Row) Descriptor() ([]byte, []int) {
	return file_eolymp_judge_scoreboard_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Scoreboard_Row) GetParticipantId() string {
	if x != nil {
		return x.ParticipantId
	}
	return ""
}

func (x *Scoreboard_Row) GetScoreboardId() string {
	if x != nil {
		return x.ScoreboardId
	}
	return ""
}

func (x *Scoreboard_Row) GetScore() float32 {
	if x != nil {
		return x.Score
	}
	return 0
}

func (x *Scoreboard_Row) GetPenalty() float32 {
	if x != nil {
		return x.Penalty
	}
	return 0
}

func (x *Scoreboard_Row) GetRank() uint32 {
	if x != nil {
		return x.Rank
	}
	return 0
}

func (x *Scoreboard_Row) GetRankLower() uint32 {
	if x != nil {
		return x.RankLower
	}
	return 0
}

func (x *Scoreboard_Row) GetBreakdown() []*Scoreboard_Score {
	if x != nil {
		return x.Breakdown
	}
	return nil
}

type Scoreboard_Score struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProblemId  string               `protobuf:"bytes,1,opt,name=problem_id,json=problemId,proto3" json:"problem_id,omitempty"`
	Score      float32              `protobuf:"fixed32,2,opt,name=score,proto3" json:"score,omitempty"` // number of points scored
	Penalty    float32              `protobuf:"fixed32,3,opt,name=penalty,proto3" json:"penalty,omitempty"`
	Solved     bool                 `protobuf:"varint,4,opt,name=solved,proto3" json:"solved,omitempty"`
	Percentage float32              `protobuf:"fixed32,5,opt,name=percentage,proto3" json:"percentage,omitempty"`            // percentage scored from 0 to 1
	Attempts   uint32               `protobuf:"varint,7,opt,name=attempts,proto3" json:"attempts,omitempty"`                 // number of submits made until problem was solved
	SolvedAt   *timestamp.Timestamp `protobuf:"bytes,8,opt,name=solved_at,json=solvedAt,proto3" json:"solved_at,omitempty"`  // time when first accepted submission is made
	SolvedIn   uint32               `protobuf:"varint,9,opt,name=solved_in,json=solvedIn,proto3" json:"solved_in,omitempty"` // amount of time it took user to get accepted
}

func (x *Scoreboard_Score) Reset() {
	*x = Scoreboard_Score{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_judge_scoreboard_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Scoreboard_Score) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Scoreboard_Score) ProtoMessage() {}

func (x *Scoreboard_Score) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_judge_scoreboard_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Scoreboard_Score.ProtoReflect.Descriptor instead.
func (*Scoreboard_Score) Descriptor() ([]byte, []int) {
	return file_eolymp_judge_scoreboard_proto_rawDescGZIP(), []int{0, 1}
}

func (x *Scoreboard_Score) GetProblemId() string {
	if x != nil {
		return x.ProblemId
	}
	return ""
}

func (x *Scoreboard_Score) GetScore() float32 {
	if x != nil {
		return x.Score
	}
	return 0
}

func (x *Scoreboard_Score) GetPenalty() float32 {
	if x != nil {
		return x.Penalty
	}
	return 0
}

func (x *Scoreboard_Score) GetSolved() bool {
	if x != nil {
		return x.Solved
	}
	return false
}

func (x *Scoreboard_Score) GetPercentage() float32 {
	if x != nil {
		return x.Percentage
	}
	return 0
}

func (x *Scoreboard_Score) GetAttempts() uint32 {
	if x != nil {
		return x.Attempts
	}
	return 0
}

func (x *Scoreboard_Score) GetSolvedAt() *timestamp.Timestamp {
	if x != nil {
		return x.SolvedAt
	}
	return nil
}

func (x *Scoreboard_Score) GetSolvedIn() uint32 {
	if x != nil {
		return x.SolvedIn
	}
	return 0
}

var File_eolymp_judge_scoreboard_proto protoreflect.FileDescriptor

var file_eolymp_judge_scoreboard_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x6a, 0x75, 0x64, 0x67, 0x65, 0x2f, 0x73,
	0x63, 0x6f, 0x72, 0x65, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0c, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x6a, 0x75, 0x64, 0x67, 0x65, 0x1a, 0x1f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe3,
	0x06, 0x0a, 0x0a, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x07, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x76,
	0x69, 0x73, 0x69, 0x62, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x76, 0x69,
	0x73, 0x69, 0x62, 0x6c, 0x65, 0x12, 0x27, 0x0a, 0x0f, 0x61, 0x74, 0x74, 0x65, 0x6d, 0x70, 0x74,
	0x5f, 0x70, 0x65, 0x6e, 0x61, 0x6c, 0x74, 0x79, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0e,
	0x61, 0x74, 0x74, 0x65, 0x6d, 0x70, 0x74, 0x50, 0x65, 0x6e, 0x61, 0x6c, 0x74, 0x79, 0x12, 0x37,
	0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1f,
	0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x6a, 0x75, 0x64, 0x67, 0x65, 0x2e, 0x53, 0x63,
	0x6f, 0x72, 0x65, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x3b, 0x0a, 0x0b, 0x66, 0x72, 0x65, 0x65, 0x7a,
	0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x15, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x66, 0x72, 0x65, 0x65, 0x7a, 0x65,
	0x54, 0x69, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x72, 0x65, 0x65, 0x7a, 0x65, 0x5f, 0x69,
	0x6e, 0x18, 0x16, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x66, 0x72, 0x65, 0x65, 0x7a, 0x65, 0x49,
	0x6e, 0x1a, 0xf2, 0x01, 0x0a, 0x03, 0x52, 0x6f, 0x77, 0x12, 0x25, 0x0a, 0x0e, 0x70, 0x61, 0x72,
	0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0d, 0x70, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x49, 0x64,
	0x12, 0x23, 0x0a, 0x0d, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x5f, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x62, 0x6f,
	0x61, 0x72, 0x64, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x0a,
	0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x70,
	0x65, 0x6e, 0x61, 0x6c, 0x74, 0x79, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x02, 0x52, 0x07, 0x70, 0x65,
	0x6e, 0x61, 0x6c, 0x74, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x61, 0x6e, 0x6b, 0x18, 0x0c, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x04, 0x72, 0x61, 0x6e, 0x6b, 0x12, 0x1d, 0x0a, 0x0a, 0x72, 0x61, 0x6e,
	0x6b, 0x5f, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x72,
	0x61, 0x6e, 0x6b, 0x4c, 0x6f, 0x77, 0x65, 0x72, 0x12, 0x3c, 0x0a, 0x09, 0x62, 0x72, 0x65, 0x61,
	0x6b, 0x64, 0x6f, 0x77, 0x6e, 0x18, 0x14, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x65, 0x6f,
	0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x6a, 0x75, 0x64, 0x67, 0x65, 0x2e, 0x53, 0x63, 0x6f, 0x72, 0x65,
	0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x52, 0x09, 0x62, 0x72, 0x65,
	0x61, 0x6b, 0x64, 0x6f, 0x77, 0x6e, 0x1a, 0x80, 0x02, 0x0a, 0x05, 0x53, 0x63, 0x6f, 0x72, 0x65,
	0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x49, 0x64, 0x12,
	0x14, 0x0a, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05,
	0x73, 0x63, 0x6f, 0x72, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x65, 0x6e, 0x61, 0x6c, 0x74, 0x79,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x07, 0x70, 0x65, 0x6e, 0x61, 0x6c, 0x74, 0x79, 0x12,
	0x16, 0x0a, 0x06, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x06, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x65, 0x72, 0x63, 0x65,
	0x6e, 0x74, 0x61, 0x67, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0a, 0x70, 0x65, 0x72,
	0x63, 0x65, 0x6e, 0x74, 0x61, 0x67, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x61, 0x74, 0x74, 0x65, 0x6d,
	0x70, 0x74, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x61, 0x74, 0x74, 0x65, 0x6d,
	0x70, 0x74, 0x73, 0x12, 0x37, 0x0a, 0x09, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x64, 0x5f, 0x61, 0x74,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x08, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1b, 0x0a, 0x09,
	0x73, 0x6f, 0x6c, 0x76, 0x65, 0x64, 0x5f, 0x69, 0x6e, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x08, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x64, 0x49, 0x6e, 0x22, 0x49, 0x0a, 0x06, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00,
	0x12, 0x0a, 0x0a, 0x06, 0x41, 0x43, 0x54, 0x49, 0x56, 0x45, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06,
	0x46, 0x52, 0x4f, 0x5a, 0x45, 0x4e, 0x10, 0x02, 0x12, 0x0e, 0x0a, 0x0a, 0x52, 0x45, 0x42, 0x55,
	0x49, 0x4c, 0x44, 0x49, 0x4e, 0x47, 0x10, 0x03, 0x12, 0x0a, 0x0a, 0x06, 0x46, 0x41, 0x49, 0x4c,
	0x45, 0x44, 0x10, 0x04, 0x42, 0x33, 0x5a, 0x31, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61,
	0x63, 0x74, 0x73, 0x2f, 0x67, 0x6f, 0x2f, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x6a, 0x75,
	0x64, 0x67, 0x65, 0x3b, 0x6a, 0x75, 0x64, 0x67, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_eolymp_judge_scoreboard_proto_rawDescOnce sync.Once
	file_eolymp_judge_scoreboard_proto_rawDescData = file_eolymp_judge_scoreboard_proto_rawDesc
)

func file_eolymp_judge_scoreboard_proto_rawDescGZIP() []byte {
	file_eolymp_judge_scoreboard_proto_rawDescOnce.Do(func() {
		file_eolymp_judge_scoreboard_proto_rawDescData = protoimpl.X.CompressGZIP(file_eolymp_judge_scoreboard_proto_rawDescData)
	})
	return file_eolymp_judge_scoreboard_proto_rawDescData
}

var file_eolymp_judge_scoreboard_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_eolymp_judge_scoreboard_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_eolymp_judge_scoreboard_proto_goTypes = []interface{}{
	(Scoreboard_Status)(0),      // 0: eolymp.judge.Scoreboard.Status
	(*Scoreboard)(nil),          // 1: eolymp.judge.Scoreboard
	(*Scoreboard_Row)(nil),      // 2: eolymp.judge.Scoreboard.Row
	(*Scoreboard_Score)(nil),    // 3: eolymp.judge.Scoreboard.Score
	(*timestamp.Timestamp)(nil), // 4: google.protobuf.Timestamp
}
var file_eolymp_judge_scoreboard_proto_depIdxs = []int32{
	0, // 0: eolymp.judge.Scoreboard.status:type_name -> eolymp.judge.Scoreboard.Status
	4, // 1: eolymp.judge.Scoreboard.freeze_time:type_name -> google.protobuf.Timestamp
	3, // 2: eolymp.judge.Scoreboard.Row.breakdown:type_name -> eolymp.judge.Scoreboard.Score
	4, // 3: eolymp.judge.Scoreboard.Score.solved_at:type_name -> google.protobuf.Timestamp
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_eolymp_judge_scoreboard_proto_init() }
func file_eolymp_judge_scoreboard_proto_init() {
	if File_eolymp_judge_scoreboard_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_eolymp_judge_scoreboard_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Scoreboard); i {
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
		file_eolymp_judge_scoreboard_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Scoreboard_Row); i {
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
		file_eolymp_judge_scoreboard_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Scoreboard_Score); i {
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
			RawDescriptor: file_eolymp_judge_scoreboard_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_eolymp_judge_scoreboard_proto_goTypes,
		DependencyIndexes: file_eolymp_judge_scoreboard_proto_depIdxs,
		EnumInfos:         file_eolymp_judge_scoreboard_proto_enumTypes,
		MessageInfos:      file_eolymp_judge_scoreboard_proto_msgTypes,
	}.Build()
	File_eolymp_judge_scoreboard_proto = out.File
	file_eolymp_judge_scoreboard_proto_rawDesc = nil
	file_eolymp_judge_scoreboard_proto_goTypes = nil
	file_eolymp_judge_scoreboard_proto_depIdxs = nil
}
