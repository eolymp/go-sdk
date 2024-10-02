// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.21.12
// source: eolymp/course/module.proto

package course

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

type Module_Extra int32

const (
	Module_UNKNOWN_EXTRA      Module_Extra = 0
	Module_DESCRIPTION_VALUE  Module_Extra = 1
	Module_DESCRIPTION_RENDER Module_Extra = 2
	Module_PROGRESS           Module_Extra = 4
	Module_ASSIGNMENT         Module_Extra = 5
)

// Enum value maps for Module_Extra.
var (
	Module_Extra_name = map[int32]string{
		0: "UNKNOWN_EXTRA",
		1: "DESCRIPTION_VALUE",
		2: "DESCRIPTION_RENDER",
		4: "PROGRESS",
		5: "ASSIGNMENT",
	}
	Module_Extra_value = map[string]int32{
		"UNKNOWN_EXTRA":      0,
		"DESCRIPTION_VALUE":  1,
		"DESCRIPTION_RENDER": 2,
		"PROGRESS":           4,
		"ASSIGNMENT":         5,
	}
)

func (x Module_Extra) Enum() *Module_Extra {
	p := new(Module_Extra)
	*p = x
	return p
}

func (x Module_Extra) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Module_Extra) Descriptor() protoreflect.EnumDescriptor {
	return file_eolymp_course_module_proto_enumTypes[0].Descriptor()
}

func (Module_Extra) Type() protoreflect.EnumType {
	return &file_eolymp_course_module_proto_enumTypes[0]
}

func (x Module_Extra) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Module_Extra.Descriptor instead.
func (Module_Extra) EnumDescriptor() ([]byte, []int) {
	return file_eolymp_course_module_proto_rawDescGZIP(), []int{0, 0}
}

