// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.12.2
// source: eolymp/judge/contest.proto

package judge

import (
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type Contest_Status int32

const (
	Contest_NO_STATUS Contest_Status = 0 // reserved, should not be used
	Contest_SCHEDULED Contest_Status = 1
	Contest_OPEN      Contest_Status = 2
	Contest_COMPLETE  Contest_Status = 3
)

// Enum value maps for Contest_Status.
var (
	Contest_Status_name = map[int32]string{
		0: "NO_STATUS",
		1: "SCHEDULED",
		2: "OPEN",
		3: "COMPLETE",
	}
	Contest_Status_value = map[string]int32{
		"NO_STATUS": 0,
		"SCHEDULED": 1,
		"OPEN":      2,
		"COMPLETE":  3,
	}
)

func (x Contest_Status) Enum() *Contest_Status {
	p := new(Contest_Status)
	*p = x
	return p
}

func (x Contest_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Contest_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_eolymp_judge_contest_proto_enumTypes[0].Descriptor()
}

func (Contest_Status) Type() protoreflect.EnumType {
	return &file_eolymp_judge_contest_proto_enumTypes[0]
}

func (x Contest_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Contest_Status.Descriptor instead.
func (Contest_Status) EnumDescriptor() ([]byte, []int) {
	return file_eolymp_judge_contest_proto_rawDescGZIP(), []int{0, 0}
}

type Contest_Visibility int32

const (
	Contest_NO_VISIBILITY Contest_Visibility = 0 // reserved, should not be used
	Contest_PUBLIC        Contest_Visibility = 1 // visible to everyone and shown on the website
	Contest_UNLISTED      Contest_Visibility = 2 // anyone can participate, but not shown on the website
	Contest_PRIVATE       Contest_Visibility = 3 // only explicitly added participants can participate
)

// Enum value maps for Contest_Visibility.
var (
	Contest_Visibility_name = map[int32]string{
		0: "NO_VISIBILITY",
		1: "PUBLIC",
		2: "UNLISTED",
		3: "PRIVATE",
	}
	Contest_Visibility_value = map[string]int32{
		"NO_VISIBILITY": 0,
		"PUBLIC":        1,
		"UNLISTED":      2,
		"PRIVATE":       3,
	}
)

func (x Contest_Visibility) Enum() *Contest_Visibility {
	p := new(Contest_Visibility)
	*p = x
	return p
}

func (x Contest_Visibility) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Contest_Visibility) Descriptor() protoreflect.EnumDescriptor {
	return file_eolymp_judge_contest_proto_enumTypes[1].Descriptor()
}

func (Contest_Visibility) Type() protoreflect.EnumType {
	return &file_eolymp_judge_contest_proto_enumTypes[1]
}

func (x Contest_Visibility) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Contest_Visibility.Descriptor instead.
func (Contest_Visibility) EnumDescriptor() ([]byte, []int) {
	return file_eolymp_judge_contest_proto_rawDescGZIP(), []int{0, 1}
}

type Contest_ParticipationMode int32

const (
	Contest_NO_PARTICIPATION Contest_ParticipationMode = 0 // reserved, should not be used
	Contest_ONLINE           Contest_ParticipationMode = 1 // everyone start and finish together
	Contest_VIRTUAL          Contest_ParticipationMode = 2 // participants can start contest individually
)

// Enum value maps for Contest_ParticipationMode.
var (
	Contest_ParticipationMode_name = map[int32]string{
		0: "NO_PARTICIPATION",
		1: "ONLINE",
		2: "VIRTUAL",
	}
	Contest_ParticipationMode_value = map[string]int32{
		"NO_PARTICIPATION": 0,
		"ONLINE":           1,
		"VIRTUAL":          2,
	}
)

func (x Contest_ParticipationMode) Enum() *Contest_ParticipationMode {
	p := new(Contest_ParticipationMode)
	*p = x
	return p
}

func (x Contest_ParticipationMode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Contest_ParticipationMode) Descriptor() protoreflect.EnumDescriptor {
	return file_eolymp_judge_contest_proto_enumTypes[2].Descriptor()
}

func (Contest_ParticipationMode) Type() protoreflect.EnumType {
	return &file_eolymp_judge_contest_proto_enumTypes[2]
}

func (x Contest_ParticipationMode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Contest_ParticipationMode.Descriptor instead.
func (Contest_ParticipationMode) EnumDescriptor() ([]byte, []int) {
	return file_eolymp_judge_contest_proto_rawDescGZIP(), []int{0, 2}
}

type Contest_Format int32

const (
	Contest_NO_FORMAT Contest_Format = 0 // reserved, should not be used
	Contest_IOI       Contest_Format = 1
	Contest_ICPC      Contest_Format = 2
)

// Enum value maps for Contest_Format.
var (
	Contest_Format_name = map[int32]string{
		0: "NO_FORMAT",
		1: "IOI",
		2: "ICPC",
	}
	Contest_Format_value = map[string]int32{
		"NO_FORMAT": 0,
		"IOI":       1,
		"ICPC":      2,
	}
)

func (x Contest_Format) Enum() *Contest_Format {
	p := new(Contest_Format)
	*p = x
	return p
}

func (x Contest_Format) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Contest_Format) Descriptor() protoreflect.EnumDescriptor {
	return file_eolymp_judge_contest_proto_enumTypes[3].Descriptor()
}

func (Contest_Format) Type() protoreflect.EnumType {
	return &file_eolymp_judge_contest_proto_enumTypes[3]
}

func (x Contest_Format) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Contest_Format.Descriptor instead.
func (Contest_Format) EnumDescriptor() ([]byte, []int) {
	return file_eolymp_judge_contest_proto_rawDescGZIP(), []int{0, 3}
}

type Contest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Contest unique identifier, automatically allocated when contest is created.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Contest name.
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// Contest starting time, after this time users will be able to see problems and make submissions.
	StartsAt *timestamp.Timestamp `protobuf:"bytes,10,opt,name=starts_at,json=startsAt,proto3" json:"starts_at,omitempty"`
	// Time in seconds until beginning of the contest
	StartsIn uint32 `protobuf:"varint,11,opt,name=starts_in,json=startsIn,proto3" json:"starts_in,omitempty"`
	// Contest ending time, after this time users submissions won't be counted to the score anymore.
	EndsAt *timestamp.Timestamp `protobuf:"bytes,15,opt,name=ends_at,json=endsAt,proto3" json:"ends_at,omitempty"`
	// Time in seconds until end of the contest
	EndsIn uint32 `protobuf:"varint,16,opt,name=ends_in,json=endsIn,proto3" json:"ends_in,omitempty"`
	// Duration in seconds for virtual participation mode. Users can start participating at any moment between startsAt
	// and endsAt, but once started they would have to finish in Duration amount of time. For example contest maybe will
	// start at midnight and finish at 23:59:59 the same day, but if duration is set to 4h it would mean users can
	// participate at any point of the day but will have only 4 hours to solve problems (4h or before 23:59:59 whichever
	// happens first).
	Duration uint32 `protobuf:"varint,12,opt,name=duration,proto3" json:"duration,omitempty"`
	// Contest status (see statuses above)
	Status Contest_Status `protobuf:"varint,20,opt,name=status,proto3,enum=eolymp.judge.Contest_Status" json:"status,omitempty"`
	// Contest visibility defines who can participate and where contest is listed.
	Visibility Contest_Visibility `protobuf:"varint,30,opt,name=visibility,proto3,enum=eolymp.judge.Contest_Visibility" json:"visibility,omitempty"`
	// Participation mode defines timeframe for participation: online or virtual.
	ParticipationMode Contest_ParticipationMode `protobuf:"varint,31,opt,name=participation_mode,json=participationMode,proto3,enum=eolymp.judge.Contest_ParticipationMode" json:"participation_mode,omitempty"`
	// Format defines competition style IOI or ICPC.
	Format Contest_Format `protobuf:"varint,32,opt,name=format,proto3,enum=eolymp.judge.Contest_Format" json:"format,omitempty"`
	// Domain for contest, used to lookup for contest by domain name.
	Domain string `protobuf:"bytes,40,opt,name=domain,proto3" json:"domain,omitempty"`
	// Deprecated, space where contest was created, should be avoided.
	SpaceId string `protobuf:"bytes,1000,opt,name=space_id,json=spaceId,proto3" json:"space_id,omitempty"`
}

