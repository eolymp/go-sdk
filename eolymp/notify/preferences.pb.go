// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v4.24.4
// source: eolymp/notify/preferences.proto

package notify

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

type Preferences_Digest int32

const (
	Preferences_UNKNOWN_DIGEST Preferences_Digest = 0
	Preferences_IMMEDIATE      Preferences_Digest = 1
	Preferences_HOURLY         Preferences_Digest = 2
	Preferences_DAILY          Preferences_Digest = 3
)

// Enum value maps for Preferences_Digest.
var (
	Preferences_Digest_name = map[int32]string{
		0: "UNKNOWN_DIGEST",
		1: "IMMEDIATE",
		2: "HOURLY",
		3: "DAILY",
	}
	Preferences_Digest_value = map[string]int32{
		"UNKNOWN_DIGEST": 0,
		"IMMEDIATE":      1,
		"HOURLY":         2,
		"DAILY":          3,
	}
)

func (x Preferences_Digest) Enum() *Preferences_Digest {
	p := new(Preferences_Digest)
	*p = x
	return p
}

func (x Preferences_Digest) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Preferences_Digest) Descriptor() protoreflect.EnumDescriptor {
	return file_eolymp_notify_preferences_proto_enumTypes[0].Descriptor()
}

func (Preferences_Digest) Type() protoreflect.EnumType {
	return &file_eolymp_notify_preferences_proto_enumTypes[0]
}

func (x Preferences_Digest) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Preferences_Digest.Descriptor instead.
func (Preferences_Digest) EnumDescriptor() ([]byte, []int) {
	return file_eolymp_notify_preferences_proto_rawDescGZIP(), []int{0, 0}
}

type Preferences struct {
	state         protoimpl.MessageState      `protogen:"open.v1"`
	Subscriptions []*Preferences_Subscription `protobuf:"bytes,10,rep,name=subscriptions,proto3" json:"subscriptions,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Preferences) Reset() {
	*x = Preferences{}
	mi := &file_eolymp_notify_preferences_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Preferences) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Preferences) ProtoMessage() {}

func (x *Preferences) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_notify_preferences_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Preferences.ProtoReflect.Descriptor instead.
func (*Preferences) Descriptor() ([]byte, []int) {
	return file_eolymp_notify_preferences_proto_rawDescGZIP(), []int{0}
}

func (x *Preferences) GetSubscriptions() []*Preferences_Subscription {
	if x != nil {
		return x.Subscriptions
	}
	return nil
}

type Preferences_Subscription struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Topic         string                 `protobuf:"bytes,1,opt,name=topic,proto3" json:"topic,omitempty"`
	Digest        Preferences_Digest     `protobuf:"varint,2,opt,name=digest,proto3,enum=eolymp.notify.Preferences_Digest" json:"digest,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Preferences_Subscription) Reset() {
	*x = Preferences_Subscription{}
	mi := &file_eolymp_notify_preferences_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Preferences_Subscription) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Preferences_Subscription) ProtoMessage() {}

func (x *Preferences_Subscription) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_notify_preferences_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Preferences_Subscription.ProtoReflect.Descriptor instead.
func (*Preferences_Subscription) Descriptor() ([]byte, []int) {
	return file_eolymp_notify_preferences_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Preferences_Subscription) GetTopic() string {
	if x != nil {
		return x.Topic
	}
	return ""
}

func (x *Preferences_Subscription) GetDigest() Preferences_Digest {
	if x != nil {
		return x.Digest
	}
	return Preferences_UNKNOWN_DIGEST
}

var File_eolymp_notify_preferences_proto protoreflect.FileDescriptor

var file_eolymp_notify_preferences_proto_rawDesc = string([]byte{
	0x0a, 0x1f, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2f,
	0x70, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0d, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x79,
	0x22, 0x81, 0x02, 0x0a, 0x0b, 0x50, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x73,
	0x12, 0x4d, 0x0a, 0x0d, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70,
	0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x50, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e,
	0x63, 0x65, 0x73, 0x2e, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x0d, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x1a,
	0x5f, 0x0a, 0x0c, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x14, 0x0a, 0x05, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x74, 0x6f, 0x70, 0x69, 0x63, 0x12, 0x39, 0x0a, 0x06, 0x64, 0x69, 0x67, 0x65, 0x73, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x21, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x6e,
	0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x50, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65,
	0x73, 0x2e, 0x44, 0x69, 0x67, 0x65, 0x73, 0x74, 0x52, 0x06, 0x64, 0x69, 0x67, 0x65, 0x73, 0x74,
	0x22, 0x42, 0x0a, 0x06, 0x44, 0x69, 0x67, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x0e, 0x55, 0x4e,
	0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x5f, 0x44, 0x49, 0x47, 0x45, 0x53, 0x54, 0x10, 0x00, 0x12, 0x0d,
	0x0a, 0x09, 0x49, 0x4d, 0x4d, 0x45, 0x44, 0x49, 0x41, 0x54, 0x45, 0x10, 0x01, 0x12, 0x0a, 0x0a,
	0x06, 0x48, 0x4f, 0x55, 0x52, 0x4c, 0x59, 0x10, 0x02, 0x12, 0x09, 0x0a, 0x05, 0x44, 0x41, 0x49,
	0x4c, 0x59, 0x10, 0x03, 0x42, 0x2f, 0x5a, 0x2d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x67, 0x6f, 0x2d, 0x73, 0x64, 0x6b,
	0x2f, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x3b, 0x6e,
	0x6f, 0x74, 0x69, 0x66, 0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_eolymp_notify_preferences_proto_rawDescOnce sync.Once
	file_eolymp_notify_preferences_proto_rawDescData []byte
)

func file_eolymp_notify_preferences_proto_rawDescGZIP() []byte {
	file_eolymp_notify_preferences_proto_rawDescOnce.Do(func() {
		file_eolymp_notify_preferences_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_eolymp_notify_preferences_proto_rawDesc), len(file_eolymp_notify_preferences_proto_rawDesc)))
	})
	return file_eolymp_notify_preferences_proto_rawDescData
}

var file_eolymp_notify_preferences_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_eolymp_notify_preferences_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_eolymp_notify_preferences_proto_goTypes = []any{
	(Preferences_Digest)(0),          // 0: eolymp.notify.Preferences.Digest
	(*Preferences)(nil),              // 1: eolymp.notify.Preferences
	(*Preferences_Subscription)(nil), // 2: eolymp.notify.Preferences.Subscription
}
var file_eolymp_notify_preferences_proto_depIdxs = []int32{
	2, // 0: eolymp.notify.Preferences.subscriptions:type_name -> eolymp.notify.Preferences.Subscription
	0, // 1: eolymp.notify.Preferences.Subscription.digest:type_name -> eolymp.notify.Preferences.Digest
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_eolymp_notify_preferences_proto_init() }
func file_eolymp_notify_preferences_proto_init() {
	if File_eolymp_notify_preferences_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_eolymp_notify_preferences_proto_rawDesc), len(file_eolymp_notify_preferences_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_eolymp_notify_preferences_proto_goTypes,
		DependencyIndexes: file_eolymp_notify_preferences_proto_depIdxs,
		EnumInfos:         file_eolymp_notify_preferences_proto_enumTypes,
		MessageInfos:      file_eolymp_notify_preferences_proto_msgTypes,
	}.Build()
	File_eolymp_notify_preferences_proto = out.File
	file_eolymp_notify_preferences_proto_goTypes = nil
	file_eolymp_notify_preferences_proto_depIdxs = nil
}
