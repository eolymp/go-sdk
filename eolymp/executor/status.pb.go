// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.18.1
// source: eolymp/executor/status.proto

package executor

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

// Type of the report
type Status_Type int32

const (
	// Report without type should be ignored.
	Status_NONE Status_Type = 0
	// Update report contains information about runs, inspect runs to see their status.
	// Update report might be partial in case task was split among multiple agents.
	Status_UPDATE Status_Type = 1
	// Error report means agent has encountered an error and can't perform runs.
	// Errors are always related to the source code of the task, such as compilation or linting error. Detailed
	// description of the error can be found in error field (ie. output of the compiler or linter).
	Status_ERROR Status_Type = 2
	// Failure report means agent has encountered an failure and can't perform runs.
	// Failure, unlike error, happens due to a problem with executor agent itself or task configuration, such as
	// invalid object id references or verifier source does not compile. Detailed description of the failure can be
	// found in the failure field.
	Status_FAILURE Status_Type = 3
)

// Enum value maps for Status_Type.
var (
	Status_Type_name = map[int32]string{
		0: "NONE",
		1: "UPDATE",
		2: "ERROR",
		3: "FAILURE",
	}
	Status_Type_value = map[string]int32{
		"NONE":    0,
		"UPDATE":  1,
		"ERROR":   2,
		"FAILURE": 3,
	}
)

func (x Status_Type) Enum() *Status_Type {
	p := new(Status_Type)
	*p = x
	return p
}

func (x Status_Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Status_Type) Descriptor() protoreflect.EnumDescriptor {
	return file_eolymp_executor_status_proto_enumTypes[0].Descriptor()
}

func (Status_Type) Type() protoreflect.EnumType {
	return &file_eolymp_executor_status_proto_enumTypes[0]
}

