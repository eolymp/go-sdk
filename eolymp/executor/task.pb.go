// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.21.12
// source: eolymp/executor/task.proto

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

// Task represents task to be executed by judge agent.
type Task struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Task unique reference so originator can look up task in its database, for example, this field might be set to
	// Submission ID. Unlike ID, reference is set by client and does not have to be unique. Although, it would make
	// sense for reference to be unique within same originator, otherwise originator won't be able to correlated tasks
	// in its database.
	Reference string `protobuf:"bytes,1,opt,name=reference,proto3" json:"reference,omitempty"`
	// Task originator (service which has created task). This field will be added to task status reports, so consumers
	// can easily filter status reports they are interested in.
	Origin string `protobuf:"bytes,2,opt,name=origin,proto3" json:"origin,omitempty"`
	// Task priority. Allowed values 1-9. Messages with higher priority are precessed first.
	// Currently not supported.
	Priority uint32 `protobuf:"varint,3,opt,name=priority,proto3" json:"priority,omitempty"`
	// Runtime which should be used to execute source code.
	Runtime string `protobuf:"bytes,10,opt,name=runtime,proto3" json:"runtime,omitempty"`
	// Source code.
	Source    string `protobuf:"bytes,11,opt,name=source,proto3" json:"source,omitempty"`                         // deprecated, use source_url instead
	SourceUrl string `protobuf:"bytes,110,opt,name=source_url,json=sourceUrl,proto3" json:"source_url,omitempty"` // source code URL (overrides source)
	HeaderUrl string `protobuf:"bytes,111,opt,name=header_url,json=headerUrl,proto3" json:"header_url,omitempty"` // prepend source code before executing
	FooterUrl string `protobuf:"bytes,112,opt,name=footer_url,json=footerUrl,proto3" json:"footer_url,omitempty"` // append source code before executing
	// Combine stderr and stdout when capturing output. Checker will use combined output as answer. Status will capture
	// both stderr and stdout in output field while stderr will be empty.
	RedirectStderrToStdout bool `protobuf:"varint,13,opt,name=redirect_stderr_to_stdout,json=redirectStderrToStdout,proto3" json:"redirect_stderr_to_stdout,omitempty"`
	// Number of times solution will be executed, after each run (except last) output.txt will be renamed to input.txt.
	RunCount uint32 `protobuf:"varint,16,opt,name=run_count,json=runCount,proto3" json:"run_count,omitempty"`
	// Precondition define conditions when each run should be executed, if runs does not satisfy preconditions it will be skipped.
	Preconditions []*Task_Precondition `protobuf:"bytes,40,rep,name=preconditions,proto3" json:"preconditions,omitempty"`
	// Execution constraints, define limits imposed on each run.
	Constraints []*Task_Constraint `protobuf:"bytes,20,rep,name=constraints,proto3" json:"constraints,omitempty"`
	// Checker configuration.
	Checker *Checker `protobuf:"bytes,24,opt,name=checker,proto3" json:"checker,omitempty"`
	// Interactor configuration
	Interactor *Interactor `protobuf:"bytes,25,opt,name=interactor,proto3" json:"interactor,omitempty"`
	// Run configurations.
	Runs []*Task_Run `protobuf:"bytes,30,rep,name=runs,proto3" json:"runs,omitempty"`
	// Additional files to be placed in the work directory during compilation and runs*
	Files []*Task_File `protobuf:"bytes,50,rep,name=files,proto3" json:"files,omitempty"`
}

func (x *Task) Reset() {
	*x = Task{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_executor_task_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Task) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Task) ProtoMessage() {}

func (x *Task) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_executor_task_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Task.ProtoReflect.Descriptor instead.
func (*Task) Descriptor() ([]byte, []int) {
	return file_eolymp_executor_task_proto_rawDescGZIP(), []int{0}
}

func (x *Task) GetReference() string {
	if x != nil {
		return x.Reference
	}
	return ""
}

func (x *Task) GetOrigin() string {
	if x != nil {
		return x.Origin
	}
	return ""
}

func (x *Task) GetPriority() uint32 {
	if x != nil {
		return x.Priority
	}
	return 0
}

