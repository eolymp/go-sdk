// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.21.12
// source: eolymp/rating/rating_service.proto

package rating

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

type ListRatingInput_Sortable int32

const (
	ListRatingInput_TIMESTAMP ListRatingInput_Sortable = 0
)

// Enum value maps for ListRatingInput_Sortable.
var (
	ListRatingInput_Sortable_name = map[int32]string{
		0: "TIMESTAMP",
	}
	ListRatingInput_Sortable_value = map[string]int32{
		"TIMESTAMP": 0,
	}
)

func (x ListRatingInput_Sortable) Enum() *ListRatingInput_Sortable {
	p := new(ListRatingInput_Sortable)
	*p = x
	return p
}

func (x ListRatingInput_Sortable) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ListRatingInput_Sortable) Descriptor() protoreflect.EnumDescriptor {
	return file_eolymp_rating_rating_service_proto_enumTypes[0].Descriptor()
}

func (ListRatingInput_Sortable) Type() protoreflect.EnumType {
	return &file_eolymp_rating_rating_service_proto_enumTypes[0]
}

func (x ListRatingInput_Sortable) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ListRatingInput_Sortable.Descriptor instead.
func (ListRatingInput_Sortable) EnumDescriptor() ([]byte, []int) {
	return file_eolymp_rating_rating_service_proto_rawDescGZIP(), []int{8, 0}
}

type SetRatingInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Rating *Rating `protobuf:"bytes,1,opt,name=rating,proto3" json:"rating,omitempty"`
}

func (x *SetRatingInput) Reset() {
	*x = SetRatingInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_rating_rating_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetRatingInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetRatingInput) ProtoMessage() {}

func (x *SetRatingInput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_rating_rating_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetRatingInput.ProtoReflect.Descriptor instead.
func (*SetRatingInput) Descriptor() ([]byte, []int) {
	return file_eolymp_rating_rating_service_proto_rawDescGZIP(), []int{0}
}

func (x *SetRatingInput) GetRating() *Rating {
	if x != nil {
		return x.Rating
	}
	return nil
}

type SetRatingOutput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RatingId string `protobuf:"bytes,1,opt,name=rating_id,json=ratingId,proto3" json:"rating_id,omitempty"`
}

func (x *SetRatingOutput) Reset() {
	*x = SetRatingOutput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_rating_rating_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetRatingOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetRatingOutput) ProtoMessage() {}

func (x *SetRatingOutput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_rating_rating_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetRatingOutput.ProtoReflect.Descriptor instead.
func (*SetRatingOutput) Descriptor() ([]byte, []int) {
	return file_eolymp_rating_rating_service_proto_rawDescGZIP(), []int{1}
}

func (x *SetRatingOutput) GetRatingId() string {
	if x != nil {
		return x.RatingId
	}
	return ""
}

type UpdateRatingInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RatingId string  `protobuf:"bytes,1,opt,name=rating_id,json=ratingId,proto3" json:"rating_id,omitempty"`
	Rating   *Rating `protobuf:"bytes,2,opt,name=rating,proto3" json:"rating,omitempty"`
}

func (x *UpdateRatingInput) Reset() {
	*x = UpdateRatingInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_rating_rating_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateRatingInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateRatingInput) ProtoMessage() {}

func (x *UpdateRatingInput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_rating_rating_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateRatingInput.ProtoReflect.Descriptor instead.
func (*UpdateRatingInput) Descriptor() ([]byte, []int) {
	return file_eolymp_rating_rating_service_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateRatingInput) GetRatingId() string {
	if x != nil {
		return x.RatingId
	}
	return ""
}

func (x *UpdateRatingInput) GetRating() *Rating {
	if x != nil {
		return x.Rating
	}
	return nil
}

type UpdateRatingOutput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateRatingOutput) Reset() {
	*x = UpdateRatingOutput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_rating_rating_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateRatingOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateRatingOutput) ProtoMessage() {}

