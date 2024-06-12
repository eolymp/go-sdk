// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.21.12
// source: eolymp/keeper/blob.proto

package keeper

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

type Blob struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ern  string `protobuf:"bytes,1,opt,name=ern,proto3" json:"ern,omitempty"`
	Key  string `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`    // Unique identifier of the blob
	Hash string `protobuf:"bytes,3,opt,name=hash,proto3" json:"hash,omitempty"`  // SHA1 hash of data
	Size uint32 `protobuf:"varint,4,opt,name=size,proto3" json:"size,omitempty"` // Size of the data
	Data []byte `protobuf:"bytes,100,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *Blob) Reset() {
	*x = Blob{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_keeper_blob_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Blob) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Blob) ProtoMessage() {}

func (x *Blob) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_keeper_blob_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Blob.ProtoReflect.Descriptor instead.
func (*Blob) Descriptor() ([]byte, []int) {
	return file_eolymp_keeper_blob_proto_rawDescGZIP(), []int{0}
}

func (x *Blob) GetErn() string {
	if x != nil {
		return x.Ern
	}
	return ""
}

func (x *Blob) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *Blob) GetHash() string {
	if x != nil {
		return x.Hash
	}
	return ""
}

func (x *Blob) GetSize() uint32 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *Blob) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_eolymp_keeper_blob_proto protoreflect.FileDescriptor

var file_eolymp_keeper_blob_proto_rawDesc = []byte{
	0x0a, 0x18, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x6b, 0x65, 0x65, 0x70, 0x65, 0x72, 0x2f,
	0x62, 0x6c, 0x6f, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x65, 0x6f, 0x6c, 0x79,
	0x6d, 0x70, 0x2e, 0x6b, 0x65, 0x65, 0x70, 0x65, 0x72, 0x22, 0x66, 0x0a, 0x04, 0x42, 0x6c, 0x6f,
	0x62, 0x12, 0x10, 0x0a, 0x03, 0x65, 0x72, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x65, 0x72, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x61, 0x73, 0x68, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x61, 0x73, 0x68, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x64, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x42, 0x2f, 0x5a, 0x2d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x67, 0x6f, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x65, 0x6f,
	0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x6b, 0x65, 0x65, 0x70, 0x65, 0x72, 0x3b, 0x6b, 0x65, 0x65, 0x70,
	0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_eolymp_keeper_blob_proto_rawDescOnce sync.Once
	file_eolymp_keeper_blob_proto_rawDescData = file_eolymp_keeper_blob_proto_rawDesc
)

func file_eolymp_keeper_blob_proto_rawDescGZIP() []byte {
	file_eolymp_keeper_blob_proto_rawDescOnce.Do(func() {
		file_eolymp_keeper_blob_proto_rawDescData = protoimpl.X.CompressGZIP(file_eolymp_keeper_blob_proto_rawDescData)
	})
	return file_eolymp_keeper_blob_proto_rawDescData
}

var file_eolymp_keeper_blob_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_eolymp_keeper_blob_proto_goTypes = []any{
	(*Blob)(nil), // 0: eolymp.keeper.Blob
}
var file_eolymp_keeper_blob_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_eolymp_keeper_blob_proto_init() }
func file_eolymp_keeper_blob_proto_init() {
	if File_eolymp_keeper_blob_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_eolymp_keeper_blob_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*Blob); i {
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
			RawDescriptor: file_eolymp_keeper_blob_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_eolymp_keeper_blob_proto_goTypes,
		DependencyIndexes: file_eolymp_keeper_blob_proto_depIdxs,
		MessageInfos:      file_eolymp_keeper_blob_proto_msgTypes,
	}.Build()
	File_eolymp_keeper_blob_proto = out.File
	file_eolymp_keeper_blob_proto_rawDesc = nil
	file_eolymp_keeper_blob_proto_goTypes = nil
	file_eolymp_keeper_blob_proto_depIdxs = nil
}