func (x *Task) GetRuntime() string {
	if x != nil {
		return x.Runtime
	}
	return ""
}

func (x *Task) GetSource() string {
	if x != nil {
		return x.Source
	}
	return ""
}

func (x *Task) GetSourceUrl() string {
	if x != nil {
		return x.SourceUrl
	}
	return ""
}

func (x *Task) GetHeaderUrl() string {
	if x != nil {
		return x.HeaderUrl
	}
	return ""
}

func (x *Task) GetFooterUrl() string {
	if x != nil {
		return x.FooterUrl
	}
	return ""
}

func (x *Task) GetRedirectStderrToStdout() bool {
	if x != nil {
		return x.RedirectStderrToStdout
	}
	return false
}

func (x *Task) GetRunCount() uint32 {
	if x != nil {
		return x.RunCount
	}
	return 0
}

func (x *Task) GetPreconditions() []*Task_Precondition {
	if x != nil {
		return x.Preconditions
	}
	return nil
}

func (x *Task) GetConstraints() []*Task_Constraint {
	if x != nil {
		return x.Constraints
	}
	return nil
}

func (x *Task) GetChecker() *Checker {
	if x != nil {
		return x.Checker
	}
	return nil
}

func (x *Task) GetInteractor() *Interactor {
	if x != nil {
		return x.Interactor
	}
	return nil
}

func (x *Task) GetRuns() []*Task_Run {
	if x != nil {
		return x.Runs
	}
	return nil
}

func (x *Task) GetFiles() []*Task_File {
	if x != nil {
		return x.Files
	}
	return nil
}

// File defines additional file which might be placed into work directory during compilation or execution
type Task_File struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Path      string `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`                            // Path where file should be placed (always relative to the workdir)
	SourceUrl string `protobuf:"bytes,3,opt,name=source_url,json=sourceUrl,proto3" json:"source_url,omitempty"` // URL of the file content
}

func (x *Task_File) Reset() {
	*x = Task_File{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_executor_task_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Task_File) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Task_File) ProtoMessage() {}

func (x *Task_File) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_executor_task_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Task_File.ProtoReflect.Descriptor instead.
func (*Task_File) Descriptor() ([]byte, []int) {
	return file_eolymp_executor_task_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Task_File) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *Task_File) GetSourceUrl() string {
	if x != nil {
		return x.SourceUrl
	}
	return ""
}

// Run defines a single execution of the task.
type Task_Run struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Run reference.
	Reference string `protobuf:"bytes,1,opt,name=reference,proto3" json:"reference,omitempty"`
	// Index defines order in which runs are executed.
	Index uint32 `protobuf:"varint,2,opt,name=index,proto3" json:"index,omitempty"`
	// Use debugging (report content of stdout, stderr and exit code).
	Debug bool `protobuf:"varint,3,opt,name=debug,proto3" json:"debug,omitempty"`
	// A number of points awarded for passing this run.
	Cost float32 `protobuf:"fixed32,4,opt,name=cost,proto3" json:"cost,omitempty"`
	// Labels used to match constraints and preconditions.
	Labels []string `protobuf:"bytes,30,rep,name=labels,proto3" json:"labels,omitempty"`
	// Input data, might be empty if checker type is NONE.
	//
	// Types that are assignable to Input:
	//
	//	*Task_Run_InputObjectId
	//	*Task_Run_InputContent
	//	*Task_Run_InputUrl
	Input isTask_Run_Input `protobuf_oneof:"input"`
	// Answer data, should be empty if checker type is NONE.
	//
	// Types that are assignable to Answer:
	//
	//	*Task_Run_AnswerObjectId
	//	*Task_Run_AnswerContent
	//	*Task_Run_AnswerUrl
	Answer isTask_Run_Answer `protobuf_oneof:"answer"`
}

func (x *Task_Run) Reset() {
	*x = Task_Run{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_executor_task_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Task_Run) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Task_Run) ProtoMessage() {}

func (x *Task_Run) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_executor_task_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Task_Run.ProtoReflect.Descriptor instead.
func (*Task_Run) Descriptor() ([]byte, []int) {
	return file_eolymp_executor_task_proto_rawDescGZIP(), []int{0, 1}
}

