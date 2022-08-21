// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.18.1
// source: eolymp/community/attribute.proto

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

// Field type defines looks and type of the data for the field.
type Attribute_Type int32

const (
	Attribute_UNKNOWN Attribute_Type = 0
	// Single line of text.
	//
	// Validation:
	//  - required - string must be non-empty
	//  - min - minimal string length
	//  - max - maximal string length
	//  - regexp - regular expression (RE2) which string must match completely (add .* to configure partial match)
	Attribute_STRING Attribute_Type = 1
	// Multiline text.
	//
	// Validation:
	//  - required - text must be non-empty
	//  - min - minimal text length
	//  - max - maximal text length
	//  - regexp - regular expression (RE2)
	Attribute_TEXT Attribute_Type = 2
	// An integer number.
	//
	// Validation:
	//  - required - field must be non-empty (0 is considered non-empty value)
	//  - min - minimal value
	//  - max - maximal value
	Attribute_NUMBER Attribute_Type = 3
	// Choice, a dropdown with options.
	//
	// Validation:
	//  - required - field must be non-empty
	//  - choices - available values
	Attribute_CHOICE Attribute_Type = 4
	// Date picker.
	//
	// Validation:
	//  - required - field must be non-empty
	Attribute_DATE Attribute_Type = 5
	// Email.
	//
	// Validation:
	//  - required - field must be non-empty
	Attribute_EMAIL Attribute_Type = 6
	// Checkbox gives simple yes/no value.
	//
	// Validation:
	//  - required - field must be checked
	Attribute_CHECKBOX Attribute_Type = 7
	// Country value.
	//
	// Validation:
	//  - required - field must be non-empty
	Attribute_COUNTRY Attribute_Type = 8
	// Country and Region value.
	//
	// Validation:
	//  - required - field must be non-empty
	//  - country  - region must belong to a specific country
	Attribute_REGION Attribute_Type = 9
)

// Enum value maps for Attribute_Type.
var (
	Attribute_Type_name = map[int32]string{
		0: "UNKNOWN",
		1: "STRING",
		2: "TEXT",
		3: "NUMBER",
		4: "CHOICE",
		5: "DATE",
		6: "EMAIL",
		7: "CHECKBOX",
		8: "COUNTRY",
		9: "REGION",
	}
	Attribute_Type_value = map[string]int32{
		"UNKNOWN":  0,
		"STRING":   1,
		"TEXT":     2,
		"NUMBER":   3,
		"CHOICE":   4,
		"DATE":     5,
		"EMAIL":    6,
		"CHECKBOX": 7,
		"COUNTRY":  8,
		"REGION":   9,
	}
)

func (x Attribute_Type) Enum() *Attribute_Type {
	p := new(Attribute_Type)
	*p = x
	return p
}

func (x Attribute_Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Attribute_Type) Descriptor() protoreflect.EnumDescriptor {
	return file_eolymp_community_attribute_proto_enumTypes[0].Descriptor()
}

func (Attribute_Type) Type() protoreflect.EnumType {
	return &file_eolymp_community_attribute_proto_enumTypes[0]
}

func (x Attribute_Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Attribute_Type.Descriptor instead.
func (Attribute_Type) EnumDescriptor() ([]byte, []int) {
	return file_eolymp_community_attribute_proto_rawDescGZIP(), []int{0, 0}
}

type Attribute struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key         string                   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"` // unique field identifier (should be unique within the form, not globally unique)
	Ern         string                   `protobuf:"bytes,9999,opt,name=ern,proto3" json:"ern,omitempty"`
	Description []*Attribute_Description `protobuf:"bytes,10,rep,name=description,proto3" json:"description,omitempty"`                         // field localized data, such as label and help message
	Type        Attribute_Type           `protobuf:"varint,20,opt,name=type,proto3,enum=eolymp.community.Attribute_Type" json:"type,omitempty"` // type of the field
	Index       uint32                   `protobuf:"varint,21,opt,name=index,proto3" json:"index,omitempty"`
	Required    bool                     `protobuf:"varint,31,opt,name=required,proto3" json:"required,omitempty"` // field is required (see field types for details)
	Hidden      bool                     `protobuf:"varint,32,opt,name=hidden,proto3" json:"hidden,omitempty"`     // field is hidden (only visible to admin)
	// validation
	Regexp  string   `protobuf:"bytes,100,opt,name=regexp,proto3" json:"regexp,omitempty"`   // regexp validation (see field types for details)
	Min     int32    `protobuf:"varint,101,opt,name=min,proto3" json:"min,omitempty"`        // min value validation (see field types for details)
	Max     int32    `protobuf:"varint,102,opt,name=max,proto3" json:"max,omitempty"`        // max value validation (see field types for details)
	Choices []string `protobuf:"bytes,103,rep,name=choices,proto3" json:"choices,omitempty"` // possible choices validation (see field types for details)
	Country string   `protobuf:"bytes,104,opt,name=country,proto3" json:"country,omitempty"` // restrict region selector to a specific country
}

