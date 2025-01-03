// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.1
// 	protoc        v4.24.4
// source: eolymp/atlas/testing_test.proto

package atlas

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

type Test struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Test unique identifier.
	Id                 string  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	TestsetId          string  `protobuf:"bytes,2,opt,name=testset_id,json=testsetId,proto3" json:"testset_id,omitempty"`
	Index              int32   `protobuf:"varint,3,opt,name=index,proto3" json:"index,omitempty"`                                                       // Index, a sequential number, of the test within testset.
	Example            bool    `protobuf:"varint,4,opt,name=example,proto3" json:"example,omitempty"`                                                   // Flag which defines if test should be shown in as an example in the problem statement
	Inactive           bool    `protobuf:"varint,7,opt,name=inactive,proto3" json:"inactive,omitempty"`                                                 // The test is inactive
	Secret             bool    `protobuf:"varint,6,opt,name=secret,proto3" json:"secret,omitempty"`                                                     // Secret test, input and answer are never populated
	Score              float32 `protobuf:"fixed32,5,opt,name=score,proto3" json:"score,omitempty"`                                                      // Score for passing this test
	ExampleInputUrl    string  `protobuf:"bytes,40,opt,name=example_input_url,json=exampleInputUrl,proto3" json:"example_input_url,omitempty"`          // Optionally, override input data for example in statement
	ExampleAnswerUrl   string  `protobuf:"bytes,41,opt,name=example_answer_url,json=exampleAnswerUrl,proto3" json:"example_answer_url,omitempty"`       // Optionally, override answer data for example in statement
	GeneratedInputUrl  string  `protobuf:"bytes,50,opt,name=generated_input_url,json=generatedInputUrl,proto3" json:"generated_input_url,omitempty"`    // Generated input data
	GeneratedAnswerUrl string  `protobuf:"bytes,51,opt,name=generated_answer_url,json=generatedAnswerUrl,proto3" json:"generated_answer_url,omitempty"` // Generated answer data
	// Types that are valid to be assigned to Input:
	//
	//	*Test_InputUrl
	//	*Test_InputGenerator
	Input isTest_Input `protobuf_oneof:"input"`
	// Types that are valid to be assigned to Answer:
	//
	//	*Test_AnswerUrl
	//	*Test_AnswerGenerator
	Answer        isTest_Answer `protobuf_oneof:"answer"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Test) Reset() {
	*x = Test{}
	mi := &file_eolymp_atlas_testing_test_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Test) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Test) ProtoMessage() {}

func (x *Test) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_atlas_testing_test_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Test.ProtoReflect.Descriptor instead.
func (*Test) Descriptor() ([]byte, []int) {
	return file_eolymp_atlas_testing_test_proto_rawDescGZIP(), []int{0}
}

func (x *Test) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Test) GetTestsetId() string {
	if x != nil {
		return x.TestsetId
	}
	return ""
}

func (x *Test) GetIndex() int32 {
	if x != nil {
		return x.Index
	}
	return 0
}

func (x *Test) GetExample() bool {
	if x != nil {
		return x.Example
	}
	return false
}

func (x *Test) GetInactive() bool {
	if x != nil {
		return x.Inactive
	}
	return false
}

func (x *Test) GetSecret() bool {
	if x != nil {
		return x.Secret
	}
	return false
}

func (x *Test) GetScore() float32 {
	if x != nil {
		return x.Score
	}
	return 0
}

func (x *Test) GetExampleInputUrl() string {
	if x != nil {
		return x.ExampleInputUrl
	}
	return ""
}

func (x *Test) GetExampleAnswerUrl() string {
	if x != nil {
		return x.ExampleAnswerUrl
	}
	return ""
}

func (x *Test) GetGeneratedInputUrl() string {
	if x != nil {
		return x.GeneratedInputUrl
	}
	return ""
}

func (x *Test) GetGeneratedAnswerUrl() string {
	if x != nil {
		return x.GeneratedAnswerUrl
	}
	return ""
}

func (x *Test) GetInput() isTest_Input {
	if x != nil {
		return x.Input
	}
	return nil
}

func (x *Test) GetInputUrl() string {
	if x != nil {
		if x, ok := x.Input.(*Test_InputUrl); ok {
			return x.InputUrl
		}
	}
	return ""
}

func (x *Test) GetInputGenerator() *Test_Generator {
	if x != nil {
		if x, ok := x.Input.(*Test_InputGenerator); ok {
			return x.InputGenerator
		}
	}
	return nil
}

func (x *Test) GetAnswer() isTest_Answer {
	if x != nil {
		return x.Answer
	}
	return nil
}

func (x *Test) GetAnswerUrl() string {
	if x != nil {
		if x, ok := x.Answer.(*Test_AnswerUrl); ok {
			return x.AnswerUrl
		}
	}
	return ""
}

func (x *Test) GetAnswerGenerator() *Test_Generator {
	if x != nil {
		if x, ok := x.Answer.(*Test_AnswerGenerator); ok {
			return x.AnswerGenerator
		}
	}
	return nil
}

type isTest_Input interface {
	isTest_Input()
}

type Test_InputUrl struct {
	InputUrl string `protobuf:"bytes,11,opt,name=input_url,json=inputUrl,proto3,oneof"`
}

type Test_InputGenerator struct {
	InputGenerator *Test_Generator `protobuf:"bytes,12,opt,name=input_generator,json=inputGenerator,proto3,oneof"`
}

func (*Test_InputUrl) isTest_Input() {}

func (*Test_InputGenerator) isTest_Input() {}

type isTest_Answer interface {
	isTest_Answer()
}

type Test_AnswerUrl struct {
	AnswerUrl string `protobuf:"bytes,21,opt,name=answer_url,json=answerUrl,proto3,oneof"`
}

type Test_AnswerGenerator struct {
	AnswerGenerator *Test_Generator `protobuf:"bytes,22,opt,name=answer_generator,json=answerGenerator,proto3,oneof"`
}

func (*Test_AnswerUrl) isTest_Answer() {}

func (*Test_AnswerGenerator) isTest_Answer() {}

type Test_Generator struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ScriptName    string                 `protobuf:"bytes,1,opt,name=script_name,json=scriptName,proto3" json:"script_name,omitempty"`
	Arguments     []string               `protobuf:"bytes,2,rep,name=arguments,proto3" json:"arguments,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Test_Generator) Reset() {
	*x = Test_Generator{}
	mi := &file_eolymp_atlas_testing_test_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Test_Generator) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Test_Generator) ProtoMessage() {}