func (x *Task_Run) GetReference() string {
	if x != nil {
		return x.Reference
	}
	return ""
}

func (x *Task_Run) GetIndex() uint32 {
	if x != nil {
		return x.Index
	}
	return 0
}

func (x *Task_Run) GetDebug() bool {
	if x != nil {
		return x.Debug
	}
	return false
}

func (x *Task_Run) GetCost() float32 {
	if x != nil {
		return x.Cost
	}
	return 0
}

func (x *Task_Run) GetLabels() []string {
	if x != nil {
		return x.Labels
	}
	return nil
}

func (m *Task_Run) GetInput() isTask_Run_Input {
	if m != nil {
		return m.Input
	}
	return nil
}

func (x *Task_Run) GetInputObjectId() string {
	if x, ok := x.GetInput().(*Task_Run_InputObjectId); ok {
		return x.InputObjectId
	}
	return ""
}

func (x *Task_Run) GetInputContent() string {
	if x, ok := x.GetInput().(*Task_Run_InputContent); ok {
		return x.InputContent
	}
	return ""
}

func (x *Task_Run) GetInputUrl() string {
	if x, ok := x.GetInput().(*Task_Run_InputUrl); ok {
		return x.InputUrl
	}
	return ""
}

func (m *Task_Run) GetAnswer() isTask_Run_Answer {
	if m != nil {
		return m.Answer
	}
	return nil
}

func (x *Task_Run) GetAnswerObjectId() string {
	if x, ok := x.GetAnswer().(*Task_Run_AnswerObjectId); ok {
		return x.AnswerObjectId
	}
	return ""
}

func (x *Task_Run) GetAnswerContent() string {
	if x, ok := x.GetAnswer().(*Task_Run_AnswerContent); ok {
		return x.AnswerContent
	}
	return ""
}

func (x *Task_Run) GetAnswerUrl() string {
	if x, ok := x.GetAnswer().(*Task_Run_AnswerUrl); ok {
		return x.AnswerUrl
	}
	return ""
}

type isTask_Run_Input interface {
	isTask_Run_Input()
}

type Task_Run_InputObjectId struct {
	InputObjectId string `protobuf:"bytes,10,opt,name=input_object_id,json=inputObjectId,proto3,oneof"` // download input from object storage
}

type Task_Run_InputContent struct {
	InputContent string `protobuf:"bytes,11,opt,name=input_content,json=inputContent,proto3,oneof"` // use input from this field (up to 1KB)
}

type Task_Run_InputUrl struct {
	InputUrl string `protobuf:"bytes,12,opt,name=input_url,json=inputUrl,proto3,oneof"` // download input via URL
}

func (*Task_Run_InputObjectId) isTask_Run_Input() {}

func (*Task_Run_InputContent) isTask_Run_Input() {}

func (*Task_Run_InputUrl) isTask_Run_Input() {}

type isTask_Run_Answer interface {
	isTask_Run_Answer()
}

type Task_Run_AnswerObjectId struct {
	AnswerObjectId string `protobuf:"bytes,20,opt,name=answer_object_id,json=answerObjectId,proto3,oneof"` // download answer from object storage
}

type Task_Run_AnswerContent struct {
	AnswerContent string `protobuf:"bytes,21,opt,name=answer_content,json=answerContent,proto3,oneof"` // use answer from this field (up to 1KB)
}

type Task_Run_AnswerUrl struct {
	AnswerUrl string `protobuf:"bytes,22,opt,name=answer_url,json=answerUrl,proto3,oneof"` // download answer via URL
}

func (*Task_Run_AnswerObjectId) isTask_Run_Answer() {}

func (*Task_Run_AnswerContent) isTask_Run_Answer() {}

func (*Task_Run_AnswerUrl) isTask_Run_Answer() {}

