// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v3.21.12
// source: eolymp/atlas/code_template.proto

package atlas

import (
	executor "github.com/eolymp/go-sdk/eolymp/executor"
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

type Template struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      string           `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Runtime string           `protobuf:"bytes,3,opt,name=runtime,proto3" json:"runtime,omitempty"`
	Secret  bool             `protobuf:"varint,4,opt,name=secret,proto3" json:"secret,omitempty"` // header, footer and files are secret
	Source  string           `protobuf:"bytes,10,opt,name=source,proto3" json:"source,omitempty"`
	Header  string           `protobuf:"bytes,11,opt,name=header,proto3" json:"header,omitempty"`
	Footer  string           `protobuf:"bytes,12,opt,name=footer,proto3" json:"footer,omitempty"`
	Files   []*executor.File `protobuf:"bytes,30,rep,name=files,proto3" json:"files,omitempty"`
}

func (x *Template) Reset() {
	*x = Template{}
	mi := &file_eolymp_atlas_code_template_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Template) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Template) ProtoMessage() {}

func (x *Template) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_atlas_code_template_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Template.ProtoReflect.Descriptor instead.
func (*Template) Descriptor() ([]byte, []int) {
	return file_eolymp_atlas_code_template_proto_rawDescGZIP(), []int{0}
}

func (x *Template) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Template) GetRuntime() string {
	if x != nil {
		return x.Runtime
	}
	return ""
}

func (x *Template) GetSecret() bool {
	if x != nil {
		return x.Secret
	}
	return false
}

func (x *Template) GetSource() string {
	if x != nil {
		return x.Source
	}
	return ""
}

func (x *Template) GetHeader() string {
	if x != nil {
		return x.Header
	}
	return ""
}

func (x *Template) GetFooter() string {
	if x != nil {
		return x.Footer
	}
	return ""
}

func (x *Template) GetFiles() []*executor.File {
	if x != nil {
		return x.Files
	}
	return nil
}

var File_eolymp_atlas_code_template_proto protoreflect.FileDescriptor

var file_eolymp_atlas_code_template_proto_rawDesc = []byte{
	0x0a, 0x20, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x61, 0x74, 0x6c, 0x61, 0x73, 0x2f, 0x63,
	0x6f, 0x64, 0x65, 0x5f, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0c, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x61, 0x74, 0x6c, 0x61, 0x73,
	0x1a, 0x1a, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x6f,
	0x72, 0x2f, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc1, 0x01, 0x0a,
	0x08, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x75, 0x6e,
	0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x72, 0x75, 0x6e, 0x74,
	0x69, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x06, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x0b, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x66,
	0x6f, 0x6f, 0x74, 0x65, 0x72, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x66, 0x6f, 0x6f,
	0x74, 0x65, 0x72, 0x12, 0x2b, 0x0a, 0x05, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x18, 0x1e, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x15, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x65, 0x78, 0x65, 0x63,
	0x75, 0x74, 0x6f, 0x72, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x05, 0x66, 0x69, 0x6c, 0x65, 0x73,
	0x42, 0x2d, 0x5a, 0x2b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65,
	0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x67, 0x6f, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x65, 0x6f, 0x6c,
	0x79, 0x6d, 0x70, 0x2f, 0x61, 0x74, 0x6c, 0x61, 0x73, 0x3b, 0x61, 0x74, 0x6c, 0x61, 0x73, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_eolymp_atlas_code_template_proto_rawDescOnce sync.Once
	file_eolymp_atlas_code_template_proto_rawDescData = file_eolymp_atlas_code_template_proto_rawDesc
)

func file_eolymp_atlas_code_template_proto_rawDescGZIP() []byte {
	file_eolymp_atlas_code_template_proto_rawDescOnce.Do(func() {
		file_eolymp_atlas_code_template_proto_rawDescData = protoimpl.X.CompressGZIP(file_eolymp_atlas_code_template_proto_rawDescData)
	})
	return file_eolymp_atlas_code_template_proto_rawDescData
}

var file_eolymp_atlas_code_template_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_eolymp_atlas_code_template_proto_goTypes = []any{
	(*Template)(nil),      // 0: eolymp.atlas.Template
	(*executor.File)(nil), // 1: eolymp.executor.File
}
var file_eolymp_atlas_code_template_proto_depIdxs = []int32{
	1, // 0: eolymp.atlas.Template.files:type_name -> eolymp.executor.File
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_eolymp_atlas_code_template_proto_init() }
func file_eolymp_atlas_code_template_proto_init() {
	if File_eolymp_atlas_code_template_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_eolymp_atlas_code_template_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_eolymp_atlas_code_template_proto_goTypes,
		DependencyIndexes: file_eolymp_atlas_code_template_proto_depIdxs,
		MessageInfos:      file_eolymp_atlas_code_template_proto_msgTypes,
	}.Build()
	File_eolymp_atlas_code_template_proto = out.File
	file_eolymp_atlas_code_template_proto_rawDesc = nil
	file_eolymp_atlas_code_template_proto_goTypes = nil
	file_eolymp_atlas_code_template_proto_depIdxs = nil
}