func (x *Contest) Reset() {
	*x = Contest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_judge_contest_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Contest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Contest) ProtoMessage() {}

func (x *Contest) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_judge_contest_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Contest.ProtoReflect.Descriptor instead.
func (*Contest) Descriptor() ([]byte, []int) {
	return file_eolymp_judge_contest_proto_rawDescGZIP(), []int{0}
}

func (x *Contest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Contest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Contest) GetStartsAt() *timestamp.Timestamp {
	if x != nil {
		return x.StartsAt
	}
	return nil
}

func (x *Contest) GetStartsIn() uint32 {
	if x != nil {
		return x.StartsIn
	}
	return 0
}

func (x *Contest) GetEndsAt() *timestamp.Timestamp {
	if x != nil {
		return x.EndsAt
	}
	return nil
}

func (x *Contest) GetEndsIn() uint32 {
	if x != nil {
		return x.EndsIn
	}
	return 0
}

func (x *Contest) GetDuration() uint32 {
	if x != nil {
		return x.Duration
	}
	return 0
}

func (x *Contest) GetStatus() Contest_Status {
	if x != nil {
		return x.Status
	}
	return Contest_NO_STATUS
}

func (x *Contest) GetVisibility() Contest_Visibility {
	if x != nil {
		return x.Visibility
	}
	return Contest_NO_VISIBILITY
}

