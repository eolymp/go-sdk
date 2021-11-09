// Code generated by protoc-gen-go-server. DO NOT EDIT.
// See https://github.com/eolymp/contracts/tree/main/cmd/protoc-gen-go-server for more details.

package community

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

// _Community_HTTPReadRequestBody parses body into proto.Message
func _Community_HTTPReadRequestBody(r *http.Request, v proto.Message) error {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}

	if err := (protojson.UnmarshalOptions{DiscardUnknown: true}.Unmarshal(data, v)); err != nil {
		return err
	}

	return nil
}

// _Community_HTTPWriteResponse writes proto.Message to HTTP response
func _Community_HTTPWriteResponse(w http.ResponseWriter, v proto.Message) {
	data, err := protojson.Marshal(v)
	if err != nil {
		_Community_HTTPWriteErrorResponse(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	_, _ = w.Write(data)
}

// _Community_HTTPWriteErrorResponse writes error to HTTP response with error status code
func _Community_HTTPWriteErrorResponse(w http.ResponseWriter, e error) {
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

// NewCommunityHandler constructs new http.Handler for CommunityServer
func NewCommunityHandler(srv CommunityServer) http.Handler {
	router := mux.NewRouter()
	router.Handle("/eolymp.community.Community/AddMember", _Community_AddMember(srv)).Methods(http.MethodPost)
	router.Handle("/eolymp.community.Community/UpdateMember", _Community_UpdateMember(srv)).Methods(http.MethodPost)
	router.Handle("/eolymp.community.Community/RemoveMember", _Community_RemoveMember(srv)).Methods(http.MethodPost)
	router.Handle("/eolymp.community.Community/DescribeMember", _Community_DescribeMember(srv)).Methods(http.MethodPost)
	router.Handle("/eolymp.community.Community/IntrospectMember", _Community_IntrospectMember(srv)).Methods(http.MethodPost)
	router.Handle("/eolymp.community.Community/ListMembers", _Community_ListMembers(srv)).Methods(http.MethodPost)
	router.Handle("/eolymp.community.Community/AddAttribute", _Community_AddAttribute(srv)).Methods(http.MethodPost)
	router.Handle("/eolymp.community.Community/UpdateAttribute", _Community_UpdateAttribute(srv)).Methods(http.MethodPost)
	router.Handle("/eolymp.community.Community/RemoveAttribute", _Community_RemoveAttribute(srv)).Methods(http.MethodPost)
	router.Handle("/eolymp.community.Community/DescribeAttribute", _Community_DescribeAttribute(srv)).Methods(http.MethodPost)
	router.Handle("/eolymp.community.Community/ListAttributes", _Community_ListAttributes(srv)).Methods(http.MethodPost)
	return router
}

func _Community_AddMember(srv CommunityServer) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &AddMemberInput{}

		if err := _Community_HTTPReadRequestBody(r, in); err != nil {
			err = status.New(codes.InvalidArgument, err.Error()).Err()
			_Community_HTTPWriteErrorResponse(w, err)
			return
		}

		out, err := srv.AddMember(r.Context(), in)
		if err != nil {
			_Community_HTTPWriteErrorResponse(w, err)
			return
		}

		_Community_HTTPWriteResponse(w, out)
	})
}

func _Community_UpdateMember(srv CommunityServer) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &UpdateMemberInput{}

		if err := _Community_HTTPReadRequestBody(r, in); err != nil {
			err = status.New(codes.InvalidArgument, err.Error()).Err()
			_Community_HTTPWriteErrorResponse(w, err)
			return
		}

		out, err := srv.UpdateMember(r.Context(), in)
		if err != nil {
			_Community_HTTPWriteErrorResponse(w, err)
			return
		}

		_Community_HTTPWriteResponse(w, out)
	})
}

