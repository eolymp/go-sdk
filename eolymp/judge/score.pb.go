// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.0
// 	protoc        v3.21.12
// source: eolymp/judge/score.proto

package judge

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Score_FetchingMode int32

const (
	// Actual score returns score at the moment of participation. This mode is intended to show scoreboard to
	// participants. This mode returns score following these rules:
	//   - During scoreboard freezing time, frozen scores are reported.
	//   - In virtual contests (everyone can start at different time), score values are returned relatively to the
	//     starting time. For instance, authenticated user who has been participating for 1 hour, will receive scores at 1
	//     hour mark, as user progresses further, more score updates will be revealed.
	//   - If authenticated user is not participating in the contest, an error will be returned (even if requested by a
	//     user with admin permissions)
	Score_ACTUAL Score_FetchingMode = 0
	// Punctual score returns score at particular moment, use time_offset parameter to specify time. This mode is
	// intended to show historical score at a given moment. Value for time_offset will be capped by the freezing time
	// for participants.
	Score_PUNCTUAL Score_FetchingMode = 1
	// Latest score returns the latest score recorded. This mode is intended for admins to see current scoreboard.
	// Users without admin permissions will get PermissionDenied error when requesting score in latest mode.
	Score_LATEST Score_FetchingMode = 2
	// Frozen score returns the latest score recorded before freezing time. This mode is intended for admins to see
	// frozen scoreboard.
	Score_FROZEN Score_FetchingMode = 3
	// Upsolve score returns score counting upsolve time.
	Score_UPSOLVE Score_FetchingMode = 4
)

// Enum value maps for Score_FetchingMode.
var (
	Score_FetchingMode_name = map[int32]string{
		0: "ACTUAL",
		1: "PUNCTUAL",
		2: "LATEST",
		3: "FROZEN",
		4: "UPSOLVE",
	}
	Score_FetchingMode_value = map[string]int32{
		"ACTUAL":   0,
		"PUNCTUAL": 1,
		"LATEST":   2,
		"FROZEN":   3,
		"UPSOLVE":  4,
	}
)

func (x Score_FetchingMode) Enum() *Score_FetchingMode {
	p := new(Score_FetchingMode)
	*p = x
	return p
}

func (x Score_FetchingMode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Score_FetchingMode) Descriptor() protoreflect.EnumDescriptor {
	return file_eolymp_judge_score_proto_enumTypes[0].Descriptor()
}

func (Score_FetchingMode) Type() protoreflect.EnumType {
	return &file_eolymp_judge_score_proto_enumTypes[0]
}

