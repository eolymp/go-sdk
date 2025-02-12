// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v4.24.4
// source: eolymp/atlas/testing_checker.proto

package atlas

import (
	executor "github.com/eolymp/go-sdk/eolymp/executor"
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

// Checker provides configuration on how to verify answers
type Checker struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	VersionId     string                 `protobuf:"bytes,11,opt,name=version_id,json=versionId,proto3" json:"version_id,omitempty"`
	Secret        bool                   `protobuf:"varint,7,opt,name=secret,proto3" json:"secret,omitempty"` // Secret means checker code and configuration must not be exposed to users
	Type          executor.Checker_Type  `protobuf:"varint,1,opt,name=type,proto3,enum=eolymp.executor.Checker_Type" json:"type,omitempty"`
	Runtime       string                 `protobuf:"bytes,2,opt,name=runtime,proto3" json:"runtime,omitempty"`                                   // Programming language for PROGRAM checker
	Source        string                 `protobuf:"bytes,3,opt,name=source,proto3" json:"source,omitempty"`                                     // Source code for PROGRAM checker
	Files         []*executor.File       `protobuf:"bytes,10,rep,name=files,proto3" json:"files,omitempty"`                                      // Additional files placed into workdir during compilation and execution
	Precision     int32                  `protobuf:"varint,4,opt,name=precision,proto3" json:"precision,omitempty"`                              // Precision for TOKEN checker
	CaseSensitive bool                   `protobuf:"varint,5,opt,name=case_sensitive,json=caseSensitive,proto3" json:"case_sensitive,omitempty"` // Case sensitivity option for TOKEN checker
	// Order sensitivity option for QUERY_RESULTS checker.
	// If set to false the rows of output and answer will be sorted before comparison.
	OrderSensitive bool `protobuf:"varint,6,opt,name=order_sensitive,json=orderSensitive,proto3" json:"order_sensitive,omitempty"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *Checker) Reset() {
	*x = Checker{}
	mi := &file_eolymp_atlas_testing_checker_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Checker) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Checker) ProtoMessage() {}

func (x *Checker) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_atlas_testing_checker_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Checker.ProtoReflect.Descriptor instead.
func (*Checker) Descriptor() ([]byte, []int) {
	return file_eolymp_atlas_testing_checker_proto_rawDescGZIP(), []int{0}
}

func (x *Checker) GetVersionId() string {
	if x != nil {
		return x.VersionId
	}
	return ""
}

func (x *Checker) GetSecret() bool {
	if x != nil {
		return x.Secret
	}
	return false
}

func (x *Checker) GetType() executor.Checker_Type {
	if x != nil {
		return x.Type
	}
	return executor.Checker_Type(0)
}

func (x *Checker) GetRuntime() string {
	if x != nil {
		return x.Runtime
	}
	return ""
}

func (x *Checker) GetSource() string {
	if x != nil {
		return x.Source
	}
	return ""
}

func (x *Checker) GetFiles() []*executor.File {
	if x != nil {
		return x.Files
	}
	return nil
}

func (x *Checker) GetPrecision() int32 {
	if x != nil {
		return x.Precision
	}
	return 0
}

func (x *Checker) GetCaseSensitive() bool {
	if x != nil {
		return x.CaseSensitive
	}
	return false
}

func (x *Checker) GetOrderSensitive() bool {
	if x != nil {
		return x.OrderSensitive
	}
	return false
}

var File_eolymp_atlas_testing_checker_proto protoreflect.FileDescriptor

var file_eolymp_atlas_testing_checker_proto_rawDesc = string([]byte{
	0x0a, 0x22, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x61, 0x74, 0x6c, 0x61, 0x73, 0x2f, 0x74,
	0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x65, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x61, 0x74, 0x6c,
	0x61, 0x73, 0x1a, 0x1d, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x65, 0x78, 0x65, 0x63, 0x75,
	0x74, 0x6f, 0x72, 0x2f, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1a, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74,
	0x6f, 0x72, 0x2f, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc0, 0x02,
	0x0a, 0x07, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x65, 0x72, 0x12, 0x1d, 0x0a, 0x0a, 0x76, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x76,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x63, 0x72,
	0x65, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74,
	0x12, 0x31, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1d,
	0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x6f, 0x72,
	0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x65, 0x72, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x2b, 0x0a, 0x05, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x18, 0x0a,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x65, 0x78,
	0x65, 0x63, 0x75, 0x74, 0x6f, 0x72, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x05, 0x66, 0x69, 0x6c,
	0x65, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x72, 0x65, 0x63, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x70, 0x72, 0x65, 0x63, 0x69, 0x73, 0x69, 0x6f, 0x6e,
	0x12, 0x25, 0x0a, 0x0e, 0x63, 0x61, 0x73, 0x65, 0x5f, 0x73, 0x65, 0x6e, 0x73, 0x69, 0x74, 0x69,
	0x76, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d, 0x63, 0x61, 0x73, 0x65, 0x53, 0x65,
	0x6e, 0x73, 0x69, 0x74, 0x69, 0x76, 0x65, 0x12, 0x27, 0x0a, 0x0f, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x5f, 0x73, 0x65, 0x6e, 0x73, 0x69, 0x74, 0x69, 0x76, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x0e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x65, 0x6e, 0x73, 0x69, 0x74, 0x69, 0x76, 0x65,
	0x42, 0x2d, 0x5a, 0x2b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65,
	0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x67, 0x6f, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x65, 0x6f, 0x6c,
	0x79, 0x6d, 0x70, 0x2f, 0x61, 0x74, 0x6c, 0x61, 0x73, 0x3b, 0x61, 0x74, 0x6c, 0x61, 0x73, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_eolymp_atlas_testing_checker_proto_rawDescOnce sync.Once
	file_eolymp_atlas_testing_checker_proto_rawDescData []byte
)

func file_eolymp_atlas_testing_checker_proto_rawDescGZIP() []byte {
	file_eolymp_atlas_testing_checker_proto_rawDescOnce.Do(func() {
		file_eolymp_atlas_testing_checker_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_eolymp_atlas_testing_checker_proto_rawDesc), len(file_eolymp_atlas_testing_checker_proto_rawDesc)))
	})
	return file_eolymp_atlas_testing_checker_proto_rawDescData
}

var file_eolymp_atlas_testing_checker_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_eolymp_atlas_testing_checker_proto_goTypes = []any{
	(*Checker)(nil),            // 0: eolymp.atlas.Checker
	(executor.Checker_Type)(0), // 1: eolymp.executor.Checker.Type
	(*executor.File)(nil),      // 2: eolymp.executor.File
}
var file_eolymp_atlas_testing_checker_proto_depIdxs = []int32{
	1, // 0: eolymp.atlas.Checker.type:type_name -> eolymp.executor.Checker.Type
	2, // 1: eolymp.atlas.Checker.files:type_name -> eolymp.executor.File
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_eolymp_atlas_testing_checker_proto_init() }
func file_eolymp_atlas_testing_checker_proto_init() {
	if File_eolymp_atlas_testing_checker_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_eolymp_atlas_testing_checker_proto_rawDesc), len(file_eolymp_atlas_testing_checker_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_eolymp_atlas_testing_checker_proto_goTypes,
		DependencyIndexes: file_eolymp_atlas_testing_checker_proto_depIdxs,
		MessageInfos:      file_eolymp_atlas_testing_checker_proto_msgTypes,
	}.Build()
	File_eolymp_atlas_testing_checker_proto = out.File
	file_eolymp_atlas_testing_checker_proto_goTypes = nil
	file_eolymp_atlas_testing_checker_proto_depIdxs = nil
}
