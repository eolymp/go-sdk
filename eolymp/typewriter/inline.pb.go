// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.18.1
// source: eolymp/typewriter/inline.proto

package typewriter

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

type Inline struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Node:
	//	*Inline_Text_
	//	*Inline_Link_
	//	*Inline_Math_
	//	*Inline_Code_
	//	*Inline_Variable_
	Node isInline_Node `protobuf_oneof:"node"`
}

func (x *Inline) Reset() {
	*x = Inline{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_typewriter_inline_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Inline) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Inline) ProtoMessage() {}

func (x *Inline) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_typewriter_inline_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Inline.ProtoReflect.Descriptor instead.
func (*Inline) Descriptor() ([]byte, []int) {
	return file_eolymp_typewriter_inline_proto_rawDescGZIP(), []int{0}
}

func (m *Inline) GetNode() isInline_Node {
	if m != nil {
		return m.Node
	}
	return nil
}

func (x *Inline) GetText() *Inline_Text {
	if x, ok := x.GetNode().(*Inline_Text_); ok {
		return x.Text
	}
	return nil
}

func (x *Inline) GetLink() *Inline_Link {
	if x, ok := x.GetNode().(*Inline_Link_); ok {
		return x.Link
	}
	return nil
}

func (x *Inline) GetMath() *Inline_Math {
	if x, ok := x.GetNode().(*Inline_Math_); ok {
		return x.Math
	}
	return nil
}

func (x *Inline) GetCode() *Inline_Code {
	if x, ok := x.GetNode().(*Inline_Code_); ok {
		return x.Code
	}
	return nil
}

func (x *Inline) GetVariable() *Inline_Variable {
	if x, ok := x.GetNode().(*Inline_Variable_); ok {
		return x.Variable
	}
	return nil
}

type isInline_Node interface {
	isInline_Node()
}

type Inline_Text_ struct {
	Text *Inline_Text `protobuf:"bytes,1,opt,name=text,proto3,oneof"`
}

type Inline_Link_ struct {
	Link *Inline_Link `protobuf:"bytes,2,opt,name=link,proto3,oneof"`
}

type Inline_Math_ struct {
	Math *Inline_Math `protobuf:"bytes,3,opt,name=math,proto3,oneof"`
}

type Inline_Code_ struct {
	Code *Inline_Code `protobuf:"bytes,4,opt,name=code,proto3,oneof"`
}

type Inline_Variable_ struct {
	Variable *Inline_Variable `protobuf:"bytes,5,opt,name=variable,proto3,oneof"`
}

func (*Inline_Text_) isInline_Node() {}

func (*Inline_Link_) isInline_Node() {}

func (*Inline_Math_) isInline_Node() {}

func (*Inline_Code_) isInline_Node() {}

func (*Inline_Variable_) isInline_Node() {}

type Inline_Style struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Bold          bool `protobuf:"varint,1,opt,name=bold,proto3" json:"bold,omitempty"`
	Italic        bool `protobuf:"varint,2,opt,name=italic,proto3" json:"italic,omitempty"`
	Underline     bool `protobuf:"varint,3,opt,name=underline,proto3" json:"underline,omitempty"`
	Strikethrough bool `protobuf:"varint,4,opt,name=strikethrough,proto3" json:"strikethrough,omitempty"`
}

func (x *Inline_Style) Reset() {
	*x = Inline_Style{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_typewriter_inline_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Inline_Style) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Inline_Style) ProtoMessage() {}

func (x *Inline_Style) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_typewriter_inline_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Inline_Style.ProtoReflect.Descriptor instead.
func (*Inline_Style) Descriptor() ([]byte, []int) {
	return file_eolymp_typewriter_inline_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Inline_Style) GetBold() bool {
	if x != nil {
		return x.Bold
	}
	return false
}

func (x *Inline_Style) GetItalic() bool {
	if x != nil {
		return x.Italic
	}
	return false
}

func (x *Inline_Style) GetUnderline() bool {
	if x != nil {
		return x.Underline
	}
	return false
}

func (x *Inline_Style) GetStrikethrough() bool {
	if x != nil {
		return x.Strikethrough
	}
	return false
}

type Inline_Text struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Str   string        `protobuf:"bytes,1,opt,name=str,proto3" json:"str,omitempty"`
	Style *Inline_Style `protobuf:"bytes,2,opt,name=style,proto3" json:"style,omitempty"`
}

