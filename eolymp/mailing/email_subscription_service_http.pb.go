// Code generated by protoc-gen-go-server. DO NOT EDIT.
// See https://github.com/eolymp/contracts/tree/main/cmd/protoc-gen-go-server for more details.

package mailing

import (
	context "context"
	fmt "fmt"
	mux "github.com/gorilla/mux"
	websocket "golang.org/x/net/websocket"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	metadata "google.golang.org/grpc/metadata"
	status "google.golang.org/grpc/status"
	protojson "google.golang.org/protobuf/encoding/protojson"
	proto "google.golang.org/protobuf/proto"
	ioutil "io/ioutil"
	http "net/http"
)

// _EmailSubscriptionService_HTTPReadQueryString parses body into proto.Message
func _EmailSubscriptionService_HTTPReadQueryString(r *http.Request, v proto.Message) error {
	query := r.URL.Query().Get("q")
	if query == "" {
		return nil
	}

	if err := (protojson.UnmarshalOptions{DiscardUnknown: true}.Unmarshal([]byte(query), v)); err != nil {
		return err
	}

	return nil
}

// _EmailSubscriptionService_HTTPReadRequestBody parses body into proto.Message
func _EmailSubscriptionService_HTTPReadRequestBody(r *http.Request, v proto.Message) error {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}

	if err := (protojson.UnmarshalOptions{DiscardUnknown: true}.Unmarshal(data, v)); err != nil {
		return err
	}

	return nil
}

// _EmailSubscriptionService_HTTPWriteResponse writes proto.Message to HTTP response
func _EmailSubscriptionService_HTTPWriteResponse(w http.ResponseWriter, v proto.Message, h, t metadata.MD) {
	data, err := protojson.Marshal(v)
	if err != nil {
		_EmailSubscriptionService_HTTPWriteErrorResponse(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	if v := append(h.Get("cache-control"), t.Get("cache-control")...); len(v) > 0 {
		w.Header().Set("Cache-Control", v[len(v)-1])
	}

	if v := append(h.Get("etag"), t.Get("etag")...); len(v) > 0 {
		w.Header().Set("ETag", v[len(v)-1])
	}

	if v := append(h.Get("last-modified"), t.Get("last-modified")...); len(v) > 0 {
		w.Header().Set("Last-Modified", v[len(v)-1])
	}

	w.WriteHeader(http.StatusOK)

	_, _ = w.Write(data)
}

// _EmailSubscriptionService_HTTPWriteErrorResponse writes error to HTTP response with error status code
func _EmailSubscriptionService_HTTPWriteErrorResponse(w http.ResponseWriter, e error) {
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

// _EmailSubscriptionService_WebsocketErrorResponse writes error to websocket connection
func _EmailSubscriptionService_WebsocketErrorResponse(conn *websocket.Conn, e error) {
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

// _EmailSubscriptionService_WebsocketCodec implements protobuf codec for websockets package
var _EmailSubscriptionService_WebsocketCodec = websocket.Codec{
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

// RegisterEmailSubscriptionServiceHttpHandlers adds handlers for for EmailSubscriptionServiceClient
// This constructor creates http.Handler, the actual implementation might change at any moment
func RegisterEmailSubscriptionServiceHttpHandlers(router *mux.Router, prefix string, cli EmailSubscriptionServiceClient) {
	router.Handle(prefix+"/email-subscription", _EmailSubscriptionService_DescribeEmailSubscription_Rule0(cli)).
		Methods("GET").
		Name("eolymp.mailing.EmailSubscriptionService.DescribeEmailSubscription")
	router.Handle(prefix+"/email-subscription", _EmailSubscriptionService_UpdateEmailSubscription_Rule0(cli)).
		Methods("POST").
		Name("eolymp.mailing.EmailSubscriptionService.UpdateEmailSubscription")
}

func _EmailSubscriptionService_DescribeEmailSubscription_Rule0(cli EmailSubscriptionServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &DescribeEmailSubscriptionInput{}

		if err := _EmailSubscriptionService_HTTPReadQueryString(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_EmailSubscriptionService_HTTPWriteErrorResponse(w, err)
			return
		}

		var header, trailer metadata.MD

		out, err := cli.DescribeEmailSubscription(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_EmailSubscriptionService_HTTPWriteErrorResponse(w, err)
			return
		}

		_EmailSubscriptionService_HTTPWriteResponse(w, out, header, trailer)
	})
}

func _EmailSubscriptionService_UpdateEmailSubscription_Rule0(cli EmailSubscriptionServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &UpdateEmailSubscriptionInput{}

		if err := _EmailSubscriptionService_HTTPReadRequestBody(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_EmailSubscriptionService_HTTPWriteErrorResponse(w, err)
			return
		}

		var header, trailer metadata.MD

		out, err := cli.UpdateEmailSubscription(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_EmailSubscriptionService_HTTPWriteErrorResponse(w, err)
			return
		}

		_EmailSubscriptionService_HTTPWriteResponse(w, out, header, trailer)
	})
}

type _EmailSubscriptionServiceHandler = func(ctx context.Context, in proto.Message) (proto.Message, error)
type _EmailSubscriptionServiceMiddleware = func(ctx context.Context, method string, in proto.Message, handler _EmailSubscriptionServiceHandler) (out proto.Message, err error)
type EmailSubscriptionServiceInterceptor struct {
	middleware []_EmailSubscriptionServiceMiddleware
	client     EmailSubscriptionServiceClient
}

// NewEmailSubscriptionServiceInterceptor constructs additional middleware for a server based on annotations in proto files
func NewEmailSubscriptionServiceInterceptor(cli EmailSubscriptionServiceClient, middleware ..._EmailSubscriptionServiceMiddleware) *EmailSubscriptionServiceInterceptor {
	return &EmailSubscriptionServiceInterceptor{client: cli, middleware: middleware}
}

func (i *EmailSubscriptionServiceInterceptor) DescribeEmailSubscription(ctx context.Context, in *DescribeEmailSubscriptionInput, opts ...grpc.CallOption) (*DescribeEmailSubscriptionOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*DescribeEmailSubscriptionInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *DescribeEmailSubscriptionInput, got %T", in))
		}

		return i.client.DescribeEmailSubscription(ctx, message, opts...)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.mailing.EmailSubscriptionService.DescribeEmailSubscription", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*DescribeEmailSubscriptionOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *DescribeEmailSubscriptionOutput, got %T", out))
	}

	return message, err
}

func (i *EmailSubscriptionServiceInterceptor) UpdateEmailSubscription(ctx context.Context, in *UpdateEmailSubscriptionInput, opts ...grpc.CallOption) (*UpdateEmailSubscriptionOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*UpdateEmailSubscriptionInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *UpdateEmailSubscriptionInput, got %T", in))
		}

		return i.client.UpdateEmailSubscription(ctx, message, opts...)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.mailing.EmailSubscriptionService.UpdateEmailSubscription", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*UpdateEmailSubscriptionOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *UpdateEmailSubscriptionOutput, got %T", out))
	}

	return message, err
}