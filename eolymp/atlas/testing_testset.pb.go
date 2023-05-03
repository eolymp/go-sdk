// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.18.1
// source: eolymp/atlas/testing_testset.proto

package atlas

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

type Testset struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Unique identifier. Keep empty when creating new testset.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Problem where this testset belongs. Keep empty when creating new testset.
	ProblemId string `protobuf:"bytes,2,opt,name=problem_id,json=problemId,proto3" json:"problem_id,omitempty"`
	// Index, a sequential number, of the testset within a problem.
	Index uint32 `protobuf:"varint,3,opt,name=index,proto3" json:"index,omitempty"`
	// Time limit in milliseconds for runs within this testset.
	TimeLimit uint32 `protobuf:"varint,10,opt,name=time_limit,json=timeLimit,proto3" json:"time_limit,omitempty"`
	// Time limit in milliseconds for CPU usage within this testset.
	CpuLimit uint32 `protobuf:"varint,13,opt,name=cpu_limit,json=cpuLimit,proto3" json:"cpu_limit,omitempty"`
	// Memory limit in bytes for runs within this testset.
	MemoryLimit uint64 `protobuf:"varint,11,opt,name=memory_limit,json=memoryLimit,proto3" json:"memory_limit,omitempty"`
	// File size limit in bytes for runs within this testset.
	FileSizeLimit uint64 `protobuf:"varint,12,opt,name=file_size_limit,json=fileSizeLimit,proto3" json:"file_size_limit,omitempty"`
	// List of other testset indices which must pass before this testset is executed.
	Dependencies []uint32 `protobuf:"varint,20,rep,packed,name=dependencies,proto3" json:"dependencies,omitempty"`
	// Score mode defines how score points are awarded for this testset. See ScoringMode enumeration for details.
	ScoringMode ScoringMode `protobuf:"varint,30,opt,name=scoring_mode,json=scoringMode,proto3,enum=eolymp.atlas.ScoringMode" json:"scoring_mode,omitempty"`
	// Feedback policy defines how result will be shown back to the user
	FeedbackPolicy FeedbackPolicy `protobuf:"varint,40,opt,name=feedback_policy,json=feedbackPolicy,proto3,enum=eolymp.atlas.FeedbackPolicy" json:"feedback_policy,omitempty"`
}

func (x *Testset) Reset() {
	*x = Testset{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_atlas_testing_testset_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Testset) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Testset) ProtoMessage() {}

func (x *Testset) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_atlas_testing_testset_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Testset.ProtoReflect.Descriptor instead.
func (*Testset) Descriptor() ([]byte, []int) {
	return file_eolymp_atlas_testing_testset_proto_rawDescGZIP(), []int{0}
}

func (x *Testset) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Testset) GetProblemId() string {
	if x != nil {
		return x.ProblemId
	}
	return ""
}

func (x *Testset) GetIndex() uint32 {
	if x != nil {
		return x.Index
	}
	return 0
}

func (x *Testset) GetTimeLimit() uint32 {
	if x != nil {
		return x.TimeLimit
	}
	return 0
}

func (x *Testset) GetCpuLimit() uint32 {
	if x != nil {
		return x.CpuLimit
	}
	return 0
}

func (x *Testset) GetMemoryLimit() uint64 {
	if x != nil {
		return x.MemoryLimit
	}
	return 0
}

func (x *Testset) GetFileSizeLimit() uint64 {
	if x != nil {
		return x.FileSizeLimit
	}
	return 0
}

func (x *Testset) GetDependencies() []uint32 {
	if x != nil {
		return x.Dependencies
	}
	return nil
}

func (x *Testset) GetScoringMode() ScoringMode {
	if x != nil {
		return x.ScoringMode
	}
	return ScoringMode_NO_SCORE
}

func (x *Testset) GetFeedbackPolicy() FeedbackPolicy {
	if x != nil {
		return x.FeedbackPolicy
	}
	return FeedbackPolicy_COMPLETE
}

var File_eolymp_atlas_testing_testset_proto protoreflect.FileDescriptor

