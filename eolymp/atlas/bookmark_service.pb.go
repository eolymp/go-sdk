// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.0
// 	protoc        v3.21.12
// source: eolymp/atlas/bookmark_service.proto

package atlas

import (
	_ "github.com/eolymp/go-sdk/eolymp/annotations"
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

type SetBookmarkInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProblemId string `protobuf:"bytes,1,opt,name=problem_id,json=problemId,proto3" json:"problem_id,omitempty"`
	Bookmark  bool   `protobuf:"varint,2,opt,name=bookmark,proto3" json:"bookmark,omitempty"`
}

func (x *SetBookmarkInput) Reset() {
	*x = SetBookmarkInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_atlas_bookmark_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetBookmarkInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetBookmarkInput) ProtoMessage() {}

func (x *SetBookmarkInput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_atlas_bookmark_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetBookmarkInput.ProtoReflect.Descriptor instead.
func (*SetBookmarkInput) Descriptor() ([]byte, []int) {
	return file_eolymp_atlas_bookmark_service_proto_rawDescGZIP(), []int{0}
}

func (x *SetBookmarkInput) GetProblemId() string {
	if x != nil {
		return x.ProblemId
	}
	return ""
}

func (x *SetBookmarkInput) GetBookmark() bool {
	if x != nil {
		return x.Bookmark
	}
	return false
}

type SetBookmarkOutput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SetBookmarkOutput) Reset() {
	*x = SetBookmarkOutput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_atlas_bookmark_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetBookmarkOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetBookmarkOutput) ProtoMessage() {}

func (x *SetBookmarkOutput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_atlas_bookmark_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetBookmarkOutput.ProtoReflect.Descriptor instead.
func (*SetBookmarkOutput) Descriptor() ([]byte, []int) {
	return file_eolymp_atlas_bookmark_service_proto_rawDescGZIP(), []int{1}
}

type GetBookmarkInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProblemId string `protobuf:"bytes,1,opt,name=problem_id,json=problemId,proto3" json:"problem_id,omitempty"`
}

func (x *GetBookmarkInput) Reset() {
	*x = GetBookmarkInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_atlas_bookmark_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBookmarkInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBookmarkInput) ProtoMessage() {}

func (x *GetBookmarkInput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_atlas_bookmark_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBookmarkInput.ProtoReflect.Descriptor instead.
func (*GetBookmarkInput) Descriptor() ([]byte, []int) {
	return file_eolymp_atlas_bookmark_service_proto_rawDescGZIP(), []int{2}
}

func (x *GetBookmarkInput) GetProblemId() string {
	if x != nil {
		return x.ProblemId
	}
	return ""
}

type GetBookmarkOutput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Bookmark bool `protobuf:"varint,1,opt,name=bookmark,proto3" json:"bookmark,omitempty"`
}

func (x *GetBookmarkOutput) Reset() {
	*x = GetBookmarkOutput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_atlas_bookmark_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBookmarkOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBookmarkOutput) ProtoMessage() {}

