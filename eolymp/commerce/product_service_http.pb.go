// Code generated by protoc-gen-go-server. DO NOT EDIT.
// See https://github.com/eolymp/contracts/tree/main/cmd/protoc-gen-go-server for more details.

package commerce

import (
	context "context"
	fmt "fmt"
	mux "github.com/gorilla/mux"
	websocket "golang.org/x/net/websocket"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protojson "google.golang.org/protobuf/encoding/protojson"
	proto "google.golang.org/protobuf/proto"
	ioutil "io/ioutil"
	http "net/http"
)

// _ProductService_HTTPReadQueryString parses body into proto.Message
func _ProductService_HTTPReadQueryString(r *http.Request, v proto.Message) error {
	query := r.URL.Query().Get("q")
	if query == "" {
		return nil
	}

	if err := (protojson.UnmarshalOptions{DiscardUnknown: true}.Unmarshal([]byte(query), v)); err != nil {
		return err
	}

	return nil
}

// _ProductService_HTTPReadRequestBody parses body into proto.Message
func _ProductService_HTTPReadRequestBody(r *http.Request, v proto.Message) error {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}

	if err := (protojson.UnmarshalOptions{DiscardUnknown: true}.Unmarshal(data, v)); err != nil {
		return err
	}

	return nil
}

// _ProductService_HTTPWriteResponse writes proto.Message to HTTP response
func _ProductService_HTTPWriteResponse(w http.ResponseWriter, v proto.Message) {
	data, err := protojson.Marshal(v)
	if err != nil {
		_ProductService_HTTPWriteErrorResponse(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	_, _ = w.Write(data)
}

// _ProductService_HTTPWriteErrorResponse writes error to HTTP response with error status code
func _ProductService_HTTPWriteErrorResponse(w http.ResponseWriter, e error) {
	s := status.Convert(e)

	w.Header().Set("Content-Type", "application/json")

	switch s.Code() {
	case codes.OK:
		w.WriteHeader(http.StatusOK)
	case codes.Canceled:
		w.WriteHeader(http.StatusGatewayTimeout)
	case codes.Unknown:
		w.WriteHeader(http.StatusInternalServerError)
	case codes.InvalidArgument:
		w.WriteHeader(http.StatusBadRequest)
	case codes.DeadlineExceeded:
		w.WriteHeader(http.StatusGatewayTimeout)
	case codes.NotFound:
		w.WriteHeader(http.StatusNotFound)
	case codes.AlreadyExists:
		w.WriteHeader(http.StatusConflict)
	case codes.PermissionDenied:
		w.WriteHeader(http.StatusForbidden)
	case codes.ResourceExhausted:
		w.WriteHeader(http.StatusTooManyRequests)
	case codes.FailedPrecondition:
		w.WriteHeader(http.StatusPreconditionFailed)
	case codes.Aborted:
		w.WriteHeader(http.StatusServiceUnavailable)
	case codes.OutOfRange:
		w.WriteHeader(http.StatusRequestedRangeNotSatisfiable)
	case codes.Unimplemented:
		w.WriteHeader(http.StatusNotImplemented)
	case codes.Internal:
		w.WriteHeader(http.StatusInternalServerError)
	case codes.Unavailable:
		w.WriteHeader(http.StatusServiceUnavailable)
	case codes.DataLoss:
		w.WriteHeader(http.StatusInternalServerError)
	case codes.Unauthenticated:
		w.WriteHeader(http.StatusUnauthorized)
	default:
		w.WriteHeader(http.StatusInternalServerError)
	}

	data, err := protojson.Marshal(s.Proto())
	if err != nil {
		panic(err)
	}

	_, _ = w.Write(data)
}

// _ProductService_WebsocketErrorResponse writes error to websocket connection
func _ProductService_WebsocketErrorResponse(conn *websocket.Conn, e error) {
	switch status.Convert(e).Code() {
	case codes.OK:
		conn.WriteClose(1000)
	case codes.Canceled:
		conn.WriteClose(1000)
	case codes.Unknown:
		conn.WriteClose(1011)
	case codes.InvalidArgument:
		conn.WriteClose(1003)
	case codes.DeadlineExceeded:
		conn.WriteClose(1000)
	case codes.NotFound:
		conn.WriteClose(1000)
	case codes.AlreadyExists:
		conn.WriteClose(1000)
	case codes.PermissionDenied:
		conn.WriteClose(1000)
	case codes.ResourceExhausted:
		conn.WriteClose(1000)
	case codes.FailedPrecondition:
		conn.WriteClose(1000)
	case codes.Aborted:
		conn.WriteClose(1000)
	case codes.OutOfRange:
		conn.WriteClose(1000)
	case codes.Unimplemented:
		conn.WriteClose(1011)
	case codes.Internal:
		conn.WriteClose(1011)
	case codes.Unavailable:
		conn.WriteClose(1011)
	case codes.DataLoss:
		conn.WriteClose(1011)
	case codes.Unauthenticated:
		conn.WriteClose(1000)
	default:
		conn.WriteClose(1000)
	}
}

// _ProductService_WebsocketCodec implements protobuf codec for websockets package
var _ProductService_WebsocketCodec = websocket.Codec{
	Marshal: func(v interface{}) ([]byte, byte, error) {
		m, ok := v.(proto.Message)
		if !ok {
			panic(fmt.Errorf("invalid message type %T", v))
		}

		d, err := protojson.Marshal(m)
		if err != nil {
			return nil, 0, err
		}

		return d, websocket.TextFrame, err
	},
	Unmarshal: func(d []byte, t byte, v interface{}) error {
		m, ok := v.(proto.Message)
		if !ok {
			panic(fmt.Errorf("invalid message type %T", v))
		}

		return protojson.UnmarshalOptions{DiscardUnknown: true}.Unmarshal(d, m)
	},
}

// RegisterProductServiceHttpHandlers adds handlers for for ProductServiceClient
// This constructor creates http.Handler, the actual implementation might change at any moment
func RegisterProductServiceHttpHandlers(router *mux.Router, prefix string, cli ProductServiceClient) {
}

type _ProductServiceHandler = func(ctx context.Context, in proto.Message) (proto.Message, error)
type _ProductServiceMiddleware = func(ctx context.Context, method string, in proto.Message, handler _ProductServiceHandler) (out proto.Message, err error)
type ProductServiceInterceptor struct {
	middleware []_ProductServiceMiddleware
	client     ProductServiceClient
}

// NewProductServiceInterceptor constructs additional middleware for a server based on annotations in proto files
func NewProductServiceInterceptor(cli ProductServiceClient, middleware ..._ProductServiceMiddleware) *ProductServiceInterceptor {
	return &ProductServiceInterceptor{client: cli, middleware: middleware}
}

func (i *ProductServiceInterceptor) CreateProduct(ctx context.Context, in *CreateProductInput, opts ...grpc.CallOption) (*CreateProductOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*CreateProductInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *CreateProductInput, got %T", in))
		}

		return i.client.CreateProduct(ctx, message, opts...)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.commerce.ProductService.CreateProduct", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*CreateProductOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *CreateProductOutput, got %T", out))
	}

	return message, err
}

