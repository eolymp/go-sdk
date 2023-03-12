// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.29.0
// 	protoc        v3.18.1
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
	Participant_NONE         Participant_Status = 0 // reserved, should not be used
	Participant_READY        Participant_Status = 1 // participant added to the contest and can start
	Participant_ACTIVE       Participant_Status = 2 // participant has started participating in the contest
	Participant_COMPLETE     Participant_Status = 3 // participant has completed contest (contest is complete, time run out etc)
	Participant_INACTIVE     Participant_Status = 4 // participant has been "disabled", profile does not show up in results and user can not participate
	Participant_UNREGISTERED Participant_Status = 5 // participant added to the contest but needs to fill registration form
	Participant_GHOST        Participant_Status = 6 // participant assigned to a ghost member
)

// Enum value maps for Participant_Status.
var (
	Participant_Status_name = map[int32]string{
		0: "NONE",
		1: "READY",
		2: "ACTIVE",
		3: "COMPLETE",
		4: "INACTIVE",
		5: "UNREGISTERED",
		6: "GHOST",
	}
	Participant_Status_value = map[string]int32{
		"NONE":         0,
		"READY":        1,
		"ACTIVE":       2,
		"COMPLETE":     3,
		"INACTIVE":     4,
		"UNREGISTERED": 5,
		"GHOST":        6,
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

	Id               string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"` // unique identifier of the participant (can not be set when creating participant)
	Ern              string `protobuf:"bytes,9999,opt,name=ern,proto3" json:"ern,omitempty"`
	ContestId        string `protobuf:"bytes,3,opt,name=contest_id,json=contestId,proto3" json:"contest_id,omitempty"`                         // contest
	MemberId         string `protobuf:"bytes,4,opt,name=member_id,json=memberId,proto3" json:"member_id,omitempty"`                            // community member
	Name             string `protobuf:"bytes,5,opt,name=name,proto3" json:"name,omitempty"`                                                    // name
	OutOfCompetition bool   `protobuf:"varint,6,opt,name=out_of_competition,json=outOfCompetition,proto3" json:"out_of_competition,omitempty"` // if true, participant won't be assigned rank in scoreboard
	// status (see explanation to enumeration values)
	Status     Participant_Status     `protobuf:"varint,20,opt,name=status,proto3,enum=eolymp.judge.Participant_Status" json:"status,omitempty"`
	StartedAt  *timestamppb.Timestamp `protobuf:"bytes,21,opt,name=started_at,json=startedAt,proto3" json:"started_at,omitempty"`     // time when user has started participating
	StartedIn  uint32                 `protobuf:"varint,22,opt,name=started_in,json=startedIn,proto3" json:"started_in,omitempty"`    // time in seconds when user has started participating
	EndAt      *timestamppb.Timestamp `protobuf:"bytes,23,opt,name=end_at,json=endAt,proto3" json:"end_at,omitempty"`                 // time when user has finished (will finish) participating (excl. bonus time)
	EndIn      uint32                 `protobuf:"varint,24,opt,name=end_in,json=endIn,proto3" json:"end_in,omitempty"`                // time in seconds when user has finished (will finish) participating (excl. bonus time)
	CompleteAt *timestamppb.Timestamp `protobuf:"bytes,25,opt,name=complete_at,json=completeAt,proto3" json:"complete_at,omitempty"`  // time when user has finished (will finish) participating (incl. bonus time)
	CompleteIn uint32                 `protobuf:"varint,26,opt,name=complete_in,json=completeIn,proto3" json:"complete_in,omitempty"` // time in seconds when user has finished (will finish) participating  (incl. bonus time)
	BonusTime  uint32                 `protobuf:"varint,27,opt,name=bonus_time,json=bonusTime,proto3" json:"bonus_time,omitempty"`    // additional time in seconds for participation (included in complete_at and complete_in)
	// passcode is a code participant has to enter before she can begin contest
	// this field is only populated when request is made by contest owner
	// passcode is read-only and should be set using ResetPasscode method.
	Passcode string `protobuf:"bytes,30,opt,name=passcode,proto3" json:"passcode,omitempty"`
	// Submit counter is used to count how many times user submitted the problem.
	Submits []*Participant_Submit `protobuf:"bytes,40,rep,name=submits,proto3" json:"submits,omitempty"`
}

func (x *Participant) Reset() {
	*x = Participant{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_judge_participant_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Participant) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Participant) ProtoMessage() {}

func (x *Participant) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_judge_participant_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
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

func (x *Participant) GetErn() string {
	if x != nil {
		return x.Ern
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

func (x *Participant) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Participant) GetOutOfCompetition() bool {
	if x != nil {
		return x.OutOfCompetition
	}
	return false
}

func (x *Participant) GetStatus() Participant_Status {
	if x != nil {
		return x.Status
	}
	return Participant_NONE
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

func (x *Participant) GetCompleteAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CompleteAt
	}
	return nil
}

func (x *Participant) GetCompleteIn() uint32 {
	if x != nil {
		return x.CompleteIn
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
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_judge_participant_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Participant_Submit) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Participant_Submit) ProtoMessage() {}

func (x *Participant_Submit) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_judge_participant_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
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
	0x12, 0x0c, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x6a, 0x75, 0x64, 0x67, 0x65, 0x1a, 0x1f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x88, 0x06, 0x0a, 0x0b, 0x50, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x11, 0x0a, 0x03, 0x65, 0x72, 0x6e, 0x18, 0x8f, 0x4e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x65,
	0x72, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x69, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x73, 0x74, 0x49,
	0x64, 0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x49, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x2c, 0x0a, 0x12, 0x6f, 0x75, 0x74, 0x5f, 0x6f, 0x66, 0x5f, 0x63, 0x6f, 0x6d,
	0x70, 0x65, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x10,
	0x6f, 0x75, 0x74, 0x4f, 0x66, 0x43, 0x6f, 0x6d, 0x70, 0x65, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x38, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x14, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x20, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x6a, 0x75, 0x64, 0x67, 0x65, 0x2e,
	0x50, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x2e, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x39, 0x0a, 0x0a, 0x73, 0x74,
	0x61, 0x72, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x15, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x65, 0x64,
	0x5f, 0x69, 0x6e, 0x18, 0x16, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74,
	0x65, 0x64, 0x49, 0x6e, 0x12, 0x31, 0x0a, 0x06, 0x65, 0x6e, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x17,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x05, 0x65, 0x6e, 0x64, 0x41, 0x74, 0x12, 0x15, 0x0a, 0x06, 0x65, 0x6e, 0x64, 0x5f, 0x69,
	0x6e, 0x18, 0x18, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x65, 0x6e, 0x64, 0x49, 0x6e, 0x12, 0x3b,
	0x0a, 0x0b, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x5f, 0x61, 0x74, 0x18, 0x19, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x0a, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x63,
	0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x5f, 0x69, 0x6e, 0x18, 0x1a, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x0a, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x49, 0x6e, 0x12, 0x1d, 0x0a, 0x0a,
	0x62, 0x6f, 0x6e, 0x75, 0x73, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x1b, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x09, 0x62, 0x6f, 0x6e, 0x75, 0x73, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70,
	0x61, 0x73, 0x73, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70,
	0x61, 0x73, 0x73, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x3a, 0x0a, 0x07, 0x73, 0x75, 0x62, 0x6d, 0x69,
	0x74, 0x73, 0x18, 0x28, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d,
	0x70, 0x2e, 0x6a, 0x75, 0x64, 0x67, 0x65, 0x2e, 0x50, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70,
	0x61, 0x6e, 0x74, 0x2e, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x52, 0x07, 0x73, 0x75, 0x62, 0x6d,
	0x69, 0x74, 0x73, 0x1a, 0x41, 0x0a, 0x06, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x12, 0x1d, 0x0a,
	0x0a, 0x70, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x22, 0x62, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x08, 0x0a, 0x04, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x52, 0x45,
	0x41, 0x44, 0x59, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x41, 0x43, 0x54, 0x49, 0x56, 0x45, 0x10,
	0x02, 0x12, 0x0c, 0x0a, 0x08, 0x43, 0x4f, 0x4d, 0x50, 0x4c, 0x45, 0x54, 0x45, 0x10, 0x03, 0x12,
	0x0c, 0x0a, 0x08, 0x49, 0x4e, 0x41, 0x43, 0x54, 0x49, 0x56, 0x45, 0x10, 0x04, 0x12, 0x10, 0x0a,
	0x0c, 0x55, 0x4e, 0x52, 0x45, 0x47, 0x49, 0x53, 0x54, 0x45, 0x52, 0x45, 0x44, 0x10, 0x05, 0x12,
	0x09, 0x0a, 0x05, 0x47, 0x48, 0x4f, 0x53, 0x54, 0x10, 0x06, 0x42, 0x2d, 0x5a, 0x2b, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f,
	0x67, 0x6f, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x6a, 0x75,
	0x64, 0x67, 0x65, 0x3b, 0x6a, 0x75, 0x64, 0x67, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
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
var file_eolymp_judge_participant_proto_goTypes = []interface{}{
	(Participant_Status)(0),       // 0: eolymp.judge.Participant.Status
	(*Participant)(nil),           // 1: eolymp.judge.Participant
	(*Participant_Submit)(nil),    // 2: eolymp.judge.Participant.Submit
	(*timestamppb.Timestamp)(nil), // 3: google.protobuf.Timestamp
}
var file_eolymp_judge_participant_proto_depIdxs = []int32{
	0, // 0: eolymp.judge.Participant.status:type_name -> eolymp.judge.Participant.Status
	3, // 1: eolymp.judge.Participant.started_at:type_name -> google.protobuf.Timestamp
	3, // 2: eolymp.judge.Participant.end_at:type_name -> google.protobuf.Timestamp
	3, // 3: eolymp.judge.Participant.complete_at:type_name -> google.protobuf.Timestamp
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
	if !protoimpl.UnsafeEnabled {
		file_eolymp_judge_participant_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Participant); i {
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
		file_eolymp_judge_participant_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Participant_Submit); i {
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
