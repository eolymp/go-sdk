// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.18.1
// source: eolymp/community/tier_service.proto

package community

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

type ListTiersInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Offset   int32  `protobuf:"varint,10,opt,name=offset,proto3" json:"offset,omitempty"`
	Size     int32  `protobuf:"varint,11,opt,name=size,proto3" json:"size,omitempty"`
	Render   bool   `protobuf:"varint,100,opt,name=render,proto3" json:"render,omitempty"`    // render description to ecm
	Locale   string `protobuf:"bytes,101,opt,name=locale,proto3" json:"locale,omitempty"`     // optionally, request name and description in particular language
	Currency string `protobuf:"bytes,102,opt,name=currency,proto3" json:"currency,omitempty"` // optionally, request prices in particular currency
}

func (x *ListTiersInput) Reset() {
	*x = ListTiersInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_community_tier_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListTiersInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListTiersInput) ProtoMessage() {}

func (x *ListTiersInput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_community_tier_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListTiersInput.ProtoReflect.Descriptor instead.
func (*ListTiersInput) Descriptor() ([]byte, []int) {
	return file_eolymp_community_tier_service_proto_rawDescGZIP(), []int{0}
}

func (x *ListTiersInput) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *ListTiersInput) GetSize() int32 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *ListTiersInput) GetRender() bool {
	if x != nil {
		return x.Render
	}
	return false
}

func (x *ListTiersInput) GetLocale() string {
	if x != nil {
		return x.Locale
	}
	return ""
}

func (x *ListTiersInput) GetCurrency() string {
	if x != nil {
		return x.Currency
	}
	return ""
}

type ListTiersOutput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Total int32   `protobuf:"varint,1,opt,name=total,proto3" json:"total,omitempty"`
	Items []*Tier `protobuf:"bytes,2,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *ListTiersOutput) Reset() {
	*x = ListTiersOutput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_community_tier_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListTiersOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListTiersOutput) ProtoMessage() {}

func (x *ListTiersOutput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_community_tier_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListTiersOutput.ProtoReflect.Descriptor instead.
func (*ListTiersOutput) Descriptor() ([]byte, []int) {
	return file_eolymp_community_tier_service_proto_rawDescGZIP(), []int{1}
}

func (x *ListTiersOutput) GetTotal() int32 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *ListTiersOutput) GetItems() []*Tier {
	if x != nil {
		return x.Items
	}
	return nil
}

type DescribeTierInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TierId   string `protobuf:"bytes,1,opt,name=tier_id,json=tierId,proto3" json:"tier_id,omitempty"`
	Render   bool   `protobuf:"varint,100,opt,name=render,proto3" json:"render,omitempty"`    // render description to ecm
	Locale   string `protobuf:"bytes,101,opt,name=locale,proto3" json:"locale,omitempty"`     // optionally, request name and description in particular language
	Currency string `protobuf:"bytes,102,opt,name=currency,proto3" json:"currency,omitempty"` // optionally, request prices in particular currency
}

func (x *DescribeTierInput) Reset() {
	*x = DescribeTierInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_community_tier_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DescribeTierInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescribeTierInput) ProtoMessage() {}

func (x *DescribeTierInput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_community_tier_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DescribeTierInput.ProtoReflect.Descriptor instead.
func (*DescribeTierInput) Descriptor() ([]byte, []int) {
	return file_eolymp_community_tier_service_proto_rawDescGZIP(), []int{2}
}

func (x *DescribeTierInput) GetTierId() string {
	if x != nil {
		return x.TierId
	}
	return ""
}

func (x *DescribeTierInput) GetRender() bool {
	if x != nil {
		return x.Render
	}
	return false
}

func (x *DescribeTierInput) GetLocale() string {
	if x != nil {
		return x.Locale
	}
	return ""
}

func (x *DescribeTierInput) GetCurrency() string {
	if x != nil {
		return x.Currency
	}
	return ""
}

type DescribeTierOutput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tier *Tier `protobuf:"bytes,1,opt,name=tier,proto3" json:"tier,omitempty"`
}

func (x *DescribeTierOutput) Reset() {
	*x = DescribeTierOutput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_community_tier_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DescribeTierOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescribeTierOutput) ProtoMessage() {}

func (x *DescribeTierOutput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_community_tier_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DescribeTierOutput.ProtoReflect.Descriptor instead.
func (*DescribeTierOutput) Descriptor() ([]byte, []int) {
	return file_eolymp_community_tier_service_proto_rawDescGZIP(), []int{3}
}

func (x *DescribeTierOutput) GetTier() *Tier {
	if x != nil {
		return x.Tier
	}
	return nil
}

type ListCurrenciesInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListCurrenciesInput) Reset() {
	*x = ListCurrenciesInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_community_tier_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListCurrenciesInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListCurrenciesInput) ProtoMessage() {}

func (x *ListCurrenciesInput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_community_tier_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListCurrenciesInput.ProtoReflect.Descriptor instead.
func (*ListCurrenciesInput) Descriptor() ([]byte, []int) {
	return file_eolymp_community_tier_service_proto_rawDescGZIP(), []int{4}
}

type ListCurrenciesOutput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items []string `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *ListCurrenciesOutput) Reset() {
	*x = ListCurrenciesOutput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_community_tier_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListCurrenciesOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListCurrenciesOutput) ProtoMessage() {}

