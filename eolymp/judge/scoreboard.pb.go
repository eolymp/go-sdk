// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.1
// 	protoc        v4.24.4
// source: eolymp/judge/scoreboard.proto

package judge

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

type Scoreboard_Mode int32

const (
	Scoreboard_RESULT  Scoreboard_Mode = 0
	Scoreboard_FROZEN  Scoreboard_Mode = 1
	Scoreboard_UPSOLVE Scoreboard_Mode = 2
)

// Enum value maps for Scoreboard_Mode.
var (
	Scoreboard_Mode_name = map[int32]string{
		0: "RESULT",
		1: "FROZEN",
		2: "UPSOLVE",
	}
	Scoreboard_Mode_value = map[string]int32{
		"RESULT":  0,
		"FROZEN":  1,
		"UPSOLVE": 2,
	}
)

func (x Scoreboard_Mode) Enum() *Scoreboard_Mode {
	p := new(Scoreboard_Mode)
	*p = x
	return p
}

func (x Scoreboard_Mode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Scoreboard_Mode) Descriptor() protoreflect.EnumDescriptor {
	return file_eolymp_judge_scoreboard_proto_enumTypes[0].Descriptor()
}

func (Scoreboard_Mode) Type() protoreflect.EnumType {
	return &file_eolymp_judge_scoreboard_proto_enumTypes[0]
}

func (x Scoreboard_Mode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Scoreboard_Mode.Descriptor instead.
func (Scoreboard_Mode) EnumDescriptor() ([]byte, []int) {
	return file_eolymp_judge_scoreboard_proto_rawDescGZIP(), []int{0, 0}
}

type Scoreboard_Column_Type int32

const (
	Scoreboard_Column_UNKNOWN_TYPE  Scoreboard_Column_Type = 0
	Scoreboard_Column_ROUND_SCORE   Scoreboard_Column_Type = 2
	Scoreboard_Column_PROBLEM_SCORE Scoreboard_Column_Type = 3
	Scoreboard_Column_STRING        Scoreboard_Column_Type = 10
	Scoreboard_Column_NUMBER        Scoreboard_Column_Type = 11
	Scoreboard_Column_CHOICE        Scoreboard_Column_Type = 12
	Scoreboard_Column_DATE          Scoreboard_Column_Type = 13
	Scoreboard_Column_EMAIL         Scoreboard_Column_Type = 14
	Scoreboard_Column_CHECKBOX      Scoreboard_Column_Type = 15
	Scoreboard_Column_COUNTRY       Scoreboard_Column_Type = 16
	Scoreboard_Column_REGION        Scoreboard_Column_Type = 17
	Scoreboard_Column_INSTITUTION   Scoreboard_Column_Type = 18
)

// Enum value maps for Scoreboard_Column_Type.
var (
	Scoreboard_Column_Type_name = map[int32]string{
		0:  "UNKNOWN_TYPE",
		2:  "ROUND_SCORE",
		3:  "PROBLEM_SCORE",
		10: "STRING",
		11: "NUMBER",
		12: "CHOICE",
		13: "DATE",
		14: "EMAIL",
		15: "CHECKBOX",
		16: "COUNTRY",
		17: "REGION",
		18: "INSTITUTION",
	}
	Scoreboard_Column_Type_value = map[string]int32{
		"UNKNOWN_TYPE":  0,
		"ROUND_SCORE":   2,
		"PROBLEM_SCORE": 3,
		"STRING":        10,
		"NUMBER":        11,
		"CHOICE":        12,
		"DATE":          13,
		"EMAIL":         14,
		"CHECKBOX":      15,
		"COUNTRY":       16,
		"REGION":        17,
		"INSTITUTION":   18,
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
	return file_eolymp_judge_scoreboard_proto_enumTypes[1].Descriptor()
}

func (Scoreboard_Column_Type) Type() protoreflect.EnumType {
	return &file_eolymp_judge_scoreboard_proto_enumTypes[1]
}

func (x Scoreboard_Column_Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Scoreboard_Column_Type.Descriptor instead.
func (Scoreboard_Column_Type) EnumDescriptor() ([]byte, []int) {
	return file_eolymp_judge_scoreboard_proto_rawDescGZIP(), []int{0, 1, 0}
}

type Scoreboard struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Modes         []Scoreboard_Mode      `protobuf:"varint,10,rep,packed,name=modes,proto3,enum=eolymp.judge.Scoreboard_Mode" json:"modes,omitempty"`
	Rounds        []*Scoreboard_Round    `protobuf:"bytes,12,rep,name=rounds,proto3" json:"rounds,omitempty"`
	Columns       []*Scoreboard_Column   `protobuf:"bytes,11,rep,name=columns,proto3" json:"columns,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Scoreboard) Reset() {
	*x = Scoreboard{}
	mi := &file_eolymp_judge_scoreboard_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Scoreboard) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Scoreboard) ProtoMessage() {}

func (x *Scoreboard) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_judge_scoreboard_proto_msgTypes[0]
	if x != nil {
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
	return file_eolymp_judge_scoreboard_proto_rawDescGZIP(), []int{0}
}

func (x *Scoreboard) GetModes() []Scoreboard_Mode {
	if x != nil {
		return x.Modes
	}
	return nil
}

func (x *Scoreboard) GetRounds() []*Scoreboard_Round {
	if x != nil {
		return x.Rounds
	}
	return nil
}

func (x *Scoreboard) GetColumns() []*Scoreboard_Column {
	if x != nil {
		return x.Columns
	}
	return nil
}

type Scoreboard_Round struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Title         string                 `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Scoreboard_Round) Reset() {
	*x = Scoreboard_Round{}
	mi := &file_eolymp_judge_scoreboard_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Scoreboard_Round) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Scoreboard_Round) ProtoMessage() {}

func (x *Scoreboard_Round) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_judge_scoreboard_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Scoreboard_Round.ProtoReflect.Descriptor instead.
func (*Scoreboard_Round) Descriptor() ([]byte, []int) {
	return file_eolymp_judge_scoreboard_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Scoreboard_Round) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Scoreboard_Round) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

