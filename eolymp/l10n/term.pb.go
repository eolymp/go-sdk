// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.4
// 	protoc        v4.24.4
// source: eolymp/l10n/term.proto

package l10n

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
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

type Term_Status int32

const (
	Term_NONE       Term_Status = 0 // reserved, not in use
	Term_ACTIVE     Term_Status = 1 // active term, currently in use
	Term_DEPRECATED Term_Status = 2 // not in use anymore (not exported)
)

// Enum value maps for Term_Status.
var (
	Term_Status_name = map[int32]string{
		0: "NONE",
		1: "ACTIVE",
		2: "DEPRECATED",
	}
	Term_Status_value = map[string]int32{
		"NONE":       0,
		"ACTIVE":     1,
		"DEPRECATED": 2,
	}
)

func (x Term_Status) Enum() *Term_Status {
	p := new(Term_Status)
	*p = x
	return p
}

func (x Term_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Term_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_eolymp_l10n_term_proto_enumTypes[0].Descriptor()
}

func (Term_Status) Type() protoreflect.EnumType {
	return &file_eolymp_l10n_term_proto_enumTypes[0]
}

func (x Term_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Term_Status.Descriptor instead.
func (Term_Status) EnumDescriptor() ([]byte, []int) {
	return file_eolymp_l10n_term_proto_rawDescGZIP(), []int{0, 0}
}

type Term struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Key           string                 `protobuf:"bytes,5,opt,name=key,proto3" json:"key,omitempty"`
	Message       string                 `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Description   string                 `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Status        Term_Status            `protobuf:"varint,4,opt,name=status,proto3,enum=eolymp.l10n.Term_Status" json:"status,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Term) Reset() {
	*x = Term{}
	mi := &file_eolymp_l10n_term_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Term) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Term) ProtoMessage() {}

func (x *Term) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_l10n_term_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Term.ProtoReflect.Descriptor instead.
func (*Term) Descriptor() ([]byte, []int) {
	return file_eolymp_l10n_term_proto_rawDescGZIP(), []int{0}
}

func (x *Term) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Term) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *Term) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *Term) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Term) GetStatus() Term_Status {
	if x != nil {
		return x.Status
	}
	return Term_NONE
}

var File_eolymp_l10n_term_proto protoreflect.FileDescriptor

var file_eolymp_l10n_term_proto_rawDesc = string([]byte{
	0x0a, 0x16, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x6c, 0x31, 0x30, 0x6e, 0x2f, 0x74, 0x65,
	0x72, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70,
	0x2e, 0x6c, 0x31, 0x30, 0x6e, 0x22, 0xc6, 0x01, 0x0a, 0x04, 0x54, 0x65, 0x72, 0x6d, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x10,
	0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79,
	0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x30, 0x0a, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x18, 0x2e, 0x65,
	0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x6c, 0x31, 0x30, 0x6e, 0x2e, 0x54, 0x65, 0x72, 0x6d, 0x2e,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x2e,
	0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x4f, 0x4e, 0x45,
	0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x41, 0x43, 0x54, 0x49, 0x56, 0x45, 0x10, 0x01, 0x12, 0x0e,
	0x0a, 0x0a, 0x44, 0x45, 0x50, 0x52, 0x45, 0x43, 0x41, 0x54, 0x45, 0x44, 0x10, 0x02, 0x42, 0x2b,
	0x5a, 0x29, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6f, 0x6c,
	0x79, 0x6d, 0x70, 0x2f, 0x67, 0x6f, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x65, 0x6f, 0x6c, 0x79, 0x6d,
	0x70, 0x2f, 0x6c, 0x31, 0x30, 0x6e, 0x3b, 0x6c, 0x31, 0x30, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
})

var (
	file_eolymp_l10n_term_proto_rawDescOnce sync.Once
	file_eolymp_l10n_term_proto_rawDescData []byte
)

func file_eolymp_l10n_term_proto_rawDescGZIP() []byte {
	file_eolymp_l10n_term_proto_rawDescOnce.Do(func() {
		file_eolymp_l10n_term_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_eolymp_l10n_term_proto_rawDesc), len(file_eolymp_l10n_term_proto_rawDesc)))
	})
	return file_eolymp_l10n_term_proto_rawDescData
}

var file_eolymp_l10n_term_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_eolymp_l10n_term_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_eolymp_l10n_term_proto_goTypes = []any{
	(Term_Status)(0), // 0: eolymp.l10n.Term.Status
	(*Term)(nil),     // 1: eolymp.l10n.Term
}
var file_eolymp_l10n_term_proto_depIdxs = []int32{
	0, // 0: eolymp.l10n.Term.status:type_name -> eolymp.l10n.Term.Status
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_eolymp_l10n_term_proto_init() }
func file_eolymp_l10n_term_proto_init() {
	if File_eolymp_l10n_term_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_eolymp_l10n_term_proto_rawDesc), len(file_eolymp_l10n_term_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_eolymp_l10n_term_proto_goTypes,
		DependencyIndexes: file_eolymp_l10n_term_proto_depIdxs,
		EnumInfos:         file_eolymp_l10n_term_proto_enumTypes,
		MessageInfos:      file_eolymp_l10n_term_proto_msgTypes,
	}.Build()
	File_eolymp_l10n_term_proto = out.File
	file_eolymp_l10n_term_proto_goTypes = nil
	file_eolymp_l10n_term_proto_depIdxs = nil
}
