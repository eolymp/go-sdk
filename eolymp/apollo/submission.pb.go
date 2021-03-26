// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.2
// source: eolymp/apollo/submission.proto

package apollo

import (
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type Submission_Status int32

const (
	Submission_NONE     Submission_Status = 0 // reserved, should not be used
	Submission_PENDING  Submission_Status = 1 // submission created and awaiting testing
	Submission_TESTING  Submission_Status = 2 // submission is being tested by an agent
	Submission_TIMEOUT  Submission_Status = 3 // submission testing took too long (too many timeouts in tests)
	Submission_COMPLETE Submission_Status = 4 // submission testing complete (score is populated)
	Submission_ERROR    Submission_Status = 5 // submission produced an error (eg. compilation error)
	Submission_FAILURE  Submission_Status = 6 // submission testing failed due to system error (or problem configuration error)
)

// Enum value maps for Submission_Status.
var (
	Submission_Status_name = map[int32]string{
		0: "NONE",
		1: "PENDING",
		2: "TESTING",
		3: "TIMEOUT",
		4: "COMPLETE",
		5: "ERROR",
		6: "FAILURE",
	}
	Submission_Status_value = map[string]int32{
		"NONE":     0,
		"PENDING":  1,
		"TESTING":  2,
		"TIMEOUT":  3,
		"COMPLETE": 4,
		"ERROR":    5,
		"FAILURE":  6,
	}
)

func (x Submission_Status) Enum() *Submission_Status {
	p := new(Submission_Status)
	*p = x
	return p
}

func (x Submission_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Submission_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_eolymp_apollo_submission_proto_enumTypes[0].Descriptor()
}

func (Submission_Status) Type() protoreflect.EnumType {
	return &file_eolymp_apollo_submission_proto_enumTypes[0]
}

func (x Submission_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Submission_Status.Descriptor instead.
func (Submission_Status) EnumDescriptor() ([]byte, []int) {
	return file_eolymp_apollo_submission_proto_rawDescGZIP(), []int{0, 0}
}

type Submission_Run_Status int32

const (
	Submission_Run_NONE               Submission_Run_Status = 0
	Submission_Run_CREATED            Submission_Run_Status = 1
	Submission_Run_EXECUTING          Submission_Run_Status = 2
	Submission_Run_TIMEOUT            Submission_Run_Status = 3
	Submission_Run_OVERFLOW           Submission_Run_Status = 4
	Submission_Run_WRONG_ANSWER       Submission_Run_Status = 5
	Submission_Run_ACCEPTED           Submission_Run_Status = 6
	Submission_Run_RUNTIME_ERROR      Submission_Run_Status = 7
	Submission_Run_VERIFICATION_ERROR Submission_Run_Status = 8
	Submission_Run_SKIPPED            Submission_Run_Status = 9
)

// Enum value maps for Submission_Run_Status.
var (
	Submission_Run_Status_name = map[int32]string{
		0: "NONE",
		1: "CREATED",
		2: "EXECUTING",
		3: "TIMEOUT",
		4: "OVERFLOW",
		5: "WRONG_ANSWER",
		6: "ACCEPTED",
		7: "RUNTIME_ERROR",
		8: "VERIFICATION_ERROR",
		9: "SKIPPED",
	}
	Submission_Run_Status_value = map[string]int32{
		"NONE":               0,
		"CREATED":            1,
		"EXECUTING":          2,
		"TIMEOUT":            3,
		"OVERFLOW":           4,
		"WRONG_ANSWER":       5,
		"ACCEPTED":           6,
		"RUNTIME_ERROR":      7,
		"VERIFICATION_ERROR": 8,
		"SKIPPED":            9,
	}
)

func (x Submission_Run_Status) Enum() *Submission_Run_Status {
	p := new(Submission_Run_Status)
	*p = x
	return p
}

func (x Submission_Run_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Submission_Run_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_eolymp_apollo_submission_proto_enumTypes[1].Descriptor()
}

func (Submission_Run_Status) Type() protoreflect.EnumType {
	return &file_eolymp_apollo_submission_proto_enumTypes[1]
}

func (x Submission_Run_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Submission_Run_Status.Descriptor instead.
func (Submission_Run_Status) EnumDescriptor() ([]byte, []int) {
	return file_eolymp_apollo_submission_proto_rawDescGZIP(), []int{0, 0, 0}
}

type Submission struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string               `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`                                                // unique identifier
	ProblemId   string               `protobuf:"bytes,2,opt,name=problem_id,json=problemId,proto3" json:"problem_id,omitempty"`                 // problem
	UserId      string               `protobuf:"bytes,3,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`                          // submitter
	SubmittedAt *timestamp.Timestamp `protobuf:"bytes,4,opt,name=submitted_at,json=submittedAt,proto3" json:"submitted_at,omitempty"`           // time when submission was created
	Lang        string               `protobuf:"bytes,10,opt,name=lang,proto3" json:"lang,omitempty"`                                           // programming language
	Source      string               `protobuf:"bytes,11,opt,name=source,proto3" json:"source,omitempty"`                                       // source code
	Status      Submission_Status    `protobuf:"varint,20,opt,name=status,proto3,enum=eolymp.apollo.Submission_Status" json:"status,omitempty"` // status (see explanation for enumeration values)
	Error       string               `protobuf:"bytes,21,opt,name=error,proto3" json:"error,omitempty"`                                         // error message in case status is ERROR
	Score       float32              `protobuf:"fixed32,30,opt,name=score,proto3" json:"score,omitempty"`                                       // score from 0 (none of the tests are passing) to 1 (all tests are passing)
	Runs        []*Submission_Run    `protobuf:"bytes,50,rep,name=runs,proto3" json:"runs,omitempty"`                                           // status for each run
}

func (x *Submission) Reset() {
	*x = Submission{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_apollo_submission_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Submission) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Submission) ProtoMessage() {}

func (x *Submission) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_apollo_submission_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Submission.ProtoReflect.Descriptor instead.
func (*Submission) Descriptor() ([]byte, []int) {
	return file_eolymp_apollo_submission_proto_rawDescGZIP(), []int{0}
}

func (x *Submission) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Submission) GetProblemId() string {
	if x != nil {
		return x.ProblemId
	}
	return ""
}

func (x *Submission) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *Submission) GetSubmittedAt() *timestamp.Timestamp {
	if x != nil {
		return x.SubmittedAt
	}
	return nil
}

