// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.21.12
// source: eolymp/community/achievement_service.proto

package community

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

type AssignAchievementInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MemberId      string `protobuf:"bytes,1,opt,name=member_id,json=memberId,proto3" json:"member_id,omitempty"`
	AchievementId string `protobuf:"bytes,2,opt,name=achievement_id,json=achievementId,proto3" json:"achievement_id,omitempty"`
	// Types that are assignable to Quantity:
	//
	//	*AssignAchievementInput_SetTo
	//	*AssignAchievementInput_IncBy
	Quantity isAssignAchievementInput_Quantity `protobuf_oneof:"quantity"`
}

func (x *AssignAchievementInput) Reset() {
	*x = AssignAchievementInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_community_achievement_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AssignAchievementInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AssignAchievementInput) ProtoMessage() {}

func (x *AssignAchievementInput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_community_achievement_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AssignAchievementInput.ProtoReflect.Descriptor instead.
func (*AssignAchievementInput) Descriptor() ([]byte, []int) {
	return file_eolymp_community_achievement_service_proto_rawDescGZIP(), []int{0}
}

func (x *AssignAchievementInput) GetMemberId() string {
	if x != nil {
		return x.MemberId
	}
	return ""
}

func (x *AssignAchievementInput) GetAchievementId() string {
	if x != nil {
		return x.AchievementId
	}
	return ""
}

func (m *AssignAchievementInput) GetQuantity() isAssignAchievementInput_Quantity {
	if m != nil {
		return m.Quantity
	}
	return nil
}

func (x *AssignAchievementInput) GetSetTo() int32 {
	if x, ok := x.GetQuantity().(*AssignAchievementInput_SetTo); ok {
		return x.SetTo
	}
	return 0
}

func (x *AssignAchievementInput) GetIncBy() int32 {
	if x, ok := x.GetQuantity().(*AssignAchievementInput_IncBy); ok {
		return x.IncBy
	}
	return 0
}

type isAssignAchievementInput_Quantity interface {
	isAssignAchievementInput_Quantity()
}

type AssignAchievementInput_SetTo struct {
	SetTo int32 `protobuf:"varint,3,opt,name=set_to,json=setTo,proto3,oneof"`
}

type AssignAchievementInput_IncBy struct {
	IncBy int32 `protobuf:"varint,4,opt,name=inc_by,json=incBy,proto3,oneof"`
}

func (*AssignAchievementInput_SetTo) isAssignAchievementInput_Quantity() {}

func (*AssignAchievementInput_IncBy) isAssignAchievementInput_Quantity() {}

type AssignAchievementOutput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Quantity int32 `protobuf:"varint,3,opt,name=quantity,proto3" json:"quantity,omitempty"`
}

func (x *AssignAchievementOutput) Reset() {
	*x = AssignAchievementOutput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_community_achievement_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AssignAchievementOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AssignAchievementOutput) ProtoMessage() {}

func (x *AssignAchievementOutput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_community_achievement_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AssignAchievementOutput.ProtoReflect.Descriptor instead.
func (*AssignAchievementOutput) Descriptor() ([]byte, []int) {
	return file_eolymp_community_achievement_service_proto_rawDescGZIP(), []int{1}
}

func (x *AssignAchievementOutput) GetQuantity() int32 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

type UnassignAchievementInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MemberId      string `protobuf:"bytes,1,opt,name=member_id,json=memberId,proto3" json:"member_id,omitempty"`
	AchievementId string `protobuf:"bytes,2,opt,name=achievement_id,json=achievementId,proto3" json:"achievement_id,omitempty"`
}

func (x *UnassignAchievementInput) Reset() {
	*x = UnassignAchievementInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_community_achievement_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UnassignAchievementInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnassignAchievementInput) ProtoMessage() {}

func (x *UnassignAchievementInput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_community_achievement_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnassignAchievementInput.ProtoReflect.Descriptor instead.
func (*UnassignAchievementInput) Descriptor() ([]byte, []int) {
	return file_eolymp_community_achievement_service_proto_rawDescGZIP(), []int{2}
}

func (x *UnassignAchievementInput) GetMemberId() string {
	if x != nil {
		return x.MemberId
	}
	return ""
}

func (x *UnassignAchievementInput) GetAchievementId() string {
	if x != nil {
		return x.AchievementId
	}
	return ""
}

type UnassignAchievementOutput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UnassignAchievementOutput) Reset() {
	*x = UnassignAchievementOutput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_community_achievement_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UnassignAchievementOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnassignAchievementOutput) ProtoMessage() {}

