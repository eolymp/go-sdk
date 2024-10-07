// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v3.21.12
// source: eolymp/annotations/endpoint.proto

package annotations

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Endpoint struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Field   string              `protobuf:"bytes,1,opt,name=field,proto3" json:"field,omitempty"`
	Service []*Endpoint_Service `protobuf:"bytes,2,rep,name=service,proto3" json:"service,omitempty"`
}

func (x *Endpoint) Reset() {
	*x = Endpoint{}
	mi := &file_eolymp_annotations_endpoint_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Endpoint) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Endpoint) ProtoMessage() {}

func (x *Endpoint) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_annotations_endpoint_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Endpoint.ProtoReflect.Descriptor instead.
func (*Endpoint) Descriptor() ([]byte, []int) {
	return file_eolymp_annotations_endpoint_proto_rawDescGZIP(), []int{0}
}

func (x *Endpoint) GetField() string {
	if x != nil {
		return x.Field
	}
	return ""
}

func (x *Endpoint) GetService() []*Endpoint_Service {
	if x != nil {
		return x.Service
	}
	return nil
}

type Endpoint_Service struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Endpoint_Service) Reset() {
	*x = Endpoint_Service{}
	mi := &file_eolymp_annotations_endpoint_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Endpoint_Service) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Endpoint_Service) ProtoMessage() {}

func (x *Endpoint_Service) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_annotations_endpoint_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Endpoint_Service.ProtoReflect.Descriptor instead.
func (*Endpoint_Service) Descriptor() ([]byte, []int) {
	return file_eolymp_annotations_endpoint_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Endpoint_Service) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var file_eolymp_annotations_endpoint_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptorpb.MessageOptions)(nil),
		ExtensionType: (*Endpoint)(nil),
		Field:         22075,
		Name:          "eolymp.api.endpoint",
		Tag:           "bytes,22075,opt,name=endpoint",
		Filename:      "eolymp/annotations/endpoint.proto",
	},
}

// Extension fields to descriptorpb.MessageOptions.
var (
	// optional eolymp.api.Endpoint endpoint = 22075;
	E_Endpoint = &file_eolymp_annotations_endpoint_proto_extTypes[0]
)

var File_eolymp_annotations_endpoint_proto protoreflect.FileDescriptor

var file_eolymp_annotations_endpoint_proto_rawDesc = []byte{
	0x0a, 0x21, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x61, 0x70, 0x69, 0x1a,
	0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x77, 0x0a, 0x08, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x14, 0x0a,
	0x05, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x66, 0x69,
	0x65, 0x6c, 0x64, 0x12, 0x36, 0x0a, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x52, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x1a, 0x1d, 0x0a, 0x07, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x3a, 0x53, 0x0a, 0x08, 0x65, 0x6e,
	0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x1f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xbb, 0xac, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x14, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x45, 0x6e, 0x64,
	0x70, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x08, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x42,
	0x39, 0x5a, 0x37, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6f,
	0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x67, 0x6f, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x65, 0x6f, 0x6c, 0x79,
	0x6d, 0x70, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x3b, 0x61,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_eolymp_annotations_endpoint_proto_rawDescOnce sync.Once
	file_eolymp_annotations_endpoint_proto_rawDescData = file_eolymp_annotations_endpoint_proto_rawDesc
)

func file_eolymp_annotations_endpoint_proto_rawDescGZIP() []byte {
	file_eolymp_annotations_endpoint_proto_rawDescOnce.Do(func() {
		file_eolymp_annotations_endpoint_proto_rawDescData = protoimpl.X.CompressGZIP(file_eolymp_annotations_endpoint_proto_rawDescData)
	})
	return file_eolymp_annotations_endpoint_proto_rawDescData
}

var file_eolymp_annotations_endpoint_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_eolymp_annotations_endpoint_proto_goTypes = []any{
	(*Endpoint)(nil),                    // 0: eolymp.api.Endpoint
	(*Endpoint_Service)(nil),            // 1: eolymp.api.Endpoint.Service
	(*descriptorpb.MessageOptions)(nil), // 2: google.protobuf.MessageOptions
}
var file_eolymp_annotations_endpoint_proto_depIdxs = []int32{
	1, // 0: eolymp.api.Endpoint.service:type_name -> eolymp.api.Endpoint.Service
	2, // 1: eolymp.api.endpoint:extendee -> google.protobuf.MessageOptions
	0, // 2: eolymp.api.endpoint:type_name -> eolymp.api.Endpoint
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	2, // [2:3] is the sub-list for extension type_name
	1, // [1:2] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_eolymp_annotations_endpoint_proto_init() }
func file_eolymp_annotations_endpoint_proto_init() {
	if File_eolymp_annotations_endpoint_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_eolymp_annotations_endpoint_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 1,
			NumServices:   0,
		},
		GoTypes:           file_eolymp_annotations_endpoint_proto_goTypes,
		DependencyIndexes: file_eolymp_annotations_endpoint_proto_depIdxs,
		MessageInfos:      file_eolymp_annotations_endpoint_proto_msgTypes,
		ExtensionInfos:    file_eolymp_annotations_endpoint_proto_extTypes,
	}.Build()
	File_eolymp_annotations_endpoint_proto = out.File
	file_eolymp_annotations_endpoint_proto_rawDesc = nil
	file_eolymp_annotations_endpoint_proto_goTypes = nil
	file_eolymp_annotations_endpoint_proto_depIdxs = nil
}
