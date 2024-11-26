// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v3.21.12
// source: eolymp/worker/worker_service.proto

package worker

import (
	_ "github.com/eolymp/go-sdk/eolymp/annotations"
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

type CreateJobInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type   string            `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Inputs map[string]string `protobuf:"bytes,2,rep,name=inputs,proto3" json:"inputs,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *CreateJobInput) Reset() {
	*x = CreateJobInput{}
	mi := &file_eolymp_worker_worker_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateJobInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateJobInput) ProtoMessage() {}

func (x *CreateJobInput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_worker_worker_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateJobInput.ProtoReflect.Descriptor instead.
func (*CreateJobInput) Descriptor() ([]byte, []int) {
	return file_eolymp_worker_worker_service_proto_rawDescGZIP(), []int{0}
}

func (x *CreateJobInput) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *CreateJobInput) GetInputs() map[string]string {
	if x != nil {
		return x.Inputs
	}
	return nil
}

type CreateJobOutput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	JobId string `protobuf:"bytes,1,opt,name=job_id,json=jobId,proto3" json:"job_id,omitempty"`
}

func (x *CreateJobOutput) Reset() {
	*x = CreateJobOutput{}
	mi := &file_eolymp_worker_worker_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateJobOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateJobOutput) ProtoMessage() {}

func (x *CreateJobOutput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_worker_worker_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateJobOutput.ProtoReflect.Descriptor instead.
func (*CreateJobOutput) Descriptor() ([]byte, []int) {
	return file_eolymp_worker_worker_service_proto_rawDescGZIP(), []int{1}
}

func (x *CreateJobOutput) GetJobId() string {
	if x != nil {
		return x.JobId
	}
	return ""
}

type UpdateJobInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Patch []Job_Patch `protobuf:"varint,1,rep,packed,name=patch,proto3,enum=eolymp.worker.Job_Patch" json:"patch,omitempty"`
	JobId string      `protobuf:"bytes,2,opt,name=job_id,json=jobId,proto3" json:"job_id,omitempty"`
	Job   *Job        `protobuf:"bytes,3,opt,name=job,proto3" json:"job,omitempty"`
}

func (x *UpdateJobInput) Reset() {
	*x = UpdateJobInput{}
	mi := &file_eolymp_worker_worker_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateJobInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateJobInput) ProtoMessage() {}

func (x *UpdateJobInput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_worker_worker_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateJobInput.ProtoReflect.Descriptor instead.
func (*UpdateJobInput) Descriptor() ([]byte, []int) {
	return file_eolymp_worker_worker_service_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateJobInput) GetPatch() []Job_Patch {
	if x != nil {
		return x.Patch
	}
	return nil
}

func (x *UpdateJobInput) GetJobId() string {
	if x != nil {
		return x.JobId
	}
	return ""
}

func (x *UpdateJobInput) GetJob() *Job {
	if x != nil {
		return x.Job
	}
	return nil
}

type UpdateJobOutput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateJobOutput) Reset() {
	*x = UpdateJobOutput{}
	mi := &file_eolymp_worker_worker_service_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateJobOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateJobOutput) ProtoMessage() {}

func (x *UpdateJobOutput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_worker_worker_service_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateJobOutput.ProtoReflect.Descriptor instead.
func (*UpdateJobOutput) Descriptor() ([]byte, []int) {
	return file_eolymp_worker_worker_service_proto_rawDescGZIP(), []int{3}
}

type DescribeJobInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	JobId string `protobuf:"bytes,1,opt,name=job_id,json=jobId,proto3" json:"job_id,omitempty"`
}

func (x *DescribeJobInput) Reset() {
	*x = DescribeJobInput{}
	mi := &file_eolymp_worker_worker_service_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DescribeJobInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescribeJobInput) ProtoMessage() {}

func (x *DescribeJobInput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_worker_worker_service_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DescribeJobInput.ProtoReflect.Descriptor instead.
func (*DescribeJobInput) Descriptor() ([]byte, []int) {
	return file_eolymp_worker_worker_service_proto_rawDescGZIP(), []int{4}
}

func (x *DescribeJobInput) GetJobId() string {
	if x != nil {
		return x.JobId
	}
	return ""
}

type DescribeJobOutput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Job *Job `protobuf:"bytes,1,opt,name=job,proto3" json:"job,omitempty"`
}