func (x *Submission) GetLang() string {
	if x != nil {
		return x.Lang
	}
	return ""
}

func (x *Submission) GetSource() string {
	if x != nil {
		return x.Source
	}
	return ""
}

func (x *Submission) GetStatus() Submission_Status {
	if x != nil {
		return x.Status
	}
	return Submission_NONE
}

func (x *Submission) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

func (x *Submission) GetScore() float32 {
	if x != nil {
		return x.Score
	}
	return 0
}

func (x *Submission) GetRuns() []*Submission_Run {
	if x != nil {
		return x.Runs
	}
	return nil
}

type Submission_Run struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string                `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	TimeUsage   int32                 `protobuf:"varint,2,opt,name=time_usage,json=timeUsage,proto3" json:"time_usage,omitempty"`
	MemoryUsage int32                 `protobuf:"varint,3,opt,name=memory_usage,json=memoryUsage,proto3" json:"memory_usage,omitempty"`
	TestIndex   int32                 `protobuf:"varint,4,opt,name=test_index,json=testIndex,proto3" json:"test_index,omitempty"`
	TestId      string                `protobuf:"bytes,5,opt,name=test_id,json=testId,proto3" json:"test_id,omitempty"`
	Status      Submission_Run_Status `protobuf:"varint,6,opt,name=status,proto3,enum=eolymp.apollo.Submission_Run_Status" json:"status,omitempty"`
}

func (x *Submission_Run) Reset() {
	*x = Submission_Run{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_apollo_submission_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Submission_Run) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Submission_Run) ProtoMessage() {}

func (x *Submission_Run) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_apollo_submission_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Submission_Run.ProtoReflect.Descriptor instead.
func (*Submission_Run) Descriptor() ([]byte, []int) {
	return file_eolymp_apollo_submission_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Submission_Run) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Submission_Run) GetTimeUsage() int32 {
	if x != nil {
		return x.TimeUsage
	}
	return 0
}

func (x *Submission_Run) GetMemoryUsage() int32 {
	if x != nil {
		return x.MemoryUsage
	}
	return 0
}

func (x *Submission_Run) GetTestIndex() int32 {
	if x != nil {
		return x.TestIndex
	}
	return 0
}

func (x *Submission_Run) GetTestId() string {
	if x != nil {
		return x.TestId
	}
	return ""
}

func (x *Submission_Run) GetStatus() Submission_Run_Status {
	if x != nil {
		return x.Status
	}
	return Submission_Run_NONE
}

var File_eolymp_apollo_submission_proto protoreflect.FileDescriptor

var file_eolymp_apollo_submission_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x61, 0x70, 0x6f, 0x6c, 0x6c, 0x6f, 0x2f,
	0x73, 0x75, 0x62, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0d, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x61, 0x70, 0x6f, 0x6c, 0x6c, 0x6f, 0x1a,
	0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xad, 0x06, 0x0a, 0x0a, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x49, 0x64, 0x12, 0x17,
	0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x3d, 0x0a, 0x0c, 0x73, 0x75, 0x62, 0x6d, 0x69,
	0x74, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0b, 0x73, 0x75, 0x62, 0x6d, 0x69,
	0x74, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x61, 0x6e, 0x67, 0x18, 0x0a,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6c, 0x61, 0x6e, 0x67, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x12, 0x38, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x14, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x20, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x61, 0x70, 0x6f, 0x6c,
	0x6c, 0x6f, 0x2e, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x14, 0x0a, 0x05,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x15, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x1e, 0x20, 0x01, 0x28,
	0x02, 0x52, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x12, 0x31, 0x0a, 0x04, 0x72, 0x75, 0x6e, 0x73,
	0x18, 0x32, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e,
	0x61, 0x70, 0x6f, 0x6c, 0x6c, 0x6f, 0x2e, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x2e, 0x52, 0x75, 0x6e, 0x52, 0x04, 0x72, 0x75, 0x6e, 0x73, 0x1a, 0xf1, 0x02, 0x0a, 0x03,
	0x52, 0x75, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x75, 0x73, 0x61, 0x67,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x55, 0x73, 0x61,
	0x67, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x5f, 0x75, 0x73, 0x61,
	0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79,
	0x55, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x69, 0x6e,
	0x64, 0x65, 0x78, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x74, 0x65, 0x73, 0x74, 0x49,
	0x6e, 0x64, 0x65, 0x78, 0x12, 0x17, 0x0a, 0x07, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x65, 0x73, 0x74, 0x49, 0x64, 0x12, 0x3c, 0x0a,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x24, 0x2e,
	0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x61, 0x70, 0x6f, 0x6c, 0x6c, 0x6f, 0x2e, 0x53, 0x75,
	0x62, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x52, 0x75, 0x6e, 0x2e, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0xa1, 0x01, 0x0a, 0x06,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00,
	0x12, 0x0b, 0x0a, 0x07, 0x43, 0x52, 0x45, 0x41, 0x54, 0x45, 0x44, 0x10, 0x01, 0x12, 0x0d, 0x0a,
	0x09, 0x45, 0x58, 0x45, 0x43, 0x55, 0x54, 0x49, 0x4e, 0x47, 0x10, 0x02, 0x12, 0x0b, 0x0a, 0x07,
	0x54, 0x49, 0x4d, 0x45, 0x4f, 0x55, 0x54, 0x10, 0x03, 0x12, 0x0c, 0x0a, 0x08, 0x4f, 0x56, 0x45,
	0x52, 0x46, 0x4c, 0x4f, 0x57, 0x10, 0x04, 0x12, 0x10, 0x0a, 0x0c, 0x57, 0x52, 0x4f, 0x4e, 0x47,
	0x5f, 0x41, 0x4e, 0x53, 0x57, 0x45, 0x52, 0x10, 0x05, 0x12, 0x0c, 0x0a, 0x08, 0x41, 0x43, 0x43,
	0x45, 0x50, 0x54, 0x45, 0x44, 0x10, 0x06, 0x12, 0x11, 0x0a, 0x0d, 0x52, 0x55, 0x4e, 0x54, 0x49,
	0x4d, 0x45, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x07, 0x12, 0x16, 0x0a, 0x12, 0x56, 0x45,
	0x52, 0x49, 0x46, 0x49, 0x43, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52,
	0x10, 0x08, 0x12, 0x0b, 0x0a, 0x07, 0x53, 0x4b, 0x49, 0x50, 0x50, 0x45, 0x44, 0x10, 0x09, 0x22,
	0x5f, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x4f, 0x4e,
	0x45, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x50, 0x45, 0x4e, 0x44, 0x49, 0x4e, 0x47, 0x10, 0x01,
	0x12, 0x0b, 0x0a, 0x07, 0x54, 0x45, 0x53, 0x54, 0x49, 0x4e, 0x47, 0x10, 0x02, 0x12, 0x0b, 0x0a,
	0x07, 0x54, 0x49, 0x4d, 0x45, 0x4f, 0x55, 0x54, 0x10, 0x03, 0x12, 0x0c, 0x0a, 0x08, 0x43, 0x4f,
	0x4d, 0x50, 0x4c, 0x45, 0x54, 0x45, 0x10, 0x04, 0x12, 0x09, 0x0a, 0x05, 0x45, 0x52, 0x52, 0x4f,
	0x52, 0x10, 0x05, 0x12, 0x0b, 0x0a, 0x07, 0x46, 0x41, 0x49, 0x4c, 0x55, 0x52, 0x45, 0x10, 0x06,
	0x42, 0x35, 0x5a, 0x33, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65,
	0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x73, 0x2f,
	0x67, 0x6f, 0x2f, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x61, 0x70, 0x6f, 0x6c, 0x6c, 0x6f,
	0x3b, 0x61, 0x70, 0x6f, 0x6c, 0x6c, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_eolymp_apollo_submission_proto_rawDescOnce sync.Once
	file_eolymp_apollo_submission_proto_rawDescData = file_eolymp_apollo_submission_proto_rawDesc
)

func file_eolymp_apollo_submission_proto_rawDescGZIP() []byte {
	file_eolymp_apollo_submission_proto_rawDescOnce.Do(func() {
		file_eolymp_apollo_submission_proto_rawDescData = protoimpl.X.CompressGZIP(file_eolymp_apollo_submission_proto_rawDescData)
	})
	return file_eolymp_apollo_submission_proto_rawDescData
}

var file_eolymp_apollo_submission_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_eolymp_apollo_submission_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_eolymp_apollo_submission_proto_goTypes = []interface{}{
	(Submission_Status)(0),      // 0: eolymp.apollo.Submission.Status
	(Submission_Run_Status)(0),  // 1: eolymp.apollo.Submission.Run.Status
	(*Submission)(nil),          // 2: eolymp.apollo.Submission
	(*Submission_Run)(nil),      // 3: eolymp.apollo.Submission.Run
	(*timestamp.Timestamp)(nil), // 4: google.protobuf.Timestamp
}
var file_eolymp_apollo_submission_proto_depIdxs = []int32{
	4, // 0: eolymp.apollo.Submission.submitted_at:type_name -> google.protobuf.Timestamp
	0, // 1: eolymp.apollo.Submission.status:type_name -> eolymp.apollo.Submission.Status
	3, // 2: eolymp.apollo.Submission.runs:type_name -> eolymp.apollo.Submission.Run
	1, // 3: eolymp.apollo.Submission.Run.status:type_name -> eolymp.apollo.Submission.Run.Status
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_eolymp_apollo_submission_proto_init() }
func file_eolymp_apollo_submission_proto_init() {
	if File_eolymp_apollo_submission_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_eolymp_apollo_submission_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Submission); i {
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
		file_eolymp_apollo_submission_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Submission_Run); i {
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
			RawDescriptor: file_eolymp_apollo_submission_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_eolymp_apollo_submission_proto_goTypes,
		DependencyIndexes: file_eolymp_apollo_submission_proto_depIdxs,
		EnumInfos:         file_eolymp_apollo_submission_proto_enumTypes,
		MessageInfos:      file_eolymp_apollo_submission_proto_msgTypes,
	}.Build()
	File_eolymp_apollo_submission_proto = out.File
	file_eolymp_apollo_submission_proto_rawDesc = nil
	file_eolymp_apollo_submission_proto_goTypes = nil
	file_eolymp_apollo_submission_proto_depIdxs = nil
}
