// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v3.21.12
// source: eolymp/course/submission_service.proto

package course

import (
	_ "github.com/eolymp/go-sdk/eolymp/annotations"
	atlas "github.com/eolymp/go-sdk/eolymp/atlas"
	wellknown "github.com/eolymp/go-sdk/eolymp/wellknown"
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

type CreateSubmissionInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MaterialId string `protobuf:"bytes,1,opt,name=material_id,json=materialId,proto3" json:"material_id,omitempty"`
	Runtime    string `protobuf:"bytes,2,opt,name=runtime,proto3" json:"runtime,omitempty"`
	Source     string `protobuf:"bytes,3,opt,name=source,proto3" json:"source,omitempty"`
}

func (x *CreateSubmissionInput) Reset() {
	*x = CreateSubmissionInput{}
	mi := &file_eolymp_course_submission_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateSubmissionInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateSubmissionInput) ProtoMessage() {}

func (x *CreateSubmissionInput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_course_submission_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateSubmissionInput.ProtoReflect.Descriptor instead.
func (*CreateSubmissionInput) Descriptor() ([]byte, []int) {
	return file_eolymp_course_submission_service_proto_rawDescGZIP(), []int{0}
}

func (x *CreateSubmissionInput) GetMaterialId() string {
	if x != nil {
		return x.MaterialId
	}
	return ""
}

func (x *CreateSubmissionInput) GetRuntime() string {
	if x != nil {
		return x.Runtime
	}
	return ""
}

func (x *CreateSubmissionInput) GetSource() string {
	if x != nil {
		return x.Source
	}
	return ""
}

type CreateSubmissionOutput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SubmissionId string `protobuf:"bytes,1,opt,name=submission_id,json=submissionId,proto3" json:"submission_id,omitempty"`
}

func (x *CreateSubmissionOutput) Reset() {
	*x = CreateSubmissionOutput{}
	mi := &file_eolymp_course_submission_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateSubmissionOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateSubmissionOutput) ProtoMessage() {}

func (x *CreateSubmissionOutput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_course_submission_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateSubmissionOutput.ProtoReflect.Descriptor instead.
func (*CreateSubmissionOutput) Descriptor() ([]byte, []int) {
	return file_eolymp_course_submission_service_proto_rawDescGZIP(), []int{1}
}

func (x *CreateSubmissionOutput) GetSubmissionId() string {
	if x != nil {
		return x.SubmissionId
	}
	return ""
}

type ListSubmissionsInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// pagination
	After string `protobuf:"bytes,12,opt,name=after,proto3" json:"after,omitempty"`
	Size  int32  `protobuf:"varint,11,opt,name=size,proto3" json:"size,omitempty"`
	// data filters
	Filters *ListSubmissionsInput_Filter `protobuf:"bytes,40,opt,name=filters,proto3" json:"filters,omitempty"`
}

func (x *ListSubmissionsInput) Reset() {
	*x = ListSubmissionsInput{}
	mi := &file_eolymp_course_submission_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListSubmissionsInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListSubmissionsInput) ProtoMessage() {}

func (x *ListSubmissionsInput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_course_submission_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListSubmissionsInput.ProtoReflect.Descriptor instead.
func (*ListSubmissionsInput) Descriptor() ([]byte, []int) {
	return file_eolymp_course_submission_service_proto_rawDescGZIP(), []int{2}
}

func (x *ListSubmissionsInput) GetAfter() string {
	if x != nil {
		return x.After
	}
	return ""
}

func (x *ListSubmissionsInput) GetSize() int32 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *ListSubmissionsInput) GetFilters() *ListSubmissionsInput_Filter {
	if x != nil {
		return x.Filters
	}
	return nil
}

type ListSubmissionsOutput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Total          int32               `protobuf:"varint,1,opt,name=total,proto3" json:"total,omitempty"`
	Items          []*atlas.Submission `protobuf:"bytes,2,rep,name=items,proto3" json:"items,omitempty"`
	NextPageCursor string              `protobuf:"bytes,4,opt,name=next_page_cursor,json=nextPageCursor,proto3" json:"next_page_cursor,omitempty"`
}

