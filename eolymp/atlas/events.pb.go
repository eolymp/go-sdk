// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.18.1
// source: eolymp/atlas/events.proto

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

type PermissionsDeletedEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProblemId string `protobuf:"bytes,1,opt,name=problem_id,json=problemId,proto3" json:"problem_id,omitempty"`
	UserId    string `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *PermissionsDeletedEvent) Reset() {
	*x = PermissionsDeletedEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_atlas_events_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PermissionsDeletedEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PermissionsDeletedEvent) ProtoMessage() {}

func (x *PermissionsDeletedEvent) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_atlas_events_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PermissionsDeletedEvent.ProtoReflect.Descriptor instead.
func (*PermissionsDeletedEvent) Descriptor() ([]byte, []int) {
	return file_eolymp_atlas_events_proto_rawDescGZIP(), []int{0}
}

func (x *PermissionsDeletedEvent) GetProblemId() string {
	if x != nil {
		return x.ProblemId
	}
	return ""
}

func (x *PermissionsDeletedEvent) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

// SubmissionCompleteEvent is dispatched every time submission is fully judged (got final result).
type SubmissionCompleteEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Submission *Submission `protobuf:"bytes,1,opt,name=submission,proto3" json:"submission,omitempty"` // submission
	Update     bool        `protobuf:"varint,2,opt,name=update,proto3" json:"update,omitempty"`        // true, if submission has been completed before, this flag is set when submission is being retested
}

func (x *SubmissionCompleteEvent) Reset() {
	*x = SubmissionCompleteEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_atlas_events_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubmissionCompleteEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubmissionCompleteEvent) ProtoMessage() {}

func (x *SubmissionCompleteEvent) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_atlas_events_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubmissionCompleteEvent.ProtoReflect.Descriptor instead.
func (*SubmissionCompleteEvent) Descriptor() ([]byte, []int) {
	return file_eolymp_atlas_events_proto_rawDescGZIP(), []int{1}
}

func (x *SubmissionCompleteEvent) GetSubmission() *Submission {
	if x != nil {
		return x.Submission
	}
	return nil
}

func (x *SubmissionCompleteEvent) GetUpdate() bool {
	if x != nil {
		return x.Update
	}
	return false
}

// ScoreUpdatedEvent is dispatched when score for a given problem and user has been updated (increased).
type ScoreUpdatedEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Score *Score `protobuf:"bytes,1,opt,name=score,proto3" json:"score,omitempty"`
}

func (x *ScoreUpdatedEvent) Reset() {
	*x = ScoreUpdatedEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_atlas_events_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ScoreUpdatedEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ScoreUpdatedEvent) ProtoMessage() {}

func (x *ScoreUpdatedEvent) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_atlas_events_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ScoreUpdatedEvent.ProtoReflect.Descriptor instead.
func (*ScoreUpdatedEvent) Descriptor() ([]byte, []int) {
	return file_eolymp_atlas_events_proto_rawDescGZIP(), []int{2}
}

func (x *ScoreUpdatedEvent) GetScore() *Score {
	if x != nil {
		return x.Score
	}
	return nil
}

var File_eolymp_atlas_events_proto protoreflect.FileDescriptor

var file_eolymp_atlas_events_proto_rawDesc = []byte{
	0x0a, 0x19, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x61, 0x74, 0x6c, 0x61, 0x73, 0x2f, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x65, 0x6f, 0x6c,
	0x79, 0x6d, 0x70, 0x2e, 0x61, 0x74, 0x6c, 0x61, 0x73, 0x1a, 0x20, 0x65, 0x6f, 0x6c, 0x79, 0x6d,
	0x70, 0x2f, 0x61, 0x74, 0x6c, 0x61, 0x73, 0x2f, 0x73, 0x63, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x5f,
	0x73, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x65, 0x6f, 0x6c,
	0x79, 0x6d, 0x70, 0x2f, 0x61, 0x74, 0x6c, 0x61, 0x73, 0x2f, 0x73, 0x75, 0x62, 0x6d, 0x69, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x51, 0x0a, 0x17, 0x50, 0x65,
	0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x62, 0x6c,
	0x65, 0x6d, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x6b, 0x0a,
	0x17, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6d, 0x70, 0x6c,
	0x65, 0x74, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x38, 0x0a, 0x0a, 0x73, 0x75, 0x62, 0x6d,
	0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x65,
	0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x61, 0x74, 0x6c, 0x61, 0x73, 0x2e, 0x53, 0x75, 0x62, 0x6d,
	0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x73, 0x75, 0x62, 0x6d, 0x69, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x06, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x22, 0x3e, 0x0a, 0x11, 0x53, 0x63,
	0x6f, 0x72, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12,
	0x29, 0x0a, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13,
	0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x61, 0x74, 0x6c, 0x61, 0x73, 0x2e, 0x53, 0x63,
	0x6f, 0x72, 0x65, 0x52, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x42, 0x2d, 0x5a, 0x2b, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f,
	0x67, 0x6f, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x61, 0x74,
	0x6c, 0x61, 0x73, 0x3b, 0x61, 0x74, 0x6c, 0x61, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_eolymp_atlas_events_proto_rawDescOnce sync.Once
	file_eolymp_atlas_events_proto_rawDescData = file_eolymp_atlas_events_proto_rawDesc
)

func file_eolymp_atlas_events_proto_rawDescGZIP() []byte {
	file_eolymp_atlas_events_proto_rawDescOnce.Do(func() {
		file_eolymp_atlas_events_proto_rawDescData = protoimpl.X.CompressGZIP(file_eolymp_atlas_events_proto_rawDescData)
	})
	return file_eolymp_atlas_events_proto_rawDescData
}

var file_eolymp_atlas_events_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_eolymp_atlas_events_proto_goTypes = []interface{}{
	(*PermissionsDeletedEvent)(nil), // 0: eolymp.atlas.PermissionsDeletedEvent
	(*SubmissionCompleteEvent)(nil), // 1: eolymp.atlas.SubmissionCompleteEvent
	(*ScoreUpdatedEvent)(nil),       // 2: eolymp.atlas.ScoreUpdatedEvent
	(*Submission)(nil),              // 3: eolymp.atlas.Submission
	(*Score)(nil),                   // 4: eolymp.atlas.Score
}
var file_eolymp_atlas_events_proto_depIdxs = []int32{
	3, // 0: eolymp.atlas.SubmissionCompleteEvent.submission:type_name -> eolymp.atlas.Submission
	4, // 1: eolymp.atlas.ScoreUpdatedEvent.score:type_name -> eolymp.atlas.Score
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_eolymp_atlas_events_proto_init() }
func file_eolymp_atlas_events_proto_init() {
	if File_eolymp_atlas_events_proto != nil {
		return
	}
	file_eolymp_atlas_scoring_score_proto_init()
	file_eolymp_atlas_submission_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_eolymp_atlas_events_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PermissionsDeletedEvent); i {
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
		file_eolymp_atlas_events_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubmissionCompleteEvent); i {
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
		file_eolymp_atlas_events_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ScoreUpdatedEvent); i {
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
			RawDescriptor: file_eolymp_atlas_events_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_eolymp_atlas_events_proto_goTypes,
		DependencyIndexes: file_eolymp_atlas_events_proto_depIdxs,
		MessageInfos:      file_eolymp_atlas_events_proto_msgTypes,
	}.Build()
	File_eolymp_atlas_events_proto = out.File
	file_eolymp_atlas_events_proto_rawDesc = nil
	file_eolymp_atlas_events_proto_goTypes = nil
	file_eolymp_atlas_events_proto_depIdxs = nil
}
