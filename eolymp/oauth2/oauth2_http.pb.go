// Code generated by protoc-gen-go-server. DO NOT EDIT.
// See https://github.com/eolymp/contracts/tree/main/cmd/protoc-gen-go-server for more details.

package oauth2

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

// _OAuth2_HTTPReadQueryString parses body into proto.Message
func _OAuth2_HTTPReadQueryString(r *http.Request, v proto.Message) error {
	query := r.URL.Query().Get("q")
	if query == "" {
		return nil
	}

	if err := (protojson.UnmarshalOptions{DiscardUnknown: true}.Unmarshal([]byte(query), v)); err != nil {
		return err
	}

	return nil
}

// _OAuth2_HTTPReadRequestBody parses body into proto.Message
func _OAuth2_HTTPReadRequestBody(r *http.Request, v proto.Message) error {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}

	if err := (protojson.UnmarshalOptions{DiscardUnknown: true}.Unmarshal(data, v)); err != nil {
		return err
	}

	return nil
}

// _OAuth2_HTTPWriteResponse writes proto.Message to HTTP response
func _OAuth2_HTTPWriteResponse(w http.ResponseWriter, v proto.Message) {
	data, err := protojson.Marshal(v)
	if err != nil {
		_OAuth2_HTTPWriteErrorResponse(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	_, _ = w.Write(data)
}

// _OAuth2_HTTPWriteErrorResponse writes error to HTTP response with error status code
func _OAuth2_HTTPWriteErrorResponse(w http.ResponseWriter, e error) {
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

// _OAuth2_WebsocketErrorResponse writes error to websocket connection
func _OAuth2_WebsocketErrorResponse(conn *websocket.Conn, e error) {
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

// _OAuth2_WebsocketCodec implements protobuf codec for websockets package
var _OAuth2_WebsocketCodec = websocket.Codec{
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

// RegisterOAuth2HttpHandlers adds handlers for for OAuth2Client
// This constructor creates http.Handler, the actual implementation might change at any moment
func RegisterOAuth2HttpHandlers(router *mux.Router, prefix string, cli OAuth2Client) {
}

type _OAuth2Handler = func(ctx context.Context, in proto.Message) (proto.Message, error)
type _OAuth2Middleware = func(ctx context.Context, method string, in proto.Message, handler _OAuth2Handler) (out proto.Message, err error)
type OAuth2Interceptor struct {
	middleware []_OAuth2Middleware
	client     OAuth2Client
}

// NewOAuth2Interceptor constructs additional middleware for a server based on annotations in proto files
func NewOAuth2Interceptor(cli OAuth2Client, middleware ..._OAuth2Middleware) *OAuth2Interceptor {
	return &OAuth2Interceptor{client: cli, middleware: middleware}
}

func (i *OAuth2Interceptor) Token(ctx context.Context, in *TokenInput, opts ...grpc.CallOption) (*TokenOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*TokenInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *TokenInput, got %T", in))
		}

		return i.client.Token(ctx, message, opts...)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.oauth2.OAuth2.Token", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*TokenOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *TokenOutput, got %T", out))
	}

	return message, err
}

func (i *OAuth2Interceptor) UserInfo(ctx context.Context, in *UserInfoInput, opts ...grpc.CallOption) (*UserInfoOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*UserInfoInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *UserInfoInput, got %T", in))
		}

		return i.client.UserInfo(ctx, message, opts...)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.oauth2.OAuth2.UserInfo", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*UserInfoOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *UserInfoOutput, got %T", out))
	}

	return message, err
}

func (i *OAuth2Interceptor) Introspect(ctx context.Context, in *IntrospectInput, opts ...grpc.CallOption) (*IntrospectOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*IntrospectInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *IntrospectInput, got %T", in))
		}

		return i.client.Introspect(ctx, message, opts...)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.oauth2.OAuth2.Introspect", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*IntrospectOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *IntrospectOutput, got %T", out))
	}

	return message, err
}

func (i *OAuth2Interceptor) Revoke(ctx context.Context, in *RevokeInput, opts ...grpc.CallOption) (*RevokeOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*RevokeInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *RevokeInput, got %T", in))
		}

		return i.client.Revoke(ctx, message, opts...)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.oauth2.OAuth2.Revoke", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*RevokeOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *RevokeOutput, got %T", out))
	}

	return message, err
}

func (i *OAuth2Interceptor) AuthCode(ctx context.Context, in *AuthCodeInput, opts ...grpc.CallOption) (*AuthCodeOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*AuthCodeInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *AuthCodeInput, got %T", in))
		}

		return i.client.AuthCode(ctx, message, opts...)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.oauth2.OAuth2.AuthCode", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*AuthCodeOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *AuthCodeOutput, got %T", out))
	}

	return message, err
}

func (i *OAuth2Interceptor) Authorize(ctx context.Context, in *AuthorizeInput, opts ...grpc.CallOption) (*AuthorizeOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*AuthorizeInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *AuthorizeInput, got %T", in))
		}

		return i.client.Authorize(ctx, message, opts...)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.oauth2.OAuth2.Authorize", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*AuthorizeOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *AuthorizeOutput, got %T", out))
	}

	return message, err
}

func (i *OAuth2Interceptor) Callback(ctx context.Context, in *CallbackInput, opts ...grpc.CallOption) (*CallbackOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*CallbackInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *CallbackInput, got %T", in))
		}

		return i.client.Callback(ctx, message, opts...)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.oauth2.OAuth2.Callback", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*CallbackOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *CallbackOutput, got %T", out))
	}

	return message, err
}
