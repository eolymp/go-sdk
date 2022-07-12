// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.12.2
// source: eolymp/executor/verifier.proto

package executor

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

type Verifier_Type int32

const (
	// None verifier does not perform verification. This type should be used when requester
	// wants to execute source code and inspect output by itself.
	Verifier_NONE Verifier_Type = 0
	// Tokens verifies individual tokens:
	//  - whitespaces and new-lines do not affect result
	//  - numbers are compared as numbers and string tokens are compared as strings
	//  - numbers are compared with predefined precision
	//  - string tokens are compared with case sensitivity
	Verifier_TOKENS Verifier_Type = 1
	// Lines verifies individual lines:
	//  - trailing whitespaces do not affect result
	//  - final empty lines do not affect result
	Verifier_LINES Verifier_Type = 2
	// Program verifies output and answer using program `source` written in `lang`
	// Program is compatible with testlib.
	// Program runs with following arguments: <input-file> <output-file> <answer-file>.
	//
	// Additionally, program receives following environment variables:
	// - EOLYMP=1
	// - INPUT_FILE=<input-file>
	// - OUTPUT_FILE=<output-file>
	// - ANSWER_FILE=<answer-file>
	Verifier_PROGRAM Verifier_Type = 3
	// Program verifies output and answer using program `source` written in `lang`
	// Program is compatible with E-Olymp checkers (uses different argument order comparing to testlib).
	// Program runs with following arguments: <input-file> <answer-file> <output-file>.
	//
	// Additionally, program receives following environment variables:
	// - EOLYMP=1
	// - INPUT_FILE=<input-file>
	// - OUTPUT_FILE=<output-file>
	// - ANSWER_FILE=<answer-file>
	Verifier_LEGACY_PROGRAM Verifier_Type = 4
	// Query results verifies JSON encoded results of a query.
	Verifier_QUERY_RESULTS Verifier_Type = 5
)

// Enum value maps for Verifier_Type.
var (
	Verifier_Type_name = map[int32]string{
		0: "NONE",
		1: "TOKENS",
		2: "LINES",
		3: "PROGRAM",
		4: "LEGACY_PROGRAM",
		5: "QUERY_RESULTS",
	}
	Verifier_Type_value = map[string]int32{
		"NONE":           0,
		"TOKENS":         1,
		"LINES":          2,
		"PROGRAM":        3,
		"LEGACY_PROGRAM": 4,
		"QUERY_RESULTS":  5,
	}
)

func (x Verifier_Type) Enum() *Verifier_Type {
	p := new(Verifier_Type)
	*p = x
	return p
}

func (x Verifier_Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Verifier_Type) Descriptor() protoreflect.EnumDescriptor {
	return file_eolymp_executor_verifier_proto_enumTypes[0].Descriptor()
}

func (Verifier_Type) Type() protoreflect.EnumType {
	return &file_eolymp_executor_verifier_proto_enumTypes[0]
}

func (x Verifier_Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Verifier_Type.Descriptor instead.
func (Verifier_Type) EnumDescriptor() ([]byte, []int) {
	return file_eolymp_executor_verifier_proto_rawDescGZIP(), []int{0, 0}
}

// Verifier provides configuration on how to verify answers
type Verifier struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Verifier type (see types enumeration for details)
	Type Verifier_Type `protobuf:"varint,1,opt,name=type,proto3,enum=eolymp.executor.Verifier_Type" json:"type,omitempty"`
	// Programming language for PROGRAM verifier
	Lang string `protobuf:"bytes,2,opt,name=lang,proto3" json:"lang,omitempty"`
	// Source code for PROGRAM verifier
	Source string `protobuf:"bytes,3,opt,name=source,proto3" json:"source,omitempty"`
	// Precision for TOKEN verifier
	Precision int32 `protobuf:"varint,4,opt,name=precision,proto3" json:"precision,omitempty"`
	// Case sensitivity option for TOKEN verifier
	CaseSensitive bool `protobuf:"varint,5,opt,name=case_sensitive,json=caseSensitive,proto3" json:"case_sensitive,omitempty"`
	// Order sensitivity option for QUERY_RESULTS verifier.
	// If set to false the rows of output and answer will be sorted before comparison.
	OrderSensitive bool `protobuf:"varint,6,opt,name=order_sensitive,json=orderSensitive,proto3" json:"order_sensitive,omitempty"`
}

func (x *Verifier) Reset() {
	*x = Verifier{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_executor_verifier_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Verifier) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Verifier) ProtoMessage() {}

func (x *Verifier) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_executor_verifier_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Verifier.ProtoReflect.Descriptor instead.
func (*Verifier) Descriptor() ([]byte, []int) {
	return file_eolymp_executor_verifier_proto_rawDescGZIP(), []int{0}
}

func (x *Verifier) GetType() Verifier_Type {
	if x != nil {
		return x.Type
	}
	return Verifier_NONE
}

func (x *Verifier) GetLang() string {
	if x != nil {
		return x.Lang
	}
	return ""
}

func (x *Verifier) GetSource() string {
	if x != nil {
		return x.Source
	}
	return ""
}

