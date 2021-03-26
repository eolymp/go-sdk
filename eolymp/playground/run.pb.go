// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.2
// source: eolymp/playground/run.proto

package playground

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Run_Status int32

const (
	Run_NONE      Run_Status = 0 // Reserved, should not be used.
	Run_PENDING   Run_Status = 1 // Program pending execution.
	Run_EXECUTING Run_Status = 2 // Program is being executed.
	Run_EXECUTED  Run_Status = 3 // Program has terminated.
	Run_TIMEOUT   Run_Status = 4 // Program was terminated due to timeout.
	Run_OVERFLOW  Run_Status = 5 // Program was terminated due to memory overflow.
	Run_ERROR     Run_Status = 6 // Program executed with an error (see error field for details). Typically this status means compiler or linter has returned an error.
	Run_FAILURE   Run_Status = 7 // Worker has failed to execute program due to an internal error, try again in this case.
)

// Enum value maps for Run_Status.
var (
	Run_Status_name = map[int32]string{
		0: "NONE",
		1: "PENDING",
		2: "EXECUTING",
		3: "EXECUTED",
		4: "TIMEOUT",
		5: "OVERFLOW",
		6: "ERROR",
		7: "FAILURE",
	}
	Run_Status_value = map[string]int32{
		"NONE":      0,
		"PENDING":   1,
		"EXECUTING": 2,
		"EXECUTED":  3,
		"TIMEOUT":   4,
		"OVERFLOW":  5,
		"ERROR":     6,
		"FAILURE":   7,
	}
)

func (x Run_Status) Enum() *Run_Status {
	p := new(Run_Status)
	*p = x
	return p
}

func (x Run_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Run_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_eolymp_playground_run_proto_enumTypes[0].Descriptor()
}

func (Run_Status) Type() protoreflect.EnumType {
	return &file_eolymp_playground_run_proto_enumTypes[0]
}

func (x Run_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Run_Status.Descriptor instead.
func (Run_Status) EnumDescriptor() ([]byte, []int) {
	return file_eolymp_playground_run_proto_rawDescGZIP(), []int{0, 0}
}

type Run struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Task unique identifier.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Programming language.
	Lang string `protobuf:"bytes,10,opt,name=lang,proto3" json:"lang,omitempty"`
	// Source code.
	Source string `protobuf:"bytes,11,opt,name=source,proto3" json:"source,omitempty"`
	// Input data (up to 1KB)
	Input string `protobuf:"bytes,12,opt,name=input,proto3" json:"input,omitempty"`
	// Status (see enumeration values for details)
	Status Run_Status `protobuf:"varint,20,opt,name=status,proto3,enum=eolymp.playground.Run_Status" json:"status,omitempty"`
	// Error message (when status is ERROR)
	Error string `protobuf:"bytes,21,opt,name=error,proto3" json:"error,omitempty"`
	// Exit code.
	ExitCode int32 `protobuf:"varint,30,opt,name=exit_code,json=exitCode,proto3" json:"exit_code,omitempty"`
	// Signal used to kill process, if any.
	Signal int32 `protobuf:"varint,31,opt,name=signal,proto3" json:"signal,omitempty"`
	// Amount of time in milliseconds program was executed.
	TimeUsage int32 `protobuf:"varint,32,opt,name=time_usage,json=timeUsage,proto3" json:"time_usage,omitempty"`
	// Peak amount of memory in bytes used during program execution.
	MemoryUsage int32 `protobuf:"varint,33,opt,name=memory_usage,json=memoryUsage,proto3" json:"memory_usage,omitempty"`
	// Combined output (stdout + stderr) generated by the program (only first 1KB, the rest is truncated).
	Output string `protobuf:"bytes,40,opt,name=output,proto3" json:"output,omitempty"`
}

func (x *Run) Reset() {
	*x = Run{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_playground_run_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Run) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Run) ProtoMessage() {}

func (x *Run) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_playground_run_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Run.ProtoReflect.Descriptor instead.
func (*Run) Descriptor() ([]byte, []int) {
	return file_eolymp_playground_run_proto_rawDescGZIP(), []int{0}
}

func (x *Run) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Run) GetLang() string {
	if x != nil {
		return x.Lang
	}
	return ""
}

func (x *Run) GetSource() string {
	if x != nil {
		return x.Source
	}
	return ""
}

func (x *Run) GetInput() string {
	if x != nil {
		return x.Input
	}
	return ""
}

func (x *Run) GetStatus() Run_Status {
	if x != nil {
		return x.Status
	}
	return Run_NONE
}

func (x *Run) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

func (x *Run) GetExitCode() int32 {
	if x != nil {
		return x.ExitCode
	}
	return 0
}

func (x *Run) GetSignal() int32 {
	if x != nil {
		return x.Signal
	}
	return 0
}

func (x *Run) GetTimeUsage() int32 {
	if x != nil {
		return x.TimeUsage
	}
	return 0
}

func (x *Run) GetMemoryUsage() int32 {
	if x != nil {
		return x.MemoryUsage
	}
	return 0
}