func (x *Contest) GetParticipationMode() Contest_ParticipationMode {
	if x != nil {
		return x.ParticipationMode
	}
	return Contest_NO_PARTICIPATION
}

func (x *Contest) GetFormat() Contest_Format {
	if x != nil {
		return x.Format
	}
	return Contest_NO_FORMAT
}

func (x *Contest) GetDomain() string {
	if x != nil {
		return x.Domain
	}
	return ""
}

func (x *Contest) GetSpaceId() string {
	if x != nil {
		return x.SpaceId
	}
	return ""
}

type Contest_Appearance struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title          string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Tagline        string `protobuf:"bytes,2,opt,name=tagline,proto3" json:"tagline,omitempty"`
	LogoImage      string `protobuf:"bytes,3,opt,name=logo_image,json=logoImage,proto3" json:"logo_image,omitempty"`
	PrimaryColor   string `protobuf:"bytes,4,opt,name=primary_color,json=primaryColor,proto3" json:"primary_color,omitempty"`
	SecondaryColor string `protobuf:"bytes,5,opt,name=secondary_color,json=secondaryColor,proto3" json:"secondary_color,omitempty"`
}

func (x *Contest_Appearance) Reset() {
	*x = Contest_Appearance{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_judge_contest_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Contest_Appearance) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Contest_Appearance) ProtoMessage() {}

func (x *Contest_Appearance) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_judge_contest_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Contest_Appearance.ProtoReflect.Descriptor instead.
func (*Contest_Appearance) Descriptor() ([]byte, []int) {
	return file_eolymp_judge_contest_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Contest_Appearance) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Contest_Appearance) GetTagline() string {
	if x != nil {
		return x.Tagline
	}
	return ""
}

func (x *Contest_Appearance) GetLogoImage() string {
	if x != nil {
		return x.LogoImage
	}
	return ""
}

func (x *Contest_Appearance) GetPrimaryColor() string {
	if x != nil {
		return x.PrimaryColor
	}
	return ""
}

func (x *Contest_Appearance) GetSecondaryColor() string {
	if x != nil {
		return x.SecondaryColor
	}
	return ""
}

type Contest_Scoring struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ShowScoreboard bool   `protobuf:"varint,1,opt,name=show_scoreboard,json=showScoreboard,proto3" json:"show_scoreboard,omitempty"`
	AttemptPenalty uint32 `protobuf:"varint,2,opt,name=attempt_penalty,json=attemptPenalty,proto3" json:"attempt_penalty,omitempty"`
	FreezingTime   uint32 `protobuf:"varint,3,opt,name=freezing_time,json=freezingTime,proto3" json:"freezing_time,omitempty"`
}