// Precondition defines criteria for a run to be executed. If this criteria does not meet, the run will be skipped.
type Task_Precondition struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Precondition will apply to all runs matching this label selector.
	// Run must have all labels defined by selector. Empty selector will match all runs.
	Selector []string `protobuf:"bytes,1,rep,name=selector,proto3" json:"selector,omitempty"`
	// Defines label selector for runs which must be ACCEPTED before.
	// Dependent runs will match if they have at least one of the labels defined in depends_on.
	// To match dependent runs with by multiple labels, use two different preconditions.
	// Empty depends_on means no dependencies.
	DependsOn []string `protobuf:"bytes,10,rep,name=depends_on,json=dependsOn,proto3" json:"depends_on,omitempty"`
	// Skip the rest of the runs if one fails.
	StopOnFailure bool `protobuf:"varint,11,opt,name=stop_on_failure,json=stopOnFailure,proto3" json:"stop_on_failure,omitempty"`
	// Defines maximum execution time for all runs matching this precondition.
	// Accepted runs are not counted towards execution time.
	// Once max_execution_time is reached, the rest of the tests are skipped.
	MaxExecutionTime uint32 `protobuf:"varint,12,opt,name=max_execution_time,json=maxExecutionTime,proto3" json:"max_execution_time,omitempty"`
}

func (x *Task_Precondition) Reset() {
	*x = Task_Precondition{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_executor_task_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Task_Precondition) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Task_Precondition) ProtoMessage() {}

func (x *Task_Precondition) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_executor_task_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Task_Precondition.ProtoReflect.Descriptor instead.
func (*Task_Precondition) Descriptor() ([]byte, []int) {
	return file_eolymp_executor_task_proto_rawDescGZIP(), []int{0, 2}
}

func (x *Task_Precondition) GetSelector() []string {
	if x != nil {
		return x.Selector
	}
	return nil
}

func (x *Task_Precondition) GetDependsOn() []string {
	if x != nil {
		return x.DependsOn
	}
	return nil
}

func (x *Task_Precondition) GetStopOnFailure() bool {
	if x != nil {
		return x.StopOnFailure
	}
	return false
}

func (x *Task_Precondition) GetMaxExecutionTime() uint32 {
	if x != nil {
		return x.MaxExecutionTime
	}
	return 0
}

// Constraint defines limitations which apply to the run.
type Task_Constraint struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Constraint will apply to all runs matching this label selector.
	// Run must have all labels defined by selector. Empty selector will match all runs.
	Selector []string `protobuf:"bytes,1,rep,name=selector,proto3" json:"selector,omitempty"`
	// Defines actor to whom the constraint is applied (used by Task V2)
	Actor string `protobuf:"bytes,2,opt,name=actor,proto3" json:"actor,omitempty"`
	// Real-world time limit in milliseconds.
	WallTimeLimit uint32 `protobuf:"varint,10,opt,name=wall_time_limit,json=wallTimeLimit,proto3" json:"wall_time_limit,omitempty"`
	// CPU time limit in milliseconds.
	CpuTimeLimit uint32 `protobuf:"varint,11,opt,name=cpu_time_limit,json=cpuTimeLimit,proto3" json:"cpu_time_limit,omitempty"`
	// Memory limit in bytes.
	MemoryLimit uint64 `protobuf:"varint,12,opt,name=memory_limit,json=memoryLimit,proto3" json:"memory_limit,omitempty"`
	// File size limit in bytes.
	FileSizeLimit uint64 `protobuf:"varint,13,opt,name=file_size_limit,json=fileSizeLimit,proto3" json:"file_size_limit,omitempty"`
}

func (x *Task_Constraint) Reset() {
	*x = Task_Constraint{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_executor_task_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Task_Constraint) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Task_Constraint) ProtoMessage() {}

func (x *Task_Constraint) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_executor_task_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Task_Constraint.ProtoReflect.Descriptor instead.
func (*Task_Constraint) Descriptor() ([]byte, []int) {
	return file_eolymp_executor_task_proto_rawDescGZIP(), []int{0, 3}
}

func (x *Task_Constraint) GetSelector() []string {
	if x != nil {
		return x.Selector
	}
	return nil
}

func (x *Task_Constraint) GetActor() string {
	if x != nil {
		return x.Actor
	}
	return ""
}

func (x *Task_Constraint) GetWallTimeLimit() uint32 {
	if x != nil {
		return x.WallTimeLimit
	}
	return 0
}

func (x *Task_Constraint) GetCpuTimeLimit() uint32 {
	if x != nil {
		return x.CpuTimeLimit
	}
	return 0
}