func (x *ListCurrenciesOutput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_community_tier_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListCurrenciesOutput.ProtoReflect.Descriptor instead.
func (*ListCurrenciesOutput) Descriptor() ([]byte, []int) {
	return file_eolymp_community_tier_service_proto_rawDescGZIP(), []int{5}
}

func (x *ListCurrenciesOutput) GetItems() []string {
	if x != nil {
		return x.Items
	}
	return nil
}

var File_eolymp_community_tier_service_proto protoreflect.FileDescriptor

var file_eolymp_community_tier_service_proto_rawDesc = []byte{
	0x0a, 0x23, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69,
	0x74, 0x79, 0x2f, 0x74, 0x69, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x1a, 0x1d, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x68, 0x74, 0x74, 0x70,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x22, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x61,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x72, 0x61, 0x74, 0x65, 0x6c,
	0x69, 0x6d, 0x69, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x65, 0x6f, 0x6c, 0x79,
	0x6d, 0x70, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73,
	0x63, 0x6f, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x65, 0x6f, 0x6c, 0x79,
	0x6d, 0x70, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x2f, 0x74, 0x69, 0x65,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x88, 0x01, 0x0a, 0x0e, 0x4c, 0x69, 0x73, 0x74,
	0x54, 0x69, 0x65, 0x72, 0x73, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66,
	0x66, 0x73, 0x65, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73,
	0x65, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x6e, 0x64, 0x65, 0x72,
	0x18, 0x64, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x72, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x16,
	0x0a, 0x06, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x65, 0x18, 0x65, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e,
	0x63, 0x79, 0x18, 0x66, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e,
	0x63, 0x79, 0x22, 0x55, 0x0a, 0x0f, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x69, 0x65, 0x72, 0x73, 0x4f,
	0x75, 0x74, 0x70, 0x75, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x2c, 0x0a, 0x05, 0x69,
	0x74, 0x65, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x65, 0x6f, 0x6c,
	0x79, 0x6d, 0x70, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x2e, 0x54, 0x69,
	0x65, 0x72, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x22, 0x78, 0x0a, 0x11, 0x44, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x62, 0x65, 0x54, 0x69, 0x65, 0x72, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x17,
	0x0a, 0x07, 0x74, 0x69, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x74, 0x69, 0x65, 0x72, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x6e, 0x64, 0x65,
	0x72, 0x18, 0x64, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x72, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12,
	0x16, 0x0a, 0x06, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x65, 0x18, 0x65, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x75, 0x72, 0x72, 0x65,
	0x6e, 0x63, 0x79, 0x18, 0x66, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x75, 0x72, 0x72, 0x65,
	0x6e, 0x63, 0x79, 0x22, 0x40, 0x0a, 0x12, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x54,
	0x69, 0x65, 0x72, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x12, 0x2a, 0x0a, 0x04, 0x74, 0x69, 0x65,
	0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x2e, 0x54, 0x69, 0x65, 0x72, 0x52,
	0x04, 0x74, 0x69, 0x65, 0x72, 0x22, 0x15, 0x0a, 0x13, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x75, 0x72,
	0x72, 0x65, 0x6e, 0x63, 0x69, 0x65, 0x73, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x22, 0x2c, 0x0a, 0x14,
	0x4c, 0x69, 0x73, 0x74, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x69, 0x65, 0x73, 0x4f, 0x75,
	0x74, 0x70, 0x75, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x32, 0xc5, 0x03, 0x0a, 0x0b, 0x54,
	0x69, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x9d, 0x01, 0x0a, 0x0c, 0x44,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x54, 0x69, 0x65, 0x72, 0x12, 0x23, 0x2e, 0x65, 0x6f,
	0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x2e, 0x44,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x54, 0x69, 0x65, 0x72, 0x49, 0x6e, 0x70, 0x75, 0x74,
	0x1a, 0x24, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e,
	0x69, 0x74, 0x79, 0x2e, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x54, 0x69, 0x65, 0x72,
	0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x22, 0x42, 0xea, 0xe2, 0x0a, 0x0b, 0xf5, 0xe2, 0x0a, 0x00,
	0x00, 0xa0, 0x40, 0xf8, 0xe2, 0x0a, 0x32, 0x82, 0xe3, 0x0a, 0x17, 0x8a, 0xe3, 0x0a, 0x13, 0x63,
	0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x3a, 0x74, 0x69, 0x65, 0x72, 0x3a, 0x72, 0x65,
	0x61, 0x64, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x12, 0x12, 0x10, 0x2f, 0x74, 0x69, 0x65, 0x72, 0x73,
	0x2f, 0x7b, 0x74, 0x69, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x7d, 0x12, 0x8a, 0x01, 0x0a, 0x09, 0x4c,
	0x69, 0x73, 0x74, 0x54, 0x69, 0x65, 0x72, 0x73, 0x12, 0x20, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d,
	0x70, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x2e, 0x4c, 0x69, 0x73, 0x74,
	0x54, 0x69, 0x65, 0x72, 0x73, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x21, 0x2e, 0x65, 0x6f, 0x6c,
	0x79, 0x6d, 0x70, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x2e, 0x4c, 0x69,
	0x73, 0x74, 0x54, 0x69, 0x65, 0x72, 0x73, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x22, 0x38, 0xea,
	0xe2, 0x0a, 0x0b, 0xf5, 0xe2, 0x0a, 0x00, 0x00, 0xa0, 0x40, 0xf8, 0xe2, 0x0a, 0x14, 0x82, 0xe3,
	0x0a, 0x17, 0x8a, 0xe3, 0x0a, 0x13, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x3a,
	0x74, 0x69, 0x65, 0x72, 0x3a, 0x72, 0x65, 0x61, 0x64, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x08, 0x12,
	0x06, 0x2f, 0x74, 0x69, 0x65, 0x72, 0x73, 0x12, 0x88, 0x01, 0x0a, 0x0e, 0x4c, 0x69, 0x73, 0x74,
	0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x69, 0x65, 0x73, 0x12, 0x25, 0x2e, 0x65, 0x6f, 0x6c,
	0x79, 0x6d, 0x70, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x2e, 0x4c, 0x69,
	0x73, 0x74, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x69, 0x65, 0x73, 0x49, 0x6e, 0x70, 0x75,
	0x74, 0x1a, 0x26, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x75,
	0x6e, 0x69, 0x74, 0x79, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63,
	0x69, 0x65, 0x73, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x22, 0x27, 0xea, 0xe2, 0x0a, 0x0b, 0xf5,
	0xe2, 0x0a, 0x00, 0x00, 0xa0, 0x40, 0xf8, 0xe2, 0x0a, 0x14, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x12,
	0x12, 0x10, 0x2f, 0x74, 0x69, 0x65, 0x72, 0x2d, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x69,
	0x65, 0x73, 0x42, 0x35, 0x5a, 0x33, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x67, 0x6f, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x65,
	0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x3b,
	0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_eolymp_community_tier_service_proto_rawDescOnce sync.Once
	file_eolymp_community_tier_service_proto_rawDescData = file_eolymp_community_tier_service_proto_rawDesc
)

func file_eolymp_community_tier_service_proto_rawDescGZIP() []byte {
	file_eolymp_community_tier_service_proto_rawDescOnce.Do(func() {
		file_eolymp_community_tier_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_eolymp_community_tier_service_proto_rawDescData)
	})
	return file_eolymp_community_tier_service_proto_rawDescData
}

