// Code generated by protoc-gen-go-server. DO NOT EDIT.
// See https://github.com/eolymp/contracts/tree/main/cmd/protoc-gen-go-server for more details.

package universe

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

// _Universe_HTTPReadRequestBody parses body into proto.Message
func _Universe_HTTPReadRequestBody(r *http.Request, v proto.Message) error {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}

	if err := (protojson.UnmarshalOptions{DiscardUnknown: true}.Unmarshal(data, v)); err != nil {
		return err
	}

	return nil
}

// _Universe_HTTPWriteResponse writes proto.Message to HTTP response
func _Universe_HTTPWriteResponse(w http.ResponseWriter, v proto.Message) {
	data, err := protojson.Marshal(v)
	if err != nil {
		_Universe_HTTPWriteErrorResponse(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	_, _ = w.Write(data)
}

// _Universe_HTTPWriteErrorResponse writes error to HTTP response with error status code
func _Universe_HTTPWriteErrorResponse(w http.ResponseWriter, e error) {
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

// NewUniverseHandler constructs new http.Handler for UniverseServer
func NewUniverseHandler(srv UniverseServer) http.Handler {
	router := mux.NewRouter()
	router.Handle("/eolymp.universe.Universe/CreateSpace", _Universe_CreateSpace(srv)).Methods(http.MethodPost)
	router.Handle("/eolymp.universe.Universe/UpdateSpace", _Universe_UpdateSpace(srv)).Methods(http.MethodPost)
	router.Handle("/eolymp.universe.Universe/DeleteSpace", _Universe_DeleteSpace(srv)).Methods(http.MethodPost)
	router.Handle("/eolymp.universe.Universe/LookupSpace", _Universe_LookupSpace(srv)).Methods(http.MethodPost)
	router.Handle("/eolymp.universe.Universe/DescribeSpace", _Universe_DescribeSpace(srv)).Methods(http.MethodPost)
	router.Handle("/eolymp.universe.Universe/ListSpaces", _Universe_ListSpaces(srv)).Methods(http.MethodPost)
	router.Handle("/eolymp.universe.Universe/GrantPermission", _Universe_GrantPermission(srv)).Methods(http.MethodPost)
	router.Handle("/eolymp.universe.Universe/RevokePermission", _Universe_RevokePermission(srv)).Methods(http.MethodPost)
	router.Handle("/eolymp.universe.Universe/DescribePermission", _Universe_DescribePermission(srv)).Methods(http.MethodPost)
	router.Handle("/eolymp.universe.Universe/IntrospectPermission", _Universe_IntrospectPermission(srv)).Methods(http.MethodPost)
	router.Handle("/eolymp.universe.Universe/ListPermissions", _Universe_ListPermissions(srv)).Methods(http.MethodPost)
	return router
}

func _Universe_CreateSpace(srv UniverseServer) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &CreateSpaceInput{}

		if err := _Universe_HTTPReadRequestBody(r, in); err != nil {
			err = status.New(codes.InvalidArgument, err.Error()).Err()
			_Universe_HTTPWriteErrorResponse(w, err)
			return
		}

		out, err := srv.CreateSpace(r.Context(), in)
		if err != nil {
			_Universe_HTTPWriteErrorResponse(w, err)
			return
		}

		_Universe_HTTPWriteResponse(w, out)
	})
}

func _Universe_UpdateSpace(srv UniverseServer) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &UpdateSpaceInput{}

		if err := _Universe_HTTPReadRequestBody(r, in); err != nil {
			err = status.New(codes.InvalidArgument, err.Error()).Err()
			_Universe_HTTPWriteErrorResponse(w, err)
			return
		}

		out, err := srv.UpdateSpace(r.Context(), in)
		if err != nil {
			_Universe_HTTPWriteErrorResponse(w, err)
			return
		}

		_Universe_HTTPWriteResponse(w, out)
	})
}

