// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v3.21.12
// source: eolymp/helpdesk/events.proto

package helpdesk

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

type TicketChangedEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Before *Ticket `protobuf:"bytes,1,opt,name=before,proto3" json:"before,omitempty"`
	After  *Ticket `protobuf:"bytes,2,opt,name=after,proto3" json:"after,omitempty"`
}

func (x *TicketChangedEvent) Reset() {
	*x = TicketChangedEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_helpdesk_events_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TicketChangedEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TicketChangedEvent) ProtoMessage() {}

func (x *TicketChangedEvent) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_helpdesk_events_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
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
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TicketId string          `protobuf:"bytes,1,opt,name=ticket_id,json=ticketId,proto3" json:"ticket_id,omitempty"`
	Before   *Ticket_Comment `protobuf:"bytes,2,opt,name=before,proto3" json:"before,omitempty"`
	After    *Ticket_Comment `protobuf:"bytes,3,opt,name=after,proto3" json:"after,omitempty"`
}

func (x *CommentChangedEvent) Reset() {
	*x = CommentChangedEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_helpdesk_events_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommentChangedEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommentChangedEvent) ProtoMessage() {}

func (x *CommentChangedEvent) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_helpdesk_events_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
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

var file_eolymp_helpdesk_events_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x68, 0x65, 0x6c, 0x70, 0x64, 0x65, 0x73,
	0x6b, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f,
	0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x68, 0x65, 0x6c, 0x70, 0x64, 0x65, 0x73, 0x6b, 0x1a,
	0x1c, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x68, 0x65, 0x6c, 0x70, 0x64, 0x65, 0x73, 0x6b,
	0x2f, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x74, 0x0a,
	0x12, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x64, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x12, 0x2f, 0x0a, 0x06, 0x62, 0x65, 0x66, 0x6f, 0x72, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x68, 0x65, 0x6c,
	0x70, 0x64, 0x65, 0x73, 0x6b, 0x2e, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x06, 0x62, 0x65,
	0x66, 0x6f, 0x72, 0x65, 0x12, 0x2d, 0x0a, 0x05, 0x61, 0x66, 0x74, 0x65, 0x72, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x68, 0x65, 0x6c,
	0x70, 0x64, 0x65, 0x73, 0x6b, 0x2e, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x05, 0x61, 0x66,
	0x74, 0x65, 0x72, 0x22, 0xa2, 0x01, 0x0a, 0x13, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x43,
	0x68, 0x61, 0x6e, 0x67, 0x65, 0x64, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x74,
	0x69, 0x63, 0x6b, 0x65, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x49, 0x64, 0x12, 0x37, 0x0a, 0x06, 0x62, 0x65, 0x66, 0x6f,
	0x72, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d,
	0x70, 0x2e, 0x68, 0x65, 0x6c, 0x70, 0x64, 0x65, 0x73, 0x6b, 0x2e, 0x54, 0x69, 0x63, 0x6b, 0x65,
	0x74, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x06, 0x62, 0x65, 0x66, 0x6f, 0x72,
	0x65, 0x12, 0x35, 0x0a, 0x05, 0x61, 0x66, 0x74, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1f, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x68, 0x65, 0x6c, 0x70, 0x64, 0x65,
	0x73, 0x6b, 0x2e, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e,
	0x74, 0x52, 0x05, 0x61, 0x66, 0x74, 0x65, 0x72, 0x42, 0x33, 0x5a, 0x31, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x67, 0x6f,
	0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x68, 0x65, 0x6c, 0x70,
	0x64, 0x65, 0x73, 0x6b, 0x3b, 0x68, 0x65, 0x6c, 0x70, 0x64, 0x65, 0x73, 0x6b, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_eolymp_helpdesk_events_proto_rawDescOnce sync.Once
	file_eolymp_helpdesk_events_proto_rawDescData = file_eolymp_helpdesk_events_proto_rawDesc
)

func file_eolymp_helpdesk_events_proto_rawDescGZIP() []byte {
	file_eolymp_helpdesk_events_proto_rawDescOnce.Do(func() {
		file_eolymp_helpdesk_events_proto_rawDescData = protoimpl.X.CompressGZIP(file_eolymp_helpdesk_events_proto_rawDescData)
	})
	return file_eolymp_helpdesk_events_proto_rawDescData
}

var file_eolymp_helpdesk_events_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_eolymp_helpdesk_events_proto_goTypes = []interface{}{
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
	if !protoimpl.UnsafeEnabled {
		file_eolymp_helpdesk_events_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TicketChangedEvent); i {
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
		file_eolymp_helpdesk_events_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommentChangedEvent); i {
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
			RawDescriptor: file_eolymp_helpdesk_events_proto_rawDesc,
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
	file_eolymp_helpdesk_events_proto_rawDesc = nil
	file_eolymp_helpdesk_events_proto_goTypes = nil
	file_eolymp_helpdesk_events_proto_depIdxs = nil
}