func (x *DescribeJobOutput) Reset() {
	*x = DescribeJobOutput{}
	mi := &file_eolymp_worker_worker_service_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DescribeJobOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescribeJobOutput) ProtoMessage() {}

func (x *DescribeJobOutput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_worker_worker_service_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DescribeJobOutput.ProtoReflect.Descriptor instead.
func (*DescribeJobOutput) Descriptor() ([]byte, []int) {
	return file_eolymp_worker_worker_service_proto_rawDescGZIP(), []int{5}
}

func (x *DescribeJobOutput) GetJob() *Job {
	if x != nil {
		return x.Job
	}
	return nil
}

type ListJobsInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// pagination
	Offset int32 `protobuf:"varint,10,opt,name=offset,proto3" json:"offset,omitempty"`
	Size   int32 `protobuf:"varint,11,opt,name=size,proto3" json:"size,omitempty"`
	// data filters
	Filters *ListJobsInput_Filter `protobuf:"bytes,40,opt,name=filters,proto3" json:"filters,omitempty"`
}

func (x *ListJobsInput) Reset() {
	*x = ListJobsInput{}
	mi := &file_eolymp_worker_worker_service_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListJobsInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListJobsInput) ProtoMessage() {}

func (x *ListJobsInput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_worker_worker_service_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListJobsInput.ProtoReflect.Descriptor instead.
func (*ListJobsInput) Descriptor() ([]byte, []int) {
	return file_eolymp_worker_worker_service_proto_rawDescGZIP(), []int{6}
}

func (x *ListJobsInput) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *ListJobsInput) GetSize() int32 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *ListJobsInput) GetFilters() *ListJobsInput_Filter {
	if x != nil {
		return x.Filters
	}
	return nil
}

