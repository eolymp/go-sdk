// Code generated by protoc-gen-go-server. DO NOT EDIT.
// See https://github.com/eolymp/contracts/tree/main/cmd/protoc-gen-go-server for more details.

package keeper

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

// _Keeper_HTTPReadRequestBody parses body into proto.Message
func _Keeper_HTTPReadRequestBody(r *http.Request, v proto.Message) error {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}

	if len(data) == 0 {
		return nil
	}

	if err := (protojson.UnmarshalOptions{DiscardUnknown: true}.Unmarshal(data, v)); err != nil {
		return err
	}

	return nil
}

// _Keeper_HTTPWriteResponse writes proto.Message to HTTP response
func _Keeper_HTTPWriteResponse(w http.ResponseWriter, v proto.Message) {
	data, err := protojson.Marshal(v)
	if err != nil {
		_Keeper_HTTPWriteErrorResponse(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	_, _ = w.Write(data)
}

// _Keeper_HTTPWriteErrorResponse writes error to HTTP response with error status code
func _Keeper_HTTPWriteErrorResponse(w http.ResponseWriter, e error) {
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

// NewKeeperHandler constructs new http.Handler for KeeperServer
func NewKeeperHandler(srv KeeperServer) http.Handler {
	router := mux.NewRouter()
	router.Handle("/eolymp.keeper.Keeper/CreateObject", _Keeper_CreateObject(srv)).Methods(http.MethodPost)
	router.Handle("/eolymp.keeper.Keeper/DescribeObject", _Keeper_DescribeObject(srv)).Methods(http.MethodPost)
	router.Handle("/eolymp.keeper.Keeper/DownloadObject", _Keeper_DownloadObject(srv)).Methods(http.MethodPost)
	return router
}

// KeeperPrefix defines prefix for routes of this service
const KeeperPrefix = ""

// NewKeeperHandlerHttp constructs new http.Handler for KeeperServer
// This constructor creates http.Handler, the actual implementation might change at any moment
func NewKeeperHandlerHttp(srv KeeperServer, prefix string) http.Handler {
	router := mux.NewRouter()
	return router
}

func _Keeper_CreateObject(srv KeeperServer) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &CreateObjectInput{}

		if err := _Keeper_HTTPReadRequestBody(r, in); err != nil {
			err = status.New(codes.InvalidArgument, err.Error()).Err()
			_Keeper_HTTPWriteErrorResponse(w, err)
			return
		}

		out, err := srv.CreateObject(r.Context(), in)
		if err != nil {
			_Keeper_HTTPWriteErrorResponse(w, err)
			return
		}

		_Keeper_HTTPWriteResponse(w, out)
	})
}

func _Keeper_DescribeObject(srv KeeperServer) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &DescribeObjectInput{}

		if err := _Keeper_HTTPReadRequestBody(r, in); err != nil {
			err = status.New(codes.InvalidArgument, err.Error()).Err()
			_Keeper_HTTPWriteErrorResponse(w, err)
			return
		}

		out, err := srv.DescribeObject(r.Context(), in)
		if err != nil {
			_Keeper_HTTPWriteErrorResponse(w, err)
			return
		}

		_Keeper_HTTPWriteResponse(w, out)
	})
}

func _Keeper_DownloadObject(srv KeeperServer) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &DownloadObjectInput{}

		if err := _Keeper_HTTPReadRequestBody(r, in); err != nil {
			err = status.New(codes.InvalidArgument, err.Error()).Err()
			_Keeper_HTTPWriteErrorResponse(w, err)
			return
		}

		out, err := srv.DownloadObject(r.Context(), in)
		if err != nil {
			_Keeper_HTTPWriteErrorResponse(w, err)
			return
		}

		_Keeper_HTTPWriteResponse(w, out)
	})
}

var promKeeperRequestLatency = promauto.NewHistogramVec(prometheus.HistogramOpts{
	Name:    "keeper_request_latency",
	Help:    "Keeper request latency",
	Buckets: []float64{0.1, 0.4, 1, 5},
}, []string{"method", "status"})

type _KeeperLimiter interface {
	Allow(context.Context, string, float64, int) bool
}

type KeeperInterceptor struct {
	limiter _KeeperLimiter
	server  KeeperServer
}

// NewKeeperInterceptor constructs additional middleware for a server based on annotations in proto files
func NewKeeperInterceptor(srv KeeperServer, lim _KeeperLimiter) *KeeperInterceptor {
	return &KeeperInterceptor{server: srv, limiter: lim}
}

func (i *KeeperInterceptor) CreateObject(ctx context.Context, in *CreateObjectInput) (out *CreateObjectOutput, err error) {
	start := time.Now()
	defer func() {
		s, _ := status.FromError(err)
		if s == nil {
			s = status.New(codes.OK, "OK")
		}

		promKeeperRequestLatency.WithLabelValues("eolymp.keeper.Keeper/CreateObject", s.Code().String()).
			Observe(time.Since(start).Seconds())
	}()

	token, ok := oauth.TokenFromContext(ctx)
	if ok && !token.Has("keeper:object:write") {
		err = status.Error(codes.PermissionDenied, "required token scopes are missing: keeper:object:write")
		return
	}

	if !i.limiter.Allow(ctx, "eolymp.keeper.Keeper/CreateObject", 2, 500) {
		err = status.Error(codes.ResourceExhausted, "too many requests")
		return
	}

	out, err = i.server.CreateObject(ctx, in)
	return
}

func (i *KeeperInterceptor) DescribeObject(ctx context.Context, in *DescribeObjectInput) (out *DescribeObjectOutput, err error) {
	start := time.Now()
	defer func() {
		s, _ := status.FromError(err)
		if s == nil {
			s = status.New(codes.OK, "OK")
		}

		promKeeperRequestLatency.WithLabelValues("eolymp.keeper.Keeper/DescribeObject", s.Code().String()).
			Observe(time.Since(start).Seconds())
	}()

	token, ok := oauth.TokenFromContext(ctx)
	if ok && !token.Has("keeper:object:read") {
		err = status.Error(codes.PermissionDenied, "required token scopes are missing: keeper:object:read")
		return
	}

	if !i.limiter.Allow(ctx, "eolymp.keeper.Keeper/DescribeObject", 50, 500) {
		err = status.Error(codes.ResourceExhausted, "too many requests")
		return
	}

	out, err = i.server.DescribeObject(ctx, in)
	return
}

func (i *KeeperInterceptor) DownloadObject(ctx context.Context, in *DownloadObjectInput) (out *DownloadObjectOutput, err error) {
	start := time.Now()
	defer func() {
		s, _ := status.FromError(err)
		if s == nil {
			s = status.New(codes.OK, "OK")
		}

		promKeeperRequestLatency.WithLabelValues("eolymp.keeper.Keeper/DownloadObject", s.Code().String()).
			Observe(time.Since(start).Seconds())
	}()

	token, ok := oauth.TokenFromContext(ctx)
	if ok && !token.Has("keeper:object:read") {
		err = status.Error(codes.PermissionDenied, "required token scopes are missing: keeper:object:read")
		return
	}

	if !i.limiter.Allow(ctx, "eolymp.keeper.Keeper/DownloadObject", 50, 500) {
		err = status.Error(codes.ResourceExhausted, "too many requests")
		return
	}

	out, err = i.server.DownloadObject(ctx, in)
	return
}
