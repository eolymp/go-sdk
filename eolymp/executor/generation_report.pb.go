// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.21.12
// source: eolymp/executor/generation_report.proto

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

type GenerationReport_Run_Status int32

const (
	GenerationReport_Run_NONE     GenerationReport_Run_Status = 0
	GenerationReport_Run_PENDING  GenerationReport_Run_Status = 1
	GenerationReport_Run_COMPLETE GenerationReport_Run_Status = 2
	GenerationReport_Run_FAILED   GenerationReport_Run_Status = 3
)

// Enum value maps for GenerationReport_Run_Status.
var (
	GenerationReport_Run_Status_name = map[int32]string{
		0: "NONE",
		1: "PENDING",
		2: "COMPLETE",
		3: "FAILED",
	}
	GenerationReport_Run_Status_value = map[string]int32{
		"NONE":     0,
		"PENDING":  1,
		"COMPLETE": 2,
		"FAILED":   3,
	}
)

func (x GenerationReport_Run_Status) Enum() *GenerationReport_Run_Status {
	p := new(GenerationReport_Run_Status)
	*p = x
	return p
}

func (x GenerationReport_Run_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (GenerationReport_Run_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_eolymp_executor_generation_report_proto_enumTypes[0].Descriptor()
}

func (GenerationReport_Run_Status) Type() protoreflect.EnumType {
	return &file_eolymp_executor_generation_report_proto_enumTypes[0]
}

func (x GenerationReport_Run_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use GenerationReport_Run_Status.Descriptor instead.
func (GenerationReport_Run_Status) EnumDescriptor() ([]byte, []int) {
	return file_eolymp_executor_generation_report_proto_rawDescGZIP(), []int{0, 0, 0}
}

type GenerationReport struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TaskId    string                  `protobuf:"bytes,1,opt,name=task_id,json=taskId,proto3" json:"task_id,omitempty"`
	Reference string                  `protobuf:"bytes,2,opt,name=reference,proto3" json:"reference,omitempty"`
	Origin    string                  `protobuf:"bytes,3,opt,name=origin,proto3" json:"origin,omitempty"`
	Agent     string                  `protobuf:"bytes,4,opt,name=agent,proto3" json:"agent,omitempty"`
	Error     string                  `protobuf:"bytes,20,opt,name=error,proto3" json:"error,omitempty"`
	Runs      []*GenerationReport_Run `protobuf:"bytes,40,rep,name=runs,proto3" json:"runs,omitempty"`
}

func (x *GenerationReport) Reset() {
	*x = GenerationReport{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_executor_generation_report_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenerationReport) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenerationReport) ProtoMessage() {}

func (x *GenerationReport) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_executor_generation_report_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenerationReport.ProtoReflect.Descriptor instead.
func (*GenerationReport) Descriptor() ([]byte, []int) {
	return file_eolymp_executor_generation_report_proto_rawDescGZIP(), []int{0}
}

func (x *GenerationReport) GetTaskId() string {
	if x != nil {
		return x.TaskId
	}
	return ""
}

func (x *GenerationReport) GetReference() string {
	if x != nil {
		return x.Reference
	}
	return ""
}

func (x *GenerationReport) GetOrigin() string {
	if x != nil {
		return x.Origin
	}
	return ""
}

func (x *GenerationReport) GetAgent() string {
	if x != nil {
		return x.Agent
	}
	return ""
}

func (x *GenerationReport) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

func (x *GenerationReport) GetRuns() []*GenerationReport_Run {
	if x != nil {
		return x.Runs
	}
	return nil
}

// Run represents a single execution
type GenerationReport_Run struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Reference            string                      `protobuf:"bytes,1,opt,name=reference,proto3" json:"reference,omitempty"`
	Status               GenerationReport_Run_Status `protobuf:"varint,2,opt,name=status,proto3,enum=eolymp.executor.GenerationReport_Run_Status" json:"status,omitempty"`
	InputUrl             string                      `protobuf:"bytes,10,opt,name=input_url,json=inputUrl,proto3" json:"input_url,omitempty"`
	AnswerUrl            string                      `protobuf:"bytes,11,opt,name=answer_url,json=answerUrl,proto3" json:"answer_url,omitempty"`
	InputGeneratorStats  *Stats                      `protobuf:"bytes,20,opt,name=input_generator_stats,json=inputGeneratorStats,proto3" json:"input_generator_stats,omitempty"`
	AnswerGeneratorStats *Stats                      `protobuf:"bytes,30,opt,name=answer_generator_stats,json=answerGeneratorStats,proto3" json:"answer_generator_stats,omitempty"`
}