func (x Score_FetchingMode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Score_FetchingMode.Descriptor instead.
func (Score_FetchingMode) EnumDescriptor() ([]byte, []int) {
	return file_eolymp_judge_score_proto_rawDescGZIP(), []int{0, 0}
}

type Score struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ValidAfter uint32                 `protobuf:"varint,1,opt,name=valid_after,json=validAfter,proto3" json:"valid_after,omitempty"` // time when score was set, number of seconds since start of contest (participation)
	ValidUntil uint32                 `protobuf:"varint,2,opt,name=valid_until,json=validUntil,proto3" json:"valid_until,omitempty"` // time when score was overridden by newer value, number of seconds since start of contest (participation)
	Timestamp  *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`                      // absolute time when score was set
	Score      float32                `protobuf:"fixed32,10,opt,name=score,proto3" json:"score,omitempty"`
	Penalty    float32                `protobuf:"fixed32,11,opt,name=penalty,proto3" json:"penalty,omitempty"`
	Upsolve    bool                   `protobuf:"varint,12,opt,name=upsolve,proto3" json:"upsolve,omitempty"`
	Breakdown  []*Score_Problem       `protobuf:"bytes,20,rep,name=breakdown,proto3" json:"breakdown,omitempty"` // breakdown of total score by problem
}

func (x *Score) Reset() {
	*x = Score{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_judge_score_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Score) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Score) ProtoMessage() {}

func (x *Score) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_judge_score_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Score.ProtoReflect.Descriptor instead.
func (*Score) Descriptor() ([]byte, []int) {
	return file_eolymp_judge_score_proto_rawDescGZIP(), []int{0}
}

func (x *Score) GetValidAfter() uint32 {
	if x != nil {
		return x.ValidAfter
	}
	return 0
}

func (x *Score) GetValidUntil() uint32 {
	if x != nil {
		return x.ValidUntil
	}
	return 0
}

func (x *Score) GetTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

func (x *Score) GetScore() float32 {
	if x != nil {
		return x.Score
	}
	return 0
}

func (x *Score) GetPenalty() float32 {
	if x != nil {
		return x.Penalty
	}
	return 0
}

func (x *Score) GetUpsolve() bool {
	if x != nil {
		return x.Upsolve
	}
	return false
}

func (x *Score) GetBreakdown() []*Score_Problem {
	if x != nil {
		return x.Breakdown
	}
	return nil
}

type Score_Problem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProblemId  string                 `protobuf:"bytes,1,opt,name=problem_id,json=problemId,proto3" json:"problem_id,omitempty"`
	Score      float32                `protobuf:"fixed32,2,opt,name=score,proto3" json:"score,omitempty"` // number of points scored
	Penalty    float32                `protobuf:"fixed32,3,opt,name=penalty,proto3" json:"penalty,omitempty"`
	Solved     bool                   `protobuf:"varint,4,opt,name=solved,proto3" json:"solved,omitempty"`
	Percentage float32                `protobuf:"fixed32,5,opt,name=percentage,proto3" json:"percentage,omitempty"`            // percentage scored from 0 to 1
	Attempts   uint32                 `protobuf:"varint,7,opt,name=attempts,proto3" json:"attempts,omitempty"`                 // number of submits made until problem was solved
	SolvedAt   *timestamppb.Timestamp `protobuf:"bytes,8,opt,name=solved_at,json=solvedAt,proto3" json:"solved_at,omitempty"`  // time when first accepted submission is made
	SolvedIn   uint32                 `protobuf:"varint,9,opt,name=solved_in,json=solvedIn,proto3" json:"solved_in,omitempty"` // amount of time it took user to get accepted
	Changed    bool                   `protobuf:"varint,10,opt,name=changed,proto3" json:"changed,omitempty"`                  // if true, means there is a newer value for this score
	Breakdown  []*Score_Testset       `protobuf:"bytes,20,rep,name=breakdown,proto3" json:"breakdown,omitempty"`               // breakdown of problem score by testset
}

func (x *Score_Problem) Reset() {
	*x = Score_Problem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_judge_score_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Score_Problem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Score_Problem) ProtoMessage() {}

func (x *Score_Problem) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_judge_score_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Score_Problem.ProtoReflect.Descriptor instead.
func (*Score_Problem) Descriptor() ([]byte, []int) {
	return file_eolymp_judge_score_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Score_Problem) GetProblemId() string {
	if x != nil {
		return x.ProblemId
	}
	return ""
}

func (x *Score_Problem) GetScore() float32 {
	if x != nil {
		return x.Score
	}
	return 0
}

func (x *Score_Problem) GetPenalty() float32 {
	if x != nil {
		return x.Penalty
	}
	return 0
}

func (x *Score_Problem) GetSolved() bool {
	if x != nil {
		return x.Solved
	}
	return false
}

func (x *Score_Problem) GetPercentage() float32 {
	if x != nil {
		return x.Percentage
	}
	return 0
}

func (x *Score_Problem) GetAttempts() uint32 {
	if x != nil {
		return x.Attempts
	}
	return 0
}

func (x *Score_Problem) GetSolvedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.SolvedAt
	}
	return nil
}

func (x *Score_Problem) GetSolvedIn() uint32 {
	if x != nil {
		return x.SolvedIn
	}
	return 0
}

func (x *Score_Problem) GetChanged() bool {
	if x != nil {
		return x.Changed
	}
	return false
}

func (x *Score_Problem) GetBreakdown() []*Score_Testset {
	if x != nil {
		return x.Breakdown
	}
	return nil
}

type Score_Testset struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TestsetId string  `protobuf:"bytes,1,opt,name=testset_id,json=testsetId,proto3" json:"testset_id,omitempty"`
	Index     uint32  `protobuf:"varint,2,opt,name=index,proto3" json:"index,omitempty"`
	Cost      float32 `protobuf:"fixed32,10,opt,name=cost,proto3" json:"cost,omitempty"`
	Score     float32 `protobuf:"fixed32,11,opt,name=score,proto3" json:"score,omitempty"`
}

func (x *Score_Testset) Reset() {
	*x = Score_Testset{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_judge_score_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Score_Testset) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Score_Testset) ProtoMessage() {}

func (x *Score_Testset) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_judge_score_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Score_Testset.ProtoReflect.Descriptor instead.
func (*Score_Testset) Descriptor() ([]byte, []int) {
	return file_eolymp_judge_score_proto_rawDescGZIP(), []int{0, 1}
}

func (x *Score_Testset) GetTestsetId() string {
	if x != nil {
		return x.TestsetId
	}
	return ""
}

func (x *Score_Testset) GetIndex() uint32 {
	if x != nil {
		return x.Index
	}
	return 0
}

func (x *Score_Testset) GetCost() float32 {
	if x != nil {
		return x.Cost
	}
	return 0
}

func (x *Score_Testset) GetScore() float32 {
	if x != nil {
		return x.Score
	}
	return 0
}

var File_eolymp_judge_score_proto protoreflect.FileDescriptor

var file_eolymp_judge_score_proto_rawDesc = []byte{
	0x0a, 0x18, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x6a, 0x75, 0x64, 0x67, 0x65, 0x2f, 0x73,
	0x63, 0x6f, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x65, 0x6f, 0x6c, 0x79,
	0x6d, 0x70, 0x2e, 0x6a, 0x75, 0x64, 0x67, 0x65, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9b, 0x06, 0x0a, 0x05, 0x53, 0x63,
	0x6f, 0x72, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x5f, 0x61, 0x66, 0x74,
	0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x41,
	0x66, 0x74, 0x65, 0x72, 0x12, 0x1f, 0x0a, 0x0b, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x5f, 0x75, 0x6e,
	0x74, 0x69, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x55, 0x6e, 0x74, 0x69, 0x6c, 0x12, 0x38, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12,
	0x14, 0x0a, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05,
	0x73, 0x63, 0x6f, 0x72, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x65, 0x6e, 0x61, 0x6c, 0x74, 0x79,
	0x18, 0x0b, 0x20, 0x01, 0x28, 0x02, 0x52, 0x07, 0x70, 0x65, 0x6e, 0x61, 0x6c, 0x74, 0x79, 0x12,
	0x18, 0x0a, 0x07, 0x75, 0x70, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x07, 0x75, 0x70, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x12, 0x39, 0x0a, 0x09, 0x62, 0x72, 0x65,
	0x61, 0x6b, 0x64, 0x6f, 0x77, 0x6e, 0x18, 0x14, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x65,
	0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x6a, 0x75, 0x64, 0x67, 0x65, 0x2e, 0x53, 0x63, 0x6f, 0x72,
	0x65, 0x2e, 0x50, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x52, 0x09, 0x62, 0x72, 0x65, 0x61, 0x6b,
	0x64, 0x6f, 0x77, 0x6e, 0x1a, 0xd7, 0x02, 0x0a, 0x07, 0x50, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d,
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
	0x08, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x64, 0x49, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x68, 0x61,
	0x6e, 0x67, 0x65, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x63, 0x68, 0x61, 0x6e,
	0x67, 0x65, 0x64, 0x12, 0x39, 0x0a, 0x09, 0x62, 0x72, 0x65, 0x61, 0x6b, 0x64, 0x6f, 0x77, 0x6e,
	0x18, 0x14, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e,
	0x6a, 0x75, 0x64, 0x67, 0x65, 0x2e, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x54, 0x65, 0x73, 0x74,
	0x73, 0x65, 0x74, 0x52, 0x09, 0x62, 0x72, 0x65, 0x61, 0x6b, 0x64, 0x6f, 0x77, 0x6e, 0x1a, 0x68,
	0x0a, 0x07, 0x54, 0x65, 0x73, 0x74, 0x73, 0x65, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x65, 0x73,
	0x74, 0x73, 0x65, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74,
	0x65, 0x73, 0x74, 0x73, 0x65, 0x74, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6e, 0x64, 0x65,
	0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x12,
	0x0a, 0x04, 0x63, 0x6f, 0x73, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x02, 0x52, 0x04, 0x63, 0x6f,
	0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28,
	0x02, 0x52, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x22, 0x4d, 0x0a, 0x0c, 0x46, 0x65, 0x74, 0x63,
	0x68, 0x69, 0x6e, 0x67, 0x4d, 0x6f, 0x64, 0x65, 0x12, 0x0a, 0x0a, 0x06, 0x41, 0x43, 0x54, 0x55,
	0x41, 0x4c, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x50, 0x55, 0x4e, 0x43, 0x54, 0x55, 0x41, 0x4c,
	0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x4c, 0x41, 0x54, 0x45, 0x53, 0x54, 0x10, 0x02, 0x12, 0x0a,
	0x0a, 0x06, 0x46, 0x52, 0x4f, 0x5a, 0x45, 0x4e, 0x10, 0x03, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x50,
	0x53, 0x4f, 0x4c, 0x56, 0x45, 0x10, 0x04, 0x42, 0x2d, 0x5a, 0x2b, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x67, 0x6f, 0x2d,
	0x73, 0x64, 0x6b, 0x2f, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x6a, 0x75, 0x64, 0x67, 0x65,
	0x3b, 0x6a, 0x75, 0x64, 0x67, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_eolymp_judge_score_proto_rawDescOnce sync.Once
	file_eolymp_judge_score_proto_rawDescData = file_eolymp_judge_score_proto_rawDesc
)

func file_eolymp_judge_score_proto_rawDescGZIP() []byte {
	file_eolymp_judge_score_proto_rawDescOnce.Do(func() {
		file_eolymp_judge_score_proto_rawDescData = protoimpl.X.CompressGZIP(file_eolymp_judge_score_proto_rawDescData)
	})
	return file_eolymp_judge_score_proto_rawDescData
}

var file_eolymp_judge_score_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_eolymp_judge_score_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_eolymp_judge_score_proto_goTypes = []interface{}{
	(Score_FetchingMode)(0),       // 0: eolymp.judge.Score.FetchingMode
	(*Score)(nil),                 // 1: eolymp.judge.Score
	(*Score_Problem)(nil),         // 2: eolymp.judge.Score.Problem
	(*Score_Testset)(nil),         // 3: eolymp.judge.Score.Testset
	(*timestamppb.Timestamp)(nil), // 4: google.protobuf.Timestamp
}
var file_eolymp_judge_score_proto_depIdxs = []int32{
	4, // 0: eolymp.judge.Score.timestamp:type_name -> google.protobuf.Timestamp
	2, // 1: eolymp.judge.Score.breakdown:type_name -> eolymp.judge.Score.Problem
	4, // 2: eolymp.judge.Score.Problem.solved_at:type_name -> google.protobuf.Timestamp
	3, // 3: eolymp.judge.Score.Problem.breakdown:type_name -> eolymp.judge.Score.Testset
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_eolymp_judge_score_proto_init() }
func file_eolymp_judge_score_proto_init() {
	if File_eolymp_judge_score_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_eolymp_judge_score_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Score); i {
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
		file_eolymp_judge_score_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Score_Problem); i {
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
		file_eolymp_judge_score_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Score_Testset); i {
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
			RawDescriptor: file_eolymp_judge_score_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_eolymp_judge_score_proto_goTypes,
		DependencyIndexes: file_eolymp_judge_score_proto_depIdxs,
		EnumInfos:         file_eolymp_judge_score_proto_enumTypes,
		MessageInfos:      file_eolymp_judge_score_proto_msgTypes,
	}.Build()
	File_eolymp_judge_score_proto = out.File
	file_eolymp_judge_score_proto_rawDesc = nil
	file_eolymp_judge_score_proto_goTypes = nil
	file_eolymp_judge_score_proto_depIdxs = nil
}
