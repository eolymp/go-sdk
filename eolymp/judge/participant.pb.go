// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v3.21.12
// source: eolymp/judge/participant.proto

package judge

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Participant_Status int32

const (
	Participant_UNKNOWN_STATUS Participant_Status = 0 // reserved, should not be used
	Participant_WAITING        Participant_Status = 5 // participant can not start yet
	Participant_READY          Participant_Status = 1 // participant is ready to start (via StartContest method)
	Participant_ACTIVE         Participant_Status = 2 // participant is participating in the contest
	Participant_COMPLETE       Participant_Status = 3 // participant has completed contest (contest is complete, time run out etc)
	Participant_UPSOLVE        Participant_Status = 7 // participant has completed contest but can continue solve problems in upsolve mode
	Participant_BLOCKED        Participant_Status = 8 // participant can not participate (participant is blocked or contest is suspended)
)

// Enum value maps for Participant_Status.
var (
	Participant_Status_name = map[int32]string{
		0: "UNKNOWN_STATUS",
		5: "WAITING",
		1: "READY",
		2: "ACTIVE",
		3: "COMPLETE",
		7: "UPSOLVE",
		8: "BLOCKED",
	}
	Participant_Status_value = map[string]int32{
		"UNKNOWN_STATUS": 0,
		"WAITING":        5,
		"READY":          1,
		"ACTIVE":         2,
		"COMPLETE":       3,
		"UPSOLVE":        7,
		"BLOCKED":        8,
	}
)

func (x Participant_Status) Enum() *Participant_Status {
	p := new(Participant_Status)
	*p = x
	return p
}

func (x Participant_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Participant_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_eolymp_judge_participant_proto_enumTypes[0].Descriptor()
}

func (Participant_Status) Type() protoreflect.EnumType {
	return &file_eolymp_judge_participant_proto_enumTypes[0]
}