var file_eolymp_atlas_testing_testset_proto_rawDesc = []byte{
	0x0a, 0x22, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x61, 0x74, 0x6c, 0x61, 0x73, 0x2f, 0x74,
	0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x74, 0x65, 0x73, 0x74, 0x73, 0x65, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x61, 0x74, 0x6c,
	0x61, 0x73, 0x1a, 0x23, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x61, 0x74, 0x6c, 0x61, 0x73,
	0x2f, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x66, 0x65, 0x65, 0x64, 0x62, 0x61, 0x63,
	0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x22, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f,
	0x61, 0x74, 0x6c, 0x61, 0x73, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x73, 0x63,
	0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xfe, 0x02, 0x0a, 0x07,
	0x54, 0x65, 0x73, 0x74, 0x73, 0x65, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x62, 0x6c,
	0x65, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f,
	0x62, 0x6c, 0x65, 0x6d, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x1d, 0x0a, 0x0a,
	0x74, 0x69, 0x6d, 0x65, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x63,
	0x70, 0x75, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08,
	0x63, 0x70, 0x75, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x6d, 0x65, 0x6d, 0x6f,
	0x72, 0x79, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0b,
	0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x26, 0x0a, 0x0f, 0x66,
	0x69, 0x6c, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x0c,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x0d, 0x66, 0x69, 0x6c, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x4c, 0x69,
	0x6d, 0x69, 0x74, 0x12, 0x22, 0x0a, 0x0c, 0x64, 0x65, 0x70, 0x65, 0x6e, 0x64, 0x65, 0x6e, 0x63,
	0x69, 0x65, 0x73, 0x18, 0x14, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x0c, 0x64, 0x65, 0x70, 0x65, 0x6e,
	0x64, 0x65, 0x6e, 0x63, 0x69, 0x65, 0x73, 0x12, 0x3c, 0x0a, 0x0c, 0x73, 0x63, 0x6f, 0x72, 0x69,
	0x6e, 0x67, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x19, 0x2e,
	0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x61, 0x74, 0x6c, 0x61, 0x73, 0x2e, 0x53, 0x63, 0x6f,
	0x72, 0x69, 0x6e, 0x67, 0x4d, 0x6f, 0x64, 0x65, 0x52, 0x0b, 0x73, 0x63, 0x6f, 0x72, 0x69, 0x6e,
	0x67, 0x4d, 0x6f, 0x64, 0x65, 0x12, 0x45, 0x0a, 0x0f, 0x66, 0x65, 0x65, 0x64, 0x62, 0x61, 0x63,
	0x6b, 0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x18, 0x28, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1c,
	0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x61, 0x74, 0x6c, 0x61, 0x73, 0x2e, 0x46, 0x65,
	0x65, 0x64, 0x62, 0x61, 0x63, 0x6b, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x0e, 0x66, 0x65,
	0x65, 0x64, 0x62, 0x61, 0x63, 0x6b, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x42, 0x2d, 0x5a, 0x2b,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6f, 0x6c, 0x79, 0x6d,
	0x70, 0x2f, 0x67, 0x6f, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f,
	0x61, 0x74, 0x6c, 0x61, 0x73, 0x3b, 0x61, 0x74, 0x6c, 0x61, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_eolymp_atlas_testing_testset_proto_rawDescOnce sync.Once
	file_eolymp_atlas_testing_testset_proto_rawDescData = file_eolymp_atlas_testing_testset_proto_rawDesc
)

func file_eolymp_atlas_testing_testset_proto_rawDescGZIP() []byte {
	file_eolymp_atlas_testing_testset_proto_rawDescOnce.Do(func() {
		file_eolymp_atlas_testing_testset_proto_rawDescData = protoimpl.X.CompressGZIP(file_eolymp_atlas_testing_testset_proto_rawDescData)
	})
	return file_eolymp_atlas_testing_testset_proto_rawDescData
}

var file_eolymp_atlas_testing_testset_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_eolymp_atlas_testing_testset_proto_goTypes = []interface{}{
	(*Testset)(nil),     // 0: eolymp.atlas.Testset
	(ScoringMode)(0),    // 1: eolymp.atlas.ScoringMode
	(FeedbackPolicy)(0), // 2: eolymp.atlas.FeedbackPolicy
}
var file_eolymp_atlas_testing_testset_proto_depIdxs = []int32{
	1, // 0: eolymp.atlas.Testset.scoring_mode:type_name -> eolymp.atlas.ScoringMode
	2, // 1: eolymp.atlas.Testset.feedback_policy:type_name -> eolymp.atlas.FeedbackPolicy
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_eolymp_atlas_testing_testset_proto_init() }
func file_eolymp_atlas_testing_testset_proto_init() {
	if File_eolymp_atlas_testing_testset_proto != nil {
		return
	}
	file_eolymp_atlas_testing_feedback_proto_init()
	file_eolymp_atlas_testing_scoring_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_eolymp_atlas_testing_testset_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Testset); i {
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
			RawDescriptor: file_eolymp_atlas_testing_testset_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_eolymp_atlas_testing_testset_proto_goTypes,
		DependencyIndexes: file_eolymp_atlas_testing_testset_proto_depIdxs,
		MessageInfos:      file_eolymp_atlas_testing_testset_proto_msgTypes,
	}.Build()
	File_eolymp_atlas_testing_testset_proto = out.File
	file_eolymp_atlas_testing_testset_proto_rawDesc = nil
	file_eolymp_atlas_testing_testset_proto_goTypes = nil
	file_eolymp_atlas_testing_testset_proto_depIdxs = nil
}