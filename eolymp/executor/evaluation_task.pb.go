// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v3.21.12
// source: eolymp/executor/evaluation_task.proto

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
type EvaluationTask struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TaskId    string `protobuf:"bytes,1,opt,name=task_id,json=taskId,proto3" json:"task_id,omitempty"`
	Reference string `protobuf:"bytes,2,opt,name=reference,proto3" json:"reference,omitempty"`
	Origin    string `protobuf:"bytes,3,opt,name=origin,proto3" json:"origin,omitempty"`
	// Task priority. Allowed values 1-9. Messages with higher priority are precessed first.
	// Currently not supported.
	Priority uint32 `protobuf:"varint,4,opt,name=priority,proto3" json:"priority,omitempty"`
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
	Preconditions []*EvaluationTask_Precondition `protobuf:"bytes,40,rep,name=preconditions,proto3" json:"preconditions,omitempty"`
	// Execution constraints, define limits imposed on each run.
	Constraints []*EvaluationTask_Constraint `protobuf:"bytes,20,rep,name=constraints,proto3" json:"constraints,omitempty"`
	// Checker configuration.
	Checker *Checker `protobuf:"bytes,24,opt,name=checker,proto3" json:"checker,omitempty"`
	// Interactor configuration
	Interactor *Interactor `protobuf:"bytes,25,opt,name=interactor,proto3" json:"interactor,omitempty"`
	// Run configurations.
	Runs []*EvaluationTask_Run `protobuf:"bytes,30,rep,name=runs,proto3" json:"runs,omitempty"`
	// Additional files to be placed in the work directory during compilation and runs*
	Files []*File `protobuf:"bytes,50,rep,name=files,proto3" json:"files,omitempty"`
	// Additional scripts to be used during execution.
	Scripts []*Script `protobuf:"bytes,60,rep,name=scripts,proto3" json:"scripts,omitempty"`
}

func (x *EvaluationTask) Reset() {
	*x = EvaluationTask{}
	mi := &file_eolymp_executor_evaluation_task_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *EvaluationTask) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EvaluationTask) ProtoMessage() {}

func (x *EvaluationTask) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_executor_evaluation_task_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EvaluationTask.ProtoReflect.Descriptor instead.
func (*EvaluationTask) Descriptor() ([]byte, []int) {
	return file_eolymp_executor_evaluation_task_proto_rawDescGZIP(), []int{0}
}

func (x *EvaluationTask) GetTaskId() string {
	if x != nil {
		return x.TaskId
	}
	return ""
}

func (x *EvaluationTask) GetReference() string {
	if x != nil {
		return x.Reference
	}
	return ""
}

func (x *EvaluationTask) GetOrigin() string {
	if x != nil {
		return x.Origin
	}
	return ""
}

func (x *EvaluationTask) GetPriority() uint32 {
	if x != nil {
		return x.Priority
	}
	return 0
}

func (x *EvaluationTask) GetRuntime() string {
	if x != nil {
		return x.Runtime
	}
	return ""
}

func (x *EvaluationTask) GetSource() string {
	if x != nil {
		return x.Source
	}
	return ""
}

func (x *EvaluationTask) GetSourceUrl() string {
	if x != nil {
		return x.SourceUrl
	}
	return ""
}

func (x *EvaluationTask) GetHeaderUrl() string {
	if x != nil {
		return x.HeaderUrl
	}
	return ""
}

func (x *EvaluationTask) GetFooterUrl() string {
	if x != nil {
		return x.FooterUrl
	}
	return ""
}

func (x *EvaluationTask) GetRedirectStderrToStdout() bool {
	if x != nil {
		return x.RedirectStderrToStdout
	}
	return false
}

func (x *EvaluationTask) GetRunCount() uint32 {
	if x != nil {
		return x.RunCount
	}
	return 0
}

func (x *EvaluationTask) GetPreconditions() []*EvaluationTask_Precondition {
	if x != nil {
		return x.Preconditions
	}
	return nil
}

func (x *EvaluationTask) GetConstraints() []*EvaluationTask_Constraint {
	if x != nil {
		return x.Constraints
	}
	return nil
}