func (x *Contest_Scoring) Reset() {
	*x = Contest_Scoring{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_judge_contest_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Contest_Scoring) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Contest_Scoring) ProtoMessage() {}

func (x *Contest_Scoring) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_judge_contest_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Contest_Scoring.ProtoReflect.Descriptor instead.
func (*Contest_Scoring) Descriptor() ([]byte, []int) {
	return file_eolymp_judge_contest_proto_rawDescGZIP(), []int{0, 1}
}

func (x *Contest_Scoring) GetShowScoreboard() bool {
	if x != nil {
		return x.ShowScoreboard
	}
	return false
}

func (x *Contest_Scoring) GetAttemptPenalty() uint32 {
	if x != nil {
		return x.AttemptPenalty
	}
	return 0
}

func (x *Contest_Scoring) GetFreezingTime() uint32 {
	if x != nil {
		return x.FreezingTime
	}
	return 0
}

var File_eolymp_judge_contest_proto protoreflect.FileDescriptor

var file_eolymp_judge_contest_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x6a, 0x75, 0x64, 0x67, 0x65, 0x2f, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x65, 0x6f,
	0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x6a, 0x75, 0x64, 0x67, 0x65, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xce, 0x08, 0x0a, 0x07,
	0x43, 0x6f, 0x6e, 0x74, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x37, 0x0a, 0x09, 0x73,
	0x74, 0x61, 0x72, 0x74, 0x73, 0x5f, 0x61, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x08, 0x73, 0x74, 0x61, 0x72,
	0x74, 0x73, 0x41, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x73, 0x5f, 0x69,
	0x6e, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x73, 0x74, 0x61, 0x72, 0x74, 0x73, 0x49,
	0x6e, 0x12, 0x33, 0x0a, 0x07, 0x65, 0x6e, 0x64, 0x73, 0x5f, 0x61, 0x74, 0x18, 0x0f, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x06,
	0x65, 0x6e, 0x64, 0x73, 0x41, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x65, 0x6e, 0x64, 0x73, 0x5f, 0x69,
	0x6e, 0x18, 0x10, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x65, 0x6e, 0x64, 0x73, 0x49, 0x6e, 0x12,
	0x1a, 0x0a, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x0c, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x34, 0x0a, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x14, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1c, 0x2e, 0x65, 0x6f,
	0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x6a, 0x75, 0x64, 0x67, 0x65, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x65,
	0x73, 0x74, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x40, 0x0a, 0x0a, 0x76, 0x69, 0x73, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x18,
	0x1e, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x20, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x6a,
	0x75, 0x64, 0x67, 0x65, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x56, 0x69, 0x73,
	0x69, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x52, 0x0a, 0x76, 0x69, 0x73, 0x69, 0x62, 0x69, 0x6c,
	0x69, 0x74, 0x79, 0x12, 0x56, 0x0a, 0x12, 0x70, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x18, 0x1f, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x27, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x6a, 0x75, 0x64, 0x67, 0x65, 0x2e, 0x43,
	0x6f, 0x6e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x50, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x6f, 0x64, 0x65, 0x52, 0x11, 0x70, 0x61, 0x72, 0x74, 0x69, 0x63,
	0x69, 0x70, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x6f, 0x64, 0x65, 0x12, 0x34, 0x0a, 0x06, 0x66,
	0x6f, 0x72, 0x6d, 0x61, 0x74, 0x18, 0x20, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1c, 0x2e, 0x65, 0x6f,
	0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x6a, 0x75, 0x64, 0x67, 0x65, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x65,
	0x73, 0x74, 0x2e, 0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x52, 0x06, 0x66, 0x6f, 0x72, 0x6d, 0x61,
	0x74, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x18, 0x28, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0xe8, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x70,
	0x61, 0x63, 0x65, 0x49, 0x64, 0x1a, 0xa9, 0x01, 0x0a, 0x0a, 0x41, 0x70, 0x70, 0x65, 0x61, 0x72,
	0x61, 0x6e, 0x63, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x74, 0x61,
	0x67, 0x6c, 0x69, 0x6e, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x74, 0x61, 0x67,
	0x6c, 0x69, 0x6e, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x6c, 0x6f, 0x67, 0x6f, 0x5f, 0x69, 0x6d, 0x61,
	0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6c, 0x6f, 0x67, 0x6f, 0x49, 0x6d,
	0x61, 0x67, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x70, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x5f, 0x63,
	0x6f, 0x6c, 0x6f, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x70, 0x72, 0x69, 0x6d,
	0x61, 0x72, 0x79, 0x43, 0x6f, 0x6c, 0x6f, 0x72, 0x12, 0x27, 0x0a, 0x0f, 0x73, 0x65, 0x63, 0x6f,
	0x6e, 0x64, 0x61, 0x72, 0x79, 0x5f, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0e, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x61, 0x72, 0x79, 0x43, 0x6f, 0x6c, 0x6f,
	0x72, 0x1a, 0x80, 0x01, 0x0a, 0x07, 0x53, 0x63, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x12, 0x27, 0x0a,
	0x0f, 0x73, 0x68, 0x6f, 0x77, 0x5f, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x62, 0x6f, 0x61, 0x72, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0e, 0x73, 0x68, 0x6f, 0x77, 0x53, 0x63, 0x6f, 0x72,
	0x65, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x12, 0x27, 0x0a, 0x0f, 0x61, 0x74, 0x74, 0x65, 0x6d, 0x70,
	0x74, 0x5f, 0x70, 0x65, 0x6e, 0x61, 0x6c, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x0e, 0x61, 0x74, 0x74, 0x65, 0x6d, 0x70, 0x74, 0x50, 0x65, 0x6e, 0x61, 0x6c, 0x74, 0x79, 0x12,
	0x23, 0x0a, 0x0d, 0x66, 0x72, 0x65, 0x65, 0x7a, 0x69, 0x6e, 0x67, 0x5f, 0x74, 0x69, 0x6d, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c, 0x66, 0x72, 0x65, 0x65, 0x7a, 0x69, 0x6e, 0x67,
	0x54, 0x69, 0x6d, 0x65, 0x22, 0x3e, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0d,
	0x0a, 0x09, 0x4e, 0x4f, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x10, 0x00, 0x12, 0x0d, 0x0a,
	0x09, 0x53, 0x43, 0x48, 0x45, 0x44, 0x55, 0x4c, 0x45, 0x44, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04,
	0x4f, 0x50, 0x45, 0x4e, 0x10, 0x02, 0x12, 0x0c, 0x0a, 0x08, 0x43, 0x4f, 0x4d, 0x50, 0x4c, 0x45,
	0x54, 0x45, 0x10, 0x03, 0x22, 0x46, 0x0a, 0x0a, 0x56, 0x69, 0x73, 0x69, 0x62, 0x69, 0x6c, 0x69,
	0x74, 0x79, 0x12, 0x11, 0x0a, 0x0d, 0x4e, 0x4f, 0x5f, 0x56, 0x49, 0x53, 0x49, 0x42, 0x49, 0x4c,
	0x49, 0x54, 0x59, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x50, 0x55, 0x42, 0x4c, 0x49, 0x43, 0x10,
	0x01, 0x12, 0x0c, 0x0a, 0x08, 0x55, 0x4e, 0x4c, 0x49, 0x53, 0x54, 0x45, 0x44, 0x10, 0x02, 0x12,
	0x0b, 0x0a, 0x07, 0x50, 0x52, 0x49, 0x56, 0x41, 0x54, 0x45, 0x10, 0x03, 0x22, 0x42, 0x0a, 0x11,
	0x50, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x6f, 0x64,
	0x65, 0x12, 0x14, 0x0a, 0x10, 0x4e, 0x4f, 0x5f, 0x50, 0x41, 0x52, 0x54, 0x49, 0x43, 0x49, 0x50,
	0x41, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x4f, 0x4e, 0x4c, 0x49, 0x4e,
	0x45, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x56, 0x49, 0x52, 0x54, 0x55, 0x41, 0x4c, 0x10, 0x02,
	0x22, 0x2a, 0x0a, 0x06, 0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x12, 0x0d, 0x0a, 0x09, 0x4e, 0x4f,
	0x5f, 0x46, 0x4f, 0x52, 0x4d, 0x41, 0x54, 0x10, 0x00, 0x12, 0x07, 0x0a, 0x03, 0x49, 0x4f, 0x49,
	0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x49, 0x43, 0x50, 0x43, 0x10, 0x02, 0x42, 0x2d, 0x5a, 0x2b,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6f, 0x6c, 0x79, 0x6d,
	0x70, 0x2f, 0x67, 0x6f, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f,
	0x6a, 0x75, 0x64, 0x67, 0x65, 0x3b, 0x6a, 0x75, 0x64, 0x67, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_eolymp_judge_contest_proto_rawDescOnce sync.Once
	file_eolymp_judge_contest_proto_rawDescData = file_eolymp_judge_contest_proto_rawDesc
)