func (x *GenerationReport_Run) Reset() {
	*x = GenerationReport_Run{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_executor_generation_report_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenerationReport_Run) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenerationReport_Run) ProtoMessage() {}

func (x *GenerationReport_Run) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_executor_generation_report_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenerationReport_Run.ProtoReflect.Descriptor instead.
func (*GenerationReport_Run) Descriptor() ([]byte, []int) {
	return file_eolymp_executor_generation_report_proto_rawDescGZIP(), []int{0, 0}
}

func (x *GenerationReport_Run) GetReference() string {
	if x != nil {
		return x.Reference
	}
	return ""
}

func (x *GenerationReport_Run) GetStatus() GenerationReport_Run_Status {
	if x != nil {
		return x.Status
	}
	return GenerationReport_Run_NONE
}

func (x *GenerationReport_Run) GetInputUrl() string {
	if x != nil {
		return x.InputUrl
	}
	return ""
}

func (x *GenerationReport_Run) GetAnswerUrl() string {
	if x != nil {
		return x.AnswerUrl
	}
	return ""
}

func (x *GenerationReport_Run) GetInputGeneratorStats() *Stats {
	if x != nil {
		return x.InputGeneratorStats
	}
	return nil
}

func (x *GenerationReport_Run) GetAnswerGeneratorStats() *Stats {
	if x != nil {
		return x.AnswerGeneratorStats
	}
	return nil
}

var File_eolymp_executor_generation_report_proto protoreflect.FileDescriptor

var file_eolymp_executor_generation_report_proto_rawDesc = []byte{
	0x0a, 0x27, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x6f,
	0x72, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x72, 0x65, 0x70,
	0x6f, 0x72, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x65, 0x6f, 0x6c, 0x79, 0x6d,
	0x70, 0x2e, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x6f, 0x72, 0x1a, 0x1b, 0x65, 0x6f, 0x6c, 0x79,
	0x6d, 0x70, 0x2f, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x6f, 0x72, 0x2f, 0x73, 0x74, 0x61, 0x74,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc5, 0x04, 0x0a, 0x10, 0x47, 0x65, 0x6e, 0x65,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x17, 0x0a, 0x07,
	0x74, 0x61, 0x73, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74,
	0x61, 0x73, 0x6b, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e,
	0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65,
	0x6e, 0x63, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x61,
	0x67, 0x65, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x61, 0x67, 0x65, 0x6e,
	0x74, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x39, 0x0a, 0x04, 0x72, 0x75, 0x6e, 0x73, 0x18,
	0x28, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x65,
	0x78, 0x65, 0x63, 0x75, 0x74, 0x6f, 0x72, 0x2e, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x52, 0x75, 0x6e, 0x52, 0x04, 0x72, 0x75,
	0x6e, 0x73, 0x1a, 0xfa, 0x02, 0x0a, 0x03, 0x52, 0x75, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x72, 0x65,
	0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x72,
	0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x44, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2c, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d,
	0x70, 0x2e, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x6f, 0x72, 0x2e, 0x47, 0x65, 0x6e, 0x65, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x52, 0x75, 0x6e, 0x2e,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1b,
	0x0a, 0x09, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x55, 0x72, 0x6c, 0x12, 0x1d, 0x0a, 0x0a, 0x61,
	0x6e, 0x73, 0x77, 0x65, 0x72, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x61, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x55, 0x72, 0x6c, 0x12, 0x4a, 0x0a, 0x15, 0x69, 0x6e,
	0x70, 0x75, 0x74, 0x5f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x73, 0x74,
	0x61, 0x74, 0x73, 0x18, 0x14, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x65, 0x6f, 0x6c, 0x79,
	0x6d, 0x70, 0x2e, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x6f, 0x72, 0x2e, 0x53, 0x74, 0x61, 0x74,
	0x73, 0x52, 0x13, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x6f,
	0x72, 0x53, 0x74, 0x61, 0x74, 0x73, 0x12, 0x4c, 0x0a, 0x16, 0x61, 0x6e, 0x73, 0x77, 0x65, 0x72,
	0x5f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x73,
	0x18, 0x1e, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e,
	0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x6f, 0x72, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x73, 0x52, 0x14,
	0x61, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x53,
	0x74, 0x61, 0x74, 0x73, 0x22, 0x39, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x08,
	0x0a, 0x04, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x50, 0x45, 0x4e, 0x44,
	0x49, 0x4e, 0x47, 0x10, 0x01, 0x12, 0x0c, 0x0a, 0x08, 0x43, 0x4f, 0x4d, 0x50, 0x4c, 0x45, 0x54,
	0x45, 0x10, 0x02, 0x12, 0x0a, 0x0a, 0x06, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x10, 0x03, 0x42,
	0x33, 0x5a, 0x31, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6f,
	0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x67, 0x6f, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x65, 0x6f, 0x6c, 0x79,
	0x6d, 0x70, 0x2f, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x6f, 0x72, 0x3b, 0x65, 0x78, 0x65, 0x63,
	0x75, 0x74, 0x6f, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_eolymp_executor_generation_report_proto_rawDescOnce sync.Once
	file_eolymp_executor_generation_report_proto_rawDescData = file_eolymp_executor_generation_report_proto_rawDesc
)