type ListJobsOutput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Total int32  `protobuf:"varint,1,opt,name=total,proto3" json:"total,omitempty"`
	Items []*Job `protobuf:"bytes,2,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *ListJobsOutput) Reset() {
	*x = ListJobsOutput{}
	mi := &file_eolymp_worker_worker_service_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListJobsOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListJobsOutput) ProtoMessage() {}

func (x *ListJobsOutput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_worker_worker_service_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListJobsOutput.ProtoReflect.Descriptor instead.
func (*ListJobsOutput) Descriptor() ([]byte, []int) {
	return file_eolymp_worker_worker_service_proto_rawDescGZIP(), []int{7}
}

func (x *ListJobsOutput) GetTotal() int32 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *ListJobsOutput) GetItems() []*Job {
	if x != nil {
		return x.Items
	}
	return nil
}

type ListJobsInput_Filter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     []*wellknown.ExpressionID   `protobuf:"bytes,1,rep,name=id,proto3" json:"id,omitempty"`
	Type   []*wellknown.ExpressionEnum `protobuf:"bytes,2,rep,name=type,proto3" json:"type,omitempty"`
	Status []*wellknown.ExpressionEnum `protobuf:"bytes,3,rep,name=status,proto3" json:"status,omitempty"`
}

func (x *ListJobsInput_Filter) Reset() {
	*x = ListJobsInput_Filter{}
	mi := &file_eolymp_worker_worker_service_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListJobsInput_Filter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListJobsInput_Filter) ProtoMessage() {}

func (x *ListJobsInput_Filter) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_worker_worker_service_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListJobsInput_Filter.ProtoReflect.Descriptor instead.
func (*ListJobsInput_Filter) Descriptor() ([]byte, []int) {
	return file_eolymp_worker_worker_service_proto_rawDescGZIP(), []int{6, 0}
}

func (x *ListJobsInput_Filter) GetId() []*wellknown.ExpressionID {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *ListJobsInput_Filter) GetType() []*wellknown.ExpressionEnum {
	if x != nil {
		return x.Type
	}
	return nil
}

func (x *ListJobsInput_Filter) GetStatus() []*wellknown.ExpressionEnum {
	if x != nil {
		return x.Status
	}
	return nil
}

var File_eolymp_worker_worker_service_proto protoreflect.FileDescriptor

var file_eolymp_worker_worker_service_proto_rawDesc = []byte{
	0x0a, 0x22, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x2f,
	0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x77, 0x6f, 0x72,
	0x6b, 0x65, 0x72, 0x1a, 0x1d, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x22, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x72, 0x61, 0x74, 0x65, 0x6c, 0x69, 0x6d, 0x69, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x61,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x63, 0x6f, 0x70, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x21, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x77,
	0x65, 0x6c, 0x6c, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x2f, 0x65, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x65, 0x6f, 0x6c, 0x79, 0x6d,
	0x70, 0x2f, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x2f, 0x6a, 0x6f, 0x62, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xa2, 0x01, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4a, 0x6f, 0x62,
	0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x41, 0x0a, 0x06, 0x69, 0x6e, 0x70,
	0x75, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x65, 0x6f, 0x6c, 0x79,
	0x6d, 0x70, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x4a, 0x6f, 0x62, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x2e, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x73, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x73, 0x1a, 0x39, 0x0a, 0x0b,
	0x49, 0x6e, 0x70, 0x75, 0x74, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x28, 0x0a, 0x0f, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x4a, 0x6f, 0x62, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x12, 0x15, 0x0a, 0x06, 0x6a, 0x6f,
	0x62, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6a, 0x6f, 0x62, 0x49,
	0x64, 0x22, 0x7d, 0x0a, 0x0e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4a, 0x6f, 0x62, 0x49, 0x6e,
	0x70, 0x75, 0x74, 0x12, 0x2e, 0x0a, 0x05, 0x70, 0x61, 0x74, 0x63, 0x68, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0e, 0x32, 0x18, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x77, 0x6f, 0x72, 0x6b,
	0x65, 0x72, 0x2e, 0x4a, 0x6f, 0x62, 0x2e, 0x50, 0x61, 0x74, 0x63, 0x68, 0x52, 0x05, 0x70, 0x61,
	0x74, 0x63, 0x68, 0x12, 0x15, 0x0a, 0x06, 0x6a, 0x6f, 0x62, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x6a, 0x6f, 0x62, 0x49, 0x64, 0x12, 0x24, 0x0a, 0x03, 0x6a, 0x6f,
	0x62, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70,
	0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x2e, 0x4a, 0x6f, 0x62, 0x52, 0x03, 0x6a, 0x6f, 0x62,
	0x22, 0x11, 0x0a, 0x0f, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4a, 0x6f, 0x62, 0x4f, 0x75, 0x74,
	0x70, 0x75, 0x74, 0x22, 0x29, 0x0a, 0x10, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x4a,
	0x6f, 0x62, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x15, 0x0a, 0x06, 0x6a, 0x6f, 0x62, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6a, 0x6f, 0x62, 0x49, 0x64, 0x22, 0x39,
	0x0a, 0x11, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x4a, 0x6f, 0x62, 0x4f, 0x75, 0x74,
	0x70, 0x75, 0x74, 0x12, 0x24, 0x0a, 0x03, 0x6a, 0x6f, 0x62, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x12, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72,
	0x2e, 0x4a, 0x6f, 0x62, 0x52, 0x03, 0x6a, 0x6f, 0x62, 0x22, 0xa5, 0x02, 0x0a, 0x0d, 0x4c, 0x69,
	0x73, 0x74, 0x4a, 0x6f, 0x62, 0x73, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6f,
	0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6f, 0x66, 0x66,
	0x73, 0x65, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x3d, 0x0a, 0x07, 0x66, 0x69, 0x6c, 0x74, 0x65,
	0x72, 0x73, 0x18, 0x28, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d,
	0x70, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4a, 0x6f, 0x62,
	0x73, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x2e, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x52, 0x07, 0x66,
	0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x1a, 0xa8, 0x01, 0x0a, 0x06, 0x46, 0x69, 0x6c, 0x74, 0x65,
	0x72, 0x12, 0x2e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e,
	0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x77, 0x65, 0x6c, 0x6c, 0x6b, 0x6e, 0x6f, 0x77, 0x6e,
	0x2e, 0x45, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x34, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x20, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x77, 0x65, 0x6c, 0x6c, 0x6b, 0x6e, 0x6f,
	0x77, 0x6e, 0x2e, 0x45, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x45, 0x6e, 0x75,
	0x6d, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x38, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70,
	0x2e, 0x77, 0x65, 0x6c, 0x6c, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x2e, 0x45, 0x78, 0x70, 0x72, 0x65,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x45, 0x6e, 0x75, 0x6d, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x22, 0x50, 0x0a, 0x0e, 0x4c, 0x69, 0x73, 0x74, 0x4a, 0x6f, 0x62, 0x73, 0x4f, 0x75, 0x74,
	0x70, 0x75, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x28, 0x0a, 0x05, 0x69, 0x74, 0x65,
	0x6d, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d,
	0x70, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x2e, 0x4a, 0x6f, 0x62, 0x52, 0x05, 0x69, 0x74,
	0x65, 0x6d, 0x73, 0x32, 0xba, 0x03, 0x0a, 0x0d, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4c, 0x0a, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4a,
	0x6f, 0x62, 0x12, 0x1d, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x77, 0x6f, 0x72, 0x6b,
	0x65, 0x72, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4a, 0x6f, 0x62, 0x49, 0x6e, 0x70, 0x75,
	0x74, 0x1a, 0x1e, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x65,
	0x72, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4a, 0x6f, 0x62, 0x4f, 0x75, 0x74, 0x70, 0x75,
	0x74, 0x22, 0x00, 0x12, 0x4c, 0x0a, 0x09, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4a, 0x6f, 0x62,
	0x12, 0x1d, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72,
	0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4a, 0x6f, 0x62, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a,
	0x1e, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x2e,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4a, 0x6f, 0x62, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x22,
	0x00, 0x12, 0x8e, 0x01, 0x0a, 0x0b, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x4a, 0x6f,
	0x62, 0x12, 0x1f, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x65,
	0x72, 0x2e, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x4a, 0x6f, 0x62, 0x49, 0x6e, 0x70,
	0x75, 0x74, 0x1a, 0x20, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x77, 0x6f, 0x72, 0x6b,
	0x65, 0x72, 0x2e, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x4a, 0x6f, 0x62, 0x4f, 0x75,
	0x74, 0x70, 0x75, 0x74, 0x22, 0x3c, 0xea, 0xe2, 0x0a, 0x0b, 0xf5, 0xe2, 0x0a, 0x00, 0x00, 0x00,
	0x40, 0xf8, 0xe2, 0x0a, 0x0a, 0x82, 0xe3, 0x0a, 0x13, 0x8a, 0xe3, 0x0a, 0x0f, 0x77, 0x6f, 0x72,
	0x6b, 0x65, 0x72, 0x3a, 0x6a, 0x6f, 0x62, 0x3a, 0x72, 0x65, 0x61, 0x64, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x10, 0x12, 0x0e, 0x2f, 0x6a, 0x6f, 0x62, 0x73, 0x2f, 0x7b, 0x6a, 0x6f, 0x62, 0x5f, 0x69,
	0x64, 0x7d, 0x12, 0x7c, 0x0a, 0x08, 0x4c, 0x69, 0x73, 0x74, 0x4a, 0x6f, 0x62, 0x73, 0x12, 0x1c,
	0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x4a, 0x6f, 0x62, 0x73, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x1d, 0x2e, 0x65,
	0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x4a, 0x6f, 0x62, 0x73, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x22, 0x33, 0xea, 0xe2, 0x0a,
	0x0b, 0xf5, 0xe2, 0x0a, 0x00, 0x00, 0x00, 0x40, 0xf8, 0xe2, 0x0a, 0x0a, 0x82, 0xe3, 0x0a, 0x13,
	0x8a, 0xe3, 0x0a, 0x0f, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x3a, 0x6a, 0x6f, 0x62, 0x3a, 0x72,
	0x65, 0x61, 0x64, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x07, 0x12, 0x05, 0x2f, 0x6a, 0x6f, 0x62, 0x73,
	0x42, 0x2f, 0x5a, 0x2d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65,
	0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x67, 0x6f, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x65, 0x6f, 0x6c,
	0x79, 0x6d, 0x70, 0x2f, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x3b, 0x77, 0x6f, 0x72, 0x6b, 0x65,
	0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_eolymp_worker_worker_service_proto_rawDescOnce sync.Once
	file_eolymp_worker_worker_service_proto_rawDescData = file_eolymp_worker_worker_service_proto_rawDesc
)

func file_eolymp_worker_worker_service_proto_rawDescGZIP() []byte {
	file_eolymp_worker_worker_service_proto_rawDescOnce.Do(func() {
		file_eolymp_worker_worker_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_eolymp_worker_worker_service_proto_rawDescData)
	})
	return file_eolymp_worker_worker_service_proto_rawDescData
}

var file_eolymp_worker_worker_service_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_eolymp_worker_worker_service_proto_goTypes = []any{
	(*CreateJobInput)(nil),           // 0: eolymp.worker.CreateJobInput
	(*CreateJobOutput)(nil),          // 1: eolymp.worker.CreateJobOutput
	(*UpdateJobInput)(nil),           // 2: eolymp.worker.UpdateJobInput
	(*UpdateJobOutput)(nil),          // 3: eolymp.worker.UpdateJobOutput
	(*DescribeJobInput)(nil),         // 4: eolymp.worker.DescribeJobInput
	(*DescribeJobOutput)(nil),        // 5: eolymp.worker.DescribeJobOutput
	(*ListJobsInput)(nil),            // 6: eolymp.worker.ListJobsInput
	(*ListJobsOutput)(nil),           // 7: eolymp.worker.ListJobsOutput
	nil,                              // 8: eolymp.worker.CreateJobInput.InputsEntry
	(*ListJobsInput_Filter)(nil),     // 9: eolymp.worker.ListJobsInput.Filter
	(Job_Patch)(0),                   // 10: eolymp.worker.Job.Patch
	(*Job)(nil),                      // 11: eolymp.worker.Job
	(*wellknown.ExpressionID)(nil),   // 12: eolymp.wellknown.ExpressionID
	(*wellknown.ExpressionEnum)(nil), // 13: eolymp.wellknown.ExpressionEnum
}
var file_eolymp_worker_worker_service_proto_depIdxs = []int32{
	8,  // 0: eolymp.worker.CreateJobInput.inputs:type_name -> eolymp.worker.CreateJobInput.InputsEntry
	10, // 1: eolymp.worker.UpdateJobInput.patch:type_name -> eolymp.worker.Job.Patch
	11, // 2: eolymp.worker.UpdateJobInput.job:type_name -> eolymp.worker.Job
	11, // 3: eolymp.worker.DescribeJobOutput.job:type_name -> eolymp.worker.Job
	9,  // 4: eolymp.worker.ListJobsInput.filters:type_name -> eolymp.worker.ListJobsInput.Filter
	11, // 5: eolymp.worker.ListJobsOutput.items:type_name -> eolymp.worker.Job
	12, // 6: eolymp.worker.ListJobsInput.Filter.id:type_name -> eolymp.wellknown.ExpressionID
	13, // 7: eolymp.worker.ListJobsInput.Filter.type:type_name -> eolymp.wellknown.ExpressionEnum
	13, // 8: eolymp.worker.ListJobsInput.Filter.status:type_name -> eolymp.wellknown.ExpressionEnum
	0,  // 9: eolymp.worker.WorkerService.CreateJob:input_type -> eolymp.worker.CreateJobInput
	2,  // 10: eolymp.worker.WorkerService.UpdateJob:input_type -> eolymp.worker.UpdateJobInput
	4,  // 11: eolymp.worker.WorkerService.DescribeJob:input_type -> eolymp.worker.DescribeJobInput
	6,  // 12: eolymp.worker.WorkerService.ListJobs:input_type -> eolymp.worker.ListJobsInput
	1,  // 13: eolymp.worker.WorkerService.CreateJob:output_type -> eolymp.worker.CreateJobOutput
	3,  // 14: eolymp.worker.WorkerService.UpdateJob:output_type -> eolymp.worker.UpdateJobOutput
	5,  // 15: eolymp.worker.WorkerService.DescribeJob:output_type -> eolymp.worker.DescribeJobOutput
	7,  // 16: eolymp.worker.WorkerService.ListJobs:output_type -> eolymp.worker.ListJobsOutput
	13, // [13:17] is the sub-list for method output_type
	9,  // [9:13] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_eolymp_worker_worker_service_proto_init() }
func file_eolymp_worker_worker_service_proto_init() {
	if File_eolymp_worker_worker_service_proto != nil {
		return
	}
	file_eolymp_worker_job_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_eolymp_worker_worker_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_eolymp_worker_worker_service_proto_goTypes,
		DependencyIndexes: file_eolymp_worker_worker_service_proto_depIdxs,
		MessageInfos:      file_eolymp_worker_worker_service_proto_msgTypes,
	}.Build()
	File_eolymp_worker_worker_service_proto = out.File
	file_eolymp_worker_worker_service_proto_rawDesc = nil
	file_eolymp_worker_worker_service_proto_goTypes = nil
	file_eolymp_worker_worker_service_proto_depIdxs = nil
}
