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

// _SubmissionService_HTTPReadQueryString parses body into proto.Message
func _SubmissionService_HTTPReadQueryString(r *http.Request, v proto.Message) error {
	query := r.URL.Query().Get("q")
	if query == "" {
		return nil
	}

	if err := (protojson.UnmarshalOptions{DiscardUnknown: true}.Unmarshal([]byte(query), v)); err != nil {
		return err
	}

	return nil
}

// _SubmissionService_HTTPReadRequestBody parses body into proto.Message
func _SubmissionService_HTTPReadRequestBody(r *http.Request, v proto.Message) error {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}

	if err := (protojson.UnmarshalOptions{DiscardUnknown: true}.Unmarshal(data, v)); err != nil {
		return err
	}

	return nil
}

// _SubmissionService_HTTPWriteResponse writes proto.Message to HTTP response
func _SubmissionService_HTTPWriteResponse(w http.ResponseWriter, v proto.Message, h, t metadata.MD) {
	data, err := protojson.Marshal(v)
	if err != nil {
		_SubmissionService_HTTPWriteErrorResponse(w, err)
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

// _SubmissionService_HTTPWriteErrorResponse writes error to HTTP response with error status code
func _SubmissionService_HTTPWriteErrorResponse(w http.ResponseWriter, e error) {
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

// _SubmissionService_WebsocketErrorResponse writes error to websocket connection
func _SubmissionService_WebsocketErrorResponse(conn *websocket.Conn, e error) {
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

// _SubmissionService_WebsocketCodec implements protobuf codec for websockets package
var _SubmissionService_WebsocketCodec = websocket.Codec{
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

type _SubmissionService_WatchSubmission_WSStream struct {
	ctx  context.Context
	conn *websocket.Conn
}

func (s *_SubmissionService_WatchSubmission_WSStream) Send(m *WatchSubmissionOutput) error {
	return s.SendMsg(m)
}

func (s *_SubmissionService_WatchSubmission_WSStream) SetHeader(metadata.MD) error {
	return nil
}

func (s *_SubmissionService_WatchSubmission_WSStream) SendHeader(metadata.MD) error {
	return nil
}

func (s *_SubmissionService_WatchSubmission_WSStream) SetTrailer(metadata.MD) {
}

func (s *_SubmissionService_WatchSubmission_WSStream) Context() context.Context {
	return s.ctx
}

func (s *_SubmissionService_WatchSubmission_WSStream) SendMsg(m interface{}) error {
	return _SubmissionService_WebsocketCodec.Send(s.conn, m)
}

func (s *_SubmissionService_WatchSubmission_WSStream) RecvMsg(m interface{}) error {
	return nil
}

// RegisterSubmissionServiceHttpHandlers adds handlers for for SubmissionServiceClient
// This constructor creates http.Handler, the actual implementation might change at any moment
func RegisterSubmissionServiceHttpHandlers(router *mux.Router, prefix string, cli SubmissionServiceClient) {
	router.Handle(prefix+"/problems/{problem_id}/submissions", _SubmissionService_CreateSubmission_Rule0(cli)).
		Methods("POST").
		Name("eolymp.judge.SubmissionService.CreateSubmission")
	router.Handle(prefix+"/submissions", _SubmissionService_ListSubmissions_Rule0(cli)).
		Methods("GET").
		Name("eolymp.judge.SubmissionService.ListSubmissions")
	router.Handle(prefix+"/submissions/{submission_id}", _SubmissionService_DescribeSubmission_Rule0(cli)).
		Methods("GET").
		Name("eolymp.judge.SubmissionService.DescribeSubmission")
	router.Handle(prefix+"/submissions/{submission_id}/retest", _SubmissionService_RetestSubmission_Rule0(cli)).
		Methods("POST").
		Name("eolymp.judge.SubmissionService.RetestSubmission")
	router.Handle(prefix+"/submissions/{submission_id}", _SubmissionService_DeleteSubmission_Rule0(cli)).
		Methods("DELETE").
		Name("eolymp.judge.SubmissionService.DeleteSubmission")
	router.Handle(prefix+"/submissions/{submission_id}/restore", _SubmissionService_RestoreSubmission_Rule0(cli)).
		Methods("POST").
		Name("eolymp.judge.SubmissionService.RestoreSubmission")
	router.Handle(prefix+"/problems/{problem_id}/retest", _SubmissionService_RetestProblem_Rule0(cli)).
		Methods("GET").
		Name("eolymp.judge.SubmissionService.RetestProblem")
}

func _SubmissionService_CreateSubmission_Rule0(cli SubmissionServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &CreateSubmissionInput{}

		if err := _SubmissionService_HTTPReadRequestBody(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_SubmissionService_HTTPWriteErrorResponse(w, err)
			return
		}

		vars := mux.Vars(r)
		in.ProblemId = vars["problem_id"]

		var header, trailer metadata.MD

		out, err := cli.CreateSubmission(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_SubmissionService_HTTPWriteErrorResponse(w, err)
			return
		}

		_SubmissionService_HTTPWriteResponse(w, out, header, trailer)
	})
}

func _SubmissionService_ListSubmissions_Rule0(cli SubmissionServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &ListSubmissionsInput{}

		if err := _SubmissionService_HTTPReadQueryString(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_SubmissionService_HTTPWriteErrorResponse(w, err)
			return
		}

		var header, trailer metadata.MD

		out, err := cli.ListSubmissions(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_SubmissionService_HTTPWriteErrorResponse(w, err)
			return
		}

		_SubmissionService_HTTPWriteResponse(w, out, header, trailer)
	})
}

func _SubmissionService_DescribeSubmission_Rule0(cli SubmissionServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &DescribeSubmissionInput{}

		if err := _SubmissionService_HTTPReadQueryString(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_SubmissionService_HTTPWriteErrorResponse(w, err)
			return
		}

		vars := mux.Vars(r)
		in.SubmissionId = vars["submission_id"]

		var header, trailer metadata.MD

		out, err := cli.DescribeSubmission(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_SubmissionService_HTTPWriteErrorResponse(w, err)
			return
		}

		_SubmissionService_HTTPWriteResponse(w, out, header, trailer)
	})
}

