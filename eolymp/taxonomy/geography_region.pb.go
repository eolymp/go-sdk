// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v4.24.4
// source: eolymp/taxonomy/geography_region.proto

package taxonomy

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

type Region struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	CountryId     string                 `protobuf:"bytes,2,opt,name=country_id,json=countryId,proto3" json:"country_id,omitempty"`
	Name          string                 `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	LocalName     string                 `protobuf:"bytes,4,opt,name=local_name,json=localName,proto3" json:"local_name,omitempty"`
	FlagUrl       string                 `protobuf:"bytes,5,opt,name=flag_url,json=flagUrl,proto3" json:"flag_url,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Region) Reset() {
	*x = Region{}
	mi := &file_eolymp_taxonomy_geography_region_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Region) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Region) ProtoMessage() {}

func (x *Region) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_taxonomy_geography_region_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Region.ProtoReflect.Descriptor instead.
func (*Region) Descriptor() ([]byte, []int) {
	return file_eolymp_taxonomy_geography_region_proto_rawDescGZIP(), []int{0}
}

func (x *Region) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Region) GetCountryId() string {
	if x != nil {
		return x.CountryId
	}
	return ""
}

func (x *Region) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Region) GetLocalName() string {
	if x != nil {
		return x.LocalName
	}
	return ""
}

func (x *Region) GetFlagUrl() string {
	if x != nil {
		return x.FlagUrl
	}
	return ""
}

var File_eolymp_taxonomy_geography_region_proto protoreflect.FileDescriptor

const file_eolymp_taxonomy_geography_region_proto_rawDesc = "" +
	"\n" +
	"&eolymp/taxonomy/geography_region.proto\x12\x0feolymp.taxonomy\"\x85\x01\n" +
	"\x06Region\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x12\x1d\n" +
	"\n" +
	"country_id\x18\x02 \x01(\tR\tcountryId\x12\x12\n" +
	"\x04name\x18\x03 \x01(\tR\x04name\x12\x1d\n" +
	"\n" +
	"local_name\x18\x04 \x01(\tR\tlocalName\x12\x19\n" +
	"\bflag_url\x18\x05 \x01(\tR\aflagUrlB3Z1github.com/eolymp/go-sdk/eolymp/taxonomy;taxonomyb\x06proto3"

var (
	file_eolymp_taxonomy_geography_region_proto_rawDescOnce sync.Once
	file_eolymp_taxonomy_geography_region_proto_rawDescData []byte
)

func file_eolymp_taxonomy_geography_region_proto_rawDescGZIP() []byte {
	file_eolymp_taxonomy_geography_region_proto_rawDescOnce.Do(func() {
		file_eolymp_taxonomy_geography_region_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_eolymp_taxonomy_geography_region_proto_rawDesc), len(file_eolymp_taxonomy_geography_region_proto_rawDesc)))
	})
	return file_eolymp_taxonomy_geography_region_proto_rawDescData
}

var file_eolymp_taxonomy_geography_region_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_eolymp_taxonomy_geography_region_proto_goTypes = []any{
	(*Region)(nil), // 0: eolymp.taxonomy.Region
}
var file_eolymp_taxonomy_geography_region_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_eolymp_taxonomy_geography_region_proto_init() }
func file_eolymp_taxonomy_geography_region_proto_init() {
	if File_eolymp_taxonomy_geography_region_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_eolymp_taxonomy_geography_region_proto_rawDesc), len(file_eolymp_taxonomy_geography_region_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_eolymp_taxonomy_geography_region_proto_goTypes,
		DependencyIndexes: file_eolymp_taxonomy_geography_region_proto_depIdxs,
		MessageInfos:      file_eolymp_taxonomy_geography_region_proto_msgTypes,
	}.Build()
	File_eolymp_taxonomy_geography_region_proto = out.File
	file_eolymp_taxonomy_geography_region_proto_goTypes = nil
	file_eolymp_taxonomy_geography_region_proto_depIdxs = nil
}
