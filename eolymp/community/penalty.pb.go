// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.3
// 	protoc        v4.24.4
// source: eolymp/community/penalty.proto

package community

import (
	ecm "github.com/eolymp/go-sdk/eolymp/ecm"
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

type Penalty_Extra int32

const (
	Penalty_NO_EXTRA           Penalty_Extra = 0
	Penalty_DESCRIPTION_VALUE  Penalty_Extra = 1
	Penalty_DESCRIPTION_RENDER Penalty_Extra = 2
)

// Enum value maps for Penalty_Extra.
var (
	Penalty_Extra_name = map[int32]string{
		0: "NO_EXTRA",
		1: "DESCRIPTION_VALUE",
		2: "DESCRIPTION_RENDER",
	}
	Penalty_Extra_value = map[string]int32{
		"NO_EXTRA":           0,
		"DESCRIPTION_VALUE":  1,
		"DESCRIPTION_RENDER": 2,
	}
)

func (x Penalty_Extra) Enum() *Penalty_Extra {
	p := new(Penalty_Extra)
	*p = x
	return p
}

func (x Penalty_Extra) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Penalty_Extra) Descriptor() protoreflect.EnumDescriptor {
	return file_eolymp_community_penalty_proto_enumTypes[0].Descriptor()
}

func (Penalty_Extra) Type() protoreflect.EnumType {
	return &file_eolymp_community_penalty_proto_enumTypes[0]
}

func (x Penalty_Extra) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Penalty_Extra.Descriptor instead.
func (Penalty_Extra) EnumDescriptor() ([]byte, []int) {
	return file_eolymp_community_penalty_proto_rawDescGZIP(), []int{0, 0}
}

