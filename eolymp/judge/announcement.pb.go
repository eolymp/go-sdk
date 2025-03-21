// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v4.24.4
// source: eolymp/judge/announcement.proto

package judge

import (
	ecm "github.com/eolymp/go-sdk/eolymp/ecm"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
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

type Announcement_Extra int32

const (
	Announcement_NO_EXTRA       Announcement_Extra = 0
	Announcement_MESSAGE_RENDER Announcement_Extra = 1
	Announcement_MESSAGE_VALUE  Announcement_Extra = 2
)

// Enum value maps for Announcement_Extra.
var (
	Announcement_Extra_name = map[int32]string{
		0: "NO_EXTRA",
		1: "MESSAGE_RENDER",
		2: "MESSAGE_VALUE",
	}
	Announcement_Extra_value = map[string]int32{
		"NO_EXTRA":       0,
		"MESSAGE_RENDER": 1,
		"MESSAGE_VALUE":  2,
	}
)

func (x Announcement_Extra) Enum() *Announcement_Extra {
	p := new(Announcement_Extra)
	*p = x
	return p
}

func (x Announcement_Extra) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Announcement_Extra) Descriptor() protoreflect.EnumDescriptor {
	return file_eolymp_judge_announcement_proto_enumTypes[0].Descriptor()
}

func (Announcement_Extra) Type() protoreflect.EnumType {
	return &file_eolymp_judge_announcement_proto_enumTypes[0]
}

func (x Announcement_Extra) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Announcement_Extra.Descriptor instead.
func (Announcement_Extra) EnumDescriptor() ([]byte, []int) {
	return file_eolymp_judge_announcement_proto_rawDescGZIP(), []int{0, 0}
}

type Announcement struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Announcement unique identifier.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Contest where announcement was published.
	ContestId string `protobuf:"bytes,2,opt,name=contest_id,json=contestId,proto3" json:"contest_id,omitempty"`
	// Timestamp when announcement was initially created.
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	// Announcement subject. Max 255 symbols.
	Subject string `protobuf:"bytes,10,opt,name=subject,proto3" json:"subject,omitempty"`
	// Announcement content.
	Message       *ecm.Content `protobuf:"bytes,11,opt,name=message,proto3" json:"message,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Announcement) Reset() {
	*x = Announcement{}
	mi := &file_eolymp_judge_announcement_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Announcement) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Announcement) ProtoMessage() {}

func (x *Announcement) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_judge_announcement_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Announcement.ProtoReflect.Descriptor instead.
func (*Announcement) Descriptor() ([]byte, []int) {
	return file_eolymp_judge_announcement_proto_rawDescGZIP(), []int{0}
}

func (x *Announcement) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Announcement) GetContestId() string {
	if x != nil {
		return x.ContestId
	}
	return ""
}

func (x *Announcement) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Announcement) GetSubject() string {
	if x != nil {
		return x.Subject
	}
	return ""
}

func (x *Announcement) GetMessage() *ecm.Content {
	if x != nil {
		return x.Message
	}
	return nil
}

var File_eolymp_judge_announcement_proto protoreflect.FileDescriptor

var file_eolymp_judge_announcement_proto_rawDesc = string([]byte{
	0x0a, 0x1f, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x6a, 0x75, 0x64, 0x67, 0x65, 0x2f, 0x61,
	0x6e, 0x6e, 0x6f, 0x75, 0x6e, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0c, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x6a, 0x75, 0x64, 0x67, 0x65, 0x1a,
	0x18, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x65, 0x63, 0x6d, 0x2f, 0x63, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xff, 0x01, 0x0a, 0x0c, 0x41,
	0x6e, 0x6e, 0x6f, 0x75, 0x6e, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x73, 0x74, 0x49, 0x64, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12,
	0x2d, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x13, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x65, 0x63, 0x6d, 0x2e, 0x43, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x3c,
	0x0a, 0x05, 0x45, 0x78, 0x74, 0x72, 0x61, 0x12, 0x0c, 0x0a, 0x08, 0x4e, 0x4f, 0x5f, 0x45, 0x58,
	0x54, 0x52, 0x41, 0x10, 0x00, 0x12, 0x12, 0x0a, 0x0e, 0x4d, 0x45, 0x53, 0x53, 0x41, 0x47, 0x45,
	0x5f, 0x52, 0x45, 0x4e, 0x44, 0x45, 0x52, 0x10, 0x01, 0x12, 0x11, 0x0a, 0x0d, 0x4d, 0x45, 0x53,
	0x53, 0x41, 0x47, 0x45, 0x5f, 0x56, 0x41, 0x4c, 0x55, 0x45, 0x10, 0x02, 0x42, 0x2d, 0x5a, 0x2b,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6f, 0x6c, 0x79, 0x6d,
	0x70, 0x2f, 0x67, 0x6f, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f,
	0x6a, 0x75, 0x64, 0x67, 0x65, 0x3b, 0x6a, 0x75, 0x64, 0x67, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
})

var (
	file_eolymp_judge_announcement_proto_rawDescOnce sync.Once
	file_eolymp_judge_announcement_proto_rawDescData []byte
)

func file_eolymp_judge_announcement_proto_rawDescGZIP() []byte {
	file_eolymp_judge_announcement_proto_rawDescOnce.Do(func() {
		file_eolymp_judge_announcement_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_eolymp_judge_announcement_proto_rawDesc), len(file_eolymp_judge_announcement_proto_rawDesc)))
	})
	return file_eolymp_judge_announcement_proto_rawDescData
}

var file_eolymp_judge_announcement_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_eolymp_judge_announcement_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_eolymp_judge_announcement_proto_goTypes = []any{
	(Announcement_Extra)(0),       // 0: eolymp.judge.Announcement.Extra
	(*Announcement)(nil),          // 1: eolymp.judge.Announcement
	(*timestamppb.Timestamp)(nil), // 2: google.protobuf.Timestamp
	(*ecm.Content)(nil),           // 3: eolymp.ecm.Content
}
var file_eolymp_judge_announcement_proto_depIdxs = []int32{
	2, // 0: eolymp.judge.Announcement.created_at:type_name -> google.protobuf.Timestamp
	3, // 1: eolymp.judge.Announcement.message:type_name -> eolymp.ecm.Content
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_eolymp_judge_announcement_proto_init() }
func file_eolymp_judge_announcement_proto_init() {
	if File_eolymp_judge_announcement_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_eolymp_judge_announcement_proto_rawDesc), len(file_eolymp_judge_announcement_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_eolymp_judge_announcement_proto_goTypes,
		DependencyIndexes: file_eolymp_judge_announcement_proto_depIdxs,
		EnumInfos:         file_eolymp_judge_announcement_proto_enumTypes,
		MessageInfos:      file_eolymp_judge_announcement_proto_msgTypes,
	}.Build()
	File_eolymp_judge_announcement_proto = out.File
	file_eolymp_judge_announcement_proto_goTypes = nil
	file_eolymp_judge_announcement_proto_depIdxs = nil
}
