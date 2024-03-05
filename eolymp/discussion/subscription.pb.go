// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v3.18.1
// source: eolymp/discussion/subscription.proto

package discussion

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

type Subscription_Type int32

const (
	Subscription_NONE       Subscription_Type = 0
	Subscription_MUTED      Subscription_Type = 1
	Subscription_SUBSCRIBED Subscription_Type = 2
	Subscription_WATCHING   Subscription_Type = 3
)

// Enum value maps for Subscription_Type.
var (
	Subscription_Type_name = map[int32]string{
		0: "NONE",
		1: "MUTED",
		2: "SUBSCRIBED",
		3: "WATCHING",
	}
	Subscription_Type_value = map[string]int32{
		"NONE":       0,
		"MUTED":      1,
		"SUBSCRIBED": 2,
		"WATCHING":   3,
	}
)

func (x Subscription_Type) Enum() *Subscription_Type {
	p := new(Subscription_Type)
	*p = x
	return p
}

func (x Subscription_Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Subscription_Type) Descriptor() protoreflect.EnumDescriptor {
	return file_eolymp_discussion_subscription_proto_enumTypes[0].Descriptor()
}

func (Subscription_Type) Type() protoreflect.EnumType {
	return &file_eolymp_discussion_subscription_proto_enumTypes[0]
}

func (x Subscription_Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Subscription_Type.Descriptor instead.
func (Subscription_Type) EnumDescriptor() ([]byte, []int) {
	return file_eolymp_discussion_subscription_proto_rawDescGZIP(), []int{0, 0}
}

type Subscription struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type Subscription_Type `protobuf:"varint,1,opt,name=type,proto3,enum=eolymp.discussion.Subscription_Type" json:"type,omitempty"`
}

func (x *Subscription) Reset() {
	*x = Subscription{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_discussion_subscription_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Subscription) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Subscription) ProtoMessage() {}

func (x *Subscription) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_discussion_subscription_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Subscription.ProtoReflect.Descriptor instead.
func (*Subscription) Descriptor() ([]byte, []int) {
	return file_eolymp_discussion_subscription_proto_rawDescGZIP(), []int{0}
}

func (x *Subscription) GetType() Subscription_Type {
	if x != nil {
		return x.Type
	}
	return Subscription_NONE
}

var File_eolymp_discussion_subscription_proto protoreflect.FileDescriptor

var file_eolymp_discussion_subscription_proto_rawDesc = []byte{
	0x0a, 0x24, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x64, 0x69, 0x73, 0x63, 0x75, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x2f, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x64,
	0x69, 0x73, 0x63, 0x75, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x83, 0x01, 0x0a, 0x0c, 0x53, 0x75,
	0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x38, 0x0a, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x24, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d,
	0x70, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x75, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x53, 0x75, 0x62,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x22, 0x39, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x08, 0x0a, 0x04,
	0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x4d, 0x55, 0x54, 0x45, 0x44, 0x10,
	0x01, 0x12, 0x0e, 0x0a, 0x0a, 0x53, 0x55, 0x42, 0x53, 0x43, 0x52, 0x49, 0x42, 0x45, 0x44, 0x10,
	0x02, 0x12, 0x0c, 0x0a, 0x08, 0x57, 0x41, 0x54, 0x43, 0x48, 0x49, 0x4e, 0x47, 0x10, 0x03, 0x42,
	0x37, 0x5a, 0x35, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6f,
	0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x67, 0x6f, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x65, 0x6f, 0x6c, 0x79,
	0x6d, 0x70, 0x2f, 0x64, 0x69, 0x73, 0x63, 0x75, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x3b, 0x64, 0x69,
	0x73, 0x63, 0x75, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_eolymp_discussion_subscription_proto_rawDescOnce sync.Once
	file_eolymp_discussion_subscription_proto_rawDescData = file_eolymp_discussion_subscription_proto_rawDesc
)

func file_eolymp_discussion_subscription_proto_rawDescGZIP() []byte {
	file_eolymp_discussion_subscription_proto_rawDescOnce.Do(func() {
		file_eolymp_discussion_subscription_proto_rawDescData = protoimpl.X.CompressGZIP(file_eolymp_discussion_subscription_proto_rawDescData)
	})
	return file_eolymp_discussion_subscription_proto_rawDescData
}

var file_eolymp_discussion_subscription_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_eolymp_discussion_subscription_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_eolymp_discussion_subscription_proto_goTypes = []interface{}{
	(Subscription_Type)(0), // 0: eolymp.discussion.Subscription.Type
	(*Subscription)(nil),   // 1: eolymp.discussion.Subscription
}
var file_eolymp_discussion_subscription_proto_depIdxs = []int32{
	0, // 0: eolymp.discussion.Subscription.type:type_name -> eolymp.discussion.Subscription.Type
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_eolymp_discussion_subscription_proto_init() }
func file_eolymp_discussion_subscription_proto_init() {
	if File_eolymp_discussion_subscription_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_eolymp_discussion_subscription_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Subscription); i {
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
			RawDescriptor: file_eolymp_discussion_subscription_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_eolymp_discussion_subscription_proto_goTypes,
		DependencyIndexes: file_eolymp_discussion_subscription_proto_depIdxs,
		EnumInfos:         file_eolymp_discussion_subscription_proto_enumTypes,
		MessageInfos:      file_eolymp_discussion_subscription_proto_msgTypes,
	}.Build()
	File_eolymp_discussion_subscription_proto = out.File
	file_eolymp_discussion_subscription_proto_rawDesc = nil
	file_eolymp_discussion_subscription_proto_goTypes = nil
	file_eolymp_discussion_subscription_proto_depIdxs = nil
}