func (x *EvaluationTask) GetChecker() *Checker {
	if x != nil {
		return x.Checker
	}
	return nil
}

func (x *EvaluationTask) GetInteractor() *Interactor {
	if x != nil {
		return x.Interactor
	}
	return nil
}

func (x *EvaluationTask) GetRuns() []*EvaluationTask_Run {
	if x != nil {
		return x.Runs
	}
	return nil
}

func (x *EvaluationTask) GetFiles() []*File {
	if x != nil {
		return x.Files
	}
	return nil
}

func (x *EvaluationTask) GetScripts() []*Script {
	if x != nil {
		return x.Scripts
	}
	return nil
}

type EvaluationTask_Generator struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ScriptName string   `protobuf:"bytes,1,opt,name=script_name,json=scriptName,proto3" json:"script_name,omitempty"`
	Arguments  []string `protobuf:"bytes,2,rep,name=arguments,proto3" json:"arguments,omitempty"`
}

func (x *EvaluationTask_Generator) Reset() {
	*x = EvaluationTask_Generator{}
	mi := &file_eolymp_executor_evaluation_task_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *EvaluationTask_Generator) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EvaluationTask_Generator) ProtoMessage() {}

func (x *EvaluationTask_Generator) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_executor_evaluation_task_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EvaluationTask_Generator.ProtoReflect.Descriptor instead.
func (*EvaluationTask_Generator) Descriptor() ([]byte, []int) {
	return file_eolymp_executor_evaluation_task_proto_rawDescGZIP(), []int{0, 0}
}

func (x *EvaluationTask_Generator) GetScriptName() string {
	if x != nil {
		return x.ScriptName
	}
	return ""
}

func (x *EvaluationTask_Generator) GetArguments() []string {
	if x != nil {
		return x.Arguments
	}
	return nil
}

// Run defines a single execution of the task.
type EvaluationTask_Run struct {
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
	//	*EvaluationTask_Run_InputUrl
	//	*EvaluationTask_Run_InputContent
	//	*EvaluationTask_Run_InputGenerator
	Input isEvaluationTask_Run_Input `protobuf_oneof:"input"`
	// Answer data, should be empty if checker type is NONE.
	//
	// Types that are assignable to Answer:
	//
	//	*EvaluationTask_Run_AnswerUrl
	//	*EvaluationTask_Run_AnswerContent
	//	*EvaluationTask_Run_AnswerGenerator
	Answer isEvaluationTask_Run_Answer `protobuf_oneof:"answer"`
}

func (x *EvaluationTask_Run) Reset() {
	*x = EvaluationTask_Run{}
	mi := &file_eolymp_executor_evaluation_task_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *EvaluationTask_Run) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EvaluationTask_Run) ProtoMessage() {}

func (x *EvaluationTask_Run) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_executor_evaluation_task_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EvaluationTask_Run.ProtoReflect.Descriptor instead.
func (*EvaluationTask_Run) Descriptor() ([]byte, []int) {
	return file_eolymp_executor_evaluation_task_proto_rawDescGZIP(), []int{0, 1}
}

func (x *EvaluationTask_Run) GetReference() string {
	if x != nil {
		return x.Reference
	}
	return ""
}

func (x *EvaluationTask_Run) GetIndex() uint32 {
	if x != nil {
		return x.Index
	}
	return 0
}

func (x *EvaluationTask_Run) GetDebug() bool {
	if x != nil {
		return x.Debug
	}
	return false
}

func (x *EvaluationTask_Run) GetCost() float32 {
	if x != nil {
		return x.Cost
	}
	return 0
}

func (x *EvaluationTask_Run) GetLabels() []string {
	if x != nil {
		return x.Labels
	}
	return nil
}

func (m *EvaluationTask_Run) GetInput() isEvaluationTask_Run_Input {
	if m != nil {
		return m.Input
	}
	return nil
}

func (x *EvaluationTask_Run) GetInputUrl() string {
	if x, ok := x.GetInput().(*EvaluationTask_Run_InputUrl); ok {
		return x.InputUrl
	}
	return ""
}

