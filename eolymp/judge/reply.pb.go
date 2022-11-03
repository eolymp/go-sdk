// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.18.1
// source: eolymp/judge/reply.proto

package judge

import (
	_ "github.com/eolymp/go-sdk/eolymp/annotations"
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

type Reply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Reply unique identifier.
	Id  string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Ern string `protobuf:"bytes,9999,opt,name=ern,proto3" json:"ern,omitempty"`
	// Ticket this reply belongs to.
	TicketId string `protobuf:"bytes,2,opt,name=ticket_id,json=ticketId,proto3" json:"ticket_id,omitempty"`
	// User (jury) or Member (participant) who added this reply.
	//
	// Types that are assignable to Author:
	//	*Reply_UserId
	//	*Reply_MemberId
	Author isReply_Author `protobuf_oneof:"author"`
	// Message of the reply.
	Message string `protobuf:"bytes,10,opt,name=message,proto3" json:"message,omitempty"`
	// Timestamp when reply was created.
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,11,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
}

func (x *Reply) Reset() {
	*x = Reply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_judge_reply_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Reply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Reply) ProtoMessage() {}

func (x *Reply) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_judge_reply_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
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
	return file_eolymp_judge_reply_proto_rawDescGZIP(), []int{0}
}

func (x *Reply) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Reply) GetErn() string {
	if x != nil {
		return x.Ern
	}
	return ""
}

func (x *Reply) GetTicketId() string {
	if x != nil {
		return x.TicketId
	}
	return ""
}

func (m *Reply) GetAuthor() isReply_Author {
	if m != nil {
		return m.Author
	}
	return nil
}

func (x *Reply) GetUserId() string {
	if x, ok := x.GetAuthor().(*Reply_UserId); ok {
		return x.UserId
	}
	return ""
}

func (x *Reply) GetMemberId() string {
	if x, ok := x.GetAuthor().(*Reply_MemberId); ok {
		return x.MemberId
	}
	return ""
}

func (x *Reply) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
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

var File_eolymp_judge_reply_proto protoreflect.FileDescriptor

var file_eolymp_judge_reply_proto_rawDesc = []byte{
	0x0a, 0x18, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x6a, 0x75, 0x64, 0x67, 0x65, 0x2f, 0x72,
	0x65, 0x70, 0x6c, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x65, 0x6f, 0x6c, 0x79,
	0x6d, 0x70, 0x2e, 0x6a, 0x75, 0x64, 0x67, 0x65, 0x1a, 0x21, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70,
	0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x86, 0x02, 0x0a,
	0x05, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x11, 0x0a, 0x03, 0x65, 0x72, 0x6e, 0x18, 0x8f, 0x4e,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x65, 0x72, 0x6e, 0x12, 0x1b, 0x0a, 0x09, 0x74, 0x69, 0x63,
	0x6b, 0x65, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x69,
	0x63, 0x6b, 0x65, 0x74, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x1d, 0x0a, 0x09, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x08, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x49, 0x64,
	0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x3a, 0x24, 0xb2, 0xe3, 0x0a, 0x20, 0xba, 0xe3, 0x0a, 0x05, 0x72,
	0x65, 0x70, 0x6c, 0x79, 0xc2, 0xe3, 0x0a, 0x13, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x6a,
	0x75, 0x64, 0x67, 0x65, 0x2e, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x42, 0x08, 0x0a, 0x06, 0x61,
	0x75, 0x74, 0x68, 0x6f, 0x72, 0x42, 0x2d, 0x5a, 0x2b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x67, 0x6f, 0x2d, 0x73, 0x64,
	0x6b, 0x2f, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x6a, 0x75, 0x64, 0x67, 0x65, 0x3b, 0x6a,
	0x75, 0x64, 0x67, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_eolymp_judge_reply_proto_rawDescOnce sync.Once
	file_eolymp_judge_reply_proto_rawDescData = file_eolymp_judge_reply_proto_rawDesc
)

func file_eolymp_judge_reply_proto_rawDescGZIP() []byte {
	file_eolymp_judge_reply_proto_rawDescOnce.Do(func() {
		file_eolymp_judge_reply_proto_rawDescData = protoimpl.X.CompressGZIP(file_eolymp_judge_reply_proto_rawDescData)
	})
	return file_eolymp_judge_reply_proto_rawDescData
}

var file_eolymp_judge_reply_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_eolymp_judge_reply_proto_goTypes = []interface{}{
	(*Reply)(nil),                 // 0: eolymp.judge.Reply
	(*timestamppb.Timestamp)(nil), // 1: google.protobuf.Timestamp
}
var file_eolymp_judge_reply_proto_depIdxs = []int32{
	1, // 0: eolymp.judge.Reply.created_at:type_name -> google.protobuf.Timestamp
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_eolymp_judge_reply_proto_init() }
func file_eolymp_judge_reply_proto_init() {
	if File_eolymp_judge_reply_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_eolymp_judge_reply_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Reply); i {
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
	file_eolymp_judge_reply_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*Reply_UserId)(nil),
		(*Reply_MemberId)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_eolymp_judge_reply_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_eolymp_judge_reply_proto_goTypes,
		DependencyIndexes: file_eolymp_judge_reply_proto_depIdxs,
		MessageInfos:      file_eolymp_judge_reply_proto_msgTypes,
	}.Build()
	File_eolymp_judge_reply_proto = out.File
	file_eolymp_judge_reply_proto_rawDesc = nil
	file_eolymp_judge_reply_proto_goTypes = nil
	file_eolymp_judge_reply_proto_depIdxs = nil
}
