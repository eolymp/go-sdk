// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v3.21.12
// source: eolymp/atlas/problem.proto

package atlas

import (
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

type Problem_Extra int32

const (
	Problem_UNKNOWN_EXTRA  Problem_Extra = 0
	Problem_VOTE           Problem_Extra = 1
	Problem_TITLE          Problem_Extra = 2
	Problem_CONTENT_VALUE  Problem_Extra = 3
	Problem_CONTENT_RENDER Problem_Extra = 4
)

// Enum value maps for Problem_Extra.
var (
	Problem_Extra_name = map[int32]string{
		0: "UNKNOWN_EXTRA",
		1: "VOTE",
		2: "TITLE",
		3: "CONTENT_VALUE",
		4: "CONTENT_RENDER",
	}
	Problem_Extra_value = map[string]int32{
		"UNKNOWN_EXTRA":  0,
		"VOTE":           1,
		"TITLE":          2,
		"CONTENT_VALUE":  3,
		"CONTENT_RENDER": 4,
	}
)

func (x Problem_Extra) Enum() *Problem_Extra {
	p := new(Problem_Extra)
	*p = x
	return p
}

func (x Problem_Extra) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Problem_Extra) Descriptor() protoreflect.EnumDescriptor {
	return file_eolymp_atlas_problem_proto_enumTypes[0].Descriptor()
}

func (Problem_Extra) Type() protoreflect.EnumType {
	return &file_eolymp_atlas_problem_proto_enumTypes[0]
}

func (x Problem_Extra) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Problem_Extra.Descriptor instead.
func (Problem_Extra) EnumDescriptor() ([]byte, []int) {
	return file_eolymp_atlas_problem_proto_rawDescGZIP(), []int{0, 0}
}

type Problem_Type int32

const (
	Problem_UNKNOWN_TYPE Problem_Type = 0
	Problem_PROGRAM      Problem_Type = 1 // program problem, user should write a program to solve it
	Problem_FUNCTION     Problem_Type = 2 // function problem, user should write a function to solve it
	Problem_OUTPUT       Problem_Type = 3 // output-only problem, user should upload answer file to solve it
	Problem_SQL          Problem_Type = 4 // SQL problem, user should write SQL queries to solve it
	Problem_ML           Problem_Type = 5 // Machine learning problem
)

// Enum value maps for Problem_Type.
var (
	Problem_Type_name = map[int32]string{
		0: "UNKNOWN_TYPE",
		1: "PROGRAM",
		2: "FUNCTION",
		3: "OUTPUT",
		4: "SQL",
		5: "ML",
	}
	Problem_Type_value = map[string]int32{
		"UNKNOWN_TYPE": 0,
		"PROGRAM":      1,
		"FUNCTION":     2,
		"OUTPUT":       3,
		"SQL":          4,
		"ML":           5,
	}
)

func (x Problem_Type) Enum() *Problem_Type {
	p := new(Problem_Type)
	*p = x
	return p
}

func (x Problem_Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Problem_Type) Descriptor() protoreflect.EnumDescriptor {
	return file_eolymp_atlas_problem_proto_enumTypes[1].Descriptor()
}

func (Problem_Type) Type() protoreflect.EnumType {
	return &file_eolymp_atlas_problem_proto_enumTypes[1]
}

