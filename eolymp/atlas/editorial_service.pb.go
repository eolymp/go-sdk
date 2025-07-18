// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.29.4
// source: eolymp/atlas/editorial_service.proto

package atlas

import (
	_ "github.com/eolymp/go-sdk/eolymp/annotations"
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

type ListEditorialsInput struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// pagination
	Offset        int32             `protobuf:"varint,10,opt,name=offset,proto3" json:"offset,omitempty"`
	Size          int32             `protobuf:"varint,11,opt,name=size,proto3" json:"size,omitempty"`
	Version       uint32            `protobuf:"varint,100,opt,name=version,proto3" json:"version,omitempty"` // request data for specific problem version
	Render        bool              `protobuf:"varint,1,opt,name=render,proto3" json:"render,omitempty"`     // deprecated
	Extra         []Editorial_Extra `protobuf:"varint,1123,rep,packed,name=extra,proto3,enum=eolymp.atlas.Editorial_Extra" json:"extra,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListEditorialsInput) Reset() {
	*x = ListEditorialsInput{}
	mi := &file_eolymp_atlas_editorial_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListEditorialsInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListEditorialsInput) ProtoMessage() {}

func (x *ListEditorialsInput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_atlas_editorial_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListEditorialsInput.ProtoReflect.Descriptor instead.
func (*ListEditorialsInput) Descriptor() ([]byte, []int) {
	return file_eolymp_atlas_editorial_service_proto_rawDescGZIP(), []int{0}
}

func (x *ListEditorialsInput) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *ListEditorialsInput) GetSize() int32 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *ListEditorialsInput) GetVersion() uint32 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *ListEditorialsInput) GetRender() bool {
	if x != nil {
		return x.Render
	}
	return false
}

func (x *ListEditorialsInput) GetExtra() []Editorial_Extra {
	if x != nil {
		return x.Extra
	}
	return nil
}

type ListEditorialsOutput struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Total         int32                  `protobuf:"varint,1,opt,name=total,proto3" json:"total,omitempty"`
	Items         []*Editorial           `protobuf:"bytes,2,rep,name=items,proto3" json:"items,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListEditorialsOutput) Reset() {
	*x = ListEditorialsOutput{}
	mi := &file_eolymp_atlas_editorial_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListEditorialsOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListEditorialsOutput) ProtoMessage() {}

func (x *ListEditorialsOutput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_atlas_editorial_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListEditorialsOutput.ProtoReflect.Descriptor instead.
func (*ListEditorialsOutput) Descriptor() ([]byte, []int) {
	return file_eolymp_atlas_editorial_service_proto_rawDescGZIP(), []int{1}
}

func (x *ListEditorialsOutput) GetTotal() int32 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *ListEditorialsOutput) GetItems() []*Editorial {
	if x != nil {
		return x.Items
	}
	return nil
}

type DescribeEditorialInput struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	EditorialId   string                 `protobuf:"bytes,2,opt,name=editorial_id,json=editorialId,proto3" json:"editorial_id,omitempty"`
	Version       uint32                 `protobuf:"varint,100,opt,name=version,proto3" json:"version,omitempty"` // request data for specific problem version
	Render        bool                   `protobuf:"varint,3,opt,name=render,proto3" json:"render,omitempty"`     // deprecated
	Extra         []Editorial_Extra      `protobuf:"varint,1123,rep,packed,name=extra,proto3,enum=eolymp.atlas.Editorial_Extra" json:"extra,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DescribeEditorialInput) Reset() {
	*x = DescribeEditorialInput{}
	mi := &file_eolymp_atlas_editorial_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DescribeEditorialInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescribeEditorialInput) ProtoMessage() {}

func (x *DescribeEditorialInput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_atlas_editorial_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DescribeEditorialInput.ProtoReflect.Descriptor instead.
func (*DescribeEditorialInput) Descriptor() ([]byte, []int) {
	return file_eolymp_atlas_editorial_service_proto_rawDescGZIP(), []int{2}
}

func (x *DescribeEditorialInput) GetEditorialId() string {
	if x != nil {
		return x.EditorialId
	}
	return ""
}

func (x *DescribeEditorialInput) GetVersion() uint32 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *DescribeEditorialInput) GetRender() bool {
	if x != nil {
		return x.Render
	}
	return false
}

func (x *DescribeEditorialInput) GetExtra() []Editorial_Extra {
	if x != nil {
		return x.Extra
	}
	return nil
}

type DescribeEditorialOutput struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Editorial     *Editorial             `protobuf:"bytes,1,opt,name=editorial,proto3" json:"editorial,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DescribeEditorialOutput) Reset() {
	*x = DescribeEditorialOutput{}
	mi := &file_eolymp_atlas_editorial_service_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DescribeEditorialOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescribeEditorialOutput) ProtoMessage() {}

