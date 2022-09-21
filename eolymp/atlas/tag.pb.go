// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.18.1
// source: eolymp/atlas/tag.proto

package atlas

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

type Tag struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           string             `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Visible      bool               `protobuf:"varint,10,opt,name=visible,proto3" json:"visible,omitempty"`
	Descriptions []*Tag_Description `protobuf:"bytes,30,rep,name=descriptions,proto3" json:"descriptions,omitempty"`
}

func (x *Tag) Reset() {
	*x = Tag{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_atlas_tag_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Tag) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Tag) ProtoMessage() {}

func (x *Tag) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_atlas_tag_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Tag.ProtoReflect.Descriptor instead.
func (*Tag) Descriptor() ([]byte, []int) {
	return file_eolymp_atlas_tag_proto_rawDescGZIP(), []int{0}
}

func (x *Tag) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Tag) GetVisible() bool {
	if x != nil {
		return x.Visible
	}
	return false
}

func (x *Tag) GetDescriptions() []*Tag_Description {
	if x != nil {
		return x.Descriptions
	}
	return nil
}

type Tag_Description struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Locale  string `protobuf:"bytes,1,opt,name=locale,proto3" json:"locale,omitempty"`
	Name    string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Summary string `protobuf:"bytes,3,opt,name=summary,proto3" json:"summary,omitempty"`
}

func (x *Tag_Description) Reset() {
	*x = Tag_Description{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_atlas_tag_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Tag_Description) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Tag_Description) ProtoMessage() {}

func (x *Tag_Description) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_atlas_tag_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Tag_Description.ProtoReflect.Descriptor instead.
func (*Tag_Description) Descriptor() ([]byte, []int) {
	return file_eolymp_atlas_tag_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Tag_Description) GetLocale() string {
	if x != nil {
		return x.Locale
	}
	return ""
}

func (x *Tag_Description) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Tag_Description) GetSummary() string {
	if x != nil {
		return x.Summary
	}
	return ""
}

var File_eolymp_atlas_tag_proto protoreflect.FileDescriptor

var file_eolymp_atlas_tag_proto_rawDesc = []byte{
	0x0a, 0x16, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x61, 0x74, 0x6c, 0x61, 0x73, 0x2f, 0x74,
	0x61, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70,
	0x2e, 0x61, 0x74, 0x6c, 0x61, 0x73, 0x22, 0xc7, 0x01, 0x0a, 0x03, 0x54, 0x61, 0x67, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x18,
	0x0a, 0x07, 0x76, 0x69, 0x73, 0x69, 0x62, 0x6c, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x07, 0x76, 0x69, 0x73, 0x69, 0x62, 0x6c, 0x65, 0x12, 0x41, 0x0a, 0x0c, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x1e, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d,
	0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x61, 0x74, 0x6c, 0x61, 0x73, 0x2e, 0x54, 0x61,
	0x67, 0x2e, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0c, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x1a, 0x53, 0x0a, 0x0b, 0x44,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x6c, 0x6f,
	0x63, 0x61, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6c, 0x6f, 0x63, 0x61,
	0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72,
	0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79,
	0x42, 0x2d, 0x5a, 0x2b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65,
	0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x67, 0x6f, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x65, 0x6f, 0x6c,
	0x79, 0x6d, 0x70, 0x2f, 0x61, 0x74, 0x6c, 0x61, 0x73, 0x3b, 0x61, 0x74, 0x6c, 0x61, 0x73, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_eolymp_atlas_tag_proto_rawDescOnce sync.Once
	file_eolymp_atlas_tag_proto_rawDescData = file_eolymp_atlas_tag_proto_rawDesc
)

func file_eolymp_atlas_tag_proto_rawDescGZIP() []byte {
	file_eolymp_atlas_tag_proto_rawDescOnce.Do(func() {
		file_eolymp_atlas_tag_proto_rawDescData = protoimpl.X.CompressGZIP(file_eolymp_atlas_tag_proto_rawDescData)
	})
	return file_eolymp_atlas_tag_proto_rawDescData
}

var file_eolymp_atlas_tag_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_eolymp_atlas_tag_proto_goTypes = []interface{}{
	(*Tag)(nil),             // 0: eolymp.atlas.Tag
	(*Tag_Description)(nil), // 1: eolymp.atlas.Tag.Description
}
var file_eolymp_atlas_tag_proto_depIdxs = []int32{
	1, // 0: eolymp.atlas.Tag.descriptions:type_name -> eolymp.atlas.Tag.Description
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_eolymp_atlas_tag_proto_init() }
func file_eolymp_atlas_tag_proto_init() {
	if File_eolymp_atlas_tag_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_eolymp_atlas_tag_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Tag); i {
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
		file_eolymp_atlas_tag_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Tag_Description); i {
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
			RawDescriptor: file_eolymp_atlas_tag_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_eolymp_atlas_tag_proto_goTypes,
		DependencyIndexes: file_eolymp_atlas_tag_proto_depIdxs,
		MessageInfos:      file_eolymp_atlas_tag_proto_msgTypes,
	}.Build()
	File_eolymp_atlas_tag_proto = out.File
	file_eolymp_atlas_tag_proto_rawDesc = nil
	file_eolymp_atlas_tag_proto_goTypes = nil
	file_eolymp_atlas_tag_proto_depIdxs = nil
}
