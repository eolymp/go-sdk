// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.29.4
// source: eolymp/atlas/scoring_service.proto

package atlas

import (
	_ "github.com/eolymp/go-sdk/eolymp/annotations"
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

type ScoreChangedEvent struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ProblemId     string                 `protobuf:"bytes,3,opt,name=problem_id,json=problemId,proto3" json:"problem_id,omitempty"`
	MemberId      string                 `protobuf:"bytes,4,opt,name=member_id,json=memberId,proto3" json:"member_id,omitempty"`
	Before        *Score                 `protobuf:"bytes,1,opt,name=before,proto3" json:"before,omitempty"`
	After         *Score                 `protobuf:"bytes,2,opt,name=after,proto3" json:"after,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ScoreChangedEvent) Reset() {
	*x = ScoreChangedEvent{}
	mi := &file_eolymp_atlas_scoring_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ScoreChangedEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ScoreChangedEvent) ProtoMessage() {}

func (x *ScoreChangedEvent) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_atlas_scoring_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ScoreChangedEvent.ProtoReflect.Descriptor instead.
func (*ScoreChangedEvent) Descriptor() ([]byte, []int) {
	return file_eolymp_atlas_scoring_service_proto_rawDescGZIP(), []int{0}
}

func (x *ScoreChangedEvent) GetProblemId() string {
	if x != nil {
		return x.ProblemId
	}
	return ""
}

func (x *ScoreChangedEvent) GetMemberId() string {
	if x != nil {
		return x.MemberId
	}
	return ""
}

func (x *ScoreChangedEvent) GetBefore() *Score {
	if x != nil {
		return x.Before
	}
	return nil
}

func (x *ScoreChangedEvent) GetAfter() *Score {
	if x != nil {
		return x.After
	}
	return nil
}

type DescribeScoreInput struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	MemberId      string                 `protobuf:"bytes,2,opt,name=member_id,json=memberId,proto3" json:"member_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DescribeScoreInput) Reset() {
	*x = DescribeScoreInput{}
	mi := &file_eolymp_atlas_scoring_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DescribeScoreInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescribeScoreInput) ProtoMessage() {}

func (x *DescribeScoreInput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_atlas_scoring_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DescribeScoreInput.ProtoReflect.Descriptor instead.
func (*DescribeScoreInput) Descriptor() ([]byte, []int) {
	return file_eolymp_atlas_scoring_service_proto_rawDescGZIP(), []int{1}
}

func (x *DescribeScoreInput) GetMemberId() string {
	if x != nil {
		return x.MemberId
	}
	return ""
}

type DescribeScoreOutput struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Score         *Score                 `protobuf:"bytes,1,opt,name=score,proto3" json:"score,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DescribeScoreOutput) Reset() {
	*x = DescribeScoreOutput{}
	mi := &file_eolymp_atlas_scoring_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DescribeScoreOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescribeScoreOutput) ProtoMessage() {}

func (x *DescribeScoreOutput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_atlas_scoring_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DescribeScoreOutput.ProtoReflect.Descriptor instead.
func (*DescribeScoreOutput) Descriptor() ([]byte, []int) {
	return file_eolymp_atlas_scoring_service_proto_rawDescGZIP(), []int{2}
}

func (x *DescribeScoreOutput) GetScore() *Score {
	if x != nil {
		return x.Score
	}
	return nil
}

type DescribeProblemGradingInput struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DescribeProblemGradingInput) Reset() {
	*x = DescribeProblemGradingInput{}
	mi := &file_eolymp_atlas_scoring_service_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DescribeProblemGradingInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescribeProblemGradingInput) ProtoMessage() {}

func (x *DescribeProblemGradingInput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_atlas_scoring_service_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DescribeProblemGradingInput.ProtoReflect.Descriptor instead.
func (*DescribeProblemGradingInput) Descriptor() ([]byte, []int) {
	return file_eolymp_atlas_scoring_service_proto_rawDescGZIP(), []int{3}
}

type DescribeProblemGradingOutput struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Problem grade ranges from highest grade to lowest
	// to grade a submission, iterate over grade ranges and compare resource_usage value with upper_bound,
	// first grade which has upper_bound higher or equal to resource usage is the grade for submission.
	//
	//	for _, range := range ranges {
	//	   if range.GetUpperBound() >= submission.GetResourceUsage() {
	//	       return range.GetGrade()
	//	   }
	//	}
	Ranges        []*DescribeProblemGradingOutput_Range `protobuf:"bytes,2,rep,name=ranges,proto3" json:"ranges,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DescribeProblemGradingOutput) Reset() {
	*x = DescribeProblemGradingOutput{}
	mi := &file_eolymp_atlas_scoring_service_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DescribeProblemGradingOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescribeProblemGradingOutput) ProtoMessage() {}

