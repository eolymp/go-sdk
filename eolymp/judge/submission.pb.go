// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v3.21.12
// source: eolymp/judge/submission.proto

package judge

import (
	atlas "github.com/eolymp/go-sdk/eolymp/atlas"
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

type Submission_Extra int32

const (
	Submission_NO_EXTRA Submission_Extra = 0
	Submission_SOURCE   Submission_Extra = 1
	Submission_GROUPS   Submission_Extra = 2
	Submission_RUNS     Submission_Extra = 3
)

// Enum value maps for Submission_Extra.
var (
	Submission_Extra_name = map[int32]string{
		0: "NO_EXTRA",
		1: "SOURCE",
		2: "GROUPS",
		3: "RUNS",
	}
	Submission_Extra_value = map[string]int32{
		"NO_EXTRA": 0,
		"SOURCE":   1,
		"GROUPS":   2,
		"RUNS":     3,
	}
)

func (x Submission_Extra) Enum() *Submission_Extra {
	p := new(Submission_Extra)
	*p = x
	return p
}

func (x Submission_Extra) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Submission_Extra) Descriptor() protoreflect.EnumDescriptor {
	return file_eolymp_judge_submission_proto_enumTypes[0].Descriptor()
}

func (Submission_Extra) Type() protoreflect.EnumType {
	return &file_eolymp_judge_submission_proto_enumTypes[0]
}

func (x Submission_Extra) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Submission_Extra.Descriptor instead.
func (Submission_Extra) EnumDescriptor() ([]byte, []int) {
	return file_eolymp_judge_submission_proto_rawDescGZIP(), []int{0, 0}
}

type Submission struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            string                   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"` // unique identifier
	Url           string                   `protobuf:"bytes,7,opt,name=url,proto3" json:"url,omitempty"`
	ContestId     string                   `protobuf:"bytes,2,opt,name=contest_id,json=contestId,proto3" json:"contest_id,omitempty"`                   // contest
	ProblemId     string                   `protobuf:"bytes,3,opt,name=problem_id,json=problemId,proto3" json:"problem_id,omitempty"`                   // problem
	ParticipantId string                   `protobuf:"bytes,4,opt,name=participant_id,json=participantId,proto3" json:"participant_id,omitempty"`       // submitter
	SubmittedAt   *timestamppb.Timestamp   `protobuf:"bytes,5,opt,name=submitted_at,json=submittedAt,proto3" json:"submitted_at,omitempty"`             // time when submission was created
	Deleted       bool                     `protobuf:"varint,6,opt,name=deleted,proto3" json:"deleted,omitempty"`                                       // mark deleted (excluded from scoring) submissions
	Lang          string                   `protobuf:"bytes,10,opt,name=lang,proto3" json:"lang,omitempty"`                                             // programming language
	Source        string                   `protobuf:"bytes,11,opt,name=source,proto3" json:"source,omitempty"`                                         // source code
	SourceUrl     string                   `protobuf:"bytes,110,opt,name=source_url,json=sourceUrl,proto3" json:"source_url,omitempty"`                 // source code URL (overrides source)
	Signature     string                   `protobuf:"bytes,12,opt,name=signature,proto3" json:"signature,omitempty"`                                   // source code
	Status        atlas.Submission_Status  `protobuf:"varint,20,opt,name=status,proto3,enum=eolymp.atlas.Submission_Status" json:"status,omitempty"`    // status (see explanation for enumeration values)
	Verdict       atlas.Submission_Verdict `protobuf:"varint,22,opt,name=verdict,proto3,enum=eolymp.atlas.Submission_Verdict" json:"verdict,omitempty"` // overall verdict based on verdicts in groups/runs
	Error         string                   `protobuf:"bytes,21,opt,name=error,proto3" json:"error,omitempty"`                                           // error message in case status is ERROR
	ErrorUrl      string                   `protobuf:"bytes,23,opt,name=error_url,json=errorUrl,proto3" json:"error_url,omitempty"`                     // a URL with error output
	Cost          float32                  `protobuf:"fixed32,30,opt,name=cost,proto3" json:"cost,omitempty"`                                           // maximum possible score for the submission
	Score         float32                  `protobuf:"fixed32,31,opt,name=score,proto3" json:"score,omitempty"`                                         // sum of earned points
	Percentage    float32                  `protobuf:"fixed32,32,opt,name=percentage,proto3" json:"percentage,omitempty"`
	Groups        []*Submission_Group      `protobuf:"bytes,50,rep,name=groups,proto3" json:"groups,omitempty"` // status for each run by group
}

