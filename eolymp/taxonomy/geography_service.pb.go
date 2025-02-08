// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v4.24.4
// source: eolymp/taxonomy/geography_service.proto

package taxonomy

import (
	_ "github.com/eolymp/go-sdk/eolymp/annotations"
	wellknown "github.com/eolymp/go-sdk/eolymp/wellknown"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
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

type DescribeCountryInput struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	CountryId     string                 `protobuf:"bytes,1,opt,name=country_id,json=countryId,proto3" json:"country_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DescribeCountryInput) Reset() {
	*x = DescribeCountryInput{}
	mi := &file_eolymp_taxonomy_geography_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DescribeCountryInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescribeCountryInput) ProtoMessage() {}

func (x *DescribeCountryInput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_taxonomy_geography_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DescribeCountryInput.ProtoReflect.Descriptor instead.
func (*DescribeCountryInput) Descriptor() ([]byte, []int) {
	return file_eolymp_taxonomy_geography_service_proto_rawDescGZIP(), []int{0}
}

func (x *DescribeCountryInput) GetCountryId() string {
	if x != nil {
		return x.CountryId
	}
	return ""
}

type DescribeCountryOutput struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Country       *Country               `protobuf:"bytes,1,opt,name=country,proto3" json:"country,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DescribeCountryOutput) Reset() {
	*x = DescribeCountryOutput{}
	mi := &file_eolymp_taxonomy_geography_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DescribeCountryOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescribeCountryOutput) ProtoMessage() {}

func (x *DescribeCountryOutput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_taxonomy_geography_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DescribeCountryOutput.ProtoReflect.Descriptor instead.
func (*DescribeCountryOutput) Descriptor() ([]byte, []int) {
	return file_eolymp_taxonomy_geography_service_proto_rawDescGZIP(), []int{1}
}

func (x *DescribeCountryOutput) GetCountry() *Country {
	if x != nil {
		return x.Country
	}
	return nil
}

type ListCountriesInput struct {
	state  protoimpl.MessageState `protogen:"open.v1"`
	Offset int32                  `protobuf:"varint,10,opt,name=offset,proto3" json:"offset,omitempty"`
	Size   int32                  `protobuf:"varint,11,opt,name=size,proto3" json:"size,omitempty"`
	// data filters
	Filters       *ListCountriesInput_Filter `protobuf:"bytes,40,opt,name=filters,proto3" json:"filters,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListCountriesInput) Reset() {
	*x = ListCountriesInput{}
	mi := &file_eolymp_taxonomy_geography_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListCountriesInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListCountriesInput) ProtoMessage() {}

func (x *ListCountriesInput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_taxonomy_geography_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListCountriesInput.ProtoReflect.Descriptor instead.
func (*ListCountriesInput) Descriptor() ([]byte, []int) {
	return file_eolymp_taxonomy_geography_service_proto_rawDescGZIP(), []int{2}
}

func (x *ListCountriesInput) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *ListCountriesInput) GetSize() int32 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *ListCountriesInput) GetFilters() *ListCountriesInput_Filter {
	if x != nil {
		return x.Filters
	}
	return nil
}

type ListCountriesOutput struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Items         []*Country             `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
	Total         int32                  `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListCountriesOutput) Reset() {
	*x = ListCountriesOutput{}
	mi := &file_eolymp_taxonomy_geography_service_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListCountriesOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListCountriesOutput) ProtoMessage() {}

func (x *ListCountriesOutput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_taxonomy_geography_service_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListCountriesOutput.ProtoReflect.Descriptor instead.
func (*ListCountriesOutput) Descriptor() ([]byte, []int) {
	return file_eolymp_taxonomy_geography_service_proto_rawDescGZIP(), []int{3}
}

func (x *ListCountriesOutput) GetItems() []*Country {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *ListCountriesOutput) GetTotal() int32 {
	if x != nil {
		return x.Total
	}
	return 0
}

type DescribeRegionInput struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	RegionId      string                 `protobuf:"bytes,1,opt,name=region_id,json=regionId,proto3" json:"region_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DescribeRegionInput) Reset() {
	*x = DescribeRegionInput{}
	mi := &file_eolymp_taxonomy_geography_service_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DescribeRegionInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescribeRegionInput) ProtoMessage() {}

func (x *DescribeRegionInput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_taxonomy_geography_service_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DescribeRegionInput.ProtoReflect.Descriptor instead.
func (*DescribeRegionInput) Descriptor() ([]byte, []int) {
	return file_eolymp_taxonomy_geography_service_proto_rawDescGZIP(), []int{4}
}

func (x *DescribeRegionInput) GetRegionId() string {
	if x != nil {
		return x.RegionId
	}
	return ""
}

type DescribeRegionOutput struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Region        *Region                `protobuf:"bytes,1,opt,name=region,proto3" json:"region,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DescribeRegionOutput) Reset() {
	*x = DescribeRegionOutput{}
	mi := &file_eolymp_taxonomy_geography_service_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DescribeRegionOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescribeRegionOutput) ProtoMessage() {}

func (x *DescribeRegionOutput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_taxonomy_geography_service_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DescribeRegionOutput.ProtoReflect.Descriptor instead.
func (*DescribeRegionOutput) Descriptor() ([]byte, []int) {
	return file_eolymp_taxonomy_geography_service_proto_rawDescGZIP(), []int{5}
}

func (x *DescribeRegionOutput) GetRegion() *Region {
	if x != nil {
		return x.Region
	}
	return nil
}

type ListRegionsInput struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// pagination
	Offset int32 `protobuf:"varint,10,opt,name=offset,proto3" json:"offset,omitempty"`
	Size   int32 `protobuf:"varint,11,opt,name=size,proto3" json:"size,omitempty"`
	// data filters
	Filters       *ListRegionsInput_Filter `protobuf:"bytes,40,opt,name=filters,proto3" json:"filters,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListRegionsInput) Reset() {
	*x = ListRegionsInput{}
	mi := &file_eolymp_taxonomy_geography_service_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListRegionsInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListRegionsInput) ProtoMessage() {}

func (x *ListRegionsInput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_taxonomy_geography_service_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListRegionsInput.ProtoReflect.Descriptor instead.
func (*ListRegionsInput) Descriptor() ([]byte, []int) {
	return file_eolymp_taxonomy_geography_service_proto_rawDescGZIP(), []int{6}
}

func (x *ListRegionsInput) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *ListRegionsInput) GetSize() int32 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *ListRegionsInput) GetFilters() *ListRegionsInput_Filter {
	if x != nil {
		return x.Filters
	}
	return nil
}

