// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v4.24.4
// source: eolymp/judge/ticket.proto

package judge

import (
	ecm "github.com/eolymp/go-sdk/eolymp/ecm"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
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

type Ticket_Extra int32

const (
	Ticket_NO_EXTRA       Ticket_Extra = 0
	Ticket_MESSAGE_RENDER Ticket_Extra = 1
	Ticket_MESSAGE_VALUE  Ticket_Extra = 2
)

// Enum value maps for Ticket_Extra.
var (
	Ticket_Extra_name = map[int32]string{
		0: "NO_EXTRA",
		1: "MESSAGE_RENDER",
		2: "MESSAGE_VALUE",
	}
	Ticket_Extra_value = map[string]int32{
		"NO_EXTRA":       0,
		"MESSAGE_RENDER": 1,
		"MESSAGE_VALUE":  2,
	}
)

func (x Ticket_Extra) Enum() *Ticket_Extra {
	p := new(Ticket_Extra)
	*p = x
	return p
}

func (x Ticket_Extra) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Ticket_Extra) Descriptor() protoreflect.EnumDescriptor {
	return file_eolymp_judge_ticket_proto_enumTypes[0].Descriptor()
}

func (Ticket_Extra) Type() protoreflect.EnumType {
	return &file_eolymp_judge_ticket_proto_enumTypes[0]
}

func (x Ticket_Extra) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Ticket_Extra.Descriptor instead.
func (Ticket_Extra) EnumDescriptor() ([]byte, []int) {
	return file_eolymp_judge_ticket_proto_rawDescGZIP(), []int{0, 0}
}

type Ticket_Status int32

const (
	Ticket_UNKNOWN_STATUS Ticket_Status = 0
	Ticket_AWAITING       Ticket_Status = 1 // ticket awaiting reply
	Ticket_RESOLVED       Ticket_Status = 2 // ticket with reply
	Ticket_CLOSED         Ticket_Status = 3 // closed ticket
)

// Enum value maps for Ticket_Status.
var (
	Ticket_Status_name = map[int32]string{
		0: "UNKNOWN_STATUS",
		1: "AWAITING",
		2: "RESOLVED",
		3: "CLOSED",
	}
	Ticket_Status_value = map[string]int32{
		"UNKNOWN_STATUS": 0,
		"AWAITING":       1,
		"RESOLVED":       2,
		"CLOSED":         3,
	}
)

func (x Ticket_Status) Enum() *Ticket_Status {
	p := new(Ticket_Status)
	*p = x
	return p
}

func (x Ticket_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Ticket_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_eolymp_judge_ticket_proto_enumTypes[1].Descriptor()
}

func (Ticket_Status) Type() protoreflect.EnumType {
	return &file_eolymp_judge_ticket_proto_enumTypes[1]
}

func (x Ticket_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Ticket_Status.Descriptor instead.
func (Ticket_Status) EnumDescriptor() ([]byte, []int) {
	return file_eolymp_judge_ticket_proto_rawDescGZIP(), []int{0, 1}
}

type Ticket struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Ticket unique identifier.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Contest where ticket was opened.
	ContestId string `protobuf:"bytes,2,opt,name=contest_id,json=contestId,proto3" json:"contest_id,omitempty"`
	// Participant who opened the ticket.
	ParticipantId string `protobuf:"bytes,3,opt,name=participant_id,json=participantId,proto3" json:"participant_id,omitempty"`
	// Member who opened the ticket.
	MemberId string `protobuf:"bytes,4,opt,name=member_id,json=memberId,proto3" json:"member_id,omitempty"`
	// Ticket status as
	Status Ticket_Status `protobuf:"varint,5,opt,name=status,proto3,enum=eolymp.judge.Ticket_Status" json:"status,omitempty"`
	// Ticket subject. Max 255 symbols.
	Subject string `protobuf:"bytes,10,opt,name=subject,proto3" json:"subject,omitempty"`
	// Ticket message.
	Message    *ecm.Content `protobuf:"bytes,12,opt,name=message,proto3" json:"message,omitempty"`
	RawMessage string       `protobuf:"bytes,11,opt,name=raw_message,json=rawMessage,proto3" json:"raw_message,omitempty"`
	// True if ticket has not been read by current user
	IsRead bool `protobuf:"varint,21,opt,name=is_read,json=isRead,proto3" json:"is_read,omitempty"`
	// Number of replies in the ticket.
	ReplyCount uint32 `protobuf:"varint,30,opt,name=reply_count,json=replyCount,proto3" json:"reply_count,omitempty"`
	// Timestamp when ticket was initially created.
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,23,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	// Timestamp when ticket was modified/replied.
	UpdatedAt *timestamppb.Timestamp `protobuf:"bytes,24,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	// Timestamp when ticket was read by current user
	ReadAt *timestamppb.Timestamp `protobuf:"bytes,25,opt,name=read_at,json=readAt,proto3" json:"read_at,omitempty"`
	// Timestamp when last reply was added
	LastReplyAt *timestamppb.Timestamp `protobuf:"bytes,26,opt,name=last_reply_at,json=lastReplyAt,proto3" json:"last_reply_at,omitempty"`
	// Cursor when ticket is queries in the list
	Cursor        string `protobuf:"bytes,100,opt,name=cursor,proto3" json:"cursor,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Ticket) Reset() {
	*x = Ticket{}
	mi := &file_eolymp_judge_ticket_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Ticket) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ticket) ProtoMessage() {}