func (x *DescribeEditorialOutput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_atlas_editorial_service_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DescribeEditorialOutput.ProtoReflect.Descriptor instead.
func (*DescribeEditorialOutput) Descriptor() ([]byte, []int) {
	return file_eolymp_atlas_editorial_service_proto_rawDescGZIP(), []int{3}
}

func (x *DescribeEditorialOutput) GetEditorial() *Editorial {
	if x != nil {
		return x.Editorial
	}
	return nil
}

type LookupEditorialInput struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Locale        string                 `protobuf:"bytes,2,opt,name=locale,proto3" json:"locale,omitempty"`
	Version       uint32                 `protobuf:"varint,100,opt,name=version,proto3" json:"version,omitempty"` // request data for specific problem version
	Render        bool                   `protobuf:"varint,3,opt,name=render,proto3" json:"render,omitempty"`     // deprecated
	Extra         []Editorial_Extra      `protobuf:"varint,1123,rep,packed,name=extra,proto3,enum=eolymp.atlas.Editorial_Extra" json:"extra,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *LookupEditorialInput) Reset() {
	*x = LookupEditorialInput{}
	mi := &file_eolymp_atlas_editorial_service_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LookupEditorialInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LookupEditorialInput) ProtoMessage() {}

func (x *LookupEditorialInput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_atlas_editorial_service_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LookupEditorialInput.ProtoReflect.Descriptor instead.
func (*LookupEditorialInput) Descriptor() ([]byte, []int) {
	return file_eolymp_atlas_editorial_service_proto_rawDescGZIP(), []int{4}
}

func (x *LookupEditorialInput) GetLocale() string {
	if x != nil {
		return x.Locale
	}
	return ""
}

func (x *LookupEditorialInput) GetVersion() uint32 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *LookupEditorialInput) GetRender() bool {
	if x != nil {
		return x.Render
	}
	return false
}

func (x *LookupEditorialInput) GetExtra() []Editorial_Extra {
	if x != nil {
		return x.Extra
	}
	return nil
}

type LookupEditorialOutput struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Editorial     *Editorial             `protobuf:"bytes,1,opt,name=editorial,proto3" json:"editorial,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *LookupEditorialOutput) Reset() {
	*x = LookupEditorialOutput{}
	mi := &file_eolymp_atlas_editorial_service_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LookupEditorialOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LookupEditorialOutput) ProtoMessage() {}

func (x *LookupEditorialOutput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_atlas_editorial_service_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LookupEditorialOutput.ProtoReflect.Descriptor instead.
func (*LookupEditorialOutput) Descriptor() ([]byte, []int) {
	return file_eolymp_atlas_editorial_service_proto_rawDescGZIP(), []int{5}
}

func (x *LookupEditorialOutput) GetEditorial() *Editorial {
	if x != nil {
		return x.Editorial
	}
	return nil
}

type PreviewEditorialInput struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Editorial     *Editorial             `protobuf:"bytes,2,opt,name=editorial,proto3" json:"editorial,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PreviewEditorialInput) Reset() {
	*x = PreviewEditorialInput{}
	mi := &file_eolymp_atlas_editorial_service_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PreviewEditorialInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PreviewEditorialInput) ProtoMessage() {}

func (x *PreviewEditorialInput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_atlas_editorial_service_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PreviewEditorialInput.ProtoReflect.Descriptor instead.
func (*PreviewEditorialInput) Descriptor() ([]byte, []int) {
	return file_eolymp_atlas_editorial_service_proto_rawDescGZIP(), []int{6}
}

func (x *PreviewEditorialInput) GetEditorial() *Editorial {
	if x != nil {
		return x.Editorial
	}
	return nil
}

type PreviewEditorialOutput struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Editorial     *Editorial             `protobuf:"bytes,1,opt,name=editorial,proto3" json:"editorial,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PreviewEditorialOutput) Reset() {
	*x = PreviewEditorialOutput{}
	mi := &file_eolymp_atlas_editorial_service_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PreviewEditorialOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PreviewEditorialOutput) ProtoMessage() {}

