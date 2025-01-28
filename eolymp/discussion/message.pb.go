// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.4
// 	protoc        v4.24.4
// source: eolymp/discussion/message.proto

package discussion

import (
	ecm "github.com/eolymp/go-sdk/eolymp/ecm"
	wellknown "github.com/eolymp/go-sdk/eolymp/wellknown"
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

type Message_Extra int32

const (
	Message_UNKNOWN_EXTRA   Message_Extra = 0
	Message_MESSAGE_VALUE   Message_Extra = 1 // include message source (original Markdown, LaTeX etc)
	Message_MESSAGE_RENDER  Message_Extra = 2 // include rendered message in ECM
	Message_MESSAGE_PREVIEW Message_Extra = 3 // include rendered but trimmed message in ECM, overrides MESSAGE_RENDER
	Message_VOTE            Message_Extra = 4 // include vote field
)

// Enum value maps for Message_Extra.
var (
	Message_Extra_name = map[int32]string{
		0: "UNKNOWN_EXTRA",
		1: "MESSAGE_VALUE",
		2: "MESSAGE_RENDER",
		3: "MESSAGE_PREVIEW",
		4: "VOTE",
	}
	Message_Extra_value = map[string]int32{
		"UNKNOWN_EXTRA":   0,
		"MESSAGE_VALUE":   1,
		"MESSAGE_RENDER":  2,
		"MESSAGE_PREVIEW": 3,
		"VOTE":            4,
	}
)

func (x Message_Extra) Enum() *Message_Extra {
	p := new(Message_Extra)
	*p = x
	return p
}

func (x Message_Extra) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Message_Extra) Descriptor() protoreflect.EnumDescriptor {
	return file_eolymp_discussion_message_proto_enumTypes[0].Descriptor()
}

func (Message_Extra) Type() protoreflect.EnumType {
	return &file_eolymp_discussion_message_proto_enumTypes[0]
}

func (x Message_Extra) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Message_Extra.Descriptor instead.
func (Message_Extra) EnumDescriptor() ([]byte, []int) {
	return file_eolymp_discussion_message_proto_rawDescGZIP(), []int{0, 0}
}

