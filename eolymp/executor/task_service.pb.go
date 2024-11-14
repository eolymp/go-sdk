// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v3.21.12
// source: eolymp/executor/task_service.proto

package executor

import (
	_ "github.com/eolymp/go-sdk/eolymp/annotations"
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

type CreateTaskInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Task:
	//
	//	*CreateTaskInput_Evaluation
	//	*CreateTaskInput_Generation
	Task isCreateTaskInput_Task `protobuf_oneof:"task"`
}

func (x *CreateTaskInput) Reset() {
	*x = CreateTaskInput{}
	mi := &file_eolymp_executor_task_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateTaskInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTaskInput) ProtoMessage() {}

func (x *CreateTaskInput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_executor_task_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTaskInput.ProtoReflect.Descriptor instead.
func (*CreateTaskInput) Descriptor() ([]byte, []int) {
	return file_eolymp_executor_task_service_proto_rawDescGZIP(), []int{0}
}

func (m *CreateTaskInput) GetTask() isCreateTaskInput_Task {
	if m != nil {
		return m.Task
	}
	return nil
}

func (x *CreateTaskInput) GetEvaluation() *EvaluationTask {
	if x, ok := x.GetTask().(*CreateTaskInput_Evaluation); ok {
		return x.Evaluation
	}
	return nil
}

func (x *CreateTaskInput) GetGeneration() *GenerationTask {
	if x, ok := x.GetTask().(*CreateTaskInput_Generation); ok {
		return x.Generation
	}
	return nil
}

type isCreateTaskInput_Task interface {
	isCreateTaskInput_Task()
}

type CreateTaskInput_Evaluation struct {
	Evaluation *EvaluationTask `protobuf:"bytes,1,opt,name=evaluation,proto3,oneof"`
}

type CreateTaskInput_Generation struct {
	Generation *GenerationTask `protobuf:"bytes,2,opt,name=generation,proto3,oneof"`
}

func (*CreateTaskInput_Evaluation) isCreateTaskInput_Task() {}

func (*CreateTaskInput_Generation) isCreateTaskInput_Task() {}

type CreateTaskOutput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TaskId string `protobuf:"bytes,1,opt,name=task_id,json=taskId,proto3" json:"task_id,omitempty"`
}

func (x *CreateTaskOutput) Reset() {
	*x = CreateTaskOutput{}
	mi := &file_eolymp_executor_task_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateTaskOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTaskOutput) ProtoMessage() {}

