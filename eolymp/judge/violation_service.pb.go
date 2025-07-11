// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.29.4
// source: eolymp/judge/violation_service.proto

package judge

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

type CreateViolationInput struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Violation     *Violation             `protobuf:"bytes,1,opt,name=violation,proto3" json:"violation,omitempty"`
	DontNotify    bool                   `protobuf:"varint,2,opt,name=dont_notify,json=dontNotify,proto3" json:"dont_notify,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateViolationInput) Reset() {
	*x = CreateViolationInput{}
	mi := &file_eolymp_judge_violation_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateViolationInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateViolationInput) ProtoMessage() {}

func (x *CreateViolationInput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_judge_violation_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateViolationInput.ProtoReflect.Descriptor instead.
func (*CreateViolationInput) Descriptor() ([]byte, []int) {
	return file_eolymp_judge_violation_service_proto_rawDescGZIP(), []int{0}
}

func (x *CreateViolationInput) GetViolation() *Violation {
	if x != nil {
		return x.Violation
	}
	return nil
}

func (x *CreateViolationInput) GetDontNotify() bool {
	if x != nil {
		return x.DontNotify
	}
	return false
}

type CreateViolationOutput struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ViolationId   string                 `protobuf:"bytes,1,opt,name=violation_id,json=violationId,proto3" json:"violation_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateViolationOutput) Reset() {
	*x = CreateViolationOutput{}
	mi := &file_eolymp_judge_violation_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateViolationOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateViolationOutput) ProtoMessage() {}

