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

// _ProblemSolver_HTTPReadQueryString parses body into proto.Message
func _ProblemSolver_HTTPReadQueryString(r *http.Request, v proto.Message) error {
	query := r.URL.Query().Get("q")
	if query == "" {
		return nil
	}

	if err := (protojson.UnmarshalOptions{DiscardUnknown: true}.Unmarshal([]byte(query), v)); err != nil {
		return err
	}

	return nil
}

// _ProblemSolver_HTTPReadRequestBody parses body into proto.Message
func _ProblemSolver_HTTPReadRequestBody(r *http.Request, v proto.Message) error {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}

	if err := (protojson.UnmarshalOptions{DiscardUnknown: true}.Unmarshal(data, v)); err != nil {
		return err
	}

	return nil
}

// _ProblemSolver_HTTPWriteResponse writes proto.Message to HTTP response
func _ProblemSolver_HTTPWriteResponse(w http.ResponseWriter, v proto.Message) {
	data, err := protojson.Marshal(v)
	if err != nil {
		_ProblemSolver_HTTPWriteErrorResponse(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	_, _ = w.Write(data)
}

// _ProblemSolver_HTTPWriteErrorResponse writes error to HTTP response with error status code
func _ProblemSolver_HTTPWriteErrorResponse(w http.ResponseWriter, e error) {
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

// _ProblemSolver_WebsocketErrorResponse writes error to websocket connection
func _ProblemSolver_WebsocketErrorResponse(conn *websocket.Conn, e error) {
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

// _ProblemSolver_WebsocketCodec implements protobuf codec for websockets package
var _ProblemSolver_WebsocketCodec = websocket.Codec{
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

type _ProblemSolver_WatchSubmission_WSStream struct {
	ctx  context.Context
	conn *websocket.Conn
}

func (s *_ProblemSolver_WatchSubmission_WSStream) Send(m *WatchSubmissionOutput) error {
	return s.SendMsg(m)
}

func (s *_ProblemSolver_WatchSubmission_WSStream) SetHeader(metadata.MD) error {
	return nil
}

func (s *_ProblemSolver_WatchSubmission_WSStream) SendHeader(metadata.MD) error {
	return nil
}

func (s *_ProblemSolver_WatchSubmission_WSStream) SetTrailer(metadata.MD) {
}

func (s *_ProblemSolver_WatchSubmission_WSStream) Context() context.Context {
	return s.ctx
}

func (s *_ProblemSolver_WatchSubmission_WSStream) SendMsg(m interface{}) error {
	return _ProblemSolver_WebsocketCodec.Send(s.conn, m)
}

func (s *_ProblemSolver_WatchSubmission_WSStream) RecvMsg(m interface{}) error {
	return nil
}

// RegisterProblemSolverHttpHandlers adds handlers for for ProblemSolverClient
// This constructor creates http.Handler, the actual implementation might change at any moment
func RegisterProblemSolverHttpHandlers(router *mux.Router, prefix string, cli ProblemSolverClient) {
	router.Handle(prefix+"/statements", _ProblemSolver_ListStatements_Rule0(cli)).
		Methods("GET").
		Name("eolymp.atlas.ProblemSolver.ListStatements")
	router.Handle(prefix+"/examples", _ProblemSolver_ListExamples_Rule0(cli)).
		Methods("GET").
		Name("eolymp.atlas.ProblemSolver.ListExamples")
	router.Handle(prefix+"/submissions", _ProblemSolver_CreateSubmission_Rule0(cli)).
		Methods("POST").
		Name("eolymp.atlas.ProblemSolver.CreateSubmission")
	router.Handle(prefix+"/submissions", _ProblemSolver_ListSubmissions_Rule0(cli)).
		Methods("GET").
		Name("eolymp.atlas.ProblemSolver.ListSubmissions")
	router.Handle(prefix+"/submissions/{submission_id}", _ProblemSolver_DescribeSubmission_Rule0(cli)).
		Methods("GET").
		Name("eolymp.atlas.ProblemSolver.DescribeSubmission")
	router.Handle(prefix+"/score", _ProblemSolver_DescribeScore_Rule0(cli)).
		Methods("GET").
		Name("eolymp.atlas.ProblemSolver.DescribeScore")
	router.Handle(prefix+"/template", _ProblemSolver_LookupCodeTemplate_Rule0(cli)).
		Methods("GET").
		Name("eolymp.atlas.ProblemSolver.LookupCodeTemplate")
}

func _ProblemSolver_ListStatements_Rule0(cli ProblemSolverClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &ListStatementsInput{}

		if err := _ProblemSolver_HTTPReadQueryString(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_ProblemSolver_HTTPWriteErrorResponse(w, err)
			return
		}

		out, err := cli.ListStatements(r.Context(), in)
		if err != nil {
			_ProblemSolver_HTTPWriteErrorResponse(w, err)
			return
		}

		_ProblemSolver_HTTPWriteResponse(w, out)
	})
}

func _ProblemSolver_ListExamples_Rule0(cli ProblemSolverClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &ListExamplesInput{}

		if err := _ProblemSolver_HTTPReadQueryString(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_ProblemSolver_HTTPWriteErrorResponse(w, err)
			return
		}

		out, err := cli.ListExamples(r.Context(), in)
		if err != nil {
			_ProblemSolver_HTTPWriteErrorResponse(w, err)
			return
		}

		_ProblemSolver_HTTPWriteResponse(w, out)
	})
}

func _ProblemSolver_CreateSubmission_Rule0(cli ProblemSolverClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &CreateSubmissionInput{}

		if err := _ProblemSolver_HTTPReadRequestBody(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_ProblemSolver_HTTPWriteErrorResponse(w, err)
			return
		}

		out, err := cli.CreateSubmission(r.Context(), in)
		if err != nil {
			_ProblemSolver_HTTPWriteErrorResponse(w, err)
			return
		}

		_ProblemSolver_HTTPWriteResponse(w, out)
	})
}

