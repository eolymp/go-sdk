// Code generated by protoc-gen-go-server. DO NOT EDIT.
// See https://github.com/eolymp/contracts/tree/main/cmd/protoc-gen-go-server for more details.

package atlas

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

// _SubmissionAssistantService_HTTPReadQueryString parses body into proto.Message
func _SubmissionAssistantService_HTTPReadQueryString(r *http.Request, v proto.Message) error {
	query := r.URL.Query().Get("q")
	if query == "" {
		return nil
	}

	if err := (protojson.UnmarshalOptions{DiscardUnknown: true}.Unmarshal([]byte(query), v)); err != nil {
		return err
	}

	return nil
}

// _SubmissionAssistantService_HTTPReadRequestBody parses body into proto.Message
func _SubmissionAssistantService_HTTPReadRequestBody(r *http.Request, v proto.Message) error {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}

	if err := (protojson.UnmarshalOptions{DiscardUnknown: true}.Unmarshal(data, v)); err != nil {
		return err
	}

	return nil
}

// _SubmissionAssistantService_HTTPWriteResponse writes proto.Message to HTTP response
func _SubmissionAssistantService_HTTPWriteResponse(w http.ResponseWriter, v proto.Message, h, t metadata.MD) {
	data, err := protojson.Marshal(v)
	if err != nil {
		_SubmissionAssistantService_HTTPWriteErrorResponse(w, err)
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

// _SubmissionAssistantService_HTTPWriteErrorResponse writes error to HTTP response with error status code
func _SubmissionAssistantService_HTTPWriteErrorResponse(w http.ResponseWriter, e error) {
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

// _SubmissionAssistantService_WebsocketErrorResponse writes error to websocket connection
func _SubmissionAssistantService_WebsocketErrorResponse(conn *websocket.Conn, e error) {
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

// _SubmissionAssistantService_WebsocketCodec implements protobuf codec for websockets package
var _SubmissionAssistantService_WebsocketCodec = websocket.Codec{
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

// RegisterSubmissionAssistantServiceHttpHandlers adds handlers for for SubmissionAssistantServiceClient
// This constructor creates http.Handler, the actual implementation might change at any moment
func RegisterSubmissionAssistantServiceHttpHandlers(router *mux.Router, prefix string, cli SubmissionAssistantServiceClient) {
	router.Handle(prefix+"/submissions/{submission_id}/assistant:debug", _SubmissionAssistantService_DebugSubmission_Rule0(cli)).
		Methods("POST").
		Name("eolymp.atlas.SubmissionAssistantService.DebugSubmission")
	router.Handle(prefix+"/submissions/{submission_id}/assistant:rate", _SubmissionAssistantService_RateDebugAssistance_Rule0(cli)).
		Methods("POST").
		Name("eolymp.atlas.SubmissionAssistantService.RateDebugAssistance")
}

func _SubmissionAssistantService_DebugSubmission_Rule0(cli SubmissionAssistantServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &DebugSubmissionInput{}

		if err := _SubmissionAssistantService_HTTPReadRequestBody(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_SubmissionAssistantService_HTTPWriteErrorResponse(w, err)
			return
		}

		vars := mux.Vars(r)
		in.SubmissionId = vars["submission_id"]

		var header, trailer metadata.MD

		out, err := cli.DebugSubmission(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_SubmissionAssistantService_HTTPWriteErrorResponse(w, err)
			return
		}

		_SubmissionAssistantService_HTTPWriteResponse(w, out, header, trailer)
	})
}

func _SubmissionAssistantService_RateDebugAssistance_Rule0(cli SubmissionAssistantServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &RateDebugAssistanceInput{}

		if err := _SubmissionAssistantService_HTTPReadRequestBody(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_SubmissionAssistantService_HTTPWriteErrorResponse(w, err)
			return
		}

		vars := mux.Vars(r)
		in.SubmissionId = vars["submission_id"]

		var header, trailer metadata.MD

		out, err := cli.RateDebugAssistance(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_SubmissionAssistantService_HTTPWriteErrorResponse(w, err)
			return
		}

		_SubmissionAssistantService_HTTPWriteResponse(w, out, header, trailer)
	})
}

type _SubmissionAssistantServiceHandler = func(ctx context.Context, in proto.Message) (proto.Message, error)
type _SubmissionAssistantServiceMiddleware = func(ctx context.Context, method string, in proto.Message, handler _SubmissionAssistantServiceHandler) (out proto.Message, err error)
type SubmissionAssistantServiceInterceptor struct {
	middleware []_SubmissionAssistantServiceMiddleware
	client     SubmissionAssistantServiceClient
}

// NewSubmissionAssistantServiceInterceptor constructs additional middleware for a server based on annotations in proto files
func NewSubmissionAssistantServiceInterceptor(cli SubmissionAssistantServiceClient, middleware ..._SubmissionAssistantServiceMiddleware) *SubmissionAssistantServiceInterceptor {
	return &SubmissionAssistantServiceInterceptor{client: cli, middleware: middleware}
}

func (i *SubmissionAssistantServiceInterceptor) DebugSubmission(ctx context.Context, in *DebugSubmissionInput, opts ...grpc.CallOption) (*DebugSubmissionOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*DebugSubmissionInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *DebugSubmissionInput, got %T", in))
		}

		return i.client.DebugSubmission(ctx, message, opts...)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.atlas.SubmissionAssistantService.DebugSubmission", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*DebugSubmissionOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *DebugSubmissionOutput, got %T", out))
	}

	return message, err
}

func (i *SubmissionAssistantServiceInterceptor) RateDebugAssistance(ctx context.Context, in *RateDebugAssistanceInput, opts ...grpc.CallOption) (*RateDebugAssistanceOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*RateDebugAssistanceInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *RateDebugAssistanceInput, got %T", in))
		}

		return i.client.RateDebugAssistance(ctx, message, opts...)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.atlas.SubmissionAssistantService.RateDebugAssistance", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*RateDebugAssistanceOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *RateDebugAssistanceOutput, got %T", out))
	}

	return message, err
}