func (x *Inline_Text) Reset() {
	*x = Inline_Text{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_typewriter_inline_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Inline_Text) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Inline_Text) ProtoMessage() {}

func (x *Inline_Text) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_typewriter_inline_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Inline_Text.ProtoReflect.Descriptor instead.
func (*Inline_Text) Descriptor() ([]byte, []int) {
	return file_eolymp_typewriter_inline_proto_rawDescGZIP(), []int{0, 1}
}

func (x *Inline_Text) GetStr() string {
	if x != nil {
		return x.Str
	}
	return ""
}

func (x *Inline_Text) GetStyle() *Inline_Style {
	if x != nil {
		return x.Style
	}
	return nil
}

type Inline_Link struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url      string    `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	Children []*Inline `protobuf:"bytes,10,rep,name=children,proto3" json:"children,omitempty"`
}

func (x *Inline_Link) Reset() {
	*x = Inline_Link{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_typewriter_inline_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Inline_Link) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Inline_Link) ProtoMessage() {}

func (x *Inline_Link) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_typewriter_inline_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Inline_Link.ProtoReflect.Descriptor instead.
func (*Inline_Link) Descriptor() ([]byte, []int) {
	return file_eolymp_typewriter_inline_proto_rawDescGZIP(), []int{0, 2}
}

func (x *Inline_Link) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *Inline_Link) GetChildren() []*Inline {
	if x != nil {
		return x.Children
	}
	return nil
}

type Inline_Math struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Content string `protobuf:"bytes,1,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *Inline_Math) Reset() {
	*x = Inline_Math{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_typewriter_inline_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Inline_Math) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Inline_Math) ProtoMessage() {}

func (x *Inline_Math) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_typewriter_inline_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Inline_Math.ProtoReflect.Descriptor instead.
func (*Inline_Math) Descriptor() ([]byte, []int) {
	return file_eolymp_typewriter_inline_proto_rawDescGZIP(), []int{0, 3}
}

func (x *Inline_Math) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

type Inline_Code struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Content string `protobuf:"bytes,1,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *Inline_Code) Reset() {
	*x = Inline_Code{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_typewriter_inline_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Inline_Code) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Inline_Code) ProtoMessage() {}

func (x *Inline_Code) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_typewriter_inline_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Inline_Code.ProtoReflect.Descriptor instead.
func (*Inline_Code) Descriptor() ([]byte, []int) {
	return file_eolymp_typewriter_inline_proto_rawDescGZIP(), []int{0, 4}
}

func (x *Inline_Code) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

type Inline_Variable struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name  string        `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Style *Inline_Style `protobuf:"bytes,2,opt,name=style,proto3" json:"style,omitempty"`
}

func (x *Inline_Variable) Reset() {
	*x = Inline_Variable{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_typewriter_inline_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Inline_Variable) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Inline_Variable) ProtoMessage() {}

func (x *Inline_Variable) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_typewriter_inline_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Inline_Variable.ProtoReflect.Descriptor instead.
func (*Inline_Variable) Descriptor() ([]byte, []int) {
	return file_eolymp_typewriter_inline_proto_rawDescGZIP(), []int{0, 5}
}

func (x *Inline_Variable) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Inline_Variable) GetStyle() *Inline_Style {
	if x != nil {
		return x.Style
	}
	return nil
}

var File_eolymp_typewriter_inline_proto protoreflect.FileDescriptor