func (i *ProductServiceInterceptor) UpdateProduct(ctx context.Context, in *UpdateProductInput, opts ...grpc.CallOption) (*UpdateProductOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*UpdateProductInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *UpdateProductInput, got %T", in))
		}

		return i.client.UpdateProduct(ctx, message, opts...)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.commerce.ProductService.UpdateProduct", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*UpdateProductOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *UpdateProductOutput, got %T", out))
	}

	return message, err
}

func (i *ProductServiceInterceptor) DescribeProduct(ctx context.Context, in *DescribeProductInput, opts ...grpc.CallOption) (*DescribeProductOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*DescribeProductInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *DescribeProductInput, got %T", in))
		}

		return i.client.DescribeProduct(ctx, message, opts...)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.commerce.ProductService.DescribeProduct", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*DescribeProductOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *DescribeProductOutput, got %T", out))
	}

	return message, err
}

func (i *ProductServiceInterceptor) ListProducts(ctx context.Context, in *ListProductsInput, opts ...grpc.CallOption) (*ListProductsOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*ListProductsInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *ListProductsInput, got %T", in))
		}

		return i.client.ListProducts(ctx, message, opts...)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.commerce.ProductService.ListProducts", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*ListProductsOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *ListProductsOutput, got %T", out))
	}

	return message, err
}

func (i *ProductServiceInterceptor) DeleteProduct(ctx context.Context, in *DeleteProductInput, opts ...grpc.CallOption) (*DeleteProductOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*DeleteProductInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *DeleteProductInput, got %T", in))
		}

		return i.client.DeleteProduct(ctx, message, opts...)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.commerce.ProductService.DeleteProduct", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*DeleteProductOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *DeleteProductOutput, got %T", out))
	}

	return message, err
}

func (i *ProductServiceInterceptor) ListProductPrices(ctx context.Context, in *ListProductPricesInput, opts ...grpc.CallOption) (*ListProductPricesOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*ListProductPricesInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *ListProductPricesInput, got %T", in))
		}

		return i.client.ListProductPrices(ctx, message, opts...)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.commerce.ProductService.ListProductPrices", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*ListProductPricesOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *ListProductPricesOutput, got %T", out))
	}

	return message, err
}

func (i *ProductServiceInterceptor) DescribeProductPrice(ctx context.Context, in *DescribeProductPriceInput, opts ...grpc.CallOption) (*DescribeProductPriceOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*DescribeProductPriceInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *DescribeProductPriceInput, got %T", in))
		}

		return i.client.DescribeProductPrice(ctx, message, opts...)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.commerce.ProductService.DescribeProductPrice", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*DescribeProductPriceOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *DescribeProductPriceOutput, got %T", out))
	}

	return message, err
}

func (i *ProductServiceInterceptor) CreateProductPrice(ctx context.Context, in *CreateProductPriceInput, opts ...grpc.CallOption) (*CreateProductPriceOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*CreateProductPriceInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *CreateProductPriceInput, got %T", in))
		}

		return i.client.CreateProductPrice(ctx, message, opts...)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.commerce.ProductService.CreateProductPrice", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*CreateProductPriceOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *CreateProductPriceOutput, got %T", out))
	}

	return message, err
}

func (i *ProductServiceInterceptor) DeleteProductPrice(ctx context.Context, in *DeleteProductPriceInput, opts ...grpc.CallOption) (*DeleteProductPriceOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*DeleteProductPriceInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *DeleteProductPriceInput, got %T", in))
		}

		return i.client.DeleteProductPrice(ctx, message, opts...)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.commerce.ProductService.DeleteProductPrice", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*DeleteProductPriceOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *DeleteProductPriceOutput, got %T", out))
	}

	return message, err
}
