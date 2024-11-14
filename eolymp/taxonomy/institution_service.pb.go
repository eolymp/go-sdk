// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v3.21.12
// source: eolymp/taxonomy/institution_service.proto

package taxonomy

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

type DescribeInstitutionInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	InstitutionId string `protobuf:"bytes,1,opt,name=institution_id,json=institutionId,proto3" json:"institution_id,omitempty"`
}

func (x *DescribeInstitutionInput) Reset() {
	*x = DescribeInstitutionInput{}
	mi := &file_eolymp_taxonomy_institution_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DescribeInstitutionInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescribeInstitutionInput) ProtoMessage() {}

func (x *DescribeInstitutionInput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_taxonomy_institution_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DescribeInstitutionInput.ProtoReflect.Descriptor instead.
func (*DescribeInstitutionInput) Descriptor() ([]byte, []int) {
	return file_eolymp_taxonomy_institution_service_proto_rawDescGZIP(), []int{0}
}

func (x *DescribeInstitutionInput) GetInstitutionId() string {
	if x != nil {
		return x.InstitutionId
	}
	return ""
}

type DescribeInstitutionOutput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Institution *Institution `protobuf:"bytes,1,opt,name=institution,proto3" json:"institution,omitempty"`
}

func (x *DescribeInstitutionOutput) Reset() {
	*x = DescribeInstitutionOutput{}
	mi := &file_eolymp_taxonomy_institution_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DescribeInstitutionOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescribeInstitutionOutput) ProtoMessage() {}

func (x *DescribeInstitutionOutput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_taxonomy_institution_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DescribeInstitutionOutput.ProtoReflect.Descriptor instead.
func (*DescribeInstitutionOutput) Descriptor() ([]byte, []int) {
	return file_eolymp_taxonomy_institution_service_proto_rawDescGZIP(), []int{1}
}

func (x *DescribeInstitutionOutput) GetInstitution() *Institution {
	if x != nil {
		return x.Institution
	}
	return nil
}

type ListInstitutionsInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// pagination
	Offset int32 `protobuf:"varint,10,opt,name=offset,proto3" json:"offset,omitempty"`
	Size   int32 `protobuf:"varint,11,opt,name=size,proto3" json:"size,omitempty"`
	// data filters
	Filters *ListInstitutionsInput_Filter `protobuf:"bytes,40,opt,name=filters,proto3" json:"filters,omitempty"`
}

func (x *ListInstitutionsInput) Reset() {
	*x = ListInstitutionsInput{}
	mi := &file_eolymp_taxonomy_institution_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListInstitutionsInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListInstitutionsInput) ProtoMessage() {}

func (x *ListInstitutionsInput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_taxonomy_institution_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListInstitutionsInput.ProtoReflect.Descriptor instead.
func (*ListInstitutionsInput) Descriptor() ([]byte, []int) {
	return file_eolymp_taxonomy_institution_service_proto_rawDescGZIP(), []int{2}
}

func (x *ListInstitutionsInput) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *ListInstitutionsInput) GetSize() int32 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *ListInstitutionsInput) GetFilters() *ListInstitutionsInput_Filter {
	if x != nil {
		return x.Filters
	}
	return nil
}

