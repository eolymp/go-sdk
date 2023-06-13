// Code generated by protoc-gen-go-server. DO NOT EDIT.
// See https://github.com/eolymp/contracts/tree/main/cmd/protoc-gen-go-server for more details.

package worker

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

// _Worker_HTTPReadQueryString parses body into proto.Message
func _Worker_HTTPReadQueryString(r *http.Request, v proto.Message) error {
	query := r.URL.Query().Get("q")
	if query == "" {
		return nil
	}

	if err := (protojson.UnmarshalOptions{DiscardUnknown: true}.Unmarshal([]byte(query), v)); err != nil {
		return err
	}

	return nil
}

// _Worker_HTTPReadRequestBody parses body into proto.Message
func _Worker_HTTPReadRequestBody(r *http.Request, v proto.Message) error {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}

	if err := (protojson.UnmarshalOptions{DiscardUnknown: true}.Unmarshal(data, v)); err != nil {
		return err
	}

	return nil
}

// _Worker_HTTPWriteResponse writes proto.Message to HTTP response
func _Worker_HTTPWriteResponse(w http.ResponseWriter, v proto.Message) {
	data, err := protojson.Marshal(v)
	if err != nil {
		_Worker_HTTPWriteErrorResponse(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	_, _ = w.Write(data)
}

// _Worker_HTTPWriteErrorResponse writes error to HTTP response with error status code
func _Worker_HTTPWriteErrorResponse(w http.ResponseWriter, e error) {
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

// _Worker_WebsocketErrorResponse writes error to websocket connection
func _Worker_WebsocketErrorResponse(conn *websocket.Conn, e error) {
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

// _Worker_WebsocketCodec implements protobuf codec for websockets package
var _Worker_WebsocketCodec = websocket.Codec{
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

// RegisterWorkerHttpHandlers adds handlers for for WorkerClient
// This constructor creates http.Handler, the actual implementation might change at any moment
func RegisterWorkerHttpHandlers(router *mux.Router, prefix string, cli WorkerClient) {
	router.Handle(prefix+"/jobs", _Worker_CreateJob_Rule0(cli)).
		Methods("POST").
		Name("eolymp.worker.Worker.CreateJob")
	router.Handle(prefix+"/jobs/{job_id}", _Worker_DescribeJob_Rule0(cli)).
		Methods("GET").
		Name("eolymp.worker.Worker.DescribeJob")
	router.Handle(prefix+"/jobs", _Worker_ListJobs_Rule0(cli)).
		Methods("GET").
		Name("eolymp.worker.Worker.ListJobs")
}

func _Worker_CreateJob_Rule0(cli WorkerClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &CreateJobInput{}

		if err := _Worker_HTTPReadRequestBody(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_Worker_HTTPWriteErrorResponse(w, err)
			return
		}

		out, err := cli.CreateJob(r.Context(), in)
		if err != nil {
			_Worker_HTTPWriteErrorResponse(w, err)
			return
		}

		_Worker_HTTPWriteResponse(w, out)
	})
}

func _Worker_DescribeJob_Rule0(cli WorkerClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &DescribeJobInput{}

		if err := _Worker_HTTPReadQueryString(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_Worker_HTTPWriteErrorResponse(w, err)
			return
		}

		vars := mux.Vars(r)
		in.JobId = vars["job_id"]

		out, err := cli.DescribeJob(r.Context(), in)
		if err != nil {
			_Worker_HTTPWriteErrorResponse(w, err)
			return
		}

		_Worker_HTTPWriteResponse(w, out)
	})
}

func _Worker_ListJobs_Rule0(cli WorkerClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &ListJobsInput{}

		if err := _Worker_HTTPReadQueryString(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_Worker_HTTPWriteErrorResponse(w, err)
			return
		}

		out, err := cli.ListJobs(r.Context(), in)
		if err != nil {
			_Worker_HTTPWriteErrorResponse(w, err)
			return
		}

		_Worker_HTTPWriteResponse(w, out)
	})
}