func (x Status_Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Status_Type.Descriptor instead.
func (Status_Type) EnumDescriptor() ([]byte, []int) {
	return file_eolymp_executor_status_proto_rawDescGZIP(), []int{0, 0}
}

type Status_Run_Status int32

const (
	Status_Run_NONE                 Status_Run_Status = 0  // should not be used
	Status_Run_PENDING              Status_Run_Status = 1  // pending to be executed
	Status_Run_EXECUTING            Status_Run_Status = 2  // executing (see output and stderr for partial data)
	Status_Run_TIMEOUT              Status_Run_Status = 3  // timeout reached (wall time usage reached)
	Status_Run_CPU_EXHAUSTED        Status_Run_Status = 4  // cpu exhausted (cpu time usage reached)
	Status_Run_MEMORY_OVERFLOW      Status_Run_Status = 5  // memory limit reached
	Status_Run_RUNTIME_ERROR        Status_Run_Status = 6  // executed, but exit code was non-zero
	Status_Run_EXECUTED             Status_Run_Status = 7  // executed (it is a final status for tasks without verifier)
	Status_Run_ACCEPTED             Status_Run_Status = 8  // executed and output matched answer
	Status_Run_WRONG_ANSWER         Status_Run_Status = 9  // executed, but output didn't match answer
	Status_Run_VERIFICATION_FAILURE Status_Run_Status = 10 // executed, but verifier returned an error during execution (use verifier_log to learn more about failure)
	Status_Run_SKIPPED              Status_Run_Status = 11 // run won't be executed due to preconditions (overall timeout, block etc)
)

// Enum value maps for Status_Run_Status.
var (
	Status_Run_Status_name = map[int32]string{
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
	}
	Status_Run_Status_value = map[string]int32{
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
	}
)

func (x Status_Run_Status) Enum() *Status_Run_Status {
	p := new(Status_Run_Status)
	*p = x
	return p
}

func (x Status_Run_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Status_Run_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_eolymp_executor_status_proto_enumTypes[1].Descriptor()
}

func (Status_Run_Status) Type() protoreflect.EnumType {
	return &file_eolymp_executor_status_proto_enumTypes[1]
}

func (x Status_Run_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Status_Run_Status.Descriptor instead.
func (Status_Run_Status) EnumDescriptor() ([]byte, []int) {
	return file_eolymp_executor_status_proto_rawDescGZIP(), []int{0, 0, 0}
}

// Status represents results of the executing task.
type Status struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Task reference as set when task was created.
	Reference string `protobuf:"bytes,1,opt,name=reference,proto3" json:"reference,omitempty"`
	// Originator of the task (service which created task)
	Origin string `protobuf:"bytes,2,opt,name=origin,proto3" json:"origin,omitempty"`
	// Report type, see explanation to Type enumeration
	Type Status_Type `protobuf:"varint,10,opt,name=type,proto3,enum=eolymp.executor.Status_Type" json:"type,omitempty"`
	// Error message for ERROR report
	Error string `protobuf:"bytes,20,opt,name=error,proto3" json:"error,omitempty"`
	// Failure message for FAILURE report
	Failure string `protobuf:"bytes,30,opt,name=failure,proto3" json:"failure,omitempty"`
	// Runs for UPDATE report
	Runs []*Status_Run `protobuf:"bytes,40,rep,name=runs,proto3" json:"runs,omitempty"`
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
	// Unique identifier of the agent which executed task.
	Agent string `protobuf:"bytes,110,opt,name=agent,proto3" json:"agent,omitempty"`
}

func (x *Status) Reset() {
	*x = Status{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_executor_status_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Status) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Status) ProtoMessage() {}

func (x *Status) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_executor_status_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Status.ProtoReflect.Descriptor instead.
func (*Status) Descriptor() ([]byte, []int) {
	return file_eolymp_executor_status_proto_rawDescGZIP(), []int{0}
}

func (x *Status) GetReference() string {
	if x != nil {
		return x.Reference
	}
	return ""
}

func (x *Status) GetOrigin() string {
	if x != nil {
		return x.Origin
	}
	return ""
}

func (x *Status) GetType() Status_Type {
	if x != nil {
		return x.Type
	}
	return Status_NONE
}

func (x *Status) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

func (x *Status) GetFailure() string {
	if x != nil {
		return x.Failure
	}
	return ""
}

func (x *Status) GetRuns() []*Status_Run {
	if x != nil {
		return x.Runs
	}
	return nil
}

func (x *Status) GetSignature() string {
	if x != nil {
		return x.Signature
	}
	return ""
}

func (x *Status) GetVersion() uint32 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *Status) GetAgent() string {
	if x != nil {
		return x.Agent
	}
	return ""
}