type ListInstitutionsOutput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Total int32          `protobuf:"varint,1,opt,name=total,proto3" json:"total,omitempty"`
	Items []*Institution `protobuf:"bytes,2,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *ListInstitutionsOutput) Reset() {
	*x = ListInstitutionsOutput{}
	mi := &file_eolymp_taxonomy_institution_service_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListInstitutionsOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListInstitutionsOutput) ProtoMessage() {}

func (x *ListInstitutionsOutput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_taxonomy_institution_service_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListInstitutionsOutput.ProtoReflect.Descriptor instead.
func (*ListInstitutionsOutput) Descriptor() ([]byte, []int) {
	return file_eolymp_taxonomy_institution_service_proto_rawDescGZIP(), []int{3}
}

func (x *ListInstitutionsOutput) GetTotal() int32 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *ListInstitutionsOutput) GetItems() []*Institution {
	if x != nil {
		return x.Items
	}
	return nil
}

type ListInstitutionsInput_Filter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Query      string                        `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
	Id         []*wellknown.ExpressionID     `protobuf:"bytes,2,rep,name=id,proto3" json:"id,omitempty"`
	Name       []*wellknown.ExpressionString `protobuf:"bytes,3,rep,name=name,proto3" json:"name,omitempty"`
	Acronym    []*wellknown.ExpressionString `protobuf:"bytes,9,rep,name=acronym,proto3" json:"acronym,omitempty"`
	CountryId  []*wellknown.ExpressionID     `protobuf:"bytes,4,rep,name=country_id,json=countryId,proto3" json:"country_id,omitempty"`
	RegionId   []*wellknown.ExpressionID     `protobuf:"bytes,5,rep,name=region_id,json=regionId,proto3" json:"region_id,omitempty"`
	Level      []*wellknown.ExpressionEnum   `protobuf:"bytes,6,rep,name=level,proto3" json:"level,omitempty"`
	Type       []*wellknown.ExpressionEnum   `protobuf:"bytes,7,rep,name=type,proto3" json:"type,omitempty"`
	Governance []*wellknown.ExpressionEnum   `protobuf:"bytes,8,rep,name=governance,proto3" json:"governance,omitempty"`
}

func (x *ListInstitutionsInput_Filter) Reset() {
	*x = ListInstitutionsInput_Filter{}
	mi := &file_eolymp_taxonomy_institution_service_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListInstitutionsInput_Filter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListInstitutionsInput_Filter) ProtoMessage() {}

func (x *ListInstitutionsInput_Filter) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_taxonomy_institution_service_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListInstitutionsInput_Filter.ProtoReflect.Descriptor instead.
func (*ListInstitutionsInput_Filter) Descriptor() ([]byte, []int) {
	return file_eolymp_taxonomy_institution_service_proto_rawDescGZIP(), []int{2, 0}
}

func (x *ListInstitutionsInput_Filter) GetQuery() string {
	if x != nil {
		return x.Query
	}
	return ""
}

func (x *ListInstitutionsInput_Filter) GetId() []*wellknown.ExpressionID {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *ListInstitutionsInput_Filter) GetName() []*wellknown.ExpressionString {
	if x != nil {
		return x.Name
	}
	return nil
}

func (x *ListInstitutionsInput_Filter) GetAcronym() []*wellknown.ExpressionString {
	if x != nil {
		return x.Acronym
	}
	return nil
}

func (x *ListInstitutionsInput_Filter) GetCountryId() []*wellknown.ExpressionID {
	if x != nil {
		return x.CountryId
	}
	return nil
}

func (x *ListInstitutionsInput_Filter) GetRegionId() []*wellknown.ExpressionID {
	if x != nil {
		return x.RegionId
	}
	return nil
}

func (x *ListInstitutionsInput_Filter) GetLevel() []*wellknown.ExpressionEnum {
	if x != nil {
		return x.Level
	}
	return nil
}

func (x *ListInstitutionsInput_Filter) GetType() []*wellknown.ExpressionEnum {
	if x != nil {
		return x.Type
	}
	return nil
}

func (x *ListInstitutionsInput_Filter) GetGovernance() []*wellknown.ExpressionEnum {
	if x != nil {
		return x.Governance
	}
	return nil
}

var File_eolymp_taxonomy_institution_service_proto protoreflect.FileDescriptor