func (x *Submission) Reset() {
	*x = Submission{}
	mi := &file_eolymp_judge_submission_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Submission) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Submission) ProtoMessage() {}

func (x *Submission) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_judge_submission_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Submission.ProtoReflect.Descriptor instead.
func (*Submission) Descriptor() ([]byte, []int) {
	return file_eolymp_judge_submission_proto_rawDescGZIP(), []int{0}
}

func (x *Submission) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Submission) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *Submission) GetContestId() string {
	if x != nil {
		return x.ContestId
	}
	return ""
}

func (x *Submission) GetProblemId() string {
	if x != nil {
		return x.ProblemId
	}
	return ""
}

func (x *Submission) GetParticipantId() string {
	if x != nil {
		return x.ParticipantId
	}
	return ""
}

func (x *Submission) GetSubmittedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.SubmittedAt
	}
	return nil
}

func (x *Submission) GetDeleted() bool {
	if x != nil {
		return x.Deleted
	}
	return false
}

func (x *Submission) GetLang() string {
	if x != nil {
		return x.Lang
	}
	return ""
}

func (x *Submission) GetSource() string {
	if x != nil {
		return x.Source
	}
	return ""
}

func (x *Submission) GetSourceUrl() string {
	if x != nil {
		return x.SourceUrl
	}
	return ""
}

func (x *Submission) GetSignature() string {
	if x != nil {
		return x.Signature
	}
	return ""
}

func (x *Submission) GetStatus() atlas.Submission_Status {
	if x != nil {
		return x.Status
	}
	return atlas.Submission_Status(0)
}

func (x *Submission) GetVerdict() atlas.Submission_Verdict {
	if x != nil {
		return x.Verdict
	}
	return atlas.Submission_Verdict(0)
}

func (x *Submission) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

func (x *Submission) GetErrorUrl() string {
	if x != nil {
		return x.ErrorUrl
	}
	return ""
}

func (x *Submission) GetCost() float32 {
	if x != nil {
		return x.Cost
	}
	return 0
}

func (x *Submission) GetScore() float32 {
	if x != nil {
		return x.Score
	}
	return 0
}

func (x *Submission) GetPercentage() float32 {
	if x != nil {
		return x.Percentage
	}
	return 0
}

func (x *Submission) GetGroups() []*Submission_Group {
	if x != nil {
		return x.Groups
	}
	return nil
}

type Submission_Run struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            string                   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Index         uint32                   `protobuf:"varint,10,opt,name=index,proto3" json:"index,omitempty"`
	TestId        string                   `protobuf:"bytes,11,opt,name=test_id,json=testId,proto3" json:"test_id,omitempty"`
	Cost          float32                  `protobuf:"fixed32,12,opt,name=cost,proto3" json:"cost,omitempty"`
	Score         float32                  `protobuf:"fixed32,13,opt,name=score,proto3" json:"score,omitempty"`
	WallTimeUsage uint32                   `protobuf:"varint,2,opt,name=wall_time_usage,json=wallTimeUsage,proto3" json:"wall_time_usage,omitempty"` // wall time (real-world time) usage
	CpuTimeUsage  uint32                   `protobuf:"varint,3,opt,name=cpu_time_usage,json=cpuTimeUsage,proto3" json:"cpu_time_usage,omitempty"`    // cpu time (time cpu was active)
	MemoryUsage   uint64                   `protobuf:"varint,4,opt,name=memory_usage,json=memoryUsage,proto3" json:"memory_usage,omitempty"`
	Status        atlas.Submission_Status  `protobuf:"varint,20,opt,name=status,proto3,enum=eolymp.atlas.Submission_Status" json:"status,omitempty"`
	Verdict       atlas.Submission_Verdict `protobuf:"varint,21,opt,name=verdict,proto3,enum=eolymp.atlas.Submission_Verdict" json:"verdict,omitempty"` // overall verdict based on verdicts in groups/runs
}