func (x *UpdateRatingOutput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_rating_rating_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateRatingOutput.ProtoReflect.Descriptor instead.
func (*UpdateRatingOutput) Descriptor() ([]byte, []int) {
	return file_eolymp_rating_rating_service_proto_rawDescGZIP(), []int{3}
}

type DeleteRatingInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RatingId string  `protobuf:"bytes,1,opt,name=rating_id,json=ratingId,proto3" json:"rating_id,omitempty"`
	Rating   *Rating `protobuf:"bytes,2,opt,name=rating,proto3" json:"rating,omitempty"`
}

func (x *DeleteRatingInput) Reset() {
	*x = DeleteRatingInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_rating_rating_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteRatingInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteRatingInput) ProtoMessage() {}

func (x *DeleteRatingInput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_rating_rating_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteRatingInput.ProtoReflect.Descriptor instead.
func (*DeleteRatingInput) Descriptor() ([]byte, []int) {
	return file_eolymp_rating_rating_service_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteRatingInput) GetRatingId() string {
	if x != nil {
		return x.RatingId
	}
	return ""
}

func (x *DeleteRatingInput) GetRating() *Rating {
	if x != nil {
		return x.Rating
	}
	return nil
}

type DeleteRatingOutput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteRatingOutput) Reset() {
	*x = DeleteRatingOutput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_rating_rating_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteRatingOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteRatingOutput) ProtoMessage() {}

func (x *DeleteRatingOutput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_rating_rating_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteRatingOutput.ProtoReflect.Descriptor instead.
func (*DeleteRatingOutput) Descriptor() ([]byte, []int) {
	return file_eolymp_rating_rating_service_proto_rawDescGZIP(), []int{5}
}

type DescribeRatingInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RatingId string `protobuf:"bytes,1,opt,name=rating_id,json=ratingId,proto3" json:"rating_id,omitempty"`
}

func (x *DescribeRatingInput) Reset() {
	*x = DescribeRatingInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_rating_rating_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DescribeRatingInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescribeRatingInput) ProtoMessage() {}

func (x *DescribeRatingInput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_rating_rating_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DescribeRatingInput.ProtoReflect.Descriptor instead.
func (*DescribeRatingInput) Descriptor() ([]byte, []int) {
	return file_eolymp_rating_rating_service_proto_rawDescGZIP(), []int{6}
}

func (x *DescribeRatingInput) GetRatingId() string {
	if x != nil {
		return x.RatingId
	}
	return ""
}

type DescribeRatingOutput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Rating *Rating `protobuf:"bytes,1,opt,name=rating,proto3" json:"rating,omitempty"`
}

func (x *DescribeRatingOutput) Reset() {
	*x = DescribeRatingOutput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_rating_rating_service_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DescribeRatingOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescribeRatingOutput) ProtoMessage() {}

func (x *DescribeRatingOutput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_rating_rating_service_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DescribeRatingOutput.ProtoReflect.Descriptor instead.
func (*DescribeRatingOutput) Descriptor() ([]byte, []int) {
	return file_eolymp_rating_rating_service_proto_rawDescGZIP(), []int{7}
}

func (x *DescribeRatingOutput) GetRating() *Rating {
	if x != nil {
		return x.Rating
	}
	return nil
}

type ListRatingInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MemberId string                   `protobuf:"bytes,1,opt,name=member_id,json=memberId,proto3" json:"member_id,omitempty"`
	Offset   int32                    `protobuf:"varint,10,opt,name=offset,proto3" json:"offset,omitempty"`
	Size     int32                    `protobuf:"varint,11,opt,name=size,proto3" json:"size,omitempty"`
	After    string                   `protobuf:"bytes,12,opt,name=after,proto3" json:"after,omitempty"`
	Filters  *ListRatingInput_Filter  `protobuf:"bytes,40,opt,name=filters,proto3" json:"filters,omitempty"`
	Sort     ListRatingInput_Sortable `protobuf:"varint,50,opt,name=sort,proto3,enum=eolymp.rating.ListRatingInput_Sortable" json:"sort,omitempty"`
	Order    wellknown.Direction      `protobuf:"varint,51,opt,name=order,proto3,enum=eolymp.wellknown.Direction" json:"order,omitempty"`
}

func (x *ListRatingInput) Reset() {
	*x = ListRatingInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_rating_rating_service_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListRatingInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListRatingInput) ProtoMessage() {}

