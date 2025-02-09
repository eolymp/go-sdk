// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v4.24.4
// source: eolymp/atlas/solution.proto

package atlas

import (
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

type Solution_Type int32

const (
	Solution_UNSET                Solution_Type = 0
	Solution_CORRECT              Solution_Type = 1
	Solution_INCORRECT            Solution_Type = 2
	Solution_WRONG_ANSWER         Solution_Type = 3
	Solution_TIMEOUT              Solution_Type = 4
	Solution_OVERFLOW             Solution_Type = 5
	Solution_TIMEOUT_OR_ACCEPTED  Solution_Type = 6
	Solution_OVERFLOW_OR_ACCEPTED Solution_Type = 7
	Solution_DONT_RUN             Solution_Type = 8
	Solution_FAILURE              Solution_Type = 9
)

// Enum value maps for Solution_Type.
var (
	Solution_Type_name = map[int32]string{
		0: "UNSET",
		1: "CORRECT",
		2: "INCORRECT",
		3: "WRONG_ANSWER",
		4: "TIMEOUT",
		5: "OVERFLOW",
		6: "TIMEOUT_OR_ACCEPTED",
		7: "OVERFLOW_OR_ACCEPTED",
		8: "DONT_RUN",
		9: "FAILURE",
	}
	Solution_Type_value = map[string]int32{
		"UNSET":                0,
		"CORRECT":              1,
		"INCORRECT":            2,
		"WRONG_ANSWER":         3,
		"TIMEOUT":              4,
		"OVERFLOW":             5,
		"TIMEOUT_OR_ACCEPTED":  6,
		"OVERFLOW_OR_ACCEPTED": 7,
		"DONT_RUN":             8,
		"FAILURE":              9,
	}
)

func (x Solution_Type) Enum() *Solution_Type {
	p := new(Solution_Type)
	*p = x
	return p
}

func (x Solution_Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Solution_Type) Descriptor() protoreflect.EnumDescriptor {
	return file_eolymp_atlas_solution_proto_enumTypes[0].Descriptor()
}

func (Solution_Type) Type() protoreflect.EnumType {
	return &file_eolymp_atlas_solution_proto_enumTypes[0]
}

func (x Solution_Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Solution_Type.Descriptor instead.
func (Solution_Type) EnumDescriptor() ([]byte, []int) {
	return file_eolymp_atlas_solution_proto_rawDescGZIP(), []int{0, 0}
}

type Solution_Status int32

const (
	Solution_UNKNOWN_STATUS Solution_Status = 0
	Solution_PASS           Solution_Status = 1
	Solution_FAIL           Solution_Status = 2
	Solution_ERROR          Solution_Status = 3
)

// Enum value maps for Solution_Status.
var (
	Solution_Status_name = map[int32]string{
		0: "UNKNOWN_STATUS",
		1: "PASS",
		2: "FAIL",
		3: "ERROR",
	}
	Solution_Status_value = map[string]int32{
		"UNKNOWN_STATUS": 0,
		"PASS":           1,
		"FAIL":           2,
		"ERROR":          3,
	}
)

func (x Solution_Status) Enum() *Solution_Status {
	p := new(Solution_Status)
	*p = x
	return p
}

func (x Solution_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Solution_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_eolymp_atlas_solution_proto_enumTypes[1].Descriptor()
}

func (Solution_Status) Type() protoreflect.EnumType {
	return &file_eolymp_atlas_solution_proto_enumTypes[1]
}

func (x Solution_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Solution_Status.Descriptor instead.
func (Solution_Status) EnumDescriptor() ([]byte, []int) {
	return file_eolymp_atlas_solution_proto_rawDescGZIP(), []int{0, 1}
}

type Solution struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`            // unique identifier
	Name          string                 `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`        // solution name
	Secret        bool                   `protobuf:"varint,4,opt,name=secret,proto3" json:"secret,omitempty"`   // runtime and source are secret
	Runtime       string                 `protobuf:"bytes,10,opt,name=runtime,proto3" json:"runtime,omitempty"` // programming language
	Source        string                 `protobuf:"bytes,11,opt,name=source,proto3" json:"source,omitempty"`   // source code
	Type          Solution_Type          `protobuf:"varint,20,opt,name=type,proto3,enum=eolymp.atlas.Solution_Type" json:"type,omitempty"`
	Status        Solution_Status        `protobuf:"varint,30,opt,name=status,proto3,enum=eolymp.atlas.Solution_Status" json:"status,omitempty"`
	SubmissionId  string                 `protobuf:"bytes,31,opt,name=submission_id,json=submissionId,proto3" json:"submission_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Solution) Reset() {
	*x = Solution{}
	mi := &file_eolymp_atlas_solution_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Solution) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Solution) ProtoMessage() {}

func (x *Solution) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_atlas_solution_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Solution.ProtoReflect.Descriptor instead.
func (*Solution) Descriptor() ([]byte, []int) {
	return file_eolymp_atlas_solution_proto_rawDescGZIP(), []int{0}
}

func (x *Solution) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Solution) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Solution) GetSecret() bool {
	if x != nil {
		return x.Secret
	}
	return false
}

func (x *Solution) GetRuntime() string {
	if x != nil {
		return x.Runtime
	}
	return ""
}

func (x *Solution) GetSource() string {
	if x != nil {
		return x.Source
	}
	return ""
}

func (x *Solution) GetType() Solution_Type {
	if x != nil {
		return x.Type
	}
	return Solution_UNSET
}

func (x *Solution) GetStatus() Solution_Status {
	if x != nil {
		return x.Status
	}
	return Solution_UNKNOWN_STATUS
}

func (x *Solution) GetSubmissionId() string {
	if x != nil {
		return x.SubmissionId
	}
	return ""
}

