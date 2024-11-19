// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v3.21.12
// source: eolymp/community/attribute.proto

package community

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

type Attribute_Patch int32

const (
	Attribute_PATCH_ALL         Attribute_Patch = 0
	Attribute_PATCH_LABEL       Attribute_Patch = 1
	Attribute_PATCH_HELP        Attribute_Patch = 2
	Attribute_PATCH_TYPE        Attribute_Patch = 3
	Attribute_PATCH_INDEX       Attribute_Patch = 4
	Attribute_PATCH_REQUIRED    Attribute_Patch = 5
	Attribute_PATCH_READONLY    Attribute_Patch = 6
	Attribute_PATCH_VISIBILITY  Attribute_Patch = 7
	Attribute_PATCH_REGEXP      Attribute_Patch = 8
	Attribute_PATCH_MIN         Attribute_Patch = 9
	Attribute_PATCH_MAX         Attribute_Patch = 10
	Attribute_PATCH_CHOICES     Attribute_Patch = 11
	Attribute_PATCH_CONSTRAINTS Attribute_Patch = 12
)

// Enum value maps for Attribute_Patch.
var (
	Attribute_Patch_name = map[int32]string{
		0:  "PATCH_ALL",
		1:  "PATCH_LABEL",
		2:  "PATCH_HELP",
		3:  "PATCH_TYPE",
		4:  "PATCH_INDEX",
		5:  "PATCH_REQUIRED",
		6:  "PATCH_READONLY",
		7:  "PATCH_VISIBILITY",
		8:  "PATCH_REGEXP",
		9:  "PATCH_MIN",
		10: "PATCH_MAX",
		11: "PATCH_CHOICES",
		12: "PATCH_CONSTRAINTS",
	}
	Attribute_Patch_value = map[string]int32{
		"PATCH_ALL":         0,
		"PATCH_LABEL":       1,
		"PATCH_HELP":        2,
		"PATCH_TYPE":        3,
		"PATCH_INDEX":       4,
		"PATCH_REQUIRED":    5,
		"PATCH_READONLY":    6,
		"PATCH_VISIBILITY":  7,
		"PATCH_REGEXP":      8,
		"PATCH_MIN":         9,
		"PATCH_MAX":         10,
		"PATCH_CHOICES":     11,
		"PATCH_CONSTRAINTS": 12,
	}
)

func (x Attribute_Patch) Enum() *Attribute_Patch {
	p := new(Attribute_Patch)
	*p = x
	return p
}

func (x Attribute_Patch) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Attribute_Patch) Descriptor() protoreflect.EnumDescriptor {
	return file_eolymp_community_attribute_proto_enumTypes[0].Descriptor()
}

func (Attribute_Patch) Type() protoreflect.EnumType {
	return &file_eolymp_community_attribute_proto_enumTypes[0]
}

func (x Attribute_Patch) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Attribute_Patch.Descriptor instead.
func (Attribute_Patch) EnumDescriptor() ([]byte, []int) {
	return file_eolymp_community_attribute_proto_rawDescGZIP(), []int{0, 0}
}

// Field type defines looks and type of the data for the field.
type Attribute_Type int32