func (x *ListRatingInput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_rating_rating_service_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListRatingInput.ProtoReflect.Descriptor instead.
func (*ListRatingInput) Descriptor() ([]byte, []int) {
	return file_eolymp_rating_rating_service_proto_rawDescGZIP(), []int{8}
}

func (x *ListRatingInput) GetMemberId() string {
	if x != nil {
		return x.MemberId
	}
	return ""
}

func (x *ListRatingInput) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *ListRatingInput) GetSize() int32 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *ListRatingInput) GetAfter() string {
	if x != nil {
		return x.After
	}
	return ""
}

func (x *ListRatingInput) GetFilters() *ListRatingInput_Filter {
	if x != nil {
		return x.Filters
	}
	return nil
}

func (x *ListRatingInput) GetSort() ListRatingInput_Sortable {
	if x != nil {
		return x.Sort
	}
	return ListRatingInput_TIMESTAMP
}

func (x *ListRatingInput) GetOrder() wellknown.Direction {
	if x != nil {
		return x.Order
	}
	return wellknown.Direction(0)
}

type ListRatingOutput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Total int32     `protobuf:"varint,1,opt,name=total,proto3" json:"total,omitempty"`
	Items []*Rating `protobuf:"bytes,2,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *ListRatingOutput) Reset() {
	*x = ListRatingOutput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_rating_rating_service_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListRatingOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListRatingOutput) ProtoMessage() {}

func (x *ListRatingOutput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_rating_rating_service_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListRatingOutput.ProtoReflect.Descriptor instead.
func (*ListRatingOutput) Descriptor() ([]byte, []int) {
	return file_eolymp_rating_rating_service_proto_rawDescGZIP(), []int{9}
}

func (x *ListRatingOutput) GetTotal() int32 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *ListRatingOutput) GetItems() []*Rating {
	if x != nil {
		return x.Items
	}
	return nil
}

type ListRatingInput_Filter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        []*wellknown.ExpressionID        `protobuf:"bytes,1,rep,name=id,proto3" json:"id,omitempty"`
	ContestId []*wellknown.ExpressionID        `protobuf:"bytes,2,rep,name=contest_id,json=contestId,proto3" json:"contest_id,omitempty"`
	Timestamp []*wellknown.ExpressionTimestamp `protobuf:"bytes,3,rep,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *ListRatingInput_Filter) Reset() {
	*x = ListRatingInput_Filter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_rating_rating_service_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListRatingInput_Filter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListRatingInput_Filter) ProtoMessage() {}

func (x *ListRatingInput_Filter) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_rating_rating_service_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListRatingInput_Filter.ProtoReflect.Descriptor instead.
func (*ListRatingInput_Filter) Descriptor() ([]byte, []int) {
	return file_eolymp_rating_rating_service_proto_rawDescGZIP(), []int{8, 0}
}

func (x *ListRatingInput_Filter) GetId() []*wellknown.ExpressionID {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *ListRatingInput_Filter) GetContestId() []*wellknown.ExpressionID {
	if x != nil {
		return x.ContestId
	}
	return nil
}

func (x *ListRatingInput_Filter) GetTimestamp() []*wellknown.ExpressionTimestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

var File_eolymp_rating_rating_service_proto protoreflect.FileDescriptor