type Scoreboard_Column struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Type          Scoreboard_Column_Type `protobuf:"varint,2,opt,name=type,proto3,enum=eolymp.judge.Scoreboard_Column_Type" json:"type,omitempty"`
	Title         string                 `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	Choices       []string               `protobuf:"bytes,10,rep,name=choices,proto3" json:"choices,omitempty"`
	Sortable      bool                   `protobuf:"varint,20,opt,name=sortable,proto3" json:"sortable,omitempty"`
	Filterable    bool                   `protobuf:"varint,21,opt,name=filterable,proto3" json:"filterable,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Scoreboard_Column) Reset() {
	*x = Scoreboard_Column{}
	mi := &file_eolymp_judge_scoreboard_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Scoreboard_Column) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Scoreboard_Column) ProtoMessage() {}

func (x *Scoreboard_Column) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_judge_scoreboard_proto_msgTypes[2]
	if x != nil {
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
	return file_eolymp_judge_scoreboard_proto_rawDescGZIP(), []int{0, 1}
}

func (x *Scoreboard_Column) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Scoreboard_Column) GetType() Scoreboard_Column_Type {
	if x != nil {
		return x.Type
	}
	return Scoreboard_Column_UNKNOWN_TYPE
}

func (x *Scoreboard_Column) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Scoreboard_Column) GetChoices() []string {
	if x != nil {
		return x.Choices
	}
	return nil
}

func (x *Scoreboard_Column) GetSortable() bool {
	if x != nil {
		return x.Sortable
	}
	return false
}

func (x *Scoreboard_Column) GetFilterable() bool {
	if x != nil {
		return x.Filterable
	}
	return false
}

