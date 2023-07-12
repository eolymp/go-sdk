// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.18.1
// source: eolymp/typewriter/events.proto

package typewriter

import (
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

type FragmentChangedEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Before *Fragment `protobuf:"bytes,1,opt,name=before,proto3" json:"before,omitempty"`
	After  *Fragment `protobuf:"bytes,2,opt,name=after,proto3" json:"after,omitempty"`
}

func (x *FragmentChangedEvent) Reset() {
	*x = FragmentChangedEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_typewriter_events_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FragmentChangedEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FragmentChangedEvent) ProtoMessage() {}

func (x *FragmentChangedEvent) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_typewriter_events_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
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
	return file_eolymp_typewriter_events_proto_rawDescGZIP(), []int{0}
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

var File_eolymp_typewriter_events_proto protoreflect.FileDescriptor

var file_eolymp_typewriter_events_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x77, 0x72, 0x69,
	0x74, 0x65, 0x72, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x11, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x77, 0x72, 0x69,
	0x74, 0x65, 0x72, 0x1a, 0x20, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x74, 0x79, 0x70, 0x65,
	0x77, 0x72, 0x69, 0x74, 0x65, 0x72, 0x2f, 0x66, 0x72, 0x61, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x7e, 0x0a, 0x14, 0x46, 0x72, 0x61, 0x67, 0x6d, 0x65, 0x6e,
	0x74, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x64, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x33, 0x0a,
	0x06, 0x62, 0x65, 0x66, 0x6f, 0x72, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e,
	0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x77, 0x72, 0x69, 0x74, 0x65,
	0x72, 0x2e, 0x46, 0x72, 0x61, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x06, 0x62, 0x65, 0x66, 0x6f,
	0x72, 0x65, 0x12, 0x31, 0x0a, 0x05, 0x61, 0x66, 0x74, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1b, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x77,
	0x72, 0x69, 0x74, 0x65, 0x72, 0x2e, 0x46, 0x72, 0x61, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x05,
	0x61, 0x66, 0x74, 0x65, 0x72, 0x42, 0x37, 0x5a, 0x35, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x67, 0x6f, 0x2d, 0x73, 0x64,
	0x6b, 0x2f, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x77, 0x72, 0x69,
	0x74, 0x65, 0x72, 0x3b, 0x74, 0x79, 0x70, 0x65, 0x77, 0x72, 0x69, 0x74, 0x65, 0x72, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_eolymp_typewriter_events_proto_rawDescOnce sync.Once
	file_eolymp_typewriter_events_proto_rawDescData = file_eolymp_typewriter_events_proto_rawDesc
)

func file_eolymp_typewriter_events_proto_rawDescGZIP() []byte {
	file_eolymp_typewriter_events_proto_rawDescOnce.Do(func() {
		file_eolymp_typewriter_events_proto_rawDescData = protoimpl.X.CompressGZIP(file_eolymp_typewriter_events_proto_rawDescData)
	})
	return file_eolymp_typewriter_events_proto_rawDescData
}

var file_eolymp_typewriter_events_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_eolymp_typewriter_events_proto_goTypes = []interface{}{
	(*FragmentChangedEvent)(nil), // 0: eolymp.typewriter.FragmentChangedEvent
	(*Fragment)(nil),             // 1: eolymp.typewriter.Fragment
}
var file_eolymp_typewriter_events_proto_depIdxs = []int32{
	1, // 0: eolymp.typewriter.FragmentChangedEvent.before:type_name -> eolymp.typewriter.Fragment
	1, // 1: eolymp.typewriter.FragmentChangedEvent.after:type_name -> eolymp.typewriter.Fragment
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_eolymp_typewriter_events_proto_init() }
func file_eolymp_typewriter_events_proto_init() {
	if File_eolymp_typewriter_events_proto != nil {
		return
	}
	file_eolymp_typewriter_fragment_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_eolymp_typewriter_events_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FragmentChangedEvent); i {
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
			RawDescriptor: file_eolymp_typewriter_events_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_eolymp_typewriter_events_proto_goTypes,
		DependencyIndexes: file_eolymp_typewriter_events_proto_depIdxs,
		MessageInfos:      file_eolymp_typewriter_events_proto_msgTypes,
	}.Build()
	File_eolymp_typewriter_events_proto = out.File
	file_eolymp_typewriter_events_proto_rawDesc = nil
	file_eolymp_typewriter_events_proto_goTypes = nil
	file_eolymp_typewriter_events_proto_depIdxs = nil
}