func _Community_RemoveMember(srv CommunityServer) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &RemoveMemberInput{}

		if err := _Community_HTTPReadRequestBody(r, in); err != nil {
			err = status.New(codes.InvalidArgument, err.Error()).Err()
			_Community_HTTPWriteErrorResponse(w, err)
			return
		}

		out, err := srv.RemoveMember(r.Context(), in)
		if err != nil {
			_Community_HTTPWriteErrorResponse(w, err)
			return
		}

		_Community_HTTPWriteResponse(w, out)
	})
}

func _Community_DescribeMember(srv CommunityServer) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &DescribeMemberInput{}

		if err := _Community_HTTPReadRequestBody(r, in); err != nil {
			err = status.New(codes.InvalidArgument, err.Error()).Err()
			_Community_HTTPWriteErrorResponse(w, err)
			return
		}

		out, err := srv.DescribeMember(r.Context(), in)
		if err != nil {
			_Community_HTTPWriteErrorResponse(w, err)
			return
		}

		_Community_HTTPWriteResponse(w, out)
	})
}

func _Community_IntrospectMember(srv CommunityServer) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &IntrospectMemberInput{}

		if err := _Community_HTTPReadRequestBody(r, in); err != nil {
			err = status.New(codes.InvalidArgument, err.Error()).Err()
			_Community_HTTPWriteErrorResponse(w, err)
			return
		}

		out, err := srv.IntrospectMember(r.Context(), in)
		if err != nil {
			_Community_HTTPWriteErrorResponse(w, err)
			return
		}

		_Community_HTTPWriteResponse(w, out)
	})
}

func _Community_ListMembers(srv CommunityServer) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &ListMembersInput{}

		if err := _Community_HTTPReadRequestBody(r, in); err != nil {
			err = status.New(codes.InvalidArgument, err.Error()).Err()
			_Community_HTTPWriteErrorResponse(w, err)
			return
		}

		out, err := srv.ListMembers(r.Context(), in)
		if err != nil {
			_Community_HTTPWriteErrorResponse(w, err)
			return
		}

		_Community_HTTPWriteResponse(w, out)
	})
}

func _Community_AddAttribute(srv CommunityServer) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &AddAttributeInput{}

		if err := _Community_HTTPReadRequestBody(r, in); err != nil {
			err = status.New(codes.InvalidArgument, err.Error()).Err()
			_Community_HTTPWriteErrorResponse(w, err)
			return
		}

		out, err := srv.AddAttribute(r.Context(), in)
		if err != nil {
			_Community_HTTPWriteErrorResponse(w, err)
			return
		}

		_Community_HTTPWriteResponse(w, out)
	})
}

func _Community_UpdateAttribute(srv CommunityServer) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &UpdateAttributeInput{}

		if err := _Community_HTTPReadRequestBody(r, in); err != nil {
			err = status.New(codes.InvalidArgument, err.Error()).Err()
			_Community_HTTPWriteErrorResponse(w, err)
			return
		}

		out, err := srv.UpdateAttribute(r.Context(), in)
		if err != nil {
			_Community_HTTPWriteErrorResponse(w, err)
			return
		}

		_Community_HTTPWriteResponse(w, out)
	})
}

func _Community_RemoveAttribute(srv CommunityServer) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &RemoveAttributeInput{}

		if err := _Community_HTTPReadRequestBody(r, in); err != nil {
			err = status.New(codes.InvalidArgument, err.Error()).Err()
			_Community_HTTPWriteErrorResponse(w, err)
			return
		}

		out, err := srv.RemoveAttribute(r.Context(), in)
		if err != nil {
			_Community_HTTPWriteErrorResponse(w, err)
			return
		}

		_Community_HTTPWriteResponse(w, out)
	})
}

func _Community_DescribeAttribute(srv CommunityServer) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &DescribeAttributeInput{}

		if err := _Community_HTTPReadRequestBody(r, in); err != nil {
			err = status.New(codes.InvalidArgument, err.Error()).Err()
			_Community_HTTPWriteErrorResponse(w, err)
			return
		}

		out, err := srv.DescribeAttribute(r.Context(), in)
		if err != nil {
			_Community_HTTPWriteErrorResponse(w, err)
			return
		}

		_Community_HTTPWriteResponse(w, out)
	})
}