const (
	Attribute_UNKNOWN_TYPE Attribute_Type = 0
	// Single line of text.
	//
	// Validation:
	//   - required - string must be non-empty
	//   - min - minimal string length
	//   - max - maximal string length
	//   - regexp - regular expression (RE2) which string must match completely (add .* to configure partial match)
	Attribute_STRING Attribute_Type = 1
	// Multiline text.
	//
	// Validation:
	//   - required - text must be non-empty
	//   - min - minimal text length
	//   - max - maximal text length
	//   - regexp - regular expression (RE2)
	Attribute_TEXT Attribute_Type = 2
	// An integer number.
	//
	// Validation:
	//   - required - field must be non-empty (0 is considered non-empty value)
	//   - min - minimal value
	//   - max - maximal value
	Attribute_NUMBER Attribute_Type = 3
	// Choice, a dropdown with options.
	//
	// Validation:
	//   - required - field must be non-empty
	//   - choices - available values
	Attribute_CHOICE Attribute_Type = 4
	// Date picker.
	//
	// Validation:
	//   - required - field must be non-empty
	Attribute_DATE Attribute_Type = 5
	// Email.
	//
	// Validation:
	//   - required - field must be non-empty
	Attribute_EMAIL Attribute_Type = 6
	// Checkbox gives simple yes/no value.
	//
	// Validation:
	//   - required - field must be checked
	Attribute_CHECKBOX Attribute_Type = 7
	// Country value.
	//
	// Validation:
	//   - required - field must be non-empty
	Attribute_COUNTRY Attribute_Type = 8
	// Country and Region value.
	//
	// Validation:
	//   - required  - field must be non-empty
	//   - countries - region must belong to a specific country
	Attribute_REGION Attribute_Type = 9
	// Country and Region value.
	//
	// Validation:
	//   - required  - field must be non-empty
	//
	// Constraints:
	//   - governance:public
	//   - governance:private
	//   - governance:charter
	//   - level:preschool
	//   - level:primary
	//   - level:secondary
	//   - level:tertiary
	//   - country:[two-letter]
	Attribute_INSTITUTION Attribute_Type = 10
	// An image as a URL to eolympusercontent.com.
	//
	// Validation:
	//   - required - field must be non-empty
	//   - max      - maximal file size in bytes
	//   - min      - minimal file size in bytes
	//
	// Constraints:
	//   - type:[mime-type]
	Attribute_IMAGE Attribute_Type = 11
	// A file value as a URL to eolympusercontent.com.
	//
	// Validation:
	//   - required - field must be non-empty
	//   - max      - maximal file size in bytes
	//   - min      - minimal file size in bytes
	//
	// Constraints:
	//   - type:[mime-type]
	Attribute_FILE Attribute_Type = 12
)

// Enum value maps for Attribute_Type.
var (
	Attribute_Type_name = map[int32]string{
		0:  "UNKNOWN_TYPE",
		1:  "STRING",
		2:  "TEXT",
		3:  "NUMBER",
		4:  "CHOICE",
		5:  "DATE",
		6:  "EMAIL",
		7:  "CHECKBOX",
		8:  "COUNTRY",
		9:  "REGION",
		10: "INSTITUTION",
		11: "IMAGE",
		12: "FILE",
	}
	Attribute_Type_value = map[string]int32{
		"UNKNOWN_TYPE": 0,
		"STRING":       1,
		"TEXT":         2,
		"NUMBER":       3,
		"CHOICE":       4,
		"DATE":         5,
		"EMAIL":        6,
		"CHECKBOX":     7,
		"COUNTRY":      8,
		"REGION":       9,
		"INSTITUTION":  10,
		"IMAGE":        11,
		"FILE":         12,
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
	return file_eolymp_community_attribute_proto_enumTypes[1].Descriptor()
}

func (Attribute_Type) Type() protoreflect.EnumType {
	return &file_eolymp_community_attribute_proto_enumTypes[1]
}

func (x Attribute_Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Attribute_Type.Descriptor instead.
func (Attribute_Type) EnumDescriptor() ([]byte, []int) {
	return file_eolymp_community_attribute_proto_rawDescGZIP(), []int{0, 1}
}

type Attribute_Visibility int32

const (
	Attribute_UNKNOWN_VISIBILITY Attribute_Visibility = 0 // unused
	Attribute_PRIVATE            Attribute_Visibility = 1 // visible only to administrators and member itself
	Attribute_PUBLIC             Attribute_Visibility = 2 // visible to everyone
	Attribute_INTERNAL           Attribute_Visibility = 3 // visible only to administrators
)

// Enum value maps for Attribute_Visibility.
var (
	Attribute_Visibility_name = map[int32]string{
		0: "UNKNOWN_VISIBILITY",
		1: "PRIVATE",
		2: "PUBLIC",
		3: "INTERNAL",
	}
	Attribute_Visibility_value = map[string]int32{
		"UNKNOWN_VISIBILITY": 0,
		"PRIVATE":            1,
		"PUBLIC":             2,
		"INTERNAL":           3,
	}
)

func (x Attribute_Visibility) Enum() *Attribute_Visibility {
	p := new(Attribute_Visibility)
	*p = x
	return p
}

func (x Attribute_Visibility) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Attribute_Visibility) Descriptor() protoreflect.EnumDescriptor {
	return file_eolymp_community_attribute_proto_enumTypes[2].Descriptor()
}

func (Attribute_Visibility) Type() protoreflect.EnumType {
	return &file_eolymp_community_attribute_proto_enumTypes[2]
}

func (x Attribute_Visibility) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Attribute_Visibility.Descriptor instead.
func (Attribute_Visibility) EnumDescriptor() ([]byte, []int) {
	return file_eolymp_community_attribute_proto_rawDescGZIP(), []int{0, 2}
}

type Attribute struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         string               `protobuf:"bytes,4,opt,name=id,proto3" json:"id,omitempty"`
	Key        string               `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`                                          // unique field identifier (should be unique within the form, not globally unique)
	Label      string               `protobuf:"bytes,2,opt,name=label,proto3" json:"label,omitempty"`                                      // label, normally displayed above the field
	Help       string               `protobuf:"bytes,3,opt,name=help,proto3" json:"help,omitempty"`                                        // help message, normally displayed right below the field
	Type       Attribute_Type       `protobuf:"varint,20,opt,name=type,proto3,enum=eolymp.community.Attribute_Type" json:"type,omitempty"` // type of the field
	Index      uint32               `protobuf:"varint,21,opt,name=index,proto3" json:"index,omitempty"`
	Required   bool                 `protobuf:"varint,31,opt,name=required,proto3" json:"required,omitempty"` // value is required (see field types for details)
	Readonly   bool                 `protobuf:"varint,33,opt,name=readonly,proto3" json:"readonly,omitempty"` // only administrator can change the value (member can be set during registration)
	Hidden     bool                 `protobuf:"varint,32,opt,name=hidden,proto3" json:"hidden,omitempty"`     // field is hidden (only visible to admin), deprecated, use visibility instead
	Visibility Attribute_Visibility `protobuf:"varint,34,opt,name=visibility,proto3,enum=eolymp.community.Attribute_Visibility" json:"visibility,omitempty"`
	// validation
	Regexp      string   `protobuf:"bytes,100,opt,name=regexp,proto3" json:"regexp,omitempty"`           // regexp validation (see field types for details)
	Min         int32    `protobuf:"varint,101,opt,name=min,proto3" json:"min,omitempty"`                // min value validation (see field types for details)
	Max         int32    `protobuf:"varint,102,opt,name=max,proto3" json:"max,omitempty"`                // max value validation (see field types for details)
	Choices     []string `protobuf:"bytes,103,rep,name=choices,proto3" json:"choices,omitempty"`         // possible choices validation (see field types for details)
	Country     string   `protobuf:"bytes,104,opt,name=country,proto3" json:"country,omitempty"`         // restrict region selector to a specific country
	Constraints []string `protobuf:"bytes,105,rep,name=constraints,proto3" json:"constraints,omitempty"` // additional constraints
}