type ListRegionsOutput struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Items         []*Region              `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
	Total         int32                  `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListRegionsOutput) Reset() {
	*x = ListRegionsOutput{}
	mi := &file_eolymp_taxonomy_geography_service_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListRegionsOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListRegionsOutput) ProtoMessage() {}

func (x *ListRegionsOutput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_taxonomy_geography_service_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListRegionsOutput.ProtoReflect.Descriptor instead.
func (*ListRegionsOutput) Descriptor() ([]byte, []int) {
	return file_eolymp_taxonomy_geography_service_proto_rawDescGZIP(), []int{7}
}

func (x *ListRegionsOutput) GetItems() []*Region {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *ListRegionsOutput) GetTotal() int32 {
	if x != nil {
		return x.Total
	}
	return 0
}

type ListCountriesInput_Filter struct {
	state         protoimpl.MessageState        `protogen:"open.v1"`
	Query         string                        `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
	Id            []*wellknown.ExpressionID     `protobuf:"bytes,2,rep,name=id,proto3" json:"id,omitempty"`
	Name          []*wellknown.ExpressionString `protobuf:"bytes,3,rep,name=name,proto3" json:"name,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListCountriesInput_Filter) Reset() {
	*x = ListCountriesInput_Filter{}
	mi := &file_eolymp_taxonomy_geography_service_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListCountriesInput_Filter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListCountriesInput_Filter) ProtoMessage() {}

func (x *ListCountriesInput_Filter) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_taxonomy_geography_service_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListCountriesInput_Filter.ProtoReflect.Descriptor instead.
func (*ListCountriesInput_Filter) Descriptor() ([]byte, []int) {
	return file_eolymp_taxonomy_geography_service_proto_rawDescGZIP(), []int{2, 0}
}

func (x *ListCountriesInput_Filter) GetQuery() string {
	if x != nil {
		return x.Query
	}
	return ""
}

func (x *ListCountriesInput_Filter) GetId() []*wellknown.ExpressionID {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *ListCountriesInput_Filter) GetName() []*wellknown.ExpressionString {
	if x != nil {
		return x.Name
	}
	return nil
}

type ListRegionsInput_Filter struct {
	state         protoimpl.MessageState        `protogen:"open.v1"`
	Query         string                        `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
	Id            []*wellknown.ExpressionID     `protobuf:"bytes,2,rep,name=id,proto3" json:"id,omitempty"`
	CountryId     []*wellknown.ExpressionID     `protobuf:"bytes,3,rep,name=country_id,json=countryId,proto3" json:"country_id,omitempty"`
	Name          []*wellknown.ExpressionString `protobuf:"bytes,4,rep,name=name,proto3" json:"name,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListRegionsInput_Filter) Reset() {
	*x = ListRegionsInput_Filter{}
	mi := &file_eolymp_taxonomy_geography_service_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListRegionsInput_Filter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListRegionsInput_Filter) ProtoMessage() {}

func (x *ListRegionsInput_Filter) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_taxonomy_geography_service_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListRegionsInput_Filter.ProtoReflect.Descriptor instead.
func (*ListRegionsInput_Filter) Descriptor() ([]byte, []int) {
	return file_eolymp_taxonomy_geography_service_proto_rawDescGZIP(), []int{6, 0}
}

func (x *ListRegionsInput_Filter) GetQuery() string {
	if x != nil {
		return x.Query
	}
	return ""
}

func (x *ListRegionsInput_Filter) GetId() []*wellknown.ExpressionID {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *ListRegionsInput_Filter) GetCountryId() []*wellknown.ExpressionID {
	if x != nil {
		return x.CountryId
	}
	return nil
}

func (x *ListRegionsInput_Filter) GetName() []*wellknown.ExpressionString {
	if x != nil {
		return x.Name
	}
	return nil
}

var File_eolymp_taxonomy_geography_service_proto protoreflect.FileDescriptor

var file_eolymp_taxonomy_geography_service_proto_rawDesc = string([]byte{
	0x0a, 0x27, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x74, 0x61, 0x78, 0x6f, 0x6e, 0x6f, 0x6d,
	0x79, 0x2f, 0x67, 0x65, 0x6f, 0x67, 0x72, 0x61, 0x70, 0x68, 0x79, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x65, 0x6f, 0x6c, 0x79, 0x6d,
	0x70, 0x2e, 0x74, 0x61, 0x78, 0x6f, 0x6e, 0x6f, 0x6d, 0x79, 0x1a, 0x1d, 0x65, 0x6f, 0x6c, 0x79,
	0x6d, 0x70, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x68,
	0x74, 0x74, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x22, 0x65, 0x6f, 0x6c, 0x79, 0x6d,
	0x70, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x72, 0x61,
	0x74, 0x65, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x27, 0x65,
	0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x74, 0x61, 0x78, 0x6f, 0x6e, 0x6f, 0x6d, 0x79, 0x2f, 0x67,
	0x65, 0x6f, 0x67, 0x72, 0x61, 0x70, 0x68, 0x79, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x26, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x74,
	0x61, 0x78, 0x6f, 0x6e, 0x6f, 0x6d, 0x79, 0x2f, 0x67, 0x65, 0x6f, 0x67, 0x72, 0x61, 0x70, 0x68,
	0x79, 0x5f, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x21,
	0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x77, 0x65, 0x6c, 0x6c, 0x6b, 0x6e, 0x6f, 0x77, 0x6e,
	0x2f, 0x65, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x35, 0x0a, 0x14, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x72, 0x79, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x72, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x49, 0x64, 0x22, 0x4b, 0x0a, 0x15, 0x44, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x62, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x4f, 0x75, 0x74, 0x70, 0x75,
	0x74, 0x12, 0x32, 0x0a, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x18, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x74, 0x61, 0x78, 0x6f,
	0x6e, 0x6f, 0x6d, 0x79, 0x2e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x72, 0x79, 0x22, 0x8f, 0x02, 0x0a, 0x12, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f,
	0x75, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x16, 0x0a, 0x06,
	0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6f, 0x66,
	0x66, 0x73, 0x65, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x0b, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x44, 0x0a, 0x07, 0x66, 0x69, 0x6c, 0x74,
	0x65, 0x72, 0x73, 0x18, 0x28, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x65, 0x6f, 0x6c, 0x79,
	0x6d, 0x70, 0x2e, 0x74, 0x61, 0x78, 0x6f, 0x6e, 0x6f, 0x6d, 0x79, 0x2e, 0x4c, 0x69, 0x73, 0x74,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x2e, 0x46,
	0x69, 0x6c, 0x74, 0x65, 0x72, 0x52, 0x07, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x1a, 0x86,
	0x01, 0x0a, 0x06, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x71, 0x75, 0x65,
	0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x12,
	0x2e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x65, 0x6f,
	0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x77, 0x65, 0x6c, 0x6c, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x2e, 0x45,
	0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x36, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e,
	0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x77, 0x65, 0x6c, 0x6c, 0x6b, 0x6e, 0x6f, 0x77, 0x6e,
	0x2e, 0x45, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x72, 0x69, 0x6e,
	0x67, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x5b, 0x0a, 0x13, 0x4c, 0x69, 0x73, 0x74, 0x43,
	0x6f, 0x75, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x12, 0x2e,
	0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e,
	0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x74, 0x61, 0x78, 0x6f, 0x6e, 0x6f, 0x6d, 0x79, 0x2e,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x14,
	0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x74,
	0x6f, 0x74, 0x61, 0x6c, 0x22, 0x32, 0x0a, 0x13, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65,
	0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x72,
	0x65, 0x67, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x22, 0x47, 0x0a, 0x14, 0x44, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x62, 0x65, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74,
	0x12, 0x2f, 0x0a, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x17, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x74, 0x61, 0x78, 0x6f, 0x6e, 0x6f,
	0x6d, 0x79, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x52, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f,
	0x6e, 0x22, 0xca, 0x02, 0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e,
	0x73, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x73, 0x69,
	0x7a, 0x65, 0x12, 0x42, 0x0a, 0x07, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x18, 0x28, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x74, 0x61, 0x78,
	0x6f, 0x6e, 0x6f, 0x6d, 0x79, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e,
	0x73, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x2e, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x52, 0x07, 0x66,
	0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x1a, 0xc5, 0x01, 0x0a, 0x06, 0x46, 0x69, 0x6c, 0x74, 0x65,
	0x72, 0x12, 0x14, 0x0a, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x12, 0x2e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x77, 0x65, 0x6c,
	0x6c, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x2e, 0x45, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x49, 0x44, 0x52, 0x02, 0x69, 0x64, 0x12, 0x3d, 0x0a, 0x0a, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x72, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x65, 0x6f,
	0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x77, 0x65, 0x6c, 0x6c, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x2e, 0x45,
	0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x52, 0x09, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x72, 0x79, 0x49, 0x64, 0x12, 0x36, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x77, 0x65,
	0x6c, 0x6c, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x2e, 0x45, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x58,
	0x0a, 0x11, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x73, 0x4f, 0x75, 0x74,
	0x70, 0x75, 0x74, 0x12, 0x2d, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x17, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x74, 0x61, 0x78, 0x6f,
	0x6e, 0x6f, 0x6d, 0x79, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x52, 0x05, 0x69, 0x74, 0x65,
	0x6d, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x32, 0xac, 0x04, 0x0a, 0x10, 0x47, 0x65, 0x6f,
	0x67, 0x72, 0x61, 0x70, 0x68, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x7e, 0x0a,
	0x0d, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x12, 0x23,
	0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x74, 0x61, 0x78, 0x6f, 0x6e, 0x6f, 0x6d, 0x79,
	0x2e, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x49, 0x6e,
	0x70, 0x75, 0x74, 0x1a, 0x24, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x74, 0x61, 0x78,
	0x6f, 0x6e, 0x6f, 0x6d, 0x79, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72,
	0x69, 0x65, 0x73, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x22, 0x22, 0xea, 0xe2, 0x0a, 0x0c, 0xf5,
	0xe2, 0x0a, 0x00, 0x00, 0x48, 0x42, 0xf8, 0xe2, 0x0a, 0xf4, 0x03, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x0c, 0x12, 0x0a, 0x2f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x12, 0x91, 0x01,
	0x0a, 0x0f, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72,
	0x79, 0x12, 0x25, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x74, 0x61, 0x78, 0x6f, 0x6e,
	0x6f, 0x6d, 0x79, 0x2e, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x43, 0x6f, 0x75, 0x6e,
	0x74, 0x72, 0x79, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x26, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d,
	0x70, 0x2e, 0x74, 0x61, 0x78, 0x6f, 0x6e, 0x6f, 0x6d, 0x79, 0x2e, 0x44, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x62, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74,
	0x22, 0x2f, 0xea, 0xe2, 0x0a, 0x0c, 0xf5, 0xe2, 0x0a, 0x00, 0x00, 0x48, 0x42, 0xf8, 0xe2, 0x0a,
	0xf4, 0x03, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x19, 0x12, 0x17, 0x2f, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x72, 0x69, 0x65, 0x73, 0x2f, 0x7b, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x5f, 0x69, 0x64,
	0x7d, 0x12, 0x76, 0x0a, 0x0b, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x73,
	0x12, 0x21, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x74, 0x61, 0x78, 0x6f, 0x6e, 0x6f,
	0x6d, 0x79, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x73, 0x49, 0x6e,
	0x70, 0x75, 0x74, 0x1a, 0x22, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x74, 0x61, 0x78,
	0x6f, 0x6e, 0x6f, 0x6d, 0x79, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e,
	0x73, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x22, 0x20, 0xea, 0xe2, 0x0a, 0x0c, 0xf5, 0xe2, 0x0a,
	0x00, 0x00, 0x48, 0x42, 0xf8, 0xe2, 0x0a, 0xf4, 0x03, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0a, 0x12,
	0x08, 0x2f, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x8b, 0x01, 0x0a, 0x0e, 0x44, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x12, 0x24, 0x2e, 0x65,
	0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x74, 0x61, 0x78, 0x6f, 0x6e, 0x6f, 0x6d, 0x79, 0x2e, 0x44,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x70,
	0x75, 0x74, 0x1a, 0x25, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x74, 0x61, 0x78, 0x6f,
	0x6e, 0x6f, 0x6d, 0x79, 0x2e, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x52, 0x65, 0x67,
	0x69, 0x6f, 0x6e, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x22, 0x2c, 0xea, 0xe2, 0x0a, 0x0c, 0xf5,
	0xe2, 0x0a, 0x00, 0x00, 0x48, 0x42, 0xf8, 0xe2, 0x0a, 0xf4, 0x03, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x16, 0x12, 0x14, 0x2f, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x7b, 0x72, 0x65, 0x67,
	0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x7d, 0x42, 0x33, 0x5a, 0x31, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x67, 0x6f, 0x2d,
	0x73, 0x64, 0x6b, 0x2f, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x74, 0x61, 0x78, 0x6f, 0x6e,
	0x6f, 0x6d, 0x79, 0x3b, 0x74, 0x61, 0x78, 0x6f, 0x6e, 0x6f, 0x6d, 0x79, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_eolymp_taxonomy_geography_service_proto_rawDescOnce sync.Once
	file_eolymp_taxonomy_geography_service_proto_rawDescData []byte
)

func file_eolymp_taxonomy_geography_service_proto_rawDescGZIP() []byte {
	file_eolymp_taxonomy_geography_service_proto_rawDescOnce.Do(func() {
		file_eolymp_taxonomy_geography_service_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_eolymp_taxonomy_geography_service_proto_rawDesc), len(file_eolymp_taxonomy_geography_service_proto_rawDesc)))
	})
	return file_eolymp_taxonomy_geography_service_proto_rawDescData
}

