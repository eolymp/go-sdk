// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v3.21.12
// source: eolymp/course/events.proto

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

type ModuleChangedEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CourseId string  `protobuf:"bytes,1,opt,name=course_id,json=courseId,proto3" json:"course_id,omitempty"`
	Before   *Module `protobuf:"bytes,2,opt,name=before,proto3" json:"before,omitempty"`
	After    *Module `protobuf:"bytes,3,opt,name=after,proto3" json:"after,omitempty"`
}

func (x *ModuleChangedEvent) Reset() {
	*x = ModuleChangedEvent{}
	mi := &file_eolymp_course_events_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ModuleChangedEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ModuleChangedEvent) ProtoMessage() {}

func (x *ModuleChangedEvent) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_course_events_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ModuleChangedEvent.ProtoReflect.Descriptor instead.
func (*ModuleChangedEvent) Descriptor() ([]byte, []int) {
	return file_eolymp_course_events_proto_rawDescGZIP(), []int{0}
}

func (x *ModuleChangedEvent) GetCourseId() string {
	if x != nil {
		return x.CourseId
	}
	return ""
}

func (x *ModuleChangedEvent) GetBefore() *Module {
	if x != nil {
		return x.Before
	}
	return nil
}

func (x *ModuleChangedEvent) GetAfter() *Module {
	if x != nil {
		return x.After
	}
	return nil
}

type MaterialChangedEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CourseId string    `protobuf:"bytes,1,opt,name=course_id,json=courseId,proto3" json:"course_id,omitempty"`
	Before   *Material `protobuf:"bytes,3,opt,name=before,proto3" json:"before,omitempty"`
	After    *Material `protobuf:"bytes,4,opt,name=after,proto3" json:"after,omitempty"`
}

func (x *MaterialChangedEvent) Reset() {
	*x = MaterialChangedEvent{}
	mi := &file_eolymp_course_events_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MaterialChangedEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MaterialChangedEvent) ProtoMessage() {}

func (x *MaterialChangedEvent) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_course_events_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MaterialChangedEvent.ProtoReflect.Descriptor instead.
func (*MaterialChangedEvent) Descriptor() ([]byte, []int) {
	return file_eolymp_course_events_proto_rawDescGZIP(), []int{1}
}

func (x *MaterialChangedEvent) GetCourseId() string {
	if x != nil {
		return x.CourseId
	}
	return ""
}

func (x *MaterialChangedEvent) GetBefore() *Material {
	if x != nil {
		return x.Before
	}
	return nil
}

func (x *MaterialChangedEvent) GetAfter() *Material {
	if x != nil {
		return x.After
	}
	return nil
}

type StudentChangedEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CourseId string   `protobuf:"bytes,1,opt,name=course_id,json=courseId,proto3" json:"course_id,omitempty"`
	Before   *Student `protobuf:"bytes,2,opt,name=before,proto3" json:"before,omitempty"`
	After    *Student `protobuf:"bytes,3,opt,name=after,proto3" json:"after,omitempty"`
}

func (x *StudentChangedEvent) Reset() {
	*x = StudentChangedEvent{}
	mi := &file_eolymp_course_events_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *StudentChangedEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StudentChangedEvent) ProtoMessage() {}

func (x *StudentChangedEvent) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_course_events_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StudentChangedEvent.ProtoReflect.Descriptor instead.
func (*StudentChangedEvent) Descriptor() ([]byte, []int) {
	return file_eolymp_course_events_proto_rawDescGZIP(), []int{2}
}

func (x *StudentChangedEvent) GetCourseId() string {
	if x != nil {
		return x.CourseId
	}
	return ""
}

func (x *StudentChangedEvent) GetBefore() *Student {
	if x != nil {
		return x.Before
	}
	return nil
}

func (x *StudentChangedEvent) GetAfter() *Student {
	if x != nil {
		return x.After
	}
	return nil
}

type AssignmentChangedEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CourseId string      `protobuf:"bytes,1,opt,name=course_id,json=courseId,proto3" json:"course_id,omitempty"`
	ModuleId string      `protobuf:"bytes,2,opt,name=module_id,json=moduleId,proto3" json:"module_id,omitempty"`
	Before   *Assignment `protobuf:"bytes,3,opt,name=before,proto3" json:"before,omitempty"`
	After    *Assignment `protobuf:"bytes,4,opt,name=after,proto3" json:"after,omitempty"`
}