func (x *UnassignAchievementOutput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_community_achievement_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnassignAchievementOutput.ProtoReflect.Descriptor instead.
func (*UnassignAchievementOutput) Descriptor() ([]byte, []int) {
	return file_eolymp_community_achievement_service_proto_rawDescGZIP(), []int{3}
}

type ListAchievementsInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MemberId string                        `protobuf:"bytes,1,opt,name=member_id,json=memberId,proto3" json:"member_id,omitempty"`
	Offset   int32                         `protobuf:"varint,10,opt,name=offset,proto3" json:"offset,omitempty"`
	Size     int32                         `protobuf:"varint,11,opt,name=size,proto3" json:"size,omitempty"`
	Filters  *ListAchievementsInput_Filter `protobuf:"bytes,40,opt,name=filters,proto3" json:"filters,omitempty"`
}

func (x *ListAchievementsInput) Reset() {
	*x = ListAchievementsInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_community_achievement_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListAchievementsInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListAchievementsInput) ProtoMessage() {}

func (x *ListAchievementsInput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_community_achievement_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListAchievementsInput.ProtoReflect.Descriptor instead.
func (*ListAchievementsInput) Descriptor() ([]byte, []int) {
	return file_eolymp_community_achievement_service_proto_rawDescGZIP(), []int{4}
}

func (x *ListAchievementsInput) GetMemberId() string {
	if x != nil {
		return x.MemberId
	}
	return ""
}

func (x *ListAchievementsInput) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *ListAchievementsInput) GetSize() int32 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *ListAchievementsInput) GetFilters() *ListAchievementsInput_Filter {
	if x != nil {
		return x.Filters
	}
	return nil
}

type ListAchievementsOutput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Total int32          `protobuf:"varint,1,opt,name=total,proto3" json:"total,omitempty"`
	Items []*Achievement `protobuf:"bytes,2,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *ListAchievementsOutput) Reset() {
	*x = ListAchievementsOutput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_community_achievement_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListAchievementsOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListAchievementsOutput) ProtoMessage() {}

func (x *ListAchievementsOutput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_community_achievement_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListAchievementsOutput.ProtoReflect.Descriptor instead.
func (*ListAchievementsOutput) Descriptor() ([]byte, []int) {
	return file_eolymp_community_achievement_service_proto_rawDescGZIP(), []int{5}
}

func (x *ListAchievementsOutput) GetTotal() int32 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *ListAchievementsOutput) GetItems() []*Achievement {
	if x != nil {
		return x.Items
	}
	return nil
}

type ListAchievementsInput_Filter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Query string                    `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
	Id    []*wellknown.ExpressionID `protobuf:"bytes,2,rep,name=id,proto3" json:"id,omitempty"`
}

func (x *ListAchievementsInput_Filter) Reset() {
	*x = ListAchievementsInput_Filter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_community_achievement_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListAchievementsInput_Filter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListAchievementsInput_Filter) ProtoMessage() {}

func (x *ListAchievementsInput_Filter) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_community_achievement_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListAchievementsInput_Filter.ProtoReflect.Descriptor instead.
func (*ListAchievementsInput_Filter) Descriptor() ([]byte, []int) {
	return file_eolymp_community_achievement_service_proto_rawDescGZIP(), []int{4, 0}
}

func (x *ListAchievementsInput_Filter) GetQuery() string {
	if x != nil {
		return x.Query
	}
	return ""
}

func (x *ListAchievementsInput_Filter) GetId() []*wellknown.ExpressionID {
	if x != nil {
		return x.Id
	}
	return nil
}

var File_eolymp_community_achievement_service_proto protoreflect.FileDescriptor

