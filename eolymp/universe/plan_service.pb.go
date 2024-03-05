// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v3.18.1
// source: eolymp/universe/plan_service.proto

package universe

import (
	_ "github.com/eolymp/go-sdk/eolymp/annotations"
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

type DescribePlanInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PlanId   string       `protobuf:"bytes,1,opt,name=plan_id,json=planId,proto3" json:"plan_id,omitempty"`
	Locale   string       `protobuf:"bytes,101,opt,name=locale,proto3" json:"locale,omitempty"`     // optionally, request name and description in particular locale
	Currency string       `protobuf:"bytes,102,opt,name=currency,proto3" json:"currency,omitempty"` // optionally, request variants in particular currency
	Extra    []Plan_Extra `protobuf:"varint,1123,rep,packed,name=extra,proto3,enum=eolymp.universe.Plan_Extra" json:"extra,omitempty"`
}

func (x *DescribePlanInput) Reset() {
	*x = DescribePlanInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_universe_plan_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DescribePlanInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescribePlanInput) ProtoMessage() {}

func (x *DescribePlanInput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_universe_plan_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DescribePlanInput.ProtoReflect.Descriptor instead.
func (*DescribePlanInput) Descriptor() ([]byte, []int) {
	return file_eolymp_universe_plan_service_proto_rawDescGZIP(), []int{0}
}

func (x *DescribePlanInput) GetPlanId() string {
	if x != nil {
		return x.PlanId
	}
	return ""
}

func (x *DescribePlanInput) GetLocale() string {
	if x != nil {
		return x.Locale
	}
	return ""
}

func (x *DescribePlanInput) GetCurrency() string {
	if x != nil {
		return x.Currency
	}
	return ""
}

func (x *DescribePlanInput) GetExtra() []Plan_Extra {
	if x != nil {
		return x.Extra
	}
	return nil
}

type DescribePlanOutput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Plan *Plan `protobuf:"bytes,1,opt,name=plan,proto3" json:"plan,omitempty"`
}

func (x *DescribePlanOutput) Reset() {
	*x = DescribePlanOutput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_universe_plan_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DescribePlanOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescribePlanOutput) ProtoMessage() {}

func (x *DescribePlanOutput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_universe_plan_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DescribePlanOutput.ProtoReflect.Descriptor instead.
func (*DescribePlanOutput) Descriptor() ([]byte, []int) {
	return file_eolymp_universe_plan_service_proto_rawDescGZIP(), []int{1}
}

func (x *DescribePlanOutput) GetPlan() *Plan {
	if x != nil {
		return x.Plan
	}
	return nil
}

type ListPlansInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Offset   int32        `protobuf:"varint,10,opt,name=offset,proto3" json:"offset,omitempty"`
	Size     int32        `protobuf:"varint,11,opt,name=size,proto3" json:"size,omitempty"`
	Locale   string       `protobuf:"bytes,101,opt,name=locale,proto3" json:"locale,omitempty"`     // optionally, request name and description in particular language
	Currency string       `protobuf:"bytes,102,opt,name=currency,proto3" json:"currency,omitempty"` // optionally, request prices in particular currency
	Extra    []Plan_Extra `protobuf:"varint,1123,rep,packed,name=extra,proto3,enum=eolymp.universe.Plan_Extra" json:"extra,omitempty"`
}

func (x *ListPlansInput) Reset() {
	*x = ListPlansInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_universe_plan_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListPlansInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListPlansInput) ProtoMessage() {}

func (x *ListPlansInput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_universe_plan_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListPlansInput.ProtoReflect.Descriptor instead.
func (*ListPlansInput) Descriptor() ([]byte, []int) {
	return file_eolymp_universe_plan_service_proto_rawDescGZIP(), []int{2}
}

func (x *ListPlansInput) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *ListPlansInput) GetSize() int32 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *ListPlansInput) GetLocale() string {
	if x != nil {
		return x.Locale
	}
	return ""
}

func (x *ListPlansInput) GetCurrency() string {
	if x != nil {
		return x.Currency
	}
	return ""
}

func (x *ListPlansInput) GetExtra() []Plan_Extra {
	if x != nil {
		return x.Extra
	}
	return nil
}

