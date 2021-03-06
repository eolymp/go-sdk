// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.18.1
// source: eolymp/ranker/scoreboard.proto

package ranker

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

type Scoreboard_Column_Type int32

const (
	Scoreboard_Column_NONE      Scoreboard_Column_Type = 0
	Scoreboard_Column_CONTEST   Scoreboard_Column_Type = 1
	Scoreboard_Column_PROBLEM   Scoreboard_Column_Type = 2
	Scoreboard_Column_ATTRIBUTE Scoreboard_Column_Type = 3
	Scoreboard_Column_NAME      Scoreboard_Column_Type = 4
)

// Enum value maps for Scoreboard_Column_Type.
var (
	Scoreboard_Column_Type_name = map[int32]string{
		0: "NONE",
		1: "CONTEST",
		2: "PROBLEM",
		3: "ATTRIBUTE",
		4: "NAME",
	}
	Scoreboard_Column_Type_value = map[string]int32{
		"NONE":      0,
		"CONTEST":   1,
		"PROBLEM":   2,
		"ATTRIBUTE": 3,
		"NAME":      4,
	}
)

func (x Scoreboard_Column_Type) Enum() *Scoreboard_Column_Type {
	p := new(Scoreboard_Column_Type)
	*p = x
	return p
}

func (x Scoreboard_Column_Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Scoreboard_Column_Type) Descriptor() protoreflect.EnumDescriptor {
	return file_eolymp_ranker_scoreboard_proto_enumTypes[0].Descriptor()
}

func (Scoreboard_Column_Type) Type() protoreflect.EnumType {
	return &file_eolymp_ranker_scoreboard_proto_enumTypes[0]
}

func (x Scoreboard_Column_Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Scoreboard_Column_Type.Descriptor instead.
func (Scoreboard_Column_Type) EnumDescriptor() ([]byte, []int) {
	return file_eolymp_ranker_scoreboard_proto_rawDescGZIP(), []int{0, 1, 0}
}

type Scoreboard struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Unique identifier. This field is readonly, it's assigned automatically when scoreboard is returned.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// A user friendly identifier for the scoreboard, used as part of the domain name for the scoreboard.
	Key string `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	// A user friendly name for the scoreboard.
	Name   string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Format Format `protobuf:"varint,10,opt,name=format,proto3,enum=eolymp.ranker.Format" json:"format,omitempty"`
}

func (x *Scoreboard) Reset() {
	*x = Scoreboard{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_ranker_scoreboard_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Scoreboard) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Scoreboard) ProtoMessage() {}

func (x *Scoreboard) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_ranker_scoreboard_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Scoreboard.ProtoReflect.Descriptor instead.
func (*Scoreboard) Descriptor() ([]byte, []int) {
	return file_eolymp_ranker_scoreboard_proto_rawDescGZIP(), []int{0}
}

func (x *Scoreboard) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Scoreboard) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *Scoreboard) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Scoreboard) GetFormat() Format {
	if x != nil {
		return x.Format
	}
	return Format_NONE
}

type Scoreboard_Row struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string                  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`                                  // Row unique identifier
	Name      string                  `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`                              // Participant name (identifier).
	MemberId  string                  `protobuf:"bytes,3,opt,name=member_id,json=memberId,proto3" json:"member_id,omitempty"`      // Member ID.
	Score     float32                 `protobuf:"fixed32,10,opt,name=score,proto3" json:"score,omitempty"`                         // Total score.
	Penalty   float32                 `protobuf:"fixed32,11,opt,name=penalty,proto3" json:"penalty,omitempty"`                     // Total penalty.
	Rank      uint32                  `protobuf:"varint,30,opt,name=rank,proto3" json:"rank,omitempty"`                            // Rank in the scoreboard
	RankLower uint32                  `protobuf:"varint,31,opt,name=rank_lower,json=rankLower,proto3" json:"rank_lower,omitempty"` // Lower bound of the rank (when shared)
	Values    []*Scoreboard_Row_Value `protobuf:"bytes,20,rep,name=values,proto3" json:"values,omitempty"`                         // Score breakdown by contest and problem.
}

func (x *Scoreboard_Row) Reset() {
	*x = Scoreboard_Row{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_ranker_scoreboard_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Scoreboard_Row) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Scoreboard_Row) ProtoMessage() {}

