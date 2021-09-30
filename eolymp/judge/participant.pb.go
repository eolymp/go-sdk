// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.12.2
// source: eolymp/judge/participant.proto

package judge

import (
	duration "github.com/golang/protobuf/ptypes/duration"
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

type Participant_Status int32

const (
	Participant_NONE         Participant_Status = 0 // reserved, should not be used
	Participant_READY        Participant_Status = 1 // participant added to the contest and can start
	Participant_ACTIVE       Participant_Status = 2 // participant has started participating in the contest
	Participant_COMPLETE     Participant_Status = 3 // participant has completed contest (contest is complete, time run out etc)
	Participant_INACTIVE     Participant_Status = 4 // participant has been "disabled", profile does not show up in results and user can not participate
	Participant_UNREGISTERED Participant_Status = 5 // participant added to the contest but needs to fill registration form
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
	}
	Participant_Status_value = map[string]int32{
		"NONE":         0,
		"READY":        1,
		"ACTIVE":       2,
		"COMPLETE":     3,
		"INACTIVE":     4,
		"UNREGISTERED": 5,
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

	Id        string              `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`                                // unique identifier of the participant (can not be set when creating participant)
	ContestId string              `protobuf:"bytes,3,opt,name=contest_id,json=contestId,proto3" json:"contest_id,omitempty"` // contest
	Name      string              `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`                            // name
	Users     []*Participant_User `protobuf:"bytes,5,rep,name=users,proto3" json:"users,omitempty"`
	UserId    string              `protobuf:"bytes,9998,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"` // deprecated: user identifier
	Username  string              `protobuf:"bytes,9999,opt,name=username,proto3" json:"username,omitempty"`           // deprecated: username
	// status (see explanation to enumeration values)
	Status     Participant_Status   `protobuf:"varint,20,opt,name=status,proto3,enum=eolymp.judge.Participant_Status" json:"status,omitempty"`
	StartedAt  *timestamp.Timestamp `protobuf:"bytes,21,opt,name=started_at,json=startedAt,proto3" json:"started_at,omitempty"`    // time when user has started participating
	StartedIn  *duration.Duration   `protobuf:"bytes,22,opt,name=started_in,json=startedIn,proto3" json:"started_in,omitempty"`    // time when user has started participating
	CompleteAt *timestamp.Timestamp `protobuf:"bytes,25,opt,name=complete_at,json=completeAt,proto3" json:"complete_at,omitempty"` // time when user has finished (will finish) participating
	CompleteIn *duration.Duration   `protobuf:"bytes,26,opt,name=complete_in,json=completeIn,proto3" json:"complete_in,omitempty"` // time when user has finished (will finish) participating
	// passcode is a code participant has to enter before she can begin contest
	// this field is only populated when request is made by contest owner
	// passcode is read-only and should be set using ResetPasscode method.
	Passcode string `protobuf:"bytes,30,opt,name=passcode,proto3" json:"passcode,omitempty"`
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

func (x *Participant) GetContestId() string {
	if x != nil {
		return x.ContestId
	}
	return ""
}

func (x *Participant) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Participant) GetUsers() []*Participant_User {
	if x != nil {
		return x.Users
	}
	return nil
}

func (x *Participant) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *Participant) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *Participant) GetStatus() Participant_Status {
	if x != nil {
		return x.Status
	}
	return Participant_NONE
}

func (x *Participant) GetStartedAt() *timestamp.Timestamp {
	if x != nil {
		return x.StartedAt
	}
	return nil
}

func (x *Participant) GetStartedIn() *duration.Duration {
	if x != nil {
		return x.StartedIn
	}
	return nil
}

func (x *Participant) GetCompleteAt() *timestamp.Timestamp {
	if x != nil {
		return x.CompleteAt
	}
	return nil
}

func (x *Participant) GetCompleteIn() *duration.Duration {
	if x != nil {
		return x.CompleteIn
	}
	return nil
}

func (x *Participant) GetPasscode() string {
	if x != nil {
		return x.Passcode
	}
	return ""
}

type Participant_User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *Participant_User) Reset() {
	*x = Participant_User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_judge_participant_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Participant_User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Participant_User) ProtoMessage() {}

