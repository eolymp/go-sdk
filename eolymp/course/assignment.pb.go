// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v3.18.1
// source: eolymp/course/assignment.proto

package course

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

type Assignment_Status int32

const (
	Assignment_UNKNOWN_STATUS Assignment_Status = 0
	Assignment_INHERITED      Assignment_Status = 1 // same as parent element (same as PENDING on course level)
	Assignment_INACTIVE       Assignment_Status = 2 // assignment is not assigned to the user
	Assignment_SCHEDULED      Assignment_Status = 3 // assignment will start automatically at start_at
	Assignment_READY          Assignment_Status = 4 // assignment is ready to be started, student must call StartAssignment to activate assignment
	Assignment_ACTIVE         Assignment_Status = 5 // assignment is active and can be seen by the student
	Assignment_COMPLETE       Assignment_Status = 6 // assignment time has run out and not shown to the student anymore
	Assignment_UPSOLVE        Assignment_Status = 7 // assignment time has run out, the result is final, but student can see course and solve tasks
)

// Enum value maps for Assignment_Status.
var (
	Assignment_Status_name = map[int32]string{
		0: "UNKNOWN_STATUS",
		1: "INHERITED",
		2: "INACTIVE",
		3: "SCHEDULED",
		4: "READY",
		5: "ACTIVE",
		6: "COMPLETE",
		7: "UPSOLVE",
	}
	Assignment_Status_value = map[string]int32{
		"UNKNOWN_STATUS": 0,
		"INHERITED":      1,
		"INACTIVE":       2,
		"SCHEDULED":      3,
		"READY":          4,
		"ACTIVE":         5,
		"COMPLETE":       6,
		"UPSOLVE":        7,
	}
)

func (x Assignment_Status) Enum() *Assignment_Status {
	p := new(Assignment_Status)
	*p = x
	return p
}

func (x Assignment_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Assignment_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_eolymp_course_assignment_proto_enumTypes[0].Descriptor()
}

func (Assignment_Status) Type() protoreflect.EnumType {
	return &file_eolymp_course_assignment_proto_enumTypes[0]
}

func (x Assignment_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Assignment_Status.Descriptor instead.
func (Assignment_Status) EnumDescriptor() ([]byte, []int) {
	return file_eolymp_course_assignment_proto_rawDescGZIP(), []int{0, 0}
}

type Assignment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status     Assignment_Status      `protobuf:"varint,1,opt,name=status,proto3,enum=eolymp.course.Assignment_Status" json:"status,omitempty"`
	StartAt    *timestamppb.Timestamp `protobuf:"bytes,10,opt,name=start_at,json=startAt,proto3" json:"start_at,omitempty"`          // optionally, time when assignment will commence
	EndAt      *timestamppb.Timestamp `protobuf:"bytes,11,opt,name=end_at,json=endAt,proto3" json:"end_at,omitempty"`                // optionally, time by when assignment should be complete
	CompleteAt *timestamppb.Timestamp `protobuf:"bytes,12,opt,name=complete_at,json=completeAt,proto3" json:"complete_at,omitempty"` // optionally, time when assignment was actually complete
	Duration   uint32                 `protobuf:"varint,15,opt,name=duration,proto3" json:"duration,omitempty"`                      // optionally, duration of the assignment in seconds
	Progress   float32                `protobuf:"fixed32,20,opt,name=progress,proto3" json:"progress,omitempty"`                     // progress from 0 to 1
	Grade      uint32                 `protobuf:"varint,30,opt,name=grade,proto3" json:"grade,omitempty"`                            // grade from 0 to 100
}

func (x *Assignment) Reset() {
	*x = Assignment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_course_assignment_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Assignment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Assignment) ProtoMessage() {}

func (x *Assignment) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_course_assignment_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Assignment.ProtoReflect.Descriptor instead.
func (*Assignment) Descriptor() ([]byte, []int) {
	return file_eolymp_course_assignment_proto_rawDescGZIP(), []int{0}
}

func (x *Assignment) GetStatus() Assignment_Status {
	if x != nil {
		return x.Status
	}
	return Assignment_UNKNOWN_STATUS
}

func (x *Assignment) GetStartAt() *timestamppb.Timestamp {
	if x != nil {
		return x.StartAt
	}
	return nil
}

func (x *Assignment) GetEndAt() *timestamppb.Timestamp {
	if x != nil {
		return x.EndAt
	}
	return nil
}

func (x *Assignment) GetCompleteAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CompleteAt
	}
	return nil
}

func (x *Assignment) GetDuration() uint32 {
	if x != nil {
		return x.Duration
	}
	return 0
}

func (x *Assignment) GetProgress() float32 {
	if x != nil {
		return x.Progress
	}
	return 0
}

func (x *Assignment) GetGrade() uint32 {
	if x != nil {
		return x.Grade
	}
	return 0
}

var File_eolymp_course_assignment_proto protoreflect.FileDescriptor

