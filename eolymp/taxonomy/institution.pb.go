// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v3.21.12
// source: eolymp/taxonomy/institution.proto

package taxonomy

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

type Institution_Governance int32

const (
	Institution_UNKNOWN_GOVERNANCE Institution_Governance = 0
	Institution_PUBLIC             Institution_Governance = 1
	Institution_PRIVATE            Institution_Governance = 2
	Institution_CHARTER            Institution_Governance = 3
)

// Enum value maps for Institution_Governance.
var (
	Institution_Governance_name = map[int32]string{
		0: "UNKNOWN_GOVERNANCE",
		1: "PUBLIC",
		2: "PRIVATE",
		3: "CHARTER",
	}
	Institution_Governance_value = map[string]int32{
		"UNKNOWN_GOVERNANCE": 0,
		"PUBLIC":             1,
		"PRIVATE":            2,
		"CHARTER":            3,
	}
)

func (x Institution_Governance) Enum() *Institution_Governance {
	p := new(Institution_Governance)
	*p = x
	return p
}

func (x Institution_Governance) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Institution_Governance) Descriptor() protoreflect.EnumDescriptor {
	return file_eolymp_taxonomy_institution_proto_enumTypes[0].Descriptor()
}

func (Institution_Governance) Type() protoreflect.EnumType {
	return &file_eolymp_taxonomy_institution_proto_enumTypes[0]
}

func (x Institution_Governance) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Institution_Governance.Descriptor instead.
func (Institution_Governance) EnumDescriptor() ([]byte, []int) {
	return file_eolymp_taxonomy_institution_proto_rawDescGZIP(), []int{0, 0}
}

type Institution_Level int32

const (
	Institution_UNKNOWN_LEVEL Institution_Level = 0
	Institution_PRESCHOOL     Institution_Level = 1 // under 5 yo.: kindergartens, etc. (here mostly for completeness)
	Institution_PRIMARY       Institution_Level = 2 // 5-11 yo.: elementary schools, etc.
	Institution_SECONDARY     Institution_Level = 3 // 11-18 yo.: high schools, gymnasiums, etc.
	Institution_TERTIARY      Institution_Level = 4 // 18+ yo.: universities, colleges, etc.
)

// Enum value maps for Institution_Level.
var (
	Institution_Level_name = map[int32]string{
		0: "UNKNOWN_LEVEL",
		1: "PRESCHOOL",
		2: "PRIMARY",
		3: "SECONDARY",
		4: "TERTIARY",
	}
	Institution_Level_value = map[string]int32{
		"UNKNOWN_LEVEL": 0,
		"PRESCHOOL":     1,
		"PRIMARY":       2,
		"SECONDARY":     3,
		"TERTIARY":      4,
	}
)

func (x Institution_Level) Enum() *Institution_Level {
	p := new(Institution_Level)
	*p = x
	return p
}

func (x Institution_Level) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Institution_Level) Descriptor() protoreflect.EnumDescriptor {
	return file_eolymp_taxonomy_institution_proto_enumTypes[1].Descriptor()
}

func (Institution_Level) Type() protoreflect.EnumType {
	return &file_eolymp_taxonomy_institution_proto_enumTypes[1]
}

func (x Institution_Level) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Institution_Level.Descriptor instead.
func (Institution_Level) EnumDescriptor() ([]byte, []int) {
	return file_eolymp_taxonomy_institution_proto_rawDescGZIP(), []int{0, 1}
}

type Institution_Type int32

const (
	Institution_UNKNOWN_TYPE Institution_Type = 0
	Institution_KINDERGARTEN Institution_Type = 1
	Institution_SCHOOL       Institution_Type = 2
	Institution_LYCEUM       Institution_Type = 3
	Institution_GYMNASIUM    Institution_Type = 4
	Institution_COLLEGE      Institution_Type = 5
	Institution_INSTITUTE    Institution_Type = 6
	Institution_UNIVERSITY   Institution_Type = 7
	Institution_ACADEMY      Institution_Type = 8
)

// Enum value maps for Institution_Type.
var (
	Institution_Type_name = map[int32]string{
		0: "UNKNOWN_TYPE",
		1: "KINDERGARTEN",
		2: "SCHOOL",
		3: "LYCEUM",
		4: "GYMNASIUM",
		5: "COLLEGE",
		6: "INSTITUTE",
		7: "UNIVERSITY",
		8: "ACADEMY",
	}
	Institution_Type_value = map[string]int32{
		"UNKNOWN_TYPE": 0,
		"KINDERGARTEN": 1,
		"SCHOOL":       2,
		"LYCEUM":       3,
		"GYMNASIUM":    4,
		"COLLEGE":      5,
		"INSTITUTE":    6,
		"UNIVERSITY":   7,
		"ACADEMY":      8,
	}
)