func (x *EvaluationTask_Run) GetInputContent() string {
	if x, ok := x.GetInput().(*EvaluationTask_Run_InputContent); ok {
		return x.InputContent
	}
	return ""
}

func (x *EvaluationTask_Run) GetInputGenerator() *EvaluationTask_Generator {
	if x, ok := x.GetInput().(*EvaluationTask_Run_InputGenerator); ok {
		return x.InputGenerator
	}
	return nil
}

func (m *EvaluationTask_Run) GetAnswer() isEvaluationTask_Run_Answer {
	if m != nil {
		return m.Answer
	}
	return nil
}

func (x *EvaluationTask_Run) GetAnswerUrl() string {
	if x, ok := x.GetAnswer().(*EvaluationTask_Run_AnswerUrl); ok {
		return x.AnswerUrl
	}
	return ""
}

func (x *EvaluationTask_Run) GetAnswerContent() string {
	if x, ok := x.GetAnswer().(*EvaluationTask_Run_AnswerContent); ok {
		return x.AnswerContent
	}
	return ""
}

func (x *EvaluationTask_Run) GetAnswerGenerator() *EvaluationTask_Generator {
	if x, ok := x.GetAnswer().(*EvaluationTask_Run_AnswerGenerator); ok {
		return x.AnswerGenerator
	}
	return nil
}

type isEvaluationTask_Run_Input interface {
	isEvaluationTask_Run_Input()
}

type EvaluationTask_Run_InputUrl struct {
	InputUrl string `protobuf:"bytes,12,opt,name=input_url,json=inputUrl,proto3,oneof"` // download input via URL
}

type EvaluationTask_Run_InputContent struct {
	InputContent string `protobuf:"bytes,11,opt,name=input_content,json=inputContent,proto3,oneof"` // use input from this field (up to 1KB)
}

type EvaluationTask_Run_InputGenerator struct {
	InputGenerator *EvaluationTask_Generator `protobuf:"bytes,13,opt,name=input_generator,json=inputGenerator,proto3,oneof"` // generate input using script
}

func (*EvaluationTask_Run_InputUrl) isEvaluationTask_Run_Input() {}

func (*EvaluationTask_Run_InputContent) isEvaluationTask_Run_Input() {}

func (*EvaluationTask_Run_InputGenerator) isEvaluationTask_Run_Input() {}

type isEvaluationTask_Run_Answer interface {
	isEvaluationTask_Run_Answer()
}

type EvaluationTask_Run_AnswerUrl struct {
	AnswerUrl string `protobuf:"bytes,22,opt,name=answer_url,json=answerUrl,proto3,oneof"` // download answer via URL
}

type EvaluationTask_Run_AnswerContent struct {
	AnswerContent string `protobuf:"bytes,21,opt,name=answer_content,json=answerContent,proto3,oneof"` // use answer from this field (up to 1KB)
}

type EvaluationTask_Run_AnswerGenerator struct {
	AnswerGenerator *EvaluationTask_Generator `protobuf:"bytes,23,opt,name=answer_generator,json=answerGenerator,proto3,oneof"` // generate input using script
}

func (*EvaluationTask_Run_AnswerUrl) isEvaluationTask_Run_Answer() {}

func (*EvaluationTask_Run_AnswerContent) isEvaluationTask_Run_Answer() {}

func (*EvaluationTask_Run_AnswerGenerator) isEvaluationTask_Run_Answer() {}

// Precondition defines criteria for a run to be executed. If this criteria does not meet, the run will be skipped.
type EvaluationTask_Precondition struct {
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

func (x *EvaluationTask_Precondition) Reset() {
	*x = EvaluationTask_Precondition{}
	mi := &file_eolymp_executor_evaluation_task_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *EvaluationTask_Precondition) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EvaluationTask_Precondition) ProtoMessage() {}

func (x *EvaluationTask_Precondition) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_executor_evaluation_task_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EvaluationTask_Precondition.ProtoReflect.Descriptor instead.
func (*EvaluationTask_Precondition) Descriptor() ([]byte, []int) {
	return file_eolymp_executor_evaluation_task_proto_rawDescGZIP(), []int{0, 2}
}