var file_eolymp_community_tier_service_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_eolymp_community_tier_service_proto_goTypes = []interface{}{
	(*ListTiersInput)(nil),       // 0: eolymp.community.ListTiersInput
	(*ListTiersOutput)(nil),      // 1: eolymp.community.ListTiersOutput
	(*DescribeTierInput)(nil),    // 2: eolymp.community.DescribeTierInput
	(*DescribeTierOutput)(nil),   // 3: eolymp.community.DescribeTierOutput
	(*ListCurrenciesInput)(nil),  // 4: eolymp.community.ListCurrenciesInput
	(*ListCurrenciesOutput)(nil), // 5: eolymp.community.ListCurrenciesOutput
	(*Tier)(nil),                 // 6: eolymp.community.Tier
}
var file_eolymp_community_tier_service_proto_depIdxs = []int32{
	6, // 0: eolymp.community.ListTiersOutput.items:type_name -> eolymp.community.Tier
	6, // 1: eolymp.community.DescribeTierOutput.tier:type_name -> eolymp.community.Tier
	2, // 2: eolymp.community.TierService.DescribeTier:input_type -> eolymp.community.DescribeTierInput
	0, // 3: eolymp.community.TierService.ListTiers:input_type -> eolymp.community.ListTiersInput
	4, // 4: eolymp.community.TierService.ListCurrencies:input_type -> eolymp.community.ListCurrenciesInput
	3, // 5: eolymp.community.TierService.DescribeTier:output_type -> eolymp.community.DescribeTierOutput
	1, // 6: eolymp.community.TierService.ListTiers:output_type -> eolymp.community.ListTiersOutput
	5, // 7: eolymp.community.TierService.ListCurrencies:output_type -> eolymp.community.ListCurrenciesOutput
	5, // [5:8] is the sub-list for method output_type
	2, // [2:5] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_eolymp_community_tier_service_proto_init() }
func file_eolymp_community_tier_service_proto_init() {
	if File_eolymp_community_tier_service_proto != nil {
		return
	}
	file_eolymp_community_tier_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_eolymp_community_tier_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListTiersInput); i {
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
		file_eolymp_community_tier_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListTiersOutput); i {
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
		file_eolymp_community_tier_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DescribeTierInput); i {
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
		file_eolymp_community_tier_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DescribeTierOutput); i {
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
		file_eolymp_community_tier_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListCurrenciesInput); i {
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
		file_eolymp_community_tier_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListCurrenciesOutput); i {
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
			RawDescriptor: file_eolymp_community_tier_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_eolymp_community_tier_service_proto_goTypes,
		DependencyIndexes: file_eolymp_community_tier_service_proto_depIdxs,
		MessageInfos:      file_eolymp_community_tier_service_proto_msgTypes,
	}.Build()
	File_eolymp_community_tier_service_proto = out.File
	file_eolymp_community_tier_service_proto_rawDesc = nil
	file_eolymp_community_tier_service_proto_goTypes = nil
	file_eolymp_community_tier_service_proto_depIdxs = nil
}