func (x Problem_Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Problem_Type.Descriptor instead.
func (Problem_Type) EnumDescriptor() ([]byte, []int) {
	return file_eolymp_atlas_problem_proto_rawDescGZIP(), []int{0, 1}
}

type Problem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Unique identifier.
	Id    string            `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Url   string            `protobuf:"bytes,682,opt,name=url,proto3" json:"url,omitempty"`
	Type  Problem_Type      `protobuf:"varint,23,opt,name=type,proto3,enum=eolymp.atlas.Problem_Type" json:"type,omitempty"`
	Links map[string]string `protobuf:"bytes,683,rep,name=links,proto3" json:"links,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Index in the public problem catalog.
	Number int32 `protobuf:"varint,10,opt,name=number,proto3" json:"number,omitempty"`
	// Problem is visible to users in public catalog.
	Visible bool `protobuf:"varint,11,opt,name=visible,proto3" json:"visible,omitempty"`
	// For imported problems, provides the source from where the problem is imported and synchronized.
	Origin  string       `protobuf:"bytes,13,opt,name=origin,proto3" json:"origin,omitempty"`
	Title   string       `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Content *ecm.Content `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
	// Problem topics (ID of topics from eolymp.taxonomy.TopicService)
	Topics      []string             `protobuf:"bytes,20,rep,name=topics,proto3" json:"topics,omitempty"`
	Score       float32              `protobuf:"fixed32,31,opt,name=score,proto3" json:"score,omitempty"`           // Total score
	Constraints *Problem_Constraints `protobuf:"bytes,30,opt,name=constraints,proto3" json:"constraints,omitempty"` // Constraints
	// Acceptance rate from 0 to 1, where 1 means that all submissions are accepted.
	AcceptanceRate      float32 `protobuf:"fixed32,40,opt,name=acceptance_rate,json=acceptanceRate,proto3" json:"acceptance_rate,omitempty"`
	SubmissionsCount    uint32  `protobuf:"varint,42,opt,name=submissions_count,json=submissionsCount,proto3" json:"submissions_count,omitempty"`
	SubmissionsAccepted uint32  `protobuf:"varint,43,opt,name=submissions_accepted,json=submissionsAccepted,proto3" json:"submissions_accepted,omitempty"`
	// Number of votes for the problem
	Vote      int32 `protobuf:"varint,50,opt,name=vote,proto3" json:"vote,omitempty"`
	VoteCount int32 `protobuf:"varint,51,opt,name=vote_count,json=voteCount,proto3" json:"vote_count,omitempty"`
	// Difficulty from 0 (very easy) to 5 (very hard)
	Difficulty uint32 `protobuf:"varint,21,opt,name=difficulty,proto3" json:"difficulty,omitempty"`
}

func (x *Problem) Reset() {
	*x = Problem{}
	mi := &file_eolymp_atlas_problem_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Problem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Problem) ProtoMessage() {}

func (x *Problem) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_atlas_problem_proto_msgTypes[0]
	if x != nil {
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
	return file_eolymp_atlas_problem_proto_rawDescGZIP(), []int{0}
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

func (x *Problem) GetType() Problem_Type {
	if x != nil {
		return x.Type
	}
	return Problem_UNKNOWN_TYPE
}

func (x *Problem) GetLinks() map[string]string {
	if x != nil {
		return x.Links
	}
	return nil
}

func (x *Problem) GetNumber() int32 {
	if x != nil {
		return x.Number
	}
	return 0
}

func (x *Problem) GetVisible() bool {
	if x != nil {
		return x.Visible
	}
	return false
}

func (x *Problem) GetOrigin() string {
	if x != nil {
		return x.Origin
	}
	return ""
}

func (x *Problem) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Problem) GetContent() *ecm.Content {
	if x != nil {
		return x.Content
	}
	return nil
}

func (x *Problem) GetTopics() []string {
	if x != nil {
		return x.Topics
	}
	return nil
}

func (x *Problem) GetScore() float32 {
	if x != nil {
		return x.Score
	}
	return 0
}

func (x *Problem) GetConstraints() *Problem_Constraints {
	if x != nil {
		return x.Constraints
	}
	return nil
}

func (x *Problem) GetAcceptanceRate() float32 {
	if x != nil {
		return x.AcceptanceRate
	}
	return 0
}

func (x *Problem) GetSubmissionsCount() uint32 {
	if x != nil {
		return x.SubmissionsCount
	}
	return 0
}

func (x *Problem) GetSubmissionsAccepted() uint32 {
	if x != nil {
		return x.SubmissionsAccepted
	}
	return 0
}

func (x *Problem) GetVote() int32 {
	if x != nil {
		return x.Vote
	}
	return 0
}

func (x *Problem) GetVoteCount() int32 {
	if x != nil {
		return x.VoteCount
	}
	return 0
}

func (x *Problem) GetDifficulty() uint32 {
	if x != nil {
		return x.Difficulty
	}
	return 0
}