var file_eolymp_taxonomy_institution_service_proto_rawDesc = []byte{
	0x0a, 0x29, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x74, 0x61, 0x78, 0x6f, 0x6e, 0x6f, 0x6d,
	0x79, 0x2f, 0x69, 0x6e, 0x73, 0x74, 0x69, 0x74, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x65, 0x6f, 0x6c,
	0x79, 0x6d, 0x70, 0x2e, 0x74, 0x61, 0x78, 0x6f, 0x6e, 0x6f, 0x6d, 0x79, 0x1a, 0x1d, 0x65, 0x6f,
	0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2f, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x22, 0x65, 0x6f, 0x6c,
	0x79, 0x6d, 0x70, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f,
	0x72, 0x61, 0x74, 0x65, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x21, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x74, 0x61, 0x78, 0x6f, 0x6e, 0x6f, 0x6d, 0x79,
	0x2f, 0x69, 0x6e, 0x73, 0x74, 0x69, 0x74, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x21, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x77, 0x65, 0x6c, 0x6c, 0x6b,
	0x6e, 0x6f, 0x77, 0x6e, 0x2f, 0x65, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x41, 0x0a, 0x18, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62,
	0x65, 0x49, 0x6e, 0x73, 0x74, 0x69, 0x74, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x70, 0x75,
	0x74, 0x12, 0x25, 0x0a, 0x0e, 0x69, 0x6e, 0x73, 0x74, 0x69, 0x74, 0x75, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x69, 0x6e, 0x73, 0x74, 0x69,
	0x74, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x22, 0x5b, 0x0a, 0x19, 0x44, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x62, 0x65, 0x49, 0x6e, 0x73, 0x74, 0x69, 0x74, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x4f,
	0x75, 0x74, 0x70, 0x75, 0x74, 0x12, 0x3e, 0x0a, 0x0b, 0x69, 0x6e, 0x73, 0x74, 0x69, 0x74, 0x75,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x65, 0x6f, 0x6c,
	0x79, 0x6d, 0x70, 0x2e, 0x74, 0x61, 0x78, 0x6f, 0x6e, 0x6f, 0x6d, 0x79, 0x2e, 0x49, 0x6e, 0x73,
	0x74, 0x69, 0x74, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x69, 0x6e, 0x73, 0x74, 0x69, 0x74,
	0x75, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0xff, 0x04, 0x0a, 0x15, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x6e,
	0x73, 0x74, 0x69, 0x74, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12,
	0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18,
	0x0b, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x47, 0x0a, 0x07, 0x66,
	0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x18, 0x28, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x65,
	0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x74, 0x61, 0x78, 0x6f, 0x6e, 0x6f, 0x6d, 0x79, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x49, 0x6e, 0x73, 0x74, 0x69, 0x74, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x49,
	0x6e, 0x70, 0x75, 0x74, 0x2e, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x52, 0x07, 0x66, 0x69, 0x6c,
	0x74, 0x65, 0x72, 0x73, 0x1a, 0xf0, 0x03, 0x0a, 0x06, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12,
	0x14, 0x0a, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x71, 0x75, 0x65, 0x72, 0x79, 0x12, 0x2e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x1e, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x77, 0x65, 0x6c, 0x6c, 0x6b,
	0x6e, 0x6f, 0x77, 0x6e, 0x2e, 0x45, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49,
	0x44, 0x52, 0x02, 0x69, 0x64, 0x12, 0x36, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x77, 0x65, 0x6c,
	0x6c, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x2e, 0x45, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x3c, 0x0a,
	0x07, 0x61, 0x63, 0x72, 0x6f, 0x6e, 0x79, 0x6d, 0x18, 0x09, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22,
	0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x77, 0x65, 0x6c, 0x6c, 0x6b, 0x6e, 0x6f, 0x77,
	0x6e, 0x2e, 0x45, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x72, 0x69,
	0x6e, 0x67, 0x52, 0x07, 0x61, 0x63, 0x72, 0x6f, 0x6e, 0x79, 0x6d, 0x12, 0x3d, 0x0a, 0x0a, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x1e, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x77, 0x65, 0x6c, 0x6c, 0x6b, 0x6e, 0x6f,
	0x77, 0x6e, 0x2e, 0x45, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x52,
	0x09, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x49, 0x64, 0x12, 0x3b, 0x0a, 0x09, 0x72, 0x65,
	0x67, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e,
	0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x77, 0x65, 0x6c, 0x6c, 0x6b, 0x6e, 0x6f, 0x77, 0x6e,
	0x2e, 0x45, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x52, 0x08, 0x72,
	0x65, 0x67, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x36, 0x0a, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c,
	0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e,
	0x77, 0x65, 0x6c, 0x6c, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x2e, 0x45, 0x78, 0x70, 0x72, 0x65, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x45, 0x6e, 0x75, 0x6d, 0x52, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x12,
	0x34, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e,
	0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x77, 0x65, 0x6c, 0x6c, 0x6b, 0x6e, 0x6f, 0x77, 0x6e,
	0x2e, 0x45, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x45, 0x6e, 0x75, 0x6d, 0x52,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x40, 0x0a, 0x0a, 0x67, 0x6f, 0x76, 0x65, 0x72, 0x6e, 0x61,
	0x6e, 0x63, 0x65, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x65, 0x6f, 0x6c, 0x79,
	0x6d, 0x70, 0x2e, 0x77, 0x65, 0x6c, 0x6c, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x2e, 0x45, 0x78, 0x70,
	0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x45, 0x6e, 0x75, 0x6d, 0x52, 0x0a, 0x67, 0x6f, 0x76,
	0x65, 0x72, 0x6e, 0x61, 0x6e, 0x63, 0x65, 0x22, 0x62, 0x0a, 0x16, 0x4c, 0x69, 0x73, 0x74, 0x49,
	0x6e, 0x73, 0x74, 0x69, 0x74, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x4f, 0x75, 0x74, 0x70, 0x75,
	0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x32, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e,
	0x74, 0x61, 0x78, 0x6f, 0x6e, 0x6f, 0x6d, 0x79, 0x2e, 0x49, 0x6e, 0x73, 0x74, 0x69, 0x74, 0x75,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x32, 0xc6, 0x02, 0x0a, 0x12,
	0x49, 0x6e, 0x73, 0x74, 0x69, 0x74, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x89, 0x01, 0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x6e, 0x73, 0x74, 0x69,
	0x74, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x26, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70,
	0x2e, 0x74, 0x61, 0x78, 0x6f, 0x6e, 0x6f, 0x6d, 0x79, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x6e,
	0x73, 0x74, 0x69, 0x74, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a,
	0x27, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x74, 0x61, 0x78, 0x6f, 0x6e, 0x6f, 0x6d,
	0x79, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x6e, 0x73, 0x74, 0x69, 0x74, 0x75, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x22, 0x24, 0xea, 0xe2, 0x0a, 0x0b, 0xf5, 0xe2,
	0x0a, 0x00, 0x00, 0xa0, 0x40, 0xf8, 0xe2, 0x0a, 0x14, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0f, 0x12,
	0x0d, 0x2f, 0x69, 0x6e, 0x73, 0x74, 0x69, 0x74, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0xa3,
	0x01, 0x0a, 0x13, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x49, 0x6e, 0x73, 0x74, 0x69,
	0x74, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x29, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e,
	0x74, 0x61, 0x78, 0x6f, 0x6e, 0x6f, 0x6d, 0x79, 0x2e, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62,
	0x65, 0x49, 0x6e, 0x73, 0x74, 0x69, 0x74, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x70, 0x75,
	0x74, 0x1a, 0x2a, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x74, 0x61, 0x78, 0x6f, 0x6e,
	0x6f, 0x6d, 0x79, 0x2e, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x49, 0x6e, 0x73, 0x74,
	0x69, 0x74, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x22, 0x35, 0xea,
	0xe2, 0x0a, 0x0b, 0xf5, 0xe2, 0x0a, 0x00, 0x00, 0xa0, 0x41, 0xf8, 0xe2, 0x0a, 0x64, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x20, 0x12, 0x1e, 0x2f, 0x69, 0x6e, 0x73, 0x74, 0x69, 0x74, 0x75, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2f, 0x7b, 0x69, 0x6e, 0x73, 0x74, 0x69, 0x74, 0x75, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x69, 0x64, 0x7d, 0x42, 0x33, 0x5a, 0x31, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x67, 0x6f, 0x2d, 0x73, 0x64, 0x6b,
	0x2f, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x74, 0x61, 0x78, 0x6f, 0x6e, 0x6f, 0x6d, 0x79,
	0x3b, 0x74, 0x61, 0x78, 0x6f, 0x6e, 0x6f, 0x6d, 0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_eolymp_taxonomy_institution_service_proto_rawDescOnce sync.Once
	file_eolymp_taxonomy_institution_service_proto_rawDescData = file_eolymp_taxonomy_institution_service_proto_rawDesc
)