func _SubmissionService_RetestSubmission_Rule0(cli SubmissionServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &RetestSubmissionInput{}

		if err := _SubmissionService_HTTPReadRequestBody(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_SubmissionService_HTTPWriteErrorResponse(w, err)
			return
		}

		vars := mux.Vars(r)
		in.SubmissionId = vars["submission_id"]

		var header, trailer metadata.MD

		out, err := cli.RetestSubmission(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_SubmissionService_HTTPWriteErrorResponse(w, err)
			return
		}

		_SubmissionService_HTTPWriteResponse(w, out, header, trailer)
	})
}

func _SubmissionService_DeleteSubmission_Rule0(cli SubmissionServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &DeleteSubmissionInput{}

		if err := _SubmissionService_HTTPReadRequestBody(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_SubmissionService_HTTPWriteErrorResponse(w, err)
			return
		}

		vars := mux.Vars(r)
		in.SubmissionId = vars["submission_id"]

		var header, trailer metadata.MD

		out, err := cli.DeleteSubmission(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_SubmissionService_HTTPWriteErrorResponse(w, err)
			return
		}

		_SubmissionService_HTTPWriteResponse(w, out, header, trailer)
	})
}

func _SubmissionService_RestoreSubmission_Rule0(cli SubmissionServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &RestoreSubmissionInput{}

		if err := _SubmissionService_HTTPReadRequestBody(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_SubmissionService_HTTPWriteErrorResponse(w, err)
			return
		}

		vars := mux.Vars(r)
		in.SubmissionId = vars["submission_id"]

		var header, trailer metadata.MD

		out, err := cli.RestoreSubmission(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_SubmissionService_HTTPWriteErrorResponse(w, err)
			return
		}

		_SubmissionService_HTTPWriteResponse(w, out, header, trailer)
	})
}

func _SubmissionService_RetestProblem_Rule0(cli SubmissionServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &RetestProblemInput{}

		if err := _SubmissionService_HTTPReadQueryString(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_SubmissionService_HTTPWriteErrorResponse(w, err)
			return
		}

		vars := mux.Vars(r)
		in.ProblemId = vars["problem_id"]

		var header, trailer metadata.MD

		out, err := cli.RetestProblem(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_SubmissionService_HTTPWriteErrorResponse(w, err)
			return
		}

		_SubmissionService_HTTPWriteResponse(w, out, header, trailer)
	})
}

