// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.18.1
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

type Submission_Status int32

const (
	Submission_NONE     Submission_Status = 0 // reserved, should not be used
	Submission_PENDING  Submission_Status = 1 // submission created and awaiting testing
	Submission_TESTING  Submission_Status = 2 // submission is being tested by an agent
	Submission_TIMEOUT  Submission_Status = 3 // submission testing took too long (too many timeouts in tests)
	Submission_COMPLETE Submission_Status = 4 // submission testing complete (score is populated)
	Submission_ERROR    Submission_Status = 5 // submission produced an error (eg. compilation error)
	Submission_FAILURE  Submission_Status = 6 // submission testing failed due to system error (or problem configuration error)
)

// Enum value maps for Submission_Status.
var (
	Submission_Status_name = map[int32]string{
		0: "NONE",
		1: "PENDING",
		2: "TESTING",
		3: "TIMEOUT",
		4: "COMPLETE",
		5: "ERROR",
		6: "FAILURE",
	}
	Submission_Status_value = map[string]int32{
		"NONE":     0,
		"PENDING":  1,
		"TESTING":  2,
		"TIMEOUT":  3,
		"COMPLETE": 4,
		"ERROR":    5,
		"FAILURE":  6,
	}
)

func (x Submission_Status) Enum() *Submission_Status {
	p := new(Submission_Status)
	*p = x
	return p
}

func (x Submission_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Submission_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_eolymp_judge_submission_proto_enumTypes[0].Descriptor()
}

func (Submission_Status) Type() protoreflect.EnumType {
	return &file_eolymp_judge_submission_proto_enumTypes[0]
}

