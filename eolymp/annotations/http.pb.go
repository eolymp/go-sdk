// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v4.24.4
// source: eolymp/annotations/http.proto

package annotations

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
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

type Http struct {
	state                        protoimpl.MessageState `protogen:"open.v1"`
	Rules                        []*HttpRule            `protobuf:"bytes,1,rep,name=rules,proto3" json:"rules,omitempty"`
	FullyDecodeReservedExpansion bool                   `protobuf:"varint,2,opt,name=fully_decode_reserved_expansion,json=fullyDecodeReservedExpansion,proto3" json:"fully_decode_reserved_expansion,omitempty"`
	unknownFields                protoimpl.UnknownFields
	sizeCache                    protoimpl.SizeCache
}

func (x *Http) Reset() {
	*x = Http{}
	mi := &file_eolymp_annotations_http_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Http) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Http) ProtoMessage() {}

func (x *Http) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_annotations_http_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Http.ProtoReflect.Descriptor instead.
func (*Http) Descriptor() ([]byte, []int) {
	return file_eolymp_annotations_http_proto_rawDescGZIP(), []int{0}
}

func (x *Http) GetRules() []*HttpRule {
	if x != nil {
		return x.Rules
	}
	return nil
}

func (x *Http) GetFullyDecodeReservedExpansion() bool {
	if x != nil {
		return x.FullyDecodeReservedExpansion
	}
	return false
}

type HttpRule struct {
	state    protoimpl.MessageState `protogen:"open.v1"`
	Selector string                 `protobuf:"bytes,1,opt,name=selector,proto3" json:"selector,omitempty"`
	// Types that are valid to be assigned to Pattern:
	//
	//	*HttpRule_Get
	//	*HttpRule_Put
	//	*HttpRule_Post
	//	*HttpRule_Delete
	//	*HttpRule_Patch
	//	*HttpRule_Custom
	Pattern            isHttpRule_Pattern `protobuf_oneof:"pattern"`
	Body               string             `protobuf:"bytes,7,opt,name=body,proto3" json:"body,omitempty"`
	ResponseBody       string             `protobuf:"bytes,12,opt,name=response_body,json=responseBody,proto3" json:"response_body,omitempty"`
	AdditionalBindings []*HttpRule        `protobuf:"bytes,11,rep,name=additional_bindings,json=additionalBindings,proto3" json:"additional_bindings,omitempty"`
	unknownFields      protoimpl.UnknownFields
	sizeCache          protoimpl.SizeCache
}

func (x *HttpRule) Reset() {
	*x = HttpRule{}
	mi := &file_eolymp_annotations_http_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *HttpRule) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HttpRule) ProtoMessage() {}

func (x *HttpRule) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_annotations_http_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HttpRule.ProtoReflect.Descriptor instead.
func (*HttpRule) Descriptor() ([]byte, []int) {
	return file_eolymp_annotations_http_proto_rawDescGZIP(), []int{1}
}

func (x *HttpRule) GetSelector() string {
	if x != nil {
		return x.Selector
	}
	return ""
}

func (x *HttpRule) GetPattern() isHttpRule_Pattern {
	if x != nil {
		return x.Pattern
	}
	return nil
}

func (x *HttpRule) GetGet() string {
	if x != nil {
		if x, ok := x.Pattern.(*HttpRule_Get); ok {
			return x.Get
		}
	}
	return ""
}

func (x *HttpRule) GetPut() string {
	if x != nil {
		if x, ok := x.Pattern.(*HttpRule_Put); ok {
			return x.Put
		}
	}
	return ""
}

func (x *HttpRule) GetPost() string {
	if x != nil {
		if x, ok := x.Pattern.(*HttpRule_Post); ok {
			return x.Post
		}
	}
	return ""
}

func (x *HttpRule) GetDelete() string {
	if x != nil {
		if x, ok := x.Pattern.(*HttpRule_Delete); ok {
			return x.Delete
		}
	}
	return ""
}

func (x *HttpRule) GetPatch() string {
	if x != nil {
		if x, ok := x.Pattern.(*HttpRule_Patch); ok {
			return x.Patch
		}
	}
	return ""
}

func (x *HttpRule) GetCustom() *CustomHttpPattern {
	if x != nil {
		if x, ok := x.Pattern.(*HttpRule_Custom); ok {
			return x.Custom
		}
	}
	return nil
}

func (x *HttpRule) GetBody() string {
	if x != nil {
		return x.Body
	}
	return ""
}

func (x *HttpRule) GetResponseBody() string {
	if x != nil {
		return x.ResponseBody
	}
	return ""
}

func (x *HttpRule) GetAdditionalBindings() []*HttpRule {
	if x != nil {
		return x.AdditionalBindings
	}
	return nil
}

