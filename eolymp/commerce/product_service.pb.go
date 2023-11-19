// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.18.1
// source: eolymp/commerce/product_service.proto

package commerce

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

type CreateProductInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Product *Product `protobuf:"bytes,1,opt,name=product,proto3" json:"product,omitempty"`
}

func (x *CreateProductInput) Reset() {
	*x = CreateProductInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_commerce_product_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateProductInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateProductInput) ProtoMessage() {}

func (x *CreateProductInput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_commerce_product_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateProductInput.ProtoReflect.Descriptor instead.
func (*CreateProductInput) Descriptor() ([]byte, []int) {
	return file_eolymp_commerce_product_service_proto_rawDescGZIP(), []int{0}
}

func (x *CreateProductInput) GetProduct() *Product {
	if x != nil {
		return x.Product
	}
	return nil
}

type CreateProductOutput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProductId string `protobuf:"bytes,1,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
}

func (x *CreateProductOutput) Reset() {
	*x = CreateProductOutput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_commerce_product_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateProductOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateProductOutput) ProtoMessage() {}

func (x *CreateProductOutput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_commerce_product_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateProductOutput.ProtoReflect.Descriptor instead.
func (*CreateProductOutput) Descriptor() ([]byte, []int) {
	return file_eolymp_commerce_product_service_proto_rawDescGZIP(), []int{1}
}

func (x *CreateProductOutput) GetProductId() string {
	if x != nil {
		return x.ProductId
	}
	return ""
}

type UpdateProductInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProductId string   `protobuf:"bytes,1,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
	Product   *Product `protobuf:"bytes,2,opt,name=product,proto3" json:"product,omitempty"`
}

func (x *UpdateProductInput) Reset() {
	*x = UpdateProductInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_commerce_product_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateProductInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateProductInput) ProtoMessage() {}

func (x *UpdateProductInput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_commerce_product_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateProductInput.ProtoReflect.Descriptor instead.
func (*UpdateProductInput) Descriptor() ([]byte, []int) {
	return file_eolymp_commerce_product_service_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateProductInput) GetProductId() string {
	if x != nil {
		return x.ProductId
	}
	return ""
}

func (x *UpdateProductInput) GetProduct() *Product {
	if x != nil {
		return x.Product
	}
	return nil
}

type UpdateProductOutput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateProductOutput) Reset() {
	*x = UpdateProductOutput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_commerce_product_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateProductOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateProductOutput) ProtoMessage() {}

func (x *UpdateProductOutput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_commerce_product_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateProductOutput.ProtoReflect.Descriptor instead.
func (*UpdateProductOutput) Descriptor() ([]byte, []int) {
	return file_eolymp_commerce_product_service_proto_rawDescGZIP(), []int{3}
}

type DescribeProductInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProductId string `protobuf:"bytes,1,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
}

func (x *DescribeProductInput) Reset() {
	*x = DescribeProductInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_commerce_product_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DescribeProductInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescribeProductInput) ProtoMessage() {}

func (x *DescribeProductInput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_commerce_product_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DescribeProductInput.ProtoReflect.Descriptor instead.
func (*DescribeProductInput) Descriptor() ([]byte, []int) {
	return file_eolymp_commerce_product_service_proto_rawDescGZIP(), []int{4}
}

func (x *DescribeProductInput) GetProductId() string {
	if x != nil {
		return x.ProductId
	}
	return ""
}

type DescribeProductOutput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Product *Product `protobuf:"bytes,1,opt,name=product,proto3" json:"product,omitempty"`
}

func (x *DescribeProductOutput) Reset() {
	*x = DescribeProductOutput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_commerce_product_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DescribeProductOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescribeProductOutput) ProtoMessage() {}

func (x *DescribeProductOutput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_commerce_product_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DescribeProductOutput.ProtoReflect.Descriptor instead.
func (*DescribeProductOutput) Descriptor() ([]byte, []int) {
	return file_eolymp_commerce_product_service_proto_rawDescGZIP(), []int{5}
}

func (x *DescribeProductOutput) GetProduct() *Product {
	if x != nil {
		return x.Product
	}
	return nil
}

type DeleteProductInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProductId string `protobuf:"bytes,1,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
}