func _Community_ListAttributes(srv CommunityServer) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &ListAttributesInput{}

		if err := _Community_HTTPReadRequestBody(r, in); err != nil {
			err = status.New(codes.InvalidArgument, err.Error()).Err()
			_Community_HTTPWriteErrorResponse(w, err)
			return
		}

		out, err := srv.ListAttributes(r.Context(), in)
		if err != nil {
			_Community_HTTPWriteErrorResponse(w, err)
			return
		}

		_Community_HTTPWriteResponse(w, out)
	})
}

var promCommunityRequestLatency = promauto.NewHistogramVec(prometheus.HistogramOpts{
	Name:    "community_request_latency",
	Help:    "Community request latency",
	Buckets: []float64{0.1, 0.4, 1, 5},
}, []string{"method", "status"})

type _CommunityLimiter interface {
	Allow(context.Context, string, float64, int) bool
}

type CommunityInterceptor struct {
	limiter _CommunityLimiter
	server  CommunityServer
}

// NewCommunityInterceptor constructs additional middleware for a server based on annotations in proto files
func NewCommunityInterceptor(srv CommunityServer, lim _CommunityLimiter) *CommunityInterceptor {
	return &CommunityInterceptor{server: srv, limiter: lim}
}

func (i *CommunityInterceptor) AddMember(ctx context.Context, in *AddMemberInput) (out *AddMemberOutput, err error) {
	start := time.Now()
	defer func() {
		s, _ := status.FromError(err)
		if s == nil {
			s = status.New(codes.OK, "OK")
		}

		promCommunityRequestLatency.WithLabelValues("eolymp.community.Community/AddMember", s.Code().String()).
			Observe(time.Since(start).Seconds())
	}()

	token, ok := oauth.TokenFromContext(ctx)
	if ok && !token.Has("community:member:write") {
		err = status.Error(codes.PermissionDenied, "required token scopes are missing: community:member:write")
		return
	}

	if !i.limiter.Allow(ctx, "eolymp.community.Community/AddMember", 5, 20) {
		err = status.Error(codes.ResourceExhausted, "too many requests")
		return
	}

	out, err = i.server.AddMember(ctx, in)
	return
}

func (i *CommunityInterceptor) UpdateMember(ctx context.Context, in *UpdateMemberInput) (out *UpdateMemberOutput, err error) {
	start := time.Now()
	defer func() {
		s, _ := status.FromError(err)
		if s == nil {
			s = status.New(codes.OK, "OK")
		}

		promCommunityRequestLatency.WithLabelValues("eolymp.community.Community/UpdateMember", s.Code().String()).
			Observe(time.Since(start).Seconds())
	}()

	token, ok := oauth.TokenFromContext(ctx)
	if ok && !token.Has("community:member:write") {
		err = status.Error(codes.PermissionDenied, "required token scopes are missing: community:member:write")
		return
	}

	if !i.limiter.Allow(ctx, "eolymp.community.Community/UpdateMember", 5, 20) {
		err = status.Error(codes.ResourceExhausted, "too many requests")
		return
	}

	out, err = i.server.UpdateMember(ctx, in)
	return
}

func (i *CommunityInterceptor) RemoveMember(ctx context.Context, in *RemoveMemberInput) (out *RemoveMemberOutput, err error) {
	start := time.Now()
	defer func() {
		s, _ := status.FromError(err)
		if s == nil {
			s = status.New(codes.OK, "OK")
		}

		promCommunityRequestLatency.WithLabelValues("eolymp.community.Community/RemoveMember", s.Code().String()).
			Observe(time.Since(start).Seconds())
	}()

	token, ok := oauth.TokenFromContext(ctx)
	if ok && !token.Has("community:member:write") {
		err = status.Error(codes.PermissionDenied, "required token scopes are missing: community:member:write")
		return
	}

	if !i.limiter.Allow(ctx, "eolymp.community.Community/RemoveMember", 5, 20) {
		err = status.Error(codes.ResourceExhausted, "too many requests")
		return
	}

	out, err = i.server.RemoveMember(ctx, in)
	return
}