func (x *DescribeProblemGradingOutput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_atlas_scoring_service_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DescribeProblemGradingOutput.ProtoReflect.Descriptor instead.
func (*DescribeProblemGradingOutput) Descriptor() ([]byte, []int) {
	return file_eolymp_atlas_scoring_service_proto_rawDescGZIP(), []int{4}
}

func (x *DescribeProblemGradingOutput) GetRanges() []*DescribeProblemGradingOutput_Range {
	if x != nil {
		return x.Ranges
	}
	return nil
}

type DescribeProblemGradingOutput_Range struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Grade         uint32                 `protobuf:"varint,1,opt,name=grade,proto3" json:"grade,omitempty"`
	UpperBound    float32                `protobuf:"fixed32,2,opt,name=upper_bound,json=upperBound,proto3" json:"upper_bound,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DescribeProblemGradingOutput_Range) Reset() {
	*x = DescribeProblemGradingOutput_Range{}
	mi := &file_eolymp_atlas_scoring_service_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DescribeProblemGradingOutput_Range) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescribeProblemGradingOutput_Range) ProtoMessage() {}

func (x *DescribeProblemGradingOutput_Range) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_atlas_scoring_service_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DescribeProblemGradingOutput_Range.ProtoReflect.Descriptor instead.
func (*DescribeProblemGradingOutput_Range) Descriptor() ([]byte, []int) {
	return file_eolymp_atlas_scoring_service_proto_rawDescGZIP(), []int{4, 0}
}

func (x *DescribeProblemGradingOutput_Range) GetGrade() uint32 {
	if x != nil {
		return x.Grade
	}
	return 0
}

func (x *DescribeProblemGradingOutput_Range) GetUpperBound() float32 {
	if x != nil {
		return x.UpperBound
	}
	return 0
}

var File_eolymp_atlas_scoring_service_proto protoreflect.FileDescriptor

const file_eolymp_atlas_scoring_service_proto_rawDesc = "" +
	"\n" +
	"\"eolymp/atlas/scoring_service.proto\x12\feolymp.atlas\x1a\x1deolymp/annotations/http.proto\x1a\x1eeolymp/annotations/scope.proto\x1a eolymp/atlas/scoring_score.proto\"\xa7\x01\n" +
	"\x11ScoreChangedEvent\x12\x1d\n" +
	"\n" +
	"problem_id\x18\x03 \x01(\tR\tproblemId\x12\x1b\n" +
	"\tmember_id\x18\x04 \x01(\tR\bmemberId\x12+\n" +
	"\x06before\x18\x01 \x01(\v2\x13.eolymp.atlas.ScoreR\x06before\x12)\n" +
	"\x05after\x18\x02 \x01(\v2\x13.eolymp.atlas.ScoreR\x05after\"1\n" +
	"\x12DescribeScoreInput\x12\x1b\n" +
	"\tmember_id\x18\x02 \x01(\tR\bmemberId\"@\n" +
	"\x13DescribeScoreOutput\x12)\n" +
	"\x05score\x18\x01 \x01(\v2\x13.eolymp.atlas.ScoreR\x05score\"\x1d\n" +
	"\x1bDescribeProblemGradingInput\"\xa8\x01\n" +
	"\x1cDescribeProblemGradingOutput\x12H\n" +
	"\x06ranges\x18\x02 \x03(\v20.eolymp.atlas.DescribeProblemGradingOutput.RangeR\x06ranges\x1a>\n" +
	"\x05Range\x12\x14\n" +
	"\x05grade\x18\x01 \x01(\rR\x05grade\x12\x1f\n" +
	"\vupper_bound\x18\x02 \x01(\x02R\n" +
	"upperBound2\xbf\x02\n" +
	"\x0eScoringService\x12\x8e\x01\n" +
	"\rDescribeScore\x12 .eolymp.atlas.DescribeScoreInput\x1a!.eolymp.atlas.DescribeScoreOutput\"8\x82\xe3\n" +
	"\x19\x8a\xe3\n" +
	"\x15atlas:submission:read\x82\xd3\xe4\x93\x02\x15\x12\x13/scores/{member_id}\x12\x9b\x01\n" +
	"\x16DescribeProblemGrading\x12).eolymp.atlas.DescribeProblemGradingInput\x1a*.eolymp.atlas.DescribeProblemGradingOutput\"*\x82\xe3\n" +
	"\x16\x8a\xe3\n" +
	"\x12atlas:problem:read\x82\xd3\xe4\x93\x02\n" +
	"\x12\b/gradingB-Z+github.com/eolymp/go-sdk/eolymp/atlas;atlasb\x06proto3"

var (
	file_eolymp_atlas_scoring_service_proto_rawDescOnce sync.Once
	file_eolymp_atlas_scoring_service_proto_rawDescData []byte
)

func file_eolymp_atlas_scoring_service_proto_rawDescGZIP() []byte {
	file_eolymp_atlas_scoring_service_proto_rawDescOnce.Do(func() {
		file_eolymp_atlas_scoring_service_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_eolymp_atlas_scoring_service_proto_rawDesc), len(file_eolymp_atlas_scoring_service_proto_rawDesc)))
	})
	return file_eolymp_atlas_scoring_service_proto_rawDescData
}

var file_eolymp_atlas_scoring_service_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_eolymp_atlas_scoring_service_proto_goTypes = []any{
	(*ScoreChangedEvent)(nil),                  // 0: eolymp.atlas.ScoreChangedEvent
	(*DescribeScoreInput)(nil),                 // 1: eolymp.atlas.DescribeScoreInput
	(*DescribeScoreOutput)(nil),                // 2: eolymp.atlas.DescribeScoreOutput
	(*DescribeProblemGradingInput)(nil),        // 3: eolymp.atlas.DescribeProblemGradingInput
	(*DescribeProblemGradingOutput)(nil),       // 4: eolymp.atlas.DescribeProblemGradingOutput
	(*DescribeProblemGradingOutput_Range)(nil), // 5: eolymp.atlas.DescribeProblemGradingOutput.Range
	(*Score)(nil),                              // 6: eolymp.atlas.Score
}
var file_eolymp_atlas_scoring_service_proto_depIdxs = []int32{
	6, // 0: eolymp.atlas.ScoreChangedEvent.before:type_name -> eolymp.atlas.Score
	6, // 1: eolymp.atlas.ScoreChangedEvent.after:type_name -> eolymp.atlas.Score
	6, // 2: eolymp.atlas.DescribeScoreOutput.score:type_name -> eolymp.atlas.Score
	5, // 3: eolymp.atlas.DescribeProblemGradingOutput.ranges:type_name -> eolymp.atlas.DescribeProblemGradingOutput.Range
	1, // 4: eolymp.atlas.ScoringService.DescribeScore:input_type -> eolymp.atlas.DescribeScoreInput
	3, // 5: eolymp.atlas.ScoringService.DescribeProblemGrading:input_type -> eolymp.atlas.DescribeProblemGradingInput
	2, // 6: eolymp.atlas.ScoringService.DescribeScore:output_type -> eolymp.atlas.DescribeScoreOutput
	4, // 7: eolymp.atlas.ScoringService.DescribeProblemGrading:output_type -> eolymp.atlas.DescribeProblemGradingOutput
	6, // [6:8] is the sub-list for method output_type
	4, // [4:6] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_eolymp_atlas_scoring_service_proto_init() }
func file_eolymp_atlas_scoring_service_proto_init() {
	if File_eolymp_atlas_scoring_service_proto != nil {
		return
	}
	file_eolymp_atlas_scoring_score_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_eolymp_atlas_scoring_service_proto_rawDesc), len(file_eolymp_atlas_scoring_service_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_eolymp_atlas_scoring_service_proto_goTypes,
		DependencyIndexes: file_eolymp_atlas_scoring_service_proto_depIdxs,
		MessageInfos:      file_eolymp_atlas_scoring_service_proto_msgTypes,
	}.Build()
	File_eolymp_atlas_scoring_service_proto = out.File
	file_eolymp_atlas_scoring_service_proto_goTypes = nil
	file_eolymp_atlas_scoring_service_proto_depIdxs = nil
}
