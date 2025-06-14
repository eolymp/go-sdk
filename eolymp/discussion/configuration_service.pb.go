// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.29.4
// source: eolymp/discussion/configuration_service.proto

package discussion

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

type UpdateDiscussionConfigInput_Patch int32

const (
	UpdateDiscussionConfigInput_ALL                             UpdateDiscussionConfigInput_Patch = 0
	UpdateDiscussionConfigInput_MEMBERS_CAN_CREATE_POSTS        UpdateDiscussionConfigInput_Patch = 1
	UpdateDiscussionConfigInput_MEMBERS_CAN_COMMENT_ON_POSTS    UpdateDiscussionConfigInput_Patch = 2
	UpdateDiscussionConfigInput_MEMBERS_CAN_COMMENT_ON_PROBLEMS UpdateDiscussionConfigInput_Patch = 3
	UpdateDiscussionConfigInput_POST_MODERATION                 UpdateDiscussionConfigInput_Patch = 4
	UpdateDiscussionConfigInput_COMMENT_MODERATION              UpdateDiscussionConfigInput_Patch = 5
)

// Enum value maps for UpdateDiscussionConfigInput_Patch.
var (
	UpdateDiscussionConfigInput_Patch_name = map[int32]string{
		0: "ALL",
		1: "MEMBERS_CAN_CREATE_POSTS",
		2: "MEMBERS_CAN_COMMENT_ON_POSTS",
		3: "MEMBERS_CAN_COMMENT_ON_PROBLEMS",
		4: "POST_MODERATION",
		5: "COMMENT_MODERATION",
	}
	UpdateDiscussionConfigInput_Patch_value = map[string]int32{
		"ALL":                             0,
		"MEMBERS_CAN_CREATE_POSTS":        1,
		"MEMBERS_CAN_COMMENT_ON_POSTS":    2,
		"MEMBERS_CAN_COMMENT_ON_PROBLEMS": 3,
		"POST_MODERATION":                 4,
		"COMMENT_MODERATION":              5,
	}
)

func (x UpdateDiscussionConfigInput_Patch) Enum() *UpdateDiscussionConfigInput_Patch {
	p := new(UpdateDiscussionConfigInput_Patch)
	*p = x
	return p
}

func (x UpdateDiscussionConfigInput_Patch) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (UpdateDiscussionConfigInput_Patch) Descriptor() protoreflect.EnumDescriptor {
	return file_eolymp_discussion_configuration_service_proto_enumTypes[0].Descriptor()
}

func (UpdateDiscussionConfigInput_Patch) Type() protoreflect.EnumType {
	return &file_eolymp_discussion_configuration_service_proto_enumTypes[0]
}

func (x UpdateDiscussionConfigInput_Patch) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use UpdateDiscussionConfigInput_Patch.Descriptor instead.
func (UpdateDiscussionConfigInput_Patch) EnumDescriptor() ([]byte, []int) {
	return file_eolymp_discussion_configuration_service_proto_rawDescGZIP(), []int{2, 0}
}

type DescribeDiscussionConfigInput struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DescribeDiscussionConfigInput) Reset() {
	*x = DescribeDiscussionConfigInput{}
	mi := &file_eolymp_discussion_configuration_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DescribeDiscussionConfigInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescribeDiscussionConfigInput) ProtoMessage() {}

