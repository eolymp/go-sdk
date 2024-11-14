// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v3.21.12
// source: eolymp/rating/rating.proto

package rating

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Rating struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Timestamp *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	MemberId  string                 `protobuf:"bytes,3,opt,name=member_id,json=memberId,proto3" json:"member_id,omitempty"`
	ContestId string                 `protobuf:"bytes,4,opt,name=contest_id,json=contestId,proto3" json:"contest_id,omitempty"`
	Value     uint32                 `protobuf:"varint,10,opt,name=value,proto3" json:"value,omitempty"`
	Level     uint32                 `protobuf:"varint,11,opt,name=level,proto3" json:"level,omitempty"`
}

func (x *Rating) Reset() {
	*x = Rating{}
	mi := &file_eolymp_rating_rating_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Rating) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Rating) ProtoMessage() {}

func (x *Rating) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_rating_rating_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Rating.ProtoReflect.Descriptor instead.
func (*Rating) Descriptor() ([]byte, []int) {
	return file_eolymp_rating_rating_proto_rawDescGZIP(), []int{0}
}

func (x *Rating) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Rating) GetTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

func (x *Rating) GetMemberId() string {
	if x != nil {
		return x.MemberId
	}
	return ""
}

func (x *Rating) GetContestId() string {
	if x != nil {
		return x.ContestId
	}
	return ""
}

func (x *Rating) GetValue() uint32 {
	if x != nil {
		return x.Value
	}
	return 0
}

func (x *Rating) GetLevel() uint32 {
	if x != nil {
		return x.Level
	}
	return 0
}

var File_eolymp_rating_rating_proto protoreflect.FileDescriptor

var file_eolymp_rating_rating_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x2f,
	0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x65, 0x6f,
	0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x1a, 0x1f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xba, 0x01, 0x0a,
	0x06, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x38, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1d,
	0x0a, 0x0a, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x73, 0x74, 0x49, 0x64, 0x12, 0x14, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x0b, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x42, 0x2f, 0x5a, 0x2d, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x67,
	0x6f, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x72, 0x61, 0x74,
	0x69, 0x6e, 0x67, 0x3b, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_eolymp_rating_rating_proto_rawDescOnce sync.Once
	file_eolymp_rating_rating_proto_rawDescData = file_eolymp_rating_rating_proto_rawDesc
)

func file_eolymp_rating_rating_proto_rawDescGZIP() []byte {
	file_eolymp_rating_rating_proto_rawDescOnce.Do(func() {
		file_eolymp_rating_rating_proto_rawDescData = protoimpl.X.CompressGZIP(file_eolymp_rating_rating_proto_rawDescData)
	})
	return file_eolymp_rating_rating_proto_rawDescData
}

var file_eolymp_rating_rating_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_eolymp_rating_rating_proto_goTypes = []any{
	(*Rating)(nil),                // 0: eolymp.rating.Rating
	(*timestamppb.Timestamp)(nil), // 1: google.protobuf.Timestamp
}
var file_eolymp_rating_rating_proto_depIdxs = []int32{
	1, // 0: eolymp.rating.Rating.timestamp:type_name -> google.protobuf.Timestamp
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_eolymp_rating_rating_proto_init() }
func file_eolymp_rating_rating_proto_init() {
	if File_eolymp_rating_rating_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_eolymp_rating_rating_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_eolymp_rating_rating_proto_goTypes,
		DependencyIndexes: file_eolymp_rating_rating_proto_depIdxs,
		MessageInfos:      file_eolymp_rating_rating_proto_msgTypes,
	}.Build()
	File_eolymp_rating_rating_proto = out.File
	file_eolymp_rating_rating_proto_rawDesc = nil
	file_eolymp_rating_rating_proto_goTypes = nil
	file_eolymp_rating_rating_proto_depIdxs = nil
}
