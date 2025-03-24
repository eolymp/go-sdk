// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v4.24.4
// source: eolymp/course/course_service.proto

package course

import (
	_ "github.com/eolymp/go-sdk/eolymp/annotations"
	wellknown "github.com/eolymp/go-sdk/eolymp/wellknown"
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

type ListCoursesInput_Sortable int32

const (
	ListCoursesInput_DEFAULT ListCoursesInput_Sortable = 0
)

// Enum value maps for ListCoursesInput_Sortable.
var (
	ListCoursesInput_Sortable_name = map[int32]string{
		0: "DEFAULT",
	}
	ListCoursesInput_Sortable_value = map[string]int32{
		"DEFAULT": 0,
	}
)

func (x ListCoursesInput_Sortable) Enum() *ListCoursesInput_Sortable {
	p := new(ListCoursesInput_Sortable)
	*p = x
	return p
}

func (x ListCoursesInput_Sortable) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ListCoursesInput_Sortable) Descriptor() protoreflect.EnumDescriptor {
	return file_eolymp_course_course_service_proto_enumTypes[0].Descriptor()
}

func (ListCoursesInput_Sortable) Type() protoreflect.EnumType {
	return &file_eolymp_course_course_service_proto_enumTypes[0]
}

func (x ListCoursesInput_Sortable) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ListCoursesInput_Sortable.Descriptor instead.
func (ListCoursesInput_Sortable) EnumDescriptor() ([]byte, []int) {
	return file_eolymp_course_course_service_proto_rawDescGZIP(), []int{0, 0}
}

type UpdateCourseInput_Patch int32

const (
	UpdateCourseInput_ALL         UpdateCourseInput_Patch = 0
	UpdateCourseInput_LOCALE      UpdateCourseInput_Patch = 1
	UpdateCourseInput_NAME        UpdateCourseInput_Patch = 2
	UpdateCourseInput_DESCRIPTION UpdateCourseInput_Patch = 3
	UpdateCourseInput_IMAGE       UpdateCourseInput_Patch = 4
	UpdateCourseInput_VISIBILITY  UpdateCourseInput_Patch = 5
	UpdateCourseInput_DURATION    UpdateCourseInput_Patch = 6
	UpdateCourseInput_TOPICS      UpdateCourseInput_Patch = 7
)

// Enum value maps for UpdateCourseInput_Patch.
var (
	UpdateCourseInput_Patch_name = map[int32]string{
		0: "ALL",
		1: "LOCALE",
		2: "NAME",
		3: "DESCRIPTION",
		4: "IMAGE",
		5: "VISIBILITY",
		6: "DURATION",
		7: "TOPICS",
	}
	UpdateCourseInput_Patch_value = map[string]int32{
		"ALL":         0,
		"LOCALE":      1,
		"NAME":        2,
		"DESCRIPTION": 3,
		"IMAGE":       4,
		"VISIBILITY":  5,
		"DURATION":    6,
		"TOPICS":      7,
	}
)

func (x UpdateCourseInput_Patch) Enum() *UpdateCourseInput_Patch {
	p := new(UpdateCourseInput_Patch)
	*p = x
	return p
}

func (x UpdateCourseInput_Patch) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (UpdateCourseInput_Patch) Descriptor() protoreflect.EnumDescriptor {
	return file_eolymp_course_course_service_proto_enumTypes[1].Descriptor()
}

func (UpdateCourseInput_Patch) Type() protoreflect.EnumType {
	return &file_eolymp_course_course_service_proto_enumTypes[1]
}

func (x UpdateCourseInput_Patch) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use UpdateCourseInput_Patch.Descriptor instead.
func (UpdateCourseInput_Patch) EnumDescriptor() ([]byte, []int) {
	return file_eolymp_course_course_service_proto_rawDescGZIP(), []int{6, 0}
}

type CopyCourseInput_Scope int32

const (
	CopyCourseInput_ALL         CopyCourseInput_Scope = 0
	CopyCourseInput_MODULES     CopyCourseInput_Scope = 1
	CopyCourseInput_STUDENTS    CopyCourseInput_Scope = 2
	CopyCourseInput_ASSIGNMENTS CopyCourseInput_Scope = 3
	CopyCourseInput_PERMISSIONS CopyCourseInput_Scope = 4
)

