// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.4
// 	protoc        v4.24.4
// source: eolymp/executor/evaluation_report.proto

package executor

import (
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

type EvaluationReport_Status int32

const (
	EvaluationReport_UNKNOWN_STATUS EvaluationReport_Status = 0 // should not be used
	EvaluationReport_PENDING        EvaluationReport_Status = 1 // not yet picked up by agent
	EvaluationReport_PROVISIONING   EvaluationReport_Status = 2 // agent is provisioning the environment
	EvaluationReport_INITIALIZING   EvaluationReport_Status = 3 // agent is initializing the task
	EvaluationReport_EXECUTING      EvaluationReport_Status = 4 // agent is evaluating the task
	EvaluationReport_COMPLETE       EvaluationReport_Status = 5 // evaluation is complete
	EvaluationReport_ERROR          EvaluationReport_Status = 6 // evaluation failed due to an error in the task
	EvaluationReport_FAILED         EvaluationReport_Status = 7 // evaluation failed due to an internal error
)

// Enum value maps for EvaluationReport_Status.
var (
	EvaluationReport_Status_name = map[int32]string{
		0: "UNKNOWN_STATUS",
		1: "PENDING",
		2: "PROVISIONING",
		3: "INITIALIZING",
		4: "EXECUTING",
		5: "COMPLETE",
		6: "ERROR",
		7: "FAILED",
	}
	EvaluationReport_Status_value = map[string]int32{
		"UNKNOWN_STATUS": 0,
		"PENDING":        1,
		"PROVISIONING":   2,
		"INITIALIZING":   3,
		"EXECUTING":      4,
		"COMPLETE":       5,
		"ERROR":          6,
		"FAILED":         7,
	}
)

func (x EvaluationReport_Status) Enum() *EvaluationReport_Status {
	p := new(EvaluationReport_Status)
	*p = x
	return p
}

func (x EvaluationReport_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EvaluationReport_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_eolymp_executor_evaluation_report_proto_enumTypes[0].Descriptor()
}

func (EvaluationReport_Status) Type() protoreflect.EnumType {
	return &file_eolymp_executor_evaluation_report_proto_enumTypes[0]
}

func (x EvaluationReport_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use EvaluationReport_Status.Descriptor instead.
func (EvaluationReport_Status) EnumDescriptor() ([]byte, []int) {
	return file_eolymp_executor_evaluation_report_proto_rawDescGZIP(), []int{0, 0}
}

// Type of the report
type EvaluationReport_Type int32

const (
	// Report without type should be ignored.
	EvaluationReport_UNKNOWN_TYPE EvaluationReport_Type = 0
	// Update report contains information about runs, inspect runs to see their status.
	// Update report might be partial in case task was split among multiple agents.
	EvaluationReport_TYPE_UPDATE EvaluationReport_Type = 1
	// Error report means agent has encountered an error and can't perform runs.
	// Errors are always related to the source code of the task, such as compilation or linting error. Detailed
	// description of the error can be found in error field (ie. output of the compiler or linter).
	EvaluationReport_TYPE_ERROR EvaluationReport_Type = 2
	// Failure report means agent has encountered an failure and can't perform runs.
	// Failure, unlike error, happens due to a problem with executor agent itself or task configuration, such as
	// invalid object id references or checker source does not compile. Detailed description of the failure can be
	// found in the failure field.
	EvaluationReport_TYPE_FAILURE EvaluationReport_Type = 3
)

// Enum value maps for EvaluationReport_Type.
var (
	EvaluationReport_Type_name = map[int32]string{
		0: "UNKNOWN_TYPE",
		1: "TYPE_UPDATE",
		2: "TYPE_ERROR",
		3: "TYPE_FAILURE",
	}
	EvaluationReport_Type_value = map[string]int32{
		"UNKNOWN_TYPE": 0,
		"TYPE_UPDATE":  1,
		"TYPE_ERROR":   2,
		"TYPE_FAILURE": 3,
	}
)

func (x EvaluationReport_Type) Enum() *EvaluationReport_Type {
	p := new(EvaluationReport_Type)
	*p = x
	return p
}

func (x EvaluationReport_Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EvaluationReport_Type) Descriptor() protoreflect.EnumDescriptor {
	return file_eolymp_executor_evaluation_report_proto_enumTypes[1].Descriptor()
}

func (EvaluationReport_Type) Type() protoreflect.EnumType {
	return &file_eolymp_executor_evaluation_report_proto_enumTypes[1]
}

func (x EvaluationReport_Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use EvaluationReport_Type.Descriptor instead.
func (EvaluationReport_Type) EnumDescriptor() ([]byte, []int) {
	return file_eolymp_executor_evaluation_report_proto_rawDescGZIP(), []int{0, 1}
}

type EvaluationReport_Run_Status int32

const (
	EvaluationReport_Run_NONE                 EvaluationReport_Run_Status = 0  // should not be used
	EvaluationReport_Run_PENDING              EvaluationReport_Run_Status = 1  // pending to be executed
	EvaluationReport_Run_EXECUTING            EvaluationReport_Run_Status = 2  // executing (see output and stderr for partial data)
	EvaluationReport_Run_TIMEOUT              EvaluationReport_Run_Status = 3  // timeout reached (wall time usage reached)
	EvaluationReport_Run_CPU_EXHAUSTED        EvaluationReport_Run_Status = 4  // cpu exhausted (cpu time usage reached)
	EvaluationReport_Run_MEMORY_OVERFLOW      EvaluationReport_Run_Status = 5  // memory limit reached
	EvaluationReport_Run_RUNTIME_ERROR        EvaluationReport_Run_Status = 6  // executed, but exit code was non-zero
	EvaluationReport_Run_EXECUTED             EvaluationReport_Run_Status = 7  // executed (it is a final status for tasks without checker)
	EvaluationReport_Run_ACCEPTED             EvaluationReport_Run_Status = 8  // executed and output matched answer
	EvaluationReport_Run_WRONG_ANSWER         EvaluationReport_Run_Status = 9  // executed, but output didn't match answer
	EvaluationReport_Run_VERIFICATION_FAILURE EvaluationReport_Run_Status = 10 // executed, but checker returned an error during execution (use checker_log to learn more about failure)
	EvaluationReport_Run_SKIPPED              EvaluationReport_Run_Status = 11 // run won't be executed due to preconditions (overall timeout, block etc)
	EvaluationReport_Run_INTERACTION_FAILURE  EvaluationReport_Run_Status = 12 // interactor failed (TL-ed, ML-ed, got a runtime error or claimed that the jury had a wrong answer), check interactor_log
)

// Enum value maps for EvaluationReport_Run_Status.
var (
	EvaluationReport_Run_Status_name = map[int32]string{
		0:  "NONE",
		1:  "PENDING",
		2:  "EXECUTING",
		3:  "TIMEOUT",
		4:  "CPU_EXHAUSTED",
		5:  "MEMORY_OVERFLOW",
		6:  "RUNTIME_ERROR",
		7:  "EXECUTED",
		8:  "ACCEPTED",
		9:  "WRONG_ANSWER",
		10: "VERIFICATION_FAILURE",
		11: "SKIPPED",
		12: "INTERACTION_FAILURE",
	}
	EvaluationReport_Run_Status_value = map[string]int32{
		"NONE":                 0,
		"PENDING":              1,
		"EXECUTING":            2,
		"TIMEOUT":              3,
		"CPU_EXHAUSTED":        4,
		"MEMORY_OVERFLOW":      5,
		"RUNTIME_ERROR":        6,
		"EXECUTED":             7,
		"ACCEPTED":             8,
		"WRONG_ANSWER":         9,
		"VERIFICATION_FAILURE": 10,
		"SKIPPED":              11,
		"INTERACTION_FAILURE":  12,
	}
)

func (x EvaluationReport_Run_Status) Enum() *EvaluationReport_Run_Status {
	p := new(EvaluationReport_Run_Status)
	*p = x
	return p
}

func (x EvaluationReport_Run_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EvaluationReport_Run_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_eolymp_executor_evaluation_report_proto_enumTypes[2].Descriptor()
}

func (EvaluationReport_Run_Status) Type() protoreflect.EnumType {
	return &file_eolymp_executor_evaluation_report_proto_enumTypes[2]
}

func (x EvaluationReport_Run_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use EvaluationReport_Run_Status.Descriptor instead.
func (EvaluationReport_Run_Status) EnumDescriptor() ([]byte, []int) {
	return file_eolymp_executor_evaluation_report_proto_rawDescGZIP(), []int{0, 0, 0}
}

// Status represents results of the executing task.
type EvaluationReport struct {
	state     protoimpl.MessageState `protogen:"open.v1"`
	TaskId    string                 `protobuf:"bytes,1,opt,name=task_id,json=taskId,proto3" json:"task_id,omitempty"`
	Reference string                 `protobuf:"bytes,2,opt,name=reference,proto3" json:"reference,omitempty"`
	Origin    string                 `protobuf:"bytes,3,opt,name=origin,proto3" json:"origin,omitempty"`
	Agent     string                 `protobuf:"bytes,4,opt,name=agent,proto3" json:"agent,omitempty"`
	// Source code signature is a unique fingerprint of the code, calculated by agent for a specific language.
	// It should be used to find identical or similar tasks. Signature can be calculated only for some languages.
	Signature string `protobuf:"bytes,50,opt,name=signature,proto3" json:"signature,omitempty"`
	// Always increasing report version.
	//
	// Each time agent emits a report it would increase version, so listener can put reports in the right order: process
	// newer and ignore older.
	//
	// In case runs of a single task are distributed among multiple agents, each agent will report version independently,
	// so listener must track versions per run (eg. run #1 last update is v.15, run #2 last update is v.41, if listener
	// receives run #1 v.20 it's newer and should be processed, but run #2 v.20 should be ignored).
	Version uint32 `protobuf:"varint,100,opt,name=version,proto3" json:"version,omitempty"`
	// Report type, see explanation to Type enumeration
	// deprecated: use status instead
	Type EvaluationReport_Type `protobuf:"varint,10,opt,name=type,proto3,enum=eolymp.executor.EvaluationReport_Type" json:"type,omitempty"`
	// Status of the evaluation
	Status EvaluationReport_Status `protobuf:"varint,11,opt,name=status,proto3,enum=eolymp.executor.EvaluationReport_Status" json:"status,omitempty"`
	// Error message
	ErrorMessage  string                  `protobuf:"bytes,20,opt,name=error_message,json=errorMessage,proto3" json:"error_message,omitempty"`
	Runs          []*EvaluationReport_Run `protobuf:"bytes,40,rep,name=runs,proto3" json:"runs,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *EvaluationReport) Reset() {
	*x = EvaluationReport{}
	mi := &file_eolymp_executor_evaluation_report_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *EvaluationReport) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EvaluationReport) ProtoMessage() {}

func (x *EvaluationReport) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_executor_evaluation_report_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EvaluationReport.ProtoReflect.Descriptor instead.
func (*EvaluationReport) Descriptor() ([]byte, []int) {
	return file_eolymp_executor_evaluation_report_proto_rawDescGZIP(), []int{0}
}

func (x *EvaluationReport) GetTaskId() string {
	if x != nil {
		return x.TaskId
	}
	return ""
}

func (x *EvaluationReport) GetReference() string {
	if x != nil {
		return x.Reference
	}
	return ""
}

func (x *EvaluationReport) GetOrigin() string {
	if x != nil {
		return x.Origin
	}
	return ""
}

func (x *EvaluationReport) GetAgent() string {
	if x != nil {
		return x.Agent
	}
	return ""
}

func (x *EvaluationReport) GetSignature() string {
	if x != nil {
		return x.Signature
	}
	return ""
}

func (x *EvaluationReport) GetVersion() uint32 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *EvaluationReport) GetType() EvaluationReport_Type {
	if x != nil {
		return x.Type
	}
	return EvaluationReport_UNKNOWN_TYPE
}

func (x *EvaluationReport) GetStatus() EvaluationReport_Status {
	if x != nil {
		return x.Status
	}
	return EvaluationReport_UNKNOWN_STATUS
}

func (x *EvaluationReport) GetErrorMessage() string {
	if x != nil {
		return x.ErrorMessage
	}
	return ""
}

func (x *EvaluationReport) GetRuns() []*EvaluationReport_Run {
	if x != nil {
		return x.Runs
	}
	return nil
}

// Run represents a single execution
type EvaluationReport_Run struct {
	state           protoimpl.MessageState      `protogen:"open.v1"`
	Reference       string                      `protobuf:"bytes,1,opt,name=reference,proto3" json:"reference,omitempty"`                                             // run reference as set by originator
	Status          EvaluationReport_Run_Status `protobuf:"varint,2,opt,name=status,proto3,enum=eolymp.executor.EvaluationReport_Run_Status" json:"status,omitempty"` // run status
	Score           float32                     `protobuf:"fixed32,80,opt,name=score,proto3" json:"score,omitempty"`                                                  // score, number of points awarded by checker
	Cost            float32                     `protobuf:"fixed32,81,opt,name=cost,proto3" json:"cost,omitempty"`                                                    // cost, maximum number of points awarded for the test (as defined in the task)
	WallTimeUsage   uint32                      `protobuf:"varint,51,opt,name=wall_time_usage,json=wallTimeUsage,proto3" json:"wall_time_usage,omitempty"`
	WallTimeLimit   uint32                      `protobuf:"varint,61,opt,name=wall_time_limit,json=wallTimeLimit,proto3" json:"wall_time_limit,omitempty"`
	CpuTimeUsage    uint32                      `protobuf:"varint,52,opt,name=cpu_time_usage,json=cpuTimeUsage,proto3" json:"cpu_time_usage,omitempty"`
	CpuTimeLimit    uint32                      `protobuf:"varint,62,opt,name=cpu_time_limit,json=cpuTimeLimit,proto3" json:"cpu_time_limit,omitempty"`
	MemoryUsage     uint64                      `protobuf:"varint,53,opt,name=memory_usage,json=memoryUsage,proto3" json:"memory_usage,omitempty"`
	MemoryLimit     uint64                      `protobuf:"varint,63,opt,name=memory_limit,json=memoryLimit,proto3" json:"memory_limit,omitempty"`
	InputUrl        string                      `protobuf:"bytes,10,opt,name=input_url,json=inputUrl,proto3" json:"input_url,omitempty"`
	OutputUrl       string                      `protobuf:"bytes,11,opt,name=output_url,json=outputUrl,proto3" json:"output_url,omitempty"`
	AnswerUrl       string                      `protobuf:"bytes,12,opt,name=answer_url,json=answerUrl,proto3" json:"answer_url,omitempty"`
	DebugStats      *Stats                      `protobuf:"bytes,90,opt,name=debug_stats,json=debugStats,proto3" json:"debug_stats,omitempty"`                // execution stats
	CheckerStats    *Stats                      `protobuf:"bytes,35,opt,name=checker_stats,json=checkerStats,proto3" json:"checker_stats,omitempty"`          // checker stats
	InteractorStats *Stats                      `protobuf:"bytes,45,opt,name=interactor_stats,json=interactorStats,proto3" json:"interactor_stats,omitempty"` // interactor stats
	unknownFields   protoimpl.UnknownFields
	sizeCache       protoimpl.SizeCache
}

func (x *EvaluationReport_Run) Reset() {
	*x = EvaluationReport_Run{}
	mi := &file_eolymp_executor_evaluation_report_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *EvaluationReport_Run) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EvaluationReport_Run) ProtoMessage() {}

func (x *EvaluationReport_Run) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_executor_evaluation_report_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EvaluationReport_Run.ProtoReflect.Descriptor instead.
func (*EvaluationReport_Run) Descriptor() ([]byte, []int) {
	return file_eolymp_executor_evaluation_report_proto_rawDescGZIP(), []int{0, 0}
}

func (x *EvaluationReport_Run) GetReference() string {
	if x != nil {
		return x.Reference
	}
	return ""
}

func (x *EvaluationReport_Run) GetStatus() EvaluationReport_Run_Status {
	if x != nil {
		return x.Status
	}
	return EvaluationReport_Run_NONE
}

func (x *EvaluationReport_Run) GetScore() float32 {
	if x != nil {
		return x.Score
	}
	return 0
}

func (x *EvaluationReport_Run) GetCost() float32 {
	if x != nil {
		return x.Cost
	}
	return 0
}

func (x *EvaluationReport_Run) GetWallTimeUsage() uint32 {
	if x != nil {
		return x.WallTimeUsage
	}
	return 0
}

func (x *EvaluationReport_Run) GetWallTimeLimit() uint32 {
	if x != nil {
		return x.WallTimeLimit
	}
	return 0
}

func (x *EvaluationReport_Run) GetCpuTimeUsage() uint32 {
	if x != nil {
		return x.CpuTimeUsage
	}
	return 0
}

func (x *EvaluationReport_Run) GetCpuTimeLimit() uint32 {
	if x != nil {
		return x.CpuTimeLimit
	}
	return 0
}

func (x *EvaluationReport_Run) GetMemoryUsage() uint64 {
	if x != nil {
		return x.MemoryUsage
	}
	return 0
}

func (x *EvaluationReport_Run) GetMemoryLimit() uint64 {
	if x != nil {
		return x.MemoryLimit
	}
	return 0
}

func (x *EvaluationReport_Run) GetInputUrl() string {
	if x != nil {
		return x.InputUrl
	}
	return ""
}

func (x *EvaluationReport_Run) GetOutputUrl() string {
	if x != nil {
		return x.OutputUrl
	}
	return ""
}

func (x *EvaluationReport_Run) GetAnswerUrl() string {
	if x != nil {
		return x.AnswerUrl
	}
	return ""
}

func (x *EvaluationReport_Run) GetDebugStats() *Stats {
	if x != nil {
		return x.DebugStats
	}
	return nil
}

func (x *EvaluationReport_Run) GetCheckerStats() *Stats {
	if x != nil {
		return x.CheckerStats
	}
	return nil
}

func (x *EvaluationReport_Run) GetInteractorStats() *Stats {
	if x != nil {
		return x.InteractorStats
	}
	return nil
}

var File_eolymp_executor_evaluation_report_proto protoreflect.FileDescriptor

var file_eolymp_executor_evaluation_report_proto_rawDesc = string([]byte{
	0x0a, 0x27, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x6f,
	0x72, 0x2f, 0x65, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x72, 0x65, 0x70,
	0x6f, 0x72, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x65, 0x6f, 0x6c, 0x79, 0x6d,
	0x70, 0x2e, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x6f, 0x72, 0x1a, 0x1b, 0x65, 0x6f, 0x6c, 0x79,
	0x6d, 0x70, 0x2f, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x6f, 0x72, 0x2f, 0x73, 0x74, 0x61, 0x74,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd1, 0x0b, 0x0a, 0x10, 0x45, 0x76, 0x61, 0x6c,
	0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x17, 0x0a, 0x07,
	0x74, 0x61, 0x73, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74,
	0x61, 0x73, 0x6b, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e,
	0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65,
	0x6e, 0x63, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x61,
	0x67, 0x65, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x61, 0x67, 0x65, 0x6e,
	0x74, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x32,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x64, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x3a, 0x0a, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x26, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70,
	0x2e, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x6f, 0x72, 0x2e, 0x45, 0x76, 0x61, 0x6c, 0x75, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x52,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x40, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x0b, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x28, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x65,
	0x78, 0x65, 0x63, 0x75, 0x74, 0x6f, 0x72, 0x2e, 0x45, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x23, 0x0a, 0x0d, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x39, 0x0a, 0x04,
	0x72, 0x75, 0x6e, 0x73, 0x18, 0x28, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x65, 0x6f, 0x6c,
	0x79, 0x6d, 0x70, 0x2e, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x6f, 0x72, 0x2e, 0x45, 0x76, 0x61,
	0x6c, 0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x52, 0x75,
	0x6e, 0x52, 0x04, 0x72, 0x75, 0x6e, 0x73, 0x1a, 0xf0, 0x06, 0x0a, 0x03, 0x52, 0x75, 0x6e, 0x12,
	0x1c, 0x0a, 0x09, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x44, 0x0a,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2c, 0x2e,
	0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x6f, 0x72, 0x2e,
	0x45, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74,
	0x2e, 0x52, 0x75, 0x6e, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x50, 0x20, 0x01,
	0x28, 0x02, 0x52, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x73,
	0x74, 0x18, 0x51, 0x20, 0x01, 0x28, 0x02, 0x52, 0x04, 0x63, 0x6f, 0x73, 0x74, 0x12, 0x26, 0x0a,
	0x0f, 0x77, 0x61, 0x6c, 0x6c, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x75, 0x73, 0x61, 0x67, 0x65,
	0x18, 0x33, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0d, 0x77, 0x61, 0x6c, 0x6c, 0x54, 0x69, 0x6d, 0x65,
	0x55, 0x73, 0x61, 0x67, 0x65, 0x12, 0x26, 0x0a, 0x0f, 0x77, 0x61, 0x6c, 0x6c, 0x5f, 0x74, 0x69,
	0x6d, 0x65, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x3d, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0d,
	0x77, 0x61, 0x6c, 0x6c, 0x54, 0x69, 0x6d, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x24, 0x0a,
	0x0e, 0x63, 0x70, 0x75, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x75, 0x73, 0x61, 0x67, 0x65, 0x18,
	0x34, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c, 0x63, 0x70, 0x75, 0x54, 0x69, 0x6d, 0x65, 0x55, 0x73,
	0x61, 0x67, 0x65, 0x12, 0x24, 0x0a, 0x0e, 0x63, 0x70, 0x75, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x5f,
	0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x3e, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c, 0x63, 0x70, 0x75,
	0x54, 0x69, 0x6d, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x6d, 0x65, 0x6d,
	0x6f, 0x72, 0x79, 0x5f, 0x75, 0x73, 0x61, 0x67, 0x65, 0x18, 0x35, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x0b, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x55, 0x73, 0x61, 0x67, 0x65, 0x12, 0x21, 0x0a, 0x0c,
	0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x3f, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x0b, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x12,
	0x1b, 0x0a, 0x09, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x55, 0x72, 0x6c, 0x12, 0x1d, 0x0a, 0x0a,
	0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x55, 0x72, 0x6c, 0x12, 0x1d, 0x0a, 0x0a, 0x61,
	0x6e, 0x73, 0x77, 0x65, 0x72, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x61, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x55, 0x72, 0x6c, 0x12, 0x37, 0x0a, 0x0b, 0x64, 0x65,
	0x62, 0x75, 0x67, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x73, 0x18, 0x5a, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x16, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x6f,
	0x72, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x73, 0x52, 0x0a, 0x64, 0x65, 0x62, 0x75, 0x67, 0x53, 0x74,
	0x61, 0x74, 0x73, 0x12, 0x3b, 0x0a, 0x0d, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x65, 0x72, 0x5f, 0x73,
	0x74, 0x61, 0x74, 0x73, 0x18, 0x23, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x65, 0x6f, 0x6c,
	0x79, 0x6d, 0x70, 0x2e, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x6f, 0x72, 0x2e, 0x53, 0x74, 0x61,
	0x74, 0x73, 0x52, 0x0c, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x73,
	0x12, 0x41, 0x0a, 0x10, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x5f, 0x73,
	0x74, 0x61, 0x74, 0x73, 0x18, 0x2d, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x65, 0x6f, 0x6c,
	0x79, 0x6d, 0x70, 0x2e, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x6f, 0x72, 0x2e, 0x53, 0x74, 0x61,
	0x74, 0x73, 0x52, 0x0f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x53, 0x74,
	0x61, 0x74, 0x73, 0x22, 0xe4, 0x01, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x08,
	0x0a, 0x04, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x50, 0x45, 0x4e, 0x44,
	0x49, 0x4e, 0x47, 0x10, 0x01, 0x12, 0x0d, 0x0a, 0x09, 0x45, 0x58, 0x45, 0x43, 0x55, 0x54, 0x49,
	0x4e, 0x47, 0x10, 0x02, 0x12, 0x0b, 0x0a, 0x07, 0x54, 0x49, 0x4d, 0x45, 0x4f, 0x55, 0x54, 0x10,
	0x03, 0x12, 0x11, 0x0a, 0x0d, 0x43, 0x50, 0x55, 0x5f, 0x45, 0x58, 0x48, 0x41, 0x55, 0x53, 0x54,
	0x45, 0x44, 0x10, 0x04, 0x12, 0x13, 0x0a, 0x0f, 0x4d, 0x45, 0x4d, 0x4f, 0x52, 0x59, 0x5f, 0x4f,
	0x56, 0x45, 0x52, 0x46, 0x4c, 0x4f, 0x57, 0x10, 0x05, 0x12, 0x11, 0x0a, 0x0d, 0x52, 0x55, 0x4e,
	0x54, 0x49, 0x4d, 0x45, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x06, 0x12, 0x0c, 0x0a, 0x08,
	0x45, 0x58, 0x45, 0x43, 0x55, 0x54, 0x45, 0x44, 0x10, 0x07, 0x12, 0x0c, 0x0a, 0x08, 0x41, 0x43,
	0x43, 0x45, 0x50, 0x54, 0x45, 0x44, 0x10, 0x08, 0x12, 0x10, 0x0a, 0x0c, 0x57, 0x52, 0x4f, 0x4e,
	0x47, 0x5f, 0x41, 0x4e, 0x53, 0x57, 0x45, 0x52, 0x10, 0x09, 0x12, 0x18, 0x0a, 0x14, 0x56, 0x45,
	0x52, 0x49, 0x46, 0x49, 0x43, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x46, 0x41, 0x49, 0x4c, 0x55,
	0x52, 0x45, 0x10, 0x0a, 0x12, 0x0b, 0x0a, 0x07, 0x53, 0x4b, 0x49, 0x50, 0x50, 0x45, 0x44, 0x10,
	0x0b, 0x12, 0x17, 0x0a, 0x13, 0x49, 0x4e, 0x54, 0x45, 0x52, 0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e,
	0x5f, 0x46, 0x41, 0x49, 0x4c, 0x55, 0x52, 0x45, 0x10, 0x0c, 0x22, 0x81, 0x01, 0x0a, 0x06, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x12, 0x0a, 0x0e, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e,
	0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x50, 0x45, 0x4e,
	0x44, 0x49, 0x4e, 0x47, 0x10, 0x01, 0x12, 0x10, 0x0a, 0x0c, 0x50, 0x52, 0x4f, 0x56, 0x49, 0x53,
	0x49, 0x4f, 0x4e, 0x49, 0x4e, 0x47, 0x10, 0x02, 0x12, 0x10, 0x0a, 0x0c, 0x49, 0x4e, 0x49, 0x54,
	0x49, 0x41, 0x4c, 0x49, 0x5a, 0x49, 0x4e, 0x47, 0x10, 0x03, 0x12, 0x0d, 0x0a, 0x09, 0x45, 0x58,
	0x45, 0x43, 0x55, 0x54, 0x49, 0x4e, 0x47, 0x10, 0x04, 0x12, 0x0c, 0x0a, 0x08, 0x43, 0x4f, 0x4d,
	0x50, 0x4c, 0x45, 0x54, 0x45, 0x10, 0x05, 0x12, 0x09, 0x0a, 0x05, 0x45, 0x52, 0x52, 0x4f, 0x52,
	0x10, 0x06, 0x12, 0x0a, 0x0a, 0x06, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x10, 0x07, 0x22, 0x4b,
	0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x10, 0x0a, 0x0c, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57,
	0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x10, 0x00, 0x12, 0x0f, 0x0a, 0x0b, 0x54, 0x59, 0x50, 0x45,
	0x5f, 0x55, 0x50, 0x44, 0x41, 0x54, 0x45, 0x10, 0x01, 0x12, 0x0e, 0x0a, 0x0a, 0x54, 0x59, 0x50,
	0x45, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x02, 0x12, 0x10, 0x0a, 0x0c, 0x54, 0x59, 0x50,
	0x45, 0x5f, 0x46, 0x41, 0x49, 0x4c, 0x55, 0x52, 0x45, 0x10, 0x03, 0x42, 0x33, 0x5a, 0x31, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70,
	0x2f, 0x67, 0x6f, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x65,
	0x78, 0x65, 0x63, 0x75, 0x74, 0x6f, 0x72, 0x3b, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x6f, 0x72,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_eolymp_executor_evaluation_report_proto_rawDescOnce sync.Once
	file_eolymp_executor_evaluation_report_proto_rawDescData []byte
)

func file_eolymp_executor_evaluation_report_proto_rawDescGZIP() []byte {
	file_eolymp_executor_evaluation_report_proto_rawDescOnce.Do(func() {
		file_eolymp_executor_evaluation_report_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_eolymp_executor_evaluation_report_proto_rawDesc), len(file_eolymp_executor_evaluation_report_proto_rawDesc)))
	})
	return file_eolymp_executor_evaluation_report_proto_rawDescData
}

var file_eolymp_executor_evaluation_report_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_eolymp_executor_evaluation_report_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_eolymp_executor_evaluation_report_proto_goTypes = []any{
	(EvaluationReport_Status)(0),     // 0: eolymp.executor.EvaluationReport.Status
	(EvaluationReport_Type)(0),       // 1: eolymp.executor.EvaluationReport.Type
	(EvaluationReport_Run_Status)(0), // 2: eolymp.executor.EvaluationReport.Run.Status
	(*EvaluationReport)(nil),         // 3: eolymp.executor.EvaluationReport
	(*EvaluationReport_Run)(nil),     // 4: eolymp.executor.EvaluationReport.Run
	(*Stats)(nil),                    // 5: eolymp.executor.Stats
}
var file_eolymp_executor_evaluation_report_proto_depIdxs = []int32{
	1, // 0: eolymp.executor.EvaluationReport.type:type_name -> eolymp.executor.EvaluationReport.Type
	0, // 1: eolymp.executor.EvaluationReport.status:type_name -> eolymp.executor.EvaluationReport.Status
	4, // 2: eolymp.executor.EvaluationReport.runs:type_name -> eolymp.executor.EvaluationReport.Run
	2, // 3: eolymp.executor.EvaluationReport.Run.status:type_name -> eolymp.executor.EvaluationReport.Run.Status
	5, // 4: eolymp.executor.EvaluationReport.Run.debug_stats:type_name -> eolymp.executor.Stats
	5, // 5: eolymp.executor.EvaluationReport.Run.checker_stats:type_name -> eolymp.executor.Stats
	5, // 6: eolymp.executor.EvaluationReport.Run.interactor_stats:type_name -> eolymp.executor.Stats
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_eolymp_executor_evaluation_report_proto_init() }
func file_eolymp_executor_evaluation_report_proto_init() {
	if File_eolymp_executor_evaluation_report_proto != nil {
		return
	}
	file_eolymp_executor_stats_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_eolymp_executor_evaluation_report_proto_rawDesc), len(file_eolymp_executor_evaluation_report_proto_rawDesc)),
			NumEnums:      3,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_eolymp_executor_evaluation_report_proto_goTypes,
		DependencyIndexes: file_eolymp_executor_evaluation_report_proto_depIdxs,
		EnumInfos:         file_eolymp_executor_evaluation_report_proto_enumTypes,
		MessageInfos:      file_eolymp_executor_evaluation_report_proto_msgTypes,
	}.Build()
	File_eolymp_executor_evaluation_report_proto = out.File
	file_eolymp_executor_evaluation_report_proto_goTypes = nil
	file_eolymp_executor_evaluation_report_proto_depIdxs = nil
}