func (x *DescribeDiscussionConfigInput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_discussion_configuration_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DescribeDiscussionConfigInput.ProtoReflect.Descriptor instead.
func (*DescribeDiscussionConfigInput) Descriptor() ([]byte, []int) {
	return file_eolymp_discussion_configuration_service_proto_rawDescGZIP(), []int{0}
}

type DescribeDiscussionConfigOutput struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Config        *Configuration         `protobuf:"bytes,1,opt,name=config,proto3" json:"config,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DescribeDiscussionConfigOutput) Reset() {
	*x = DescribeDiscussionConfigOutput{}
	mi := &file_eolymp_discussion_configuration_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DescribeDiscussionConfigOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescribeDiscussionConfigOutput) ProtoMessage() {}

func (x *DescribeDiscussionConfigOutput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_discussion_configuration_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DescribeDiscussionConfigOutput.ProtoReflect.Descriptor instead.
func (*DescribeDiscussionConfigOutput) Descriptor() ([]byte, []int) {
	return file_eolymp_discussion_configuration_service_proto_rawDescGZIP(), []int{1}
}

func (x *DescribeDiscussionConfigOutput) GetConfig() *Configuration {
	if x != nil {
		return x.Config
	}
	return nil
}

type UpdateDiscussionConfigInput struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Config        *Configuration         `protobuf:"bytes,1,opt,name=config,proto3" json:"config,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateDiscussionConfigInput) Reset() {
	*x = UpdateDiscussionConfigInput{}
	mi := &file_eolymp_discussion_configuration_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateDiscussionConfigInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateDiscussionConfigInput) ProtoMessage() {}

func (x *UpdateDiscussionConfigInput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_discussion_configuration_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateDiscussionConfigInput.ProtoReflect.Descriptor instead.
func (*UpdateDiscussionConfigInput) Descriptor() ([]byte, []int) {
	return file_eolymp_discussion_configuration_service_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateDiscussionConfigInput) GetConfig() *Configuration {
	if x != nil {
		return x.Config
	}
	return nil
}

type UpdateDiscussionConfigOutput struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateDiscussionConfigOutput) Reset() {
	*x = UpdateDiscussionConfigOutput{}
	mi := &file_eolymp_discussion_configuration_service_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateDiscussionConfigOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateDiscussionConfigOutput) ProtoMessage() {}

func (x *UpdateDiscussionConfigOutput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_discussion_configuration_service_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateDiscussionConfigOutput.ProtoReflect.Descriptor instead.
func (*UpdateDiscussionConfigOutput) Descriptor() ([]byte, []int) {
	return file_eolymp_discussion_configuration_service_proto_rawDescGZIP(), []int{3}
}

var File_eolymp_discussion_configuration_service_proto protoreflect.FileDescriptor

const file_eolymp_discussion_configuration_service_proto_rawDesc = "" +
	"\n" +
	"-eolymp/discussion/configuration_service.proto\x12\x11eolymp.discussion\x1a\x1deolymp/annotations/http.proto\x1a\"eolymp/annotations/ratelimit.proto\x1a%eolymp/discussion/configuration.proto\"\x1f\n" +
	"\x1dDescribeDiscussionConfigInput\"Z\n" +
	"\x1eDescribeDiscussionConfigOutput\x128\n" +
	"\x06config\x18\x01 \x01(\v2 .eolymp.discussion.ConfigurationR\x06config\"\xfc\x01\n" +
	"\x1bUpdateDiscussionConfigInput\x128\n" +
	"\x06config\x18\x01 \x01(\v2 .eolymp.discussion.ConfigurationR\x06config\"\xa2\x01\n" +
	"\x05Patch\x12\a\n" +
	"\x03ALL\x10\x00\x12\x1c\n" +
	"\x18MEMBERS_CAN_CREATE_POSTS\x10\x01\x12 \n" +
	"\x1cMEMBERS_CAN_COMMENT_ON_POSTS\x10\x02\x12#\n" +
	"\x1fMEMBERS_CAN_COMMENT_ON_PROBLEMS\x10\x03\x12\x13\n" +
	"\x0fPOST_MODERATION\x10\x04\x12\x16\n" +
	"\x12COMMENT_MODERATION\x10\x05\"\x1e\n" +
	"\x1cUpdateDiscussionConfigOutput2\xee\x02\n" +
	"\x14ConfigurationService\x12\xac\x01\n" +
	"\x18DescribeDiscussionConfig\x120.eolymp.discussion.DescribeDiscussionConfigInput\x1a1.eolymp.discussion.DescribeDiscussionConfigOutput\"+\xea\xe2\n" +
	"\v\xf5\xe2\n" +
	"\x00\x00 A\xf8\xe2\n" +
	"d\x82\xd3\xe4\x93\x02\x16\x12\x14/configs/discussions\x12\xa6\x01\n" +
	"\x16UpdateDiscussionConfig\x12..eolymp.discussion.UpdateDiscussionConfigInput\x1a/.eolymp.discussion.UpdateDiscussionConfigOutput\"+\xea\xe2\n" +
	"\v\xf5\xe2\n" +
	"\x00\x00\x80?\xf8\xe2\n" +
	"\x05\x82\xd3\xe4\x93\x02\x16\x1a\x14/configs/discussionsB7Z5github.com/eolymp/go-sdk/eolymp/discussion;discussionb\x06proto3"

var (
	file_eolymp_discussion_configuration_service_proto_rawDescOnce sync.Once
	file_eolymp_discussion_configuration_service_proto_rawDescData []byte
)

func file_eolymp_discussion_configuration_service_proto_rawDescGZIP() []byte {
	file_eolymp_discussion_configuration_service_proto_rawDescOnce.Do(func() {
		file_eolymp_discussion_configuration_service_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_eolymp_discussion_configuration_service_proto_rawDesc), len(file_eolymp_discussion_configuration_service_proto_rawDesc)))
	})
	return file_eolymp_discussion_configuration_service_proto_rawDescData
}