func (x *GetBookmarkOutput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_atlas_bookmark_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBookmarkOutput.ProtoReflect.Descriptor instead.
func (*GetBookmarkOutput) Descriptor() ([]byte, []int) {
	return file_eolymp_atlas_bookmark_service_proto_rawDescGZIP(), []int{3}
}

func (x *GetBookmarkOutput) GetBookmark() bool {
	if x != nil {
		return x.Bookmark
	}
	return false
}

var File_eolymp_atlas_bookmark_service_proto protoreflect.FileDescriptor

var file_eolymp_atlas_bookmark_service_proto_rawDesc = []byte{
	0x0a, 0x23, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x61, 0x74, 0x6c, 0x61, 0x73, 0x2f, 0x62,
	0x6f, 0x6f, 0x6b, 0x6d, 0x61, 0x72, 0x6b, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x61, 0x74,
	0x6c, 0x61, 0x73, 0x1a, 0x1d, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x22, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x72, 0x61, 0x74, 0x65, 0x6c, 0x69, 0x6d, 0x69, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4d, 0x0a, 0x10, 0x53, 0x65, 0x74, 0x42, 0x6f, 0x6f,
	0x6b, 0x6d, 0x61, 0x72, 0x6b, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72,
	0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x70, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x62, 0x6f, 0x6f,
	0x6b, 0x6d, 0x61, 0x72, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x62, 0x6f, 0x6f,
	0x6b, 0x6d, 0x61, 0x72, 0x6b, 0x22, 0x13, 0x0a, 0x11, 0x53, 0x65, 0x74, 0x42, 0x6f, 0x6f, 0x6b,
	0x6d, 0x61, 0x72, 0x6b, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x22, 0x31, 0x0a, 0x10, 0x47, 0x65,
	0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x6d, 0x61, 0x72, 0x6b, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x1d,
	0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x49, 0x64, 0x22, 0x2f, 0x0a,
	0x11, 0x47, 0x65, 0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x6d, 0x61, 0x72, 0x6b, 0x4f, 0x75, 0x74, 0x70,
	0x75, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x62, 0x6f, 0x6f, 0x6b, 0x6d, 0x61, 0x72, 0x6b, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x62, 0x6f, 0x6f, 0x6b, 0x6d, 0x61, 0x72, 0x6b, 0x32, 0xf5,
	0x01, 0x0a, 0x0f, 0x42, 0x6f, 0x6f, 0x6b, 0x6d, 0x61, 0x72, 0x6b, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x70, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x6d, 0x61, 0x72,
	0x6b, 0x12, 0x1e, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x61, 0x74, 0x6c, 0x61, 0x73,
	0x2e, 0x47, 0x65, 0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x6d, 0x61, 0x72, 0x6b, 0x49, 0x6e, 0x70, 0x75,
	0x74, 0x1a, 0x1f, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x61, 0x74, 0x6c, 0x61, 0x73,
	0x2e, 0x47, 0x65, 0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x6d, 0x61, 0x72, 0x6b, 0x4f, 0x75, 0x74, 0x70,
	0x75, 0x74, 0x22, 0x20, 0xea, 0xe2, 0x0a, 0x0b, 0xf5, 0xe2, 0x0a, 0x00, 0x00, 0xc8, 0x41, 0xf8,
	0xe2, 0x0a, 0x64, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0b, 0x12, 0x09, 0x2f, 0x62, 0x6f, 0x6f, 0x6b,
	0x6d, 0x61, 0x72, 0x6b, 0x12, 0x70, 0x0a, 0x0b, 0x53, 0x65, 0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x6d,
	0x61, 0x72, 0x6b, 0x12, 0x1e, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x61, 0x74, 0x6c,
	0x61, 0x73, 0x2e, 0x53, 0x65, 0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x6d, 0x61, 0x72, 0x6b, 0x49, 0x6e,
	0x70, 0x75, 0x74, 0x1a, 0x1f, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x61, 0x74, 0x6c,
	0x61, 0x73, 0x2e, 0x53, 0x65, 0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x6d, 0x61, 0x72, 0x6b, 0x4f, 0x75,
	0x74, 0x70, 0x75, 0x74, 0x22, 0x20, 0xea, 0xe2, 0x0a, 0x0b, 0xf5, 0xe2, 0x0a, 0x00, 0x00, 0x00,
	0x40, 0xf8, 0xe2, 0x0a, 0x0a, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0b, 0x22, 0x09, 0x2f, 0x62, 0x6f,
	0x6f, 0x6b, 0x6d, 0x61, 0x72, 0x6b, 0x42, 0x2d, 0x5a, 0x2b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x67, 0x6f, 0x2d, 0x73,
	0x64, 0x6b, 0x2f, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x61, 0x74, 0x6c, 0x61, 0x73, 0x3b,
	0x61, 0x74, 0x6c, 0x61, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_eolymp_atlas_bookmark_service_proto_rawDescOnce sync.Once
	file_eolymp_atlas_bookmark_service_proto_rawDescData = file_eolymp_atlas_bookmark_service_proto_rawDesc
)

func file_eolymp_atlas_bookmark_service_proto_rawDescGZIP() []byte {
	file_eolymp_atlas_bookmark_service_proto_rawDescOnce.Do(func() {
		file_eolymp_atlas_bookmark_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_eolymp_atlas_bookmark_service_proto_rawDescData)
	})
	return file_eolymp_atlas_bookmark_service_proto_rawDescData
}

var file_eolymp_atlas_bookmark_service_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_eolymp_atlas_bookmark_service_proto_goTypes = []interface{}{
	(*SetBookmarkInput)(nil),  // 0: eolymp.atlas.SetBookmarkInput
	(*SetBookmarkOutput)(nil), // 1: eolymp.atlas.SetBookmarkOutput
	(*GetBookmarkInput)(nil),  // 2: eolymp.atlas.GetBookmarkInput
	(*GetBookmarkOutput)(nil), // 3: eolymp.atlas.GetBookmarkOutput
}
var file_eolymp_atlas_bookmark_service_proto_depIdxs = []int32{
	2, // 0: eolymp.atlas.BookmarkService.GetBookmark:input_type -> eolymp.atlas.GetBookmarkInput
	0, // 1: eolymp.atlas.BookmarkService.SetBookmark:input_type -> eolymp.atlas.SetBookmarkInput
	3, // 2: eolymp.atlas.BookmarkService.GetBookmark:output_type -> eolymp.atlas.GetBookmarkOutput
	1, // 3: eolymp.atlas.BookmarkService.SetBookmark:output_type -> eolymp.atlas.SetBookmarkOutput
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_eolymp_atlas_bookmark_service_proto_init() }
func file_eolymp_atlas_bookmark_service_proto_init() {
	if File_eolymp_atlas_bookmark_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_eolymp_atlas_bookmark_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetBookmarkInput); i {
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
		file_eolymp_atlas_bookmark_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetBookmarkOutput); i {
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
		file_eolymp_atlas_bookmark_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBookmarkInput); i {
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
		file_eolymp_atlas_bookmark_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBookmarkOutput); i {
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
			RawDescriptor: file_eolymp_atlas_bookmark_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_eolymp_atlas_bookmark_service_proto_goTypes,
		DependencyIndexes: file_eolymp_atlas_bookmark_service_proto_depIdxs,
		MessageInfos:      file_eolymp_atlas_bookmark_service_proto_msgTypes,
	}.Build()
	File_eolymp_atlas_bookmark_service_proto = out.File
	file_eolymp_atlas_bookmark_service_proto_rawDesc = nil
	file_eolymp_atlas_bookmark_service_proto_goTypes = nil
	file_eolymp_atlas_bookmark_service_proto_depIdxs = nil
}