func (x *DeleteProductInput) Reset() {
	*x = DeleteProductInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_commerce_product_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteProductInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteProductInput) ProtoMessage() {}

func (x *DeleteProductInput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_commerce_product_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteProductInput.ProtoReflect.Descriptor instead.
func (*DeleteProductInput) Descriptor() ([]byte, []int) {
	return file_eolymp_commerce_product_service_proto_rawDescGZIP(), []int{6}
}

func (x *DeleteProductInput) GetProductId() string {
	if x != nil {
		return x.ProductId
	}
	return ""
}

type DeleteProductOutput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteProductOutput) Reset() {
	*x = DeleteProductOutput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_commerce_product_service_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteProductOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteProductOutput) ProtoMessage() {}

func (x *DeleteProductOutput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_commerce_product_service_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteProductOutput.ProtoReflect.Descriptor instead.
func (*DeleteProductOutput) Descriptor() ([]byte, []int) {
	return file_eolymp_commerce_product_service_proto_rawDescGZIP(), []int{7}
}

type ListProductPricesInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Currency string `protobuf:"bytes,1,opt,name=currency,proto3" json:"currency,omitempty"`
}

func (x *ListProductPricesInput) Reset() {
	*x = ListProductPricesInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_commerce_product_service_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListProductPricesInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListProductPricesInput) ProtoMessage() {}

func (x *ListProductPricesInput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_commerce_product_service_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListProductPricesInput.ProtoReflect.Descriptor instead.
func (*ListProductPricesInput) Descriptor() ([]byte, []int) {
	return file_eolymp_commerce_product_service_proto_rawDescGZIP(), []int{8}
}

func (x *ListProductPricesInput) GetCurrency() string {
	if x != nil {
		return x.Currency
	}
	return ""
}

type ListProductPricesOutput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items []*Product_Price `protobuf:"bytes,2,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *ListProductPricesOutput) Reset() {
	*x = ListProductPricesOutput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_commerce_product_service_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListProductPricesOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListProductPricesOutput) ProtoMessage() {}

func (x *ListProductPricesOutput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_commerce_product_service_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListProductPricesOutput.ProtoReflect.Descriptor instead.
func (*ListProductPricesOutput) Descriptor() ([]byte, []int) {
	return file_eolymp_commerce_product_service_proto_rawDescGZIP(), []int{9}
}

func (x *ListProductPricesOutput) GetItems() []*Product_Price {
	if x != nil {
		return x.Items
	}
	return nil
}

type AddProductPriceInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProductId  string         `protobuf:"bytes,1,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
	Price      *Product_Price `protobuf:"bytes,2,opt,name=price,proto3" json:"price,omitempty"`
	SetDefault bool           `protobuf:"varint,3,opt,name=set_default,json=setDefault,proto3" json:"set_default,omitempty"`
}

func (x *AddProductPriceInput) Reset() {
	*x = AddProductPriceInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_commerce_product_service_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddProductPriceInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddProductPriceInput) ProtoMessage() {}

func (x *AddProductPriceInput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_commerce_product_service_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddProductPriceInput.ProtoReflect.Descriptor instead.
func (*AddProductPriceInput) Descriptor() ([]byte, []int) {
	return file_eolymp_commerce_product_service_proto_rawDescGZIP(), []int{10}
}

func (x *AddProductPriceInput) GetProductId() string {
	if x != nil {
		return x.ProductId
	}
	return ""
}

func (x *AddProductPriceInput) GetPrice() *Product_Price {
	if x != nil {
		return x.Price
	}
	return nil
}

func (x *AddProductPriceInput) GetSetDefault() bool {
	if x != nil {
		return x.SetDefault
	}
	return false
}

type AddProductPriceOutput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AddProductPriceOutput) Reset() {
	*x = AddProductPriceOutput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_commerce_product_service_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddProductPriceOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddProductPriceOutput) ProtoMessage() {}