type _WorkerHandler = func(ctx context.Context, in proto.Message) (proto.Message, error)
type _WorkerMiddleware = func(ctx context.Context, method string, in proto.Message, handler _WorkerHandler) (out proto.Message, err error)
type WorkerInterceptor struct {
	middleware []_WorkerMiddleware
	client     WorkerClient
}

// NewWorkerInterceptor constructs additional middleware for a server based on annotations in proto files
func NewWorkerInterceptor(cli WorkerClient, middleware ..._WorkerMiddleware) *WorkerInterceptor {
	return &WorkerInterceptor{client: cli, middleware: middleware}
}

func (i *WorkerInterceptor) CreateJob(ctx context.Context, in *CreateJobInput, opts ...grpc.CallOption) (*CreateJobOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*CreateJobInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *CreateJobInput, got %T", in))
		}

		return i.client.CreateJob(ctx, message, opts...)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.worker.Worker.CreateJob", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*CreateJobOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *CreateJobOutput, got %T", out))
	}

	return message, err
}

func (i *WorkerInterceptor) DescribeJob(ctx context.Context, in *DescribeJobInput, opts ...grpc.CallOption) (*DescribeJobOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*DescribeJobInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *DescribeJobInput, got %T", in))
		}

		return i.client.DescribeJob(ctx, message, opts...)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.worker.Worker.DescribeJob", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*DescribeJobOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *DescribeJobOutput, got %T", out))
	}

	return message, err
}

func (i *WorkerInterceptor) ListJobs(ctx context.Context, in *ListJobsInput, opts ...grpc.CallOption) (*ListJobsOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*ListJobsInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *ListJobsInput, got %T", in))
		}

		return i.client.ListJobs(ctx, message, opts...)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.worker.Worker.ListJobs", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*ListJobsOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *ListJobsOutput, got %T", out))
	}

	return message, err
}

// _WorkerService_HTTPReadQueryString parses body into proto.Message
func _WorkerService_HTTPReadQueryString(r *http.Request, v proto.Message) error {
	query := r.URL.Query().Get("q")
	if query == "" {
		return nil
	}

	if err := (protojson.UnmarshalOptions{DiscardUnknown: true}.Unmarshal([]byte(query), v)); err != nil {
		return err
	}

	return nil
}

// _WorkerService_HTTPReadRequestBody parses body into proto.Message
func _WorkerService_HTTPReadRequestBody(r *http.Request, v proto.Message) error {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}

	if err := (protojson.UnmarshalOptions{DiscardUnknown: true}.Unmarshal(data, v)); err != nil {
		return err
	}

	return nil
}