var file_eolymp_taxonomy_geography_service_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_eolymp_taxonomy_geography_service_proto_goTypes = []any{
	(*DescribeCountryInput)(nil),       // 0: eolymp.taxonomy.DescribeCountryInput
	(*DescribeCountryOutput)(nil),      // 1: eolymp.taxonomy.DescribeCountryOutput
	(*ListCountriesInput)(nil),         // 2: eolymp.taxonomy.ListCountriesInput
	(*ListCountriesOutput)(nil),        // 3: eolymp.taxonomy.ListCountriesOutput
	(*DescribeRegionInput)(nil),        // 4: eolymp.taxonomy.DescribeRegionInput
	(*DescribeRegionOutput)(nil),       // 5: eolymp.taxonomy.DescribeRegionOutput
	(*ListRegionsInput)(nil),           // 6: eolymp.taxonomy.ListRegionsInput
	(*ListRegionsOutput)(nil),          // 7: eolymp.taxonomy.ListRegionsOutput
	(*ListCountriesInput_Filter)(nil),  // 8: eolymp.taxonomy.ListCountriesInput.Filter
	(*ListRegionsInput_Filter)(nil),    // 9: eolymp.taxonomy.ListRegionsInput.Filter
	(*Country)(nil),                    // 10: eolymp.taxonomy.Country
	(*Region)(nil),                     // 11: eolymp.taxonomy.Region
	(*wellknown.ExpressionID)(nil),     // 12: eolymp.wellknown.ExpressionID
	(*wellknown.ExpressionString)(nil), // 13: eolymp.wellknown.ExpressionString
}
var file_eolymp_taxonomy_geography_service_proto_depIdxs = []int32{
	10, // 0: eolymp.taxonomy.DescribeCountryOutput.country:type_name -> eolymp.taxonomy.Country
	8,  // 1: eolymp.taxonomy.ListCountriesInput.filters:type_name -> eolymp.taxonomy.ListCountriesInput.Filter
	10, // 2: eolymp.taxonomy.ListCountriesOutput.items:type_name -> eolymp.taxonomy.Country
	11, // 3: eolymp.taxonomy.DescribeRegionOutput.region:type_name -> eolymp.taxonomy.Region
	9,  // 4: eolymp.taxonomy.ListRegionsInput.filters:type_name -> eolymp.taxonomy.ListRegionsInput.Filter
	11, // 5: eolymp.taxonomy.ListRegionsOutput.items:type_name -> eolymp.taxonomy.Region
	12, // 6: eolymp.taxonomy.ListCountriesInput.Filter.id:type_name -> eolymp.wellknown.ExpressionID
	13, // 7: eolymp.taxonomy.ListCountriesInput.Filter.name:type_name -> eolymp.wellknown.ExpressionString
	12, // 8: eolymp.taxonomy.ListRegionsInput.Filter.id:type_name -> eolymp.wellknown.ExpressionID
	12, // 9: eolymp.taxonomy.ListRegionsInput.Filter.country_id:type_name -> eolymp.wellknown.ExpressionID
	13, // 10: eolymp.taxonomy.ListRegionsInput.Filter.name:type_name -> eolymp.wellknown.ExpressionString
	2,  // 11: eolymp.taxonomy.GeographyService.ListCountries:input_type -> eolymp.taxonomy.ListCountriesInput
	0,  // 12: eolymp.taxonomy.GeographyService.DescribeCountry:input_type -> eolymp.taxonomy.DescribeCountryInput
	6,  // 13: eolymp.taxonomy.GeographyService.ListRegions:input_type -> eolymp.taxonomy.ListRegionsInput
	4,  // 14: eolymp.taxonomy.GeographyService.DescribeRegion:input_type -> eolymp.taxonomy.DescribeRegionInput
	3,  // 15: eolymp.taxonomy.GeographyService.ListCountries:output_type -> eolymp.taxonomy.ListCountriesOutput
	1,  // 16: eolymp.taxonomy.GeographyService.DescribeCountry:output_type -> eolymp.taxonomy.DescribeCountryOutput
	7,  // 17: eolymp.taxonomy.GeographyService.ListRegions:output_type -> eolymp.taxonomy.ListRegionsOutput
	5,  // 18: eolymp.taxonomy.GeographyService.DescribeRegion:output_type -> eolymp.taxonomy.DescribeRegionOutput
	15, // [15:19] is the sub-list for method output_type
	11, // [11:15] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_eolymp_taxonomy_geography_service_proto_init() }
func file_eolymp_taxonomy_geography_service_proto_init() {
	if File_eolymp_taxonomy_geography_service_proto != nil {
		return
	}
	file_eolymp_taxonomy_geography_country_proto_init()
	file_eolymp_taxonomy_geography_region_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_eolymp_taxonomy_geography_service_proto_rawDesc), len(file_eolymp_taxonomy_geography_service_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_eolymp_taxonomy_geography_service_proto_goTypes,
		DependencyIndexes: file_eolymp_taxonomy_geography_service_proto_depIdxs,
		MessageInfos:      file_eolymp_taxonomy_geography_service_proto_msgTypes,
	}.Build()
	File_eolymp_taxonomy_geography_service_proto = out.File
	file_eolymp_taxonomy_geography_service_proto_goTypes = nil
	file_eolymp_taxonomy_geography_service_proto_depIdxs = nil
}
