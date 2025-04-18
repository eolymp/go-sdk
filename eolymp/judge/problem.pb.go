// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v4.24.4
// source: eolymp/judge/problem.proto

package judge

import (
	atlas "github.com/eolymp/go-sdk/eolymp/atlas"
	ecm "github.com/eolymp/go-sdk/eolymp/ecm"
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

type Problem struct {
	state              protoimpl.MessageState `protogen:"open.v1"`
	Id                 string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"` // globally problem ID
	Url                string                 `protobuf:"bytes,6,opt,name=url,proto3" json:"url,omitempty"`
	Type               atlas.Problem_Type     `protobuf:"varint,23,opt,name=type,proto3,enum=eolymp.atlas.Problem_Type" json:"type,omitempty"`
	Index              uint32                 `protobuf:"varint,2,opt,name=index,proto3" json:"index,omitempty"`                // problem index within contest
	Score              float32                `protobuf:"fixed32,3,opt,name=score,proto3" json:"score,omitempty"`               // score for solving problem
	BaseId             string                 `protobuf:"bytes,4,opt,name=base_id,json=baseId,proto3" json:"base_id,omitempty"` // ID of the problem in database achieve
	BaseNumber         uint32                 `protobuf:"varint,45,opt,name=base_number,json=baseNumber,proto3" json:"base_number,omitempty"`
	BaseSpaceId        string                 `protobuf:"bytes,40,opt,name=base_space_id,json=baseSpaceId,proto3" json:"base_space_id,omitempty"`
	ContestId          string                 `protobuf:"bytes,5,opt,name=contest_id,json=contestId,proto3" json:"contest_id,omitempty"`                                                   // contest
	FeedbackPolicy     atlas.FeedbackPolicy   `protobuf:"varint,10,opt,name=feedback_policy,json=feedbackPolicy,proto3,enum=eolymp.atlas.FeedbackPolicy" json:"feedback_policy,omitempty"` // defines feedback policy for the problem
	TimeLimit          uint32                 `protobuf:"varint,200,opt,name=time_limit,json=timeLimit,proto3" json:"time_limit,omitempty"`                                                // in milliseconds
	CpuLimit           uint32                 `protobuf:"varint,204,opt,name=cpu_limit,json=cpuLimit,proto3" json:"cpu_limit,omitempty"`                                                   // in milliseconds
	MemoryLimit        uint64                 `protobuf:"varint,201,opt,name=memory_limit,json=memoryLimit,proto3" json:"memory_limit,omitempty"`                                          // in bytes
	FileSizeLimit      uint64                 `protobuf:"varint,202,opt,name=file_size_limit,json=fileSizeLimit,proto3" json:"file_size_limit,omitempty"`                                  // in bytes
	SubmitLimit        uint32                 `protobuf:"varint,203,opt,name=submit_limit,json=submitLimit,proto3" json:"submit_limit,omitempty"`                                          // in number of submits
	ScoreByBestTestset bool                   `protobuf:"varint,210,opt,name=score_by_best_testset,json=scoreByBestTestset,proto3" json:"score_by_best_testset,omitempty"`                 // problem score is calculated as sum of best score in each testset (best among all submissions)
	unknownFields      protoimpl.UnknownFields
	sizeCache          protoimpl.SizeCache
}

func (x *Problem) Reset() {
	*x = Problem{}
	mi := &file_eolymp_judge_problem_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Problem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Problem) ProtoMessage() {}

func (x *Problem) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_judge_problem_proto_msgTypes[0]
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

func (x *Problem) GetType() atlas.Problem_Type {
	if x != nil {
		return x.Type
	}
	return atlas.Problem_Type(0)
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

func (x *Problem) GetBaseNumber() uint32 {
	if x != nil {
		return x.BaseNumber
	}
	return 0
}

func (x *Problem) GetBaseSpaceId() string {
	if x != nil {
		return x.BaseSpaceId
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
	state         protoimpl.MessageState `protogen:"open.v1"`
	Locale        string                 `protobuf:"bytes,1,opt,name=locale,proto3" json:"locale,omitempty"`
	Title         string                 `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Content       *ecm.Node              `protobuf:"bytes,4,opt,name=content,proto3" json:"content,omitempty"`
	DownloadLink  string                 `protobuf:"bytes,6,opt,name=download_link,json=downloadLink,proto3" json:"download_link,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Problem_Statement) Reset() {
	*x = Problem_Statement{}
	mi := &file_eolymp_judge_problem_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Problem_Statement) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Problem_Statement) ProtoMessage() {}

func (x *Problem_Statement) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_judge_problem_proto_msgTypes[1]
	if x != nil {
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

func (x *Problem_Statement) GetContent() *ecm.Node {
	if x != nil {
		return x.Content
	}
	return nil
}

func (x *Problem_Statement) GetDownloadLink() string {
	if x != nil {
		return x.DownloadLink
	}
	return ""
}

// Test defines set of input and answer data
type Problem_Test struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Index         uint32                 `protobuf:"varint,1,opt,name=index,proto3" json:"index,omitempty"`
	Example       bool                   `protobuf:"varint,2,opt,name=example,proto3" json:"example,omitempty"`
	InputUrl      string                 `protobuf:"bytes,13,opt,name=input_url,json=inputUrl,proto3" json:"input_url,omitempty"`
	AnswerUrl     string                 `protobuf:"bytes,14,opt,name=answer_url,json=answerUrl,proto3" json:"answer_url,omitempty"`
	Score         float32                `protobuf:"fixed32,5,opt,name=score,proto3" json:"score,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Problem_Test) Reset() {
	*x = Problem_Test{}
	mi := &file_eolymp_judge_problem_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Problem_Test) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Problem_Test) ProtoMessage() {}

func (x *Problem_Test) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_judge_problem_proto_msgTypes[2]
	if x != nil {
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

func (x *Problem_Test) GetInputUrl() string {
	if x != nil {
		return x.InputUrl
	}
	return ""
}

func (x *Problem_Test) GetAnswerUrl() string {
	if x != nil {
		return x.AnswerUrl
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
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name          string                 `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Link          string                 `protobuf:"bytes,4,opt,name=link,proto3" json:"link,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Problem_Attachment) Reset() {
	*x = Problem_Attachment{}
	mi := &file_eolymp_judge_problem_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Problem_Attachment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Problem_Attachment) ProtoMessage() {}

func (x *Problem_Attachment) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_judge_problem_proto_msgTypes[3]
	if x != nil {
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

const file_eolymp_judge_problem_proto_rawDesc = "" +
	"\n" +
	"\x1aeolymp/judge/problem.proto\x12\feolymp.judge\x1a\x1aeolymp/atlas/problem.proto\x1a#eolymp/atlas/testing_feedback.proto\x1a\x15eolymp/ecm/node.proto\"\x8c\a\n" +
	"\aProblem\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x12\x10\n" +
	"\x03url\x18\x06 \x01(\tR\x03url\x12.\n" +
	"\x04type\x18\x17 \x01(\x0e2\x1a.eolymp.atlas.Problem.TypeR\x04type\x12\x14\n" +
	"\x05index\x18\x02 \x01(\rR\x05index\x12\x14\n" +
	"\x05score\x18\x03 \x01(\x02R\x05score\x12\x17\n" +
	"\abase_id\x18\x04 \x01(\tR\x06baseId\x12\x1f\n" +
	"\vbase_number\x18- \x01(\rR\n" +
	"baseNumber\x12\"\n" +
	"\rbase_space_id\x18( \x01(\tR\vbaseSpaceId\x12\x1d\n" +
	"\n" +
	"contest_id\x18\x05 \x01(\tR\tcontestId\x12E\n" +
	"\x0ffeedback_policy\x18\n" +
	" \x01(\x0e2\x1c.eolymp.atlas.FeedbackPolicyR\x0efeedbackPolicy\x12\x1e\n" +
	"\n" +
	"time_limit\x18\xc8\x01 \x01(\rR\ttimeLimit\x12\x1c\n" +
	"\tcpu_limit\x18\xcc\x01 \x01(\rR\bcpuLimit\x12\"\n" +
	"\fmemory_limit\x18\xc9\x01 \x01(\x04R\vmemoryLimit\x12'\n" +
	"\x0ffile_size_limit\x18\xca\x01 \x01(\x04R\rfileSizeLimit\x12\"\n" +
	"\fsubmit_limit\x18\xcb\x01 \x01(\rR\vsubmitLimit\x122\n" +
	"\x15score_by_best_testset\x18\xd2\x01 \x01(\bR\x12scoreByBestTestset\x1a\x8a\x01\n" +
	"\tStatement\x12\x16\n" +
	"\x06locale\x18\x01 \x01(\tR\x06locale\x12\x14\n" +
	"\x05title\x18\x02 \x01(\tR\x05title\x12*\n" +
	"\acontent\x18\x04 \x01(\v2\x10.eolymp.ecm.NodeR\acontent\x12#\n" +
	"\rdownload_link\x18\x06 \x01(\tR\fdownloadLink\x1a\x88\x01\n" +
	"\x04Test\x12\x14\n" +
	"\x05index\x18\x01 \x01(\rR\x05index\x12\x18\n" +
	"\aexample\x18\x02 \x01(\bR\aexample\x12\x1b\n" +
	"\tinput_url\x18\r \x01(\tR\binputUrl\x12\x1d\n" +
	"\n" +
	"answer_url\x18\x0e \x01(\tR\tanswerUrl\x12\x14\n" +
	"\x05score\x18\x05 \x01(\x02R\x05score\x1aD\n" +
	"\n" +
	"Attachment\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x12\x12\n" +
	"\x04name\x18\x03 \x01(\tR\x04name\x12\x12\n" +
	"\x04link\x18\x04 \x01(\tR\x04linkB-Z+github.com/eolymp/go-sdk/eolymp/judge;judgeb\x06proto3"

var (
	file_eolymp_judge_problem_proto_rawDescOnce sync.Once
	file_eolymp_judge_problem_proto_rawDescData []byte
)

func file_eolymp_judge_problem_proto_rawDescGZIP() []byte {
	file_eolymp_judge_problem_proto_rawDescOnce.Do(func() {
		file_eolymp_judge_problem_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_eolymp_judge_problem_proto_rawDesc), len(file_eolymp_judge_problem_proto_rawDesc)))
	})
	return file_eolymp_judge_problem_proto_rawDescData
}

var file_eolymp_judge_problem_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_eolymp_judge_problem_proto_goTypes = []any{
	(*Problem)(nil),            // 0: eolymp.judge.Problem
	(*Problem_Statement)(nil),  // 1: eolymp.judge.Problem.Statement
	(*Problem_Test)(nil),       // 2: eolymp.judge.Problem.Test
	(*Problem_Attachment)(nil), // 3: eolymp.judge.Problem.Attachment
	(atlas.Problem_Type)(0),    // 4: eolymp.atlas.Problem.Type
	(atlas.FeedbackPolicy)(0),  // 5: eolymp.atlas.FeedbackPolicy
	(*ecm.Node)(nil),           // 6: eolymp.ecm.Node
}
var file_eolymp_judge_problem_proto_depIdxs = []int32{
	4, // 0: eolymp.judge.Problem.type:type_name -> eolymp.atlas.Problem.Type
	5, // 1: eolymp.judge.Problem.feedback_policy:type_name -> eolymp.atlas.FeedbackPolicy
	6, // 2: eolymp.judge.Problem.Statement.content:type_name -> eolymp.ecm.Node
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_eolymp_judge_problem_proto_rawDesc), len(file_eolymp_judge_problem_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_eolymp_judge_problem_proto_goTypes,
		DependencyIndexes: file_eolymp_judge_problem_proto_depIdxs,
		MessageInfos:      file_eolymp_judge_problem_proto_msgTypes,
	}.Build()
	File_eolymp_judge_problem_proto = out.File
	file_eolymp_judge_problem_proto_goTypes = nil
	file_eolymp_judge_problem_proto_depIdxs = nil
}
