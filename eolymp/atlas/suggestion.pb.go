// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v3.21.12
// source: eolymp/atlas/suggestion.proto

package atlas

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

type Suggestion_Status int32

const (
	Suggestion_UNKNOWN_STATUS Suggestion_Status = 0
	Suggestion_PENDING        Suggestion_Status = 1
	Suggestion_IN_REVIEW      Suggestion_Status = 2
	Suggestion_APPROVED       Suggestion_Status = 3
	Suggestion_REJECTED       Suggestion_Status = 4
)

// Enum value maps for Suggestion_Status.
var (
	Suggestion_Status_name = map[int32]string{
		0: "UNKNOWN_STATUS",
		1: "PENDING",
		2: "IN_REVIEW",
		3: "APPROVED",
		4: "REJECTED",
	}
	Suggestion_Status_value = map[string]int32{
		"UNKNOWN_STATUS": 0,
		"PENDING":        1,
		"IN_REVIEW":      2,
		"APPROVED":       3,
		"REJECTED":       4,
	}
)

func (x Suggestion_Status) Enum() *Suggestion_Status {
	p := new(Suggestion_Status)
	*p = x
	return p
}

func (x Suggestion_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Suggestion_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_eolymp_atlas_suggestion_proto_enumTypes[0].Descriptor()
}

func (Suggestion_Status) Type() protoreflect.EnumType {
	return &file_eolymp_atlas_suggestion_proto_enumTypes[0]
}

func (x Suggestion_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Suggestion_Status.Descriptor instead.
func (Suggestion_Status) EnumDescriptor() ([]byte, []int) {
	return file_eolymp_atlas_suggestion_proto_rawDescGZIP(), []int{0, 0}
}

type Suggestion struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Status     Suggestion_Status      `protobuf:"varint,2,opt,name=status,proto3,enum=eolymp.atlas.Suggestion_Status" json:"status,omitempty"`
	Locale     string                 `protobuf:"bytes,4,opt,name=locale,proto3" json:"locale,omitempty"`
	Title      string                 `protobuf:"bytes,6,opt,name=title,proto3" json:"title,omitempty"`
	MemberId   string                 `protobuf:"bytes,5,opt,name=member_id,json=memberId,proto3" json:"member_id,omitempty"`
	CreatedAt  *timestamppb.Timestamp `protobuf:"bytes,10,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt  *timestamppb.Timestamp `protobuf:"bytes,11,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	Topics     []string               `protobuf:"bytes,100,rep,name=topics,proto3" json:"topics,omitempty"`
	Difficulty uint32                 `protobuf:"varint,101,opt,name=difficulty,proto3" json:"difficulty,omitempty"`
	Statement  *ecm.Content           `protobuf:"bytes,102,opt,name=statement,proto3" json:"statement,omitempty"`
	Editorial  *ecm.Content           `protobuf:"bytes,103,opt,name=editorial,proto3" json:"editorial,omitempty"`
}

func (x *Suggestion) Reset() {
	*x = Suggestion{}
	mi := &file_eolymp_atlas_suggestion_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Suggestion) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Suggestion) ProtoMessage() {}

func (x *Suggestion) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_atlas_suggestion_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Suggestion.ProtoReflect.Descriptor instead.
func (*Suggestion) Descriptor() ([]byte, []int) {
	return file_eolymp_atlas_suggestion_proto_rawDescGZIP(), []int{0}
}

func (x *Suggestion) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Suggestion) GetStatus() Suggestion_Status {
	if x != nil {
		return x.Status
	}
	return Suggestion_UNKNOWN_STATUS
}

func (x *Suggestion) GetLocale() string {
	if x != nil {
		return x.Locale
	}
	return ""
}

func (x *Suggestion) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Suggestion) GetMemberId() string {
	if x != nil {
		return x.MemberId
	}
	return ""
}

func (x *Suggestion) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Suggestion) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *Suggestion) GetTopics() []string {
	if x != nil {
		return x.Topics
	}
	return nil
}

func (x *Suggestion) GetDifficulty() uint32 {
	if x != nil {
		return x.Difficulty
	}
	return 0
}

func (x *Suggestion) GetStatement() *ecm.Content {
	if x != nil {
		return x.Statement
	}
	return nil
}

func (x *Suggestion) GetEditorial() *ecm.Content {
	if x != nil {
		return x.Editorial
	}
	return nil
}

var File_eolymp_atlas_suggestion_proto protoreflect.FileDescriptor

