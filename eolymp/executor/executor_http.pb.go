// Code generated by protoc-gen-go-server. DO NOT EDIT.
// See https://github.com/eolymp/contracts/tree/main/cmd/protoc-gen-go-server for more details.

package executor

import (
	context "context"
	fmt "fmt"
	mux "github.com/gorilla/mux"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protojson "google.golang.org/protobuf/encoding/protojson"
	proto "google.golang.org/protobuf/proto"
	ioutil "io/ioutil"
	http "net/http"
)

// _Executor_HTTPReadQueryString parses body into proto.Message
func _Executor_HTTPReadQueryString(r *http.Request, v proto.Message) error {
	query := r.URL.Query().Get("q")
	if query == "" {
		return nil
	}

	if err := (protojson.UnmarshalOptions{DiscardUnknown: true}.Unmarshal([]byte(query), v)); err != nil {
		return err
	}

	return nil
}

// _Executor_HTTPReadRequestBody parses body into proto.Message
func _Executor_HTTPReadRequestBody(r *http.Request, v proto.Message) error {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}

	if err := (protojson.UnmarshalOptions{DiscardUnknown: true}.Unmarshal(data, v)); err != nil {
		return err
	}

	return nil
}

// _Executor_HTTPWriteResponse writes proto.Message to HTTP response
func _Executor_HTTPWriteResponse(w http.ResponseWriter, v proto.Message) {
	data, err := protojson.Marshal(v)
	if err != nil {
		_Executor_HTTPWriteErrorResponse(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	_, _ = w.Write(data)
}

// _Executor_HTTPWriteErrorResponse writes error to HTTP response with error status code
func _Executor_HTTPWriteErrorResponse(w http.ResponseWriter, e error) {
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

// NewExecutorHandler constructs new http.Handler for ExecutorServer
func NewExecutorHandler(srv ExecutorServer) http.Handler {
	router := mux.NewRouter()
	router.Handle("/eolymp.executor.Executor/CreateTask", _Executor_CreateTask(srv)).Methods(http.MethodPost)
	router.Handle("/eolymp.executor.Executor/DescribeLanguage", _Executor_DescribeLanguage(srv)).Methods(http.MethodPost)
	router.Handle("/eolymp.executor.Executor/ListLanguages", _Executor_ListLanguages(srv)).Methods(http.MethodPost)
	router.Handle("/eolymp.executor.Executor/DescribeRuntime", _Executor_DescribeRuntime(srv)).Methods(http.MethodPost)
	router.Handle("/eolymp.executor.Executor/ListRuntime", _Executor_ListRuntime(srv)).Methods(http.MethodPost)
	router.Handle("/eolymp.executor.Executor/DescribeCodeTemplate", _Executor_DescribeCodeTemplate(srv)).Methods(http.MethodPost)
	return router
}

// NewExecutorHandlerHttp constructs new http.Handler for ExecutorServer
// This constructor creates http.Handler, the actual implementation might change at any moment
func NewExecutorHandlerHttp(srv ExecutorServer, prefix string) http.Handler {
	router := mux.NewRouter()

	router.Handle(prefix+"/exec/languages/{language_id}", _Executor_DescribeLanguage_Rule0(srv)).
		Methods("GET").
		Name("eolymp.executor.Executor.DescribeLanguage")

	router.Handle(prefix+"/exec/languages", _Executor_ListLanguages_Rule0(srv)).
		Methods("GET").
		Name("eolymp.executor.Executor.ListLanguages")

	router.Handle(prefix+"/exec/runtime/{runtime_id}", _Executor_DescribeRuntime_Rule0(srv)).
		Methods("GET").
		Name("eolymp.executor.Executor.DescribeRuntime")

	router.Handle(prefix+"/exec/runtime", _Executor_ListRuntime_Rule0(srv)).
		Methods("GET").
		Name("eolymp.executor.Executor.ListRuntime")

	router.Handle(prefix+"/exec/runtime/{runtime_id}/template", _Executor_DescribeCodeTemplate_Rule0(srv)).
		Methods("GET").
		Name("eolymp.executor.Executor.DescribeCodeTemplate")

	return router
}

func _Executor_CreateTask(srv ExecutorServer) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &CreateTaskInput{}

		if err := _Executor_HTTPReadRequestBody(r, in); err != nil {
			err = status.New(codes.InvalidArgument, err.Error()).Err()
			_Executor_HTTPWriteErrorResponse(w, err)
			return
		}

		out, err := srv.CreateTask(r.Context(), in)
		if err != nil {
			_Executor_HTTPWriteErrorResponse(w, err)
			return
		}

		_Executor_HTTPWriteResponse(w, out)
	})
}

