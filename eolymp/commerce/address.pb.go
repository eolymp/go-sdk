// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.4
// 	protoc        v4.24.4
// source: eolymp/commerce/address.proto

package commerce

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

type Address struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Country       string                 `protobuf:"bytes,200,opt,name=country,proto3" json:"country,omitempty"` // two letter code, lowercase
	State         string                 `protobuf:"bytes,201,opt,name=state,proto3" json:"state,omitempty"`
	PostalCode    string                 `protobuf:"bytes,202,opt,name=postal_code,json=postalCode,proto3" json:"postal_code,omitempty"`
	City          string                 `protobuf:"bytes,203,opt,name=city,proto3" json:"city,omitempty"`
	Line1         string                 `protobuf:"bytes,210,opt,name=line1,proto3" json:"line1,omitempty"`
	Line2         string                 `protobuf:"bytes,211,opt,name=line2,proto3" json:"line2,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Address) Reset() {
	*x = Address{}
	mi := &file_eolymp_commerce_address_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Address) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Address) ProtoMessage() {}

func (x *Address) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_commerce_address_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Address.ProtoReflect.Descriptor instead.
func (*Address) Descriptor() ([]byte, []int) {
	return file_eolymp_commerce_address_proto_rawDescGZIP(), []int{0}
}

func (x *Address) GetCountry() string {
	if x != nil {
		return x.Country
	}
	return ""
}

func (x *Address) GetState() string {
	if x != nil {
		return x.State
	}
	return ""
}

func (x *Address) GetPostalCode() string {
	if x != nil {
		return x.PostalCode
	}
	return ""
}

func (x *Address) GetCity() string {
	if x != nil {
		return x.City
	}
	return ""
}

func (x *Address) GetLine1() string {
	if x != nil {
		return x.Line1
	}
	return ""
}

func (x *Address) GetLine2() string {
	if x != nil {
		return x.Line2
	}
	return ""
}

var File_eolymp_commerce_address_proto protoreflect.FileDescriptor

var file_eolymp_commerce_address_proto_rawDesc = string([]byte{
	0x0a, 0x1d, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x63,
	0x65, 0x2f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0f, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x63, 0x65,
	0x22, 0xa0, 0x01, 0x0a, 0x07, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x19, 0x0a, 0x07,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x18, 0xc8, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x15, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65,
	0x18, 0xc9, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x20,
	0x0a, 0x0b, 0x70, 0x6f, 0x73, 0x74, 0x61, 0x6c, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0xca, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x6f, 0x73, 0x74, 0x61, 0x6c, 0x43, 0x6f, 0x64, 0x65,
	0x12, 0x13, 0x0a, 0x04, 0x63, 0x69, 0x74, 0x79, 0x18, 0xcb, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x63, 0x69, 0x74, 0x79, 0x12, 0x15, 0x0a, 0x05, 0x6c, 0x69, 0x6e, 0x65, 0x31, 0x18, 0xd2,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x69, 0x6e, 0x65, 0x31, 0x12, 0x15, 0x0a, 0x05,
	0x6c, 0x69, 0x6e, 0x65, 0x32, 0x18, 0xd3, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x69,
	0x6e, 0x65, 0x32, 0x42, 0x33, 0x5a, 0x31, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x67, 0x6f, 0x2d, 0x73, 0x64, 0x6b, 0x2f,
	0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x63, 0x65, 0x3b,
	0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_eolymp_commerce_address_proto_rawDescOnce sync.Once
	file_eolymp_commerce_address_proto_rawDescData []byte
)

func file_eolymp_commerce_address_proto_rawDescGZIP() []byte {
	file_eolymp_commerce_address_proto_rawDescOnce.Do(func() {
		file_eolymp_commerce_address_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_eolymp_commerce_address_proto_rawDesc), len(file_eolymp_commerce_address_proto_rawDesc)))
	})
	return file_eolymp_commerce_address_proto_rawDescData
}

var file_eolymp_commerce_address_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_eolymp_commerce_address_proto_goTypes = []any{
	(*Address)(nil), // 0: eolymp.commerce.Address
}
var file_eolymp_commerce_address_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_eolymp_commerce_address_proto_init() }
func file_eolymp_commerce_address_proto_init() {
	if File_eolymp_commerce_address_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_eolymp_commerce_address_proto_rawDesc), len(file_eolymp_commerce_address_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_eolymp_commerce_address_proto_goTypes,
		DependencyIndexes: file_eolymp_commerce_address_proto_depIdxs,
		MessageInfos:      file_eolymp_commerce_address_proto_msgTypes,
	}.Build()
	File_eolymp_commerce_address_proto = out.File
	file_eolymp_commerce_address_proto_goTypes = nil
	file_eolymp_commerce_address_proto_depIdxs = nil
}
