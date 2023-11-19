// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.18.1
// source: eolymp/commerce/product_service.proto

package commerce

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	ProductService_CreateProduct_FullMethodName      = "/eolymp.commerce.ProductService/CreateProduct"
	ProductService_UpdateProduct_FullMethodName      = "/eolymp.commerce.ProductService/UpdateProduct"
	ProductService_DescribeProduct_FullMethodName    = "/eolymp.commerce.ProductService/DescribeProduct"
	ProductService_DeleteProduct_FullMethodName      = "/eolymp.commerce.ProductService/DeleteProduct"
	ProductService_ListProductPrices_FullMethodName  = "/eolymp.commerce.ProductService/ListProductPrices"
	ProductService_AddProductPrice_FullMethodName    = "/eolymp.commerce.ProductService/AddProductPrice"
	ProductService_RemoveProductPrice_FullMethodName = "/eolymp.commerce.ProductService/RemoveProductPrice"
)

// ProductServiceClient is the client API for ProductService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProductServiceClient interface {
	CreateProduct(ctx context.Context, in *CreateProductInput, opts ...grpc.CallOption) (*CreateProductOutput, error)
	UpdateProduct(ctx context.Context, in *UpdateProductInput, opts ...grpc.CallOption) (*UpdateProductOutput, error)
	DescribeProduct(ctx context.Context, in *DescribeProductInput, opts ...grpc.CallOption) (*DescribeProductOutput, error)
	DeleteProduct(ctx context.Context, in *DeleteProductInput, opts ...grpc.CallOption) (*DeleteProductOutput, error)
	ListProductPrices(ctx context.Context, in *ListProductPricesInput, opts ...grpc.CallOption) (*ListProductPricesOutput, error)
	AddProductPrice(ctx context.Context, in *AddProductPriceInput, opts ...grpc.CallOption) (*AddProductPriceOutput, error)
	RemoveProductPrice(ctx context.Context, in *RemoveProductPriceInput, opts ...grpc.CallOption) (*RemoveProductPriceOutput, error)
}

type productServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewProductServiceClient(cc grpc.ClientConnInterface) ProductServiceClient {
	return &productServiceClient{cc}
}