// Enum value maps for CopyCourseInput_Scope.
var (
	CopyCourseInput_Scope_name = map[int32]string{
		0: "ALL",
		1: "MODULES",
		2: "STUDENTS",
		3: "ASSIGNMENTS",
		4: "PERMISSIONS",
	}
	CopyCourseInput_Scope_value = map[string]int32{
		"ALL":         0,
		"MODULES":     1,
		"STUDENTS":    2,
		"ASSIGNMENTS": 3,
		"PERMISSIONS": 4,
	}
)

func (x CopyCourseInput_Scope) Enum() *CopyCourseInput_Scope {
	p := new(CopyCourseInput_Scope)
	*p = x
	return p
}

func (x CopyCourseInput_Scope) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CopyCourseInput_Scope) Descriptor() protoreflect.EnumDescriptor {
	return file_eolymp_course_course_service_proto_enumTypes[2].Descriptor()
}

func (CopyCourseInput_Scope) Type() protoreflect.EnumType {
	return &file_eolymp_course_course_service_proto_enumTypes[2]
}

func (x CopyCourseInput_Scope) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CopyCourseInput_Scope.Descriptor instead.
func (CopyCourseInput_Scope) EnumDescriptor() ([]byte, []int) {
	return file_eolymp_course_course_service_proto_rawDescGZIP(), []int{10, 0}
}

type ListCoursesInput struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// pagination
	Offset int32 `protobuf:"varint,10,opt,name=offset,proto3" json:"offset,omitempty"`
	Size   int32 `protobuf:"varint,11,opt,name=size,proto3" json:"size,omitempty"`
	// data filters
	Search        string                    `protobuf:"bytes,20,opt,name=search,proto3" json:"search,omitempty"`
	Filters       *ListCoursesInput_Filter  `protobuf:"bytes,40,opt,name=filters,proto3" json:"filters,omitempty"`
	Sort          ListCoursesInput_Sortable `protobuf:"varint,50,opt,name=sort,proto3,enum=eolymp.course.ListCoursesInput_Sortable" json:"sort,omitempty"`
	Order         wellknown.Direction       `protobuf:"varint,51,opt,name=order,proto3,enum=eolymp.wellknown.Direction" json:"order,omitempty"`
	Extra         []Course_Extra            `protobuf:"varint,1123,rep,packed,name=extra,proto3,enum=eolymp.course.Course_Extra" json:"extra,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListCoursesInput) Reset() {
	*x = ListCoursesInput{}
	mi := &file_eolymp_course_course_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListCoursesInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListCoursesInput) ProtoMessage() {}

func (x *ListCoursesInput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_course_course_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListCoursesInput.ProtoReflect.Descriptor instead.
func (*ListCoursesInput) Descriptor() ([]byte, []int) {
	return file_eolymp_course_course_service_proto_rawDescGZIP(), []int{0}
}

func (x *ListCoursesInput) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *ListCoursesInput) GetSize() int32 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *ListCoursesInput) GetSearch() string {
	if x != nil {
		return x.Search
	}
	return ""
}

func (x *ListCoursesInput) GetFilters() *ListCoursesInput_Filter {
	if x != nil {
		return x.Filters
	}
	return nil
}

func (x *ListCoursesInput) GetSort() ListCoursesInput_Sortable {
	if x != nil {
		return x.Sort
	}
	return ListCoursesInput_DEFAULT
}

func (x *ListCoursesInput) GetOrder() wellknown.Direction {
	if x != nil {
		return x.Order
	}
	return wellknown.Direction(0)
}

func (x *ListCoursesInput) GetExtra() []Course_Extra {
	if x != nil {
		return x.Extra
	}
	return nil
}

type ListCoursesOutput struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Total         int32                  `protobuf:"varint,1,opt,name=total,proto3" json:"total,omitempty"`
	Items         []*Course              `protobuf:"bytes,2,rep,name=items,proto3" json:"items,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListCoursesOutput) Reset() {
	*x = ListCoursesOutput{}
	mi := &file_eolymp_course_course_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListCoursesOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListCoursesOutput) ProtoMessage() {}

func (x *ListCoursesOutput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_course_course_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListCoursesOutput.ProtoReflect.Descriptor instead.
func (*ListCoursesOutput) Descriptor() ([]byte, []int) {
	return file_eolymp_course_course_service_proto_rawDescGZIP(), []int{1}
}

func (x *ListCoursesOutput) GetTotal() int32 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *ListCoursesOutput) GetItems() []*Course {
	if x != nil {
		return x.Items
	}
	return nil
}