func _Executor_DescribeLanguage(srv ExecutorServer) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &DescribeLanguageInput{}

		if err := _Executor_HTTPReadRequestBody(r, in); err != nil {
			err = status.New(codes.InvalidArgument, err.Error()).Err()
			_Executor_HTTPWriteErrorResponse(w, err)
			return
		}

		out, err := srv.DescribeLanguage(r.Context(), in)
		if err != nil {
			_Executor_HTTPWriteErrorResponse(w, err)
			return
		}

		_Executor_HTTPWriteResponse(w, out)
	})
}

func _Executor_ListLanguages(srv ExecutorServer) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &ListLanguagesInput{}

		if err := _Executor_HTTPReadRequestBody(r, in); err != nil {
			err = status.New(codes.InvalidArgument, err.Error()).Err()
			_Executor_HTTPWriteErrorResponse(w, err)
			return
		}

		out, err := srv.ListLanguages(r.Context(), in)
		if err != nil {
			_Executor_HTTPWriteErrorResponse(w, err)
			return
		}

		_Executor_HTTPWriteResponse(w, out)
	})
}

func _Executor_DescribeRuntime(srv ExecutorServer) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &DescribeRuntimeInput{}

		if err := _Executor_HTTPReadRequestBody(r, in); err != nil {
			err = status.New(codes.InvalidArgument, err.Error()).Err()
			_Executor_HTTPWriteErrorResponse(w, err)
			return
		}

		out, err := srv.DescribeRuntime(r.Context(), in)
		if err != nil {
			_Executor_HTTPWriteErrorResponse(w, err)
			return
		}

		_Executor_HTTPWriteResponse(w, out)
	})
}

func _Executor_ListRuntime(srv ExecutorServer) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &ListRuntimeInput{}

		if err := _Executor_HTTPReadRequestBody(r, in); err != nil {
			err = status.New(codes.InvalidArgument, err.Error()).Err()
			_Executor_HTTPWriteErrorResponse(w, err)
			return
		}

		out, err := srv.ListRuntime(r.Context(), in)
		if err != nil {
			_Executor_HTTPWriteErrorResponse(w, err)
			return
		}

		_Executor_HTTPWriteResponse(w, out)
	})
}

func _Executor_DescribeCodeTemplate(srv ExecutorServer) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &DescribeCodeTemplateInput{}

		if err := _Executor_HTTPReadRequestBody(r, in); err != nil {
			err = status.New(codes.InvalidArgument, err.Error()).Err()
			_Executor_HTTPWriteErrorResponse(w, err)
			return
		}

		out, err := srv.DescribeCodeTemplate(r.Context(), in)
		if err != nil {
			_Executor_HTTPWriteErrorResponse(w, err)
			return
		}

		_Executor_HTTPWriteResponse(w, out)
	})
}

func _Executor_DescribeLanguage_Rule0(srv ExecutorServer) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &DescribeLanguageInput{}

		if err := _Executor_HTTPReadQueryString(r, in); err != nil {
			err = status.New(codes.InvalidArgument, err.Error()).Err()
			_Executor_HTTPWriteErrorResponse(w, err)
			return
		}

		vars := mux.Vars(r)
		in.LanguageId = vars["language_id"]

		out, err := srv.DescribeLanguage(r.Context(), in)
		if err != nil {
			_Executor_HTTPWriteErrorResponse(w, err)
			return
		}

		_Executor_HTTPWriteResponse(w, out)
	})
}

func _Executor_ListLanguages_Rule0(srv ExecutorServer) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &ListLanguagesInput{}

		if err := _Executor_HTTPReadQueryString(r, in); err != nil {
			err = status.New(codes.InvalidArgument, err.Error()).Err()
			_Executor_HTTPWriteErrorResponse(w, err)
			return
		}

		out, err := srv.ListLanguages(r.Context(), in)
		if err != nil {
			_Executor_HTTPWriteErrorResponse(w, err)
			return
		}

		_Executor_HTTPWriteResponse(w, out)
	})
}

func _Executor_DescribeRuntime_Rule0(srv ExecutorServer) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &DescribeRuntimeInput{}

		if err := _Executor_HTTPReadQueryString(r, in); err != nil {
			err = status.New(codes.InvalidArgument, err.Error()).Err()
			_Executor_HTTPWriteErrorResponse(w, err)
			return
		}

		vars := mux.Vars(r)
		in.RuntimeId = vars["runtime_id"]

		out, err := srv.DescribeRuntime(r.Context(), in)
		if err != nil {
			_Executor_HTTPWriteErrorResponse(w, err)
			return
		}

		_Executor_HTTPWriteResponse(w, out)
	})
}