type ListPlansOutput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Total int32   `protobuf:"varint,1,opt,name=total,proto3" json:"total,omitempty"`
	Items []*Plan `protobuf:"bytes,2,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *ListPlansOutput) Reset() {
	*x = ListPlansOutput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_universe_plan_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListPlansOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListPlansOutput) ProtoMessage() {}

func (x *ListPlansOutput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_universe_plan_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListPlansOutput.ProtoReflect.Descriptor instead.
func (*ListPlansOutput) Descriptor() ([]byte, []int) {
	return file_eolymp_universe_plan_service_proto_rawDescGZIP(), []int{3}
}

func (x *ListPlansOutput) GetTotal() int32 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *ListPlansOutput) GetItems() []*Plan {
	if x != nil {
		return x.Items
	}
	return nil
}

var File_eolymp_universe_plan_service_proto protoreflect.FileDescriptor

var file_eolymp_universe_plan_service_proto_rawDesc = []byte{
	0x0a, 0x22, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x75, 0x6e, 0x69, 0x76, 0x65, 0x72, 0x73,
	0x65, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x75, 0x6e, 0x69,
	0x76, 0x65, 0x72, 0x73, 0x65, 0x1a, 0x1d, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x61, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x22, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x72, 0x61, 0x74, 0x65, 0x6c, 0x69, 0x6d,
	0x69, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1a, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70,
	0x2f, 0x75, 0x6e, 0x69, 0x76, 0x65, 0x72, 0x73, 0x65, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x94, 0x01, 0x0a, 0x11, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62,
	0x65, 0x50, 0x6c, 0x61, 0x6e, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x70, 0x6c,
	0x61, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x6c, 0x61,
	0x6e, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x65, 0x18, 0x65, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x63,
	0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x66, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63,
	0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x32, 0x0a, 0x05, 0x65, 0x78, 0x74, 0x72, 0x61,
	0x18, 0xe3, 0x08, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x1b, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70,
	0x2e, 0x75, 0x6e, 0x69, 0x76, 0x65, 0x72, 0x73, 0x65, 0x2e, 0x50, 0x6c, 0x61, 0x6e, 0x2e, 0x45,
	0x78, 0x74, 0x72, 0x61, 0x52, 0x05, 0x65, 0x78, 0x74, 0x72, 0x61, 0x22, 0x3f, 0x0a, 0x12, 0x44,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x50, 0x6c, 0x61, 0x6e, 0x4f, 0x75, 0x74, 0x70, 0x75,
	0x74, 0x12, 0x29, 0x0a, 0x04, 0x70, 0x6c, 0x61, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x15, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x75, 0x6e, 0x69, 0x76, 0x65, 0x72, 0x73,
	0x65, 0x2e, 0x50, 0x6c, 0x61, 0x6e, 0x52, 0x04, 0x70, 0x6c, 0x61, 0x6e, 0x22, 0xa4, 0x01, 0x0a,
	0x0e, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x6c, 0x61, 0x6e, 0x73, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12,
	0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18,
	0x0b, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6c,
	0x6f, 0x63, 0x61, 0x6c, 0x65, 0x18, 0x65, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6c, 0x6f, 0x63,
	0x61, 0x6c, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x18,
	0x66, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x12,
	0x32, 0x0a, 0x05, 0x65, 0x78, 0x74, 0x72, 0x61, 0x18, 0xe3, 0x08, 0x20, 0x03, 0x28, 0x0e, 0x32,
	0x1b, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x75, 0x6e, 0x69, 0x76, 0x65, 0x72, 0x73,
	0x65, 0x2e, 0x50, 0x6c, 0x61, 0x6e, 0x2e, 0x45, 0x78, 0x74, 0x72, 0x61, 0x52, 0x05, 0x65, 0x78,
	0x74, 0x72, 0x61, 0x22, 0x54, 0x0a, 0x0f, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x6c, 0x61, 0x6e, 0x73,
	0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x2b, 0x0a, 0x05,
	0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x65, 0x6f,
	0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x75, 0x6e, 0x69, 0x76, 0x65, 0x72, 0x73, 0x65, 0x2e, 0x50, 0x6c,
	0x61, 0x6e, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x32, 0xff, 0x01, 0x0a, 0x0b, 0x50, 0x6c,
	0x61, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x80, 0x01, 0x0a, 0x0c, 0x44, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x50, 0x6c, 0x61, 0x6e, 0x12, 0x22, 0x2e, 0x65, 0x6f, 0x6c,
	0x79, 0x6d, 0x70, 0x2e, 0x75, 0x6e, 0x69, 0x76, 0x65, 0x72, 0x73, 0x65, 0x2e, 0x44, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x62, 0x65, 0x50, 0x6c, 0x61, 0x6e, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x23,
	0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x75, 0x6e, 0x69, 0x76, 0x65, 0x72, 0x73, 0x65,
	0x2e, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x50, 0x6c, 0x61, 0x6e, 0x4f, 0x75, 0x74,
	0x70, 0x75, 0x74, 0x22, 0x27, 0xea, 0xe2, 0x0a, 0x0b, 0xf5, 0xe2, 0x0a, 0x00, 0x00, 0xa0, 0x40,
	0xf8, 0xe2, 0x0a, 0x32, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x12, 0x12, 0x10, 0x2f, 0x70, 0x6c, 0x61,
	0x6e, 0x73, 0x2f, 0x7b, 0x70, 0x6c, 0x61, 0x6e, 0x5f, 0x69, 0x64, 0x7d, 0x12, 0x6d, 0x0a, 0x09,
	0x4c, 0x69, 0x73, 0x74, 0x50, 0x6c, 0x61, 0x6e, 0x73, 0x12, 0x1f, 0x2e, 0x65, 0x6f, 0x6c, 0x79,
	0x6d, 0x70, 0x2e, 0x75, 0x6e, 0x69, 0x76, 0x65, 0x72, 0x73, 0x65, 0x2e, 0x4c, 0x69, 0x73, 0x74,
	0x50, 0x6c, 0x61, 0x6e, 0x73, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x20, 0x2e, 0x65, 0x6f, 0x6c,
	0x79, 0x6d, 0x70, 0x2e, 0x75, 0x6e, 0x69, 0x76, 0x65, 0x72, 0x73, 0x65, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x50, 0x6c, 0x61, 0x6e, 0x73, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x22, 0x1d, 0xea, 0xe2,
	0x0a, 0x0b, 0xf5, 0xe2, 0x0a, 0x00, 0x00, 0xa0, 0x40, 0xf8, 0xe2, 0x0a, 0x14, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x08, 0x12, 0x06, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x73, 0x42, 0x33, 0x5a, 0x31, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70,
	0x2f, 0x67, 0x6f, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x75,
	0x6e, 0x69, 0x76, 0x65, 0x72, 0x73, 0x65, 0x3b, 0x75, 0x6e, 0x69, 0x76, 0x65, 0x72, 0x73, 0x65,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_eolymp_universe_plan_service_proto_rawDescOnce sync.Once
	file_eolymp_universe_plan_service_proto_rawDescData = file_eolymp_universe_plan_service_proto_rawDesc
)

func file_eolymp_universe_plan_service_proto_rawDescGZIP() []byte {
	file_eolymp_universe_plan_service_proto_rawDescOnce.Do(func() {
		file_eolymp_universe_plan_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_eolymp_universe_plan_service_proto_rawDescData)
	})
	return file_eolymp_universe_plan_service_proto_rawDescData
}

var file_eolymp_universe_plan_service_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_eolymp_universe_plan_service_proto_goTypes = []interface{}{
	(*DescribePlanInput)(nil),  // 0: eolymp.universe.DescribePlanInput
	(*DescribePlanOutput)(nil), // 1: eolymp.universe.DescribePlanOutput
	(*ListPlansInput)(nil),     // 2: eolymp.universe.ListPlansInput
	(*ListPlansOutput)(nil),    // 3: eolymp.universe.ListPlansOutput
	(Plan_Extra)(0),            // 4: eolymp.universe.Plan.Extra
	(*Plan)(nil),               // 5: eolymp.universe.Plan
}
var file_eolymp_universe_plan_service_proto_depIdxs = []int32{
	4, // 0: eolymp.universe.DescribePlanInput.extra:type_name -> eolymp.universe.Plan.Extra
	5, // 1: eolymp.universe.DescribePlanOutput.plan:type_name -> eolymp.universe.Plan
	4, // 2: eolymp.universe.ListPlansInput.extra:type_name -> eolymp.universe.Plan.Extra
	5, // 3: eolymp.universe.ListPlansOutput.items:type_name -> eolymp.universe.Plan
	0, // 4: eolymp.universe.PlanService.DescribePlan:input_type -> eolymp.universe.DescribePlanInput
	2, // 5: eolymp.universe.PlanService.ListPlans:input_type -> eolymp.universe.ListPlansInput
	1, // 6: eolymp.universe.PlanService.DescribePlan:output_type -> eolymp.universe.DescribePlanOutput
	3, // 7: eolymp.universe.PlanService.ListPlans:output_type -> eolymp.universe.ListPlansOutput
	6, // [6:8] is the sub-list for method output_type
	4, // [4:6] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_eolymp_universe_plan_service_proto_init() }
func file_eolymp_universe_plan_service_proto_init() {
	if File_eolymp_universe_plan_service_proto != nil {
		return
	}
	file_eolymp_universe_plan_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_eolymp_universe_plan_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DescribePlanInput); i {
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
		file_eolymp_universe_plan_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DescribePlanOutput); i {
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
		file_eolymp_universe_plan_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListPlansInput); i {
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
		file_eolymp_universe_plan_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListPlansOutput); i {
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
			RawDescriptor: file_eolymp_universe_plan_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_eolymp_universe_plan_service_proto_goTypes,
		DependencyIndexes: file_eolymp_universe_plan_service_proto_depIdxs,
		MessageInfos:      file_eolymp_universe_plan_service_proto_msgTypes,
	}.Build()
	File_eolymp_universe_plan_service_proto = out.File
	file_eolymp_universe_plan_service_proto_rawDesc = nil
	file_eolymp_universe_plan_service_proto_goTypes = nil
	file_eolymp_universe_plan_service_proto_depIdxs = nil
}
