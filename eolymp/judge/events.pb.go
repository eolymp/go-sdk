// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.18.1
// source: eolymp/judge/events.proto

package judge

import (
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

type SubmissionCompletedEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Submission *Submission `protobuf:"bytes,1,opt,name=submission,proto3" json:"submission,omitempty"`
}

func (x *SubmissionCompletedEvent) Reset() {
	*x = SubmissionCompletedEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_judge_events_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubmissionCompletedEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubmissionCompletedEvent) ProtoMessage() {}

func (x *SubmissionCompletedEvent) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_judge_events_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubmissionCompletedEvent.ProtoReflect.Descriptor instead.
func (*SubmissionCompletedEvent) Descriptor() ([]byte, []int) {
	return file_eolymp_judge_events_proto_rawDescGZIP(), []int{0}
}

func (x *SubmissionCompletedEvent) GetSubmission() *Submission {
	if x != nil {
		return x.Submission
	}
	return nil
}

type RebuildScoreEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ContestId  string `protobuf:"bytes,1,opt,name=contest_id,json=contestId,proto3" json:"contest_id,omitempty"`
	ActivityId string `protobuf:"bytes,2,opt,name=activity_id,json=activityId,proto3" json:"activity_id,omitempty"`
}

func (x *RebuildScoreEvent) Reset() {
	*x = RebuildScoreEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_judge_events_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RebuildScoreEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RebuildScoreEvent) ProtoMessage() {}

func (x *RebuildScoreEvent) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_judge_events_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RebuildScoreEvent.ProtoReflect.Descriptor instead.
func (*RebuildScoreEvent) Descriptor() ([]byte, []int) {
	return file_eolymp_judge_events_proto_rawDescGZIP(), []int{1}
}

func (x *RebuildScoreEvent) GetContestId() string {
	if x != nil {
		return x.ContestId
	}
	return ""
}

func (x *RebuildScoreEvent) GetActivityId() string {
	if x != nil {
		return x.ActivityId
	}
	return ""
}

type ScoreUpdatedEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ContestId        string `protobuf:"bytes,1,opt,name=contest_id,json=contestId,proto3" json:"contest_id,omitempty"`
	ParticipantId    string `protobuf:"bytes,2,opt,name=participant_id,json=participantId,proto3" json:"participant_id,omitempty"`
	OutOfCompetition bool   `protobuf:"varint,4,opt,name=out_of_competition,json=outOfCompetition,proto3" json:"out_of_competition,omitempty"`
	Score            *Score `protobuf:"bytes,3,opt,name=score,proto3" json:"score,omitempty"`
}

func (x *ScoreUpdatedEvent) Reset() {
	*x = ScoreUpdatedEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_judge_events_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ScoreUpdatedEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ScoreUpdatedEvent) ProtoMessage() {}

func (x *ScoreUpdatedEvent) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_judge_events_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ScoreUpdatedEvent.ProtoReflect.Descriptor instead.
func (*ScoreUpdatedEvent) Descriptor() ([]byte, []int) {
	return file_eolymp_judge_events_proto_rawDescGZIP(), []int{2}
}

func (x *ScoreUpdatedEvent) GetContestId() string {
	if x != nil {
		return x.ContestId
	}
	return ""
}

func (x *ScoreUpdatedEvent) GetParticipantId() string {
	if x != nil {
		return x.ParticipantId
	}
	return ""
}

func (x *ScoreUpdatedEvent) GetOutOfCompetition() bool {
	if x != nil {
		return x.OutOfCompetition
	}
	return false
}

func (x *ScoreUpdatedEvent) GetScore() *Score {
	if x != nil {
		return x.Score
	}
	return nil
}

type TicketCreatedEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ticket *Ticket `protobuf:"bytes,1,opt,name=ticket,proto3" json:"ticket,omitempty"`
}

func (x *TicketCreatedEvent) Reset() {
	*x = TicketCreatedEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_judge_events_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TicketCreatedEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TicketCreatedEvent) ProtoMessage() {}

func (x *TicketCreatedEvent) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_judge_events_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TicketCreatedEvent.ProtoReflect.Descriptor instead.
func (*TicketCreatedEvent) Descriptor() ([]byte, []int) {
	return file_eolymp_judge_events_proto_rawDescGZIP(), []int{3}
}