var file_eolymp_community_achievement_service_proto_rawDesc = []byte{
	0x0a, 0x2a, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69,
	0x74, 0x79, 0x2f, 0x61, 0x63, 0x68, 0x69, 0x65, 0x76, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x65, 0x6f,
	0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x1a, 0x1d,
	0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x22, 0x65,
	0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2f, 0x72, 0x61, 0x74, 0x65, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x22, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e,
	0x69, 0x74, 0x79, 0x2f, 0x61, 0x63, 0x68, 0x69, 0x65, 0x76, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x21, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x77, 0x65,
	0x6c, 0x6c, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x2f, 0x65, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9a, 0x01, 0x0a, 0x16, 0x41, 0x73, 0x73,
	0x69, 0x67, 0x6e, 0x41, 0x63, 0x68, 0x69, 0x65, 0x76, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x6e,
	0x70, 0x75, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x49, 0x64,
	0x12, 0x25, 0x0a, 0x0e, 0x61, 0x63, 0x68, 0x69, 0x65, 0x76, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x61, 0x63, 0x68, 0x69, 0x65, 0x76,
	0x65, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x06, 0x73, 0x65, 0x74, 0x5f, 0x74,
	0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x48, 0x00, 0x52, 0x05, 0x73, 0x65, 0x74, 0x54, 0x6f,
	0x12, 0x17, 0x0a, 0x06, 0x69, 0x6e, 0x63, 0x5f, 0x62, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05,
	0x48, 0x00, 0x52, 0x05, 0x69, 0x6e, 0x63, 0x42, 0x79, 0x42, 0x0a, 0x0a, 0x08, 0x71, 0x75, 0x61,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x22, 0x35, 0x0a, 0x17, 0x41, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x41,
	0x63, 0x68, 0x69, 0x65, 0x76, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74,
	0x12, 0x1a, 0x0a, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x22, 0x5e, 0x0a, 0x18,
	0x55, 0x6e, 0x61, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x41, 0x63, 0x68, 0x69, 0x65, 0x76, 0x65, 0x6d,
	0x65, 0x6e, 0x74, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x65, 0x6d, 0x62,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6d, 0x65, 0x6d,
	0x62, 0x65, 0x72, 0x49, 0x64, 0x12, 0x25, 0x0a, 0x0e, 0x61, 0x63, 0x68, 0x69, 0x65, 0x76, 0x65,
	0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x61,
	0x63, 0x68, 0x69, 0x65, 0x76, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x22, 0x1b, 0x0a, 0x19,
	0x55, 0x6e, 0x61, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x41, 0x63, 0x68, 0x69, 0x65, 0x76, 0x65, 0x6d,
	0x65, 0x6e, 0x74, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x22, 0xfa, 0x01, 0x0a, 0x15, 0x4c, 0x69,
	0x73, 0x74, 0x41, 0x63, 0x68, 0x69, 0x65, 0x76, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x49, 0x6e,
	0x70, 0x75, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x49, 0x64,
	0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65,
	0x18, 0x0b, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x48, 0x0a, 0x07,
	0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x18, 0x28, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2e, 0x2e,
	0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79,
	0x2e, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x63, 0x68, 0x69, 0x65, 0x76, 0x65, 0x6d, 0x65, 0x6e, 0x74,
	0x73, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x2e, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x52, 0x07, 0x66,
	0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x1a, 0x4e, 0x0a, 0x06, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72,
	0x12, 0x14, 0x0a, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x12, 0x2e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x77, 0x65, 0x6c, 0x6c,
	0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x2e, 0x45, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x49, 0x44, 0x52, 0x02, 0x69, 0x64, 0x22, 0x63, 0x0a, 0x16, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x63,
	0x68, 0x69, 0x65, 0x76, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74,
	0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x33, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x2e, 0x41, 0x63, 0x68, 0x69, 0x65, 0x76, 0x65,
	0x6d, 0x65, 0x6e, 0x74, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x32, 0x81, 0x05, 0x0a, 0x12,
	0x41, 0x63, 0x68, 0x69, 0x65, 0x76, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0xd1, 0x01, 0x0a, 0x11, 0x41, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x41, 0x63, 0x68,
	0x69, 0x65, 0x76, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x28, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d,
	0x70, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x2e, 0x41, 0x73, 0x73, 0x69,
	0x67, 0x6e, 0x41, 0x63, 0x68, 0x69, 0x65, 0x76, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x70,
	0x75, 0x74, 0x1a, 0x29, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x75, 0x6e, 0x69, 0x74, 0x79, 0x2e, 0x41, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x41, 0x63, 0x68, 0x69,
	0x65, 0x76, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x22, 0x67, 0xea,
	0xe2, 0x0a, 0x0b, 0xf5, 0xe2, 0x0a, 0x00, 0x00, 0xa0, 0x40, 0xf8, 0xe2, 0x0a, 0x14, 0x82, 0xe3,
	0x0a, 0x1a, 0x8a, 0xe3, 0x0a, 0x16, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x3a,
	0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x3a, 0x77, 0x72, 0x69, 0x74, 0x65, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x34, 0x1a, 0x32, 0x2f, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x2f, 0x7b, 0x6d, 0x65,
	0x6d, 0x62, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x7d, 0x2f, 0x61, 0x63, 0x68, 0x69, 0x65, 0x76, 0x65,
	0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x7b, 0x61, 0x63, 0x68, 0x69, 0x65, 0x76, 0x65, 0x6d, 0x65,
	0x6e, 0x74, 0x5f, 0x69, 0x64, 0x7d, 0x12, 0xd7, 0x01, 0x0a, 0x13, 0x55, 0x6e, 0x61, 0x73, 0x73,
	0x69, 0x67, 0x6e, 0x41, 0x63, 0x68, 0x69, 0x65, 0x76, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x2a,
	0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74,
	0x79, 0x2e, 0x55, 0x6e, 0x61, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x41, 0x63, 0x68, 0x69, 0x65, 0x76,
	0x65, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x2b, 0x2e, 0x65, 0x6f, 0x6c,
	0x79, 0x6d, 0x70, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x2e, 0x55, 0x6e,
	0x61, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x41, 0x63, 0x68, 0x69, 0x65, 0x76, 0x65, 0x6d, 0x65, 0x6e,
	0x74, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x22, 0x67, 0xea, 0xe2, 0x0a, 0x0b, 0xf5, 0xe2, 0x0a,
	0x00, 0x00, 0xa0, 0x40, 0xf8, 0xe2, 0x0a, 0x14, 0x82, 0xe3, 0x0a, 0x1a, 0x8a, 0xe3, 0x0a, 0x16,
	0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x3a, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72,
	0x3a, 0x77, 0x72, 0x69, 0x74, 0x65, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x34, 0x2a, 0x32, 0x2f, 0x6d,
	0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x2f, 0x7b, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x7d, 0x2f, 0x61, 0x63, 0x68, 0x69, 0x65, 0x76, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2f,
	0x7b, 0x61, 0x63, 0x68, 0x69, 0x65, 0x76, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x7d,
	0x12, 0xbc, 0x01, 0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x63, 0x68, 0x69, 0x65, 0x76, 0x65,
	0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x27, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x63, 0x68,
	0x69, 0x65, 0x76, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x28,
	0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74,
	0x79, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x63, 0x68, 0x69, 0x65, 0x76, 0x65, 0x6d, 0x65, 0x6e,
	0x74, 0x73, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x22, 0x55, 0xea, 0xe2, 0x0a, 0x0b, 0xf5, 0xe2,
	0x0a, 0x00, 0x00, 0xa0, 0x40, 0xf8, 0xe2, 0x0a, 0x14, 0x82, 0xe3, 0x0a, 0x19, 0x8a, 0xe3, 0x0a,
	0x15, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x3a, 0x6d, 0x65, 0x6d, 0x62, 0x65,
	0x72, 0x3a, 0x72, 0x65, 0x61, 0x64, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x23, 0x12, 0x21, 0x2f, 0x6d,
	0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x2f, 0x7b, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x7d, 0x2f, 0x61, 0x63, 0x68, 0x69, 0x65, 0x76, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x42,
	0x35, 0x5a, 0x33, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6f,
	0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x67, 0x6f, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x65, 0x6f, 0x6c, 0x79,
	0x6d, 0x70, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x3b, 0x63, 0x6f, 0x6d,
	0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_eolymp_community_achievement_service_proto_rawDescOnce sync.Once
	file_eolymp_community_achievement_service_proto_rawDescData = file_eolymp_community_achievement_service_proto_rawDesc
)