func (x *AssignmentChangedEvent) Reset() {
	*x = AssignmentChangedEvent{}
	mi := &file_eolymp_course_events_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AssignmentChangedEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AssignmentChangedEvent) ProtoMessage() {}

func (x *AssignmentChangedEvent) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_course_events_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AssignmentChangedEvent.ProtoReflect.Descriptor instead.
func (*AssignmentChangedEvent) Descriptor() ([]byte, []int) {
	return file_eolymp_course_events_proto_rawDescGZIP(), []int{3}
}

func (x *AssignmentChangedEvent) GetCourseId() string {
	if x != nil {
		return x.CourseId
	}
	return ""
}

func (x *AssignmentChangedEvent) GetModuleId() string {
	if x != nil {
		return x.ModuleId
	}
	return ""
}

func (x *AssignmentChangedEvent) GetBefore() *Assignment {
	if x != nil {
		return x.Before
	}
	return nil
}

func (x *AssignmentChangedEvent) GetAfter() *Assignment {
	if x != nil {
		return x.After
	}
	return nil
}

var File_eolymp_course_events_proto protoreflect.FileDescriptor

var file_eolymp_course_events_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x2f,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x65, 0x6f,
	0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x1a, 0x1e, 0x65, 0x6f, 0x6c,
	0x79, 0x6d, 0x70, 0x2f, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x2f, 0x61, 0x73, 0x73, 0x69, 0x67,
	0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x65, 0x6f, 0x6c,
	0x79, 0x6d, 0x70, 0x2f, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x2f, 0x6d, 0x61, 0x74, 0x65, 0x72,
	0x69, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1a, 0x65, 0x6f, 0x6c, 0x79, 0x6d,
	0x70, 0x2f, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x2f, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x63, 0x6f,
	0x75, 0x72, 0x73, 0x65, 0x2f, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x8d, 0x01, 0x0a, 0x12, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x43, 0x68, 0x61,
	0x6e, 0x67, 0x65, 0x64, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6f, 0x75,
	0x72, 0x73, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6f,
	0x75, 0x72, 0x73, 0x65, 0x49, 0x64, 0x12, 0x2d, 0x0a, 0x06, 0x62, 0x65, 0x66, 0x6f, 0x72, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e,
	0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x2e, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x52, 0x06, 0x62,
	0x65, 0x66, 0x6f, 0x72, 0x65, 0x12, 0x2b, 0x0a, 0x05, 0x61, 0x66, 0x74, 0x65, 0x72, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x63, 0x6f,
	0x75, 0x72, 0x73, 0x65, 0x2e, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x52, 0x05, 0x61, 0x66, 0x74,
	0x65, 0x72, 0x22, 0x93, 0x01, 0x0a, 0x14, 0x4d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x43,
	0x68, 0x61, 0x6e, 0x67, 0x65, 0x64, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x63,
	0x6f, 0x75, 0x72, 0x73, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x49, 0x64, 0x12, 0x2f, 0x0a, 0x06, 0x62, 0x65, 0x66, 0x6f,
	0x72, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d,
	0x70, 0x2e, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x2e, 0x4d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61,
	0x6c, 0x52, 0x06, 0x62, 0x65, 0x66, 0x6f, 0x72, 0x65, 0x12, 0x2d, 0x0a, 0x05, 0x61, 0x66, 0x74,
	0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d,
	0x70, 0x2e, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x2e, 0x4d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61,
	0x6c, 0x52, 0x05, 0x61, 0x66, 0x74, 0x65, 0x72, 0x22, 0x90, 0x01, 0x0a, 0x13, 0x53, 0x74, 0x75,
	0x64, 0x65, 0x6e, 0x74, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x64, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x49, 0x64, 0x12, 0x2e, 0x0a,
	0x06, 0x62, 0x65, 0x66, 0x6f, 0x72, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e,
	0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x2e, 0x53, 0x74,
	0x75, 0x64, 0x65, 0x6e, 0x74, 0x52, 0x06, 0x62, 0x65, 0x66, 0x6f, 0x72, 0x65, 0x12, 0x2c, 0x0a,
	0x05, 0x61, 0x66, 0x74, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x65,
	0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x2e, 0x53, 0x74, 0x75,
	0x64, 0x65, 0x6e, 0x74, 0x52, 0x05, 0x61, 0x66, 0x74, 0x65, 0x72, 0x22, 0xb6, 0x01, 0x0a, 0x16,
	0x41, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65,
	0x64, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6f, 0x75, 0x72, 0x73,
	0x65, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x5f, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x49, 0x64,
	0x12, 0x31, 0x0a, 0x06, 0x62, 0x65, 0x66, 0x6f, 0x72, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x19, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65,
	0x2e, 0x41, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x06, 0x62, 0x65, 0x66,
	0x6f, 0x72, 0x65, 0x12, 0x2f, 0x0a, 0x05, 0x61, 0x66, 0x74, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x19, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x63, 0x6f, 0x75, 0x72,
	0x73, 0x65, 0x2e, 0x41, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x05, 0x61,
	0x66, 0x74, 0x65, 0x72, 0x42, 0x2f, 0x5a, 0x2d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x67, 0x6f, 0x2d, 0x73, 0x64, 0x6b,
	0x2f, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x3b, 0x63,
	0x6f, 0x75, 0x72, 0x73, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_eolymp_course_events_proto_rawDescOnce sync.Once
	file_eolymp_course_events_proto_rawDescData = file_eolymp_course_events_proto_rawDesc
)

