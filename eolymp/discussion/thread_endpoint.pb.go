// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.1
// 	protoc        v4.24.4
// source: eolymp/discussion/thread_endpoint.proto

package discussion

import (
	_ "github.com/eolymp/go-sdk/eolymp/annotations"
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

type DescribeThreadInput struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Locale        string                 `protobuf:"bytes,10,opt,name=locale,proto3" json:"locale,omitempty"` // locale preference
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DescribeThreadInput) Reset() {
	*x = DescribeThreadInput{}
	mi := &file_eolymp_discussion_thread_endpoint_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DescribeThreadInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescribeThreadInput) ProtoMessage() {}

func (x *DescribeThreadInput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_discussion_thread_endpoint_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DescribeThreadInput.ProtoReflect.Descriptor instead.
func (*DescribeThreadInput) Descriptor() ([]byte, []int) {
	return file_eolymp_discussion_thread_endpoint_proto_rawDescGZIP(), []int{0}
}

func (x *DescribeThreadInput) GetLocale() string {
	if x != nil {
		return x.Locale
	}
	return ""
}

type DescribeThreadOutput struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Thread        *Thread                `protobuf:"bytes,1,opt,name=thread,proto3" json:"thread,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DescribeThreadOutput) Reset() {
	*x = DescribeThreadOutput{}
	mi := &file_eolymp_discussion_thread_endpoint_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DescribeThreadOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescribeThreadOutput) ProtoMessage() {}

func (x *DescribeThreadOutput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_discussion_thread_endpoint_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DescribeThreadOutput.ProtoReflect.Descriptor instead.
func (*DescribeThreadOutput) Descriptor() ([]byte, []int) {
	return file_eolymp_discussion_thread_endpoint_proto_rawDescGZIP(), []int{1}
}

func (x *DescribeThreadOutput) GetThread() *Thread {
	if x != nil {
		return x.Thread
	}
	return nil
}

type VoteThreadInput struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	MessageId     string                 `protobuf:"bytes,1,opt,name=message_id,json=messageId,proto3" json:"message_id,omitempty"`
	Vote          int32                  `protobuf:"varint,2,opt,name=vote,proto3" json:"vote,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *VoteThreadInput) Reset() {
	*x = VoteThreadInput{}
	mi := &file_eolymp_discussion_thread_endpoint_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *VoteThreadInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VoteThreadInput) ProtoMessage() {}

