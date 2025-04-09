// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v4.24.4
// source: eolymp/judge/activity.proto

package judge

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
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

type Activity_Type int32

const (
	Activity_NONE               Activity_Type = 0
	Activity_PROBLEM_RETEST     Activity_Type = 1
	Activity_SCOREBOARD_REBUILD Activity_Type = 2
)

// Enum value maps for Activity_Type.
var (
	Activity_Type_name = map[int32]string{
		0: "NONE",
		1: "PROBLEM_RETEST",
		2: "SCOREBOARD_REBUILD",
	}
	Activity_Type_value = map[string]int32{
		"NONE":               0,
		"PROBLEM_RETEST":     1,
		"SCOREBOARD_REBUILD": 2,
	}
)

func (x Activity_Type) Enum() *Activity_Type {
	p := new(Activity_Type)
	*p = x
	return p
}

func (x Activity_Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Activity_Type) Descriptor() protoreflect.EnumDescriptor {
	return file_eolymp_judge_activity_proto_enumTypes[0].Descriptor()
}

func (Activity_Type) Type() protoreflect.EnumType {
	return &file_eolymp_judge_activity_proto_enumTypes[0]
}

func (x Activity_Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Activity_Type.Descriptor instead.
func (Activity_Type) EnumDescriptor() ([]byte, []int) {
	return file_eolymp_judge_activity_proto_rawDescGZIP(), []int{0, 0}
}

type Activity_Status int32

const (
	Activity_UNKNOWN  Activity_Status = 0
	Activity_CREATED  Activity_Status = 1
	Activity_STARTED  Activity_Status = 2
	Activity_COMPLETE Activity_Status = 3
	Activity_ERROR    Activity_Status = 4
)

// Enum value maps for Activity_Status.
var (
	Activity_Status_name = map[int32]string{
		0: "UNKNOWN",
		1: "CREATED",
		2: "STARTED",
		3: "COMPLETE",
		4: "ERROR",
	}
	Activity_Status_value = map[string]int32{
		"UNKNOWN":  0,
		"CREATED":  1,
		"STARTED":  2,
		"COMPLETE": 3,
		"ERROR":    4,
	}
)

func (x Activity_Status) Enum() *Activity_Status {
	p := new(Activity_Status)
	*p = x
	return p
}

func (x Activity_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Activity_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_eolymp_judge_activity_proto_enumTypes[1].Descriptor()
}

func (Activity_Status) Type() protoreflect.EnumType {
	return &file_eolymp_judge_activity_proto_enumTypes[1]
}

func (x Activity_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Activity_Status.Descriptor instead.
func (Activity_Status) EnumDescriptor() ([]byte, []int) {
	return file_eolymp_judge_activity_proto_rawDescGZIP(), []int{0, 1}
}

type Activity struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Activity unique identifier.
	Id     string          `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Type   Activity_Type   `protobuf:"varint,2,opt,name=type,proto3,enum=eolymp.judge.Activity_Type" json:"type,omitempty"`
	Status Activity_Status `protobuf:"varint,3,opt,name=status,proto3,enum=eolymp.judge.Activity_Status" json:"status,omitempty"`
	// Contest where activity is running.
	ContestId string `protobuf:"bytes,5,opt,name=contest_id,json=contestId,proto3" json:"contest_id,omitempty"`
	// Context for activity, for example scoreboard being rebuild, or problem retested.
	ScoreboardId string `protobuf:"bytes,100,opt,name=scoreboard_id,json=scoreboardId,proto3" json:"scoreboard_id,omitempty"`
	ProblemId    string `protobuf:"bytes,101,opt,name=problem_id,json=problemId,proto3" json:"problem_id,omitempty"`
	// Timestamp when activity was created, started, updated and complete.
	CreatedAt  *timestamppb.Timestamp `protobuf:"bytes,10,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	StartedAt  *timestamppb.Timestamp `protobuf:"bytes,11,opt,name=started_at,json=startedAt,proto3" json:"started_at,omitempty"`
	ProgressAt *timestamppb.Timestamp `protobuf:"bytes,12,opt,name=progress_at,json=progressAt,proto3" json:"progress_at,omitempty"`
	CompleteAt *timestamppb.Timestamp `protobuf:"bytes,13,opt,name=complete_at,json=completeAt,proto3" json:"complete_at,omitempty"`
	// progress is a number from 0 to `total`, showing amount of work complete
	Progress uint32 `protobuf:"varint,20,opt,name=progress,proto3" json:"progress,omitempty"`
	// total is a number indicating total amount of work to be complete, 0 means it's unknown
	Total uint32 `protobuf:"varint,21,opt,name=total,proto3" json:"total,omitempty"`
	// error is a string indicating why job has failed
	Error         string `protobuf:"bytes,30,opt,name=error,proto3" json:"error,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Activity) Reset() {
	*x = Activity{}
	mi := &file_eolymp_judge_activity_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Activity) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Activity) ProtoMessage() {}

func (x *Activity) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_judge_activity_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Activity.ProtoReflect.Descriptor instead.
func (*Activity) Descriptor() ([]byte, []int) {
	return file_eolymp_judge_activity_proto_rawDescGZIP(), []int{0}
}

func (x *Activity) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Activity) GetType() Activity_Type {
	if x != nil {
		return x.Type
	}
	return Activity_NONE
}

func (x *Activity) GetStatus() Activity_Status {
	if x != nil {
		return x.Status
	}
	return Activity_UNKNOWN
}

func (x *Activity) GetContestId() string {
	if x != nil {
		return x.ContestId
	}
	return ""
}

func (x *Activity) GetScoreboardId() string {
	if x != nil {
		return x.ScoreboardId
	}
	return ""
}

func (x *Activity) GetProblemId() string {
	if x != nil {
		return x.ProblemId
	}
	return ""
}

func (x *Activity) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Activity) GetStartedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.StartedAt
	}
	return nil
}

func (x *Activity) GetProgressAt() *timestamppb.Timestamp {
	if x != nil {
		return x.ProgressAt
	}
	return nil
}

func (x *Activity) GetCompleteAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CompleteAt
	}
	return nil
}

func (x *Activity) GetProgress() uint32 {
	if x != nil {
		return x.Progress
	}
	return 0
}

func (x *Activity) GetTotal() uint32 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *Activity) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

var File_eolymp_judge_activity_proto protoreflect.FileDescriptor

const file_eolymp_judge_activity_proto_rawDesc = "" +
	"\n" +
	"\x1beolymp/judge/activity.proto\x12\feolymp.judge\x1a\x1fgoogle/protobuf/timestamp.proto\"\xa5\x05\n" +
	"\bActivity\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x12/\n" +
	"\x04type\x18\x02 \x01(\x0e2\x1b.eolymp.judge.Activity.TypeR\x04type\x125\n" +
	"\x06status\x18\x03 \x01(\x0e2\x1d.eolymp.judge.Activity.StatusR\x06status\x12\x1d\n" +
	"\n" +
	"contest_id\x18\x05 \x01(\tR\tcontestId\x12#\n" +
	"\rscoreboard_id\x18d \x01(\tR\fscoreboardId\x12\x1d\n" +
	"\n" +
	"problem_id\x18e \x01(\tR\tproblemId\x129\n" +
	"\n" +
	"created_at\x18\n" +
	" \x01(\v2\x1a.google.protobuf.TimestampR\tcreatedAt\x129\n" +
	"\n" +
	"started_at\x18\v \x01(\v2\x1a.google.protobuf.TimestampR\tstartedAt\x12;\n" +
	"\vprogress_at\x18\f \x01(\v2\x1a.google.protobuf.TimestampR\n" +
	"progressAt\x12;\n" +
	"\vcomplete_at\x18\r \x01(\v2\x1a.google.protobuf.TimestampR\n" +
	"completeAt\x12\x1a\n" +
	"\bprogress\x18\x14 \x01(\rR\bprogress\x12\x14\n" +
	"\x05total\x18\x15 \x01(\rR\x05total\x12\x14\n" +
	"\x05error\x18\x1e \x01(\tR\x05error\"<\n" +
	"\x04Type\x12\b\n" +
	"\x04NONE\x10\x00\x12\x12\n" +
	"\x0ePROBLEM_RETEST\x10\x01\x12\x16\n" +
	"\x12SCOREBOARD_REBUILD\x10\x02\"H\n" +
	"\x06Status\x12\v\n" +
	"\aUNKNOWN\x10\x00\x12\v\n" +
	"\aCREATED\x10\x01\x12\v\n" +
	"\aSTARTED\x10\x02\x12\f\n" +
	"\bCOMPLETE\x10\x03\x12\t\n" +
	"\x05ERROR\x10\x04B-Z+github.com/eolymp/go-sdk/eolymp/judge;judgeb\x06proto3"

var (
	file_eolymp_judge_activity_proto_rawDescOnce sync.Once
	file_eolymp_judge_activity_proto_rawDescData []byte
)

func file_eolymp_judge_activity_proto_rawDescGZIP() []byte {
	file_eolymp_judge_activity_proto_rawDescOnce.Do(func() {
		file_eolymp_judge_activity_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_eolymp_judge_activity_proto_rawDesc), len(file_eolymp_judge_activity_proto_rawDesc)))
	})
	return file_eolymp_judge_activity_proto_rawDescData
}

var file_eolymp_judge_activity_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_eolymp_judge_activity_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_eolymp_judge_activity_proto_goTypes = []any{
	(Activity_Type)(0),            // 0: eolymp.judge.Activity.Type
	(Activity_Status)(0),          // 1: eolymp.judge.Activity.Status
	(*Activity)(nil),              // 2: eolymp.judge.Activity
	(*timestamppb.Timestamp)(nil), // 3: google.protobuf.Timestamp
}
var file_eolymp_judge_activity_proto_depIdxs = []int32{
	0, // 0: eolymp.judge.Activity.type:type_name -> eolymp.judge.Activity.Type
	1, // 1: eolymp.judge.Activity.status:type_name -> eolymp.judge.Activity.Status
	3, // 2: eolymp.judge.Activity.created_at:type_name -> google.protobuf.Timestamp
	3, // 3: eolymp.judge.Activity.started_at:type_name -> google.protobuf.Timestamp
	3, // 4: eolymp.judge.Activity.progress_at:type_name -> google.protobuf.Timestamp
	3, // 5: eolymp.judge.Activity.complete_at:type_name -> google.protobuf.Timestamp
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_eolymp_judge_activity_proto_init() }
func file_eolymp_judge_activity_proto_init() {
	if File_eolymp_judge_activity_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_eolymp_judge_activity_proto_rawDesc), len(file_eolymp_judge_activity_proto_rawDesc)),
			NumEnums:      2,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_eolymp_judge_activity_proto_goTypes,
		DependencyIndexes: file_eolymp_judge_activity_proto_depIdxs,
		EnumInfos:         file_eolymp_judge_activity_proto_enumTypes,
		MessageInfos:      file_eolymp_judge_activity_proto_msgTypes,
	}.Build()
	File_eolymp_judge_activity_proto = out.File
	file_eolymp_judge_activity_proto_goTypes = nil
	file_eolymp_judge_activity_proto_depIdxs = nil
}
