// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v4.24.4
// source: eolymp/printer/printer_job.proto

package printer

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
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

type Job_Status int32

const (
	Job_UNKNOWN_STATUS Job_Status = 0
	Job_PENDING        Job_Status = 1
	Job_PRINTING       Job_Status = 2
	Job_COMPLETE       Job_Status = 3
	Job_ERROR          Job_Status = 4
	Job_CANCELLED      Job_Status = 5
)

// Enum value maps for Job_Status.
var (
	Job_Status_name = map[int32]string{
		0: "UNKNOWN_STATUS",
		1: "PENDING",
		2: "PRINTING",
		3: "COMPLETE",
		4: "ERROR",
		5: "CANCELLED",
	}
	Job_Status_value = map[string]int32{
		"UNKNOWN_STATUS": 0,
		"PENDING":        1,
		"PRINTING":       2,
		"COMPLETE":       3,
		"ERROR":          4,
		"CANCELLED":      5,
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
	Creator       isJob_Creator          `protobuf_oneof:"creator"`
	DocumentUrl   string                 `protobuf:"bytes,20,opt,name=document_url,json=documentUrl,proto3" json:"document_url,omitempty"`
	CreatedAt     *timestamppb.Timestamp `protobuf:"bytes,30,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt     *timestamppb.Timestamp `protobuf:"bytes,31,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
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

func (x *Job) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Job) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
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

const file_eolymp_printer_printer_job_proto_rawDesc = "" +
	"\n" +
	" eolymp/printer/printer_job.proto\x12\x0eeolymp.printer\x1a\x1fgoogle/protobuf/timestamp.proto\"\x88\x03\n" +
	"\x03Job\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x122\n" +
	"\x06status\x18\x02 \x01(\x0e2\x1a.eolymp.printer.Job.StatusR\x06status\x12\x19\n" +
	"\auser_id\x18\n" +
	" \x01(\tH\x00R\x06userId\x12\x1d\n" +
	"\tmember_id\x18\v \x01(\tH\x00R\bmemberId\x12!\n" +
	"\fdocument_url\x18\x14 \x01(\tR\vdocumentUrl\x129\n" +
	"\n" +
	"created_at\x18\x1e \x01(\v2\x1a.google.protobuf.TimestampR\tcreatedAt\x129\n" +
	"\n" +
	"updated_at\x18\x1f \x01(\v2\x1a.google.protobuf.TimestampR\tupdatedAt\"_\n" +
	"\x06Status\x12\x12\n" +
	"\x0eUNKNOWN_STATUS\x10\x00\x12\v\n" +
	"\aPENDING\x10\x01\x12\f\n" +
	"\bPRINTING\x10\x02\x12\f\n" +
	"\bCOMPLETE\x10\x03\x12\t\n" +
	"\x05ERROR\x10\x04\x12\r\n" +
	"\tCANCELLED\x10\x05B\t\n" +
	"\acreatorB1Z/github.com/eolymp/go-sdk/eolymp/printer;printerb\x06proto3"

var (
	file_eolymp_printer_printer_job_proto_rawDescOnce sync.Once
	file_eolymp_printer_printer_job_proto_rawDescData []byte
)

func file_eolymp_printer_printer_job_proto_rawDescGZIP() []byte {
	file_eolymp_printer_printer_job_proto_rawDescOnce.Do(func() {
		file_eolymp_printer_printer_job_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_eolymp_printer_printer_job_proto_rawDesc), len(file_eolymp_printer_printer_job_proto_rawDesc)))
	})
	return file_eolymp_printer_printer_job_proto_rawDescData
}

var file_eolymp_printer_printer_job_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_eolymp_printer_printer_job_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_eolymp_printer_printer_job_proto_goTypes = []any{
	(Job_Status)(0),               // 0: eolymp.printer.Job.Status
	(*Job)(nil),                   // 1: eolymp.printer.Job
	(*timestamppb.Timestamp)(nil), // 2: google.protobuf.Timestamp
}
var file_eolymp_printer_printer_job_proto_depIdxs = []int32{
	0, // 0: eolymp.printer.Job.status:type_name -> eolymp.printer.Job.Status
	2, // 1: eolymp.printer.Job.created_at:type_name -> google.protobuf.Timestamp
	2, // 2: eolymp.printer.Job.updated_at:type_name -> google.protobuf.Timestamp
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
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
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_eolymp_printer_printer_job_proto_rawDesc), len(file_eolymp_printer_printer_job_proto_rawDesc)),
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
	file_eolymp_printer_printer_job_proto_goTypes = nil
	file_eolymp_printer_printer_job_proto_depIdxs = nil
}