func (x *Attribute) Reset() {
	*x = Attribute{}
	mi := &file_eolymp_community_attribute_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Attribute) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Attribute) ProtoMessage() {}

func (x *Attribute) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_community_attribute_proto_msgTypes[0]
	if x != nil {
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

func (x *Attribute) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Attribute) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *Attribute) GetLabel() string {
	if x != nil {
		return x.Label
	}
	return ""
}

func (x *Attribute) GetHelp() string {
	if x != nil {
		return x.Help
	}
	return ""
}

func (x *Attribute) GetType() Attribute_Type {
	if x != nil {
		return x.Type
	}
	return Attribute_UNKNOWN_TYPE
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

func (x *Attribute) GetReadonly() bool {
	if x != nil {
		return x.Readonly
	}
	return false
}

func (x *Attribute) GetHidden() bool {
	if x != nil {
		return x.Hidden
	}
	return false
}

func (x *Attribute) GetVisibility() Attribute_Visibility {
	if x != nil {
		return x.Visibility
	}
	return Attribute_UNKNOWN_VISIBILITY
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

func (x *Attribute) GetConstraints() []string {
	if x != nil {
		return x.Constraints
	}
	return nil
}

// Description provides localized information about field.
type Attribute_Description struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Locale  string   `protobuf:"bytes,2,opt,name=locale,proto3" json:"locale,omitempty"`   // locale
	Label   string   `protobuf:"bytes,3,opt,name=label,proto3" json:"label,omitempty"`     // label, normally displayed above the field
	Help    string   `protobuf:"bytes,4,opt,name=help,proto3" json:"help,omitempty"`       // help message, normally displayed right below the field
	Choices []string `protobuf:"bytes,5,rep,name=choices,proto3" json:"choices,omitempty"` // translation for choices (must be in the same order)
}

func (x *Attribute_Description) Reset() {
	*x = Attribute_Description{}
	mi := &file_eolymp_community_attribute_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Attribute_Description) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Attribute_Description) ProtoMessage() {}