func file_eolymp_taxonomy_institution_service_proto_rawDescGZIP() []byte {
	file_eolymp_taxonomy_institution_service_proto_rawDescOnce.Do(func() {
		file_eolymp_taxonomy_institution_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_eolymp_taxonomy_institution_service_proto_rawDescData)
	})
	return file_eolymp_taxonomy_institution_service_proto_rawDescData
}

var file_eolymp_taxonomy_institution_service_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_eolymp_taxonomy_institution_service_proto_goTypes = []any{
	(*DescribeInstitutionInput)(nil),     // 0: eolymp.taxonomy.DescribeInstitutionInput
	(*DescribeInstitutionOutput)(nil),    // 1: eolymp.taxonomy.DescribeInstitutionOutput
	(*ListInstitutionsInput)(nil),        // 2: eolymp.taxonomy.ListInstitutionsInput
	(*ListInstitutionsOutput)(nil),       // 3: eolymp.taxonomy.ListInstitutionsOutput
	(*ListInstitutionsInput_Filter)(nil), // 4: eolymp.taxonomy.ListInstitutionsInput.Filter
	(*Institution)(nil),                  // 5: eolymp.taxonomy.Institution
	(*wellknown.ExpressionID)(nil),       // 6: eolymp.wellknown.ExpressionID
	(*wellknown.ExpressionString)(nil),   // 7: eolymp.wellknown.ExpressionString
	(*wellknown.ExpressionEnum)(nil),     // 8: eolymp.wellknown.ExpressionEnum
}
var file_eolymp_taxonomy_institution_service_proto_depIdxs = []int32{
	5,  // 0: eolymp.taxonomy.DescribeInstitutionOutput.institution:type_name -> eolymp.taxonomy.Institution
	4,  // 1: eolymp.taxonomy.ListInstitutionsInput.filters:type_name -> eolymp.taxonomy.ListInstitutionsInput.Filter
	5,  // 2: eolymp.taxonomy.ListInstitutionsOutput.items:type_name -> eolymp.taxonomy.Institution
	6,  // 3: eolymp.taxonomy.ListInstitutionsInput.Filter.id:type_name -> eolymp.wellknown.ExpressionID
	7,  // 4: eolymp.taxonomy.ListInstitutionsInput.Filter.name:type_name -> eolymp.wellknown.ExpressionString
	7,  // 5: eolymp.taxonomy.ListInstitutionsInput.Filter.acronym:type_name -> eolymp.wellknown.ExpressionString
	6,  // 6: eolymp.taxonomy.ListInstitutionsInput.Filter.country_id:type_name -> eolymp.wellknown.ExpressionID
	6,  // 7: eolymp.taxonomy.ListInstitutionsInput.Filter.region_id:type_name -> eolymp.wellknown.ExpressionID
	8,  // 8: eolymp.taxonomy.ListInstitutionsInput.Filter.level:type_name -> eolymp.wellknown.ExpressionEnum
	8,  // 9: eolymp.taxonomy.ListInstitutionsInput.Filter.type:type_name -> eolymp.wellknown.ExpressionEnum
	8,  // 10: eolymp.taxonomy.ListInstitutionsInput.Filter.governance:type_name -> eolymp.wellknown.ExpressionEnum
	2,  // 11: eolymp.taxonomy.InstitutionService.ListInstitutions:input_type -> eolymp.taxonomy.ListInstitutionsInput
	0,  // 12: eolymp.taxonomy.InstitutionService.DescribeInstitution:input_type -> eolymp.taxonomy.DescribeInstitutionInput
	3,  // 13: eolymp.taxonomy.InstitutionService.ListInstitutions:output_type -> eolymp.taxonomy.ListInstitutionsOutput
	1,  // 14: eolymp.taxonomy.InstitutionService.DescribeInstitution:output_type -> eolymp.taxonomy.DescribeInstitutionOutput
	13, // [13:15] is the sub-list for method output_type
	11, // [11:13] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_eolymp_taxonomy_institution_service_proto_init() }
func file_eolymp_taxonomy_institution_service_proto_init() {
	if File_eolymp_taxonomy_institution_service_proto != nil {
		return
	}
	file_eolymp_taxonomy_institution_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_eolymp_taxonomy_institution_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_eolymp_taxonomy_institution_service_proto_goTypes,
		DependencyIndexes: file_eolymp_taxonomy_institution_service_proto_depIdxs,
		MessageInfos:      file_eolymp_taxonomy_institution_service_proto_msgTypes,
	}.Build()
	File_eolymp_taxonomy_institution_service_proto = out.File
	file_eolymp_taxonomy_institution_service_proto_rawDesc = nil
	file_eolymp_taxonomy_institution_service_proto_goTypes = nil
	file_eolymp_taxonomy_institution_service_proto_depIdxs = nil
}
