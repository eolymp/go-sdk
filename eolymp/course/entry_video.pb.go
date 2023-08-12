// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.18.1
// source: eolymp/course/entry_video.proto

package course

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

type Video struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ImageUrl string `protobuf:"bytes,1,opt,name=image_url,json=imageUrl,proto3" json:"image_url,omitempty"`
	VideoUrl string `protobuf:"bytes,2,opt,name=video_url,json=videoUrl,proto3" json:"video_url,omitempty"`
	Duration uint32 `protobuf:"varint,3,opt,name=duration,proto3" json:"duration,omitempty"`
}

func (x *Video) Reset() {
	*x = Video{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_course_entry_video_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Video) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Video) ProtoMessage() {}

func (x *Video) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_course_entry_video_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Video.ProtoReflect.Descriptor instead.
func (*Video) Descriptor() ([]byte, []int) {
	return file_eolymp_course_entry_video_proto_rawDescGZIP(), []int{0}
}

func (x *Video) GetImageUrl() string {
	if x != nil {
		return x.ImageUrl
	}
	return ""
}

func (x *Video) GetVideoUrl() string {
	if x != nil {
		return x.VideoUrl
	}
	return ""
}

func (x *Video) GetDuration() uint32 {
	if x != nil {
		return x.Duration
	}
	return 0
}

var File_eolymp_course_entry_video_proto protoreflect.FileDescriptor

var file_eolymp_course_entry_video_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x2f,
	0x65, 0x6e, 0x74, 0x72, 0x79, 0x5f, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0d, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65,
	0x22, 0x5d, 0x0a, 0x05, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x6d, 0x61,
	0x67, 0x65, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x69, 0x6d,
	0x61, 0x67, 0x65, 0x55, 0x72, 0x6c, 0x12, 0x1b, 0x0a, 0x09, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x5f,
	0x75, 0x72, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x76, 0x69, 0x64, 0x65, 0x6f,
	0x55, 0x72, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42,
	0x2f, 0x5a, 0x2d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6f,
	0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x67, 0x6f, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x65, 0x6f, 0x6c, 0x79,
	0x6d, 0x70, 0x2f, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x3b, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_eolymp_course_entry_video_proto_rawDescOnce sync.Once
	file_eolymp_course_entry_video_proto_rawDescData = file_eolymp_course_entry_video_proto_rawDesc
)

func file_eolymp_course_entry_video_proto_rawDescGZIP() []byte {
	file_eolymp_course_entry_video_proto_rawDescOnce.Do(func() {
		file_eolymp_course_entry_video_proto_rawDescData = protoimpl.X.CompressGZIP(file_eolymp_course_entry_video_proto_rawDescData)
	})
	return file_eolymp_course_entry_video_proto_rawDescData
}

var file_eolymp_course_entry_video_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_eolymp_course_entry_video_proto_goTypes = []interface{}{
	(*Video)(nil), // 0: eolymp.course.Video
}
var file_eolymp_course_entry_video_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_eolymp_course_entry_video_proto_init() }
func file_eolymp_course_entry_video_proto_init() {
	if File_eolymp_course_entry_video_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_eolymp_course_entry_video_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Video); i {
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
			RawDescriptor: file_eolymp_course_entry_video_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_eolymp_course_entry_video_proto_goTypes,
		DependencyIndexes: file_eolymp_course_entry_video_proto_depIdxs,
		MessageInfos:      file_eolymp_course_entry_video_proto_msgTypes,
	}.Build()
	File_eolymp_course_entry_video_proto = out.File
	file_eolymp_course_entry_video_proto_rawDesc = nil
	file_eolymp_course_entry_video_proto_goTypes = nil
	file_eolymp_course_entry_video_proto_depIdxs = nil
}