var file_eolymp_typewriter_inline_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x77, 0x72, 0x69,
	0x74, 0x65, 0x72, 0x2f, 0x69, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x11, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x77, 0x72, 0x69,
	0x74, 0x65, 0x72, 0x22, 0xe0, 0x05, 0x0a, 0x06, 0x49, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x12, 0x34,
	0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x65,
	0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x77, 0x72, 0x69, 0x74, 0x65, 0x72,
	0x2e, 0x49, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x54, 0x65, 0x78, 0x74, 0x48, 0x00, 0x52, 0x04,
	0x74, 0x65, 0x78, 0x74, 0x12, 0x34, 0x0a, 0x04, 0x6c, 0x69, 0x6e, 0x6b, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x74, 0x79, 0x70, 0x65,
	0x77, 0x72, 0x69, 0x74, 0x65, 0x72, 0x2e, 0x49, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x4c, 0x69,
	0x6e, 0x6b, 0x48, 0x00, 0x52, 0x04, 0x6c, 0x69, 0x6e, 0x6b, 0x12, 0x34, 0x0a, 0x04, 0x6d, 0x61,
	0x74, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d,
	0x70, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x77, 0x72, 0x69, 0x74, 0x65, 0x72, 0x2e, 0x49, 0x6e, 0x6c,
	0x69, 0x6e, 0x65, 0x2e, 0x4d, 0x61, 0x74, 0x68, 0x48, 0x00, 0x52, 0x04, 0x6d, 0x61, 0x74, 0x68,
	0x12, 0x34, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e,
	0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x77, 0x72, 0x69, 0x74,
	0x65, 0x72, 0x2e, 0x49, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x48, 0x00,
	0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x40, 0x0a, 0x08, 0x76, 0x61, 0x72, 0x69, 0x61, 0x62,
	0x6c, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d,
	0x70, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x77, 0x72, 0x69, 0x74, 0x65, 0x72, 0x2e, 0x49, 0x6e, 0x6c,
	0x69, 0x6e, 0x65, 0x2e, 0x56, 0x61, 0x72, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x48, 0x00, 0x52, 0x08,
	0x76, 0x61, 0x72, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x1a, 0x77, 0x0a, 0x05, 0x53, 0x74, 0x79, 0x6c,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x6f, 0x6c, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x04, 0x62, 0x6f, 0x6c, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x69, 0x74, 0x61, 0x6c, 0x69, 0x63, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x69, 0x74, 0x61, 0x6c, 0x69, 0x63, 0x12, 0x1c, 0x0a,
	0x09, 0x75, 0x6e, 0x64, 0x65, 0x72, 0x6c, 0x69, 0x6e, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x09, 0x75, 0x6e, 0x64, 0x65, 0x72, 0x6c, 0x69, 0x6e, 0x65, 0x12, 0x24, 0x0a, 0x0d, 0x73,
	0x74, 0x72, 0x69, 0x6b, 0x65, 0x74, 0x68, 0x72, 0x6f, 0x75, 0x67, 0x68, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x0d, 0x73, 0x74, 0x72, 0x69, 0x6b, 0x65, 0x74, 0x68, 0x72, 0x6f, 0x75, 0x67,
	0x68, 0x1a, 0x4f, 0x0a, 0x04, 0x54, 0x65, 0x78, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x74, 0x72,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x73, 0x74, 0x72, 0x12, 0x35, 0x0a, 0x05, 0x73,
	0x74, 0x79, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x65, 0x6f, 0x6c,
	0x79, 0x6d, 0x70, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x77, 0x72, 0x69, 0x74, 0x65, 0x72, 0x2e, 0x49,
	0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x53, 0x74, 0x79, 0x6c, 0x65, 0x52, 0x05, 0x73, 0x74, 0x79,
	0x6c, 0x65, 0x1a, 0x4f, 0x0a, 0x04, 0x4c, 0x69, 0x6e, 0x6b, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72,
	0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x35, 0x0a, 0x08,
	0x63, 0x68, 0x69, 0x6c, 0x64, 0x72, 0x65, 0x6e, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19,
	0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x77, 0x72, 0x69, 0x74,
	0x65, 0x72, 0x2e, 0x49, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x52, 0x08, 0x63, 0x68, 0x69, 0x6c, 0x64,
	0x72, 0x65, 0x6e, 0x1a, 0x20, 0x0a, 0x04, 0x4d, 0x61, 0x74, 0x68, 0x12, 0x18, 0x0a, 0x07, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x1a, 0x20, 0x0a, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x1a, 0x55, 0x0a, 0x08, 0x56, 0x61, 0x72, 0x69, 0x61,
	0x62, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x35, 0x0a, 0x05, 0x73, 0x74, 0x79, 0x6c, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e,
	0x74, 0x79, 0x70, 0x65, 0x77, 0x72, 0x69, 0x74, 0x65, 0x72, 0x2e, 0x49, 0x6e, 0x6c, 0x69, 0x6e,
	0x65, 0x2e, 0x53, 0x74, 0x79, 0x6c, 0x65, 0x52, 0x05, 0x73, 0x74, 0x79, 0x6c, 0x65, 0x42, 0x06,
	0x0a, 0x04, 0x6e, 0x6f, 0x64, 0x65, 0x42, 0x37, 0x5a, 0x35, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x67, 0x6f, 0x2d, 0x73,
	0x64, 0x6b, 0x2f, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x77, 0x72,
	0x69, 0x74, 0x65, 0x72, 0x3b, 0x74, 0x79, 0x70, 0x65, 0x77, 0x72, 0x69, 0x74, 0x65, 0x72, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_eolymp_typewriter_inline_proto_rawDescOnce sync.Once
	file_eolymp_typewriter_inline_proto_rawDescData = file_eolymp_typewriter_inline_proto_rawDesc
)

