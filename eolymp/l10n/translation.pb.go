// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.1
// 	protoc        v4.24.4
// source: eolymp/l10n/translation.proto

package l10n

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

type Translation_Status int32

const (
	Translation_NONE    Translation_Status = 0 // reserved, not in use
	Translation_PENDING Translation_Status = 1 // suggestion pending review
	Translation_ACTIVE  Translation_Status = 2 // active translation
	Translation_UNUSED  Translation_Status = 4 // unused translation
)

// Enum value maps for Translation_Status.
var (
	Translation_Status_name = map[int32]string{
		0: "NONE",
		1: "PENDING",
		2: "ACTIVE",
		4: "UNUSED",
	}
	Translation_Status_value = map[string]int32{
		"NONE":    0,
		"PENDING": 1,
		"ACTIVE":  2,
		"UNUSED":  4,
	}
)

func (x Translation_Status) Enum() *Translation_Status {
	p := new(Translation_Status)
	*p = x
	return p
}

func (x Translation_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Translation_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_eolymp_l10n_translation_proto_enumTypes[0].Descriptor()
}

func (Translation_Status) Type() protoreflect.EnumType {
	return &file_eolymp_l10n_translation_proto_enumTypes[0]
}

func (x Translation_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Translation_Status.Descriptor instead.
func (Translation_Status) EnumDescriptor() ([]byte, []int) {
	return file_eolymp_l10n_translation_proto_rawDescGZIP(), []int{0, 0}
}

type Translation struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Locale        string                 `protobuf:"bytes,2,opt,name=locale,proto3" json:"locale,omitempty"`
	Message       string                 `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
	Status        Translation_Status     `protobuf:"varint,4,opt,name=status,proto3,enum=eolymp.l10n.Translation_Status" json:"status,omitempty"`
	NeedsReview   bool                   `protobuf:"varint,5,opt,name=needs_review,json=needsReview,proto3" json:"needs_review,omitempty"`
	CreatedBy     string                 `protobuf:"bytes,10,opt,name=created_by,json=createdBy,proto3" json:"created_by,omitempty"`
	CreatedAt     *timestamppb.Timestamp `protobuf:"bytes,11,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	ApprovedAt    *timestamppb.Timestamp `protobuf:"bytes,12,opt,name=approved_at,json=approvedAt,proto3" json:"approved_at,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Translation) Reset() {
	*x = Translation{}
	mi := &file_eolymp_l10n_translation_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Translation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Translation) ProtoMessage() {}

func (x *Translation) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_l10n_translation_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Translation.ProtoReflect.Descriptor instead.
func (*Translation) Descriptor() ([]byte, []int) {
	return file_eolymp_l10n_translation_proto_rawDescGZIP(), []int{0}
}

func (x *Translation) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Translation) GetLocale() string {
	if x != nil {
		return x.Locale
	}
	return ""
}

func (x *Translation) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *Translation) GetStatus() Translation_Status {
	if x != nil {
		return x.Status
	}
	return Translation_NONE
}

func (x *Translation) GetNeedsReview() bool {
	if x != nil {
		return x.NeedsReview
	}
	return false
}

func (x *Translation) GetCreatedBy() string {
	if x != nil {
		return x.CreatedBy
	}
	return ""
}

func (x *Translation) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Translation) GetApprovedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.ApprovedAt
	}
	return nil
}

var File_eolymp_l10n_translation_proto protoreflect.FileDescriptor

var file_eolymp_l10n_translation_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x6c, 0x31, 0x30, 0x6e, 0x2f, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0b, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x6c, 0x31, 0x30, 0x6e, 0x1a, 0x1f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xfb, 0x02,
	0x0a, 0x0b, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a,
	0x06, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6c,
	0x6f, 0x63, 0x61, 0x6c, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12,
	0x37, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x1f, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x6c, 0x31, 0x30, 0x6e, 0x2e, 0x54, 0x72,
	0x61, 0x6e, 0x73, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x21, 0x0a, 0x0c, 0x6e, 0x65, 0x65, 0x64,
	0x73, 0x5f, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b,
	0x6e, 0x65, 0x65, 0x64, 0x73, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x12, 0x1d, 0x0a, 0x0a, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x62, 0x79, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x3b, 0x0a, 0x0b, 0x61, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x65,
	0x64, 0x5f, 0x61, 0x74, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x61, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x65, 0x64,
	0x41, 0x74, 0x22, 0x37, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x08, 0x0a, 0x04,
	0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x50, 0x45, 0x4e, 0x44, 0x49, 0x4e,
	0x47, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x41, 0x43, 0x54, 0x49, 0x56, 0x45, 0x10, 0x02, 0x12,
	0x0a, 0x0a, 0x06, 0x55, 0x4e, 0x55, 0x53, 0x45, 0x44, 0x10, 0x04, 0x42, 0x2b, 0x5a, 0x29, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70,
	0x2f, 0x67, 0x6f, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x6c,
	0x31, 0x30, 0x6e, 0x3b, 0x6c, 0x31, 0x30, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_eolymp_l10n_translation_proto_rawDescOnce sync.Once
	file_eolymp_l10n_translation_proto_rawDescData = file_eolymp_l10n_translation_proto_rawDesc
)

func file_eolymp_l10n_translation_proto_rawDescGZIP() []byte {
	file_eolymp_l10n_translation_proto_rawDescOnce.Do(func() {
		file_eolymp_l10n_translation_proto_rawDescData = protoimpl.X.CompressGZIP(file_eolymp_l10n_translation_proto_rawDescData)
	})
	return file_eolymp_l10n_translation_proto_rawDescData
}

var file_eolymp_l10n_translation_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_eolymp_l10n_translation_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_eolymp_l10n_translation_proto_goTypes = []any{
	(Translation_Status)(0),       // 0: eolymp.l10n.Translation.Status
	(*Translation)(nil),           // 1: eolymp.l10n.Translation
	(*timestamppb.Timestamp)(nil), // 2: google.protobuf.Timestamp
}
var file_eolymp_l10n_translation_proto_depIdxs = []int32{
	0, // 0: eolymp.l10n.Translation.status:type_name -> eolymp.l10n.Translation.Status
	2, // 1: eolymp.l10n.Translation.created_at:type_name -> google.protobuf.Timestamp
	2, // 2: eolymp.l10n.Translation.approved_at:type_name -> google.protobuf.Timestamp
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_eolymp_l10n_translation_proto_init() }
func file_eolymp_l10n_translation_proto_init() {
	if File_eolymp_l10n_translation_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_eolymp_l10n_translation_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_eolymp_l10n_translation_proto_goTypes,
		DependencyIndexes: file_eolymp_l10n_translation_proto_depIdxs,
		EnumInfos:         file_eolymp_l10n_translation_proto_enumTypes,
		MessageInfos:      file_eolymp_l10n_translation_proto_msgTypes,
	}.Build()
	File_eolymp_l10n_translation_proto = out.File
	file_eolymp_l10n_translation_proto_rawDesc = nil
	file_eolymp_l10n_translation_proto_goTypes = nil
	file_eolymp_l10n_translation_proto_depIdxs = nil
}
