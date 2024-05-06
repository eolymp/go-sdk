// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.0
// 	protoc        v3.21.12
// source: eolymp/judge/problem.proto

package judge

import (
	atlas "github.com/eolymp/go-sdk/eolymp/atlas"
	ecm "github.com/eolymp/go-sdk/eolymp/ecm"
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

type Problem_Statement_Format int32

const (
	Problem_Statement_NONE     Problem_Statement_Format = 0
	Problem_Statement_TEX      Problem_Statement_Format = 1
	Problem_Statement_MARKDOWN Problem_Statement_Format = 2
	Problem_Statement_HTML     Problem_Statement_Format = 3
	Problem_Statement_RICH     Problem_Statement_Format = 4
)

// Enum value maps for Problem_Statement_Format.
var (
	Problem_Statement_Format_name = map[int32]string{
		0: "NONE",
		1: "TEX",
		2: "MARKDOWN",
		3: "HTML",
		4: "RICH",
	}
	Problem_Statement_Format_value = map[string]int32{
		"NONE":     0,
		"TEX":      1,
		"MARKDOWN": 2,
		"HTML":     3,
		"RICH":     4,
	}
)

func (x Problem_Statement_Format) Enum() *Problem_Statement_Format {
	p := new(Problem_Statement_Format)
	*p = x
	return p
}

func (x Problem_Statement_Format) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Problem_Statement_Format) Descriptor() protoreflect.EnumDescriptor {
	return file_eolymp_judge_problem_proto_enumTypes[0].Descriptor()
}

func (Problem_Statement_Format) Type() protoreflect.EnumType {
	return &file_eolymp_judge_problem_proto_enumTypes[0]
}

func (x Problem_Statement_Format) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Problem_Statement_Format.Descriptor instead.
func (Problem_Statement_Format) EnumDescriptor() ([]byte, []int) {
	return file_eolymp_judge_problem_proto_rawDescGZIP(), []int{0, 0, 0}
}

type Problem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                 string               `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"` // globally problem ID
	Url                string               `protobuf:"bytes,6,opt,name=url,proto3" json:"url,omitempty"`
	Index              uint32               `protobuf:"varint,2,opt,name=index,proto3" json:"index,omitempty"`                                                                           // problem index within contest
	Score              float32              `protobuf:"fixed32,3,opt,name=score,proto3" json:"score,omitempty"`                                                                          // score for solving problem
	BaseId             string               `protobuf:"bytes,4,opt,name=base_id,json=baseId,proto3" json:"base_id,omitempty"`                                                            // ID of the problem in database achieve
	ContestId          string               `protobuf:"bytes,5,opt,name=contest_id,json=contestId,proto3" json:"contest_id,omitempty"`                                                   // contest
	FeedbackPolicy     atlas.FeedbackPolicy `protobuf:"varint,10,opt,name=feedback_policy,json=feedbackPolicy,proto3,enum=eolymp.atlas.FeedbackPolicy" json:"feedback_policy,omitempty"` // defines feedback policy for the problem
	TimeLimit          uint32               `protobuf:"varint,200,opt,name=time_limit,json=timeLimit,proto3" json:"time_limit,omitempty"`                                                // in milliseconds
	CpuLimit           uint32               `protobuf:"varint,204,opt,name=cpu_limit,json=cpuLimit,proto3" json:"cpu_limit,omitempty"`                                                   // in milliseconds
	MemoryLimit        uint64               `protobuf:"varint,201,opt,name=memory_limit,json=memoryLimit,proto3" json:"memory_limit,omitempty"`                                          // in bytes
	FileSizeLimit      uint64               `protobuf:"varint,202,opt,name=file_size_limit,json=fileSizeLimit,proto3" json:"file_size_limit,omitempty"`                                  // in bytes
	SubmitLimit        uint32               `protobuf:"varint,203,opt,name=submit_limit,json=submitLimit,proto3" json:"submit_limit,omitempty"`                                          // in number of submits
	ScoreByBestTestset bool                 `protobuf:"varint,210,opt,name=score_by_best_testset,json=scoreByBestTestset,proto3" json:"score_by_best_testset,omitempty"`                 // problem score is calculated as sum of best score in each testset (best among all submissions)
}

func (x *Problem) Reset() {
	*x = Problem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_judge_problem_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Problem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Problem) ProtoMessage() {}

func (x *Problem) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_judge_problem_proto_msgTypes[0]
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
	return file_eolymp_judge_problem_proto_rawDescGZIP(), []int{0}
}

func (x *Problem) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Problem) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *Problem) GetIndex() uint32 {
	if x != nil {
		return x.Index
	}
	return 0
}

func (x *Problem) GetScore() float32 {
	if x != nil {
		return x.Score
	}
	return 0
}