// Run represents a single execution
type Status_Run struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Reference     string            `protobuf:"bytes,1,opt,name=reference,proto3" json:"reference,omitempty"`                                   // run reference as set by originator
	Status        Status_Run_Status `protobuf:"varint,2,opt,name=status,proto3,enum=eolymp.executor.Status_Run_Status" json:"status,omitempty"` // run status
	Score         float32           `protobuf:"fixed32,80,opt,name=score,proto3" json:"score,omitempty"`                                        // score, number of points awarded by verifier
	Cost          float32           `protobuf:"fixed32,81,opt,name=cost,proto3" json:"cost,omitempty"`                                          // cost, maximum number of points awarded for the test (as defined in the task)
	WallTimeUsage uint32            `protobuf:"varint,51,opt,name=wall_time_usage,json=wallTimeUsage,proto3" json:"wall_time_usage,omitempty"`  // wall time usage in milliseconds
	WallTimeLimit uint32            `protobuf:"varint,61,opt,name=wall_time_limit,json=wallTimeLimit,proto3" json:"wall_time_limit,omitempty"`  // wall time limit in milliseconds
	CpuTimeUsage  uint32            `protobuf:"varint,52,opt,name=cpu_time_usage,json=cpuTimeUsage,proto3" json:"cpu_time_usage,omitempty"`     // cpu time usage in milliseconds
	CpuTimeLimit  uint32            `protobuf:"varint,62,opt,name=cpu_time_limit,json=cpuTimeLimit,proto3" json:"cpu_time_limit,omitempty"`     // cpu time limit in milliseconds
	MemoryUsage   uint64            `protobuf:"varint,53,opt,name=memory_usage,json=memoryUsage,proto3" json:"memory_usage,omitempty"`          // memory usage in bytes
	MemoryLimit   uint64            `protobuf:"varint,63,opt,name=memory_limit,json=memoryLimit,proto3" json:"memory_limit,omitempty"`          // memory limit in bytes
	ExitCode      uint32            `protobuf:"varint,70,opt,name=exit_code,json=exitCode,proto3" json:"exit_code,omitempty"`                   // program exit code
	Signal        uint32            `protobuf:"varint,71,opt,name=signal,proto3" json:"signal,omitempty"`                                       // might contain signal used to kill program
	// output data
	// only populated if debug was set to true
	// up to 5KB, the rest is truncated
	Output string `protobuf:"bytes,10,opt,name=output,proto3" json:"output,omitempty"`
	// stderr data
	// only populated if debug was set to true
	// up to 5KB, the rest is truncated
	Stderr string `protobuf:"bytes,20,opt,name=stderr,proto3" json:"stderr,omitempty"`
	// output (incl. stderr) produced by verifier during execution
	// only populated if debug was set to true
	// up to 5KB, the rest is truncated
	VerifierLog string `protobuf:"bytes,30,opt,name=verifier_log,json=verifierLog,proto3" json:"verifier_log,omitempty"`
	// output (incl. stderr) produced by interactor during execution
	// only populated if debug was set to true
	// up to 5KB, the rest is truncated
	InteractorLog string `protobuf:"bytes,40,opt,name=interactor_log,json=interactorLog,proto3" json:"interactor_log,omitempty"`
}

func (x *Status_Run) Reset() {
	*x = Status_Run{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_executor_status_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Status_Run) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Status_Run) ProtoMessage() {}

func (x *Status_Run) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_executor_status_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Status_Run.ProtoReflect.Descriptor instead.
func (*Status_Run) Descriptor() ([]byte, []int) {
	return file_eolymp_executor_status_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Status_Run) GetReference() string {
	if x != nil {
		return x.Reference
	}
	return ""
}

func (x *Status_Run) GetStatus() Status_Run_Status {
	if x != nil {
		return x.Status
	}
	return Status_Run_NONE
}

func (x *Status_Run) GetScore() float32 {
	if x != nil {
		return x.Score
	}
	return 0
}

func (x *Status_Run) GetCost() float32 {
	if x != nil {
		return x.Cost
	}
	return 0
}

func (x *Status_Run) GetWallTimeUsage() uint32 {
	if x != nil {
		return x.WallTimeUsage
	}
	return 0
}

func (x *Status_Run) GetWallTimeLimit() uint32 {
	if x != nil {
		return x.WallTimeLimit
	}
	return 0
}

func (x *Status_Run) GetCpuTimeUsage() uint32 {
	if x != nil {
		return x.CpuTimeUsage
	}
	return 0
}

func (x *Status_Run) GetCpuTimeLimit() uint32 {
	if x != nil {
		return x.CpuTimeLimit
	}
	return 0
}

func (x *Status_Run) GetMemoryUsage() uint64 {
	if x != nil {
		return x.MemoryUsage
	}
	return 0
}

func (x *Status_Run) GetMemoryLimit() uint64 {
	if x != nil {
		return x.MemoryLimit
	}
	return 0
}

func (x *Status_Run) GetExitCode() uint32 {
	if x != nil {
		return x.ExitCode
	}
	return 0
}

func (x *Status_Run) GetSignal() uint32 {
	if x != nil {
		return x.Signal
	}
	return 0
}

func (x *Status_Run) GetOutput() string {
	if x != nil {
		return x.Output
	}
	return ""
}

func (x *Status_Run) GetStderr() string {
	if x != nil {
		return x.Stderr
	}
	return ""
}

func (x *Status_Run) GetVerifierLog() string {
	if x != nil {
		return x.VerifierLog
	}
	return ""
}