type Scoreboard_Row struct {
	state         protoimpl.MessageState  `protogen:"open.v1"`
	Id            string                  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	MemberId      string                  `protobuf:"bytes,2,opt,name=member_id,json=memberId,proto3" json:"member_id,omitempty"`
	Index         uint32                  `protobuf:"varint,10,opt,name=index,proto3" json:"index,omitempty"`
	Rank          uint32                  `protobuf:"varint,11,opt,name=rank,proto3" json:"rank,omitempty"`
	RankLength    uint32                  `protobuf:"varint,12,opt,name=rank_length,json=rankLength,proto3" json:"rank_length,omitempty"`
	Score         float32                 `protobuf:"fixed32,20,opt,name=score,proto3" json:"score,omitempty"`
	Penalty       float32                 `protobuf:"fixed32,21,opt,name=penalty,proto3" json:"penalty,omitempty"`
	TieBreaker    uint32                  `protobuf:"varint,22,opt,name=tie_breaker,json=tieBreaker,proto3" json:"tie_breaker,omitempty"`
	Unofficial    bool                    `protobuf:"varint,30,opt,name=unofficial,proto3" json:"unofficial,omitempty"`
	Disqualified  bool                    `protobuf:"varint,31,opt,name=disqualified,proto3" json:"disqualified,omitempty"`
	Medal         Medal                   `protobuf:"varint,32,opt,name=medal,proto3,enum=eolymp.judge.Medal" json:"medal,omitempty"`
	Values        []*Scoreboard_Row_Value `protobuf:"bytes,50,rep,name=values,proto3" json:"values,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Scoreboard_Row) Reset() {
	*x = Scoreboard_Row{}
	mi := &file_eolymp_judge_scoreboard_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Scoreboard_Row) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Scoreboard_Row) ProtoMessage() {}

func (x *Scoreboard_Row) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_judge_scoreboard_proto_msgTypes[3]
	if x != nil {
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
	return file_eolymp_judge_scoreboard_proto_rawDescGZIP(), []int{0, 2}
}

func (x *Scoreboard_Row) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Scoreboard_Row) GetMemberId() string {
	if x != nil {
		return x.MemberId
	}
	return ""
}

func (x *Scoreboard_Row) GetIndex() uint32 {
	if x != nil {
		return x.Index
	}
	return 0
}

func (x *Scoreboard_Row) GetRank() uint32 {
	if x != nil {
		return x.Rank
	}
	return 0
}

func (x *Scoreboard_Row) GetRankLength() uint32 {
	if x != nil {
		return x.RankLength
	}
	return 0
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

func (x *Scoreboard_Row) GetTieBreaker() uint32 {
	if x != nil {
		return x.TieBreaker
	}
	return 0
}

func (x *Scoreboard_Row) GetUnofficial() bool {
	if x != nil {
		return x.Unofficial
	}
	return false
}

func (x *Scoreboard_Row) GetDisqualified() bool {
	if x != nil {
		return x.Disqualified
	}
	return false
}

func (x *Scoreboard_Row) GetMedal() Medal {
	if x != nil {
		return x.Medal
	}
	return Medal_NO_MEDAL
}

func (x *Scoreboard_Row) GetValues() []*Scoreboard_Row_Value {
	if x != nil {
		return x.Values
	}
	return nil
}

type Scoreboard_Row_Value struct {
	state    protoimpl.MessageState `protogen:"open.v1"`
	ColumnId string                 `protobuf:"bytes,1,opt,name=column_id,json=columnId,proto3" json:"column_id,omitempty"`
	// Types that are valid to be assigned to Value:
	//
	//	*Scoreboard_Row_Value_RoundScore
	//	*Scoreboard_Row_Value_ProblemScore
	//	*Scoreboard_Row_Value_String_
	//	*Scoreboard_Row_Value_Number
	Value         isScoreboard_Row_Value_Value `protobuf_oneof:"value"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Scoreboard_Row_Value) Reset() {
	*x = Scoreboard_Row_Value{}
	mi := &file_eolymp_judge_scoreboard_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Scoreboard_Row_Value) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Scoreboard_Row_Value) ProtoMessage() {}

func (x *Scoreboard_Row_Value) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_judge_scoreboard_proto_msgTypes[4]
	if x != nil {
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
	return file_eolymp_judge_scoreboard_proto_rawDescGZIP(), []int{0, 2, 0}
}

func (x *Scoreboard_Row_Value) GetColumnId() string {
	if x != nil {
		return x.ColumnId
	}
	return ""
}

func (x *Scoreboard_Row_Value) GetValue() isScoreboard_Row_Value_Value {
	if x != nil {
		return x.Value
	}
	return nil
}

func (x *Scoreboard_Row_Value) GetRoundScore() *Scoreboard_Row_RoundScore {
	if x != nil {
		if x, ok := x.Value.(*Scoreboard_Row_Value_RoundScore); ok {
			return x.RoundScore
		}
	}
	return nil
}

func (x *Scoreboard_Row_Value) GetProblemScore() *Scoreboard_Row_ProblemScore {
	if x != nil {
		if x, ok := x.Value.(*Scoreboard_Row_Value_ProblemScore); ok {
			return x.ProblemScore
		}
	}
	return nil
}

func (x *Scoreboard_Row_Value) GetString_() string {
	if x != nil {
		if x, ok := x.Value.(*Scoreboard_Row_Value_String_); ok {
			return x.String_
		}
	}
	return ""
}

func (x *Scoreboard_Row_Value) GetNumber() string {
	if x != nil {
		if x, ok := x.Value.(*Scoreboard_Row_Value_Number); ok {
			return x.Number
		}
	}
	return ""
}