var file_eolymp_atlas_suggestion_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x61, 0x74, 0x6c, 0x61, 0x73, 0x2f, 0x73,
	0x75, 0x67, 0x67, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0c, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x61, 0x74, 0x6c, 0x61, 0x73, 0x1a, 0x18, 0x65,
	0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x65, 0x63, 0x6d, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8a, 0x04, 0x0a, 0x0a, 0x53, 0x75, 0x67,
	0x67, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x37, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1f, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70,
	0x2e, 0x61, 0x74, 0x6c, 0x61, 0x73, 0x2e, 0x53, 0x75, 0x67, 0x67, 0x65, 0x73, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x16, 0x0a, 0x06, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c,
	0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x1b,
	0x0a, 0x09, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x49, 0x64, 0x12, 0x39, 0x0a, 0x0a, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x64, 0x5f, 0x61, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x73, 0x18, 0x64, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x06, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x64, 0x69, 0x66,
	0x66, 0x69, 0x63, 0x75, 0x6c, 0x74, 0x79, 0x18, 0x65, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x64,
	0x69, 0x66, 0x66, 0x69, 0x63, 0x75, 0x6c, 0x74, 0x79, 0x12, 0x31, 0x0a, 0x09, 0x73, 0x74, 0x61,
	0x74, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x66, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x65,
	0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x65, 0x63, 0x6d, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x52, 0x09, 0x73, 0x74, 0x61, 0x74, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x31, 0x0a, 0x09,
	0x65, 0x64, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x61, 0x6c, 0x18, 0x67, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x13, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x65, 0x63, 0x6d, 0x2e, 0x43, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x52, 0x09, 0x65, 0x64, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x61, 0x6c, 0x22,
	0x54, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x12, 0x0a, 0x0e, 0x55, 0x4e, 0x4b,
	0x4e, 0x4f, 0x57, 0x4e, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x10, 0x00, 0x12, 0x0b, 0x0a,
	0x07, 0x50, 0x45, 0x4e, 0x44, 0x49, 0x4e, 0x47, 0x10, 0x01, 0x12, 0x0d, 0x0a, 0x09, 0x49, 0x4e,
	0x5f, 0x52, 0x45, 0x56, 0x49, 0x45, 0x57, 0x10, 0x02, 0x12, 0x0c, 0x0a, 0x08, 0x41, 0x50, 0x50,
	0x52, 0x4f, 0x56, 0x45, 0x44, 0x10, 0x03, 0x12, 0x0c, 0x0a, 0x08, 0x52, 0x45, 0x4a, 0x45, 0x43,
	0x54, 0x45, 0x44, 0x10, 0x04, 0x42, 0x2d, 0x5a, 0x2b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x67, 0x6f, 0x2d, 0x73, 0x64,
	0x6b, 0x2f, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x61, 0x74, 0x6c, 0x61, 0x73, 0x3b, 0x61,
	0x74, 0x6c, 0x61, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_eolymp_atlas_suggestion_proto_rawDescOnce sync.Once
	file_eolymp_atlas_suggestion_proto_rawDescData = file_eolymp_atlas_suggestion_proto_rawDesc
)

func file_eolymp_atlas_suggestion_proto_rawDescGZIP() []byte {
	file_eolymp_atlas_suggestion_proto_rawDescOnce.Do(func() {
		file_eolymp_atlas_suggestion_proto_rawDescData = protoimpl.X.CompressGZIP(file_eolymp_atlas_suggestion_proto_rawDescData)
	})
	return file_eolymp_atlas_suggestion_proto_rawDescData
}

var file_eolymp_atlas_suggestion_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_eolymp_atlas_suggestion_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_eolymp_atlas_suggestion_proto_goTypes = []any{
	(Suggestion_Status)(0),        // 0: eolymp.atlas.Suggestion.Status
	(*Suggestion)(nil),            // 1: eolymp.atlas.Suggestion
	(*timestamppb.Timestamp)(nil), // 2: google.protobuf.Timestamp
	(*ecm.Content)(nil),           // 3: eolymp.ecm.Content
}
var file_eolymp_atlas_suggestion_proto_depIdxs = []int32{
	0, // 0: eolymp.atlas.Suggestion.status:type_name -> eolymp.atlas.Suggestion.Status
	2, // 1: eolymp.atlas.Suggestion.created_at:type_name -> google.protobuf.Timestamp
	2, // 2: eolymp.atlas.Suggestion.updated_at:type_name -> google.protobuf.Timestamp
	3, // 3: eolymp.atlas.Suggestion.statement:type_name -> eolymp.ecm.Content
	3, // 4: eolymp.atlas.Suggestion.editorial:type_name -> eolymp.ecm.Content
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_eolymp_atlas_suggestion_proto_init() }
func file_eolymp_atlas_suggestion_proto_init() {
	if File_eolymp_atlas_suggestion_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_eolymp_atlas_suggestion_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_eolymp_atlas_suggestion_proto_goTypes,
		DependencyIndexes: file_eolymp_atlas_suggestion_proto_depIdxs,
		EnumInfos:         file_eolymp_atlas_suggestion_proto_enumTypes,
		MessageInfos:      file_eolymp_atlas_suggestion_proto_msgTypes,
	}.Build()
	File_eolymp_atlas_suggestion_proto = out.File
	file_eolymp_atlas_suggestion_proto_rawDesc = nil
	file_eolymp_atlas_suggestion_proto_goTypes = nil
	file_eolymp_atlas_suggestion_proto_depIdxs = nil
}