func (x *VoteThreadInput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_discussion_thread_endpoint_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VoteThreadInput.ProtoReflect.Descriptor instead.
func (*VoteThreadInput) Descriptor() ([]byte, []int) {
	return file_eolymp_discussion_thread_endpoint_proto_rawDescGZIP(), []int{2}
}

func (x *VoteThreadInput) GetMessageId() string {
	if x != nil {
		return x.MessageId
	}
	return ""
}

func (x *VoteThreadInput) GetVote() int32 {
	if x != nil {
		return x.Vote
	}
	return 0
}

type VoteThreadOutput struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	VoteCount     int32                  `protobuf:"varint,1,opt,name=vote_count,json=voteCount,proto3" json:"vote_count,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *VoteThreadOutput) Reset() {
	*x = VoteThreadOutput{}
	mi := &file_eolymp_discussion_thread_endpoint_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *VoteThreadOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VoteThreadOutput) ProtoMessage() {}

func (x *VoteThreadOutput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_discussion_thread_endpoint_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VoteThreadOutput.ProtoReflect.Descriptor instead.
func (*VoteThreadOutput) Descriptor() ([]byte, []int) {
	return file_eolymp_discussion_thread_endpoint_proto_rawDescGZIP(), []int{3}
}

func (x *VoteThreadOutput) GetVoteCount() int32 {
	if x != nil {
		return x.VoteCount
	}
	return 0
}

var File_eolymp_discussion_thread_endpoint_proto protoreflect.FileDescriptor

var file_eolymp_discussion_thread_endpoint_proto_rawDesc = []byte{
	0x0a, 0x27, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x64, 0x69, 0x73, 0x63, 0x75, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x2f, 0x74, 0x68, 0x72, 0x65, 0x61, 0x64, 0x5f, 0x65, 0x6e, 0x64, 0x70, 0x6f,
	0x69, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11, 0x65, 0x6f, 0x6c, 0x79, 0x6d,
	0x70, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x75, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x1a, 0x1d, 0x65, 0x6f,
	0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2f, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x22, 0x65, 0x6f, 0x6c,
	0x79, 0x6d, 0x70, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f,
	0x72, 0x61, 0x74, 0x65, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x64, 0x69, 0x73, 0x63, 0x75, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x2f, 0x74, 0x68, 0x72, 0x65, 0x61, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x2d, 0x0a, 0x13, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x54, 0x68, 0x72, 0x65, 0x61,
	0x64, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x65,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x65, 0x22, 0x49,
	0x0a, 0x14, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x54, 0x68, 0x72, 0x65, 0x61, 0x64,
	0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x12, 0x31, 0x0a, 0x06, 0x74, 0x68, 0x72, 0x65, 0x61, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e,
	0x64, 0x69, 0x73, 0x63, 0x75, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x54, 0x68, 0x72, 0x65, 0x61,
	0x64, 0x52, 0x06, 0x74, 0x68, 0x72, 0x65, 0x61, 0x64, 0x22, 0x44, 0x0a, 0x0f, 0x56, 0x6f, 0x74,
	0x65, 0x54, 0x68, 0x72, 0x65, 0x61, 0x64, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x1d, 0x0a, 0x0a,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x76,
	0x6f, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x76, 0x6f, 0x74, 0x65, 0x22,
	0x31, 0x0a, 0x10, 0x56, 0x6f, 0x74, 0x65, 0x54, 0x68, 0x72, 0x65, 0x61, 0x64, 0x4f, 0x75, 0x74,
	0x70, 0x75, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x76, 0x6f, 0x74, 0x65, 0x5f, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x76, 0x6f, 0x74, 0x65, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x32, 0x81, 0x02, 0x0a, 0x0e, 0x54, 0x68, 0x72, 0x65, 0x61, 0x64, 0x45, 0x6e, 0x64,
	0x70, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x7a, 0x0a, 0x0e, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62,
	0x65, 0x54, 0x68, 0x72, 0x65, 0x61, 0x64, 0x12, 0x26, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70,
	0x2e, 0x64, 0x69, 0x73, 0x63, 0x75, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x44, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x62, 0x65, 0x54, 0x68, 0x72, 0x65, 0x61, 0x64, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a,
	0x27, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x75, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x2e, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x54, 0x68, 0x72, 0x65,
	0x61, 0x64, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x22, 0x17, 0xea, 0xe2, 0x0a, 0x0b, 0xf5, 0xe2,
	0x0a, 0x00, 0x00, 0xa0, 0x41, 0xf8, 0xe2, 0x0a, 0x64, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x02, 0x12,
	0x00, 0x12, 0x73, 0x0a, 0x0a, 0x56, 0x6f, 0x74, 0x65, 0x54, 0x68, 0x72, 0x65, 0x61, 0x64, 0x12,
	0x22, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x75, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x2e, 0x56, 0x6f, 0x74, 0x65, 0x54, 0x68, 0x72, 0x65, 0x61, 0x64, 0x49, 0x6e,
	0x70, 0x75, 0x74, 0x1a, 0x23, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x64, 0x69, 0x73,
	0x63, 0x75, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x56, 0x6f, 0x74, 0x65, 0x54, 0x68, 0x72, 0x65,
	0x61, 0x64, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x22, 0x1c, 0xea, 0xe2, 0x0a, 0x0b, 0xf5, 0xe2,
	0x0a, 0x00, 0x00, 0xa0, 0x40, 0xf8, 0xe2, 0x0a, 0x32, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x07, 0x22,
	0x05, 0x2f, 0x76, 0x6f, 0x74, 0x65, 0x42, 0x37, 0x5a, 0x35, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x67, 0x6f, 0x2d, 0x73,
	0x64, 0x6b, 0x2f, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x64, 0x69, 0x73, 0x63, 0x75, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x3b, 0x64, 0x69, 0x73, 0x63, 0x75, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_eolymp_discussion_thread_endpoint_proto_rawDescOnce sync.Once
	file_eolymp_discussion_thread_endpoint_proto_rawDescData = file_eolymp_discussion_thread_endpoint_proto_rawDesc
)

func file_eolymp_discussion_thread_endpoint_proto_rawDescGZIP() []byte {
	file_eolymp_discussion_thread_endpoint_proto_rawDescOnce.Do(func() {
		file_eolymp_discussion_thread_endpoint_proto_rawDescData = protoimpl.X.CompressGZIP(file_eolymp_discussion_thread_endpoint_proto_rawDescData)
	})
	return file_eolymp_discussion_thread_endpoint_proto_rawDescData
}

var file_eolymp_discussion_thread_endpoint_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_eolymp_discussion_thread_endpoint_proto_goTypes = []any{
	(*DescribeThreadInput)(nil),  // 0: eolymp.discussion.DescribeThreadInput
	(*DescribeThreadOutput)(nil), // 1: eolymp.discussion.DescribeThreadOutput
	(*VoteThreadInput)(nil),      // 2: eolymp.discussion.VoteThreadInput
	(*VoteThreadOutput)(nil),     // 3: eolymp.discussion.VoteThreadOutput
	(*Thread)(nil),               // 4: eolymp.discussion.Thread
}
var file_eolymp_discussion_thread_endpoint_proto_depIdxs = []int32{
	4, // 0: eolymp.discussion.DescribeThreadOutput.thread:type_name -> eolymp.discussion.Thread
	0, // 1: eolymp.discussion.ThreadEndpoint.DescribeThread:input_type -> eolymp.discussion.DescribeThreadInput
	2, // 2: eolymp.discussion.ThreadEndpoint.VoteThread:input_type -> eolymp.discussion.VoteThreadInput
	1, // 3: eolymp.discussion.ThreadEndpoint.DescribeThread:output_type -> eolymp.discussion.DescribeThreadOutput
	3, // 4: eolymp.discussion.ThreadEndpoint.VoteThread:output_type -> eolymp.discussion.VoteThreadOutput
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_eolymp_discussion_thread_endpoint_proto_init() }
func file_eolymp_discussion_thread_endpoint_proto_init() {
	if File_eolymp_discussion_thread_endpoint_proto != nil {
		return
	}
	file_eolymp_discussion_thread_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_eolymp_discussion_thread_endpoint_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_eolymp_discussion_thread_endpoint_proto_goTypes,
		DependencyIndexes: file_eolymp_discussion_thread_endpoint_proto_depIdxs,
		MessageInfos:      file_eolymp_discussion_thread_endpoint_proto_msgTypes,
	}.Build()
	File_eolymp_discussion_thread_endpoint_proto = out.File
	file_eolymp_discussion_thread_endpoint_proto_rawDesc = nil
	file_eolymp_discussion_thread_endpoint_proto_goTypes = nil
	file_eolymp_discussion_thread_endpoint_proto_depIdxs = nil
}