func (x Institution_Type) Enum() *Institution_Type {
	p := new(Institution_Type)
	*p = x
	return p
}

func (x Institution_Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Institution_Type) Descriptor() protoreflect.EnumDescriptor {
	return file_eolymp_taxonomy_institution_proto_enumTypes[2].Descriptor()
}

func (Institution_Type) Type() protoreflect.EnumType {
	return &file_eolymp_taxonomy_institution_proto_enumTypes[2]
}

func (x Institution_Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Institution_Type.Descriptor instead.
func (Institution_Type) EnumDescriptor() ([]byte, []int) {
	return file_eolymp_taxonomy_institution_proto_rawDescGZIP(), []int{0, 2}
}

type Institution struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name         string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Acronym      string                 `protobuf:"bytes,4,opt,name=acronym,proto3" json:"acronym,omitempty"`
	LocalName    string                 `protobuf:"bytes,3,opt,name=local_name,json=localName,proto3" json:"local_name,omitempty"`
	LocalAcronym string                 `protobuf:"bytes,5,opt,name=local_acronym,json=localAcronym,proto3" json:"local_acronym,omitempty"`
	LogoUrl      string                 `protobuf:"bytes,20,opt,name=logo_url,json=logoUrl,proto3" json:"logo_url,omitempty"`
	Governance   Institution_Governance `protobuf:"varint,10,opt,name=governance,proto3,enum=eolymp.taxonomy.Institution_Governance" json:"governance,omitempty"`
	Level        Institution_Level      `protobuf:"varint,11,opt,name=level,proto3,enum=eolymp.taxonomy.Institution_Level" json:"level,omitempty"`
	Type         Institution_Type       `protobuf:"varint,12,opt,name=type,proto3,enum=eolymp.taxonomy.Institution_Type" json:"type,omitempty"`
	CountryId    string                 `protobuf:"bytes,30,opt,name=country_id,json=countryId,proto3" json:"country_id,omitempty"`
	RegionId     string                 `protobuf:"bytes,31,opt,name=region_id,json=regionId,proto3" json:"region_id,omitempty"`
	Address      string                 `protobuf:"bytes,32,opt,name=address,proto3" json:"address,omitempty"`
	Links        map[string]string      `protobuf:"bytes,41,rep,name=links,proto3" json:"links,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Institution) Reset() {
	*x = Institution{}
	mi := &file_eolymp_taxonomy_institution_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Institution) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Institution) ProtoMessage() {}

func (x *Institution) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_taxonomy_institution_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Institution.ProtoReflect.Descriptor instead.
func (*Institution) Descriptor() ([]byte, []int) {
	return file_eolymp_taxonomy_institution_proto_rawDescGZIP(), []int{0}
}

func (x *Institution) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Institution) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Institution) GetAcronym() string {
	if x != nil {
		return x.Acronym
	}
	return ""
}

func (x *Institution) GetLocalName() string {
	if x != nil {
		return x.LocalName
	}
	return ""
}

func (x *Institution) GetLocalAcronym() string {
	if x != nil {
		return x.LocalAcronym
	}
	return ""
}

func (x *Institution) GetLogoUrl() string {
	if x != nil {
		return x.LogoUrl
	}
	return ""
}

func (x *Institution) GetGovernance() Institution_Governance {
	if x != nil {
		return x.Governance
	}
	return Institution_UNKNOWN_GOVERNANCE
}

func (x *Institution) GetLevel() Institution_Level {
	if x != nil {
		return x.Level
	}
	return Institution_UNKNOWN_LEVEL
}

func (x *Institution) GetType() Institution_Type {
	if x != nil {
		return x.Type
	}
	return Institution_UNKNOWN_TYPE
}

func (x *Institution) GetCountryId() string {
	if x != nil {
		return x.CountryId
	}
	return ""
}

func (x *Institution) GetRegionId() string {
	if x != nil {
		return x.RegionId
	}
	return ""
}

func (x *Institution) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *Institution) GetLinks() map[string]string {
	if x != nil {
		return x.Links
	}
	return nil
}

var File_eolymp_taxonomy_institution_proto protoreflect.FileDescriptor

var file_eolymp_taxonomy_institution_proto_rawDesc = []byte{
	0x0a, 0x21, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x74, 0x61, 0x78, 0x6f, 0x6e, 0x6f, 0x6d,
	0x79, 0x2f, 0x69, 0x6e, 0x73, 0x74, 0x69, 0x74, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x74, 0x61, 0x78, 0x6f,
	0x6e, 0x6f, 0x6d, 0x79, 0x22, 0xe1, 0x06, 0x0a, 0x0b, 0x49, 0x6e, 0x73, 0x74, 0x69, 0x74, 0x75,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x63, 0x72, 0x6f,
	0x6e, 0x79, 0x6d, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x63, 0x72, 0x6f, 0x6e,
	0x79, 0x6d, 0x12, 0x1d, 0x0a, 0x0a, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x23, 0x0a, 0x0d, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x5f, 0x61, 0x63, 0x72, 0x6f, 0x6e,
	0x79, 0x6d, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x41,
	0x63, 0x72, 0x6f, 0x6e, 0x79, 0x6d, 0x12, 0x19, 0x0a, 0x08, 0x6c, 0x6f, 0x67, 0x6f, 0x5f, 0x75,
	0x72, 0x6c, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6c, 0x6f, 0x67, 0x6f, 0x55, 0x72,
	0x6c, 0x12, 0x47, 0x0a, 0x0a, 0x67, 0x6f, 0x76, 0x65, 0x72, 0x6e, 0x61, 0x6e, 0x63, 0x65, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x27, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x74,
	0x61, 0x78, 0x6f, 0x6e, 0x6f, 0x6d, 0x79, 0x2e, 0x49, 0x6e, 0x73, 0x74, 0x69, 0x74, 0x75, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x47, 0x6f, 0x76, 0x65, 0x72, 0x6e, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x0a,
	0x67, 0x6f, 0x76, 0x65, 0x72, 0x6e, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x38, 0x0a, 0x05, 0x6c, 0x65,
	0x76, 0x65, 0x6c, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x22, 0x2e, 0x65, 0x6f, 0x6c, 0x79,
	0x6d, 0x70, 0x2e, 0x74, 0x61, 0x78, 0x6f, 0x6e, 0x6f, 0x6d, 0x79, 0x2e, 0x49, 0x6e, 0x73, 0x74,
	0x69, 0x74, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x52, 0x05, 0x6c,
	0x65, 0x76, 0x65, 0x6c, 0x12, 0x35, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x0c, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x21, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x74, 0x61, 0x78, 0x6f,
	0x6e, 0x6f, 0x6d, 0x79, 0x2e, 0x49, 0x6e, 0x73, 0x74, 0x69, 0x74, 0x75, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x72, 0x65,
	0x67, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x1f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72,
	0x65, 0x67, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x18, 0x20, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x12, 0x3d, 0x0a, 0x05, 0x6c, 0x69, 0x6e, 0x6b, 0x73, 0x18, 0x29, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x27, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x74, 0x61, 0x78, 0x6f, 0x6e, 0x6f,
	0x6d, 0x79, 0x2e, 0x49, 0x6e, 0x73, 0x74, 0x69, 0x74, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x4c,
	0x69, 0x6e, 0x6b, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x05, 0x6c, 0x69, 0x6e, 0x6b, 0x73,
	0x1a, 0x38, 0x0a, 0x0a, 0x4c, 0x69, 0x6e, 0x6b, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10,
	0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79,
	0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x4a, 0x0a, 0x0a, 0x47, 0x6f,
	0x76, 0x65, 0x72, 0x6e, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x16, 0x0a, 0x12, 0x55, 0x4e, 0x4b, 0x4e,
	0x4f, 0x57, 0x4e, 0x5f, 0x47, 0x4f, 0x56, 0x45, 0x52, 0x4e, 0x41, 0x4e, 0x43, 0x45, 0x10, 0x00,
	0x12, 0x0a, 0x0a, 0x06, 0x50, 0x55, 0x42, 0x4c, 0x49, 0x43, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07,
	0x50, 0x52, 0x49, 0x56, 0x41, 0x54, 0x45, 0x10, 0x02, 0x12, 0x0b, 0x0a, 0x07, 0x43, 0x48, 0x41,
	0x52, 0x54, 0x45, 0x52, 0x10, 0x03, 0x22, 0x53, 0x0a, 0x05, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12,
	0x11, 0x0a, 0x0d, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x5f, 0x4c, 0x45, 0x56, 0x45, 0x4c,
	0x10, 0x00, 0x12, 0x0d, 0x0a, 0x09, 0x50, 0x52, 0x45, 0x53, 0x43, 0x48, 0x4f, 0x4f, 0x4c, 0x10,
	0x01, 0x12, 0x0b, 0x0a, 0x07, 0x50, 0x52, 0x49, 0x4d, 0x41, 0x52, 0x59, 0x10, 0x02, 0x12, 0x0d,
	0x0a, 0x09, 0x53, 0x45, 0x43, 0x4f, 0x4e, 0x44, 0x41, 0x52, 0x59, 0x10, 0x03, 0x12, 0x0c, 0x0a,
	0x08, 0x54, 0x45, 0x52, 0x54, 0x49, 0x41, 0x52, 0x59, 0x10, 0x04, 0x22, 0x8a, 0x01, 0x0a, 0x04,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x10, 0x0a, 0x0c, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x5f,
	0x54, 0x59, 0x50, 0x45, 0x10, 0x00, 0x12, 0x10, 0x0a, 0x0c, 0x4b, 0x49, 0x4e, 0x44, 0x45, 0x52,
	0x47, 0x41, 0x52, 0x54, 0x45, 0x4e, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x53, 0x43, 0x48, 0x4f,
	0x4f, 0x4c, 0x10, 0x02, 0x12, 0x0a, 0x0a, 0x06, 0x4c, 0x59, 0x43, 0x45, 0x55, 0x4d, 0x10, 0x03,
	0x12, 0x0d, 0x0a, 0x09, 0x47, 0x59, 0x4d, 0x4e, 0x41, 0x53, 0x49, 0x55, 0x4d, 0x10, 0x04, 0x12,
	0x0b, 0x0a, 0x07, 0x43, 0x4f, 0x4c, 0x4c, 0x45, 0x47, 0x45, 0x10, 0x05, 0x12, 0x0d, 0x0a, 0x09,
	0x49, 0x4e, 0x53, 0x54, 0x49, 0x54, 0x55, 0x54, 0x45, 0x10, 0x06, 0x12, 0x0e, 0x0a, 0x0a, 0x55,
	0x4e, 0x49, 0x56, 0x45, 0x52, 0x53, 0x49, 0x54, 0x59, 0x10, 0x07, 0x12, 0x0b, 0x0a, 0x07, 0x41,
	0x43, 0x41, 0x44, 0x45, 0x4d, 0x59, 0x10, 0x08, 0x42, 0x33, 0x5a, 0x31, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x67, 0x6f,
	0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x74, 0x61, 0x78, 0x6f,
	0x6e, 0x6f, 0x6d, 0x79, 0x3b, 0x74, 0x61, 0x78, 0x6f, 0x6e, 0x6f, 0x6d, 0x79, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_eolymp_taxonomy_institution_proto_rawDescOnce sync.Once
	file_eolymp_taxonomy_institution_proto_rawDescData = file_eolymp_taxonomy_institution_proto_rawDesc
)

func file_eolymp_taxonomy_institution_proto_rawDescGZIP() []byte {
	file_eolymp_taxonomy_institution_proto_rawDescOnce.Do(func() {
		file_eolymp_taxonomy_institution_proto_rawDescData = protoimpl.X.CompressGZIP(file_eolymp_taxonomy_institution_proto_rawDescData)
	})
	return file_eolymp_taxonomy_institution_proto_rawDescData
}

var file_eolymp_taxonomy_institution_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_eolymp_taxonomy_institution_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_eolymp_taxonomy_institution_proto_goTypes = []any{
	(Institution_Governance)(0), // 0: eolymp.taxonomy.Institution.Governance
	(Institution_Level)(0),      // 1: eolymp.taxonomy.Institution.Level
	(Institution_Type)(0),       // 2: eolymp.taxonomy.Institution.Type
	(*Institution)(nil),         // 3: eolymp.taxonomy.Institution
	nil,                         // 4: eolymp.taxonomy.Institution.LinksEntry
}
var file_eolymp_taxonomy_institution_proto_depIdxs = []int32{
	0, // 0: eolymp.taxonomy.Institution.governance:type_name -> eolymp.taxonomy.Institution.Governance
	1, // 1: eolymp.taxonomy.Institution.level:type_name -> eolymp.taxonomy.Institution.Level
	2, // 2: eolymp.taxonomy.Institution.type:type_name -> eolymp.taxonomy.Institution.Type
	4, // 3: eolymp.taxonomy.Institution.links:type_name -> eolymp.taxonomy.Institution.LinksEntry
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_eolymp_taxonomy_institution_proto_init() }
func file_eolymp_taxonomy_institution_proto_init() {
	if File_eolymp_taxonomy_institution_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_eolymp_taxonomy_institution_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_eolymp_taxonomy_institution_proto_goTypes,
		DependencyIndexes: file_eolymp_taxonomy_institution_proto_depIdxs,
		EnumInfos:         file_eolymp_taxonomy_institution_proto_enumTypes,
		MessageInfos:      file_eolymp_taxonomy_institution_proto_msgTypes,
	}.Build()
	File_eolymp_taxonomy_institution_proto = out.File
	file_eolymp_taxonomy_institution_proto_rawDesc = nil
	file_eolymp_taxonomy_institution_proto_goTypes = nil
	file_eolymp_taxonomy_institution_proto_depIdxs = nil
}
