// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.1
// 	protoc        v4.24.4
// source: eolymp/worker/worker_job.proto

package worker

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

type Job_Status int32

const (
	Job_UNKNOWN  Job_Status = 0
	Job_CREATED  Job_Status = 1
	Job_STARTED  Job_Status = 2
	Job_COMPLETE Job_Status = 3
	Job_ERROR    Job_Status = 4
)

// Enum value maps for Job_Status.
var (
	Job_Status_name = map[int32]string{
		0: "UNKNOWN",
		1: "CREATED",
		2: "STARTED",
		3: "COMPLETE",
		4: "ERROR",
	}
	Job_Status_value = map[string]int32{
		"UNKNOWN":  0,
		"CREATED":  1,
		"STARTED":  2,
		"COMPLETE": 3,
		"ERROR":    4,
	}
)

func (x Job_Status) Enum() *Job_Status {
	p := new(Job_Status)
	*p = x
	return p
}

func (x Job_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Job_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_eolymp_worker_worker_job_proto_enumTypes[0].Descriptor()
}

func (Job_Status) Type() protoreflect.EnumType {
	return &file_eolymp_worker_worker_job_proto_enumTypes[0]
}

func (x Job_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Job_Status.Descriptor instead.
func (Job_Status) EnumDescriptor() ([]byte, []int) {
	return file_eolymp_worker_worker_job_proto_rawDescGZIP(), []int{0, 0}
}

type Job_Patch int32

const (
	Job_UNKNOWN_PATCH  Job_Patch = 0
	Job_PATCH_ALL      Job_Patch = 1
	Job_PATCH_PROGRESS Job_Patch = 2
	Job_PATCH_OUTPUTS  Job_Patch = 3
	Job_PATCH_LOGS_URL Job_Patch = 4
	Job_PATCH_STATUS   Job_Patch = 5
)

// Enum value maps for Job_Patch.
var (
	Job_Patch_name = map[int32]string{
		0: "UNKNOWN_PATCH",
		1: "PATCH_ALL",
		2: "PATCH_PROGRESS",
		3: "PATCH_OUTPUTS",
		4: "PATCH_LOGS_URL",
		5: "PATCH_STATUS",
	}
	Job_Patch_value = map[string]int32{
		"UNKNOWN_PATCH":  0,
		"PATCH_ALL":      1,
		"PATCH_PROGRESS": 2,
		"PATCH_OUTPUTS":  3,
		"PATCH_LOGS_URL": 4,
		"PATCH_STATUS":   5,
	}
)

func (x Job_Patch) Enum() *Job_Patch {
	p := new(Job_Patch)
	*p = x
	return p
}

func (x Job_Patch) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Job_Patch) Descriptor() protoreflect.EnumDescriptor {
	return file_eolymp_worker_worker_job_proto_enumTypes[1].Descriptor()
}

func (Job_Patch) Type() protoreflect.EnumType {
	return &file_eolymp_worker_worker_job_proto_enumTypes[1]
}

func (x Job_Patch) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Job_Patch.Descriptor instead.
func (Job_Patch) EnumDescriptor() ([]byte, []int) {
	return file_eolymp_worker_worker_job_proto_rawDescGZIP(), []int{0, 1}
}