func (x *Participant_User) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use Participant_User.ProtoReflect.Descriptor instead.
func (*Participant_User) Descriptor() ([]byte, []int) {
	return file_eolymp_judge_participant_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Participant_User) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

var File_eolymp_judge_participant_proto protoreflect.FileDescriptor

var file_eolymp_judge_participant_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x6a, 0x75, 0x64, 0x67, 0x65, 0x2f, 0x70,
	0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0c, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x6a, 0x75, 0x64, 0x67, 0x65, 0x1a, 0x1f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xfb, 0x04, 0x0a, 0x0b, 0x50, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x1d, 0x0a, 0x0a, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x73, 0x74, 0x49, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x34, 0x0a, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x1e, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x6a, 0x75, 0x64, 0x67, 0x65,
	0x2e, 0x50, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x2e, 0x55, 0x73, 0x65,
	0x72, 0x52, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x8e, 0x4e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x12, 0x1b, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x8f,
	0x4e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x38, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x14, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x20, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x6a, 0x75, 0x64, 0x67, 0x65, 0x2e, 0x50,
	0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x39, 0x0a, 0x0a, 0x73, 0x74, 0x61,
	0x72, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x15, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x12, 0x38, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x65, 0x64, 0x5f,
	0x69, 0x6e, 0x18, 0x16, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x65, 0x64, 0x49, 0x6e, 0x12, 0x3b,
	0x0a, 0x0b, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x5f, 0x61, 0x74, 0x18, 0x19, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x0a, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x74, 0x12, 0x3a, 0x0a, 0x0b, 0x63,
	0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x5f, 0x69, 0x6e, 0x18, 0x1a, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x63, 0x6f, 0x6d,
	0x70, 0x6c, 0x65, 0x74, 0x65, 0x49, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x63,
	0x6f, 0x64, 0x65, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x63,
	0x6f, 0x64, 0x65, 0x1a, 0x1f, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x17, 0x0a, 0x07, 0x75,
	0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x22, 0x57, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x08,
	0x0a, 0x04, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x52, 0x45, 0x41, 0x44,
	0x59, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x41, 0x43, 0x54, 0x49, 0x56, 0x45, 0x10, 0x02, 0x12,
	0x0c, 0x0a, 0x08, 0x43, 0x4f, 0x4d, 0x50, 0x4c, 0x45, 0x54, 0x45, 0x10, 0x03, 0x12, 0x0c, 0x0a,
	0x08, 0x49, 0x4e, 0x41, 0x43, 0x54, 0x49, 0x56, 0x45, 0x10, 0x04, 0x12, 0x10, 0x0a, 0x0c, 0x55,
	0x4e, 0x52, 0x45, 0x47, 0x49, 0x53, 0x54, 0x45, 0x52, 0x45, 0x44, 0x10, 0x05, 0x42, 0x33, 0x5a,
	0x31, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6f, 0x6c, 0x79,
	0x6d, 0x70, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x73, 0x2f, 0x67, 0x6f, 0x2f,
	0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x6a, 0x75, 0x64, 0x67, 0x65, 0x3b, 0x6a, 0x75, 0x64,
	0x67, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
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
	(Participant_Status)(0),     // 0: eolymp.judge.Participant.Status
	(*Participant)(nil),         // 1: eolymp.judge.Participant
	(*Participant_User)(nil),    // 2: eolymp.judge.Participant.User
	(*timestamp.Timestamp)(nil), // 3: google.protobuf.Timestamp
	(*duration.Duration)(nil),   // 4: google.protobuf.Duration
}
var file_eolymp_judge_participant_proto_depIdxs = []int32{
	2, // 0: eolymp.judge.Participant.users:type_name -> eolymp.judge.Participant.User
	0, // 1: eolymp.judge.Participant.status:type_name -> eolymp.judge.Participant.Status
	3, // 2: eolymp.judge.Participant.started_at:type_name -> google.protobuf.Timestamp
	4, // 3: eolymp.judge.Participant.started_in:type_name -> google.protobuf.Duration
	3, // 4: eolymp.judge.Participant.complete_at:type_name -> google.protobuf.Timestamp
	4, // 5: eolymp.judge.Participant.complete_in:type_name -> google.protobuf.Duration
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
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
			switch v := v.(*Participant_User); i {
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