type DescribeCourseInput struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	CourseId      string                 `protobuf:"bytes,1,opt,name=course_id,json=courseId,proto3" json:"course_id,omitempty"`
	Extra         []Course_Extra         `protobuf:"varint,1123,rep,packed,name=extra,proto3,enum=eolymp.course.Course_Extra" json:"extra,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DescribeCourseInput) Reset() {
	*x = DescribeCourseInput{}
	mi := &file_eolymp_course_course_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DescribeCourseInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescribeCourseInput) ProtoMessage() {}

func (x *DescribeCourseInput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_course_course_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DescribeCourseInput.ProtoReflect.Descriptor instead.
func (*DescribeCourseInput) Descriptor() ([]byte, []int) {
	return file_eolymp_course_course_service_proto_rawDescGZIP(), []int{2}
}

func (x *DescribeCourseInput) GetCourseId() string {
	if x != nil {
		return x.CourseId
	}
	return ""
}

func (x *DescribeCourseInput) GetExtra() []Course_Extra {
	if x != nil {
		return x.Extra
	}
	return nil
}

type DescribeCourseOutput struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Course        *Course                `protobuf:"bytes,1,opt,name=course,proto3" json:"course,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DescribeCourseOutput) Reset() {
	*x = DescribeCourseOutput{}
	mi := &file_eolymp_course_course_service_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DescribeCourseOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescribeCourseOutput) ProtoMessage() {}

func (x *DescribeCourseOutput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_course_course_service_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DescribeCourseOutput.ProtoReflect.Descriptor instead.
func (*DescribeCourseOutput) Descriptor() ([]byte, []int) {
	return file_eolymp_course_course_service_proto_rawDescGZIP(), []int{3}
}

func (x *DescribeCourseOutput) GetCourse() *Course {
	if x != nil {
		return x.Course
	}
	return nil
}

type CreateCourseInput struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Course        *Course                `protobuf:"bytes,1,opt,name=course,proto3" json:"course,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateCourseInput) Reset() {
	*x = CreateCourseInput{}
	mi := &file_eolymp_course_course_service_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateCourseInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCourseInput) ProtoMessage() {}

func (x *CreateCourseInput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_course_course_service_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCourseInput.ProtoReflect.Descriptor instead.
func (*CreateCourseInput) Descriptor() ([]byte, []int) {
	return file_eolymp_course_course_service_proto_rawDescGZIP(), []int{4}
}

func (x *CreateCourseInput) GetCourse() *Course {
	if x != nil {
		return x.Course
	}
	return nil
}

type CreateCourseOutput struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	CourseId      string                 `protobuf:"bytes,1,opt,name=course_id,json=courseId,proto3" json:"course_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateCourseOutput) Reset() {
	*x = CreateCourseOutput{}
	mi := &file_eolymp_course_course_service_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateCourseOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCourseOutput) ProtoMessage() {}

func (x *CreateCourseOutput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_course_course_service_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCourseOutput.ProtoReflect.Descriptor instead.
func (*CreateCourseOutput) Descriptor() ([]byte, []int) {
	return file_eolymp_course_course_service_proto_rawDescGZIP(), []int{5}
}

func (x *CreateCourseOutput) GetCourseId() string {
	if x != nil {
		return x.CourseId
	}
	return ""
}

type UpdateCourseInput struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// specify list of fields to update, if empty all fields are updated
	Patch         []UpdateCourseInput_Patch `protobuf:"varint,1,rep,packed,name=patch,proto3,enum=eolymp.course.UpdateCourseInput_Patch" json:"patch,omitempty"`
	CourseId      string                    `protobuf:"bytes,2,opt,name=course_id,json=courseId,proto3" json:"course_id,omitempty"`
	Course        *Course                   `protobuf:"bytes,3,opt,name=course,proto3" json:"course,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateCourseInput) Reset() {
	*x = UpdateCourseInput{}
	mi := &file_eolymp_course_course_service_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateCourseInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateCourseInput) ProtoMessage() {}

func (x *UpdateCourseInput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_course_course_service_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateCourseInput.ProtoReflect.Descriptor instead.
func (*UpdateCourseInput) Descriptor() ([]byte, []int) {
	return file_eolymp_course_course_service_proto_rawDescGZIP(), []int{6}
}

func (x *UpdateCourseInput) GetPatch() []UpdateCourseInput_Patch {
	if x != nil {
		return x.Patch
	}
	return nil
}

func (x *UpdateCourseInput) GetCourseId() string {
	if x != nil {
		return x.CourseId
	}
	return ""
}

func (x *UpdateCourseInput) GetCourse() *Course {
	if x != nil {
		return x.Course
	}
	return nil
}

type UpdateCourseOutput struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateCourseOutput) Reset() {
	*x = UpdateCourseOutput{}
	mi := &file_eolymp_course_course_service_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateCourseOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateCourseOutput) ProtoMessage() {}

func (x *UpdateCourseOutput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_course_course_service_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateCourseOutput.ProtoReflect.Descriptor instead.
func (*UpdateCourseOutput) Descriptor() ([]byte, []int) {
	return file_eolymp_course_course_service_proto_rawDescGZIP(), []int{7}
}

type DeleteCourseInput struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	CourseId      string                 `protobuf:"bytes,1,opt,name=course_id,json=courseId,proto3" json:"course_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteCourseInput) Reset() {
	*x = DeleteCourseInput{}
	mi := &file_eolymp_course_course_service_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteCourseInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteCourseInput) ProtoMessage() {}