func (x *Problem) GetBaseId() string {
	if x != nil {
		return x.BaseId
	}
	return ""
}

func (x *Problem) GetContestId() string {
	if x != nil {
		return x.ContestId
	}
	return ""
}

func (x *Problem) GetFeedbackPolicy() atlas.FeedbackPolicy {
	if x != nil {
		return x.FeedbackPolicy
	}
	return atlas.FeedbackPolicy(0)
}

func (x *Problem) GetTimeLimit() uint32 {
	if x != nil {
		return x.TimeLimit
	}
	return 0
}

func (x *Problem) GetCpuLimit() uint32 {
	if x != nil {
		return x.CpuLimit
	}
	return 0
}

func (x *Problem) GetMemoryLimit() uint64 {
	if x != nil {
		return x.MemoryLimit
	}
	return 0
}

func (x *Problem) GetFileSizeLimit() uint64 {
	if x != nil {
		return x.FileSizeLimit
	}
	return 0
}

func (x *Problem) GetSubmitLimit() uint32 {
	if x != nil {
		return x.SubmitLimit
	}
	return 0
}

func (x *Problem) GetScoreByBestTestset() bool {
	if x != nil {
		return x.ScoreByBestTestset
	}
	return false
}

// Statement is localized problem statement
type Problem_Statement struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Locale       string                   `protobuf:"bytes,1,opt,name=locale,proto3" json:"locale,omitempty"`
	Title        string                   `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	ContentRaw   string                   `protobuf:"bytes,3,opt,name=content_raw,json=contentRaw,proto3" json:"content_raw,omitempty"` // deprecated: use content instead
	Content      *ecm.Node                `protobuf:"bytes,4,opt,name=content,proto3" json:"content,omitempty"`
	Format       Problem_Statement_Format `protobuf:"varint,5,opt,name=format,proto3,enum=eolymp.judge.Problem_Statement_Format" json:"format,omitempty"` // deprecated: content always uses ECM format
	DownloadLink string                   `protobuf:"bytes,6,opt,name=download_link,json=downloadLink,proto3" json:"download_link,omitempty"`
}

func (x *Problem_Statement) Reset() {
	*x = Problem_Statement{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_judge_problem_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Problem_Statement) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Problem_Statement) ProtoMessage() {}

func (x *Problem_Statement) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_judge_problem_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Problem_Statement.ProtoReflect.Descriptor instead.
func (*Problem_Statement) Descriptor() ([]byte, []int) {
	return file_eolymp_judge_problem_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Problem_Statement) GetLocale() string {
	if x != nil {
		return x.Locale
	}
	return ""
}

func (x *Problem_Statement) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Problem_Statement) GetContentRaw() string {
	if x != nil {
		return x.ContentRaw
	}
	return ""
}

func (x *Problem_Statement) GetContent() *ecm.Node {
	if x != nil {
		return x.Content
	}
	return nil
}

func (x *Problem_Statement) GetFormat() Problem_Statement_Format {
	if x != nil {
		return x.Format
	}
	return Problem_Statement_NONE
}

func (x *Problem_Statement) GetDownloadLink() string {
	if x != nil {
		return x.DownloadLink
	}
	return ""
}

// Test defines set of input and answer data
type Problem_Test struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Index          uint32  `protobuf:"varint,1,opt,name=index,proto3" json:"index,omitempty"`
	Example        bool    `protobuf:"varint,2,opt,name=example,proto3" json:"example,omitempty"`
	InputObjectId  string  `protobuf:"bytes,3,opt,name=input_object_id,json=inputObjectId,proto3" json:"input_object_id,omitempty"`
	AnswerObjectId string  `protobuf:"bytes,4,opt,name=answer_object_id,json=answerObjectId,proto3" json:"answer_object_id,omitempty"`
	Score          float32 `protobuf:"fixed32,5,opt,name=score,proto3" json:"score,omitempty"`
}

func (x *Problem_Test) Reset() {
	*x = Problem_Test{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_judge_problem_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Problem_Test) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Problem_Test) ProtoMessage() {}

func (x *Problem_Test) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_judge_problem_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Problem_Test.ProtoReflect.Descriptor instead.
func (*Problem_Test) Descriptor() ([]byte, []int) {
	return file_eolymp_judge_problem_proto_rawDescGZIP(), []int{0, 1}
}

func (x *Problem_Test) GetIndex() uint32 {
	if x != nil {
		return x.Index
	}
	return 0
}

func (x *Problem_Test) GetExample() bool {
	if x != nil {
		return x.Example
	}
	return false
}

func (x *Problem_Test) GetInputObjectId() string {
	if x != nil {
		return x.InputObjectId
	}
	return ""
}

func (x *Problem_Test) GetAnswerObjectId() string {
	if x != nil {
		return x.AnswerObjectId
	}
	return ""
}

func (x *Problem_Test) GetScore() float32 {
	if x != nil {
		return x.Score
	}
	return 0
}

// attachment to the problem statement
type Problem_Attachment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Link string `protobuf:"bytes,4,opt,name=link,proto3" json:"link,omitempty"`
}

func (x *Problem_Attachment) Reset() {
	*x = Problem_Attachment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_judge_problem_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Problem_Attachment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Problem_Attachment) ProtoMessage() {}

func (x *Problem_Attachment) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_judge_problem_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Problem_Attachment.ProtoReflect.Descriptor instead.
func (*Problem_Attachment) Descriptor() ([]byte, []int) {
	return file_eolymp_judge_problem_proto_rawDescGZIP(), []int{0, 2}
}

func (x *Problem_Attachment) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Problem_Attachment) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Problem_Attachment) GetLink() string {
	if x != nil {
		return x.Link
	}
	return ""
}

var File_eolymp_judge_problem_proto protoreflect.FileDescriptor

var file_eolymp_judge_problem_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x6a, 0x75, 0x64, 0x67, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x65, 0x6f,
	0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x6a, 0x75, 0x64, 0x67, 0x65, 0x1a, 0x23, 0x65, 0x6f, 0x6c, 0x79,
	0x6d, 0x70, 0x2f, 0x61, 0x74, 0x6c, 0x61, 0x73, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67,
	0x5f, 0x66, 0x65, 0x65, 0x64, 0x62, 0x61, 0x63, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x15, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x65, 0x63, 0x6d, 0x2f, 0x6e, 0x6f, 0x64, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xcd, 0x07, 0x0a, 0x07, 0x50, 0x72, 0x6f, 0x62, 0x6c,
	0x65, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x75, 0x72, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x63,
	0x6f, 0x72, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65,
	0x12, 0x17, 0x0a, 0x07, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x62, 0x61, 0x73, 0x65, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6f, 0x6e,
	0x74, 0x65, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x73, 0x74, 0x49, 0x64, 0x12, 0x45, 0x0a, 0x0f, 0x66, 0x65, 0x65, 0x64,
	0x62, 0x61, 0x63, 0x6b, 0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x1c, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x61, 0x74, 0x6c, 0x61, 0x73,
	0x2e, 0x46, 0x65, 0x65, 0x64, 0x62, 0x61, 0x63, 0x6b, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52,
	0x0e, 0x66, 0x65, 0x65, 0x64, 0x62, 0x61, 0x63, 0x6b, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12,
	0x1e, 0x0a, 0x0a, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0xc8, 0x01,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x12,
	0x1c, 0x0a, 0x09, 0x63, 0x70, 0x75, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0xcc, 0x01, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x08, 0x63, 0x70, 0x75, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x22, 0x0a,
	0x0c, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0xc9, 0x01,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x0b, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x4c, 0x69, 0x6d, 0x69,
	0x74, 0x12, 0x27, 0x0a, 0x0f, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x5f, 0x6c,
	0x69, 0x6d, 0x69, 0x74, 0x18, 0xca, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0d, 0x66, 0x69, 0x6c,
	0x65, 0x53, 0x69, 0x7a, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x22, 0x0a, 0x0c, 0x73, 0x75,
	0x62, 0x6d, 0x69, 0x74, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0xcb, 0x01, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x0b, 0x73, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x32,
	0x0a, 0x15, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x5f, 0x62, 0x79, 0x5f, 0x62, 0x65, 0x73, 0x74, 0x5f,
	0x74, 0x65, 0x73, 0x74, 0x73, 0x65, 0x74, 0x18, 0xd2, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x12,
	0x73, 0x63, 0x6f, 0x72, 0x65, 0x42, 0x79, 0x42, 0x65, 0x73, 0x74, 0x54, 0x65, 0x73, 0x74, 0x73,
	0x65, 0x74, 0x1a, 0xaa, 0x02, 0x0a, 0x09, 0x53, 0x74, 0x61, 0x74, 0x65, 0x6d, 0x65, 0x6e, 0x74,
	0x12, 0x16, 0x0a, 0x06, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x1f,
	0x0a, 0x0b, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x5f, 0x72, 0x61, 0x77, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x52, 0x61, 0x77, 0x12,
	0x2a, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x10, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x65, 0x63, 0x6d, 0x2e, 0x4e, 0x6f,
	0x64, 0x65, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x3e, 0x0a, 0x06, 0x66,
	0x6f, 0x72, 0x6d, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x26, 0x2e, 0x65, 0x6f,
	0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x6a, 0x75, 0x64, 0x67, 0x65, 0x2e, 0x50, 0x72, 0x6f, 0x62, 0x6c,
	0x65, 0x6d, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x46, 0x6f, 0x72,
	0x6d, 0x61, 0x74, 0x52, 0x06, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x64,
	0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x6c, 0x69, 0x6e, 0x6b, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0c, 0x64, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x4c, 0x69, 0x6e, 0x6b,
	0x22, 0x3d, 0x0a, 0x06, 0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x4f,
	0x4e, 0x45, 0x10, 0x00, 0x12, 0x07, 0x0a, 0x03, 0x54, 0x45, 0x58, 0x10, 0x01, 0x12, 0x0c, 0x0a,
	0x08, 0x4d, 0x41, 0x52, 0x4b, 0x44, 0x4f, 0x57, 0x4e, 0x10, 0x02, 0x12, 0x08, 0x0a, 0x04, 0x48,
	0x54, 0x4d, 0x4c, 0x10, 0x03, 0x12, 0x08, 0x0a, 0x04, 0x52, 0x49, 0x43, 0x48, 0x10, 0x04, 0x1a,
	0x9e, 0x01, 0x0a, 0x04, 0x54, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6e, 0x64, 0x65,
	0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x18,
	0x0a, 0x07, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x07, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x12, 0x26, 0x0a, 0x0f, 0x69, 0x6e, 0x70, 0x75,
	0x74, 0x5f, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0d, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64,
	0x12, 0x28, 0x0a, 0x10, 0x61, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x5f, 0x6f, 0x62, 0x6a, 0x65, 0x63,
	0x74, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x61, 0x6e, 0x73, 0x77,
	0x65, 0x72, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x63,
	0x6f, 0x72, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65,
	0x1a, 0x44, 0x0a, 0x0a, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x69, 0x6e, 0x6b, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6c, 0x69, 0x6e, 0x6b, 0x42, 0x2d, 0x5a, 0x2b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x67, 0x6f, 0x2d, 0x73,
	0x64, 0x6b, 0x2f, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x6a, 0x75, 0x64, 0x67, 0x65, 0x3b,
	0x6a, 0x75, 0x64, 0x67, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_eolymp_judge_problem_proto_rawDescOnce sync.Once
	file_eolymp_judge_problem_proto_rawDescData = file_eolymp_judge_problem_proto_rawDesc
)

func file_eolymp_judge_problem_proto_rawDescGZIP() []byte {
	file_eolymp_judge_problem_proto_rawDescOnce.Do(func() {
		file_eolymp_judge_problem_proto_rawDescData = protoimpl.X.CompressGZIP(file_eolymp_judge_problem_proto_rawDescData)
	})
	return file_eolymp_judge_problem_proto_rawDescData
}

var file_eolymp_judge_problem_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_eolymp_judge_problem_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_eolymp_judge_problem_proto_goTypes = []interface{}{
	(Problem_Statement_Format)(0), // 0: eolymp.judge.Problem.Statement.Format
	(*Problem)(nil),               // 1: eolymp.judge.Problem
	(*Problem_Statement)(nil),     // 2: eolymp.judge.Problem.Statement
	(*Problem_Test)(nil),          // 3: eolymp.judge.Problem.Test
	(*Problem_Attachment)(nil),    // 4: eolymp.judge.Problem.Attachment
	(atlas.FeedbackPolicy)(0),     // 5: eolymp.atlas.FeedbackPolicy
	(*ecm.Node)(nil),              // 6: eolymp.ecm.Node
}
var file_eolymp_judge_problem_proto_depIdxs = []int32{
	5, // 0: eolymp.judge.Problem.feedback_policy:type_name -> eolymp.atlas.FeedbackPolicy
	6, // 1: eolymp.judge.Problem.Statement.content:type_name -> eolymp.ecm.Node
	0, // 2: eolymp.judge.Problem.Statement.format:type_name -> eolymp.judge.Problem.Statement.Format
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_eolymp_judge_problem_proto_init() }
func file_eolymp_judge_problem_proto_init() {
	if File_eolymp_judge_problem_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_eolymp_judge_problem_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_eolymp_judge_problem_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Problem_Statement); i {
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
		file_eolymp_judge_problem_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Problem_Test); i {
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
		file_eolymp_judge_problem_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Problem_Attachment); i {
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
			RawDescriptor: file_eolymp_judge_problem_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_eolymp_judge_problem_proto_goTypes,
		DependencyIndexes: file_eolymp_judge_problem_proto_depIdxs,
		EnumInfos:         file_eolymp_judge_problem_proto_enumTypes,
		MessageInfos:      file_eolymp_judge_problem_proto_msgTypes,
	}.Build()
	File_eolymp_judge_problem_proto = out.File
	file_eolymp_judge_problem_proto_rawDesc = nil
	file_eolymp_judge_problem_proto_goTypes = nil
	file_eolymp_judge_problem_proto_depIdxs = nil
}