func file_eolymp_executor_generation_report_proto_rawDescGZIP() []byte {
	file_eolymp_executor_generation_report_proto_rawDescOnce.Do(func() {
		file_eolymp_executor_generation_report_proto_rawDescData = protoimpl.X.CompressGZIP(file_eolymp_executor_generation_report_proto_rawDescData)
	})
	return file_eolymp_executor_generation_report_proto_rawDescData
}

var file_eolymp_executor_generation_report_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_eolymp_executor_generation_report_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_eolymp_executor_generation_report_proto_goTypes = []any{
	(GenerationReport_Run_Status)(0), // 0: eolymp.executor.GenerationReport.Run.Status
	(*GenerationReport)(nil),         // 1: eolymp.executor.GenerationReport
	(*GenerationReport_Run)(nil),     // 2: eolymp.executor.GenerationReport.Run
	(*Stats)(nil),                    // 3: eolymp.executor.Stats
}
var file_eolymp_executor_generation_report_proto_depIdxs = []int32{
	2, // 0: eolymp.executor.GenerationReport.runs:type_name -> eolymp.executor.GenerationReport.Run
	0, // 1: eolymp.executor.GenerationReport.Run.status:type_name -> eolymp.executor.GenerationReport.Run.Status
	3, // 2: eolymp.executor.GenerationReport.Run.input_generator_stats:type_name -> eolymp.executor.Stats
	3, // 3: eolymp.executor.GenerationReport.Run.answer_generator_stats:type_name -> eolymp.executor.Stats
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_eolymp_executor_generation_report_proto_init() }
func file_eolymp_executor_generation_report_proto_init() {
	if File_eolymp_executor_generation_report_proto != nil {
		return
	}
	file_eolymp_executor_stats_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_eolymp_executor_generation_report_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*GenerationReport); i {
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
		file_eolymp_executor_generation_report_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*GenerationReport_Run); i {
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
			RawDescriptor: file_eolymp_executor_generation_report_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_eolymp_executor_generation_report_proto_goTypes,
		DependencyIndexes: file_eolymp_executor_generation_report_proto_depIdxs,
		EnumInfos:         file_eolymp_executor_generation_report_proto_enumTypes,
		MessageInfos:      file_eolymp_executor_generation_report_proto_msgTypes,
	}.Build()
	File_eolymp_executor_generation_report_proto = out.File
	file_eolymp_executor_generation_report_proto_rawDesc = nil
	file_eolymp_executor_generation_report_proto_goTypes = nil
	file_eolymp_executor_generation_report_proto_depIdxs = nil
}
