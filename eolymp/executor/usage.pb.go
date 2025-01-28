// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.4
// 	protoc        v4.24.4
// source: eolymp/executor/usage.proto

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

type ResourceUsage struct {
	state                      protoimpl.MessageState `protogen:"open.v1"`
	UserCpuTime                uint32                 `protobuf:"varint,1,opt,name=user_cpu_time,json=userCpuTime,proto3" json:"user_cpu_time,omitempty"`                                              // user CPU time used (in milliseconds)
	SystemCpuTime              uint32                 `protobuf:"varint,2,opt,name=system_cpu_time,json=systemCpuTime,proto3" json:"system_cpu_time,omitempty"`                                        // system CPU time used (in milliseconds)
	MaxResidentSetSize         uint32                 `protobuf:"varint,3,opt,name=max_resident_set_size,json=maxResidentSetSize,proto3" json:"max_resident_set_size,omitempty"`                       // This is the maximum resident set size used (in kilobytes).
	PageReclaims               uint32                 `protobuf:"varint,4,opt,name=page_reclaims,json=pageReclaims,proto3" json:"page_reclaims,omitempty"`                                             // The number of page faults serviced without any I/O activity; here I/O activity is avoided by “reclaiming” a page frame from the list of pages awaiting reallocation.
	PageFaults                 uint32                 `protobuf:"varint,5,opt,name=page_faults,json=pageFaults,proto3" json:"page_faults,omitempty"`                                                   // The number of page faults serviced that required I/O activity.
	BlockInputOperations       uint32                 `protobuf:"varint,6,opt,name=block_input_operations,json=blockInputOperations,proto3" json:"block_input_operations,omitempty"`                   // The number of times the filesystem had to perform input.
	BlockOutputOperations      uint32                 `protobuf:"varint,7,opt,name=block_output_operations,json=blockOutputOperations,proto3" json:"block_output_operations,omitempty"`                // The number of times the filesystem had to perform output.
	VoluntaryContextSwitches   uint32                 `protobuf:"varint,8,opt,name=voluntary_context_switches,json=voluntaryContextSwitches,proto3" json:"voluntary_context_switches,omitempty"`       // The number of times a context switch resulted due to a process voluntarily giving up the processor before its time slice was completed (usually to await availability of a resource).
	InvoluntaryContextSwitches uint32                 `protobuf:"varint,9,opt,name=involuntary_context_switches,json=involuntaryContextSwitches,proto3" json:"involuntary_context_switches,omitempty"` // The number of times a context switch resulted due to a higher priority process becoming runnable or because the current process exceeded its time slice.
	unknownFields              protoimpl.UnknownFields
	sizeCache                  protoimpl.SizeCache
}

func (x *ResourceUsage) Reset() {
	*x = ResourceUsage{}
	mi := &file_eolymp_executor_usage_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ResourceUsage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResourceUsage) ProtoMessage() {}

func (x *ResourceUsage) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_executor_usage_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResourceUsage.ProtoReflect.Descriptor instead.
func (*ResourceUsage) Descriptor() ([]byte, []int) {
	return file_eolymp_executor_usage_proto_rawDescGZIP(), []int{0}
}

func (x *ResourceUsage) GetUserCpuTime() uint32 {
	if x != nil {
		return x.UserCpuTime
	}
	return 0
}

func (x *ResourceUsage) GetSystemCpuTime() uint32 {
	if x != nil {
		return x.SystemCpuTime
	}
	return 0
}

func (x *ResourceUsage) GetMaxResidentSetSize() uint32 {
	if x != nil {
		return x.MaxResidentSetSize
	}
	return 0
}

func (x *ResourceUsage) GetPageReclaims() uint32 {
	if x != nil {
		return x.PageReclaims
	}
	return 0
}

func (x *ResourceUsage) GetPageFaults() uint32 {
	if x != nil {
		return x.PageFaults
	}
	return 0
}

func (x *ResourceUsage) GetBlockInputOperations() uint32 {
	if x != nil {
		return x.BlockInputOperations
	}
	return 0
}

func (x *ResourceUsage) GetBlockOutputOperations() uint32 {
	if x != nil {
		return x.BlockOutputOperations
	}
	return 0
}

func (x *ResourceUsage) GetVoluntaryContextSwitches() uint32 {
	if x != nil {
		return x.VoluntaryContextSwitches
	}
	return 0
}

func (x *ResourceUsage) GetInvoluntaryContextSwitches() uint32 {
	if x != nil {
		return x.InvoluntaryContextSwitches
	}
	return 0
}

var File_eolymp_executor_usage_proto protoreflect.FileDescriptor