var file_eolymp_course_assignment_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x2f,
	0x61, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0d, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x1a,
	0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xb7, 0x03, 0x0a, 0x0a, 0x41, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x12,
	0x38, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x20, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x2e,
	0x41, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x35, 0x0a, 0x08, 0x73, 0x74, 0x61,
	0x72, 0x74, 0x5f, 0x61, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x07, 0x73, 0x74, 0x61, 0x72, 0x74, 0x41, 0x74,
	0x12, 0x31, 0x0a, 0x06, 0x65, 0x6e, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x05, 0x65, 0x6e,
	0x64, 0x41, 0x74, 0x12, 0x3b, 0x0a, 0x0b, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x5f,
	0x61, 0x74, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x74,
	0x12, 0x1a, 0x0a, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x0f, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08,
	0x70, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x18, 0x14, 0x20, 0x01, 0x28, 0x02, 0x52, 0x08,
	0x70, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x67, 0x72, 0x61, 0x64,
	0x65, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x67, 0x72, 0x61, 0x64, 0x65, 0x22, 0x7a,
	0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x12, 0x0a, 0x0e, 0x55, 0x4e, 0x4b, 0x4e,
	0x4f, 0x57, 0x4e, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x10, 0x00, 0x12, 0x0d, 0x0a, 0x09,
	0x49, 0x4e, 0x48, 0x45, 0x52, 0x49, 0x54, 0x45, 0x44, 0x10, 0x01, 0x12, 0x0c, 0x0a, 0x08, 0x49,
	0x4e, 0x41, 0x43, 0x54, 0x49, 0x56, 0x45, 0x10, 0x02, 0x12, 0x0d, 0x0a, 0x09, 0x53, 0x43, 0x48,
	0x45, 0x44, 0x55, 0x4c, 0x45, 0x44, 0x10, 0x03, 0x12, 0x09, 0x0a, 0x05, 0x52, 0x45, 0x41, 0x44,
	0x59, 0x10, 0x04, 0x12, 0x0a, 0x0a, 0x06, 0x41, 0x43, 0x54, 0x49, 0x56, 0x45, 0x10, 0x05, 0x12,
	0x0c, 0x0a, 0x08, 0x43, 0x4f, 0x4d, 0x50, 0x4c, 0x45, 0x54, 0x45, 0x10, 0x06, 0x12, 0x0b, 0x0a,
	0x07, 0x55, 0x50, 0x53, 0x4f, 0x4c, 0x56, 0x45, 0x10, 0x07, 0x42, 0x2f, 0x5a, 0x2d, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f,
	0x67, 0x6f, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x63, 0x6f,
	0x75, 0x72, 0x73, 0x65, 0x3b, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_eolymp_course_assignment_proto_rawDescOnce sync.Once
	file_eolymp_course_assignment_proto_rawDescData = file_eolymp_course_assignment_proto_rawDesc
)

func file_eolymp_course_assignment_proto_rawDescGZIP() []byte {
	file_eolymp_course_assignment_proto_rawDescOnce.Do(func() {
		file_eolymp_course_assignment_proto_rawDescData = protoimpl.X.CompressGZIP(file_eolymp_course_assignment_proto_rawDescData)
	})
	return file_eolymp_course_assignment_proto_rawDescData
}

var file_eolymp_course_assignment_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_eolymp_course_assignment_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_eolymp_course_assignment_proto_goTypes = []interface{}{
	(Assignment_Status)(0),        // 0: eolymp.course.Assignment.Status
	(*Assignment)(nil),            // 1: eolymp.course.Assignment
	(*timestamppb.Timestamp)(nil), // 2: google.protobuf.Timestamp
}
var file_eolymp_course_assignment_proto_depIdxs = []int32{
	0, // 0: eolymp.course.Assignment.status:type_name -> eolymp.course.Assignment.Status
	2, // 1: eolymp.course.Assignment.start_at:type_name -> google.protobuf.Timestamp
	2, // 2: eolymp.course.Assignment.end_at:type_name -> google.protobuf.Timestamp
	2, // 3: eolymp.course.Assignment.complete_at:type_name -> google.protobuf.Timestamp
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_eolymp_course_assignment_proto_init() }
func file_eolymp_course_assignment_proto_init() {
	if File_eolymp_course_assignment_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_eolymp_course_assignment_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Assignment); i {
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
			RawDescriptor: file_eolymp_course_assignment_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_eolymp_course_assignment_proto_goTypes,
		DependencyIndexes: file_eolymp_course_assignment_proto_depIdxs,
		EnumInfos:         file_eolymp_course_assignment_proto_enumTypes,
		MessageInfos:      file_eolymp_course_assignment_proto_msgTypes,
	}.Build()
	File_eolymp_course_assignment_proto = out.File
	file_eolymp_course_assignment_proto_rawDesc = nil
	file_eolymp_course_assignment_proto_goTypes = nil
	file_eolymp_course_assignment_proto_depIdxs = nil
}