func (x *Task_Constraint) GetMemoryLimit() uint64 {
	if x != nil {
		return x.MemoryLimit
	}
	return 0
}

func (x *Task_Constraint) GetFileSizeLimit() uint64 {
	if x != nil {
		return x.FileSizeLimit
	}
	return 0
}

var File_eolymp_executor_task_proto protoreflect.FileDescriptor

var file_eolymp_executor_task_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x6f,
	0x72, 0x2f, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x65, 0x6f,
	0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x6f, 0x72, 0x1a, 0x1d, 0x65,
	0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x6f, 0x72, 0x2f, 0x63,
	0x68, 0x65, 0x63, 0x6b, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x65, 0x6f,
	0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x6f, 0x72, 0x2f, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xcd,
	0x0b, 0x0a, 0x04, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x1c, 0x0a, 0x09, 0x72, 0x65, 0x66, 0x65, 0x72,
	0x65, 0x6e, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x72, 0x65, 0x66, 0x65,
	0x72, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x12, 0x1a, 0x0a,
	0x08, 0x70, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x08, 0x70, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x75, 0x6e,
	0x74, 0x69, 0x6d, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x72, 0x75, 0x6e, 0x74,
	0x69, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x0b, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x6e, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x55, 0x72, 0x6c, 0x12, 0x1d, 0x0a, 0x0a, 0x68, 0x65,
	0x61, 0x64, 0x65, 0x72, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x6f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x55, 0x72, 0x6c, 0x12, 0x1d, 0x0a, 0x0a, 0x66, 0x6f, 0x6f,
	0x74, 0x65, 0x72, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x70, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66,
	0x6f, 0x6f, 0x74, 0x65, 0x72, 0x55, 0x72, 0x6c, 0x12, 0x39, 0x0a, 0x19, 0x72, 0x65, 0x64, 0x69,
	0x72, 0x65, 0x63, 0x74, 0x5f, 0x73, 0x74, 0x64, 0x65, 0x72, 0x72, 0x5f, 0x74, 0x6f, 0x5f, 0x73,
	0x74, 0x64, 0x6f, 0x75, 0x74, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x08, 0x52, 0x16, 0x72, 0x65, 0x64,
	0x69, 0x72, 0x65, 0x63, 0x74, 0x53, 0x74, 0x64, 0x65, 0x72, 0x72, 0x54, 0x6f, 0x53, 0x74, 0x64,
	0x6f, 0x75, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x72, 0x75, 0x6e, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x10, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x72, 0x75, 0x6e, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x12, 0x48, 0x0a, 0x0d, 0x70, 0x72, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x18, 0x28, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70,
	0x2e, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x6f, 0x72, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x2e, 0x50,
	0x72, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0d, 0x70, 0x72, 0x65,
	0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x42, 0x0a, 0x0b, 0x63, 0x6f,
	0x6e, 0x73, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x74, 0x73, 0x18, 0x14, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x20, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x6f,
	0x72, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x2e, 0x43, 0x6f, 0x6e, 0x73, 0x74, 0x72, 0x61, 0x69, 0x6e,
	0x74, 0x52, 0x0b, 0x63, 0x6f, 0x6e, 0x73, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x74, 0x73, 0x12, 0x32,
	0x0a, 0x07, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x65, 0x72, 0x18, 0x18, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x18, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x6f,
	0x72, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x65, 0x72, 0x52, 0x07, 0x63, 0x68, 0x65, 0x63, 0x6b,
	0x65, 0x72, 0x12, 0x3b, 0x0a, 0x0a, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x6f, 0x72,
	0x18, 0x19, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e,
	0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x6f, 0x72, 0x2e, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63,
	0x74, 0x6f, 0x72, 0x52, 0x0a, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x12,
	0x2d, 0x0a, 0x04, 0x72, 0x75, 0x6e, 0x73, 0x18, 0x1e, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e,
	0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x6f, 0x72, 0x2e,
	0x54, 0x61, 0x73, 0x6b, 0x2e, 0x52, 0x75, 0x6e, 0x52, 0x04, 0x72, 0x75, 0x6e, 0x73, 0x12, 0x30,
	0x0a, 0x05, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x18, 0x32, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x6f, 0x72, 0x2e,
	0x54, 0x61, 0x73, 0x6b, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x05, 0x66, 0x69, 0x6c, 0x65, 0x73,
	0x1a, 0x39, 0x0a, 0x04, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x1d, 0x0a, 0x0a,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x55, 0x72, 0x6c, 0x1a, 0xf4, 0x02, 0x0a, 0x03,
	0x52, 0x75, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x14, 0x0a, 0x05, 0x64, 0x65, 0x62, 0x75, 0x67,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x64, 0x65, 0x62, 0x75, 0x67, 0x12, 0x12, 0x0a,
	0x04, 0x63, 0x6f, 0x73, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x02, 0x52, 0x04, 0x63, 0x6f, 0x73,
	0x74, 0x12, 0x16, 0x0a, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x18, 0x1e, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x12, 0x28, 0x0a, 0x0f, 0x69, 0x6e, 0x70,
	0x75, 0x74, 0x5f, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x09, 0x48, 0x00, 0x52, 0x0d, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x4f, 0x62, 0x6a, 0x65, 0x63,
	0x74, 0x49, 0x64, 0x12, 0x25, 0x0a, 0x0d, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x5f, 0x63, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0c, 0x69, 0x6e,
	0x70, 0x75, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x1d, 0x0a, 0x09, 0x69, 0x6e,
	0x70, 0x75, 0x74, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52,
	0x08, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x55, 0x72, 0x6c, 0x12, 0x2a, 0x0a, 0x10, 0x61, 0x6e, 0x73,
	0x77, 0x65, 0x72, 0x5f, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x14, 0x20,
	0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x0e, 0x61, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x4f, 0x62, 0x6a,
	0x65, 0x63, 0x74, 0x49, 0x64, 0x12, 0x27, 0x0a, 0x0e, 0x61, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x5f,
	0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x15, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52,
	0x0d, 0x61, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x1f,
	0x0a, 0x0a, 0x61, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x16, 0x20, 0x01,
	0x28, 0x09, 0x48, 0x01, 0x52, 0x09, 0x61, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x55, 0x72, 0x6c, 0x42,
	0x07, 0x0a, 0x05, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x42, 0x08, 0x0a, 0x06, 0x61, 0x6e, 0x73, 0x77,
	0x65, 0x72, 0x1a, 0x9f, 0x01, 0x0a, 0x0c, 0x50, 0x72, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x12,
	0x1d, 0x0a, 0x0a, 0x64, 0x65, 0x70, 0x65, 0x6e, 0x64, 0x73, 0x5f, 0x6f, 0x6e, 0x18, 0x0a, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x09, 0x64, 0x65, 0x70, 0x65, 0x6e, 0x64, 0x73, 0x4f, 0x6e, 0x12, 0x26,
	0x0a, 0x0f, 0x73, 0x74, 0x6f, 0x70, 0x5f, 0x6f, 0x6e, 0x5f, 0x66, 0x61, 0x69, 0x6c, 0x75, 0x72,
	0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d, 0x73, 0x74, 0x6f, 0x70, 0x4f, 0x6e, 0x46,
	0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x12, 0x2c, 0x0a, 0x12, 0x6d, 0x61, 0x78, 0x5f, 0x65, 0x78,
	0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x0c, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x10, 0x6d, 0x61, 0x78, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e,
	0x54, 0x69, 0x6d, 0x65, 0x1a, 0xd7, 0x01, 0x0a, 0x0a, 0x43, 0x6f, 0x6e, 0x73, 0x74, 0x72, 0x61,
	0x69, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x12,
	0x14, 0x0a, 0x05, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x61, 0x63, 0x74, 0x6f, 0x72, 0x12, 0x26, 0x0a, 0x0f, 0x77, 0x61, 0x6c, 0x6c, 0x5f, 0x74, 0x69,
	0x6d, 0x65, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0d,
	0x77, 0x61, 0x6c, 0x6c, 0x54, 0x69, 0x6d, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x24, 0x0a,
	0x0e, 0x63, 0x70, 0x75, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18,
	0x0b, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c, 0x63, 0x70, 0x75, 0x54, 0x69, 0x6d, 0x65, 0x4c, 0x69,
	0x6d, 0x69, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x5f, 0x6c, 0x69,
	0x6d, 0x69, 0x74, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0b, 0x6d, 0x65, 0x6d, 0x6f, 0x72,
	0x79, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x26, 0x0a, 0x0f, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x73,
	0x69, 0x7a, 0x65, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x0d, 0x66, 0x69, 0x6c, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x42, 0x33,
	0x5a, 0x31, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6f, 0x6c,
	0x79, 0x6d, 0x70, 0x2f, 0x67, 0x6f, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x65, 0x6f, 0x6c, 0x79, 0x6d,
	0x70, 0x2f, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x6f, 0x72, 0x3b, 0x65, 0x78, 0x65, 0x63, 0x75,
	0x74, 0x6f, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_eolymp_executor_task_proto_rawDescOnce sync.Once
	file_eolymp_executor_task_proto_rawDescData = file_eolymp_executor_task_proto_rawDesc
)