func (c *productServiceClient) CreateProduct(ctx context.Context, in *CreateProductInput, opts ...grpc.CallOption) (*CreateProductOutput, error) {
	out := new(CreateProductOutput)
	err := c.cc.Invoke(ctx, ProductService_CreateProduct_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productServiceClient) UpdateProduct(ctx context.Context, in *UpdateProductInput, opts ...grpc.CallOption) (*UpdateProductOutput, error) {
	out := new(UpdateProductOutput)
	err := c.cc.Invoke(ctx, ProductService_UpdateProduct_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productServiceClient) DescribeProduct(ctx context.Context, in *DescribeProductInput, opts ...grpc.CallOption) (*DescribeProductOutput, error) {
	out := new(DescribeProductOutput)
	err := c.cc.Invoke(ctx, ProductService_DescribeProduct_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productServiceClient) DeleteProduct(ctx context.Context, in *DeleteProductInput, opts ...grpc.CallOption) (*DeleteProductOutput, error) {
	out := new(DeleteProductOutput)
	err := c.cc.Invoke(ctx, ProductService_DeleteProduct_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productServiceClient) ListProductPrices(ctx context.Context, in *ListProductPricesInput, opts ...grpc.CallOption) (*ListProductPricesOutput, error) {
	out := new(ListProductPricesOutput)
	err := c.cc.Invoke(ctx, ProductService_ListProductPrices_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productServiceClient) AddProductPrice(ctx context.Context, in *AddProductPriceInput, opts ...grpc.CallOption) (*AddProductPriceOutput, error) {
	out := new(AddProductPriceOutput)
	err := c.cc.Invoke(ctx, ProductService_AddProductPrice_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productServiceClient) RemoveProductPrice(ctx context.Context, in *RemoveProductPriceInput, opts ...grpc.CallOption) (*RemoveProductPriceOutput, error) {
	out := new(RemoveProductPriceOutput)
	err := c.cc.Invoke(ctx, ProductService_RemoveProductPrice_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProductServiceServer is the server API for ProductService service.
// All implementations should embed UnimplementedProductServiceServer
// for forward compatibility
type ProductServiceServer interface {
	CreateProduct(context.Context, *CreateProductInput) (*CreateProductOutput, error)
	UpdateProduct(context.Context, *UpdateProductInput) (*UpdateProductOutput, error)
	DescribeProduct(context.Context, *DescribeProductInput) (*DescribeProductOutput, error)
	DeleteProduct(context.Context, *DeleteProductInput) (*DeleteProductOutput, error)
	ListProductPrices(context.Context, *ListProductPricesInput) (*ListProductPricesOutput, error)
	AddProductPrice(context.Context, *AddProductPriceInput) (*AddProductPriceOutput, error)
	RemoveProductPrice(context.Context, *RemoveProductPriceInput) (*RemoveProductPriceOutput, error)
}

// UnimplementedProductServiceServer should be embedded to have forward compatible implementations.
type UnimplementedProductServiceServer struct {
}

func (UnimplementedProductServiceServer) CreateProduct(context.Context, *CreateProductInput) (*CreateProductOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateProduct not implemented")
}
func (UnimplementedProductServiceServer) UpdateProduct(context.Context, *UpdateProductInput) (*UpdateProductOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateProduct not implemented")
}
func (UnimplementedProductServiceServer) DescribeProduct(context.Context, *DescribeProductInput) (*DescribeProductOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DescribeProduct not implemented")
}
func (UnimplementedProductServiceServer) DeleteProduct(context.Context, *DeleteProductInput) (*DeleteProductOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteProduct not implemented")
}
func (UnimplementedProductServiceServer) ListProductPrices(context.Context, *ListProductPricesInput) (*ListProductPricesOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListProductPrices not implemented")
}
func (UnimplementedProductServiceServer) AddProductPrice(context.Context, *AddProductPriceInput) (*AddProductPriceOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddProductPrice not implemented")
}
func (UnimplementedProductServiceServer) RemoveProductPrice(context.Context, *RemoveProductPriceInput) (*RemoveProductPriceOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveProductPrice not implemented")
}

// UnsafeProductServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProductServiceServer will
// result in compilation errors.
type UnsafeProductServiceServer interface {
	mustEmbedUnimplementedProductServiceServer()
}

func RegisterProductServiceServer(s grpc.ServiceRegistrar, srv ProductServiceServer) {
	s.RegisterService(&ProductService_ServiceDesc, srv)
}

func _ProductService_CreateProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateProductInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).CreateProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProductService_CreateProduct_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).CreateProduct(ctx, req.(*CreateProductInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductService_UpdateProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateProductInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).UpdateProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProductService_UpdateProduct_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).UpdateProduct(ctx, req.(*UpdateProductInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductService_DescribeProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DescribeProductInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).DescribeProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProductService_DescribeProduct_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).DescribeProduct(ctx, req.(*DescribeProductInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductService_DeleteProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteProductInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).DeleteProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProductService_DeleteProduct_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).DeleteProduct(ctx, req.(*DeleteProductInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductService_ListProductPrices_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListProductPricesInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).ListProductPrices(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProductService_ListProductPrices_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).ListProductPrices(ctx, req.(*ListProductPricesInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductService_AddProductPrice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddProductPriceInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).AddProductPrice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProductService_AddProductPrice_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).AddProductPrice(ctx, req.(*AddProductPriceInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductService_RemoveProductPrice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveProductPriceInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).RemoveProductPrice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProductService_RemoveProductPrice_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).RemoveProductPrice(ctx, req.(*RemoveProductPriceInput))
	}
	return interceptor(ctx, in, info, handler)
}

// ProductService_ServiceDesc is the grpc.ServiceDesc for ProductService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ProductService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "eolymp.commerce.ProductService",
	HandlerType: (*ProductServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateProduct",
			Handler:    _ProductService_CreateProduct_Handler,
		},
		{
			MethodName: "UpdateProduct",
			Handler:    _ProductService_UpdateProduct_Handler,
		},
		{
			MethodName: "DescribeProduct",
			Handler:    _ProductService_DescribeProduct_Handler,
		},
		{
			MethodName: "DeleteProduct",
			Handler:    _ProductService_DeleteProduct_Handler,
		},
		{
			MethodName: "ListProductPrices",
			Handler:    _ProductService_ListProductPrices_Handler,
		},
		{
			MethodName: "AddProductPrice",
			Handler:    _ProductService_AddProductPrice_Handler,
		},
		{
			MethodName: "RemoveProductPrice",
			Handler:    _ProductService_RemoveProductPrice_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "eolymp/commerce/product_service.proto",
}