type isScoreboard_Row_Value_Value interface {
	isScoreboard_Row_Value_Value()
}

type Scoreboard_Row_Value_RoundScore struct {
	RoundScore *Scoreboard_Row_RoundScore `protobuf:"bytes,11,opt,name=round_score,json=roundScore,proto3,oneof"`
}

type Scoreboard_Row_Value_ProblemScore struct {
	ProblemScore *Scoreboard_Row_ProblemScore `protobuf:"bytes,12,opt,name=problem_score,json=problemScore,proto3,oneof"`
}

type Scoreboard_Row_Value_String_ struct {
	String_ string `protobuf:"bytes,13,opt,name=string,proto3,oneof"`
}

type Scoreboard_Row_Value_Number struct {
	Number string `protobuf:"bytes,14,opt,name=number,proto3,oneof"`
}

func (*Scoreboard_Row_Value_RoundScore) isScoreboard_Row_Value_Value() {}

func (*Scoreboard_Row_Value_ProblemScore) isScoreboard_Row_Value_Value() {}

func (*Scoreboard_Row_Value_String_) isScoreboard_Row_Value_Value() {}

func (*Scoreboard_Row_Value_Number) isScoreboard_Row_Value_Value() {}

type Scoreboard_Row_RoundScore struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Score         float32                `protobuf:"fixed32,1,opt,name=score,proto3" json:"score,omitempty"`
	Penalty       float32                `protobuf:"fixed32,2,opt,name=penalty,proto3" json:"penalty,omitempty"`
	TieBreaker    uint32                 `protobuf:"varint,3,opt,name=tie_breaker,json=tieBreaker,proto3" json:"tie_breaker,omitempty"`
	Unofficial    bool                   `protobuf:"varint,10,opt,name=unofficial,proto3" json:"unofficial,omitempty"`
	Disqualified  bool                   `protobuf:"varint,11,opt,name=disqualified,proto3" json:"disqualified,omitempty"`
	Medal         Medal                  `protobuf:"varint,32,opt,name=medal,proto3,enum=eolymp.judge.Medal" json:"medal,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Scoreboard_Row_RoundScore) Reset() {
	*x = Scoreboard_Row_RoundScore{}
	mi := &file_eolymp_judge_scoreboard_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Scoreboard_Row_RoundScore) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Scoreboard_Row_RoundScore) ProtoMessage() {}

func (x *Scoreboard_Row_RoundScore) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_judge_scoreboard_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Scoreboard_Row_RoundScore.ProtoReflect.Descriptor instead.
func (*Scoreboard_Row_RoundScore) Descriptor() ([]byte, []int) {
	return file_eolymp_judge_scoreboard_proto_rawDescGZIP(), []int{0, 2, 1}
}

func (x *Scoreboard_Row_RoundScore) GetScore() float32 {
	if x != nil {
		return x.Score
	}
	return 0
}

func (x *Scoreboard_Row_RoundScore) GetPenalty() float32 {
	if x != nil {
		return x.Penalty
	}
	return 0
}

func (x *Scoreboard_Row_RoundScore) GetTieBreaker() uint32 {
	if x != nil {
		return x.TieBreaker
	}
	return 0
}

func (x *Scoreboard_Row_RoundScore) GetUnofficial() bool {
	if x != nil {
		return x.Unofficial
	}
	return false
}

func (x *Scoreboard_Row_RoundScore) GetDisqualified() bool {
	if x != nil {
		return x.Disqualified
	}
	return false
}

func (x *Scoreboard_Row_RoundScore) GetMedal() Medal {
	if x != nil {
		return x.Medal
	}
	return Medal_NO_MEDAL
}

type Scoreboard_Row_ProblemScore struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Score         float32                `protobuf:"fixed32,1,opt,name=score,proto3" json:"score,omitempty"`
	Penalty       float32                `protobuf:"fixed32,2,opt,name=penalty,proto3" json:"penalty,omitempty"`
	Attempts      uint32                 `protobuf:"varint,3,opt,name=attempts,proto3" json:"attempts,omitempty"`
	Percentage    float32                `protobuf:"fixed32,4,opt,name=percentage,proto3" json:"percentage,omitempty"`
	Time          uint32                 `protobuf:"varint,5,opt,name=time,proto3" json:"time,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Scoreboard_Row_ProblemScore) Reset() {
	*x = Scoreboard_Row_ProblemScore{}
	mi := &file_eolymp_judge_scoreboard_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Scoreboard_Row_ProblemScore) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Scoreboard_Row_ProblemScore) ProtoMessage() {}