func _ProblemSolver_ListSubmissions_Rule0(cli ProblemSolverClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &ListSubmissionsInput{}

		if err := _ProblemSolver_HTTPReadQueryString(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_ProblemSolver_HTTPWriteErrorResponse(w, err)
			return
		}

		out, err := cli.ListSubmissions(r.Context(), in)
		if err != nil {
			_ProblemSolver_HTTPWriteErrorResponse(w, err)
			return
		}

		_ProblemSolver_HTTPWriteResponse(w, out)
	})
}

func _ProblemSolver_DescribeSubmission_Rule0(cli ProblemSolverClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &DescribeSubmissionInput{}

		if err := _ProblemSolver_HTTPReadQueryString(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_ProblemSolver_HTTPWriteErrorResponse(w, err)
			return
		}

		vars := mux.Vars(r)
		in.SubmissionId = vars["submission_id"]

		out, err := cli.DescribeSubmission(r.Context(), in)
		if err != nil {
			_ProblemSolver_HTTPWriteErrorResponse(w, err)
			return
		}

		_ProblemSolver_HTTPWriteResponse(w, out)
	})
}

func _ProblemSolver_DescribeScore_Rule0(cli ProblemSolverClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &DescribeScoreInput{}

		if err := _ProblemSolver_HTTPReadQueryString(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_ProblemSolver_HTTPWriteErrorResponse(w, err)
			return
		}

		out, err := cli.DescribeScore(r.Context(), in)
		if err != nil {
			_ProblemSolver_HTTPWriteErrorResponse(w, err)
			return
		}

		_ProblemSolver_HTTPWriteResponse(w, out)
	})
}

func _ProblemSolver_LookupCodeTemplate_Rule0(cli ProblemSolverClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &LookupCodeTemplateInput{}

		if err := _ProblemSolver_HTTPReadQueryString(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_ProblemSolver_HTTPWriteErrorResponse(w, err)
			return
		}

		out, err := cli.LookupCodeTemplate(r.Context(), in)
		if err != nil {
			_ProblemSolver_HTTPWriteErrorResponse(w, err)
			return
		}

		_ProblemSolver_HTTPWriteResponse(w, out)
	})
}