var file_eolymp_rating_rating_service_proto_rawDesc = []byte{
	0x0a, 0x22, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x2f,
	0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x72, 0x61, 0x74,
	0x69, 0x6e, 0x67, 0x1a, 0x1d, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x22, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x72, 0x61, 0x74, 0x65, 0x6c, 0x69, 0x6d, 0x69, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x61,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x63, 0x6f, 0x70, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1a, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x72,
	0x61, 0x74, 0x69, 0x6e, 0x67, 0x2f, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x20, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x77, 0x65, 0x6c, 0x6c, 0x6b,
	0x6e, 0x6f, 0x77, 0x6e, 0x2f, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x21, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x77, 0x65, 0x6c,
	0x6c, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x2f, 0x65, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x3f, 0x0a, 0x0e, 0x53, 0x65, 0x74, 0x52, 0x61,
	0x74, 0x69, 0x6e, 0x67, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x2d, 0x0a, 0x06, 0x72, 0x61, 0x74,
	0x69, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x65, 0x6f, 0x6c, 0x79,
	0x6d, 0x70, 0x2e, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67,
	0x52, 0x06, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x22, 0x2e, 0x0a, 0x0f, 0x53, 0x65, 0x74, 0x52,
	0x61, 0x74, 0x69, 0x6e, 0x67, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x72,
	0x61, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x49, 0x64, 0x22, 0x5f, 0x0a, 0x11, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x1b, 0x0a,
	0x09, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x49, 0x64, 0x12, 0x2d, 0x0a, 0x06, 0x72, 0x61,
	0x74, 0x69, 0x6e, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x65, 0x6f, 0x6c,
	0x79, 0x6d, 0x70, 0x2e, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x52, 0x61, 0x74, 0x69, 0x6e,
	0x67, 0x52, 0x06, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x22, 0x14, 0x0a, 0x12, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x22,
	0x5f, 0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x49,
	0x6e, 0x70, 0x75, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x49,
	0x64, 0x12, 0x2d, 0x0a, 0x06, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x15, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x72, 0x61, 0x74, 0x69, 0x6e,
	0x67, 0x2e, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x06, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67,
	0x22, 0x14, 0x0a, 0x12, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67,
	0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x22, 0x32, 0x0a, 0x13, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x62, 0x65, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x1b, 0x0a,
	0x09, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x49, 0x64, 0x22, 0x45, 0x0a, 0x14, 0x44, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x4f, 0x75, 0x74, 0x70,
	0x75, 0x74, 0x12, 0x2d, 0x0a, 0x06, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x15, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x72, 0x61, 0x74, 0x69,
	0x6e, 0x67, 0x2e, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x06, 0x72, 0x61, 0x74, 0x69, 0x6e,
	0x67, 0x22, 0xfb, 0x03, 0x0a, 0x0f, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67,
	0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72,
	0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69,
	0x7a, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x61, 0x66, 0x74, 0x65, 0x72, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x61,
	0x66, 0x74, 0x65, 0x72, 0x12, 0x3f, 0x0a, 0x07, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x18,
	0x28, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x72,
	0x61, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67,
	0x49, 0x6e, 0x70, 0x75, 0x74, 0x2e, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x52, 0x07, 0x66, 0x69,
	0x6c, 0x74, 0x65, 0x72, 0x73, 0x12, 0x3b, 0x0a, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x18, 0x32, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x27, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x72, 0x61, 0x74,
	0x69, 0x6e, 0x67, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x49, 0x6e,
	0x70, 0x75, 0x74, 0x2e, 0x53, 0x6f, 0x72, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x52, 0x04, 0x73, 0x6f,
	0x72, 0x74, 0x12, 0x31, 0x0a, 0x05, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x18, 0x33, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x1b, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x77, 0x65, 0x6c, 0x6c, 0x6b,
	0x6e, 0x6f, 0x77, 0x6e, 0x2e, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x05,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x1a, 0xbc, 0x01, 0x0a, 0x06, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72,
	0x12, 0x2e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x65,
	0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x77, 0x65, 0x6c, 0x6c, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x2e,
	0x45, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x3d, 0x0a, 0x0a, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x77, 0x65,
	0x6c, 0x6c, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x2e, 0x45, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x49, 0x44, 0x52, 0x09, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x73, 0x74, 0x49, 0x64, 0x12,
	0x43, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x03, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x25, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x77, 0x65, 0x6c, 0x6c,
	0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x2e, 0x45, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x22, 0x19, 0x0a, 0x08, 0x53, 0x6f, 0x72, 0x74, 0x61, 0x62, 0x6c, 0x65,
	0x12, 0x0d, 0x0a, 0x09, 0x54, 0x49, 0x4d, 0x45, 0x53, 0x54, 0x41, 0x4d, 0x50, 0x10, 0x00, 0x22,
	0x55, 0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x4f, 0x75, 0x74,
	0x70, 0x75, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x2b, 0x0a, 0x05, 0x69, 0x74, 0x65,
	0x6d, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d,
	0x70, 0x2e, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x52,
	0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x32, 0xa0, 0x06, 0x0a, 0x0d, 0x52, 0x61, 0x74, 0x69, 0x6e,
	0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x88, 0x01, 0x0a, 0x09, 0x53, 0x65, 0x74,
	0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x1d, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e,
	0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x53, 0x65, 0x74, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67,
	0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x1e, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x72,
	0x61, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x53, 0x65, 0x74, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x4f,
	0x75, 0x74, 0x70, 0x75, 0x74, 0x22, 0x3c, 0xea, 0xe2, 0x0a, 0x0b, 0xf5, 0xe2, 0x0a, 0x00, 0x00,
	0x20, 0x41, 0xf8, 0xe2, 0x0a, 0x64, 0x82, 0xe3, 0x0a, 0x1a, 0x8a, 0xe3, 0x0a, 0x16, 0x63, 0x6f,
	0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x3a, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x3a, 0x77,
	0x72, 0x69, 0x74, 0x65, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x09, 0x22, 0x07, 0x2f, 0x72, 0x61, 0x74,
	0x69, 0x6e, 0x67, 0x12, 0x9d, 0x01, 0x0a, 0x0c, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x61,
	0x74, 0x69, 0x6e, 0x67, 0x12, 0x20, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x72, 0x61,
	0x74, 0x69, 0x6e, 0x67, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x61, 0x74, 0x69, 0x6e,
	0x67, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x21, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e,
	0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x61, 0x74,
	0x69, 0x6e, 0x67, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x22, 0x48, 0xea, 0xe2, 0x0a, 0x0b, 0xf5,
	0xe2, 0x0a, 0x00, 0x00, 0x20, 0x41, 0xf8, 0xe2, 0x0a, 0x64, 0x82, 0xe3, 0x0a, 0x1a, 0x8a, 0xe3,
	0x0a, 0x16, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x3a, 0x6d, 0x65, 0x6d, 0x62,
	0x65, 0x72, 0x3a, 0x77, 0x72, 0x69, 0x74, 0x65, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x15, 0x1a, 0x13,
	0x2f, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x2f, 0x7b, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x5f,
	0x69, 0x64, 0x7d, 0x12, 0x9d, 0x01, 0x0a, 0x0c, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x61,
	0x74, 0x69, 0x6e, 0x67, 0x12, 0x20, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x72, 0x61,
	0x74, 0x69, 0x6e, 0x67, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x61, 0x74, 0x69, 0x6e,
	0x67, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x21, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e,
	0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x61, 0x74,
	0x69, 0x6e, 0x67, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x22, 0x48, 0xea, 0xe2, 0x0a, 0x0b, 0xf5,
	0xe2, 0x0a, 0x00, 0x00, 0x20, 0x41, 0xf8, 0xe2, 0x0a, 0x64, 0x82, 0xe3, 0x0a, 0x1a, 0x8a, 0xe3,
	0x0a, 0x16, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x3a, 0x6d, 0x65, 0x6d, 0x62,
	0x65, 0x72, 0x3a, 0x77, 0x72, 0x69, 0x74, 0x65, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x15, 0x2a, 0x13,
	0x2f, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x2f, 0x7b, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x5f,
	0x69, 0x64, 0x7d, 0x12, 0xa2, 0x01, 0x0a, 0x0e, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65,
	0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x22, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e,
	0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x52,
	0x61, 0x74, 0x69, 0x6e, 0x67, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x23, 0x2e, 0x65, 0x6f, 0x6c,
	0x79, 0x6d, 0x70, 0x2e, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x44, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x62, 0x65, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x22,
	0x47, 0xea, 0xe2, 0x0a, 0x0b, 0xf5, 0xe2, 0x0a, 0x00, 0x00, 0xa0, 0x41, 0xf8, 0xe2, 0x0a, 0x64,
	0x82, 0xe3, 0x0a, 0x19, 0x8a, 0xe3, 0x0a, 0x15, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74,
	0x79, 0x3a, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x3a, 0x72, 0x65, 0x61, 0x64, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x15, 0x12, 0x13, 0x2f, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x2f, 0x7b, 0x72, 0x61,
	0x74, 0x69, 0x6e, 0x67, 0x5f, 0x69, 0x64, 0x7d, 0x12, 0x9e, 0x01, 0x0a, 0x0a, 0x4c, 0x69, 0x73,
	0x74, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x1e, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70,
	0x2e, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x61, 0x74, 0x69,
	0x6e, 0x67, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x1f, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70,
	0x2e, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x61, 0x74, 0x69,
	0x6e, 0x67, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x22, 0x4f, 0xea, 0xe2, 0x0a, 0x0b, 0xf5, 0xe2,
	0x0a, 0x00, 0x00, 0xa0, 0x40, 0xf8, 0xe2, 0x0a, 0x14, 0x82, 0xe3, 0x0a, 0x19, 0x8a, 0xe3, 0x0a,
	0x15, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x3a, 0x6d, 0x65, 0x6d, 0x62, 0x65,
	0x72, 0x3a, 0x72, 0x65, 0x61, 0x64, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1d, 0x12, 0x1b, 0x2f, 0x6d,
	0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x2f, 0x7b, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x7d, 0x2f, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x42, 0x2f, 0x5a, 0x2d, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x67,
	0x6f, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x72, 0x61, 0x74,
	0x69, 0x6e, 0x67, 0x3b, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_eolymp_rating_rating_service_proto_rawDescOnce sync.Once
	file_eolymp_rating_rating_service_proto_rawDescData = file_eolymp_rating_rating_service_proto_rawDesc
)