func (i *CommunityInterceptor) DescribeMember(ctx context.Context, in *DescribeMemberInput) (out *DescribeMemberOutput, err error) {
	start := time.Now()
	defer func() {
		s, _ := status.FromError(err)
		if s == nil {
			s = status.New(codes.OK, "OK")
		}

		promCommunityRequestLatency.WithLabelValues("eolymp.community.Community/DescribeMember", s.Code().String()).
			Observe(time.Since(start).Seconds())
	}()

	token, ok := oauth.TokenFromContext(ctx)
	if ok && !token.Has("community:member:read") {
		err = status.Error(codes.PermissionDenied, "required token scopes are missing: community:member:read")
		return
	}

	if !i.limiter.Allow(ctx, "eolymp.community.Community/DescribeMember", 5, 20) {
		err = status.Error(codes.ResourceExhausted, "too many requests")
		return
	}

	out, err = i.server.DescribeMember(ctx, in)
	return
}

func (i *CommunityInterceptor) IntrospectMember(ctx context.Context, in *IntrospectMemberInput) (out *IntrospectMemberOutput, err error) {
	start := time.Now()
	defer func() {
		s, _ := status.FromError(err)
		if s == nil {
			s = status.New(codes.OK, "OK")
		}

		promCommunityRequestLatency.WithLabelValues("eolymp.community.Community/IntrospectMember", s.Code().String()).
			Observe(time.Since(start).Seconds())
	}()

	token, ok := oauth.TokenFromContext(ctx)
	if ok && !token.Has("community:member:read") {
		err = status.Error(codes.PermissionDenied, "required token scopes are missing: community:member:read")
		return
	}

	if !i.limiter.Allow(ctx, "eolymp.community.Community/IntrospectMember", 5, 20) {
		err = status.Error(codes.ResourceExhausted, "too many requests")
		return
	}

	out, err = i.server.IntrospectMember(ctx, in)
	return
}

func (i *CommunityInterceptor) ListMembers(ctx context.Context, in *ListMembersInput) (out *ListMembersOutput, err error) {
	start := time.Now()
	defer func() {
		s, _ := status.FromError(err)
		if s == nil {
			s = status.New(codes.OK, "OK")
		}

		promCommunityRequestLatency.WithLabelValues("eolymp.community.Community/ListMembers", s.Code().String()).
			Observe(time.Since(start).Seconds())
	}()

	token, ok := oauth.TokenFromContext(ctx)
	if ok && !token.Has("community:member:read") {
		err = status.Error(codes.PermissionDenied, "required token scopes are missing: community:member:read")
		return
	}

	if !i.limiter.Allow(ctx, "eolymp.community.Community/ListMembers", 5, 20) {
		err = status.Error(codes.ResourceExhausted, "too many requests")
		return
	}

	out, err = i.server.ListMembers(ctx, in)
	return
}

func (i *CommunityInterceptor) AddAttribute(ctx context.Context, in *AddAttributeInput) (out *AddAttributeOutput, err error) {
	start := time.Now()
	defer func() {
		s, _ := status.FromError(err)
		if s == nil {
			s = status.New(codes.OK, "OK")
		}

		promCommunityRequestLatency.WithLabelValues("eolymp.community.Community/AddAttribute", s.Code().String()).
			Observe(time.Since(start).Seconds())
	}()

	token, ok := oauth.TokenFromContext(ctx)
	if ok && !token.Has("community:member:write") {
		err = status.Error(codes.PermissionDenied, "required token scopes are missing: community:member:write")
		return
	}

	if !i.limiter.Allow(ctx, "eolymp.community.Community/AddAttribute", 5, 20) {
		err = status.Error(codes.ResourceExhausted, "too many requests")
		return
	}

	out, err = i.server.AddAttribute(ctx, in)
	return
}

