// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v3.21.12
// source: eolymp/content/fragment.proto

package content

import (
	ecm "github.com/eolymp/go-sdk/eolymp/ecm"
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

type Fragment_Extra int32

const (
	Fragment_NO_EXTRA       Fragment_Extra = 0
	Fragment_CONTENT_RENDER Fragment_Extra = 1
	Fragment_CONTENT_VALUE  Fragment_Extra = 2
)

// Enum value maps for Fragment_Extra.
var (
	Fragment_Extra_name = map[int32]string{
		0: "NO_EXTRA",
		1: "CONTENT_RENDER",
		2: "CONTENT_VALUE",
	}
	Fragment_Extra_value = map[string]int32{
		"NO_EXTRA":       0,
		"CONTENT_RENDER": 1,
		"CONTENT_VALUE":  2,
	}
)

func (x Fragment_Extra) Enum() *Fragment_Extra {
	p := new(Fragment_Extra)
	*p = x
	return p
}

func (x Fragment_Extra) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Fragment_Extra) Descriptor() protoreflect.EnumDescriptor {
	return file_eolymp_content_fragment_proto_enumTypes[0].Descriptor()
}

func (Fragment_Extra) Type() protoreflect.EnumType {
	return &file_eolymp_content_fragment_proto_enumTypes[0]
}

func (x Fragment_Extra) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Fragment_Extra.Descriptor instead.
func (Fragment_Extra) EnumDescriptor() ([]byte, []int) {
	return file_eolymp_content_fragment_proto_rawDescGZIP(), []int{0, 0}
}

type Fragment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Path      string                 `protobuf:"bytes,10,opt,name=path,proto3" json:"path,omitempty"`
	Locale    string                 `protobuf:"bytes,11,opt,name=locale,proto3" json:"locale,omitempty"`
	Title     string                 `protobuf:"bytes,12,opt,name=title,proto3" json:"title,omitempty"`
	Public    string                 `protobuf:"bytes,13,opt,name=public,proto3" json:"public,omitempty"` // true if fragment is public and available, otherwise it's considered to be draft
	Content   *ecm.Content           `protobuf:"bytes,51,opt,name=content,proto3" json:"content,omitempty"`
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,60,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt *timestamppb.Timestamp `protobuf:"bytes,61,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	Labels    []string               `protobuf:"bytes,100,rep,name=labels,proto3" json:"labels,omitempty"`
}

func (x *Fragment) Reset() {
	*x = Fragment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_content_fragment_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Fragment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Fragment) ProtoMessage() {}

func (x *Fragment) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_content_fragment_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Fragment.ProtoReflect.Descriptor instead.
func (*Fragment) Descriptor() ([]byte, []int) {
	return file_eolymp_content_fragment_proto_rawDescGZIP(), []int{0}
}

func (x *Fragment) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Fragment) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *Fragment) GetLocale() string {
	if x != nil {
		return x.Locale
	}
	return ""
}

func (x *Fragment) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Fragment) GetPublic() string {
	if x != nil {
		return x.Public
	}
	return ""
}

func (x *Fragment) GetContent() *ecm.Content {
	if x != nil {
		return x.Content
	}
	return nil
}

func (x *Fragment) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Fragment) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *Fragment) GetLabels() []string {
	if x != nil {
		return x.Labels
	}
	return nil
}

var File_eolymp_content_fragment_proto protoreflect.FileDescriptor

var file_eolymp_content_fragment_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x2f, 0x66, 0x72, 0x61, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x1a,
	0x18, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x65, 0x63, 0x6d, 0x2f, 0x63, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xef, 0x02, 0x0a, 0x08, 0x46,
	0x72, 0x61, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x16, 0x0a, 0x06, 0x6c,
	0x6f, 0x63, 0x61, 0x6c, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6c, 0x6f, 0x63,
	0x61, 0x6c, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x0c, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x75, 0x62,
	0x6c, 0x69, 0x63, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x75, 0x62, 0x6c, 0x69,
	0x63, 0x12, 0x2d, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x33, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x13, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x65, 0x63, 0x6d, 0x2e,
	0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x3c,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x3d, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73,
	0x18, 0x64, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x22, 0x3c,
	0x0a, 0x05, 0x45, 0x78, 0x74, 0x72, 0x61, 0x12, 0x0c, 0x0a, 0x08, 0x4e, 0x4f, 0x5f, 0x45, 0x58,
	0x54, 0x52, 0x41, 0x10, 0x00, 0x12, 0x12, 0x0a, 0x0e, 0x43, 0x4f, 0x4e, 0x54, 0x45, 0x4e, 0x54,
	0x5f, 0x52, 0x45, 0x4e, 0x44, 0x45, 0x52, 0x10, 0x01, 0x12, 0x11, 0x0a, 0x0d, 0x43, 0x4f, 0x4e,
	0x54, 0x45, 0x4e, 0x54, 0x5f, 0x56, 0x41, 0x4c, 0x55, 0x45, 0x10, 0x02, 0x42, 0x31, 0x5a, 0x2f,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6f, 0x6c, 0x79, 0x6d,
	0x70, 0x2f, 0x67, 0x6f, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f,
	0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x3b, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_eolymp_content_fragment_proto_rawDescOnce sync.Once
	file_eolymp_content_fragment_proto_rawDescData = file_eolymp_content_fragment_proto_rawDesc
)

func file_eolymp_content_fragment_proto_rawDescGZIP() []byte {
	file_eolymp_content_fragment_proto_rawDescOnce.Do(func() {
		file_eolymp_content_fragment_proto_rawDescData = protoimpl.X.CompressGZIP(file_eolymp_content_fragment_proto_rawDescData)
	})
	return file_eolymp_content_fragment_proto_rawDescData
}

var file_eolymp_content_fragment_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_eolymp_content_fragment_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_eolymp_content_fragment_proto_goTypes = []interface{}{
	(Fragment_Extra)(0),           // 0: eolymp.content.Fragment.Extra
	(*Fragment)(nil),              // 1: eolymp.content.Fragment
	(*ecm.Content)(nil),           // 2: eolymp.ecm.Content
	(*timestamppb.Timestamp)(nil), // 3: google.protobuf.Timestamp
}
var file_eolymp_content_fragment_proto_depIdxs = []int32{
	2, // 0: eolymp.content.Fragment.content:type_name -> eolymp.ecm.Content
	3, // 1: eolymp.content.Fragment.created_at:type_name -> google.protobuf.Timestamp
	3, // 2: eolymp.content.Fragment.updated_at:type_name -> google.protobuf.Timestamp
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_eolymp_content_fragment_proto_init() }
func file_eolymp_content_fragment_proto_init() {
	if File_eolymp_content_fragment_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_eolymp_content_fragment_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Fragment); i {
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
			RawDescriptor: file_eolymp_content_fragment_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_eolymp_content_fragment_proto_goTypes,
		DependencyIndexes: file_eolymp_content_fragment_proto_depIdxs,
		EnumInfos:         file_eolymp_content_fragment_proto_enumTypes,
		MessageInfos:      file_eolymp_content_fragment_proto_msgTypes,
	}.Build()
	File_eolymp_content_fragment_proto = out.File
	file_eolymp_content_fragment_proto_rawDesc = nil
	file_eolymp_content_fragment_proto_goTypes = nil
	file_eolymp_content_fragment_proto_depIdxs = nil
}