func (x *Submission_Run) Reset() {
	*x = Submission_Run{}
	mi := &file_eolymp_judge_submission_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Submission_Run) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Submission_Run) ProtoMessage() {}

func (x *Submission_Run) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_judge_submission_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Submission_Run.ProtoReflect.Descriptor instead.
func (*Submission_Run) Descriptor() ([]byte, []int) {
	return file_eolymp_judge_submission_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Submission_Run) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Submission_Run) GetIndex() uint32 {
	if x != nil {
		return x.Index
	}
	return 0
}

func (x *Submission_Run) GetTestId() string {
	if x != nil {
		return x.TestId
	}
	return ""
}

func (x *Submission_Run) GetCost() float32 {
	if x != nil {
		return x.Cost
	}
	return 0
}

func (x *Submission_Run) GetScore() float32 {
	if x != nil {
		return x.Score
	}
	return 0
}

func (x *Submission_Run) GetWallTimeUsage() uint32 {
	if x != nil {
		return x.WallTimeUsage
	}
	return 0
}

func (x *Submission_Run) GetCpuTimeUsage() uint32 {
	if x != nil {
		return x.CpuTimeUsage
	}
	return 0
}

func (x *Submission_Run) GetMemoryUsage() uint64 {
	if x != nil {
		return x.MemoryUsage
	}
	return 0
}

func (x *Submission_Run) GetStatus() atlas.Submission_Status {
	if x != nil {
		return x.Status
	}
	return atlas.Submission_Status(0)
}

func (x *Submission_Run) GetVerdict() atlas.Submission_Verdict {
	if x != nil {
		return x.Verdict
	}
	return atlas.Submission_Verdict(0)
}

type Submission_Group struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Index          uint32                   `protobuf:"varint,1,opt,name=index,proto3" json:"index,omitempty"`                                                                           // group index
	TestsetId      string                   `protobuf:"bytes,2,opt,name=testset_id,json=testsetId,proto3" json:"testset_id,omitempty"`                                                   // testset associated with the group
	Status         atlas.Submission_Status  `protobuf:"varint,10,opt,name=status,proto3,enum=eolymp.atlas.Submission_Status" json:"status,omitempty"`                                    // status of the group
	Verdict        atlas.Submission_Verdict `protobuf:"varint,12,opt,name=verdict,proto3,enum=eolymp.atlas.Submission_Verdict" json:"verdict,omitempty"`                                 // overall verdict based on verdicts in groups/runs
	Dependencies   []uint32                 `protobuf:"varint,11,rep,packed,name=dependencies,proto3" json:"dependencies,omitempty"`                                                     // other group indices which need to pass for this group to run
	Cost           float32                  `protobuf:"fixed32,20,opt,name=cost,proto3" json:"cost,omitempty"`                                                                           // max possible score for passing all tests in the group
	Score          float32                  `protobuf:"fixed32,21,opt,name=score,proto3" json:"score,omitempty"`                                                                         // sum of earned points within a group
	ScoringMode    atlas.ScoringMode        `protobuf:"varint,22,opt,name=scoring_mode,json=scoringMode,proto3,enum=eolymp.atlas.ScoringMode" json:"scoring_mode,omitempty"`             // how group is scored
	FeedbackPolicy atlas.FeedbackPolicy     `protobuf:"varint,30,opt,name=feedback_policy,json=feedbackPolicy,proto3,enum=eolymp.atlas.FeedbackPolicy" json:"feedback_policy,omitempty"` // how tests are shown to the user
	WallTimeUsage  uint32                   `protobuf:"varint,41,opt,name=wall_time_usage,json=wallTimeUsage,proto3" json:"wall_time_usage,omitempty"`                                   // provides feedback on wall time usage within the group, depending on feedback mode it might be max execution time in group or time usage in the first non-accepted test
	CpuTimeUsage   uint32                   `protobuf:"varint,42,opt,name=cpu_time_usage,json=cpuTimeUsage,proto3" json:"cpu_time_usage,omitempty"`                                      // provides feedback on CPU time usage within the group, depending on feedback mode it might be max execution time in group or time usage in the first non-accepted test
	MemoryUsage    uint64                   `protobuf:"varint,46,opt,name=memory_usage,json=memoryUsage,proto3" json:"memory_usage,omitempty"`                                           // provides feedback on memory usage within the group, depending on feedback mode it might be memory usage peak in group or memory usage in the first non-accepted test
	Runs           []*Submission_Run        `protobuf:"bytes,100,rep,name=runs,proto3" json:"runs,omitempty"`                                                                            // runs of the group
}

