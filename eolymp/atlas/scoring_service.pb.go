// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.1
// 	protoc        v4.24.4
// source: eolymp/atlas/scoring_service.proto

package atlas

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

var file_eolymp_atlas_scoring_service_proto_rawDesc = []byte{
	0x0a, 0x22, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x61, 0x74, 0x6c, 0x61, 0x73, 0x2f, 0x73,
	0x63, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x61, 0x74, 0x6c,
	0x61, 0x73, 0x1a, 0x1d, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x20, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x61, 0x74, 0x6c, 0x61, 0x73, 0x2f,
	0x73, 0x63, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xa7, 0x01, 0x0a, 0x11, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x43, 0x68, 0x61,
	0x6e, 0x67, 0x65, 0x64, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f,
	0x62, 0x6c, 0x65, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70,
	0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x65, 0x6d, 0x62,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6d, 0x65, 0x6d,
	0x62, 0x65, 0x72, 0x49, 0x64, 0x12, 0x2b, 0x0a, 0x06, 0x62, 0x65, 0x66, 0x6f, 0x72, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x61,
	0x74, 0x6c, 0x61, 0x73, 0x2e, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x52, 0x06, 0x62, 0x65, 0x66, 0x6f,
	0x72, 0x65, 0x12, 0x29, 0x0a, 0x05, 0x61, 0x66, 0x74, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x13, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x61, 0x74, 0x6c, 0x61, 0x73,
	0x2e, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x52, 0x05, 0x61, 0x66, 0x74, 0x65, 0x72, 0x22, 0x31, 0x0a,
	0x12, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x49, 0x6e,
	0x70, 0x75, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x49, 0x64,
	0x22, 0x40, 0x0a, 0x13, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x53, 0x63, 0x6f, 0x72,
	0x65, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x12, 0x29, 0x0a, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e,
	0x61, 0x74, 0x6c, 0x61, 0x73, 0x2e, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x52, 0x05, 0x73, 0x63, 0x6f,
	0x72, 0x65, 0x22, 0x1d, 0x0a, 0x1b, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x50, 0x72,
	0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x47, 0x72, 0x61, 0x64, 0x69, 0x6e, 0x67, 0x49, 0x6e, 0x70, 0x75,
	0x74, 0x22, 0xa8, 0x01, 0x0a, 0x1c, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x50, 0x72,
	0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x47, 0x72, 0x61, 0x64, 0x69, 0x6e, 0x67, 0x4f, 0x75, 0x74, 0x70,
	0x75, 0x74, 0x12, 0x48, 0x0a, 0x06, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x30, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x61, 0x74, 0x6c, 0x61,
	0x73, 0x2e, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x50, 0x72, 0x6f, 0x62, 0x6c, 0x65,
	0x6d, 0x47, 0x72, 0x61, 0x64, 0x69, 0x6e, 0x67, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x2e, 0x52,
	0x61, 0x6e, 0x67, 0x65, 0x52, 0x06, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x73, 0x1a, 0x3e, 0x0a, 0x05,
	0x52, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x67, 0x72, 0x61, 0x64, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x67, 0x72, 0x61, 0x64, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x75,
	0x70, 0x70, 0x65, 0x72, 0x5f, 0x62, 0x6f, 0x75, 0x6e, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02,
	0x52, 0x0a, 0x75, 0x70, 0x70, 0x65, 0x72, 0x42, 0x6f, 0x75, 0x6e, 0x64, 0x32, 0xbf, 0x02, 0x0a,
	0x0e, 0x53, 0x63, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x8e, 0x01, 0x0a, 0x0d, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x53, 0x63, 0x6f, 0x72,
	0x65, 0x12, 0x20, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x61, 0x74, 0x6c, 0x61, 0x73,
	0x2e, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x49, 0x6e,
	0x70, 0x75, 0x74, 0x1a, 0x21, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x61, 0x74, 0x6c,
	0x61, 0x73, 0x2e, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x53, 0x63, 0x6f, 0x72, 0x65,
	0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x22, 0x38, 0x82, 0xe3, 0x0a, 0x19, 0x8a, 0xe3, 0x0a, 0x15,
	0x61, 0x74, 0x6c, 0x61, 0x73, 0x3a, 0x73, 0x75, 0x62, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x3a, 0x72, 0x65, 0x61, 0x64, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x15, 0x12, 0x13, 0x2f, 0x73, 0x63,
	0x6f, 0x72, 0x65, 0x73, 0x2f, 0x7b, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x7d,
	0x12, 0x9b, 0x01, 0x0a, 0x16, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x50, 0x72, 0x6f,
	0x62, 0x6c, 0x65, 0x6d, 0x47, 0x72, 0x61, 0x64, 0x69, 0x6e, 0x67, 0x12, 0x29, 0x2e, 0x65, 0x6f,
	0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x61, 0x74, 0x6c, 0x61, 0x73, 0x2e, 0x44, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x62, 0x65, 0x50, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x47, 0x72, 0x61, 0x64, 0x69, 0x6e,
	0x67, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x2a, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e,
	0x61, 0x74, 0x6c, 0x61, 0x73, 0x2e, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x50, 0x72,
	0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x47, 0x72, 0x61, 0x64, 0x69, 0x6e, 0x67, 0x4f, 0x75, 0x74, 0x70,
	0x75, 0x74, 0x22, 0x2a, 0x82, 0xe3, 0x0a, 0x16, 0x8a, 0xe3, 0x0a, 0x12, 0x61, 0x74, 0x6c, 0x61,
	0x73, 0x3a, 0x70, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x3a, 0x72, 0x65, 0x61, 0x64, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x0a, 0x12, 0x08, 0x2f, 0x67, 0x72, 0x61, 0x64, 0x69, 0x6e, 0x67, 0x42, 0x2d,
	0x5a, 0x2b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6f, 0x6c,
	0x79, 0x6d, 0x70, 0x2f, 0x67, 0x6f, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x65, 0x6f, 0x6c, 0x79, 0x6d,
	0x70, 0x2f, 0x61, 0x74, 0x6c, 0x61, 0x73, 0x3b, 0x61, 0x74, 0x6c, 0x61, 0x73, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_eolymp_atlas_scoring_service_proto_rawDescOnce sync.Once
	file_eolymp_atlas_scoring_service_proto_rawDescData = file_eolymp_atlas_scoring_service_proto_rawDesc
)

func file_eolymp_atlas_scoring_service_proto_rawDescGZIP() []byte {
	file_eolymp_atlas_scoring_service_proto_rawDescOnce.Do(func() {
		file_eolymp_atlas_scoring_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_eolymp_atlas_scoring_service_proto_rawDescData)
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
			RawDescriptor: file_eolymp_atlas_scoring_service_proto_rawDesc,
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
	file_eolymp_atlas_scoring_service_proto_rawDesc = nil
	file_eolymp_atlas_scoring_service_proto_goTypes = nil
	file_eolymp_atlas_scoring_service_proto_depIdxs = nil
}