type Job struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Type          string                 `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	Namespace     string                 `protobuf:"bytes,3,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Status        Job_Status             `protobuf:"varint,30,opt,name=status,proto3,enum=eolymp.worker.Job_Status" json:"status,omitempty"`
	Progress      uint32                 `protobuf:"varint,20,opt,name=progress,proto3" json:"progress,omitempty"`
	Total         uint32                 `protobuf:"varint,21,opt,name=total,proto3" json:"total,omitempty"`
	Inputs        map[string]string      `protobuf:"bytes,4,rep,name=inputs,proto3" json:"inputs,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Outputs       map[string]string      `protobuf:"bytes,31,rep,name=outputs,proto3" json:"outputs,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	CreatedAt     *timestamppb.Timestamp `protobuf:"bytes,10,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	StartedAt     *timestamppb.Timestamp `protobuf:"bytes,11,opt,name=started_at,json=startedAt,proto3" json:"started_at,omitempty"`
	ProgressAt    *timestamppb.Timestamp `protobuf:"bytes,12,opt,name=progress_at,json=progressAt,proto3" json:"progress_at,omitempty"`
	CompleteAt    *timestamppb.Timestamp `protobuf:"bytes,13,opt,name=complete_at,json=completeAt,proto3" json:"complete_at,omitempty"`
	LogsUrl       string                 `protobuf:"bytes,100,opt,name=logs_url,json=logsUrl,proto3" json:"logs_url,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Job) Reset() {
	*x = Job{}
	mi := &file_eolymp_worker_worker_job_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Job) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Job) ProtoMessage() {}

func (x *Job) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_worker_worker_job_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Job.ProtoReflect.Descriptor instead.
func (*Job) Descriptor() ([]byte, []int) {
	return file_eolymp_worker_worker_job_proto_rawDescGZIP(), []int{0}
}

func (x *Job) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Job) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Job) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *Job) GetStatus() Job_Status {
	if x != nil {
		return x.Status
	}
	return Job_UNKNOWN
}

func (x *Job) GetProgress() uint32 {
	if x != nil {
		return x.Progress
	}
	return 0
}

func (x *Job) GetTotal() uint32 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *Job) GetInputs() map[string]string {
	if x != nil {
		return x.Inputs
	}
	return nil
}

func (x *Job) GetOutputs() map[string]string {
	if x != nil {
		return x.Outputs
	}
	return nil
}

func (x *Job) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Job) GetStartedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.StartedAt
	}
	return nil
}

func (x *Job) GetProgressAt() *timestamppb.Timestamp {
	if x != nil {
		return x.ProgressAt
	}
	return nil
}

func (x *Job) GetCompleteAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CompleteAt
	}
	return nil
}

func (x *Job) GetLogsUrl() string {
	if x != nil {
		return x.LogsUrl
	}
	return ""
}

var File_eolymp_worker_worker_job_proto protoreflect.FileDescriptor

var file_eolymp_worker_worker_job_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x2f,
	0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x5f, 0x6a, 0x6f, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0d, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x1a,
	0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xe3, 0x06, 0x0a, 0x03, 0x4a, 0x6f, 0x62, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1c, 0x0a, 0x09,
	0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x31, 0x0a, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x19, 0x2e, 0x65, 0x6f, 0x6c,
	0x79, 0x6d, 0x70, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x2e, 0x4a, 0x6f, 0x62, 0x2e, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1a, 0x0a,
	0x08, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x18, 0x14, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x08, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74,
	0x61, 0x6c, 0x18, 0x15, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12,
	0x36, 0x0a, 0x06, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x1e, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x2e,
	0x4a, 0x6f, 0x62, 0x2e, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52,
	0x06, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x73, 0x12, 0x39, 0x0a, 0x07, 0x6f, 0x75, 0x74, 0x70, 0x75,
	0x74, 0x73, 0x18, 0x1f, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d,
	0x70, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x2e, 0x4a, 0x6f, 0x62, 0x2e, 0x4f, 0x75, 0x74,
	0x70, 0x75, 0x74, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x6f, 0x75, 0x74, 0x70, 0x75,
	0x74, 0x73, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x39, 0x0a,
	0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x73,
	0x74, 0x61, 0x72, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x3b, 0x0a, 0x0b, 0x70, 0x72, 0x6f, 0x67,
	0x72, 0x65, 0x73, 0x73, 0x5f, 0x61, 0x74, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x70, 0x72, 0x6f, 0x67, 0x72,
	0x65, 0x73, 0x73, 0x41, 0x74, 0x12, 0x3b, 0x0a, 0x0b, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74,
	0x65, 0x5f, 0x61, 0x74, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65,
	0x41, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x6c, 0x6f, 0x67, 0x73, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x64,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6c, 0x6f, 0x67, 0x73, 0x55, 0x72, 0x6c, 0x1a, 0x39, 0x0a,
	0x0b, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x3a, 0x0a, 0x0c, 0x4f, 0x75, 0x74, 0x70,
	0x75, 0x74, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x3a, 0x02, 0x38, 0x01, 0x22, 0x48, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0b,
	0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x43,
	0x52, 0x45, 0x41, 0x54, 0x45, 0x44, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x53, 0x54, 0x41, 0x52,
	0x54, 0x45, 0x44, 0x10, 0x02, 0x12, 0x0c, 0x0a, 0x08, 0x43, 0x4f, 0x4d, 0x50, 0x4c, 0x45, 0x54,
	0x45, 0x10, 0x03, 0x12, 0x09, 0x0a, 0x05, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x04, 0x22, 0x76,
	0x0a, 0x05, 0x50, 0x61, 0x74, 0x63, 0x68, 0x12, 0x11, 0x0a, 0x0d, 0x55, 0x4e, 0x4b, 0x4e, 0x4f,
	0x57, 0x4e, 0x5f, 0x50, 0x41, 0x54, 0x43, 0x48, 0x10, 0x00, 0x12, 0x0d, 0x0a, 0x09, 0x50, 0x41,
	0x54, 0x43, 0x48, 0x5f, 0x41, 0x4c, 0x4c, 0x10, 0x01, 0x12, 0x12, 0x0a, 0x0e, 0x50, 0x41, 0x54,
	0x43, 0x48, 0x5f, 0x50, 0x52, 0x4f, 0x47, 0x52, 0x45, 0x53, 0x53, 0x10, 0x02, 0x12, 0x11, 0x0a,
	0x0d, 0x50, 0x41, 0x54, 0x43, 0x48, 0x5f, 0x4f, 0x55, 0x54, 0x50, 0x55, 0x54, 0x53, 0x10, 0x03,
	0x12, 0x12, 0x0a, 0x0e, 0x50, 0x41, 0x54, 0x43, 0x48, 0x5f, 0x4c, 0x4f, 0x47, 0x53, 0x5f, 0x55,
	0x52, 0x4c, 0x10, 0x04, 0x12, 0x10, 0x0a, 0x0c, 0x50, 0x41, 0x54, 0x43, 0x48, 0x5f, 0x53, 0x54,
	0x41, 0x54, 0x55, 0x53, 0x10, 0x05, 0x42, 0x2f, 0x5a, 0x2d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x67, 0x6f, 0x2d, 0x73,
	0x64, 0x6b, 0x2f, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72,
	0x3b, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_eolymp_worker_worker_job_proto_rawDescOnce sync.Once
	file_eolymp_worker_worker_job_proto_rawDescData = file_eolymp_worker_worker_job_proto_rawDesc
)

func file_eolymp_worker_worker_job_proto_rawDescGZIP() []byte {
	file_eolymp_worker_worker_job_proto_rawDescOnce.Do(func() {
		file_eolymp_worker_worker_job_proto_rawDescData = protoimpl.X.CompressGZIP(file_eolymp_worker_worker_job_proto_rawDescData)
	})
	return file_eolymp_worker_worker_job_proto_rawDescData
}

var file_eolymp_worker_worker_job_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_eolymp_worker_worker_job_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_eolymp_worker_worker_job_proto_goTypes = []any{
	(Job_Status)(0),               // 0: eolymp.worker.Job.Status
	(Job_Patch)(0),                // 1: eolymp.worker.Job.Patch
	(*Job)(nil),                   // 2: eolymp.worker.Job
	nil,                           // 3: eolymp.worker.Job.InputsEntry
	nil,                           // 4: eolymp.worker.Job.OutputsEntry
	(*timestamppb.Timestamp)(nil), // 5: google.protobuf.Timestamp
}
var file_eolymp_worker_worker_job_proto_depIdxs = []int32{
	0, // 0: eolymp.worker.Job.status:type_name -> eolymp.worker.Job.Status
	3, // 1: eolymp.worker.Job.inputs:type_name -> eolymp.worker.Job.InputsEntry
	4, // 2: eolymp.worker.Job.outputs:type_name -> eolymp.worker.Job.OutputsEntry
	5, // 3: eolymp.worker.Job.created_at:type_name -> google.protobuf.Timestamp
	5, // 4: eolymp.worker.Job.started_at:type_name -> google.protobuf.Timestamp
	5, // 5: eolymp.worker.Job.progress_at:type_name -> google.protobuf.Timestamp
	5, // 6: eolymp.worker.Job.complete_at:type_name -> google.protobuf.Timestamp
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_eolymp_worker_worker_job_proto_init() }
func file_eolymp_worker_worker_job_proto_init() {
	if File_eolymp_worker_worker_job_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_eolymp_worker_worker_job_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_eolymp_worker_worker_job_proto_goTypes,
		DependencyIndexes: file_eolymp_worker_worker_job_proto_depIdxs,
		EnumInfos:         file_eolymp_worker_worker_job_proto_enumTypes,
		MessageInfos:      file_eolymp_worker_worker_job_proto_msgTypes,
	}.Build()
	File_eolymp_worker_worker_job_proto = out.File
	file_eolymp_worker_worker_job_proto_rawDesc = nil
	file_eolymp_worker_worker_job_proto_goTypes = nil
	file_eolymp_worker_worker_job_proto_depIdxs = nil
}