type Problem_Constraints struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TimeLimitMin   uint32 `protobuf:"varint,60,opt,name=time_limit_min,json=timeLimitMin,proto3" json:"time_limit_min,omitempty"`
	TimeLimitMax   uint32 `protobuf:"varint,61,opt,name=time_limit_max,json=timeLimitMax,proto3" json:"time_limit_max,omitempty"`
	CpuLimitMin    uint32 `protobuf:"varint,62,opt,name=cpu_limit_min,json=cpuLimitMin,proto3" json:"cpu_limit_min,omitempty"`
	CpuLimitMax    uint32 `protobuf:"varint,63,opt,name=cpu_limit_max,json=cpuLimitMax,proto3" json:"cpu_limit_max,omitempty"`
	MemoryLimitMin uint64 `protobuf:"varint,64,opt,name=memory_limit_min,json=memoryLimitMin,proto3" json:"memory_limit_min,omitempty"`
	MemoryLimitMax uint64 `protobuf:"varint,65,opt,name=memory_limit_max,json=memoryLimitMax,proto3" json:"memory_limit_max,omitempty"`
}

func (x *Problem_Constraints) Reset() {
	*x = Problem_Constraints{}
	mi := &file_eolymp_atlas_problem_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Problem_Constraints) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Problem_Constraints) ProtoMessage() {}

func (x *Problem_Constraints) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_atlas_problem_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Problem_Constraints.ProtoReflect.Descriptor instead.
func (*Problem_Constraints) Descriptor() ([]byte, []int) {
	return file_eolymp_atlas_problem_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Problem_Constraints) GetTimeLimitMin() uint32 {
	if x != nil {
		return x.TimeLimitMin
	}
	return 0
}

func (x *Problem_Constraints) GetTimeLimitMax() uint32 {
	if x != nil {
		return x.TimeLimitMax
	}
	return 0
}

func (x *Problem_Constraints) GetCpuLimitMin() uint32 {
	if x != nil {
		return x.CpuLimitMin
	}
	return 0
}

func (x *Problem_Constraints) GetCpuLimitMax() uint32 {
	if x != nil {
		return x.CpuLimitMax
	}
	return 0
}

func (x *Problem_Constraints) GetMemoryLimitMin() uint64 {
	if x != nil {
		return x.MemoryLimitMin
	}
	return 0
}

func (x *Problem_Constraints) GetMemoryLimitMax() uint64 {
	if x != nil {
		return x.MemoryLimitMax
	}
	return 0
}

var File_eolymp_atlas_problem_proto protoreflect.FileDescriptor