func (x *EvaluationTask_Precondition) GetSelector() []string {
	if x != nil {
		return x.Selector
	}
	return nil
}

func (x *EvaluationTask_Precondition) GetDependsOn() []string {
	if x != nil {
		return x.DependsOn
	}
	return nil
}

func (x *EvaluationTask_Precondition) GetStopOnFailure() bool {
	if x != nil {
		return x.StopOnFailure
	}
	return false
}

func (x *EvaluationTask_Precondition) GetMaxExecutionTime() uint32 {
	if x != nil {
		return x.MaxExecutionTime
	}
	return 0
}

// Constraint defines limitations which apply to the run.
type EvaluationTask_Constraint struct {
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

func (x *EvaluationTask_Constraint) Reset() {
	*x = EvaluationTask_Constraint{}
	mi := &file_eolymp_executor_evaluation_task_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *EvaluationTask_Constraint) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EvaluationTask_Constraint) ProtoMessage() {}

func (x *EvaluationTask_Constraint) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_executor_evaluation_task_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EvaluationTask_Constraint.ProtoReflect.Descriptor instead.
func (*EvaluationTask_Constraint) Descriptor() ([]byte, []int) {
	return file_eolymp_executor_evaluation_task_proto_rawDescGZIP(), []int{0, 3}
}

func (x *EvaluationTask_Constraint) GetSelector() []string {
	if x != nil {
		return x.Selector
	}
	return nil
}

func (x *EvaluationTask_Constraint) GetActor() string {
	if x != nil {
		return x.Actor
	}
	return ""
}

func (x *EvaluationTask_Constraint) GetWallTimeLimit() uint32 {
	if x != nil {
		return x.WallTimeLimit
	}
	return 0
}

func (x *EvaluationTask_Constraint) GetCpuTimeLimit() uint32 {
	if x != nil {
		return x.CpuTimeLimit
	}
	return 0
}

func (x *EvaluationTask_Constraint) GetMemoryLimit() uint64 {
	if x != nil {
		return x.MemoryLimit
	}
	return 0
}

func (x *EvaluationTask_Constraint) GetFileSizeLimit() uint64 {
	if x != nil {
		return x.FileSizeLimit
	}
	return 0
}

var File_eolymp_executor_evaluation_task_proto protoreflect.FileDescriptor