func (x *ListSubmissionsOutput) Reset() {
	*x = ListSubmissionsOutput{}
	mi := &file_eolymp_course_submission_service_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListSubmissionsOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListSubmissionsOutput) ProtoMessage() {}

func (x *ListSubmissionsOutput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_course_submission_service_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListSubmissionsOutput.ProtoReflect.Descriptor instead.
func (*ListSubmissionsOutput) Descriptor() ([]byte, []int) {
	return file_eolymp_course_submission_service_proto_rawDescGZIP(), []int{3}
}

func (x *ListSubmissionsOutput) GetTotal() int32 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *ListSubmissionsOutput) GetItems() []*atlas.Submission {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *ListSubmissionsOutput) GetNextPageCursor() string {
	if x != nil {
		return x.NextPageCursor
	}
	return ""
}

type DescribeSubmissionInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SubmissionId string `protobuf:"bytes,1,opt,name=submission_id,json=submissionId,proto3" json:"submission_id,omitempty"`
}

func (x *DescribeSubmissionInput) Reset() {
	*x = DescribeSubmissionInput{}
	mi := &file_eolymp_course_submission_service_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DescribeSubmissionInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescribeSubmissionInput) ProtoMessage() {}

func (x *DescribeSubmissionInput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_course_submission_service_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DescribeSubmissionInput.ProtoReflect.Descriptor instead.
func (*DescribeSubmissionInput) Descriptor() ([]byte, []int) {
	return file_eolymp_course_submission_service_proto_rawDescGZIP(), []int{4}
}

func (x *DescribeSubmissionInput) GetSubmissionId() string {
	if x != nil {
		return x.SubmissionId
	}
	return ""
}

type DescribeSubmissionOutput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Submission *atlas.Submission `protobuf:"bytes,1,opt,name=submission,proto3" json:"submission,omitempty"`
}

func (x *DescribeSubmissionOutput) Reset() {
	*x = DescribeSubmissionOutput{}
	mi := &file_eolymp_course_submission_service_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DescribeSubmissionOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescribeSubmissionOutput) ProtoMessage() {}

func (x *DescribeSubmissionOutput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_course_submission_service_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DescribeSubmissionOutput.ProtoReflect.Descriptor instead.
func (*DescribeSubmissionOutput) Descriptor() ([]byte, []int) {
	return file_eolymp_course_submission_service_proto_rawDescGZIP(), []int{5}
}

func (x *DescribeSubmissionOutput) GetSubmission() *atlas.Submission {
	if x != nil {
		return x.Submission
	}
	return nil
}

type WatchSubmissionInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SubmissionId string `protobuf:"bytes,1,opt,name=submission_id,json=submissionId,proto3" json:"submission_id,omitempty"`
}

func (x *WatchSubmissionInput) Reset() {
	*x = WatchSubmissionInput{}
	mi := &file_eolymp_course_submission_service_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *WatchSubmissionInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WatchSubmissionInput) ProtoMessage() {}

func (x *WatchSubmissionInput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_course_submission_service_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WatchSubmissionInput.ProtoReflect.Descriptor instead.
func (*WatchSubmissionInput) Descriptor() ([]byte, []int) {
	return file_eolymp_course_submission_service_proto_rawDescGZIP(), []int{6}
}

func (x *WatchSubmissionInput) GetSubmissionId() string {
	if x != nil {
		return x.SubmissionId
	}
	return ""
}

type WatchSubmissionOutput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Submission *atlas.Submission `protobuf:"bytes,1,opt,name=submission,proto3" json:"submission,omitempty"`
}

func (x *WatchSubmissionOutput) Reset() {
	*x = WatchSubmissionOutput{}
	mi := &file_eolymp_course_submission_service_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *WatchSubmissionOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WatchSubmissionOutput) ProtoMessage() {}