func _Universe_DeleteSpace(srv UniverseServer) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &DeleteSpaceInput{}

		if err := _Universe_HTTPReadRequestBody(r, in); err != nil {
			err = status.New(codes.InvalidArgument, err.Error()).Err()
			_Universe_HTTPWriteErrorResponse(w, err)
			return
		}

		out, err := srv.DeleteSpace(r.Context(), in)
		if err != nil {
			_Universe_HTTPWriteErrorResponse(w, err)
			return
		}

		_Universe_HTTPWriteResponse(w, out)
	})
}

func _Universe_LookupSpace(srv UniverseServer) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &LookupSpaceInput{}

		if err := _Universe_HTTPReadRequestBody(r, in); err != nil {
			err = status.New(codes.InvalidArgument, err.Error()).Err()
			_Universe_HTTPWriteErrorResponse(w, err)
			return
		}

		out, err := srv.LookupSpace(r.Context(), in)
		if err != nil {
			_Universe_HTTPWriteErrorResponse(w, err)
			return
		}

		_Universe_HTTPWriteResponse(w, out)
	})
}

func _Universe_DescribeSpace(srv UniverseServer) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &DescribeSpaceInput{}

		if err := _Universe_HTTPReadRequestBody(r, in); err != nil {
			err = status.New(codes.InvalidArgument, err.Error()).Err()
			_Universe_HTTPWriteErrorResponse(w, err)
			return
		}

		out, err := srv.DescribeSpace(r.Context(), in)
		if err != nil {
			_Universe_HTTPWriteErrorResponse(w, err)
			return
		}

		_Universe_HTTPWriteResponse(w, out)
	})
}

func _Universe_ListSpaces(srv UniverseServer) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &ListSpacesInput{}

		if err := _Universe_HTTPReadRequestBody(r, in); err != nil {
			err = status.New(codes.InvalidArgument, err.Error()).Err()
			_Universe_HTTPWriteErrorResponse(w, err)
			return
		}

		out, err := srv.ListSpaces(r.Context(), in)
		if err != nil {
			_Universe_HTTPWriteErrorResponse(w, err)
			return
		}

		_Universe_HTTPWriteResponse(w, out)
	})
}

func _Universe_GrantPermission(srv UniverseServer) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &GrantPermissionInput{}

		if err := _Universe_HTTPReadRequestBody(r, in); err != nil {
			err = status.New(codes.InvalidArgument, err.Error()).Err()
			_Universe_HTTPWriteErrorResponse(w, err)
			return
		}

		out, err := srv.GrantPermission(r.Context(), in)
		if err != nil {
			_Universe_HTTPWriteErrorResponse(w, err)
			return
		}

		_Universe_HTTPWriteResponse(w, out)
	})
}

func _Universe_RevokePermission(srv UniverseServer) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &RevokePermissionInput{}

		if err := _Universe_HTTPReadRequestBody(r, in); err != nil {
			err = status.New(codes.InvalidArgument, err.Error()).Err()
			_Universe_HTTPWriteErrorResponse(w, err)
			return
		}

		out, err := srv.RevokePermission(r.Context(), in)
		if err != nil {
			_Universe_HTTPWriteErrorResponse(w, err)
			return
		}

		_Universe_HTTPWriteResponse(w, out)
	})
}

func _Universe_DescribePermission(srv UniverseServer) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &DescribePermissionInput{}

		if err := _Universe_HTTPReadRequestBody(r, in); err != nil {
			err = status.New(codes.InvalidArgument, err.Error()).Err()
			_Universe_HTTPWriteErrorResponse(w, err)
			return
		}

		out, err := srv.DescribePermission(r.Context(), in)
		if err != nil {
			_Universe_HTTPWriteErrorResponse(w, err)
			return
		}

		_Universe_HTTPWriteResponse(w, out)
	})
}