type _ProblemSolverHandler = func(ctx context.Context, in proto.Message) (proto.Message, error)
type _ProblemSolverMiddleware = func(ctx context.Context, method string, in proto.Message, handler _ProblemSolverHandler) (out proto.Message, err error)
type ProblemSolverInterceptor struct {
	middleware []_ProblemSolverMiddleware
	client     ProblemSolverClient
}

// NewProblemSolverInterceptor constructs additional middleware for a server based on annotations in proto files
func NewProblemSolverInterceptor(cli ProblemSolverClient, middleware ..._ProblemSolverMiddleware) *ProblemSolverInterceptor {
	return &ProblemSolverInterceptor{client: cli, middleware: middleware}
}

func (i *ProblemSolverInterceptor) ListStatements(ctx context.Context, in *ListStatementsInput, opts ...grpc.CallOption) (*ListStatementsOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*ListStatementsInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *ListStatementsInput, got %T", in))
		}

		return i.client.ListStatements(ctx, message, opts...)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.atlas.ProblemSolver.ListStatements", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*ListStatementsOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *ListStatementsOutput, got %T", out))
	}

	return message, err
}

func (i *ProblemSolverInterceptor) ListExamples(ctx context.Context, in *ListExamplesInput, opts ...grpc.CallOption) (*ListExamplesOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*ListExamplesInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *ListExamplesInput, got %T", in))
		}

		return i.client.ListExamples(ctx, message, opts...)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.atlas.ProblemSolver.ListExamples", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*ListExamplesOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *ListExamplesOutput, got %T", out))
	}

	return message, err
}

func (i *ProblemSolverInterceptor) CreateSubmission(ctx context.Context, in *CreateSubmissionInput, opts ...grpc.CallOption) (*CreateSubmissionOutput, error) {
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
			return mw(ctx, "eolymp.atlas.ProblemSolver.CreateSubmission", in, next)
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

func (i *ProblemSolverInterceptor) ListSubmissions(ctx context.Context, in *ListSubmissionsInput, opts ...grpc.CallOption) (*ListSubmissionsOutput, error) {
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
			return mw(ctx, "eolymp.atlas.ProblemSolver.ListSubmissions", in, next)
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

func (i *ProblemSolverInterceptor) DescribeSubmission(ctx context.Context, in *DescribeSubmissionInput, opts ...grpc.CallOption) (*DescribeSubmissionOutput, error) {
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
			return mw(ctx, "eolymp.atlas.ProblemSolver.DescribeSubmission", in, next)
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

func (i *ProblemSolverInterceptor) WatchSubmission(ctx context.Context, in *WatchSubmissionInput, opts ...grpc.CallOption) (ProblemSolver_WatchSubmissionClient, error) {
	return i.client.WatchSubmission(ctx, in, opts...)
}

func (i *ProblemSolverInterceptor) DescribeScore(ctx context.Context, in *DescribeScoreInput, opts ...grpc.CallOption) (*DescribeScoreOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*DescribeScoreInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *DescribeScoreInput, got %T", in))
		}

		return i.client.DescribeScore(ctx, message, opts...)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.atlas.ProblemSolver.DescribeScore", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*DescribeScoreOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *DescribeScoreOutput, got %T", out))
	}

	return message, err
}

func (i *ProblemSolverInterceptor) LookupCodeTemplate(ctx context.Context, in *LookupCodeTemplateInput, opts ...grpc.CallOption) (*LookupCodeTemplateOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*LookupCodeTemplateInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *LookupCodeTemplateInput, got %T", in))
		}

		return i.client.LookupCodeTemplate(ctx, message, opts...)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.atlas.ProblemSolver.LookupCodeTemplate", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*LookupCodeTemplateOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *LookupCodeTemplateOutput, got %T", out))
	}

	return message, err
}