// _WorkerService_HTTPWriteResponse writes proto.Message to HTTP response
func _WorkerService_HTTPWriteResponse(w http.ResponseWriter, v proto.Message) {
	data, err := protojson.Marshal(v)
	if err != nil {
		_WorkerService_HTTPWriteErrorResponse(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	_, _ = w.Write(data)
}

// _WorkerService_HTTPWriteErrorResponse writes error to HTTP response with error status code
func _WorkerService_HTTPWriteErrorResponse(w http.ResponseWriter, e error) {
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

// _WorkerService_WebsocketErrorResponse writes error to websocket connection
func _WorkerService_WebsocketErrorResponse(conn *websocket.Conn, e error) {
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

// _WorkerService_WebsocketCodec implements protobuf codec for websockets package
var _WorkerService_WebsocketCodec = websocket.Codec{
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

// RegisterWorkerServiceHttpHandlers adds handlers for for WorkerServiceClient
// This constructor creates http.Handler, the actual implementation might change at any moment
func RegisterWorkerServiceHttpHandlers(router *mux.Router, prefix string, cli WorkerServiceClient) {
	router.Handle(prefix+"/jobs", _WorkerService_CreateJob_Rule0(cli)).
		Methods("POST").
		Name("eolymp.worker.WorkerService.CreateJob")
	router.Handle(prefix+"/jobs/{job_id}", _WorkerService_DescribeJob_Rule0(cli)).
		Methods("GET").
		Name("eolymp.worker.WorkerService.DescribeJob")
	router.Handle(prefix+"/jobs", _WorkerService_ListJobs_Rule0(cli)).
		Methods("GET").
		Name("eolymp.worker.WorkerService.ListJobs")
}

func _WorkerService_CreateJob_Rule0(cli WorkerServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &CreateJobInput{}

		if err := _WorkerService_HTTPReadRequestBody(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_WorkerService_HTTPWriteErrorResponse(w, err)
			return
		}

		out, err := cli.CreateJob(r.Context(), in)
		if err != nil {
			_WorkerService_HTTPWriteErrorResponse(w, err)
			return
		}

		_WorkerService_HTTPWriteResponse(w, out)
	})
}

func _WorkerService_DescribeJob_Rule0(cli WorkerServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &DescribeJobInput{}

		if err := _WorkerService_HTTPReadQueryString(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_WorkerService_HTTPWriteErrorResponse(w, err)
			return
		}

		vars := mux.Vars(r)
		in.JobId = vars["job_id"]

		out, err := cli.DescribeJob(r.Context(), in)
		if err != nil {
			_WorkerService_HTTPWriteErrorResponse(w, err)
			return
		}

		_WorkerService_HTTPWriteResponse(w, out)
	})
}

func _WorkerService_ListJobs_Rule0(cli WorkerServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &ListJobsInput{}

		if err := _WorkerService_HTTPReadQueryString(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_WorkerService_HTTPWriteErrorResponse(w, err)
			return
		}

		out, err := cli.ListJobs(r.Context(), in)
		if err != nil {
			_WorkerService_HTTPWriteErrorResponse(w, err)
			return
		}

		_WorkerService_HTTPWriteResponse(w, out)
	})
}

type _WorkerServiceHandler = func(ctx context.Context, in proto.Message) (proto.Message, error)
type _WorkerServiceMiddleware = func(ctx context.Context, method string, in proto.Message, handler _WorkerServiceHandler) (out proto.Message, err error)
type WorkerServiceInterceptor struct {
	middleware []_WorkerServiceMiddleware
	client     WorkerServiceClient
}

// NewWorkerServiceInterceptor constructs additional middleware for a server based on annotations in proto files
func NewWorkerServiceInterceptor(cli WorkerServiceClient, middleware ..._WorkerServiceMiddleware) *WorkerServiceInterceptor {
	return &WorkerServiceInterceptor{client: cli, middleware: middleware}
}

func (i *WorkerServiceInterceptor) CreateJob(ctx context.Context, in *CreateJobInput, opts ...grpc.CallOption) (*CreateJobOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*CreateJobInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *CreateJobInput, got %T", in))
		}

		return i.client.CreateJob(ctx, message, opts...)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.worker.WorkerService.CreateJob", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*CreateJobOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *CreateJobOutput, got %T", out))
	}

	return message, err
}

func (i *WorkerServiceInterceptor) DescribeJob(ctx context.Context, in *DescribeJobInput, opts ...grpc.CallOption) (*DescribeJobOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*DescribeJobInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *DescribeJobInput, got %T", in))
		}

		return i.client.DescribeJob(ctx, message, opts...)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.worker.WorkerService.DescribeJob", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*DescribeJobOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *DescribeJobOutput, got %T", out))
	}

	return message, err
}

func (i *WorkerServiceInterceptor) ListJobs(ctx context.Context, in *ListJobsInput, opts ...grpc.CallOption) (*ListJobsOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*ListJobsInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *ListJobsInput, got %T", in))
		}

		return i.client.ListJobs(ctx, message, opts...)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.worker.WorkerService.ListJobs", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*ListJobsOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *ListJobsOutput, got %T", out))
	}

	return message, err
}
