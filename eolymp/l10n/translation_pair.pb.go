// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v4.24.4
// source: eolymp/l10n/translation_pair.proto

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

type TranslationPair struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Term          *Term                  `protobuf:"bytes,1,opt,name=term,proto3" json:"term,omitempty"`
	Source        *Translation           `protobuf:"bytes,2,opt,name=source,proto3" json:"source,omitempty"`
	Translation   *Translation           `protobuf:"bytes,3,opt,name=translation,proto3" json:"translation,omitempty"`
	Suggestion    *Translation           `protobuf:"bytes,4,opt,name=suggestion,proto3" json:"suggestion,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *TranslationPair) Reset() {
	*x = TranslationPair{}
	mi := &file_eolymp_l10n_translation_pair_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TranslationPair) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TranslationPair) ProtoMessage() {}

func (x *TranslationPair) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_l10n_translation_pair_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TranslationPair.ProtoReflect.Descriptor instead.
func (*TranslationPair) Descriptor() ([]byte, []int) {
	return file_eolymp_l10n_translation_pair_proto_rawDescGZIP(), []int{0}
}

func (x *TranslationPair) GetTerm() *Term {
	if x != nil {
		return x.Term
	}
	return nil
}

func (x *TranslationPair) GetSource() *Translation {
	if x != nil {
		return x.Source
	}
	return nil
}

func (x *TranslationPair) GetTranslation() *Translation {
	if x != nil {
		return x.Translation
	}
	return nil
}

func (x *TranslationPair) GetSuggestion() *Translation {
	if x != nil {
		return x.Suggestion
	}
	return nil
}

var File_eolymp_l10n_translation_pair_proto protoreflect.FileDescriptor

var file_eolymp_l10n_translation_pair_proto_rawDesc = string([]byte{
	0x0a, 0x22, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x6c, 0x31, 0x30, 0x6e, 0x2f, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x70, 0x61, 0x69, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x6c, 0x31, 0x30,
	0x6e, 0x1a, 0x16, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x6c, 0x31, 0x30, 0x6e, 0x2f, 0x74,
	0x65, 0x72, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x65, 0x6f, 0x6c, 0x79, 0x6d,
	0x70, 0x2f, 0x6c, 0x31, 0x30, 0x6e, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x6c, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe0, 0x01, 0x0a, 0x0f, 0x54, 0x72, 0x61,
	0x6e, 0x73, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x61, 0x69, 0x72, 0x12, 0x25, 0x0a, 0x04,
	0x74, 0x65, 0x72, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x65, 0x6f, 0x6c,
	0x79, 0x6d, 0x70, 0x2e, 0x6c, 0x31, 0x30, 0x6e, 0x2e, 0x54, 0x65, 0x72, 0x6d, 0x52, 0x04, 0x74,
	0x65, 0x72, 0x6d, 0x12, 0x30, 0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x6c, 0x31, 0x30,
	0x6e, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x06, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x3a, 0x0a, 0x0b, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x6c, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x65, 0x6f, 0x6c,
	0x79, 0x6d, 0x70, 0x2e, 0x6c, 0x31, 0x30, 0x6e, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x6c, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x6c, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x38, 0x0a, 0x0a, 0x73, 0x75, 0x67, 0x67, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x6c,
	0x31, 0x30, 0x6e, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x0a, 0x73, 0x75, 0x67, 0x67, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x2b, 0x5a, 0x29, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70,
	0x2f, 0x67, 0x6f, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x6c,
	0x31, 0x30, 0x6e, 0x3b, 0x6c, 0x31, 0x30, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_eolymp_l10n_translation_pair_proto_rawDescOnce sync.Once
	file_eolymp_l10n_translation_pair_proto_rawDescData []byte
)

func file_eolymp_l10n_translation_pair_proto_rawDescGZIP() []byte {
	file_eolymp_l10n_translation_pair_proto_rawDescOnce.Do(func() {
		file_eolymp_l10n_translation_pair_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_eolymp_l10n_translation_pair_proto_rawDesc), len(file_eolymp_l10n_translation_pair_proto_rawDesc)))
	})
	return file_eolymp_l10n_translation_pair_proto_rawDescData
}

var file_eolymp_l10n_translation_pair_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_eolymp_l10n_translation_pair_proto_goTypes = []any{
	(*TranslationPair)(nil), // 0: eolymp.l10n.TranslationPair
	(*Term)(nil),            // 1: eolymp.l10n.Term
	(*Translation)(nil),     // 2: eolymp.l10n.Translation
}
var file_eolymp_l10n_translation_pair_proto_depIdxs = []int32{
	1, // 0: eolymp.l10n.TranslationPair.term:type_name -> eolymp.l10n.Term
	2, // 1: eolymp.l10n.TranslationPair.source:type_name -> eolymp.l10n.Translation
	2, // 2: eolymp.l10n.TranslationPair.translation:type_name -> eolymp.l10n.Translation
	2, // 3: eolymp.l10n.TranslationPair.suggestion:type_name -> eolymp.l10n.Translation
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_eolymp_l10n_translation_pair_proto_init() }
func file_eolymp_l10n_translation_pair_proto_init() {
	if File_eolymp_l10n_translation_pair_proto != nil {
		return
	}
	file_eolymp_l10n_term_proto_init()
	file_eolymp_l10n_translation_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_eolymp_l10n_translation_pair_proto_rawDesc), len(file_eolymp_l10n_translation_pair_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_eolymp_l10n_translation_pair_proto_goTypes,
		DependencyIndexes: file_eolymp_l10n_translation_pair_proto_depIdxs,
		MessageInfos:      file_eolymp_l10n_translation_pair_proto_msgTypes,
	}.Build()
	File_eolymp_l10n_translation_pair_proto = out.File
	file_eolymp_l10n_translation_pair_proto_goTypes = nil
	file_eolymp_l10n_translation_pair_proto_depIdxs = nil
}
