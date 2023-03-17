// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.18.1
// source: eolymp/geography/country.proto

package geography

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

type Country struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Flag string `protobuf:"bytes,3,opt,name=flag,proto3" json:"flag,omitempty"`
}

func (x *Country) Reset() {
	*x = Country{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_geography_country_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Country) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Country) ProtoMessage() {}

func (x *Country) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_geography_country_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Country.ProtoReflect.Descriptor instead.
func (*Country) Descriptor() ([]byte, []int) {
	return file_eolymp_geography_country_proto_rawDescGZIP(), []int{0}
}

func (x *Country) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Country) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Country) GetFlag() string {
	if x != nil {
		return x.Flag
	}
	return ""
}

var File_eolymp_geography_country_proto protoreflect.FileDescriptor

var file_eolymp_geography_country_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x67, 0x65, 0x6f, 0x67, 0x72, 0x61, 0x70,
	0x68, 0x79, 0x2f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x10, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x67, 0x65, 0x6f, 0x67, 0x72, 0x61, 0x70,
	0x68, 0x79, 0x22, 0x41, 0x0a, 0x07, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x6c, 0x61, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x66, 0x6c, 0x61, 0x67, 0x42, 0x35, 0x5a, 0x33, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x67, 0x6f, 0x2d, 0x73, 0x64,
	0x6b, 0x2f, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x67, 0x65, 0x6f, 0x67, 0x72, 0x61, 0x70,
	0x68, 0x79, 0x3b, 0x67, 0x65, 0x6f, 0x67, 0x72, 0x61, 0x70, 0x68, 0x79, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_eolymp_geography_country_proto_rawDescOnce sync.Once
	file_eolymp_geography_country_proto_rawDescData = file_eolymp_geography_country_proto_rawDesc
)

func file_eolymp_geography_country_proto_rawDescGZIP() []byte {
	file_eolymp_geography_country_proto_rawDescOnce.Do(func() {
		file_eolymp_geography_country_proto_rawDescData = protoimpl.X.CompressGZIP(file_eolymp_geography_country_proto_rawDescData)
	})
	return file_eolymp_geography_country_proto_rawDescData
}

var file_eolymp_geography_country_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_eolymp_geography_country_proto_goTypes = []interface{}{
	(*Country)(nil), // 0: eolymp.geography.Country
}
var file_eolymp_geography_country_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_eolymp_geography_country_proto_init() }
func file_eolymp_geography_country_proto_init() {
	if File_eolymp_geography_country_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_eolymp_geography_country_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Country); i {
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
			RawDescriptor: file_eolymp_geography_country_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_eolymp_geography_country_proto_goTypes,
		DependencyIndexes: file_eolymp_geography_country_proto_depIdxs,
		MessageInfos:      file_eolymp_geography_country_proto_msgTypes,
	}.Build()
	File_eolymp_geography_country_proto = out.File
	file_eolymp_geography_country_proto_rawDesc = nil
	file_eolymp_geography_country_proto_goTypes = nil
	file_eolymp_geography_country_proto_depIdxs = nil
}