func (x *DeleteCourseInput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_course_course_service_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteCourseInput.ProtoReflect.Descriptor instead.
func (*DeleteCourseInput) Descriptor() ([]byte, []int) {
	return file_eolymp_course_course_service_proto_rawDescGZIP(), []int{8}
}

func (x *DeleteCourseInput) GetCourseId() string {
	if x != nil {
		return x.CourseId
	}
	return ""
}

type DeleteCourseOutput struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteCourseOutput) Reset() {
	*x = DeleteCourseOutput{}
	mi := &file_eolymp_course_course_service_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteCourseOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteCourseOutput) ProtoMessage() {}

func (x *DeleteCourseOutput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_course_course_service_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteCourseOutput.ProtoReflect.Descriptor instead.
func (*DeleteCourseOutput) Descriptor() ([]byte, []int) {
	return file_eolymp_course_course_service_proto_rawDescGZIP(), []int{9}
}

type CopyCourseInput struct {
	state         protoimpl.MessageState  `protogen:"open.v1"`
	CourseId      string                  `protobuf:"bytes,1,opt,name=course_id,json=courseId,proto3" json:"course_id,omitempty"`
	CopyScope     []CopyCourseInput_Scope `protobuf:"varint,2,rep,packed,name=copy_scope,json=copyScope,proto3,enum=eolymp.course.CopyCourseInput_Scope" json:"copy_scope,omitempty"`
	CopyName      string                  `protobuf:"bytes,3,opt,name=copy_name,json=copyName,proto3" json:"copy_name,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CopyCourseInput) Reset() {
	*x = CopyCourseInput{}
	mi := &file_eolymp_course_course_service_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CopyCourseInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CopyCourseInput) ProtoMessage() {}

func (x *CopyCourseInput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_course_course_service_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CopyCourseInput.ProtoReflect.Descriptor instead.
func (*CopyCourseInput) Descriptor() ([]byte, []int) {
	return file_eolymp_course_course_service_proto_rawDescGZIP(), []int{10}
}

func (x *CopyCourseInput) GetCourseId() string {
	if x != nil {
		return x.CourseId
	}
	return ""
}

func (x *CopyCourseInput) GetCopyScope() []CopyCourseInput_Scope {
	if x != nil {
		return x.CopyScope
	}
	return nil
}

func (x *CopyCourseInput) GetCopyName() string {
	if x != nil {
		return x.CopyName
	}
	return ""
}

type CopyCourseOutput struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	CopyCourseId  string                 `protobuf:"bytes,1,opt,name=copy_course_id,json=copyCourseId,proto3" json:"copy_course_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CopyCourseOutput) Reset() {
	*x = CopyCourseOutput{}
	mi := &file_eolymp_course_course_service_proto_msgTypes[11]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CopyCourseOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CopyCourseOutput) ProtoMessage() {}

func (x *CopyCourseOutput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_course_course_service_proto_msgTypes[11]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CopyCourseOutput.ProtoReflect.Descriptor instead.
func (*CopyCourseOutput) Descriptor() ([]byte, []int) {
	return file_eolymp_course_course_service_proto_rawDescGZIP(), []int{11}
}

func (x *CopyCourseOutput) GetCopyCourseId() string {
	if x != nil {
		return x.CopyCourseId
	}
	return ""
}

type ListCoursesInput_Filter struct {
	state         protoimpl.MessageState      `protogen:"open.v1"`
	Id            []*wellknown.ExpressionID   `protobuf:"bytes,10,rep,name=id,proto3" json:"id,omitempty"`
	TopicId       []*wellknown.ExpressionID   `protobuf:"bytes,11,rep,name=topic_id,json=topicId,proto3" json:"topic_id,omitempty"`
	Locale        []*wellknown.ExpressionEnum `protobuf:"bytes,12,rep,name=locale,proto3" json:"locale,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListCoursesInput_Filter) Reset() {
	*x = ListCoursesInput_Filter{}
	mi := &file_eolymp_course_course_service_proto_msgTypes[12]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListCoursesInput_Filter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListCoursesInput_Filter) ProtoMessage() {}

func (x *ListCoursesInput_Filter) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_course_course_service_proto_msgTypes[12]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListCoursesInput_Filter.ProtoReflect.Descriptor instead.
func (*ListCoursesInput_Filter) Descriptor() ([]byte, []int) {
	return file_eolymp_course_course_service_proto_rawDescGZIP(), []int{0, 0}
}

func (x *ListCoursesInput_Filter) GetId() []*wellknown.ExpressionID {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *ListCoursesInput_Filter) GetTopicId() []*wellknown.ExpressionID {
	if x != nil {
		return x.TopicId
	}
	return nil
}

func (x *ListCoursesInput_Filter) GetLocale() []*wellknown.ExpressionEnum {
	if x != nil {
		return x.Locale
	}
	return nil
}

var File_eolymp_course_course_service_proto protoreflect.FileDescriptor

const file_eolymp_course_course_service_proto_rawDesc = "" +
	"\n" +
	"\"eolymp/course/course_service.proto\x12\reolymp.course\x1a\x1deolymp/annotations/http.proto\x1a\"eolymp/annotations/ratelimit.proto\x1a\x1eeolymp/annotations/scope.proto\x1a\x1aeolymp/course/course.proto\x1a eolymp/wellknown/direction.proto\x1a!eolymp/wellknown/expression.proto\"\x86\x04\n" +
	"\x10ListCoursesInput\x12\x16\n" +
	"\x06offset\x18\n" +
	" \x01(\x05R\x06offset\x12\x12\n" +
	"\x04size\x18\v \x01(\x05R\x04size\x12\x16\n" +
	"\x06search\x18\x14 \x01(\tR\x06search\x12@\n" +
	"\afilters\x18( \x01(\v2&.eolymp.course.ListCoursesInput.FilterR\afilters\x12<\n" +
	"\x04sort\x182 \x01(\x0e2(.eolymp.course.ListCoursesInput.SortableR\x04sort\x121\n" +
	"\x05order\x183 \x01(\x0e2\x1b.eolymp.wellknown.DirectionR\x05order\x122\n" +
	"\x05extra\x18\xe3\b \x03(\x0e2\x1b.eolymp.course.Course.ExtraR\x05extra\x1a\xad\x01\n" +
	"\x06Filter\x12.\n" +
	"\x02id\x18\n" +
	" \x03(\v2\x1e.eolymp.wellknown.ExpressionIDR\x02id\x129\n" +
	"\btopic_id\x18\v \x03(\v2\x1e.eolymp.wellknown.ExpressionIDR\atopicId\x128\n" +
	"\x06locale\x18\f \x03(\v2 .eolymp.wellknown.ExpressionEnumR\x06locale\"\x17\n" +
	"\bSortable\x12\v\n" +
	"\aDEFAULT\x10\x00\"V\n" +
	"\x11ListCoursesOutput\x12\x14\n" +
	"\x05total\x18\x01 \x01(\x05R\x05total\x12+\n" +
	"\x05items\x18\x02 \x03(\v2\x15.eolymp.course.CourseR\x05items\"f\n" +
	"\x13DescribeCourseInput\x12\x1b\n" +
	"\tcourse_id\x18\x01 \x01(\tR\bcourseId\x122\n" +
	"\x05extra\x18\xe3\b \x03(\x0e2\x1b.eolymp.course.Course.ExtraR\x05extra\"E\n" +
	"\x14DescribeCourseOutput\x12-\n" +
	"\x06course\x18\x01 \x01(\v2\x15.eolymp.course.CourseR\x06course\"B\n" +
	"\x11CreateCourseInput\x12-\n" +
	"\x06course\x18\x01 \x01(\v2\x15.eolymp.course.CourseR\x06course\"1\n" +
	"\x12CreateCourseOutput\x12\x1b\n" +
	"\tcourse_id\x18\x01 \x01(\tR\bcourseId\"\x8b\x02\n" +
	"\x11UpdateCourseInput\x12<\n" +
	"\x05patch\x18\x01 \x03(\x0e2&.eolymp.course.UpdateCourseInput.PatchR\x05patch\x12\x1b\n" +
	"\tcourse_id\x18\x02 \x01(\tR\bcourseId\x12-\n" +
	"\x06course\x18\x03 \x01(\v2\x15.eolymp.course.CourseR\x06course\"l\n" +
	"\x05Patch\x12\a\n" +
	"\x03ALL\x10\x00\x12\n" +
	"\n" +
	"\x06LOCALE\x10\x01\x12\b\n" +
	"\x04NAME\x10\x02\x12\x0f\n" +
	"\vDESCRIPTION\x10\x03\x12\t\n" +
	"\x05IMAGE\x10\x04\x12\x0e\n" +
	"\n" +
	"VISIBILITY\x10\x05\x12\f\n" +
	"\bDURATION\x10\x06\x12\n" +
	"\n" +
	"\x06TOPICS\x10\a\"\x14\n" +
	"\x12UpdateCourseOutput\"0\n" +
	"\x11DeleteCourseInput\x12\x1b\n" +
	"\tcourse_id\x18\x01 \x01(\tR\bcourseId\"\x14\n" +
	"\x12DeleteCourseOutput\"\xdf\x01\n" +
	"\x0fCopyCourseInput\x12\x1b\n" +
	"\tcourse_id\x18\x01 \x01(\tR\bcourseId\x12C\n" +
	"\n" +
	"copy_scope\x18\x02 \x03(\x0e2$.eolymp.course.CopyCourseInput.ScopeR\tcopyScope\x12\x1b\n" +
	"\tcopy_name\x18\x03 \x01(\tR\bcopyName\"M\n" +
	"\x05Scope\x12\a\n" +
	"\x03ALL\x10\x00\x12\v\n" +
	"\aMODULES\x10\x01\x12\f\n" +
	"\bSTUDENTS\x10\x02\x12\x0f\n" +
	"\vASSIGNMENTS\x10\x03\x12\x0f\n" +
	"\vPERMISSIONS\x10\x04\"8\n" +
	"\x10CopyCourseOutput\x12$\n" +
	"\x0ecopy_course_id\x18\x01 \x01(\tR\fcopyCourseId2\xab\a\n" +
	"\rCourseService\x12\x8f\x01\n" +
	"\fCreateCourse\x12 .eolymp.course.CreateCourseInput\x1a!.eolymp.course.CreateCourseOutput\":\xea\xe2\n" +
	"\v\xf5\xe2\n" +
	"\x00\x00\x80?\xf8\xe2\n" +
	"\x05\x82\xe3\n" +
	"\x17\x8a\xe3\n" +
	"\x13course:course:write\x82\xd3\xe4\x93\x02\n" +
	"\x1a\b/courses\x12\x9b\x01\n" +
	"\fUpdateCourse\x12 .eolymp.course.UpdateCourseInput\x1a!.eolymp.course.UpdateCourseOutput\"F\xea\xe2\n" +
	"\v\xf5\xe2\n" +
	"\x00\x00\x80?\xf8\xe2\n" +
	"\x05\x82\xe3\n" +
	"\x17\x8a\xe3\n" +
	"\x13course:course:write\x82\xd3\xe4\x93\x02\x16\x1a\x14/courses/{course_id}\x12\x9b\x01\n" +
	"\fDeleteCourse\x12 .eolymp.course.DeleteCourseInput\x1a!.eolymp.course.DeleteCourseOutput\"F\xea\xe2\n" +
	"\v\xf5\xe2\n" +
	"\x00\x00\x80?\xf8\xe2\n" +
	"\x05\x82\xe3\n" +
	"\x17\x8a\xe3\n" +
	"\x13course:course:write\x82\xd3\xe4\x93\x02\x16*\x14/courses/{course_id}\x12\xa0\x01\n" +
	"\x0eDescribeCourse\x12\".eolymp.course.DescribeCourseInput\x1a#.eolymp.course.DescribeCourseOutput\"E\xea\xe2\n" +
	"\v\xf5\xe2\n" +
	"\x00\x00\xa0A\xf8\xe2\n" +
	"d\x82\xe3\n" +
	"\x16\x8a\xe3\n" +
	"\x12course:course:read\x82\xd3\xe4\x93\x02\x16\x12\x14/courses/{course_id}\x12\x8b\x01\n" +
	"\vListCourses\x12\x1f.eolymp.course.ListCoursesInput\x1a .eolymp.course.ListCoursesOutput\"9\xea\xe2\n" +
	"\v\xf5\xe2\n" +
	"\x00\x00\xa0A\xf8\xe2\n" +
	"d\x82\xe3\n" +
	"\x16\x8a\xe3\n" +
	"\x12course:course:read\x82\xd3\xe4\x93\x02\n" +
	"\x12\b/courses\x12\x9a\x01\n" +
	"\n" +
	"CopyCourse\x12\x1e.eolymp.course.CopyCourseInput\x1a\x1f.eolymp.course.CopyCourseOutput\"K\xea\xe2\n" +
	"\v\xf5\xe2\n" +
	"\x00\x00\x80?\xf8\xe2\n" +
	"\x05\x82\xe3\n" +
	"\x17\x8a\xe3\n" +
	"\x13course:course:write\x82\xd3\xe4\x93\x02\x1b\"\x19/courses/{course_id}/copyB/Z-github.com/eolymp/go-sdk/eolymp/course;courseb\x06proto3"

var (
	file_eolymp_course_course_service_proto_rawDescOnce sync.Once
	file_eolymp_course_course_service_proto_rawDescData []byte
)

func file_eolymp_course_course_service_proto_rawDescGZIP() []byte {
	file_eolymp_course_course_service_proto_rawDescOnce.Do(func() {
		file_eolymp_course_course_service_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_eolymp_course_course_service_proto_rawDesc), len(file_eolymp_course_course_service_proto_rawDesc)))
	})
	return file_eolymp_course_course_service_proto_rawDescData
}

var file_eolymp_course_course_service_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_eolymp_course_course_service_proto_msgTypes = make([]protoimpl.MessageInfo, 13)
var file_eolymp_course_course_service_proto_goTypes = []any{
	(ListCoursesInput_Sortable)(0),   // 0: eolymp.course.ListCoursesInput.Sortable
	(UpdateCourseInput_Patch)(0),     // 1: eolymp.course.UpdateCourseInput.Patch
	(CopyCourseInput_Scope)(0),       // 2: eolymp.course.CopyCourseInput.Scope
	(*ListCoursesInput)(nil),         // 3: eolymp.course.ListCoursesInput
	(*ListCoursesOutput)(nil),        // 4: eolymp.course.ListCoursesOutput
	(*DescribeCourseInput)(nil),      // 5: eolymp.course.DescribeCourseInput
	(*DescribeCourseOutput)(nil),     // 6: eolymp.course.DescribeCourseOutput
	(*CreateCourseInput)(nil),        // 7: eolymp.course.CreateCourseInput
	(*CreateCourseOutput)(nil),       // 8: eolymp.course.CreateCourseOutput
	(*UpdateCourseInput)(nil),        // 9: eolymp.course.UpdateCourseInput
	(*UpdateCourseOutput)(nil),       // 10: eolymp.course.UpdateCourseOutput
	(*DeleteCourseInput)(nil),        // 11: eolymp.course.DeleteCourseInput
	(*DeleteCourseOutput)(nil),       // 12: eolymp.course.DeleteCourseOutput
	(*CopyCourseInput)(nil),          // 13: eolymp.course.CopyCourseInput
	(*CopyCourseOutput)(nil),         // 14: eolymp.course.CopyCourseOutput
	(*ListCoursesInput_Filter)(nil),  // 15: eolymp.course.ListCoursesInput.Filter
	(wellknown.Direction)(0),         // 16: eolymp.wellknown.Direction
	(Course_Extra)(0),                // 17: eolymp.course.Course.Extra
	(*Course)(nil),                   // 18: eolymp.course.Course
	(*wellknown.ExpressionID)(nil),   // 19: eolymp.wellknown.ExpressionID
	(*wellknown.ExpressionEnum)(nil), // 20: eolymp.wellknown.ExpressionEnum
}
var file_eolymp_course_course_service_proto_depIdxs = []int32{
	15, // 0: eolymp.course.ListCoursesInput.filters:type_name -> eolymp.course.ListCoursesInput.Filter
	0,  // 1: eolymp.course.ListCoursesInput.sort:type_name -> eolymp.course.ListCoursesInput.Sortable
	16, // 2: eolymp.course.ListCoursesInput.order:type_name -> eolymp.wellknown.Direction
	17, // 3: eolymp.course.ListCoursesInput.extra:type_name -> eolymp.course.Course.Extra
	18, // 4: eolymp.course.ListCoursesOutput.items:type_name -> eolymp.course.Course
	17, // 5: eolymp.course.DescribeCourseInput.extra:type_name -> eolymp.course.Course.Extra
	18, // 6: eolymp.course.DescribeCourseOutput.course:type_name -> eolymp.course.Course
	18, // 7: eolymp.course.CreateCourseInput.course:type_name -> eolymp.course.Course
	1,  // 8: eolymp.course.UpdateCourseInput.patch:type_name -> eolymp.course.UpdateCourseInput.Patch
	18, // 9: eolymp.course.UpdateCourseInput.course:type_name -> eolymp.course.Course
	2,  // 10: eolymp.course.CopyCourseInput.copy_scope:type_name -> eolymp.course.CopyCourseInput.Scope
	19, // 11: eolymp.course.ListCoursesInput.Filter.id:type_name -> eolymp.wellknown.ExpressionID
	19, // 12: eolymp.course.ListCoursesInput.Filter.topic_id:type_name -> eolymp.wellknown.ExpressionID
	20, // 13: eolymp.course.ListCoursesInput.Filter.locale:type_name -> eolymp.wellknown.ExpressionEnum
	7,  // 14: eolymp.course.CourseService.CreateCourse:input_type -> eolymp.course.CreateCourseInput
	9,  // 15: eolymp.course.CourseService.UpdateCourse:input_type -> eolymp.course.UpdateCourseInput
	11, // 16: eolymp.course.CourseService.DeleteCourse:input_type -> eolymp.course.DeleteCourseInput
	5,  // 17: eolymp.course.CourseService.DescribeCourse:input_type -> eolymp.course.DescribeCourseInput
	3,  // 18: eolymp.course.CourseService.ListCourses:input_type -> eolymp.course.ListCoursesInput
	13, // 19: eolymp.course.CourseService.CopyCourse:input_type -> eolymp.course.CopyCourseInput
	8,  // 20: eolymp.course.CourseService.CreateCourse:output_type -> eolymp.course.CreateCourseOutput
	10, // 21: eolymp.course.CourseService.UpdateCourse:output_type -> eolymp.course.UpdateCourseOutput
	12, // 22: eolymp.course.CourseService.DeleteCourse:output_type -> eolymp.course.DeleteCourseOutput
	6,  // 23: eolymp.course.CourseService.DescribeCourse:output_type -> eolymp.course.DescribeCourseOutput
	4,  // 24: eolymp.course.CourseService.ListCourses:output_type -> eolymp.course.ListCoursesOutput
	14, // 25: eolymp.course.CourseService.CopyCourse:output_type -> eolymp.course.CopyCourseOutput
	20, // [20:26] is the sub-list for method output_type
	14, // [14:20] is the sub-list for method input_type
	14, // [14:14] is the sub-list for extension type_name
	14, // [14:14] is the sub-list for extension extendee
	0,  // [0:14] is the sub-list for field type_name
}

func init() { file_eolymp_course_course_service_proto_init() }
func file_eolymp_course_course_service_proto_init() {
	if File_eolymp_course_course_service_proto != nil {
		return
	}
	file_eolymp_course_course_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_eolymp_course_course_service_proto_rawDesc), len(file_eolymp_course_course_service_proto_rawDesc)),
			NumEnums:      3,
			NumMessages:   13,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_eolymp_course_course_service_proto_goTypes,
		DependencyIndexes: file_eolymp_course_course_service_proto_depIdxs,
		EnumInfos:         file_eolymp_course_course_service_proto_enumTypes,
		MessageInfos:      file_eolymp_course_course_service_proto_msgTypes,
	}.Build()
	File_eolymp_course_course_service_proto = out.File
	file_eolymp_course_course_service_proto_goTypes = nil
	file_eolymp_course_course_service_proto_depIdxs = nil
}
