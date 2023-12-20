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

// _SubscriptionService_HTTPReadQueryString parses body into proto.Message
func _SubscriptionService_HTTPReadQueryString(r *http.Request, v proto.Message) error {
	query := r.URL.Query().Get("q")
	if query == "" {
		return nil
	}

	if err := (protojson.UnmarshalOptions{DiscardUnknown: true}.Unmarshal([]byte(query), v)); err != nil {
		return err
	}

	return nil
}

// _SubscriptionService_HTTPReadRequestBody parses body into proto.Message
func _SubscriptionService_HTTPReadRequestBody(r *http.Request, v proto.Message) error {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}

	if err := (protojson.UnmarshalOptions{DiscardUnknown: true}.Unmarshal(data, v)); err != nil {
		return err
	}

	return nil
}

// _SubscriptionService_HTTPWriteResponse writes proto.Message to HTTP response
func _SubscriptionService_HTTPWriteResponse(w http.ResponseWriter, v proto.Message) {
	data, err := protojson.Marshal(v)
	if err != nil {
		_SubscriptionService_HTTPWriteErrorResponse(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	_, _ = w.Write(data)
}

// _SubscriptionService_HTTPWriteErrorResponse writes error to HTTP response with error status code
func _SubscriptionService_HTTPWriteErrorResponse(w http.ResponseWriter, e error) {
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

// _SubscriptionService_WebsocketErrorResponse writes error to websocket connection
func _SubscriptionService_WebsocketErrorResponse(conn *websocket.Conn, e error) {
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

// _SubscriptionService_WebsocketCodec implements protobuf codec for websockets package
var _SubscriptionService_WebsocketCodec = websocket.Codec{
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

// RegisterSubscriptionServiceHttpHandlers adds handlers for for SubscriptionServiceClient
// This constructor creates http.Handler, the actual implementation might change at any moment
func RegisterSubscriptionServiceHttpHandlers(router *mux.Router, prefix string, cli SubscriptionServiceClient) {
}

type _SubscriptionServiceHandler = func(ctx context.Context, in proto.Message) (proto.Message, error)
type _SubscriptionServiceMiddleware = func(ctx context.Context, method string, in proto.Message, handler _SubscriptionServiceHandler) (out proto.Message, err error)
type SubscriptionServiceInterceptor struct {
	middleware []_SubscriptionServiceMiddleware
	client     SubscriptionServiceClient
}

// NewSubscriptionServiceInterceptor constructs additional middleware for a server based on annotations in proto files
func NewSubscriptionServiceInterceptor(cli SubscriptionServiceClient, middleware ..._SubscriptionServiceMiddleware) *SubscriptionServiceInterceptor {
	return &SubscriptionServiceInterceptor{client: cli, middleware: middleware}
}

func (i *SubscriptionServiceInterceptor) CreateSubscription(ctx context.Context, in *CreateSubscriptionInput, opts ...grpc.CallOption) (*CreateSubscriptionOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*CreateSubscriptionInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *CreateSubscriptionInput, got %T", in))
		}

		return i.client.CreateSubscription(ctx, message, opts...)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.commerce.SubscriptionService.CreateSubscription", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*CreateSubscriptionOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *CreateSubscriptionOutput, got %T", out))
	}

	return message, err
}

func (i *SubscriptionServiceInterceptor) UpdateSubscription(ctx context.Context, in *UpdateSubscriptionInput, opts ...grpc.CallOption) (*UpdateSubscriptionOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*UpdateSubscriptionInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *UpdateSubscriptionInput, got %T", in))
		}

		return i.client.UpdateSubscription(ctx, message, opts...)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.commerce.SubscriptionService.UpdateSubscription", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*UpdateSubscriptionOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *UpdateSubscriptionOutput, got %T", out))
	}

	return message, err
}

func (i *SubscriptionServiceInterceptor) CancelSubscription(ctx context.Context, in *CancelSubscriptionInput, opts ...grpc.CallOption) (*CancelSubscriptionOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*CancelSubscriptionInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *CancelSubscriptionInput, got %T", in))
		}

		return i.client.CancelSubscription(ctx, message, opts...)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.commerce.SubscriptionService.CancelSubscription", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*CancelSubscriptionOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *CancelSubscriptionOutput, got %T", out))
	}

	return message, err
}

func (i *SubscriptionServiceInterceptor) DescribeSubscription(ctx context.Context, in *DescribeSubscriptionInput, opts ...grpc.CallOption) (*DescribeSubscriptionOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*DescribeSubscriptionInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *DescribeSubscriptionInput, got %T", in))
		}

		return i.client.DescribeSubscription(ctx, message, opts...)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.commerce.SubscriptionService.DescribeSubscription", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*DescribeSubscriptionOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *DescribeSubscriptionOutput, got %T", out))
	}

	return message, err
}

func (i *SubscriptionServiceInterceptor) ListSubscriptions(ctx context.Context, in *ListSubscriptionsInput, opts ...grpc.CallOption) (*ListSubscriptionsOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*ListSubscriptionsInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *ListSubscriptionsInput, got %T", in))
		}

		return i.client.ListSubscriptions(ctx, message, opts...)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.commerce.SubscriptionService.ListSubscriptions", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*ListSubscriptionsOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *ListSubscriptionsOutput, got %T", out))
	}

	return message, err
}
