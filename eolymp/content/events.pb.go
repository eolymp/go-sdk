// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v4.24.4
// source: eolymp/content/events.proto

package content

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

type FragmentChangedEvent struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Scope         string                 `protobuf:"bytes,10,opt,name=scope,proto3" json:"scope,omitempty"`
	Before        *Fragment              `protobuf:"bytes,1,opt,name=before,proto3" json:"before,omitempty"`
	After         *Fragment              `protobuf:"bytes,2,opt,name=after,proto3" json:"after,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *FragmentChangedEvent) Reset() {
	*x = FragmentChangedEvent{}
	mi := &file_eolymp_content_events_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FragmentChangedEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FragmentChangedEvent) ProtoMessage() {}

func (x *FragmentChangedEvent) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_content_events_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FragmentChangedEvent.ProtoReflect.Descriptor instead.
func (*FragmentChangedEvent) Descriptor() ([]byte, []int) {
	return file_eolymp_content_events_proto_rawDescGZIP(), []int{0}
}

func (x *FragmentChangedEvent) GetScope() string {
	if x != nil {
		return x.Scope
	}
	return ""
}

func (x *FragmentChangedEvent) GetBefore() *Fragment {
	if x != nil {
		return x.Before
	}
	return nil
}

func (x *FragmentChangedEvent) GetAfter() *Fragment {
	if x != nil {
		return x.After
	}
	return nil
}

type VariantChangedEvent struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	FragmentId    string                 `protobuf:"bytes,10,opt,name=fragment_id,json=fragmentId,proto3" json:"fragment_id,omitempty"`
	Before        *Variant               `protobuf:"bytes,1,opt,name=before,proto3" json:"before,omitempty"`
	After         *Variant               `protobuf:"bytes,2,opt,name=after,proto3" json:"after,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *VariantChangedEvent) Reset() {
	*x = VariantChangedEvent{}
	mi := &file_eolymp_content_events_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *VariantChangedEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VariantChangedEvent) ProtoMessage() {}

func (x *VariantChangedEvent) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_content_events_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VariantChangedEvent.ProtoReflect.Descriptor instead.
func (*VariantChangedEvent) Descriptor() ([]byte, []int) {
	return file_eolymp_content_events_proto_rawDescGZIP(), []int{1}
}

func (x *VariantChangedEvent) GetFragmentId() string {
	if x != nil {
		return x.FragmentId
	}
	return ""
}

func (x *VariantChangedEvent) GetBefore() *Variant {
	if x != nil {
		return x.Before
	}
	return nil
}

func (x *VariantChangedEvent) GetAfter() *Variant {
	if x != nil {
		return x.After
	}
	return nil
}

var File_eolymp_content_events_proto protoreflect.FileDescriptor

var file_eolymp_content_events_proto_rawDesc = string([]byte{
	0x0a, 0x1b, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x65,
	0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x1a, 0x1d, 0x65,
	0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x2f, 0x66, 0x72,
	0x61, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x65, 0x6f,
	0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x2f, 0x76, 0x61, 0x72,
	0x69, 0x61, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8e, 0x01, 0x0a, 0x14, 0x46,
	0x72, 0x61, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x64, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x12, 0x30, 0x0a, 0x06, 0x62, 0x65, 0x66,
	0x6f, 0x72, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x65, 0x6f, 0x6c, 0x79,
	0x6d, 0x70, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x2e, 0x46, 0x72, 0x61, 0x67, 0x6d,
	0x65, 0x6e, 0x74, 0x52, 0x06, 0x62, 0x65, 0x66, 0x6f, 0x72, 0x65, 0x12, 0x2e, 0x0a, 0x05, 0x61,
	0x66, 0x74, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x65, 0x6f, 0x6c,
	0x79, 0x6d, 0x70, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x2e, 0x46, 0x72, 0x61, 0x67,
	0x6d, 0x65, 0x6e, 0x74, 0x52, 0x05, 0x61, 0x66, 0x74, 0x65, 0x72, 0x22, 0x96, 0x01, 0x0a, 0x13,
	0x56, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x74, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x64, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x66, 0x72, 0x61, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x5f,
	0x69, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x66, 0x72, 0x61, 0x67, 0x6d, 0x65,
	0x6e, 0x74, 0x49, 0x64, 0x12, 0x2f, 0x0a, 0x06, 0x62, 0x65, 0x66, 0x6f, 0x72, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x2e, 0x56, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x74, 0x52, 0x06, 0x62,
	0x65, 0x66, 0x6f, 0x72, 0x65, 0x12, 0x2d, 0x0a, 0x05, 0x61, 0x66, 0x74, 0x65, 0x72, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x2e, 0x56, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x74, 0x52, 0x05, 0x61,
	0x66, 0x74, 0x65, 0x72, 0x42, 0x31, 0x5a, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x67, 0x6f, 0x2d, 0x73, 0x64, 0x6b,
	0x2f, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x3b,
	0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_eolymp_content_events_proto_rawDescOnce sync.Once
	file_eolymp_content_events_proto_rawDescData []byte
)

func file_eolymp_content_events_proto_rawDescGZIP() []byte {
	file_eolymp_content_events_proto_rawDescOnce.Do(func() {
		file_eolymp_content_events_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_eolymp_content_events_proto_rawDesc), len(file_eolymp_content_events_proto_rawDesc)))
	})
	return file_eolymp_content_events_proto_rawDescData
}

var file_eolymp_content_events_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_eolymp_content_events_proto_goTypes = []any{
	(*FragmentChangedEvent)(nil), // 0: eolymp.content.FragmentChangedEvent
	(*VariantChangedEvent)(nil),  // 1: eolymp.content.VariantChangedEvent
	(*Fragment)(nil),             // 2: eolymp.content.Fragment
	(*Variant)(nil),              // 3: eolymp.content.Variant
}
var file_eolymp_content_events_proto_depIdxs = []int32{
	2, // 0: eolymp.content.FragmentChangedEvent.before:type_name -> eolymp.content.Fragment
	2, // 1: eolymp.content.FragmentChangedEvent.after:type_name -> eolymp.content.Fragment
	3, // 2: eolymp.content.VariantChangedEvent.before:type_name -> eolymp.content.Variant
	3, // 3: eolymp.content.VariantChangedEvent.after:type_name -> eolymp.content.Variant
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_eolymp_content_events_proto_init() }
func file_eolymp_content_events_proto_init() {
	if File_eolymp_content_events_proto != nil {
		return
	}
	file_eolymp_content_fragment_proto_init()
	file_eolymp_content_variant_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_eolymp_content_events_proto_rawDesc), len(file_eolymp_content_events_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_eolymp_content_events_proto_goTypes,
		DependencyIndexes: file_eolymp_content_events_proto_depIdxs,
		MessageInfos:      file_eolymp_content_events_proto_msgTypes,
	}.Build()
	File_eolymp_content_events_proto = out.File
	file_eolymp_content_events_proto_goTypes = nil
	file_eolymp_content_events_proto_depIdxs = nil
}
