// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.21.12
// source: eolymp/course/course.proto

package course

import (
	ecm "github.com/eolymp/go-sdk/eolymp/ecm"
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

type Course_Extra int32

const (
	Course_UNKNOWN_EXTRA      Course_Extra = 0
	Course_DESCRIPTION_VALUE  Course_Extra = 1
	Course_DESCRIPTION_RENDER Course_Extra = 2
)

// Enum value maps for Course_Extra.
var (
	Course_Extra_name = map[int32]string{
		0: "UNKNOWN_EXTRA",
		1: "DESCRIPTION_VALUE",
		2: "DESCRIPTION_RENDER",
	}
	Course_Extra_value = map[string]int32{
		"UNKNOWN_EXTRA":      0,
		"DESCRIPTION_VALUE":  1,
		"DESCRIPTION_RENDER": 2,
	}
)

func (x Course_Extra) Enum() *Course_Extra {
	p := new(Course_Extra)
	*p = x
	return p
}

func (x Course_Extra) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Course_Extra) Descriptor() protoreflect.EnumDescriptor {
	return file_eolymp_course_course_proto_enumTypes[0].Descriptor()
}

func (Course_Extra) Type() protoreflect.EnumType {
	return &file_eolymp_course_course_proto_enumTypes[0]
}

func (x Course_Extra) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Course_Extra.Descriptor instead.
func (Course_Extra) EnumDescriptor() ([]byte, []int) {
	return file_eolymp_course_course_proto_rawDescGZIP(), []int{0, 0}
}

type Course_Visibility int32

const (
	Course_UNKNOWN_VISIBILITY Course_Visibility = 0
	Course_PUBLIC             Course_Visibility = 1 // anyone can join, displayed in the list
	Course_UNLISTED           Course_Visibility = 2 // anyone can join, not displayed in the list
	Course_PRIVATE            Course_Visibility = 3 // course must be explicitly assigned
)

// Enum value maps for Course_Visibility.
var (
	Course_Visibility_name = map[int32]string{
		0: "UNKNOWN_VISIBILITY",
		1: "PUBLIC",
		2: "UNLISTED",
		3: "PRIVATE",
	}
	Course_Visibility_value = map[string]int32{
		"UNKNOWN_VISIBILITY": 0,
		"PUBLIC":             1,
		"UNLISTED":           2,
		"PRIVATE":            3,
	}
)

func (x Course_Visibility) Enum() *Course_Visibility {
	p := new(Course_Visibility)
	*p = x
	return p
}

func (x Course_Visibility) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Course_Visibility) Descriptor() protoreflect.EnumDescriptor {
	return file_eolymp_course_course_proto_enumTypes[1].Descriptor()
}

func (Course_Visibility) Type() protoreflect.EnumType {
	return &file_eolymp_course_course_proto_enumTypes[1]
}

func (x Course_Visibility) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Course_Visibility.Descriptor instead.
func (Course_Visibility) EnumDescriptor() ([]byte, []int) {
	return file_eolymp_course_course_proto_rawDescGZIP(), []int{0, 1}
}

