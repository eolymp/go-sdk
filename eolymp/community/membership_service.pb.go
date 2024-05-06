// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.0
// 	protoc        v3.21.12
// source: eolymp/community/membership_service.proto

package community

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

type DescribeMembershipInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DescribeMembershipInput) Reset() {
	*x = DescribeMembershipInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_community_membership_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DescribeMembershipInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescribeMembershipInput) ProtoMessage() {}

func (x *DescribeMembershipInput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_community_membership_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DescribeMembershipInput.ProtoReflect.Descriptor instead.
func (*DescribeMembershipInput) Descriptor() ([]byte, []int) {
	return file_eolymp_community_membership_service_proto_rawDescGZIP(), []int{0}
}

type DescribeMembershipOutput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Member *Member `protobuf:"bytes,1,opt,name=member,proto3" json:"member,omitempty"`
}

func (x *DescribeMembershipOutput) Reset() {
	*x = DescribeMembershipOutput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_community_membership_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DescribeMembershipOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescribeMembershipOutput) ProtoMessage() {}

func (x *DescribeMembershipOutput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_community_membership_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DescribeMembershipOutput.ProtoReflect.Descriptor instead.
func (*DescribeMembershipOutput) Descriptor() ([]byte, []int) {
	return file_eolymp_community_membership_service_proto_rawDescGZIP(), []int{1}
}

func (x *DescribeMembershipOutput) GetMember() *Member {
	if x != nil {
		return x.Member
	}
	return nil
}

type UpdateMembershipInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Attributes []*Attribute_Value `protobuf:"bytes,1,rep,name=attributes,proto3" json:"attributes,omitempty"`
}

func (x *UpdateMembershipInput) Reset() {
	*x = UpdateMembershipInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_community_membership_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateMembershipInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateMembershipInput) ProtoMessage() {}

func (x *UpdateMembershipInput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_community_membership_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateMembershipInput.ProtoReflect.Descriptor instead.
func (*UpdateMembershipInput) Descriptor() ([]byte, []int) {
	return file_eolymp_community_membership_service_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateMembershipInput) GetAttributes() []*Attribute_Value {
	if x != nil {
		return x.Attributes
	}
	return nil
}

type UpdateMembershipOutput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateMembershipOutput) Reset() {
	*x = UpdateMembershipOutput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_community_membership_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateMembershipOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateMembershipOutput) ProtoMessage() {}

func (x *UpdateMembershipOutput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_community_membership_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateMembershipOutput.ProtoReflect.Descriptor instead.
func (*UpdateMembershipOutput) Descriptor() ([]byte, []int) {
	return file_eolymp_community_membership_service_proto_rawDescGZIP(), []int{3}
}

var File_eolymp_community_membership_service_proto protoreflect.FileDescriptor

var file_eolymp_community_membership_service_proto_rawDesc = []byte{
	0x0a, 0x29, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69,
	0x74, 0x79, 0x2f, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x68, 0x69, 0x70, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x65, 0x6f, 0x6c,
	0x79, 0x6d, 0x70, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x1a, 0x1d, 0x65,
	0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2f, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x22, 0x65, 0x6f,
	0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2f, 0x72, 0x61, 0x74, 0x65, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x20, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69,
	0x74, 0x79, 0x2f, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1d, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x75,
	0x6e, 0x69, 0x74, 0x79, 0x2f, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x19, 0x0a, 0x17, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x4d, 0x65, 0x6d,
	0x62, 0x65, 0x72, 0x73, 0x68, 0x69, 0x70, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x22, 0x4c, 0x0a, 0x18,
	0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x68,
	0x69, 0x70, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x12, 0x30, 0x0a, 0x06, 0x6d, 0x65, 0x6d, 0x62,
	0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d,
	0x70, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x2e, 0x4d, 0x65, 0x6d, 0x62,
	0x65, 0x72, 0x52, 0x06, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x22, 0x5a, 0x0a, 0x15, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x68, 0x69, 0x70, 0x49, 0x6e,
	0x70, 0x75, 0x74, 0x12, 0x41, 0x0a, 0x0a, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x2e, 0x41, 0x74, 0x74, 0x72, 0x69,
	0x62, 0x75, 0x74, 0x65, 0x2e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0a, 0x61, 0x74, 0x74, 0x72,
	0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x22, 0x18, 0x0a, 0x16, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x68, 0x69, 0x70, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74,
	0x32, 0xe0, 0x02, 0x0a, 0x11, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x68, 0x69, 0x70, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x92, 0x01, 0x0a, 0x12, 0x44, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x62, 0x65, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x68, 0x69, 0x70, 0x12, 0x29, 0x2e,
	0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79,
	0x2e, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73,
	0x68, 0x69, 0x70, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x2a, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d,
	0x70, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x2e, 0x44, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x62, 0x65, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x68, 0x69, 0x70, 0x4f, 0x75,
	0x74, 0x70, 0x75, 0x74, 0x22, 0x25, 0xea, 0xe2, 0x0a, 0x0b, 0xf5, 0xe2, 0x0a, 0x00, 0x00, 0x20,
	0x41, 0xf8, 0xe2, 0x0a, 0x32, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x10, 0x12, 0x0e, 0x2f, 0x6d, 0x65,
	0x6d, 0x62, 0x65, 0x72, 0x73, 0x2f, 0x5f, 0x73, 0x65, 0x6c, 0x66, 0x12, 0xb5, 0x01, 0x0a, 0x10,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x68, 0x69, 0x70,
	0x12, 0x27, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e,
	0x69, 0x74, 0x79, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72,
	0x73, 0x68, 0x69, 0x70, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x28, 0x2e, 0x65, 0x6f, 0x6c, 0x79,
	0x6d, 0x70, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x2e, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x68, 0x69, 0x70, 0x4f, 0x75, 0x74,
	0x70, 0x75, 0x74, 0x22, 0x4e, 0xea, 0xe2, 0x0a, 0x0b, 0xf5, 0xe2, 0x0a, 0x00, 0x00, 0xa0, 0x40,
	0xf8, 0xe2, 0x0a, 0x14, 0x82, 0xe3, 0x0a, 0x1a, 0x8a, 0xe3, 0x0a, 0x16, 0x63, 0x6f, 0x6d, 0x6d,
	0x75, 0x6e, 0x69, 0x74, 0x79, 0x3a, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x3a, 0x77, 0x72, 0x69,
	0x74, 0x65, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1b, 0x22, 0x19, 0x2f, 0x6d, 0x65, 0x6d, 0x62, 0x65,
	0x72, 0x73, 0x2f, 0x5f, 0x73, 0x65, 0x6c, 0x66, 0x2f, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75,
	0x74, 0x65, 0x73, 0x42, 0x35, 0x5a, 0x33, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x67, 0x6f, 0x2d, 0x73, 0x64, 0x6b, 0x2f,
	0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79,
	0x3b, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_eolymp_community_membership_service_proto_rawDescOnce sync.Once
	file_eolymp_community_membership_service_proto_rawDescData = file_eolymp_community_membership_service_proto_rawDesc
)

func file_eolymp_community_membership_service_proto_rawDescGZIP() []byte {
	file_eolymp_community_membership_service_proto_rawDescOnce.Do(func() {
		file_eolymp_community_membership_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_eolymp_community_membership_service_proto_rawDescData)
	})
	return file_eolymp_community_membership_service_proto_rawDescData
}