func (x Participant_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Participant_Status.Descriptor instead.
func (Participant_Status) EnumDescriptor() ([]byte, []int) {
	return file_eolymp_judge_participant_proto_rawDescGZIP(), []int{0, 0}
}

type Participant struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`                                      // unique identifier of the participant (can not be set when creating participant)
	ContestId    string `protobuf:"bytes,3,opt,name=contest_id,json=contestId,proto3" json:"contest_id,omitempty"`       // contest
	MemberId     string `protobuf:"bytes,4,opt,name=member_id,json=memberId,proto3" json:"member_id,omitempty"`          // community member
	DisplayName  string `protobuf:"bytes,5,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"` // display name
	Unofficial   bool   `protobuf:"varint,6,opt,name=unofficial,proto3" json:"unofficial,omitempty"`                     // if true, participant won't be assigned rank in scoreboard
	Inactive     bool   `protobuf:"varint,10,opt,name=inactive,proto3" json:"inactive,omitempty"`                        // participant is inactive
	Disqualified bool   `protobuf:"varint,11,opt,name=disqualified,proto3" json:"disqualified,omitempty"`                // participant is disqualified
	Ghost        bool   `protobuf:"varint,8,opt,name=ghost,proto3" json:"ghost,omitempty"`                               // participant is a ghost
	Medal        Medal  `protobuf:"varint,9,opt,name=medal,proto3,enum=eolymp.judge.Medal" json:"medal,omitempty"`
	// status (see explanation to enumeration values)
	Status    Participant_Status     `protobuf:"varint,20,opt,name=status,proto3,enum=eolymp.judge.Participant_Status" json:"status,omitempty"`
	StartedAt *timestamppb.Timestamp `protobuf:"bytes,21,opt,name=started_at,json=startedAt,proto3" json:"started_at,omitempty"`  // time when user has started participating
	StartedIn uint32                 `protobuf:"varint,22,opt,name=started_in,json=startedIn,proto3" json:"started_in,omitempty"` // time in seconds when user has started participating
	EndAt     *timestamppb.Timestamp `protobuf:"bytes,25,opt,name=end_at,json=endAt,proto3" json:"end_at,omitempty"`              // time when user has finished (will finish) participating (incl. bonus time)
	EndIn     uint32                 `protobuf:"varint,26,opt,name=end_in,json=endIn,proto3" json:"end_in,omitempty"`             // time in seconds when user has finished (will finish) participating (incl. bonus time)
	BonusTime uint32                 `protobuf:"varint,27,opt,name=bonus_time,json=bonusTime,proto3" json:"bonus_time,omitempty"` // additional time in seconds for participation
	// passcode is a code participant has to enter before she can begin contest
	// this field is only populated when request is made by contest owner
	// passcode is read-only and should be set using ResetPasscode method.
	Passcode string `protobuf:"bytes,30,opt,name=passcode,proto3" json:"passcode,omitempty"`
	// Submit counter is used to count how many times user submitted the problem.
	Submits []*Participant_Submit `protobuf:"bytes,40,rep,name=submits,proto3" json:"submits,omitempty"`
}

func (x *Participant) Reset() {
	*x = Participant{}
	mi := &file_eolymp_judge_participant_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Participant) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Participant) ProtoMessage() {}

func (x *Participant) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_judge_participant_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Participant.ProtoReflect.Descriptor instead.
func (*Participant) Descriptor() ([]byte, []int) {
	return file_eolymp_judge_participant_proto_rawDescGZIP(), []int{0}
}

func (x *Participant) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Participant) GetContestId() string {
	if x != nil {
		return x.ContestId
	}
	return ""
}

func (x *Participant) GetMemberId() string {
	if x != nil {
		return x.MemberId
	}
	return ""
}

func (x *Participant) GetDisplayName() string {
	if x != nil {
		return x.DisplayName
	}
	return ""
}

func (x *Participant) GetUnofficial() bool {
	if x != nil {
		return x.Unofficial
	}
	return false
}

func (x *Participant) GetInactive() bool {
	if x != nil {
		return x.Inactive
	}
	return false
}

func (x *Participant) GetDisqualified() bool {
	if x != nil {
		return x.Disqualified
	}
	return false
}

func (x *Participant) GetGhost() bool {
	if x != nil {
		return x.Ghost
	}
	return false
}

func (x *Participant) GetMedal() Medal {
	if x != nil {
		return x.Medal
	}
	return Medal_NO_MEDAL
}

func (x *Participant) GetStatus() Participant_Status {
	if x != nil {
		return x.Status
	}
	return Participant_UNKNOWN_STATUS
}

func (x *Participant) GetStartedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.StartedAt
	}
	return nil
}

func (x *Participant) GetStartedIn() uint32 {
	if x != nil {
		return x.StartedIn
	}
	return 0
}

func (x *Participant) GetEndAt() *timestamppb.Timestamp {
	if x != nil {
		return x.EndAt
	}
	return nil
}

func (x *Participant) GetEndIn() uint32 {
	if x != nil {
		return x.EndIn
	}
	return 0
}

func (x *Participant) GetBonusTime() uint32 {
	if x != nil {
		return x.BonusTime
	}
	return 0
}

func (x *Participant) GetPasscode() string {
	if x != nil {
		return x.Passcode
	}
	return ""
}

func (x *Participant) GetSubmits() []*Participant_Submit {
	if x != nil {
		return x.Submits
	}
	return nil
}

type Participant_Submit struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProblemId string `protobuf:"bytes,1,opt,name=problem_id,json=problemId,proto3" json:"problem_id,omitempty"`
	Counter   uint32 `protobuf:"varint,2,opt,name=counter,proto3" json:"counter,omitempty"`
}

func (x *Participant_Submit) Reset() {
	*x = Participant_Submit{}
	mi := &file_eolymp_judge_participant_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Participant_Submit) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Participant_Submit) ProtoMessage() {}

func (x *Participant_Submit) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_judge_participant_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Participant_Submit.ProtoReflect.Descriptor instead.
func (*Participant_Submit) Descriptor() ([]byte, []int) {
	return file_eolymp_judge_participant_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Participant_Submit) GetProblemId() string {
	if x != nil {
		return x.ProblemId
	}
	return ""
}

func (x *Participant_Submit) GetCounter() uint32 {
	if x != nil {
		return x.Counter
	}
	return 0
}

var File_eolymp_judge_participant_proto protoreflect.FileDescriptor

var file_eolymp_judge_participant_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x6a, 0x75, 0x64, 0x67, 0x65, 0x2f, 0x70,
	0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0c, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x6a, 0x75, 0x64, 0x67, 0x65, 0x1a, 0x18,
	0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x6a, 0x75, 0x64, 0x67, 0x65, 0x2f, 0x6d, 0x65, 0x64,
	0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9f, 0x06, 0x0a, 0x0b, 0x50, 0x61,
	0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6f, 0x6e,
	0x74, 0x65, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x73, 0x74, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x65, 0x6d, 0x62,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6d, 0x65, 0x6d,
	0x62, 0x65, 0x72, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x69, 0x73,
	0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x75, 0x6e, 0x6f, 0x66,
	0x66, 0x69, 0x63, 0x69, 0x61, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x75, 0x6e,
	0x6f, 0x66, 0x66, 0x69, 0x63, 0x69, 0x61, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x6e, 0x61, 0x63,
	0x74, 0x69, 0x76, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x69, 0x6e, 0x61, 0x63,
	0x74, 0x69, 0x76, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x64, 0x69, 0x73, 0x71, 0x75, 0x61, 0x6c, 0x69,
	0x66, 0x69, 0x65, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x64, 0x69, 0x73, 0x71,
	0x75, 0x61, 0x6c, 0x69, 0x66, 0x69, 0x65, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x67, 0x68, 0x6f, 0x73,
	0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x67, 0x68, 0x6f, 0x73, 0x74, 0x12, 0x29,
	0x0a, 0x05, 0x6d, 0x65, 0x64, 0x61, 0x6c, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x13, 0x2e,
	0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x6a, 0x75, 0x64, 0x67, 0x65, 0x2e, 0x4d, 0x65, 0x64,
	0x61, 0x6c, 0x52, 0x05, 0x6d, 0x65, 0x64, 0x61, 0x6c, 0x12, 0x38, 0x0a, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x14, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x20, 0x2e, 0x65, 0x6f, 0x6c, 0x79,
	0x6d, 0x70, 0x2e, 0x6a, 0x75, 0x64, 0x67, 0x65, 0x2e, 0x50, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69,
	0x70, 0x61, 0x6e, 0x74, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x39, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x65, 0x64, 0x5f, 0x61,
	0x74, 0x18, 0x15, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d,
	0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x65, 0x64, 0x5f, 0x69, 0x6e, 0x18, 0x16, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x65, 0x64, 0x49, 0x6e, 0x12, 0x31, 0x0a,
	0x06, 0x65, 0x6e, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x19, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x05, 0x65, 0x6e, 0x64, 0x41, 0x74,
	0x12, 0x15, 0x0a, 0x06, 0x65, 0x6e, 0x64, 0x5f, 0x69, 0x6e, 0x18, 0x1a, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x05, 0x65, 0x6e, 0x64, 0x49, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x62, 0x6f, 0x6e, 0x75, 0x73,
	0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x1b, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x62, 0x6f, 0x6e,
	0x75, 0x73, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x63, 0x6f,
	0x64, 0x65, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x63, 0x6f,
	0x64, 0x65, 0x12, 0x3a, 0x0a, 0x07, 0x73, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x73, 0x18, 0x28, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x6a, 0x75, 0x64,
	0x67, 0x65, 0x2e, 0x50, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x2e, 0x53,
	0x75, 0x62, 0x6d, 0x69, 0x74, 0x52, 0x07, 0x73, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x73, 0x1a, 0x41,
	0x0a, 0x06, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x62,
	0x6c, 0x65, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72,
	0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65,
	0x72, 0x22, 0x68, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x12, 0x0a, 0x0e, 0x55,
	0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x10, 0x00, 0x12,
	0x0b, 0x0a, 0x07, 0x57, 0x41, 0x49, 0x54, 0x49, 0x4e, 0x47, 0x10, 0x05, 0x12, 0x09, 0x0a, 0x05,
	0x52, 0x45, 0x41, 0x44, 0x59, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x41, 0x43, 0x54, 0x49, 0x56,
	0x45, 0x10, 0x02, 0x12, 0x0c, 0x0a, 0x08, 0x43, 0x4f, 0x4d, 0x50, 0x4c, 0x45, 0x54, 0x45, 0x10,
	0x03, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x50, 0x53, 0x4f, 0x4c, 0x56, 0x45, 0x10, 0x07, 0x12, 0x0b,
	0x0a, 0x07, 0x42, 0x4c, 0x4f, 0x43, 0x4b, 0x45, 0x44, 0x10, 0x08, 0x42, 0x2d, 0x5a, 0x2b, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70,
	0x2f, 0x67, 0x6f, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x6a,
	0x75, 0x64, 0x67, 0x65, 0x3b, 0x6a, 0x75, 0x64, 0x67, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_eolymp_judge_participant_proto_rawDescOnce sync.Once
	file_eolymp_judge_participant_proto_rawDescData = file_eolymp_judge_participant_proto_rawDesc
)

func file_eolymp_judge_participant_proto_rawDescGZIP() []byte {
	file_eolymp_judge_participant_proto_rawDescOnce.Do(func() {
		file_eolymp_judge_participant_proto_rawDescData = protoimpl.X.CompressGZIP(file_eolymp_judge_participant_proto_rawDescData)
	})
	return file_eolymp_judge_participant_proto_rawDescData
}

var file_eolymp_judge_participant_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_eolymp_judge_participant_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_eolymp_judge_participant_proto_goTypes = []any{
	(Participant_Status)(0),       // 0: eolymp.judge.Participant.Status
	(*Participant)(nil),           // 1: eolymp.judge.Participant
	(*Participant_Submit)(nil),    // 2: eolymp.judge.Participant.Submit
	(Medal)(0),                    // 3: eolymp.judge.Medal
	(*timestamppb.Timestamp)(nil), // 4: google.protobuf.Timestamp
}
var file_eolymp_judge_participant_proto_depIdxs = []int32{
	3, // 0: eolymp.judge.Participant.medal:type_name -> eolymp.judge.Medal
	0, // 1: eolymp.judge.Participant.status:type_name -> eolymp.judge.Participant.Status
	4, // 2: eolymp.judge.Participant.started_at:type_name -> google.protobuf.Timestamp
	4, // 3: eolymp.judge.Participant.end_at:type_name -> google.protobuf.Timestamp
	2, // 4: eolymp.judge.Participant.submits:type_name -> eolymp.judge.Participant.Submit
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_eolymp_judge_participant_proto_init() }
func file_eolymp_judge_participant_proto_init() {
	if File_eolymp_judge_participant_proto != nil {
		return
	}
	file_eolymp_judge_medal_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_eolymp_judge_participant_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_eolymp_judge_participant_proto_goTypes,
		DependencyIndexes: file_eolymp_judge_participant_proto_depIdxs,
		EnumInfos:         file_eolymp_judge_participant_proto_enumTypes,
		MessageInfos:      file_eolymp_judge_participant_proto_msgTypes,
	}.Build()
	File_eolymp_judge_participant_proto = out.File
	file_eolymp_judge_participant_proto_rawDesc = nil
	file_eolymp_judge_participant_proto_goTypes = nil
	file_eolymp_judge_participant_proto_depIdxs = nil
}