type Course struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string            `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Url         string            `protobuf:"bytes,2,opt,name=url,proto3" json:"url,omitempty"`
	Locale      string            `protobuf:"bytes,10,opt,name=locale,proto3" json:"locale,omitempty"`
	Name        string            `protobuf:"bytes,11,opt,name=name,proto3" json:"name,omitempty"`
	Description *ecm.Content      `protobuf:"bytes,12,opt,name=description,proto3" json:"description,omitempty"`
	Image       string            `protobuf:"bytes,13,opt,name=image,proto3" json:"image,omitempty"`
	Visibility  Course_Visibility `protobuf:"varint,14,opt,name=visibility,proto3,enum=eolymp.course.Course_Visibility" json:"visibility,omitempty"`
	Duration    uint32            `protobuf:"varint,15,opt,name=duration,proto3" json:"duration,omitempty"`
	// Problem topics (ID of topics from eolymp.taxonomy.TopicService)
	Topics   []string `protobuf:"bytes,16,rep,name=topics,proto3" json:"topics,omitempty"`
	Estimate uint32   `protobuf:"varint,20,opt,name=estimate,proto3" json:"estimate,omitempty"`
}

func (x *Course) Reset() {
	*x = Course{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_course_course_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Course) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Course) ProtoMessage() {}

func (x *Course) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_course_course_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Course.ProtoReflect.Descriptor instead.
func (*Course) Descriptor() ([]byte, []int) {
	return file_eolymp_course_course_proto_rawDescGZIP(), []int{0}
}

func (x *Course) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Course) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *Course) GetLocale() string {
	if x != nil {
		return x.Locale
	}
	return ""
}

func (x *Course) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Course) GetDescription() *ecm.Content {
	if x != nil {
		return x.Description
	}
	return nil
}

func (x *Course) GetImage() string {
	if x != nil {
		return x.Image
	}
	return ""
}

func (x *Course) GetVisibility() Course_Visibility {
	if x != nil {
		return x.Visibility
	}
	return Course_UNKNOWN_VISIBILITY
}

func (x *Course) GetDuration() uint32 {
	if x != nil {
		return x.Duration
	}
	return 0
}

func (x *Course) GetTopics() []string {
	if x != nil {
		return x.Topics
	}
	return nil
}

func (x *Course) GetEstimate() uint32 {
	if x != nil {
		return x.Estimate
	}
	return 0
}

var File_eolymp_course_course_proto protoreflect.FileDescriptor

var file_eolymp_course_course_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x2f,
	0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x65, 0x6f,
	0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x1a, 0x18, 0x65, 0x6f, 0x6c,
	0x79, 0x6d, 0x70, 0x2f, 0x65, 0x63, 0x6d, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xcd, 0x03, 0x0a, 0x06, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75,
	0x72, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x65, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x35,
	0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x0c, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x65, 0x63, 0x6d,
	0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x0d,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x40, 0x0a, 0x0a, 0x76,
	0x69, 0x73, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x20, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x2e,
	0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x2e, 0x56, 0x69, 0x73, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x74,
	0x79, 0x52, 0x0a, 0x76, 0x69, 0x73, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x12, 0x1a, 0x0a,
	0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x6f, 0x70,
	0x69, 0x63, 0x73, 0x18, 0x10, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x74, 0x6f, 0x70, 0x69, 0x63,
	0x73, 0x12, 0x1a, 0x0a, 0x08, 0x65, 0x73, 0x74, 0x69, 0x6d, 0x61, 0x74, 0x65, 0x18, 0x14, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x08, 0x65, 0x73, 0x74, 0x69, 0x6d, 0x61, 0x74, 0x65, 0x22, 0x49, 0x0a,
	0x05, 0x45, 0x78, 0x74, 0x72, 0x61, 0x12, 0x11, 0x0a, 0x0d, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57,
	0x4e, 0x5f, 0x45, 0x58, 0x54, 0x52, 0x41, 0x10, 0x00, 0x12, 0x15, 0x0a, 0x11, 0x44, 0x45, 0x53,
	0x43, 0x52, 0x49, 0x50, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x56, 0x41, 0x4c, 0x55, 0x45, 0x10, 0x01,
	0x12, 0x16, 0x0a, 0x12, 0x44, 0x45, 0x53, 0x43, 0x52, 0x49, 0x50, 0x54, 0x49, 0x4f, 0x4e, 0x5f,
	0x52, 0x45, 0x4e, 0x44, 0x45, 0x52, 0x10, 0x02, 0x22, 0x4b, 0x0a, 0x0a, 0x56, 0x69, 0x73, 0x69,
	0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x12, 0x16, 0x0a, 0x12, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57,
	0x4e, 0x5f, 0x56, 0x49, 0x53, 0x49, 0x42, 0x49, 0x4c, 0x49, 0x54, 0x59, 0x10, 0x00, 0x12, 0x0a,
	0x0a, 0x06, 0x50, 0x55, 0x42, 0x4c, 0x49, 0x43, 0x10, 0x01, 0x12, 0x0c, 0x0a, 0x08, 0x55, 0x4e,
	0x4c, 0x49, 0x53, 0x54, 0x45, 0x44, 0x10, 0x02, 0x12, 0x0b, 0x0a, 0x07, 0x50, 0x52, 0x49, 0x56,
	0x41, 0x54, 0x45, 0x10, 0x03, 0x42, 0x2f, 0x5a, 0x2d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x67, 0x6f, 0x2d, 0x73, 0x64,
	0x6b, 0x2f, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x3b,
	0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_eolymp_course_course_proto_rawDescOnce sync.Once
	file_eolymp_course_course_proto_rawDescData = file_eolymp_course_course_proto_rawDesc
)

func file_eolymp_course_course_proto_rawDescGZIP() []byte {
	file_eolymp_course_course_proto_rawDescOnce.Do(func() {
		file_eolymp_course_course_proto_rawDescData = protoimpl.X.CompressGZIP(file_eolymp_course_course_proto_rawDescData)
	})
	return file_eolymp_course_course_proto_rawDescData
}

var file_eolymp_course_course_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_eolymp_course_course_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_eolymp_course_course_proto_goTypes = []any{
	(Course_Extra)(0),      // 0: eolymp.course.Course.Extra
	(Course_Visibility)(0), // 1: eolymp.course.Course.Visibility
	(*Course)(nil),         // 2: eolymp.course.Course
	(*ecm.Content)(nil),    // 3: eolymp.ecm.Content
}
var file_eolymp_course_course_proto_depIdxs = []int32{
	3, // 0: eolymp.course.Course.description:type_name -> eolymp.ecm.Content
	1, // 1: eolymp.course.Course.visibility:type_name -> eolymp.course.Course.Visibility
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_eolymp_course_course_proto_init() }
func file_eolymp_course_course_proto_init() {
	if File_eolymp_course_course_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_eolymp_course_course_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*Course); i {
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
			RawDescriptor: file_eolymp_course_course_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_eolymp_course_course_proto_goTypes,
		DependencyIndexes: file_eolymp_course_course_proto_depIdxs,
		EnumInfos:         file_eolymp_course_course_proto_enumTypes,
		MessageInfos:      file_eolymp_course_course_proto_msgTypes,
	}.Build()
	File_eolymp_course_course_proto = out.File
	file_eolymp_course_course_proto_rawDesc = nil
	file_eolymp_course_course_proto_goTypes = nil
	file_eolymp_course_course_proto_depIdxs = nil
}