func file_eolymp_executor_task_proto_rawDescGZIP() []byte {
	file_eolymp_executor_task_proto_rawDescOnce.Do(func() {
		file_eolymp_executor_task_proto_rawDescData = protoimpl.X.CompressGZIP(file_eolymp_executor_task_proto_rawDescData)
	})
	return file_eolymp_executor_task_proto_rawDescData
}

var file_eolymp_executor_task_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_eolymp_executor_task_proto_goTypes = []any{
	(*Task)(nil),              // 0: eolymp.executor.Task
	(*Task_File)(nil),         // 1: eolymp.executor.Task.File
	(*Task_Run)(nil),          // 2: eolymp.executor.Task.Run
	(*Task_Precondition)(nil), // 3: eolymp.executor.Task.Precondition
	(*Task_Constraint)(nil),   // 4: eolymp.executor.Task.Constraint
	(*Checker)(nil),           // 5: eolymp.executor.Checker
	(*Interactor)(nil),        // 6: eolymp.executor.Interactor
}
var file_eolymp_executor_task_proto_depIdxs = []int32{
	3, // 0: eolymp.executor.Task.preconditions:type_name -> eolymp.executor.Task.Precondition
	4, // 1: eolymp.executor.Task.constraints:type_name -> eolymp.executor.Task.Constraint
	5, // 2: eolymp.executor.Task.checker:type_name -> eolymp.executor.Checker
	6, // 3: eolymp.executor.Task.interactor:type_name -> eolymp.executor.Interactor
	2, // 4: eolymp.executor.Task.runs:type_name -> eolymp.executor.Task.Run
	1, // 5: eolymp.executor.Task.files:type_name -> eolymp.executor.Task.File
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_eolymp_executor_task_proto_init() }
func file_eolymp_executor_task_proto_init() {
	if File_eolymp_executor_task_proto != nil {
		return
	}
	file_eolymp_executor_checker_proto_init()
	file_eolymp_executor_interactor_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_eolymp_executor_task_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*Task); i {
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
		file_eolymp_executor_task_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*Task_File); i {
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
		file_eolymp_executor_task_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*Task_Run); i {
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
		file_eolymp_executor_task_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*Task_Precondition); i {
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
		file_eolymp_executor_task_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*Task_Constraint); i {
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
	file_eolymp_executor_task_proto_msgTypes[2].OneofWrappers = []any{
		(*Task_Run_InputObjectId)(nil),
		(*Task_Run_InputContent)(nil),
		(*Task_Run_InputUrl)(nil),
		(*Task_Run_AnswerObjectId)(nil),
		(*Task_Run_AnswerContent)(nil),
		(*Task_Run_AnswerUrl)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_eolymp_executor_task_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_eolymp_executor_task_proto_goTypes,
		DependencyIndexes: file_eolymp_executor_task_proto_depIdxs,
		MessageInfos:      file_eolymp_executor_task_proto_msgTypes,
	}.Build()
	File_eolymp_executor_task_proto = out.File
	file_eolymp_executor_task_proto_rawDesc = nil
	file_eolymp_executor_task_proto_goTypes = nil
	file_eolymp_executor_task_proto_depIdxs = nil
}