type _SubmissionServiceHandler = func(ctx context.Context, in proto.Message) (proto.Message, error)
type _SubmissionServiceMiddleware = func(ctx context.Context, method string, in proto.Message, handler _SubmissionServiceHandler) (out proto.Message, err error)
type SubmissionServiceInterceptor struct {
	middleware []_SubmissionServiceMiddleware
	client     SubmissionServiceClient
}

// NewSubmissionServiceInterceptor constructs additional middleware for a server based on annotations in proto files
func NewSubmissionServiceInterceptor(cli SubmissionServiceClient, middleware ..._SubmissionServiceMiddleware) *SubmissionServiceInterceptor {
	return &SubmissionServiceInterceptor{client: cli, middleware: middleware}
}

func (i *SubmissionServiceInterceptor) CreateSubmission(ctx context.Context, in *CreateSubmissionInput, opts ...grpc.CallOption) (*CreateSubmissionOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*CreateSubmissionInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *CreateSubmissionInput, got %T", in))
		}

		return i.client.CreateSubmission(ctx, message, opts...)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.judge.SubmissionService.CreateSubmission", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*CreateSubmissionOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *CreateSubmissionOutput, got %T", out))
	}

	return message, err
}

func (i *SubmissionServiceInterceptor) ListSubmissions(ctx context.Context, in *ListSubmissionsInput, opts ...grpc.CallOption) (*ListSubmissionsOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*ListSubmissionsInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *ListSubmissionsInput, got %T", in))
		}

		return i.client.ListSubmissions(ctx, message, opts...)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.judge.SubmissionService.ListSubmissions", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*ListSubmissionsOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *ListSubmissionsOutput, got %T", out))
	}

	return message, err
}

func (i *SubmissionServiceInterceptor) DescribeSubmission(ctx context.Context, in *DescribeSubmissionInput, opts ...grpc.CallOption) (*DescribeSubmissionOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*DescribeSubmissionInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *DescribeSubmissionInput, got %T", in))
		}

		return i.client.DescribeSubmission(ctx, message, opts...)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.judge.SubmissionService.DescribeSubmission", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*DescribeSubmissionOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *DescribeSubmissionOutput, got %T", out))
	}

	return message, err
}

func (i *SubmissionServiceInterceptor) WatchSubmission(ctx context.Context, in *WatchSubmissionInput, opts ...grpc.CallOption) (SubmissionService_WatchSubmissionClient, error) {
	return i.client.WatchSubmission(ctx, in, opts...)
}

func (i *SubmissionServiceInterceptor) RetestSubmission(ctx context.Context, in *RetestSubmissionInput, opts ...grpc.CallOption) (*RetestSubmissionOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*RetestSubmissionInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *RetestSubmissionInput, got %T", in))
		}

		return i.client.RetestSubmission(ctx, message, opts...)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.judge.SubmissionService.RetestSubmission", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*RetestSubmissionOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *RetestSubmissionOutput, got %T", out))
	}

	return message, err
}

func (i *SubmissionServiceInterceptor) DeleteSubmission(ctx context.Context, in *DeleteSubmissionInput, opts ...grpc.CallOption) (*DeleteSubmissionOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*DeleteSubmissionInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *DeleteSubmissionInput, got %T", in))
		}

		return i.client.DeleteSubmission(ctx, message, opts...)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.judge.SubmissionService.DeleteSubmission", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*DeleteSubmissionOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *DeleteSubmissionOutput, got %T", out))
	}

	return message, err
}

func (i *SubmissionServiceInterceptor) RestoreSubmission(ctx context.Context, in *RestoreSubmissionInput, opts ...grpc.CallOption) (*RestoreSubmissionOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*RestoreSubmissionInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *RestoreSubmissionInput, got %T", in))
		}

		return i.client.RestoreSubmission(ctx, message, opts...)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.judge.SubmissionService.RestoreSubmission", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*RestoreSubmissionOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *RestoreSubmissionOutput, got %T", out))
	}

	return message, err
}

func (i *SubmissionServiceInterceptor) RetestProblem(ctx context.Context, in *RetestProblemInput, opts ...grpc.CallOption) (*RetestProblemOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*RetestProblemInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *RetestProblemInput, got %T", in))
		}

		return i.client.RetestProblem(ctx, message, opts...)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.judge.SubmissionService.RetestProblem", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*RetestProblemOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *RetestProblemOutput, got %T", out))
	}

	return message, err
}