func (x *CreateViolationOutput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_judge_violation_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateViolationOutput.ProtoReflect.Descriptor instead.
func (*CreateViolationOutput) Descriptor() ([]byte, []int) {
	return file_eolymp_judge_violation_service_proto_rawDescGZIP(), []int{1}
}

func (x *CreateViolationOutput) GetViolationId() string {
	if x != nil {
		return x.ViolationId
	}
	return ""
}

type UpdateViolationInput struct {
	state         protoimpl.MessageState  `protogen:"open.v1"`
	Patch         []Violation_Patch_Field `protobuf:"varint,3,rep,packed,name=patch,proto3,enum=eolymp.judge.Violation_Patch_Field" json:"patch,omitempty"`
	ViolationId   string                  `protobuf:"bytes,1,opt,name=violation_id,json=violationId,proto3" json:"violation_id,omitempty"`
	Violation     *Violation              `protobuf:"bytes,2,opt,name=violation,proto3" json:"violation,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateViolationInput) Reset() {
	*x = UpdateViolationInput{}
	mi := &file_eolymp_judge_violation_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateViolationInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateViolationInput) ProtoMessage() {}

func (x *UpdateViolationInput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_judge_violation_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateViolationInput.ProtoReflect.Descriptor instead.
func (*UpdateViolationInput) Descriptor() ([]byte, []int) {
	return file_eolymp_judge_violation_service_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateViolationInput) GetPatch() []Violation_Patch_Field {
	if x != nil {
		return x.Patch
	}
	return nil
}

func (x *UpdateViolationInput) GetViolationId() string {
	if x != nil {
		return x.ViolationId
	}
	return ""
}

func (x *UpdateViolationInput) GetViolation() *Violation {
	if x != nil {
		return x.Violation
	}
	return nil
}

type UpdateViolationOutput struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateViolationOutput) Reset() {
	*x = UpdateViolationOutput{}
	mi := &file_eolymp_judge_violation_service_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateViolationOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateViolationOutput) ProtoMessage() {}

func (x *UpdateViolationOutput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_judge_violation_service_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateViolationOutput.ProtoReflect.Descriptor instead.
func (*UpdateViolationOutput) Descriptor() ([]byte, []int) {
	return file_eolymp_judge_violation_service_proto_rawDescGZIP(), []int{3}
}

type DeleteViolationInput struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ViolationId   string                 `protobuf:"bytes,1,opt,name=violation_id,json=violationId,proto3" json:"violation_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteViolationInput) Reset() {
	*x = DeleteViolationInput{}
	mi := &file_eolymp_judge_violation_service_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteViolationInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteViolationInput) ProtoMessage() {}

func (x *DeleteViolationInput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_judge_violation_service_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteViolationInput.ProtoReflect.Descriptor instead.
func (*DeleteViolationInput) Descriptor() ([]byte, []int) {
	return file_eolymp_judge_violation_service_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteViolationInput) GetViolationId() string {
	if x != nil {
		return x.ViolationId
	}
	return ""
}

type DeleteViolationOutput struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteViolationOutput) Reset() {
	*x = DeleteViolationOutput{}
	mi := &file_eolymp_judge_violation_service_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteViolationOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteViolationOutput) ProtoMessage() {}

func (x *DeleteViolationOutput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_judge_violation_service_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteViolationOutput.ProtoReflect.Descriptor instead.
func (*DeleteViolationOutput) Descriptor() ([]byte, []int) {
	return file_eolymp_judge_violation_service_proto_rawDescGZIP(), []int{5}
}

type DescribeViolationInput struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ViolationId   string                 `protobuf:"bytes,1,opt,name=violation_id,json=violationId,proto3" json:"violation_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DescribeViolationInput) Reset() {
	*x = DescribeViolationInput{}
	mi := &file_eolymp_judge_violation_service_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DescribeViolationInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescribeViolationInput) ProtoMessage() {}

func (x *DescribeViolationInput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_judge_violation_service_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DescribeViolationInput.ProtoReflect.Descriptor instead.
func (*DescribeViolationInput) Descriptor() ([]byte, []int) {
	return file_eolymp_judge_violation_service_proto_rawDescGZIP(), []int{6}
}

func (x *DescribeViolationInput) GetViolationId() string {
	if x != nil {
		return x.ViolationId
	}
	return ""
}

type DescribeViolationOutput struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Violation     *Violation             `protobuf:"bytes,1,opt,name=violation,proto3" json:"violation,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DescribeViolationOutput) Reset() {
	*x = DescribeViolationOutput{}
	mi := &file_eolymp_judge_violation_service_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DescribeViolationOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescribeViolationOutput) ProtoMessage() {}

func (x *DescribeViolationOutput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_judge_violation_service_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DescribeViolationOutput.ProtoReflect.Descriptor instead.
func (*DescribeViolationOutput) Descriptor() ([]byte, []int) {
	return file_eolymp_judge_violation_service_proto_rawDescGZIP(), []int{7}
}

func (x *DescribeViolationOutput) GetViolation() *Violation {
	if x != nil {
		return x.Violation
	}
	return nil
}

type ListViolationsInput struct {
	state         protoimpl.MessageState      `protogen:"open.v1"`
	Offset        int32                       `protobuf:"varint,10,opt,name=offset,proto3" json:"offset,omitempty"`
	Size          int32                       `protobuf:"varint,11,opt,name=size,proto3" json:"size,omitempty"`
	Filters       *ListViolationsInput_Filter `protobuf:"bytes,40,opt,name=filters,proto3" json:"filters,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListViolationsInput) Reset() {
	*x = ListViolationsInput{}
	mi := &file_eolymp_judge_violation_service_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListViolationsInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListViolationsInput) ProtoMessage() {}

func (x *ListViolationsInput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_judge_violation_service_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListViolationsInput.ProtoReflect.Descriptor instead.
func (*ListViolationsInput) Descriptor() ([]byte, []int) {
	return file_eolymp_judge_violation_service_proto_rawDescGZIP(), []int{8}
}

func (x *ListViolationsInput) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *ListViolationsInput) GetSize() int32 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *ListViolationsInput) GetFilters() *ListViolationsInput_Filter {
	if x != nil {
		return x.Filters
	}
	return nil
}

type ListViolationsOutput struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Total         int32                  `protobuf:"varint,1,opt,name=total,proto3" json:"total,omitempty"`
	Items         []*Violation           `protobuf:"bytes,2,rep,name=items,proto3" json:"items,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListViolationsOutput) Reset() {
	*x = ListViolationsOutput{}
	mi := &file_eolymp_judge_violation_service_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListViolationsOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListViolationsOutput) ProtoMessage() {}

func (x *ListViolationsOutput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_judge_violation_service_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListViolationsOutput.ProtoReflect.Descriptor instead.
func (*ListViolationsOutput) Descriptor() ([]byte, []int) {
	return file_eolymp_judge_violation_service_proto_rawDescGZIP(), []int{9}
}

func (x *ListViolationsOutput) GetTotal() int32 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *ListViolationsOutput) GetItems() []*Violation {
	if x != nil {
		return x.Items
	}
	return nil
}

type ListViolationsInput_Filter struct {
	state         protoimpl.MessageState        `protogen:"open.v1"`
	Id            []*wellknown.ExpressionID     `protobuf:"bytes,2,rep,name=id,proto3" json:"id,omitempty"`
	Status        []*wellknown.ExpressionEnum   `protobuf:"bytes,3,rep,name=status,proto3" json:"status,omitempty"`
	Type          []*wellknown.ExpressionEnum   `protobuf:"bytes,5,rep,name=type,proto3" json:"type,omitempty"`
	Summary       []*wellknown.ExpressionString `protobuf:"bytes,4,rep,name=summary,proto3" json:"summary,omitempty"`
	Automatic     []*wellknown.ExpressionBool   `protobuf:"bytes,10,rep,name=automatic,proto3" json:"automatic,omitempty"`
	ParticipantId []*wellknown.ExpressionID     `protobuf:"bytes,6,rep,name=participant_id,json=participantId,proto3" json:"participant_id,omitempty"`
	SubmissionId  []*wellknown.ExpressionID     `protobuf:"bytes,8,rep,name=submission_id,json=submissionId,proto3" json:"submission_id,omitempty"`
	CreatedBy     []*wellknown.ExpressionID     `protobuf:"bytes,7,rep,name=created_by,json=createdBy,proto3" json:"created_by,omitempty"`
	ConfirmedBy   []*wellknown.ExpressionID     `protobuf:"bytes,9,rep,name=confirmed_by,json=confirmedBy,proto3" json:"confirmed_by,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListViolationsInput_Filter) Reset() {
	*x = ListViolationsInput_Filter{}
	mi := &file_eolymp_judge_violation_service_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListViolationsInput_Filter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListViolationsInput_Filter) ProtoMessage() {}

func (x *ListViolationsInput_Filter) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_judge_violation_service_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListViolationsInput_Filter.ProtoReflect.Descriptor instead.
func (*ListViolationsInput_Filter) Descriptor() ([]byte, []int) {
	return file_eolymp_judge_violation_service_proto_rawDescGZIP(), []int{8, 0}
}

func (x *ListViolationsInput_Filter) GetId() []*wellknown.ExpressionID {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *ListViolationsInput_Filter) GetStatus() []*wellknown.ExpressionEnum {
	if x != nil {
		return x.Status
	}
	return nil
}

func (x *ListViolationsInput_Filter) GetType() []*wellknown.ExpressionEnum {
	if x != nil {
		return x.Type
	}
	return nil
}

func (x *ListViolationsInput_Filter) GetSummary() []*wellknown.ExpressionString {
	if x != nil {
		return x.Summary
	}
	return nil
}

func (x *ListViolationsInput_Filter) GetAutomatic() []*wellknown.ExpressionBool {
	if x != nil {
		return x.Automatic
	}
	return nil
}

func (x *ListViolationsInput_Filter) GetParticipantId() []*wellknown.ExpressionID {
	if x != nil {
		return x.ParticipantId
	}
	return nil
}

func (x *ListViolationsInput_Filter) GetSubmissionId() []*wellknown.ExpressionID {
	if x != nil {
		return x.SubmissionId
	}
	return nil
}

func (x *ListViolationsInput_Filter) GetCreatedBy() []*wellknown.ExpressionID {
	if x != nil {
		return x.CreatedBy
	}
	return nil
}

func (x *ListViolationsInput_Filter) GetConfirmedBy() []*wellknown.ExpressionID {
	if x != nil {
		return x.ConfirmedBy
	}
	return nil
}

var File_eolymp_judge_violation_service_proto protoreflect.FileDescriptor

const file_eolymp_judge_violation_service_proto_rawDesc = "" +
	"\n" +
	"$eolymp/judge/violation_service.proto\x12\feolymp.judge\x1a\x1deolymp/annotations/http.proto\x1a\"eolymp/annotations/ratelimit.proto\x1a\x1eeolymp/annotations/scope.proto\x1a\x1ceolymp/judge/violation.proto\x1a!eolymp/wellknown/expression.proto\"n\n" +
	"\x14CreateViolationInput\x125\n" +
	"\tviolation\x18\x01 \x01(\v2\x17.eolymp.judge.ViolationR\tviolation\x12\x1f\n" +
	"\vdont_notify\x18\x02 \x01(\bR\n" +
	"dontNotify\":\n" +
	"\x15CreateViolationOutput\x12!\n" +
	"\fviolation_id\x18\x01 \x01(\tR\vviolationId\"\xab\x01\n" +
	"\x14UpdateViolationInput\x129\n" +
	"\x05patch\x18\x03 \x03(\x0e2#.eolymp.judge.Violation.Patch.FieldR\x05patch\x12!\n" +
	"\fviolation_id\x18\x01 \x01(\tR\vviolationId\x125\n" +
	"\tviolation\x18\x02 \x01(\v2\x17.eolymp.judge.ViolationR\tviolation\"\x17\n" +
	"\x15UpdateViolationOutput\"9\n" +
	"\x14DeleteViolationInput\x12!\n" +
	"\fviolation_id\x18\x01 \x01(\tR\vviolationId\"\x17\n" +
	"\x15DeleteViolationOutput\";\n" +
	"\x16DescribeViolationInput\x12!\n" +
	"\fviolation_id\x18\x01 \x01(\tR\vviolationId\"P\n" +
	"\x17DescribeViolationOutput\x125\n" +
	"\tviolation\x18\x01 \x01(\v2\x17.eolymp.judge.ViolationR\tviolation\"\xbc\x05\n" +
	"\x13ListViolationsInput\x12\x16\n" +
	"\x06offset\x18\n" +
	" \x01(\x05R\x06offset\x12\x12\n" +
	"\x04size\x18\v \x01(\x05R\x04size\x12B\n" +
	"\afilters\x18( \x01(\v2(.eolymp.judge.ListViolationsInput.FilterR\afilters\x1a\xb4\x04\n" +
	"\x06Filter\x12.\n" +
	"\x02id\x18\x02 \x03(\v2\x1e.eolymp.wellknown.ExpressionIDR\x02id\x128\n" +
	"\x06status\x18\x03 \x03(\v2 .eolymp.wellknown.ExpressionEnumR\x06status\x124\n" +
	"\x04type\x18\x05 \x03(\v2 .eolymp.wellknown.ExpressionEnumR\x04type\x12<\n" +
	"\asummary\x18\x04 \x03(\v2\".eolymp.wellknown.ExpressionStringR\asummary\x12>\n" +
	"\tautomatic\x18\n" +
	" \x03(\v2 .eolymp.wellknown.ExpressionBoolR\tautomatic\x12E\n" +
	"\x0eparticipant_id\x18\x06 \x03(\v2\x1e.eolymp.wellknown.ExpressionIDR\rparticipantId\x12C\n" +
	"\rsubmission_id\x18\b \x03(\v2\x1e.eolymp.wellknown.ExpressionIDR\fsubmissionId\x12=\n" +
	"\n" +
	"created_by\x18\a \x03(\v2\x1e.eolymp.wellknown.ExpressionIDR\tcreatedBy\x12A\n" +
	"\fconfirmed_by\x18\t \x03(\v2\x1e.eolymp.wellknown.ExpressionIDR\vconfirmedBy\"[\n" +
	"\x14ListViolationsOutput\x12\x14\n" +
	"\x05total\x18\x01 \x01(\x05R\x05total\x12-\n" +
	"\x05items\x18\x02 \x03(\v2\x17.eolymp.judge.ViolationR\x05items2\xcc\x06\n" +
	"\x10ViolationService\x12\x99\x01\n" +
	"\x0fCreateViolation\x12\".eolymp.judge.CreateViolationInput\x1a#.eolymp.judge.CreateViolationOutput\"=\xea\xe2\n" +
	"\v\xf5\xe2\n" +
	"\x00\x00 A\xf8\xe2\n" +
	"d\x82\xe3\n" +
	"\x17\x8a\xe3\n" +
	"\x13judge:contest:write\x82\xd3\xe4\x93\x02\r\"\v/violations\x12\xa8\x01\n" +
	"\x0fUpdateViolation\x12\".eolymp.judge.UpdateViolationInput\x1a#.eolymp.judge.UpdateViolationOutput\"L\xea\xe2\n" +
	"\v\xf5\xe2\n" +
	"\x00\x00\xa0@\xf8\xe2\n" +
	"\x14\x82\xe3\n" +
	"\x17\x8a\xe3\n" +
	"\x13judge:contest:write\x82\xd3\xe4\x93\x02\x1c\"\x1a/violations/{violation_id}\x12\xa8\x01\n" +
	"\x0fDeleteViolation\x12\".eolymp.judge.DeleteViolationInput\x1a#.eolymp.judge.DeleteViolationOutput\"L\xea\xe2\n" +
	"\v\xf5\xe2\n" +
	"\x00\x00 A\xf8\xe2\n" +
	"d\x82\xe3\n" +
	"\x17\x8a\xe3\n" +
	"\x13judge:contest:write\x82\xd3\xe4\x93\x02\x1c*\x1a/violations/{violation_id}\x12\xad\x01\n" +
	"\x11DescribeViolation\x12$.eolymp.judge.DescribeViolationInput\x1a%.eolymp.judge.DescribeViolationOutput\"K\xea\xe2\n" +
	"\v\xf5\xe2\n" +
	"\x00\x00\xa0@\xf8\xe2\n" +
	"\x14\x82\xe3\n" +
	"\x16\x8a\xe3\n" +
	"\x12judge:contest:read\x82\xd3\xe4\x93\x02\x1c\x12\x1a/violations/{violation_id}\x12\x95\x01\n" +
	"\x0eListViolations\x12!.eolymp.judge.ListViolationsInput\x1a\".eolymp.judge.ListViolationsOutput\"<\xea\xe2\n" +
	"\v\xf5\xe2\n" +
	"\x00\x00\xa0@\xf8\xe2\n" +
	"\x14\x82\xe3\n" +
	"\x16\x8a\xe3\n" +
	"\x12judge:contest:read\x82\xd3\xe4\x93\x02\r\x12\v/violationsB-Z+github.com/eolymp/go-sdk/eolymp/judge;judgeb\x06proto3"

var (
	file_eolymp_judge_violation_service_proto_rawDescOnce sync.Once
	file_eolymp_judge_violation_service_proto_rawDescData []byte
)

func file_eolymp_judge_violation_service_proto_rawDescGZIP() []byte {
	file_eolymp_judge_violation_service_proto_rawDescOnce.Do(func() {
		file_eolymp_judge_violation_service_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_eolymp_judge_violation_service_proto_rawDesc), len(file_eolymp_judge_violation_service_proto_rawDesc)))
	})
	return file_eolymp_judge_violation_service_proto_rawDescData
}

var file_eolymp_judge_violation_service_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_eolymp_judge_violation_service_proto_goTypes = []any{
	(*CreateViolationInput)(nil),       // 0: eolymp.judge.CreateViolationInput
	(*CreateViolationOutput)(nil),      // 1: eolymp.judge.CreateViolationOutput
	(*UpdateViolationInput)(nil),       // 2: eolymp.judge.UpdateViolationInput
	(*UpdateViolationOutput)(nil),      // 3: eolymp.judge.UpdateViolationOutput
	(*DeleteViolationInput)(nil),       // 4: eolymp.judge.DeleteViolationInput
	(*DeleteViolationOutput)(nil),      // 5: eolymp.judge.DeleteViolationOutput
	(*DescribeViolationInput)(nil),     // 6: eolymp.judge.DescribeViolationInput
	(*DescribeViolationOutput)(nil),    // 7: eolymp.judge.DescribeViolationOutput
	(*ListViolationsInput)(nil),        // 8: eolymp.judge.ListViolationsInput
	(*ListViolationsOutput)(nil),       // 9: eolymp.judge.ListViolationsOutput
	(*ListViolationsInput_Filter)(nil), // 10: eolymp.judge.ListViolationsInput.Filter
	(*Violation)(nil),                  // 11: eolymp.judge.Violation
	(Violation_Patch_Field)(0),         // 12: eolymp.judge.Violation.Patch.Field
	(*wellknown.ExpressionID)(nil),     // 13: eolymp.wellknown.ExpressionID
	(*wellknown.ExpressionEnum)(nil),   // 14: eolymp.wellknown.ExpressionEnum
	(*wellknown.ExpressionString)(nil), // 15: eolymp.wellknown.ExpressionString
	(*wellknown.ExpressionBool)(nil),   // 16: eolymp.wellknown.ExpressionBool
}
var file_eolymp_judge_violation_service_proto_depIdxs = []int32{
	11, // 0: eolymp.judge.CreateViolationInput.violation:type_name -> eolymp.judge.Violation
	12, // 1: eolymp.judge.UpdateViolationInput.patch:type_name -> eolymp.judge.Violation.Patch.Field
	11, // 2: eolymp.judge.UpdateViolationInput.violation:type_name -> eolymp.judge.Violation
	11, // 3: eolymp.judge.DescribeViolationOutput.violation:type_name -> eolymp.judge.Violation
	10, // 4: eolymp.judge.ListViolationsInput.filters:type_name -> eolymp.judge.ListViolationsInput.Filter
	11, // 5: eolymp.judge.ListViolationsOutput.items:type_name -> eolymp.judge.Violation
	13, // 6: eolymp.judge.ListViolationsInput.Filter.id:type_name -> eolymp.wellknown.ExpressionID
	14, // 7: eolymp.judge.ListViolationsInput.Filter.status:type_name -> eolymp.wellknown.ExpressionEnum
	14, // 8: eolymp.judge.ListViolationsInput.Filter.type:type_name -> eolymp.wellknown.ExpressionEnum
	15, // 9: eolymp.judge.ListViolationsInput.Filter.summary:type_name -> eolymp.wellknown.ExpressionString
	16, // 10: eolymp.judge.ListViolationsInput.Filter.automatic:type_name -> eolymp.wellknown.ExpressionBool
	13, // 11: eolymp.judge.ListViolationsInput.Filter.participant_id:type_name -> eolymp.wellknown.ExpressionID
	13, // 12: eolymp.judge.ListViolationsInput.Filter.submission_id:type_name -> eolymp.wellknown.ExpressionID
	13, // 13: eolymp.judge.ListViolationsInput.Filter.created_by:type_name -> eolymp.wellknown.ExpressionID
	13, // 14: eolymp.judge.ListViolationsInput.Filter.confirmed_by:type_name -> eolymp.wellknown.ExpressionID
	0,  // 15: eolymp.judge.ViolationService.CreateViolation:input_type -> eolymp.judge.CreateViolationInput
	2,  // 16: eolymp.judge.ViolationService.UpdateViolation:input_type -> eolymp.judge.UpdateViolationInput
	4,  // 17: eolymp.judge.ViolationService.DeleteViolation:input_type -> eolymp.judge.DeleteViolationInput
	6,  // 18: eolymp.judge.ViolationService.DescribeViolation:input_type -> eolymp.judge.DescribeViolationInput
	8,  // 19: eolymp.judge.ViolationService.ListViolations:input_type -> eolymp.judge.ListViolationsInput
	1,  // 20: eolymp.judge.ViolationService.CreateViolation:output_type -> eolymp.judge.CreateViolationOutput
	3,  // 21: eolymp.judge.ViolationService.UpdateViolation:output_type -> eolymp.judge.UpdateViolationOutput
	5,  // 22: eolymp.judge.ViolationService.DeleteViolation:output_type -> eolymp.judge.DeleteViolationOutput
	7,  // 23: eolymp.judge.ViolationService.DescribeViolation:output_type -> eolymp.judge.DescribeViolationOutput
	9,  // 24: eolymp.judge.ViolationService.ListViolations:output_type -> eolymp.judge.ListViolationsOutput
	20, // [20:25] is the sub-list for method output_type
	15, // [15:20] is the sub-list for method input_type
	15, // [15:15] is the sub-list for extension type_name
	15, // [15:15] is the sub-list for extension extendee
	0,  // [0:15] is the sub-list for field type_name
}

func init() { file_eolymp_judge_violation_service_proto_init() }
func file_eolymp_judge_violation_service_proto_init() {
	if File_eolymp_judge_violation_service_proto != nil {
		return
	}
	file_eolymp_judge_violation_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_eolymp_judge_violation_service_proto_rawDesc), len(file_eolymp_judge_violation_service_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_eolymp_judge_violation_service_proto_goTypes,
		DependencyIndexes: file_eolymp_judge_violation_service_proto_depIdxs,
		MessageInfos:      file_eolymp_judge_violation_service_proto_msgTypes,
	}.Build()
	File_eolymp_judge_violation_service_proto = out.File
	file_eolymp_judge_violation_service_proto_goTypes = nil
	file_eolymp_judge_violation_service_proto_depIdxs = nil
}
