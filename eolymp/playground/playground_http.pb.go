// Code generated by protoc-gen-go-server. DO NOT EDIT.
// See https://github.com/eolymp/contracts/tree/main/cmd/protoc-gen-go-server for more details.

package playground

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

// _Playground_HTTPReadQueryString parses body into proto.Message
func _Playground_HTTPReadQueryString(r *http.Request, v proto.Message) error {
	query := r.URL.Query().Get("q")
	if query == "" {
		return nil
	}

	if err := (protojson.UnmarshalOptions{DiscardUnknown: true}.Unmarshal([]byte(query), v)); err != nil {
		return err
	}

	return nil
}

// _Playground_HTTPReadRequestBody parses body into proto.Message
func _Playground_HTTPReadRequestBody(r *http.Request, v proto.Message) error {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}

	if err := (protojson.UnmarshalOptions{DiscardUnknown: true}.Unmarshal(data, v)); err != nil {
		return err
	}

	return nil
}

// _Playground_HTTPWriteResponse writes proto.Message to HTTP response
func _Playground_HTTPWriteResponse(w http.ResponseWriter, v proto.Message) {
	data, err := protojson.Marshal(v)
	if err != nil {
		_Playground_HTTPWriteErrorResponse(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	_, _ = w.Write(data)
}

// _Playground_HTTPWriteErrorResponse writes error to HTTP response with error status code
func _Playground_HTTPWriteErrorResponse(w http.ResponseWriter, e error) {
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

// _Playground_WebsocketErrorResponse writes error to websocket connection
func _Playground_WebsocketErrorResponse(conn *websocket.Conn, e error) {
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

// _Playground_WebsocketCodec implements protobuf codec for websockets package
var _Playground_WebsocketCodec = websocket.Codec{
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

type _Playground_WatchRun_WSStream struct {
	ctx  context.Context
	conn *websocket.Conn
}

func (s *_Playground_WatchRun_WSStream) Send(m *WatchRunOutput) error {
	return s.SendMsg(m)
}

func (s *_Playground_WatchRun_WSStream) SetHeader(metadata.MD) error {
	return nil
}

func (s *_Playground_WatchRun_WSStream) SendHeader(metadata.MD) error {
	return nil
}

func (s *_Playground_WatchRun_WSStream) SetTrailer(metadata.MD) {
}

func (s *_Playground_WatchRun_WSStream) Context() context.Context {
	return s.ctx
}

func (s *_Playground_WatchRun_WSStream) SendMsg(m interface{}) error {
	return _Playground_WebsocketCodec.Send(s.conn, m)
}

func (s *_Playground_WatchRun_WSStream) RecvMsg(m interface{}) error {
	return nil
}

// RegisterPlaygroundHttpHandlers adds handlers for for PlaygroundClient
// This constructor creates http.Handler, the actual implementation might change at any moment
func RegisterPlaygroundHttpHandlers(router *mux.Router, prefix string, cli PlaygroundClient) {
	router.Handle(prefix+"/runs", _Playground_CreateRun_Rule0(cli)).
		Methods("POST").
		Name("eolymp.playground.Playground.CreateRun")
	router.Handle(prefix+"/runs/{run_id}", _Playground_DescribeRun_Rule0(cli)).
		Methods("GET").
		Name("eolymp.playground.Playground.DescribeRun")
	router.Handle(prefix+"/runs/{run_id}/watch", _Playground_WatchRun_Rule0(cli)).
		Methods("GET").
		Name("eolymp.playground.Playground.WatchRun")
}

func _Playground_CreateRun_Rule0(cli PlaygroundClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &CreateRunInput{}

		if err := _Playground_HTTPReadRequestBody(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_Playground_HTTPWriteErrorResponse(w, err)
			return
		}

		out, err := cli.CreateRun(r.Context(), in)
		if err != nil {
			_Playground_HTTPWriteErrorResponse(w, err)
			return
		}

		_Playground_HTTPWriteResponse(w, out)
	})
}

func _Playground_DescribeRun_Rule0(cli PlaygroundClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &DescribeRunInput{}

		if err := _Playground_HTTPReadQueryString(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_Playground_HTTPWriteErrorResponse(w, err)
			return
		}

		vars := mux.Vars(r)
		in.RunId = vars["run_id"]

		out, err := cli.DescribeRun(r.Context(), in)
		if err != nil {
			_Playground_HTTPWriteErrorResponse(w, err)
			return
		}

		_Playground_HTTPWriteResponse(w, out)
	})
}

func _Playground_WatchRun_Rule0(cli PlaygroundClient) http.Handler {
	return websocket.Handler(func(ws *websocket.Conn) {
		if err := ws.WriteClose(1000); err != nil {
			panic(err)
		}
	})
}

type _PlaygroundHandler = func(ctx context.Context, in proto.Message) (proto.Message, error)
type _PlaygroundMiddleware = func(ctx context.Context, method string, in proto.Message, handler _PlaygroundHandler) (out proto.Message, err error)
type PlaygroundInterceptor struct {
	middleware []_PlaygroundMiddleware
	client     PlaygroundClient
}

// NewPlaygroundInterceptor constructs additional middleware for a server based on annotations in proto files
func NewPlaygroundInterceptor(cli PlaygroundClient, middleware ..._PlaygroundMiddleware) *PlaygroundInterceptor {
	return &PlaygroundInterceptor{client: cli, middleware: middleware}
}

func (i *PlaygroundInterceptor) CreateRun(ctx context.Context, in *CreateRunInput, opts ...grpc.CallOption) (*CreateRunOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*CreateRunInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *CreateRunInput, got %T", in))
		}

		return i.client.CreateRun(ctx, message, opts...)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.playground.Playground.CreateRun", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*CreateRunOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *CreateRunOutput, got %T", out))
	}

	return message, err
}

func (i *PlaygroundInterceptor) DescribeRun(ctx context.Context, in *DescribeRunInput, opts ...grpc.CallOption) (*DescribeRunOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*DescribeRunInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *DescribeRunInput, got %T", in))
		}

		return i.client.DescribeRun(ctx, message, opts...)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.playground.Playground.DescribeRun", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*DescribeRunOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *DescribeRunOutput, got %T", out))
	}

	return message, err
}

func (i *PlaygroundInterceptor) WatchRun(ctx context.Context, in *WatchRunInput, opts ...grpc.CallOption) (Playground_WatchRunClient, error) {
	return i.client.WatchRun(ctx, in, opts...)
}
