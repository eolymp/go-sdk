// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.18.1
// source: eolymp/community/events.proto

package community

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

type MemberCreatedEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Member *Member `protobuf:"bytes,1,opt,name=member,proto3" json:"member,omitempty"`
}

func (x *MemberCreatedEvent) Reset() {
	*x = MemberCreatedEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_community_events_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MemberCreatedEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MemberCreatedEvent) ProtoMessage() {}

func (x *MemberCreatedEvent) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_community_events_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MemberCreatedEvent.ProtoReflect.Descriptor instead.
func (*MemberCreatedEvent) Descriptor() ([]byte, []int) {
	return file_eolymp_community_events_proto_rawDescGZIP(), []int{0}
}

func (x *MemberCreatedEvent) GetMember() *Member {
	if x != nil {
		return x.Member
	}
	return nil
}

type MemberUpdatedEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Member *Member `protobuf:"bytes,1,opt,name=member,proto3" json:"member,omitempty"`
}

func (x *MemberUpdatedEvent) Reset() {
	*x = MemberUpdatedEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_community_events_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MemberUpdatedEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MemberUpdatedEvent) ProtoMessage() {}

func (x *MemberUpdatedEvent) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_community_events_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MemberUpdatedEvent.ProtoReflect.Descriptor instead.
func (*MemberUpdatedEvent) Descriptor() ([]byte, []int) {
	return file_eolymp_community_events_proto_rawDescGZIP(), []int{1}
}

func (x *MemberUpdatedEvent) GetMember() *Member {
	if x != nil {
		return x.Member
	}
	return nil
}

type MemberDeletedEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Member *Member `protobuf:"bytes,1,opt,name=member,proto3" json:"member,omitempty"`
}

func (x *MemberDeletedEvent) Reset() {
	*x = MemberDeletedEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_community_events_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MemberDeletedEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MemberDeletedEvent) ProtoMessage() {}

func (x *MemberDeletedEvent) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_community_events_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MemberDeletedEvent.ProtoReflect.Descriptor instead.
func (*MemberDeletedEvent) Descriptor() ([]byte, []int) {
	return file_eolymp_community_events_proto_rawDescGZIP(), []int{2}
}

func (x *MemberDeletedEvent) GetMember() *Member {
	if x != nil {
		return x.Member
	}
	return nil
}

var File_eolymp_community_events_proto protoreflect.FileDescriptor

var file_eolymp_community_events_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69,
	0x74, 0x79, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x10, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74,
	0x79, 0x1a, 0x1d, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e,
	0x69, 0x74, 0x79, 0x2f, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x46, 0x0a, 0x12, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x30, 0x0a, 0x06, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x2e, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72,
	0x52, 0x06, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x22, 0x46, 0x0a, 0x12, 0x4d, 0x65, 0x6d, 0x62,
	0x65, 0x72, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x30,
	0x0a, 0x06, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18,
	0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74,
	0x79, 0x2e, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x06, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72,
	0x22, 0x46, 0x0a, 0x12, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x64, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x30, 0x0a, 0x06, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x2e, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72,
	0x52, 0x06, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x42, 0x35, 0x5a, 0x33, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x67, 0x6f,
	0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x63, 0x6f, 0x6d, 0x6d,
	0x75, 0x6e, 0x69, 0x74, 0x79, 0x3b, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_eolymp_community_events_proto_rawDescOnce sync.Once
	file_eolymp_community_events_proto_rawDescData = file_eolymp_community_events_proto_rawDesc
)

func file_eolymp_community_events_proto_rawDescGZIP() []byte {
	file_eolymp_community_events_proto_rawDescOnce.Do(func() {
		file_eolymp_community_events_proto_rawDescData = protoimpl.X.CompressGZIP(file_eolymp_community_events_proto_rawDescData)
	})
	return file_eolymp_community_events_proto_rawDescData
}

var file_eolymp_community_events_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_eolymp_community_events_proto_goTypes = []interface{}{
	(*MemberCreatedEvent)(nil), // 0: eolymp.community.MemberCreatedEvent
	(*MemberUpdatedEvent)(nil), // 1: eolymp.community.MemberUpdatedEvent
	(*MemberDeletedEvent)(nil), // 2: eolymp.community.MemberDeletedEvent
	(*Member)(nil),             // 3: eolymp.community.Member
}
var file_eolymp_community_events_proto_depIdxs = []int32{
	3, // 0: eolymp.community.MemberCreatedEvent.member:type_name -> eolymp.community.Member
	3, // 1: eolymp.community.MemberUpdatedEvent.member:type_name -> eolymp.community.Member
	3, // 2: eolymp.community.MemberDeletedEvent.member:type_name -> eolymp.community.Member
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_eolymp_community_events_proto_init() }
func file_eolymp_community_events_proto_init() {
	if File_eolymp_community_events_proto != nil {
		return
	}
	file_eolymp_community_member_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_eolymp_community_events_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MemberCreatedEvent); i {
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
		file_eolymp_community_events_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MemberUpdatedEvent); i {
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
		file_eolymp_community_events_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MemberDeletedEvent); i {
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
			RawDescriptor: file_eolymp_community_events_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_eolymp_community_events_proto_goTypes,
		DependencyIndexes: file_eolymp_community_events_proto_depIdxs,
		MessageInfos:      file_eolymp_community_events_proto_msgTypes,
	}.Build()
	File_eolymp_community_events_proto = out.File
	file_eolymp_community_events_proto_rawDesc = nil
	file_eolymp_community_events_proto_goTypes = nil
	file_eolymp_community_events_proto_depIdxs = nil
}
