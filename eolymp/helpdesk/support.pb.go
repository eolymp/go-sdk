// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v4.24.4
// source: eolymp/helpdesk/support.proto

package helpdesk

import (
	_ "github.com/eolymp/go-sdk/eolymp/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_eolymp_helpdesk_support_proto protoreflect.FileDescriptor

const file_eolymp_helpdesk_support_proto_rawDesc = "" +
	"\n" +
	"\x1deolymp/helpdesk/support.proto\x12\x0feolymp.helpdesk\x1a\x1deolymp/annotations/http.proto\x1a\"eolymp/annotations/ratelimit.proto\x1a\x1eeolymp/annotations/scope.proto\x1a(eolymp/helpdesk/auto_reply_service.proto\x1a$eolymp/helpdesk/ticket_service.proto2\xfd\x1a\n" +
	"\aSupport\x12\x9e\x01\n" +
	"\fCreateTicket\x12\".eolymp.helpdesk.CreateTicketInput\x1a#.eolymp.helpdesk.CreateTicketOutput\"E\xea\xe2\n" +
	"\v\xf5\xe2\n" +
	"\x00\x00\x00@\xf8\xe2\n" +
	"\n" +
	"\x82\xe3\n" +
	"\x19\x8a\xe3\n" +
	"\x15helpdesk:ticket:write\x82\xd3\xe4\x93\x02\x13\"\x11/helpdesk/tickets\x12\xaa\x01\n" +
	"\fUpdateTicket\x12\".eolymp.helpdesk.UpdateTicketInput\x1a#.eolymp.helpdesk.UpdateTicketOutput\"Q\xea\xe2\n" +
	"\v\xf5\xe2\n" +
	"\x00\x00\x00@\xf8\xe2\n" +
	"\n" +
	"\x82\xe3\n" +
	"\x19\x8a\xe3\n" +
	"\x15helpdesk:ticket:write\x82\xd3\xe4\x93\x02\x1f\x1a\x1d/helpdesk/tickets/{ticket_id}\x12\xaa\x01\n" +
	"\fDeleteTicket\x12\".eolymp.helpdesk.DeleteTicketInput\x1a#.eolymp.helpdesk.DeleteTicketOutput\"Q\xea\xe2\n" +
	"\v\xf5\xe2\n" +
	"\x00\x00\x00@\xf8\xe2\n" +
	"\n" +
	"\x82\xe3\n" +
	"\x19\x8a\xe3\n" +
	"\x15helpdesk:ticket:write\x82\xd3\xe4\x93\x02\x1f*\x1d/helpdesk/tickets/{ticket_id}\x12\xaf\x01\n" +
	"\x0eDescribeTicket\x12$.eolymp.helpdesk.DescribeTicketInput\x1a%.eolymp.helpdesk.DescribeTicketOutput\"P\xea\xe2\n" +
	"\v\xf5\xe2\n" +
	"\x00\x00\x00@\xf8\xe2\n" +
	"\n" +
	"\x82\xe3\n" +
	"\x18\x8a\xe3\n" +
	"\x14helpdesk:ticket:read\x82\xd3\xe4\x93\x02\x1f\x12\x1d/helpdesk/tickets/{ticket_id}\x12\x9a\x01\n" +
	"\vListTickets\x12!.eolymp.helpdesk.ListTicketsInput\x1a\".eolymp.helpdesk.ListTicketsOutput\"D\xea\xe2\n" +
	"\v\xf5\xe2\n" +
	"\x00\x00\x00@\xf8\xe2\n" +
	"\n" +
	"\x82\xe3\n" +
	"\x18\x8a\xe3\n" +
	"\x14helpdesk:ticket:read\x82\xd3\xe4\x93\x02\x13\x12\x11/helpdesk/tickets\x12\xb5\x01\n" +
	"\rApproveTicket\x12#.eolymp.helpdesk.ApproveTicketInput\x1a$.eolymp.helpdesk.ApproveTicketOutput\"Y\xea\xe2\n" +
	"\v\xf5\xe2\n" +
	"\x00\x00\x00@\xf8\xe2\n" +
	"\n" +
	"\x82\xe3\n" +
	"\x19\x8a\xe3\n" +
	"\x15helpdesk:ticket:write\x82\xd3\xe4\x93\x02'\"%/helpdesk/tickets/{ticket_id}/approve\x12\xb1\x01\n" +
	"\fRejectTicket\x12\".eolymp.helpdesk.RejectTicketInput\x1a#.eolymp.helpdesk.RejectTicketOutput\"X\xea\xe2\n" +
	"\v\xf5\xe2\n" +
	"\x00\x00\x00@\xf8\xe2\n" +
	"\n" +
	"\x82\xe3\n" +
	"\x19\x8a\xe3\n" +
	"\x15helpdesk:ticket:write\x82\xd3\xe4\x93\x02&\"$/helpdesk/tickets/{ticket_id}/reject\x12\xad\x01\n" +
	"\vCloseTicket\x12!.eolymp.helpdesk.CloseTicketInput\x1a\".eolymp.helpdesk.CloseTicketOutput\"W\xea\xe2\n" +
	"\v\xf5\xe2\n" +
	"\x00\x00\x00@\xf8\xe2\n" +
	"\n" +
	"\x82\xe3\n" +
	"\x19\x8a\xe3\n" +
	"\x15helpdesk:ticket:write\x82\xd3\xe4\x93\x02%\"#/helpdesk/tickets/{ticket_id}/close\x12\xad\x01\n" +
	"\n" +
	"AddComment\x12 .eolymp.helpdesk.AddCommentInput\x1a!.eolymp.helpdesk.AddCommentOutput\"Z\xea\xe2\n" +
	"\v\xf5\xe2\n" +
	"\x00\x00\x00@\xf8\xe2\n" +
	"\n" +
	"\x82\xe3\n" +
	"\x19\x8a\xe3\n" +
	"\x15helpdesk:ticket:write\x82\xd3\xe4\x93\x02(\"&/helpdesk/tickets/{ticket_id}/comments\x12\xc3\x01\n" +
	"\rUpdateComment\x12#.eolymp.helpdesk.UpdateCommentInput\x1a$.eolymp.helpdesk.UpdateCommentOutput\"g\xea\xe2\n" +
	"\v\xf5\xe2\n" +
	"\x00\x00\x00@\xf8\xe2\n" +
	"\n" +
	"\x82\xe3\n" +
	"\x19\x8a\xe3\n" +
	"\x15helpdesk:ticket:write\x82\xd3\xe4\x93\x025\x1a3/helpdesk/tickets/{ticket_id}/comments/{comment_id}\x12\xc3\x01\n" +
	"\rDeleteComment\x12#.eolymp.helpdesk.DeleteCommentInput\x1a$.eolymp.helpdesk.DeleteCommentOutput\"g\xea\xe2\n" +
	"\v\xf5\xe2\n" +
	"\x00\x00\x00@\xf8\xe2\n" +
	"\n" +
	"\x82\xe3\n" +
	"\x19\x8a\xe3\n" +
	"\x15helpdesk:ticket:write\x82\xd3\xe4\x93\x025*3/helpdesk/tickets/{ticket_id}/comments/{comment_id}\x12\xb2\x01\n" +
	"\fListComments\x12\".eolymp.helpdesk.ListCommentsInput\x1a#.eolymp.helpdesk.ListCommentsOutput\"Y\xea\xe2\n" +
	"\v\xf5\xe2\n" +
	"\x00\x00\xa0@\xf8\xe2\n" +
	"\x14\x82\xe3\n" +
	"\x18\x8a\xe3\n" +
	"\x14helpdesk:ticket:read\x82\xd3\xe4\x93\x02(\x12&/helpdesk/tickets/{ticket_id}/comments\x12\xc9\x01\n" +
	"\x0fDescribeComment\x12%.eolymp.helpdesk.DescribeCommentInput\x1a&.eolymp.helpdesk.DescribeCommentOutput\"g\xea\xe2\n" +
	"\v\xf5\xe2\n" +
	"\x00\x00\xa0@\xf8\xe2\n" +
	"2\x82\xe3\n" +
	"\x19\x8a\xe3\n" +
	"\x15helpdesk:ticket:write\x82\xd3\xe4\x93\x025\x123/helpdesk/tickets/{ticket_id}/comments/{comment_id}\x12\xaf\x01\n" +
	"\x0fCreateAutoReply\x12%.eolymp.helpdesk.CreateAutoReplyInput\x1a&.eolymp.helpdesk.CreateAutoReplyOutput\"M\xea\xe2\n" +
	"\v\xf5\xe2\n" +
	"\x00\x00\x00@\xf8\xe2\n" +
	"\n" +
	"\x82\xe3\n" +
	"\x1d\x8a\xe3\n" +
	"\x19helpdesk:auto-reply:write\x82\xd3\xe4\x93\x02\x17\"\x15/helpdesk/autoreplies\x12\xba\x01\n" +
	"\x0fUpdateAutoReply\x12%.eolymp.helpdesk.UpdateAutoReplyInput\x1a&.eolymp.helpdesk.UpdateAutoReplyOutput\"X\xea\xe2\n" +
	"\v\xf5\xe2\n" +
	"\x00\x00\x00@\xf8\xe2\n" +
	"\n" +
	"\x82\xe3\n" +
	"\x1d\x8a\xe3\n" +
	"\x19helpdesk:auto-reply:write\x82\xd3\xe4\x93\x02\"\x1a /helpdesk/autoreplies/{reply_id}\x12\xba\x01\n" +
	"\x0fDeleteAutoReply\x12%.eolymp.helpdesk.DeleteAutoReplyInput\x1a&.eolymp.helpdesk.DeleteAutoReplyOutput\"X\xea\xe2\n" +
	"\v\xf5\xe2\n" +
	"\x00\x00\x00@\xf8\xe2\n" +
	"\n" +
	"\x82\xe3\n" +
	"\x1d\x8a\xe3\n" +
	"\x19helpdesk:auto-reply:write\x82\xd3\xe4\x93\x02\"* /helpdesk/autoreplies/{reply_id}\x12\xae\x01\n" +
	"\x0fListAutoReplies\x12%.eolymp.helpdesk.ListAutoRepliesInput\x1a&.eolymp.helpdesk.ListAutoRepliesOutput\"L\xea\xe2\n" +
	"\v\xf5\xe2\n" +
	"\x00\x00\xa0@\xf8\xe2\n" +
	"\x14\x82\xe3\n" +
	"\x1c\x8a\xe3\n" +
	"\x18helpdesk:auto-reply:read\x82\xd3\xe4\x93\x02\x17\x12\x15/helpdesk/autoreplies\x12\xc0\x01\n" +
	"\x11DescribeAutoReply\x12'.eolymp.helpdesk.DescribeAutoReplyInput\x1a(.eolymp.helpdesk.DescribeAutoReplyOutput\"X\xea\xe2\n" +
	"\v\xf5\xe2\n" +
	"\x00\x00\xa0@\xf8\xe2\n" +
	"2\x82\xe3\n" +
	"\x1d\x8a\xe3\n" +
	"\x19helpdesk:auto-reply:write\x82\xd3\xe4\x93\x02\"\x12 /helpdesk/autoreplies/{reply_id}\x12\xae\x01\n" +
	"\x10UploadAttachment\x12&.eolymp.helpdesk.UploadAttachmentInput\x1a'.eolymp.helpdesk.UploadAttachmentOutput\"I\xea\xe2\n" +
	"\v\xf5\xe2\n" +
	"\x00\x00\xa0@\xf8\xe2\n" +
	"2\x82\xe3\n" +
	"\x19\x8a\xe3\n" +
	"\x15helpdesk:ticket:write\x82\xd3\xe4\x93\x02\x17\"\x15/helpdesk/attachmentsB3Z1github.com/eolymp/go-sdk/eolymp/helpdesk;helpdeskb\x06proto3"

var file_eolymp_helpdesk_support_proto_goTypes = []any{
	(*CreateTicketInput)(nil),       // 0: eolymp.helpdesk.CreateTicketInput
	(*UpdateTicketInput)(nil),       // 1: eolymp.helpdesk.UpdateTicketInput
	(*DeleteTicketInput)(nil),       // 2: eolymp.helpdesk.DeleteTicketInput
	(*DescribeTicketInput)(nil),     // 3: eolymp.helpdesk.DescribeTicketInput
	(*ListTicketsInput)(nil),        // 4: eolymp.helpdesk.ListTicketsInput
	(*ApproveTicketInput)(nil),      // 5: eolymp.helpdesk.ApproveTicketInput
	(*RejectTicketInput)(nil),       // 6: eolymp.helpdesk.RejectTicketInput
	(*CloseTicketInput)(nil),        // 7: eolymp.helpdesk.CloseTicketInput
	(*AddCommentInput)(nil),         // 8: eolymp.helpdesk.AddCommentInput
	(*UpdateCommentInput)(nil),      // 9: eolymp.helpdesk.UpdateCommentInput
	(*DeleteCommentInput)(nil),      // 10: eolymp.helpdesk.DeleteCommentInput
	(*ListCommentsInput)(nil),       // 11: eolymp.helpdesk.ListCommentsInput
	(*DescribeCommentInput)(nil),    // 12: eolymp.helpdesk.DescribeCommentInput
	(*CreateAutoReplyInput)(nil),    // 13: eolymp.helpdesk.CreateAutoReplyInput
	(*UpdateAutoReplyInput)(nil),    // 14: eolymp.helpdesk.UpdateAutoReplyInput
	(*DeleteAutoReplyInput)(nil),    // 15: eolymp.helpdesk.DeleteAutoReplyInput
	(*ListAutoRepliesInput)(nil),    // 16: eolymp.helpdesk.ListAutoRepliesInput
	(*DescribeAutoReplyInput)(nil),  // 17: eolymp.helpdesk.DescribeAutoReplyInput
	(*UploadAttachmentInput)(nil),   // 18: eolymp.helpdesk.UploadAttachmentInput
	(*CreateTicketOutput)(nil),      // 19: eolymp.helpdesk.CreateTicketOutput
	(*UpdateTicketOutput)(nil),      // 20: eolymp.helpdesk.UpdateTicketOutput
	(*DeleteTicketOutput)(nil),      // 21: eolymp.helpdesk.DeleteTicketOutput
	(*DescribeTicketOutput)(nil),    // 22: eolymp.helpdesk.DescribeTicketOutput
	(*ListTicketsOutput)(nil),       // 23: eolymp.helpdesk.ListTicketsOutput
	(*ApproveTicketOutput)(nil),     // 24: eolymp.helpdesk.ApproveTicketOutput
	(*RejectTicketOutput)(nil),      // 25: eolymp.helpdesk.RejectTicketOutput
	(*CloseTicketOutput)(nil),       // 26: eolymp.helpdesk.CloseTicketOutput
	(*AddCommentOutput)(nil),        // 27: eolymp.helpdesk.AddCommentOutput
	(*UpdateCommentOutput)(nil),     // 28: eolymp.helpdesk.UpdateCommentOutput
	(*DeleteCommentOutput)(nil),     // 29: eolymp.helpdesk.DeleteCommentOutput
	(*ListCommentsOutput)(nil),      // 30: eolymp.helpdesk.ListCommentsOutput
	(*DescribeCommentOutput)(nil),   // 31: eolymp.helpdesk.DescribeCommentOutput
	(*CreateAutoReplyOutput)(nil),   // 32: eolymp.helpdesk.CreateAutoReplyOutput
	(*UpdateAutoReplyOutput)(nil),   // 33: eolymp.helpdesk.UpdateAutoReplyOutput
	(*DeleteAutoReplyOutput)(nil),   // 34: eolymp.helpdesk.DeleteAutoReplyOutput
	(*ListAutoRepliesOutput)(nil),   // 35: eolymp.helpdesk.ListAutoRepliesOutput
	(*DescribeAutoReplyOutput)(nil), // 36: eolymp.helpdesk.DescribeAutoReplyOutput
	(*UploadAttachmentOutput)(nil),  // 37: eolymp.helpdesk.UploadAttachmentOutput
}
var file_eolymp_helpdesk_support_proto_depIdxs = []int32{
	0,  // 0: eolymp.helpdesk.Support.CreateTicket:input_type -> eolymp.helpdesk.CreateTicketInput
	1,  // 1: eolymp.helpdesk.Support.UpdateTicket:input_type -> eolymp.helpdesk.UpdateTicketInput
	2,  // 2: eolymp.helpdesk.Support.DeleteTicket:input_type -> eolymp.helpdesk.DeleteTicketInput
	3,  // 3: eolymp.helpdesk.Support.DescribeTicket:input_type -> eolymp.helpdesk.DescribeTicketInput
	4,  // 4: eolymp.helpdesk.Support.ListTickets:input_type -> eolymp.helpdesk.ListTicketsInput
	5,  // 5: eolymp.helpdesk.Support.ApproveTicket:input_type -> eolymp.helpdesk.ApproveTicketInput
	6,  // 6: eolymp.helpdesk.Support.RejectTicket:input_type -> eolymp.helpdesk.RejectTicketInput
	7,  // 7: eolymp.helpdesk.Support.CloseTicket:input_type -> eolymp.helpdesk.CloseTicketInput
	8,  // 8: eolymp.helpdesk.Support.AddComment:input_type -> eolymp.helpdesk.AddCommentInput
	9,  // 9: eolymp.helpdesk.Support.UpdateComment:input_type -> eolymp.helpdesk.UpdateCommentInput
	10, // 10: eolymp.helpdesk.Support.DeleteComment:input_type -> eolymp.helpdesk.DeleteCommentInput
	11, // 11: eolymp.helpdesk.Support.ListComments:input_type -> eolymp.helpdesk.ListCommentsInput
	12, // 12: eolymp.helpdesk.Support.DescribeComment:input_type -> eolymp.helpdesk.DescribeCommentInput
	13, // 13: eolymp.helpdesk.Support.CreateAutoReply:input_type -> eolymp.helpdesk.CreateAutoReplyInput
	14, // 14: eolymp.helpdesk.Support.UpdateAutoReply:input_type -> eolymp.helpdesk.UpdateAutoReplyInput
	15, // 15: eolymp.helpdesk.Support.DeleteAutoReply:input_type -> eolymp.helpdesk.DeleteAutoReplyInput
	16, // 16: eolymp.helpdesk.Support.ListAutoReplies:input_type -> eolymp.helpdesk.ListAutoRepliesInput
	17, // 17: eolymp.helpdesk.Support.DescribeAutoReply:input_type -> eolymp.helpdesk.DescribeAutoReplyInput
	18, // 18: eolymp.helpdesk.Support.UploadAttachment:input_type -> eolymp.helpdesk.UploadAttachmentInput
	19, // 19: eolymp.helpdesk.Support.CreateTicket:output_type -> eolymp.helpdesk.CreateTicketOutput
	20, // 20: eolymp.helpdesk.Support.UpdateTicket:output_type -> eolymp.helpdesk.UpdateTicketOutput
	21, // 21: eolymp.helpdesk.Support.DeleteTicket:output_type -> eolymp.helpdesk.DeleteTicketOutput
	22, // 22: eolymp.helpdesk.Support.DescribeTicket:output_type -> eolymp.helpdesk.DescribeTicketOutput
	23, // 23: eolymp.helpdesk.Support.ListTickets:output_type -> eolymp.helpdesk.ListTicketsOutput
	24, // 24: eolymp.helpdesk.Support.ApproveTicket:output_type -> eolymp.helpdesk.ApproveTicketOutput
	25, // 25: eolymp.helpdesk.Support.RejectTicket:output_type -> eolymp.helpdesk.RejectTicketOutput
	26, // 26: eolymp.helpdesk.Support.CloseTicket:output_type -> eolymp.helpdesk.CloseTicketOutput
	27, // 27: eolymp.helpdesk.Support.AddComment:output_type -> eolymp.helpdesk.AddCommentOutput
	28, // 28: eolymp.helpdesk.Support.UpdateComment:output_type -> eolymp.helpdesk.UpdateCommentOutput
	29, // 29: eolymp.helpdesk.Support.DeleteComment:output_type -> eolymp.helpdesk.DeleteCommentOutput
	30, // 30: eolymp.helpdesk.Support.ListComments:output_type -> eolymp.helpdesk.ListCommentsOutput
	31, // 31: eolymp.helpdesk.Support.DescribeComment:output_type -> eolymp.helpdesk.DescribeCommentOutput
	32, // 32: eolymp.helpdesk.Support.CreateAutoReply:output_type -> eolymp.helpdesk.CreateAutoReplyOutput
	33, // 33: eolymp.helpdesk.Support.UpdateAutoReply:output_type -> eolymp.helpdesk.UpdateAutoReplyOutput
	34, // 34: eolymp.helpdesk.Support.DeleteAutoReply:output_type -> eolymp.helpdesk.DeleteAutoReplyOutput
	35, // 35: eolymp.helpdesk.Support.ListAutoReplies:output_type -> eolymp.helpdesk.ListAutoRepliesOutput
	36, // 36: eolymp.helpdesk.Support.DescribeAutoReply:output_type -> eolymp.helpdesk.DescribeAutoReplyOutput
	37, // 37: eolymp.helpdesk.Support.UploadAttachment:output_type -> eolymp.helpdesk.UploadAttachmentOutput
	19, // [19:38] is the sub-list for method output_type
	0,  // [0:19] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_eolymp_helpdesk_support_proto_init() }
func file_eolymp_helpdesk_support_proto_init() {
	if File_eolymp_helpdesk_support_proto != nil {
		return
	}
	file_eolymp_helpdesk_auto_reply_service_proto_init()
	file_eolymp_helpdesk_ticket_service_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_eolymp_helpdesk_support_proto_rawDesc), len(file_eolymp_helpdesk_support_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_eolymp_helpdesk_support_proto_goTypes,
		DependencyIndexes: file_eolymp_helpdesk_support_proto_depIdxs,
	}.Build()
	File_eolymp_helpdesk_support_proto = out.File
	file_eolymp_helpdesk_support_proto_goTypes = nil
	file_eolymp_helpdesk_support_proto_depIdxs = nil
}