func (x *Submission_Group) Reset() {
	*x = Submission_Group{}
	mi := &file_eolymp_judge_submission_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Submission_Group) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Submission_Group) ProtoMessage() {}

func (x *Submission_Group) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_judge_submission_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Submission_Group.ProtoReflect.Descriptor instead.
func (*Submission_Group) Descriptor() ([]byte, []int) {
	return file_eolymp_judge_submission_proto_rawDescGZIP(), []int{0, 1}
}

func (x *Submission_Group) GetIndex() uint32 {
	if x != nil {
		return x.Index
	}
	return 0
}

func (x *Submission_Group) GetTestsetId() string {
	if x != nil {
		return x.TestsetId
	}
	return ""
}

func (x *Submission_Group) GetStatus() atlas.Submission_Status {
	if x != nil {
		return x.Status
	}
	return atlas.Submission_Status(0)
}

func (x *Submission_Group) GetVerdict() atlas.Submission_Verdict {
	if x != nil {
		return x.Verdict
	}
	return atlas.Submission_Verdict(0)
}

func (x *Submission_Group) GetDependencies() []uint32 {
	if x != nil {
		return x.Dependencies
	}
	return nil
}

func (x *Submission_Group) GetCost() float32 {
	if x != nil {
		return x.Cost
	}
	return 0
}

func (x *Submission_Group) GetScore() float32 {
	if x != nil {
		return x.Score
	}
	return 0
}

func (x *Submission_Group) GetScoringMode() atlas.ScoringMode {
	if x != nil {
		return x.ScoringMode
	}
	return atlas.ScoringMode(0)
}

func (x *Submission_Group) GetFeedbackPolicy() atlas.FeedbackPolicy {
	if x != nil {
		return x.FeedbackPolicy
	}
	return atlas.FeedbackPolicy(0)
}

func (x *Submission_Group) GetWallTimeUsage() uint32 {
	if x != nil {
		return x.WallTimeUsage
	}
	return 0
}

func (x *Submission_Group) GetCpuTimeUsage() uint32 {
	if x != nil {
		return x.CpuTimeUsage
	}
	return 0
}

func (x *Submission_Group) GetMemoryUsage() uint64 {
	if x != nil {
		return x.MemoryUsage
	}
	return 0
}

func (x *Submission_Group) GetRuns() []*Submission_Run {
	if x != nil {
		return x.Runs
	}
	return nil
}

var File_eolymp_judge_submission_proto protoreflect.FileDescriptor