func (x *Ticket) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_judge_ticket_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Ticket.ProtoReflect.Descriptor instead.
func (*Ticket) Descriptor() ([]byte, []int) {
	return file_eolymp_judge_ticket_proto_rawDescGZIP(), []int{0}
}

func (x *Ticket) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Ticket) GetContestId() string {
	if x != nil {
		return x.ContestId
	}
	return ""
}

func (x *Ticket) GetParticipantId() string {
	if x != nil {
		return x.ParticipantId
	}
	return ""
}

func (x *Ticket) GetMemberId() string {
	if x != nil {
		return x.MemberId
	}
	return ""
}

func (x *Ticket) GetStatus() Ticket_Status {
	if x != nil {
		return x.Status
	}
	return Ticket_UNKNOWN_STATUS
}

func (x *Ticket) GetSubject() string {
	if x != nil {
		return x.Subject
	}
	return ""
}

func (x *Ticket) GetMessage() *ecm.Content {
	if x != nil {
		return x.Message
	}
	return nil
}

func (x *Ticket) GetRawMessage() string {
	if x != nil {
		return x.RawMessage
	}
	return ""
}

func (x *Ticket) GetIsRead() bool {
	if x != nil {
		return x.IsRead
	}
	return false
}

func (x *Ticket) GetReplyCount() uint32 {
	if x != nil {
		return x.ReplyCount
	}
	return 0
}

func (x *Ticket) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Ticket) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *Ticket) GetReadAt() *timestamppb.Timestamp {
	if x != nil {
		return x.ReadAt
	}
	return nil
}

func (x *Ticket) GetLastReplyAt() *timestamppb.Timestamp {
	if x != nil {
		return x.LastReplyAt
	}
	return nil
}

func (x *Ticket) GetCursor() string {
	if x != nil {
		return x.Cursor
	}
	return ""
}

var File_eolymp_judge_ticket_proto protoreflect.FileDescriptor

