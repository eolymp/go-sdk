// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v3.21.12
// source: eolymp/notify/notification.proto

package notify

import (
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

type Notification struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Read       bool                   `protobuf:"varint,2,opt,name=read,proto3" json:"read,omitempty"`
	Topic      string                 `protobuf:"bytes,10,opt,name=topic,proto3" json:"topic,omitempty"`
	Parameters map[string]string      `protobuf:"bytes,11,rep,name=parameters,proto3" json:"parameters,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	CreatedAt  *timestamppb.Timestamp `protobuf:"bytes,20,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	ReadAt     *timestamppb.Timestamp `protobuf:"bytes,21,opt,name=read_at,json=readAt,proto3" json:"read_at,omitempty"`
}

func (x *Notification) Reset() {
	*x = Notification{}
	mi := &file_eolymp_notify_notification_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Notification) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Notification) ProtoMessage() {}

func (x *Notification) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_notify_notification_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Notification.ProtoReflect.Descriptor instead.
func (*Notification) Descriptor() ([]byte, []int) {
	return file_eolymp_notify_notification_proto_rawDescGZIP(), []int{0}
}

func (x *Notification) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Notification) GetRead() bool {
	if x != nil {
		return x.Read
	}
	return false
}

func (x *Notification) GetTopic() string {
	if x != nil {
		return x.Topic
	}
	return ""
}

func (x *Notification) GetParameters() map[string]string {
	if x != nil {
		return x.Parameters
	}
	return nil
}

func (x *Notification) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Notification) GetReadAt() *timestamppb.Timestamp {
	if x != nil {
		return x.ReadAt
	}
	return nil
}

var File_eolymp_notify_notification_proto protoreflect.FileDescriptor

var file_eolymp_notify_notification_proto_rawDesc = []byte{
	0x0a, 0x20, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2f,
	0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0d, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66,
	0x79, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xc4, 0x02, 0x0a, 0x0c, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x65, 0x61, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x04, 0x72, 0x65, 0x61, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x70, 0x69, 0x63,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x12, 0x4b, 0x0a,
	0x0a, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x18, 0x0b, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x2b, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66,
	0x79, 0x2e, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x50,
	0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0a,
	0x70, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x14, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x33, 0x0a, 0x07, 0x72, 0x65, 0x61, 0x64, 0x5f, 0x61, 0x74,
	0x18, 0x15, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x06, 0x72, 0x65, 0x61, 0x64, 0x41, 0x74, 0x1a, 0x3d, 0x0a, 0x0f, 0x50, 0x61,
	0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x2f, 0x5a, 0x2d, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x67,
	0x6f, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x6e, 0x6f, 0x74,
	0x69, 0x66, 0x79, 0x3b, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_eolymp_notify_notification_proto_rawDescOnce sync.Once
	file_eolymp_notify_notification_proto_rawDescData = file_eolymp_notify_notification_proto_rawDesc
)

func file_eolymp_notify_notification_proto_rawDescGZIP() []byte {
	file_eolymp_notify_notification_proto_rawDescOnce.Do(func() {
		file_eolymp_notify_notification_proto_rawDescData = protoimpl.X.CompressGZIP(file_eolymp_notify_notification_proto_rawDescData)
	})
	return file_eolymp_notify_notification_proto_rawDescData
}

var file_eolymp_notify_notification_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_eolymp_notify_notification_proto_goTypes = []any{
	(*Notification)(nil),          // 0: eolymp.notify.Notification
	nil,                           // 1: eolymp.notify.Notification.ParametersEntry
	(*timestamppb.Timestamp)(nil), // 2: google.protobuf.Timestamp
}
var file_eolymp_notify_notification_proto_depIdxs = []int32{
	1, // 0: eolymp.notify.Notification.parameters:type_name -> eolymp.notify.Notification.ParametersEntry
	2, // 1: eolymp.notify.Notification.created_at:type_name -> google.protobuf.Timestamp
	2, // 2: eolymp.notify.Notification.read_at:type_name -> google.protobuf.Timestamp
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_eolymp_notify_notification_proto_init() }
func file_eolymp_notify_notification_proto_init() {
	if File_eolymp_notify_notification_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_eolymp_notify_notification_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_eolymp_notify_notification_proto_goTypes,
		DependencyIndexes: file_eolymp_notify_notification_proto_depIdxs,
		MessageInfos:      file_eolymp_notify_notification_proto_msgTypes,
	}.Build()
	File_eolymp_notify_notification_proto = out.File
	file_eolymp_notify_notification_proto_rawDesc = nil
	file_eolymp_notify_notification_proto_goTypes = nil
	file_eolymp_notify_notification_proto_depIdxs = nil
}
