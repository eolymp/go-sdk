// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v3.21.12
// source: eolymp/executor/interactor.proto

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

type Interactor_Type int32

const (
	Interactor_NONE    Interactor_Type = 0
	Interactor_PROGRAM Interactor_Type = 1
)

// Enum value maps for Interactor_Type.
var (
	Interactor_Type_name = map[int32]string{
		0: "NONE",
		1: "PROGRAM",
	}
	Interactor_Type_value = map[string]int32{
		"NONE":    0,
		"PROGRAM": 1,
	}
)

func (x Interactor_Type) Enum() *Interactor_Type {
	p := new(Interactor_Type)
	*p = x
	return p
}

func (x Interactor_Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Interactor_Type) Descriptor() protoreflect.EnumDescriptor {
	return file_eolymp_executor_interactor_proto_enumTypes[0].Descriptor()
}

func (Interactor_Type) Type() protoreflect.EnumType {
	return &file_eolymp_executor_interactor_proto_enumTypes[0]
}

func (x Interactor_Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Interactor_Type.Descriptor instead.
func (Interactor_Type) EnumDescriptor() ([]byte, []int) {
	return file_eolymp_executor_interactor_proto_rawDescGZIP(), []int{0, 0}
}

// Interactor provides configuration for program which would interact with task during execution
type Interactor struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Interactor type (see types enumeration for details)
	Type Interactor_Type `protobuf:"varint,1,opt,name=type,proto3,enum=eolymp.executor.Interactor_Type" json:"type,omitempty"`
	// Programming language in which interactor is written
	Lang string `protobuf:"bytes,2,opt,name=lang,proto3" json:"lang,omitempty"`
	// deprecated
	Source string `protobuf:"bytes,3,opt,name=source,proto3" json:"source,omitempty"`
	// Source code for interactor
	SourceUrl string `protobuf:"bytes,8,opt,name=source_url,json=sourceUrl,proto3" json:"source_url,omitempty"`
	// Secret means checker code and configuration must not be exposed to users
	Secret bool `protobuf:"varint,7,opt,name=secret,proto3" json:"secret,omitempty"`
	// Additional files placed into workdir during compilation and execution
	Files []*Interactor_File `protobuf:"bytes,10,rep,name=files,proto3" json:"files,omitempty"`
}

func (x *Interactor) Reset() {
	*x = Interactor{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_executor_interactor_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Interactor) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Interactor) ProtoMessage() {}

func (x *Interactor) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_executor_interactor_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Interactor.ProtoReflect.Descriptor instead.
func (*Interactor) Descriptor() ([]byte, []int) {
	return file_eolymp_executor_interactor_proto_rawDescGZIP(), []int{0}
}

func (x *Interactor) GetType() Interactor_Type {
	if x != nil {
		return x.Type
	}
	return Interactor_NONE
}

func (x *Interactor) GetLang() string {
	if x != nil {
		return x.Lang
	}
	return ""
}

func (x *Interactor) GetSource() string {
	if x != nil {
		return x.Source
	}
	return ""
}

func (x *Interactor) GetSourceUrl() string {
	if x != nil {
		return x.SourceUrl
	}
	return ""
}

func (x *Interactor) GetSecret() bool {
	if x != nil {
		return x.Secret
	}
	return false
}

func (x *Interactor) GetFiles() []*Interactor_File {
	if x != nil {
		return x.Files
	}
	return nil
}

// File defines additional file which might be placed into work directory during compilation or execution
type Interactor_File struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Path      string `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"` // Path where file should be placed (always relative to the workdir)
	SourceUrl string `protobuf:"bytes,3,opt,name=source_url,json=sourceUrl,proto3" json:"source_url,omitempty"`
}

func (x *Interactor_File) Reset() {
	*x = Interactor_File{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_executor_interactor_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Interactor_File) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Interactor_File) ProtoMessage() {}

func (x *Interactor_File) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_executor_interactor_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Interactor_File.ProtoReflect.Descriptor instead.
func (*Interactor_File) Descriptor() ([]byte, []int) {
	return file_eolymp_executor_interactor_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Interactor_File) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *Interactor_File) GetSourceUrl() string {
	if x != nil {
		return x.SourceUrl
	}
	return ""
}

