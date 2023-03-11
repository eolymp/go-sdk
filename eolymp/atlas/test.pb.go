// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.29.0
// 	protoc        v3.18.1
// source: eolymp/atlas/test.proto

package atlas

import (
	_ "github.com/eolymp/go-sdk/eolymp/annotations"
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
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Test unique identifier.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Testset where this test belongs.
	TestsetId string `protobuf:"bytes,2,opt,name=testset_id,json=testsetId,proto3" json:"testset_id,omitempty"`
	// Index, a sequential number, of the test within testset.
	Index int32 `protobuf:"varint,3,opt,name=index,proto3" json:"index,omitempty"`
	// Flag which defines if test should be shown in as an example in the problem statement.
	Example bool `protobuf:"varint,4,opt,name=example,proto3" json:"example,omitempty"`
	// Score for passing this test.
	Score float32 `protobuf:"fixed32,5,opt,name=score,proto3" json:"score,omitempty"`
	// Input Object ID, use keeper service to upload or download actual test data.
	InputObjectId string `protobuf:"bytes,10,opt,name=input_object_id,json=inputObjectId,proto3" json:"input_object_id,omitempty"`
	// Answer Object ID, use keeper service to upload or download actual test data.
	AnswerObjectId string `protobuf:"bytes,20,opt,name=answer_object_id,json=answerObjectId,proto3" json:"answer_object_id,omitempty"`
}

func (x *Test) Reset() {
	*x = Test{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_atlas_test_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Test) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Test) ProtoMessage() {}

func (x *Test) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_atlas_test_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
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
	return file_eolymp_atlas_test_proto_rawDescGZIP(), []int{0}
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

func (x *Test) GetScore() float32 {
	if x != nil {
		return x.Score
	}
	return 0
}

func (x *Test) GetInputObjectId() string {
	if x != nil {
		return x.InputObjectId
	}
	return ""
}

func (x *Test) GetAnswerObjectId() string {
	if x != nil {
		return x.AnswerObjectId
	}
	return ""
}

var File_eolymp_atlas_test_proto protoreflect.FileDescriptor

var file_eolymp_atlas_test_proto_rawDesc = []byte{
	0x0a, 0x17, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x61, 0x74, 0x6c, 0x61, 0x73, 0x2f, 0x74,
	0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x65, 0x6f, 0x6c, 0x79, 0x6d,
	0x70, 0x2e, 0x61, 0x74, 0x6c, 0x61, 0x73, 0x1a, 0x21, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf3, 0x01, 0x0a, 0x04, 0x54,
	0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x65, 0x73, 0x74, 0x73, 0x65, 0x74, 0x5f, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x65, 0x73, 0x74, 0x73, 0x65, 0x74,
	0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x78, 0x61, 0x6d,
	0x70, 0x6c, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x65, 0x78, 0x61, 0x6d, 0x70,
	0x6c, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x02, 0x52, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x12, 0x26, 0x0a, 0x0f, 0x69, 0x6e, 0x70, 0x75,
	0x74, 0x5f, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0d, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64,
	0x12, 0x28, 0x0a, 0x10, 0x61, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x5f, 0x6f, 0x62, 0x6a, 0x65, 0x63,
	0x74, 0x5f, 0x69, 0x64, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x61, 0x6e, 0x73, 0x77,
	0x65, 0x72, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x3a, 0x24, 0xb2, 0xe3, 0x0a, 0x20,
	0xba, 0xe3, 0x0a, 0x04, 0x74, 0x65, 0x73, 0x74, 0xc2, 0xe3, 0x0a, 0x14, 0x65, 0x6f, 0x6c, 0x79,
	0x6d, 0x70, 0x2e, 0x61, 0x74, 0x6c, 0x61, 0x73, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x73, 0x65, 0x74,
	0x42, 0x2d, 0x5a, 0x2b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65,
	0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x67, 0x6f, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x65, 0x6f, 0x6c,
	0x79, 0x6d, 0x70, 0x2f, 0x61, 0x74, 0x6c, 0x61, 0x73, 0x3b, 0x61, 0x74, 0x6c, 0x61, 0x73, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_eolymp_atlas_test_proto_rawDescOnce sync.Once
	file_eolymp_atlas_test_proto_rawDescData = file_eolymp_atlas_test_proto_rawDesc
)

func file_eolymp_atlas_test_proto_rawDescGZIP() []byte {
	file_eolymp_atlas_test_proto_rawDescOnce.Do(func() {
		file_eolymp_atlas_test_proto_rawDescData = protoimpl.X.CompressGZIP(file_eolymp_atlas_test_proto_rawDescData)
	})
	return file_eolymp_atlas_test_proto_rawDescData
}

var file_eolymp_atlas_test_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_eolymp_atlas_test_proto_goTypes = []interface{}{
	(*Test)(nil), // 0: eolymp.atlas.Test
}
var file_eolymp_atlas_test_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_eolymp_atlas_test_proto_init() }
func file_eolymp_atlas_test_proto_init() {
	if File_eolymp_atlas_test_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_eolymp_atlas_test_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Test); i {
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
			RawDescriptor: file_eolymp_atlas_test_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_eolymp_atlas_test_proto_goTypes,
		DependencyIndexes: file_eolymp_atlas_test_proto_depIdxs,
		MessageInfos:      file_eolymp_atlas_test_proto_msgTypes,
	}.Build()
	File_eolymp_atlas_test_proto = out.File
	file_eolymp_atlas_test_proto_rawDesc = nil
	file_eolymp_atlas_test_proto_goTypes = nil
	file_eolymp_atlas_test_proto_depIdxs = nil
}