type Penalty struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Summary       string                 `protobuf:"bytes,2,opt,name=summary,proto3" json:"summary,omitempty"`
	Description   *ecm.Content           `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Scope         []string               `protobuf:"bytes,20,rep,name=scope,proto3" json:"scope,omitempty"`
	CreatedAt     *timestamppb.Timestamp `protobuf:"bytes,10,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	ExpiresAt     *timestamppb.Timestamp `protobuf:"bytes,11,opt,name=expires_at,json=expiresAt,proto3" json:"expires_at,omitempty"`
	CancelledAt   *timestamppb.Timestamp `protobuf:"bytes,12,opt,name=cancelled_at,json=cancelledAt,proto3" json:"cancelled_at,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Penalty) Reset() {
	*x = Penalty{}
	mi := &file_eolymp_community_penalty_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Penalty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Penalty) ProtoMessage() {}

func (x *Penalty) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_community_penalty_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Penalty.ProtoReflect.Descriptor instead.
func (*Penalty) Descriptor() ([]byte, []int) {
	return file_eolymp_community_penalty_proto_rawDescGZIP(), []int{0}
}

func (x *Penalty) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Penalty) GetSummary() string {
	if x != nil {
		return x.Summary
	}
	return ""
}

func (x *Penalty) GetDescription() *ecm.Content {
	if x != nil {
		return x.Description
	}
	return nil
}

func (x *Penalty) GetScope() []string {
	if x != nil {
		return x.Scope
	}
	return nil
}

func (x *Penalty) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Penalty) GetExpiresAt() *timestamppb.Timestamp {
	if x != nil {
		return x.ExpiresAt
	}
	return nil
}

func (x *Penalty) GetCancelledAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CancelledAt
	}
	return nil
}

var File_eolymp_community_penalty_proto protoreflect.FileDescriptor

var file_eolymp_community_penalty_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69,
	0x74, 0x79, 0x2f, 0x70, 0x65, 0x6e, 0x61, 0x6c, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x10, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69,
	0x74, 0x79, 0x1a, 0x18, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x65, 0x63, 0x6d, 0x2f, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xfb, 0x02,
	0x0a, 0x07, 0x50, 0x65, 0x6e, 0x61, 0x6c, 0x74, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x6d,
	0x6d, 0x61, 0x72, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x75, 0x6d, 0x6d,
	0x61, 0x72, 0x79, 0x12, 0x35, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d,
	0x70, 0x2e, 0x65, 0x63, 0x6d, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x52, 0x0b, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x63,
	0x6f, 0x70, 0x65, 0x18, 0x14, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x73, 0x63, 0x6f, 0x70, 0x65,
	0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0a,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x65,
	0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x5f, 0x61, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x65, 0x78, 0x70,
	0x69, 0x72, 0x65, 0x73, 0x41, 0x74, 0x12, 0x3d, 0x0a, 0x0c, 0x63, 0x61, 0x6e, 0x63, 0x65, 0x6c,
	0x6c, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0b, 0x63, 0x61, 0x6e, 0x63, 0x65, 0x6c,
	0x6c, 0x65, 0x64, 0x41, 0x74, 0x22, 0x44, 0x0a, 0x05, 0x45, 0x78, 0x74, 0x72, 0x61, 0x12, 0x0c,
	0x0a, 0x08, 0x4e, 0x4f, 0x5f, 0x45, 0x58, 0x54, 0x52, 0x41, 0x10, 0x00, 0x12, 0x15, 0x0a, 0x11,
	0x44, 0x45, 0x53, 0x43, 0x52, 0x49, 0x50, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x56, 0x41, 0x4c, 0x55,
	0x45, 0x10, 0x01, 0x12, 0x16, 0x0a, 0x12, 0x44, 0x45, 0x53, 0x43, 0x52, 0x49, 0x50, 0x54, 0x49,
	0x4f, 0x4e, 0x5f, 0x52, 0x45, 0x4e, 0x44, 0x45, 0x52, 0x10, 0x02, 0x42, 0x35, 0x5a, 0x33, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70,
	0x2f, 0x67, 0x6f, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x63,
	0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x3b, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69,
	0x74, 0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_eolymp_community_penalty_proto_rawDescOnce sync.Once
	file_eolymp_community_penalty_proto_rawDescData = file_eolymp_community_penalty_proto_rawDesc
)

func file_eolymp_community_penalty_proto_rawDescGZIP() []byte {
	file_eolymp_community_penalty_proto_rawDescOnce.Do(func() {
		file_eolymp_community_penalty_proto_rawDescData = protoimpl.X.CompressGZIP(file_eolymp_community_penalty_proto_rawDescData)
	})
	return file_eolymp_community_penalty_proto_rawDescData
}

var file_eolymp_community_penalty_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_eolymp_community_penalty_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_eolymp_community_penalty_proto_goTypes = []any{
	(Penalty_Extra)(0),            // 0: eolymp.community.Penalty.Extra
	(*Penalty)(nil),               // 1: eolymp.community.Penalty
	(*ecm.Content)(nil),           // 2: eolymp.ecm.Content
	(*timestamppb.Timestamp)(nil), // 3: google.protobuf.Timestamp
}
var file_eolymp_community_penalty_proto_depIdxs = []int32{
	2, // 0: eolymp.community.Penalty.description:type_name -> eolymp.ecm.Content
	3, // 1: eolymp.community.Penalty.created_at:type_name -> google.protobuf.Timestamp
	3, // 2: eolymp.community.Penalty.expires_at:type_name -> google.protobuf.Timestamp
	3, // 3: eolymp.community.Penalty.cancelled_at:type_name -> google.protobuf.Timestamp
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_eolymp_community_penalty_proto_init() }
func file_eolymp_community_penalty_proto_init() {
	if File_eolymp_community_penalty_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_eolymp_community_penalty_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_eolymp_community_penalty_proto_goTypes,
		DependencyIndexes: file_eolymp_community_penalty_proto_depIdxs,
		EnumInfos:         file_eolymp_community_penalty_proto_enumTypes,
		MessageInfos:      file_eolymp_community_penalty_proto_msgTypes,
	}.Build()
	File_eolymp_community_penalty_proto = out.File
	file_eolymp_community_penalty_proto_rawDesc = nil
	file_eolymp_community_penalty_proto_goTypes = nil
	file_eolymp_community_penalty_proto_depIdxs = nil
}