func (x *Attribute_Description) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_community_attribute_proto_msgTypes[1]
	if x != nil {
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

// Attribute value
type Attribute_Value struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AttributeKey  string         `protobuf:"bytes,1,opt,name=attribute_key,json=attributeKey,proto3" json:"attribute_key,omitempty"`
	AttributeType Attribute_Type `protobuf:"varint,2,opt,name=attribute_type,json=attributeType,proto3,enum=eolymp.community.Attribute_Type" json:"attribute_type,omitempty"`
	// Types that are assignable to Value:
	//
	//	*Attribute_Value_String_
	//	*Attribute_Value_Number
	Value isAttribute_Value_Value `protobuf_oneof:"value"`
}

func (x *Attribute_Value) Reset() {
	*x = Attribute_Value{}
	mi := &file_eolymp_community_attribute_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Attribute_Value) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Attribute_Value) ProtoMessage() {}

func (x *Attribute_Value) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_community_attribute_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Attribute_Value.ProtoReflect.Descriptor instead.
func (*Attribute_Value) Descriptor() ([]byte, []int) {
	return file_eolymp_community_attribute_proto_rawDescGZIP(), []int{0, 1}
}

func (x *Attribute_Value) GetAttributeKey() string {
	if x != nil {
		return x.AttributeKey
	}
	return ""
}

func (x *Attribute_Value) GetAttributeType() Attribute_Type {
	if x != nil {
		return x.AttributeType
	}
	return Attribute_UNKNOWN_TYPE
}

func (m *Attribute_Value) GetValue() isAttribute_Value_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (x *Attribute_Value) GetString_() string {
	if x, ok := x.GetValue().(*Attribute_Value_String_); ok {
		return x.String_
	}
	return ""
}

func (x *Attribute_Value) GetNumber() int32 {
	if x, ok := x.GetValue().(*Attribute_Value_Number); ok {
		return x.Number
	}
	return 0
}

type isAttribute_Value_Value interface {
	isAttribute_Value_Value()
}

type Attribute_Value_String_ struct {
	String_ string `protobuf:"bytes,10,opt,name=string,proto3,oneof"`
}

type Attribute_Value_Number struct {
	Number int32 `protobuf:"varint,11,opt,name=number,proto3,oneof"`
}

func (*Attribute_Value_String_) isAttribute_Value_Value() {}

func (*Attribute_Value_Number) isAttribute_Value_Value() {}

var File_eolymp_community_attribute_proto protoreflect.FileDescriptor