type Module struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id             string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Url            string                 `protobuf:"bytes,2,opt,name=url,proto3" json:"url,omitempty"`
	Draft          bool                   `protobuf:"varint,3,opt,name=draft,proto3" json:"draft,omitempty"`
	Extra          bool                   `protobuf:"varint,7,opt,name=extra,proto3" json:"extra,omitempty"`    // extra module, will only be shown when explicitly assigned
	Weight         float32                `protobuf:"fixed32,8,opt,name=weight,proto3" json:"weight,omitempty"` // weight of the module when calculating course score
	Name           string                 `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	ImageUrl       string                 `protobuf:"bytes,5,opt,name=image_url,json=imageUrl,proto3" json:"image_url,omitempty"`
	Index          uint32                 `protobuf:"varint,6,opt,name=index,proto3" json:"index,omitempty"`
	Description    *ecm.Content           `protobuf:"bytes,10,opt,name=description,proto3" json:"description,omitempty"`
	StartAfter     *timestamppb.Timestamp `protobuf:"bytes,11,opt,name=start_after,json=startAfter,proto3" json:"start_after,omitempty"`             // optionally, default time by when module should be complete
	CompleteBefore *timestamppb.Timestamp `protobuf:"bytes,12,opt,name=complete_before,json=completeBefore,proto3" json:"complete_before,omitempty"` // optionally, default time by when module should be complete
	Duration       uint32                 `protobuf:"varint,13,opt,name=duration,proto3" json:"duration,omitempty"`                                  // optionally, default duration of the module in seconds
	Progress       *Module_Progress       `protobuf:"bytes,31,opt,name=progress,proto3" json:"progress,omitempty"`                                   // progress carries information about module status and grade for a specific student
	Assignment     *Assignment            `protobuf:"bytes,32,opt,name=assignment,proto3" json:"assignment,omitempty"`                               // assignment carries parameters of the assignment for a given student (member) or a class (group)
}

func (x *Module) Reset() {
	*x = Module{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_course_module_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Module) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Module) ProtoMessage() {}

func (x *Module) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_course_module_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Module.ProtoReflect.Descriptor instead.
func (*Module) Descriptor() ([]byte, []int) {
	return file_eolymp_course_module_proto_rawDescGZIP(), []int{0}
}

func (x *Module) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Module) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *Module) GetDraft() bool {
	if x != nil {
		return x.Draft
	}
	return false
}

func (x *Module) GetExtra() bool {
	if x != nil {
		return x.Extra
	}
	return false
}

func (x *Module) GetWeight() float32 {
	if x != nil {
		return x.Weight
	}
	return 0
}

func (x *Module) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Module) GetImageUrl() string {
	if x != nil {
		return x.ImageUrl
	}
	return ""
}

func (x *Module) GetIndex() uint32 {
	if x != nil {
		return x.Index
	}
	return 0
}

func (x *Module) GetDescription() *ecm.Content {
	if x != nil {
		return x.Description
	}
	return nil
}

func (x *Module) GetStartAfter() *timestamppb.Timestamp {
	if x != nil {
		return x.StartAfter
	}
	return nil
}

func (x *Module) GetCompleteBefore() *timestamppb.Timestamp {
	if x != nil {
		return x.CompleteBefore
	}
	return nil
}

func (x *Module) GetDuration() uint32 {
	if x != nil {
		return x.Duration
	}
	return 0
}

func (x *Module) GetProgress() *Module_Progress {
	if x != nil {
		return x.Progress
	}
	return nil
}

func (x *Module) GetAssignment() *Assignment {
	if x != nil {
		return x.Assignment
	}
	return nil
}

type Module_Progress struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status         Assignment_Status      `protobuf:"varint,10,opt,name=status,proto3,enum=eolymp.course.Assignment_Status" json:"status,omitempty"`
	Percentage     float32                `protobuf:"fixed32,40,opt,name=percentage,proto3" json:"percentage,omitempty"`
	Grade          uint32                 `protobuf:"varint,41,opt,name=grade,proto3" json:"grade,omitempty"`
	GradeAutomatic uint32                 `protobuf:"varint,42,opt,name=grade_automatic,json=gradeAutomatic,proto3" json:"grade_automatic,omitempty"`
	GradeOverride  uint32                 `protobuf:"varint,43,opt,name=grade_override,json=gradeOverride,proto3" json:"grade_override,omitempty"`
	Excused        bool                   `protobuf:"varint,44,opt,name=excused,proto3" json:"excused,omitempty"`
	StartAfter     *timestamppb.Timestamp `protobuf:"bytes,11,opt,name=start_after,json=startAfter,proto3" json:"start_after,omitempty"`             // optionally, time by when assignment should be complete
	CompleteBefore *timestamppb.Timestamp `protobuf:"bytes,12,opt,name=complete_before,json=completeBefore,proto3" json:"complete_before,omitempty"` // optionally, time by when assignment should be complete
	Duration       uint32                 `protobuf:"varint,13,opt,name=duration,proto3" json:"duration,omitempty"`                                  // optionally, duration of the assignment in seconds
	Upsolve        bool                   `protobuf:"varint,30,opt,name=upsolve,proto3" json:"upsolve,omitempty"`                                    // if true the task will be available after time runs out
	AssignedAt     *timestamppb.Timestamp `protobuf:"bytes,20,opt,name=assigned_at,json=assignedAt,proto3" json:"assigned_at,omitempty"`             // read-only, time when assignment was created
	StartedAt      *timestamppb.Timestamp `protobuf:"bytes,25,opt,name=started_at,json=startedAt,proto3" json:"started_at,omitempty"`                // read-only, time when assignment has been started (via StartAssignment API)
	CompletedAt    *timestamppb.Timestamp `protobuf:"bytes,26,opt,name=completed_at,json=completedAt,proto3" json:"completed_at,omitempty"`          // read-only, time when assignment will complete (started_at + duration)
}

func (x *Module_Progress) Reset() {
	*x = Module_Progress{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_course_module_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Module_Progress) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Module_Progress) ProtoMessage() {}

func (x *Module_Progress) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_course_module_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Module_Progress.ProtoReflect.Descriptor instead.
func (*Module_Progress) Descriptor() ([]byte, []int) {
	return file_eolymp_course_module_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Module_Progress) GetStatus() Assignment_Status {
	if x != nil {
		return x.Status
	}
	return Assignment_UNKNOWN_STATUS
}

func (x *Module_Progress) GetPercentage() float32 {
	if x != nil {
		return x.Percentage
	}
	return 0
}

func (x *Module_Progress) GetGrade() uint32 {
	if x != nil {
		return x.Grade
	}
	return 0
}

func (x *Module_Progress) GetGradeAutomatic() uint32 {
	if x != nil {
		return x.GradeAutomatic
	}
	return 0
}

func (x *Module_Progress) GetGradeOverride() uint32 {
	if x != nil {
		return x.GradeOverride
	}
	return 0
}

func (x *Module_Progress) GetExcused() bool {
	if x != nil {
		return x.Excused
	}
	return false
}

func (x *Module_Progress) GetStartAfter() *timestamppb.Timestamp {
	if x != nil {
		return x.StartAfter
	}
	return nil
}

func (x *Module_Progress) GetCompleteBefore() *timestamppb.Timestamp {
	if x != nil {
		return x.CompleteBefore
	}
	return nil
}

func (x *Module_Progress) GetDuration() uint32 {
	if x != nil {
		return x.Duration
	}
	return 0
}

func (x *Module_Progress) GetUpsolve() bool {
	if x != nil {
		return x.Upsolve
	}
	return false
}

func (x *Module_Progress) GetAssignedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.AssignedAt
	}
	return nil
}

func (x *Module_Progress) GetStartedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.StartedAt
	}
	return nil
}

func (x *Module_Progress) GetCompletedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CompletedAt
	}
	return nil
}

var File_eolymp_course_module_proto protoreflect.FileDescriptor

var file_eolymp_course_module_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x2f,
	0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x65, 0x6f,
	0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x1a, 0x1e, 0x65, 0x6f, 0x6c,
	0x79, 0x6d, 0x70, 0x2f, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x2f, 0x61, 0x73, 0x73, 0x69, 0x67,
	0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x18, 0x65, 0x6f, 0x6c,
	0x79, 0x6d, 0x70, 0x2f, 0x65, 0x63, 0x6d, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc0, 0x09, 0x0a, 0x06, 0x4d, 0x6f, 0x64, 0x75, 0x6c,
	0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x75, 0x72, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x64, 0x72, 0x61, 0x66, 0x74, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x05, 0x64, 0x72, 0x61, 0x66, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x78, 0x74,
	0x72, 0x61, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x65, 0x78, 0x74, 0x72, 0x61, 0x12,
	0x16, 0x0a, 0x06, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x02, 0x52,
	0x06, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x69,
	0x6d, 0x61, 0x67, 0x65, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x69, 0x6d, 0x61, 0x67, 0x65, 0x55, 0x72, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6e, 0x64, 0x65,
	0x78, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x35,
	0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x65, 0x63, 0x6d,
	0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x3b, 0x0a, 0x0b, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x61,
	0x66, 0x74, 0x65, 0x72, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x41, 0x66, 0x74,
	0x65, 0x72, 0x12, 0x43, 0x0a, 0x0f, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x5f, 0x62,
	0x65, 0x66, 0x6f, 0x72, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0e, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74,
	0x65, 0x42, 0x65, 0x66, 0x6f, 0x72, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x3a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x18,
	0x1f, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x63,
	0x6f, 0x75, 0x72, 0x73, 0x65, 0x2e, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x2e, 0x50, 0x72, 0x6f,
	0x67, 0x72, 0x65, 0x73, 0x73, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x12,
	0x39, 0x0a, 0x0a, 0x61, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x20, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x63, 0x6f, 0x75,
	0x72, 0x73, 0x65, 0x2e, 0x41, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x0a,
	0x61, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x1a, 0xd3, 0x04, 0x0a, 0x08, 0x50,
	0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x12, 0x38, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x20, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70,
	0x2e, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x2e, 0x41, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x6d, 0x65,
	0x6e, 0x74, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x61, 0x67, 0x65, 0x18,
	0x28, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0a, 0x70, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x61, 0x67,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x67, 0x72, 0x61, 0x64, 0x65, 0x18, 0x29, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x05, 0x67, 0x72, 0x61, 0x64, 0x65, 0x12, 0x27, 0x0a, 0x0f, 0x67, 0x72, 0x61, 0x64, 0x65,
	0x5f, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x69, 0x63, 0x18, 0x2a, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x0e, 0x67, 0x72, 0x61, 0x64, 0x65, 0x41, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x69, 0x63,
	0x12, 0x25, 0x0a, 0x0e, 0x67, 0x72, 0x61, 0x64, 0x65, 0x5f, 0x6f, 0x76, 0x65, 0x72, 0x72, 0x69,
	0x64, 0x65, 0x18, 0x2b, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0d, 0x67, 0x72, 0x61, 0x64, 0x65, 0x4f,
	0x76, 0x65, 0x72, 0x72, 0x69, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x78, 0x63, 0x75, 0x73,
	0x65, 0x64, 0x18, 0x2c, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x65, 0x78, 0x63, 0x75, 0x73, 0x65,
	0x64, 0x12, 0x3b, 0x0a, 0x0b, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x61, 0x66, 0x74, 0x65, 0x72,
	0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x41, 0x66, 0x74, 0x65, 0x72, 0x12, 0x43,
	0x0a, 0x0f, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x5f, 0x62, 0x65, 0x66, 0x6f, 0x72,
	0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x0e, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x65, 0x66,
	0x6f, 0x72, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x0d, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x18, 0x0a, 0x07, 0x75, 0x70, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x07, 0x75, 0x70, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x12, 0x3b, 0x0a, 0x0b, 0x61, 0x73, 0x73,
	0x69, 0x67, 0x6e, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x14, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x61, 0x73, 0x73, 0x69,
	0x67, 0x6e, 0x65, 0x64, 0x41, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x65,
	0x64, 0x5f, 0x61, 0x74, 0x18, 0x19, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x12, 0x3d, 0x0a, 0x0c, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x5f, 0x61,
	0x74, 0x18, 0x1a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x0b, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x22, 0x67, 0x0a, 0x05, 0x45, 0x78, 0x74, 0x72, 0x61, 0x12, 0x11, 0x0a, 0x0d, 0x55, 0x4e, 0x4b,
	0x4e, 0x4f, 0x57, 0x4e, 0x5f, 0x45, 0x58, 0x54, 0x52, 0x41, 0x10, 0x00, 0x12, 0x15, 0x0a, 0x11,
	0x44, 0x45, 0x53, 0x43, 0x52, 0x49, 0x50, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x56, 0x41, 0x4c, 0x55,
	0x45, 0x10, 0x01, 0x12, 0x16, 0x0a, 0x12, 0x44, 0x45, 0x53, 0x43, 0x52, 0x49, 0x50, 0x54, 0x49,
	0x4f, 0x4e, 0x5f, 0x52, 0x45, 0x4e, 0x44, 0x45, 0x52, 0x10, 0x02, 0x12, 0x0c, 0x0a, 0x08, 0x50,
	0x52, 0x4f, 0x47, 0x52, 0x45, 0x53, 0x53, 0x10, 0x04, 0x12, 0x0e, 0x0a, 0x0a, 0x41, 0x53, 0x53,
	0x49, 0x47, 0x4e, 0x4d, 0x45, 0x4e, 0x54, 0x10, 0x05, 0x42, 0x2f, 0x5a, 0x2d, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x67,
	0x6f, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x63, 0x6f, 0x75,
	0x72, 0x73, 0x65, 0x3b, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_eolymp_course_module_proto_rawDescOnce sync.Once
	file_eolymp_course_module_proto_rawDescData = file_eolymp_course_module_proto_rawDesc
)

func file_eolymp_course_module_proto_rawDescGZIP() []byte {
	file_eolymp_course_module_proto_rawDescOnce.Do(func() {
		file_eolymp_course_module_proto_rawDescData = protoimpl.X.CompressGZIP(file_eolymp_course_module_proto_rawDescData)
	})
	return file_eolymp_course_module_proto_rawDescData
}

var file_eolymp_course_module_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_eolymp_course_module_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_eolymp_course_module_proto_goTypes = []any{
	(Module_Extra)(0),             // 0: eolymp.course.Module.Extra
	(*Module)(nil),                // 1: eolymp.course.Module
	(*Module_Progress)(nil),       // 2: eolymp.course.Module.Progress
	(*ecm.Content)(nil),           // 3: eolymp.ecm.Content
	(*timestamppb.Timestamp)(nil), // 4: google.protobuf.Timestamp
	(*Assignment)(nil),            // 5: eolymp.course.Assignment
	(Assignment_Status)(0),        // 6: eolymp.course.Assignment.Status
}
var file_eolymp_course_module_proto_depIdxs = []int32{
	3,  // 0: eolymp.course.Module.description:type_name -> eolymp.ecm.Content
	4,  // 1: eolymp.course.Module.start_after:type_name -> google.protobuf.Timestamp
	4,  // 2: eolymp.course.Module.complete_before:type_name -> google.protobuf.Timestamp
	2,  // 3: eolymp.course.Module.progress:type_name -> eolymp.course.Module.Progress
	5,  // 4: eolymp.course.Module.assignment:type_name -> eolymp.course.Assignment
	6,  // 5: eolymp.course.Module.Progress.status:type_name -> eolymp.course.Assignment.Status
	4,  // 6: eolymp.course.Module.Progress.start_after:type_name -> google.protobuf.Timestamp
	4,  // 7: eolymp.course.Module.Progress.complete_before:type_name -> google.protobuf.Timestamp
	4,  // 8: eolymp.course.Module.Progress.assigned_at:type_name -> google.protobuf.Timestamp
	4,  // 9: eolymp.course.Module.Progress.started_at:type_name -> google.protobuf.Timestamp
	4,  // 10: eolymp.course.Module.Progress.completed_at:type_name -> google.protobuf.Timestamp
	11, // [11:11] is the sub-list for method output_type
	11, // [11:11] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_eolymp_course_module_proto_init() }
func file_eolymp_course_module_proto_init() {
	if File_eolymp_course_module_proto != nil {
		return
	}
	file_eolymp_course_assignment_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_eolymp_course_module_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*Module); i {
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
		file_eolymp_course_module_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*Module_Progress); i {
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
			RawDescriptor: file_eolymp_course_module_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_eolymp_course_module_proto_goTypes,
		DependencyIndexes: file_eolymp_course_module_proto_depIdxs,
		EnumInfos:         file_eolymp_course_module_proto_enumTypes,
		MessageInfos:      file_eolymp_course_module_proto_msgTypes,
	}.Build()
	File_eolymp_course_module_proto = out.File
	file_eolymp_course_module_proto_rawDesc = nil
	file_eolymp_course_module_proto_goTypes = nil
	file_eolymp_course_module_proto_depIdxs = nil
}
