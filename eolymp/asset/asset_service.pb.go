// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v3.18.1
// source: eolymp/asset/asset_service.proto

package asset

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

type UploadImageInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name     string                   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Type     string                   `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	Crop     *UploadImageInput_Crop   `protobuf:"bytes,10,opt,name=crop,proto3" json:"crop,omitempty"`
	Size     *UploadImageInput_Size   `protobuf:"bytes,11,opt,name=size,proto3" json:"size,omitempty"`
	Variants []*UploadImageInput_Size `protobuf:"bytes,20,rep,name=variants,proto3" json:"variants,omitempty"`
	Data     []byte                   `protobuf:"bytes,100,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *UploadImageInput) Reset() {
	*x = UploadImageInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_asset_asset_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UploadImageInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadImageInput) ProtoMessage() {}

func (x *UploadImageInput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_asset_asset_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadImageInput.ProtoReflect.Descriptor instead.
func (*UploadImageInput) Descriptor() ([]byte, []int) {
	return file_eolymp_asset_asset_service_proto_rawDescGZIP(), []int{0}
}

func (x *UploadImageInput) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UploadImageInput) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *UploadImageInput) GetCrop() *UploadImageInput_Crop {
	if x != nil {
		return x.Crop
	}
	return nil
}

func (x *UploadImageInput) GetSize() *UploadImageInput_Size {
	if x != nil {
		return x.Size
	}
	return nil
}

func (x *UploadImageInput) GetVariants() []*UploadImageInput_Size {
	if x != nil {
		return x.Variants
	}
	return nil
}

func (x *UploadImageInput) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

type UploadImageOutput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ImageUrl string `protobuf:"bytes,1,opt,name=image_url,json=imageUrl,proto3" json:"image_url,omitempty"`
}

func (x *UploadImageOutput) Reset() {
	*x = UploadImageOutput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_asset_asset_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UploadImageOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadImageOutput) ProtoMessage() {}

func (x *UploadImageOutput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_asset_asset_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadImageOutput.ProtoReflect.Descriptor instead.
func (*UploadImageOutput) Descriptor() ([]byte, []int) {
	return file_eolymp_asset_asset_service_proto_rawDescGZIP(), []int{1}
}

func (x *UploadImageOutput) GetImageUrl() string {
	if x != nil {
		return x.ImageUrl
	}
	return ""
}

type UploadFileInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Type string `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	Data []byte `protobuf:"bytes,100,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *UploadFileInput) Reset() {
	*x = UploadFileInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_asset_asset_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UploadFileInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadFileInput) ProtoMessage() {}

func (x *UploadFileInput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_asset_asset_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadFileInput.ProtoReflect.Descriptor instead.
func (*UploadFileInput) Descriptor() ([]byte, []int) {
	return file_eolymp_asset_asset_service_proto_rawDescGZIP(), []int{2}
}

func (x *UploadFileInput) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UploadFileInput) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *UploadFileInput) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

type UploadFileOutput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FileUrl string `protobuf:"bytes,1,opt,name=file_url,json=fileUrl,proto3" json:"file_url,omitempty"`
}

func (x *UploadFileOutput) Reset() {
	*x = UploadFileOutput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_asset_asset_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UploadFileOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadFileOutput) ProtoMessage() {}

func (x *UploadFileOutput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_asset_asset_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadFileOutput.ProtoReflect.Descriptor instead.
func (*UploadFileOutput) Descriptor() ([]byte, []int) {
	return file_eolymp_asset_asset_service_proto_rawDescGZIP(), []int{3}
}

func (x *UploadFileOutput) GetFileUrl() string {
	if x != nil {
		return x.FileUrl
	}
	return ""
}

type UploadImageInput_Size struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Width  uint32 `protobuf:"varint,1,opt,name=width,proto3" json:"width,omitempty"`
	Height uint32 `protobuf:"varint,2,opt,name=height,proto3" json:"height,omitempty"`
}

func (x *UploadImageInput_Size) Reset() {
	*x = UploadImageInput_Size{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_asset_asset_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UploadImageInput_Size) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadImageInput_Size) ProtoMessage() {}