func (x *Status_Run) GetInteractorLog() string {
	if x != nil {
		return x.InteractorLog
	}
	return ""
}

var File_eolymp_executor_status_proto protoreflect.FileDescriptor

var file_eolymp_executor_status_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x6f,
	0x72, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f,
	0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x6f, 0x72, 0x22,
	0xc0, 0x08, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x72, 0x65,
	0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x72,
	0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x72, 0x69, 0x67,
	0x69, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e,
	0x12, 0x30, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1c,
	0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x6f, 0x72,
	0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x14, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x66, 0x61, 0x69, 0x6c,
	0x75, 0x72, 0x65, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x66, 0x61, 0x69, 0x6c, 0x75,
	0x72, 0x65, 0x12, 0x2f, 0x0a, 0x04, 0x72, 0x75, 0x6e, 0x73, 0x18, 0x28, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x1b, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74,
	0x6f, 0x72, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x52, 0x75, 0x6e, 0x52, 0x04, 0x72,
	0x75, 0x6e, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65,
	0x18, 0x32, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x64, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x61,
	0x67, 0x65, 0x6e, 0x74, 0x18, 0x6e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x61, 0x67, 0x65, 0x6e,
	0x74, 0x1a, 0xe8, 0x05, 0x0a, 0x03, 0x52, 0x75, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x72, 0x65, 0x66,
	0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x72, 0x65,
	0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x3a, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x22, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70,
	0x2e, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x6f, 0x72, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
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
	0x1b, 0x0a, 0x09, 0x65, 0x78, 0x69, 0x74, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x46, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x08, 0x65, 0x78, 0x69, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x73, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x18, 0x47, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x73, 0x69,
	0x67, 0x6e, 0x61, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x18, 0x0a,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x12, 0x16, 0x0a, 0x06,
	0x73, 0x74, 0x64, 0x65, 0x72, 0x72, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74,
	0x64, 0x65, 0x72, 0x72, 0x12, 0x21, 0x0a, 0x0c, 0x76, 0x65, 0x72, 0x69, 0x66, 0x69, 0x65, 0x72,
	0x5f, 0x6c, 0x6f, 0x67, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x76, 0x65, 0x72, 0x69,
	0x66, 0x69, 0x65, 0x72, 0x4c, 0x6f, 0x67, 0x12, 0x25, 0x0a, 0x0e, 0x69, 0x6e, 0x74, 0x65, 0x72,
	0x61, 0x63, 0x74, 0x6f, 0x72, 0x5f, 0x6c, 0x6f, 0x67, 0x18, 0x28, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0d, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x4c, 0x6f, 0x67, 0x22, 0xcb,
	0x01, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x4f, 0x4e,
	0x45, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x50, 0x45, 0x4e, 0x44, 0x49, 0x4e, 0x47, 0x10, 0x01,
	0x12, 0x0d, 0x0a, 0x09, 0x45, 0x58, 0x45, 0x43, 0x55, 0x54, 0x49, 0x4e, 0x47, 0x10, 0x02, 0x12,
	0x0b, 0x0a, 0x07, 0x54, 0x49, 0x4d, 0x45, 0x4f, 0x55, 0x54, 0x10, 0x03, 0x12, 0x11, 0x0a, 0x0d,
	0x43, 0x50, 0x55, 0x5f, 0x45, 0x58, 0x48, 0x41, 0x55, 0x53, 0x54, 0x45, 0x44, 0x10, 0x04, 0x12,
	0x13, 0x0a, 0x0f, 0x4d, 0x45, 0x4d, 0x4f, 0x52, 0x59, 0x5f, 0x4f, 0x56, 0x45, 0x52, 0x46, 0x4c,
	0x4f, 0x57, 0x10, 0x05, 0x12, 0x11, 0x0a, 0x0d, 0x52, 0x55, 0x4e, 0x54, 0x49, 0x4d, 0x45, 0x5f,
	0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x06, 0x12, 0x0c, 0x0a, 0x08, 0x45, 0x58, 0x45, 0x43, 0x55,
	0x54, 0x45, 0x44, 0x10, 0x07, 0x12, 0x0c, 0x0a, 0x08, 0x41, 0x43, 0x43, 0x45, 0x50, 0x54, 0x45,
	0x44, 0x10, 0x08, 0x12, 0x10, 0x0a, 0x0c, 0x57, 0x52, 0x4f, 0x4e, 0x47, 0x5f, 0x41, 0x4e, 0x53,
	0x57, 0x45, 0x52, 0x10, 0x09, 0x12, 0x18, 0x0a, 0x14, 0x56, 0x45, 0x52, 0x49, 0x46, 0x49, 0x43,
	0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x46, 0x41, 0x49, 0x4c, 0x55, 0x52, 0x45, 0x10, 0x0a, 0x12,
	0x0b, 0x0a, 0x07, 0x53, 0x4b, 0x49, 0x50, 0x50, 0x45, 0x44, 0x10, 0x0b, 0x22, 0x34, 0x0a, 0x04,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x0a,
	0x0a, 0x06, 0x55, 0x50, 0x44, 0x41, 0x54, 0x45, 0x10, 0x01, 0x12, 0x09, 0x0a, 0x05, 0x45, 0x52,
	0x52, 0x4f, 0x52, 0x10, 0x02, 0x12, 0x0b, 0x0a, 0x07, 0x46, 0x41, 0x49, 0x4c, 0x55, 0x52, 0x45,
	0x10, 0x03, 0x42, 0x33, 0x5a, 0x31, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x67, 0x6f, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x65,
	0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x6f, 0x72, 0x3b, 0x65,
	0x78, 0x65, 0x63, 0x75, 0x74, 0x6f, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_eolymp_executor_status_proto_rawDescOnce sync.Once
	file_eolymp_executor_status_proto_rawDescData = file_eolymp_executor_status_proto_rawDesc
)