func (x *PreviewEditorialOutput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_atlas_editorial_service_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PreviewEditorialOutput.ProtoReflect.Descriptor instead.
func (*PreviewEditorialOutput) Descriptor() ([]byte, []int) {
	return file_eolymp_atlas_editorial_service_proto_rawDescGZIP(), []int{7}
}

func (x *PreviewEditorialOutput) GetEditorial() *Editorial {
	if x != nil {
		return x.Editorial
	}
	return nil
}

type CreateEditorialInput struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Editorial     *Editorial             `protobuf:"bytes,2,opt,name=editorial,proto3" json:"editorial,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateEditorialInput) Reset() {
	*x = CreateEditorialInput{}
	mi := &file_eolymp_atlas_editorial_service_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateEditorialInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateEditorialInput) ProtoMessage() {}

func (x *CreateEditorialInput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_atlas_editorial_service_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateEditorialInput.ProtoReflect.Descriptor instead.
func (*CreateEditorialInput) Descriptor() ([]byte, []int) {
	return file_eolymp_atlas_editorial_service_proto_rawDescGZIP(), []int{8}
}

func (x *CreateEditorialInput) GetEditorial() *Editorial {
	if x != nil {
		return x.Editorial
	}
	return nil
}

type CreateEditorialOutput struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	EditorialId   string                 `protobuf:"bytes,1,opt,name=editorial_id,json=editorialId,proto3" json:"editorial_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateEditorialOutput) Reset() {
	*x = CreateEditorialOutput{}
	mi := &file_eolymp_atlas_editorial_service_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateEditorialOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateEditorialOutput) ProtoMessage() {}

func (x *CreateEditorialOutput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_atlas_editorial_service_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateEditorialOutput.ProtoReflect.Descriptor instead.
func (*CreateEditorialOutput) Descriptor() ([]byte, []int) {
	return file_eolymp_atlas_editorial_service_proto_rawDescGZIP(), []int{9}
}

func (x *CreateEditorialOutput) GetEditorialId() string {
	if x != nil {
		return x.EditorialId
	}
	return ""
}

type UpdateEditorialInput struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Patch         []Editorial_Patch      `protobuf:"varint,10,rep,packed,name=patch,proto3,enum=eolymp.atlas.Editorial_Patch" json:"patch,omitempty"`
	EditorialId   string                 `protobuf:"bytes,2,opt,name=editorial_id,json=editorialId,proto3" json:"editorial_id,omitempty"`
	Editorial     *Editorial             `protobuf:"bytes,3,opt,name=editorial,proto3" json:"editorial,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateEditorialInput) Reset() {
	*x = UpdateEditorialInput{}
	mi := &file_eolymp_atlas_editorial_service_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateEditorialInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateEditorialInput) ProtoMessage() {}

func (x *UpdateEditorialInput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_atlas_editorial_service_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateEditorialInput.ProtoReflect.Descriptor instead.
func (*UpdateEditorialInput) Descriptor() ([]byte, []int) {
	return file_eolymp_atlas_editorial_service_proto_rawDescGZIP(), []int{10}
}

func (x *UpdateEditorialInput) GetPatch() []Editorial_Patch {
	if x != nil {
		return x.Patch
	}
	return nil
}

func (x *UpdateEditorialInput) GetEditorialId() string {
	if x != nil {
		return x.EditorialId
	}
	return ""
}

func (x *UpdateEditorialInput) GetEditorial() *Editorial {
	if x != nil {
		return x.Editorial
	}
	return nil
}

type UpdateEditorialOutput struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateEditorialOutput) Reset() {
	*x = UpdateEditorialOutput{}
	mi := &file_eolymp_atlas_editorial_service_proto_msgTypes[11]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateEditorialOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateEditorialOutput) ProtoMessage() {}

func (x *UpdateEditorialOutput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_atlas_editorial_service_proto_msgTypes[11]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateEditorialOutput.ProtoReflect.Descriptor instead.
func (*UpdateEditorialOutput) Descriptor() ([]byte, []int) {
	return file_eolymp_atlas_editorial_service_proto_rawDescGZIP(), []int{11}
}

type DeleteEditorialInput struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	EditorialId   string                 `protobuf:"bytes,2,opt,name=editorial_id,json=editorialId,proto3" json:"editorial_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteEditorialInput) Reset() {
	*x = DeleteEditorialInput{}
	mi := &file_eolymp_atlas_editorial_service_proto_msgTypes[12]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteEditorialInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteEditorialInput) ProtoMessage() {}