var file_eolymp_community_attribute_proto_rawDesc = []byte{
	0x0a, 0x20, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69,
	0x74, 0x79, 0x2f, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x10, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x75,
	0x6e, 0x69, 0x74, 0x79, 0x22, 0xd8, 0x09, 0x0a, 0x09, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75,
	0x74, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x65,
	0x6c, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x65, 0x6c, 0x70, 0x12, 0x34,
	0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x14, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x20, 0x2e, 0x65,
	0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x2e,
	0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x15, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65,
	0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x18, 0x1f, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x72, 0x65,
	0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x61, 0x64, 0x6f, 0x6e,
	0x6c, 0x79, 0x18, 0x21, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x72, 0x65, 0x61, 0x64, 0x6f, 0x6e,
	0x6c, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x68, 0x69, 0x64, 0x64, 0x65, 0x6e, 0x18, 0x20, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x06, 0x68, 0x69, 0x64, 0x64, 0x65, 0x6e, 0x12, 0x46, 0x0a, 0x0a, 0x76, 0x69,
	0x73, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x18, 0x22, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x26,
	0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74,
	0x79, 0x2e, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x2e, 0x56, 0x69, 0x73, 0x69,
	0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x52, 0x0a, 0x76, 0x69, 0x73, 0x69, 0x62, 0x69, 0x6c, 0x69,
	0x74, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x67, 0x65, 0x78, 0x70, 0x18, 0x64, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x67, 0x65, 0x78, 0x70, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x69,
	0x6e, 0x18, 0x65, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x6d, 0x69, 0x6e, 0x12, 0x10, 0x0a, 0x03,
	0x6d, 0x61, 0x78, 0x18, 0x66, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x6d, 0x61, 0x78, 0x12, 0x18,
	0x0a, 0x07, 0x63, 0x68, 0x6f, 0x69, 0x63, 0x65, 0x73, 0x18, 0x67, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x07, 0x63, 0x68, 0x6f, 0x69, 0x63, 0x65, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x72, 0x79, 0x18, 0x68, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x72, 0x79, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x6f, 0x6e, 0x73, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x74,
	0x73, 0x18, 0x69, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6f, 0x6e, 0x73, 0x74, 0x72, 0x61,
	0x69, 0x6e, 0x74, 0x73, 0x1a, 0x69, 0x0a, 0x0b, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6c,
	0x61, 0x62, 0x65, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x61, 0x62, 0x65,
	0x6c, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x65, 0x6c, 0x70, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x68, 0x65, 0x6c, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x68, 0x6f, 0x69, 0x63, 0x65, 0x73,
	0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x63, 0x68, 0x6f, 0x69, 0x63, 0x65, 0x73, 0x1a,
	0xb2, 0x01, 0x0a, 0x05, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x61, 0x74, 0x74,
	0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0c, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x4b, 0x65, 0x79, 0x12, 0x47,
	0x0a, 0x0e, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x20, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x2e, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62,
	0x75, 0x74, 0x65, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0d, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62,
	0x75, 0x74, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x06, 0x73, 0x74, 0x72, 0x69, 0x6e,
	0x67, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x06, 0x73, 0x74, 0x72, 0x69, 0x6e,
	0x67, 0x12, 0x18, 0x0a, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x0b, 0x20, 0x01, 0x28,
	0x05, 0x48, 0x00, 0x52, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x42, 0x07, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x22, 0xf0, 0x01, 0x0a, 0x05, 0x50, 0x61, 0x74, 0x63, 0x68, 0x12, 0x0d,
	0x0a, 0x09, 0x50, 0x41, 0x54, 0x43, 0x48, 0x5f, 0x41, 0x4c, 0x4c, 0x10, 0x00, 0x12, 0x0f, 0x0a,
	0x0b, 0x50, 0x41, 0x54, 0x43, 0x48, 0x5f, 0x4c, 0x41, 0x42, 0x45, 0x4c, 0x10, 0x01, 0x12, 0x0e,
	0x0a, 0x0a, 0x50, 0x41, 0x54, 0x43, 0x48, 0x5f, 0x48, 0x45, 0x4c, 0x50, 0x10, 0x02, 0x12, 0x0e,
	0x0a, 0x0a, 0x50, 0x41, 0x54, 0x43, 0x48, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x10, 0x03, 0x12, 0x0f,
	0x0a, 0x0b, 0x50, 0x41, 0x54, 0x43, 0x48, 0x5f, 0x49, 0x4e, 0x44, 0x45, 0x58, 0x10, 0x04, 0x12,
	0x12, 0x0a, 0x0e, 0x50, 0x41, 0x54, 0x43, 0x48, 0x5f, 0x52, 0x45, 0x51, 0x55, 0x49, 0x52, 0x45,
	0x44, 0x10, 0x05, 0x12, 0x12, 0x0a, 0x0e, 0x50, 0x41, 0x54, 0x43, 0x48, 0x5f, 0x52, 0x45, 0x41,
	0x44, 0x4f, 0x4e, 0x4c, 0x59, 0x10, 0x06, 0x12, 0x14, 0x0a, 0x10, 0x50, 0x41, 0x54, 0x43, 0x48,
	0x5f, 0x56, 0x49, 0x53, 0x49, 0x42, 0x49, 0x4c, 0x49, 0x54, 0x59, 0x10, 0x07, 0x12, 0x10, 0x0a,
	0x0c, 0x50, 0x41, 0x54, 0x43, 0x48, 0x5f, 0x52, 0x45, 0x47, 0x45, 0x58, 0x50, 0x10, 0x08, 0x12,
	0x0d, 0x0a, 0x09, 0x50, 0x41, 0x54, 0x43, 0x48, 0x5f, 0x4d, 0x49, 0x4e, 0x10, 0x09, 0x12, 0x0d,
	0x0a, 0x09, 0x50, 0x41, 0x54, 0x43, 0x48, 0x5f, 0x4d, 0x41, 0x58, 0x10, 0x0a, 0x12, 0x11, 0x0a,
	0x0d, 0x50, 0x41, 0x54, 0x43, 0x48, 0x5f, 0x43, 0x48, 0x4f, 0x49, 0x43, 0x45, 0x53, 0x10, 0x0b,
	0x12, 0x15, 0x0a, 0x11, 0x50, 0x41, 0x54, 0x43, 0x48, 0x5f, 0x43, 0x4f, 0x4e, 0x53, 0x54, 0x52,
	0x41, 0x49, 0x4e, 0x54, 0x53, 0x10, 0x0c, 0x22, 0xa8, 0x01, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x10, 0x0a, 0x0c, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45,
	0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x53, 0x54, 0x52, 0x49, 0x4e, 0x47, 0x10, 0x01, 0x12, 0x08,
	0x0a, 0x04, 0x54, 0x45, 0x58, 0x54, 0x10, 0x02, 0x12, 0x0a, 0x0a, 0x06, 0x4e, 0x55, 0x4d, 0x42,
	0x45, 0x52, 0x10, 0x03, 0x12, 0x0a, 0x0a, 0x06, 0x43, 0x48, 0x4f, 0x49, 0x43, 0x45, 0x10, 0x04,
	0x12, 0x08, 0x0a, 0x04, 0x44, 0x41, 0x54, 0x45, 0x10, 0x05, 0x12, 0x09, 0x0a, 0x05, 0x45, 0x4d,
	0x41, 0x49, 0x4c, 0x10, 0x06, 0x12, 0x0c, 0x0a, 0x08, 0x43, 0x48, 0x45, 0x43, 0x4b, 0x42, 0x4f,
	0x58, 0x10, 0x07, 0x12, 0x0b, 0x0a, 0x07, 0x43, 0x4f, 0x55, 0x4e, 0x54, 0x52, 0x59, 0x10, 0x08,
	0x12, 0x0a, 0x0a, 0x06, 0x52, 0x45, 0x47, 0x49, 0x4f, 0x4e, 0x10, 0x09, 0x12, 0x0f, 0x0a, 0x0b,
	0x49, 0x4e, 0x53, 0x54, 0x49, 0x54, 0x55, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x0a, 0x12, 0x09, 0x0a,
	0x05, 0x49, 0x4d, 0x41, 0x47, 0x45, 0x10, 0x0b, 0x12, 0x08, 0x0a, 0x04, 0x46, 0x49, 0x4c, 0x45,
	0x10, 0x0c, 0x22, 0x4b, 0x0a, 0x0a, 0x56, 0x69, 0x73, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79,
	0x12, 0x16, 0x0a, 0x12, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x5f, 0x56, 0x49, 0x53, 0x49,
	0x42, 0x49, 0x4c, 0x49, 0x54, 0x59, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x50, 0x52, 0x49, 0x56,
	0x41, 0x54, 0x45, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x50, 0x55, 0x42, 0x4c, 0x49, 0x43, 0x10,
	0x02, 0x12, 0x0c, 0x0a, 0x08, 0x49, 0x4e, 0x54, 0x45, 0x52, 0x4e, 0x41, 0x4c, 0x10, 0x03, 0x42,
	0x35, 0x5a, 0x33, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6f,
	0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x67, 0x6f, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x65, 0x6f, 0x6c, 0x79,
	0x6d, 0x70, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x3b, 0x63, 0x6f, 0x6d,
	0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
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

var file_eolymp_community_attribute_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_eolymp_community_attribute_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_eolymp_community_attribute_proto_goTypes = []any{
	(Attribute_Patch)(0),          // 0: eolymp.community.Attribute.Patch
	(Attribute_Type)(0),           // 1: eolymp.community.Attribute.Type
	(Attribute_Visibility)(0),     // 2: eolymp.community.Attribute.Visibility
	(*Attribute)(nil),             // 3: eolymp.community.Attribute
	(*Attribute_Description)(nil), // 4: eolymp.community.Attribute.Description
	(*Attribute_Value)(nil),       // 5: eolymp.community.Attribute.Value
}
var file_eolymp_community_attribute_proto_depIdxs = []int32{
	1, // 0: eolymp.community.Attribute.type:type_name -> eolymp.community.Attribute.Type
	2, // 1: eolymp.community.Attribute.visibility:type_name -> eolymp.community.Attribute.Visibility
	1, // 2: eolymp.community.Attribute.Value.attribute_type:type_name -> eolymp.community.Attribute.Type
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_eolymp_community_attribute_proto_init() }
func file_eolymp_community_attribute_proto_init() {
	if File_eolymp_community_attribute_proto != nil {
		return
	}
	file_eolymp_community_attribute_proto_msgTypes[2].OneofWrappers = []any{
		(*Attribute_Value_String_)(nil),
		(*Attribute_Value_Number)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_eolymp_community_attribute_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   3,
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