const file_eolymp_judge_ticket_proto_rawDesc = "" +
	"\n" +
	"\x19eolymp/judge/ticket.proto\x12\feolymp.judge\x1a\x18eolymp/ecm/content.proto\x1a\x1fgoogle/protobuf/timestamp.proto\"\xdb\x05\n" +
	"\x06Ticket\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x12\x1d\n" +
	"\n" +
	"contest_id\x18\x02 \x01(\tR\tcontestId\x12%\n" +
	"\x0eparticipant_id\x18\x03 \x01(\tR\rparticipantId\x12\x1b\n" +
	"\tmember_id\x18\x04 \x01(\tR\bmemberId\x123\n" +
	"\x06status\x18\x05 \x01(\x0e2\x1b.eolymp.judge.Ticket.StatusR\x06status\x12\x18\n" +
	"\asubject\x18\n" +
	" \x01(\tR\asubject\x12-\n" +
	"\amessage\x18\f \x01(\v2\x13.eolymp.ecm.ContentR\amessage\x12\x1f\n" +
	"\vraw_message\x18\v \x01(\tR\n" +
	"rawMessage\x12\x17\n" +
	"\ais_read\x18\x15 \x01(\bR\x06isRead\x12\x1f\n" +
	"\vreply_count\x18\x1e \x01(\rR\n" +
	"replyCount\x129\n" +
	"\n" +
	"created_at\x18\x17 \x01(\v2\x1a.google.protobuf.TimestampR\tcreatedAt\x129\n" +
	"\n" +
	"updated_at\x18\x18 \x01(\v2\x1a.google.protobuf.TimestampR\tupdatedAt\x123\n" +
	"\aread_at\x18\x19 \x01(\v2\x1a.google.protobuf.TimestampR\x06readAt\x12>\n" +
	"\rlast_reply_at\x18\x1a \x01(\v2\x1a.google.protobuf.TimestampR\vlastReplyAt\x12\x16\n" +
	"\x06cursor\x18d \x01(\tR\x06cursor\"<\n" +
	"\x05Extra\x12\f\n" +
	"\bNO_EXTRA\x10\x00\x12\x12\n" +
	"\x0eMESSAGE_RENDER\x10\x01\x12\x11\n" +
	"\rMESSAGE_VALUE\x10\x02\"D\n" +
	"\x06Status\x12\x12\n" +
	"\x0eUNKNOWN_STATUS\x10\x00\x12\f\n" +
	"\bAWAITING\x10\x01\x12\f\n" +
	"\bRESOLVED\x10\x02\x12\n" +
	"\n" +
	"\x06CLOSED\x10\x03B-Z+github.com/eolymp/go-sdk/eolymp/judge;judgeb\x06proto3"

var (
	file_eolymp_judge_ticket_proto_rawDescOnce sync.Once
	file_eolymp_judge_ticket_proto_rawDescData []byte
)

func file_eolymp_judge_ticket_proto_rawDescGZIP() []byte {
	file_eolymp_judge_ticket_proto_rawDescOnce.Do(func() {
		file_eolymp_judge_ticket_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_eolymp_judge_ticket_proto_rawDesc), len(file_eolymp_judge_ticket_proto_rawDesc)))
	})
	return file_eolymp_judge_ticket_proto_rawDescData
}

var file_eolymp_judge_ticket_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_eolymp_judge_ticket_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_eolymp_judge_ticket_proto_goTypes = []any{
	(Ticket_Extra)(0),             // 0: eolymp.judge.Ticket.Extra
	(Ticket_Status)(0),            // 1: eolymp.judge.Ticket.Status
	(*Ticket)(nil),                // 2: eolymp.judge.Ticket
	(*ecm.Content)(nil),           // 3: eolymp.ecm.Content
	(*timestamppb.Timestamp)(nil), // 4: google.protobuf.Timestamp
}
var file_eolymp_judge_ticket_proto_depIdxs = []int32{
	1, // 0: eolymp.judge.Ticket.status:type_name -> eolymp.judge.Ticket.Status
	3, // 1: eolymp.judge.Ticket.message:type_name -> eolymp.ecm.Content
	4, // 2: eolymp.judge.Ticket.created_at:type_name -> google.protobuf.Timestamp
	4, // 3: eolymp.judge.Ticket.updated_at:type_name -> google.protobuf.Timestamp
	4, // 4: eolymp.judge.Ticket.read_at:type_name -> google.protobuf.Timestamp
	4, // 5: eolymp.judge.Ticket.last_reply_at:type_name -> google.protobuf.Timestamp
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_eolymp_judge_ticket_proto_init() }
func file_eolymp_judge_ticket_proto_init() {
	if File_eolymp_judge_ticket_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_eolymp_judge_ticket_proto_rawDesc), len(file_eolymp_judge_ticket_proto_rawDesc)),
			NumEnums:      2,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_eolymp_judge_ticket_proto_goTypes,
		DependencyIndexes: file_eolymp_judge_ticket_proto_depIdxs,
		EnumInfos:         file_eolymp_judge_ticket_proto_enumTypes,
		MessageInfos:      file_eolymp_judge_ticket_proto_msgTypes,
	}.Build()
	File_eolymp_judge_ticket_proto = out.File
	file_eolymp_judge_ticket_proto_goTypes = nil
	file_eolymp_judge_ticket_proto_depIdxs = nil
}
