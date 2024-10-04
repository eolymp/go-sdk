// Code generated by protoc-gen-go-server. DO NOT EDIT.
// See https://github.com/eolymp/contracts/tree/main/cmd/protoc-gen-go-server for more details.

package judge

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

// _ScoreboardService_HTTPReadQueryString parses body into proto.Message
func _ScoreboardService_HTTPReadQueryString(r *http.Request, v proto.Message) error {
	query := r.URL.Query().Get("q")
	if query == "" {
		return nil
	}

	if err := (protojson.UnmarshalOptions{DiscardUnknown: true}.Unmarshal([]byte(query), v)); err != nil {
		return err
	}

	return nil
}

// _ScoreboardService_HTTPReadRequestBody parses body into proto.Message
func _ScoreboardService_HTTPReadRequestBody(r *http.Request, v proto.Message) error {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}

	if err := (protojson.UnmarshalOptions{DiscardUnknown: true}.Unmarshal(data, v)); err != nil {
		return err
	}

	return nil
}

// _ScoreboardService_HTTPWriteResponse writes proto.Message to HTTP response
func _ScoreboardService_HTTPWriteResponse(w http.ResponseWriter, v proto.Message, h, t metadata.MD) {
	data, err := protojson.Marshal(v)
	if err != nil {
		_ScoreboardService_HTTPWriteErrorResponse(w, err)
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

// _ScoreboardService_HTTPWriteErrorResponse writes error to HTTP response with error status code
func _ScoreboardService_HTTPWriteErrorResponse(w http.ResponseWriter, e error) {
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

// _ScoreboardService_WebsocketErrorResponse writes error to websocket connection
func _ScoreboardService_WebsocketErrorResponse(conn *websocket.Conn, e error) {
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

// _ScoreboardService_WebsocketCodec implements protobuf codec for websockets package
var _ScoreboardService_WebsocketCodec = websocket.Codec{
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

// RegisterScoreboardServiceHttpHandlers adds handlers for for ScoreboardServiceClient
// This constructor creates http.Handler, the actual implementation might change at any moment
func RegisterScoreboardServiceHttpHandlers(router *mux.Router, prefix string, cli ScoreboardServiceClient) {
	router.Handle(prefix+"/scoreboard", _ScoreboardService_DescribeScoreboard_Rule0(cli)).
		Methods("GET").
		Name("eolymp.judge.ScoreboardService.DescribeScoreboard")
	router.Handle(prefix+"/scoreboard/rows", _ScoreboardService_ListScoreboardRows_Rule0(cli)).
		Methods("GET").
		Name("eolymp.judge.ScoreboardService.ListScoreboardRows")
}

func _ScoreboardService_DescribeScoreboard_Rule0(cli ScoreboardServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &DescribeScoreboardInput{}

		if err := _ScoreboardService_HTTPReadQueryString(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_ScoreboardService_HTTPWriteErrorResponse(w, err)
			return
		}

		var header, trailer metadata.MD

		out, err := cli.DescribeScoreboard(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_ScoreboardService_HTTPWriteErrorResponse(w, err)
			return
		}

		_ScoreboardService_HTTPWriteResponse(w, out, header, trailer)
	})
}

func _ScoreboardService_ListScoreboardRows_Rule0(cli ScoreboardServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &ListScoreboardRowsInput{}

		if err := _ScoreboardService_HTTPReadQueryString(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_ScoreboardService_HTTPWriteErrorResponse(w, err)
			return
		}

		var header, trailer metadata.MD

		out, err := cli.ListScoreboardRows(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_ScoreboardService_HTTPWriteErrorResponse(w, err)
			return
		}

		_ScoreboardService_HTTPWriteResponse(w, out, header, trailer)
	})
}

type _ScoreboardServiceHandler = func(ctx context.Context, in proto.Message) (proto.Message, error)
type _ScoreboardServiceMiddleware = func(ctx context.Context, method string, in proto.Message, handler _ScoreboardServiceHandler) (out proto.Message, err error)
type ScoreboardServiceInterceptor struct {
	middleware []_ScoreboardServiceMiddleware
	client     ScoreboardServiceClient
}

// NewScoreboardServiceInterceptor constructs additional middleware for a server based on annotations in proto files
func NewScoreboardServiceInterceptor(cli ScoreboardServiceClient, middleware ..._ScoreboardServiceMiddleware) *ScoreboardServiceInterceptor {
	return &ScoreboardServiceInterceptor{client: cli, middleware: middleware}
}

func (i *ScoreboardServiceInterceptor) DescribeScoreboard(ctx context.Context, in *DescribeScoreboardInput, opts ...grpc.CallOption) (*DescribeScoreboardOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*DescribeScoreboardInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *DescribeScoreboardInput, got %T", in))
		}

		return i.client.DescribeScoreboard(ctx, message, opts...)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.judge.ScoreboardService.DescribeScoreboard", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*DescribeScoreboardOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *DescribeScoreboardOutput, got %T", out))
	}

	return message, err
}

func (i *ScoreboardServiceInterceptor) ListScoreboardRows(ctx context.Context, in *ListScoreboardRowsInput, opts ...grpc.CallOption) (*ListScoreboardRowsOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*ListScoreboardRowsInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *ListScoreboardRowsInput, got %T", in))
		}

		return i.client.ListScoreboardRows(ctx, message, opts...)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.judge.ScoreboardService.ListScoreboardRows", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*ListScoreboardRowsOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *ListScoreboardRowsOutput, got %T", out))
	}

	return message, err
}