func (x *Run) GetOutput() string {
	if x != nil {
		return x.Output
	}
	return ""
}

var File_eolymp_playground_run_proto protoreflect.FileDescriptor

var file_eolymp_playground_run_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x70, 0x6c, 0x61, 0x79, 0x67, 0x72, 0x6f,
	0x75, 0x6e, 0x64, 0x2f, 0x72, 0x75, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11, 0x65,
	0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x70, 0x6c, 0x61, 0x79, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64,
	0x22, 0xa4, 0x03, 0x0a, 0x03, 0x52, 0x75, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x61, 0x6e, 0x67,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6c, 0x61, 0x6e, 0x67, 0x12, 0x16, 0x0a, 0x06,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x18, 0x0c, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x35, 0x0a, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x14, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1d, 0x2e, 0x65, 0x6f, 0x6c,
	0x79, 0x6d, 0x70, 0x2e, 0x70, 0x6c, 0x61, 0x79, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x2e, 0x52,
	0x75, 0x6e, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x15, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x1b, 0x0a, 0x09, 0x65, 0x78, 0x69, 0x74, 0x5f,
	0x63, 0x6f, 0x64, 0x65, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x65, 0x78, 0x69, 0x74,
	0x43, 0x6f, 0x64, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x18, 0x1f,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x12, 0x1d, 0x0a, 0x0a,
	0x74, 0x69, 0x6d, 0x65, 0x5f, 0x75, 0x73, 0x61, 0x67, 0x65, 0x18, 0x20, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x55, 0x73, 0x61, 0x67, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x6d,
	0x65, 0x6d, 0x6f, 0x72, 0x79, 0x5f, 0x75, 0x73, 0x61, 0x67, 0x65, 0x18, 0x21, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x0b, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x55, 0x73, 0x61, 0x67, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x18, 0x28, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x22, 0x6f, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x08, 0x0a, 0x04, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x50, 0x45,
	0x4e, 0x44, 0x49, 0x4e, 0x47, 0x10, 0x01, 0x12, 0x0d, 0x0a, 0x09, 0x45, 0x58, 0x45, 0x43, 0x55,
	0x54, 0x49, 0x4e, 0x47, 0x10, 0x02, 0x12, 0x0c, 0x0a, 0x08, 0x45, 0x58, 0x45, 0x43, 0x55, 0x54,
	0x45, 0x44, 0x10, 0x03, 0x12, 0x0b, 0x0a, 0x07, 0x54, 0x49, 0x4d, 0x45, 0x4f, 0x55, 0x54, 0x10,
	0x04, 0x12, 0x0c, 0x0a, 0x08, 0x4f, 0x56, 0x45, 0x52, 0x46, 0x4c, 0x4f, 0x57, 0x10, 0x05, 0x12,
	0x09, 0x0a, 0x05, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x06, 0x12, 0x0b, 0x0a, 0x07, 0x46, 0x41,
	0x49, 0x4c, 0x55, 0x52, 0x45, 0x10, 0x07, 0x42, 0x3d, 0x5a, 0x3b, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x63, 0x6f, 0x6e,
	0x74, 0x72, 0x61, 0x63, 0x74, 0x73, 0x2f, 0x67, 0x6f, 0x2f, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70,
	0x2f, 0x70, 0x6c, 0x61, 0x79, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x3b, 0x70, 0x6c, 0x61, 0x79,
	0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_eolymp_playground_run_proto_rawDescOnce sync.Once
	file_eolymp_playground_run_proto_rawDescData = file_eolymp_playground_run_proto_rawDesc
)

func file_eolymp_playground_run_proto_rawDescGZIP() []byte {
	file_eolymp_playground_run_proto_rawDescOnce.Do(func() {
		file_eolymp_playground_run_proto_rawDescData = protoimpl.X.CompressGZIP(file_eolymp_playground_run_proto_rawDescData)
	})
	return file_eolymp_playground_run_proto_rawDescData
}

var file_eolymp_playground_run_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_eolymp_playground_run_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_eolymp_playground_run_proto_goTypes = []interface{}{
	(Run_Status)(0), // 0: eolymp.playground.Run.Status
	(*Run)(nil),     // 1: eolymp.playground.Run
}
var file_eolymp_playground_run_proto_depIdxs = []int32{
	0, // 0: eolymp.playground.Run.status:type_name -> eolymp.playground.Run.Status
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_eolymp_playground_run_proto_init() }
func file_eolymp_playground_run_proto_init() {
	if File_eolymp_playground_run_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_eolymp_playground_run_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Run); i {
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
			RawDescriptor: file_eolymp_playground_run_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_eolymp_playground_run_proto_goTypes,
		DependencyIndexes: file_eolymp_playground_run_proto_depIdxs,
		EnumInfos:         file_eolymp_playground_run_proto_enumTypes,
		MessageInfos:      file_eolymp_playground_run_proto_msgTypes,
	}.Build()
	File_eolymp_playground_run_proto = out.File
	file_eolymp_playground_run_proto_rawDesc = nil
	file_eolymp_playground_run_proto_goTypes = nil
	file_eolymp_playground_run_proto_depIdxs = nil
}
