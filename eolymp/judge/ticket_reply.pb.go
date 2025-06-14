// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.29.4
// source: eolymp/judge/ticket_reply.proto

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

type Reply_Extra int32

const (
	Reply_NO_EXTRA       Reply_Extra = 0
	Reply_MESSAGE_RENDER Reply_Extra = 1
	Reply_MESSAGE_VALUE  Reply_Extra = 2
)

// Enum value maps for Reply_Extra.
var (
	Reply_Extra_name = map[int32]string{
		0: "NO_EXTRA",
		1: "MESSAGE_RENDER",
		2: "MESSAGE_VALUE",
	}
	Reply_Extra_value = map[string]int32{
		"NO_EXTRA":       0,
		"MESSAGE_RENDER": 1,
		"MESSAGE_VALUE":  2,
	}
)

func (x Reply_Extra) Enum() *Reply_Extra {
	p := new(Reply_Extra)
	*p = x
	return p
}

func (x Reply_Extra) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Reply_Extra) Descriptor() protoreflect.EnumDescriptor {
	return file_eolymp_judge_ticket_reply_proto_enumTypes[0].Descriptor()
}

func (Reply_Extra) Type() protoreflect.EnumType {
	return &file_eolymp_judge_ticket_reply_proto_enumTypes[0]
}

func (x Reply_Extra) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Reply_Extra.Descriptor instead.
func (Reply_Extra) EnumDescriptor() ([]byte, []int) {
	return file_eolymp_judge_ticket_reply_proto_rawDescGZIP(), []int{0, 0}
}