func (x *Attribute) Reset() {
	*x = Attribute{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_community_attribute_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Attribute) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Attribute) ProtoMessage() {}

func (x *Attribute) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_community_attribute_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Attribute.ProtoReflect.Descriptor instead.
func (*Attribute) Descriptor() ([]byte, []int) {
	return file_eolymp_community_attribute_proto_rawDescGZIP(), []int{0}
}

func (x *Attribute) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *Attribute) GetErn() string {
	if x != nil {
		return x.Ern
	}
	return ""
}

func (x *Attribute) GetDescription() []*Attribute_Description {
	if x != nil {
		return x.Description
	}
	return nil
}

func (x *Attribute) GetType() Attribute_Type {
	if x != nil {
		return x.Type
	}
	return Attribute_UNKNOWN
}

func (x *Attribute) GetIndex() uint32 {
	if x != nil {
		return x.Index
	}
	return 0
}

func (x *Attribute) GetRequired() bool {
	if x != nil {
		return x.Required
	}
	return false
}

func (x *Attribute) GetHidden() bool {
	if x != nil {
		return x.Hidden
	}
	return false
}

func (x *Attribute) GetRegexp() string {
	if x != nil {
		return x.Regexp
	}
	return ""
}

func (x *Attribute) GetMin() int32 {
	if x != nil {
		return x.Min
	}
	return 0
}

func (x *Attribute) GetMax() int32 {
	if x != nil {
		return x.Max
	}
	return 0
}

func (x *Attribute) GetChoices() []string {
	if x != nil {
		return x.Choices
	}
	return nil
}

func (x *Attribute) GetCountry() string {
	if x != nil {
		return x.Country
	}
	return ""
}

// Description provides localized information about field.
type Attribute_Description struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Default bool     `protobuf:"varint,1,opt,name=default,proto3" json:"default,omitempty"` // default description used in case translation is not available
	Locale  string   `protobuf:"bytes,2,opt,name=locale,proto3" json:"locale,omitempty"`    // locale
	Label   string   `protobuf:"bytes,3,opt,name=label,proto3" json:"label,omitempty"`      // label, normally displayed above the field
	Help    string   `protobuf:"bytes,4,opt,name=help,proto3" json:"help,omitempty"`        // help message, normally displayed right below the field
	Choices []string `protobuf:"bytes,5,rep,name=choices,proto3" json:"choices,omitempty"`  // translation for choices (must be in the same order)
}

func (x *Attribute_Description) Reset() {
	*x = Attribute_Description{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_community_attribute_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Attribute_Description) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Attribute_Description) ProtoMessage() {}

func (x *Attribute_Description) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_community_attribute_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Attribute_Description.ProtoReflect.Descriptor instead.
func (*Attribute_Description) Descriptor() ([]byte, []int) {
	return file_eolymp_community_attribute_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Attribute_Description) GetDefault() bool {
	if x != nil {
		return x.Default
	}
	return false
}

func (x *Attribute_Description) GetLocale() string {
	if x != nil {
		return x.Locale
	}
	return ""
}

func (x *Attribute_Description) GetLabel() string {
	if x != nil {
		return x.Label
	}
	return ""
}

func (x *Attribute_Description) GetHelp() string {
	if x != nil {
		return x.Help
	}
	return ""
}

func (x *Attribute_Description) GetChoices() []string {
	if x != nil {
		return x.Choices
	}
	return nil
}

var File_eolymp_community_attribute_proto protoreflect.FileDescriptor