func (x Submission_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Submission_Status.Descriptor instead.
func (Submission_Status) EnumDescriptor() ([]byte, []int) {
	return file_eolymp_judge_submission_proto_rawDescGZIP(), []int{0, 0}
}

type Submission_Run_Status int32

const (
	Submission_Run_NONE               Submission_Run_Status = 0  // reserved, should not be used
	Submission_Run_CREATED            Submission_Run_Status = 1  // awaiting execution
	Submission_Run_EXECUTING          Submission_Run_Status = 2  // running
	Submission_Run_TIMEOUT            Submission_Run_Status = 3  // wall time limit exceeded
	Submission_Run_CPU_EXHAUSTED      Submission_Run_Status = 4  // CPU time limit exceeded
	Submission_Run_MEMORY_OVERFLOW    Submission_Run_Status = 5  // memory limit exceeded
	Submission_Run_WRONG_ANSWER       Submission_Run_Status = 6  // wrong answer
	Submission_Run_ACCEPTED           Submission_Run_Status = 7  // answer accepted
	Submission_Run_RUNTIME_ERROR      Submission_Run_Status = 8  // program returned non 0 exit code
	Submission_Run_VERIFICATION_ERROR Submission_Run_Status = 9  // verifier couldn't run
	Submission_Run_SKIPPED            Submission_Run_Status = 10 // run won't be executed
	Submission_Run_BLOCKED            Submission_Run_Status = 11 // dependant test did not pass
)

// Enum value maps for Submission_Run_Status.
var (
	Submission_Run_Status_name = map[int32]string{
		0:  "NONE",
		1:  "CREATED",
		2:  "EXECUTING",
		3:  "TIMEOUT",
		4:  "CPU_EXHAUSTED",
		5:  "MEMORY_OVERFLOW",
		6:  "WRONG_ANSWER",
		7:  "ACCEPTED",
		8:  "RUNTIME_ERROR",
		9:  "VERIFICATION_ERROR",
		10: "SKIPPED",
		11: "BLOCKED",
	}
	Submission_Run_Status_value = map[string]int32{
		"NONE":               0,
		"CREATED":            1,
		"EXECUTING":          2,
		"TIMEOUT":            3,
		"CPU_EXHAUSTED":      4,
		"MEMORY_OVERFLOW":    5,
		"WRONG_ANSWER":       6,
		"ACCEPTED":           7,
		"RUNTIME_ERROR":      8,
		"VERIFICATION_ERROR": 9,
		"SKIPPED":            10,
		"BLOCKED":            11,
	}
)

func (x Submission_Run_Status) Enum() *Submission_Run_Status {
	p := new(Submission_Run_Status)
	*p = x
	return p
}

func (x Submission_Run_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Submission_Run_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_eolymp_judge_submission_proto_enumTypes[1].Descriptor()
}

func (Submission_Run_Status) Type() protoreflect.EnumType {
	return &file_eolymp_judge_submission_proto_enumTypes[1]
}

func (x Submission_Run_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Submission_Run_Status.Descriptor instead.
func (Submission_Run_Status) EnumDescriptor() ([]byte, []int) {
	return file_eolymp_judge_submission_proto_rawDescGZIP(), []int{0, 0, 0}
}

type Submission_Group_Status int32

const (
	Submission_Group_UNKNOWN            Submission_Group_Status = 0  // reserved, not supposed to be used
	Submission_Group_BLOCKED            Submission_Group_Status = 1  // group is waiting for dependant groups to pass
	Submission_Group_PENDING            Submission_Group_Status = 2  // none of the runs of the group has executed
	Submission_Group_TESTING            Submission_Group_Status = 3  // not all runs of the group has been executed
	Submission_Group_ACCEPTED           Submission_Group_Status = 4  // all tests in the group are Accepted
	Submission_Group_TIMEOUT            Submission_Group_Status = 5  // first not accepted run has status Timeout
	Submission_Group_CPU_EXHAUSTED      Submission_Group_Status = 6  // CPU time limit exceeded
	Submission_Group_MEMORY_OVERFLOW    Submission_Group_Status = 7  // first not accepted run has status Memory limit exceeded
	Submission_Group_WRONG_ANSWER       Submission_Group_Status = 8  // first not accepted run has status Wrong Answer
	Submission_Group_RUNTIME_ERROR      Submission_Group_Status = 9  // first not accepted run has status Runtime Error
	Submission_Group_VERIFICATION_ERROR Submission_Group_Status = 10 // first not accepted run has status Verification Error
	Submission_Group_SKIPPED            Submission_Group_Status = 11 // first not accepted run has status Skipped
)

// Enum value maps for Submission_Group_Status.
var (
	Submission_Group_Status_name = map[int32]string{
		0:  "UNKNOWN",
		1:  "BLOCKED",
		2:  "PENDING",
		3:  "TESTING",
		4:  "ACCEPTED",
		5:  "TIMEOUT",
		6:  "CPU_EXHAUSTED",
		7:  "MEMORY_OVERFLOW",
		8:  "WRONG_ANSWER",
		9:  "RUNTIME_ERROR",
		10: "VERIFICATION_ERROR",
		11: "SKIPPED",
	}
	Submission_Group_Status_value = map[string]int32{
		"UNKNOWN":            0,
		"BLOCKED":            1,
		"PENDING":            2,
		"TESTING":            3,
		"ACCEPTED":           4,
		"TIMEOUT":            5,
		"CPU_EXHAUSTED":      6,
		"MEMORY_OVERFLOW":    7,
		"WRONG_ANSWER":       8,
		"RUNTIME_ERROR":      9,
		"VERIFICATION_ERROR": 10,
		"SKIPPED":            11,
	}
)

func (x Submission_Group_Status) Enum() *Submission_Group_Status {
	p := new(Submission_Group_Status)
	*p = x
	return p
}

func (x Submission_Group_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Submission_Group_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_eolymp_judge_submission_proto_enumTypes[2].Descriptor()
}

func (Submission_Group_Status) Type() protoreflect.EnumType {
	return &file_eolymp_judge_submission_proto_enumTypes[2]
}

func (x Submission_Group_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Submission_Group_Status.Descriptor instead.
func (Submission_Group_Status) EnumDescriptor() ([]byte, []int) {
	return file_eolymp_judge_submission_proto_rawDescGZIP(), []int{0, 1, 0}
}

type Submission struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"` // unique identifier
	Ern           string                 `protobuf:"bytes,9999,opt,name=ern,proto3" json:"ern,omitempty"`
	ContestId     string                 `protobuf:"bytes,2,opt,name=contest_id,json=contestId,proto3" json:"contest_id,omitempty"`                // contest
	ProblemId     string                 `protobuf:"bytes,3,opt,name=problem_id,json=problemId,proto3" json:"problem_id,omitempty"`                // problem
	ParticipantId string                 `protobuf:"bytes,4,opt,name=participant_id,json=participantId,proto3" json:"participant_id,omitempty"`    // submitter
	SubmittedAt   *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=submitted_at,json=submittedAt,proto3" json:"submitted_at,omitempty"`          // time when submission was created
	Deleted       bool                   `protobuf:"varint,6,opt,name=deleted,proto3" json:"deleted,omitempty"`                                    // mark deleted (excluded from scoring) submissions
	Lang          string                 `protobuf:"bytes,10,opt,name=lang,proto3" json:"lang,omitempty"`                                          // programming language
	Source        string                 `protobuf:"bytes,11,opt,name=source,proto3" json:"source,omitempty"`                                      // source code
	Signature     string                 `protobuf:"bytes,12,opt,name=signature,proto3" json:"signature,omitempty"`                                // source code
	Status        Submission_Status      `protobuf:"varint,20,opt,name=status,proto3,enum=eolymp.judge.Submission_Status" json:"status,omitempty"` // status (see explanation for enumeration values)
	Error         string                 `protobuf:"bytes,21,opt,name=error,proto3" json:"error,omitempty"`                                        // error message in case status is ERROR
	Cost          float32                `protobuf:"fixed32,30,opt,name=cost,proto3" json:"cost,omitempty"`                                        // maximum possible score for the submission
	Score         float32                `protobuf:"fixed32,31,opt,name=score,proto3" json:"score,omitempty"`                                      // sum of earned points
	Percentage    float32                `protobuf:"fixed32,32,opt,name=percentage,proto3" json:"percentage,omitempty"`
	Groups        []*Submission_Group    `protobuf:"bytes,50,rep,name=groups,proto3" json:"groups,omitempty"` // status for each run by group
}

func (x *Submission) Reset() {
	*x = Submission{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_judge_submission_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Submission) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Submission) ProtoMessage() {}

func (x *Submission) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_judge_submission_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
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

func (x *Submission) GetErn() string {
	if x != nil {
		return x.Ern
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

func (x *Submission) GetSignature() string {
	if x != nil {
		return x.Signature
	}
	return ""
}

func (x *Submission) GetStatus() Submission_Status {
	if x != nil {
		return x.Status
	}
	return Submission_NONE
}

func (x *Submission) GetError() string {
	if x != nil {
		return x.Error
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

	Id            string                `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	WallTimeUsage uint32                `protobuf:"varint,2,opt,name=wall_time_usage,json=wallTimeUsage,proto3" json:"wall_time_usage,omitempty"` // wall time (real-world time) usage
	CpuTimeUsage  uint32                `protobuf:"varint,3,opt,name=cpu_time_usage,json=cpuTimeUsage,proto3" json:"cpu_time_usage,omitempty"`    // cpu time (time cpu was active)
	MemoryUsage   uint64                `protobuf:"varint,4,opt,name=memory_usage,json=memoryUsage,proto3" json:"memory_usage,omitempty"`
	Index         uint32                `protobuf:"varint,10,opt,name=index,proto3" json:"index,omitempty"`
	TestId        string                `protobuf:"bytes,11,opt,name=test_id,json=testId,proto3" json:"test_id,omitempty"`
	Cost          float32               `protobuf:"fixed32,12,opt,name=cost,proto3" json:"cost,omitempty"`
	Score         float32               `protobuf:"fixed32,13,opt,name=score,proto3" json:"score,omitempty"`
	Status        Submission_Run_Status `protobuf:"varint,20,opt,name=status,proto3,enum=eolymp.judge.Submission_Run_Status" json:"status,omitempty"`
}

func (x *Submission_Run) Reset() {
	*x = Submission_Run{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_judge_submission_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Submission_Run) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Submission_Run) ProtoMessage() {}

func (x *Submission_Run) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_judge_submission_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
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

func (x *Submission_Run) GetStatus() Submission_Run_Status {
	if x != nil {
		return x.Status
	}
	return Submission_Run_NONE
}

type Submission_Group struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Index          uint32                  `protobuf:"varint,1,opt,name=index,proto3" json:"index,omitempty"`                                                                           // group index
	TestsetId      string                  `protobuf:"bytes,2,opt,name=testset_id,json=testsetId,proto3" json:"testset_id,omitempty"`                                                   // testset associated with the group
	Status         Submission_Group_Status `protobuf:"varint,10,opt,name=status,proto3,enum=eolymp.judge.Submission_Group_Status" json:"status,omitempty"`                              // status of the group
	Dependencies   []uint32                `protobuf:"varint,11,rep,packed,name=dependencies,proto3" json:"dependencies,omitempty"`                                                     // other group indices which need to pass for this group to run
	Cost           float32                 `protobuf:"fixed32,20,opt,name=cost,proto3" json:"cost,omitempty"`                                                                           // max possible score for passing all tests in the group
	Score          float32                 `protobuf:"fixed32,21,opt,name=score,proto3" json:"score,omitempty"`                                                                         // sum of earned points within a group
	ScoringMode    atlas.ScoringMode       `protobuf:"varint,22,opt,name=scoring_mode,json=scoringMode,proto3,enum=eolymp.atlas.ScoringMode" json:"scoring_mode,omitempty"`             // how group is scored
	FeedbackPolicy atlas.FeedbackPolicy    `protobuf:"varint,30,opt,name=feedback_policy,json=feedbackPolicy,proto3,enum=eolymp.atlas.FeedbackPolicy" json:"feedback_policy,omitempty"` // how tests are shown to the user
	WallTimeUsage  uint32                  `protobuf:"varint,41,opt,name=wall_time_usage,json=wallTimeUsage,proto3" json:"wall_time_usage,omitempty"`                                   // provides feedback on wall time usage within the group, depending on feedback mode it might be max execution time in group or time usage in the first non-accepted test
	CpuTimeUsage   uint32                  `protobuf:"varint,42,opt,name=cpu_time_usage,json=cpuTimeUsage,proto3" json:"cpu_time_usage,omitempty"`                                      // provides feedback on CPU time usage within the group, depending on feedback mode it might be max execution time in group or time usage in the first non-accepted test
	MemoryUsage    uint64                  `protobuf:"varint,46,opt,name=memory_usage,json=memoryUsage,proto3" json:"memory_usage,omitempty"`                                           // provides feedback on memory usage within the group, depending on feedback mode it might be memory usage peak in group or memory usage in the first non-accepted test
	Runs           []*Submission_Run       `protobuf:"bytes,100,rep,name=runs,proto3" json:"runs,omitempty"`                                                                            // runs of the group
}

func (x *Submission_Group) Reset() {
	*x = Submission_Group{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_judge_submission_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Submission_Group) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Submission_Group) ProtoMessage() {}

func (x *Submission_Group) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_judge_submission_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
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

func (x *Submission_Group) GetStatus() Submission_Group_Status {
	if x != nil {
		return x.Status
	}
	return Submission_Group_UNKNOWN
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
	0x0c, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x6a, 0x75, 0x64, 0x67, 0x65, 0x1a, 0x23, 0x65,
	0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x61, 0x74, 0x6c, 0x61, 0x73, 0x2f, 0x74, 0x65, 0x73, 0x74,
	0x69, 0x6e, 0x67, 0x5f, 0x66, 0x65, 0x65, 0x64, 0x62, 0x61, 0x63, 0x6b, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x22, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x61, 0x74, 0x6c, 0x61, 0x73,
	0x2f, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x73, 0x63, 0x6f, 0x72, 0x69, 0x6e, 0x67,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x93, 0x0e, 0x0a, 0x0a, 0x53, 0x75, 0x62, 0x6d,
	0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x11, 0x0a, 0x03, 0x65, 0x72, 0x6e, 0x18, 0x8f, 0x4e,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x65, 0x72, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6f, 0x6e,
	0x74, 0x65, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x73, 0x74, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x62,
	0x6c, 0x65, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72,
	0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x49, 0x64, 0x12, 0x25, 0x0a, 0x0e, 0x70, 0x61, 0x72, 0x74, 0x69,
	0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0d, 0x70, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x3d,
	0x0a, 0x0c, 0x73, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x0b, 0x73, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x18, 0x0a,
	0x07, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07,
	0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x61, 0x6e, 0x67, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6c, 0x61, 0x6e, 0x67, 0x12, 0x16, 0x0a, 0x06, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65,
	0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72,
	0x65, 0x12, 0x37, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x14, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x1f, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x6a, 0x75, 0x64, 0x67, 0x65,
	0x2e, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x18, 0x15, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x73, 0x74, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x02, 0x52, 0x04,
	0x63, 0x6f, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x1f, 0x20,
	0x01, 0x28, 0x02, 0x52, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x65,
	0x72, 0x63, 0x65, 0x6e, 0x74, 0x61, 0x67, 0x65, 0x18, 0x20, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0a,
	0x70, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x61, 0x67, 0x65, 0x12, 0x36, 0x0a, 0x06, 0x67, 0x72,
	0x6f, 0x75, 0x70, 0x73, 0x18, 0x32, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x65, 0x6f, 0x6c,
	0x79, 0x6d, 0x70, 0x2e, 0x6a, 0x75, 0x64, 0x67, 0x65, 0x2e, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x06, 0x67, 0x72, 0x6f, 0x75,
	0x70, 0x73, 0x1a, 0xe7, 0x03, 0x0a, 0x03, 0x52, 0x75, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x26, 0x0a, 0x0f, 0x77, 0x61,
	0x6c, 0x6c, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x75, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x0d, 0x77, 0x61, 0x6c, 0x6c, 0x54, 0x69, 0x6d, 0x65, 0x55, 0x73, 0x61,
	0x67, 0x65, 0x12, 0x24, 0x0a, 0x0e, 0x63, 0x70, 0x75, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x75,
	0x73, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c, 0x63, 0x70, 0x75, 0x54,
	0x69, 0x6d, 0x65, 0x55, 0x73, 0x61, 0x67, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x6d, 0x65, 0x6d, 0x6f,
	0x72, 0x79, 0x5f, 0x75, 0x73, 0x61, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0b,
	0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x55, 0x73, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x69,
	0x6e, 0x64, 0x65, 0x78, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x69, 0x6e, 0x64, 0x65,
	0x78, 0x12, 0x17, 0x0a, 0x07, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x0b, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x74, 0x65, 0x73, 0x74, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f,
	0x73, 0x74, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x02, 0x52, 0x04, 0x63, 0x6f, 0x73, 0x74, 0x12, 0x14,
	0x0a, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x73,
	0x63, 0x6f, 0x72, 0x65, 0x12, 0x3b, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x14,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x23, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x6a, 0x75,
	0x64, 0x67, 0x65, 0x2e, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x52,
	0x75, 0x6e, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x22, 0xc8, 0x01, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x08, 0x0a, 0x04,
	0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x43, 0x52, 0x45, 0x41, 0x54, 0x45,
	0x44, 0x10, 0x01, 0x12, 0x0d, 0x0a, 0x09, 0x45, 0x58, 0x45, 0x43, 0x55, 0x54, 0x49, 0x4e, 0x47,
	0x10, 0x02, 0x12, 0x0b, 0x0a, 0x07, 0x54, 0x49, 0x4d, 0x45, 0x4f, 0x55, 0x54, 0x10, 0x03, 0x12,
	0x11, 0x0a, 0x0d, 0x43, 0x50, 0x55, 0x5f, 0x45, 0x58, 0x48, 0x41, 0x55, 0x53, 0x54, 0x45, 0x44,
	0x10, 0x04, 0x12, 0x13, 0x0a, 0x0f, 0x4d, 0x45, 0x4d, 0x4f, 0x52, 0x59, 0x5f, 0x4f, 0x56, 0x45,
	0x52, 0x46, 0x4c, 0x4f, 0x57, 0x10, 0x05, 0x12, 0x10, 0x0a, 0x0c, 0x57, 0x52, 0x4f, 0x4e, 0x47,
	0x5f, 0x41, 0x4e, 0x53, 0x57, 0x45, 0x52, 0x10, 0x06, 0x12, 0x0c, 0x0a, 0x08, 0x41, 0x43, 0x43,
	0x45, 0x50, 0x54, 0x45, 0x44, 0x10, 0x07, 0x12, 0x11, 0x0a, 0x0d, 0x52, 0x55, 0x4e, 0x54, 0x49,
	0x4d, 0x45, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x08, 0x12, 0x16, 0x0a, 0x12, 0x56, 0x45,
	0x52, 0x49, 0x46, 0x49, 0x43, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52,
	0x10, 0x09, 0x12, 0x0b, 0x0a, 0x07, 0x53, 0x4b, 0x49, 0x50, 0x50, 0x45, 0x44, 0x10, 0x0a, 0x12,
	0x0b, 0x0a, 0x07, 0x42, 0x4c, 0x4f, 0x43, 0x4b, 0x45, 0x44, 0x10, 0x0b, 0x1a, 0xbd, 0x05, 0x0a,
	0x05, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x1d, 0x0a, 0x0a,
	0x74, 0x65, 0x73, 0x74, 0x73, 0x65, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x74, 0x65, 0x73, 0x74, 0x73, 0x65, 0x74, 0x49, 0x64, 0x12, 0x3d, 0x0a, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x25, 0x2e, 0x65, 0x6f,
	0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x6a, 0x75, 0x64, 0x67, 0x65, 0x2e, 0x53, 0x75, 0x62, 0x6d, 0x69,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x22, 0x0a, 0x0c, 0x64, 0x65,
	0x70, 0x65, 0x6e, 0x64, 0x65, 0x6e, 0x63, 0x69, 0x65, 0x73, 0x18, 0x0b, 0x20, 0x03, 0x28, 0x0d,
	0x52, 0x0c, 0x64, 0x65, 0x70, 0x65, 0x6e, 0x64, 0x65, 0x6e, 0x63, 0x69, 0x65, 0x73, 0x12, 0x12,
	0x0a, 0x04, 0x63, 0x6f, 0x73, 0x74, 0x18, 0x14, 0x20, 0x01, 0x28, 0x02, 0x52, 0x04, 0x63, 0x6f,
	0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x15, 0x20, 0x01, 0x28,
	0x02, 0x52, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x12, 0x3c, 0x0a, 0x0c, 0x73, 0x63, 0x6f, 0x72,
	0x69, 0x6e, 0x67, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x18, 0x16, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x19,
	0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x61, 0x74, 0x6c, 0x61, 0x73, 0x2e, 0x53, 0x63,
	0x6f, 0x72, 0x69, 0x6e, 0x67, 0x4d, 0x6f, 0x64, 0x65, 0x52, 0x0b, 0x73, 0x63, 0x6f, 0x72, 0x69,
	0x6e, 0x67, 0x4d, 0x6f, 0x64, 0x65, 0x12, 0x45, 0x0a, 0x0f, 0x66, 0x65, 0x65, 0x64, 0x62, 0x61,
	0x63, 0x6b, 0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x1c, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x61, 0x74, 0x6c, 0x61, 0x73, 0x2e, 0x46,
	0x65, 0x65, 0x64, 0x62, 0x61, 0x63, 0x6b, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x0e, 0x66,
	0x65, 0x65, 0x64, 0x62, 0x61, 0x63, 0x6b, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0x26, 0x0a,
	0x0f, 0x77, 0x61, 0x6c, 0x6c, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x75, 0x73, 0x61, 0x67, 0x65,
	0x18, 0x29, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0d, 0x77, 0x61, 0x6c, 0x6c, 0x54, 0x69, 0x6d, 0x65,
	0x55, 0x73, 0x61, 0x67, 0x65, 0x12, 0x24, 0x0a, 0x0e, 0x63, 0x70, 0x75, 0x5f, 0x74, 0x69, 0x6d,
	0x65, 0x5f, 0x75, 0x73, 0x61, 0x67, 0x65, 0x18, 0x2a, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c, 0x63,
	0x70, 0x75, 0x54, 0x69, 0x6d, 0x65, 0x55, 0x73, 0x61, 0x67, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x6d,
	0x65, 0x6d, 0x6f, 0x72, 0x79, 0x5f, 0x75, 0x73, 0x61, 0x67, 0x65, 0x18, 0x2e, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x0b, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x55, 0x73, 0x61, 0x67, 0x65, 0x12, 0x30,
	0x0a, 0x04, 0x72, 0x75, 0x6e, 0x73, 0x18, 0x64, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x65,
	0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x6a, 0x75, 0x64, 0x67, 0x65, 0x2e, 0x53, 0x75, 0x62, 0x6d,
	0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x52, 0x75, 0x6e, 0x52, 0x04, 0x72, 0x75, 0x6e, 0x73,
	0x22, 0xc9, 0x01, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0b, 0x0a, 0x07, 0x55,
	0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x42, 0x4c, 0x4f, 0x43,
	0x4b, 0x45, 0x44, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x50, 0x45, 0x4e, 0x44, 0x49, 0x4e, 0x47,
	0x10, 0x02, 0x12, 0x0b, 0x0a, 0x07, 0x54, 0x45, 0x53, 0x54, 0x49, 0x4e, 0x47, 0x10, 0x03, 0x12,
	0x0c, 0x0a, 0x08, 0x41, 0x43, 0x43, 0x45, 0x50, 0x54, 0x45, 0x44, 0x10, 0x04, 0x12, 0x0b, 0x0a,
	0x07, 0x54, 0x49, 0x4d, 0x45, 0x4f, 0x55, 0x54, 0x10, 0x05, 0x12, 0x11, 0x0a, 0x0d, 0x43, 0x50,
	0x55, 0x5f, 0x45, 0x58, 0x48, 0x41, 0x55, 0x53, 0x54, 0x45, 0x44, 0x10, 0x06, 0x12, 0x13, 0x0a,
	0x0f, 0x4d, 0x45, 0x4d, 0x4f, 0x52, 0x59, 0x5f, 0x4f, 0x56, 0x45, 0x52, 0x46, 0x4c, 0x4f, 0x57,
	0x10, 0x07, 0x12, 0x10, 0x0a, 0x0c, 0x57, 0x52, 0x4f, 0x4e, 0x47, 0x5f, 0x41, 0x4e, 0x53, 0x57,
	0x45, 0x52, 0x10, 0x08, 0x12, 0x11, 0x0a, 0x0d, 0x52, 0x55, 0x4e, 0x54, 0x49, 0x4d, 0x45, 0x5f,
	0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x09, 0x12, 0x16, 0x0a, 0x12, 0x56, 0x45, 0x52, 0x49, 0x46,
	0x49, 0x43, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x0a, 0x12,
	0x0b, 0x0a, 0x07, 0x53, 0x4b, 0x49, 0x50, 0x50, 0x45, 0x44, 0x10, 0x0b, 0x22, 0x5f, 0x0a, 0x06,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00,
	0x12, 0x0b, 0x0a, 0x07, 0x50, 0x45, 0x4e, 0x44, 0x49, 0x4e, 0x47, 0x10, 0x01, 0x12, 0x0b, 0x0a,
	0x07, 0x54, 0x45, 0x53, 0x54, 0x49, 0x4e, 0x47, 0x10, 0x02, 0x12, 0x0b, 0x0a, 0x07, 0x54, 0x49,
	0x4d, 0x45, 0x4f, 0x55, 0x54, 0x10, 0x03, 0x12, 0x0c, 0x0a, 0x08, 0x43, 0x4f, 0x4d, 0x50, 0x4c,
	0x45, 0x54, 0x45, 0x10, 0x04, 0x12, 0x09, 0x0a, 0x05, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x05,
	0x12, 0x0b, 0x0a, 0x07, 0x46, 0x41, 0x49, 0x4c, 0x55, 0x52, 0x45, 0x10, 0x06, 0x42, 0x2d, 0x5a,
	0x2b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6f, 0x6c, 0x79,
	0x6d, 0x70, 0x2f, 0x67, 0x6f, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70,
	0x2f, 0x6a, 0x75, 0x64, 0x67, 0x65, 0x3b, 0x6a, 0x75, 0x64, 0x67, 0x65, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
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

var file_eolymp_judge_submission_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_eolymp_judge_submission_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_eolymp_judge_submission_proto_goTypes = []interface{}{
	(Submission_Status)(0),        // 0: eolymp.judge.Submission.Status
	(Submission_Run_Status)(0),    // 1: eolymp.judge.Submission.Run.Status
	(Submission_Group_Status)(0),  // 2: eolymp.judge.Submission.Group.Status
	(*Submission)(nil),            // 3: eolymp.judge.Submission
	(*Submission_Run)(nil),        // 4: eolymp.judge.Submission.Run
	(*Submission_Group)(nil),      // 5: eolymp.judge.Submission.Group
	(*timestamppb.Timestamp)(nil), // 6: google.protobuf.Timestamp
	(atlas.ScoringMode)(0),        // 7: eolymp.atlas.ScoringMode
	(atlas.FeedbackPolicy)(0),     // 8: eolymp.atlas.FeedbackPolicy
}
var file_eolymp_judge_submission_proto_depIdxs = []int32{
	6, // 0: eolymp.judge.Submission.submitted_at:type_name -> google.protobuf.Timestamp
	0, // 1: eolymp.judge.Submission.status:type_name -> eolymp.judge.Submission.Status
	5, // 2: eolymp.judge.Submission.groups:type_name -> eolymp.judge.Submission.Group
	1, // 3: eolymp.judge.Submission.Run.status:type_name -> eolymp.judge.Submission.Run.Status
	2, // 4: eolymp.judge.Submission.Group.status:type_name -> eolymp.judge.Submission.Group.Status
	7, // 5: eolymp.judge.Submission.Group.scoring_mode:type_name -> eolymp.atlas.ScoringMode
	8, // 6: eolymp.judge.Submission.Group.feedback_policy:type_name -> eolymp.atlas.FeedbackPolicy
	4, // 7: eolymp.judge.Submission.Group.runs:type_name -> eolymp.judge.Submission.Run
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_eolymp_judge_submission_proto_init() }
func file_eolymp_judge_submission_proto_init() {
	if File_eolymp_judge_submission_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_eolymp_judge_submission_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Submission); i {
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
		file_eolymp_judge_submission_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Submission_Run); i {
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
		file_eolymp_judge_submission_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Submission_Group); i {
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
			RawDescriptor: file_eolymp_judge_submission_proto_rawDesc,
			NumEnums:      3,
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