func (x *TicketCreatedEvent) GetTicket() *Ticket {
	if x != nil {
		return x.Ticket
	}
	return nil
}

type TicketUpdatedEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ticket *Ticket `protobuf:"bytes,1,opt,name=ticket,proto3" json:"ticket,omitempty"`
	Reply  *Reply  `protobuf:"bytes,2,opt,name=reply,proto3" json:"reply,omitempty"`
}

func (x *TicketUpdatedEvent) Reset() {
	*x = TicketUpdatedEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_judge_events_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TicketUpdatedEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TicketUpdatedEvent) ProtoMessage() {}

func (x *TicketUpdatedEvent) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_judge_events_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TicketUpdatedEvent.ProtoReflect.Descriptor instead.
func (*TicketUpdatedEvent) Descriptor() ([]byte, []int) {
	return file_eolymp_judge_events_proto_rawDescGZIP(), []int{4}
}

func (x *TicketUpdatedEvent) GetTicket() *Ticket {
	if x != nil {
		return x.Ticket
	}
	return nil
}

func (x *TicketUpdatedEvent) GetReply() *Reply {
	if x != nil {
		return x.Reply
	}
	return nil
}

var File_eolymp_judge_events_proto protoreflect.FileDescriptor

var file_eolymp_judge_events_proto_rawDesc = []byte{
	0x0a, 0x19, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x6a, 0x75, 0x64, 0x67, 0x65, 0x2f, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x65, 0x6f, 0x6c,
	0x79, 0x6d, 0x70, 0x2e, 0x6a, 0x75, 0x64, 0x67, 0x65, 0x1a, 0x18, 0x65, 0x6f, 0x6c, 0x79, 0x6d,
	0x70, 0x2f, 0x6a, 0x75, 0x64, 0x67, 0x65, 0x2f, 0x72, 0x65, 0x70, 0x6c, 0x79, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x18, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x6a, 0x75, 0x64, 0x67,
	0x65, 0x2f, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x65,
	0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x6a, 0x75, 0x64, 0x67, 0x65, 0x2f, 0x73, 0x75, 0x62, 0x6d,
	0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x65, 0x6f,
	0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x6a, 0x75, 0x64, 0x67, 0x65, 0x2f, 0x74, 0x69, 0x63, 0x6b, 0x65,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x54, 0x0a, 0x18, 0x53, 0x75, 0x62, 0x6d, 0x69,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x12, 0x38, 0x0a, 0x0a, 0x73, 0x75, 0x62, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70,
	0x2e, 0x6a, 0x75, 0x64, 0x67, 0x65, 0x2e, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x52, 0x0a, 0x73, 0x75, 0x62, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x53, 0x0a,
	0x11, 0x52, 0x65, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x73, 0x74, 0x49,
	0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x5f, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79,
	0x49, 0x64, 0x22, 0xb2, 0x01, 0x0a, 0x11, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x64, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6f, 0x6e, 0x74,
	0x65, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x73, 0x74, 0x49, 0x64, 0x12, 0x25, 0x0a, 0x0e, 0x70, 0x61, 0x72, 0x74, 0x69,
	0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0d, 0x70, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x2c,
	0x0a, 0x12, 0x6f, 0x75, 0x74, 0x5f, 0x6f, 0x66, 0x5f, 0x63, 0x6f, 0x6d, 0x70, 0x65, 0x74, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x10, 0x6f, 0x75, 0x74, 0x4f,
	0x66, 0x43, 0x6f, 0x6d, 0x70, 0x65, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x29, 0x0a, 0x05,
	0x73, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x65, 0x6f,
	0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x6a, 0x75, 0x64, 0x67, 0x65, 0x2e, 0x53, 0x63, 0x6f, 0x72, 0x65,
	0x52, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x22, 0x42, 0x0a, 0x12, 0x54, 0x69, 0x63, 0x6b, 0x65,
	0x74, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x2c, 0x0a,
	0x06, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e,
	0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x6a, 0x75, 0x64, 0x67, 0x65, 0x2e, 0x54, 0x69, 0x63,
	0x6b, 0x65, 0x74, 0x52, 0x06, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x22, 0x6d, 0x0a, 0x12, 0x54,
	0x69, 0x63, 0x6b, 0x65, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x12, 0x2c, 0x0a, 0x06, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x14, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x6a, 0x75, 0x64, 0x67, 0x65,
	0x2e, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x06, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x12,
	0x29, 0x0a, 0x05, 0x72, 0x65, 0x70, 0x6c, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13,
	0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x6a, 0x75, 0x64, 0x67, 0x65, 0x2e, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x52, 0x05, 0x72, 0x65, 0x70, 0x6c, 0x79, 0x42, 0x2d, 0x5a, 0x2b, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f,
	0x67, 0x6f, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x6a, 0x75,
	0x64, 0x67, 0x65, 0x3b, 0x6a, 0x75, 0x64, 0x67, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_eolymp_judge_events_proto_rawDescOnce sync.Once
	file_eolymp_judge_events_proto_rawDescData = file_eolymp_judge_events_proto_rawDesc
)

