// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v4.24.4
// source: eolymp/helpdesk/events.proto

package helpdesk

import (
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

type TicketChangedEvent struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Scope         string                 `protobuf:"bytes,10,opt,name=scope,proto3" json:"scope,omitempty"`
	Before        *Ticket                `protobuf:"bytes,1,opt,name=before,proto3" json:"before,omitempty"`
	After         *Ticket                `protobuf:"bytes,2,opt,name=after,proto3" json:"after,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *TicketChangedEvent) Reset() {
	*x = TicketChangedEvent{}
	mi := &file_eolymp_helpdesk_events_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TicketChangedEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TicketChangedEvent) ProtoMessage() {}

func (x *TicketChangedEvent) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_helpdesk_events_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TicketChangedEvent.ProtoReflect.Descriptor instead.
func (*TicketChangedEvent) Descriptor() ([]byte, []int) {
	return file_eolymp_helpdesk_events_proto_rawDescGZIP(), []int{0}
}

func (x *TicketChangedEvent) GetScope() string {
	if x != nil {
		return x.Scope
	}
	return ""
}

func (x *TicketChangedEvent) GetBefore() *Ticket {
	if x != nil {
		return x.Before
	}
	return nil
}

func (x *TicketChangedEvent) GetAfter() *Ticket {
	if x != nil {
		return x.After
	}
	return nil
}

type CommentChangedEvent struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	TicketId      string                 `protobuf:"bytes,1,opt,name=ticket_id,json=ticketId,proto3" json:"ticket_id,omitempty"`
	Before        *Ticket_Comment        `protobuf:"bytes,2,opt,name=before,proto3" json:"before,omitempty"`
	After         *Ticket_Comment        `protobuf:"bytes,3,opt,name=after,proto3" json:"after,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CommentChangedEvent) Reset() {
	*x = CommentChangedEvent{}
	mi := &file_eolymp_helpdesk_events_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CommentChangedEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommentChangedEvent) ProtoMessage() {}

func (x *CommentChangedEvent) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_helpdesk_events_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommentChangedEvent.ProtoReflect.Descriptor instead.
func (*CommentChangedEvent) Descriptor() ([]byte, []int) {
	return file_eolymp_helpdesk_events_proto_rawDescGZIP(), []int{1}
}

func (x *CommentChangedEvent) GetTicketId() string {
	if x != nil {
		return x.TicketId
	}
	return ""
}

func (x *CommentChangedEvent) GetBefore() *Ticket_Comment {
	if x != nil {
		return x.Before
	}
	return nil
}

func (x *CommentChangedEvent) GetAfter() *Ticket_Comment {
	if x != nil {
		return x.After
	}
	return nil
}

var File_eolymp_helpdesk_events_proto protoreflect.FileDescriptor

const file_eolymp_helpdesk_events_proto_rawDesc = "" +
	"\n" +
	"\x1ceolymp/helpdesk/events.proto\x12\x0feolymp.helpdesk\x1a\x1ceolymp/helpdesk/ticket.proto\"\x8a\x01\n" +
	"\x12TicketChangedEvent\x12\x14\n" +
	"\x05scope\x18\n" +
	" \x01(\tR\x05scope\x12/\n" +
	"\x06before\x18\x01 \x01(\v2\x17.eolymp.helpdesk.TicketR\x06before\x12-\n" +
	"\x05after\x18\x02 \x01(\v2\x17.eolymp.helpdesk.TicketR\x05after\"\xa2\x01\n" +
	"\x13CommentChangedEvent\x12\x1b\n" +
	"\tticket_id\x18\x01 \x01(\tR\bticketId\x127\n" +
	"\x06before\x18\x02 \x01(\v2\x1f.eolymp.helpdesk.Ticket.CommentR\x06before\x125\n" +
	"\x05after\x18\x03 \x01(\v2\x1f.eolymp.helpdesk.Ticket.CommentR\x05afterB3Z1github.com/eolymp/go-sdk/eolymp/helpdesk;helpdeskb\x06proto3"

var (
	file_eolymp_helpdesk_events_proto_rawDescOnce sync.Once
	file_eolymp_helpdesk_events_proto_rawDescData []byte
)

func file_eolymp_helpdesk_events_proto_rawDescGZIP() []byte {
	file_eolymp_helpdesk_events_proto_rawDescOnce.Do(func() {
		file_eolymp_helpdesk_events_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_eolymp_helpdesk_events_proto_rawDesc), len(file_eolymp_helpdesk_events_proto_rawDesc)))
	})
	return file_eolymp_helpdesk_events_proto_rawDescData
}

var file_eolymp_helpdesk_events_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_eolymp_helpdesk_events_proto_goTypes = []any{
	(*TicketChangedEvent)(nil),  // 0: eolymp.helpdesk.TicketChangedEvent
	(*CommentChangedEvent)(nil), // 1: eolymp.helpdesk.CommentChangedEvent
	(*Ticket)(nil),              // 2: eolymp.helpdesk.Ticket
	(*Ticket_Comment)(nil),      // 3: eolymp.helpdesk.Ticket.Comment
}
var file_eolymp_helpdesk_events_proto_depIdxs = []int32{
	2, // 0: eolymp.helpdesk.TicketChangedEvent.before:type_name -> eolymp.helpdesk.Ticket
	2, // 1: eolymp.helpdesk.TicketChangedEvent.after:type_name -> eolymp.helpdesk.Ticket
	3, // 2: eolymp.helpdesk.CommentChangedEvent.before:type_name -> eolymp.helpdesk.Ticket.Comment
	3, // 3: eolymp.helpdesk.CommentChangedEvent.after:type_name -> eolymp.helpdesk.Ticket.Comment
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_eolymp_helpdesk_events_proto_init() }
func file_eolymp_helpdesk_events_proto_init() {
	if File_eolymp_helpdesk_events_proto != nil {
		return
	}
	file_eolymp_helpdesk_ticket_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_eolymp_helpdesk_events_proto_rawDesc), len(file_eolymp_helpdesk_events_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_eolymp_helpdesk_events_proto_goTypes,
		DependencyIndexes: file_eolymp_helpdesk_events_proto_depIdxs,
		MessageInfos:      file_eolymp_helpdesk_events_proto_msgTypes,
	}.Build()
	File_eolymp_helpdesk_events_proto = out.File
	file_eolymp_helpdesk_events_proto_goTypes = nil
	file_eolymp_helpdesk_events_proto_depIdxs = nil
}