func (x *WatchSubmissionOutput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_course_submission_service_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WatchSubmissionOutput.ProtoReflect.Descriptor instead.
func (*WatchSubmissionOutput) Descriptor() ([]byte, []int) {
	return file_eolymp_course_submission_service_proto_rawDescGZIP(), []int{7}
}

func (x *WatchSubmissionOutput) GetSubmission() *atlas.Submission {
	if x != nil {
		return x.Submission
	}
	return nil
}

type ListSubmissionsInput_Filter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          []*wellknown.ExpressionID        `protobuf:"bytes,1,rep,name=id,proto3" json:"id,omitempty"`
	MemberId    []*wellknown.ExpressionID        `protobuf:"bytes,9,rep,name=member_id,json=memberId,proto3" json:"member_id,omitempty"`
	MaterialId  []*wellknown.ExpressionID        `protobuf:"bytes,11,rep,name=material_id,json=materialId,proto3" json:"material_id,omitempty"`
	SubmittedAt []*wellknown.ExpressionTimestamp `protobuf:"bytes,4,rep,name=submitted_at,json=submittedAt,proto3" json:"submitted_at,omitempty"`
	Runtime     []*wellknown.ExpressionEnum      `protobuf:"bytes,5,rep,name=runtime,proto3" json:"runtime,omitempty"`
	Status      []*wellknown.ExpressionEnum      `protobuf:"bytes,6,rep,name=status,proto3" json:"status,omitempty"`
	Verdict     []*wellknown.ExpressionEnum      `protobuf:"bytes,10,rep,name=verdict,proto3" json:"verdict,omitempty"`
	Score       []*wellknown.ExpressionFloat     `protobuf:"bytes,7,rep,name=score,proto3" json:"score,omitempty"`
	Percentage  []*wellknown.ExpressionFloat     `protobuf:"bytes,8,rep,name=percentage,proto3" json:"percentage,omitempty"`
}

func (x *ListSubmissionsInput_Filter) Reset() {
	*x = ListSubmissionsInput_Filter{}
	mi := &file_eolymp_course_submission_service_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListSubmissionsInput_Filter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListSubmissionsInput_Filter) ProtoMessage() {}

func (x *ListSubmissionsInput_Filter) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_course_submission_service_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListSubmissionsInput_Filter.ProtoReflect.Descriptor instead.
func (*ListSubmissionsInput_Filter) Descriptor() ([]byte, []int) {
	return file_eolymp_course_submission_service_proto_rawDescGZIP(), []int{2, 0}
}

func (x *ListSubmissionsInput_Filter) GetId() []*wellknown.ExpressionID {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *ListSubmissionsInput_Filter) GetMemberId() []*wellknown.ExpressionID {
	if x != nil {
		return x.MemberId
	}
	return nil
}

func (x *ListSubmissionsInput_Filter) GetMaterialId() []*wellknown.ExpressionID {
	if x != nil {
		return x.MaterialId
	}
	return nil
}

func (x *ListSubmissionsInput_Filter) GetSubmittedAt() []*wellknown.ExpressionTimestamp {
	if x != nil {
		return x.SubmittedAt
	}
	return nil
}

func (x *ListSubmissionsInput_Filter) GetRuntime() []*wellknown.ExpressionEnum {
	if x != nil {
		return x.Runtime
	}
	return nil
}

func (x *ListSubmissionsInput_Filter) GetStatus() []*wellknown.ExpressionEnum {
	if x != nil {
		return x.Status
	}
	return nil
}

func (x *ListSubmissionsInput_Filter) GetVerdict() []*wellknown.ExpressionEnum {
	if x != nil {
		return x.Verdict
	}
	return nil
}

func (x *ListSubmissionsInput_Filter) GetScore() []*wellknown.ExpressionFloat {
	if x != nil {
		return x.Score
	}
	return nil
}

func (x *ListSubmissionsInput_Filter) GetPercentage() []*wellknown.ExpressionFloat {
	if x != nil {
		return x.Percentage
	}
	return nil
}

var File_eolymp_course_submission_service_proto protoreflect.FileDescriptor