func (x *Scoreboard_Row_ProblemScore) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_judge_scoreboard_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Scoreboard_Row_ProblemScore.ProtoReflect.Descriptor instead.
func (*Scoreboard_Row_ProblemScore) Descriptor() ([]byte, []int) {
	return file_eolymp_judge_scoreboard_proto_rawDescGZIP(), []int{0, 2, 2}
}

func (x *Scoreboard_Row_ProblemScore) GetScore() float32 {
	if x != nil {
		return x.Score
	}
	return 0
}

func (x *Scoreboard_Row_ProblemScore) GetPenalty() float32 {
	if x != nil {
		return x.Penalty
	}
	return 0
}

func (x *Scoreboard_Row_ProblemScore) GetAttempts() uint32 {
	if x != nil {
		return x.Attempts
	}
	return 0
}

func (x *Scoreboard_Row_ProblemScore) GetPercentage() float32 {
	if x != nil {
		return x.Percentage
	}
	return 0
}

func (x *Scoreboard_Row_ProblemScore) GetTime() uint32 {
	if x != nil {
		return x.Time
	}
	return 0
}

var File_eolymp_judge_scoreboard_proto protoreflect.FileDescriptor

var file_eolymp_judge_scoreboard_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x6a, 0x75, 0x64, 0x67, 0x65, 0x2f, 0x73,
	0x63, 0x6f, 0x72, 0x65, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0c, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x6a, 0x75, 0x64, 0x67, 0x65, 0x1a, 0x18, 0x65,
	0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x6a, 0x75, 0x64, 0x67, 0x65, 0x2f, 0x6d, 0x65, 0x64, 0x61,
	0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xdf, 0x0c, 0x0a, 0x0a, 0x53, 0x63, 0x6f, 0x72,
	0x65, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x12, 0x33, 0x0a, 0x05, 0x6d, 0x6f, 0x64, 0x65, 0x73, 0x18,
	0x0a, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x1d, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x6a,
	0x75, 0x64, 0x67, 0x65, 0x2e, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e,
	0x4d, 0x6f, 0x64, 0x65, 0x52, 0x05, 0x6d, 0x6f, 0x64, 0x65, 0x73, 0x12, 0x36, 0x0a, 0x06, 0x72,
	0x6f, 0x75, 0x6e, 0x64, 0x73, 0x18, 0x0c, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x65, 0x6f,
	0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x6a, 0x75, 0x64, 0x67, 0x65, 0x2e, 0x53, 0x63, 0x6f, 0x72, 0x65,
	0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x52, 0x6f, 0x75, 0x6e, 0x64, 0x52, 0x06, 0x72, 0x6f, 0x75,
	0x6e, 0x64, 0x73, 0x12, 0x39, 0x0a, 0x07, 0x63, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x73, 0x18, 0x0b,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x6a, 0x75,
	0x64, 0x67, 0x65, 0x2e, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x43,
	0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x52, 0x07, 0x63, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x73, 0x1a, 0x2d,
	0x0a, 0x05, 0x52, 0x6f, 0x75, 0x6e, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x1a, 0xee, 0x02,
	0x0a, 0x06, 0x43, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x38, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x24, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e,
	0x6a, 0x75, 0x64, 0x67, 0x65, 0x2e, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x62, 0x6f, 0x61, 0x72, 0x64,
	0x2e, 0x43, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x68, 0x6f, 0x69,
	0x63, 0x65, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x63, 0x68, 0x6f, 0x69, 0x63,
	0x65, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x6f, 0x72, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x14,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x73, 0x6f, 0x72, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x1e,
	0x0a, 0x0a, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x15, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x0a, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x61, 0x62, 0x6c, 0x65, 0x22, 0xad,
	0x01, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x10, 0x0a, 0x0c, 0x55, 0x4e, 0x4b, 0x4e, 0x4f,
	0x57, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x10, 0x00, 0x12, 0x0f, 0x0a, 0x0b, 0x52, 0x4f, 0x55,
	0x4e, 0x44, 0x5f, 0x53, 0x43, 0x4f, 0x52, 0x45, 0x10, 0x02, 0x12, 0x11, 0x0a, 0x0d, 0x50, 0x52,
	0x4f, 0x42, 0x4c, 0x45, 0x4d, 0x5f, 0x53, 0x43, 0x4f, 0x52, 0x45, 0x10, 0x03, 0x12, 0x0a, 0x0a,
	0x06, 0x53, 0x54, 0x52, 0x49, 0x4e, 0x47, 0x10, 0x0a, 0x12, 0x0a, 0x0a, 0x06, 0x4e, 0x55, 0x4d,
	0x42, 0x45, 0x52, 0x10, 0x0b, 0x12, 0x0a, 0x0a, 0x06, 0x43, 0x48, 0x4f, 0x49, 0x43, 0x45, 0x10,
	0x0c, 0x12, 0x08, 0x0a, 0x04, 0x44, 0x41, 0x54, 0x45, 0x10, 0x0d, 0x12, 0x09, 0x0a, 0x05, 0x45,
	0x4d, 0x41, 0x49, 0x4c, 0x10, 0x0e, 0x12, 0x0c, 0x0a, 0x08, 0x43, 0x48, 0x45, 0x43, 0x4b, 0x42,
	0x4f, 0x58, 0x10, 0x0f, 0x12, 0x0b, 0x0a, 0x07, 0x43, 0x4f, 0x55, 0x4e, 0x54, 0x52, 0x59, 0x10,
	0x10, 0x12, 0x0a, 0x0a, 0x06, 0x52, 0x45, 0x47, 0x49, 0x4f, 0x4e, 0x10, 0x11, 0x12, 0x0f, 0x0a,
	0x0b, 0x49, 0x4e, 0x53, 0x54, 0x49, 0x54, 0x55, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x12, 0x1a, 0xdb,
	0x07, 0x0a, 0x03, 0x52, 0x6f, 0x77, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6d, 0x65, 0x6d, 0x62, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x61, 0x6e,
	0x6b, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x72, 0x61, 0x6e, 0x6b, 0x12, 0x1f, 0x0a,
	0x0b, 0x72, 0x61, 0x6e, 0x6b, 0x5f, 0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x18, 0x0c, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x0a, 0x72, 0x61, 0x6e, 0x6b, 0x4c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x12, 0x14,
	0x0a, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x14, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x73,
	0x63, 0x6f, 0x72, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x65, 0x6e, 0x61, 0x6c, 0x74, 0x79, 0x18,
	0x15, 0x20, 0x01, 0x28, 0x02, 0x52, 0x07, 0x70, 0x65, 0x6e, 0x61, 0x6c, 0x74, 0x79, 0x12, 0x1f,
	0x0a, 0x0b, 0x74, 0x69, 0x65, 0x5f, 0x62, 0x72, 0x65, 0x61, 0x6b, 0x65, 0x72, 0x18, 0x16, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x0a, 0x74, 0x69, 0x65, 0x42, 0x72, 0x65, 0x61, 0x6b, 0x65, 0x72, 0x12,
	0x1e, 0x0a, 0x0a, 0x75, 0x6e, 0x6f, 0x66, 0x66, 0x69, 0x63, 0x69, 0x61, 0x6c, 0x18, 0x1e, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x0a, 0x75, 0x6e, 0x6f, 0x66, 0x66, 0x69, 0x63, 0x69, 0x61, 0x6c, 0x12,
	0x22, 0x0a, 0x0c, 0x64, 0x69, 0x73, 0x71, 0x75, 0x61, 0x6c, 0x69, 0x66, 0x69, 0x65, 0x64, 0x18,
	0x1f, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x64, 0x69, 0x73, 0x71, 0x75, 0x61, 0x6c, 0x69, 0x66,
	0x69, 0x65, 0x64, 0x12, 0x29, 0x0a, 0x05, 0x6d, 0x65, 0x64, 0x61, 0x6c, 0x18, 0x20, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x13, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x6a, 0x75, 0x64, 0x67,
	0x65, 0x2e, 0x4d, 0x65, 0x64, 0x61, 0x6c, 0x52, 0x05, 0x6d, 0x65, 0x64, 0x61, 0x6c, 0x12, 0x3a,
	0x0a, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x32, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22,
	0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x6a, 0x75, 0x64, 0x67, 0x65, 0x2e, 0x53, 0x63,
	0x6f, 0x72, 0x65, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x52, 0x6f, 0x77, 0x2e, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x52, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x1a, 0xff, 0x01, 0x0a, 0x05, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x49,
	0x64, 0x12, 0x4a, 0x0a, 0x0b, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x5f, 0x73, 0x63, 0x6f, 0x72, 0x65,
	0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e,
	0x6a, 0x75, 0x64, 0x67, 0x65, 0x2e, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x62, 0x6f, 0x61, 0x72, 0x64,
	0x2e, 0x52, 0x6f, 0x77, 0x2e, 0x52, 0x6f, 0x75, 0x6e, 0x64, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x48,
	0x00, 0x52, 0x0a, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x12, 0x50, 0x0a,
	0x0d, 0x70, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x5f, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x0c,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x6a, 0x75,
	0x64, 0x67, 0x65, 0x2e, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x52,
	0x6f, 0x77, 0x2e, 0x50, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x48,
	0x00, 0x52, 0x0c, 0x70, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x12,
	0x18, 0x0a, 0x06, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x48,
	0x00, 0x52, 0x06, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x12, 0x18, 0x0a, 0x06, 0x6e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x06, 0x6e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x42, 0x07, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x1a, 0xcc, 0x01, 0x0a,
	0x0a, 0x52, 0x6f, 0x75, 0x6e, 0x64, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x73,
	0x63, 0x6f, 0x72, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x73, 0x63, 0x6f, 0x72,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x65, 0x6e, 0x61, 0x6c, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x02, 0x52, 0x07, 0x70, 0x65, 0x6e, 0x61, 0x6c, 0x74, 0x79, 0x12, 0x1f, 0x0a, 0x0b, 0x74,
	0x69, 0x65, 0x5f, 0x62, 0x72, 0x65, 0x61, 0x6b, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x0a, 0x74, 0x69, 0x65, 0x42, 0x72, 0x65, 0x61, 0x6b, 0x65, 0x72, 0x12, 0x1e, 0x0a, 0x0a,
	0x75, 0x6e, 0x6f, 0x66, 0x66, 0x69, 0x63, 0x69, 0x61, 0x6c, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x0a, 0x75, 0x6e, 0x6f, 0x66, 0x66, 0x69, 0x63, 0x69, 0x61, 0x6c, 0x12, 0x22, 0x0a, 0x0c,
	0x64, 0x69, 0x73, 0x71, 0x75, 0x61, 0x6c, 0x69, 0x66, 0x69, 0x65, 0x64, 0x18, 0x0b, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x0c, 0x64, 0x69, 0x73, 0x71, 0x75, 0x61, 0x6c, 0x69, 0x66, 0x69, 0x65, 0x64,
	0x12, 0x29, 0x0a, 0x05, 0x6d, 0x65, 0x64, 0x61, 0x6c, 0x18, 0x20, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x13, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x6a, 0x75, 0x64, 0x67, 0x65, 0x2e, 0x4d,
	0x65, 0x64, 0x61, 0x6c, 0x52, 0x05, 0x6d, 0x65, 0x64, 0x61, 0x6c, 0x1a, 0x8e, 0x01, 0x0a, 0x0c,
	0x50, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x73, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x73, 0x63, 0x6f,
	0x72, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x65, 0x6e, 0x61, 0x6c, 0x74, 0x79, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x02, 0x52, 0x07, 0x70, 0x65, 0x6e, 0x61, 0x6c, 0x74, 0x79, 0x12, 0x1a, 0x0a, 0x08,
	0x61, 0x74, 0x74, 0x65, 0x6d, 0x70, 0x74, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08,
	0x61, 0x74, 0x74, 0x65, 0x6d, 0x70, 0x74, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x65, 0x72, 0x63,
	0x65, 0x6e, 0x74, 0x61, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0a, 0x70, 0x65,
	0x72, 0x63, 0x65, 0x6e, 0x74, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x22, 0x2b, 0x0a, 0x04,
	0x4d, 0x6f, 0x64, 0x65, 0x12, 0x0a, 0x0a, 0x06, 0x52, 0x45, 0x53, 0x55, 0x4c, 0x54, 0x10, 0x00,
	0x12, 0x0a, 0x0a, 0x06, 0x46, 0x52, 0x4f, 0x5a, 0x45, 0x4e, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07,
	0x55, 0x50, 0x53, 0x4f, 0x4c, 0x56, 0x45, 0x10, 0x02, 0x42, 0x2d, 0x5a, 0x2b, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x67,
	0x6f, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x6a, 0x75, 0x64,
	0x67, 0x65, 0x3b, 0x6a, 0x75, 0x64, 0x67, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_eolymp_judge_scoreboard_proto_rawDescOnce sync.Once
	file_eolymp_judge_scoreboard_proto_rawDescData = file_eolymp_judge_scoreboard_proto_rawDesc
)