func file_eolymp_typewriter_inline_proto_rawDescGZIP() []byte {
	file_eolymp_typewriter_inline_proto_rawDescOnce.Do(func() {
		file_eolymp_typewriter_inline_proto_rawDescData = protoimpl.X.CompressGZIP(file_eolymp_typewriter_inline_proto_rawDescData)
	})
	return file_eolymp_typewriter_inline_proto_rawDescData
}

var file_eolymp_typewriter_inline_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_eolymp_typewriter_inline_proto_goTypes = []interface{}{
	(*Inline)(nil),          // 0: eolymp.typewriter.Inline
	(*Inline_Style)(nil),    // 1: eolymp.typewriter.Inline.Style
	(*Inline_Text)(nil),     // 2: eolymp.typewriter.Inline.Text
	(*Inline_Link)(nil),     // 3: eolymp.typewriter.Inline.Link
	(*Inline_Math)(nil),     // 4: eolymp.typewriter.Inline.Math
	(*Inline_Code)(nil),     // 5: eolymp.typewriter.Inline.Code
	(*Inline_Variable)(nil), // 6: eolymp.typewriter.Inline.Variable
}
var file_eolymp_typewriter_inline_proto_depIdxs = []int32{
	2, // 0: eolymp.typewriter.Inline.text:type_name -> eolymp.typewriter.Inline.Text
	3, // 1: eolymp.typewriter.Inline.link:type_name -> eolymp.typewriter.Inline.Link
	4, // 2: eolymp.typewriter.Inline.math:type_name -> eolymp.typewriter.Inline.Math
	5, // 3: eolymp.typewriter.Inline.code:type_name -> eolymp.typewriter.Inline.Code
	6, // 4: eolymp.typewriter.Inline.variable:type_name -> eolymp.typewriter.Inline.Variable
	1, // 5: eolymp.typewriter.Inline.Text.style:type_name -> eolymp.typewriter.Inline.Style
	0, // 6: eolymp.typewriter.Inline.Link.children:type_name -> eolymp.typewriter.Inline
	1, // 7: eolymp.typewriter.Inline.Variable.style:type_name -> eolymp.typewriter.Inline.Style
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_eolymp_typewriter_inline_proto_init() }
func file_eolymp_typewriter_inline_proto_init() {
	if File_eolymp_typewriter_inline_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_eolymp_typewriter_inline_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Inline); i {
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
		file_eolymp_typewriter_inline_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Inline_Style); i {
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
		file_eolymp_typewriter_inline_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Inline_Text); i {
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
		file_eolymp_typewriter_inline_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Inline_Link); i {
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
		file_eolymp_typewriter_inline_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Inline_Math); i {
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
		file_eolymp_typewriter_inline_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Inline_Code); i {
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
		file_eolymp_typewriter_inline_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Inline_Variable); i {
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
	file_eolymp_typewriter_inline_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*Inline_Text_)(nil),
		(*Inline_Link_)(nil),
		(*Inline_Math_)(nil),
		(*Inline_Code_)(nil),
		(*Inline_Variable_)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_eolymp_typewriter_inline_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_eolymp_typewriter_inline_proto_goTypes,
		DependencyIndexes: file_eolymp_typewriter_inline_proto_depIdxs,
		MessageInfos:      file_eolymp_typewriter_inline_proto_msgTypes,
	}.Build()
	File_eolymp_typewriter_inline_proto = out.File
	file_eolymp_typewriter_inline_proto_rawDesc = nil
	file_eolymp_typewriter_inline_proto_goTypes = nil
	file_eolymp_typewriter_inline_proto_depIdxs = nil
}