var file_eolymp_course_submission_service_proto_rawDesc = []byte{
	0x0a, 0x26, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x2f,
	0x73, 0x75, 0x62, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70,
	0x2e, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x1a, 0x1d, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x68, 0x74, 0x74, 0x70,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x22, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x61,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x72, 0x61, 0x74, 0x65, 0x6c,
	0x69, 0x6d, 0x69, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x65, 0x6f, 0x6c, 0x79,
	0x6d, 0x70, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73,
	0x63, 0x6f, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x65, 0x6f, 0x6c, 0x79,
	0x6d, 0x70, 0x2f, 0x61, 0x74, 0x6c, 0x61, 0x73, 0x2f, 0x73, 0x75, 0x62, 0x6d, 0x69, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x21, 0x65, 0x6f, 0x6c, 0x79, 0x6d,
	0x70, 0x2f, 0x77, 0x65, 0x6c, 0x6c, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x2f, 0x65, 0x78, 0x70, 0x72,
	0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x6a, 0x0a, 0x15,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x6d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61,
	0x6c, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6d, 0x61, 0x74, 0x65,
	0x72, 0x69, 0x61, 0x6c, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x22, 0x3d, 0x0a, 0x16, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x4f, 0x75, 0x74, 0x70,
	0x75, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x73, 0x75, 0x62, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x73, 0x75, 0x62, 0x6d, 0x69,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x22, 0xb7, 0x05, 0x0a, 0x14, 0x4c, 0x69, 0x73, 0x74,
	0x53, 0x75, 0x62, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x49, 0x6e, 0x70, 0x75, 0x74,
	0x12, 0x14, 0x0a, 0x05, 0x61, 0x66, 0x74, 0x65, 0x72, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x61, 0x66, 0x74, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x0b,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x44, 0x0a, 0x07, 0x66, 0x69,
	0x6c, 0x74, 0x65, 0x72, 0x73, 0x18, 0x28, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x65, 0x6f,
	0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x2e, 0x4c, 0x69, 0x73, 0x74,
	0x53, 0x75, 0x62, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x49, 0x6e, 0x70, 0x75, 0x74,
	0x2e, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x52, 0x07, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73,
	0x1a, 0xae, 0x04, 0x0a, 0x06, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x2e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70,
	0x2e, 0x77, 0x65, 0x6c, 0x6c, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x2e, 0x45, 0x78, 0x70, 0x72, 0x65,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x52, 0x02, 0x69, 0x64, 0x12, 0x3b, 0x0a, 0x09, 0x6d,
	0x65, 0x6d, 0x62, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x09, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e,
	0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x77, 0x65, 0x6c, 0x6c, 0x6b, 0x6e, 0x6f, 0x77,
	0x6e, 0x2e, 0x45, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x52, 0x08,
	0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x49, 0x64, 0x12, 0x3f, 0x0a, 0x0b, 0x6d, 0x61, 0x74, 0x65,
	0x72, 0x69, 0x61, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x0b, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e,
	0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x77, 0x65, 0x6c, 0x6c, 0x6b, 0x6e, 0x6f, 0x77, 0x6e,
	0x2e, 0x45, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x52, 0x0a, 0x6d,
	0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x49, 0x64, 0x12, 0x48, 0x0a, 0x0c, 0x73, 0x75, 0x62,
	0x6d, 0x69, 0x74, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x25, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x77, 0x65, 0x6c, 0x6c, 0x6b, 0x6e, 0x6f,
	0x77, 0x6e, 0x2e, 0x45, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0b, 0x73, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x12, 0x3a, 0x0a, 0x07, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x05,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x77, 0x65,
	0x6c, 0x6c, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x2e, 0x45, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x45, 0x6e, 0x75, 0x6d, 0x52, 0x07, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x12,
	0x38, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x20, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x77, 0x65, 0x6c, 0x6c, 0x6b, 0x6e, 0x6f,
	0x77, 0x6e, 0x2e, 0x45, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x45, 0x6e, 0x75,
	0x6d, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x3a, 0x0a, 0x07, 0x76, 0x65, 0x72,
	0x64, 0x69, 0x63, 0x74, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x65, 0x6f, 0x6c,
	0x79, 0x6d, 0x70, 0x2e, 0x77, 0x65, 0x6c, 0x6c, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x2e, 0x45, 0x78,
	0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x45, 0x6e, 0x75, 0x6d, 0x52, 0x07, 0x76, 0x65,
	0x72, 0x64, 0x69, 0x63, 0x74, 0x12, 0x37, 0x0a, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x07,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x77, 0x65,
	0x6c, 0x6c, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x2e, 0x45, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x46, 0x6c, 0x6f, 0x61, 0x74, 0x52, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x12, 0x41,
	0x0a, 0x0a, 0x70, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x61, 0x67, 0x65, 0x18, 0x08, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x21, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x77, 0x65, 0x6c, 0x6c,
	0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x2e, 0x45, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x46, 0x6c, 0x6f, 0x61, 0x74, 0x52, 0x0a, 0x70, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x61, 0x67,
	0x65, 0x22, 0x87, 0x01, 0x0a, 0x15, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x73, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74,
	0x6f, 0x74, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61,
	0x6c, 0x12, 0x2e, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x18, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x61, 0x74, 0x6c, 0x61, 0x73, 0x2e,
	0x53, 0x75, 0x62, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d,
	0x73, 0x12, 0x28, 0x0a, 0x10, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x63,
	0x75, 0x72, 0x73, 0x6f, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x6e, 0x65, 0x78,
	0x74, 0x50, 0x61, 0x67, 0x65, 0x43, 0x75, 0x72, 0x73, 0x6f, 0x72, 0x22, 0x3e, 0x0a, 0x17, 0x44,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x73, 0x75, 0x62, 0x6d, 0x69, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x73,
	0x75, 0x62, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x22, 0x54, 0x0a, 0x18, 0x44,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x12, 0x38, 0x0a, 0x0a, 0x73, 0x75, 0x62, 0x6d, 0x69,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x65, 0x6f,
	0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x61, 0x74, 0x6c, 0x61, 0x73, 0x2e, 0x53, 0x75, 0x62, 0x6d, 0x69,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x73, 0x75, 0x62, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x22, 0x3b, 0x0a, 0x14, 0x57, 0x61, 0x74, 0x63, 0x68, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x73, 0x75, 0x62,
	0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0c, 0x73, 0x75, 0x62, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x22, 0x51,
	0x0a, 0x15, 0x57, 0x61, 0x74, 0x63, 0x68, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x12, 0x38, 0x0a, 0x0a, 0x73, 0x75, 0x62, 0x6d, 0x69,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x65, 0x6f,
	0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x61, 0x74, 0x6c, 0x61, 0x73, 0x2e, 0x53, 0x75, 0x62, 0x6d, 0x69,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x73, 0x75, 0x62, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x32, 0x87, 0x05, 0x0a, 0x11, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0xa0, 0x01, 0x0a, 0x10, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x24, 0x2e, 0x65,
	0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x70,
	0x75, 0x74, 0x1a, 0x25, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x63, 0x6f, 0x75, 0x72,
	0x73, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x22, 0x3f, 0xea, 0xe2, 0x0a, 0x0b, 0xf5,
	0xe2, 0x0a, 0x0a, 0xd7, 0x23, 0x3e, 0xf8, 0xe2, 0x0a, 0x0a, 0x82, 0xe3, 0x0a, 0x18, 0x8a, 0xe3,
	0x0a, 0x14, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x3a, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x3a,
	0x73, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0e, 0x22, 0x0c, 0x2f, 0x73,
	0x75, 0x62, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x9b, 0x01, 0x0a, 0x0f, 0x4c,
	0x69, 0x73, 0x74, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x23,
	0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x49, 0x6e,
	0x70, 0x75, 0x74, 0x1a, 0x24, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x63, 0x6f, 0x75,
	0x72, 0x73, 0x65, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x73, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x22, 0x3d, 0xea, 0xe2, 0x0a, 0x0b, 0xf5,
	0xe2, 0x0a, 0x00, 0x00, 0x80, 0x3f, 0xf8, 0xe2, 0x0a, 0x05, 0x82, 0xe3, 0x0a, 0x16, 0x8a, 0xe3,
	0x0a, 0x12, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x3a, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x3a,
	0x72, 0x65, 0x61, 0x64, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0e, 0x12, 0x0c, 0x2f, 0x73, 0x75, 0x62,
	0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0xb4, 0x01, 0x0a, 0x12, 0x44, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x62, 0x65, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12,
	0x26, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x2e,
	0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x27, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70,
	0x2e, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x2e, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65,
	0x53, 0x75, 0x62, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74,
	0x22, 0x4d, 0xea, 0xe2, 0x0a, 0x0b, 0xf5, 0xe2, 0x0a, 0x00, 0x00, 0x00, 0x40, 0xf8, 0xe2, 0x0a,
	0x0a, 0x82, 0xe3, 0x0a, 0x16, 0x8a, 0xe3, 0x0a, 0x12, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x3a,
	0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x3a, 0x72, 0x65, 0x61, 0x64, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x1e, 0x12, 0x1c, 0x2f, 0x73, 0x75, 0x62, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2f,
	0x7b, 0x73, 0x75, 0x62, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x7d, 0x12,
	0x7a, 0x0a, 0x0f, 0x57, 0x61, 0x74, 0x63, 0x68, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x12, 0x23, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x63, 0x6f, 0x75, 0x72,
	0x73, 0x65, 0x2e, 0x57, 0x61, 0x74, 0x63, 0x68, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x24, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70,
	0x2e, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x2e, 0x57, 0x61, 0x74, 0x63, 0x68, 0x53, 0x75, 0x62,
	0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x22, 0x1a, 0x82,
	0xe3, 0x0a, 0x16, 0x8a, 0xe3, 0x0a, 0x12, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x3a, 0x63, 0x6f,
	0x75, 0x72, 0x73, 0x65, 0x3a, 0x72, 0x65, 0x61, 0x64, 0x30, 0x01, 0x42, 0x2f, 0x5a, 0x2d, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70,
	0x2f, 0x67, 0x6f, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x63,
	0x6f, 0x75, 0x72, 0x73, 0x65, 0x3b, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_eolymp_course_submission_service_proto_rawDescOnce sync.Once
	file_eolymp_course_submission_service_proto_rawDescData = file_eolymp_course_submission_service_proto_rawDesc
)

