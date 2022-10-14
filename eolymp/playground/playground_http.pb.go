// Code generated by protoc-gen-go-server. DO NOT EDIT.
// See https://github.com/eolymp/contracts/tree/main/cmd/protoc-gen-go-server for more details.

package playground

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

// NewPlaygroundHandler constructs new http.Handler for PlaygroundServer
func NewPlaygroundHandler(srv PlaygroundServer) http.Handler {
	router := mux.NewRouter()
	router.Handle("/eolymp.playground.Playground/CreateRun", _Playground_CreateRun(srv)).Methods(http.MethodPost)
	router.Handle("/eolymp.playground.Playground/DescribeRun", _Playground_DescribeRun(srv)).Methods(http.MethodPost)
	return router
}

// NewPlaygroundHandlerHttp constructs new http.Handler for PlaygroundServer
// This constructor creates http.Handler, the actual implementation might change at any moment
func NewPlaygroundHandlerHttp(srv PlaygroundServer, prefix string) http.Handler {
	router := mux.NewRouter()

	router.Handle(prefix+"/playground/runs", _Playground_CreateRun_Rule0(srv)).
		Methods("POST").
		Name("eolymp.playground.Playground.CreateRun")

	router.Handle(prefix+"/playground/runs/{run_id}", _Playground_DescribeRun_Rule0(srv)).
		Methods("GET").
		Name("eolymp.playground.Playground.DescribeRun")

	return router
}

func _Playground_CreateRun(srv PlaygroundServer) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &CreateRunInput{}

		if err := _Playground_HTTPReadRequestBody(r, in); err != nil {
			err = status.New(codes.InvalidArgument, err.Error()).Err()
			_Playground_HTTPWriteErrorResponse(w, err)
			return
		}

		out, err := srv.CreateRun(r.Context(), in)
		if err != nil {
			_Playground_HTTPWriteErrorResponse(w, err)
			return
		}

		_Playground_HTTPWriteResponse(w, out)
	})
}

func _Playground_DescribeRun(srv PlaygroundServer) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &DescribeRunInput{}

		if err := _Playground_HTTPReadRequestBody(r, in); err != nil {
			err = status.New(codes.InvalidArgument, err.Error()).Err()
			_Playground_HTTPWriteErrorResponse(w, err)
			return
		}

		out, err := srv.DescribeRun(r.Context(), in)
		if err != nil {
			_Playground_HTTPWriteErrorResponse(w, err)
			return
		}

		_Playground_HTTPWriteResponse(w, out)
	})
}

func _Playground_CreateRun_Rule0(srv PlaygroundServer) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &CreateRunInput{}

		if err := _Playground_HTTPReadRequestBody(r, in); err != nil {
			err = status.New(codes.InvalidArgument, err.Error()).Err()
			_Playground_HTTPWriteErrorResponse(w, err)
			return
		}

		out, err := srv.CreateRun(r.Context(), in)
		if err != nil {
			_Playground_HTTPWriteErrorResponse(w, err)
			return
		}

		_Playground_HTTPWriteResponse(w, out)
	})
}

func _Playground_DescribeRun_Rule0(srv PlaygroundServer) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &DescribeRunInput{}

		if err := _Playground_HTTPReadQueryString(r, in); err != nil {
			err = status.New(codes.InvalidArgument, err.Error()).Err()
			_Playground_HTTPWriteErrorResponse(w, err)
			return
		}

		vars := mux.Vars(r)
		in.RunId = vars["run_id"]

		out, err := srv.DescribeRun(r.Context(), in)
		if err != nil {
			_Playground_HTTPWriteErrorResponse(w, err)
			return
		}

		_Playground_HTTPWriteResponse(w, out)
	})
}

var promPlaygroundRequestLatency = promauto.NewHistogramVec(prometheus.HistogramOpts{
	Name:    "playground_request_latency",
	Help:    "Playground request latency",
	Buckets: []float64{0.1, 0.4, 1, 5},
}, []string{"method", "status"})

type _PlaygroundLimiter interface {
	Allow(context.Context, string, float64, int) bool
}

type PlaygroundInterceptor struct {
	limiter _PlaygroundLimiter
	server  PlaygroundServer
}

// NewPlaygroundInterceptor constructs additional middleware for a server based on annotations in proto files
func NewPlaygroundInterceptor(srv PlaygroundServer, lim _PlaygroundLimiter) *PlaygroundInterceptor {
	return &PlaygroundInterceptor{server: srv, limiter: lim}
}

func (i *PlaygroundInterceptor) CreateRun(ctx context.Context, in *CreateRunInput) (out *CreateRunOutput, err error) {
	start := time.Now()
	defer func() {
		s, _ := status.FromError(err)
		if s == nil {
			s = status.New(codes.OK, "OK")
		}

		promPlaygroundRequestLatency.WithLabelValues("eolymp.playground.Playground/CreateRun", s.Code().String()).
			Observe(time.Since(start).Seconds())
	}()

	token, ok := oauth.TokenFromContext(ctx)
	if ok && !token.Has("playground:run:write") {
		err = status.Error(codes.PermissionDenied, "required token scopes are missing: playground:run:write")
		return
	}

	if !i.limiter.Allow(ctx, "eolymp.playground.Playground/CreateRun", 0.16, 5) {
		err = status.Error(codes.ResourceExhausted, "too many requests")
		return
	}

	out, err = i.server.CreateRun(ctx, in)
	return
}

func (i *PlaygroundInterceptor) DescribeRun(ctx context.Context, in *DescribeRunInput) (out *DescribeRunOutput, err error) {
	start := time.Now()
	defer func() {
		s, _ := status.FromError(err)
		if s == nil {
			s = status.New(codes.OK, "OK")
		}

		promPlaygroundRequestLatency.WithLabelValues("eolymp.playground.Playground/DescribeRun", s.Code().String()).
			Observe(time.Since(start).Seconds())
	}()

	token, ok := oauth.TokenFromContext(ctx)
	if ok && !token.Has("playground:run:read") {
		err = status.Error(codes.PermissionDenied, "required token scopes are missing: playground:run:read")
		return
	}

	if !i.limiter.Allow(ctx, "eolymp.playground.Playground/DescribeRun", 2, 10) {
		err = status.Error(codes.ResourceExhausted, "too many requests")
		return
	}

	out, err = i.server.DescribeRun(ctx, in)
	return
}