func file_eolymp_judge_scoreboard_proto_rawDescGZIP() []byte {
	file_eolymp_judge_scoreboard_proto_rawDescOnce.Do(func() {
		file_eolymp_judge_scoreboard_proto_rawDescData = protoimpl.X.CompressGZIP(file_eolymp_judge_scoreboard_proto_rawDescData)
	})
	return file_eolymp_judge_scoreboard_proto_rawDescData
}

var file_eolymp_judge_scoreboard_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_eolymp_judge_scoreboard_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_eolymp_judge_scoreboard_proto_goTypes = []any{
	(Scoreboard_Mode)(0),                // 0: eolymp.judge.Scoreboard.Mode
	(Scoreboard_Column_Type)(0),         // 1: eolymp.judge.Scoreboard.Column.Type
	(*Scoreboard)(nil),                  // 2: eolymp.judge.Scoreboard
	(*Scoreboard_Round)(nil),            // 3: eolymp.judge.Scoreboard.Round
	(*Scoreboard_Column)(nil),           // 4: eolymp.judge.Scoreboard.Column
	(*Scoreboard_Row)(nil),              // 5: eolymp.judge.Scoreboard.Row
	(*Scoreboard_Row_Value)(nil),        // 6: eolymp.judge.Scoreboard.Row.Value
	(*Scoreboard_Row_RoundScore)(nil),   // 7: eolymp.judge.Scoreboard.Row.RoundScore
	(*Scoreboard_Row_ProblemScore)(nil), // 8: eolymp.judge.Scoreboard.Row.ProblemScore
	(Medal)(0),                          // 9: eolymp.judge.Medal
}
var file_eolymp_judge_scoreboard_proto_depIdxs = []int32{
	0, // 0: eolymp.judge.Scoreboard.modes:type_name -> eolymp.judge.Scoreboard.Mode
	3, // 1: eolymp.judge.Scoreboard.rounds:type_name -> eolymp.judge.Scoreboard.Round
	4, // 2: eolymp.judge.Scoreboard.columns:type_name -> eolymp.judge.Scoreboard.Column
	1, // 3: eolymp.judge.Scoreboard.Column.type:type_name -> eolymp.judge.Scoreboard.Column.Type
	9, // 4: eolymp.judge.Scoreboard.Row.medal:type_name -> eolymp.judge.Medal
	6, // 5: eolymp.judge.Scoreboard.Row.values:type_name -> eolymp.judge.Scoreboard.Row.Value
	7, // 6: eolymp.judge.Scoreboard.Row.Value.round_score:type_name -> eolymp.judge.Scoreboard.Row.RoundScore
	8, // 7: eolymp.judge.Scoreboard.Row.Value.problem_score:type_name -> eolymp.judge.Scoreboard.Row.ProblemScore
	9, // 8: eolymp.judge.Scoreboard.Row.RoundScore.medal:type_name -> eolymp.judge.Medal
	9, // [9:9] is the sub-list for method output_type
	9, // [9:9] is the sub-list for method input_type
	9, // [9:9] is the sub-list for extension type_name
	9, // [9:9] is the sub-list for extension extendee
	0, // [0:9] is the sub-list for field type_name
}

func init() { file_eolymp_judge_scoreboard_proto_init() }
func file_eolymp_judge_scoreboard_proto_init() {
	if File_eolymp_judge_scoreboard_proto != nil {
		return
	}
	file_eolymp_judge_medal_proto_init()
	file_eolymp_judge_scoreboard_proto_msgTypes[4].OneofWrappers = []any{
		(*Scoreboard_Row_Value_RoundScore)(nil),
		(*Scoreboard_Row_Value_ProblemScore)(nil),
		(*Scoreboard_Row_Value_String_)(nil),
		(*Scoreboard_Row_Value_Number)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_eolymp_judge_scoreboard_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_eolymp_judge_scoreboard_proto_goTypes,
		DependencyIndexes: file_eolymp_judge_scoreboard_proto_depIdxs,
		EnumInfos:         file_eolymp_judge_scoreboard_proto_enumTypes,
		MessageInfos:      file_eolymp_judge_scoreboard_proto_msgTypes,
	}.Build()
	File_eolymp_judge_scoreboard_proto = out.File
	file_eolymp_judge_scoreboard_proto_rawDesc = nil
	file_eolymp_judge_scoreboard_proto_goTypes = nil
	file_eolymp_judge_scoreboard_proto_depIdxs = nil
}
