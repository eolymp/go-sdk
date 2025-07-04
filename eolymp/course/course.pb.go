// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.29.4
// source: eolymp/course/course.proto

package course

import (
	ecm "github.com/eolymp/go-sdk/eolymp/ecm"
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
	state       protoimpl.MessageState `protogen:"open.v1"`
	Id          string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Url         string                 `protobuf:"bytes,2,opt,name=url,proto3" json:"url,omitempty"`
	Locale      string                 `protobuf:"bytes,10,opt,name=locale,proto3" json:"locale,omitempty"`
	Name        string                 `protobuf:"bytes,11,opt,name=name,proto3" json:"name,omitempty"`
	Description *ecm.Content           `protobuf:"bytes,12,opt,name=description,proto3" json:"description,omitempty"`
	ImageUrl    string                 `protobuf:"bytes,13,opt,name=image_url,json=imageUrl,proto3" json:"image_url,omitempty"`
	Visibility  Course_Visibility      `protobuf:"varint,14,opt,name=visibility,proto3,enum=eolymp.course.Course_Visibility" json:"visibility,omitempty"`
	Duration    uint32                 `protobuf:"varint,15,opt,name=duration,proto3" json:"duration,omitempty"`
	// Problem topics (ID of topics from eolymp.taxonomy.TopicService)
	Topics        []string `protobuf:"bytes,16,rep,name=topics,proto3" json:"topics,omitempty"`
	Estimate      uint32   `protobuf:"varint,20,opt,name=estimate,proto3" json:"estimate,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Course) Reset() {
	*x = Course{}
	mi := &file_eolymp_course_course_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Course) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Course) ProtoMessage() {}

func (x *Course) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_course_course_proto_msgTypes[0]
	if x != nil {
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

func (x *Course) GetImageUrl() string {
	if x != nil {
		return x.ImageUrl
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

const file_eolymp_course_course_proto_rawDesc = "" +
	"\n" +
	"\x1aeolymp/course/course.proto\x12\reolymp.course\x1a\x18eolymp/ecm/content.proto\"\xd4\x03\n" +
	"\x06Course\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x12\x10\n" +
	"\x03url\x18\x02 \x01(\tR\x03url\x12\x16\n" +
	"\x06locale\x18\n" +
	" \x01(\tR\x06locale\x12\x12\n" +
	"\x04name\x18\v \x01(\tR\x04name\x125\n" +
	"\vdescription\x18\f \x01(\v2\x13.eolymp.ecm.ContentR\vdescription\x12\x1b\n" +
	"\timage_url\x18\r \x01(\tR\bimageUrl\x12@\n" +
	"\n" +
	"visibility\x18\x0e \x01(\x0e2 .eolymp.course.Course.VisibilityR\n" +
	"visibility\x12\x1a\n" +
	"\bduration\x18\x0f \x01(\rR\bduration\x12\x16\n" +
	"\x06topics\x18\x10 \x03(\tR\x06topics\x12\x1a\n" +
	"\bestimate\x18\x14 \x01(\rR\bestimate\"I\n" +
	"\x05Extra\x12\x11\n" +
	"\rUNKNOWN_EXTRA\x10\x00\x12\x15\n" +
	"\x11DESCRIPTION_VALUE\x10\x01\x12\x16\n" +
	"\x12DESCRIPTION_RENDER\x10\x02\"K\n" +
	"\n" +
	"Visibility\x12\x16\n" +
	"\x12UNKNOWN_VISIBILITY\x10\x00\x12\n" +
	"\n" +
	"\x06PUBLIC\x10\x01\x12\f\n" +
	"\bUNLISTED\x10\x02\x12\v\n" +
	"\aPRIVATE\x10\x03B/Z-github.com/eolymp/go-sdk/eolymp/course;courseb\x06proto3"

var (
	file_eolymp_course_course_proto_rawDescOnce sync.Once
	file_eolymp_course_course_proto_rawDescData []byte
)

func file_eolymp_course_course_proto_rawDescGZIP() []byte {
	file_eolymp_course_course_proto_rawDescOnce.Do(func() {
		file_eolymp_course_course_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_eolymp_course_course_proto_rawDesc), len(file_eolymp_course_course_proto_rawDesc)))
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_eolymp_course_course_proto_rawDesc), len(file_eolymp_course_course_proto_rawDesc)),
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
	file_eolymp_course_course_proto_goTypes = nil
	file_eolymp_course_course_proto_depIdxs = nil
}