func (x *AddProductPriceOutput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_commerce_product_service_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddProductPriceOutput.ProtoReflect.Descriptor instead.
func (*AddProductPriceOutput) Descriptor() ([]byte, []int) {
	return file_eolymp_commerce_product_service_proto_rawDescGZIP(), []int{11}
}

type RemoveProductPriceInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProductId string `protobuf:"bytes,1,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
	PriceId   string `protobuf:"bytes,2,opt,name=price_id,json=priceId,proto3" json:"price_id,omitempty"`
}

func (x *RemoveProductPriceInput) Reset() {
	*x = RemoveProductPriceInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_commerce_product_service_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveProductPriceInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveProductPriceInput) ProtoMessage() {}

func (x *RemoveProductPriceInput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_commerce_product_service_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveProductPriceInput.ProtoReflect.Descriptor instead.
func (*RemoveProductPriceInput) Descriptor() ([]byte, []int) {
	return file_eolymp_commerce_product_service_proto_rawDescGZIP(), []int{12}
}

func (x *RemoveProductPriceInput) GetProductId() string {
	if x != nil {
		return x.ProductId
	}
	return ""
}

func (x *RemoveProductPriceInput) GetPriceId() string {
	if x != nil {
		return x.PriceId
	}
	return ""
}

type RemoveProductPriceOutput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *RemoveProductPriceOutput) Reset() {
	*x = RemoveProductPriceOutput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_commerce_product_service_proto_msgTypes[13]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveProductPriceOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveProductPriceOutput) ProtoMessage() {}

func (x *RemoveProductPriceOutput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_commerce_product_service_proto_msgTypes[13]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveProductPriceOutput.ProtoReflect.Descriptor instead.
func (*RemoveProductPriceOutput) Descriptor() ([]byte, []int) {
	return file_eolymp_commerce_product_service_proto_rawDescGZIP(), []int{13}
}

var File_eolymp_commerce_product_service_proto protoreflect.FileDescriptor

var file_eolymp_commerce_product_service_proto_rawDesc = []byte{
	0x0a, 0x25, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x63,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x63, 0x65, 0x1a, 0x22, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70,
	0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x72, 0x61, 0x74,
	0x65, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x65, 0x6f,
	0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x63, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x48, 0x0a, 0x12, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x6e, 0x70, 0x75,
	0x74, 0x12, 0x32, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x18, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x65, 0x72, 0x63, 0x65, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x07, 0x70, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x22, 0x34, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x12, 0x1d, 0x0a, 0x0a,
	0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x22, 0x67, 0x0a, 0x12, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x6e, 0x70, 0x75,
	0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64,
	0x12, 0x32, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x18, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x65,
	0x72, 0x63, 0x65, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x07, 0x70, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x22, 0x15, 0x0a, 0x13, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x22, 0x35, 0x0a, 0x14, 0x44,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x6e,
	0x70, 0x75, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x49, 0x64, 0x22, 0x4b, 0x0a, 0x15, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x50, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x12, 0x32, 0x0a, 0x07, 0x70,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x65,
	0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x63, 0x65, 0x2e, 0x50,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x22,
	0x33, 0x0a, 0x12, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x49, 0x64, 0x22, 0x15, 0x0a, 0x13, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x22, 0x34, 0x0a, 0x16, 0x4c,
	0x69, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x50, 0x72, 0x69, 0x63, 0x65, 0x73,
	0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63,
	0x79, 0x22, 0x4f, 0x0a, 0x17, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x50, 0x72, 0x69, 0x63, 0x65, 0x73, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x12, 0x34, 0x0a, 0x05,
	0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x65, 0x6f,
	0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x63, 0x65, 0x2e, 0x50, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x50, 0x72, 0x69, 0x63, 0x65, 0x52, 0x05, 0x69, 0x74, 0x65,
	0x6d, 0x73, 0x22, 0x8c, 0x01, 0x0a, 0x14, 0x41, 0x64, 0x64, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x50, 0x72, 0x69, 0x63, 0x65, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x70,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x12, 0x34, 0x0a, 0x05, 0x70, 0x72,
	0x69, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x65, 0x6f, 0x6c, 0x79,
	0x6d, 0x70, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x63, 0x65, 0x2e, 0x50, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x2e, 0x50, 0x72, 0x69, 0x63, 0x65, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65,
	0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x65, 0x74, 0x5f, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x73, 0x65, 0x74, 0x44, 0x65, 0x66, 0x61, 0x75, 0x6c,
	0x74, 0x22, 0x17, 0x0a, 0x15, 0x41, 0x64, 0x64, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x50,
	0x72, 0x69, 0x63, 0x65, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x22, 0x53, 0x0a, 0x17, 0x52, 0x65,
	0x6d, 0x6f, 0x76, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x50, 0x72, 0x69, 0x63, 0x65,
	0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x70, 0x72, 0x69, 0x63, 0x65, 0x5f, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x72, 0x69, 0x63, 0x65, 0x49, 0x64, 0x22,
	0x1a, 0x0a, 0x18, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x50, 0x72, 0x69, 0x63, 0x65, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x32, 0xb2, 0x06, 0x0a, 0x0e,
	0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x6b,
	0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x12,
	0x23, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x63,
	0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49,
	0x6e, 0x70, 0x75, 0x74, 0x1a, 0x24, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x65, 0x72, 0x63, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x22, 0x0f, 0xea, 0xe2, 0x0a, 0x0b,
	0xf5, 0xe2, 0x0a, 0x00, 0x00, 0xa0, 0x40, 0xf8, 0xe2, 0x0a, 0x32, 0x12, 0x6b, 0x0a, 0x0d, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x12, 0x23, 0x2e, 0x65,
	0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x63, 0x65, 0x2e, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x6e, 0x70, 0x75,
	0x74, 0x1a, 0x24, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x65,
	0x72, 0x63, 0x65, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x22, 0x0f, 0xea, 0xe2, 0x0a, 0x0b, 0xf5, 0xe2, 0x0a,
	0x00, 0x00, 0xa0, 0x40, 0xf8, 0xe2, 0x0a, 0x32, 0x12, 0x71, 0x0a, 0x0f, 0x44, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x62, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x12, 0x25, 0x2e, 0x65, 0x6f,
	0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x63, 0x65, 0x2e, 0x44, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x6e, 0x70,
	0x75, 0x74, 0x1a, 0x26, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x65, 0x72, 0x63, 0x65, 0x2e, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x50, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x22, 0x0f, 0xea, 0xe2, 0x0a, 0x0b,
	0xf5, 0xe2, 0x0a, 0x00, 0x00, 0xa0, 0x40, 0xf8, 0xe2, 0x0a, 0x32, 0x12, 0x6b, 0x0a, 0x0d, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x12, 0x23, 0x2e, 0x65,
	0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x63, 0x65, 0x2e, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x6e, 0x70, 0x75,
	0x74, 0x1a, 0x24, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x65,
	0x72, 0x63, 0x65, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x22, 0x0f, 0xea, 0xe2, 0x0a, 0x0b, 0xf5, 0xe2, 0x0a,
	0x00, 0x00, 0xa0, 0x40, 0xf8, 0xe2, 0x0a, 0x32, 0x12, 0x77, 0x0a, 0x11, 0x4c, 0x69, 0x73, 0x74,
	0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x50, 0x72, 0x69, 0x63, 0x65, 0x73, 0x12, 0x27, 0x2e,
	0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x63, 0x65, 0x2e,
	0x4c, 0x69, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x50, 0x72, 0x69, 0x63, 0x65,
	0x73, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x28, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x63, 0x65, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x50, 0x72, 0x69, 0x63, 0x65, 0x73, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74,
	0x22, 0x0f, 0xea, 0xe2, 0x0a, 0x0b, 0xf5, 0xe2, 0x0a, 0x00, 0x00, 0xa0, 0x40, 0xf8, 0xe2, 0x0a,
	0x32, 0x12, 0x71, 0x0a, 0x0f, 0x41, 0x64, 0x64, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x50,
	0x72, 0x69, 0x63, 0x65, 0x12, 0x25, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x65, 0x72, 0x63, 0x65, 0x2e, 0x41, 0x64, 0x64, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x50, 0x72, 0x69, 0x63, 0x65, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x26, 0x2e, 0x65, 0x6f,
	0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x63, 0x65, 0x2e, 0x41, 0x64,
	0x64, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x50, 0x72, 0x69, 0x63, 0x65, 0x4f, 0x75, 0x74,
	0x70, 0x75, 0x74, 0x22, 0x0f, 0xea, 0xe2, 0x0a, 0x0b, 0xf5, 0xe2, 0x0a, 0x00, 0x00, 0xa0, 0x40,
	0xf8, 0xe2, 0x0a, 0x32, 0x12, 0x7a, 0x0a, 0x12, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x50, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x28, 0x2e, 0x65, 0x6f, 0x6c,
	0x79, 0x6d, 0x70, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x63, 0x65, 0x2e, 0x52, 0x65, 0x6d,
	0x6f, 0x76, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x50, 0x72, 0x69, 0x63, 0x65, 0x49,
	0x6e, 0x70, 0x75, 0x74, 0x1a, 0x29, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x65, 0x72, 0x63, 0x65, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x50, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x50, 0x72, 0x69, 0x63, 0x65, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x22,
	0x0f, 0xea, 0xe2, 0x0a, 0x0b, 0xf5, 0xe2, 0x0a, 0x00, 0x00, 0xa0, 0x40, 0xf8, 0xe2, 0x0a, 0x32,
	0x42, 0x33, 0x5a, 0x31, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65,
	0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x67, 0x6f, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x65, 0x6f, 0x6c,
	0x79, 0x6d, 0x70, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x63, 0x65, 0x3b, 0x63, 0x6f, 0x6d,
	0x6d, 0x65, 0x72, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_eolymp_commerce_product_service_proto_rawDescOnce sync.Once
	file_eolymp_commerce_product_service_proto_rawDescData = file_eolymp_commerce_product_service_proto_rawDesc
)

func file_eolymp_commerce_product_service_proto_rawDescGZIP() []byte {
	file_eolymp_commerce_product_service_proto_rawDescOnce.Do(func() {
		file_eolymp_commerce_product_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_eolymp_commerce_product_service_proto_rawDescData)
	})
	return file_eolymp_commerce_product_service_proto_rawDescData
}

var file_eolymp_commerce_product_service_proto_msgTypes = make([]protoimpl.MessageInfo, 14)
var file_eolymp_commerce_product_service_proto_goTypes = []interface{}{
	(*CreateProductInput)(nil),       // 0: eolymp.commerce.CreateProductInput
	(*CreateProductOutput)(nil),      // 1: eolymp.commerce.CreateProductOutput
	(*UpdateProductInput)(nil),       // 2: eolymp.commerce.UpdateProductInput
	(*UpdateProductOutput)(nil),      // 3: eolymp.commerce.UpdateProductOutput
	(*DescribeProductInput)(nil),     // 4: eolymp.commerce.DescribeProductInput
	(*DescribeProductOutput)(nil),    // 5: eolymp.commerce.DescribeProductOutput
	(*DeleteProductInput)(nil),       // 6: eolymp.commerce.DeleteProductInput
	(*DeleteProductOutput)(nil),      // 7: eolymp.commerce.DeleteProductOutput
	(*ListProductPricesInput)(nil),   // 8: eolymp.commerce.ListProductPricesInput
	(*ListProductPricesOutput)(nil),  // 9: eolymp.commerce.ListProductPricesOutput
	(*AddProductPriceInput)(nil),     // 10: eolymp.commerce.AddProductPriceInput
	(*AddProductPriceOutput)(nil),    // 11: eolymp.commerce.AddProductPriceOutput
	(*RemoveProductPriceInput)(nil),  // 12: eolymp.commerce.RemoveProductPriceInput
	(*RemoveProductPriceOutput)(nil), // 13: eolymp.commerce.RemoveProductPriceOutput
	(*Product)(nil),                  // 14: eolymp.commerce.Product
	(*Product_Price)(nil),            // 15: eolymp.commerce.Product.Price
}
var file_eolymp_commerce_product_service_proto_depIdxs = []int32{
	14, // 0: eolymp.commerce.CreateProductInput.product:type_name -> eolymp.commerce.Product
	14, // 1: eolymp.commerce.UpdateProductInput.product:type_name -> eolymp.commerce.Product
	14, // 2: eolymp.commerce.DescribeProductOutput.product:type_name -> eolymp.commerce.Product
	15, // 3: eolymp.commerce.ListProductPricesOutput.items:type_name -> eolymp.commerce.Product.Price
	15, // 4: eolymp.commerce.AddProductPriceInput.price:type_name -> eolymp.commerce.Product.Price
	0,  // 5: eolymp.commerce.ProductService.CreateProduct:input_type -> eolymp.commerce.CreateProductInput
	2,  // 6: eolymp.commerce.ProductService.UpdateProduct:input_type -> eolymp.commerce.UpdateProductInput
	4,  // 7: eolymp.commerce.ProductService.DescribeProduct:input_type -> eolymp.commerce.DescribeProductInput
	6,  // 8: eolymp.commerce.ProductService.DeleteProduct:input_type -> eolymp.commerce.DeleteProductInput
	8,  // 9: eolymp.commerce.ProductService.ListProductPrices:input_type -> eolymp.commerce.ListProductPricesInput
	10, // 10: eolymp.commerce.ProductService.AddProductPrice:input_type -> eolymp.commerce.AddProductPriceInput
	12, // 11: eolymp.commerce.ProductService.RemoveProductPrice:input_type -> eolymp.commerce.RemoveProductPriceInput
	1,  // 12: eolymp.commerce.ProductService.CreateProduct:output_type -> eolymp.commerce.CreateProductOutput
	3,  // 13: eolymp.commerce.ProductService.UpdateProduct:output_type -> eolymp.commerce.UpdateProductOutput
	5,  // 14: eolymp.commerce.ProductService.DescribeProduct:output_type -> eolymp.commerce.DescribeProductOutput
	7,  // 15: eolymp.commerce.ProductService.DeleteProduct:output_type -> eolymp.commerce.DeleteProductOutput
	9,  // 16: eolymp.commerce.ProductService.ListProductPrices:output_type -> eolymp.commerce.ListProductPricesOutput
	11, // 17: eolymp.commerce.ProductService.AddProductPrice:output_type -> eolymp.commerce.AddProductPriceOutput
	13, // 18: eolymp.commerce.ProductService.RemoveProductPrice:output_type -> eolymp.commerce.RemoveProductPriceOutput
	12, // [12:19] is the sub-list for method output_type
	5,  // [5:12] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_eolymp_commerce_product_service_proto_init() }
func file_eolymp_commerce_product_service_proto_init() {
	if File_eolymp_commerce_product_service_proto != nil {
		return
	}
	file_eolymp_commerce_product_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_eolymp_commerce_product_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateProductInput); i {
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
		file_eolymp_commerce_product_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateProductOutput); i {
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
		file_eolymp_commerce_product_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateProductInput); i {
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
		file_eolymp_commerce_product_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateProductOutput); i {
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
		file_eolymp_commerce_product_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DescribeProductInput); i {
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
		file_eolymp_commerce_product_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DescribeProductOutput); i {
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
		file_eolymp_commerce_product_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteProductInput); i {
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
		file_eolymp_commerce_product_service_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteProductOutput); i {
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
		file_eolymp_commerce_product_service_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListProductPricesInput); i {
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
		file_eolymp_commerce_product_service_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListProductPricesOutput); i {
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
		file_eolymp_commerce_product_service_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddProductPriceInput); i {
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
		file_eolymp_commerce_product_service_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddProductPriceOutput); i {
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
		file_eolymp_commerce_product_service_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveProductPriceInput); i {
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
		file_eolymp_commerce_product_service_proto_msgTypes[13].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveProductPriceOutput); i {
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
			RawDescriptor: file_eolymp_commerce_product_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   14,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_eolymp_commerce_product_service_proto_goTypes,
		DependencyIndexes: file_eolymp_commerce_product_service_proto_depIdxs,
		MessageInfos:      file_eolymp_commerce_product_service_proto_msgTypes,
	}.Build()
	File_eolymp_commerce_product_service_proto = out.File
	file_eolymp_commerce_product_service_proto_rawDesc = nil
	file_eolymp_commerce_product_service_proto_goTypes = nil
	file_eolymp_commerce_product_service_proto_depIdxs = nil
}