var file_eolymp_judge_submission_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x6a, 0x75, 0x64, 0x67, 0x65, 0x2f, 0x73,
	0x75, 0x62, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0c, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x6a, 0x75, 0x64, 0x67, 0x65, 0x1a, 0x1d, 0x65,
	0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x61, 0x74, 0x6c, 0x61, 0x73, 0x2f, 0x73, 0x75, 0x62, 0x6d,
	0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x23, 0x65, 0x6f,
	0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x61, 0x74, 0x6c, 0x61, 0x73, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x69,
	0x6e, 0x67, 0x5f, 0x66, 0x65, 0x65, 0x64, 0x62, 0x61, 0x63, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x22, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x61, 0x74, 0x6c, 0x61, 0x73, 0x2f,
	0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x73, 0x63, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb9, 0x0c, 0x0a, 0x0a, 0x53, 0x75, 0x62, 0x6d, 0x69,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6f, 0x6e, 0x74, 0x65,
	0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6f, 0x6e,
	0x74, 0x65, 0x73, 0x74, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x62, 0x6c, 0x65,
	0x6d, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x62,
	0x6c, 0x65, 0x6d, 0x49, 0x64, 0x12, 0x25, 0x0a, 0x0e, 0x70, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69,
	0x70, 0x61, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x70,
	0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x3d, 0x0a, 0x0c,
	0x73, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0b,
	0x73, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x64,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x64, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x61, 0x6e, 0x67, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6c, 0x61, 0x6e, 0x67, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x75, 0x72, 0x6c, 0x18,
	0x6e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x55, 0x72, 0x6c,
	0x12, 0x1c, 0x0a, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x0c, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x37,
	0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x14, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1f,
	0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x61, 0x74, 0x6c, 0x61, 0x73, 0x2e, 0x53, 0x75,
	0x62, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x3a, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x64, 0x69,
	0x63, 0x74, 0x18, 0x16, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x20, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d,
	0x70, 0x2e, 0x61, 0x74, 0x6c, 0x61, 0x73, 0x2e, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x2e, 0x56, 0x65, 0x72, 0x64, 0x69, 0x63, 0x74, 0x52, 0x07, 0x76, 0x65, 0x72, 0x64,
	0x69, 0x63, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x15, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x1b, 0x0a, 0x09, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x17, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x55, 0x72, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x73, 0x74, 0x18, 0x1e,
	0x20, 0x01, 0x28, 0x02, 0x52, 0x04, 0x63, 0x6f, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x63,
	0x6f, 0x72, 0x65, 0x18, 0x1f, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65,
	0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x61, 0x67, 0x65, 0x18, 0x20,
	0x20, 0x01, 0x28, 0x02, 0x52, 0x0a, 0x70, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x61, 0x67, 0x65,
	0x12, 0x36, 0x0a, 0x06, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x18, 0x32, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x1e, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x6a, 0x75, 0x64, 0x67, 0x65, 0x2e,
	0x53, 0x75, 0x62, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x52, 0x06, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x1a, 0xd4, 0x02, 0x0a, 0x03, 0x52, 0x75, 0x6e,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x14, 0x0a, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x17, 0x0a, 0x07, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x69,
	0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x65, 0x73, 0x74, 0x49, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x63, 0x6f, 0x73, 0x74, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x02, 0x52, 0x04, 0x63,
	0x6f, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x0d, 0x20, 0x01,
	0x28, 0x02, 0x52, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x12, 0x26, 0x0a, 0x0f, 0x77, 0x61, 0x6c,
	0x6c, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x75, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x0d, 0x77, 0x61, 0x6c, 0x6c, 0x54, 0x69, 0x6d, 0x65, 0x55, 0x73, 0x61, 0x67,
	0x65, 0x12, 0x24, 0x0a, 0x0e, 0x63, 0x70, 0x75, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x75, 0x73,
	0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c, 0x63, 0x70, 0x75, 0x54, 0x69,
	0x6d, 0x65, 0x55, 0x73, 0x61, 0x67, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x6d, 0x65, 0x6d, 0x6f, 0x72,
	0x79, 0x5f, 0x75, 0x73, 0x61, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0b, 0x6d,
	0x65, 0x6d, 0x6f, 0x72, 0x79, 0x55, 0x73, 0x61, 0x67, 0x65, 0x12, 0x37, 0x0a, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x14, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1f, 0x2e, 0x65, 0x6f, 0x6c,
	0x79, 0x6d, 0x70, 0x2e, 0x61, 0x74, 0x6c, 0x61, 0x73, 0x2e, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x3a, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x64, 0x69, 0x63, 0x74, 0x18, 0x15,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x20, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x61, 0x74,
	0x6c, 0x61, 0x73, 0x2e, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x56,
	0x65, 0x72, 0x64, 0x69, 0x63, 0x74, 0x52, 0x07, 0x76, 0x65, 0x72, 0x64, 0x69, 0x63, 0x74, 0x1a,
	0xa7, 0x04, 0x0a, 0x05, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6e, 0x64,
	0x65, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x12,
	0x1d, 0x0a, 0x0a, 0x74, 0x65, 0x73, 0x74, 0x73, 0x65, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x65, 0x73, 0x74, 0x73, 0x65, 0x74, 0x49, 0x64, 0x12, 0x37,
	0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1f,
	0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x61, 0x74, 0x6c, 0x61, 0x73, 0x2e, 0x53, 0x75,
	0x62, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x3a, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x64, 0x69,
	0x63, 0x74, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x20, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d,
	0x70, 0x2e, 0x61, 0x74, 0x6c, 0x61, 0x73, 0x2e, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x2e, 0x56, 0x65, 0x72, 0x64, 0x69, 0x63, 0x74, 0x52, 0x07, 0x76, 0x65, 0x72, 0x64,
	0x69, 0x63, 0x74, 0x12, 0x22, 0x0a, 0x0c, 0x64, 0x65, 0x70, 0x65, 0x6e, 0x64, 0x65, 0x6e, 0x63,
	0x69, 0x65, 0x73, 0x18, 0x0b, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x0c, 0x64, 0x65, 0x70, 0x65, 0x6e,
	0x64, 0x65, 0x6e, 0x63, 0x69, 0x65, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x73, 0x74, 0x18,
	0x14, 0x20, 0x01, 0x28, 0x02, 0x52, 0x04, 0x63, 0x6f, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x73,
	0x63, 0x6f, 0x72, 0x65, 0x18, 0x15, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x73, 0x63, 0x6f, 0x72,
	0x65, 0x12, 0x3c, 0x0a, 0x0c, 0x73, 0x63, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x6d, 0x6f, 0x64,
	0x65, 0x18, 0x16, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x19, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70,
	0x2e, 0x61, 0x74, 0x6c, 0x61, 0x73, 0x2e, 0x53, 0x63, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x4d, 0x6f,
	0x64, 0x65, 0x52, 0x0b, 0x73, 0x63, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x4d, 0x6f, 0x64, 0x65, 0x12,
	0x45, 0x0a, 0x0f, 0x66, 0x65, 0x65, 0x64, 0x62, 0x61, 0x63, 0x6b, 0x5f, 0x70, 0x6f, 0x6c, 0x69,
	0x63, 0x79, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1c, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d,
	0x70, 0x2e, 0x61, 0x74, 0x6c, 0x61, 0x73, 0x2e, 0x46, 0x65, 0x65, 0x64, 0x62, 0x61, 0x63, 0x6b,
	0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x0e, 0x66, 0x65, 0x65, 0x64, 0x62, 0x61, 0x63, 0x6b,
	0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0x26, 0x0a, 0x0f, 0x77, 0x61, 0x6c, 0x6c, 0x5f, 0x74,
	0x69, 0x6d, 0x65, 0x5f, 0x75, 0x73, 0x61, 0x67, 0x65, 0x18, 0x29, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x0d, 0x77, 0x61, 0x6c, 0x6c, 0x54, 0x69, 0x6d, 0x65, 0x55, 0x73, 0x61, 0x67, 0x65, 0x12, 0x24,
	0x0a, 0x0e, 0x63, 0x70, 0x75, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x75, 0x73, 0x61, 0x67, 0x65,
	0x18, 0x2a, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c, 0x63, 0x70, 0x75, 0x54, 0x69, 0x6d, 0x65, 0x55,
	0x73, 0x61, 0x67, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x5f, 0x75,
	0x73, 0x61, 0x67, 0x65, 0x18, 0x2e, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0b, 0x6d, 0x65, 0x6d, 0x6f,
	0x72, 0x79, 0x55, 0x73, 0x61, 0x67, 0x65, 0x12, 0x30, 0x0a, 0x04, 0x72, 0x75, 0x6e, 0x73, 0x18,
	0x64, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x6a,
	0x75, 0x64, 0x67, 0x65, 0x2e, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e,
	0x52, 0x75, 0x6e, 0x52, 0x04, 0x72, 0x75, 0x6e, 0x73, 0x22, 0x37, 0x0a, 0x05, 0x45, 0x78, 0x74,
	0x72, 0x61, 0x12, 0x0c, 0x0a, 0x08, 0x4e, 0x4f, 0x5f, 0x45, 0x58, 0x54, 0x52, 0x41, 0x10, 0x00,
	0x12, 0x0a, 0x0a, 0x06, 0x53, 0x4f, 0x55, 0x52, 0x43, 0x45, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06,
	0x47, 0x52, 0x4f, 0x55, 0x50, 0x53, 0x10, 0x02, 0x12, 0x08, 0x0a, 0x04, 0x52, 0x55, 0x4e, 0x53,
	0x10, 0x03, 0x42, 0x2d, 0x5a, 0x2b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x67, 0x6f, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x65,
	0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x6a, 0x75, 0x64, 0x67, 0x65, 0x3b, 0x6a, 0x75, 0x64, 0x67,
	0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_eolymp_judge_submission_proto_rawDescOnce sync.Once
	file_eolymp_judge_submission_proto_rawDescData = file_eolymp_judge_submission_proto_rawDesc
)