var file_eolymp_community_membership_service_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_eolymp_community_membership_service_proto_goTypes = []interface{}{
	(*DescribeMembershipInput)(nil),  // 0: eolymp.community.DescribeMembershipInput
	(*DescribeMembershipOutput)(nil), // 1: eolymp.community.DescribeMembershipOutput
	(*UpdateMembershipInput)(nil),    // 2: eolymp.community.UpdateMembershipInput
	(*UpdateMembershipOutput)(nil),   // 3: eolymp.community.UpdateMembershipOutput
	(*Member)(nil),                   // 4: eolymp.community.Member
	(*Attribute_Value)(nil),          // 5: eolymp.community.Attribute.Value
}
var file_eolymp_community_membership_service_proto_depIdxs = []int32{
	4, // 0: eolymp.community.DescribeMembershipOutput.member:type_name -> eolymp.community.Member
	5, // 1: eolymp.community.UpdateMembershipInput.attributes:type_name -> eolymp.community.Attribute.Value
	0, // 2: eolymp.community.MembershipService.DescribeMembership:input_type -> eolymp.community.DescribeMembershipInput
	2, // 3: eolymp.community.MembershipService.UpdateMembership:input_type -> eolymp.community.UpdateMembershipInput
	1, // 4: eolymp.community.MembershipService.DescribeMembership:output_type -> eolymp.community.DescribeMembershipOutput
	3, // 5: eolymp.community.MembershipService.UpdateMembership:output_type -> eolymp.community.UpdateMembershipOutput
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_eolymp_community_membership_service_proto_init() }
func file_eolymp_community_membership_service_proto_init() {
	if File_eolymp_community_membership_service_proto != nil {
		return
	}
	file_eolymp_community_attribute_proto_init()
	file_eolymp_community_member_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_eolymp_community_membership_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DescribeMembershipInput); i {
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
		file_eolymp_community_membership_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DescribeMembershipOutput); i {
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
		file_eolymp_community_membership_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateMembershipInput); i {
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
		file_eolymp_community_membership_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateMembershipOutput); i {
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
			RawDescriptor: file_eolymp_community_membership_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_eolymp_community_membership_service_proto_goTypes,
		DependencyIndexes: file_eolymp_community_membership_service_proto_depIdxs,
		MessageInfos:      file_eolymp_community_membership_service_proto_msgTypes,
	}.Build()
	File_eolymp_community_membership_service_proto = out.File
	file_eolymp_community_membership_service_proto_rawDesc = nil
	file_eolymp_community_membership_service_proto_goTypes = nil
	file_eolymp_community_membership_service_proto_depIdxs = nil
}