func file_eolymp_course_submission_service_proto_rawDescGZIP() []byte {
	file_eolymp_course_submission_service_proto_rawDescOnce.Do(func() {
		file_eolymp_course_submission_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_eolymp_course_submission_service_proto_rawDescData)
	})
	return file_eolymp_course_submission_service_proto_rawDescData
}

var file_eolymp_course_submission_service_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_eolymp_course_submission_service_proto_goTypes = []any{
	(*CreateSubmissionInput)(nil),         // 0: eolymp.course.CreateSubmissionInput
	(*CreateSubmissionOutput)(nil),        // 1: eolymp.course.CreateSubmissionOutput
	(*ListSubmissionsInput)(nil),          // 2: eolymp.course.ListSubmissionsInput
	(*ListSubmissionsOutput)(nil),         // 3: eolymp.course.ListSubmissionsOutput
	(*DescribeSubmissionInput)(nil),       // 4: eolymp.course.DescribeSubmissionInput
	(*DescribeSubmissionOutput)(nil),      // 5: eolymp.course.DescribeSubmissionOutput
	(*WatchSubmissionInput)(nil),          // 6: eolymp.course.WatchSubmissionInput
	(*WatchSubmissionOutput)(nil),         // 7: eolymp.course.WatchSubmissionOutput
	(*ListSubmissionsInput_Filter)(nil),   // 8: eolymp.course.ListSubmissionsInput.Filter
	(*atlas.Submission)(nil),              // 9: eolymp.atlas.Submission
	(*wellknown.ExpressionID)(nil),        // 10: eolymp.wellknown.ExpressionID
	(*wellknown.ExpressionTimestamp)(nil), // 11: eolymp.wellknown.ExpressionTimestamp
	(*wellknown.ExpressionEnum)(nil),      // 12: eolymp.wellknown.ExpressionEnum
	(*wellknown.ExpressionFloat)(nil),     // 13: eolymp.wellknown.ExpressionFloat
}
var file_eolymp_course_submission_service_proto_depIdxs = []int32{
	8,  // 0: eolymp.course.ListSubmissionsInput.filters:type_name -> eolymp.course.ListSubmissionsInput.Filter
	9,  // 1: eolymp.course.ListSubmissionsOutput.items:type_name -> eolymp.atlas.Submission
	9,  // 2: eolymp.course.DescribeSubmissionOutput.submission:type_name -> eolymp.atlas.Submission
	9,  // 3: eolymp.course.WatchSubmissionOutput.submission:type_name -> eolymp.atlas.Submission
	10, // 4: eolymp.course.ListSubmissionsInput.Filter.id:type_name -> eolymp.wellknown.ExpressionID
	10, // 5: eolymp.course.ListSubmissionsInput.Filter.member_id:type_name -> eolymp.wellknown.ExpressionID
	10, // 6: eolymp.course.ListSubmissionsInput.Filter.material_id:type_name -> eolymp.wellknown.ExpressionID
	11, // 7: eolymp.course.ListSubmissionsInput.Filter.submitted_at:type_name -> eolymp.wellknown.ExpressionTimestamp
	12, // 8: eolymp.course.ListSubmissionsInput.Filter.runtime:type_name -> eolymp.wellknown.ExpressionEnum
	12, // 9: eolymp.course.ListSubmissionsInput.Filter.status:type_name -> eolymp.wellknown.ExpressionEnum
	12, // 10: eolymp.course.ListSubmissionsInput.Filter.verdict:type_name -> eolymp.wellknown.ExpressionEnum
	13, // 11: eolymp.course.ListSubmissionsInput.Filter.score:type_name -> eolymp.wellknown.ExpressionFloat
	13, // 12: eolymp.course.ListSubmissionsInput.Filter.percentage:type_name -> eolymp.wellknown.ExpressionFloat
	0,  // 13: eolymp.course.SubmissionService.CreateSubmission:input_type -> eolymp.course.CreateSubmissionInput
	2,  // 14: eolymp.course.SubmissionService.ListSubmissions:input_type -> eolymp.course.ListSubmissionsInput
	4,  // 15: eolymp.course.SubmissionService.DescribeSubmission:input_type -> eolymp.course.DescribeSubmissionInput
	6,  // 16: eolymp.course.SubmissionService.WatchSubmission:input_type -> eolymp.course.WatchSubmissionInput
	1,  // 17: eolymp.course.SubmissionService.CreateSubmission:output_type -> eolymp.course.CreateSubmissionOutput
	3,  // 18: eolymp.course.SubmissionService.ListSubmissions:output_type -> eolymp.course.ListSubmissionsOutput
	5,  // 19: eolymp.course.SubmissionService.DescribeSubmission:output_type -> eolymp.course.DescribeSubmissionOutput
	7,  // 20: eolymp.course.SubmissionService.WatchSubmission:output_type -> eolymp.course.WatchSubmissionOutput
	17, // [17:21] is the sub-list for method output_type
	13, // [13:17] is the sub-list for method input_type
	13, // [13:13] is the sub-list for extension type_name
	13, // [13:13] is the sub-list for extension extendee
	0,  // [0:13] is the sub-list for field type_name
}

func init() { file_eolymp_course_submission_service_proto_init() }
func file_eolymp_course_submission_service_proto_init() {
	if File_eolymp_course_submission_service_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_eolymp_course_submission_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_eolymp_course_submission_service_proto_goTypes,
		DependencyIndexes: file_eolymp_course_submission_service_proto_depIdxs,
		MessageInfos:      file_eolymp_course_submission_service_proto_msgTypes,
	}.Build()
	File_eolymp_course_submission_service_proto = out.File
	file_eolymp_course_submission_service_proto_rawDesc = nil
	file_eolymp_course_submission_service_proto_goTypes = nil
	file_eolymp_course_submission_service_proto_depIdxs = nil
}