var file_eolymp_executor_evaluation_task_proto_rawDesc = []byte{
	0x0a, 0x25, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x6f,
	0x72, 0x2f, 0x65, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x61, 0x73,
	0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e,
	0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x6f, 0x72, 0x1a, 0x1d, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70,
	0x2f, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x6f, 0x72, 0x2f, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x65,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1a, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f,
	0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x6f, 0x72, 0x2f, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x65, 0x78, 0x65, 0x63,
	0x75, 0x74, 0x6f, 0x72, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x65, 0x78,
	0x65, 0x63, 0x75, 0x74, 0x6f, 0x72, 0x2f, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xa5, 0x0d, 0x0a, 0x0e, 0x45, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x17, 0x0a, 0x07, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x61, 0x73, 0x6b, 0x49, 0x64, 0x12,
	0x1c, 0x0a, 0x09, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6f,
	0x72, 0x69, 0x67, 0x69, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74,
	0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x70, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74,
	0x79, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x75, 0x72,
	0x6c, 0x18, 0x6e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x55,
	0x72, 0x6c, 0x12, 0x1d, 0x0a, 0x0a, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x5f, 0x75, 0x72, 0x6c,
	0x18, 0x6f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x55, 0x72,
	0x6c, 0x12, 0x1d, 0x0a, 0x0a, 0x66, 0x6f, 0x6f, 0x74, 0x65, 0x72, 0x5f, 0x75, 0x72, 0x6c, 0x18,
	0x70, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x6f, 0x6f, 0x74, 0x65, 0x72, 0x55, 0x72, 0x6c,
	0x12, 0x39, 0x0a, 0x19, 0x72, 0x65, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x5f, 0x73, 0x74, 0x64,
	0x65, 0x72, 0x72, 0x5f, 0x74, 0x6f, 0x5f, 0x73, 0x74, 0x64, 0x6f, 0x75, 0x74, 0x18, 0x0d, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x16, 0x72, 0x65, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x53, 0x74, 0x64,
	0x65, 0x72, 0x72, 0x54, 0x6f, 0x53, 0x74, 0x64, 0x6f, 0x75, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x72,
	0x75, 0x6e, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x10, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08,
	0x72, 0x75, 0x6e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x52, 0x0a, 0x0d, 0x70, 0x72, 0x65, 0x63,
	0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x28, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x2c, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x6f,
	0x72, 0x2e, 0x45, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x61, 0x73, 0x6b,
	0x2e, 0x50, 0x72, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0d, 0x70,
	0x72, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x4c, 0x0a, 0x0b,
	0x63, 0x6f, 0x6e, 0x73, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x74, 0x73, 0x18, 0x14, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x2a, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x65, 0x78, 0x65, 0x63, 0x75,
	0x74, 0x6f, 0x72, 0x2e, 0x45, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x61,
	0x73, 0x6b, 0x2e, 0x43, 0x6f, 0x6e, 0x73, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x74, 0x52, 0x0b, 0x63,
	0x6f, 0x6e, 0x73, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x74, 0x73, 0x12, 0x32, 0x0a, 0x07, 0x63, 0x68,
	0x65, 0x63, 0x6b, 0x65, 0x72, 0x18, 0x18, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x65, 0x6f,
	0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x6f, 0x72, 0x2e, 0x43, 0x68,
	0x65, 0x63, 0x6b, 0x65, 0x72, 0x52, 0x07, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x65, 0x72, 0x12, 0x3b,
	0x0a, 0x0a, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x18, 0x19, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x65, 0x78, 0x65, 0x63,
	0x75, 0x74, 0x6f, 0x72, 0x2e, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x52,
	0x0a, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x12, 0x37, 0x0a, 0x04, 0x72,
	0x75, 0x6e, 0x73, 0x18, 0x1e, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x65, 0x6f, 0x6c, 0x79,
	0x6d, 0x70, 0x2e, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x6f, 0x72, 0x2e, 0x45, 0x76, 0x61, 0x6c,
	0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x61, 0x73, 0x6b, 0x2e, 0x52, 0x75, 0x6e, 0x52, 0x04,
	0x72, 0x75, 0x6e, 0x73, 0x12, 0x2b, 0x0a, 0x05, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x18, 0x32, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x65, 0x78, 0x65,
	0x63, 0x75, 0x74, 0x6f, 0x72, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x05, 0x66, 0x69, 0x6c, 0x65,
	0x73, 0x12, 0x31, 0x0a, 0x07, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x73, 0x18, 0x3c, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x17, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x65, 0x78, 0x65, 0x63,
	0x75, 0x74, 0x6f, 0x72, 0x2e, 0x53, 0x63, 0x72, 0x69, 0x70, 0x74, 0x52, 0x07, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x73, 0x1a, 0x4a, 0x0a, 0x09, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x6f,
	0x72, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x61, 0x72, 0x67, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x61, 0x72, 0x67, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x73,
	0x1a, 0xcc, 0x03, 0x0a, 0x03, 0x52, 0x75, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x72, 0x65, 0x66, 0x65,
	0x72, 0x65, 0x6e, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x72, 0x65, 0x66,
	0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x14, 0x0a, 0x05,
	0x64, 0x65, 0x62, 0x75, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x64, 0x65, 0x62,
	0x75, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x73, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x02,
	0x52, 0x04, 0x63, 0x6f, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73,
	0x18, 0x1e, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x12, 0x1d,
	0x0a, 0x09, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x0c, 0x20, 0x01, 0x28,
	0x09, 0x48, 0x00, 0x52, 0x08, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x55, 0x72, 0x6c, 0x12, 0x25, 0x0a,
	0x0d, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x0b,
	0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0c, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x43, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x12, 0x54, 0x0a, 0x0f, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x5f, 0x67, 0x65,
	0x6e, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e,
	0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x6f, 0x72, 0x2e,
	0x45, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x61, 0x73, 0x6b, 0x2e, 0x47,
	0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x48, 0x00, 0x52, 0x0e, 0x69, 0x6e, 0x70, 0x75,
	0x74, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x1f, 0x0a, 0x0a, 0x61, 0x6e,
	0x73, 0x77, 0x65, 0x72, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x16, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01,
	0x52, 0x09, 0x61, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x55, 0x72, 0x6c, 0x12, 0x27, 0x0a, 0x0e, 0x61,
	0x6e, 0x73, 0x77, 0x65, 0x72, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x15, 0x20,
	0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x0d, 0x61, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x43, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x12, 0x56, 0x0a, 0x10, 0x61, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x5f, 0x67,
	0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x18, 0x17, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29,
	0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x6f, 0x72,
	0x2e, 0x45, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x61, 0x73, 0x6b, 0x2e,
	0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x48, 0x01, 0x52, 0x0f, 0x61, 0x6e, 0x73,
	0x77, 0x65, 0x72, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x42, 0x07, 0x0a, 0x05,
	0x69, 0x6e, 0x70, 0x75, 0x74, 0x42, 0x08, 0x0a, 0x06, 0x61, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x1a,
	0x9f, 0x01, 0x0a, 0x0c, 0x50, 0x72, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x1a, 0x0a, 0x08, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x08, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x12, 0x1d, 0x0a, 0x0a,
	0x64, 0x65, 0x70, 0x65, 0x6e, 0x64, 0x73, 0x5f, 0x6f, 0x6e, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x09, 0x64, 0x65, 0x70, 0x65, 0x6e, 0x64, 0x73, 0x4f, 0x6e, 0x12, 0x26, 0x0a, 0x0f, 0x73,
	0x74, 0x6f, 0x70, 0x5f, 0x6f, 0x6e, 0x5f, 0x66, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x18, 0x0b,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x0d, 0x73, 0x74, 0x6f, 0x70, 0x4f, 0x6e, 0x46, 0x61, 0x69, 0x6c,
	0x75, 0x72, 0x65, 0x12, 0x2c, 0x0a, 0x12, 0x6d, 0x61, 0x78, 0x5f, 0x65, 0x78, 0x65, 0x63, 0x75,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x10, 0x6d, 0x61, 0x78, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d,
	0x65, 0x1a, 0xd7, 0x01, 0x0a, 0x0a, 0x43, 0x6f, 0x6e, 0x73, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x74,
	0x12, 0x1a, 0x0a, 0x08, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x08, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x12, 0x14, 0x0a, 0x05,
	0x61, 0x63, 0x74, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x61, 0x63, 0x74,
	0x6f, 0x72, 0x12, 0x26, 0x0a, 0x0f, 0x77, 0x61, 0x6c, 0x6c, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x5f,
	0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0d, 0x77, 0x61, 0x6c,
	0x6c, 0x54, 0x69, 0x6d, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x24, 0x0a, 0x0e, 0x63, 0x70,
	0x75, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x0b, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x0c, 0x63, 0x70, 0x75, 0x54, 0x69, 0x6d, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74,
	0x12, 0x21, 0x0a, 0x0c, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74,
	0x18, 0x0c, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0b, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x4c, 0x69,
	0x6d, 0x69, 0x74, 0x12, 0x26, 0x0a, 0x0f, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65,
	0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0d, 0x66, 0x69,
	0x6c, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x42, 0x33, 0x5a, 0x31, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70,
	0x2f, 0x67, 0x6f, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x65,
	0x78, 0x65, 0x63, 0x75, 0x74, 0x6f, 0x72, 0x3b, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x6f, 0x72,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_eolymp_executor_evaluation_task_proto_rawDescOnce sync.Once
	file_eolymp_executor_evaluation_task_proto_rawDescData = file_eolymp_executor_evaluation_task_proto_rawDesc
)