func (x *DeleteEditorialInput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_atlas_editorial_service_proto_msgTypes[12]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteEditorialInput.ProtoReflect.Descriptor instead.
func (*DeleteEditorialInput) Descriptor() ([]byte, []int) {
	return file_eolymp_atlas_editorial_service_proto_rawDescGZIP(), []int{12}
}

func (x *DeleteEditorialInput) GetEditorialId() string {
	if x != nil {
		return x.EditorialId
	}
	return ""
}

type DeleteEditorialOutput struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteEditorialOutput) Reset() {
	*x = DeleteEditorialOutput{}
	mi := &file_eolymp_atlas_editorial_service_proto_msgTypes[13]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteEditorialOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteEditorialOutput) ProtoMessage() {}

func (x *DeleteEditorialOutput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_atlas_editorial_service_proto_msgTypes[13]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteEditorialOutput.ProtoReflect.Descriptor instead.
func (*DeleteEditorialOutput) Descriptor() ([]byte, []int) {
	return file_eolymp_atlas_editorial_service_proto_rawDescGZIP(), []int{13}
}

type TranslateEditorialsInput struct {
	state           protoimpl.MessageState `protogen:"open.v1"`
	Source          string                 `protobuf:"bytes,1,opt,name=source,proto3" json:"source,omitempty"`                                           // if empty, english translation will be used, if english translation is not available, first available translation will be used
	Target          []string               `protobuf:"bytes,2,rep,name=target,proto3" json:"target,omitempty"`                                           // list of target languages, if statement exists and it has automatic=true (or override_manual=true) it will be updated
	TargetAutomatic bool                   `protobuf:"varint,3,opt,name=target_automatic,json=targetAutomatic,proto3" json:"target_automatic,omitempty"` // add to targets all editorials with automatic=true
	OverrideManual  bool                   `protobuf:"varint,4,opt,name=override_manual,json=overrideManual,proto3" json:"override_manual,omitempty"`    // update editorials even if automatic=false
	unknownFields   protoimpl.UnknownFields
	sizeCache       protoimpl.SizeCache
}

func (x *TranslateEditorialsInput) Reset() {
	*x = TranslateEditorialsInput{}
	mi := &file_eolymp_atlas_editorial_service_proto_msgTypes[14]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TranslateEditorialsInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TranslateEditorialsInput) ProtoMessage() {}

func (x *TranslateEditorialsInput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_atlas_editorial_service_proto_msgTypes[14]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TranslateEditorialsInput.ProtoReflect.Descriptor instead.
func (*TranslateEditorialsInput) Descriptor() ([]byte, []int) {
	return file_eolymp_atlas_editorial_service_proto_rawDescGZIP(), []int{14}
}

func (x *TranslateEditorialsInput) GetSource() string {
	if x != nil {
		return x.Source
	}
	return ""
}

func (x *TranslateEditorialsInput) GetTarget() []string {
	if x != nil {
		return x.Target
	}
	return nil
}

func (x *TranslateEditorialsInput) GetTargetAutomatic() bool {
	if x != nil {
		return x.TargetAutomatic
	}
	return false
}

func (x *TranslateEditorialsInput) GetOverrideManual() bool {
	if x != nil {
		return x.OverrideManual
	}
	return false
}

type TranslateEditorialsOutput struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	JobId         string                 `protobuf:"bytes,1,opt,name=job_id,json=jobId,proto3" json:"job_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *TranslateEditorialsOutput) Reset() {
	*x = TranslateEditorialsOutput{}
	mi := &file_eolymp_atlas_editorial_service_proto_msgTypes[15]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TranslateEditorialsOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TranslateEditorialsOutput) ProtoMessage() {}

func (x *TranslateEditorialsOutput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_atlas_editorial_service_proto_msgTypes[15]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TranslateEditorialsOutput.ProtoReflect.Descriptor instead.
func (*TranslateEditorialsOutput) Descriptor() ([]byte, []int) {
	return file_eolymp_atlas_editorial_service_proto_rawDescGZIP(), []int{15}
}

func (x *TranslateEditorialsOutput) GetJobId() string {
	if x != nil {
		return x.JobId
	}
	return ""
}

var File_eolymp_atlas_editorial_service_proto protoreflect.FileDescriptor

const file_eolymp_atlas_editorial_service_proto_rawDesc = "" +
	"\n" +
	"$eolymp/atlas/editorial_service.proto\x12\feolymp.atlas\x1a\x1deolymp/annotations/http.proto\x1a\"eolymp/annotations/ratelimit.proto\x1a\x1eeolymp/annotations/scope.proto\x1a\x1ceolymp/atlas/editorial.proto\"\xa9\x01\n" +
	"\x13ListEditorialsInput\x12\x16\n" +
	"\x06offset\x18\n" +
	" \x01(\x05R\x06offset\x12\x12\n" +
	"\x04size\x18\v \x01(\x05R\x04size\x12\x18\n" +
	"\aversion\x18d \x01(\rR\aversion\x12\x16\n" +
	"\x06render\x18\x01 \x01(\bR\x06render\x124\n" +
	"\x05extra\x18\xe3\b \x03(\x0e2\x1d.eolymp.atlas.Editorial.ExtraR\x05extra\"[\n" +
	"\x14ListEditorialsOutput\x12\x14\n" +
	"\x05total\x18\x01 \x01(\x05R\x05total\x12-\n" +
	"\x05items\x18\x02 \x03(\v2\x17.eolymp.atlas.EditorialR\x05items\"\xa3\x01\n" +
	"\x16DescribeEditorialInput\x12!\n" +
	"\feditorial_id\x18\x02 \x01(\tR\veditorialId\x12\x18\n" +
	"\aversion\x18d \x01(\rR\aversion\x12\x16\n" +
	"\x06render\x18\x03 \x01(\bR\x06render\x124\n" +
	"\x05extra\x18\xe3\b \x03(\x0e2\x1d.eolymp.atlas.Editorial.ExtraR\x05extra\"P\n" +
	"\x17DescribeEditorialOutput\x125\n" +
	"\teditorial\x18\x01 \x01(\v2\x17.eolymp.atlas.EditorialR\teditorial\"\x96\x01\n" +
	"\x14LookupEditorialInput\x12\x16\n" +
	"\x06locale\x18\x02 \x01(\tR\x06locale\x12\x18\n" +
	"\aversion\x18d \x01(\rR\aversion\x12\x16\n" +
	"\x06render\x18\x03 \x01(\bR\x06render\x124\n" +
	"\x05extra\x18\xe3\b \x03(\x0e2\x1d.eolymp.atlas.Editorial.ExtraR\x05extra\"N\n" +
	"\x15LookupEditorialOutput\x125\n" +
	"\teditorial\x18\x01 \x01(\v2\x17.eolymp.atlas.EditorialR\teditorial\"N\n" +
	"\x15PreviewEditorialInput\x125\n" +
	"\teditorial\x18\x02 \x01(\v2\x17.eolymp.atlas.EditorialR\teditorial\"O\n" +
	"\x16PreviewEditorialOutput\x125\n" +
	"\teditorial\x18\x01 \x01(\v2\x17.eolymp.atlas.EditorialR\teditorial\"M\n" +
	"\x14CreateEditorialInput\x125\n" +
	"\teditorial\x18\x02 \x01(\v2\x17.eolymp.atlas.EditorialR\teditorial\":\n" +
	"\x15CreateEditorialOutput\x12!\n" +
	"\feditorial_id\x18\x01 \x01(\tR\veditorialId\"\xa5\x01\n" +
	"\x14UpdateEditorialInput\x123\n" +
	"\x05patch\x18\n" +
	" \x03(\x0e2\x1d.eolymp.atlas.Editorial.PatchR\x05patch\x12!\n" +
	"\feditorial_id\x18\x02 \x01(\tR\veditorialId\x125\n" +
	"\teditorial\x18\x03 \x01(\v2\x17.eolymp.atlas.EditorialR\teditorial\"\x17\n" +
	"\x15UpdateEditorialOutput\"9\n" +
	"\x14DeleteEditorialInput\x12!\n" +
	"\feditorial_id\x18\x02 \x01(\tR\veditorialId\"\x17\n" +
	"\x15DeleteEditorialOutput\"\x9e\x01\n" +
	"\x18TranslateEditorialsInput\x12\x16\n" +
	"\x06source\x18\x01 \x01(\tR\x06source\x12\x16\n" +
	"\x06target\x18\x02 \x03(\tR\x06target\x12)\n" +
	"\x10target_automatic\x18\x03 \x01(\bR\x0ftargetAutomatic\x12'\n" +
	"\x0foverride_manual\x18\x04 \x01(\bR\x0eoverrideManual\"2\n" +
	"\x19TranslateEditorialsOutput\x12\x15\n" +
	"\x06job_id\x18\x01 \x01(\tR\x05jobId2\xbd\n" +
	"\n" +
	"\x10EditorialService\x12\x99\x01\n" +
	"\x0fCreateEditorial\x12\".eolymp.atlas.CreateEditorialInput\x1a#.eolymp.atlas.CreateEditorialOutput\"=\xea\xe2\n" +
	"\v\xf5\xe2\n" +
	"\x00\x00\x80?\xf8\xe2\n" +
	"\x05\x82\xe3\n" +
	"\x17\x8a\xe3\n" +
	"\x13atlas:problem:write\x82\xd3\xe4\x93\x02\r\x1a\v/editorials\x12\xa8\x01\n" +
	"\x0fUpdateEditorial\x12\".eolymp.atlas.UpdateEditorialInput\x1a#.eolymp.atlas.UpdateEditorialOutput\"L\xea\xe2\n" +
	"\v\xf5\xe2\n" +
	"\x00\x00\x80?\xf8\xe2\n" +
	"\x05\x82\xe3\n" +
	"\x17\x8a\xe3\n" +
	"\x13atlas:problem:write\x82\xd3\xe4\x93\x02\x1c\x1a\x1a/editorials/{editorial_id}\x12\xa8\x01\n" +
	"\x0fDeleteEditorial\x12\".eolymp.atlas.DeleteEditorialInput\x1a#.eolymp.atlas.DeleteEditorialOutput\"L\xea\xe2\n" +
	"\v\xf5\xe2\n" +
	"\x00\x00\x80?\xf8\xe2\n" +
	"\x05\x82\xe3\n" +
	"\x17\x8a\xe3\n" +
	"\x13atlas:problem:write\x82\xd3\xe4\x93\x02\x1c*\x1a/editorials/{editorial_id}\x12\xad\x01\n" +
	"\x11DescribeEditorial\x12$.eolymp.atlas.DescribeEditorialInput\x1a%.eolymp.atlas.DescribeEditorialOutput\"K\xea\xe2\n" +
	"\v\xf5\xe2\n" +
	"\x00\x00\xa0A\xf8\xe2\n" +
	"d\x82\xe3\n" +
	"\x16\x8a\xe3\n" +
	"\x12atlas:problem:read\x82\xd3\xe4\x93\x02\x1c\x12\x1a/editorials/{editorial_id}\x12\x97\x01\n" +
	"\x0fLookupEditorial\x12\".eolymp.atlas.LookupEditorialInput\x1a#.eolymp.atlas.LookupEditorialOutput\";\xea\xe2\n" +
	"\v\xf5\xe2\n" +
	"\x00\x00\xa0A\xf8\xe2\n" +
	"d\x82\xe3\n" +
	"\x16\x8a\xe3\n" +
	"\x12atlas:problem:read\x82\xd3\xe4\x93\x02\f\x12\n" +
	"/editorial\x12\xa2\x01\n" +
	"\x10PreviewEditorial\x12#.eolymp.atlas.PreviewEditorialInput\x1a$.eolymp.atlas.PreviewEditorialOutput\"C\xea\xe2\n" +
	"\v\xf5\xe2\n" +
	"\x00\x00\xa0A\xf8\xe2\n" +
	"d\x82\xe3\n" +
	"\x16\x8a\xe3\n" +
	"\x12atlas:problem:read\x82\xd3\xe4\x93\x02\x14\"\x12/editorial/preview\x12\x95\x01\n" +
	"\x0eListEditorials\x12!.eolymp.atlas.ListEditorialsInput\x1a\".eolymp.atlas.ListEditorialsOutput\"<\xea\xe2\n" +
	"\v\xf5\xe2\n" +
	"\x00\x00\xa0A\xf8\xe2\n" +
	"d\x82\xe3\n" +
	"\x16\x8a\xe3\n" +
	"\x12atlas:problem:read\x82\xd3\xe4\x93\x02\r\x12\v/editorials\x12\xaf\x01\n" +
	"\x13TranslateEditorials\x12&.eolymp.atlas.TranslateEditorialsInput\x1a'.eolymp.atlas.TranslateEditorialsOutput\"G\xea\xe2\n" +
	"\v\xf5\xe2\n" +
	"\x00\x00\x00@\xf8\xe2\n" +
	"\n" +
	"\x82\xe3\n" +
	"\x17\x8a\xe3\n" +
	"\x13atlas:problem:write\x82\xd3\xe4\x93\x02\x17\"\x15/editorials:translateB-Z+github.com/eolymp/go-sdk/eolymp/atlas;atlasb\x06proto3"

var (
	file_eolymp_atlas_editorial_service_proto_rawDescOnce sync.Once
	file_eolymp_atlas_editorial_service_proto_rawDescData []byte
)

func file_eolymp_atlas_editorial_service_proto_rawDescGZIP() []byte {
	file_eolymp_atlas_editorial_service_proto_rawDescOnce.Do(func() {
		file_eolymp_atlas_editorial_service_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_eolymp_atlas_editorial_service_proto_rawDesc), len(file_eolymp_atlas_editorial_service_proto_rawDesc)))
	})
	return file_eolymp_atlas_editorial_service_proto_rawDescData
}

var file_eolymp_atlas_editorial_service_proto_msgTypes = make([]protoimpl.MessageInfo, 16)
var file_eolymp_atlas_editorial_service_proto_goTypes = []any{
	(*ListEditorialsInput)(nil),       // 0: eolymp.atlas.ListEditorialsInput
	(*ListEditorialsOutput)(nil),      // 1: eolymp.atlas.ListEditorialsOutput
	(*DescribeEditorialInput)(nil),    // 2: eolymp.atlas.DescribeEditorialInput
	(*DescribeEditorialOutput)(nil),   // 3: eolymp.atlas.DescribeEditorialOutput
	(*LookupEditorialInput)(nil),      // 4: eolymp.atlas.LookupEditorialInput
	(*LookupEditorialOutput)(nil),     // 5: eolymp.atlas.LookupEditorialOutput
	(*PreviewEditorialInput)(nil),     // 6: eolymp.atlas.PreviewEditorialInput
	(*PreviewEditorialOutput)(nil),    // 7: eolymp.atlas.PreviewEditorialOutput
	(*CreateEditorialInput)(nil),      // 8: eolymp.atlas.CreateEditorialInput
	(*CreateEditorialOutput)(nil),     // 9: eolymp.atlas.CreateEditorialOutput
	(*UpdateEditorialInput)(nil),      // 10: eolymp.atlas.UpdateEditorialInput
	(*UpdateEditorialOutput)(nil),     // 11: eolymp.atlas.UpdateEditorialOutput
	(*DeleteEditorialInput)(nil),      // 12: eolymp.atlas.DeleteEditorialInput
	(*DeleteEditorialOutput)(nil),     // 13: eolymp.atlas.DeleteEditorialOutput
	(*TranslateEditorialsInput)(nil),  // 14: eolymp.atlas.TranslateEditorialsInput
	(*TranslateEditorialsOutput)(nil), // 15: eolymp.atlas.TranslateEditorialsOutput
	(Editorial_Extra)(0),              // 16: eolymp.atlas.Editorial.Extra
	(*Editorial)(nil),                 // 17: eolymp.atlas.Editorial
	(Editorial_Patch)(0),              // 18: eolymp.atlas.Editorial.Patch
}
var file_eolymp_atlas_editorial_service_proto_depIdxs = []int32{
	16, // 0: eolymp.atlas.ListEditorialsInput.extra:type_name -> eolymp.atlas.Editorial.Extra
	17, // 1: eolymp.atlas.ListEditorialsOutput.items:type_name -> eolymp.atlas.Editorial
	16, // 2: eolymp.atlas.DescribeEditorialInput.extra:type_name -> eolymp.atlas.Editorial.Extra
	17, // 3: eolymp.atlas.DescribeEditorialOutput.editorial:type_name -> eolymp.atlas.Editorial
	16, // 4: eolymp.atlas.LookupEditorialInput.extra:type_name -> eolymp.atlas.Editorial.Extra
	17, // 5: eolymp.atlas.LookupEditorialOutput.editorial:type_name -> eolymp.atlas.Editorial
	17, // 6: eolymp.atlas.PreviewEditorialInput.editorial:type_name -> eolymp.atlas.Editorial
	17, // 7: eolymp.atlas.PreviewEditorialOutput.editorial:type_name -> eolymp.atlas.Editorial
	17, // 8: eolymp.atlas.CreateEditorialInput.editorial:type_name -> eolymp.atlas.Editorial
	18, // 9: eolymp.atlas.UpdateEditorialInput.patch:type_name -> eolymp.atlas.Editorial.Patch
	17, // 10: eolymp.atlas.UpdateEditorialInput.editorial:type_name -> eolymp.atlas.Editorial
	8,  // 11: eolymp.atlas.EditorialService.CreateEditorial:input_type -> eolymp.atlas.CreateEditorialInput
	10, // 12: eolymp.atlas.EditorialService.UpdateEditorial:input_type -> eolymp.atlas.UpdateEditorialInput
	12, // 13: eolymp.atlas.EditorialService.DeleteEditorial:input_type -> eolymp.atlas.DeleteEditorialInput
	2,  // 14: eolymp.atlas.EditorialService.DescribeEditorial:input_type -> eolymp.atlas.DescribeEditorialInput
	4,  // 15: eolymp.atlas.EditorialService.LookupEditorial:input_type -> eolymp.atlas.LookupEditorialInput
	6,  // 16: eolymp.atlas.EditorialService.PreviewEditorial:input_type -> eolymp.atlas.PreviewEditorialInput
	0,  // 17: eolymp.atlas.EditorialService.ListEditorials:input_type -> eolymp.atlas.ListEditorialsInput
	14, // 18: eolymp.atlas.EditorialService.TranslateEditorials:input_type -> eolymp.atlas.TranslateEditorialsInput
	9,  // 19: eolymp.atlas.EditorialService.CreateEditorial:output_type -> eolymp.atlas.CreateEditorialOutput
	11, // 20: eolymp.atlas.EditorialService.UpdateEditorial:output_type -> eolymp.atlas.UpdateEditorialOutput
	13, // 21: eolymp.atlas.EditorialService.DeleteEditorial:output_type -> eolymp.atlas.DeleteEditorialOutput
	3,  // 22: eolymp.atlas.EditorialService.DescribeEditorial:output_type -> eolymp.atlas.DescribeEditorialOutput
	5,  // 23: eolymp.atlas.EditorialService.LookupEditorial:output_type -> eolymp.atlas.LookupEditorialOutput
	7,  // 24: eolymp.atlas.EditorialService.PreviewEditorial:output_type -> eolymp.atlas.PreviewEditorialOutput
	1,  // 25: eolymp.atlas.EditorialService.ListEditorials:output_type -> eolymp.atlas.ListEditorialsOutput
	15, // 26: eolymp.atlas.EditorialService.TranslateEditorials:output_type -> eolymp.atlas.TranslateEditorialsOutput
	19, // [19:27] is the sub-list for method output_type
	11, // [11:19] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_eolymp_atlas_editorial_service_proto_init() }
func file_eolymp_atlas_editorial_service_proto_init() {
	if File_eolymp_atlas_editorial_service_proto != nil {
		return
	}
	file_eolymp_atlas_editorial_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_eolymp_atlas_editorial_service_proto_rawDesc), len(file_eolymp_atlas_editorial_service_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   16,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_eolymp_atlas_editorial_service_proto_goTypes,
		DependencyIndexes: file_eolymp_atlas_editorial_service_proto_depIdxs,
		MessageInfos:      file_eolymp_atlas_editorial_service_proto_msgTypes,
	}.Build()
	File_eolymp_atlas_editorial_service_proto = out.File
	file_eolymp_atlas_editorial_service_proto_goTypes = nil
	file_eolymp_atlas_editorial_service_proto_depIdxs = nil
}