var file_eolymp_atlas_problem_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x61, 0x74, 0x6c, 0x61, 0x73, 0x2f, 0x70,
	0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x65, 0x6f,
	0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x61, 0x74, 0x6c, 0x61, 0x73, 0x1a, 0x18, 0x65, 0x6f, 0x6c, 0x79,
	0x6d, 0x70, 0x2f, 0x65, 0x63, 0x6d, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xcf, 0x08, 0x0a, 0x07, 0x50, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x11, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0xaa, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x75, 0x72, 0x6c, 0x12, 0x2e, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x17, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x1a, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x61, 0x74, 0x6c, 0x61, 0x73,
	0x2e, 0x50, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x12, 0x37, 0x0a, 0x05, 0x6c, 0x69, 0x6e, 0x6b, 0x73, 0x18, 0xab, 0x05, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x61, 0x74, 0x6c,
	0x61, 0x73, 0x2e, 0x50, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x2e, 0x4c, 0x69, 0x6e, 0x6b, 0x73,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x05, 0x6c, 0x69, 0x6e, 0x6b, 0x73, 0x12, 0x16, 0x0a, 0x06,
	0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x69, 0x73, 0x69, 0x62, 0x6c, 0x65, 0x18,
	0x0b, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x76, 0x69, 0x73, 0x69, 0x62, 0x6c, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x2d, 0x0a, 0x07,
	0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e,
	0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x65, 0x63, 0x6d, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x74,
	0x6f, 0x70, 0x69, 0x63, 0x73, 0x18, 0x14, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x74, 0x6f, 0x70,
	0x69, 0x63, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x1f, 0x20, 0x01,
	0x28, 0x02, 0x52, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x12, 0x43, 0x0a, 0x0b, 0x63, 0x6f, 0x6e,
	0x73, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x74, 0x73, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21,
	0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x61, 0x74, 0x6c, 0x61, 0x73, 0x2e, 0x50, 0x72,
	0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x2e, 0x43, 0x6f, 0x6e, 0x73, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x74,
	0x73, 0x52, 0x0b, 0x63, 0x6f, 0x6e, 0x73, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x74, 0x73, 0x12, 0x27,
	0x0a, 0x0f, 0x61, 0x63, 0x63, 0x65, 0x70, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x72, 0x61, 0x74,
	0x65, 0x18, 0x28, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0e, 0x61, 0x63, 0x63, 0x65, 0x70, 0x74, 0x61,
	0x6e, 0x63, 0x65, 0x52, 0x61, 0x74, 0x65, 0x12, 0x2b, 0x0a, 0x11, 0x73, 0x75, 0x62, 0x6d, 0x69,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x2a, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x10, 0x73, 0x75, 0x62, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x43,
	0x6f, 0x75, 0x6e, 0x74, 0x12, 0x31, 0x0a, 0x14, 0x73, 0x75, 0x62, 0x6d, 0x69, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x73, 0x5f, 0x61, 0x63, 0x63, 0x65, 0x70, 0x74, 0x65, 0x64, 0x18, 0x2b, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x13, 0x73, 0x75, 0x62, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x41,
	0x63, 0x63, 0x65, 0x70, 0x74, 0x65, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x76, 0x6f, 0x74, 0x65, 0x18,
	0x32, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x76, 0x6f, 0x74, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x76,
	0x6f, 0x74, 0x65, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x33, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x09, 0x76, 0x6f, 0x74, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x64, 0x69,
	0x66, 0x66, 0x69, 0x63, 0x75, 0x6c, 0x74, 0x79, 0x18, 0x15, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a,
	0x64, 0x69, 0x66, 0x66, 0x69, 0x63, 0x75, 0x6c, 0x74, 0x79, 0x1a, 0xf5, 0x01, 0x0a, 0x0b, 0x43,
	0x6f, 0x6e, 0x73, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x74, 0x73, 0x12, 0x24, 0x0a, 0x0e, 0x74, 0x69,
	0x6d, 0x65, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x5f, 0x6d, 0x69, 0x6e, 0x18, 0x3c, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x0c, 0x74, 0x69, 0x6d, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x4d, 0x69, 0x6e,
	0x12, 0x24, 0x0a, 0x0e, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x5f, 0x6d,
	0x61, 0x78, 0x18, 0x3d, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c, 0x74, 0x69, 0x6d, 0x65, 0x4c, 0x69,
	0x6d, 0x69, 0x74, 0x4d, 0x61, 0x78, 0x12, 0x22, 0x0a, 0x0d, 0x63, 0x70, 0x75, 0x5f, 0x6c, 0x69,
	0x6d, 0x69, 0x74, 0x5f, 0x6d, 0x69, 0x6e, 0x18, 0x3e, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x63,
	0x70, 0x75, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x4d, 0x69, 0x6e, 0x12, 0x22, 0x0a, 0x0d, 0x63, 0x70,
	0x75, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x5f, 0x6d, 0x61, 0x78, 0x18, 0x3f, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x0b, 0x63, 0x70, 0x75, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x4d, 0x61, 0x78, 0x12, 0x28,
	0x0a, 0x10, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x5f, 0x6d,
	0x69, 0x6e, 0x18, 0x40, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0e, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79,
	0x4c, 0x69, 0x6d, 0x69, 0x74, 0x4d, 0x69, 0x6e, 0x12, 0x28, 0x0a, 0x10, 0x6d, 0x65, 0x6d, 0x6f,
	0x72, 0x79, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x5f, 0x6d, 0x61, 0x78, 0x18, 0x41, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x0e, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x4d,
	0x61, 0x78, 0x1a, 0x38, 0x0a, 0x0a, 0x4c, 0x69, 0x6e, 0x6b, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x56, 0x0a, 0x05,
	0x45, 0x78, 0x74, 0x72, 0x61, 0x12, 0x11, 0x0a, 0x0d, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e,
	0x5f, 0x45, 0x58, 0x54, 0x52, 0x41, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x56, 0x4f, 0x54, 0x45,
	0x10, 0x01, 0x12, 0x09, 0x0a, 0x05, 0x54, 0x49, 0x54, 0x4c, 0x45, 0x10, 0x02, 0x12, 0x11, 0x0a,
	0x0d, 0x43, 0x4f, 0x4e, 0x54, 0x45, 0x4e, 0x54, 0x5f, 0x56, 0x41, 0x4c, 0x55, 0x45, 0x10, 0x03,
	0x12, 0x12, 0x0a, 0x0e, 0x43, 0x4f, 0x4e, 0x54, 0x45, 0x4e, 0x54, 0x5f, 0x52, 0x45, 0x4e, 0x44,
	0x45, 0x52, 0x10, 0x04, 0x22, 0x50, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x10, 0x0a, 0x0c,
	0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x10, 0x00, 0x12, 0x0b,
	0x0a, 0x07, 0x50, 0x52, 0x4f, 0x47, 0x52, 0x41, 0x4d, 0x10, 0x01, 0x12, 0x0c, 0x0a, 0x08, 0x46,
	0x55, 0x4e, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x02, 0x12, 0x0a, 0x0a, 0x06, 0x4f, 0x55, 0x54,
	0x50, 0x55, 0x54, 0x10, 0x03, 0x12, 0x07, 0x0a, 0x03, 0x53, 0x51, 0x4c, 0x10, 0x04, 0x12, 0x06,
	0x0a, 0x02, 0x4d, 0x4c, 0x10, 0x05, 0x42, 0x2d, 0x5a, 0x2b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x67, 0x6f, 0x2d, 0x73,
	0x64, 0x6b, 0x2f, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x61, 0x74, 0x6c, 0x61, 0x73, 0x3b,
	0x61, 0x74, 0x6c, 0x61, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_eolymp_atlas_problem_proto_rawDescOnce sync.Once
	file_eolymp_atlas_problem_proto_rawDescData = file_eolymp_atlas_problem_proto_rawDesc
)