var File_eolymp_atlas_solution_proto protoreflect.FileDescriptor

var file_eolymp_atlas_solution_proto_rawDesc = string([]byte{
	0x0a, 0x1b, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x61, 0x74, 0x6c, 0x61, 0x73, 0x2f, 0x73,
	0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x65,
	0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x61, 0x74, 0x6c, 0x61, 0x73, 0x22, 0xed, 0x03, 0x0a, 0x08,
	0x53, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x73, 0x65,
	0x63, 0x72, 0x65, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x2f, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x14,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x1b, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x61, 0x74,
	0x6c, 0x61, 0x73, 0x2e, 0x53, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x54, 0x79, 0x70,
	0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x35, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1d, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70,
	0x2e, 0x61, 0x74, 0x6c, 0x61, 0x73, 0x2e, 0x53, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x23,
	0x0a, 0x0d, 0x73, 0x75, 0x62, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18,
	0x1f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x73, 0x75, 0x62, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x49, 0x64, 0x22, 0xa8, 0x01, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x09, 0x0a, 0x05,
	0x55, 0x4e, 0x53, 0x45, 0x54, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x43, 0x4f, 0x52, 0x52, 0x45,
	0x43, 0x54, 0x10, 0x01, 0x12, 0x0d, 0x0a, 0x09, 0x49, 0x4e, 0x43, 0x4f, 0x52, 0x52, 0x45, 0x43,
	0x54, 0x10, 0x02, 0x12, 0x10, 0x0a, 0x0c, 0x57, 0x52, 0x4f, 0x4e, 0x47, 0x5f, 0x41, 0x4e, 0x53,
	0x57, 0x45, 0x52, 0x10, 0x03, 0x12, 0x0b, 0x0a, 0x07, 0x54, 0x49, 0x4d, 0x45, 0x4f, 0x55, 0x54,
	0x10, 0x04, 0x12, 0x0c, 0x0a, 0x08, 0x4f, 0x56, 0x45, 0x52, 0x46, 0x4c, 0x4f, 0x57, 0x10, 0x05,
	0x12, 0x17, 0x0a, 0x13, 0x54, 0x49, 0x4d, 0x45, 0x4f, 0x55, 0x54, 0x5f, 0x4f, 0x52, 0x5f, 0x41,
	0x43, 0x43, 0x45, 0x50, 0x54, 0x45, 0x44, 0x10, 0x06, 0x12, 0x18, 0x0a, 0x14, 0x4f, 0x56, 0x45,
	0x52, 0x46, 0x4c, 0x4f, 0x57, 0x5f, 0x4f, 0x52, 0x5f, 0x41, 0x43, 0x43, 0x45, 0x50, 0x54, 0x45,
	0x44, 0x10, 0x07, 0x12, 0x0c, 0x0a, 0x08, 0x44, 0x4f, 0x4e, 0x54, 0x5f, 0x52, 0x55, 0x4e, 0x10,
	0x08, 0x12, 0x0b, 0x0a, 0x07, 0x46, 0x41, 0x49, 0x4c, 0x55, 0x52, 0x45, 0x10, 0x09, 0x22, 0x3b,
	0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x12, 0x0a, 0x0e, 0x55, 0x4e, 0x4b, 0x4e,
	0x4f, 0x57, 0x4e, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04,
	0x50, 0x41, 0x53, 0x53, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x46, 0x41, 0x49, 0x4c, 0x10, 0x02,
	0x12, 0x09, 0x0a, 0x05, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x03, 0x42, 0x2d, 0x5a, 0x2b, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70,
	0x2f, 0x67, 0x6f, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x61,
	0x74, 0x6c, 0x61, 0x73, 0x3b, 0x61, 0x74, 0x6c, 0x61, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
})

var (
	file_eolymp_atlas_solution_proto_rawDescOnce sync.Once
	file_eolymp_atlas_solution_proto_rawDescData []byte
)

func file_eolymp_atlas_solution_proto_rawDescGZIP() []byte {
	file_eolymp_atlas_solution_proto_rawDescOnce.Do(func() {
		file_eolymp_atlas_solution_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_eolymp_atlas_solution_proto_rawDesc), len(file_eolymp_atlas_solution_proto_rawDesc)))
	})
	return file_eolymp_atlas_solution_proto_rawDescData
}

var file_eolymp_atlas_solution_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_eolymp_atlas_solution_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_eolymp_atlas_solution_proto_goTypes = []any{
	(Solution_Type)(0),   // 0: eolymp.atlas.Solution.Type
	(Solution_Status)(0), // 1: eolymp.atlas.Solution.Status
	(*Solution)(nil),     // 2: eolymp.atlas.Solution
}
var file_eolymp_atlas_solution_proto_depIdxs = []int32{
	0, // 0: eolymp.atlas.Solution.type:type_name -> eolymp.atlas.Solution.Type
	1, // 1: eolymp.atlas.Solution.status:type_name -> eolymp.atlas.Solution.Status
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_eolymp_atlas_solution_proto_init() }
func file_eolymp_atlas_solution_proto_init() {
	if File_eolymp_atlas_solution_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_eolymp_atlas_solution_proto_rawDesc), len(file_eolymp_atlas_solution_proto_rawDesc)),
			NumEnums:      2,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_eolymp_atlas_solution_proto_goTypes,
		DependencyIndexes: file_eolymp_atlas_solution_proto_depIdxs,
		EnumInfos:         file_eolymp_atlas_solution_proto_enumTypes,
		MessageInfos:      file_eolymp_atlas_solution_proto_msgTypes,
	}.Build()
	File_eolymp_atlas_solution_proto = out.File
	file_eolymp_atlas_solution_proto_goTypes = nil
	file_eolymp_atlas_solution_proto_depIdxs = nil
}