func (x *Test_Generator) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_atlas_testing_test_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Test_Generator.ProtoReflect.Descriptor instead.
func (*Test_Generator) Descriptor() ([]byte, []int) {
	return file_eolymp_atlas_testing_test_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Test_Generator) GetScriptName() string {
	if x != nil {
		return x.ScriptName
	}
	return ""
}

func (x *Test_Generator) GetArguments() []string {
	if x != nil {
		return x.Arguments
	}
	return nil
}

var File_eolymp_atlas_testing_test_proto protoreflect.FileDescriptor

var file_eolymp_atlas_testing_test_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x61, 0x74, 0x6c, 0x61, 0x73, 0x2f, 0x74,
	0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0c, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x61, 0x74, 0x6c, 0x61, 0x73, 0x22,
	0x9e, 0x05, 0x0a, 0x04, 0x54, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x65, 0x73, 0x74,
	0x73, 0x65, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x65,
	0x73, 0x74, 0x73, 0x65, 0x74, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x18, 0x0a,
	0x07, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07,
	0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x6e, 0x61, 0x63, 0x74,
	0x69, 0x76, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x69, 0x6e, 0x61, 0x63, 0x74,
	0x69, 0x76, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x06, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x73,
	0x63, 0x6f, 0x72, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x73, 0x63, 0x6f, 0x72,
	0x65, 0x12, 0x2a, 0x0a, 0x11, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x5f, 0x69, 0x6e, 0x70,
	0x75, 0x74, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x28, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x65, 0x78,
	0x61, 0x6d, 0x70, 0x6c, 0x65, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x55, 0x72, 0x6c, 0x12, 0x2c, 0x0a,
	0x12, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x5f, 0x61, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x5f,
	0x75, 0x72, 0x6c, 0x18, 0x29, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x65, 0x78, 0x61, 0x6d, 0x70,
	0x6c, 0x65, 0x41, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x55, 0x72, 0x6c, 0x12, 0x2e, 0x0a, 0x13, 0x67,
	0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x5f, 0x75,
	0x72, 0x6c, 0x18, 0x32, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61,
	0x74, 0x65, 0x64, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x55, 0x72, 0x6c, 0x12, 0x30, 0x0a, 0x14, 0x67,
	0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x5f,
	0x75, 0x72, 0x6c, 0x18, 0x33, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x67, 0x65, 0x6e, 0x65, 0x72,
	0x61, 0x74, 0x65, 0x64, 0x41, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x55, 0x72, 0x6c, 0x12, 0x1d, 0x0a,
	0x09, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09,
	0x48, 0x00, 0x52, 0x08, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x55, 0x72, 0x6c, 0x12, 0x47, 0x0a, 0x0f,
	0x69, 0x6e, 0x70, 0x75, 0x74, 0x5f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x18,
	0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x61,
	0x74, 0x6c, 0x61, 0x73, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x2e, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61,
	0x74, 0x6f, 0x72, 0x48, 0x00, 0x52, 0x0e, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x47, 0x65, 0x6e, 0x65,
	0x72, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x1f, 0x0a, 0x0a, 0x61, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x5f,
	0x75, 0x72, 0x6c, 0x18, 0x15, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x09, 0x61, 0x6e, 0x73,
	0x77, 0x65, 0x72, 0x55, 0x72, 0x6c, 0x12, 0x49, 0x0a, 0x10, 0x61, 0x6e, 0x73, 0x77, 0x65, 0x72,
	0x5f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x18, 0x16, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1c, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x61, 0x74, 0x6c, 0x61, 0x73, 0x2e,
	0x54, 0x65, 0x73, 0x74, 0x2e, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x48, 0x01,
	0x52, 0x0f, 0x61, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x6f,
	0x72, 0x1a, 0x4a, 0x0a, 0x09, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x1f,
	0x0a, 0x0b, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x1c, 0x0a, 0x09, 0x61, 0x72, 0x67, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x09, 0x61, 0x72, 0x67, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x42, 0x07, 0x0a,
	0x05, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x42, 0x08, 0x0a, 0x06, 0x61, 0x6e, 0x73, 0x77, 0x65, 0x72,
	0x42, 0x2d, 0x5a, 0x2b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65,
	0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x67, 0x6f, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x65, 0x6f, 0x6c,
	0x79, 0x6d, 0x70, 0x2f, 0x61, 0x74, 0x6c, 0x61, 0x73, 0x3b, 0x61, 0x74, 0x6c, 0x61, 0x73, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_eolymp_atlas_testing_test_proto_rawDescOnce sync.Once
	file_eolymp_atlas_testing_test_proto_rawDescData = file_eolymp_atlas_testing_test_proto_rawDesc
)

