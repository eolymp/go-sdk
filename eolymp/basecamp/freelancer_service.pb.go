// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v4.24.4
// source: eolymp/basecamp/freelancer_service.proto

package basecamp

import (
	_ "github.com/eolymp/go-sdk/eolymp/annotations"
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

type DescribeFreelancerStatusOutput_Status int32

const (
	DescribeFreelancerStatusOutput_UNKNOWN_STATUS DescribeFreelancerStatusOutput_Status = 0
	DescribeFreelancerStatusOutput_UNSIGNED       DescribeFreelancerStatusOutput_Status = 1 // freelancer hasn't sign the agreement
	DescribeFreelancerStatusOutput_PENDING        DescribeFreelancerStatusOutput_Status = 2 // agreement pending approval from Eolymp
	DescribeFreelancerStatusOutput_COMPLETE       DescribeFreelancerStatusOutput_Status = 3 // agreement is signed and complete
)

// Enum value maps for DescribeFreelancerStatusOutput_Status.
var (
	DescribeFreelancerStatusOutput_Status_name = map[int32]string{
		0: "UNKNOWN_STATUS",
		1: "UNSIGNED",
		2: "PENDING",
		3: "COMPLETE",
	}
	DescribeFreelancerStatusOutput_Status_value = map[string]int32{
		"UNKNOWN_STATUS": 0,
		"UNSIGNED":       1,
		"PENDING":        2,
		"COMPLETE":       3,
	}
)

func (x DescribeFreelancerStatusOutput_Status) Enum() *DescribeFreelancerStatusOutput_Status {
	p := new(DescribeFreelancerStatusOutput_Status)
	*p = x
	return p
}

func (x DescribeFreelancerStatusOutput_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DescribeFreelancerStatusOutput_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_eolymp_basecamp_freelancer_service_proto_enumTypes[0].Descriptor()
}

func (DescribeFreelancerStatusOutput_Status) Type() protoreflect.EnumType {
	return &file_eolymp_basecamp_freelancer_service_proto_enumTypes[0]
}

func (x DescribeFreelancerStatusOutput_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DescribeFreelancerStatusOutput_Status.Descriptor instead.
func (DescribeFreelancerStatusOutput_Status) EnumDescriptor() ([]byte, []int) {
	return file_eolymp_basecamp_freelancer_service_proto_rawDescGZIP(), []int{1, 0}
}

type DescribeFreelancerStatusInput struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DescribeFreelancerStatusInput) Reset() {
	*x = DescribeFreelancerStatusInput{}
	mi := &file_eolymp_basecamp_freelancer_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DescribeFreelancerStatusInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescribeFreelancerStatusInput) ProtoMessage() {}

func (x *DescribeFreelancerStatusInput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_basecamp_freelancer_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DescribeFreelancerStatusInput.ProtoReflect.Descriptor instead.
func (*DescribeFreelancerStatusInput) Descriptor() ([]byte, []int) {
	return file_eolymp_basecamp_freelancer_service_proto_rawDescGZIP(), []int{0}
}

type DescribeFreelancerStatusOutput struct {
	state         protoimpl.MessageState                `protogen:"open.v1"`
	Status        DescribeFreelancerStatusOutput_Status `protobuf:"varint,1,opt,name=status,proto3,enum=eolymp.basecamp.DescribeFreelancerStatusOutput_Status" json:"status,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DescribeFreelancerStatusOutput) Reset() {
	*x = DescribeFreelancerStatusOutput{}
	mi := &file_eolymp_basecamp_freelancer_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DescribeFreelancerStatusOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescribeFreelancerStatusOutput) ProtoMessage() {}

func (x *DescribeFreelancerStatusOutput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_basecamp_freelancer_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DescribeFreelancerStatusOutput.ProtoReflect.Descriptor instead.
func (*DescribeFreelancerStatusOutput) Descriptor() ([]byte, []int) {
	return file_eolymp_basecamp_freelancer_service_proto_rawDescGZIP(), []int{1}
}

func (x *DescribeFreelancerStatusOutput) GetStatus() DescribeFreelancerStatusOutput_Status {
	if x != nil {
		return x.Status
	}
	return DescribeFreelancerStatusOutput_UNKNOWN_STATUS
}

type UpdateFreelancerStatusInput struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	FirstName     string                 `protobuf:"bytes,1,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName      string                 `protobuf:"bytes,2,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	Address       string                 `protobuf:"bytes,3,opt,name=address,proto3" json:"address,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateFreelancerStatusInput) Reset() {
	*x = UpdateFreelancerStatusInput{}
	mi := &file_eolymp_basecamp_freelancer_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateFreelancerStatusInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateFreelancerStatusInput) ProtoMessage() {}

func (x *UpdateFreelancerStatusInput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_basecamp_freelancer_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateFreelancerStatusInput.ProtoReflect.Descriptor instead.
func (*UpdateFreelancerStatusInput) Descriptor() ([]byte, []int) {
	return file_eolymp_basecamp_freelancer_service_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateFreelancerStatusInput) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *UpdateFreelancerStatusInput) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *UpdateFreelancerStatusInput) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

type UpdateFreelancerStatusOutput struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	SignUrl       string                 `protobuf:"bytes,1,opt,name=sign_url,json=signUrl,proto3" json:"sign_url,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateFreelancerStatusOutput) Reset() {
	*x = UpdateFreelancerStatusOutput{}
	mi := &file_eolymp_basecamp_freelancer_service_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateFreelancerStatusOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateFreelancerStatusOutput) ProtoMessage() {}

func (x *UpdateFreelancerStatusOutput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_basecamp_freelancer_service_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateFreelancerStatusOutput.ProtoReflect.Descriptor instead.
func (*UpdateFreelancerStatusOutput) Descriptor() ([]byte, []int) {
	return file_eolymp_basecamp_freelancer_service_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateFreelancerStatusOutput) GetSignUrl() string {
	if x != nil {
		return x.SignUrl
	}
	return ""
}

var File_eolymp_basecamp_freelancer_service_proto protoreflect.FileDescriptor

const file_eolymp_basecamp_freelancer_service_proto_rawDesc = "" +
	"\n" +
	"(eolymp/basecamp/freelancer_service.proto\x12\x0feolymp.basecamp\x1a\x1deolymp/annotations/http.proto\x1a\"eolymp/annotations/ratelimit.proto\x1a\x1eeolymp/annotations/scope.proto\"\x1f\n" +
	"\x1dDescribeFreelancerStatusInput\"\xb7\x01\n" +
	"\x1eDescribeFreelancerStatusOutput\x12N\n" +
	"\x06status\x18\x01 \x01(\x0e26.eolymp.basecamp.DescribeFreelancerStatusOutput.StatusR\x06status\"E\n" +
	"\x06Status\x12\x12\n" +
	"\x0eUNKNOWN_STATUS\x10\x00\x12\f\n" +
	"\bUNSIGNED\x10\x01\x12\v\n" +
	"\aPENDING\x10\x02\x12\f\n" +
	"\bCOMPLETE\x10\x03\"s\n" +
	"\x1bUpdateFreelancerStatusInput\x12\x1d\n" +
	"\n" +
	"first_name\x18\x01 \x01(\tR\tfirstName\x12\x1b\n" +
	"\tlast_name\x18\x02 \x01(\tR\blastName\x12\x18\n" +
	"\aaddress\x18\x03 \x01(\tR\aaddress\"9\n" +
	"\x1cUpdateFreelancerStatusOutput\x12\x19\n" +
	"\bsign_url\x18\x01 \x01(\tR\asignUrl2\xf1\x02\n" +
	"\x11FreelancerService\x12\xaf\x01\n" +
	"\x18DescribeFreelancerStatus\x12..eolymp.basecamp.DescribeFreelancerStatusInput\x1a/.eolymp.basecamp.DescribeFreelancerStatusOutput\"2\xea\xe2\n" +
	"\v\xf5\xe2\n" +
	"\x00\x00\x80?\xf8\xe2\n" +
	"\n" +
	"\x82\xd3\xe4\x93\x02\x1d\x12\x1b/basecamp/freelancer-status\x12\xa9\x01\n" +
	"\x16UpdateFreelancerStatus\x12,.eolymp.basecamp.UpdateFreelancerStatusInput\x1a-.eolymp.basecamp.UpdateFreelancerStatusOutput\"2\xea\xe2\n" +
	"\v\xf5\xe2\n" +
	"\x00\x00\x80?\xf8\xe2\n" +
	"\n" +
	"\x82\xd3\xe4\x93\x02\x1d\"\x1b/basecamp/freelancer-statusB3Z1github.com/eolymp/go-sdk/eolymp/basecamp;basecampb\x06proto3"

var (
	file_eolymp_basecamp_freelancer_service_proto_rawDescOnce sync.Once
	file_eolymp_basecamp_freelancer_service_proto_rawDescData []byte
)

func file_eolymp_basecamp_freelancer_service_proto_rawDescGZIP() []byte {
	file_eolymp_basecamp_freelancer_service_proto_rawDescOnce.Do(func() {
		file_eolymp_basecamp_freelancer_service_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_eolymp_basecamp_freelancer_service_proto_rawDesc), len(file_eolymp_basecamp_freelancer_service_proto_rawDesc)))
	})
	return file_eolymp_basecamp_freelancer_service_proto_rawDescData
}

