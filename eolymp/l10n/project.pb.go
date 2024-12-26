// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.1
// 	protoc        v4.24.4
// source: eolymp/l10n/project.proto

package l10n

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

type Project struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Url           string                 `protobuf:"bytes,4,opt,name=url,proto3" json:"url,omitempty"`
	LogoUrl       string                 `protobuf:"bytes,3,opt,name=logo_url,json=logoUrl,proto3" json:"logo_url,omitempty"`
	HomeUrl       string                 `protobuf:"bytes,5,opt,name=home_url,json=homeUrl,proto3" json:"home_url,omitempty"`
	Description   string                 `protobuf:"bytes,6,opt,name=description,proto3" json:"description,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Project) Reset() {
	*x = Project{}
	mi := &file_eolymp_l10n_project_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Project) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Project) ProtoMessage() {}

func (x *Project) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_l10n_project_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Project.ProtoReflect.Descriptor instead.
func (*Project) Descriptor() ([]byte, []int) {
	return file_eolymp_l10n_project_proto_rawDescGZIP(), []int{0}
}

func (x *Project) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Project) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Project) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *Project) GetLogoUrl() string {
	if x != nil {
		return x.LogoUrl
	}
	return ""
}

func (x *Project) GetHomeUrl() string {
	if x != nil {
		return x.HomeUrl
	}
	return ""
}

func (x *Project) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

var File_eolymp_l10n_project_proto protoreflect.FileDescriptor

var file_eolymp_l10n_project_proto_rawDesc = []byte{
	0x0a, 0x19, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x6c, 0x31, 0x30, 0x6e, 0x2f, 0x70, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x65, 0x6f, 0x6c,
	0x79, 0x6d, 0x70, 0x2e, 0x6c, 0x31, 0x30, 0x6e, 0x22, 0x97, 0x01, 0x0a, 0x07, 0x50, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x19, 0x0a, 0x08, 0x6c, 0x6f,
	0x67, 0x6f, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6c, 0x6f,
	0x67, 0x6f, 0x55, 0x72, 0x6c, 0x12, 0x19, 0x0a, 0x08, 0x68, 0x6f, 0x6d, 0x65, 0x5f, 0x75, 0x72,
	0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x68, 0x6f, 0x6d, 0x65, 0x55, 0x72, 0x6c,
	0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x42, 0x2b, 0x5a, 0x29, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x67, 0x6f, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x65,
	0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x6c, 0x31, 0x30, 0x6e, 0x3b, 0x6c, 0x31, 0x30, 0x6e, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_eolymp_l10n_project_proto_rawDescOnce sync.Once
	file_eolymp_l10n_project_proto_rawDescData = file_eolymp_l10n_project_proto_rawDesc
)

func file_eolymp_l10n_project_proto_rawDescGZIP() []byte {
	file_eolymp_l10n_project_proto_rawDescOnce.Do(func() {
		file_eolymp_l10n_project_proto_rawDescData = protoimpl.X.CompressGZIP(file_eolymp_l10n_project_proto_rawDescData)
	})
	return file_eolymp_l10n_project_proto_rawDescData
}

var file_eolymp_l10n_project_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_eolymp_l10n_project_proto_goTypes = []any{
	(*Project)(nil), // 0: eolymp.l10n.Project
}
var file_eolymp_l10n_project_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_eolymp_l10n_project_proto_init() }
func file_eolymp_l10n_project_proto_init() {
	if File_eolymp_l10n_project_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_eolymp_l10n_project_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_eolymp_l10n_project_proto_goTypes,
		DependencyIndexes: file_eolymp_l10n_project_proto_depIdxs,
		MessageInfos:      file_eolymp_l10n_project_proto_msgTypes,
	}.Build()
	File_eolymp_l10n_project_proto = out.File
	file_eolymp_l10n_project_proto_rawDesc = nil
	file_eolymp_l10n_project_proto_goTypes = nil
	file_eolymp_l10n_project_proto_depIdxs = nil
}