func file_eolymp_judge_events_proto_rawDescGZIP() []byte {
	file_eolymp_judge_events_proto_rawDescOnce.Do(func() {
		file_eolymp_judge_events_proto_rawDescData = protoimpl.X.CompressGZIP(file_eolymp_judge_events_proto_rawDescData)
	})
	return file_eolymp_judge_events_proto_rawDescData
}

var file_eolymp_judge_events_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_eolymp_judge_events_proto_goTypes = []interface{}{
	(*SubmissionCompletedEvent)(nil), // 0: eolymp.judge.SubmissionCompletedEvent
	(*RebuildScoreEvent)(nil),        // 1: eolymp.judge.RebuildScoreEvent
	(*ScoreUpdatedEvent)(nil),        // 2: eolymp.judge.ScoreUpdatedEvent
	(*TicketCreatedEvent)(nil),       // 3: eolymp.judge.TicketCreatedEvent
	(*TicketUpdatedEvent)(nil),       // 4: eolymp.judge.TicketUpdatedEvent
	(*Submission)(nil),               // 5: eolymp.judge.Submission
	(*Score)(nil),                    // 6: eolymp.judge.Score
	(*Ticket)(nil),                   // 7: eolymp.judge.Ticket
	(*Reply)(nil),                    // 8: eolymp.judge.Reply
}
var file_eolymp_judge_events_proto_depIdxs = []int32{
	5, // 0: eolymp.judge.SubmissionCompletedEvent.submission:type_name -> eolymp.judge.Submission
	6, // 1: eolymp.judge.ScoreUpdatedEvent.score:type_name -> eolymp.judge.Score
	7, // 2: eolymp.judge.TicketCreatedEvent.ticket:type_name -> eolymp.judge.Ticket
	7, // 3: eolymp.judge.TicketUpdatedEvent.ticket:type_name -> eolymp.judge.Ticket
	8, // 4: eolymp.judge.TicketUpdatedEvent.reply:type_name -> eolymp.judge.Reply
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_eolymp_judge_events_proto_init() }
func file_eolymp_judge_events_proto_init() {
	if File_eolymp_judge_events_proto != nil {
		return
	}
	file_eolymp_judge_reply_proto_init()
	file_eolymp_judge_score_proto_init()
	file_eolymp_judge_submission_proto_init()
	file_eolymp_judge_ticket_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_eolymp_judge_events_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubmissionCompletedEvent); i {
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
		file_eolymp_judge_events_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RebuildScoreEvent); i {
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
		file_eolymp_judge_events_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ScoreUpdatedEvent); i {
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
		file_eolymp_judge_events_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TicketCreatedEvent); i {
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
		file_eolymp_judge_events_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TicketUpdatedEvent); i {
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
			RawDescriptor: file_eolymp_judge_events_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_eolymp_judge_events_proto_goTypes,
		DependencyIndexes: file_eolymp_judge_events_proto_depIdxs,
		MessageInfos:      file_eolymp_judge_events_proto_msgTypes,
	}.Build()
	File_eolymp_judge_events_proto = out.File
	file_eolymp_judge_events_proto_rawDesc = nil
	file_eolymp_judge_events_proto_goTypes = nil
	file_eolymp_judge_events_proto_depIdxs = nil
}