var file_eolymp_basecamp_freelancer_service_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_eolymp_basecamp_freelancer_service_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_eolymp_basecamp_freelancer_service_proto_goTypes = []any{
	(DescribeFreelancerStatusOutput_Status)(0), // 0: eolymp.basecamp.DescribeFreelancerStatusOutput.Status
	(*DescribeFreelancerStatusInput)(nil),      // 1: eolymp.basecamp.DescribeFreelancerStatusInput
	(*DescribeFreelancerStatusOutput)(nil),     // 2: eolymp.basecamp.DescribeFreelancerStatusOutput
	(*UpdateFreelancerStatusInput)(nil),        // 3: eolymp.basecamp.UpdateFreelancerStatusInput
	(*UpdateFreelancerStatusOutput)(nil),       // 4: eolymp.basecamp.UpdateFreelancerStatusOutput
}
var file_eolymp_basecamp_freelancer_service_proto_depIdxs = []int32{
	0, // 0: eolymp.basecamp.DescribeFreelancerStatusOutput.status:type_name -> eolymp.basecamp.DescribeFreelancerStatusOutput.Status
	1, // 1: eolymp.basecamp.FreelancerService.DescribeFreelancerStatus:input_type -> eolymp.basecamp.DescribeFreelancerStatusInput
	3, // 2: eolymp.basecamp.FreelancerService.UpdateFreelancerStatus:input_type -> eolymp.basecamp.UpdateFreelancerStatusInput
	2, // 3: eolymp.basecamp.FreelancerService.DescribeFreelancerStatus:output_type -> eolymp.basecamp.DescribeFreelancerStatusOutput
	4, // 4: eolymp.basecamp.FreelancerService.UpdateFreelancerStatus:output_type -> eolymp.basecamp.UpdateFreelancerStatusOutput
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_eolymp_basecamp_freelancer_service_proto_init() }
func file_eolymp_basecamp_freelancer_service_proto_init() {
	if File_eolymp_basecamp_freelancer_service_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_eolymp_basecamp_freelancer_service_proto_rawDesc), len(file_eolymp_basecamp_freelancer_service_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_eolymp_basecamp_freelancer_service_proto_goTypes,
		DependencyIndexes: file_eolymp_basecamp_freelancer_service_proto_depIdxs,
		EnumInfos:         file_eolymp_basecamp_freelancer_service_proto_enumTypes,
		MessageInfos:      file_eolymp_basecamp_freelancer_service_proto_msgTypes,
	}.Build()
	File_eolymp_basecamp_freelancer_service_proto = out.File
	file_eolymp_basecamp_freelancer_service_proto_goTypes = nil
	file_eolymp_basecamp_freelancer_service_proto_depIdxs = nil
}
