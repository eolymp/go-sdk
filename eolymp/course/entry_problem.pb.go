// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v3.18.1
// source: eolymp/course/entry_problem.proto

package course

import (
	atlas "github.com/eolymp/go-sdk/eolymp/atlas"
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

type Problem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProblemUrl             string               `protobuf:"bytes,1,opt,name=problem_url,json=problemUrl,proto3" json:"problem_url,omitempty"`
	ProblemId              string               `protobuf:"bytes,2,opt,name=problem_id,json=problemId,proto3" json:"problem_id,omitempty"`
	Cost                   uint32               `protobuf:"varint,10,opt,name=cost,proto3" json:"cost,omitempty"`
	SubmitLimit            uint32               `protobuf:"varint,20,opt,name=submit_limit,json=submitLimit,proto3" json:"submit_limit,omitempty"`
	OverrideFeedbackPolicy atlas.FeedbackPolicy `protobuf:"varint,21,opt,name=override_feedback_policy,json=overrideFeedbackPolicy,proto3,enum=eolymp.atlas.FeedbackPolicy" json:"override_feedback_policy,omitempty"`
}

func (x *Problem) Reset() {
	*x = Problem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_course_entry_problem_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Problem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Problem) ProtoMessage() {}

func (x *Problem) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_course_entry_problem_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Problem.ProtoReflect.Descriptor instead.
func (*Problem) Descriptor() ([]byte, []int) {
	return file_eolymp_course_entry_problem_proto_rawDescGZIP(), []int{0}
}

func (x *Problem) GetProblemUrl() string {
	if x != nil {
		return x.ProblemUrl
	}
	return ""
}

func (x *Problem) GetProblemId() string {
	if x != nil {
		return x.ProblemId
	}
	return ""
}

func (x *Problem) GetCost() uint32 {
	if x != nil {
		return x.Cost
	}
	return 0
}

func (x *Problem) GetSubmitLimit() uint32 {
	if x != nil {
		return x.SubmitLimit
	}
	return 0
}

func (x *Problem) GetOverrideFeedbackPolicy() atlas.FeedbackPolicy {
	if x != nil {
		return x.OverrideFeedbackPolicy
	}
	return atlas.FeedbackPolicy(0)
}

var File_eolymp_course_entry_problem_proto protoreflect.FileDescriptor

var file_eolymp_course_entry_problem_proto_rawDesc = []byte{
	0x0a, 0x21, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x2f,
	0x65, 0x6e, 0x74, 0x72, 0x79, 0x5f, 0x70, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x63, 0x6f, 0x75, 0x72,
	0x73, 0x65, 0x1a, 0x23, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x61, 0x74, 0x6c, 0x61, 0x73,
	0x2f, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x66, 0x65, 0x65, 0x64, 0x62, 0x61, 0x63,
	0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd8, 0x01, 0x0a, 0x07, 0x50, 0x72, 0x6f, 0x62,
	0x6c, 0x65, 0x6d, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x5f, 0x75,
	0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x72, 0x6f, 0x62, 0x6c, 0x65,
	0x6d, 0x55, 0x72, 0x6c, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x62, 0x6c, 0x65,
	0x6d, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x73, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x04, 0x63, 0x6f, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x75, 0x62, 0x6d, 0x69,
	0x74, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x14, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x73,
	0x75, 0x62, 0x6d, 0x69, 0x74, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x56, 0x0a, 0x18, 0x6f, 0x76,
	0x65, 0x72, 0x72, 0x69, 0x64, 0x65, 0x5f, 0x66, 0x65, 0x65, 0x64, 0x62, 0x61, 0x63, 0x6b, 0x5f,
	0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x18, 0x15, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1c, 0x2e, 0x65,
	0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x61, 0x74, 0x6c, 0x61, 0x73, 0x2e, 0x46, 0x65, 0x65, 0x64,
	0x62, 0x61, 0x63, 0x6b, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x16, 0x6f, 0x76, 0x65, 0x72,
	0x72, 0x69, 0x64, 0x65, 0x46, 0x65, 0x65, 0x64, 0x62, 0x61, 0x63, 0x6b, 0x50, 0x6f, 0x6c, 0x69,
	0x63, 0x79, 0x42, 0x2f, 0x5a, 0x2d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x67, 0x6f, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x65,
	0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x3b, 0x63, 0x6f, 0x75,
	0x72, 0x73, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_eolymp_course_entry_problem_proto_rawDescOnce sync.Once
	file_eolymp_course_entry_problem_proto_rawDescData = file_eolymp_course_entry_problem_proto_rawDesc
)

func file_eolymp_course_entry_problem_proto_rawDescGZIP() []byte {
	file_eolymp_course_entry_problem_proto_rawDescOnce.Do(func() {
		file_eolymp_course_entry_problem_proto_rawDescData = protoimpl.X.CompressGZIP(file_eolymp_course_entry_problem_proto_rawDescData)
	})
	return file_eolymp_course_entry_problem_proto_rawDescData
}

var file_eolymp_course_entry_problem_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_eolymp_course_entry_problem_proto_goTypes = []interface{}{
	(*Problem)(nil),           // 0: eolymp.course.Problem
	(atlas.FeedbackPolicy)(0), // 1: eolymp.atlas.FeedbackPolicy
}
var file_eolymp_course_entry_problem_proto_depIdxs = []int32{
	1, // 0: eolymp.course.Problem.override_feedback_policy:type_name -> eolymp.atlas.FeedbackPolicy
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_eolymp_course_entry_problem_proto_init() }
func file_eolymp_course_entry_problem_proto_init() {
	if File_eolymp_course_entry_problem_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_eolymp_course_entry_problem_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Problem); i {
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
			RawDescriptor: file_eolymp_course_entry_problem_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_eolymp_course_entry_problem_proto_goTypes,
		DependencyIndexes: file_eolymp_course_entry_problem_proto_depIdxs,
		MessageInfos:      file_eolymp_course_entry_problem_proto_msgTypes,
	}.Build()
	File_eolymp_course_entry_problem_proto = out.File
	file_eolymp_course_entry_problem_proto_rawDesc = nil
	file_eolymp_course_entry_problem_proto_goTypes = nil
	file_eolymp_course_entry_problem_proto_depIdxs = nil
}