var file_eolymp_community_attribute_proto_rawDesc = []byte{
	0x0a, 0x20, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69,
	0x74, 0x79, 0x2f, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x10, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x75,
	0x6e, 0x69, 0x74, 0x79, 0x1a, 0x21, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa3, 0x05, 0x0a, 0x09, 0x41, 0x74, 0x74, 0x72,
	0x69, 0x62, 0x75, 0x74, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x11, 0x0a, 0x03, 0x65, 0x72, 0x6e, 0x18, 0x8f,
	0x4e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x65, 0x72, 0x6e, 0x12, 0x49, 0x0a, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x27, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69,
	0x74, 0x79, 0x2e, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x2e, 0x44, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x34, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x14, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x20, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x2e, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65,
	0x2e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x69,
	0x6e, 0x64, 0x65, 0x78, 0x18, 0x15, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x69, 0x6e, 0x64, 0x65,
	0x78, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x18, 0x1f, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x08, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x12, 0x16, 0x0a,
	0x06, 0x68, 0x69, 0x64, 0x64, 0x65, 0x6e, 0x18, 0x20, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x68,
	0x69, 0x64, 0x64, 0x65, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x67, 0x65, 0x78, 0x70, 0x18,
	0x64, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x67, 0x65, 0x78, 0x70, 0x12, 0x10, 0x0a,
	0x03, 0x6d, 0x69, 0x6e, 0x18, 0x65, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x6d, 0x69, 0x6e, 0x12,
	0x10, 0x0a, 0x03, 0x6d, 0x61, 0x78, 0x18, 0x66, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x6d, 0x61,
	0x78, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x68, 0x6f, 0x69, 0x63, 0x65, 0x73, 0x18, 0x67, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x07, 0x63, 0x68, 0x6f, 0x69, 0x63, 0x65, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x68, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x72, 0x79, 0x1a, 0x83, 0x01, 0x0a, 0x0b, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x12,
	0x16, 0x0a, 0x06, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x12, 0x12, 0x0a,
	0x04, 0x68, 0x65, 0x6c, 0x70, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x65, 0x6c,
	0x70, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x68, 0x6f, 0x69, 0x63, 0x65, 0x73, 0x18, 0x05, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x07, 0x63, 0x68, 0x6f, 0x69, 0x63, 0x65, 0x73, 0x22, 0x7d, 0x0a, 0x04, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00,
	0x12, 0x0a, 0x0a, 0x06, 0x53, 0x54, 0x52, 0x49, 0x4e, 0x47, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04,
	0x54, 0x45, 0x58, 0x54, 0x10, 0x02, 0x12, 0x0a, 0x0a, 0x06, 0x4e, 0x55, 0x4d, 0x42, 0x45, 0x52,
	0x10, 0x03, 0x12, 0x0a, 0x0a, 0x06, 0x43, 0x48, 0x4f, 0x49, 0x43, 0x45, 0x10, 0x04, 0x12, 0x08,
	0x0a, 0x04, 0x44, 0x41, 0x54, 0x45, 0x10, 0x05, 0x12, 0x09, 0x0a, 0x05, 0x45, 0x4d, 0x41, 0x49,
	0x4c, 0x10, 0x06, 0x12, 0x0c, 0x0a, 0x08, 0x43, 0x48, 0x45, 0x43, 0x4b, 0x42, 0x4f, 0x58, 0x10,
	0x07, 0x12, 0x0b, 0x0a, 0x07, 0x43, 0x4f, 0x55, 0x4e, 0x54, 0x52, 0x59, 0x10, 0x08, 0x12, 0x0a,
	0x0a, 0x06, 0x52, 0x45, 0x47, 0x49, 0x4f, 0x4e, 0x10, 0x09, 0x3a, 0x31, 0xb2, 0xe3, 0x0a, 0x2d,
	0xba, 0xe3, 0x0a, 0x09, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0xc2, 0xe3, 0x0a,
	0x15, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x75, 0x6e, 0x69, 0x76, 0x65, 0x72, 0x73, 0x65,
	0x2e, 0x53, 0x70, 0x61, 0x63, 0x65, 0xca, 0xe3, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x42, 0x35, 0x5a,
	0x33, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6f, 0x6c, 0x79,
	0x6d, 0x70, 0x2f, 0x67, 0x6f, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70,
	0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x3b, 0x63, 0x6f, 0x6d, 0x6d, 0x75,
	0x6e, 0x69, 0x74, 0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_eolymp_community_attribute_proto_rawDescOnce sync.Once
	file_eolymp_community_attribute_proto_rawDescData = file_eolymp_community_attribute_proto_rawDesc
)

func file_eolymp_community_attribute_proto_rawDescGZIP() []byte {
	file_eolymp_community_attribute_proto_rawDescOnce.Do(func() {
		file_eolymp_community_attribute_proto_rawDescData = protoimpl.X.CompressGZIP(file_eolymp_community_attribute_proto_rawDescData)
	})
	return file_eolymp_community_attribute_proto_rawDescData
}

var file_eolymp_community_attribute_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_eolymp_community_attribute_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_eolymp_community_attribute_proto_goTypes = []interface{}{
	(Attribute_Type)(0),           // 0: eolymp.community.Attribute.Type
	(*Attribute)(nil),             // 1: eolymp.community.Attribute
	(*Attribute_Description)(nil), // 2: eolymp.community.Attribute.Description
}
var file_eolymp_community_attribute_proto_depIdxs = []int32{
	2, // 0: eolymp.community.Attribute.description:type_name -> eolymp.community.Attribute.Description
	0, // 1: eolymp.community.Attribute.type:type_name -> eolymp.community.Attribute.Type
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_eolymp_community_attribute_proto_init() }
func file_eolymp_community_attribute_proto_init() {
	if File_eolymp_community_attribute_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_eolymp_community_attribute_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Attribute); i {
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
		file_eolymp_community_attribute_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Attribute_Description); i {
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
			RawDescriptor: file_eolymp_community_attribute_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_eolymp_community_attribute_proto_goTypes,
		DependencyIndexes: file_eolymp_community_attribute_proto_depIdxs,
		EnumInfos:         file_eolymp_community_attribute_proto_enumTypes,
		MessageInfos:      file_eolymp_community_attribute_proto_msgTypes,
	}.Build()
	File_eolymp_community_attribute_proto = out.File
	file_eolymp_community_attribute_proto_rawDesc = nil
	file_eolymp_community_attribute_proto_goTypes = nil
	file_eolymp_community_attribute_proto_depIdxs = nil
}