func (x *UploadImageInput_Size) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_asset_asset_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadImageInput_Size.ProtoReflect.Descriptor instead.
func (*UploadImageInput_Size) Descriptor() ([]byte, []int) {
	return file_eolymp_asset_asset_service_proto_rawDescGZIP(), []int{0, 0}
}

func (x *UploadImageInput_Size) GetWidth() uint32 {
	if x != nil {
		return x.Width
	}
	return 0
}

func (x *UploadImageInput_Size) GetHeight() uint32 {
	if x != nil {
		return x.Height
	}
	return 0
}

type UploadImageInput_Crop struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Top    uint32 `protobuf:"varint,1,opt,name=top,proto3" json:"top,omitempty"`
	Right  uint32 `protobuf:"varint,2,opt,name=right,proto3" json:"right,omitempty"`
	Bottom uint32 `protobuf:"varint,3,opt,name=bottom,proto3" json:"bottom,omitempty"`
	Left   uint32 `protobuf:"varint,4,opt,name=left,proto3" json:"left,omitempty"`
}

func (x *UploadImageInput_Crop) Reset() {
	*x = UploadImageInput_Crop{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_asset_asset_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UploadImageInput_Crop) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadImageInput_Crop) ProtoMessage() {}

func (x *UploadImageInput_Crop) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_asset_asset_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadImageInput_Crop.ProtoReflect.Descriptor instead.
func (*UploadImageInput_Crop) Descriptor() ([]byte, []int) {
	return file_eolymp_asset_asset_service_proto_rawDescGZIP(), []int{0, 1}
}

func (x *UploadImageInput_Crop) GetTop() uint32 {
	if x != nil {
		return x.Top
	}
	return 0
}

func (x *UploadImageInput_Crop) GetRight() uint32 {
	if x != nil {
		return x.Right
	}
	return 0
}

func (x *UploadImageInput_Crop) GetBottom() uint32 {
	if x != nil {
		return x.Bottom
	}
	return 0
}

func (x *UploadImageInput_Crop) GetLeft() uint32 {
	if x != nil {
		return x.Left
	}
	return 0
}

var File_eolymp_asset_asset_service_proto protoreflect.FileDescriptor

