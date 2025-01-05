// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.1
// 	protoc        v4.24.4
// source: eolymp/content/print_service.proto

package content

import (
	_ "github.com/eolymp/go-sdk/eolymp/annotations"
	ecm "github.com/eolymp/go-sdk/eolymp/ecm"
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

type PrintContentInput struct {
	state           protoimpl.MessageState `protogen:"open.v1"`
	SinglePage      bool                   `protobuf:"varint,90,opt,name=single_page,json=singlePage,proto3" json:"single_page,omitempty"`
	HeaderPath      string                 `protobuf:"bytes,10,opt,name=header_path,json=headerPath,proto3" json:"header_path,omitempty"`
	FooterPath      string                 `protobuf:"bytes,11,opt,name=footer_path,json=footerPath,proto3" json:"footer_path,omitempty"`
	PrintBackground bool                   `protobuf:"varint,20,opt,name=print_background,json=printBackground,proto3" json:"print_background,omitempty"`
	PaperWidth      uint32                 `protobuf:"varint,30,opt,name=paper_width,json=paperWidth,proto3" json:"paper_width,omitempty"`    // in cm, A4 by default
	PaperHeight     uint32                 `protobuf:"varint,31,opt,name=paper_height,json=paperHeight,proto3" json:"paper_height,omitempty"` // in cm, A4 by default
	Margins         []uint32               `protobuf:"varint,35,rep,packed,name=margins,proto3" json:"margins,omitempty"`                     // one value for all sides, two values for vertical and horizontal, four values for top, right, bottom, left
	Landscape       bool                   `protobuf:"varint,39,opt,name=landscape,proto3" json:"landscape,omitempty"`
	Content         *ecm.Content           `protobuf:"bytes,1,opt,name=content,proto3" json:"content,omitempty"` // document to print
	unknownFields   protoimpl.UnknownFields
	sizeCache       protoimpl.SizeCache
}

func (x *PrintContentInput) Reset() {
	*x = PrintContentInput{}
	mi := &file_eolymp_content_print_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PrintContentInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PrintContentInput) ProtoMessage() {}

func (x *PrintContentInput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_content_print_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PrintContentInput.ProtoReflect.Descriptor instead.
func (*PrintContentInput) Descriptor() ([]byte, []int) {
	return file_eolymp_content_print_service_proto_rawDescGZIP(), []int{0}
}

func (x *PrintContentInput) GetSinglePage() bool {
	if x != nil {
		return x.SinglePage
	}
	return false
}

func (x *PrintContentInput) GetHeaderPath() string {
	if x != nil {
		return x.HeaderPath
	}
	return ""
}

func (x *PrintContentInput) GetFooterPath() string {
	if x != nil {
		return x.FooterPath
	}
	return ""
}

func (x *PrintContentInput) GetPrintBackground() bool {
	if x != nil {
		return x.PrintBackground
	}
	return false
}

func (x *PrintContentInput) GetPaperWidth() uint32 {
	if x != nil {
		return x.PaperWidth
	}
	return 0
}

func (x *PrintContentInput) GetPaperHeight() uint32 {
	if x != nil {
		return x.PaperHeight
	}
	return 0
}

func (x *PrintContentInput) GetMargins() []uint32 {
	if x != nil {
		return x.Margins
	}
	return nil
}

func (x *PrintContentInput) GetLandscape() bool {
	if x != nil {
		return x.Landscape
	}
	return false
}

func (x *PrintContentInput) GetContent() *ecm.Content {
	if x != nil {
		return x.Content
	}
	return nil
}

type PrintContentOutput struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	DocumentUrl   string                 `protobuf:"bytes,1,opt,name=document_url,json=documentUrl,proto3" json:"document_url,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PrintContentOutput) Reset() {
	*x = PrintContentOutput{}
	mi := &file_eolymp_content_print_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PrintContentOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PrintContentOutput) ProtoMessage() {}

func (x *PrintContentOutput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_content_print_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PrintContentOutput.ProtoReflect.Descriptor instead.
func (*PrintContentOutput) Descriptor() ([]byte, []int) {
	return file_eolymp_content_print_service_proto_rawDescGZIP(), []int{1}
}

func (x *PrintContentOutput) GetDocumentUrl() string {
	if x != nil {
		return x.DocumentUrl
	}
	return ""
}

var File_eolymp_content_print_service_proto protoreflect.FileDescriptor

var file_eolymp_content_print_service_proto_rawDesc = []byte{
	0x0a, 0x22, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x2f, 0x70, 0x72, 0x69, 0x6e, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x63, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x1a, 0x1d, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x22, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x72, 0x61, 0x74, 0x65, 0x6c, 0x69, 0x6d, 0x69,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x18, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f,
	0x65, 0x63, 0x6d, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x15, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x65, 0x63, 0x6d, 0x2f, 0x6e, 0x6f,
	0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xcc, 0x02, 0x0a, 0x11, 0x50, 0x72, 0x69,
	0x6e, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x1f,
	0x0a, 0x0b, 0x73, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x18, 0x5a, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x0a, 0x73, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x50, 0x61, 0x67, 0x65, 0x12,
	0x1f, 0x0a, 0x0b, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x0a,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x50, 0x61, 0x74, 0x68,
	0x12, 0x1f, 0x0a, 0x0b, 0x66, 0x6f, 0x6f, 0x74, 0x65, 0x72, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18,
	0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x66, 0x6f, 0x6f, 0x74, 0x65, 0x72, 0x50, 0x61, 0x74,
	0x68, 0x12, 0x29, 0x0a, 0x10, 0x70, 0x72, 0x69, 0x6e, 0x74, 0x5f, 0x62, 0x61, 0x63, 0x6b, 0x67,
	0x72, 0x6f, 0x75, 0x6e, 0x64, 0x18, 0x14, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0f, 0x70, 0x72, 0x69,
	0x6e, 0x74, 0x42, 0x61, 0x63, 0x6b, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x12, 0x1f, 0x0a, 0x0b,
	0x70, 0x61, 0x70, 0x65, 0x72, 0x5f, 0x77, 0x69, 0x64, 0x74, 0x68, 0x18, 0x1e, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x0a, 0x70, 0x61, 0x70, 0x65, 0x72, 0x57, 0x69, 0x64, 0x74, 0x68, 0x12, 0x21, 0x0a,
	0x0c, 0x70, 0x61, 0x70, 0x65, 0x72, 0x5f, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x1f, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x0b, 0x70, 0x61, 0x70, 0x65, 0x72, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74,
	0x12, 0x18, 0x0a, 0x07, 0x6d, 0x61, 0x72, 0x67, 0x69, 0x6e, 0x73, 0x18, 0x23, 0x20, 0x03, 0x28,
	0x0d, 0x52, 0x07, 0x6d, 0x61, 0x72, 0x67, 0x69, 0x6e, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x6c, 0x61,
	0x6e, 0x64, 0x73, 0x63, 0x61, 0x70, 0x65, 0x18, 0x27, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x6c,
	0x61, 0x6e, 0x64, 0x73, 0x63, 0x61, 0x70, 0x65, 0x12, 0x2d, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x65, 0x6f, 0x6c, 0x79,
	0x6d, 0x70, 0x2e, 0x65, 0x63, 0x6d, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x52, 0x07,
	0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0x37, 0x0a, 0x12, 0x50, 0x72, 0x69, 0x6e, 0x74,
	0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x12, 0x21, 0x0a,
	0x0c, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x55, 0x72, 0x6c,
	0x32, 0x84, 0x01, 0x0a, 0x0c, 0x50, 0x72, 0x69, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x74, 0x0a, 0x0c, 0x50, 0x72, 0x69, 0x6e, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x12, 0x21, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x2e, 0x50, 0x72, 0x69, 0x6e, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x49,
	0x6e, 0x70, 0x75, 0x74, 0x1a, 0x22, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x2e, 0x50, 0x72, 0x69, 0x6e, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x22, 0x1d, 0xea, 0xe2, 0x0a, 0x0b, 0xf5, 0xe2,
	0x0a, 0x00, 0x00, 0x00, 0x40, 0xf8, 0xe2, 0x0a, 0x0a, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x08, 0x22,
	0x06, 0x2f, 0x70, 0x72, 0x69, 0x6e, 0x74, 0x42, 0x31, 0x5a, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x67, 0x6f, 0x2d,
	0x73, 0x64, 0x6b, 0x2f, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x3b, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_eolymp_content_print_service_proto_rawDescOnce sync.Once
	file_eolymp_content_print_service_proto_rawDescData = file_eolymp_content_print_service_proto_rawDesc
)

func file_eolymp_content_print_service_proto_rawDescGZIP() []byte {
	file_eolymp_content_print_service_proto_rawDescOnce.Do(func() {
		file_eolymp_content_print_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_eolymp_content_print_service_proto_rawDescData)
	})
	return file_eolymp_content_print_service_proto_rawDescData
}

var file_eolymp_content_print_service_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_eolymp_content_print_service_proto_goTypes = []any{
	(*PrintContentInput)(nil),  // 0: eolymp.content.PrintContentInput
	(*PrintContentOutput)(nil), // 1: eolymp.content.PrintContentOutput
	(*ecm.Content)(nil),        // 2: eolymp.ecm.Content
}
var file_eolymp_content_print_service_proto_depIdxs = []int32{
	2, // 0: eolymp.content.PrintContentInput.content:type_name -> eolymp.ecm.Content
	0, // 1: eolymp.content.PrintService.PrintContent:input_type -> eolymp.content.PrintContentInput
	1, // 2: eolymp.content.PrintService.PrintContent:output_type -> eolymp.content.PrintContentOutput
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_eolymp_content_print_service_proto_init() }
func file_eolymp_content_print_service_proto_init() {
	if File_eolymp_content_print_service_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_eolymp_content_print_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_eolymp_content_print_service_proto_goTypes,
		DependencyIndexes: file_eolymp_content_print_service_proto_depIdxs,
		MessageInfos:      file_eolymp_content_print_service_proto_msgTypes,
	}.Build()
	File_eolymp_content_print_service_proto = out.File
	file_eolymp_content_print_service_proto_rawDesc = nil
	file_eolymp_content_print_service_proto_goTypes = nil
	file_eolymp_content_print_service_proto_depIdxs = nil
}