func (x *Verifier) GetPrecision() int32 {
	if x != nil {
		return x.Precision
	}
	return 0
}

func (x *Verifier) GetCaseSensitive() bool {
	if x != nil {
		return x.CaseSensitive
	}
	return false
}

func (x *Verifier) GetOrderSensitive() bool {
	if x != nil {
		return x.OrderSensitive
	}
	return false
}

var File_eolymp_executor_verifier_proto protoreflect.FileDescriptor

var file_eolymp_executor_verifier_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x6f,
	0x72, 0x2f, 0x76, 0x65, 0x72, 0x69, 0x66, 0x69, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0f, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x6f,
	0x72, 0x22, 0xb5, 0x02, 0x0a, 0x08, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x65, 0x72, 0x12, 0x32,
	0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1e, 0x2e, 0x65,
	0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x6f, 0x72, 0x2e, 0x56,
	0x65, 0x72, 0x69, 0x66, 0x69, 0x65, 0x72, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x61, 0x6e, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6c, 0x61, 0x6e, 0x67, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x1c,
	0x0a, 0x09, 0x70, 0x72, 0x65, 0x63, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x09, 0x70, 0x72, 0x65, 0x63, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x25, 0x0a, 0x0e,
	0x63, 0x61, 0x73, 0x65, 0x5f, 0x73, 0x65, 0x6e, 0x73, 0x69, 0x74, 0x69, 0x76, 0x65, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x0d, 0x63, 0x61, 0x73, 0x65, 0x53, 0x65, 0x6e, 0x73, 0x69, 0x74,
	0x69, 0x76, 0x65, 0x12, 0x27, 0x0a, 0x0f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x6e,
	0x73, 0x69, 0x74, 0x69, 0x76, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0e, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x53, 0x65, 0x6e, 0x73, 0x69, 0x74, 0x69, 0x76, 0x65, 0x22, 0x5b, 0x0a, 0x04,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x0a,
	0x0a, 0x06, 0x54, 0x4f, 0x4b, 0x45, 0x4e, 0x53, 0x10, 0x01, 0x12, 0x09, 0x0a, 0x05, 0x4c, 0x49,
	0x4e, 0x45, 0x53, 0x10, 0x02, 0x12, 0x0b, 0x0a, 0x07, 0x50, 0x52, 0x4f, 0x47, 0x52, 0x41, 0x4d,
	0x10, 0x03, 0x12, 0x12, 0x0a, 0x0e, 0x4c, 0x45, 0x47, 0x41, 0x43, 0x59, 0x5f, 0x50, 0x52, 0x4f,
	0x47, 0x52, 0x41, 0x4d, 0x10, 0x04, 0x12, 0x11, 0x0a, 0x0d, 0x51, 0x55, 0x45, 0x52, 0x59, 0x5f,
	0x52, 0x45, 0x53, 0x55, 0x4c, 0x54, 0x53, 0x10, 0x05, 0x42, 0x33, 0x5a, 0x31, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x67,
	0x6f, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x65, 0x78, 0x65,
	0x63, 0x75, 0x74, 0x6f, 0x72, 0x3b, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x6f, 0x72, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_eolymp_executor_verifier_proto_rawDescOnce sync.Once
	file_eolymp_executor_verifier_proto_rawDescData = file_eolymp_executor_verifier_proto_rawDesc
)

func file_eolymp_executor_verifier_proto_rawDescGZIP() []byte {
	file_eolymp_executor_verifier_proto_rawDescOnce.Do(func() {
		file_eolymp_executor_verifier_proto_rawDescData = protoimpl.X.CompressGZIP(file_eolymp_executor_verifier_proto_rawDescData)
	})
	return file_eolymp_executor_verifier_proto_rawDescData
}

var file_eolymp_executor_verifier_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_eolymp_executor_verifier_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_eolymp_executor_verifier_proto_goTypes = []interface{}{
	(Verifier_Type)(0), // 0: eolymp.executor.Verifier.Type
	(*Verifier)(nil),   // 1: eolymp.executor.Verifier
}
var file_eolymp_executor_verifier_proto_depIdxs = []int32{
	0, // 0: eolymp.executor.Verifier.type:type_name -> eolymp.executor.Verifier.Type
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_eolymp_executor_verifier_proto_init() }
func file_eolymp_executor_verifier_proto_init() {
	if File_eolymp_executor_verifier_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_eolymp_executor_verifier_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Verifier); i {
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
			RawDescriptor: file_eolymp_executor_verifier_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_eolymp_executor_verifier_proto_goTypes,
		DependencyIndexes: file_eolymp_executor_verifier_proto_depIdxs,
		EnumInfos:         file_eolymp_executor_verifier_proto_enumTypes,
		MessageInfos:      file_eolymp_executor_verifier_proto_msgTypes,
	}.Build()
	File_eolymp_executor_verifier_proto = out.File
	file_eolymp_executor_verifier_proto_rawDesc = nil
	file_eolymp_executor_verifier_proto_goTypes = nil
	file_eolymp_executor_verifier_proto_depIdxs = nil
}