var file_eolymp_executor_usage_proto_rawDesc = string([]byte{
	0x0a, 0x1b, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x6f,
	0x72, 0x2f, 0x75, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x65,
	0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x6f, 0x72, 0x22, 0xc2,
	0x03, 0x0a, 0x0d, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x55, 0x73, 0x61, 0x67, 0x65,
	0x12, 0x22, 0x0a, 0x0d, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x63, 0x70, 0x75, 0x5f, 0x74, 0x69, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x75, 0x73, 0x65, 0x72, 0x43, 0x70, 0x75,
	0x54, 0x69, 0x6d, 0x65, 0x12, 0x26, 0x0a, 0x0f, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x5f, 0x63,
	0x70, 0x75, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0d, 0x73,
	0x79, 0x73, 0x74, 0x65, 0x6d, 0x43, 0x70, 0x75, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x31, 0x0a, 0x15,
	0x6d, 0x61, 0x78, 0x5f, 0x72, 0x65, 0x73, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x65, 0x74,
	0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x12, 0x6d, 0x61, 0x78,
	0x52, 0x65, 0x73, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x74, 0x53, 0x69, 0x7a, 0x65, 0x12,
	0x23, 0x0a, 0x0d, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x72, 0x65, 0x63, 0x6c, 0x61, 0x69, 0x6d, 0x73,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c, 0x70, 0x61, 0x67, 0x65, 0x52, 0x65, 0x63, 0x6c,
	0x61, 0x69, 0x6d, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x66, 0x61, 0x75,
	0x6c, 0x74, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x70, 0x61, 0x67, 0x65, 0x46,
	0x61, 0x75, 0x6c, 0x74, 0x73, 0x12, 0x34, 0x0a, 0x16, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x69,
	0x6e, 0x70, 0x75, 0x74, 0x5f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x14, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x49, 0x6e, 0x70, 0x75,
	0x74, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x36, 0x0a, 0x17, 0x62,
	0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x5f, 0x6f, 0x70, 0x65, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x15, 0x62, 0x6c,
	0x6f, 0x63, 0x6b, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x12, 0x3c, 0x0a, 0x1a, 0x76, 0x6f, 0x6c, 0x75, 0x6e, 0x74, 0x61, 0x72, 0x79,
	0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x5f, 0x73, 0x77, 0x69, 0x74, 0x63, 0x68, 0x65,
	0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x18, 0x76, 0x6f, 0x6c, 0x75, 0x6e, 0x74, 0x61,
	0x72, 0x79, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x53, 0x77, 0x69, 0x74, 0x63, 0x68, 0x65,
	0x73, 0x12, 0x40, 0x0a, 0x1c, 0x69, 0x6e, 0x76, 0x6f, 0x6c, 0x75, 0x6e, 0x74, 0x61, 0x72, 0x79,
	0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x5f, 0x73, 0x77, 0x69, 0x74, 0x63, 0x68, 0x65,
	0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x1a, 0x69, 0x6e, 0x76, 0x6f, 0x6c, 0x75, 0x6e,
	0x74, 0x61, 0x72, 0x79, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x53, 0x77, 0x69, 0x74, 0x63,
	0x68, 0x65, 0x73, 0x42, 0x33, 0x5a, 0x31, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x67, 0x6f, 0x2d, 0x73, 0x64, 0x6b, 0x2f,
	0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x6f, 0x72, 0x3b,
	0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x6f, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_eolymp_executor_usage_proto_rawDescOnce sync.Once
	file_eolymp_executor_usage_proto_rawDescData []byte
)

func file_eolymp_executor_usage_proto_rawDescGZIP() []byte {
	file_eolymp_executor_usage_proto_rawDescOnce.Do(func() {
		file_eolymp_executor_usage_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_eolymp_executor_usage_proto_rawDesc), len(file_eolymp_executor_usage_proto_rawDesc)))
	})
	return file_eolymp_executor_usage_proto_rawDescData
}

var file_eolymp_executor_usage_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_eolymp_executor_usage_proto_goTypes = []any{
	(*ResourceUsage)(nil), // 0: eolymp.executor.ResourceUsage
}
var file_eolymp_executor_usage_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_eolymp_executor_usage_proto_init() }
func file_eolymp_executor_usage_proto_init() {
	if File_eolymp_executor_usage_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_eolymp_executor_usage_proto_rawDesc), len(file_eolymp_executor_usage_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_eolymp_executor_usage_proto_goTypes,
		DependencyIndexes: file_eolymp_executor_usage_proto_depIdxs,
		MessageInfos:      file_eolymp_executor_usage_proto_msgTypes,
	}.Build()
	File_eolymp_executor_usage_proto = out.File
	file_eolymp_executor_usage_proto_goTypes = nil
	file_eolymp_executor_usage_proto_depIdxs = nil
}