type Reply struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Reply unique identifier.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Ticket this reply belongs to.
	TicketId string `protobuf:"bytes,2,opt,name=ticket_id,json=ticketId,proto3" json:"ticket_id,omitempty"`
	// User (jury) or Member (participant) who added this reply.
	//
	// Types that are valid to be assigned to Author:
	//
	//	*Reply_UserId
	//	*Reply_MemberId
	Author isReply_Author `protobuf_oneof:"author"`
	// Message of the reply.
	Message *ecm.Content `protobuf:"bytes,10,opt,name=message,proto3" json:"message,omitempty"`
	// Timestamp when reply was created.
	CreatedAt     *timestamppb.Timestamp `protobuf:"bytes,11,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Reply) Reset() {
	*x = Reply{}
	mi := &file_eolymp_judge_ticket_reply_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Reply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Reply) ProtoMessage() {}

func (x *Reply) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_judge_ticket_reply_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Reply.ProtoReflect.Descriptor instead.
func (*Reply) Descriptor() ([]byte, []int) {
	return file_eolymp_judge_ticket_reply_proto_rawDescGZIP(), []int{0}
}

func (x *Reply) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Reply) GetTicketId() string {
	if x != nil {
		return x.TicketId
	}
	return ""
}

func (x *Reply) GetAuthor() isReply_Author {
	if x != nil {
		return x.Author
	}
	return nil
}

func (x *Reply) GetUserId() string {
	if x != nil {
		if x, ok := x.Author.(*Reply_UserId); ok {
			return x.UserId
		}
	}
	return ""
}

func (x *Reply) GetMemberId() string {
	if x != nil {
		if x, ok := x.Author.(*Reply_MemberId); ok {
			return x.MemberId
		}
	}
	return ""
}

func (x *Reply) GetMessage() *ecm.Content {
	if x != nil {
		return x.Message
	}
	return nil
}

func (x *Reply) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

type isReply_Author interface {
	isReply_Author()
}

type Reply_UserId struct {
	UserId string `protobuf:"bytes,3,opt,name=user_id,json=userId,proto3,oneof"`
}

type Reply_MemberId struct {
	MemberId string `protobuf:"bytes,4,opt,name=member_id,json=memberId,proto3,oneof"`
}

func (*Reply_UserId) isReply_Author() {}

func (*Reply_MemberId) isReply_Author() {}

var File_eolymp_judge_ticket_reply_proto protoreflect.FileDescriptor

const file_eolymp_judge_ticket_reply_proto_rawDesc = "" +
	"\n" +
	"\x1feolymp/judge/ticket_reply.proto\x12\feolymp.judge\x1a\x18eolymp/ecm/content.proto\x1a\x1fgoogle/protobuf/timestamp.proto\"\xa0\x02\n" +
	"\x05Reply\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x12\x1b\n" +
	"\tticket_id\x18\x02 \x01(\tR\bticketId\x12\x19\n" +
	"\auser_id\x18\x03 \x01(\tH\x00R\x06userId\x12\x1d\n" +
	"\tmember_id\x18\x04 \x01(\tH\x00R\bmemberId\x12-\n" +
	"\amessage\x18\n" +
	" \x01(\v2\x13.eolymp.ecm.ContentR\amessage\x129\n" +
	"\n" +
	"created_at\x18\v \x01(\v2\x1a.google.protobuf.TimestampR\tcreatedAt\"<\n" +
	"\x05Extra\x12\f\n" +
	"\bNO_EXTRA\x10\x00\x12\x12\n" +
	"\x0eMESSAGE_RENDER\x10\x01\x12\x11\n" +
	"\rMESSAGE_VALUE\x10\x02B\b\n" +
	"\x06authorB-Z+github.com/eolymp/go-sdk/eolymp/judge;judgeb\x06proto3"

var (
	file_eolymp_judge_ticket_reply_proto_rawDescOnce sync.Once
	file_eolymp_judge_ticket_reply_proto_rawDescData []byte
)

func file_eolymp_judge_ticket_reply_proto_rawDescGZIP() []byte {
	file_eolymp_judge_ticket_reply_proto_rawDescOnce.Do(func() {
		file_eolymp_judge_ticket_reply_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_eolymp_judge_ticket_reply_proto_rawDesc), len(file_eolymp_judge_ticket_reply_proto_rawDesc)))
	})
	return file_eolymp_judge_ticket_reply_proto_rawDescData
}

var file_eolymp_judge_ticket_reply_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_eolymp_judge_ticket_reply_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_eolymp_judge_ticket_reply_proto_goTypes = []any{
	(Reply_Extra)(0),              // 0: eolymp.judge.Reply.Extra
	(*Reply)(nil),                 // 1: eolymp.judge.Reply
	(*ecm.Content)(nil),           // 2: eolymp.ecm.Content
	(*timestamppb.Timestamp)(nil), // 3: google.protobuf.Timestamp
}
var file_eolymp_judge_ticket_reply_proto_depIdxs = []int32{
	2, // 0: eolymp.judge.Reply.message:type_name -> eolymp.ecm.Content
	3, // 1: eolymp.judge.Reply.created_at:type_name -> google.protobuf.Timestamp
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_eolymp_judge_ticket_reply_proto_init() }
func file_eolymp_judge_ticket_reply_proto_init() {
	if File_eolymp_judge_ticket_reply_proto != nil {
		return
	}
	file_eolymp_judge_ticket_reply_proto_msgTypes[0].OneofWrappers = []any{
		(*Reply_UserId)(nil),
		(*Reply_MemberId)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_eolymp_judge_ticket_reply_proto_rawDesc), len(file_eolymp_judge_ticket_reply_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_eolymp_judge_ticket_reply_proto_goTypes,
		DependencyIndexes: file_eolymp_judge_ticket_reply_proto_depIdxs,
		EnumInfos:         file_eolymp_judge_ticket_reply_proto_enumTypes,
		MessageInfos:      file_eolymp_judge_ticket_reply_proto_msgTypes,
	}.Build()
	File_eolymp_judge_ticket_reply_proto = out.File
	file_eolymp_judge_ticket_reply_proto_goTypes = nil
	file_eolymp_judge_ticket_reply_proto_depIdxs = nil
}