func file_eolymp_executor_evaluation_task_proto_rawDescGZIP() []byte {
	file_eolymp_executor_evaluation_task_proto_rawDescOnce.Do(func() {
		file_eolymp_executor_evaluation_task_proto_rawDescData = protoimpl.X.CompressGZIP(file_eolymp_executor_evaluation_task_proto_rawDescData)
	})
	return file_eolymp_executor_evaluation_task_proto_rawDescData
}

var file_eolymp_executor_evaluation_task_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_eolymp_executor_evaluation_task_proto_goTypes = []any{
	(*EvaluationTask)(nil),              // 0: eolymp.executor.EvaluationTask
	(*EvaluationTask_Generator)(nil),    // 1: eolymp.executor.EvaluationTask.Generator
	(*EvaluationTask_Run)(nil),          // 2: eolymp.executor.EvaluationTask.Run
	(*EvaluationTask_Precondition)(nil), // 3: eolymp.executor.EvaluationTask.Precondition
	(*EvaluationTask_Constraint)(nil),   // 4: eolymp.executor.EvaluationTask.Constraint
	(*Checker)(nil),                     // 5: eolymp.executor.Checker
	(*Interactor)(nil),                  // 6: eolymp.executor.Interactor
	(*File)(nil),                        // 7: eolymp.executor.File
	(*Script)(nil),                      // 8: eolymp.executor.Script
}
var file_eolymp_executor_evaluation_task_proto_depIdxs = []int32{
	3, // 0: eolymp.executor.EvaluationTask.preconditions:type_name -> eolymp.executor.EvaluationTask.Precondition
	4, // 1: eolymp.executor.EvaluationTask.constraints:type_name -> eolymp.executor.EvaluationTask.Constraint
	5, // 2: eolymp.executor.EvaluationTask.checker:type_name -> eolymp.executor.Checker
	6, // 3: eolymp.executor.EvaluationTask.interactor:type_name -> eolymp.executor.Interactor
	2, // 4: eolymp.executor.EvaluationTask.runs:type_name -> eolymp.executor.EvaluationTask.Run
	7, // 5: eolymp.executor.EvaluationTask.files:type_name -> eolymp.executor.File
	8, // 6: eolymp.executor.EvaluationTask.scripts:type_name -> eolymp.executor.Script
	1, // 7: eolymp.executor.EvaluationTask.Run.input_generator:type_name -> eolymp.executor.EvaluationTask.Generator
	1, // 8: eolymp.executor.EvaluationTask.Run.answer_generator:type_name -> eolymp.executor.EvaluationTask.Generator
	9, // [9:9] is the sub-list for method output_type
	9, // [9:9] is the sub-list for method input_type
	9, // [9:9] is the sub-list for extension type_name
	9, // [9:9] is the sub-list for extension extendee
	0, // [0:9] is the sub-list for field type_name
}

