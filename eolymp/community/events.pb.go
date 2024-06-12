// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.21.12
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

type MemberChangeRecord_Operation int32

const (
	MemberChangeRecord_NO_OPERATION MemberChangeRecord_Operation = 0
	MemberChangeRecord_CREATE       MemberChangeRecord_Operation = 1
	MemberChangeRecord_UPDATE       MemberChangeRecord_Operation = 2
	MemberChangeRecord_DELETE       MemberChangeRecord_Operation = 3
)

// Enum value maps for MemberChangeRecord_Operation.
var (
	MemberChangeRecord_Operation_name = map[int32]string{
		0: "NO_OPERATION",
		1: "CREATE",
		2: "UPDATE",
		3: "DELETE",
	}
	MemberChangeRecord_Operation_value = map[string]int32{
		"NO_OPERATION": 0,
		"CREATE":       1,
		"UPDATE":       2,
		"DELETE":       3,
	}
)

func (x MemberChangeRecord_Operation) Enum() *MemberChangeRecord_Operation {
	p := new(MemberChangeRecord_Operation)
	*p = x
	return p
}

func (x MemberChangeRecord_Operation) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MemberChangeRecord_Operation) Descriptor() protoreflect.EnumDescriptor {
	return file_eolymp_community_events_proto_enumTypes[0].Descriptor()
}

func (MemberChangeRecord_Operation) Type() protoreflect.EnumType {
	return &file_eolymp_community_events_proto_enumTypes[0]
}

func (x MemberChangeRecord_Operation) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MemberChangeRecord_Operation.Descriptor instead.
func (MemberChangeRecord_Operation) EnumDescriptor() ([]byte, []int) {
	return file_eolymp_community_events_proto_rawDescGZIP(), []int{3, 0}
}

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

type MemberChangeRecord struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SpaceId string                       `protobuf:"bytes,1,opt,name=space_id,json=spaceId,proto3" json:"space_id,omitempty"`
	Op      MemberChangeRecord_Operation `protobuf:"varint,2,opt,name=op,proto3,enum=eolymp.community.MemberChangeRecord_Operation" json:"op,omitempty"`
	Member  *Member                      `protobuf:"bytes,3,opt,name=member,proto3" json:"member,omitempty"`
}

func (x *MemberChangeRecord) Reset() {
	*x = MemberChangeRecord{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_community_events_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MemberChangeRecord) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MemberChangeRecord) ProtoMessage() {}

func (x *MemberChangeRecord) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_community_events_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MemberChangeRecord.ProtoReflect.Descriptor instead.
func (*MemberChangeRecord) Descriptor() ([]byte, []int) {
	return file_eolymp_community_events_proto_rawDescGZIP(), []int{3}
}

func (x *MemberChangeRecord) GetSpaceId() string {
	if x != nil {
		return x.SpaceId
	}
	return ""
}

func (x *MemberChangeRecord) GetOp() MemberChangeRecord_Operation {
	if x != nil {
		return x.Op
	}
	return MemberChangeRecord_NO_OPERATION
}

func (x *MemberChangeRecord) GetMember() *Member {
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
	0x52, 0x06, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x22, 0xe4, 0x01, 0x0a, 0x12, 0x4d, 0x65, 0x6d,
	0x62, 0x65, 0x72, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x12,
	0x19, 0x0a, 0x08, 0x73, 0x70, 0x61, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x73, 0x70, 0x61, 0x63, 0x65, 0x49, 0x64, 0x12, 0x3e, 0x0a, 0x02, 0x6f, 0x70,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2e, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x2e, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72,
	0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x2e, 0x4f, 0x70, 0x65,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x02, 0x6f, 0x70, 0x12, 0x30, 0x0a, 0x06, 0x6d, 0x65,
	0x6d, 0x62, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x65, 0x6f, 0x6c,
	0x79, 0x6d, 0x70, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x2e, 0x4d, 0x65,
	0x6d, 0x62, 0x65, 0x72, 0x52, 0x06, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x22, 0x41, 0x0a, 0x09,
	0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x10, 0x0a, 0x0c, 0x4e, 0x4f, 0x5f,
	0x4f, 0x50, 0x45, 0x52, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x43,
	0x52, 0x45, 0x41, 0x54, 0x45, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x55, 0x50, 0x44, 0x41, 0x54,
	0x45, 0x10, 0x02, 0x12, 0x0a, 0x0a, 0x06, 0x44, 0x45, 0x4c, 0x45, 0x54, 0x45, 0x10, 0x03, 0x42,
	0x35, 0x5a, 0x33, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6f,
	0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x67, 0x6f, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x65, 0x6f, 0x6c, 0x79,
	0x6d, 0x70, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x3b, 0x63, 0x6f, 0x6d,
	0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
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

var file_eolymp_community_events_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_eolymp_community_events_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_eolymp_community_events_proto_goTypes = []any{
	(MemberChangeRecord_Operation)(0), // 0: eolymp.community.MemberChangeRecord.Operation
	(*MemberCreatedEvent)(nil),        // 1: eolymp.community.MemberCreatedEvent
	(*MemberUpdatedEvent)(nil),        // 2: eolymp.community.MemberUpdatedEvent
	(*MemberDeletedEvent)(nil),        // 3: eolymp.community.MemberDeletedEvent
	(*MemberChangeRecord)(nil),        // 4: eolymp.community.MemberChangeRecord
	(*Member)(nil),                    // 5: eolymp.community.Member
}
var file_eolymp_community_events_proto_depIdxs = []int32{
	5, // 0: eolymp.community.MemberCreatedEvent.member:type_name -> eolymp.community.Member
	5, // 1: eolymp.community.MemberUpdatedEvent.member:type_name -> eolymp.community.Member
	5, // 2: eolymp.community.MemberDeletedEvent.member:type_name -> eolymp.community.Member
	0, // 3: eolymp.community.MemberChangeRecord.op:type_name -> eolymp.community.MemberChangeRecord.Operation
	5, // 4: eolymp.community.MemberChangeRecord.member:type_name -> eolymp.community.Member
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_eolymp_community_events_proto_init() }
func file_eolymp_community_events_proto_init() {
	if File_eolymp_community_events_proto != nil {
		return
	}
	file_eolymp_community_member_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_eolymp_community_events_proto_msgTypes[0].Exporter = func(v any, i int) any {
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
		file_eolymp_community_events_proto_msgTypes[1].Exporter = func(v any, i int) any {
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
		file_eolymp_community_events_proto_msgTypes[2].Exporter = func(v any, i int) any {
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
		file_eolymp_community_events_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*MemberChangeRecord); i {
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
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_eolymp_community_events_proto_goTypes,
		DependencyIndexes: file_eolymp_community_events_proto_depIdxs,
		EnumInfos:         file_eolymp_community_events_proto_enumTypes,
		MessageInfos:      file_eolymp_community_events_proto_msgTypes,
	}.Build()
	File_eolymp_community_events_proto = out.File
	file_eolymp_community_events_proto_rawDesc = nil
	file_eolymp_community_events_proto_goTypes = nil
	file_eolymp_community_events_proto_depIdxs = nil
}