func file_eolymp_rating_rating_service_proto_rawDescGZIP() []byte {
	file_eolymp_rating_rating_service_proto_rawDescOnce.Do(func() {
		file_eolymp_rating_rating_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_eolymp_rating_rating_service_proto_rawDescData)
	})
	return file_eolymp_rating_rating_service_proto_rawDescData
}

var file_eolymp_rating_rating_service_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_eolymp_rating_rating_service_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_eolymp_rating_rating_service_proto_goTypes = []any{
	(ListRatingInput_Sortable)(0),         // 0: eolymp.rating.ListRatingInput.Sortable
	(*SetRatingInput)(nil),                // 1: eolymp.rating.SetRatingInput
	(*SetRatingOutput)(nil),               // 2: eolymp.rating.SetRatingOutput
	(*UpdateRatingInput)(nil),             // 3: eolymp.rating.UpdateRatingInput
	(*UpdateRatingOutput)(nil),            // 4: eolymp.rating.UpdateRatingOutput
	(*DeleteRatingInput)(nil),             // 5: eolymp.rating.DeleteRatingInput
	(*DeleteRatingOutput)(nil),            // 6: eolymp.rating.DeleteRatingOutput
	(*DescribeRatingInput)(nil),           // 7: eolymp.rating.DescribeRatingInput
	(*DescribeRatingOutput)(nil),          // 8: eolymp.rating.DescribeRatingOutput
	(*ListRatingInput)(nil),               // 9: eolymp.rating.ListRatingInput
	(*ListRatingOutput)(nil),              // 10: eolymp.rating.ListRatingOutput
	(*ListRatingInput_Filter)(nil),        // 11: eolymp.rating.ListRatingInput.Filter
	(*Rating)(nil),                        // 12: eolymp.rating.Rating
	(wellknown.Direction)(0),              // 13: eolymp.wellknown.Direction
	(*wellknown.ExpressionID)(nil),        // 14: eolymp.wellknown.ExpressionID
	(*wellknown.ExpressionTimestamp)(nil), // 15: eolymp.wellknown.ExpressionTimestamp
}
var file_eolymp_rating_rating_service_proto_depIdxs = []int32{
	12, // 0: eolymp.rating.SetRatingInput.rating:type_name -> eolymp.rating.Rating
	12, // 1: eolymp.rating.UpdateRatingInput.rating:type_name -> eolymp.rating.Rating
	12, // 2: eolymp.rating.DeleteRatingInput.rating:type_name -> eolymp.rating.Rating
	12, // 3: eolymp.rating.DescribeRatingOutput.rating:type_name -> eolymp.rating.Rating
	11, // 4: eolymp.rating.ListRatingInput.filters:type_name -> eolymp.rating.ListRatingInput.Filter
	0,  // 5: eolymp.rating.ListRatingInput.sort:type_name -> eolymp.rating.ListRatingInput.Sortable
	13, // 6: eolymp.rating.ListRatingInput.order:type_name -> eolymp.wellknown.Direction
	12, // 7: eolymp.rating.ListRatingOutput.items:type_name -> eolymp.rating.Rating
	14, // 8: eolymp.rating.ListRatingInput.Filter.id:type_name -> eolymp.wellknown.ExpressionID
	14, // 9: eolymp.rating.ListRatingInput.Filter.contest_id:type_name -> eolymp.wellknown.ExpressionID
	15, // 10: eolymp.rating.ListRatingInput.Filter.timestamp:type_name -> eolymp.wellknown.ExpressionTimestamp
	1,  // 11: eolymp.rating.RatingService.SetRating:input_type -> eolymp.rating.SetRatingInput
	3,  // 12: eolymp.rating.RatingService.UpdateRating:input_type -> eolymp.rating.UpdateRatingInput
	5,  // 13: eolymp.rating.RatingService.DeleteRating:input_type -> eolymp.rating.DeleteRatingInput
	7,  // 14: eolymp.rating.RatingService.DescribeRating:input_type -> eolymp.rating.DescribeRatingInput
	9,  // 15: eolymp.rating.RatingService.ListRating:input_type -> eolymp.rating.ListRatingInput
	2,  // 16: eolymp.rating.RatingService.SetRating:output_type -> eolymp.rating.SetRatingOutput
	4,  // 17: eolymp.rating.RatingService.UpdateRating:output_type -> eolymp.rating.UpdateRatingOutput
	6,  // 18: eolymp.rating.RatingService.DeleteRating:output_type -> eolymp.rating.DeleteRatingOutput
	8,  // 19: eolymp.rating.RatingService.DescribeRating:output_type -> eolymp.rating.DescribeRatingOutput
	10, // 20: eolymp.rating.RatingService.ListRating:output_type -> eolymp.rating.ListRatingOutput
	16, // [16:21] is the sub-list for method output_type
	11, // [11:16] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_eolymp_rating_rating_service_proto_init() }
func file_eolymp_rating_rating_service_proto_init() {
	if File_eolymp_rating_rating_service_proto != nil {
		return
	}
	file_eolymp_rating_rating_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_eolymp_rating_rating_service_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*SetRatingInput); i {
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
		file_eolymp_rating_rating_service_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*SetRatingOutput); i {
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
		file_eolymp_rating_rating_service_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*UpdateRatingInput); i {
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
		file_eolymp_rating_rating_service_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*UpdateRatingOutput); i {
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
		file_eolymp_rating_rating_service_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*DeleteRatingInput); i {
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
		file_eolymp_rating_rating_service_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*DeleteRatingOutput); i {
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
		file_eolymp_rating_rating_service_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*DescribeRatingInput); i {
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
		file_eolymp_rating_rating_service_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*DescribeRatingOutput); i {
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
		file_eolymp_rating_rating_service_proto_msgTypes[8].Exporter = func(v any, i int) any {
			switch v := v.(*ListRatingInput); i {
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
		file_eolymp_rating_rating_service_proto_msgTypes[9].Exporter = func(v any, i int) any {
			switch v := v.(*ListRatingOutput); i {
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
		file_eolymp_rating_rating_service_proto_msgTypes[10].Exporter = func(v any, i int) any {
			switch v := v.(*ListRatingInput_Filter); i {
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
			RawDescriptor: file_eolymp_rating_rating_service_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_eolymp_rating_rating_service_proto_goTypes,
		DependencyIndexes: file_eolymp_rating_rating_service_proto_depIdxs,
		EnumInfos:         file_eolymp_rating_rating_service_proto_enumTypes,
		MessageInfos:      file_eolymp_rating_rating_service_proto_msgTypes,
	}.Build()
	File_eolymp_rating_rating_service_proto = out.File
	file_eolymp_rating_rating_service_proto_rawDesc = nil
	file_eolymp_rating_rating_service_proto_goTypes = nil
	file_eolymp_rating_rating_service_proto_depIdxs = nil
}