func file_eolymp_atlas_problem_proto_rawDescGZIP() []byte {
	file_eolymp_atlas_problem_proto_rawDescOnce.Do(func() {
		file_eolymp_atlas_problem_proto_rawDescData = protoimpl.X.CompressGZIP(file_eolymp_atlas_problem_proto_rawDescData)
	})
	return file_eolymp_atlas_problem_proto_rawDescData
}

var file_eolymp_atlas_problem_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_eolymp_atlas_problem_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_eolymp_atlas_problem_proto_goTypes = []any{
	(Problem_Extra)(0),          // 0: eolymp.atlas.Problem.Extra
	(Problem_Type)(0),           // 1: eolymp.atlas.Problem.Type
	(*Problem)(nil),             // 2: eolymp.atlas.Problem
	(*Problem_Constraints)(nil), // 3: eolymp.atlas.Problem.Constraints
	nil,                         // 4: eolymp.atlas.Problem.LinksEntry
	(*ecm.Content)(nil),         // 5: eolymp.ecm.Content
}
var file_eolymp_atlas_problem_proto_depIdxs = []int32{
	1, // 0: eolymp.atlas.Problem.type:type_name -> eolymp.atlas.Problem.Type
	4, // 1: eolymp.atlas.Problem.links:type_name -> eolymp.atlas.Problem.LinksEntry
	5, // 2: eolymp.atlas.Problem.content:type_name -> eolymp.ecm.Content
	3, // 3: eolymp.atlas.Problem.constraints:type_name -> eolymp.atlas.Problem.Constraints
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_eolymp_atlas_problem_proto_init() }
func file_eolymp_atlas_problem_proto_init() {
	if File_eolymp_atlas_problem_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_eolymp_atlas_problem_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_eolymp_atlas_problem_proto_goTypes,
		DependencyIndexes: file_eolymp_atlas_problem_proto_depIdxs,
		EnumInfos:         file_eolymp_atlas_problem_proto_enumTypes,
		MessageInfos:      file_eolymp_atlas_problem_proto_msgTypes,
	}.Build()
	File_eolymp_atlas_problem_proto = out.File
	file_eolymp_atlas_problem_proto_rawDesc = nil
	file_eolymp_atlas_problem_proto_goTypes = nil
	file_eolymp_atlas_problem_proto_depIdxs = nil
}