type isHttpRule_Pattern interface {
	isHttpRule_Pattern()
}

type HttpRule_Get struct {
	Get string `protobuf:"bytes,2,opt,name=get,proto3,oneof"`
}

type HttpRule_Put struct {
	Put string `protobuf:"bytes,3,opt,name=put,proto3,oneof"`
}

type HttpRule_Post struct {
	Post string `protobuf:"bytes,4,opt,name=post,proto3,oneof"`
}

type HttpRule_Delete struct {
	Delete string `protobuf:"bytes,5,opt,name=delete,proto3,oneof"`
}

type HttpRule_Patch struct {
	Patch string `protobuf:"bytes,6,opt,name=patch,proto3,oneof"`
}

type HttpRule_Custom struct {
	Custom *CustomHttpPattern `protobuf:"bytes,8,opt,name=custom,proto3,oneof"`
}

func (*HttpRule_Get) isHttpRule_Pattern() {}

func (*HttpRule_Put) isHttpRule_Pattern() {}

func (*HttpRule_Post) isHttpRule_Pattern() {}

func (*HttpRule_Delete) isHttpRule_Pattern() {}

func (*HttpRule_Patch) isHttpRule_Pattern() {}

func (*HttpRule_Custom) isHttpRule_Pattern() {}

type CustomHttpPattern struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Kind          string                 `protobuf:"bytes,1,opt,name=kind,proto3" json:"kind,omitempty"`
	Path          string                 `protobuf:"bytes,2,opt,name=path,proto3" json:"path,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CustomHttpPattern) Reset() {
	*x = CustomHttpPattern{}
	mi := &file_eolymp_annotations_http_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CustomHttpPattern) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CustomHttpPattern) ProtoMessage() {}

func (x *CustomHttpPattern) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_annotations_http_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CustomHttpPattern.ProtoReflect.Descriptor instead.
func (*CustomHttpPattern) Descriptor() ([]byte, []int) {
	return file_eolymp_annotations_http_proto_rawDescGZIP(), []int{2}
}

func (x *CustomHttpPattern) GetKind() string {
	if x != nil {
		return x.Kind
	}
	return ""
}

func (x *CustomHttpPattern) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

var file_eolymp_annotations_http_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptorpb.MethodOptions)(nil),
		ExtensionType: (*HttpRule)(nil),
		Field:         72295728,
		Name:          "eolymp.api.http",
		Tag:           "bytes,72295728,opt,name=http",
		Filename:      "eolymp/annotations/http.proto",
	},
}

// Extension fields to descriptorpb.MethodOptions.
var (
	// optional eolymp.api.HttpRule http = 72295728;
	E_Http = &file_eolymp_annotations_http_proto_extTypes[0] // keep extension number the same as one in google.api for compatibility
)

var File_eolymp_annotations_http_proto protoreflect.FileDescriptor

var file_eolymp_annotations_http_proto_rawDesc = string([]byte{
	0x0a, 0x1d, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0a, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x61, 0x70, 0x69, 0x1a, 0x20, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x79, 0x0a,
	0x04, 0x48, 0x74, 0x74, 0x70, 0x12, 0x2a, 0x0a, 0x05, 0x72, 0x75, 0x6c, 0x65, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x48, 0x74, 0x74, 0x70, 0x52, 0x75, 0x6c, 0x65, 0x52, 0x05, 0x72, 0x75, 0x6c, 0x65,
	0x73, 0x12, 0x45, 0x0a, 0x1f, 0x66, 0x75, 0x6c, 0x6c, 0x79, 0x5f, 0x64, 0x65, 0x63, 0x6f, 0x64,
	0x65, 0x5f, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x64, 0x5f, 0x65, 0x78, 0x70, 0x61, 0x6e,
	0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x1c, 0x66, 0x75, 0x6c, 0x6c,
	0x79, 0x44, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x64, 0x45,
	0x78, 0x70, 0x61, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0xda, 0x02, 0x0a, 0x08, 0x48, 0x74, 0x74,
	0x70, 0x52, 0x75, 0x6c, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f,
	0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f,
	0x72, 0x12, 0x12, 0x0a, 0x03, 0x67, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00,
	0x52, 0x03, 0x67, 0x65, 0x74, 0x12, 0x12, 0x0a, 0x03, 0x70, 0x75, 0x74, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x48, 0x00, 0x52, 0x03, 0x70, 0x75, 0x74, 0x12, 0x14, 0x0a, 0x04, 0x70, 0x6f, 0x73,
	0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x04, 0x70, 0x6f, 0x73, 0x74, 0x12,
	0x18, 0x0a, 0x06, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x48,
	0x00, 0x52, 0x06, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x16, 0x0a, 0x05, 0x70, 0x61, 0x74,
	0x63, 0x68, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x05, 0x70, 0x61, 0x74, 0x63,
	0x68, 0x12, 0x37, 0x0a, 0x06, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1d, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x43,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x48, 0x74, 0x74, 0x70, 0x50, 0x61, 0x74, 0x74, 0x65, 0x72, 0x6e,
	0x48, 0x00, 0x52, 0x06, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x6f,
	0x64, 0x79, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x12, 0x23,
	0x0a, 0x0d, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x62, 0x6f, 0x64, 0x79, 0x18,
	0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42,
	0x6f, 0x64, 0x79, 0x12, 0x45, 0x0a, 0x13, 0x61, 0x64, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x61,
	0x6c, 0x5f, 0x62, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x73, 0x18, 0x0b, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x14, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x48, 0x74,
	0x74, 0x70, 0x52, 0x75, 0x6c, 0x65, 0x52, 0x12, 0x61, 0x64, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x61, 0x6c, 0x42, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x73, 0x42, 0x09, 0x0a, 0x07, 0x70, 0x61,
	0x74, 0x74, 0x65, 0x72, 0x6e, 0x22, 0x3b, 0x0a, 0x11, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x48,
	0x74, 0x74, 0x70, 0x50, 0x61, 0x74, 0x74, 0x65, 0x72, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6b, 0x69,
	0x6e, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61,
	0x74, 0x68, 0x3a, 0x4b, 0x0a, 0x04, 0x68, 0x74, 0x74, 0x70, 0x12, 0x1e, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4d, 0x65, 0x74,
	0x68, 0x6f, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xb0, 0xca, 0xbc, 0x22, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x48, 0x74, 0x74, 0x70, 0x52, 0x75, 0x6c, 0x65, 0x52, 0x04, 0x68, 0x74, 0x74, 0x70, 0x42,
	0x39, 0x5a, 0x37, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6f,
	0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x67, 0x6f, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x65, 0x6f, 0x6c, 0x79,
	0x6d, 0x70, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x3b, 0x61,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
})

var (
	file_eolymp_annotations_http_proto_rawDescOnce sync.Once
	file_eolymp_annotations_http_proto_rawDescData []byte
)

func file_eolymp_annotations_http_proto_rawDescGZIP() []byte {
	file_eolymp_annotations_http_proto_rawDescOnce.Do(func() {
		file_eolymp_annotations_http_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_eolymp_annotations_http_proto_rawDesc), len(file_eolymp_annotations_http_proto_rawDesc)))
	})
	return file_eolymp_annotations_http_proto_rawDescData
}

var file_eolymp_annotations_http_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_eolymp_annotations_http_proto_goTypes = []any{
	(*Http)(nil),                       // 0: eolymp.api.Http
	(*HttpRule)(nil),                   // 1: eolymp.api.HttpRule
	(*CustomHttpPattern)(nil),          // 2: eolymp.api.CustomHttpPattern
	(*descriptorpb.MethodOptions)(nil), // 3: google.protobuf.MethodOptions
}
var file_eolymp_annotations_http_proto_depIdxs = []int32{
	1, // 0: eolymp.api.Http.rules:type_name -> eolymp.api.HttpRule
	2, // 1: eolymp.api.HttpRule.custom:type_name -> eolymp.api.CustomHttpPattern
	1, // 2: eolymp.api.HttpRule.additional_bindings:type_name -> eolymp.api.HttpRule
	3, // 3: eolymp.api.http:extendee -> google.protobuf.MethodOptions
	1, // 4: eolymp.api.http:type_name -> eolymp.api.HttpRule
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	4, // [4:5] is the sub-list for extension type_name
	3, // [3:4] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_eolymp_annotations_http_proto_init() }
func file_eolymp_annotations_http_proto_init() {
	if File_eolymp_annotations_http_proto != nil {
		return
	}
	file_eolymp_annotations_http_proto_msgTypes[1].OneofWrappers = []any{
		(*HttpRule_Get)(nil),
		(*HttpRule_Put)(nil),
		(*HttpRule_Post)(nil),
		(*HttpRule_Delete)(nil),
		(*HttpRule_Patch)(nil),
		(*HttpRule_Custom)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_eolymp_annotations_http_proto_rawDesc), len(file_eolymp_annotations_http_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 1,
			NumServices:   0,
		},
		GoTypes:           file_eolymp_annotations_http_proto_goTypes,
		DependencyIndexes: file_eolymp_annotations_http_proto_depIdxs,
		MessageInfos:      file_eolymp_annotations_http_proto_msgTypes,
		ExtensionInfos:    file_eolymp_annotations_http_proto_extTypes,
	}.Build()
	File_eolymp_annotations_http_proto = out.File
	file_eolymp_annotations_http_proto_goTypes = nil
	file_eolymp_annotations_http_proto_depIdxs = nil
}