var file_eolymp_asset_asset_service_proto_rawDesc = []byte{
	0x0a, 0x20, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x61, 0x73, 0x73, 0x65, 0x74, 0x2f, 0x61,
	0x73, 0x73, 0x65, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0c, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x61, 0x73, 0x73, 0x65, 0x74,
	0x1a, 0x1d, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x22, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2f, 0x72, 0x61, 0x74, 0x65, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x93, 0x03, 0x0a, 0x10, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x49, 0x6d,
	0x61, 0x67, 0x65, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x12, 0x37, 0x0a, 0x04, 0x63, 0x72, 0x6f, 0x70, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23,
	0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x61, 0x73, 0x73, 0x65, 0x74, 0x2e, 0x55, 0x70,
	0x6c, 0x6f, 0x61, 0x64, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x2e, 0x43,
	0x72, 0x6f, 0x70, 0x52, 0x04, 0x63, 0x72, 0x6f, 0x70, 0x12, 0x37, 0x0a, 0x04, 0x73, 0x69, 0x7a,
	0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70,
	0x2e, 0x61, 0x73, 0x73, 0x65, 0x74, 0x2e, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x49, 0x6d, 0x61,
	0x67, 0x65, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x2e, 0x53, 0x69, 0x7a, 0x65, 0x52, 0x04, 0x73, 0x69,
	0x7a, 0x65, 0x12, 0x3f, 0x0a, 0x08, 0x76, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x74, 0x73, 0x18, 0x14,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x61, 0x73,
	0x73, 0x65, 0x74, 0x2e, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x49,
	0x6e, 0x70, 0x75, 0x74, 0x2e, 0x53, 0x69, 0x7a, 0x65, 0x52, 0x08, 0x76, 0x61, 0x72, 0x69, 0x61,
	0x6e, 0x74, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x64, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x1a, 0x34, 0x0a, 0x04, 0x53, 0x69, 0x7a, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x77, 0x69, 0x64, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05,
	0x77, 0x69, 0x64, 0x74, 0x68, 0x12, 0x16, 0x0a, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x1a, 0x5a, 0x0a,
	0x04, 0x43, 0x72, 0x6f, 0x70, 0x12, 0x10, 0x0a, 0x03, 0x74, 0x6f, 0x70, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x03, 0x74, 0x6f, 0x70, 0x12, 0x14, 0x0a, 0x05, 0x72, 0x69, 0x67, 0x68, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x72, 0x69, 0x67, 0x68, 0x74, 0x12, 0x16, 0x0a,
	0x06, 0x62, 0x6f, 0x74, 0x74, 0x6f, 0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x62,
	0x6f, 0x74, 0x74, 0x6f, 0x6d, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x65, 0x66, 0x74, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x04, 0x6c, 0x65, 0x66, 0x74, 0x22, 0x30, 0x0a, 0x11, 0x55, 0x70, 0x6c,
	0x6f, 0x61, 0x64, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x12, 0x1b,
	0x0a, 0x09, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x55, 0x72, 0x6c, 0x22, 0x4d, 0x0a, 0x0f, 0x55,
	0x70, 0x6c, 0x6f, 0x61, 0x64, 0x46, 0x69, 0x6c, 0x65, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x64,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x2d, 0x0a, 0x10, 0x55, 0x70,
	0x6c, 0x6f, 0x61, 0x64, 0x46, 0x69, 0x6c, 0x65, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x12, 0x19,
	0x0a, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x66, 0x69, 0x6c, 0x65, 0x55, 0x72, 0x6c, 0x32, 0xac, 0x02, 0x0a, 0x0c, 0x41, 0x73,
	0x73, 0x65, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x8e, 0x01, 0x0a, 0x0b, 0x55,
	0x70, 0x6c, 0x6f, 0x61, 0x64, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x1e, 0x2e, 0x65, 0x6f, 0x6c,
	0x79, 0x6d, 0x70, 0x2e, 0x61, 0x73, 0x73, 0x65, 0x74, 0x2e, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64,
	0x49, 0x6d, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x1f, 0x2e, 0x65, 0x6f, 0x6c,
	0x79, 0x6d, 0x70, 0x2e, 0x61, 0x73, 0x73, 0x65, 0x74, 0x2e, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64,
	0x49, 0x6d, 0x61, 0x67, 0x65, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x22, 0x3e, 0xea, 0xe2, 0x0a,
	0x0b, 0xf5, 0xe2, 0x0a, 0x00, 0x00, 0x00, 0x40, 0xf8, 0xe2, 0x0a, 0x64, 0x82, 0xe3, 0x0a, 0x15,
	0x8a, 0xe3, 0x0a, 0x11, 0x61, 0x73, 0x73, 0x65, 0x74, 0x3a, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x3a,
	0x77, 0x72, 0x69, 0x74, 0x65, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x10, 0x22, 0x0e, 0x2f, 0x61, 0x73,
	0x73, 0x65, 0x74, 0x73, 0x2f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x12, 0x8a, 0x01, 0x0a, 0x0a,
	0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x1d, 0x2e, 0x65, 0x6f, 0x6c,
	0x79, 0x6d, 0x70, 0x2e, 0x61, 0x73, 0x73, 0x65, 0x74, 0x2e, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64,
	0x46, 0x69, 0x6c, 0x65, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x1e, 0x2e, 0x65, 0x6f, 0x6c, 0x79,
	0x6d, 0x70, 0x2e, 0x61, 0x73, 0x73, 0x65, 0x74, 0x2e, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x46,
	0x69, 0x6c, 0x65, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x22, 0x3d, 0xea, 0xe2, 0x0a, 0x0b, 0xf5,
	0xe2, 0x0a, 0x00, 0x00, 0x00, 0x40, 0xf8, 0xe2, 0x0a, 0x64, 0x82, 0xe3, 0x0a, 0x15, 0x8a, 0xe3,
	0x0a, 0x11, 0x61, 0x73, 0x73, 0x65, 0x74, 0x3a, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x3a, 0x77, 0x72,
	0x69, 0x74, 0x65, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0f, 0x22, 0x0d, 0x2f, 0x61, 0x73, 0x73, 0x65,
	0x74, 0x73, 0x2f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x42, 0x2d, 0x5a, 0x2b, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x67, 0x6f,
	0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x61, 0x73, 0x73, 0x65,
	0x74, 0x3b, 0x61, 0x73, 0x73, 0x65, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_eolymp_asset_asset_service_proto_rawDescOnce sync.Once
	file_eolymp_asset_asset_service_proto_rawDescData = file_eolymp_asset_asset_service_proto_rawDesc
)