func file_eolymp_community_achievement_service_proto_rawDescGZIP() []byte {
	file_eolymp_community_achievement_service_proto_rawDescOnce.Do(func() {
		file_eolymp_community_achievement_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_eolymp_community_achievement_service_proto_rawDescData)
	})
	return file_eolymp_community_achievement_service_proto_rawDescData
}

var file_eolymp_community_achievement_service_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_eolymp_community_achievement_service_proto_goTypes = []any{
	(*AssignAchievementInput)(nil),       // 0: eolymp.community.AssignAchievementInput
	(*AssignAchievementOutput)(nil),      // 1: eolymp.community.AssignAchievementOutput
	(*UnassignAchievementInput)(nil),     // 2: eolymp.community.UnassignAchievementInput
	(*UnassignAchievementOutput)(nil),    // 3: eolymp.community.UnassignAchievementOutput
	(*ListAchievementsInput)(nil),        // 4: eolymp.community.ListAchievementsInput
	(*ListAchievementsOutput)(nil),       // 5: eolymp.community.ListAchievementsOutput
	(*ListAchievementsInput_Filter)(nil), // 6: eolymp.community.ListAchievementsInput.Filter
	(*Achievement)(nil),                  // 7: eolymp.community.Achievement
	(*wellknown.ExpressionID)(nil),       // 8: eolymp.wellknown.ExpressionID
}
var file_eolymp_community_achievement_service_proto_depIdxs = []int32{
	6, // 0: eolymp.community.ListAchievementsInput.filters:type_name -> eolymp.community.ListAchievementsInput.Filter
	7, // 1: eolymp.community.ListAchievementsOutput.items:type_name -> eolymp.community.Achievement
	8, // 2: eolymp.community.ListAchievementsInput.Filter.id:type_name -> eolymp.wellknown.ExpressionID
	0, // 3: eolymp.community.AchievementService.AssignAchievement:input_type -> eolymp.community.AssignAchievementInput
	2, // 4: eolymp.community.AchievementService.UnassignAchievement:input_type -> eolymp.community.UnassignAchievementInput
	4, // 5: eolymp.community.AchievementService.ListAchievements:input_type -> eolymp.community.ListAchievementsInput
	1, // 6: eolymp.community.AchievementService.AssignAchievement:output_type -> eolymp.community.AssignAchievementOutput
	3, // 7: eolymp.community.AchievementService.UnassignAchievement:output_type -> eolymp.community.UnassignAchievementOutput
	5, // 8: eolymp.community.AchievementService.ListAchievements:output_type -> eolymp.community.ListAchievementsOutput
	6, // [6:9] is the sub-list for method output_type
	3, // [3:6] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_eolymp_community_achievement_service_proto_init() }
func file_eolymp_community_achievement_service_proto_init() {
	if File_eolymp_community_achievement_service_proto != nil {
		return
	}
	file_eolymp_community_achievement_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_eolymp_community_achievement_service_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*AssignAchievementInput); i {
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
		file_eolymp_community_achievement_service_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*AssignAchievementOutput); i {
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
		file_eolymp_community_achievement_service_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*UnassignAchievementInput); i {
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
		file_eolymp_community_achievement_service_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*UnassignAchievementOutput); i {
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
		file_eolymp_community_achievement_service_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*ListAchievementsInput); i {
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
		file_eolymp_community_achievement_service_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*ListAchievementsOutput); i {
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
		file_eolymp_community_achievement_service_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*ListAchievementsInput_Filter); i {
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
	file_eolymp_community_achievement_service_proto_msgTypes[0].OneofWrappers = []any{
		(*AssignAchievementInput_SetTo)(nil),
		(*AssignAchievementInput_IncBy)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_eolymp_community_achievement_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_eolymp_community_achievement_service_proto_goTypes,
		DependencyIndexes: file_eolymp_community_achievement_service_proto_depIdxs,
		MessageInfos:      file_eolymp_community_achievement_service_proto_msgTypes,
	}.Build()
	File_eolymp_community_achievement_service_proto = out.File
	file_eolymp_community_achievement_service_proto_rawDesc = nil
	file_eolymp_community_achievement_service_proto_goTypes = nil
	file_eolymp_community_achievement_service_proto_depIdxs = nil
}