func init() { file_eolymp_executor_evaluation_task_proto_init() }
func file_eolymp_executor_evaluation_task_proto_init() {
	if File_eolymp_executor_evaluation_task_proto != nil {
		return
	}
	file_eolymp_executor_checker_proto_init()
	file_eolymp_executor_file_proto_init()
	file_eolymp_executor_interactor_proto_init()
	file_eolymp_executor_script_proto_init()
	file_eolymp_executor_evaluation_task_proto_msgTypes[2].OneofWrappers = []any{
		(*EvaluationTask_Run_InputUrl)(nil),
		(*EvaluationTask_Run_InputContent)(nil),
		(*EvaluationTask_Run_InputGenerator)(nil),
		(*EvaluationTask_Run_AnswerUrl)(nil),
		(*EvaluationTask_Run_AnswerContent)(nil),
		(*EvaluationTask_Run_AnswerGenerator)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_eolymp_executor_evaluation_task_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_eolymp_executor_evaluation_task_proto_goTypes,
		DependencyIndexes: file_eolymp_executor_evaluation_task_proto_depIdxs,
		MessageInfos:      file_eolymp_executor_evaluation_task_proto_msgTypes,
	}.Build()
	File_eolymp_executor_evaluation_task_proto = out.File
	file_eolymp_executor_evaluation_task_proto_rawDesc = nil
	file_eolymp_executor_evaluation_task_proto_goTypes = nil
	file_eolymp_executor_evaluation_task_proto_depIdxs = nil
}