func file_eolymp_asset_asset_service_proto_rawDescGZIP() []byte {
	file_eolymp_asset_asset_service_proto_rawDescOnce.Do(func() {
		file_eolymp_asset_asset_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_eolymp_asset_asset_service_proto_rawDescData)
	})
	return file_eolymp_asset_asset_service_proto_rawDescData
}

var file_eolymp_asset_asset_service_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_eolymp_asset_asset_service_proto_goTypes = []interface{}{
	(*UploadImageInput)(nil),      // 0: eolymp.asset.UploadImageInput
	(*UploadImageOutput)(nil),     // 1: eolymp.asset.UploadImageOutput
	(*UploadFileInput)(nil),       // 2: eolymp.asset.UploadFileInput
	(*UploadFileOutput)(nil),      // 3: eolymp.asset.UploadFileOutput
	(*UploadImageInput_Size)(nil), // 4: eolymp.asset.UploadImageInput.Size
	(*UploadImageInput_Crop)(nil), // 5: eolymp.asset.UploadImageInput.Crop
}
var file_eolymp_asset_asset_service_proto_depIdxs = []int32{
	5, // 0: eolymp.asset.UploadImageInput.crop:type_name -> eolymp.asset.UploadImageInput.Crop
	4, // 1: eolymp.asset.UploadImageInput.size:type_name -> eolymp.asset.UploadImageInput.Size
	4, // 2: eolymp.asset.UploadImageInput.variants:type_name -> eolymp.asset.UploadImageInput.Size
	0, // 3: eolymp.asset.AssetService.UploadImage:input_type -> eolymp.asset.UploadImageInput
	2, // 4: eolymp.asset.AssetService.UploadFile:input_type -> eolymp.asset.UploadFileInput
	1, // 5: eolymp.asset.AssetService.UploadImage:output_type -> eolymp.asset.UploadImageOutput
	3, // 6: eolymp.asset.AssetService.UploadFile:output_type -> eolymp.asset.UploadFileOutput
	5, // [5:7] is the sub-list for method output_type
	3, // [3:5] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_eolymp_asset_asset_service_proto_init() }
func file_eolymp_asset_asset_service_proto_init() {
	if File_eolymp_asset_asset_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_eolymp_asset_asset_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UploadImageInput); i {
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
		file_eolymp_asset_asset_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UploadImageOutput); i {
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
		file_eolymp_asset_asset_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UploadFileInput); i {
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
		file_eolymp_asset_asset_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UploadFileOutput); i {
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
		file_eolymp_asset_asset_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UploadImageInput_Size); i {
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
		file_eolymp_asset_asset_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UploadImageInput_Crop); i {
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
			RawDescriptor: file_eolymp_asset_asset_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_eolymp_asset_asset_service_proto_goTypes,
		DependencyIndexes: file_eolymp_asset_asset_service_proto_depIdxs,
		MessageInfos:      file_eolymp_asset_asset_service_proto_msgTypes,
	}.Build()
	File_eolymp_asset_asset_service_proto = out.File
	file_eolymp_asset_asset_service_proto_rawDesc = nil
	file_eolymp_asset_asset_service_proto_goTypes = nil
	file_eolymp_asset_asset_service_proto_depIdxs = nil
}
