// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v3.21.12
// source: eolymp/judge/scoreboard_service.proto

package judge

import (
	_ "github.com/eolymp/go-sdk/eolymp/annotations"
	wellknown "github.com/eolymp/go-sdk/eolymp/wellknown"
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

type DescribeScoreboardInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RoundId string `protobuf:"bytes,1,opt,name=round_id,json=roundId,proto3" json:"round_id,omitempty"`
}

func (x *DescribeScoreboardInput) Reset() {
	*x = DescribeScoreboardInput{}
	mi := &file_eolymp_judge_scoreboard_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DescribeScoreboardInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescribeScoreboardInput) ProtoMessage() {}

func (x *DescribeScoreboardInput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_judge_scoreboard_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DescribeScoreboardInput.ProtoReflect.Descriptor instead.
func (*DescribeScoreboardInput) Descriptor() ([]byte, []int) {
	return file_eolymp_judge_scoreboard_service_proto_rawDescGZIP(), []int{0}
}

func (x *DescribeScoreboardInput) GetRoundId() string {
	if x != nil {
		return x.RoundId
	}
	return ""
}

type DescribeScoreboardOutput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Scoreboard *Scoreboard `protobuf:"bytes,1,opt,name=scoreboard,proto3" json:"scoreboard,omitempty"`
}

func (x *DescribeScoreboardOutput) Reset() {
	*x = DescribeScoreboardOutput{}
	mi := &file_eolymp_judge_scoreboard_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DescribeScoreboardOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescribeScoreboardOutput) ProtoMessage() {}

func (x *DescribeScoreboardOutput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_judge_scoreboard_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DescribeScoreboardOutput.ProtoReflect.Descriptor instead.
func (*DescribeScoreboardOutput) Descriptor() ([]byte, []int) {
	return file_eolymp_judge_scoreboard_service_proto_rawDescGZIP(), []int{1}
}

func (x *DescribeScoreboardOutput) GetScoreboard() *Scoreboard {
	if x != nil {
		return x.Scoreboard
	}
	return nil
}

type ListScoreboardRowsInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Mode    Scoreboard_Mode                 `protobuf:"varint,1,opt,name=mode,proto3,enum=eolymp.judge.Scoreboard_Mode" json:"mode,omitempty"` // Mode for fetching score value (see enum description).
	RoundId string                          `protobuf:"bytes,2,opt,name=round_id,json=roundId,proto3" json:"round_id,omitempty"`
	Size    int32                           `protobuf:"varint,10,opt,name=size,proto3" json:"size,omitempty"` // Number of results per page, (max. 100, default 100)
	Offset  int32                           `protobuf:"varint,12,opt,name=offset,proto3" json:"offset,omitempty"`
	Filters *ListScoreboardRowsInput_Filter `protobuf:"bytes,40,opt,name=filters,proto3" json:"filters,omitempty"`
	Sort    string                          `protobuf:"bytes,50,opt,name=sort,proto3" json:"sort,omitempty"`
	Order   wellknown.Direction             `protobuf:"varint,60,opt,name=order,proto3,enum=eolymp.wellknown.Direction" json:"order,omitempty"`
}

func (x *ListScoreboardRowsInput) Reset() {
	*x = ListScoreboardRowsInput{}
	mi := &file_eolymp_judge_scoreboard_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListScoreboardRowsInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListScoreboardRowsInput) ProtoMessage() {}

func (x *ListScoreboardRowsInput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_judge_scoreboard_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListScoreboardRowsInput.ProtoReflect.Descriptor instead.
func (*ListScoreboardRowsInput) Descriptor() ([]byte, []int) {
	return file_eolymp_judge_scoreboard_service_proto_rawDescGZIP(), []int{2}
}

func (x *ListScoreboardRowsInput) GetMode() Scoreboard_Mode {
	if x != nil {
		return x.Mode
	}
	return Scoreboard_RESULT
}

func (x *ListScoreboardRowsInput) GetRoundId() string {
	if x != nil {
		return x.RoundId
	}
	return ""
}

func (x *ListScoreboardRowsInput) GetSize() int32 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *ListScoreboardRowsInput) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *ListScoreboardRowsInput) GetFilters() *ListScoreboardRowsInput_Filter {
	if x != nil {
		return x.Filters
	}
	return nil
}

func (x *ListScoreboardRowsInput) GetSort() string {
	if x != nil {
		return x.Sort
	}
	return ""
}

func (x *ListScoreboardRowsInput) GetOrder() wellknown.Direction {
	if x != nil {
		return x.Order
	}
	return wellknown.Direction(0)
}