func _Universe_IntrospectPermission(srv UniverseServer) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &IntrospectPermissionInput{}

		if err := _Universe_HTTPReadRequestBody(r, in); err != nil {
			err = status.New(codes.InvalidArgument, err.Error()).Err()
			_Universe_HTTPWriteErrorResponse(w, err)
			return
		}

		out, err := srv.IntrospectPermission(r.Context(), in)
		if err != nil {
			_Universe_HTTPWriteErrorResponse(w, err)
			return
		}

		_Universe_HTTPWriteResponse(w, out)
	})
}

func _Universe_ListPermissions(srv UniverseServer) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &ListPermissionsInput{}

		if err := _Universe_HTTPReadRequestBody(r, in); err != nil {
			err = status.New(codes.InvalidArgument, err.Error()).Err()
			_Universe_HTTPWriteErrorResponse(w, err)
			return
		}

		out, err := srv.ListPermissions(r.Context(), in)
		if err != nil {
			_Universe_HTTPWriteErrorResponse(w, err)
			return
		}

		_Universe_HTTPWriteResponse(w, out)
	})
}

var promUniverseRequestLatency = promauto.NewHistogramVec(prometheus.HistogramOpts{
	Name:    "universe_request_latency",
	Help:    "Universe request latency",
	Buckets: []float64{0.1, 0.4, 1, 5},
}, []string{"method", "status"})

type _UniverseLimiter interface {
	Allow(context.Context, string, float64, int) bool
}

type UniverseInterceptor struct {
	limiter _UniverseLimiter
	server  UniverseServer
}

// NewUniverseInterceptor constructs additional middleware for a server based on annotations in proto files
func NewUniverseInterceptor(srv UniverseServer, lim _UniverseLimiter) *UniverseInterceptor {
	return &UniverseInterceptor{server: srv, limiter: lim}
}

func (i *UniverseInterceptor) CreateSpace(ctx context.Context, in *CreateSpaceInput) (out *CreateSpaceOutput, err error) {
	start := time.Now()
	defer func() {
		s, _ := status.FromError(err)
		if s == nil {
			s = status.New(codes.OK, "OK")
		}

		promUniverseRequestLatency.WithLabelValues("eolymp.universe.Universe/CreateSpace", s.Code().String()).
			Observe(time.Since(start).Seconds())
	}()

	token, ok := oauth.TokenFromContext(ctx)
	if ok && !token.Has("universe:space:write") {
		err = status.Error(codes.PermissionDenied, "required token scopes are missing: universe:space:write")
		return
	}

	if !i.limiter.Allow(ctx, "eolymp.universe.Universe/CreateSpace", 1, 5) {
		err = status.Error(codes.ResourceExhausted, "too many requests")
		return
	}

	out, err = i.server.CreateSpace(ctx, in)
	return
}

func (i *UniverseInterceptor) UpdateSpace(ctx context.Context, in *UpdateSpaceInput) (out *UpdateSpaceOutput, err error) {
	start := time.Now()
	defer func() {
		s, _ := status.FromError(err)
		if s == nil {
			s = status.New(codes.OK, "OK")
		}

		promUniverseRequestLatency.WithLabelValues("eolymp.universe.Universe/UpdateSpace", s.Code().String()).
			Observe(time.Since(start).Seconds())
	}()

	token, ok := oauth.TokenFromContext(ctx)
	if ok && !token.Has("universe:space:write") {
		err = status.Error(codes.PermissionDenied, "required token scopes are missing: universe:space:write")
		return
	}

	if !i.limiter.Allow(ctx, "eolymp.universe.Universe/UpdateSpace", 1, 5) {
		err = status.Error(codes.ResourceExhausted, "too many requests")
		return
	}

	out, err = i.server.UpdateSpace(ctx, in)
	return
}