func file_eolymp_course_events_proto_rawDescGZIP() []byte {
	file_eolymp_course_events_proto_rawDescOnce.Do(func() {
		file_eolymp_course_events_proto_rawDescData = protoimpl.X.CompressGZIP(file_eolymp_course_events_proto_rawDescData)
	})
	return file_eolymp_course_events_proto_rawDescData
}

var file_eolymp_course_events_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_eolymp_course_events_proto_goTypes = []any{
	(*ModuleChangedEvent)(nil),     // 0: eolymp.course.ModuleChangedEvent
	(*MaterialChangedEvent)(nil),   // 1: eolymp.course.MaterialChangedEvent
	(*StudentChangedEvent)(nil),    // 2: eolymp.course.StudentChangedEvent
	(*AssignmentChangedEvent)(nil), // 3: eolymp.course.AssignmentChangedEvent
	(*Module)(nil),                 // 4: eolymp.course.Module
	(*Material)(nil),               // 5: eolymp.course.Material
	(*Student)(nil),                // 6: eolymp.course.Student
	(*Assignment)(nil),             // 7: eolymp.course.Assignment
}
var file_eolymp_course_events_proto_depIdxs = []int32{
	4, // 0: eolymp.course.ModuleChangedEvent.before:type_name -> eolymp.course.Module
	4, // 1: eolymp.course.ModuleChangedEvent.after:type_name -> eolymp.course.Module
	5, // 2: eolymp.course.MaterialChangedEvent.before:type_name -> eolymp.course.Material
	5, // 3: eolymp.course.MaterialChangedEvent.after:type_name -> eolymp.course.Material
	6, // 4: eolymp.course.StudentChangedEvent.before:type_name -> eolymp.course.Student
	6, // 5: eolymp.course.StudentChangedEvent.after:type_name -> eolymp.course.Student
	7, // 6: eolymp.course.AssignmentChangedEvent.before:type_name -> eolymp.course.Assignment
	7, // 7: eolymp.course.AssignmentChangedEvent.after:type_name -> eolymp.course.Assignment
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_eolymp_course_events_proto_init() }
func file_eolymp_course_events_proto_init() {
	if File_eolymp_course_events_proto != nil {
		return
	}
	file_eolymp_course_assignment_proto_init()
	file_eolymp_course_material_proto_init()
	file_eolymp_course_module_proto_init()
	file_eolymp_course_student_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_eolymp_course_events_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_eolymp_course_events_proto_goTypes,
		DependencyIndexes: file_eolymp_course_events_proto_depIdxs,
		MessageInfos:      file_eolymp_course_events_proto_msgTypes,
	}.Build()
	File_eolymp_course_events_proto = out.File
	file_eolymp_course_events_proto_rawDesc = nil
	file_eolymp_course_events_proto_goTypes = nil
	file_eolymp_course_events_proto_depIdxs = nil
}