func (x *Scoreboard_Row) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_ranker_scoreboard_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Scoreboard_Row.ProtoReflect.Descriptor instead.
func (*Scoreboard_Row) Descriptor() ([]byte, []int) {
	return file_eolymp_ranker_scoreboard_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Scoreboard_Row) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Scoreboard_Row) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Scoreboard_Row) GetMemberId() string {
	if x != nil {
		return x.MemberId
	}
	return ""
}

func (x *Scoreboard_Row) GetScore() float32 {
	if x != nil {
		return x.Score
	}
	return 0
}

func (x *Scoreboard_Row) GetPenalty() float32 {
	if x != nil {
		return x.Penalty
	}
	return 0
}

func (x *Scoreboard_Row) GetRank() uint32 {
	if x != nil {
		return x.Rank
	}
	return 0
}

func (x *Scoreboard_Row) GetRankLower() uint32 {
	if x != nil {
		return x.RankLower
	}
	return 0
}

func (x *Scoreboard_Row) GetValues() []*Scoreboard_Row_Value {
	if x != nil {
		return x.Values
	}
	return nil
}

type Scoreboard_Column struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                    string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	ParentId              string                 `protobuf:"bytes,2,opt,name=parent_id,json=parentId,proto3" json:"parent_id,omitempty"`
	Key                   string                 `protobuf:"bytes,3,opt,name=key,proto3" json:"key,omitempty"`
	Name                  string                 `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	ShortName             string                 `protobuf:"bytes,5,opt,name=short_name,json=shortName,proto3" json:"short_name,omitempty"`
	Type                  Scoreboard_Column_Type `protobuf:"varint,6,opt,name=type,proto3,enum=eolymp.ranker.Scoreboard_Column_Type" json:"type,omitempty"`
	Index                 uint32                 `protobuf:"varint,7,opt,name=index,proto3" json:"index,omitempty"`
	Visible               bool                   `protobuf:"varint,8,opt,name=visible,proto3" json:"visible,omitempty"`
	JudgeContestId        string                 `protobuf:"bytes,20,opt,name=judge_contest_id,json=judgeContestId,proto3" json:"judge_contest_id,omitempty"`
	JudgeProblemId        string                 `protobuf:"bytes,21,opt,name=judge_problem_id,json=judgeProblemId,proto3" json:"judge_problem_id,omitempty"`
	CommunityAttributeKey string                 `protobuf:"bytes,22,opt,name=community_attribute_key,json=communityAttributeKey,proto3" json:"community_attribute_key,omitempty"`
	Columns               []*Scoreboard_Column   `protobuf:"bytes,100,rep,name=columns,proto3" json:"columns,omitempty"`
}

func (x *Scoreboard_Column) Reset() {
	*x = Scoreboard_Column{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_ranker_scoreboard_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Scoreboard_Column) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Scoreboard_Column) ProtoMessage() {}

func (x *Scoreboard_Column) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_ranker_scoreboard_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Scoreboard_Column.ProtoReflect.Descriptor instead.
func (*Scoreboard_Column) Descriptor() ([]byte, []int) {
	return file_eolymp_ranker_scoreboard_proto_rawDescGZIP(), []int{0, 1}
}

func (x *Scoreboard_Column) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Scoreboard_Column) GetParentId() string {
	if x != nil {
		return x.ParentId
	}
	return ""
}

func (x *Scoreboard_Column) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *Scoreboard_Column) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Scoreboard_Column) GetShortName() string {
	if x != nil {
		return x.ShortName
	}
	return ""
}

func (x *Scoreboard_Column) GetType() Scoreboard_Column_Type {
	if x != nil {
		return x.Type
	}
	return Scoreboard_Column_NONE
}

func (x *Scoreboard_Column) GetIndex() uint32 {
	if x != nil {
		return x.Index
	}
	return 0
}

func (x *Scoreboard_Column) GetVisible() bool {
	if x != nil {
		return x.Visible
	}
	return false
}

func (x *Scoreboard_Column) GetJudgeContestId() string {
	if x != nil {
		return x.JudgeContestId
	}
	return ""
}

func (x *Scoreboard_Column) GetJudgeProblemId() string {
	if x != nil {
		return x.JudgeProblemId
	}
	return ""
}

func (x *Scoreboard_Column) GetCommunityAttributeKey() string {
	if x != nil {
		return x.CommunityAttributeKey
	}
	return ""
}

func (x *Scoreboard_Column) GetColumns() []*Scoreboard_Column {
	if x != nil {
		return x.Columns
	}
	return nil
}

type Scoreboard_Row_Value struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	ColumnId    string  `protobuf:"bytes,2,opt,name=column_id,json=columnId,proto3" json:"column_id,omitempty"`
	Score       float32 `protobuf:"fixed32,10,opt,name=score,proto3" json:"score,omitempty"`
	Penalty     float32 `protobuf:"fixed32,11,opt,name=penalty,proto3" json:"penalty,omitempty"`
	Percentage  float32 `protobuf:"fixed32,12,opt,name=percentage,proto3" json:"percentage,omitempty"`            // Percentage of points scored, from 0 to 1.
	Attempts    uint32  `protobuf:"varint,13,opt,name=attempts,proto3" json:"attempts,omitempty"`                 // Number of attempts before problem was solved.
	SolvedIn    uint32  `protobuf:"varint,14,opt,name=solved_in,json=solvedIn,proto3" json:"solved_in,omitempty"` // How much time it took to solve problem since beginning of the contest, in seconds.
	ValueString string  `protobuf:"bytes,20,opt,name=value_string,json=valueString,proto3" json:"value_string,omitempty"`
	ValueNumber int32   `protobuf:"varint,21,opt,name=value_number,json=valueNumber,proto3" json:"value_number,omitempty"`
}

func (x *Scoreboard_Row_Value) Reset() {
	*x = Scoreboard_Row_Value{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_ranker_scoreboard_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Scoreboard_Row_Value) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Scoreboard_Row_Value) ProtoMessage() {}

func (x *Scoreboard_Row_Value) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_ranker_scoreboard_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Scoreboard_Row_Value.ProtoReflect.Descriptor instead.
func (*Scoreboard_Row_Value) Descriptor() ([]byte, []int) {
	return file_eolymp_ranker_scoreboard_proto_rawDescGZIP(), []int{0, 0, 0}
}

func (x *Scoreboard_Row_Value) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Scoreboard_Row_Value) GetColumnId() string {
	if x != nil {
		return x.ColumnId
	}
	return ""
}

func (x *Scoreboard_Row_Value) GetScore() float32 {
	if x != nil {
		return x.Score
	}
	return 0
}

func (x *Scoreboard_Row_Value) GetPenalty() float32 {
	if x != nil {
		return x.Penalty
	}
	return 0
}

func (x *Scoreboard_Row_Value) GetPercentage() float32 {
	if x != nil {
		return x.Percentage
	}
	return 0
}

func (x *Scoreboard_Row_Value) GetAttempts() uint32 {
	if x != nil {
		return x.Attempts
	}
	return 0
}

func (x *Scoreboard_Row_Value) GetSolvedIn() uint32 {
	if x != nil {
		return x.SolvedIn
	}
	return 0
}

func (x *Scoreboard_Row_Value) GetValueString() string {
	if x != nil {
		return x.ValueString
	}
	return ""
}

func (x *Scoreboard_Row_Value) GetValueNumber() int32 {
	if x != nil {
		return x.ValueNumber
	}
	return 0
}

var File_eolymp_ranker_scoreboard_proto protoreflect.FileDescriptor

var file_eolymp_ranker_scoreboard_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x72, 0x61, 0x6e, 0x6b, 0x65, 0x72, 0x2f,
	0x73, 0x63, 0x6f, 0x72, 0x65, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0d, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x72, 0x61, 0x6e, 0x6b, 0x65, 0x72, 0x1a,
	0x1a, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x72, 0x61, 0x6e, 0x6b, 0x65, 0x72, 0x2f, 0x66,
	0x6f, 0x72, 0x6d, 0x61, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd5, 0x08, 0x0a, 0x0a,
	0x53, 0x63, 0x6f, 0x72, 0x65, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x2d, 0x0a, 0x06, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x15, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x72, 0x61, 0x6e, 0x6b, 0x65, 0x72,
	0x2e, 0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x52, 0x06, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x1a,
	0xec, 0x03, 0x0a, 0x03, 0x52, 0x6f, 0x77, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6d,
	0x65, 0x6d, 0x62, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x63, 0x6f, 0x72,
	0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x70, 0x65, 0x6e, 0x61, 0x6c, 0x74, 0x79, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x02, 0x52,
	0x07, 0x70, 0x65, 0x6e, 0x61, 0x6c, 0x74, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x61, 0x6e, 0x6b,
	0x18, 0x1e, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x72, 0x61, 0x6e, 0x6b, 0x12, 0x1d, 0x0a, 0x0a,
	0x72, 0x61, 0x6e, 0x6b, 0x5f, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x18, 0x1f, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x09, 0x72, 0x61, 0x6e, 0x6b, 0x4c, 0x6f, 0x77, 0x65, 0x72, 0x12, 0x3b, 0x0a, 0x06, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x14, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x65, 0x6f,
	0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x72, 0x61, 0x6e, 0x6b, 0x65, 0x72, 0x2e, 0x53, 0x63, 0x6f, 0x72,
	0x65, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x52, 0x6f, 0x77, 0x2e, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x52, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x1a, 0x83, 0x02, 0x0a, 0x05, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x5f, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x49, 0x64, 0x12,
	0x14, 0x0a, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05,
	0x73, 0x63, 0x6f, 0x72, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x65, 0x6e, 0x61, 0x6c, 0x74, 0x79,
	0x18, 0x0b, 0x20, 0x01, 0x28, 0x02, 0x52, 0x07, 0x70, 0x65, 0x6e, 0x61, 0x6c, 0x74, 0x79, 0x12,
	0x1e, 0x0a, 0x0a, 0x70, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x61, 0x67, 0x65, 0x18, 0x0c, 0x20,
	0x01, 0x28, 0x02, 0x52, 0x0a, 0x70, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x61, 0x67, 0x65, 0x12,
	0x1a, 0x0a, 0x08, 0x61, 0x74, 0x74, 0x65, 0x6d, 0x70, 0x74, 0x73, 0x18, 0x0d, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x08, 0x61, 0x74, 0x74, 0x65, 0x6d, 0x70, 0x74, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x73,
	0x6f, 0x6c, 0x76, 0x65, 0x64, 0x5f, 0x69, 0x6e, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08,
	0x73, 0x6f, 0x6c, 0x76, 0x65, 0x64, 0x49, 0x6e, 0x12, 0x21, 0x0a, 0x0c, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x5f, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x12, 0x21, 0x0a, 0x0c, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x15, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x0b, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x1a, 0xf2,
	0x03, 0x0a, 0x06, 0x43, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x72,
	0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61,
	0x72, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a,
	0x73, 0x68, 0x6f, 0x72, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x39, 0x0a, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x25, 0x2e, 0x65, 0x6f, 0x6c, 0x79,
	0x6d, 0x70, 0x2e, 0x72, 0x61, 0x6e, 0x6b, 0x65, 0x72, 0x2e, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x62,
	0x6f, 0x61, 0x72, 0x64, 0x2e, 0x43, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x2e, 0x54, 0x79, 0x70, 0x65,
	0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x18, 0x0a, 0x07,
	0x76, 0x69, 0x73, 0x69, 0x62, 0x6c, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x76,
	0x69, 0x73, 0x69, 0x62, 0x6c, 0x65, 0x12, 0x28, 0x0a, 0x10, 0x6a, 0x75, 0x64, 0x67, 0x65, 0x5f,
	0x63, 0x6f, 0x6e, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0e, 0x6a, 0x75, 0x64, 0x67, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x73, 0x74, 0x49, 0x64,
	0x12, 0x28, 0x0a, 0x10, 0x6a, 0x75, 0x64, 0x67, 0x65, 0x5f, 0x70, 0x72, 0x6f, 0x62, 0x6c, 0x65,
	0x6d, 0x5f, 0x69, 0x64, 0x18, 0x15, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x6a, 0x75, 0x64, 0x67,
	0x65, 0x50, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x49, 0x64, 0x12, 0x36, 0x0a, 0x17, 0x63, 0x6f,
	0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x5f, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74,
	0x65, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x16, 0x20, 0x01, 0x28, 0x09, 0x52, 0x15, 0x63, 0x6f, 0x6d,
	0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x4b,
	0x65, 0x79, 0x12, 0x3a, 0x0a, 0x07, 0x63, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x73, 0x18, 0x64, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x72, 0x61, 0x6e,
	0x6b, 0x65, 0x72, 0x2e, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x43,
	0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x52, 0x07, 0x63, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x73, 0x22, 0x43,
	0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00,
	0x12, 0x0b, 0x0a, 0x07, 0x43, 0x4f, 0x4e, 0x54, 0x45, 0x53, 0x54, 0x10, 0x01, 0x12, 0x0b, 0x0a,
	0x07, 0x50, 0x52, 0x4f, 0x42, 0x4c, 0x45, 0x4d, 0x10, 0x02, 0x12, 0x0d, 0x0a, 0x09, 0x41, 0x54,
	0x54, 0x52, 0x49, 0x42, 0x55, 0x54, 0x45, 0x10, 0x03, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x41, 0x4d,
	0x45, 0x10, 0x04, 0x42, 0x2f, 0x5a, 0x2d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x67, 0x6f, 0x2d, 0x73, 0x64, 0x6b, 0x2f,
	0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x72, 0x61, 0x6e, 0x6b, 0x65, 0x72, 0x3b, 0x72, 0x61,
	0x6e, 0x6b, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_eolymp_ranker_scoreboard_proto_rawDescOnce sync.Once
	file_eolymp_ranker_scoreboard_proto_rawDescData = file_eolymp_ranker_scoreboard_proto_rawDesc
)

func file_eolymp_ranker_scoreboard_proto_rawDescGZIP() []byte {
	file_eolymp_ranker_scoreboard_proto_rawDescOnce.Do(func() {
		file_eolymp_ranker_scoreboard_proto_rawDescData = protoimpl.X.CompressGZIP(file_eolymp_ranker_scoreboard_proto_rawDescData)
	})
	return file_eolymp_ranker_scoreboard_proto_rawDescData
}

var file_eolymp_ranker_scoreboard_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_eolymp_ranker_scoreboard_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_eolymp_ranker_scoreboard_proto_goTypes = []interface{}{
	(Scoreboard_Column_Type)(0),  // 0: eolymp.ranker.Scoreboard.Column.Type
	(*Scoreboard)(nil),           // 1: eolymp.ranker.Scoreboard
	(*Scoreboard_Row)(nil),       // 2: eolymp.ranker.Scoreboard.Row
	(*Scoreboard_Column)(nil),    // 3: eolymp.ranker.Scoreboard.Column
	(*Scoreboard_Row_Value)(nil), // 4: eolymp.ranker.Scoreboard.Row.Value
	(Format)(0),                  // 5: eolymp.ranker.Format
}
var file_eolymp_ranker_scoreboard_proto_depIdxs = []int32{
	5, // 0: eolymp.ranker.Scoreboard.format:type_name -> eolymp.ranker.Format
	4, // 1: eolymp.ranker.Scoreboard.Row.values:type_name -> eolymp.ranker.Scoreboard.Row.Value
	0, // 2: eolymp.ranker.Scoreboard.Column.type:type_name -> eolymp.ranker.Scoreboard.Column.Type
	3, // 3: eolymp.ranker.Scoreboard.Column.columns:type_name -> eolymp.ranker.Scoreboard.Column
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_eolymp_ranker_scoreboard_proto_init() }
func file_eolymp_ranker_scoreboard_proto_init() {
	if File_eolymp_ranker_scoreboard_proto != nil {
		return
	}
	file_eolymp_ranker_format_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_eolymp_ranker_scoreboard_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Scoreboard); i {
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
		file_eolymp_ranker_scoreboard_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Scoreboard_Row); i {
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
		file_eolymp_ranker_scoreboard_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Scoreboard_Column); i {
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
		file_eolymp_ranker_scoreboard_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Scoreboard_Row_Value); i {
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
			RawDescriptor: file_eolymp_ranker_scoreboard_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_eolymp_ranker_scoreboard_proto_goTypes,
		DependencyIndexes: file_eolymp_ranker_scoreboard_proto_depIdxs,
		EnumInfos:         file_eolymp_ranker_scoreboard_proto_enumTypes,
		MessageInfos:      file_eolymp_ranker_scoreboard_proto_msgTypes,
	}.Build()
	File_eolymp_ranker_scoreboard_proto = out.File
	file_eolymp_ranker_scoreboard_proto_rawDesc = nil
	file_eolymp_ranker_scoreboard_proto_goTypes = nil
	file_eolymp_ranker_scoreboard_proto_depIdxs = nil
}