type Message struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Url           string                 `protobuf:"bytes,5,opt,name=url,proto3" json:"url,omitempty"`
	ThreadUrl     string                 `protobuf:"bytes,2,opt,name=thread_url,json=threadUrl,proto3" json:"thread_url,omitempty"`
	MemberId      string                 `protobuf:"bytes,3,opt,name=member_id,json=memberId,proto3" json:"member_id,omitempty"`
	ReplyTo       string                 `protobuf:"bytes,4,opt,name=reply_to,json=replyTo,proto3" json:"reply_to,omitempty"`
	Vote          int32                  `protobuf:"varint,12,opt,name=vote,proto3" json:"vote,omitempty"`                               // vote of authenticated user (+1 or -1)
	VoteCount     int32                  `protobuf:"varint,10,opt,name=vote_count,json=voteCount,proto3" json:"vote_count,omitempty"`    // total vote count
	ReplyCount    int32                  `protobuf:"varint,11,opt,name=reply_count,json=replyCount,proto3" json:"reply_count,omitempty"` // total number of replies
	PostedAt      *timestamppb.Timestamp `protobuf:"bytes,20,opt,name=posted_at,json=postedAt,proto3" json:"posted_at,omitempty"`
	EditedAt      *timestamppb.Timestamp `protobuf:"bytes,21,opt,name=edited_at,json=editedAt,proto3" json:"edited_at,omitempty"`
	DeletedAt     *timestamppb.Timestamp `protobuf:"bytes,22,opt,name=deleted_at,json=deletedAt,proto3" json:"deleted_at,omitempty"`
	Revision      int32                  `protobuf:"varint,13,opt,name=revision,proto3" json:"revision,omitempty"`
	Message       *ecm.Content           `protobuf:"bytes,100,opt,name=message,proto3" json:"message,omitempty"`
	Links         []*wellknown.Link      `protobuf:"bytes,200,rep,name=links,proto3" json:"links,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Message) Reset() {
	*x = Message{}
	mi := &file_eolymp_discussion_message_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message) ProtoMessage() {}

func (x *Message) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_discussion_message_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Message.ProtoReflect.Descriptor instead.
func (*Message) Descriptor() ([]byte, []int) {
	return file_eolymp_discussion_message_proto_rawDescGZIP(), []int{0}
}

func (x *Message) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Message) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *Message) GetThreadUrl() string {
	if x != nil {
		return x.ThreadUrl
	}
	return ""
}

func (x *Message) GetMemberId() string {
	if x != nil {
		return x.MemberId
	}
	return ""
}

func (x *Message) GetReplyTo() string {
	if x != nil {
		return x.ReplyTo
	}
	return ""
}

func (x *Message) GetVote() int32 {
	if x != nil {
		return x.Vote
	}
	return 0
}

func (x *Message) GetVoteCount() int32 {
	if x != nil {
		return x.VoteCount
	}
	return 0
}

func (x *Message) GetReplyCount() int32 {
	if x != nil {
		return x.ReplyCount
	}
	return 0
}

func (x *Message) GetPostedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.PostedAt
	}
	return nil
}

func (x *Message) GetEditedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.EditedAt
	}
	return nil
}

func (x *Message) GetDeletedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.DeletedAt
	}
	return nil
}

func (x *Message) GetRevision() int32 {
	if x != nil {
		return x.Revision
	}
	return 0
}

func (x *Message) GetMessage() *ecm.Content {
	if x != nil {
		return x.Message
	}
	return nil
}

func (x *Message) GetLinks() []*wellknown.Link {
	if x != nil {
		return x.Links
	}
	return nil
}

var File_eolymp_discussion_message_proto protoreflect.FileDescriptor

var file_eolymp_discussion_message_proto_rawDesc = string([]byte{
	0x0a, 0x1f, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x64, 0x69, 0x73, 0x63, 0x75, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x11, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x75, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x1a, 0x18, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x65, 0x63, 0x6d,
	0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b,
	0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x77, 0x65, 0x6c, 0x6c, 0x6b, 0x6e, 0x6f, 0x77, 0x6e,
	0x2f, 0x6c, 0x69, 0x6e, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xdf, 0x04, 0x0a,
	0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x68,
	0x72, 0x65, 0x61, 0x64, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x74, 0x68, 0x72, 0x65, 0x61, 0x64, 0x55, 0x72, 0x6c, 0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x65, 0x6d,
	0x62, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6d, 0x65,
	0x6d, 0x62, 0x65, 0x72, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x72, 0x65, 0x70, 0x6c, 0x79, 0x5f,
	0x74, 0x6f, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x72, 0x65, 0x70, 0x6c, 0x79, 0x54,
	0x6f, 0x12, 0x12, 0x0a, 0x04, 0x76, 0x6f, 0x74, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x04, 0x76, 0x6f, 0x74, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x76, 0x6f, 0x74, 0x65, 0x5f, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x76, 0x6f, 0x74, 0x65, 0x43,
	0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x70, 0x6c, 0x79, 0x5f, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x72, 0x65, 0x70, 0x6c, 0x79,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x37, 0x0a, 0x09, 0x70, 0x6f, 0x73, 0x74, 0x65, 0x64, 0x5f,
	0x61, 0x74, 0x18, 0x14, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x08, 0x70, 0x6f, 0x73, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x37,
	0x0a, 0x09, 0x65, 0x64, 0x69, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x15, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x08, 0x65,
	0x64, 0x69, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x64, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x16, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x0d,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x72, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x2d,
	0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x64, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x13, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x65, 0x63, 0x6d, 0x2e, 0x43, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x2d, 0x0a,
	0x05, 0x6c, 0x69, 0x6e, 0x6b, 0x73, 0x18, 0xc8, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e,
	0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x77, 0x65, 0x6c, 0x6c, 0x6b, 0x6e, 0x6f, 0x77, 0x6e,
	0x2e, 0x4c, 0x69, 0x6e, 0x6b, 0x52, 0x05, 0x6c, 0x69, 0x6e, 0x6b, 0x73, 0x22, 0x60, 0x0a, 0x05,
	0x45, 0x78, 0x74, 0x72, 0x61, 0x12, 0x11, 0x0a, 0x0d, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e,
	0x5f, 0x45, 0x58, 0x54, 0x52, 0x41, 0x10, 0x00, 0x12, 0x11, 0x0a, 0x0d, 0x4d, 0x45, 0x53, 0x53,
	0x41, 0x47, 0x45, 0x5f, 0x56, 0x41, 0x4c, 0x55, 0x45, 0x10, 0x01, 0x12, 0x12, 0x0a, 0x0e, 0x4d,
	0x45, 0x53, 0x53, 0x41, 0x47, 0x45, 0x5f, 0x52, 0x45, 0x4e, 0x44, 0x45, 0x52, 0x10, 0x02, 0x12,
	0x13, 0x0a, 0x0f, 0x4d, 0x45, 0x53, 0x53, 0x41, 0x47, 0x45, 0x5f, 0x50, 0x52, 0x45, 0x56, 0x49,
	0x45, 0x57, 0x10, 0x03, 0x12, 0x08, 0x0a, 0x04, 0x56, 0x4f, 0x54, 0x45, 0x10, 0x04, 0x42, 0x37,
	0x5a, 0x35, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6f, 0x6c,
	0x79, 0x6d, 0x70, 0x2f, 0x67, 0x6f, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x65, 0x6f, 0x6c, 0x79, 0x6d,
	0x70, 0x2f, 0x64, 0x69, 0x73, 0x63, 0x75, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x3b, 0x64, 0x69, 0x73,
	0x63, 0x75, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_eolymp_discussion_message_proto_rawDescOnce sync.Once
	file_eolymp_discussion_message_proto_rawDescData []byte
)

func file_eolymp_discussion_message_proto_rawDescGZIP() []byte {
	file_eolymp_discussion_message_proto_rawDescOnce.Do(func() {
		file_eolymp_discussion_message_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_eolymp_discussion_message_proto_rawDesc), len(file_eolymp_discussion_message_proto_rawDesc)))
	})
	return file_eolymp_discussion_message_proto_rawDescData
}

var file_eolymp_discussion_message_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_eolymp_discussion_message_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_eolymp_discussion_message_proto_goTypes = []any{
	(Message_Extra)(0),            // 0: eolymp.discussion.Message.Extra
	(*Message)(nil),               // 1: eolymp.discussion.Message
	(*timestamppb.Timestamp)(nil), // 2: google.protobuf.Timestamp
	(*ecm.Content)(nil),           // 3: eolymp.ecm.Content
	(*wellknown.Link)(nil),        // 4: eolymp.wellknown.Link
}
var file_eolymp_discussion_message_proto_depIdxs = []int32{
	2, // 0: eolymp.discussion.Message.posted_at:type_name -> google.protobuf.Timestamp
	2, // 1: eolymp.discussion.Message.edited_at:type_name -> google.protobuf.Timestamp
	2, // 2: eolymp.discussion.Message.deleted_at:type_name -> google.protobuf.Timestamp
	3, // 3: eolymp.discussion.Message.message:type_name -> eolymp.ecm.Content
	4, // 4: eolymp.discussion.Message.links:type_name -> eolymp.wellknown.Link
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_eolymp_discussion_message_proto_init() }
func file_eolymp_discussion_message_proto_init() {
	if File_eolymp_discussion_message_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_eolymp_discussion_message_proto_rawDesc), len(file_eolymp_discussion_message_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_eolymp_discussion_message_proto_goTypes,
		DependencyIndexes: file_eolymp_discussion_message_proto_depIdxs,
		EnumInfos:         file_eolymp_discussion_message_proto_enumTypes,
		MessageInfos:      file_eolymp_discussion_message_proto_msgTypes,
	}.Build()
	File_eolymp_discussion_message_proto = out.File
	file_eolymp_discussion_message_proto_goTypes = nil
	file_eolymp_discussion_message_proto_depIdxs = nil
}
