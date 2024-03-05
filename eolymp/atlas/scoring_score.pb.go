// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v3.18.1
// source: eolymp/atlas/scoring_score.proto

package atlas

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

// Score for a problem
type Score struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`                                // unique identifier
	ProblemId string                 `protobuf:"bytes,2,opt,name=problem_id,json=problemId,proto3" json:"problem_id,omitempty"` // problem
	UserId    string                 `protobuf:"bytes,3,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`          // user
	MemberId  string                 `protobuf:"bytes,7,opt,name=member_id,json=memberId,proto3" json:"member_id,omitempty"`    // member
	SolvedAt  *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=solved_at,json=solvedAt,proto3" json:"solved_at,omitempty"`    // time when submission was created
	Score     float32                `protobuf:"fixed32,5,opt,name=score,proto3" json:"score,omitempty"`                        // score from 0 (none of the tests are passing) to 1 (all tests are passing)
	Attempts  uint32                 `protobuf:"varint,8,opt,name=attempts,proto3" json:"attempts,omitempty"`
}

func (x *Score) Reset() {
	*x = Score{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_atlas_scoring_score_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Score) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Score) ProtoMessage() {}

func (x *Score) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_atlas_scoring_score_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Score.ProtoReflect.Descriptor instead.
func (*Score) Descriptor() ([]byte, []int) {
	return file_eolymp_atlas_scoring_score_proto_rawDescGZIP(), []int{0}
}

func (x *Score) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Score) GetProblemId() string {
	if x != nil {
		return x.ProblemId
	}
	return ""
}

func (x *Score) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *Score) GetMemberId() string {
	if x != nil {
		return x.MemberId
	}
	return ""
}

func (x *Score) GetSolvedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.SolvedAt
	}
	return nil
}

func (x *Score) GetScore() float32 {
	if x != nil {
		return x.Score
	}
	return 0
}

func (x *Score) GetAttempts() uint32 {
	if x != nil {
		return x.Attempts
	}
	return 0
}

var File_eolymp_atlas_scoring_score_proto protoreflect.FileDescriptor

var file_eolymp_atlas_scoring_score_proto_rawDesc = []byte{
	0x0a, 0x20, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x61, 0x74, 0x6c, 0x61, 0x73, 0x2f, 0x73,
	0x63, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0c, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x61, 0x74, 0x6c, 0x61, 0x73,
	0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xd7, 0x01, 0x0a, 0x05, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x70,
	0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x70, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x49, 0x64,
	0x12, 0x37, 0x0a, 0x09, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x08, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x64, 0x41, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x63, 0x6f,
	0x72, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x12,
	0x1a, 0x0a, 0x08, 0x61, 0x74, 0x74, 0x65, 0x6d, 0x70, 0x74, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x08, 0x61, 0x74, 0x74, 0x65, 0x6d, 0x70, 0x74, 0x73, 0x42, 0x2d, 0x5a, 0x2b, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70,
	0x2f, 0x67, 0x6f, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x61,
	0x74, 0x6c, 0x61, 0x73, 0x3b, 0x61, 0x74, 0x6c, 0x61, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_eolymp_atlas_scoring_score_proto_rawDescOnce sync.Once
	file_eolymp_atlas_scoring_score_proto_rawDescData = file_eolymp_atlas_scoring_score_proto_rawDesc
)

func file_eolymp_atlas_scoring_score_proto_rawDescGZIP() []byte {
	file_eolymp_atlas_scoring_score_proto_rawDescOnce.Do(func() {
		file_eolymp_atlas_scoring_score_proto_rawDescData = protoimpl.X.CompressGZIP(file_eolymp_atlas_scoring_score_proto_rawDescData)
	})
	return file_eolymp_atlas_scoring_score_proto_rawDescData
}

var file_eolymp_atlas_scoring_score_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_eolymp_atlas_scoring_score_proto_goTypes = []interface{}{
	(*Score)(nil),                 // 0: eolymp.atlas.Score
	(*timestamppb.Timestamp)(nil), // 1: google.protobuf.Timestamp
}
var file_eolymp_atlas_scoring_score_proto_depIdxs = []int32{
	1, // 0: eolymp.atlas.Score.solved_at:type_name -> google.protobuf.Timestamp
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_eolymp_atlas_scoring_score_proto_init() }
func file_eolymp_atlas_scoring_score_proto_init() {
	if File_eolymp_atlas_scoring_score_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_eolymp_atlas_scoring_score_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Score); i {
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
			RawDescriptor: file_eolymp_atlas_scoring_score_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_eolymp_atlas_scoring_score_proto_goTypes,
		DependencyIndexes: file_eolymp_atlas_scoring_score_proto_depIdxs,
		MessageInfos:      file_eolymp_atlas_scoring_score_proto_msgTypes,
	}.Build()
	File_eolymp_atlas_scoring_score_proto = out.File
	file_eolymp_atlas_scoring_score_proto_rawDesc = nil
	file_eolymp_atlas_scoring_score_proto_goTypes = nil
	file_eolymp_atlas_scoring_score_proto_depIdxs = nil
}