var File_eolymp_executor_interactor_proto protoreflect.FileDescriptor

var file_eolymp_executor_interactor_proto_rawDesc = []byte{
	0x0a, 0x20, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x6f,
	0x72, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0f, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x65, 0x78, 0x65, 0x63, 0x75,
	0x74, 0x6f, 0x72, 0x22, 0xb7, 0x02, 0x0a, 0x0a, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74,
	0x6f, 0x72, 0x12, 0x34, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x20, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74,
	0x6f, 0x72, 0x2e, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x54, 0x79,
	0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x61, 0x6e, 0x67,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6c, 0x61, 0x6e, 0x67, 0x12, 0x16, 0x0a, 0x06,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x75,
	0x72, 0x6c, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x55, 0x72, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x06, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x12, 0x36, 0x0a, 0x05, 0x66,
	0x69, 0x6c, 0x65, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x65, 0x6f, 0x6c,
	0x79, 0x6d, 0x70, 0x2e, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x6f, 0x72, 0x2e, 0x49, 0x6e, 0x74,
	0x65, 0x72, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x05, 0x66, 0x69,
	0x6c, 0x65, 0x73, 0x1a, 0x39, 0x0a, 0x04, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x70,
	0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12,
	0x1d, 0x0a, 0x0a, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x55, 0x72, 0x6c, 0x22, 0x1d,
	0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00,
	0x12, 0x0b, 0x0a, 0x07, 0x50, 0x52, 0x4f, 0x47, 0x52, 0x41, 0x4d, 0x10, 0x01, 0x42, 0x33, 0x5a,
	0x31, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6f, 0x6c, 0x79,
	0x6d, 0x70, 0x2f, 0x67, 0x6f, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70,
	0x2f, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x6f, 0x72, 0x3b, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74,
	0x6f, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_eolymp_executor_interactor_proto_rawDescOnce sync.Once
	file_eolymp_executor_interactor_proto_rawDescData = file_eolymp_executor_interactor_proto_rawDesc
)

func file_eolymp_executor_interactor_proto_rawDescGZIP() []byte {
	file_eolymp_executor_interactor_proto_rawDescOnce.Do(func() {
		file_eolymp_executor_interactor_proto_rawDescData = protoimpl.X.CompressGZIP(file_eolymp_executor_interactor_proto_rawDescData)
	})
	return file_eolymp_executor_interactor_proto_rawDescData
}

var file_eolymp_executor_interactor_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_eolymp_executor_interactor_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_eolymp_executor_interactor_proto_goTypes = []interface{}{
	(Interactor_Type)(0),    // 0: eolymp.executor.Interactor.Type
	(*Interactor)(nil),      // 1: eolymp.executor.Interactor
	(*Interactor_File)(nil), // 2: eolymp.executor.Interactor.File
}
var file_eolymp_executor_interactor_proto_depIdxs = []int32{
	0, // 0: eolymp.executor.Interactor.type:type_name -> eolymp.executor.Interactor.Type
	2, // 1: eolymp.executor.Interactor.files:type_name -> eolymp.executor.Interactor.File
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_eolymp_executor_interactor_proto_init() }
func file_eolymp_executor_interactor_proto_init() {
	if File_eolymp_executor_interactor_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_eolymp_executor_interactor_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Interactor); i {
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
		file_eolymp_executor_interactor_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Interactor_File); i {
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
			RawDescriptor: file_eolymp_executor_interactor_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_eolymp_executor_interactor_proto_goTypes,
		DependencyIndexes: file_eolymp_executor_interactor_proto_depIdxs,
		EnumInfos:         file_eolymp_executor_interactor_proto_enumTypes,
		MessageInfos:      file_eolymp_executor_interactor_proto_msgTypes,
	}.Build()
	File_eolymp_executor_interactor_proto = out.File
	file_eolymp_executor_interactor_proto_rawDesc = nil
	file_eolymp_executor_interactor_proto_goTypes = nil
	file_eolymp_executor_interactor_proto_depIdxs = nil
}