func file_eolymp_atlas_testing_test_proto_rawDescGZIP() []byte {
	file_eolymp_atlas_testing_test_proto_rawDescOnce.Do(func() {
		file_eolymp_atlas_testing_test_proto_rawDescData = protoimpl.X.CompressGZIP(file_eolymp_atlas_testing_test_proto_rawDescData)
	})
	return file_eolymp_atlas_testing_test_proto_rawDescData
}

var file_eolymp_atlas_testing_test_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_eolymp_atlas_testing_test_proto_goTypes = []any{
	(*Test)(nil),           // 0: eolymp.atlas.Test
	(*Test_Generator)(nil), // 1: eolymp.atlas.Test.Generator
}
var file_eolymp_atlas_testing_test_proto_depIdxs = []int32{
	1, // 0: eolymp.atlas.Test.input_generator:type_name -> eolymp.atlas.Test.Generator
	1, // 1: eolymp.atlas.Test.answer_generator:type_name -> eolymp.atlas.Test.Generator
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_eolymp_atlas_testing_test_proto_init() }
func file_eolymp_atlas_testing_test_proto_init() {
	if File_eolymp_atlas_testing_test_proto != nil {
		return
	}
	file_eolymp_atlas_testing_test_proto_msgTypes[0].OneofWrappers = []any{
		(*Test_InputUrl)(nil),
		(*Test_InputGenerator)(nil),
		(*Test_AnswerUrl)(nil),
		(*Test_AnswerGenerator)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_eolymp_atlas_testing_test_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_eolymp_atlas_testing_test_proto_goTypes,
		DependencyIndexes: file_eolymp_atlas_testing_test_proto_depIdxs,
		MessageInfos:      file_eolymp_atlas_testing_test_proto_msgTypes,
	}.Build()
	File_eolymp_atlas_testing_test_proto = out.File
	file_eolymp_atlas_testing_test_proto_rawDesc = nil
	file_eolymp_atlas_testing_test_proto_goTypes = nil
	file_eolymp_atlas_testing_test_proto_depIdxs = nil
}
