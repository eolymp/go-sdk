// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.18.1
// source: eolymp/course/entry.proto

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

type Entry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Title    string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	ParentId string `protobuf:"bytes,10,opt,name=parent_id,json=parentId,proto3" json:"parent_id,omitempty"`
	Index    uint32 `protobuf:"varint,11,opt,name=index,proto3" json:"index,omitempty"`
	Estimate uint32 `protobuf:"varint,21,opt,name=estimate,proto3" json:"estimate,omitempty"`
	// Types that are assignable to Content:
	//	*Entry_Section
	//	*Entry_Document
	//	*Entry_Video
	Content isEntry_Content `protobuf_oneof:"content"`
}

func (x *Entry) Reset() {
	*x = Entry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_course_entry_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Entry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Entry) ProtoMessage() {}

func (x *Entry) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_course_entry_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Entry.ProtoReflect.Descriptor instead.
func (*Entry) Descriptor() ([]byte, []int) {
	return file_eolymp_course_entry_proto_rawDescGZIP(), []int{0}
}

func (x *Entry) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Entry) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Entry) GetParentId() string {
	if x != nil {
		return x.ParentId
	}
	return ""
}

func (x *Entry) GetIndex() uint32 {
	if x != nil {
		return x.Index
	}
	return 0
}

func (x *Entry) GetEstimate() uint32 {
	if x != nil {
		return x.Estimate
	}
	return 0
}

func (m *Entry) GetContent() isEntry_Content {
	if m != nil {
		return m.Content
	}
	return nil
}

func (x *Entry) GetSection() *Section {
	if x, ok := x.GetContent().(*Entry_Section); ok {
		return x.Section
	}
	return nil
}

func (x *Entry) GetDocument() *ecm.Content {
	if x, ok := x.GetContent().(*Entry_Document); ok {
		return x.Document
	}
	return nil
}

func (x *Entry) GetVideo() *Video {
	if x, ok := x.GetContent().(*Entry_Video); ok {
		return x.Video
	}
	return nil
}

type isEntry_Content interface {
	isEntry_Content()
}

type Entry_Section struct {
	Section *Section `protobuf:"bytes,100,opt,name=section,proto3,oneof"`
}

type Entry_Document struct {
	Document *ecm.Content `protobuf:"bytes,101,opt,name=document,proto3,oneof"`
}

type Entry_Video struct {
	Video *Video `protobuf:"bytes,102,opt,name=video,proto3,oneof"`
}

func (*Entry_Section) isEntry_Content() {}

func (*Entry_Document) isEntry_Content() {}

func (*Entry_Video) isEntry_Content() {}

var File_eolymp_course_entry_proto protoreflect.FileDescriptor

var file_eolymp_course_entry_proto_rawDesc = []byte{
	0x0a, 0x19, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x2f,
	0x65, 0x6e, 0x74, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x65, 0x6f, 0x6c,
	0x79, 0x6d, 0x70, 0x2e, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x1a, 0x21, 0x65, 0x6f, 0x6c, 0x79,
	0x6d, 0x70, 0x2f, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x2f, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x5f,
	0x73, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x65,
	0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x2f, 0x65, 0x6e, 0x74,
	0x72, 0x79, 0x5f, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x18,
	0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x65, 0x63, 0x6d, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9c, 0x02, 0x0a, 0x05, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x72, 0x65,
	0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x72,
	0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x0b,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x1a, 0x0a, 0x08, 0x65,
	0x73, 0x74, 0x69, 0x6d, 0x61, 0x74, 0x65, 0x18, 0x15, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x65,
	0x73, 0x74, 0x69, 0x6d, 0x61, 0x74, 0x65, 0x12, 0x32, 0x0a, 0x07, 0x73, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x64, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d,
	0x70, 0x2e, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x2e, 0x53, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x48, 0x00, 0x52, 0x07, 0x73, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x31, 0x0a, 0x08, 0x64,
	0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x65, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e,
	0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x65, 0x63, 0x6d, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x48, 0x00, 0x52, 0x08, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x2c,
	0x0a, 0x05, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x18, 0x66, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e,
	0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x2e, 0x56, 0x69,
	0x64, 0x65, 0x6f, 0x48, 0x00, 0x52, 0x05, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x42, 0x09, 0x0a, 0x07,
	0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x42, 0x2f, 0x5a, 0x2d, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x67, 0x6f, 0x2d,
	0x73, 0x64, 0x6b, 0x2f, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x63, 0x6f, 0x75, 0x72, 0x73,
	0x65, 0x3b, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_eolymp_course_entry_proto_rawDescOnce sync.Once
	file_eolymp_course_entry_proto_rawDescData = file_eolymp_course_entry_proto_rawDesc
)

func file_eolymp_course_entry_proto_rawDescGZIP() []byte {
	file_eolymp_course_entry_proto_rawDescOnce.Do(func() {
		file_eolymp_course_entry_proto_rawDescData = protoimpl.X.CompressGZIP(file_eolymp_course_entry_proto_rawDescData)
	})
	return file_eolymp_course_entry_proto_rawDescData
}

var file_eolymp_course_entry_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_eolymp_course_entry_proto_goTypes = []interface{}{
	(*Entry)(nil),       // 0: eolymp.course.Entry
	(*Section)(nil),     // 1: eolymp.course.Section
	(*ecm.Content)(nil), // 2: eolymp.ecm.Content
	(*Video)(nil),       // 3: eolymp.course.Video
}
var file_eolymp_course_entry_proto_depIdxs = []int32{
	1, // 0: eolymp.course.Entry.section:type_name -> eolymp.course.Section
	2, // 1: eolymp.course.Entry.document:type_name -> eolymp.ecm.Content
	3, // 2: eolymp.course.Entry.video:type_name -> eolymp.course.Video
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_eolymp_course_entry_proto_init() }
func file_eolymp_course_entry_proto_init() {
	if File_eolymp_course_entry_proto != nil {
		return
	}
	file_eolymp_course_entry_section_proto_init()
	file_eolymp_course_entry_video_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_eolymp_course_entry_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Entry); i {
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
	file_eolymp_course_entry_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*Entry_Section)(nil),
		(*Entry_Document)(nil),
		(*Entry_Video)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_eolymp_course_entry_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_eolymp_course_entry_proto_goTypes,
		DependencyIndexes: file_eolymp_course_entry_proto_depIdxs,
		MessageInfos:      file_eolymp_course_entry_proto_msgTypes,
	}.Build()
	File_eolymp_course_entry_proto = out.File
	file_eolymp_course_entry_proto_rawDesc = nil
	file_eolymp_course_entry_proto_goTypes = nil
	file_eolymp_course_entry_proto_depIdxs = nil
}
