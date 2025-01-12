// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.2
// 	protoc        v4.24.4
// source: eolymp/printer/printer_job.proto

package printer

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

type Job_Status int32

const (
	Job_UNKNOWN_STATUS Job_Status = 0
	Job_PENDING        Job_Status = 1
	Job_PRINTING       Job_Status = 2
	Job_COMPLETE       Job_Status = 3
	Job_ERROR          Job_Status = 4
	Job_EXPIRED        Job_Status = 5
)

// Enum value maps for Job_Status.
var (
	Job_Status_name = map[int32]string{
		0: "UNKNOWN_STATUS",
		1: "PENDING",
		2: "PRINTING",
		3: "COMPLETE",
		4: "ERROR",
		5: "EXPIRED",
	}
	Job_Status_value = map[string]int32{
		"UNKNOWN_STATUS": 0,
		"PENDING":        1,
		"PRINTING":       2,
		"COMPLETE":       3,
		"ERROR":          4,
		"EXPIRED":        5,
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
	return file_eolymp_printer_printer_job_proto_enumTypes[0].Descriptor()
}

func (Job_Status) Type() protoreflect.EnumType {
	return &file_eolymp_printer_printer_job_proto_enumTypes[0]
}

func (x Job_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Job_Status.Descriptor instead.
func (Job_Status) EnumDescriptor() ([]byte, []int) {
	return file_eolymp_printer_printer_job_proto_rawDescGZIP(), []int{0, 0}
}

type Job struct {
	state  protoimpl.MessageState `protogen:"open.v1"`
	Id     string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Status Job_Status             `protobuf:"varint,2,opt,name=status,proto3,enum=eolymp.printer.Job_Status" json:"status,omitempty"`
	// Types that are valid to be assigned to Creator:
	//
	//	*Job_UserId
	//	*Job_MemberId
	Creator       isJob_Creator `protobuf_oneof:"creator"`
	DocumentUrl   string        `protobuf:"bytes,20,opt,name=document_url,json=documentUrl,proto3" json:"document_url,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Job) Reset() {
	*x = Job{}
	mi := &file_eolymp_printer_printer_job_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Job) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Job) ProtoMessage() {}

func (x *Job) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_printer_printer_job_proto_msgTypes[0]
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
	return file_eolymp_printer_printer_job_proto_rawDescGZIP(), []int{0}
}

func (x *Job) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Job) GetStatus() Job_Status {
	if x != nil {
		return x.Status
	}
	return Job_UNKNOWN_STATUS
}

func (x *Job) GetCreator() isJob_Creator {
	if x != nil {
		return x.Creator
	}
	return nil
}

func (x *Job) GetUserId() string {
	if x != nil {
		if x, ok := x.Creator.(*Job_UserId); ok {
			return x.UserId
		}
	}
	return ""
}

func (x *Job) GetMemberId() string {
	if x != nil {
		if x, ok := x.Creator.(*Job_MemberId); ok {
			return x.MemberId
		}
	}
	return ""
}

func (x *Job) GetDocumentUrl() string {
	if x != nil {
		return x.DocumentUrl
	}
	return ""
}

type isJob_Creator interface {
	isJob_Creator()
}

type Job_UserId struct {
	UserId string `protobuf:"bytes,10,opt,name=user_id,json=userId,proto3,oneof"`
}

type Job_MemberId struct {
	MemberId string `protobuf:"bytes,11,opt,name=member_id,json=memberId,proto3,oneof"`
}

func (*Job_UserId) isJob_Creator() {}

func (*Job_MemberId) isJob_Creator() {}

var File_eolymp_printer_printer_job_proto protoreflect.FileDescriptor

var file_eolymp_printer_printer_job_proto_rawDesc = []byte{
	0x0a, 0x20, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x70, 0x72, 0x69, 0x6e, 0x74, 0x65, 0x72,
	0x2f, 0x70, 0x72, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x5f, 0x6a, 0x6f, 0x62, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x69, 0x6e, 0x74,
	0x65, 0x72, 0x22, 0x90, 0x02, 0x0a, 0x03, 0x4a, 0x6f, 0x62, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x32, 0x0a, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1a, 0x2e, 0x65, 0x6f, 0x6c,
	0x79, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x4a, 0x6f, 0x62, 0x2e,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x19,
	0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x48,
	0x00, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x09, 0x6d, 0x65, 0x6d,
	0x62, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x08,
	0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x64, 0x6f, 0x63, 0x75,
	0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x55, 0x72, 0x6c, 0x22, 0x5d, 0x0a, 0x06, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x12, 0x0a, 0x0e, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e,
	0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x50, 0x45, 0x4e,
	0x44, 0x49, 0x4e, 0x47, 0x10, 0x01, 0x12, 0x0c, 0x0a, 0x08, 0x50, 0x52, 0x49, 0x4e, 0x54, 0x49,
	0x4e, 0x47, 0x10, 0x02, 0x12, 0x0c, 0x0a, 0x08, 0x43, 0x4f, 0x4d, 0x50, 0x4c, 0x45, 0x54, 0x45,
	0x10, 0x03, 0x12, 0x09, 0x0a, 0x05, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x04, 0x12, 0x0b, 0x0a,
	0x07, 0x45, 0x58, 0x50, 0x49, 0x52, 0x45, 0x44, 0x10, 0x05, 0x42, 0x09, 0x0a, 0x07, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x6f, 0x72, 0x42, 0x31, 0x5a, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x67, 0x6f, 0x2d, 0x73, 0x64,
	0x6b, 0x2f, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x70, 0x72, 0x69, 0x6e, 0x74, 0x65, 0x72,
	0x3b, 0x70, 0x72, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_eolymp_printer_printer_job_proto_rawDescOnce sync.Once
	file_eolymp_printer_printer_job_proto_rawDescData = file_eolymp_printer_printer_job_proto_rawDesc
)

func file_eolymp_printer_printer_job_proto_rawDescGZIP() []byte {
	file_eolymp_printer_printer_job_proto_rawDescOnce.Do(func() {
		file_eolymp_printer_printer_job_proto_rawDescData = protoimpl.X.CompressGZIP(file_eolymp_printer_printer_job_proto_rawDescData)
	})
	return file_eolymp_printer_printer_job_proto_rawDescData
}

var file_eolymp_printer_printer_job_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_eolymp_printer_printer_job_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_eolymp_printer_printer_job_proto_goTypes = []any{
	(Job_Status)(0), // 0: eolymp.printer.Job.Status
	(*Job)(nil),     // 1: eolymp.printer.Job
}
var file_eolymp_printer_printer_job_proto_depIdxs = []int32{
	0, // 0: eolymp.printer.Job.status:type_name -> eolymp.printer.Job.Status
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_eolymp_printer_printer_job_proto_init() }
func file_eolymp_printer_printer_job_proto_init() {
	if File_eolymp_printer_printer_job_proto != nil {
		return
	}
	file_eolymp_printer_printer_job_proto_msgTypes[0].OneofWrappers = []any{
		(*Job_UserId)(nil),
		(*Job_MemberId)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_eolymp_printer_printer_job_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_eolymp_printer_printer_job_proto_goTypes,
		DependencyIndexes: file_eolymp_printer_printer_job_proto_depIdxs,
		EnumInfos:         file_eolymp_printer_printer_job_proto_enumTypes,
		MessageInfos:      file_eolymp_printer_printer_job_proto_msgTypes,
	}.Build()
	File_eolymp_printer_printer_job_proto = out.File
	file_eolymp_printer_printer_job_proto_rawDesc = nil
	file_eolymp_printer_printer_job_proto_goTypes = nil
	file_eolymp_printer_printer_job_proto_depIdxs = nil
}