func file_eolymp_judge_submission_proto_rawDescGZIP() []byte {
	file_eolymp_judge_submission_proto_rawDescOnce.Do(func() {
		file_eolymp_judge_submission_proto_rawDescData = protoimpl.X.CompressGZIP(file_eolymp_judge_submission_proto_rawDescData)
	})
	return file_eolymp_judge_submission_proto_rawDescData
}

var file_eolymp_judge_submission_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_eolymp_judge_submission_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_eolymp_judge_submission_proto_goTypes = []any{
	(Submission_Extra)(0),         // 0: eolymp.judge.Submission.Extra
	(*Submission)(nil),            // 1: eolymp.judge.Submission
	(*Submission_Run)(nil),        // 2: eolymp.judge.Submission.Run
	(*Submission_Group)(nil),      // 3: eolymp.judge.Submission.Group
	(*timestamppb.Timestamp)(nil), // 4: google.protobuf.Timestamp
	(atlas.Submission_Status)(0),  // 5: eolymp.atlas.Submission.Status
	(atlas.Submission_Verdict)(0), // 6: eolymp.atlas.Submission.Verdict
	(atlas.ScoringMode)(0),        // 7: eolymp.atlas.ScoringMode
	(atlas.FeedbackPolicy)(0),     // 8: eolymp.atlas.FeedbackPolicy
}
var file_eolymp_judge_submission_proto_depIdxs = []int32{
	4,  // 0: eolymp.judge.Submission.submitted_at:type_name -> google.protobuf.Timestamp
	5,  // 1: eolymp.judge.Submission.status:type_name -> eolymp.atlas.Submission.Status
	6,  // 2: eolymp.judge.Submission.verdict:type_name -> eolymp.atlas.Submission.Verdict
	3,  // 3: eolymp.judge.Submission.groups:type_name -> eolymp.judge.Submission.Group
	5,  // 4: eolymp.judge.Submission.Run.status:type_name -> eolymp.atlas.Submission.Status
	6,  // 5: eolymp.judge.Submission.Run.verdict:type_name -> eolymp.atlas.Submission.Verdict
	5,  // 6: eolymp.judge.Submission.Group.status:type_name -> eolymp.atlas.Submission.Status
	6,  // 7: eolymp.judge.Submission.Group.verdict:type_name -> eolymp.atlas.Submission.Verdict
	7,  // 8: eolymp.judge.Submission.Group.scoring_mode:type_name -> eolymp.atlas.ScoringMode
	8,  // 9: eolymp.judge.Submission.Group.feedback_policy:type_name -> eolymp.atlas.FeedbackPolicy
	2,  // 10: eolymp.judge.Submission.Group.runs:type_name -> eolymp.judge.Submission.Run
	11, // [11:11] is the sub-list for method output_type
	11, // [11:11] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_eolymp_judge_submission_proto_init() }
func file_eolymp_judge_submission_proto_init() {
	if File_eolymp_judge_submission_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_eolymp_judge_submission_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_eolymp_judge_submission_proto_goTypes,
		DependencyIndexes: file_eolymp_judge_submission_proto_depIdxs,
		EnumInfos:         file_eolymp_judge_submission_proto_enumTypes,
		MessageInfos:      file_eolymp_judge_submission_proto_msgTypes,
	}.Build()
	File_eolymp_judge_submission_proto = out.File
	file_eolymp_judge_submission_proto_rawDesc = nil
	file_eolymp_judge_submission_proto_goTypes = nil
	file_eolymp_judge_submission_proto_depIdxs = nil
}
