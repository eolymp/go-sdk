// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v4.24.4
// source: eolymp/newsletter/recipient.proto

package newsletter

import (
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

type Recipient_Status int32

const (
	Recipient_UNKNOWN_STATUS Recipient_Status = 0
	Recipient_CREATED        Recipient_Status = 1
	Recipient_PENDING        Recipient_Status = 2
	Recipient_SENT           Recipient_Status = 3
	Recipient_DELIVERED      Recipient_Status = 4
	Recipient_BOUNCED        Recipient_Status = 5
)

// Enum value maps for Recipient_Status.
var (
	Recipient_Status_name = map[int32]string{
		0: "UNKNOWN_STATUS",
		1: "CREATED",
		2: "PENDING",
		3: "SENT",
		4: "DELIVERED",
		5: "BOUNCED",
	}
	Recipient_Status_value = map[string]int32{
		"UNKNOWN_STATUS": 0,
		"CREATED":        1,
		"PENDING":        2,
		"SENT":           3,
		"DELIVERED":      4,
		"BOUNCED":        5,
	}
)

func (x Recipient_Status) Enum() *Recipient_Status {
	p := new(Recipient_Status)
	*p = x
	return p
}

func (x Recipient_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Recipient_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_eolymp_newsletter_recipient_proto_enumTypes[0].Descriptor()
}

func (Recipient_Status) Type() protoreflect.EnumType {
	return &file_eolymp_newsletter_recipient_proto_enumTypes[0]
}

func (x Recipient_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Recipient_Status.Descriptor instead.
func (Recipient_Status) EnumDescriptor() ([]byte, []int) {
	return file_eolymp_newsletter_recipient_proto_rawDescGZIP(), []int{0, 0}
}

type Recipient struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	CreatedAt     *timestamppb.Timestamp `protobuf:"bytes,10,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt     *timestamppb.Timestamp `protobuf:"bytes,11,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	SentAt        *timestamppb.Timestamp `protobuf:"bytes,12,opt,name=sent_at,json=sentAt,proto3" json:"sent_at,omitempty"`
	DeliveredAt   *timestamppb.Timestamp `protobuf:"bytes,13,opt,name=delivered_at,json=deliveredAt,proto3" json:"delivered_at,omitempty"`
	MemberId      string                 `protobuf:"bytes,2,opt,name=member_id,json=memberId,proto3" json:"member_id,omitempty"`
	Status        Recipient_Status       `protobuf:"varint,3,opt,name=status,proto3,enum=eolymp.newsletter.Recipient_Status" json:"status,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Recipient) Reset() {
	*x = Recipient{}
	mi := &file_eolymp_newsletter_recipient_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Recipient) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Recipient) ProtoMessage() {}

func (x *Recipient) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_newsletter_recipient_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Recipient.ProtoReflect.Descriptor instead.
func (*Recipient) Descriptor() ([]byte, []int) {
	return file_eolymp_newsletter_recipient_proto_rawDescGZIP(), []int{0}
}

func (x *Recipient) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Recipient) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Recipient) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *Recipient) GetSentAt() *timestamppb.Timestamp {
	if x != nil {
		return x.SentAt
	}
	return nil
}

func (x *Recipient) GetDeliveredAt() *timestamppb.Timestamp {
	if x != nil {
		return x.DeliveredAt
	}
	return nil
}

func (x *Recipient) GetMemberId() string {
	if x != nil {
		return x.MemberId
	}
	return ""
}

func (x *Recipient) GetStatus() Recipient_Status {
	if x != nil {
		return x.Status
	}
	return Recipient_UNKNOWN_STATUS
}

var File_eolymp_newsletter_recipient_proto protoreflect.FileDescriptor

var file_eolymp_newsletter_recipient_proto_rawDesc = string([]byte{
	0x0a, 0x21, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x6e, 0x65, 0x77, 0x73, 0x6c, 0x65, 0x74,
	0x74, 0x65, 0x72, 0x2f, 0x72, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x11, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x6e, 0x65, 0x77, 0x73,
	0x6c, 0x65, 0x74, 0x74, 0x65, 0x72, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xbd, 0x03, 0x0a, 0x09, 0x52, 0x65, 0x63, 0x69,
	0x70, 0x69, 0x65, 0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x5f, 0x61, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x12, 0x39, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0b,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x33, 0x0a, 0x07, 0x73,
	0x65, 0x6e, 0x74, 0x5f, 0x61, 0x74, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x06, 0x73, 0x65, 0x6e, 0x74, 0x41, 0x74,
	0x12, 0x3d, 0x0a, 0x0c, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x65, 0x64, 0x5f, 0x61, 0x74,
	0x18, 0x0d, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x0b, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x65, 0x64, 0x41, 0x74, 0x12,
	0x1b, 0x0a, 0x09, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x49, 0x64, 0x12, 0x3b, 0x0a, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x23, 0x2e, 0x65,
	0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x6e, 0x65, 0x77, 0x73, 0x6c, 0x65, 0x74, 0x74, 0x65, 0x72,
	0x2e, 0x52, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x5c, 0x0a, 0x06, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x12, 0x0a, 0x0e, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x5f, 0x53,
	0x54, 0x41, 0x54, 0x55, 0x53, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x43, 0x52, 0x45, 0x41, 0x54,
	0x45, 0x44, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x50, 0x45, 0x4e, 0x44, 0x49, 0x4e, 0x47, 0x10,
	0x02, 0x12, 0x08, 0x0a, 0x04, 0x53, 0x45, 0x4e, 0x54, 0x10, 0x03, 0x12, 0x0d, 0x0a, 0x09, 0x44,
	0x45, 0x4c, 0x49, 0x56, 0x45, 0x52, 0x45, 0x44, 0x10, 0x04, 0x12, 0x0b, 0x0a, 0x07, 0x42, 0x4f,
	0x55, 0x4e, 0x43, 0x45, 0x44, 0x10, 0x05, 0x42, 0x37, 0x5a, 0x35, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x67, 0x6f, 0x2d,
	0x73, 0x64, 0x6b, 0x2f, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x6e, 0x65, 0x77, 0x73, 0x6c,
	0x65, 0x74, 0x74, 0x65, 0x72, 0x3b, 0x6e, 0x65, 0x77, 0x73, 0x6c, 0x65, 0x74, 0x74, 0x65, 0x72,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_eolymp_newsletter_recipient_proto_rawDescOnce sync.Once
	file_eolymp_newsletter_recipient_proto_rawDescData []byte
)

func file_eolymp_newsletter_recipient_proto_rawDescGZIP() []byte {
	file_eolymp_newsletter_recipient_proto_rawDescOnce.Do(func() {
		file_eolymp_newsletter_recipient_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_eolymp_newsletter_recipient_proto_rawDesc), len(file_eolymp_newsletter_recipient_proto_rawDesc)))
	})
	return file_eolymp_newsletter_recipient_proto_rawDescData
}

var file_eolymp_newsletter_recipient_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_eolymp_newsletter_recipient_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_eolymp_newsletter_recipient_proto_goTypes = []any{
	(Recipient_Status)(0),         // 0: eolymp.newsletter.Recipient.Status
	(*Recipient)(nil),             // 1: eolymp.newsletter.Recipient
	(*timestamppb.Timestamp)(nil), // 2: google.protobuf.Timestamp
}
var file_eolymp_newsletter_recipient_proto_depIdxs = []int32{
	2, // 0: eolymp.newsletter.Recipient.created_at:type_name -> google.protobuf.Timestamp
	2, // 1: eolymp.newsletter.Recipient.updated_at:type_name -> google.protobuf.Timestamp
	2, // 2: eolymp.newsletter.Recipient.sent_at:type_name -> google.protobuf.Timestamp
	2, // 3: eolymp.newsletter.Recipient.delivered_at:type_name -> google.protobuf.Timestamp
	0, // 4: eolymp.newsletter.Recipient.status:type_name -> eolymp.newsletter.Recipient.Status
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_eolymp_newsletter_recipient_proto_init() }
func file_eolymp_newsletter_recipient_proto_init() {
	if File_eolymp_newsletter_recipient_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_eolymp_newsletter_recipient_proto_rawDesc), len(file_eolymp_newsletter_recipient_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_eolymp_newsletter_recipient_proto_goTypes,
		DependencyIndexes: file_eolymp_newsletter_recipient_proto_depIdxs,
		EnumInfos:         file_eolymp_newsletter_recipient_proto_enumTypes,
		MessageInfos:      file_eolymp_newsletter_recipient_proto_msgTypes,
	}.Build()
	File_eolymp_newsletter_recipient_proto = out.File
	file_eolymp_newsletter_recipient_proto_goTypes = nil
	file_eolymp_newsletter_recipient_proto_depIdxs = nil
}