func (i *UniverseInterceptor) DeleteSpace(ctx context.Context, in *DeleteSpaceInput) (out *DeleteSpaceOutput, err error) {
	start := time.Now()
	defer func() {
		s, _ := status.FromError(err)
		if s == nil {
			s = status.New(codes.OK, "OK")
		}

		promUniverseRequestLatency.WithLabelValues("eolymp.universe.Universe/DeleteSpace", s.Code().String()).
			Observe(time.Since(start).Seconds())
	}()

	token, ok := oauth.TokenFromContext(ctx)
	if ok && !token.Has("universe:space:write") {
		err = status.Error(codes.PermissionDenied, "required token scopes are missing: universe:space:write")
		return
	}

	if !i.limiter.Allow(ctx, "eolymp.universe.Universe/DeleteSpace", 2, 10) {
		err = status.Error(codes.ResourceExhausted, "too many requests")
		return
	}

	out, err = i.server.DeleteSpace(ctx, in)
	return
}

func (i *UniverseInterceptor) LookupSpace(ctx context.Context, in *LookupSpaceInput) (out *LookupSpaceOutput, err error) {
	start := time.Now()
	defer func() {
		s, _ := status.FromError(err)
		if s == nil {
			s = status.New(codes.OK, "OK")
		}

		promUniverseRequestLatency.WithLabelValues("eolymp.universe.Universe/LookupSpace", s.Code().String()).
			Observe(time.Since(start).Seconds())
	}()

	if !i.limiter.Allow(ctx, "eolymp.universe.Universe/LookupSpace", 10, 100) {
		err = status.Error(codes.ResourceExhausted, "too many requests")
		return
	}

	out, err = i.server.LookupSpace(ctx, in)
	return
}

func (i *UniverseInterceptor) DescribeSpace(ctx context.Context, in *DescribeSpaceInput) (out *DescribeSpaceOutput, err error) {
	start := time.Now()
	defer func() {
		s, _ := status.FromError(err)
		if s == nil {
			s = status.New(codes.OK, "OK")
		}

		promUniverseRequestLatency.WithLabelValues("eolymp.universe.Universe/DescribeSpace", s.Code().String()).
			Observe(time.Since(start).Seconds())
	}()

	if !i.limiter.Allow(ctx, "eolymp.universe.Universe/DescribeSpace", 10, 100) {
		err = status.Error(codes.ResourceExhausted, "too many requests")
		return
	}

	out, err = i.server.DescribeSpace(ctx, in)
	return
}

func (i *UniverseInterceptor) ListSpaces(ctx context.Context, in *ListSpacesInput) (out *ListSpacesOutput, err error) {
	start := time.Now()
	defer func() {
		s, _ := status.FromError(err)
		if s == nil {
			s = status.New(codes.OK, "OK")
		}

		promUniverseRequestLatency.WithLabelValues("eolymp.universe.Universe/ListSpaces", s.Code().String()).
			Observe(time.Since(start).Seconds())
	}()

	token, ok := oauth.TokenFromContext(ctx)
	if ok && !token.Has("universe:space:read") {
		err = status.Error(codes.PermissionDenied, "required token scopes are missing: universe:space:read")
		return
	}

	if !i.limiter.Allow(ctx, "eolymp.universe.Universe/ListSpaces", 5, 20) {
		err = status.Error(codes.ResourceExhausted, "too many requests")
		return
	}

	out, err = i.server.ListSpaces(ctx, in)
	return
}

func (i *UniverseInterceptor) GrantPermission(ctx context.Context, in *GrantPermissionInput) (out *GrantPermissionOutput, err error) {
	start := time.Now()
	defer func() {
		s, _ := status.FromError(err)
		if s == nil {
			s = status.New(codes.OK, "OK")
		}

		promUniverseRequestLatency.WithLabelValues("eolymp.universe.Universe/GrantPermission", s.Code().String()).
			Observe(time.Since(start).Seconds())
	}()

	token, ok := oauth.TokenFromContext(ctx)
	if ok && !token.Has("universe:space:write") {
		err = status.Error(codes.PermissionDenied, "required token scopes are missing: universe:space:write")
		return
	}

	if !i.limiter.Allow(ctx, "eolymp.universe.Universe/GrantPermission", 5, 20) {
		err = status.Error(codes.ResourceExhausted, "too many requests")
		return
	}

	out, err = i.server.GrantPermission(ctx, in)
	return
}