type ListScoreboardRowsOutput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Total int32             `protobuf:"varint,1,opt,name=total,proto3" json:"total,omitempty"`
	Items []*Scoreboard_Row `protobuf:"bytes,2,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *ListScoreboardRowsOutput) Reset() {
	*x = ListScoreboardRowsOutput{}
	mi := &file_eolymp_judge_scoreboard_service_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListScoreboardRowsOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListScoreboardRowsOutput) ProtoMessage() {}

func (x *ListScoreboardRowsOutput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_judge_scoreboard_service_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListScoreboardRowsOutput.ProtoReflect.Descriptor instead.
func (*ListScoreboardRowsOutput) Descriptor() ([]byte, []int) {
	return file_eolymp_judge_scoreboard_service_proto_rawDescGZIP(), []int{3}
}

func (x *ListScoreboardRowsOutput) GetTotal() int32 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *ListScoreboardRowsOutput) GetItems() []*Scoreboard_Row {
	if x != nil {
		return x.Items
	}
	return nil
}

type ListScoreboardRowsInput_Filter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Unofficial   []*wellknown.ExpressionBool `protobuf:"bytes,10,rep,name=unofficial,proto3" json:"unofficial,omitempty"`
	Disqualified []*wellknown.ExpressionBool `protobuf:"bytes,11,rep,name=disqualified,proto3" json:"disqualified,omitempty"`
}

func (x *ListScoreboardRowsInput_Filter) Reset() {
	*x = ListScoreboardRowsInput_Filter{}
	mi := &file_eolymp_judge_scoreboard_service_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListScoreboardRowsInput_Filter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListScoreboardRowsInput_Filter) ProtoMessage() {}

func (x *ListScoreboardRowsInput_Filter) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_judge_scoreboard_service_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListScoreboardRowsInput_Filter.ProtoReflect.Descriptor instead.
func (*ListScoreboardRowsInput_Filter) Descriptor() ([]byte, []int) {
	return file_eolymp_judge_scoreboard_service_proto_rawDescGZIP(), []int{2, 0}
}

func (x *ListScoreboardRowsInput_Filter) GetUnofficial() []*wellknown.ExpressionBool {
	if x != nil {
		return x.Unofficial
	}
	return nil
}

func (x *ListScoreboardRowsInput_Filter) GetDisqualified() []*wellknown.ExpressionBool {
	if x != nil {
		return x.Disqualified
	}
	return nil
}

var File_eolymp_judge_scoreboard_service_proto protoreflect.FileDescriptor

var file_eolymp_judge_scoreboard_service_proto_rawDesc = []byte{
	0x0a, 0x25, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x6a, 0x75, 0x64, 0x67, 0x65, 0x2f, 0x73,
	0x63, 0x6f, 0x72, 0x65, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e,
	0x6a, 0x75, 0x64, 0x67, 0x65, 0x1a, 0x1d, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x61, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x22, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x72, 0x61, 0x74, 0x65, 0x6c, 0x69, 0x6d,
	0x69, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70,
	0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x63, 0x6f,
	0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70,
	0x2f, 0x6a, 0x75, 0x64, 0x67, 0x65, 0x2f, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x62, 0x6f, 0x61, 0x72,
	0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f,
	0x77, 0x65, 0x6c, 0x6c, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x2f, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x21, 0x65, 0x6f, 0x6c, 0x79, 0x6d,
	0x70, 0x2f, 0x77, 0x65, 0x6c, 0x6c, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x2f, 0x65, 0x78, 0x70, 0x72,
	0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x34, 0x0a, 0x17,
	0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x62, 0x6f, 0x61,
	0x72, 0x64, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x72, 0x6f, 0x75, 0x6e, 0x64,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x72, 0x6f, 0x75, 0x6e, 0x64,
	0x49, 0x64, 0x22, 0x54, 0x0a, 0x18, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x53, 0x63,
	0x6f, 0x72, 0x65, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x12, 0x38,
	0x0a, 0x0a, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x18, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x6a, 0x75, 0x64, 0x67,
	0x65, 0x2e, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x52, 0x0a, 0x73, 0x63,
	0x6f, 0x72, 0x65, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x22, 0xb5, 0x03, 0x0a, 0x17, 0x4c, 0x69, 0x73,
	0x74, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x52, 0x6f, 0x77, 0x73, 0x49,
	0x6e, 0x70, 0x75, 0x74, 0x12, 0x31, 0x0a, 0x04, 0x6d, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x1d, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x6a, 0x75, 0x64, 0x67,
	0x65, 0x2e, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x4d, 0x6f, 0x64,
	0x65, 0x52, 0x04, 0x6d, 0x6f, 0x64, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x72, 0x6f, 0x75, 0x6e, 0x64,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x72, 0x6f, 0x75, 0x6e, 0x64,
	0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74,
	0x18, 0x0c, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x46,
	0x0a, 0x07, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x18, 0x28, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x2c, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x6a, 0x75, 0x64, 0x67, 0x65, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x52, 0x6f, 0x77,
	0x73, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x2e, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x52, 0x07, 0x66,
	0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x18, 0x32,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x12, 0x31, 0x0a, 0x05, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x18, 0x3c, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1b, 0x2e, 0x65, 0x6f, 0x6c, 0x79,
	0x6d, 0x70, 0x2e, 0x77, 0x65, 0x6c, 0x6c, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x2e, 0x44, 0x69, 0x72,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x05, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x1a, 0x90, 0x01,
	0x0a, 0x06, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x40, 0x0a, 0x0a, 0x75, 0x6e, 0x6f, 0x66,
	0x66, 0x69, 0x63, 0x69, 0x61, 0x6c, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x65,
	0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x77, 0x65, 0x6c, 0x6c, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x2e,
	0x45, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x42, 0x6f, 0x6f, 0x6c, 0x52, 0x0a,
	0x75, 0x6e, 0x6f, 0x66, 0x66, 0x69, 0x63, 0x69, 0x61, 0x6c, 0x12, 0x44, 0x0a, 0x0c, 0x64, 0x69,
	0x73, 0x71, 0x75, 0x61, 0x6c, 0x69, 0x66, 0x69, 0x65, 0x64, 0x18, 0x0b, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x20, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x77, 0x65, 0x6c, 0x6c, 0x6b, 0x6e,
	0x6f, 0x77, 0x6e, 0x2e, 0x45, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x42, 0x6f,
	0x6f, 0x6c, 0x52, 0x0c, 0x64, 0x69, 0x73, 0x71, 0x75, 0x61, 0x6c, 0x69, 0x66, 0x69, 0x65, 0x64,
	0x22, 0x64, 0x0a, 0x18, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x62, 0x6f, 0x61,
	0x72, 0x64, 0x52, 0x6f, 0x77, 0x73, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x12, 0x14, 0x0a, 0x05,
	0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x74, 0x6f, 0x74,
	0x61, 0x6c, 0x12, 0x32, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x1c, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x6a, 0x75, 0x64, 0x67, 0x65,
	0x2e, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x52, 0x6f, 0x77, 0x52,
	0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x32, 0xe2, 0x02, 0x0a, 0x11, 0x53, 0x63, 0x6f, 0x72, 0x65,
	0x62, 0x6f, 0x61, 0x72, 0x64, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0xa2, 0x01, 0x0a,
	0x12, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x62, 0x6f,
	0x61, 0x72, 0x64, 0x12, 0x25, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x6a, 0x75, 0x64,
	0x67, 0x65, 0x2e, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x53, 0x63, 0x6f, 0x72, 0x65,
	0x62, 0x6f, 0x61, 0x72, 0x64, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x26, 0x2e, 0x65, 0x6f, 0x6c,
	0x79, 0x6d, 0x70, 0x2e, 0x6a, 0x75, 0x64, 0x67, 0x65, 0x2e, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x62, 0x65, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x4f, 0x75, 0x74, 0x70,
	0x75, 0x74, 0x22, 0x3d, 0xea, 0xe2, 0x0a, 0x0c, 0xf5, 0xe2, 0x0a, 0x00, 0x00, 0x48, 0x42, 0xf8,
	0xe2, 0x0a, 0xc8, 0x01, 0x82, 0xe3, 0x0a, 0x16, 0x8a, 0xe3, 0x0a, 0x12, 0x6a, 0x75, 0x64, 0x67,
	0x65, 0x3a, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x73, 0x74, 0x3a, 0x72, 0x65, 0x61, 0x64, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x0d, 0x12, 0x0b, 0x2f, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x62, 0x6f, 0x61, 0x72,
	0x64, 0x12, 0xa7, 0x01, 0x0a, 0x12, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x62,
	0x6f, 0x61, 0x72, 0x64, 0x52, 0x6f, 0x77, 0x73, 0x12, 0x25, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d,
	0x70, 0x2e, 0x6a, 0x75, 0x64, 0x67, 0x65, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x63, 0x6f, 0x72,
	0x65, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x52, 0x6f, 0x77, 0x73, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a,
	0x26, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x6a, 0x75, 0x64, 0x67, 0x65, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x52, 0x6f, 0x77,
	0x73, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x22, 0x42, 0xea, 0xe2, 0x0a, 0x0c, 0xf5, 0xe2, 0x0a,
	0x00, 0x00, 0x48, 0x42, 0xf8, 0xe2, 0x0a, 0xc8, 0x01, 0x82, 0xe3, 0x0a, 0x16, 0x8a, 0xe3, 0x0a,
	0x12, 0x6a, 0x75, 0x64, 0x67, 0x65, 0x3a, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x73, 0x74, 0x3a, 0x72,
	0x65, 0x61, 0x64, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x12, 0x12, 0x10, 0x2f, 0x73, 0x63, 0x6f, 0x72,
	0x65, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2f, 0x72, 0x6f, 0x77, 0x73, 0x42, 0x2d, 0x5a, 0x2b, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70,
	0x2f, 0x67, 0x6f, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x6a,
	0x75, 0x64, 0x67, 0x65, 0x3b, 0x6a, 0x75, 0x64, 0x67, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_eolymp_judge_scoreboard_service_proto_rawDescOnce sync.Once
	file_eolymp_judge_scoreboard_service_proto_rawDescData = file_eolymp_judge_scoreboard_service_proto_rawDesc
)

func file_eolymp_judge_scoreboard_service_proto_rawDescGZIP() []byte {
	file_eolymp_judge_scoreboard_service_proto_rawDescOnce.Do(func() {
		file_eolymp_judge_scoreboard_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_eolymp_judge_scoreboard_service_proto_rawDescData)
	})
	return file_eolymp_judge_scoreboard_service_proto_rawDescData
}

var file_eolymp_judge_scoreboard_service_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_eolymp_judge_scoreboard_service_proto_goTypes = []any{
	(*DescribeScoreboardInput)(nil),        // 0: eolymp.judge.DescribeScoreboardInput
	(*DescribeScoreboardOutput)(nil),       // 1: eolymp.judge.DescribeScoreboardOutput
	(*ListScoreboardRowsInput)(nil),        // 2: eolymp.judge.ListScoreboardRowsInput
	(*ListScoreboardRowsOutput)(nil),       // 3: eolymp.judge.ListScoreboardRowsOutput
	(*ListScoreboardRowsInput_Filter)(nil), // 4: eolymp.judge.ListScoreboardRowsInput.Filter
	(*Scoreboard)(nil),                     // 5: eolymp.judge.Scoreboard
	(Scoreboard_Mode)(0),                   // 6: eolymp.judge.Scoreboard.Mode
	(wellknown.Direction)(0),               // 7: eolymp.wellknown.Direction
	(*Scoreboard_Row)(nil),                 // 8: eolymp.judge.Scoreboard.Row
	(*wellknown.ExpressionBool)(nil),       // 9: eolymp.wellknown.ExpressionBool
}
var file_eolymp_judge_scoreboard_service_proto_depIdxs = []int32{
	5, // 0: eolymp.judge.DescribeScoreboardOutput.scoreboard:type_name -> eolymp.judge.Scoreboard
	6, // 1: eolymp.judge.ListScoreboardRowsInput.mode:type_name -> eolymp.judge.Scoreboard.Mode
	4, // 2: eolymp.judge.ListScoreboardRowsInput.filters:type_name -> eolymp.judge.ListScoreboardRowsInput.Filter
	7, // 3: eolymp.judge.ListScoreboardRowsInput.order:type_name -> eolymp.wellknown.Direction
	8, // 4: eolymp.judge.ListScoreboardRowsOutput.items:type_name -> eolymp.judge.Scoreboard.Row
	9, // 5: eolymp.judge.ListScoreboardRowsInput.Filter.unofficial:type_name -> eolymp.wellknown.ExpressionBool
	9, // 6: eolymp.judge.ListScoreboardRowsInput.Filter.disqualified:type_name -> eolymp.wellknown.ExpressionBool
	0, // 7: eolymp.judge.ScoreboardService.DescribeScoreboard:input_type -> eolymp.judge.DescribeScoreboardInput
	2, // 8: eolymp.judge.ScoreboardService.ListScoreboardRows:input_type -> eolymp.judge.ListScoreboardRowsInput
	1, // 9: eolymp.judge.ScoreboardService.DescribeScoreboard:output_type -> eolymp.judge.DescribeScoreboardOutput
	3, // 10: eolymp.judge.ScoreboardService.ListScoreboardRows:output_type -> eolymp.judge.ListScoreboardRowsOutput
	9, // [9:11] is the sub-list for method output_type
	7, // [7:9] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_eolymp_judge_scoreboard_service_proto_init() }
func file_eolymp_judge_scoreboard_service_proto_init() {
	if File_eolymp_judge_scoreboard_service_proto != nil {
		return
	}
	file_eolymp_judge_scoreboard_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_eolymp_judge_scoreboard_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_eolymp_judge_scoreboard_service_proto_goTypes,
		DependencyIndexes: file_eolymp_judge_scoreboard_service_proto_depIdxs,
		MessageInfos:      file_eolymp_judge_scoreboard_service_proto_msgTypes,
	}.Build()
	File_eolymp_judge_scoreboard_service_proto = out.File
	file_eolymp_judge_scoreboard_service_proto_rawDesc = nil
	file_eolymp_judge_scoreboard_service_proto_goTypes = nil
	file_eolymp_judge_scoreboard_service_proto_depIdxs = nil
}