func _Executor_ListRuntime_Rule0(srv ExecutorServer) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &ListRuntimeInput{}

		if err := _Executor_HTTPReadQueryString(r, in); err != nil {
			err = status.New(codes.InvalidArgument, err.Error()).Err()
			_Executor_HTTPWriteErrorResponse(w, err)
			return
		}

		out, err := srv.ListRuntime(r.Context(), in)
		if err != nil {
			_Executor_HTTPWriteErrorResponse(w, err)
			return
		}

		_Executor_HTTPWriteResponse(w, out)
	})
}

func _Executor_DescribeCodeTemplate_Rule0(srv ExecutorServer) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &DescribeCodeTemplateInput{}

		if err := _Executor_HTTPReadQueryString(r, in); err != nil {
			err = status.New(codes.InvalidArgument, err.Error()).Err()
			_Executor_HTTPWriteErrorResponse(w, err)
			return
		}

		vars := mux.Vars(r)
		in.RuntimeId = vars["runtime_id"]

		out, err := srv.DescribeCodeTemplate(r.Context(), in)
		if err != nil {
			_Executor_HTTPWriteErrorResponse(w, err)
			return
		}

		_Executor_HTTPWriteResponse(w, out)
	})
}

type _ExecutorHandler = func(ctx context.Context, in proto.Message) (proto.Message, error)
type _ExecutorMiddleware = func(ctx context.Context, method string, in proto.Message, handler _ExecutorHandler) (out proto.Message, err error)
type ExecutorInterceptor struct {
	middleware []_ExecutorMiddleware
	server     ExecutorServer
}

// NewExecutorInterceptor constructs additional middleware for a server based on annotations in proto files
func NewExecutorInterceptor(srv ExecutorServer, middleware ..._ExecutorMiddleware) *ExecutorInterceptor {
	return &ExecutorInterceptor{server: srv, middleware: middleware}
}

func (i *ExecutorInterceptor) CreateTask(ctx context.Context, in *CreateTaskInput) (*CreateTaskOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*CreateTaskInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *CreateTaskInput, got %T", in))
		}

		return i.server.CreateTask(ctx, message)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.executor.Executor.CreateTask", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*CreateTaskOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *CreateTaskOutput, got %T", out))
	}

	return message, err
}

func (i *ExecutorInterceptor) DescribeLanguage(ctx context.Context, in *DescribeLanguageInput) (*DescribeLanguageOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*DescribeLanguageInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *DescribeLanguageInput, got %T", in))
		}

		return i.server.DescribeLanguage(ctx, message)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.executor.Executor.DescribeLanguage", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*DescribeLanguageOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *DescribeLanguageOutput, got %T", out))
	}

	return message, err
}

func (i *ExecutorInterceptor) ListLanguages(ctx context.Context, in *ListLanguagesInput) (*ListLanguagesOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*ListLanguagesInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *ListLanguagesInput, got %T", in))
		}

		return i.server.ListLanguages(ctx, message)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.executor.Executor.ListLanguages", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*ListLanguagesOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *ListLanguagesOutput, got %T", out))
	}

	return message, err
}

func (i *ExecutorInterceptor) DescribeRuntime(ctx context.Context, in *DescribeRuntimeInput) (*DescribeRuntimeOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*DescribeRuntimeInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *DescribeRuntimeInput, got %T", in))
		}

		return i.server.DescribeRuntime(ctx, message)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.executor.Executor.DescribeRuntime", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*DescribeRuntimeOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *DescribeRuntimeOutput, got %T", out))
	}

	return message, err
}

func (i *ExecutorInterceptor) ListRuntime(ctx context.Context, in *ListRuntimeInput) (*ListRuntimeOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*ListRuntimeInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *ListRuntimeInput, got %T", in))
		}

		return i.server.ListRuntime(ctx, message)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.executor.Executor.ListRuntime", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*ListRuntimeOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *ListRuntimeOutput, got %T", out))
	}

	return message, err
}

func (i *ExecutorInterceptor) DescribeCodeTemplate(ctx context.Context, in *DescribeCodeTemplateInput) (*DescribeCodeTemplateOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*DescribeCodeTemplateInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *DescribeCodeTemplateInput, got %T", in))
		}

		return i.server.DescribeCodeTemplate(ctx, message)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.executor.Executor.DescribeCodeTemplate", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*DescribeCodeTemplateOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *DescribeCodeTemplateOutput, got %T", out))
	}

	return message, err
}