func (i *UniverseInterceptor) RevokePermission(ctx context.Context, in *RevokePermissionInput) (out *RevokePermissionOutput, err error) {
	start := time.Now()
	defer func() {
		s, _ := status.FromError(err)
		if s == nil {
			s = status.New(codes.OK, "OK")
		}

		promUniverseRequestLatency.WithLabelValues("eolymp.universe.Universe/RevokePermission", s.Code().String()).
			Observe(time.Since(start).Seconds())
	}()

	token, ok := oauth.TokenFromContext(ctx)
	if ok && !token.Has("universe:space:write") {
		err = status.Error(codes.PermissionDenied, "required token scopes are missing: universe:space:write")
		return
	}

	if !i.limiter.Allow(ctx, "eolymp.universe.Universe/RevokePermission", 5, 20) {
		err = status.Error(codes.ResourceExhausted, "too many requests")
		return
	}

	out, err = i.server.RevokePermission(ctx, in)
	return
}

func (i *UniverseInterceptor) DescribePermission(ctx context.Context, in *DescribePermissionInput) (out *DescribePermissionOutput, err error) {
	start := time.Now()
	defer func() {
		s, _ := status.FromError(err)
		if s == nil {
			s = status.New(codes.OK, "OK")
		}

		promUniverseRequestLatency.WithLabelValues("eolymp.universe.Universe/DescribePermission", s.Code().String()).
			Observe(time.Since(start).Seconds())
	}()

	token, ok := oauth.TokenFromContext(ctx)
	if ok && !token.Has("universe:space:read") {
		err = status.Error(codes.PermissionDenied, "required token scopes are missing: universe:space:read")
		return
	}

	if !i.limiter.Allow(ctx, "eolymp.universe.Universe/DescribePermission", 5, 20) {
		err = status.Error(codes.ResourceExhausted, "too many requests")
		return
	}

	out, err = i.server.DescribePermission(ctx, in)
	return
}

func (i *UniverseInterceptor) IntrospectPermission(ctx context.Context, in *IntrospectPermissionInput) (out *IntrospectPermissionOutput, err error) {
	start := time.Now()
	defer func() {
		s, _ := status.FromError(err)
		if s == nil {
			s = status.New(codes.OK, "OK")
		}

		promUniverseRequestLatency.WithLabelValues("eolymp.universe.Universe/IntrospectPermission", s.Code().String()).
			Observe(time.Since(start).Seconds())
	}()

	token, ok := oauth.TokenFromContext(ctx)
	if ok && !token.Has("universe:space:read") {
		err = status.Error(codes.PermissionDenied, "required token scopes are missing: universe:space:read")
		return
	}

	if !i.limiter.Allow(ctx, "eolymp.universe.Universe/IntrospectPermission", 5, 20) {
		err = status.Error(codes.ResourceExhausted, "too many requests")
		return
	}

	out, err = i.server.IntrospectPermission(ctx, in)
	return
}

func (i *UniverseInterceptor) ListPermissions(ctx context.Context, in *ListPermissionsInput) (out *ListPermissionsOutput, err error) {
	start := time.Now()
	defer func() {
		s, _ := status.FromError(err)
		if s == nil {
			s = status.New(codes.OK, "OK")
		}

		promUniverseRequestLatency.WithLabelValues("eolymp.universe.Universe/ListPermissions", s.Code().String()).
			Observe(time.Since(start).Seconds())
	}()

	token, ok := oauth.TokenFromContext(ctx)
	if ok && !token.Has("universe:space:read") {
		err = status.Error(codes.PermissionDenied, "required token scopes are missing: universe:space:read")
		return
	}

	if !i.limiter.Allow(ctx, "eolymp.universe.Universe/ListPermissions", 5, 20) {
		err = status.Error(codes.ResourceExhausted, "too many requests")
		return
	}

	out, err = i.server.ListPermissions(ctx, in)
	return
}
