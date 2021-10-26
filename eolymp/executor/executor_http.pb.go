// Code generated by protoc-gen-go-server. DO NOT EDIT.
// See https://github.com/eolymp/contracts/tree/main/cmd/protoc-gen-go-server for more details.

package executor

import (
	context "context"
	oauth "github.com/eolymp/go-packages/oauth"
	mux "github.com/gorilla/mux"
	prometheus "github.com/prometheus/client_golang/prometheus"
	promauto "github.com/prometheus/client_golang/prometheus/promauto"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protojson "google.golang.org/protobuf/encoding/protojson"
	proto "google.golang.org/protobuf/proto"
	ioutil "io/ioutil"
	http "net/http"
	time "time"
)

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
	router.Handle("/eolymp.executor.Executor/DescribeLanguage", _Executor_DescribeLanguage(srv)).Methods(http.MethodPost)
	router.Handle("/eolymp.executor.Executor/ListLanguages", _Executor_ListLanguages(srv)).Methods(http.MethodPost)
	router.Handle("/eolymp.executor.Executor/DescribeRuntime", _Executor_DescribeRuntime(srv)).Methods(http.MethodPost)
	router.Handle("/eolymp.executor.Executor/ListRuntime", _Executor_ListRuntime(srv)).Methods(http.MethodPost)
	router.Handle("/eolymp.executor.Executor/DescribeCodeTemplate", _Executor_DescribeCodeTemplate(srv)).Methods(http.MethodPost)
	return router
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

var promExecutorRequestLatency = promauto.NewHistogramVec(prometheus.HistogramOpts{
	Name:    "executor_request_latency",
	Help:    "Executor request latency",
	Buckets: []float64{0.1, 0.4, 1, 5},
}, []string{"method", "status"})

type _ExecutorLimiter interface {
	Allow(context.Context, string, float64, int) bool
}

type ExecutorInterceptor struct {
	limiter _ExecutorLimiter
	server  ExecutorServer
}

// NewExecutorInterceptor constructs additional middleware for a server based on annotations in proto files
func NewExecutorInterceptor(srv ExecutorServer, lim _ExecutorLimiter) *ExecutorInterceptor {
	return &ExecutorInterceptor{server: srv, limiter: lim}
}

func (i *ExecutorInterceptor) DescribeLanguage(ctx context.Context, in *DescribeLanguageInput) (out *DescribeLanguageOutput, err error) {
	start := time.Now()
	defer func() {
		s, _ := status.FromError(err)
		if s == nil {
			s = status.New(codes.OK, "OK")
		}

		promExecutorRequestLatency.WithLabelValues("eolymp.executor.Executor/DescribeLanguage", s.Code().String()).
			Observe(time.Since(start).Seconds())
	}()

	token, ok := oauth.TokenFromContext(ctx)
	if ok && !token.Has("executor:runtime:read") {
		err = status.Error(codes.PermissionDenied, "required token scopes are missing: executor:runtime:read")
		return
	}

	if !i.limiter.Allow(ctx, "eolymp.executor.Executor/DescribeLanguage", 20, 200) {
		err = status.Error(codes.ResourceExhausted, "too many requests")
		return
	}

	out, err = i.server.DescribeLanguage(ctx, in)
	return
}

func (i *ExecutorInterceptor) ListLanguages(ctx context.Context, in *ListLanguagesInput) (out *ListLanguagesOutput, err error) {
	start := time.Now()
	defer func() {
		s, _ := status.FromError(err)
		if s == nil {
			s = status.New(codes.OK, "OK")
		}

		promExecutorRequestLatency.WithLabelValues("eolymp.executor.Executor/ListLanguages", s.Code().String()).
			Observe(time.Since(start).Seconds())
	}()

	token, ok := oauth.TokenFromContext(ctx)
	if ok && !token.Has("executor:runtime:read") {
		err = status.Error(codes.PermissionDenied, "required token scopes are missing: executor:runtime:read")
		return
	}

	if !i.limiter.Allow(ctx, "eolymp.executor.Executor/ListLanguages", 20, 200) {
		err = status.Error(codes.ResourceExhausted, "too many requests")
		return
	}

	out, err = i.server.ListLanguages(ctx, in)
	return
}

func (i *ExecutorInterceptor) DescribeRuntime(ctx context.Context, in *DescribeRuntimeInput) (out *DescribeRuntimeOutput, err error) {
	start := time.Now()
	defer func() {
		s, _ := status.FromError(err)
		if s == nil {
			s = status.New(codes.OK, "OK")
		}

		promExecutorRequestLatency.WithLabelValues("eolymp.executor.Executor/DescribeRuntime", s.Code().String()).
			Observe(time.Since(start).Seconds())
	}()

	token, ok := oauth.TokenFromContext(ctx)
	if ok && !token.Has("executor:runtime:read") {
		err = status.Error(codes.PermissionDenied, "required token scopes are missing: executor:runtime:read")
		return
	}

	if !i.limiter.Allow(ctx, "eolymp.executor.Executor/DescribeRuntime", 20, 200) {
		err = status.Error(codes.ResourceExhausted, "too many requests")
		return
	}

	out, err = i.server.DescribeRuntime(ctx, in)
	return
}

func (i *ExecutorInterceptor) ListRuntime(ctx context.Context, in *ListRuntimeInput) (out *ListRuntimeOutput, err error) {
	start := time.Now()
	defer func() {
		s, _ := status.FromError(err)
		if s == nil {
			s = status.New(codes.OK, "OK")
		}

		promExecutorRequestLatency.WithLabelValues("eolymp.executor.Executor/ListRuntime", s.Code().String()).
			Observe(time.Since(start).Seconds())
	}()

	token, ok := oauth.TokenFromContext(ctx)
	if ok && !token.Has("executor:runtime:read") {
		err = status.Error(codes.PermissionDenied, "required token scopes are missing: executor:runtime:read")
		return
	}

	if !i.limiter.Allow(ctx, "eolymp.executor.Executor/ListRuntime", 20, 200) {
		err = status.Error(codes.ResourceExhausted, "too many requests")
		return
	}

	out, err = i.server.ListRuntime(ctx, in)
	return
}

func (i *ExecutorInterceptor) DescribeCodeTemplate(ctx context.Context, in *DescribeCodeTemplateInput) (out *DescribeCodeTemplateOutput, err error) {
	start := time.Now()
	defer func() {
		s, _ := status.FromError(err)
		if s == nil {
			s = status.New(codes.OK, "OK")
		}

		promExecutorRequestLatency.WithLabelValues("eolymp.executor.Executor/DescribeCodeTemplate", s.Code().String()).
			Observe(time.Since(start).Seconds())
	}()

	token, ok := oauth.TokenFromContext(ctx)
	if ok && !token.Has("executor:runtime:read") {
		err = status.Error(codes.PermissionDenied, "required token scopes are missing: executor:runtime:read")
		return
	}

	if !i.limiter.Allow(ctx, "eolymp.executor.Executor/DescribeCodeTemplate", 20, 200) {
		err = status.Error(codes.ResourceExhausted, "too many requests")
		return
	}

	out, err = i.server.DescribeCodeTemplate(ctx, in)
	return
}
