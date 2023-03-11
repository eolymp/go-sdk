// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.29.0
// 	protoc        v3.18.1
// source: eolymp/drive/file.proto

package drive

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type File struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Path       string                 `protobuf:"bytes,2,opt,name=path,proto3" json:"path,omitempty"`  // File path
	Hash       string                 `protobuf:"bytes,3,opt,name=hash,proto3" json:"hash,omitempty"`  // SHA1 hash of data
	Size       uint32                 `protobuf:"varint,4,opt,name=size,proto3" json:"size,omitempty"` // Size of the data
	Type       string                 `protobuf:"bytes,5,opt,name=type,proto3" json:"type,omitempty"`  // File mime type
	CreatedAt  *timestamppb.Timestamp `protobuf:"bytes,10,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt  *timestamppb.Timestamp `protobuf:"bytes,11,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	Attributes map[string]string      `protobuf:"bytes,20,rep,name=attributes,proto3" json:"attributes,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Url        string                 `protobuf:"bytes,100,opt,name=url,proto3" json:"url,omitempty"`
}

func (x *File) Reset() {
	*x = File{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_drive_file_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *File) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*File) ProtoMessage() {}

func (x *File) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_drive_file_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use File.ProtoReflect.Descriptor instead.
func (*File) Descriptor() ([]byte, []int) {
	return file_eolymp_drive_file_proto_rawDescGZIP(), []int{0}
}

func (x *File) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *File) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *File) GetHash() string {
	if x != nil {
		return x.Hash
	}
	return ""
}

func (x *File) GetSize() uint32 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *File) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *File) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *File) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *File) GetAttributes() map[string]string {
	if x != nil {
		return x.Attributes
	}
	return nil
}

func (x *File) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

var File_eolymp_drive_file_proto protoreflect.FileDescriptor

var file_eolymp_drive_file_proto_rawDesc = []byte{
	0x0a, 0x17, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x64, 0x72, 0x69, 0x76, 0x65, 0x2f, 0x66,
	0x69, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x65, 0x6f, 0x6c, 0x79, 0x6d,
	0x70, 0x2e, 0x64, 0x72, 0x69, 0x76, 0x65, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf1, 0x02, 0x0a, 0x04, 0x46, 0x69, 0x6c,
	0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x61, 0x73, 0x68, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x61, 0x73, 0x68, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x39, 0x0a, 0x0a,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x42, 0x0a, 0x0a, 0x61, 0x74, 0x74, 0x72, 0x69,
	0x62, 0x75, 0x74, 0x65, 0x73, 0x18, 0x14, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x65, 0x6f,
	0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x64, 0x72, 0x69, 0x76, 0x65, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x2e,
	0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52,
	0x0a, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x75,
	0x72, 0x6c, 0x18, 0x64, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x1a, 0x3d, 0x0a,
	0x0f, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x2d, 0x5a, 0x2b,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6f, 0x6c, 0x79, 0x6d,
	0x70, 0x2f, 0x67, 0x6f, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f,
	0x64, 0x72, 0x69, 0x76, 0x65, 0x3b, 0x64, 0x72, 0x69, 0x76, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_eolymp_drive_file_proto_rawDescOnce sync.Once
	file_eolymp_drive_file_proto_rawDescData = file_eolymp_drive_file_proto_rawDesc
)

func file_eolymp_drive_file_proto_rawDescGZIP() []byte {
	file_eolymp_drive_file_proto_rawDescOnce.Do(func() {
		file_eolymp_drive_file_proto_rawDescData = protoimpl.X.CompressGZIP(file_eolymp_drive_file_proto_rawDescData)
	})
	return file_eolymp_drive_file_proto_rawDescData
}

var file_eolymp_drive_file_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_eolymp_drive_file_proto_goTypes = []interface{}{
	(*File)(nil),                  // 0: eolymp.drive.File
	nil,                           // 1: eolymp.drive.File.AttributesEntry
	(*timestamppb.Timestamp)(nil), // 2: google.protobuf.Timestamp
}
var file_eolymp_drive_file_proto_depIdxs = []int32{
	2, // 0: eolymp.drive.File.created_at:type_name -> google.protobuf.Timestamp
	2, // 1: eolymp.drive.File.updated_at:type_name -> google.protobuf.Timestamp
	1, // 2: eolymp.drive.File.attributes:type_name -> eolymp.drive.File.AttributesEntry
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_eolymp_drive_file_proto_init() }
func file_eolymp_drive_file_proto_init() {
	if File_eolymp_drive_file_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_eolymp_drive_file_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*File); i {
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
			RawDescriptor: file_eolymp_drive_file_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_eolymp_drive_file_proto_goTypes,
		DependencyIndexes: file_eolymp_drive_file_proto_depIdxs,
		MessageInfos:      file_eolymp_drive_file_proto_msgTypes,
	}.Build()
	File_eolymp_drive_file_proto = out.File
	file_eolymp_drive_file_proto_rawDesc = nil
	file_eolymp_drive_file_proto_goTypes = nil
	file_eolymp_drive_file_proto_depIdxs = nil
}
