// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.29.0
// 	protoc        v3.18.1
// source: eolymp/harmony/consent.proto

package harmony

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

type Consent_Status int32

const (
	Consent_PENDING  Consent_Status = 0
	Consent_ACCEPTED Consent_Status = 1
	Consent_REJECTED Consent_Status = 2
)

// Enum value maps for Consent_Status.
var (
	Consent_Status_name = map[int32]string{
		0: "PENDING",
		1: "ACCEPTED",
		2: "REJECTED",
	}
	Consent_Status_value = map[string]int32{
		"PENDING":  0,
		"ACCEPTED": 1,
		"REJECTED": 2,
	}
)

func (x Consent_Status) Enum() *Consent_Status {
	p := new(Consent_Status)
	*p = x
	return p
}

func (x Consent_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Consent_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_eolymp_harmony_consent_proto_enumTypes[0].Descriptor()
}

func (Consent_Status) Type() protoreflect.EnumType {
	return &file_eolymp_harmony_consent_proto_enumTypes[0]
}

func (x Consent_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Consent_Status.Descriptor instead.
func (Consent_Status) EnumDescriptor() ([]byte, []int) {
	return file_eolymp_harmony_consent_proto_rawDescGZIP(), []int{0, 0}
}

type Consent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	AgreementId string                 `protobuf:"bytes,2,opt,name=agreement_id,json=agreementId,proto3" json:"agreement_id,omitempty"`
	UserId      string                 `protobuf:"bytes,3,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Revocable   bool                   `protobuf:"varint,4,opt,name=revocable,proto3" json:"revocable,omitempty"`
	Status      Consent_Status         `protobuf:"varint,10,opt,name=status,proto3,enum=eolymp.harmony.Consent_Status" json:"status,omitempty"`
	SetOn       *timestamppb.Timestamp `protobuf:"bytes,11,opt,name=set_on,json=setOn,proto3" json:"set_on,omitempty"`
}

func (x *Consent) Reset() {
	*x = Consent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_harmony_consent_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Consent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Consent) ProtoMessage() {}

func (x *Consent) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_harmony_consent_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Consent.ProtoReflect.Descriptor instead.
func (*Consent) Descriptor() ([]byte, []int) {
	return file_eolymp_harmony_consent_proto_rawDescGZIP(), []int{0}
}

func (x *Consent) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Consent) GetAgreementId() string {
	if x != nil {
		return x.AgreementId
	}
	return ""
}

func (x *Consent) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *Consent) GetRevocable() bool {
	if x != nil {
		return x.Revocable
	}
	return false
}

func (x *Consent) GetStatus() Consent_Status {
	if x != nil {
		return x.Status
	}
	return Consent_PENDING
}

func (x *Consent) GetSetOn() *timestamppb.Timestamp {
	if x != nil {
		return x.SetOn
	}
	return nil
}

var File_eolymp_harmony_consent_proto protoreflect.FileDescriptor

var file_eolymp_harmony_consent_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x68, 0x61, 0x72, 0x6d, 0x6f, 0x6e, 0x79,
	0x2f, 0x63, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e,
	0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x68, 0x61, 0x72, 0x6d, 0x6f, 0x6e, 0x79, 0x1a, 0x1f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x91, 0x02, 0x0a, 0x07, 0x43, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x61,
	0x67, 0x72, 0x65, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x61, 0x67, 0x72, 0x65, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x17,
	0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x72, 0x65, 0x76, 0x6f, 0x63,
	0x61, 0x62, 0x6c, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x72, 0x65, 0x76, 0x6f,
	0x63, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x36, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1e, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x68,
	0x61, 0x72, 0x6d, 0x6f, 0x6e, 0x79, 0x2e, 0x43, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x74, 0x2e, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x31, 0x0a,
	0x06, 0x73, 0x65, 0x74, 0x5f, 0x6f, 0x6e, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x05, 0x73, 0x65, 0x74, 0x4f, 0x6e,
	0x22, 0x31, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0b, 0x0a, 0x07, 0x50, 0x45,
	0x4e, 0x44, 0x49, 0x4e, 0x47, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x41, 0x43, 0x43, 0x45, 0x50,
	0x54, 0x45, 0x44, 0x10, 0x01, 0x12, 0x0c, 0x0a, 0x08, 0x52, 0x45, 0x4a, 0x45, 0x43, 0x54, 0x45,
	0x44, 0x10, 0x02, 0x42, 0x31, 0x5a, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x67, 0x6f, 0x2d, 0x73, 0x64, 0x6b, 0x2f,
	0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x68, 0x61, 0x72, 0x6d, 0x6f, 0x6e, 0x79, 0x3b, 0x68,
	0x61, 0x72, 0x6d, 0x6f, 0x6e, 0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_eolymp_harmony_consent_proto_rawDescOnce sync.Once
	file_eolymp_harmony_consent_proto_rawDescData = file_eolymp_harmony_consent_proto_rawDesc
)

func file_eolymp_harmony_consent_proto_rawDescGZIP() []byte {
	file_eolymp_harmony_consent_proto_rawDescOnce.Do(func() {
		file_eolymp_harmony_consent_proto_rawDescData = protoimpl.X.CompressGZIP(file_eolymp_harmony_consent_proto_rawDescData)
	})
	return file_eolymp_harmony_consent_proto_rawDescData
}

var file_eolymp_harmony_consent_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_eolymp_harmony_consent_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_eolymp_harmony_consent_proto_goTypes = []interface{}{
	(Consent_Status)(0),           // 0: eolymp.harmony.Consent.Status
	(*Consent)(nil),               // 1: eolymp.harmony.Consent
	(*timestamppb.Timestamp)(nil), // 2: google.protobuf.Timestamp
}
var file_eolymp_harmony_consent_proto_depIdxs = []int32{
	0, // 0: eolymp.harmony.Consent.status:type_name -> eolymp.harmony.Consent.Status
	2, // 1: eolymp.harmony.Consent.set_on:type_name -> google.protobuf.Timestamp
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_eolymp_harmony_consent_proto_init() }
func file_eolymp_harmony_consent_proto_init() {
	if File_eolymp_harmony_consent_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_eolymp_harmony_consent_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Consent); i {
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
			RawDescriptor: file_eolymp_harmony_consent_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_eolymp_harmony_consent_proto_goTypes,
		DependencyIndexes: file_eolymp_harmony_consent_proto_depIdxs,
		EnumInfos:         file_eolymp_harmony_consent_proto_enumTypes,
		MessageInfos:      file_eolymp_harmony_consent_proto_msgTypes,
	}.Build()
	File_eolymp_harmony_consent_proto = out.File
	file_eolymp_harmony_consent_proto_rawDesc = nil
	file_eolymp_harmony_consent_proto_goTypes = nil
	file_eolymp_harmony_consent_proto_depIdxs = nil
}
