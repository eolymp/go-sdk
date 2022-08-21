// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.18.1
// source: eolymp/wellknown/expression.proto

package wellknown

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ExpressionID_Type int32

const (
	ExpressionID_NONE      ExpressionID_Type = 0
	ExpressionID_EQUAL     ExpressionID_Type = 1
	ExpressionID_NOT_EQUAL ExpressionID_Type = 2
)

// Enum value maps for ExpressionID_Type.
var (
	ExpressionID_Type_name = map[int32]string{
		0: "NONE",
		1: "EQUAL",
		2: "NOT_EQUAL",
	}
	ExpressionID_Type_value = map[string]int32{
		"NONE":      0,
		"EQUAL":     1,
		"NOT_EQUAL": 2,
	}
)

func (x ExpressionID_Type) Enum() *ExpressionID_Type {
	p := new(ExpressionID_Type)
	*p = x
	return p
}

func (x ExpressionID_Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ExpressionID_Type) Descriptor() protoreflect.EnumDescriptor {
	return file_eolymp_wellknown_expression_proto_enumTypes[0].Descriptor()
}

func (ExpressionID_Type) Type() protoreflect.EnumType {
	return &file_eolymp_wellknown_expression_proto_enumTypes[0]
}

func (x ExpressionID_Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ExpressionID_Type.Descriptor instead.
func (ExpressionID_Type) EnumDescriptor() ([]byte, []int) {
	return file_eolymp_wellknown_expression_proto_rawDescGZIP(), []int{0, 0}
}

type ExpressionInt_Type int32

const (
	ExpressionInt_NONE               ExpressionInt_Type = 0
	ExpressionInt_EQUAL              ExpressionInt_Type = 1
	ExpressionInt_NOT_EQUAL          ExpressionInt_Type = 2
	ExpressionInt_GREATER_THAN       ExpressionInt_Type = 3
	ExpressionInt_GREATER_THAN_EQUAL ExpressionInt_Type = 4
	ExpressionInt_LESS_THAN          ExpressionInt_Type = 5
	ExpressionInt_LESS_THAN_EQUAL    ExpressionInt_Type = 6
)

// Enum value maps for ExpressionInt_Type.
var (
	ExpressionInt_Type_name = map[int32]string{
		0: "NONE",
		1: "EQUAL",
		2: "NOT_EQUAL",
		3: "GREATER_THAN",
		4: "GREATER_THAN_EQUAL",
		5: "LESS_THAN",
		6: "LESS_THAN_EQUAL",
	}
	ExpressionInt_Type_value = map[string]int32{
		"NONE":               0,
		"EQUAL":              1,
		"NOT_EQUAL":          2,
		"GREATER_THAN":       3,
		"GREATER_THAN_EQUAL": 4,
		"LESS_THAN":          5,
		"LESS_THAN_EQUAL":    6,
	}
)

func (x ExpressionInt_Type) Enum() *ExpressionInt_Type {
	p := new(ExpressionInt_Type)
	*p = x
	return p
}

func (x ExpressionInt_Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ExpressionInt_Type) Descriptor() protoreflect.EnumDescriptor {
	return file_eolymp_wellknown_expression_proto_enumTypes[1].Descriptor()
}

func (ExpressionInt_Type) Type() protoreflect.EnumType {
	return &file_eolymp_wellknown_expression_proto_enumTypes[1]
}

func (x ExpressionInt_Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ExpressionInt_Type.Descriptor instead.
func (ExpressionInt_Type) EnumDescriptor() ([]byte, []int) {
	return file_eolymp_wellknown_expression_proto_rawDescGZIP(), []int{1, 0}
}

type ExpressionFloat_Type int32

const (
	ExpressionFloat_NONE               ExpressionFloat_Type = 0
	ExpressionFloat_EQUAL              ExpressionFloat_Type = 1
	ExpressionFloat_NOT_EQUAL          ExpressionFloat_Type = 2
	ExpressionFloat_GREATER_THAN       ExpressionFloat_Type = 3
	ExpressionFloat_GREATER_THAN_EQUAL ExpressionFloat_Type = 4
	ExpressionFloat_LESS_THAN          ExpressionFloat_Type = 5
	ExpressionFloat_LESS_THAN_EQUAL    ExpressionFloat_Type = 6
)

// Enum value maps for ExpressionFloat_Type.
var (
	ExpressionFloat_Type_name = map[int32]string{
		0: "NONE",
		1: "EQUAL",
		2: "NOT_EQUAL",
		3: "GREATER_THAN",
		4: "GREATER_THAN_EQUAL",
		5: "LESS_THAN",
		6: "LESS_THAN_EQUAL",
	}
	ExpressionFloat_Type_value = map[string]int32{
		"NONE":               0,
		"EQUAL":              1,
		"NOT_EQUAL":          2,
		"GREATER_THAN":       3,
		"GREATER_THAN_EQUAL": 4,
		"LESS_THAN":          5,
		"LESS_THAN_EQUAL":    6,
	}
)

func (x ExpressionFloat_Type) Enum() *ExpressionFloat_Type {
	p := new(ExpressionFloat_Type)
	*p = x
	return p
}

func (x ExpressionFloat_Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ExpressionFloat_Type) Descriptor() protoreflect.EnumDescriptor {
	return file_eolymp_wellknown_expression_proto_enumTypes[2].Descriptor()
}

func (ExpressionFloat_Type) Type() protoreflect.EnumType {
	return &file_eolymp_wellknown_expression_proto_enumTypes[2]
}

func (x ExpressionFloat_Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ExpressionFloat_Type.Descriptor instead.
func (ExpressionFloat_Type) EnumDescriptor() ([]byte, []int) {
	return file_eolymp_wellknown_expression_proto_rawDescGZIP(), []int{2, 0}
}

type ExpressionString_Type int32

const (
	ExpressionString_NONE       ExpressionString_Type = 0
	ExpressionString_EQUAL      ExpressionString_Type = 1
	ExpressionString_NOT_EQUAL  ExpressionString_Type = 2
	ExpressionString_CONTAINING ExpressionString_Type = 3
	ExpressionString_STARTING   ExpressionString_Type = 4
)

// Enum value maps for ExpressionString_Type.
var (
	ExpressionString_Type_name = map[int32]string{
		0: "NONE",
		1: "EQUAL",
		2: "NOT_EQUAL",
		3: "CONTAINING",
		4: "STARTING",
	}
	ExpressionString_Type_value = map[string]int32{
		"NONE":       0,
		"EQUAL":      1,
		"NOT_EQUAL":  2,
		"CONTAINING": 3,
		"STARTING":   4,
	}
)

func (x ExpressionString_Type) Enum() *ExpressionString_Type {
	p := new(ExpressionString_Type)
	*p = x
	return p
}

func (x ExpressionString_Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ExpressionString_Type) Descriptor() protoreflect.EnumDescriptor {
	return file_eolymp_wellknown_expression_proto_enumTypes[3].Descriptor()
}

func (ExpressionString_Type) Type() protoreflect.EnumType {
	return &file_eolymp_wellknown_expression_proto_enumTypes[3]
}

func (x ExpressionString_Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ExpressionString_Type.Descriptor instead.
func (ExpressionString_Type) EnumDescriptor() ([]byte, []int) {
	return file_eolymp_wellknown_expression_proto_rawDescGZIP(), []int{3, 0}
}

type ExpressionEnum_Type int32

const (
	ExpressionEnum_NONE      ExpressionEnum_Type = 0
	ExpressionEnum_EQUAL     ExpressionEnum_Type = 1
	ExpressionEnum_NOT_EQUAL ExpressionEnum_Type = 2
)

// Enum value maps for ExpressionEnum_Type.
var (
	ExpressionEnum_Type_name = map[int32]string{
		0: "NONE",
		1: "EQUAL",
		2: "NOT_EQUAL",
	}
	ExpressionEnum_Type_value = map[string]int32{
		"NONE":      0,
		"EQUAL":     1,
		"NOT_EQUAL": 2,
	}
)

func (x ExpressionEnum_Type) Enum() *ExpressionEnum_Type {
	p := new(ExpressionEnum_Type)
	*p = x
	return p
}

func (x ExpressionEnum_Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ExpressionEnum_Type) Descriptor() protoreflect.EnumDescriptor {
	return file_eolymp_wellknown_expression_proto_enumTypes[4].Descriptor()
}

func (ExpressionEnum_Type) Type() protoreflect.EnumType {
	return &file_eolymp_wellknown_expression_proto_enumTypes[4]
}

func (x ExpressionEnum_Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ExpressionEnum_Type.Descriptor instead.
func (ExpressionEnum_Type) EnumDescriptor() ([]byte, []int) {
	return file_eolymp_wellknown_expression_proto_rawDescGZIP(), []int{4, 0}
}

type ExpressionBool_Type int32

const (
	ExpressionBool_NONE  ExpressionBool_Type = 0
	ExpressionBool_EQUAL ExpressionBool_Type = 1
)

// Enum value maps for ExpressionBool_Type.
var (
	ExpressionBool_Type_name = map[int32]string{
		0: "NONE",
		1: "EQUAL",
	}
	ExpressionBool_Type_value = map[string]int32{
		"NONE":  0,
		"EQUAL": 1,
	}
)

func (x ExpressionBool_Type) Enum() *ExpressionBool_Type {
	p := new(ExpressionBool_Type)
	*p = x
	return p
}

func (x ExpressionBool_Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ExpressionBool_Type) Descriptor() protoreflect.EnumDescriptor {
	return file_eolymp_wellknown_expression_proto_enumTypes[5].Descriptor()
}

func (ExpressionBool_Type) Type() protoreflect.EnumType {
	return &file_eolymp_wellknown_expression_proto_enumTypes[5]
}

func (x ExpressionBool_Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ExpressionBool_Type.Descriptor instead.
func (ExpressionBool_Type) EnumDescriptor() ([]byte, []int) {
	return file_eolymp_wellknown_expression_proto_rawDescGZIP(), []int{5, 0}
}

type ExpressionTimestamp_Type int32

const (
	ExpressionTimestamp_NONE               ExpressionTimestamp_Type = 0
	ExpressionTimestamp_EQUAL              ExpressionTimestamp_Type = 1
	ExpressionTimestamp_NOT_EQUAL          ExpressionTimestamp_Type = 2
	ExpressionTimestamp_GREATER_THAN       ExpressionTimestamp_Type = 3
	ExpressionTimestamp_GREATER_THAN_EQUAL ExpressionTimestamp_Type = 4
	ExpressionTimestamp_LESS_THAN          ExpressionTimestamp_Type = 5
	ExpressionTimestamp_LESS_THAN_EQUAL    ExpressionTimestamp_Type = 6
)

// Enum value maps for ExpressionTimestamp_Type.
var (
	ExpressionTimestamp_Type_name = map[int32]string{
		0: "NONE",
		1: "EQUAL",
		2: "NOT_EQUAL",
		3: "GREATER_THAN",
		4: "GREATER_THAN_EQUAL",
		5: "LESS_THAN",
		6: "LESS_THAN_EQUAL",
	}
	ExpressionTimestamp_Type_value = map[string]int32{
		"NONE":               0,
		"EQUAL":              1,
		"NOT_EQUAL":          2,
		"GREATER_THAN":       3,
		"GREATER_THAN_EQUAL": 4,
		"LESS_THAN":          5,
		"LESS_THAN_EQUAL":    6,
	}
)

func (x ExpressionTimestamp_Type) Enum() *ExpressionTimestamp_Type {
	p := new(ExpressionTimestamp_Type)
	*p = x
	return p
}

func (x ExpressionTimestamp_Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ExpressionTimestamp_Type) Descriptor() protoreflect.EnumDescriptor {
	return file_eolymp_wellknown_expression_proto_enumTypes[6].Descriptor()
}

func (ExpressionTimestamp_Type) Type() protoreflect.EnumType {
	return &file_eolymp_wellknown_expression_proto_enumTypes[6]
}

func (x ExpressionTimestamp_Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ExpressionTimestamp_Type.Descriptor instead.
func (ExpressionTimestamp_Type) EnumDescriptor() ([]byte, []int) {
	return file_eolymp_wellknown_expression_proto_rawDescGZIP(), []int{6, 0}
}

// ExpressionID specifies match criteria for unique identifier
type ExpressionID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Is    ExpressionID_Type `protobuf:"varint,1,opt,name=is,proto3,enum=eolymp.wellknown.ExpressionID_Type" json:"is,omitempty"`
	Value string            `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *ExpressionID) Reset() {
	*x = ExpressionID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_wellknown_expression_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExpressionID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExpressionID) ProtoMessage() {}

func (x *ExpressionID) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_wellknown_expression_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExpressionID.ProtoReflect.Descriptor instead.
func (*ExpressionID) Descriptor() ([]byte, []int) {
	return file_eolymp_wellknown_expression_proto_rawDescGZIP(), []int{0}
}

func (x *ExpressionID) GetIs() ExpressionID_Type {
	if x != nil {
		return x.Is
	}
	return ExpressionID_NONE
}

func (x *ExpressionID) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

// ExpressionInt specifies match criteria for an integer
type ExpressionInt struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Is    ExpressionInt_Type `protobuf:"varint,1,opt,name=is,proto3,enum=eolymp.wellknown.ExpressionInt_Type" json:"is,omitempty"`
	Value int32              `protobuf:"varint,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *ExpressionInt) Reset() {
	*x = ExpressionInt{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_wellknown_expression_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExpressionInt) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExpressionInt) ProtoMessage() {}

func (x *ExpressionInt) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_wellknown_expression_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExpressionInt.ProtoReflect.Descriptor instead.
func (*ExpressionInt) Descriptor() ([]byte, []int) {
	return file_eolymp_wellknown_expression_proto_rawDescGZIP(), []int{1}
}

func (x *ExpressionInt) GetIs() ExpressionInt_Type {
	if x != nil {
		return x.Is
	}
	return ExpressionInt_NONE
}

func (x *ExpressionInt) GetValue() int32 {
	if x != nil {
		return x.Value
	}
	return 0
}

// ExpressionFloat specifies match criteria for an floateger
type ExpressionFloat struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Is    ExpressionFloat_Type `protobuf:"varint,1,opt,name=is,proto3,enum=eolymp.wellknown.ExpressionFloat_Type" json:"is,omitempty"`
	Value float32              `protobuf:"fixed32,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *ExpressionFloat) Reset() {
	*x = ExpressionFloat{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_wellknown_expression_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExpressionFloat) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExpressionFloat) ProtoMessage() {}

func (x *ExpressionFloat) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_wellknown_expression_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExpressionFloat.ProtoReflect.Descriptor instead.
func (*ExpressionFloat) Descriptor() ([]byte, []int) {
	return file_eolymp_wellknown_expression_proto_rawDescGZIP(), []int{2}
}

func (x *ExpressionFloat) GetIs() ExpressionFloat_Type {
	if x != nil {
		return x.Is
	}
	return ExpressionFloat_NONE
}

func (x *ExpressionFloat) GetValue() float32 {
	if x != nil {
		return x.Value
	}
	return 0
}

// ExpressionString specifies match criteria for a string
type ExpressionString struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Is    ExpressionString_Type `protobuf:"varint,1,opt,name=is,proto3,enum=eolymp.wellknown.ExpressionString_Type" json:"is,omitempty"`
	Value string                `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *ExpressionString) Reset() {
	*x = ExpressionString{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_wellknown_expression_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExpressionString) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExpressionString) ProtoMessage() {}

func (x *ExpressionString) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_wellknown_expression_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExpressionString.ProtoReflect.Descriptor instead.
func (*ExpressionString) Descriptor() ([]byte, []int) {
	return file_eolymp_wellknown_expression_proto_rawDescGZIP(), []int{3}
}

func (x *ExpressionString) GetIs() ExpressionString_Type {
	if x != nil {
		return x.Is
	}
	return ExpressionString_NONE
}

func (x *ExpressionString) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

// ExpressionEnum specifies match criteria for an enumeration (unlike string does not allow partial match)
type ExpressionEnum struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Is    ExpressionEnum_Type `protobuf:"varint,1,opt,name=is,proto3,enum=eolymp.wellknown.ExpressionEnum_Type" json:"is,omitempty"`
	Value string              `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *ExpressionEnum) Reset() {
	*x = ExpressionEnum{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_wellknown_expression_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExpressionEnum) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExpressionEnum) ProtoMessage() {}

func (x *ExpressionEnum) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_wellknown_expression_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExpressionEnum.ProtoReflect.Descriptor instead.
func (*ExpressionEnum) Descriptor() ([]byte, []int) {
	return file_eolymp_wellknown_expression_proto_rawDescGZIP(), []int{4}
}

func (x *ExpressionEnum) GetIs() ExpressionEnum_Type {
	if x != nil {
		return x.Is
	}
	return ExpressionEnum_NONE
}

func (x *ExpressionEnum) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

// ExpressionBool specifies match criteria for a boolean
type ExpressionBool struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Is    ExpressionBool_Type `protobuf:"varint,1,opt,name=is,proto3,enum=eolymp.wellknown.ExpressionBool_Type" json:"is,omitempty"`
	Value bool                `protobuf:"varint,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *ExpressionBool) Reset() {
	*x = ExpressionBool{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_wellknown_expression_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExpressionBool) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExpressionBool) ProtoMessage() {}

func (x *ExpressionBool) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_wellknown_expression_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExpressionBool.ProtoReflect.Descriptor instead.
func (*ExpressionBool) Descriptor() ([]byte, []int) {
	return file_eolymp_wellknown_expression_proto_rawDescGZIP(), []int{5}
}

func (x *ExpressionBool) GetIs() ExpressionBool_Type {
	if x != nil {
		return x.Is
	}
	return ExpressionBool_NONE
}

func (x *ExpressionBool) GetValue() bool {
	if x != nil {
		return x.Value
	}
	return false
}

// ExpressionTimestamp specifies match criteria for a timestamp
type ExpressionTimestamp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Is    ExpressionTimestamp_Type `protobuf:"varint,1,opt,name=is,proto3,enum=eolymp.wellknown.ExpressionTimestamp_Type" json:"is,omitempty"`
	Value *timestamppb.Timestamp   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *ExpressionTimestamp) Reset() {
	*x = ExpressionTimestamp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_wellknown_expression_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExpressionTimestamp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExpressionTimestamp) ProtoMessage() {}

func (x *ExpressionTimestamp) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_wellknown_expression_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExpressionTimestamp.ProtoReflect.Descriptor instead.
func (*ExpressionTimestamp) Descriptor() ([]byte, []int) {
	return file_eolymp_wellknown_expression_proto_rawDescGZIP(), []int{6}
}

func (x *ExpressionTimestamp) GetIs() ExpressionTimestamp_Type {
	if x != nil {
		return x.Is
	}
	return ExpressionTimestamp_NONE
}

func (x *ExpressionTimestamp) GetValue() *timestamppb.Timestamp {
	if x != nil {
		return x.Value
	}
	return nil
}

var File_eolymp_wellknown_expression_proto protoreflect.FileDescriptor

var file_eolymp_wellknown_expression_proto_rawDesc = []byte{
	0x0a, 0x21, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x77, 0x65, 0x6c, 0x6c, 0x6b, 0x6e, 0x6f,
	0x77, 0x6e, 0x2f, 0x65, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x10, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x77, 0x65, 0x6c, 0x6c,
	0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x85, 0x01, 0x0a, 0x0c, 0x45, 0x78, 0x70, 0x72, 0x65,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x12, 0x33, 0x0a, 0x02, 0x69, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x23, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x77, 0x65, 0x6c,
	0x6c, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x2e, 0x45, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x49, 0x44, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x02, 0x69, 0x73, 0x12, 0x14, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x22, 0x2a, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x4f,
	0x4e, 0x45, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x45, 0x51, 0x55, 0x41, 0x4c, 0x10, 0x01, 0x12,
	0x0d, 0x0a, 0x09, 0x4e, 0x4f, 0x54, 0x5f, 0x45, 0x51, 0x55, 0x41, 0x4c, 0x10, 0x02, 0x22, 0xd5,
	0x01, 0x0a, 0x0d, 0x45, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x74,
	0x12, 0x34, 0x0a, 0x02, 0x69, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x24, 0x2e, 0x65,
	0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x77, 0x65, 0x6c, 0x6c, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x2e,
	0x45, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x74, 0x2e, 0x54, 0x79,
	0x70, 0x65, 0x52, 0x02, 0x69, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x78, 0x0a, 0x04,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x09,
	0x0a, 0x05, 0x45, 0x51, 0x55, 0x41, 0x4c, 0x10, 0x01, 0x12, 0x0d, 0x0a, 0x09, 0x4e, 0x4f, 0x54,
	0x5f, 0x45, 0x51, 0x55, 0x41, 0x4c, 0x10, 0x02, 0x12, 0x10, 0x0a, 0x0c, 0x47, 0x52, 0x45, 0x41,
	0x54, 0x45, 0x52, 0x5f, 0x54, 0x48, 0x41, 0x4e, 0x10, 0x03, 0x12, 0x16, 0x0a, 0x12, 0x47, 0x52,
	0x45, 0x41, 0x54, 0x45, 0x52, 0x5f, 0x54, 0x48, 0x41, 0x4e, 0x5f, 0x45, 0x51, 0x55, 0x41, 0x4c,
	0x10, 0x04, 0x12, 0x0d, 0x0a, 0x09, 0x4c, 0x45, 0x53, 0x53, 0x5f, 0x54, 0x48, 0x41, 0x4e, 0x10,
	0x05, 0x12, 0x13, 0x0a, 0x0f, 0x4c, 0x45, 0x53, 0x53, 0x5f, 0x54, 0x48, 0x41, 0x4e, 0x5f, 0x45,
	0x51, 0x55, 0x41, 0x4c, 0x10, 0x06, 0x22, 0xd9, 0x01, 0x0a, 0x0f, 0x45, 0x78, 0x70, 0x72, 0x65,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x46, 0x6c, 0x6f, 0x61, 0x74, 0x12, 0x36, 0x0a, 0x02, 0x69, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x26, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e,
	0x77, 0x65, 0x6c, 0x6c, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x2e, 0x45, 0x78, 0x70, 0x72, 0x65, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x46, 0x6c, 0x6f, 0x61, 0x74, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x02,
	0x69, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x02, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x78, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x08, 0x0a, 0x04, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x45, 0x51,
	0x55, 0x41, 0x4c, 0x10, 0x01, 0x12, 0x0d, 0x0a, 0x09, 0x4e, 0x4f, 0x54, 0x5f, 0x45, 0x51, 0x55,
	0x41, 0x4c, 0x10, 0x02, 0x12, 0x10, 0x0a, 0x0c, 0x47, 0x52, 0x45, 0x41, 0x54, 0x45, 0x52, 0x5f,
	0x54, 0x48, 0x41, 0x4e, 0x10, 0x03, 0x12, 0x16, 0x0a, 0x12, 0x47, 0x52, 0x45, 0x41, 0x54, 0x45,
	0x52, 0x5f, 0x54, 0x48, 0x41, 0x4e, 0x5f, 0x45, 0x51, 0x55, 0x41, 0x4c, 0x10, 0x04, 0x12, 0x0d,
	0x0a, 0x09, 0x4c, 0x45, 0x53, 0x53, 0x5f, 0x54, 0x48, 0x41, 0x4e, 0x10, 0x05, 0x12, 0x13, 0x0a,
	0x0f, 0x4c, 0x45, 0x53, 0x53, 0x5f, 0x54, 0x48, 0x41, 0x4e, 0x5f, 0x45, 0x51, 0x55, 0x41, 0x4c,
	0x10, 0x06, 0x22, 0xab, 0x01, 0x0a, 0x10, 0x45, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x12, 0x37, 0x0a, 0x02, 0x69, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x27, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x77, 0x65, 0x6c,
	0x6c, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x2e, 0x45, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x02, 0x69, 0x73,
	0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x48, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x08,
	0x0a, 0x04, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x45, 0x51, 0x55, 0x41,
	0x4c, 0x10, 0x01, 0x12, 0x0d, 0x0a, 0x09, 0x4e, 0x4f, 0x54, 0x5f, 0x45, 0x51, 0x55, 0x41, 0x4c,
	0x10, 0x02, 0x12, 0x0e, 0x0a, 0x0a, 0x43, 0x4f, 0x4e, 0x54, 0x41, 0x49, 0x4e, 0x49, 0x4e, 0x47,
	0x10, 0x03, 0x12, 0x0c, 0x0a, 0x08, 0x53, 0x54, 0x41, 0x52, 0x54, 0x49, 0x4e, 0x47, 0x10, 0x04,
	0x22, 0x89, 0x01, 0x0a, 0x0e, 0x45, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x45,
	0x6e, 0x75, 0x6d, 0x12, 0x35, 0x0a, 0x02, 0x69, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x25, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x77, 0x65, 0x6c, 0x6c, 0x6b, 0x6e, 0x6f,
	0x77, 0x6e, 0x2e, 0x45, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x45, 0x6e, 0x75,
	0x6d, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x02, 0x69, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x22, 0x2a, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x4f, 0x4e, 0x45,
	0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x45, 0x51, 0x55, 0x41, 0x4c, 0x10, 0x01, 0x12, 0x0d, 0x0a,
	0x09, 0x4e, 0x4f, 0x54, 0x5f, 0x45, 0x51, 0x55, 0x41, 0x4c, 0x10, 0x02, 0x22, 0x7a, 0x0a, 0x0e,
	0x45, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x42, 0x6f, 0x6f, 0x6c, 0x12, 0x35,
	0x0a, 0x02, 0x69, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x25, 0x2e, 0x65, 0x6f, 0x6c,
	0x79, 0x6d, 0x70, 0x2e, 0x77, 0x65, 0x6c, 0x6c, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x2e, 0x45, 0x78,
	0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x42, 0x6f, 0x6f, 0x6c, 0x2e, 0x54, 0x79, 0x70,
	0x65, 0x52, 0x02, 0x69, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x1b, 0x0a, 0x04, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x09, 0x0a,
	0x05, 0x45, 0x51, 0x55, 0x41, 0x4c, 0x10, 0x01, 0x22, 0xfd, 0x01, 0x0a, 0x13, 0x45, 0x78, 0x70,
	0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x12, 0x3a, 0x0a, 0x02, 0x69, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2a, 0x2e, 0x65,
	0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x77, 0x65, 0x6c, 0x6c, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x2e,
	0x45, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x02, 0x69, 0x73, 0x12, 0x30, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x78,
	0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00,
	0x12, 0x09, 0x0a, 0x05, 0x45, 0x51, 0x55, 0x41, 0x4c, 0x10, 0x01, 0x12, 0x0d, 0x0a, 0x09, 0x4e,
	0x4f, 0x54, 0x5f, 0x45, 0x51, 0x55, 0x41, 0x4c, 0x10, 0x02, 0x12, 0x10, 0x0a, 0x0c, 0x47, 0x52,
	0x45, 0x41, 0x54, 0x45, 0x52, 0x5f, 0x54, 0x48, 0x41, 0x4e, 0x10, 0x03, 0x12, 0x16, 0x0a, 0x12,
	0x47, 0x52, 0x45, 0x41, 0x54, 0x45, 0x52, 0x5f, 0x54, 0x48, 0x41, 0x4e, 0x5f, 0x45, 0x51, 0x55,
	0x41, 0x4c, 0x10, 0x04, 0x12, 0x0d, 0x0a, 0x09, 0x4c, 0x45, 0x53, 0x53, 0x5f, 0x54, 0x48, 0x41,
	0x4e, 0x10, 0x05, 0x12, 0x13, 0x0a, 0x0f, 0x4c, 0x45, 0x53, 0x53, 0x5f, 0x54, 0x48, 0x41, 0x4e,
	0x5f, 0x45, 0x51, 0x55, 0x41, 0x4c, 0x10, 0x06, 0x42, 0x35, 0x5a, 0x33, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x67, 0x6f,
	0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x77, 0x65, 0x6c, 0x6c,
	0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x3b, 0x77, 0x65, 0x6c, 0x6c, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_eolymp_wellknown_expression_proto_rawDescOnce sync.Once
	file_eolymp_wellknown_expression_proto_rawDescData = file_eolymp_wellknown_expression_proto_rawDesc
)

func file_eolymp_wellknown_expression_proto_rawDescGZIP() []byte {
	file_eolymp_wellknown_expression_proto_rawDescOnce.Do(func() {
		file_eolymp_wellknown_expression_proto_rawDescData = protoimpl.X.CompressGZIP(file_eolymp_wellknown_expression_proto_rawDescData)
	})
	return file_eolymp_wellknown_expression_proto_rawDescData
}

var file_eolymp_wellknown_expression_proto_enumTypes = make([]protoimpl.EnumInfo, 7)
var file_eolymp_wellknown_expression_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_eolymp_wellknown_expression_proto_goTypes = []interface{}{
	(ExpressionID_Type)(0),        // 0: eolymp.wellknown.ExpressionID.Type
	(ExpressionInt_Type)(0),       // 1: eolymp.wellknown.ExpressionInt.Type
	(ExpressionFloat_Type)(0),     // 2: eolymp.wellknown.ExpressionFloat.Type
	(ExpressionString_Type)(0),    // 3: eolymp.wellknown.ExpressionString.Type
	(ExpressionEnum_Type)(0),      // 4: eolymp.wellknown.ExpressionEnum.Type
	(ExpressionBool_Type)(0),      // 5: eolymp.wellknown.ExpressionBool.Type
	(ExpressionTimestamp_Type)(0), // 6: eolymp.wellknown.ExpressionTimestamp.Type
	(*ExpressionID)(nil),          // 7: eolymp.wellknown.ExpressionID
	(*ExpressionInt)(nil),         // 8: eolymp.wellknown.ExpressionInt
	(*ExpressionFloat)(nil),       // 9: eolymp.wellknown.ExpressionFloat
	(*ExpressionString)(nil),      // 10: eolymp.wellknown.ExpressionString
	(*ExpressionEnum)(nil),        // 11: eolymp.wellknown.ExpressionEnum
	(*ExpressionBool)(nil),        // 12: eolymp.wellknown.ExpressionBool
	(*ExpressionTimestamp)(nil),   // 13: eolymp.wellknown.ExpressionTimestamp
	(*timestamppb.Timestamp)(nil), // 14: google.protobuf.Timestamp
}
var file_eolymp_wellknown_expression_proto_depIdxs = []int32{
	0,  // 0: eolymp.wellknown.ExpressionID.is:type_name -> eolymp.wellknown.ExpressionID.Type
	1,  // 1: eolymp.wellknown.ExpressionInt.is:type_name -> eolymp.wellknown.ExpressionInt.Type
	2,  // 2: eolymp.wellknown.ExpressionFloat.is:type_name -> eolymp.wellknown.ExpressionFloat.Type
	3,  // 3: eolymp.wellknown.ExpressionString.is:type_name -> eolymp.wellknown.ExpressionString.Type
	4,  // 4: eolymp.wellknown.ExpressionEnum.is:type_name -> eolymp.wellknown.ExpressionEnum.Type
	5,  // 5: eolymp.wellknown.ExpressionBool.is:type_name -> eolymp.wellknown.ExpressionBool.Type
	6,  // 6: eolymp.wellknown.ExpressionTimestamp.is:type_name -> eolymp.wellknown.ExpressionTimestamp.Type
	14, // 7: eolymp.wellknown.ExpressionTimestamp.value:type_name -> google.protobuf.Timestamp
	8,  // [8:8] is the sub-list for method output_type
	8,  // [8:8] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_eolymp_wellknown_expression_proto_init() }
func file_eolymp_wellknown_expression_proto_init() {
	if File_eolymp_wellknown_expression_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_eolymp_wellknown_expression_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExpressionID); i {
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
		file_eolymp_wellknown_expression_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExpressionInt); i {
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
		file_eolymp_wellknown_expression_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExpressionFloat); i {
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
		file_eolymp_wellknown_expression_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExpressionString); i {
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
		file_eolymp_wellknown_expression_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExpressionEnum); i {
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
		file_eolymp_wellknown_expression_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExpressionBool); i {
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
		file_eolymp_wellknown_expression_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExpressionTimestamp); i {
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
			RawDescriptor: file_eolymp_wellknown_expression_proto_rawDesc,
			NumEnums:      7,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_eolymp_wellknown_expression_proto_goTypes,
		DependencyIndexes: file_eolymp_wellknown_expression_proto_depIdxs,
		EnumInfos:         file_eolymp_wellknown_expression_proto_enumTypes,
		MessageInfos:      file_eolymp_wellknown_expression_proto_msgTypes,
	}.Build()
	File_eolymp_wellknown_expression_proto = out.File
	file_eolymp_wellknown_expression_proto_rawDesc = nil
	file_eolymp_wellknown_expression_proto_goTypes = nil
	file_eolymp_wellknown_expression_proto_depIdxs = nil
}