func (x *CreateTaskOutput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_executor_task_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTaskOutput.ProtoReflect.Descriptor instead.
func (*CreateTaskOutput) Descriptor() ([]byte, []int) {
	return file_eolymp_executor_task_service_proto_rawDescGZIP(), []int{1}
}

func (x *CreateTaskOutput) GetTaskId() string {
	if x != nil {
		return x.TaskId
	}
	return ""
}

var File_eolymp_executor_task_service_proto protoreflect.FileDescriptor

var file_eolymp_executor_task_service_proto_rawDesc = []byte{
	0x0a, 0x22, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x6f,
	0x72, 0x2f, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x65, 0x78, 0x65,
	0x63, 0x75, 0x74, 0x6f, 0x72, 0x1a, 0x22, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x61, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x72, 0x61, 0x74, 0x65, 0x6c, 0x69,
	0x6d, 0x69, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x25, 0x65, 0x6f, 0x6c, 0x79, 0x6d,
	0x70, 0x2f, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x6f, 0x72, 0x2f, 0x65, 0x76, 0x61, 0x6c, 0x75,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x25, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x6f,
	0x72, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x61, 0x73,
	0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9f, 0x01, 0x0a, 0x0f, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x41, 0x0a, 0x0a, 0x65,
	0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1f, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x6f,
	0x72, 0x2e, 0x45, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x61, 0x73, 0x6b,
	0x48, 0x00, 0x52, 0x0a, 0x65, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x41,
	0x0a, 0x0a, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x65, 0x78, 0x65, 0x63,
	0x75, 0x74, 0x6f, 0x72, 0x2e, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54,
	0x61, 0x73, 0x6b, 0x48, 0x00, 0x52, 0x0a, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x42, 0x06, 0x0a, 0x04, 0x74, 0x61, 0x73, 0x6b, 0x22, 0x2b, 0x0a, 0x10, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x12, 0x17, 0x0a,
	0x07, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x74, 0x61, 0x73, 0x6b, 0x49, 0x64, 0x32, 0x72, 0x0a, 0x0b, 0x54, 0x61, 0x73, 0x6b, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x63, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54,
	0x61, 0x73, 0x6b, 0x12, 0x20, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x65, 0x78, 0x65,
	0x63, 0x75, 0x74, 0x6f, 0x72, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x61, 0x73, 0x6b,
	0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x21, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x65,
	0x78, 0x65, 0x63, 0x75, 0x74, 0x6f, 0x72, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x61,
	0x73, 0x6b, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x22, 0x10, 0xea, 0xe2, 0x0a, 0x0c, 0xf5, 0xe2,
	0x0a, 0x00, 0x00, 0xa0, 0x41, 0xf8, 0xe2, 0x0a, 0xc8, 0x01, 0x42, 0x33, 0x5a, 0x31, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f,
	0x67, 0x6f, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x65, 0x78,
	0x65, 0x63, 0x75, 0x74, 0x6f, 0x72, 0x3b, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x6f, 0x72, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_eolymp_executor_task_service_proto_rawDescOnce sync.Once
	file_eolymp_executor_task_service_proto_rawDescData = file_eolymp_executor_task_service_proto_rawDesc
)

func file_eolymp_executor_task_service_proto_rawDescGZIP() []byte {
	file_eolymp_executor_task_service_proto_rawDescOnce.Do(func() {
		file_eolymp_executor_task_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_eolymp_executor_task_service_proto_rawDescData)
	})
	return file_eolymp_executor_task_service_proto_rawDescData
}

var file_eolymp_executor_task_service_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_eolymp_executor_task_service_proto_goTypes = []any{
	(*CreateTaskInput)(nil),  // 0: eolymp.executor.CreateTaskInput
	(*CreateTaskOutput)(nil), // 1: eolymp.executor.CreateTaskOutput
	(*EvaluationTask)(nil),   // 2: eolymp.executor.EvaluationTask
	(*GenerationTask)(nil),   // 3: eolymp.executor.GenerationTask
}
var file_eolymp_executor_task_service_proto_depIdxs = []int32{
	2, // 0: eolymp.executor.CreateTaskInput.evaluation:type_name -> eolymp.executor.EvaluationTask
	3, // 1: eolymp.executor.CreateTaskInput.generation:type_name -> eolymp.executor.GenerationTask
	0, // 2: eolymp.executor.TaskService.CreateTask:input_type -> eolymp.executor.CreateTaskInput
	1, // 3: eolymp.executor.TaskService.CreateTask:output_type -> eolymp.executor.CreateTaskOutput
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_eolymp_executor_task_service_proto_init() }
func file_eolymp_executor_task_service_proto_init() {
	if File_eolymp_executor_task_service_proto != nil {
		return
	}
	file_eolymp_executor_evaluation_task_proto_init()
	file_eolymp_executor_generation_task_proto_init()
	file_eolymp_executor_task_service_proto_msgTypes[0].OneofWrappers = []any{
		(*CreateTaskInput_Evaluation)(nil),
		(*CreateTaskInput_Generation)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_eolymp_executor_task_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_eolymp_executor_task_service_proto_goTypes,
		DependencyIndexes: file_eolymp_executor_task_service_proto_depIdxs,
		MessageInfos:      file_eolymp_executor_task_service_proto_msgTypes,
	}.Build()
	File_eolymp_executor_task_service_proto = out.File
	file_eolymp_executor_task_service_proto_rawDesc = nil
	file_eolymp_executor_task_service_proto_goTypes = nil
	file_eolymp_executor_task_service_proto_depIdxs = nil
}