func (i *CommunityInterceptor) UpdateAttribute(ctx context.Context, in *UpdateAttributeInput) (out *UpdateAttributeOutput, err error) {
	start := time.Now()
	defer func() {
		s, _ := status.FromError(err)
		if s == nil {
			s = status.New(codes.OK, "OK")
		}

		promCommunityRequestLatency.WithLabelValues("eolymp.community.Community/UpdateAttribute", s.Code().String()).
			Observe(time.Since(start).Seconds())
	}()

	token, ok := oauth.TokenFromContext(ctx)
	if ok && !token.Has("community:member:write") {
		err = status.Error(codes.PermissionDenied, "required token scopes are missing: community:member:write")
		return
	}

	if !i.limiter.Allow(ctx, "eolymp.community.Community/UpdateAttribute", 5, 20) {
		err = status.Error(codes.ResourceExhausted, "too many requests")
		return
	}

	out, err = i.server.UpdateAttribute(ctx, in)
	return
}

func (i *CommunityInterceptor) RemoveAttribute(ctx context.Context, in *RemoveAttributeInput) (out *RemoveAttributeOutput, err error) {
	start := time.Now()
	defer func() {
		s, _ := status.FromError(err)
		if s == nil {
			s = status.New(codes.OK, "OK")
		}

		promCommunityRequestLatency.WithLabelValues("eolymp.community.Community/RemoveAttribute", s.Code().String()).
			Observe(time.Since(start).Seconds())
	}()

	token, ok := oauth.TokenFromContext(ctx)
	if ok && !token.Has("community:member:write") {
		err = status.Error(codes.PermissionDenied, "required token scopes are missing: community:member:write")
		return
	}

	if !i.limiter.Allow(ctx, "eolymp.community.Community/RemoveAttribute", 5, 20) {
		err = status.Error(codes.ResourceExhausted, "too many requests")
		return
	}

	out, err = i.server.RemoveAttribute(ctx, in)
	return
}

func (i *CommunityInterceptor) DescribeAttribute(ctx context.Context, in *DescribeAttributeInput) (out *DescribeAttributeOutput, err error) {
	start := time.Now()
	defer func() {
		s, _ := status.FromError(err)
		if s == nil {
			s = status.New(codes.OK, "OK")
		}

		promCommunityRequestLatency.WithLabelValues("eolymp.community.Community/DescribeAttribute", s.Code().String()).
			Observe(time.Since(start).Seconds())
	}()

	token, ok := oauth.TokenFromContext(ctx)
	if ok && !token.Has("community:member:read") {
		err = status.Error(codes.PermissionDenied, "required token scopes are missing: community:member:read")
		return
	}

	if !i.limiter.Allow(ctx, "eolymp.community.Community/DescribeAttribute", 5, 20) {
		err = status.Error(codes.ResourceExhausted, "too many requests")
		return
	}

	out, err = i.server.DescribeAttribute(ctx, in)
	return
}

func (i *CommunityInterceptor) ListAttributes(ctx context.Context, in *ListAttributesInput) (out *ListAttributesOutput, err error) {
	start := time.Now()
	defer func() {
		s, _ := status.FromError(err)
		if s == nil {
			s = status.New(codes.OK, "OK")
		}

		promCommunityRequestLatency.WithLabelValues("eolymp.community.Community/ListAttributes", s.Code().String()).
			Observe(time.Since(start).Seconds())
	}()

	token, ok := oauth.TokenFromContext(ctx)
	if ok && !token.Has("community:member:read") {
		err = status.Error(codes.PermissionDenied, "required token scopes are missing: community:member:read")
		return
	}

	if !i.limiter.Allow(ctx, "eolymp.community.Community/ListAttributes", 5, 20) {
		err = status.Error(codes.ResourceExhausted, "too many requests")
		return
	}

	out, err = i.server.ListAttributes(ctx, in)
	return
}