func file_eolymp_judge_contest_proto_rawDescGZIP() []byte {
	file_eolymp_judge_contest_proto_rawDescOnce.Do(func() {
		file_eolymp_judge_contest_proto_rawDescData = protoimpl.X.CompressGZIP(file_eolymp_judge_contest_proto_rawDescData)
	})
	return file_eolymp_judge_contest_proto_rawDescData
}

var file_eolymp_judge_contest_proto_enumTypes = make([]protoimpl.EnumInfo, 4)
var file_eolymp_judge_contest_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_eolymp_judge_contest_proto_goTypes = []interface{}{
	(Contest_Status)(0),            // 0: eolymp.judge.Contest.Status
	(Contest_Visibility)(0),        // 1: eolymp.judge.Contest.Visibility
	(Contest_ParticipationMode)(0), // 2: eolymp.judge.Contest.ParticipationMode
	(Contest_Format)(0),            // 3: eolymp.judge.Contest.Format
	(*Contest)(nil),                // 4: eolymp.judge.Contest
	(*Contest_Appearance)(nil),     // 5: eolymp.judge.Contest.Appearance
	(*Contest_Scoring)(nil),        // 6: eolymp.judge.Contest.Scoring
	(*timestamp.Timestamp)(nil),    // 7: google.protobuf.Timestamp
}
var file_eolymp_judge_contest_proto_depIdxs = []int32{
	7, // 0: eolymp.judge.Contest.starts_at:type_name -> google.protobuf.Timestamp
	7, // 1: eolymp.judge.Contest.ends_at:type_name -> google.protobuf.Timestamp
	0, // 2: eolymp.judge.Contest.status:type_name -> eolymp.judge.Contest.Status
	1, // 3: eolymp.judge.Contest.visibility:type_name -> eolymp.judge.Contest.Visibility
	2, // 4: eolymp.judge.Contest.participation_mode:type_name -> eolymp.judge.Contest.ParticipationMode
	3, // 5: eolymp.judge.Contest.format:type_name -> eolymp.judge.Contest.Format
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_eolymp_judge_contest_proto_init() }
func file_eolymp_judge_contest_proto_init() {
	if File_eolymp_judge_contest_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_eolymp_judge_contest_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Contest); i {
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
		file_eolymp_judge_contest_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Contest_Appearance); i {
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
		file_eolymp_judge_contest_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Contest_Scoring); i {
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
			RawDescriptor: file_eolymp_judge_contest_proto_rawDesc,
			NumEnums:      4,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_eolymp_judge_contest_proto_goTypes,
		DependencyIndexes: file_eolymp_judge_contest_proto_depIdxs,
		EnumInfos:         file_eolymp_judge_contest_proto_enumTypes,
		MessageInfos:      file_eolymp_judge_contest_proto_msgTypes,
	}.Build()
	File_eolymp_judge_contest_proto = out.File
	file_eolymp_judge_contest_proto_rawDesc = nil
	file_eolymp_judge_contest_proto_goTypes = nil
	file_eolymp_judge_contest_proto_depIdxs = nil
}