func file_eolymp_executor_status_proto_rawDescGZIP() []byte {
	file_eolymp_executor_status_proto_rawDescOnce.Do(func() {
		file_eolymp_executor_status_proto_rawDescData = protoimpl.X.CompressGZIP(file_eolymp_executor_status_proto_rawDescData)
	})
	return file_eolymp_executor_status_proto_rawDescData
}

var file_eolymp_executor_status_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_eolymp_executor_status_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_eolymp_executor_status_proto_goTypes = []interface{}{
	(Status_Type)(0),       // 0: eolymp.executor.Status.Type
	(Status_Run_Status)(0), // 1: eolymp.executor.Status.Run.Status
	(*Status)(nil),         // 2: eolymp.executor.Status
	(*Status_Run)(nil),     // 3: eolymp.executor.Status.Run
}
var file_eolymp_executor_status_proto_depIdxs = []int32{
	0, // 0: eolymp.executor.Status.type:type_name -> eolymp.executor.Status.Type
	3, // 1: eolymp.executor.Status.runs:type_name -> eolymp.executor.Status.Run
	1, // 2: eolymp.executor.Status.Run.status:type_name -> eolymp.executor.Status.Run.Status
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_eolymp_executor_status_proto_init() }
func file_eolymp_executor_status_proto_init() {
	if File_eolymp_executor_status_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_eolymp_executor_status_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Status); i {
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
		file_eolymp_executor_status_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Status_Run); i {
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
			RawDescriptor: file_eolymp_executor_status_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_eolymp_executor_status_proto_goTypes,
		DependencyIndexes: file_eolymp_executor_status_proto_depIdxs,
		EnumInfos:         file_eolymp_executor_status_proto_enumTypes,
		MessageInfos:      file_eolymp_executor_status_proto_msgTypes,
	}.Build()
	File_eolymp_executor_status_proto = out.File
	file_eolymp_executor_status_proto_rawDesc = nil
	file_eolymp_executor_status_proto_goTypes = nil
	file_eolymp_executor_status_proto_depIdxs = nil
}