var file_eolymp_discussion_configuration_service_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_eolymp_discussion_configuration_service_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_eolymp_discussion_configuration_service_proto_goTypes = []any{
	(UpdateDiscussionConfigInput_Patch)(0), // 0: eolymp.discussion.UpdateDiscussionConfigInput.Patch
	(*DescribeDiscussionConfigInput)(nil),  // 1: eolymp.discussion.DescribeDiscussionConfigInput
	(*DescribeDiscussionConfigOutput)(nil), // 2: eolymp.discussion.DescribeDiscussionConfigOutput
	(*UpdateDiscussionConfigInput)(nil),    // 3: eolymp.discussion.UpdateDiscussionConfigInput
	(*UpdateDiscussionConfigOutput)(nil),   // 4: eolymp.discussion.UpdateDiscussionConfigOutput
	(*Configuration)(nil),                  // 5: eolymp.discussion.Configuration
}
var file_eolymp_discussion_configuration_service_proto_depIdxs = []int32{
	5, // 0: eolymp.discussion.DescribeDiscussionConfigOutput.config:type_name -> eolymp.discussion.Configuration
	5, // 1: eolymp.discussion.UpdateDiscussionConfigInput.config:type_name -> eolymp.discussion.Configuration
	1, // 2: eolymp.discussion.ConfigurationService.DescribeDiscussionConfig:input_type -> eolymp.discussion.DescribeDiscussionConfigInput
	3, // 3: eolymp.discussion.ConfigurationService.UpdateDiscussionConfig:input_type -> eolymp.discussion.UpdateDiscussionConfigInput
	2, // 4: eolymp.discussion.ConfigurationService.DescribeDiscussionConfig:output_type -> eolymp.discussion.DescribeDiscussionConfigOutput
	4, // 5: eolymp.discussion.ConfigurationService.UpdateDiscussionConfig:output_type -> eolymp.discussion.UpdateDiscussionConfigOutput
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_eolymp_discussion_configuration_service_proto_init() }
func file_eolymp_discussion_configuration_service_proto_init() {
	if File_eolymp_discussion_configuration_service_proto != nil {
		return
	}
	file_eolymp_discussion_configuration_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_eolymp_discussion_configuration_service_proto_rawDesc), len(file_eolymp_discussion_configuration_service_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_eolymp_discussion_configuration_service_proto_goTypes,
		DependencyIndexes: file_eolymp_discussion_configuration_service_proto_depIdxs,
		EnumInfos:         file_eolymp_discussion_configuration_service_proto_enumTypes,
		MessageInfos:      file_eolymp_discussion_configuration_service_proto_msgTypes,
	}.Build()
	File_eolymp_discussion_configuration_service_proto = out.File
	file_eolymp_discussion_configuration_service_proto_goTypes = nil
	file_eolymp_discussion_configuration_service_proto_depIdxs = nil
}
